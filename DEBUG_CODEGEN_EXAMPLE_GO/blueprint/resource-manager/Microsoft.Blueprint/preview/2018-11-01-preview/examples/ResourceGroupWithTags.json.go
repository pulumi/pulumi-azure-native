package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewBlueprint(ctx, "blueprint", &blueprint.BlueprintArgs{
			BlueprintName: pulumi.String("simpleBlueprint"),
			Description:   pulumi.String("An example blueprint containing an RG with two tags."),
			ResourceGroups: blueprint.ResourceGroupDefinitionMap{
				"myRGName": &blueprint.ResourceGroupDefinitionArgs{
					DisplayName: pulumi.String("My Resource Group"),
					Location:    pulumi.String("westus"),
					Name:        pulumi.String("myRGName"),
					Tags: pulumi.StringMap{
						"costcenter":  pulumi.String("123456"),
						"nameOnlyTag": pulumi.String(""),
					},
				},
			},
			ResourceScope: pulumi.String("providers/Microsoft.Management/managementGroups/{ManagementGroupId}"),
			TargetScope:   pulumi.String("subscription"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
