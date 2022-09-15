


package enterpriseknowledgegraph

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EnterpriseKnowledgeGraph struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                           `pulumi:"location"`
	Name       pulumi.StringOutput                              `pulumi:"name"`
	Properties EnterpriseKnowledgeGraphPropertiesResponseOutput `pulumi:"properties"`
	Sku        SkuResponsePtrOutput                             `pulumi:"sku"`
	Tags       pulumi.StringMapOutput                           `pulumi:"tags"`
	Type       pulumi.StringOutput                              `pulumi:"type"`
}


func NewEnterpriseKnowledgeGraph(ctx *pulumi.Context,
	name string, args *EnterpriseKnowledgeGraphArgs, opts ...pulumi.ResourceOption) (*EnterpriseKnowledgeGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:enterpriseknowledgegraph/v20181203:EnterpriseKnowledgeGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource EnterpriseKnowledgeGraph
	err := ctx.RegisterResource("azure-native:enterpriseknowledgegraph:EnterpriseKnowledgeGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetEnterpriseKnowledgeGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *EnterpriseKnowledgeGraphState, opts ...pulumi.ResourceOption) (*EnterpriseKnowledgeGraph, error) {
	var resource EnterpriseKnowledgeGraph
	err := ctx.ReadResource("azure-native:enterpriseknowledgegraph:EnterpriseKnowledgeGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type enterpriseKnowledgeGraphState struct {
}

type EnterpriseKnowledgeGraphState struct {
}

func (EnterpriseKnowledgeGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseKnowledgeGraphState)(nil)).Elem()
}

type enterpriseKnowledgeGraphArgs struct {
	Location          *string                             `pulumi:"location"`
	Properties        *EnterpriseKnowledgeGraphProperties `pulumi:"properties"`
	ResourceGroupName string                              `pulumi:"resourceGroupName"`
	ResourceName      *string                             `pulumi:"resourceName"`
	Sku               *Sku                                `pulumi:"sku"`
	Tags              map[string]string                   `pulumi:"tags"`
}


type EnterpriseKnowledgeGraphArgs struct {
	Location          pulumi.StringPtrInput
	Properties        EnterpriseKnowledgeGraphPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Sku               SkuPtrInput
	Tags              pulumi.StringMapInput
}

func (EnterpriseKnowledgeGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*enterpriseKnowledgeGraphArgs)(nil)).Elem()
}

type EnterpriseKnowledgeGraphInput interface {
	pulumi.Input

	ToEnterpriseKnowledgeGraphOutput() EnterpriseKnowledgeGraphOutput
	ToEnterpriseKnowledgeGraphOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphOutput
}

func (*EnterpriseKnowledgeGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraph)(nil)).Elem()
}

func (i *EnterpriseKnowledgeGraph) ToEnterpriseKnowledgeGraphOutput() EnterpriseKnowledgeGraphOutput {
	return i.ToEnterpriseKnowledgeGraphOutputWithContext(context.Background())
}

func (i *EnterpriseKnowledgeGraph) ToEnterpriseKnowledgeGraphOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseKnowledgeGraphOutput)
}

type EnterpriseKnowledgeGraphOutput struct{ *pulumi.OutputState }

func (EnterpriseKnowledgeGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseKnowledgeGraph)(nil)).Elem()
}

func (o EnterpriseKnowledgeGraphOutput) ToEnterpriseKnowledgeGraphOutput() EnterpriseKnowledgeGraphOutput {
	return o
}

func (o EnterpriseKnowledgeGraphOutput) ToEnterpriseKnowledgeGraphOutputWithContext(ctx context.Context) EnterpriseKnowledgeGraphOutput {
	return o
}

func (o EnterpriseKnowledgeGraphOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o EnterpriseKnowledgeGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o EnterpriseKnowledgeGraphOutput) Properties() EnterpriseKnowledgeGraphPropertiesResponseOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) EnterpriseKnowledgeGraphPropertiesResponseOutput {
		return v.Properties
	}).(EnterpriseKnowledgeGraphPropertiesResponseOutput)
}

func (o EnterpriseKnowledgeGraphOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) SkuResponsePtrOutput { return v.Sku }).(SkuResponsePtrOutput)
}

func (o EnterpriseKnowledgeGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o EnterpriseKnowledgeGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *EnterpriseKnowledgeGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(EnterpriseKnowledgeGraphOutput{})
}
