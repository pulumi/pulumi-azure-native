


package resourcegraph

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GraphQuery struct {
	pulumi.CustomResourceState

	Description  pulumi.StringPtrOutput `pulumi:"description"`
	Etag         pulumi.StringPtrOutput `pulumi:"etag"`
	Location     pulumi.StringPtrOutput `pulumi:"location"`
	Name         pulumi.StringOutput    `pulumi:"name"`
	Query        pulumi.StringOutput    `pulumi:"query"`
	ResultKind   pulumi.StringOutput    `pulumi:"resultKind"`
	Tags         pulumi.StringMapOutput `pulumi:"tags"`
	TimeModified pulumi.StringOutput    `pulumi:"timeModified"`
	Type         pulumi.StringOutput    `pulumi:"type"`
}


func NewGraphQuery(ctx *pulumi.Context,
	name string, args *GraphQueryArgs, opts ...pulumi.ResourceOption) (*GraphQuery, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Query == nil {
		return nil, errors.New("invalid value for required argument 'Query'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:resourcegraph/v20180901preview:GraphQuery"),
		},
		{
			Type: pulumi.String("azure-native:resourcegraph/v20200401preview:GraphQuery"),
		},
	})
	opts = append(opts, aliases)
	var resource GraphQuery
	err := ctx.RegisterResource("azure-native:resourcegraph:GraphQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGraphQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQueryState, opts ...pulumi.ResourceOption) (*GraphQuery, error) {
	var resource GraphQuery
	err := ctx.ReadResource("azure-native:resourcegraph:GraphQuery", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type graphQueryState struct {
}

type GraphQueryState struct {
}

func (GraphQueryState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQueryState)(nil)).Elem()
}

type graphQueryArgs struct {
	Description       *string           `pulumi:"description"`
	Location          *string           `pulumi:"location"`
	Query             string            `pulumi:"query"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Tags              map[string]string `pulumi:"tags"`
}


type GraphQueryArgs struct {
	Description       pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Query             pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ResourceName      pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
}

func (GraphQueryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphQueryArgs)(nil)).Elem()
}

type GraphQueryInput interface {
	pulumi.Input

	ToGraphQueryOutput() GraphQueryOutput
	ToGraphQueryOutputWithContext(ctx context.Context) GraphQueryOutput
}

func (*GraphQuery) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQuery)(nil)).Elem()
}

func (i *GraphQuery) ToGraphQueryOutput() GraphQueryOutput {
	return i.ToGraphQueryOutputWithContext(context.Background())
}

func (i *GraphQuery) ToGraphQueryOutputWithContext(ctx context.Context) GraphQueryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphQueryOutput)
}

type GraphQueryOutput struct{ *pulumi.OutputState }

func (GraphQueryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphQuery)(nil)).Elem()
}

func (o GraphQueryOutput) ToGraphQueryOutput() GraphQueryOutput {
	return o
}

func (o GraphQueryOutput) ToGraphQueryOutputWithContext(ctx context.Context) GraphQueryOutput {
	return o
}

func (o GraphQueryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GraphQueryOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o GraphQueryOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GraphQueryOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GraphQueryOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.Query }).(pulumi.StringOutput)
}

func (o GraphQueryOutput) ResultKind() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.ResultKind }).(pulumi.StringOutput)
}

func (o GraphQueryOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GraphQueryOutput) TimeModified() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.TimeModified }).(pulumi.StringOutput)
}

func (o GraphQueryOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphQuery) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphQueryOutput{})
}
