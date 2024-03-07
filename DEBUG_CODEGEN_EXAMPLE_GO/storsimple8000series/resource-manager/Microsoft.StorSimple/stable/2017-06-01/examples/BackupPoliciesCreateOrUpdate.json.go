package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewBackupPolicy(ctx, "backupPolicy", &storsimple.BackupPolicyArgs{
			BackupPolicyName:  pulumi.String("BkUpPolicy01ForSDKTest"),
			DeviceName:        pulumi.String("Device05ForSDKTest"),
			Kind:              storsimple.KindSeries8000,
			ManagerName:       pulumi.String("ManagerForSDKTest1"),
			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
			VolumeIds: pulumi.StringArray{
				pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/Clonedvolume1"),
				pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/devices/Device05ForSDKTest/volumeContainers/volumeContainerForSDKTest/volumes/volume1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
