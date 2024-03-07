package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewSnapshot(ctx, "snapshot", &netapp.SnapshotArgs{
			AccountName:       pulumi.String("account1"),
			Location:          pulumi.String("eastus"),
			PoolName:          pulumi.String("pool1"),
			ResourceGroupName: pulumi.String("myRG"),
			SnapshotName:      pulumi.String("snapshot1"),
			VolumeName:        pulumi.String("volume1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
