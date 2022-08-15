


package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipelineJob(ctx *pulumi.Context, args *LookupPipelineJobArgs, opts ...pulumi.InvokeOption) (*LookupPipelineJobResult, error) {
	var rv LookupPipelineJobResult
	err := ctx.Invoke("azure-native:videoanalyzer:getPipelineJob", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineJobArgs struct {
	AccountName       string `pulumi:"accountName"`
	PipelineJobName   string `pulumi:"pipelineJobName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupPipelineJobResult struct {
	Description  *string                       `pulumi:"description"`
	Error        PipelineJobErrorResponse      `pulumi:"error"`
	Expiration   string                        `pulumi:"expiration"`
	Id           string                        `pulumi:"id"`
	Name         string                        `pulumi:"name"`
	Parameters   []ParameterDefinitionResponse `pulumi:"parameters"`
	State        string                        `pulumi:"state"`
	SystemData   SystemDataResponse            `pulumi:"systemData"`
	TopologyName string                        `pulumi:"topologyName"`
	Type         string                        `pulumi:"type"`
}

func LookupPipelineJobOutput(ctx *pulumi.Context, args LookupPipelineJobOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineJobResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineJobResult, error) {
			args := v.(LookupPipelineJobArgs)
			r, err := LookupPipelineJob(ctx, &args, opts...)
			var s LookupPipelineJobResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineJobResultOutput)
}

type LookupPipelineJobOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	PipelineJobName   pulumi.StringInput `pulumi:"pipelineJobName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineJobOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineJobArgs)(nil)).Elem()
}


type LookupPipelineJobResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineJobResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineJobResult)(nil)).Elem()
}

func (o LookupPipelineJobResultOutput) ToLookupPipelineJobResultOutput() LookupPipelineJobResultOutput {
	return o
}

func (o LookupPipelineJobResultOutput) ToLookupPipelineJobResultOutputWithContext(ctx context.Context) LookupPipelineJobResultOutput {
	return o
}

func (o LookupPipelineJobResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineJobResultOutput) Error() PipelineJobErrorResponseOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) PipelineJobErrorResponse { return v.Error }).(PipelineJobErrorResponseOutput)
}

func (o LookupPipelineJobResultOutput) Expiration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) string { return v.Expiration }).(pulumi.StringOutput)
}

func (o LookupPipelineJobResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPipelineJobResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPipelineJobResultOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) []ParameterDefinitionResponse { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

func (o LookupPipelineJobResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupPipelineJobResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPipelineJobResultOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) string { return v.TopologyName }).(pulumi.StringOutput)
}

func (o LookupPipelineJobResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineJobResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineJobResultOutput{})
}
