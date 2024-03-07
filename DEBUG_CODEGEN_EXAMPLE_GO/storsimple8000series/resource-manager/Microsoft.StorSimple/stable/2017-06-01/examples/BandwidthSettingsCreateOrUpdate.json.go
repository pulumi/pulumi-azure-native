package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewBandwidthSetting(ctx, "bandwidthSetting", &storsimple.BandwidthSettingArgs{
			BandwidthSettingName: pulumi.String("BWSForTest"),
			ManagerName:          pulumi.String("ManagerForSDKTest1"),
			ResourceGroupName:    pulumi.String("ResourceGroupForSDKTest"),
			Schedules: []storsimple.BandwidthScheduleArgs{
				{
					Days: storsimple.DayOfWeekArray{
						storsimple.DayOfWeekSaturday,
						storsimple.DayOfWeekSunday,
					},
					RateInMbps: pulumi.Int(10),
					Start: {
						Hours:   pulumi.Int(10),
						Minutes: pulumi.Int(0),
						Seconds: pulumi.Int(0),
					},
					Stop: {
						Hours:   pulumi.Int(20),
						Minutes: pulumi.Int(0),
						Seconds: pulumi.Int(0),
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
