


package v20201201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiDiagnostic(ctx *pulumi.Context, args *LookupApiDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupApiDiagnosticResult, error) {
	var rv LookupApiDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20201201:getApiDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupApiDiagnosticArgs struct {
	ApiId             string `pulumi:"apiId"`
	DiagnosticId      string `pulumi:"diagnosticId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupApiDiagnosticResult struct {
	AlwaysLog               *string                             `pulumi:"alwaysLog"`
	Backend                 *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	Frontend                *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	HttpCorrelationProtocol *string                             `pulumi:"httpCorrelationProtocol"`
	Id                      string                              `pulumi:"id"`
	LogClientIp             *bool                               `pulumi:"logClientIp"`
	LoggerId                string                              `pulumi:"loggerId"`
	Metrics                 *bool                               `pulumi:"metrics"`
	Name                    string                              `pulumi:"name"`
	OperationNameFormat     *string                             `pulumi:"operationNameFormat"`
	Sampling                *SamplingSettingsResponse           `pulumi:"sampling"`
	Type                    string                              `pulumi:"type"`
	Verbosity               *string                             `pulumi:"verbosity"`
}

func LookupApiDiagnosticOutput(ctx *pulumi.Context, args LookupApiDiagnosticOutputArgs, opts ...pulumi.InvokeOption) LookupApiDiagnosticResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupApiDiagnosticResult, error) {
			args := v.(LookupApiDiagnosticArgs)
			r, err := LookupApiDiagnostic(ctx, &args, opts...)
			var s LookupApiDiagnosticResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupApiDiagnosticResultOutput)
}

type LookupApiDiagnosticOutputArgs struct {
	ApiId             pulumi.StringInput `pulumi:"apiId"`
	DiagnosticId      pulumi.StringInput `pulumi:"diagnosticId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupApiDiagnosticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDiagnosticArgs)(nil)).Elem()
}


type LookupApiDiagnosticResultOutput struct{ *pulumi.OutputState }

func (LookupApiDiagnosticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupApiDiagnosticResult)(nil)).Elem()
}

func (o LookupApiDiagnosticResultOutput) ToLookupApiDiagnosticResultOutput() LookupApiDiagnosticResultOutput {
	return o
}

func (o LookupApiDiagnosticResultOutput) ToLookupApiDiagnosticResultOutputWithContext(ctx context.Context) LookupApiDiagnosticResultOutput {
	return o
}

func (o LookupApiDiagnosticResultOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

func (o LookupApiDiagnosticResultOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o LookupApiDiagnosticResultOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o LookupApiDiagnosticResultOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupApiDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiDiagnosticResultOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *bool { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

func (o LookupApiDiagnosticResultOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.LoggerId }).(pulumi.StringOutput)
}

func (o LookupApiDiagnosticResultOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *bool { return v.Metrics }).(pulumi.BoolPtrOutput)
}

func (o LookupApiDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiDiagnosticResultOutput) OperationNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.OperationNameFormat }).(pulumi.StringPtrOutput)
}

func (o LookupApiDiagnosticResultOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *SamplingSettingsResponse { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

func (o LookupApiDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupApiDiagnosticResultOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) *string { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiDiagnosticResultOutput{})
}
