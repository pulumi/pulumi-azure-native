


package v20180601

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DataFlow struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringOutput `pulumi:"etag"`
	Name       pulumi.StringOutput `pulumi:"name"`
	Properties pulumi.AnyOutput    `pulumi:"properties"`
	Type       pulumi.StringOutput `pulumi:"type"`
}


func NewDataFlow(ctx *pulumi.Context,
	name string, args *DataFlowArgs, opts ...pulumi.ResourceOption) (*DataFlow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.Properties == nil {
		return nil, errors.New("invalid value for required argument 'Properties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory:DataFlow"),
		},
	})
	opts = append(opts, aliases)
	var resource DataFlow
	err := ctx.RegisterResource("azure-native:datafactory/v20180601:DataFlow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetDataFlow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *DataFlowState, opts ...pulumi.ResourceOption) (*DataFlow, error) {
	var resource DataFlow
	err := ctx.ReadResource("azure-native:datafactory/v20180601:DataFlow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type dataFlowState struct {
}

type DataFlowState struct {
}

func (DataFlowState) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowState)(nil)).Elem()
}

type dataFlowArgs struct {
	DataFlowName      *string     `pulumi:"dataFlowName"`
	FactoryName       string      `pulumi:"factoryName"`
	Properties        interface{} `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
}


type DataFlowArgs struct {
	DataFlowName      pulumi.StringPtrInput
	FactoryName       pulumi.StringInput
	Properties        pulumi.Input
	ResourceGroupName pulumi.StringInput
}

func (DataFlowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*dataFlowArgs)(nil)).Elem()
}

type DataFlowInput interface {
	pulumi.Input

	ToDataFlowOutput() DataFlowOutput
	ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput
}

func (*DataFlow) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlow)(nil))
}

func (i *DataFlow) ToDataFlowOutput() DataFlowOutput {
	return i.ToDataFlowOutputWithContext(context.Background())
}

func (i *DataFlow) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataFlowOutput)
}

type DataFlowOutput struct{ *pulumi.OutputState }

func (DataFlowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataFlow)(nil))
}

func (o DataFlowOutput) ToDataFlowOutput() DataFlowOutput {
	return o
}

func (o DataFlowOutput) ToDataFlowOutputWithContext(ctx context.Context) DataFlowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(DataFlowOutput{})
}
