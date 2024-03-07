package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVpnGateway(ctx, "vpnGateway", &network.VpnGatewayArgs{
			BgpSettings: network.BgpSettingsResponse{
				Asn: pulumi.Float64(65515),
				BgpPeeringAddresses: network.IPConfigurationBgpPeeringAddressArray{
					&network.IPConfigurationBgpPeeringAddressArgs{
						CustomBgpIpAddresses: pulumi.StringArray{
							pulumi.String("169.254.21.5"),
						},
						IpconfigurationId: pulumi.String("Instance0"),
					},
					&network.IPConfigurationBgpPeeringAddressArgs{
						CustomBgpIpAddresses: pulumi.StringArray{
							pulumi.String("169.254.21.10"),
						},
						IpconfigurationId: pulumi.String("Instance1"),
					},
				},
				PeerWeight: pulumi.Int(0),
			},
			Connections: []network.VpnConnectionTypeArgs{
				{
					Name: pulumi.String("vpnConnection1"),
					RemoteVpnSite: {
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1"),
					},
					VpnLinkConnections: network.VpnSiteLinkConnectionArray{
						{
							ConnectionBandwidth: pulumi.Int(200),
							EgressNatRules: network.SubResourceArray{
								{
									Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnGateways/gateway1/natRules/nat03"),
								},
							},
							Name:                      pulumi.String("Connection-Link1"),
							SharedKey:                 pulumi.String("key"),
							VpnConnectionProtocolType: pulumi.String("IKEv2"),
							VpnSiteLink: {
								Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnSites/vpnSite1/vpnSiteLinks/siteLink1"),
							},
						},
					},
				},
			},
			EnableBgpRouteTranslationForNat: pulumi.Bool(false),
			GatewayName:                     pulumi.String("gateway1"),
			IsRoutingPreferenceInternet:     pulumi.Bool(false),
			Location:                        pulumi.String("westcentralus"),
			NatRules: []network.VpnGatewayNatRuleArgs{
				{
					ExternalMappings: network.VpnNatRuleMappingArray{
						{
							AddressSpace: pulumi.String("192.168.0.0/26"),
						},
					},
					InternalMappings: network.VpnNatRuleMappingArray{
						{
							AddressSpace: pulumi.String("0.0.0.0/26"),
						},
					},
					IpConfigurationId: pulumi.String(""),
					Mode:              pulumi.String("EgressSnat"),
					Name:              pulumi.String("nat03"),
					Type:              pulumi.String("Static"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualHub: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/virtualHub1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
