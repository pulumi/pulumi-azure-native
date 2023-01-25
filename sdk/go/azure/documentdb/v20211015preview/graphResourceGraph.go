


package v20211015preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GraphResourceGraph struct {
	pulumi.CustomResourceState

	Identity ManagedServiceIdentityResponsePtrOutput             `pulumi:"identity"`
	Location pulumi.StringPtrOutput                              `pulumi:"location"`
	Name     pulumi.StringOutput                                 `pulumi:"name"`
	Options  GraphResourceGetPropertiesResponseOptionsPtrOutput  `pulumi:"options"`
	Resource GraphResourceGetPropertiesResponseResourcePtrOutput `pulumi:"resource"`
	Tags     pulumi.StringMapOutput                              `pulumi:"tags"`
	Type     pulumi.StringOutput                                 `pulumi:"type"`
}


func NewGraphResourceGraph(ctx *pulumi.Context,
	name string, args *GraphResourceGraphArgs, opts ...pulumi.ResourceOption) (*GraphResourceGraph, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Resource == nil {
		return nil, errors.New("invalid value for required argument 'Resource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:documentdb:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20210701preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20211115preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220215preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220515preview:GraphResourceGraph"),
		},
		{
			Type: pulumi.String("azure-native:documentdb/v20220815preview:GraphResourceGraph"),
		},
	})
	opts = append(opts, aliases)
	var resource GraphResourceGraph
	err := ctx.RegisterResource("azure-native:documentdb/v20211015preview:GraphResourceGraph", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGraphResourceGraph(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GraphResourceGraphState, opts ...pulumi.ResourceOption) (*GraphResourceGraph, error) {
	var resource GraphResourceGraph
	err := ctx.ReadResource("azure-native:documentdb/v20211015preview:GraphResourceGraph", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type graphResourceGraphState struct {
}

type GraphResourceGraphState struct {
}

func (GraphResourceGraphState) ElementType() reflect.Type {
	return reflect.TypeOf((*graphResourceGraphState)(nil)).Elem()
}

type graphResourceGraphArgs struct {
	AccountName       string                  `pulumi:"accountName"`
	GraphName         *string                 `pulumi:"graphName"`
	Identity          *ManagedServiceIdentity `pulumi:"identity"`
	Location          *string                 `pulumi:"location"`
	Options           *CreateUpdateOptions    `pulumi:"options"`
	Resource          GraphResource           `pulumi:"resource"`
	ResourceGroupName string                  `pulumi:"resourceGroupName"`
	Tags              map[string]string       `pulumi:"tags"`
}


type GraphResourceGraphArgs struct {
	AccountName       pulumi.StringInput
	GraphName         pulumi.StringPtrInput
	Identity          ManagedServiceIdentityPtrInput
	Location          pulumi.StringPtrInput
	Options           CreateUpdateOptionsPtrInput
	Resource          GraphResourceInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (GraphResourceGraphArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*graphResourceGraphArgs)(nil)).Elem()
}

type GraphResourceGraphInput interface {
	pulumi.Input

	ToGraphResourceGraphOutput() GraphResourceGraphOutput
	ToGraphResourceGraphOutputWithContext(ctx context.Context) GraphResourceGraphOutput
}

func (*GraphResourceGraph) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphResourceGraph)(nil)).Elem()
}

func (i *GraphResourceGraph) ToGraphResourceGraphOutput() GraphResourceGraphOutput {
	return i.ToGraphResourceGraphOutputWithContext(context.Background())
}

func (i *GraphResourceGraph) ToGraphResourceGraphOutputWithContext(ctx context.Context) GraphResourceGraphOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GraphResourceGraphOutput)
}

type GraphResourceGraphOutput struct{ *pulumi.OutputState }

func (GraphResourceGraphOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GraphResourceGraph)(nil)).Elem()
}

func (o GraphResourceGraphOutput) ToGraphResourceGraphOutput() GraphResourceGraphOutput {
	return o
}

func (o GraphResourceGraphOutput) ToGraphResourceGraphOutputWithContext(ctx context.Context) GraphResourceGraphOutput {
	return o
}

func (o GraphResourceGraphOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o GraphResourceGraphOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o GraphResourceGraphOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GraphResourceGraphOutput) Options() GraphResourceGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) GraphResourceGetPropertiesResponseOptionsPtrOutput { return v.Options }).(GraphResourceGetPropertiesResponseOptionsPtrOutput)
}

func (o GraphResourceGraphOutput) Resource() GraphResourceGetPropertiesResponseResourcePtrOutput {
	return o.ApplyT(func(v *GraphResourceGraph) GraphResourceGetPropertiesResponseResourcePtrOutput { return v.Resource }).(GraphResourceGetPropertiesResponseResourcePtrOutput)
}

func (o GraphResourceGraphOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o GraphResourceGraphOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GraphResourceGraph) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GraphResourceGraphOutput{})
}
