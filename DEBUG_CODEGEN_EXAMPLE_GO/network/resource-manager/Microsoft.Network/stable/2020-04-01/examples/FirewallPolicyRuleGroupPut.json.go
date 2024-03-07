package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewFirewallPolicyRuleGroup(ctx, "firewallPolicyRuleGroup", &network.FirewallPolicyRuleGroupArgs{
			FirewallPolicyName: pulumi.String("firewallPolicy"),
			Priority:           pulumi.Int(110),
			ResourceGroupName:  pulumi.String("rg1"),
			RuleGroupName:      pulumi.String("ruleGroup1"),
			Rules: pulumi.Array{
				network.FirewallPolicyFilterRule{
					Action: network.FirewallPolicyFilterRuleAction{
						Type: "Deny",
					},
					Name: "Example-Filter-Rule",
					RuleConditions: []interface{}{
						network.NetworkRuleCondition{
							DestinationAddresses: []string{
								"*",
							},
							DestinationPorts: []string{
								"*",
							},
							IpProtocols: []network.FirewallPolicyRuleConditionNetworkProtocol{
								"TCP",
							},
							Name:              "network-condition1",
							RuleConditionType: "NetworkRuleCondition",
							SourceAddresses: []string{
								"10.1.25.0/24",
							},
						},
					},
					RuleType: "FirewallPolicyFilterRule",
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
