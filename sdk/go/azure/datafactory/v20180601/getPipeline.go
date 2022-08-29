


package v20180601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipeline(ctx *pulumi.Context, args *LookupPipelineArgs, opts ...pulumi.InvokeOption) (*LookupPipelineResult, error) {
	var rv LookupPipelineResult
	err := ctx.Invoke("azure-native:datafactory/v20180601:getPipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineArgs struct {
	FactoryName       string `pulumi:"factoryName"`
	PipelineName      string `pulumi:"pipelineName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPipelineResult struct {
	Activities    []interface{}                             `pulumi:"activities"`
	Annotations   []interface{}                             `pulumi:"annotations"`
	Concurrency   *int                                      `pulumi:"concurrency"`
	Description   *string                                   `pulumi:"description"`
	Etag          string                                    `pulumi:"etag"`
	Folder        *PipelineResponseFolder                   `pulumi:"folder"`
	Id            string                                    `pulumi:"id"`
	Name          string                                    `pulumi:"name"`
	Parameters    map[string]ParameterSpecificationResponse `pulumi:"parameters"`
	Policy        *PipelinePolicyResponse                   `pulumi:"policy"`
	RunDimensions map[string]interface{}                    `pulumi:"runDimensions"`
	Type          string                                    `pulumi:"type"`
	Variables     map[string]VariableSpecificationResponse  `pulumi:"variables"`
}

func LookupPipelineOutput(ctx *pulumi.Context, args LookupPipelineOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineResult, error) {
			args := v.(LookupPipelineArgs)
			r, err := LookupPipeline(ctx, &args, opts...)
			var s LookupPipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineResultOutput)
}

type LookupPipelineOutputArgs struct {
	FactoryName       pulumi.StringInput `pulumi:"factoryName"`
	PipelineName      pulumi.StringInput `pulumi:"pipelineName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineArgs)(nil)).Elem()
}


type LookupPipelineResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineResult)(nil)).Elem()
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutput() LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) ToLookupPipelineResultOutputWithContext(ctx context.Context) LookupPipelineResultOutput {
	return o
}

func (o LookupPipelineResultOutput) Activities() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []interface{} { return v.Activities }).(pulumi.ArrayOutput)
}

func (o LookupPipelineResultOutput) Annotations() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupPipelineResult) []interface{} { return v.Annotations }).(pulumi.ArrayOutput)
}

func (o LookupPipelineResultOutput) Concurrency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *int { return v.Concurrency }).(pulumi.IntPtrOutput)
}

func (o LookupPipelineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupPipelineResultOutput) Folder() PipelineResponseFolderPtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *PipelineResponseFolder { return v.Folder }).(PipelineResponseFolderPtrOutput)
}

func (o LookupPipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPipelineResultOutput) Parameters() ParameterSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupPipelineResult) map[string]ParameterSpecificationResponse { return v.Parameters }).(ParameterSpecificationResponseMapOutput)
}

func (o LookupPipelineResultOutput) Policy() PipelinePolicyResponsePtrOutput {
	return o.ApplyT(func(v LookupPipelineResult) *PipelinePolicyResponse { return v.Policy }).(PipelinePolicyResponsePtrOutput)
}

func (o LookupPipelineResultOutput) RunDimensions() pulumi.MapOutput {
	return o.ApplyT(func(v LookupPipelineResult) map[string]interface{} { return v.RunDimensions }).(pulumi.MapOutput)
}

func (o LookupPipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupPipelineResultOutput) Variables() VariableSpecificationResponseMapOutput {
	return o.ApplyT(func(v LookupPipelineResult) map[string]VariableSpecificationResponse { return v.Variables }).(VariableSpecificationResponseMapOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineResultOutput{})
}
