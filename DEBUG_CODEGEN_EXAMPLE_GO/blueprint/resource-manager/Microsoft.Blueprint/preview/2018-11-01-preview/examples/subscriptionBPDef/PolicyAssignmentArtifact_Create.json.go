package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewTemplateArtifact(ctx, "templateArtifact", &blueprint.TemplateArtifactArgs{
			ArtifactName:  pulumi.String("costCenterPolicy"),
			BlueprintName: pulumi.String("simpleBlueprint"),
			ResourceScope: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
