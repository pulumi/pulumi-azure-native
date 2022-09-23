


package v20200301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:storagecache/v20200301:getCache", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupCacheArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupCacheResult struct {
	CacheSizeGB        *int                             `pulumi:"cacheSizeGB"`
	EncryptionSettings *CacheEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Health             CacheHealthResponse              `pulumi:"health"`
	Id                 string                           `pulumi:"id"`
	Identity           *CacheIdentityResponse           `pulumi:"identity"`
	Location           *string                          `pulumi:"location"`
	MountAddresses     []string                         `pulumi:"mountAddresses"`
	Name               string                           `pulumi:"name"`
	NetworkSettings    *CacheNetworkSettingsResponse    `pulumi:"networkSettings"`
	ProvisioningState  *string                          `pulumi:"provisioningState"`
	SecuritySettings   *CacheSecuritySettingsResponse   `pulumi:"securitySettings"`
	Sku                *CacheResponseSku                `pulumi:"sku"`
	Subnet             *string                          `pulumi:"subnet"`
	SystemData         SystemDataResponse               `pulumi:"systemData"`
	Tags               interface{}                      `pulumi:"tags"`
	Type               string                           `pulumi:"type"`
	UpgradeStatus      *CacheUpgradeStatusResponse      `pulumi:"upgradeStatus"`
}


func (val *LookupCacheResult) Defaults() *LookupCacheResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.NetworkSettings = tmp.NetworkSettings.Defaults()

	return &tmp
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

func (o LookupCacheResultOutput) EncryptionSettings() CacheEncryptionSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheEncryptionSettingsResponse { return v.EncryptionSettings }).(CacheEncryptionSettingsResponsePtrOutput)
}

func (o LookupCacheResultOutput) Health() CacheHealthResponseOutput {
	return o.ApplyT(func(v LookupCacheResult) CacheHealthResponse { return v.Health }).(CacheHealthResponseOutput)
}

func (o LookupCacheResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCacheResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCacheResultOutput) Identity() CacheIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheIdentityResponse { return v.Identity }).(CacheIdentityResponsePtrOutput)
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

func (o LookupCacheResultOutput) NetworkSettings() CacheNetworkSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheNetworkSettingsResponse { return v.NetworkSettings }).(CacheNetworkSettingsResponsePtrOutput)
}

func (o LookupCacheResultOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) SecuritySettings() CacheSecuritySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheSecuritySettingsResponse { return v.SecuritySettings }).(CacheSecuritySettingsResponsePtrOutput)
}

func (o LookupCacheResultOutput) Sku() CacheResponseSkuPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *CacheResponseSku { return v.Sku }).(CacheResponseSkuPtrOutput)
}

func (o LookupCacheResultOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupCacheResult) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o LookupCacheResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupCacheResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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
