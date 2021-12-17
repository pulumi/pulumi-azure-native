


package v20210301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCache(ctx *pulumi.Context, args *LookupCacheArgs, opts ...pulumi.InvokeOption) (*LookupCacheResult, error) {
	var rv LookupCacheResult
	err := ctx.Invoke("azure-native:storagecache/v20210301:getCache", args, &rv, opts...)
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
	CacheSizeGB               *int                             `pulumi:"cacheSizeGB"`
	DirectoryServicesSettings *CacheDirectorySettingsResponse  `pulumi:"directoryServicesSettings"`
	EncryptionSettings        *CacheEncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Health                    CacheHealthResponse              `pulumi:"health"`
	Id                        string                           `pulumi:"id"`
	Identity                  *CacheIdentityResponse           `pulumi:"identity"`
	Location                  *string                          `pulumi:"location"`
	MountAddresses            []string                         `pulumi:"mountAddresses"`
	Name                      string                           `pulumi:"name"`
	NetworkSettings           *CacheNetworkSettingsResponse    `pulumi:"networkSettings"`
	ProvisioningState         *string                          `pulumi:"provisioningState"`
	SecuritySettings          *CacheSecuritySettingsResponse   `pulumi:"securitySettings"`
	Sku                       *CacheResponseSku                `pulumi:"sku"`
	Subnet                    *string                          `pulumi:"subnet"`
	SystemData                SystemDataResponse               `pulumi:"systemData"`
	Tags                      map[string]string                `pulumi:"tags"`
	Type                      string                           `pulumi:"type"`
	UpgradeStatus             *CacheUpgradeStatusResponse      `pulumi:"upgradeStatus"`
}


func (val *LookupCacheResult) Defaults() *LookupCacheResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.DirectoryServicesSettings = tmp.DirectoryServicesSettings.Defaults()

	tmp.NetworkSettings = tmp.NetworkSettings.Defaults()

	return &tmp
}
