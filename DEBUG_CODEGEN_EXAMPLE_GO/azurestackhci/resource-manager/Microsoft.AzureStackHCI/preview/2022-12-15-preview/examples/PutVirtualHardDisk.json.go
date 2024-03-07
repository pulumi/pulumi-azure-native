package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewVirtualHardDisk(ctx, "virtualHardDisk", &azurestackhci.VirtualHardDiskArgs{
			DiskSizeGB: pulumi.Float64(32),
			ExtendedLocation: &azurestackhci.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
				Type: pulumi.String("CustomLocation"),
			},
			Location:            pulumi.String("West US2"),
			ResourceGroupName:   pulumi.String("test-rg"),
			VirtualHardDiskName: pulumi.String("test-vhd"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
