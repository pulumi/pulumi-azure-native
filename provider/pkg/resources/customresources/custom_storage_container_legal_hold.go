// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	rpc "github.com/pulumi/pulumi/sdk/v3/proto/go"
)

const legalHold = "legalHold"

func storageContainerWithLegalHold(azureClient azure.AzureClient) *CustomResource {
	return &CustomResource{
		tok: "azure-native:storage:BlobContainer",
		//path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}",

		SchemaOverlay: &schema.ResourceSpec{
			ObjectTypeSpec: schema.ObjectTypeSpec{
				Properties: map[string]schema.PropertySpec{
					legalHold: {TypeSpec: schema.TypeSpec{Ref: "#/types/azure-native:storage:LegalHold"}},
				},
			},
			InputProperties: map[string]schema.PropertySpec{
				legalHold: {TypeSpec: schema.TypeSpec{Ref: "#/types/azure-native:storage:LegalHold"}},
			},
		},

		MetaOverlay: &resources.AzureAPIResource{
			PutParameters: []resources.AzureAPIParameter{
				{
					Location: "body",
					Body: &resources.AzureAPIType{
						Properties: map[string]resources.AzureAPIProperty{
							legalHold: {Ref: "#/types/azure-native:storage:LegalHold"},
						},
					},
				},
			},
		},

		Create: func(ctx context.Context, req *rpc.CreateRequest, crudClient crud.ResourceCrudClient) (map[string]interface{}, error) {
			return nil, nil
		},

		Read: func(ctx context.Context, id string, properties resource.PropertyMap) (map[string]interface{}, bool, error) {
			return nil, false, nil
		},

		Update: func(ctx context.Context, req *rpc.UpdateRequest, crudClient crud.ResourceCrudClient) (map[string]interface{}, error) {
			id, bodyParams, queryParams, err := crudClient.PrepareAzureRESTInputs()
			if err != nil {
				return nil, err
			}

			err = crudClient.MaintainSubResourcePropertiesIfNotSet(ctx, id, bodyParams)
			if err != nil {
				return nil, fmt.Errorf("failed maintaining unset sub-resource properties: %w", err)
			}

			response, updated, err := crudClient.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				if updated {
					return nil, crudClient.HandleErrorWithCheckpoint(ctx, err, id, req.GetNews())
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
