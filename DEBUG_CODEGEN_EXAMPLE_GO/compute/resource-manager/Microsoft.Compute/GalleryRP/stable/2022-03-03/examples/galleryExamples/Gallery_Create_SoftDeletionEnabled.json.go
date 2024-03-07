package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewGallery(ctx, "gallery", &compute.GalleryArgs{
			Description:       pulumi.String("This is the gallery description."),
			GalleryName:       pulumi.String("myGalleryName"),
			Location:          pulumi.String("West US"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			SoftDeletePolicy: &compute.SoftDeletePolicyArgs{
				IsSoftDeleteEnabled: pulumi.Bool(true),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
