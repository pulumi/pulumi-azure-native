


package v20190801preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:storagecache/v20190801preview:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCacheResult struct {
	CacheSizeGB       *int                        `pulumi:"cacheSizeGB"`
	Health            CacheHealthResponse         `pulumi:"health"`
	Id                string                      `pulumi:"id"`
	Location          *string                     `pulumi:"location"`
	MountAddresses    []string                    `pulumi:"mountAddresses"`
	Name              string                      `pulumi:"name"`
	ProvisioningState *string                     `pulumi:"provisioningState"`
	Sku               *CacheResponseSku           `pulumi:"sku"`
	Subnet            *string                     `pulumi:"subnet"`
	Tags              interface{}                 `pulumi:"tags"`
	Type              string                      `pulumi:"type"`
	UpgradeStatus     *CacheUpgradeStatusResponse `pulumi:"upgradeStatus"`
}
