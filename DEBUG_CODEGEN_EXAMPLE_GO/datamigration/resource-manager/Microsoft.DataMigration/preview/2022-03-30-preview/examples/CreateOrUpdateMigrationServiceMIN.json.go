package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewSqlMigrationService(ctx, "sqlMigrationService", &datamigration.SqlMigrationServiceArgs{
			Location:                pulumi.String("northeurope"),
			ResourceGroupName:       pulumi.String("testrg"),
			SqlMigrationServiceName: pulumi.String("testagent"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
