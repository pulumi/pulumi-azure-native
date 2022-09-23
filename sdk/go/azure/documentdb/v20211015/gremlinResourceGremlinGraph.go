


package v20211015

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GremlinResourceGremlinGraph struct {
	pulumi.CustomResourceState

	Location pulumi.StringPtrOutput                             `pulumi:"location"`
	Name     pulumi.StringOutput                                `pulumi:"name"`
	Options  GremlinGraphGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource GremlinGraphGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                             `pulumi:"tags"`
	Type     pulumi.StringOutput                                `pulumi:"type"`
}


func NewGremlinResourceGremlinGraph(ctx *pulumi.Context,
	name string, args *GremlinResourceGremlinGraphArgs, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.DatabaseName == nil {
		return nil, errors.New("invalid value for required argument 'DatabaseName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	args.Resource = args.Resource.ToGremlinGraphResourceOutput().ApplyT(func(v GremlinGraphResource) GremlinGraphResource { return *v.Defaults() }).(GremlinGraphResourceOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150401:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20150408:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20151106:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160319:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20160331:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20190801:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20191212:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200301:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200401:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200601preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20200901:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210115:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210301preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210315:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210401preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210415:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210515:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210615:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211015preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815:GremlinResourceGremlinGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:GremlinResourceGremlinGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource GremlinResourceGremlinGraph
	err := ctx.RegisterResource("azure-native:documentdb/v20211015:GremlinResourceGremlinGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGremlinResourceGremlinGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GremlinResourceGremlinGraphState, opts ...pulumi.ResourceOption) (*GremlinResourceGremlinGraph, error) {
	var resource GremlinResourceGremlinGraph
	err := ctx.ReadResource("azure-native:documentdb/v20211015:GremlinResourceGremlinGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type gremlinResourceGremlinGraphState struct {
}

type GremlinResourceGremlinGraphState struct {
}

func (GremlinResourceGremlinGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinGraphState)(nil)).Elem()
}

type gremlinResourceGremlinGraphArgs struct {
	AccountName       string               `pulumi:"accountName"`
	DatabaseName      string               `pulumi:"databaseName"`
	GraphName         *string              `pulumi:"graphName"`
	Location          *string              `pulumi:"location"`
	Options           *CreateUpdateOptions `pulumi:"options"`
	Resource          GremlinGraphResource `pulumi:"resource"`
	ResourceGroupName string               `pulumi:"resourceGroupName"`
	Tags              map[string]string    `pulumi:"tags"`
}


type GremlinResourceGremlinGraphArgs struct {
	AccountName       pulumi.StringInput
	DatabaseName      pulumi.StringInput
	GraphName         pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          GremlinGraphResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (GremlinResourceGremlinGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*gremlinResourceGremlinGraphArgs)(nil)).Elem()
}

type GremlinResourceGremlinGraphInput interface {
	pulumi.Input

	ToGremlinResourceGremlinGraphOutput() GremlinResourceGremlinGraphOutput
	ToGremlinResourceGremlinGraphOutputWithContext(ctx context.Context) GremlinResourceGremlinGraphOutput
}

func (*GremlinResourceGremlinGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinResourceGremlinGraph)(nil)).Elem()
}

func (i *GremlinResourceGremlinGraph) ToGremlinResourceGremlinGraphOutput() GremlinResourceGremlinGraphOutput {
	return i.ToGremlinResourceGremlinGraphOutputWithContext(context.Background())
}

func (i *GremlinResourceGremlinGraph) ToGremlinResourceGremlinGraphOutputWithContext(ctx context.Context) GremlinResourceGremlinGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinResourceGremlinGraphOutput)
}

type GremlinResourceGremlinGraphOutput struct{ *pulumi.OutputState }

func (GremlinResourceGremlinGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinResourceGremlinGraph)(nil)).Elem()
}

func (o GremlinResourceGremlinGraphOutput) ToGremlinResourceGremlinGraphOutput() GremlinResourceGremlinGraphOutput {
	return o
}

func (o GremlinResourceGremlinGraphOutput) ToGremlinResourceGremlinGraphOutputWithContext(ctx context.Context) GremlinResourceGremlinGraphOutput {
	return o
}

func (o GremlinResourceGremlinGraphOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GremlinResourceGremlinGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GremlinResourceGremlinGraphOutput) Options() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
		return v.Options
	}).(GremlinGraphGetPropertiesResponseOptionsPtrOutput)
}

func (o GremlinResourceGremlinGraphOutput) Resource() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) GremlinGraphGetPropertiesResponseResourcePtrOutput {
		return v.Resource
	}).(GremlinGraphGetPropertiesResponseResourcePtrOutput)
}

func (o GremlinResourceGremlinGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GremlinResourceGremlinGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GremlinResourceGremlinGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GremlinResourceGremlinGraphOutput{})
}
