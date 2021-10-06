


package v20200901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20200901:getVolume", args, &rv, opts...)
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
	BackupId                 *string                                 `pulumi:"backupId"`
	BaremetalTenantId        string                                  `pulumi:"baremetalTenantId"`
	CreationToken            string                                  `pulumi:"creationToken"`
	DataProtection           *VolumePropertiesResponseDataProtection `pulumi:"dataProtection"`
	ExportPolicy             *VolumePropertiesResponseExportPolicy   `pulumi:"exportPolicy"`
	FileSystemId             string                                  `pulumi:"fileSystemId"`
	Id                       string                                  `pulumi:"id"`
	IsRestoring              *bool                                   `pulumi:"isRestoring"`
	KerberosEnabled          *bool                                   `pulumi:"kerberosEnabled"`
	Location                 string                                  `pulumi:"location"`
	MountTargets             []MountTargetPropertiesResponse         `pulumi:"mountTargets"`
	Name                     string                                  `pulumi:"name"`
	ProtocolTypes            []string                                `pulumi:"protocolTypes"`
	ProvisioningState        string                                  `pulumi:"provisioningState"`
	SecurityStyle            *string                                 `pulumi:"securityStyle"`
	ServiceLevel             *string                                 `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable *bool                                   `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption            *bool                                   `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible *bool                                   `pulumi:"snapshotDirectoryVisible"`
	SnapshotId               *string                                 `pulumi:"snapshotId"`
	SubnetId                 string                                  `pulumi:"subnetId"`
	Tags                     map[string]string                       `pulumi:"tags"`
	ThroughputMibps          *float64                                `pulumi:"throughputMibps"`
	Type                     string                                  `pulumi:"type"`
	UsageThreshold           float64                                 `pulumi:"usageThreshold"`
	VolumeType               *string                                 `pulumi:"volumeType"`
}
