package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewBandwidthSchedule(ctx, "bandwidthSchedule", &databoxedge.BandwidthScheduleArgs{
			Days: pulumi.StringArray{
				pulumi.String("Sunday"),
				pulumi.String("Monday"),
			},
			DeviceName:        pulumi.String("testedgedevice"),
			Name:              pulumi.String("bandwidth-1"),
			RateInMbps:        pulumi.Int(100),
			ResourceGroupName: pulumi.String("GroupForEdgeAutomation"),
			Start:             pulumi.String("0:0:0"),
			Stop:              pulumi.String("13:59:0"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
