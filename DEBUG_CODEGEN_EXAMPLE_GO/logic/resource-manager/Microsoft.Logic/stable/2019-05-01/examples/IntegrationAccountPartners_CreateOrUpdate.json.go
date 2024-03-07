package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccountPartner(ctx, "integrationAccountPartner", &logic.IntegrationAccountPartnerArgs{
			Content: &logic.PartnerContentArgs{
				B2b: &logic.B2BPartnerContentArgs{
					BusinessIdentities: logic.BusinessIdentityArray{
						&logic.BusinessIdentityArgs{
							Qualifier: pulumi.String("AA"),
							Value:     pulumi.String("ZZ"),
						},
					},
				},
			},
			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
			Location:               pulumi.String("westus"),
			Metadata:               nil,
			PartnerName:            pulumi.String("testPartner"),
			PartnerType:            pulumi.String("B2B"),
			ResourceGroupName:      pulumi.String("testResourceGroup"),
			Tags:                   nil,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
