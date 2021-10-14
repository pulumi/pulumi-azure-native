


package v20190801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	BaremetalTenantId pulumi.StringOutput                             `pulumi:"baremetalTenantId"`
	CreationToken     pulumi.StringOutput                             `pulumi:"creationToken"`
	DataProtection    VolumePropertiesResponseDataProtectionPtrOutput `pulumi:"dataProtection"`
	ExportPolicy      VolumePropertiesResponseExportPolicyPtrOutput   `pulumi:"exportPolicy"`
	FileSystemId      pulumi.StringOutput                             `pulumi:"fileSystemId"`
	Location          pulumi.StringOutput                             `pulumi:"location"`
	MountTargets      MountTargetPropertiesResponseArrayOutput        `pulumi:"mountTargets"`
	Name              pulumi.StringOutput                             `pulumi:"name"`
	ProtocolTypes     pulumi.StringArrayOutput                        `pulumi:"protocolTypes"`
	ProvisioningState pulumi.StringOutput                             `pulumi:"provisioningState"`
	ServiceLevel      pulumi.StringPtrOutput                          `pulumi:"serviceLevel"`
	SnapshotId        pulumi.StringPtrOutput                          `pulumi:"snapshotId"`
	SubnetId          pulumi.StringOutput                             `pulumi:"subnetId"`
	Tags              pulumi.StringMapOutput                          `pulumi:"tags"`
	Type              pulumi.StringOutput                             `pulumi:"type"`
	UsageThreshold    pulumi.Float64Output                            `pulumi:"usageThreshold"`
	UsedBytes         pulumi.Float64Output                            `pulumi:"usedBytes"`
	VolumeType        pulumi.StringPtrOutput                          `pulumi:"volumeType"`
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
	if args.ServiceLevel == nil {
		args.ServiceLevel = pulumi.StringPtr("Premium")
	}
	if args.UsageThreshold == nil {
		args.UsageThreshold = pulumi.Float64(107374182400)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20170815:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20170815:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190501:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190601:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20190701:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20190701:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191001:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191001:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20191101:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20191101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200201:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200301:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200301:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200501:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200501:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200601:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200601:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200701:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200701:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200801:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200801:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20200901:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20200901:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201101:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201101:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20201201:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20201201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210201:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210201:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210401preview:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210401preview:Volume"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20210601:Volume"),
		},
		{
			Type: pulumi.String("azure-nextgen:netapp/v20210601:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-native:netapp/v20190801:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:netapp/v20190801:Volume", name, id, state, &resource, opts...)
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
	AccountName       string                          `pulumi:"accountName"`
	CreationToken     string                          `pulumi:"creationToken"`
	DataProtection    *VolumePropertiesDataProtection `pulumi:"dataProtection"`
	ExportPolicy      *VolumePropertiesExportPolicy   `pulumi:"exportPolicy"`
	Location          *string                         `pulumi:"location"`
	MountTargets      []MountTargetProperties         `pulumi:"mountTargets"`
	PoolName          string                          `pulumi:"poolName"`
	ProtocolTypes     []string                        `pulumi:"protocolTypes"`
	ResourceGroupName string                          `pulumi:"resourceGroupName"`
	ServiceLevel      *string                         `pulumi:"serviceLevel"`
	SnapshotId        *string                         `pulumi:"snapshotId"`
	SubnetId          string                          `pulumi:"subnetId"`
	Tags              map[string]string               `pulumi:"tags"`
	UsageThreshold    float64                         `pulumi:"usageThreshold"`
	VolumeName        *string                         `pulumi:"volumeName"`
	VolumeType        *string                         `pulumi:"volumeType"`
}


type VolumeArgs struct {
	AccountName       pulumi.StringInput
	CreationToken     pulumi.StringInput
	DataProtection    VolumePropertiesDataProtectionPtrInput
	ExportPolicy      VolumePropertiesExportPolicyPtrInput
	Location          pulumi.StringPtrInput
	MountTargets      MountTargetPropertiesArrayInput
	PoolName          pulumi.StringInput
	ProtocolTypes     pulumi.StringArrayInput
	ResourceGroupName pulumi.StringInput
	ServiceLevel      pulumi.StringPtrInput
	SnapshotId        pulumi.StringPtrInput
	SubnetId          pulumi.StringInput
	Tags              pulumi.StringMapInput
	UsageThreshold    pulumi.Float64Input
	VolumeName        pulumi.StringPtrInput
	VolumeType        pulumi.StringPtrInput
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
	return reflect.TypeOf((*Volume)(nil))
}

func (i *Volume) ToVolumeOutput() VolumeOutput {
	return i.ToVolumeOutputWithContext(context.Background())
}

func (i *Volume) ToVolumeOutputWithContext(ctx context.Context) VolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeOutput)
}

type VolumeOutput struct{ *pulumi.OutputState }

func (VolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Volume)(nil))
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
