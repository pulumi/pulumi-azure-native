


package v20210808preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AddActionGroups struct {
	ActionGroupIds []string `pulumi:"actionGroupIds"`
	ActionType     string   `pulumi:"actionType"`
}

type AddActionGroupsResponse struct {
	ActionGroupIds []string `pulumi:"actionGroupIds"`
	ActionType     string   `pulumi:"actionType"`
}

type AlertProcessingRuleProperties struct {
	Actions     []interface{} `pulumi:"actions"`
	Conditions  []Condition   `pulumi:"conditions"`
	Description *string       `pulumi:"description"`
	Enabled     *bool         `pulumi:"enabled"`
	Schedule    *Schedule     `pulumi:"schedule"`
	Scopes      []string      `pulumi:"scopes"`
}


func (val *AlertProcessingRuleProperties) Defaults() *AlertProcessingRuleProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := true
		tmp.Enabled = &enabled_
	}
	return &tmp
}





type AlertProcessingRulePropertiesInput interface {
	pulumi.Input

	ToAlertProcessingRulePropertiesOutput() AlertProcessingRulePropertiesOutput
	ToAlertProcessingRulePropertiesOutputWithContext(context.Context) AlertProcessingRulePropertiesOutput
}

type AlertProcessingRulePropertiesArgs struct {
	Actions     pulumi.ArrayInput       `pulumi:"actions"`
	Conditions  ConditionArrayInput     `pulumi:"conditions"`
	Description pulumi.StringPtrInput   `pulumi:"description"`
	Enabled     pulumi.BoolPtrInput     `pulumi:"enabled"`
	Schedule    SchedulePtrInput        `pulumi:"schedule"`
	Scopes      pulumi.StringArrayInput `pulumi:"scopes"`
}


func (val *AlertProcessingRulePropertiesArgs) Defaults() *AlertProcessingRulePropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		tmp.Enabled = pulumi.BoolPtr(true)
	}
	return &tmp
}
func (AlertProcessingRulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertProcessingRuleProperties)(nil)).Elem()
}

func (i AlertProcessingRulePropertiesArgs) ToAlertProcessingRulePropertiesOutput() AlertProcessingRulePropertiesOutput {
	return i.ToAlertProcessingRulePropertiesOutputWithContext(context.Background())
}

func (i AlertProcessingRulePropertiesArgs) ToAlertProcessingRulePropertiesOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertProcessingRulePropertiesOutput)
}

func (i AlertProcessingRulePropertiesArgs) ToAlertProcessingRulePropertiesPtrOutput() AlertProcessingRulePropertiesPtrOutput {
	return i.ToAlertProcessingRulePropertiesPtrOutputWithContext(context.Background())
}

func (i AlertProcessingRulePropertiesArgs) ToAlertProcessingRulePropertiesPtrOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertProcessingRulePropertiesOutput).ToAlertProcessingRulePropertiesPtrOutputWithContext(ctx)
}









type AlertProcessingRulePropertiesPtrInput interface {
	pulumi.Input

	ToAlertProcessingRulePropertiesPtrOutput() AlertProcessingRulePropertiesPtrOutput
	ToAlertProcessingRulePropertiesPtrOutputWithContext(context.Context) AlertProcessingRulePropertiesPtrOutput
}

type alertProcessingRulePropertiesPtrType AlertProcessingRulePropertiesArgs

func AlertProcessingRulePropertiesPtr(v *AlertProcessingRulePropertiesArgs) AlertProcessingRulePropertiesPtrInput {
	return (*alertProcessingRulePropertiesPtrType)(v)
}

func (*alertProcessingRulePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertProcessingRuleProperties)(nil)).Elem()
}

func (i *alertProcessingRulePropertiesPtrType) ToAlertProcessingRulePropertiesPtrOutput() AlertProcessingRulePropertiesPtrOutput {
	return i.ToAlertProcessingRulePropertiesPtrOutputWithContext(context.Background())
}

func (i *alertProcessingRulePropertiesPtrType) ToAlertProcessingRulePropertiesPtrOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertProcessingRulePropertiesPtrOutput)
}

type AlertProcessingRulePropertiesOutput struct{ *pulumi.OutputState }

func (AlertProcessingRulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertProcessingRuleProperties)(nil)).Elem()
}

func (o AlertProcessingRulePropertiesOutput) ToAlertProcessingRulePropertiesOutput() AlertProcessingRulePropertiesOutput {
	return o
}

func (o AlertProcessingRulePropertiesOutput) ToAlertProcessingRulePropertiesOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesOutput {
	return o
}

func (o AlertProcessingRulePropertiesOutput) ToAlertProcessingRulePropertiesPtrOutput() AlertProcessingRulePropertiesPtrOutput {
	return o.ToAlertProcessingRulePropertiesPtrOutputWithContext(context.Background())
}

func (o AlertProcessingRulePropertiesOutput) ToAlertProcessingRulePropertiesPtrOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertProcessingRuleProperties) *AlertProcessingRuleProperties {
		return &v
	}).(AlertProcessingRulePropertiesPtrOutput)
}

func (o AlertProcessingRulePropertiesOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v AlertProcessingRuleProperties) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o AlertProcessingRulePropertiesOutput) Conditions() ConditionArrayOutput {
	return o.ApplyT(func(v AlertProcessingRuleProperties) []Condition { return v.Conditions }).(ConditionArrayOutput)
}

func (o AlertProcessingRulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertProcessingRuleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AlertProcessingRulePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AlertProcessingRuleProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AlertProcessingRulePropertiesOutput) Schedule() SchedulePtrOutput {
	return o.ApplyT(func(v AlertProcessingRuleProperties) *Schedule { return v.Schedule }).(SchedulePtrOutput)
}

func (o AlertProcessingRulePropertiesOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertProcessingRuleProperties) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type AlertProcessingRulePropertiesPtrOutput struct{ *pulumi.OutputState }

func (AlertProcessingRulePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertProcessingRuleProperties)(nil)).Elem()
}

func (o AlertProcessingRulePropertiesPtrOutput) ToAlertProcessingRulePropertiesPtrOutput() AlertProcessingRulePropertiesPtrOutput {
	return o
}

func (o AlertProcessingRulePropertiesPtrOutput) ToAlertProcessingRulePropertiesPtrOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesPtrOutput {
	return o
}

func (o AlertProcessingRulePropertiesPtrOutput) Elem() AlertProcessingRulePropertiesOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) AlertProcessingRuleProperties {
		if v != nil {
			return *v
		}
		var ret AlertProcessingRuleProperties
		return ret
	}).(AlertProcessingRulePropertiesOutput)
}

func (o AlertProcessingRulePropertiesPtrOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) []interface{} {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(pulumi.ArrayOutput)
}

func (o AlertProcessingRulePropertiesPtrOutput) Conditions() ConditionArrayOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) []Condition {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(ConditionArrayOutput)
}

func (o AlertProcessingRulePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o AlertProcessingRulePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AlertProcessingRulePropertiesPtrOutput) Schedule() SchedulePtrOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) *Schedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(SchedulePtrOutput)
}

func (o AlertProcessingRulePropertiesPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AlertProcessingRuleProperties) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type AlertProcessingRulePropertiesResponse struct {
	Actions     []interface{}       `pulumi:"actions"`
	Conditions  []ConditionResponse `pulumi:"conditions"`
	Description *string             `pulumi:"description"`
	Enabled     *bool               `pulumi:"enabled"`
	Schedule    *ScheduleResponse   `pulumi:"schedule"`
	Scopes      []string            `pulumi:"scopes"`
}


func (val *AlertProcessingRulePropertiesResponse) Defaults() *AlertProcessingRulePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Enabled) {
		enabled_ := true
		tmp.Enabled = &enabled_
	}
	return &tmp
}

type AlertProcessingRulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (AlertProcessingRulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertProcessingRulePropertiesResponse)(nil)).Elem()
}

func (o AlertProcessingRulePropertiesResponseOutput) ToAlertProcessingRulePropertiesResponseOutput() AlertProcessingRulePropertiesResponseOutput {
	return o
}

func (o AlertProcessingRulePropertiesResponseOutput) ToAlertProcessingRulePropertiesResponseOutputWithContext(ctx context.Context) AlertProcessingRulePropertiesResponseOutput {
	return o
}

func (o AlertProcessingRulePropertiesResponseOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v AlertProcessingRulePropertiesResponse) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o AlertProcessingRulePropertiesResponseOutput) Conditions() ConditionResponseArrayOutput {
	return o.ApplyT(func(v AlertProcessingRulePropertiesResponse) []ConditionResponse { return v.Conditions }).(ConditionResponseArrayOutput)
}

func (o AlertProcessingRulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertProcessingRulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AlertProcessingRulePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AlertProcessingRulePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AlertProcessingRulePropertiesResponseOutput) Schedule() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v AlertProcessingRulePropertiesResponse) *ScheduleResponse { return v.Schedule }).(ScheduleResponsePtrOutput)
}

func (o AlertProcessingRulePropertiesResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AlertProcessingRulePropertiesResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type Condition struct {
	Field    *string  `pulumi:"field"`
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ConditionInput interface {
	pulumi.Input

	ToConditionOutput() ConditionOutput
	ToConditionOutputWithContext(context.Context) ConditionOutput
}

type ConditionArgs struct {
	Field    pulumi.StringPtrInput   `pulumi:"field"`
	Operator pulumi.StringPtrInput   `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Condition)(nil)).Elem()
}

func (i ConditionArgs) ToConditionOutput() ConditionOutput {
	return i.ToConditionOutputWithContext(context.Background())
}

func (i ConditionArgs) ToConditionOutputWithContext(ctx context.Context) ConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionOutput)
}





type ConditionArrayInput interface {
	pulumi.Input

	ToConditionArrayOutput() ConditionArrayOutput
	ToConditionArrayOutputWithContext(context.Context) ConditionArrayOutput
}

type ConditionArray []ConditionInput

func (ConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Condition)(nil)).Elem()
}

func (i ConditionArray) ToConditionArrayOutput() ConditionArrayOutput {
	return i.ToConditionArrayOutputWithContext(context.Background())
}

func (i ConditionArray) ToConditionArrayOutputWithContext(ctx context.Context) ConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionArrayOutput)
}

type ConditionOutput struct{ *pulumi.OutputState }

func (ConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Condition)(nil)).Elem()
}

func (o ConditionOutput) ToConditionOutput() ConditionOutput {
	return o
}

func (o ConditionOutput) ToConditionOutputWithContext(ctx context.Context) ConditionOutput {
	return o
}

func (o ConditionOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Field }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Condition) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ConditionArrayOutput struct{ *pulumi.OutputState }

func (ConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Condition)(nil)).Elem()
}

func (o ConditionArrayOutput) ToConditionArrayOutput() ConditionArrayOutput {
	return o
}

func (o ConditionArrayOutput) ToConditionArrayOutputWithContext(ctx context.Context) ConditionArrayOutput {
	return o
}

func (o ConditionArrayOutput) Index(i pulumi.IntInput) ConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Condition {
		return vs[0].([]Condition)[vs[1].(int)]
	}).(ConditionOutput)
}

type ConditionResponse struct {
	Field    *string  `pulumi:"field"`
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}

type ConditionResponseOutput struct{ *pulumi.OutputState }

func (ConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseOutput) ToConditionResponseOutput() ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return o
}

func (o ConditionResponseOutput) Field() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Field }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConditionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (ConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return o
}

func (o ConditionResponseArrayOutput) Index(i pulumi.IntInput) ConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConditionResponse {
		return vs[0].([]ConditionResponse)[vs[1].(int)]
	}).(ConditionResponseOutput)
}

type DailyRecurrence struct {
	EndTime        string `pulumi:"endTime"`
	RecurrenceType string `pulumi:"recurrenceType"`
	StartTime      string `pulumi:"startTime"`
}

type DailyRecurrenceResponse struct {
	EndTime        string `pulumi:"endTime"`
	RecurrenceType string `pulumi:"recurrenceType"`
	StartTime      string `pulumi:"startTime"`
}

type MonthlyRecurrence struct {
	DaysOfMonth    []int   `pulumi:"daysOfMonth"`
	EndTime        *string `pulumi:"endTime"`
	RecurrenceType string  `pulumi:"recurrenceType"`
	StartTime      *string `pulumi:"startTime"`
}

type MonthlyRecurrenceResponse struct {
	DaysOfMonth    []int   `pulumi:"daysOfMonth"`
	EndTime        *string `pulumi:"endTime"`
	RecurrenceType string  `pulumi:"recurrenceType"`
	StartTime      *string `pulumi:"startTime"`
}

type RemoveAllActionGroups struct {
	ActionType string `pulumi:"actionType"`
}

type RemoveAllActionGroupsResponse struct {
	ActionType string `pulumi:"actionType"`
}

type Schedule struct {
	EffectiveFrom  *string       `pulumi:"effectiveFrom"`
	EffectiveUntil *string       `pulumi:"effectiveUntil"`
	Recurrences    []interface{} `pulumi:"recurrences"`
	TimeZone       *string       `pulumi:"timeZone"`
}





type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

type ScheduleArgs struct {
	EffectiveFrom  pulumi.StringPtrInput `pulumi:"effectiveFrom"`
	EffectiveUntil pulumi.StringPtrInput `pulumi:"effectiveUntil"`
	Recurrences    pulumi.ArrayInput     `pulumi:"recurrences"`
	TimeZone       pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i ScheduleArgs) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput).ToSchedulePtrOutputWithContext(ctx)
}









type SchedulePtrInput interface {
	pulumi.Input

	ToSchedulePtrOutput() SchedulePtrOutput
	ToSchedulePtrOutputWithContext(context.Context) SchedulePtrOutput
}

type schedulePtrType ScheduleArgs

func SchedulePtr(v *ScheduleArgs) SchedulePtrInput {
	return (*schedulePtrType)(v)
}

func (*schedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *schedulePtrType) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *schedulePtrType) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o.ToSchedulePtrOutputWithContext(context.Background())
}

func (o ScheduleOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schedule) *Schedule {
		return &v
	}).(SchedulePtrOutput)
}

func (o ScheduleOutput) EffectiveFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schedule) *string { return v.EffectiveFrom }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) EffectiveUntil() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schedule) *string { return v.EffectiveUntil }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) Recurrences() pulumi.ArrayOutput {
	return o.ApplyT(func(v Schedule) []interface{} { return v.Recurrences }).(pulumi.ArrayOutput)
}

func (o ScheduleOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schedule) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type SchedulePtrOutput struct{ *pulumi.OutputState }

func (SchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o SchedulePtrOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) Elem() ScheduleOutput {
	return o.ApplyT(func(v *Schedule) Schedule {
		if v != nil {
			return *v
		}
		var ret Schedule
		return ret
	}).(ScheduleOutput)
}

func (o SchedulePtrOutput) EffectiveFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) *string {
		if v == nil {
			return nil
		}
		return v.EffectiveFrom
	}).(pulumi.StringPtrOutput)
}

func (o SchedulePtrOutput) EffectiveUntil() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) *string {
		if v == nil {
			return nil
		}
		return v.EffectiveUntil
	}).(pulumi.StringPtrOutput)
}

func (o SchedulePtrOutput) Recurrences() pulumi.ArrayOutput {
	return o.ApplyT(func(v *Schedule) []interface{} {
		if v == nil {
			return nil
		}
		return v.Recurrences
	}).(pulumi.ArrayOutput)
}

func (o SchedulePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
	}).(pulumi.StringPtrOutput)
}

type ScheduleResponse struct {
	EffectiveFrom  *string       `pulumi:"effectiveFrom"`
	EffectiveUntil *string       `pulumi:"effectiveUntil"`
	Recurrences    []interface{} `pulumi:"recurrences"`
	TimeZone       *string       `pulumi:"timeZone"`
}

type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) EffectiveFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.EffectiveFrom }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) EffectiveUntil() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.EffectiveUntil }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) Recurrences() pulumi.ArrayOutput {
	return o.ApplyT(func(v ScheduleResponse) []interface{} { return v.Recurrences }).(pulumi.ArrayOutput)
}

func (o ScheduleResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type ScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) Elem() ScheduleResponseOutput {
	return o.ApplyT(func(v *ScheduleResponse) ScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleResponse
		return ret
	}).(ScheduleResponseOutput)
}

func (o ScheduleResponsePtrOutput) EffectiveFrom() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.EffectiveFrom
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) EffectiveUntil() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.EffectiveUntil
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Recurrences() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ScheduleResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Recurrences
	}).(pulumi.ArrayOutput)
}

func (o ScheduleResponsePtrOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeZone
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

type WeeklyRecurrence struct {
	DaysOfWeek     []string `pulumi:"daysOfWeek"`
	EndTime        *string  `pulumi:"endTime"`
	RecurrenceType string   `pulumi:"recurrenceType"`
	StartTime      *string  `pulumi:"startTime"`
}

type WeeklyRecurrenceResponse struct {
	DaysOfWeek     []string `pulumi:"daysOfWeek"`
	EndTime        *string  `pulumi:"endTime"`
	RecurrenceType string   `pulumi:"recurrenceType"`
	StartTime      *string  `pulumi:"startTime"`
}

func init() {
	pulumi.RegisterOutputType(AlertProcessingRulePropertiesOutput{})
	pulumi.RegisterOutputType(AlertProcessingRulePropertiesPtrOutput{})
	pulumi.RegisterOutputType(AlertProcessingRulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionArrayOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
