package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/eventhub/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := eventhub.NewNamespaceVirtualNetworkRule(ctx, "namespaceVirtualNetworkRule", &eventhub.NamespaceVirtualNetworkRuleArgs{
			NamespaceName:          pulumi.String("sdk-Namespace-6019"),
			ResourceGroupName:      pulumi.String("ResourceGroup"),
			VirtualNetworkRuleName: pulumi.String("sdk-VirtualNetworkRules-9191"),
			VirtualNetworkSubnetId: pulumi.String("/subscriptions/Subscription/resourceGroups/sbehvnettest/providers/Microsoft.Network/virtualNetworks/sbehvnettest/subnets/default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
