package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
			CreationData: &compute.CreationDataArgs{
				CreateOption:     pulumi.String("Import"),
				SourceUri:        pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
				StorageAccountId: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SnapshotName:      pulumi.String("mySnapshot1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
