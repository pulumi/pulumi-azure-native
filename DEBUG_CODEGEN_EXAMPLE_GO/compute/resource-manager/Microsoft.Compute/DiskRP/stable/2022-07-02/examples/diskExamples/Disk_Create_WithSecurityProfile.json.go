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
					Id: pulumi.String("/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}"),
				},
			},
			DiskName:          pulumi.String("myDisk"),
			Location:          pulumi.String("North Central US"),
			OsType:            compute.OperatingSystemTypesWindows,
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SecurityProfile: &compute.DiskSecurityProfileArgs{
				SecurityType: pulumi.String("TrustedLaunch"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
