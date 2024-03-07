package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/kubernetesconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := kubernetesconfiguration.NewPrivateLinkScope(ctx, "privateLinkScope", &kubernetesconfiguration.PrivateLinkScopeArgs{
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("my-resource-group"),
			ScopeName:         pulumi.String("my-privatelinkscope"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
