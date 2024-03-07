package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewPatchSchedule(ctx, "patchSchedule", &cache.PatchScheduleArgs{
			Default:           pulumi.String("default"),
			Name:              pulumi.String("cache1"),
			ResourceGroupName: pulumi.String("rg1"),
			ScheduleEntries: []cache.ScheduleEntryArgs{
				{
					DayOfWeek:         cache.DayOfWeekMonday,
					MaintenanceWindow: pulumi.String("PT5H"),
					StartHourUtc:      pulumi.Int(12),
				},
				{
					DayOfWeek:    cache.DayOfWeekTuesday,
					StartHourUtc: pulumi.Int(12),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
