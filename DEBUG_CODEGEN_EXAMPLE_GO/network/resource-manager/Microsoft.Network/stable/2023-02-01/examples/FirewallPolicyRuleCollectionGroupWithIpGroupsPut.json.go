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
						network.NetworkRule{
							DestinationIpGroups: []string{
								"/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups2",
							},
							DestinationPorts: []string{
								"*",
							},
							IpProtocols: []network.FirewallPolicyRuleNetworkProtocol{
								"TCP",
							},
							Name:     "network-1",
							RuleType: "NetworkRule",
							SourceIpGroups: []string{
								"/subscriptions/subid/providers/Microsoft.Network/resourceGroup/rg1/ipGroups/ipGroups1",
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
