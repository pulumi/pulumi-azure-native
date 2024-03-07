package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewWebAppDiscoverySiteDataSourcesController(ctx, "webAppDiscoverySiteDataSourcesController", &offazure.WebAppDiscoverySiteDataSourcesControllerArgs{
			DiscoverySiteDataSourceName: pulumi.String("Q-38555Y-2-8-6-bdZk2y"),
			ResourceGroupName:           pulumi.String("rgmigrate"),
			SiteName:                    pulumi.String("257-4BP-1j"),
			WebAppSiteName:              pulumi.String("S--3265vli3j4X--Vy-J"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
