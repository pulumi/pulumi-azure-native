


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CapacityReservationGroup struct {
	pulumi.CustomResourceState

	CapacityReservations      SubResourceReadOnlyResponseArrayOutput             `pulumi:"capacityReservations"`
	InstanceView              CapacityReservationGroupInstanceViewResponseOutput `pulumi:"instanceView"`
	Location                  pulumi.StringOutput                                `pulumi:"location"`
	Name                      pulumi.StringOutput                                `pulumi:"name"`
	Tags                      pulumi.StringMapOutput                             `pulumi:"tags"`
	Type                      pulumi.StringOutput                                `pulumi:"type"`
	VirtualMachinesAssociated SubResourceReadOnlyResponseArrayOutput             `pulumi:"virtualMachinesAssociated"`
	Zones                     pulumi.StringArrayOutput                           `pulumi:"zones"`
}


func NewCapacityReservationGroup(ctx *pulumi.Context,
	name string, args *CapacityReservationGroupArgs, opts ...pulumi.ResourceOption) (*CapacityReservationGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:CapacityReservationGroup"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:CapacityReservationGroup"),
		},
	})
	opts = append(opts, aliases)
	var resource CapacityReservationGroup
	err := ctx.RegisterResource("azure-native:compute/v20210701:CapacityReservationGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCapacityReservationGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationGroupState, opts ...pulumi.ResourceOption) (*CapacityReservationGroup, error) {
	var resource CapacityReservationGroup
	err := ctx.ReadResource("azure-native:compute/v20210701:CapacityReservationGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type capacityReservationGroupState struct {
}

type CapacityReservationGroupState struct {
}

func (CapacityReservationGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationGroupState)(nil)).Elem()
}

type capacityReservationGroupArgs struct {
	CapacityReservationGroupName *string           `pulumi:"capacityReservationGroupName"`
	Location                     *string           `pulumi:"location"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Tags                         map[string]string `pulumi:"tags"`
	Zones                        []string          `pulumi:"zones"`
}


type CapacityReservationGroupArgs struct {
	CapacityReservationGroupName pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Tags                         pulumi.StringMapInput
	Zones                        pulumi.StringArrayInput
}

func (CapacityReservationGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationGroupArgs)(nil)).Elem()
}

type CapacityReservationGroupInput interface {
	pulumi.Input

	ToCapacityReservationGroupOutput() CapacityReservationGroupOutput
	ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput
}

func (*CapacityReservationGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationGroup)(nil)).Elem()
}

func (i *CapacityReservationGroup) ToCapacityReservationGroupOutput() CapacityReservationGroupOutput {
	return i.ToCapacityReservationGroupOutputWithContext(context.Background())
}

func (i *CapacityReservationGroup) ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationGroupOutput)
}

type CapacityReservationGroupOutput struct{ *pulumi.OutputState }

func (CapacityReservationGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservationGroup)(nil)).Elem()
}

func (o CapacityReservationGroupOutput) ToCapacityReservationGroupOutput() CapacityReservationGroupOutput {
	return o
}

func (o CapacityReservationGroupOutput) ToCapacityReservationGroupOutputWithContext(ctx context.Context) CapacityReservationGroupOutput {
	return o
}

func (o CapacityReservationGroupOutput) CapacityReservations() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) SubResourceReadOnlyResponseArrayOutput {
		return v.CapacityReservations
	}).(SubResourceReadOnlyResponseArrayOutput)
}

func (o CapacityReservationGroupOutput) InstanceView() CapacityReservationGroupInstanceViewResponseOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) CapacityReservationGroupInstanceViewResponseOutput {
		return v.InstanceView
	}).(CapacityReservationGroupInstanceViewResponseOutput)
}

func (o CapacityReservationGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CapacityReservationGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CapacityReservationGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CapacityReservationGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CapacityReservationGroupOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) SubResourceReadOnlyResponseArrayOutput {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

func (o CapacityReservationGroupOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CapacityReservationGroup) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationGroupOutput{})
}
