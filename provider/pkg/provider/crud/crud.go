package crud

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/convert"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource/plugin"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/rpcutil/rpcerror"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
	"google.golang.org/grpc/codes"
)

// AzureRESTConverter is an interface for preparing Azure inputs from Pulumi data and for converting from Azure outputs to Pulumi SDK shape.
// It operates in the context of a specific kind of Azure resource of type resources.AzureAPIResource.
type AzureRESTConverter interface {
	// PrepareAzureRESTIdAndQuery prepares the ID and query parameters for an Azure REST API request.
	PrepareAzureRESTIdAndQuery(inputs resource.PropertyMap) (string, map[string]any, error)

	// PrepareAzureRESTBody prepares the body of an Azure REST API request.
	// It returns the body as a map of strings to any, which can be marshaled to JSON.
	// An error might be returned instead of or as well as the body. If the body is not nil, the error can be treated as a warning.
	PrepareAzureRESTBody(id string, inputs resource.PropertyMap) (map[string]any, error)

	// ResponseBodyToSdkOutputs converts an Azure REST API response to Pulumi SDK outputs.
	ResponseBodyToSdkOutputs(response map[string]any) map[string]any
	ResponseToSdkInputs(pathValues map[string]string, responseBody map[string]any) map[string]any
	SdkInputsToRequestBody(values map[string]any, id string) (map[string]any, error)
}

// ResourceCrudOperations is an interface for performing CRUD operations on Azure resources of a certain kind.
// See AzureRESTConverter for creating proper inputs and converting outputs.
// It operates in the context of a specific kind of Azure resource of type resources.AzureAPIResource.
type ResourceCrudOperations interface {
	CanCreate(ctx context.Context, id string) error
	CreateOrUpdate(ctx context.Context, id string, bodyParams, queryParams map[string]any) (map[string]any, bool, error)
	Read(ctx context.Context, id string) (map[string]any, error)
	HandleErrorWithCheckpoint(ctx context.Context, err error, id string, inputs resource.PropertyMap, properties *structpb.Struct) error
}

// SubresourceMaintainer is an interface for handling sub-resource properties.
// It operates in the context of a specific kind of Azure resource of type resources.AzureAPIResource.
type SubresourceMaintainer interface {
	MaintainSubResourcePropertiesIfNotSet(ctx context.Context, id string, bodyParams map[string]any) error
	// Properties pointing to sub-resources that can be maintained as separate resources might not be
	// present in the inputs because the user wants to manage them as standalone resources. However,
	// such a property might be required by Azure even if it's not annotated as such in the spec, e.g.,
	// Key Vault's accessPolicies. Therefore, we set these properties to their default value here,
	// an empty array. For more details, see section "Sub-resources" in CONTRIBUTING.md.
	//
	// During create, no sub-resources can exist yet so there's no danger of overwriting existing values.
	//
	// The `input` param is used to determine the unset sub-resource properties. They are then reset in
	// the `output` parameter which is modified in-place.
	//
	// Implementation note: we should make it possible to write custom resources that call code from
	// the default implementation as needed. This would allow us to cleanly implement special logic
	// like for Key Vault into custom resources without duplicating much code. In the Key Vault case,
	// the custom Read() would look like
	//
	//	provider.azureCanCreate(ctx, id, &res)
	//	setUnsetSubresourcePropertiesToDefaults(res, bodyParams) // custom
	//	k.azureCreateOrUpdate
	//	...
	SetUnsetSubresourcePropertiesToDefaults(input, output map[string]any, outputIsInApiShape bool)
}

// ResourceCrudClient is a client for performing CRUD operations on Azure resources.
// It encapsulates both common logic and instances of helpers such as azure.AzureClient and convert.SdkShapeConverter.
// It operates in the context of a specific kind of Azure resource of type resources.AzureAPIResource.
type ResourceCrudClient interface {
	ResourceCrudOperations
	AzureRESTConverter
	SubresourceMaintainer
	ApiVersion() string
}

type resourceCrudClient struct {
	azureClient    azure.AzureClient
	lookupType     resources.TypeLookupFunc
	converter      *convert.SdkShapeConverter
	subscriptionID string

	res *resources.AzureAPIResource
}

type ResourceCrudClientFactory func(
	res *resources.AzureAPIResource,
) ResourceCrudClient

func NewResourceCrudClient(
	azureClient azure.AzureClient,
	lookupType resources.TypeLookupFunc,
	converter *convert.SdkShapeConverter,
	subscriptionID string,
	res *resources.AzureAPIResource,
) ResourceCrudClient {
	return &resourceCrudClient{
		azureClient:    azureClient,
		lookupType:     lookupType,
		converter:      converter,
		subscriptionID: subscriptionID,
		res:            res,
	}
}

func (r *resourceCrudClient) ApiVersion() string {
	return r.res.APIVersion
}

func (r *resourceCrudClient) PrepareAzureRESTIdAndQuery(inputs resource.PropertyMap) (string, map[string]any, error) {
	return PrepareAzureRESTIdAndQuery(r.res.Path, r.res.PutParameters, inputs.Mappable(), map[string]any{
		"subscriptionId": r.subscriptionID,
		"api-version":    r.res.APIVersion,
	})
}

func (r *resourceCrudClient) PrepareAzureRESTBody(id string, inputs resource.PropertyMap) (map[string]any, error) {
	return PrepareAzureRESTBody(id, r.res.PutParameters, r.res.RequiredContainers, inputs.Mappable(), r.converter)
}

func PrepareAzureRESTIdAndQuery(path string, parameters []resources.AzureAPIParameter, methodInputs, clientInputs map[string]any) (string, map[string]any, error) {
	params := map[string]map[string]interface{}{
		"query": {
			"api-version": clientInputs["api-version"],
		},
		"path": {},
	}

	// Build maps of path and query parameters.
	for _, param := range parameters {
		if param.Location == "body" {
			continue
		}
		var val interface{}
		var has bool
		sdkName := param.Name
		if param.Value.SdkName != "" {
			sdkName = param.Value.SdkName
		}
		// Look in both `method` and `client` inputs with `method` first
		// Navigate into each level of the path to find the value.

		val, has = findInContainer(methodInputs, strings.Split(sdkName, "."))
		if !has {
			val, has = clientInputs[sdkName]
		}
		if has {
			params[param.Location][param.Name] = val
		}
	}

	// Calculate resource ID based on path parameter values.
	id := path
	for key, value := range params["path"] {
		strVal, isString := value.(string)
		if !isString {
			return "", nil, errors.Errorf("expected string value for path parameter '%s', got %T", key, value)
		}
		encodedVal := strings.Replace(url.QueryEscape(strVal), "+", "%20", -1)
		id = strings.Replace(id, "{"+key+"}", encodedVal, -1)
	}

	return id, params["query"], nil
}

func findInContainer(container map[string]any, path []string) (interface{}, bool) {
	currentContainer := container
	for i, pathPart := range path {
		if currentContainer == nil {
			return nil, false
		}
		inner, has := currentContainer[pathPart]
		if !has {
			return nil, false
		}
		if i == len(path)-1 {
			return inner, true
		}
		currentContainer, _ = inner.(map[string]any)
	}
	return nil, false
}

func PrepareAzureRESTBody(id string, parameters []resources.AzureAPIParameter, requiredContainers [][]string,
	methodInputs map[string]any, converter *convert.SdkShapeConverter) (map[string]any, error) {
	// Build the body JSON.
	var body map[string]interface{}
	var err error
	for _, param := range parameters {
		if param.Location != "body" {
			continue
		}
		body, err = converter.SdkInputsToRequestBody(param.Body.Properties, methodInputs, id)
		break
	}

	// Ensure all required containers are created.
	for _, containers := range requiredContainers {
		currentContainer := body
		for _, containerName := range containers {
			innerContainer, ok := currentContainer[containerName]
			if !ok {
				innerContainer = map[string]interface{}{}
				currentContainer[containerName] = innerContainer
			}
			currentContainer, ok = innerContainer.(map[string]interface{})
			if !ok {
				break
			}
		}
	}

	return body, err
}

// partialError creates an error for resources that did not complete an operation in progress.
// The last known state of the object is included in the error so that it can be checkpointed.
func partialError(id string, err error, state *structpb.Struct, inputs *structpb.Struct) error {
	detail := rpc.ErrorResourceInitFailed{
		Id:         id,
		Properties: state,
		Reasons:    []string{err.Error()},
		Inputs:     inputs,
	}
	return rpcerror.WithDetails(rpcerror.New(codes.Unknown, err.Error()), &detail)
}

func (r *resourceCrudClient) HandleErrorWithCheckpoint(ctx context.Context, err error, id string, inputs resource.PropertyMap, properties *structpb.Struct) error {
	// Resource was partially updated but the operation failed to complete.
	// Try reading its state by ID and return a partial error if succeeded.
	checkpoint, getErr := r.currentResourceStateCheckpoint(ctx, id, inputs)
	if getErr != nil {
		// If reading the state failed, combine the errors but still return the partial state.
		err = azure.AzureError(errors.Wrapf(err, "resource created but read failed %s", getErr))
	}
	return partialError(id, err, checkpoint, properties)
}

// currentResourceStateCheckpoint reads the resource state by ID, converts it to outputs map, and
// produces a checkpoint with these outputs and given inputs.
func (r *resourceCrudClient) currentResourceStateCheckpoint(ctx context.Context, id string, inputs resource.PropertyMap) (*structpb.Struct, error) {
	getResp, getErr := r.azureClient.Get(ctx, id, r.res.APIVersion, r.res.ReadQueryParams)
	if getErr != nil {
		return nil, getErr
	}
	outputs := r.converter.ResponseBodyToSdkOutputs(r.res.Response, getResp.(map[string]any))
	obj := checkpointObject(inputs, outputs)
	return plugin.MarshalProperties(
		obj,
		plugin.MarshalOptions{
			Label:        "currentResourceStateCheckpoint.checkpoint",
			KeepSecrets:  true,
			KeepUnknowns: true,
			SkipNulls:    true,
		},
	)
}

// checkpointObject puts inputs in the `__inputs` field of the state.
func checkpointObject(inputs resource.PropertyMap, outputs map[string]interface{}) resource.PropertyMap {
	object := resource.NewPropertyMapFromMap(outputs)
	object["__inputs"] = resource.MakeSecret(resource.NewObjectProperty(inputs))
	return object
}

func (r *resourceCrudClient) MaintainSubResourcePropertiesIfNotSet(ctx context.Context, id string, bodyParams map[string]interface{}) error {
	// Identify the properties we need to read
	missingProperties := findUnsetPropertiesToMaintain(r.res, bodyParams, true /* returnApiShapePaths */, r.lookupType)

	if len(missingProperties) == 0 {
		// Everything's already specified explicitly by the user, no need to do read.
		return nil
	}

	// Read the current resource state.
	state, err := r.azureClient.Get(ctx, id+r.res.ReadPath, r.res.APIVersion, r.res.ReadQueryParams)
	if err != nil {
		return fmt.Errorf("reading cloud state: %w", err)
	}

	writtenProperties := writePropertiesToBody(missingProperties, bodyParams, state.(map[string]any))
	for writtenProperty, writtenValue := range writtenProperties {
		logging.V(9).Infof("Maintaining remote value for property: %s.%s = %v", id, writtenProperty, writtenValue)
	}

	return nil
}

// SetUnsetSubresourcePropertiesToDefaults is the standard implementation of SubresourceMaintainer.SetUnsetSubresourcePropertiesToDefaults.
func (r *resourceCrudClient) SetUnsetSubresourcePropertiesToDefaults(input, output map[string]interface{}, outputIsInApiShape bool) {
	unset := findUnsetPropertiesToMaintain(r.res, input, outputIsInApiShape, r.lookupType)

	for _, p := range unset {
		cur := output
		for _, pathEl := range p.path[:len(p.path)-1] {
			curObj, ok := cur[pathEl]
			if !ok {
				newContainer := map[string]any{}
				cur[pathEl] = newContainer
				cur = newContainer
				continue
			}
			cur, ok = curObj.(map[string]any)
			if !ok {
				break
			}
		}

		cur[p.path[len(p.path)-1]] = []any{}
	}
}

func findUnsetPropertiesToMaintain(res *resources.AzureAPIResource, bodyParams map[string]interface{}, returnApiShapePaths bool, lookupType resources.TypeLookupFunc) []propertyPath {
	missingProperties := []propertyPath{}
	for _, path := range res.PathsToSubResourcePropertiesToMaintain(returnApiShapePaths, lookupType) {
		curBody := bodyParams
		for i, pathEl := range path {
			v, ok := curBody[pathEl]
			if !ok {
				missingProperties = append(missingProperties, propertyPath{
					path:         path,
					propertyName: pathEl,
				})
				break
			}

			// At the end of the path we don't need to go deeper via references and map lookups.
			if i == len(path)-1 {
				break
			}

			curBody, ok = v.(map[string]interface{})
			if !ok {
				missingProperties = append(missingProperties, propertyPath{
					path:         path,
					propertyName: pathEl,
				})
				break
			}
		}
	}

	return missingProperties
}

func (r *resourceCrudClient) ResponseBodyToSdkOutputs(response map[string]any) map[string]any {
	return r.converter.ResponseBodyToSdkOutputs(r.res.Response, response)
}

func (r *resourceCrudClient) ResponseToSdkInputs(pathValues map[string]string, responseBody map[string]any) map[string]any {
	return r.converter.ResponseToSdkInputs(r.res.PutParameters, pathValues, responseBody)
}

func (r *resourceCrudClient) SdkInputsToRequestBody(properties map[string]any, id string) (map[string]any, error) {
	if bodyParam, ok := r.res.BodyParameter(); ok {
		return r.converter.SdkInputsToRequestBody(bodyParam.Body.Properties, properties, id)
	}
	return nil, nil
}

func (r *resourceCrudClient) CanCreate(ctx context.Context, id string) error {
	return r.azureClient.CanCreate(ctx, id, r.res.ReadPath, r.res.APIVersion, r.res.ReadMethod, r.res.Singleton, r.res.DefaultBody != nil, func(outputs map[string]any) bool {
		return r.converter.IsDefaultResponse(r.res.PutParameters, outputs, r.res.DefaultBody)
	})
}

func (r *resourceCrudClient) CreateOrUpdate(ctx context.Context, id string, bodyParams, queryParams map[string]any) (map[string]any, bool, error) {
	// Submit the `PUT` or `PATCH` against the ARM endpoint
	op := r.azureClient.Put
	if r.res.UpdateMethod == "PATCH" {
		op = r.azureClient.Patch
	}
	return op(ctx, id, bodyParams, queryParams, r.res.PutAsyncStyle)
}

func (r *resourceCrudClient) Read(ctx context.Context, id string) (map[string]any, error) {
	url := id + r.res.ReadPath

	var resp any
	var err error
	switch r.res.ReadMethod {
	case "HEAD":
		err = r.azureClient.Head(ctx, url, r.res.APIVersion)
		return nil, err
	case "POST":
		bodyParams := map[string]interface{}{}
		queryParams := map[string]interface{}{
			"api-version": r.res.APIVersion,
		}
		resp, err = r.azureClient.Post(ctx, url, bodyParams, queryParams)
	default:
		resp, err = r.azureClient.Get(ctx, url, r.res.APIVersion, r.res.ReadQueryParams)
	}

	if err != nil {
		return nil, err
	}
	// Cast should be safe because we're reading a resource, only invokes return non-objects.
	if respMap, ok := resp.(map[string]any); ok {
		return respMap, nil
	}
	return nil, errors.Errorf("expected object from Read of '%s', got %T", url, resp)
}

type propertyPath struct {
	path         []string
	propertyName string
}

func writePropertiesToBody(missingProperties []propertyPath, bodyParams map[string]interface{}, remoteState map[string]interface{}) map[string]interface{} {
	writtenProperties := map[string]interface{}{}
	for _, prop := range missingProperties {
		currentBodyContainer := bodyParams
		currentStateContainer := remoteState
		for _, containerName := range prop.path {
			innerBodyContainer, bodyOk := currentBodyContainer[containerName]
			innerStateContainer, stateOk := currentStateContainer[containerName]
			// If the container doesn't exist in either body or state, create it and continue iterating.
			// But if it doesn't exist in either, there is no point in continuing.
			if !bodyOk && !stateOk {
				break
			}
			if !bodyOk {
				innerBodyContainer = map[string]interface{}{}
				currentBodyContainer[containerName] = innerBodyContainer
			}
			if !stateOk {
				innerStateContainer = map[string]interface{}{}
				currentStateContainer[containerName] = innerStateContainer
			}
			innerBodyObj, innerBodyIsObject := innerBodyContainer.(map[string]interface{})
			innerStateObj, innerStateIsObject := innerStateContainer.(map[string]interface{})
			if !innerBodyIsObject || !innerStateIsObject { // we've reached a leaf node (primitive type)
				break
			}
			currentBodyContainer = innerBodyObj
			currentStateContainer = innerStateObj
		}

		stateValue, ok := currentStateContainer[prop.propertyName]
		if ok {
			currentBodyContainer[prop.propertyName] = stateValue
			writtenProperties[fmt.Sprintf("%s.%s", strings.Join(prop.path, "."), prop.propertyName)] = stateValue
		}
	}
	return writtenProperties
}
