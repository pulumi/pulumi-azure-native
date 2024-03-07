package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datamigration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datamigration.NewDatabaseMigrationsSqlDb(ctx, "databaseMigrationsSqlDb", &datamigration.DatabaseMigrationsSqlDbArgs{
			Properties: &datamigration.DatabaseMigrationPropertiesSqlDbArgs{
				Kind:               pulumi.String("SqlDb"),
				MigrationService:   pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.DataMigration/sqlMigrationServices/testagent"),
				Scope:              pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Sql/servers/sqldbinstance"),
				SourceDatabaseName: pulumi.String("aaa"),
				SourceSqlConnection: &datamigration.SqlConnectionInformationArgs{
					Authentication:         pulumi.String("WindowsAuthentication"),
					DataSource:             pulumi.String("aaa"),
					EncryptConnection:      pulumi.Bool(true),
					Password:               pulumi.String("placeholder"),
					TrustServerCertificate: pulumi.Bool(true),
					UserName:               pulumi.String("bbb"),
				},
				TargetSqlConnection: &datamigration.SqlConnectionInformationArgs{
					Authentication:         pulumi.String("SqlAuthentication"),
					DataSource:             pulumi.String("sqldbinstance"),
					EncryptConnection:      pulumi.Bool(true),
					Password:               pulumi.String("placeholder"),
					TrustServerCertificate: pulumi.Bool(true),
					UserName:               pulumi.String("bbb"),
				},
			},
			ResourceGroupName: pulumi.String("testrg"),
			SqlDbInstanceName: pulumi.String("sqldbinstance"),
			TargetDbName:      pulumi.String("db1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
