package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewDocumentation(ctx, "documentation", &apimanagement.DocumentationArgs{
			Content:           pulumi.String("content"),
			DocumentationId:   pulumi.String("57d1f7558aa04f15146d9d8a"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			Title:             pulumi.String("Title"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
