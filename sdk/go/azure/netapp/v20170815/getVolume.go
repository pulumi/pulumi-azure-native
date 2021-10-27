


package v20170815

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20170815:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVolumeArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupVolumeResult struct {
	CreationToken     string                                `pulumi:"creationToken"`
	ExportPolicy      *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	FileSystemId      string                                `pulumi:"fileSystemId"`
	Id                string                                `pulumi:"id"`
	Location          string                                `pulumi:"location"`
	Name              string                                `pulumi:"name"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	ServiceLevel      string                                `pulumi:"serviceLevel"`
	SubnetId          *string                               `pulumi:"subnetId"`
	Tags              interface{}                           `pulumi:"tags"`
	Type              string                                `pulumi:"type"`
	UsageThreshold    *float64                              `pulumi:"usageThreshold"`
}
