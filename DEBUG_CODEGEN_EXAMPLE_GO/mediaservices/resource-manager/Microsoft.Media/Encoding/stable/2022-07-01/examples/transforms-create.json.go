package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewTransform(ctx, "transform", &media.TransformArgs{
			AccountName: pulumi.String("contosomedia"),
			Description: pulumi.String("Example Transform to illustrate create and update."),
			Outputs: media.TransformOutputTypeArray{
				&media.TransformOutputTypeArgs{
					Preset: media.BuiltInStandardEncoderPreset{
						OdataType:  "#Microsoft.Media.BuiltInStandardEncoderPreset",
						PresetName: "AdaptiveStreaming",
					},
				},
			},
			ResourceGroupName: pulumi.String("contosoresources"),
			TransformName:     pulumi.String("createdTransform"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
