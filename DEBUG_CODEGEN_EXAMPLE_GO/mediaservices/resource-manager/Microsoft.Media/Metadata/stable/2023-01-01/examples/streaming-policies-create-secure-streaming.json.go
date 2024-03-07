package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewStreamingPolicy(ctx, "streamingPolicy", &media.StreamingPolicyArgs{
			AccountName: pulumi.String("contosomedia"),
			CommonEncryptionCbcs: &media.CommonEncryptionCbcsArgs{
				ContentKeys: &media.StreamingPolicyContentKeysArgs{
					DefaultKey: &media.DefaultKeyArgs{
						Label: pulumi.String("cbcsDefaultKey"),
					},
				},
				Drm: &media.CbcsDrmConfigurationArgs{
					FairPlay: &media.StreamingPolicyFairPlayConfigurationArgs{
						AllowPersistentLicense:              pulumi.Bool(true),
						CustomLicenseAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/fairplay/{ContentKeyId}"),
					},
				},
				EnabledProtocols: &media.EnabledProtocolsArgs{
					Dash:            pulumi.Bool(false),
					Download:        pulumi.Bool(false),
					Hls:             pulumi.Bool(true),
					SmoothStreaming: pulumi.Bool(false),
				},
			},
			CommonEncryptionCenc: &media.CommonEncryptionCencArgs{
				ClearTracks: media.TrackSelectionArray{
					&media.TrackSelectionArgs{
						TrackSelections: media.TrackPropertyConditionArray{
							&media.TrackPropertyConditionArgs{
								Operation: pulumi.String("Equal"),
								Property:  pulumi.String("FourCC"),
								Value:     pulumi.String("hev1"),
							},
						},
					},
				},
				ContentKeys: &media.StreamingPolicyContentKeysArgs{
					DefaultKey: &media.DefaultKeyArgs{
						Label: pulumi.String("cencDefaultKey"),
					},
				},
				Drm: &media.CencDrmConfigurationArgs{
					PlayReady: &media.StreamingPolicyPlayReadyConfigurationArgs{
						CustomLicenseAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/playready/{ContentKeyId}"),
						PlayReadyCustomAttributes:           pulumi.String("PlayReady CustomAttributes"),
					},
					Widevine: &media.StreamingPolicyWidevineConfigurationArgs{
						CustomLicenseAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/widevine/{ContentKeyId"),
					},
				},
				EnabledProtocols: &media.EnabledProtocolsArgs{
					Dash:            pulumi.Bool(true),
					Download:        pulumi.Bool(false),
					Hls:             pulumi.Bool(false),
					SmoothStreaming: pulumi.Bool(true),
				},
			},
			DefaultContentKeyPolicyName: pulumi.String("PolicyWithMultipleOptions"),
			EnvelopeEncryption: &media.EnvelopeEncryptionArgs{
				ContentKeys: &media.StreamingPolicyContentKeysArgs{
					DefaultKey: &media.DefaultKeyArgs{
						Label: pulumi.String("aesDefaultKey"),
					},
				},
				CustomKeyAcquisitionUrlTemplate: pulumi.String("https://contoso.com/{AssetAlternativeId}/envelope/{ContentKeyId}"),
				EnabledProtocols: &media.EnabledProtocolsArgs{
					Dash:            pulumi.Bool(true),
					Download:        pulumi.Bool(false),
					Hls:             pulumi.Bool(true),
					SmoothStreaming: pulumi.Bool(true),
				},
			},
			ResourceGroupName:   pulumi.String("contosorg"),
			StreamingPolicyName: pulumi.String("UserCreatedSecureStreamingPolicy"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
