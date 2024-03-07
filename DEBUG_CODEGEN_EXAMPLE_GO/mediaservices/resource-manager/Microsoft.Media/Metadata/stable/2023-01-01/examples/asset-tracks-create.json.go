package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewTrack(ctx, "track", &media.TrackArgs{
			AccountName:       pulumi.String("contosomedia"),
			AssetName:         pulumi.String("ClimbingMountRainer"),
			ResourceGroupName: pulumi.String("contosorg"),
			Track: media.TextTrack{
				DisplayName:      "A new track",
				FileName:         "text3.ttml",
				OdataType:        "#Microsoft.Media.TextTrack",
				PlayerVisibility: "Visible",
			},
			TrackName: pulumi.String("text3"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
