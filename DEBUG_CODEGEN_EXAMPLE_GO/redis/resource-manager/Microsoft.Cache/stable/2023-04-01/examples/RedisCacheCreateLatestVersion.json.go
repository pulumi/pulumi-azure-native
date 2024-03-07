package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewRedis(ctx, "redis", &cache.RedisArgs{
			EnableNonSslPort:  pulumi.Bool(true),
			Location:          pulumi.String("West US"),
			MinimumTlsVersion: pulumi.String("1.2"),
			Name:              pulumi.String("cache1"),
			RedisConfiguration: &cache.RedisCommonPropertiesRedisConfigurationArgs{
				MaxmemoryPolicy: pulumi.String("allkeys-lru"),
			},
			RedisVersion:       pulumi.String("Latest"),
			ReplicasPerPrimary: pulumi.Int(2),
			ResourceGroupName:  pulumi.String("rg1"),
			ShardCount:         pulumi.Int(2),
			Sku: &cache.SkuArgs{
				Capacity: pulumi.Int(1),
				Family:   pulumi.String("P"),
				Name:     pulumi.String("Premium"),
			},
			StaticIP: pulumi.String("192.168.0.5"),
			SubnetId: pulumi.String("/subscriptions/subid/resourceGroups/rg2/providers/Microsoft.Network/virtualNetworks/network1/subnets/subnet1"),
			Zones: pulumi.StringArray{
				pulumi.String("1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
