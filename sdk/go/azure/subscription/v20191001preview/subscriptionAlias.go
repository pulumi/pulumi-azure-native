


package v20191001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriptionAlias struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                      `pulumi:"name"`
	Properties PutAliasResponsePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                      `pulumi:"type"`
}


func NewSubscriptionAlias(ctx *pulumi.Context,
	name string, args *SubscriptionAliasArgs, opts ...pulumi.ResourceOption) (*SubscriptionAlias, error) {
	if args == nil {
		args = &SubscriptionAliasArgs{}
	}

	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:subscription/v20191001preview:SubscriptionAlias"),
		},
		{
			Type: pulumi.String("azure-native:subscription:SubscriptionAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:subscription:SubscriptionAlias"),
		},
		{
			Type: pulumi.String("azure-native:subscription/v20200901:SubscriptionAlias"),
		},
		{
			Type: pulumi.String("azure-nextgen:subscription/v20200901:SubscriptionAlias"),
		},
	})
	opts = append(opts, aliases)
	var resource SubscriptionAlias
	err := ctx.RegisterResource("azure-native:subscription/v20191001preview:SubscriptionAlias", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSubscriptionAlias(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SubscriptionAliasState, opts ...pulumi.ResourceOption) (*SubscriptionAlias, error) {
	var resource SubscriptionAlias
	err := ctx.ReadResource("azure-native:subscription/v20191001preview:SubscriptionAlias", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type subscriptionAliasState struct {
}

type SubscriptionAliasState struct {
}

func (SubscriptionAliasState) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionAliasState)(nil)).Elem()
}

type subscriptionAliasArgs struct {
	AliasName  *string                    `pulumi:"aliasName"`
	Properties *PutAliasRequestProperties `pulumi:"properties"`
}


type SubscriptionAliasArgs struct {
	AliasName  pulumi.StringPtrInput
	Properties PutAliasRequestPropertiesPtrInput
}

func (SubscriptionAliasArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*subscriptionAliasArgs)(nil)).Elem()
}

type SubscriptionAliasInput interface {
	pulumi.Input

	ToSubscriptionAliasOutput() SubscriptionAliasOutput
	ToSubscriptionAliasOutputWithContext(ctx context.Context) SubscriptionAliasOutput
}

func (*SubscriptionAlias) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionAlias)(nil))
}

func (i *SubscriptionAlias) ToSubscriptionAliasOutput() SubscriptionAliasOutput {
	return i.ToSubscriptionAliasOutputWithContext(context.Background())
}

func (i *SubscriptionAlias) ToSubscriptionAliasOutputWithContext(ctx context.Context) SubscriptionAliasOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionAliasOutput)
}

type SubscriptionAliasOutput struct{ *pulumi.OutputState }

func (SubscriptionAliasOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionAlias)(nil))
}

func (o SubscriptionAliasOutput) ToSubscriptionAliasOutput() SubscriptionAliasOutput {
	return o
}

func (o SubscriptionAliasOutput) ToSubscriptionAliasOutputWithContext(ctx context.Context) SubscriptionAliasOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SubscriptionAliasInput)(nil)).Elem(), &SubscriptionAlias{})
	pulumi.RegisterOutputType(SubscriptionAliasOutput{})
}
