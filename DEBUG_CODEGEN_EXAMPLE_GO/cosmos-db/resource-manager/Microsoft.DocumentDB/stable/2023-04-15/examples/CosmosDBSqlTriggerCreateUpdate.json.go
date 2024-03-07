package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/documentdb/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := documentdb.NewSqlResourceSqlTrigger(ctx, "sqlResourceSqlTrigger", &documentdb.SqlResourceSqlTriggerArgs{
			AccountName:   pulumi.String("ddb1"),
			ContainerName: pulumi.String("containerName"),
			DatabaseName:  pulumi.String("databaseName"),
			Options:       nil,
			Resource: &documentdb.SqlTriggerResourceArgs{
				Body:             pulumi.String("body"),
				Id:               pulumi.String("triggerName"),
				TriggerOperation: pulumi.String("triggerOperation"),
				TriggerType:      pulumi.String("triggerType"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			TriggerName:       pulumi.String("triggerName"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
