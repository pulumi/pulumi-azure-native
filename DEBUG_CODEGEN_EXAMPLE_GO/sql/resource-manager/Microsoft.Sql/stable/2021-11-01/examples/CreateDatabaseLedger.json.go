package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewDatabase(ctx, "database", &sql.DatabaseArgs{
			DatabaseName:      pulumi.String("testdb"),
			IsLedgerOn:        pulumi.Bool(true),
			Location:          pulumi.String("southeastasia"),
			ResourceGroupName: pulumi.String("Default-SQL-SouthEastAsia"),
			ServerName:        pulumi.String("testsvr"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
