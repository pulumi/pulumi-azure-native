package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualHubIpConfiguration(ctx, "virtualHubIpConfiguration", &network.VirtualHubIpConfigurationArgs{
			IpConfigName:      pulumi.String("ipconfig1"),
			ResourceGroupName: pulumi.String("rg1"),
			Subnet: &network.SubnetTypeArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/subnet1"),
			},
			VirtualHubName: pulumi.String("hub1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
