


package v20170301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnostic(ctx *pulumi.Context, args *LookupDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticResult, error) {
	var rv LookupDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20170301:getDiagnostic", args, &rv, opts...)
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
	Enabled bool   `pulumi:"enabled"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Type    string `pulumi:"type"`
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

func (o LookupDiagnosticResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticResultOutput{})
}
