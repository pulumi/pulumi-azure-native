package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewNetworkTapRule(ctx, "networkTapRule", &managednetworkfabric.NetworkTapRuleArgs{
			Annotation:        pulumi.String("annotation"),
			ConfigurationType: pulumi.String("File"),
			DynamicMatchConfigurations: managednetworkfabric.CommonDynamicMatchConfigurationArray{
				&managednetworkfabric.CommonDynamicMatchConfigurationArgs{
					IpGroups: managednetworkfabric.IpGroupPropertiesArray{
						&managednetworkfabric.IpGroupPropertiesArgs{
							IpAddressType: pulumi.String("IPv4"),
							IpPrefixes: pulumi.StringArray{
								pulumi.String("10.10.10.10/30"),
							},
							Name: pulumi.String("example-ipGroup1"),
						},
					},
					PortGroups: managednetworkfabric.PortGroupPropertiesArray{
						&managednetworkfabric.PortGroupPropertiesArgs{
							Name: pulumi.String("example-portGroup1"),
							Ports: pulumi.StringArray{
								pulumi.String("100-200"),
							},
						},
						&managednetworkfabric.PortGroupPropertiesArgs{
							Name: pulumi.String("example-portGroup2"),
							Ports: pulumi.StringArray{
								pulumi.String("900"),
								pulumi.String("1000-2000"),
							},
						},
					},
					VlanGroups: managednetworkfabric.VlanGroupPropertiesArray{
						&managednetworkfabric.VlanGroupPropertiesArgs{
							Name: pulumi.String("exmaple-vlanGroup"),
							Vlans: pulumi.StringArray{
								pulumi.String("10"),
								pulumi.String("100-200"),
							},
						},
					},
				},
			},
			Location: pulumi.String("eastus"),
			MatchConfigurations: managednetworkfabric.NetworkTapRuleMatchConfigurationArray{
				&managednetworkfabric.NetworkTapRuleMatchConfigurationArgs{
					Actions: managednetworkfabric.NetworkTapRuleActionArray{
						&managednetworkfabric.NetworkTapRuleActionArgs{
							DestinationId:          pulumi.String("/subscriptions/1234ABCD-0A1B-1234-5678-123456ABCDEF/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/neighborGroups/example-neighborGroup"),
							IsTimestampEnabled:     pulumi.String("True"),
							MatchConfigurationName: pulumi.String("match1"),
							Truncate:               pulumi.String("100"),
							Type:                   pulumi.String("Drop"),
						},
					},
					IpAddressType: pulumi.String("IPv4"),
					MatchConditions: managednetworkfabric.NetworkTapRuleMatchConditionArray{
						&managednetworkfabric.NetworkTapRuleMatchConditionArgs{
							EncapsulationType: pulumi.String("None"),
							IpCondition: &managednetworkfabric.IpMatchConditionArgs{
								IpGroupNames: pulumi.StringArray{
									pulumi.String("example-ipGroup"),
								},
								IpPrefixValues: pulumi.StringArray{
									pulumi.String("10.10.10.10/20"),
								},
								PrefixType: pulumi.String("Prefix"),
								Type:       pulumi.String("SourceIP"),
							},
							PortCondition: &managednetworkfabric.PortConditionArgs{
								Layer4Protocol: pulumi.String("TCP"),
								PortGroupNames: pulumi.StringArray{
									pulumi.String("example-portGroup1"),
								},
								PortType: pulumi.String("SourcePort"),
								Ports: pulumi.StringArray{
									pulumi.String("100"),
								},
							},
							ProtocolTypes: pulumi.StringArray{
								pulumi.String("TCP"),
							},
							VlanMatchCondition: &managednetworkfabric.VlanMatchConditionArgs{
								InnerVlans: pulumi.StringArray{
									pulumi.String("11-20"),
								},
								VlanGroupNames: pulumi.StringArray{
									pulumi.String("exmaple-vlanGroup"),
								},
								Vlans: pulumi.StringArray{
									pulumi.String("10"),
								},
							},
						},
					},
					MatchConfigurationName: pulumi.String("config1"),
					SequenceNumber:         pulumi.Float64(10),
				},
			},
			NetworkTapRuleName:       pulumi.String("example-tapRule"),
			PollingIntervalInSeconds: pulumi.Int(30),
			ResourceGroupName:        pulumi.String("example-rg"),
			Tags: pulumi.StringMap{
				"keyID": pulumi.String("keyValue"),
			},
			TapRulesUrl: pulumi.String("https://microsoft.com/a"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
