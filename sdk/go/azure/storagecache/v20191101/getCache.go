


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:storagecache/v20191101:getCache", args, &rv, opts...)
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

func LookupCacheOutput(ctx *pulumi.Context, args LookupCacheOutputArgs, opts ...pulumi.InvokeOption) LookupCacheResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCacheResult, error) {
			args := v.(LookupCacheArgs)
			r, err := LookupCache(ctx, &args, opts...)
			var s LookupCacheResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCacheResultOutput)
}

type LookupCacheOutputArgs struct {
	CacheName         pulumi.StringInput `pulumi:"cacheName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupCacheOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheArgs)(nil)).Elem()
}


type LookupCacheResultOutput struct{ *pulumi.OutputState }

func (LookupCacheResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCacheResult)(nil)).Elem()
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutput() LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) ToLookupCacheResultOutputWithContext(ctx context.Context) LookupCacheResultOutput {
	return o
}

func (o LookupCacheResultOutput) CacheSizeGB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *int { return v.CacheSizeGB }).(pulumi.IntPtrOutput)
}

func (o LookupCacheResultOutput) Health() CacheHealthResponseOutput {
	return o.ApplyT(func(v LookupCacheResult) CacheHealthResponse { return v.Health }).(CacheHealthResponseOutput)
}

func (o LookupCacheResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) MountAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupCacheResult) []string { return v.MountAddresses }).(pulumi.StringArrayOutput)
}

func (o LookupCacheResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) Sku() CacheResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheResponseSku { return v.Sku }).(CacheResponseSkuPtrOutput)
}

func (o LookupCacheResultOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupCacheResult) interface{} { return v.Tags }).(pulumi.AnyOutput)
}

func (o LookupCacheResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) UpgradeStatus() CacheUpgradeStatusResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheUpgradeStatusResponse { return v.UpgradeStatus }).(CacheUpgradeStatusResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCacheResultOutput{})
}
