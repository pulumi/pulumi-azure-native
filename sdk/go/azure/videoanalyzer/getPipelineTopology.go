


package videoanalyzer

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPipelineTopology(ctx *pulumi.Context, args *LookupPipelineTopologyArgs, opts ...pulumi.InvokeOption) (*LookupPipelineTopologyResult, error) {
	var rv LookupPipelineTopologyResult
	err := ctx.Invoke("azure-native:videoanalyzer:getPipelineTopology", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPipelineTopologyArgs struct {
	AccountName          string `pulumi:"accountName"`
	PipelineTopologyName string `pulumi:"pipelineTopologyName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}







type LookupPipelineTopologyResult struct {
	Description *string                        `pulumi:"description"`
	Id          string                         `pulumi:"id"`
	Kind        string                         `pulumi:"kind"`
	Name        string                         `pulumi:"name"`
	Parameters  []ParameterDeclarationResponse `pulumi:"parameters"`
	Processors  []EncoderProcessorResponse     `pulumi:"processors"`
	Sinks       []VideoSinkResponse            `pulumi:"sinks"`
	Sku         SkuResponse                    `pulumi:"sku"`
	Sources     []interface{}                  `pulumi:"sources"`
	SystemData  SystemDataResponse             `pulumi:"systemData"`
	Type        string                         `pulumi:"type"`
}

func LookupPipelineTopologyOutput(ctx *pulumi.Context, args LookupPipelineTopologyOutputArgs, opts ...pulumi.InvokeOption) LookupPipelineTopologyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPipelineTopologyResult, error) {
			args := v.(LookupPipelineTopologyArgs)
			r, err := LookupPipelineTopology(ctx, &args, opts...)
			var s LookupPipelineTopologyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPipelineTopologyResultOutput)
}

type LookupPipelineTopologyOutputArgs struct {
	AccountName          pulumi.StringInput `pulumi:"accountName"`
	PipelineTopologyName pulumi.StringInput `pulumi:"pipelineTopologyName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupPipelineTopologyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineTopologyArgs)(nil)).Elem()
}







type LookupPipelineTopologyResultOutput struct{ *pulumi.OutputState }

func (LookupPipelineTopologyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPipelineTopologyResult)(nil)).Elem()
}

func (o LookupPipelineTopologyResultOutput) ToLookupPipelineTopologyResultOutput() LookupPipelineTopologyResultOutput {
	return o
}

func (o LookupPipelineTopologyResultOutput) ToLookupPipelineTopologyResultOutputWithContext(ctx context.Context) LookupPipelineTopologyResultOutput {
	return o
}

func (o LookupPipelineTopologyResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPipelineTopologyResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPipelineTopologyResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupPipelineTopologyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPipelineTopologyResultOutput) Parameters() ParameterDeclarationResponseArrayOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) []ParameterDeclarationResponse { return v.Parameters }).(ParameterDeclarationResponseArrayOutput)
}

func (o LookupPipelineTopologyResultOutput) Processors() EncoderProcessorResponseArrayOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) []EncoderProcessorResponse { return v.Processors }).(EncoderProcessorResponseArrayOutput)
}

func (o LookupPipelineTopologyResultOutput) Sinks() VideoSinkResponseArrayOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) []VideoSinkResponse { return v.Sinks }).(VideoSinkResponseArrayOutput)
}

func (o LookupPipelineTopologyResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupPipelineTopologyResultOutput) Sources() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) []interface{} { return v.Sources }).(pulumi.ArrayOutput)
}

func (o LookupPipelineTopologyResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPipelineTopologyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPipelineTopologyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPipelineTopologyResultOutput{})
}
