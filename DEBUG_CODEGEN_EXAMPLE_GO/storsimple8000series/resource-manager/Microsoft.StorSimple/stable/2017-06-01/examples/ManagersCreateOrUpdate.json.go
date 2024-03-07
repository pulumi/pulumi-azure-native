package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewManager(ctx, "manager", &storsimple.ManagerArgs{
			CisIntrinsicSettings: &storsimple.ManagerIntrinsicSettingsArgs{
				Type: storsimple.ManagerTypeGardaV1,
			},
			Location:          pulumi.String("westus"),
			ManagerName:       pulumi.String("ManagerForSDKTest2"),
			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
			Sku: &storsimple.ManagerSkuArgs{
				Name: storsimple.ManagerSkuTypeStandard,
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
