package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewACSSBackupConnection(ctx, "acssBackupConnection", &workloads.ACSSBackupConnectionArgs{
			BackupData: workloads.VMBackupData{
				BackupPolicy: workloads.VMBackupPolicyProperties{
					BackupManagementType:          "AzureIaasVM",
					InstantRPDetails:              nil,
					InstantRpRetentionRangeInDays: 2,
					Name:                          "defaultVmPolicy",
					ProtectedItemsCount:           0,
					RetentionPolicy: workloads.LongTermRetentionPolicy{
						DailySchedule: workloads.DailyRetentionSchedule{
							RetentionDuration: workloads.RetentionDuration{
								Count:        30,
								DurationType: "Days",
							},
							RetentionTimes: []string{
								"2018-01-10T18:30:00Z",
							},
						},
						RetentionPolicyType: "LongTermRetentionPolicy",
					},
					SchedulePolicy: workloads.SimpleSchedulePolicy{
						SchedulePolicyType:   "SimpleSchedulePolicy",
						ScheduleRunFrequency: "Daily",
						ScheduleRunTimes: []string{
							"2018-01-10T18:30:00Z",
						},
						ScheduleWeeklyFrequency: 0,
					},
				},
				BackupType: "VM",
				DiskExclusionProperties: workloads.DiskExclusionProperties{
					DiskLunList:     []interface{}{},
					IsInclusionList: true,
				},
				RecoveryServicesVault: workloads.NewRecoveryServicesVault{
					Name:          "test-vault",
					ResourceGroup: "test-rg",
					VaultType:     "New",
				},
			},
			BackupName:        pulumi.String("vmBackup"),
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
