


package v20210601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SkusNestedResourceTypeThird struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties SkuResourceResponsePropertiesOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewSkusNestedResourceTypeThird(ctx *pulumi.Context,
	name string, args *SkusNestedResourceTypeThirdArgs, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeThird, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.NestedResourceTypeFirst == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeFirst'")
	}
	if args.NestedResourceTypeSecond == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeSecond'")
	}
	if args.NestedResourceTypeThird == nil {
		return nil, errors.New("invalid value for required argument 'NestedResourceTypeThird'")
	}
	if args.ProviderNamespace == nil {
		return nil, errors.New("invalid value for required argument 'ProviderNamespace'")
	}
	if args.ResourceType == nil {
		return nil, errors.New("invalid value for required argument 'ResourceType'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210601preview:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20201120:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20201120:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210501preview:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210501preview:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-native:providerhub/v20210901preview:SkusNestedResourceTypeThird"),
		},
		{
			Type: pulumi.String("azure-nextgen:providerhub/v20210901preview:SkusNestedResourceTypeThird"),
		},
	})
	opts = append(opts, aliases)
	var resource SkusNestedResourceTypeThird
	err := ctx.RegisterResource("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeThird", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSkusNestedResourceTypeThird(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SkusNestedResourceTypeThirdState, opts ...pulumi.ResourceOption) (*SkusNestedResourceTypeThird, error) {
	var resource SkusNestedResourceTypeThird
	err := ctx.ReadResource("azure-native:providerhub/v20210601preview:SkusNestedResourceTypeThird", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type skusNestedResourceTypeThirdState struct {
}

type SkusNestedResourceTypeThirdState struct {
}

func (SkusNestedResourceTypeThirdState) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeThirdState)(nil)).Elem()
}

type skusNestedResourceTypeThirdArgs struct {
	NestedResourceTypeFirst  string                 `pulumi:"nestedResourceTypeFirst"`
	NestedResourceTypeSecond string                 `pulumi:"nestedResourceTypeSecond"`
	NestedResourceTypeThird  string                 `pulumi:"nestedResourceTypeThird"`
	Properties               *SkuResourceProperties `pulumi:"properties"`
	ProviderNamespace        string                 `pulumi:"providerNamespace"`
	ResourceType             string                 `pulumi:"resourceType"`
	Sku                      *string                `pulumi:"sku"`
}


type SkusNestedResourceTypeThirdArgs struct {
	NestedResourceTypeFirst  pulumi.StringInput
	NestedResourceTypeSecond pulumi.StringInput
	NestedResourceTypeThird  pulumi.StringInput
	Properties               SkuResourcePropertiesPtrInput
	ProviderNamespace        pulumi.StringInput
	ResourceType             pulumi.StringInput
	Sku                      pulumi.StringPtrInput
}

func (SkusNestedResourceTypeThirdArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*skusNestedResourceTypeThirdArgs)(nil)).Elem()
}

type SkusNestedResourceTypeThirdInput interface {
	pulumi.Input

	ToSkusNestedResourceTypeThirdOutput() SkusNestedResourceTypeThirdOutput
	ToSkusNestedResourceTypeThirdOutputWithContext(ctx context.Context) SkusNestedResourceTypeThirdOutput
}

func (*SkusNestedResourceTypeThird) ElementType() reflect.Type {
	return reflect.TypeOf((*SkusNestedResourceTypeThird)(nil))
}

func (i *SkusNestedResourceTypeThird) ToSkusNestedResourceTypeThirdOutput() SkusNestedResourceTypeThirdOutput {
	return i.ToSkusNestedResourceTypeThirdOutputWithContext(context.Background())
}

func (i *SkusNestedResourceTypeThird) ToSkusNestedResourceTypeThirdOutputWithContext(ctx context.Context) SkusNestedResourceTypeThirdOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkusNestedResourceTypeThirdOutput)
}

type SkusNestedResourceTypeThirdOutput struct{ *pulumi.OutputState }

func (SkusNestedResourceTypeThirdOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkusNestedResourceTypeThird)(nil))
}

func (o SkusNestedResourceTypeThirdOutput) ToSkusNestedResourceTypeThirdOutput() SkusNestedResourceTypeThirdOutput {
	return o
}

func (o SkusNestedResourceTypeThirdOutput) ToSkusNestedResourceTypeThirdOutputWithContext(ctx context.Context) SkusNestedResourceTypeThirdOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SkusNestedResourceTypeThirdOutput{})
}
