package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			CreateMode:        pulumi.String("Secondary"),
			DatabaseName:      pulumi.String("testdb"),
			Location:          pulumi.String("southeastasia"),
			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
			SecondaryType:     pulumi.String("Named"),
			ServerName:        pulumi.String("testsvr"),
			Sku: &sql.SkuArgs{
				Capacity: pulumi.Int(2),
				Name:     pulumi.String("HS_Gen4"),
				Tier:     pulumi.String("Hyperscale"),
			},
			SourceDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-NorthEurope/providers/Microsoft.Sql/servers/testsvr1/databases/primarydb"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
