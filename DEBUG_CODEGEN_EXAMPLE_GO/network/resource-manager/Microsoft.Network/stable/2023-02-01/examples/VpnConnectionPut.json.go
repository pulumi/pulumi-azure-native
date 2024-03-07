package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVpnConnection(ctx, "vpnConnection", &network.VpnConnectionArgs{
			ConnectionName: pulumi.String("vpnConnection1"),
			GatewayName:    pulumi.String("gateway1"),
			RemoteVpnSite: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			RoutingConfiguration: &network.RoutingConfigurationArgs{
				AssociatedRouteTable: &network.SubResourceArgs{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
				},
				InboundRouteMap: &network.SubResourceArgs{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap1"),
				},
				OutboundRouteMap: &network.SubResourceArgs{
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1/routeMaps/routeMap2"),
				},
				PropagatedRouteTables: &network.PropagatedRouteTableArgs{
					Ids: network.SubResourceArray{
						&network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable1"),
						},
						&network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable2"),
						},
						&network.SubResourceArgs{
							Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1/hubRouteTables/hubRouteTable3"),
						},
					},
					Labels: pulumi.StringArray{
						pulumi.String("label1"),
						pulumi.String("label2"),
					},
				},
			},
			TrafficSelectorPolicies: network.TrafficSelectorPolicyArray{},
			VpnLinkConnections: network.VpnSiteLinkConnectionArray{
				&network.VpnSiteLinkConnectionArgs{
					ConnectionBandwidth:            pulumi.Int(200),
					Name:                           pulumi.String("Connection-Link1"),
					SharedKey:                      pulumi.String("key"),
					UsePolicyBasedTrafficSelectors: pulumi.Bool(false),
					VpnConnectionProtocolType:      pulumi.String("IKEv2"),
					VpnLinkConnectionMode:          pulumi.String("Default"),
					VpnSiteLink: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
