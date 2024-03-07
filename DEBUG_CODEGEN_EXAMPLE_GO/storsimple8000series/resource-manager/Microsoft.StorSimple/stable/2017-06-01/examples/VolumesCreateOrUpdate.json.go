package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewVolume(ctx, "volume", &storsimple.VolumeArgs{
			AccessControlRecordIds: pulumi.StringArray{
				pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/accessControlRecords/ACR2"),
			},
			DeviceName:          pulumi.String("Device05ForSDKTest"),
			ManagerName:         pulumi.String("ManagerForSDKTest1"),
			MonitoringStatus:    storsimple.MonitoringStatusEnabled,
			ResourceGroupName:   pulumi.String("ResourceGroupForSDKTest"),
			SizeInBytes:         pulumi.Float64(5368709120),
			VolumeContainerName: pulumi.String("VolumeContainerForSDKTest"),
			VolumeName:          pulumi.String("Volume1ForSDKTest"),
			VolumeStatus:        storsimple.VolumeStatusOffline,
			VolumeType:          storsimple.VolumeTypeTiered,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
