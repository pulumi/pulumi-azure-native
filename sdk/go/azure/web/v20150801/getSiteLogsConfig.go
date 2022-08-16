


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupSiteLogsConfig(ctx *pulumi.Context, args *LookupSiteLogsConfigArgs, opts ...pulumi.InvokeOption) (*LookupSiteLogsConfigResult, error) {
	var rv LookupSiteLogsConfigResult
	err := ctx.Invoke("azure-native:web/v20150801:getSiteLogsConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSiteLogsConfigArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupSiteLogsConfigResult struct {
	ApplicationLogs       *ApplicationLogsConfigResponse `pulumi:"applicationLogs"`
	DetailedErrorMessages *EnabledConfigResponse         `pulumi:"detailedErrorMessages"`
	FailedRequestsTracing *EnabledConfigResponse         `pulumi:"failedRequestsTracing"`
	HttpLogs              *HttpLogsConfigResponse        `pulumi:"httpLogs"`
	Id                    *string                        `pulumi:"id"`
	Kind                  *string                        `pulumi:"kind"`
	Location              string                         `pulumi:"location"`
	Name                  *string                        `pulumi:"name"`
	Tags                  map[string]string              `pulumi:"tags"`
	Type                  *string                        `pulumi:"type"`
}

func LookupSiteLogsConfigOutput(ctx *pulumi.Context, args LookupSiteLogsConfigOutputArgs, opts ...pulumi.InvokeOption) LookupSiteLogsConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSiteLogsConfigResult, error) {
			args := v.(LookupSiteLogsConfigArgs)
			r, err := LookupSiteLogsConfig(ctx, &args, opts...)
			var s LookupSiteLogsConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSiteLogsConfigResultOutput)
}

type LookupSiteLogsConfigOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupSiteLogsConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteLogsConfigArgs)(nil)).Elem()
}


type LookupSiteLogsConfigResultOutput struct{ *pulumi.OutputState }

func (LookupSiteLogsConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSiteLogsConfigResult)(nil)).Elem()
}

func (o LookupSiteLogsConfigResultOutput) ToLookupSiteLogsConfigResultOutput() LookupSiteLogsConfigResultOutput {
	return o
}

func (o LookupSiteLogsConfigResultOutput) ToLookupSiteLogsConfigResultOutputWithContext(ctx context.Context) LookupSiteLogsConfigResultOutput {
	return o
}

func (o LookupSiteLogsConfigResultOutput) ApplicationLogs() ApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *ApplicationLogsConfigResponse { return v.ApplicationLogs }).(ApplicationLogsConfigResponsePtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) DetailedErrorMessages() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *EnabledConfigResponse { return v.DetailedErrorMessages }).(EnabledConfigResponsePtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) FailedRequestsTracing() EnabledConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *EnabledConfigResponse { return v.FailedRequestsTracing }).(EnabledConfigResponsePtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) HttpLogs() HttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *HttpLogsConfigResponse { return v.HttpLogs }).(HttpLogsConfigResponsePtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupSiteLogsConfigResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupSiteLogsConfigResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupSiteLogsConfigResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSiteLogsConfigResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSiteLogsConfigResultOutput{})
}
