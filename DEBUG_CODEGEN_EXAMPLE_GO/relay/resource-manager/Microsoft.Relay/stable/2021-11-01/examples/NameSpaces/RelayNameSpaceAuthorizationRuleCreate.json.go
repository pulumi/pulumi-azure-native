package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/relay/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := relay.NewNamespaceAuthorizationRule(ctx, "namespaceAuthorizationRule", &relay.NamespaceAuthorizationRuleArgs{
			AuthorizationRuleName: pulumi.String("example-RelayAuthRules-01"),
			NamespaceName:         pulumi.String("example-RelayNamespace-01"),
			ResourceGroupName:     pulumi.String("resourcegroup"),
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
