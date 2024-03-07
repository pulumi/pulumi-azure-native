package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/hybridcompute/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := hybridcompute.NewPrivateLinkScope(ctx, "privateLinkScope", &hybridcompute.PrivateLinkScopeArgs{
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ScopeName:         pulumi.String("my-privatelinkscope"),
			Tags: pulumi.StringMap{
				"Tag1": pulumi.String("Value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
