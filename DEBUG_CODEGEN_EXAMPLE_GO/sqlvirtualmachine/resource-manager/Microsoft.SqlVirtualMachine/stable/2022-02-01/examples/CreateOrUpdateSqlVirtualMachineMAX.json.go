package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sqlvirtualmachine/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sqlvirtualmachine.NewSqlVirtualMachine(ctx, "sqlVirtualMachine", &sqlvirtualmachine.SqlVirtualMachineArgs{
			AssessmentSettings: &sqlvirtualmachine.AssessmentSettingsArgs{
				Enable:         pulumi.Bool(true),
				RunImmediately: pulumi.Bool(true),
				Schedule: &sqlvirtualmachine.ScheduleArgs{
					DayOfWeek:      sqlvirtualmachine.AssessmentDayOfWeekSunday,
					Enable:         pulumi.Bool(true),
					StartTime:      pulumi.String("23:17"),
					WeeklyInterval: pulumi.Int(1),
				},
			},
			AutoBackupSettings: &sqlvirtualmachine.AutoBackupSettingsArgs{
				BackupScheduleType:    pulumi.String("Manual"),
				BackupSystemDbs:       pulumi.Bool(true),
				Enable:                pulumi.Bool(true),
				EnableEncryption:      pulumi.Bool(true),
				FullBackupFrequency:   pulumi.String("Daily"),
				FullBackupStartTime:   pulumi.Int(6),
				FullBackupWindowHours: pulumi.Int(11),
				LogBackupFrequency:    pulumi.Int(10),
				Password:              pulumi.String("<Password>"),
				RetentionPeriod:       pulumi.Int(17),
				StorageAccessKey:      pulumi.String("<primary storage access key>"),
				StorageAccountUrl:     pulumi.String("https://teststorage.blob.core.windows.net/"),
				StorageContainerName:  pulumi.String("testcontainer"),
			},
			AutoPatchingSettings: &sqlvirtualmachine.AutoPatchingSettingsArgs{
				DayOfWeek:                     sqlvirtualmachine.DayOfWeekSunday,
				Enable:                        pulumi.Bool(true),
				MaintenanceWindowDuration:     pulumi.Int(60),
				MaintenanceWindowStartingHour: pulumi.Int(2),
			},
			KeyVaultCredentialSettings: &sqlvirtualmachine.KeyVaultCredentialSettingsArgs{
				Enable: pulumi.Bool(false),
			},
			Location:          pulumi.String("northeurope"),
			ResourceGroupName: pulumi.String("testrg"),
			ServerConfigurationsManagementSettings: &sqlvirtualmachine.ServerConfigurationsManagementSettingsArgs{
				AdditionalFeaturesServerConfigurations: &sqlvirtualmachine.AdditionalFeaturesServerConfigurationsArgs{
					IsRServicesEnabled: pulumi.Bool(false),
				},
				SqlConnectivityUpdateSettings: &sqlvirtualmachine.SqlConnectivityUpdateSettingsArgs{
					ConnectivityType:      pulumi.String("PRIVATE"),
					Port:                  pulumi.Int(1433),
					SqlAuthUpdatePassword: pulumi.String("<password>"),
					SqlAuthUpdateUserName: pulumi.String("sqllogin"),
				},
				SqlInstanceSettings: &sqlvirtualmachine.SQLInstanceSettingsArgs{
					Collation:                          pulumi.String("SQL_Latin1_General_CP1_CI_AS"),
					IsIfiEnabled:                       pulumi.Bool(true),
					IsLpimEnabled:                      pulumi.Bool(true),
					IsOptimizeForAdHocWorkloadsEnabled: pulumi.Bool(true),
					MaxDop:                             pulumi.Int(8),
					MaxServerMemoryMB:                  pulumi.Int(128),
					MinServerMemoryMB:                  pulumi.Int(0),
				},
				SqlStorageUpdateSettings: &sqlvirtualmachine.SqlStorageUpdateSettingsArgs{
					DiskConfigurationType: pulumi.String("NEW"),
					DiskCount:             pulumi.Int(1),
					StartingDeviceId:      pulumi.Int(2),
				},
				SqlWorkloadTypeUpdateSettings: &sqlvirtualmachine.SqlWorkloadTypeUpdateSettingsArgs{
					SqlWorkloadType: pulumi.String("OLTP"),
				},
			},
			SqlImageSku:              pulumi.String("Enterprise"),
			SqlManagement:            pulumi.String("Full"),
			SqlServerLicenseType:     pulumi.String("PAYG"),
			SqlVirtualMachineName:    pulumi.String("testvm"),
			VirtualMachineResourceId: pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/testrg/providers/Microsoft.Compute/virtualMachines/testvm"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
