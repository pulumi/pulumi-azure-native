package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/devcenter/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := devcenter.NewGallery(ctx, "gallery", &devcenter.GalleryArgs{
			DevCenterName:     pulumi.String("Contoso"),
			GalleryName:       pulumi.String("StandardGallery"),
			GalleryResourceId: pulumi.String("/subscriptions/0ac520ee-14c0-480f-b6c9-0a90c58ffff/resourceGroups/rg1/providers/Microsoft.Compute/galleries/StandardGallery"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
