package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewTableResourceTable(ctx, "tableResourceTable", &documentdb.TableResourceTableArgs{
			AccountName: pulumi.String("ddb1"),
			Location:    pulumi.String("West US"),
			Options:     nil,
			Resource: &documentdb.TableResourceArgs{
				Id: pulumi.String("tableName"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			TableName:         pulumi.String("tableName"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
