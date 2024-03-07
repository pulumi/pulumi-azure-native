package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSqlDiscoverySiteDataSourceController(ctx, "sqlDiscoverySiteDataSourceController", &offazure.SqlDiscoverySiteDataSourceControllerArgs{
			DiscoverySiteDataSourceName: pulumi.String("Sw-cch"),
			ResourceGroupName:           pulumi.String("rgmigrate"),
			SiteName:                    pulumi.String("-2n-j-5O"),
			SqlSiteName:                 pulumi.String("32F0K64"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
