package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagecache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagecache.NewStorageTarget(ctx, "storageTarget", &storagecache.StorageTargetArgs{
			BlobNfs: &storagecache.BlobNfsTargetArgs{
				Target:            pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/scgroup/providers/Microsoft.Storage/storageAccounts/blofnfs/blobServices/default/containers/blobnfs"),
				UsageModel:        pulumi.String("READ_WRITE"),
				VerificationTimer: pulumi.Int(28800),
				WriteBackTimer:    pulumi.Int(3600),
			},
			CacheName: pulumi.String("sc1"),
			Junctions: storagecache.NamespaceJunctionArray{
				&storagecache.NamespaceJunctionArgs{
					NamespacePath: pulumi.String("/blobnfs"),
				},
			},
			ResourceGroupName: pulumi.String("scgroup"),
			StorageTargetName: pulumi.String("st1"),
			TargetType:        pulumi.String("blobNfs"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
