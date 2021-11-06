


package v20210808preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionRuleProperties struct {
	Actions     []interface{} `pulumi:"actions"`
	Conditions  []Condition   `pulumi:"conditions"`
	Description *string       `pulumi:"description"`
	Enabled     *bool         `pulumi:"enabled"`
	Schedule    *Schedule     `pulumi:"schedule"`
	Scopes      []string      `pulumi:"scopes"`
}





type ActionRulePropertiesInput interface {
	pulumi.Input

	ToActionRulePropertiesOutput() ActionRulePropertiesOutput
	ToActionRulePropertiesOutputWithContext(context.Context) ActionRulePropertiesOutput
}

type ActionRulePropertiesArgs struct {
	Actions     pulumi.ArrayInput       `pulumi:"actions"`
	Conditions  ConditionArrayInput     `pulumi:"conditions"`
	Description pulumi.StringPtrInput   `pulumi:"description"`
	Enabled     pulumi.BoolPtrInput     `pulumi:"enabled"`
	Schedule    SchedulePtrInput        `pulumi:"schedule"`
	Scopes      pulumi.StringArrayInput `pulumi:"scopes"`
}

func (ActionRulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleProperties)(nil)).Elem()
}

func (i ActionRulePropertiesArgs) ToActionRulePropertiesOutput() ActionRulePropertiesOutput {
	return i.ToActionRulePropertiesOutputWithContext(context.Background())
}

func (i ActionRulePropertiesArgs) ToActionRulePropertiesOutputWithContext(ctx context.Context) ActionRulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRulePropertiesOutput)
}

func (i ActionRulePropertiesArgs) ToActionRulePropertiesPtrOutput() ActionRulePropertiesPtrOutput {
	return i.ToActionRulePropertiesPtrOutputWithContext(context.Background())
}

func (i ActionRulePropertiesArgs) ToActionRulePropertiesPtrOutputWithContext(ctx context.Context) ActionRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRulePropertiesOutput).ToActionRulePropertiesPtrOutputWithContext(ctx)
}









type ActionRulePropertiesPtrInput interface {
	pulumi.Input

	ToActionRulePropertiesPtrOutput() ActionRulePropertiesPtrOutput
	ToActionRulePropertiesPtrOutputWithContext(context.Context) ActionRulePropertiesPtrOutput
}

type actionRulePropertiesPtrType ActionRulePropertiesArgs

func ActionRulePropertiesPtr(v *ActionRulePropertiesArgs) ActionRulePropertiesPtrInput {
	return (*actionRulePropertiesPtrType)(v)
}

func (*actionRulePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleProperties)(nil)).Elem()
}

func (i *actionRulePropertiesPtrType) ToActionRulePropertiesPtrOutput() ActionRulePropertiesPtrOutput {
	return i.ToActionRulePropertiesPtrOutputWithContext(context.Background())
}

func (i *actionRulePropertiesPtrType) ToActionRulePropertiesPtrOutputWithContext(ctx context.Context) ActionRulePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRulePropertiesPtrOutput)
}

type ActionRulePropertiesOutput struct{ *pulumi.OutputState }

func (ActionRulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRuleProperties)(nil)).Elem()
}

func (o ActionRulePropertiesOutput) ToActionRulePropertiesOutput() ActionRulePropertiesOutput {
	return o
}

func (o ActionRulePropertiesOutput) ToActionRulePropertiesOutputWithContext(ctx context.Context) ActionRulePropertiesOutput {
	return o
}

func (o ActionRulePropertiesOutput) ToActionRulePropertiesPtrOutput() ActionRulePropertiesPtrOutput {
	return o.ToActionRulePropertiesPtrOutputWithContext(context.Background())
}

func (o ActionRulePropertiesOutput) ToActionRulePropertiesPtrOutputWithContext(ctx context.Context) ActionRulePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionRuleProperties) *ActionRuleProperties {
		return &v
	}).(ActionRulePropertiesPtrOutput)
}

func (o ActionRulePropertiesOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v ActionRuleProperties) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o ActionRulePropertiesOutput) Conditions() ConditionArrayOutput {
	return o.ApplyT(func(v ActionRuleProperties) []Condition { return v.Conditions }).(ConditionArrayOutput)
}

func (o ActionRulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ActionRulePropertiesOutput) Schedule() SchedulePtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *Schedule { return v.Schedule }).(SchedulePtrOutput)
}

func (o ActionRulePropertiesOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionRuleProperties) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ActionRulePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ActionRulePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRuleProperties)(nil)).Elem()
}

func (o ActionRulePropertiesPtrOutput) ToActionRulePropertiesPtrOutput() ActionRulePropertiesPtrOutput {
	return o
}

func (o ActionRulePropertiesPtrOutput) ToActionRulePropertiesPtrOutputWithContext(ctx context.Context) ActionRulePropertiesPtrOutput {
	return o
}

func (o ActionRulePropertiesPtrOutput) Elem() ActionRulePropertiesOutput {
	return o.ApplyT(func(v *ActionRuleProperties) ActionRuleProperties {
		if v != nil {
			return *v
		}
		var ret ActionRuleProperties
		return ret
	}).(ActionRulePropertiesOutput)
}

func (o ActionRulePropertiesPtrOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ActionRuleProperties) []interface{} {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(pulumi.ArrayOutput)
}

func (o ActionRulePropertiesPtrOutput) Conditions() ConditionArrayOutput {
	return o.ApplyT(func(v *ActionRuleProperties) []Condition {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(ConditionArrayOutput)
}

func (o ActionRulePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ActionRulePropertiesPtrOutput) Schedule() SchedulePtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *Schedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(SchedulePtrOutput)
}

func (o ActionRulePropertiesPtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionRuleProperties) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type ActionRulePropertiesResponse struct {
	Actions     []interface{}       `pulumi:"actions"`
	Conditions  []ConditionResponse `pulumi:"conditions"`
	Description *string             `pulumi:"description"`
	Enabled     *bool               `pulumi:"enabled"`
	Schedule    *ScheduleResponse   `pulumi:"schedule"`
	Scopes      []string            `pulumi:"scopes"`
}





type ActionRulePropertiesResponseInput interface {
	pulumi.Input

	ToActionRulePropertiesResponseOutput() ActionRulePropertiesResponseOutput
	ToActionRulePropertiesResponseOutputWithContext(context.Context) ActionRulePropertiesResponseOutput
}

type ActionRulePropertiesResponseArgs struct {
	Actions     pulumi.ArrayInput           `pulumi:"actions"`
	Conditions  ConditionResponseArrayInput `pulumi:"conditions"`
	Description pulumi.StringPtrInput       `pulumi:"description"`
	Enabled     pulumi.BoolPtrInput         `pulumi:"enabled"`
	Schedule    ScheduleResponsePtrInput    `pulumi:"schedule"`
	Scopes      pulumi.StringArrayInput     `pulumi:"scopes"`
}

func (ActionRulePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRulePropertiesResponse)(nil)).Elem()
}

func (i ActionRulePropertiesResponseArgs) ToActionRulePropertiesResponseOutput() ActionRulePropertiesResponseOutput {
	return i.ToActionRulePropertiesResponseOutputWithContext(context.Background())
}

func (i ActionRulePropertiesResponseArgs) ToActionRulePropertiesResponseOutputWithContext(ctx context.Context) ActionRulePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRulePropertiesResponseOutput)
}

func (i ActionRulePropertiesResponseArgs) ToActionRulePropertiesResponsePtrOutput() ActionRulePropertiesResponsePtrOutput {
	return i.ToActionRulePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ActionRulePropertiesResponseArgs) ToActionRulePropertiesResponsePtrOutputWithContext(ctx context.Context) ActionRulePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRulePropertiesResponseOutput).ToActionRulePropertiesResponsePtrOutputWithContext(ctx)
}









type ActionRulePropertiesResponsePtrInput interface {
	pulumi.Input

	ToActionRulePropertiesResponsePtrOutput() ActionRulePropertiesResponsePtrOutput
	ToActionRulePropertiesResponsePtrOutputWithContext(context.Context) ActionRulePropertiesResponsePtrOutput
}

type actionRulePropertiesResponsePtrType ActionRulePropertiesResponseArgs

func ActionRulePropertiesResponsePtr(v *ActionRulePropertiesResponseArgs) ActionRulePropertiesResponsePtrInput {
	return (*actionRulePropertiesResponsePtrType)(v)
}

func (*actionRulePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRulePropertiesResponse)(nil)).Elem()
}

func (i *actionRulePropertiesResponsePtrType) ToActionRulePropertiesResponsePtrOutput() ActionRulePropertiesResponsePtrOutput {
	return i.ToActionRulePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *actionRulePropertiesResponsePtrType) ToActionRulePropertiesResponsePtrOutputWithContext(ctx context.Context) ActionRulePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionRulePropertiesResponsePtrOutput)
}

type ActionRulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ActionRulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionRulePropertiesResponse)(nil)).Elem()
}

func (o ActionRulePropertiesResponseOutput) ToActionRulePropertiesResponseOutput() ActionRulePropertiesResponseOutput {
	return o
}

func (o ActionRulePropertiesResponseOutput) ToActionRulePropertiesResponseOutputWithContext(ctx context.Context) ActionRulePropertiesResponseOutput {
	return o
}

func (o ActionRulePropertiesResponseOutput) ToActionRulePropertiesResponsePtrOutput() ActionRulePropertiesResponsePtrOutput {
	return o.ToActionRulePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ActionRulePropertiesResponseOutput) ToActionRulePropertiesResponsePtrOutputWithContext(ctx context.Context) ActionRulePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionRulePropertiesResponse) *ActionRulePropertiesResponse {
		return &v
	}).(ActionRulePropertiesResponsePtrOutput)
}

func (o ActionRulePropertiesResponseOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o ActionRulePropertiesResponseOutput) Conditions() ConditionResponseArrayOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) []ConditionResponse { return v.Conditions }).(ConditionResponseArrayOutput)
}

func (o ActionRulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ActionRulePropertiesResponseOutput) Schedule() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *ScheduleResponse { return v.Schedule }).(ScheduleResponsePtrOutput)
}

func (o ActionRulePropertiesResponseOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

type ActionRulePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionRulePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionRulePropertiesResponse)(nil)).Elem()
}

func (o ActionRulePropertiesResponsePtrOutput) ToActionRulePropertiesResponsePtrOutput() ActionRulePropertiesResponsePtrOutput {
	return o
}

func (o ActionRulePropertiesResponsePtrOutput) ToActionRulePropertiesResponsePtrOutputWithContext(ctx context.Context) ActionRulePropertiesResponsePtrOutput {
	return o
}

func (o ActionRulePropertiesResponsePtrOutput) Elem() ActionRulePropertiesResponseOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) ActionRulePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ActionRulePropertiesResponse
		return ret
	}).(ActionRulePropertiesResponseOutput)
}

func (o ActionRulePropertiesResponsePtrOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(pulumi.ArrayOutput)
}

func (o ActionRulePropertiesResponsePtrOutput) Conditions() ConditionResponseArrayOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) []ConditionResponse {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(ConditionResponseArrayOutput)
}

func (o ActionRulePropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o ActionRulePropertiesResponsePtrOutput) Schedule() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) *ScheduleResponse {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(ScheduleResponsePtrOutput)
}

func (o ActionRulePropertiesResponsePtrOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionRulePropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringArrayOutput)
}

type AddActionGroups struct {
	ActionGroupIds []string `pulumi:"actionGroupIds"`
	ActionType     string   `pulumi:"actionType"`
}





type AddActionGroupsInput interface {
	pulumi.Input

	ToAddActionGroupsOutput() AddActionGroupsOutput
	ToAddActionGroupsOutputWithContext(context.Context) AddActionGroupsOutput
}

type AddActionGroupsArgs struct {
	ActionGroupIds pulumi.StringArrayInput `pulumi:"actionGroupIds"`
	ActionType     pulumi.StringInput      `pulumi:"actionType"`
}

func (AddActionGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddActionGroups)(nil)).Elem()
}

func (i AddActionGroupsArgs) ToAddActionGroupsOutput() AddActionGroupsOutput {
	return i.ToAddActionGroupsOutputWithContext(context.Background())
}

func (i AddActionGroupsArgs) ToAddActionGroupsOutputWithContext(ctx context.Context) AddActionGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddActionGroupsOutput)
}

type AddActionGroupsOutput struct{ *pulumi.OutputState }

func (AddActionGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddActionGroups)(nil)).Elem()
}

func (o AddActionGroupsOutput) ToAddActionGroupsOutput() AddActionGroupsOutput {
	return o
}

func (o AddActionGroupsOutput) ToAddActionGroupsOutputWithContext(ctx context.Context) AddActionGroupsOutput {
	return o
}

func (o AddActionGroupsOutput) ActionGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddActionGroups) []string { return v.ActionGroupIds }).(pulumi.StringArrayOutput)
}

func (o AddActionGroupsOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AddActionGroups) string { return v.ActionType }).(pulumi.StringOutput)
}

type AddActionGroupsResponse struct {
	ActionGroupIds []string `pulumi:"actionGroupIds"`
	ActionType     string   `pulumi:"actionType"`
}





type AddActionGroupsResponseInput interface {
	pulumi.Input

	ToAddActionGroupsResponseOutput() AddActionGroupsResponseOutput
	ToAddActionGroupsResponseOutputWithContext(context.Context) AddActionGroupsResponseOutput
}

type AddActionGroupsResponseArgs struct {
	ActionGroupIds pulumi.StringArrayInput `pulumi:"actionGroupIds"`
	ActionType     pulumi.StringInput      `pulumi:"actionType"`
}

func (AddActionGroupsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AddActionGroupsResponse)(nil)).Elem()
}

func (i AddActionGroupsResponseArgs) ToAddActionGroupsResponseOutput() AddActionGroupsResponseOutput {
	return i.ToAddActionGroupsResponseOutputWithContext(context.Background())
}

func (i AddActionGroupsResponseArgs) ToAddActionGroupsResponseOutputWithContext(ctx context.Context) AddActionGroupsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AddActionGroupsResponseOutput)
}

type AddActionGroupsResponseOutput struct{ *pulumi.OutputState }

func (AddActionGroupsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AddActionGroupsResponse)(nil)).Elem()
}

func (o AddActionGroupsResponseOutput) ToAddActionGroupsResponseOutput() AddActionGroupsResponseOutput {
	return o
}

func (o AddActionGroupsResponseOutput) ToAddActionGroupsResponseOutputWithContext(ctx context.Context) AddActionGroupsResponseOutput {
	return o
}

func (o AddActionGroupsResponseOutput) ActionGroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AddActionGroupsResponse) []string { return v.ActionGroupIds }).(pulumi.StringArrayOutput)
}

func (o AddActionGroupsResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AddActionGroupsResponse) string { return v.ActionType }).(pulumi.StringOutput)
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





type ConditionResponseInput interface {
	pulumi.Input

	ToConditionResponseOutput() ConditionResponseOutput
	ToConditionResponseOutputWithContext(context.Context) ConditionResponseOutput
}

type ConditionResponseArgs struct {
	Field    pulumi.StringPtrInput   `pulumi:"field"`
	Operator pulumi.StringPtrInput   `pulumi:"operator"`
	Values   pulumi.StringArrayInput `pulumi:"values"`
}

func (ConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionResponse)(nil)).Elem()
}

func (i ConditionResponseArgs) ToConditionResponseOutput() ConditionResponseOutput {
	return i.ToConditionResponseOutputWithContext(context.Background())
}

func (i ConditionResponseArgs) ToConditionResponseOutputWithContext(ctx context.Context) ConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseOutput)
}





type ConditionResponseArrayInput interface {
	pulumi.Input

	ToConditionResponseArrayOutput() ConditionResponseArrayOutput
	ToConditionResponseArrayOutputWithContext(context.Context) ConditionResponseArrayOutput
}

type ConditionResponseArray []ConditionResponseInput

func (ConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConditionResponse)(nil)).Elem()
}

func (i ConditionResponseArray) ToConditionResponseArrayOutput() ConditionResponseArrayOutput {
	return i.ToConditionResponseArrayOutputWithContext(context.Background())
}

func (i ConditionResponseArray) ToConditionResponseArrayOutputWithContext(ctx context.Context) ConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseArrayOutput)
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





type DailyRecurrenceInput interface {
	pulumi.Input

	ToDailyRecurrenceOutput() DailyRecurrenceOutput
	ToDailyRecurrenceOutputWithContext(context.Context) DailyRecurrenceOutput
}

type DailyRecurrenceArgs struct {
	EndTime        pulumi.StringInput `pulumi:"endTime"`
	RecurrenceType pulumi.StringInput `pulumi:"recurrenceType"`
	StartTime      pulumi.StringInput `pulumi:"startTime"`
}

func (DailyRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRecurrence)(nil)).Elem()
}

func (i DailyRecurrenceArgs) ToDailyRecurrenceOutput() DailyRecurrenceOutput {
	return i.ToDailyRecurrenceOutputWithContext(context.Background())
}

func (i DailyRecurrenceArgs) ToDailyRecurrenceOutputWithContext(ctx context.Context) DailyRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRecurrenceOutput)
}

type DailyRecurrenceOutput struct{ *pulumi.OutputState }

func (DailyRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRecurrence)(nil)).Elem()
}

func (o DailyRecurrenceOutput) ToDailyRecurrenceOutput() DailyRecurrenceOutput {
	return o
}

func (o DailyRecurrenceOutput) ToDailyRecurrenceOutputWithContext(ctx context.Context) DailyRecurrenceOutput {
	return o
}

func (o DailyRecurrenceOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v DailyRecurrence) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o DailyRecurrenceOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v DailyRecurrence) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o DailyRecurrenceOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v DailyRecurrence) string { return v.StartTime }).(pulumi.StringOutput)
}

type DailyRecurrenceResponse struct {
	EndTime        string `pulumi:"endTime"`
	RecurrenceType string `pulumi:"recurrenceType"`
	StartTime      string `pulumi:"startTime"`
}





type DailyRecurrenceResponseInput interface {
	pulumi.Input

	ToDailyRecurrenceResponseOutput() DailyRecurrenceResponseOutput
	ToDailyRecurrenceResponseOutputWithContext(context.Context) DailyRecurrenceResponseOutput
}

type DailyRecurrenceResponseArgs struct {
	EndTime        pulumi.StringInput `pulumi:"endTime"`
	RecurrenceType pulumi.StringInput `pulumi:"recurrenceType"`
	StartTime      pulumi.StringInput `pulumi:"startTime"`
}

func (DailyRecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRecurrenceResponse)(nil)).Elem()
}

func (i DailyRecurrenceResponseArgs) ToDailyRecurrenceResponseOutput() DailyRecurrenceResponseOutput {
	return i.ToDailyRecurrenceResponseOutputWithContext(context.Background())
}

func (i DailyRecurrenceResponseArgs) ToDailyRecurrenceResponseOutputWithContext(ctx context.Context) DailyRecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRecurrenceResponseOutput)
}

type DailyRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (DailyRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRecurrenceResponse)(nil)).Elem()
}

func (o DailyRecurrenceResponseOutput) ToDailyRecurrenceResponseOutput() DailyRecurrenceResponseOutput {
	return o
}

func (o DailyRecurrenceResponseOutput) ToDailyRecurrenceResponseOutputWithContext(ctx context.Context) DailyRecurrenceResponseOutput {
	return o
}

func (o DailyRecurrenceResponseOutput) EndTime() pulumi.StringOutput {
	return o.ApplyT(func(v DailyRecurrenceResponse) string { return v.EndTime }).(pulumi.StringOutput)
}

func (o DailyRecurrenceResponseOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v DailyRecurrenceResponse) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o DailyRecurrenceResponseOutput) StartTime() pulumi.StringOutput {
	return o.ApplyT(func(v DailyRecurrenceResponse) string { return v.StartTime }).(pulumi.StringOutput)
}

type MonthlyRecurrence struct {
	DaysOfMonth    []int   `pulumi:"daysOfMonth"`
	EndTime        *string `pulumi:"endTime"`
	RecurrenceType string  `pulumi:"recurrenceType"`
	StartTime      *string `pulumi:"startTime"`
}





type MonthlyRecurrenceInput interface {
	pulumi.Input

	ToMonthlyRecurrenceOutput() MonthlyRecurrenceOutput
	ToMonthlyRecurrenceOutputWithContext(context.Context) MonthlyRecurrenceOutput
}

type MonthlyRecurrenceArgs struct {
	DaysOfMonth    pulumi.IntArrayInput  `pulumi:"daysOfMonth"`
	EndTime        pulumi.StringPtrInput `pulumi:"endTime"`
	RecurrenceType pulumi.StringInput    `pulumi:"recurrenceType"`
	StartTime      pulumi.StringPtrInput `pulumi:"startTime"`
}

func (MonthlyRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRecurrence)(nil)).Elem()
}

func (i MonthlyRecurrenceArgs) ToMonthlyRecurrenceOutput() MonthlyRecurrenceOutput {
	return i.ToMonthlyRecurrenceOutputWithContext(context.Background())
}

func (i MonthlyRecurrenceArgs) ToMonthlyRecurrenceOutputWithContext(ctx context.Context) MonthlyRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRecurrenceOutput)
}

type MonthlyRecurrenceOutput struct{ *pulumi.OutputState }

func (MonthlyRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRecurrence)(nil)).Elem()
}

func (o MonthlyRecurrenceOutput) ToMonthlyRecurrenceOutput() MonthlyRecurrenceOutput {
	return o
}

func (o MonthlyRecurrenceOutput) ToMonthlyRecurrenceOutputWithContext(ctx context.Context) MonthlyRecurrenceOutput {
	return o
}

func (o MonthlyRecurrenceOutput) DaysOfMonth() pulumi.IntArrayOutput {
	return o.ApplyT(func(v MonthlyRecurrence) []int { return v.DaysOfMonth }).(pulumi.IntArrayOutput)
}

func (o MonthlyRecurrenceOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRecurrence) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o MonthlyRecurrenceOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v MonthlyRecurrence) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o MonthlyRecurrenceOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRecurrence) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type MonthlyRecurrenceResponse struct {
	DaysOfMonth    []int   `pulumi:"daysOfMonth"`
	EndTime        *string `pulumi:"endTime"`
	RecurrenceType string  `pulumi:"recurrenceType"`
	StartTime      *string `pulumi:"startTime"`
}





type MonthlyRecurrenceResponseInput interface {
	pulumi.Input

	ToMonthlyRecurrenceResponseOutput() MonthlyRecurrenceResponseOutput
	ToMonthlyRecurrenceResponseOutputWithContext(context.Context) MonthlyRecurrenceResponseOutput
}

type MonthlyRecurrenceResponseArgs struct {
	DaysOfMonth    pulumi.IntArrayInput  `pulumi:"daysOfMonth"`
	EndTime        pulumi.StringPtrInput `pulumi:"endTime"`
	RecurrenceType pulumi.StringInput    `pulumi:"recurrenceType"`
	StartTime      pulumi.StringPtrInput `pulumi:"startTime"`
}

func (MonthlyRecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRecurrenceResponse)(nil)).Elem()
}

func (i MonthlyRecurrenceResponseArgs) ToMonthlyRecurrenceResponseOutput() MonthlyRecurrenceResponseOutput {
	return i.ToMonthlyRecurrenceResponseOutputWithContext(context.Background())
}

func (i MonthlyRecurrenceResponseArgs) ToMonthlyRecurrenceResponseOutputWithContext(ctx context.Context) MonthlyRecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRecurrenceResponseOutput)
}

type MonthlyRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (MonthlyRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRecurrenceResponse)(nil)).Elem()
}

func (o MonthlyRecurrenceResponseOutput) ToMonthlyRecurrenceResponseOutput() MonthlyRecurrenceResponseOutput {
	return o
}

func (o MonthlyRecurrenceResponseOutput) ToMonthlyRecurrenceResponseOutputWithContext(ctx context.Context) MonthlyRecurrenceResponseOutput {
	return o
}

func (o MonthlyRecurrenceResponseOutput) DaysOfMonth() pulumi.IntArrayOutput {
	return o.ApplyT(func(v MonthlyRecurrenceResponse) []int { return v.DaysOfMonth }).(pulumi.IntArrayOutput)
}

func (o MonthlyRecurrenceResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRecurrenceResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o MonthlyRecurrenceResponseOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v MonthlyRecurrenceResponse) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o MonthlyRecurrenceResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRecurrenceResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type RemoveAllActionGroups struct {
	ActionType string `pulumi:"actionType"`
}





type RemoveAllActionGroupsInput interface {
	pulumi.Input

	ToRemoveAllActionGroupsOutput() RemoveAllActionGroupsOutput
	ToRemoveAllActionGroupsOutputWithContext(context.Context) RemoveAllActionGroupsOutput
}

type RemoveAllActionGroupsArgs struct {
	ActionType pulumi.StringInput `pulumi:"actionType"`
}

func (RemoveAllActionGroupsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoveAllActionGroups)(nil)).Elem()
}

func (i RemoveAllActionGroupsArgs) ToRemoveAllActionGroupsOutput() RemoveAllActionGroupsOutput {
	return i.ToRemoveAllActionGroupsOutputWithContext(context.Background())
}

func (i RemoveAllActionGroupsArgs) ToRemoveAllActionGroupsOutputWithContext(ctx context.Context) RemoveAllActionGroupsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoveAllActionGroupsOutput)
}

type RemoveAllActionGroupsOutput struct{ *pulumi.OutputState }

func (RemoveAllActionGroupsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoveAllActionGroups)(nil)).Elem()
}

func (o RemoveAllActionGroupsOutput) ToRemoveAllActionGroupsOutput() RemoveAllActionGroupsOutput {
	return o
}

func (o RemoveAllActionGroupsOutput) ToRemoveAllActionGroupsOutputWithContext(ctx context.Context) RemoveAllActionGroupsOutput {
	return o
}

func (o RemoveAllActionGroupsOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v RemoveAllActionGroups) string { return v.ActionType }).(pulumi.StringOutput)
}

type RemoveAllActionGroupsResponse struct {
	ActionType string `pulumi:"actionType"`
}





type RemoveAllActionGroupsResponseInput interface {
	pulumi.Input

	ToRemoveAllActionGroupsResponseOutput() RemoveAllActionGroupsResponseOutput
	ToRemoveAllActionGroupsResponseOutputWithContext(context.Context) RemoveAllActionGroupsResponseOutput
}

type RemoveAllActionGroupsResponseArgs struct {
	ActionType pulumi.StringInput `pulumi:"actionType"`
}

func (RemoveAllActionGroupsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoveAllActionGroupsResponse)(nil)).Elem()
}

func (i RemoveAllActionGroupsResponseArgs) ToRemoveAllActionGroupsResponseOutput() RemoveAllActionGroupsResponseOutput {
	return i.ToRemoveAllActionGroupsResponseOutputWithContext(context.Background())
}

func (i RemoveAllActionGroupsResponseArgs) ToRemoveAllActionGroupsResponseOutputWithContext(ctx context.Context) RemoveAllActionGroupsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RemoveAllActionGroupsResponseOutput)
}

type RemoveAllActionGroupsResponseOutput struct{ *pulumi.OutputState }

func (RemoveAllActionGroupsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemoveAllActionGroupsResponse)(nil)).Elem()
}

func (o RemoveAllActionGroupsResponseOutput) ToRemoveAllActionGroupsResponseOutput() RemoveAllActionGroupsResponseOutput {
	return o
}

func (o RemoveAllActionGroupsResponseOutput) ToRemoveAllActionGroupsResponseOutputWithContext(ctx context.Context) RemoveAllActionGroupsResponseOutput {
	return o
}

func (o RemoveAllActionGroupsResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v RemoveAllActionGroupsResponse) string { return v.ActionType }).(pulumi.StringOutput)
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





type ScheduleResponseInput interface {
	pulumi.Input

	ToScheduleResponseOutput() ScheduleResponseOutput
	ToScheduleResponseOutputWithContext(context.Context) ScheduleResponseOutput
}

type ScheduleResponseArgs struct {
	EffectiveFrom  pulumi.StringPtrInput `pulumi:"effectiveFrom"`
	EffectiveUntil pulumi.StringPtrInput `pulumi:"effectiveUntil"`
	Recurrences    pulumi.ArrayInput     `pulumi:"recurrences"`
	TimeZone       pulumi.StringPtrInput `pulumi:"timeZone"`
}

func (ScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (i ScheduleResponseArgs) ToScheduleResponseOutput() ScheduleResponseOutput {
	return i.ToScheduleResponseOutputWithContext(context.Background())
}

func (i ScheduleResponseArgs) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResponseOutput)
}

func (i ScheduleResponseArgs) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return i.ToScheduleResponsePtrOutputWithContext(context.Background())
}

func (i ScheduleResponseArgs) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResponseOutput).ToScheduleResponsePtrOutputWithContext(ctx)
}









type ScheduleResponsePtrInput interface {
	pulumi.Input

	ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput
	ToScheduleResponsePtrOutputWithContext(context.Context) ScheduleResponsePtrOutput
}

type scheduleResponsePtrType ScheduleResponseArgs

func ScheduleResponsePtr(v *ScheduleResponseArgs) ScheduleResponsePtrInput {
	return (*scheduleResponsePtrType)(v)
}

func (*scheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (i *scheduleResponsePtrType) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return i.ToScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *scheduleResponsePtrType) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleResponsePtrOutput)
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

func (o ScheduleResponseOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o.ToScheduleResponsePtrOutputWithContext(context.Background())
}

func (o ScheduleResponseOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleResponse) *ScheduleResponse {
		return &v
	}).(ScheduleResponsePtrOutput)
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

type WeeklyRecurrence struct {
	DaysOfWeek     []string `pulumi:"daysOfWeek"`
	EndTime        *string  `pulumi:"endTime"`
	RecurrenceType string   `pulumi:"recurrenceType"`
	StartTime      *string  `pulumi:"startTime"`
}





type WeeklyRecurrenceInput interface {
	pulumi.Input

	ToWeeklyRecurrenceOutput() WeeklyRecurrenceOutput
	ToWeeklyRecurrenceOutputWithContext(context.Context) WeeklyRecurrenceOutput
}

type WeeklyRecurrenceArgs struct {
	DaysOfWeek     pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	EndTime        pulumi.StringPtrInput   `pulumi:"endTime"`
	RecurrenceType pulumi.StringInput      `pulumi:"recurrenceType"`
	StartTime      pulumi.StringPtrInput   `pulumi:"startTime"`
}

func (WeeklyRecurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRecurrence)(nil)).Elem()
}

func (i WeeklyRecurrenceArgs) ToWeeklyRecurrenceOutput() WeeklyRecurrenceOutput {
	return i.ToWeeklyRecurrenceOutputWithContext(context.Background())
}

func (i WeeklyRecurrenceArgs) ToWeeklyRecurrenceOutputWithContext(ctx context.Context) WeeklyRecurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRecurrenceOutput)
}

type WeeklyRecurrenceOutput struct{ *pulumi.OutputState }

func (WeeklyRecurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRecurrence)(nil)).Elem()
}

func (o WeeklyRecurrenceOutput) ToWeeklyRecurrenceOutput() WeeklyRecurrenceOutput {
	return o
}

func (o WeeklyRecurrenceOutput) ToWeeklyRecurrenceOutputWithContext(ctx context.Context) WeeklyRecurrenceOutput {
	return o
}

func (o WeeklyRecurrenceOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRecurrence) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o WeeklyRecurrenceOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeeklyRecurrence) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o WeeklyRecurrenceOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v WeeklyRecurrence) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o WeeklyRecurrenceOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeeklyRecurrence) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type WeeklyRecurrenceResponse struct {
	DaysOfWeek     []string `pulumi:"daysOfWeek"`
	EndTime        *string  `pulumi:"endTime"`
	RecurrenceType string   `pulumi:"recurrenceType"`
	StartTime      *string  `pulumi:"startTime"`
}





type WeeklyRecurrenceResponseInput interface {
	pulumi.Input

	ToWeeklyRecurrenceResponseOutput() WeeklyRecurrenceResponseOutput
	ToWeeklyRecurrenceResponseOutputWithContext(context.Context) WeeklyRecurrenceResponseOutput
}

type WeeklyRecurrenceResponseArgs struct {
	DaysOfWeek     pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	EndTime        pulumi.StringPtrInput   `pulumi:"endTime"`
	RecurrenceType pulumi.StringInput      `pulumi:"recurrenceType"`
	StartTime      pulumi.StringPtrInput   `pulumi:"startTime"`
}

func (WeeklyRecurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRecurrenceResponse)(nil)).Elem()
}

func (i WeeklyRecurrenceResponseArgs) ToWeeklyRecurrenceResponseOutput() WeeklyRecurrenceResponseOutput {
	return i.ToWeeklyRecurrenceResponseOutputWithContext(context.Background())
}

func (i WeeklyRecurrenceResponseArgs) ToWeeklyRecurrenceResponseOutputWithContext(ctx context.Context) WeeklyRecurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRecurrenceResponseOutput)
}

type WeeklyRecurrenceResponseOutput struct{ *pulumi.OutputState }

func (WeeklyRecurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRecurrenceResponse)(nil)).Elem()
}

func (o WeeklyRecurrenceResponseOutput) ToWeeklyRecurrenceResponseOutput() WeeklyRecurrenceResponseOutput {
	return o
}

func (o WeeklyRecurrenceResponseOutput) ToWeeklyRecurrenceResponseOutputWithContext(ctx context.Context) WeeklyRecurrenceResponseOutput {
	return o
}

func (o WeeklyRecurrenceResponseOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRecurrenceResponse) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o WeeklyRecurrenceResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeeklyRecurrenceResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o WeeklyRecurrenceResponseOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v WeeklyRecurrenceResponse) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o WeeklyRecurrenceResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WeeklyRecurrenceResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionRulePropertiesOutput{})
	pulumi.RegisterOutputType(ActionRulePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ActionRulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ActionRulePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AddActionGroupsOutput{})
	pulumi.RegisterOutputType(AddActionGroupsResponseOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionArrayOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(DailyRecurrenceOutput{})
	pulumi.RegisterOutputType(DailyRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(MonthlyRecurrenceOutput{})
	pulumi.RegisterOutputType(MonthlyRecurrenceResponseOutput{})
	pulumi.RegisterOutputType(RemoveAllActionGroupsOutput{})
	pulumi.RegisterOutputType(RemoveAllActionGroupsResponseOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(WeeklyRecurrenceOutput{})
	pulumi.RegisterOutputType(WeeklyRecurrenceResponseOutput{})
}
