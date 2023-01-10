


package v20200101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementGroupDiagnosticSetting(ctx *pulumi.Context, args *LookupManagementGroupDiagnosticSettingArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupDiagnosticSettingResult, error) {
	var rv LookupManagementGroupDiagnosticSettingResult
	err := ctx.Invoke("azure-native:insights/v20200101preview:getManagementGroupDiagnosticSetting", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupDiagnosticSettingArgs struct {
	ManagementGroupId string `pulumi:"managementGroupId"`
	Name              string `pulumi:"name"`
}


type LookupManagementGroupDiagnosticSettingResult struct {
	EventHubAuthorizationRuleId *string                              `pulumi:"eventHubAuthorizationRuleId"`
	EventHubName                *string                              `pulumi:"eventHubName"`
	Id                          string                               `pulumi:"id"`
	Location                    *string                              `pulumi:"location"`
	Logs                        []ManagementGroupLogSettingsResponse `pulumi:"logs"`
	Name                        string                               `pulumi:"name"`
	ServiceBusRuleId            *string                              `pulumi:"serviceBusRuleId"`
	StorageAccountId            *string                              `pulumi:"storageAccountId"`
	Type                        string                               `pulumi:"type"`
	WorkspaceId                 *string                              `pulumi:"workspaceId"`
}

func LookupManagementGroupDiagnosticSettingOutput(ctx *pulumi.Context, args LookupManagementGroupDiagnosticSettingOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupDiagnosticSettingResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupDiagnosticSettingResult, error) {
			args := v.(LookupManagementGroupDiagnosticSettingArgs)
			r, err := LookupManagementGroupDiagnosticSetting(ctx, &args, opts...)
			var s LookupManagementGroupDiagnosticSettingResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupDiagnosticSettingResultOutput)
}

type LookupManagementGroupDiagnosticSettingOutputArgs struct {
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	Name              pulumi.StringInput `pulumi:"name"`
}

func (LookupManagementGroupDiagnosticSettingOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupDiagnosticSettingArgs)(nil)).Elem()
}


type LookupManagementGroupDiagnosticSettingResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupDiagnosticSettingResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupDiagnosticSettingResult)(nil)).Elem()
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) ToLookupManagementGroupDiagnosticSettingResultOutput() LookupManagementGroupDiagnosticSettingResultOutput {
	return o
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) ToLookupManagementGroupDiagnosticSettingResultOutputWithContext(ctx context.Context) LookupManagementGroupDiagnosticSettingResultOutput {
	return o
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) EventHubAuthorizationRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.EventHubAuthorizationRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) EventHubName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.EventHubName }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) Logs() ManagementGroupLogSettingsResponseArrayOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) []ManagementGroupLogSettingsResponse {
		return v.Logs
	}).(ManagementGroupLogSettingsResponseArrayOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) ServiceBusRuleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.ServiceBusRuleId }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) StorageAccountId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.StorageAccountId }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupManagementGroupDiagnosticSettingResultOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupDiagnosticSettingResult) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupDiagnosticSettingResultOutput{})
}
