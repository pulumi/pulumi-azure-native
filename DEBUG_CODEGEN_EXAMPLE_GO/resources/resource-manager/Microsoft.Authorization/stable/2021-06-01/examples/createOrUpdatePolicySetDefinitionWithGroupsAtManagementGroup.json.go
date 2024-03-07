package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicySetDefinitionAtManagementGroup(ctx, "policySetDefinitionAtManagementGroup", &authorization.PolicySetDefinitionAtManagementGroupArgs{
			Description:       pulumi.String("Policies to enforce low cost storage SKUs"),
			DisplayName:       pulumi.String("Cost Management"),
			ManagementGroupId: pulumi.String("MyManagementGroup"),
			Metadata: pulumi.Any{
				Category: "Cost Management",
			},
			PolicyDefinitionGroups: authorization.PolicyDefinitionGroupArray{
				&authorization.PolicyDefinitionGroupArgs{
					Description: pulumi.String("Policies designed to control spend within a subscription."),
					DisplayName: pulumi.String("Cost Management Policies"),
					Name:        pulumi.String("CostSaving"),
				},
				&authorization.PolicyDefinitionGroupArgs{
					Description: pulumi.String("Policies that help enforce resource organization standards within a subscription."),
					DisplayName: pulumi.String("Organizational Policies"),
					Name:        pulumi.String("Organizational"),
				},
			},
			PolicyDefinitions: authorization.PolicyDefinitionReferenceArray{
				&authorization.PolicyDefinitionReferenceArgs{
					GroupNames: pulumi.StringArray{
						pulumi.String("CostSaving"),
					},
					Parameters: authorization.ParameterValuesValueMap{
						"listOfAllowedSKUs": &authorization.ParameterValuesValueArgs{
							Value: pulumi.Any{
								"Standard_GRS",
								"Standard_LRS",
							},
						},
					},
					PolicyDefinitionId:          pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
					PolicyDefinitionReferenceId: pulumi.String("Limit_Skus"),
				},
				&authorization.PolicyDefinitionReferenceArgs{
					GroupNames: pulumi.StringArray{
						pulumi.String("Organizational"),
					},
					Parameters: authorization.ParameterValuesValueMap{
						"prefix": &authorization.ParameterValuesValueArgs{
							Value: pulumi.Any("DeptA"),
						},
						"suffix": &authorization.ParameterValuesValueArgs{
							Value: pulumi.Any("-LC"),
						},
					},
					PolicyDefinitionId:          pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
					PolicyDefinitionReferenceId: pulumi.String("Resource_Naming"),
				},
			},
			PolicySetDefinitionName: pulumi.String("CostManagement"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
