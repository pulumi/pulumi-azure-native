


package v20210808preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionType string

const (
	ActionTypeAddActionGroups       = ActionType("AddActionGroups")
	ActionTypeRemoveAllActionGroups = ActionType("RemoveAllActionGroups")
)

func (ActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (e ActionType) ToActionTypeOutput() ActionTypeOutput {
	return pulumi.ToOutput(e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ActionTypeOutput)
}

func (e ActionType) ToActionTypePtrOutput() ActionTypePtrOutput {
	return e.ToActionTypePtrOutputWithContext(context.Background())
}

func (e ActionType) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return ActionType(e).ToActionTypeOutputWithContext(ctx).ToActionTypePtrOutputWithContext(ctx)
}

func (e ActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ActionTypeOutput struct{ *pulumi.OutputState }

func (ActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionType)(nil)).Elem()
}

func (o ActionTypeOutput) ToActionTypeOutput() ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypeOutputWithContext(ctx context.Context) ActionTypeOutput {
	return o
}

func (o ActionTypeOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o.ToActionTypePtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionType) *ActionType {
		return &v
	}).(ActionTypePtrOutput)
}

func (o ActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ActionTypePtrOutput struct{ *pulumi.OutputState }

func (ActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionType)(nil)).Elem()
}

func (o ActionTypePtrOutput) ToActionTypePtrOutput() ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return o
}

func (o ActionTypePtrOutput) Elem() ActionTypeOutput {
	return o.ApplyT(func(v *ActionType) ActionType {
		if v != nil {
			return *v
		}
		var ret ActionType
		return ret
	}).(ActionTypeOutput)
}

func (o ActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ActionTypeInput interface {
	pulumi.Input

	ToActionTypeOutput() ActionTypeOutput
	ToActionTypeOutputWithContext(context.Context) ActionTypeOutput
}

var actionTypePtrType = reflect.TypeOf((**ActionType)(nil)).Elem()

type ActionTypePtrInput interface {
	pulumi.Input

	ToActionTypePtrOutput() ActionTypePtrOutput
	ToActionTypePtrOutputWithContext(context.Context) ActionTypePtrOutput
}

type actionTypePtr string

func ActionTypePtr(v string) ActionTypePtrInput {
	return (*actionTypePtr)(&v)
}

func (*actionTypePtr) ElementType() reflect.Type {
	return actionTypePtrType
}

func (in *actionTypePtr) ToActionTypePtrOutput() ActionTypePtrOutput {
	return pulumi.ToOutput(in).(ActionTypePtrOutput)
}

func (in *actionTypePtr) ToActionTypePtrOutputWithContext(ctx context.Context) ActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ActionTypePtrOutput)
}

type DaysOfWeek string

const (
	DaysOfWeekSunday    = DaysOfWeek("Sunday")
	DaysOfWeekMonday    = DaysOfWeek("Monday")
	DaysOfWeekTuesday   = DaysOfWeek("Tuesday")
	DaysOfWeekWednesday = DaysOfWeek("Wednesday")
	DaysOfWeekThursday  = DaysOfWeek("Thursday")
	DaysOfWeekFriday    = DaysOfWeek("Friday")
	DaysOfWeekSaturday  = DaysOfWeek("Saturday")
)

func (DaysOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DaysOfWeek)(nil)).Elem()
}

func (e DaysOfWeek) ToDaysOfWeekOutput() DaysOfWeekOutput {
	return pulumi.ToOutput(e).(DaysOfWeekOutput)
}

func (e DaysOfWeek) ToDaysOfWeekOutputWithContext(ctx context.Context) DaysOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DaysOfWeekOutput)
}

func (e DaysOfWeek) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return e.ToDaysOfWeekPtrOutputWithContext(context.Background())
}

func (e DaysOfWeek) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return DaysOfWeek(e).ToDaysOfWeekOutputWithContext(ctx).ToDaysOfWeekPtrOutputWithContext(ctx)
}

func (e DaysOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DaysOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DaysOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DaysOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DaysOfWeekOutput struct{ *pulumi.OutputState }

func (DaysOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekOutput) ToDaysOfWeekOutput() DaysOfWeekOutput {
	return o
}

func (o DaysOfWeekOutput) ToDaysOfWeekOutputWithContext(ctx context.Context) DaysOfWeekOutput {
	return o
}

func (o DaysOfWeekOutput) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return o.ToDaysOfWeekPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DaysOfWeek) *DaysOfWeek {
		return &v
	}).(DaysOfWeekPtrOutput)
}

func (o DaysOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DaysOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DaysOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DaysOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DaysOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DaysOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaysOfWeek)(nil)).Elem()
}

func (o DaysOfWeekPtrOutput) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return o
}

func (o DaysOfWeekPtrOutput) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return o
}

func (o DaysOfWeekPtrOutput) Elem() DaysOfWeekOutput {
	return o.ApplyT(func(v *DaysOfWeek) DaysOfWeek {
		if v != nil {
			return *v
		}
		var ret DaysOfWeek
		return ret
	}).(DaysOfWeekOutput)
}

func (o DaysOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DaysOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DaysOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DaysOfWeekInput interface {
	pulumi.Input

	ToDaysOfWeekOutput() DaysOfWeekOutput
	ToDaysOfWeekOutputWithContext(context.Context) DaysOfWeekOutput
}

var daysOfWeekPtrType = reflect.TypeOf((**DaysOfWeek)(nil)).Elem()

type DaysOfWeekPtrInput interface {
	pulumi.Input

	ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput
	ToDaysOfWeekPtrOutputWithContext(context.Context) DaysOfWeekPtrOutput
}

type daysOfWeekPtr string

func DaysOfWeekPtr(v string) DaysOfWeekPtrInput {
	return (*daysOfWeekPtr)(&v)
}

func (*daysOfWeekPtr) ElementType() reflect.Type {
	return daysOfWeekPtrType
}

func (in *daysOfWeekPtr) ToDaysOfWeekPtrOutput() DaysOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DaysOfWeekPtrOutput)
}

func (in *daysOfWeekPtr) ToDaysOfWeekPtrOutputWithContext(ctx context.Context) DaysOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DaysOfWeekPtrOutput)
}

type Field string

const (
	FieldSeverity            = Field("Severity")
	FieldMonitorService      = Field("MonitorService")
	FieldMonitorCondition    = Field("MonitorCondition")
	FieldSignalType          = Field("SignalType")
	FieldTargetResourceType  = Field("TargetResourceType")
	FieldTargetResource      = Field("TargetResource")
	FieldTargetResourceGroup = Field("TargetResourceGroup")
	FieldAlertRuleId         = Field("AlertRuleId")
	FieldAlertRuleName       = Field("AlertRuleName")
	FieldDescription         = Field("Description")
	FieldAlertContext        = Field("AlertContext")
)

func (Field) ElementType() reflect.Type {
	return reflect.TypeOf((*Field)(nil)).Elem()
}

func (e Field) ToFieldOutput() FieldOutput {
	return pulumi.ToOutput(e).(FieldOutput)
}

func (e Field) ToFieldOutputWithContext(ctx context.Context) FieldOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FieldOutput)
}

func (e Field) ToFieldPtrOutput() FieldPtrOutput {
	return e.ToFieldPtrOutputWithContext(context.Background())
}

func (e Field) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return Field(e).ToFieldOutputWithContext(ctx).ToFieldPtrOutputWithContext(ctx)
}

func (e Field) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Field) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Field) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Field) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FieldOutput struct{ *pulumi.OutputState }

func (FieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Field)(nil)).Elem()
}

func (o FieldOutput) ToFieldOutput() FieldOutput {
	return o
}

func (o FieldOutput) ToFieldOutputWithContext(ctx context.Context) FieldOutput {
	return o
}

func (o FieldOutput) ToFieldPtrOutput() FieldPtrOutput {
	return o.ToFieldPtrOutputWithContext(context.Background())
}

func (o FieldOutput) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Field) *Field {
		return &v
	}).(FieldPtrOutput)
}

func (o FieldOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FieldOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Field) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FieldOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FieldOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Field) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FieldPtrOutput struct{ *pulumi.OutputState }

func (FieldPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Field)(nil)).Elem()
}

func (o FieldPtrOutput) ToFieldPtrOutput() FieldPtrOutput {
	return o
}

func (o FieldPtrOutput) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return o
}

func (o FieldPtrOutput) Elem() FieldOutput {
	return o.ApplyT(func(v *Field) Field {
		if v != nil {
			return *v
		}
		var ret Field
		return ret
	}).(FieldOutput)
}

func (o FieldPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FieldPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Field) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FieldInput interface {
	pulumi.Input

	ToFieldOutput() FieldOutput
	ToFieldOutputWithContext(context.Context) FieldOutput
}

var fieldPtrType = reflect.TypeOf((**Field)(nil)).Elem()

type FieldPtrInput interface {
	pulumi.Input

	ToFieldPtrOutput() FieldPtrOutput
	ToFieldPtrOutputWithContext(context.Context) FieldPtrOutput
}

type fieldPtr string

func FieldPtr(v string) FieldPtrInput {
	return (*fieldPtr)(&v)
}

func (*fieldPtr) ElementType() reflect.Type {
	return fieldPtrType
}

func (in *fieldPtr) ToFieldPtrOutput() FieldPtrOutput {
	return pulumi.ToOutput(in).(FieldPtrOutput)
}

func (in *fieldPtr) ToFieldPtrOutputWithContext(ctx context.Context) FieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FieldPtrOutput)
}

type Operator string

const (
	OperatorEquals         = Operator("Equals")
	OperatorNotEquals      = Operator("NotEquals")
	OperatorContains       = Operator("Contains")
	OperatorDoesNotContain = Operator("DoesNotContain")
)

func (Operator) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (e Operator) ToOperatorOutput() OperatorOutput {
	return pulumi.ToOutput(e).(OperatorOutput)
}

func (e Operator) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorOutput)
}

func (e Operator) ToOperatorPtrOutput() OperatorPtrOutput {
	return e.ToOperatorPtrOutputWithContext(context.Background())
}

func (e Operator) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return Operator(e).ToOperatorOutputWithContext(ctx).ToOperatorPtrOutputWithContext(ctx)
}

func (e Operator) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Operator) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Operator) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorOutput struct{ *pulumi.OutputState }

func (OperatorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Operator)(nil)).Elem()
}

func (o OperatorOutput) ToOperatorOutput() OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorOutputWithContext(ctx context.Context) OperatorOutput {
	return o
}

func (o OperatorOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o.ToOperatorPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Operator) *Operator {
		return &v
	}).(OperatorPtrOutput)
}

func (o OperatorOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Operator) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorPtrOutput struct{ *pulumi.OutputState }

func (OperatorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Operator)(nil)).Elem()
}

func (o OperatorPtrOutput) ToOperatorPtrOutput() OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return o
}

func (o OperatorPtrOutput) Elem() OperatorOutput {
	return o.ApplyT(func(v *Operator) Operator {
		if v != nil {
			return *v
		}
		var ret Operator
		return ret
	}).(OperatorOutput)
}

func (o OperatorPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Operator) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatorInput interface {
	pulumi.Input

	ToOperatorOutput() OperatorOutput
	ToOperatorOutputWithContext(context.Context) OperatorOutput
}

var operatorPtrType = reflect.TypeOf((**Operator)(nil)).Elem()

type OperatorPtrInput interface {
	pulumi.Input

	ToOperatorPtrOutput() OperatorPtrOutput
	ToOperatorPtrOutputWithContext(context.Context) OperatorPtrOutput
}

type operatorPtr string

func OperatorPtr(v string) OperatorPtrInput {
	return (*operatorPtr)(&v)
}

func (*operatorPtr) ElementType() reflect.Type {
	return operatorPtrType
}

func (in *operatorPtr) ToOperatorPtrOutput() OperatorPtrOutput {
	return pulumi.ToOutput(in).(OperatorPtrOutput)
}

func (in *operatorPtr) ToOperatorPtrOutputWithContext(ctx context.Context) OperatorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorPtrOutput)
}

type RecurrenceType string

const (
	RecurrenceTypeDaily   = RecurrenceType("Daily")
	RecurrenceTypeWeekly  = RecurrenceType("Weekly")
	RecurrenceTypeMonthly = RecurrenceType("Monthly")
)

func (RecurrenceType) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceType)(nil)).Elem()
}

func (e RecurrenceType) ToRecurrenceTypeOutput() RecurrenceTypeOutput {
	return pulumi.ToOutput(e).(RecurrenceTypeOutput)
}

func (e RecurrenceType) ToRecurrenceTypeOutputWithContext(ctx context.Context) RecurrenceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecurrenceTypeOutput)
}

func (e RecurrenceType) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return e.ToRecurrenceTypePtrOutputWithContext(context.Background())
}

func (e RecurrenceType) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return RecurrenceType(e).ToRecurrenceTypeOutputWithContext(ctx).ToRecurrenceTypePtrOutputWithContext(ctx)
}

func (e RecurrenceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurrenceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecurrenceTypeOutput struct{ *pulumi.OutputState }

func (RecurrenceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceType)(nil)).Elem()
}

func (o RecurrenceTypeOutput) ToRecurrenceTypeOutput() RecurrenceTypeOutput {
	return o
}

func (o RecurrenceTypeOutput) ToRecurrenceTypeOutputWithContext(ctx context.Context) RecurrenceTypeOutput {
	return o
}

func (o RecurrenceTypeOutput) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return o.ToRecurrenceTypePtrOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceType) *RecurrenceType {
		return &v
	}).(RecurrenceTypePtrOutput)
}

func (o RecurrenceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecurrenceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecurrenceTypePtrOutput struct{ *pulumi.OutputState }

func (RecurrenceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceType)(nil)).Elem()
}

func (o RecurrenceTypePtrOutput) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return o
}

func (o RecurrenceTypePtrOutput) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return o
}

func (o RecurrenceTypePtrOutput) Elem() RecurrenceTypeOutput {
	return o.ApplyT(func(v *RecurrenceType) RecurrenceType {
		if v != nil {
			return *v
		}
		var ret RecurrenceType
		return ret
	}).(RecurrenceTypeOutput)
}

func (o RecurrenceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecurrenceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecurrenceTypeInput interface {
	pulumi.Input

	ToRecurrenceTypeOutput() RecurrenceTypeOutput
	ToRecurrenceTypeOutputWithContext(context.Context) RecurrenceTypeOutput
}

var recurrenceTypePtrType = reflect.TypeOf((**RecurrenceType)(nil)).Elem()

type RecurrenceTypePtrInput interface {
	pulumi.Input

	ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput
	ToRecurrenceTypePtrOutputWithContext(context.Context) RecurrenceTypePtrOutput
}

type recurrenceTypePtr string

func RecurrenceTypePtr(v string) RecurrenceTypePtrInput {
	return (*recurrenceTypePtr)(&v)
}

func (*recurrenceTypePtr) ElementType() reflect.Type {
	return recurrenceTypePtrType
}

func (in *recurrenceTypePtr) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return pulumi.ToOutput(in).(RecurrenceTypePtrOutput)
}

func (in *recurrenceTypePtr) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecurrenceTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionTypeOutput{})
	pulumi.RegisterOutputType(ActionTypePtrOutput{})
	pulumi.RegisterOutputType(DaysOfWeekOutput{})
	pulumi.RegisterOutputType(DaysOfWeekPtrOutput{})
	pulumi.RegisterOutputType(FieldOutput{})
	pulumi.RegisterOutputType(FieldPtrOutput{})
	pulumi.RegisterOutputType(OperatorOutput{})
	pulumi.RegisterOutputType(OperatorPtrOutput{})
	pulumi.RegisterOutputType(RecurrenceTypeOutput{})
	pulumi.RegisterOutputType(RecurrenceTypePtrOutput{})
}
