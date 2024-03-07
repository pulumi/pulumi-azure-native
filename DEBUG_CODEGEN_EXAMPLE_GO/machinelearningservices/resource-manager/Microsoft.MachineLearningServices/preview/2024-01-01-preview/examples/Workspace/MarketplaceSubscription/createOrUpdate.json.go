package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewMarketplaceSubscription(ctx, "marketplaceSubscription", &machinelearningservices.MarketplaceSubscriptionArgs{
			MarketplaceSubscriptionProperties: &machinelearningservices.MarketplaceSubscriptionTypeArgs{
				ModelId: pulumi.String("string"),
			},
			Name:              pulumi.String("string"),
			ResourceGroupName: pulumi.String("test-rg"),
			WorkspaceName:     pulumi.String("my-aml-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
