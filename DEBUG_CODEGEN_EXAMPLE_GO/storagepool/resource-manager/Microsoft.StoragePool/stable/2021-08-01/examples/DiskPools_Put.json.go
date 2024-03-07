package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagepool/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagepool.NewDiskPool(ctx, "diskPool", &storagepool.DiskPoolArgs{
			AvailabilityZones: pulumi.StringArray{
				pulumi.String("1"),
			},
			DiskPoolName: pulumi.String("myDiskPool"),
			Disks: storagepool.DiskArray{
				&storagepool.DiskArgs{
					Id: pulumi.String("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_0"),
				},
				&storagepool.DiskArgs{
					Id: pulumi.String("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
				},
			},
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &storagepool.SkuArgs{
				Name: pulumi.String("Basic_V1"),
				Tier: pulumi.String("Basic"),
			},
			SubnetId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/myvnet/subnets/mysubnet"),
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
