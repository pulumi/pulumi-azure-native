package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewAssignment(ctx, "assignment", &blueprint.AssignmentArgs{
			AssignmentName: pulumi.String("assignSimpleBlueprint"),
			BlueprintId:    pulumi.String("/providers/Microsoft.Management/managementGroups/ContosoOnlineGroup/providers/Microsoft.Blueprint/blueprints/simpleBlueprint"),
			Description:    pulumi.String("enforce pre-defined simpleBlueprint to this XXXXXXXX subscription."),
			Identity: &blueprint.ManagedServiceIdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location: pulumi.String("eastus"),
			Parameters: blueprint.ParameterValueMap{
				"costCenter": &blueprint.ParameterValueArgs{
					Value: pulumi.Any("Contoso/Online/Shopping/Production"),
				},
				"owners": &blueprint.ParameterValueArgs{
					Value: pulumi.Any{
						"johnDoe@contoso.com",
						"johnsteam@contoso.com",
					},
				},
				"storageAccountType": &blueprint.ParameterValueArgs{
					Value: pulumi.Any("Standard_LRS"),
				},
			},
			ResourceGroups: blueprint.ResourceGroupValueMap{
				"storageRG": &blueprint.ResourceGroupValueArgs{
					Location: pulumi.String("eastus"),
					Name:     pulumi.String("defaultRG"),
				},
			},
			ResourceScope: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
