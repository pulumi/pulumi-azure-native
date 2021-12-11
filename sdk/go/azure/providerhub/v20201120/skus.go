


package v20201120

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Skus struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewSkus(ctx *pulumi.Context,
	name string, args *SkusArgs, opts ...pulumi.ResourceOption) (*Skus, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:providerhub:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210601preview:Skus"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:Skus"),
		},
	})
	opts = append(opts, aliases)
	var resource Skus
	err := ctx.RegisterResource("azure-native:providerhub/v20201120:Skus", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSkus(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusState, opts ...pulumi.ResourceOption) (*Skus, error) {
	var resource Skus
	err := ctx.ReadResource("azure-native:providerhub/v20201120:Skus", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type skusState struct {
}

type SkusState struct {
}

func (SkusState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusState)(nil)).Elem()
}

type skusArgs struct {
	Properties        *SkuResourceProperties `pulumi:"properties"`
	ProviderNamespace string                 `pulumi:"providerNamespace"`
	ResourceType      string                 `pulumi:"resourceType"`
	Sku               *string                `pulumi:"sku"`
}


type SkusArgs struct {
	Properties        SkuResourcePropertiesPtrInput
	ProviderNamespace pulumi.StringInput
	ResourceType      pulumi.StringInput
	Sku               pulumi.StringPtrInput
}

func (SkusArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusArgs)(nil)).Elem()
}

type SkusInput interface {
	pulumi.Input

	ToSkusOutput() SkusOutput
	ToSkusOutputWithContext(ctx context.Context) SkusOutput
}

func (*Skus) ElementType() reflect.Type {
	return reflect.TypeOf((*Skus)(nil))
}

func (i *Skus) ToSkusOutput() SkusOutput {
	return i.ToSkusOutputWithContext(context.Background())
}

func (i *Skus) ToSkusOutputWithContext(ctx context.Context) SkusOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusOutput)
}

type SkusOutput struct{ *pulumi.OutputState }

func (SkusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Skus)(nil))
}

func (o SkusOutput) ToSkusOutput() SkusOutput {
	return o
}

func (o SkusOutput) ToSkusOutputWithContext(ctx context.Context) SkusOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SkusOutput{})
}
