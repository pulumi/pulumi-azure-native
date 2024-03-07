package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetwork.NewManagedNetworkGroup(ctx, "managedNetworkGroup", &managednetwork.ManagedNetworkGroupArgs{
			ManagedNetworkGroupName: pulumi.String("myManagedNetworkGroup1"),
			ManagedNetworkName:      pulumi.String("myManagedNetwork"),
			ManagementGroups:        managednetwork.ResourceIdArray{},
			ResourceGroupName:       pulumi.String("myResourceGroup"),
			Subnets: managednetwork.ResourceIdArray{
				&managednetwork.ResourceIdArgs{
					Id: pulumi.String("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA/subnets/subnetA"),
				},
			},
			Subscriptions: managednetwork.ResourceIdArray{},
			VirtualNetworks: managednetwork.ResourceIdArray{
				&managednetwork.ResourceIdArgs{
					Id: pulumi.String("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA"),
				},
				&managednetwork.ResourceIdArgs{
					Id: pulumi.String("/subscriptionB/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetB"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
