package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewVpnSite(ctx, "vpnSite", &network.VpnSiteArgs{
			AddressSpace: &network.AddressSpaceArgs{
				AddressPrefixes: pulumi.StringArray{
					pulumi.String("10.0.0.0/16"),
				},
			},
			IsSecuritySite: pulumi.Bool(false),
			Location:       pulumi.String("West US"),
			O365Policy: &network.O365PolicyPropertiesArgs{
				BreakOutCategories: &network.O365BreakOutCategoryPoliciesArgs{
					Allow:    pulumi.Bool(true),
					Default:  pulumi.Bool(false),
					Optimize: pulumi.Bool(true),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			VirtualWan: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualWANs/wan1"),
			},
			VpnSiteLinks: network.VpnSiteLinkArray{
				&network.VpnSiteLinkArgs{
					BgpProperties: &network.VpnLinkBgpSettingsArgs{
						Asn:               pulumi.Float64(1234),
						BgpPeeringAddress: pulumi.String("192.168.0.0"),
					},
					Fqdn:      pulumi.String("link1.vpnsite1.contoso.com"),
					IpAddress: pulumi.String("50.50.50.56"),
					LinkProperties: &network.VpnLinkProviderPropertiesArgs{
						LinkProviderName: pulumi.String("vendor1"),
						LinkSpeedInMbps:  pulumi.Int(0),
					},
					Name: pulumi.String("vpnSiteLink1"),
				},
			},
			VpnSiteName: pulumi.String("vpnSite1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
