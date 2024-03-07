package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewAzureFirewall(ctx, "azureFirewall", &network.AzureFirewallArgs{
			AdditionalProperties: pulumi.StringMap{
				"key1": pulumi.String("value1"),
				"key2": pulumi.String("value2"),
			},
			ApplicationRuleCollections: network.AzureFirewallApplicationRuleCollectionArray{
				&network.AzureFirewallApplicationRuleCollectionArgs{
					Action: &network.AzureFirewallRCActionArgs{
						Type: pulumi.String("Deny"),
					},
					Name:     pulumi.String("apprulecoll"),
					Priority: pulumi.Int(110),
					Rules: network.AzureFirewallApplicationRuleArray{
						&network.AzureFirewallApplicationRuleArgs{
							Description: pulumi.String("Deny inbound rule"),
							Name:        pulumi.String("rule1"),
							Protocols: network.AzureFirewallApplicationRuleProtocolArray{
								&network.AzureFirewallApplicationRuleProtocolArgs{
									Port:         pulumi.Int(443),
									ProtocolType: pulumi.String("Https"),
								},
							},
							SourceAddresses: pulumi.StringArray{
								pulumi.String("216.58.216.164"),
								pulumi.String("10.0.0.0/24"),
							},
							TargetFqdns: pulumi.StringArray{
								pulumi.String("www.test.com"),
							},
						},
					},
				},
			},
			AzureFirewallName: pulumi.String("azurefirewall"),
			IpConfigurations: network.AzureFirewallIPConfigurationArray{
				&network.AzureFirewallIPConfigurationArgs{
					Name: pulumi.String("azureFirewallIpConfiguration"),
					PublicIPAddress: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/publicIPAddresses/pipName"),
					},
					Subnet: &network.SubResourceArgs{
						Id: pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualNetworks/vnet2/subnets/AzureFirewallSubnet"),
					},
				},
			},
			Location: pulumi.String("West US"),
			NatRuleCollections: network.AzureFirewallNatRuleCollectionArray{
				&network.AzureFirewallNatRuleCollectionArgs{
					Action: &network.AzureFirewallNatRCActionArgs{
						Type: pulumi.String("Dnat"),
					},
					Name:     pulumi.String("natrulecoll"),
					Priority: pulumi.Int(112),
					Rules: network.AzureFirewallNatRuleArray{
						&network.AzureFirewallNatRuleArgs{
							Description: pulumi.String("D-NAT all outbound web traffic for inspection"),
							DestinationAddresses: pulumi.StringArray{
								pulumi.String("1.2.3.4"),
							},
							DestinationPorts: pulumi.StringArray{
								pulumi.String("443"),
							},
							Name: pulumi.String("DNAT-HTTPS-traffic"),
							Protocols: pulumi.StringArray{
								pulumi.String("TCP"),
							},
							SourceAddresses: pulumi.StringArray{
								pulumi.String("*"),
							},
							TranslatedAddress: pulumi.String("1.2.3.5"),
							TranslatedPort:    pulumi.String("8443"),
						},
						&network.AzureFirewallNatRuleArgs{
							Description: pulumi.String("D-NAT all inbound web traffic for inspection"),
							DestinationAddresses: pulumi.StringArray{
								pulumi.String("1.2.3.4"),
							},
							DestinationPorts: pulumi.StringArray{
								pulumi.String("80"),
							},
							Name: pulumi.String("DNAT-HTTP-traffic-With-FQDN"),
							Protocols: pulumi.StringArray{
								pulumi.String("TCP"),
							},
							SourceAddresses: pulumi.StringArray{
								pulumi.String("*"),
							},
							TranslatedFqdn: pulumi.String("internalhttpserver"),
							TranslatedPort: pulumi.String("880"),
						},
					},
				},
			},
			NetworkRuleCollections: network.AzureFirewallNetworkRuleCollectionArray{
				&network.AzureFirewallNetworkRuleCollectionArgs{
					Action: &network.AzureFirewallRCActionArgs{
						Type: pulumi.String("Deny"),
					},
					Name:     pulumi.String("netrulecoll"),
					Priority: pulumi.Int(112),
					Rules: network.AzureFirewallNetworkRuleArray{
						&network.AzureFirewallNetworkRuleArgs{
							Description: pulumi.String("Block traffic based on source IPs and ports"),
							DestinationAddresses: pulumi.StringArray{
								pulumi.String("*"),
							},
							DestinationPorts: pulumi.StringArray{
								pulumi.String("443-444"),
								pulumi.String("8443"),
							},
							Name: pulumi.String("L4-traffic"),
							Protocols: pulumi.StringArray{
								pulumi.String("TCP"),
							},
							SourceAddresses: pulumi.StringArray{
								pulumi.String("192.168.1.1-192.168.1.12"),
								pulumi.String("10.1.4.12-10.1.4.255"),
							},
						},
						&network.AzureFirewallNetworkRuleArgs{
							Description: pulumi.String("Block traffic based on source IPs and ports to amazon"),
							DestinationFqdns: pulumi.StringArray{
								pulumi.String("www.amazon.com"),
							},
							DestinationPorts: pulumi.StringArray{
								pulumi.String("443-444"),
								pulumi.String("8443"),
							},
							Name: pulumi.String("L4-traffic-with-FQDN"),
							Protocols: pulumi.StringArray{
								pulumi.String("TCP"),
							},
							SourceAddresses: pulumi.StringArray{
								pulumi.String("10.2.4.12-10.2.4.255"),
							},
						},
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &network.AzureFirewallSkuArgs{
				Name: pulumi.String("AZFW_VNet"),
				Tier: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
			ThreatIntelMode: pulumi.String("Alert"),
			Zones:           pulumi.StringArray{},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
