package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/servicebus/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := servicebus.NewNamespace(ctx, "namespace", &servicebus.NamespaceArgs{
			Location:          pulumi.String("South Central US"),
			NamespaceName:     pulumi.String("sdk-Namespace2924"),
			ResourceGroupName: pulumi.String("ArunMonocle"),
			Sku: &servicebus.SBSkuArgs{
				Name: pulumi.String("Standard"),
				Tier: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("value1"),
				"tag2": pulumi.String("value2"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
