package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewContentKeyPolicy(ctx, "contentKeyPolicy", &media.ContentKeyPolicyArgs{
			AccountName:          pulumi.String("contosomedia"),
			ContentKeyPolicyName: pulumi.String("PolicyWithPlayReadyOptionAndOpenRestriction"),
			Description:          pulumi.String("ArmPolicyDescription"),
			Options: media.ContentKeyPolicyOptionArray{
				&media.ContentKeyPolicyOptionArgs{
					Configuration: media.ContentKeyPolicyPlayReadyConfiguration{
						Licenses: []media.ContentKeyPolicyPlayReadyLicense{
							{
								AllowTestDevices: true,
								BeginDate:        "2017-10-16T18:22:53.46Z",
								ContentKeyLocation: {
									OdataType: "#Microsoft.Media.ContentKeyPolicyPlayReadyContentEncryptionKeyFromHeader",
								},
								ContentType: "UltraVioletDownload",
								LicenseType: "Persistent",
								PlayRight: {
									AllowPassingVideoContentToUnknownOutput:            "NotAllowed",
									DigitalVideoOnlyContentRestriction:                 false,
									ImageConstraintForAnalogComponentVideoRestriction:  true,
									ImageConstraintForAnalogComputerMonitorRestriction: false,
									ScmsRestriction: 2,
								},
								SecurityLevel: "SL150",
							},
						},
						OdataType: "#Microsoft.Media.ContentKeyPolicyPlayReadyConfiguration",
					},
					Name: pulumi.String("ArmPolicyOptionName"),
					Restriction: media.ContentKeyPolicyOpenRestriction{
						OdataType: "#Microsoft.Media.ContentKeyPolicyOpenRestriction",
					},
				},
			},
			ResourceGroupName: pulumi.String("contosorg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
