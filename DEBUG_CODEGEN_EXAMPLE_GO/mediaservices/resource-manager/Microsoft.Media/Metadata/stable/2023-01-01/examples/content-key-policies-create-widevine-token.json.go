package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewContentKeyPolicy(ctx, "contentKeyPolicy", &media.ContentKeyPolicyArgs{
			AccountName:          pulumi.String("contosomedia"),
			ContentKeyPolicyName: pulumi.String("PolicyWithWidevineOptionAndJwtTokenRestriction"),
			Description:          pulumi.String("ArmPolicyDescription"),
			Options: []media.ContentKeyPolicyOptionArgs{
				{
					Configuration: {
						OdataType:        "#Microsoft.Media.ContentKeyPolicyWidevineConfiguration",
						WidevineTemplate: "{\"allowed_track_types\":\"SD_HD\",\"content_key_specs\":[{\"track_type\":\"SD\",\"security_level\":1,\"required_output_protection\":{\"hdcp\":\"HDCP_V2\"}}],\"policy_overrides\":{\"can_play\":true,\"can_persist\":true,\"can_renew\":false}}",
					},
					Name: pulumi.String("widevineoption"),
					Restriction: {
						AlternateVerificationKeys: []interface{}{
							{
								KeyValue:  "AAAAAAAAAAAAAAAAAAAAAA==",
								OdataType: "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey",
							},
						},
						Audience:  "urn:audience",
						Issuer:    "urn:issuer",
						OdataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
						PrimaryVerificationKey: {
							Exponent:  "AQAB",
							Modulus:   "AQAD",
							OdataType: "#Microsoft.Media.ContentKeyPolicyRsaTokenKey",
						},
						RestrictionTokenType: "Jwt",
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
