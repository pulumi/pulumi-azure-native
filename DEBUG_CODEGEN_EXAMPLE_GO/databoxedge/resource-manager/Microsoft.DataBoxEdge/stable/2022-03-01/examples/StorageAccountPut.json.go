package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/databoxedge/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := databoxedge.NewStorageAccount(ctx, "storageAccount", &databoxedge.StorageAccountArgs{
			DataPolicy:                 pulumi.String("Cloud"),
			Description:                pulumi.String("It's an awesome storage account"),
			DeviceName:                 pulumi.String("testedgedevice"),
			ResourceGroupName:          pulumi.String("GroupForEdgeAutomation"),
			StorageAccountCredentialId: pulumi.String("/subscriptions/4385cf00-2d3a-425a-832f-f4285b1c9dce/resourceGroups/GroupForDataBoxEdgeAutomation/providers/Microsoft.DataBoxEdge/dataBoxEdgeDevices/testedgedevice/storageAccountCredentials/cisbvt"),
			StorageAccountName:         pulumi.String("blobstorageaccount1"),
			StorageAccountStatus:       pulumi.String("OK"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
