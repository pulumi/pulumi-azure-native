


package v20210125

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AssignmentInfoResponse struct {
	Configuration *ConfigurationInfoResponse `pulumi:"configuration"`
	Name          string                     `pulumi:"name"`
}





type AssignmentInfoResponseInput interface {
	pulumi.Input

	ToAssignmentInfoResponseOutput() AssignmentInfoResponseOutput
	ToAssignmentInfoResponseOutputWithContext(context.Context) AssignmentInfoResponseOutput
}

type AssignmentInfoResponseArgs struct {
	Configuration ConfigurationInfoResponsePtrInput `pulumi:"configuration"`
	Name          pulumi.StringInput                `pulumi:"name"`
}

func (AssignmentInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentInfoResponse)(nil)).Elem()
}

func (i AssignmentInfoResponseArgs) ToAssignmentInfoResponseOutput() AssignmentInfoResponseOutput {
	return i.ToAssignmentInfoResponseOutputWithContext(context.Background())
}

func (i AssignmentInfoResponseArgs) ToAssignmentInfoResponseOutputWithContext(ctx context.Context) AssignmentInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentInfoResponseOutput)
}

func (i AssignmentInfoResponseArgs) ToAssignmentInfoResponsePtrOutput() AssignmentInfoResponsePtrOutput {
	return i.ToAssignmentInfoResponsePtrOutputWithContext(context.Background())
}

func (i AssignmentInfoResponseArgs) ToAssignmentInfoResponsePtrOutputWithContext(ctx context.Context) AssignmentInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentInfoResponseOutput).ToAssignmentInfoResponsePtrOutputWithContext(ctx)
}









type AssignmentInfoResponsePtrInput interface {
	pulumi.Input

	ToAssignmentInfoResponsePtrOutput() AssignmentInfoResponsePtrOutput
	ToAssignmentInfoResponsePtrOutputWithContext(context.Context) AssignmentInfoResponsePtrOutput
}

type assignmentInfoResponsePtrType AssignmentInfoResponseArgs

func AssignmentInfoResponsePtr(v *AssignmentInfoResponseArgs) AssignmentInfoResponsePtrInput {
	return (*assignmentInfoResponsePtrType)(v)
}

func (*assignmentInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentInfoResponse)(nil)).Elem()
}

func (i *assignmentInfoResponsePtrType) ToAssignmentInfoResponsePtrOutput() AssignmentInfoResponsePtrOutput {
	return i.ToAssignmentInfoResponsePtrOutputWithContext(context.Background())
}

func (i *assignmentInfoResponsePtrType) ToAssignmentInfoResponsePtrOutputWithContext(ctx context.Context) AssignmentInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentInfoResponsePtrOutput)
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

func (o AssignmentInfoResponseOutput) ToAssignmentInfoResponsePtrOutput() AssignmentInfoResponsePtrOutput {
	return o.ToAssignmentInfoResponsePtrOutputWithContext(context.Background())
}

func (o AssignmentInfoResponseOutput) ToAssignmentInfoResponsePtrOutputWithContext(ctx context.Context) AssignmentInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentInfoResponse) *AssignmentInfoResponse {
		return &v
	}).(AssignmentInfoResponsePtrOutput)
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





type AssignmentReportResourceComplianceReasonResponseInput interface {
	pulumi.Input

	ToAssignmentReportResourceComplianceReasonResponseOutput() AssignmentReportResourceComplianceReasonResponseOutput
	ToAssignmentReportResourceComplianceReasonResponseOutputWithContext(context.Context) AssignmentReportResourceComplianceReasonResponseOutput
}

type AssignmentReportResourceComplianceReasonResponseArgs struct {
	Code   pulumi.StringInput `pulumi:"code"`
	Phrase pulumi.StringInput `pulumi:"phrase"`
}

func (AssignmentReportResourceComplianceReasonResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentReportResourceComplianceReasonResponse)(nil)).Elem()
}

func (i AssignmentReportResourceComplianceReasonResponseArgs) ToAssignmentReportResourceComplianceReasonResponseOutput() AssignmentReportResourceComplianceReasonResponseOutput {
	return i.ToAssignmentReportResourceComplianceReasonResponseOutputWithContext(context.Background())
}

func (i AssignmentReportResourceComplianceReasonResponseArgs) ToAssignmentReportResourceComplianceReasonResponseOutputWithContext(ctx context.Context) AssignmentReportResourceComplianceReasonResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResourceComplianceReasonResponseOutput)
}





type AssignmentReportResourceComplianceReasonResponseArrayInput interface {
	pulumi.Input

	ToAssignmentReportResourceComplianceReasonResponseArrayOutput() AssignmentReportResourceComplianceReasonResponseArrayOutput
	ToAssignmentReportResourceComplianceReasonResponseArrayOutputWithContext(context.Context) AssignmentReportResourceComplianceReasonResponseArrayOutput
}

type AssignmentReportResourceComplianceReasonResponseArray []AssignmentReportResourceComplianceReasonResponseInput

func (AssignmentReportResourceComplianceReasonResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentReportResourceComplianceReasonResponse)(nil)).Elem()
}

func (i AssignmentReportResourceComplianceReasonResponseArray) ToAssignmentReportResourceComplianceReasonResponseArrayOutput() AssignmentReportResourceComplianceReasonResponseArrayOutput {
	return i.ToAssignmentReportResourceComplianceReasonResponseArrayOutputWithContext(context.Background())
}

func (i AssignmentReportResourceComplianceReasonResponseArray) ToAssignmentReportResourceComplianceReasonResponseArrayOutputWithContext(ctx context.Context) AssignmentReportResourceComplianceReasonResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResourceComplianceReasonResponseArrayOutput)
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





type AssignmentReportResourceResponseInput interface {
	pulumi.Input

	ToAssignmentReportResourceResponseOutput() AssignmentReportResourceResponseOutput
	ToAssignmentReportResourceResponseOutputWithContext(context.Context) AssignmentReportResourceResponseOutput
}

type AssignmentReportResourceResponseArgs struct {
	ComplianceStatus pulumi.StringInput                                         `pulumi:"complianceStatus"`
	Properties       pulumi.Input                                               `pulumi:"properties"`
	Reasons          AssignmentReportResourceComplianceReasonResponseArrayInput `pulumi:"reasons"`
	ResourceId       pulumi.StringInput                                         `pulumi:"resourceId"`
}

func (AssignmentReportResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentReportResourceResponse)(nil)).Elem()
}

func (i AssignmentReportResourceResponseArgs) ToAssignmentReportResourceResponseOutput() AssignmentReportResourceResponseOutput {
	return i.ToAssignmentReportResourceResponseOutputWithContext(context.Background())
}

func (i AssignmentReportResourceResponseArgs) ToAssignmentReportResourceResponseOutputWithContext(ctx context.Context) AssignmentReportResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResourceResponseOutput)
}





type AssignmentReportResourceResponseArrayInput interface {
	pulumi.Input

	ToAssignmentReportResourceResponseArrayOutput() AssignmentReportResourceResponseArrayOutput
	ToAssignmentReportResourceResponseArrayOutputWithContext(context.Context) AssignmentReportResourceResponseArrayOutput
}

type AssignmentReportResourceResponseArray []AssignmentReportResourceResponseInput

func (AssignmentReportResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AssignmentReportResourceResponse)(nil)).Elem()
}

func (i AssignmentReportResourceResponseArray) ToAssignmentReportResourceResponseArrayOutput() AssignmentReportResourceResponseArrayOutput {
	return i.ToAssignmentReportResourceResponseArrayOutputWithContext(context.Background())
}

func (i AssignmentReportResourceResponseArray) ToAssignmentReportResourceResponseArrayOutputWithContext(ctx context.Context) AssignmentReportResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResourceResponseArrayOutput)
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





type AssignmentReportResponseInput interface {
	pulumi.Input

	ToAssignmentReportResponseOutput() AssignmentReportResponseOutput
	ToAssignmentReportResponseOutputWithContext(context.Context) AssignmentReportResponseOutput
}

type AssignmentReportResponseArgs struct {
	Assignment       AssignmentInfoResponsePtrInput             `pulumi:"assignment"`
	ComplianceStatus pulumi.StringInput                         `pulumi:"complianceStatus"`
	EndTime          pulumi.StringInput                         `pulumi:"endTime"`
	Id               pulumi.StringInput                         `pulumi:"id"`
	OperationType    pulumi.StringInput                         `pulumi:"operationType"`
	ReportId         pulumi.StringInput                         `pulumi:"reportId"`
	Resources        AssignmentReportResourceResponseArrayInput `pulumi:"resources"`
	StartTime        pulumi.StringInput                         `pulumi:"startTime"`
	Vm               VMInfoResponsePtrInput                     `pulumi:"vm"`
}

func (AssignmentReportResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssignmentReportResponse)(nil)).Elem()
}

func (i AssignmentReportResponseArgs) ToAssignmentReportResponseOutput() AssignmentReportResponseOutput {
	return i.ToAssignmentReportResponseOutputWithContext(context.Background())
}

func (i AssignmentReportResponseArgs) ToAssignmentReportResponseOutputWithContext(ctx context.Context) AssignmentReportResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResponseOutput)
}

func (i AssignmentReportResponseArgs) ToAssignmentReportResponsePtrOutput() AssignmentReportResponsePtrOutput {
	return i.ToAssignmentReportResponsePtrOutputWithContext(context.Background())
}

func (i AssignmentReportResponseArgs) ToAssignmentReportResponsePtrOutputWithContext(ctx context.Context) AssignmentReportResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResponseOutput).ToAssignmentReportResponsePtrOutputWithContext(ctx)
}









type AssignmentReportResponsePtrInput interface {
	pulumi.Input

	ToAssignmentReportResponsePtrOutput() AssignmentReportResponsePtrOutput
	ToAssignmentReportResponsePtrOutputWithContext(context.Context) AssignmentReportResponsePtrOutput
}

type assignmentReportResponsePtrType AssignmentReportResponseArgs

func AssignmentReportResponsePtr(v *AssignmentReportResponseArgs) AssignmentReportResponsePtrInput {
	return (*assignmentReportResponsePtrType)(v)
}

func (*assignmentReportResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssignmentReportResponse)(nil)).Elem()
}

func (i *assignmentReportResponsePtrType) ToAssignmentReportResponsePtrOutput() AssignmentReportResponsePtrOutput {
	return i.ToAssignmentReportResponsePtrOutputWithContext(context.Background())
}

func (i *assignmentReportResponsePtrType) ToAssignmentReportResponsePtrOutputWithContext(ctx context.Context) AssignmentReportResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssignmentReportResponsePtrOutput)
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

func (o AssignmentReportResponseOutput) ToAssignmentReportResponsePtrOutput() AssignmentReportResponsePtrOutput {
	return o.ToAssignmentReportResponsePtrOutputWithContext(context.Background())
}

func (o AssignmentReportResponseOutput) ToAssignmentReportResponsePtrOutputWithContext(ctx context.Context) AssignmentReportResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssignmentReportResponse) *AssignmentReportResponse {
		return &v
	}).(AssignmentReportResponsePtrOutput)
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





type ConfigurationInfoResponseInput interface {
	pulumi.Input

	ToConfigurationInfoResponseOutput() ConfigurationInfoResponseOutput
	ToConfigurationInfoResponseOutputWithContext(context.Context) ConfigurationInfoResponseOutput
}

type ConfigurationInfoResponseArgs struct {
	Name    pulumi.StringInput `pulumi:"name"`
	Version pulumi.StringInput `pulumi:"version"`
}

func (ConfigurationInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationInfoResponse)(nil)).Elem()
}

func (i ConfigurationInfoResponseArgs) ToConfigurationInfoResponseOutput() ConfigurationInfoResponseOutput {
	return i.ToConfigurationInfoResponseOutputWithContext(context.Background())
}

func (i ConfigurationInfoResponseArgs) ToConfigurationInfoResponseOutputWithContext(ctx context.Context) ConfigurationInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationInfoResponseOutput)
}

func (i ConfigurationInfoResponseArgs) ToConfigurationInfoResponsePtrOutput() ConfigurationInfoResponsePtrOutput {
	return i.ToConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationInfoResponseArgs) ToConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationInfoResponseOutput).ToConfigurationInfoResponsePtrOutputWithContext(ctx)
}









type ConfigurationInfoResponsePtrInput interface {
	pulumi.Input

	ToConfigurationInfoResponsePtrOutput() ConfigurationInfoResponsePtrOutput
	ToConfigurationInfoResponsePtrOutputWithContext(context.Context) ConfigurationInfoResponsePtrOutput
}

type configurationInfoResponsePtrType ConfigurationInfoResponseArgs

func ConfigurationInfoResponsePtr(v *ConfigurationInfoResponseArgs) ConfigurationInfoResponsePtrInput {
	return (*configurationInfoResponsePtrType)(v)
}

func (*configurationInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationInfoResponse)(nil)).Elem()
}

func (i *configurationInfoResponsePtrType) ToConfigurationInfoResponsePtrOutput() ConfigurationInfoResponsePtrOutput {
	return i.ToConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (i *configurationInfoResponsePtrType) ToConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ConfigurationInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationInfoResponsePtrOutput)
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

func (o ConfigurationInfoResponseOutput) ToConfigurationInfoResponsePtrOutput() ConfigurationInfoResponsePtrOutput {
	return o.ToConfigurationInfoResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationInfoResponseOutput) ToConfigurationInfoResponsePtrOutputWithContext(ctx context.Context) ConfigurationInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationInfoResponse) *ConfigurationInfoResponse {
		return &v
	}).(ConfigurationInfoResponsePtrOutput)
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





type ConfigurationParameterResponseInput interface {
	pulumi.Input

	ToConfigurationParameterResponseOutput() ConfigurationParameterResponseOutput
	ToConfigurationParameterResponseOutputWithContext(context.Context) ConfigurationParameterResponseOutput
}

type ConfigurationParameterResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConfigurationParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameterResponse)(nil)).Elem()
}

func (i ConfigurationParameterResponseArgs) ToConfigurationParameterResponseOutput() ConfigurationParameterResponseOutput {
	return i.ToConfigurationParameterResponseOutputWithContext(context.Background())
}

func (i ConfigurationParameterResponseArgs) ToConfigurationParameterResponseOutputWithContext(ctx context.Context) ConfigurationParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterResponseOutput)
}





type ConfigurationParameterResponseArrayInput interface {
	pulumi.Input

	ToConfigurationParameterResponseArrayOutput() ConfigurationParameterResponseArrayOutput
	ToConfigurationParameterResponseArrayOutputWithContext(context.Context) ConfigurationParameterResponseArrayOutput
}

type ConfigurationParameterResponseArray []ConfigurationParameterResponseInput

func (ConfigurationParameterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameterResponse)(nil)).Elem()
}

func (i ConfigurationParameterResponseArray) ToConfigurationParameterResponseArrayOutput() ConfigurationParameterResponseArrayOutput {
	return i.ToConfigurationParameterResponseArrayOutputWithContext(context.Background())
}

func (i ConfigurationParameterResponseArray) ToConfigurationParameterResponseArrayOutputWithContext(ctx context.Context) ConfigurationParameterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterResponseArrayOutput)
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

type ConfigurationSetting struct {
	ActionAfterReboot              *string  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           *bool    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              *string  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins *float64 `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 *bool    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           *float64 `pulumi:"refreshFrequencyMins"`
}





type ConfigurationSettingInput interface {
	pulumi.Input

	ToConfigurationSettingOutput() ConfigurationSettingOutput
	ToConfigurationSettingOutputWithContext(context.Context) ConfigurationSettingOutput
}

type ConfigurationSettingArgs struct {
	ActionAfterReboot              pulumi.StringPtrInput  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           pulumi.BoolPtrInput    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              pulumi.StringPtrInput  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins pulumi.Float64PtrInput `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 pulumi.BoolPtrInput    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           pulumi.Float64PtrInput `pulumi:"refreshFrequencyMins"`
}

func (ConfigurationSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSetting)(nil)).Elem()
}

func (i ConfigurationSettingArgs) ToConfigurationSettingOutput() ConfigurationSettingOutput {
	return i.ToConfigurationSettingOutputWithContext(context.Background())
}

func (i ConfigurationSettingArgs) ToConfigurationSettingOutputWithContext(ctx context.Context) ConfigurationSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingOutput)
}

func (i ConfigurationSettingArgs) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return i.ToConfigurationSettingPtrOutputWithContext(context.Background())
}

func (i ConfigurationSettingArgs) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingOutput).ToConfigurationSettingPtrOutputWithContext(ctx)
}









type ConfigurationSettingPtrInput interface {
	pulumi.Input

	ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput
	ToConfigurationSettingPtrOutputWithContext(context.Context) ConfigurationSettingPtrOutput
}

type configurationSettingPtrType ConfigurationSettingArgs

func ConfigurationSettingPtr(v *ConfigurationSettingArgs) ConfigurationSettingPtrInput {
	return (*configurationSettingPtrType)(v)
}

func (*configurationSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSetting)(nil)).Elem()
}

func (i *configurationSettingPtrType) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return i.ToConfigurationSettingPtrOutputWithContext(context.Background())
}

func (i *configurationSettingPtrType) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingPtrOutput)
}

type ConfigurationSettingOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSetting)(nil)).Elem()
}

func (o ConfigurationSettingOutput) ToConfigurationSettingOutput() ConfigurationSettingOutput {
	return o
}

func (o ConfigurationSettingOutput) ToConfigurationSettingOutputWithContext(ctx context.Context) ConfigurationSettingOutput {
	return o
}

func (o ConfigurationSettingOutput) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return o.ToConfigurationSettingPtrOutputWithContext(context.Background())
}

func (o ConfigurationSettingOutput) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationSetting) *ConfigurationSetting {
		return &v
	}).(ConfigurationSettingPtrOutput)
}

func (o ConfigurationSettingOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *string { return v.ActionAfterReboot }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *bool { return v.AllowModuleOverwrite }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *string { return v.ConfigurationMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *float64 { return v.ConfigurationModeFrequencyMins }).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *bool { return v.RebootIfNeeded }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *float64 { return v.RefreshFrequencyMins }).(pulumi.Float64PtrOutput)
}

type ConfigurationSettingPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSetting)(nil)).Elem()
}

func (o ConfigurationSettingPtrOutput) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return o
}

func (o ConfigurationSettingPtrOutput) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return o
}

func (o ConfigurationSettingPtrOutput) Elem() ConfigurationSettingOutput {
	return o.ApplyT(func(v *ConfigurationSetting) ConfigurationSetting {
		if v != nil {
			return *v
		}
		var ret ConfigurationSetting
		return ret
	}).(ConfigurationSettingOutput)
}

func (o ConfigurationSettingPtrOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *string {
		if v == nil {
			return nil
		}
		return v.ActionAfterReboot
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingPtrOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *bool {
		if v == nil {
			return nil
		}
		return v.AllowModuleOverwrite
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingPtrOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingPtrOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *float64 {
		if v == nil {
			return nil
		}
		return v.ConfigurationModeFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingPtrOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *bool {
		if v == nil {
			return nil
		}
		return v.RebootIfNeeded
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingPtrOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

type ConfigurationSettingResponse struct {
	ActionAfterReboot              *string  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           *bool    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              *string  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins *float64 `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 *bool    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           *float64 `pulumi:"refreshFrequencyMins"`
}





type ConfigurationSettingResponseInput interface {
	pulumi.Input

	ToConfigurationSettingResponseOutput() ConfigurationSettingResponseOutput
	ToConfigurationSettingResponseOutputWithContext(context.Context) ConfigurationSettingResponseOutput
}

type ConfigurationSettingResponseArgs struct {
	ActionAfterReboot              pulumi.StringPtrInput  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           pulumi.BoolPtrInput    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              pulumi.StringPtrInput  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins pulumi.Float64PtrInput `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 pulumi.BoolPtrInput    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           pulumi.Float64PtrInput `pulumi:"refreshFrequencyMins"`
}

func (ConfigurationSettingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSettingResponse)(nil)).Elem()
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponseOutput() ConfigurationSettingResponseOutput {
	return i.ToConfigurationSettingResponseOutputWithContext(context.Background())
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponseOutputWithContext(ctx context.Context) ConfigurationSettingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingResponseOutput)
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return i.ToConfigurationSettingResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingResponseOutput).ToConfigurationSettingResponsePtrOutputWithContext(ctx)
}









type ConfigurationSettingResponsePtrInput interface {
	pulumi.Input

	ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput
	ToConfigurationSettingResponsePtrOutputWithContext(context.Context) ConfigurationSettingResponsePtrOutput
}

type configurationSettingResponsePtrType ConfigurationSettingResponseArgs

func ConfigurationSettingResponsePtr(v *ConfigurationSettingResponseArgs) ConfigurationSettingResponsePtrInput {
	return (*configurationSettingResponsePtrType)(v)
}

func (*configurationSettingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSettingResponse)(nil)).Elem()
}

func (i *configurationSettingResponsePtrType) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return i.ToConfigurationSettingResponsePtrOutputWithContext(context.Background())
}

func (i *configurationSettingResponsePtrType) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingResponsePtrOutput)
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

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return o.ToConfigurationSettingResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationSettingResponse) *ConfigurationSettingResponse {
		return &v
	}).(ConfigurationSettingResponsePtrOutput)
}

func (o ConfigurationSettingResponseOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *string { return v.ActionAfterReboot }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponseOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *bool { return v.AllowModuleOverwrite }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponseOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *string { return v.ConfigurationMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponseOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *float64 { return v.ConfigurationModeFrequencyMins }).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingResponseOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *bool { return v.RebootIfNeeded }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponseOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *float64 { return v.RefreshFrequencyMins }).(pulumi.Float64PtrOutput)
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
		return v.ActionAfterReboot
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowModuleOverwrite
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ConfigurationModeFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RebootIfNeeded
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshFrequencyMins
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





type GuestConfigurationAssignmentPropertiesResponseInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesResponseOutput() GuestConfigurationAssignmentPropertiesResponseOutput
	ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesResponseOutput
}

type GuestConfigurationAssignmentPropertiesResponseArgs struct {
	AssignmentHash              pulumi.StringInput                           `pulumi:"assignmentHash"`
	ComplianceStatus            pulumi.StringInput                           `pulumi:"complianceStatus"`
	Context                     pulumi.StringPtrInput                        `pulumi:"context"`
	GuestConfiguration          GuestConfigurationNavigationResponsePtrInput `pulumi:"guestConfiguration"`
	LastComplianceStatusChecked pulumi.StringInput                           `pulumi:"lastComplianceStatusChecked"`
	LatestAssignmentReport      AssignmentReportResponsePtrInput             `pulumi:"latestAssignmentReport"`
	LatestReportId              pulumi.StringInput                           `pulumi:"latestReportId"`
	ParameterHash               pulumi.StringInput                           `pulumi:"parameterHash"`
	ProvisioningState           pulumi.StringInput                           `pulumi:"provisioningState"`
	ResourceType                pulumi.StringInput                           `pulumi:"resourceType"`
	TargetResourceId            pulumi.StringInput                           `pulumi:"targetResourceId"`
	VmssVMList                  VMSSVMInfoResponseArrayInput                 `pulumi:"vmssVMList"`
}

func (GuestConfigurationAssignmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponseOutput() GuestConfigurationAssignmentPropertiesResponseOutput {
	return i.ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesResponseOutput).ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx)
}









type GuestConfigurationAssignmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput
	ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput
}

type guestConfigurationAssignmentPropertiesResponsePtrType GuestConfigurationAssignmentPropertiesResponseArgs

func GuestConfigurationAssignmentPropertiesResponsePtr(v *GuestConfigurationAssignmentPropertiesResponseArgs) GuestConfigurationAssignmentPropertiesResponsePtrInput {
	return (*guestConfigurationAssignmentPropertiesResponsePtrType)(v)
}

func (*guestConfigurationAssignmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (i *guestConfigurationAssignmentPropertiesResponsePtrType) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *guestConfigurationAssignmentPropertiesResponsePtrType) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesResponsePtrOutput)
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

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o.ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationAssignmentPropertiesResponse) *GuestConfigurationAssignmentPropertiesResponse {
		return &v
	}).(GuestConfigurationAssignmentPropertiesResponsePtrOutput)
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

type GuestConfigurationAssignmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) Elem() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) GuestConfigurationAssignmentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationAssignmentPropertiesResponse
		return ret
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) AssignmentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssignmentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ComplianceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComplianceStatus
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Context
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) GuestConfiguration() GuestConfigurationNavigationResponsePtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *GuestConfigurationNavigationResponse {
		if v == nil {
			return nil
		}
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) LastComplianceStatusChecked() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastComplianceStatusChecked
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) LatestAssignmentReport() AssignmentReportResponsePtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *AssignmentReportResponse {
		if v == nil {
			return nil
		}
		return v.LatestAssignmentReport
	}).(AssignmentReportResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) LatestReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LatestReportId
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ParameterHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ParameterHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ResourceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ResourceType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TargetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) VmssVMList() VMSSVMInfoResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) []VMSSVMInfoResponse {
		if v == nil {
			return nil
		}
		return v.VmssVMList
	}).(VMSSVMInfoResponseArrayOutput)
}

type GuestConfigurationNavigation struct {
	AssignmentType                  *string                  `pulumi:"assignmentType"`
	ConfigurationParameter          []ConfigurationParameter `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter []ConfigurationParameter `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            *ConfigurationSetting    `pulumi:"configurationSetting"`
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
	ConfigurationSetting            ConfigurationSettingPtrInput     `pulumi:"configurationSetting"`
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

func (o GuestConfigurationNavigationOutput) ConfigurationSetting() ConfigurationSettingPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *ConfigurationSetting { return v.ConfigurationSetting }).(ConfigurationSettingPtrOutput)
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

func (o GuestConfigurationNavigationPtrOutput) ConfigurationSetting() ConfigurationSettingPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *ConfigurationSetting {
		if v == nil {
			return nil
		}
		return v.ConfigurationSetting
	}).(ConfigurationSettingPtrOutput)
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
	AssignmentType                  *string                          `pulumi:"assignmentType"`
	ConfigurationParameter          []ConfigurationParameterResponse `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter []ConfigurationParameterResponse `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            *ConfigurationSettingResponse    `pulumi:"configurationSetting"`
	ContentHash                     *string                          `pulumi:"contentHash"`
	ContentType                     string                           `pulumi:"contentType"`
	ContentUri                      *string                          `pulumi:"contentUri"`
	Kind                            *string                          `pulumi:"kind"`
	Name                            *string                          `pulumi:"name"`
	Version                         *string                          `pulumi:"version"`
}





type GuestConfigurationNavigationResponseInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationResponseOutput() GuestConfigurationNavigationResponseOutput
	ToGuestConfigurationNavigationResponseOutputWithContext(context.Context) GuestConfigurationNavigationResponseOutput
}

type GuestConfigurationNavigationResponseArgs struct {
	AssignmentType                  pulumi.StringPtrInput                    `pulumi:"assignmentType"`
	ConfigurationParameter          ConfigurationParameterResponseArrayInput `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter ConfigurationParameterResponseArrayInput `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            ConfigurationSettingResponsePtrInput     `pulumi:"configurationSetting"`
	ContentHash                     pulumi.StringPtrInput                    `pulumi:"contentHash"`
	ContentType                     pulumi.StringInput                       `pulumi:"contentType"`
	ContentUri                      pulumi.StringPtrInput                    `pulumi:"contentUri"`
	Kind                            pulumi.StringPtrInput                    `pulumi:"kind"`
	Name                            pulumi.StringPtrInput                    `pulumi:"name"`
	Version                         pulumi.StringPtrInput                    `pulumi:"version"`
}

func (GuestConfigurationNavigationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponseOutput() GuestConfigurationNavigationResponseOutput {
	return i.ToGuestConfigurationNavigationResponseOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponseOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationResponseOutput)
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return i.ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationResponseOutput).ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx)
}









type GuestConfigurationNavigationResponsePtrInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput
	ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Context) GuestConfigurationNavigationResponsePtrOutput
}

type guestConfigurationNavigationResponsePtrType GuestConfigurationNavigationResponseArgs

func GuestConfigurationNavigationResponsePtr(v *GuestConfigurationNavigationResponseArgs) GuestConfigurationNavigationResponsePtrInput {
	return (*guestConfigurationNavigationResponsePtrType)(v)
}

func (*guestConfigurationNavigationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (i *guestConfigurationNavigationResponsePtrType) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return i.ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Background())
}

func (i *guestConfigurationNavigationResponsePtrType) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationResponsePtrOutput)
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

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return o.ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Background())
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationNavigationResponse) *GuestConfigurationNavigationResponse {
		return &v
	}).(GuestConfigurationNavigationResponsePtrOutput)
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

func (o GuestConfigurationNavigationResponseOutput) ConfigurationSetting() ConfigurationSettingResponsePtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *ConfigurationSettingResponse {
		return v.ConfigurationSetting
	}).(ConfigurationSettingResponsePtrOutput)
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
		return v.ConfigurationSetting
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





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
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

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
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

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

type VMInfoResponse struct {
	Id   string `pulumi:"id"`
	Uuid string `pulumi:"uuid"`
}





type VMInfoResponseInput interface {
	pulumi.Input

	ToVMInfoResponseOutput() VMInfoResponseOutput
	ToVMInfoResponseOutputWithContext(context.Context) VMInfoResponseOutput
}

type VMInfoResponseArgs struct {
	Id   pulumi.StringInput `pulumi:"id"`
	Uuid pulumi.StringInput `pulumi:"uuid"`
}

func (VMInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMInfoResponse)(nil)).Elem()
}

func (i VMInfoResponseArgs) ToVMInfoResponseOutput() VMInfoResponseOutput {
	return i.ToVMInfoResponseOutputWithContext(context.Background())
}

func (i VMInfoResponseArgs) ToVMInfoResponseOutputWithContext(ctx context.Context) VMInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMInfoResponseOutput)
}

func (i VMInfoResponseArgs) ToVMInfoResponsePtrOutput() VMInfoResponsePtrOutput {
	return i.ToVMInfoResponsePtrOutputWithContext(context.Background())
}

func (i VMInfoResponseArgs) ToVMInfoResponsePtrOutputWithContext(ctx context.Context) VMInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMInfoResponseOutput).ToVMInfoResponsePtrOutputWithContext(ctx)
}









type VMInfoResponsePtrInput interface {
	pulumi.Input

	ToVMInfoResponsePtrOutput() VMInfoResponsePtrOutput
	ToVMInfoResponsePtrOutputWithContext(context.Context) VMInfoResponsePtrOutput
}

type vminfoResponsePtrType VMInfoResponseArgs

func VMInfoResponsePtr(v *VMInfoResponseArgs) VMInfoResponsePtrInput {
	return (*vminfoResponsePtrType)(v)
}

func (*vminfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VMInfoResponse)(nil)).Elem()
}

func (i *vminfoResponsePtrType) ToVMInfoResponsePtrOutput() VMInfoResponsePtrOutput {
	return i.ToVMInfoResponsePtrOutputWithContext(context.Background())
}

func (i *vminfoResponsePtrType) ToVMInfoResponsePtrOutputWithContext(ctx context.Context) VMInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMInfoResponsePtrOutput)
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

func (o VMInfoResponseOutput) ToVMInfoResponsePtrOutput() VMInfoResponsePtrOutput {
	return o.ToVMInfoResponsePtrOutputWithContext(context.Background())
}

func (o VMInfoResponseOutput) ToVMInfoResponsePtrOutputWithContext(ctx context.Context) VMInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VMInfoResponse) *VMInfoResponse {
		return &v
	}).(VMInfoResponsePtrOutput)
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





type VMSSVMInfoResponseInput interface {
	pulumi.Input

	ToVMSSVMInfoResponseOutput() VMSSVMInfoResponseOutput
	ToVMSSVMInfoResponseOutputWithContext(context.Context) VMSSVMInfoResponseOutput
}

type VMSSVMInfoResponseArgs struct {
	ComplianceStatus      pulumi.StringInput `pulumi:"complianceStatus"`
	LastComplianceChecked pulumi.StringInput `pulumi:"lastComplianceChecked"`
	LatestReportId        pulumi.StringInput `pulumi:"latestReportId"`
	VmId                  pulumi.StringInput `pulumi:"vmId"`
	VmResourceId          pulumi.StringInput `pulumi:"vmResourceId"`
}

func (VMSSVMInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSVMInfoResponse)(nil)).Elem()
}

func (i VMSSVMInfoResponseArgs) ToVMSSVMInfoResponseOutput() VMSSVMInfoResponseOutput {
	return i.ToVMSSVMInfoResponseOutputWithContext(context.Background())
}

func (i VMSSVMInfoResponseArgs) ToVMSSVMInfoResponseOutputWithContext(ctx context.Context) VMSSVMInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSVMInfoResponseOutput)
}





type VMSSVMInfoResponseArrayInput interface {
	pulumi.Input

	ToVMSSVMInfoResponseArrayOutput() VMSSVMInfoResponseArrayOutput
	ToVMSSVMInfoResponseArrayOutputWithContext(context.Context) VMSSVMInfoResponseArrayOutput
}

type VMSSVMInfoResponseArray []VMSSVMInfoResponseInput

func (VMSSVMInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSVMInfoResponse)(nil)).Elem()
}

func (i VMSSVMInfoResponseArray) ToVMSSVMInfoResponseArrayOutput() VMSSVMInfoResponseArrayOutput {
	return i.ToVMSSVMInfoResponseArrayOutputWithContext(context.Background())
}

func (i VMSSVMInfoResponseArray) ToVMSSVMInfoResponseArrayOutputWithContext(ctx context.Context) VMSSVMInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSVMInfoResponseArrayOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentInfoResponseInput)(nil)).Elem(), AssignmentInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentInfoResponsePtrInput)(nil)).Elem(), AssignmentInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentReportResourceComplianceReasonResponseInput)(nil)).Elem(), AssignmentReportResourceComplianceReasonResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentReportResourceComplianceReasonResponseArrayInput)(nil)).Elem(), AssignmentReportResourceComplianceReasonResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentReportResourceResponseInput)(nil)).Elem(), AssignmentReportResourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentReportResourceResponseArrayInput)(nil)).Elem(), AssignmentReportResourceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentReportResponseInput)(nil)).Elem(), AssignmentReportResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AssignmentReportResponsePtrInput)(nil)).Elem(), AssignmentReportResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationInfoResponseInput)(nil)).Elem(), ConfigurationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationInfoResponsePtrInput)(nil)).Elem(), ConfigurationInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationParameterInput)(nil)).Elem(), ConfigurationParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationParameterArrayInput)(nil)).Elem(), ConfigurationParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationParameterResponseInput)(nil)).Elem(), ConfigurationParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationParameterResponseArrayInput)(nil)).Elem(), ConfigurationParameterResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSettingInput)(nil)).Elem(), ConfigurationSettingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSettingPtrInput)(nil)).Elem(), ConfigurationSettingArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSettingResponseInput)(nil)).Elem(), ConfigurationSettingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConfigurationSettingResponsePtrInput)(nil)).Elem(), ConfigurationSettingResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationAssignmentPropertiesInput)(nil)).Elem(), GuestConfigurationAssignmentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationAssignmentPropertiesPtrInput)(nil)).Elem(), GuestConfigurationAssignmentPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationAssignmentPropertiesResponseInput)(nil)).Elem(), GuestConfigurationAssignmentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationAssignmentPropertiesResponsePtrInput)(nil)).Elem(), GuestConfigurationAssignmentPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationNavigationInput)(nil)).Elem(), GuestConfigurationNavigationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationNavigationPtrInput)(nil)).Elem(), GuestConfigurationNavigationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationNavigationResponseInput)(nil)).Elem(), GuestConfigurationNavigationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GuestConfigurationNavigationResponsePtrInput)(nil)).Elem(), GuestConfigurationNavigationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponseInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SystemDataResponsePtrInput)(nil)).Elem(), SystemDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMInfoResponseInput)(nil)).Elem(), VMInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMInfoResponsePtrInput)(nil)).Elem(), VMInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSSVMInfoResponseInput)(nil)).Elem(), VMSSVMInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VMSSVMInfoResponseArrayInput)(nil)).Elem(), VMSSVMInfoResponseArray{})
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
	pulumi.RegisterOutputType(ConfigurationSettingOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(VMInfoResponseOutput{})
	pulumi.RegisterOutputType(VMInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(VMSSVMInfoResponseOutput{})
	pulumi.RegisterOutputType(VMSSVMInfoResponseArrayOutput{})
}
