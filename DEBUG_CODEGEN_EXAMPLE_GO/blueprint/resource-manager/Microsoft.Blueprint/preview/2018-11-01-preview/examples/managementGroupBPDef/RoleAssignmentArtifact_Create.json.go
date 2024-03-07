package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewTemplateArtifact(ctx, "templateArtifact", &blueprint.TemplateArtifactArgs{
			ArtifactName:  pulumi.String("ownerAssignment"),
			BlueprintName: pulumi.String("simpleBlueprint"),
			ResourceScope: pulumi.String("providers/Microsoft.Management/managementGroups/ContosoOnlineGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
