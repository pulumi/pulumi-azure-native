package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/features/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := features.NewSubscriptionFeatureRegistration(ctx, "subscriptionFeatureRegistration", &features.SubscriptionFeatureRegistrationArgs{
			FeatureName:       pulumi.String("testFeature"),
			Properties:        nil,
			ProviderNamespace: pulumi.String("subscriptionFeatureRegistrationGroupTestRG"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
