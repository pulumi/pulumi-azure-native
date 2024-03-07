package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/costmanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := costmanagement.NewScheduledActionByScope(ctx, "scheduledActionByScope", &costmanagement.ScheduledActionByScopeArgs{
			DisplayName: pulumi.String("Daily anomaly by resource"),
			Kind:        pulumi.String("InsightAlert"),
			Name:        pulumi.String("dailyAnomalyByResource"),
			Notification: &costmanagement.NotificationPropertiesArgs{
				Subject: pulumi.String("Cost anomaly detected in the resource"),
				To: pulumi.StringArray{
					pulumi.String("user@gmail.com"),
					pulumi.String("team@gmail.com"),
				},
			},
			Schedule: &costmanagement.SchedulePropertiesArgs{
				EndDate:   pulumi.String("2021-06-19T22:21:51.1287144Z"),
				Frequency: pulumi.String("Daily"),
				StartDate: pulumi.String("2020-06-19T22:21:51.1287144Z"),
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
