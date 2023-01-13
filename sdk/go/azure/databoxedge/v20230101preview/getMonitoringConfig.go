


package v20230101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitoringConfig(ctx *pulumi.Context, args *LookupMonitoringConfigArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringConfigResult, error) {
	var rv LookupMonitoringConfigResult
	err := ctx.Invoke("azure-native:databoxedge/v20230101preview:getMonitoringConfig", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringConfigArgs struct {
	DeviceName        string `pulumi:"deviceName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RoleName          string `pulumi:"roleName"`
}


type LookupMonitoringConfigResult struct {
	Id                   string                        `pulumi:"id"`
	MetricConfigurations []MetricConfigurationResponse `pulumi:"metricConfigurations"`
	Name                 string                        `pulumi:"name"`
	SystemData           SystemDataResponse            `pulumi:"systemData"`
	Type                 string                        `pulumi:"type"`
}

func LookupMonitoringConfigOutput(ctx *pulumi.Context, args LookupMonitoringConfigOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringConfigResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoringConfigResult, error) {
			args := v.(LookupMonitoringConfigArgs)
			r, err := LookupMonitoringConfig(ctx, &args, opts...)
			var s LookupMonitoringConfigResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoringConfigResultOutput)
}

type LookupMonitoringConfigOutputArgs struct {
	DeviceName        pulumi.StringInput `pulumi:"deviceName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RoleName          pulumi.StringInput `pulumi:"roleName"`
}

func (LookupMonitoringConfigOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringConfigArgs)(nil)).Elem()
}


type LookupMonitoringConfigResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringConfigResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringConfigResult)(nil)).Elem()
}

func (o LookupMonitoringConfigResultOutput) ToLookupMonitoringConfigResultOutput() LookupMonitoringConfigResultOutput {
	return o
}

func (o LookupMonitoringConfigResultOutput) ToLookupMonitoringConfigResultOutputWithContext(ctx context.Context) LookupMonitoringConfigResultOutput {
	return o
}

func (o LookupMonitoringConfigResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMonitoringConfigResultOutput) MetricConfigurations() MetricConfigurationResponseArrayOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) []MetricConfigurationResponse { return v.MetricConfigurations }).(MetricConfigurationResponseArrayOutput)
}

func (o LookupMonitoringConfigResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMonitoringConfigResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMonitoringConfigResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringConfigResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringConfigResultOutput{})
}
