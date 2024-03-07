package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewSubscription(ctx, "subscription", &servicebus.SubscriptionArgs{
			EnableBatchedOperations: pulumi.Bool(true),
			NamespaceName:           pulumi.String("sdk-Namespace-1349"),
			ResourceGroupName:       pulumi.String("ResourceGroup"),
			SubscriptionName:        pulumi.String("sdk-Subscriptions-2178"),
			TopicName:               pulumi.String("sdk-Topics-8740"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
