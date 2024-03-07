package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/compute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := compute.NewGalleryApplication(ctx, "galleryApplication", &compute.GalleryApplicationArgs{
			CustomActions: []compute.GalleryApplicationCustomActionArgs{
				{
					Description: pulumi.String("This is the custom action description."),
					Name:        pulumi.String("myCustomAction"),
					Parameters: compute.GalleryApplicationCustomActionParameterArray{
						{
							DefaultValue: pulumi.String("default value of parameter."),
							Description:  pulumi.String("This is the description of the parameter"),
							Name:         pulumi.String("myCustomActionParameter"),
							Required:     pulumi.Bool(false),
							Type:         compute.GalleryApplicationCustomActionParameterTypeString,
						},
					},
					Script: pulumi.String("myCustomActionScript"),
				},
			},
			Description:            pulumi.String("This is the gallery application description."),
			Eula:                   pulumi.String("This is the gallery application EULA."),
			GalleryApplicationName: pulumi.String("myGalleryApplicationName"),
			GalleryName:            pulumi.String("myGalleryName"),
			Location:               pulumi.String("West US"),
			PrivacyStatementUri:    pulumi.String("myPrivacyStatementUri}"),
			ReleaseNoteUri:         pulumi.String("myReleaseNoteUri"),
			ResourceGroupName:      pulumi.String("myResourceGroup"),
			SupportedOSType:        compute.OperatingSystemTypesWindows,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
