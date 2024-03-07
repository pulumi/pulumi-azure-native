package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/storsimple/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := storsimple.NewAccessControlRecord(ctx, "accessControlRecord", &storsimple.AccessControlRecordArgs{
			AccessControlRecordName: pulumi.String("ACRForTest"),
			InitiatorName:           pulumi.String("iqn.2017-06.com.contoso:ForTest"),
			ManagerName:             pulumi.String("ManagerForSDKTest1"),
			ResourceGroupName:       pulumi.String("ResourceGroupForSDKTest"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
