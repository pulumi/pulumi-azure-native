package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewWebAppSitesController(ctx, "webAppSitesController", &offazure.WebAppSitesControllerArgs{
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("s13669b---4sI"),
			WebAppSiteName:    pulumi.String("3BXk-O8O6W3-GB7J"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
