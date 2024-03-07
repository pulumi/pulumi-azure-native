package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedDatabase(ctx, "managedDatabase", &sql.ManagedDatabaseArgs{
			CreateMode:            pulumi.String("Recovery"),
			DatabaseName:          pulumi.String("testdb_recovered"),
			Location:              pulumi.String("southeastasia"),
			ManagedInstanceName:   pulumi.String("server1"),
			RecoverableDatabaseId: pulumi.String("/subscriptions/11111111-2222-3333-4444-555555555555/resourceGroups/Default-SQL-WestEurope/providers/Microsoft.Sql/managedInstances/testsvr/recoverableDatabases/testdb"),
			ResourceGroupName:     pulumi.String("Default-SQL-SouthEastAsia"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
