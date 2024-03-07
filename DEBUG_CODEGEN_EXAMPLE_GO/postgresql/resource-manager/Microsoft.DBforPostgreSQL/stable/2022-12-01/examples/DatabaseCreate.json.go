package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewDatabase(ctx, "database", &dbforpostgresql.DatabaseArgs{
			Charset:           pulumi.String("utf8"),
			Collation:         pulumi.String("en_US.utf8"),
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
