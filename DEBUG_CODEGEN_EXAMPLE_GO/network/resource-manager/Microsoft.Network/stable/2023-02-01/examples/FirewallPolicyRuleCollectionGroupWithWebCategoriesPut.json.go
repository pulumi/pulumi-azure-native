package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewFirewallPolicyRuleCollectionGroup(ctx, "firewallPolicyRuleCollectionGroup", &network.FirewallPolicyRuleCollectionGroupArgs{
			FirewallPolicyName:      pulumi.String("firewallPolicy"),
			Priority:                pulumi.Int(110),
			ResourceGroupName:       pulumi.String("rg1"),
			RuleCollectionGroupName: pulumi.String("ruleCollectionGroup1"),
			RuleCollections: pulumi.Array{
				network.FirewallPolicyFilterRuleCollection{
					Action: network.FirewallPolicyFilterRuleCollectionAction{
						Type: "Deny",
					},
					Name:               "Example-Filter-Rule-Collection",
					RuleCollectionType: "FirewallPolicyFilterRuleCollection",
					Rules: []interface{}{
						network.ApplicationRule{
							Description: "Deny inbound rule",
							Name:        "rule1",
							Protocols: []network.FirewallPolicyRuleApplicationProtocol{
								{
									Port:         443,
									ProtocolType: "Https",
								},
							},
							RuleType: "ApplicationRule",
							SourceAddresses: []string{
								"216.58.216.164",
								"10.0.0.0/24",
							},
							WebCategories: []string{
								"Hacking",
							},
						},
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
