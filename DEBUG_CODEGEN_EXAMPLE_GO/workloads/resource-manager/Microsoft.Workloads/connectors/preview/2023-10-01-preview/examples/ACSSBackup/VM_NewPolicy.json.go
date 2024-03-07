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
BackupManagementType: "AzureIaasVM",
InstantRPDetails: workloads.InstantRPAdditionalDetails{
AzureBackupRGNamePrefix: "dasas",
AzureBackupRGNameSuffix: "a",
},
InstantRpRetentionRangeInDays: 2,
Name: "defaultVmPolicy",
PolicyType: "V1",
ProtectedItemsCount: 0,
RetentionPolicy: workloads.LongTermRetentionPolicy{
DailySchedule: workloads.DailyRetentionSchedule{
RetentionDuration: workloads.RetentionDuration{
Count: 30,
DurationType: "Days",
},
RetentionTimes: []string{
"2022-11-29T19:30:00.000Z",
},
},
MonthlySchedule: workloads.MonthlyRetentionSchedule{
RetentionDuration: workloads.RetentionDuration{
Count: 60,
DurationType: "Months",
},
RetentionScheduleFormatType: "Weekly",
RetentionScheduleWeekly: workloads.WeeklyRetentionFormat{
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
WeeklySchedule: workloads.WeeklyRetentionSchedule{
DaysOfTheWeek: []workloads.DayOfWeek{
workloads.DayOfWeekSunday,
},
RetentionDuration: workloads.RetentionDuration{
Count: 12,
DurationType: "Weeks",
},
RetentionTimes: []string{
"2022-11-29T19:30:00.000Z",
},
},
YearlySchedule: workloads.YearlyRetentionSchedule{
MonthsOfYear: []workloads.MonthOfYear{
workloads.MonthOfYearJanuary,
},
RetentionDuration: workloads.RetentionDuration{
Count: 10,
DurationType: "Years",
},
RetentionScheduleFormatType: "Weekly",
RetentionScheduleWeekly: workloads.WeeklyRetentionFormat{
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
SchedulePolicy: workloads.SimpleSchedulePolicy{
SchedulePolicyType: "SimpleSchedulePolicy",
ScheduleRunFrequency: "Daily",
ScheduleRunTimes: []string{
"2022-11-29T19:30:00.000Z",
},
},
TieringPolicy: interface{}{
ArchivedRP: workloads.TieringPolicy{
Duration: 3,
DurationType: "Months",
TieringMode: "TierAfter",
},
},
TimeZone: "UTC",
},
BackupType: "VM",
DiskExclusionProperties: workloads.DiskExclusionProperties{
DiskLunList: []interface{}{
},
IsInclusionList: true,
},
RecoveryServicesVault: workloads.NewRecoveryServicesVault{
Name: "test-vault",
ResourceGroup: "test-rg",
VaultType: "New",
},
},
BackupName: pulumi.String("vmBackup"),
ConnectorName: pulumi.String("C1"),
Location: pulumi.String("westcentralus"),
ResourceGroupName: pulumi.String("test-rg"),
Tags: nil,
})
if err != nil {
return err
}
return nil
})
}
