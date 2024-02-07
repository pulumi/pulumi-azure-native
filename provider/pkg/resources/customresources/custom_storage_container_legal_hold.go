// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	// . "github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
)

func storageContainerWithLegalHold(azureClient azure.AzureClient) *CustomResource {
	return &CustomResource{
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}",
		Create: func(ctx context.Context, properties resource.PropertyMap) (map[string]interface{}, error) {
			return nil, nil
		},
		Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
			return nil, false, nil
		},
		Update: func(ctx context.Context, properties resource.PropertyMap) (map[string]interface{}, error) {
			err := crudClient.MaintainSubResourcePropertiesIfNotSet(ctx, bodyParams)
			if err != nil {
				return nil, fmt.Errorf("failed maintaining unset sub-resource properties: %w", err)
			}

			response, updated, err := crudClient.Put(ctx, bodyParams, queryParams)
			if err != nil {
				if updated {
					return nil, crudClient.HandleErrorWithCheckpoint(ctx, err, req.GetNews())
				}
				return nil, azure.AzureError(err)
			}

			// Map the raw response to the shape of outputs that the SDKs expect.
			outputs := crudClient.ResponseBodyToSdkOutputs(response)

			return outputs, nil
		},
		Delete: func(ctx context.Context, id string, properties resource.PropertyMap) error {
			return nil
		},
	}
}
