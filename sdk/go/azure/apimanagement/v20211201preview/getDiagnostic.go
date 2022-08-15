


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20211201preview:getDiagnostic", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticArgs struct {
	DiagnosticId      string `pulumi:"diagnosticId"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupDiagnosticResult struct {
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

func LookupDiagnosticOutput(ctx *pulumi.Context, args LookupDiagnosticOutputArgs, opts ...pulumi.InvokeOption) LookupDiagnosticResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiagnosticResult, error) {
			args := v.(LookupDiagnosticArgs)
			r, err := LookupDiagnostic(ctx, &args, opts...)
			var s LookupDiagnosticResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiagnosticResultOutput)
}

type LookupDiagnosticOutputArgs struct {
	DiagnosticId      pulumi.StringInput `pulumi:"diagnosticId"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupDiagnosticOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticArgs)(nil)).Elem()
}


type LookupDiagnosticResultOutput struct{ *pulumi.OutputState }

func (LookupDiagnosticResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticResult)(nil)).Elem()
}

func (o LookupDiagnosticResultOutput) ToLookupDiagnosticResultOutput() LookupDiagnosticResultOutput {
	return o
}

func (o LookupDiagnosticResultOutput) ToLookupDiagnosticResultOutputWithContext(ctx context.Context) LookupDiagnosticResultOutput {
	return o
}

func (o LookupDiagnosticResultOutput) AlwaysLog() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.AlwaysLog }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticResultOutput) Backend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Backend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o LookupDiagnosticResultOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o LookupDiagnosticResultOutput) HttpCorrelationProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.HttpCorrelationProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) LogClientIp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *bool { return v.LogClientIp }).(pulumi.BoolPtrOutput)
}

func (o LookupDiagnosticResultOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.LoggerId }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) Metrics() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *bool { return v.Metrics }).(pulumi.BoolPtrOutput)
}

func (o LookupDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) OperationNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.OperationNameFormat }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticResultOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *SamplingSettingsResponse { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

func (o LookupDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) Verbosity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *string { return v.Verbosity }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticResultOutput{})
}
