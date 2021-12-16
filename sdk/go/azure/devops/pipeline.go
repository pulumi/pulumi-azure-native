


package devops

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pipeline struct {
	pulumi.CustomResourceState

	BootstrapConfiguration BootstrapConfigurationResponseOutput `pulumi:"bootstrapConfiguration"`
	Location               pulumi.StringPtrOutput               `pulumi:"location"`
	Name                   pulumi.StringOutput                  `pulumi:"name"`
	PipelineId             pulumi.IntOutput                     `pulumi:"pipelineId"`
	PipelineType           pulumi.StringOutput                  `pulumi:"pipelineType"`
	SystemData             SystemDataResponseOutput             `pulumi:"systemData"`
	Tags                   pulumi.StringMapOutput               `pulumi:"tags"`
	Type                   pulumi.StringOutput                  `pulumi:"type"`
}


func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BootstrapConfiguration == nil {
		return nil, errors.New("invalid value for required argument 'BootstrapConfiguration'")
	}
	if args.PipelineType == nil {
		return nil, errors.New("invalid value for required argument 'PipelineType'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devops/v20190701preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-native:devops/v20200713preview:Pipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource Pipeline
	err := ctx.RegisterResource("azure-native:devops:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("azure-native:devops:Pipeline", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type pipelineState struct {
}

type PipelineState struct {
}

func (PipelineState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineState)(nil)).Elem()
}

type pipelineArgs struct {
	BootstrapConfiguration BootstrapConfiguration `pulumi:"bootstrapConfiguration"`
	Location               *string                `pulumi:"location"`
	PipelineName           *string                `pulumi:"pipelineName"`
	PipelineType           string                 `pulumi:"pipelineType"`
	ResourceGroupName      string                 `pulumi:"resourceGroupName"`
	Tags                   map[string]string      `pulumi:"tags"`
}


type PipelineArgs struct {
	BootstrapConfiguration BootstrapConfigurationInput
	Location               pulumi.StringPtrInput
	PipelineName           pulumi.StringPtrInput
	PipelineType           pulumi.StringInput
	ResourceGroupName      pulumi.StringInput
	Tags                   pulumi.StringMapInput
}

func (PipelineArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineArgs)(nil)).Elem()
}

type PipelineInput interface {
	pulumi.Input

	ToPipelineOutput() PipelineOutput
	ToPipelineOutputWithContext(ctx context.Context) PipelineOutput
}

func (*Pipeline) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (i *Pipeline) ToPipelineOutput() PipelineOutput {
	return i.ToPipelineOutputWithContext(context.Background())
}

func (i *Pipeline) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineOutput)
}

type PipelineOutput struct{ *pulumi.OutputState }

func (PipelineOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Pipeline)(nil)).Elem()
}

func (o PipelineOutput) ToPipelineOutput() PipelineOutput {
	return o
}

func (o PipelineOutput) ToPipelineOutputWithContext(ctx context.Context) PipelineOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PipelineOutput{})
}
