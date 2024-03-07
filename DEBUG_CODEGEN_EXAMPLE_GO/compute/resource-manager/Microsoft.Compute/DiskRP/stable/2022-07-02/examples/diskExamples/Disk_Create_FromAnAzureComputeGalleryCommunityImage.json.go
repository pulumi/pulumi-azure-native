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
				GalleryImageReference: &compute.ImageDiskReferenceArgs{
					CommunityGalleryImageId: pulumi.String("/CommunityGalleries/{communityGalleryPublicGalleryName}/Images/{imageName}/Versions/1.0.0"),
				},
			},
			DiskName:          pulumi.String("myDisk"),
			Location:          pulumi.String("West US"),
			OsType:            compute.OperatingSystemTypesWindows,
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
