package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewGlobalSchedule(ctx, "globalSchedule", &devtestlab.GlobalScheduleArgs{
			Name:              pulumi.String("labvmautostart"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Status:            pulumi.String("Enabled"),
			TaskType:          pulumi.String("LabVmsStartupTask"),
			TimeZoneId:        pulumi.String("Hawaiian Standard Time"),
			WeeklyRecurrence: &devtestlab.WeekDetailsArgs{
				Time: pulumi.String("0700"),
				Weekdays: pulumi.StringArray{
					pulumi.String("Monday"),
					pulumi.String("Tuesday"),
					pulumi.String("Wednesday"),
					pulumi.String("Thursday"),
					pulumi.String("Friday"),
					pulumi.String("Saturday"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
