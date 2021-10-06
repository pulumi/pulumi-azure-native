


package v20190505preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionGroup struct {
	ActionGroupId string      `pulumi:"actionGroupId"`
	Conditions    *Conditions `pulumi:"conditions"`
	Description   *string     `pulumi:"description"`
	Scope         *Scope      `pulumi:"scope"`
	Status        *string     `pulumi:"status"`
	Type          string      `pulumi:"type"`
}





type ActionGroupInput interface {
	pulumi.Input

	ToActionGroupOutput() ActionGroupOutput
	ToActionGroupOutputWithContext(context.Context) ActionGroupOutput
}

type ActionGroupArgs struct {
	ActionGroupId pulumi.StringInput    `pulumi:"actionGroupId"`
	Conditions    ConditionsPtrInput    `pulumi:"conditions"`
	Description   pulumi.StringPtrInput `pulumi:"description"`
	Scope         ScopePtrInput         `pulumi:"scope"`
	Status        pulumi.StringPtrInput `pulumi:"status"`
	Type          pulumi.StringInput    `pulumi:"type"`
}

func (ActionGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil)).Elem()
}

func (i ActionGroupArgs) ToActionGroupOutput() ActionGroupOutput {
	return i.ToActionGroupOutputWithContext(context.Background())
}

func (i ActionGroupArgs) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupOutput)
}

type ActionGroupOutput struct{ *pulumi.OutputState }

func (ActionGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroup)(nil)).Elem()
}

func (o ActionGroupOutput) ToActionGroupOutput() ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ToActionGroupOutputWithContext(ctx context.Context) ActionGroupOutput {
	return o
}

func (o ActionGroupOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroup) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActionGroupOutput) Conditions() ConditionsPtrOutput {
	return o.ApplyT(func(v ActionGroup) *Conditions { return v.Conditions }).(ConditionsPtrOutput)
}

func (o ActionGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroup) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActionGroupOutput) Scope() ScopePtrOutput {
	return o.ApplyT(func(v ActionGroup) *Scope { return v.Scope }).(ScopePtrOutput)
}

func (o ActionGroupOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroup) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ActionGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroup) string { return v.Type }).(pulumi.StringOutput)
}

type ActionGroupResponse struct {
	ActionGroupId  string              `pulumi:"actionGroupId"`
	Conditions     *ConditionsResponse `pulumi:"conditions"`
	CreatedAt      string              `pulumi:"createdAt"`
	CreatedBy      string              `pulumi:"createdBy"`
	Description    *string             `pulumi:"description"`
	LastModifiedAt string              `pulumi:"lastModifiedAt"`
	LastModifiedBy string              `pulumi:"lastModifiedBy"`
	Scope          *ScopeResponse      `pulumi:"scope"`
	Status         *string             `pulumi:"status"`
	Type           string              `pulumi:"type"`
}





type ActionGroupResponseInput interface {
	pulumi.Input

	ToActionGroupResponseOutput() ActionGroupResponseOutput
	ToActionGroupResponseOutputWithContext(context.Context) ActionGroupResponseOutput
}

type ActionGroupResponseArgs struct {
	ActionGroupId  pulumi.StringInput         `pulumi:"actionGroupId"`
	Conditions     ConditionsResponsePtrInput `pulumi:"conditions"`
	CreatedAt      pulumi.StringInput         `pulumi:"createdAt"`
	CreatedBy      pulumi.StringInput         `pulumi:"createdBy"`
	Description    pulumi.StringPtrInput      `pulumi:"description"`
	LastModifiedAt pulumi.StringInput         `pulumi:"lastModifiedAt"`
	LastModifiedBy pulumi.StringInput         `pulumi:"lastModifiedBy"`
	Scope          ScopeResponsePtrInput      `pulumi:"scope"`
	Status         pulumi.StringPtrInput      `pulumi:"status"`
	Type           pulumi.StringInput         `pulumi:"type"`
}

func (ActionGroupResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupResponse)(nil)).Elem()
}

func (i ActionGroupResponseArgs) ToActionGroupResponseOutput() ActionGroupResponseOutput {
	return i.ToActionGroupResponseOutputWithContext(context.Background())
}

func (i ActionGroupResponseArgs) ToActionGroupResponseOutputWithContext(ctx context.Context) ActionGroupResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupResponseOutput)
}

type ActionGroupResponseOutput struct{ *pulumi.OutputState }

func (ActionGroupResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupResponse)(nil)).Elem()
}

func (o ActionGroupResponseOutput) ToActionGroupResponseOutput() ActionGroupResponseOutput {
	return o
}

func (o ActionGroupResponseOutput) ToActionGroupResponseOutputWithContext(ctx context.Context) ActionGroupResponseOutput {
	return o
}

func (o ActionGroupResponseOutput) ActionGroupId() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.ActionGroupId }).(pulumi.StringOutput)
}

func (o ActionGroupResponseOutput) Conditions() ConditionsResponsePtrOutput {
	return o.ApplyT(func(v ActionGroupResponse) *ConditionsResponse { return v.Conditions }).(ConditionsResponsePtrOutput)
}

func (o ActionGroupResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ActionGroupResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o ActionGroupResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActionGroupResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o ActionGroupResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o ActionGroupResponseOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v ActionGroupResponse) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o ActionGroupResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ActionGroupResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ActionGroupResponse) string { return v.Type }).(pulumi.StringOutput)
}

type Condition struct {
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ConditionInput interface {
	pulumi.Input

	ToConditionOutput() ConditionOutput
	ToConditionOutputWithContext(context.Context) ConditionOutput
}

type ConditionArgs struct {
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

func (i ConditionArgs) ToConditionPtrOutput() ConditionPtrOutput {
	return i.ToConditionPtrOutputWithContext(context.Background())
}

func (i ConditionArgs) ToConditionPtrOutputWithContext(ctx context.Context) ConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionOutput).ToConditionPtrOutputWithContext(ctx)
}









type ConditionPtrInput interface {
	pulumi.Input

	ToConditionPtrOutput() ConditionPtrOutput
	ToConditionPtrOutputWithContext(context.Context) ConditionPtrOutput
}

type conditionPtrType ConditionArgs

func ConditionPtr(v *ConditionArgs) ConditionPtrInput {
	return (*conditionPtrType)(v)
}

func (*conditionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Condition)(nil)).Elem()
}

func (i *conditionPtrType) ToConditionPtrOutput() ConditionPtrOutput {
	return i.ToConditionPtrOutputWithContext(context.Background())
}

func (i *conditionPtrType) ToConditionPtrOutputWithContext(ctx context.Context) ConditionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionPtrOutput)
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

func (o ConditionOutput) ToConditionPtrOutput() ConditionPtrOutput {
	return o.ToConditionPtrOutputWithContext(context.Background())
}

func (o ConditionOutput) ToConditionPtrOutputWithContext(ctx context.Context) ConditionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Condition) *Condition {
		return &v
	}).(ConditionPtrOutput)
}

func (o ConditionOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Condition) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ConditionOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Condition) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ConditionPtrOutput struct{ *pulumi.OutputState }

func (ConditionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Condition)(nil)).Elem()
}

func (o ConditionPtrOutput) ToConditionPtrOutput() ConditionPtrOutput {
	return o
}

func (o ConditionPtrOutput) ToConditionPtrOutputWithContext(ctx context.Context) ConditionPtrOutput {
	return o
}

func (o ConditionPtrOutput) Elem() ConditionOutput {
	return o.ApplyT(func(v *Condition) Condition {
		if v != nil {
			return *v
		}
		var ret Condition
		return ret
	}).(ConditionOutput)
}

func (o ConditionPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Condition) *string {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ConditionPtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Condition) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ConditionResponse struct {
	Operator *string  `pulumi:"operator"`
	Values   []string `pulumi:"values"`
}





type ConditionResponseInput interface {
	pulumi.Input

	ToConditionResponseOutput() ConditionResponseOutput
	ToConditionResponseOutputWithContext(context.Context) ConditionResponseOutput
}

type ConditionResponseArgs struct {
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

func (i ConditionResponseArgs) ToConditionResponsePtrOutput() ConditionResponsePtrOutput {
	return i.ToConditionResponsePtrOutputWithContext(context.Background())
}

func (i ConditionResponseArgs) ToConditionResponsePtrOutputWithContext(ctx context.Context) ConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponseOutput).ToConditionResponsePtrOutputWithContext(ctx)
}









type ConditionResponsePtrInput interface {
	pulumi.Input

	ToConditionResponsePtrOutput() ConditionResponsePtrOutput
	ToConditionResponsePtrOutputWithContext(context.Context) ConditionResponsePtrOutput
}

type conditionResponsePtrType ConditionResponseArgs

func ConditionResponsePtr(v *ConditionResponseArgs) ConditionResponsePtrInput {
	return (*conditionResponsePtrType)(v)
}

func (*conditionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionResponse)(nil)).Elem()
}

func (i *conditionResponsePtrType) ToConditionResponsePtrOutput() ConditionResponsePtrOutput {
	return i.ToConditionResponsePtrOutputWithContext(context.Background())
}

func (i *conditionResponsePtrType) ToConditionResponsePtrOutputWithContext(ctx context.Context) ConditionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionResponsePtrOutput)
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

func (o ConditionResponseOutput) ToConditionResponsePtrOutput() ConditionResponsePtrOutput {
	return o.ToConditionResponsePtrOutputWithContext(context.Background())
}

func (o ConditionResponseOutput) ToConditionResponsePtrOutputWithContext(ctx context.Context) ConditionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionResponse) *ConditionResponse {
		return &v
	}).(ConditionResponsePtrOutput)
}

func (o ConditionResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConditionResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o ConditionResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ConditionResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ConditionResponsePtrOutput struct{ *pulumi.OutputState }

func (ConditionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionResponse)(nil)).Elem()
}

func (o ConditionResponsePtrOutput) ToConditionResponsePtrOutput() ConditionResponsePtrOutput {
	return o
}

func (o ConditionResponsePtrOutput) ToConditionResponsePtrOutputWithContext(ctx context.Context) ConditionResponsePtrOutput {
	return o
}

func (o ConditionResponsePtrOutput) Elem() ConditionResponseOutput {
	return o.ApplyT(func(v *ConditionResponse) ConditionResponse {
		if v != nil {
			return *v
		}
		var ret ConditionResponse
		return ret
	}).(ConditionResponseOutput)
}

func (o ConditionResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConditionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o ConditionResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ConditionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type Conditions struct {
	AlertContext       *Condition `pulumi:"alertContext"`
	AlertRuleId        *Condition `pulumi:"alertRuleId"`
	Description        *Condition `pulumi:"description"`
	MonitorCondition   *Condition `pulumi:"monitorCondition"`
	MonitorService     *Condition `pulumi:"monitorService"`
	Severity           *Condition `pulumi:"severity"`
	TargetResourceType *Condition `pulumi:"targetResourceType"`
}





type ConditionsInput interface {
	pulumi.Input

	ToConditionsOutput() ConditionsOutput
	ToConditionsOutputWithContext(context.Context) ConditionsOutput
}

type ConditionsArgs struct {
	AlertContext       ConditionPtrInput `pulumi:"alertContext"`
	AlertRuleId        ConditionPtrInput `pulumi:"alertRuleId"`
	Description        ConditionPtrInput `pulumi:"description"`
	MonitorCondition   ConditionPtrInput `pulumi:"monitorCondition"`
	MonitorService     ConditionPtrInput `pulumi:"monitorService"`
	Severity           ConditionPtrInput `pulumi:"severity"`
	TargetResourceType ConditionPtrInput `pulumi:"targetResourceType"`
}

func (ConditionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Conditions)(nil)).Elem()
}

func (i ConditionsArgs) ToConditionsOutput() ConditionsOutput {
	return i.ToConditionsOutputWithContext(context.Background())
}

func (i ConditionsArgs) ToConditionsOutputWithContext(ctx context.Context) ConditionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionsOutput)
}

func (i ConditionsArgs) ToConditionsPtrOutput() ConditionsPtrOutput {
	return i.ToConditionsPtrOutputWithContext(context.Background())
}

func (i ConditionsArgs) ToConditionsPtrOutputWithContext(ctx context.Context) ConditionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionsOutput).ToConditionsPtrOutputWithContext(ctx)
}









type ConditionsPtrInput interface {
	pulumi.Input

	ToConditionsPtrOutput() ConditionsPtrOutput
	ToConditionsPtrOutputWithContext(context.Context) ConditionsPtrOutput
}

type conditionsPtrType ConditionsArgs

func ConditionsPtr(v *ConditionsArgs) ConditionsPtrInput {
	return (*conditionsPtrType)(v)
}

func (*conditionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Conditions)(nil)).Elem()
}

func (i *conditionsPtrType) ToConditionsPtrOutput() ConditionsPtrOutput {
	return i.ToConditionsPtrOutputWithContext(context.Background())
}

func (i *conditionsPtrType) ToConditionsPtrOutputWithContext(ctx context.Context) ConditionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionsPtrOutput)
}

type ConditionsOutput struct{ *pulumi.OutputState }

func (ConditionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Conditions)(nil)).Elem()
}

func (o ConditionsOutput) ToConditionsOutput() ConditionsOutput {
	return o
}

func (o ConditionsOutput) ToConditionsOutputWithContext(ctx context.Context) ConditionsOutput {
	return o
}

func (o ConditionsOutput) ToConditionsPtrOutput() ConditionsPtrOutput {
	return o.ToConditionsPtrOutputWithContext(context.Background())
}

func (o ConditionsOutput) ToConditionsPtrOutputWithContext(ctx context.Context) ConditionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Conditions) *Conditions {
		return &v
	}).(ConditionsPtrOutput)
}

func (o ConditionsOutput) AlertContext() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.AlertContext }).(ConditionPtrOutput)
}

func (o ConditionsOutput) AlertRuleId() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.AlertRuleId }).(ConditionPtrOutput)
}

func (o ConditionsOutput) Description() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.Description }).(ConditionPtrOutput)
}

func (o ConditionsOutput) MonitorCondition() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.MonitorCondition }).(ConditionPtrOutput)
}

func (o ConditionsOutput) MonitorService() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.MonitorService }).(ConditionPtrOutput)
}

func (o ConditionsOutput) Severity() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.Severity }).(ConditionPtrOutput)
}

func (o ConditionsOutput) TargetResourceType() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.TargetResourceType }).(ConditionPtrOutput)
}

type ConditionsPtrOutput struct{ *pulumi.OutputState }

func (ConditionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Conditions)(nil)).Elem()
}

func (o ConditionsPtrOutput) ToConditionsPtrOutput() ConditionsPtrOutput {
	return o
}

func (o ConditionsPtrOutput) ToConditionsPtrOutputWithContext(ctx context.Context) ConditionsPtrOutput {
	return o
}

func (o ConditionsPtrOutput) Elem() ConditionsOutput {
	return o.ApplyT(func(v *Conditions) Conditions {
		if v != nil {
			return *v
		}
		var ret Conditions
		return ret
	}).(ConditionsOutput)
}

func (o ConditionsPtrOutput) AlertContext() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.AlertContext
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) AlertRuleId() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.AlertRuleId
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) Description() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.Description
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) MonitorCondition() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.MonitorCondition
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) MonitorService() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.MonitorService
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) Severity() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.Severity
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) TargetResourceType() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.TargetResourceType
	}).(ConditionPtrOutput)
}

type ConditionsResponse struct {
	AlertContext       *ConditionResponse `pulumi:"alertContext"`
	AlertRuleId        *ConditionResponse `pulumi:"alertRuleId"`
	Description        *ConditionResponse `pulumi:"description"`
	MonitorCondition   *ConditionResponse `pulumi:"monitorCondition"`
	MonitorService     *ConditionResponse `pulumi:"monitorService"`
	Severity           *ConditionResponse `pulumi:"severity"`
	TargetResourceType *ConditionResponse `pulumi:"targetResourceType"`
}





type ConditionsResponseInput interface {
	pulumi.Input

	ToConditionsResponseOutput() ConditionsResponseOutput
	ToConditionsResponseOutputWithContext(context.Context) ConditionsResponseOutput
}

type ConditionsResponseArgs struct {
	AlertContext       ConditionResponsePtrInput `pulumi:"alertContext"`
	AlertRuleId        ConditionResponsePtrInput `pulumi:"alertRuleId"`
	Description        ConditionResponsePtrInput `pulumi:"description"`
	MonitorCondition   ConditionResponsePtrInput `pulumi:"monitorCondition"`
	MonitorService     ConditionResponsePtrInput `pulumi:"monitorService"`
	Severity           ConditionResponsePtrInput `pulumi:"severity"`
	TargetResourceType ConditionResponsePtrInput `pulumi:"targetResourceType"`
}

func (ConditionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionsResponse)(nil)).Elem()
}

func (i ConditionsResponseArgs) ToConditionsResponseOutput() ConditionsResponseOutput {
	return i.ToConditionsResponseOutputWithContext(context.Background())
}

func (i ConditionsResponseArgs) ToConditionsResponseOutputWithContext(ctx context.Context) ConditionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionsResponseOutput)
}

func (i ConditionsResponseArgs) ToConditionsResponsePtrOutput() ConditionsResponsePtrOutput {
	return i.ToConditionsResponsePtrOutputWithContext(context.Background())
}

func (i ConditionsResponseArgs) ToConditionsResponsePtrOutputWithContext(ctx context.Context) ConditionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionsResponseOutput).ToConditionsResponsePtrOutputWithContext(ctx)
}









type ConditionsResponsePtrInput interface {
	pulumi.Input

	ToConditionsResponsePtrOutput() ConditionsResponsePtrOutput
	ToConditionsResponsePtrOutputWithContext(context.Context) ConditionsResponsePtrOutput
}

type conditionsResponsePtrType ConditionsResponseArgs

func ConditionsResponsePtr(v *ConditionsResponseArgs) ConditionsResponsePtrInput {
	return (*conditionsResponsePtrType)(v)
}

func (*conditionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionsResponse)(nil)).Elem()
}

func (i *conditionsResponsePtrType) ToConditionsResponsePtrOutput() ConditionsResponsePtrOutput {
	return i.ToConditionsResponsePtrOutputWithContext(context.Background())
}

func (i *conditionsResponsePtrType) ToConditionsResponsePtrOutputWithContext(ctx context.Context) ConditionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConditionsResponsePtrOutput)
}

type ConditionsResponseOutput struct{ *pulumi.OutputState }

func (ConditionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConditionsResponse)(nil)).Elem()
}

func (o ConditionsResponseOutput) ToConditionsResponseOutput() ConditionsResponseOutput {
	return o
}

func (o ConditionsResponseOutput) ToConditionsResponseOutputWithContext(ctx context.Context) ConditionsResponseOutput {
	return o
}

func (o ConditionsResponseOutput) ToConditionsResponsePtrOutput() ConditionsResponsePtrOutput {
	return o.ToConditionsResponsePtrOutputWithContext(context.Background())
}

func (o ConditionsResponseOutput) ToConditionsResponsePtrOutputWithContext(ctx context.Context) ConditionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConditionsResponse) *ConditionsResponse {
		return &v
	}).(ConditionsResponsePtrOutput)
}

func (o ConditionsResponseOutput) AlertContext() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.AlertContext }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) AlertRuleId() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.AlertRuleId }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) Description() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.Description }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) MonitorCondition() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.MonitorCondition }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) MonitorService() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.MonitorService }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) Severity() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.Severity }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) TargetResourceType() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.TargetResourceType }).(ConditionResponsePtrOutput)
}

type ConditionsResponsePtrOutput struct{ *pulumi.OutputState }

func (ConditionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConditionsResponse)(nil)).Elem()
}

func (o ConditionsResponsePtrOutput) ToConditionsResponsePtrOutput() ConditionsResponsePtrOutput {
	return o
}

func (o ConditionsResponsePtrOutput) ToConditionsResponsePtrOutputWithContext(ctx context.Context) ConditionsResponsePtrOutput {
	return o
}

func (o ConditionsResponsePtrOutput) Elem() ConditionsResponseOutput {
	return o.ApplyT(func(v *ConditionsResponse) ConditionsResponse {
		if v != nil {
			return *v
		}
		var ret ConditionsResponse
		return ret
	}).(ConditionsResponseOutput)
}

func (o ConditionsResponsePtrOutput) AlertContext() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.AlertContext
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) AlertRuleId() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.AlertRuleId
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) Description() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.Description
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) MonitorCondition() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.MonitorCondition
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) MonitorService() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.MonitorService
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) Severity() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.Severity
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) TargetResourceType() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.TargetResourceType
	}).(ConditionResponsePtrOutput)
}

type Diagnostics struct {
	Conditions  *Conditions `pulumi:"conditions"`
	Description *string     `pulumi:"description"`
	Scope       *Scope      `pulumi:"scope"`
	Status      *string     `pulumi:"status"`
	Type        string      `pulumi:"type"`
}





type DiagnosticsInput interface {
	pulumi.Input

	ToDiagnosticsOutput() DiagnosticsOutput
	ToDiagnosticsOutputWithContext(context.Context) DiagnosticsOutput
}

type DiagnosticsArgs struct {
	Conditions  ConditionsPtrInput    `pulumi:"conditions"`
	Description pulumi.StringPtrInput `pulumi:"description"`
	Scope       ScopePtrInput         `pulumi:"scope"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
	Type        pulumi.StringInput    `pulumi:"type"`
}

func (DiagnosticsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Diagnostics)(nil)).Elem()
}

func (i DiagnosticsArgs) ToDiagnosticsOutput() DiagnosticsOutput {
	return i.ToDiagnosticsOutputWithContext(context.Background())
}

func (i DiagnosticsArgs) ToDiagnosticsOutputWithContext(ctx context.Context) DiagnosticsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsOutput)
}

type DiagnosticsOutput struct{ *pulumi.OutputState }

func (DiagnosticsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Diagnostics)(nil)).Elem()
}

func (o DiagnosticsOutput) ToDiagnosticsOutput() DiagnosticsOutput {
	return o
}

func (o DiagnosticsOutput) ToDiagnosticsOutputWithContext(ctx context.Context) DiagnosticsOutput {
	return o
}

func (o DiagnosticsOutput) Conditions() ConditionsPtrOutput {
	return o.ApplyT(func(v Diagnostics) *Conditions { return v.Conditions }).(ConditionsPtrOutput)
}

func (o DiagnosticsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Diagnostics) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsOutput) Scope() ScopePtrOutput {
	return o.ApplyT(func(v Diagnostics) *Scope { return v.Scope }).(ScopePtrOutput)
}

func (o DiagnosticsOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Diagnostics) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Diagnostics) string { return v.Type }).(pulumi.StringOutput)
}

type DiagnosticsResponse struct {
	Conditions     *ConditionsResponse `pulumi:"conditions"`
	CreatedAt      string              `pulumi:"createdAt"`
	CreatedBy      string              `pulumi:"createdBy"`
	Description    *string             `pulumi:"description"`
	LastModifiedAt string              `pulumi:"lastModifiedAt"`
	LastModifiedBy string              `pulumi:"lastModifiedBy"`
	Scope          *ScopeResponse      `pulumi:"scope"`
	Status         *string             `pulumi:"status"`
	Type           string              `pulumi:"type"`
}





type DiagnosticsResponseInput interface {
	pulumi.Input

	ToDiagnosticsResponseOutput() DiagnosticsResponseOutput
	ToDiagnosticsResponseOutputWithContext(context.Context) DiagnosticsResponseOutput
}

type DiagnosticsResponseArgs struct {
	Conditions     ConditionsResponsePtrInput `pulumi:"conditions"`
	CreatedAt      pulumi.StringInput         `pulumi:"createdAt"`
	CreatedBy      pulumi.StringInput         `pulumi:"createdBy"`
	Description    pulumi.StringPtrInput      `pulumi:"description"`
	LastModifiedAt pulumi.StringInput         `pulumi:"lastModifiedAt"`
	LastModifiedBy pulumi.StringInput         `pulumi:"lastModifiedBy"`
	Scope          ScopeResponsePtrInput      `pulumi:"scope"`
	Status         pulumi.StringPtrInput      `pulumi:"status"`
	Type           pulumi.StringInput         `pulumi:"type"`
}

func (DiagnosticsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsResponse)(nil)).Elem()
}

func (i DiagnosticsResponseArgs) ToDiagnosticsResponseOutput() DiagnosticsResponseOutput {
	return i.ToDiagnosticsResponseOutputWithContext(context.Background())
}

func (i DiagnosticsResponseArgs) ToDiagnosticsResponseOutputWithContext(ctx context.Context) DiagnosticsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsResponseOutput)
}

type DiagnosticsResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsResponse)(nil)).Elem()
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutput() DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) ToDiagnosticsResponseOutputWithContext(ctx context.Context) DiagnosticsResponseOutput {
	return o
}

func (o DiagnosticsResponseOutput) Conditions() ConditionsResponsePtrOutput {
	return o.ApplyT(func(v DiagnosticsResponse) *ConditionsResponse { return v.Conditions }).(ConditionsResponsePtrOutput)
}

func (o DiagnosticsResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o DiagnosticsResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o DiagnosticsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o DiagnosticsResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o DiagnosticsResponseOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v DiagnosticsResponse) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o DiagnosticsResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DiagnosticsResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o DiagnosticsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v DiagnosticsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type Scope struct {
	ScopeType *string  `pulumi:"scopeType"`
	Values    []string `pulumi:"values"`
}





type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(context.Context) ScopeOutput
}

type ScopeArgs struct {
	ScopeType pulumi.StringPtrInput   `pulumi:"scopeType"`
	Values    pulumi.StringArrayInput `pulumi:"values"`
}

func (ScopeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (i ScopeArgs) ToScopeOutput() ScopeOutput {
	return i.ToScopeOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput)
}

func (i ScopeArgs) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i ScopeArgs) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeOutput).ToScopePtrOutputWithContext(ctx)
}









type ScopePtrInput interface {
	pulumi.Input

	ToScopePtrOutput() ScopePtrOutput
	ToScopePtrOutputWithContext(context.Context) ScopePtrOutput
}

type scopePtrType ScopeArgs

func ScopePtr(v *ScopeArgs) ScopePtrInput {
	return (*scopePtrType)(v)
}

func (*scopePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (i *scopePtrType) ToScopePtrOutput() ScopePtrOutput {
	return i.ToScopePtrOutputWithContext(context.Background())
}

func (i *scopePtrType) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopePtrOutput)
}

type ScopeOutput struct{ *pulumi.OutputState }

func (ScopeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scope)(nil)).Elem()
}

func (o ScopeOutput) ToScopeOutput() ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopeOutputWithContext(ctx context.Context) ScopeOutput {
	return o
}

func (o ScopeOutput) ToScopePtrOutput() ScopePtrOutput {
	return o.ToScopePtrOutputWithContext(context.Background())
}

func (o ScopeOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scope) *Scope {
		return &v
	}).(ScopePtrOutput)
}

func (o ScopeOutput) ScopeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Scope) *string { return v.ScopeType }).(pulumi.StringPtrOutput)
}

func (o ScopeOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Scope) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ScopePtrOutput struct{ *pulumi.OutputState }

func (ScopePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scope)(nil)).Elem()
}

func (o ScopePtrOutput) ToScopePtrOutput() ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) ToScopePtrOutputWithContext(ctx context.Context) ScopePtrOutput {
	return o
}

func (o ScopePtrOutput) Elem() ScopeOutput {
	return o.ApplyT(func(v *Scope) Scope {
		if v != nil {
			return *v
		}
		var ret Scope
		return ret
	}).(ScopeOutput)
}

func (o ScopePtrOutput) ScopeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scope) *string {
		if v == nil {
			return nil
		}
		return v.ScopeType
	}).(pulumi.StringPtrOutput)
}

func (o ScopePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Scope) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type ScopeResponse struct {
	ScopeType *string  `pulumi:"scopeType"`
	Values    []string `pulumi:"values"`
}





type ScopeResponseInput interface {
	pulumi.Input

	ToScopeResponseOutput() ScopeResponseOutput
	ToScopeResponseOutputWithContext(context.Context) ScopeResponseOutput
}

type ScopeResponseArgs struct {
	ScopeType pulumi.StringPtrInput   `pulumi:"scopeType"`
	Values    pulumi.StringArrayInput `pulumi:"values"`
}

func (ScopeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (i ScopeResponseArgs) ToScopeResponseOutput() ScopeResponseOutput {
	return i.ToScopeResponseOutputWithContext(context.Background())
}

func (i ScopeResponseArgs) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponseOutput)
}

func (i ScopeResponseArgs) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return i.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (i ScopeResponseArgs) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponseOutput).ToScopeResponsePtrOutputWithContext(ctx)
}









type ScopeResponsePtrInput interface {
	pulumi.Input

	ToScopeResponsePtrOutput() ScopeResponsePtrOutput
	ToScopeResponsePtrOutputWithContext(context.Context) ScopeResponsePtrOutput
}

type scopeResponsePtrType ScopeResponseArgs

func ScopeResponsePtr(v *ScopeResponseArgs) ScopeResponsePtrInput {
	return (*scopeResponsePtrType)(v)
}

func (*scopeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (i *scopeResponsePtrType) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return i.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (i *scopeResponsePtrType) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScopeResponsePtrOutput)
}

type ScopeResponseOutput struct{ *pulumi.OutputState }

func (ScopeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeResponse)(nil)).Elem()
}

func (o ScopeResponseOutput) ToScopeResponseOutput() ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponseOutputWithContext(ctx context.Context) ScopeResponseOutput {
	return o
}

func (o ScopeResponseOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o.ToScopeResponsePtrOutputWithContext(context.Background())
}

func (o ScopeResponseOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeResponse) *ScopeResponse {
		return &v
	}).(ScopeResponsePtrOutput)
}

func (o ScopeResponseOutput) ScopeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeResponse) *string { return v.ScopeType }).(pulumi.StringPtrOutput)
}

func (o ScopeResponseOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScopeResponse) []string { return v.Values }).(pulumi.StringArrayOutput)
}

type ScopeResponsePtrOutput struct{ *pulumi.OutputState }

func (ScopeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeResponse)(nil)).Elem()
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutput() ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) ToScopeResponsePtrOutputWithContext(ctx context.Context) ScopeResponsePtrOutput {
	return o
}

func (o ScopeResponsePtrOutput) Elem() ScopeResponseOutput {
	return o.ApplyT(func(v *ScopeResponse) ScopeResponse {
		if v != nil {
			return *v
		}
		var ret ScopeResponse
		return ret
	}).(ScopeResponseOutput)
}

func (o ScopeResponsePtrOutput) ScopeType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScopeType
	}).(pulumi.StringPtrOutput)
}

func (o ScopeResponsePtrOutput) Values() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ScopeResponse) []string {
		if v == nil {
			return nil
		}
		return v.Values
	}).(pulumi.StringArrayOutput)
}

type Suppression struct {
	Conditions        *Conditions       `pulumi:"conditions"`
	Description       *string           `pulumi:"description"`
	Scope             *Scope            `pulumi:"scope"`
	Status            *string           `pulumi:"status"`
	SuppressionConfig SuppressionConfig `pulumi:"suppressionConfig"`
	Type              string            `pulumi:"type"`
}





type SuppressionInput interface {
	pulumi.Input

	ToSuppressionOutput() SuppressionOutput
	ToSuppressionOutputWithContext(context.Context) SuppressionOutput
}

type SuppressionArgs struct {
	Conditions        ConditionsPtrInput     `pulumi:"conditions"`
	Description       pulumi.StringPtrInput  `pulumi:"description"`
	Scope             ScopePtrInput          `pulumi:"scope"`
	Status            pulumi.StringPtrInput  `pulumi:"status"`
	SuppressionConfig SuppressionConfigInput `pulumi:"suppressionConfig"`
	Type              pulumi.StringInput     `pulumi:"type"`
}

func (SuppressionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Suppression)(nil)).Elem()
}

func (i SuppressionArgs) ToSuppressionOutput() SuppressionOutput {
	return i.ToSuppressionOutputWithContext(context.Background())
}

func (i SuppressionArgs) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionOutput)
}

type SuppressionOutput struct{ *pulumi.OutputState }

func (SuppressionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Suppression)(nil)).Elem()
}

func (o SuppressionOutput) ToSuppressionOutput() SuppressionOutput {
	return o
}

func (o SuppressionOutput) ToSuppressionOutputWithContext(ctx context.Context) SuppressionOutput {
	return o
}

func (o SuppressionOutput) Conditions() ConditionsPtrOutput {
	return o.ApplyT(func(v Suppression) *Conditions { return v.Conditions }).(ConditionsPtrOutput)
}

func (o SuppressionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Suppression) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SuppressionOutput) Scope() ScopePtrOutput {
	return o.ApplyT(func(v Suppression) *Scope { return v.Scope }).(ScopePtrOutput)
}

func (o SuppressionOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Suppression) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SuppressionOutput) SuppressionConfig() SuppressionConfigOutput {
	return o.ApplyT(func(v Suppression) SuppressionConfig { return v.SuppressionConfig }).(SuppressionConfigOutput)
}

func (o SuppressionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v Suppression) string { return v.Type }).(pulumi.StringOutput)
}

type SuppressionConfig struct {
	RecurrenceType string               `pulumi:"recurrenceType"`
	Schedule       *SuppressionSchedule `pulumi:"schedule"`
}





type SuppressionConfigInput interface {
	pulumi.Input

	ToSuppressionConfigOutput() SuppressionConfigOutput
	ToSuppressionConfigOutputWithContext(context.Context) SuppressionConfigOutput
}

type SuppressionConfigArgs struct {
	RecurrenceType pulumi.StringInput          `pulumi:"recurrenceType"`
	Schedule       SuppressionSchedulePtrInput `pulumi:"schedule"`
}

func (SuppressionConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionConfig)(nil)).Elem()
}

func (i SuppressionConfigArgs) ToSuppressionConfigOutput() SuppressionConfigOutput {
	return i.ToSuppressionConfigOutputWithContext(context.Background())
}

func (i SuppressionConfigArgs) ToSuppressionConfigOutputWithContext(ctx context.Context) SuppressionConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionConfigOutput)
}

type SuppressionConfigOutput struct{ *pulumi.OutputState }

func (SuppressionConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionConfig)(nil)).Elem()
}

func (o SuppressionConfigOutput) ToSuppressionConfigOutput() SuppressionConfigOutput {
	return o
}

func (o SuppressionConfigOutput) ToSuppressionConfigOutputWithContext(ctx context.Context) SuppressionConfigOutput {
	return o
}

func (o SuppressionConfigOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionConfig) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o SuppressionConfigOutput) Schedule() SuppressionSchedulePtrOutput {
	return o.ApplyT(func(v SuppressionConfig) *SuppressionSchedule { return v.Schedule }).(SuppressionSchedulePtrOutput)
}

type SuppressionConfigResponse struct {
	RecurrenceType string                       `pulumi:"recurrenceType"`
	Schedule       *SuppressionScheduleResponse `pulumi:"schedule"`
}





type SuppressionConfigResponseInput interface {
	pulumi.Input

	ToSuppressionConfigResponseOutput() SuppressionConfigResponseOutput
	ToSuppressionConfigResponseOutputWithContext(context.Context) SuppressionConfigResponseOutput
}

type SuppressionConfigResponseArgs struct {
	RecurrenceType pulumi.StringInput                  `pulumi:"recurrenceType"`
	Schedule       SuppressionScheduleResponsePtrInput `pulumi:"schedule"`
}

func (SuppressionConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionConfigResponse)(nil)).Elem()
}

func (i SuppressionConfigResponseArgs) ToSuppressionConfigResponseOutput() SuppressionConfigResponseOutput {
	return i.ToSuppressionConfigResponseOutputWithContext(context.Background())
}

func (i SuppressionConfigResponseArgs) ToSuppressionConfigResponseOutputWithContext(ctx context.Context) SuppressionConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionConfigResponseOutput)
}

type SuppressionConfigResponseOutput struct{ *pulumi.OutputState }

func (SuppressionConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionConfigResponse)(nil)).Elem()
}

func (o SuppressionConfigResponseOutput) ToSuppressionConfigResponseOutput() SuppressionConfigResponseOutput {
	return o
}

func (o SuppressionConfigResponseOutput) ToSuppressionConfigResponseOutputWithContext(ctx context.Context) SuppressionConfigResponseOutput {
	return o
}

func (o SuppressionConfigResponseOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionConfigResponse) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o SuppressionConfigResponseOutput) Schedule() SuppressionScheduleResponsePtrOutput {
	return o.ApplyT(func(v SuppressionConfigResponse) *SuppressionScheduleResponse { return v.Schedule }).(SuppressionScheduleResponsePtrOutput)
}

type SuppressionResponse struct {
	Conditions        *ConditionsResponse       `pulumi:"conditions"`
	CreatedAt         string                    `pulumi:"createdAt"`
	CreatedBy         string                    `pulumi:"createdBy"`
	Description       *string                   `pulumi:"description"`
	LastModifiedAt    string                    `pulumi:"lastModifiedAt"`
	LastModifiedBy    string                    `pulumi:"lastModifiedBy"`
	Scope             *ScopeResponse            `pulumi:"scope"`
	Status            *string                   `pulumi:"status"`
	SuppressionConfig SuppressionConfigResponse `pulumi:"suppressionConfig"`
	Type              string                    `pulumi:"type"`
}





type SuppressionResponseInput interface {
	pulumi.Input

	ToSuppressionResponseOutput() SuppressionResponseOutput
	ToSuppressionResponseOutputWithContext(context.Context) SuppressionResponseOutput
}

type SuppressionResponseArgs struct {
	Conditions        ConditionsResponsePtrInput     `pulumi:"conditions"`
	CreatedAt         pulumi.StringInput             `pulumi:"createdAt"`
	CreatedBy         pulumi.StringInput             `pulumi:"createdBy"`
	Description       pulumi.StringPtrInput          `pulumi:"description"`
	LastModifiedAt    pulumi.StringInput             `pulumi:"lastModifiedAt"`
	LastModifiedBy    pulumi.StringInput             `pulumi:"lastModifiedBy"`
	Scope             ScopeResponsePtrInput          `pulumi:"scope"`
	Status            pulumi.StringPtrInput          `pulumi:"status"`
	SuppressionConfig SuppressionConfigResponseInput `pulumi:"suppressionConfig"`
	Type              pulumi.StringInput             `pulumi:"type"`
}

func (SuppressionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionResponse)(nil)).Elem()
}

func (i SuppressionResponseArgs) ToSuppressionResponseOutput() SuppressionResponseOutput {
	return i.ToSuppressionResponseOutputWithContext(context.Background())
}

func (i SuppressionResponseArgs) ToSuppressionResponseOutputWithContext(ctx context.Context) SuppressionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionResponseOutput)
}

type SuppressionResponseOutput struct{ *pulumi.OutputState }

func (SuppressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionResponse)(nil)).Elem()
}

func (o SuppressionResponseOutput) ToSuppressionResponseOutput() SuppressionResponseOutput {
	return o
}

func (o SuppressionResponseOutput) ToSuppressionResponseOutputWithContext(ctx context.Context) SuppressionResponseOutput {
	return o
}

func (o SuppressionResponseOutput) Conditions() ConditionsResponsePtrOutput {
	return o.ApplyT(func(v SuppressionResponse) *ConditionsResponse { return v.Conditions }).(ConditionsResponsePtrOutput)
}

func (o SuppressionResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o SuppressionResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o SuppressionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SuppressionResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o SuppressionResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o SuppressionResponseOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v SuppressionResponse) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o SuppressionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o SuppressionResponseOutput) SuppressionConfig() SuppressionConfigResponseOutput {
	return o.ApplyT(func(v SuppressionResponse) SuppressionConfigResponse { return v.SuppressionConfig }).(SuppressionConfigResponseOutput)
}

func (o SuppressionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type SuppressionSchedule struct {
	EndDate          *string `pulumi:"endDate"`
	EndTime          *string `pulumi:"endTime"`
	RecurrenceValues []int   `pulumi:"recurrenceValues"`
	StartDate        *string `pulumi:"startDate"`
	StartTime        *string `pulumi:"startTime"`
}





type SuppressionScheduleInput interface {
	pulumi.Input

	ToSuppressionScheduleOutput() SuppressionScheduleOutput
	ToSuppressionScheduleOutputWithContext(context.Context) SuppressionScheduleOutput
}

type SuppressionScheduleArgs struct {
	EndDate          pulumi.StringPtrInput `pulumi:"endDate"`
	EndTime          pulumi.StringPtrInput `pulumi:"endTime"`
	RecurrenceValues pulumi.IntArrayInput  `pulumi:"recurrenceValues"`
	StartDate        pulumi.StringPtrInput `pulumi:"startDate"`
	StartTime        pulumi.StringPtrInput `pulumi:"startTime"`
}

func (SuppressionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionSchedule)(nil)).Elem()
}

func (i SuppressionScheduleArgs) ToSuppressionScheduleOutput() SuppressionScheduleOutput {
	return i.ToSuppressionScheduleOutputWithContext(context.Background())
}

func (i SuppressionScheduleArgs) ToSuppressionScheduleOutputWithContext(ctx context.Context) SuppressionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionScheduleOutput)
}

func (i SuppressionScheduleArgs) ToSuppressionSchedulePtrOutput() SuppressionSchedulePtrOutput {
	return i.ToSuppressionSchedulePtrOutputWithContext(context.Background())
}

func (i SuppressionScheduleArgs) ToSuppressionSchedulePtrOutputWithContext(ctx context.Context) SuppressionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionScheduleOutput).ToSuppressionSchedulePtrOutputWithContext(ctx)
}









type SuppressionSchedulePtrInput interface {
	pulumi.Input

	ToSuppressionSchedulePtrOutput() SuppressionSchedulePtrOutput
	ToSuppressionSchedulePtrOutputWithContext(context.Context) SuppressionSchedulePtrOutput
}

type suppressionSchedulePtrType SuppressionScheduleArgs

func SuppressionSchedulePtr(v *SuppressionScheduleArgs) SuppressionSchedulePtrInput {
	return (*suppressionSchedulePtrType)(v)
}

func (*suppressionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionSchedule)(nil)).Elem()
}

func (i *suppressionSchedulePtrType) ToSuppressionSchedulePtrOutput() SuppressionSchedulePtrOutput {
	return i.ToSuppressionSchedulePtrOutputWithContext(context.Background())
}

func (i *suppressionSchedulePtrType) ToSuppressionSchedulePtrOutputWithContext(ctx context.Context) SuppressionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionSchedulePtrOutput)
}

type SuppressionScheduleOutput struct{ *pulumi.OutputState }

func (SuppressionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionSchedule)(nil)).Elem()
}

func (o SuppressionScheduleOutput) ToSuppressionScheduleOutput() SuppressionScheduleOutput {
	return o
}

func (o SuppressionScheduleOutput) ToSuppressionScheduleOutputWithContext(ctx context.Context) SuppressionScheduleOutput {
	return o
}

func (o SuppressionScheduleOutput) ToSuppressionSchedulePtrOutput() SuppressionSchedulePtrOutput {
	return o.ToSuppressionSchedulePtrOutputWithContext(context.Background())
}

func (o SuppressionScheduleOutput) ToSuppressionSchedulePtrOutputWithContext(ctx context.Context) SuppressionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionSchedule) *SuppressionSchedule {
		return &v
	}).(SuppressionSchedulePtrOutput)
}

func (o SuppressionScheduleOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionSchedule) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionSchedule) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleOutput) RecurrenceValues() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SuppressionSchedule) []int { return v.RecurrenceValues }).(pulumi.IntArrayOutput)
}

func (o SuppressionScheduleOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionSchedule) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionSchedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type SuppressionSchedulePtrOutput struct{ *pulumi.OutputState }

func (SuppressionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionSchedule)(nil)).Elem()
}

func (o SuppressionSchedulePtrOutput) ToSuppressionSchedulePtrOutput() SuppressionSchedulePtrOutput {
	return o
}

func (o SuppressionSchedulePtrOutput) ToSuppressionSchedulePtrOutputWithContext(ctx context.Context) SuppressionSchedulePtrOutput {
	return o
}

func (o SuppressionSchedulePtrOutput) Elem() SuppressionScheduleOutput {
	return o.ApplyT(func(v *SuppressionSchedule) SuppressionSchedule {
		if v != nil {
			return *v
		}
		var ret SuppressionSchedule
		return ret
	}).(SuppressionScheduleOutput)
}

func (o SuppressionSchedulePtrOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionSchedule) *string {
		if v == nil {
			return nil
		}
		return v.EndDate
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionSchedulePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionSchedule) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionSchedulePtrOutput) RecurrenceValues() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SuppressionSchedule) []int {
		if v == nil {
			return nil
		}
		return v.RecurrenceValues
	}).(pulumi.IntArrayOutput)
}

func (o SuppressionSchedulePtrOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartDate
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionSchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type SuppressionScheduleResponse struct {
	EndDate          *string `pulumi:"endDate"`
	EndTime          *string `pulumi:"endTime"`
	RecurrenceValues []int   `pulumi:"recurrenceValues"`
	StartDate        *string `pulumi:"startDate"`
	StartTime        *string `pulumi:"startTime"`
}





type SuppressionScheduleResponseInput interface {
	pulumi.Input

	ToSuppressionScheduleResponseOutput() SuppressionScheduleResponseOutput
	ToSuppressionScheduleResponseOutputWithContext(context.Context) SuppressionScheduleResponseOutput
}

type SuppressionScheduleResponseArgs struct {
	EndDate          pulumi.StringPtrInput `pulumi:"endDate"`
	EndTime          pulumi.StringPtrInput `pulumi:"endTime"`
	RecurrenceValues pulumi.IntArrayInput  `pulumi:"recurrenceValues"`
	StartDate        pulumi.StringPtrInput `pulumi:"startDate"`
	StartTime        pulumi.StringPtrInput `pulumi:"startTime"`
}

func (SuppressionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionScheduleResponse)(nil)).Elem()
}

func (i SuppressionScheduleResponseArgs) ToSuppressionScheduleResponseOutput() SuppressionScheduleResponseOutput {
	return i.ToSuppressionScheduleResponseOutputWithContext(context.Background())
}

func (i SuppressionScheduleResponseArgs) ToSuppressionScheduleResponseOutputWithContext(ctx context.Context) SuppressionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionScheduleResponseOutput)
}

func (i SuppressionScheduleResponseArgs) ToSuppressionScheduleResponsePtrOutput() SuppressionScheduleResponsePtrOutput {
	return i.ToSuppressionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i SuppressionScheduleResponseArgs) ToSuppressionScheduleResponsePtrOutputWithContext(ctx context.Context) SuppressionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionScheduleResponseOutput).ToSuppressionScheduleResponsePtrOutputWithContext(ctx)
}









type SuppressionScheduleResponsePtrInput interface {
	pulumi.Input

	ToSuppressionScheduleResponsePtrOutput() SuppressionScheduleResponsePtrOutput
	ToSuppressionScheduleResponsePtrOutputWithContext(context.Context) SuppressionScheduleResponsePtrOutput
}

type suppressionScheduleResponsePtrType SuppressionScheduleResponseArgs

func SuppressionScheduleResponsePtr(v *SuppressionScheduleResponseArgs) SuppressionScheduleResponsePtrInput {
	return (*suppressionScheduleResponsePtrType)(v)
}

func (*suppressionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionScheduleResponse)(nil)).Elem()
}

func (i *suppressionScheduleResponsePtrType) ToSuppressionScheduleResponsePtrOutput() SuppressionScheduleResponsePtrOutput {
	return i.ToSuppressionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *suppressionScheduleResponsePtrType) ToSuppressionScheduleResponsePtrOutputWithContext(ctx context.Context) SuppressionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionScheduleResponsePtrOutput)
}

type SuppressionScheduleResponseOutput struct{ *pulumi.OutputState }

func (SuppressionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SuppressionScheduleResponse)(nil)).Elem()
}

func (o SuppressionScheduleResponseOutput) ToSuppressionScheduleResponseOutput() SuppressionScheduleResponseOutput {
	return o
}

func (o SuppressionScheduleResponseOutput) ToSuppressionScheduleResponseOutputWithContext(ctx context.Context) SuppressionScheduleResponseOutput {
	return o
}

func (o SuppressionScheduleResponseOutput) ToSuppressionScheduleResponsePtrOutput() SuppressionScheduleResponsePtrOutput {
	return o.ToSuppressionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o SuppressionScheduleResponseOutput) ToSuppressionScheduleResponsePtrOutputWithContext(ctx context.Context) SuppressionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionScheduleResponse) *SuppressionScheduleResponse {
		return &v
	}).(SuppressionScheduleResponsePtrOutput)
}

func (o SuppressionScheduleResponseOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionScheduleResponse) *string { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleResponseOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionScheduleResponse) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleResponseOutput) RecurrenceValues() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SuppressionScheduleResponse) []int { return v.RecurrenceValues }).(pulumi.IntArrayOutput)
}

func (o SuppressionScheduleResponseOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionScheduleResponse) *string { return v.StartDate }).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SuppressionScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type SuppressionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (SuppressionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionScheduleResponse)(nil)).Elem()
}

func (o SuppressionScheduleResponsePtrOutput) ToSuppressionScheduleResponsePtrOutput() SuppressionScheduleResponsePtrOutput {
	return o
}

func (o SuppressionScheduleResponsePtrOutput) ToSuppressionScheduleResponsePtrOutputWithContext(ctx context.Context) SuppressionScheduleResponsePtrOutput {
	return o
}

func (o SuppressionScheduleResponsePtrOutput) Elem() SuppressionScheduleResponseOutput {
	return o.ApplyT(func(v *SuppressionScheduleResponse) SuppressionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret SuppressionScheduleResponse
		return ret
	}).(SuppressionScheduleResponseOutput)
}

func (o SuppressionScheduleResponsePtrOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndDate
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleResponsePtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleResponsePtrOutput) RecurrenceValues() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SuppressionScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.RecurrenceValues
	}).(pulumi.IntArrayOutput)
}

func (o SuppressionScheduleResponsePtrOutput) StartDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartDate
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActionGroupOutput{})
	pulumi.RegisterOutputType(ActionGroupResponseOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionPtrOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConditionsOutput{})
	pulumi.RegisterOutputType(ConditionsPtrOutput{})
	pulumi.RegisterOutputType(ConditionsResponseOutput{})
	pulumi.RegisterOutputType(ConditionsResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopePtrOutput{})
	pulumi.RegisterOutputType(ScopeResponseOutput{})
	pulumi.RegisterOutputType(ScopeResponsePtrOutput{})
	pulumi.RegisterOutputType(SuppressionOutput{})
	pulumi.RegisterOutputType(SuppressionConfigOutput{})
	pulumi.RegisterOutputType(SuppressionConfigResponseOutput{})
	pulumi.RegisterOutputType(SuppressionResponseOutput{})
	pulumi.RegisterOutputType(SuppressionScheduleOutput{})
	pulumi.RegisterOutputType(SuppressionSchedulePtrOutput{})
	pulumi.RegisterOutputType(SuppressionScheduleResponseOutput{})
	pulumi.RegisterOutputType(SuppressionScheduleResponsePtrOutput{})
}
