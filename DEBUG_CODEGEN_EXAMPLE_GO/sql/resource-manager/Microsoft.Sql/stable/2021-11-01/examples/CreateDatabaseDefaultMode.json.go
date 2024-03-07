package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			Collation:         pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
			CreateMode:        pulumi.String("Default"),
			DatabaseName:      pulumi.String("testdb"),
			Location:          pulumi.String("southeastasia"),
			MaxSizeBytes:      pulumi.Float64(1073741824),
			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:        pulumi.String("testsvr"),
			Sku: &sql.SkuArgs{
				Name: pulumi.String("S0"),
				Tier: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
