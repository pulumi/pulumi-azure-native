package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewDisk(ctx, "disk", &compute.DiskArgs{
			CreationData: compute.CreationDataResponse{
				CreateOption: pulumi.String("FromImage"),
				ImageReference: &compute.ImageDiskReferenceArgs{
					Id: pulumi.String("/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/westus/Publishers/{publisher}/ArtifactTypes/VMImage/Offers/{offer}/Skus/{sku}/Versions/1.0.0"),
				},
			},
			DiskName:          pulumi.String("myDisk"),
			Location:          pulumi.String("West US"),
			OsType:            compute.OperatingSystemTypesWindows,
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SecurityProfile: &compute.DiskSecurityProfileArgs{
				SecureVMDiskEncryptionSetId: pulumi.String("/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{diskEncryptionSetName}"),
				SecurityType:                pulumi.String("ConfidentialVM_DiskEncryptedWithCustomerKey"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
