package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewRule(ctx, "rule", &cdn.RuleArgs{
			Actions: pulumi.Array{
				cdn.DeliveryRuleResponseHeaderAction{
					Name: "ModifyResponseHeader",
					Parameters: cdn.HeaderActionParameters{
						HeaderAction: "Overwrite",
						HeaderName:   "X-CDN",
						TypeName:     "DeliveryRuleHeaderActionParameters",
						Value:        "MSFT",
					},
				},
			},
			Conditions: pulumi.Array{
				cdn.DeliveryRuleRequestMethodCondition{
					Name: "RequestMethod",
					Parameters: cdn.RequestMethodMatchConditionParameters{
						MatchValues: []string{
							"GET",
						},
						NegateCondition: false,
						Operator:        "Equal",
						TypeName:        "DeliveryRuleRequestMethodConditionParameters",
					},
				},
			},
			Order:             pulumi.Int(1),
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			RuleName:          pulumi.String("rule1"),
			RuleSetName:       pulumi.String("ruleSet1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
