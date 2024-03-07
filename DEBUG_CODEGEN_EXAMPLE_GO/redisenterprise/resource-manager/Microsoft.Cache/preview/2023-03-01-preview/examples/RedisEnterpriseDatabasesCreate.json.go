package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/cache/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := cache.NewDatabase(ctx, "database", &cache.DatabaseArgs{
			ClientProtocol:   pulumi.String("Encrypted"),
			ClusterName:      pulumi.String("cache1"),
			ClusteringPolicy: pulumi.String("EnterpriseCluster"),
			DatabaseName:     pulumi.String("default"),
			EvictionPolicy:   pulumi.String("AllKeysLRU"),
			Modules: cache.ModuleArray{
				&cache.ModuleArgs{
					Args: pulumi.String("ERROR_RATE 0.00 INITIAL_SIZE 400"),
					Name: pulumi.String("RedisBloom"),
				},
				&cache.ModuleArgs{
					Args: pulumi.String("RETENTION_POLICY 20"),
					Name: pulumi.String("RedisTimeSeries"),
				},
				&cache.ModuleArgs{
					Name: pulumi.String("RediSearch"),
				},
			},
			Persistence: &cache.PersistenceArgs{
				AofEnabled:   pulumi.Bool(true),
				AofFrequency: pulumi.String("1s"),
			},
			Port:              pulumi.Int(10000),
			ResourceGroupName: pulumi.String("rg1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
