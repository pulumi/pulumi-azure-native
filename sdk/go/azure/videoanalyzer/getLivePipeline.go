


package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupLivePipeline(ctx *pulumi.Context, args *LookupLivePipelineArgs, opts ...pulumi.InvokeOption) (*LookupLivePipelineResult, error) {
	var rv LookupLivePipelineResult
	err := ctx.Invoke("azure-native:videoanalyzer:getLivePipeline", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupLivePipelineArgs struct {
	AccountName       string `pulumi:"accountName"`
	LivePipelineName  string `pulumi:"livePipelineName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupLivePipelineResult struct {
	BitrateKbps  int                           `pulumi:"bitrateKbps"`
	Description  *string                       `pulumi:"description"`
	Id           string                        `pulumi:"id"`
	Name         string                        `pulumi:"name"`
	Parameters   []ParameterDefinitionResponse `pulumi:"parameters"`
	State        string                        `pulumi:"state"`
	SystemData   SystemDataResponse            `pulumi:"systemData"`
	TopologyName string                        `pulumi:"topologyName"`
	Type         string                        `pulumi:"type"`
}

func LookupLivePipelineOutput(ctx *pulumi.Context, args LookupLivePipelineOutputArgs, opts ...pulumi.InvokeOption) LookupLivePipelineResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupLivePipelineResult, error) {
			args := v.(LookupLivePipelineArgs)
			r, err := LookupLivePipeline(ctx, &args, opts...)
			var s LookupLivePipelineResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupLivePipelineResultOutput)
}

type LookupLivePipelineOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	LivePipelineName  pulumi.StringInput `pulumi:"livePipelineName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupLivePipelineOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLivePipelineArgs)(nil)).Elem()
}


type LookupLivePipelineResultOutput struct{ *pulumi.OutputState }

func (LookupLivePipelineResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupLivePipelineResult)(nil)).Elem()
}

func (o LookupLivePipelineResultOutput) ToLookupLivePipelineResultOutput() LookupLivePipelineResultOutput {
	return o
}

func (o LookupLivePipelineResultOutput) ToLookupLivePipelineResultOutputWithContext(ctx context.Context) LookupLivePipelineResultOutput {
	return o
}

func (o LookupLivePipelineResultOutput) BitrateKbps() pulumi.IntOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) int { return v.BitrateKbps }).(pulumi.IntOutput)
}

func (o LookupLivePipelineResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupLivePipelineResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupLivePipelineResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupLivePipelineResultOutput) Parameters() ParameterDefinitionResponseArrayOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) []ParameterDefinitionResponse { return v.Parameters }).(ParameterDefinitionResponseArrayOutput)
}

func (o LookupLivePipelineResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupLivePipelineResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupLivePipelineResultOutput) TopologyName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.TopologyName }).(pulumi.StringOutput)
}

func (o LookupLivePipelineResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupLivePipelineResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupLivePipelineResultOutput{})
}
