


package v20181102privatepreview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActionRuleProperties struct {
	Conditions        *Conditions        `pulumi:"conditions"`
	Description       *string            `pulumi:"description"`
	Scope             *Scope             `pulumi:"scope"`
	Status            *string            `pulumi:"status"`
	SuppressionConfig *SuppressionConfig `pulumi:"suppressionConfig"`
}





type ActionRulePropertiesInput interface {
	pulumi.Input

	ToActionRulePropertiesOutput() ActionRulePropertiesOutput
	ToActionRulePropertiesOutputWithContext(context.Context) ActionRulePropertiesOutput
}

type ActionRulePropertiesArgs struct {
	Conditions        ConditionsPtrInput        `pulumi:"conditions"`
	Description       pulumi.StringPtrInput     `pulumi:"description"`
	Scope             ScopePtrInput             `pulumi:"scope"`
	Status            pulumi.StringPtrInput     `pulumi:"status"`
	SuppressionConfig SuppressionConfigPtrInput `pulumi:"suppressionConfig"`
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

func (o ActionRulePropertiesOutput) Conditions() ConditionsPtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *Conditions { return v.Conditions }).(ConditionsPtrOutput)
}

func (o ActionRulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesOutput) Scope() ScopePtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *Scope { return v.Scope }).(ScopePtrOutput)
}

func (o ActionRulePropertiesOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesOutput) SuppressionConfig() SuppressionConfigPtrOutput {
	return o.ApplyT(func(v ActionRuleProperties) *SuppressionConfig { return v.SuppressionConfig }).(SuppressionConfigPtrOutput)
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

func (o ActionRulePropertiesPtrOutput) Conditions() ConditionsPtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *Conditions {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(ConditionsPtrOutput)
}

func (o ActionRulePropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesPtrOutput) Scope() ScopePtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *Scope {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(ScopePtrOutput)
}

func (o ActionRulePropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesPtrOutput) SuppressionConfig() SuppressionConfigPtrOutput {
	return o.ApplyT(func(v *ActionRuleProperties) *SuppressionConfig {
		if v == nil {
			return nil
		}
		return v.SuppressionConfig
	}).(SuppressionConfigPtrOutput)
}

type ActionRulePropertiesResponse struct {
	Conditions        *ConditionsResponse        `pulumi:"conditions"`
	CreatedAt         string                     `pulumi:"createdAt"`
	CreatedBy         string                     `pulumi:"createdBy"`
	Description       *string                    `pulumi:"description"`
	LastModifiedAt    string                     `pulumi:"lastModifiedAt"`
	LastModifiedBy    string                     `pulumi:"lastModifiedBy"`
	ResourceGroup     string                     `pulumi:"resourceGroup"`
	Scope             *ScopeResponse             `pulumi:"scope"`
	Status            *string                    `pulumi:"status"`
	SuppressionConfig *SuppressionConfigResponse `pulumi:"suppressionConfig"`
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

func (o ActionRulePropertiesResponseOutput) Conditions() ConditionsResponsePtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *ConditionsResponse { return v.Conditions }).(ConditionsResponsePtrOutput)
}

func (o ActionRulePropertiesResponseOutput) CreatedAt() pulumi.StringOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) string { return v.CreatedAt }).(pulumi.StringOutput)
}

func (o ActionRulePropertiesResponseOutput) CreatedBy() pulumi.StringOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) string { return v.CreatedBy }).(pulumi.StringOutput)
}

func (o ActionRulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesResponseOutput) LastModifiedAt() pulumi.StringOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) string { return v.LastModifiedAt }).(pulumi.StringOutput)
}

func (o ActionRulePropertiesResponseOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o ActionRulePropertiesResponseOutput) ResourceGroup() pulumi.StringOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) string { return v.ResourceGroup }).(pulumi.StringOutput)
}

func (o ActionRulePropertiesResponseOutput) Scope() ScopeResponsePtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *ScopeResponse { return v.Scope }).(ScopeResponsePtrOutput)
}

func (o ActionRulePropertiesResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ActionRulePropertiesResponseOutput) SuppressionConfig() SuppressionConfigResponsePtrOutput {
	return o.ApplyT(func(v ActionRulePropertiesResponse) *SuppressionConfigResponse { return v.SuppressionConfig }).(SuppressionConfigResponsePtrOutput)
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
	AlertRuleId                      *Condition `pulumi:"alertRuleId"`
	ApplicationInsightsSearchResults *Condition `pulumi:"applicationInsightsSearchResults"`
	Description                      *Condition `pulumi:"description"`
	LogAnalyticsSearchResults        *Condition `pulumi:"logAnalyticsSearchResults"`
	MonitorCondition                 *Condition `pulumi:"monitorCondition"`
	MonitorService                   *Condition `pulumi:"monitorService"`
	Severity                         *Condition `pulumi:"severity"`
	SignalType                       *Condition `pulumi:"signalType"`
	TargetResource                   *Condition `pulumi:"targetResource"`
	TargetResourceGroup              *Condition `pulumi:"targetResourceGroup"`
	TargetResourceType               *Condition `pulumi:"targetResourceType"`
}





type ConditionsInput interface {
	pulumi.Input

	ToConditionsOutput() ConditionsOutput
	ToConditionsOutputWithContext(context.Context) ConditionsOutput
}

type ConditionsArgs struct {
	AlertRuleId                      ConditionPtrInput `pulumi:"alertRuleId"`
	ApplicationInsightsSearchResults ConditionPtrInput `pulumi:"applicationInsightsSearchResults"`
	Description                      ConditionPtrInput `pulumi:"description"`
	LogAnalyticsSearchResults        ConditionPtrInput `pulumi:"logAnalyticsSearchResults"`
	MonitorCondition                 ConditionPtrInput `pulumi:"monitorCondition"`
	MonitorService                   ConditionPtrInput `pulumi:"monitorService"`
	Severity                         ConditionPtrInput `pulumi:"severity"`
	SignalType                       ConditionPtrInput `pulumi:"signalType"`
	TargetResource                   ConditionPtrInput `pulumi:"targetResource"`
	TargetResourceGroup              ConditionPtrInput `pulumi:"targetResourceGroup"`
	TargetResourceType               ConditionPtrInput `pulumi:"targetResourceType"`
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

func (o ConditionsOutput) AlertRuleId() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.AlertRuleId }).(ConditionPtrOutput)
}

func (o ConditionsOutput) ApplicationInsightsSearchResults() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.ApplicationInsightsSearchResults }).(ConditionPtrOutput)
}

func (o ConditionsOutput) Description() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.Description }).(ConditionPtrOutput)
}

func (o ConditionsOutput) LogAnalyticsSearchResults() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.LogAnalyticsSearchResults }).(ConditionPtrOutput)
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

func (o ConditionsOutput) SignalType() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.SignalType }).(ConditionPtrOutput)
}

func (o ConditionsOutput) TargetResource() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.TargetResource }).(ConditionPtrOutput)
}

func (o ConditionsOutput) TargetResourceGroup() ConditionPtrOutput {
	return o.ApplyT(func(v Conditions) *Condition { return v.TargetResourceGroup }).(ConditionPtrOutput)
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

func (o ConditionsPtrOutput) AlertRuleId() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.AlertRuleId
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) ApplicationInsightsSearchResults() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.ApplicationInsightsSearchResults
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

func (o ConditionsPtrOutput) LogAnalyticsSearchResults() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsSearchResults
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

func (o ConditionsPtrOutput) SignalType() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.SignalType
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) TargetResource() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.TargetResource
	}).(ConditionPtrOutput)
}

func (o ConditionsPtrOutput) TargetResourceGroup() ConditionPtrOutput {
	return o.ApplyT(func(v *Conditions) *Condition {
		if v == nil {
			return nil
		}
		return v.TargetResourceGroup
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
	AlertRuleId                      *ConditionResponse `pulumi:"alertRuleId"`
	ApplicationInsightsSearchResults *ConditionResponse `pulumi:"applicationInsightsSearchResults"`
	Description                      *ConditionResponse `pulumi:"description"`
	LogAnalyticsSearchResults        *ConditionResponse `pulumi:"logAnalyticsSearchResults"`
	MonitorCondition                 *ConditionResponse `pulumi:"monitorCondition"`
	MonitorService                   *ConditionResponse `pulumi:"monitorService"`
	Severity                         *ConditionResponse `pulumi:"severity"`
	SignalType                       *ConditionResponse `pulumi:"signalType"`
	TargetResource                   *ConditionResponse `pulumi:"targetResource"`
	TargetResourceGroup              *ConditionResponse `pulumi:"targetResourceGroup"`
	TargetResourceType               *ConditionResponse `pulumi:"targetResourceType"`
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

func (o ConditionsResponseOutput) AlertRuleId() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.AlertRuleId }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) ApplicationInsightsSearchResults() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.ApplicationInsightsSearchResults }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) Description() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.Description }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) LogAnalyticsSearchResults() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.LogAnalyticsSearchResults }).(ConditionResponsePtrOutput)
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

func (o ConditionsResponseOutput) SignalType() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.SignalType }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) TargetResource() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.TargetResource }).(ConditionResponsePtrOutput)
}

func (o ConditionsResponseOutput) TargetResourceGroup() ConditionResponsePtrOutput {
	return o.ApplyT(func(v ConditionsResponse) *ConditionResponse { return v.TargetResourceGroup }).(ConditionResponsePtrOutput)
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

func (o ConditionsResponsePtrOutput) AlertRuleId() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.AlertRuleId
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) ApplicationInsightsSearchResults() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.ApplicationInsightsSearchResults
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

func (o ConditionsResponsePtrOutput) LogAnalyticsSearchResults() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsSearchResults
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

func (o ConditionsResponsePtrOutput) SignalType() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.SignalType
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) TargetResource() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.TargetResource
	}).(ConditionResponsePtrOutput)
}

func (o ConditionsResponsePtrOutput) TargetResourceGroup() ConditionResponsePtrOutput {
	return o.ApplyT(func(v *ConditionsResponse) *ConditionResponse {
		if v == nil {
			return nil
		}
		return v.TargetResourceGroup
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

type Scope struct {
	Type   *string  `pulumi:"type"`
	Values []string `pulumi:"values"`
}





type ScopeInput interface {
	pulumi.Input

	ToScopeOutput() ScopeOutput
	ToScopeOutputWithContext(context.Context) ScopeOutput
}

type ScopeArgs struct {
	Type   pulumi.StringPtrInput   `pulumi:"type"`
	Values pulumi.StringArrayInput `pulumi:"values"`
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

func (o ScopeOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Scope) *string { return v.Type }).(pulumi.StringPtrOutput)
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

func (o ScopePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Scope) *string {
		if v == nil {
			return nil
		}
		return v.Type
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
	Type   *string  `pulumi:"type"`
	Values []string `pulumi:"values"`
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

func (o ScopeResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScopeResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
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

func (o ScopeResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScopeResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
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

func (i SuppressionConfigArgs) ToSuppressionConfigPtrOutput() SuppressionConfigPtrOutput {
	return i.ToSuppressionConfigPtrOutputWithContext(context.Background())
}

func (i SuppressionConfigArgs) ToSuppressionConfigPtrOutputWithContext(ctx context.Context) SuppressionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionConfigOutput).ToSuppressionConfigPtrOutputWithContext(ctx)
}









type SuppressionConfigPtrInput interface {
	pulumi.Input

	ToSuppressionConfigPtrOutput() SuppressionConfigPtrOutput
	ToSuppressionConfigPtrOutputWithContext(context.Context) SuppressionConfigPtrOutput
}

type suppressionConfigPtrType SuppressionConfigArgs

func SuppressionConfigPtr(v *SuppressionConfigArgs) SuppressionConfigPtrInput {
	return (*suppressionConfigPtrType)(v)
}

func (*suppressionConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionConfig)(nil)).Elem()
}

func (i *suppressionConfigPtrType) ToSuppressionConfigPtrOutput() SuppressionConfigPtrOutput {
	return i.ToSuppressionConfigPtrOutputWithContext(context.Background())
}

func (i *suppressionConfigPtrType) ToSuppressionConfigPtrOutputWithContext(ctx context.Context) SuppressionConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SuppressionConfigPtrOutput)
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

func (o SuppressionConfigOutput) ToSuppressionConfigPtrOutput() SuppressionConfigPtrOutput {
	return o.ToSuppressionConfigPtrOutputWithContext(context.Background())
}

func (o SuppressionConfigOutput) ToSuppressionConfigPtrOutputWithContext(ctx context.Context) SuppressionConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SuppressionConfig) *SuppressionConfig {
		return &v
	}).(SuppressionConfigPtrOutput)
}

func (o SuppressionConfigOutput) RecurrenceType() pulumi.StringOutput {
	return o.ApplyT(func(v SuppressionConfig) string { return v.RecurrenceType }).(pulumi.StringOutput)
}

func (o SuppressionConfigOutput) Schedule() SuppressionSchedulePtrOutput {
	return o.ApplyT(func(v SuppressionConfig) *SuppressionSchedule { return v.Schedule }).(SuppressionSchedulePtrOutput)
}

type SuppressionConfigPtrOutput struct{ *pulumi.OutputState }

func (SuppressionConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionConfig)(nil)).Elem()
}

func (o SuppressionConfigPtrOutput) ToSuppressionConfigPtrOutput() SuppressionConfigPtrOutput {
	return o
}

func (o SuppressionConfigPtrOutput) ToSuppressionConfigPtrOutputWithContext(ctx context.Context) SuppressionConfigPtrOutput {
	return o
}

func (o SuppressionConfigPtrOutput) Elem() SuppressionConfigOutput {
	return o.ApplyT(func(v *SuppressionConfig) SuppressionConfig {
		if v != nil {
			return *v
		}
		var ret SuppressionConfig
		return ret
	}).(SuppressionConfigOutput)
}

func (o SuppressionConfigPtrOutput) RecurrenceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionConfig) *string {
		if v == nil {
			return nil
		}
		return &v.RecurrenceType
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionConfigPtrOutput) Schedule() SuppressionSchedulePtrOutput {
	return o.ApplyT(func(v *SuppressionConfig) *SuppressionSchedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(SuppressionSchedulePtrOutput)
}

type SuppressionConfigResponse struct {
	RecurrenceType string                       `pulumi:"recurrenceType"`
	Schedule       *SuppressionScheduleResponse `pulumi:"schedule"`
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

type SuppressionConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (SuppressionConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SuppressionConfigResponse)(nil)).Elem()
}

func (o SuppressionConfigResponsePtrOutput) ToSuppressionConfigResponsePtrOutput() SuppressionConfigResponsePtrOutput {
	return o
}

func (o SuppressionConfigResponsePtrOutput) ToSuppressionConfigResponsePtrOutputWithContext(ctx context.Context) SuppressionConfigResponsePtrOutput {
	return o
}

func (o SuppressionConfigResponsePtrOutput) Elem() SuppressionConfigResponseOutput {
	return o.ApplyT(func(v *SuppressionConfigResponse) SuppressionConfigResponse {
		if v != nil {
			return *v
		}
		var ret SuppressionConfigResponse
		return ret
	}).(SuppressionConfigResponseOutput)
}

func (o SuppressionConfigResponsePtrOutput) RecurrenceType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SuppressionConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RecurrenceType
	}).(pulumi.StringPtrOutput)
}

func (o SuppressionConfigResponsePtrOutput) Schedule() SuppressionScheduleResponsePtrOutput {
	return o.ApplyT(func(v *SuppressionConfigResponse) *SuppressionScheduleResponse {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(SuppressionScheduleResponsePtrOutput)
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
	pulumi.RegisterOutputType(ActionRulePropertiesOutput{})
	pulumi.RegisterOutputType(ActionRulePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ActionRulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConditionOutput{})
	pulumi.RegisterOutputType(ConditionPtrOutput{})
	pulumi.RegisterOutputType(ConditionResponseOutput{})
	pulumi.RegisterOutputType(ConditionResponsePtrOutput{})
	pulumi.RegisterOutputType(ConditionsOutput{})
	pulumi.RegisterOutputType(ConditionsPtrOutput{})
	pulumi.RegisterOutputType(ConditionsResponseOutput{})
	pulumi.RegisterOutputType(ConditionsResponsePtrOutput{})
	pulumi.RegisterOutputType(ScopeOutput{})
	pulumi.RegisterOutputType(ScopePtrOutput{})
	pulumi.RegisterOutputType(ScopeResponseOutput{})
	pulumi.RegisterOutputType(ScopeResponsePtrOutput{})
	pulumi.RegisterOutputType(SuppressionConfigOutput{})
	pulumi.RegisterOutputType(SuppressionConfigPtrOutput{})
	pulumi.RegisterOutputType(SuppressionConfigResponseOutput{})
	pulumi.RegisterOutputType(SuppressionConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(SuppressionScheduleOutput{})
	pulumi.RegisterOutputType(SuppressionSchedulePtrOutput{})
	pulumi.RegisterOutputType(SuppressionScheduleResponseOutput{})
	pulumi.RegisterOutputType(SuppressionScheduleResponsePtrOutput{})
}
