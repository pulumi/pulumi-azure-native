package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewManagerExtendedInfo(ctx, "managerExtendedInfo", &storsimple.ManagerExtendedInfoArgs{
			Algorithm:         pulumi.String("None"),
			IntegrityKey:      pulumi.String("BIl+RHqO8PZ6DRvuXTTK7g=="),
			ManagerName:       pulumi.String("ManagerForSDKTest2"),
			ResourceGroupName: pulumi.String("ResourceGroupForSDKTest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
