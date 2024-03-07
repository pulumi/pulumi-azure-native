package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicySetDefinitionVersionAtManagementGroup(ctx, "policySetDefinitionVersionAtManagementGroup", &authorization.PolicySetDefinitionVersionAtManagementGroupArgs{
			Description:         pulumi.String("Policies to enforce low cost storage SKUs"),
			DisplayName:         pulumi.String("Cost Management"),
			ManagementGroupName: pulumi.String("MyManagementGroup"),
			Metadata: pulumi.Any{
				Category: "Cost Management",
			},
			PolicyDefinitionVersion: pulumi.String("1.2.1"),
			PolicyDefinitions: []authorization.PolicyDefinitionReferenceArgs{
				{
					Parameters: {
						"listOfAllowedSKUs": {
							Value: pulumi.Any{
								"Standard_GRS",
								"Standard_LRS",
							},
						},
					},
					PolicyDefinitionId:          pulumi.String("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
					PolicyDefinitionReferenceId: pulumi.String("Limit_Skus"),
				},
				{
					Parameters: {
						"prefix": {
							Value: pulumi.Any("DeptA"),
						},
						"suffix": {
							Value: pulumi.Any("-LC"),
						},
					},
					PolicyDefinitionId:          pulumi.String("/providers/Microsoft.Management/managementgroups/MyManagementGroup/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
					PolicyDefinitionReferenceId: pulumi.String("Resource_Naming"),
				},
			},
			PolicySetDefinitionName: pulumi.String("CostManagement"),
			Version:                 pulumi.String("1.2.1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
