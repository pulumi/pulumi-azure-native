package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionPolicy(ctx, "protectionPolicy", &recoveryservices.ProtectionPolicyArgs{
			PolicyName: pulumi.String("newPolicy2"),
			Properties: recoveryservices.AzureFileShareProtectionPolicy{
				BackupManagementType: "AzureStorage",
				RetentionPolicy: recoveryservices.LongTermRetentionPolicy{
					DailySchedule: recoveryservices.DailyRetentionSchedule{
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        5,
							DurationType: "Days",
						},
					},
					MonthlySchedule: recoveryservices.MonthlyRetentionSchedule{
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        60,
							DurationType: "Months",
						},
						RetentionScheduleFormatType: "Weekly",
						RetentionScheduleWeekly: recoveryservices.WeeklyRetentionFormat{
							DaysOfTheWeek: []recoveryservices.DayOfWeek{
								recoveryservices.DayOfWeekSunday,
							},
							WeeksOfTheMonth: []recoveryservices.WeekOfMonth{
								recoveryservices.WeekOfMonthFirst,
							},
						},
					},
					RetentionPolicyType: "LongTermRetentionPolicy",
					WeeklySchedule: recoveryservices.WeeklyRetentionSchedule{
						DaysOfTheWeek: []recoveryservices.DayOfWeek{
							recoveryservices.DayOfWeekSunday,
						},
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        12,
							DurationType: "Weeks",
						},
					},
					YearlySchedule: recoveryservices.YearlyRetentionSchedule{
						MonthsOfYear: []recoveryservices.MonthOfYear{
							recoveryservices.MonthOfYearJanuary,
						},
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        10,
							DurationType: "Years",
						},
						RetentionScheduleFormatType: "Weekly",
						RetentionScheduleWeekly: recoveryservices.WeeklyRetentionFormat{
							DaysOfTheWeek: []recoveryservices.DayOfWeek{
								recoveryservices.DayOfWeekSunday,
							},
							WeeksOfTheMonth: []recoveryservices.WeekOfMonth{
								recoveryservices.WeekOfMonthFirst,
							},
						},
					},
				},
				SchedulePolicy: recoveryservices.SimpleSchedulePolicy{
					HourlySchedule: recoveryservices.HourlySchedule{
						Interval:                4,
						ScheduleWindowDuration:  12,
						ScheduleWindowStartTime: "2021-09-29T08:00:00.000Z",
					},
					SchedulePolicyType:   "SimpleSchedulePolicy",
					ScheduleRunFrequency: "Hourly",
				},
				TimeZone:     "UTC",
				WorkLoadType: "AzureFileShare",
			},
			ResourceGroupName: pulumi.String("SwaggerTestRg"),
			VaultName:         pulumi.String("swaggertestvault"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
