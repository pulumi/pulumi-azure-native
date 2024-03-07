package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagesync/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagesync.NewStorageSyncService(ctx, "storageSyncService", &storagesync.StorageSyncServiceArgs{
			IncomingTrafficPolicy:  pulumi.String("AllowAllTraffic"),
			Location:               pulumi.String("WestUS"),
			ResourceGroupName:      pulumi.String("SampleResourceGroup_1"),
			StorageSyncServiceName: pulumi.String("SampleStorageSyncService_1"),
			Tags:                   nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
