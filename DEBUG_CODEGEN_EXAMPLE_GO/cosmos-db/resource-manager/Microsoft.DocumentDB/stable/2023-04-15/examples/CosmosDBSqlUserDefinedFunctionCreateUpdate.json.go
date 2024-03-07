package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlUserDefinedFunction(ctx, "sqlResourceSqlUserDefinedFunction", &documentdb.SqlResourceSqlUserDefinedFunctionArgs{
			AccountName:   pulumi.String("ddb1"),
			ContainerName: pulumi.String("containerName"),
			DatabaseName:  pulumi.String("databaseName"),
			Options:       nil,
			Resource: &documentdb.SqlUserDefinedFunctionResourceArgs{
				Body: pulumi.String("body"),
				Id:   pulumi.String("userDefinedFunctionName"),
			},
			ResourceGroupName:       pulumi.String("rg1"),
			UserDefinedFunctionName: pulumi.String("userDefinedFunctionName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
