package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewTemplateSpecVersion(ctx, "templateSpecVersion", &resources.TemplateSpecVersionArgs{
			Description: pulumi.String("This is version v1.0 of our template content"),
			Location:    pulumi.String("eastus"),
			MainTemplate: pulumi.Any{
				Schema:         "http://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
				ContentVersion: "1.0.0.0",
				Parameters:     nil,
				Resources:      []interface{}{},
			},
			ResourceGroupName:   pulumi.String("templateSpecRG"),
			TemplateSpecName:    pulumi.String("simpleTemplateSpec"),
			TemplateSpecVersion: pulumi.String("v1.0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
