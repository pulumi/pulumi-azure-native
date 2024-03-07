package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurearcdata/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurearcdata.NewSqlServerDatabase(ctx, "sqlServerDatabase", &azurearcdata.SqlServerDatabaseArgs{
			DatabaseName: pulumi.String("testdb"),
			Location:     pulumi.String("southeastasia"),
			Properties: azurearcdata.SqlServerDatabaseResourcePropertiesResponse{
				BackupInformation: &azurearcdata.SqlServerDatabaseResourcePropertiesBackupInformationArgs{
					LastFullBackup: pulumi.String("2022-05-05T16:26:33.883Z"),
					LastLogBackup:  pulumi.String("2022-05-10T16:26:33.883Z"),
				},
				CollationName:        pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
				CompatibilityLevel:   pulumi.Int(150),
				DatabaseCreationDate: pulumi.String("2022-04-05T16:26:33.883Z"),
				DatabaseOptions: &azurearcdata.SqlServerDatabaseResourcePropertiesDatabaseOptionsArgs{
					IsAutoCloseOn:               pulumi.Bool(true),
					IsAutoCreateStatsOn:         pulumi.Bool(true),
					IsAutoShrinkOn:              pulumi.Bool(true),
					IsAutoUpdateStatsOn:         pulumi.Bool(true),
					IsEncrypted:                 pulumi.Bool(true),
					IsMemoryOptimizationEnabled: pulumi.Bool(true),
					IsRemoteDataArchiveEnabled:  pulumi.Bool(true),
					IsTrustworthyOn:             pulumi.Bool(true),
				},
				IsReadOnly:       pulumi.Bool(true),
				RecoveryMode:     pulumi.String("Full"),
				SizeMB:           pulumi.Float64(150),
				SpaceAvailableMB: pulumi.Float64(100),
				State:            pulumi.String("Online"),
			},
			ResourceGroupName:     pulumi.String("testrg"),
			SqlServerInstanceName: pulumi.String("testSqlServerInstance"),
			Tags: pulumi.StringMap{
				"mytag": pulumi.String("myval"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
