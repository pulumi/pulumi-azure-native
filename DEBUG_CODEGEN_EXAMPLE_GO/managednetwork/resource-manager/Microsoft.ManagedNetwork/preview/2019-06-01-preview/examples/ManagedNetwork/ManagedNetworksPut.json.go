package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetwork/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetwork.NewManagedNetwork(ctx, "managedNetwork", &managednetwork.ManagedNetworkArgs{
			Location:           pulumi.String("eastus"),
			ManagedNetworkName: pulumi.String("myManagedNetwork"),
			ResourceGroupName:  pulumi.String("myResourceGroup"),
			Scope: &managednetwork.ScopeArgs{
				ManagementGroups: managednetwork.ResourceIdArray{
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/providers/Microsoft.Management/managementGroups/20000000-0001-0000-0000-000000000000"),
					},
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/providers/Microsoft.Management/managementGroups/20000000-0002-0000-0000-000000000000"),
					},
				},
				Subnets: managednetwork.ResourceIdArray{
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetC/subnets/subnetA"),
					},
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetC/subnets/subnetB"),
					},
				},
				Subscriptions: managednetwork.ResourceIdArray{
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("subscriptionA"),
					},
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("subscriptionB"),
					},
				},
				VirtualNetworks: managednetwork.ResourceIdArray{
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetA"),
					},
					&managednetwork.ResourceIdArgs{
						Id: pulumi.String("/subscriptions/subscriptionC/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/VnetB"),
					},
				},
			},
			Tags: nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
