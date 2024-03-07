package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cdn/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cdn.NewRuleSet(ctx, "ruleSet", &cdn.RuleSetArgs{
			ProfileName:       pulumi.String("profile1"),
			ResourceGroupName: pulumi.String("RG"),
			RuleSetName:       pulumi.String("ruleSet1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
