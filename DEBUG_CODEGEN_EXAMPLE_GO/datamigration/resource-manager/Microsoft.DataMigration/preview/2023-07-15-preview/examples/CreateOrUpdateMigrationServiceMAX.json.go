package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewMigrationService(ctx, "migrationService", &datamigration.MigrationServiceArgs{
			Location:             pulumi.String("northeurope"),
			MigrationServiceName: pulumi.String("testagent"),
			ResourceGroupName:    pulumi.String("testrg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
