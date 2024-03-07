package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			CreateMode:         pulumi.String("PointInTimeRestore"),
			DatabaseName:       pulumi.String("testdw"),
			Location:           pulumi.String("southeastasia"),
			ResourceGroupName:  pulumi.String("Default-SQL-SouthEastAsia"),
			RestorePointInTime: pulumi.String("2022-01-22T05:35:31.503Z"),
			ServerName:         pulumi.String("testsvr"),
			SourceResourceId:   pulumi.String("/subscriptions/55555555-6666-7777-8888-999999999999/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/servers/srcsvr/databases/srcdw"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
