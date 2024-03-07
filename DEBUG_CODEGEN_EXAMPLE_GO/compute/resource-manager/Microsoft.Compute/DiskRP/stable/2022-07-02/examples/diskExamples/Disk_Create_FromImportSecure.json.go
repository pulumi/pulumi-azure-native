package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
			CreationData: &compute.CreationDataArgs{
				CreateOption:     pulumi.String("ImportSecure"),
				SecurityDataUri:  pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/vmgs.vhd"),
				SourceUri:        pulumi.String("https://mystorageaccount.blob.core.windows.net/osimages/osimage.vhd"),
				StorageAccountId: pulumi.String("subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Storage/storageAccounts/myStorageAccount"),
			},
			DiskName:          pulumi.String("myDisk"),
			Location:          pulumi.String("West US"),
			OsType:            compute.OperatingSystemTypesWindows,
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SecurityProfile: &compute.DiskSecurityProfileArgs{
				SecurityType: pulumi.String("ConfidentialVM_VMGuestStateOnlyEncryptedWithPlatformKey"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
