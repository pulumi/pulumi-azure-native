


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20190701:getVolume", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupVolumeArgs struct {
	AccountName       string `pulumi:"accountName"`
	PoolName          string `pulumi:"poolName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VolumeName        string `pulumi:"volumeName"`
}


type LookupVolumeResult struct {
	BaremetalTenantId string                                `pulumi:"baremetalTenantId"`
	CreationToken     string                                `pulumi:"creationToken"`
	ExportPolicy      *VolumePropertiesResponseExportPolicy `pulumi:"exportPolicy"`
	FileSystemId      string                                `pulumi:"fileSystemId"`
	Id                string                                `pulumi:"id"`
	Location          string                                `pulumi:"location"`
	MountTargets      []MountTargetPropertiesResponse       `pulumi:"mountTargets"`
	Name              string                                `pulumi:"name"`
	ProtocolTypes     []string                              `pulumi:"protocolTypes"`
	ProvisioningState string                                `pulumi:"provisioningState"`
	ServiceLevel      *string                               `pulumi:"serviceLevel"`
	SnapshotId        *string                               `pulumi:"snapshotId"`
	SubnetId          string                                `pulumi:"subnetId"`
	Tags              interface{}                           `pulumi:"tags"`
	Type              string                                `pulumi:"type"`
	UsageThreshold    float64                               `pulumi:"usageThreshold"`
}


func (val *LookupVolumeResult) Defaults() *LookupVolumeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ServiceLevel) {
		serviceLevel_ := "Premium"
		tmp.ServiceLevel = &serviceLevel_
	}
	if isZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
	}
	return &tmp
}
