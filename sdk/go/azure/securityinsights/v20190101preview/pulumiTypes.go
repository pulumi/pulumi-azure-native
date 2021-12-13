


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivityTimelineItemResponse struct {
	BucketEndTimeUTC     string `pulumi:"bucketEndTimeUTC"`
	BucketStartTimeUTC   string `pulumi:"bucketStartTimeUTC"`
	Content              string `pulumi:"content"`
	FirstActivityTimeUTC string `pulumi:"firstActivityTimeUTC"`
	Kind                 string `pulumi:"kind"`
	LastActivityTimeUTC  string `pulumi:"lastActivityTimeUTC"`
	QueryId              string `pulumi:"queryId"`
	Title                string `pulumi:"title"`
}

type AlertsDataTypeOfDataConnector struct {
	Alerts AlertsDataTypeOfDataConnectorAlerts `pulumi:"alerts"`
}





type AlertsDataTypeOfDataConnectorInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorOutput() AlertsDataTypeOfDataConnectorOutput
	ToAlertsDataTypeOfDataConnectorOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorOutput
}

type AlertsDataTypeOfDataConnectorArgs struct {
	Alerts AlertsDataTypeOfDataConnectorAlertsInput `pulumi:"alerts"`
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

func (o AlertsDataTypeOfDataConnectorOutput) Alerts() AlertsDataTypeOfDataConnectorAlertsOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnector) AlertsDataTypeOfDataConnectorAlerts { return v.Alerts }).(AlertsDataTypeOfDataConnectorAlertsOutput)
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

func (o AlertsDataTypeOfDataConnectorPtrOutput) Alerts() AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnector) *AlertsDataTypeOfDataConnectorAlerts {
		if v == nil {
			return nil
		}
		return &v.Alerts
	}).(AlertsDataTypeOfDataConnectorAlertsPtrOutput)
}

type AlertsDataTypeOfDataConnectorAlerts struct {
	State string `pulumi:"state"`
}





type AlertsDataTypeOfDataConnectorAlertsInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorAlertsOutput() AlertsDataTypeOfDataConnectorAlertsOutput
	ToAlertsDataTypeOfDataConnectorAlertsOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorAlertsOutput
}

type AlertsDataTypeOfDataConnectorAlertsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (AlertsDataTypeOfDataConnectorAlertsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnectorAlerts)(nil)).Elem()
}

func (i AlertsDataTypeOfDataConnectorAlertsArgs) ToAlertsDataTypeOfDataConnectorAlertsOutput() AlertsDataTypeOfDataConnectorAlertsOutput {
	return i.ToAlertsDataTypeOfDataConnectorAlertsOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorAlertsArgs) ToAlertsDataTypeOfDataConnectorAlertsOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorAlertsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorAlertsOutput)
}

func (i AlertsDataTypeOfDataConnectorAlertsArgs) ToAlertsDataTypeOfDataConnectorAlertsPtrOutput() AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorAlertsArgs) ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorAlertsOutput).ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(ctx)
}









type AlertsDataTypeOfDataConnectorAlertsPtrInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorAlertsPtrOutput() AlertsDataTypeOfDataConnectorAlertsPtrOutput
	ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorAlertsPtrOutput
}

type alertsDataTypeOfDataConnectorAlertsPtrType AlertsDataTypeOfDataConnectorAlertsArgs

func AlertsDataTypeOfDataConnectorAlertsPtr(v *AlertsDataTypeOfDataConnectorAlertsArgs) AlertsDataTypeOfDataConnectorAlertsPtrInput {
	return (*alertsDataTypeOfDataConnectorAlertsPtrType)(v)
}

func (*alertsDataTypeOfDataConnectorAlertsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnectorAlerts)(nil)).Elem()
}

func (i *alertsDataTypeOfDataConnectorAlertsPtrType) ToAlertsDataTypeOfDataConnectorAlertsPtrOutput() AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(context.Background())
}

func (i *alertsDataTypeOfDataConnectorAlertsPtrType) ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorAlertsPtrOutput)
}

type AlertsDataTypeOfDataConnectorAlertsOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorAlertsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnectorAlerts)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorAlertsOutput) ToAlertsDataTypeOfDataConnectorAlertsOutput() AlertsDataTypeOfDataConnectorAlertsOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorAlertsOutput) ToAlertsDataTypeOfDataConnectorAlertsOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorAlertsOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorAlertsOutput) ToAlertsDataTypeOfDataConnectorAlertsPtrOutput() AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return o.ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(context.Background())
}

func (o AlertsDataTypeOfDataConnectorAlertsOutput) ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertsDataTypeOfDataConnectorAlerts) *AlertsDataTypeOfDataConnectorAlerts {
		return &v
	}).(AlertsDataTypeOfDataConnectorAlertsPtrOutput)
}

func (o AlertsDataTypeOfDataConnectorAlertsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnectorAlerts) string { return v.State }).(pulumi.StringOutput)
}

type AlertsDataTypeOfDataConnectorAlertsPtrOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorAlertsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnectorAlerts)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorAlertsPtrOutput) ToAlertsDataTypeOfDataConnectorAlertsPtrOutput() AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorAlertsPtrOutput) ToAlertsDataTypeOfDataConnectorAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorAlertsPtrOutput) Elem() AlertsDataTypeOfDataConnectorAlertsOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorAlerts) AlertsDataTypeOfDataConnectorAlerts {
		if v != nil {
			return *v
		}
		var ret AlertsDataTypeOfDataConnectorAlerts
		return ret
	}).(AlertsDataTypeOfDataConnectorAlertsOutput)
}

func (o AlertsDataTypeOfDataConnectorAlertsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorAlerts) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type AlertsDataTypeOfDataConnectorResponse struct {
	Alerts AlertsDataTypeOfDataConnectorResponseAlerts `pulumi:"alerts"`
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

func (o AlertsDataTypeOfDataConnectorResponseOutput) Alerts() AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnectorResponse) AlertsDataTypeOfDataConnectorResponseAlerts {
		return v.Alerts
	}).(AlertsDataTypeOfDataConnectorResponseAlertsOutput)
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

func (o AlertsDataTypeOfDataConnectorResponsePtrOutput) Alerts() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorResponse) *AlertsDataTypeOfDataConnectorResponseAlerts {
		if v == nil {
			return nil
		}
		return &v.Alerts
	}).(AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput)
}

type AlertsDataTypeOfDataConnectorResponseAlerts struct {
	State string `pulumi:"state"`
}

type AlertsDataTypeOfDataConnectorResponseAlertsOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorResponseAlertsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponseAlerts)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsOutput) ToAlertsDataTypeOfDataConnectorResponseAlertsOutput() AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsOutput) ToAlertsDataTypeOfDataConnectorResponseAlertsOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnectorResponseAlerts) string { return v.State }).(pulumi.StringOutput)
}

type AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput struct{ *pulumi.OutputState }

func (AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnectorResponseAlerts)(nil)).Elem()
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutput() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return o
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput) Elem() AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorResponseAlerts) AlertsDataTypeOfDataConnectorResponseAlerts {
		if v != nil {
			return *v
		}
		var ret AlertsDataTypeOfDataConnectorResponseAlerts
		return ret
	}).(AlertsDataTypeOfDataConnectorResponseAlertsOutput)
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertsDataTypeOfDataConnectorResponseAlerts) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type AutomationRuleModifyPropertiesAction struct {
	ActionConfiguration AutomationRuleModifyPropertiesActionActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                                  `pulumi:"actionType"`
	Order               int                                                     `pulumi:"order"`
}

type AutomationRuleModifyPropertiesActionActionConfiguration struct {
	Classification        *string            `pulumi:"classification"`
	ClassificationComment *string            `pulumi:"classificationComment"`
	ClassificationReason  *string            `pulumi:"classificationReason"`
	Labels                []IncidentLabel    `pulumi:"labels"`
	Owner                 *IncidentOwnerInfo `pulumi:"owner"`
	Severity              *string            `pulumi:"severity"`
	Status                *string            `pulumi:"status"`
}

type AutomationRuleModifyPropertiesActionResponse struct {
	ActionConfiguration AutomationRuleModifyPropertiesActionResponseActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                                          `pulumi:"actionType"`
	Order               int                                                             `pulumi:"order"`
}

type AutomationRuleModifyPropertiesActionResponseActionConfiguration struct {
	Classification        *string                    `pulumi:"classification"`
	ClassificationComment *string                    `pulumi:"classificationComment"`
	ClassificationReason  *string                    `pulumi:"classificationReason"`
	Labels                []IncidentLabelResponse    `pulumi:"labels"`
	Owner                 *IncidentOwnerInfoResponse `pulumi:"owner"`
	Severity              *string                    `pulumi:"severity"`
	Status                *string                    `pulumi:"status"`
}

type AutomationRulePropertyValuesCondition struct {
	ConditionProperties AutomationRulePropertyValuesConditionConditionProperties `pulumi:"conditionProperties"`
	ConditionType       string                                                   `pulumi:"conditionType"`
}





type AutomationRulePropertyValuesConditionInput interface {
	pulumi.Input

	ToAutomationRulePropertyValuesConditionOutput() AutomationRulePropertyValuesConditionOutput
	ToAutomationRulePropertyValuesConditionOutputWithContext(context.Context) AutomationRulePropertyValuesConditionOutput
}

type AutomationRulePropertyValuesConditionArgs struct {
	ConditionProperties AutomationRulePropertyValuesConditionConditionPropertiesInput `pulumi:"conditionProperties"`
	ConditionType       pulumi.StringInput                                            `pulumi:"conditionType"`
}

func (AutomationRulePropertyValuesConditionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesCondition)(nil)).Elem()
}

func (i AutomationRulePropertyValuesConditionArgs) ToAutomationRulePropertyValuesConditionOutput() AutomationRulePropertyValuesConditionOutput {
	return i.ToAutomationRulePropertyValuesConditionOutputWithContext(context.Background())
}

func (i AutomationRulePropertyValuesConditionArgs) ToAutomationRulePropertyValuesConditionOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRulePropertyValuesConditionOutput)
}





type AutomationRulePropertyValuesConditionArrayInput interface {
	pulumi.Input

	ToAutomationRulePropertyValuesConditionArrayOutput() AutomationRulePropertyValuesConditionArrayOutput
	ToAutomationRulePropertyValuesConditionArrayOutputWithContext(context.Context) AutomationRulePropertyValuesConditionArrayOutput
}

type AutomationRulePropertyValuesConditionArray []AutomationRulePropertyValuesConditionInput

func (AutomationRulePropertyValuesConditionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRulePropertyValuesCondition)(nil)).Elem()
}

func (i AutomationRulePropertyValuesConditionArray) ToAutomationRulePropertyValuesConditionArrayOutput() AutomationRulePropertyValuesConditionArrayOutput {
	return i.ToAutomationRulePropertyValuesConditionArrayOutputWithContext(context.Background())
}

func (i AutomationRulePropertyValuesConditionArray) ToAutomationRulePropertyValuesConditionArrayOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRulePropertyValuesConditionArrayOutput)
}

type AutomationRulePropertyValuesConditionOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyValuesConditionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesCondition)(nil)).Elem()
}

func (o AutomationRulePropertyValuesConditionOutput) ToAutomationRulePropertyValuesConditionOutput() AutomationRulePropertyValuesConditionOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionOutput) ToAutomationRulePropertyValuesConditionOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionOutput) ConditionProperties() AutomationRulePropertyValuesConditionConditionPropertiesOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesCondition) AutomationRulePropertyValuesConditionConditionProperties {
		return v.ConditionProperties
	}).(AutomationRulePropertyValuesConditionConditionPropertiesOutput)
}

func (o AutomationRulePropertyValuesConditionOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesCondition) string { return v.ConditionType }).(pulumi.StringOutput)
}

type AutomationRulePropertyValuesConditionArrayOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyValuesConditionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRulePropertyValuesCondition)(nil)).Elem()
}

func (o AutomationRulePropertyValuesConditionArrayOutput) ToAutomationRulePropertyValuesConditionArrayOutput() AutomationRulePropertyValuesConditionArrayOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionArrayOutput) ToAutomationRulePropertyValuesConditionArrayOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionArrayOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionArrayOutput) Index(i pulumi.IntInput) AutomationRulePropertyValuesConditionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRulePropertyValuesCondition {
		return vs[0].([]AutomationRulePropertyValuesCondition)[vs[1].(int)]
	}).(AutomationRulePropertyValuesConditionOutput)
}

type AutomationRulePropertyValuesConditionConditionProperties struct {
	Operator       *string  `pulumi:"operator"`
	PropertyName   *string  `pulumi:"propertyName"`
	PropertyValues []string `pulumi:"propertyValues"`
}





type AutomationRulePropertyValuesConditionConditionPropertiesInput interface {
	pulumi.Input

	ToAutomationRulePropertyValuesConditionConditionPropertiesOutput() AutomationRulePropertyValuesConditionConditionPropertiesOutput
	ToAutomationRulePropertyValuesConditionConditionPropertiesOutputWithContext(context.Context) AutomationRulePropertyValuesConditionConditionPropertiesOutput
}

type AutomationRulePropertyValuesConditionConditionPropertiesArgs struct {
	Operator       pulumi.StringPtrInput   `pulumi:"operator"`
	PropertyName   pulumi.StringPtrInput   `pulumi:"propertyName"`
	PropertyValues pulumi.StringArrayInput `pulumi:"propertyValues"`
}

func (AutomationRulePropertyValuesConditionConditionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesConditionConditionProperties)(nil)).Elem()
}

func (i AutomationRulePropertyValuesConditionConditionPropertiesArgs) ToAutomationRulePropertyValuesConditionConditionPropertiesOutput() AutomationRulePropertyValuesConditionConditionPropertiesOutput {
	return i.ToAutomationRulePropertyValuesConditionConditionPropertiesOutputWithContext(context.Background())
}

func (i AutomationRulePropertyValuesConditionConditionPropertiesArgs) ToAutomationRulePropertyValuesConditionConditionPropertiesOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionConditionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRulePropertyValuesConditionConditionPropertiesOutput)
}

type AutomationRulePropertyValuesConditionConditionPropertiesOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyValuesConditionConditionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesConditionConditionProperties)(nil)).Elem()
}

func (o AutomationRulePropertyValuesConditionConditionPropertiesOutput) ToAutomationRulePropertyValuesConditionConditionPropertiesOutput() AutomationRulePropertyValuesConditionConditionPropertiesOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionConditionPropertiesOutput) ToAutomationRulePropertyValuesConditionConditionPropertiesOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionConditionPropertiesOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionConditionPropertiesOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionConditionProperties) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o AutomationRulePropertyValuesConditionConditionPropertiesOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionConditionProperties) *string { return v.PropertyName }).(pulumi.StringPtrOutput)
}

func (o AutomationRulePropertyValuesConditionConditionPropertiesOutput) PropertyValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionConditionProperties) []string { return v.PropertyValues }).(pulumi.StringArrayOutput)
}

type AutomationRulePropertyValuesConditionResponse struct {
	ConditionProperties AutomationRulePropertyValuesConditionResponseConditionProperties `pulumi:"conditionProperties"`
	ConditionType       string                                                           `pulumi:"conditionType"`
}

type AutomationRulePropertyValuesConditionResponseOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyValuesConditionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesConditionResponse)(nil)).Elem()
}

func (o AutomationRulePropertyValuesConditionResponseOutput) ToAutomationRulePropertyValuesConditionResponseOutput() AutomationRulePropertyValuesConditionResponseOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionResponseOutput) ToAutomationRulePropertyValuesConditionResponseOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionResponseOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionResponseOutput) ConditionProperties() AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionResponse) AutomationRulePropertyValuesConditionResponseConditionProperties {
		return v.ConditionProperties
	}).(AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput)
}

func (o AutomationRulePropertyValuesConditionResponseOutput) ConditionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionResponse) string { return v.ConditionType }).(pulumi.StringOutput)
}

type AutomationRulePropertyValuesConditionResponseArrayOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyValuesConditionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRulePropertyValuesConditionResponse)(nil)).Elem()
}

func (o AutomationRulePropertyValuesConditionResponseArrayOutput) ToAutomationRulePropertyValuesConditionResponseArrayOutput() AutomationRulePropertyValuesConditionResponseArrayOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionResponseArrayOutput) ToAutomationRulePropertyValuesConditionResponseArrayOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionResponseArrayOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionResponseArrayOutput) Index(i pulumi.IntInput) AutomationRulePropertyValuesConditionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AutomationRulePropertyValuesConditionResponse {
		return vs[0].([]AutomationRulePropertyValuesConditionResponse)[vs[1].(int)]
	}).(AutomationRulePropertyValuesConditionResponseOutput)
}

type AutomationRulePropertyValuesConditionResponseConditionProperties struct {
	Operator       *string  `pulumi:"operator"`
	PropertyName   *string  `pulumi:"propertyName"`
	PropertyValues []string `pulumi:"propertyValues"`
}

type AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput struct{ *pulumi.OutputState }

func (AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesConditionResponseConditionProperties)(nil)).Elem()
}

func (o AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput) ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutput() AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput) ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput {
	return o
}

func (o AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionResponseConditionProperties) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput) PropertyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionResponseConditionProperties) *string {
		return v.PropertyName
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput) PropertyValues() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AutomationRulePropertyValuesConditionResponseConditionProperties) []string {
		return v.PropertyValues
	}).(pulumi.StringArrayOutput)
}

type AutomationRuleRunPlaybookAction struct {
	ActionConfiguration AutomationRuleRunPlaybookActionActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                             `pulumi:"actionType"`
	Order               int                                                `pulumi:"order"`
}

type AutomationRuleRunPlaybookActionActionConfiguration struct {
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}

type AutomationRuleRunPlaybookActionResponse struct {
	ActionConfiguration AutomationRuleRunPlaybookActionResponseActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                                     `pulumi:"actionType"`
	Order               int                                                        `pulumi:"order"`
}

type AutomationRuleRunPlaybookActionResponseActionConfiguration struct {
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}

type AutomationRuleTriggeringLogic struct {
	Conditions        []AutomationRulePropertyValuesCondition `pulumi:"conditions"`
	ExpirationTimeUtc *string                                 `pulumi:"expirationTimeUtc"`
	IsEnabled         bool                                    `pulumi:"isEnabled"`
	TriggersOn        string                                  `pulumi:"triggersOn"`
	TriggersWhen      string                                  `pulumi:"triggersWhen"`
}





type AutomationRuleTriggeringLogicInput interface {
	pulumi.Input

	ToAutomationRuleTriggeringLogicOutput() AutomationRuleTriggeringLogicOutput
	ToAutomationRuleTriggeringLogicOutputWithContext(context.Context) AutomationRuleTriggeringLogicOutput
}

type AutomationRuleTriggeringLogicArgs struct {
	Conditions        AutomationRulePropertyValuesConditionArrayInput `pulumi:"conditions"`
	ExpirationTimeUtc pulumi.StringPtrInput                           `pulumi:"expirationTimeUtc"`
	IsEnabled         pulumi.BoolInput                                `pulumi:"isEnabled"`
	TriggersOn        pulumi.StringInput                              `pulumi:"triggersOn"`
	TriggersWhen      pulumi.StringInput                              `pulumi:"triggersWhen"`
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

func (o AutomationRuleTriggeringLogicOutput) Conditions() AutomationRulePropertyValuesConditionArrayOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogic) []AutomationRulePropertyValuesCondition { return v.Conditions }).(AutomationRulePropertyValuesConditionArrayOutput)
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
	Conditions        []AutomationRulePropertyValuesConditionResponse `pulumi:"conditions"`
	ExpirationTimeUtc *string                                         `pulumi:"expirationTimeUtc"`
	IsEnabled         bool                                            `pulumi:"isEnabled"`
	TriggersOn        string                                          `pulumi:"triggersOn"`
	TriggersWhen      string                                          `pulumi:"triggersWhen"`
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

func (o AutomationRuleTriggeringLogicResponseOutput) Conditions() AutomationRulePropertyValuesConditionResponseArrayOutput {
	return o.ApplyT(func(v AutomationRuleTriggeringLogicResponse) []AutomationRulePropertyValuesConditionResponse {
		return v.Conditions
	}).(AutomationRulePropertyValuesConditionResponseArrayOutput)
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
	Logs AwsCloudTrailDataConnectorDataTypesLogs `pulumi:"logs"`
}





type AwsCloudTrailDataConnectorDataTypesInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesOutput() AwsCloudTrailDataConnectorDataTypesOutput
	ToAwsCloudTrailDataConnectorDataTypesOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesOutput
}

type AwsCloudTrailDataConnectorDataTypesArgs struct {
	Logs AwsCloudTrailDataConnectorDataTypesLogsInput `pulumi:"logs"`
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

func (o AwsCloudTrailDataConnectorDataTypesOutput) Logs() AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypes) AwsCloudTrailDataConnectorDataTypesLogs { return v.Logs }).(AwsCloudTrailDataConnectorDataTypesLogsOutput)
}

type AwsCloudTrailDataConnectorDataTypesLogs struct {
	State string `pulumi:"state"`
}





type AwsCloudTrailDataConnectorDataTypesLogsInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesLogsOutput() AwsCloudTrailDataConnectorDataTypesLogsOutput
	ToAwsCloudTrailDataConnectorDataTypesLogsOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesLogsOutput
}

type AwsCloudTrailDataConnectorDataTypesLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
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

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesLogs) string { return v.State }).(pulumi.StringOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponse struct {
	Logs AwsCloudTrailDataConnectorDataTypesResponseLogs `pulumi:"logs"`
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

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) Logs() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesResponse) AwsCloudTrailDataConnectorDataTypesResponseLogs {
		return v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponseLogs struct {
	State string `pulumi:"state"`
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

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesResponseLogs) string { return v.State }).(pulumi.StringOutput)
}

type BookmarkTimelineItemResponse struct {
	AzureResourceId string            `pulumi:"azureResourceId"`
	CreatedBy       *UserInfoResponse `pulumi:"createdBy"`
	DisplayName     *string           `pulumi:"displayName"`
	EndTimeUtc      *string           `pulumi:"endTimeUtc"`
	EventTime       *string           `pulumi:"eventTime"`
	Kind            string            `pulumi:"kind"`
	Labels          []string          `pulumi:"labels"`
	Notes           *string           `pulumi:"notes"`
	StartTimeUtc    *string           `pulumi:"startTimeUtc"`
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

type Dynamics365DataConnectorDataTypes struct {
	Dynamics365CdsActivities Dynamics365DataConnectorDataTypesDynamics365CdsActivities `pulumi:"dynamics365CdsActivities"`
}





type Dynamics365DataConnectorDataTypesInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesOutput() Dynamics365DataConnectorDataTypesOutput
	ToDynamics365DataConnectorDataTypesOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesOutput
}

type Dynamics365DataConnectorDataTypesArgs struct {
	Dynamics365CdsActivities Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesInput `pulumi:"dynamics365CdsActivities"`
}

func (Dynamics365DataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypes)(nil)).Elem()
}

func (i Dynamics365DataConnectorDataTypesArgs) ToDynamics365DataConnectorDataTypesOutput() Dynamics365DataConnectorDataTypesOutput {
	return i.ToDynamics365DataConnectorDataTypesOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesArgs) ToDynamics365DataConnectorDataTypesOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesOutput)
}

type Dynamics365DataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypes)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesOutput) ToDynamics365DataConnectorDataTypesOutput() Dynamics365DataConnectorDataTypesOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesOutput) ToDynamics365DataConnectorDataTypesOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesOutput) Dynamics365CdsActivities() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypes) Dynamics365DataConnectorDataTypesDynamics365CdsActivities {
		return v.Dynamics365CdsActivities
	}).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput)
}

type Dynamics365DataConnectorDataTypesDynamics365CdsActivities struct {
	State string `pulumi:"state"`
}





type Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput
	ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput
}

type Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypesDynamics365CdsActivities)(nil)).Elem()
}

func (i Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return i.ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput)
}

type Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypesDynamics365CdsActivities)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypesDynamics365CdsActivities) string { return v.State }).(pulumi.StringOutput)
}

type Dynamics365DataConnectorDataTypesResponse struct {
	Dynamics365CdsActivities Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities `pulumi:"dynamics365CdsActivities"`
}

type Dynamics365DataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponse)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesResponseOutput) ToDynamics365DataConnectorDataTypesResponseOutput() Dynamics365DataConnectorDataTypesResponseOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponseOutput) ToDynamics365DataConnectorDataTypesResponseOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponseOutput) Dynamics365CdsActivities() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypesResponse) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities {
		return v.Dynamics365CdsActivities
	}).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput)
}

type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities struct {
	State string `pulumi:"state"`
}

type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities) string { return v.State }).(pulumi.StringOutput)
}

type EntityInsightItemResponse struct {
	ChartQueryResults []InsightsTableResultResponse               `pulumi:"chartQueryResults"`
	QueryId           *string                                     `pulumi:"queryId"`
	QueryTimeInterval *EntityInsightItemResponseQueryTimeInterval `pulumi:"queryTimeInterval"`
	TableQueryResults *InsightsTableResultResponse                `pulumi:"tableQueryResults"`
}

type EntityInsightItemResponseQueryTimeInterval struct {
	EndTime   *string `pulumi:"endTime"`
	StartTime *string `pulumi:"startTime"`
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

type GetInsightsErrorResponse struct {
	ErrorMessage string  `pulumi:"errorMessage"`
	Kind         string  `pulumi:"kind"`
	QueryId      *string `pulumi:"queryId"`
}

type GetInsightsResultsMetadataResponse struct {
	Errors     []GetInsightsErrorResponse `pulumi:"errors"`
	TotalCount int                        `pulumi:"totalCount"`
}

type GroupingConfiguration struct {
	Enabled                bool     `pulumi:"enabled"`
	EntitiesMatchingMethod string   `pulumi:"entitiesMatchingMethod"`
	GroupByEntities        []string `pulumi:"groupByEntities"`
	LookbackDuration       string   `pulumi:"lookbackDuration"`
	ReopenClosedIncident   bool     `pulumi:"reopenClosedIncident"`
}





type GroupingConfigurationInput interface {
	pulumi.Input

	ToGroupingConfigurationOutput() GroupingConfigurationOutput
	ToGroupingConfigurationOutputWithContext(context.Context) GroupingConfigurationOutput
}

type GroupingConfigurationArgs struct {
	Enabled                pulumi.BoolInput        `pulumi:"enabled"`
	EntitiesMatchingMethod pulumi.StringInput      `pulumi:"entitiesMatchingMethod"`
	GroupByEntities        pulumi.StringArrayInput `pulumi:"groupByEntities"`
	LookbackDuration       pulumi.StringInput      `pulumi:"lookbackDuration"`
	ReopenClosedIncident   pulumi.BoolInput        `pulumi:"reopenClosedIncident"`
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

func (o GroupingConfigurationOutput) EntitiesMatchingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfiguration) string { return v.EntitiesMatchingMethod }).(pulumi.StringOutput)
}

func (o GroupingConfigurationOutput) GroupByEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfiguration) []string { return v.GroupByEntities }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationOutput) LookbackDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfiguration) string { return v.LookbackDuration }).(pulumi.StringOutput)
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

func (o GroupingConfigurationPtrOutput) EntitiesMatchingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingConfiguration) *string {
		if v == nil {
			return nil
		}
		return &v.EntitiesMatchingMethod
	}).(pulumi.StringPtrOutput)
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

func (o GroupingConfigurationPtrOutput) ReopenClosedIncident() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GroupingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return &v.ReopenClosedIncident
	}).(pulumi.BoolPtrOutput)
}

type GroupingConfigurationResponse struct {
	Enabled                bool     `pulumi:"enabled"`
	EntitiesMatchingMethod string   `pulumi:"entitiesMatchingMethod"`
	GroupByEntities        []string `pulumi:"groupByEntities"`
	LookbackDuration       string   `pulumi:"lookbackDuration"`
	ReopenClosedIncident   bool     `pulumi:"reopenClosedIncident"`
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

func (o GroupingConfigurationResponseOutput) EntitiesMatchingMethod() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) string { return v.EntitiesMatchingMethod }).(pulumi.StringOutput)
}

func (o GroupingConfigurationResponseOutput) GroupByEntities() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) []string { return v.GroupByEntities }).(pulumi.StringArrayOutput)
}

func (o GroupingConfigurationResponseOutput) LookbackDuration() pulumi.StringOutput {
	return o.ApplyT(func(v GroupingConfigurationResponse) string { return v.LookbackDuration }).(pulumi.StringOutput)
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

func (o GroupingConfigurationResponsePtrOutput) EntitiesMatchingMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GroupingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EntitiesMatchingMethod
	}).(pulumi.StringPtrOutput)
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

func (o IncidentOwnerInfoResponsePtrOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IncidentOwnerInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPrincipalName
	}).(pulumi.StringPtrOutput)
}

type InsightsTableResultResponse struct {
	Columns []InsightsTableResultResponseColumns `pulumi:"columns"`
	Rows    [][]string                           `pulumi:"rows"`
}

type InsightsTableResultResponseColumns struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type MCASDataConnectorDataTypes struct {
	Alerts        AlertsDataTypeOfDataConnectorAlerts      `pulumi:"alerts"`
	DiscoveryLogs *MCASDataConnectorDataTypesDiscoveryLogs `pulumi:"discoveryLogs"`
}





type MCASDataConnectorDataTypesInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesOutput() MCASDataConnectorDataTypesOutput
	ToMCASDataConnectorDataTypesOutputWithContext(context.Context) MCASDataConnectorDataTypesOutput
}

type MCASDataConnectorDataTypesArgs struct {
	Alerts        AlertsDataTypeOfDataConnectorAlertsInput        `pulumi:"alerts"`
	DiscoveryLogs MCASDataConnectorDataTypesDiscoveryLogsPtrInput `pulumi:"discoveryLogs"`
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

func (o MCASDataConnectorDataTypesOutput) Alerts() AlertsDataTypeOfDataConnectorAlertsOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) AlertsDataTypeOfDataConnectorAlerts { return v.Alerts }).(AlertsDataTypeOfDataConnectorAlertsOutput)
}

func (o MCASDataConnectorDataTypesOutput) DiscoveryLogs() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) *MCASDataConnectorDataTypesDiscoveryLogs { return v.DiscoveryLogs }).(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput)
}

type MCASDataConnectorDataTypesDiscoveryLogs struct {
	State string `pulumi:"state"`
}





type MCASDataConnectorDataTypesDiscoveryLogsInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesDiscoveryLogsOutput() MCASDataConnectorDataTypesDiscoveryLogsOutput
	ToMCASDataConnectorDataTypesDiscoveryLogsOutputWithContext(context.Context) MCASDataConnectorDataTypesDiscoveryLogsOutput
}

type MCASDataConnectorDataTypesDiscoveryLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (MCASDataConnectorDataTypesDiscoveryLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypesDiscoveryLogs)(nil)).Elem()
}

func (i MCASDataConnectorDataTypesDiscoveryLogsArgs) ToMCASDataConnectorDataTypesDiscoveryLogsOutput() MCASDataConnectorDataTypesDiscoveryLogsOutput {
	return i.ToMCASDataConnectorDataTypesDiscoveryLogsOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesDiscoveryLogsArgs) ToMCASDataConnectorDataTypesDiscoveryLogsOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesDiscoveryLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesDiscoveryLogsOutput)
}

func (i MCASDataConnectorDataTypesDiscoveryLogsArgs) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return i.ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesDiscoveryLogsArgs) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesDiscoveryLogsOutput).ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(ctx)
}









type MCASDataConnectorDataTypesDiscoveryLogsPtrInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput
	ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(context.Context) MCASDataConnectorDataTypesDiscoveryLogsPtrOutput
}

type mcasdataConnectorDataTypesDiscoveryLogsPtrType MCASDataConnectorDataTypesDiscoveryLogsArgs

func MCASDataConnectorDataTypesDiscoveryLogsPtr(v *MCASDataConnectorDataTypesDiscoveryLogsArgs) MCASDataConnectorDataTypesDiscoveryLogsPtrInput {
	return (*mcasdataConnectorDataTypesDiscoveryLogsPtrType)(v)
}

func (*mcasdataConnectorDataTypesDiscoveryLogsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypesDiscoveryLogs)(nil)).Elem()
}

func (i *mcasdataConnectorDataTypesDiscoveryLogsPtrType) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return i.ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(context.Background())
}

func (i *mcasdataConnectorDataTypesDiscoveryLogsPtrType) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput)
}

type MCASDataConnectorDataTypesDiscoveryLogsOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesDiscoveryLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypesDiscoveryLogs)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesDiscoveryLogsOutput) ToMCASDataConnectorDataTypesDiscoveryLogsOutput() MCASDataConnectorDataTypesDiscoveryLogsOutput {
	return o
}

func (o MCASDataConnectorDataTypesDiscoveryLogsOutput) ToMCASDataConnectorDataTypesDiscoveryLogsOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesDiscoveryLogsOutput {
	return o
}

func (o MCASDataConnectorDataTypesDiscoveryLogsOutput) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o.ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(context.Background())
}

func (o MCASDataConnectorDataTypesDiscoveryLogsOutput) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MCASDataConnectorDataTypesDiscoveryLogs) *MCASDataConnectorDataTypesDiscoveryLogs {
		return &v
	}).(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput)
}

func (o MCASDataConnectorDataTypesDiscoveryLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesDiscoveryLogs) string { return v.State }).(pulumi.StringOutput)
}

type MCASDataConnectorDataTypesDiscoveryLogsPtrOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesDiscoveryLogsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypesDiscoveryLogs)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesDiscoveryLogsPtrOutput) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesDiscoveryLogsPtrOutput) ToMCASDataConnectorDataTypesDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesDiscoveryLogsPtrOutput) Elem() MCASDataConnectorDataTypesDiscoveryLogsOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesDiscoveryLogs) MCASDataConnectorDataTypesDiscoveryLogs {
		if v != nil {
			return *v
		}
		var ret MCASDataConnectorDataTypesDiscoveryLogs
		return ret
	}).(MCASDataConnectorDataTypesDiscoveryLogsOutput)
}

func (o MCASDataConnectorDataTypesDiscoveryLogsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesDiscoveryLogs) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MCASDataConnectorDataTypesResponse struct {
	Alerts        AlertsDataTypeOfDataConnectorResponseAlerts      `pulumi:"alerts"`
	DiscoveryLogs *MCASDataConnectorDataTypesResponseDiscoveryLogs `pulumi:"discoveryLogs"`
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

func (o MCASDataConnectorDataTypesResponseOutput) Alerts() AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) AlertsDataTypeOfDataConnectorResponseAlerts {
		return v.Alerts
	}).(AlertsDataTypeOfDataConnectorResponseAlertsOutput)
}

func (o MCASDataConnectorDataTypesResponseOutput) DiscoveryLogs() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) *MCASDataConnectorDataTypesResponseDiscoveryLogs {
		return v.DiscoveryLogs
	}).(MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput)
}

type MCASDataConnectorDataTypesResponseDiscoveryLogs struct {
	State string `pulumi:"state"`
}

type MCASDataConnectorDataTypesResponseDiscoveryLogsOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesResponseDiscoveryLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypesResponseDiscoveryLogs)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsOutput) ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsOutput) ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponseDiscoveryLogs) string { return v.State }).(pulumi.StringOutput)
}

type MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput struct{ *pulumi.OutputState }

func (MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypesResponseDiscoveryLogs)(nil)).Elem()
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return o
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput) Elem() MCASDataConnectorDataTypesResponseDiscoveryLogsOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponseDiscoveryLogs) MCASDataConnectorDataTypesResponseDiscoveryLogs {
		if v != nil {
			return *v
		}
		var ret MCASDataConnectorDataTypesResponseDiscoveryLogs
		return ret
	}).(MCASDataConnectorDataTypesResponseDiscoveryLogsOutput)
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponseDiscoveryLogs) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MSTIDataConnectorDataTypes struct {
	BingSafetyPhishingURL       MSTIDataConnectorDataTypesBingSafetyPhishingURL       `pulumi:"bingSafetyPhishingURL"`
	MicrosoftEmergingThreatFeed MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed `pulumi:"microsoftEmergingThreatFeed"`
}





type MSTIDataConnectorDataTypesInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesOutput() MSTIDataConnectorDataTypesOutput
	ToMSTIDataConnectorDataTypesOutputWithContext(context.Context) MSTIDataConnectorDataTypesOutput
}

type MSTIDataConnectorDataTypesArgs struct {
	BingSafetyPhishingURL       MSTIDataConnectorDataTypesBingSafetyPhishingURLInput       `pulumi:"bingSafetyPhishingURL"`
	MicrosoftEmergingThreatFeed MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedInput `pulumi:"microsoftEmergingThreatFeed"`
}

func (MSTIDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypes)(nil)).Elem()
}

func (i MSTIDataConnectorDataTypesArgs) ToMSTIDataConnectorDataTypesOutput() MSTIDataConnectorDataTypesOutput {
	return i.ToMSTIDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesArgs) ToMSTIDataConnectorDataTypesOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesOutput)
}

type MSTIDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypes)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesOutput) ToMSTIDataConnectorDataTypesOutput() MSTIDataConnectorDataTypesOutput {
	return o
}

func (o MSTIDataConnectorDataTypesOutput) ToMSTIDataConnectorDataTypesOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesOutput {
	return o
}

func (o MSTIDataConnectorDataTypesOutput) BingSafetyPhishingURL() MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypes) MSTIDataConnectorDataTypesBingSafetyPhishingURL {
		return v.BingSafetyPhishingURL
	}).(MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput)
}

func (o MSTIDataConnectorDataTypesOutput) MicrosoftEmergingThreatFeed() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypes) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed {
		return v.MicrosoftEmergingThreatFeed
	}).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput)
}

type MSTIDataConnectorDataTypesBingSafetyPhishingURL struct {
	LookbackPeriod string `pulumi:"lookbackPeriod"`
	State          string `pulumi:"state"`
}





type MSTIDataConnectorDataTypesBingSafetyPhishingURLInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput
	ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutputWithContext(context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput
}

type MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs struct {
	LookbackPeriod pulumi.StringInput `pulumi:"lookbackPeriod"`
	State          pulumi.StringInput `pulumi:"state"`
}

func (MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesBingSafetyPhishingURL)(nil)).Elem()
}

func (i MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput {
	return i.ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput)
}

type MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesBingSafetyPhishingURL)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput {
	return o
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput {
	return o
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesBingSafetyPhishingURL) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesBingSafetyPhishingURL) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed struct {
	LookbackPeriod string `pulumi:"lookbackPeriod"`
	State          string `pulumi:"state"`
}





type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput
	ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutputWithContext(context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput
}

type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs struct {
	LookbackPeriod pulumi.StringInput `pulumi:"lookbackPeriod"`
	State          pulumi.StringInput `pulumi:"state"`
}

func (MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (i MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput {
	return i.ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput)
}

type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput {
	return o
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput {
	return o
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesResponse struct {
	BingSafetyPhishingURL       MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL       `pulumi:"bingSafetyPhishingURL"`
	MicrosoftEmergingThreatFeed MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed `pulumi:"microsoftEmergingThreatFeed"`
}

type MSTIDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesResponseOutput) ToMSTIDataConnectorDataTypesResponseOutput() MSTIDataConnectorDataTypesResponseOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseOutput) ToMSTIDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseOutput) BingSafetyPhishingURL() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponse) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL {
		return v.BingSafetyPhishingURL
	}).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput)
}

func (o MSTIDataConnectorDataTypesResponseOutput) MicrosoftEmergingThreatFeed() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponse) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed {
		return v.MicrosoftEmergingThreatFeed
	}).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput)
}

type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL struct {
	LookbackPeriod string `pulumi:"lookbackPeriod"`
	State          string `pulumi:"state"`
}

type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed struct {
	LookbackPeriod string `pulumi:"lookbackPeriod"`
	State          string `pulumi:"state"`
}

type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) string { return v.State }).(pulumi.StringOutput)
}

type MTPDataConnectorDataTypes struct {
	Incidents MTPDataConnectorDataTypesIncidents `pulumi:"incidents"`
}





type MTPDataConnectorDataTypesInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesOutput() MTPDataConnectorDataTypesOutput
	ToMTPDataConnectorDataTypesOutputWithContext(context.Context) MTPDataConnectorDataTypesOutput
}

type MTPDataConnectorDataTypesArgs struct {
	Incidents MTPDataConnectorDataTypesIncidentsInput `pulumi:"incidents"`
}

func (MTPDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypes)(nil)).Elem()
}

func (i MTPDataConnectorDataTypesArgs) ToMTPDataConnectorDataTypesOutput() MTPDataConnectorDataTypesOutput {
	return i.ToMTPDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesArgs) ToMTPDataConnectorDataTypesOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesOutput)
}

type MTPDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypes)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesOutput) ToMTPDataConnectorDataTypesOutput() MTPDataConnectorDataTypesOutput {
	return o
}

func (o MTPDataConnectorDataTypesOutput) ToMTPDataConnectorDataTypesOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesOutput {
	return o
}

func (o MTPDataConnectorDataTypesOutput) Incidents() MTPDataConnectorDataTypesIncidentsOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypes) MTPDataConnectorDataTypesIncidents { return v.Incidents }).(MTPDataConnectorDataTypesIncidentsOutput)
}

type MTPDataConnectorDataTypesIncidents struct {
	State string `pulumi:"state"`
}





type MTPDataConnectorDataTypesIncidentsInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesIncidentsOutput() MTPDataConnectorDataTypesIncidentsOutput
	ToMTPDataConnectorDataTypesIncidentsOutputWithContext(context.Context) MTPDataConnectorDataTypesIncidentsOutput
}

type MTPDataConnectorDataTypesIncidentsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (MTPDataConnectorDataTypesIncidentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypesIncidents)(nil)).Elem()
}

func (i MTPDataConnectorDataTypesIncidentsArgs) ToMTPDataConnectorDataTypesIncidentsOutput() MTPDataConnectorDataTypesIncidentsOutput {
	return i.ToMTPDataConnectorDataTypesIncidentsOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesIncidentsArgs) ToMTPDataConnectorDataTypesIncidentsOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesIncidentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesIncidentsOutput)
}

type MTPDataConnectorDataTypesIncidentsOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesIncidentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypesIncidents)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesIncidentsOutput) ToMTPDataConnectorDataTypesIncidentsOutput() MTPDataConnectorDataTypesIncidentsOutput {
	return o
}

func (o MTPDataConnectorDataTypesIncidentsOutput) ToMTPDataConnectorDataTypesIncidentsOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesIncidentsOutput {
	return o
}

func (o MTPDataConnectorDataTypesIncidentsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypesIncidents) string { return v.State }).(pulumi.StringOutput)
}

type MTPDataConnectorDataTypesResponse struct {
	Incidents MTPDataConnectorDataTypesResponseIncidents `pulumi:"incidents"`
}

type MTPDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesResponseOutput) ToMTPDataConnectorDataTypesResponseOutput() MTPDataConnectorDataTypesResponseOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponseOutput) ToMTPDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponseOutput) Incidents() MTPDataConnectorDataTypesResponseIncidentsOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypesResponse) MTPDataConnectorDataTypesResponseIncidents {
		return v.Incidents
	}).(MTPDataConnectorDataTypesResponseIncidentsOutput)
}

type MTPDataConnectorDataTypesResponseIncidents struct {
	State string `pulumi:"state"`
}

type MTPDataConnectorDataTypesResponseIncidentsOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesResponseIncidentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypesResponseIncidents)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesResponseIncidentsOutput) ToMTPDataConnectorDataTypesResponseIncidentsOutput() MTPDataConnectorDataTypesResponseIncidentsOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponseIncidentsOutput) ToMTPDataConnectorDataTypesResponseIncidentsOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseIncidentsOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponseIncidentsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypesResponseIncidents) string { return v.State }).(pulumi.StringOutput)
}

type OfficeDataConnectorDataTypes struct {
	Exchange   OfficeDataConnectorDataTypesExchange   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesSharePoint `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesTeams      `pulumi:"teams"`
}





type OfficeDataConnectorDataTypesInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesOutput() OfficeDataConnectorDataTypesOutput
	ToOfficeDataConnectorDataTypesOutputWithContext(context.Context) OfficeDataConnectorDataTypesOutput
}

type OfficeDataConnectorDataTypesArgs struct {
	Exchange   OfficeDataConnectorDataTypesExchangeInput   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesSharePointInput `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesTeamsInput      `pulumi:"teams"`
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

func (o OfficeDataConnectorDataTypesOutput) Exchange() OfficeDataConnectorDataTypesExchangeOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypesExchange { return v.Exchange }).(OfficeDataConnectorDataTypesExchangeOutput)
}

func (o OfficeDataConnectorDataTypesOutput) SharePoint() OfficeDataConnectorDataTypesSharePointOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypesSharePoint { return v.SharePoint }).(OfficeDataConnectorDataTypesSharePointOutput)
}

func (o OfficeDataConnectorDataTypesOutput) Teams() OfficeDataConnectorDataTypesTeamsOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypesTeams { return v.Teams }).(OfficeDataConnectorDataTypesTeamsOutput)
}

type OfficeDataConnectorDataTypesExchange struct {
	State string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesExchangeInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesExchangeOutput() OfficeDataConnectorDataTypesExchangeOutput
	ToOfficeDataConnectorDataTypesExchangeOutputWithContext(context.Context) OfficeDataConnectorDataTypesExchangeOutput
}

type OfficeDataConnectorDataTypesExchangeArgs struct {
	State pulumi.StringInput `pulumi:"state"`
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

func (o OfficeDataConnectorDataTypesExchangeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesExchange) string { return v.State }).(pulumi.StringOutput)
}

type OfficeDataConnectorDataTypesResponse struct {
	Exchange   OfficeDataConnectorDataTypesResponseExchange   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesResponseSharePoint `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesResponseTeams      `pulumi:"teams"`
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

func (o OfficeDataConnectorDataTypesResponseOutput) Exchange() OfficeDataConnectorDataTypesResponseExchangeOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponse) OfficeDataConnectorDataTypesResponseExchange {
		return v.Exchange
	}).(OfficeDataConnectorDataTypesResponseExchangeOutput)
}

func (o OfficeDataConnectorDataTypesResponseOutput) SharePoint() OfficeDataConnectorDataTypesResponseSharePointOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponse) OfficeDataConnectorDataTypesResponseSharePoint {
		return v.SharePoint
	}).(OfficeDataConnectorDataTypesResponseSharePointOutput)
}

func (o OfficeDataConnectorDataTypesResponseOutput) Teams() OfficeDataConnectorDataTypesResponseTeamsOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponse) OfficeDataConnectorDataTypesResponseTeams { return v.Teams }).(OfficeDataConnectorDataTypesResponseTeamsOutput)
}

type OfficeDataConnectorDataTypesResponseExchange struct {
	State string `pulumi:"state"`
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

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseExchange) string { return v.State }).(pulumi.StringOutput)
}

type OfficeDataConnectorDataTypesResponseSharePoint struct {
	State string `pulumi:"state"`
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

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseSharePoint) string { return v.State }).(pulumi.StringOutput)
}

type OfficeDataConnectorDataTypesResponseTeams struct {
	State string `pulumi:"state"`
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

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseTeams) string { return v.State }).(pulumi.StringOutput)
}

type OfficeDataConnectorDataTypesSharePoint struct {
	State string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesSharePointInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesSharePointOutput() OfficeDataConnectorDataTypesSharePointOutput
	ToOfficeDataConnectorDataTypesSharePointOutputWithContext(context.Context) OfficeDataConnectorDataTypesSharePointOutput
}

type OfficeDataConnectorDataTypesSharePointArgs struct {
	State pulumi.StringInput `pulumi:"state"`
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

func (o OfficeDataConnectorDataTypesSharePointOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesSharePoint) string { return v.State }).(pulumi.StringOutput)
}

type OfficeDataConnectorDataTypesTeams struct {
	State string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesTeamsInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesTeamsOutput() OfficeDataConnectorDataTypesTeamsOutput
	ToOfficeDataConnectorDataTypesTeamsOutputWithContext(context.Context) OfficeDataConnectorDataTypesTeamsOutput
}

type OfficeDataConnectorDataTypesTeamsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
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

func (o OfficeDataConnectorDataTypesTeamsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesTeams) string { return v.State }).(pulumi.StringOutput)
}

type SecurityAlertTimelineItemResponse struct {
	AlertType       string  `pulumi:"alertType"`
	AzureResourceId string  `pulumi:"azureResourceId"`
	Description     *string `pulumi:"description"`
	DisplayName     string  `pulumi:"displayName"`
	EndTimeUtc      string  `pulumi:"endTimeUtc"`
	Kind            string  `pulumi:"kind"`
	ProductName     *string `pulumi:"productName"`
	Severity        string  `pulumi:"severity"`
	StartTimeUtc    string  `pulumi:"startTimeUtc"`
	TimeGenerated   string  `pulumi:"timeGenerated"`
}

type TIDataConnectorDataTypes struct {
	Indicators TIDataConnectorDataTypesIndicators `pulumi:"indicators"`
}





type TIDataConnectorDataTypesInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesOutput() TIDataConnectorDataTypesOutput
	ToTIDataConnectorDataTypesOutputWithContext(context.Context) TIDataConnectorDataTypesOutput
}

type TIDataConnectorDataTypesArgs struct {
	Indicators TIDataConnectorDataTypesIndicatorsInput `pulumi:"indicators"`
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

func (o TIDataConnectorDataTypesOutput) Indicators() TIDataConnectorDataTypesIndicatorsOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypes) TIDataConnectorDataTypesIndicators { return v.Indicators }).(TIDataConnectorDataTypesIndicatorsOutput)
}

type TIDataConnectorDataTypesIndicators struct {
	State string `pulumi:"state"`
}





type TIDataConnectorDataTypesIndicatorsInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesIndicatorsOutput() TIDataConnectorDataTypesIndicatorsOutput
	ToTIDataConnectorDataTypesIndicatorsOutputWithContext(context.Context) TIDataConnectorDataTypesIndicatorsOutput
}

type TIDataConnectorDataTypesIndicatorsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
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

func (o TIDataConnectorDataTypesIndicatorsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesIndicators) string { return v.State }).(pulumi.StringOutput)
}

type TIDataConnectorDataTypesResponse struct {
	Indicators TIDataConnectorDataTypesResponseIndicators `pulumi:"indicators"`
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

func (o TIDataConnectorDataTypesResponseOutput) Indicators() TIDataConnectorDataTypesResponseIndicatorsOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesResponse) TIDataConnectorDataTypesResponseIndicators {
		return v.Indicators
	}).(TIDataConnectorDataTypesResponseIndicatorsOutput)
}

type TIDataConnectorDataTypesResponseIndicators struct {
	State string `pulumi:"state"`
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

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesResponseIndicators) string { return v.State }).(pulumi.StringOutput)
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

type TiTaxiiDataConnectorDataTypes struct {
	TaxiiClient TiTaxiiDataConnectorDataTypesTaxiiClient `pulumi:"taxiiClient"`
}





type TiTaxiiDataConnectorDataTypesInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesOutput() TiTaxiiDataConnectorDataTypesOutput
	ToTiTaxiiDataConnectorDataTypesOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesOutput
}

type TiTaxiiDataConnectorDataTypesArgs struct {
	TaxiiClient TiTaxiiDataConnectorDataTypesTaxiiClientInput `pulumi:"taxiiClient"`
}

func (TiTaxiiDataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypes)(nil)).Elem()
}

func (i TiTaxiiDataConnectorDataTypesArgs) ToTiTaxiiDataConnectorDataTypesOutput() TiTaxiiDataConnectorDataTypesOutput {
	return i.ToTiTaxiiDataConnectorDataTypesOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesArgs) ToTiTaxiiDataConnectorDataTypesOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesOutput)
}

type TiTaxiiDataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypes)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesOutput) ToTiTaxiiDataConnectorDataTypesOutput() TiTaxiiDataConnectorDataTypesOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesOutput) ToTiTaxiiDataConnectorDataTypesOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesOutput) TaxiiClient() TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypes) TiTaxiiDataConnectorDataTypesTaxiiClient { return v.TaxiiClient }).(TiTaxiiDataConnectorDataTypesTaxiiClientOutput)
}

type TiTaxiiDataConnectorDataTypesResponse struct {
	TaxiiClient TiTaxiiDataConnectorDataTypesResponseTaxiiClient `pulumi:"taxiiClient"`
}

type TiTaxiiDataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesResponseOutput) ToTiTaxiiDataConnectorDataTypesResponseOutput() TiTaxiiDataConnectorDataTypesResponseOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponseOutput) ToTiTaxiiDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponseOutput) TaxiiClient() TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypesResponse) TiTaxiiDataConnectorDataTypesResponseTaxiiClient {
		return v.TaxiiClient
	}).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput)
}

type TiTaxiiDataConnectorDataTypesResponseTaxiiClient struct {
	State string `pulumi:"state"`
}

type TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponseTaxiiClient)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypesResponseTaxiiClient) string { return v.State }).(pulumi.StringOutput)
}

type TiTaxiiDataConnectorDataTypesTaxiiClient struct {
	State string `pulumi:"state"`
}





type TiTaxiiDataConnectorDataTypesTaxiiClientInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesTaxiiClientOutput() TiTaxiiDataConnectorDataTypesTaxiiClientOutput
	ToTiTaxiiDataConnectorDataTypesTaxiiClientOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientOutput
}

type TiTaxiiDataConnectorDataTypesTaxiiClientArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (TiTaxiiDataConnectorDataTypesTaxiiClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypesTaxiiClient)(nil)).Elem()
}

func (i TiTaxiiDataConnectorDataTypesTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesTaxiiClientOutput() TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return i.ToTiTaxiiDataConnectorDataTypesTaxiiClientOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesTaxiiClientOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesTaxiiClientOutput)
}

type TiTaxiiDataConnectorDataTypesTaxiiClientOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesTaxiiClientOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypesTaxiiClient)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesTaxiiClientOutput() TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesTaxiiClientOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypesTaxiiClient) string { return v.State }).(pulumi.StringOutput)
}

type TimelineAggregationResponse struct {
	Count int    `pulumi:"count"`
	Kind  string `pulumi:"kind"`
}

type TimelineErrorResponse struct {
	ErrorMessage string  `pulumi:"errorMessage"`
	Kind         string  `pulumi:"kind"`
	QueryId      *string `pulumi:"queryId"`
}

type TimelineResultsMetadataResponse struct {
	Aggregations []TimelineAggregationResponse `pulumi:"aggregations"`
	Errors       []TimelineErrorResponse       `pulumi:"errors"`
	TotalCount   int                           `pulumi:"totalCount"`
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
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorPtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorAlertsOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorAlertsPtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponsePtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseAlertsOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionArrayOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionConditionPropertiesOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionResponseOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(ClientInfoResponseOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsResponsePtrOutput{})
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
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesDiscoveryLogsOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseDiscoveryLogsOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesIncidentsOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesResponseIncidentsOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesExchangeOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseExchangeOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseSharePointOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseTeamsOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesSharePointOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesTeamsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesIndicatorsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseIndicatorsOutput{})
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
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesTaxiiClientOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoPtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponseOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponsePtrOutput{})
}
