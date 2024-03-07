package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewLinkedServer(ctx, "linkedServer", &cache.LinkedServerArgs{
			LinkedRedisCacheId:       pulumi.String("/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/Redis/cache2"),
			LinkedRedisCacheLocation: pulumi.String("West US"),
			LinkedServerName:         pulumi.String("cache2"),
			Name:                     pulumi.String("cache1"),
			ResourceGroupName:        pulumi.String("rg1"),
			ServerRole:               cache.ReplicationRoleSecondary,
		})
		if err != nil {
			return err
		}
		return nil
	})
}
