package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewNatGateway(ctx, "natGateway", &network.NatGatewayArgs{
			Location:       pulumi.String("westus"),
			NatGatewayName: pulumi.String("test-natgateway"),
			PublicIpAddresses: network.SubResourceArray{
				&network.SubResourceArgs{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/PublicIpAddress1"),
				},
			},
			PublicIpPrefixes: network.SubResourceArray{
				&network.SubResourceArgs{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPPrefixes/PublicIpPrefix1"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &network.NatGatewaySkuArgs{
				Name: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
