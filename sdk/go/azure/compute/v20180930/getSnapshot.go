


package v20180930

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:compute/v20180930:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SnapshotName      string `pulumi:"snapshotName"`
}


type LookupSnapshotResult struct {
	CreationData                 CreationDataResponse                  `pulumi:"creationData"`
	DiskSizeGB                   *int                                  `pulumi:"diskSizeGB"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                               `pulumi:"hyperVGeneration"`
	Id                           string                                `pulumi:"id"`
	Location                     string                                `pulumi:"location"`
	ManagedBy                    string                                `pulumi:"managedBy"`
	Name                         string                                `pulumi:"name"`
	OsType                       *string                               `pulumi:"osType"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	Sku                          *SnapshotSkuResponse                  `pulumi:"sku"`
	Tags                         map[string]string                     `pulumi:"tags"`
	TimeCreated                  string                                `pulumi:"timeCreated"`
	Type                         string                                `pulumi:"type"`
}


func (val *LookupSnapshotResult) Defaults() *LookupSnapshotResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Sku = tmp.Sku.Defaults()

	return &tmp
}
