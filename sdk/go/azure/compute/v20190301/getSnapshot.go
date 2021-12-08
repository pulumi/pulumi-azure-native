


package v20190301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSnapshot(ctx *pulumi.Context, args *LookupSnapshotArgs, opts ...pulumi.InvokeOption) (*LookupSnapshotResult, error) {
	var rv LookupSnapshotResult
	err := ctx.Invoke("azure-native:compute/v20190301:getSnapshot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSnapshotArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SnapshotName      string `pulumi:"snapshotName"`
}


type LookupSnapshotResult struct {
	CreationData                 CreationDataResponse                  `pulumi:"creationData"`
	DiskSizeBytes                float64                               `pulumi:"diskSizeBytes"`
	DiskSizeGB                   *int                                  `pulumi:"diskSizeGB"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                               `pulumi:"hyperVGeneration"`
	Id                           string                                `pulumi:"id"`
	Incremental                  *bool                                 `pulumi:"incremental"`
	Location                     string                                `pulumi:"location"`
	ManagedBy                    string                                `pulumi:"managedBy"`
	Name                         string                                `pulumi:"name"`
	OsType                       *string                               `pulumi:"osType"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	Sku                          *SnapshotSkuResponse                  `pulumi:"sku"`
	Tags                         map[string]string                     `pulumi:"tags"`
	TimeCreated                  string                                `pulumi:"timeCreated"`
	Type                         string                                `pulumi:"type"`
	UniqueId                     string                                `pulumi:"uniqueId"`
}
