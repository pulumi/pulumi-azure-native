package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/resources/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := resources.NewTemplateSpec(ctx, "templateSpec", &resources.TemplateSpecArgs{
			Description:       pulumi.String("A very simple Template Spec"),
			Location:          pulumi.String("eastus"),
			ResourceGroupName: pulumi.String("templateSpecRG"),
			TemplateSpecName:  pulumi.String("simpleTemplateSpec"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
