


package v20200701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Volume struct {
	pulumi.CustomResourceState

	BackupId                 pulumi.StringPtrOutput                          `pulumi:"backupId"`
	BaremetalTenantId        pulumi.StringOutput                             `pulumi:"baremetalTenantId"`
	CreationToken            pulumi.StringOutput                             `pulumi:"creationToken"`
	DataProtection           VolumePropertiesResponseDataProtectionPtrOutput `pulumi:"dataProtection"`
	ExportPolicy             VolumePropertiesResponseExportPolicyPtrOutput   `pulumi:"exportPolicy"`
	FileSystemId             pulumi.StringOutput                             `pulumi:"fileSystemId"`
	IsRestoring              pulumi.BoolPtrOutput                            `pulumi:"isRestoring"`
	KerberosEnabled          pulumi.BoolPtrOutput                            `pulumi:"kerberosEnabled"`
	Location                 pulumi.StringOutput                             `pulumi:"location"`
	MountTargets             MountTargetPropertiesResponseArrayOutput        `pulumi:"mountTargets"`
	Name                     pulumi.StringOutput                             `pulumi:"name"`
	ProtocolTypes            pulumi.StringArrayOutput                        `pulumi:"protocolTypes"`
	ProvisioningState        pulumi.StringOutput                             `pulumi:"provisioningState"`
	SecurityStyle            pulumi.StringPtrOutput                          `pulumi:"securityStyle"`
	ServiceLevel             pulumi.StringPtrOutput                          `pulumi:"serviceLevel"`
	SnapshotDirectoryVisible pulumi.BoolPtrOutput                            `pulumi:"snapshotDirectoryVisible"`
	SnapshotId               pulumi.StringPtrOutput                          `pulumi:"snapshotId"`
	SubnetId                 pulumi.StringOutput                             `pulumi:"subnetId"`
	Tags                     pulumi.StringMapOutput                          `pulumi:"tags"`
	ThroughputMibps          pulumi.Float64PtrOutput                         `pulumi:"throughputMibps"`
	Type                     pulumi.StringOutput                             `pulumi:"type"`
	UsageThreshold           pulumi.Float64Output                            `pulumi:"usageThreshold"`
	VolumeType               pulumi.StringPtrOutput                          `pulumi:"volumeType"`
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
	if isZero(args.KerberosEnabled) {
		args.KerberosEnabled = pulumi.BoolPtr(false)
	}
	if isZero(args.SecurityStyle) {
		args.SecurityStyle = pulumi.StringPtr("unix")
	}
	if isZero(args.ServiceLevel) {
		args.ServiceLevel = pulumi.StringPtr("Premium")
	}
	if isZero(args.SnapshotDirectoryVisible) {
		args.SnapshotDirectoryVisible = pulumi.BoolPtr(true)
	}
	if isZero(args.ThroughputMibps) {
		args.ThroughputMibps = pulumi.Float64Ptr(0.0)
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
			Type: pulumi.String("azure-native:netapp/v20211001:Volume"),
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
	err := ctx.RegisterResource("azure-native:netapp/v20200701:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:netapp/v20200701:Volume", name, id, state, &resource, opts...)
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
	BackupId                 *string                         `pulumi:"backupId"`
	CreationToken            string                          `pulumi:"creationToken"`
	DataProtection           *VolumePropertiesDataProtection `pulumi:"dataProtection"`
	ExportPolicy             *VolumePropertiesExportPolicy   `pulumi:"exportPolicy"`
	IsRestoring              *bool                           `pulumi:"isRestoring"`
	KerberosEnabled          *bool                           `pulumi:"kerberosEnabled"`
	Location                 *string                         `pulumi:"location"`
	PoolName                 string                          `pulumi:"poolName"`
	ProtocolTypes            []string                        `pulumi:"protocolTypes"`
	ResourceGroupName        string                          `pulumi:"resourceGroupName"`
	SecurityStyle            *string                         `pulumi:"securityStyle"`
	ServiceLevel             *string                         `pulumi:"serviceLevel"`
	SnapshotDirectoryVisible *bool                           `pulumi:"snapshotDirectoryVisible"`
	SnapshotId               *string                         `pulumi:"snapshotId"`
	SubnetId                 string                          `pulumi:"subnetId"`
	Tags                     map[string]string               `pulumi:"tags"`
	ThroughputMibps          *float64                        `pulumi:"throughputMibps"`
	UsageThreshold           float64                         `pulumi:"usageThreshold"`
	VolumeName               *string                         `pulumi:"volumeName"`
	VolumeType               *string                         `pulumi:"volumeType"`
}


type VolumeArgs struct {
	AccountName              pulumi.StringInput
	BackupId                 pulumi.StringPtrInput
	CreationToken            pulumi.StringInput
	DataProtection           VolumePropertiesDataProtectionPtrInput
	ExportPolicy             VolumePropertiesExportPolicyPtrInput
	IsRestoring              pulumi.BoolPtrInput
	KerberosEnabled          pulumi.BoolPtrInput
	Location                 pulumi.StringPtrInput
	PoolName                 pulumi.StringInput
	ProtocolTypes            pulumi.StringArrayInput
	ResourceGroupName        pulumi.StringInput
	SecurityStyle            pulumi.StringPtrInput
	ServiceLevel             pulumi.StringPtrInput
	SnapshotDirectoryVisible pulumi.BoolPtrInput
	SnapshotId               pulumi.StringPtrInput
	SubnetId                 pulumi.StringInput
	Tags                     pulumi.StringMapInput
	ThroughputMibps          pulumi.Float64PtrInput
	UsageThreshold           pulumi.Float64Input
	VolumeName               pulumi.StringPtrInput
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

func (o VolumeOutput) BackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.BackupId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) BaremetalTenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.BaremetalTenantId }).(pulumi.StringOutput)
}

func (o VolumeOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CreationToken }).(pulumi.StringOutput)
}

func (o VolumeOutput) DataProtection() VolumePropertiesResponseDataProtectionPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseDataProtectionPtrOutput { return v.DataProtection }).(VolumePropertiesResponseDataProtectionPtrOutput)
}

func (o VolumeOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseExportPolicyPtrOutput { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o VolumeOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o VolumeOutput) IsRestoring() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.IsRestoring }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) KerberosEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.KerberosEnabled }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VolumeOutput) MountTargets() MountTargetPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *Volume) MountTargetPropertiesResponseArrayOutput { return v.MountTargets }).(MountTargetPropertiesResponseArrayOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeOutput) ProtocolTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringArrayOutput { return v.ProtocolTypes }).(pulumi.StringArrayOutput)
}

func (o VolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeOutput) SecurityStyle() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SecurityStyle }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) ServiceLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.ServiceLevel }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) SnapshotDirectoryVisible() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.BoolPtrOutput { return v.SnapshotDirectoryVisible }).(pulumi.BoolPtrOutput)
}

func (o VolumeOutput) SnapshotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SnapshotId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) SubnetId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.SubnetId }).(pulumi.StringOutput)
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

func (o VolumeOutput) UsageThreshold() pulumi.Float64Output {
	return o.ApplyT(func(v *Volume) pulumi.Float64Output { return v.UsageThreshold }).(pulumi.Float64Output)
}

func (o VolumeOutput) VolumeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.VolumeType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
