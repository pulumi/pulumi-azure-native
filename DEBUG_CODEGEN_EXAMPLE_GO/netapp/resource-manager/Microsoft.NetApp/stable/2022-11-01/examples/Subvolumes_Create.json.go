package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewSubvolume(ctx, "subvolume", &netapp.SubvolumeArgs{
			AccountName:       pulumi.String("account1"),
			Path:              pulumi.String("/subvolumePath"),
			PoolName:          pulumi.String("pool1"),
			ResourceGroupName: pulumi.String("myRG"),
			SubvolumeName:     pulumi.String("subvolume1"),
			VolumeName:        pulumi.String("volume1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
