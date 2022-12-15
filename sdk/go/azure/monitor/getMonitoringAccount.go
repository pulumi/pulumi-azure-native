


package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMonitoringAccount(ctx *pulumi.Context, args *LookupMonitoringAccountArgs, opts ...pulumi.InvokeOption) (*LookupMonitoringAccountResult, error) {
	var rv LookupMonitoringAccountResult
	err := ctx.Invoke("azure-native:monitor:getMonitoringAccount", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMonitoringAccountArgs struct {
	MonitoringAccountName string `pulumi:"monitoringAccountName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupMonitoringAccountResult struct {
	AccountId                string                                            `pulumi:"accountId"`
	DefaultIngestionSettings MonitoringAccountResponseDefaultIngestionSettings `pulumi:"defaultIngestionSettings"`
	Etag                     string                                            `pulumi:"etag"`
	Id                       string                                            `pulumi:"id"`
	Location                 string                                            `pulumi:"location"`
	Metrics                  MonitoringAccountResponseMetrics                  `pulumi:"metrics"`
	Name                     string                                            `pulumi:"name"`
	ProvisioningState        string                                            `pulumi:"provisioningState"`
	SystemData               SystemDataResponse                                `pulumi:"systemData"`
	Tags                     map[string]string                                 `pulumi:"tags"`
	Type                     string                                            `pulumi:"type"`
}

func LookupMonitoringAccountOutput(ctx *pulumi.Context, args LookupMonitoringAccountOutputArgs, opts ...pulumi.InvokeOption) LookupMonitoringAccountResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMonitoringAccountResult, error) {
			args := v.(LookupMonitoringAccountArgs)
			r, err := LookupMonitoringAccount(ctx, &args, opts...)
			var s LookupMonitoringAccountResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMonitoringAccountResultOutput)
}

type LookupMonitoringAccountOutputArgs struct {
	MonitoringAccountName pulumi.StringInput `pulumi:"monitoringAccountName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupMonitoringAccountOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringAccountArgs)(nil)).Elem()
}


type LookupMonitoringAccountResultOutput struct{ *pulumi.OutputState }

func (LookupMonitoringAccountResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMonitoringAccountResult)(nil)).Elem()
}

func (o LookupMonitoringAccountResultOutput) ToLookupMonitoringAccountResultOutput() LookupMonitoringAccountResultOutput {
	return o
}

func (o LookupMonitoringAccountResultOutput) ToLookupMonitoringAccountResultOutputWithContext(ctx context.Context) LookupMonitoringAccountResultOutput {
	return o
}

func (o LookupMonitoringAccountResultOutput) AccountId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.AccountId }).(pulumi.StringOutput)
}

func (o LookupMonitoringAccountResultOutput) DefaultIngestionSettings() MonitoringAccountResponseDefaultIngestionSettingsOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) MonitoringAccountResponseDefaultIngestionSettings {
		return v.DefaultIngestionSettings
	}).(MonitoringAccountResponseDefaultIngestionSettingsOutput)
}

func (o LookupMonitoringAccountResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupMonitoringAccountResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMonitoringAccountResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupMonitoringAccountResultOutput) Metrics() MonitoringAccountResponseMetricsOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) MonitoringAccountResponseMetrics { return v.Metrics }).(MonitoringAccountResponseMetricsOutput)
}

func (o LookupMonitoringAccountResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMonitoringAccountResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupMonitoringAccountResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMonitoringAccountResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupMonitoringAccountResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMonitoringAccountResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMonitoringAccountResultOutput{})
}
