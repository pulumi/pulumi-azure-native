package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNetworkProfile(ctx, "networkProfile", &network.NetworkProfileArgs{
			ContainerNetworkInterfaceConfigurations: network.ContainerNetworkInterfaceConfigurationArray{
				&network.ContainerNetworkInterfaceConfigurationArgs{
					IpConfigurations: network.IPConfigurationProfileArray{
						&network.IPConfigurationProfileArgs{
							Name: pulumi.String("ipconfig1"),
							Subnet: &network.SubnetTypeArgs{
								Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/networkProfileVnet/subnets/networkProfileSubnet1"),
							},
						},
					},
					Name: pulumi.String("eth1"),
				},
			},
			Location:           pulumi.String("westus"),
			NetworkProfileName: pulumi.String("networkProfile1"),
			ResourceGroupName:  pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
