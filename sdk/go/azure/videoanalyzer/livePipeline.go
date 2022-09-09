


package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LivePipeline struct {
	pulumi.CustomResourceState

	BitrateKbps  pulumi.IntOutput                       `pulumi:"bitrateKbps"`
	Description  pulumi.StringPtrOutput                 `pulumi:"description"`
	Name         pulumi.StringOutput                    `pulumi:"name"`
	Parameters   ParameterDefinitionResponseArrayOutput `pulumi:"parameters"`
	State        pulumi.StringOutput                    `pulumi:"state"`
	SystemData   SystemDataResponseOutput               `pulumi:"systemData"`
	TopologyName pulumi.StringOutput                    `pulumi:"topologyName"`
	Type         pulumi.StringOutput                    `pulumi:"type"`
}


func NewLivePipeline(ctx *pulumi.Context,
	name string, args *LivePipelineArgs, opts ...pulumi.ResourceOption) (*LivePipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.BitrateKbps == nil {
		return nil, errors.New("invalid value for required argument 'BitrateKbps'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopologyName == nil {
		return nil, errors.New("invalid value for required argument 'TopologyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer/v20211101preview:LivePipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource LivePipeline
	err := ctx.RegisterResource("azure-native:videoanalyzer:LivePipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLivePipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LivePipelineState, opts ...pulumi.ResourceOption) (*LivePipeline, error) {
	var resource LivePipeline
	err := ctx.ReadResource("azure-native:videoanalyzer:LivePipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type livePipelineState struct {
}

type LivePipelineState struct {
}

func (LivePipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*livePipelineState)(nil)).Elem()
}

type livePipelineArgs struct {
	AccountName       string                `pulumi:"accountName"`
	BitrateKbps       int                   `pulumi:"bitrateKbps"`
	Description       *string               `pulumi:"description"`
	LivePipelineName  *string               `pulumi:"livePipelineName"`
	Parameters        []ParameterDefinition `pulumi:"parameters"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	TopologyName      string                `pulumi:"topologyName"`
}


type LivePipelineArgs struct {
	AccountName       pulumi.StringInput
	BitrateKbps       pulumi.IntInput
	Description       pulumi.StringPtrInput
	LivePipelineName  pulumi.StringPtrInput
	Parameters        ParameterDefinitionArrayInput
	ResourceGroupName pulumi.StringInput
	TopologyName      pulumi.StringInput
}

func (LivePipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*livePipelineArgs)(nil)).Elem()
}

type LivePipelineInput interface {
	pulumi.Input

	ToLivePipelineOutput() LivePipelineOutput
	ToLivePipelineOutputWithContext(ctx context.Context) LivePipelineOutput
}

func (*LivePipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**LivePipeline)(nil)).Elem()
}

func (i *LivePipeline) ToLivePipelineOutput() LivePipelineOutput {
	return i.ToLivePipelineOutputWithContext(context.Background())
}

func (i *LivePipeline) ToLivePipelineOutputWithContext(ctx context.Context) LivePipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LivePipelineOutput)
}

type LivePipelineOutput struct{ *pulumi.OutputState }

func (LivePipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LivePipeline)(nil)).Elem()
}

func (o LivePipelineOutput) ToLivePipelineOutput() LivePipelineOutput {
	return o
}

func (o LivePipelineOutput) ToLivePipelineOutputWithContext(ctx context.Context) LivePipelineOutput {
	return o
}

func (o LivePipelineOutput) BitrateKbps() pulumi.IntOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.IntOutput { return v.BitrateKbps }).(pulumi.IntOutput)
}

func (o LivePipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LivePipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LivePipelineOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *LivePipeline) ParameterDefinitionResponseArrayOutput { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

func (o LivePipelineOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o LivePipelineOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LivePipeline) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LivePipelineOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.TopologyName }).(pulumi.StringOutput)
}

func (o LivePipelineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LivePipeline) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LivePipelineOutput{})
}
