package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewP2sVpnGateway(ctx, "p2sVpnGateway", &network.P2sVpnGatewayArgs{
			CustomDnsServers: pulumi.StringArray{
				pulumi.String("1.1.1.1"),
				pulumi.String("2.2.2.2"),
			},
			GatewayName:                 pulumi.String("p2sVpnGateway1"),
			IsRoutingPreferenceInternet: pulumi.Bool(false),
			Location:                    pulumi.String("West US"),
			P2SConnectionConfigurations: network.P2SConnectionConfigurationArray{
				&network.P2SConnectionConfigurationArgs{
					Id:   pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/p2sVpnGateways/p2sVpnGateway1/p2sConnectionConfigurations/P2SConnectionConfig1"),
					Name: pulumi.String("P2SConnectionConfig1"),
					RoutingConfiguration: &network.RoutingConfigurationArgs{
						AssociatedRouteTable: &network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
						},
						PropagatedRouteTables: &network.PropagatedRouteTableArgs{
							Ids: network.SubResourceArray{
								&network.SubResourceArgs{
									Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable1"),
								},
								&network.SubResourceArgs{
									Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable2"),
								},
								&network.SubResourceArgs{
									Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/hubRouteTables/hubRouteTable3"),
								},
							},
							Labels: pulumi.StringArray{
								pulumi.String("label1"),
								pulumi.String("label2"),
							},
						},
						VnetRoutes: &network.VnetRouteArgs{
							StaticRoutes: network.StaticRouteArray{},
						},
					},
					VpnClientAddressPool: &network.AddressSpaceArgs{
						AddressPrefixes: pulumi.StringArray{
							pulumi.String("101.3.0.0/16"),
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualHub: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
			},
			VpnGatewayScaleUnit: pulumi.Int(1),
			VpnServerConfiguration: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnServerConfiguration1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
