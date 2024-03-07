package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewBlueprint(ctx, "blueprint", &blueprint.BlueprintArgs{
			BlueprintName: pulumi.String("simpleBlueprint"),
			Description:   pulumi.String("blueprint contains all artifact kinds {'template', 'rbac', 'policy'}"),
			Parameters: blueprint.ParameterDefinitionMap{
				"costCenter": &blueprint.ParameterDefinitionArgs{
					DisplayName: pulumi.String("force cost center tag for all resources under given subscription."),
					Type:        pulumi.String("string"),
				},
				"owners": &blueprint.ParameterDefinitionArgs{
					DisplayName: pulumi.String("assign owners to subscription along with blueprint assignment."),
					Type:        pulumi.String("array"),
				},
				"storageAccountType": &blueprint.ParameterDefinitionArgs{
					DisplayName: pulumi.String("storage account type."),
					Type:        pulumi.String("string"),
				},
			},
			ResourceGroups: blueprint.ResourceGroupDefinitionMap{
				"storageRG": &blueprint.ResourceGroupDefinitionArgs{
					Description: pulumi.String("Contains storageAccounts that collect all shoebox logs."),
					DisplayName: pulumi.String("storage resource group"),
				},
			},
			ResourceScope: pulumi.String("providers/Microsoft.Management/managementGroups/ContosoOnlineGroup"),
			TargetScope:   pulumi.String("subscription"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
