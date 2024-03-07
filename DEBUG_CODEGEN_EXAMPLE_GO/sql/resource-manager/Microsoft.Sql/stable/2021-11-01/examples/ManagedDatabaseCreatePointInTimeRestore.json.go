package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedDatabase(ctx, "managedDatabase", &sql.ManagedDatabaseArgs{
			CreateMode:          pulumi.String("PointInTimeRestore"),
			DatabaseName:        pulumi.String("managedDatabase"),
			Location:            pulumi.String("southeastasia"),
			ManagedInstanceName: pulumi.String("managedInstance"),
			ResourceGroupName:   pulumi.String("Default-SQL-SouthEastAsia"),
			RestorePointInTime:  pulumi.String("2017-07-14T05:35:31.503Z"),
			SourceDatabaseId:    pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/Default-SQL-SouthEastAsia/providers/Microsoft.Sql/managedInstances/testsvr/databases/testdb"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
