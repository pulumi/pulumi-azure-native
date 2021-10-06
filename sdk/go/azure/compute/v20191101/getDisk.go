


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDisk(ctx *pulumi.Context, args *LookupDiskArgs, opts ...pulumi.InvokeOption) (*LookupDiskResult, error) {
	var rv LookupDiskResult
	err := ctx.Invoke("azure-native:compute/v20191101:getDisk", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiskArgs struct {
	DiskName          string `pulumi:"diskName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupDiskResult struct {
	CreationData                 CreationDataResponse                  `pulumi:"creationData"`
	DiskIOPSReadOnly             *float64                              `pulumi:"diskIOPSReadOnly"`
	DiskIOPSReadWrite            *float64                              `pulumi:"diskIOPSReadWrite"`
	DiskMBpsReadOnly             *float64                              `pulumi:"diskMBpsReadOnly"`
	DiskMBpsReadWrite            *float64                              `pulumi:"diskMBpsReadWrite"`
	DiskSizeBytes                float64                               `pulumi:"diskSizeBytes"`
	DiskSizeGB                   *int                                  `pulumi:"diskSizeGB"`
	DiskState                    string                                `pulumi:"diskState"`
	Encryption                   *EncryptionResponse                   `pulumi:"encryption"`
	EncryptionSettingsCollection *EncryptionSettingsCollectionResponse `pulumi:"encryptionSettingsCollection"`
	HyperVGeneration             *string                               `pulumi:"hyperVGeneration"`
	Id                           string                                `pulumi:"id"`
	Location                     string                                `pulumi:"location"`
	ManagedBy                    string                                `pulumi:"managedBy"`
	ManagedByExtended            []string                              `pulumi:"managedByExtended"`
	MaxShares                    *int                                  `pulumi:"maxShares"`
	Name                         string                                `pulumi:"name"`
	OsType                       *string                               `pulumi:"osType"`
	ProvisioningState            string                                `pulumi:"provisioningState"`
	ShareInfo                    []ShareInfoElementResponse            `pulumi:"shareInfo"`
	Sku                          *DiskSkuResponse                      `pulumi:"sku"`
	Tags                         map[string]string                     `pulumi:"tags"`
	TimeCreated                  string                                `pulumi:"timeCreated"`
	Type                         string                                `pulumi:"type"`
	UniqueId                     string                                `pulumi:"uniqueId"`
	Zones                        []string                              `pulumi:"zones"`
}
