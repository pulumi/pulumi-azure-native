package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

const webAppResourceType = "azure-native:web:WebApp"

// https://github.com/pulumi/pulumi-azure-native/issues/1529
func customWebAppDelete(resourceLookupper ResourceLookupper, deleter AzureDeleter) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Web/sites/{name}",
		tok:  webAppResourceType,
		Delete: func(ctx context.Context, id string, properties resource.PropertyMap) error {
			res, ok, err := resourceLookupper.LookupResource(webAppResourceType)
			if err != nil {
				return err
			}
			if !ok {
				return fmt.Errorf("resource %q not found", webAppResourceType)
			}

			return deleter.AzureDelete(ctx, id, res.APIVersion, res.DeleteAsyncStyle, map[string]any{
				// Don't delete the app service plan even if this was the last web app in it. It's
				// a separate Pulumi resource.
				"deleteEmptyServerFarm": false,
			})
		},
	}
}
