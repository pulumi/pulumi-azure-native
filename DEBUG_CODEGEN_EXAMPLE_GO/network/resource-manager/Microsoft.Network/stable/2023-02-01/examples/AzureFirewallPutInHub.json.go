package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewAzureFirewall(ctx, "azureFirewall", &network.AzureFirewallArgs{
			AzureFirewallName: pulumi.String("azurefirewall"),
			FirewallPolicy: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/firewallPolicies/policy1"),
			},
			HubIPAddresses: &network.HubIPAddressesArgs{
				PublicIPs: &network.HubPublicIPAddressesArgs{
					Addresses: network.AzureFirewallPublicIPAddressArray{},
					Count:     pulumi.Int(1),
				},
			},
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &network.AzureFirewallSkuArgs{
				Name: pulumi.String("AZFW_Hub"),
				Tier: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			ThreatIntelMode: pulumi.String("Alert"),
			VirtualHub: &network.SubResourceArgs{
				Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1"),
			},
			Zones: pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
