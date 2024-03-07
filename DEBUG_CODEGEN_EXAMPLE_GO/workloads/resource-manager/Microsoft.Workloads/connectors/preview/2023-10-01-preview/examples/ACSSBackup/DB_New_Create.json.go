package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewACSSBackupConnection(ctx, "acssBackupConnection", &workloads.ACSSBackupConnectionArgs{
			BackupData: workloads.SqlBackupData{
				BackupPolicy: workloads.DBBackupPolicyProperties{
					BackupManagementType: "AzureWorkload",
					Name:                 "defaultDbPolicy",
					ProtectedItemsCount:  0,
					Settings: workloads.Settings{
						IsCompression:    false,
						Issqlcompression: false,
						TimeZone:         "UTC",
					},
					SubProtectionPolicy: []workloads.SubProtectionPolicy{
						{
							PolicyType: "Full",
							RetentionPolicy: {
								DailySchedule: {
									RetentionDuration: {
										Count:        30,
										DurationType: "Days",
									},
									RetentionTimes: []string{
										"2018-01-10T18:30:00Z",
									},
								},
								RetentionPolicyType: "LongTermRetentionPolicy",
							},
							SchedulePolicy: {
								SchedulePolicyType:   "SimpleSchedulePolicy",
								ScheduleRunFrequency: "Daily",
								ScheduleRunTimes: []string{
									"2018-01-10T18:30:00Z",
								},
								ScheduleWeeklyFrequency: 0,
							},
						},
						{
							PolicyType: "Log",
							RetentionPolicy: {
								RetentionDuration: {
									Count:        30,
									DurationType: "Days",
								},
								RetentionPolicyType: "SimpleRetentionPolicy",
							},
							SchedulePolicy: {
								ScheduleFrequencyInMins: 60,
								SchedulePolicyType:      "LogSchedulePolicy",
							},
						},
					},
					WorkLoadType: "SQLDataBase",
				},
				BackupType: "SQL",
				RecoveryServicesVault: workloads.ExistingRecoveryServicesVault{
					Id:        "/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourceGroups/test-rg/providers/Microsoft.RecoveryServices/vaults/test-vault",
					VaultType: "Existing",
				},
			},
			BackupName:        pulumi.String("dbBackup"),
			ConnectorName:     pulumi.String("C1"),
			Location:          pulumi.String("westcentralus"),
			ResourceGroupName: pulumi.String("test-rg"),
			Tags:              nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
