


package v20210801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVolume(ctx *pulumi.Context, args *LookupVolumeArgs, opts ...pulumi.InvokeOption) (*LookupVolumeResult, error) {
	var rv LookupVolumeResult
	err := ctx.Invoke("azure-native:netapp/v20210801:getVolume", args, &rv, opts...)
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
	AvsDataStore              *string                                 `pulumi:"avsDataStore"`
	BackupId                  *string                                 `pulumi:"backupId"`
	BaremetalTenantId         string                                  `pulumi:"baremetalTenantId"`
	CapacityPoolResourceId    *string                                 `pulumi:"capacityPoolResourceId"`
	CloneProgress             int                                     `pulumi:"cloneProgress"`
	CoolAccess                *bool                                   `pulumi:"coolAccess"`
	CoolnessPeriod            *int                                    `pulumi:"coolnessPeriod"`
	CreationToken             string                                  `pulumi:"creationToken"`
	DataProtection            *VolumePropertiesResponseDataProtection `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs   *float64                                `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs    *float64                                `pulumi:"defaultUserQuotaInKiBs"`
	EncryptionKeySource       *string                                 `pulumi:"encryptionKeySource"`
	Etag                      string                                  `pulumi:"etag"`
	ExportPolicy              *VolumePropertiesResponseExportPolicy   `pulumi:"exportPolicy"`
	FileSystemId              string                                  `pulumi:"fileSystemId"`
	Id                        string                                  `pulumi:"id"`
	IsDefaultQuotaEnabled     *bool                                   `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring               *bool                                   `pulumi:"isRestoring"`
	KerberosEnabled           *bool                                   `pulumi:"kerberosEnabled"`
	LdapEnabled               *bool                                   `pulumi:"ldapEnabled"`
	Location                  string                                  `pulumi:"location"`
	MountTargets              []MountTargetPropertiesResponse         `pulumi:"mountTargets"`
	Name                      string                                  `pulumi:"name"`
	NetworkFeatures           *string                                 `pulumi:"networkFeatures"`
	NetworkSiblingSetId       string                                  `pulumi:"networkSiblingSetId"`
	PlacementRules            []PlacementKeyValuePairsResponse        `pulumi:"placementRules"`
	ProtocolTypes             []string                                `pulumi:"protocolTypes"`
	ProvisioningState         string                                  `pulumi:"provisioningState"`
	ProximityPlacementGroup   *string                                 `pulumi:"proximityPlacementGroup"`
	SecurityStyle             *string                                 `pulumi:"securityStyle"`
	ServiceLevel              *string                                 `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable  *bool                                   `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption             *bool                                   `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible  *bool                                   `pulumi:"snapshotDirectoryVisible"`
	SnapshotId                *string                                 `pulumi:"snapshotId"`
	StorageToNetworkProximity string                                  `pulumi:"storageToNetworkProximity"`
	SubnetId                  string                                  `pulumi:"subnetId"`
	T2Network                 string                                  `pulumi:"t2Network"`
	Tags                      map[string]string                       `pulumi:"tags"`
	ThroughputMibps           *float64                                `pulumi:"throughputMibps"`
	Type                      string                                  `pulumi:"type"`
	UnixPermissions           *string                                 `pulumi:"unixPermissions"`
	UsageThreshold            float64                                 `pulumi:"usageThreshold"`
	VolumeGroupName           string                                  `pulumi:"volumeGroupName"`
	VolumeSpecName            *string                                 `pulumi:"volumeSpecName"`
	VolumeType                *string                                 `pulumi:"volumeType"`
}


func (val *LookupVolumeResult) Defaults() *LookupVolumeResult {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AvsDataStore) {
		avsDataStore_ := "Disabled"
		tmp.AvsDataStore = &avsDataStore_
	}
	if isZero(tmp.CoolAccess) {
		coolAccess_ := false
		tmp.CoolAccess = &coolAccess_
	}
	if isZero(tmp.DefaultGroupQuotaInKiBs) {
		defaultGroupQuotaInKiBs_ := 0.0
		tmp.DefaultGroupQuotaInKiBs = &defaultGroupQuotaInKiBs_
	}
	if isZero(tmp.DefaultUserQuotaInKiBs) {
		defaultUserQuotaInKiBs_ := 0.0
		tmp.DefaultUserQuotaInKiBs = &defaultUserQuotaInKiBs_
	}
	if isZero(tmp.IsDefaultQuotaEnabled) {
		isDefaultQuotaEnabled_ := false
		tmp.IsDefaultQuotaEnabled = &isDefaultQuotaEnabled_
	}
	if isZero(tmp.KerberosEnabled) {
		kerberosEnabled_ := false
		tmp.KerberosEnabled = &kerberosEnabled_
	}
	if isZero(tmp.LdapEnabled) {
		ldapEnabled_ := false
		tmp.LdapEnabled = &ldapEnabled_
	}
	if isZero(tmp.NetworkFeatures) {
		networkFeatures_ := "Basic"
		tmp.NetworkFeatures = &networkFeatures_
	}
	if isZero(tmp.SecurityStyle) {
		securityStyle_ := "unix"
		tmp.SecurityStyle = &securityStyle_
	}
	if isZero(tmp.SmbContinuouslyAvailable) {
		smbContinuouslyAvailable_ := false
		tmp.SmbContinuouslyAvailable = &smbContinuouslyAvailable_
	}
	if isZero(tmp.SmbEncryption) {
		smbEncryption_ := false
		tmp.SmbEncryption = &smbEncryption_
	}
	if isZero(tmp.SnapshotDirectoryVisible) {
		snapshotDirectoryVisible_ := true
		tmp.SnapshotDirectoryVisible = &snapshotDirectoryVisible_
	}
	if isZero(tmp.ThroughputMibps) {
		throughputMibps_ := 0.0
		tmp.ThroughputMibps = &throughputMibps_
	}
	if isZero(tmp.UnixPermissions) {
		unixPermissions_ := "0770"
		tmp.UnixPermissions = &unixPermissions_
	}
	if isZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
	}
	return &tmp
}
