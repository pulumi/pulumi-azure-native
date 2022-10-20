


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180601preview:getDiagnostic", args, &rv, opts...)
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
	AlwaysLog                    *string                             `pulumi:"alwaysLog"`
	Backend                      *PipelineDiagnosticSettingsResponse `pulumi:"backend"`
	EnableHttpCorrelationHeaders *bool                               `pulumi:"enableHttpCorrelationHeaders"`
	Frontend                     *PipelineDiagnosticSettingsResponse `pulumi:"frontend"`
	Id                           string                              `pulumi:"id"`
	LoggerId                     string                              `pulumi:"loggerId"`
	Name                         string                              `pulumi:"name"`
	Sampling                     *SamplingSettingsResponse           `pulumi:"sampling"`
	Type                         string                              `pulumi:"type"`
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

func (o LookupDiagnosticResultOutput) EnableHttpCorrelationHeaders() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *bool { return v.EnableHttpCorrelationHeaders }).(pulumi.BoolPtrOutput)
}

func (o LookupDiagnosticResultOutput) Frontend() PipelineDiagnosticSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *PipelineDiagnosticSettingsResponse { return v.Frontend }).(PipelineDiagnosticSettingsResponsePtrOutput)
}

func (o LookupDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) LoggerId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.LoggerId }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) Sampling() SamplingSettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) *SamplingSettingsResponse { return v.Sampling }).(SamplingSettingsResponsePtrOutput)
}

func (o LookupDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticResultOutput{})
}
