package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicyAssignment(ctx, "policyAssignment", &authorization.PolicyAssignmentArgs{
			Description:     pulumi.String("Force resource names to begin with given DeptA and end with -LC"),
			DisplayName:     pulumi.String("Enforce resource naming rules"),
			EnforcementMode: pulumi.String("Default"),
			Identity: &authorization.IdentityArgs{
				Type: authorization.ResourceIdentityTypeSystemAssigned,
			},
			Location: pulumi.String("eastus"),
			Metadata: pulumi.Any{
				AssignedBy: "Foo Bar",
			},
			Parameters: authorization.ParameterValuesValueMap{
				"prefix": &authorization.ParameterValuesValueArgs{
					Value: pulumi.Any("DeptA"),
				},
				"suffix": &authorization.ParameterValuesValueArgs{
					Value: pulumi.Any("-LC"),
				},
			},
			PolicyAssignmentName: pulumi.String("EnforceNaming"),
			PolicyDefinitionId:   pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
			Scope:                pulumi.String("subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
