


package alertsmanagement

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

type ActionGroupsInformation struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}





type ActionGroupsInformationInput interface {
	pulumi.Input

	ToActionGroupsInformationOutput() ActionGroupsInformationOutput
	ToActionGroupsInformationOutputWithContext(context.Context) ActionGroupsInformationOutput
}

type ActionGroupsInformationArgs struct {
	CustomEmailSubject   pulumi.StringPtrInput   `pulumi:"customEmailSubject"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
}

func (ActionGroupsInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformation)(nil)).Elem()
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationOutput() ActionGroupsInformationOutput {
	return i.ToActionGroupsInformationOutputWithContext(context.Background())
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationOutputWithContext(ctx context.Context) ActionGroupsInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationOutput)
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return i.ToActionGroupsInformationPtrOutputWithContext(context.Background())
}

func (i ActionGroupsInformationArgs) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationOutput).ToActionGroupsInformationPtrOutputWithContext(ctx)
}









type ActionGroupsInformationPtrInput interface {
	pulumi.Input

	ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput
	ToActionGroupsInformationPtrOutputWithContext(context.Context) ActionGroupsInformationPtrOutput
}

type actionGroupsInformationPtrType ActionGroupsInformationArgs

func ActionGroupsInformationPtr(v *ActionGroupsInformationArgs) ActionGroupsInformationPtrInput {
	return (*actionGroupsInformationPtrType)(v)
}

func (*actionGroupsInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformation)(nil)).Elem()
}

func (i *actionGroupsInformationPtrType) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return i.ToActionGroupsInformationPtrOutputWithContext(context.Background())
}

func (i *actionGroupsInformationPtrType) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationPtrOutput)
}

type ActionGroupsInformationOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformation)(nil)).Elem()
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationOutput() ActionGroupsInformationOutput {
	return o
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationOutputWithContext(ctx context.Context) ActionGroupsInformationOutput {
	return o
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return o.ToActionGroupsInformationPtrOutputWithContext(context.Background())
}

func (o ActionGroupsInformationOutput) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionGroupsInformation) *ActionGroupsInformation {
		return &v
	}).(ActionGroupsInformationPtrOutput)
}

func (o ActionGroupsInformationOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformation) *string { return v.CustomEmailSubject }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformation) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionGroupsInformation) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

type ActionGroupsInformationPtrOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformation)(nil)).Elem()
}

func (o ActionGroupsInformationPtrOutput) ToActionGroupsInformationPtrOutput() ActionGroupsInformationPtrOutput {
	return o
}

func (o ActionGroupsInformationPtrOutput) ToActionGroupsInformationPtrOutputWithContext(ctx context.Context) ActionGroupsInformationPtrOutput {
	return o
}

func (o ActionGroupsInformationPtrOutput) Elem() ActionGroupsInformationOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) ActionGroupsInformation {
		if v != nil {
			return *v
		}
		var ret ActionGroupsInformation
		return ret
	}).(ActionGroupsInformationOutput)
}

func (o ActionGroupsInformationPtrOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) *string {
		if v == nil {
			return nil
		}
		return v.CustomEmailSubject
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationPtrOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) *string {
		if v == nil {
			return nil
		}
		return v.CustomWebhookPayload
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationPtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionGroupsInformation) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
}

type ActionGroupsInformationResponse struct {
	CustomEmailSubject   *string  `pulumi:"customEmailSubject"`
	CustomWebhookPayload *string  `pulumi:"customWebhookPayload"`
	GroupIds             []string `pulumi:"groupIds"`
}





type ActionGroupsInformationResponseInput interface {
	pulumi.Input

	ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput
	ToActionGroupsInformationResponseOutputWithContext(context.Context) ActionGroupsInformationResponseOutput
}

type ActionGroupsInformationResponseArgs struct {
	CustomEmailSubject   pulumi.StringPtrInput   `pulumi:"customEmailSubject"`
	CustomWebhookPayload pulumi.StringPtrInput   `pulumi:"customWebhookPayload"`
	GroupIds             pulumi.StringArrayInput `pulumi:"groupIds"`
}

func (ActionGroupsInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformationResponse)(nil)).Elem()
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput {
	return i.ToActionGroupsInformationResponseOutputWithContext(context.Background())
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponseOutputWithContext(ctx context.Context) ActionGroupsInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationResponseOutput)
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return i.ToActionGroupsInformationResponsePtrOutputWithContext(context.Background())
}

func (i ActionGroupsInformationResponseArgs) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationResponseOutput).ToActionGroupsInformationResponsePtrOutputWithContext(ctx)
}









type ActionGroupsInformationResponsePtrInput interface {
	pulumi.Input

	ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput
	ToActionGroupsInformationResponsePtrOutputWithContext(context.Context) ActionGroupsInformationResponsePtrOutput
}

type actionGroupsInformationResponsePtrType ActionGroupsInformationResponseArgs

func ActionGroupsInformationResponsePtr(v *ActionGroupsInformationResponseArgs) ActionGroupsInformationResponsePtrInput {
	return (*actionGroupsInformationResponsePtrType)(v)
}

func (*actionGroupsInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformationResponse)(nil)).Elem()
}

func (i *actionGroupsInformationResponsePtrType) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return i.ToActionGroupsInformationResponsePtrOutputWithContext(context.Background())
}

func (i *actionGroupsInformationResponsePtrType) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActionGroupsInformationResponsePtrOutput)
}

type ActionGroupsInformationResponseOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActionGroupsInformationResponse)(nil)).Elem()
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponseOutput() ActionGroupsInformationResponseOutput {
	return o
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponseOutputWithContext(ctx context.Context) ActionGroupsInformationResponseOutput {
	return o
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return o.ToActionGroupsInformationResponsePtrOutputWithContext(context.Background())
}

func (o ActionGroupsInformationResponseOutput) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActionGroupsInformationResponse) *ActionGroupsInformationResponse {
		return &v
	}).(ActionGroupsInformationResponsePtrOutput)
}

func (o ActionGroupsInformationResponseOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) *string { return v.CustomEmailSubject }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponseOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) *string { return v.CustomWebhookPayload }).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ActionGroupsInformationResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

type ActionGroupsInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ActionGroupsInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActionGroupsInformationResponse)(nil)).Elem()
}

func (o ActionGroupsInformationResponsePtrOutput) ToActionGroupsInformationResponsePtrOutput() ActionGroupsInformationResponsePtrOutput {
	return o
}

func (o ActionGroupsInformationResponsePtrOutput) ToActionGroupsInformationResponsePtrOutputWithContext(ctx context.Context) ActionGroupsInformationResponsePtrOutput {
	return o
}

func (o ActionGroupsInformationResponsePtrOutput) Elem() ActionGroupsInformationResponseOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) ActionGroupsInformationResponse {
		if v != nil {
			return *v
		}
		var ret ActionGroupsInformationResponse
		return ret
	}).(ActionGroupsInformationResponseOutput)
}

func (o ActionGroupsInformationResponsePtrOutput) CustomEmailSubject() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomEmailSubject
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponsePtrOutput) CustomWebhookPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomWebhookPayload
	}).(pulumi.StringPtrOutput)
}

func (o ActionGroupsInformationResponsePtrOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ActionGroupsInformationResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupIds
	}).(pulumi.StringArrayOutput)
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

type Detector struct {
	Description            *string                `pulumi:"description"`
	Id                     string                 `pulumi:"id"`
	ImagePaths             []string               `pulumi:"imagePaths"`
	Name                   *string                `pulumi:"name"`
	Parameters             map[string]interface{} `pulumi:"parameters"`
	SupportedResourceTypes []string               `pulumi:"supportedResourceTypes"`
}





type DetectorInput interface {
	pulumi.Input

	ToDetectorOutput() DetectorOutput
	ToDetectorOutputWithContext(context.Context) DetectorOutput
}

type DetectorArgs struct {
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	Id                     pulumi.StringInput      `pulumi:"id"`
	ImagePaths             pulumi.StringArrayInput `pulumi:"imagePaths"`
	Name                   pulumi.StringPtrInput   `pulumi:"name"`
	Parameters             pulumi.MapInput         `pulumi:"parameters"`
	SupportedResourceTypes pulumi.StringArrayInput `pulumi:"supportedResourceTypes"`
}

func (DetectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil)).Elem()
}

func (i DetectorArgs) ToDetectorOutput() DetectorOutput {
	return i.ToDetectorOutputWithContext(context.Background())
}

func (i DetectorArgs) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput)
}

func (i DetectorArgs) ToDetectorPtrOutput() DetectorPtrOutput {
	return i.ToDetectorPtrOutputWithContext(context.Background())
}

func (i DetectorArgs) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorOutput).ToDetectorPtrOutputWithContext(ctx)
}









type DetectorPtrInput interface {
	pulumi.Input

	ToDetectorPtrOutput() DetectorPtrOutput
	ToDetectorPtrOutputWithContext(context.Context) DetectorPtrOutput
}

type detectorPtrType DetectorArgs

func DetectorPtr(v *DetectorArgs) DetectorPtrInput {
	return (*detectorPtrType)(v)
}

func (*detectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (i *detectorPtrType) ToDetectorPtrOutput() DetectorPtrOutput {
	return i.ToDetectorPtrOutputWithContext(context.Background())
}

func (i *detectorPtrType) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorPtrOutput)
}

type DetectorOutput struct{ *pulumi.OutputState }

func (DetectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Detector)(nil)).Elem()
}

func (o DetectorOutput) ToDetectorOutput() DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorOutputWithContext(ctx context.Context) DetectorOutput {
	return o
}

func (o DetectorOutput) ToDetectorPtrOutput() DetectorPtrOutput {
	return o.ToDetectorPtrOutputWithContext(context.Background())
}

func (o DetectorOutput) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Detector) *Detector {
		return &v
	}).(DetectorPtrOutput)
}

func (o DetectorOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Detector) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v Detector) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Detector) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Detector) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v Detector) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Detector) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type DetectorPtrOutput struct{ *pulumi.OutputState }

func (DetectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Detector)(nil)).Elem()
}

func (o DetectorPtrOutput) ToDetectorPtrOutput() DetectorPtrOutput {
	return o
}

func (o DetectorPtrOutput) ToDetectorPtrOutputWithContext(ctx context.Context) DetectorPtrOutput {
	return o
}

func (o DetectorPtrOutput) Elem() DetectorOutput {
	return o.ApplyT(func(v *Detector) Detector {
		if v != nil {
			return *v
		}
		var ret Detector
		return ret
	}).(DetectorOutput)
}

func (o DetectorPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o DetectorPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DetectorPtrOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Detector) []string {
		if v == nil {
			return nil
		}
		return v.ImagePaths
	}).(pulumi.StringArrayOutput)
}

func (o DetectorPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Detector) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DetectorPtrOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v *Detector) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.MapOutput)
}

func (o DetectorPtrOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *Detector) []string {
		if v == nil {
			return nil
		}
		return v.SupportedResourceTypes
	}).(pulumi.StringArrayOutput)
}

type DetectorResponse struct {
	Description            *string                `pulumi:"description"`
	Id                     string                 `pulumi:"id"`
	ImagePaths             []string               `pulumi:"imagePaths"`
	Name                   *string                `pulumi:"name"`
	Parameters             map[string]interface{} `pulumi:"parameters"`
	SupportedResourceTypes []string               `pulumi:"supportedResourceTypes"`
}





type DetectorResponseInput interface {
	pulumi.Input

	ToDetectorResponseOutput() DetectorResponseOutput
	ToDetectorResponseOutputWithContext(context.Context) DetectorResponseOutput
}

type DetectorResponseArgs struct {
	Description            pulumi.StringPtrInput   `pulumi:"description"`
	Id                     pulumi.StringInput      `pulumi:"id"`
	ImagePaths             pulumi.StringArrayInput `pulumi:"imagePaths"`
	Name                   pulumi.StringPtrInput   `pulumi:"name"`
	Parameters             pulumi.MapInput         `pulumi:"parameters"`
	SupportedResourceTypes pulumi.StringArrayInput `pulumi:"supportedResourceTypes"`
}

func (DetectorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorResponse)(nil)).Elem()
}

func (i DetectorResponseArgs) ToDetectorResponseOutput() DetectorResponseOutput {
	return i.ToDetectorResponseOutputWithContext(context.Background())
}

func (i DetectorResponseArgs) ToDetectorResponseOutputWithContext(ctx context.Context) DetectorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorResponseOutput)
}

func (i DetectorResponseArgs) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return i.ToDetectorResponsePtrOutputWithContext(context.Background())
}

func (i DetectorResponseArgs) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorResponseOutput).ToDetectorResponsePtrOutputWithContext(ctx)
}









type DetectorResponsePtrInput interface {
	pulumi.Input

	ToDetectorResponsePtrOutput() DetectorResponsePtrOutput
	ToDetectorResponsePtrOutputWithContext(context.Context) DetectorResponsePtrOutput
}

type detectorResponsePtrType DetectorResponseArgs

func DetectorResponsePtr(v *DetectorResponseArgs) DetectorResponsePtrInput {
	return (*detectorResponsePtrType)(v)
}

func (*detectorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorResponse)(nil)).Elem()
}

func (i *detectorResponsePtrType) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return i.ToDetectorResponsePtrOutputWithContext(context.Background())
}

func (i *detectorResponsePtrType) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DetectorResponsePtrOutput)
}

type DetectorResponseOutput struct{ *pulumi.OutputState }

func (DetectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DetectorResponse)(nil)).Elem()
}

func (o DetectorResponseOutput) ToDetectorResponseOutput() DetectorResponseOutput {
	return o
}

func (o DetectorResponseOutput) ToDetectorResponseOutputWithContext(ctx context.Context) DetectorResponseOutput {
	return o
}

func (o DetectorResponseOutput) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return o.ToDetectorResponsePtrOutputWithContext(context.Background())
}

func (o DetectorResponseOutput) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DetectorResponse) *DetectorResponse {
		return &v
	}).(DetectorResponsePtrOutput)
}

func (o DetectorResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o DetectorResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v DetectorResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o DetectorResponseOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.ImagePaths }).(pulumi.StringArrayOutput)
}

func (o DetectorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DetectorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DetectorResponseOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v DetectorResponse) map[string]interface{} { return v.Parameters }).(pulumi.MapOutput)
}

func (o DetectorResponseOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DetectorResponse) []string { return v.SupportedResourceTypes }).(pulumi.StringArrayOutput)
}

type DetectorResponsePtrOutput struct{ *pulumi.OutputState }

func (DetectorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DetectorResponse)(nil)).Elem()
}

func (o DetectorResponsePtrOutput) ToDetectorResponsePtrOutput() DetectorResponsePtrOutput {
	return o
}

func (o DetectorResponsePtrOutput) ToDetectorResponsePtrOutputWithContext(ctx context.Context) DetectorResponsePtrOutput {
	return o
}

func (o DetectorResponsePtrOutput) Elem() DetectorResponseOutput {
	return o.ApplyT(func(v *DetectorResponse) DetectorResponse {
		if v != nil {
			return *v
		}
		var ret DetectorResponse
		return ret
	}).(DetectorResponseOutput)
}

func (o DetectorResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DetectorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o DetectorResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DetectorResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o DetectorResponsePtrOutput) ImagePaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DetectorResponse) []string {
		if v == nil {
			return nil
		}
		return v.ImagePaths
	}).(pulumi.StringArrayOutput)
}

func (o DetectorResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DetectorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o DetectorResponsePtrOutput) Parameters() pulumi.MapOutput {
	return o.ApplyT(func(v *DetectorResponse) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.MapOutput)
}

func (o DetectorResponsePtrOutput) SupportedResourceTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DetectorResponse) []string {
		if v == nil {
			return nil
		}
		return v.SupportedResourceTypes
	}).(pulumi.StringArrayOutput)
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

type HealthAlertAction struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type HealthAlertActionInput interface {
	pulumi.Input

	ToHealthAlertActionOutput() HealthAlertActionOutput
	ToHealthAlertActionOutputWithContext(context.Context) HealthAlertActionOutput
}

type HealthAlertActionArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (HealthAlertActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertAction)(nil)).Elem()
}

func (i HealthAlertActionArgs) ToHealthAlertActionOutput() HealthAlertActionOutput {
	return i.ToHealthAlertActionOutputWithContext(context.Background())
}

func (i HealthAlertActionArgs) ToHealthAlertActionOutputWithContext(ctx context.Context) HealthAlertActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionOutput)
}





type HealthAlertActionArrayInput interface {
	pulumi.Input

	ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput
	ToHealthAlertActionArrayOutputWithContext(context.Context) HealthAlertActionArrayOutput
}

type HealthAlertActionArray []HealthAlertActionInput

func (HealthAlertActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertAction)(nil)).Elem()
}

func (i HealthAlertActionArray) ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput {
	return i.ToHealthAlertActionArrayOutputWithContext(context.Background())
}

func (i HealthAlertActionArray) ToHealthAlertActionArrayOutputWithContext(ctx context.Context) HealthAlertActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionArrayOutput)
}

type HealthAlertActionOutput struct{ *pulumi.OutputState }

func (HealthAlertActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertAction)(nil)).Elem()
}

func (o HealthAlertActionOutput) ToHealthAlertActionOutput() HealthAlertActionOutput {
	return o
}

func (o HealthAlertActionOutput) ToHealthAlertActionOutputWithContext(ctx context.Context) HealthAlertActionOutput {
	return o
}

func (o HealthAlertActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthAlertAction) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o HealthAlertActionOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v HealthAlertAction) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type HealthAlertActionArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertAction)(nil)).Elem()
}

func (o HealthAlertActionArrayOutput) ToHealthAlertActionArrayOutput() HealthAlertActionArrayOutput {
	return o
}

func (o HealthAlertActionArrayOutput) ToHealthAlertActionArrayOutputWithContext(ctx context.Context) HealthAlertActionArrayOutput {
	return o
}

func (o HealthAlertActionArrayOutput) Index(i pulumi.IntInput) HealthAlertActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertAction {
		return vs[0].([]HealthAlertAction)[vs[1].(int)]
	}).(HealthAlertActionOutput)
}

type HealthAlertActionResponse struct {
	ActionGroupId     *string           `pulumi:"actionGroupId"`
	WebHookProperties map[string]string `pulumi:"webHookProperties"`
}





type HealthAlertActionResponseInput interface {
	pulumi.Input

	ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput
	ToHealthAlertActionResponseOutputWithContext(context.Context) HealthAlertActionResponseOutput
}

type HealthAlertActionResponseArgs struct {
	ActionGroupId     pulumi.StringPtrInput `pulumi:"actionGroupId"`
	WebHookProperties pulumi.StringMapInput `pulumi:"webHookProperties"`
}

func (HealthAlertActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertActionResponse)(nil)).Elem()
}

func (i HealthAlertActionResponseArgs) ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput {
	return i.ToHealthAlertActionResponseOutputWithContext(context.Background())
}

func (i HealthAlertActionResponseArgs) ToHealthAlertActionResponseOutputWithContext(ctx context.Context) HealthAlertActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionResponseOutput)
}





type HealthAlertActionResponseArrayInput interface {
	pulumi.Input

	ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput
	ToHealthAlertActionResponseArrayOutputWithContext(context.Context) HealthAlertActionResponseArrayOutput
}

type HealthAlertActionResponseArray []HealthAlertActionResponseInput

func (HealthAlertActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertActionResponse)(nil)).Elem()
}

func (i HealthAlertActionResponseArray) ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput {
	return i.ToHealthAlertActionResponseArrayOutputWithContext(context.Background())
}

func (i HealthAlertActionResponseArray) ToHealthAlertActionResponseArrayOutputWithContext(ctx context.Context) HealthAlertActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertActionResponseArrayOutput)
}

type HealthAlertActionResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertActionResponse)(nil)).Elem()
}

func (o HealthAlertActionResponseOutput) ToHealthAlertActionResponseOutput() HealthAlertActionResponseOutput {
	return o
}

func (o HealthAlertActionResponseOutput) ToHealthAlertActionResponseOutputWithContext(ctx context.Context) HealthAlertActionResponseOutput {
	return o
}

func (o HealthAlertActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HealthAlertActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o HealthAlertActionResponseOutput) WebHookProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v HealthAlertActionResponse) map[string]string { return v.WebHookProperties }).(pulumi.StringMapOutput)
}

type HealthAlertActionResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthAlertActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthAlertActionResponse)(nil)).Elem()
}

func (o HealthAlertActionResponseArrayOutput) ToHealthAlertActionResponseArrayOutput() HealthAlertActionResponseArrayOutput {
	return o
}

func (o HealthAlertActionResponseArrayOutput) ToHealthAlertActionResponseArrayOutputWithContext(ctx context.Context) HealthAlertActionResponseArrayOutput {
	return o
}

func (o HealthAlertActionResponseArrayOutput) Index(i pulumi.IntInput) HealthAlertActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthAlertActionResponse {
		return vs[0].([]HealthAlertActionResponse)[vs[1].(int)]
	}).(HealthAlertActionResponseOutput)
}

type HealthAlertCriteria struct {
	AllOf []VmGuestHealthAlertCriterion `pulumi:"allOf"`
}





type HealthAlertCriteriaInput interface {
	pulumi.Input

	ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput
	ToHealthAlertCriteriaOutputWithContext(context.Context) HealthAlertCriteriaOutput
}

type HealthAlertCriteriaArgs struct {
	AllOf VmGuestHealthAlertCriterionArrayInput `pulumi:"allOf"`
}

func (HealthAlertCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteria)(nil)).Elem()
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput {
	return i.ToHealthAlertCriteriaOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaOutputWithContext(ctx context.Context) HealthAlertCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaOutput)
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return i.ToHealthAlertCriteriaPtrOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaArgs) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaOutput).ToHealthAlertCriteriaPtrOutputWithContext(ctx)
}









type HealthAlertCriteriaPtrInput interface {
	pulumi.Input

	ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput
	ToHealthAlertCriteriaPtrOutputWithContext(context.Context) HealthAlertCriteriaPtrOutput
}

type healthAlertCriteriaPtrType HealthAlertCriteriaArgs

func HealthAlertCriteriaPtr(v *HealthAlertCriteriaArgs) HealthAlertCriteriaPtrInput {
	return (*healthAlertCriteriaPtrType)(v)
}

func (*healthAlertCriteriaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteria)(nil)).Elem()
}

func (i *healthAlertCriteriaPtrType) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return i.ToHealthAlertCriteriaPtrOutputWithContext(context.Background())
}

func (i *healthAlertCriteriaPtrType) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaPtrOutput)
}

type HealthAlertCriteriaOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteria)(nil)).Elem()
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaOutput() HealthAlertCriteriaOutput {
	return o
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaOutputWithContext(ctx context.Context) HealthAlertCriteriaOutput {
	return o
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return o.ToHealthAlertCriteriaPtrOutputWithContext(context.Background())
}

func (o HealthAlertCriteriaOutput) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthAlertCriteria) *HealthAlertCriteria {
		return &v
	}).(HealthAlertCriteriaPtrOutput)
}

func (o HealthAlertCriteriaOutput) AllOf() VmGuestHealthAlertCriterionArrayOutput {
	return o.ApplyT(func(v HealthAlertCriteria) []VmGuestHealthAlertCriterion { return v.AllOf }).(VmGuestHealthAlertCriterionArrayOutput)
}

type HealthAlertCriteriaPtrOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteria)(nil)).Elem()
}

func (o HealthAlertCriteriaPtrOutput) ToHealthAlertCriteriaPtrOutput() HealthAlertCriteriaPtrOutput {
	return o
}

func (o HealthAlertCriteriaPtrOutput) ToHealthAlertCriteriaPtrOutputWithContext(ctx context.Context) HealthAlertCriteriaPtrOutput {
	return o
}

func (o HealthAlertCriteriaPtrOutput) Elem() HealthAlertCriteriaOutput {
	return o.ApplyT(func(v *HealthAlertCriteria) HealthAlertCriteria {
		if v != nil {
			return *v
		}
		var ret HealthAlertCriteria
		return ret
	}).(HealthAlertCriteriaOutput)
}

func (o HealthAlertCriteriaPtrOutput) AllOf() VmGuestHealthAlertCriterionArrayOutput {
	return o.ApplyT(func(v *HealthAlertCriteria) []VmGuestHealthAlertCriterion {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(VmGuestHealthAlertCriterionArrayOutput)
}

type HealthAlertCriteriaResponse struct {
	AllOf []VmGuestHealthAlertCriterionResponse `pulumi:"allOf"`
}





type HealthAlertCriteriaResponseInput interface {
	pulumi.Input

	ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput
	ToHealthAlertCriteriaResponseOutputWithContext(context.Context) HealthAlertCriteriaResponseOutput
}

type HealthAlertCriteriaResponseArgs struct {
	AllOf VmGuestHealthAlertCriterionResponseArrayInput `pulumi:"allOf"`
}

func (HealthAlertCriteriaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteriaResponse)(nil)).Elem()
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput {
	return i.ToHealthAlertCriteriaResponseOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponseOutputWithContext(ctx context.Context) HealthAlertCriteriaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaResponseOutput)
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return i.ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Background())
}

func (i HealthAlertCriteriaResponseArgs) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaResponseOutput).ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx)
}









type HealthAlertCriteriaResponsePtrInput interface {
	pulumi.Input

	ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput
	ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Context) HealthAlertCriteriaResponsePtrOutput
}

type healthAlertCriteriaResponsePtrType HealthAlertCriteriaResponseArgs

func HealthAlertCriteriaResponsePtr(v *HealthAlertCriteriaResponseArgs) HealthAlertCriteriaResponsePtrInput {
	return (*healthAlertCriteriaResponsePtrType)(v)
}

func (*healthAlertCriteriaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteriaResponse)(nil)).Elem()
}

func (i *healthAlertCriteriaResponsePtrType) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return i.ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Background())
}

func (i *healthAlertCriteriaResponsePtrType) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthAlertCriteriaResponsePtrOutput)
}

type HealthAlertCriteriaResponseOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthAlertCriteriaResponse)(nil)).Elem()
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponseOutput() HealthAlertCriteriaResponseOutput {
	return o
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponseOutputWithContext(ctx context.Context) HealthAlertCriteriaResponseOutput {
	return o
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return o.ToHealthAlertCriteriaResponsePtrOutputWithContext(context.Background())
}

func (o HealthAlertCriteriaResponseOutput) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HealthAlertCriteriaResponse) *HealthAlertCriteriaResponse {
		return &v
	}).(HealthAlertCriteriaResponsePtrOutput)
}

func (o HealthAlertCriteriaResponseOutput) AllOf() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o.ApplyT(func(v HealthAlertCriteriaResponse) []VmGuestHealthAlertCriterionResponse { return v.AllOf }).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type HealthAlertCriteriaResponsePtrOutput struct{ *pulumi.OutputState }

func (HealthAlertCriteriaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HealthAlertCriteriaResponse)(nil)).Elem()
}

func (o HealthAlertCriteriaResponsePtrOutput) ToHealthAlertCriteriaResponsePtrOutput() HealthAlertCriteriaResponsePtrOutput {
	return o
}

func (o HealthAlertCriteriaResponsePtrOutput) ToHealthAlertCriteriaResponsePtrOutputWithContext(ctx context.Context) HealthAlertCriteriaResponsePtrOutput {
	return o
}

func (o HealthAlertCriteriaResponsePtrOutput) Elem() HealthAlertCriteriaResponseOutput {
	return o.ApplyT(func(v *HealthAlertCriteriaResponse) HealthAlertCriteriaResponse {
		if v != nil {
			return *v
		}
		var ret HealthAlertCriteriaResponse
		return ret
	}).(HealthAlertCriteriaResponseOutput)
}

func (o HealthAlertCriteriaResponsePtrOutput) AllOf() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o.ApplyT(func(v *HealthAlertCriteriaResponse) []VmGuestHealthAlertCriterionResponse {
		if v == nil {
			return nil
		}
		return v.AllOf
	}).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type HealthState struct {
	HealthStateName string  `pulumi:"healthStateName"`
	Severity        float64 `pulumi:"severity"`
}





type HealthStateInput interface {
	pulumi.Input

	ToHealthStateOutput() HealthStateOutput
	ToHealthStateOutputWithContext(context.Context) HealthStateOutput
}

type HealthStateArgs struct {
	HealthStateName pulumi.StringInput  `pulumi:"healthStateName"`
	Severity        pulumi.Float64Input `pulumi:"severity"`
}

func (HealthStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthState)(nil)).Elem()
}

func (i HealthStateArgs) ToHealthStateOutput() HealthStateOutput {
	return i.ToHealthStateOutputWithContext(context.Background())
}

func (i HealthStateArgs) ToHealthStateOutputWithContext(ctx context.Context) HealthStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateOutput)
}





type HealthStateArrayInput interface {
	pulumi.Input

	ToHealthStateArrayOutput() HealthStateArrayOutput
	ToHealthStateArrayOutputWithContext(context.Context) HealthStateArrayOutput
}

type HealthStateArray []HealthStateInput

func (HealthStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthState)(nil)).Elem()
}

func (i HealthStateArray) ToHealthStateArrayOutput() HealthStateArrayOutput {
	return i.ToHealthStateArrayOutputWithContext(context.Background())
}

func (i HealthStateArray) ToHealthStateArrayOutputWithContext(ctx context.Context) HealthStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateArrayOutput)
}

type HealthStateOutput struct{ *pulumi.OutputState }

func (HealthStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthState)(nil)).Elem()
}

func (o HealthStateOutput) ToHealthStateOutput() HealthStateOutput {
	return o
}

func (o HealthStateOutput) ToHealthStateOutputWithContext(ctx context.Context) HealthStateOutput {
	return o
}

func (o HealthStateOutput) HealthStateName() pulumi.StringOutput {
	return o.ApplyT(func(v HealthState) string { return v.HealthStateName }).(pulumi.StringOutput)
}

func (o HealthStateOutput) Severity() pulumi.Float64Output {
	return o.ApplyT(func(v HealthState) float64 { return v.Severity }).(pulumi.Float64Output)
}

type HealthStateArrayOutput struct{ *pulumi.OutputState }

func (HealthStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthState)(nil)).Elem()
}

func (o HealthStateArrayOutput) ToHealthStateArrayOutput() HealthStateArrayOutput {
	return o
}

func (o HealthStateArrayOutput) ToHealthStateArrayOutputWithContext(ctx context.Context) HealthStateArrayOutput {
	return o
}

func (o HealthStateArrayOutput) Index(i pulumi.IntInput) HealthStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthState {
		return vs[0].([]HealthState)[vs[1].(int)]
	}).(HealthStateOutput)
}

type HealthStateResponse struct {
	HealthStateName string  `pulumi:"healthStateName"`
	Severity        float64 `pulumi:"severity"`
}





type HealthStateResponseInput interface {
	pulumi.Input

	ToHealthStateResponseOutput() HealthStateResponseOutput
	ToHealthStateResponseOutputWithContext(context.Context) HealthStateResponseOutput
}

type HealthStateResponseArgs struct {
	HealthStateName pulumi.StringInput  `pulumi:"healthStateName"`
	Severity        pulumi.Float64Input `pulumi:"severity"`
}

func (HealthStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStateResponse)(nil)).Elem()
}

func (i HealthStateResponseArgs) ToHealthStateResponseOutput() HealthStateResponseOutput {
	return i.ToHealthStateResponseOutputWithContext(context.Background())
}

func (i HealthStateResponseArgs) ToHealthStateResponseOutputWithContext(ctx context.Context) HealthStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateResponseOutput)
}





type HealthStateResponseArrayInput interface {
	pulumi.Input

	ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput
	ToHealthStateResponseArrayOutputWithContext(context.Context) HealthStateResponseArrayOutput
}

type HealthStateResponseArray []HealthStateResponseInput

func (HealthStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthStateResponse)(nil)).Elem()
}

func (i HealthStateResponseArray) ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput {
	return i.ToHealthStateResponseArrayOutputWithContext(context.Background())
}

func (i HealthStateResponseArray) ToHealthStateResponseArrayOutputWithContext(ctx context.Context) HealthStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HealthStateResponseArrayOutput)
}

type HealthStateResponseOutput struct{ *pulumi.OutputState }

func (HealthStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HealthStateResponse)(nil)).Elem()
}

func (o HealthStateResponseOutput) ToHealthStateResponseOutput() HealthStateResponseOutput {
	return o
}

func (o HealthStateResponseOutput) ToHealthStateResponseOutputWithContext(ctx context.Context) HealthStateResponseOutput {
	return o
}

func (o HealthStateResponseOutput) HealthStateName() pulumi.StringOutput {
	return o.ApplyT(func(v HealthStateResponse) string { return v.HealthStateName }).(pulumi.StringOutput)
}

func (o HealthStateResponseOutput) Severity() pulumi.Float64Output {
	return o.ApplyT(func(v HealthStateResponse) float64 { return v.Severity }).(pulumi.Float64Output)
}

type HealthStateResponseArrayOutput struct{ *pulumi.OutputState }

func (HealthStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HealthStateResponse)(nil)).Elem()
}

func (o HealthStateResponseArrayOutput) ToHealthStateResponseArrayOutput() HealthStateResponseArrayOutput {
	return o
}

func (o HealthStateResponseArrayOutput) ToHealthStateResponseArrayOutputWithContext(ctx context.Context) HealthStateResponseArrayOutput {
	return o
}

func (o HealthStateResponseArrayOutput) Index(i pulumi.IntInput) HealthStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HealthStateResponse {
		return vs[0].([]HealthStateResponse)[vs[1].(int)]
	}).(HealthStateResponseOutput)
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

type ThrottlingInformation struct {
	Duration *string `pulumi:"duration"`
}





type ThrottlingInformationInput interface {
	pulumi.Input

	ToThrottlingInformationOutput() ThrottlingInformationOutput
	ToThrottlingInformationOutputWithContext(context.Context) ThrottlingInformationOutput
}

type ThrottlingInformationArgs struct {
	Duration pulumi.StringPtrInput `pulumi:"duration"`
}

func (ThrottlingInformationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformation)(nil)).Elem()
}

func (i ThrottlingInformationArgs) ToThrottlingInformationOutput() ThrottlingInformationOutput {
	return i.ToThrottlingInformationOutputWithContext(context.Background())
}

func (i ThrottlingInformationArgs) ToThrottlingInformationOutputWithContext(ctx context.Context) ThrottlingInformationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationOutput)
}

func (i ThrottlingInformationArgs) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return i.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (i ThrottlingInformationArgs) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationOutput).ToThrottlingInformationPtrOutputWithContext(ctx)
}









type ThrottlingInformationPtrInput interface {
	pulumi.Input

	ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput
	ToThrottlingInformationPtrOutputWithContext(context.Context) ThrottlingInformationPtrOutput
}

type throttlingInformationPtrType ThrottlingInformationArgs

func ThrottlingInformationPtr(v *ThrottlingInformationArgs) ThrottlingInformationPtrInput {
	return (*throttlingInformationPtrType)(v)
}

func (*throttlingInformationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformation)(nil)).Elem()
}

func (i *throttlingInformationPtrType) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return i.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (i *throttlingInformationPtrType) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationPtrOutput)
}

type ThrottlingInformationOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformation)(nil)).Elem()
}

func (o ThrottlingInformationOutput) ToThrottlingInformationOutput() ThrottlingInformationOutput {
	return o
}

func (o ThrottlingInformationOutput) ToThrottlingInformationOutputWithContext(ctx context.Context) ThrottlingInformationOutput {
	return o
}

func (o ThrottlingInformationOutput) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return o.ToThrottlingInformationPtrOutputWithContext(context.Background())
}

func (o ThrottlingInformationOutput) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThrottlingInformation) *ThrottlingInformation {
		return &v
	}).(ThrottlingInformationPtrOutput)
}

func (o ThrottlingInformationOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingInformation) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

type ThrottlingInformationPtrOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformation)(nil)).Elem()
}

func (o ThrottlingInformationPtrOutput) ToThrottlingInformationPtrOutput() ThrottlingInformationPtrOutput {
	return o
}

func (o ThrottlingInformationPtrOutput) ToThrottlingInformationPtrOutputWithContext(ctx context.Context) ThrottlingInformationPtrOutput {
	return o
}

func (o ThrottlingInformationPtrOutput) Elem() ThrottlingInformationOutput {
	return o.ApplyT(func(v *ThrottlingInformation) ThrottlingInformation {
		if v != nil {
			return *v
		}
		var ret ThrottlingInformation
		return ret
	}).(ThrottlingInformationOutput)
}

func (o ThrottlingInformationPtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThrottlingInformation) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

type ThrottlingInformationResponse struct {
	Duration *string `pulumi:"duration"`
}





type ThrottlingInformationResponseInput interface {
	pulumi.Input

	ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput
	ToThrottlingInformationResponseOutputWithContext(context.Context) ThrottlingInformationResponseOutput
}

type ThrottlingInformationResponseArgs struct {
	Duration pulumi.StringPtrInput `pulumi:"duration"`
}

func (ThrottlingInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformationResponse)(nil)).Elem()
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput {
	return i.ToThrottlingInformationResponseOutputWithContext(context.Background())
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponseOutputWithContext(ctx context.Context) ThrottlingInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationResponseOutput)
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return i.ToThrottlingInformationResponsePtrOutputWithContext(context.Background())
}

func (i ThrottlingInformationResponseArgs) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationResponseOutput).ToThrottlingInformationResponsePtrOutputWithContext(ctx)
}









type ThrottlingInformationResponsePtrInput interface {
	pulumi.Input

	ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput
	ToThrottlingInformationResponsePtrOutputWithContext(context.Context) ThrottlingInformationResponsePtrOutput
}

type throttlingInformationResponsePtrType ThrottlingInformationResponseArgs

func ThrottlingInformationResponsePtr(v *ThrottlingInformationResponseArgs) ThrottlingInformationResponsePtrInput {
	return (*throttlingInformationResponsePtrType)(v)
}

func (*throttlingInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformationResponse)(nil)).Elem()
}

func (i *throttlingInformationResponsePtrType) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return i.ToThrottlingInformationResponsePtrOutputWithContext(context.Background())
}

func (i *throttlingInformationResponsePtrType) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThrottlingInformationResponsePtrOutput)
}

type ThrottlingInformationResponseOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThrottlingInformationResponse)(nil)).Elem()
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponseOutput() ThrottlingInformationResponseOutput {
	return o
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponseOutputWithContext(ctx context.Context) ThrottlingInformationResponseOutput {
	return o
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return o.ToThrottlingInformationResponsePtrOutputWithContext(context.Background())
}

func (o ThrottlingInformationResponseOutput) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThrottlingInformationResponse) *ThrottlingInformationResponse {
		return &v
	}).(ThrottlingInformationResponsePtrOutput)
}

func (o ThrottlingInformationResponseOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThrottlingInformationResponse) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

type ThrottlingInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (ThrottlingInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThrottlingInformationResponse)(nil)).Elem()
}

func (o ThrottlingInformationResponsePtrOutput) ToThrottlingInformationResponsePtrOutput() ThrottlingInformationResponsePtrOutput {
	return o
}

func (o ThrottlingInformationResponsePtrOutput) ToThrottlingInformationResponsePtrOutputWithContext(ctx context.Context) ThrottlingInformationResponsePtrOutput {
	return o
}

func (o ThrottlingInformationResponsePtrOutput) Elem() ThrottlingInformationResponseOutput {
	return o.ApplyT(func(v *ThrottlingInformationResponse) ThrottlingInformationResponse {
		if v != nil {
			return *v
		}
		var ret ThrottlingInformationResponse
		return ret
	}).(ThrottlingInformationResponseOutput)
}

func (o ThrottlingInformationResponsePtrOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThrottlingInformationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Duration
	}).(pulumi.StringPtrOutput)
}

type VmGuestHealthAlertCriterion struct {
	HealthStates []HealthState `pulumi:"healthStates"`
	MonitorNames []string      `pulumi:"monitorNames"`
	MonitorTypes []string      `pulumi:"monitorTypes"`
	Namespace    string        `pulumi:"namespace"`
}





type VmGuestHealthAlertCriterionInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput
	ToVmGuestHealthAlertCriterionOutputWithContext(context.Context) VmGuestHealthAlertCriterionOutput
}

type VmGuestHealthAlertCriterionArgs struct {
	HealthStates HealthStateArrayInput   `pulumi:"healthStates"`
	MonitorNames pulumi.StringArrayInput `pulumi:"monitorNames"`
	MonitorTypes pulumi.StringArrayInput `pulumi:"monitorTypes"`
	Namespace    pulumi.StringInput      `pulumi:"namespace"`
}

func (VmGuestHealthAlertCriterionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionArgs) ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput {
	return i.ToVmGuestHealthAlertCriterionOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionArgs) ToVmGuestHealthAlertCriterionOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionOutput)
}





type VmGuestHealthAlertCriterionArrayInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput
	ToVmGuestHealthAlertCriterionArrayOutputWithContext(context.Context) VmGuestHealthAlertCriterionArrayOutput
}

type VmGuestHealthAlertCriterionArray []VmGuestHealthAlertCriterionInput

func (VmGuestHealthAlertCriterionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionArray) ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput {
	return i.ToVmGuestHealthAlertCriterionArrayOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionArray) ToVmGuestHealthAlertCriterionArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionArrayOutput)
}

type VmGuestHealthAlertCriterionOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionOutput) ToVmGuestHealthAlertCriterionOutput() VmGuestHealthAlertCriterionOutput {
	return o
}

func (o VmGuestHealthAlertCriterionOutput) ToVmGuestHealthAlertCriterionOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionOutput {
	return o
}

func (o VmGuestHealthAlertCriterionOutput) HealthStates() HealthStateArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []HealthState { return v.HealthStates }).(HealthStateArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) MonitorNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []string { return v.MonitorNames }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterion) string { return v.Namespace }).(pulumi.StringOutput)
}

type VmGuestHealthAlertCriterionArrayOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterion)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionArrayOutput) ToVmGuestHealthAlertCriterionArrayOutput() VmGuestHealthAlertCriterionArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionArrayOutput) ToVmGuestHealthAlertCriterionArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionArrayOutput) Index(i pulumi.IntInput) VmGuestHealthAlertCriterionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmGuestHealthAlertCriterion {
		return vs[0].([]VmGuestHealthAlertCriterion)[vs[1].(int)]
	}).(VmGuestHealthAlertCriterionOutput)
}

type VmGuestHealthAlertCriterionResponse struct {
	HealthStates []HealthStateResponse `pulumi:"healthStates"`
	MonitorNames []string              `pulumi:"monitorNames"`
	MonitorTypes []string              `pulumi:"monitorTypes"`
	Namespace    string                `pulumi:"namespace"`
}





type VmGuestHealthAlertCriterionResponseInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput
	ToVmGuestHealthAlertCriterionResponseOutputWithContext(context.Context) VmGuestHealthAlertCriterionResponseOutput
}

type VmGuestHealthAlertCriterionResponseArgs struct {
	HealthStates HealthStateResponseArrayInput `pulumi:"healthStates"`
	MonitorNames pulumi.StringArrayInput       `pulumi:"monitorNames"`
	MonitorTypes pulumi.StringArrayInput       `pulumi:"monitorTypes"`
	Namespace    pulumi.StringInput            `pulumi:"namespace"`
}

func (VmGuestHealthAlertCriterionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionResponseArgs) ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput {
	return i.ToVmGuestHealthAlertCriterionResponseOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionResponseArgs) ToVmGuestHealthAlertCriterionResponseOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionResponseOutput)
}





type VmGuestHealthAlertCriterionResponseArrayInput interface {
	pulumi.Input

	ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput
	ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(context.Context) VmGuestHealthAlertCriterionResponseArrayOutput
}

type VmGuestHealthAlertCriterionResponseArray []VmGuestHealthAlertCriterionResponseInput

func (VmGuestHealthAlertCriterionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (i VmGuestHealthAlertCriterionResponseArray) ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput {
	return i.ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(context.Background())
}

func (i VmGuestHealthAlertCriterionResponseArray) ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VmGuestHealthAlertCriterionResponseArrayOutput)
}

type VmGuestHealthAlertCriterionResponseOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionResponseOutput) ToVmGuestHealthAlertCriterionResponseOutput() VmGuestHealthAlertCriterionResponseOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseOutput) ToVmGuestHealthAlertCriterionResponseOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseOutput) HealthStates() HealthStateResponseArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []HealthStateResponse { return v.HealthStates }).(HealthStateResponseArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) MonitorNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []string { return v.MonitorNames }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) MonitorTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) []string { return v.MonitorTypes }).(pulumi.StringArrayOutput)
}

func (o VmGuestHealthAlertCriterionResponseOutput) Namespace() pulumi.StringOutput {
	return o.ApplyT(func(v VmGuestHealthAlertCriterionResponse) string { return v.Namespace }).(pulumi.StringOutput)
}

type VmGuestHealthAlertCriterionResponseArrayOutput struct{ *pulumi.OutputState }

func (VmGuestHealthAlertCriterionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VmGuestHealthAlertCriterionResponse)(nil)).Elem()
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) ToVmGuestHealthAlertCriterionResponseArrayOutput() VmGuestHealthAlertCriterionResponseArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) ToVmGuestHealthAlertCriterionResponseArrayOutputWithContext(ctx context.Context) VmGuestHealthAlertCriterionResponseArrayOutput {
	return o
}

func (o VmGuestHealthAlertCriterionResponseArrayOutput) Index(i pulumi.IntInput) VmGuestHealthAlertCriterionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VmGuestHealthAlertCriterionResponse {
		return vs[0].([]VmGuestHealthAlertCriterionResponse)[vs[1].(int)]
	}).(VmGuestHealthAlertCriterionResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupInput)(nil)).Elem(), ActionGroupArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupResponseInput)(nil)).Elem(), ActionGroupResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupsInformationInput)(nil)).Elem(), ActionGroupsInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupsInformationPtrInput)(nil)).Elem(), ActionGroupsInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupsInformationResponseInput)(nil)).Elem(), ActionGroupsInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ActionGroupsInformationResponsePtrInput)(nil)).Elem(), ActionGroupsInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionInput)(nil)).Elem(), ConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionPtrInput)(nil)).Elem(), ConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionResponseInput)(nil)).Elem(), ConditionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionResponsePtrInput)(nil)).Elem(), ConditionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionsInput)(nil)).Elem(), ConditionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionsPtrInput)(nil)).Elem(), ConditionsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionsResponseInput)(nil)).Elem(), ConditionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConditionsResponsePtrInput)(nil)).Elem(), ConditionsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorInput)(nil)).Elem(), DetectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorPtrInput)(nil)).Elem(), DetectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorResponseInput)(nil)).Elem(), DetectorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DetectorResponsePtrInput)(nil)).Elem(), DetectorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiagnosticsInput)(nil)).Elem(), DiagnosticsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DiagnosticsResponseInput)(nil)).Elem(), DiagnosticsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertActionInput)(nil)).Elem(), HealthAlertActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertActionArrayInput)(nil)).Elem(), HealthAlertActionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertActionResponseInput)(nil)).Elem(), HealthAlertActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertActionResponseArrayInput)(nil)).Elem(), HealthAlertActionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertCriteriaInput)(nil)).Elem(), HealthAlertCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertCriteriaPtrInput)(nil)).Elem(), HealthAlertCriteriaArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertCriteriaResponseInput)(nil)).Elem(), HealthAlertCriteriaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthAlertCriteriaResponsePtrInput)(nil)).Elem(), HealthAlertCriteriaResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthStateInput)(nil)).Elem(), HealthStateArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthStateArrayInput)(nil)).Elem(), HealthStateArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthStateResponseInput)(nil)).Elem(), HealthStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*HealthStateResponseArrayInput)(nil)).Elem(), HealthStateResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopeInput)(nil)).Elem(), ScopeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopePtrInput)(nil)).Elem(), ScopeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopeResponseInput)(nil)).Elem(), ScopeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScopeResponsePtrInput)(nil)).Elem(), ScopeResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionInput)(nil)).Elem(), SuppressionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionConfigInput)(nil)).Elem(), SuppressionConfigArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionConfigResponseInput)(nil)).Elem(), SuppressionConfigResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionResponseInput)(nil)).Elem(), SuppressionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionScheduleInput)(nil)).Elem(), SuppressionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionSchedulePtrInput)(nil)).Elem(), SuppressionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionScheduleResponseInput)(nil)).Elem(), SuppressionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SuppressionScheduleResponsePtrInput)(nil)).Elem(), SuppressionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThrottlingInformationInput)(nil)).Elem(), ThrottlingInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThrottlingInformationPtrInput)(nil)).Elem(), ThrottlingInformationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThrottlingInformationResponseInput)(nil)).Elem(), ThrottlingInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThrottlingInformationResponsePtrInput)(nil)).Elem(), ThrottlingInformationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmGuestHealthAlertCriterionInput)(nil)).Elem(), VmGuestHealthAlertCriterionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmGuestHealthAlertCriterionArrayInput)(nil)).Elem(), VmGuestHealthAlertCriterionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmGuestHealthAlertCriterionResponseInput)(nil)).Elem(), VmGuestHealthAlertCriterionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VmGuestHealthAlertCriterionResponseArrayInput)(nil)).Elem(), VmGuestHealthAlertCriterionResponseArray{})
	pulumi.RegisterOutputType(ActionGroupOutput{})
	pulumi.RegisterOutputType(ActionGroupResponseOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationPtrOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationResponseOutput{})
	pulumi.RegisterOutputType(ActionGroupsInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionPtrOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConditionsOutput{})
	pulumi.RegisterOutputType(ConditionsPtrOutput{})
	pulumi.RegisterOutputType(ConditionsResponseOutput{})
	pulumi.RegisterOutputType(ConditionsResponsePtrOutput{})
	pulumi.RegisterOutputType(DetectorOutput{})
	pulumi.RegisterOutputType(DetectorPtrOutput{})
	pulumi.RegisterOutputType(DetectorResponseOutput{})
	pulumi.RegisterOutputType(DetectorResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsOutput{})
	pulumi.RegisterOutputType(DiagnosticsResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertActionOutput{})
	pulumi.RegisterOutputType(HealthAlertActionArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertActionResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertActionResponseArrayOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaPtrOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaResponseOutput{})
	pulumi.RegisterOutputType(HealthAlertCriteriaResponsePtrOutput{})
	pulumi.RegisterOutputType(HealthStateOutput{})
	pulumi.RegisterOutputType(HealthStateArrayOutput{})
	pulumi.RegisterOutputType(HealthStateResponseOutput{})
	pulumi.RegisterOutputType(HealthStateResponseArrayOutput{})
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
	pulumi.RegisterOutputType(ThrottlingInformationOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationPtrOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponseOutput{})
	pulumi.RegisterOutputType(ThrottlingInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionArrayOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionResponseOutput{})
	pulumi.RegisterOutputType(VmGuestHealthAlertCriterionResponseArrayOutput{})
}
