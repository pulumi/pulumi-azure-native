package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			Collation:                  pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
			CreateMode:                 pulumi.String("Default"),
			DatabaseName:               pulumi.String("testdb"),
			Location:                   pulumi.String("southeastasia"),
			MaintenanceConfigurationId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/providers/Microsoft.Maintenance/publicMaintenanceConfigurations/SQL_SouthEastAsia_1"),
			MaxSizeBytes:               pulumi.Float64(1073741824),
			ResourceGroupName:          pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:                 pulumi.String("testsvr"),
			Sku: &sql.SkuArgs{
				Name: pulumi.String("S2"),
				Tier: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
