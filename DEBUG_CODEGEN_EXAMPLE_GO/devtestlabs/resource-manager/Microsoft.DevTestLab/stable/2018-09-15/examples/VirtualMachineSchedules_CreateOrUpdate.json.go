package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devtestlab/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devtestlab.NewVirtualMachineSchedule(ctx, "virtualMachineSchedule", &devtestlab.VirtualMachineScheduleArgs{
			DailyRecurrence: &devtestlab.DayDetailsArgs{
				Time: pulumi.String("1900"),
			},
			HourlyRecurrence: &devtestlab.HourDetailsArgs{
				Minute: pulumi.Int(30),
			},
			LabName:  pulumi.String("{labName}"),
			Location: pulumi.String("{location}"),
			Name:     pulumi.String("LabVmsShutdown"),
			NotificationSettings: &devtestlab.NotificationSettingsArgs{
				EmailRecipient:     pulumi.String("{email}"),
				NotificationLocale: pulumi.String("EN"),
				Status:             pulumi.String("Enabled"),
				TimeInMinutes:      pulumi.Int(30),
				WebhookUrl:         pulumi.String("{webhookUrl}"),
			},
			ResourceGroupName: pulumi.String("resourceGroupName"),
			Status:            pulumi.String("Enabled"),
			Tags: pulumi.StringMap{
				"tagName1": pulumi.String("tagValue1"),
			},
			TargetResourceId:   pulumi.String("/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualMachines/{vmName}"),
			TaskType:           pulumi.String("LabVmsShutdownTask"),
			TimeZoneId:         pulumi.String("Pacific Standard Time"),
			VirtualMachineName: pulumi.String("{vmName}"),
			WeeklyRecurrence: &devtestlab.WeekDetailsArgs{
				Time: pulumi.String("1700"),
				Weekdays: pulumi.StringArray{
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
