package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewStaticSiteUserProvidedFunctionAppForStaticSite(ctx, "staticSiteUserProvidedFunctionAppForStaticSite", &web.StaticSiteUserProvidedFunctionAppForStaticSiteArgs{
			FunctionAppName:       pulumi.String("testFunctionApp"),
			FunctionAppRegion:     pulumi.String("West US 2"),
			FunctionAppResourceId: pulumi.String("/subscription/34adfa4f-cedf-4dc0-ba29-b6d1a69ab345/resourceGroups/functionRG/providers/Microsoft.Web/sites/testFunctionApp"),
			IsForced:              pulumi.Bool(true),
			Name:                  pulumi.String("testStaticSite0"),
			ResourceGroupName:     pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
