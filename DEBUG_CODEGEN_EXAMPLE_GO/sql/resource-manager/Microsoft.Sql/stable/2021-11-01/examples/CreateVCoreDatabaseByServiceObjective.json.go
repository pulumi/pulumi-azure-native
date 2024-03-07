package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			DatabaseName:      pulumi.String("testdb"),
			Location:          pulumi.String("southeastasia"),
			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:        pulumi.String("testsvr"),
			Sku: &sql.SkuArgs{
				Capacity: pulumi.Int(2),
				Family:   pulumi.String("Gen4"),
				Name:     pulumi.String("BC"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
