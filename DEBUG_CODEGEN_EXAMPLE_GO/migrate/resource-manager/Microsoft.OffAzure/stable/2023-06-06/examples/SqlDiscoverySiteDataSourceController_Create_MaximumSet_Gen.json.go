package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSqlDiscoverySiteDataSourceController(ctx, "sqlDiscoverySiteDataSourceController", &offazure.SqlDiscoverySiteDataSourceControllerArgs{
			DiscoverySiteDataSourceName: pulumi.String("q3--4O9O5vc-2"),
			DiscoverySiteId:             pulumi.String("xvtylcidvhdspuw"),
			ResourceGroupName:           pulumi.String("rgmigrate"),
			SiteName:                    pulumi.String("qU5b-1JBE45EC6Z-IF"),
			SqlSiteName:                 pulumi.String("I5-O3912-L3y2Q57"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
