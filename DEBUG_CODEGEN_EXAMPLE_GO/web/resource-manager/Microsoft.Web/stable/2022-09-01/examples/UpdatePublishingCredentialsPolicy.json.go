package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/web/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := web.NewWebAppScmAllowed(ctx, "webAppScmAllowed", &web.WebAppScmAllowedArgs{
			Allow:             pulumi.Bool(true),
			Name:              pulumi.String("testSite"),
			ResourceGroupName: pulumi.String("rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
