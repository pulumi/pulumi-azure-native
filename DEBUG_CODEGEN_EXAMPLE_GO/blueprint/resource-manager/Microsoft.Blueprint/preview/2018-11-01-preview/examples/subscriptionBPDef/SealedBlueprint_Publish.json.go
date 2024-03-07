package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/blueprint/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := blueprint.NewPublishedBlueprint(ctx, "publishedBlueprint", &blueprint.PublishedBlueprintArgs{
			BlueprintName: pulumi.String("simpleBlueprint"),
			ResourceScope: pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
			VersionId:     pulumi.String("v2"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
