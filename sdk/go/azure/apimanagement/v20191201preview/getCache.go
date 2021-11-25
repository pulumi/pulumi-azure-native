


package v20191201preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:apimanagement/v20191201preview:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCacheArgs struct {
	CacheId           string `pulumi:"cacheId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupCacheResult struct {
	ConnectionString string  `pulumi:"connectionString"`
	Description      *string `pulumi:"description"`
	Id               string  `pulumi:"id"`
	Name             string  `pulumi:"name"`
	ResourceId       *string `pulumi:"resourceId"`
	Type             string  `pulumi:"type"`
}
