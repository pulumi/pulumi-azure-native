


package v20170601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRegistry(ctx *pulumi.Context, args *LookupRegistryArgs, opts ...pulumi.InvokeOption) (*LookupRegistryResult, error) {
	var rv LookupRegistryResult
	err := ctx.Invoke("azure-native:containerregistry/v20170601preview:getRegistry", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupRegistryArgs struct {
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupRegistryResult struct {
	AdminUserEnabled  *bool                             `pulumi:"adminUserEnabled"`
	CreationDate      string                            `pulumi:"creationDate"`
	Id                string                            `pulumi:"id"`
	Location          string                            `pulumi:"location"`
	LoginServer       string                            `pulumi:"loginServer"`
	Name              string                            `pulumi:"name"`
	ProvisioningState string                            `pulumi:"provisioningState"`
	Sku               SkuResponse                       `pulumi:"sku"`
	Status            StatusResponse                    `pulumi:"status"`
	StorageAccount    *StorageAccountPropertiesResponse `pulumi:"storageAccount"`
	Tags              map[string]string                 `pulumi:"tags"`
	Type              string                            `pulumi:"type"`
}


func (val *LookupRegistryResult) Defaults() *LookupRegistryResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AdminUserEnabled) {
		adminUserEnabled_ := false
		tmp.AdminUserEnabled = &adminUserEnabled_
	}
	return &tmp
}
