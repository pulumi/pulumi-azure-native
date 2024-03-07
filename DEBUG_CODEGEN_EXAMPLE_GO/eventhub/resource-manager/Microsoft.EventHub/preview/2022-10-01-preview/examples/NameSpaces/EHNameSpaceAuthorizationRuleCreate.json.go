package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewNamespaceAuthorizationRule(ctx, "namespaceAuthorizationRule", &eventhub.NamespaceAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("sdk-Authrules-1746"),
			NamespaceName:         pulumi.String("sdk-Namespace-2702"),
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
