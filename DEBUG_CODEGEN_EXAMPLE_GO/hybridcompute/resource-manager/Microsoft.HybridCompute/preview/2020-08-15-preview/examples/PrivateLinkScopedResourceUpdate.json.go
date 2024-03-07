package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewPrivateLinkScopedResource(ctx, "privateLinkScopedResource", &hybridcompute.PrivateLinkScopedResourceArgs{
			LinkedResourceId:  pulumi.String("/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/myResourceGroup/providers/Microsoft.HybridCompute/Machines/machineName1"),
			Name:              pulumi.String("scoped-resource-name"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ScopeName:         pulumi.String("myPrivateLinkScope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
