


package peering

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type RegisteredAsn struct {
	pulumi.CustomResourceState

	Asn                     pulumi.IntPtrOutput `pulumi:"asn"`
	Name                    pulumi.StringOutput `pulumi:"name"`
	PeeringServicePrefixKey pulumi.StringOutput `pulumi:"peeringServicePrefixKey"`
	ProvisioningState       pulumi.StringOutput `pulumi:"provisioningState"`
	Type                    pulumi.StringOutput `pulumi:"type"`
}


func NewRegisteredAsn(ctx *pulumi.Context,
	name string, args *RegisteredAsnArgs, opts ...pulumi.ResourceOption) (*RegisteredAsn, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PeeringName == nil {
		return nil, errors.New("invalid value for required argument 'PeeringName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:peering/v20200101preview:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20200401:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20201001:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210101:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20210601:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220101:RegisteredAsn"),
		},
		{
			Type: pulumi.String("azure-native:peering/v20220601:RegisteredAsn"),
		},
	})
	opts = append(opts, aliases)
	var resource RegisteredAsn
	err := ctx.RegisterResource("azure-native:peering:RegisteredAsn", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetRegisteredAsn(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *RegisteredAsnState, opts ...pulumi.ResourceOption) (*RegisteredAsn, error) {
	var resource RegisteredAsn
	err := ctx.ReadResource("azure-native:peering:RegisteredAsn", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type registeredAsnState struct {
}

type RegisteredAsnState struct {
}

func (RegisteredAsnState) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredAsnState)(nil)).Elem()
}

type registeredAsnArgs struct {
	Asn               *int    `pulumi:"asn"`
	PeeringName       string  `pulumi:"peeringName"`
	RegisteredAsnName *string `pulumi:"registeredAsnName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type RegisteredAsnArgs struct {
	Asn               pulumi.IntPtrInput
	PeeringName       pulumi.StringInput
	RegisteredAsnName pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
}

func (RegisteredAsnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*registeredAsnArgs)(nil)).Elem()
}

type RegisteredAsnInput interface {
	pulumi.Input

	ToRegisteredAsnOutput() RegisteredAsnOutput
	ToRegisteredAsnOutputWithContext(ctx context.Context) RegisteredAsnOutput
}

func (*RegisteredAsn) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredAsn)(nil)).Elem()
}

func (i *RegisteredAsn) ToRegisteredAsnOutput() RegisteredAsnOutput {
	return i.ToRegisteredAsnOutputWithContext(context.Background())
}

func (i *RegisteredAsn) ToRegisteredAsnOutputWithContext(ctx context.Context) RegisteredAsnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegisteredAsnOutput)
}

type RegisteredAsnOutput struct{ *pulumi.OutputState }

func (RegisteredAsnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RegisteredAsn)(nil)).Elem()
}

func (o RegisteredAsnOutput) ToRegisteredAsnOutput() RegisteredAsnOutput {
	return o
}

func (o RegisteredAsnOutput) ToRegisteredAsnOutputWithContext(ctx context.Context) RegisteredAsnOutput {
	return o
}

func (o RegisteredAsnOutput) Asn() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RegisteredAsn) pulumi.IntPtrOutput { return v.Asn }).(pulumi.IntPtrOutput)
}

func (o RegisteredAsnOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredAsn) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o RegisteredAsnOutput) PeeringServicePrefixKey() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredAsn) pulumi.StringOutput { return v.PeeringServicePrefixKey }).(pulumi.StringOutput)
}

func (o RegisteredAsnOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredAsn) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RegisteredAsnOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *RegisteredAsn) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(RegisteredAsnOutput{})
}
