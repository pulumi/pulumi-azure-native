package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/recoveryservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := recoveryservices.NewProtectionPolicy(ctx, "protectionPolicy", &recoveryservices.ProtectionPolicyArgs{
			PolicyName: pulumi.String("testPolicy1"),
			Properties: recoveryservices.AzureIaaSVMProtectionPolicy{
				BackupManagementType: "AzureIaasVM",
				RetentionPolicy: recoveryservices.LongTermRetentionPolicy{
					MonthlySchedule: recoveryservices.MonthlyRetentionSchedule{
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        2,
							DurationType: "Months",
						},
						RetentionScheduleFormatType: "Weekly",
						RetentionScheduleWeekly: recoveryservices.WeeklyRetentionFormat{
							DaysOfTheWeek: []recoveryservices.DayOfWeek{
								recoveryservices.DayOfWeekWednesday,
								recoveryservices.DayOfWeekThursday,
							},
							WeeksOfTheMonth: []recoveryservices.WeekOfMonth{
								recoveryservices.WeekOfMonthFirst,
								recoveryservices.WeekOfMonthThird,
							},
						},
						RetentionTimes: []string{
							"2018-01-24T10:00:00Z",
						},
					},
					RetentionPolicyType: "LongTermRetentionPolicy",
					WeeklySchedule: recoveryservices.WeeklyRetentionSchedule{
						DaysOfTheWeek: []recoveryservices.DayOfWeek{
							recoveryservices.DayOfWeekMonday,
							recoveryservices.DayOfWeekWednesday,
							recoveryservices.DayOfWeekThursday,
						},
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        1,
							DurationType: "Weeks",
						},
						RetentionTimes: []string{
							"2018-01-24T10:00:00Z",
						},
					},
					YearlySchedule: recoveryservices.YearlyRetentionSchedule{
						MonthsOfYear: []recoveryservices.MonthOfYear{
							recoveryservices.MonthOfYearFebruary,
							recoveryservices.MonthOfYearNovember,
						},
						RetentionDuration: recoveryservices.RetentionDuration{
							Count:        4,
							DurationType: "Years",
						},
						RetentionScheduleFormatType: "Weekly",
						RetentionScheduleWeekly: recoveryservices.WeeklyRetentionFormat{
							DaysOfTheWeek: []recoveryservices.DayOfWeek{
								recoveryservices.DayOfWeekMonday,
								recoveryservices.DayOfWeekThursday,
							},
							WeeksOfTheMonth: []recoveryservices.WeekOfMonth{
								recoveryservices.WeekOfMonthFourth,
							},
						},
						RetentionTimes: []string{
							"2018-01-24T10:00:00Z",
						},
					},
				},
				SchedulePolicy: recoveryservices.SimpleSchedulePolicy{
					SchedulePolicyType: "SimpleSchedulePolicy",
					ScheduleRunDays: []recoveryservices.DayOfWeek{
						recoveryservices.DayOfWeekMonday,
						recoveryservices.DayOfWeekWednesday,
						recoveryservices.DayOfWeekThursday,
					},
					ScheduleRunFrequency: "Weekly",
					ScheduleRunTimes: []string{
						"2018-01-24T10:00:00Z",
					},
				},
				TimeZone: "Pacific Standard Time",
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
