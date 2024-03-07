package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azurestackhci/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azurestackhci.NewMarketplaceGalleryImage(ctx, "marketplaceGalleryImage", &azurestackhci.MarketplaceGalleryImageArgs{
			CloudInitDataSource: pulumi.String("Azure"),
			ContainerName:       pulumi.String("Default_Container"),
			ExtendedLocation: &azurestackhci.ExtendedLocationArgs{
				Name: pulumi.String("/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location"),
				Type: pulumi.String("CustomLocation"),
			},
			HyperVGeneration: pulumi.String("V1"),
			Identifier: &azurestackhci.GalleryImageIdentifierArgs{
				Offer:     pulumi.String("myOfferName"),
				Publisher: pulumi.String("myPublisherName"),
				Sku:       pulumi.String("mySkuName"),
			},
			Location:                    pulumi.String("West US2"),
			MarketplaceGalleryImageName: pulumi.String("test-marketplace-gallery-image"),
			OsType:                      azurestackhci.OperatingSystemTypesWindows,
			ResourceGroupName:           pulumi.String("test-rg"),
			Version: &azurestackhci.GalleryImageVersionArgs{
				Name: pulumi.String("1.0.0"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
