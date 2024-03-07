package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerstorage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerstorage.NewVolume(ctx, "volume", &containerstorage.VolumeArgs{
			CapacityGiB: pulumi.Float64(25838),
			Labels: pulumi.StringMap{
				"key2039": pulumi.String("value2039"),
			},
			PoolName:          pulumi.String("test-pool"),
			ResourceGroupName: pulumi.String("test-rg"),
			VolumeName:        pulumi.String("test-volume"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
