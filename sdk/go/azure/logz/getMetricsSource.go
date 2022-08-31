


package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMetricsSource(ctx *pulumi.Context, args *LookupMetricsSourceArgs, opts ...pulumi.InvokeOption) (*LookupMetricsSourceResult, error) {
	var rv LookupMetricsSourceResult
	err := ctx.Invoke("azure-native:logz:getMetricsSource", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricsSourceArgs struct {
	MetricsSourceName string `pulumi:"metricsSourceName"`
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}

type LookupMetricsSourceResult struct {
	Id         string                      `pulumi:"id"`
	Identity   *IdentityPropertiesResponse `pulumi:"identity"`
	Location   string                      `pulumi:"location"`
	Name       string                      `pulumi:"name"`
	Properties MonitorPropertiesResponse   `pulumi:"properties"`
	SystemData SystemDataResponse          `pulumi:"systemData"`
	Tags       map[string]string           `pulumi:"tags"`
	Type       string                      `pulumi:"type"`
}

func LookupMetricsSourceOutput(ctx *pulumi.Context, args LookupMetricsSourceOutputArgs, opts ...pulumi.InvokeOption) LookupMetricsSourceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetricsSourceResult, error) {
			args := v.(LookupMetricsSourceArgs)
			r, err := LookupMetricsSource(ctx, &args, opts...)
			var s LookupMetricsSourceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetricsSourceResultOutput)
}

type LookupMetricsSourceOutputArgs struct {
	MetricsSourceName pulumi.StringInput `pulumi:"metricsSourceName"`
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMetricsSourceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceArgs)(nil)).Elem()
}

type LookupMetricsSourceResultOutput struct{ *pulumi.OutputState }

func (LookupMetricsSourceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceResult)(nil)).Elem()
}

func (o LookupMetricsSourceResultOutput) ToLookupMetricsSourceResultOutput() LookupMetricsSourceResultOutput {
	return o
}

func (o LookupMetricsSourceResultOutput) ToLookupMetricsSourceResultOutputWithContext(ctx context.Context) LookupMetricsSourceResultOutput {
	return o
}

func (o LookupMetricsSourceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMetricsSourceResultOutput) Identity() IdentityPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) *IdentityPropertiesResponse { return v.Identity }).(IdentityPropertiesResponsePtrOutput)
}

func (o LookupMetricsSourceResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMetricsSourceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMetricsSourceResultOutput) Properties() MonitorPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) MonitorPropertiesResponse { return v.Properties }).(MonitorPropertiesResponseOutput)
}

func (o LookupMetricsSourceResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMetricsSourceResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMetricsSourceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricsSourceResultOutput{})
}
