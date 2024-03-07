package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/media/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := media.NewContentKeyPolicy(ctx, "contentKeyPolicy", &media.ContentKeyPolicyArgs{
			AccountName:          pulumi.String("contosomedia"),
			ContentKeyPolicyName: pulumi.String("PolicyWithClearKeyOptionAndSwtTokenRestriction"),
			Description:          pulumi.String("ArmPolicyDescription"),
			Options: media.ContentKeyPolicyOptionArray{
				&media.ContentKeyPolicyOptionArgs{
					Configuration: media.ContentKeyPolicyClearKeyConfiguration{
						OdataType: "#Microsoft.Media.ContentKeyPolicyClearKeyConfiguration",
					},
					Name: pulumi.String("ClearKeyOption"),
					Restriction: media.ContentKeyPolicyTokenRestriction{
						Audience:  "urn:audience",
						Issuer:    "urn:issuer",
						OdataType: "#Microsoft.Media.ContentKeyPolicyTokenRestriction",
						PrimaryVerificationKey: media.ContentKeyPolicySymmetricTokenKey{
							KeyValue:  "AAAAAAAAAAAAAAAAAAAAAA==",
							OdataType: "#Microsoft.Media.ContentKeyPolicySymmetricTokenKey",
						},
						RestrictionTokenType: "Swt",
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
