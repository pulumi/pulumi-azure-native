package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/relay/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := relay.NewNamespace(ctx, "namespace", &relay.NamespaceArgs{
			Location:          pulumi.String("South Central US"),
			NamespaceName:     pulumi.String("example-RelayNamespace-5849"),
			ResourceGroupName: pulumi.String("resourcegroup"),
			Sku: &relay.SkuArgs{
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
