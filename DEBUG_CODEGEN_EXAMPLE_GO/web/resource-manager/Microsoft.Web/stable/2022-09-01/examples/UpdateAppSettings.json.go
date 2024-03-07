package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppApplicationSettings(ctx, "webAppApplicationSettings", &web.WebAppApplicationSettingsArgs{
			Name: pulumi.String("sitef6141"),
			Properties: pulumi.StringMap{
				"Setting1": pulumi.String("Value1"),
				"Setting2": pulumi.String("Value2"),
			},
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
