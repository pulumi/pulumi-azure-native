package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerstorage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerstorage.NewSnapshot(ctx, "snapshot", &containerstorage.SnapshotArgs{
			PoolName:          pulumi.String("test-pool"),
			ResourceGroupName: pulumi.String("test-rg"),
			SnapshotName:      pulumi.String("test-snapshot"),
			Source:            pulumi.String("C0C6I6"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
