package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlStoredProcedure(ctx, "sqlResourceSqlStoredProcedure", &documentdb.SqlResourceSqlStoredProcedureArgs{
			AccountName:   pulumi.String("ddb1"),
			ContainerName: pulumi.String("containerName"),
			DatabaseName:  pulumi.String("databaseName"),
			Options:       nil,
			Resource: &documentdb.SqlStoredProcedureResourceArgs{
				Body: pulumi.String("body"),
				Id:   pulumi.String("storedProcedureName"),
			},
			ResourceGroupName:   pulumi.String("rg1"),
			StoredProcedureName: pulumi.String("storedProcedureName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
