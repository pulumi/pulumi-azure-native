


package v20220125

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignmentInfoResponse struct {
	Configuration *ConfigurationInfoResponse `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
}

type AssignmentInfoResponseOutput struct{ *pulumi.OutputState }

func (AssignmentInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentInfoResponse)(nil)).Elem()
}

func (o AssignmentInfoResponseOutput) ToAssignmentInfoResponseOutput() AssignmentInfoResponseOutput {
	return o
}

func (o AssignmentInfoResponseOutput) ToAssignmentInfoResponseOutputWithContext(ctx context.Context) AssignmentInfoResponseOutput {
	return o
}

func (o AssignmentInfoResponseOutput) Configuration() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v AssignmentInfoResponse) *ConfigurationInfoResponse { return v.Configuration }).(ConfigurationInfoResponsePtrOutput)
}

func (o AssignmentInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

type AssignmentInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignmentInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentInfoResponse)(nil)).Elem()
}

func (o AssignmentInfoResponsePtrOutput) ToAssignmentInfoResponsePtrOutput() AssignmentInfoResponsePtrOutput {
	return o
}

func (o AssignmentInfoResponsePtrOutput) ToAssignmentInfoResponsePtrOutputWithContext(ctx context.Context) AssignmentInfoResponsePtrOutput {
	return o
}

func (o AssignmentInfoResponsePtrOutput) Elem() AssignmentInfoResponseOutput {
	return o.ApplyT(func(v *AssignmentInfoResponse) AssignmentInfoResponse {
		if v != nil {
			return *v
		}
		var ret AssignmentInfoResponse
		return ret
	}).(AssignmentInfoResponseOutput)
}

func (o AssignmentInfoResponsePtrOutput) Configuration() ConfigurationInfoResponsePtrOutput {
	return o.ApplyT(func(v *AssignmentInfoResponse) *ConfigurationInfoResponse {
		if v == nil {
			return nil
		}
		return v.Configuration
	}).(ConfigurationInfoResponsePtrOutput)
}

func (o AssignmentInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type AssignmentReportResourceComplianceReasonResponse struct {
	Code   string `pulumi:"code"`
	Phrase string `pulumi:"phrase"`
}

type AssignmentReportResourceComplianceReasonResponseOutput struct{ *pulumi.OutputState }

func (AssignmentReportResourceComplianceReasonResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentReportResourceComplianceReasonResponse)(nil)).Elem()
}

func (o AssignmentReportResourceComplianceReasonResponseOutput) ToAssignmentReportResourceComplianceReasonResponseOutput() AssignmentReportResourceComplianceReasonResponseOutput {
	return o
}

func (o AssignmentReportResourceComplianceReasonResponseOutput) ToAssignmentReportResourceComplianceReasonResponseOutputWithContext(ctx context.Context) AssignmentReportResourceComplianceReasonResponseOutput {
	return o
}

func (o AssignmentReportResourceComplianceReasonResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResourceComplianceReasonResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AssignmentReportResourceComplianceReasonResponseOutput) Phrase() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResourceComplianceReasonResponse) string { return v.Phrase }).(pulumi.StringOutput)
}

type AssignmentReportResourceComplianceReasonResponseArrayOutput struct{ *pulumi.OutputState }

func (AssignmentReportResourceComplianceReasonResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentReportResourceComplianceReasonResponse)(nil)).Elem()
}

func (o AssignmentReportResourceComplianceReasonResponseArrayOutput) ToAssignmentReportResourceComplianceReasonResponseArrayOutput() AssignmentReportResourceComplianceReasonResponseArrayOutput {
	return o
}

func (o AssignmentReportResourceComplianceReasonResponseArrayOutput) ToAssignmentReportResourceComplianceReasonResponseArrayOutputWithContext(ctx context.Context) AssignmentReportResourceComplianceReasonResponseArrayOutput {
	return o
}

func (o AssignmentReportResourceComplianceReasonResponseArrayOutput) Index(i pulumi.IntInput) AssignmentReportResourceComplianceReasonResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssignmentReportResourceComplianceReasonResponse {
		return vs[0].([]AssignmentReportResourceComplianceReasonResponse)[vs[1].(int)]
	}).(AssignmentReportResourceComplianceReasonResponseOutput)
}

type AssignmentReportResourceResponse struct {
	ComplianceStatus string                                             `pulumi:"complianceStatus"`
	Properties       interface{}                                        `pulumi:"properties"`
	Reasons          []AssignmentReportResourceComplianceReasonResponse `pulumi:"reasons"`
	ResourceId       string                                             `pulumi:"resourceId"`
}

type AssignmentReportResourceResponseOutput struct{ *pulumi.OutputState }

func (AssignmentReportResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentReportResourceResponse)(nil)).Elem()
}

func (o AssignmentReportResourceResponseOutput) ToAssignmentReportResourceResponseOutput() AssignmentReportResourceResponseOutput {
	return o
}

func (o AssignmentReportResourceResponseOutput) ToAssignmentReportResourceResponseOutputWithContext(ctx context.Context) AssignmentReportResourceResponseOutput {
	return o
}

func (o AssignmentReportResourceResponseOutput) ComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResourceResponse) string { return v.ComplianceStatus }).(pulumi.StringOutput)
}

func (o AssignmentReportResourceResponseOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v AssignmentReportResourceResponse) interface{} { return v.Properties }).(pulumi.AnyOutput)
}

func (o AssignmentReportResourceResponseOutput) Reasons() AssignmentReportResourceComplianceReasonResponseArrayOutput {
	return o.ApplyT(func(v AssignmentReportResourceResponse) []AssignmentReportResourceComplianceReasonResponse {
		return v.Reasons
	}).(AssignmentReportResourceComplianceReasonResponseArrayOutput)
}

func (o AssignmentReportResourceResponseOutput) ResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResourceResponse) string { return v.ResourceId }).(pulumi.StringOutput)
}

type AssignmentReportResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (AssignmentReportResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentReportResourceResponse)(nil)).Elem()
}

func (o AssignmentReportResourceResponseArrayOutput) ToAssignmentReportResourceResponseArrayOutput() AssignmentReportResourceResponseArrayOutput {
	return o
}

func (o AssignmentReportResourceResponseArrayOutput) ToAssignmentReportResourceResponseArrayOutputWithContext(ctx context.Context) AssignmentReportResourceResponseArrayOutput {
	return o
}

func (o AssignmentReportResourceResponseArrayOutput) Index(i pulumi.IntInput) AssignmentReportResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AssignmentReportResourceResponse {
		return vs[0].([]AssignmentReportResourceResponse)[vs[1].(int)]
	}).(AssignmentReportResourceResponseOutput)
}

type AssignmentReportResponse struct {
	Assignment       *AssignmentInfoResponse            `pulumi:"assignment"`
	ComplianceStatus string                             `pulumi:"complianceStatus"`
	EndTime          string                             `pulumi:"endTime"`
	Id               string                             `pulumi:"id"`
	OperationType    string                             `pulumi:"operationType"`
	ReportId         string                             `pulumi:"reportId"`
	Resources        []AssignmentReportResourceResponse `pulumi:"resources"`
	StartTime        string                             `pulumi:"startTime"`
	Vm               *VMInfoResponse                    `pulumi:"vm"`
}

type AssignmentReportResponseOutput struct{ *pulumi.OutputState }

func (AssignmentReportResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentReportResponse)(nil)).Elem()
}

func (o AssignmentReportResponseOutput) ToAssignmentReportResponseOutput() AssignmentReportResponseOutput {
	return o
}

func (o AssignmentReportResponseOutput) ToAssignmentReportResponseOutputWithContext(ctx context.Context) AssignmentReportResponseOutput {
	return o
}

func (o AssignmentReportResponseOutput) Assignment() AssignmentInfoResponsePtrOutput {
	return o.ApplyT(func(v AssignmentReportResponse) *AssignmentInfoResponse { return v.Assignment }).(AssignmentInfoResponsePtrOutput)
}

func (o AssignmentReportResponseOutput) ComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResponse) string { return v.ComplianceStatus }).(pulumi.StringOutput)
}

func (o AssignmentReportResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o AssignmentReportResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AssignmentReportResponseOutput) OperationType() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResponse) string { return v.OperationType }).(pulumi.StringOutput)
}

func (o AssignmentReportResponseOutput) ReportId() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResponse) string { return v.ReportId }).(pulumi.StringOutput)
}

func (o AssignmentReportResponseOutput) Resources() AssignmentReportResourceResponseArrayOutput {
	return o.ApplyT(func(v AssignmentReportResponse) []AssignmentReportResourceResponse { return v.Resources }).(AssignmentReportResourceResponseArrayOutput)
}

func (o AssignmentReportResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v AssignmentReportResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

func (o AssignmentReportResponseOutput) Vm() VMInfoResponsePtrOutput {
	return o.ApplyT(func(v AssignmentReportResponse) *VMInfoResponse { return v.Vm }).(VMInfoResponsePtrOutput)
}

type AssignmentReportResponsePtrOutput struct{ *pulumi.OutputState }

func (AssignmentReportResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentReportResponse)(nil)).Elem()
}

func (o AssignmentReportResponsePtrOutput) ToAssignmentReportResponsePtrOutput() AssignmentReportResponsePtrOutput {
	return o
}

func (o AssignmentReportResponsePtrOutput) ToAssignmentReportResponsePtrOutputWithContext(ctx context.Context) AssignmentReportResponsePtrOutput {
	return o
}

func (o AssignmentReportResponsePtrOutput) Elem() AssignmentReportResponseOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) AssignmentReportResponse {
		if v != nil {
			return *v
		}
		var ret AssignmentReportResponse
		return ret
	}).(AssignmentReportResponseOutput)
}

func (o AssignmentReportResponsePtrOutput) Assignment() AssignmentInfoResponsePtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *AssignmentInfoResponse {
		if v == nil {
			return nil
		}
		return v.Assignment
	}).(AssignmentInfoResponsePtrOutput)
}

func (o AssignmentReportResponsePtrOutput) ComplianceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComplianceStatus
	}).(pulumi.StringPtrOutput)
}

func (o AssignmentReportResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o AssignmentReportResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o AssignmentReportResponsePtrOutput) OperationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OperationType
	}).(pulumi.StringPtrOutput)
}

func (o AssignmentReportResponsePtrOutput) ReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ReportId
	}).(pulumi.StringPtrOutput)
}

func (o AssignmentReportResponsePtrOutput) Resources() AssignmentReportResourceResponseArrayOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) []AssignmentReportResourceResponse {
		if v == nil {
			return nil
		}
		return v.Resources
	}).(AssignmentReportResourceResponseArrayOutput)
}

func (o AssignmentReportResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o AssignmentReportResponsePtrOutput) Vm() VMInfoResponsePtrOutput {
	return o.ApplyT(func(v *AssignmentReportResponse) *VMInfoResponse {
		if v == nil {
			return nil
		}
		return v.Vm
	}).(VMInfoResponsePtrOutput)
}

type ConfigurationInfoResponse struct {
	Name    string `pulumi:"name"`
	Version string `pulumi:"version"`
}

type ConfigurationInfoResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationInfoResponse)(nil)).Elem()
}

func (o ConfigurationInfoResponseOutput) ToConfigurationInfoResponseOutput() ConfigurationInfoResponseOutput {
	return o
}

func (o ConfigurationInfoResponseOutput) ToConfigurationInfoResponseOutputWithContext(ctx context.Context) ConfigurationInfoResponseOutput {
	return o
}

func (o ConfigurationInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationInfoResponseOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationInfoResponse) string { return v.Version }).(pulumi.StringOutput)
}

type ConfigurationInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationInfoResponse)(nil)).Elem()
}

func (o ConfigurationInfoResponsePtrOutput) ToConfigurationInfoResponsePtrOutput() ConfigurationInfoResponsePtrOutput {
	return o
}

func (o ConfigurationInfoResponsePtrOutput) ToConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ConfigurationInfoResponsePtrOutput {
	return o
}

func (o ConfigurationInfoResponsePtrOutput) Elem() ConfigurationInfoResponseOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) ConfigurationInfoResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationInfoResponse
		return ret
	}).(ConfigurationInfoResponseOutput)
}

func (o ConfigurationInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationInfoResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Version
	}).(pulumi.StringPtrOutput)
}

type ConfigurationParameter struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ConfigurationParameterInput interface {
	pulumi.Input

	ToConfigurationParameterOutput() ConfigurationParameterOutput
	ToConfigurationParameterOutputWithContext(context.Context) ConfigurationParameterOutput
}

type ConfigurationParameterArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConfigurationParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameter)(nil)).Elem()
}

func (i ConfigurationParameterArgs) ToConfigurationParameterOutput() ConfigurationParameterOutput {
	return i.ToConfigurationParameterOutputWithContext(context.Background())
}

func (i ConfigurationParameterArgs) ToConfigurationParameterOutputWithContext(ctx context.Context) ConfigurationParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterOutput)
}





type ConfigurationParameterArrayInput interface {
	pulumi.Input

	ToConfigurationParameterArrayOutput() ConfigurationParameterArrayOutput
	ToConfigurationParameterArrayOutputWithContext(context.Context) ConfigurationParameterArrayOutput
}

type ConfigurationParameterArray []ConfigurationParameterInput

func (ConfigurationParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameter)(nil)).Elem()
}

func (i ConfigurationParameterArray) ToConfigurationParameterArrayOutput() ConfigurationParameterArrayOutput {
	return i.ToConfigurationParameterArrayOutputWithContext(context.Background())
}

func (i ConfigurationParameterArray) ToConfigurationParameterArrayOutputWithContext(ctx context.Context) ConfigurationParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterArrayOutput)
}

type ConfigurationParameterOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameter)(nil)).Elem()
}

func (o ConfigurationParameterOutput) ToConfigurationParameterOutput() ConfigurationParameterOutput {
	return o
}

func (o ConfigurationParameterOutput) ToConfigurationParameterOutputWithContext(ctx context.Context) ConfigurationParameterOutput {
	return o
}

func (o ConfigurationParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConfigurationParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConfigurationParameterArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameter)(nil)).Elem()
}

func (o ConfigurationParameterArrayOutput) ToConfigurationParameterArrayOutput() ConfigurationParameterArrayOutput {
	return o
}

func (o ConfigurationParameterArrayOutput) ToConfigurationParameterArrayOutputWithContext(ctx context.Context) ConfigurationParameterArrayOutput {
	return o
}

func (o ConfigurationParameterArrayOutput) Index(i pulumi.IntInput) ConfigurationParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationParameter {
		return vs[0].([]ConfigurationParameter)[vs[1].(int)]
	}).(ConfigurationParameterOutput)
}

type ConfigurationParameterResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type ConfigurationParameterResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameterResponse)(nil)).Elem()
}

func (o ConfigurationParameterResponseOutput) ToConfigurationParameterResponseOutput() ConfigurationParameterResponseOutput {
	return o
}

func (o ConfigurationParameterResponseOutput) ToConfigurationParameterResponseOutputWithContext(ctx context.Context) ConfigurationParameterResponseOutput {
	return o
}

func (o ConfigurationParameterResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameterResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConfigurationParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConfigurationParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameterResponse)(nil)).Elem()
}

func (o ConfigurationParameterResponseArrayOutput) ToConfigurationParameterResponseArrayOutput() ConfigurationParameterResponseArrayOutput {
	return o
}

func (o ConfigurationParameterResponseArrayOutput) ToConfigurationParameterResponseArrayOutputWithContext(ctx context.Context) ConfigurationParameterResponseArrayOutput {
	return o
}

func (o ConfigurationParameterResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationParameterResponse {
		return vs[0].([]ConfigurationParameterResponse)[vs[1].(int)]
	}).(ConfigurationParameterResponseOutput)
}

type ConfigurationSettingResponse struct {
	ActionAfterReboot              string  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           bool    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              string  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins float64 `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 bool    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           float64 `pulumi:"refreshFrequencyMins"`
}


func (val *ConfigurationSettingResponse) Defaults() *ConfigurationSettingResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ConfigurationModeFrequencyMins) {
		tmp.ConfigurationModeFrequencyMins = 15.0
	}
	if isZero(tmp.RefreshFrequencyMins) {
		tmp.RefreshFrequencyMins = 30.0
	}
	return &tmp
}

type ConfigurationSettingResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSettingResponse)(nil)).Elem()
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponseOutput() ConfigurationSettingResponseOutput {
	return o
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponseOutputWithContext(ctx context.Context) ConfigurationSettingResponseOutput {
	return o
}

func (o ConfigurationSettingResponseOutput) ActionAfterReboot() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) string { return v.ActionAfterReboot }).(pulumi.StringOutput)
}

func (o ConfigurationSettingResponseOutput) AllowModuleOverwrite() pulumi.BoolOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) bool { return v.AllowModuleOverwrite }).(pulumi.BoolOutput)
}

func (o ConfigurationSettingResponseOutput) ConfigurationMode() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) string { return v.ConfigurationMode }).(pulumi.StringOutput)
}

func (o ConfigurationSettingResponseOutput) ConfigurationModeFrequencyMins() pulumi.Float64Output {
	return o.ApplyT(func(v ConfigurationSettingResponse) float64 { return v.ConfigurationModeFrequencyMins }).(pulumi.Float64Output)
}

func (o ConfigurationSettingResponseOutput) RebootIfNeeded() pulumi.BoolOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) bool { return v.RebootIfNeeded }).(pulumi.BoolOutput)
}

func (o ConfigurationSettingResponseOutput) RefreshFrequencyMins() pulumi.Float64Output {
	return o.ApplyT(func(v ConfigurationSettingResponse) float64 { return v.RefreshFrequencyMins }).(pulumi.Float64Output)
}

type ConfigurationSettingResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSettingResponse)(nil)).Elem()
}

func (o ConfigurationSettingResponsePtrOutput) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return o
}

func (o ConfigurationSettingResponsePtrOutput) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return o
}

func (o ConfigurationSettingResponsePtrOutput) Elem() ConfigurationSettingResponseOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) ConfigurationSettingResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationSettingResponse
		return ret
	}).(ConfigurationSettingResponseOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionAfterReboot
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.AllowModuleOverwrite
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ConfigurationMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.ConfigurationModeFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RebootIfNeeded
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *float64 {
		if v == nil {
			return nil
		}
		return &v.RefreshFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

type GuestConfigurationAssignmentProperties struct {
	Context            *string                       `pulumi:"context"`
	GuestConfiguration *GuestConfigurationNavigation `pulumi:"guestConfiguration"`
}





type GuestConfigurationAssignmentPropertiesInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesOutput() GuestConfigurationAssignmentPropertiesOutput
	ToGuestConfigurationAssignmentPropertiesOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesOutput
}

type GuestConfigurationAssignmentPropertiesArgs struct {
	Context            pulumi.StringPtrInput                `pulumi:"context"`
	GuestConfiguration GuestConfigurationNavigationPtrInput `pulumi:"guestConfiguration"`
}

func (GuestConfigurationAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesOutput() GuestConfigurationAssignmentPropertiesOutput {
	return i.ToGuestConfigurationAssignmentPropertiesOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesOutput)
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesOutput).ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx)
}









type GuestConfigurationAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput
	ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesPtrOutput
}

type guestConfigurationAssignmentPropertiesPtrType GuestConfigurationAssignmentPropertiesArgs

func GuestConfigurationAssignmentPropertiesPtr(v *GuestConfigurationAssignmentPropertiesArgs) GuestConfigurationAssignmentPropertiesPtrInput {
	return (*guestConfigurationAssignmentPropertiesPtrType)(v)
}

func (*guestConfigurationAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (i *guestConfigurationAssignmentPropertiesPtrType) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *guestConfigurationAssignmentPropertiesPtrType) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesPtrOutput)
}

type GuestConfigurationAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesOutput() GuestConfigurationAssignmentPropertiesOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return o.ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationAssignmentProperties) *GuestConfigurationAssignmentProperties {
		return &v
	}).(GuestConfigurationAssignmentPropertiesPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentProperties) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesOutput) GuestConfiguration() GuestConfigurationNavigationPtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentProperties) *GuestConfigurationNavigation {
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationPtrOutput)
}

type GuestConfigurationAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) Elem() GuestConfigurationAssignmentPropertiesOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentProperties) GuestConfigurationAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationAssignmentProperties
		return ret
	}).(GuestConfigurationAssignmentPropertiesOutput)
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Context
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) GuestConfiguration() GuestConfigurationNavigationPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentProperties) *GuestConfigurationNavigation {
		if v == nil {
			return nil
		}
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationPtrOutput)
}

type GuestConfigurationAssignmentPropertiesResponse struct {
	AssignmentHash              string                                `pulumi:"assignmentHash"`
	ComplianceStatus            string                                `pulumi:"complianceStatus"`
	Context                     *string                               `pulumi:"context"`
	GuestConfiguration          *GuestConfigurationNavigationResponse `pulumi:"guestConfiguration"`
	LastComplianceStatusChecked string                                `pulumi:"lastComplianceStatusChecked"`
	LatestAssignmentReport      *AssignmentReportResponse             `pulumi:"latestAssignmentReport"`
	LatestReportId              string                                `pulumi:"latestReportId"`
	ParameterHash               string                                `pulumi:"parameterHash"`
	ProvisioningState           string                                `pulumi:"provisioningState"`
	ResourceType                string                                `pulumi:"resourceType"`
	TargetResourceId            string                                `pulumi:"targetResourceId"`
	VmssVMList                  []VMSSVMInfoResponse                  `pulumi:"vmssVMList"`
}


func (val *GuestConfigurationAssignmentPropertiesResponse) Defaults() *GuestConfigurationAssignmentPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.GuestConfiguration = tmp.GuestConfiguration.Defaults()

	return &tmp
}

type GuestConfigurationAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponseOutput() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponseOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) AssignmentHash() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.AssignmentHash }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ComplianceStatus }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) GuestConfiguration() GuestConfigurationNavigationResponsePtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) *GuestConfigurationNavigationResponse {
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) LastComplianceStatusChecked() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.LastComplianceStatusChecked }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) LatestAssignmentReport() AssignmentReportResponsePtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) *AssignmentReportResponse {
		return v.LatestAssignmentReport
	}).(AssignmentReportResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) LatestReportId() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.LatestReportId }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ParameterHash() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ParameterHash }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ResourceType }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.TargetResourceId }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) VmssVMList() VMSSVMInfoResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) []VMSSVMInfoResponse { return v.VmssVMList }).(VMSSVMInfoResponseArrayOutput)
}

type GuestConfigurationNavigation struct {
	AssignmentType                  *string                  `pulumi:"assignmentType"`
	ConfigurationParameter          []ConfigurationParameter `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter []ConfigurationParameter `pulumi:"configurationProtectedParameter"`
	ContentHash                     *string                  `pulumi:"contentHash"`
	ContentUri                      *string                  `pulumi:"contentUri"`
	Kind                            *string                  `pulumi:"kind"`
	Name                            *string                  `pulumi:"name"`
	Version                         *string                  `pulumi:"version"`
}





type GuestConfigurationNavigationInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput
	ToGuestConfigurationNavigationOutputWithContext(context.Context) GuestConfigurationNavigationOutput
}

type GuestConfigurationNavigationArgs struct {
	AssignmentType                  pulumi.StringPtrInput            `pulumi:"assignmentType"`
	ConfigurationParameter          ConfigurationParameterArrayInput `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter ConfigurationParameterArrayInput `pulumi:"configurationProtectedParameter"`
	ContentHash                     pulumi.StringPtrInput            `pulumi:"contentHash"`
	ContentUri                      pulumi.StringPtrInput            `pulumi:"contentUri"`
	Kind                            pulumi.StringPtrInput            `pulumi:"kind"`
	Name                            pulumi.StringPtrInput            `pulumi:"name"`
	Version                         pulumi.StringPtrInput            `pulumi:"version"`
}

func (GuestConfigurationNavigationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigation)(nil)).Elem()
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput {
	return i.ToGuestConfigurationNavigationOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationOutputWithContext(ctx context.Context) GuestConfigurationNavigationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationOutput)
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return i.ToGuestConfigurationNavigationPtrOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationOutput).ToGuestConfigurationNavigationPtrOutputWithContext(ctx)
}









type GuestConfigurationNavigationPtrInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput
	ToGuestConfigurationNavigationPtrOutputWithContext(context.Context) GuestConfigurationNavigationPtrOutput
}

type guestConfigurationNavigationPtrType GuestConfigurationNavigationArgs

func GuestConfigurationNavigationPtr(v *GuestConfigurationNavigationArgs) GuestConfigurationNavigationPtrInput {
	return (*guestConfigurationNavigationPtrType)(v)
}

func (*guestConfigurationNavigationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigation)(nil)).Elem()
}

func (i *guestConfigurationNavigationPtrType) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return i.ToGuestConfigurationNavigationPtrOutputWithContext(context.Background())
}

func (i *guestConfigurationNavigationPtrType) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationPtrOutput)
}

type GuestConfigurationNavigationOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigation)(nil)).Elem()
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput {
	return o
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationOutputWithContext(ctx context.Context) GuestConfigurationNavigationOutput {
	return o
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return o.ToGuestConfigurationNavigationPtrOutputWithContext(context.Background())
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationNavigation) *GuestConfigurationNavigation {
		return &v
	}).(GuestConfigurationNavigationPtrOutput)
}

func (o GuestConfigurationNavigationOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.AssignmentType }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) ConfigurationParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) []ConfigurationParameter { return v.ConfigurationParameter }).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationOutput) ConfigurationProtectedParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) []ConfigurationParameter {
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.ContentHash }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.ContentUri }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GuestConfigurationNavigationPtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigation)(nil)).Elem()
}

func (o GuestConfigurationNavigationPtrOutput) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return o
}

func (o GuestConfigurationNavigationPtrOutput) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return o
}

func (o GuestConfigurationNavigationPtrOutput) Elem() GuestConfigurationNavigationOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) GuestConfigurationNavigation {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationNavigation
		return ret
	}).(GuestConfigurationNavigationOutput)
}

func (o GuestConfigurationNavigationPtrOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.AssignmentType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ConfigurationParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) []ConfigurationParameter {
		if v == nil {
			return nil
		}
		return v.ConfigurationParameter
	}).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ConfigurationProtectedParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) []ConfigurationParameter {
		if v == nil {
			return nil
		}
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.ContentUri
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type GuestConfigurationNavigationResponse struct {
	AssignmentSource                string                           `pulumi:"assignmentSource"`
	AssignmentType                  *string                          `pulumi:"assignmentType"`
	ConfigurationParameter          []ConfigurationParameterResponse `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter []ConfigurationParameterResponse `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            ConfigurationSettingResponse     `pulumi:"configurationSetting"`
	ContentHash                     *string                          `pulumi:"contentHash"`
	ContentType                     string                           `pulumi:"contentType"`
	ContentUri                      *string                          `pulumi:"contentUri"`
	Kind                            *string                          `pulumi:"kind"`
	Name                            *string                          `pulumi:"name"`
	Version                         *string                          `pulumi:"version"`
}


func (val *GuestConfigurationNavigationResponse) Defaults() *GuestConfigurationNavigationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConfigurationSetting = *tmp.ConfigurationSetting.Defaults()

	return &tmp
}

type GuestConfigurationNavigationResponseOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponseOutput() GuestConfigurationNavigationResponseOutput {
	return o
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponseOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponseOutput {
	return o
}

func (o GuestConfigurationNavigationResponseOutput) AssignmentSource() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) string { return v.AssignmentSource }).(pulumi.StringOutput)
}

func (o GuestConfigurationNavigationResponseOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.AssignmentType }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		return v.ConfigurationParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationProtectedParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationSetting() ConfigurationSettingResponseOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) ConfigurationSettingResponse {
		return v.ConfigurationSetting
	}).(ConfigurationSettingResponseOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.ContentHash }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.ContentUri }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GuestConfigurationNavigationResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (o GuestConfigurationNavigationResponsePtrOutput) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return o
}

func (o GuestConfigurationNavigationResponsePtrOutput) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return o
}

func (o GuestConfigurationNavigationResponsePtrOutput) Elem() GuestConfigurationNavigationResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) GuestConfigurationNavigationResponse {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationNavigationResponse
		return ret
	}).(GuestConfigurationNavigationResponseOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) AssignmentSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssignmentSource
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssignmentType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationProtectedParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationSetting() ConfigurationSettingResponsePtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *ConfigurationSettingResponse {
		if v == nil {
			return nil
		}
		return &v.ConfigurationSetting
	}).(ConfigurationSettingResponsePtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentUri
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
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

type VMInfoResponse struct {
	Id   string `pulumi:"id"`
	Uuid string `pulumi:"uuid"`
}

type VMInfoResponseOutput struct{ *pulumi.OutputState }

func (VMInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMInfoResponse)(nil)).Elem()
}

func (o VMInfoResponseOutput) ToVMInfoResponseOutput() VMInfoResponseOutput {
	return o
}

func (o VMInfoResponseOutput) ToVMInfoResponseOutputWithContext(ctx context.Context) VMInfoResponseOutput {
	return o
}

func (o VMInfoResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VMInfoResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VMInfoResponseOutput) Uuid() pulumi.StringOutput {
	return o.ApplyT(func(v VMInfoResponse) string { return v.Uuid }).(pulumi.StringOutput)
}

type VMInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (VMInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VMInfoResponse)(nil)).Elem()
}

func (o VMInfoResponsePtrOutput) ToVMInfoResponsePtrOutput() VMInfoResponsePtrOutput {
	return o
}

func (o VMInfoResponsePtrOutput) ToVMInfoResponsePtrOutputWithContext(ctx context.Context) VMInfoResponsePtrOutput {
	return o
}

func (o VMInfoResponsePtrOutput) Elem() VMInfoResponseOutput {
	return o.ApplyT(func(v *VMInfoResponse) VMInfoResponse {
		if v != nil {
			return *v
		}
		var ret VMInfoResponse
		return ret
	}).(VMInfoResponseOutput)
}

func (o VMInfoResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VMInfoResponsePtrOutput) Uuid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VMInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Uuid
	}).(pulumi.StringPtrOutput)
}

type VMSSVMInfoResponse struct {
	ComplianceStatus      string `pulumi:"complianceStatus"`
	LastComplianceChecked string `pulumi:"lastComplianceChecked"`
	LatestReportId        string `pulumi:"latestReportId"`
	VmId                  string `pulumi:"vmId"`
	VmResourceId          string `pulumi:"vmResourceId"`
}

type VMSSVMInfoResponseOutput struct{ *pulumi.OutputState }

func (VMSSVMInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSVMInfoResponse)(nil)).Elem()
}

func (o VMSSVMInfoResponseOutput) ToVMSSVMInfoResponseOutput() VMSSVMInfoResponseOutput {
	return o
}

func (o VMSSVMInfoResponseOutput) ToVMSSVMInfoResponseOutputWithContext(ctx context.Context) VMSSVMInfoResponseOutput {
	return o
}

func (o VMSSVMInfoResponseOutput) ComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.ComplianceStatus }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) LastComplianceChecked() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.LastComplianceChecked }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) LatestReportId() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.LatestReportId }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.VmId }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) VmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.VmResourceId }).(pulumi.StringOutput)
}

type VMSSVMInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (VMSSVMInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSVMInfoResponse)(nil)).Elem()
}

func (o VMSSVMInfoResponseArrayOutput) ToVMSSVMInfoResponseArrayOutput() VMSSVMInfoResponseArrayOutput {
	return o
}

func (o VMSSVMInfoResponseArrayOutput) ToVMSSVMInfoResponseArrayOutputWithContext(ctx context.Context) VMSSVMInfoResponseArrayOutput {
	return o
}

func (o VMSSVMInfoResponseArrayOutput) Index(i pulumi.IntInput) VMSSVMInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSVMInfoResponse {
		return vs[0].([]VMSSVMInfoResponse)[vs[1].(int)]
	}).(VMSSVMInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AssignmentInfoResponseOutput{})
	pulumi.RegisterOutputType(AssignmentInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AssignmentReportResourceComplianceReasonResponseOutput{})
	pulumi.RegisterOutputType(AssignmentReportResourceComplianceReasonResponseArrayOutput{})
	pulumi.RegisterOutputType(AssignmentReportResourceResponseOutput{})
	pulumi.RegisterOutputType(AssignmentReportResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(AssignmentReportResponseOutput{})
	pulumi.RegisterOutputType(AssignmentReportResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationInfoResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(VMInfoResponseOutput{})
	pulumi.RegisterOutputType(VMInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(VMSSVMInfoResponseOutput{})
	pulumi.RegisterOutputType(VMSSVMInfoResponseArrayOutput{})
}
