


package v20210601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Peering struct {
	pulumi.CustomResourceState

	Direct            PeeringPropertiesDirectResponsePtrOutput   `pulumi:"direct"`
	Exchange          PeeringPropertiesExchangeResponsePtrOutput `pulumi:"exchange"`
	Kind              pulumi.StringOutput                        `pulumi:"kind"`
	Location          pulumi.StringOutput                        `pulumi:"location"`
	Name              pulumi.StringOutput                        `pulumi:"name"`
	PeeringLocation   pulumi.StringPtrOutput                     `pulumi:"peeringLocation"`
	ProvisioningState pulumi.StringOutput                        `pulumi:"provisioningState"`
	Sku               PeeringSkuResponseOutput                   `pulumi:"sku"`
	Tags              pulumi.StringMapOutput                     `pulumi:"tags"`
	Type              pulumi.StringOutput                        `pulumi:"type"`
}


func NewPeering(ctx *pulumi.Context,
	name string, args *PeeringArgs, opts ...pulumi.ResourceOption) (*Peering, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190801preview:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20190901preview:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:Peering"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:Peering"),
		},
	})
	opts = append(opts, aliases)
	var resource Peering
	err := ctx.RegisterResource("azure-native:peering/v20210601:Peering", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPeering(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PeeringState, opts ...pulumi.ResourceOption) (*Peering, error) {
	var resource Peering
	err := ctx.ReadResource("azure-native:peering/v20210601:Peering", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type peeringState struct {
}

type PeeringState struct {
}

func (PeeringState) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringState)(nil)).Elem()
}

type peeringArgs struct {
	Direct            *PeeringPropertiesDirect   `pulumi:"direct"`
	Exchange          *PeeringPropertiesExchange `pulumi:"exchange"`
	Kind              string                     `pulumi:"kind"`
	Location          *string                    `pulumi:"location"`
	PeeringLocation   *string                    `pulumi:"peeringLocation"`
	PeeringName       *string                    `pulumi:"peeringName"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	Sku               PeeringSku                 `pulumi:"sku"`
	Tags              map[string]string          `pulumi:"tags"`
}


type PeeringArgs struct {
	Direct            PeeringPropertiesDirectPtrInput
	Exchange          PeeringPropertiesExchangePtrInput
	Kind              pulumi.StringInput
	Location          pulumi.StringPtrInput
	PeeringLocation   pulumi.StringPtrInput
	PeeringName       pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               PeeringSkuInput
	Tags              pulumi.StringMapInput
}

func (PeeringArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*peeringArgs)(nil)).Elem()
}

type PeeringInput interface {
	pulumi.Input

	ToPeeringOutput() PeeringOutput
	ToPeeringOutputWithContext(ctx context.Context) PeeringOutput
}

func (*Peering) ElementType() reflect.Type {
	return reflect.TypeOf((**Peering)(nil)).Elem()
}

func (i *Peering) ToPeeringOutput() PeeringOutput {
	return i.ToPeeringOutputWithContext(context.Background())
}

func (i *Peering) ToPeeringOutputWithContext(ctx context.Context) PeeringOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeeringOutput)
}

type PeeringOutput struct{ *pulumi.OutputState }

func (PeeringOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Peering)(nil)).Elem()
}

func (o PeeringOutput) ToPeeringOutput() PeeringOutput {
	return o
}

func (o PeeringOutput) ToPeeringOutputWithContext(ctx context.Context) PeeringOutput {
	return o
}

func (o PeeringOutput) Direct() PeeringPropertiesDirectResponsePtrOutput {
	return o.ApplyT(func(v *Peering) PeeringPropertiesDirectResponsePtrOutput { return v.Direct }).(PeeringPropertiesDirectResponsePtrOutput)
}

func (o PeeringOutput) Exchange() PeeringPropertiesExchangeResponsePtrOutput {
	return o.ApplyT(func(v *Peering) PeeringPropertiesExchangeResponsePtrOutput { return v.Exchange }).(PeeringPropertiesExchangeResponsePtrOutput)
}

func (o PeeringOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o PeeringOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PeeringOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PeeringOutput) PeeringLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringPtrOutput { return v.PeeringLocation }).(pulumi.StringPtrOutput)
}

func (o PeeringOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PeeringOutput) Sku() PeeringSkuResponseOutput {
	return o.ApplyT(func(v *Peering) PeeringSkuResponseOutput { return v.Sku }).(PeeringSkuResponseOutput)
}

func (o PeeringOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PeeringOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Peering) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PeeringOutput{})
}
