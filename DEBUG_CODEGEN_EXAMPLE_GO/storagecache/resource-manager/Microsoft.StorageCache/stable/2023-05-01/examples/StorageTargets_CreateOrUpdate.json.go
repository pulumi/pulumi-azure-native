package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagecache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagecache.NewStorageTarget(ctx, "storageTarget", &storagecache.StorageTargetArgs{
			CacheName: pulumi.String("sc1"),
			Junctions: storagecache.NamespaceJunctionArray{
				&storagecache.NamespaceJunctionArgs{
					NamespacePath:   pulumi.String("/path/on/cache"),
					NfsAccessPolicy: pulumi.String("default"),
					NfsExport:       pulumi.String("exp1"),
					TargetPath:      pulumi.String("/path/on/exp1"),
				},
				&storagecache.NamespaceJunctionArgs{
					NamespacePath:   pulumi.String("/path2/on/cache"),
					NfsAccessPolicy: pulumi.String("rootSquash"),
					NfsExport:       pulumi.String("exp2"),
					TargetPath:      pulumi.String("/path2/on/exp2"),
				},
			},
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
