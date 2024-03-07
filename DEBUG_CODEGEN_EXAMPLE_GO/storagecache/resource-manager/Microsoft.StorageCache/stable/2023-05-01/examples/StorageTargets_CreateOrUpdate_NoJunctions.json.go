package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagecache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagecache.NewStorageTarget(ctx, "storageTarget", &storagecache.StorageTargetArgs{
			CacheName: pulumi.String("sc1"),
			Nfs3: &storagecache.Nfs3TargetArgs{
				Target:            pulumi.String("10.0.44.44"),
				UsageModel:        pulumi.String("READ_ONLY"),
				VerificationTimer: pulumi.Int(30),
			},
			ResourceGroupName: pulumi.String("scgroup"),
			StorageTargetName: pulumi.String("st1"),
			TargetType:        pulumi.String("nfs3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
