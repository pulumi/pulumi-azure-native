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

const webAppTok = "azure-native:web:WebApp"

// webApp is a custom resource for web:WebApp. It overrides Read and Delete, for independent reasons.
func webApp(crudClientFactory crud.ResourceCrudClientFactory, azureClient azure.AzureClient, lookupResource resources.ResourceLookupFunc) (*CustomResource, error) {
	// The arguments can be nil when the custom resource is constructed only for feature lookups
	var crudClient crud.ResourceCrudClient
	if crudClientFactory != nil && lookupResource != nil {
		var err error
		crudClient, err = createCrudClient(crudClientFactory, lookupResource)
		if err != nil {
			return nil, err
		}
	}

	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}",
		tok:  webAppTok,

		// WebApp.SiteConfig is created and updated via a PUT to the WebApp itself, but a GET of
		// the WebApp does not return the SiteConfig. We need to make a separate GET request to
		// /config/web and merge the results. #1468
		Read: func(ctx context.Context, id string, inputs resource.PropertyMap) (map[string]any, bool, error) {
			apiVersion, ok := versionLookup.GetDefaultApiVersionForResource("Storage", "BlobContainer")
			if !ok {
				// default as of 2024-02-16
				apiVersion = "2022-09-01"
				logging.V(3).Infof("Warning: could not find default API version for %s. Using %s", webAppTok, apiVersion)
			}

			webAppResponse, err := crudClient.Read(ctx, id)
			if err != nil {
				return nil, false, err
			}

			siteConfigResponseRaw, err := azureClient.Get(ctx, id+"/config/web", apiVersion)
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

			return crudClient.ResponseBodyToSdkOutputs(webAppResponse), true, nil
		},

		// https://github.com/pulumi/pulumi-azure-native/issues/1529
		Delete: func(ctx context.Context, id string, properties resource.PropertyMap) error {
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

func mergeWebAppSiteConfig(webApp, siteConfig map[string]any) error {
	if outerProperties, ok := util.GetInnerMap(webApp, "properties"); ok {
		outerProperties["siteConfig"] = siteConfig["properties"]
		return nil
	}
	return fmt.Errorf("'properties' not found in response, cannot update siteConfig")
}

func createCrudClient(crudClientFactory crud.ResourceCrudClientFactory, lookupResource resources.ResourceLookupFunc) (crud.ResourceCrudClient, error) {
	res, ok, err := lookupResource(webAppTok)
	if err != nil {
		return nil, err
	}
	if !ok {
		return nil, fmt.Errorf("resource %s not found", webAppTok)
	}

	return crudClientFactory(&res), nil
}
