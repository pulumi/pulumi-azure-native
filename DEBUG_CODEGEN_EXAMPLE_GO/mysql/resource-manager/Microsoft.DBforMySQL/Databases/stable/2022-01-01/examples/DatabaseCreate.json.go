package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbformysql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbformysql.NewDatabase(ctx, "database", &dbformysql.DatabaseArgs{
			Charset:           pulumi.String("utf8"),
			Collation:         pulumi.String("utf8_general_ci"),
			DatabaseName:      pulumi.String("db1"),
			ResourceGroupName: pulumi.String("TestGroup"),
			ServerName:        pulumi.String("testserver"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
