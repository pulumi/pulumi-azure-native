


package v20180901preview

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
			Type: pulumi.String("azure-native:resourcegraph:GraphQuery"),
		},
		{
			Type: pulumi.String("azure-native:resourcegraph/v20200401preview:GraphQuery"),
		},
	})
	opts = append(opts, aliases)
	var resource GraphQuery
	err := ctx.RegisterResource("azure-native:resourcegraph/v20180901preview:GraphQuery", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGraphQuery(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphQueryState, opts ...pulumi.ResourceOption) (*GraphQuery, error) {
	var resource GraphQuery
	err := ctx.ReadResource("azure-native:resourcegraph/v20180901preview:GraphQuery", name, id, state, &resource, opts...)
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
	Etag              *string           `pulumi:"etag"`
	Location          *string           `pulumi:"location"`
	Query             string            `pulumi:"query"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	ResourceName      *string           `pulumi:"resourceName"`
	Tags              map[string]string `pulumi:"tags"`
}


type GraphQueryArgs struct {
	Description       pulumi.StringPtrInput
	Etag              pulumi.StringPtrInput
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

func init() {
	pulumi.RegisterOutputType(GraphQueryOutput{})
}
