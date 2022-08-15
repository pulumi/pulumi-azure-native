


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebAppDiagnosticLogsConfiguration(ctx *pulumi.Context, args *LookupWebAppDiagnosticLogsConfigurationArgs, opts ...pulumi.InvokeOption) (*LookupWebAppDiagnosticLogsConfigurationResult, error) {
	var rv LookupWebAppDiagnosticLogsConfigurationResult
	err := ctx.Invoke("azure-native:web/v20190801:getWebAppDiagnosticLogsConfiguration", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return rv.Defaults(), nil
}

type LookupWebAppDiagnosticLogsConfigurationArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebAppDiagnosticLogsConfigurationResult struct {
	ApplicationLogs       *ApplicationLogsConfigResponse `pulumi:"applicationLogs"`
	DetailedErrorMessages *EnabledConfigResponse         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing *EnabledConfigResponse         `pulumi:"failedRequestsTracing"`
	HttpLogs              *HttpLogsConfigResponse        `pulumi:"httpLogs"`
	Id                    string                         `pulumi:"id"`
	Kind                  *string                        `pulumi:"kind"`
	Name                  string                         `pulumi:"name"`
	Type                  string                         `pulumi:"type"`
}


func (val *LookupWebAppDiagnosticLogsConfigurationResult) Defaults() *LookupWebAppDiagnosticLogsConfigurationResult {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ApplicationLogs = tmp.ApplicationLogs.Defaults()

	return &tmp
}

func LookupWebAppDiagnosticLogsConfigurationOutput(ctx *pulumi.Context, args LookupWebAppDiagnosticLogsConfigurationOutputArgs, opts ...pulumi.InvokeOption) LookupWebAppDiagnosticLogsConfigurationResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebAppDiagnosticLogsConfigurationResult, error) {
			args := v.(LookupWebAppDiagnosticLogsConfigurationArgs)
			r, err := LookupWebAppDiagnosticLogsConfiguration(ctx, &args, opts...)
			var s LookupWebAppDiagnosticLogsConfigurationResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebAppDiagnosticLogsConfigurationResultOutput)
}

type LookupWebAppDiagnosticLogsConfigurationOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebAppDiagnosticLogsConfigurationOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDiagnosticLogsConfigurationArgs)(nil)).Elem()
}


type LookupWebAppDiagnosticLogsConfigurationResultOutput struct{ *pulumi.OutputState }

func (LookupWebAppDiagnosticLogsConfigurationResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebAppDiagnosticLogsConfigurationResult)(nil)).Elem()
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ToLookupWebAppDiagnosticLogsConfigurationResultOutput() LookupWebAppDiagnosticLogsConfigurationResultOutput {
	return o
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ToLookupWebAppDiagnosticLogsConfigurationResultOutputWithContext(ctx context.Context) LookupWebAppDiagnosticLogsConfigurationResultOutput {
	return o
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) ApplicationLogs() ApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *ApplicationLogsConfigResponse {
		return v.ApplicationLogs
	}).(ApplicationLogsConfigResponsePtrOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) DetailedErrorMessages() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *EnabledConfigResponse {
		return v.DetailedErrorMessages
	}).(EnabledConfigResponsePtrOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) FailedRequestsTracing() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *EnabledConfigResponse {
		return v.FailedRequestsTracing
	}).(EnabledConfigResponsePtrOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) HttpLogs() HttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *HttpLogsConfigResponse { return v.HttpLogs }).(HttpLogsConfigResponsePtrOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebAppDiagnosticLogsConfigurationResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebAppDiagnosticLogsConfigurationResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebAppDiagnosticLogsConfigurationResultOutput{})
}
