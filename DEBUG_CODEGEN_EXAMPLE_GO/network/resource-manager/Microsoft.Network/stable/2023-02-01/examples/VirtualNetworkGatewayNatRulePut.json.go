package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualNetworkGatewayNatRule(ctx, "virtualNetworkGatewayNatRule", &network.VirtualNetworkGatewayNatRuleArgs{
			ExternalMappings: []network.VpnNatRuleMappingArgs{
				{
					AddressSpace: pulumi.String("192.168.21.0/24"),
					PortRange:    pulumi.String("300-400"),
				},
			},
			InternalMappings: []network.VpnNatRuleMappingArgs{
				{
					AddressSpace: pulumi.String("10.4.0.0/24"),
					PortRange:    pulumi.String("200-300"),
				},
			},
			IpConfigurationId:         pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/gateway1/ipConfigurations/default"),
			Mode:                      pulumi.String("EgressSnat"),
			NatRuleName:               pulumi.String("natRule1"),
			ResourceGroupName:         pulumi.String("rg1"),
			Type:                      pulumi.String("Static"),
			VirtualNetworkGatewayName: pulumi.String("gateway1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
