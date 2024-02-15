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

type ResourceCrudClient struct {
	azureClient    azure.AzureClient
	lookupType     resources.TypeLookupFunc
	converter      *convert.SdkShapeConverter
	subscriptionID string

	res *resources.AzureAPIResource
}

func NewResourceCrudClient(
	azureClient azure.AzureClient,
	lookupType resources.TypeLookupFunc,
	converter *convert.SdkShapeConverter,
	subscriptionID string,
	res *resources.AzureAPIResource,
) *ResourceCrudClient {
	return &ResourceCrudClient{
		azureClient:    azureClient,
		lookupType:     lookupType,
		converter:      converter,
		subscriptionID: subscriptionID,
		res:            res,
	}
}

func (r *ResourceCrudClient) PrepareAzureRESTInputs(inputs resource.PropertyMap) (string, map[string]any, map[string]any, error) {
	return PrepareAzureRESTInputs(r.res.Path, r.res.PutParameters, r.res.RequiredContainers, inputs.Mappable(), map[string]any{
		"subscriptionId": r.subscriptionID,
		"api-version":    r.res.APIVersion,
	}, r.converter)
}

func PrepareAzureRESTInputs(path string, parameters []resources.AzureAPIParameter, requiredContainers [][]string, methodInputs,
	clientInputs map[string]interface{}, converter *convert.SdkShapeConverter) (string, map[string]interface{}, map[string]interface{}, error) {
	// Schema-driven mapping of inputs into Autorest id/body/query
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
		val, has = methodInputs[sdkName]
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
		encodedVal := strings.Replace(url.QueryEscape(value.(string)), "+", "%20", -1)
		id = strings.Replace(id, "{"+key+"}", encodedVal, -1)
	}

	// Build the body JSON.
	var body map[string]interface{}
	for _, param := range parameters {
		if param.Location != "body" {
			continue
		}
		body = converter.SdkInputsToRequestBody(param.Body.Properties, methodInputs, id)
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

	return id, body, params["query"], nil
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

func (r *ResourceCrudClient) HandleErrorWithCheckpoint(ctx context.Context, err error, id string, inputs resource.PropertyMap, properties *structpb.Struct) error {
	// Resource was partially updated but the operation failed to complete.
	// Try reading its state by ID and return a partial error if succeeded.
	checkpoint, getErr := r.currentResourceStateCheckpoint(ctx, id, inputs)
	if getErr != nil {
		return azure.AzureError(errors.Wrapf(err, "resource updated but read failed %s", getErr))
	}
	return partialError(id, err, checkpoint, properties)
}

// currentResourceStateCheckpoint reads the resource state by ID, converts it to outputs map, and
// produces a checkpoint with these outputs and given inputs.
func (r *ResourceCrudClient) currentResourceStateCheckpoint(ctx context.Context, id string, inputs resource.PropertyMap) (*structpb.Struct, error) {
	getResp, getErr := r.azureClient.Get(ctx, id, r.res.APIVersion)
	if getErr != nil {
		return nil, getErr
	}
	outputs := r.converter.ResponseBodyToSdkOutputs(r.res.Response, getResp)
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

func (r *ResourceCrudClient) MaintainSubResourcePropertiesIfNotSet(ctx context.Context, id string, bodyParams map[string]interface{}) error {
	// Identify the properties we need to read
	missingProperties := findUnsetPropertiesToMaintain(r.res, bodyParams, true /* returnApiShapePaths */, r.lookupType)

	if len(missingProperties) == 0 {
		// Everything's already specified explicitly by the user, no need to do read.
		return nil
	}

	// Read the current resource state.
	state, err := r.azureClient.Get(ctx, id+r.res.ReadPath, r.res.APIVersion)
	if err != nil {
		return fmt.Errorf("reading cloud state: %w", err)
	}

	writtenProperties := writePropertiesToBody(missingProperties, bodyParams, state)
	for writtenProperty, writtenValue := range writtenProperties {
		logging.V(9).Infof("Maintaining remote value for property: %s.%s = %v", id, writtenProperty, writtenValue)
	}

	return nil
}

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
func (r *ResourceCrudClient) SetUnsetSubresourcePropertiesToDefaults(input, output map[string]interface{}, outputIsInApiShape bool) {
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

func (r *ResourceCrudClient) ResponseBodyToSdkOutputs(
	response map[string]any,
) map[string]any {
	return r.converter.ResponseBodyToSdkOutputs(r.res.Response, response)
}

func (r *ResourceCrudClient) CanCreate(ctx context.Context, id string) error {
	return r.azureClient.CanCreate(ctx, id, r.res.ReadMethod, r.res.APIVersion, r.res.ReadMethod, r.res.Singleton, r.res.DefaultBody != nil, func(outputs map[string]any) bool {
		return r.converter.IsDefaultResponse(r.res.PutParameters, outputs, r.res.DefaultBody)
	})
}

func (r *ResourceCrudClient) CreateOrUpdate(ctx context.Context, id string, bodyParams, queryParams map[string]any) (map[string]any, bool, error) {
	// Submit the `PUT` or `PATCH` against the ARM endpoint
	op := r.azureClient.Put
	if r.res.UpdateMethod == "PATCH" {
		op = r.azureClient.Patch
	}
	return op(ctx, id, bodyParams, queryParams, r.res.PutAsyncStyle)
}

func (r *ResourceCrudClient) Read(ctx context.Context, id string) (map[string]any, error) {
	url := id + r.res.ReadPath

	switch r.res.ReadMethod {
	case "HEAD":
		err := r.azureClient.Head(ctx, url, r.res.APIVersion)
		return nil, err
	case "POST":
		bodyParams := map[string]interface{}{}
		queryParams := map[string]interface{}{
			"api-version": r.res.APIVersion,
		}
		return r.azureClient.Post(ctx, url, bodyParams, queryParams)
	default:
		return r.azureClient.Get(ctx, url, r.res.APIVersion)
	}
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
