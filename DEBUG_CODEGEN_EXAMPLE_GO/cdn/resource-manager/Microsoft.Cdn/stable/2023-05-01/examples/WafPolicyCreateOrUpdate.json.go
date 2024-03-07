package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewPolicy(ctx, "policy", &cdn.PolicyArgs{
			CustomRules: &cdn.CustomRuleListArgs{
				Rules: cdn.CustomRuleArray{
					&cdn.CustomRuleArgs{
						Action:       pulumi.String("Block"),
						EnabledState: pulumi.String("Enabled"),
						MatchConditions: cdn.MatchConditionArray{
							&cdn.MatchConditionArgs{
								MatchValue: pulumi.StringArray{
									pulumi.String("CH"),
								},
								MatchVariable:   pulumi.String("RemoteAddr"),
								NegateCondition: pulumi.Bool(false),
								Operator:        pulumi.String("GeoMatch"),
								Transforms:      pulumi.StringArray{},
							},
							&cdn.MatchConditionArgs{
								MatchValue: pulumi.StringArray{
									pulumi.String("windows"),
								},
								MatchVariable:   pulumi.String("RequestHeader"),
								NegateCondition: pulumi.Bool(false),
								Operator:        pulumi.String("Contains"),
								Selector:        pulumi.String("UserAgent"),
								Transforms:      pulumi.StringArray{},
							},
							&cdn.MatchConditionArgs{
								MatchValue: pulumi.StringArray{
									pulumi.String("<?php"),
									pulumi.String("?>"),
								},
								MatchVariable:   pulumi.String("QueryString"),
								NegateCondition: pulumi.Bool(false),
								Operator:        pulumi.String("Contains"),
								Selector:        pulumi.String("search"),
								Transforms: pulumi.StringArray{
									pulumi.String("UrlDecode"),
									pulumi.String("Lowercase"),
								},
							},
						},
						Name:     pulumi.String("CustomRule1"),
						Priority: pulumi.Int(2),
					},
				},
			},
			Location: pulumi.String("global"),
			ManagedRules: &cdn.ManagedRuleSetListArgs{
				ManagedRuleSets: cdn.ManagedRuleSetArray{
					&cdn.ManagedRuleSetArgs{
						RuleGroupOverrides: cdn.ManagedRuleGroupOverrideArray{
							&cdn.ManagedRuleGroupOverrideArgs{
								RuleGroupName: pulumi.String("Group1"),
								Rules: cdn.ManagedRuleOverrideArray{
									&cdn.ManagedRuleOverrideArgs{
										Action:       pulumi.String("Redirect"),
										EnabledState: pulumi.String("Enabled"),
										RuleId:       pulumi.String("GROUP1-0001"),
									},
									&cdn.ManagedRuleOverrideArgs{
										EnabledState: pulumi.String("Disabled"),
										RuleId:       pulumi.String("GROUP1-0002"),
									},
								},
							},
						},
						RuleSetType:    pulumi.String("DefaultRuleSet"),
						RuleSetVersion: pulumi.String("preview-1.0"),
					},
				},
			},
			PolicyName: pulumi.String("MicrosoftCdnWafPolicy"),
			PolicySettings: &cdn.PolicySettingsArgs{
				DefaultCustomBlockResponseBody:       pulumi.String("PGh0bWw+CjxoZWFkZXI+PHRpdGxlPkhlbGxvPC90aXRsZT48L2hlYWRlcj4KPGJvZHk+CkhlbGxvIHdvcmxkCjwvYm9keT4KPC9odG1sPg=="),
				DefaultCustomBlockResponseStatusCode: pulumi.Int(200),
				DefaultRedirectUrl:                   pulumi.String("http://www.bing.com"),
			},
			RateLimitRules: &cdn.RateLimitRuleListArgs{
				Rules: cdn.RateLimitRuleArray{
					&cdn.RateLimitRuleArgs{
						Action:       pulumi.String("Block"),
						EnabledState: pulumi.String("Enabled"),
						MatchConditions: cdn.MatchConditionArray{
							&cdn.MatchConditionArgs{
								MatchValue: pulumi.StringArray{
									pulumi.String("192.168.1.0/24"),
									pulumi.String("10.0.0.0/24"),
								},
								MatchVariable:   pulumi.String("RemoteAddr"),
								NegateCondition: pulumi.Bool(false),
								Operator:        pulumi.String("IPMatch"),
								Transforms:      pulumi.StringArray{},
							},
						},
						Name:                       pulumi.String("RateLimitRule1"),
						Priority:                   pulumi.Int(1),
						RateLimitDurationInMinutes: pulumi.Int(0),
						RateLimitThreshold:         pulumi.Int(1000),
					},
				},
			},
			ResourceGroupName: pulumi.String("rg1"),
			Sku: &cdn.SkuArgs{
				Name: pulumi.String("Standard_Microsoft"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
