


package netapp

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VolumeGroup struct {
	pulumi.CustomResourceState

	GroupMetaData     VolumeGroupMetaDataResponsePtrOutput           `pulumi:"groupMetaData"`
	Location          pulumi.StringPtrOutput                         `pulumi:"location"`
	Name              pulumi.StringOutput                            `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                            `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput                         `pulumi:"tags"`
	Type              pulumi.StringOutput                            `pulumi:"type"`
	Volumes           VolumeGroupVolumePropertiesResponseArrayOutput `pulumi:"volumes"`
}


func NewVolumeGroup(ctx *pulumi.Context,
	name string, args *VolumeGroupArgs, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:netapp/v20210801:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20211001:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220101:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220301:VolumeGroup"),
		},
		{
			Type: pulumi.String("azure-native:netapp/v20220501:VolumeGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeGroup
	err := ctx.RegisterResource("azure-native:netapp:VolumeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolumeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeGroupState, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	var resource VolumeGroup
	err := ctx.ReadResource("azure-native:netapp:VolumeGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type volumeGroupState struct {
}

type VolumeGroupState struct {
}

func (VolumeGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeGroupState)(nil)).Elem()
}

type volumeGroupArgs struct {
	AccountName       string                        `pulumi:"accountName"`
	GroupMetaData     *VolumeGroupMetaData          `pulumi:"groupMetaData"`
	Location          *string                       `pulumi:"location"`
	ResourceGroupName string                        `pulumi:"resourceGroupName"`
	Tags              map[string]string             `pulumi:"tags"`
	VolumeGroupName   *string                       `pulumi:"volumeGroupName"`
	Volumes           []VolumeGroupVolumeProperties `pulumi:"volumes"`
}


type VolumeGroupArgs struct {
	AccountName       pulumi.StringInput
	GroupMetaData     VolumeGroupMetaDataPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VolumeGroupName   pulumi.StringPtrInput
	Volumes           VolumeGroupVolumePropertiesArrayInput
}

func (VolumeGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*volumeGroupArgs)(nil)).Elem()
}

type VolumeGroupInput interface {
	pulumi.Input

	ToVolumeGroupOutput() VolumeGroupOutput
	ToVolumeGroupOutputWithContext(ctx context.Context) VolumeGroupOutput
}

func (*VolumeGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroup)(nil)).Elem()
}

func (i *VolumeGroup) ToVolumeGroupOutput() VolumeGroupOutput {
	return i.ToVolumeGroupOutputWithContext(context.Background())
}

func (i *VolumeGroup) ToVolumeGroupOutputWithContext(ctx context.Context) VolumeGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeGroupOutput)
}

type VolumeGroupOutput struct{ *pulumi.OutputState }

func (VolumeGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeGroup)(nil)).Elem()
}

func (o VolumeGroupOutput) ToVolumeGroupOutput() VolumeGroupOutput {
	return o
}

func (o VolumeGroupOutput) ToVolumeGroupOutputWithContext(ctx context.Context) VolumeGroupOutput {
	return o
}

func (o VolumeGroupOutput) GroupMetaData() VolumeGroupMetaDataResponsePtrOutput {
	return o.ApplyT(func(v *VolumeGroup) VolumeGroupMetaDataResponsePtrOutput { return v.GroupMetaData }).(VolumeGroupMetaDataResponsePtrOutput)
}

func (o VolumeGroupOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o VolumeGroupOutput) Volumes() VolumeGroupVolumePropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VolumeGroup) VolumeGroupVolumePropertiesResponseArrayOutput { return v.Volumes }).(VolumeGroupVolumePropertiesResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeGroupOutput{})
}
