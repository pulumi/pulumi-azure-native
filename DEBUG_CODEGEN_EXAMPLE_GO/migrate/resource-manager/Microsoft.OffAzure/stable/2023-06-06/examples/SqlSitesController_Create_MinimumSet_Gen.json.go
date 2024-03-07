package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/offazure/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := offazure.NewSqlSitesController(ctx, "sqlSitesController", &offazure.SqlSitesControllerArgs{
			ResourceGroupName: pulumi.String("rgmigrate"),
			SiteName:          pulumi.String("Ag5-Ex4n7G5Qn-5-yHL5"),
			SqlSiteName:       pulumi.String("-03--VkVPfDq5"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
