package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagemover/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagemover.NewEndpoint(ctx, "endpoint", &storagemover.EndpointArgs{
			EndpointName: pulumi.String("examples-endpointName"),
			Properties: storagemover.AzureStorageBlobContainerEndpointProperties{
				BlobContainerName:        "examples-blobContainerName",
				Description:              "Example Storage Container Endpoint Description",
				EndpointType:             "AzureStorageBlobContainer",
				StorageAccountResourceId: "/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/examples-rg/providers/Microsoft.Storage/storageAccounts/examples-storageAccountName/",
			},
			ResourceGroupName: pulumi.String("examples-rg"),
			StorageMoverName:  pulumi.String("examples-storageMoverName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
