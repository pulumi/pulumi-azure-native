


package v20170501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubscriptionDiagnosticSetting(ctx *pulumi.Context, args *LookupSubscriptionDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupSubscriptionDiagnosticSettingResult, error) {
	var rv LookupSubscriptionDiagnosticSettingResult
	err := ctx.Invoke("azure-native:insights/v20170501preview:getSubscriptionDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubscriptionDiagnosticSettingArgs struct {
	Name string `pulumi:"name"`
}


type LookupSubscriptionDiagnosticSettingResult struct {
	EventHubAuthorizationRuleId *string                           `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                           `pulumi:"eventHubName"`
	Id                          string                            `pulumi:"id"`
	Location                    *string                           `pulumi:"location"`
	Logs                        []SubscriptionLogSettingsResponse `pulumi:"logs"`
	Name                        string                            `pulumi:"name"`
	ServiceBusRuleId            *string                           `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                           `pulumi:"storageAccountId"`
	Type                        string                            `pulumi:"type"`
	WorkspaceId                 *string                           `pulumi:"workspaceId"`
}

func LookupSubscriptionDiagnosticSettingOutput(ctx *pulumi.Context, args LookupSubscriptionDiagnosticSettingOutputArgs, opts ...pulumi.InvokeOption) LookupSubscriptionDiagnosticSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubscriptionDiagnosticSettingResult, error) {
			args := v.(LookupSubscriptionDiagnosticSettingArgs)
			r, err := LookupSubscriptionDiagnosticSetting(ctx, &args, opts...)
			var s LookupSubscriptionDiagnosticSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubscriptionDiagnosticSettingResultOutput)
}

type LookupSubscriptionDiagnosticSettingOutputArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (LookupSubscriptionDiagnosticSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionDiagnosticSettingArgs)(nil)).Elem()
}


type LookupSubscriptionDiagnosticSettingResultOutput struct{ *pulumi.OutputState }

func (LookupSubscriptionDiagnosticSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubscriptionDiagnosticSettingResult)(nil)).Elem()
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) ToLookupSubscriptionDiagnosticSettingResultOutput() LookupSubscriptionDiagnosticSettingResultOutput {
	return o
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) ToLookupSubscriptionDiagnosticSettingResultOutputWithContext(ctx context.Context) LookupSubscriptionDiagnosticSettingResultOutput {
	return o
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) *string { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) Logs() SubscriptionLogSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) []SubscriptionLogSettingsResponse { return v.Logs }).(SubscriptionLogSettingsResponseArrayOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) *string { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupSubscriptionDiagnosticSettingResultOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSubscriptionDiagnosticSettingResult) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubscriptionDiagnosticSettingResultOutput{})
}
