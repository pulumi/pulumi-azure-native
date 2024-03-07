package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlDatabase(ctx, "sqlResourceSqlDatabase", &documentdb.SqlResourceSqlDatabaseArgs{
			AccountName:  pulumi.String("ddb1"),
			DatabaseName: pulumi.String("databaseName"),
			Location:     pulumi.String("West US"),
			Options:      nil,
			Resource: &documentdb.SqlDatabaseResourceArgs{
				Id: pulumi.String("databaseName"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
