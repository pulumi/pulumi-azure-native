package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewManagedDatabase(ctx, "managedDatabase", &sql.ManagedDatabaseArgs{
			AutoCompleteRestore:      pulumi.Bool(true),
			Collation:                pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
			CreateMode:               pulumi.String("RestoreExternalBackup"),
			DatabaseName:             pulumi.String("managedDatabase"),
			LastBackupName:           pulumi.String("last_backup_name"),
			Location:                 pulumi.String("southeastasia"),
			ManagedInstanceName:      pulumi.String("managedInstance"),
			ResourceGroupName:        pulumi.String("Default-SQL-SouthEastAsia"),
			StorageContainerSasToken: pulumi.String("sv=2015-12-11&sr=c&sp=rl&sig=1234"),
			StorageContainerUri:      pulumi.String("https://myaccountname.blob.core.windows.net/backups"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
