


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivityEntityQueriesPropertiesQueryDefinitions struct {
	Query *string `pulumi:"query"`
}





type ActivityEntityQueriesPropertiesQueryDefinitionsInput interface {
	pulumi.Input

	ToActivityEntityQueriesPropertiesQueryDefinitionsOutput() ActivityEntityQueriesPropertiesQueryDefinitionsOutput
	ToActivityEntityQueriesPropertiesQueryDefinitionsOutputWithContext(context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsOutput
}

type ActivityEntityQueriesPropertiesQueryDefinitionsArgs struct {
	Query pulumi.StringPtrInput `pulumi:"query"`
}

func (ActivityEntityQueriesPropertiesQueryDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityEntityQueriesPropertiesQueryDefinitions)(nil)).Elem()
}

func (i ActivityEntityQueriesPropertiesQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesQueryDefinitionsOutput() ActivityEntityQueriesPropertiesQueryDefinitionsOutput {
	return i.ToActivityEntityQueriesPropertiesQueryDefinitionsOutputWithContext(context.Background())
}

func (i ActivityEntityQueriesPropertiesQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesQueryDefinitionsOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityEntityQueriesPropertiesQueryDefinitionsOutput)
}

func (i ActivityEntityQueriesPropertiesQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return i.ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(context.Background())
}

func (i ActivityEntityQueriesPropertiesQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityEntityQueriesPropertiesQueryDefinitionsOutput).ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(ctx)
}









type ActivityEntityQueriesPropertiesQueryDefinitionsPtrInput interface {
	pulumi.Input

	ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput
	ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput
}

type activityEntityQueriesPropertiesQueryDefinitionsPtrType ActivityEntityQueriesPropertiesQueryDefinitionsArgs

func ActivityEntityQueriesPropertiesQueryDefinitionsPtr(v *ActivityEntityQueriesPropertiesQueryDefinitionsArgs) ActivityEntityQueriesPropertiesQueryDefinitionsPtrInput {
	return (*activityEntityQueriesPropertiesQueryDefinitionsPtrType)(v)
}

func (*activityEntityQueriesPropertiesQueryDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityEntityQueriesPropertiesQueryDefinitions)(nil)).Elem()
}

func (i *activityEntityQueriesPropertiesQueryDefinitionsPtrType) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return i.ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(context.Background())
}

func (i *activityEntityQueriesPropertiesQueryDefinitionsPtrType) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput)
}

type ActivityEntityQueriesPropertiesQueryDefinitionsOutput struct{ *pulumi.OutputState }

func (ActivityEntityQueriesPropertiesQueryDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityEntityQueriesPropertiesQueryDefinitions)(nil)).Elem()
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesQueryDefinitionsOutput() ActivityEntityQueriesPropertiesQueryDefinitionsOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesQueryDefinitionsOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return o.ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(context.Background())
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActivityEntityQueriesPropertiesQueryDefinitions) *ActivityEntityQueriesPropertiesQueryDefinitions {
		return &v
	}).(ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput)
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActivityEntityQueriesPropertiesQueryDefinitions) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityEntityQueriesPropertiesQueryDefinitions)(nil)).Elem()
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput) ToActivityEntityQueriesPropertiesQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput) Elem() ActivityEntityQueriesPropertiesQueryDefinitionsOutput {
	return o.ApplyT(func(v *ActivityEntityQueriesPropertiesQueryDefinitions) ActivityEntityQueriesPropertiesQueryDefinitions {
		if v != nil {
			return *v
		}
		var ret ActivityEntityQueriesPropertiesQueryDefinitions
		return ret
	}).(ActivityEntityQueriesPropertiesQueryDefinitionsOutput)
}

func (o ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityEntityQueriesPropertiesQueryDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

type ActivityEntityQueriesPropertiesResponseQueryDefinitions struct {
	Query *string `pulumi:"query"`
}

type ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput struct{ *pulumi.OutputState }

func (ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityEntityQueriesPropertiesResponseQueryDefinitions)(nil)).Elem()
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActivityEntityQueriesPropertiesResponseQueryDefinitions) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput struct{ *pulumi.OutputState }

func (ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityEntityQueriesPropertiesResponseQueryDefinitions)(nil)).Elem()
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput) Elem() ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput {
	return o.ApplyT(func(v *ActivityEntityQueriesPropertiesResponseQueryDefinitions) ActivityEntityQueriesPropertiesResponseQueryDefinitions {
		if v != nil {
			return *v
		}
		var ret ActivityEntityQueriesPropertiesResponseQueryDefinitions
		return ret
	}).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput)
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivityEntityQueriesPropertiesResponseQueryDefinitions) *string {
		if v == nil {
			return nil
		}
		return v.Query
	}).(pulumi.StringPtrOutput)
}

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
	Alerts DataConnectorDataTypeCommon `pulumi:"alerts"`
}





type AlertsDataTypeOfDataConnectorInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorOutput() AlertsDataTypeOfDataConnectorOutput
	ToAlertsDataTypeOfDataConnectorOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorOutput
}

type AlertsDataTypeOfDataConnectorArgs struct {
	Alerts DataConnectorDataTypeCommonInput `pulumi:"alerts"`
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

func (o AlertsDataTypeOfDataConnectorOutput) Alerts() DataConnectorDataTypeCommonOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnector) DataConnectorDataTypeCommon { return v.Alerts }).(DataConnectorDataTypeCommonOutput)
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
		return &v.Alerts
	}).(DataConnectorDataTypeCommonPtrOutput)
}

type AlertsDataTypeOfDataConnectorResponse struct {
	Alerts DataConnectorDataTypeCommonResponse `pulumi:"alerts"`
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

func (o AlertsDataTypeOfDataConnectorResponseOutput) Alerts() DataConnectorDataTypeCommonResponseOutput {
	return o.ApplyT(func(v AlertsDataTypeOfDataConnectorResponse) DataConnectorDataTypeCommonResponse { return v.Alerts }).(DataConnectorDataTypeCommonResponseOutput)
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
		return &v.Alerts
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
}

type AnomalyTimelineItemResponse struct {
	AzureResourceId string   `pulumi:"azureResourceId"`
	Description     *string  `pulumi:"description"`
	DisplayName     string   `pulumi:"displayName"`
	EndTimeUtc      string   `pulumi:"endTimeUtc"`
	Intent          *string  `pulumi:"intent"`
	Kind            string   `pulumi:"kind"`
	ProductName     *string  `pulumi:"productName"`
	Reasons         []string `pulumi:"reasons"`
	StartTimeUtc    string   `pulumi:"startTimeUtc"`
	Techniques      []string `pulumi:"techniques"`
	TimeGenerated   string   `pulumi:"timeGenerated"`
	Vendor          *string  `pulumi:"vendor"`
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

type Availability struct {
	IsPreview *bool `pulumi:"isPreview"`
	Status    *int  `pulumi:"status"`
}





type AvailabilityInput interface {
	pulumi.Input

	ToAvailabilityOutput() AvailabilityOutput
	ToAvailabilityOutputWithContext(context.Context) AvailabilityOutput
}

type AvailabilityArgs struct {
	IsPreview pulumi.BoolPtrInput `pulumi:"isPreview"`
	Status    pulumi.IntPtrInput  `pulumi:"status"`
}

func (AvailabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Availability)(nil)).Elem()
}

func (i AvailabilityArgs) ToAvailabilityOutput() AvailabilityOutput {
	return i.ToAvailabilityOutputWithContext(context.Background())
}

func (i AvailabilityArgs) ToAvailabilityOutputWithContext(ctx context.Context) AvailabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityOutput)
}

func (i AvailabilityArgs) ToAvailabilityPtrOutput() AvailabilityPtrOutput {
	return i.ToAvailabilityPtrOutputWithContext(context.Background())
}

func (i AvailabilityArgs) ToAvailabilityPtrOutputWithContext(ctx context.Context) AvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityOutput).ToAvailabilityPtrOutputWithContext(ctx)
}









type AvailabilityPtrInput interface {
	pulumi.Input

	ToAvailabilityPtrOutput() AvailabilityPtrOutput
	ToAvailabilityPtrOutputWithContext(context.Context) AvailabilityPtrOutput
}

type availabilityPtrType AvailabilityArgs

func AvailabilityPtr(v *AvailabilityArgs) AvailabilityPtrInput {
	return (*availabilityPtrType)(v)
}

func (*availabilityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Availability)(nil)).Elem()
}

func (i *availabilityPtrType) ToAvailabilityPtrOutput() AvailabilityPtrOutput {
	return i.ToAvailabilityPtrOutputWithContext(context.Background())
}

func (i *availabilityPtrType) ToAvailabilityPtrOutputWithContext(ctx context.Context) AvailabilityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityPtrOutput)
}

type AvailabilityOutput struct{ *pulumi.OutputState }

func (AvailabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Availability)(nil)).Elem()
}

func (o AvailabilityOutput) ToAvailabilityOutput() AvailabilityOutput {
	return o
}

func (o AvailabilityOutput) ToAvailabilityOutputWithContext(ctx context.Context) AvailabilityOutput {
	return o
}

func (o AvailabilityOutput) ToAvailabilityPtrOutput() AvailabilityPtrOutput {
	return o.ToAvailabilityPtrOutputWithContext(context.Background())
}

func (o AvailabilityOutput) ToAvailabilityPtrOutputWithContext(ctx context.Context) AvailabilityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Availability) *Availability {
		return &v
	}).(AvailabilityPtrOutput)
}

func (o AvailabilityOutput) IsPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Availability) *bool { return v.IsPreview }).(pulumi.BoolPtrOutput)
}

func (o AvailabilityOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Availability) *int { return v.Status }).(pulumi.IntPtrOutput)
}

type AvailabilityPtrOutput struct{ *pulumi.OutputState }

func (AvailabilityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Availability)(nil)).Elem()
}

func (o AvailabilityPtrOutput) ToAvailabilityPtrOutput() AvailabilityPtrOutput {
	return o
}

func (o AvailabilityPtrOutput) ToAvailabilityPtrOutputWithContext(ctx context.Context) AvailabilityPtrOutput {
	return o
}

func (o AvailabilityPtrOutput) Elem() AvailabilityOutput {
	return o.ApplyT(func(v *Availability) Availability {
		if v != nil {
			return *v
		}
		var ret Availability
		return ret
	}).(AvailabilityOutput)
}

func (o AvailabilityPtrOutput) IsPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Availability) *bool {
		if v == nil {
			return nil
		}
		return v.IsPreview
	}).(pulumi.BoolPtrOutput)
}

func (o AvailabilityPtrOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Availability) *int {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.IntPtrOutput)
}

type AvailabilityResponse struct {
	IsPreview *bool `pulumi:"isPreview"`
	Status    *int  `pulumi:"status"`
}

type AvailabilityResponseOutput struct{ *pulumi.OutputState }

func (AvailabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityResponse)(nil)).Elem()
}

func (o AvailabilityResponseOutput) ToAvailabilityResponseOutput() AvailabilityResponseOutput {
	return o
}

func (o AvailabilityResponseOutput) ToAvailabilityResponseOutputWithContext(ctx context.Context) AvailabilityResponseOutput {
	return o
}

func (o AvailabilityResponseOutput) IsPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AvailabilityResponse) *bool { return v.IsPreview }).(pulumi.BoolPtrOutput)
}

func (o AvailabilityResponseOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AvailabilityResponse) *int { return v.Status }).(pulumi.IntPtrOutput)
}

type AvailabilityResponsePtrOutput struct{ *pulumi.OutputState }

func (AvailabilityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityResponse)(nil)).Elem()
}

func (o AvailabilityResponsePtrOutput) ToAvailabilityResponsePtrOutput() AvailabilityResponsePtrOutput {
	return o
}

func (o AvailabilityResponsePtrOutput) ToAvailabilityResponsePtrOutputWithContext(ctx context.Context) AvailabilityResponsePtrOutput {
	return o
}

func (o AvailabilityResponsePtrOutput) Elem() AvailabilityResponseOutput {
	return o.ApplyT(func(v *AvailabilityResponse) AvailabilityResponse {
		if v != nil {
			return *v
		}
		var ret AvailabilityResponse
		return ret
	}).(AvailabilityResponseOutput)
}

func (o AvailabilityResponsePtrOutput) IsPreview() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AvailabilityResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsPreview
	}).(pulumi.BoolPtrOutput)
}

func (o AvailabilityResponsePtrOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AvailabilityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.IntPtrOutput)
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

type AwsS3DataConnectorDataTypes struct {
	Logs AwsS3DataConnectorDataTypesLogs `pulumi:"logs"`
}





type AwsS3DataConnectorDataTypesInput interface {
	pulumi.Input

	ToAwsS3DataConnectorDataTypesOutput() AwsS3DataConnectorDataTypesOutput
	ToAwsS3DataConnectorDataTypesOutputWithContext(context.Context) AwsS3DataConnectorDataTypesOutput
}

type AwsS3DataConnectorDataTypesArgs struct {
	Logs AwsS3DataConnectorDataTypesLogsInput `pulumi:"logs"`
}

func (AwsS3DataConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsS3DataConnectorDataTypes)(nil)).Elem()
}

func (i AwsS3DataConnectorDataTypesArgs) ToAwsS3DataConnectorDataTypesOutput() AwsS3DataConnectorDataTypesOutput {
	return i.ToAwsS3DataConnectorDataTypesOutputWithContext(context.Background())
}

func (i AwsS3DataConnectorDataTypesArgs) ToAwsS3DataConnectorDataTypesOutputWithContext(ctx context.Context) AwsS3DataConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsS3DataConnectorDataTypesOutput)
}

type AwsS3DataConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (AwsS3DataConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsS3DataConnectorDataTypes)(nil)).Elem()
}

func (o AwsS3DataConnectorDataTypesOutput) ToAwsS3DataConnectorDataTypesOutput() AwsS3DataConnectorDataTypesOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesOutput) ToAwsS3DataConnectorDataTypesOutputWithContext(ctx context.Context) AwsS3DataConnectorDataTypesOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesOutput) Logs() AwsS3DataConnectorDataTypesLogsOutput {
	return o.ApplyT(func(v AwsS3DataConnectorDataTypes) AwsS3DataConnectorDataTypesLogs { return v.Logs }).(AwsS3DataConnectorDataTypesLogsOutput)
}

type AwsS3DataConnectorDataTypesLogs struct {
	State string `pulumi:"state"`
}





type AwsS3DataConnectorDataTypesLogsInput interface {
	pulumi.Input

	ToAwsS3DataConnectorDataTypesLogsOutput() AwsS3DataConnectorDataTypesLogsOutput
	ToAwsS3DataConnectorDataTypesLogsOutputWithContext(context.Context) AwsS3DataConnectorDataTypesLogsOutput
}

type AwsS3DataConnectorDataTypesLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (AwsS3DataConnectorDataTypesLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsS3DataConnectorDataTypesLogs)(nil)).Elem()
}

func (i AwsS3DataConnectorDataTypesLogsArgs) ToAwsS3DataConnectorDataTypesLogsOutput() AwsS3DataConnectorDataTypesLogsOutput {
	return i.ToAwsS3DataConnectorDataTypesLogsOutputWithContext(context.Background())
}

func (i AwsS3DataConnectorDataTypesLogsArgs) ToAwsS3DataConnectorDataTypesLogsOutputWithContext(ctx context.Context) AwsS3DataConnectorDataTypesLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsS3DataConnectorDataTypesLogsOutput)
}

type AwsS3DataConnectorDataTypesLogsOutput struct{ *pulumi.OutputState }

func (AwsS3DataConnectorDataTypesLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsS3DataConnectorDataTypesLogs)(nil)).Elem()
}

func (o AwsS3DataConnectorDataTypesLogsOutput) ToAwsS3DataConnectorDataTypesLogsOutput() AwsS3DataConnectorDataTypesLogsOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesLogsOutput) ToAwsS3DataConnectorDataTypesLogsOutputWithContext(ctx context.Context) AwsS3DataConnectorDataTypesLogsOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AwsS3DataConnectorDataTypesLogs) string { return v.State }).(pulumi.StringOutput)
}

type AwsS3DataConnectorDataTypesResponse struct {
	Logs AwsS3DataConnectorDataTypesResponseLogs `pulumi:"logs"`
}

type AwsS3DataConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (AwsS3DataConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsS3DataConnectorDataTypesResponse)(nil)).Elem()
}

func (o AwsS3DataConnectorDataTypesResponseOutput) ToAwsS3DataConnectorDataTypesResponseOutput() AwsS3DataConnectorDataTypesResponseOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesResponseOutput) ToAwsS3DataConnectorDataTypesResponseOutputWithContext(ctx context.Context) AwsS3DataConnectorDataTypesResponseOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesResponseOutput) Logs() AwsS3DataConnectorDataTypesResponseLogsOutput {
	return o.ApplyT(func(v AwsS3DataConnectorDataTypesResponse) AwsS3DataConnectorDataTypesResponseLogs { return v.Logs }).(AwsS3DataConnectorDataTypesResponseLogsOutput)
}

type AwsS3DataConnectorDataTypesResponseLogs struct {
	State string `pulumi:"state"`
}

type AwsS3DataConnectorDataTypesResponseLogsOutput struct{ *pulumi.OutputState }

func (AwsS3DataConnectorDataTypesResponseLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsS3DataConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (o AwsS3DataConnectorDataTypesResponseLogsOutput) ToAwsS3DataConnectorDataTypesResponseLogsOutput() AwsS3DataConnectorDataTypesResponseLogsOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesResponseLogsOutput) ToAwsS3DataConnectorDataTypesResponseLogsOutputWithContext(ctx context.Context) AwsS3DataConnectorDataTypesResponseLogsOutput {
	return o
}

func (o AwsS3DataConnectorDataTypesResponseLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AwsS3DataConnectorDataTypesResponseLogs) string { return v.State }).(pulumi.StringOutput)
}

type AzureDevOpsResourceInfo struct {
	PipelineId          *string `pulumi:"pipelineId"`
	ServiceConnectionId *string `pulumi:"serviceConnectionId"`
}





type AzureDevOpsResourceInfoInput interface {
	pulumi.Input

	ToAzureDevOpsResourceInfoOutput() AzureDevOpsResourceInfoOutput
	ToAzureDevOpsResourceInfoOutputWithContext(context.Context) AzureDevOpsResourceInfoOutput
}

type AzureDevOpsResourceInfoArgs struct {
	PipelineId          pulumi.StringPtrInput `pulumi:"pipelineId"`
	ServiceConnectionId pulumi.StringPtrInput `pulumi:"serviceConnectionId"`
}

func (AzureDevOpsResourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsResourceInfo)(nil)).Elem()
}

func (i AzureDevOpsResourceInfoArgs) ToAzureDevOpsResourceInfoOutput() AzureDevOpsResourceInfoOutput {
	return i.ToAzureDevOpsResourceInfoOutputWithContext(context.Background())
}

func (i AzureDevOpsResourceInfoArgs) ToAzureDevOpsResourceInfoOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsResourceInfoOutput)
}

func (i AzureDevOpsResourceInfoArgs) ToAzureDevOpsResourceInfoPtrOutput() AzureDevOpsResourceInfoPtrOutput {
	return i.ToAzureDevOpsResourceInfoPtrOutputWithContext(context.Background())
}

func (i AzureDevOpsResourceInfoArgs) ToAzureDevOpsResourceInfoPtrOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsResourceInfoOutput).ToAzureDevOpsResourceInfoPtrOutputWithContext(ctx)
}









type AzureDevOpsResourceInfoPtrInput interface {
	pulumi.Input

	ToAzureDevOpsResourceInfoPtrOutput() AzureDevOpsResourceInfoPtrOutput
	ToAzureDevOpsResourceInfoPtrOutputWithContext(context.Context) AzureDevOpsResourceInfoPtrOutput
}

type azureDevOpsResourceInfoPtrType AzureDevOpsResourceInfoArgs

func AzureDevOpsResourceInfoPtr(v *AzureDevOpsResourceInfoArgs) AzureDevOpsResourceInfoPtrInput {
	return (*azureDevOpsResourceInfoPtrType)(v)
}

func (*azureDevOpsResourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsResourceInfo)(nil)).Elem()
}

func (i *azureDevOpsResourceInfoPtrType) ToAzureDevOpsResourceInfoPtrOutput() AzureDevOpsResourceInfoPtrOutput {
	return i.ToAzureDevOpsResourceInfoPtrOutputWithContext(context.Background())
}

func (i *azureDevOpsResourceInfoPtrType) ToAzureDevOpsResourceInfoPtrOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureDevOpsResourceInfoPtrOutput)
}

type AzureDevOpsResourceInfoOutput struct{ *pulumi.OutputState }

func (AzureDevOpsResourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsResourceInfo)(nil)).Elem()
}

func (o AzureDevOpsResourceInfoOutput) ToAzureDevOpsResourceInfoOutput() AzureDevOpsResourceInfoOutput {
	return o
}

func (o AzureDevOpsResourceInfoOutput) ToAzureDevOpsResourceInfoOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoOutput {
	return o
}

func (o AzureDevOpsResourceInfoOutput) ToAzureDevOpsResourceInfoPtrOutput() AzureDevOpsResourceInfoPtrOutput {
	return o.ToAzureDevOpsResourceInfoPtrOutputWithContext(context.Background())
}

func (o AzureDevOpsResourceInfoOutput) ToAzureDevOpsResourceInfoPtrOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureDevOpsResourceInfo) *AzureDevOpsResourceInfo {
		return &v
	}).(AzureDevOpsResourceInfoPtrOutput)
}

func (o AzureDevOpsResourceInfoOutput) PipelineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsResourceInfo) *string { return v.PipelineId }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsResourceInfoOutput) ServiceConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsResourceInfo) *string { return v.ServiceConnectionId }).(pulumi.StringPtrOutput)
}

type AzureDevOpsResourceInfoPtrOutput struct{ *pulumi.OutputState }

func (AzureDevOpsResourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsResourceInfo)(nil)).Elem()
}

func (o AzureDevOpsResourceInfoPtrOutput) ToAzureDevOpsResourceInfoPtrOutput() AzureDevOpsResourceInfoPtrOutput {
	return o
}

func (o AzureDevOpsResourceInfoPtrOutput) ToAzureDevOpsResourceInfoPtrOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoPtrOutput {
	return o
}

func (o AzureDevOpsResourceInfoPtrOutput) Elem() AzureDevOpsResourceInfoOutput {
	return o.ApplyT(func(v *AzureDevOpsResourceInfo) AzureDevOpsResourceInfo {
		if v != nil {
			return *v
		}
		var ret AzureDevOpsResourceInfo
		return ret
	}).(AzureDevOpsResourceInfoOutput)
}

func (o AzureDevOpsResourceInfoPtrOutput) PipelineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDevOpsResourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.PipelineId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsResourceInfoPtrOutput) ServiceConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDevOpsResourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.ServiceConnectionId
	}).(pulumi.StringPtrOutput)
}

type AzureDevOpsResourceInfoResponse struct {
	PipelineId          *string `pulumi:"pipelineId"`
	ServiceConnectionId *string `pulumi:"serviceConnectionId"`
}

type AzureDevOpsResourceInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureDevOpsResourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureDevOpsResourceInfoResponse)(nil)).Elem()
}

func (o AzureDevOpsResourceInfoResponseOutput) ToAzureDevOpsResourceInfoResponseOutput() AzureDevOpsResourceInfoResponseOutput {
	return o
}

func (o AzureDevOpsResourceInfoResponseOutput) ToAzureDevOpsResourceInfoResponseOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoResponseOutput {
	return o
}

func (o AzureDevOpsResourceInfoResponseOutput) PipelineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsResourceInfoResponse) *string { return v.PipelineId }).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsResourceInfoResponseOutput) ServiceConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureDevOpsResourceInfoResponse) *string { return v.ServiceConnectionId }).(pulumi.StringPtrOutput)
}

type AzureDevOpsResourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureDevOpsResourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureDevOpsResourceInfoResponse)(nil)).Elem()
}

func (o AzureDevOpsResourceInfoResponsePtrOutput) ToAzureDevOpsResourceInfoResponsePtrOutput() AzureDevOpsResourceInfoResponsePtrOutput {
	return o
}

func (o AzureDevOpsResourceInfoResponsePtrOutput) ToAzureDevOpsResourceInfoResponsePtrOutputWithContext(ctx context.Context) AzureDevOpsResourceInfoResponsePtrOutput {
	return o
}

func (o AzureDevOpsResourceInfoResponsePtrOutput) Elem() AzureDevOpsResourceInfoResponseOutput {
	return o.ApplyT(func(v *AzureDevOpsResourceInfoResponse) AzureDevOpsResourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureDevOpsResourceInfoResponse
		return ret
	}).(AzureDevOpsResourceInfoResponseOutput)
}

func (o AzureDevOpsResourceInfoResponsePtrOutput) PipelineId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDevOpsResourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.PipelineId
	}).(pulumi.StringPtrOutput)
}

func (o AzureDevOpsResourceInfoResponsePtrOutput) ServiceConnectionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureDevOpsResourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceConnectionId
	}).(pulumi.StringPtrOutput)
}

type BookmarkEntityMappings struct {
	EntityType    *string              `pulumi:"entityType"`
	FieldMappings []EntityFieldMapping `pulumi:"fieldMappings"`
}





type BookmarkEntityMappingsInput interface {
	pulumi.Input

	ToBookmarkEntityMappingsOutput() BookmarkEntityMappingsOutput
	ToBookmarkEntityMappingsOutputWithContext(context.Context) BookmarkEntityMappingsOutput
}

type BookmarkEntityMappingsArgs struct {
	EntityType    pulumi.StringPtrInput        `pulumi:"entityType"`
	FieldMappings EntityFieldMappingArrayInput `pulumi:"fieldMappings"`
}

func (BookmarkEntityMappingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BookmarkEntityMappings)(nil)).Elem()
}

func (i BookmarkEntityMappingsArgs) ToBookmarkEntityMappingsOutput() BookmarkEntityMappingsOutput {
	return i.ToBookmarkEntityMappingsOutputWithContext(context.Background())
}

func (i BookmarkEntityMappingsArgs) ToBookmarkEntityMappingsOutputWithContext(ctx context.Context) BookmarkEntityMappingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkEntityMappingsOutput)
}





type BookmarkEntityMappingsArrayInput interface {
	pulumi.Input

	ToBookmarkEntityMappingsArrayOutput() BookmarkEntityMappingsArrayOutput
	ToBookmarkEntityMappingsArrayOutputWithContext(context.Context) BookmarkEntityMappingsArrayOutput
}

type BookmarkEntityMappingsArray []BookmarkEntityMappingsInput

func (BookmarkEntityMappingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BookmarkEntityMappings)(nil)).Elem()
}

func (i BookmarkEntityMappingsArray) ToBookmarkEntityMappingsArrayOutput() BookmarkEntityMappingsArrayOutput {
	return i.ToBookmarkEntityMappingsArrayOutputWithContext(context.Background())
}

func (i BookmarkEntityMappingsArray) ToBookmarkEntityMappingsArrayOutputWithContext(ctx context.Context) BookmarkEntityMappingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkEntityMappingsArrayOutput)
}

type BookmarkEntityMappingsOutput struct{ *pulumi.OutputState }

func (BookmarkEntityMappingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BookmarkEntityMappings)(nil)).Elem()
}

func (o BookmarkEntityMappingsOutput) ToBookmarkEntityMappingsOutput() BookmarkEntityMappingsOutput {
	return o
}

func (o BookmarkEntityMappingsOutput) ToBookmarkEntityMappingsOutputWithContext(ctx context.Context) BookmarkEntityMappingsOutput {
	return o
}

func (o BookmarkEntityMappingsOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkEntityMappings) *string { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o BookmarkEntityMappingsOutput) FieldMappings() EntityFieldMappingArrayOutput {
	return o.ApplyT(func(v BookmarkEntityMappings) []EntityFieldMapping { return v.FieldMappings }).(EntityFieldMappingArrayOutput)
}

type BookmarkEntityMappingsArrayOutput struct{ *pulumi.OutputState }

func (BookmarkEntityMappingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BookmarkEntityMappings)(nil)).Elem()
}

func (o BookmarkEntityMappingsArrayOutput) ToBookmarkEntityMappingsArrayOutput() BookmarkEntityMappingsArrayOutput {
	return o
}

func (o BookmarkEntityMappingsArrayOutput) ToBookmarkEntityMappingsArrayOutputWithContext(ctx context.Context) BookmarkEntityMappingsArrayOutput {
	return o
}

func (o BookmarkEntityMappingsArrayOutput) Index(i pulumi.IntInput) BookmarkEntityMappingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BookmarkEntityMappings {
		return vs[0].([]BookmarkEntityMappings)[vs[1].(int)]
	}).(BookmarkEntityMappingsOutput)
}

type BookmarkEntityMappingsResponse struct {
	EntityType    *string                      `pulumi:"entityType"`
	FieldMappings []EntityFieldMappingResponse `pulumi:"fieldMappings"`
}

type BookmarkEntityMappingsResponseOutput struct{ *pulumi.OutputState }

func (BookmarkEntityMappingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BookmarkEntityMappingsResponse)(nil)).Elem()
}

func (o BookmarkEntityMappingsResponseOutput) ToBookmarkEntityMappingsResponseOutput() BookmarkEntityMappingsResponseOutput {
	return o
}

func (o BookmarkEntityMappingsResponseOutput) ToBookmarkEntityMappingsResponseOutputWithContext(ctx context.Context) BookmarkEntityMappingsResponseOutput {
	return o
}

func (o BookmarkEntityMappingsResponseOutput) EntityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkEntityMappingsResponse) *string { return v.EntityType }).(pulumi.StringPtrOutput)
}

func (o BookmarkEntityMappingsResponseOutput) FieldMappings() EntityFieldMappingResponseArrayOutput {
	return o.ApplyT(func(v BookmarkEntityMappingsResponse) []EntityFieldMappingResponse { return v.FieldMappings }).(EntityFieldMappingResponseArrayOutput)
}

type BookmarkEntityMappingsResponseArrayOutput struct{ *pulumi.OutputState }

func (BookmarkEntityMappingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BookmarkEntityMappingsResponse)(nil)).Elem()
}

func (o BookmarkEntityMappingsResponseArrayOutput) ToBookmarkEntityMappingsResponseArrayOutput() BookmarkEntityMappingsResponseArrayOutput {
	return o
}

func (o BookmarkEntityMappingsResponseArrayOutput) ToBookmarkEntityMappingsResponseArrayOutputWithContext(ctx context.Context) BookmarkEntityMappingsResponseArrayOutput {
	return o
}

func (o BookmarkEntityMappingsResponseArrayOutput) Index(i pulumi.IntInput) BookmarkEntityMappingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BookmarkEntityMappingsResponse {
		return vs[0].([]BookmarkEntityMappingsResponse)[vs[1].(int)]
	}).(BookmarkEntityMappingsResponseOutput)
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

type CodelessConnectorPollingAuthProperties struct {
	ApiKeyIdentifier                     *string     `pulumi:"apiKeyIdentifier"`
	ApiKeyName                           *string     `pulumi:"apiKeyName"`
	AuthType                             string      `pulumi:"authType"`
	AuthorizationEndpoint                *string     `pulumi:"authorizationEndpoint"`
	AuthorizationEndpointQueryParameters interface{} `pulumi:"authorizationEndpointQueryParameters"`
	FlowName                             *string     `pulumi:"flowName"`
	IsApiKeyInPostPayload                *string     `pulumi:"isApiKeyInPostPayload"`
	IsClientSecretInHeader               *bool       `pulumi:"isClientSecretInHeader"`
	RedirectionEndpoint                  *string     `pulumi:"redirectionEndpoint"`
	Scope                                *string     `pulumi:"scope"`
	TokenEndpoint                        *string     `pulumi:"tokenEndpoint"`
	TokenEndpointHeaders                 interface{} `pulumi:"tokenEndpointHeaders"`
	TokenEndpointQueryParameters         interface{} `pulumi:"tokenEndpointQueryParameters"`
}





type CodelessConnectorPollingAuthPropertiesInput interface {
	pulumi.Input

	ToCodelessConnectorPollingAuthPropertiesOutput() CodelessConnectorPollingAuthPropertiesOutput
	ToCodelessConnectorPollingAuthPropertiesOutputWithContext(context.Context) CodelessConnectorPollingAuthPropertiesOutput
}

type CodelessConnectorPollingAuthPropertiesArgs struct {
	ApiKeyIdentifier                     pulumi.StringPtrInput `pulumi:"apiKeyIdentifier"`
	ApiKeyName                           pulumi.StringPtrInput `pulumi:"apiKeyName"`
	AuthType                             pulumi.StringInput    `pulumi:"authType"`
	AuthorizationEndpoint                pulumi.StringPtrInput `pulumi:"authorizationEndpoint"`
	AuthorizationEndpointQueryParameters pulumi.Input          `pulumi:"authorizationEndpointQueryParameters"`
	FlowName                             pulumi.StringPtrInput `pulumi:"flowName"`
	IsApiKeyInPostPayload                pulumi.StringPtrInput `pulumi:"isApiKeyInPostPayload"`
	IsClientSecretInHeader               pulumi.BoolPtrInput   `pulumi:"isClientSecretInHeader"`
	RedirectionEndpoint                  pulumi.StringPtrInput `pulumi:"redirectionEndpoint"`
	Scope                                pulumi.StringPtrInput `pulumi:"scope"`
	TokenEndpoint                        pulumi.StringPtrInput `pulumi:"tokenEndpoint"`
	TokenEndpointHeaders                 pulumi.Input          `pulumi:"tokenEndpointHeaders"`
	TokenEndpointQueryParameters         pulumi.Input          `pulumi:"tokenEndpointQueryParameters"`
}

func (CodelessConnectorPollingAuthPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingAuthProperties)(nil)).Elem()
}

func (i CodelessConnectorPollingAuthPropertiesArgs) ToCodelessConnectorPollingAuthPropertiesOutput() CodelessConnectorPollingAuthPropertiesOutput {
	return i.ToCodelessConnectorPollingAuthPropertiesOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingAuthPropertiesArgs) ToCodelessConnectorPollingAuthPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingAuthPropertiesOutput)
}

func (i CodelessConnectorPollingAuthPropertiesArgs) ToCodelessConnectorPollingAuthPropertiesPtrOutput() CodelessConnectorPollingAuthPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingAuthPropertiesArgs) ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingAuthPropertiesOutput).ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(ctx)
}









type CodelessConnectorPollingAuthPropertiesPtrInput interface {
	pulumi.Input

	ToCodelessConnectorPollingAuthPropertiesPtrOutput() CodelessConnectorPollingAuthPropertiesPtrOutput
	ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(context.Context) CodelessConnectorPollingAuthPropertiesPtrOutput
}

type codelessConnectorPollingAuthPropertiesPtrType CodelessConnectorPollingAuthPropertiesArgs

func CodelessConnectorPollingAuthPropertiesPtr(v *CodelessConnectorPollingAuthPropertiesArgs) CodelessConnectorPollingAuthPropertiesPtrInput {
	return (*codelessConnectorPollingAuthPropertiesPtrType)(v)
}

func (*codelessConnectorPollingAuthPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingAuthProperties)(nil)).Elem()
}

func (i *codelessConnectorPollingAuthPropertiesPtrType) ToCodelessConnectorPollingAuthPropertiesPtrOutput() CodelessConnectorPollingAuthPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(context.Background())
}

func (i *codelessConnectorPollingAuthPropertiesPtrType) ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingAuthPropertiesPtrOutput)
}

type CodelessConnectorPollingAuthPropertiesOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingAuthPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingAuthProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingAuthPropertiesOutput) ToCodelessConnectorPollingAuthPropertiesOutput() CodelessConnectorPollingAuthPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesOutput) ToCodelessConnectorPollingAuthPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesOutput) ToCodelessConnectorPollingAuthPropertiesPtrOutput() CodelessConnectorPollingAuthPropertiesPtrOutput {
	return o.ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(context.Background())
}

func (o CodelessConnectorPollingAuthPropertiesOutput) ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessConnectorPollingAuthProperties) *CodelessConnectorPollingAuthProperties {
		return &v
	}).(CodelessConnectorPollingAuthPropertiesPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) ApiKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.ApiKeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) ApiKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.ApiKeyName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) AuthorizationEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) interface{} {
		return v.AuthorizationEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) FlowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.FlowName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) IsApiKeyInPostPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.IsApiKeyInPostPayload }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) IsClientSecretInHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *bool { return v.IsClientSecretInHeader }).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) RedirectionEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.RedirectionEndpoint }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) TokenEndpointHeaders() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) interface{} { return v.TokenEndpointHeaders }).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesOutput) TokenEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthProperties) interface{} { return v.TokenEndpointQueryParameters }).(pulumi.AnyOutput)
}

type CodelessConnectorPollingAuthPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingAuthPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingAuthProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) ToCodelessConnectorPollingAuthPropertiesPtrOutput() CodelessConnectorPollingAuthPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) ToCodelessConnectorPollingAuthPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) Elem() CodelessConnectorPollingAuthPropertiesOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) CodelessConnectorPollingAuthProperties {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingAuthProperties
		return ret
	}).(CodelessConnectorPollingAuthPropertiesOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) ApiKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) ApiKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiKeyName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AuthType
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) AuthorizationEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) FlowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.FlowName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) IsApiKeyInPostPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.IsApiKeyInPostPayload
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) IsClientSecretInHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsClientSecretInHeader
	}).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) RedirectionEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.RedirectionEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) *string {
		if v == nil {
			return nil
		}
		return v.TokenEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) TokenEndpointHeaders() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.TokenEndpointHeaders
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesPtrOutput) TokenEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.TokenEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

type CodelessConnectorPollingAuthPropertiesResponse struct {
	ApiKeyIdentifier                     *string     `pulumi:"apiKeyIdentifier"`
	ApiKeyName                           *string     `pulumi:"apiKeyName"`
	AuthType                             string      `pulumi:"authType"`
	AuthorizationEndpoint                *string     `pulumi:"authorizationEndpoint"`
	AuthorizationEndpointQueryParameters interface{} `pulumi:"authorizationEndpointQueryParameters"`
	FlowName                             *string     `pulumi:"flowName"`
	IsApiKeyInPostPayload                *string     `pulumi:"isApiKeyInPostPayload"`
	IsClientSecretInHeader               *bool       `pulumi:"isClientSecretInHeader"`
	RedirectionEndpoint                  *string     `pulumi:"redirectionEndpoint"`
	Scope                                *string     `pulumi:"scope"`
	TokenEndpoint                        *string     `pulumi:"tokenEndpoint"`
	TokenEndpointHeaders                 interface{} `pulumi:"tokenEndpointHeaders"`
	TokenEndpointQueryParameters         interface{} `pulumi:"tokenEndpointQueryParameters"`
}

type CodelessConnectorPollingAuthPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingAuthPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingAuthPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) ToCodelessConnectorPollingAuthPropertiesResponseOutput() CodelessConnectorPollingAuthPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) ToCodelessConnectorPollingAuthPropertiesResponseOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) ApiKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.ApiKeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) ApiKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.ApiKeyName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) AuthType() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) string { return v.AuthType }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.AuthorizationEndpoint }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) AuthorizationEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) interface{} {
		return v.AuthorizationEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) FlowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.FlowName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) IsApiKeyInPostPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.IsApiKeyInPostPayload }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) IsClientSecretInHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *bool { return v.IsClientSecretInHeader }).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) RedirectionEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.RedirectionEndpoint }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) *string { return v.TokenEndpoint }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) TokenEndpointHeaders() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) interface{} { return v.TokenEndpointHeaders }).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponseOutput) TokenEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingAuthPropertiesResponse) interface{} {
		return v.TokenEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

type CodelessConnectorPollingAuthPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingAuthPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingAuthPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) ToCodelessConnectorPollingAuthPropertiesResponsePtrOutput() CodelessConnectorPollingAuthPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) ToCodelessConnectorPollingAuthPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessConnectorPollingAuthPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) Elem() CodelessConnectorPollingAuthPropertiesResponseOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) CodelessConnectorPollingAuthPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingAuthPropertiesResponse
		return ret
	}).(CodelessConnectorPollingAuthPropertiesResponseOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) ApiKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) ApiKeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiKeyName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) AuthType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AuthType
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) AuthorizationEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) AuthorizationEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.AuthorizationEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) FlowName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FlowName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) IsApiKeyInPostPayload() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IsApiKeyInPostPayload
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) IsClientSecretInHeader() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsClientSecretInHeader
	}).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) RedirectionEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RedirectionEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scope
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) TokenEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TokenEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) TokenEndpointHeaders() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.TokenEndpointHeaders
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingAuthPropertiesResponsePtrOutput) TokenEndpointQueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingAuthPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.TokenEndpointQueryParameters
	}).(pulumi.AnyOutput)
}

type CodelessConnectorPollingConfigProperties struct {
	Auth     CodelessConnectorPollingAuthProperties      `pulumi:"auth"`
	IsActive *bool                                       `pulumi:"isActive"`
	Paging   *CodelessConnectorPollingPagingProperties   `pulumi:"paging"`
	Request  CodelessConnectorPollingRequestProperties   `pulumi:"request"`
	Response *CodelessConnectorPollingResponseProperties `pulumi:"response"`
}





type CodelessConnectorPollingConfigPropertiesInput interface {
	pulumi.Input

	ToCodelessConnectorPollingConfigPropertiesOutput() CodelessConnectorPollingConfigPropertiesOutput
	ToCodelessConnectorPollingConfigPropertiesOutputWithContext(context.Context) CodelessConnectorPollingConfigPropertiesOutput
}

type CodelessConnectorPollingConfigPropertiesArgs struct {
	Auth     CodelessConnectorPollingAuthPropertiesInput        `pulumi:"auth"`
	IsActive pulumi.BoolPtrInput                                `pulumi:"isActive"`
	Paging   CodelessConnectorPollingPagingPropertiesPtrInput   `pulumi:"paging"`
	Request  CodelessConnectorPollingRequestPropertiesInput     `pulumi:"request"`
	Response CodelessConnectorPollingResponsePropertiesPtrInput `pulumi:"response"`
}

func (CodelessConnectorPollingConfigPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingConfigProperties)(nil)).Elem()
}

func (i CodelessConnectorPollingConfigPropertiesArgs) ToCodelessConnectorPollingConfigPropertiesOutput() CodelessConnectorPollingConfigPropertiesOutput {
	return i.ToCodelessConnectorPollingConfigPropertiesOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingConfigPropertiesArgs) ToCodelessConnectorPollingConfigPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingConfigPropertiesOutput)
}

func (i CodelessConnectorPollingConfigPropertiesArgs) ToCodelessConnectorPollingConfigPropertiesPtrOutput() CodelessConnectorPollingConfigPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingConfigPropertiesArgs) ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingConfigPropertiesOutput).ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(ctx)
}









type CodelessConnectorPollingConfigPropertiesPtrInput interface {
	pulumi.Input

	ToCodelessConnectorPollingConfigPropertiesPtrOutput() CodelessConnectorPollingConfigPropertiesPtrOutput
	ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(context.Context) CodelessConnectorPollingConfigPropertiesPtrOutput
}

type codelessConnectorPollingConfigPropertiesPtrType CodelessConnectorPollingConfigPropertiesArgs

func CodelessConnectorPollingConfigPropertiesPtr(v *CodelessConnectorPollingConfigPropertiesArgs) CodelessConnectorPollingConfigPropertiesPtrInput {
	return (*codelessConnectorPollingConfigPropertiesPtrType)(v)
}

func (*codelessConnectorPollingConfigPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingConfigProperties)(nil)).Elem()
}

func (i *codelessConnectorPollingConfigPropertiesPtrType) ToCodelessConnectorPollingConfigPropertiesPtrOutput() CodelessConnectorPollingConfigPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i *codelessConnectorPollingConfigPropertiesPtrType) ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingConfigPropertiesPtrOutput)
}

type CodelessConnectorPollingConfigPropertiesOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingConfigPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingConfigProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingConfigPropertiesOutput) ToCodelessConnectorPollingConfigPropertiesOutput() CodelessConnectorPollingConfigPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesOutput) ToCodelessConnectorPollingConfigPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesOutput) ToCodelessConnectorPollingConfigPropertiesPtrOutput() CodelessConnectorPollingConfigPropertiesPtrOutput {
	return o.ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(context.Background())
}

func (o CodelessConnectorPollingConfigPropertiesOutput) ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingConfigProperties {
		return &v
	}).(CodelessConnectorPollingConfigPropertiesPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesOutput) Auth() CodelessConnectorPollingAuthPropertiesOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigProperties) CodelessConnectorPollingAuthProperties { return v.Auth }).(CodelessConnectorPollingAuthPropertiesOutput)
}

func (o CodelessConnectorPollingConfigPropertiesOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigProperties) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesOutput) Paging() CodelessConnectorPollingPagingPropertiesPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingPagingProperties {
		return v.Paging
	}).(CodelessConnectorPollingPagingPropertiesPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesOutput) Request() CodelessConnectorPollingRequestPropertiesOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigProperties) CodelessConnectorPollingRequestProperties {
		return v.Request
	}).(CodelessConnectorPollingRequestPropertiesOutput)
}

func (o CodelessConnectorPollingConfigPropertiesOutput) Response() CodelessConnectorPollingResponsePropertiesPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingResponseProperties {
		return v.Response
	}).(CodelessConnectorPollingResponsePropertiesPtrOutput)
}

type CodelessConnectorPollingConfigPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingConfigPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingConfigProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) ToCodelessConnectorPollingConfigPropertiesPtrOutput() CodelessConnectorPollingConfigPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) ToCodelessConnectorPollingConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) Elem() CodelessConnectorPollingConfigPropertiesOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigProperties) CodelessConnectorPollingConfigProperties {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingConfigProperties
		return ret
	}).(CodelessConnectorPollingConfigPropertiesOutput)
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) Auth() CodelessConnectorPollingAuthPropertiesPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingAuthProperties {
		if v == nil {
			return nil
		}
		return &v.Auth
	}).(CodelessConnectorPollingAuthPropertiesPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsActive
	}).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) Paging() CodelessConnectorPollingPagingPropertiesPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingPagingProperties {
		if v == nil {
			return nil
		}
		return v.Paging
	}).(CodelessConnectorPollingPagingPropertiesPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) Request() CodelessConnectorPollingRequestPropertiesPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingRequestProperties {
		if v == nil {
			return nil
		}
		return &v.Request
	}).(CodelessConnectorPollingRequestPropertiesPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesPtrOutput) Response() CodelessConnectorPollingResponsePropertiesPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigProperties) *CodelessConnectorPollingResponseProperties {
		if v == nil {
			return nil
		}
		return v.Response
	}).(CodelessConnectorPollingResponsePropertiesPtrOutput)
}

type CodelessConnectorPollingConfigPropertiesResponse struct {
	Auth     CodelessConnectorPollingAuthPropertiesResponse      `pulumi:"auth"`
	IsActive *bool                                               `pulumi:"isActive"`
	Paging   *CodelessConnectorPollingPagingPropertiesResponse   `pulumi:"paging"`
	Request  CodelessConnectorPollingRequestPropertiesResponse   `pulumi:"request"`
	Response *CodelessConnectorPollingResponsePropertiesResponse `pulumi:"response"`
}

type CodelessConnectorPollingConfigPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingConfigPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingConfigPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) ToCodelessConnectorPollingConfigPropertiesResponseOutput() CodelessConnectorPollingConfigPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) ToCodelessConnectorPollingConfigPropertiesResponseOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) Auth() CodelessConnectorPollingAuthPropertiesResponseOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigPropertiesResponse) CodelessConnectorPollingAuthPropertiesResponse {
		return v.Auth
	}).(CodelessConnectorPollingAuthPropertiesResponseOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigPropertiesResponse) *bool { return v.IsActive }).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) Paging() CodelessConnectorPollingPagingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigPropertiesResponse) *CodelessConnectorPollingPagingPropertiesResponse {
		return v.Paging
	}).(CodelessConnectorPollingPagingPropertiesResponsePtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) Request() CodelessConnectorPollingRequestPropertiesResponseOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigPropertiesResponse) CodelessConnectorPollingRequestPropertiesResponse {
		return v.Request
	}).(CodelessConnectorPollingRequestPropertiesResponseOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponseOutput) Response() CodelessConnectorPollingResponsePropertiesResponsePtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingConfigPropertiesResponse) *CodelessConnectorPollingResponsePropertiesResponse {
		return v.Response
	}).(CodelessConnectorPollingResponsePropertiesResponsePtrOutput)
}

type CodelessConnectorPollingConfigPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingConfigPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingConfigPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) ToCodelessConnectorPollingConfigPropertiesResponsePtrOutput() CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) ToCodelessConnectorPollingConfigPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessConnectorPollingConfigPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) Elem() CodelessConnectorPollingConfigPropertiesResponseOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigPropertiesResponse) CodelessConnectorPollingConfigPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingConfigPropertiesResponse
		return ret
	}).(CodelessConnectorPollingConfigPropertiesResponseOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) Auth() CodelessConnectorPollingAuthPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigPropertiesResponse) *CodelessConnectorPollingAuthPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.Auth
	}).(CodelessConnectorPollingAuthPropertiesResponsePtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) IsActive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsActive
	}).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) Paging() CodelessConnectorPollingPagingPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigPropertiesResponse) *CodelessConnectorPollingPagingPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Paging
	}).(CodelessConnectorPollingPagingPropertiesResponsePtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) Request() CodelessConnectorPollingRequestPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigPropertiesResponse) *CodelessConnectorPollingRequestPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.Request
	}).(CodelessConnectorPollingRequestPropertiesResponsePtrOutput)
}

func (o CodelessConnectorPollingConfigPropertiesResponsePtrOutput) Response() CodelessConnectorPollingResponsePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingConfigPropertiesResponse) *CodelessConnectorPollingResponsePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Response
	}).(CodelessConnectorPollingResponsePropertiesResponsePtrOutput)
}

type CodelessConnectorPollingPagingProperties struct {
	NextPageParaName                       *string `pulumi:"nextPageParaName"`
	NextPageTokenJsonPath                  *string `pulumi:"nextPageTokenJsonPath"`
	PageCountAttributePath                 *string `pulumi:"pageCountAttributePath"`
	PageSize                               *int    `pulumi:"pageSize"`
	PageSizeParaName                       *string `pulumi:"pageSizeParaName"`
	PageTimeStampAttributePath             *string `pulumi:"pageTimeStampAttributePath"`
	PageTotalCountAttributePath            *string `pulumi:"pageTotalCountAttributePath"`
	PagingType                             string  `pulumi:"pagingType"`
	SearchTheLatestTimeStampFromEventsList *string `pulumi:"searchTheLatestTimeStampFromEventsList"`
}





type CodelessConnectorPollingPagingPropertiesInput interface {
	pulumi.Input

	ToCodelessConnectorPollingPagingPropertiesOutput() CodelessConnectorPollingPagingPropertiesOutput
	ToCodelessConnectorPollingPagingPropertiesOutputWithContext(context.Context) CodelessConnectorPollingPagingPropertiesOutput
}

type CodelessConnectorPollingPagingPropertiesArgs struct {
	NextPageParaName                       pulumi.StringPtrInput `pulumi:"nextPageParaName"`
	NextPageTokenJsonPath                  pulumi.StringPtrInput `pulumi:"nextPageTokenJsonPath"`
	PageCountAttributePath                 pulumi.StringPtrInput `pulumi:"pageCountAttributePath"`
	PageSize                               pulumi.IntPtrInput    `pulumi:"pageSize"`
	PageSizeParaName                       pulumi.StringPtrInput `pulumi:"pageSizeParaName"`
	PageTimeStampAttributePath             pulumi.StringPtrInput `pulumi:"pageTimeStampAttributePath"`
	PageTotalCountAttributePath            pulumi.StringPtrInput `pulumi:"pageTotalCountAttributePath"`
	PagingType                             pulumi.StringInput    `pulumi:"pagingType"`
	SearchTheLatestTimeStampFromEventsList pulumi.StringPtrInput `pulumi:"searchTheLatestTimeStampFromEventsList"`
}

func (CodelessConnectorPollingPagingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingPagingProperties)(nil)).Elem()
}

func (i CodelessConnectorPollingPagingPropertiesArgs) ToCodelessConnectorPollingPagingPropertiesOutput() CodelessConnectorPollingPagingPropertiesOutput {
	return i.ToCodelessConnectorPollingPagingPropertiesOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingPagingPropertiesArgs) ToCodelessConnectorPollingPagingPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingPagingPropertiesOutput)
}

func (i CodelessConnectorPollingPagingPropertiesArgs) ToCodelessConnectorPollingPagingPropertiesPtrOutput() CodelessConnectorPollingPagingPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingPagingPropertiesArgs) ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingPagingPropertiesOutput).ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(ctx)
}









type CodelessConnectorPollingPagingPropertiesPtrInput interface {
	pulumi.Input

	ToCodelessConnectorPollingPagingPropertiesPtrOutput() CodelessConnectorPollingPagingPropertiesPtrOutput
	ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(context.Context) CodelessConnectorPollingPagingPropertiesPtrOutput
}

type codelessConnectorPollingPagingPropertiesPtrType CodelessConnectorPollingPagingPropertiesArgs

func CodelessConnectorPollingPagingPropertiesPtr(v *CodelessConnectorPollingPagingPropertiesArgs) CodelessConnectorPollingPagingPropertiesPtrInput {
	return (*codelessConnectorPollingPagingPropertiesPtrType)(v)
}

func (*codelessConnectorPollingPagingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingPagingProperties)(nil)).Elem()
}

func (i *codelessConnectorPollingPagingPropertiesPtrType) ToCodelessConnectorPollingPagingPropertiesPtrOutput() CodelessConnectorPollingPagingPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(context.Background())
}

func (i *codelessConnectorPollingPagingPropertiesPtrType) ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingPagingPropertiesPtrOutput)
}

type CodelessConnectorPollingPagingPropertiesOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingPagingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingPagingProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingPagingPropertiesOutput) ToCodelessConnectorPollingPagingPropertiesOutput() CodelessConnectorPollingPagingPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesOutput) ToCodelessConnectorPollingPagingPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesOutput) ToCodelessConnectorPollingPagingPropertiesPtrOutput() CodelessConnectorPollingPagingPropertiesPtrOutput {
	return o.ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(context.Background())
}

func (o CodelessConnectorPollingPagingPropertiesOutput) ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessConnectorPollingPagingProperties) *CodelessConnectorPollingPagingProperties {
		return &v
	}).(CodelessConnectorPollingPagingPropertiesPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) NextPageParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string { return v.NextPageParaName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) NextPageTokenJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string { return v.NextPageTokenJsonPath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) PageCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string { return v.PageCountAttributePath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) PageSizeParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string { return v.PageSizeParaName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) PageTimeStampAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string { return v.PageTimeStampAttributePath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) PageTotalCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string { return v.PageTotalCountAttributePath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) PagingType() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) string { return v.PagingType }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingPagingPropertiesOutput) SearchTheLatestTimeStampFromEventsList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingProperties) *string {
		return v.SearchTheLatestTimeStampFromEventsList
	}).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingPagingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingPagingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingPagingProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) ToCodelessConnectorPollingPagingPropertiesPtrOutput() CodelessConnectorPollingPagingPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) ToCodelessConnectorPollingPagingPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) Elem() CodelessConnectorPollingPagingPropertiesOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) CodelessConnectorPollingPagingProperties {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingPagingProperties
		return ret
	}).(CodelessConnectorPollingPagingPropertiesOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) NextPageParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.NextPageParaName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) NextPageTokenJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.NextPageTokenJsonPath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) PageCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.PageCountAttributePath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *int {
		if v == nil {
			return nil
		}
		return v.PageSize
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) PageSizeParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.PageSizeParaName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) PageTimeStampAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.PageTimeStampAttributePath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) PageTotalCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.PageTotalCountAttributePath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) PagingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return &v.PagingType
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesPtrOutput) SearchTheLatestTimeStampFromEventsList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingProperties) *string {
		if v == nil {
			return nil
		}
		return v.SearchTheLatestTimeStampFromEventsList
	}).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingPagingPropertiesResponse struct {
	NextPageParaName                       *string `pulumi:"nextPageParaName"`
	NextPageTokenJsonPath                  *string `pulumi:"nextPageTokenJsonPath"`
	PageCountAttributePath                 *string `pulumi:"pageCountAttributePath"`
	PageSize                               *int    `pulumi:"pageSize"`
	PageSizeParaName                       *string `pulumi:"pageSizeParaName"`
	PageTimeStampAttributePath             *string `pulumi:"pageTimeStampAttributePath"`
	PageTotalCountAttributePath            *string `pulumi:"pageTotalCountAttributePath"`
	PagingType                             string  `pulumi:"pagingType"`
	SearchTheLatestTimeStampFromEventsList *string `pulumi:"searchTheLatestTimeStampFromEventsList"`
}

type CodelessConnectorPollingPagingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingPagingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingPagingPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) ToCodelessConnectorPollingPagingPropertiesResponseOutput() CodelessConnectorPollingPagingPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) ToCodelessConnectorPollingPagingPropertiesResponseOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) NextPageParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string { return v.NextPageParaName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) NextPageTokenJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string { return v.NextPageTokenJsonPath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) PageCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string { return v.PageCountAttributePath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *int { return v.PageSize }).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) PageSizeParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string { return v.PageSizeParaName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) PageTimeStampAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string { return v.PageTimeStampAttributePath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) PageTotalCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string { return v.PageTotalCountAttributePath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) PagingType() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) string { return v.PagingType }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponseOutput) SearchTheLatestTimeStampFromEventsList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingPagingPropertiesResponse) *string {
		return v.SearchTheLatestTimeStampFromEventsList
	}).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingPagingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingPagingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingPagingPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) ToCodelessConnectorPollingPagingPropertiesResponsePtrOutput() CodelessConnectorPollingPagingPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) ToCodelessConnectorPollingPagingPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessConnectorPollingPagingPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) Elem() CodelessConnectorPollingPagingPropertiesResponseOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) CodelessConnectorPollingPagingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingPagingPropertiesResponse
		return ret
	}).(CodelessConnectorPollingPagingPropertiesResponseOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) NextPageParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.NextPageParaName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) NextPageTokenJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.NextPageTokenJsonPath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) PageCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PageCountAttributePath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) PageSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.PageSize
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) PageSizeParaName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PageSizeParaName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) PageTimeStampAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PageTimeStampAttributePath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) PageTotalCountAttributePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PageTotalCountAttributePath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) PagingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PagingType
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingPagingPropertiesResponsePtrOutput) SearchTheLatestTimeStampFromEventsList() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingPagingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SearchTheLatestTimeStampFromEventsList
	}).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingRequestProperties struct {
	ApiEndpoint             string      `pulumi:"apiEndpoint"`
	EndTimeAttributeName    *string     `pulumi:"endTimeAttributeName"`
	Headers                 interface{} `pulumi:"headers"`
	HttpMethod              string      `pulumi:"httpMethod"`
	QueryParameters         interface{} `pulumi:"queryParameters"`
	QueryParametersTemplate *string     `pulumi:"queryParametersTemplate"`
	QueryTimeFormat         string      `pulumi:"queryTimeFormat"`
	QueryWindowInMin        int         `pulumi:"queryWindowInMin"`
	RateLimitQps            *int        `pulumi:"rateLimitQps"`
	RetryCount              *int        `pulumi:"retryCount"`
	StartTimeAttributeName  *string     `pulumi:"startTimeAttributeName"`
	TimeoutInSeconds        *int        `pulumi:"timeoutInSeconds"`
}





type CodelessConnectorPollingRequestPropertiesInput interface {
	pulumi.Input

	ToCodelessConnectorPollingRequestPropertiesOutput() CodelessConnectorPollingRequestPropertiesOutput
	ToCodelessConnectorPollingRequestPropertiesOutputWithContext(context.Context) CodelessConnectorPollingRequestPropertiesOutput
}

type CodelessConnectorPollingRequestPropertiesArgs struct {
	ApiEndpoint             pulumi.StringInput    `pulumi:"apiEndpoint"`
	EndTimeAttributeName    pulumi.StringPtrInput `pulumi:"endTimeAttributeName"`
	Headers                 pulumi.Input          `pulumi:"headers"`
	HttpMethod              pulumi.StringInput    `pulumi:"httpMethod"`
	QueryParameters         pulumi.Input          `pulumi:"queryParameters"`
	QueryParametersTemplate pulumi.StringPtrInput `pulumi:"queryParametersTemplate"`
	QueryTimeFormat         pulumi.StringInput    `pulumi:"queryTimeFormat"`
	QueryWindowInMin        pulumi.IntInput       `pulumi:"queryWindowInMin"`
	RateLimitQps            pulumi.IntPtrInput    `pulumi:"rateLimitQps"`
	RetryCount              pulumi.IntPtrInput    `pulumi:"retryCount"`
	StartTimeAttributeName  pulumi.StringPtrInput `pulumi:"startTimeAttributeName"`
	TimeoutInSeconds        pulumi.IntPtrInput    `pulumi:"timeoutInSeconds"`
}

func (CodelessConnectorPollingRequestPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingRequestProperties)(nil)).Elem()
}

func (i CodelessConnectorPollingRequestPropertiesArgs) ToCodelessConnectorPollingRequestPropertiesOutput() CodelessConnectorPollingRequestPropertiesOutput {
	return i.ToCodelessConnectorPollingRequestPropertiesOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingRequestPropertiesArgs) ToCodelessConnectorPollingRequestPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingRequestPropertiesOutput)
}

func (i CodelessConnectorPollingRequestPropertiesArgs) ToCodelessConnectorPollingRequestPropertiesPtrOutput() CodelessConnectorPollingRequestPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingRequestPropertiesArgs) ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingRequestPropertiesOutput).ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(ctx)
}









type CodelessConnectorPollingRequestPropertiesPtrInput interface {
	pulumi.Input

	ToCodelessConnectorPollingRequestPropertiesPtrOutput() CodelessConnectorPollingRequestPropertiesPtrOutput
	ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(context.Context) CodelessConnectorPollingRequestPropertiesPtrOutput
}

type codelessConnectorPollingRequestPropertiesPtrType CodelessConnectorPollingRequestPropertiesArgs

func CodelessConnectorPollingRequestPropertiesPtr(v *CodelessConnectorPollingRequestPropertiesArgs) CodelessConnectorPollingRequestPropertiesPtrInput {
	return (*codelessConnectorPollingRequestPropertiesPtrType)(v)
}

func (*codelessConnectorPollingRequestPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingRequestProperties)(nil)).Elem()
}

func (i *codelessConnectorPollingRequestPropertiesPtrType) ToCodelessConnectorPollingRequestPropertiesPtrOutput() CodelessConnectorPollingRequestPropertiesPtrOutput {
	return i.ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(context.Background())
}

func (i *codelessConnectorPollingRequestPropertiesPtrType) ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingRequestPropertiesPtrOutput)
}

type CodelessConnectorPollingRequestPropertiesOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingRequestPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingRequestProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingRequestPropertiesOutput) ToCodelessConnectorPollingRequestPropertiesOutput() CodelessConnectorPollingRequestPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesOutput) ToCodelessConnectorPollingRequestPropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesOutput) ToCodelessConnectorPollingRequestPropertiesPtrOutput() CodelessConnectorPollingRequestPropertiesPtrOutput {
	return o.ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(context.Background())
}

func (o CodelessConnectorPollingRequestPropertiesOutput) ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessConnectorPollingRequestProperties) *CodelessConnectorPollingRequestProperties {
		return &v
	}).(CodelessConnectorPollingRequestPropertiesPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) string { return v.ApiEndpoint }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) EndTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) *string { return v.EndTimeAttributeName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) Headers() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) interface{} { return v.Headers }).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) string { return v.HttpMethod }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) QueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) interface{} { return v.QueryParameters }).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) QueryParametersTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) *string { return v.QueryParametersTemplate }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) QueryTimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) string { return v.QueryTimeFormat }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) QueryWindowInMin() pulumi.IntOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) int { return v.QueryWindowInMin }).(pulumi.IntOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) RateLimitQps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) *int { return v.RateLimitQps }).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) *int { return v.RetryCount }).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) StartTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) *string { return v.StartTimeAttributeName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestProperties) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

type CodelessConnectorPollingRequestPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingRequestPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingRequestProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) ToCodelessConnectorPollingRequestPropertiesPtrOutput() CodelessConnectorPollingRequestPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) ToCodelessConnectorPollingRequestPropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) Elem() CodelessConnectorPollingRequestPropertiesOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) CodelessConnectorPollingRequestProperties {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingRequestProperties
		return ret
	}).(CodelessConnectorPollingRequestPropertiesOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) ApiEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *string {
		if v == nil {
			return nil
		}
		return &v.ApiEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) EndTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.EndTimeAttributeName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) Headers() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) HttpMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *string {
		if v == nil {
			return nil
		}
		return &v.HttpMethod
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) QueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) interface{} {
		if v == nil {
			return nil
		}
		return v.QueryParameters
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) QueryParametersTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.QueryParametersTemplate
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) QueryTimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *string {
		if v == nil {
			return nil
		}
		return &v.QueryTimeFormat
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) QueryWindowInMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *int {
		if v == nil {
			return nil
		}
		return &v.QueryWindowInMin
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) RateLimitQps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *int {
		if v == nil {
			return nil
		}
		return v.RateLimitQps
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *int {
		if v == nil {
			return nil
		}
		return v.RetryCount
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) StartTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *string {
		if v == nil {
			return nil
		}
		return v.StartTimeAttributeName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesPtrOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestProperties) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.IntPtrOutput)
}

type CodelessConnectorPollingRequestPropertiesResponse struct {
	ApiEndpoint             string      `pulumi:"apiEndpoint"`
	EndTimeAttributeName    *string     `pulumi:"endTimeAttributeName"`
	Headers                 interface{} `pulumi:"headers"`
	HttpMethod              string      `pulumi:"httpMethod"`
	QueryParameters         interface{} `pulumi:"queryParameters"`
	QueryParametersTemplate *string     `pulumi:"queryParametersTemplate"`
	QueryTimeFormat         string      `pulumi:"queryTimeFormat"`
	QueryWindowInMin        int         `pulumi:"queryWindowInMin"`
	RateLimitQps            *int        `pulumi:"rateLimitQps"`
	RetryCount              *int        `pulumi:"retryCount"`
	StartTimeAttributeName  *string     `pulumi:"startTimeAttributeName"`
	TimeoutInSeconds        *int        `pulumi:"timeoutInSeconds"`
}

type CodelessConnectorPollingRequestPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingRequestPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingRequestPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) ToCodelessConnectorPollingRequestPropertiesResponseOutput() CodelessConnectorPollingRequestPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) ToCodelessConnectorPollingRequestPropertiesResponseOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) ApiEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) string { return v.ApiEndpoint }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) EndTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) *string { return v.EndTimeAttributeName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) Headers() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) interface{} { return v.Headers }).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) HttpMethod() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) string { return v.HttpMethod }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) QueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) interface{} { return v.QueryParameters }).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) QueryParametersTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) *string { return v.QueryParametersTemplate }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) QueryTimeFormat() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) string { return v.QueryTimeFormat }).(pulumi.StringOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) QueryWindowInMin() pulumi.IntOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) int { return v.QueryWindowInMin }).(pulumi.IntOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) RateLimitQps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) *int { return v.RateLimitQps }).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) *int { return v.RetryCount }).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) StartTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) *string { return v.StartTimeAttributeName }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponseOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingRequestPropertiesResponse) *int { return v.TimeoutInSeconds }).(pulumi.IntPtrOutput)
}

type CodelessConnectorPollingRequestPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingRequestPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingRequestPropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) ToCodelessConnectorPollingRequestPropertiesResponsePtrOutput() CodelessConnectorPollingRequestPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) ToCodelessConnectorPollingRequestPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessConnectorPollingRequestPropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) Elem() CodelessConnectorPollingRequestPropertiesResponseOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) CodelessConnectorPollingRequestPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingRequestPropertiesResponse
		return ret
	}).(CodelessConnectorPollingRequestPropertiesResponseOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) ApiEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ApiEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) EndTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.EndTimeAttributeName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) Headers() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) HttpMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.HttpMethod
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) QueryParameters() pulumi.AnyOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) interface{} {
		if v == nil {
			return nil
		}
		return v.QueryParameters
	}).(pulumi.AnyOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) QueryParametersTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueryParametersTemplate
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) QueryTimeFormat() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.QueryTimeFormat
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) QueryWindowInMin() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return &v.QueryWindowInMin
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) RateLimitQps() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.RateLimitQps
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) RetryCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetryCount
	}).(pulumi.IntPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) StartTimeAttributeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTimeAttributeName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingRequestPropertiesResponsePtrOutput) TimeoutInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingRequestPropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.TimeoutInSeconds
	}).(pulumi.IntPtrOutput)
}

type CodelessConnectorPollingResponseProperties struct {
	EventsJsonPaths       []string `pulumi:"eventsJsonPaths"`
	IsGzipCompressed      *bool    `pulumi:"isGzipCompressed"`
	SuccessStatusJsonPath *string  `pulumi:"successStatusJsonPath"`
	SuccessStatusValue    *string  `pulumi:"successStatusValue"`
}





type CodelessConnectorPollingResponsePropertiesInput interface {
	pulumi.Input

	ToCodelessConnectorPollingResponsePropertiesOutput() CodelessConnectorPollingResponsePropertiesOutput
	ToCodelessConnectorPollingResponsePropertiesOutputWithContext(context.Context) CodelessConnectorPollingResponsePropertiesOutput
}

type CodelessConnectorPollingResponsePropertiesArgs struct {
	EventsJsonPaths       pulumi.StringArrayInput `pulumi:"eventsJsonPaths"`
	IsGzipCompressed      pulumi.BoolPtrInput     `pulumi:"isGzipCompressed"`
	SuccessStatusJsonPath pulumi.StringPtrInput   `pulumi:"successStatusJsonPath"`
	SuccessStatusValue    pulumi.StringPtrInput   `pulumi:"successStatusValue"`
}

func (CodelessConnectorPollingResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingResponseProperties)(nil)).Elem()
}

func (i CodelessConnectorPollingResponsePropertiesArgs) ToCodelessConnectorPollingResponsePropertiesOutput() CodelessConnectorPollingResponsePropertiesOutput {
	return i.ToCodelessConnectorPollingResponsePropertiesOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingResponsePropertiesArgs) ToCodelessConnectorPollingResponsePropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingResponsePropertiesOutput)
}

func (i CodelessConnectorPollingResponsePropertiesArgs) ToCodelessConnectorPollingResponsePropertiesPtrOutput() CodelessConnectorPollingResponsePropertiesPtrOutput {
	return i.ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i CodelessConnectorPollingResponsePropertiesArgs) ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingResponsePropertiesOutput).ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(ctx)
}









type CodelessConnectorPollingResponsePropertiesPtrInput interface {
	pulumi.Input

	ToCodelessConnectorPollingResponsePropertiesPtrOutput() CodelessConnectorPollingResponsePropertiesPtrOutput
	ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(context.Context) CodelessConnectorPollingResponsePropertiesPtrOutput
}

type codelessConnectorPollingResponsePropertiesPtrType CodelessConnectorPollingResponsePropertiesArgs

func CodelessConnectorPollingResponsePropertiesPtr(v *CodelessConnectorPollingResponsePropertiesArgs) CodelessConnectorPollingResponsePropertiesPtrInput {
	return (*codelessConnectorPollingResponsePropertiesPtrType)(v)
}

func (*codelessConnectorPollingResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingResponseProperties)(nil)).Elem()
}

func (i *codelessConnectorPollingResponsePropertiesPtrType) ToCodelessConnectorPollingResponsePropertiesPtrOutput() CodelessConnectorPollingResponsePropertiesPtrOutput {
	return i.ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *codelessConnectorPollingResponsePropertiesPtrType) ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessConnectorPollingResponsePropertiesPtrOutput)
}

type CodelessConnectorPollingResponsePropertiesOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingResponseProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingResponsePropertiesOutput) ToCodelessConnectorPollingResponsePropertiesOutput() CodelessConnectorPollingResponsePropertiesOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesOutput) ToCodelessConnectorPollingResponsePropertiesOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesOutput) ToCodelessConnectorPollingResponsePropertiesPtrOutput() CodelessConnectorPollingResponsePropertiesPtrOutput {
	return o.ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o CodelessConnectorPollingResponsePropertiesOutput) ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessConnectorPollingResponseProperties) *CodelessConnectorPollingResponseProperties {
		return &v
	}).(CodelessConnectorPollingResponsePropertiesPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesOutput) EventsJsonPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponseProperties) []string { return v.EventsJsonPaths }).(pulumi.StringArrayOutput)
}

func (o CodelessConnectorPollingResponsePropertiesOutput) IsGzipCompressed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponseProperties) *bool { return v.IsGzipCompressed }).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesOutput) SuccessStatusJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponseProperties) *string { return v.SuccessStatusJsonPath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesOutput) SuccessStatusValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponseProperties) *string { return v.SuccessStatusValue }).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingResponseProperties)(nil)).Elem()
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) ToCodelessConnectorPollingResponsePropertiesPtrOutput() CodelessConnectorPollingResponsePropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) ToCodelessConnectorPollingResponsePropertiesPtrOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesPtrOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) Elem() CodelessConnectorPollingResponsePropertiesOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponseProperties) CodelessConnectorPollingResponseProperties {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingResponseProperties
		return ret
	}).(CodelessConnectorPollingResponsePropertiesOutput)
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) EventsJsonPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponseProperties) []string {
		if v == nil {
			return nil
		}
		return v.EventsJsonPaths
	}).(pulumi.StringArrayOutput)
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) IsGzipCompressed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsGzipCompressed
	}).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) SuccessStatusJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.SuccessStatusJsonPath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesPtrOutput) SuccessStatusValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.SuccessStatusValue
	}).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingResponsePropertiesResponse struct {
	EventsJsonPaths       []string `pulumi:"eventsJsonPaths"`
	IsGzipCompressed      *bool    `pulumi:"isGzipCompressed"`
	SuccessStatusJsonPath *string  `pulumi:"successStatusJsonPath"`
	SuccessStatusValue    *string  `pulumi:"successStatusValue"`
}

type CodelessConnectorPollingResponsePropertiesResponseOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingResponsePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessConnectorPollingResponsePropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingResponsePropertiesResponseOutput) ToCodelessConnectorPollingResponsePropertiesResponseOutput() CodelessConnectorPollingResponsePropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesResponseOutput) ToCodelessConnectorPollingResponsePropertiesResponseOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesResponseOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesResponseOutput) EventsJsonPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponsePropertiesResponse) []string { return v.EventsJsonPaths }).(pulumi.StringArrayOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponseOutput) IsGzipCompressed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponsePropertiesResponse) *bool { return v.IsGzipCompressed }).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponseOutput) SuccessStatusJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponsePropertiesResponse) *string { return v.SuccessStatusJsonPath }).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponseOutput) SuccessStatusValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessConnectorPollingResponsePropertiesResponse) *string { return v.SuccessStatusValue }).(pulumi.StringPtrOutput)
}

type CodelessConnectorPollingResponsePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CodelessConnectorPollingResponsePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessConnectorPollingResponsePropertiesResponse)(nil)).Elem()
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) ToCodelessConnectorPollingResponsePropertiesResponsePtrOutput() CodelessConnectorPollingResponsePropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) ToCodelessConnectorPollingResponsePropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessConnectorPollingResponsePropertiesResponsePtrOutput {
	return o
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) Elem() CodelessConnectorPollingResponsePropertiesResponseOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponsePropertiesResponse) CodelessConnectorPollingResponsePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CodelessConnectorPollingResponsePropertiesResponse
		return ret
	}).(CodelessConnectorPollingResponsePropertiesResponseOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) EventsJsonPaths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponsePropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.EventsJsonPaths
	}).(pulumi.StringArrayOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) IsGzipCompressed() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponsePropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsGzipCompressed
	}).(pulumi.BoolPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) SuccessStatusJsonPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponsePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SuccessStatusJsonPath
	}).(pulumi.StringPtrOutput)
}

func (o CodelessConnectorPollingResponsePropertiesResponsePtrOutput) SuccessStatusValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessConnectorPollingResponsePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SuccessStatusValue
	}).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigProperties struct {
	Availability          Availability                                              `pulumi:"availability"`
	ConnectivityCriteria  []CodelessUiConnectorConfigPropertiesConnectivityCriteria `pulumi:"connectivityCriteria"`
	CustomImage           *string                                                   `pulumi:"customImage"`
	DataTypes             []CodelessUiConnectorConfigPropertiesDataTypes            `pulumi:"dataTypes"`
	DescriptionMarkdown   string                                                    `pulumi:"descriptionMarkdown"`
	GraphQueries          []CodelessUiConnectorConfigPropertiesGraphQueries         `pulumi:"graphQueries"`
	GraphQueriesTableName string                                                    `pulumi:"graphQueriesTableName"`
	InstructionSteps      []CodelessUiConnectorConfigPropertiesInstructionSteps     `pulumi:"instructionSteps"`
	Permissions           Permissions                                               `pulumi:"permissions"`
	Publisher             string                                                    `pulumi:"publisher"`
	SampleQueries         []CodelessUiConnectorConfigPropertiesSampleQueries        `pulumi:"sampleQueries"`
	Title                 string                                                    `pulumi:"title"`
}





type CodelessUiConnectorConfigPropertiesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesOutput() CodelessUiConnectorConfigPropertiesOutput
	ToCodelessUiConnectorConfigPropertiesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesOutput
}

type CodelessUiConnectorConfigPropertiesArgs struct {
	Availability          AvailabilityInput                                                 `pulumi:"availability"`
	ConnectivityCriteria  CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayInput `pulumi:"connectivityCriteria"`
	CustomImage           pulumi.StringPtrInput                                             `pulumi:"customImage"`
	DataTypes             CodelessUiConnectorConfigPropertiesDataTypesArrayInput            `pulumi:"dataTypes"`
	DescriptionMarkdown   pulumi.StringInput                                                `pulumi:"descriptionMarkdown"`
	GraphQueries          CodelessUiConnectorConfigPropertiesGraphQueriesArrayInput         `pulumi:"graphQueries"`
	GraphQueriesTableName pulumi.StringInput                                                `pulumi:"graphQueriesTableName"`
	InstructionSteps      CodelessUiConnectorConfigPropertiesInstructionStepsArrayInput     `pulumi:"instructionSteps"`
	Permissions           PermissionsInput                                                  `pulumi:"permissions"`
	Publisher             pulumi.StringInput                                                `pulumi:"publisher"`
	SampleQueries         CodelessUiConnectorConfigPropertiesSampleQueriesArrayInput        `pulumi:"sampleQueries"`
	Title                 pulumi.StringInput                                                `pulumi:"title"`
}

func (CodelessUiConnectorConfigPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigProperties)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesArgs) ToCodelessUiConnectorConfigPropertiesOutput() CodelessUiConnectorConfigPropertiesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesArgs) ToCodelessUiConnectorConfigPropertiesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesOutput)
}

func (i CodelessUiConnectorConfigPropertiesArgs) ToCodelessUiConnectorConfigPropertiesPtrOutput() CodelessUiConnectorConfigPropertiesPtrOutput {
	return i.ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesArgs) ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesOutput).ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(ctx)
}









type CodelessUiConnectorConfigPropertiesPtrInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesPtrOutput() CodelessUiConnectorConfigPropertiesPtrOutput
	ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesPtrOutput
}

type codelessUiConnectorConfigPropertiesPtrType CodelessUiConnectorConfigPropertiesArgs

func CodelessUiConnectorConfigPropertiesPtr(v *CodelessUiConnectorConfigPropertiesArgs) CodelessUiConnectorConfigPropertiesPtrInput {
	return (*codelessUiConnectorConfigPropertiesPtrType)(v)
}

func (*codelessUiConnectorConfigPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiConnectorConfigProperties)(nil)).Elem()
}

func (i *codelessUiConnectorConfigPropertiesPtrType) ToCodelessUiConnectorConfigPropertiesPtrOutput() CodelessUiConnectorConfigPropertiesPtrOutput {
	return i.ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(context.Background())
}

func (i *codelessUiConnectorConfigPropertiesPtrType) ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesPtrOutput)
}

type CodelessUiConnectorConfigPropertiesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigProperties)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesOutput) ToCodelessUiConnectorConfigPropertiesOutput() CodelessUiConnectorConfigPropertiesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesOutput) ToCodelessUiConnectorConfigPropertiesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesOutput) ToCodelessUiConnectorConfigPropertiesPtrOutput() CodelessUiConnectorConfigPropertiesPtrOutput {
	return o.ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(context.Background())
}

func (o CodelessUiConnectorConfigPropertiesOutput) ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessUiConnectorConfigProperties) *CodelessUiConnectorConfigProperties {
		return &v
	}).(CodelessUiConnectorConfigPropertiesPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) Availability() AvailabilityOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) Availability { return v.Availability }).(AvailabilityOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) ConnectivityCriteria() CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesConnectivityCriteria {
		return v.ConnectivityCriteria
	}).(CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) CustomImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) *string { return v.CustomImage }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) DataTypes() CodelessUiConnectorConfigPropertiesDataTypesArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesDataTypes {
		return v.DataTypes
	}).(CodelessUiConnectorConfigPropertiesDataTypesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) DescriptionMarkdown() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) string { return v.DescriptionMarkdown }).(pulumi.StringOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) GraphQueries() CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesGraphQueries {
		return v.GraphQueries
	}).(CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) GraphQueriesTableName() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) string { return v.GraphQueriesTableName }).(pulumi.StringOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) InstructionSteps() CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesInstructionSteps {
		return v.InstructionSteps
	}).(CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) Permissions() PermissionsOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) Permissions { return v.Permissions }).(PermissionsOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) SampleQueries() CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesSampleQueries {
		return v.SampleQueries
	}).(CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigProperties) string { return v.Title }).(pulumi.StringOutput)
}

type CodelessUiConnectorConfigPropertiesPtrOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiConnectorConfigProperties)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) ToCodelessUiConnectorConfigPropertiesPtrOutput() CodelessUiConnectorConfigPropertiesPtrOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) ToCodelessUiConnectorConfigPropertiesPtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesPtrOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) Elem() CodelessUiConnectorConfigPropertiesOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) CodelessUiConnectorConfigProperties {
		if v != nil {
			return *v
		}
		var ret CodelessUiConnectorConfigProperties
		return ret
	}).(CodelessUiConnectorConfigPropertiesOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) Availability() AvailabilityPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *Availability {
		if v == nil {
			return nil
		}
		return &v.Availability
	}).(AvailabilityPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) ConnectivityCriteria() CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesConnectivityCriteria {
		if v == nil {
			return nil
		}
		return v.ConnectivityCriteria
	}).(CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) CustomImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomImage
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) DataTypes() CodelessUiConnectorConfigPropertiesDataTypesArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesDataTypes {
		if v == nil {
			return nil
		}
		return v.DataTypes
	}).(CodelessUiConnectorConfigPropertiesDataTypesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) DescriptionMarkdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DescriptionMarkdown
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) GraphQueries() CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesGraphQueries {
		if v == nil {
			return nil
		}
		return v.GraphQueries
	}).(CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) GraphQueriesTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *string {
		if v == nil {
			return nil
		}
		return &v.GraphQueriesTableName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) InstructionSteps() CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesInstructionSteps {
		if v == nil {
			return nil
		}
		return v.InstructionSteps
	}).(CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) Permissions() PermissionsPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *Permissions {
		if v == nil {
			return nil
		}
		return &v.Permissions
	}).(PermissionsPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) SampleQueries() CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) []CodelessUiConnectorConfigPropertiesSampleQueries {
		if v == nil {
			return nil
		}
		return v.SampleQueries
	}).(CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesPtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesConnectivityCriteria struct {
	Type  *string  `pulumi:"type"`
	Value []string `pulumi:"value"`
}





type CodelessUiConnectorConfigPropertiesConnectivityCriteriaInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput() CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput
	ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput
}

type CodelessUiConnectorConfigPropertiesConnectivityCriteriaArgs struct {
	Type  pulumi.StringPtrInput   `pulumi:"type"`
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (CodelessUiConnectorConfigPropertiesConnectivityCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesConnectivityCriteria)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesConnectivityCriteriaArgs) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput() CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput {
	return i.ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesConnectivityCriteriaArgs) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput)
}





type CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput() CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput
	ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput
}

type CodelessUiConnectorConfigPropertiesConnectivityCriteriaArray []CodelessUiConnectorConfigPropertiesConnectivityCriteriaInput

func (CodelessUiConnectorConfigPropertiesConnectivityCriteriaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesConnectivityCriteria)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesConnectivityCriteriaArray) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput() CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesConnectivityCriteriaArray) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput)
}

type CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesConnectivityCriteria)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput() CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesConnectivityCriteria) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesConnectivityCriteria) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesConnectivityCriteria)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput() CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput) ToCodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesConnectivityCriteria {
		return vs[0].([]CodelessUiConnectorConfigPropertiesConnectivityCriteria)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput)
}

type CodelessUiConnectorConfigPropertiesDataTypes struct {
	LastDataReceivedQuery *string `pulumi:"lastDataReceivedQuery"`
	Name                  *string `pulumi:"name"`
}





type CodelessUiConnectorConfigPropertiesDataTypesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesDataTypesOutput() CodelessUiConnectorConfigPropertiesDataTypesOutput
	ToCodelessUiConnectorConfigPropertiesDataTypesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesDataTypesOutput
}

type CodelessUiConnectorConfigPropertiesDataTypesArgs struct {
	LastDataReceivedQuery pulumi.StringPtrInput `pulumi:"lastDataReceivedQuery"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
}

func (CodelessUiConnectorConfigPropertiesDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesDataTypes)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesDataTypesArgs) ToCodelessUiConnectorConfigPropertiesDataTypesOutput() CodelessUiConnectorConfigPropertiesDataTypesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesDataTypesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesDataTypesArgs) ToCodelessUiConnectorConfigPropertiesDataTypesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesDataTypesOutput)
}





type CodelessUiConnectorConfigPropertiesDataTypesArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutput() CodelessUiConnectorConfigPropertiesDataTypesArrayOutput
	ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesDataTypesArrayOutput
}

type CodelessUiConnectorConfigPropertiesDataTypesArray []CodelessUiConnectorConfigPropertiesDataTypesInput

func (CodelessUiConnectorConfigPropertiesDataTypesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesDataTypes)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesDataTypesArray) ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutput() CodelessUiConnectorConfigPropertiesDataTypesArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesDataTypesArray) ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesDataTypesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesDataTypesArrayOutput)
}

type CodelessUiConnectorConfigPropertiesDataTypesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesDataTypes)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesDataTypesOutput) ToCodelessUiConnectorConfigPropertiesDataTypesOutput() CodelessUiConnectorConfigPropertiesDataTypesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesDataTypesOutput) ToCodelessUiConnectorConfigPropertiesDataTypesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesDataTypesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesDataTypesOutput) LastDataReceivedQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesDataTypes) *string { return v.LastDataReceivedQuery }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesDataTypesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesDataTypes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesDataTypesArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesDataTypesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesDataTypes)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesDataTypesArrayOutput) ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutput() CodelessUiConnectorConfigPropertiesDataTypesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesDataTypesArrayOutput) ToCodelessUiConnectorConfigPropertiesDataTypesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesDataTypesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesDataTypesArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesDataTypesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesDataTypes {
		return vs[0].([]CodelessUiConnectorConfigPropertiesDataTypes)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesDataTypesOutput)
}

type CodelessUiConnectorConfigPropertiesGraphQueries struct {
	BaseQuery  *string `pulumi:"baseQuery"`
	Legend     *string `pulumi:"legend"`
	MetricName *string `pulumi:"metricName"`
}





type CodelessUiConnectorConfigPropertiesGraphQueriesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesGraphQueriesOutput() CodelessUiConnectorConfigPropertiesGraphQueriesOutput
	ToCodelessUiConnectorConfigPropertiesGraphQueriesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesGraphQueriesOutput
}

type CodelessUiConnectorConfigPropertiesGraphQueriesArgs struct {
	BaseQuery  pulumi.StringPtrInput `pulumi:"baseQuery"`
	Legend     pulumi.StringPtrInput `pulumi:"legend"`
	MetricName pulumi.StringPtrInput `pulumi:"metricName"`
}

func (CodelessUiConnectorConfigPropertiesGraphQueriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesGraphQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesGraphQueriesArgs) ToCodelessUiConnectorConfigPropertiesGraphQueriesOutput() CodelessUiConnectorConfigPropertiesGraphQueriesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesGraphQueriesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesGraphQueriesArgs) ToCodelessUiConnectorConfigPropertiesGraphQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesGraphQueriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesGraphQueriesOutput)
}





type CodelessUiConnectorConfigPropertiesGraphQueriesArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput() CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput
	ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput
}

type CodelessUiConnectorConfigPropertiesGraphQueriesArray []CodelessUiConnectorConfigPropertiesGraphQueriesInput

func (CodelessUiConnectorConfigPropertiesGraphQueriesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesGraphQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesGraphQueriesArray) ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput() CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesGraphQueriesArray) ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput)
}

type CodelessUiConnectorConfigPropertiesGraphQueriesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesGraphQueriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesGraphQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesOutput) ToCodelessUiConnectorConfigPropertiesGraphQueriesOutput() CodelessUiConnectorConfigPropertiesGraphQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesOutput) ToCodelessUiConnectorConfigPropertiesGraphQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesGraphQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesOutput) BaseQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesGraphQueries) *string { return v.BaseQuery }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesOutput) Legend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesGraphQueries) *string { return v.Legend }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesGraphQueries) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesGraphQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput() CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesGraphQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesGraphQueriesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesGraphQueries {
		return vs[0].([]CodelessUiConnectorConfigPropertiesGraphQueries)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesGraphQueriesOutput)
}

type CodelessUiConnectorConfigPropertiesInstructionSteps struct {
	Description  *string                        `pulumi:"description"`
	Instructions []InstructionStepsInstructions `pulumi:"instructions"`
	Title        *string                        `pulumi:"title"`
}





type CodelessUiConnectorConfigPropertiesInstructionStepsInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesInstructionStepsOutput() CodelessUiConnectorConfigPropertiesInstructionStepsOutput
	ToCodelessUiConnectorConfigPropertiesInstructionStepsOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesInstructionStepsOutput
}

type CodelessUiConnectorConfigPropertiesInstructionStepsArgs struct {
	Description  pulumi.StringPtrInput                  `pulumi:"description"`
	Instructions InstructionStepsInstructionsArrayInput `pulumi:"instructions"`
	Title        pulumi.StringPtrInput                  `pulumi:"title"`
}

func (CodelessUiConnectorConfigPropertiesInstructionStepsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesInstructionSteps)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesInstructionStepsArgs) ToCodelessUiConnectorConfigPropertiesInstructionStepsOutput() CodelessUiConnectorConfigPropertiesInstructionStepsOutput {
	return i.ToCodelessUiConnectorConfigPropertiesInstructionStepsOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesInstructionStepsArgs) ToCodelessUiConnectorConfigPropertiesInstructionStepsOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesInstructionStepsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesInstructionStepsOutput)
}





type CodelessUiConnectorConfigPropertiesInstructionStepsArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput() CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput
	ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput
}

type CodelessUiConnectorConfigPropertiesInstructionStepsArray []CodelessUiConnectorConfigPropertiesInstructionStepsInput

func (CodelessUiConnectorConfigPropertiesInstructionStepsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesInstructionSteps)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesInstructionStepsArray) ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput() CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesInstructionStepsArray) ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput)
}

type CodelessUiConnectorConfigPropertiesInstructionStepsOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesInstructionStepsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesInstructionSteps)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsOutput) ToCodelessUiConnectorConfigPropertiesInstructionStepsOutput() CodelessUiConnectorConfigPropertiesInstructionStepsOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsOutput) ToCodelessUiConnectorConfigPropertiesInstructionStepsOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesInstructionStepsOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesInstructionSteps) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsOutput) Instructions() InstructionStepsInstructionsArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesInstructionSteps) []InstructionStepsInstructions {
		return v.Instructions
	}).(InstructionStepsInstructionsArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesInstructionSteps) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesInstructionSteps)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput) ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput() CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput) ToCodelessUiConnectorConfigPropertiesInstructionStepsArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesInstructionStepsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesInstructionSteps {
		return vs[0].([]CodelessUiConnectorConfigPropertiesInstructionSteps)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesInstructionStepsOutput)
}

type CodelessUiConnectorConfigPropertiesResponse struct {
	Availability          AvailabilityResponse                                              `pulumi:"availability"`
	ConnectivityCriteria  []CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria `pulumi:"connectivityCriteria"`
	CustomImage           *string                                                           `pulumi:"customImage"`
	DataTypes             []CodelessUiConnectorConfigPropertiesResponseDataTypes            `pulumi:"dataTypes"`
	DescriptionMarkdown   string                                                            `pulumi:"descriptionMarkdown"`
	GraphQueries          []CodelessUiConnectorConfigPropertiesResponseGraphQueries         `pulumi:"graphQueries"`
	GraphQueriesTableName string                                                            `pulumi:"graphQueriesTableName"`
	InstructionSteps      []CodelessUiConnectorConfigPropertiesResponseInstructionSteps     `pulumi:"instructionSteps"`
	Permissions           PermissionsResponse                                               `pulumi:"permissions"`
	Publisher             string                                                            `pulumi:"publisher"`
	SampleQueries         []CodelessUiConnectorConfigPropertiesResponseSampleQueries        `pulumi:"sampleQueries"`
	Title                 string                                                            `pulumi:"title"`
}

type CodelessUiConnectorConfigPropertiesResponseOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponse)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) ToCodelessUiConnectorConfigPropertiesResponseOutput() CodelessUiConnectorConfigPropertiesResponseOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) ToCodelessUiConnectorConfigPropertiesResponseOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) Availability() AvailabilityResponseOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) AvailabilityResponse { return v.Availability }).(AvailabilityResponseOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) ConnectivityCriteria() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria {
		return v.ConnectivityCriteria
	}).(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) CustomImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) *string { return v.CustomImage }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) DataTypes() CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseDataTypes {
		return v.DataTypes
	}).(CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) DescriptionMarkdown() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) string { return v.DescriptionMarkdown }).(pulumi.StringOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) GraphQueries() CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseGraphQueries {
		return v.GraphQueries
	}).(CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) GraphQueriesTableName() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) string { return v.GraphQueriesTableName }).(pulumi.StringOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) InstructionSteps() CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseInstructionSteps {
		return v.InstructionSteps
	}).(CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) Permissions() PermissionsResponseOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) PermissionsResponse { return v.Permissions }).(PermissionsResponseOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) Publisher() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) string { return v.Publisher }).(pulumi.StringOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) SampleQueries() CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseSampleQueries {
		return v.SampleQueries
	}).(CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponse) string { return v.Title }).(pulumi.StringOutput)
}

type CodelessUiConnectorConfigPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiConnectorConfigPropertiesResponse)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) ToCodelessUiConnectorConfigPropertiesResponsePtrOutput() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) Elem() CodelessUiConnectorConfigPropertiesResponseOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) CodelessUiConnectorConfigPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret CodelessUiConnectorConfigPropertiesResponse
		return ret
	}).(CodelessUiConnectorConfigPropertiesResponseOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) Availability() AvailabilityResponsePtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *AvailabilityResponse {
		if v == nil {
			return nil
		}
		return &v.Availability
	}).(AvailabilityResponsePtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) ConnectivityCriteria() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria {
		if v == nil {
			return nil
		}
		return v.ConnectivityCriteria
	}).(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) CustomImage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomImage
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) DataTypes() CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseDataTypes {
		if v == nil {
			return nil
		}
		return v.DataTypes
	}).(CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) DescriptionMarkdown() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DescriptionMarkdown
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) GraphQueries() CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseGraphQueries {
		if v == nil {
			return nil
		}
		return v.GraphQueries
	}).(CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) GraphQueriesTableName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GraphQueriesTableName
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) InstructionSteps() CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseInstructionSteps {
		if v == nil {
			return nil
		}
		return v.InstructionSteps
	}).(CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) Permissions() PermissionsResponsePtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *PermissionsResponse {
		if v == nil {
			return nil
		}
		return &v.Permissions
	}).(PermissionsResponsePtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) SampleQueries() CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) []CodelessUiConnectorConfigPropertiesResponseSampleQueries {
		if v == nil {
			return nil
		}
		return v.SampleQueries
	}).(CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponsePtrOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CodelessUiConnectorConfigPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Title
	}).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria struct {
	Type  *string  `pulumi:"type"`
	Value []string `pulumi:"value"`
}

type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput) Value() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria) []string { return v.Value }).(pulumi.StringArrayOutput)
}

type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria {
		return vs[0].([]CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput)
}

type CodelessUiConnectorConfigPropertiesResponseDataTypes struct {
	LastDataReceivedQuery *string `pulumi:"lastDataReceivedQuery"`
	Name                  *string `pulumi:"name"`
}

type CodelessUiConnectorConfigPropertiesResponseDataTypesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseDataTypes)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesOutput) ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutput() CodelessUiConnectorConfigPropertiesResponseDataTypesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesOutput) ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseDataTypesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesOutput) LastDataReceivedQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseDataTypes) *string { return v.LastDataReceivedQuery }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseDataTypes) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseDataTypes)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput() CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesResponseDataTypesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesResponseDataTypes {
		return vs[0].([]CodelessUiConnectorConfigPropertiesResponseDataTypes)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesResponseDataTypesOutput)
}

type CodelessUiConnectorConfigPropertiesResponseGraphQueries struct {
	BaseQuery  *string `pulumi:"baseQuery"`
	Legend     *string `pulumi:"legend"`
	MetricName *string `pulumi:"metricName"`
}

type CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseGraphQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput() CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput) BaseQuery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseGraphQueries) *string { return v.BaseQuery }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput) Legend() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseGraphQueries) *string { return v.Legend }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput) MetricName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseGraphQueries) *string { return v.MetricName }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseGraphQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput() CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesResponseGraphQueries {
		return vs[0].([]CodelessUiConnectorConfigPropertiesResponseGraphQueries)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput)
}

type CodelessUiConnectorConfigPropertiesResponseInstructionSteps struct {
	Description  *string                                `pulumi:"description"`
	Instructions []InstructionStepsResponseInstructions `pulumi:"instructions"`
	Title        *string                                `pulumi:"title"`
}

type CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseInstructionSteps)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput() CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseInstructionSteps) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput) Instructions() InstructionStepsResponseInstructionsArrayOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseInstructionSteps) []InstructionStepsResponseInstructions {
		return v.Instructions
	}).(InstructionStepsResponseInstructionsArrayOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput) Title() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseInstructionSteps) *string { return v.Title }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseInstructionSteps)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput() CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesResponseInstructionSteps {
		return vs[0].([]CodelessUiConnectorConfigPropertiesResponseInstructionSteps)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput)
}

type CodelessUiConnectorConfigPropertiesResponseSampleQueries struct {
	Description *string `pulumi:"description"`
	Query       *string `pulumi:"query"`
}

type CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseSampleQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput() CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseSampleQueries) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesResponseSampleQueries) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseSampleQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput() CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesResponseSampleQueries {
		return vs[0].([]CodelessUiConnectorConfigPropertiesResponseSampleQueries)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput)
}

type CodelessUiConnectorConfigPropertiesSampleQueries struct {
	Description *string `pulumi:"description"`
	Query       *string `pulumi:"query"`
}





type CodelessUiConnectorConfigPropertiesSampleQueriesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesSampleQueriesOutput() CodelessUiConnectorConfigPropertiesSampleQueriesOutput
	ToCodelessUiConnectorConfigPropertiesSampleQueriesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesSampleQueriesOutput
}

type CodelessUiConnectorConfigPropertiesSampleQueriesArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Query       pulumi.StringPtrInput `pulumi:"query"`
}

func (CodelessUiConnectorConfigPropertiesSampleQueriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesSampleQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesSampleQueriesArgs) ToCodelessUiConnectorConfigPropertiesSampleQueriesOutput() CodelessUiConnectorConfigPropertiesSampleQueriesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesSampleQueriesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesSampleQueriesArgs) ToCodelessUiConnectorConfigPropertiesSampleQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesSampleQueriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesSampleQueriesOutput)
}





type CodelessUiConnectorConfigPropertiesSampleQueriesArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput() CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput
	ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput
}

type CodelessUiConnectorConfigPropertiesSampleQueriesArray []CodelessUiConnectorConfigPropertiesSampleQueriesInput

func (CodelessUiConnectorConfigPropertiesSampleQueriesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesSampleQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesSampleQueriesArray) ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput() CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesSampleQueriesArray) ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput)
}

type CodelessUiConnectorConfigPropertiesSampleQueriesOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesSampleQueriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesSampleQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesOutput) ToCodelessUiConnectorConfigPropertiesSampleQueriesOutput() CodelessUiConnectorConfigPropertiesSampleQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesOutput) ToCodelessUiConnectorConfigPropertiesSampleQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesSampleQueriesOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesSampleQueries) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesOutput) Query() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CodelessUiConnectorConfigPropertiesSampleQueries) *string { return v.Query }).(pulumi.StringPtrOutput)
}

type CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput struct{ *pulumi.OutputState }

func (CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesSampleQueries)(nil)).Elem()
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput() CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput) ToCodelessUiConnectorConfigPropertiesSampleQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput {
	return o
}

func (o CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput) Index(i pulumi.IntInput) CodelessUiConnectorConfigPropertiesSampleQueriesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CodelessUiConnectorConfigPropertiesSampleQueries {
		return vs[0].([]CodelessUiConnectorConfigPropertiesSampleQueries)[vs[1].(int)]
	}).(CodelessUiConnectorConfigPropertiesSampleQueriesOutput)
}

type ContentPathMap struct {
	ContentType *string `pulumi:"contentType"`
	Path        *string `pulumi:"path"`
}





type ContentPathMapInput interface {
	pulumi.Input

	ToContentPathMapOutput() ContentPathMapOutput
	ToContentPathMapOutputWithContext(context.Context) ContentPathMapOutput
}

type ContentPathMapArgs struct {
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	Path        pulumi.StringPtrInput `pulumi:"path"`
}

func (ContentPathMapArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentPathMap)(nil)).Elem()
}

func (i ContentPathMapArgs) ToContentPathMapOutput() ContentPathMapOutput {
	return i.ToContentPathMapOutputWithContext(context.Background())
}

func (i ContentPathMapArgs) ToContentPathMapOutputWithContext(ctx context.Context) ContentPathMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentPathMapOutput)
}





type ContentPathMapArrayInput interface {
	pulumi.Input

	ToContentPathMapArrayOutput() ContentPathMapArrayOutput
	ToContentPathMapArrayOutputWithContext(context.Context) ContentPathMapArrayOutput
}

type ContentPathMapArray []ContentPathMapInput

func (ContentPathMapArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentPathMap)(nil)).Elem()
}

func (i ContentPathMapArray) ToContentPathMapArrayOutput() ContentPathMapArrayOutput {
	return i.ToContentPathMapArrayOutputWithContext(context.Background())
}

func (i ContentPathMapArray) ToContentPathMapArrayOutputWithContext(ctx context.Context) ContentPathMapArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentPathMapArrayOutput)
}

type ContentPathMapOutput struct{ *pulumi.OutputState }

func (ContentPathMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentPathMap)(nil)).Elem()
}

func (o ContentPathMapOutput) ToContentPathMapOutput() ContentPathMapOutput {
	return o
}

func (o ContentPathMapOutput) ToContentPathMapOutputWithContext(ctx context.Context) ContentPathMapOutput {
	return o
}

func (o ContentPathMapOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentPathMap) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o ContentPathMapOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentPathMap) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ContentPathMapArrayOutput struct{ *pulumi.OutputState }

func (ContentPathMapArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentPathMap)(nil)).Elem()
}

func (o ContentPathMapArrayOutput) ToContentPathMapArrayOutput() ContentPathMapArrayOutput {
	return o
}

func (o ContentPathMapArrayOutput) ToContentPathMapArrayOutputWithContext(ctx context.Context) ContentPathMapArrayOutput {
	return o
}

func (o ContentPathMapArrayOutput) Index(i pulumi.IntInput) ContentPathMapOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentPathMap {
		return vs[0].([]ContentPathMap)[vs[1].(int)]
	}).(ContentPathMapOutput)
}

type ContentPathMapResponse struct {
	ContentType *string `pulumi:"contentType"`
	Path        *string `pulumi:"path"`
}

type ContentPathMapResponseOutput struct{ *pulumi.OutputState }

func (ContentPathMapResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentPathMapResponse)(nil)).Elem()
}

func (o ContentPathMapResponseOutput) ToContentPathMapResponseOutput() ContentPathMapResponseOutput {
	return o
}

func (o ContentPathMapResponseOutput) ToContentPathMapResponseOutputWithContext(ctx context.Context) ContentPathMapResponseOutput {
	return o
}

func (o ContentPathMapResponseOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentPathMapResponse) *string { return v.ContentType }).(pulumi.StringPtrOutput)
}

func (o ContentPathMapResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentPathMapResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ContentPathMapResponseArrayOutput struct{ *pulumi.OutputState }

func (ContentPathMapResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentPathMapResponse)(nil)).Elem()
}

func (o ContentPathMapResponseArrayOutput) ToContentPathMapResponseArrayOutput() ContentPathMapResponseArrayOutput {
	return o
}

func (o ContentPathMapResponseArrayOutput) ToContentPathMapResponseArrayOutputWithContext(ctx context.Context) ContentPathMapResponseArrayOutput {
	return o
}

func (o ContentPathMapResponseArrayOutput) Index(i pulumi.IntInput) ContentPathMapResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContentPathMapResponse {
		return vs[0].([]ContentPathMapResponse)[vs[1].(int)]
	}).(ContentPathMapResponseOutput)
}

type DataConnectorDataTypeCommon struct {
	State string `pulumi:"state"`
}





type DataConnectorDataTypeCommonInput interface {
	pulumi.Input

	ToDataConnectorDataTypeCommonOutput() DataConnectorDataTypeCommonOutput
	ToDataConnectorDataTypeCommonOutputWithContext(context.Context) DataConnectorDataTypeCommonOutput
}

type DataConnectorDataTypeCommonArgs struct {
	State pulumi.StringInput `pulumi:"state"`
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

func (o DataConnectorDataTypeCommonOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v DataConnectorDataTypeCommon) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type DataConnectorDataTypeCommonResponse struct {
	State string `pulumi:"state"`
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

func (o DataConnectorDataTypeCommonResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v DataConnectorDataTypeCommonResponse) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type Deployment struct {
	DeploymentId      *string `pulumi:"deploymentId"`
	DeploymentLogsUrl *string `pulumi:"deploymentLogsUrl"`
	DeploymentResult  *string `pulumi:"deploymentResult"`
	DeploymentState   *string `pulumi:"deploymentState"`
	DeploymentTime    *string `pulumi:"deploymentTime"`
}





type DeploymentInput interface {
	pulumi.Input

	ToDeploymentOutput() DeploymentOutput
	ToDeploymentOutputWithContext(context.Context) DeploymentOutput
}

type DeploymentArgs struct {
	DeploymentId      pulumi.StringPtrInput `pulumi:"deploymentId"`
	DeploymentLogsUrl pulumi.StringPtrInput `pulumi:"deploymentLogsUrl"`
	DeploymentResult  pulumi.StringPtrInput `pulumi:"deploymentResult"`
	DeploymentState   pulumi.StringPtrInput `pulumi:"deploymentState"`
	DeploymentTime    pulumi.StringPtrInput `pulumi:"deploymentTime"`
}

func (DeploymentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil)).Elem()
}

func (i DeploymentArgs) ToDeploymentOutput() DeploymentOutput {
	return i.ToDeploymentOutputWithContext(context.Background())
}

func (i DeploymentArgs) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput)
}

func (i DeploymentArgs) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i DeploymentArgs) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentOutput).ToDeploymentPtrOutputWithContext(ctx)
}









type DeploymentPtrInput interface {
	pulumi.Input

	ToDeploymentPtrOutput() DeploymentPtrOutput
	ToDeploymentPtrOutputWithContext(context.Context) DeploymentPtrOutput
}

type deploymentPtrType DeploymentArgs

func DeploymentPtr(v *DeploymentArgs) DeploymentPtrInput {
	return (*deploymentPtrType)(v)
}

func (*deploymentPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (i *deploymentPtrType) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return i.ToDeploymentPtrOutputWithContext(context.Background())
}

func (i *deploymentPtrType) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentPtrOutput)
}

type DeploymentOutput struct{ *pulumi.OutputState }

func (DeploymentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Deployment)(nil)).Elem()
}

func (o DeploymentOutput) ToDeploymentOutput() DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentOutputWithContext(ctx context.Context) DeploymentOutput {
	return o
}

func (o DeploymentOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o.ToDeploymentPtrOutputWithContext(context.Background())
}

func (o DeploymentOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Deployment) *Deployment {
		return &v
	}).(DeploymentPtrOutput)
}

func (o DeploymentOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deployment) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

func (o DeploymentOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deployment) *string { return v.DeploymentLogsUrl }).(pulumi.StringPtrOutput)
}

func (o DeploymentOutput) DeploymentResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deployment) *string { return v.DeploymentResult }).(pulumi.StringPtrOutput)
}

func (o DeploymentOutput) DeploymentState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deployment) *string { return v.DeploymentState }).(pulumi.StringPtrOutput)
}

func (o DeploymentOutput) DeploymentTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Deployment) *string { return v.DeploymentTime }).(pulumi.StringPtrOutput)
}

type DeploymentPtrOutput struct{ *pulumi.OutputState }

func (DeploymentPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Deployment)(nil)).Elem()
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutput() DeploymentPtrOutput {
	return o
}

func (o DeploymentPtrOutput) ToDeploymentPtrOutputWithContext(ctx context.Context) DeploymentPtrOutput {
	return o
}

func (o DeploymentPtrOutput) Elem() DeploymentOutput {
	return o.ApplyT(func(v *Deployment) Deployment {
		if v != nil {
			return *v
		}
		var ret Deployment
		return ret
	}).(DeploymentOutput)
}

func (o DeploymentPtrOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentId
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPtrOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentLogsUrl
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPtrOutput) DeploymentResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentResult
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPtrOutput) DeploymentState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentState
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentPtrOutput) DeploymentTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Deployment) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentTime
	}).(pulumi.StringPtrOutput)
}

type DeploymentInfo struct {
	Deployment            *Deployment `pulumi:"deployment"`
	DeploymentFetchStatus *string     `pulumi:"deploymentFetchStatus"`
	Message               *string     `pulumi:"message"`
}





type DeploymentInfoInput interface {
	pulumi.Input

	ToDeploymentInfoOutput() DeploymentInfoOutput
	ToDeploymentInfoOutputWithContext(context.Context) DeploymentInfoOutput
}

type DeploymentInfoArgs struct {
	Deployment            DeploymentPtrInput    `pulumi:"deployment"`
	DeploymentFetchStatus pulumi.StringPtrInput `pulumi:"deploymentFetchStatus"`
	Message               pulumi.StringPtrInput `pulumi:"message"`
}

func (DeploymentInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentInfo)(nil)).Elem()
}

func (i DeploymentInfoArgs) ToDeploymentInfoOutput() DeploymentInfoOutput {
	return i.ToDeploymentInfoOutputWithContext(context.Background())
}

func (i DeploymentInfoArgs) ToDeploymentInfoOutputWithContext(ctx context.Context) DeploymentInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentInfoOutput)
}

func (i DeploymentInfoArgs) ToDeploymentInfoPtrOutput() DeploymentInfoPtrOutput {
	return i.ToDeploymentInfoPtrOutputWithContext(context.Background())
}

func (i DeploymentInfoArgs) ToDeploymentInfoPtrOutputWithContext(ctx context.Context) DeploymentInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentInfoOutput).ToDeploymentInfoPtrOutputWithContext(ctx)
}









type DeploymentInfoPtrInput interface {
	pulumi.Input

	ToDeploymentInfoPtrOutput() DeploymentInfoPtrOutput
	ToDeploymentInfoPtrOutputWithContext(context.Context) DeploymentInfoPtrOutput
}

type deploymentInfoPtrType DeploymentInfoArgs

func DeploymentInfoPtr(v *DeploymentInfoArgs) DeploymentInfoPtrInput {
	return (*deploymentInfoPtrType)(v)
}

func (*deploymentInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentInfo)(nil)).Elem()
}

func (i *deploymentInfoPtrType) ToDeploymentInfoPtrOutput() DeploymentInfoPtrOutput {
	return i.ToDeploymentInfoPtrOutputWithContext(context.Background())
}

func (i *deploymentInfoPtrType) ToDeploymentInfoPtrOutputWithContext(ctx context.Context) DeploymentInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DeploymentInfoPtrOutput)
}

type DeploymentInfoOutput struct{ *pulumi.OutputState }

func (DeploymentInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentInfo)(nil)).Elem()
}

func (o DeploymentInfoOutput) ToDeploymentInfoOutput() DeploymentInfoOutput {
	return o
}

func (o DeploymentInfoOutput) ToDeploymentInfoOutputWithContext(ctx context.Context) DeploymentInfoOutput {
	return o
}

func (o DeploymentInfoOutput) ToDeploymentInfoPtrOutput() DeploymentInfoPtrOutput {
	return o.ToDeploymentInfoPtrOutputWithContext(context.Background())
}

func (o DeploymentInfoOutput) ToDeploymentInfoPtrOutputWithContext(ctx context.Context) DeploymentInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeploymentInfo) *DeploymentInfo {
		return &v
	}).(DeploymentInfoPtrOutput)
}

func (o DeploymentInfoOutput) Deployment() DeploymentPtrOutput {
	return o.ApplyT(func(v DeploymentInfo) *Deployment { return v.Deployment }).(DeploymentPtrOutput)
}

func (o DeploymentInfoOutput) DeploymentFetchStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentInfo) *string { return v.DeploymentFetchStatus }).(pulumi.StringPtrOutput)
}

func (o DeploymentInfoOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentInfo) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type DeploymentInfoPtrOutput struct{ *pulumi.OutputState }

func (DeploymentInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentInfo)(nil)).Elem()
}

func (o DeploymentInfoPtrOutput) ToDeploymentInfoPtrOutput() DeploymentInfoPtrOutput {
	return o
}

func (o DeploymentInfoPtrOutput) ToDeploymentInfoPtrOutputWithContext(ctx context.Context) DeploymentInfoPtrOutput {
	return o
}

func (o DeploymentInfoPtrOutput) Elem() DeploymentInfoOutput {
	return o.ApplyT(func(v *DeploymentInfo) DeploymentInfo {
		if v != nil {
			return *v
		}
		var ret DeploymentInfo
		return ret
	}).(DeploymentInfoOutput)
}

func (o DeploymentInfoPtrOutput) Deployment() DeploymentPtrOutput {
	return o.ApplyT(func(v *DeploymentInfo) *Deployment {
		if v == nil {
			return nil
		}
		return v.Deployment
	}).(DeploymentPtrOutput)
}

func (o DeploymentInfoPtrOutput) DeploymentFetchStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentInfo) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentFetchStatus
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentInfoPtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentInfo) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type DeploymentInfoResponse struct {
	Deployment            *DeploymentResponse `pulumi:"deployment"`
	DeploymentFetchStatus *string             `pulumi:"deploymentFetchStatus"`
	Message               *string             `pulumi:"message"`
}

type DeploymentInfoResponseOutput struct{ *pulumi.OutputState }

func (DeploymentInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentInfoResponse)(nil)).Elem()
}

func (o DeploymentInfoResponseOutput) ToDeploymentInfoResponseOutput() DeploymentInfoResponseOutput {
	return o
}

func (o DeploymentInfoResponseOutput) ToDeploymentInfoResponseOutputWithContext(ctx context.Context) DeploymentInfoResponseOutput {
	return o
}

func (o DeploymentInfoResponseOutput) Deployment() DeploymentResponsePtrOutput {
	return o.ApplyT(func(v DeploymentInfoResponse) *DeploymentResponse { return v.Deployment }).(DeploymentResponsePtrOutput)
}

func (o DeploymentInfoResponseOutput) DeploymentFetchStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentInfoResponse) *string { return v.DeploymentFetchStatus }).(pulumi.StringPtrOutput)
}

func (o DeploymentInfoResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentInfoResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type DeploymentInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentInfoResponse)(nil)).Elem()
}

func (o DeploymentInfoResponsePtrOutput) ToDeploymentInfoResponsePtrOutput() DeploymentInfoResponsePtrOutput {
	return o
}

func (o DeploymentInfoResponsePtrOutput) ToDeploymentInfoResponsePtrOutputWithContext(ctx context.Context) DeploymentInfoResponsePtrOutput {
	return o
}

func (o DeploymentInfoResponsePtrOutput) Elem() DeploymentInfoResponseOutput {
	return o.ApplyT(func(v *DeploymentInfoResponse) DeploymentInfoResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentInfoResponse
		return ret
	}).(DeploymentInfoResponseOutput)
}

func (o DeploymentInfoResponsePtrOutput) Deployment() DeploymentResponsePtrOutput {
	return o.ApplyT(func(v *DeploymentInfoResponse) *DeploymentResponse {
		if v == nil {
			return nil
		}
		return v.Deployment
	}).(DeploymentResponsePtrOutput)
}

func (o DeploymentInfoResponsePtrOutput) DeploymentFetchStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentFetchStatus
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type DeploymentResponse struct {
	DeploymentId      *string `pulumi:"deploymentId"`
	DeploymentLogsUrl *string `pulumi:"deploymentLogsUrl"`
	DeploymentResult  *string `pulumi:"deploymentResult"`
	DeploymentState   *string `pulumi:"deploymentState"`
	DeploymentTime    *string `pulumi:"deploymentTime"`
}

type DeploymentResponseOutput struct{ *pulumi.OutputState }

func (DeploymentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeploymentResponse)(nil)).Elem()
}

func (o DeploymentResponseOutput) ToDeploymentResponseOutput() DeploymentResponseOutput {
	return o
}

func (o DeploymentResponseOutput) ToDeploymentResponseOutputWithContext(ctx context.Context) DeploymentResponseOutput {
	return o
}

func (o DeploymentResponseOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentResponse) *string { return v.DeploymentId }).(pulumi.StringPtrOutput)
}

func (o DeploymentResponseOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentResponse) *string { return v.DeploymentLogsUrl }).(pulumi.StringPtrOutput)
}

func (o DeploymentResponseOutput) DeploymentResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentResponse) *string { return v.DeploymentResult }).(pulumi.StringPtrOutput)
}

func (o DeploymentResponseOutput) DeploymentState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentResponse) *string { return v.DeploymentState }).(pulumi.StringPtrOutput)
}

func (o DeploymentResponseOutput) DeploymentTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DeploymentResponse) *string { return v.DeploymentTime }).(pulumi.StringPtrOutput)
}

type DeploymentResponsePtrOutput struct{ *pulumi.OutputState }

func (DeploymentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeploymentResponse)(nil)).Elem()
}

func (o DeploymentResponsePtrOutput) ToDeploymentResponsePtrOutput() DeploymentResponsePtrOutput {
	return o
}

func (o DeploymentResponsePtrOutput) ToDeploymentResponsePtrOutputWithContext(ctx context.Context) DeploymentResponsePtrOutput {
	return o
}

func (o DeploymentResponsePtrOutput) Elem() DeploymentResponseOutput {
	return o.ApplyT(func(v *DeploymentResponse) DeploymentResponse {
		if v != nil {
			return *v
		}
		var ret DeploymentResponse
		return ret
	}).(DeploymentResponseOutput)
}

func (o DeploymentResponsePtrOutput) DeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentId
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResponsePtrOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentLogsUrl
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResponsePtrOutput) DeploymentResult() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentResult
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResponsePtrOutput) DeploymentState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentState
	}).(pulumi.StringPtrOutput)
}

func (o DeploymentResponsePtrOutput) DeploymentTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DeploymentResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentTime
	}).(pulumi.StringPtrOutput)
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

type EntityFieldMapping struct {
	Identifier *string `pulumi:"identifier"`
	Value      *string `pulumi:"value"`
}





type EntityFieldMappingInput interface {
	pulumi.Input

	ToEntityFieldMappingOutput() EntityFieldMappingOutput
	ToEntityFieldMappingOutputWithContext(context.Context) EntityFieldMappingOutput
}

type EntityFieldMappingArgs struct {
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
	Value      pulumi.StringPtrInput `pulumi:"value"`
}

func (EntityFieldMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityFieldMapping)(nil)).Elem()
}

func (i EntityFieldMappingArgs) ToEntityFieldMappingOutput() EntityFieldMappingOutput {
	return i.ToEntityFieldMappingOutputWithContext(context.Background())
}

func (i EntityFieldMappingArgs) ToEntityFieldMappingOutputWithContext(ctx context.Context) EntityFieldMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityFieldMappingOutput)
}





type EntityFieldMappingArrayInput interface {
	pulumi.Input

	ToEntityFieldMappingArrayOutput() EntityFieldMappingArrayOutput
	ToEntityFieldMappingArrayOutputWithContext(context.Context) EntityFieldMappingArrayOutput
}

type EntityFieldMappingArray []EntityFieldMappingInput

func (EntityFieldMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityFieldMapping)(nil)).Elem()
}

func (i EntityFieldMappingArray) ToEntityFieldMappingArrayOutput() EntityFieldMappingArrayOutput {
	return i.ToEntityFieldMappingArrayOutputWithContext(context.Background())
}

func (i EntityFieldMappingArray) ToEntityFieldMappingArrayOutputWithContext(ctx context.Context) EntityFieldMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityFieldMappingArrayOutput)
}

type EntityFieldMappingOutput struct{ *pulumi.OutputState }

func (EntityFieldMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityFieldMapping)(nil)).Elem()
}

func (o EntityFieldMappingOutput) ToEntityFieldMappingOutput() EntityFieldMappingOutput {
	return o
}

func (o EntityFieldMappingOutput) ToEntityFieldMappingOutputWithContext(ctx context.Context) EntityFieldMappingOutput {
	return o
}

func (o EntityFieldMappingOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityFieldMapping) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

func (o EntityFieldMappingOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityFieldMapping) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EntityFieldMappingArrayOutput struct{ *pulumi.OutputState }

func (EntityFieldMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityFieldMapping)(nil)).Elem()
}

func (o EntityFieldMappingArrayOutput) ToEntityFieldMappingArrayOutput() EntityFieldMappingArrayOutput {
	return o
}

func (o EntityFieldMappingArrayOutput) ToEntityFieldMappingArrayOutputWithContext(ctx context.Context) EntityFieldMappingArrayOutput {
	return o
}

func (o EntityFieldMappingArrayOutput) Index(i pulumi.IntInput) EntityFieldMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityFieldMapping {
		return vs[0].([]EntityFieldMapping)[vs[1].(int)]
	}).(EntityFieldMappingOutput)
}

type EntityFieldMappingResponse struct {
	Identifier *string `pulumi:"identifier"`
	Value      *string `pulumi:"value"`
}

type EntityFieldMappingResponseOutput struct{ *pulumi.OutputState }

func (EntityFieldMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityFieldMappingResponse)(nil)).Elem()
}

func (o EntityFieldMappingResponseOutput) ToEntityFieldMappingResponseOutput() EntityFieldMappingResponseOutput {
	return o
}

func (o EntityFieldMappingResponseOutput) ToEntityFieldMappingResponseOutputWithContext(ctx context.Context) EntityFieldMappingResponseOutput {
	return o
}

func (o EntityFieldMappingResponseOutput) Identifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityFieldMappingResponse) *string { return v.Identifier }).(pulumi.StringPtrOutput)
}

func (o EntityFieldMappingResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityFieldMappingResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EntityFieldMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (EntityFieldMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityFieldMappingResponse)(nil)).Elem()
}

func (o EntityFieldMappingResponseArrayOutput) ToEntityFieldMappingResponseArrayOutput() EntityFieldMappingResponseArrayOutput {
	return o
}

func (o EntityFieldMappingResponseArrayOutput) ToEntityFieldMappingResponseArrayOutputWithContext(ctx context.Context) EntityFieldMappingResponseArrayOutput {
	return o
}

func (o EntityFieldMappingResponseArrayOutput) Index(i pulumi.IntInput) EntityFieldMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityFieldMappingResponse {
		return vs[0].([]EntityFieldMappingResponse)[vs[1].(int)]
	}).(EntityFieldMappingResponseOutput)
}

type EntityInsightItemResponse struct {
	ChartQueryResults []InsightsTableResultResponse               `pulumi:"chartQueryResults"`
	QueryId           *string                                     `pulumi:"queryId"`
	QueryTimeInterval *EntityInsightItemResponseQueryTimeInterval `pulumi:"queryTimeInterval"`
	TableQueryResults *InsightsTableResultResponse                `pulumi:"tableQueryResults"`
}

type EntityInsightItemResponseOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInsightItemResponse)(nil)).Elem()
}

func (o EntityInsightItemResponseOutput) ToEntityInsightItemResponseOutput() EntityInsightItemResponseOutput {
	return o
}

func (o EntityInsightItemResponseOutput) ToEntityInsightItemResponseOutputWithContext(ctx context.Context) EntityInsightItemResponseOutput {
	return o
}

func (o EntityInsightItemResponseOutput) ChartQueryResults() InsightsTableResultResponseArrayOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) []InsightsTableResultResponse { return v.ChartQueryResults }).(InsightsTableResultResponseArrayOutput)
}

func (o EntityInsightItemResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

func (o EntityInsightItemResponseOutput) QueryTimeInterval() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) *EntityInsightItemResponseQueryTimeInterval {
		return v.QueryTimeInterval
	}).(EntityInsightItemResponseQueryTimeIntervalPtrOutput)
}

func (o EntityInsightItemResponseOutput) TableQueryResults() InsightsTableResultResponsePtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponse) *InsightsTableResultResponse { return v.TableQueryResults }).(InsightsTableResultResponsePtrOutput)
}

type EntityInsightItemResponseArrayOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityInsightItemResponse)(nil)).Elem()
}

func (o EntityInsightItemResponseArrayOutput) ToEntityInsightItemResponseArrayOutput() EntityInsightItemResponseArrayOutput {
	return o
}

func (o EntityInsightItemResponseArrayOutput) ToEntityInsightItemResponseArrayOutputWithContext(ctx context.Context) EntityInsightItemResponseArrayOutput {
	return o
}

func (o EntityInsightItemResponseArrayOutput) Index(i pulumi.IntInput) EntityInsightItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EntityInsightItemResponse {
		return vs[0].([]EntityInsightItemResponse)[vs[1].(int)]
	}).(EntityInsightItemResponseOutput)
}

type EntityInsightItemResponseQueryTimeInterval struct {
	EndTime   *string `pulumi:"endTime"`
	StartTime *string `pulumi:"startTime"`
}

type EntityInsightItemResponseQueryTimeIntervalOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseQueryTimeIntervalOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInsightItemResponseQueryTimeInterval)(nil)).Elem()
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToEntityInsightItemResponseQueryTimeIntervalOutput() EntityInsightItemResponseQueryTimeIntervalOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToEntityInsightItemResponseQueryTimeIntervalOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponseQueryTimeInterval) *string { return v.EndTime }).(pulumi.StringPtrOutput)
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EntityInsightItemResponseQueryTimeInterval) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type EntityInsightItemResponseQueryTimeIntervalPtrOutput struct{ *pulumi.OutputState }

func (EntityInsightItemResponseQueryTimeIntervalPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityInsightItemResponseQueryTimeInterval)(nil)).Elem()
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) ToEntityInsightItemResponseQueryTimeIntervalPtrOutput() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) Elem() EntityInsightItemResponseQueryTimeIntervalOutput {
	return o.ApplyT(func(v *EntityInsightItemResponseQueryTimeInterval) EntityInsightItemResponseQueryTimeInterval {
		if v != nil {
			return *v
		}
		var ret EntityInsightItemResponseQueryTimeInterval
		return ret
	}).(EntityInsightItemResponseQueryTimeIntervalOutput)
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) EndTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityInsightItemResponseQueryTimeInterval) *string {
		if v == nil {
			return nil
		}
		return v.EndTime
	}).(pulumi.StringPtrOutput)
}

func (o EntityInsightItemResponseQueryTimeIntervalPtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EntityInsightItemResponseQueryTimeInterval) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
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

type FusionScenarioExclusionPattern struct {
	DateAddedInUTC   string `pulumi:"dateAddedInUTC"`
	ExclusionPattern string `pulumi:"exclusionPattern"`
}





type FusionScenarioExclusionPatternInput interface {
	pulumi.Input

	ToFusionScenarioExclusionPatternOutput() FusionScenarioExclusionPatternOutput
	ToFusionScenarioExclusionPatternOutputWithContext(context.Context) FusionScenarioExclusionPatternOutput
}

type FusionScenarioExclusionPatternArgs struct {
	DateAddedInUTC   pulumi.StringInput `pulumi:"dateAddedInUTC"`
	ExclusionPattern pulumi.StringInput `pulumi:"exclusionPattern"`
}

func (FusionScenarioExclusionPatternArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionScenarioExclusionPattern)(nil)).Elem()
}

func (i FusionScenarioExclusionPatternArgs) ToFusionScenarioExclusionPatternOutput() FusionScenarioExclusionPatternOutput {
	return i.ToFusionScenarioExclusionPatternOutputWithContext(context.Background())
}

func (i FusionScenarioExclusionPatternArgs) ToFusionScenarioExclusionPatternOutputWithContext(ctx context.Context) FusionScenarioExclusionPatternOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionScenarioExclusionPatternOutput)
}





type FusionScenarioExclusionPatternArrayInput interface {
	pulumi.Input

	ToFusionScenarioExclusionPatternArrayOutput() FusionScenarioExclusionPatternArrayOutput
	ToFusionScenarioExclusionPatternArrayOutputWithContext(context.Context) FusionScenarioExclusionPatternArrayOutput
}

type FusionScenarioExclusionPatternArray []FusionScenarioExclusionPatternInput

func (FusionScenarioExclusionPatternArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionScenarioExclusionPattern)(nil)).Elem()
}

func (i FusionScenarioExclusionPatternArray) ToFusionScenarioExclusionPatternArrayOutput() FusionScenarioExclusionPatternArrayOutput {
	return i.ToFusionScenarioExclusionPatternArrayOutputWithContext(context.Background())
}

func (i FusionScenarioExclusionPatternArray) ToFusionScenarioExclusionPatternArrayOutputWithContext(ctx context.Context) FusionScenarioExclusionPatternArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionScenarioExclusionPatternArrayOutput)
}

type FusionScenarioExclusionPatternOutput struct{ *pulumi.OutputState }

func (FusionScenarioExclusionPatternOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionScenarioExclusionPattern)(nil)).Elem()
}

func (o FusionScenarioExclusionPatternOutput) ToFusionScenarioExclusionPatternOutput() FusionScenarioExclusionPatternOutput {
	return o
}

func (o FusionScenarioExclusionPatternOutput) ToFusionScenarioExclusionPatternOutputWithContext(ctx context.Context) FusionScenarioExclusionPatternOutput {
	return o
}

func (o FusionScenarioExclusionPatternOutput) DateAddedInUTC() pulumi.StringOutput {
	return o.ApplyT(func(v FusionScenarioExclusionPattern) string { return v.DateAddedInUTC }).(pulumi.StringOutput)
}

func (o FusionScenarioExclusionPatternOutput) ExclusionPattern() pulumi.StringOutput {
	return o.ApplyT(func(v FusionScenarioExclusionPattern) string { return v.ExclusionPattern }).(pulumi.StringOutput)
}

type FusionScenarioExclusionPatternArrayOutput struct{ *pulumi.OutputState }

func (FusionScenarioExclusionPatternArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionScenarioExclusionPattern)(nil)).Elem()
}

func (o FusionScenarioExclusionPatternArrayOutput) ToFusionScenarioExclusionPatternArrayOutput() FusionScenarioExclusionPatternArrayOutput {
	return o
}

func (o FusionScenarioExclusionPatternArrayOutput) ToFusionScenarioExclusionPatternArrayOutputWithContext(ctx context.Context) FusionScenarioExclusionPatternArrayOutput {
	return o
}

func (o FusionScenarioExclusionPatternArrayOutput) Index(i pulumi.IntInput) FusionScenarioExclusionPatternOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionScenarioExclusionPattern {
		return vs[0].([]FusionScenarioExclusionPattern)[vs[1].(int)]
	}).(FusionScenarioExclusionPatternOutput)
}

type FusionScenarioExclusionPatternResponse struct {
	DateAddedInUTC   string `pulumi:"dateAddedInUTC"`
	ExclusionPattern string `pulumi:"exclusionPattern"`
}

type FusionScenarioExclusionPatternResponseOutput struct{ *pulumi.OutputState }

func (FusionScenarioExclusionPatternResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionScenarioExclusionPatternResponse)(nil)).Elem()
}

func (o FusionScenarioExclusionPatternResponseOutput) ToFusionScenarioExclusionPatternResponseOutput() FusionScenarioExclusionPatternResponseOutput {
	return o
}

func (o FusionScenarioExclusionPatternResponseOutput) ToFusionScenarioExclusionPatternResponseOutputWithContext(ctx context.Context) FusionScenarioExclusionPatternResponseOutput {
	return o
}

func (o FusionScenarioExclusionPatternResponseOutput) DateAddedInUTC() pulumi.StringOutput {
	return o.ApplyT(func(v FusionScenarioExclusionPatternResponse) string { return v.DateAddedInUTC }).(pulumi.StringOutput)
}

func (o FusionScenarioExclusionPatternResponseOutput) ExclusionPattern() pulumi.StringOutput {
	return o.ApplyT(func(v FusionScenarioExclusionPatternResponse) string { return v.ExclusionPattern }).(pulumi.StringOutput)
}

type FusionScenarioExclusionPatternResponseArrayOutput struct{ *pulumi.OutputState }

func (FusionScenarioExclusionPatternResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionScenarioExclusionPatternResponse)(nil)).Elem()
}

func (o FusionScenarioExclusionPatternResponseArrayOutput) ToFusionScenarioExclusionPatternResponseArrayOutput() FusionScenarioExclusionPatternResponseArrayOutput {
	return o
}

func (o FusionScenarioExclusionPatternResponseArrayOutput) ToFusionScenarioExclusionPatternResponseArrayOutputWithContext(ctx context.Context) FusionScenarioExclusionPatternResponseArrayOutput {
	return o
}

func (o FusionScenarioExclusionPatternResponseArrayOutput) Index(i pulumi.IntInput) FusionScenarioExclusionPatternResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionScenarioExclusionPatternResponse {
		return vs[0].([]FusionScenarioExclusionPatternResponse)[vs[1].(int)]
	}).(FusionScenarioExclusionPatternResponseOutput)
}

type FusionSourceSettings struct {
	Enabled        bool                         `pulumi:"enabled"`
	SourceName     string                       `pulumi:"sourceName"`
	SourceSubTypes []FusionSourceSubTypeSetting `pulumi:"sourceSubTypes"`
}





type FusionSourceSettingsInput interface {
	pulumi.Input

	ToFusionSourceSettingsOutput() FusionSourceSettingsOutput
	ToFusionSourceSettingsOutputWithContext(context.Context) FusionSourceSettingsOutput
}

type FusionSourceSettingsArgs struct {
	Enabled        pulumi.BoolInput                     `pulumi:"enabled"`
	SourceName     pulumi.StringInput                   `pulumi:"sourceName"`
	SourceSubTypes FusionSourceSubTypeSettingArrayInput `pulumi:"sourceSubTypes"`
}

func (FusionSourceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSourceSettings)(nil)).Elem()
}

func (i FusionSourceSettingsArgs) ToFusionSourceSettingsOutput() FusionSourceSettingsOutput {
	return i.ToFusionSourceSettingsOutputWithContext(context.Background())
}

func (i FusionSourceSettingsArgs) ToFusionSourceSettingsOutputWithContext(ctx context.Context) FusionSourceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSourceSettingsOutput)
}





type FusionSourceSettingsArrayInput interface {
	pulumi.Input

	ToFusionSourceSettingsArrayOutput() FusionSourceSettingsArrayOutput
	ToFusionSourceSettingsArrayOutputWithContext(context.Context) FusionSourceSettingsArrayOutput
}

type FusionSourceSettingsArray []FusionSourceSettingsInput

func (FusionSourceSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSourceSettings)(nil)).Elem()
}

func (i FusionSourceSettingsArray) ToFusionSourceSettingsArrayOutput() FusionSourceSettingsArrayOutput {
	return i.ToFusionSourceSettingsArrayOutputWithContext(context.Background())
}

func (i FusionSourceSettingsArray) ToFusionSourceSettingsArrayOutputWithContext(ctx context.Context) FusionSourceSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSourceSettingsArrayOutput)
}

type FusionSourceSettingsOutput struct{ *pulumi.OutputState }

func (FusionSourceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSourceSettings)(nil)).Elem()
}

func (o FusionSourceSettingsOutput) ToFusionSourceSettingsOutput() FusionSourceSettingsOutput {
	return o
}

func (o FusionSourceSettingsOutput) ToFusionSourceSettingsOutputWithContext(ctx context.Context) FusionSourceSettingsOutput {
	return o
}

func (o FusionSourceSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSourceSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionSourceSettingsOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSourceSettings) string { return v.SourceName }).(pulumi.StringOutput)
}

func (o FusionSourceSettingsOutput) SourceSubTypes() FusionSourceSubTypeSettingArrayOutput {
	return o.ApplyT(func(v FusionSourceSettings) []FusionSourceSubTypeSetting { return v.SourceSubTypes }).(FusionSourceSubTypeSettingArrayOutput)
}

type FusionSourceSettingsArrayOutput struct{ *pulumi.OutputState }

func (FusionSourceSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSourceSettings)(nil)).Elem()
}

func (o FusionSourceSettingsArrayOutput) ToFusionSourceSettingsArrayOutput() FusionSourceSettingsArrayOutput {
	return o
}

func (o FusionSourceSettingsArrayOutput) ToFusionSourceSettingsArrayOutputWithContext(ctx context.Context) FusionSourceSettingsArrayOutput {
	return o
}

func (o FusionSourceSettingsArrayOutput) Index(i pulumi.IntInput) FusionSourceSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionSourceSettings {
		return vs[0].([]FusionSourceSettings)[vs[1].(int)]
	}).(FusionSourceSettingsOutput)
}

type FusionSourceSettingsResponse struct {
	Enabled        bool                                 `pulumi:"enabled"`
	SourceName     string                               `pulumi:"sourceName"`
	SourceSubTypes []FusionSourceSubTypeSettingResponse `pulumi:"sourceSubTypes"`
}

type FusionSourceSettingsResponseOutput struct{ *pulumi.OutputState }

func (FusionSourceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSourceSettingsResponse)(nil)).Elem()
}

func (o FusionSourceSettingsResponseOutput) ToFusionSourceSettingsResponseOutput() FusionSourceSettingsResponseOutput {
	return o
}

func (o FusionSourceSettingsResponseOutput) ToFusionSourceSettingsResponseOutputWithContext(ctx context.Context) FusionSourceSettingsResponseOutput {
	return o
}

func (o FusionSourceSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSourceSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionSourceSettingsResponseOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSourceSettingsResponse) string { return v.SourceName }).(pulumi.StringOutput)
}

func (o FusionSourceSettingsResponseOutput) SourceSubTypes() FusionSourceSubTypeSettingResponseArrayOutput {
	return o.ApplyT(func(v FusionSourceSettingsResponse) []FusionSourceSubTypeSettingResponse { return v.SourceSubTypes }).(FusionSourceSubTypeSettingResponseArrayOutput)
}

type FusionSourceSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (FusionSourceSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSourceSettingsResponse)(nil)).Elem()
}

func (o FusionSourceSettingsResponseArrayOutput) ToFusionSourceSettingsResponseArrayOutput() FusionSourceSettingsResponseArrayOutput {
	return o
}

func (o FusionSourceSettingsResponseArrayOutput) ToFusionSourceSettingsResponseArrayOutputWithContext(ctx context.Context) FusionSourceSettingsResponseArrayOutput {
	return o
}

func (o FusionSourceSettingsResponseArrayOutput) Index(i pulumi.IntInput) FusionSourceSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionSourceSettingsResponse {
		return vs[0].([]FusionSourceSettingsResponse)[vs[1].(int)]
	}).(FusionSourceSettingsResponseOutput)
}

type FusionSourceSubTypeSetting struct {
	Enabled           bool                        `pulumi:"enabled"`
	SeverityFilters   FusionSubTypeSeverityFilter `pulumi:"severityFilters"`
	SourceSubTypeName string                      `pulumi:"sourceSubTypeName"`
}





type FusionSourceSubTypeSettingInput interface {
	pulumi.Input

	ToFusionSourceSubTypeSettingOutput() FusionSourceSubTypeSettingOutput
	ToFusionSourceSubTypeSettingOutputWithContext(context.Context) FusionSourceSubTypeSettingOutput
}

type FusionSourceSubTypeSettingArgs struct {
	Enabled           pulumi.BoolInput                 `pulumi:"enabled"`
	SeverityFilters   FusionSubTypeSeverityFilterInput `pulumi:"severityFilters"`
	SourceSubTypeName pulumi.StringInput               `pulumi:"sourceSubTypeName"`
}

func (FusionSourceSubTypeSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSourceSubTypeSetting)(nil)).Elem()
}

func (i FusionSourceSubTypeSettingArgs) ToFusionSourceSubTypeSettingOutput() FusionSourceSubTypeSettingOutput {
	return i.ToFusionSourceSubTypeSettingOutputWithContext(context.Background())
}

func (i FusionSourceSubTypeSettingArgs) ToFusionSourceSubTypeSettingOutputWithContext(ctx context.Context) FusionSourceSubTypeSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSourceSubTypeSettingOutput)
}





type FusionSourceSubTypeSettingArrayInput interface {
	pulumi.Input

	ToFusionSourceSubTypeSettingArrayOutput() FusionSourceSubTypeSettingArrayOutput
	ToFusionSourceSubTypeSettingArrayOutputWithContext(context.Context) FusionSourceSubTypeSettingArrayOutput
}

type FusionSourceSubTypeSettingArray []FusionSourceSubTypeSettingInput

func (FusionSourceSubTypeSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSourceSubTypeSetting)(nil)).Elem()
}

func (i FusionSourceSubTypeSettingArray) ToFusionSourceSubTypeSettingArrayOutput() FusionSourceSubTypeSettingArrayOutput {
	return i.ToFusionSourceSubTypeSettingArrayOutputWithContext(context.Background())
}

func (i FusionSourceSubTypeSettingArray) ToFusionSourceSubTypeSettingArrayOutputWithContext(ctx context.Context) FusionSourceSubTypeSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSourceSubTypeSettingArrayOutput)
}

type FusionSourceSubTypeSettingOutput struct{ *pulumi.OutputState }

func (FusionSourceSubTypeSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSourceSubTypeSetting)(nil)).Elem()
}

func (o FusionSourceSubTypeSettingOutput) ToFusionSourceSubTypeSettingOutput() FusionSourceSubTypeSettingOutput {
	return o
}

func (o FusionSourceSubTypeSettingOutput) ToFusionSourceSubTypeSettingOutputWithContext(ctx context.Context) FusionSourceSubTypeSettingOutput {
	return o
}

func (o FusionSourceSubTypeSettingOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSetting) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionSourceSubTypeSettingOutput) SeverityFilters() FusionSubTypeSeverityFilterOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSetting) FusionSubTypeSeverityFilter { return v.SeverityFilters }).(FusionSubTypeSeverityFilterOutput)
}

func (o FusionSourceSubTypeSettingOutput) SourceSubTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSetting) string { return v.SourceSubTypeName }).(pulumi.StringOutput)
}

type FusionSourceSubTypeSettingArrayOutput struct{ *pulumi.OutputState }

func (FusionSourceSubTypeSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSourceSubTypeSetting)(nil)).Elem()
}

func (o FusionSourceSubTypeSettingArrayOutput) ToFusionSourceSubTypeSettingArrayOutput() FusionSourceSubTypeSettingArrayOutput {
	return o
}

func (o FusionSourceSubTypeSettingArrayOutput) ToFusionSourceSubTypeSettingArrayOutputWithContext(ctx context.Context) FusionSourceSubTypeSettingArrayOutput {
	return o
}

func (o FusionSourceSubTypeSettingArrayOutput) Index(i pulumi.IntInput) FusionSourceSubTypeSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionSourceSubTypeSetting {
		return vs[0].([]FusionSourceSubTypeSetting)[vs[1].(int)]
	}).(FusionSourceSubTypeSettingOutput)
}

type FusionSourceSubTypeSettingResponse struct {
	Enabled                  bool                                `pulumi:"enabled"`
	SeverityFilters          FusionSubTypeSeverityFilterResponse `pulumi:"severityFilters"`
	SourceSubTypeDisplayName string                              `pulumi:"sourceSubTypeDisplayName"`
	SourceSubTypeName        string                              `pulumi:"sourceSubTypeName"`
}

type FusionSourceSubTypeSettingResponseOutput struct{ *pulumi.OutputState }

func (FusionSourceSubTypeSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSourceSubTypeSettingResponse)(nil)).Elem()
}

func (o FusionSourceSubTypeSettingResponseOutput) ToFusionSourceSubTypeSettingResponseOutput() FusionSourceSubTypeSettingResponseOutput {
	return o
}

func (o FusionSourceSubTypeSettingResponseOutput) ToFusionSourceSubTypeSettingResponseOutputWithContext(ctx context.Context) FusionSourceSubTypeSettingResponseOutput {
	return o
}

func (o FusionSourceSubTypeSettingResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSettingResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionSourceSubTypeSettingResponseOutput) SeverityFilters() FusionSubTypeSeverityFilterResponseOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSettingResponse) FusionSubTypeSeverityFilterResponse {
		return v.SeverityFilters
	}).(FusionSubTypeSeverityFilterResponseOutput)
}

func (o FusionSourceSubTypeSettingResponseOutput) SourceSubTypeDisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSettingResponse) string { return v.SourceSubTypeDisplayName }).(pulumi.StringOutput)
}

func (o FusionSourceSubTypeSettingResponseOutput) SourceSubTypeName() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSourceSubTypeSettingResponse) string { return v.SourceSubTypeName }).(pulumi.StringOutput)
}

type FusionSourceSubTypeSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (FusionSourceSubTypeSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSourceSubTypeSettingResponse)(nil)).Elem()
}

func (o FusionSourceSubTypeSettingResponseArrayOutput) ToFusionSourceSubTypeSettingResponseArrayOutput() FusionSourceSubTypeSettingResponseArrayOutput {
	return o
}

func (o FusionSourceSubTypeSettingResponseArrayOutput) ToFusionSourceSubTypeSettingResponseArrayOutputWithContext(ctx context.Context) FusionSourceSubTypeSettingResponseArrayOutput {
	return o
}

func (o FusionSourceSubTypeSettingResponseArrayOutput) Index(i pulumi.IntInput) FusionSourceSubTypeSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionSourceSubTypeSettingResponse {
		return vs[0].([]FusionSourceSubTypeSettingResponse)[vs[1].(int)]
	}).(FusionSourceSubTypeSettingResponseOutput)
}

type FusionSubTypeSeverityFilter struct {
	Filters []FusionSubTypeSeverityFiltersItem `pulumi:"filters"`
}





type FusionSubTypeSeverityFilterInput interface {
	pulumi.Input

	ToFusionSubTypeSeverityFilterOutput() FusionSubTypeSeverityFilterOutput
	ToFusionSubTypeSeverityFilterOutputWithContext(context.Context) FusionSubTypeSeverityFilterOutput
}

type FusionSubTypeSeverityFilterArgs struct {
	Filters FusionSubTypeSeverityFiltersItemArrayInput `pulumi:"filters"`
}

func (FusionSubTypeSeverityFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSubTypeSeverityFilter)(nil)).Elem()
}

func (i FusionSubTypeSeverityFilterArgs) ToFusionSubTypeSeverityFilterOutput() FusionSubTypeSeverityFilterOutput {
	return i.ToFusionSubTypeSeverityFilterOutputWithContext(context.Background())
}

func (i FusionSubTypeSeverityFilterArgs) ToFusionSubTypeSeverityFilterOutputWithContext(ctx context.Context) FusionSubTypeSeverityFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSubTypeSeverityFilterOutput)
}

type FusionSubTypeSeverityFilterOutput struct{ *pulumi.OutputState }

func (FusionSubTypeSeverityFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSubTypeSeverityFilter)(nil)).Elem()
}

func (o FusionSubTypeSeverityFilterOutput) ToFusionSubTypeSeverityFilterOutput() FusionSubTypeSeverityFilterOutput {
	return o
}

func (o FusionSubTypeSeverityFilterOutput) ToFusionSubTypeSeverityFilterOutputWithContext(ctx context.Context) FusionSubTypeSeverityFilterOutput {
	return o
}

func (o FusionSubTypeSeverityFilterOutput) Filters() FusionSubTypeSeverityFiltersItemArrayOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFilter) []FusionSubTypeSeverityFiltersItem { return v.Filters }).(FusionSubTypeSeverityFiltersItemArrayOutput)
}

type FusionSubTypeSeverityFilterResponse struct {
	Filters     []FusionSubTypeSeverityFiltersItemResponse `pulumi:"filters"`
	IsSupported bool                                       `pulumi:"isSupported"`
}

type FusionSubTypeSeverityFilterResponseOutput struct{ *pulumi.OutputState }

func (FusionSubTypeSeverityFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSubTypeSeverityFilterResponse)(nil)).Elem()
}

func (o FusionSubTypeSeverityFilterResponseOutput) ToFusionSubTypeSeverityFilterResponseOutput() FusionSubTypeSeverityFilterResponseOutput {
	return o
}

func (o FusionSubTypeSeverityFilterResponseOutput) ToFusionSubTypeSeverityFilterResponseOutputWithContext(ctx context.Context) FusionSubTypeSeverityFilterResponseOutput {
	return o
}

func (o FusionSubTypeSeverityFilterResponseOutput) Filters() FusionSubTypeSeverityFiltersItemResponseArrayOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFilterResponse) []FusionSubTypeSeverityFiltersItemResponse {
		return v.Filters
	}).(FusionSubTypeSeverityFiltersItemResponseArrayOutput)
}

func (o FusionSubTypeSeverityFilterResponseOutput) IsSupported() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFilterResponse) bool { return v.IsSupported }).(pulumi.BoolOutput)
}

type FusionSubTypeSeverityFiltersItem struct {
	Enabled  bool   `pulumi:"enabled"`
	Severity string `pulumi:"severity"`
}





type FusionSubTypeSeverityFiltersItemInput interface {
	pulumi.Input

	ToFusionSubTypeSeverityFiltersItemOutput() FusionSubTypeSeverityFiltersItemOutput
	ToFusionSubTypeSeverityFiltersItemOutputWithContext(context.Context) FusionSubTypeSeverityFiltersItemOutput
}

type FusionSubTypeSeverityFiltersItemArgs struct {
	Enabled  pulumi.BoolInput   `pulumi:"enabled"`
	Severity pulumi.StringInput `pulumi:"severity"`
}

func (FusionSubTypeSeverityFiltersItemArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSubTypeSeverityFiltersItem)(nil)).Elem()
}

func (i FusionSubTypeSeverityFiltersItemArgs) ToFusionSubTypeSeverityFiltersItemOutput() FusionSubTypeSeverityFiltersItemOutput {
	return i.ToFusionSubTypeSeverityFiltersItemOutputWithContext(context.Background())
}

func (i FusionSubTypeSeverityFiltersItemArgs) ToFusionSubTypeSeverityFiltersItemOutputWithContext(ctx context.Context) FusionSubTypeSeverityFiltersItemOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSubTypeSeverityFiltersItemOutput)
}





type FusionSubTypeSeverityFiltersItemArrayInput interface {
	pulumi.Input

	ToFusionSubTypeSeverityFiltersItemArrayOutput() FusionSubTypeSeverityFiltersItemArrayOutput
	ToFusionSubTypeSeverityFiltersItemArrayOutputWithContext(context.Context) FusionSubTypeSeverityFiltersItemArrayOutput
}

type FusionSubTypeSeverityFiltersItemArray []FusionSubTypeSeverityFiltersItemInput

func (FusionSubTypeSeverityFiltersItemArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSubTypeSeverityFiltersItem)(nil)).Elem()
}

func (i FusionSubTypeSeverityFiltersItemArray) ToFusionSubTypeSeverityFiltersItemArrayOutput() FusionSubTypeSeverityFiltersItemArrayOutput {
	return i.ToFusionSubTypeSeverityFiltersItemArrayOutputWithContext(context.Background())
}

func (i FusionSubTypeSeverityFiltersItemArray) ToFusionSubTypeSeverityFiltersItemArrayOutputWithContext(ctx context.Context) FusionSubTypeSeverityFiltersItemArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionSubTypeSeverityFiltersItemArrayOutput)
}

type FusionSubTypeSeverityFiltersItemOutput struct{ *pulumi.OutputState }

func (FusionSubTypeSeverityFiltersItemOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSubTypeSeverityFiltersItem)(nil)).Elem()
}

func (o FusionSubTypeSeverityFiltersItemOutput) ToFusionSubTypeSeverityFiltersItemOutput() FusionSubTypeSeverityFiltersItemOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemOutput) ToFusionSubTypeSeverityFiltersItemOutputWithContext(ctx context.Context) FusionSubTypeSeverityFiltersItemOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFiltersItem) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionSubTypeSeverityFiltersItemOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFiltersItem) string { return v.Severity }).(pulumi.StringOutput)
}

type FusionSubTypeSeverityFiltersItemArrayOutput struct{ *pulumi.OutputState }

func (FusionSubTypeSeverityFiltersItemArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSubTypeSeverityFiltersItem)(nil)).Elem()
}

func (o FusionSubTypeSeverityFiltersItemArrayOutput) ToFusionSubTypeSeverityFiltersItemArrayOutput() FusionSubTypeSeverityFiltersItemArrayOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemArrayOutput) ToFusionSubTypeSeverityFiltersItemArrayOutputWithContext(ctx context.Context) FusionSubTypeSeverityFiltersItemArrayOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemArrayOutput) Index(i pulumi.IntInput) FusionSubTypeSeverityFiltersItemOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionSubTypeSeverityFiltersItem {
		return vs[0].([]FusionSubTypeSeverityFiltersItem)[vs[1].(int)]
	}).(FusionSubTypeSeverityFiltersItemOutput)
}

type FusionSubTypeSeverityFiltersItemResponse struct {
	Enabled  bool   `pulumi:"enabled"`
	Severity string `pulumi:"severity"`
}

type FusionSubTypeSeverityFiltersItemResponseOutput struct{ *pulumi.OutputState }

func (FusionSubTypeSeverityFiltersItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionSubTypeSeverityFiltersItemResponse)(nil)).Elem()
}

func (o FusionSubTypeSeverityFiltersItemResponseOutput) ToFusionSubTypeSeverityFiltersItemResponseOutput() FusionSubTypeSeverityFiltersItemResponseOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemResponseOutput) ToFusionSubTypeSeverityFiltersItemResponseOutputWithContext(ctx context.Context) FusionSubTypeSeverityFiltersItemResponseOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFiltersItemResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionSubTypeSeverityFiltersItemResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v FusionSubTypeSeverityFiltersItemResponse) string { return v.Severity }).(pulumi.StringOutput)
}

type FusionSubTypeSeverityFiltersItemResponseArrayOutput struct{ *pulumi.OutputState }

func (FusionSubTypeSeverityFiltersItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FusionSubTypeSeverityFiltersItemResponse)(nil)).Elem()
}

func (o FusionSubTypeSeverityFiltersItemResponseArrayOutput) ToFusionSubTypeSeverityFiltersItemResponseArrayOutput() FusionSubTypeSeverityFiltersItemResponseArrayOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemResponseArrayOutput) ToFusionSubTypeSeverityFiltersItemResponseArrayOutputWithContext(ctx context.Context) FusionSubTypeSeverityFiltersItemResponseArrayOutput {
	return o
}

func (o FusionSubTypeSeverityFiltersItemResponseArrayOutput) Index(i pulumi.IntInput) FusionSubTypeSeverityFiltersItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FusionSubTypeSeverityFiltersItemResponse {
		return vs[0].([]FusionSubTypeSeverityFiltersItemResponse)[vs[1].(int)]
	}).(FusionSubTypeSeverityFiltersItemResponseOutput)
}

type GetInsightsErrorKindResponse struct {
	ErrorMessage string  `pulumi:"errorMessage"`
	Kind         string  `pulumi:"kind"`
	QueryId      *string `pulumi:"queryId"`
}

type GetInsightsErrorKindResponseOutput struct{ *pulumi.OutputState }

func (GetInsightsErrorKindResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsErrorKindResponse)(nil)).Elem()
}

func (o GetInsightsErrorKindResponseOutput) ToGetInsightsErrorKindResponseOutput() GetInsightsErrorKindResponseOutput {
	return o
}

func (o GetInsightsErrorKindResponseOutput) ToGetInsightsErrorKindResponseOutputWithContext(ctx context.Context) GetInsightsErrorKindResponseOutput {
	return o
}

func (o GetInsightsErrorKindResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v GetInsightsErrorKindResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o GetInsightsErrorKindResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetInsightsErrorKindResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetInsightsErrorKindResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInsightsErrorKindResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

type GetInsightsErrorKindResponseArrayOutput struct{ *pulumi.OutputState }

func (GetInsightsErrorKindResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInsightsErrorKindResponse)(nil)).Elem()
}

func (o GetInsightsErrorKindResponseArrayOutput) ToGetInsightsErrorKindResponseArrayOutput() GetInsightsErrorKindResponseArrayOutput {
	return o
}

func (o GetInsightsErrorKindResponseArrayOutput) ToGetInsightsErrorKindResponseArrayOutputWithContext(ctx context.Context) GetInsightsErrorKindResponseArrayOutput {
	return o
}

func (o GetInsightsErrorKindResponseArrayOutput) Index(i pulumi.IntInput) GetInsightsErrorKindResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInsightsErrorKindResponse {
		return vs[0].([]GetInsightsErrorKindResponse)[vs[1].(int)]
	}).(GetInsightsErrorKindResponseOutput)
}

type GetInsightsResultsMetadataResponse struct {
	Errors     []GetInsightsErrorKindResponse `pulumi:"errors"`
	TotalCount int                            `pulumi:"totalCount"`
}

type GetInsightsResultsMetadataResponseOutput struct{ *pulumi.OutputState }

func (GetInsightsResultsMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsResultsMetadataResponse)(nil)).Elem()
}

func (o GetInsightsResultsMetadataResponseOutput) ToGetInsightsResultsMetadataResponseOutput() GetInsightsResultsMetadataResponseOutput {
	return o
}

func (o GetInsightsResultsMetadataResponseOutput) ToGetInsightsResultsMetadataResponseOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponseOutput {
	return o
}

func (o GetInsightsResultsMetadataResponseOutput) Errors() GetInsightsErrorKindResponseArrayOutput {
	return o.ApplyT(func(v GetInsightsResultsMetadataResponse) []GetInsightsErrorKindResponse { return v.Errors }).(GetInsightsErrorKindResponseArrayOutput)
}

func (o GetInsightsResultsMetadataResponseOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v GetInsightsResultsMetadataResponse) int { return v.TotalCount }).(pulumi.IntOutput)
}

type GetInsightsResultsMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (GetInsightsResultsMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GetInsightsResultsMetadataResponse)(nil)).Elem()
}

func (o GetInsightsResultsMetadataResponsePtrOutput) ToGetInsightsResultsMetadataResponsePtrOutput() GetInsightsResultsMetadataResponsePtrOutput {
	return o
}

func (o GetInsightsResultsMetadataResponsePtrOutput) ToGetInsightsResultsMetadataResponsePtrOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponsePtrOutput {
	return o
}

func (o GetInsightsResultsMetadataResponsePtrOutput) Elem() GetInsightsResultsMetadataResponseOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) GetInsightsResultsMetadataResponse {
		if v != nil {
			return *v
		}
		var ret GetInsightsResultsMetadataResponse
		return ret
	}).(GetInsightsResultsMetadataResponseOutput)
}

func (o GetInsightsResultsMetadataResponsePtrOutput) Errors() GetInsightsErrorKindResponseArrayOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) []GetInsightsErrorKindResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(GetInsightsErrorKindResponseArrayOutput)
}

func (o GetInsightsResultsMetadataResponsePtrOutput) TotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TotalCount
	}).(pulumi.IntPtrOutput)
}

type GitHubResourceInfo struct {
	AppInstallationId *string `pulumi:"appInstallationId"`
}





type GitHubResourceInfoInput interface {
	pulumi.Input

	ToGitHubResourceInfoOutput() GitHubResourceInfoOutput
	ToGitHubResourceInfoOutputWithContext(context.Context) GitHubResourceInfoOutput
}

type GitHubResourceInfoArgs struct {
	AppInstallationId pulumi.StringPtrInput `pulumi:"appInstallationId"`
}

func (GitHubResourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubResourceInfo)(nil)).Elem()
}

func (i GitHubResourceInfoArgs) ToGitHubResourceInfoOutput() GitHubResourceInfoOutput {
	return i.ToGitHubResourceInfoOutputWithContext(context.Background())
}

func (i GitHubResourceInfoArgs) ToGitHubResourceInfoOutputWithContext(ctx context.Context) GitHubResourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubResourceInfoOutput)
}

func (i GitHubResourceInfoArgs) ToGitHubResourceInfoPtrOutput() GitHubResourceInfoPtrOutput {
	return i.ToGitHubResourceInfoPtrOutputWithContext(context.Background())
}

func (i GitHubResourceInfoArgs) ToGitHubResourceInfoPtrOutputWithContext(ctx context.Context) GitHubResourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubResourceInfoOutput).ToGitHubResourceInfoPtrOutputWithContext(ctx)
}









type GitHubResourceInfoPtrInput interface {
	pulumi.Input

	ToGitHubResourceInfoPtrOutput() GitHubResourceInfoPtrOutput
	ToGitHubResourceInfoPtrOutputWithContext(context.Context) GitHubResourceInfoPtrOutput
}

type gitHubResourceInfoPtrType GitHubResourceInfoArgs

func GitHubResourceInfoPtr(v *GitHubResourceInfoArgs) GitHubResourceInfoPtrInput {
	return (*gitHubResourceInfoPtrType)(v)
}

func (*gitHubResourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubResourceInfo)(nil)).Elem()
}

func (i *gitHubResourceInfoPtrType) ToGitHubResourceInfoPtrOutput() GitHubResourceInfoPtrOutput {
	return i.ToGitHubResourceInfoPtrOutputWithContext(context.Background())
}

func (i *gitHubResourceInfoPtrType) ToGitHubResourceInfoPtrOutputWithContext(ctx context.Context) GitHubResourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubResourceInfoPtrOutput)
}

type GitHubResourceInfoOutput struct{ *pulumi.OutputState }

func (GitHubResourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubResourceInfo)(nil)).Elem()
}

func (o GitHubResourceInfoOutput) ToGitHubResourceInfoOutput() GitHubResourceInfoOutput {
	return o
}

func (o GitHubResourceInfoOutput) ToGitHubResourceInfoOutputWithContext(ctx context.Context) GitHubResourceInfoOutput {
	return o
}

func (o GitHubResourceInfoOutput) ToGitHubResourceInfoPtrOutput() GitHubResourceInfoPtrOutput {
	return o.ToGitHubResourceInfoPtrOutputWithContext(context.Background())
}

func (o GitHubResourceInfoOutput) ToGitHubResourceInfoPtrOutputWithContext(ctx context.Context) GitHubResourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubResourceInfo) *GitHubResourceInfo {
		return &v
	}).(GitHubResourceInfoPtrOutput)
}

func (o GitHubResourceInfoOutput) AppInstallationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubResourceInfo) *string { return v.AppInstallationId }).(pulumi.StringPtrOutput)
}

type GitHubResourceInfoPtrOutput struct{ *pulumi.OutputState }

func (GitHubResourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubResourceInfo)(nil)).Elem()
}

func (o GitHubResourceInfoPtrOutput) ToGitHubResourceInfoPtrOutput() GitHubResourceInfoPtrOutput {
	return o
}

func (o GitHubResourceInfoPtrOutput) ToGitHubResourceInfoPtrOutputWithContext(ctx context.Context) GitHubResourceInfoPtrOutput {
	return o
}

func (o GitHubResourceInfoPtrOutput) Elem() GitHubResourceInfoOutput {
	return o.ApplyT(func(v *GitHubResourceInfo) GitHubResourceInfo {
		if v != nil {
			return *v
		}
		var ret GitHubResourceInfo
		return ret
	}).(GitHubResourceInfoOutput)
}

func (o GitHubResourceInfoPtrOutput) AppInstallationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubResourceInfo) *string {
		if v == nil {
			return nil
		}
		return v.AppInstallationId
	}).(pulumi.StringPtrOutput)
}

type GitHubResourceInfoResponse struct {
	AppInstallationId *string `pulumi:"appInstallationId"`
}

type GitHubResourceInfoResponseOutput struct{ *pulumi.OutputState }

func (GitHubResourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubResourceInfoResponse)(nil)).Elem()
}

func (o GitHubResourceInfoResponseOutput) ToGitHubResourceInfoResponseOutput() GitHubResourceInfoResponseOutput {
	return o
}

func (o GitHubResourceInfoResponseOutput) ToGitHubResourceInfoResponseOutputWithContext(ctx context.Context) GitHubResourceInfoResponseOutput {
	return o
}

func (o GitHubResourceInfoResponseOutput) AppInstallationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubResourceInfoResponse) *string { return v.AppInstallationId }).(pulumi.StringPtrOutput)
}

type GitHubResourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubResourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubResourceInfoResponse)(nil)).Elem()
}

func (o GitHubResourceInfoResponsePtrOutput) ToGitHubResourceInfoResponsePtrOutput() GitHubResourceInfoResponsePtrOutput {
	return o
}

func (o GitHubResourceInfoResponsePtrOutput) ToGitHubResourceInfoResponsePtrOutputWithContext(ctx context.Context) GitHubResourceInfoResponsePtrOutput {
	return o
}

func (o GitHubResourceInfoResponsePtrOutput) Elem() GitHubResourceInfoResponseOutput {
	return o.ApplyT(func(v *GitHubResourceInfoResponse) GitHubResourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret GitHubResourceInfoResponse
		return ret
	}).(GitHubResourceInfoResponseOutput)
}

func (o GitHubResourceInfoResponsePtrOutput) AppInstallationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubResourceInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppInstallationId
	}).(pulumi.StringPtrOutput)
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
	AlertProductNames   []string `pulumi:"alertProductNames"`
	AlertsCount         int      `pulumi:"alertsCount"`
	BookmarksCount      int      `pulumi:"bookmarksCount"`
	CommentsCount       int      `pulumi:"commentsCount"`
	ProviderIncidentUrl string   `pulumi:"providerIncidentUrl"`
	Tactics             []string `pulumi:"tactics"`
	Techniques          []string `pulumi:"techniques"`
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

func (o IncidentAdditionalDataResponseOutput) ProviderIncidentUrl() pulumi.StringOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) string { return v.ProviderIncidentUrl }).(pulumi.StringOutput)
}

func (o IncidentAdditionalDataResponseOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o IncidentAdditionalDataResponseOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v IncidentAdditionalDataResponse) []string { return v.Techniques }).(pulumi.StringArrayOutput)
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

type InsightsTableResultResponse struct {
	Columns []InsightsTableResultResponseColumns `pulumi:"columns"`
	Rows    [][]string                           `pulumi:"rows"`
}

type InsightsTableResultResponseOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InsightsTableResultResponse)(nil)).Elem()
}

func (o InsightsTableResultResponseOutput) ToInsightsTableResultResponseOutput() InsightsTableResultResponseOutput {
	return o
}

func (o InsightsTableResultResponseOutput) ToInsightsTableResultResponseOutputWithContext(ctx context.Context) InsightsTableResultResponseOutput {
	return o
}

func (o InsightsTableResultResponseOutput) Columns() InsightsTableResultResponseColumnsArrayOutput {
	return o.ApplyT(func(v InsightsTableResultResponse) []InsightsTableResultResponseColumns { return v.Columns }).(InsightsTableResultResponseColumnsArrayOutput)
}

func (o InsightsTableResultResponseOutput) Rows() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v InsightsTableResultResponse) [][]string { return v.Rows }).(pulumi.StringArrayArrayOutput)
}

type InsightsTableResultResponsePtrOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InsightsTableResultResponse)(nil)).Elem()
}

func (o InsightsTableResultResponsePtrOutput) ToInsightsTableResultResponsePtrOutput() InsightsTableResultResponsePtrOutput {
	return o
}

func (o InsightsTableResultResponsePtrOutput) ToInsightsTableResultResponsePtrOutputWithContext(ctx context.Context) InsightsTableResultResponsePtrOutput {
	return o
}

func (o InsightsTableResultResponsePtrOutput) Elem() InsightsTableResultResponseOutput {
	return o.ApplyT(func(v *InsightsTableResultResponse) InsightsTableResultResponse {
		if v != nil {
			return *v
		}
		var ret InsightsTableResultResponse
		return ret
	}).(InsightsTableResultResponseOutput)
}

func (o InsightsTableResultResponsePtrOutput) Columns() InsightsTableResultResponseColumnsArrayOutput {
	return o.ApplyT(func(v *InsightsTableResultResponse) []InsightsTableResultResponseColumns {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(InsightsTableResultResponseColumnsArrayOutput)
}

func (o InsightsTableResultResponsePtrOutput) Rows() pulumi.StringArrayArrayOutput {
	return o.ApplyT(func(v *InsightsTableResultResponse) [][]string {
		if v == nil {
			return nil
		}
		return v.Rows
	}).(pulumi.StringArrayArrayOutput)
}

type InsightsTableResultResponseArrayOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InsightsTableResultResponse)(nil)).Elem()
}

func (o InsightsTableResultResponseArrayOutput) ToInsightsTableResultResponseArrayOutput() InsightsTableResultResponseArrayOutput {
	return o
}

func (o InsightsTableResultResponseArrayOutput) ToInsightsTableResultResponseArrayOutputWithContext(ctx context.Context) InsightsTableResultResponseArrayOutput {
	return o
}

func (o InsightsTableResultResponseArrayOutput) Index(i pulumi.IntInput) InsightsTableResultResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InsightsTableResultResponse {
		return vs[0].([]InsightsTableResultResponse)[vs[1].(int)]
	}).(InsightsTableResultResponseOutput)
}

type InsightsTableResultResponseColumns struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}

type InsightsTableResultResponseColumnsOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseColumnsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InsightsTableResultResponseColumns)(nil)).Elem()
}

func (o InsightsTableResultResponseColumnsOutput) ToInsightsTableResultResponseColumnsOutput() InsightsTableResultResponseColumnsOutput {
	return o
}

func (o InsightsTableResultResponseColumnsOutput) ToInsightsTableResultResponseColumnsOutputWithContext(ctx context.Context) InsightsTableResultResponseColumnsOutput {
	return o
}

func (o InsightsTableResultResponseColumnsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InsightsTableResultResponseColumns) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o InsightsTableResultResponseColumnsOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InsightsTableResultResponseColumns) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type InsightsTableResultResponseColumnsArrayOutput struct{ *pulumi.OutputState }

func (InsightsTableResultResponseColumnsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InsightsTableResultResponseColumns)(nil)).Elem()
}

func (o InsightsTableResultResponseColumnsArrayOutput) ToInsightsTableResultResponseColumnsArrayOutput() InsightsTableResultResponseColumnsArrayOutput {
	return o
}

func (o InsightsTableResultResponseColumnsArrayOutput) ToInsightsTableResultResponseColumnsArrayOutputWithContext(ctx context.Context) InsightsTableResultResponseColumnsArrayOutput {
	return o
}

func (o InsightsTableResultResponseColumnsArrayOutput) Index(i pulumi.IntInput) InsightsTableResultResponseColumnsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InsightsTableResultResponseColumns {
		return vs[0].([]InsightsTableResultResponseColumns)[vs[1].(int)]
	}).(InsightsTableResultResponseColumnsOutput)
}

type InstructionStepsInstructions struct {
	Parameters interface{} `pulumi:"parameters"`
	Type       string      `pulumi:"type"`
}





type InstructionStepsInstructionsInput interface {
	pulumi.Input

	ToInstructionStepsInstructionsOutput() InstructionStepsInstructionsOutput
	ToInstructionStepsInstructionsOutputWithContext(context.Context) InstructionStepsInstructionsOutput
}

type InstructionStepsInstructionsArgs struct {
	Parameters pulumi.Input       `pulumi:"parameters"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (InstructionStepsInstructionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstructionStepsInstructions)(nil)).Elem()
}

func (i InstructionStepsInstructionsArgs) ToInstructionStepsInstructionsOutput() InstructionStepsInstructionsOutput {
	return i.ToInstructionStepsInstructionsOutputWithContext(context.Background())
}

func (i InstructionStepsInstructionsArgs) ToInstructionStepsInstructionsOutputWithContext(ctx context.Context) InstructionStepsInstructionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstructionStepsInstructionsOutput)
}





type InstructionStepsInstructionsArrayInput interface {
	pulumi.Input

	ToInstructionStepsInstructionsArrayOutput() InstructionStepsInstructionsArrayOutput
	ToInstructionStepsInstructionsArrayOutputWithContext(context.Context) InstructionStepsInstructionsArrayOutput
}

type InstructionStepsInstructionsArray []InstructionStepsInstructionsInput

func (InstructionStepsInstructionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstructionStepsInstructions)(nil)).Elem()
}

func (i InstructionStepsInstructionsArray) ToInstructionStepsInstructionsArrayOutput() InstructionStepsInstructionsArrayOutput {
	return i.ToInstructionStepsInstructionsArrayOutputWithContext(context.Background())
}

func (i InstructionStepsInstructionsArray) ToInstructionStepsInstructionsArrayOutputWithContext(ctx context.Context) InstructionStepsInstructionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstructionStepsInstructionsArrayOutput)
}

type InstructionStepsInstructionsOutput struct{ *pulumi.OutputState }

func (InstructionStepsInstructionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstructionStepsInstructions)(nil)).Elem()
}

func (o InstructionStepsInstructionsOutput) ToInstructionStepsInstructionsOutput() InstructionStepsInstructionsOutput {
	return o
}

func (o InstructionStepsInstructionsOutput) ToInstructionStepsInstructionsOutputWithContext(ctx context.Context) InstructionStepsInstructionsOutput {
	return o
}

func (o InstructionStepsInstructionsOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v InstructionStepsInstructions) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o InstructionStepsInstructionsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InstructionStepsInstructions) string { return v.Type }).(pulumi.StringOutput)
}

type InstructionStepsInstructionsArrayOutput struct{ *pulumi.OutputState }

func (InstructionStepsInstructionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstructionStepsInstructions)(nil)).Elem()
}

func (o InstructionStepsInstructionsArrayOutput) ToInstructionStepsInstructionsArrayOutput() InstructionStepsInstructionsArrayOutput {
	return o
}

func (o InstructionStepsInstructionsArrayOutput) ToInstructionStepsInstructionsArrayOutputWithContext(ctx context.Context) InstructionStepsInstructionsArrayOutput {
	return o
}

func (o InstructionStepsInstructionsArrayOutput) Index(i pulumi.IntInput) InstructionStepsInstructionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstructionStepsInstructions {
		return vs[0].([]InstructionStepsInstructions)[vs[1].(int)]
	}).(InstructionStepsInstructionsOutput)
}

type InstructionStepsResponseInstructions struct {
	Parameters interface{} `pulumi:"parameters"`
	Type       string      `pulumi:"type"`
}

type InstructionStepsResponseInstructionsOutput struct{ *pulumi.OutputState }

func (InstructionStepsResponseInstructionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InstructionStepsResponseInstructions)(nil)).Elem()
}

func (o InstructionStepsResponseInstructionsOutput) ToInstructionStepsResponseInstructionsOutput() InstructionStepsResponseInstructionsOutput {
	return o
}

func (o InstructionStepsResponseInstructionsOutput) ToInstructionStepsResponseInstructionsOutputWithContext(ctx context.Context) InstructionStepsResponseInstructionsOutput {
	return o
}

func (o InstructionStepsResponseInstructionsOutput) Parameters() pulumi.AnyOutput {
	return o.ApplyT(func(v InstructionStepsResponseInstructions) interface{} { return v.Parameters }).(pulumi.AnyOutput)
}

func (o InstructionStepsResponseInstructionsOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v InstructionStepsResponseInstructions) string { return v.Type }).(pulumi.StringOutput)
}

type InstructionStepsResponseInstructionsArrayOutput struct{ *pulumi.OutputState }

func (InstructionStepsResponseInstructionsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstructionStepsResponseInstructions)(nil)).Elem()
}

func (o InstructionStepsResponseInstructionsArrayOutput) ToInstructionStepsResponseInstructionsArrayOutput() InstructionStepsResponseInstructionsArrayOutput {
	return o
}

func (o InstructionStepsResponseInstructionsArrayOutput) ToInstructionStepsResponseInstructionsArrayOutputWithContext(ctx context.Context) InstructionStepsResponseInstructionsArrayOutput {
	return o
}

func (o InstructionStepsResponseInstructionsArrayOutput) Index(i pulumi.IntInput) InstructionStepsResponseInstructionsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InstructionStepsResponseInstructions {
		return vs[0].([]InstructionStepsResponseInstructions)[vs[1].(int)]
	}).(InstructionStepsResponseInstructionsOutput)
}

type MCASDataConnectorDataTypes struct {
	Alerts        DataConnectorDataTypeCommon  `pulumi:"alerts"`
	DiscoveryLogs *DataConnectorDataTypeCommon `pulumi:"discoveryLogs"`
}





type MCASDataConnectorDataTypesInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesOutput() MCASDataConnectorDataTypesOutput
	ToMCASDataConnectorDataTypesOutputWithContext(context.Context) MCASDataConnectorDataTypesOutput
}

type MCASDataConnectorDataTypesArgs struct {
	Alerts        DataConnectorDataTypeCommonInput    `pulumi:"alerts"`
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

func (o MCASDataConnectorDataTypesOutput) Alerts() DataConnectorDataTypeCommonOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) DataConnectorDataTypeCommon { return v.Alerts }).(DataConnectorDataTypeCommonOutput)
}

func (o MCASDataConnectorDataTypesOutput) DiscoveryLogs() DataConnectorDataTypeCommonPtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) *DataConnectorDataTypeCommon { return v.DiscoveryLogs }).(DataConnectorDataTypeCommonPtrOutput)
}

type MCASDataConnectorDataTypesResponse struct {
	Alerts        DataConnectorDataTypeCommonResponse  `pulumi:"alerts"`
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

func (o MCASDataConnectorDataTypesResponseOutput) Alerts() DataConnectorDataTypeCommonResponseOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) DataConnectorDataTypeCommonResponse { return v.Alerts }).(DataConnectorDataTypeCommonResponseOutput)
}

func (o MCASDataConnectorDataTypesResponseOutput) DiscoveryLogs() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) *DataConnectorDataTypeCommonResponse {
		return v.DiscoveryLogs
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
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

type MetadataAuthor struct {
	Email *string `pulumi:"email"`
	Link  *string `pulumi:"link"`
	Name  *string `pulumi:"name"`
}





type MetadataAuthorInput interface {
	pulumi.Input

	ToMetadataAuthorOutput() MetadataAuthorOutput
	ToMetadataAuthorOutputWithContext(context.Context) MetadataAuthorOutput
}

type MetadataAuthorArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Link  pulumi.StringPtrInput `pulumi:"link"`
	Name  pulumi.StringPtrInput `pulumi:"name"`
}

func (MetadataAuthorArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataAuthor)(nil)).Elem()
}

func (i MetadataAuthorArgs) ToMetadataAuthorOutput() MetadataAuthorOutput {
	return i.ToMetadataAuthorOutputWithContext(context.Background())
}

func (i MetadataAuthorArgs) ToMetadataAuthorOutputWithContext(ctx context.Context) MetadataAuthorOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataAuthorOutput)
}

func (i MetadataAuthorArgs) ToMetadataAuthorPtrOutput() MetadataAuthorPtrOutput {
	return i.ToMetadataAuthorPtrOutputWithContext(context.Background())
}

func (i MetadataAuthorArgs) ToMetadataAuthorPtrOutputWithContext(ctx context.Context) MetadataAuthorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataAuthorOutput).ToMetadataAuthorPtrOutputWithContext(ctx)
}









type MetadataAuthorPtrInput interface {
	pulumi.Input

	ToMetadataAuthorPtrOutput() MetadataAuthorPtrOutput
	ToMetadataAuthorPtrOutputWithContext(context.Context) MetadataAuthorPtrOutput
}

type metadataAuthorPtrType MetadataAuthorArgs

func MetadataAuthorPtr(v *MetadataAuthorArgs) MetadataAuthorPtrInput {
	return (*metadataAuthorPtrType)(v)
}

func (*metadataAuthorPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataAuthor)(nil)).Elem()
}

func (i *metadataAuthorPtrType) ToMetadataAuthorPtrOutput() MetadataAuthorPtrOutput {
	return i.ToMetadataAuthorPtrOutputWithContext(context.Background())
}

func (i *metadataAuthorPtrType) ToMetadataAuthorPtrOutputWithContext(ctx context.Context) MetadataAuthorPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataAuthorPtrOutput)
}

type MetadataAuthorOutput struct{ *pulumi.OutputState }

func (MetadataAuthorOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataAuthor)(nil)).Elem()
}

func (o MetadataAuthorOutput) ToMetadataAuthorOutput() MetadataAuthorOutput {
	return o
}

func (o MetadataAuthorOutput) ToMetadataAuthorOutputWithContext(ctx context.Context) MetadataAuthorOutput {
	return o
}

func (o MetadataAuthorOutput) ToMetadataAuthorPtrOutput() MetadataAuthorPtrOutput {
	return o.ToMetadataAuthorPtrOutputWithContext(context.Background())
}

func (o MetadataAuthorOutput) ToMetadataAuthorPtrOutputWithContext(ctx context.Context) MetadataAuthorPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataAuthor) *MetadataAuthor {
		return &v
	}).(MetadataAuthorPtrOutput)
}

func (o MetadataAuthorOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataAuthor) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataAuthor) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataAuthor) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type MetadataAuthorPtrOutput struct{ *pulumi.OutputState }

func (MetadataAuthorPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataAuthor)(nil)).Elem()
}

func (o MetadataAuthorPtrOutput) ToMetadataAuthorPtrOutput() MetadataAuthorPtrOutput {
	return o
}

func (o MetadataAuthorPtrOutput) ToMetadataAuthorPtrOutputWithContext(ctx context.Context) MetadataAuthorPtrOutput {
	return o
}

func (o MetadataAuthorPtrOutput) Elem() MetadataAuthorOutput {
	return o.ApplyT(func(v *MetadataAuthor) MetadataAuthor {
		if v != nil {
			return *v
		}
		var ret MetadataAuthor
		return ret
	}).(MetadataAuthorOutput)
}

func (o MetadataAuthorPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataAuthor) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorPtrOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataAuthor) *string {
		if v == nil {
			return nil
		}
		return v.Link
	}).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataAuthor) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type MetadataAuthorResponse struct {
	Email *string `pulumi:"email"`
	Link  *string `pulumi:"link"`
	Name  *string `pulumi:"name"`
}

type MetadataAuthorResponseOutput struct{ *pulumi.OutputState }

func (MetadataAuthorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataAuthorResponse)(nil)).Elem()
}

func (o MetadataAuthorResponseOutput) ToMetadataAuthorResponseOutput() MetadataAuthorResponseOutput {
	return o
}

func (o MetadataAuthorResponseOutput) ToMetadataAuthorResponseOutputWithContext(ctx context.Context) MetadataAuthorResponseOutput {
	return o
}

func (o MetadataAuthorResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataAuthorResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataAuthorResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataAuthorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type MetadataAuthorResponsePtrOutput struct{ *pulumi.OutputState }

func (MetadataAuthorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataAuthorResponse)(nil)).Elem()
}

func (o MetadataAuthorResponsePtrOutput) ToMetadataAuthorResponsePtrOutput() MetadataAuthorResponsePtrOutput {
	return o
}

func (o MetadataAuthorResponsePtrOutput) ToMetadataAuthorResponsePtrOutputWithContext(ctx context.Context) MetadataAuthorResponsePtrOutput {
	return o
}

func (o MetadataAuthorResponsePtrOutput) Elem() MetadataAuthorResponseOutput {
	return o.ApplyT(func(v *MetadataAuthorResponse) MetadataAuthorResponse {
		if v != nil {
			return *v
		}
		var ret MetadataAuthorResponse
		return ret
	}).(MetadataAuthorResponseOutput)
}

func (o MetadataAuthorResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataAuthorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorResponsePtrOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataAuthorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Link
	}).(pulumi.StringPtrOutput)
}

func (o MetadataAuthorResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataAuthorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type MetadataCategories struct {
	Domains   []string `pulumi:"domains"`
	Verticals []string `pulumi:"verticals"`
}





type MetadataCategoriesInput interface {
	pulumi.Input

	ToMetadataCategoriesOutput() MetadataCategoriesOutput
	ToMetadataCategoriesOutputWithContext(context.Context) MetadataCategoriesOutput
}

type MetadataCategoriesArgs struct {
	Domains   pulumi.StringArrayInput `pulumi:"domains"`
	Verticals pulumi.StringArrayInput `pulumi:"verticals"`
}

func (MetadataCategoriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataCategories)(nil)).Elem()
}

func (i MetadataCategoriesArgs) ToMetadataCategoriesOutput() MetadataCategoriesOutput {
	return i.ToMetadataCategoriesOutputWithContext(context.Background())
}

func (i MetadataCategoriesArgs) ToMetadataCategoriesOutputWithContext(ctx context.Context) MetadataCategoriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataCategoriesOutput)
}

func (i MetadataCategoriesArgs) ToMetadataCategoriesPtrOutput() MetadataCategoriesPtrOutput {
	return i.ToMetadataCategoriesPtrOutputWithContext(context.Background())
}

func (i MetadataCategoriesArgs) ToMetadataCategoriesPtrOutputWithContext(ctx context.Context) MetadataCategoriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataCategoriesOutput).ToMetadataCategoriesPtrOutputWithContext(ctx)
}









type MetadataCategoriesPtrInput interface {
	pulumi.Input

	ToMetadataCategoriesPtrOutput() MetadataCategoriesPtrOutput
	ToMetadataCategoriesPtrOutputWithContext(context.Context) MetadataCategoriesPtrOutput
}

type metadataCategoriesPtrType MetadataCategoriesArgs

func MetadataCategoriesPtr(v *MetadataCategoriesArgs) MetadataCategoriesPtrInput {
	return (*metadataCategoriesPtrType)(v)
}

func (*metadataCategoriesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataCategories)(nil)).Elem()
}

func (i *metadataCategoriesPtrType) ToMetadataCategoriesPtrOutput() MetadataCategoriesPtrOutput {
	return i.ToMetadataCategoriesPtrOutputWithContext(context.Background())
}

func (i *metadataCategoriesPtrType) ToMetadataCategoriesPtrOutputWithContext(ctx context.Context) MetadataCategoriesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataCategoriesPtrOutput)
}

type MetadataCategoriesOutput struct{ *pulumi.OutputState }

func (MetadataCategoriesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataCategories)(nil)).Elem()
}

func (o MetadataCategoriesOutput) ToMetadataCategoriesOutput() MetadataCategoriesOutput {
	return o
}

func (o MetadataCategoriesOutput) ToMetadataCategoriesOutputWithContext(ctx context.Context) MetadataCategoriesOutput {
	return o
}

func (o MetadataCategoriesOutput) ToMetadataCategoriesPtrOutput() MetadataCategoriesPtrOutput {
	return o.ToMetadataCategoriesPtrOutputWithContext(context.Background())
}

func (o MetadataCategoriesOutput) ToMetadataCategoriesPtrOutputWithContext(ctx context.Context) MetadataCategoriesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataCategories) *MetadataCategories {
		return &v
	}).(MetadataCategoriesPtrOutput)
}

func (o MetadataCategoriesOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetadataCategories) []string { return v.Domains }).(pulumi.StringArrayOutput)
}

func (o MetadataCategoriesOutput) Verticals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetadataCategories) []string { return v.Verticals }).(pulumi.StringArrayOutput)
}

type MetadataCategoriesPtrOutput struct{ *pulumi.OutputState }

func (MetadataCategoriesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataCategories)(nil)).Elem()
}

func (o MetadataCategoriesPtrOutput) ToMetadataCategoriesPtrOutput() MetadataCategoriesPtrOutput {
	return o
}

func (o MetadataCategoriesPtrOutput) ToMetadataCategoriesPtrOutputWithContext(ctx context.Context) MetadataCategoriesPtrOutput {
	return o
}

func (o MetadataCategoriesPtrOutput) Elem() MetadataCategoriesOutput {
	return o.ApplyT(func(v *MetadataCategories) MetadataCategories {
		if v != nil {
			return *v
		}
		var ret MetadataCategories
		return ret
	}).(MetadataCategoriesOutput)
}

func (o MetadataCategoriesPtrOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetadataCategories) []string {
		if v == nil {
			return nil
		}
		return v.Domains
	}).(pulumi.StringArrayOutput)
}

func (o MetadataCategoriesPtrOutput) Verticals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetadataCategories) []string {
		if v == nil {
			return nil
		}
		return v.Verticals
	}).(pulumi.StringArrayOutput)
}

type MetadataCategoriesResponse struct {
	Domains   []string `pulumi:"domains"`
	Verticals []string `pulumi:"verticals"`
}

type MetadataCategoriesResponseOutput struct{ *pulumi.OutputState }

func (MetadataCategoriesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataCategoriesResponse)(nil)).Elem()
}

func (o MetadataCategoriesResponseOutput) ToMetadataCategoriesResponseOutput() MetadataCategoriesResponseOutput {
	return o
}

func (o MetadataCategoriesResponseOutput) ToMetadataCategoriesResponseOutputWithContext(ctx context.Context) MetadataCategoriesResponseOutput {
	return o
}

func (o MetadataCategoriesResponseOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetadataCategoriesResponse) []string { return v.Domains }).(pulumi.StringArrayOutput)
}

func (o MetadataCategoriesResponseOutput) Verticals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MetadataCategoriesResponse) []string { return v.Verticals }).(pulumi.StringArrayOutput)
}

type MetadataCategoriesResponsePtrOutput struct{ *pulumi.OutputState }

func (MetadataCategoriesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataCategoriesResponse)(nil)).Elem()
}

func (o MetadataCategoriesResponsePtrOutput) ToMetadataCategoriesResponsePtrOutput() MetadataCategoriesResponsePtrOutput {
	return o
}

func (o MetadataCategoriesResponsePtrOutput) ToMetadataCategoriesResponsePtrOutputWithContext(ctx context.Context) MetadataCategoriesResponsePtrOutput {
	return o
}

func (o MetadataCategoriesResponsePtrOutput) Elem() MetadataCategoriesResponseOutput {
	return o.ApplyT(func(v *MetadataCategoriesResponse) MetadataCategoriesResponse {
		if v != nil {
			return *v
		}
		var ret MetadataCategoriesResponse
		return ret
	}).(MetadataCategoriesResponseOutput)
}

func (o MetadataCategoriesResponsePtrOutput) Domains() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetadataCategoriesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Domains
	}).(pulumi.StringArrayOutput)
}

func (o MetadataCategoriesResponsePtrOutput) Verticals() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MetadataCategoriesResponse) []string {
		if v == nil {
			return nil
		}
		return v.Verticals
	}).(pulumi.StringArrayOutput)
}

type MetadataDependencies struct {
	ContentId *string                `pulumi:"contentId"`
	Criteria  []MetadataDependencies `pulumi:"criteria"`
	Kind      *string                `pulumi:"kind"`
	Name      *string                `pulumi:"name"`
	Operator  *string                `pulumi:"operator"`
	Version   *string                `pulumi:"version"`
}





type MetadataDependenciesInput interface {
	pulumi.Input

	ToMetadataDependenciesOutput() MetadataDependenciesOutput
	ToMetadataDependenciesOutputWithContext(context.Context) MetadataDependenciesOutput
}

type MetadataDependenciesArgs struct {
	ContentId pulumi.StringPtrInput          `pulumi:"contentId"`
	Criteria  MetadataDependenciesArrayInput `pulumi:"criteria"`
	Kind      pulumi.StringPtrInput          `pulumi:"kind"`
	Name      pulumi.StringPtrInput          `pulumi:"name"`
	Operator  pulumi.StringPtrInput          `pulumi:"operator"`
	Version   pulumi.StringPtrInput          `pulumi:"version"`
}

func (MetadataDependenciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataDependencies)(nil)).Elem()
}

func (i MetadataDependenciesArgs) ToMetadataDependenciesOutput() MetadataDependenciesOutput {
	return i.ToMetadataDependenciesOutputWithContext(context.Background())
}

func (i MetadataDependenciesArgs) ToMetadataDependenciesOutputWithContext(ctx context.Context) MetadataDependenciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesOutput)
}

func (i MetadataDependenciesArgs) ToMetadataDependenciesPtrOutput() MetadataDependenciesPtrOutput {
	return i.ToMetadataDependenciesPtrOutputWithContext(context.Background())
}

func (i MetadataDependenciesArgs) ToMetadataDependenciesPtrOutputWithContext(ctx context.Context) MetadataDependenciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesOutput).ToMetadataDependenciesPtrOutputWithContext(ctx)
}









type MetadataDependenciesPtrInput interface {
	pulumi.Input

	ToMetadataDependenciesPtrOutput() MetadataDependenciesPtrOutput
	ToMetadataDependenciesPtrOutputWithContext(context.Context) MetadataDependenciesPtrOutput
}

type metadataDependenciesPtrType MetadataDependenciesArgs

func MetadataDependenciesPtr(v *MetadataDependenciesArgs) MetadataDependenciesPtrInput {
	return (*metadataDependenciesPtrType)(v)
}

func (*metadataDependenciesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataDependencies)(nil)).Elem()
}

func (i *metadataDependenciesPtrType) ToMetadataDependenciesPtrOutput() MetadataDependenciesPtrOutput {
	return i.ToMetadataDependenciesPtrOutputWithContext(context.Background())
}

func (i *metadataDependenciesPtrType) ToMetadataDependenciesPtrOutputWithContext(ctx context.Context) MetadataDependenciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesPtrOutput)
}





type MetadataDependenciesArrayInput interface {
	pulumi.Input

	ToMetadataDependenciesArrayOutput() MetadataDependenciesArrayOutput
	ToMetadataDependenciesArrayOutputWithContext(context.Context) MetadataDependenciesArrayOutput
}

type MetadataDependenciesArray []MetadataDependenciesInput

func (MetadataDependenciesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataDependencies)(nil)).Elem()
}

func (i MetadataDependenciesArray) ToMetadataDependenciesArrayOutput() MetadataDependenciesArrayOutput {
	return i.ToMetadataDependenciesArrayOutputWithContext(context.Background())
}

func (i MetadataDependenciesArray) ToMetadataDependenciesArrayOutputWithContext(ctx context.Context) MetadataDependenciesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesArrayOutput)
}

type MetadataDependenciesOutput struct{ *pulumi.OutputState }

func (MetadataDependenciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataDependencies)(nil)).Elem()
}

func (o MetadataDependenciesOutput) ToMetadataDependenciesOutput() MetadataDependenciesOutput {
	return o
}

func (o MetadataDependenciesOutput) ToMetadataDependenciesOutputWithContext(ctx context.Context) MetadataDependenciesOutput {
	return o
}

func (o MetadataDependenciesOutput) ToMetadataDependenciesPtrOutput() MetadataDependenciesPtrOutput {
	return o.ToMetadataDependenciesPtrOutputWithContext(context.Background())
}

func (o MetadataDependenciesOutput) ToMetadataDependenciesPtrOutputWithContext(ctx context.Context) MetadataDependenciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataDependencies) *MetadataDependencies {
		return &v
	}).(MetadataDependenciesPtrOutput)
}

func (o MetadataDependenciesOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependencies) *string { return v.ContentId }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesOutput) Criteria() MetadataDependenciesArrayOutput {
	return o.ApplyT(func(v MetadataDependencies) []MetadataDependencies { return v.Criteria }).(MetadataDependenciesArrayOutput)
}

func (o MetadataDependenciesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependencies) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependencies) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependencies) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependencies) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type MetadataDependenciesPtrOutput struct{ *pulumi.OutputState }

func (MetadataDependenciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataDependencies)(nil)).Elem()
}

func (o MetadataDependenciesPtrOutput) ToMetadataDependenciesPtrOutput() MetadataDependenciesPtrOutput {
	return o
}

func (o MetadataDependenciesPtrOutput) ToMetadataDependenciesPtrOutputWithContext(ctx context.Context) MetadataDependenciesPtrOutput {
	return o
}

func (o MetadataDependenciesPtrOutput) Elem() MetadataDependenciesOutput {
	return o.ApplyT(func(v *MetadataDependencies) MetadataDependencies {
		if v != nil {
			return *v
		}
		var ret MetadataDependencies
		return ret
	}).(MetadataDependenciesOutput)
}

func (o MetadataDependenciesPtrOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependencies) *string {
		if v == nil {
			return nil
		}
		return v.ContentId
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesPtrOutput) Criteria() MetadataDependenciesArrayOutput {
	return o.ApplyT(func(v *MetadataDependencies) []MetadataDependencies {
		if v == nil {
			return nil
		}
		return v.Criteria
	}).(MetadataDependenciesArrayOutput)
}

func (o MetadataDependenciesPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependencies) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependencies) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesPtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependencies) *string {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependencies) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type MetadataDependenciesArrayOutput struct{ *pulumi.OutputState }

func (MetadataDependenciesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataDependencies)(nil)).Elem()
}

func (o MetadataDependenciesArrayOutput) ToMetadataDependenciesArrayOutput() MetadataDependenciesArrayOutput {
	return o
}

func (o MetadataDependenciesArrayOutput) ToMetadataDependenciesArrayOutputWithContext(ctx context.Context) MetadataDependenciesArrayOutput {
	return o
}

func (o MetadataDependenciesArrayOutput) Index(i pulumi.IntInput) MetadataDependenciesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetadataDependencies {
		return vs[0].([]MetadataDependencies)[vs[1].(int)]
	}).(MetadataDependenciesOutput)
}

type MetadataDependenciesResponse struct {
	ContentId *string                        `pulumi:"contentId"`
	Criteria  []MetadataDependenciesResponse `pulumi:"criteria"`
	Kind      *string                        `pulumi:"kind"`
	Name      *string                        `pulumi:"name"`
	Operator  *string                        `pulumi:"operator"`
	Version   *string                        `pulumi:"version"`
}

type MetadataDependenciesResponseOutput struct{ *pulumi.OutputState }

func (MetadataDependenciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataDependenciesResponse)(nil)).Elem()
}

func (o MetadataDependenciesResponseOutput) ToMetadataDependenciesResponseOutput() MetadataDependenciesResponseOutput {
	return o
}

func (o MetadataDependenciesResponseOutput) ToMetadataDependenciesResponseOutputWithContext(ctx context.Context) MetadataDependenciesResponseOutput {
	return o
}

func (o MetadataDependenciesResponseOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependenciesResponse) *string { return v.ContentId }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponseOutput) Criteria() MetadataDependenciesResponseArrayOutput {
	return o.ApplyT(func(v MetadataDependenciesResponse) []MetadataDependenciesResponse { return v.Criteria }).(MetadataDependenciesResponseArrayOutput)
}

func (o MetadataDependenciesResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependenciesResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependenciesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponseOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependenciesResponse) *string { return v.Operator }).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataDependenciesResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type MetadataDependenciesResponsePtrOutput struct{ *pulumi.OutputState }

func (MetadataDependenciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataDependenciesResponse)(nil)).Elem()
}

func (o MetadataDependenciesResponsePtrOutput) ToMetadataDependenciesResponsePtrOutput() MetadataDependenciesResponsePtrOutput {
	return o
}

func (o MetadataDependenciesResponsePtrOutput) ToMetadataDependenciesResponsePtrOutputWithContext(ctx context.Context) MetadataDependenciesResponsePtrOutput {
	return o
}

func (o MetadataDependenciesResponsePtrOutput) Elem() MetadataDependenciesResponseOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) MetadataDependenciesResponse {
		if v != nil {
			return *v
		}
		var ret MetadataDependenciesResponse
		return ret
	}).(MetadataDependenciesResponseOutput)
}

func (o MetadataDependenciesResponsePtrOutput) ContentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentId
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponsePtrOutput) Criteria() MetadataDependenciesResponseArrayOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) []MetadataDependenciesResponse {
		if v == nil {
			return nil
		}
		return v.Criteria
	}).(MetadataDependenciesResponseArrayOutput)
}

func (o MetadataDependenciesResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponsePtrOutput) Operator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Operator
	}).(pulumi.StringPtrOutput)
}

func (o MetadataDependenciesResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataDependenciesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type MetadataDependenciesResponseArrayOutput struct{ *pulumi.OutputState }

func (MetadataDependenciesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataDependenciesResponse)(nil)).Elem()
}

func (o MetadataDependenciesResponseArrayOutput) ToMetadataDependenciesResponseArrayOutput() MetadataDependenciesResponseArrayOutput {
	return o
}

func (o MetadataDependenciesResponseArrayOutput) ToMetadataDependenciesResponseArrayOutputWithContext(ctx context.Context) MetadataDependenciesResponseArrayOutput {
	return o
}

func (o MetadataDependenciesResponseArrayOutput) Index(i pulumi.IntInput) MetadataDependenciesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetadataDependenciesResponse {
		return vs[0].([]MetadataDependenciesResponse)[vs[1].(int)]
	}).(MetadataDependenciesResponseOutput)
}

type MetadataSource struct {
	Kind     string  `pulumi:"kind"`
	Name     *string `pulumi:"name"`
	SourceId *string `pulumi:"sourceId"`
}





type MetadataSourceInput interface {
	pulumi.Input

	ToMetadataSourceOutput() MetadataSourceOutput
	ToMetadataSourceOutputWithContext(context.Context) MetadataSourceOutput
}

type MetadataSourceArgs struct {
	Kind     pulumi.StringInput    `pulumi:"kind"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	SourceId pulumi.StringPtrInput `pulumi:"sourceId"`
}

func (MetadataSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSource)(nil)).Elem()
}

func (i MetadataSourceArgs) ToMetadataSourceOutput() MetadataSourceOutput {
	return i.ToMetadataSourceOutputWithContext(context.Background())
}

func (i MetadataSourceArgs) ToMetadataSourceOutputWithContext(ctx context.Context) MetadataSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSourceOutput)
}

func (i MetadataSourceArgs) ToMetadataSourcePtrOutput() MetadataSourcePtrOutput {
	return i.ToMetadataSourcePtrOutputWithContext(context.Background())
}

func (i MetadataSourceArgs) ToMetadataSourcePtrOutputWithContext(ctx context.Context) MetadataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSourceOutput).ToMetadataSourcePtrOutputWithContext(ctx)
}









type MetadataSourcePtrInput interface {
	pulumi.Input

	ToMetadataSourcePtrOutput() MetadataSourcePtrOutput
	ToMetadataSourcePtrOutputWithContext(context.Context) MetadataSourcePtrOutput
}

type metadataSourcePtrType MetadataSourceArgs

func MetadataSourcePtr(v *MetadataSourceArgs) MetadataSourcePtrInput {
	return (*metadataSourcePtrType)(v)
}

func (*metadataSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSource)(nil)).Elem()
}

func (i *metadataSourcePtrType) ToMetadataSourcePtrOutput() MetadataSourcePtrOutput {
	return i.ToMetadataSourcePtrOutputWithContext(context.Background())
}

func (i *metadataSourcePtrType) ToMetadataSourcePtrOutputWithContext(ctx context.Context) MetadataSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSourcePtrOutput)
}

type MetadataSourceOutput struct{ *pulumi.OutputState }

func (MetadataSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSource)(nil)).Elem()
}

func (o MetadataSourceOutput) ToMetadataSourceOutput() MetadataSourceOutput {
	return o
}

func (o MetadataSourceOutput) ToMetadataSourceOutputWithContext(ctx context.Context) MetadataSourceOutput {
	return o
}

func (o MetadataSourceOutput) ToMetadataSourcePtrOutput() MetadataSourcePtrOutput {
	return o.ToMetadataSourcePtrOutputWithContext(context.Background())
}

func (o MetadataSourceOutput) ToMetadataSourcePtrOutputWithContext(ctx context.Context) MetadataSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataSource) *MetadataSource {
		return &v
	}).(MetadataSourcePtrOutput)
}

func (o MetadataSourceOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataSource) string { return v.Kind }).(pulumi.StringOutput)
}

func (o MetadataSourceOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSource) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MetadataSourceOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSource) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

type MetadataSourcePtrOutput struct{ *pulumi.OutputState }

func (MetadataSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSource)(nil)).Elem()
}

func (o MetadataSourcePtrOutput) ToMetadataSourcePtrOutput() MetadataSourcePtrOutput {
	return o
}

func (o MetadataSourcePtrOutput) ToMetadataSourcePtrOutputWithContext(ctx context.Context) MetadataSourcePtrOutput {
	return o
}

func (o MetadataSourcePtrOutput) Elem() MetadataSourceOutput {
	return o.ApplyT(func(v *MetadataSource) MetadataSource {
		if v != nil {
			return *v
		}
		var ret MetadataSource
		return ret
	}).(MetadataSourceOutput)
}

func (o MetadataSourcePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSource) *string {
		if v == nil {
			return nil
		}
		return &v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSourcePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSource) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSourcePtrOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSource) *string {
		if v == nil {
			return nil
		}
		return v.SourceId
	}).(pulumi.StringPtrOutput)
}

type MetadataSourceResponse struct {
	Kind     string  `pulumi:"kind"`
	Name     *string `pulumi:"name"`
	SourceId *string `pulumi:"sourceId"`
}

type MetadataSourceResponseOutput struct{ *pulumi.OutputState }

func (MetadataSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSourceResponse)(nil)).Elem()
}

func (o MetadataSourceResponseOutput) ToMetadataSourceResponseOutput() MetadataSourceResponseOutput {
	return o
}

func (o MetadataSourceResponseOutput) ToMetadataSourceResponseOutputWithContext(ctx context.Context) MetadataSourceResponseOutput {
	return o
}

func (o MetadataSourceResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataSourceResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o MetadataSourceResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSourceResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MetadataSourceResponseOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSourceResponse) *string { return v.SourceId }).(pulumi.StringPtrOutput)
}

type MetadataSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (MetadataSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSourceResponse)(nil)).Elem()
}

func (o MetadataSourceResponsePtrOutput) ToMetadataSourceResponsePtrOutput() MetadataSourceResponsePtrOutput {
	return o
}

func (o MetadataSourceResponsePtrOutput) ToMetadataSourceResponsePtrOutputWithContext(ctx context.Context) MetadataSourceResponsePtrOutput {
	return o
}

func (o MetadataSourceResponsePtrOutput) Elem() MetadataSourceResponseOutput {
	return o.ApplyT(func(v *MetadataSourceResponse) MetadataSourceResponse {
		if v != nil {
			return *v
		}
		var ret MetadataSourceResponse
		return ret
	}).(MetadataSourceResponseOutput)
}

func (o MetadataSourceResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSourceResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSourceResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSourceResponsePtrOutput) SourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceId
	}).(pulumi.StringPtrOutput)
}

type MetadataSupport struct {
	Email *string `pulumi:"email"`
	Link  *string `pulumi:"link"`
	Name  *string `pulumi:"name"`
	Tier  string  `pulumi:"tier"`
}





type MetadataSupportInput interface {
	pulumi.Input

	ToMetadataSupportOutput() MetadataSupportOutput
	ToMetadataSupportOutputWithContext(context.Context) MetadataSupportOutput
}

type MetadataSupportArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Link  pulumi.StringPtrInput `pulumi:"link"`
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Tier  pulumi.StringInput    `pulumi:"tier"`
}

func (MetadataSupportArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSupport)(nil)).Elem()
}

func (i MetadataSupportArgs) ToMetadataSupportOutput() MetadataSupportOutput {
	return i.ToMetadataSupportOutputWithContext(context.Background())
}

func (i MetadataSupportArgs) ToMetadataSupportOutputWithContext(ctx context.Context) MetadataSupportOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSupportOutput)
}

func (i MetadataSupportArgs) ToMetadataSupportPtrOutput() MetadataSupportPtrOutput {
	return i.ToMetadataSupportPtrOutputWithContext(context.Background())
}

func (i MetadataSupportArgs) ToMetadataSupportPtrOutputWithContext(ctx context.Context) MetadataSupportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSupportOutput).ToMetadataSupportPtrOutputWithContext(ctx)
}









type MetadataSupportPtrInput interface {
	pulumi.Input

	ToMetadataSupportPtrOutput() MetadataSupportPtrOutput
	ToMetadataSupportPtrOutputWithContext(context.Context) MetadataSupportPtrOutput
}

type metadataSupportPtrType MetadataSupportArgs

func MetadataSupportPtr(v *MetadataSupportArgs) MetadataSupportPtrInput {
	return (*metadataSupportPtrType)(v)
}

func (*metadataSupportPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSupport)(nil)).Elem()
}

func (i *metadataSupportPtrType) ToMetadataSupportPtrOutput() MetadataSupportPtrOutput {
	return i.ToMetadataSupportPtrOutputWithContext(context.Background())
}

func (i *metadataSupportPtrType) ToMetadataSupportPtrOutputWithContext(ctx context.Context) MetadataSupportPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSupportPtrOutput)
}

type MetadataSupportOutput struct{ *pulumi.OutputState }

func (MetadataSupportOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSupport)(nil)).Elem()
}

func (o MetadataSupportOutput) ToMetadataSupportOutput() MetadataSupportOutput {
	return o
}

func (o MetadataSupportOutput) ToMetadataSupportOutputWithContext(ctx context.Context) MetadataSupportOutput {
	return o
}

func (o MetadataSupportOutput) ToMetadataSupportPtrOutput() MetadataSupportPtrOutput {
	return o.ToMetadataSupportPtrOutputWithContext(context.Background())
}

func (o MetadataSupportOutput) ToMetadataSupportPtrOutputWithContext(ctx context.Context) MetadataSupportPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataSupport) *MetadataSupport {
		return &v
	}).(MetadataSupportPtrOutput)
}

func (o MetadataSupportOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSupport) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o MetadataSupportOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSupport) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o MetadataSupportOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSupport) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MetadataSupportOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataSupport) string { return v.Tier }).(pulumi.StringOutput)
}

type MetadataSupportPtrOutput struct{ *pulumi.OutputState }

func (MetadataSupportPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSupport)(nil)).Elem()
}

func (o MetadataSupportPtrOutput) ToMetadataSupportPtrOutput() MetadataSupportPtrOutput {
	return o
}

func (o MetadataSupportPtrOutput) ToMetadataSupportPtrOutputWithContext(ctx context.Context) MetadataSupportPtrOutput {
	return o
}

func (o MetadataSupportPtrOutput) Elem() MetadataSupportOutput {
	return o.ApplyT(func(v *MetadataSupport) MetadataSupport {
		if v != nil {
			return *v
		}
		var ret MetadataSupport
		return ret
	}).(MetadataSupportOutput)
}

func (o MetadataSupportPtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupport) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSupportPtrOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupport) *string {
		if v == nil {
			return nil
		}
		return v.Link
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSupportPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupport) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSupportPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupport) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type MetadataSupportResponse struct {
	Email *string `pulumi:"email"`
	Link  *string `pulumi:"link"`
	Name  *string `pulumi:"name"`
	Tier  string  `pulumi:"tier"`
}

type MetadataSupportResponseOutput struct{ *pulumi.OutputState }

func (MetadataSupportResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSupportResponse)(nil)).Elem()
}

func (o MetadataSupportResponseOutput) ToMetadataSupportResponseOutput() MetadataSupportResponseOutput {
	return o
}

func (o MetadataSupportResponseOutput) ToMetadataSupportResponseOutputWithContext(ctx context.Context) MetadataSupportResponseOutput {
	return o
}

func (o MetadataSupportResponseOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSupportResponse) *string { return v.Email }).(pulumi.StringPtrOutput)
}

func (o MetadataSupportResponseOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSupportResponse) *string { return v.Link }).(pulumi.StringPtrOutput)
}

func (o MetadataSupportResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetadataSupportResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o MetadataSupportResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v MetadataSupportResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type MetadataSupportResponsePtrOutput struct{ *pulumi.OutputState }

func (MetadataSupportResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSupportResponse)(nil)).Elem()
}

func (o MetadataSupportResponsePtrOutput) ToMetadataSupportResponsePtrOutput() MetadataSupportResponsePtrOutput {
	return o
}

func (o MetadataSupportResponsePtrOutput) ToMetadataSupportResponsePtrOutputWithContext(ctx context.Context) MetadataSupportResponsePtrOutput {
	return o
}

func (o MetadataSupportResponsePtrOutput) Elem() MetadataSupportResponseOutput {
	return o.ApplyT(func(v *MetadataSupportResponse) MetadataSupportResponse {
		if v != nil {
			return *v
		}
		var ret MetadataSupportResponse
		return ret
	}).(MetadataSupportResponseOutput)
}

func (o MetadataSupportResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupportResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSupportResponsePtrOutput) Link() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupportResponse) *string {
		if v == nil {
			return nil
		}
		return v.Link
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSupportResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupportResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o MetadataSupportResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MetadataSupportResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type Office365ProjectConnectorDataTypes struct {
	Logs Office365ProjectConnectorDataTypesLogs `pulumi:"logs"`
}





type Office365ProjectConnectorDataTypesInput interface {
	pulumi.Input

	ToOffice365ProjectConnectorDataTypesOutput() Office365ProjectConnectorDataTypesOutput
	ToOffice365ProjectConnectorDataTypesOutputWithContext(context.Context) Office365ProjectConnectorDataTypesOutput
}

type Office365ProjectConnectorDataTypesArgs struct {
	Logs Office365ProjectConnectorDataTypesLogsInput `pulumi:"logs"`
}

func (Office365ProjectConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Office365ProjectConnectorDataTypes)(nil)).Elem()
}

func (i Office365ProjectConnectorDataTypesArgs) ToOffice365ProjectConnectorDataTypesOutput() Office365ProjectConnectorDataTypesOutput {
	return i.ToOffice365ProjectConnectorDataTypesOutputWithContext(context.Background())
}

func (i Office365ProjectConnectorDataTypesArgs) ToOffice365ProjectConnectorDataTypesOutputWithContext(ctx context.Context) Office365ProjectConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Office365ProjectConnectorDataTypesOutput)
}

type Office365ProjectConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (Office365ProjectConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Office365ProjectConnectorDataTypes)(nil)).Elem()
}

func (o Office365ProjectConnectorDataTypesOutput) ToOffice365ProjectConnectorDataTypesOutput() Office365ProjectConnectorDataTypesOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesOutput) ToOffice365ProjectConnectorDataTypesOutputWithContext(ctx context.Context) Office365ProjectConnectorDataTypesOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesOutput) Logs() Office365ProjectConnectorDataTypesLogsOutput {
	return o.ApplyT(func(v Office365ProjectConnectorDataTypes) Office365ProjectConnectorDataTypesLogs { return v.Logs }).(Office365ProjectConnectorDataTypesLogsOutput)
}

type Office365ProjectConnectorDataTypesLogs struct {
	State string `pulumi:"state"`
}





type Office365ProjectConnectorDataTypesLogsInput interface {
	pulumi.Input

	ToOffice365ProjectConnectorDataTypesLogsOutput() Office365ProjectConnectorDataTypesLogsOutput
	ToOffice365ProjectConnectorDataTypesLogsOutputWithContext(context.Context) Office365ProjectConnectorDataTypesLogsOutput
}

type Office365ProjectConnectorDataTypesLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (Office365ProjectConnectorDataTypesLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Office365ProjectConnectorDataTypesLogs)(nil)).Elem()
}

func (i Office365ProjectConnectorDataTypesLogsArgs) ToOffice365ProjectConnectorDataTypesLogsOutput() Office365ProjectConnectorDataTypesLogsOutput {
	return i.ToOffice365ProjectConnectorDataTypesLogsOutputWithContext(context.Background())
}

func (i Office365ProjectConnectorDataTypesLogsArgs) ToOffice365ProjectConnectorDataTypesLogsOutputWithContext(ctx context.Context) Office365ProjectConnectorDataTypesLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Office365ProjectConnectorDataTypesLogsOutput)
}

type Office365ProjectConnectorDataTypesLogsOutput struct{ *pulumi.OutputState }

func (Office365ProjectConnectorDataTypesLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Office365ProjectConnectorDataTypesLogs)(nil)).Elem()
}

func (o Office365ProjectConnectorDataTypesLogsOutput) ToOffice365ProjectConnectorDataTypesLogsOutput() Office365ProjectConnectorDataTypesLogsOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesLogsOutput) ToOffice365ProjectConnectorDataTypesLogsOutputWithContext(ctx context.Context) Office365ProjectConnectorDataTypesLogsOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Office365ProjectConnectorDataTypesLogs) string { return v.State }).(pulumi.StringOutput)
}

type Office365ProjectConnectorDataTypesResponse struct {
	Logs Office365ProjectConnectorDataTypesResponseLogs `pulumi:"logs"`
}

type Office365ProjectConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (Office365ProjectConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Office365ProjectConnectorDataTypesResponse)(nil)).Elem()
}

func (o Office365ProjectConnectorDataTypesResponseOutput) ToOffice365ProjectConnectorDataTypesResponseOutput() Office365ProjectConnectorDataTypesResponseOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesResponseOutput) ToOffice365ProjectConnectorDataTypesResponseOutputWithContext(ctx context.Context) Office365ProjectConnectorDataTypesResponseOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesResponseOutput) Logs() Office365ProjectConnectorDataTypesResponseLogsOutput {
	return o.ApplyT(func(v Office365ProjectConnectorDataTypesResponse) Office365ProjectConnectorDataTypesResponseLogs {
		return v.Logs
	}).(Office365ProjectConnectorDataTypesResponseLogsOutput)
}

type Office365ProjectConnectorDataTypesResponseLogs struct {
	State string `pulumi:"state"`
}

type Office365ProjectConnectorDataTypesResponseLogsOutput struct{ *pulumi.OutputState }

func (Office365ProjectConnectorDataTypesResponseLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Office365ProjectConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (o Office365ProjectConnectorDataTypesResponseLogsOutput) ToOffice365ProjectConnectorDataTypesResponseLogsOutput() Office365ProjectConnectorDataTypesResponseLogsOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesResponseLogsOutput) ToOffice365ProjectConnectorDataTypesResponseLogsOutputWithContext(ctx context.Context) Office365ProjectConnectorDataTypesResponseLogsOutput {
	return o
}

func (o Office365ProjectConnectorDataTypesResponseLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Office365ProjectConnectorDataTypesResponseLogs) string { return v.State }).(pulumi.StringOutput)
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

type OfficePowerBIConnectorDataTypes struct {
	Logs OfficePowerBIConnectorDataTypesLogs `pulumi:"logs"`
}





type OfficePowerBIConnectorDataTypesInput interface {
	pulumi.Input

	ToOfficePowerBIConnectorDataTypesOutput() OfficePowerBIConnectorDataTypesOutput
	ToOfficePowerBIConnectorDataTypesOutputWithContext(context.Context) OfficePowerBIConnectorDataTypesOutput
}

type OfficePowerBIConnectorDataTypesArgs struct {
	Logs OfficePowerBIConnectorDataTypesLogsInput `pulumi:"logs"`
}

func (OfficePowerBIConnectorDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficePowerBIConnectorDataTypes)(nil)).Elem()
}

func (i OfficePowerBIConnectorDataTypesArgs) ToOfficePowerBIConnectorDataTypesOutput() OfficePowerBIConnectorDataTypesOutput {
	return i.ToOfficePowerBIConnectorDataTypesOutputWithContext(context.Background())
}

func (i OfficePowerBIConnectorDataTypesArgs) ToOfficePowerBIConnectorDataTypesOutputWithContext(ctx context.Context) OfficePowerBIConnectorDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficePowerBIConnectorDataTypesOutput)
}

type OfficePowerBIConnectorDataTypesOutput struct{ *pulumi.OutputState }

func (OfficePowerBIConnectorDataTypesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficePowerBIConnectorDataTypes)(nil)).Elem()
}

func (o OfficePowerBIConnectorDataTypesOutput) ToOfficePowerBIConnectorDataTypesOutput() OfficePowerBIConnectorDataTypesOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesOutput) ToOfficePowerBIConnectorDataTypesOutputWithContext(ctx context.Context) OfficePowerBIConnectorDataTypesOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesOutput) Logs() OfficePowerBIConnectorDataTypesLogsOutput {
	return o.ApplyT(func(v OfficePowerBIConnectorDataTypes) OfficePowerBIConnectorDataTypesLogs { return v.Logs }).(OfficePowerBIConnectorDataTypesLogsOutput)
}

type OfficePowerBIConnectorDataTypesLogs struct {
	State string `pulumi:"state"`
}





type OfficePowerBIConnectorDataTypesLogsInput interface {
	pulumi.Input

	ToOfficePowerBIConnectorDataTypesLogsOutput() OfficePowerBIConnectorDataTypesLogsOutput
	ToOfficePowerBIConnectorDataTypesLogsOutputWithContext(context.Context) OfficePowerBIConnectorDataTypesLogsOutput
}

type OfficePowerBIConnectorDataTypesLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (OfficePowerBIConnectorDataTypesLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficePowerBIConnectorDataTypesLogs)(nil)).Elem()
}

func (i OfficePowerBIConnectorDataTypesLogsArgs) ToOfficePowerBIConnectorDataTypesLogsOutput() OfficePowerBIConnectorDataTypesLogsOutput {
	return i.ToOfficePowerBIConnectorDataTypesLogsOutputWithContext(context.Background())
}

func (i OfficePowerBIConnectorDataTypesLogsArgs) ToOfficePowerBIConnectorDataTypesLogsOutputWithContext(ctx context.Context) OfficePowerBIConnectorDataTypesLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficePowerBIConnectorDataTypesLogsOutput)
}

type OfficePowerBIConnectorDataTypesLogsOutput struct{ *pulumi.OutputState }

func (OfficePowerBIConnectorDataTypesLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficePowerBIConnectorDataTypesLogs)(nil)).Elem()
}

func (o OfficePowerBIConnectorDataTypesLogsOutput) ToOfficePowerBIConnectorDataTypesLogsOutput() OfficePowerBIConnectorDataTypesLogsOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesLogsOutput) ToOfficePowerBIConnectorDataTypesLogsOutputWithContext(ctx context.Context) OfficePowerBIConnectorDataTypesLogsOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficePowerBIConnectorDataTypesLogs) string { return v.State }).(pulumi.StringOutput)
}

type OfficePowerBIConnectorDataTypesResponse struct {
	Logs OfficePowerBIConnectorDataTypesResponseLogs `pulumi:"logs"`
}

type OfficePowerBIConnectorDataTypesResponseOutput struct{ *pulumi.OutputState }

func (OfficePowerBIConnectorDataTypesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficePowerBIConnectorDataTypesResponse)(nil)).Elem()
}

func (o OfficePowerBIConnectorDataTypesResponseOutput) ToOfficePowerBIConnectorDataTypesResponseOutput() OfficePowerBIConnectorDataTypesResponseOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesResponseOutput) ToOfficePowerBIConnectorDataTypesResponseOutputWithContext(ctx context.Context) OfficePowerBIConnectorDataTypesResponseOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesResponseOutput) Logs() OfficePowerBIConnectorDataTypesResponseLogsOutput {
	return o.ApplyT(func(v OfficePowerBIConnectorDataTypesResponse) OfficePowerBIConnectorDataTypesResponseLogs {
		return v.Logs
	}).(OfficePowerBIConnectorDataTypesResponseLogsOutput)
}

type OfficePowerBIConnectorDataTypesResponseLogs struct {
	State string `pulumi:"state"`
}

type OfficePowerBIConnectorDataTypesResponseLogsOutput struct{ *pulumi.OutputState }

func (OfficePowerBIConnectorDataTypesResponseLogsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficePowerBIConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (o OfficePowerBIConnectorDataTypesResponseLogsOutput) ToOfficePowerBIConnectorDataTypesResponseLogsOutput() OfficePowerBIConnectorDataTypesResponseLogsOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesResponseLogsOutput) ToOfficePowerBIConnectorDataTypesResponseLogsOutputWithContext(ctx context.Context) OfficePowerBIConnectorDataTypesResponseLogsOutput {
	return o
}

func (o OfficePowerBIConnectorDataTypesResponseLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficePowerBIConnectorDataTypesResponseLogs) string { return v.State }).(pulumi.StringOutput)
}

type Permissions struct {
	Customs          []PermissionsCustoms          `pulumi:"customs"`
	ResourceProvider []PermissionsResourceProvider `pulumi:"resourceProvider"`
}





type PermissionsInput interface {
	pulumi.Input

	ToPermissionsOutput() PermissionsOutput
	ToPermissionsOutputWithContext(context.Context) PermissionsOutput
}

type PermissionsArgs struct {
	Customs          PermissionsCustomsArrayInput          `pulumi:"customs"`
	ResourceProvider PermissionsResourceProviderArrayInput `pulumi:"resourceProvider"`
}

func (PermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (i PermissionsArgs) ToPermissionsOutput() PermissionsOutput {
	return i.ToPermissionsOutputWithContext(context.Background())
}

func (i PermissionsArgs) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsOutput)
}

func (i PermissionsArgs) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return i.ToPermissionsPtrOutputWithContext(context.Background())
}

func (i PermissionsArgs) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsOutput).ToPermissionsPtrOutputWithContext(ctx)
}









type PermissionsPtrInput interface {
	pulumi.Input

	ToPermissionsPtrOutput() PermissionsPtrOutput
	ToPermissionsPtrOutputWithContext(context.Context) PermissionsPtrOutput
}

type permissionsPtrType PermissionsArgs

func PermissionsPtr(v *PermissionsArgs) PermissionsPtrInput {
	return (*permissionsPtrType)(v)
}

func (*permissionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (i *permissionsPtrType) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return i.ToPermissionsPtrOutputWithContext(context.Background())
}

func (i *permissionsPtrType) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsPtrOutput)
}

type PermissionsOutput struct{ *pulumi.OutputState }

func (PermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permissions)(nil)).Elem()
}

func (o PermissionsOutput) ToPermissionsOutput() PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsOutputWithContext(ctx context.Context) PermissionsOutput {
	return o
}

func (o PermissionsOutput) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return o.ToPermissionsPtrOutputWithContext(context.Background())
}

func (o PermissionsOutput) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Permissions) *Permissions {
		return &v
	}).(PermissionsPtrOutput)
}

func (o PermissionsOutput) Customs() PermissionsCustomsArrayOutput {
	return o.ApplyT(func(v Permissions) []PermissionsCustoms { return v.Customs }).(PermissionsCustomsArrayOutput)
}

func (o PermissionsOutput) ResourceProvider() PermissionsResourceProviderArrayOutput {
	return o.ApplyT(func(v Permissions) []PermissionsResourceProvider { return v.ResourceProvider }).(PermissionsResourceProviderArrayOutput)
}

type PermissionsPtrOutput struct{ *pulumi.OutputState }

func (PermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Permissions)(nil)).Elem()
}

func (o PermissionsPtrOutput) ToPermissionsPtrOutput() PermissionsPtrOutput {
	return o
}

func (o PermissionsPtrOutput) ToPermissionsPtrOutputWithContext(ctx context.Context) PermissionsPtrOutput {
	return o
}

func (o PermissionsPtrOutput) Elem() PermissionsOutput {
	return o.ApplyT(func(v *Permissions) Permissions {
		if v != nil {
			return *v
		}
		var ret Permissions
		return ret
	}).(PermissionsOutput)
}

func (o PermissionsPtrOutput) Customs() PermissionsCustomsArrayOutput {
	return o.ApplyT(func(v *Permissions) []PermissionsCustoms {
		if v == nil {
			return nil
		}
		return v.Customs
	}).(PermissionsCustomsArrayOutput)
}

func (o PermissionsPtrOutput) ResourceProvider() PermissionsResourceProviderArrayOutput {
	return o.ApplyT(func(v *Permissions) []PermissionsResourceProvider {
		if v == nil {
			return nil
		}
		return v.ResourceProvider
	}).(PermissionsResourceProviderArrayOutput)
}

type PermissionsCustoms struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
}





type PermissionsCustomsInput interface {
	pulumi.Input

	ToPermissionsCustomsOutput() PermissionsCustomsOutput
	ToPermissionsCustomsOutputWithContext(context.Context) PermissionsCustomsOutput
}

type PermissionsCustomsArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (PermissionsCustomsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsCustoms)(nil)).Elem()
}

func (i PermissionsCustomsArgs) ToPermissionsCustomsOutput() PermissionsCustomsOutput {
	return i.ToPermissionsCustomsOutputWithContext(context.Background())
}

func (i PermissionsCustomsArgs) ToPermissionsCustomsOutputWithContext(ctx context.Context) PermissionsCustomsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsCustomsOutput)
}





type PermissionsCustomsArrayInput interface {
	pulumi.Input

	ToPermissionsCustomsArrayOutput() PermissionsCustomsArrayOutput
	ToPermissionsCustomsArrayOutputWithContext(context.Context) PermissionsCustomsArrayOutput
}

type PermissionsCustomsArray []PermissionsCustomsInput

func (PermissionsCustomsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsCustoms)(nil)).Elem()
}

func (i PermissionsCustomsArray) ToPermissionsCustomsArrayOutput() PermissionsCustomsArrayOutput {
	return i.ToPermissionsCustomsArrayOutputWithContext(context.Background())
}

func (i PermissionsCustomsArray) ToPermissionsCustomsArrayOutputWithContext(ctx context.Context) PermissionsCustomsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsCustomsArrayOutput)
}

type PermissionsCustomsOutput struct{ *pulumi.OutputState }

func (PermissionsCustomsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsCustoms)(nil)).Elem()
}

func (o PermissionsCustomsOutput) ToPermissionsCustomsOutput() PermissionsCustomsOutput {
	return o
}

func (o PermissionsCustomsOutput) ToPermissionsCustomsOutputWithContext(ctx context.Context) PermissionsCustomsOutput {
	return o
}

func (o PermissionsCustomsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsCustoms) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PermissionsCustomsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsCustoms) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PermissionsCustomsArrayOutput struct{ *pulumi.OutputState }

func (PermissionsCustomsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsCustoms)(nil)).Elem()
}

func (o PermissionsCustomsArrayOutput) ToPermissionsCustomsArrayOutput() PermissionsCustomsArrayOutput {
	return o
}

func (o PermissionsCustomsArrayOutput) ToPermissionsCustomsArrayOutputWithContext(ctx context.Context) PermissionsCustomsArrayOutput {
	return o
}

func (o PermissionsCustomsArrayOutput) Index(i pulumi.IntInput) PermissionsCustomsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionsCustoms {
		return vs[0].([]PermissionsCustoms)[vs[1].(int)]
	}).(PermissionsCustomsOutput)
}

type PermissionsResourceProvider struct {
	PermissionsDisplayText *string              `pulumi:"permissionsDisplayText"`
	Provider               *string              `pulumi:"provider"`
	ProviderDisplayName    *string              `pulumi:"providerDisplayName"`
	RequiredPermissions    *RequiredPermissions `pulumi:"requiredPermissions"`
	Scope                  *string              `pulumi:"scope"`
}





type PermissionsResourceProviderInput interface {
	pulumi.Input

	ToPermissionsResourceProviderOutput() PermissionsResourceProviderOutput
	ToPermissionsResourceProviderOutputWithContext(context.Context) PermissionsResourceProviderOutput
}

type PermissionsResourceProviderArgs struct {
	PermissionsDisplayText pulumi.StringPtrInput       `pulumi:"permissionsDisplayText"`
	Provider               pulumi.StringPtrInput       `pulumi:"provider"`
	ProviderDisplayName    pulumi.StringPtrInput       `pulumi:"providerDisplayName"`
	RequiredPermissions    RequiredPermissionsPtrInput `pulumi:"requiredPermissions"`
	Scope                  pulumi.StringPtrInput       `pulumi:"scope"`
}

func (PermissionsResourceProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResourceProvider)(nil)).Elem()
}

func (i PermissionsResourceProviderArgs) ToPermissionsResourceProviderOutput() PermissionsResourceProviderOutput {
	return i.ToPermissionsResourceProviderOutputWithContext(context.Background())
}

func (i PermissionsResourceProviderArgs) ToPermissionsResourceProviderOutputWithContext(ctx context.Context) PermissionsResourceProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResourceProviderOutput)
}





type PermissionsResourceProviderArrayInput interface {
	pulumi.Input

	ToPermissionsResourceProviderArrayOutput() PermissionsResourceProviderArrayOutput
	ToPermissionsResourceProviderArrayOutputWithContext(context.Context) PermissionsResourceProviderArrayOutput
}

type PermissionsResourceProviderArray []PermissionsResourceProviderInput

func (PermissionsResourceProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsResourceProvider)(nil)).Elem()
}

func (i PermissionsResourceProviderArray) ToPermissionsResourceProviderArrayOutput() PermissionsResourceProviderArrayOutput {
	return i.ToPermissionsResourceProviderArrayOutputWithContext(context.Background())
}

func (i PermissionsResourceProviderArray) ToPermissionsResourceProviderArrayOutputWithContext(ctx context.Context) PermissionsResourceProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResourceProviderArrayOutput)
}

type PermissionsResourceProviderOutput struct{ *pulumi.OutputState }

func (PermissionsResourceProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResourceProvider)(nil)).Elem()
}

func (o PermissionsResourceProviderOutput) ToPermissionsResourceProviderOutput() PermissionsResourceProviderOutput {
	return o
}

func (o PermissionsResourceProviderOutput) ToPermissionsResourceProviderOutputWithContext(ctx context.Context) PermissionsResourceProviderOutput {
	return o
}

func (o PermissionsResourceProviderOutput) PermissionsDisplayText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResourceProvider) *string { return v.PermissionsDisplayText }).(pulumi.StringPtrOutput)
}

func (o PermissionsResourceProviderOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResourceProvider) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

func (o PermissionsResourceProviderOutput) ProviderDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResourceProvider) *string { return v.ProviderDisplayName }).(pulumi.StringPtrOutput)
}

func (o PermissionsResourceProviderOutput) RequiredPermissions() RequiredPermissionsPtrOutput {
	return o.ApplyT(func(v PermissionsResourceProvider) *RequiredPermissions { return v.RequiredPermissions }).(RequiredPermissionsPtrOutput)
}

func (o PermissionsResourceProviderOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResourceProvider) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type PermissionsResourceProviderArrayOutput struct{ *pulumi.OutputState }

func (PermissionsResourceProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsResourceProvider)(nil)).Elem()
}

func (o PermissionsResourceProviderArrayOutput) ToPermissionsResourceProviderArrayOutput() PermissionsResourceProviderArrayOutput {
	return o
}

func (o PermissionsResourceProviderArrayOutput) ToPermissionsResourceProviderArrayOutputWithContext(ctx context.Context) PermissionsResourceProviderArrayOutput {
	return o
}

func (o PermissionsResourceProviderArrayOutput) Index(i pulumi.IntInput) PermissionsResourceProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionsResourceProvider {
		return vs[0].([]PermissionsResourceProvider)[vs[1].(int)]
	}).(PermissionsResourceProviderOutput)
}

type PermissionsResponse struct {
	Customs          []PermissionsResponseCustoms          `pulumi:"customs"`
	ResourceProvider []PermissionsResponseResourceProvider `pulumi:"resourceProvider"`
}

type PermissionsResponseOutput struct{ *pulumi.OutputState }

func (PermissionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponse)(nil)).Elem()
}

func (o PermissionsResponseOutput) ToPermissionsResponseOutput() PermissionsResponseOutput {
	return o
}

func (o PermissionsResponseOutput) ToPermissionsResponseOutputWithContext(ctx context.Context) PermissionsResponseOutput {
	return o
}

func (o PermissionsResponseOutput) Customs() PermissionsResponseCustomsArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []PermissionsResponseCustoms { return v.Customs }).(PermissionsResponseCustomsArrayOutput)
}

func (o PermissionsResponseOutput) ResourceProvider() PermissionsResponseResourceProviderArrayOutput {
	return o.ApplyT(func(v PermissionsResponse) []PermissionsResponseResourceProvider { return v.ResourceProvider }).(PermissionsResponseResourceProviderArrayOutput)
}

type PermissionsResponsePtrOutput struct{ *pulumi.OutputState }

func (PermissionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionsResponse)(nil)).Elem()
}

func (o PermissionsResponsePtrOutput) ToPermissionsResponsePtrOutput() PermissionsResponsePtrOutput {
	return o
}

func (o PermissionsResponsePtrOutput) ToPermissionsResponsePtrOutputWithContext(ctx context.Context) PermissionsResponsePtrOutput {
	return o
}

func (o PermissionsResponsePtrOutput) Elem() PermissionsResponseOutput {
	return o.ApplyT(func(v *PermissionsResponse) PermissionsResponse {
		if v != nil {
			return *v
		}
		var ret PermissionsResponse
		return ret
	}).(PermissionsResponseOutput)
}

func (o PermissionsResponsePtrOutput) Customs() PermissionsResponseCustomsArrayOutput {
	return o.ApplyT(func(v *PermissionsResponse) []PermissionsResponseCustoms {
		if v == nil {
			return nil
		}
		return v.Customs
	}).(PermissionsResponseCustomsArrayOutput)
}

func (o PermissionsResponsePtrOutput) ResourceProvider() PermissionsResponseResourceProviderArrayOutput {
	return o.ApplyT(func(v *PermissionsResponse) []PermissionsResponseResourceProvider {
		if v == nil {
			return nil
		}
		return v.ResourceProvider
	}).(PermissionsResponseResourceProviderArrayOutput)
}

type PermissionsResponseCustoms struct {
	Description *string `pulumi:"description"`
	Name        *string `pulumi:"name"`
}

type PermissionsResponseCustomsOutput struct{ *pulumi.OutputState }

func (PermissionsResponseCustomsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponseCustoms)(nil)).Elem()
}

func (o PermissionsResponseCustomsOutput) ToPermissionsResponseCustomsOutput() PermissionsResponseCustomsOutput {
	return o
}

func (o PermissionsResponseCustomsOutput) ToPermissionsResponseCustomsOutputWithContext(ctx context.Context) PermissionsResponseCustomsOutput {
	return o
}

func (o PermissionsResponseCustomsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResponseCustoms) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PermissionsResponseCustomsOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResponseCustoms) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type PermissionsResponseCustomsArrayOutput struct{ *pulumi.OutputState }

func (PermissionsResponseCustomsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsResponseCustoms)(nil)).Elem()
}

func (o PermissionsResponseCustomsArrayOutput) ToPermissionsResponseCustomsArrayOutput() PermissionsResponseCustomsArrayOutput {
	return o
}

func (o PermissionsResponseCustomsArrayOutput) ToPermissionsResponseCustomsArrayOutputWithContext(ctx context.Context) PermissionsResponseCustomsArrayOutput {
	return o
}

func (o PermissionsResponseCustomsArrayOutput) Index(i pulumi.IntInput) PermissionsResponseCustomsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionsResponseCustoms {
		return vs[0].([]PermissionsResponseCustoms)[vs[1].(int)]
	}).(PermissionsResponseCustomsOutput)
}

type PermissionsResponseResourceProvider struct {
	PermissionsDisplayText *string                      `pulumi:"permissionsDisplayText"`
	Provider               *string                      `pulumi:"provider"`
	ProviderDisplayName    *string                      `pulumi:"providerDisplayName"`
	RequiredPermissions    *RequiredPermissionsResponse `pulumi:"requiredPermissions"`
	Scope                  *string                      `pulumi:"scope"`
}

type PermissionsResponseResourceProviderOutput struct{ *pulumi.OutputState }

func (PermissionsResponseResourceProviderOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponseResourceProvider)(nil)).Elem()
}

func (o PermissionsResponseResourceProviderOutput) ToPermissionsResponseResourceProviderOutput() PermissionsResponseResourceProviderOutput {
	return o
}

func (o PermissionsResponseResourceProviderOutput) ToPermissionsResponseResourceProviderOutputWithContext(ctx context.Context) PermissionsResponseResourceProviderOutput {
	return o
}

func (o PermissionsResponseResourceProviderOutput) PermissionsDisplayText() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResponseResourceProvider) *string { return v.PermissionsDisplayText }).(pulumi.StringPtrOutput)
}

func (o PermissionsResponseResourceProviderOutput) Provider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResponseResourceProvider) *string { return v.Provider }).(pulumi.StringPtrOutput)
}

func (o PermissionsResponseResourceProviderOutput) ProviderDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResponseResourceProvider) *string { return v.ProviderDisplayName }).(pulumi.StringPtrOutput)
}

func (o PermissionsResponseResourceProviderOutput) RequiredPermissions() RequiredPermissionsResponsePtrOutput {
	return o.ApplyT(func(v PermissionsResponseResourceProvider) *RequiredPermissionsResponse { return v.RequiredPermissions }).(RequiredPermissionsResponsePtrOutput)
}

func (o PermissionsResponseResourceProviderOutput) Scope() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PermissionsResponseResourceProvider) *string { return v.Scope }).(pulumi.StringPtrOutput)
}

type PermissionsResponseResourceProviderArrayOutput struct{ *pulumi.OutputState }

func (PermissionsResponseResourceProviderArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsResponseResourceProvider)(nil)).Elem()
}

func (o PermissionsResponseResourceProviderArrayOutput) ToPermissionsResponseResourceProviderArrayOutput() PermissionsResponseResourceProviderArrayOutput {
	return o
}

func (o PermissionsResponseResourceProviderArrayOutput) ToPermissionsResponseResourceProviderArrayOutputWithContext(ctx context.Context) PermissionsResponseResourceProviderArrayOutput {
	return o
}

func (o PermissionsResponseResourceProviderArrayOutput) Index(i pulumi.IntInput) PermissionsResponseResourceProviderOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionsResponseResourceProvider {
		return vs[0].([]PermissionsResponseResourceProvider)[vs[1].(int)]
	}).(PermissionsResponseResourceProviderOutput)
}

type PlaybookActionProperties struct {
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}

type PlaybookActionPropertiesResponse struct {
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
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

type RepoResponse struct {
	Branches []string `pulumi:"branches"`
	FullName *string  `pulumi:"fullName"`
	Url      *string  `pulumi:"url"`
}

type RepoResponseOutput struct{ *pulumi.OutputState }

func (RepoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepoResponse)(nil)).Elem()
}

func (o RepoResponseOutput) ToRepoResponseOutput() RepoResponseOutput {
	return o
}

func (o RepoResponseOutput) ToRepoResponseOutputWithContext(ctx context.Context) RepoResponseOutput {
	return o
}

func (o RepoResponseOutput) Branches() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RepoResponse) []string { return v.Branches }).(pulumi.StringArrayOutput)
}

func (o RepoResponseOutput) FullName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepoResponse) *string { return v.FullName }).(pulumi.StringPtrOutput)
}

func (o RepoResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepoResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type RepoResponseArrayOutput struct{ *pulumi.OutputState }

func (RepoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepoResponse)(nil)).Elem()
}

func (o RepoResponseArrayOutput) ToRepoResponseArrayOutput() RepoResponseArrayOutput {
	return o
}

func (o RepoResponseArrayOutput) ToRepoResponseArrayOutputWithContext(ctx context.Context) RepoResponseArrayOutput {
	return o
}

func (o RepoResponseArrayOutput) Index(i pulumi.IntInput) RepoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RepoResponse {
		return vs[0].([]RepoResponse)[vs[1].(int)]
	}).(RepoResponseOutput)
}

type Repository struct {
	Branch            *string          `pulumi:"branch"`
	DeploymentLogsUrl *string          `pulumi:"deploymentLogsUrl"`
	DisplayUrl        *string          `pulumi:"displayUrl"`
	PathMapping       []ContentPathMap `pulumi:"pathMapping"`
	Url               *string          `pulumi:"url"`
}





type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(context.Context) RepositoryOutput
}

type RepositoryArgs struct {
	Branch            pulumi.StringPtrInput    `pulumi:"branch"`
	DeploymentLogsUrl pulumi.StringPtrInput    `pulumi:"deploymentLogsUrl"`
	DisplayUrl        pulumi.StringPtrInput    `pulumi:"displayUrl"`
	PathMapping       ContentPathMapArrayInput `pulumi:"pathMapping"`
	Url               pulumi.StringPtrInput    `pulumi:"url"`
}

func (RepositoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil)).Elem()
}

func (i RepositoryArgs) ToRepositoryOutput() RepositoryOutput {
	return i.ToRepositoryOutputWithContext(context.Background())
}

func (i RepositoryArgs) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput)
}

type RepositoryOutput struct{ *pulumi.OutputState }

func (RepositoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Repository)(nil)).Elem()
}

func (o RepositoryOutput) ToRepositoryOutput() RepositoryOutput {
	return o
}

func (o RepositoryOutput) ToRepositoryOutputWithContext(ctx context.Context) RepositoryOutput {
	return o
}

func (o RepositoryOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Repository) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o RepositoryOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Repository) *string { return v.DeploymentLogsUrl }).(pulumi.StringPtrOutput)
}

func (o RepositoryOutput) DisplayUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Repository) *string { return v.DisplayUrl }).(pulumi.StringPtrOutput)
}

func (o RepositoryOutput) PathMapping() ContentPathMapArrayOutput {
	return o.ApplyT(func(v Repository) []ContentPathMap { return v.PathMapping }).(ContentPathMapArrayOutput)
}

func (o RepositoryOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Repository) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type RepositoryResourceInfo struct {
	AzureDevOpsResourceInfo *AzureDevOpsResourceInfo `pulumi:"azureDevOpsResourceInfo"`
	GitHubResourceInfo      *GitHubResourceInfo      `pulumi:"gitHubResourceInfo"`
	Webhook                 *Webhook                 `pulumi:"webhook"`
}





type RepositoryResourceInfoInput interface {
	pulumi.Input

	ToRepositoryResourceInfoOutput() RepositoryResourceInfoOutput
	ToRepositoryResourceInfoOutputWithContext(context.Context) RepositoryResourceInfoOutput
}

type RepositoryResourceInfoArgs struct {
	AzureDevOpsResourceInfo AzureDevOpsResourceInfoPtrInput `pulumi:"azureDevOpsResourceInfo"`
	GitHubResourceInfo      GitHubResourceInfoPtrInput      `pulumi:"gitHubResourceInfo"`
	Webhook                 WebhookPtrInput                 `pulumi:"webhook"`
}

func (RepositoryResourceInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryResourceInfo)(nil)).Elem()
}

func (i RepositoryResourceInfoArgs) ToRepositoryResourceInfoOutput() RepositoryResourceInfoOutput {
	return i.ToRepositoryResourceInfoOutputWithContext(context.Background())
}

func (i RepositoryResourceInfoArgs) ToRepositoryResourceInfoOutputWithContext(ctx context.Context) RepositoryResourceInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryResourceInfoOutput)
}

func (i RepositoryResourceInfoArgs) ToRepositoryResourceInfoPtrOutput() RepositoryResourceInfoPtrOutput {
	return i.ToRepositoryResourceInfoPtrOutputWithContext(context.Background())
}

func (i RepositoryResourceInfoArgs) ToRepositoryResourceInfoPtrOutputWithContext(ctx context.Context) RepositoryResourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryResourceInfoOutput).ToRepositoryResourceInfoPtrOutputWithContext(ctx)
}









type RepositoryResourceInfoPtrInput interface {
	pulumi.Input

	ToRepositoryResourceInfoPtrOutput() RepositoryResourceInfoPtrOutput
	ToRepositoryResourceInfoPtrOutputWithContext(context.Context) RepositoryResourceInfoPtrOutput
}

type repositoryResourceInfoPtrType RepositoryResourceInfoArgs

func RepositoryResourceInfoPtr(v *RepositoryResourceInfoArgs) RepositoryResourceInfoPtrInput {
	return (*repositoryResourceInfoPtrType)(v)
}

func (*repositoryResourceInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryResourceInfo)(nil)).Elem()
}

func (i *repositoryResourceInfoPtrType) ToRepositoryResourceInfoPtrOutput() RepositoryResourceInfoPtrOutput {
	return i.ToRepositoryResourceInfoPtrOutputWithContext(context.Background())
}

func (i *repositoryResourceInfoPtrType) ToRepositoryResourceInfoPtrOutputWithContext(ctx context.Context) RepositoryResourceInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryResourceInfoPtrOutput)
}

type RepositoryResourceInfoOutput struct{ *pulumi.OutputState }

func (RepositoryResourceInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryResourceInfo)(nil)).Elem()
}

func (o RepositoryResourceInfoOutput) ToRepositoryResourceInfoOutput() RepositoryResourceInfoOutput {
	return o
}

func (o RepositoryResourceInfoOutput) ToRepositoryResourceInfoOutputWithContext(ctx context.Context) RepositoryResourceInfoOutput {
	return o
}

func (o RepositoryResourceInfoOutput) ToRepositoryResourceInfoPtrOutput() RepositoryResourceInfoPtrOutput {
	return o.ToRepositoryResourceInfoPtrOutputWithContext(context.Background())
}

func (o RepositoryResourceInfoOutput) ToRepositoryResourceInfoPtrOutputWithContext(ctx context.Context) RepositoryResourceInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryResourceInfo) *RepositoryResourceInfo {
		return &v
	}).(RepositoryResourceInfoPtrOutput)
}

func (o RepositoryResourceInfoOutput) AzureDevOpsResourceInfo() AzureDevOpsResourceInfoPtrOutput {
	return o.ApplyT(func(v RepositoryResourceInfo) *AzureDevOpsResourceInfo { return v.AzureDevOpsResourceInfo }).(AzureDevOpsResourceInfoPtrOutput)
}

func (o RepositoryResourceInfoOutput) GitHubResourceInfo() GitHubResourceInfoPtrOutput {
	return o.ApplyT(func(v RepositoryResourceInfo) *GitHubResourceInfo { return v.GitHubResourceInfo }).(GitHubResourceInfoPtrOutput)
}

func (o RepositoryResourceInfoOutput) Webhook() WebhookPtrOutput {
	return o.ApplyT(func(v RepositoryResourceInfo) *Webhook { return v.Webhook }).(WebhookPtrOutput)
}

type RepositoryResourceInfoPtrOutput struct{ *pulumi.OutputState }

func (RepositoryResourceInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryResourceInfo)(nil)).Elem()
}

func (o RepositoryResourceInfoPtrOutput) ToRepositoryResourceInfoPtrOutput() RepositoryResourceInfoPtrOutput {
	return o
}

func (o RepositoryResourceInfoPtrOutput) ToRepositoryResourceInfoPtrOutputWithContext(ctx context.Context) RepositoryResourceInfoPtrOutput {
	return o
}

func (o RepositoryResourceInfoPtrOutput) Elem() RepositoryResourceInfoOutput {
	return o.ApplyT(func(v *RepositoryResourceInfo) RepositoryResourceInfo {
		if v != nil {
			return *v
		}
		var ret RepositoryResourceInfo
		return ret
	}).(RepositoryResourceInfoOutput)
}

func (o RepositoryResourceInfoPtrOutput) AzureDevOpsResourceInfo() AzureDevOpsResourceInfoPtrOutput {
	return o.ApplyT(func(v *RepositoryResourceInfo) *AzureDevOpsResourceInfo {
		if v == nil {
			return nil
		}
		return v.AzureDevOpsResourceInfo
	}).(AzureDevOpsResourceInfoPtrOutput)
}

func (o RepositoryResourceInfoPtrOutput) GitHubResourceInfo() GitHubResourceInfoPtrOutput {
	return o.ApplyT(func(v *RepositoryResourceInfo) *GitHubResourceInfo {
		if v == nil {
			return nil
		}
		return v.GitHubResourceInfo
	}).(GitHubResourceInfoPtrOutput)
}

func (o RepositoryResourceInfoPtrOutput) Webhook() WebhookPtrOutput {
	return o.ApplyT(func(v *RepositoryResourceInfo) *Webhook {
		if v == nil {
			return nil
		}
		return v.Webhook
	}).(WebhookPtrOutput)
}

type RepositoryResourceInfoResponse struct {
	AzureDevOpsResourceInfo *AzureDevOpsResourceInfoResponse `pulumi:"azureDevOpsResourceInfo"`
	GitHubResourceInfo      *GitHubResourceInfoResponse      `pulumi:"gitHubResourceInfo"`
	Webhook                 *WebhookResponse                 `pulumi:"webhook"`
}

type RepositoryResourceInfoResponseOutput struct{ *pulumi.OutputState }

func (RepositoryResourceInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryResourceInfoResponse)(nil)).Elem()
}

func (o RepositoryResourceInfoResponseOutput) ToRepositoryResourceInfoResponseOutput() RepositoryResourceInfoResponseOutput {
	return o
}

func (o RepositoryResourceInfoResponseOutput) ToRepositoryResourceInfoResponseOutputWithContext(ctx context.Context) RepositoryResourceInfoResponseOutput {
	return o
}

func (o RepositoryResourceInfoResponseOutput) AzureDevOpsResourceInfo() AzureDevOpsResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v RepositoryResourceInfoResponse) *AzureDevOpsResourceInfoResponse {
		return v.AzureDevOpsResourceInfo
	}).(AzureDevOpsResourceInfoResponsePtrOutput)
}

func (o RepositoryResourceInfoResponseOutput) GitHubResourceInfo() GitHubResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v RepositoryResourceInfoResponse) *GitHubResourceInfoResponse { return v.GitHubResourceInfo }).(GitHubResourceInfoResponsePtrOutput)
}

func (o RepositoryResourceInfoResponseOutput) Webhook() WebhookResponsePtrOutput {
	return o.ApplyT(func(v RepositoryResourceInfoResponse) *WebhookResponse { return v.Webhook }).(WebhookResponsePtrOutput)
}

type RepositoryResourceInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (RepositoryResourceInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryResourceInfoResponse)(nil)).Elem()
}

func (o RepositoryResourceInfoResponsePtrOutput) ToRepositoryResourceInfoResponsePtrOutput() RepositoryResourceInfoResponsePtrOutput {
	return o
}

func (o RepositoryResourceInfoResponsePtrOutput) ToRepositoryResourceInfoResponsePtrOutputWithContext(ctx context.Context) RepositoryResourceInfoResponsePtrOutput {
	return o
}

func (o RepositoryResourceInfoResponsePtrOutput) Elem() RepositoryResourceInfoResponseOutput {
	return o.ApplyT(func(v *RepositoryResourceInfoResponse) RepositoryResourceInfoResponse {
		if v != nil {
			return *v
		}
		var ret RepositoryResourceInfoResponse
		return ret
	}).(RepositoryResourceInfoResponseOutput)
}

func (o RepositoryResourceInfoResponsePtrOutput) AzureDevOpsResourceInfo() AzureDevOpsResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v *RepositoryResourceInfoResponse) *AzureDevOpsResourceInfoResponse {
		if v == nil {
			return nil
		}
		return v.AzureDevOpsResourceInfo
	}).(AzureDevOpsResourceInfoResponsePtrOutput)
}

func (o RepositoryResourceInfoResponsePtrOutput) GitHubResourceInfo() GitHubResourceInfoResponsePtrOutput {
	return o.ApplyT(func(v *RepositoryResourceInfoResponse) *GitHubResourceInfoResponse {
		if v == nil {
			return nil
		}
		return v.GitHubResourceInfo
	}).(GitHubResourceInfoResponsePtrOutput)
}

func (o RepositoryResourceInfoResponsePtrOutput) Webhook() WebhookResponsePtrOutput {
	return o.ApplyT(func(v *RepositoryResourceInfoResponse) *WebhookResponse {
		if v == nil {
			return nil
		}
		return v.Webhook
	}).(WebhookResponsePtrOutput)
}

type RepositoryResponse struct {
	Branch            *string                  `pulumi:"branch"`
	DeploymentLogsUrl *string                  `pulumi:"deploymentLogsUrl"`
	DisplayUrl        *string                  `pulumi:"displayUrl"`
	PathMapping       []ContentPathMapResponse `pulumi:"pathMapping"`
	Url               *string                  `pulumi:"url"`
}

type RepositoryResponseOutput struct{ *pulumi.OutputState }

func (RepositoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryResponse)(nil)).Elem()
}

func (o RepositoryResponseOutput) ToRepositoryResponseOutput() RepositoryResponseOutput {
	return o
}

func (o RepositoryResponseOutput) ToRepositoryResponseOutputWithContext(ctx context.Context) RepositoryResponseOutput {
	return o
}

func (o RepositoryResponseOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryResponse) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o RepositoryResponseOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryResponse) *string { return v.DeploymentLogsUrl }).(pulumi.StringPtrOutput)
}

func (o RepositoryResponseOutput) DisplayUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryResponse) *string { return v.DisplayUrl }).(pulumi.StringPtrOutput)
}

func (o RepositoryResponseOutput) PathMapping() ContentPathMapResponseArrayOutput {
	return o.ApplyT(func(v RepositoryResponse) []ContentPathMapResponse { return v.PathMapping }).(ContentPathMapResponseArrayOutput)
}

func (o RepositoryResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RepositoryResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type RequiredPermissions struct {
	Action *bool `pulumi:"action"`
	Delete *bool `pulumi:"delete"`
	Read   *bool `pulumi:"read"`
	Write  *bool `pulumi:"write"`
}





type RequiredPermissionsInput interface {
	pulumi.Input

	ToRequiredPermissionsOutput() RequiredPermissionsOutput
	ToRequiredPermissionsOutputWithContext(context.Context) RequiredPermissionsOutput
}

type RequiredPermissionsArgs struct {
	Action pulumi.BoolPtrInput `pulumi:"action"`
	Delete pulumi.BoolPtrInput `pulumi:"delete"`
	Read   pulumi.BoolPtrInput `pulumi:"read"`
	Write  pulumi.BoolPtrInput `pulumi:"write"`
}

func (RequiredPermissionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredPermissions)(nil)).Elem()
}

func (i RequiredPermissionsArgs) ToRequiredPermissionsOutput() RequiredPermissionsOutput {
	return i.ToRequiredPermissionsOutputWithContext(context.Background())
}

func (i RequiredPermissionsArgs) ToRequiredPermissionsOutputWithContext(ctx context.Context) RequiredPermissionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredPermissionsOutput)
}

func (i RequiredPermissionsArgs) ToRequiredPermissionsPtrOutput() RequiredPermissionsPtrOutput {
	return i.ToRequiredPermissionsPtrOutputWithContext(context.Background())
}

func (i RequiredPermissionsArgs) ToRequiredPermissionsPtrOutputWithContext(ctx context.Context) RequiredPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredPermissionsOutput).ToRequiredPermissionsPtrOutputWithContext(ctx)
}









type RequiredPermissionsPtrInput interface {
	pulumi.Input

	ToRequiredPermissionsPtrOutput() RequiredPermissionsPtrOutput
	ToRequiredPermissionsPtrOutputWithContext(context.Context) RequiredPermissionsPtrOutput
}

type requiredPermissionsPtrType RequiredPermissionsArgs

func RequiredPermissionsPtr(v *RequiredPermissionsArgs) RequiredPermissionsPtrInput {
	return (*requiredPermissionsPtrType)(v)
}

func (*requiredPermissionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequiredPermissions)(nil)).Elem()
}

func (i *requiredPermissionsPtrType) ToRequiredPermissionsPtrOutput() RequiredPermissionsPtrOutput {
	return i.ToRequiredPermissionsPtrOutputWithContext(context.Background())
}

func (i *requiredPermissionsPtrType) ToRequiredPermissionsPtrOutputWithContext(ctx context.Context) RequiredPermissionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredPermissionsPtrOutput)
}

type RequiredPermissionsOutput struct{ *pulumi.OutputState }

func (RequiredPermissionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredPermissions)(nil)).Elem()
}

func (o RequiredPermissionsOutput) ToRequiredPermissionsOutput() RequiredPermissionsOutput {
	return o
}

func (o RequiredPermissionsOutput) ToRequiredPermissionsOutputWithContext(ctx context.Context) RequiredPermissionsOutput {
	return o
}

func (o RequiredPermissionsOutput) ToRequiredPermissionsPtrOutput() RequiredPermissionsPtrOutput {
	return o.ToRequiredPermissionsPtrOutputWithContext(context.Background())
}

func (o RequiredPermissionsOutput) ToRequiredPermissionsPtrOutputWithContext(ctx context.Context) RequiredPermissionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequiredPermissions) *RequiredPermissions {
		return &v
	}).(RequiredPermissionsPtrOutput)
}

func (o RequiredPermissionsOutput) Action() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissions) *bool { return v.Action }).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsOutput) Delete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissions) *bool { return v.Delete }).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsOutput) Read() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissions) *bool { return v.Read }).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsOutput) Write() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissions) *bool { return v.Write }).(pulumi.BoolPtrOutput)
}

type RequiredPermissionsPtrOutput struct{ *pulumi.OutputState }

func (RequiredPermissionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequiredPermissions)(nil)).Elem()
}

func (o RequiredPermissionsPtrOutput) ToRequiredPermissionsPtrOutput() RequiredPermissionsPtrOutput {
	return o
}

func (o RequiredPermissionsPtrOutput) ToRequiredPermissionsPtrOutputWithContext(ctx context.Context) RequiredPermissionsPtrOutput {
	return o
}

func (o RequiredPermissionsPtrOutput) Elem() RequiredPermissionsOutput {
	return o.ApplyT(func(v *RequiredPermissions) RequiredPermissions {
		if v != nil {
			return *v
		}
		var ret RequiredPermissions
		return ret
	}).(RequiredPermissionsOutput)
}

func (o RequiredPermissionsPtrOutput) Action() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissions) *bool {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsPtrOutput) Delete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissions) *bool {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsPtrOutput) Read() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissions) *bool {
		if v == nil {
			return nil
		}
		return v.Read
	}).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsPtrOutput) Write() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissions) *bool {
		if v == nil {
			return nil
		}
		return v.Write
	}).(pulumi.BoolPtrOutput)
}

type RequiredPermissionsResponse struct {
	Action *bool `pulumi:"action"`
	Delete *bool `pulumi:"delete"`
	Read   *bool `pulumi:"read"`
	Write  *bool `pulumi:"write"`
}

type RequiredPermissionsResponseOutput struct{ *pulumi.OutputState }

func (RequiredPermissionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredPermissionsResponse)(nil)).Elem()
}

func (o RequiredPermissionsResponseOutput) ToRequiredPermissionsResponseOutput() RequiredPermissionsResponseOutput {
	return o
}

func (o RequiredPermissionsResponseOutput) ToRequiredPermissionsResponseOutputWithContext(ctx context.Context) RequiredPermissionsResponseOutput {
	return o
}

func (o RequiredPermissionsResponseOutput) Action() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissionsResponse) *bool { return v.Action }).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsResponseOutput) Delete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissionsResponse) *bool { return v.Delete }).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsResponseOutput) Read() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissionsResponse) *bool { return v.Read }).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsResponseOutput) Write() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RequiredPermissionsResponse) *bool { return v.Write }).(pulumi.BoolPtrOutput)
}

type RequiredPermissionsResponsePtrOutput struct{ *pulumi.OutputState }

func (RequiredPermissionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequiredPermissionsResponse)(nil)).Elem()
}

func (o RequiredPermissionsResponsePtrOutput) ToRequiredPermissionsResponsePtrOutput() RequiredPermissionsResponsePtrOutput {
	return o
}

func (o RequiredPermissionsResponsePtrOutput) ToRequiredPermissionsResponsePtrOutputWithContext(ctx context.Context) RequiredPermissionsResponsePtrOutput {
	return o
}

func (o RequiredPermissionsResponsePtrOutput) Elem() RequiredPermissionsResponseOutput {
	return o.ApplyT(func(v *RequiredPermissionsResponse) RequiredPermissionsResponse {
		if v != nil {
			return *v
		}
		var ret RequiredPermissionsResponse
		return ret
	}).(RequiredPermissionsResponseOutput)
}

func (o RequiredPermissionsResponsePtrOutput) Action() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsResponsePtrOutput) Delete() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Delete
	}).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsResponsePtrOutput) Read() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Read
	}).(pulumi.BoolPtrOutput)
}

func (o RequiredPermissionsResponsePtrOutput) Write() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RequiredPermissionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Write
	}).(pulumi.BoolPtrOutput)
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

type TeamInformationResponse struct {
	Description         string `pulumi:"description"`
	Name                string `pulumi:"name"`
	PrimaryChannelUrl   string `pulumi:"primaryChannelUrl"`
	TeamCreationTimeUtc string `pulumi:"teamCreationTimeUtc"`
	TeamId              string `pulumi:"teamId"`
}

type TeamInformationResponseOutput struct{ *pulumi.OutputState }

func (TeamInformationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamInformationResponse)(nil)).Elem()
}

func (o TeamInformationResponseOutput) ToTeamInformationResponseOutput() TeamInformationResponseOutput {
	return o
}

func (o TeamInformationResponseOutput) ToTeamInformationResponseOutputWithContext(ctx context.Context) TeamInformationResponseOutput {
	return o
}

func (o TeamInformationResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v TeamInformationResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o TeamInformationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v TeamInformationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o TeamInformationResponseOutput) PrimaryChannelUrl() pulumi.StringOutput {
	return o.ApplyT(func(v TeamInformationResponse) string { return v.PrimaryChannelUrl }).(pulumi.StringOutput)
}

func (o TeamInformationResponseOutput) TeamCreationTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v TeamInformationResponse) string { return v.TeamCreationTimeUtc }).(pulumi.StringOutput)
}

func (o TeamInformationResponseOutput) TeamId() pulumi.StringOutput {
	return o.ApplyT(func(v TeamInformationResponse) string { return v.TeamId }).(pulumi.StringOutput)
}

type TeamInformationResponsePtrOutput struct{ *pulumi.OutputState }

func (TeamInformationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamInformationResponse)(nil)).Elem()
}

func (o TeamInformationResponsePtrOutput) ToTeamInformationResponsePtrOutput() TeamInformationResponsePtrOutput {
	return o
}

func (o TeamInformationResponsePtrOutput) ToTeamInformationResponsePtrOutputWithContext(ctx context.Context) TeamInformationResponsePtrOutput {
	return o
}

func (o TeamInformationResponsePtrOutput) Elem() TeamInformationResponseOutput {
	return o.ApplyT(func(v *TeamInformationResponse) TeamInformationResponse {
		if v != nil {
			return *v
		}
		var ret TeamInformationResponse
		return ret
	}).(TeamInformationResponseOutput)
}

func (o TeamInformationResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Description
	}).(pulumi.StringPtrOutput)
}

func (o TeamInformationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TeamInformationResponsePtrOutput) PrimaryChannelUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrimaryChannelUrl
	}).(pulumi.StringPtrOutput)
}

func (o TeamInformationResponsePtrOutput) TeamCreationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TeamCreationTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o TeamInformationResponsePtrOutput) TeamId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TeamInformationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TeamId
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

type TimelineAggregationResponseOutput struct{ *pulumi.OutputState }

func (TimelineAggregationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineAggregationResponse)(nil)).Elem()
}

func (o TimelineAggregationResponseOutput) ToTimelineAggregationResponseOutput() TimelineAggregationResponseOutput {
	return o
}

func (o TimelineAggregationResponseOutput) ToTimelineAggregationResponseOutputWithContext(ctx context.Context) TimelineAggregationResponseOutput {
	return o
}

func (o TimelineAggregationResponseOutput) Count() pulumi.IntOutput {
	return o.ApplyT(func(v TimelineAggregationResponse) int { return v.Count }).(pulumi.IntOutput)
}

func (o TimelineAggregationResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TimelineAggregationResponse) string { return v.Kind }).(pulumi.StringOutput)
}

type TimelineAggregationResponseArrayOutput struct{ *pulumi.OutputState }

func (TimelineAggregationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimelineAggregationResponse)(nil)).Elem()
}

func (o TimelineAggregationResponseArrayOutput) ToTimelineAggregationResponseArrayOutput() TimelineAggregationResponseArrayOutput {
	return o
}

func (o TimelineAggregationResponseArrayOutput) ToTimelineAggregationResponseArrayOutputWithContext(ctx context.Context) TimelineAggregationResponseArrayOutput {
	return o
}

func (o TimelineAggregationResponseArrayOutput) Index(i pulumi.IntInput) TimelineAggregationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimelineAggregationResponse {
		return vs[0].([]TimelineAggregationResponse)[vs[1].(int)]
	}).(TimelineAggregationResponseOutput)
}

type TimelineErrorResponse struct {
	ErrorMessage string  `pulumi:"errorMessage"`
	Kind         string  `pulumi:"kind"`
	QueryId      *string `pulumi:"queryId"`
}

type TimelineErrorResponseOutput struct{ *pulumi.OutputState }

func (TimelineErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineErrorResponse)(nil)).Elem()
}

func (o TimelineErrorResponseOutput) ToTimelineErrorResponseOutput() TimelineErrorResponseOutput {
	return o
}

func (o TimelineErrorResponseOutput) ToTimelineErrorResponseOutputWithContext(ctx context.Context) TimelineErrorResponseOutput {
	return o
}

func (o TimelineErrorResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v TimelineErrorResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o TimelineErrorResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v TimelineErrorResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o TimelineErrorResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TimelineErrorResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

type TimelineErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (TimelineErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimelineErrorResponse)(nil)).Elem()
}

func (o TimelineErrorResponseArrayOutput) ToTimelineErrorResponseArrayOutput() TimelineErrorResponseArrayOutput {
	return o
}

func (o TimelineErrorResponseArrayOutput) ToTimelineErrorResponseArrayOutputWithContext(ctx context.Context) TimelineErrorResponseArrayOutput {
	return o
}

func (o TimelineErrorResponseArrayOutput) Index(i pulumi.IntInput) TimelineErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TimelineErrorResponse {
		return vs[0].([]TimelineErrorResponse)[vs[1].(int)]
	}).(TimelineErrorResponseOutput)
}

type TimelineResultsMetadataResponse struct {
	Aggregations []TimelineAggregationResponse `pulumi:"aggregations"`
	Errors       []TimelineErrorResponse       `pulumi:"errors"`
	TotalCount   int                           `pulumi:"totalCount"`
}

type TimelineResultsMetadataResponseOutput struct{ *pulumi.OutputState }

func (TimelineResultsMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineResultsMetadataResponse)(nil)).Elem()
}

func (o TimelineResultsMetadataResponseOutput) ToTimelineResultsMetadataResponseOutput() TimelineResultsMetadataResponseOutput {
	return o
}

func (o TimelineResultsMetadataResponseOutput) ToTimelineResultsMetadataResponseOutputWithContext(ctx context.Context) TimelineResultsMetadataResponseOutput {
	return o
}

func (o TimelineResultsMetadataResponseOutput) Aggregations() TimelineAggregationResponseArrayOutput {
	return o.ApplyT(func(v TimelineResultsMetadataResponse) []TimelineAggregationResponse { return v.Aggregations }).(TimelineAggregationResponseArrayOutput)
}

func (o TimelineResultsMetadataResponseOutput) Errors() TimelineErrorResponseArrayOutput {
	return o.ApplyT(func(v TimelineResultsMetadataResponse) []TimelineErrorResponse { return v.Errors }).(TimelineErrorResponseArrayOutput)
}

func (o TimelineResultsMetadataResponseOutput) TotalCount() pulumi.IntOutput {
	return o.ApplyT(func(v TimelineResultsMetadataResponse) int { return v.TotalCount }).(pulumi.IntOutput)
}

type TimelineResultsMetadataResponsePtrOutput struct{ *pulumi.OutputState }

func (TimelineResultsMetadataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimelineResultsMetadataResponse)(nil)).Elem()
}

func (o TimelineResultsMetadataResponsePtrOutput) ToTimelineResultsMetadataResponsePtrOutput() TimelineResultsMetadataResponsePtrOutput {
	return o
}

func (o TimelineResultsMetadataResponsePtrOutput) ToTimelineResultsMetadataResponsePtrOutputWithContext(ctx context.Context) TimelineResultsMetadataResponsePtrOutput {
	return o
}

func (o TimelineResultsMetadataResponsePtrOutput) Elem() TimelineResultsMetadataResponseOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) TimelineResultsMetadataResponse {
		if v != nil {
			return *v
		}
		var ret TimelineResultsMetadataResponse
		return ret
	}).(TimelineResultsMetadataResponseOutput)
}

func (o TimelineResultsMetadataResponsePtrOutput) Aggregations() TimelineAggregationResponseArrayOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) []TimelineAggregationResponse {
		if v == nil {
			return nil
		}
		return v.Aggregations
	}).(TimelineAggregationResponseArrayOutput)
}

func (o TimelineResultsMetadataResponsePtrOutput) Errors() TimelineErrorResponseArrayOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) []TimelineErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(TimelineErrorResponseArrayOutput)
}

func (o TimelineResultsMetadataResponsePtrOutput) TotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TimelineResultsMetadataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TotalCount
	}).(pulumi.IntPtrOutput)
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

type Webhook struct {
	RotateWebhookSecret     *bool   `pulumi:"rotateWebhookSecret"`
	WebhookId               *string `pulumi:"webhookId"`
	WebhookSecretUpdateTime *string `pulumi:"webhookSecretUpdateTime"`
	WebhookUrl              *string `pulumi:"webhookUrl"`
}





type WebhookInput interface {
	pulumi.Input

	ToWebhookOutput() WebhookOutput
	ToWebhookOutputWithContext(context.Context) WebhookOutput
}

type WebhookArgs struct {
	RotateWebhookSecret     pulumi.BoolPtrInput   `pulumi:"rotateWebhookSecret"`
	WebhookId               pulumi.StringPtrInput `pulumi:"webhookId"`
	WebhookSecretUpdateTime pulumi.StringPtrInput `pulumi:"webhookSecretUpdateTime"`
	WebhookUrl              pulumi.StringPtrInput `pulumi:"webhookUrl"`
}

func (WebhookArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil)).Elem()
}

func (i WebhookArgs) ToWebhookOutput() WebhookOutput {
	return i.ToWebhookOutputWithContext(context.Background())
}

func (i WebhookArgs) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput)
}

func (i WebhookArgs) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i WebhookArgs) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookOutput).ToWebhookPtrOutputWithContext(ctx)
}









type WebhookPtrInput interface {
	pulumi.Input

	ToWebhookPtrOutput() WebhookPtrOutput
	ToWebhookPtrOutputWithContext(context.Context) WebhookPtrOutput
}

type webhookPtrType WebhookArgs

func WebhookPtr(v *WebhookArgs) WebhookPtrInput {
	return (*webhookPtrType)(v)
}

func (*webhookPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (i *webhookPtrType) ToWebhookPtrOutput() WebhookPtrOutput {
	return i.ToWebhookPtrOutputWithContext(context.Background())
}

func (i *webhookPtrType) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebhookPtrOutput)
}

type WebhookOutput struct{ *pulumi.OutputState }

func (WebhookOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Webhook)(nil)).Elem()
}

func (o WebhookOutput) ToWebhookOutput() WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookOutputWithContext(ctx context.Context) WebhookOutput {
	return o
}

func (o WebhookOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o.ToWebhookPtrOutputWithContext(context.Background())
}

func (o WebhookOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Webhook) *Webhook {
		return &v
	}).(WebhookPtrOutput)
}

func (o WebhookOutput) RotateWebhookSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Webhook) *bool { return v.RotateWebhookSecret }).(pulumi.BoolPtrOutput)
}

func (o WebhookOutput) WebhookId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Webhook) *string { return v.WebhookId }).(pulumi.StringPtrOutput)
}

func (o WebhookOutput) WebhookSecretUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Webhook) *string { return v.WebhookSecretUpdateTime }).(pulumi.StringPtrOutput)
}

func (o WebhookOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Webhook) *string { return v.WebhookUrl }).(pulumi.StringPtrOutput)
}

type WebhookPtrOutput struct{ *pulumi.OutputState }

func (WebhookPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Webhook)(nil)).Elem()
}

func (o WebhookPtrOutput) ToWebhookPtrOutput() WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) ToWebhookPtrOutputWithContext(ctx context.Context) WebhookPtrOutput {
	return o
}

func (o WebhookPtrOutput) Elem() WebhookOutput {
	return o.ApplyT(func(v *Webhook) Webhook {
		if v != nil {
			return *v
		}
		var ret Webhook
		return ret
	}).(WebhookOutput)
}

func (o WebhookPtrOutput) RotateWebhookSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Webhook) *bool {
		if v == nil {
			return nil
		}
		return v.RotateWebhookSecret
	}).(pulumi.BoolPtrOutput)
}

func (o WebhookPtrOutput) WebhookId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) *string {
		if v == nil {
			return nil
		}
		return v.WebhookId
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPtrOutput) WebhookSecretUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) *string {
		if v == nil {
			return nil
		}
		return v.WebhookSecretUpdateTime
	}).(pulumi.StringPtrOutput)
}

func (o WebhookPtrOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Webhook) *string {
		if v == nil {
			return nil
		}
		return v.WebhookUrl
	}).(pulumi.StringPtrOutput)
}

type WebhookResponse struct {
	RotateWebhookSecret     *bool   `pulumi:"rotateWebhookSecret"`
	WebhookId               *string `pulumi:"webhookId"`
	WebhookSecretUpdateTime *string `pulumi:"webhookSecretUpdateTime"`
	WebhookUrl              *string `pulumi:"webhookUrl"`
}

type WebhookResponseOutput struct{ *pulumi.OutputState }

func (WebhookResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebhookResponse)(nil)).Elem()
}

func (o WebhookResponseOutput) ToWebhookResponseOutput() WebhookResponseOutput {
	return o
}

func (o WebhookResponseOutput) ToWebhookResponseOutputWithContext(ctx context.Context) WebhookResponseOutput {
	return o
}

func (o WebhookResponseOutput) RotateWebhookSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v WebhookResponse) *bool { return v.RotateWebhookSecret }).(pulumi.BoolPtrOutput)
}

func (o WebhookResponseOutput) WebhookId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookResponse) *string { return v.WebhookId }).(pulumi.StringPtrOutput)
}

func (o WebhookResponseOutput) WebhookSecretUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookResponse) *string { return v.WebhookSecretUpdateTime }).(pulumi.StringPtrOutput)
}

func (o WebhookResponseOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebhookResponse) *string { return v.WebhookUrl }).(pulumi.StringPtrOutput)
}

type WebhookResponsePtrOutput struct{ *pulumi.OutputState }

func (WebhookResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebhookResponse)(nil)).Elem()
}

func (o WebhookResponsePtrOutput) ToWebhookResponsePtrOutput() WebhookResponsePtrOutput {
	return o
}

func (o WebhookResponsePtrOutput) ToWebhookResponsePtrOutputWithContext(ctx context.Context) WebhookResponsePtrOutput {
	return o
}

func (o WebhookResponsePtrOutput) Elem() WebhookResponseOutput {
	return o.ApplyT(func(v *WebhookResponse) WebhookResponse {
		if v != nil {
			return *v
		}
		var ret WebhookResponse
		return ret
	}).(WebhookResponseOutput)
}

func (o WebhookResponsePtrOutput) RotateWebhookSecret() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebhookResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RotateWebhookSecret
	}).(pulumi.BoolPtrOutput)
}

func (o WebhookResponsePtrOutput) WebhookId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebhookId
	}).(pulumi.StringPtrOutput)
}

func (o WebhookResponsePtrOutput) WebhookSecretUpdateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebhookSecretUpdateTime
	}).(pulumi.StringPtrOutput)
}

func (o WebhookResponsePtrOutput) WebhookUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebhookResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebhookUrl
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ActivityEntityQueriesPropertiesQueryDefinitionsOutput{})
	pulumi.RegisterOutputType(ActivityEntityQueriesPropertiesQueryDefinitionsPtrOutput{})
	pulumi.RegisterOutputType(ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput{})
	pulumi.RegisterOutputType(ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput{})
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
	pulumi.RegisterOutputType(AvailabilityOutput{})
	pulumi.RegisterOutputType(AvailabilityPtrOutput{})
	pulumi.RegisterOutputType(AvailabilityResponseOutput{})
	pulumi.RegisterOutputType(AvailabilityResponsePtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(AwsS3DataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(AwsS3DataConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(AwsS3DataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(AwsS3DataConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(AzureDevOpsResourceInfoOutput{})
	pulumi.RegisterOutputType(AzureDevOpsResourceInfoPtrOutput{})
	pulumi.RegisterOutputType(AzureDevOpsResourceInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureDevOpsResourceInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(BookmarkEntityMappingsOutput{})
	pulumi.RegisterOutputType(BookmarkEntityMappingsArrayOutput{})
	pulumi.RegisterOutputType(BookmarkEntityMappingsResponseOutput{})
	pulumi.RegisterOutputType(BookmarkEntityMappingsResponseArrayOutput{})
	pulumi.RegisterOutputType(ClientInfoResponseOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingAuthPropertiesOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingAuthPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingAuthPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingAuthPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingConfigPropertiesOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingConfigPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingConfigPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingConfigPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingPagingPropertiesOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingPagingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingPagingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingPagingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingRequestPropertiesOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingRequestPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingRequestPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingRequestPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingResponsePropertiesOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingResponsePropertiesResponseOutput{})
	pulumi.RegisterOutputType(CodelessConnectorPollingResponsePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesPtrOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesConnectivityCriteriaOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesConnectivityCriteriaArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesDataTypesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesDataTypesArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesGraphQueriesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesGraphQueriesArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesInstructionStepsOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesInstructionStepsArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseDataTypesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesSampleQueriesOutput{})
	pulumi.RegisterOutputType(CodelessUiConnectorConfigPropertiesSampleQueriesArrayOutput{})
	pulumi.RegisterOutputType(ContentPathMapOutput{})
	pulumi.RegisterOutputType(ContentPathMapArrayOutput{})
	pulumi.RegisterOutputType(ContentPathMapResponseOutput{})
	pulumi.RegisterOutputType(ContentPathMapResponseArrayOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonPtrOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonResponseOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentOutput{})
	pulumi.RegisterOutputType(DeploymentPtrOutput{})
	pulumi.RegisterOutputType(DeploymentInfoOutput{})
	pulumi.RegisterOutputType(DeploymentInfoPtrOutput{})
	pulumi.RegisterOutputType(DeploymentInfoResponseOutput{})
	pulumi.RegisterOutputType(DeploymentInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(DeploymentResponseOutput{})
	pulumi.RegisterOutputType(DeploymentResponsePtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(EntityFieldMappingOutput{})
	pulumi.RegisterOutputType(EntityFieldMappingArrayOutput{})
	pulumi.RegisterOutputType(EntityFieldMappingResponseOutput{})
	pulumi.RegisterOutputType(EntityFieldMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseArrayOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalPtrOutput{})
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
	pulumi.RegisterOutputType(FusionScenarioExclusionPatternOutput{})
	pulumi.RegisterOutputType(FusionScenarioExclusionPatternArrayOutput{})
	pulumi.RegisterOutputType(FusionScenarioExclusionPatternResponseOutput{})
	pulumi.RegisterOutputType(FusionScenarioExclusionPatternResponseArrayOutput{})
	pulumi.RegisterOutputType(FusionSourceSettingsOutput{})
	pulumi.RegisterOutputType(FusionSourceSettingsArrayOutput{})
	pulumi.RegisterOutputType(FusionSourceSettingsResponseOutput{})
	pulumi.RegisterOutputType(FusionSourceSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(FusionSourceSubTypeSettingOutput{})
	pulumi.RegisterOutputType(FusionSourceSubTypeSettingArrayOutput{})
	pulumi.RegisterOutputType(FusionSourceSubTypeSettingResponseOutput{})
	pulumi.RegisterOutputType(FusionSourceSubTypeSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(FusionSubTypeSeverityFilterOutput{})
	pulumi.RegisterOutputType(FusionSubTypeSeverityFilterResponseOutput{})
	pulumi.RegisterOutputType(FusionSubTypeSeverityFiltersItemOutput{})
	pulumi.RegisterOutputType(FusionSubTypeSeverityFiltersItemArrayOutput{})
	pulumi.RegisterOutputType(FusionSubTypeSeverityFiltersItemResponseOutput{})
	pulumi.RegisterOutputType(FusionSubTypeSeverityFiltersItemResponseArrayOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorKindResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorKindResponseArrayOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubResourceInfoOutput{})
	pulumi.RegisterOutputType(GitHubResourceInfoPtrOutput{})
	pulumi.RegisterOutputType(GitHubResourceInfoResponseOutput{})
	pulumi.RegisterOutputType(GitHubResourceInfoResponsePtrOutput{})
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
	pulumi.RegisterOutputType(InsightsTableResultResponseOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponsePtrOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseArrayOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseColumnsOutput{})
	pulumi.RegisterOutputType(InsightsTableResultResponseColumnsArrayOutput{})
	pulumi.RegisterOutputType(InstructionStepsInstructionsOutput{})
	pulumi.RegisterOutputType(InstructionStepsInstructionsArrayOutput{})
	pulumi.RegisterOutputType(InstructionStepsResponseInstructionsOutput{})
	pulumi.RegisterOutputType(InstructionStepsResponseInstructionsArrayOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseOutput{})
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
	pulumi.RegisterOutputType(MetadataAuthorOutput{})
	pulumi.RegisterOutputType(MetadataAuthorPtrOutput{})
	pulumi.RegisterOutputType(MetadataAuthorResponseOutput{})
	pulumi.RegisterOutputType(MetadataAuthorResponsePtrOutput{})
	pulumi.RegisterOutputType(MetadataCategoriesOutput{})
	pulumi.RegisterOutputType(MetadataCategoriesPtrOutput{})
	pulumi.RegisterOutputType(MetadataCategoriesResponseOutput{})
	pulumi.RegisterOutputType(MetadataCategoriesResponsePtrOutput{})
	pulumi.RegisterOutputType(MetadataDependenciesOutput{})
	pulumi.RegisterOutputType(MetadataDependenciesPtrOutput{})
	pulumi.RegisterOutputType(MetadataDependenciesArrayOutput{})
	pulumi.RegisterOutputType(MetadataDependenciesResponseOutput{})
	pulumi.RegisterOutputType(MetadataDependenciesResponsePtrOutput{})
	pulumi.RegisterOutputType(MetadataDependenciesResponseArrayOutput{})
	pulumi.RegisterOutputType(MetadataSourceOutput{})
	pulumi.RegisterOutputType(MetadataSourcePtrOutput{})
	pulumi.RegisterOutputType(MetadataSourceResponseOutput{})
	pulumi.RegisterOutputType(MetadataSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(MetadataSupportOutput{})
	pulumi.RegisterOutputType(MetadataSupportPtrOutput{})
	pulumi.RegisterOutputType(MetadataSupportResponseOutput{})
	pulumi.RegisterOutputType(MetadataSupportResponsePtrOutput{})
	pulumi.RegisterOutputType(Office365ProjectConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(Office365ProjectConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(Office365ProjectConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(Office365ProjectConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesExchangeOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseExchangeOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseSharePointOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesResponseTeamsOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesSharePointOutput{})
	pulumi.RegisterOutputType(OfficeDataConnectorDataTypesTeamsOutput{})
	pulumi.RegisterOutputType(OfficePowerBIConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(OfficePowerBIConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(OfficePowerBIConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(OfficePowerBIConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(PermissionsOutput{})
	pulumi.RegisterOutputType(PermissionsPtrOutput{})
	pulumi.RegisterOutputType(PermissionsCustomsOutput{})
	pulumi.RegisterOutputType(PermissionsCustomsArrayOutput{})
	pulumi.RegisterOutputType(PermissionsResourceProviderOutput{})
	pulumi.RegisterOutputType(PermissionsResourceProviderArrayOutput{})
	pulumi.RegisterOutputType(PermissionsResponseOutput{})
	pulumi.RegisterOutputType(PermissionsResponsePtrOutput{})
	pulumi.RegisterOutputType(PermissionsResponseCustomsOutput{})
	pulumi.RegisterOutputType(PermissionsResponseCustomsArrayOutput{})
	pulumi.RegisterOutputType(PermissionsResponseResourceProviderOutput{})
	pulumi.RegisterOutputType(PermissionsResponseResourceProviderArrayOutput{})
	pulumi.RegisterOutputType(RepoResponseOutput{})
	pulumi.RegisterOutputType(RepoResponseArrayOutput{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryResourceInfoOutput{})
	pulumi.RegisterOutputType(RepositoryResourceInfoPtrOutput{})
	pulumi.RegisterOutputType(RepositoryResourceInfoResponseOutput{})
	pulumi.RegisterOutputType(RepositoryResourceInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(RepositoryResponseOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsPtrOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsResponseOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceArrayOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceResponseOutput{})
	pulumi.RegisterOutputType(SecurityMLAnalyticsSettingsDataSourceResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesIndicatorsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseIndicatorsOutput{})
	pulumi.RegisterOutputType(TeamInformationResponseOutput{})
	pulumi.RegisterOutputType(TeamInformationResponsePtrOutput{})
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
	pulumi.RegisterOutputType(TimelineAggregationResponseOutput{})
	pulumi.RegisterOutputType(TimelineAggregationResponseArrayOutput{})
	pulumi.RegisterOutputType(TimelineErrorResponseOutput{})
	pulumi.RegisterOutputType(TimelineErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(TimelineResultsMetadataResponseOutput{})
	pulumi.RegisterOutputType(TimelineResultsMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(UserInfoOutput{})
	pulumi.RegisterOutputType(UserInfoPtrOutput{})
	pulumi.RegisterOutputType(UserInfoResponseOutput{})
	pulumi.RegisterOutputType(UserInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoPtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponseOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(WebhookOutput{})
	pulumi.RegisterOutputType(WebhookPtrOutput{})
	pulumi.RegisterOutputType(WebhookResponseOutput{})
	pulumi.RegisterOutputType(WebhookResponsePtrOutput{})
}
