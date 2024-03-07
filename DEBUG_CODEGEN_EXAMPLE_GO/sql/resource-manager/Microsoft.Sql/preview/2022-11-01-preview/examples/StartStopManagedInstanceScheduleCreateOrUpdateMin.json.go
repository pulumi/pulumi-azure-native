package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/sql/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := sql.NewStartStopManagedInstanceSchedule(ctx, "startStopManagedInstanceSchedule", &sql.StartStopManagedInstanceScheduleArgs{
			ManagedInstanceName: pulumi.String("schedulemi"),
			ResourceGroupName:   pulumi.String("schedulerg"),
			ScheduleList: sql.ScheduleItemArray{
				&sql.ScheduleItemArgs{
					StartDay:  pulumi.String("Thursday"),
					StartTime: pulumi.String("18:00"),
					StopDay:   pulumi.String("Thursday"),
					StopTime:  pulumi.String("17:00"),
				},
				&sql.ScheduleItemArgs{
					StartDay:  pulumi.String("Thursday"),
					StartTime: pulumi.String("15:00"),
					StopDay:   pulumi.String("Thursday"),
					StopTime:  pulumi.String("14:00"),
				},
			},
			StartStopScheduleName: pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
