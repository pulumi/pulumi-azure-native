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
					Name:                 "defaultSqlPolicy",
					ProtectedItemsCount:  0,
					Settings: workloads.Settings{
						IsCompression:    true,
						Issqlcompression: true,
						TimeZone:         "UTC",
					},
					SubProtectionPolicy: []workloads.SubProtectionPolicy{
						{
							PolicyType: "Full",
							RetentionPolicy: {
								MonthlySchedule: {
									RetentionDuration: {
										Count:        60,
										DurationType: "Months",
									},
									RetentionScheduleFormatType: "Weekly",
									RetentionScheduleWeekly: {
										DaysOfTheWeek: []workloads.DayOfWeek{
											workloads.DayOfWeekSunday,
										},
										WeeksOfTheMonth: []workloads.WeekOfMonth{
											workloads.WeekOfMonthFirst,
										},
									},
									RetentionTimes: []string{
										"2022-11-29T19:30:00.000Z",
									},
								},
								RetentionPolicyType: "LongTermRetentionPolicy",
								WeeklySchedule: {
									DaysOfTheWeek: []workloads.DayOfWeek{
										workloads.DayOfWeekSunday,
									},
									RetentionDuration: {
										Count:        104,
										DurationType: "Weeks",
									},
									RetentionTimes: []string{
										"2022-11-29T19:30:00.000Z",
									},
								},
								YearlySchedule: {
									MonthsOfYear: []workloads.MonthOfYear{
										workloads.MonthOfYearJanuary,
									},
									RetentionDuration: {
										Count:        10,
										DurationType: "Years",
									},
									RetentionScheduleFormatType: "Weekly",
									RetentionScheduleWeekly: {
										DaysOfTheWeek: []workloads.DayOfWeek{
											workloads.DayOfWeekSunday,
										},
										WeeksOfTheMonth: []workloads.WeekOfMonth{
											workloads.WeekOfMonthFirst,
										},
									},
									RetentionTimes: []string{
										"2022-11-29T19:30:00.000Z",
									},
								},
							},
							SchedulePolicy: {
								SchedulePolicyType: "SimpleSchedulePolicy",
								ScheduleRunDays: []workloads.DayOfWeek{
									workloads.DayOfWeekSunday,
								},
								ScheduleRunFrequency: "Weekly",
								ScheduleRunTimes: []string{
									"2022-11-29T19:30:00.000Z",
								},
							},
							TieringPolicy: {
								ArchivedRP: {
									Duration:     45,
									DurationType: "Days",
									TieringMode:  "TierAfter",
								},
							},
						},
						{
							PolicyType: "Differential",
							RetentionPolicy: {
								RetentionDuration: {
									Count:        30,
									DurationType: "Days",
								},
								RetentionPolicyType: "SimpleRetentionPolicy",
							},
							SchedulePolicy: {
								SchedulePolicyType: "SimpleSchedulePolicy",
								ScheduleRunDays: []workloads.DayOfWeek{
									workloads.DayOfWeekMonday,
								},
								ScheduleRunFrequency: "Weekly",
								ScheduleRunTimes: []string{
									"2022-09-29T02:00:00Z",
								},
								ScheduleWeeklyFrequency: 0,
							},
						},
						{
							PolicyType: "Log",
							RetentionPolicy: {
								RetentionDuration: {
									Count:        20,
									DurationType: "Days",
								},
								RetentionPolicyType: "SimpleRetentionPolicy",
							},
							SchedulePolicy: {
								ScheduleFrequencyInMins: 120,
								SchedulePolicyType:      "LogSchedulePolicy",
							},
						},
					},
					WorkLoadType: "SQLDataBase",
				},
				BackupType: "SQL",
				RecoveryServicesVault: workloads.NewRecoveryServicesVault{
					Name:          "test-vault",
					ResourceGroup: "test-rg",
					VaultType:     "New",
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
