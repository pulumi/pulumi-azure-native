// Copyright 2024, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

// Server configuration cannot be deleted, only reset to the default value, and to do that we need
// to retrieve the default value first. See
// [the implementation in `az`](https://github.com/Azure/azure-cli/blob/f92f7233659890b427f4f6f486db5610cdca1020/src/azure-cli/azure/cli/command_modules/rdbms/flexible_server_custom_postgres.py#L473)
// which does the same thing.
// The Azure Postgres team confirmed that both source and value must be set:
// https://github.com/Azure/azure-rest-api-specs/issues/30143#issuecomment-2379605761.
func postgresFlexibleServerConfiguration(crudClientFactory crud.ResourceCrudClientFactory, lookupResource resources.ResourceLookupFunc) (*CustomResource, error) {
	// The arguments can be nil when the custom resource is constructed only for feature lookups
	var crudClient crud.ResourceCrudClient
	if crudClientFactory != nil && lookupResource != nil {
		var err error
		crudClient, err = createCrudClient(crudClientFactory, lookupResource, "azure-native:dbforpostgresql:Configuration")
		if err != nil {
			return nil, err
		}
	}

	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.DBforPostgreSQL/flexibleServers/{serverName}/configurations/{configurationName}",
		Delete: func(ctx context.Context, id string, inputs, state resource.PropertyMap) error {
			conf, err := crudClient.Read(ctx, id)
			if err != nil {
				return err
			}

			defaultValue := getValueFromAzureResponse(conf, "defaultValue")
			if defaultValue == nil {
				return fmt.Errorf("unexpected Postgres Configuration does not have a default value: %v", conf)
			}

			// Also read the configuration "source" and send it back. This is actually wrong, but
			// quoting the az implementation:
			// > "this should be 'system-default' but there is currently a bug in PG, so keeping as
			// > what it is for now this will reset source to be 'system-default' anyway"
			source := getValueFromAzureResponse(conf, "source")
			if source == nil {
				return fmt.Errorf("unexpected Postgres Configuration does not have a source: %v", conf)
			}

			_, queryParams, err := crudClient.PrepareAzureRESTIdAndQuery(nil)
			if err != nil {
				return err
			}
			result, _, err := crudClient.Update(ctx, id, map[string]any{
				"properties": map[string]any{
					"value":  defaultValue,
					"source": source,
				},
			}, queryParams)
			if err != nil {
				return err
			}

			newValue := getValueFromAzureResponse(result, "value")
			if newValue != defaultValue {
				return fmt.Errorf("failed to reset Postgres Configuration to default value '%v': %v", defaultValue, newValue)
			}
			return nil
		},
		isSingleton: true,
	}, nil
}

func getValueFromAzureResponse(resp map[string]any, key string) any {
	if props, ok := resp["properties"]; ok {
		if propsMap, ok := props.(map[string]any); ok {
			return propsMap[key]
		}
	}
	return nil
}
