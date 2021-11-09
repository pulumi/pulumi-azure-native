


package v20210301preview

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





type ActivityEntityQueriesPropertiesResponseQueryDefinitionsInput interface {
	pulumi.Input

	ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput
	ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutputWithContext(context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput
}

type ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs struct {
	Query pulumi.StringPtrInput `pulumi:"query"`
}

func (ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityEntityQueriesPropertiesResponseQueryDefinitions)(nil)).Elem()
}

func (i ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput {
	return i.ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutputWithContext(context.Background())
}

func (i ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput)
}

func (i ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return i.ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(context.Background())
}

func (i ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput).ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(ctx)
}









type ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrInput interface {
	pulumi.Input

	ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput
	ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput
}

type activityEntityQueriesPropertiesResponseQueryDefinitionsPtrType ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs

func ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtr(v *ActivityEntityQueriesPropertiesResponseQueryDefinitionsArgs) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrInput {
	return (*activityEntityQueriesPropertiesResponseQueryDefinitionsPtrType)(v)
}

func (*activityEntityQueriesPropertiesResponseQueryDefinitionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivityEntityQueriesPropertiesResponseQueryDefinitions)(nil)).Elem()
}

func (i *activityEntityQueriesPropertiesResponseQueryDefinitionsPtrType) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return i.ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(context.Background())
}

func (i *activityEntityQueriesPropertiesResponseQueryDefinitionsPtrType) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput)
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

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput() ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o.ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(context.Background())
}

func (o ActivityEntityQueriesPropertiesResponseQueryDefinitionsOutput) ToActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutputWithContext(ctx context.Context) ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActivityEntityQueriesPropertiesResponseQueryDefinitions) *ActivityEntityQueriesPropertiesResponseQueryDefinitions {
		return &v
	}).(ActivityEntityQueriesPropertiesResponseQueryDefinitionsPtrOutput)
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





type AlertDetailsOverrideResponseInput interface {
	pulumi.Input

	ToAlertDetailsOverrideResponseOutput() AlertDetailsOverrideResponseOutput
	ToAlertDetailsOverrideResponseOutputWithContext(context.Context) AlertDetailsOverrideResponseOutput
}

type AlertDetailsOverrideResponseArgs struct {
	AlertDescriptionFormat  pulumi.StringPtrInput `pulumi:"alertDescriptionFormat"`
	AlertDisplayNameFormat  pulumi.StringPtrInput `pulumi:"alertDisplayNameFormat"`
	AlertSeverityColumnName pulumi.StringPtrInput `pulumi:"alertSeverityColumnName"`
	AlertTacticsColumnName  pulumi.StringPtrInput `pulumi:"alertTacticsColumnName"`
}

func (AlertDetailsOverrideResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertDetailsOverrideResponse)(nil)).Elem()
}

func (i AlertDetailsOverrideResponseArgs) ToAlertDetailsOverrideResponseOutput() AlertDetailsOverrideResponseOutput {
	return i.ToAlertDetailsOverrideResponseOutputWithContext(context.Background())
}

func (i AlertDetailsOverrideResponseArgs) ToAlertDetailsOverrideResponseOutputWithContext(ctx context.Context) AlertDetailsOverrideResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDetailsOverrideResponseOutput)
}

func (i AlertDetailsOverrideResponseArgs) ToAlertDetailsOverrideResponsePtrOutput() AlertDetailsOverrideResponsePtrOutput {
	return i.ToAlertDetailsOverrideResponsePtrOutputWithContext(context.Background())
}

func (i AlertDetailsOverrideResponseArgs) ToAlertDetailsOverrideResponsePtrOutputWithContext(ctx context.Context) AlertDetailsOverrideResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDetailsOverrideResponseOutput).ToAlertDetailsOverrideResponsePtrOutputWithContext(ctx)
}









type AlertDetailsOverrideResponsePtrInput interface {
	pulumi.Input

	ToAlertDetailsOverrideResponsePtrOutput() AlertDetailsOverrideResponsePtrOutput
	ToAlertDetailsOverrideResponsePtrOutputWithContext(context.Context) AlertDetailsOverrideResponsePtrOutput
}

type alertDetailsOverrideResponsePtrType AlertDetailsOverrideResponseArgs

func AlertDetailsOverrideResponsePtr(v *AlertDetailsOverrideResponseArgs) AlertDetailsOverrideResponsePtrInput {
	return (*alertDetailsOverrideResponsePtrType)(v)
}

func (*alertDetailsOverrideResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertDetailsOverrideResponse)(nil)).Elem()
}

func (i *alertDetailsOverrideResponsePtrType) ToAlertDetailsOverrideResponsePtrOutput() AlertDetailsOverrideResponsePtrOutput {
	return i.ToAlertDetailsOverrideResponsePtrOutputWithContext(context.Background())
}

func (i *alertDetailsOverrideResponsePtrType) ToAlertDetailsOverrideResponsePtrOutputWithContext(ctx context.Context) AlertDetailsOverrideResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertDetailsOverrideResponsePtrOutput)
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

func (o AlertDetailsOverrideResponseOutput) ToAlertDetailsOverrideResponsePtrOutput() AlertDetailsOverrideResponsePtrOutput {
	return o.ToAlertDetailsOverrideResponsePtrOutputWithContext(context.Background())
}

func (o AlertDetailsOverrideResponseOutput) ToAlertDetailsOverrideResponsePtrOutputWithContext(ctx context.Context) AlertDetailsOverrideResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertDetailsOverrideResponse) *AlertDetailsOverrideResponse {
		return &v
	}).(AlertDetailsOverrideResponsePtrOutput)
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





type AlertsDataTypeOfDataConnectorResponseInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorResponseOutput() AlertsDataTypeOfDataConnectorResponseOutput
	ToAlertsDataTypeOfDataConnectorResponseOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorResponseOutput
}

type AlertsDataTypeOfDataConnectorResponseArgs struct {
	Alerts DataConnectorDataTypeCommonResponseInput `pulumi:"alerts"`
}

func (AlertsDataTypeOfDataConnectorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponse)(nil)).Elem()
}

func (i AlertsDataTypeOfDataConnectorResponseArgs) ToAlertsDataTypeOfDataConnectorResponseOutput() AlertsDataTypeOfDataConnectorResponseOutput {
	return i.ToAlertsDataTypeOfDataConnectorResponseOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorResponseArgs) ToAlertsDataTypeOfDataConnectorResponseOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorResponseOutput)
}

func (i AlertsDataTypeOfDataConnectorResponseArgs) ToAlertsDataTypeOfDataConnectorResponsePtrOutput() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorResponseArgs) ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorResponseOutput).ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(ctx)
}









type AlertsDataTypeOfDataConnectorResponsePtrInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorResponsePtrOutput() AlertsDataTypeOfDataConnectorResponsePtrOutput
	ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorResponsePtrOutput
}

type alertsDataTypeOfDataConnectorResponsePtrType AlertsDataTypeOfDataConnectorResponseArgs

func AlertsDataTypeOfDataConnectorResponsePtr(v *AlertsDataTypeOfDataConnectorResponseArgs) AlertsDataTypeOfDataConnectorResponsePtrInput {
	return (*alertsDataTypeOfDataConnectorResponsePtrType)(v)
}

func (*alertsDataTypeOfDataConnectorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnectorResponse)(nil)).Elem()
}

func (i *alertsDataTypeOfDataConnectorResponsePtrType) ToAlertsDataTypeOfDataConnectorResponsePtrOutput() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(context.Background())
}

func (i *alertsDataTypeOfDataConnectorResponsePtrType) ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
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

func (o AlertsDataTypeOfDataConnectorResponseOutput) ToAlertsDataTypeOfDataConnectorResponsePtrOutput() AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(context.Background())
}

func (o AlertsDataTypeOfDataConnectorResponseOutput) ToAlertsDataTypeOfDataConnectorResponsePtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertsDataTypeOfDataConnectorResponse) *AlertsDataTypeOfDataConnectorResponse {
		return &v
	}).(AlertsDataTypeOfDataConnectorResponsePtrOutput)
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





type AvailabilityResponseInput interface {
	pulumi.Input

	ToAvailabilityResponseOutput() AvailabilityResponseOutput
	ToAvailabilityResponseOutputWithContext(context.Context) AvailabilityResponseOutput
}

type AvailabilityResponseArgs struct {
	IsPreview pulumi.BoolPtrInput `pulumi:"isPreview"`
	Status    pulumi.IntPtrInput  `pulumi:"status"`
}

func (AvailabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AvailabilityResponse)(nil)).Elem()
}

func (i AvailabilityResponseArgs) ToAvailabilityResponseOutput() AvailabilityResponseOutput {
	return i.ToAvailabilityResponseOutputWithContext(context.Background())
}

func (i AvailabilityResponseArgs) ToAvailabilityResponseOutputWithContext(ctx context.Context) AvailabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityResponseOutput)
}

func (i AvailabilityResponseArgs) ToAvailabilityResponsePtrOutput() AvailabilityResponsePtrOutput {
	return i.ToAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (i AvailabilityResponseArgs) ToAvailabilityResponsePtrOutputWithContext(ctx context.Context) AvailabilityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityResponseOutput).ToAvailabilityResponsePtrOutputWithContext(ctx)
}









type AvailabilityResponsePtrInput interface {
	pulumi.Input

	ToAvailabilityResponsePtrOutput() AvailabilityResponsePtrOutput
	ToAvailabilityResponsePtrOutputWithContext(context.Context) AvailabilityResponsePtrOutput
}

type availabilityResponsePtrType AvailabilityResponseArgs

func AvailabilityResponsePtr(v *AvailabilityResponseArgs) AvailabilityResponsePtrInput {
	return (*availabilityResponsePtrType)(v)
}

func (*availabilityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AvailabilityResponse)(nil)).Elem()
}

func (i *availabilityResponsePtrType) ToAvailabilityResponsePtrOutput() AvailabilityResponsePtrOutput {
	return i.ToAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (i *availabilityResponsePtrType) ToAvailabilityResponsePtrOutputWithContext(ctx context.Context) AvailabilityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AvailabilityResponsePtrOutput)
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

func (o AvailabilityResponseOutput) ToAvailabilityResponsePtrOutput() AvailabilityResponsePtrOutput {
	return o.ToAvailabilityResponsePtrOutputWithContext(context.Background())
}

func (o AvailabilityResponseOutput) ToAvailabilityResponsePtrOutputWithContext(ctx context.Context) AvailabilityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AvailabilityResponse) *AvailabilityResponse {
		return &v
	}).(AvailabilityResponsePtrOutput)
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

func (o AwsCloudTrailDataConnectorDataTypesOutput) Logs() AwsCloudTrailDataConnectorDataTypesLogsOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypes) AwsCloudTrailDataConnectorDataTypesLogs { return v.Logs }).(AwsCloudTrailDataConnectorDataTypesLogsOutput)
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
		return &v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput)
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

func (o AwsCloudTrailDataConnectorDataTypesLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesLogs) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponse struct {
	Logs AwsCloudTrailDataConnectorDataTypesResponseLogs `pulumi:"logs"`
}





type AwsCloudTrailDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesResponseOutput() AwsCloudTrailDataConnectorDataTypesResponseOutput
	ToAwsCloudTrailDataConnectorDataTypesResponseOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesResponseOutput
}

type AwsCloudTrailDataConnectorDataTypesResponseArgs struct {
	Logs AwsCloudTrailDataConnectorDataTypesResponseLogsInput `pulumi:"logs"`
}

func (AwsCloudTrailDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i AwsCloudTrailDataConnectorDataTypesResponseArgs) ToAwsCloudTrailDataConnectorDataTypesResponseOutput() AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesResponseArgs) ToAwsCloudTrailDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesResponseOutput)
}

func (i AwsCloudTrailDataConnectorDataTypesResponseArgs) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutput() AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesResponseArgs) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesResponseOutput).ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type AwsCloudTrailDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutput() AwsCloudTrailDataConnectorDataTypesResponsePtrOutput
	ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesResponsePtrOutput
}

type awsCloudTrailDataConnectorDataTypesResponsePtrType AwsCloudTrailDataConnectorDataTypesResponseArgs

func AwsCloudTrailDataConnectorDataTypesResponsePtr(v *AwsCloudTrailDataConnectorDataTypesResponseArgs) AwsCloudTrailDataConnectorDataTypesResponsePtrInput {
	return (*awsCloudTrailDataConnectorDataTypesResponsePtrType)(v)
}

func (*awsCloudTrailDataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *awsCloudTrailDataConnectorDataTypesResponsePtrType) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutput() AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *awsCloudTrailDataConnectorDataTypesResponsePtrType) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesResponsePtrOutput)
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

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutput() AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return o.ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) ToAwsCloudTrailDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AwsCloudTrailDataConnectorDataTypesResponse) *AwsCloudTrailDataConnectorDataTypesResponse {
		return &v
	}).(AwsCloudTrailDataConnectorDataTypesResponsePtrOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesResponseOutput) Logs() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesResponse) AwsCloudTrailDataConnectorDataTypesResponseLogs {
		return v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput)
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
		return &v.Logs
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput)
}

type AwsCloudTrailDataConnectorDataTypesResponseLogs struct {
	State string `pulumi:"state"`
}





type AwsCloudTrailDataConnectorDataTypesResponseLogsInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput
	ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsOutput
}

type AwsCloudTrailDataConnectorDataTypesResponseLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (AwsCloudTrailDataConnectorDataTypesResponseLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (i AwsCloudTrailDataConnectorDataTypesResponseLogsArgs) ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesResponseLogsArgs) ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput)
}

func (i AwsCloudTrailDataConnectorDataTypesResponseLogsArgs) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(context.Background())
}

func (i AwsCloudTrailDataConnectorDataTypesResponseLogsArgs) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput).ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(ctx)
}









type AwsCloudTrailDataConnectorDataTypesResponseLogsPtrInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput
	ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput
}

type awsCloudTrailDataConnectorDataTypesResponseLogsPtrType AwsCloudTrailDataConnectorDataTypesResponseLogsArgs

func AwsCloudTrailDataConnectorDataTypesResponseLogsPtr(v *AwsCloudTrailDataConnectorDataTypesResponseLogsArgs) AwsCloudTrailDataConnectorDataTypesResponseLogsPtrInput {
	return (*awsCloudTrailDataConnectorDataTypesResponseLogsPtrType)(v)
}

func (*awsCloudTrailDataConnectorDataTypesResponseLogsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AwsCloudTrailDataConnectorDataTypesResponseLogs)(nil)).Elem()
}

func (i *awsCloudTrailDataConnectorDataTypesResponseLogsPtrType) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return i.ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(context.Background())
}

func (i *awsCloudTrailDataConnectorDataTypesResponseLogsPtrType) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput)
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

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return o.ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(context.Background())
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) ToAwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutputWithContext(ctx context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AwsCloudTrailDataConnectorDataTypesResponseLogs) *AwsCloudTrailDataConnectorDataTypesResponseLogs {
		return &v
	}).(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput)
}

func (o AwsCloudTrailDataConnectorDataTypesResponseLogsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AwsCloudTrailDataConnectorDataTypesResponseLogs) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type ClientInfoResponse struct {
	Email             *string `pulumi:"email"`
	Name              *string `pulumi:"name"`
	ObjectId          *string `pulumi:"objectId"`
	UserPrincipalName *string `pulumi:"userPrincipalName"`
}





type ClientInfoResponseInput interface {
	pulumi.Input

	ToClientInfoResponseOutput() ClientInfoResponseOutput
	ToClientInfoResponseOutputWithContext(context.Context) ClientInfoResponseOutput
}

type ClientInfoResponseArgs struct {
	Email             pulumi.StringPtrInput `pulumi:"email"`
	Name              pulumi.StringPtrInput `pulumi:"name"`
	ObjectId          pulumi.StringPtrInput `pulumi:"objectId"`
	UserPrincipalName pulumi.StringPtrInput `pulumi:"userPrincipalName"`
}

func (ClientInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClientInfoResponse)(nil)).Elem()
}

func (i ClientInfoResponseArgs) ToClientInfoResponseOutput() ClientInfoResponseOutput {
	return i.ToClientInfoResponseOutputWithContext(context.Background())
}

func (i ClientInfoResponseArgs) ToClientInfoResponseOutputWithContext(ctx context.Context) ClientInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientInfoResponseOutput)
}

func (i ClientInfoResponseArgs) ToClientInfoResponsePtrOutput() ClientInfoResponsePtrOutput {
	return i.ToClientInfoResponsePtrOutputWithContext(context.Background())
}

func (i ClientInfoResponseArgs) ToClientInfoResponsePtrOutputWithContext(ctx context.Context) ClientInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientInfoResponseOutput).ToClientInfoResponsePtrOutputWithContext(ctx)
}









type ClientInfoResponsePtrInput interface {
	pulumi.Input

	ToClientInfoResponsePtrOutput() ClientInfoResponsePtrOutput
	ToClientInfoResponsePtrOutputWithContext(context.Context) ClientInfoResponsePtrOutput
}

type clientInfoResponsePtrType ClientInfoResponseArgs

func ClientInfoResponsePtr(v *ClientInfoResponseArgs) ClientInfoResponsePtrInput {
	return (*clientInfoResponsePtrType)(v)
}

func (*clientInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientInfoResponse)(nil)).Elem()
}

func (i *clientInfoResponsePtrType) ToClientInfoResponsePtrOutput() ClientInfoResponsePtrOutput {
	return i.ToClientInfoResponsePtrOutputWithContext(context.Background())
}

func (i *clientInfoResponsePtrType) ToClientInfoResponsePtrOutputWithContext(ctx context.Context) ClientInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClientInfoResponsePtrOutput)
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

func (o ClientInfoResponseOutput) ToClientInfoResponsePtrOutput() ClientInfoResponsePtrOutput {
	return o.ToClientInfoResponsePtrOutputWithContext(context.Background())
}

func (o ClientInfoResponseOutput) ToClientInfoResponsePtrOutputWithContext(ctx context.Context) ClientInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClientInfoResponse) *ClientInfoResponse {
		return &v
	}).(ClientInfoResponsePtrOutput)
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

type ClientInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ClientInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClientInfoResponse)(nil)).Elem()
}

func (o ClientInfoResponsePtrOutput) ToClientInfoResponsePtrOutput() ClientInfoResponsePtrOutput {
	return o
}

func (o ClientInfoResponsePtrOutput) ToClientInfoResponsePtrOutputWithContext(ctx context.Context) ClientInfoResponsePtrOutput {
	return o
}

func (o ClientInfoResponsePtrOutput) Elem() ClientInfoResponseOutput {
	return o.ApplyT(func(v *ClientInfoResponse) ClientInfoResponse {
		if v != nil {
			return *v
		}
		var ret ClientInfoResponse
		return ret
	}).(ClientInfoResponseOutput)
}

func (o ClientInfoResponsePtrOutput) Email() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Email
	}).(pulumi.StringPtrOutput)
}

func (o ClientInfoResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ClientInfoResponsePtrOutput) ObjectId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.ObjectId
	}).(pulumi.StringPtrOutput)
}

func (o ClientInfoResponsePtrOutput) UserPrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClientInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.UserPrincipalName
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





type CodelessUiConnectorConfigPropertiesResponseInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseOutput() CodelessUiConnectorConfigPropertiesResponseOutput
	ToCodelessUiConnectorConfigPropertiesResponseOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseOutput
}

type CodelessUiConnectorConfigPropertiesResponseArgs struct {
	Availability          AvailabilityResponseInput                                                 `pulumi:"availability"`
	ConnectivityCriteria  CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayInput `pulumi:"connectivityCriteria"`
	CustomImage           pulumi.StringPtrInput                                                     `pulumi:"customImage"`
	DataTypes             CodelessUiConnectorConfigPropertiesResponseDataTypesArrayInput            `pulumi:"dataTypes"`
	DescriptionMarkdown   pulumi.StringInput                                                        `pulumi:"descriptionMarkdown"`
	GraphQueries          CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayInput         `pulumi:"graphQueries"`
	GraphQueriesTableName pulumi.StringInput                                                        `pulumi:"graphQueriesTableName"`
	InstructionSteps      CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayInput     `pulumi:"instructionSteps"`
	Permissions           PermissionsResponseInput                                                  `pulumi:"permissions"`
	Publisher             pulumi.StringInput                                                        `pulumi:"publisher"`
	SampleQueries         CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayInput        `pulumi:"sampleQueries"`
	Title                 pulumi.StringInput                                                        `pulumi:"title"`
}

func (CodelessUiConnectorConfigPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponse)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseArgs) ToCodelessUiConnectorConfigPropertiesResponseOutput() CodelessUiConnectorConfigPropertiesResponseOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseArgs) ToCodelessUiConnectorConfigPropertiesResponseOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseOutput)
}

func (i CodelessUiConnectorConfigPropertiesResponseArgs) ToCodelessUiConnectorConfigPropertiesResponsePtrOutput() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseArgs) ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseOutput).ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(ctx)
}









type CodelessUiConnectorConfigPropertiesResponsePtrInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponsePtrOutput() CodelessUiConnectorConfigPropertiesResponsePtrOutput
	ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponsePtrOutput
}

type codelessUiConnectorConfigPropertiesResponsePtrType CodelessUiConnectorConfigPropertiesResponseArgs

func CodelessUiConnectorConfigPropertiesResponsePtr(v *CodelessUiConnectorConfigPropertiesResponseArgs) CodelessUiConnectorConfigPropertiesResponsePtrInput {
	return (*codelessUiConnectorConfigPropertiesResponsePtrType)(v)
}

func (*codelessUiConnectorConfigPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CodelessUiConnectorConfigPropertiesResponse)(nil)).Elem()
}

func (i *codelessUiConnectorConfigPropertiesResponsePtrType) ToCodelessUiConnectorConfigPropertiesResponsePtrOutput() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *codelessUiConnectorConfigPropertiesResponsePtrType) ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
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

func (o CodelessUiConnectorConfigPropertiesResponseOutput) ToCodelessUiConnectorConfigPropertiesResponsePtrOutput() CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o CodelessUiConnectorConfigPropertiesResponseOutput) ToCodelessUiConnectorConfigPropertiesResponsePtrOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CodelessUiConnectorConfigPropertiesResponse) *CodelessUiConnectorConfigPropertiesResponse {
		return &v
	}).(CodelessUiConnectorConfigPropertiesResponsePtrOutput)
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





type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput
	ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput
}

type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArgs struct {
	Type  pulumi.StringPtrInput   `pulumi:"type"`
	Value pulumi.StringArrayInput `pulumi:"value"`
}

func (CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArgs) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArgs) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaOutput)
}





type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput
	ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput
}

type CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArray []CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaInput

func (CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseConnectivityCriteria)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArray) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput() CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArray) ToCodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseConnectivityCriteriaArrayOutput)
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





type CodelessUiConnectorConfigPropertiesResponseDataTypesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutput() CodelessUiConnectorConfigPropertiesResponseDataTypesOutput
	ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseDataTypesOutput
}

type CodelessUiConnectorConfigPropertiesResponseDataTypesArgs struct {
	LastDataReceivedQuery pulumi.StringPtrInput `pulumi:"lastDataReceivedQuery"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
}

func (CodelessUiConnectorConfigPropertiesResponseDataTypesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseDataTypes)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseDataTypesArgs) ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutput() CodelessUiConnectorConfigPropertiesResponseDataTypesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseDataTypesArgs) ToCodelessUiConnectorConfigPropertiesResponseDataTypesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseDataTypesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseDataTypesOutput)
}





type CodelessUiConnectorConfigPropertiesResponseDataTypesArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput() CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput
	ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput
}

type CodelessUiConnectorConfigPropertiesResponseDataTypesArray []CodelessUiConnectorConfigPropertiesResponseDataTypesInput

func (CodelessUiConnectorConfigPropertiesResponseDataTypesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseDataTypes)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseDataTypesArray) ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput() CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseDataTypesArray) ToCodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseDataTypesArrayOutput)
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





type CodelessUiConnectorConfigPropertiesResponseGraphQueriesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput() CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput
	ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput
}

type CodelessUiConnectorConfigPropertiesResponseGraphQueriesArgs struct {
	BaseQuery  pulumi.StringPtrInput `pulumi:"baseQuery"`
	Legend     pulumi.StringPtrInput `pulumi:"legend"`
	MetricName pulumi.StringPtrInput `pulumi:"metricName"`
}

func (CodelessUiConnectorConfigPropertiesResponseGraphQueriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseGraphQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseGraphQueriesArgs) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput() CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseGraphQueriesArgs) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseGraphQueriesOutput)
}





type CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput() CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput
	ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput
}

type CodelessUiConnectorConfigPropertiesResponseGraphQueriesArray []CodelessUiConnectorConfigPropertiesResponseGraphQueriesInput

func (CodelessUiConnectorConfigPropertiesResponseGraphQueriesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseGraphQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseGraphQueriesArray) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput() CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseGraphQueriesArray) ToCodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseGraphQueriesArrayOutput)
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





type CodelessUiConnectorConfigPropertiesResponseInstructionStepsInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput() CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput
	ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput
}

type CodelessUiConnectorConfigPropertiesResponseInstructionStepsArgs struct {
	Description  pulumi.StringPtrInput                          `pulumi:"description"`
	Instructions InstructionStepsResponseInstructionsArrayInput `pulumi:"instructions"`
	Title        pulumi.StringPtrInput                          `pulumi:"title"`
}

func (CodelessUiConnectorConfigPropertiesResponseInstructionStepsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseInstructionSteps)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseInstructionStepsArgs) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput() CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseInstructionStepsArgs) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseInstructionStepsOutput)
}





type CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput() CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput
	ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput
}

type CodelessUiConnectorConfigPropertiesResponseInstructionStepsArray []CodelessUiConnectorConfigPropertiesResponseInstructionStepsInput

func (CodelessUiConnectorConfigPropertiesResponseInstructionStepsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseInstructionSteps)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseInstructionStepsArray) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput() CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseInstructionStepsArray) ToCodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseInstructionStepsArrayOutput)
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





type CodelessUiConnectorConfigPropertiesResponseSampleQueriesInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput() CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput
	ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput
}

type CodelessUiConnectorConfigPropertiesResponseSampleQueriesArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Query       pulumi.StringPtrInput `pulumi:"query"`
}

func (CodelessUiConnectorConfigPropertiesResponseSampleQueriesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CodelessUiConnectorConfigPropertiesResponseSampleQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseSampleQueriesArgs) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput() CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseSampleQueriesArgs) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseSampleQueriesOutput)
}





type CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayInput interface {
	pulumi.Input

	ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput() CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput
	ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutputWithContext(context.Context) CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput
}

type CodelessUiConnectorConfigPropertiesResponseSampleQueriesArray []CodelessUiConnectorConfigPropertiesResponseSampleQueriesInput

func (CodelessUiConnectorConfigPropertiesResponseSampleQueriesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CodelessUiConnectorConfigPropertiesResponseSampleQueries)(nil)).Elem()
}

func (i CodelessUiConnectorConfigPropertiesResponseSampleQueriesArray) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput() CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput {
	return i.ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutputWithContext(context.Background())
}

func (i CodelessUiConnectorConfigPropertiesResponseSampleQueriesArray) ToCodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutputWithContext(ctx context.Context) CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CodelessUiConnectorConfigPropertiesResponseSampleQueriesArrayOutput)
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





type ContentPathMapResponseInput interface {
	pulumi.Input

	ToContentPathMapResponseOutput() ContentPathMapResponseOutput
	ToContentPathMapResponseOutputWithContext(context.Context) ContentPathMapResponseOutput
}

type ContentPathMapResponseArgs struct {
	ContentType pulumi.StringPtrInput `pulumi:"contentType"`
	Path        pulumi.StringPtrInput `pulumi:"path"`
}

func (ContentPathMapResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentPathMapResponse)(nil)).Elem()
}

func (i ContentPathMapResponseArgs) ToContentPathMapResponseOutput() ContentPathMapResponseOutput {
	return i.ToContentPathMapResponseOutputWithContext(context.Background())
}

func (i ContentPathMapResponseArgs) ToContentPathMapResponseOutputWithContext(ctx context.Context) ContentPathMapResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentPathMapResponseOutput)
}





type ContentPathMapResponseArrayInput interface {
	pulumi.Input

	ToContentPathMapResponseArrayOutput() ContentPathMapResponseArrayOutput
	ToContentPathMapResponseArrayOutputWithContext(context.Context) ContentPathMapResponseArrayOutput
}

type ContentPathMapResponseArray []ContentPathMapResponseInput

func (ContentPathMapResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContentPathMapResponse)(nil)).Elem()
}

func (i ContentPathMapResponseArray) ToContentPathMapResponseArrayOutput() ContentPathMapResponseArrayOutput {
	return i.ToContentPathMapResponseArrayOutputWithContext(context.Background())
}

func (i ContentPathMapResponseArray) ToContentPathMapResponseArrayOutputWithContext(ctx context.Context) ContentPathMapResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentPathMapResponseArrayOutput)
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





type DataConnectorDataTypeCommonResponseInput interface {
	pulumi.Input

	ToDataConnectorDataTypeCommonResponseOutput() DataConnectorDataTypeCommonResponseOutput
	ToDataConnectorDataTypeCommonResponseOutputWithContext(context.Context) DataConnectorDataTypeCommonResponseOutput
}

type DataConnectorDataTypeCommonResponseArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (DataConnectorDataTypeCommonResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataConnectorDataTypeCommonResponse)(nil)).Elem()
}

func (i DataConnectorDataTypeCommonResponseArgs) ToDataConnectorDataTypeCommonResponseOutput() DataConnectorDataTypeCommonResponseOutput {
	return i.ToDataConnectorDataTypeCommonResponseOutputWithContext(context.Background())
}

func (i DataConnectorDataTypeCommonResponseArgs) ToDataConnectorDataTypeCommonResponseOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorDataTypeCommonResponseOutput)
}

func (i DataConnectorDataTypeCommonResponseArgs) ToDataConnectorDataTypeCommonResponsePtrOutput() DataConnectorDataTypeCommonResponsePtrOutput {
	return i.ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(context.Background())
}

func (i DataConnectorDataTypeCommonResponseArgs) ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorDataTypeCommonResponseOutput).ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(ctx)
}









type DataConnectorDataTypeCommonResponsePtrInput interface {
	pulumi.Input

	ToDataConnectorDataTypeCommonResponsePtrOutput() DataConnectorDataTypeCommonResponsePtrOutput
	ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(context.Context) DataConnectorDataTypeCommonResponsePtrOutput
}

type dataConnectorDataTypeCommonResponsePtrType DataConnectorDataTypeCommonResponseArgs

func DataConnectorDataTypeCommonResponsePtr(v *DataConnectorDataTypeCommonResponseArgs) DataConnectorDataTypeCommonResponsePtrInput {
	return (*dataConnectorDataTypeCommonResponsePtrType)(v)
}

func (*dataConnectorDataTypeCommonResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataConnectorDataTypeCommonResponse)(nil)).Elem()
}

func (i *dataConnectorDataTypeCommonResponsePtrType) ToDataConnectorDataTypeCommonResponsePtrOutput() DataConnectorDataTypeCommonResponsePtrOutput {
	return i.ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(context.Background())
}

func (i *dataConnectorDataTypeCommonResponsePtrType) ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataConnectorDataTypeCommonResponsePtrOutput)
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

func (o DataConnectorDataTypeCommonResponseOutput) ToDataConnectorDataTypeCommonResponsePtrOutput() DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(context.Background())
}

func (o DataConnectorDataTypeCommonResponseOutput) ToDataConnectorDataTypeCommonResponsePtrOutputWithContext(ctx context.Context) DataConnectorDataTypeCommonResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataConnectorDataTypeCommonResponse) *DataConnectorDataTypeCommonResponse {
		return &v
	}).(DataConnectorDataTypeCommonResponsePtrOutput)
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

func (i Dynamics365DataConnectorDataTypesArgs) ToDynamics365DataConnectorDataTypesPtrOutput() Dynamics365DataConnectorDataTypesPtrOutput {
	return i.ToDynamics365DataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesArgs) ToDynamics365DataConnectorDataTypesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesOutput).ToDynamics365DataConnectorDataTypesPtrOutputWithContext(ctx)
}









type Dynamics365DataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesPtrOutput() Dynamics365DataConnectorDataTypesPtrOutput
	ToDynamics365DataConnectorDataTypesPtrOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesPtrOutput
}

type dynamics365DataConnectorDataTypesPtrType Dynamics365DataConnectorDataTypesArgs

func Dynamics365DataConnectorDataTypesPtr(v *Dynamics365DataConnectorDataTypesArgs) Dynamics365DataConnectorDataTypesPtrInput {
	return (*dynamics365DataConnectorDataTypesPtrType)(v)
}

func (*dynamics365DataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypes)(nil)).Elem()
}

func (i *dynamics365DataConnectorDataTypesPtrType) ToDynamics365DataConnectorDataTypesPtrOutput() Dynamics365DataConnectorDataTypesPtrOutput {
	return i.ToDynamics365DataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *dynamics365DataConnectorDataTypesPtrType) ToDynamics365DataConnectorDataTypesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesPtrOutput)
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

func (o Dynamics365DataConnectorDataTypesOutput) ToDynamics365DataConnectorDataTypesPtrOutput() Dynamics365DataConnectorDataTypesPtrOutput {
	return o.ToDynamics365DataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o Dynamics365DataConnectorDataTypesOutput) ToDynamics365DataConnectorDataTypesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Dynamics365DataConnectorDataTypes) *Dynamics365DataConnectorDataTypes {
		return &v
	}).(Dynamics365DataConnectorDataTypesPtrOutput)
}

func (o Dynamics365DataConnectorDataTypesOutput) Dynamics365CdsActivities() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypes) Dynamics365DataConnectorDataTypesDynamics365CdsActivities {
		return v.Dynamics365CdsActivities
	}).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput)
}

type Dynamics365DataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypes)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesPtrOutput) ToDynamics365DataConnectorDataTypesPtrOutput() Dynamics365DataConnectorDataTypesPtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesPtrOutput) ToDynamics365DataConnectorDataTypesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesPtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesPtrOutput) Elem() Dynamics365DataConnectorDataTypesOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypes) Dynamics365DataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret Dynamics365DataConnectorDataTypes
		return ret
	}).(Dynamics365DataConnectorDataTypesOutput)
}

func (o Dynamics365DataConnectorDataTypesPtrOutput) Dynamics365CdsActivities() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypes) *Dynamics365DataConnectorDataTypesDynamics365CdsActivities {
		if v == nil {
			return nil
		}
		return &v.Dynamics365CdsActivities
	}).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput)
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

func (i Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return i.ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput).ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(ctx)
}









type Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput
	ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput
}

type dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrType Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs

func Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtr(v *Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrInput {
	return (*dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrType)(v)
}

func (*dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypesDynamics365CdsActivities)(nil)).Elem()
}

func (i *dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrType) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return i.ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(context.Background())
}

func (i *dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrType) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput)
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

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return o.ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(context.Background())
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Dynamics365DataConnectorDataTypesDynamics365CdsActivities) *Dynamics365DataConnectorDataTypesDynamics365CdsActivities {
		return &v
	}).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput)
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypesDynamics365CdsActivities) string { return v.State }).(pulumi.StringOutput)
}

type Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypesDynamics365CdsActivities)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput) ToDynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput) Elem() Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypesDynamics365CdsActivities) Dynamics365DataConnectorDataTypesDynamics365CdsActivities {
		if v != nil {
			return *v
		}
		var ret Dynamics365DataConnectorDataTypesDynamics365CdsActivities
		return ret
	}).(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput)
}

func (o Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypesDynamics365CdsActivities) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type Dynamics365DataConnectorDataTypesResponse struct {
	Dynamics365CdsActivities Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities `pulumi:"dynamics365CdsActivities"`
}





type Dynamics365DataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesResponseOutput() Dynamics365DataConnectorDataTypesResponseOutput
	ToDynamics365DataConnectorDataTypesResponseOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesResponseOutput
}

type Dynamics365DataConnectorDataTypesResponseArgs struct {
	Dynamics365CdsActivities Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesInput `pulumi:"dynamics365CdsActivities"`
}

func (Dynamics365DataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponse)(nil)).Elem()
}

func (i Dynamics365DataConnectorDataTypesResponseArgs) ToDynamics365DataConnectorDataTypesResponseOutput() Dynamics365DataConnectorDataTypesResponseOutput {
	return i.ToDynamics365DataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesResponseArgs) ToDynamics365DataConnectorDataTypesResponseOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesResponseOutput)
}

func (i Dynamics365DataConnectorDataTypesResponseArgs) ToDynamics365DataConnectorDataTypesResponsePtrOutput() Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return i.ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesResponseArgs) ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesResponseOutput).ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type Dynamics365DataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesResponsePtrOutput() Dynamics365DataConnectorDataTypesResponsePtrOutput
	ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesResponsePtrOutput
}

type dynamics365DataConnectorDataTypesResponsePtrType Dynamics365DataConnectorDataTypesResponseArgs

func Dynamics365DataConnectorDataTypesResponsePtr(v *Dynamics365DataConnectorDataTypesResponseArgs) Dynamics365DataConnectorDataTypesResponsePtrInput {
	return (*dynamics365DataConnectorDataTypesResponsePtrType)(v)
}

func (*dynamics365DataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *dynamics365DataConnectorDataTypesResponsePtrType) ToDynamics365DataConnectorDataTypesResponsePtrOutput() Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return i.ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *dynamics365DataConnectorDataTypesResponsePtrType) ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesResponsePtrOutput)
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

func (o Dynamics365DataConnectorDataTypesResponseOutput) ToDynamics365DataConnectorDataTypesResponsePtrOutput() Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return o.ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o Dynamics365DataConnectorDataTypesResponseOutput) ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Dynamics365DataConnectorDataTypesResponse) *Dynamics365DataConnectorDataTypesResponse {
		return &v
	}).(Dynamics365DataConnectorDataTypesResponsePtrOutput)
}

func (o Dynamics365DataConnectorDataTypesResponseOutput) Dynamics365CdsActivities() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypesResponse) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities {
		return v.Dynamics365CdsActivities
	}).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput)
}

type Dynamics365DataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypesResponse)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesResponsePtrOutput) ToDynamics365DataConnectorDataTypesResponsePtrOutput() Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponsePtrOutput) ToDynamics365DataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponsePtrOutput) Elem() Dynamics365DataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypesResponse) Dynamics365DataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret Dynamics365DataConnectorDataTypesResponse
		return ret
	}).(Dynamics365DataConnectorDataTypesResponseOutput)
}

func (o Dynamics365DataConnectorDataTypesResponsePtrOutput) Dynamics365CdsActivities() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypesResponse) *Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities {
		if v == nil {
			return nil
		}
		return &v.Dynamics365CdsActivities
	}).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput)
}

type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities struct {
	State string `pulumi:"state"`
}





type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput
	ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput
}

type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities)(nil)).Elem()
}

func (i Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return i.ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput)
}

func (i Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return i.ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(context.Background())
}

func (i Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput).ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(ctx)
}









type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrInput interface {
	pulumi.Input

	ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput
	ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput
}

type dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrType Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs

func Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtr(v *Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrInput {
	return (*dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrType)(v)
}

func (*dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities)(nil)).Elem()
}

func (i *dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrType) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return i.ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(context.Background())
}

func (i *dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrType) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput)
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

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return o.ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(context.Background())
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities) *Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities {
		return &v
	}).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput)
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities) string { return v.State }).(pulumi.StringOutput)
}

type Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput struct{ *pulumi.OutputState }

func (Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities)(nil)).Elem()
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput) ToDynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutputWithContext(ctx context.Context) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput {
	return o
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput) Elem() Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities) Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities {
		if v != nil {
			return *v
		}
		var ret Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities
		return ret
	}).(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput)
}

func (o Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivities) *string {
		if v == nil {
			return nil
		}
		return &v.State
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





type EntityMappingResponseInput interface {
	pulumi.Input

	ToEntityMappingResponseOutput() EntityMappingResponseOutput
	ToEntityMappingResponseOutputWithContext(context.Context) EntityMappingResponseOutput
}

type EntityMappingResponseArgs struct {
	EntityType    pulumi.StringPtrInput          `pulumi:"entityType"`
	FieldMappings FieldMappingResponseArrayInput `pulumi:"fieldMappings"`
}

func (EntityMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityMappingResponse)(nil)).Elem()
}

func (i EntityMappingResponseArgs) ToEntityMappingResponseOutput() EntityMappingResponseOutput {
	return i.ToEntityMappingResponseOutputWithContext(context.Background())
}

func (i EntityMappingResponseArgs) ToEntityMappingResponseOutputWithContext(ctx context.Context) EntityMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityMappingResponseOutput)
}





type EntityMappingResponseArrayInput interface {
	pulumi.Input

	ToEntityMappingResponseArrayOutput() EntityMappingResponseArrayOutput
	ToEntityMappingResponseArrayOutputWithContext(context.Context) EntityMappingResponseArrayOutput
}

type EntityMappingResponseArray []EntityMappingResponseInput

func (EntityMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityMappingResponse)(nil)).Elem()
}

func (i EntityMappingResponseArray) ToEntityMappingResponseArrayOutput() EntityMappingResponseArrayOutput {
	return i.ToEntityMappingResponseArrayOutputWithContext(context.Background())
}

func (i EntityMappingResponseArray) ToEntityMappingResponseArrayOutputWithContext(ctx context.Context) EntityMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityMappingResponseArrayOutput)
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





type EventGroupingSettingsResponseInput interface {
	pulumi.Input

	ToEventGroupingSettingsResponseOutput() EventGroupingSettingsResponseOutput
	ToEventGroupingSettingsResponseOutputWithContext(context.Context) EventGroupingSettingsResponseOutput
}

type EventGroupingSettingsResponseArgs struct {
	AggregationKind pulumi.StringPtrInput `pulumi:"aggregationKind"`
}

func (EventGroupingSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventGroupingSettingsResponse)(nil)).Elem()
}

func (i EventGroupingSettingsResponseArgs) ToEventGroupingSettingsResponseOutput() EventGroupingSettingsResponseOutput {
	return i.ToEventGroupingSettingsResponseOutputWithContext(context.Background())
}

func (i EventGroupingSettingsResponseArgs) ToEventGroupingSettingsResponseOutputWithContext(ctx context.Context) EventGroupingSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGroupingSettingsResponseOutput)
}

func (i EventGroupingSettingsResponseArgs) ToEventGroupingSettingsResponsePtrOutput() EventGroupingSettingsResponsePtrOutput {
	return i.ToEventGroupingSettingsResponsePtrOutputWithContext(context.Background())
}

func (i EventGroupingSettingsResponseArgs) ToEventGroupingSettingsResponsePtrOutputWithContext(ctx context.Context) EventGroupingSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGroupingSettingsResponseOutput).ToEventGroupingSettingsResponsePtrOutputWithContext(ctx)
}









type EventGroupingSettingsResponsePtrInput interface {
	pulumi.Input

	ToEventGroupingSettingsResponsePtrOutput() EventGroupingSettingsResponsePtrOutput
	ToEventGroupingSettingsResponsePtrOutputWithContext(context.Context) EventGroupingSettingsResponsePtrOutput
}

type eventGroupingSettingsResponsePtrType EventGroupingSettingsResponseArgs

func EventGroupingSettingsResponsePtr(v *EventGroupingSettingsResponseArgs) EventGroupingSettingsResponsePtrInput {
	return (*eventGroupingSettingsResponsePtrType)(v)
}

func (*eventGroupingSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventGroupingSettingsResponse)(nil)).Elem()
}

func (i *eventGroupingSettingsResponsePtrType) ToEventGroupingSettingsResponsePtrOutput() EventGroupingSettingsResponsePtrOutput {
	return i.ToEventGroupingSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *eventGroupingSettingsResponsePtrType) ToEventGroupingSettingsResponsePtrOutputWithContext(ctx context.Context) EventGroupingSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventGroupingSettingsResponsePtrOutput)
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

func (o EventGroupingSettingsResponseOutput) ToEventGroupingSettingsResponsePtrOutput() EventGroupingSettingsResponsePtrOutput {
	return o.ToEventGroupingSettingsResponsePtrOutputWithContext(context.Background())
}

func (o EventGroupingSettingsResponseOutput) ToEventGroupingSettingsResponsePtrOutputWithContext(ctx context.Context) EventGroupingSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventGroupingSettingsResponse) *EventGroupingSettingsResponse {
		return &v
	}).(EventGroupingSettingsResponsePtrOutput)
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





type FieldMappingResponseInput interface {
	pulumi.Input

	ToFieldMappingResponseOutput() FieldMappingResponseOutput
	ToFieldMappingResponseOutputWithContext(context.Context) FieldMappingResponseOutput
}

type FieldMappingResponseArgs struct {
	ColumnName pulumi.StringPtrInput `pulumi:"columnName"`
	Identifier pulumi.StringPtrInput `pulumi:"identifier"`
}

func (FieldMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldMappingResponse)(nil)).Elem()
}

func (i FieldMappingResponseArgs) ToFieldMappingResponseOutput() FieldMappingResponseOutput {
	return i.ToFieldMappingResponseOutputWithContext(context.Background())
}

func (i FieldMappingResponseArgs) ToFieldMappingResponseOutputWithContext(ctx context.Context) FieldMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldMappingResponseOutput)
}





type FieldMappingResponseArrayInput interface {
	pulumi.Input

	ToFieldMappingResponseArrayOutput() FieldMappingResponseArrayOutput
	ToFieldMappingResponseArrayOutputWithContext(context.Context) FieldMappingResponseArrayOutput
}

type FieldMappingResponseArray []FieldMappingResponseInput

func (FieldMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FieldMappingResponse)(nil)).Elem()
}

func (i FieldMappingResponseArray) ToFieldMappingResponseArrayOutput() FieldMappingResponseArrayOutput {
	return i.ToFieldMappingResponseArrayOutputWithContext(context.Background())
}

func (i FieldMappingResponseArray) ToFieldMappingResponseArrayOutputWithContext(ctx context.Context) FieldMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldMappingResponseArrayOutput)
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





type GroupingConfigurationResponseInput interface {
	pulumi.Input

	ToGroupingConfigurationResponseOutput() GroupingConfigurationResponseOutput
	ToGroupingConfigurationResponseOutputWithContext(context.Context) GroupingConfigurationResponseOutput
}

type GroupingConfigurationResponseArgs struct {
	Enabled              pulumi.BoolInput        `pulumi:"enabled"`
	GroupByAlertDetails  pulumi.StringArrayInput `pulumi:"groupByAlertDetails"`
	GroupByCustomDetails pulumi.StringArrayInput `pulumi:"groupByCustomDetails"`
	GroupByEntities      pulumi.StringArrayInput `pulumi:"groupByEntities"`
	LookbackDuration     pulumi.StringInput      `pulumi:"lookbackDuration"`
	MatchingMethod       pulumi.StringInput      `pulumi:"matchingMethod"`
	ReopenClosedIncident pulumi.BoolInput        `pulumi:"reopenClosedIncident"`
}

func (GroupingConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GroupingConfigurationResponse)(nil)).Elem()
}

func (i GroupingConfigurationResponseArgs) ToGroupingConfigurationResponseOutput() GroupingConfigurationResponseOutput {
	return i.ToGroupingConfigurationResponseOutputWithContext(context.Background())
}

func (i GroupingConfigurationResponseArgs) ToGroupingConfigurationResponseOutputWithContext(ctx context.Context) GroupingConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingConfigurationResponseOutput)
}

func (i GroupingConfigurationResponseArgs) ToGroupingConfigurationResponsePtrOutput() GroupingConfigurationResponsePtrOutput {
	return i.ToGroupingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i GroupingConfigurationResponseArgs) ToGroupingConfigurationResponsePtrOutputWithContext(ctx context.Context) GroupingConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingConfigurationResponseOutput).ToGroupingConfigurationResponsePtrOutputWithContext(ctx)
}









type GroupingConfigurationResponsePtrInput interface {
	pulumi.Input

	ToGroupingConfigurationResponsePtrOutput() GroupingConfigurationResponsePtrOutput
	ToGroupingConfigurationResponsePtrOutputWithContext(context.Context) GroupingConfigurationResponsePtrOutput
}

type groupingConfigurationResponsePtrType GroupingConfigurationResponseArgs

func GroupingConfigurationResponsePtr(v *GroupingConfigurationResponseArgs) GroupingConfigurationResponsePtrInput {
	return (*groupingConfigurationResponsePtrType)(v)
}

func (*groupingConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GroupingConfigurationResponse)(nil)).Elem()
}

func (i *groupingConfigurationResponsePtrType) ToGroupingConfigurationResponsePtrOutput() GroupingConfigurationResponsePtrOutput {
	return i.ToGroupingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *groupingConfigurationResponsePtrType) ToGroupingConfigurationResponsePtrOutputWithContext(ctx context.Context) GroupingConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GroupingConfigurationResponsePtrOutput)
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

func (o GroupingConfigurationResponseOutput) ToGroupingConfigurationResponsePtrOutput() GroupingConfigurationResponsePtrOutput {
	return o.ToGroupingConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o GroupingConfigurationResponseOutput) ToGroupingConfigurationResponsePtrOutputWithContext(ctx context.Context) GroupingConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GroupingConfigurationResponse) *GroupingConfigurationResponse {
		return &v
	}).(GroupingConfigurationResponsePtrOutput)
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





type IncidentAdditionalDataResponseInput interface {
	pulumi.Input

	ToIncidentAdditionalDataResponseOutput() IncidentAdditionalDataResponseOutput
	ToIncidentAdditionalDataResponseOutputWithContext(context.Context) IncidentAdditionalDataResponseOutput
}

type IncidentAdditionalDataResponseArgs struct {
	AlertProductNames pulumi.StringArrayInput `pulumi:"alertProductNames"`
	AlertsCount       pulumi.IntInput         `pulumi:"alertsCount"`
	BookmarksCount    pulumi.IntInput         `pulumi:"bookmarksCount"`
	CommentsCount     pulumi.IntInput         `pulumi:"commentsCount"`
	Tactics           pulumi.StringArrayInput `pulumi:"tactics"`
}

func (IncidentAdditionalDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentAdditionalDataResponse)(nil)).Elem()
}

func (i IncidentAdditionalDataResponseArgs) ToIncidentAdditionalDataResponseOutput() IncidentAdditionalDataResponseOutput {
	return i.ToIncidentAdditionalDataResponseOutputWithContext(context.Background())
}

func (i IncidentAdditionalDataResponseArgs) ToIncidentAdditionalDataResponseOutputWithContext(ctx context.Context) IncidentAdditionalDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentAdditionalDataResponseOutput)
}

func (i IncidentAdditionalDataResponseArgs) ToIncidentAdditionalDataResponsePtrOutput() IncidentAdditionalDataResponsePtrOutput {
	return i.ToIncidentAdditionalDataResponsePtrOutputWithContext(context.Background())
}

func (i IncidentAdditionalDataResponseArgs) ToIncidentAdditionalDataResponsePtrOutputWithContext(ctx context.Context) IncidentAdditionalDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentAdditionalDataResponseOutput).ToIncidentAdditionalDataResponsePtrOutputWithContext(ctx)
}









type IncidentAdditionalDataResponsePtrInput interface {
	pulumi.Input

	ToIncidentAdditionalDataResponsePtrOutput() IncidentAdditionalDataResponsePtrOutput
	ToIncidentAdditionalDataResponsePtrOutputWithContext(context.Context) IncidentAdditionalDataResponsePtrOutput
}

type incidentAdditionalDataResponsePtrType IncidentAdditionalDataResponseArgs

func IncidentAdditionalDataResponsePtr(v *IncidentAdditionalDataResponseArgs) IncidentAdditionalDataResponsePtrInput {
	return (*incidentAdditionalDataResponsePtrType)(v)
}

func (*incidentAdditionalDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentAdditionalDataResponse)(nil)).Elem()
}

func (i *incidentAdditionalDataResponsePtrType) ToIncidentAdditionalDataResponsePtrOutput() IncidentAdditionalDataResponsePtrOutput {
	return i.ToIncidentAdditionalDataResponsePtrOutputWithContext(context.Background())
}

func (i *incidentAdditionalDataResponsePtrType) ToIncidentAdditionalDataResponsePtrOutputWithContext(ctx context.Context) IncidentAdditionalDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentAdditionalDataResponsePtrOutput)
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

func (o IncidentAdditionalDataResponseOutput) ToIncidentAdditionalDataResponsePtrOutput() IncidentAdditionalDataResponsePtrOutput {
	return o.ToIncidentAdditionalDataResponsePtrOutputWithContext(context.Background())
}

func (o IncidentAdditionalDataResponseOutput) ToIncidentAdditionalDataResponsePtrOutputWithContext(ctx context.Context) IncidentAdditionalDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentAdditionalDataResponse) *IncidentAdditionalDataResponse {
		return &v
	}).(IncidentAdditionalDataResponsePtrOutput)
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

type IncidentAdditionalDataResponsePtrOutput struct{ *pulumi.OutputState }

func (IncidentAdditionalDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentAdditionalDataResponse)(nil)).Elem()
}

func (o IncidentAdditionalDataResponsePtrOutput) ToIncidentAdditionalDataResponsePtrOutput() IncidentAdditionalDataResponsePtrOutput {
	return o
}

func (o IncidentAdditionalDataResponsePtrOutput) ToIncidentAdditionalDataResponsePtrOutputWithContext(ctx context.Context) IncidentAdditionalDataResponsePtrOutput {
	return o
}

func (o IncidentAdditionalDataResponsePtrOutput) Elem() IncidentAdditionalDataResponseOutput {
	return o.ApplyT(func(v *IncidentAdditionalDataResponse) IncidentAdditionalDataResponse {
		if v != nil {
			return *v
		}
		var ret IncidentAdditionalDataResponse
		return ret
	}).(IncidentAdditionalDataResponseOutput)
}

func (o IncidentAdditionalDataResponsePtrOutput) AlertProductNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IncidentAdditionalDataResponse) []string {
		if v == nil {
			return nil
		}
		return v.AlertProductNames
	}).(pulumi.StringArrayOutput)
}

func (o IncidentAdditionalDataResponsePtrOutput) AlertsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IncidentAdditionalDataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.AlertsCount
	}).(pulumi.IntPtrOutput)
}

func (o IncidentAdditionalDataResponsePtrOutput) BookmarksCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IncidentAdditionalDataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.BookmarksCount
	}).(pulumi.IntPtrOutput)
}

func (o IncidentAdditionalDataResponsePtrOutput) CommentsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IncidentAdditionalDataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.CommentsCount
	}).(pulumi.IntPtrOutput)
}

func (o IncidentAdditionalDataResponsePtrOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *IncidentAdditionalDataResponse) []string {
		if v == nil {
			return nil
		}
		return v.Tactics
	}).(pulumi.StringArrayOutput)
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





type IncidentConfigurationResponseInput interface {
	pulumi.Input

	ToIncidentConfigurationResponseOutput() IncidentConfigurationResponseOutput
	ToIncidentConfigurationResponseOutputWithContext(context.Context) IncidentConfigurationResponseOutput
}

type IncidentConfigurationResponseArgs struct {
	CreateIncident        pulumi.BoolInput                      `pulumi:"createIncident"`
	GroupingConfiguration GroupingConfigurationResponsePtrInput `pulumi:"groupingConfiguration"`
}

func (IncidentConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentConfigurationResponse)(nil)).Elem()
}

func (i IncidentConfigurationResponseArgs) ToIncidentConfigurationResponseOutput() IncidentConfigurationResponseOutput {
	return i.ToIncidentConfigurationResponseOutputWithContext(context.Background())
}

func (i IncidentConfigurationResponseArgs) ToIncidentConfigurationResponseOutputWithContext(ctx context.Context) IncidentConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentConfigurationResponseOutput)
}

func (i IncidentConfigurationResponseArgs) ToIncidentConfigurationResponsePtrOutput() IncidentConfigurationResponsePtrOutput {
	return i.ToIncidentConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i IncidentConfigurationResponseArgs) ToIncidentConfigurationResponsePtrOutputWithContext(ctx context.Context) IncidentConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentConfigurationResponseOutput).ToIncidentConfigurationResponsePtrOutputWithContext(ctx)
}









type IncidentConfigurationResponsePtrInput interface {
	pulumi.Input

	ToIncidentConfigurationResponsePtrOutput() IncidentConfigurationResponsePtrOutput
	ToIncidentConfigurationResponsePtrOutputWithContext(context.Context) IncidentConfigurationResponsePtrOutput
}

type incidentConfigurationResponsePtrType IncidentConfigurationResponseArgs

func IncidentConfigurationResponsePtr(v *IncidentConfigurationResponseArgs) IncidentConfigurationResponsePtrInput {
	return (*incidentConfigurationResponsePtrType)(v)
}

func (*incidentConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentConfigurationResponse)(nil)).Elem()
}

func (i *incidentConfigurationResponsePtrType) ToIncidentConfigurationResponsePtrOutput() IncidentConfigurationResponsePtrOutput {
	return i.ToIncidentConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *incidentConfigurationResponsePtrType) ToIncidentConfigurationResponsePtrOutputWithContext(ctx context.Context) IncidentConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentConfigurationResponsePtrOutput)
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

func (o IncidentConfigurationResponseOutput) ToIncidentConfigurationResponsePtrOutput() IncidentConfigurationResponsePtrOutput {
	return o.ToIncidentConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o IncidentConfigurationResponseOutput) ToIncidentConfigurationResponsePtrOutputWithContext(ctx context.Context) IncidentConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentConfigurationResponse) *IncidentConfigurationResponse {
		return &v
	}).(IncidentConfigurationResponsePtrOutput)
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





type IncidentLabelResponseInput interface {
	pulumi.Input

	ToIncidentLabelResponseOutput() IncidentLabelResponseOutput
	ToIncidentLabelResponseOutputWithContext(context.Context) IncidentLabelResponseOutput
}

type IncidentLabelResponseArgs struct {
	LabelName pulumi.StringInput `pulumi:"labelName"`
	LabelType pulumi.StringInput `pulumi:"labelType"`
}

func (IncidentLabelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentLabelResponse)(nil)).Elem()
}

func (i IncidentLabelResponseArgs) ToIncidentLabelResponseOutput() IncidentLabelResponseOutput {
	return i.ToIncidentLabelResponseOutputWithContext(context.Background())
}

func (i IncidentLabelResponseArgs) ToIncidentLabelResponseOutputWithContext(ctx context.Context) IncidentLabelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentLabelResponseOutput)
}





type IncidentLabelResponseArrayInput interface {
	pulumi.Input

	ToIncidentLabelResponseArrayOutput() IncidentLabelResponseArrayOutput
	ToIncidentLabelResponseArrayOutputWithContext(context.Context) IncidentLabelResponseArrayOutput
}

type IncidentLabelResponseArray []IncidentLabelResponseInput

func (IncidentLabelResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncidentLabelResponse)(nil)).Elem()
}

func (i IncidentLabelResponseArray) ToIncidentLabelResponseArrayOutput() IncidentLabelResponseArrayOutput {
	return i.ToIncidentLabelResponseArrayOutputWithContext(context.Background())
}

func (i IncidentLabelResponseArray) ToIncidentLabelResponseArrayOutputWithContext(ctx context.Context) IncidentLabelResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentLabelResponseArrayOutput)
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
	OwnerType         string  `pulumi:"ownerType"`
	UserPrincipalName *string `pulumi:"userPrincipalName"`
}





type IncidentOwnerInfoResponseInput interface {
	pulumi.Input

	ToIncidentOwnerInfoResponseOutput() IncidentOwnerInfoResponseOutput
	ToIncidentOwnerInfoResponseOutputWithContext(context.Context) IncidentOwnerInfoResponseOutput
}

type IncidentOwnerInfoResponseArgs struct {
	AssignedTo        pulumi.StringPtrInput `pulumi:"assignedTo"`
	Email             pulumi.StringPtrInput `pulumi:"email"`
	ObjectId          pulumi.StringPtrInput `pulumi:"objectId"`
	OwnerType         pulumi.StringInput    `pulumi:"ownerType"`
	UserPrincipalName pulumi.StringPtrInput `pulumi:"userPrincipalName"`
}

func (IncidentOwnerInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentOwnerInfoResponse)(nil)).Elem()
}

func (i IncidentOwnerInfoResponseArgs) ToIncidentOwnerInfoResponseOutput() IncidentOwnerInfoResponseOutput {
	return i.ToIncidentOwnerInfoResponseOutputWithContext(context.Background())
}

func (i IncidentOwnerInfoResponseArgs) ToIncidentOwnerInfoResponseOutputWithContext(ctx context.Context) IncidentOwnerInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOwnerInfoResponseOutput)
}

func (i IncidentOwnerInfoResponseArgs) ToIncidentOwnerInfoResponsePtrOutput() IncidentOwnerInfoResponsePtrOutput {
	return i.ToIncidentOwnerInfoResponsePtrOutputWithContext(context.Background())
}

func (i IncidentOwnerInfoResponseArgs) ToIncidentOwnerInfoResponsePtrOutputWithContext(ctx context.Context) IncidentOwnerInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOwnerInfoResponseOutput).ToIncidentOwnerInfoResponsePtrOutputWithContext(ctx)
}









type IncidentOwnerInfoResponsePtrInput interface {
	pulumi.Input

	ToIncidentOwnerInfoResponsePtrOutput() IncidentOwnerInfoResponsePtrOutput
	ToIncidentOwnerInfoResponsePtrOutputWithContext(context.Context) IncidentOwnerInfoResponsePtrOutput
}

type incidentOwnerInfoResponsePtrType IncidentOwnerInfoResponseArgs

func IncidentOwnerInfoResponsePtr(v *IncidentOwnerInfoResponseArgs) IncidentOwnerInfoResponsePtrInput {
	return (*incidentOwnerInfoResponsePtrType)(v)
}

func (*incidentOwnerInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentOwnerInfoResponse)(nil)).Elem()
}

func (i *incidentOwnerInfoResponsePtrType) ToIncidentOwnerInfoResponsePtrOutput() IncidentOwnerInfoResponsePtrOutput {
	return i.ToIncidentOwnerInfoResponsePtrOutputWithContext(context.Background())
}

func (i *incidentOwnerInfoResponsePtrType) ToIncidentOwnerInfoResponsePtrOutputWithContext(ctx context.Context) IncidentOwnerInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentOwnerInfoResponsePtrOutput)
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

func (o IncidentOwnerInfoResponseOutput) ToIncidentOwnerInfoResponsePtrOutput() IncidentOwnerInfoResponsePtrOutput {
	return o.ToIncidentOwnerInfoResponsePtrOutputWithContext(context.Background())
}

func (o IncidentOwnerInfoResponseOutput) ToIncidentOwnerInfoResponsePtrOutputWithContext(ctx context.Context) IncidentOwnerInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentOwnerInfoResponse) *IncidentOwnerInfoResponse {
		return &v
	}).(IncidentOwnerInfoResponsePtrOutput)
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

func (o IncidentOwnerInfoResponseOutput) OwnerType() pulumi.StringOutput {
	return o.ApplyT(func(v IncidentOwnerInfoResponse) string { return v.OwnerType }).(pulumi.StringOutput)
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
		return &v.OwnerType
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





type InstructionStepsResponseInstructionsInput interface {
	pulumi.Input

	ToInstructionStepsResponseInstructionsOutput() InstructionStepsResponseInstructionsOutput
	ToInstructionStepsResponseInstructionsOutputWithContext(context.Context) InstructionStepsResponseInstructionsOutput
}

type InstructionStepsResponseInstructionsArgs struct {
	Parameters pulumi.Input       `pulumi:"parameters"`
	Type       pulumi.StringInput `pulumi:"type"`
}

func (InstructionStepsResponseInstructionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InstructionStepsResponseInstructions)(nil)).Elem()
}

func (i InstructionStepsResponseInstructionsArgs) ToInstructionStepsResponseInstructionsOutput() InstructionStepsResponseInstructionsOutput {
	return i.ToInstructionStepsResponseInstructionsOutputWithContext(context.Background())
}

func (i InstructionStepsResponseInstructionsArgs) ToInstructionStepsResponseInstructionsOutputWithContext(ctx context.Context) InstructionStepsResponseInstructionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstructionStepsResponseInstructionsOutput)
}





type InstructionStepsResponseInstructionsArrayInput interface {
	pulumi.Input

	ToInstructionStepsResponseInstructionsArrayOutput() InstructionStepsResponseInstructionsArrayOutput
	ToInstructionStepsResponseInstructionsArrayOutputWithContext(context.Context) InstructionStepsResponseInstructionsArrayOutput
}

type InstructionStepsResponseInstructionsArray []InstructionStepsResponseInstructionsInput

func (InstructionStepsResponseInstructionsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InstructionStepsResponseInstructions)(nil)).Elem()
}

func (i InstructionStepsResponseInstructionsArray) ToInstructionStepsResponseInstructionsArrayOutput() InstructionStepsResponseInstructionsArrayOutput {
	return i.ToInstructionStepsResponseInstructionsArrayOutputWithContext(context.Background())
}

func (i InstructionStepsResponseInstructionsArray) ToInstructionStepsResponseInstructionsArrayOutputWithContext(ctx context.Context) InstructionStepsResponseInstructionsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InstructionStepsResponseInstructionsArrayOutput)
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

func (o MCASDataConnectorDataTypesOutput) Alerts() DataConnectorDataTypeCommonOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) DataConnectorDataTypeCommon { return v.Alerts }).(DataConnectorDataTypeCommonOutput)
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
		return &v.Alerts
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
	Alerts        DataConnectorDataTypeCommonResponse  `pulumi:"alerts"`
	DiscoveryLogs *DataConnectorDataTypeCommonResponse `pulumi:"discoveryLogs"`
}





type MCASDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesResponseOutput() MCASDataConnectorDataTypesResponseOutput
	ToMCASDataConnectorDataTypesResponseOutputWithContext(context.Context) MCASDataConnectorDataTypesResponseOutput
}

type MCASDataConnectorDataTypesResponseArgs struct {
	Alerts        DataConnectorDataTypeCommonResponseInput    `pulumi:"alerts"`
	DiscoveryLogs DataConnectorDataTypeCommonResponsePtrInput `pulumi:"discoveryLogs"`
}

func (MCASDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i MCASDataConnectorDataTypesResponseArgs) ToMCASDataConnectorDataTypesResponseOutput() MCASDataConnectorDataTypesResponseOutput {
	return i.ToMCASDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesResponseArgs) ToMCASDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesResponseOutput)
}

func (i MCASDataConnectorDataTypesResponseArgs) ToMCASDataConnectorDataTypesResponsePtrOutput() MCASDataConnectorDataTypesResponsePtrOutput {
	return i.ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesResponseArgs) ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesResponseOutput).ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type MCASDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesResponsePtrOutput() MCASDataConnectorDataTypesResponsePtrOutput
	ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) MCASDataConnectorDataTypesResponsePtrOutput
}

type mcasdataConnectorDataTypesResponsePtrType MCASDataConnectorDataTypesResponseArgs

func MCASDataConnectorDataTypesResponsePtr(v *MCASDataConnectorDataTypesResponseArgs) MCASDataConnectorDataTypesResponsePtrInput {
	return (*mcasdataConnectorDataTypesResponsePtrType)(v)
}

func (*mcasdataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *mcasdataConnectorDataTypesResponsePtrType) ToMCASDataConnectorDataTypesResponsePtrOutput() MCASDataConnectorDataTypesResponsePtrOutput {
	return i.ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *mcasdataConnectorDataTypesResponsePtrType) ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesResponsePtrOutput)
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

func (o MCASDataConnectorDataTypesResponseOutput) ToMCASDataConnectorDataTypesResponsePtrOutput() MCASDataConnectorDataTypesResponsePtrOutput {
	return o.ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o MCASDataConnectorDataTypesResponseOutput) ToMCASDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MCASDataConnectorDataTypesResponse) *MCASDataConnectorDataTypesResponse {
		return &v
	}).(MCASDataConnectorDataTypesResponsePtrOutput)
}

func (o MCASDataConnectorDataTypesResponseOutput) Alerts() DataConnectorDataTypeCommonResponseOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypesResponse) DataConnectorDataTypeCommonResponse { return v.Alerts }).(DataConnectorDataTypeCommonResponseOutput)
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
		return &v.Alerts
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

func (i MSTIDataConnectorDataTypesArgs) ToMSTIDataConnectorDataTypesPtrOutput() MSTIDataConnectorDataTypesPtrOutput {
	return i.ToMSTIDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesArgs) ToMSTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesOutput).ToMSTIDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type MSTIDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesPtrOutput() MSTIDataConnectorDataTypesPtrOutput
	ToMSTIDataConnectorDataTypesPtrOutputWithContext(context.Context) MSTIDataConnectorDataTypesPtrOutput
}

type mstidataConnectorDataTypesPtrType MSTIDataConnectorDataTypesArgs

func MSTIDataConnectorDataTypesPtr(v *MSTIDataConnectorDataTypesArgs) MSTIDataConnectorDataTypesPtrInput {
	return (*mstidataConnectorDataTypesPtrType)(v)
}

func (*mstidataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypes)(nil)).Elem()
}

func (i *mstidataConnectorDataTypesPtrType) ToMSTIDataConnectorDataTypesPtrOutput() MSTIDataConnectorDataTypesPtrOutput {
	return i.ToMSTIDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *mstidataConnectorDataTypesPtrType) ToMSTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesPtrOutput)
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

func (o MSTIDataConnectorDataTypesOutput) ToMSTIDataConnectorDataTypesPtrOutput() MSTIDataConnectorDataTypesPtrOutput {
	return o.ToMSTIDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o MSTIDataConnectorDataTypesOutput) ToMSTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MSTIDataConnectorDataTypes) *MSTIDataConnectorDataTypes {
		return &v
	}).(MSTIDataConnectorDataTypesPtrOutput)
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

type MSTIDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypes)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesPtrOutput) ToMSTIDataConnectorDataTypesPtrOutput() MSTIDataConnectorDataTypesPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesPtrOutput) ToMSTIDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesPtrOutput) Elem() MSTIDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypes) MSTIDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret MSTIDataConnectorDataTypes
		return ret
	}).(MSTIDataConnectorDataTypesOutput)
}

func (o MSTIDataConnectorDataTypesPtrOutput) BingSafetyPhishingURL() MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypes) *MSTIDataConnectorDataTypesBingSafetyPhishingURL {
		if v == nil {
			return nil
		}
		return &v.BingSafetyPhishingURL
	}).(MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput)
}

func (o MSTIDataConnectorDataTypesPtrOutput) MicrosoftEmergingThreatFeed() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypes) *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed {
		if v == nil {
			return nil
		}
		return &v.MicrosoftEmergingThreatFeed
	}).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput)
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

func (i MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return i.ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput).ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(ctx)
}









type MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput
	ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput
}

type mstidataConnectorDataTypesBingSafetyPhishingURLPtrType MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs

func MSTIDataConnectorDataTypesBingSafetyPhishingURLPtr(v *MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs) MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrInput {
	return (*mstidataConnectorDataTypesBingSafetyPhishingURLPtrType)(v)
}

func (*mstidataConnectorDataTypesBingSafetyPhishingURLPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesBingSafetyPhishingURL)(nil)).Elem()
}

func (i *mstidataConnectorDataTypesBingSafetyPhishingURLPtrType) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return i.ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(context.Background())
}

func (i *mstidataConnectorDataTypesBingSafetyPhishingURLPtrType) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput)
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

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return o.ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(context.Background())
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MSTIDataConnectorDataTypesBingSafetyPhishingURL) *MSTIDataConnectorDataTypesBingSafetyPhishingURL {
		return &v
	}).(MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput)
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesBingSafetyPhishingURL) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesBingSafetyPhishingURL) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesBingSafetyPhishingURL)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput) ToMSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput) Elem() MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesBingSafetyPhishingURL) MSTIDataConnectorDataTypesBingSafetyPhishingURL {
		if v != nil {
			return *v
		}
		var ret MSTIDataConnectorDataTypesBingSafetyPhishingURL
		return ret
	}).(MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput)
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput) LookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesBingSafetyPhishingURL) *string {
		if v == nil {
			return nil
		}
		return &v.LookbackPeriod
	}).(pulumi.StringPtrOutput)
}

func (o MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesBingSafetyPhishingURL) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
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

func (i MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return i.ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput).ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx)
}









type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput
	ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput
}

type mstidataConnectorDataTypesMicrosoftEmergingThreatFeedPtrType MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs

func MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtr(v *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrInput {
	return (*mstidataConnectorDataTypesMicrosoftEmergingThreatFeedPtrType)(v)
}

func (*mstidataConnectorDataTypesMicrosoftEmergingThreatFeedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (i *mstidataConnectorDataTypesMicrosoftEmergingThreatFeedPtrType) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return i.ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Background())
}

func (i *mstidataConnectorDataTypesMicrosoftEmergingThreatFeedPtrType) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput)
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

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return o.ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Background())
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed {
		return &v
	}).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput)
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput) ToMSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput) Elem() MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed {
		if v != nil {
			return *v
		}
		var ret MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed
		return ret
	}).(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput)
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput) LookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) *string {
		if v == nil {
			return nil
		}
		return &v.LookbackPeriod
	}).(pulumi.StringPtrOutput)
}

func (o MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeed) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MSTIDataConnectorDataTypesResponse struct {
	BingSafetyPhishingURL       MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL       `pulumi:"bingSafetyPhishingURL"`
	MicrosoftEmergingThreatFeed MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed `pulumi:"microsoftEmergingThreatFeed"`
}





type MSTIDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesResponseOutput() MSTIDataConnectorDataTypesResponseOutput
	ToMSTIDataConnectorDataTypesResponseOutputWithContext(context.Context) MSTIDataConnectorDataTypesResponseOutput
}

type MSTIDataConnectorDataTypesResponseArgs struct {
	BingSafetyPhishingURL       MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLInput       `pulumi:"bingSafetyPhishingURL"`
	MicrosoftEmergingThreatFeed MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedInput `pulumi:"microsoftEmergingThreatFeed"`
}

func (MSTIDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i MSTIDataConnectorDataTypesResponseArgs) ToMSTIDataConnectorDataTypesResponseOutput() MSTIDataConnectorDataTypesResponseOutput {
	return i.ToMSTIDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesResponseArgs) ToMSTIDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseOutput)
}

func (i MSTIDataConnectorDataTypesResponseArgs) ToMSTIDataConnectorDataTypesResponsePtrOutput() MSTIDataConnectorDataTypesResponsePtrOutput {
	return i.ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesResponseArgs) ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseOutput).ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type MSTIDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesResponsePtrOutput() MSTIDataConnectorDataTypesResponsePtrOutput
	ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) MSTIDataConnectorDataTypesResponsePtrOutput
}

type mstidataConnectorDataTypesResponsePtrType MSTIDataConnectorDataTypesResponseArgs

func MSTIDataConnectorDataTypesResponsePtr(v *MSTIDataConnectorDataTypesResponseArgs) MSTIDataConnectorDataTypesResponsePtrInput {
	return (*mstidataConnectorDataTypesResponsePtrType)(v)
}

func (*mstidataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *mstidataConnectorDataTypesResponsePtrType) ToMSTIDataConnectorDataTypesResponsePtrOutput() MSTIDataConnectorDataTypesResponsePtrOutput {
	return i.ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *mstidataConnectorDataTypesResponsePtrType) ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponsePtrOutput)
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

func (o MSTIDataConnectorDataTypesResponseOutput) ToMSTIDataConnectorDataTypesResponsePtrOutput() MSTIDataConnectorDataTypesResponsePtrOutput {
	return o.ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o MSTIDataConnectorDataTypesResponseOutput) ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MSTIDataConnectorDataTypesResponse) *MSTIDataConnectorDataTypesResponse {
		return &v
	}).(MSTIDataConnectorDataTypesResponsePtrOutput)
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

type MSTIDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesResponsePtrOutput) ToMSTIDataConnectorDataTypesResponsePtrOutput() MSTIDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponsePtrOutput) ToMSTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponsePtrOutput) Elem() MSTIDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponse) MSTIDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret MSTIDataConnectorDataTypesResponse
		return ret
	}).(MSTIDataConnectorDataTypesResponseOutput)
}

func (o MSTIDataConnectorDataTypesResponsePtrOutput) BingSafetyPhishingURL() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponse) *MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL {
		if v == nil {
			return nil
		}
		return &v.BingSafetyPhishingURL
	}).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput)
}

func (o MSTIDataConnectorDataTypesResponsePtrOutput) MicrosoftEmergingThreatFeed() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponse) *MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed {
		if v == nil {
			return nil
		}
		return &v.MicrosoftEmergingThreatFeed
	}).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput)
}

type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL struct {
	LookbackPeriod string `pulumi:"lookbackPeriod"`
	State          string `pulumi:"state"`
}





type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput
	ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutputWithContext(context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput
}

type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs struct {
	LookbackPeriod pulumi.StringInput `pulumi:"lookbackPeriod"`
	State          pulumi.StringInput `pulumi:"state"`
}

func (MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL)(nil)).Elem()
}

func (i MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput {
	return i.ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput)
}

func (i MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return i.ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput).ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(ctx)
}









type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput
	ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput
}

type mstidataConnectorDataTypesResponseBingSafetyPhishingURLPtrType MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs

func MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtr(v *MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrInput {
	return (*mstidataConnectorDataTypesResponseBingSafetyPhishingURLPtrType)(v)
}

func (*mstidataConnectorDataTypesResponseBingSafetyPhishingURLPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL)(nil)).Elem()
}

func (i *mstidataConnectorDataTypesResponseBingSafetyPhishingURLPtrType) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return i.ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(context.Background())
}

func (i *mstidataConnectorDataTypesResponseBingSafetyPhishingURLPtrType) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput)
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

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return o.ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(context.Background())
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) *MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL {
		return &v
	}).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput)
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput) ToMSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput) Elem() MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL {
		if v != nil {
			return *v
		}
		var ret MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL
		return ret
	}).(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput)
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput) LookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) *string {
		if v == nil {
			return nil
		}
		return &v.LookbackPeriod
	}).(pulumi.StringPtrOutput)
}

func (o MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponseBingSafetyPhishingURL) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed struct {
	LookbackPeriod string `pulumi:"lookbackPeriod"`
	State          string `pulumi:"state"`
}





type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput
	ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutputWithContext(context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput
}

type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs struct {
	LookbackPeriod pulumi.StringInput `pulumi:"lookbackPeriod"`
	State          pulumi.StringInput `pulumi:"state"`
}

func (MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (i MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput {
	return i.ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput)
}

func (i MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return i.ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Background())
}

func (i MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput).ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx)
}









type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrInput interface {
	pulumi.Input

	ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput
	ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput
}

type mstidataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrType MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs

func MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtr(v *MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrInput {
	return (*mstidataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrType)(v)
}

func (*mstidataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (i *mstidataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrType) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return i.ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Background())
}

func (i *mstidataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrType) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput)
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

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return o.ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(context.Background())
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) *MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed {
		return &v
	}).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput)
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) LookbackPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) string { return v.LookbackPeriod }).(pulumi.StringOutput)
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) string { return v.State }).(pulumi.StringOutput)
}

type MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput struct{ *pulumi.OutputState }

func (MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed)(nil)).Elem()
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput) ToMSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutputWithContext(ctx context.Context) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput {
	return o
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput) Elem() MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed {
		if v != nil {
			return *v
		}
		var ret MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed
		return ret
	}).(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput)
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput) LookbackPeriod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) *string {
		if v == nil {
			return nil
		}
		return &v.LookbackPeriod
	}).(pulumi.StringPtrOutput)
}

func (o MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeed) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
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

func (i MTPDataConnectorDataTypesArgs) ToMTPDataConnectorDataTypesPtrOutput() MTPDataConnectorDataTypesPtrOutput {
	return i.ToMTPDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesArgs) ToMTPDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesOutput).ToMTPDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type MTPDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesPtrOutput() MTPDataConnectorDataTypesPtrOutput
	ToMTPDataConnectorDataTypesPtrOutputWithContext(context.Context) MTPDataConnectorDataTypesPtrOutput
}

type mtpdataConnectorDataTypesPtrType MTPDataConnectorDataTypesArgs

func MTPDataConnectorDataTypesPtr(v *MTPDataConnectorDataTypesArgs) MTPDataConnectorDataTypesPtrInput {
	return (*mtpdataConnectorDataTypesPtrType)(v)
}

func (*mtpdataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypes)(nil)).Elem()
}

func (i *mtpdataConnectorDataTypesPtrType) ToMTPDataConnectorDataTypesPtrOutput() MTPDataConnectorDataTypesPtrOutput {
	return i.ToMTPDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *mtpdataConnectorDataTypesPtrType) ToMTPDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesPtrOutput)
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

func (o MTPDataConnectorDataTypesOutput) ToMTPDataConnectorDataTypesPtrOutput() MTPDataConnectorDataTypesPtrOutput {
	return o.ToMTPDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o MTPDataConnectorDataTypesOutput) ToMTPDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MTPDataConnectorDataTypes) *MTPDataConnectorDataTypes {
		return &v
	}).(MTPDataConnectorDataTypesPtrOutput)
}

func (o MTPDataConnectorDataTypesOutput) Incidents() MTPDataConnectorDataTypesIncidentsOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypes) MTPDataConnectorDataTypesIncidents { return v.Incidents }).(MTPDataConnectorDataTypesIncidentsOutput)
}

type MTPDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypes)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesPtrOutput) ToMTPDataConnectorDataTypesPtrOutput() MTPDataConnectorDataTypesPtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesPtrOutput) ToMTPDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesPtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesPtrOutput) Elem() MTPDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypes) MTPDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret MTPDataConnectorDataTypes
		return ret
	}).(MTPDataConnectorDataTypesOutput)
}

func (o MTPDataConnectorDataTypesPtrOutput) Incidents() MTPDataConnectorDataTypesIncidentsPtrOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypes) *MTPDataConnectorDataTypesIncidents {
		if v == nil {
			return nil
		}
		return &v.Incidents
	}).(MTPDataConnectorDataTypesIncidentsPtrOutput)
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

func (i MTPDataConnectorDataTypesIncidentsArgs) ToMTPDataConnectorDataTypesIncidentsPtrOutput() MTPDataConnectorDataTypesIncidentsPtrOutput {
	return i.ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesIncidentsArgs) ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesIncidentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesIncidentsOutput).ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(ctx)
}









type MTPDataConnectorDataTypesIncidentsPtrInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesIncidentsPtrOutput() MTPDataConnectorDataTypesIncidentsPtrOutput
	ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(context.Context) MTPDataConnectorDataTypesIncidentsPtrOutput
}

type mtpdataConnectorDataTypesIncidentsPtrType MTPDataConnectorDataTypesIncidentsArgs

func MTPDataConnectorDataTypesIncidentsPtr(v *MTPDataConnectorDataTypesIncidentsArgs) MTPDataConnectorDataTypesIncidentsPtrInput {
	return (*mtpdataConnectorDataTypesIncidentsPtrType)(v)
}

func (*mtpdataConnectorDataTypesIncidentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypesIncidents)(nil)).Elem()
}

func (i *mtpdataConnectorDataTypesIncidentsPtrType) ToMTPDataConnectorDataTypesIncidentsPtrOutput() MTPDataConnectorDataTypesIncidentsPtrOutput {
	return i.ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(context.Background())
}

func (i *mtpdataConnectorDataTypesIncidentsPtrType) ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesIncidentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesIncidentsPtrOutput)
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

func (o MTPDataConnectorDataTypesIncidentsOutput) ToMTPDataConnectorDataTypesIncidentsPtrOutput() MTPDataConnectorDataTypesIncidentsPtrOutput {
	return o.ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(context.Background())
}

func (o MTPDataConnectorDataTypesIncidentsOutput) ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesIncidentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MTPDataConnectorDataTypesIncidents) *MTPDataConnectorDataTypesIncidents {
		return &v
	}).(MTPDataConnectorDataTypesIncidentsPtrOutput)
}

func (o MTPDataConnectorDataTypesIncidentsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypesIncidents) string { return v.State }).(pulumi.StringOutput)
}

type MTPDataConnectorDataTypesIncidentsPtrOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesIncidentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypesIncidents)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesIncidentsPtrOutput) ToMTPDataConnectorDataTypesIncidentsPtrOutput() MTPDataConnectorDataTypesIncidentsPtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesIncidentsPtrOutput) ToMTPDataConnectorDataTypesIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesIncidentsPtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesIncidentsPtrOutput) Elem() MTPDataConnectorDataTypesIncidentsOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypesIncidents) MTPDataConnectorDataTypesIncidents {
		if v != nil {
			return *v
		}
		var ret MTPDataConnectorDataTypesIncidents
		return ret
	}).(MTPDataConnectorDataTypesIncidentsOutput)
}

func (o MTPDataConnectorDataTypesIncidentsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypesIncidents) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type MTPDataConnectorDataTypesResponse struct {
	Incidents MTPDataConnectorDataTypesResponseIncidents `pulumi:"incidents"`
}





type MTPDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesResponseOutput() MTPDataConnectorDataTypesResponseOutput
	ToMTPDataConnectorDataTypesResponseOutputWithContext(context.Context) MTPDataConnectorDataTypesResponseOutput
}

type MTPDataConnectorDataTypesResponseArgs struct {
	Incidents MTPDataConnectorDataTypesResponseIncidentsInput `pulumi:"incidents"`
}

func (MTPDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i MTPDataConnectorDataTypesResponseArgs) ToMTPDataConnectorDataTypesResponseOutput() MTPDataConnectorDataTypesResponseOutput {
	return i.ToMTPDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesResponseArgs) ToMTPDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesResponseOutput)
}

func (i MTPDataConnectorDataTypesResponseArgs) ToMTPDataConnectorDataTypesResponsePtrOutput() MTPDataConnectorDataTypesResponsePtrOutput {
	return i.ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesResponseArgs) ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesResponseOutput).ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type MTPDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesResponsePtrOutput() MTPDataConnectorDataTypesResponsePtrOutput
	ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) MTPDataConnectorDataTypesResponsePtrOutput
}

type mtpdataConnectorDataTypesResponsePtrType MTPDataConnectorDataTypesResponseArgs

func MTPDataConnectorDataTypesResponsePtr(v *MTPDataConnectorDataTypesResponseArgs) MTPDataConnectorDataTypesResponsePtrInput {
	return (*mtpdataConnectorDataTypesResponsePtrType)(v)
}

func (*mtpdataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *mtpdataConnectorDataTypesResponsePtrType) ToMTPDataConnectorDataTypesResponsePtrOutput() MTPDataConnectorDataTypesResponsePtrOutput {
	return i.ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *mtpdataConnectorDataTypesResponsePtrType) ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesResponsePtrOutput)
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

func (o MTPDataConnectorDataTypesResponseOutput) ToMTPDataConnectorDataTypesResponsePtrOutput() MTPDataConnectorDataTypesResponsePtrOutput {
	return o.ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o MTPDataConnectorDataTypesResponseOutput) ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MTPDataConnectorDataTypesResponse) *MTPDataConnectorDataTypesResponse {
		return &v
	}).(MTPDataConnectorDataTypesResponsePtrOutput)
}

func (o MTPDataConnectorDataTypesResponseOutput) Incidents() MTPDataConnectorDataTypesResponseIncidentsOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypesResponse) MTPDataConnectorDataTypesResponseIncidents {
		return v.Incidents
	}).(MTPDataConnectorDataTypesResponseIncidentsOutput)
}

type MTPDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesResponsePtrOutput) ToMTPDataConnectorDataTypesResponsePtrOutput() MTPDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponsePtrOutput) ToMTPDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponsePtrOutput) Elem() MTPDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypesResponse) MTPDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret MTPDataConnectorDataTypesResponse
		return ret
	}).(MTPDataConnectorDataTypesResponseOutput)
}

func (o MTPDataConnectorDataTypesResponsePtrOutput) Incidents() MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypesResponse) *MTPDataConnectorDataTypesResponseIncidents {
		if v == nil {
			return nil
		}
		return &v.Incidents
	}).(MTPDataConnectorDataTypesResponseIncidentsPtrOutput)
}

type MTPDataConnectorDataTypesResponseIncidents struct {
	State string `pulumi:"state"`
}





type MTPDataConnectorDataTypesResponseIncidentsInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesResponseIncidentsOutput() MTPDataConnectorDataTypesResponseIncidentsOutput
	ToMTPDataConnectorDataTypesResponseIncidentsOutputWithContext(context.Context) MTPDataConnectorDataTypesResponseIncidentsOutput
}

type MTPDataConnectorDataTypesResponseIncidentsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (MTPDataConnectorDataTypesResponseIncidentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MTPDataConnectorDataTypesResponseIncidents)(nil)).Elem()
}

func (i MTPDataConnectorDataTypesResponseIncidentsArgs) ToMTPDataConnectorDataTypesResponseIncidentsOutput() MTPDataConnectorDataTypesResponseIncidentsOutput {
	return i.ToMTPDataConnectorDataTypesResponseIncidentsOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesResponseIncidentsArgs) ToMTPDataConnectorDataTypesResponseIncidentsOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseIncidentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesResponseIncidentsOutput)
}

func (i MTPDataConnectorDataTypesResponseIncidentsArgs) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutput() MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return i.ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(context.Background())
}

func (i MTPDataConnectorDataTypesResponseIncidentsArgs) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesResponseIncidentsOutput).ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(ctx)
}









type MTPDataConnectorDataTypesResponseIncidentsPtrInput interface {
	pulumi.Input

	ToMTPDataConnectorDataTypesResponseIncidentsPtrOutput() MTPDataConnectorDataTypesResponseIncidentsPtrOutput
	ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(context.Context) MTPDataConnectorDataTypesResponseIncidentsPtrOutput
}

type mtpdataConnectorDataTypesResponseIncidentsPtrType MTPDataConnectorDataTypesResponseIncidentsArgs

func MTPDataConnectorDataTypesResponseIncidentsPtr(v *MTPDataConnectorDataTypesResponseIncidentsArgs) MTPDataConnectorDataTypesResponseIncidentsPtrInput {
	return (*mtpdataConnectorDataTypesResponseIncidentsPtrType)(v)
}

func (*mtpdataConnectorDataTypesResponseIncidentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypesResponseIncidents)(nil)).Elem()
}

func (i *mtpdataConnectorDataTypesResponseIncidentsPtrType) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutput() MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return i.ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(context.Background())
}

func (i *mtpdataConnectorDataTypesResponseIncidentsPtrType) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MTPDataConnectorDataTypesResponseIncidentsPtrOutput)
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

func (o MTPDataConnectorDataTypesResponseIncidentsOutput) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutput() MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return o.ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(context.Background())
}

func (o MTPDataConnectorDataTypesResponseIncidentsOutput) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MTPDataConnectorDataTypesResponseIncidents) *MTPDataConnectorDataTypesResponseIncidents {
		return &v
	}).(MTPDataConnectorDataTypesResponseIncidentsPtrOutput)
}

func (o MTPDataConnectorDataTypesResponseIncidentsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v MTPDataConnectorDataTypesResponseIncidents) string { return v.State }).(pulumi.StringOutput)
}

type MTPDataConnectorDataTypesResponseIncidentsPtrOutput struct{ *pulumi.OutputState }

func (MTPDataConnectorDataTypesResponseIncidentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MTPDataConnectorDataTypesResponseIncidents)(nil)).Elem()
}

func (o MTPDataConnectorDataTypesResponseIncidentsPtrOutput) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutput() MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponseIncidentsPtrOutput) ToMTPDataConnectorDataTypesResponseIncidentsPtrOutputWithContext(ctx context.Context) MTPDataConnectorDataTypesResponseIncidentsPtrOutput {
	return o
}

func (o MTPDataConnectorDataTypesResponseIncidentsPtrOutput) Elem() MTPDataConnectorDataTypesResponseIncidentsOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypesResponseIncidents) MTPDataConnectorDataTypesResponseIncidents {
		if v != nil {
			return *v
		}
		var ret MTPDataConnectorDataTypesResponseIncidents
		return ret
	}).(MTPDataConnectorDataTypesResponseIncidentsOutput)
}

func (o MTPDataConnectorDataTypesResponseIncidentsPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MTPDataConnectorDataTypesResponseIncidents) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
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





type MetadataAuthorResponseInput interface {
	pulumi.Input

	ToMetadataAuthorResponseOutput() MetadataAuthorResponseOutput
	ToMetadataAuthorResponseOutputWithContext(context.Context) MetadataAuthorResponseOutput
}

type MetadataAuthorResponseArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Link  pulumi.StringPtrInput `pulumi:"link"`
	Name  pulumi.StringPtrInput `pulumi:"name"`
}

func (MetadataAuthorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataAuthorResponse)(nil)).Elem()
}

func (i MetadataAuthorResponseArgs) ToMetadataAuthorResponseOutput() MetadataAuthorResponseOutput {
	return i.ToMetadataAuthorResponseOutputWithContext(context.Background())
}

func (i MetadataAuthorResponseArgs) ToMetadataAuthorResponseOutputWithContext(ctx context.Context) MetadataAuthorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataAuthorResponseOutput)
}

func (i MetadataAuthorResponseArgs) ToMetadataAuthorResponsePtrOutput() MetadataAuthorResponsePtrOutput {
	return i.ToMetadataAuthorResponsePtrOutputWithContext(context.Background())
}

func (i MetadataAuthorResponseArgs) ToMetadataAuthorResponsePtrOutputWithContext(ctx context.Context) MetadataAuthorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataAuthorResponseOutput).ToMetadataAuthorResponsePtrOutputWithContext(ctx)
}









type MetadataAuthorResponsePtrInput interface {
	pulumi.Input

	ToMetadataAuthorResponsePtrOutput() MetadataAuthorResponsePtrOutput
	ToMetadataAuthorResponsePtrOutputWithContext(context.Context) MetadataAuthorResponsePtrOutput
}

type metadataAuthorResponsePtrType MetadataAuthorResponseArgs

func MetadataAuthorResponsePtr(v *MetadataAuthorResponseArgs) MetadataAuthorResponsePtrInput {
	return (*metadataAuthorResponsePtrType)(v)
}

func (*metadataAuthorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataAuthorResponse)(nil)).Elem()
}

func (i *metadataAuthorResponsePtrType) ToMetadataAuthorResponsePtrOutput() MetadataAuthorResponsePtrOutput {
	return i.ToMetadataAuthorResponsePtrOutputWithContext(context.Background())
}

func (i *metadataAuthorResponsePtrType) ToMetadataAuthorResponsePtrOutputWithContext(ctx context.Context) MetadataAuthorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataAuthorResponsePtrOutput)
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

func (o MetadataAuthorResponseOutput) ToMetadataAuthorResponsePtrOutput() MetadataAuthorResponsePtrOutput {
	return o.ToMetadataAuthorResponsePtrOutputWithContext(context.Background())
}

func (o MetadataAuthorResponseOutput) ToMetadataAuthorResponsePtrOutputWithContext(ctx context.Context) MetadataAuthorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataAuthorResponse) *MetadataAuthorResponse {
		return &v
	}).(MetadataAuthorResponsePtrOutput)
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





type MetadataCategoriesResponseInput interface {
	pulumi.Input

	ToMetadataCategoriesResponseOutput() MetadataCategoriesResponseOutput
	ToMetadataCategoriesResponseOutputWithContext(context.Context) MetadataCategoriesResponseOutput
}

type MetadataCategoriesResponseArgs struct {
	Domains   pulumi.StringArrayInput `pulumi:"domains"`
	Verticals pulumi.StringArrayInput `pulumi:"verticals"`
}

func (MetadataCategoriesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataCategoriesResponse)(nil)).Elem()
}

func (i MetadataCategoriesResponseArgs) ToMetadataCategoriesResponseOutput() MetadataCategoriesResponseOutput {
	return i.ToMetadataCategoriesResponseOutputWithContext(context.Background())
}

func (i MetadataCategoriesResponseArgs) ToMetadataCategoriesResponseOutputWithContext(ctx context.Context) MetadataCategoriesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataCategoriesResponseOutput)
}

func (i MetadataCategoriesResponseArgs) ToMetadataCategoriesResponsePtrOutput() MetadataCategoriesResponsePtrOutput {
	return i.ToMetadataCategoriesResponsePtrOutputWithContext(context.Background())
}

func (i MetadataCategoriesResponseArgs) ToMetadataCategoriesResponsePtrOutputWithContext(ctx context.Context) MetadataCategoriesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataCategoriesResponseOutput).ToMetadataCategoriesResponsePtrOutputWithContext(ctx)
}









type MetadataCategoriesResponsePtrInput interface {
	pulumi.Input

	ToMetadataCategoriesResponsePtrOutput() MetadataCategoriesResponsePtrOutput
	ToMetadataCategoriesResponsePtrOutputWithContext(context.Context) MetadataCategoriesResponsePtrOutput
}

type metadataCategoriesResponsePtrType MetadataCategoriesResponseArgs

func MetadataCategoriesResponsePtr(v *MetadataCategoriesResponseArgs) MetadataCategoriesResponsePtrInput {
	return (*metadataCategoriesResponsePtrType)(v)
}

func (*metadataCategoriesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataCategoriesResponse)(nil)).Elem()
}

func (i *metadataCategoriesResponsePtrType) ToMetadataCategoriesResponsePtrOutput() MetadataCategoriesResponsePtrOutput {
	return i.ToMetadataCategoriesResponsePtrOutputWithContext(context.Background())
}

func (i *metadataCategoriesResponsePtrType) ToMetadataCategoriesResponsePtrOutputWithContext(ctx context.Context) MetadataCategoriesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataCategoriesResponsePtrOutput)
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

func (o MetadataCategoriesResponseOutput) ToMetadataCategoriesResponsePtrOutput() MetadataCategoriesResponsePtrOutput {
	return o.ToMetadataCategoriesResponsePtrOutputWithContext(context.Background())
}

func (o MetadataCategoriesResponseOutput) ToMetadataCategoriesResponsePtrOutputWithContext(ctx context.Context) MetadataCategoriesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataCategoriesResponse) *MetadataCategoriesResponse {
		return &v
	}).(MetadataCategoriesResponsePtrOutput)
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





type MetadataDependenciesResponseInput interface {
	pulumi.Input

	ToMetadataDependenciesResponseOutput() MetadataDependenciesResponseOutput
	ToMetadataDependenciesResponseOutputWithContext(context.Context) MetadataDependenciesResponseOutput
}

type MetadataDependenciesResponseArgs struct {
	ContentId pulumi.StringPtrInput                  `pulumi:"contentId"`
	Criteria  MetadataDependenciesResponseArrayInput `pulumi:"criteria"`
	Kind      pulumi.StringPtrInput                  `pulumi:"kind"`
	Name      pulumi.StringPtrInput                  `pulumi:"name"`
	Operator  pulumi.StringPtrInput                  `pulumi:"operator"`
	Version   pulumi.StringPtrInput                  `pulumi:"version"`
}

func (MetadataDependenciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataDependenciesResponse)(nil)).Elem()
}

func (i MetadataDependenciesResponseArgs) ToMetadataDependenciesResponseOutput() MetadataDependenciesResponseOutput {
	return i.ToMetadataDependenciesResponseOutputWithContext(context.Background())
}

func (i MetadataDependenciesResponseArgs) ToMetadataDependenciesResponseOutputWithContext(ctx context.Context) MetadataDependenciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesResponseOutput)
}

func (i MetadataDependenciesResponseArgs) ToMetadataDependenciesResponsePtrOutput() MetadataDependenciesResponsePtrOutput {
	return i.ToMetadataDependenciesResponsePtrOutputWithContext(context.Background())
}

func (i MetadataDependenciesResponseArgs) ToMetadataDependenciesResponsePtrOutputWithContext(ctx context.Context) MetadataDependenciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesResponseOutput).ToMetadataDependenciesResponsePtrOutputWithContext(ctx)
}









type MetadataDependenciesResponsePtrInput interface {
	pulumi.Input

	ToMetadataDependenciesResponsePtrOutput() MetadataDependenciesResponsePtrOutput
	ToMetadataDependenciesResponsePtrOutputWithContext(context.Context) MetadataDependenciesResponsePtrOutput
}

type metadataDependenciesResponsePtrType MetadataDependenciesResponseArgs

func MetadataDependenciesResponsePtr(v *MetadataDependenciesResponseArgs) MetadataDependenciesResponsePtrInput {
	return (*metadataDependenciesResponsePtrType)(v)
}

func (*metadataDependenciesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataDependenciesResponse)(nil)).Elem()
}

func (i *metadataDependenciesResponsePtrType) ToMetadataDependenciesResponsePtrOutput() MetadataDependenciesResponsePtrOutput {
	return i.ToMetadataDependenciesResponsePtrOutputWithContext(context.Background())
}

func (i *metadataDependenciesResponsePtrType) ToMetadataDependenciesResponsePtrOutputWithContext(ctx context.Context) MetadataDependenciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesResponsePtrOutput)
}





type MetadataDependenciesResponseArrayInput interface {
	pulumi.Input

	ToMetadataDependenciesResponseArrayOutput() MetadataDependenciesResponseArrayOutput
	ToMetadataDependenciesResponseArrayOutputWithContext(context.Context) MetadataDependenciesResponseArrayOutput
}

type MetadataDependenciesResponseArray []MetadataDependenciesResponseInput

func (MetadataDependenciesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetadataDependenciesResponse)(nil)).Elem()
}

func (i MetadataDependenciesResponseArray) ToMetadataDependenciesResponseArrayOutput() MetadataDependenciesResponseArrayOutput {
	return i.ToMetadataDependenciesResponseArrayOutputWithContext(context.Background())
}

func (i MetadataDependenciesResponseArray) ToMetadataDependenciesResponseArrayOutputWithContext(ctx context.Context) MetadataDependenciesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataDependenciesResponseArrayOutput)
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

func (o MetadataDependenciesResponseOutput) ToMetadataDependenciesResponsePtrOutput() MetadataDependenciesResponsePtrOutput {
	return o.ToMetadataDependenciesResponsePtrOutputWithContext(context.Background())
}

func (o MetadataDependenciesResponseOutput) ToMetadataDependenciesResponsePtrOutputWithContext(ctx context.Context) MetadataDependenciesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataDependenciesResponse) *MetadataDependenciesResponse {
		return &v
	}).(MetadataDependenciesResponsePtrOutput)
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





type MetadataSourceResponseInput interface {
	pulumi.Input

	ToMetadataSourceResponseOutput() MetadataSourceResponseOutput
	ToMetadataSourceResponseOutputWithContext(context.Context) MetadataSourceResponseOutput
}

type MetadataSourceResponseArgs struct {
	Kind     pulumi.StringInput    `pulumi:"kind"`
	Name     pulumi.StringPtrInput `pulumi:"name"`
	SourceId pulumi.StringPtrInput `pulumi:"sourceId"`
}

func (MetadataSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSourceResponse)(nil)).Elem()
}

func (i MetadataSourceResponseArgs) ToMetadataSourceResponseOutput() MetadataSourceResponseOutput {
	return i.ToMetadataSourceResponseOutputWithContext(context.Background())
}

func (i MetadataSourceResponseArgs) ToMetadataSourceResponseOutputWithContext(ctx context.Context) MetadataSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSourceResponseOutput)
}

func (i MetadataSourceResponseArgs) ToMetadataSourceResponsePtrOutput() MetadataSourceResponsePtrOutput {
	return i.ToMetadataSourceResponsePtrOutputWithContext(context.Background())
}

func (i MetadataSourceResponseArgs) ToMetadataSourceResponsePtrOutputWithContext(ctx context.Context) MetadataSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSourceResponseOutput).ToMetadataSourceResponsePtrOutputWithContext(ctx)
}









type MetadataSourceResponsePtrInput interface {
	pulumi.Input

	ToMetadataSourceResponsePtrOutput() MetadataSourceResponsePtrOutput
	ToMetadataSourceResponsePtrOutputWithContext(context.Context) MetadataSourceResponsePtrOutput
}

type metadataSourceResponsePtrType MetadataSourceResponseArgs

func MetadataSourceResponsePtr(v *MetadataSourceResponseArgs) MetadataSourceResponsePtrInput {
	return (*metadataSourceResponsePtrType)(v)
}

func (*metadataSourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSourceResponse)(nil)).Elem()
}

func (i *metadataSourceResponsePtrType) ToMetadataSourceResponsePtrOutput() MetadataSourceResponsePtrOutput {
	return i.ToMetadataSourceResponsePtrOutputWithContext(context.Background())
}

func (i *metadataSourceResponsePtrType) ToMetadataSourceResponsePtrOutputWithContext(ctx context.Context) MetadataSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSourceResponsePtrOutput)
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

func (o MetadataSourceResponseOutput) ToMetadataSourceResponsePtrOutput() MetadataSourceResponsePtrOutput {
	return o.ToMetadataSourceResponsePtrOutputWithContext(context.Background())
}

func (o MetadataSourceResponseOutput) ToMetadataSourceResponsePtrOutputWithContext(ctx context.Context) MetadataSourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataSourceResponse) *MetadataSourceResponse {
		return &v
	}).(MetadataSourceResponsePtrOutput)
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





type MetadataSupportResponseInput interface {
	pulumi.Input

	ToMetadataSupportResponseOutput() MetadataSupportResponseOutput
	ToMetadataSupportResponseOutputWithContext(context.Context) MetadataSupportResponseOutput
}

type MetadataSupportResponseArgs struct {
	Email pulumi.StringPtrInput `pulumi:"email"`
	Link  pulumi.StringPtrInput `pulumi:"link"`
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Tier  pulumi.StringInput    `pulumi:"tier"`
}

func (MetadataSupportResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetadataSupportResponse)(nil)).Elem()
}

func (i MetadataSupportResponseArgs) ToMetadataSupportResponseOutput() MetadataSupportResponseOutput {
	return i.ToMetadataSupportResponseOutputWithContext(context.Background())
}

func (i MetadataSupportResponseArgs) ToMetadataSupportResponseOutputWithContext(ctx context.Context) MetadataSupportResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSupportResponseOutput)
}

func (i MetadataSupportResponseArgs) ToMetadataSupportResponsePtrOutput() MetadataSupportResponsePtrOutput {
	return i.ToMetadataSupportResponsePtrOutputWithContext(context.Background())
}

func (i MetadataSupportResponseArgs) ToMetadataSupportResponsePtrOutputWithContext(ctx context.Context) MetadataSupportResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSupportResponseOutput).ToMetadataSupportResponsePtrOutputWithContext(ctx)
}









type MetadataSupportResponsePtrInput interface {
	pulumi.Input

	ToMetadataSupportResponsePtrOutput() MetadataSupportResponsePtrOutput
	ToMetadataSupportResponsePtrOutputWithContext(context.Context) MetadataSupportResponsePtrOutput
}

type metadataSupportResponsePtrType MetadataSupportResponseArgs

func MetadataSupportResponsePtr(v *MetadataSupportResponseArgs) MetadataSupportResponsePtrInput {
	return (*metadataSupportResponsePtrType)(v)
}

func (*metadataSupportResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MetadataSupportResponse)(nil)).Elem()
}

func (i *metadataSupportResponsePtrType) ToMetadataSupportResponsePtrOutput() MetadataSupportResponsePtrOutput {
	return i.ToMetadataSupportResponsePtrOutputWithContext(context.Background())
}

func (i *metadataSupportResponsePtrType) ToMetadataSupportResponsePtrOutputWithContext(ctx context.Context) MetadataSupportResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetadataSupportResponsePtrOutput)
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

func (o MetadataSupportResponseOutput) ToMetadataSupportResponsePtrOutput() MetadataSupportResponsePtrOutput {
	return o.ToMetadataSupportResponsePtrOutputWithContext(context.Background())
}

func (o MetadataSupportResponseOutput) ToMetadataSupportResponsePtrOutputWithContext(ctx context.Context) MetadataSupportResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetadataSupportResponse) *MetadataSupportResponse {
		return &v
	}).(MetadataSupportResponsePtrOutput)
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

func (o OfficeDataConnectorDataTypesOutput) Exchange() OfficeDataConnectorDataTypesExchangeOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypesExchange { return v.Exchange }).(OfficeDataConnectorDataTypesExchangeOutput)
}

func (o OfficeDataConnectorDataTypesOutput) SharePoint() OfficeDataConnectorDataTypesSharePointOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypesSharePoint { return v.SharePoint }).(OfficeDataConnectorDataTypesSharePointOutput)
}

func (o OfficeDataConnectorDataTypesOutput) Teams() OfficeDataConnectorDataTypesTeamsOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypes) OfficeDataConnectorDataTypesTeams { return v.Teams }).(OfficeDataConnectorDataTypesTeamsOutput)
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
		return &v.Exchange
	}).(OfficeDataConnectorDataTypesExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesPtrOutput) SharePoint() OfficeDataConnectorDataTypesSharePointPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesSharePoint {
		if v == nil {
			return nil
		}
		return &v.SharePoint
	}).(OfficeDataConnectorDataTypesSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesPtrOutput) Teams() OfficeDataConnectorDataTypesTeamsPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypes) *OfficeDataConnectorDataTypesTeams {
		if v == nil {
			return nil
		}
		return &v.Teams
	}).(OfficeDataConnectorDataTypesTeamsPtrOutput)
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

func (o OfficeDataConnectorDataTypesExchangeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesExchange) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponse struct {
	Exchange   OfficeDataConnectorDataTypesResponseExchange   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesResponseSharePoint `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesResponseTeams      `pulumi:"teams"`
}





type OfficeDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseOutput() OfficeDataConnectorDataTypesResponseOutput
	ToOfficeDataConnectorDataTypesResponseOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseOutput
}

type OfficeDataConnectorDataTypesResponseArgs struct {
	Exchange   OfficeDataConnectorDataTypesResponseExchangeInput   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesResponseSharePointInput `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesResponseTeamsInput      `pulumi:"teams"`
}

func (OfficeDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesResponseArgs) ToOfficeDataConnectorDataTypesResponseOutput() OfficeDataConnectorDataTypesResponseOutput {
	return i.ToOfficeDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseArgs) ToOfficeDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseOutput)
}

func (i OfficeDataConnectorDataTypesResponseArgs) ToOfficeDataConnectorDataTypesResponsePtrOutput() OfficeDataConnectorDataTypesResponsePtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseArgs) ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseOutput).ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponsePtrOutput() OfficeDataConnectorDataTypesResponsePtrOutput
	ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponsePtrOutput
}

type officeDataConnectorDataTypesResponsePtrType OfficeDataConnectorDataTypesResponseArgs

func OfficeDataConnectorDataTypesResponsePtr(v *OfficeDataConnectorDataTypesResponseArgs) OfficeDataConnectorDataTypesResponsePtrInput {
	return (*officeDataConnectorDataTypesResponsePtrType)(v)
}

func (*officeDataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesResponsePtrType) ToOfficeDataConnectorDataTypesResponsePtrOutput() OfficeDataConnectorDataTypesResponsePtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesResponsePtrType) ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponsePtrOutput)
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

func (o OfficeDataConnectorDataTypesResponseOutput) ToOfficeDataConnectorDataTypesResponsePtrOutput() OfficeDataConnectorDataTypesResponsePtrOutput {
	return o.ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesResponseOutput) ToOfficeDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponse {
		return &v
	}).(OfficeDataConnectorDataTypesResponsePtrOutput)
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
		return &v.Exchange
	}).(OfficeDataConnectorDataTypesResponseExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) SharePoint() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseSharePoint {
		if v == nil {
			return nil
		}
		return &v.SharePoint
	}).(OfficeDataConnectorDataTypesResponseSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesResponsePtrOutput) Teams() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o.ApplyT(func(v *OfficeDataConnectorDataTypesResponse) *OfficeDataConnectorDataTypesResponseTeams {
		if v == nil {
			return nil
		}
		return &v.Teams
	}).(OfficeDataConnectorDataTypesResponseTeamsPtrOutput)
}

type OfficeDataConnectorDataTypesResponseExchange struct {
	State string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesResponseExchangeInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseExchangeOutput() OfficeDataConnectorDataTypesResponseExchangeOutput
	ToOfficeDataConnectorDataTypesResponseExchangeOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseExchangeOutput
}

type OfficeDataConnectorDataTypesResponseExchangeArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (OfficeDataConnectorDataTypesResponseExchangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponseExchange)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesResponseExchangeArgs) ToOfficeDataConnectorDataTypesResponseExchangeOutput() OfficeDataConnectorDataTypesResponseExchangeOutput {
	return i.ToOfficeDataConnectorDataTypesResponseExchangeOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseExchangeArgs) ToOfficeDataConnectorDataTypesResponseExchangeOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseExchangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseExchangeOutput)
}

func (i OfficeDataConnectorDataTypesResponseExchangeArgs) ToOfficeDataConnectorDataTypesResponseExchangePtrOutput() OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseExchangeArgs) ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseExchangeOutput).ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesResponseExchangePtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseExchangePtrOutput() OfficeDataConnectorDataTypesResponseExchangePtrOutput
	ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseExchangePtrOutput
}

type officeDataConnectorDataTypesResponseExchangePtrType OfficeDataConnectorDataTypesResponseExchangeArgs

func OfficeDataConnectorDataTypesResponseExchangePtr(v *OfficeDataConnectorDataTypesResponseExchangeArgs) OfficeDataConnectorDataTypesResponseExchangePtrInput {
	return (*officeDataConnectorDataTypesResponseExchangePtrType)(v)
}

func (*officeDataConnectorDataTypesResponseExchangePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponseExchange)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesResponseExchangePtrType) ToOfficeDataConnectorDataTypesResponseExchangePtrOutput() OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesResponseExchangePtrType) ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseExchangePtrOutput)
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

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) ToOfficeDataConnectorDataTypesResponseExchangePtrOutput() OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return o.ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) ToOfficeDataConnectorDataTypesResponseExchangePtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseExchangePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesResponseExchange) *OfficeDataConnectorDataTypesResponseExchange {
		return &v
	}).(OfficeDataConnectorDataTypesResponseExchangePtrOutput)
}

func (o OfficeDataConnectorDataTypesResponseExchangeOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseExchange) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseSharePoint struct {
	State string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesResponseSharePointInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseSharePointOutput() OfficeDataConnectorDataTypesResponseSharePointOutput
	ToOfficeDataConnectorDataTypesResponseSharePointOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseSharePointOutput
}

type OfficeDataConnectorDataTypesResponseSharePointArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (OfficeDataConnectorDataTypesResponseSharePointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponseSharePoint)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesResponseSharePointArgs) ToOfficeDataConnectorDataTypesResponseSharePointOutput() OfficeDataConnectorDataTypesResponseSharePointOutput {
	return i.ToOfficeDataConnectorDataTypesResponseSharePointOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseSharePointArgs) ToOfficeDataConnectorDataTypesResponseSharePointOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseSharePointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseSharePointOutput)
}

func (i OfficeDataConnectorDataTypesResponseSharePointArgs) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutput() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseSharePointArgs) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseSharePointOutput).ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesResponseSharePointPtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseSharePointPtrOutput() OfficeDataConnectorDataTypesResponseSharePointPtrOutput
	ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseSharePointPtrOutput
}

type officeDataConnectorDataTypesResponseSharePointPtrType OfficeDataConnectorDataTypesResponseSharePointArgs

func OfficeDataConnectorDataTypesResponseSharePointPtr(v *OfficeDataConnectorDataTypesResponseSharePointArgs) OfficeDataConnectorDataTypesResponseSharePointPtrInput {
	return (*officeDataConnectorDataTypesResponseSharePointPtrType)(v)
}

func (*officeDataConnectorDataTypesResponseSharePointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponseSharePoint)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesResponseSharePointPtrType) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutput() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesResponseSharePointPtrType) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseSharePointPtrOutput)
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

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutput() OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o.ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) ToOfficeDataConnectorDataTypesResponseSharePointPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseSharePointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesResponseSharePoint) *OfficeDataConnectorDataTypesResponseSharePoint {
		return &v
	}).(OfficeDataConnectorDataTypesResponseSharePointPtrOutput)
}

func (o OfficeDataConnectorDataTypesResponseSharePointOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseSharePoint) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type OfficeDataConnectorDataTypesResponseTeams struct {
	State string `pulumi:"state"`
}





type OfficeDataConnectorDataTypesResponseTeamsInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseTeamsOutput() OfficeDataConnectorDataTypesResponseTeamsOutput
	ToOfficeDataConnectorDataTypesResponseTeamsOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseTeamsOutput
}

type OfficeDataConnectorDataTypesResponseTeamsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (OfficeDataConnectorDataTypesResponseTeamsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*OfficeDataConnectorDataTypesResponseTeams)(nil)).Elem()
}

func (i OfficeDataConnectorDataTypesResponseTeamsArgs) ToOfficeDataConnectorDataTypesResponseTeamsOutput() OfficeDataConnectorDataTypesResponseTeamsOutput {
	return i.ToOfficeDataConnectorDataTypesResponseTeamsOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseTeamsArgs) ToOfficeDataConnectorDataTypesResponseTeamsOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseTeamsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseTeamsOutput)
}

func (i OfficeDataConnectorDataTypesResponseTeamsArgs) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutput() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(context.Background())
}

func (i OfficeDataConnectorDataTypesResponseTeamsArgs) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseTeamsOutput).ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(ctx)
}









type OfficeDataConnectorDataTypesResponseTeamsPtrInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseTeamsPtrOutput() OfficeDataConnectorDataTypesResponseTeamsPtrOutput
	ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseTeamsPtrOutput
}

type officeDataConnectorDataTypesResponseTeamsPtrType OfficeDataConnectorDataTypesResponseTeamsArgs

func OfficeDataConnectorDataTypesResponseTeamsPtr(v *OfficeDataConnectorDataTypesResponseTeamsArgs) OfficeDataConnectorDataTypesResponseTeamsPtrInput {
	return (*officeDataConnectorDataTypesResponseTeamsPtrType)(v)
}

func (*officeDataConnectorDataTypesResponseTeamsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**OfficeDataConnectorDataTypesResponseTeams)(nil)).Elem()
}

func (i *officeDataConnectorDataTypesResponseTeamsPtrType) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutput() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return i.ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(context.Background())
}

func (i *officeDataConnectorDataTypesResponseTeamsPtrType) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(OfficeDataConnectorDataTypesResponseTeamsPtrOutput)
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

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutput() OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o.ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(context.Background())
}

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) ToOfficeDataConnectorDataTypesResponseTeamsPtrOutputWithContext(ctx context.Context) OfficeDataConnectorDataTypesResponseTeamsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OfficeDataConnectorDataTypesResponseTeams) *OfficeDataConnectorDataTypesResponseTeams {
		return &v
	}).(OfficeDataConnectorDataTypesResponseTeamsPtrOutput)
}

func (o OfficeDataConnectorDataTypesResponseTeamsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesResponseTeams) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
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

func (o OfficeDataConnectorDataTypesSharePointOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesSharePoint) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
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

func (o OfficeDataConnectorDataTypesTeamsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v OfficeDataConnectorDataTypesTeams) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
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





type PermissionsResponseInput interface {
	pulumi.Input

	ToPermissionsResponseOutput() PermissionsResponseOutput
	ToPermissionsResponseOutputWithContext(context.Context) PermissionsResponseOutput
}

type PermissionsResponseArgs struct {
	Customs          PermissionsResponseCustomsArrayInput          `pulumi:"customs"`
	ResourceProvider PermissionsResponseResourceProviderArrayInput `pulumi:"resourceProvider"`
}

func (PermissionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponse)(nil)).Elem()
}

func (i PermissionsResponseArgs) ToPermissionsResponseOutput() PermissionsResponseOutput {
	return i.ToPermissionsResponseOutputWithContext(context.Background())
}

func (i PermissionsResponseArgs) ToPermissionsResponseOutputWithContext(ctx context.Context) PermissionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseOutput)
}

func (i PermissionsResponseArgs) ToPermissionsResponsePtrOutput() PermissionsResponsePtrOutput {
	return i.ToPermissionsResponsePtrOutputWithContext(context.Background())
}

func (i PermissionsResponseArgs) ToPermissionsResponsePtrOutputWithContext(ctx context.Context) PermissionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseOutput).ToPermissionsResponsePtrOutputWithContext(ctx)
}









type PermissionsResponsePtrInput interface {
	pulumi.Input

	ToPermissionsResponsePtrOutput() PermissionsResponsePtrOutput
	ToPermissionsResponsePtrOutputWithContext(context.Context) PermissionsResponsePtrOutput
}

type permissionsResponsePtrType PermissionsResponseArgs

func PermissionsResponsePtr(v *PermissionsResponseArgs) PermissionsResponsePtrInput {
	return (*permissionsResponsePtrType)(v)
}

func (*permissionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PermissionsResponse)(nil)).Elem()
}

func (i *permissionsResponsePtrType) ToPermissionsResponsePtrOutput() PermissionsResponsePtrOutput {
	return i.ToPermissionsResponsePtrOutputWithContext(context.Background())
}

func (i *permissionsResponsePtrType) ToPermissionsResponsePtrOutputWithContext(ctx context.Context) PermissionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponsePtrOutput)
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

func (o PermissionsResponseOutput) ToPermissionsResponsePtrOutput() PermissionsResponsePtrOutput {
	return o.ToPermissionsResponsePtrOutputWithContext(context.Background())
}

func (o PermissionsResponseOutput) ToPermissionsResponsePtrOutputWithContext(ctx context.Context) PermissionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PermissionsResponse) *PermissionsResponse {
		return &v
	}).(PermissionsResponsePtrOutput)
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





type PermissionsResponseCustomsInput interface {
	pulumi.Input

	ToPermissionsResponseCustomsOutput() PermissionsResponseCustomsOutput
	ToPermissionsResponseCustomsOutputWithContext(context.Context) PermissionsResponseCustomsOutput
}

type PermissionsResponseCustomsArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Name        pulumi.StringPtrInput `pulumi:"name"`
}

func (PermissionsResponseCustomsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponseCustoms)(nil)).Elem()
}

func (i PermissionsResponseCustomsArgs) ToPermissionsResponseCustomsOutput() PermissionsResponseCustomsOutput {
	return i.ToPermissionsResponseCustomsOutputWithContext(context.Background())
}

func (i PermissionsResponseCustomsArgs) ToPermissionsResponseCustomsOutputWithContext(ctx context.Context) PermissionsResponseCustomsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseCustomsOutput)
}





type PermissionsResponseCustomsArrayInput interface {
	pulumi.Input

	ToPermissionsResponseCustomsArrayOutput() PermissionsResponseCustomsArrayOutput
	ToPermissionsResponseCustomsArrayOutputWithContext(context.Context) PermissionsResponseCustomsArrayOutput
}

type PermissionsResponseCustomsArray []PermissionsResponseCustomsInput

func (PermissionsResponseCustomsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsResponseCustoms)(nil)).Elem()
}

func (i PermissionsResponseCustomsArray) ToPermissionsResponseCustomsArrayOutput() PermissionsResponseCustomsArrayOutput {
	return i.ToPermissionsResponseCustomsArrayOutputWithContext(context.Background())
}

func (i PermissionsResponseCustomsArray) ToPermissionsResponseCustomsArrayOutputWithContext(ctx context.Context) PermissionsResponseCustomsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseCustomsArrayOutput)
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





type PermissionsResponseResourceProviderInput interface {
	pulumi.Input

	ToPermissionsResponseResourceProviderOutput() PermissionsResponseResourceProviderOutput
	ToPermissionsResponseResourceProviderOutputWithContext(context.Context) PermissionsResponseResourceProviderOutput
}

type PermissionsResponseResourceProviderArgs struct {
	PermissionsDisplayText pulumi.StringPtrInput               `pulumi:"permissionsDisplayText"`
	Provider               pulumi.StringPtrInput               `pulumi:"provider"`
	ProviderDisplayName    pulumi.StringPtrInput               `pulumi:"providerDisplayName"`
	RequiredPermissions    RequiredPermissionsResponsePtrInput `pulumi:"requiredPermissions"`
	Scope                  pulumi.StringPtrInput               `pulumi:"scope"`
}

func (PermissionsResponseResourceProviderArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionsResponseResourceProvider)(nil)).Elem()
}

func (i PermissionsResponseResourceProviderArgs) ToPermissionsResponseResourceProviderOutput() PermissionsResponseResourceProviderOutput {
	return i.ToPermissionsResponseResourceProviderOutputWithContext(context.Background())
}

func (i PermissionsResponseResourceProviderArgs) ToPermissionsResponseResourceProviderOutputWithContext(ctx context.Context) PermissionsResponseResourceProviderOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseResourceProviderOutput)
}





type PermissionsResponseResourceProviderArrayInput interface {
	pulumi.Input

	ToPermissionsResponseResourceProviderArrayOutput() PermissionsResponseResourceProviderArrayOutput
	ToPermissionsResponseResourceProviderArrayOutputWithContext(context.Context) PermissionsResponseResourceProviderArrayOutput
}

type PermissionsResponseResourceProviderArray []PermissionsResponseResourceProviderInput

func (PermissionsResponseResourceProviderArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionsResponseResourceProvider)(nil)).Elem()
}

func (i PermissionsResponseResourceProviderArray) ToPermissionsResponseResourceProviderArrayOutput() PermissionsResponseResourceProviderArrayOutput {
	return i.ToPermissionsResponseResourceProviderArrayOutputWithContext(context.Background())
}

func (i PermissionsResponseResourceProviderArray) ToPermissionsResponseResourceProviderArrayOutputWithContext(ctx context.Context) PermissionsResponseResourceProviderArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionsResponseResourceProviderArrayOutput)
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

type RepoResponse struct {
	Branches []string `pulumi:"branches"`
	FullName *string  `pulumi:"fullName"`
	Url      *string  `pulumi:"url"`
}





type RepoResponseInput interface {
	pulumi.Input

	ToRepoResponseOutput() RepoResponseOutput
	ToRepoResponseOutputWithContext(context.Context) RepoResponseOutput
}

type RepoResponseArgs struct {
	Branches pulumi.StringArrayInput `pulumi:"branches"`
	FullName pulumi.StringPtrInput   `pulumi:"fullName"`
	Url      pulumi.StringPtrInput   `pulumi:"url"`
}

func (RepoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepoResponse)(nil)).Elem()
}

func (i RepoResponseArgs) ToRepoResponseOutput() RepoResponseOutput {
	return i.ToRepoResponseOutputWithContext(context.Background())
}

func (i RepoResponseArgs) ToRepoResponseOutputWithContext(ctx context.Context) RepoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepoResponseOutput)
}





type RepoResponseArrayInput interface {
	pulumi.Input

	ToRepoResponseArrayOutput() RepoResponseArrayOutput
	ToRepoResponseArrayOutputWithContext(context.Context) RepoResponseArrayOutput
}

type RepoResponseArray []RepoResponseInput

func (RepoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RepoResponse)(nil)).Elem()
}

func (i RepoResponseArray) ToRepoResponseArrayOutput() RepoResponseArrayOutput {
	return i.ToRepoResponseArrayOutputWithContext(context.Background())
}

func (i RepoResponseArray) ToRepoResponseArrayOutputWithContext(ctx context.Context) RepoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepoResponseArrayOutput)
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

func (i RepositoryArgs) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return i.ToRepositoryPtrOutputWithContext(context.Background())
}

func (i RepositoryArgs) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryOutput).ToRepositoryPtrOutputWithContext(ctx)
}









type RepositoryPtrInput interface {
	pulumi.Input

	ToRepositoryPtrOutput() RepositoryPtrOutput
	ToRepositoryPtrOutputWithContext(context.Context) RepositoryPtrOutput
}

type repositoryPtrType RepositoryArgs

func RepositoryPtr(v *RepositoryArgs) RepositoryPtrInput {
	return (*repositoryPtrType)(v)
}

func (*repositoryPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (i *repositoryPtrType) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return i.ToRepositoryPtrOutputWithContext(context.Background())
}

func (i *repositoryPtrType) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryPtrOutput)
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

func (o RepositoryOutput) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return o.ToRepositoryPtrOutputWithContext(context.Background())
}

func (o RepositoryOutput) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Repository) *Repository {
		return &v
	}).(RepositoryPtrOutput)
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

type RepositoryPtrOutput struct{ *pulumi.OutputState }

func (RepositoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Repository)(nil)).Elem()
}

func (o RepositoryPtrOutput) ToRepositoryPtrOutput() RepositoryPtrOutput {
	return o
}

func (o RepositoryPtrOutput) ToRepositoryPtrOutputWithContext(ctx context.Context) RepositoryPtrOutput {
	return o
}

func (o RepositoryPtrOutput) Elem() RepositoryOutput {
	return o.ApplyT(func(v *Repository) Repository {
		if v != nil {
			return *v
		}
		var ret Repository
		return ret
	}).(RepositoryOutput)
}

func (o RepositoryPtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryPtrOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentLogsUrl
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryPtrOutput) DisplayUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) *string {
		if v == nil {
			return nil
		}
		return v.DisplayUrl
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryPtrOutput) PathMapping() ContentPathMapArrayOutput {
	return o.ApplyT(func(v *Repository) []ContentPathMap {
		if v == nil {
			return nil
		}
		return v.PathMapping
	}).(ContentPathMapArrayOutput)
}

func (o RepositoryPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Repository) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type RepositoryResponse struct {
	Branch            *string                  `pulumi:"branch"`
	DeploymentLogsUrl *string                  `pulumi:"deploymentLogsUrl"`
	DisplayUrl        *string                  `pulumi:"displayUrl"`
	PathMapping       []ContentPathMapResponse `pulumi:"pathMapping"`
	Url               *string                  `pulumi:"url"`
}





type RepositoryResponseInput interface {
	pulumi.Input

	ToRepositoryResponseOutput() RepositoryResponseOutput
	ToRepositoryResponseOutputWithContext(context.Context) RepositoryResponseOutput
}

type RepositoryResponseArgs struct {
	Branch            pulumi.StringPtrInput            `pulumi:"branch"`
	DeploymentLogsUrl pulumi.StringPtrInput            `pulumi:"deploymentLogsUrl"`
	DisplayUrl        pulumi.StringPtrInput            `pulumi:"displayUrl"`
	PathMapping       ContentPathMapResponseArrayInput `pulumi:"pathMapping"`
	Url               pulumi.StringPtrInput            `pulumi:"url"`
}

func (RepositoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RepositoryResponse)(nil)).Elem()
}

func (i RepositoryResponseArgs) ToRepositoryResponseOutput() RepositoryResponseOutput {
	return i.ToRepositoryResponseOutputWithContext(context.Background())
}

func (i RepositoryResponseArgs) ToRepositoryResponseOutputWithContext(ctx context.Context) RepositoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryResponseOutput)
}

func (i RepositoryResponseArgs) ToRepositoryResponsePtrOutput() RepositoryResponsePtrOutput {
	return i.ToRepositoryResponsePtrOutputWithContext(context.Background())
}

func (i RepositoryResponseArgs) ToRepositoryResponsePtrOutputWithContext(ctx context.Context) RepositoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryResponseOutput).ToRepositoryResponsePtrOutputWithContext(ctx)
}









type RepositoryResponsePtrInput interface {
	pulumi.Input

	ToRepositoryResponsePtrOutput() RepositoryResponsePtrOutput
	ToRepositoryResponsePtrOutputWithContext(context.Context) RepositoryResponsePtrOutput
}

type repositoryResponsePtrType RepositoryResponseArgs

func RepositoryResponsePtr(v *RepositoryResponseArgs) RepositoryResponsePtrInput {
	return (*repositoryResponsePtrType)(v)
}

func (*repositoryResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryResponse)(nil)).Elem()
}

func (i *repositoryResponsePtrType) ToRepositoryResponsePtrOutput() RepositoryResponsePtrOutput {
	return i.ToRepositoryResponsePtrOutputWithContext(context.Background())
}

func (i *repositoryResponsePtrType) ToRepositoryResponsePtrOutputWithContext(ctx context.Context) RepositoryResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RepositoryResponsePtrOutput)
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

func (o RepositoryResponseOutput) ToRepositoryResponsePtrOutput() RepositoryResponsePtrOutput {
	return o.ToRepositoryResponsePtrOutputWithContext(context.Background())
}

func (o RepositoryResponseOutput) ToRepositoryResponsePtrOutputWithContext(ctx context.Context) RepositoryResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RepositoryResponse) *RepositoryResponse {
		return &v
	}).(RepositoryResponsePtrOutput)
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

type RepositoryResponsePtrOutput struct{ *pulumi.OutputState }

func (RepositoryResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RepositoryResponse)(nil)).Elem()
}

func (o RepositoryResponsePtrOutput) ToRepositoryResponsePtrOutput() RepositoryResponsePtrOutput {
	return o
}

func (o RepositoryResponsePtrOutput) ToRepositoryResponsePtrOutputWithContext(ctx context.Context) RepositoryResponsePtrOutput {
	return o
}

func (o RepositoryResponsePtrOutput) Elem() RepositoryResponseOutput {
	return o.ApplyT(func(v *RepositoryResponse) RepositoryResponse {
		if v != nil {
			return *v
		}
		var ret RepositoryResponse
		return ret
	}).(RepositoryResponseOutput)
}

func (o RepositoryResponsePtrOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.Branch
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryResponsePtrOutput) DeploymentLogsUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeploymentLogsUrl
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryResponsePtrOutput) DisplayUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.DisplayUrl
	}).(pulumi.StringPtrOutput)
}

func (o RepositoryResponsePtrOutput) PathMapping() ContentPathMapResponseArrayOutput {
	return o.ApplyT(func(v *RepositoryResponse) []ContentPathMapResponse {
		if v == nil {
			return nil
		}
		return v.PathMapping
	}).(ContentPathMapResponseArrayOutput)
}

func (o RepositoryResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RepositoryResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
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





type RequiredPermissionsResponseInput interface {
	pulumi.Input

	ToRequiredPermissionsResponseOutput() RequiredPermissionsResponseOutput
	ToRequiredPermissionsResponseOutputWithContext(context.Context) RequiredPermissionsResponseOutput
}

type RequiredPermissionsResponseArgs struct {
	Action pulumi.BoolPtrInput `pulumi:"action"`
	Delete pulumi.BoolPtrInput `pulumi:"delete"`
	Read   pulumi.BoolPtrInput `pulumi:"read"`
	Write  pulumi.BoolPtrInput `pulumi:"write"`
}

func (RequiredPermissionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequiredPermissionsResponse)(nil)).Elem()
}

func (i RequiredPermissionsResponseArgs) ToRequiredPermissionsResponseOutput() RequiredPermissionsResponseOutput {
	return i.ToRequiredPermissionsResponseOutputWithContext(context.Background())
}

func (i RequiredPermissionsResponseArgs) ToRequiredPermissionsResponseOutputWithContext(ctx context.Context) RequiredPermissionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredPermissionsResponseOutput)
}

func (i RequiredPermissionsResponseArgs) ToRequiredPermissionsResponsePtrOutput() RequiredPermissionsResponsePtrOutput {
	return i.ToRequiredPermissionsResponsePtrOutputWithContext(context.Background())
}

func (i RequiredPermissionsResponseArgs) ToRequiredPermissionsResponsePtrOutputWithContext(ctx context.Context) RequiredPermissionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredPermissionsResponseOutput).ToRequiredPermissionsResponsePtrOutputWithContext(ctx)
}









type RequiredPermissionsResponsePtrInput interface {
	pulumi.Input

	ToRequiredPermissionsResponsePtrOutput() RequiredPermissionsResponsePtrOutput
	ToRequiredPermissionsResponsePtrOutputWithContext(context.Context) RequiredPermissionsResponsePtrOutput
}

type requiredPermissionsResponsePtrType RequiredPermissionsResponseArgs

func RequiredPermissionsResponsePtr(v *RequiredPermissionsResponseArgs) RequiredPermissionsResponsePtrInput {
	return (*requiredPermissionsResponsePtrType)(v)
}

func (*requiredPermissionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequiredPermissionsResponse)(nil)).Elem()
}

func (i *requiredPermissionsResponsePtrType) ToRequiredPermissionsResponsePtrOutput() RequiredPermissionsResponsePtrOutput {
	return i.ToRequiredPermissionsResponsePtrOutputWithContext(context.Background())
}

func (i *requiredPermissionsResponsePtrType) ToRequiredPermissionsResponsePtrOutputWithContext(ctx context.Context) RequiredPermissionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequiredPermissionsResponsePtrOutput)
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

func (o RequiredPermissionsResponseOutput) ToRequiredPermissionsResponsePtrOutput() RequiredPermissionsResponsePtrOutput {
	return o.ToRequiredPermissionsResponsePtrOutputWithContext(context.Background())
}

func (o RequiredPermissionsResponseOutput) ToRequiredPermissionsResponsePtrOutputWithContext(ctx context.Context) RequiredPermissionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequiredPermissionsResponse) *RequiredPermissionsResponse {
		return &v
	}).(RequiredPermissionsResponsePtrOutput)
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

func (o TIDataConnectorDataTypesOutput) Indicators() TIDataConnectorDataTypesIndicatorsOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypes) TIDataConnectorDataTypesIndicators { return v.Indicators }).(TIDataConnectorDataTypesIndicatorsOutput)
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
		return &v.Indicators
	}).(TIDataConnectorDataTypesIndicatorsPtrOutput)
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

func (o TIDataConnectorDataTypesIndicatorsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesIndicators) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type TIDataConnectorDataTypesResponse struct {
	Indicators TIDataConnectorDataTypesResponseIndicators `pulumi:"indicators"`
}





type TIDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesResponseOutput() TIDataConnectorDataTypesResponseOutput
	ToTIDataConnectorDataTypesResponseOutputWithContext(context.Context) TIDataConnectorDataTypesResponseOutput
}

type TIDataConnectorDataTypesResponseArgs struct {
	Indicators TIDataConnectorDataTypesResponseIndicatorsInput `pulumi:"indicators"`
}

func (TIDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i TIDataConnectorDataTypesResponseArgs) ToTIDataConnectorDataTypesResponseOutput() TIDataConnectorDataTypesResponseOutput {
	return i.ToTIDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesResponseArgs) ToTIDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesResponseOutput)
}

func (i TIDataConnectorDataTypesResponseArgs) ToTIDataConnectorDataTypesResponsePtrOutput() TIDataConnectorDataTypesResponsePtrOutput {
	return i.ToTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesResponseArgs) ToTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesResponseOutput).ToTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type TIDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesResponsePtrOutput() TIDataConnectorDataTypesResponsePtrOutput
	ToTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) TIDataConnectorDataTypesResponsePtrOutput
}

type tidataConnectorDataTypesResponsePtrType TIDataConnectorDataTypesResponseArgs

func TIDataConnectorDataTypesResponsePtr(v *TIDataConnectorDataTypesResponseArgs) TIDataConnectorDataTypesResponsePtrInput {
	return (*tidataConnectorDataTypesResponsePtrType)(v)
}

func (*tidataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *tidataConnectorDataTypesResponsePtrType) ToTIDataConnectorDataTypesResponsePtrOutput() TIDataConnectorDataTypesResponsePtrOutput {
	return i.ToTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *tidataConnectorDataTypesResponsePtrType) ToTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesResponsePtrOutput)
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

func (o TIDataConnectorDataTypesResponseOutput) ToTIDataConnectorDataTypesResponsePtrOutput() TIDataConnectorDataTypesResponsePtrOutput {
	return o.ToTIDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o TIDataConnectorDataTypesResponseOutput) ToTIDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TIDataConnectorDataTypesResponse) *TIDataConnectorDataTypesResponse {
		return &v
	}).(TIDataConnectorDataTypesResponsePtrOutput)
}

func (o TIDataConnectorDataTypesResponseOutput) Indicators() TIDataConnectorDataTypesResponseIndicatorsOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesResponse) TIDataConnectorDataTypesResponseIndicators {
		return v.Indicators
	}).(TIDataConnectorDataTypesResponseIndicatorsOutput)
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
		return &v.Indicators
	}).(TIDataConnectorDataTypesResponseIndicatorsPtrOutput)
}

type TIDataConnectorDataTypesResponseIndicators struct {
	State string `pulumi:"state"`
}





type TIDataConnectorDataTypesResponseIndicatorsInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesResponseIndicatorsOutput() TIDataConnectorDataTypesResponseIndicatorsOutput
	ToTIDataConnectorDataTypesResponseIndicatorsOutputWithContext(context.Context) TIDataConnectorDataTypesResponseIndicatorsOutput
}

type TIDataConnectorDataTypesResponseIndicatorsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (TIDataConnectorDataTypesResponseIndicatorsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TIDataConnectorDataTypesResponseIndicators)(nil)).Elem()
}

func (i TIDataConnectorDataTypesResponseIndicatorsArgs) ToTIDataConnectorDataTypesResponseIndicatorsOutput() TIDataConnectorDataTypesResponseIndicatorsOutput {
	return i.ToTIDataConnectorDataTypesResponseIndicatorsOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesResponseIndicatorsArgs) ToTIDataConnectorDataTypesResponseIndicatorsOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseIndicatorsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesResponseIndicatorsOutput)
}

func (i TIDataConnectorDataTypesResponseIndicatorsArgs) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutput() TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return i.ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(context.Background())
}

func (i TIDataConnectorDataTypesResponseIndicatorsArgs) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesResponseIndicatorsOutput).ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(ctx)
}









type TIDataConnectorDataTypesResponseIndicatorsPtrInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesResponseIndicatorsPtrOutput() TIDataConnectorDataTypesResponseIndicatorsPtrOutput
	ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(context.Context) TIDataConnectorDataTypesResponseIndicatorsPtrOutput
}

type tidataConnectorDataTypesResponseIndicatorsPtrType TIDataConnectorDataTypesResponseIndicatorsArgs

func TIDataConnectorDataTypesResponseIndicatorsPtr(v *TIDataConnectorDataTypesResponseIndicatorsArgs) TIDataConnectorDataTypesResponseIndicatorsPtrInput {
	return (*tidataConnectorDataTypesResponseIndicatorsPtrType)(v)
}

func (*tidataConnectorDataTypesResponseIndicatorsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TIDataConnectorDataTypesResponseIndicators)(nil)).Elem()
}

func (i *tidataConnectorDataTypesResponseIndicatorsPtrType) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutput() TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return i.ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(context.Background())
}

func (i *tidataConnectorDataTypesResponseIndicatorsPtrType) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TIDataConnectorDataTypesResponseIndicatorsPtrOutput)
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

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutput() TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return o.ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(context.Background())
}

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) ToTIDataConnectorDataTypesResponseIndicatorsPtrOutputWithContext(ctx context.Context) TIDataConnectorDataTypesResponseIndicatorsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TIDataConnectorDataTypesResponseIndicators) *TIDataConnectorDataTypesResponseIndicators {
		return &v
	}).(TIDataConnectorDataTypesResponseIndicatorsPtrOutput)
}

func (o TIDataConnectorDataTypesResponseIndicatorsOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TIDataConnectorDataTypesResponseIndicators) string { return v.State }).(pulumi.StringOutput)
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
		return &v.State
	}).(pulumi.StringPtrOutput)
}

type TeamInformationResponse struct {
	Description         string `pulumi:"description"`
	Name                string `pulumi:"name"`
	PrimaryChannelUrl   string `pulumi:"primaryChannelUrl"`
	TeamCreationTimeUtc string `pulumi:"teamCreationTimeUtc"`
	TeamId              string `pulumi:"teamId"`
}





type TeamInformationResponseInput interface {
	pulumi.Input

	ToTeamInformationResponseOutput() TeamInformationResponseOutput
	ToTeamInformationResponseOutputWithContext(context.Context) TeamInformationResponseOutput
}

type TeamInformationResponseArgs struct {
	Description         pulumi.StringInput `pulumi:"description"`
	Name                pulumi.StringInput `pulumi:"name"`
	PrimaryChannelUrl   pulumi.StringInput `pulumi:"primaryChannelUrl"`
	TeamCreationTimeUtc pulumi.StringInput `pulumi:"teamCreationTimeUtc"`
	TeamId              pulumi.StringInput `pulumi:"teamId"`
}

func (TeamInformationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TeamInformationResponse)(nil)).Elem()
}

func (i TeamInformationResponseArgs) ToTeamInformationResponseOutput() TeamInformationResponseOutput {
	return i.ToTeamInformationResponseOutputWithContext(context.Background())
}

func (i TeamInformationResponseArgs) ToTeamInformationResponseOutputWithContext(ctx context.Context) TeamInformationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamInformationResponseOutput)
}

func (i TeamInformationResponseArgs) ToTeamInformationResponsePtrOutput() TeamInformationResponsePtrOutput {
	return i.ToTeamInformationResponsePtrOutputWithContext(context.Background())
}

func (i TeamInformationResponseArgs) ToTeamInformationResponsePtrOutputWithContext(ctx context.Context) TeamInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamInformationResponseOutput).ToTeamInformationResponsePtrOutputWithContext(ctx)
}









type TeamInformationResponsePtrInput interface {
	pulumi.Input

	ToTeamInformationResponsePtrOutput() TeamInformationResponsePtrOutput
	ToTeamInformationResponsePtrOutputWithContext(context.Context) TeamInformationResponsePtrOutput
}

type teamInformationResponsePtrType TeamInformationResponseArgs

func TeamInformationResponsePtr(v *TeamInformationResponseArgs) TeamInformationResponsePtrInput {
	return (*teamInformationResponsePtrType)(v)
}

func (*teamInformationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TeamInformationResponse)(nil)).Elem()
}

func (i *teamInformationResponsePtrType) ToTeamInformationResponsePtrOutput() TeamInformationResponsePtrOutput {
	return i.ToTeamInformationResponsePtrOutputWithContext(context.Background())
}

func (i *teamInformationResponsePtrType) ToTeamInformationResponsePtrOutputWithContext(ctx context.Context) TeamInformationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TeamInformationResponsePtrOutput)
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

func (o TeamInformationResponseOutput) ToTeamInformationResponsePtrOutput() TeamInformationResponsePtrOutput {
	return o.ToTeamInformationResponsePtrOutputWithContext(context.Background())
}

func (o TeamInformationResponseOutput) ToTeamInformationResponsePtrOutputWithContext(ctx context.Context) TeamInformationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TeamInformationResponse) *TeamInformationResponse {
		return &v
	}).(TeamInformationResponsePtrOutput)
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

func (i TiTaxiiDataConnectorDataTypesArgs) ToTiTaxiiDataConnectorDataTypesPtrOutput() TiTaxiiDataConnectorDataTypesPtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesArgs) ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesOutput).ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(ctx)
}









type TiTaxiiDataConnectorDataTypesPtrInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesPtrOutput() TiTaxiiDataConnectorDataTypesPtrOutput
	ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesPtrOutput
}

type tiTaxiiDataConnectorDataTypesPtrType TiTaxiiDataConnectorDataTypesArgs

func TiTaxiiDataConnectorDataTypesPtr(v *TiTaxiiDataConnectorDataTypesArgs) TiTaxiiDataConnectorDataTypesPtrInput {
	return (*tiTaxiiDataConnectorDataTypesPtrType)(v)
}

func (*tiTaxiiDataConnectorDataTypesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypes)(nil)).Elem()
}

func (i *tiTaxiiDataConnectorDataTypesPtrType) ToTiTaxiiDataConnectorDataTypesPtrOutput() TiTaxiiDataConnectorDataTypesPtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (i *tiTaxiiDataConnectorDataTypesPtrType) ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesPtrOutput)
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

func (o TiTaxiiDataConnectorDataTypesOutput) ToTiTaxiiDataConnectorDataTypesPtrOutput() TiTaxiiDataConnectorDataTypesPtrOutput {
	return o.ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(context.Background())
}

func (o TiTaxiiDataConnectorDataTypesOutput) ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TiTaxiiDataConnectorDataTypes) *TiTaxiiDataConnectorDataTypes {
		return &v
	}).(TiTaxiiDataConnectorDataTypesPtrOutput)
}

func (o TiTaxiiDataConnectorDataTypesOutput) TaxiiClient() TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypes) TiTaxiiDataConnectorDataTypesTaxiiClient { return v.TaxiiClient }).(TiTaxiiDataConnectorDataTypesTaxiiClientOutput)
}

type TiTaxiiDataConnectorDataTypesPtrOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypes)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesPtrOutput) ToTiTaxiiDataConnectorDataTypesPtrOutput() TiTaxiiDataConnectorDataTypesPtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesPtrOutput) ToTiTaxiiDataConnectorDataTypesPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesPtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesPtrOutput) Elem() TiTaxiiDataConnectorDataTypesOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypes) TiTaxiiDataConnectorDataTypes {
		if v != nil {
			return *v
		}
		var ret TiTaxiiDataConnectorDataTypes
		return ret
	}).(TiTaxiiDataConnectorDataTypesOutput)
}

func (o TiTaxiiDataConnectorDataTypesPtrOutput) TaxiiClient() TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypes) *TiTaxiiDataConnectorDataTypesTaxiiClient {
		if v == nil {
			return nil
		}
		return &v.TaxiiClient
	}).(TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput)
}

type TiTaxiiDataConnectorDataTypesResponse struct {
	TaxiiClient TiTaxiiDataConnectorDataTypesResponseTaxiiClient `pulumi:"taxiiClient"`
}





type TiTaxiiDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesResponseOutput() TiTaxiiDataConnectorDataTypesResponseOutput
	ToTiTaxiiDataConnectorDataTypesResponseOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesResponseOutput
}

type TiTaxiiDataConnectorDataTypesResponseArgs struct {
	TaxiiClient TiTaxiiDataConnectorDataTypesResponseTaxiiClientInput `pulumi:"taxiiClient"`
}

func (TiTaxiiDataConnectorDataTypesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i TiTaxiiDataConnectorDataTypesResponseArgs) ToTiTaxiiDataConnectorDataTypesResponseOutput() TiTaxiiDataConnectorDataTypesResponseOutput {
	return i.ToTiTaxiiDataConnectorDataTypesResponseOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesResponseArgs) ToTiTaxiiDataConnectorDataTypesResponseOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesResponseOutput)
}

func (i TiTaxiiDataConnectorDataTypesResponseArgs) ToTiTaxiiDataConnectorDataTypesResponsePtrOutput() TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesResponseArgs) ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesResponseOutput).ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(ctx)
}









type TiTaxiiDataConnectorDataTypesResponsePtrInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesResponsePtrOutput() TiTaxiiDataConnectorDataTypesResponsePtrOutput
	ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesResponsePtrOutput
}

type tiTaxiiDataConnectorDataTypesResponsePtrType TiTaxiiDataConnectorDataTypesResponseArgs

func TiTaxiiDataConnectorDataTypesResponsePtr(v *TiTaxiiDataConnectorDataTypesResponseArgs) TiTaxiiDataConnectorDataTypesResponsePtrInput {
	return (*tiTaxiiDataConnectorDataTypesResponsePtrType)(v)
}

func (*tiTaxiiDataConnectorDataTypesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypesResponse)(nil)).Elem()
}

func (i *tiTaxiiDataConnectorDataTypesResponsePtrType) ToTiTaxiiDataConnectorDataTypesResponsePtrOutput() TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (i *tiTaxiiDataConnectorDataTypesResponsePtrType) ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesResponsePtrOutput)
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

func (o TiTaxiiDataConnectorDataTypesResponseOutput) ToTiTaxiiDataConnectorDataTypesResponsePtrOutput() TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return o.ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(context.Background())
}

func (o TiTaxiiDataConnectorDataTypesResponseOutput) ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TiTaxiiDataConnectorDataTypesResponse) *TiTaxiiDataConnectorDataTypesResponse {
		return &v
	}).(TiTaxiiDataConnectorDataTypesResponsePtrOutput)
}

func (o TiTaxiiDataConnectorDataTypesResponseOutput) TaxiiClient() TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypesResponse) TiTaxiiDataConnectorDataTypesResponseTaxiiClient {
		return v.TaxiiClient
	}).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput)
}

type TiTaxiiDataConnectorDataTypesResponsePtrOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypesResponse)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesResponsePtrOutput) ToTiTaxiiDataConnectorDataTypesResponsePtrOutput() TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponsePtrOutput) ToTiTaxiiDataConnectorDataTypesResponsePtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponsePtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponsePtrOutput) Elem() TiTaxiiDataConnectorDataTypesResponseOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypesResponse) TiTaxiiDataConnectorDataTypesResponse {
		if v != nil {
			return *v
		}
		var ret TiTaxiiDataConnectorDataTypesResponse
		return ret
	}).(TiTaxiiDataConnectorDataTypesResponseOutput)
}

func (o TiTaxiiDataConnectorDataTypesResponsePtrOutput) TaxiiClient() TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypesResponse) *TiTaxiiDataConnectorDataTypesResponseTaxiiClient {
		if v == nil {
			return nil
		}
		return &v.TaxiiClient
	}).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput)
}

type TiTaxiiDataConnectorDataTypesResponseTaxiiClient struct {
	State string `pulumi:"state"`
}





type TiTaxiiDataConnectorDataTypesResponseTaxiiClientInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput
	ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput
}

type TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponseTaxiiClient)(nil)).Elem()
}

func (i TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return i.ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput)
}

func (i TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput).ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(ctx)
}









type TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput
	ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput
}

type tiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrType TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs

func TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtr(v *TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs) TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrInput {
	return (*tiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrType)(v)
}

func (*tiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypesResponseTaxiiClient)(nil)).Elem()
}

func (i *tiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrType) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(context.Background())
}

func (i *tiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrType) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput)
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

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return o.ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(context.Background())
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TiTaxiiDataConnectorDataTypesResponseTaxiiClient) *TiTaxiiDataConnectorDataTypesResponseTaxiiClient {
		return &v
	}).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput)
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypesResponseTaxiiClient) string { return v.State }).(pulumi.StringOutput)
}

type TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypesResponseTaxiiClient)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput) ToTiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput) Elem() TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypesResponseTaxiiClient) TiTaxiiDataConnectorDataTypesResponseTaxiiClient {
		if v != nil {
			return *v
		}
		var ret TiTaxiiDataConnectorDataTypesResponseTaxiiClient
		return ret
	}).(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput)
}

func (o TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypesResponseTaxiiClient) *string {
		if v == nil {
			return nil
		}
		return &v.State
	}).(pulumi.StringPtrOutput)
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

func (i TiTaxiiDataConnectorDataTypesTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(context.Background())
}

func (i TiTaxiiDataConnectorDataTypesTaxiiClientArgs) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesTaxiiClientOutput).ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(ctx)
}









type TiTaxiiDataConnectorDataTypesTaxiiClientPtrInput interface {
	pulumi.Input

	ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput
	ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput
}

type tiTaxiiDataConnectorDataTypesTaxiiClientPtrType TiTaxiiDataConnectorDataTypesTaxiiClientArgs

func TiTaxiiDataConnectorDataTypesTaxiiClientPtr(v *TiTaxiiDataConnectorDataTypesTaxiiClientArgs) TiTaxiiDataConnectorDataTypesTaxiiClientPtrInput {
	return (*tiTaxiiDataConnectorDataTypesTaxiiClientPtrType)(v)
}

func (*tiTaxiiDataConnectorDataTypesTaxiiClientPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypesTaxiiClient)(nil)).Elem()
}

func (i *tiTaxiiDataConnectorDataTypesTaxiiClientPtrType) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return i.ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(context.Background())
}

func (i *tiTaxiiDataConnectorDataTypesTaxiiClientPtrType) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput)
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

func (o TiTaxiiDataConnectorDataTypesTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return o.ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(context.Background())
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientOutput) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TiTaxiiDataConnectorDataTypesTaxiiClient) *TiTaxiiDataConnectorDataTypesTaxiiClient {
		return &v
	}).(TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput)
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v TiTaxiiDataConnectorDataTypesTaxiiClient) string { return v.State }).(pulumi.StringOutput)
}

type TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput struct{ *pulumi.OutputState }

func (TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TiTaxiiDataConnectorDataTypesTaxiiClient)(nil)).Elem()
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput() TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput) ToTiTaxiiDataConnectorDataTypesTaxiiClientPtrOutputWithContext(ctx context.Context) TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput {
	return o
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput) Elem() TiTaxiiDataConnectorDataTypesTaxiiClientOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypesTaxiiClient) TiTaxiiDataConnectorDataTypesTaxiiClient {
		if v != nil {
			return *v
		}
		var ret TiTaxiiDataConnectorDataTypesTaxiiClient
		return ret
	}).(TiTaxiiDataConnectorDataTypesTaxiiClientOutput)
}

func (o TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TiTaxiiDataConnectorDataTypesTaxiiClient) *string {
		if v == nil {
			return nil
		}
		return &v.State
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





type WatchlistUserInfoResponseInput interface {
	pulumi.Input

	ToWatchlistUserInfoResponseOutput() WatchlistUserInfoResponseOutput
	ToWatchlistUserInfoResponseOutputWithContext(context.Context) WatchlistUserInfoResponseOutput
}

type WatchlistUserInfoResponseArgs struct {
	Email    pulumi.StringInput    `pulumi:"email"`
	Name     pulumi.StringInput    `pulumi:"name"`
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (WatchlistUserInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WatchlistUserInfoResponse)(nil)).Elem()
}

func (i WatchlistUserInfoResponseArgs) ToWatchlistUserInfoResponseOutput() WatchlistUserInfoResponseOutput {
	return i.ToWatchlistUserInfoResponseOutputWithContext(context.Background())
}

func (i WatchlistUserInfoResponseArgs) ToWatchlistUserInfoResponseOutputWithContext(ctx context.Context) WatchlistUserInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistUserInfoResponseOutput)
}

func (i WatchlistUserInfoResponseArgs) ToWatchlistUserInfoResponsePtrOutput() WatchlistUserInfoResponsePtrOutput {
	return i.ToWatchlistUserInfoResponsePtrOutputWithContext(context.Background())
}

func (i WatchlistUserInfoResponseArgs) ToWatchlistUserInfoResponsePtrOutputWithContext(ctx context.Context) WatchlistUserInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistUserInfoResponseOutput).ToWatchlistUserInfoResponsePtrOutputWithContext(ctx)
}









type WatchlistUserInfoResponsePtrInput interface {
	pulumi.Input

	ToWatchlistUserInfoResponsePtrOutput() WatchlistUserInfoResponsePtrOutput
	ToWatchlistUserInfoResponsePtrOutputWithContext(context.Context) WatchlistUserInfoResponsePtrOutput
}

type watchlistUserInfoResponsePtrType WatchlistUserInfoResponseArgs

func WatchlistUserInfoResponsePtr(v *WatchlistUserInfoResponseArgs) WatchlistUserInfoResponsePtrInput {
	return (*watchlistUserInfoResponsePtrType)(v)
}

func (*watchlistUserInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WatchlistUserInfoResponse)(nil)).Elem()
}

func (i *watchlistUserInfoResponsePtrType) ToWatchlistUserInfoResponsePtrOutput() WatchlistUserInfoResponsePtrOutput {
	return i.ToWatchlistUserInfoResponsePtrOutputWithContext(context.Background())
}

func (i *watchlistUserInfoResponsePtrType) ToWatchlistUserInfoResponsePtrOutputWithContext(ctx context.Context) WatchlistUserInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WatchlistUserInfoResponsePtrOutput)
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

func (o WatchlistUserInfoResponseOutput) ToWatchlistUserInfoResponsePtrOutput() WatchlistUserInfoResponsePtrOutput {
	return o.ToWatchlistUserInfoResponsePtrOutputWithContext(context.Background())
}

func (o WatchlistUserInfoResponseOutput) ToWatchlistUserInfoResponsePtrOutputWithContext(ctx context.Context) WatchlistUserInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WatchlistUserInfoResponse) *WatchlistUserInfoResponse {
		return &v
	}).(WatchlistUserInfoResponsePtrOutput)
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
	pulumi.RegisterOutputType(AvailabilityOutput{})
	pulumi.RegisterOutputType(AvailabilityPtrOutput{})
	pulumi.RegisterOutputType(AvailabilityResponseOutput{})
	pulumi.RegisterOutputType(AvailabilityResponsePtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput{})
	pulumi.RegisterOutputType(ClientInfoResponseOutput{})
	pulumi.RegisterOutputType(ClientInfoResponsePtrOutput{})
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
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput{})
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
	pulumi.RegisterOutputType(IncidentAdditionalDataResponsePtrOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationPtrOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationResponseOutput{})
	pulumi.RegisterOutputType(IncidentConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(IncidentLabelOutput{})
	pulumi.RegisterOutputType(IncidentLabelArrayOutput{})
	pulumi.RegisterOutputType(IncidentLabelResponseOutput{})
	pulumi.RegisterOutputType(IncidentLabelResponseArrayOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoPtrOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoResponseOutput{})
	pulumi.RegisterOutputType(IncidentOwnerInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(InstructionStepsInstructionsOutput{})
	pulumi.RegisterOutputType(InstructionStepsInstructionsArrayOutput{})
	pulumi.RegisterOutputType(InstructionStepsResponseInstructionsOutput{})
	pulumi.RegisterOutputType(InstructionStepsResponseInstructionsArrayOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesBingSafetyPhishingURLOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedOutput{})
	pulumi.RegisterOutputType(MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesIncidentsOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesIncidentsPtrOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesResponseIncidentsOutput{})
	pulumi.RegisterOutputType(MTPDataConnectorDataTypesResponseIncidentsPtrOutput{})
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
	pulumi.RegisterOutputType(RepositoryPtrOutput{})
	pulumi.RegisterOutputType(RepositoryResponseOutput{})
	pulumi.RegisterOutputType(RepositoryResponsePtrOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsPtrOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsResponseOutput{})
	pulumi.RegisterOutputType(RequiredPermissionsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesIndicatorsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesIndicatorsPtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseIndicatorsOutput{})
	pulumi.RegisterOutputType(TIDataConnectorDataTypesResponseIndicatorsPtrOutput{})
	pulumi.RegisterOutputType(TeamInformationResponseOutput{})
	pulumi.RegisterOutputType(TeamInformationResponsePtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesTaxiiClientOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoPtrOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponseOutput{})
	pulumi.RegisterOutputType(WatchlistUserInfoResponsePtrOutput{})
}
