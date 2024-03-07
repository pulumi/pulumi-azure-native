package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automation/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automation.NewSchedule(ctx, "schedule", &automation.ScheduleArgs{
			AdvancedSchedule:      nil,
			AutomationAccountName: pulumi.String("myAutomationAccount33"),
			Description:           pulumi.String("my description of schedule goes here"),
			ExpiryTime:            pulumi.String("2017-04-01T17:28:57.2494819Z"),
			Frequency:             pulumi.String("Hour"),
			Interval:              pulumi.Any(1),
			Name:                  pulumi.String("mySchedule"),
			ResourceGroupName:     pulumi.String("rg"),
			ScheduleName:          pulumi.String("mySchedule"),
			StartTime:             pulumi.String("2017-03-27T17:28:57.2494819Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
