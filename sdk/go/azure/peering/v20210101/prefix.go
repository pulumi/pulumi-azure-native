


package v20210101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Prefix struct {
	pulumi.CustomResourceState

	ErrorMessage            pulumi.StringOutput                          `pulumi:"errorMessage"`
	Events                  PeeringServicePrefixEventResponseArrayOutput `pulumi:"events"`
	LearnedType             pulumi.StringOutput                          `pulumi:"learnedType"`
	Name                    pulumi.StringOutput                          `pulumi:"name"`
	PeeringServicePrefixKey pulumi.StringPtrOutput                       `pulumi:"peeringServicePrefixKey"`
	Prefix                  pulumi.StringPtrOutput                       `pulumi:"prefix"`
	PrefixValidationState   pulumi.StringOutput                          `pulumi:"prefixValidationState"`
	ProvisioningState       pulumi.StringOutput                          `pulumi:"provisioningState"`
	Type                    pulumi.StringOutput                          `pulumi:"type"`
}


func NewPrefix(ctx *pulumi.Context,
	name string, args *PrefixArgs, opts ...pulumi.ResourceOption) (*Prefix, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringServiceName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringServiceName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:Prefix"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20221001:Prefix"),
		},
	})
	opts = append(opts, aliases)
	var resource Prefix
	err := ctx.RegisterResource("azure-native:peering/v20210101:Prefix", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrefix(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrefixState, opts ...pulumi.ResourceOption) (*Prefix, error) {
	var resource Prefix
	err := ctx.ReadResource("azure-native:peering/v20210101:Prefix", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type prefixState struct {
}

type PrefixState struct {
}

func (PrefixState) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixState)(nil)).Elem()
}

type prefixArgs struct {
	PeeringServiceName      string  `pulumi:"peeringServiceName"`
	PeeringServicePrefixKey *string `pulumi:"peeringServicePrefixKey"`
	Prefix                  *string `pulumi:"prefix"`
	PrefixName              *string `pulumi:"prefixName"`
	ResourceGroupName       string  `pulumi:"resourceGroupName"`
}


type PrefixArgs struct {
	PeeringServiceName      pulumi.StringInput
	PeeringServicePrefixKey pulumi.StringPtrInput
	Prefix                  pulumi.StringPtrInput
	PrefixName              pulumi.StringPtrInput
	ResourceGroupName       pulumi.StringInput
}

func (PrefixArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prefixArgs)(nil)).Elem()
}

type PrefixInput interface {
	pulumi.Input

	ToPrefixOutput() PrefixOutput
	ToPrefixOutputWithContext(ctx context.Context) PrefixOutput
}

func (*Prefix) ElementType() reflect.Type {
	return reflect.TypeOf((**Prefix)(nil)).Elem()
}

func (i *Prefix) ToPrefixOutput() PrefixOutput {
	return i.ToPrefixOutputWithContext(context.Background())
}

func (i *Prefix) ToPrefixOutputWithContext(ctx context.Context) PrefixOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrefixOutput)
}

type PrefixOutput struct{ *pulumi.OutputState }

func (PrefixOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Prefix)(nil)).Elem()
}

func (o PrefixOutput) ToPrefixOutput() PrefixOutput {
	return o
}

func (o PrefixOutput) ToPrefixOutputWithContext(ctx context.Context) PrefixOutput {
	return o
}

func (o PrefixOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringOutput { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o PrefixOutput) Events() PeeringServicePrefixEventResponseArrayOutput {
	return o.ApplyT(func(v *Prefix) PeeringServicePrefixEventResponseArrayOutput { return v.Events }).(PeeringServicePrefixEventResponseArrayOutput)
}

func (o PrefixOutput) LearnedType() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringOutput { return v.LearnedType }).(pulumi.StringOutput)
}

func (o PrefixOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrefixOutput) PeeringServicePrefixKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringPtrOutput { return v.PeeringServicePrefixKey }).(pulumi.StringPtrOutput)
}

func (o PrefixOutput) Prefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringPtrOutput { return v.Prefix }).(pulumi.StringPtrOutput)
}

func (o PrefixOutput) PrefixValidationState() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringOutput { return v.PrefixValidationState }).(pulumi.StringOutput)
}

func (o PrefixOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrefixOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Prefix) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrefixOutput{})
}
