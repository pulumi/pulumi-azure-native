


package v20211120preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VolumeGroup struct {
	pulumi.CustomResourceState

	Encryption        pulumi.StringPtrOutput          `pulumi:"encryption"`
	Name              pulumi.StringOutput             `pulumi:"name"`
	NetworkAcls       NetworkRuleSetResponsePtrOutput `pulumi:"networkAcls"`
	ProtocolType      pulumi.StringPtrOutput          `pulumi:"protocolType"`
	ProvisioningState pulumi.StringOutput             `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput        `pulumi:"systemData"`
	Tags              pulumi.StringMapOutput          `pulumi:"tags"`
	Type              pulumi.StringOutput             `pulumi:"type"`
}


func NewVolumeGroup(ctx *pulumi.Context,
	name string, args *VolumeGroupArgs, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ElasticSanName == nil {
		return nil, errors.New("invalid value for required argument 'ElasticSanName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:elasticsan:VolumeGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource VolumeGroup
	err := ctx.RegisterResource("azure-native:elasticsan/v20211120preview:VolumeGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVolumeGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VolumeGroupState, opts ...pulumi.ResourceOption) (*VolumeGroup, error) {
	var resource VolumeGroup
	err := ctx.ReadResource("azure-native:elasticsan/v20211120preview:VolumeGroup", name, id, state, &resource, opts...)
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
	ElasticSanName    string            `pulumi:"elasticSanName"`
	Encryption        *string           `pulumi:"encryption"`
	NetworkAcls       *NetworkRuleSet   `pulumi:"networkAcls"`
	ProtocolType      *string           `pulumi:"protocolType"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	Tags              map[string]string `pulumi:"tags"`
	VolumeGroupName   *string           `pulumi:"volumeGroupName"`
}


type VolumeGroupArgs struct {
	ElasticSanName    pulumi.StringInput
	Encryption        pulumi.StringPtrInput
	NetworkAcls       NetworkRuleSetPtrInput
	ProtocolType      pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
	VolumeGroupName   pulumi.StringPtrInput
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

func (o VolumeGroupOutput) Encryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringPtrOutput { return v.Encryption }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o VolumeGroupOutput) NetworkAcls() NetworkRuleSetResponsePtrOutput {
	return o.ApplyT(func(v *VolumeGroup) NetworkRuleSetResponsePtrOutput { return v.NetworkAcls }).(NetworkRuleSetResponsePtrOutput)
}

func (o VolumeGroupOutput) ProtocolType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringPtrOutput { return v.ProtocolType }).(pulumi.StringPtrOutput)
}

func (o VolumeGroupOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VolumeGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *VolumeGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o VolumeGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o VolumeGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *VolumeGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(VolumeGroupOutput{})
}
