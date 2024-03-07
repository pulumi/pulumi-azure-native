package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewScheduledActionByScope(ctx, "scheduledActionByScope", &costmanagement.ScheduledActionByScopeArgs{
			DisplayName: pulumi.String("Monthly Cost By Resource"),
			FileDestination: &costmanagement.FileDestinationArgs{
				FileFormats: pulumi.StringArray{
					pulumi.String("Csv"),
				},
			},
			Kind: pulumi.String("Email"),
			Name: pulumi.String("monthlyCostByResource"),
			Notification: &costmanagement.NotificationPropertiesArgs{
				Subject: pulumi.String("Cost by resource this month"),
				To: pulumi.StringArray{
					pulumi.String("user@gmail.com"),
					pulumi.String("team@gmail.com"),
				},
			},
			Schedule: &costmanagement.SchedulePropertiesArgs{
				DaysOfWeek: pulumi.StringArray{
					pulumi.String("Monday"),
				},
				EndDate:   pulumi.String("2021-06-19T22:21:51.1287144Z"),
				Frequency: pulumi.String("Monthly"),
				HourOfDay: pulumi.Int(10),
				StartDate: pulumi.String("2020-06-19T22:21:51.1287144Z"),
				WeeksOfMonth: pulumi.StringArray{
					pulumi.String("First"),
					pulumi.String("Third"),
				},
			},
			Scope:  pulumi.String("subscriptions/00000000-0000-0000-0000-000000000000"),
			Status: pulumi.String("Enabled"),
			ViewId: pulumi.String("/providers/Microsoft.CostManagement/views/swaggerExample"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
