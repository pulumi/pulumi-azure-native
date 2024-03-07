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
				network.FirewallPolicyFilterRuleCollection{
					Action: network.FirewallPolicyFilterRuleCollectionAction{
						Type: "Deny",
					},
					Name:               "Example-Filter-Rule-Collection",
					Priority:           100,
					RuleCollectionType: "FirewallPolicyFilterRuleCollection",
					Rules: []interface{}{
						network.NetworkRule{
							DestinationAddresses: []string{
								"*",
							},
							DestinationPorts: []string{
								"*",
							},
							IpProtocols: []network.FirewallPolicyRuleNetworkProtocol{
								"TCP",
							},
							Name:     "network-rule1",
							RuleType: "NetworkRule",
							SourceAddresses: []string{
								"10.1.25.0/24",
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
