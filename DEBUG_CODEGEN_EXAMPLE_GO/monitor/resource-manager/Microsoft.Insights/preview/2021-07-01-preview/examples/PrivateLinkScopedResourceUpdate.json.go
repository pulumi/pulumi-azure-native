package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/insights/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := insights.NewPrivateLinkScopedResource(ctx, "privateLinkScopedResource", &insights.PrivateLinkScopedResourceArgs{
			LinkedResourceId:  pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/MyResourceGroup/providers/Microsoft.Insights/components/my-component"),
			Name:              pulumi.String("scoped-resource-name"),
			ResourceGroupName: pulumi.String("MyResourceGroup"),
			ScopeName:         pulumi.String("MyPrivateLinkScope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
