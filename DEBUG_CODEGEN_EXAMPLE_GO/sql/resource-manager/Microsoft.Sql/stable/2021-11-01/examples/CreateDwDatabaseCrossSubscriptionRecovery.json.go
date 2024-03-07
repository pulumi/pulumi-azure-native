package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			CreateMode:        pulumi.String("Recovery"),
			DatabaseName:      pulumi.String("testdw"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("Default-SQL-WestUS"),
			ServerName:        pulumi.String("testsvr"),
			SourceResourceId:  pulumi.String("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-EastUS/providers/Microsoft.Sql/servers/srcsvr/recoverabledatabases/srcdw"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
