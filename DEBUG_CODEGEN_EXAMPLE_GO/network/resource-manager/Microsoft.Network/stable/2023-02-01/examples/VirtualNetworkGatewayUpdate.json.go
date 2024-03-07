package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVirtualNetworkGateway(ctx, "virtualNetworkGateway", &network.VirtualNetworkGatewayArgs{
			ActiveActive:           pulumi.Bool(false),
			AllowRemoteVnetTraffic: pulumi.Bool(false),
			AllowVirtualWanTraffic: pulumi.Bool(false),
			BgpSettings: &network.BgpSettingsArgs{
				Asn:               pulumi.Float64(65515),
				BgpPeeringAddress: pulumi.String("10.0.1.30"),
				PeerWeight:        pulumi.Int(0),
			},
			CustomRoutes: &network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{
					pulumi.String("101.168.0.6/32"),
				},
			},
			DisableIPSecReplayProtection:    pulumi.Bool(false),
			EnableBgp:                       pulumi.Bool(false),
			EnableBgpRouteTranslationForNat: pulumi.Bool(false),
			EnableDnsForwarding:             pulumi.Bool(true),
			GatewayType:                     pulumi.String("Vpn"),
			IpConfigurations: network.VirtualNetworkGatewayIPConfigurationArray{
				&network.VirtualNetworkGatewayIPConfigurationArgs{
					Name:                      pulumi.String("gwipconfig1"),
					PrivateIPAllocationMethod: pulumi.String("Dynamic"),
					PublicIPAddress: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/gwpip"),
					},
					Subnet: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet1/subnets/GatewaySubnet"),
					},
				},
			},
			Location: pulumi.String("centralus"),
			NatRules: network.VirtualNetworkGatewayNatRuleTypeArray{
				&network.VirtualNetworkGatewayNatRuleTypeArgs{
					ExternalMappings: network.VpnNatRuleMappingArray{
						&network.VpnNatRuleMappingArgs{
							AddressSpace: pulumi.String("50.0.0.0/24"),
						},
					},
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule1"),
					InternalMappings: network.VpnNatRuleMappingArray{
						&network.VpnNatRuleMappingArgs{
							AddressSpace: pulumi.String("10.10.0.0/24"),
						},
					},
					IpConfigurationId: pulumi.String(""),
					Mode:              pulumi.String("EgressSnat"),
					Name:              pulumi.String("natRule1"),
					Type:              pulumi.String("Static"),
				},
				&network.VirtualNetworkGatewayNatRuleTypeArgs{
					ExternalMappings: network.VpnNatRuleMappingArray{
						&network.VpnNatRuleMappingArgs{
							AddressSpace: pulumi.String("30.0.0.0/24"),
						},
					},
					Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworkGateways/vpngw/natRules/natRule2"),
					InternalMappings: network.VpnNatRuleMappingArray{
						&network.VpnNatRuleMappingArgs{
							AddressSpace: pulumi.String("20.10.0.0/24"),
						},
					},
					IpConfigurationId: pulumi.String(""),
					Mode:              pulumi.String("IngressSnat"),
					Name:              pulumi.String("natRule2"),
					Type:              pulumi.String("Static"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &network.VirtualNetworkGatewaySkuArgs{
				Name: pulumi.String("VpnGw1"),
				Tier: pulumi.String("VpnGw1"),
			},
			VirtualNetworkGatewayName: pulumi.String("vpngw"),
			VpnClientConfiguration: &network.VpnClientConfigurationArgs{
				RadiusServers: network.RadiusServerArray{
					&network.RadiusServerArgs{
						RadiusServerAddress: pulumi.String("10.2.0.0"),
						RadiusServerScore:   pulumi.Float64(20),
						RadiusServerSecret:  pulumi.String("radiusServerSecret"),
					},
				},
				VpnClientProtocols: pulumi.StringArray{
					pulumi.String("OpenVPN"),
				},
				VpnClientRevokedCertificates: network.VpnClientRevokedCertificateArray{},
				VpnClientRootCertificates:    network.VpnClientRootCertificateArray{},
			},
			VpnType: pulumi.String("RouteBased"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
