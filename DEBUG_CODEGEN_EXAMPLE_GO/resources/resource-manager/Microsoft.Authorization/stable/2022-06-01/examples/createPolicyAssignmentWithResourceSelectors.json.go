package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicyAssignment(ctx, "policyAssignment", &authorization.PolicyAssignmentArgs{
			Description: pulumi.String("Limit the resource location and resource SKU"),
			DisplayName: pulumi.String("Limit the resource location and resource SKU"),
			Metadata: pulumi.Any{
				AssignedBy: "Special Someone",
			},
			PolicyAssignmentName: pulumi.String("CostManagement"),
			PolicyDefinitionId:   pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policySetDefinitions/CostManagement"),
			ResourceSelectors: []authorization.ResourceSelectorArgs{
				{
					Name: pulumi.String("SDPRegions"),
					Selectors: authorization.SelectorArray{
						{
							In: pulumi.StringArray{
								pulumi.String("eastus2euap"),
								pulumi.String("centraluseuap"),
							},
							Kind: pulumi.String("resourceLocation"),
						},
					},
				},
			},
			Scope: pulumi.String("subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
