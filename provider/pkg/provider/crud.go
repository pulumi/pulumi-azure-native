package provider

import (
	"context"
	"fmt"
	"net/url"
	"strings"

	structpb "github.com/golang/protobuf/ptypes/struct"
	"github.com/pkg/errors"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

type ResourceCrudClient struct {
	azureClient azure.AzureClient
	provider    *azureNativeProvider
	res         *resources.AzureAPIResource
	id          string
	inputs      resource.PropertyMap
}

func (r *ResourceCrudClient) PrepareAzureRESTInputs(methodInputs map[string]any) (string, map[string]any, map[string]any, error) {
	clientInputs := map[string]interface{}{
		"subscriptionId": r.provider.subscriptionID,
		"api-version":    r.res.APIVersion,
	}

	// Schema-driven mapping of inputs into Autorest id/body/query
	params := map[string]map[string]interface{}{
		"query": {
			"api-version": clientInputs["api-version"],
		},
		"path": {},
	}

	// Build maps of path and query parameters.
	for _, param := range r.res.PutParameters {
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
	id := r.res.Path
	for key, value := range params["path"] {
		encodedVal := strings.Replace(url.QueryEscape(value.(string)), "+", "%20", -1)
		id = strings.Replace(id, "{"+key+"}", encodedVal, -1)
	}

	// Build the body JSON.
	var body map[string]interface{}
	for _, param := range r.res.PutParameters {
		if param.Location != "body" {
			continue
		}
		body = r.provider.converter.SdkInputsToRequestBody(param.Body.Properties, methodInputs, r.id)
		break
	}

	// Ensure all required containers are created.
	for _, containers := range r.res.RequiredContainers {
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

func (r *ResourceCrudClient) HandleErrorWithCheckpoint(ctx context.Context, err error, properties *structpb.Struct) error {
	// Resource was partially updated but the operation failed to complete.
	// Try reading its state by ID and return a partial error if succeeded.
	checkpoint, getErr := r.provider.currentResourceStateCheckpoint(ctx, r.id, *r.res, r.inputs)
	if getErr != nil {
		return azure.AzureError(errors.Wrapf(err, "resource updated but read failed %s", getErr))
	}
	return partialError(r.id, err, checkpoint, properties)
}

func (r *ResourceCrudClient) MaintainSubResourcePropertiesIfNotSet(ctx context.Context, bodyParams map[string]interface{}) error {
	// Identify the properties we need to read
	missingProperties := r.provider.findUnsetPropertiesToMaintain(r.res, bodyParams, true /* returnApiShapePaths */)

	if len(missingProperties) == 0 {
		// Everything's already specified explicitly by the user, no need to do read.
		return nil
	}

	// Read the current resource state.
	state, err := r.provider.azureClient.Get(ctx, r.id+r.res.ReadPath, r.res.APIVersion)
	if err != nil {
		return fmt.Errorf("reading cloud state: %w", err)
	}

	writtenProperties := writePropertiesToBody2(missingProperties, bodyParams, state)
	for writtenProperty, writtenValue := range writtenProperties {
		logging.V(9).Infof("Maintaining remote value for property: %s.%s = %v", r.id, writtenProperty, writtenValue)
	}

	return nil
}

func (r *ResourceCrudClient) ResponseBodyToSdkOutputs(
	response map[string]any,
) map[string]any {
	return r.provider.converter.ResponseBodyToSdkOutputs(r.res.Response, response)
}

func (r *ResourceCrudClient) Put(ctx context.Context, bodyProps map[string]interface{}, queryParameters map[string]interface{}) (map[string]interface{}, bool, error) {
	return r.azureClient.Put(ctx, r.id, bodyProps, queryParameters, r.res.UpdateMethod, r.res.PutAsyncStyle)
}

func writePropertiesToBody2(missingProperties []propertyPath, bodyParams map[string]interface{}, remoteState map[string]interface{}) map[string]interface{} {
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

func (k *azureNativeProvider) findUnsetPropertiesToMaintain(res *resources.AzureAPIResource, bodyParams map[string]interface{}, returnApiShapePaths bool) []propertyPath {
	missingProperties := []propertyPath{}
	for _, path := range res.PathsToSubResourcePropertiesToMaintain(returnApiShapePaths, k.lookupType) {
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
