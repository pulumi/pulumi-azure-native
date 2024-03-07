package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewBackupSchedule(ctx, "backupSchedule", &storsimple.BackupScheduleArgs{
			BackupPolicyName:   pulumi.String("BkUpPolicy01ForSDKTest"),
			BackupScheduleName: pulumi.String("schedule2"),
			BackupType:         storsimple.BackupTypeCloudSnapshot,
			DeviceName:         pulumi.String("Device05ForSDKTest"),
			Kind:               storsimple.KindSeries8000,
			ManagerName:        pulumi.String("ManagerForSDKTest1"),
			ResourceGroupName:  pulumi.String("ResourceGroupForSDKTest"),
			RetentionCount:     pulumi.Float64(1),
			ScheduleRecurrence: &storsimple.ScheduleRecurrenceArgs{
				RecurrenceType:  storsimple.RecurrenceTypeWeekly,
				RecurrenceValue: pulumi.Int(1),
				WeeklyDaysList: storsimple.DayOfWeekArray{
					storsimple.DayOfWeekFriday,
					storsimple.DayOfWeekThursday,
					storsimple.DayOfWeekMonday,
				},
			},
			ScheduleStatus: storsimple.ScheduleStatusEnabled,
			StartTime:      pulumi.String("2017-06-24T01:00:00Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
