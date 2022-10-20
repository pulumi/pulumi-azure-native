


package v20171201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type AvailabilitySet struct {
	pulumi.CustomResourceState

	Location                  pulumi.StringOutput                   `pulumi:"location"`
	Name                      pulumi.StringOutput                   `pulumi:"name"`
	PlatformFaultDomainCount  pulumi.IntPtrOutput                   `pulumi:"platformFaultDomainCount"`
	PlatformUpdateDomainCount pulumi.IntPtrOutput                   `pulumi:"platformUpdateDomainCount"`
	Sku                       SkuResponsePtrOutput                  `pulumi:"sku"`
	Statuses                  InstanceViewStatusResponseArrayOutput `pulumi:"statuses"`
	Tags                      pulumi.StringMapOutput                `pulumi:"tags"`
	Type                      pulumi.StringOutput                   `pulumi:"type"`
	VirtualMachines           SubResourceResponseArrayOutput        `pulumi:"virtualMachines"`
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
			Type: pulumi.String("azure-native:compute:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20150615:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160330:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20160430preview:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20170330:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180401:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20180601:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20181001:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20190701:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20191201:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20200601:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20201201:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210301:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210401:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20210701:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20211101:AvailabilitySet"),
		},
		{
			Type: pulumi.String("azure-native:compute/v20220301:AvailabilitySet"),
		},
	})
	opts = append(opts, aliases)
	var resource AvailabilitySet
	err := ctx.RegisterResource("azure-native:compute/v20171201:AvailabilitySet", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAvailabilitySet(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AvailabilitySetState, opts ...pulumi.ResourceOption) (*AvailabilitySet, error) {
	var resource AvailabilitySet
	err := ctx.ReadResource("azure-native:compute/v20171201:AvailabilitySet", name, id, state, &resource, opts...)
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
	AvailabilitySetName       *string           `pulumi:"availabilitySetName"`
	Location                  *string           `pulumi:"location"`
	PlatformFaultDomainCount  *int              `pulumi:"platformFaultDomainCount"`
	PlatformUpdateDomainCount *int              `pulumi:"platformUpdateDomainCount"`
	ResourceGroupName         string            `pulumi:"resourceGroupName"`
	Sku                       *Sku              `pulumi:"sku"`
	Tags                      map[string]string `pulumi:"tags"`
	VirtualMachines           []SubResource     `pulumi:"virtualMachines"`
}


type AvailabilitySetArgs struct {
	AvailabilitySetName       pulumi.StringPtrInput
	Location                  pulumi.StringPtrInput
	PlatformFaultDomainCount  pulumi.IntPtrInput
	PlatformUpdateDomainCount pulumi.IntPtrInput
	ResourceGroupName         pulumi.StringInput
	Sku                       SkuPtrInput
	Tags                      pulumi.StringMapInput
	VirtualMachines           SubResourceArrayInput
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

func (o AvailabilitySetOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AvailabilitySetOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AvailabilitySetOutput) PlatformFaultDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.IntPtrOutput { return v.PlatformFaultDomainCount }).(pulumi.IntPtrOutput)
}

func (o AvailabilitySetOutput) PlatformUpdateDomainCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.IntPtrOutput { return v.PlatformUpdateDomainCount }).(pulumi.IntPtrOutput)
}

func (o AvailabilitySetOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *AvailabilitySet) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o AvailabilitySetOutput) Statuses() InstanceViewStatusResponseArrayOutput {
	return o.ApplyT(func(v *AvailabilitySet) InstanceViewStatusResponseArrayOutput { return v.Statuses }).(InstanceViewStatusResponseArrayOutput)
}

func (o AvailabilitySetOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AvailabilitySetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AvailabilitySet) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o AvailabilitySetOutput) VirtualMachines() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v *AvailabilitySet) SubResourceResponseArrayOutput { return v.VirtualMachines }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(AvailabilitySetOutput{})
}
