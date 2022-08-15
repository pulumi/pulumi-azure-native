


package v20180401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type ProximityPlacementGroup struct {
	pulumi.CustomResourceState

	AvailabilitySets            SubResourceResponseArrayOutput `pulumi:"availabilitySets"`
	Location                    pulumi.StringOutput            `pulumi:"location"`
	Name                        pulumi.StringOutput            `pulumi:"name"`
	ProximityPlacementGroupType pulumi.StringPtrOutput         `pulumi:"proximityPlacementGroupType"`
	Tags                        pulumi.StringMapOutput         `pulumi:"tags"`
	Type                        pulumi.StringOutput            `pulumi:"type"`
	VirtualMachineScaleSets     SubResourceResponseArrayOutput `pulumi:"virtualMachineScaleSets"`
	VirtualMachines             SubResourceResponseArrayOutput `pulumi:"virtualMachines"`
}


func NewProximityPlacementGroup(ctx *pulumi.Context,
	name string, args *ProximityPlacementGroupArgs, opts ...pulumi.ResourceOption) (*ProximityPlacementGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:ProximityPlacementGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:ProximityPlacementGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource ProximityPlacementGroup
	err := ctx.RegisterResource("azure-native:compute/v20180401:ProximityPlacementGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetProximityPlacementGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ProximityPlacementGroupState, opts ...pulumi.ResourceOption) (*ProximityPlacementGroup, error) {
	var resource ProximityPlacementGroup
	err := ctx.ReadResource("azure-native:compute/v20180401:ProximityPlacementGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type proximityPlacementGroupState struct {
}

type ProximityPlacementGroupState struct {
}

func (ProximityPlacementGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*proximityPlacementGroupState)(nil)).Elem()
}

type proximityPlacementGroupArgs struct {
	Location                    *string           `pulumi:"location"`
	ProximityPlacementGroupName *string           `pulumi:"proximityPlacementGroupName"`
	ProximityPlacementGroupType *string           `pulumi:"proximityPlacementGroupType"`
	ResourceGroupName           string            `pulumi:"resourceGroupName"`
	Tags                        map[string]string `pulumi:"tags"`
}


type ProximityPlacementGroupArgs struct {
	Location                    pulumi.StringPtrInput
	ProximityPlacementGroupName pulumi.StringPtrInput
	ProximityPlacementGroupType pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	Tags                        pulumi.StringMapInput
}

func (ProximityPlacementGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*proximityPlacementGroupArgs)(nil)).Elem()
}

type ProximityPlacementGroupInput interface {
	pulumi.Input

	ToProximityPlacementGroupOutput() ProximityPlacementGroupOutput
	ToProximityPlacementGroupOutputWithContext(ctx context.Context) ProximityPlacementGroupOutput
}

func (*ProximityPlacementGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroup)(nil)).Elem()
}

func (i *ProximityPlacementGroup) ToProximityPlacementGroupOutput() ProximityPlacementGroupOutput {
	return i.ToProximityPlacementGroupOutputWithContext(context.Background())
}

func (i *ProximityPlacementGroup) ToProximityPlacementGroupOutputWithContext(ctx context.Context) ProximityPlacementGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProximityPlacementGroupOutput)
}

type ProximityPlacementGroupOutput struct{ *pulumi.OutputState }

func (ProximityPlacementGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProximityPlacementGroup)(nil)).Elem()
}

func (o ProximityPlacementGroupOutput) ToProximityPlacementGroupOutput() ProximityPlacementGroupOutput {
	return o
}

func (o ProximityPlacementGroupOutput) ToProximityPlacementGroupOutputWithContext(ctx context.Context) ProximityPlacementGroupOutput {
	return o
}

func (o ProximityPlacementGroupOutput) AvailabilitySets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) SubResourceResponseArrayOutput { return v.AvailabilitySets }).(SubResourceResponseArrayOutput)
}

func (o ProximityPlacementGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o ProximityPlacementGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ProximityPlacementGroupOutput) ProximityPlacementGroupType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringPtrOutput { return v.ProximityPlacementGroupType }).(pulumi.StringPtrOutput)
}

func (o ProximityPlacementGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ProximityPlacementGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ProximityPlacementGroupOutput) VirtualMachineScaleSets() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) SubResourceResponseArrayOutput { return v.VirtualMachineScaleSets }).(SubResourceResponseArrayOutput)
}

func (o ProximityPlacementGroupOutput) VirtualMachines() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *ProximityPlacementGroup) SubResourceResponseArrayOutput { return v.VirtualMachines }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ProximityPlacementGroupOutput{})
}
