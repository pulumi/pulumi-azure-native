package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewWebAppDiscoverySiteDataSourcesController(ctx, "webAppDiscoverySiteDataSourcesController", &offazure.WebAppDiscoverySiteDataSourcesControllerArgs{
			DiscoverySiteDataSourceName: pulumi.String("171iiIY1M39-M--VU2x8uOy"),
			DiscoverySiteId:             pulumi.String("fwkwva"),
			ResourceGroupName:           pulumi.String("rgmigrate"),
			SiteName:                    pulumi.String("0M4C6JZc"),
			WebAppSiteName:              pulumi.String("QhlGHU7obm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
