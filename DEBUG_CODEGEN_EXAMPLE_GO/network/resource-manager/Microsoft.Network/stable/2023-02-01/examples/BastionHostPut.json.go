package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewBastionHost(ctx, "bastionHost", &network.BastionHostArgs{
			BastionHostName: pulumi.String("bastionhosttenant"),
			IpConfigurations: []network.BastionHostIPConfigurationArgs{
				{
					Name: pulumi.String("bastionHostIpConfiguration"),
					PublicIPAddress: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
					},
					Subnet: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/BastionHostSubnet"),
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
