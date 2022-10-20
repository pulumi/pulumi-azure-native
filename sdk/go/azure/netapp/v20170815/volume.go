


package v20170815

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type Volume struct {
	pulumi.CustomResourceState

	CreationToken     pulumi.StringOutput                           `pulumi:"creationToken"`
	ExportPolicy      VolumePropertiesResponseExportPolicyPtrOutput `pulumi:"exportPolicy"`
	FileSystemId      pulumi.StringOutput                           `pulumi:"fileSystemId"`
	Location          pulumi.StringOutput                           `pulumi:"location"`
	Name              pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                           `pulumi:"provisioningState"`
	ServiceLevel      pulumi.StringOutput                           `pulumi:"serviceLevel"`
	SubnetId          pulumi.StringPtrOutput                        `pulumi:"subnetId"`
	Tags              pulumi.AnyOutput                              `pulumi:"tags"`
	Type              pulumi.StringOutput                           `pulumi:"type"`
	UsageThreshold    pulumi.Float64PtrOutput                       `pulumi:"usageThreshold"`
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
	if isZero(args.ServiceLevel) {
		args.ServiceLevel = pulumi.String("Premium")
	}
	if isZero(args.UsageThreshold) {
		args.UsageThreshold = pulumi.Float64Ptr(107374182400.0)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp:Volume"),
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
	err := ctx.RegisterResource("azure-native:netapp/v20170815:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:netapp/v20170815:Volume", name, id, state, &resource, opts...)
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
	AccountName       string                        `pulumi:"accountName"`
	CreationToken     string                        `pulumi:"creationToken"`
	ExportPolicy      *VolumePropertiesExportPolicy `pulumi:"exportPolicy"`
	Location          *string                       `pulumi:"location"`
	PoolName          string                        `pulumi:"poolName"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	ServiceLevel      string                        `pulumi:"serviceLevel"`
	SubnetId          *string                       `pulumi:"subnetId"`
	Tags              interface{}                   `pulumi:"tags"`
	UsageThreshold    *float64                      `pulumi:"usageThreshold"`
	VolumeName        *string                       `pulumi:"volumeName"`
}


type VolumeArgs struct {
	AccountName       pulumi.StringInput
	CreationToken     pulumi.StringInput
	ExportPolicy      VolumePropertiesExportPolicyPtrInput
	Location          pulumi.StringPtrInput
	PoolName          pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ServiceLevel      pulumi.StringInput
	SubnetId          pulumi.StringPtrInput
	Tags              pulumi.Input
	UsageThreshold    pulumi.Float64PtrInput
	VolumeName        pulumi.StringPtrInput
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

func (o VolumeOutput) CreationToken() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.CreationToken }).(pulumi.StringOutput)
}

func (o VolumeOutput) ExportPolicy() VolumePropertiesResponseExportPolicyPtrOutput {
	return o.ApplyT(func(v *Volume) VolumePropertiesResponseExportPolicyPtrOutput { return v.ExportPolicy }).(VolumePropertiesResponseExportPolicyPtrOutput)
}

func (o VolumeOutput) FileSystemId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.FileSystemId }).(pulumi.StringOutput)
}

func (o VolumeOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeOutput) ServiceLevel() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.ServiceLevel }).(pulumi.StringOutput)
}

func (o VolumeOutput) SubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringPtrOutput { return v.SubnetId }).(pulumi.StringPtrOutput)
}

func (o VolumeOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Volume) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeOutput) UsageThreshold() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.UsageThreshold }).(pulumi.Float64PtrOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
