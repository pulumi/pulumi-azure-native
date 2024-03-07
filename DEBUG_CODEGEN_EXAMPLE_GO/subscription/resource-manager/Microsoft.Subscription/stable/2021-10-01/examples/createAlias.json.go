package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/subscription/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := subscription.NewAlias(ctx, "alias", &subscription.AliasArgs{
			AliasName: pulumi.String("aliasForNewSub"),
			Properties: subscription.SubscriptionAliasResponsePropertiesResponse{
				AdditionalProperties: &subscription.PutAliasRequestAdditionalPropertiesArgs{
					SubscriptionOwnerId:  pulumi.String("f09b39eb-c496-482c-9ab9-afd799572f4c"),
					SubscriptionTenantId: pulumi.String("66f6e4d6-07dc-4aea-94ea-e12d3026a3c8"),
					Tags: pulumi.StringMap{
						"tag1": pulumi.String("Messi"),
						"tag2": pulumi.String("Ronaldo"),
						"tag3": pulumi.String("Lebron"),
					},
				},
				BillingScope: pulumi.String("/billingAccounts/af6231a7-7f8d-4fcc-a993-dd8466108d07:c663dac6-a9a5-405a-8938-cd903e12ab5b_2019_05_31/billingProfiles/QWDQ-QWHI-AUW-SJDO-DJH/invoiceSections/FEUF-EUHE-ISJ-SKDW-DJH"),
				DisplayName:  pulumi.String("Test Subscription"),
				Workload:     pulumi.String("Production"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
