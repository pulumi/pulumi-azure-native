package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewSchedule(ctx, "schedule", &devcenter.ScheduleArgs{
			Frequency:         pulumi.String("Daily"),
			PoolName:          pulumi.String("DevPool"),
			ProjectName:       pulumi.String("DevProject"),
			ResourceGroupName: pulumi.String("rg1"),
			ScheduleName:      pulumi.String("autoShutdown"),
			State:             pulumi.String("Enabled"),
			Time:              pulumi.String("17:30"),
			TimeZone:          pulumi.String("America/Los_Angeles"),
			Type:              pulumi.String("StopDevBox"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
