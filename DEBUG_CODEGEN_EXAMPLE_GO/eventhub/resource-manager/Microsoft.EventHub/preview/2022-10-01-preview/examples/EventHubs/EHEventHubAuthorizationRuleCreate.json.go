package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewEventHubAuthorizationRule(ctx, "eventHubAuthorizationRule", &eventhub.EventHubAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("sdk-Authrules-2513"),
			EventHubName:          pulumi.String("sdk-EventHub-532"),
			NamespaceName:         pulumi.String("sdk-Namespace-960"),
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
