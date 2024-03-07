package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewQueueAuthorizationRule(ctx, "queueAuthorizationRule", &servicebus.QueueAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("sdk-AuthRules-5800"),
			NamespaceName:         pulumi.String("sdk-Namespace-7982"),
			QueueName:             pulumi.String("sdk-Queues-2317"),
			ResourceGroupName:     pulumi.String("ArunMonocle"),
			Rights: pulumi.StringArray{
				pulumi.String("Listen"),
				pulumi.String("Send"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
