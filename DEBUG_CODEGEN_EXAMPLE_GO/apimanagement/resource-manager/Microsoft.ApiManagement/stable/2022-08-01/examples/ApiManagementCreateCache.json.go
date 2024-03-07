package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewCache(ctx, "cache", &apimanagement.CacheArgs{
			CacheId:           pulumi.String("c1"),
			ConnectionString:  pulumi.String("apim.redis.cache.windows.net:6380,password=xc,ssl=True,abortConnect=False"),
			Description:       pulumi.String("Redis cache instances in West India"),
			ResourceGroupName: pulumi.String("rg1"),
			ResourceId:        pulumi.String("https://management.azure.com/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Cache/redis/apimservice1"),
			ServiceName:       pulumi.String("apimService1"),
			UseFromLocation:   pulumi.String("default"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
