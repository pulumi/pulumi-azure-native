package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionPolicy(ctx, "protectionPolicy", &recoveryservices.ProtectionPolicyArgs{
			PolicyName: pulumi.String("testPolicy1"),
			Properties: recoveryservices.AzureVmWorkloadProtectionPolicy{
				BackupManagementType: "AzureWorkload",
				Settings: recoveryservices.Settings{
					Issqlcompression: false,
					TimeZone:         "Pacific Standard Time",
				},
				SubProtectionPolicy: []recoveryservices.SubProtectionPolicy{
					{
						PolicyType: "Full",
						RetentionPolicy: {
							MonthlySchedule: {
								RetentionDuration: {
									Count:        1,
									DurationType: "Months",
								},
								RetentionScheduleFormatType: "Weekly",
								RetentionScheduleWeekly: {
									DaysOfTheWeek: []recoveryservices.DayOfWeek{
										recoveryservices.DayOfWeekSunday,
									},
									WeeksOfTheMonth: []recoveryservices.WeekOfMonth{
										recoveryservices.WeekOfMonthSecond,
									},
								},
								RetentionTimes: []string{
									"2018-01-24T10:00:00Z",
								},
							},
							RetentionPolicyType: "LongTermRetentionPolicy",
							WeeklySchedule: {
								DaysOfTheWeek: []recoveryservices.DayOfWeek{
									recoveryservices.DayOfWeekSunday,
									recoveryservices.DayOfWeekTuesday,
								},
								RetentionDuration: {
									Count:        2,
									DurationType: "Weeks",
								},
								RetentionTimes: []string{
									"2018-01-24T10:00:00Z",
								},
							},
							YearlySchedule: {
								MonthsOfYear: []recoveryservices.MonthOfYear{
									recoveryservices.MonthOfYearJanuary,
									recoveryservices.MonthOfYearJune,
									recoveryservices.MonthOfYearDecember,
								},
								RetentionDuration: {
									Count:        1,
									DurationType: "Years",
								},
								RetentionScheduleFormatType: "Weekly",
								RetentionScheduleWeekly: {
									DaysOfTheWeek: []recoveryservices.DayOfWeek{
										recoveryservices.DayOfWeekSunday,
									},
									WeeksOfTheMonth: []recoveryservices.WeekOfMonth{
										recoveryservices.WeekOfMonthLast,
									},
								},
								RetentionTimes: []string{
									"2018-01-24T10:00:00Z",
								},
							},
						},
						SchedulePolicy: {
							SchedulePolicyType: "SimpleSchedulePolicy",
							ScheduleRunDays: []recoveryservices.DayOfWeek{
								recoveryservices.DayOfWeekSunday,
								recoveryservices.DayOfWeekTuesday,
							},
							ScheduleRunFrequency: "Weekly",
							ScheduleRunTimes: []string{
								"2018-01-24T10:00:00Z",
							},
						},
					},
					{
						PolicyType: "Differential",
						RetentionPolicy: {
							RetentionDuration: {
								Count:        8,
								DurationType: "Days",
							},
							RetentionPolicyType: "SimpleRetentionPolicy",
						},
						SchedulePolicy: {
							SchedulePolicyType: "SimpleSchedulePolicy",
							ScheduleRunDays: []recoveryservices.DayOfWeek{
								recoveryservices.DayOfWeekFriday,
							},
							ScheduleRunFrequency: "Weekly",
							ScheduleRunTimes: []string{
								"2018-01-24T10:00:00Z",
							},
						},
					},
					{
						PolicyType: "Log",
						RetentionPolicy: {
							RetentionDuration: {
								Count:        7,
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
			ResourceGroupName: pulumi.String("SwaggerTestRg"),
			VaultName:         pulumi.String("NetSDKTestRsVault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
