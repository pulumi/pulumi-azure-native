package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storagepool/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storagepool.NewIscsiTarget(ctx, "iscsiTarget", &storagepool.IscsiTargetArgs{
			AclMode:         pulumi.String("Dynamic"),
			DiskPoolName:    pulumi.String("myDiskPool"),
			IscsiTargetName: pulumi.String("myIscsiTarget"),
			Luns: []storagepool.IscsiLunArgs{
				{
					ManagedDiskAzureResourceId: pulumi.String("/subscriptions/11111111-1111-1111-1111-111111111111/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/vm-name_DataDisk_1"),
					Name:                       pulumi.String("lun0"),
				},
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			TargetIqn:         pulumi.String("iqn.2005-03.org.iscsi:server1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
