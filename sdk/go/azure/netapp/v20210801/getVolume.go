


package v20210801

import (
	"context"
	"reflect"

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
	SystemData                SystemDataResponse                      `pulumi:"systemData"`
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
	if isZero(tmp.UnixPermissions) {
		unixPermissions_ := "0770"
		tmp.UnixPermissions = &unixPermissions_
	}
	if isZero(tmp.UsageThreshold) {
		tmp.UsageThreshold = 107374182400.0
	}
	return &tmp
}

func LookupVolumeOutput(ctx *pulumi.Context, args LookupVolumeOutputArgs, opts ...pulumi.InvokeOption) LookupVolumeResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVolumeResult, error) {
			args := v.(LookupVolumeArgs)
			r, err := LookupVolume(ctx, &args, opts...)
			var s LookupVolumeResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVolumeResultOutput)
}

type LookupVolumeOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PoolName          pulumi.StringInput `pulumi:"poolName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	VolumeName        pulumi.StringInput `pulumi:"volumeName"`
}

func (LookupVolumeOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeArgs)(nil)).Elem()
}


type LookupVolumeResultOutput struct{ *pulumi.OutputState }

func (LookupVolumeResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVolumeResult)(nil)).Elem()
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutput() LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) ToLookupVolumeResultOutputWithContext(ctx context.Context) LookupVolumeResultOutput {
	return o
}

func (o LookupVolumeResultOutput) AvsDataStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.AvsDataStore }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) CapacityPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.CapacityPoolResourceId }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) CloneProgress() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVolumeResult) int { return v.CloneProgress }).(pulumi.IntOutput)
}

func (o LookupVolumeResultOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) CoolnessPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *int { return v.CoolnessPeriod }).(pulumi.IntPtrOutput)
}

func (o LookupVolumeResultOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.CreationToken }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseDataProtection { return v.DataProtection }).(VolumePropertiesResponseDataProtectionPtrOutput)
}

func (o LookupVolumeResultOutput) DefaultGroupQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.DefaultGroupQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o LookupVolumeResultOutput) DefaultUserQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.DefaultUserQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o LookupVolumeResultOutput) EncryptionKeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.EncryptionKeySource }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *VolumePropertiesResponseExportPolicy { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o LookupVolumeResultOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) IsDefaultQuotaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.IsDefaultQuotaEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) LdapEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.LdapEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []MountTargetPropertiesResponse { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

func (o LookupVolumeResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) NetworkFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.NetworkFeatures }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) NetworkSiblingSetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.NetworkSiblingSetId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) PlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []PlacementKeyValuePairsResponse { return v.PlacementRules }).(PlacementKeyValuePairsResponseArrayOutput)
}

func (o LookupVolumeResultOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupVolumeResult) []string { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

func (o LookupVolumeResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) ProximityPlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.ProximityPlacementGroup }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) SmbContinuouslyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.SmbContinuouslyAvailable }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) SmbEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.SmbEncryption }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *bool { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

func (o LookupVolumeResultOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) StorageToNetworkProximity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.StorageToNetworkProximity }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.SubnetId }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupVolumeResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupVolumeResultOutput) T2Network() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.T2Network }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVolumeResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVolumeResultOutput) ThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *float64 { return v.ThroughputMibps }).(pulumi.Float64PtrOutput)
}

func (o LookupVolumeResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) UnixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.UnixPermissions }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVolumeResult) float64 { return v.UsageThreshold }).(pulumi.Float64Output)
}

func (o LookupVolumeResultOutput) VolumeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVolumeResult) string { return v.VolumeGroupName }).(pulumi.StringOutput)
}

func (o LookupVolumeResultOutput) VolumeSpecName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.VolumeSpecName }).(pulumi.StringPtrOutput)
}

func (o LookupVolumeResultOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVolumeResult) *string { return v.VolumeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVolumeResultOutput{})
}
