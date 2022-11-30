


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitoringSetting(ctx *pulumi.Context, args *LookupMonitoringSettingArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringSettingResult, error) {
	var rv LookupMonitoringSettingResult
	err := ctx.Invoke("azure-native:appplatform/v20210901preview:getMonitoringSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringSettingArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServiceName       string `pulumi:"serviceName"`
}


type LookupMonitoringSettingResult struct {
	Id         string                              `pulumi:"id"`
	Name       string                              `pulumi:"name"`
	Properties MonitoringSettingPropertiesResponse `pulumi:"properties"`
	Type       string                              `pulumi:"type"`
}

func LookupMonitoringSettingOutput(ctx *pulumi.Context, args LookupMonitoringSettingOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoringSettingResult, error) {
			args := v.(LookupMonitoringSettingArgs)
			r, err := LookupMonitoringSetting(ctx, &args, opts...)
			var s LookupMonitoringSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoringSettingResultOutput)
}

type LookupMonitoringSettingOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServiceName       pulumi.StringInput `pulumi:"serviceName"`
}

func (LookupMonitoringSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringSettingArgs)(nil)).Elem()
}


type LookupMonitoringSettingResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringSettingResult)(nil)).Elem()
}

func (o LookupMonitoringSettingResultOutput) ToLookupMonitoringSettingResultOutput() LookupMonitoringSettingResultOutput {
	return o
}

func (o LookupMonitoringSettingResultOutput) ToLookupMonitoringSettingResultOutputWithContext(ctx context.Context) LookupMonitoringSettingResultOutput {
	return o
}

func (o LookupMonitoringSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMonitoringSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMonitoringSettingResultOutput) Properties() MonitoringSettingPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMonitoringSettingResult) MonitoringSettingPropertiesResponse { return v.Properties }).(MonitoringSettingPropertiesResponseOutput)
}

func (o LookupMonitoringSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringSettingResultOutput{})
}
