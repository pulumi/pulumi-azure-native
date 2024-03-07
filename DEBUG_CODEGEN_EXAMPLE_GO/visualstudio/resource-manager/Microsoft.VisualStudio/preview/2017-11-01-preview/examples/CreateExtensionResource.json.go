package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/visualstudio/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := visualstudio.NewExtension(ctx, "extension", &visualstudio.ExtensionArgs{
			AccountResourceName:   pulumi.String("ExampleAccount"),
			ExtensionResourceName: pulumi.String("ms.example"),
			Location:              pulumi.String("Central US"),
			Plan: &visualstudio.ExtensionResourcePlanArgs{
				Name:          pulumi.String("ExamplePlan"),
				Product:       pulumi.String("ExampleExtensionName"),
				PromotionCode: pulumi.String(""),
				Publisher:     pulumi.String("ExampleExtensionPublisher"),
				Version:       pulumi.String("1.0"),
			},
			Properties:        nil,
			ResourceGroupName: pulumi.String("VS-Example-Group"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
