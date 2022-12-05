


package insights

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupDiagnosticSetting(ctx *pulumi.Context, args *LookupDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupDiagnosticSettingResult, error) {
	var rv LookupDiagnosticSettingResult
	err := ctx.Invoke("azure-native:insights:getDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupDiagnosticSettingArgs struct {
	Name        string `pulumi:"name"`
	ResourceUri string `pulumi:"resourceUri"`
}


type LookupDiagnosticSettingResult struct {
	EventHubAuthorizationRuleId *string                  `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                  `pulumi:"eventHubName"`
	Id                          string                   `pulumi:"id"`
	LogAnalyticsDestinationType *string                  `pulumi:"logAnalyticsDestinationType"`
	Logs                        []LogSettingsResponse    `pulumi:"logs"`
	Metrics                     []MetricSettingsResponse `pulumi:"metrics"`
	Name                        string                   `pulumi:"name"`
	ServiceBusRuleId            *string                  `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                  `pulumi:"storageAccountId"`
	Type                        string                   `pulumi:"type"`
	WorkspaceId                 *string                  `pulumi:"workspaceId"`
}

func LookupDiagnosticSettingOutput(ctx *pulumi.Context, args LookupDiagnosticSettingOutputArgs, opts ...pulumi.InvokeOption) LookupDiagnosticSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupDiagnosticSettingResult, error) {
			args := v.(LookupDiagnosticSettingArgs)
			r, err := LookupDiagnosticSetting(ctx, &args, opts...)
			var s LookupDiagnosticSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupDiagnosticSettingResultOutput)
}

type LookupDiagnosticSettingOutputArgs struct {
	Name        pulumi.StringInput `pulumi:"name"`
	ResourceUri pulumi.StringInput `pulumi:"resourceUri"`
}

func (LookupDiagnosticSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticSettingArgs)(nil)).Elem()
}


type LookupDiagnosticSettingResultOutput struct{ *pulumi.OutputState }

func (LookupDiagnosticSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupDiagnosticSettingResult)(nil)).Elem()
}

func (o LookupDiagnosticSettingResultOutput) ToLookupDiagnosticSettingResultOutput() LookupDiagnosticSettingResultOutput {
	return o
}

func (o LookupDiagnosticSettingResultOutput) ToLookupDiagnosticSettingResultOutputWithContext(ctx context.Context) LookupDiagnosticSettingResultOutput {
	return o
}

func (o LookupDiagnosticSettingResultOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticSettingResultOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupDiagnosticSettingResultOutput) LogAnalyticsDestinationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.LogAnalyticsDestinationType }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticSettingResultOutput) Logs() LogSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) []LogSettingsResponse { return v.Logs }).(LogSettingsResponseArrayOutput)
}

func (o LookupDiagnosticSettingResultOutput) Metrics() MetricSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) []MetricSettingsResponse { return v.Metrics }).(MetricSettingsResponseArrayOutput)
}

func (o LookupDiagnosticSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupDiagnosticSettingResultOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticSettingResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupDiagnosticSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupDiagnosticSettingResultOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupDiagnosticSettingResult) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupDiagnosticSettingResultOutput{})
}
