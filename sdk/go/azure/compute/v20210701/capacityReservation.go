


package v20210701

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CapacityReservation struct {
	pulumi.CustomResourceState

	InstanceView              CapacityReservationInstanceViewResponseOutput `pulumi:"instanceView"`
	Location                  pulumi.StringOutput                           `pulumi:"location"`
	Name                      pulumi.StringOutput                           `pulumi:"name"`
	ProvisioningState         pulumi.StringOutput                           `pulumi:"provisioningState"`
	ProvisioningTime          pulumi.StringOutput                           `pulumi:"provisioningTime"`
	ReservationId             pulumi.StringOutput                           `pulumi:"reservationId"`
	Sku                       SkuResponseOutput                             `pulumi:"sku"`
	Tags                      pulumi.StringMapOutput                        `pulumi:"tags"`
	Type                      pulumi.StringOutput                           `pulumi:"type"`
	VirtualMachinesAssociated SubResourceReadOnlyResponseArrayOutput        `pulumi:"virtualMachinesAssociated"`
	Zones                     pulumi.StringArrayOutput                      `pulumi:"zones"`
}


func NewCapacityReservation(ctx *pulumi.Context,
	name string, args *CapacityReservationArgs, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CapacityReservationGroupName == nil {
		return nil, errors.New("invalid value for required argument 'CapacityReservationGroupName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:compute:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:CapacityReservation"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:CapacityReservation"),
		},
	})
	opts = append(opts, aliases)
	var resource CapacityReservation
	err := ctx.RegisterResource("azure-native:compute/v20210701:CapacityReservation", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCapacityReservation(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CapacityReservationState, opts ...pulumi.ResourceOption) (*CapacityReservation, error) {
	var resource CapacityReservation
	err := ctx.ReadResource("azure-native:compute/v20210701:CapacityReservation", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type capacityReservationState struct {
}

type CapacityReservationState struct {
}

func (CapacityReservationState) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationState)(nil)).Elem()
}

type capacityReservationArgs struct {
	CapacityReservationGroupName string            `pulumi:"capacityReservationGroupName"`
	CapacityReservationName      *string           `pulumi:"capacityReservationName"`
	Location                     *string           `pulumi:"location"`
	ResourceGroupName            string            `pulumi:"resourceGroupName"`
	Sku                          Sku               `pulumi:"sku"`
	Tags                         map[string]string `pulumi:"tags"`
	Zones                        []string          `pulumi:"zones"`
}


type CapacityReservationArgs struct {
	CapacityReservationGroupName pulumi.StringInput
	CapacityReservationName      pulumi.StringPtrInput
	Location                     pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	Sku                          SkuInput
	Tags                         pulumi.StringMapInput
	Zones                        pulumi.StringArrayInput
}

func (CapacityReservationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*capacityReservationArgs)(nil)).Elem()
}

type CapacityReservationInput interface {
	pulumi.Input

	ToCapacityReservationOutput() CapacityReservationOutput
	ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput
}

func (*CapacityReservation) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil)).Elem()
}

func (i *CapacityReservation) ToCapacityReservationOutput() CapacityReservationOutput {
	return i.ToCapacityReservationOutputWithContext(context.Background())
}

func (i *CapacityReservation) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityReservationOutput)
}

type CapacityReservationOutput struct{ *pulumi.OutputState }

func (CapacityReservationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityReservation)(nil)).Elem()
}

func (o CapacityReservationOutput) ToCapacityReservationOutput() CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) ToCapacityReservationOutputWithContext(ctx context.Context) CapacityReservationOutput {
	return o
}

func (o CapacityReservationOutput) InstanceView() CapacityReservationInstanceViewResponseOutput {
	return o.ApplyT(func(v *CapacityReservation) CapacityReservationInstanceViewResponseOutput { return v.InstanceView }).(CapacityReservationInstanceViewResponseOutput)
}

func (o CapacityReservationOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) ProvisioningTime() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.ProvisioningTime }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) ReservationId() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.ReservationId }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *CapacityReservation) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o CapacityReservationOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o CapacityReservationOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o CapacityReservationOutput) VirtualMachinesAssociated() SubResourceReadOnlyResponseArrayOutput {
	return o.ApplyT(func(v *CapacityReservation) SubResourceReadOnlyResponseArrayOutput {
		return v.VirtualMachinesAssociated
	}).(SubResourceReadOnlyResponseArrayOutput)
}

func (o CapacityReservationOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CapacityReservation) pulumi.StringArrayOutput { return v.Zones }).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(CapacityReservationOutput{})
}
