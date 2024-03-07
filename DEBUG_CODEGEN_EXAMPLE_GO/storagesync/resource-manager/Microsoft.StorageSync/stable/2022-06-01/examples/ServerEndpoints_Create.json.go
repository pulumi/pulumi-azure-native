package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagesync/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagesync.NewServerEndpoint(ctx, "serverEndpoint", &storagesync.ServerEndpointArgs{
			CloudTiering:                 pulumi.String("off"),
			InitialDownloadPolicy:        pulumi.String("NamespaceThenModifiedFiles"),
			InitialUploadPolicy:          pulumi.String("ServerAuthoritative"),
			LocalCacheMode:               pulumi.String("UpdateLocallyCachedFiles"),
			OfflineDataTransfer:          pulumi.String("on"),
			OfflineDataTransferShareName: pulumi.String("myfileshare"),
			ResourceGroupName:            pulumi.String("SampleResourceGroup_1"),
			ServerEndpointName:           pulumi.String("SampleServerEndpoint_1"),
			ServerLocalPath:              pulumi.String("D:\\SampleServerEndpoint_1"),
			ServerResourceId:             pulumi.String("/subscriptions/52b8da2f-61e0-4a1f-8dde-336911f367fb/resourceGroups/SampleResourceGroup_1/providers/Microsoft.StorageSync/storageSyncServices/SampleStorageSyncService_1/registeredServers/080d4133-bdb5-40a0-96a0-71a6057bfe9a"),
			StorageSyncServiceName:       pulumi.String("SampleStorageSyncService_1"),
			SyncGroupName:                pulumi.String("SampleSyncGroup_1"),
			TierFilesOlderThanDays:       pulumi.Int(0),
			VolumeFreeSpacePercent:       pulumi.Int(100),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
