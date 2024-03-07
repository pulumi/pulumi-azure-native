package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagesync/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagesync.NewCloudEndpoint(ctx, "cloudEndpoint", &storagesync.CloudEndpointArgs{
			AzureFileShareName:       pulumi.String("cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4"),
			CloudEndpointName:        pulumi.String("SampleCloudEndpoint_1"),
			FriendlyName:             pulumi.String("ankushbsubscriptionmgmtmab"),
			ResourceGroupName:        pulumi.String("SampleResourceGroup_1"),
			StorageAccountResourceId: pulumi.String("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage"),
			StorageAccountTenantId:   pulumi.String("\"72f988bf-86f1-41af-91ab-2d7cd011db47\""),
			StorageSyncServiceName:   pulumi.String("SampleStorageSyncService_1"),
			SyncGroupName:            pulumi.String("SampleSyncGroup_1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
