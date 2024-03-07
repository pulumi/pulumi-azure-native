package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewNamespaceAuthorizationRule(ctx, "namespaceAuthorizationRule", &servicebus.NamespaceAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("sdk-AuthRules-1788"),
			NamespaceName:         pulumi.String("sdk-Namespace-6914"),
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
