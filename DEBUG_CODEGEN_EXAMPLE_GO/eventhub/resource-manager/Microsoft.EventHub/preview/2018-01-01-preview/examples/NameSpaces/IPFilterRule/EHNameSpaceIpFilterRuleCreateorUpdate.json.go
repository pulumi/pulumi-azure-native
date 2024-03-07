package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewNamespaceIpFilterRule(ctx, "namespaceIpFilterRule", &eventhub.NamespaceIpFilterRuleArgs{
			Action:            pulumi.String("Accept"),
			FilterName:        pulumi.String("sdk-IPFilterRules-7337"),
			IpFilterRuleName:  pulumi.String("sdk-IPFilterRules-7337"),
			IpMask:            pulumi.String("13.78.143.246/32"),
			NamespaceName:     pulumi.String("sdk-Namespace-5232"),
			ResourceGroupName: pulumi.String("ResourceGroup"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
