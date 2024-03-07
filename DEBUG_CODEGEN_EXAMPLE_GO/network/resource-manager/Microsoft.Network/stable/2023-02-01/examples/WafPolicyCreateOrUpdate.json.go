package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/network/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := network.NewWebApplicationFirewallPolicy(ctx, "webApplicationFirewallPolicy", &network.WebApplicationFirewallPolicyArgs{
			CustomRules: network.WebApplicationFirewallCustomRuleArray{
				&network.WebApplicationFirewallCustomRuleArgs{
					Action: pulumi.String("Block"),
					MatchConditions: network.MatchConditionArray{
						&network.MatchConditionArgs{
							MatchValues: pulumi.StringArray{
								pulumi.String("192.168.1.0/24"),
								pulumi.String("10.0.0.0/24"),
							},
							MatchVariables: network.MatchVariableArray{
								&network.MatchVariableArgs{
									VariableName: pulumi.String("RemoteAddr"),
								},
							},
							Operator: pulumi.String("IPMatch"),
						},
					},
					Name:     pulumi.String("Rule1"),
					Priority: pulumi.Int(1),
					RuleType: pulumi.String("MatchRule"),
				},
				&network.WebApplicationFirewallCustomRuleArgs{
					Action: pulumi.String("Block"),
					MatchConditions: network.MatchConditionArray{
						&network.MatchConditionArgs{
							MatchValues: pulumi.StringArray{
								pulumi.String("192.168.1.0/24"),
							},
							MatchVariables: network.MatchVariableArray{
								&network.MatchVariableArgs{
									VariableName: pulumi.String("RemoteAddr"),
								},
							},
							Operator: pulumi.String("IPMatch"),
						},
						&network.MatchConditionArgs{
							MatchValues: pulumi.StringArray{
								pulumi.String("Windows"),
							},
							MatchVariables: network.MatchVariableArray{
								&network.MatchVariableArgs{
									Selector:     pulumi.String("UserAgent"),
									VariableName: pulumi.String("RequestHeaders"),
								},
							},
							Operator: pulumi.String("Contains"),
						},
					},
					Name:     pulumi.String("Rule2"),
					Priority: pulumi.Int(2),
					RuleType: pulumi.String("MatchRule"),
				},
				&network.WebApplicationFirewallCustomRuleArgs{
					Action: pulumi.String("Block"),
					GroupByUserSession: network.GroupByUserSessionArray{
						&network.GroupByUserSessionArgs{
							GroupByVariables: network.GroupByVariableArray{
								&network.GroupByVariableArgs{
									VariableName: pulumi.String("ClientAddr"),
								},
							},
						},
					},
					MatchConditions: network.MatchConditionArray{
						&network.MatchConditionArgs{
							MatchValues: pulumi.StringArray{
								pulumi.String("192.168.1.0/24"),
								pulumi.String("10.0.0.0/24"),
							},
							MatchVariables: network.MatchVariableArray{
								&network.MatchVariableArgs{
									VariableName: pulumi.String("RemoteAddr"),
								},
							},
							NegationConditon: pulumi.Bool(true),
							Operator:         pulumi.String("IPMatch"),
						},
					},
					Name:               pulumi.String("RateLimitRule3"),
					Priority:           pulumi.Int(3),
					RateLimitDuration:  pulumi.String("OneMin"),
					RateLimitThreshold: pulumi.Int(10),
					RuleType:           pulumi.String("RateLimitRule"),
				},
			},
			Location: pulumi.String("WestUs"),
			ManagedRules: &network.ManagedRulesDefinitionArgs{
				Exclusions: network.OwaspCrsExclusionEntryArray{
					&network.OwaspCrsExclusionEntryArgs{
						ExclusionManagedRuleSets: network.ExclusionManagedRuleSetArray{
							&network.ExclusionManagedRuleSetArgs{
								RuleGroups: network.ExclusionManagedRuleGroupArray{
									&network.ExclusionManagedRuleGroupArgs{
										RuleGroupName: pulumi.String("REQUEST-930-APPLICATION-ATTACK-LFI"),
										Rules: network.ExclusionManagedRuleArray{
											&network.ExclusionManagedRuleArgs{
												RuleId: pulumi.String("930120"),
											},
										},
									},
									&network.ExclusionManagedRuleGroupArgs{
										RuleGroupName: pulumi.String("REQUEST-932-APPLICATION-ATTACK-RCE"),
									},
								},
								RuleSetType:    pulumi.String("OWASP"),
								RuleSetVersion: pulumi.String("3.2"),
							},
						},
						MatchVariable:         pulumi.String("RequestArgNames"),
						Selector:              pulumi.String("hello"),
						SelectorMatchOperator: pulumi.String("StartsWith"),
					},
					&network.OwaspCrsExclusionEntryArgs{
						ExclusionManagedRuleSets: network.ExclusionManagedRuleSetArray{
							&network.ExclusionManagedRuleSetArgs{
								RuleGroups:     network.ExclusionManagedRuleGroupArray{},
								RuleSetType:    pulumi.String("OWASP"),
								RuleSetVersion: pulumi.String("3.1"),
							},
						},
						MatchVariable:         pulumi.String("RequestArgNames"),
						Selector:              pulumi.String("hello"),
						SelectorMatchOperator: pulumi.String("EndsWith"),
					},
					&network.OwaspCrsExclusionEntryArgs{
						MatchVariable:         pulumi.String("RequestArgNames"),
						Selector:              pulumi.String("test"),
						SelectorMatchOperator: pulumi.String("StartsWith"),
					},
					&network.OwaspCrsExclusionEntryArgs{
						MatchVariable:         pulumi.String("RequestArgValues"),
						Selector:              pulumi.String("test"),
						SelectorMatchOperator: pulumi.String("StartsWith"),
					},
				},
				ManagedRuleSets: network.ManagedRuleSetArray{
					&network.ManagedRuleSetArgs{
						RuleGroupOverrides: network.ManagedRuleGroupOverrideArray{
							&network.ManagedRuleGroupOverrideArgs{
								RuleGroupName: pulumi.String("REQUEST-931-APPLICATION-ATTACK-RFI"),
								Rules: network.ManagedRuleOverrideArray{
									&network.ManagedRuleOverrideArgs{
										Action: pulumi.String("Log"),
										RuleId: pulumi.String("931120"),
										State:  pulumi.String("Enabled"),
									},
									&network.ManagedRuleOverrideArgs{
										Action: pulumi.String("AnomalyScoring"),
										RuleId: pulumi.String("931130"),
										State:  pulumi.String("Disabled"),
									},
								},
							},
						},
						RuleSetType:    pulumi.String("OWASP"),
						RuleSetVersion: pulumi.String("3.2"),
					},
				},
			},
			PolicyName: pulumi.String("Policy1"),
			PolicySettings: &network.PolicySettingsArgs{
				LogScrubbing: &network.PolicySettingsLogScrubbingArgs{
					ScrubbingRules: network.WebApplicationFirewallScrubbingRulesArray{
						&network.WebApplicationFirewallScrubbingRulesArgs{
							MatchVariable:         pulumi.String("RequestArgNames"),
							Selector:              pulumi.String("test"),
							SelectorMatchOperator: pulumi.String("Equals"),
							State:                 pulumi.String("Enabled"),
						},
						&network.WebApplicationFirewallScrubbingRulesArgs{
							MatchVariable:         pulumi.String("RequestIPAddress"),
							SelectorMatchOperator: pulumi.String("EqualsAny"),
							State:                 pulumi.String("Enabled"),
						},
					},
					State: pulumi.String("Enabled"),
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
