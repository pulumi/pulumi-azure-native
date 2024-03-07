package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dbforpostgresql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dbforpostgresql.NewMigration(ctx, "migration", &dbforpostgresql.MigrationArgs{
			DbsToMigrate: pulumi.StringArray{
				pulumi.String("db1"),
				pulumi.String("db2"),
				pulumi.String("db3"),
				pulumi.String("db4"),
			},
			Location:             pulumi.String("westus"),
			MigrationMode:        pulumi.String("Offline"),
			MigrationName:        pulumi.String("testmigration"),
			OverwriteDbsInTarget: pulumi.String("True"),
			ResourceGroupName:    pulumi.String("testrg"),
			SecretParameters: &dbforpostgresql.MigrationSecretParametersArgs{
				AdminCredentials: &dbforpostgresql.AdminCredentialsArgs{
					SourceServerPassword: pulumi.String("xxxxxxxx"),
					TargetServerPassword: pulumi.String("xxxxxxxx"),
				},
			},
			SourceDbServerResourceId: pulumi.String("/subscriptions/ffffffff-ffff-ffff-ffff-ffffffffffff/resourceGroups/testrg/providers/Microsoft.DBForPostgreSql/servers/testsource"),
			TargetDbServerName:       pulumi.String("testtarget"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
