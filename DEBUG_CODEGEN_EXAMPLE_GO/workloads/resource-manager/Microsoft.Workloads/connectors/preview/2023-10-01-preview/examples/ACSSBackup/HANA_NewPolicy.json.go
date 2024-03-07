package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/workloads/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := workloads.NewACSSBackupConnection(ctx, "acssBackupConnection", &workloads.ACSSBackupConnectionArgs{
			BackupData: workloads.HanaBackupData{
				BackupPolicy: workloads.DBBackupPolicyProperties{
					BackupManagementType: "AzureWorkload",
					Name:                 "defaultHanaPolicy",
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
									TieringMode: "DoNotTier",
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
					WorkLoadType: "SAPHanaDatabase",
				},
				BackupType: "HANA",
				DbInstanceSnapshotBackupPolicy: workloads.DBBackupPolicyProperties{
					BackupManagementType: "AzureWorkload",
					Name:                 "defaultDbInstanceSnapshotPolicy",
					Settings: workloads.Settings{
						IsCompression:    false,
						Issqlcompression: false,
						TimeZone:         "UTC",
					},
					SubProtectionPolicy: []workloads.SubProtectionPolicy{
						{
							PolicyType: "SnapshotFull",
							SchedulePolicy: {
								SchedulePolicyType:   "SimpleSchedulePolicy",
								ScheduleRunFrequency: "Daily",
								ScheduleRunTimes: []string{
									"2023-09-18T06:30:00.000Z",
								},
							},
							SnapshotBackupAdditionalDetails: {
								InstantRPDetails:              "test-rg",
								InstantRpRetentionRangeInDays: 1,
								UserAssignedManagedIdentityDetails: {
									IdentityArmId: "/subscriptions/6d875e77-e412-4d7d-9af4-8895278b4443/resourcegroups/test-rg/providers/Microsoft.ManagedIdentity/userAssignedIdentities/testMsi",
									IdentityName:  "testMsi",
									UserAssignedIdentityProperties: {
										ClientId:    "c3a877cf-51f8-4031-8f17-ab562d1e7737",
										PrincipalId: "2f5834bd-4b86-4d85-a8df-6dd829a6418c",
									},
								},
							},
						},
					},
					WorkLoadType: "SAPHanaDBInstance",
				},
				HdbuserstoreKeyName: "abcd",
				InstanceNumber:      "00",
				RecoveryServicesVault: workloads.NewRecoveryServicesVault{
					Name:          "test-vault",
					ResourceGroup: "test-rg",
					VaultType:     "New",
				},
				SslConfiguration: workloads.SSLConfiguration{
					SslCryptoProvider:        "commoncrypto",
					SslHostNameInCertificate: "hostname",
					SslKeyStore:              "sapsrv.pse",
					SslTrustStore:            "sapsrv.pse",
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
