// Copyright 2021, Pulumi Corporation.  All rights reserved.

package customresources

import (
	"context"
	"fmt"

	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/azure"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/provider/crud"
	"github.com/pulumi/pulumi-azure-native/v2/provider/pkg/resources"
	"google.golang.org/protobuf/types/known/structpb"

	"github.com/pulumi/pulumi/pkg/v3/codegen/schema"
	"github.com/pulumi/pulumi/sdk/v3/go/common/resource"
	"github.com/pulumi/pulumi/sdk/v3/go/common/util/logging"
)

const legalHold = "legalHold"

func storageContainerWithLegalHold() *CustomResource {
	return &CustomResource{
		tok:  "azure-native:storage:BlobContainer",
		path: "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Storage/storageAccounts/{accountName}/blobServices/default/containers/{containerName}",
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

		Create: func(ctx context.Context, properties resource.PropertyMap, rawInputs *structpb.Struct, crudClient crud.ResourceCrudClient) (map[string]interface{}, error) {
			id, bodyParams, queryParams, err := crudClient.PrepareAzureRESTInputs(properties)
			if err != nil {
				return nil, err
			}

			// First check if the resource already exists - we want to try our best to avoid updating instead of creating here
			// (though it's technically impossible since the only operation supported is an upsert).
			err = crudClient.CanCreate(ctx, id)
			if err != nil {
				return nil, err
			}

			response, created, err := crudClient.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				if created {
					return nil, crudClient.HandleErrorWithCheckpoint(ctx, err, id, properties, rawInputs)
				}
				return nil, azure.AzureError(err)
			}

			if lh, ok := bodyParams[legalHold]; ok {
				logging.V(3).Infof("Container with Legal Hold: %v", bodyParams)
				if lhObj, ok := lh.(map[string]any); ok {
					_, err := crudClient.Post(ctx, id+"/setLegalHold", lhObj, queryParams)
					if err != nil {
						return nil, crudClient.HandleErrorWithCheckpoint(ctx, err, id, properties, rawInputs)
					}

					response, err = crudClient.Get(ctx, id)
					if err != nil {
						return nil, fmt.Errorf("failed to get the container after setting legal hold: %w", err)
					}
				}
			}

			// Map the raw response to the shape of outputs that the SDKs expect.
			outputs := crudClient.ResponseBodyToSdkOutputs(response)
			return outputs, nil
		},

		Update: func(ctx context.Context, properties resource.PropertyMap, rawInputs *structpb.Struct, crudClient crud.ResourceCrudClient) (map[string]interface{}, error) {
			id, bodyParams, queryParams, err := crudClient.PrepareAzureRESTInputs(properties)
			if err != nil {
				return nil, err
			}

			response, updated, err := crudClient.CreateOrUpdate(ctx, id, bodyParams, queryParams)
			if err != nil {
				if updated {
					return nil, crudClient.HandleErrorWithCheckpoint(ctx, err, id, properties, rawInputs)
				}
				return nil, azure.AzureError(err)
			}

			// Map the raw response to the shape of outputs that the SDKs expect.
			outputs := crudClient.ResponseBodyToSdkOutputs(response)

			return outputs, nil
		},
	}
}
