package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewScheduledSynchronizationSetting(ctx, "scheduledSynchronizationSetting", &datashare.ScheduledSynchronizationSettingArgs{
			AccountName:                pulumi.String("Account1"),
			Kind:                       pulumi.String("ScheduleBased"),
			RecurrenceInterval:         pulumi.String("Day"),
			ResourceGroupName:          pulumi.String("SampleResourceGroup"),
			ShareName:                  pulumi.String("Share1"),
			SynchronizationSettingName: pulumi.String("Dataset1"),
			SynchronizationTime:        pulumi.String("2018-11-14T04:47:52.9614956Z"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
