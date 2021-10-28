


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SkusNestedResourceTypeSecond struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewSkusNestedResourceTypeSecond(ctx *pulumi.Context,
	name string, args *SkusNestedResourceTypeSecondArgs, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeSecond, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NestedResourceTypeFirst == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeFirst'")
	}
	if args.NestedResourceTypeSecond == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeSecond'")
	}
	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210601preview:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-native:providerhub:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20201120:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210501preview:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:SkusNestedResourceTypeSecond"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210901preview:SkusNestedResourceTypeSecond"),
		},
	})
	opts = append(opts, aliases)
	var resource SkusNestedResourceTypeSecond
	err := ctx.RegisterResource("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeSecond", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSkusNestedResourceTypeSecond(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusNestedResourceTypeSecondState, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeSecond, error) {
	var resource SkusNestedResourceTypeSecond
	err := ctx.ReadResource("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeSecond", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type skusNestedResourceTypeSecondState struct {
}

type SkusNestedResourceTypeSecondState struct {
}

func (SkusNestedResourceTypeSecondState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeSecondState)(nil)).Elem()
}

type skusNestedResourceTypeSecondArgs struct {
	NestedResourceTypeFirst  string                 `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond string                 `pulumi:"nestedResourceTypeSecond"`
	Properties               *SkuResourceProperties `pulumi:"properties"`
	ProviderNamespace        string                 `pulumi:"providerNamespace"`
	ResourceType             string                 `pulumi:"resourceType"`
	Sku                      *string                `pulumi:"sku"`
}


type SkusNestedResourceTypeSecondArgs struct {
	NestedResourceTypeFirst  pulumi.StringInput
	NestedResourceTypeSecond pulumi.StringInput
	Properties               SkuResourcePropertiesPtrInput
	ProviderNamespace        pulumi.StringInput
	ResourceType             pulumi.StringInput
	Sku                      pulumi.StringPtrInput
}

func (SkusNestedResourceTypeSecondArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeSecondArgs)(nil)).Elem()
}

type SkusNestedResourceTypeSecondInput interface {
	pulumi.Input

	ToSkusNestedResourceTypeSecondOutput() SkusNestedResourceTypeSecondOutput
	ToSkusNestedResourceTypeSecondOutputWithContext(ctx context.Context) SkusNestedResourceTypeSecondOutput
}

func (*SkusNestedResourceTypeSecond) ElementType() reflect.Type {
	return reflect.TypeOf((*SkusNestedResourceTypeSecond)(nil))
}

func (i *SkusNestedResourceTypeSecond) ToSkusNestedResourceTypeSecondOutput() SkusNestedResourceTypeSecondOutput {
	return i.ToSkusNestedResourceTypeSecondOutputWithContext(context.Background())
}

func (i *SkusNestedResourceTypeSecond) ToSkusNestedResourceTypeSecondOutputWithContext(ctx context.Context) SkusNestedResourceTypeSecondOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusNestedResourceTypeSecondOutput)
}

type SkusNestedResourceTypeSecondOutput struct{ *pulumi.OutputState }

func (SkusNestedResourceTypeSecondOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkusNestedResourceTypeSecond)(nil))
}

func (o SkusNestedResourceTypeSecondOutput) ToSkusNestedResourceTypeSecondOutput() SkusNestedResourceTypeSecondOutput {
	return o
}

func (o SkusNestedResourceTypeSecondOutput) ToSkusNestedResourceTypeSecondOutputWithContext(ctx context.Context) SkusNestedResourceTypeSecondOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SkusNestedResourceTypeSecondOutput{})
}
