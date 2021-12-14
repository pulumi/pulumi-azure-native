


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	BackupId                  pulumi.StringPtrOutput                          `pulumi:"backupId"`
	BaremetalTenantId         pulumi.StringOutput                             `pulumi:"baremetalTenantId"`
	CreationToken             pulumi.StringOutput                             `pulumi:"creationToken"`
	DataProtection            VolumePropertiesResponseDataProtectionPtrOutput `pulumi:"dataProtection"`
	EncryptionKeySource       pulumi.StringPtrOutput                          `pulumi:"encryptionKeySource"`
	ExportPolicy              VolumePropertiesResponseExportPolicyPtrOutput   `pulumi:"exportPolicy"`
	FileSystemId              pulumi.StringOutput                             `pulumi:"fileSystemId"`
	IsRestoring               pulumi.BoolPtrOutput                            `pulumi:"isRestoring"`
	KerberosEnabled           pulumi.BoolPtrOutput                            `pulumi:"kerberosEnabled"`
	LdapEnabled               pulumi.BoolPtrOutput                            `pulumi:"ldapEnabled"`
	Location                  pulumi.StringOutput                             `pulumi:"location"`
	MountTargets              MountTargetPropertiesResponseArrayOutput        `pulumi:"mountTargets"`
	Name                      pulumi.StringOutput                             `pulumi:"name"`
	NetworkFeatures           pulumi.StringPtrOutput                          `pulumi:"networkFeatures"`
	NetworkSiblingSetId       pulumi.StringOutput                             `pulumi:"networkSiblingSetId"`
	ProtocolTypes             pulumi.StringArrayOutput                        `pulumi:"protocolTypes"`
	ProvisioningState         pulumi.StringOutput                             `pulumi:"provisioningState"`
	SecurityStyle             pulumi.StringPtrOutput                          `pulumi:"securityStyle"`
	ServiceLevel              pulumi.StringPtrOutput                          `pulumi:"serviceLevel"`
	SmbContinuouslyAvailable  pulumi.BoolPtrOutput                            `pulumi:"smbContinuouslyAvailable"`
	SmbEncryption             pulumi.BoolPtrOutput                            `pulumi:"smbEncryption"`
	SnapshotDirectoryVisible  pulumi.BoolPtrOutput                            `pulumi:"snapshotDirectoryVisible"`
	SnapshotId                pulumi.StringPtrOutput                          `pulumi:"snapshotId"`
	StorageToNetworkProximity pulumi.StringOutput                             `pulumi:"storageToNetworkProximity"`
	SubnetId                  pulumi.StringOutput                             `pulumi:"subnetId"`
	Tags                      pulumi.StringMapOutput                          `pulumi:"tags"`
	ThroughputMibps           pulumi.Float64PtrOutput                         `pulumi:"throughputMibps"`
	Type                      pulumi.StringOutput                             `pulumi:"type"`
	UsageThreshold            pulumi.Float64Output                            `pulumi:"usageThreshold"`
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
	if isZero(args.ServiceLevel) {
		args.ServiceLevel = pulumi.StringPtr("Premium")
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
			Type: pulumi.String("azure-native:netapp/v20210601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210801:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-native:netapp/v20210401preview:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:netapp/v20210401preview:Volume", name, id, state, &resource, opts...)
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
	EncryptionKeySource      *string                         `pulumi:"encryptionKeySource"`
	ExportPolicy             *VolumePropertiesExportPolicy   `pulumi:"exportPolicy"`
	IsRestoring              *bool                           `pulumi:"isRestoring"`
	KerberosEnabled          *bool                           `pulumi:"kerberosEnabled"`
	LdapEnabled              *bool                           `pulumi:"ldapEnabled"`
	Location                 *string                         `pulumi:"location"`
	NetworkFeatures          *string                         `pulumi:"networkFeatures"`
	PoolName                 string                          `pulumi:"poolName"`
	ProtocolTypes            []string                        `pulumi:"protocolTypes"`
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
	UsageThreshold           float64                         `pulumi:"usageThreshold"`
	VolumeName               *string                         `pulumi:"volumeName"`
	VolumeType               *string                         `pulumi:"volumeType"`
}


type VolumeArgs struct {
	AccountName              pulumi.StringInput
	BackupId                 pulumi.StringPtrInput
	CreationToken            pulumi.StringInput
	DataProtection           VolumePropertiesDataProtectionPtrInput
	EncryptionKeySource      pulumi.StringPtrInput
	ExportPolicy             VolumePropertiesExportPolicyPtrInput
	IsRestoring              pulumi.BoolPtrInput
	KerberosEnabled          pulumi.BoolPtrInput
	LdapEnabled              pulumi.BoolPtrInput
	Location                 pulumi.StringPtrInput
	NetworkFeatures          pulumi.StringPtrInput
	PoolName                 pulumi.StringInput
	ProtocolTypes            pulumi.StringArrayInput
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

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
