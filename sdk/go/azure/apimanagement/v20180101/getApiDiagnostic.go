


package v20180101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupApiDiagnostic(ctx *pulumi.Context, args *LookupApiDiagnosticArgs, opts ...pulumi.InvokeOption) (*LookupApiDiagnosticResult, error) {
	var rv LookupApiDiagnosticResult
	err := ctx.Invoke("azure-native:apimanagement/v20180101:getApiDiagnostic", args, &rv, opts...)
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
	Enabled bool   `pulumi:"enabled"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	Type    string `pulumi:"type"`
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

func (o LookupApiDiagnosticResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupApiDiagnosticResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupApiDiagnosticResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupApiDiagnosticResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupApiDiagnosticResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupApiDiagnosticResultOutput{})
}
