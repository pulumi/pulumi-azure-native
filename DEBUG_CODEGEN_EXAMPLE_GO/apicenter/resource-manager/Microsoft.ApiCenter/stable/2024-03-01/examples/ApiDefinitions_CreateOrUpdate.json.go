package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apicenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apicenter.NewApiDefinition(ctx, "apiDefinition", &apicenter.ApiDefinitionArgs{
			ApiName:           pulumi.String("openapi"),
			DefinitionName:    pulumi.String("openapi"),
			Description:       pulumi.String("Default spec"),
			ResourceGroupName: pulumi.String("contoso-resources"),
			ServiceName:       pulumi.String("contoso"),
			Title:             pulumi.String("OpenAPI"),
			VersionName:       pulumi.String("2023-01-01"),
			WorkspaceName:     pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
