package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewServiceFabricSchedule(ctx, "serviceFabricSchedule", &devtestlab.ServiceFabricScheduleArgs{
			DailyRecurrence: &devtestlab.DayDetailsArgs{
				Time: pulumi.String("19:00"),
			},
			HourlyRecurrence: &devtestlab.HourDetailsArgs{
				Minute: pulumi.Int(0),
			},
			LabName:  pulumi.String("{labName}"),
			Location: pulumi.String("{location}"),
			Name:     pulumi.String("{scheduleName}"),
			NotificationSettings: &devtestlab.NotificationSettingsArgs{
				EmailRecipient:     pulumi.String("{email}"),
				NotificationLocale: pulumi.String("EN"),
				Status:             pulumi.String("{Enabled|Disabled}"),
				TimeInMinutes:      pulumi.Int(15),
				WebhookUrl:         pulumi.String("{webhoolUrl}"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			ServiceFabricName: pulumi.String("{serviceFrabicName}"),
			Status:            pulumi.String("{Enabled|Disabled}"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			TargetResourceId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/resourceGroupName/providers/Microsoft.DevTestLab/labs/{labName}/users/{uniqueIdentifier}/servicefabrics/{serviceFrabicName}"),
			TaskType:         pulumi.String("{Unknown|LabVmsShutdownTask|LabVmsStartupTask|LabVmReclamationTask|ComputeVmShutdownTask}"),
			TimeZoneId:       pulumi.String("Pacific Standard Time"),
			UserName:         pulumi.String("@me"),
			WeeklyRecurrence: &devtestlab.WeekDetailsArgs{
				Time: pulumi.String("19:00"),
				Weekdays: pulumi.StringArray{
					pulumi.String("Monday"),
					pulumi.String("Tuesday"),
					pulumi.String("Wednesday"),
					pulumi.String("Thursday"),
					pulumi.String("Friday"),
					pulumi.String("Saturday"),
					pulumi.String("Sunday"),
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
