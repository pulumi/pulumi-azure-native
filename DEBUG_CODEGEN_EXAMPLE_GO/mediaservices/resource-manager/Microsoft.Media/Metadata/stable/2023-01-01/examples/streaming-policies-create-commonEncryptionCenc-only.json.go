package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewStreamingPolicy(ctx, "streamingPolicy", &media.StreamingPolicyArgs{
			AccountName: pulumi.String("contosomedia"),
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
			DefaultContentKeyPolicyName: pulumi.String("PolicyWithPlayReadyOptionAndOpenRestriction"),
			ResourceGroupName:           pulumi.String("contosorg"),
			StreamingPolicyName:         pulumi.String("UserCreatedSecureStreamingPolicyWithCommonEncryptionCencOnly"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
