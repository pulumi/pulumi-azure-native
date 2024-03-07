package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewGalleryImage(ctx, "galleryImage", &compute.GalleryImageArgs{
			GalleryImageName: pulumi.String("myGalleryImageName"),
			GalleryName:      pulumi.String("myGalleryName"),
			HyperVGeneration: pulumi.String("V1"),
			Identifier: &compute.GalleryImageIdentifierArgs{
				Offer:     pulumi.String("myOfferName"),
				Publisher: pulumi.String("myPublisherName"),
				Sku:       pulumi.String("mySkuName"),
			},
			Location:          pulumi.String("West US"),
			OsState:           compute.OperatingSystemStateTypesGeneralized,
			OsType:            compute.OperatingSystemTypesWindows,
			ResourceGroupName: pulumi.String("myResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
