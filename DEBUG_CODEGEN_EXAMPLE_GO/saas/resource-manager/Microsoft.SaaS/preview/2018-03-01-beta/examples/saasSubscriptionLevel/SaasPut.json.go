package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/saas/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := saas.NewSaasSubscriptionLevel(ctx, "saasSubscriptionLevel", &saas.SaasSubscriptionLevelArgs{
			Location: pulumi.String("global"),
			Name:     pulumi.String("MyContosoSubscription"),
			Properties: &saas.SaasCreationPropertiesArgs{
				OfferId: pulumi.String("contosoOffer"),
				PaymentChannelMetadata: pulumi.StringMap{
					"AzureSubscriptionId": pulumi.String("155af98a-3205-47e7-883b-a2ab9db9f88d"),
				},
				PaymentChannelType: pulumi.String("SubscriptionDelegated"),
				PublisherId:        pulumi.String("microsoft-contoso"),
				SaasResourceName:   pulumi.String("MyContosoSubscription"),
				SkuId:              pulumi.String("free"),
				TermId:             pulumi.String("hjdtn7tfnxcy"),
			},
			ResourceGroupName: pulumi.String("my-saas-rg"),
			ResourceName:      pulumi.String("MyContosoSubscription"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
