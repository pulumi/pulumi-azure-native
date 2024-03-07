package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/managednetworkfabric/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := managednetworkfabric.NewInternetGateway(ctx, "internetGateway", &managednetworkfabric.InternetGatewayArgs{
			Annotation:                pulumi.String("annotation"),
			InternetGatewayName:       pulumi.String("example-internetGateway"),
			InternetGatewayRuleId:     pulumi.String("/subscriptions/xxxx-xxxx-xxxx-xxxx/providers/Microsoft.ManagedNetworkFabric/internetGatewayRules/example-internetGatewayRule"),
			Location:                  pulumi.String("eastus"),
			NetworkFabricControllerId: pulumi.String("/subscriptions/xxxx-xxxx-xxxx-xxxx/resourcegroups/example-rg/providers/Microsoft.ManagedNetworkFabric/networkFabricControllers/example-networkFabricController"),
			ResourceGroupName:         pulumi.String("example-rg"),
			Tags: pulumi.StringMap{
				"key3540": pulumi.String("1234"),
			},
			Type: pulumi.String("Infrastructure"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
