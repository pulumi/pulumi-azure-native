


package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PipelineTopology struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput                  `pulumi:"description"`
	Kind        pulumi.StringOutput                     `pulumi:"kind"`
	Name        pulumi.StringOutput                     `pulumi:"name"`
	Parameters  ParameterDeclarationResponseArrayOutput `pulumi:"parameters"`
	Processors  EncoderProcessorResponseArrayOutput     `pulumi:"processors"`
	Sinks       VideoSinkResponseArrayOutput            `pulumi:"sinks"`
	Sku         SkuResponseOutput                       `pulumi:"sku"`
	Sources     pulumi.ArrayOutput                      `pulumi:"sources"`
	SystemData  SystemDataResponseOutput                `pulumi:"systemData"`
	Type        pulumi.StringOutput                     `pulumi:"type"`
}


func NewPipelineTopology(ctx *pulumi.Context,
	name string, args *PipelineTopologyArgs, opts ...pulumi.ResourceOption) (*PipelineTopology, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Sinks == nil {
		return nil, errors.New("invalid value for required argument 'Sinks'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	if args.Sources == nil {
		return nil, errors.New("invalid value for required argument 'Sources'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:PipelineTopology"),
		},
	})
	opts = append(opts, aliases)
	var resource PipelineTopology
	err := ctx.RegisterResource("azure-native:videoanalyzer:PipelineTopology", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPipelineTopology(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineTopologyState, opts ...pulumi.ResourceOption) (*PipelineTopology, error) {
	var resource PipelineTopology
	err := ctx.ReadResource("azure-native:videoanalyzer:PipelineTopology", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type pipelineTopologyState struct {
}

type PipelineTopologyState struct {
}

func (PipelineTopologyState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTopologyState)(nil)).Elem()
}

type pipelineTopologyArgs struct {
	AccountName          string                 `pulumi:"accountName"`
	Description          *string                `pulumi:"description"`
	Kind                 string                 `pulumi:"kind"`
	Parameters           []ParameterDeclaration `pulumi:"parameters"`
	PipelineTopologyName *string                `pulumi:"pipelineTopologyName"`
	Processors           []EncoderProcessor     `pulumi:"processors"`
	ResourceGroupName    string                 `pulumi:"resourceGroupName"`
	Sinks                []VideoSink            `pulumi:"sinks"`
	Sku                  Sku                    `pulumi:"sku"`
	Sources              []interface{}          `pulumi:"sources"`
}


type PipelineTopologyArgs struct {
	AccountName          pulumi.StringInput
	Description          pulumi.StringPtrInput
	Kind                 pulumi.StringInput
	Parameters           ParameterDeclarationArrayInput
	PipelineTopologyName pulumi.StringPtrInput
	Processors           EncoderProcessorArrayInput
	ResourceGroupName    pulumi.StringInput
	Sinks                VideoSinkArrayInput
	Sku                  SkuInput
	Sources              pulumi.ArrayInput
}

func (PipelineTopologyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineTopologyArgs)(nil)).Elem()
}

type PipelineTopologyInput interface {
	pulumi.Input

	ToPipelineTopologyOutput() PipelineTopologyOutput
	ToPipelineTopologyOutputWithContext(ctx context.Context) PipelineTopologyOutput
}

func (*PipelineTopology) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTopology)(nil))
}

func (i *PipelineTopology) ToPipelineTopologyOutput() PipelineTopologyOutput {
	return i.ToPipelineTopologyOutputWithContext(context.Background())
}

func (i *PipelineTopology) ToPipelineTopologyOutputWithContext(ctx context.Context) PipelineTopologyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTopologyOutput)
}

type PipelineTopologyOutput struct{ *pulumi.OutputState }

func (PipelineTopologyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTopology)(nil))
}

func (o PipelineTopologyOutput) ToPipelineTopologyOutput() PipelineTopologyOutput {
	return o
}

func (o PipelineTopologyOutput) ToPipelineTopologyOutputWithContext(ctx context.Context) PipelineTopologyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PipelineTopologyOutput{})
}
