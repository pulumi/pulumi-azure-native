package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagesync/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagesync.NewSyncGroup(ctx, "syncGroup", &storagesync.SyncGroupArgs{
			ResourceGroupName:      pulumi.String("SampleResourceGroup_1"),
			StorageSyncServiceName: pulumi.String("SampleStorageSyncService_1"),
			SyncGroupName:          pulumi.String("SampleSyncGroup_1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
