package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/labservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := labservices.NewSchedule(ctx, "schedule", &labservices.ScheduleArgs{
			LabName: pulumi.String("testlab"),
			Notes:   pulumi.String("Schedule 1 for students"),
			RecurrencePattern: &labservices.RecurrencePatternArgs{
				ExpirationDate: pulumi.String("2020-08-14T23:59:59Z"),
				Frequency:      labservices.RecurrenceFrequencyDaily,
				Interval:       pulumi.Int(2),
			},
			ResourceGroupName: pulumi.String("testrg123"),
			ScheduleName:      pulumi.String("schedule1"),
			StartAt:           pulumi.String("2020-05-26T12:00:00Z"),
			StopAt:            pulumi.String("2020-05-26T18:00:00Z"),
			TimeZoneId:        pulumi.String("America/Los_Angeles"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
