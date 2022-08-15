


package scvmm

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AvailabilitySet struct {
	pulumi.CustomResourceState

	AvailabilitySetName pulumi.StringPtrOutput            `pulumi:"availabilitySetName"`
	ExtendedLocation    ExtendedLocationResponsePtrOutput `pulumi:"extendedLocation"`
	Location            pulumi.StringPtrOutput            `pulumi:"location"`
	Name                pulumi.StringOutput               `pulumi:"name"`
	ProvisioningState   pulumi.StringOutput               `pulumi:"provisioningState"`
	SystemData          SystemDataResponseOutput          `pulumi:"systemData"`
	Tags                pulumi.StringMapOutput            `pulumi:"tags"`
	Type                pulumi.StringOutput               `pulumi:"type"`
	VmmServerId         pulumi.StringPtrOutput            `pulumi:"vmmServerId"`
}


func NewAvailabilitySet(ctx *pulumi.Context,
	name string, args *AvailabilitySetArgs, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:scvmm/v20200605preview:AvailabilitySet"),
		},
	})
	opts = append(opts, aliases)
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure-native:scvmm:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure-native:scvmm:AvailabilitySet", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type availabilitySetState struct {
}

type AvailabilitySetState struct {
}

func (AvailabilitySetState) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetState)(nil)).Elem()
}

type availabilitySetArgs struct {
	AvailabilitySetName *string           `pulumi:"availabilitySetName"`
	ExtendedLocation    *ExtendedLocation `pulumi:"extendedLocation"`
	Location            *string           `pulumi:"location"`
	ResourceGroupName   string            `pulumi:"resourceGroupName"`
	Tags                map[string]string `pulumi:"tags"`
	VmmServerId         *string           `pulumi:"vmmServerId"`
}


type AvailabilitySetArgs struct {
	AvailabilitySetName pulumi.StringPtrInput
	ExtendedLocation    ExtendedLocationPtrInput
	Location            pulumi.StringPtrInput
	ResourceGroupName   pulumi.StringInput
	Tags                pulumi.StringMapInput
	VmmServerId         pulumi.StringPtrInput
}

func (AvailabilitySetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*availabilitySetArgs)(nil)).Elem()
}

type AvailabilitySetInput interface {
	pulumi.Input

	ToAvailabilitySetOutput() AvailabilitySetOutput
	ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput
}

func (*AvailabilitySet) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil)).Elem()
}

func (i *AvailabilitySet) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return i.ToAvailabilitySetOutputWithContext(context.Background())
}

func (i *AvailabilitySet) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilitySetOutput)
}

type AvailabilitySetOutput struct{ *pulumi.OutputState }

func (AvailabilitySetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilitySet)(nil)).Elem()
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutput() AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) ToAvailabilitySetOutputWithContext(ctx context.Context) AvailabilitySetOutput {
	return o
}

func (o AvailabilitySetOutput) AvailabilitySetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringPtrOutput { return v.AvailabilitySetName }).(pulumi.StringPtrOutput)
}

func (o AvailabilitySetOutput) ExtendedLocation() ExtendedLocationResponsePtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) ExtendedLocationResponsePtrOutput { return v.ExtendedLocation }).(ExtendedLocationResponsePtrOutput)
}

func (o AvailabilitySetOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o AvailabilitySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AvailabilitySetOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o AvailabilitySetOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *AvailabilitySet) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o AvailabilitySetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AvailabilitySetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AvailabilitySetOutput) VmmServerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringPtrOutput { return v.VmmServerId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilitySetOutput{})
}
