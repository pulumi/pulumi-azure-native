


package monitor

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MonitoringAccountResponseDefaultIngestionSettings struct {
	DataCollectionEndpointResourceId string `pulumi:"dataCollectionEndpointResourceId"`
	DataCollectionRuleResourceId     string `pulumi:"dataCollectionRuleResourceId"`
}

type MonitoringAccountResponseDefaultIngestionSettingsOutput struct{ *pulumi.OutputState }

func (MonitoringAccountResponseDefaultIngestionSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringAccountResponseDefaultIngestionSettings)(nil)).Elem()
}

func (o MonitoringAccountResponseDefaultIngestionSettingsOutput) ToMonitoringAccountResponseDefaultIngestionSettingsOutput() MonitoringAccountResponseDefaultIngestionSettingsOutput {
	return o
}

func (o MonitoringAccountResponseDefaultIngestionSettingsOutput) ToMonitoringAccountResponseDefaultIngestionSettingsOutputWithContext(ctx context.Context) MonitoringAccountResponseDefaultIngestionSettingsOutput {
	return o
}

func (o MonitoringAccountResponseDefaultIngestionSettingsOutput) DataCollectionEndpointResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MonitoringAccountResponseDefaultIngestionSettings) string {
		return v.DataCollectionEndpointResourceId
	}).(pulumi.StringOutput)
}

func (o MonitoringAccountResponseDefaultIngestionSettingsOutput) DataCollectionRuleResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v MonitoringAccountResponseDefaultIngestionSettings) string {
		return v.DataCollectionRuleResourceId
	}).(pulumi.StringOutput)
}

type MonitoringAccountResponseMetrics struct {
	InternalId              string `pulumi:"internalId"`
	PrometheusQueryEndpoint string `pulumi:"prometheusQueryEndpoint"`
}

type MonitoringAccountResponseMetricsOutput struct{ *pulumi.OutputState }

func (MonitoringAccountResponseMetricsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonitoringAccountResponseMetrics)(nil)).Elem()
}

func (o MonitoringAccountResponseMetricsOutput) ToMonitoringAccountResponseMetricsOutput() MonitoringAccountResponseMetricsOutput {
	return o
}

func (o MonitoringAccountResponseMetricsOutput) ToMonitoringAccountResponseMetricsOutputWithContext(ctx context.Context) MonitoringAccountResponseMetricsOutput {
	return o
}

func (o MonitoringAccountResponseMetricsOutput) InternalId() pulumi.StringOutput {
	return o.ApplyT(func(v MonitoringAccountResponseMetrics) string { return v.InternalId }).(pulumi.StringOutput)
}

func (o MonitoringAccountResponseMetricsOutput) PrometheusQueryEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v MonitoringAccountResponseMetrics) string { return v.PrometheusQueryEndpoint }).(pulumi.StringOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(MonitoringAccountResponseDefaultIngestionSettingsOutput{})
	pulumi.RegisterOutputType(MonitoringAccountResponseMetricsOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
