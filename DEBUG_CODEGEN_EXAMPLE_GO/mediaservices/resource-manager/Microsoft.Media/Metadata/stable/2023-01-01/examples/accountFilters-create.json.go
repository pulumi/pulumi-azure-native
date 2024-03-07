package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewAccountFilter(ctx, "accountFilter", &media.AccountFilterArgs{
			AccountName: pulumi.String("contosomedia"),
			FilterName:  pulumi.String("newAccountFilter"),
			FirstQuality: &media.FirstQualityArgs{
				Bitrate: pulumi.Int(128000),
			},
			PresentationTimeRange: &media.PresentationTimeRangeArgs{
				EndTimestamp:               pulumi.Float64(170000000),
				ForceEndTimestamp:          pulumi.Bool(false),
				LiveBackoffDuration:        pulumi.Float64(0),
				PresentationWindowDuration: pulumi.Float64(9223372036854774784),
				StartTimestamp:             pulumi.Float64(0),
				Timescale:                  pulumi.Float64(10000000),
			},
			ResourceGroupName: pulumi.String("contosorg"),
			Tracks: media.FilterTrackSelectionArray{
				&media.FilterTrackSelectionArgs{
					TrackSelections: media.FilterTrackPropertyConditionArray{
						&media.FilterTrackPropertyConditionArgs{
							Operation: pulumi.String("Equal"),
							Property:  pulumi.String("Type"),
							Value:     pulumi.String("Audio"),
						},
						&media.FilterTrackPropertyConditionArgs{
							Operation: pulumi.String("NotEqual"),
							Property:  pulumi.String("Language"),
							Value:     pulumi.String("en"),
						},
						&media.FilterTrackPropertyConditionArgs{
							Operation: pulumi.String("NotEqual"),
							Property:  pulumi.String("FourCC"),
							Value:     pulumi.String("EC-3"),
						},
					},
				},
				&media.FilterTrackSelectionArgs{
					TrackSelections: media.FilterTrackPropertyConditionArray{
						&media.FilterTrackPropertyConditionArgs{
							Operation: pulumi.String("Equal"),
							Property:  pulumi.String("Type"),
							Value:     pulumi.String("Video"),
						},
						&media.FilterTrackPropertyConditionArgs{
							Operation: pulumi.String("Equal"),
							Property:  pulumi.String("Bitrate"),
							Value:     pulumi.String("3000000-5000000"),
						},
					},
				},
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
