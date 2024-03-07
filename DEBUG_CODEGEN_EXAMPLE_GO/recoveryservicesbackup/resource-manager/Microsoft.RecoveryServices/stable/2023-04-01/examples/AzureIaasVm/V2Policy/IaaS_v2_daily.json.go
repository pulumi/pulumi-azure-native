package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionPolicy(ctx, "protectionPolicy", &recoveryservices.ProtectionPolicyArgs{
			PolicyName: pulumi.String("v2-daily-sample"),
			Properties: recoveryservices.AzureIaaSVMProtectionPolicy{
				BackupManagementType:          "AzureIaasVM",
				InstantRpRetentionRangeInDays: 30,
				PolicyType:                    "V2",
				RetentionPolicy: recoveryservices.LongTermRetentionPolicy{
					DailySchedule: recoveryservices.DailyRetentionSchedule{
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        180,
							DurationType: "Days",
						},
						RetentionTimes: []string{
							"2021-12-17T08:00:00+00:00",
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
						RetentionTimes: []string{
							"2021-12-17T08:00:00+00:00",
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
						RetentionTimes: []string{
							"2021-12-17T08:00:00+00:00",
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
						RetentionTimes: []string{
							"2021-12-17T08:00:00+00:00",
						},
					},
				},
				SchedulePolicy: recoveryservices.SimpleSchedulePolicyV2{
					DailySchedule: recoveryservices.DailySchedule{
						ScheduleRunTimes: []string{
							"2018-01-24T10:00:00Z",
						},
					},
					SchedulePolicyType:   "SimpleSchedulePolicyV2",
					ScheduleRunFrequency: "Daily",
				},
				TimeZone: "India Standard Time",
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
