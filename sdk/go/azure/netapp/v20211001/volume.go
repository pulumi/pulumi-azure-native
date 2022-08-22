


package v20211001

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	AvsDataStore              pulumi.StringPtrOutput                          `pulumi:"avsDataStore"`
	BackupId                  pulumi.StringPtrOutput                          `pulumi:"backupId"`
	BaremetalTenantId         pulumi.StringOutput                             `pulumi:"baremetalTenantId"`
	CapacityPoolResourceId    pulumi.StringPtrOutput                          `pulumi:"capacityPoolResourceId"`
	CloneProgress             pulumi.IntOutput                                `pulumi:"cloneProgress"`
	CoolAccess                pulumi.BoolPtrOutput                            `pulumi:"coolAccess"`
	CoolnessPeriod            pulumi.IntPtrOutput                             `pulumi:"coolnessPeriod"`
	CreationToken             pulumi.StringOutput                             `pulumi:"creationToken"`
	DataProtection            VolumePropertiesResponseDataProtectionPtrOutput `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs   pulumi.Float64PtrOutput                         `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs    pulumi.Float64PtrOutput                         `pulumi:"defaultUserQuotaInKiBs"`
	EnableSubvolumes          pulumi.StringPtrOutput                          `pulumi:"enableSubvolumes"`
	EncryptionKeySource       pulumi.StringPtrOutput                          `pulumi:"encryptionKeySource"`
	Etag                      pulumi.StringOutput                             `pulumi:"etag"`
	ExportPolicy              VolumePropertiesResponseExportPolicyPtrOutput   `pulumi:"exportPolicy"`
	FileSystemId              pulumi.StringOutput                             `pulumi:"fileSystemId"`
	IsDefaultQuotaEnabled     pulumi.BoolPtrOutput                            `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring               pulumi.BoolPtrOutput                            `pulumi:"isRestoring"`
	KerberosEnabled           pulumi.BoolPtrOutput                            `pulumi:"kerberosEnabled"`
	LdapEnabled               pulumi.BoolPtrOutput                            `pulumi:"ldapEnabled"`
	Location                  pulumi.StringOutput                             `pulumi:"location"`
	MaximumNumberOfFiles      pulumi.Float64Output                            `pulumi:"maximumNumberOfFiles"`
	MountTargets              MountTargetPropertiesResponseArrayOutput        `pulumi:"mountTargets"`
	Name                      pulumi.StringOutput                             `pulumi:"name"`
	NetworkFeatures           pulumi.StringPtrOutput                          `pulumi:"networkFeatures"`
	NetworkSiblingSetId       pulumi.StringOutput                             `pulumi:"networkSiblingSetId"`
	PlacementRules            PlacementKeyValuePairsResponseArrayOutput       `pulumi:"placementRules"`
	ProtocolTypes             pulumi.StringArrayOutput                        `pulumi:"protocolTypes"`
	ProvisioningState         pulumi.StringOutput                             `pulumi:"provisioningState"`
	ProximityPlacementGroup   pulumi.StringPtrOutput                          `pulumi:"proximityPlacementGroup"`
	SecurityStyle             pulumi.StringPtrOutput                          `pulumi:"securityStyle"`
	ServiceLevel              pulumi.StringPtrOutput                          `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable  pulumi.BoolPtrOutput                            `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption             pulumi.BoolPtrOutput                            `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible  pulumi.BoolPtrOutput                            `pulumi:"snapshotDirectoryVisible"`
	SnapshotId                pulumi.StringPtrOutput                          `pulumi:"snapshotId"`
	StorageToNetworkProximity pulumi.StringOutput                             `pulumi:"storageToNetworkProximity"`
	SubnetId                  pulumi.StringOutput                             `pulumi:"subnetId"`
	SystemData                SystemDataResponseOutput                        `pulumi:"systemData"`
	T2Network                 pulumi.StringOutput                             `pulumi:"t2Network"`
	Tags                      pulumi.StringMapOutput                          `pulumi:"tags"`
	ThroughputMibps           pulumi.Float64PtrOutput                         `pulumi:"throughputMibps"`
	Type                      pulumi.StringOutput                             `pulumi:"type"`
	UnixPermissions           pulumi.StringPtrOutput                          `pulumi:"unixPermissions"`
	UsageThreshold            pulumi.Float64Output                            `pulumi:"usageThreshold"`
	VolumeGroupName           pulumi.StringOutput                             `pulumi:"volumeGroupName"`
	VolumeSpecName            pulumi.StringPtrOutput                          `pulumi:"volumeSpecName"`
	VolumeType                pulumi.StringPtrOutput                          `pulumi:"volumeType"`
}


func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.CreationToken == nil {
		return nil, errors.New("invalid value for required argument 'CreationToken'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SubnetId == nil {
		return nil, errors.New("invalid value for required argument 'SubnetId'")
	}
	if isZero(args.AvsDataStore) {
		args.AvsDataStore = pulumi.StringPtr("Disabled")
	}
	if isZero(args.CoolAccess) {
		args.CoolAccess = pulumi.BoolPtr(false)
	}
	if isZero(args.DefaultGroupQuotaInKiBs) {
		args.DefaultGroupQuotaInKiBs = pulumi.Float64Ptr(0.0)
	}
	if isZero(args.DefaultUserQuotaInKiBs) {
		args.DefaultUserQuotaInKiBs = pulumi.Float64Ptr(0.0)
	}
	if isZero(args.EnableSubvolumes) {
		args.EnableSubvolumes = pulumi.StringPtr("Disabled")
	}
	if isZero(args.IsDefaultQuotaEnabled) {
		args.IsDefaultQuotaEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.KerberosEnabled) {
		args.KerberosEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.LdapEnabled) {
		args.LdapEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.NetworkFeatures) {
		args.NetworkFeatures = pulumi.StringPtr("Basic")
	}
	if isZero(args.SecurityStyle) {
		args.SecurityStyle = pulumi.StringPtr("unix")
	}
	if isZero(args.SmbContinuouslyAvailable) {
		args.SmbContinuouslyAvailable = pulumi.BoolPtr(false)
	}
	if isZero(args.SmbEncryption) {
		args.SmbEncryption = pulumi.BoolPtr(false)
	}
	if isZero(args.SnapshotDirectoryVisible) {
		args.SnapshotDirectoryVisible = pulumi.BoolPtr(true)
	}
	if isZero(args.UnixPermissions) {
		args.UnixPermissions = pulumi.StringPtr("0770")
	}
	if isZero(args.UsageThreshold) {
		args.UsageThreshold = pulumi.Float64(107374182400.0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-native:netapp/v20211001:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:netapp/v20211001:Volume", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type volumeState struct {
}

type VolumeState struct {
}

func (VolumeState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeState)(nil)).Elem()
}

type volumeArgs struct {
	AccountName              string                          `pulumi:"accountName"`
	AvsDataStore             *string                         `pulumi:"avsDataStore"`
	BackupId                 *string                         `pulumi:"backupId"`
	CapacityPoolResourceId   *string                         `pulumi:"capacityPoolResourceId"`
	CoolAccess               *bool                           `pulumi:"coolAccess"`
	CoolnessPeriod           *int                            `pulumi:"coolnessPeriod"`
	CreationToken            string                          `pulumi:"creationToken"`
	DataProtection           *VolumePropertiesDataProtection `pulumi:"dataProtection"`
	DefaultGroupQuotaInKiBs  *float64                        `pulumi:"defaultGroupQuotaInKiBs"`
	DefaultUserQuotaInKiBs   *float64                        `pulumi:"defaultUserQuotaInKiBs"`
	EnableSubvolumes         *string                         `pulumi:"enableSubvolumes"`
	EncryptionKeySource      *string                         `pulumi:"encryptionKeySource"`
	ExportPolicy             *VolumePropertiesExportPolicy   `pulumi:"exportPolicy"`
	IsDefaultQuotaEnabled    *bool                           `pulumi:"isDefaultQuotaEnabled"`
	IsRestoring              *bool                           `pulumi:"isRestoring"`
	KerberosEnabled          *bool                           `pulumi:"kerberosEnabled"`
	LdapEnabled              *bool                           `pulumi:"ldapEnabled"`
	Location                 *string                         `pulumi:"location"`
	NetworkFeatures          *string                         `pulumi:"networkFeatures"`
	PlacementRules           []PlacementKeyValuePairs        `pulumi:"placementRules"`
	PoolName                 string                          `pulumi:"poolName"`
	ProtocolTypes            []string                        `pulumi:"protocolTypes"`
	ProximityPlacementGroup  *string                         `pulumi:"proximityPlacementGroup"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	SecurityStyle            *string                         `pulumi:"securityStyle"`
	ServiceLevel             *string                         `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable *bool                           `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption            *bool                           `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible *bool                           `pulumi:"snapshotDirectoryVisible"`
	SnapshotId               *string                         `pulumi:"snapshotId"`
	SubnetId                 string                          `pulumi:"subnetId"`
	Tags                     map[string]string               `pulumi:"tags"`
	ThroughputMibps          *float64                        `pulumi:"throughputMibps"`
	UnixPermissions          *string                         `pulumi:"unixPermissions"`
	UsageThreshold           float64                         `pulumi:"usageThreshold"`
	VolumeName               *string                         `pulumi:"volumeName"`
	VolumeSpecName           *string                         `pulumi:"volumeSpecName"`
	VolumeType               *string                         `pulumi:"volumeType"`
}


type VolumeArgs struct {
	AccountName              pulumi.StringInput
	AvsDataStore             pulumi.StringPtrInput
	BackupId                 pulumi.StringPtrInput
	CapacityPoolResourceId   pulumi.StringPtrInput
	CoolAccess               pulumi.BoolPtrInput
	CoolnessPeriod           pulumi.IntPtrInput
	CreationToken            pulumi.StringInput
	DataProtection           VolumePropertiesDataProtectionPtrInput
	DefaultGroupQuotaInKiBs  pulumi.Float64PtrInput
	DefaultUserQuotaInKiBs   pulumi.Float64PtrInput
	EnableSubvolumes         pulumi.StringPtrInput
	EncryptionKeySource      pulumi.StringPtrInput
	ExportPolicy             VolumePropertiesExportPolicyPtrInput
	IsDefaultQuotaEnabled    pulumi.BoolPtrInput
	IsRestoring              pulumi.BoolPtrInput
	KerberosEnabled          pulumi.BoolPtrInput
	LdapEnabled              pulumi.BoolPtrInput
	Location                 pulumi.StringPtrInput
	NetworkFeatures          pulumi.StringPtrInput
	PlacementRules           PlacementKeyValuePairsArrayInput
	PoolName                 pulumi.StringInput
	ProtocolTypes            pulumi.StringArrayInput
	ProximityPlacementGroup  pulumi.StringPtrInput
	ResourceGroupName        pulumi.StringInput
	SecurityStyle            pulumi.StringPtrInput
	ServiceLevel             pulumi.StringPtrInput
	SmbContinuouslyAvailable pulumi.BoolPtrInput
	SmbEncryption            pulumi.BoolPtrInput
	SnapshotDirectoryVisible pulumi.BoolPtrInput
	SnapshotId               pulumi.StringPtrInput
	SubnetId                 pulumi.StringInput
	Tags                     pulumi.StringMapInput
	ThroughputMibps          pulumi.Float64PtrInput
	UnixPermissions          pulumi.StringPtrInput
	UsageThreshold           pulumi.Float64Input
	VolumeName               pulumi.StringPtrInput
	VolumeSpecName           pulumi.StringPtrInput
	VolumeType               pulumi.StringPtrInput
}

func (VolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeArgs)(nil)).Elem()
}

type VolumeInput interface {
	pulumi.Input

	ToVolumeOutput() VolumeOutput
	ToVolumeOutputWithContext(ctx context.Context) VolumeOutput
}

func (*Volume) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Volume)(nil)).Elem()
}

func (o VolumeOutput) ToVolumeOutput() VolumeOutput {
	return o
}

func (o VolumeOutput) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return o
}

func (o VolumeOutput) AvsDataStore() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.AvsDataStore }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

func (o VolumeOutput) CapacityPoolResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.CapacityPoolResourceId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) CloneProgress() pulumi.IntOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntOutput { return v.CloneProgress }).(pulumi.IntOutput)
}

func (o VolumeOutput) CoolAccess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.CoolAccess }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) CoolnessPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.IntPtrOutput { return v.CoolnessPeriod }).(pulumi.IntPtrOutput)
}

func (o VolumeOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CreationToken }).(pulumi.StringOutput)
}

func (o VolumeOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseDataProtectionPtrOutput { return v.DataProtection }).(VolumePropertiesResponseDataProtectionPtrOutput)
}

func (o VolumeOutput) DefaultGroupQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.DefaultGroupQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeOutput) DefaultUserQuotaInKiBs() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.DefaultUserQuotaInKiBs }).(pulumi.Float64PtrOutput)
}

func (o VolumeOutput) EnableSubvolumes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.EnableSubvolumes }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) EncryptionKeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.EncryptionKeySource }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o VolumeOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseExportPolicyPtrOutput { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o VolumeOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o VolumeOutput) IsDefaultQuotaEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.IsDefaultQuotaEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) LdapEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.LdapEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VolumeOutput) MaximumNumberOfFiles() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.MaximumNumberOfFiles }).(pulumi.Float64Output)
}

func (o VolumeOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *Volume) MountTargetPropertiesResponseArrayOutput { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeOutput) NetworkFeatures() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.NetworkFeatures }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) NetworkSiblingSetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.NetworkSiblingSetId }).(pulumi.StringOutput)
}

func (o VolumeOutput) PlacementRules() PlacementKeyValuePairsResponseArrayOutput {
	return o.ApplyT(func(v *Volume) PlacementKeyValuePairsResponseArrayOutput { return v.PlacementRules }).(PlacementKeyValuePairsResponseArrayOutput)
}

func (o VolumeOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

func (o VolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeOutput) ProximityPlacementGroup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ProximityPlacementGroup }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) SmbContinuouslyAvailable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SmbContinuouslyAvailable }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) SmbEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SmbEncryption }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) StorageToNetworkProximity() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.StorageToNetworkProximity }).(pulumi.StringOutput)
}

func (o VolumeOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
}

func (o VolumeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Volume) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VolumeOutput) T2Network() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.T2Network }).(pulumi.StringOutput)
}

func (o VolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeOutput) ThroughputMibps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.ThroughputMibps }).(pulumi.Float64PtrOutput)
}

func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeOutput) UnixPermissions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.UnixPermissions }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.UsageThreshold }).(pulumi.Float64Output)
}

func (o VolumeOutput) VolumeGroupName() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeGroupName }).(pulumi.StringOutput)
}

func (o VolumeOutput) VolumeSpecName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.VolumeSpecName }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
