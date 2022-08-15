


package datafactory

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Pipeline struct {
	pulumi.CustomResourceState

	Activities    pulumi.ArrayOutput                      `pulumi:"activities"`
	Annotations   pulumi.ArrayOutput                      `pulumi:"annotations"`
	Concurrency   pulumi.IntPtrOutput                     `pulumi:"concurrency"`
	Description   pulumi.StringPtrOutput                  `pulumi:"description"`
	Etag          pulumi.StringOutput                     `pulumi:"etag"`
	Folder        PipelineResponseFolderPtrOutput         `pulumi:"folder"`
	Name          pulumi.StringOutput                     `pulumi:"name"`
	Parameters    ParameterSpecificationResponseMapOutput `pulumi:"parameters"`
	Policy        PipelinePolicyResponsePtrOutput         `pulumi:"policy"`
	RunDimensions pulumi.MapOutput                        `pulumi:"runDimensions"`
	Type          pulumi.StringOutput                     `pulumi:"type"`
	Variables     VariableSpecificationResponseMapOutput  `pulumi:"variables"`
}


func NewPipeline(ctx *pulumi.Context,
	name string, args *PipelineArgs, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.FactoryName == nil {
		return nil, errors.New("invalid value for required argument 'FactoryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datafactory/v20170901preview:Pipeline"),
		},
		{
			Type: pulumi.String("azure-native:datafactory/v20180601:Pipeline"),
		},
	})
	opts = append(opts, aliases)
	var resource Pipeline
	err := ctx.RegisterResource("azure-native:datafactory:Pipeline", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPipeline(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PipelineState, opts ...pulumi.ResourceOption) (*Pipeline, error) {
	var resource Pipeline
	err := ctx.ReadResource("azure-native:datafactory:Pipeline", name, id, state, &resource, opts...)
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
	Activities        []interface{}                     `pulumi:"activities"`
	Annotations       []interface{}                     `pulumi:"annotations"`
	Concurrency       *int                              `pulumi:"concurrency"`
	Description       *string                           `pulumi:"description"`
	FactoryName       string                            `pulumi:"factoryName"`
	Folder            *PipelineFolder                   `pulumi:"folder"`
	Parameters        map[string]ParameterSpecification `pulumi:"parameters"`
	PipelineName      *string                           `pulumi:"pipelineName"`
	Policy            *PipelinePolicy                   `pulumi:"policy"`
	ResourceGroupName string                            `pulumi:"resourceGroupName"`
	RunDimensions     map[string]interface{}            `pulumi:"runDimensions"`
	Variables         map[string]VariableSpecification  `pulumi:"variables"`
}


type PipelineArgs struct {
	Activities        pulumi.ArrayInput
	Annotations       pulumi.ArrayInput
	Concurrency       pulumi.IntPtrInput
	Description       pulumi.StringPtrInput
	FactoryName       pulumi.StringInput
	Folder            PipelineFolderPtrInput
	Parameters        ParameterSpecificationMapInput
	PipelineName      pulumi.StringPtrInput
	Policy            PipelinePolicyPtrInput
	ResourceGroupName pulumi.StringInput
	RunDimensions     pulumi.MapInput
	Variables         VariableSpecificationMapInput
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

func (o PipelineOutput) Activities() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.ArrayOutput { return v.Activities }).(pulumi.ArrayOutput)
}

func (o PipelineOutput) Annotations() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.ArrayOutput { return v.Annotations }).(pulumi.ArrayOutput)
}

func (o PipelineOutput) Concurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.IntPtrOutput { return v.Concurrency }).(pulumi.IntPtrOutput)
}

func (o PipelineOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PipelineOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Etag }).(pulumi.StringOutput)
}

func (o PipelineOutput) Folder() PipelineResponseFolderPtrOutput {
	return o.ApplyT(func(v *Pipeline) PipelineResponseFolderPtrOutput { return v.Folder }).(PipelineResponseFolderPtrOutput)
}

func (o PipelineOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PipelineOutput) Parameters() ParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v *Pipeline) ParameterSpecificationResponseMapOutput { return v.Parameters }).(ParameterSpecificationResponseMapOutput)
}

func (o PipelineOutput) Policy() PipelinePolicyResponsePtrOutput {
	return o.ApplyT(func(v *Pipeline) PipelinePolicyResponsePtrOutput { return v.Policy }).(PipelinePolicyResponsePtrOutput)
}

func (o PipelineOutput) RunDimensions() pulumi.MapOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.MapOutput { return v.RunDimensions }).(pulumi.MapOutput)
}

func (o PipelineOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Pipeline) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o PipelineOutput) Variables() VariableSpecificationResponseMapOutput {
	return o.ApplyT(func(v *Pipeline) VariableSpecificationResponseMapOutput { return v.Variables }).(VariableSpecificationResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(PipelineOutput{})
}
