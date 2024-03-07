package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			CreateMode:        pulumi.String("Copy"),
			DatabaseName:      pulumi.String("dbcopy"),
			Location:          pulumi.String("southeastasia"),
			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:        pulumi.String("testsvr"),
			Sku: &sql.SkuArgs{
				Name: pulumi.String("S0"),
				Tier: pulumi.String("Standard"),
			},
			SourceDatabaseId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/testsvr/databases/testdb"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
