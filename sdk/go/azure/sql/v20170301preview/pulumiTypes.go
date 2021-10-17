


package v20170301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DatabaseVulnerabilityAssessmentRuleBaselineItem struct {
	Result []string `pulumi:"result"`
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemArgs struct {
	Result pulumi.StringArrayInput `pulumi:"result"`
}

func (DatabaseVulnerabilityAssessmentRuleBaselineItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput)
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemArrayInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemArray []DatabaseVulnerabilityAssessmentRuleBaselineItemInput

func (DatabaseVulnerabilityAssessmentRuleBaselineItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseVulnerabilityAssessmentRuleBaselineItem) []string { return v.Result }).(pulumi.StringArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItem)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput) Index(i pulumi.IntInput) DatabaseVulnerabilityAssessmentRuleBaselineItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseVulnerabilityAssessmentRuleBaselineItem {
		return vs[0].([]DatabaseVulnerabilityAssessmentRuleBaselineItem)[vs[1].(int)]
	}).(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponse struct {
	Result []string `pulumi:"result"`
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs struct {
	Result pulumi.StringArrayInput `pulumi:"result"`
}

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput)
}





type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayInput interface {
	pulumi.Input

	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput
	ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray []DatabaseVulnerabilityAssessmentRuleBaselineItemResponseInput

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return i.ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(context.Background())
}

func (i DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput) Result() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseVulnerabilityAssessmentRuleBaselineItemResponse) []string { return v.Result }).(pulumi.StringArrayOutput)
}

type DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)(nil)).Elem()
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput() DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) ToDatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutputWithContext(ctx context.Context) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput {
	return o
}

func (o DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput) Index(i pulumi.IntInput) DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseVulnerabilityAssessmentRuleBaselineItemResponse {
		return vs[0].([]DatabaseVulnerabilityAssessmentRuleBaselineItemResponse)[vs[1].(int)]
	}).(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput)
}

type JobSchedule struct {
	Enabled   *bool            `pulumi:"enabled"`
	EndTime   *string          `pulumi:"endTime"`
	Interval  *string          `pulumi:"interval"`
	StartTime *string          `pulumi:"startTime"`
	Type      *JobScheduleType `pulumi:"type"`
}





type JobScheduleInput interface {
	pulumi.Input

	ToJobScheduleOutput() JobScheduleOutput
	ToJobScheduleOutputWithContext(context.Context) JobScheduleOutput
}

type JobScheduleArgs struct {
	Enabled   pulumi.BoolPtrInput     `pulumi:"enabled"`
	EndTime   pulumi.StringPtrInput   `pulumi:"endTime"`
	Interval  pulumi.StringPtrInput   `pulumi:"interval"`
	StartTime pulumi.StringPtrInput   `pulumi:"startTime"`
	Type      JobScheduleTypePtrInput `pulumi:"type"`
}

func (JobScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSchedule)(nil)).Elem()
}

func (i JobScheduleArgs) ToJobScheduleOutput() JobScheduleOutput {
	return i.ToJobScheduleOutputWithContext(context.Background())
}

func (i JobScheduleArgs) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleOutput)
}

func (i JobScheduleArgs) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return i.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (i JobScheduleArgs) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleOutput).ToJobSchedulePtrOutputWithContext(ctx)
}









type JobSchedulePtrInput interface {
	pulumi.Input

	ToJobSchedulePtrOutput() JobSchedulePtrOutput
	ToJobSchedulePtrOutputWithContext(context.Context) JobSchedulePtrOutput
}

type jobSchedulePtrType JobScheduleArgs

func JobSchedulePtr(v *JobScheduleArgs) JobSchedulePtrInput {
	return (*jobSchedulePtrType)(v)
}

func (*jobSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil)).Elem()
}

func (i *jobSchedulePtrType) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return i.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (i *jobSchedulePtrType) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobSchedulePtrOutput)
}

type JobScheduleOutput struct{ *pulumi.OutputState }

func (JobScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobSchedule)(nil)).Elem()
}

func (o JobScheduleOutput) ToJobScheduleOutput() JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return o.ToJobSchedulePtrOutputWithContext(context.Background())
}

func (o JobScheduleOutput) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobSchedule) *JobSchedule {
		return &v
	}).(JobSchedulePtrOutput)
}

func (o JobScheduleOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobSchedule) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o JobScheduleOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobSchedule) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobSchedule) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobSchedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) Type() JobScheduleTypePtrOutput {
	return o.ApplyT(func(v JobSchedule) *JobScheduleType { return v.Type }).(JobScheduleTypePtrOutput)
}

type JobSchedulePtrOutput struct{ *pulumi.OutputState }

func (JobSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil)).Elem()
}

func (o JobSchedulePtrOutput) ToJobSchedulePtrOutput() JobSchedulePtrOutput {
	return o
}

func (o JobSchedulePtrOutput) ToJobSchedulePtrOutputWithContext(ctx context.Context) JobSchedulePtrOutput {
	return o
}

func (o JobSchedulePtrOutput) Elem() JobScheduleOutput {
	return o.ApplyT(func(v *JobSchedule) JobSchedule {
		if v != nil {
			return *v
		}
		var ret JobSchedule
		return ret
	}).(JobScheduleOutput)
}

func (o JobSchedulePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o JobSchedulePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o JobSchedulePtrOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *string {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.StringPtrOutput)
}

func (o JobSchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JobSchedulePtrOutput) Type() JobScheduleTypePtrOutput {
	return o.ApplyT(func(v *JobSchedule) *JobScheduleType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(JobScheduleTypePtrOutput)
}

type JobScheduleResponse struct {
	Enabled   *bool   `pulumi:"enabled"`
	EndTime   *string `pulumi:"endTime"`
	Interval  *string `pulumi:"interval"`
	StartTime *string `pulumi:"startTime"`
	Type      *string `pulumi:"type"`
}





type JobScheduleResponseInput interface {
	pulumi.Input

	ToJobScheduleResponseOutput() JobScheduleResponseOutput
	ToJobScheduleResponseOutputWithContext(context.Context) JobScheduleResponseOutput
}

type JobScheduleResponseArgs struct {
	Enabled   pulumi.BoolPtrInput   `pulumi:"enabled"`
	EndTime   pulumi.StringPtrInput `pulumi:"endTime"`
	Interval  pulumi.StringPtrInput `pulumi:"interval"`
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
	Type      pulumi.StringPtrInput `pulumi:"type"`
}

func (JobScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobScheduleResponse)(nil)).Elem()
}

func (i JobScheduleResponseArgs) ToJobScheduleResponseOutput() JobScheduleResponseOutput {
	return i.ToJobScheduleResponseOutputWithContext(context.Background())
}

func (i JobScheduleResponseArgs) ToJobScheduleResponseOutputWithContext(ctx context.Context) JobScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleResponseOutput)
}

func (i JobScheduleResponseArgs) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return i.ToJobScheduleResponsePtrOutputWithContext(context.Background())
}

func (i JobScheduleResponseArgs) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleResponseOutput).ToJobScheduleResponsePtrOutputWithContext(ctx)
}









type JobScheduleResponsePtrInput interface {
	pulumi.Input

	ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput
	ToJobScheduleResponsePtrOutputWithContext(context.Context) JobScheduleResponsePtrOutput
}

type jobScheduleResponsePtrType JobScheduleResponseArgs

func JobScheduleResponsePtr(v *JobScheduleResponseArgs) JobScheduleResponsePtrInput {
	return (*jobScheduleResponsePtrType)(v)
}

func (*jobScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobScheduleResponse)(nil)).Elem()
}

func (i *jobScheduleResponsePtrType) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return i.ToJobScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *jobScheduleResponsePtrType) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleResponsePtrOutput)
}

type JobScheduleResponseOutput struct{ *pulumi.OutputState }

func (JobScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobScheduleResponse)(nil)).Elem()
}

func (o JobScheduleResponseOutput) ToJobScheduleResponseOutput() JobScheduleResponseOutput {
	return o
}

func (o JobScheduleResponseOutput) ToJobScheduleResponseOutputWithContext(ctx context.Context) JobScheduleResponseOutput {
	return o
}

func (o JobScheduleResponseOutput) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return o.ToJobScheduleResponsePtrOutputWithContext(context.Background())
}

func (o JobScheduleResponseOutput) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobScheduleResponse) *JobScheduleResponse {
		return &v
	}).(JobScheduleResponsePtrOutput)
}

func (o JobScheduleResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o JobScheduleResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponseOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobScheduleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (JobScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobScheduleResponse)(nil)).Elem()
}

func (o JobScheduleResponsePtrOutput) ToJobScheduleResponsePtrOutput() JobScheduleResponsePtrOutput {
	return o
}

func (o JobScheduleResponsePtrOutput) ToJobScheduleResponsePtrOutputWithContext(ctx context.Context) JobScheduleResponsePtrOutput {
	return o
}

func (o JobScheduleResponsePtrOutput) Elem() JobScheduleResponseOutput {
	return o.ApplyT(func(v *JobScheduleResponse) JobScheduleResponse {
		if v != nil {
			return *v
		}
		var ret JobScheduleResponse
		return ret
	}).(JobScheduleResponseOutput)
}

func (o JobScheduleResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o JobScheduleResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponsePtrOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Interval
	}).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o JobScheduleResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobStepAction struct {
	Source *string `pulumi:"source"`
	Type   *string `pulumi:"type"`
	Value  string  `pulumi:"value"`
}





type JobStepActionInput interface {
	pulumi.Input

	ToJobStepActionOutput() JobStepActionOutput
	ToJobStepActionOutputWithContext(context.Context) JobStepActionOutput
}

type JobStepActionArgs struct {
	Source pulumi.StringPtrInput `pulumi:"source"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
	Value  pulumi.StringInput    `pulumi:"value"`
}

func (JobStepActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepAction)(nil)).Elem()
}

func (i JobStepActionArgs) ToJobStepActionOutput() JobStepActionOutput {
	return i.ToJobStepActionOutputWithContext(context.Background())
}

func (i JobStepActionArgs) ToJobStepActionOutputWithContext(ctx context.Context) JobStepActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionOutput)
}

func (i JobStepActionArgs) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return i.ToJobStepActionPtrOutputWithContext(context.Background())
}

func (i JobStepActionArgs) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionOutput).ToJobStepActionPtrOutputWithContext(ctx)
}









type JobStepActionPtrInput interface {
	pulumi.Input

	ToJobStepActionPtrOutput() JobStepActionPtrOutput
	ToJobStepActionPtrOutputWithContext(context.Context) JobStepActionPtrOutput
}

type jobStepActionPtrType JobStepActionArgs

func JobStepActionPtr(v *JobStepActionArgs) JobStepActionPtrInput {
	return (*jobStepActionPtrType)(v)
}

func (*jobStepActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepAction)(nil)).Elem()
}

func (i *jobStepActionPtrType) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return i.ToJobStepActionPtrOutputWithContext(context.Background())
}

func (i *jobStepActionPtrType) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionPtrOutput)
}

type JobStepActionOutput struct{ *pulumi.OutputState }

func (JobStepActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepAction)(nil)).Elem()
}

func (o JobStepActionOutput) ToJobStepActionOutput() JobStepActionOutput {
	return o
}

func (o JobStepActionOutput) ToJobStepActionOutputWithContext(ctx context.Context) JobStepActionOutput {
	return o
}

func (o JobStepActionOutput) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return o.ToJobStepActionPtrOutputWithContext(context.Background())
}

func (o JobStepActionOutput) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepAction) *JobStepAction {
		return &v
	}).(JobStepActionPtrOutput)
}

func (o JobStepActionOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepAction) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o JobStepActionOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepAction) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o JobStepActionOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepAction) string { return v.Value }).(pulumi.StringOutput)
}

type JobStepActionPtrOutput struct{ *pulumi.OutputState }

func (JobStepActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepAction)(nil)).Elem()
}

func (o JobStepActionPtrOutput) ToJobStepActionPtrOutput() JobStepActionPtrOutput {
	return o
}

func (o JobStepActionPtrOutput) ToJobStepActionPtrOutputWithContext(ctx context.Context) JobStepActionPtrOutput {
	return o
}

func (o JobStepActionPtrOutput) Elem() JobStepActionOutput {
	return o.ApplyT(func(v *JobStepAction) JobStepAction {
		if v != nil {
			return *v
		}
		var ret JobStepAction
		return ret
	}).(JobStepActionOutput)
}

func (o JobStepActionPtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepAction) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepAction) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepAction) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type JobStepActionResponse struct {
	Source *string `pulumi:"source"`
	Type   *string `pulumi:"type"`
	Value  string  `pulumi:"value"`
}





type JobStepActionResponseInput interface {
	pulumi.Input

	ToJobStepActionResponseOutput() JobStepActionResponseOutput
	ToJobStepActionResponseOutputWithContext(context.Context) JobStepActionResponseOutput
}

type JobStepActionResponseArgs struct {
	Source pulumi.StringPtrInput `pulumi:"source"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
	Value  pulumi.StringInput    `pulumi:"value"`
}

func (JobStepActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionResponse)(nil)).Elem()
}

func (i JobStepActionResponseArgs) ToJobStepActionResponseOutput() JobStepActionResponseOutput {
	return i.ToJobStepActionResponseOutputWithContext(context.Background())
}

func (i JobStepActionResponseArgs) ToJobStepActionResponseOutputWithContext(ctx context.Context) JobStepActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionResponseOutput)
}

func (i JobStepActionResponseArgs) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return i.ToJobStepActionResponsePtrOutputWithContext(context.Background())
}

func (i JobStepActionResponseArgs) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionResponseOutput).ToJobStepActionResponsePtrOutputWithContext(ctx)
}









type JobStepActionResponsePtrInput interface {
	pulumi.Input

	ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput
	ToJobStepActionResponsePtrOutputWithContext(context.Context) JobStepActionResponsePtrOutput
}

type jobStepActionResponsePtrType JobStepActionResponseArgs

func JobStepActionResponsePtr(v *JobStepActionResponseArgs) JobStepActionResponsePtrInput {
	return (*jobStepActionResponsePtrType)(v)
}

func (*jobStepActionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepActionResponse)(nil)).Elem()
}

func (i *jobStepActionResponsePtrType) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return i.ToJobStepActionResponsePtrOutputWithContext(context.Background())
}

func (i *jobStepActionResponsePtrType) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepActionResponsePtrOutput)
}

type JobStepActionResponseOutput struct{ *pulumi.OutputState }

func (JobStepActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepActionResponse)(nil)).Elem()
}

func (o JobStepActionResponseOutput) ToJobStepActionResponseOutput() JobStepActionResponseOutput {
	return o
}

func (o JobStepActionResponseOutput) ToJobStepActionResponseOutputWithContext(ctx context.Context) JobStepActionResponseOutput {
	return o
}

func (o JobStepActionResponseOutput) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return o.ToJobStepActionResponsePtrOutputWithContext(context.Background())
}

func (o JobStepActionResponseOutput) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepActionResponse) *JobStepActionResponse {
		return &v
	}).(JobStepActionResponsePtrOutput)
}

func (o JobStepActionResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepActionResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepActionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepActionResponse) string { return v.Value }).(pulumi.StringOutput)
}

type JobStepActionResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStepActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepActionResponse)(nil)).Elem()
}

func (o JobStepActionResponsePtrOutput) ToJobStepActionResponsePtrOutput() JobStepActionResponsePtrOutput {
	return o
}

func (o JobStepActionResponsePtrOutput) ToJobStepActionResponsePtrOutputWithContext(ctx context.Context) JobStepActionResponsePtrOutput {
	return o
}

func (o JobStepActionResponsePtrOutput) Elem() JobStepActionResponseOutput {
	return o.ApplyT(func(v *JobStepActionResponse) JobStepActionResponse {
		if v != nil {
			return *v
		}
		var ret JobStepActionResponse
		return ret
	}).(JobStepActionResponseOutput)
}

func (o JobStepActionResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o JobStepActionResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepActionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type JobStepExecutionOptions struct {
	InitialRetryIntervalSeconds    *int     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    *int     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  *int     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier *float64 `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 *int     `pulumi:"timeoutSeconds"`
}





type JobStepExecutionOptionsInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsOutput() JobStepExecutionOptionsOutput
	ToJobStepExecutionOptionsOutputWithContext(context.Context) JobStepExecutionOptionsOutput
}

type JobStepExecutionOptionsArgs struct {
	InitialRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  pulumi.IntPtrInput     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier pulumi.Float64PtrInput `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 pulumi.IntPtrInput     `pulumi:"timeoutSeconds"`
}

func (JobStepExecutionOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptions)(nil)).Elem()
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsOutput() JobStepExecutionOptionsOutput {
	return i.ToJobStepExecutionOptionsOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsOutputWithContext(ctx context.Context) JobStepExecutionOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsOutput)
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return i.ToJobStepExecutionOptionsPtrOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsArgs) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsOutput).ToJobStepExecutionOptionsPtrOutputWithContext(ctx)
}









type JobStepExecutionOptionsPtrInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput
	ToJobStepExecutionOptionsPtrOutputWithContext(context.Context) JobStepExecutionOptionsPtrOutput
}

type jobStepExecutionOptionsPtrType JobStepExecutionOptionsArgs

func JobStepExecutionOptionsPtr(v *JobStepExecutionOptionsArgs) JobStepExecutionOptionsPtrInput {
	return (*jobStepExecutionOptionsPtrType)(v)
}

func (*jobStepExecutionOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptions)(nil)).Elem()
}

func (i *jobStepExecutionOptionsPtrType) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return i.ToJobStepExecutionOptionsPtrOutputWithContext(context.Background())
}

func (i *jobStepExecutionOptionsPtrType) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsPtrOutput)
}

type JobStepExecutionOptionsOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptions)(nil)).Elem()
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsOutput() JobStepExecutionOptionsOutput {
	return o
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsOutputWithContext(ctx context.Context) JobStepExecutionOptionsOutput {
	return o
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return o.ToJobStepExecutionOptionsPtrOutputWithContext(context.Background())
}

func (o JobStepExecutionOptionsOutput) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepExecutionOptions) *JobStepExecutionOptions {
		return &v
	}).(JobStepExecutionOptionsPtrOutput)
}

func (o JobStepExecutionOptionsOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.InitialRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.MaximumRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.RetryAttempts }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *float64 { return v.RetryIntervalBackoffMultiplier }).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptions) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type JobStepExecutionOptionsPtrOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptions)(nil)).Elem()
}

func (o JobStepExecutionOptionsPtrOutput) ToJobStepExecutionOptionsPtrOutput() JobStepExecutionOptionsPtrOutput {
	return o
}

func (o JobStepExecutionOptionsPtrOutput) ToJobStepExecutionOptionsPtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsPtrOutput {
	return o
}

func (o JobStepExecutionOptionsPtrOutput) Elem() JobStepExecutionOptionsOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) JobStepExecutionOptions {
		if v != nil {
			return *v
		}
		var ret JobStepExecutionOptions
		return ret
	}).(JobStepExecutionOptionsOutput)
}

func (o JobStepExecutionOptionsPtrOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.InitialRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.MaximumRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.RetryAttempts
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryIntervalBackoffMultiplier
	}).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsPtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptions) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type JobStepExecutionOptionsResponse struct {
	InitialRetryIntervalSeconds    *int     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    *int     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  *int     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier *float64 `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 *int     `pulumi:"timeoutSeconds"`
}





type JobStepExecutionOptionsResponseInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsResponseOutput() JobStepExecutionOptionsResponseOutput
	ToJobStepExecutionOptionsResponseOutputWithContext(context.Context) JobStepExecutionOptionsResponseOutput
}

type JobStepExecutionOptionsResponseArgs struct {
	InitialRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"initialRetryIntervalSeconds"`
	MaximumRetryIntervalSeconds    pulumi.IntPtrInput     `pulumi:"maximumRetryIntervalSeconds"`
	RetryAttempts                  pulumi.IntPtrInput     `pulumi:"retryAttempts"`
	RetryIntervalBackoffMultiplier pulumi.Float64PtrInput `pulumi:"retryIntervalBackoffMultiplier"`
	TimeoutSeconds                 pulumi.IntPtrInput     `pulumi:"timeoutSeconds"`
}

func (JobStepExecutionOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponseOutput() JobStepExecutionOptionsResponseOutput {
	return i.ToJobStepExecutionOptionsResponseOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponseOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsResponseOutput)
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return i.ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Background())
}

func (i JobStepExecutionOptionsResponseArgs) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsResponseOutput).ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx)
}









type JobStepExecutionOptionsResponsePtrInput interface {
	pulumi.Input

	ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput
	ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Context) JobStepExecutionOptionsResponsePtrOutput
}

type jobStepExecutionOptionsResponsePtrType JobStepExecutionOptionsResponseArgs

func JobStepExecutionOptionsResponsePtr(v *JobStepExecutionOptionsResponseArgs) JobStepExecutionOptionsResponsePtrInput {
	return (*jobStepExecutionOptionsResponsePtrType)(v)
}

func (*jobStepExecutionOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (i *jobStepExecutionOptionsResponsePtrType) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return i.ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *jobStepExecutionOptionsResponsePtrType) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepExecutionOptionsResponsePtrOutput)
}

type JobStepExecutionOptionsResponseOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponseOutput() JobStepExecutionOptionsResponseOutput {
	return o
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponseOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponseOutput {
	return o
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return o.ToJobStepExecutionOptionsResponsePtrOutputWithContext(context.Background())
}

func (o JobStepExecutionOptionsResponseOutput) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepExecutionOptionsResponse) *JobStepExecutionOptionsResponse {
		return &v
	}).(JobStepExecutionOptionsResponsePtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.InitialRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.MaximumRetryIntervalSeconds }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.RetryAttempts }).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *float64 { return v.RetryIntervalBackoffMultiplier }).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsResponseOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v JobStepExecutionOptionsResponse) *int { return v.TimeoutSeconds }).(pulumi.IntPtrOutput)
}

type JobStepExecutionOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStepExecutionOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepExecutionOptionsResponse)(nil)).Elem()
}

func (o JobStepExecutionOptionsResponsePtrOutput) ToJobStepExecutionOptionsResponsePtrOutput() JobStepExecutionOptionsResponsePtrOutput {
	return o
}

func (o JobStepExecutionOptionsResponsePtrOutput) ToJobStepExecutionOptionsResponsePtrOutputWithContext(ctx context.Context) JobStepExecutionOptionsResponsePtrOutput {
	return o
}

func (o JobStepExecutionOptionsResponsePtrOutput) Elem() JobStepExecutionOptionsResponseOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) JobStepExecutionOptionsResponse {
		if v != nil {
			return *v
		}
		var ret JobStepExecutionOptionsResponse
		return ret
	}).(JobStepExecutionOptionsResponseOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) InitialRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.InitialRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) MaximumRetryIntervalSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaximumRetryIntervalSeconds
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) RetryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetryAttempts
	}).(pulumi.IntPtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) RetryIntervalBackoffMultiplier() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RetryIntervalBackoffMultiplier
	}).(pulumi.Float64PtrOutput)
}

func (o JobStepExecutionOptionsResponsePtrOutput) TimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *JobStepExecutionOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutSeconds
	}).(pulumi.IntPtrOutput)
}

type JobStepOutputType struct {
	Credential        string  `pulumi:"credential"`
	DatabaseName      string  `pulumi:"databaseName"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SchemaName        *string `pulumi:"schemaName"`
	ServerName        string  `pulumi:"serverName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
	TableName         string  `pulumi:"tableName"`
	Type              *string `pulumi:"type"`
}





type JobStepOutputTypeInput interface {
	pulumi.Input

	ToJobStepOutputTypeOutput() JobStepOutputTypeOutput
	ToJobStepOutputTypeOutputWithContext(context.Context) JobStepOutputTypeOutput
}

type JobStepOutputTypeArgs struct {
	Credential        pulumi.StringInput    `pulumi:"credential"`
	DatabaseName      pulumi.StringInput    `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringPtrInput `pulumi:"resourceGroupName"`
	SchemaName        pulumi.StringPtrInput `pulumi:"schemaName"`
	ServerName        pulumi.StringInput    `pulumi:"serverName"`
	SubscriptionId    pulumi.StringPtrInput `pulumi:"subscriptionId"`
	TableName         pulumi.StringInput    `pulumi:"tableName"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (JobStepOutputTypeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputType)(nil)).Elem()
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypeOutput() JobStepOutputTypeOutput {
	return i.ToJobStepOutputTypeOutputWithContext(context.Background())
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypeOutputWithContext(ctx context.Context) JobStepOutputTypeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputTypeOutput)
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return i.ToJobStepOutputTypePtrOutputWithContext(context.Background())
}

func (i JobStepOutputTypeArgs) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputTypeOutput).ToJobStepOutputTypePtrOutputWithContext(ctx)
}









type JobStepOutputTypePtrInput interface {
	pulumi.Input

	ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput
	ToJobStepOutputTypePtrOutputWithContext(context.Context) JobStepOutputTypePtrOutput
}

type jobStepOutputTypePtrType JobStepOutputTypeArgs

func JobStepOutputTypePtr(v *JobStepOutputTypeArgs) JobStepOutputTypePtrInput {
	return (*jobStepOutputTypePtrType)(v)
}

func (*jobStepOutputTypePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputType)(nil)).Elem()
}

func (i *jobStepOutputTypePtrType) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return i.ToJobStepOutputTypePtrOutputWithContext(context.Background())
}

func (i *jobStepOutputTypePtrType) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputTypePtrOutput)
}

type JobStepOutputTypeOutput struct{ *pulumi.OutputState }

func (JobStepOutputTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputType)(nil)).Elem()
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypeOutput() JobStepOutputTypeOutput {
	return o
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypeOutputWithContext(ctx context.Context) JobStepOutputTypeOutput {
	return o
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return o.ToJobStepOutputTypePtrOutputWithContext(context.Background())
}

func (o JobStepOutputTypeOutput) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepOutputType) *JobStepOutputType {
		return &v
	}).(JobStepOutputTypePtrOutput)
}

func (o JobStepOutputTypeOutput) Credential() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.Credential }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypeOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypeOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypeOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputType) string { return v.TableName }).(pulumi.StringOutput)
}

func (o JobStepOutputTypeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputType) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobStepOutputTypePtrOutput struct{ *pulumi.OutputState }

func (JobStepOutputTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputType)(nil)).Elem()
}

func (o JobStepOutputTypePtrOutput) ToJobStepOutputTypePtrOutput() JobStepOutputTypePtrOutput {
	return o
}

func (o JobStepOutputTypePtrOutput) ToJobStepOutputTypePtrOutputWithContext(ctx context.Context) JobStepOutputTypePtrOutput {
	return o
}

func (o JobStepOutputTypePtrOutput) Elem() JobStepOutputTypeOutput {
	return o.ApplyT(func(v *JobStepOutputType) JobStepOutputType {
		if v != nil {
			return *v
		}
		var ret JobStepOutputType
		return ret
	}).(JobStepOutputTypeOutput)
}

func (o JobStepOutputTypePtrOutput) Credential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.Credential
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.SchemaName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return &v.TableName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputTypePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputType) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobStepOutputResponse struct {
	Credential        string  `pulumi:"credential"`
	DatabaseName      string  `pulumi:"databaseName"`
	ResourceGroupName *string `pulumi:"resourceGroupName"`
	SchemaName        *string `pulumi:"schemaName"`
	ServerName        string  `pulumi:"serverName"`
	SubscriptionId    *string `pulumi:"subscriptionId"`
	TableName         string  `pulumi:"tableName"`
	Type              *string `pulumi:"type"`
}





type JobStepOutputResponseInput interface {
	pulumi.Input

	ToJobStepOutputResponseOutput() JobStepOutputResponseOutput
	ToJobStepOutputResponseOutputWithContext(context.Context) JobStepOutputResponseOutput
}

type JobStepOutputResponseArgs struct {
	Credential        pulumi.StringInput    `pulumi:"credential"`
	DatabaseName      pulumi.StringInput    `pulumi:"databaseName"`
	ResourceGroupName pulumi.StringPtrInput `pulumi:"resourceGroupName"`
	SchemaName        pulumi.StringPtrInput `pulumi:"schemaName"`
	ServerName        pulumi.StringInput    `pulumi:"serverName"`
	SubscriptionId    pulumi.StringPtrInput `pulumi:"subscriptionId"`
	TableName         pulumi.StringInput    `pulumi:"tableName"`
	Type              pulumi.StringPtrInput `pulumi:"type"`
}

func (JobStepOutputResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputResponse)(nil)).Elem()
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponseOutput() JobStepOutputResponseOutput {
	return i.ToJobStepOutputResponseOutputWithContext(context.Background())
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponseOutputWithContext(ctx context.Context) JobStepOutputResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputResponseOutput)
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return i.ToJobStepOutputResponsePtrOutputWithContext(context.Background())
}

func (i JobStepOutputResponseArgs) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputResponseOutput).ToJobStepOutputResponsePtrOutputWithContext(ctx)
}









type JobStepOutputResponsePtrInput interface {
	pulumi.Input

	ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput
	ToJobStepOutputResponsePtrOutputWithContext(context.Context) JobStepOutputResponsePtrOutput
}

type jobStepOutputResponsePtrType JobStepOutputResponseArgs

func JobStepOutputResponsePtr(v *JobStepOutputResponseArgs) JobStepOutputResponsePtrInput {
	return (*jobStepOutputResponsePtrType)(v)
}

func (*jobStepOutputResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputResponse)(nil)).Elem()
}

func (i *jobStepOutputResponsePtrType) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return i.ToJobStepOutputResponsePtrOutputWithContext(context.Background())
}

func (i *jobStepOutputResponsePtrType) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutputResponsePtrOutput)
}

type JobStepOutputResponseOutput struct{ *pulumi.OutputState }

func (JobStepOutputResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStepOutputResponse)(nil)).Elem()
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponseOutput() JobStepOutputResponseOutput {
	return o
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponseOutputWithContext(ctx context.Context) JobStepOutputResponseOutput {
	return o
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return o.ToJobStepOutputResponsePtrOutputWithContext(context.Background())
}

func (o JobStepOutputResponseOutput) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JobStepOutputResponse) *JobStepOutputResponse {
		return &v
	}).(JobStepOutputResponsePtrOutput)
}

func (o JobStepOutputResponseOutput) Credential() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.Credential }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) DatabaseName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.DatabaseName }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.ResourceGroupName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponseOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.SchemaName }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponseOutput) ServerName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.ServerName }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.SubscriptionId }).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponseOutput) TableName() pulumi.StringOutput {
	return o.ApplyT(func(v JobStepOutputResponse) string { return v.TableName }).(pulumi.StringOutput)
}

func (o JobStepOutputResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobStepOutputResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type JobStepOutputResponsePtrOutput struct{ *pulumi.OutputState }

func (JobStepOutputResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobStepOutputResponse)(nil)).Elem()
}

func (o JobStepOutputResponsePtrOutput) ToJobStepOutputResponsePtrOutput() JobStepOutputResponsePtrOutput {
	return o
}

func (o JobStepOutputResponsePtrOutput) ToJobStepOutputResponsePtrOutputWithContext(ctx context.Context) JobStepOutputResponsePtrOutput {
	return o
}

func (o JobStepOutputResponsePtrOutput) Elem() JobStepOutputResponseOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) JobStepOutputResponse {
		if v != nil {
			return *v
		}
		var ret JobStepOutputResponse
		return ret
	}).(JobStepOutputResponseOutput)
}

func (o JobStepOutputResponsePtrOutput) Credential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Credential
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DatabaseName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) ResourceGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceGroupName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) SchemaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.SchemaName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServerName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) SubscriptionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubscriptionId
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) TableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TableName
	}).(pulumi.StringPtrOutput)
}

func (o JobStepOutputResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobStepOutputResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type JobTarget struct {
	DatabaseName      *string                       `pulumi:"databaseName"`
	ElasticPoolName   *string                       `pulumi:"elasticPoolName"`
	MembershipType    *JobTargetGroupMembershipType `pulumi:"membershipType"`
	RefreshCredential *string                       `pulumi:"refreshCredential"`
	ServerName        *string                       `pulumi:"serverName"`
	ShardMapName      *string                       `pulumi:"shardMapName"`
	Type              string                        `pulumi:"type"`
}





type JobTargetInput interface {
	pulumi.Input

	ToJobTargetOutput() JobTargetOutput
	ToJobTargetOutputWithContext(context.Context) JobTargetOutput
}

type JobTargetArgs struct {
	DatabaseName      pulumi.StringPtrInput                `pulumi:"databaseName"`
	ElasticPoolName   pulumi.StringPtrInput                `pulumi:"elasticPoolName"`
	MembershipType    JobTargetGroupMembershipTypePtrInput `pulumi:"membershipType"`
	RefreshCredential pulumi.StringPtrInput                `pulumi:"refreshCredential"`
	ServerName        pulumi.StringPtrInput                `pulumi:"serverName"`
	ShardMapName      pulumi.StringPtrInput                `pulumi:"shardMapName"`
	Type              pulumi.StringInput                   `pulumi:"type"`
}

func (JobTargetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTarget)(nil)).Elem()
}

func (i JobTargetArgs) ToJobTargetOutput() JobTargetOutput {
	return i.ToJobTargetOutputWithContext(context.Background())
}

func (i JobTargetArgs) ToJobTargetOutputWithContext(ctx context.Context) JobTargetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetOutput)
}





type JobTargetArrayInput interface {
	pulumi.Input

	ToJobTargetArrayOutput() JobTargetArrayOutput
	ToJobTargetArrayOutputWithContext(context.Context) JobTargetArrayOutput
}

type JobTargetArray []JobTargetInput

func (JobTargetArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTarget)(nil)).Elem()
}

func (i JobTargetArray) ToJobTargetArrayOutput() JobTargetArrayOutput {
	return i.ToJobTargetArrayOutputWithContext(context.Background())
}

func (i JobTargetArray) ToJobTargetArrayOutputWithContext(ctx context.Context) JobTargetArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetArrayOutput)
}

type JobTargetOutput struct{ *pulumi.OutputState }

func (JobTargetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTarget)(nil)).Elem()
}

func (o JobTargetOutput) ToJobTargetOutput() JobTargetOutput {
	return o
}

func (o JobTargetOutput) ToJobTargetOutputWithContext(ctx context.Context) JobTargetOutput {
	return o
}

func (o JobTargetOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) MembershipType() JobTargetGroupMembershipTypePtrOutput {
	return o.ApplyT(func(v JobTarget) *JobTargetGroupMembershipType { return v.MembershipType }).(JobTargetGroupMembershipTypePtrOutput)
}

func (o JobTargetOutput) RefreshCredential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.RefreshCredential }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) ShardMapName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTarget) *string { return v.ShardMapName }).(pulumi.StringPtrOutput)
}

func (o JobTargetOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JobTarget) string { return v.Type }).(pulumi.StringOutput)
}

type JobTargetArrayOutput struct{ *pulumi.OutputState }

func (JobTargetArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTarget)(nil)).Elem()
}

func (o JobTargetArrayOutput) ToJobTargetArrayOutput() JobTargetArrayOutput {
	return o
}

func (o JobTargetArrayOutput) ToJobTargetArrayOutputWithContext(ctx context.Context) JobTargetArrayOutput {
	return o
}

func (o JobTargetArrayOutput) Index(i pulumi.IntInput) JobTargetOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobTarget {
		return vs[0].([]JobTarget)[vs[1].(int)]
	}).(JobTargetOutput)
}

type JobTargetResponse struct {
	DatabaseName      *string `pulumi:"databaseName"`
	ElasticPoolName   *string `pulumi:"elasticPoolName"`
	MembershipType    *string `pulumi:"membershipType"`
	RefreshCredential *string `pulumi:"refreshCredential"`
	ServerName        *string `pulumi:"serverName"`
	ShardMapName      *string `pulumi:"shardMapName"`
	Type              string  `pulumi:"type"`
}





type JobTargetResponseInput interface {
	pulumi.Input

	ToJobTargetResponseOutput() JobTargetResponseOutput
	ToJobTargetResponseOutputWithContext(context.Context) JobTargetResponseOutput
}

type JobTargetResponseArgs struct {
	DatabaseName      pulumi.StringPtrInput `pulumi:"databaseName"`
	ElasticPoolName   pulumi.StringPtrInput `pulumi:"elasticPoolName"`
	MembershipType    pulumi.StringPtrInput `pulumi:"membershipType"`
	RefreshCredential pulumi.StringPtrInput `pulumi:"refreshCredential"`
	ServerName        pulumi.StringPtrInput `pulumi:"serverName"`
	ShardMapName      pulumi.StringPtrInput `pulumi:"shardMapName"`
	Type              pulumi.StringInput    `pulumi:"type"`
}

func (JobTargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetResponse)(nil)).Elem()
}

func (i JobTargetResponseArgs) ToJobTargetResponseOutput() JobTargetResponseOutput {
	return i.ToJobTargetResponseOutputWithContext(context.Background())
}

func (i JobTargetResponseArgs) ToJobTargetResponseOutputWithContext(ctx context.Context) JobTargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetResponseOutput)
}





type JobTargetResponseArrayInput interface {
	pulumi.Input

	ToJobTargetResponseArrayOutput() JobTargetResponseArrayOutput
	ToJobTargetResponseArrayOutputWithContext(context.Context) JobTargetResponseArrayOutput
}

type JobTargetResponseArray []JobTargetResponseInput

func (JobTargetResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTargetResponse)(nil)).Elem()
}

func (i JobTargetResponseArray) ToJobTargetResponseArrayOutput() JobTargetResponseArrayOutput {
	return i.ToJobTargetResponseArrayOutputWithContext(context.Background())
}

func (i JobTargetResponseArray) ToJobTargetResponseArrayOutputWithContext(ctx context.Context) JobTargetResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobTargetResponseArrayOutput)
}

type JobTargetResponseOutput struct{ *pulumi.OutputState }

func (JobTargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobTargetResponse)(nil)).Elem()
}

func (o JobTargetResponseOutput) ToJobTargetResponseOutput() JobTargetResponseOutput {
	return o
}

func (o JobTargetResponseOutput) ToJobTargetResponseOutputWithContext(ctx context.Context) JobTargetResponseOutput {
	return o
}

func (o JobTargetResponseOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) ElasticPoolName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.ElasticPoolName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) MembershipType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.MembershipType }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) RefreshCredential() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.RefreshCredential }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) ServerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.ServerName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) ShardMapName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JobTargetResponse) *string { return v.ShardMapName }).(pulumi.StringPtrOutput)
}

func (o JobTargetResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v JobTargetResponse) string { return v.Type }).(pulumi.StringOutput)
}

type JobTargetResponseArrayOutput struct{ *pulumi.OutputState }

func (JobTargetResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]JobTargetResponse)(nil)).Elem()
}

func (o JobTargetResponseArrayOutput) ToJobTargetResponseArrayOutput() JobTargetResponseArrayOutput {
	return o
}

func (o JobTargetResponseArrayOutput) ToJobTargetResponseArrayOutputWithContext(ctx context.Context) JobTargetResponseArrayOutput {
	return o
}

func (o JobTargetResponseArrayOutput) Index(i pulumi.IntInput) JobTargetResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) JobTargetResponse {
		return vs[0].([]JobTargetResponse)[vs[1].(int)]
	}).(JobTargetResponseOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
	Size     *string `pulumi:"size"`
	Tier     *string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
	Size     pulumi.StringPtrInput `pulumi:"size"`
	Tier     pulumi.StringPtrInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type VulnerabilityAssessmentRecurringScansProperties struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}





type VulnerabilityAssessmentRecurringScansPropertiesInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput
}

type VulnerabilityAssessmentRecurringScansPropertiesArgs struct {
	EmailSubscriptionAdmins pulumi.BoolPtrInput     `pulumi:"emailSubscriptionAdmins"`
	Emails                  pulumi.StringArrayInput `pulumi:"emails"`
	IsEnabled               pulumi.BoolPtrInput     `pulumi:"isEnabled"`
}

func (VulnerabilityAssessmentRecurringScansPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesOutput)
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesArgs) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesOutput).ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx)
}









type VulnerabilityAssessmentRecurringScansPropertiesPtrInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput
}

type vulnerabilityAssessmentRecurringScansPropertiesPtrType VulnerabilityAssessmentRecurringScansPropertiesArgs

func VulnerabilityAssessmentRecurringScansPropertiesPtr(v *VulnerabilityAssessmentRecurringScansPropertiesArgs) VulnerabilityAssessmentRecurringScansPropertiesPtrInput {
	return (*vulnerabilityAssessmentRecurringScansPropertiesPtrType)(v)
}

func (*vulnerabilityAssessmentRecurringScansPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesPtrType) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesPtrType) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesOutput() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o.ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(context.Background())
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VulnerabilityAssessmentRecurringScansProperties) *VulnerabilityAssessmentRecurringScansProperties {
		return &v
	}).(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) *bool { return v.EmailSubscriptionAdmins }).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansProperties) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesPtrOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansProperties)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutput() VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesPtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesPtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) Elem() VulnerabilityAssessmentRecurringScansPropertiesOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) VulnerabilityAssessmentRecurringScansProperties {
		if v != nil {
			return *v
		}
		var ret VulnerabilityAssessmentRecurringScansProperties
		return ret
	}).(VulnerabilityAssessmentRecurringScansPropertiesOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponse struct {
	EmailSubscriptionAdmins *bool    `pulumi:"emailSubscriptionAdmins"`
	Emails                  []string `pulumi:"emails"`
	IsEnabled               *bool    `pulumi:"isEnabled"`
}





type VulnerabilityAssessmentRecurringScansPropertiesResponseInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput
}

type VulnerabilityAssessmentRecurringScansPropertiesResponseArgs struct {
	EmailSubscriptionAdmins pulumi.BoolPtrInput     `pulumi:"emailSubscriptionAdmins"`
	Emails                  pulumi.StringArrayInput `pulumi:"emails"`
	IsEnabled               pulumi.BoolPtrInput     `pulumi:"isEnabled"`
}

func (VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput)
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput).ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx)
}









type VulnerabilityAssessmentRecurringScansPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput
	ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput
}

type vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType VulnerabilityAssessmentRecurringScansPropertiesResponseArgs

func VulnerabilityAssessmentRecurringScansPropertiesResponsePtr(v *VulnerabilityAssessmentRecurringScansPropertiesResponseArgs) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrInput {
	return (*vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType)(v)
}

func (*vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return i.ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *vulnerabilityAssessmentRecurringScansPropertiesResponsePtrType) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutput() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponseOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o.ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VulnerabilityAssessmentRecurringScansPropertiesResponse) *VulnerabilityAssessmentRecurringScansPropertiesResponse {
		return &v
	}).(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) []string { return v.Emails }).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VulnerabilityAssessmentRecurringScansPropertiesResponse)(nil)).Elem()
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput() VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) ToVulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutputWithContext(ctx context.Context) VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput {
	return o
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) Elem() VulnerabilityAssessmentRecurringScansPropertiesResponseOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) VulnerabilityAssessmentRecurringScansPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VulnerabilityAssessmentRecurringScansPropertiesResponse
		return ret
	}).(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) EmailSubscriptionAdmins() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EmailSubscriptionAdmins
	}).(pulumi.BoolPtrOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) Emails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Emails
	}).(pulumi.StringArrayOutput)
}

func (o VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *VulnerabilityAssessmentRecurringScansPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemArrayInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponseInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayInput)(nil)).Elem(), DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobScheduleInput)(nil)).Elem(), JobScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobSchedulePtrInput)(nil)).Elem(), JobScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobScheduleResponseInput)(nil)).Elem(), JobScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobScheduleResponsePtrInput)(nil)).Elem(), JobScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionInput)(nil)).Elem(), JobStepActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionPtrInput)(nil)).Elem(), JobStepActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionResponseInput)(nil)).Elem(), JobStepActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepActionResponsePtrInput)(nil)).Elem(), JobStepActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsInput)(nil)).Elem(), JobStepExecutionOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsPtrInput)(nil)).Elem(), JobStepExecutionOptionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsResponseInput)(nil)).Elem(), JobStepExecutionOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepExecutionOptionsResponsePtrInput)(nil)).Elem(), JobStepExecutionOptionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputTypeInput)(nil)).Elem(), JobStepOutputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputTypePtrInput)(nil)).Elem(), JobStepOutputTypeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputResponseInput)(nil)).Elem(), JobStepOutputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepOutputResponsePtrInput)(nil)).Elem(), JobStepOutputResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetInput)(nil)).Elem(), JobTargetArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetArrayInput)(nil)).Elem(), JobTargetArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetResponseInput)(nil)).Elem(), JobTargetResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*JobTargetResponseArrayInput)(nil)).Elem(), JobTargetResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesPtrInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponseInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VulnerabilityAssessmentRecurringScansPropertiesResponsePtrInput)(nil)).Elem(), VulnerabilityAssessmentRecurringScansPropertiesResponseArgs{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemArrayOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseOutput{})
	pulumi.RegisterOutputType(DatabaseVulnerabilityAssessmentRuleBaselineItemResponseArrayOutput{})
	pulumi.RegisterOutputType(JobScheduleOutput{})
	pulumi.RegisterOutputType(JobSchedulePtrOutput{})
	pulumi.RegisterOutputType(JobScheduleResponseOutput{})
	pulumi.RegisterOutputType(JobScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepActionOutput{})
	pulumi.RegisterOutputType(JobStepActionPtrOutput{})
	pulumi.RegisterOutputType(JobStepActionResponseOutput{})
	pulumi.RegisterOutputType(JobStepActionResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsPtrOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsResponseOutput{})
	pulumi.RegisterOutputType(JobStepExecutionOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(JobStepOutputTypeOutput{})
	pulumi.RegisterOutputType(JobStepOutputTypePtrOutput{})
	pulumi.RegisterOutputType(JobStepOutputResponseOutput{})
	pulumi.RegisterOutputType(JobStepOutputResponsePtrOutput{})
	pulumi.RegisterOutputType(JobTargetOutput{})
	pulumi.RegisterOutputType(JobTargetArrayOutput{})
	pulumi.RegisterOutputType(JobTargetResponseOutput{})
	pulumi.RegisterOutputType(JobTargetResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesPtrOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VulnerabilityAssessmentRecurringScansPropertiesResponsePtrOutput{})
}
