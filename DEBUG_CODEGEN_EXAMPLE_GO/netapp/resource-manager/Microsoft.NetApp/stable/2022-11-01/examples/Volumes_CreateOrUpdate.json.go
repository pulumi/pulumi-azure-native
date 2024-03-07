package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewVolume(ctx, "volume", &netapp.VolumeArgs{
			AccountName:       pulumi.String("account1"),
			CreationToken:     pulumi.String("my-unique-file-path"),
			Location:          pulumi.String("eastus"),
			PoolName:          pulumi.String("pool1"),
			ResourceGroupName: pulumi.String("myRG"),
			ServiceLevel:      pulumi.String("Premium"),
			SubnetId:          pulumi.String("/subscriptions/9760acf5-4638-11e7-9bdb-020073ca7778/resourceGroups/myRP/providers/Microsoft.Network/virtualNetworks/testvnet3/subnets/testsubnet3"),
			UsageThreshold:    pulumi.Float64(107374182400),
			VolumeName:        pulumi.String("volume1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
