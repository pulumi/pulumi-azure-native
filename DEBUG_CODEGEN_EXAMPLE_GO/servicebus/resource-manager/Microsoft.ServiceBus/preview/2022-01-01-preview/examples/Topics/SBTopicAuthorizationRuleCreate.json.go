package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewTopicAuthorizationRule(ctx, "topicAuthorizationRule", &servicebus.TopicAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("sdk-AuthRules-4310"),
			NamespaceName:         pulumi.String("sdk-Namespace-6261"),
			ResourceGroupName:     pulumi.String("ArunMonocle"),
			Rights: pulumi.StringArray{
				pulumi.String("Listen"),
				pulumi.String("Send"),
			},
			TopicName: pulumi.String("sdk-Topics-1984"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
