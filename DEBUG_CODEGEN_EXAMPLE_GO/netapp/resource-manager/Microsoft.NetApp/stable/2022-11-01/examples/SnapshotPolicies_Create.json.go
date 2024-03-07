package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/netapp/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := netapp.NewSnapshotPolicy(ctx, "snapshotPolicy", &netapp.SnapshotPolicyArgs{
			AccountName: pulumi.String("account1"),
			DailySchedule: &netapp.DailyScheduleArgs{
				Hour:            pulumi.Int(14),
				Minute:          pulumi.Int(30),
				SnapshotsToKeep: pulumi.Int(4),
			},
			Enabled: pulumi.Bool(true),
			HourlySchedule: &netapp.HourlyScheduleArgs{
				Minute:          pulumi.Int(50),
				SnapshotsToKeep: pulumi.Int(2),
			},
			Location: pulumi.String("eastus"),
			MonthlySchedule: &netapp.MonthlyScheduleArgs{
				DaysOfMonth:     pulumi.String("10,11,12"),
				Hour:            pulumi.Int(14),
				Minute:          pulumi.Int(15),
				SnapshotsToKeep: pulumi.Int(5),
			},
			ResourceGroupName:  pulumi.String("myRG"),
			SnapshotPolicyName: pulumi.String("snapshotPolicyName"),
			WeeklySchedule: &netapp.WeeklyScheduleArgs{
				Day:             pulumi.String("Wednesday"),
				Hour:            pulumi.Int(14),
				Minute:          pulumi.Int(45),
				SnapshotsToKeep: pulumi.Int(3),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
