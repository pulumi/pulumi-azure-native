package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewVolumeContainer(ctx, "volumeContainer", &storsimple.VolumeContainerArgs{
			BandwidthSettingId: pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/bandwidthSettings/bandwidthSetting1"),
			DeviceName:         pulumi.String("Device05ForSDKTest"),
			EncryptionKey: &storsimple.AsymmetricEncryptedSecretArgs{
				EncryptionAlgorithm:      storsimple.EncryptionAlgorithm_RSAES_PKCS1_v_1_5,
				EncryptionCertThumbprint: pulumi.String("A872A2DF196AC7682EE24791E7DE2E2A360F5926"),
				Value:                    pulumi.String("R//pyVLx/fn58ia098JiLgZB5RY7fVT+6o8a4fmsvjy+ls2UgJphMf25XVqEQCZnsp/5uxteN1M/9ArPIICdhM7M1+b/Ur7kJ0FH0ktxfk7CrPWWJLI4q20LZoduJGI56lREav1VpuLdqw5F9fRcq7zbfgPQ3B/SD0mfumNRiV+AnwbC6msfavIuWrhVDl9iSzEPE+zU06/kpsexnrS81yYT2QlVVUbvpY4F3zfH8TQPpAROTbv2pld6JO4eGOrZ5O1iOr6XCg2TY2W/jf+Ev4z5tqC9VWXE5kh65gjBfpWN0bDWXKekqEhor2crHAxZi4dybdY8Ok1MDWd1CSU8kw=="),
			},
			ManagerName:                pulumi.String("ManagerForSDKTest1"),
			ResourceGroupName:          pulumi.String("ResourceGroupForSDKTest"),
			StorageAccountCredentialId: pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/ResourceGroupForSDKTest/providers/Microsoft.StorSimple/managers/ManagerForSDKTest1/storageAccountCredentials/safortestrecording"),
			VolumeContainerName:        pulumi.String("VolumeContainerForSDKTest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
