


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PipelineJob struct {
	pulumi.CustomResourceState

	Description  pulumi.StringPtrOutput                 `pulumi:"description"`
	Error        PipelineJobErrorResponseOutput         `pulumi:"error"`
	Expiration   pulumi.StringOutput                    `pulumi:"expiration"`
	Name         pulumi.StringOutput                    `pulumi:"name"`
	Parameters   ParameterDefinitionResponseArrayOutput `pulumi:"parameters"`
	State        pulumi.StringOutput                    `pulumi:"state"`
	SystemData   SystemDataResponseOutput               `pulumi:"systemData"`
	TopologyName pulumi.StringOutput                    `pulumi:"topologyName"`
	Type         pulumi.StringOutput                    `pulumi:"type"`
}


func NewPipelineJob(ctx *pulumi.Context,
	name string, args *PipelineJobArgs, opts ...pulumi.ResourceOption) (*PipelineJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AccountName == nil {
		return nil, errors.New("invalid value for required argument 'AccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.TopologyName == nil {
		return nil, errors.New("invalid value for required argument 'TopologyName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:videoanalyzer:PipelineJob"),
		},
	})
	opts = append(opts, aliases)
	var resource PipelineJob
	err := ctx.RegisterResource("azure-native:videoanalyzer/v20211101preview:PipelineJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPipelineJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineJobState, opts ...pulumi.ResourceOption) (*PipelineJob, error) {
	var resource PipelineJob
	err := ctx.ReadResource("azure-native:videoanalyzer/v20211101preview:PipelineJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type pipelineJobState struct {
}

type PipelineJobState struct {
}

func (PipelineJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineJobState)(nil)).Elem()
}

type pipelineJobArgs struct {
	AccountName       string                `pulumi:"accountName"`
	Description       *string               `pulumi:"description"`
	Parameters        []ParameterDefinition `pulumi:"parameters"`
	PipelineJobName   *string               `pulumi:"pipelineJobName"`
	ResourceGroupName string                `pulumi:"resourceGroupName"`
	TopologyName      string                `pulumi:"topologyName"`
}


type PipelineJobArgs struct {
	AccountName       pulumi.StringInput
	Description       pulumi.StringPtrInput
	Parameters        ParameterDefinitionArrayInput
	PipelineJobName   pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	TopologyName      pulumi.StringInput
}

func (PipelineJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*pipelineJobArgs)(nil)).Elem()
}

type PipelineJobInput interface {
	pulumi.Input

	ToPipelineJobOutput() PipelineJobOutput
	ToPipelineJobOutputWithContext(ctx context.Context) PipelineJobOutput
}

func (*PipelineJob) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineJob)(nil)).Elem()
}

func (i *PipelineJob) ToPipelineJobOutput() PipelineJobOutput {
	return i.ToPipelineJobOutputWithContext(context.Background())
}

func (i *PipelineJob) ToPipelineJobOutputWithContext(ctx context.Context) PipelineJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineJobOutput)
}

type PipelineJobOutput struct{ *pulumi.OutputState }

func (PipelineJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineJob)(nil)).Elem()
}

func (o PipelineJobOutput) ToPipelineJobOutput() PipelineJobOutput {
	return o
}

func (o PipelineJobOutput) ToPipelineJobOutputWithContext(ctx context.Context) PipelineJobOutput {
	return o
}

func (o PipelineJobOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PipelineJobOutput) Error() PipelineJobErrorResponseOutput {
	return o.ApplyT(func(v *PipelineJob) PipelineJobErrorResponseOutput { return v.Error }).(PipelineJobErrorResponseOutput)
}

func (o PipelineJobOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.Expiration }).(pulumi.StringOutput)
}

func (o PipelineJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PipelineJobOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v *PipelineJob) ParameterDefinitionResponseArrayOutput { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

func (o PipelineJobOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.State }).(pulumi.StringOutput)
}

func (o PipelineJobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PipelineJob) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PipelineJobOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.TopologyName }).(pulumi.StringOutput)
}

func (o PipelineJobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PipelineJob) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineJobOutput{})
}
