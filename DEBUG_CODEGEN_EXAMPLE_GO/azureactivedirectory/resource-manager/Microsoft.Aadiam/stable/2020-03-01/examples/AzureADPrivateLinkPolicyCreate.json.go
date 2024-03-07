package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/aadiam/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := aadiam.NewPrivateLinkForAzureAd(ctx, "privateLinkForAzureAd", &aadiam.PrivateLinkForAzureAdArgs{
			AllTenants:        pulumi.Bool(false),
			Name:              pulumi.String("myOrgPrivateLinkPolicy"),
			OwnerTenantId:     pulumi.String("950f8bca-bf4d-4a41-ad10-034e792a243d"),
			PolicyName:        pulumi.String("ddb1"),
			ResourceGroup:     pulumi.String("myOrgVnetRG"),
			ResourceGroupName: pulumi.String("rg1"),
			ResourceName:      pulumi.String("myOrgVnetPrivateLink"),
			SubscriptionId:    pulumi.String("57849194-ea1f-470b-abda-d195b25634c1"),
			Tenants: pulumi.StringArray{
				pulumi.String("3616657d-1c80-41ae-9d83-2a2776f2c9be"),
				pulumi.String("727b6ef1-18ab-4627-ac95-3f9cd945ed87"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
