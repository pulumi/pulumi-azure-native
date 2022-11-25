


package v20221101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertDetailsOverride struct {
	AlertDescriptionFormat  *string `pulumi:"alertDescriptionFormat"`
	AlertDisplayNameFormat  *string `pulumi:"alertDisplayNameFormat"`
	AlertSeverityColumnName *string `pulumi:"alertSeverityColumnName"`
	AlertTacticsColumnName  *string `pulumi:"alertTacticsColumnName"`
}





type AlertDetailsOverrideInput interface {
	pulumi.Input

	ToAlertDetailsOverrideOutput() AlertDetailsOverrideOutput
	ToAlertDetailsOverrideOutputWithContext(context.Context) AlertDetailsOverrideOutput
}

type AlertDetailsOverrideArgs struct {
	AlertDescriptionFormat  pulumi.StringPtrInput `pulumi:"alertDescriptionFormat"`
	AlertDisplayNameFormat  pulumi.StringPtrInput `pulumi:"alertDisplayNameFormat"`
	AlertSeverityColumnName pulumi.StringPtrInput `pulumi:"alertSeverityColumnName"`
	AlertTacticsColumnName  pulumi.StringPtrInput `pulumi:"alertTacticsColumnName"`
}

func (AlertDetailsOverrideArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDetailsOverride)(nil)).Elem()
}

func (i AlertDetailsOverrideArgs) ToAlertDetailsOverrideOutput() AlertDetailsOverrideOutput {
	return i.ToAlertDetailsOverrideOutputWithContext(context.Background())
}

func (i AlertDetailsOverrideArgs) ToAlertDetailsOverrideOutputWithContext(ctx context.Context) AlertDetailsOverrideOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDetailsOverrideOutput)
}

func (i AlertDetailsOverrideArgs) ToAlertDetailsOverridePtrOutput() AlertDetailsOverridePtrOutput {
	return i.ToAlertDetailsOverridePtrOutputWithContext(context.Background())
}

func (i AlertDetailsOverrideArgs) ToAlertDetailsOverridePtrOutputWithContext(ctx context.Context) AlertDetailsOverridePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDetailsOverrideOutput).ToAlertDetailsOverridePtrOutputWithContext(ctx)
}









type AlertDetailsOverridePtrInput interface {
	pulumi.Input

	ToAlertDetailsOverridePtrOutput() AlertDetailsOverridePtrOutput
	ToAlertDetailsOverridePtrOutputWithContext(context.Context) AlertDetailsOverridePtrOutput
}

type alertDetailsOverridePtrType AlertDetailsOverrideArgs

func AlertDetailsOverridePtr(v *AlertDetailsOverrideArgs) AlertDetailsOverridePtrInput {
	return (*alertDetailsOverridePtrType)(v)
}

func (*alertDetailsOverridePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDetailsOverride)(nil)).Elem()
}

func (i *alertDetailsOverridePtrType) ToAlertDetailsOverridePtrOutput() AlertDetailsOverridePtrOutput {
	return i.ToAlertDetailsOverridePtrOutputWithContext(context.Background())
}

func (i *alertDetailsOverridePtrType) ToAlertDetailsOverridePtrOutputWithContext(ctx context.Context) AlertDetailsOverridePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDetailsOverridePtrOutput)
}

type AlertDetailsOverrideOutput struct{ *pulumi.OutputState }

func (AlertDetailsOverrideOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDetailsOverride)(nil)).Elem()
}

func (o AlertDetailsOverrideOutput) ToAlertDetailsOverrideOutput() AlertDetailsOverrideOutput {
	return o
}

func (o AlertDetailsOverrideOutput) ToAlertDetailsOverrideOutputWithContext(ctx context.Context) AlertDetailsOverrideOutput {
	return o
}

func (o AlertDetailsOverrideOutput) ToAlertDetailsOverridePtrOutput() AlertDetailsOverridePtrOutput {
	return o.ToAlertDetailsOverridePtrOutputWithContext(context.Background())
}

func (o AlertDetailsOverrideOutput) ToAlertDetailsOverridePtrOutputWithContext(ctx context.Context) AlertDetailsOverridePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertDetailsOverride) *AlertDetailsOverride {
		return &v
	}).(AlertDetailsOverridePtrOutput)
}

func (o AlertDetailsOverrideOutput) AlertDescriptionFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverride) *string { return v.AlertDescriptionFormat }).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideOutput) AlertDisplayNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverride) *string { return v.AlertDisplayNameFormat }).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideOutput) AlertSeverityColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverride) *string { return v.AlertSeverityColumnName }).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideOutput) AlertTacticsColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverride) *string { return v.AlertTacticsColumnName }).(pulumi.StringPtrOutput)
}

type AlertDetailsOverridePtrOutput struct{ *pulumi.OutputState }

func (AlertDetailsOverridePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDetailsOverride)(nil)).Elem()
}

func (o AlertDetailsOverridePtrOutput) ToAlertDetailsOverridePtrOutput() AlertDetailsOverridePtrOutput {
	return o
}

func (o AlertDetailsOverridePtrOutput) ToAlertDetailsOverridePtrOutputWithContext(ctx context.Context) AlertDetailsOverridePtrOutput {
	return o
}

func (o AlertDetailsOverridePtrOutput) Elem() AlertDetailsOverrideOutput {
	return o.ApplyT(func(v *AlertDetailsOverride) AlertDetailsOverride {
		if v != nil {
			return *v
		}
		var ret AlertDetailsOverride
		return ret
	}).(AlertDetailsOverrideOutput)
}

func (o AlertDetailsOverridePtrOutput) AlertDescriptionFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverride) *string {
		if v == nil {
			return nil
		}
		return v.AlertDescriptionFormat
	}).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverridePtrOutput) AlertDisplayNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverride) *string {
		if v == nil {
			return nil
		}
		return v.AlertDisplayNameFormat
	}).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverridePtrOutput) AlertSeverityColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverride) *string {
		if v == nil {
			return nil
		}
		return v.AlertSeverityColumnName
	}).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverridePtrOutput) AlertTacticsColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverride) *string {
		if v == nil {
			return nil
		}
		return v.AlertTacticsColumnName
	}).(pulumi.StringPtrOutput)
}

type AlertDetailsOverrideResponse struct {
	AlertDescriptionFormat  *string `pulumi:"alertDescriptionFormat"`
	AlertDisplayNameFormat  *string `pulumi:"alertDisplayNameFormat"`
	AlertSeverityColumnName *string `pulumi:"alertSeverityColumnName"`
	AlertTacticsColumnName  *string `pulumi:"alertTacticsColumnName"`
}

type AlertDetailsOverrideResponseOutput struct{ *pulumi.OutputState }

func (AlertDetailsOverrideResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDetailsOverrideResponse)(nil)).Elem()
}

func (o AlertDetailsOverrideResponseOutput) ToAlertDetailsOverrideResponseOutput() AlertDetailsOverrideResponseOutput {
	return o
}

func (o AlertDetailsOverrideResponseOutput) ToAlertDetailsOverrideResponseOutputWithContext(ctx context.Context) AlertDetailsOverrideResponseOutput {
	return o
}

func (o AlertDetailsOverrideResponseOutput) AlertDescriptionFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverrideResponse) *string { return v.AlertDescriptionFormat }).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideResponseOutput) AlertDisplayNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverrideResponse) *string { return v.AlertDisplayNameFormat }).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideResponseOutput) AlertSeverityColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverrideResponse) *string { return v.AlertSeverityColumnName }).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideResponseOutput) AlertTacticsColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AlertDetailsOverrideResponse) *string { return v.AlertTacticsColumnName }).(pulumi.StringPtrOutput)
}

type AlertDetailsOverrideResponsePtrOutput struct{ *pulumi.OutputState }

func (AlertDetailsOverrideResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDetailsOverrideResponse)(nil)).Elem()
}

func (o AlertDetailsOverrideResponsePtrOutput) ToAlertDetailsOverrideResponsePtrOutput() AlertDetailsOverrideResponsePtrOutput {
	return o
}

func (o AlertDetailsOverrideResponsePtrOutput) ToAlertDetailsOverrideResponsePtrOutputWithContext(ctx context.Context) AlertDetailsOverrideResponsePtrOutput {
	return o
}

func (o AlertDetailsOverrideResponsePtrOutput) Elem() AlertDetailsOverrideResponseOutput {
	return o.ApplyT(func(v *AlertDetailsOverrideResponse) AlertDetailsOverrideResponse {
		if v != nil {
			return *v
		}
		var ret AlertDetailsOverrideResponse
		return ret
	}).(AlertDetailsOverrideResponseOutput)
}

func (o AlertDetailsOverrideResponsePtrOutput) AlertDescriptionFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverrideResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertDescriptionFormat
	}).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideResponsePtrOutput) AlertDisplayNameFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverrideResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertDisplayNameFormat
	}).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideResponsePtrOutput) AlertSeverityColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverrideResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertSeverityColumnName
	}).(pulumi.StringPtrOutput)
}

func (o AlertDetailsOverrideResponsePtrOutput) AlertTacticsColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertDetailsOverrideResponse) *string {
		if v == nil {
			return nil
		}
		return v.AlertTacticsColumnName
	}).(pulumi.StringPtrOutput)
}

type AlertsDataTypeOfDataConnector struct {
	Alerts *DataConnectorDataTypeCommon `pulumi:"alerts"`
}





type AlertsDataTypeOfDataConnectorInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorOutput() AlertsDataTypeOfDataConnectorOutput
	ToAlertsDataTypeOfDataConnectorOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorOutput
}

type AlertsDataTypeOfDataConnectorArgs struct {
	Alerts DataConnectorDataTypeCommonPtrInput `pulumi:"alerts"`
}

func (AlertsDataTypeOfDataConnectorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnector)(nil)).Elem()
}

func (i AlertsDataTypeOfDataConnectorArgs) ToAlertsDataTypeOfDataConnectorOutput() AlertsDataTypeOfDataConnectorOutput {
	return i.ToAlertsDataTypeOfDataConnectorOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorArgs) ToAlertsDataTypeOfDataConnectorOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorOutput)
}

func (i AlertsDataTypeOfDataConnectorArgs) ToAlertsDataTypeOfDataConnectorPtrOutput() AlertsDataTypeOfDataConnectorPtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorArgs) ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorOutput).ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(ctx)
}









type AlertsDataTypeOfDataConnectorPtrInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorPtrOutput() AlertsDataTypeOfDataConnectorPtrOutput
	ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorPtrOutput
}

type alertsDataTypeOfDataConnectorPtrType AlertsDataTypeOfDataConnectorArgs

func AlertsDataTypeOfDataConnectorPtr(v *AlertsDataTypeOfDataConnectorArgs) AlertsDataTypeOfDataConnectorPtrInput {
	return (*alertsDataTypeOfDataConnectorPtrType)(v)
}

func (*alertsDataTypeOfDataConnectorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnector)(nil)).Elem()
}

func (i *alertsDataTypeOfDataConnectorPtrType) ToAlertsDataTypeOfDataConnectorPtrOutput() AlertsDataTypeOfDataConnectorPtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(context.Background())
}

func (i *alertsDataTypeOfDataConnectorPtrType) ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorPtrOutput)
}

type AlertsDataTypeOfDataConnectorOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnector)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorOutput) ToAlertsDataTypeOfDataConnectorOutput() AlertsDataTypeOfDataConnectorOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorOutput) ToAlertsDataTypeOfDataConnectorOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorOutput) ToAlertsDataTypeOfDataConnectorPtrOutput() AlertsDataTypeOfDataConnectorPtrOutput {
	return o.ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(context.Background())
}

func (o AlertsDataTypeOfDataConnectorOutput) ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertsDataTypeOfDataConnector) *AlertsDataTypeOfDataConnector {
		return &v
	}).(AlertsDataTypeOfDataConnectorPtrOutput)
}

func (o AlertsDataTypeOfDataConnectorOutput) Alerts() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnector) *DataConnectorDataTypeCommon { return v.Alerts }).(DataConnectorDataTypeCommonPtrOutput)
}

type AlertsDataTypeOfDataConnectorPtrOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnector)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorPtrOutput) ToAlertsDataTypeOfDataConnectorPtrOutput() AlertsDataTypeOfDataConnectorPtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorPtrOutput) ToAlertsDataTypeOfDataConnectorPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorPtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorPtrOutput) Elem() AlertsDataTypeOfDataConnectorOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnector) AlertsDataTypeOfDataConnector {
		if v != nil {
			return *v
		}
		var ret AlertsDataTypeOfDataConnector
		return ret
	}).(AlertsDataTypeOfDataConnectorOutput)
}

func (o AlertsDataTypeOfDataConnectorPtrOutput) Alerts() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnector) *DataConnectorDataTypeCommon {
		if v == nil {
			return nil
		}
		return v.Alerts
	}).(DataConnectorDataTypeCommonPtrOutput)
}

type AlertsDataTypeOfDataConnectorResponse struct {
	Alerts *DataConnectorDataTypeCommonResponse `pulumi:"alerts"`
}

type AlertsDataTypeOfDataConnectorResponseOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponse)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorResponseOutput) ToAlertsDataTypeOfDataConnectorResponseOutput() AlertsDataTypeOfDataConnectorResponseOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponseOutput) ToAlertsDataTypeOfDataConnectorResponseOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponseOutput) Alerts() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnectorResponse) *DataConnectorDataTypeCommonResponse { return v.Alerts }).(DataConnectorDataTypeCommonResponsePtrOutput)
}

type AlertsDataTypeOfDataConnectorResponsePtrOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnectorResponse)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorResponsePtrOutput) ToAlertsDataTypeOfDataConnectorResponsePtrOutput() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponsePtrOutput) ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponsePtrOutput) Elem() AlertsDataTypeOfDataConnectorResponseOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorResponse) AlertsDataTypeOfDataConnectorResponse {
		if v != nil {
			return *v
		}
		var ret AlertsDataTypeOfDataConnectorResponse
		return ret
	}).(AlertsDataTypeOfDataConnectorResponseOutput)
}

func (o AlertsDataTypeOfDataConnectorResponsePtrOutput) Alerts() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorResponse) *DataConnectorDataTypeCommonResponse {
		if v == nil {
			return nil
		}
		return v.Alerts
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
}

type AutomationRuleModifyPropertiesAction struct {
	ActionConfiguration *IncidentPropertiesAction `pulumi:"actionConfiguration"`
	ActionType          string                    `pulumi:"actionType"`
	Order               int                       `pulumi:"order"`
}

type AutomationRuleModifyPropertiesActionResponse struct {
	ActionConfiguration *IncidentPropertiesActionResponse `pulumi:"actionConfiguration"`
	ActionType          string                            `pulumi:"actionType"`
	Order               int                               `pulumi:"order"`
}

type AutomationRulePropertyArrayChangedValuesCondition struct {
	ArrayType  *string `pulumi:"arrayType"`
	ChangeType *string `pulumi:"changeType"`
}

type AutomationRulePropertyArrayChangedValuesConditionResponse struct {
	ArrayType  *string `pulumi:"arrayType"`
	ChangeType *string `pulumi:"changeType"`
}

type AutomationRulePropertyValuesChangedCondition struct {
	ChangeType     *string  `pulumi:"changeType"`
	Operator       *string  `pulumi:"operator"`
	PropertyName   *string  `pulumi:"propertyName"`
	PropertyValues []string `pulumi:"propertyValues"`
}

type AutomationRulePropertyValuesChangedConditionResponse struct {
	ChangeType     *string  `pulumi:"changeType"`
	Operator       *string  `pulumi:"operator"`
	PropertyName   *string  `pulumi:"propertyName"`
	PropertyValues []string `pulumi:"propertyValues"`
}

type AutomationRulePropertyValuesCondition struct {
	Operator       *string  `pulumi:"operator"`
	PropertyName   *string  `pulumi:"propertyName"`
	PropertyValues []string `pulumi:"propertyValues"`
}

type AutomationRulePropertyValuesConditionResponse struct {
	Operator       *string  `pulumi:"operator"`
	PropertyName   *string  `pulumi:"propertyName"`
	PropertyValues []string `pulumi:"propertyValues"`
}

type AutomationRuleRunPlaybookAction struct {
	ActionConfiguration *PlaybookActionProperties `pulumi:"actionConfiguration"`
	ActionType          string                    `pulumi:"actionType"`
	Order               int                       `pulumi:"order"`
}

type AutomationRuleRunPlaybookActionResponse struct {
	ActionConfiguration *PlaybookActionPropertiesResponse `pulumi:"actionConfiguration"`
	ActionType          string                            `pulumi:"actionType"`
	Order               int                               `pulumi:"order"`
}

type AutomationRuleTriggeringLogic struct {
	Conditions        []interface{} `pulumi:"conditions"`
	ExpirationTimeUtc *string       `pulumi:"expirationTimeUtc"`
	IsEnabled         bool          `pulumi:"isEnabled"`
	TriggersOn        string        `pulumi:"triggersOn"`
	TriggersWhen      string        `pulumi:"triggersWhen"`
}





type AutomationRuleTriggeringLogicInput interface {
	pulumi.Input

	ToAutomationRuleTriggeringLogicOutput() AutomationRuleTriggeringLogicOutput
	ToAutomationRuleTriggeringLogicOutputWithContext(context.Context) AutomationRuleTriggeringLogicOutput
}

type AutomationRuleTriggeringLogicArgs struct {
	Conditions        pulumi.ArrayInput     `pulumi:"conditions"`
	ExpirationTimeUtc pulumi.StringPtrInput `pulumi:"expirationTimeUtc"`
	IsEnabled         pulumi.BoolInput      `pulumi:"isEnabled"`
	TriggersOn        pulumi.StringInput    `pulumi:"triggersOn"`
	TriggersWhen      pulumi.StringInput    `pulumi:"triggersWhen"`
}

func (AutomationRuleTriggeringLogicArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleTriggeringLogic)(nil)).Elem()
}

func (i AutomationRuleTriggeringLogicArgs) ToAutomationRuleTriggeringLogicOutput() AutomationRuleTriggeringLogicOutput {
	return i.ToAutomationRuleTriggeringLogicOutputWithContext(context.Background())
}

func (i AutomationRuleTriggeringLogicArgs) ToAutomationRuleTriggeringLogicOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleTriggeringLogicOutput)
}

type AutomationRuleTriggeringLogicOutput struct{ *pulumi.OutputState }

func (AutomationRuleTriggeringLogicOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleTriggeringLogic)(nil)).Elem()
}

func (o AutomationRuleTriggeringLogicOutput) ToAutomationRuleTriggeringLogicOutput() AutomationRuleTriggeringLogicOutput {
	return o
}

func (o AutomationRuleTriggeringLogicOutput) ToAutomationRuleTriggeringLogicOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicOutput {
	return o
}

func (o AutomationRuleTriggeringLogicOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogic) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
}

func (o AutomationRuleTriggeringLogicOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogic) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleTriggeringLogicOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogic) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AutomationRuleTriggeringLogicOutput) TriggersOn() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogic) string { return v.TriggersOn }).(pulumi.StringOutput)
}

func (o AutomationRuleTriggeringLogicOutput) TriggersWhen() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogic) string { return v.TriggersWhen }).(pulumi.StringOutput)
}

type AutomationRuleTriggeringLogicResponse struct {
	Conditions        []interface{} `pulumi:"conditions"`
	ExpirationTimeUtc *string       `pulumi:"expirationTimeUtc"`
	IsEnabled         bool          `pulumi:"isEnabled"`
	TriggersOn        string        `pulumi:"triggersOn"`
	TriggersWhen      string        `pulumi:"triggersWhen"`
}

type AutomationRuleTriggeringLogicResponseOutput struct{ *pulumi.OutputState }

func (AutomationRuleTriggeringLogicResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleTriggeringLogicResponse)(nil)).Elem()
}

func (o AutomationRuleTriggeringLogicResponseOutput) ToAutomationRuleTriggeringLogicResponseOutput() AutomationRuleTriggeringLogicResponseOutput {
	return o
}

func (o AutomationRuleTriggeringLogicResponseOutput) ToAutomationRuleTriggeringLogicResponseOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicResponseOutput {
	return o
}

func (o AutomationRuleTriggeringLogicResponseOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogicResponse) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
}

func (o AutomationRuleTriggeringLogicResponseOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogicResponse) *string { return v.ExpirationTimeUtc }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleTriggeringLogicResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogicResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AutomationRuleTriggeringLogicResponseOutput) TriggersOn() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogicResponse) string { return v.TriggersOn }).(pulumi.StringOutput)
}

func (o AutomationRuleTriggeringLogicResponseOutput) TriggersWhen() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogicResponse) string { return v.TriggersWhen }).(pulumi.StringOutput)
}

type AwsCloudTrailDataConnectorDataTypes struct {
	Logs *AwsCloudTrailDataConnectorDataTypesLogs `pulumi:"logs"`
}





type AwsCloudTrailDataConnectorDataTypesInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesOutput() AwsCloudTrailDataConnectorDataTypesOutput
	ToAwsCloudTrailDataConnectorDataTypesOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesOutput
}

type AwsCloudTrailDataConnectorDataTypesArgs struct {
	Logs AwsCloudTrailDataConnectorDataTypesLogsPtrInput `pulumi:"logs"`
}

func (AwsCloudTrailDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypes)(nil)).Elem()
}

func (i AwsCloudTrailDataConnectorDataTypesArgs) ToAwsCloudTrailDataConnectorDataTypesOutput() AwsCloudTrailDataConnectorDataTypesOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesArgs) ToAwsCloudTrailDataConnectorDataTypesOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesOutput)
}

func (i AwsCloudTrailDataConnectorDataTypesArgs) ToAwsCloudTrailDataConnectorDataTypesPtrOutput() AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesArgs) ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesOutput).ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type AwsCloudTrailDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesPtrOutput() AwsCloudTrailDataConnectorDataTypesPtrOutput
	ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesPtrOutput
}

type awsCloudTrailDataConnectorDataTypesPtrType AwsCloudTrailDataConnectorDataTypesArgs

func AwsCloudTrailDataConnectorDataTypesPtr(v *AwsCloudTrailDataConnectorDataTypesArgs) AwsCloudTrailDataConnectorDataTypesPtrInput {
	return (*awsCloudTrailDataConnectorDataTypesPtrType)(v)
}

func (*awsCloudTrailDataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypes)(nil)).Elem()
}

func (i *awsCloudTrailDataConnectorDataTypesPtrType) ToAwsCloudTrailDataConnectorDataTypesPtrOutput() AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *awsCloudTrailDataConnectorDataTypesPtrType) ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypes)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesOutput) ToAwsCloudTrailDataConnectorDataTypesOutput() AwsCloudTrailDataConnectorDataTypesOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesOutput) ToAwsCloudTrailDataConnectorDataTypesOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesOutput) ToAwsCloudTrailDataConnectorDataTypesPtrOutput() AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return o.ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o AwsCloudTrailDataConnectorDataTypesOutput) ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AwsCloudTrailDataConnectorDataTypes) *AwsCloudTrailDataConnectorDataTypes {
		return &v
	}).(AwsCloudTrailDataConnectorDataTypesPtrOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesOutput) Logs() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypes) *AwsCloudTrailDataConnectorDataTypesLogs { return v.Logs }).(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypes)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesPtrOutput) ToAwsCloudTrailDataConnectorDataTypesPtrOutput() AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesPtrOutput) ToAwsCloudTrailDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesPtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesPtrOutput) Elem() AwsCloudTrailDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypes) AwsCloudTrailDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret AwsCloudTrailDataConnectorDataTypes
		return ret
	}).(AwsCloudTrailDataConnectorDataTypesOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesPtrOutput) Logs() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypes) *AwsCloudTrailDataConnectorDataTypesLogs {
		if v == nil {
			return nil
		}
		return v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesLogs struct {
	State *string `pulumi:"state"`
}





type AwsCloudTrailDataConnectorDataTypesLogsInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesLogsOutput() AwsCloudTrailDataConnectorDataTypesLogsOutput
	ToAwsCloudTrailDataConnectorDataTypesLogsOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesLogsOutput
}

type AwsCloudTrailDataConnectorDataTypesLogsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (AwsCloudTrailDataConnectorDataTypesLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesLogs)(nil)).Elem()
}

func (i AwsCloudTrailDataConnectorDataTypesLogsArgs) ToAwsCloudTrailDataConnectorDataTypesLogsOutput() AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesLogsOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesLogsArgs) ToAwsCloudTrailDataConnectorDataTypesLogsOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesLogsOutput)
}

func (i AwsCloudTrailDataConnectorDataTypesLogsArgs) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesLogsArgs) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesLogsOutput).ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(ctx)
}









type AwsCloudTrailDataConnectorDataTypesLogsPtrInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput
	ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesLogsPtrOutput
}

type awsCloudTrailDataConnectorDataTypesLogsPtrType AwsCloudTrailDataConnectorDataTypesLogsArgs

func AwsCloudTrailDataConnectorDataTypesLogsPtr(v *AwsCloudTrailDataConnectorDataTypesLogsArgs) AwsCloudTrailDataConnectorDataTypesLogsPtrInput {
	return (*awsCloudTrailDataConnectorDataTypesLogsPtrType)(v)
}

func (*awsCloudTrailDataConnectorDataTypesLogsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypesLogs)(nil)).Elem()
}

func (i *awsCloudTrailDataConnectorDataTypesLogsPtrType) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(context.Background())
}

func (i *awsCloudTrailDataConnectorDataTypesLogsPtrType) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesLogsOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesLogs)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) ToAwsCloudTrailDataConnectorDataTypesLogsOutput() AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) ToAwsCloudTrailDataConnectorDataTypesLogsOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return o.ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(context.Background())
}

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AwsCloudTrailDataConnectorDataTypesLogs) *AwsCloudTrailDataConnectorDataTypesLogs {
		return &v
	}).(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesLogs) *string { return v.State }).(pulumi.StringPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesLogsPtrOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesLogsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypesLogs)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesLogsPtrOutput) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesLogsPtrOutput) ToAwsCloudTrailDataConnectorDataTypesLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesLogsPtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesLogsPtrOutput) Elem() AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypesLogs) AwsCloudTrailDataConnectorDataTypesLogs {
		if v != nil {
			return *v
		}
		var ret AwsCloudTrailDataConnectorDataTypesLogs
		return ret
	}).(AwsCloudTrailDataConnectorDataTypesLogsOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesLogsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypesLogs) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponse struct {
	Logs *AwsCloudTrailDataConnectorDataTypesResponseLogs `pulumi:"logs"`
}

type AwsCloudTrailDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) ToAwsCloudTrailDataConnectorDataTypesResponseOutput() AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) ToAwsCloudTrailDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) Logs() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesResponse) *AwsCloudTrailDataConnectorDataTypesResponseLogs {
		return v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesResponsePtrOutput) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutput() AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponsePtrOutput) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponsePtrOutput) Elem() AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypesResponse) AwsCloudTrailDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret AwsCloudTrailDataConnectorDataTypesResponse
		return ret
	}).(AwsCloudTrailDataConnectorDataTypesResponseOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesResponsePtrOutput) Logs() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypesResponse) *AwsCloudTrailDataConnectorDataTypesResponseLogs {
		if v == nil {
			return nil
		}
		return v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponseLogs struct {
	State *string `pulumi:"state"`
}

type AwsCloudTrailDataConnectorDataTypesResponseLogsOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesResponseLogs) *string { return v.State }).(pulumi.StringPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput struct{ *pulumi.OutputState }

func (AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return o
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput) Elem() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypesResponseLogs) AwsCloudTrailDataConnectorDataTypesResponseLogs {
		if v != nil {
			return *v
		}
		var ret AwsCloudTrailDataConnectorDataTypesResponseLogs
		return ret
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AwsCloudTrailDataConnectorDataTypesResponseLogs) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ClientInfoResponse struct {
	Email             *string `pulumi:"email"`
	Name              *string `pulumi:"name"`
	ObjectId          *string `pulumi:"objectId"`
	UserPrincipalName *string `pulumi:"userPrincipalName"`
}

type ClientInfoResponseOutput struct{ *pulumi.OutputState }

func (ClientInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientInfoResponse)(nil)).Elem()
}

func (o ClientInfoResponseOutput) ToClientInfoResponseOutput() ClientInfoResponseOutput {
	return o
}

func (o ClientInfoResponseOutput) ToClientInfoResponseOutputWithContext(ctx context.Context) ClientInfoResponseOutput {
	return o
}

func (o ClientInfoResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientInfoResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o ClientInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClientInfoResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientInfoResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o ClientInfoResponseOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClientInfoResponse) *string { return v.UserPrincipalName }).(pulumi.StringPtrOutput)
}

type DataConnectorDataTypeCommon struct {
	State *string `pulumi:"state"`
}





type DataConnectorDataTypeCommonInput interface {
	pulumi.Input

	ToDataConnectorDataTypeCommonOutput() DataConnectorDataTypeCommonOutput
	ToDataConnectorDataTypeCommonOutputWithContext(context.Context) DataConnectorDataTypeCommonOutput
}

type DataConnectorDataTypeCommonArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (DataConnectorDataTypeCommonArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorDataTypeCommon)(nil)).Elem()
}

func (i DataConnectorDataTypeCommonArgs) ToDataConnectorDataTypeCommonOutput() DataConnectorDataTypeCommonOutput {
	return i.ToDataConnectorDataTypeCommonOutputWithContext(context.Background())
}

func (i DataConnectorDataTypeCommonArgs) ToDataConnectorDataTypeCommonOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorDataTypeCommonOutput)
}

func (i DataConnectorDataTypeCommonArgs) ToDataConnectorDataTypeCommonPtrOutput() DataConnectorDataTypeCommonPtrOutput {
	return i.ToDataConnectorDataTypeCommonPtrOutputWithContext(context.Background())
}

func (i DataConnectorDataTypeCommonArgs) ToDataConnectorDataTypeCommonPtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorDataTypeCommonOutput).ToDataConnectorDataTypeCommonPtrOutputWithContext(ctx)
}









type DataConnectorDataTypeCommonPtrInput interface {
	pulumi.Input

	ToDataConnectorDataTypeCommonPtrOutput() DataConnectorDataTypeCommonPtrOutput
	ToDataConnectorDataTypeCommonPtrOutputWithContext(context.Context) DataConnectorDataTypeCommonPtrOutput
}

type dataConnectorDataTypeCommonPtrType DataConnectorDataTypeCommonArgs

func DataConnectorDataTypeCommonPtr(v *DataConnectorDataTypeCommonArgs) DataConnectorDataTypeCommonPtrInput {
	return (*dataConnectorDataTypeCommonPtrType)(v)
}

func (*dataConnectorDataTypeCommonPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorDataTypeCommon)(nil)).Elem()
}

func (i *dataConnectorDataTypeCommonPtrType) ToDataConnectorDataTypeCommonPtrOutput() DataConnectorDataTypeCommonPtrOutput {
	return i.ToDataConnectorDataTypeCommonPtrOutputWithContext(context.Background())
}

func (i *dataConnectorDataTypeCommonPtrType) ToDataConnectorDataTypeCommonPtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorDataTypeCommonPtrOutput)
}

type DataConnectorDataTypeCommonOutput struct{ *pulumi.OutputState }

func (DataConnectorDataTypeCommonOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorDataTypeCommon)(nil)).Elem()
}

func (o DataConnectorDataTypeCommonOutput) ToDataConnectorDataTypeCommonOutput() DataConnectorDataTypeCommonOutput {
	return o
}

func (o DataConnectorDataTypeCommonOutput) ToDataConnectorDataTypeCommonOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonOutput {
	return o
}

func (o DataConnectorDataTypeCommonOutput) ToDataConnectorDataTypeCommonPtrOutput() DataConnectorDataTypeCommonPtrOutput {
	return o.ToDataConnectorDataTypeCommonPtrOutputWithContext(context.Background())
}

func (o DataConnectorDataTypeCommonOutput) ToDataConnectorDataTypeCommonPtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataConnectorDataTypeCommon) *DataConnectorDataTypeCommon {
		return &v
	}).(DataConnectorDataTypeCommonPtrOutput)
}

func (o DataConnectorDataTypeCommonOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataConnectorDataTypeCommon) *string { return v.State }).(pulumi.StringPtrOutput)
}

type DataConnectorDataTypeCommonPtrOutput struct{ *pulumi.OutputState }

func (DataConnectorDataTypeCommonPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorDataTypeCommon)(nil)).Elem()
}

func (o DataConnectorDataTypeCommonPtrOutput) ToDataConnectorDataTypeCommonPtrOutput() DataConnectorDataTypeCommonPtrOutput {
	return o
}

func (o DataConnectorDataTypeCommonPtrOutput) ToDataConnectorDataTypeCommonPtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonPtrOutput {
	return o
}

func (o DataConnectorDataTypeCommonPtrOutput) Elem() DataConnectorDataTypeCommonOutput {
	return o.ApplyT(func(v *DataConnectorDataTypeCommon) DataConnectorDataTypeCommon {
		if v != nil {
			return *v
		}
		var ret DataConnectorDataTypeCommon
		return ret
	}).(DataConnectorDataTypeCommonOutput)
}

func (o DataConnectorDataTypeCommonPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataConnectorDataTypeCommon) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type DataConnectorDataTypeCommonResponse struct {
	State *string `pulumi:"state"`
}

type DataConnectorDataTypeCommonResponseOutput struct{ *pulumi.OutputState }

func (DataConnectorDataTypeCommonResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorDataTypeCommonResponse)(nil)).Elem()
}

func (o DataConnectorDataTypeCommonResponseOutput) ToDataConnectorDataTypeCommonResponseOutput() DataConnectorDataTypeCommonResponseOutput {
	return o
}

func (o DataConnectorDataTypeCommonResponseOutput) ToDataConnectorDataTypeCommonResponseOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonResponseOutput {
	return o
}

func (o DataConnectorDataTypeCommonResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataConnectorDataTypeCommonResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type DataConnectorDataTypeCommonResponsePtrOutput struct{ *pulumi.OutputState }

func (DataConnectorDataTypeCommonResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorDataTypeCommonResponse)(nil)).Elem()
}

func (o DataConnectorDataTypeCommonResponsePtrOutput) ToDataConnectorDataTypeCommonResponsePtrOutput() DataConnectorDataTypeCommonResponsePtrOutput {
	return o
}

func (o DataConnectorDataTypeCommonResponsePtrOutput) ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonResponsePtrOutput {
	return o
}

func (o DataConnectorDataTypeCommonResponsePtrOutput) Elem() DataConnectorDataTypeCommonResponseOutput {
	return o.ApplyT(func(v *DataConnectorDataTypeCommonResponse) DataConnectorDataTypeCommonResponse {
		if v != nil {
			return *v
		}
		var ret DataConnectorDataTypeCommonResponse
		return ret
	}).(DataConnectorDataTypeCommonResponseOutput)
}

func (o DataConnectorDataTypeCommonResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataConnectorDataTypeCommonResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type EntityMapping struct {
	EntityType    *string        `pulumi:"entityType"`
	FieldMappings []FieldMapping `pulumi:"fieldMappings"`
}





type EntityMappingInput interface {
	pulumi.Input

	ToEntityMappingOutput() EntityMappingOutput
	ToEntityMappingOutputWithContext(context.Context) EntityMappingOutput
}

type EntityMappingArgs struct {
	EntityType    pulumi.StringPtrInput  `pulumi:"entityType"`
	FieldMappings FieldMappingArrayInput `pulumi:"fieldMappings"`
}

func (EntityMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityMapping)(nil)).Elem()
}

func (i EntityMappingArgs) ToEntityMappingOutput() EntityMappingOutput {
	return i.ToEntityMappingOutputWithContext(context.Background())
}

func (i EntityMappingArgs) ToEntityMappingOutputWithContext(ctx context.Context) EntityMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityMappingOutput)
}





type EntityMappingArrayInput interface {
	pulumi.Input

	ToEntityMappingArrayOutput() EntityMappingArrayOutput
	ToEntityMappingArrayOutputWithContext(context.Context) EntityMappingArrayOutput
}

type EntityMappingArray []EntityMappingInput

func (EntityMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityMapping)(nil)).Elem()
}

func (i EntityMappingArray) ToEntityMappingArrayOutput() EntityMappingArrayOutput {
	return i.ToEntityMappingArrayOutputWithContext(context.Background())
}

func (i EntityMappingArray) ToEntityMappingArrayOutputWithContext(ctx context.Context) EntityMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityMappingArrayOutput)
}

type EntityMappingOutput struct{ *pulumi.OutputState }

func (EntityMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityMapping)(nil)).Elem()
}

func (o EntityMappingOutput) ToEntityMappingOutput() EntityMappingOutput {
	return o
}

func (o EntityMappingOutput) ToEntityMappingOutputWithContext(ctx context.Context) EntityMappingOutput {
	return o
}

func (o EntityMappingOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityMapping) *string { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o EntityMappingOutput) FieldMappings() FieldMappingArrayOutput {
	return o.ApplyT(func(v EntityMapping) []FieldMapping { return v.FieldMappings }).(FieldMappingArrayOutput)
}

type EntityMappingArrayOutput struct{ *pulumi.OutputState }

func (EntityMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityMapping)(nil)).Elem()
}

func (o EntityMappingArrayOutput) ToEntityMappingArrayOutput() EntityMappingArrayOutput {
	return o
}

func (o EntityMappingArrayOutput) ToEntityMappingArrayOutputWithContext(ctx context.Context) EntityMappingArrayOutput {
	return o
}

func (o EntityMappingArrayOutput) Index(i pulumi.IntInput) EntityMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityMapping {
		return vs[0].([]EntityMapping)[vs[1].(int)]
	}).(EntityMappingOutput)
}

type EntityMappingResponse struct {
	EntityType    *string                `pulumi:"entityType"`
	FieldMappings []FieldMappingResponse `pulumi:"fieldMappings"`
}

type EntityMappingResponseOutput struct{ *pulumi.OutputState }

func (EntityMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityMappingResponse)(nil)).Elem()
}

func (o EntityMappingResponseOutput) ToEntityMappingResponseOutput() EntityMappingResponseOutput {
	return o
}

func (o EntityMappingResponseOutput) ToEntityMappingResponseOutputWithContext(ctx context.Context) EntityMappingResponseOutput {
	return o
}

func (o EntityMappingResponseOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityMappingResponse) *string { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o EntityMappingResponseOutput) FieldMappings() FieldMappingResponseArrayOutput {
	return o.ApplyT(func(v EntityMappingResponse) []FieldMappingResponse { return v.FieldMappings }).(FieldMappingResponseArrayOutput)
}

type EntityMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (EntityMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityMappingResponse)(nil)).Elem()
}

func (o EntityMappingResponseArrayOutput) ToEntityMappingResponseArrayOutput() EntityMappingResponseArrayOutput {
	return o
}

func (o EntityMappingResponseArrayOutput) ToEntityMappingResponseArrayOutputWithContext(ctx context.Context) EntityMappingResponseArrayOutput {
	return o
}

func (o EntityMappingResponseArrayOutput) Index(i pulumi.IntInput) EntityMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityMappingResponse {
		return vs[0].([]EntityMappingResponse)[vs[1].(int)]
	}).(EntityMappingResponseOutput)
}

type EventGroupingSettings struct {
	AggregationKind *string `pulumi:"aggregationKind"`
}





type EventGroupingSettingsInput interface {
	pulumi.Input

	ToEventGroupingSettingsOutput() EventGroupingSettingsOutput
	ToEventGroupingSettingsOutputWithContext(context.Context) EventGroupingSettingsOutput
}

type EventGroupingSettingsArgs struct {
	AggregationKind pulumi.StringPtrInput `pulumi:"aggregationKind"`
}

func (EventGroupingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGroupingSettings)(nil)).Elem()
}

func (i EventGroupingSettingsArgs) ToEventGroupingSettingsOutput() EventGroupingSettingsOutput {
	return i.ToEventGroupingSettingsOutputWithContext(context.Background())
}

func (i EventGroupingSettingsArgs) ToEventGroupingSettingsOutputWithContext(ctx context.Context) EventGroupingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGroupingSettingsOutput)
}

func (i EventGroupingSettingsArgs) ToEventGroupingSettingsPtrOutput() EventGroupingSettingsPtrOutput {
	return i.ToEventGroupingSettingsPtrOutputWithContext(context.Background())
}

func (i EventGroupingSettingsArgs) ToEventGroupingSettingsPtrOutputWithContext(ctx context.Context) EventGroupingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGroupingSettingsOutput).ToEventGroupingSettingsPtrOutputWithContext(ctx)
}









type EventGroupingSettingsPtrInput interface {
	pulumi.Input

	ToEventGroupingSettingsPtrOutput() EventGroupingSettingsPtrOutput
	ToEventGroupingSettingsPtrOutputWithContext(context.Context) EventGroupingSettingsPtrOutput
}

type eventGroupingSettingsPtrType EventGroupingSettingsArgs

func EventGroupingSettingsPtr(v *EventGroupingSettingsArgs) EventGroupingSettingsPtrInput {
	return (*eventGroupingSettingsPtrType)(v)
}

func (*eventGroupingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGroupingSettings)(nil)).Elem()
}

func (i *eventGroupingSettingsPtrType) ToEventGroupingSettingsPtrOutput() EventGroupingSettingsPtrOutput {
	return i.ToEventGroupingSettingsPtrOutputWithContext(context.Background())
}

func (i *eventGroupingSettingsPtrType) ToEventGroupingSettingsPtrOutputWithContext(ctx context.Context) EventGroupingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGroupingSettingsPtrOutput)
}

type EventGroupingSettingsOutput struct{ *pulumi.OutputState }

func (EventGroupingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGroupingSettings)(nil)).Elem()
}

func (o EventGroupingSettingsOutput) ToEventGroupingSettingsOutput() EventGroupingSettingsOutput {
	return o
}

func (o EventGroupingSettingsOutput) ToEventGroupingSettingsOutputWithContext(ctx context.Context) EventGroupingSettingsOutput {
	return o
}

func (o EventGroupingSettingsOutput) ToEventGroupingSettingsPtrOutput() EventGroupingSettingsPtrOutput {
	return o.ToEventGroupingSettingsPtrOutputWithContext(context.Background())
}

func (o EventGroupingSettingsOutput) ToEventGroupingSettingsPtrOutputWithContext(ctx context.Context) EventGroupingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventGroupingSettings) *EventGroupingSettings {
		return &v
	}).(EventGroupingSettingsPtrOutput)
}

func (o EventGroupingSettingsOutput) AggregationKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGroupingSettings) *string { return v.AggregationKind }).(pulumi.StringPtrOutput)
}

type EventGroupingSettingsPtrOutput struct{ *pulumi.OutputState }

func (EventGroupingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGroupingSettings)(nil)).Elem()
}

func (o EventGroupingSettingsPtrOutput) ToEventGroupingSettingsPtrOutput() EventGroupingSettingsPtrOutput {
	return o
}

func (o EventGroupingSettingsPtrOutput) ToEventGroupingSettingsPtrOutputWithContext(ctx context.Context) EventGroupingSettingsPtrOutput {
	return o
}

func (o EventGroupingSettingsPtrOutput) Elem() EventGroupingSettingsOutput {
	return o.ApplyT(func(v *EventGroupingSettings) EventGroupingSettings {
		if v != nil {
			return *v
		}
		var ret EventGroupingSettings
		return ret
	}).(EventGroupingSettingsOutput)
}

func (o EventGroupingSettingsPtrOutput) AggregationKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGroupingSettings) *string {
		if v == nil {
			return nil
		}
		return v.AggregationKind
	}).(pulumi.StringPtrOutput)
}

type EventGroupingSettingsResponse struct {
	AggregationKind *string `pulumi:"aggregationKind"`
}

type EventGroupingSettingsResponseOutput struct{ *pulumi.OutputState }

func (EventGroupingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGroupingSettingsResponse)(nil)).Elem()
}

func (o EventGroupingSettingsResponseOutput) ToEventGroupingSettingsResponseOutput() EventGroupingSettingsResponseOutput {
	return o
}

func (o EventGroupingSettingsResponseOutput) ToEventGroupingSettingsResponseOutputWithContext(ctx context.Context) EventGroupingSettingsResponseOutput {
	return o
}

func (o EventGroupingSettingsResponseOutput) AggregationKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventGroupingSettingsResponse) *string { return v.AggregationKind }).(pulumi.StringPtrOutput)
}

type EventGroupingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (EventGroupingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGroupingSettingsResponse)(nil)).Elem()
}

func (o EventGroupingSettingsResponsePtrOutput) ToEventGroupingSettingsResponsePtrOutput() EventGroupingSettingsResponsePtrOutput {
	return o
}

func (o EventGroupingSettingsResponsePtrOutput) ToEventGroupingSettingsResponsePtrOutputWithContext(ctx context.Context) EventGroupingSettingsResponsePtrOutput {
	return o
}

func (o EventGroupingSettingsResponsePtrOutput) Elem() EventGroupingSettingsResponseOutput {
	return o.ApplyT(func(v *EventGroupingSettingsResponse) EventGroupingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret EventGroupingSettingsResponse
		return ret
	}).(EventGroupingSettingsResponseOutput)
}

func (o EventGroupingSettingsResponsePtrOutput) AggregationKind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventGroupingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AggregationKind
	}).(pulumi.StringPtrOutput)
}

type FieldMapping struct {
	ColumnName *string `pulumi:"columnName"`
	Identifier *string `pulumi:"identifier"`
}





type FieldMappingInput interface {
	pulumi.Input

	ToFieldMappingOutput() FieldMappingOutput
	ToFieldMappingOutputWithContext(context.Context) FieldMappingOutput
}

type FieldMappingArgs struct {
	ColumnName pulumi.StringPtrInput `pulumi:"columnName"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (FieldMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldMapping)(nil)).Elem()
}

func (i FieldMappingArgs) ToFieldMappingOutput() FieldMappingOutput {
	return i.ToFieldMappingOutputWithContext(context.Background())
}

func (i FieldMappingArgs) ToFieldMappingOutputWithContext(ctx context.Context) FieldMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldMappingOutput)
}





type FieldMappingArrayInput interface {
	pulumi.Input

	ToFieldMappingArrayOutput() FieldMappingArrayOutput
	ToFieldMappingArrayOutputWithContext(context.Context) FieldMappingArrayOutput
}

type FieldMappingArray []FieldMappingInput

func (FieldMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FieldMapping)(nil)).Elem()
}

func (i FieldMappingArray) ToFieldMappingArrayOutput() FieldMappingArrayOutput {
	return i.ToFieldMappingArrayOutputWithContext(context.Background())
}

func (i FieldMappingArray) ToFieldMappingArrayOutputWithContext(ctx context.Context) FieldMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldMappingArrayOutput)
}

type FieldMappingOutput struct{ *pulumi.OutputState }

func (FieldMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldMapping)(nil)).Elem()
}

func (o FieldMappingOutput) ToFieldMappingOutput() FieldMappingOutput {
	return o
}

func (o FieldMappingOutput) ToFieldMappingOutputWithContext(ctx context.Context) FieldMappingOutput {
	return o
}

func (o FieldMappingOutput) ColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FieldMapping) *string { return v.ColumnName }).(pulumi.StringPtrOutput)
}

func (o FieldMappingOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FieldMapping) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type FieldMappingArrayOutput struct{ *pulumi.OutputState }

func (FieldMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FieldMapping)(nil)).Elem()
}

func (o FieldMappingArrayOutput) ToFieldMappingArrayOutput() FieldMappingArrayOutput {
	return o
}

func (o FieldMappingArrayOutput) ToFieldMappingArrayOutputWithContext(ctx context.Context) FieldMappingArrayOutput {
	return o
}

func (o FieldMappingArrayOutput) Index(i pulumi.IntInput) FieldMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FieldMapping {
		return vs[0].([]FieldMapping)[vs[1].(int)]
	}).(FieldMappingOutput)
}

type FieldMappingResponse struct {
	ColumnName *string `pulumi:"columnName"`
	Identifier *string `pulumi:"identifier"`
}

type FieldMappingResponseOutput struct{ *pulumi.OutputState }

func (FieldMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldMappingResponse)(nil)).Elem()
}

func (o FieldMappingResponseOutput) ToFieldMappingResponseOutput() FieldMappingResponseOutput {
	return o
}

func (o FieldMappingResponseOutput) ToFieldMappingResponseOutputWithContext(ctx context.Context) FieldMappingResponseOutput {
	return o
}

func (o FieldMappingResponseOutput) ColumnName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FieldMappingResponse) *string { return v.ColumnName }).(pulumi.StringPtrOutput)
}

func (o FieldMappingResponseOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FieldMappingResponse) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

type FieldMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (FieldMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FieldMappingResponse)(nil)).Elem()
}

func (o FieldMappingResponseArrayOutput) ToFieldMappingResponseArrayOutput() FieldMappingResponseArrayOutput {
	return o
}

func (o FieldMappingResponseArrayOutput) ToFieldMappingResponseArrayOutputWithContext(ctx context.Context) FieldMappingResponseArrayOutput {
	return o
}

func (o FieldMappingResponseArrayOutput) Index(i pulumi.IntInput) FieldMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FieldMappingResponse {
		return vs[0].([]FieldMappingResponse)[vs[1].(int)]
	}).(FieldMappingResponseOutput)
}

type GroupingConfiguration struct {
	Enabled              bool     `pulumi:"enabled"`
	GroupByAlertDetails  []string `pulumi:"groupByAlertDetails"`
	GroupByCustomDetails []string `pulumi:"groupByCustomDetails"`
	GroupByEntities      []string `pulumi:"groupByEntities"`
	LookbackDuration     string   `pulumi:"lookbackDuration"`
	MatchingMethod       string   `pulumi:"matchingMethod"`
	ReopenClosedIncident bool     `pulumi:"reopenClosedIncident"`
}





type GroupingConfigurationInput interface {
	pulumi.Input

	ToGroupingConfigurationOutput() GroupingConfigurationOutput
	ToGroupingConfigurationOutputWithContext(context.Context) GroupingConfigurationOutput
}

type GroupingConfigurationArgs struct {
	Enabled              pulumi.BoolInput        `pulumi:"enabled"`
	GroupByAlertDetails  pulumi.StringArrayInput `pulumi:"groupByAlertDetails"`
	GroupByCustomDetails pulumi.StringArrayInput `pulumi:"groupByCustomDetails"`
	GroupByEntities      pulumi.StringArrayInput `pulumi:"groupByEntities"`
	LookbackDuration     pulumi.StringInput      `pulumi:"lookbackDuration"`
	MatchingMethod       pulumi.StringInput      `pulumi:"matchingMethod"`
	ReopenClosedIncident pulumi.BoolInput        `pulumi:"reopenClosedIncident"`
}

func (GroupingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingConfiguration)(nil)).Elem()
}

func (i GroupingConfigurationArgs) ToGroupingConfigurationOutput() GroupingConfigurationOutput {
	return i.ToGroupingConfigurationOutputWithContext(context.Background())
}

func (i GroupingConfigurationArgs) ToGroupingConfigurationOutputWithContext(ctx context.Context) GroupingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingConfigurationOutput)
}

func (i GroupingConfigurationArgs) ToGroupingConfigurationPtrOutput() GroupingConfigurationPtrOutput {
	return i.ToGroupingConfigurationPtrOutputWithContext(context.Background())
}

func (i GroupingConfigurationArgs) ToGroupingConfigurationPtrOutputWithContext(ctx context.Context) GroupingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingConfigurationOutput).ToGroupingConfigurationPtrOutputWithContext(ctx)
}









type GroupingConfigurationPtrInput interface {
	pulumi.Input

	ToGroupingConfigurationPtrOutput() GroupingConfigurationPtrOutput
	ToGroupingConfigurationPtrOutputWithContext(context.Context) GroupingConfigurationPtrOutput
}

type groupingConfigurationPtrType GroupingConfigurationArgs

func GroupingConfigurationPtr(v *GroupingConfigurationArgs) GroupingConfigurationPtrInput {
	return (*groupingConfigurationPtrType)(v)
}

func (*groupingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingConfiguration)(nil)).Elem()
}

func (i *groupingConfigurationPtrType) ToGroupingConfigurationPtrOutput() GroupingConfigurationPtrOutput {
	return i.ToGroupingConfigurationPtrOutputWithContext(context.Background())
}

func (i *groupingConfigurationPtrType) ToGroupingConfigurationPtrOutputWithContext(ctx context.Context) GroupingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingConfigurationPtrOutput)
}

type GroupingConfigurationOutput struct{ *pulumi.OutputState }

func (GroupingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingConfiguration)(nil)).Elem()
}

func (o GroupingConfigurationOutput) ToGroupingConfigurationOutput() GroupingConfigurationOutput {
	return o
}

func (o GroupingConfigurationOutput) ToGroupingConfigurationOutputWithContext(ctx context.Context) GroupingConfigurationOutput {
	return o
}

func (o GroupingConfigurationOutput) ToGroupingConfigurationPtrOutput() GroupingConfigurationPtrOutput {
	return o.ToGroupingConfigurationPtrOutputWithContext(context.Background())
}

func (o GroupingConfigurationOutput) ToGroupingConfigurationPtrOutputWithContext(ctx context.Context) GroupingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupingConfiguration) *GroupingConfiguration {
		return &v
	}).(GroupingConfigurationPtrOutput)
}

func (o GroupingConfigurationOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupingConfiguration) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GroupingConfigurationOutput) GroupByAlertDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfiguration) []string { return v.GroupByAlertDetails }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationOutput) GroupByCustomDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfiguration) []string { return v.GroupByCustomDetails }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationOutput) GroupByEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfiguration) []string { return v.GroupByEntities }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationOutput) LookbackDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfiguration) string { return v.LookbackDuration }).(pulumi.StringOutput)
}

func (o GroupingConfigurationOutput) MatchingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfiguration) string { return v.MatchingMethod }).(pulumi.StringOutput)
}

func (o GroupingConfigurationOutput) ReopenClosedIncident() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupingConfiguration) bool { return v.ReopenClosedIncident }).(pulumi.BoolOutput)
}

type GroupingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GroupingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingConfiguration)(nil)).Elem()
}

func (o GroupingConfigurationPtrOutput) ToGroupingConfigurationPtrOutput() GroupingConfigurationPtrOutput {
	return o
}

func (o GroupingConfigurationPtrOutput) ToGroupingConfigurationPtrOutputWithContext(ctx context.Context) GroupingConfigurationPtrOutput {
	return o
}

func (o GroupingConfigurationPtrOutput) Elem() GroupingConfigurationOutput {
	return o.ApplyT(func(v *GroupingConfiguration) GroupingConfiguration {
		if v != nil {
			return *v
		}
		var ret GroupingConfiguration
		return ret
	}).(GroupingConfigurationOutput)
}

func (o GroupingConfigurationPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GroupingConfigurationPtrOutput) GroupByAlertDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupingConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.GroupByAlertDetails
	}).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationPtrOutput) GroupByCustomDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupingConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.GroupByCustomDetails
	}).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationPtrOutput) GroupByEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupingConfiguration) []string {
		if v == nil {
			return nil
		}
		return v.GroupByEntities
	}).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationPtrOutput) LookbackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.LookbackDuration
	}).(pulumi.StringPtrOutput)
}

func (o GroupingConfigurationPtrOutput) MatchingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.MatchingMethod
	}).(pulumi.StringPtrOutput)
}

func (o GroupingConfigurationPtrOutput) ReopenClosedIncident() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.ReopenClosedIncident
	}).(pulumi.BoolPtrOutput)
}

type GroupingConfigurationResponse struct {
	Enabled              bool     `pulumi:"enabled"`
	GroupByAlertDetails  []string `pulumi:"groupByAlertDetails"`
	GroupByCustomDetails []string `pulumi:"groupByCustomDetails"`
	GroupByEntities      []string `pulumi:"groupByEntities"`
	LookbackDuration     string   `pulumi:"lookbackDuration"`
	MatchingMethod       string   `pulumi:"matchingMethod"`
	ReopenClosedIncident bool     `pulumi:"reopenClosedIncident"`
}

type GroupingConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GroupingConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingConfigurationResponse)(nil)).Elem()
}

func (o GroupingConfigurationResponseOutput) ToGroupingConfigurationResponseOutput() GroupingConfigurationResponseOutput {
	return o
}

func (o GroupingConfigurationResponseOutput) ToGroupingConfigurationResponseOutputWithContext(ctx context.Context) GroupingConfigurationResponseOutput {
	return o
}

func (o GroupingConfigurationResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o GroupingConfigurationResponseOutput) GroupByAlertDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) []string { return v.GroupByAlertDetails }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponseOutput) GroupByCustomDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) []string { return v.GroupByCustomDetails }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponseOutput) GroupByEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) []string { return v.GroupByEntities }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponseOutput) LookbackDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) string { return v.LookbackDuration }).(pulumi.StringOutput)
}

func (o GroupingConfigurationResponseOutput) MatchingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) string { return v.MatchingMethod }).(pulumi.StringOutput)
}

func (o GroupingConfigurationResponseOutput) ReopenClosedIncident() pulumi.BoolOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) bool { return v.ReopenClosedIncident }).(pulumi.BoolOutput)
}

type GroupingConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GroupingConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingConfigurationResponse)(nil)).Elem()
}

func (o GroupingConfigurationResponsePtrOutput) ToGroupingConfigurationResponsePtrOutput() GroupingConfigurationResponsePtrOutput {
	return o
}

func (o GroupingConfigurationResponsePtrOutput) ToGroupingConfigurationResponsePtrOutputWithContext(ctx context.Context) GroupingConfigurationResponsePtrOutput {
	return o
}

func (o GroupingConfigurationResponsePtrOutput) Elem() GroupingConfigurationResponseOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) GroupingConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GroupingConfigurationResponse
		return ret
	}).(GroupingConfigurationResponseOutput)
}

func (o GroupingConfigurationResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o GroupingConfigurationResponsePtrOutput) GroupByAlertDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupByAlertDetails
	}).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponsePtrOutput) GroupByCustomDetails() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupByCustomDetails
	}).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponsePtrOutput) GroupByEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.GroupByEntities
	}).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponsePtrOutput) LookbackDuration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LookbackDuration
	}).(pulumi.StringPtrOutput)
}

func (o GroupingConfigurationResponsePtrOutput) MatchingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MatchingMethod
	}).(pulumi.StringPtrOutput)
}

func (o GroupingConfigurationResponsePtrOutput) ReopenClosedIncident() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.ReopenClosedIncident
	}).(pulumi.BoolPtrOutput)
}

type IncidentAdditionalDataResponse struct {
	AlertProductNames []string `pulumi:"alertProductNames"`
	AlertsCount       int      `pulumi:"alertsCount"`
	BookmarksCount    int      `pulumi:"bookmarksCount"`
	CommentsCount     int      `pulumi:"commentsCount"`
	Tactics           []string `pulumi:"tactics"`
}

type IncidentAdditionalDataResponseOutput struct{ *pulumi.OutputState }

func (IncidentAdditionalDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentAdditionalDataResponse)(nil)).Elem()
}

func (o IncidentAdditionalDataResponseOutput) ToIncidentAdditionalDataResponseOutput() IncidentAdditionalDataResponseOutput {
	return o
}

func (o IncidentAdditionalDataResponseOutput) ToIncidentAdditionalDataResponseOutputWithContext(ctx context.Context) IncidentAdditionalDataResponseOutput {
	return o
}

func (o IncidentAdditionalDataResponseOutput) AlertProductNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) []string { return v.AlertProductNames }).(pulumi.StringArrayOutput)
}

func (o IncidentAdditionalDataResponseOutput) AlertsCount() pulumi.IntOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) int { return v.AlertsCount }).(pulumi.IntOutput)
}

func (o IncidentAdditionalDataResponseOutput) BookmarksCount() pulumi.IntOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) int { return v.BookmarksCount }).(pulumi.IntOutput)
}

func (o IncidentAdditionalDataResponseOutput) CommentsCount() pulumi.IntOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) int { return v.CommentsCount }).(pulumi.IntOutput)
}

func (o IncidentAdditionalDataResponseOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

type IncidentConfiguration struct {
	CreateIncident        bool                   `pulumi:"createIncident"`
	GroupingConfiguration *GroupingConfiguration `pulumi:"groupingConfiguration"`
}





type IncidentConfigurationInput interface {
	pulumi.Input

	ToIncidentConfigurationOutput() IncidentConfigurationOutput
	ToIncidentConfigurationOutputWithContext(context.Context) IncidentConfigurationOutput
}

type IncidentConfigurationArgs struct {
	CreateIncident        pulumi.BoolInput              `pulumi:"createIncident"`
	GroupingConfiguration GroupingConfigurationPtrInput `pulumi:"groupingConfiguration"`
}

func (IncidentConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentConfiguration)(nil)).Elem()
}

func (i IncidentConfigurationArgs) ToIncidentConfigurationOutput() IncidentConfigurationOutput {
	return i.ToIncidentConfigurationOutputWithContext(context.Background())
}

func (i IncidentConfigurationArgs) ToIncidentConfigurationOutputWithContext(ctx context.Context) IncidentConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentConfigurationOutput)
}

func (i IncidentConfigurationArgs) ToIncidentConfigurationPtrOutput() IncidentConfigurationPtrOutput {
	return i.ToIncidentConfigurationPtrOutputWithContext(context.Background())
}

func (i IncidentConfigurationArgs) ToIncidentConfigurationPtrOutputWithContext(ctx context.Context) IncidentConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentConfigurationOutput).ToIncidentConfigurationPtrOutputWithContext(ctx)
}









type IncidentConfigurationPtrInput interface {
	pulumi.Input

	ToIncidentConfigurationPtrOutput() IncidentConfigurationPtrOutput
	ToIncidentConfigurationPtrOutputWithContext(context.Context) IncidentConfigurationPtrOutput
}

type incidentConfigurationPtrType IncidentConfigurationArgs

func IncidentConfigurationPtr(v *IncidentConfigurationArgs) IncidentConfigurationPtrInput {
	return (*incidentConfigurationPtrType)(v)
}

func (*incidentConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentConfiguration)(nil)).Elem()
}

func (i *incidentConfigurationPtrType) ToIncidentConfigurationPtrOutput() IncidentConfigurationPtrOutput {
	return i.ToIncidentConfigurationPtrOutputWithContext(context.Background())
}

func (i *incidentConfigurationPtrType) ToIncidentConfigurationPtrOutputWithContext(ctx context.Context) IncidentConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentConfigurationPtrOutput)
}

type IncidentConfigurationOutput struct{ *pulumi.OutputState }

func (IncidentConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentConfiguration)(nil)).Elem()
}

func (o IncidentConfigurationOutput) ToIncidentConfigurationOutput() IncidentConfigurationOutput {
	return o
}

func (o IncidentConfigurationOutput) ToIncidentConfigurationOutputWithContext(ctx context.Context) IncidentConfigurationOutput {
	return o
}

func (o IncidentConfigurationOutput) ToIncidentConfigurationPtrOutput() IncidentConfigurationPtrOutput {
	return o.ToIncidentConfigurationPtrOutputWithContext(context.Background())
}

func (o IncidentConfigurationOutput) ToIncidentConfigurationPtrOutputWithContext(ctx context.Context) IncidentConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentConfiguration) *IncidentConfiguration {
		return &v
	}).(IncidentConfigurationPtrOutput)
}

func (o IncidentConfigurationOutput) CreateIncident() pulumi.BoolOutput {
	return o.ApplyT(func(v IncidentConfiguration) bool { return v.CreateIncident }).(pulumi.BoolOutput)
}

func (o IncidentConfigurationOutput) GroupingConfiguration() GroupingConfigurationPtrOutput {
	return o.ApplyT(func(v IncidentConfiguration) *GroupingConfiguration { return v.GroupingConfiguration }).(GroupingConfigurationPtrOutput)
}

type IncidentConfigurationPtrOutput struct{ *pulumi.OutputState }

func (IncidentConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentConfiguration)(nil)).Elem()
}

func (o IncidentConfigurationPtrOutput) ToIncidentConfigurationPtrOutput() IncidentConfigurationPtrOutput {
	return o
}

func (o IncidentConfigurationPtrOutput) ToIncidentConfigurationPtrOutputWithContext(ctx context.Context) IncidentConfigurationPtrOutput {
	return o
}

func (o IncidentConfigurationPtrOutput) Elem() IncidentConfigurationOutput {
	return o.ApplyT(func(v *IncidentConfiguration) IncidentConfiguration {
		if v != nil {
			return *v
		}
		var ret IncidentConfiguration
		return ret
	}).(IncidentConfigurationOutput)
}

func (o IncidentConfigurationPtrOutput) CreateIncident() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IncidentConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateIncident
	}).(pulumi.BoolPtrOutput)
}

func (o IncidentConfigurationPtrOutput) GroupingConfiguration() GroupingConfigurationPtrOutput {
	return o.ApplyT(func(v *IncidentConfiguration) *GroupingConfiguration {
		if v == nil {
			return nil
		}
		return v.GroupingConfiguration
	}).(GroupingConfigurationPtrOutput)
}

type IncidentConfigurationResponse struct {
	CreateIncident        bool                           `pulumi:"createIncident"`
	GroupingConfiguration *GroupingConfigurationResponse `pulumi:"groupingConfiguration"`
}

type IncidentConfigurationResponseOutput struct{ *pulumi.OutputState }

func (IncidentConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentConfigurationResponse)(nil)).Elem()
}

func (o IncidentConfigurationResponseOutput) ToIncidentConfigurationResponseOutput() IncidentConfigurationResponseOutput {
	return o
}

func (o IncidentConfigurationResponseOutput) ToIncidentConfigurationResponseOutputWithContext(ctx context.Context) IncidentConfigurationResponseOutput {
	return o
}

func (o IncidentConfigurationResponseOutput) CreateIncident() pulumi.BoolOutput {
	return o.ApplyT(func(v IncidentConfigurationResponse) bool { return v.CreateIncident }).(pulumi.BoolOutput)
}

func (o IncidentConfigurationResponseOutput) GroupingConfiguration() GroupingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v IncidentConfigurationResponse) *GroupingConfigurationResponse { return v.GroupingConfiguration }).(GroupingConfigurationResponsePtrOutput)
}

type IncidentConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (IncidentConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentConfigurationResponse)(nil)).Elem()
}

func (o IncidentConfigurationResponsePtrOutput) ToIncidentConfigurationResponsePtrOutput() IncidentConfigurationResponsePtrOutput {
	return o
}

func (o IncidentConfigurationResponsePtrOutput) ToIncidentConfigurationResponsePtrOutputWithContext(ctx context.Context) IncidentConfigurationResponsePtrOutput {
	return o
}

func (o IncidentConfigurationResponsePtrOutput) Elem() IncidentConfigurationResponseOutput {
	return o.ApplyT(func(v *IncidentConfigurationResponse) IncidentConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret IncidentConfigurationResponse
		return ret
	}).(IncidentConfigurationResponseOutput)
}

func (o IncidentConfigurationResponsePtrOutput) CreateIncident() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IncidentConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.CreateIncident
	}).(pulumi.BoolPtrOutput)
}

func (o IncidentConfigurationResponsePtrOutput) GroupingConfiguration() GroupingConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *IncidentConfigurationResponse) *GroupingConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.GroupingConfiguration
	}).(GroupingConfigurationResponsePtrOutput)
}

type IncidentInfo struct {
	IncidentId   *string `pulumi:"incidentId"`
	RelationName *string `pulumi:"relationName"`
	Severity     *string `pulumi:"severity"`
	Title        *string `pulumi:"title"`
}





type IncidentInfoInput interface {
	pulumi.Input

	ToIncidentInfoOutput() IncidentInfoOutput
	ToIncidentInfoOutputWithContext(context.Context) IncidentInfoOutput
}

type IncidentInfoArgs struct {
	IncidentId   pulumi.StringPtrInput `pulumi:"incidentId"`
	RelationName pulumi.StringPtrInput `pulumi:"relationName"`
	Severity     pulumi.StringPtrInput `pulumi:"severity"`
	Title        pulumi.StringPtrInput `pulumi:"title"`
}

func (IncidentInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentInfo)(nil)).Elem()
}

func (i IncidentInfoArgs) ToIncidentInfoOutput() IncidentInfoOutput {
	return i.ToIncidentInfoOutputWithContext(context.Background())
}

func (i IncidentInfoArgs) ToIncidentInfoOutputWithContext(ctx context.Context) IncidentInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentInfoOutput)
}

func (i IncidentInfoArgs) ToIncidentInfoPtrOutput() IncidentInfoPtrOutput {
	return i.ToIncidentInfoPtrOutputWithContext(context.Background())
}

func (i IncidentInfoArgs) ToIncidentInfoPtrOutputWithContext(ctx context.Context) IncidentInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentInfoOutput).ToIncidentInfoPtrOutputWithContext(ctx)
}









type IncidentInfoPtrInput interface {
	pulumi.Input

	ToIncidentInfoPtrOutput() IncidentInfoPtrOutput
	ToIncidentInfoPtrOutputWithContext(context.Context) IncidentInfoPtrOutput
}

type incidentInfoPtrType IncidentInfoArgs

func IncidentInfoPtr(v *IncidentInfoArgs) IncidentInfoPtrInput {
	return (*incidentInfoPtrType)(v)
}

func (*incidentInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentInfo)(nil)).Elem()
}

func (i *incidentInfoPtrType) ToIncidentInfoPtrOutput() IncidentInfoPtrOutput {
	return i.ToIncidentInfoPtrOutputWithContext(context.Background())
}

func (i *incidentInfoPtrType) ToIncidentInfoPtrOutputWithContext(ctx context.Context) IncidentInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentInfoPtrOutput)
}

type IncidentInfoOutput struct{ *pulumi.OutputState }

func (IncidentInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentInfo)(nil)).Elem()
}

func (o IncidentInfoOutput) ToIncidentInfoOutput() IncidentInfoOutput {
	return o
}

func (o IncidentInfoOutput) ToIncidentInfoOutputWithContext(ctx context.Context) IncidentInfoOutput {
	return o
}

func (o IncidentInfoOutput) ToIncidentInfoPtrOutput() IncidentInfoPtrOutput {
	return o.ToIncidentInfoPtrOutputWithContext(context.Background())
}

func (o IncidentInfoOutput) ToIncidentInfoPtrOutputWithContext(ctx context.Context) IncidentInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentInfo) *IncidentInfo {
		return &v
	}).(IncidentInfoPtrOutput)
}

func (o IncidentInfoOutput) IncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfo) *string { return v.IncidentId }).(pulumi.StringPtrOutput)
}

func (o IncidentInfoOutput) RelationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfo) *string { return v.RelationName }).(pulumi.StringPtrOutput)
}

func (o IncidentInfoOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfo) *string { return v.Severity }).(pulumi.StringPtrOutput)
}

func (o IncidentInfoOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfo) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type IncidentInfoPtrOutput struct{ *pulumi.OutputState }

func (IncidentInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentInfo)(nil)).Elem()
}

func (o IncidentInfoPtrOutput) ToIncidentInfoPtrOutput() IncidentInfoPtrOutput {
	return o
}

func (o IncidentInfoPtrOutput) ToIncidentInfoPtrOutputWithContext(ctx context.Context) IncidentInfoPtrOutput {
	return o
}

func (o IncidentInfoPtrOutput) Elem() IncidentInfoOutput {
	return o.ApplyT(func(v *IncidentInfo) IncidentInfo {
		if v != nil {
			return *v
		}
		var ret IncidentInfo
		return ret
	}).(IncidentInfoOutput)
}

func (o IncidentInfoPtrOutput) IncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfo) *string {
		if v == nil {
			return nil
		}
		return v.IncidentId
	}).(pulumi.StringPtrOutput)
}

func (o IncidentInfoPtrOutput) RelationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfo) *string {
		if v == nil {
			return nil
		}
		return v.RelationName
	}).(pulumi.StringPtrOutput)
}

func (o IncidentInfoPtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfo) *string {
		if v == nil {
			return nil
		}
		return v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o IncidentInfoPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfo) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type IncidentInfoResponse struct {
	IncidentId   *string `pulumi:"incidentId"`
	RelationName *string `pulumi:"relationName"`
	Severity     *string `pulumi:"severity"`
	Title        *string `pulumi:"title"`
}

type IncidentInfoResponseOutput struct{ *pulumi.OutputState }

func (IncidentInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentInfoResponse)(nil)).Elem()
}

func (o IncidentInfoResponseOutput) ToIncidentInfoResponseOutput() IncidentInfoResponseOutput {
	return o
}

func (o IncidentInfoResponseOutput) ToIncidentInfoResponseOutputWithContext(ctx context.Context) IncidentInfoResponseOutput {
	return o
}

func (o IncidentInfoResponseOutput) IncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfoResponse) *string { return v.IncidentId }).(pulumi.StringPtrOutput)
}

func (o IncidentInfoResponseOutput) RelationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfoResponse) *string { return v.RelationName }).(pulumi.StringPtrOutput)
}

func (o IncidentInfoResponseOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfoResponse) *string { return v.Severity }).(pulumi.StringPtrOutput)
}

func (o IncidentInfoResponseOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentInfoResponse) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type IncidentInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IncidentInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentInfoResponse)(nil)).Elem()
}

func (o IncidentInfoResponsePtrOutput) ToIncidentInfoResponsePtrOutput() IncidentInfoResponsePtrOutput {
	return o
}

func (o IncidentInfoResponsePtrOutput) ToIncidentInfoResponsePtrOutputWithContext(ctx context.Context) IncidentInfoResponsePtrOutput {
	return o
}

func (o IncidentInfoResponsePtrOutput) Elem() IncidentInfoResponseOutput {
	return o.ApplyT(func(v *IncidentInfoResponse) IncidentInfoResponse {
		if v != nil {
			return *v
		}
		var ret IncidentInfoResponse
		return ret
	}).(IncidentInfoResponseOutput)
}

func (o IncidentInfoResponsePtrOutput) IncidentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.IncidentId
	}).(pulumi.StringPtrOutput)
}

func (o IncidentInfoResponsePtrOutput) RelationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.RelationName
	}).(pulumi.StringPtrOutput)
}

func (o IncidentInfoResponsePtrOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Severity
	}).(pulumi.StringPtrOutput)
}

func (o IncidentInfoResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Title
	}).(pulumi.StringPtrOutput)
}

type IncidentLabel struct {
	LabelName string `pulumi:"labelName"`
}





type IncidentLabelInput interface {
	pulumi.Input

	ToIncidentLabelOutput() IncidentLabelOutput
	ToIncidentLabelOutputWithContext(context.Context) IncidentLabelOutput
}

type IncidentLabelArgs struct {
	LabelName pulumi.StringInput `pulumi:"labelName"`
}

func (IncidentLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentLabel)(nil)).Elem()
}

func (i IncidentLabelArgs) ToIncidentLabelOutput() IncidentLabelOutput {
	return i.ToIncidentLabelOutputWithContext(context.Background())
}

func (i IncidentLabelArgs) ToIncidentLabelOutputWithContext(ctx context.Context) IncidentLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentLabelOutput)
}





type IncidentLabelArrayInput interface {
	pulumi.Input

	ToIncidentLabelArrayOutput() IncidentLabelArrayOutput
	ToIncidentLabelArrayOutputWithContext(context.Context) IncidentLabelArrayOutput
}

type IncidentLabelArray []IncidentLabelInput

func (IncidentLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncidentLabel)(nil)).Elem()
}

func (i IncidentLabelArray) ToIncidentLabelArrayOutput() IncidentLabelArrayOutput {
	return i.ToIncidentLabelArrayOutputWithContext(context.Background())
}

func (i IncidentLabelArray) ToIncidentLabelArrayOutputWithContext(ctx context.Context) IncidentLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentLabelArrayOutput)
}

type IncidentLabelOutput struct{ *pulumi.OutputState }

func (IncidentLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentLabel)(nil)).Elem()
}

func (o IncidentLabelOutput) ToIncidentLabelOutput() IncidentLabelOutput {
	return o
}

func (o IncidentLabelOutput) ToIncidentLabelOutputWithContext(ctx context.Context) IncidentLabelOutput {
	return o
}

func (o IncidentLabelOutput) LabelName() pulumi.StringOutput {
	return o.ApplyT(func(v IncidentLabel) string { return v.LabelName }).(pulumi.StringOutput)
}

type IncidentLabelArrayOutput struct{ *pulumi.OutputState }

func (IncidentLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncidentLabel)(nil)).Elem()
}

func (o IncidentLabelArrayOutput) ToIncidentLabelArrayOutput() IncidentLabelArrayOutput {
	return o
}

func (o IncidentLabelArrayOutput) ToIncidentLabelArrayOutputWithContext(ctx context.Context) IncidentLabelArrayOutput {
	return o
}

func (o IncidentLabelArrayOutput) Index(i pulumi.IntInput) IncidentLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncidentLabel {
		return vs[0].([]IncidentLabel)[vs[1].(int)]
	}).(IncidentLabelOutput)
}

type IncidentLabelResponse struct {
	LabelName string `pulumi:"labelName"`
	LabelType string `pulumi:"labelType"`
}

type IncidentLabelResponseOutput struct{ *pulumi.OutputState }

func (IncidentLabelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentLabelResponse)(nil)).Elem()
}

func (o IncidentLabelResponseOutput) ToIncidentLabelResponseOutput() IncidentLabelResponseOutput {
	return o
}

func (o IncidentLabelResponseOutput) ToIncidentLabelResponseOutputWithContext(ctx context.Context) IncidentLabelResponseOutput {
	return o
}

func (o IncidentLabelResponseOutput) LabelName() pulumi.StringOutput {
	return o.ApplyT(func(v IncidentLabelResponse) string { return v.LabelName }).(pulumi.StringOutput)
}

func (o IncidentLabelResponseOutput) LabelType() pulumi.StringOutput {
	return o.ApplyT(func(v IncidentLabelResponse) string { return v.LabelType }).(pulumi.StringOutput)
}

type IncidentLabelResponseArrayOutput struct{ *pulumi.OutputState }

func (IncidentLabelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncidentLabelResponse)(nil)).Elem()
}

func (o IncidentLabelResponseArrayOutput) ToIncidentLabelResponseArrayOutput() IncidentLabelResponseArrayOutput {
	return o
}

func (o IncidentLabelResponseArrayOutput) ToIncidentLabelResponseArrayOutputWithContext(ctx context.Context) IncidentLabelResponseArrayOutput {
	return o
}

func (o IncidentLabelResponseArrayOutput) Index(i pulumi.IntInput) IncidentLabelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncidentLabelResponse {
		return vs[0].([]IncidentLabelResponse)[vs[1].(int)]
	}).(IncidentLabelResponseOutput)
}

type IncidentOwnerInfo struct {
	AssignedTo        *string `pulumi:"assignedTo"`
	Email             *string `pulumi:"email"`
	ObjectId          *string `pulumi:"objectId"`
	OwnerType         *string `pulumi:"ownerType"`
	UserPrincipalName *string `pulumi:"userPrincipalName"`
}





type IncidentOwnerInfoInput interface {
	pulumi.Input

	ToIncidentOwnerInfoOutput() IncidentOwnerInfoOutput
	ToIncidentOwnerInfoOutputWithContext(context.Context) IncidentOwnerInfoOutput
}

type IncidentOwnerInfoArgs struct {
	AssignedTo        pulumi.StringPtrInput `pulumi:"assignedTo"`
	Email             pulumi.StringPtrInput `pulumi:"email"`
	ObjectId          pulumi.StringPtrInput `pulumi:"objectId"`
	OwnerType         pulumi.StringPtrInput `pulumi:"ownerType"`
	UserPrincipalName pulumi.StringPtrInput `pulumi:"userPrincipalName"`
}

func (IncidentOwnerInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentOwnerInfo)(nil)).Elem()
}

func (i IncidentOwnerInfoArgs) ToIncidentOwnerInfoOutput() IncidentOwnerInfoOutput {
	return i.ToIncidentOwnerInfoOutputWithContext(context.Background())
}

func (i IncidentOwnerInfoArgs) ToIncidentOwnerInfoOutputWithContext(ctx context.Context) IncidentOwnerInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOwnerInfoOutput)
}

func (i IncidentOwnerInfoArgs) ToIncidentOwnerInfoPtrOutput() IncidentOwnerInfoPtrOutput {
	return i.ToIncidentOwnerInfoPtrOutputWithContext(context.Background())
}

func (i IncidentOwnerInfoArgs) ToIncidentOwnerInfoPtrOutputWithContext(ctx context.Context) IncidentOwnerInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOwnerInfoOutput).ToIncidentOwnerInfoPtrOutputWithContext(ctx)
}









type IncidentOwnerInfoPtrInput interface {
	pulumi.Input

	ToIncidentOwnerInfoPtrOutput() IncidentOwnerInfoPtrOutput
	ToIncidentOwnerInfoPtrOutputWithContext(context.Context) IncidentOwnerInfoPtrOutput
}

type incidentOwnerInfoPtrType IncidentOwnerInfoArgs

func IncidentOwnerInfoPtr(v *IncidentOwnerInfoArgs) IncidentOwnerInfoPtrInput {
	return (*incidentOwnerInfoPtrType)(v)
}

func (*incidentOwnerInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentOwnerInfo)(nil)).Elem()
}

func (i *incidentOwnerInfoPtrType) ToIncidentOwnerInfoPtrOutput() IncidentOwnerInfoPtrOutput {
	return i.ToIncidentOwnerInfoPtrOutputWithContext(context.Background())
}

func (i *incidentOwnerInfoPtrType) ToIncidentOwnerInfoPtrOutputWithContext(ctx context.Context) IncidentOwnerInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOwnerInfoPtrOutput)
}

type IncidentOwnerInfoOutput struct{ *pulumi.OutputState }

func (IncidentOwnerInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentOwnerInfo)(nil)).Elem()
}

func (o IncidentOwnerInfoOutput) ToIncidentOwnerInfoOutput() IncidentOwnerInfoOutput {
	return o
}

func (o IncidentOwnerInfoOutput) ToIncidentOwnerInfoOutputWithContext(ctx context.Context) IncidentOwnerInfoOutput {
	return o
}

func (o IncidentOwnerInfoOutput) ToIncidentOwnerInfoPtrOutput() IncidentOwnerInfoPtrOutput {
	return o.ToIncidentOwnerInfoPtrOutputWithContext(context.Background())
}

func (o IncidentOwnerInfoOutput) ToIncidentOwnerInfoPtrOutputWithContext(ctx context.Context) IncidentOwnerInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentOwnerInfo) *IncidentOwnerInfo {
		return &v
	}).(IncidentOwnerInfoPtrOutput)
}

func (o IncidentOwnerInfoOutput) AssignedTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfo) *string { return v.AssignedTo }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfo) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfo) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoOutput) OwnerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfo) *string { return v.OwnerType }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfo) *string { return v.UserPrincipalName }).(pulumi.StringPtrOutput)
}

type IncidentOwnerInfoPtrOutput struct{ *pulumi.OutputState }

func (IncidentOwnerInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentOwnerInfo)(nil)).Elem()
}

func (o IncidentOwnerInfoPtrOutput) ToIncidentOwnerInfoPtrOutput() IncidentOwnerInfoPtrOutput {
	return o
}

func (o IncidentOwnerInfoPtrOutput) ToIncidentOwnerInfoPtrOutputWithContext(ctx context.Context) IncidentOwnerInfoPtrOutput {
	return o
}

func (o IncidentOwnerInfoPtrOutput) Elem() IncidentOwnerInfoOutput {
	return o.ApplyT(func(v *IncidentOwnerInfo) IncidentOwnerInfo {
		if v != nil {
			return *v
		}
		var ret IncidentOwnerInfo
		return ret
	}).(IncidentOwnerInfoOutput)
}

func (o IncidentOwnerInfoPtrOutput) AssignedTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfo) *string {
		if v == nil {
			return nil
		}
		return v.AssignedTo
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfo) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfo) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoPtrOutput) OwnerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfo) *string {
		if v == nil {
			return nil
		}
		return v.OwnerType
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoPtrOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfo) *string {
		if v == nil {
			return nil
		}
		return v.UserPrincipalName
	}).(pulumi.StringPtrOutput)
}

type IncidentOwnerInfoResponse struct {
	AssignedTo        *string `pulumi:"assignedTo"`
	Email             *string `pulumi:"email"`
	ObjectId          *string `pulumi:"objectId"`
	OwnerType         *string `pulumi:"ownerType"`
	UserPrincipalName *string `pulumi:"userPrincipalName"`
}

type IncidentOwnerInfoResponseOutput struct{ *pulumi.OutputState }

func (IncidentOwnerInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentOwnerInfoResponse)(nil)).Elem()
}

func (o IncidentOwnerInfoResponseOutput) ToIncidentOwnerInfoResponseOutput() IncidentOwnerInfoResponseOutput {
	return o
}

func (o IncidentOwnerInfoResponseOutput) ToIncidentOwnerInfoResponseOutputWithContext(ctx context.Context) IncidentOwnerInfoResponseOutput {
	return o
}

func (o IncidentOwnerInfoResponseOutput) AssignedTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfoResponse) *string { return v.AssignedTo }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfoResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfoResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponseOutput) OwnerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfoResponse) *string { return v.OwnerType }).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponseOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncidentOwnerInfoResponse) *string { return v.UserPrincipalName }).(pulumi.StringPtrOutput)
}

type IncidentOwnerInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (IncidentOwnerInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentOwnerInfoResponse)(nil)).Elem()
}

func (o IncidentOwnerInfoResponsePtrOutput) ToIncidentOwnerInfoResponsePtrOutput() IncidentOwnerInfoResponsePtrOutput {
	return o
}

func (o IncidentOwnerInfoResponsePtrOutput) ToIncidentOwnerInfoResponsePtrOutputWithContext(ctx context.Context) IncidentOwnerInfoResponsePtrOutput {
	return o
}

func (o IncidentOwnerInfoResponsePtrOutput) Elem() IncidentOwnerInfoResponseOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) IncidentOwnerInfoResponse {
		if v != nil {
			return *v
		}
		var ret IncidentOwnerInfoResponse
		return ret
	}).(IncidentOwnerInfoResponseOutput)
}

func (o IncidentOwnerInfoResponsePtrOutput) AssignedTo() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssignedTo
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponsePtrOutput) OwnerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.OwnerType
	}).(pulumi.StringPtrOutput)
}

func (o IncidentOwnerInfoResponsePtrOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPrincipalName
	}).(pulumi.StringPtrOutput)
}

type IncidentPropertiesAction struct {
	Classification        *string            `pulumi:"classification"`
	ClassificationComment *string            `pulumi:"classificationComment"`
	ClassificationReason  *string            `pulumi:"classificationReason"`
	Labels                []IncidentLabel    `pulumi:"labels"`
	Owner                 *IncidentOwnerInfo `pulumi:"owner"`
	Severity              *string            `pulumi:"severity"`
	Status                *string            `pulumi:"status"`
}

type IncidentPropertiesActionResponse struct {
	Classification        *string                    `pulumi:"classification"`
	ClassificationComment *string                    `pulumi:"classificationComment"`
	ClassificationReason  *string                    `pulumi:"classificationReason"`
	Labels                []IncidentLabelResponse    `pulumi:"labels"`
	Owner                 *IncidentOwnerInfoResponse `pulumi:"owner"`
	Severity              *string                    `pulumi:"severity"`
	Status                *string                    `pulumi:"status"`
}

type MCASDataConnectorDataTypes struct {
	Alerts        *DataConnectorDataTypeCommon `pulumi:"alerts"`
	DiscoveryLogs *DataConnectorDataTypeCommon `pulumi:"discoveryLogs"`
}





type MCASDataConnectorDataTypesInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesOutput() MCASDataConnectorDataTypesOutput
	ToMCASDataConnectorDataTypesOutputWithContext(context.Context) MCASDataConnectorDataTypesOutput
}

type MCASDataConnectorDataTypesArgs struct {
	Alerts        DataConnectorDataTypeCommonPtrInput `pulumi:"alerts"`
	DiscoveryLogs DataConnectorDataTypeCommonPtrInput `pulumi:"discoveryLogs"`
}

func (MCASDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypes)(nil)).Elem()
}

func (i MCASDataConnectorDataTypesArgs) ToMCASDataConnectorDataTypesOutput() MCASDataConnectorDataTypesOutput {
	return i.ToMCASDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesArgs) ToMCASDataConnectorDataTypesOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesOutput)
}

func (i MCASDataConnectorDataTypesArgs) ToMCASDataConnectorDataTypesPtrOutput() MCASDataConnectorDataTypesPtrOutput {
	return i.ToMCASDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesArgs) ToMCASDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesOutput).ToMCASDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type MCASDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesPtrOutput() MCASDataConnectorDataTypesPtrOutput
	ToMCASDataConnectorDataTypesPtrOutputWithContext(context.Context) MCASDataConnectorDataTypesPtrOutput
}

type mcasdataConnectorDataTypesPtrType MCASDataConnectorDataTypesArgs

func MCASDataConnectorDataTypesPtr(v *MCASDataConnectorDataTypesArgs) MCASDataConnectorDataTypesPtrInput {
	return (*mcasdataConnectorDataTypesPtrType)(v)
}

func (*mcasdataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypes)(nil)).Elem()
}

func (i *mcasdataConnectorDataTypesPtrType) ToMCASDataConnectorDataTypesPtrOutput() MCASDataConnectorDataTypesPtrOutput {
	return i.ToMCASDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *mcasdataConnectorDataTypesPtrType) ToMCASDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesPtrOutput)
}

type MCASDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypes)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesOutput) ToMCASDataConnectorDataTypesOutput() MCASDataConnectorDataTypesOutput {
	return o
}

func (o MCASDataConnectorDataTypesOutput) ToMCASDataConnectorDataTypesOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesOutput {
	return o
}

func (o MCASDataConnectorDataTypesOutput) ToMCASDataConnectorDataTypesPtrOutput() MCASDataConnectorDataTypesPtrOutput {
	return o.ToMCASDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o MCASDataConnectorDataTypesOutput) ToMCASDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MCASDataConnectorDataTypes) *MCASDataConnectorDataTypes {
		return &v
	}).(MCASDataConnectorDataTypesPtrOutput)
}

func (o MCASDataConnectorDataTypesOutput) Alerts() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) *DataConnectorDataTypeCommon { return v.Alerts }).(DataConnectorDataTypeCommonPtrOutput)
}

func (o MCASDataConnectorDataTypesOutput) DiscoveryLogs() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) *DataConnectorDataTypeCommon { return v.DiscoveryLogs }).(DataConnectorDataTypeCommonPtrOutput)
}

type MCASDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypes)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesPtrOutput) ToMCASDataConnectorDataTypesPtrOutput() MCASDataConnectorDataTypesPtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesPtrOutput) ToMCASDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesPtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesPtrOutput) Elem() MCASDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypes) MCASDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret MCASDataConnectorDataTypes
		return ret
	}).(MCASDataConnectorDataTypesOutput)
}

func (o MCASDataConnectorDataTypesPtrOutput) Alerts() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypes) *DataConnectorDataTypeCommon {
		if v == nil {
			return nil
		}
		return v.Alerts
	}).(DataConnectorDataTypeCommonPtrOutput)
}

func (o MCASDataConnectorDataTypesPtrOutput) DiscoveryLogs() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypes) *DataConnectorDataTypeCommon {
		if v == nil {
			return nil
		}
		return v.DiscoveryLogs
	}).(DataConnectorDataTypeCommonPtrOutput)
}

type MCASDataConnectorDataTypesResponse struct {
	Alerts        *DataConnectorDataTypeCommonResponse `pulumi:"alerts"`
	DiscoveryLogs *DataConnectorDataTypeCommonResponse `pulumi:"discoveryLogs"`
}

type MCASDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesResponseOutput) ToMCASDataConnectorDataTypesResponseOutput() MCASDataConnectorDataTypesResponseOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponseOutput) ToMCASDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponseOutput) Alerts() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) *DataConnectorDataTypeCommonResponse { return v.Alerts }).(DataConnectorDataTypeCommonResponsePtrOutput)
}

func (o MCASDataConnectorDataTypesResponseOutput) DiscoveryLogs() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) *DataConnectorDataTypeCommonResponse {
		return v.DiscoveryLogs
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
}

type MCASDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesResponsePtrOutput) ToMCASDataConnectorDataTypesResponsePtrOutput() MCASDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponsePtrOutput) ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponsePtrOutput) Elem() MCASDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponse) MCASDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret MCASDataConnectorDataTypesResponse
		return ret
	}).(MCASDataConnectorDataTypesResponseOutput)
}

func (o MCASDataConnectorDataTypesResponsePtrOutput) Alerts() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponse) *DataConnectorDataTypeCommonResponse {
		if v == nil {
			return nil
		}
		return v.Alerts
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
}

func (o MCASDataConnectorDataTypesResponsePtrOutput) DiscoveryLogs() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponse) *DataConnectorDataTypeCommonResponse {
		if v == nil {
			return nil
		}
		return v.DiscoveryLogs
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
}

type OfficeDataConnectorDataTypes struct {
	Exchange   *OfficeDataConnectorDataTypesExchange   `pulumi:"exchange"`
	SharePoint *OfficeDataConnectorDataTypesSharePoint `pulumi:"sharePoint"`
	Teams      *OfficeDataConnectorDataTypesTeams      `pulumi:"teams"`
}





type OfficeDataConnectorDataTypesInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesOutput() OfficeDataConnectorDataTypesOutput
	ToOfficeDataConnectorDataTypesOutputWithContext(context.Context) OfficeDataConnectorDataTypesOutput
}

type OfficeDataConnectorDataTypesArgs struct {
	Exchange   OfficeDataConnectorDataTypesExchangePtrInput   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesSharePointPtrInput `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesTeamsPtrInput      `pulumi:"teams"`
}

func (OfficeDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypes)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesArgs) ToOfficeDataConnectorDataTypesOutput() OfficeDataConnectorDataTypesOutput {
	return i.ToOfficeDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesArgs) ToOfficeDataConnectorDataTypesOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesOutput)
}

func (i OfficeDataConnectorDataTypesArgs) ToOfficeDataConnectorDataTypesPtrOutput() OfficeDataConnectorDataTypesPtrOutput {
	return i.ToOfficeDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesArgs) ToOfficeDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesOutput).ToOfficeDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesPtrOutput() OfficeDataConnectorDataTypesPtrOutput
	ToOfficeDataConnectorDataTypesPtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesPtrOutput
}

type officeDataConnectorDataTypesPtrType OfficeDataConnectorDataTypesArgs

func OfficeDataConnectorDataTypesPtr(v *OfficeDataConnectorDataTypesArgs) OfficeDataConnectorDataTypesPtrInput {
	return (*officeDataConnectorDataTypesPtrType)(v)
}

func (*officeDataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypes)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesPtrType) ToOfficeDataConnectorDataTypesPtrOutput() OfficeDataConnectorDataTypesPtrOutput {
	return i.ToOfficeDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesPtrType) ToOfficeDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesPtrOutput)
}

type OfficeDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypes)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesOutput) ToOfficeDataConnectorDataTypesOutput() OfficeDataConnectorDataTypesOutput {
	return o
}

func (o OfficeDataConnectorDataTypesOutput) ToOfficeDataConnectorDataTypesOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesOutput {
	return o
}

func (o OfficeDataConnectorDataTypesOutput) ToOfficeDataConnectorDataTypesPtrOutput() OfficeDataConnectorDataTypesPtrOutput {
	return o.ToOfficeDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesOutput) ToOfficeDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypes {
		return &v
	}).(OfficeDataConnectorDataTypesPtrOutput)
}

func (o OfficeDataConnectorDataTypesOutput) Exchange() OfficeDataConnectorDataTypesExchangePtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesExchange { return v.Exchange }).(OfficeDataConnectorDataTypesExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesOutput) SharePoint() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesSharePoint { return v.SharePoint }).(OfficeDataConnectorDataTypesSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesOutput) Teams() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesTeams { return v.Teams }).(OfficeDataConnectorDataTypesTeamsPtrOutput)
}

type OfficeDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypes)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesPtrOutput) ToOfficeDataConnectorDataTypesPtrOutput() OfficeDataConnectorDataTypesPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesPtrOutput) ToOfficeDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesPtrOutput) Elem() OfficeDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypes
		return ret
	}).(OfficeDataConnectorDataTypesOutput)
}

func (o OfficeDataConnectorDataTypesPtrOutput) Exchange() OfficeDataConnectorDataTypesExchangePtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesExchange {
		if v == nil {
			return nil
		}
		return v.Exchange
	}).(OfficeDataConnectorDataTypesExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesPtrOutput) SharePoint() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesSharePoint {
		if v == nil {
			return nil
		}
		return v.SharePoint
	}).(OfficeDataConnectorDataTypesSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesPtrOutput) Teams() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesTeams {
		if v == nil {
			return nil
		}
		return v.Teams
	}).(OfficeDataConnectorDataTypesTeamsPtrOutput)
}

type OfficeDataConnectorDataTypesExchange struct {
	State *string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesExchangeInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesExchangeOutput() OfficeDataConnectorDataTypesExchangeOutput
	ToOfficeDataConnectorDataTypesExchangeOutputWithContext(context.Context) OfficeDataConnectorDataTypesExchangeOutput
}

type OfficeDataConnectorDataTypesExchangeArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (OfficeDataConnectorDataTypesExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesExchange)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesExchangeArgs) ToOfficeDataConnectorDataTypesExchangeOutput() OfficeDataConnectorDataTypesExchangeOutput {
	return i.ToOfficeDataConnectorDataTypesExchangeOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesExchangeArgs) ToOfficeDataConnectorDataTypesExchangeOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesExchangeOutput)
}

func (i OfficeDataConnectorDataTypesExchangeArgs) ToOfficeDataConnectorDataTypesExchangePtrOutput() OfficeDataConnectorDataTypesExchangePtrOutput {
	return i.ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesExchangeArgs) ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesExchangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesExchangeOutput).ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesExchangePtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesExchangePtrOutput() OfficeDataConnectorDataTypesExchangePtrOutput
	ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesExchangePtrOutput
}

type officeDataConnectorDataTypesExchangePtrType OfficeDataConnectorDataTypesExchangeArgs

func OfficeDataConnectorDataTypesExchangePtr(v *OfficeDataConnectorDataTypesExchangeArgs) OfficeDataConnectorDataTypesExchangePtrInput {
	return (*officeDataConnectorDataTypesExchangePtrType)(v)
}

func (*officeDataConnectorDataTypesExchangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesExchange)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesExchangePtrType) ToOfficeDataConnectorDataTypesExchangePtrOutput() OfficeDataConnectorDataTypesExchangePtrOutput {
	return i.ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesExchangePtrType) ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesExchangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesExchangePtrOutput)
}

type OfficeDataConnectorDataTypesExchangeOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesExchange)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesExchangeOutput) ToOfficeDataConnectorDataTypesExchangeOutput() OfficeDataConnectorDataTypesExchangeOutput {
	return o
}

func (o OfficeDataConnectorDataTypesExchangeOutput) ToOfficeDataConnectorDataTypesExchangeOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesExchangeOutput {
	return o
}

func (o OfficeDataConnectorDataTypesExchangeOutput) ToOfficeDataConnectorDataTypesExchangePtrOutput() OfficeDataConnectorDataTypesExchangePtrOutput {
	return o.ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesExchangeOutput) ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesExchangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesExchange) *OfficeDataConnectorDataTypesExchange {
		return &v
	}).(OfficeDataConnectorDataTypesExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesExchangeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesExchange) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesExchangePtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesExchangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesExchange)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesExchangePtrOutput) ToOfficeDataConnectorDataTypesExchangePtrOutput() OfficeDataConnectorDataTypesExchangePtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesExchangePtrOutput) ToOfficeDataConnectorDataTypesExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesExchangePtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesExchangePtrOutput) Elem() OfficeDataConnectorDataTypesExchangeOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesExchange) OfficeDataConnectorDataTypesExchange {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesExchange
		return ret
	}).(OfficeDataConnectorDataTypesExchangeOutput)
}

func (o OfficeDataConnectorDataTypesExchangePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesExchange) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponse struct {
	Exchange   *OfficeDataConnectorDataTypesResponseExchange   `pulumi:"exchange"`
	SharePoint *OfficeDataConnectorDataTypesResponseSharePoint `pulumi:"sharePoint"`
	Teams      *OfficeDataConnectorDataTypesResponseTeams      `pulumi:"teams"`
}

type OfficeDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseOutput) ToOfficeDataConnectorDataTypesResponseOutput() OfficeDataConnectorDataTypesResponseOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseOutput) ToOfficeDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseOutput) Exchange() OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseExchange {
		return v.Exchange
	}).(OfficeDataConnectorDataTypesResponseExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesResponseOutput) SharePoint() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseSharePoint {
		return v.SharePoint
	}).(OfficeDataConnectorDataTypesResponseSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesResponseOutput) Teams() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseTeams {
		return v.Teams
	}).(OfficeDataConnectorDataTypesResponseTeamsPtrOutput)
}

type OfficeDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) ToOfficeDataConnectorDataTypesResponsePtrOutput() OfficeDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) Elem() OfficeDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponse) OfficeDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesResponse
		return ret
	}).(OfficeDataConnectorDataTypesResponseOutput)
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) Exchange() OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseExchange {
		if v == nil {
			return nil
		}
		return v.Exchange
	}).(OfficeDataConnectorDataTypesResponseExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) SharePoint() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseSharePoint {
		if v == nil {
			return nil
		}
		return v.SharePoint
	}).(OfficeDataConnectorDataTypesResponseSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) Teams() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseTeams {
		if v == nil {
			return nil
		}
		return v.Teams
	}).(OfficeDataConnectorDataTypesResponseTeamsPtrOutput)
}

type OfficeDataConnectorDataTypesResponseExchange struct {
	State *string `pulumi:"state"`
}

type OfficeDataConnectorDataTypesResponseExchangeOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseExchangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponseExchange)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) ToOfficeDataConnectorDataTypesResponseExchangeOutput() OfficeDataConnectorDataTypesResponseExchangeOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) ToOfficeDataConnectorDataTypesResponseExchangeOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseExchangeOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseExchange) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseExchangePtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseExchangePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponseExchange)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseExchangePtrOutput) ToOfficeDataConnectorDataTypesResponseExchangePtrOutput() OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseExchangePtrOutput) ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseExchangePtrOutput) Elem() OfficeDataConnectorDataTypesResponseExchangeOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponseExchange) OfficeDataConnectorDataTypesResponseExchange {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesResponseExchange
		return ret
	}).(OfficeDataConnectorDataTypesResponseExchangeOutput)
}

func (o OfficeDataConnectorDataTypesResponseExchangePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponseExchange) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseSharePoint struct {
	State *string `pulumi:"state"`
}

type OfficeDataConnectorDataTypesResponseSharePointOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseSharePointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponseSharePoint)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) ToOfficeDataConnectorDataTypesResponseSharePointOutput() OfficeDataConnectorDataTypesResponseSharePointOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) ToOfficeDataConnectorDataTypesResponseSharePointOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseSharePointOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseSharePoint) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseSharePointPtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseSharePointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponseSharePoint)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseSharePointPtrOutput) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutput() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseSharePointPtrOutput) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseSharePointPtrOutput) Elem() OfficeDataConnectorDataTypesResponseSharePointOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponseSharePoint) OfficeDataConnectorDataTypesResponseSharePoint {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesResponseSharePoint
		return ret
	}).(OfficeDataConnectorDataTypesResponseSharePointOutput)
}

func (o OfficeDataConnectorDataTypesResponseSharePointPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponseSharePoint) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseTeams struct {
	State *string `pulumi:"state"`
}

type OfficeDataConnectorDataTypesResponseTeamsOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseTeamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponseTeams)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) ToOfficeDataConnectorDataTypesResponseTeamsOutput() OfficeDataConnectorDataTypesResponseTeamsOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) ToOfficeDataConnectorDataTypesResponseTeamsOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseTeamsOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseTeams) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseTeamsPtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesResponseTeamsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponseTeams)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesResponseTeamsPtrOutput) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutput() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseTeamsPtrOutput) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesResponseTeamsPtrOutput) Elem() OfficeDataConnectorDataTypesResponseTeamsOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponseTeams) OfficeDataConnectorDataTypesResponseTeams {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesResponseTeams
		return ret
	}).(OfficeDataConnectorDataTypesResponseTeamsOutput)
}

func (o OfficeDataConnectorDataTypesResponseTeamsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponseTeams) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesSharePoint struct {
	State *string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesSharePointInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesSharePointOutput() OfficeDataConnectorDataTypesSharePointOutput
	ToOfficeDataConnectorDataTypesSharePointOutputWithContext(context.Context) OfficeDataConnectorDataTypesSharePointOutput
}

type OfficeDataConnectorDataTypesSharePointArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (OfficeDataConnectorDataTypesSharePointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesSharePoint)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesSharePointArgs) ToOfficeDataConnectorDataTypesSharePointOutput() OfficeDataConnectorDataTypesSharePointOutput {
	return i.ToOfficeDataConnectorDataTypesSharePointOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesSharePointArgs) ToOfficeDataConnectorDataTypesSharePointOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesSharePointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesSharePointOutput)
}

func (i OfficeDataConnectorDataTypesSharePointArgs) ToOfficeDataConnectorDataTypesSharePointPtrOutput() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return i.ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesSharePointArgs) ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesSharePointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesSharePointOutput).ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesSharePointPtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesSharePointPtrOutput() OfficeDataConnectorDataTypesSharePointPtrOutput
	ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesSharePointPtrOutput
}

type officeDataConnectorDataTypesSharePointPtrType OfficeDataConnectorDataTypesSharePointArgs

func OfficeDataConnectorDataTypesSharePointPtr(v *OfficeDataConnectorDataTypesSharePointArgs) OfficeDataConnectorDataTypesSharePointPtrInput {
	return (*officeDataConnectorDataTypesSharePointPtrType)(v)
}

func (*officeDataConnectorDataTypesSharePointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesSharePoint)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesSharePointPtrType) ToOfficeDataConnectorDataTypesSharePointPtrOutput() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return i.ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesSharePointPtrType) ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesSharePointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesSharePointPtrOutput)
}

type OfficeDataConnectorDataTypesSharePointOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesSharePointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesSharePoint)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesSharePointOutput) ToOfficeDataConnectorDataTypesSharePointOutput() OfficeDataConnectorDataTypesSharePointOutput {
	return o
}

func (o OfficeDataConnectorDataTypesSharePointOutput) ToOfficeDataConnectorDataTypesSharePointOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesSharePointOutput {
	return o
}

func (o OfficeDataConnectorDataTypesSharePointOutput) ToOfficeDataConnectorDataTypesSharePointPtrOutput() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o.ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesSharePointOutput) ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesSharePoint) *OfficeDataConnectorDataTypesSharePoint {
		return &v
	}).(OfficeDataConnectorDataTypesSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesSharePointOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesSharePoint) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesSharePointPtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesSharePointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesSharePoint)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesSharePointPtrOutput) ToOfficeDataConnectorDataTypesSharePointPtrOutput() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesSharePointPtrOutput) ToOfficeDataConnectorDataTypesSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesSharePointPtrOutput) Elem() OfficeDataConnectorDataTypesSharePointOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesSharePoint) OfficeDataConnectorDataTypesSharePoint {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesSharePoint
		return ret
	}).(OfficeDataConnectorDataTypesSharePointOutput)
}

func (o OfficeDataConnectorDataTypesSharePointPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesSharePoint) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesTeams struct {
	State *string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesTeamsInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesTeamsOutput() OfficeDataConnectorDataTypesTeamsOutput
	ToOfficeDataConnectorDataTypesTeamsOutputWithContext(context.Context) OfficeDataConnectorDataTypesTeamsOutput
}

type OfficeDataConnectorDataTypesTeamsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (OfficeDataConnectorDataTypesTeamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesTeams)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesTeamsArgs) ToOfficeDataConnectorDataTypesTeamsOutput() OfficeDataConnectorDataTypesTeamsOutput {
	return i.ToOfficeDataConnectorDataTypesTeamsOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesTeamsArgs) ToOfficeDataConnectorDataTypesTeamsOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesTeamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesTeamsOutput)
}

func (i OfficeDataConnectorDataTypesTeamsArgs) ToOfficeDataConnectorDataTypesTeamsPtrOutput() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return i.ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesTeamsArgs) ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesTeamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesTeamsOutput).ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesTeamsPtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesTeamsPtrOutput() OfficeDataConnectorDataTypesTeamsPtrOutput
	ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesTeamsPtrOutput
}

type officeDataConnectorDataTypesTeamsPtrType OfficeDataConnectorDataTypesTeamsArgs

func OfficeDataConnectorDataTypesTeamsPtr(v *OfficeDataConnectorDataTypesTeamsArgs) OfficeDataConnectorDataTypesTeamsPtrInput {
	return (*officeDataConnectorDataTypesTeamsPtrType)(v)
}

func (*officeDataConnectorDataTypesTeamsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesTeams)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesTeamsPtrType) ToOfficeDataConnectorDataTypesTeamsPtrOutput() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return i.ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesTeamsPtrType) ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesTeamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesTeamsPtrOutput)
}

type OfficeDataConnectorDataTypesTeamsOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesTeamsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesTeams)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesTeamsOutput) ToOfficeDataConnectorDataTypesTeamsOutput() OfficeDataConnectorDataTypesTeamsOutput {
	return o
}

func (o OfficeDataConnectorDataTypesTeamsOutput) ToOfficeDataConnectorDataTypesTeamsOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesTeamsOutput {
	return o
}

func (o OfficeDataConnectorDataTypesTeamsOutput) ToOfficeDataConnectorDataTypesTeamsPtrOutput() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o.ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesTeamsOutput) ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesTeams) *OfficeDataConnectorDataTypesTeams {
		return &v
	}).(OfficeDataConnectorDataTypesTeamsPtrOutput)
}

func (o OfficeDataConnectorDataTypesTeamsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesTeams) *string { return v.State }).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesTeamsPtrOutput struct{ *pulumi.OutputState }

func (OfficeDataConnectorDataTypesTeamsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesTeams)(nil)).Elem()
}

func (o OfficeDataConnectorDataTypesTeamsPtrOutput) ToOfficeDataConnectorDataTypesTeamsPtrOutput() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesTeamsPtrOutput) ToOfficeDataConnectorDataTypesTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o
}

func (o OfficeDataConnectorDataTypesTeamsPtrOutput) Elem() OfficeDataConnectorDataTypesTeamsOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesTeams) OfficeDataConnectorDataTypesTeams {
		if v != nil {
			return *v
		}
		var ret OfficeDataConnectorDataTypesTeams
		return ret
	}).(OfficeDataConnectorDataTypesTeamsOutput)
}

func (o OfficeDataConnectorDataTypesTeamsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesTeams) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type PlaybookActionProperties struct {
	LogicAppResourceId string  `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}

type PlaybookActionPropertiesResponse struct {
	LogicAppResourceId string  `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}

type PropertyArrayChangedConditionProperties struct {
	ConditionProperties *AutomationRulePropertyArrayChangedValuesCondition `pulumi:"conditionProperties"`
	ConditionType       string                                             `pulumi:"conditionType"`
}

type PropertyArrayChangedConditionPropertiesResponse struct {
	ConditionProperties *AutomationRulePropertyArrayChangedValuesConditionResponse `pulumi:"conditionProperties"`
	ConditionType       string                                                     `pulumi:"conditionType"`
}

type PropertyChangedConditionProperties struct {
	ConditionProperties *AutomationRulePropertyValuesChangedCondition `pulumi:"conditionProperties"`
	ConditionType       string                                        `pulumi:"conditionType"`
}

type PropertyChangedConditionPropertiesResponse struct {
	ConditionProperties *AutomationRulePropertyValuesChangedConditionResponse `pulumi:"conditionProperties"`
	ConditionType       string                                                `pulumi:"conditionType"`
}

type PropertyConditionProperties struct {
	ConditionProperties *AutomationRulePropertyValuesCondition `pulumi:"conditionProperties"`
	ConditionType       string                                 `pulumi:"conditionType"`
}

type PropertyConditionPropertiesResponse struct {
	ConditionProperties *AutomationRulePropertyValuesConditionResponse `pulumi:"conditionProperties"`
	ConditionType       string                                         `pulumi:"conditionType"`
}

type SecurityMLAnalyticsSettingsDataSource struct {
	ConnectorId *string  `pulumi:"connectorId"`
	DataTypes   []string `pulumi:"dataTypes"`
}





type SecurityMLAnalyticsSettingsDataSourceInput interface {
	pulumi.Input

	ToSecurityMLAnalyticsSettingsDataSourceOutput() SecurityMLAnalyticsSettingsDataSourceOutput
	ToSecurityMLAnalyticsSettingsDataSourceOutputWithContext(context.Context) SecurityMLAnalyticsSettingsDataSourceOutput
}

type SecurityMLAnalyticsSettingsDataSourceArgs struct {
	ConnectorId pulumi.StringPtrInput   `pulumi:"connectorId"`
	DataTypes   pulumi.StringArrayInput `pulumi:"dataTypes"`
}

func (SecurityMLAnalyticsSettingsDataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityMLAnalyticsSettingsDataSource)(nil)).Elem()
}

func (i SecurityMLAnalyticsSettingsDataSourceArgs) ToSecurityMLAnalyticsSettingsDataSourceOutput() SecurityMLAnalyticsSettingsDataSourceOutput {
	return i.ToSecurityMLAnalyticsSettingsDataSourceOutputWithContext(context.Background())
}

func (i SecurityMLAnalyticsSettingsDataSourceArgs) ToSecurityMLAnalyticsSettingsDataSourceOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingsDataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityMLAnalyticsSettingsDataSourceOutput)
}





type SecurityMLAnalyticsSettingsDataSourceArrayInput interface {
	pulumi.Input

	ToSecurityMLAnalyticsSettingsDataSourceArrayOutput() SecurityMLAnalyticsSettingsDataSourceArrayOutput
	ToSecurityMLAnalyticsSettingsDataSourceArrayOutputWithContext(context.Context) SecurityMLAnalyticsSettingsDataSourceArrayOutput
}

type SecurityMLAnalyticsSettingsDataSourceArray []SecurityMLAnalyticsSettingsDataSourceInput

func (SecurityMLAnalyticsSettingsDataSourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityMLAnalyticsSettingsDataSource)(nil)).Elem()
}

func (i SecurityMLAnalyticsSettingsDataSourceArray) ToSecurityMLAnalyticsSettingsDataSourceArrayOutput() SecurityMLAnalyticsSettingsDataSourceArrayOutput {
	return i.ToSecurityMLAnalyticsSettingsDataSourceArrayOutputWithContext(context.Background())
}

func (i SecurityMLAnalyticsSettingsDataSourceArray) ToSecurityMLAnalyticsSettingsDataSourceArrayOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingsDataSourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityMLAnalyticsSettingsDataSourceArrayOutput)
}

type SecurityMLAnalyticsSettingsDataSourceOutput struct{ *pulumi.OutputState }

func (SecurityMLAnalyticsSettingsDataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityMLAnalyticsSettingsDataSource)(nil)).Elem()
}

func (o SecurityMLAnalyticsSettingsDataSourceOutput) ToSecurityMLAnalyticsSettingsDataSourceOutput() SecurityMLAnalyticsSettingsDataSourceOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceOutput) ToSecurityMLAnalyticsSettingsDataSourceOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingsDataSourceOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceOutput) ConnectorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityMLAnalyticsSettingsDataSource) *string { return v.ConnectorId }).(pulumi.StringPtrOutput)
}

func (o SecurityMLAnalyticsSettingsDataSourceOutput) DataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityMLAnalyticsSettingsDataSource) []string { return v.DataTypes }).(pulumi.StringArrayOutput)
}

type SecurityMLAnalyticsSettingsDataSourceArrayOutput struct{ *pulumi.OutputState }

func (SecurityMLAnalyticsSettingsDataSourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityMLAnalyticsSettingsDataSource)(nil)).Elem()
}

func (o SecurityMLAnalyticsSettingsDataSourceArrayOutput) ToSecurityMLAnalyticsSettingsDataSourceArrayOutput() SecurityMLAnalyticsSettingsDataSourceArrayOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceArrayOutput) ToSecurityMLAnalyticsSettingsDataSourceArrayOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingsDataSourceArrayOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceArrayOutput) Index(i pulumi.IntInput) SecurityMLAnalyticsSettingsDataSourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityMLAnalyticsSettingsDataSource {
		return vs[0].([]SecurityMLAnalyticsSettingsDataSource)[vs[1].(int)]
	}).(SecurityMLAnalyticsSettingsDataSourceOutput)
}

type SecurityMLAnalyticsSettingsDataSourceResponse struct {
	ConnectorId *string  `pulumi:"connectorId"`
	DataTypes   []string `pulumi:"dataTypes"`
}

type SecurityMLAnalyticsSettingsDataSourceResponseOutput struct{ *pulumi.OutputState }

func (SecurityMLAnalyticsSettingsDataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityMLAnalyticsSettingsDataSourceResponse)(nil)).Elem()
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseOutput) ToSecurityMLAnalyticsSettingsDataSourceResponseOutput() SecurityMLAnalyticsSettingsDataSourceResponseOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseOutput) ToSecurityMLAnalyticsSettingsDataSourceResponseOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingsDataSourceResponseOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseOutput) ConnectorId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityMLAnalyticsSettingsDataSourceResponse) *string { return v.ConnectorId }).(pulumi.StringPtrOutput)
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseOutput) DataTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SecurityMLAnalyticsSettingsDataSourceResponse) []string { return v.DataTypes }).(pulumi.StringArrayOutput)
}

type SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput struct{ *pulumi.OutputState }

func (SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecurityMLAnalyticsSettingsDataSourceResponse)(nil)).Elem()
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput) ToSecurityMLAnalyticsSettingsDataSourceResponseArrayOutput() SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput) ToSecurityMLAnalyticsSettingsDataSourceResponseArrayOutputWithContext(ctx context.Context) SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput {
	return o
}

func (o SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput) Index(i pulumi.IntInput) SecurityMLAnalyticsSettingsDataSourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecurityMLAnalyticsSettingsDataSourceResponse {
		return vs[0].([]SecurityMLAnalyticsSettingsDataSourceResponse)[vs[1].(int)]
	}).(SecurityMLAnalyticsSettingsDataSourceResponseOutput)
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

type TIDataConnectorDataTypes struct {
	Indicators *TIDataConnectorDataTypesIndicators `pulumi:"indicators"`
}





type TIDataConnectorDataTypesInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesOutput() TIDataConnectorDataTypesOutput
	ToTIDataConnectorDataTypesOutputWithContext(context.Context) TIDataConnectorDataTypesOutput
}

type TIDataConnectorDataTypesArgs struct {
	Indicators TIDataConnectorDataTypesIndicatorsPtrInput `pulumi:"indicators"`
}

func (TIDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypes)(nil)).Elem()
}

func (i TIDataConnectorDataTypesArgs) ToTIDataConnectorDataTypesOutput() TIDataConnectorDataTypesOutput {
	return i.ToTIDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesArgs) ToTIDataConnectorDataTypesOutputWithContext(ctx context.Context) TIDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesOutput)
}

func (i TIDataConnectorDataTypesArgs) ToTIDataConnectorDataTypesPtrOutput() TIDataConnectorDataTypesPtrOutput {
	return i.ToTIDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesArgs) ToTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesOutput).ToTIDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type TIDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesPtrOutput() TIDataConnectorDataTypesPtrOutput
	ToTIDataConnectorDataTypesPtrOutputWithContext(context.Context) TIDataConnectorDataTypesPtrOutput
}

type tidataConnectorDataTypesPtrType TIDataConnectorDataTypesArgs

func TIDataConnectorDataTypesPtr(v *TIDataConnectorDataTypesArgs) TIDataConnectorDataTypesPtrInput {
	return (*tidataConnectorDataTypesPtrType)(v)
}

func (*tidataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypes)(nil)).Elem()
}

func (i *tidataConnectorDataTypesPtrType) ToTIDataConnectorDataTypesPtrOutput() TIDataConnectorDataTypesPtrOutput {
	return i.ToTIDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *tidataConnectorDataTypesPtrType) ToTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesPtrOutput)
}

type TIDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypes)(nil)).Elem()
}

func (o TIDataConnectorDataTypesOutput) ToTIDataConnectorDataTypesOutput() TIDataConnectorDataTypesOutput {
	return o
}

func (o TIDataConnectorDataTypesOutput) ToTIDataConnectorDataTypesOutputWithContext(ctx context.Context) TIDataConnectorDataTypesOutput {
	return o
}

func (o TIDataConnectorDataTypesOutput) ToTIDataConnectorDataTypesPtrOutput() TIDataConnectorDataTypesPtrOutput {
	return o.ToTIDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o TIDataConnectorDataTypesOutput) ToTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TIDataConnectorDataTypes) *TIDataConnectorDataTypes {
		return &v
	}).(TIDataConnectorDataTypesPtrOutput)
}

func (o TIDataConnectorDataTypesOutput) Indicators() TIDataConnectorDataTypesIndicatorsPtrOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypes) *TIDataConnectorDataTypesIndicators { return v.Indicators }).(TIDataConnectorDataTypesIndicatorsPtrOutput)
}

type TIDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypes)(nil)).Elem()
}

func (o TIDataConnectorDataTypesPtrOutput) ToTIDataConnectorDataTypesPtrOutput() TIDataConnectorDataTypesPtrOutput {
	return o
}

func (o TIDataConnectorDataTypesPtrOutput) ToTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesPtrOutput {
	return o
}

func (o TIDataConnectorDataTypesPtrOutput) Elem() TIDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypes) TIDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret TIDataConnectorDataTypes
		return ret
	}).(TIDataConnectorDataTypesOutput)
}

func (o TIDataConnectorDataTypesPtrOutput) Indicators() TIDataConnectorDataTypesIndicatorsPtrOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypes) *TIDataConnectorDataTypesIndicators {
		if v == nil {
			return nil
		}
		return v.Indicators
	}).(TIDataConnectorDataTypesIndicatorsPtrOutput)
}

type TIDataConnectorDataTypesIndicators struct {
	State *string `pulumi:"state"`
}





type TIDataConnectorDataTypesIndicatorsInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesIndicatorsOutput() TIDataConnectorDataTypesIndicatorsOutput
	ToTIDataConnectorDataTypesIndicatorsOutputWithContext(context.Context) TIDataConnectorDataTypesIndicatorsOutput
}

type TIDataConnectorDataTypesIndicatorsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
}

func (TIDataConnectorDataTypesIndicatorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypesIndicators)(nil)).Elem()
}

func (i TIDataConnectorDataTypesIndicatorsArgs) ToTIDataConnectorDataTypesIndicatorsOutput() TIDataConnectorDataTypesIndicatorsOutput {
	return i.ToTIDataConnectorDataTypesIndicatorsOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesIndicatorsArgs) ToTIDataConnectorDataTypesIndicatorsOutputWithContext(ctx context.Context) TIDataConnectorDataTypesIndicatorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesIndicatorsOutput)
}

func (i TIDataConnectorDataTypesIndicatorsArgs) ToTIDataConnectorDataTypesIndicatorsPtrOutput() TIDataConnectorDataTypesIndicatorsPtrOutput {
	return i.ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesIndicatorsArgs) ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesIndicatorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesIndicatorsOutput).ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(ctx)
}









type TIDataConnectorDataTypesIndicatorsPtrInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesIndicatorsPtrOutput() TIDataConnectorDataTypesIndicatorsPtrOutput
	ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(context.Context) TIDataConnectorDataTypesIndicatorsPtrOutput
}

type tidataConnectorDataTypesIndicatorsPtrType TIDataConnectorDataTypesIndicatorsArgs

func TIDataConnectorDataTypesIndicatorsPtr(v *TIDataConnectorDataTypesIndicatorsArgs) TIDataConnectorDataTypesIndicatorsPtrInput {
	return (*tidataConnectorDataTypesIndicatorsPtrType)(v)
}

func (*tidataConnectorDataTypesIndicatorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypesIndicators)(nil)).Elem()
}

func (i *tidataConnectorDataTypesIndicatorsPtrType) ToTIDataConnectorDataTypesIndicatorsPtrOutput() TIDataConnectorDataTypesIndicatorsPtrOutput {
	return i.ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(context.Background())
}

func (i *tidataConnectorDataTypesIndicatorsPtrType) ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesIndicatorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesIndicatorsPtrOutput)
}

type TIDataConnectorDataTypesIndicatorsOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesIndicatorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypesIndicators)(nil)).Elem()
}

func (o TIDataConnectorDataTypesIndicatorsOutput) ToTIDataConnectorDataTypesIndicatorsOutput() TIDataConnectorDataTypesIndicatorsOutput {
	return o
}

func (o TIDataConnectorDataTypesIndicatorsOutput) ToTIDataConnectorDataTypesIndicatorsOutputWithContext(ctx context.Context) TIDataConnectorDataTypesIndicatorsOutput {
	return o
}

func (o TIDataConnectorDataTypesIndicatorsOutput) ToTIDataConnectorDataTypesIndicatorsPtrOutput() TIDataConnectorDataTypesIndicatorsPtrOutput {
	return o.ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(context.Background())
}

func (o TIDataConnectorDataTypesIndicatorsOutput) ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesIndicatorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TIDataConnectorDataTypesIndicators) *TIDataConnectorDataTypesIndicators {
		return &v
	}).(TIDataConnectorDataTypesIndicatorsPtrOutput)
}

func (o TIDataConnectorDataTypesIndicatorsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesIndicators) *string { return v.State }).(pulumi.StringPtrOutput)
}

type TIDataConnectorDataTypesIndicatorsPtrOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesIndicatorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypesIndicators)(nil)).Elem()
}

func (o TIDataConnectorDataTypesIndicatorsPtrOutput) ToTIDataConnectorDataTypesIndicatorsPtrOutput() TIDataConnectorDataTypesIndicatorsPtrOutput {
	return o
}

func (o TIDataConnectorDataTypesIndicatorsPtrOutput) ToTIDataConnectorDataTypesIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesIndicatorsPtrOutput {
	return o
}

func (o TIDataConnectorDataTypesIndicatorsPtrOutput) Elem() TIDataConnectorDataTypesIndicatorsOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypesIndicators) TIDataConnectorDataTypesIndicators {
		if v != nil {
			return *v
		}
		var ret TIDataConnectorDataTypesIndicators
		return ret
	}).(TIDataConnectorDataTypesIndicatorsOutput)
}

func (o TIDataConnectorDataTypesIndicatorsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypesIndicators) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type TIDataConnectorDataTypesResponse struct {
	Indicators *TIDataConnectorDataTypesResponseIndicators `pulumi:"indicators"`
}

type TIDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o TIDataConnectorDataTypesResponseOutput) ToTIDataConnectorDataTypesResponseOutput() TIDataConnectorDataTypesResponseOutput {
	return o
}

func (o TIDataConnectorDataTypesResponseOutput) ToTIDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseOutput {
	return o
}

func (o TIDataConnectorDataTypesResponseOutput) Indicators() TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesResponse) *TIDataConnectorDataTypesResponseIndicators {
		return v.Indicators
	}).(TIDataConnectorDataTypesResponseIndicatorsPtrOutput)
}

type TIDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o TIDataConnectorDataTypesResponsePtrOutput) ToTIDataConnectorDataTypesResponsePtrOutput() TIDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o TIDataConnectorDataTypesResponsePtrOutput) ToTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o TIDataConnectorDataTypesResponsePtrOutput) Elem() TIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypesResponse) TIDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret TIDataConnectorDataTypesResponse
		return ret
	}).(TIDataConnectorDataTypesResponseOutput)
}

func (o TIDataConnectorDataTypesResponsePtrOutput) Indicators() TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypesResponse) *TIDataConnectorDataTypesResponseIndicators {
		if v == nil {
			return nil
		}
		return v.Indicators
	}).(TIDataConnectorDataTypesResponseIndicatorsPtrOutput)
}

type TIDataConnectorDataTypesResponseIndicators struct {
	State *string `pulumi:"state"`
}

type TIDataConnectorDataTypesResponseIndicatorsOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesResponseIndicatorsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypesResponseIndicators)(nil)).Elem()
}

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) ToTIDataConnectorDataTypesResponseIndicatorsOutput() TIDataConnectorDataTypesResponseIndicatorsOutput {
	return o
}

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) ToTIDataConnectorDataTypesResponseIndicatorsOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseIndicatorsOutput {
	return o
}

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesResponseIndicators) *string { return v.State }).(pulumi.StringPtrOutput)
}

type TIDataConnectorDataTypesResponseIndicatorsPtrOutput struct{ *pulumi.OutputState }

func (TIDataConnectorDataTypesResponseIndicatorsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypesResponseIndicators)(nil)).Elem()
}

func (o TIDataConnectorDataTypesResponseIndicatorsPtrOutput) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutput() TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return o
}

func (o TIDataConnectorDataTypesResponseIndicatorsPtrOutput) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return o
}

func (o TIDataConnectorDataTypesResponseIndicatorsPtrOutput) Elem() TIDataConnectorDataTypesResponseIndicatorsOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypesResponseIndicators) TIDataConnectorDataTypesResponseIndicators {
		if v != nil {
			return *v
		}
		var ret TIDataConnectorDataTypesResponseIndicators
		return ret
	}).(TIDataConnectorDataTypesResponseIndicatorsOutput)
}

func (o TIDataConnectorDataTypesResponseIndicatorsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TIDataConnectorDataTypesResponseIndicators) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ThreatIntelligenceExternalReference struct {
	Description *string           `pulumi:"description"`
	ExternalId  *string           `pulumi:"externalId"`
	Hashes      map[string]string `pulumi:"hashes"`
	SourceName  *string           `pulumi:"sourceName"`
	Url         *string           `pulumi:"url"`
}





type ThreatIntelligenceExternalReferenceInput interface {
	pulumi.Input

	ToThreatIntelligenceExternalReferenceOutput() ThreatIntelligenceExternalReferenceOutput
	ToThreatIntelligenceExternalReferenceOutputWithContext(context.Context) ThreatIntelligenceExternalReferenceOutput
}

type ThreatIntelligenceExternalReferenceArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	ExternalId  pulumi.StringPtrInput `pulumi:"externalId"`
	Hashes      pulumi.StringMapInput `pulumi:"hashes"`
	SourceName  pulumi.StringPtrInput `pulumi:"sourceName"`
	Url         pulumi.StringPtrInput `pulumi:"url"`
}

func (ThreatIntelligenceExternalReferenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceExternalReference)(nil)).Elem()
}

func (i ThreatIntelligenceExternalReferenceArgs) ToThreatIntelligenceExternalReferenceOutput() ThreatIntelligenceExternalReferenceOutput {
	return i.ToThreatIntelligenceExternalReferenceOutputWithContext(context.Background())
}

func (i ThreatIntelligenceExternalReferenceArgs) ToThreatIntelligenceExternalReferenceOutputWithContext(ctx context.Context) ThreatIntelligenceExternalReferenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceExternalReferenceOutput)
}





type ThreatIntelligenceExternalReferenceArrayInput interface {
	pulumi.Input

	ToThreatIntelligenceExternalReferenceArrayOutput() ThreatIntelligenceExternalReferenceArrayOutput
	ToThreatIntelligenceExternalReferenceArrayOutputWithContext(context.Context) ThreatIntelligenceExternalReferenceArrayOutput
}

type ThreatIntelligenceExternalReferenceArray []ThreatIntelligenceExternalReferenceInput

func (ThreatIntelligenceExternalReferenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceExternalReference)(nil)).Elem()
}

func (i ThreatIntelligenceExternalReferenceArray) ToThreatIntelligenceExternalReferenceArrayOutput() ThreatIntelligenceExternalReferenceArrayOutput {
	return i.ToThreatIntelligenceExternalReferenceArrayOutputWithContext(context.Background())
}

func (i ThreatIntelligenceExternalReferenceArray) ToThreatIntelligenceExternalReferenceArrayOutputWithContext(ctx context.Context) ThreatIntelligenceExternalReferenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceExternalReferenceArrayOutput)
}

type ThreatIntelligenceExternalReferenceOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceExternalReferenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceExternalReference)(nil)).Elem()
}

func (o ThreatIntelligenceExternalReferenceOutput) ToThreatIntelligenceExternalReferenceOutput() ThreatIntelligenceExternalReferenceOutput {
	return o
}

func (o ThreatIntelligenceExternalReferenceOutput) ToThreatIntelligenceExternalReferenceOutputWithContext(ctx context.Context) ThreatIntelligenceExternalReferenceOutput {
	return o
}

func (o ThreatIntelligenceExternalReferenceOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceExternalReference) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceExternalReferenceOutput) ExternalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceExternalReference) *string { return v.ExternalId }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceExternalReferenceOutput) Hashes() pulumi.StringMapOutput {
	return o.ApplyT(func(v ThreatIntelligenceExternalReference) map[string]string { return v.Hashes }).(pulumi.StringMapOutput)
}

func (o ThreatIntelligenceExternalReferenceOutput) SourceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceExternalReference) *string { return v.SourceName }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceExternalReferenceOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceExternalReference) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ThreatIntelligenceExternalReferenceArrayOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceExternalReferenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceExternalReference)(nil)).Elem()
}

func (o ThreatIntelligenceExternalReferenceArrayOutput) ToThreatIntelligenceExternalReferenceArrayOutput() ThreatIntelligenceExternalReferenceArrayOutput {
	return o
}

func (o ThreatIntelligenceExternalReferenceArrayOutput) ToThreatIntelligenceExternalReferenceArrayOutputWithContext(ctx context.Context) ThreatIntelligenceExternalReferenceArrayOutput {
	return o
}

func (o ThreatIntelligenceExternalReferenceArrayOutput) Index(i pulumi.IntInput) ThreatIntelligenceExternalReferenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThreatIntelligenceExternalReference {
		return vs[0].([]ThreatIntelligenceExternalReference)[vs[1].(int)]
	}).(ThreatIntelligenceExternalReferenceOutput)
}

type ThreatIntelligenceGranularMarkingModel struct {
	Language   *string  `pulumi:"language"`
	MarkingRef *int     `pulumi:"markingRef"`
	Selectors  []string `pulumi:"selectors"`
}





type ThreatIntelligenceGranularMarkingModelInput interface {
	pulumi.Input

	ToThreatIntelligenceGranularMarkingModelOutput() ThreatIntelligenceGranularMarkingModelOutput
	ToThreatIntelligenceGranularMarkingModelOutputWithContext(context.Context) ThreatIntelligenceGranularMarkingModelOutput
}

type ThreatIntelligenceGranularMarkingModelArgs struct {
	Language   pulumi.StringPtrInput   `pulumi:"language"`
	MarkingRef pulumi.IntPtrInput      `pulumi:"markingRef"`
	Selectors  pulumi.StringArrayInput `pulumi:"selectors"`
}

func (ThreatIntelligenceGranularMarkingModelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceGranularMarkingModel)(nil)).Elem()
}

func (i ThreatIntelligenceGranularMarkingModelArgs) ToThreatIntelligenceGranularMarkingModelOutput() ThreatIntelligenceGranularMarkingModelOutput {
	return i.ToThreatIntelligenceGranularMarkingModelOutputWithContext(context.Background())
}

func (i ThreatIntelligenceGranularMarkingModelArgs) ToThreatIntelligenceGranularMarkingModelOutputWithContext(ctx context.Context) ThreatIntelligenceGranularMarkingModelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceGranularMarkingModelOutput)
}





type ThreatIntelligenceGranularMarkingModelArrayInput interface {
	pulumi.Input

	ToThreatIntelligenceGranularMarkingModelArrayOutput() ThreatIntelligenceGranularMarkingModelArrayOutput
	ToThreatIntelligenceGranularMarkingModelArrayOutputWithContext(context.Context) ThreatIntelligenceGranularMarkingModelArrayOutput
}

type ThreatIntelligenceGranularMarkingModelArray []ThreatIntelligenceGranularMarkingModelInput

func (ThreatIntelligenceGranularMarkingModelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceGranularMarkingModel)(nil)).Elem()
}

func (i ThreatIntelligenceGranularMarkingModelArray) ToThreatIntelligenceGranularMarkingModelArrayOutput() ThreatIntelligenceGranularMarkingModelArrayOutput {
	return i.ToThreatIntelligenceGranularMarkingModelArrayOutputWithContext(context.Background())
}

func (i ThreatIntelligenceGranularMarkingModelArray) ToThreatIntelligenceGranularMarkingModelArrayOutputWithContext(ctx context.Context) ThreatIntelligenceGranularMarkingModelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceGranularMarkingModelArrayOutput)
}

type ThreatIntelligenceGranularMarkingModelOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceGranularMarkingModelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceGranularMarkingModel)(nil)).Elem()
}

func (o ThreatIntelligenceGranularMarkingModelOutput) ToThreatIntelligenceGranularMarkingModelOutput() ThreatIntelligenceGranularMarkingModelOutput {
	return o
}

func (o ThreatIntelligenceGranularMarkingModelOutput) ToThreatIntelligenceGranularMarkingModelOutputWithContext(ctx context.Context) ThreatIntelligenceGranularMarkingModelOutput {
	return o
}

func (o ThreatIntelligenceGranularMarkingModelOutput) Language() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceGranularMarkingModel) *string { return v.Language }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceGranularMarkingModelOutput) MarkingRef() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceGranularMarkingModel) *int { return v.MarkingRef }).(pulumi.IntPtrOutput)
}

func (o ThreatIntelligenceGranularMarkingModelOutput) Selectors() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ThreatIntelligenceGranularMarkingModel) []string { return v.Selectors }).(pulumi.StringArrayOutput)
}

type ThreatIntelligenceGranularMarkingModelArrayOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceGranularMarkingModelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceGranularMarkingModel)(nil)).Elem()
}

func (o ThreatIntelligenceGranularMarkingModelArrayOutput) ToThreatIntelligenceGranularMarkingModelArrayOutput() ThreatIntelligenceGranularMarkingModelArrayOutput {
	return o
}

func (o ThreatIntelligenceGranularMarkingModelArrayOutput) ToThreatIntelligenceGranularMarkingModelArrayOutputWithContext(ctx context.Context) ThreatIntelligenceGranularMarkingModelArrayOutput {
	return o
}

func (o ThreatIntelligenceGranularMarkingModelArrayOutput) Index(i pulumi.IntInput) ThreatIntelligenceGranularMarkingModelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThreatIntelligenceGranularMarkingModel {
		return vs[0].([]ThreatIntelligenceGranularMarkingModel)[vs[1].(int)]
	}).(ThreatIntelligenceGranularMarkingModelOutput)
}

type ThreatIntelligenceKillChainPhase struct {
	KillChainName *string `pulumi:"killChainName"`
	PhaseName     *string `pulumi:"phaseName"`
}





type ThreatIntelligenceKillChainPhaseInput interface {
	pulumi.Input

	ToThreatIntelligenceKillChainPhaseOutput() ThreatIntelligenceKillChainPhaseOutput
	ToThreatIntelligenceKillChainPhaseOutputWithContext(context.Context) ThreatIntelligenceKillChainPhaseOutput
}

type ThreatIntelligenceKillChainPhaseArgs struct {
	KillChainName pulumi.StringPtrInput `pulumi:"killChainName"`
	PhaseName     pulumi.StringPtrInput `pulumi:"phaseName"`
}

func (ThreatIntelligenceKillChainPhaseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceKillChainPhase)(nil)).Elem()
}

func (i ThreatIntelligenceKillChainPhaseArgs) ToThreatIntelligenceKillChainPhaseOutput() ThreatIntelligenceKillChainPhaseOutput {
	return i.ToThreatIntelligenceKillChainPhaseOutputWithContext(context.Background())
}

func (i ThreatIntelligenceKillChainPhaseArgs) ToThreatIntelligenceKillChainPhaseOutputWithContext(ctx context.Context) ThreatIntelligenceKillChainPhaseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceKillChainPhaseOutput)
}





type ThreatIntelligenceKillChainPhaseArrayInput interface {
	pulumi.Input

	ToThreatIntelligenceKillChainPhaseArrayOutput() ThreatIntelligenceKillChainPhaseArrayOutput
	ToThreatIntelligenceKillChainPhaseArrayOutputWithContext(context.Context) ThreatIntelligenceKillChainPhaseArrayOutput
}

type ThreatIntelligenceKillChainPhaseArray []ThreatIntelligenceKillChainPhaseInput

func (ThreatIntelligenceKillChainPhaseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceKillChainPhase)(nil)).Elem()
}

func (i ThreatIntelligenceKillChainPhaseArray) ToThreatIntelligenceKillChainPhaseArrayOutput() ThreatIntelligenceKillChainPhaseArrayOutput {
	return i.ToThreatIntelligenceKillChainPhaseArrayOutputWithContext(context.Background())
}

func (i ThreatIntelligenceKillChainPhaseArray) ToThreatIntelligenceKillChainPhaseArrayOutputWithContext(ctx context.Context) ThreatIntelligenceKillChainPhaseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceKillChainPhaseArrayOutput)
}

type ThreatIntelligenceKillChainPhaseOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceKillChainPhaseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceKillChainPhase)(nil)).Elem()
}

func (o ThreatIntelligenceKillChainPhaseOutput) ToThreatIntelligenceKillChainPhaseOutput() ThreatIntelligenceKillChainPhaseOutput {
	return o
}

func (o ThreatIntelligenceKillChainPhaseOutput) ToThreatIntelligenceKillChainPhaseOutputWithContext(ctx context.Context) ThreatIntelligenceKillChainPhaseOutput {
	return o
}

func (o ThreatIntelligenceKillChainPhaseOutput) KillChainName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceKillChainPhase) *string { return v.KillChainName }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceKillChainPhaseOutput) PhaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceKillChainPhase) *string { return v.PhaseName }).(pulumi.StringPtrOutput)
}

type ThreatIntelligenceKillChainPhaseArrayOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceKillChainPhaseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceKillChainPhase)(nil)).Elem()
}

func (o ThreatIntelligenceKillChainPhaseArrayOutput) ToThreatIntelligenceKillChainPhaseArrayOutput() ThreatIntelligenceKillChainPhaseArrayOutput {
	return o
}

func (o ThreatIntelligenceKillChainPhaseArrayOutput) ToThreatIntelligenceKillChainPhaseArrayOutputWithContext(ctx context.Context) ThreatIntelligenceKillChainPhaseArrayOutput {
	return o
}

func (o ThreatIntelligenceKillChainPhaseArrayOutput) Index(i pulumi.IntInput) ThreatIntelligenceKillChainPhaseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThreatIntelligenceKillChainPhase {
		return vs[0].([]ThreatIntelligenceKillChainPhase)[vs[1].(int)]
	}).(ThreatIntelligenceKillChainPhaseOutput)
}

type ThreatIntelligenceParsedPattern struct {
	PatternTypeKey    *string                                    `pulumi:"patternTypeKey"`
	PatternTypeValues []ThreatIntelligenceParsedPatternTypeValue `pulumi:"patternTypeValues"`
}





type ThreatIntelligenceParsedPatternInput interface {
	pulumi.Input

	ToThreatIntelligenceParsedPatternOutput() ThreatIntelligenceParsedPatternOutput
	ToThreatIntelligenceParsedPatternOutputWithContext(context.Context) ThreatIntelligenceParsedPatternOutput
}

type ThreatIntelligenceParsedPatternArgs struct {
	PatternTypeKey    pulumi.StringPtrInput                              `pulumi:"patternTypeKey"`
	PatternTypeValues ThreatIntelligenceParsedPatternTypeValueArrayInput `pulumi:"patternTypeValues"`
}

func (ThreatIntelligenceParsedPatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceParsedPattern)(nil)).Elem()
}

func (i ThreatIntelligenceParsedPatternArgs) ToThreatIntelligenceParsedPatternOutput() ThreatIntelligenceParsedPatternOutput {
	return i.ToThreatIntelligenceParsedPatternOutputWithContext(context.Background())
}

func (i ThreatIntelligenceParsedPatternArgs) ToThreatIntelligenceParsedPatternOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceParsedPatternOutput)
}





type ThreatIntelligenceParsedPatternArrayInput interface {
	pulumi.Input

	ToThreatIntelligenceParsedPatternArrayOutput() ThreatIntelligenceParsedPatternArrayOutput
	ToThreatIntelligenceParsedPatternArrayOutputWithContext(context.Context) ThreatIntelligenceParsedPatternArrayOutput
}

type ThreatIntelligenceParsedPatternArray []ThreatIntelligenceParsedPatternInput

func (ThreatIntelligenceParsedPatternArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceParsedPattern)(nil)).Elem()
}

func (i ThreatIntelligenceParsedPatternArray) ToThreatIntelligenceParsedPatternArrayOutput() ThreatIntelligenceParsedPatternArrayOutput {
	return i.ToThreatIntelligenceParsedPatternArrayOutputWithContext(context.Background())
}

func (i ThreatIntelligenceParsedPatternArray) ToThreatIntelligenceParsedPatternArrayOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceParsedPatternArrayOutput)
}

type ThreatIntelligenceParsedPatternOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceParsedPatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceParsedPattern)(nil)).Elem()
}

func (o ThreatIntelligenceParsedPatternOutput) ToThreatIntelligenceParsedPatternOutput() ThreatIntelligenceParsedPatternOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternOutput) ToThreatIntelligenceParsedPatternOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternOutput) PatternTypeKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceParsedPattern) *string { return v.PatternTypeKey }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceParsedPatternOutput) PatternTypeValues() ThreatIntelligenceParsedPatternTypeValueArrayOutput {
	return o.ApplyT(func(v ThreatIntelligenceParsedPattern) []ThreatIntelligenceParsedPatternTypeValue {
		return v.PatternTypeValues
	}).(ThreatIntelligenceParsedPatternTypeValueArrayOutput)
}

type ThreatIntelligenceParsedPatternArrayOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceParsedPatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceParsedPattern)(nil)).Elem()
}

func (o ThreatIntelligenceParsedPatternArrayOutput) ToThreatIntelligenceParsedPatternArrayOutput() ThreatIntelligenceParsedPatternArrayOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternArrayOutput) ToThreatIntelligenceParsedPatternArrayOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternArrayOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternArrayOutput) Index(i pulumi.IntInput) ThreatIntelligenceParsedPatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThreatIntelligenceParsedPattern {
		return vs[0].([]ThreatIntelligenceParsedPattern)[vs[1].(int)]
	}).(ThreatIntelligenceParsedPatternOutput)
}

type ThreatIntelligenceParsedPatternTypeValue struct {
	Value     *string `pulumi:"value"`
	ValueType *string `pulumi:"valueType"`
}





type ThreatIntelligenceParsedPatternTypeValueInput interface {
	pulumi.Input

	ToThreatIntelligenceParsedPatternTypeValueOutput() ThreatIntelligenceParsedPatternTypeValueOutput
	ToThreatIntelligenceParsedPatternTypeValueOutputWithContext(context.Context) ThreatIntelligenceParsedPatternTypeValueOutput
}

type ThreatIntelligenceParsedPatternTypeValueArgs struct {
	Value     pulumi.StringPtrInput `pulumi:"value"`
	ValueType pulumi.StringPtrInput `pulumi:"valueType"`
}

func (ThreatIntelligenceParsedPatternTypeValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceParsedPatternTypeValue)(nil)).Elem()
}

func (i ThreatIntelligenceParsedPatternTypeValueArgs) ToThreatIntelligenceParsedPatternTypeValueOutput() ThreatIntelligenceParsedPatternTypeValueOutput {
	return i.ToThreatIntelligenceParsedPatternTypeValueOutputWithContext(context.Background())
}

func (i ThreatIntelligenceParsedPatternTypeValueArgs) ToThreatIntelligenceParsedPatternTypeValueOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternTypeValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceParsedPatternTypeValueOutput)
}





type ThreatIntelligenceParsedPatternTypeValueArrayInput interface {
	pulumi.Input

	ToThreatIntelligenceParsedPatternTypeValueArrayOutput() ThreatIntelligenceParsedPatternTypeValueArrayOutput
	ToThreatIntelligenceParsedPatternTypeValueArrayOutputWithContext(context.Context) ThreatIntelligenceParsedPatternTypeValueArrayOutput
}

type ThreatIntelligenceParsedPatternTypeValueArray []ThreatIntelligenceParsedPatternTypeValueInput

func (ThreatIntelligenceParsedPatternTypeValueArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceParsedPatternTypeValue)(nil)).Elem()
}

func (i ThreatIntelligenceParsedPatternTypeValueArray) ToThreatIntelligenceParsedPatternTypeValueArrayOutput() ThreatIntelligenceParsedPatternTypeValueArrayOutput {
	return i.ToThreatIntelligenceParsedPatternTypeValueArrayOutputWithContext(context.Background())
}

func (i ThreatIntelligenceParsedPatternTypeValueArray) ToThreatIntelligenceParsedPatternTypeValueArrayOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternTypeValueArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceParsedPatternTypeValueArrayOutput)
}

type ThreatIntelligenceParsedPatternTypeValueOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceParsedPatternTypeValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceParsedPatternTypeValue)(nil)).Elem()
}

func (o ThreatIntelligenceParsedPatternTypeValueOutput) ToThreatIntelligenceParsedPatternTypeValueOutput() ThreatIntelligenceParsedPatternTypeValueOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternTypeValueOutput) ToThreatIntelligenceParsedPatternTypeValueOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternTypeValueOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternTypeValueOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceParsedPatternTypeValue) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceParsedPatternTypeValueOutput) ValueType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ThreatIntelligenceParsedPatternTypeValue) *string { return v.ValueType }).(pulumi.StringPtrOutput)
}

type ThreatIntelligenceParsedPatternTypeValueArrayOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceParsedPatternTypeValueArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ThreatIntelligenceParsedPatternTypeValue)(nil)).Elem()
}

func (o ThreatIntelligenceParsedPatternTypeValueArrayOutput) ToThreatIntelligenceParsedPatternTypeValueArrayOutput() ThreatIntelligenceParsedPatternTypeValueArrayOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternTypeValueArrayOutput) ToThreatIntelligenceParsedPatternTypeValueArrayOutputWithContext(ctx context.Context) ThreatIntelligenceParsedPatternTypeValueArrayOutput {
	return o
}

func (o ThreatIntelligenceParsedPatternTypeValueArrayOutput) Index(i pulumi.IntInput) ThreatIntelligenceParsedPatternTypeValueOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ThreatIntelligenceParsedPatternTypeValue {
		return vs[0].([]ThreatIntelligenceParsedPatternTypeValue)[vs[1].(int)]
	}).(ThreatIntelligenceParsedPatternTypeValueOutput)
}

type UserInfo struct {
	ObjectId *string `pulumi:"objectId"`
}





type UserInfoInput interface {
	pulumi.Input

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(context.Context) UserInfoOutput
}

type UserInfoArgs struct {
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (UserInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (i UserInfoArgs) ToUserInfoOutput() UserInfoOutput {
	return i.ToUserInfoOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput)
}

func (i UserInfoArgs) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i UserInfoArgs) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoOutput).ToUserInfoPtrOutputWithContext(ctx)
}









type UserInfoPtrInput interface {
	pulumi.Input

	ToUserInfoPtrOutput() UserInfoPtrOutput
	ToUserInfoPtrOutputWithContext(context.Context) UserInfoPtrOutput
}

type userInfoPtrType UserInfoArgs

func UserInfoPtr(v *UserInfoArgs) UserInfoPtrInput {
	return (*userInfoPtrType)(v)
}

func (*userInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (i *userInfoPtrType) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return i.ToUserInfoPtrOutputWithContext(context.Background())
}

func (i *userInfoPtrType) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoPtrOutput)
}

type UserInfoOutput struct{ *pulumi.OutputState }

func (UserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfo)(nil)).Elem()
}

func (o UserInfoOutput) ToUserInfoOutput() UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoOutputWithContext(ctx context.Context) UserInfoOutput {
	return o
}

func (o UserInfoOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o.ToUserInfoPtrOutputWithContext(context.Background())
}

func (o UserInfoOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfo) *UserInfo {
		return &v
	}).(UserInfoPtrOutput)
}

func (o UserInfoOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfo) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

type UserInfoPtrOutput struct{ *pulumi.OutputState }

func (UserInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfo)(nil)).Elem()
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutput() UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) ToUserInfoPtrOutputWithContext(ctx context.Context) UserInfoPtrOutput {
	return o
}

func (o UserInfoPtrOutput) Elem() UserInfoOutput {
	return o.ApplyT(func(v *UserInfo) UserInfo {
		if v != nil {
			return *v
		}
		var ret UserInfo
		return ret
	}).(UserInfoOutput)
}

func (o UserInfoPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfo) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

type UserInfoResponse struct {
	Email    string  `pulumi:"email"`
	Name     string  `pulumi:"name"`
	ObjectId *string `pulumi:"objectId"`
}

type UserInfoResponseOutput struct{ *pulumi.OutputState }

func (UserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return o
}

func (o UserInfoResponseOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v UserInfoResponse) string { return v.Email }).(pulumi.StringOutput)
}

func (o UserInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v UserInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o UserInfoResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserInfoResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

type UserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o
}

func (o UserInfoResponsePtrOutput) Elem() UserInfoResponseOutput {
	return o.ApplyT(func(v *UserInfoResponse) UserInfoResponse {
		if v != nil {
			return *v
		}
		var ret UserInfoResponse
		return ret
	}).(UserInfoResponseOutput)
}

func (o UserInfoResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o UserInfoResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

type WatchlistUserInfo struct {
	ObjectId *string `pulumi:"objectId"`
}





type WatchlistUserInfoInput interface {
	pulumi.Input

	ToWatchlistUserInfoOutput() WatchlistUserInfoOutput
	ToWatchlistUserInfoOutputWithContext(context.Context) WatchlistUserInfoOutput
}

type WatchlistUserInfoArgs struct {
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (WatchlistUserInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WatchlistUserInfo)(nil)).Elem()
}

func (i WatchlistUserInfoArgs) ToWatchlistUserInfoOutput() WatchlistUserInfoOutput {
	return i.ToWatchlistUserInfoOutputWithContext(context.Background())
}

func (i WatchlistUserInfoArgs) ToWatchlistUserInfoOutputWithContext(ctx context.Context) WatchlistUserInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistUserInfoOutput)
}

func (i WatchlistUserInfoArgs) ToWatchlistUserInfoPtrOutput() WatchlistUserInfoPtrOutput {
	return i.ToWatchlistUserInfoPtrOutputWithContext(context.Background())
}

func (i WatchlistUserInfoArgs) ToWatchlistUserInfoPtrOutputWithContext(ctx context.Context) WatchlistUserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistUserInfoOutput).ToWatchlistUserInfoPtrOutputWithContext(ctx)
}









type WatchlistUserInfoPtrInput interface {
	pulumi.Input

	ToWatchlistUserInfoPtrOutput() WatchlistUserInfoPtrOutput
	ToWatchlistUserInfoPtrOutputWithContext(context.Context) WatchlistUserInfoPtrOutput
}

type watchlistUserInfoPtrType WatchlistUserInfoArgs

func WatchlistUserInfoPtr(v *WatchlistUserInfoArgs) WatchlistUserInfoPtrInput {
	return (*watchlistUserInfoPtrType)(v)
}

func (*watchlistUserInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistUserInfo)(nil)).Elem()
}

func (i *watchlistUserInfoPtrType) ToWatchlistUserInfoPtrOutput() WatchlistUserInfoPtrOutput {
	return i.ToWatchlistUserInfoPtrOutputWithContext(context.Background())
}

func (i *watchlistUserInfoPtrType) ToWatchlistUserInfoPtrOutputWithContext(ctx context.Context) WatchlistUserInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistUserInfoPtrOutput)
}

type WatchlistUserInfoOutput struct{ *pulumi.OutputState }

func (WatchlistUserInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WatchlistUserInfo)(nil)).Elem()
}

func (o WatchlistUserInfoOutput) ToWatchlistUserInfoOutput() WatchlistUserInfoOutput {
	return o
}

func (o WatchlistUserInfoOutput) ToWatchlistUserInfoOutputWithContext(ctx context.Context) WatchlistUserInfoOutput {
	return o
}

func (o WatchlistUserInfoOutput) ToWatchlistUserInfoPtrOutput() WatchlistUserInfoPtrOutput {
	return o.ToWatchlistUserInfoPtrOutputWithContext(context.Background())
}

func (o WatchlistUserInfoOutput) ToWatchlistUserInfoPtrOutputWithContext(ctx context.Context) WatchlistUserInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WatchlistUserInfo) *WatchlistUserInfo {
		return &v
	}).(WatchlistUserInfoPtrOutput)
}

func (o WatchlistUserInfoOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WatchlistUserInfo) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

type WatchlistUserInfoPtrOutput struct{ *pulumi.OutputState }

func (WatchlistUserInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistUserInfo)(nil)).Elem()
}

func (o WatchlistUserInfoPtrOutput) ToWatchlistUserInfoPtrOutput() WatchlistUserInfoPtrOutput {
	return o
}

func (o WatchlistUserInfoPtrOutput) ToWatchlistUserInfoPtrOutputWithContext(ctx context.Context) WatchlistUserInfoPtrOutput {
	return o
}

func (o WatchlistUserInfoPtrOutput) Elem() WatchlistUserInfoOutput {
	return o.ApplyT(func(v *WatchlistUserInfo) WatchlistUserInfo {
		if v != nil {
			return *v
		}
		var ret WatchlistUserInfo
		return ret
	}).(WatchlistUserInfoOutput)
}

func (o WatchlistUserInfoPtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistUserInfo) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

type WatchlistUserInfoResponse struct {
	Email    string  `pulumi:"email"`
	Name     string  `pulumi:"name"`
	ObjectId *string `pulumi:"objectId"`
}

type WatchlistUserInfoResponseOutput struct{ *pulumi.OutputState }

func (WatchlistUserInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WatchlistUserInfoResponse)(nil)).Elem()
}

func (o WatchlistUserInfoResponseOutput) ToWatchlistUserInfoResponseOutput() WatchlistUserInfoResponseOutput {
	return o
}

func (o WatchlistUserInfoResponseOutput) ToWatchlistUserInfoResponseOutputWithContext(ctx context.Context) WatchlistUserInfoResponseOutput {
	return o
}

func (o WatchlistUserInfoResponseOutput) Email() pulumi.StringOutput {
	return o.ApplyT(func(v WatchlistUserInfoResponse) string { return v.Email }).(pulumi.StringOutput)
}

func (o WatchlistUserInfoResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v WatchlistUserInfoResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o WatchlistUserInfoResponseOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WatchlistUserInfoResponse) *string { return v.ObjectId }).(pulumi.StringPtrOutput)
}

type WatchlistUserInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (WatchlistUserInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistUserInfoResponse)(nil)).Elem()
}

func (o WatchlistUserInfoResponsePtrOutput) ToWatchlistUserInfoResponsePtrOutput() WatchlistUserInfoResponsePtrOutput {
	return o
}

func (o WatchlistUserInfoResponsePtrOutput) ToWatchlistUserInfoResponsePtrOutputWithContext(ctx context.Context) WatchlistUserInfoResponsePtrOutput {
	return o
}

func (o WatchlistUserInfoResponsePtrOutput) Elem() WatchlistUserInfoResponseOutput {
	return o.ApplyT(func(v *WatchlistUserInfoResponse) WatchlistUserInfoResponse {
		if v != nil {
			return *v
		}
		var ret WatchlistUserInfoResponse
		return ret
	}).(WatchlistUserInfoResponseOutput)
}

func (o WatchlistUserInfoResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistUserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Email
	}).(pulumi.StringPtrOutput)
}

func (o WatchlistUserInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistUserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o WatchlistUserInfoResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WatchlistUserInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertDetailsOverrideOutput{})
	pulumi.RegisterOutputType(AlertDetailsOverridePtrOutput{})
	pulumi.RegisterOutputType(AlertDetailsOverrideResponseOutput{})
	pulumi.RegisterOutputType(AlertDetailsOverrideResponsePtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorPtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponsePtrOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput{})
	pulumi.RegisterOutputType(ClientInfoResponseOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonPtrOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonResponseOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonResponsePtrOutput{})
	pulumi.RegisterOutputType(EntityMappingOutput{})
	pulumi.RegisterOutputType(EntityMappingArrayOutput{})
	pulumi.RegisterOutputType(EntityMappingResponseOutput{})
	pulumi.RegisterOutputType(EntityMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(FieldMappingOutput{})
	pulumi.RegisterOutputType(FieldMappingArrayOutput{})
	pulumi.RegisterOutputType(FieldMappingResponseOutput{})
	pulumi.RegisterOutputType(FieldMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(GroupingConfigurationOutput{})
	pulumi.RegisterOutputType(GroupingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GroupingConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GroupingConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IncidentAdditionalDataResponseOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationPtrOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IncidentInfoOutput{})
	pulumi.RegisterOutputType(IncidentInfoPtrOutput{})
	pulumi.RegisterOutputType(IncidentInfoResponseOutput{})
	pulumi.RegisterOutputType(IncidentInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(IncidentLabelOutput{})
	pulumi.RegisterOutputType(IncidentLabelArrayOutput{})
	pulumi.RegisterOutputType(IncidentLabelResponseOutput{})
	pulumi.RegisterOutputType(IncidentLabelResponseArrayOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoPtrOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoResponseOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesExchangeOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesExchangePtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseExchangeOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseExchangePtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseSharePointOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseSharePointPtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseTeamsOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseTeamsPtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesSharePointOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesSharePointPtrOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesTeamsOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesTeamsPtrOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceArrayOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesIndicatorsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesIndicatorsPtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseIndicatorsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseIndicatorsPtrOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceExternalReferenceOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceExternalReferenceArrayOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceGranularMarkingModelOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceGranularMarkingModelArrayOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceKillChainPhaseOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceKillChainPhaseArrayOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceParsedPatternOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceParsedPatternArrayOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceParsedPatternTypeValueOutput{})
	pulumi.RegisterOutputType(ThreatIntelligenceParsedPatternTypeValueArrayOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoPtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponseOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponsePtrOutput{})
}
