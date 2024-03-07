package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedDatabase(ctx, "managedDatabase", &sql.ManagedDatabaseArgs{
			DatabaseName:        pulumi.String("managedDatabase"),
			Location:            pulumi.String("southeastasia"),
			ManagedInstanceName: pulumi.String("managedInstance"),
			ResourceGroupName:   pulumi.String("Default-SQL-SouthEastAsia"),
			Tags: pulumi.StringMap{
				"tagKey1": pulumi.String("TagValue1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
