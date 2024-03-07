package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/dataprotection/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := dataprotection.NewBackupVault(ctx, "backupVault", &dataprotection.BackupVaultArgs{
			Identity: &dataprotection.DppIdentityDetailsArgs{
				Type: pulumi.String("None"),
			},
			Location: pulumi.String("WestUS"),
			Properties: &dataprotection.BackupVaultTypeArgs{
				MonitoringSettings: &dataprotection.MonitoringSettingsArgs{
					AzureMonitorAlertSettings: &dataprotection.AzureMonitorAlertSettingsArgs{
						AlertsForAllJobFailures: pulumi.String("Enabled"),
					},
				},
				StorageSettings: dataprotection.StorageSettingArray{
					&dataprotection.StorageSettingArgs{
						DatastoreType: pulumi.String("VaultStore"),
						Type:          pulumi.String("LocallyRedundant"),
					},
				},
			},
			ResourceGroupName: pulumi.String("SampleResourceGroup"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("val1"),
			},
			VaultName: pulumi.String("swaggerExample"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
