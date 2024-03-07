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
			EvictionPolicy:   pulumi.String("NoEviction"),
			GeoReplication: cache.DatabasePropertiesResponseGeoReplication{
				GroupNickname: pulumi.String("groupName"),
				LinkedDatabases: cache.LinkedDatabaseArray{
					&cache.LinkedDatabaseArgs{
						Id: pulumi.String("/subscriptions/subid1/resourceGroups/rg1/providers/Microsoft.Cache/redisEnterprise/cache1/databases/default"),
					},
					&cache.LinkedDatabaseArgs{
						Id: pulumi.String("/subscriptions/subid2/resourceGroups/rg2/providers/Microsoft.Cache/redisEnterprise/cache2/databases/default"),
					},
				},
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
