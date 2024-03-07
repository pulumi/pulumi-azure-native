package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewScheduledTrigger(ctx, "scheduledTrigger", &datashare.ScheduledTriggerArgs{
			AccountName:           pulumi.String("Account1"),
			Kind:                  pulumi.String("ScheduleBased"),
			RecurrenceInterval:    pulumi.String("Day"),
			ResourceGroupName:     pulumi.String("SampleResourceGroup"),
			ShareSubscriptionName: pulumi.String("ShareSubscription1"),
			SynchronizationMode:   pulumi.String("Incremental"),
			SynchronizationTime:   pulumi.String("2018-11-14T04:47:52.9614956Z"),
			TriggerName:           pulumi.String("Trigger1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
