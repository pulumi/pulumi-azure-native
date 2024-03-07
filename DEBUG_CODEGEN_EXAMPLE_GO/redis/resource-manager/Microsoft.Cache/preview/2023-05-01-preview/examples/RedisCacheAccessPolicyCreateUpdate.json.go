package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewAccessPolicy(ctx, "accessPolicy", &cache.AccessPolicyArgs{
			AccessPolicyName:  pulumi.String("accessPolicy1"),
			CacheName:         pulumi.String("cache1"),
			Permissions:       pulumi.String("+get +hget"),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
