


package v20180401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:compute/v20180401:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupDiskArgs struct {
	DiskName          string `pulumi:"diskName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskResult struct {
	CreationData       CreationDataResponse        `pulumi:"creationData"`
	DiskSizeGB         *int                        `pulumi:"diskSizeGB"`
	EncryptionSettings *EncryptionSettingsResponse `pulumi:"encryptionSettings"`
	Id                 string                      `pulumi:"id"`
	Location           string                      `pulumi:"location"`
	ManagedBy          string                      `pulumi:"managedBy"`
	Name               string                      `pulumi:"name"`
	OsType             *string                     `pulumi:"osType"`
	ProvisioningState  string                      `pulumi:"provisioningState"`
	Sku                *DiskSkuResponse            `pulumi:"sku"`
	Tags               map[string]string           `pulumi:"tags"`
	TimeCreated        string                      `pulumi:"timeCreated"`
	Type               string                      `pulumi:"type"`
	Zones              []string                    `pulumi:"zones"`
}


func (val *LookupDiskResult) Defaults() *LookupDiskResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
