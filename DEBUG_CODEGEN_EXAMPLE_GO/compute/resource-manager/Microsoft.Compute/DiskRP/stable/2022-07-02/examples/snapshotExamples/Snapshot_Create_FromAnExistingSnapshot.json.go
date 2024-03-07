package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewSnapshot(ctx, "snapshot", &compute.SnapshotArgs{
			CreationData: &compute.CreationDataArgs{
				CreateOption:     pulumi.String("Copy"),
				SourceResourceId: pulumi.String("subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot1"),
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SnapshotName:      pulumi.String("mySnapshot2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
