package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewFirewallPolicyRuleCollectionGroup(ctx, "firewallPolicyRuleCollectionGroup", &network.FirewallPolicyRuleCollectionGroupArgs{
			FirewallPolicyName:      pulumi.String("firewallPolicy"),
			Priority:                pulumi.Int(100),
			ResourceGroupName:       pulumi.String("rg1"),
			RuleCollectionGroupName: pulumi.String("ruleCollectionGroup1"),
			RuleCollections: pulumi.Array{
				network.FirewallPolicyNatRuleCollection{
					Action: network.FirewallPolicyNatRuleCollectionAction{
						Type: "DNAT",
					},
					Name:               "Example-Nat-Rule-Collection",
					Priority:           100,
					RuleCollectionType: "FirewallPolicyNatRuleCollection",
					Rules: []interface{}{
						network.NatRuleType{
							DestinationAddresses: []string{
								"152.23.32.23",
							},
							DestinationPorts: []string{
								"8080",
							},
							IpProtocols: []network.FirewallPolicyRuleNetworkProtocol{
								"TCP",
								"UDP",
							},
							Name:     "nat-rule1",
							RuleType: "NatRule",
							SourceAddresses: []string{
								"2.2.2.2",
							},
							SourceIpGroups: []interface{}{},
							TranslatedFqdn: "internalhttp.server.net",
							TranslatedPort: "8080",
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
