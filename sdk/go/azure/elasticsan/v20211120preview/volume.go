


package v20211120preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Volume struct {
	pulumi.CustomResourceState

	CreationData  SourceCreationDataResponsePtrOutput `pulumi:"creationData"`
	Name          pulumi.StringOutput                 `pulumi:"name"`
	SizeGiB       pulumi.Float64PtrOutput             `pulumi:"sizeGiB"`
	StorageTarget IscsiTargetInfoResponseOutput       `pulumi:"storageTarget"`
	SystemData    SystemDataResponseOutput            `pulumi:"systemData"`
	Tags          pulumi.StringMapOutput              `pulumi:"tags"`
	Type          pulumi.StringOutput                 `pulumi:"type"`
	VolumeId      pulumi.StringOutput                 `pulumi:"volumeId"`
}


func NewVolume(ctx *pulumi.Context,
	name string, args *VolumeArgs, opts ...pulumi.ResourceOption) (*Volume, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElasticSanName == nil {
		return nil, errors.New("invalid value for required argument 'ElasticSanName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VolumeGroupName == nil {
		return nil, errors.New("invalid value for required argument 'VolumeGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elasticsan:Volume"),
		},
	})
	opts = append(opts, aliases)
	var resource Volume
	err := ctx.RegisterResource("azure-native:elasticsan/v20211120preview:Volume", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolume(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeState, opts ...pulumi.ResourceOption) (*Volume, error) {
	var resource Volume
	err := ctx.ReadResource("azure-native:elasticsan/v20211120preview:Volume", name, id, state, &resource, opts...)
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
	CreationData      *SourceCreationData `pulumi:"creationData"`
	ElasticSanName    string              `pulumi:"elasticSanName"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	SizeGiB           *float64            `pulumi:"sizeGiB"`
	Tags              map[string]string   `pulumi:"tags"`
	VolumeGroupName   string              `pulumi:"volumeGroupName"`
	VolumeName        *string             `pulumi:"volumeName"`
}


type VolumeArgs struct {
	CreationData      SourceCreationDataPtrInput
	ElasticSanName    pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SizeGiB           pulumi.Float64PtrInput
	Tags              pulumi.StringMapInput
	VolumeGroupName   pulumi.StringInput
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

func (o VolumeOutput) CreationData() SourceCreationDataResponsePtrOutput {
	return o.ApplyT(func(v *Volume) SourceCreationDataResponsePtrOutput { return v.CreationData }).(SourceCreationDataResponsePtrOutput)
}

func (o VolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeOutput) SizeGiB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *Volume) pulumi.Float64PtrOutput { return v.SizeGiB }).(pulumi.Float64PtrOutput)
}

func (o VolumeOutput) StorageTarget() IscsiTargetInfoResponseOutput {
	return o.ApplyT(func(v *Volume) IscsiTargetInfoResponseOutput { return v.StorageTarget }).(IscsiTargetInfoResponseOutput)
}

func (o VolumeOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Volume) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VolumeOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeOutput) VolumeId() pulumi.StringOutput {
	return o.ApplyT(func(v *Volume) pulumi.StringOutput { return v.VolumeId }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeOutput{})
}
