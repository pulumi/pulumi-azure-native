package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewPeriodicTimerEventTrigger(ctx, "periodicTimerEventTrigger", &databoxedge.PeriodicTimerEventTriggerArgs{
			DeviceName:        pulumi.String("testedgedevice"),
			Name:              pulumi.String("trigger1"),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
