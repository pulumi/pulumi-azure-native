// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/util"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/versionLookup"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const webAppResourceType = "azure-native:web:WebApp"
const webAppSlotResourceType = "azure-native:web:WebAppSlot"

// webApp is a custom resource for web:WebApp. It overrides Read and Delete, for independent reasons.
func webApp(crudClientFactory crud.ResourceCrudClientFactory, azureClient azure.AzureClient, lookupResource resources.ResourceLookupFunc) (*CustomResource, error) {
	return makeWebAppResource(webAppResourceType,
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}",
		crudClientFactory, azureClient, lookupResource)
}

// webAppSlot is a custom resource for web:WebAppSlot. It overrides Read and Delete, for independent reasons.
func webAppSlot(crudClientFactory crud.ResourceCrudClientFactory, azureClient azure.AzureClient, lookupResource resources.ResourceLookupFunc) (*CustomResource, error) {
	return makeWebAppResource(webAppSlotResourceType,
		"/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}/slots/{slot}",
		crudClientFactory, azureClient, lookupResource)
}

// webApp is a custom resource for web:WebApp. It overrides Read and Delete, for independent reasons.
func makeWebAppResource(resourceType, path string, crudClientFactory crud.ResourceCrudClientFactory, azureClient azure.AzureClient, lookupResource resources.ResourceLookupFunc) (*CustomResource, error) {
	// The arguments can be nil when the custom resource is constructed only for feature lookups
	var crudClient crud.ResourceCrudClient
	if crudClientFactory != nil && lookupResource != nil {
		var err error
		crudClient, err = createCrudClient(crudClientFactory, lookupResource, resourceType)
		if err != nil {
			return nil, err
		}
	}

	return &CustomResource{
		path: path,
		tok:  resourceType,

		// WebApp.SiteConfig is created and updated via a PUT to the WebApp itself, but a GET of
		// the WebApp does not return the SiteConfig. We need to make a separate GET request to
		// /config/web and merge the results. #1468
		Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
			// We assume that WebApp and WebAppSlot share their API version since they are almost identical.
			apiVersion, ok := versionLookup.GetDefaultApiVersionForResource("Web", "WebApp")
			if !ok {
				apiVersion = "2022-09-01" // default as of 2024-07
				logging.V(3).Infof("Warning: could not find default API version for %s. Using %s", resourceType, apiVersion)
			}

			webAppResponse, err := crudClient.Read(ctx, id)
			if err != nil {
				if azure.IsNotFound(err) {
					return nil, false, nil
				}
				return nil, false, err
			}

			siteConfigResponseRaw, err := azureClient.Get(ctx, id+"/config/web", apiVersion, nil)
			if err != nil {
				return nil, false, err
			}
			siteConfigResponse, ok := siteConfigResponseRaw.(map[string]any)
			if !ok {
				return nil, true, fmt.Errorf("unexpected response type %T", siteConfigResponseRaw)
			}

			if err := mergeWebAppSiteConfig(webAppResponse, siteConfigResponse); err != nil {
				return nil, true, err
			}

			responseInSdkShape := crudClient.ResponseBodyToSdkOutputs(webAppResponse)
			filterResponse(responseInSdkShape)
			return responseInSdkShape, true, nil
		},

		// https://github.com/pulumi/pulumi-azure-native/issues/1529
		Delete: func(ctx context.Context, id string, previousInputs resource.PropertyMap, state resource.PropertyMap) error {
			res, ok, err := lookupResource(webAppResourceType)
			if err != nil {
				return err
			}
			if !ok {
				return fmt.Errorf("resource %q not found", webAppResourceType)
			}

			return azureClient.Delete(ctx, id, res.APIVersion, res.DeleteAsyncStyle, map[string]any{
				// Don't delete the app service plan even if this was the last web app in it. It's
				// a separate Pulumi resource.
				"deleteEmptyServerFarm": false,
			})
		},
	}, nil
}

// A generic hook to modify the response returned after Read() of a WebApp.
// Removes redacted publishingUsername from siteConfig to avoid meaningless diffs, #1468.
func filterResponse(response map[string]any) {
	if siteConfig, ok := util.GetInnerMap(response, "siteConfig"); ok {
		if username, ok := siteConfig["publishingUsername"]; ok {
			if usernameStr, ok := username.(string); ok && usernameStr == "REDACTED" {
				delete(siteConfig, "publishingUsername")
			}
		}
	}
}

// Add siteConfig, that we got from a separate request, to the webApp response.
func mergeWebAppSiteConfig(webApp, siteConfig map[string]any) error {
	if outerProperties, ok := util.GetInnerMap(webApp, "properties"); ok {
		outerProperties["siteConfig"] = siteConfig["properties"]
		return nil
	}
	return fmt.Errorf("'properties' not found in response, cannot update siteConfig")
}
