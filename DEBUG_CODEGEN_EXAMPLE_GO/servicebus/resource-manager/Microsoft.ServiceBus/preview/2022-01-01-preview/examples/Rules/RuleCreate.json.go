package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewRule(ctx, "rule", &servicebus.RuleArgs{
			NamespaceName:     pulumi.String("sdk-Namespace-1319"),
			ResourceGroupName: pulumi.String("resourceGroupName"),
			RuleName:          pulumi.String("sdk-Rules-6571"),
			SubscriptionName:  pulumi.String("sdk-Subscriptions-8691"),
			TopicName:         pulumi.String("sdk-Topics-2081"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
