package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicyExemption(ctx, "policyExemption", &authorization.PolicyExemptionArgs{
			Description:       pulumi.String("Exempt demo cluster from limit sku"),
			DisplayName:       pulumi.String("Exempt demo cluster"),
			ExemptionCategory: pulumi.String("Waiver"),
			Metadata: pulumi.Any{
				Reason: "Temporary exemption for a expensive VM demo",
			},
			PolicyAssignmentId: pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyAssignments/CostManagement"),
			PolicyDefinitionReferenceIds: pulumi.StringArray{
				pulumi.String("Limit_Skus"),
			},
			PolicyExemptionName: pulumi.String("DemoExpensiveVM"),
			Scope:               pulumi.String("subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/resourceGroups/demoCluster"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
