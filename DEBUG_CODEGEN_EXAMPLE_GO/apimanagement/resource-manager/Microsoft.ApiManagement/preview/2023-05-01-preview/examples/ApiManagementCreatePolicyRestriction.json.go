package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewPolicyRestriction(ctx, "policyRestriction", &apimanagement.PolicyRestrictionArgs{
			PolicyRestrictionId: pulumi.String("policyRestriction1"),
			RequireBase:         pulumi.String("true"),
			ResourceGroupName:   pulumi.String("rg1"),
			Scope:               pulumi.String("Sample Path to the policy document."),
			ServiceName:         pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
