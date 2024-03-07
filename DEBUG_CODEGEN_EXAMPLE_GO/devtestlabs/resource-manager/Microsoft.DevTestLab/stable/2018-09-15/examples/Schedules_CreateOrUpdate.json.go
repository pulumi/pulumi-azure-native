package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewSchedule(ctx, "schedule", &devtestlab.ScheduleArgs{
			DailyRecurrence: &devtestlab.DayDetailsArgs{
				Time: pulumi.String("{timeOfTheDayTheScheduleWillOccurEveryDay}"),
			},
			HourlyRecurrence: &devtestlab.HourDetailsArgs{
				Minute: pulumi.Int(30),
			},
			LabName:  pulumi.String("{labName}"),
			Location: pulumi.String("{location}"),
			Name:     pulumi.String("{scheduleName}"),
			NotificationSettings: &devtestlab.NotificationSettingsArgs{
				EmailRecipient:     pulumi.String("{email}"),
				NotificationLocale: pulumi.String("EN"),
				Status:             pulumi.String("{Enabled|Disabled}"),
				TimeInMinutes:      pulumi.Int(15),
				WebhookUrl:         pulumi.String("{webhookUrl}"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Status:            pulumi.String("{Enabled|Disabled}"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			TargetResourceId: pulumi.String("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}"),
			TaskType:         pulumi.String("{myLabVmTaskType}"),
			TimeZoneId:       pulumi.String("Pacific Standard Time"),
			WeeklyRecurrence: &devtestlab.WeekDetailsArgs{
				Time: pulumi.String("{timeOfTheDayTheScheduleWillOccurOnThoseDays}"),
				Weekdays: pulumi.StringArray{
					pulumi.String("Monday"),
					pulumi.String("Wednesday"),
					pulumi.String("Friday"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
