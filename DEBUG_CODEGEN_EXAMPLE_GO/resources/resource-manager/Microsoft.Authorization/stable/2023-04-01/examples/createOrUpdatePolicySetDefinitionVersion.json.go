package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/authorization/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := authorization.NewPolicySetDefinitionVersion(ctx, "policySetDefinitionVersion", &authorization.PolicySetDefinitionVersionArgs{
			Description: pulumi.String("Policies to enforce low cost storage SKUs"),
			DisplayName: pulumi.String("Cost Management"),
			Metadata: pulumi.Any{
				Category: "Cost Management",
			},
			Parameters: authorization.ParameterDefinitionsValueMap{
				"namePrefix": &authorization.ParameterDefinitionsValueArgs{
					DefaultValue: pulumi.Any("myPrefix"),
					Metadata: &authorization.ParameterDefinitionsValueMetadataArgs{
						DisplayName: pulumi.String("Prefix to enforce on resource names"),
					},
					Type: pulumi.String("String"),
				},
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
					PolicyDefinitionId:          pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/7433c107-6db4-4ad1-b57a-a76dce0154a1"),
					PolicyDefinitionReferenceId: pulumi.String("Limit_Skus"),
				},
				{
					Parameters: {
						"prefix": {
							Value: pulumi.Any("[parameters('namePrefix')]"),
						},
						"suffix": {
							Value: pulumi.Any("-LC"),
						},
					},
					PolicyDefinitionId:          pulumi.String("/subscriptions/ae640e6b-ba3e-4256-9d62-2993eecfa6f2/providers/Microsoft.Authorization/policyDefinitions/ResourceNaming"),
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
