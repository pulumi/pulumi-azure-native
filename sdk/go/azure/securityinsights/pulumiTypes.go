


package securityinsights

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





type ActivityTimelineItemResponseInput interface {
	pulumi.Input

	ToActivityTimelineItemResponseOutput() ActivityTimelineItemResponseOutput
	ToActivityTimelineItemResponseOutputWithContext(context.Context) ActivityTimelineItemResponseOutput
}

type ActivityTimelineItemResponseArgs struct {
	BucketEndTimeUTC     pulumi.StringInput `pulumi:"bucketEndTimeUTC"`
	BucketStartTimeUTC   pulumi.StringInput `pulumi:"bucketStartTimeUTC"`
	Content              pulumi.StringInput `pulumi:"content"`
	FirstActivityTimeUTC pulumi.StringInput `pulumi:"firstActivityTimeUTC"`
	Kind                 pulumi.StringInput `pulumi:"kind"`
	LastActivityTimeUTC  pulumi.StringInput `pulumi:"lastActivityTimeUTC"`
	QueryId              pulumi.StringInput `pulumi:"queryId"`
	Title                pulumi.StringInput `pulumi:"title"`
}

func (ActivityTimelineItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityTimelineItemResponse)(nil)).Elem()
}

func (i ActivityTimelineItemResponseArgs) ToActivityTimelineItemResponseOutput() ActivityTimelineItemResponseOutput {
	return i.ToActivityTimelineItemResponseOutputWithContext(context.Background())
}

func (i ActivityTimelineItemResponseArgs) ToActivityTimelineItemResponseOutputWithContext(ctx context.Context) ActivityTimelineItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivityTimelineItemResponseOutput)
}

type ActivityTimelineItemResponseOutput struct{ *pulumi.OutputState }

func (ActivityTimelineItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivityTimelineItemResponse)(nil)).Elem()
}

func (o ActivityTimelineItemResponseOutput) ToActivityTimelineItemResponseOutput() ActivityTimelineItemResponseOutput {
	return o
}

func (o ActivityTimelineItemResponseOutput) ToActivityTimelineItemResponseOutputWithContext(ctx context.Context) ActivityTimelineItemResponseOutput {
	return o
}

func (o ActivityTimelineItemResponseOutput) BucketEndTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.BucketEndTimeUTC }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) BucketStartTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.BucketStartTimeUTC }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.Content }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) FirstActivityTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.FirstActivityTimeUTC }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) LastActivityTimeUTC() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.LastActivityTimeUTC }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) QueryId() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.QueryId }).(pulumi.StringOutput)
}

func (o ActivityTimelineItemResponseOutput) Title() pulumi.StringOutput {
	return o.ApplyT(func(v ActivityTimelineItemResponse) string { return v.Title }).(pulumi.StringOutput)
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





type AlertsDataTypeOfDataConnectorResponseInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorResponseOutput() AlertsDataTypeOfDataConnectorResponseOutput
	ToAlertsDataTypeOfDataConnectorResponseOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorResponseOutput
}

type AlertsDataTypeOfDataConnectorResponseArgs struct {
	Alerts DataConnectorDataTypeCommonResponsePtrInput `pulumi:"alerts"`
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
	ActionConfiguration AutomationRuleModifyPropertiesActionActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                                  `pulumi:"actionType"`
	Order               int                                                     `pulumi:"order"`
}





type AutomationRuleModifyPropertiesActionInput interface {
	pulumi.Input

	ToAutomationRuleModifyPropertiesActionOutput() AutomationRuleModifyPropertiesActionOutput
	ToAutomationRuleModifyPropertiesActionOutputWithContext(context.Context) AutomationRuleModifyPropertiesActionOutput
}

type AutomationRuleModifyPropertiesActionArgs struct {
	ActionConfiguration AutomationRuleModifyPropertiesActionActionConfigurationInput `pulumi:"actionConfiguration"`
	ActionType          pulumi.StringInput                                           `pulumi:"actionType"`
	Order               pulumi.IntInput                                              `pulumi:"order"`
}

func (AutomationRuleModifyPropertiesActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesAction)(nil)).Elem()
}

func (i AutomationRuleModifyPropertiesActionArgs) ToAutomationRuleModifyPropertiesActionOutput() AutomationRuleModifyPropertiesActionOutput {
	return i.ToAutomationRuleModifyPropertiesActionOutputWithContext(context.Background())
}

func (i AutomationRuleModifyPropertiesActionArgs) ToAutomationRuleModifyPropertiesActionOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleModifyPropertiesActionOutput)
}

type AutomationRuleModifyPropertiesActionOutput struct{ *pulumi.OutputState }

func (AutomationRuleModifyPropertiesActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesAction)(nil)).Elem()
}

func (o AutomationRuleModifyPropertiesActionOutput) ToAutomationRuleModifyPropertiesActionOutput() AutomationRuleModifyPropertiesActionOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionOutput) ToAutomationRuleModifyPropertiesActionOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionOutput) ActionConfiguration() AutomationRuleModifyPropertiesActionActionConfigurationOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesAction) AutomationRuleModifyPropertiesActionActionConfiguration {
		return v.ActionConfiguration
	}).(AutomationRuleModifyPropertiesActionActionConfigurationOutput)
}

func (o AutomationRuleModifyPropertiesActionOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesAction) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationRuleModifyPropertiesActionOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesAction) int { return v.Order }).(pulumi.IntOutput)
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





type AutomationRuleModifyPropertiesActionActionConfigurationInput interface {
	pulumi.Input

	ToAutomationRuleModifyPropertiesActionActionConfigurationOutput() AutomationRuleModifyPropertiesActionActionConfigurationOutput
	ToAutomationRuleModifyPropertiesActionActionConfigurationOutputWithContext(context.Context) AutomationRuleModifyPropertiesActionActionConfigurationOutput
}

type AutomationRuleModifyPropertiesActionActionConfigurationArgs struct {
	Classification        pulumi.StringPtrInput     `pulumi:"classification"`
	ClassificationComment pulumi.StringPtrInput     `pulumi:"classificationComment"`
	ClassificationReason  pulumi.StringPtrInput     `pulumi:"classificationReason"`
	Labels                IncidentLabelArrayInput   `pulumi:"labels"`
	Owner                 IncidentOwnerInfoPtrInput `pulumi:"owner"`
	Severity              pulumi.StringPtrInput     `pulumi:"severity"`
	Status                pulumi.StringPtrInput     `pulumi:"status"`
}

func (AutomationRuleModifyPropertiesActionActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesActionActionConfiguration)(nil)).Elem()
}

func (i AutomationRuleModifyPropertiesActionActionConfigurationArgs) ToAutomationRuleModifyPropertiesActionActionConfigurationOutput() AutomationRuleModifyPropertiesActionActionConfigurationOutput {
	return i.ToAutomationRuleModifyPropertiesActionActionConfigurationOutputWithContext(context.Background())
}

func (i AutomationRuleModifyPropertiesActionActionConfigurationArgs) ToAutomationRuleModifyPropertiesActionActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleModifyPropertiesActionActionConfigurationOutput)
}

type AutomationRuleModifyPropertiesActionActionConfigurationOutput struct{ *pulumi.OutputState }

func (AutomationRuleModifyPropertiesActionActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesActionActionConfiguration)(nil)).Elem()
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) ToAutomationRuleModifyPropertiesActionActionConfigurationOutput() AutomationRuleModifyPropertiesActionActionConfigurationOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) ToAutomationRuleModifyPropertiesActionActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionActionConfigurationOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) Classification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) *string { return v.Classification }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) ClassificationComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) *string {
		return v.ClassificationComment
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) ClassificationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) *string { return v.ClassificationReason }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) Labels() IncidentLabelArrayOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) []IncidentLabel { return v.Labels }).(IncidentLabelArrayOutput)
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) Owner() IncidentOwnerInfoPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) *IncidentOwnerInfo { return v.Owner }).(IncidentOwnerInfoPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) *string { return v.Severity }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionActionConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionActionConfiguration) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type AutomationRuleModifyPropertiesActionResponse struct {
	ActionConfiguration AutomationRuleModifyPropertiesActionResponseActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                                          `pulumi:"actionType"`
	Order               int                                                             `pulumi:"order"`
}





type AutomationRuleModifyPropertiesActionResponseInput interface {
	pulumi.Input

	ToAutomationRuleModifyPropertiesActionResponseOutput() AutomationRuleModifyPropertiesActionResponseOutput
	ToAutomationRuleModifyPropertiesActionResponseOutputWithContext(context.Context) AutomationRuleModifyPropertiesActionResponseOutput
}

type AutomationRuleModifyPropertiesActionResponseArgs struct {
	ActionConfiguration AutomationRuleModifyPropertiesActionResponseActionConfigurationInput `pulumi:"actionConfiguration"`
	ActionType          pulumi.StringInput                                                   `pulumi:"actionType"`
	Order               pulumi.IntInput                                                      `pulumi:"order"`
}

func (AutomationRuleModifyPropertiesActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesActionResponse)(nil)).Elem()
}

func (i AutomationRuleModifyPropertiesActionResponseArgs) ToAutomationRuleModifyPropertiesActionResponseOutput() AutomationRuleModifyPropertiesActionResponseOutput {
	return i.ToAutomationRuleModifyPropertiesActionResponseOutputWithContext(context.Background())
}

func (i AutomationRuleModifyPropertiesActionResponseArgs) ToAutomationRuleModifyPropertiesActionResponseOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleModifyPropertiesActionResponseOutput)
}

type AutomationRuleModifyPropertiesActionResponseOutput struct{ *pulumi.OutputState }

func (AutomationRuleModifyPropertiesActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesActionResponse)(nil)).Elem()
}

func (o AutomationRuleModifyPropertiesActionResponseOutput) ToAutomationRuleModifyPropertiesActionResponseOutput() AutomationRuleModifyPropertiesActionResponseOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionResponseOutput) ToAutomationRuleModifyPropertiesActionResponseOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionResponseOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionResponseOutput) ActionConfiguration() AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponse) AutomationRuleModifyPropertiesActionResponseActionConfiguration {
		return v.ActionConfiguration
	}).(AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponse) int { return v.Order }).(pulumi.IntOutput)
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





type AutomationRuleModifyPropertiesActionResponseActionConfigurationInput interface {
	pulumi.Input

	ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutput() AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput
	ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutputWithContext(context.Context) AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput
}

type AutomationRuleModifyPropertiesActionResponseActionConfigurationArgs struct {
	Classification        pulumi.StringPtrInput             `pulumi:"classification"`
	ClassificationComment pulumi.StringPtrInput             `pulumi:"classificationComment"`
	ClassificationReason  pulumi.StringPtrInput             `pulumi:"classificationReason"`
	Labels                IncidentLabelResponseArrayInput   `pulumi:"labels"`
	Owner                 IncidentOwnerInfoResponsePtrInput `pulumi:"owner"`
	Severity              pulumi.StringPtrInput             `pulumi:"severity"`
	Status                pulumi.StringPtrInput             `pulumi:"status"`
}

func (AutomationRuleModifyPropertiesActionResponseActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesActionResponseActionConfiguration)(nil)).Elem()
}

func (i AutomationRuleModifyPropertiesActionResponseActionConfigurationArgs) ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutput() AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput {
	return i.ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutputWithContext(context.Background())
}

func (i AutomationRuleModifyPropertiesActionResponseActionConfigurationArgs) ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput)
}

type AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput struct{ *pulumi.OutputState }

func (AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleModifyPropertiesActionResponseActionConfiguration)(nil)).Elem()
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutput() AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) ToAutomationRuleModifyPropertiesActionResponseActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput {
	return o
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) Classification() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) *string {
		return v.Classification
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) ClassificationComment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) *string {
		return v.ClassificationComment
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) ClassificationReason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) *string {
		return v.ClassificationReason
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) Labels() IncidentLabelResponseArrayOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) []IncidentLabelResponse {
		return v.Labels
	}).(IncidentLabelResponseArrayOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) Owner() IncidentOwnerInfoResponsePtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) *IncidentOwnerInfoResponse {
		return v.Owner
	}).(IncidentOwnerInfoResponsePtrOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) Severity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) *string { return v.Severity }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleModifyPropertiesActionResponseActionConfiguration) *string { return v.Status }).(pulumi.StringPtrOutput)
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





type AutomationRulePropertyValuesConditionResponseInput interface {
	pulumi.Input

	ToAutomationRulePropertyValuesConditionResponseOutput() AutomationRulePropertyValuesConditionResponseOutput
	ToAutomationRulePropertyValuesConditionResponseOutputWithContext(context.Context) AutomationRulePropertyValuesConditionResponseOutput
}

type AutomationRulePropertyValuesConditionResponseArgs struct {
	ConditionProperties AutomationRulePropertyValuesConditionResponseConditionPropertiesInput `pulumi:"conditionProperties"`
	ConditionType       pulumi.StringInput                                                    `pulumi:"conditionType"`
}

func (AutomationRulePropertyValuesConditionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesConditionResponse)(nil)).Elem()
}

func (i AutomationRulePropertyValuesConditionResponseArgs) ToAutomationRulePropertyValuesConditionResponseOutput() AutomationRulePropertyValuesConditionResponseOutput {
	return i.ToAutomationRulePropertyValuesConditionResponseOutputWithContext(context.Background())
}

func (i AutomationRulePropertyValuesConditionResponseArgs) ToAutomationRulePropertyValuesConditionResponseOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRulePropertyValuesConditionResponseOutput)
}





type AutomationRulePropertyValuesConditionResponseArrayInput interface {
	pulumi.Input

	ToAutomationRulePropertyValuesConditionResponseArrayOutput() AutomationRulePropertyValuesConditionResponseArrayOutput
	ToAutomationRulePropertyValuesConditionResponseArrayOutputWithContext(context.Context) AutomationRulePropertyValuesConditionResponseArrayOutput
}

type AutomationRulePropertyValuesConditionResponseArray []AutomationRulePropertyValuesConditionResponseInput

func (AutomationRulePropertyValuesConditionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AutomationRulePropertyValuesConditionResponse)(nil)).Elem()
}

func (i AutomationRulePropertyValuesConditionResponseArray) ToAutomationRulePropertyValuesConditionResponseArrayOutput() AutomationRulePropertyValuesConditionResponseArrayOutput {
	return i.ToAutomationRulePropertyValuesConditionResponseArrayOutputWithContext(context.Background())
}

func (i AutomationRulePropertyValuesConditionResponseArray) ToAutomationRulePropertyValuesConditionResponseArrayOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRulePropertyValuesConditionResponseArrayOutput)
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





type AutomationRulePropertyValuesConditionResponseConditionPropertiesInput interface {
	pulumi.Input

	ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutput() AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput
	ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutputWithContext(context.Context) AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput
}

type AutomationRulePropertyValuesConditionResponseConditionPropertiesArgs struct {
	Operator       pulumi.StringPtrInput   `pulumi:"operator"`
	PropertyName   pulumi.StringPtrInput   `pulumi:"propertyName"`
	PropertyValues pulumi.StringArrayInput `pulumi:"propertyValues"`
}

func (AutomationRulePropertyValuesConditionResponseConditionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRulePropertyValuesConditionResponseConditionProperties)(nil)).Elem()
}

func (i AutomationRulePropertyValuesConditionResponseConditionPropertiesArgs) ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutput() AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput {
	return i.ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutputWithContext(context.Background())
}

func (i AutomationRulePropertyValuesConditionResponseConditionPropertiesArgs) ToAutomationRulePropertyValuesConditionResponseConditionPropertiesOutputWithContext(ctx context.Context) AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput)
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





type AutomationRuleRunPlaybookActionInput interface {
	pulumi.Input

	ToAutomationRuleRunPlaybookActionOutput() AutomationRuleRunPlaybookActionOutput
	ToAutomationRuleRunPlaybookActionOutputWithContext(context.Context) AutomationRuleRunPlaybookActionOutput
}

type AutomationRuleRunPlaybookActionArgs struct {
	ActionConfiguration AutomationRuleRunPlaybookActionActionConfigurationInput `pulumi:"actionConfiguration"`
	ActionType          pulumi.StringInput                                      `pulumi:"actionType"`
	Order               pulumi.IntInput                                         `pulumi:"order"`
}

func (AutomationRuleRunPlaybookActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookAction)(nil)).Elem()
}

func (i AutomationRuleRunPlaybookActionArgs) ToAutomationRuleRunPlaybookActionOutput() AutomationRuleRunPlaybookActionOutput {
	return i.ToAutomationRuleRunPlaybookActionOutputWithContext(context.Background())
}

func (i AutomationRuleRunPlaybookActionArgs) ToAutomationRuleRunPlaybookActionOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleRunPlaybookActionOutput)
}

type AutomationRuleRunPlaybookActionOutput struct{ *pulumi.OutputState }

func (AutomationRuleRunPlaybookActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookAction)(nil)).Elem()
}

func (o AutomationRuleRunPlaybookActionOutput) ToAutomationRuleRunPlaybookActionOutput() AutomationRuleRunPlaybookActionOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionOutput) ToAutomationRuleRunPlaybookActionOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionOutput) ActionConfiguration() AutomationRuleRunPlaybookActionActionConfigurationOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookAction) AutomationRuleRunPlaybookActionActionConfiguration {
		return v.ActionConfiguration
	}).(AutomationRuleRunPlaybookActionActionConfigurationOutput)
}

func (o AutomationRuleRunPlaybookActionOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookAction) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationRuleRunPlaybookActionOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookAction) int { return v.Order }).(pulumi.IntOutput)
}

type AutomationRuleRunPlaybookActionActionConfiguration struct {
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}





type AutomationRuleRunPlaybookActionActionConfigurationInput interface {
	pulumi.Input

	ToAutomationRuleRunPlaybookActionActionConfigurationOutput() AutomationRuleRunPlaybookActionActionConfigurationOutput
	ToAutomationRuleRunPlaybookActionActionConfigurationOutputWithContext(context.Context) AutomationRuleRunPlaybookActionActionConfigurationOutput
}

type AutomationRuleRunPlaybookActionActionConfigurationArgs struct {
	LogicAppResourceId pulumi.StringPtrInput `pulumi:"logicAppResourceId"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AutomationRuleRunPlaybookActionActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookActionActionConfiguration)(nil)).Elem()
}

func (i AutomationRuleRunPlaybookActionActionConfigurationArgs) ToAutomationRuleRunPlaybookActionActionConfigurationOutput() AutomationRuleRunPlaybookActionActionConfigurationOutput {
	return i.ToAutomationRuleRunPlaybookActionActionConfigurationOutputWithContext(context.Background())
}

func (i AutomationRuleRunPlaybookActionActionConfigurationArgs) ToAutomationRuleRunPlaybookActionActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleRunPlaybookActionActionConfigurationOutput)
}

type AutomationRuleRunPlaybookActionActionConfigurationOutput struct{ *pulumi.OutputState }

func (AutomationRuleRunPlaybookActionActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookActionActionConfiguration)(nil)).Elem()
}

func (o AutomationRuleRunPlaybookActionActionConfigurationOutput) ToAutomationRuleRunPlaybookActionActionConfigurationOutput() AutomationRuleRunPlaybookActionActionConfigurationOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionActionConfigurationOutput) ToAutomationRuleRunPlaybookActionActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionActionConfigurationOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionActionConfigurationOutput) LogicAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionActionConfiguration) *string { return v.LogicAppResourceId }).(pulumi.StringPtrOutput)
}

func (o AutomationRuleRunPlaybookActionActionConfigurationOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionActionConfiguration) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type AutomationRuleRunPlaybookActionResponse struct {
	ActionConfiguration AutomationRuleRunPlaybookActionResponseActionConfiguration `pulumi:"actionConfiguration"`
	ActionType          string                                                     `pulumi:"actionType"`
	Order               int                                                        `pulumi:"order"`
}





type AutomationRuleRunPlaybookActionResponseInput interface {
	pulumi.Input

	ToAutomationRuleRunPlaybookActionResponseOutput() AutomationRuleRunPlaybookActionResponseOutput
	ToAutomationRuleRunPlaybookActionResponseOutputWithContext(context.Context) AutomationRuleRunPlaybookActionResponseOutput
}

type AutomationRuleRunPlaybookActionResponseArgs struct {
	ActionConfiguration AutomationRuleRunPlaybookActionResponseActionConfigurationInput `pulumi:"actionConfiguration"`
	ActionType          pulumi.StringInput                                              `pulumi:"actionType"`
	Order               pulumi.IntInput                                                 `pulumi:"order"`
}

func (AutomationRuleRunPlaybookActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookActionResponse)(nil)).Elem()
}

func (i AutomationRuleRunPlaybookActionResponseArgs) ToAutomationRuleRunPlaybookActionResponseOutput() AutomationRuleRunPlaybookActionResponseOutput {
	return i.ToAutomationRuleRunPlaybookActionResponseOutputWithContext(context.Background())
}

func (i AutomationRuleRunPlaybookActionResponseArgs) ToAutomationRuleRunPlaybookActionResponseOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleRunPlaybookActionResponseOutput)
}

type AutomationRuleRunPlaybookActionResponseOutput struct{ *pulumi.OutputState }

func (AutomationRuleRunPlaybookActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookActionResponse)(nil)).Elem()
}

func (o AutomationRuleRunPlaybookActionResponseOutput) ToAutomationRuleRunPlaybookActionResponseOutput() AutomationRuleRunPlaybookActionResponseOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionResponseOutput) ToAutomationRuleRunPlaybookActionResponseOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionResponseOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionResponseOutput) ActionConfiguration() AutomationRuleRunPlaybookActionResponseActionConfigurationOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionResponse) AutomationRuleRunPlaybookActionResponseActionConfiguration {
		return v.ActionConfiguration
	}).(AutomationRuleRunPlaybookActionResponseActionConfigurationOutput)
}

func (o AutomationRuleRunPlaybookActionResponseOutput) ActionType() pulumi.StringOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionResponse) string { return v.ActionType }).(pulumi.StringOutput)
}

func (o AutomationRuleRunPlaybookActionResponseOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionResponse) int { return v.Order }).(pulumi.IntOutput)
}

type AutomationRuleRunPlaybookActionResponseActionConfiguration struct {
	LogicAppResourceId *string `pulumi:"logicAppResourceId"`
	TenantId           *string `pulumi:"tenantId"`
}





type AutomationRuleRunPlaybookActionResponseActionConfigurationInput interface {
	pulumi.Input

	ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutput() AutomationRuleRunPlaybookActionResponseActionConfigurationOutput
	ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutputWithContext(context.Context) AutomationRuleRunPlaybookActionResponseActionConfigurationOutput
}

type AutomationRuleRunPlaybookActionResponseActionConfigurationArgs struct {
	LogicAppResourceId pulumi.StringPtrInput `pulumi:"logicAppResourceId"`
	TenantId           pulumi.StringPtrInput `pulumi:"tenantId"`
}

func (AutomationRuleRunPlaybookActionResponseActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookActionResponseActionConfiguration)(nil)).Elem()
}

func (i AutomationRuleRunPlaybookActionResponseActionConfigurationArgs) ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutput() AutomationRuleRunPlaybookActionResponseActionConfigurationOutput {
	return i.ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutputWithContext(context.Background())
}

func (i AutomationRuleRunPlaybookActionResponseActionConfigurationArgs) ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionResponseActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleRunPlaybookActionResponseActionConfigurationOutput)
}

type AutomationRuleRunPlaybookActionResponseActionConfigurationOutput struct{ *pulumi.OutputState }

func (AutomationRuleRunPlaybookActionResponseActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleRunPlaybookActionResponseActionConfiguration)(nil)).Elem()
}

func (o AutomationRuleRunPlaybookActionResponseActionConfigurationOutput) ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutput() AutomationRuleRunPlaybookActionResponseActionConfigurationOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionResponseActionConfigurationOutput) ToAutomationRuleRunPlaybookActionResponseActionConfigurationOutputWithContext(ctx context.Context) AutomationRuleRunPlaybookActionResponseActionConfigurationOutput {
	return o
}

func (o AutomationRuleRunPlaybookActionResponseActionConfigurationOutput) LogicAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionResponseActionConfiguration) *string {
		return v.LogicAppResourceId
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleRunPlaybookActionResponseActionConfigurationOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutomationRuleRunPlaybookActionResponseActionConfiguration) *string { return v.TenantId }).(pulumi.StringPtrOutput)
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

func (i AutomationRuleTriggeringLogicArgs) ToAutomationRuleTriggeringLogicPtrOutput() AutomationRuleTriggeringLogicPtrOutput {
	return i.ToAutomationRuleTriggeringLogicPtrOutputWithContext(context.Background())
}

func (i AutomationRuleTriggeringLogicArgs) ToAutomationRuleTriggeringLogicPtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleTriggeringLogicOutput).ToAutomationRuleTriggeringLogicPtrOutputWithContext(ctx)
}









type AutomationRuleTriggeringLogicPtrInput interface {
	pulumi.Input

	ToAutomationRuleTriggeringLogicPtrOutput() AutomationRuleTriggeringLogicPtrOutput
	ToAutomationRuleTriggeringLogicPtrOutputWithContext(context.Context) AutomationRuleTriggeringLogicPtrOutput
}

type automationRuleTriggeringLogicPtrType AutomationRuleTriggeringLogicArgs

func AutomationRuleTriggeringLogicPtr(v *AutomationRuleTriggeringLogicArgs) AutomationRuleTriggeringLogicPtrInput {
	return (*automationRuleTriggeringLogicPtrType)(v)
}

func (*automationRuleTriggeringLogicPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleTriggeringLogic)(nil)).Elem()
}

func (i *automationRuleTriggeringLogicPtrType) ToAutomationRuleTriggeringLogicPtrOutput() AutomationRuleTriggeringLogicPtrOutput {
	return i.ToAutomationRuleTriggeringLogicPtrOutputWithContext(context.Background())
}

func (i *automationRuleTriggeringLogicPtrType) ToAutomationRuleTriggeringLogicPtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleTriggeringLogicPtrOutput)
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

func (o AutomationRuleTriggeringLogicOutput) ToAutomationRuleTriggeringLogicPtrOutput() AutomationRuleTriggeringLogicPtrOutput {
	return o.ToAutomationRuleTriggeringLogicPtrOutputWithContext(context.Background())
}

func (o AutomationRuleTriggeringLogicOutput) ToAutomationRuleTriggeringLogicPtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationRuleTriggeringLogic) *AutomationRuleTriggeringLogic {
		return &v
	}).(AutomationRuleTriggeringLogicPtrOutput)
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

type AutomationRuleTriggeringLogicPtrOutput struct{ *pulumi.OutputState }

func (AutomationRuleTriggeringLogicPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleTriggeringLogic)(nil)).Elem()
}

func (o AutomationRuleTriggeringLogicPtrOutput) ToAutomationRuleTriggeringLogicPtrOutput() AutomationRuleTriggeringLogicPtrOutput {
	return o
}

func (o AutomationRuleTriggeringLogicPtrOutput) ToAutomationRuleTriggeringLogicPtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicPtrOutput {
	return o
}

func (o AutomationRuleTriggeringLogicPtrOutput) Elem() AutomationRuleTriggeringLogicOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogic) AutomationRuleTriggeringLogic {
		if v != nil {
			return *v
		}
		var ret AutomationRuleTriggeringLogic
		return ret
	}).(AutomationRuleTriggeringLogicOutput)
}

func (o AutomationRuleTriggeringLogicPtrOutput) Conditions() AutomationRulePropertyValuesConditionArrayOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogic) []AutomationRulePropertyValuesCondition {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(AutomationRulePropertyValuesConditionArrayOutput)
}

func (o AutomationRuleTriggeringLogicPtrOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogic) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleTriggeringLogicPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogic) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutomationRuleTriggeringLogicPtrOutput) TriggersOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogic) *string {
		if v == nil {
			return nil
		}
		return &v.TriggersOn
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleTriggeringLogicPtrOutput) TriggersWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogic) *string {
		if v == nil {
			return nil
		}
		return &v.TriggersWhen
	}).(pulumi.StringPtrOutput)
}

type AutomationRuleTriggeringLogicResponse struct {
	Conditions        []AutomationRulePropertyValuesConditionResponse `pulumi:"conditions"`
	ExpirationTimeUtc *string                                         `pulumi:"expirationTimeUtc"`
	IsEnabled         bool                                            `pulumi:"isEnabled"`
	TriggersOn        string                                          `pulumi:"triggersOn"`
	TriggersWhen      string                                          `pulumi:"triggersWhen"`
}





type AutomationRuleTriggeringLogicResponseInput interface {
	pulumi.Input

	ToAutomationRuleTriggeringLogicResponseOutput() AutomationRuleTriggeringLogicResponseOutput
	ToAutomationRuleTriggeringLogicResponseOutputWithContext(context.Context) AutomationRuleTriggeringLogicResponseOutput
}

type AutomationRuleTriggeringLogicResponseArgs struct {
	Conditions        AutomationRulePropertyValuesConditionResponseArrayInput `pulumi:"conditions"`
	ExpirationTimeUtc pulumi.StringPtrInput                                   `pulumi:"expirationTimeUtc"`
	IsEnabled         pulumi.BoolInput                                        `pulumi:"isEnabled"`
	TriggersOn        pulumi.StringInput                                      `pulumi:"triggersOn"`
	TriggersWhen      pulumi.StringInput                                      `pulumi:"triggersWhen"`
}

func (AutomationRuleTriggeringLogicResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutomationRuleTriggeringLogicResponse)(nil)).Elem()
}

func (i AutomationRuleTriggeringLogicResponseArgs) ToAutomationRuleTriggeringLogicResponseOutput() AutomationRuleTriggeringLogicResponseOutput {
	return i.ToAutomationRuleTriggeringLogicResponseOutputWithContext(context.Background())
}

func (i AutomationRuleTriggeringLogicResponseArgs) ToAutomationRuleTriggeringLogicResponseOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleTriggeringLogicResponseOutput)
}

func (i AutomationRuleTriggeringLogicResponseArgs) ToAutomationRuleTriggeringLogicResponsePtrOutput() AutomationRuleTriggeringLogicResponsePtrOutput {
	return i.ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(context.Background())
}

func (i AutomationRuleTriggeringLogicResponseArgs) ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleTriggeringLogicResponseOutput).ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(ctx)
}









type AutomationRuleTriggeringLogicResponsePtrInput interface {
	pulumi.Input

	ToAutomationRuleTriggeringLogicResponsePtrOutput() AutomationRuleTriggeringLogicResponsePtrOutput
	ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(context.Context) AutomationRuleTriggeringLogicResponsePtrOutput
}

type automationRuleTriggeringLogicResponsePtrType AutomationRuleTriggeringLogicResponseArgs

func AutomationRuleTriggeringLogicResponsePtr(v *AutomationRuleTriggeringLogicResponseArgs) AutomationRuleTriggeringLogicResponsePtrInput {
	return (*automationRuleTriggeringLogicResponsePtrType)(v)
}

func (*automationRuleTriggeringLogicResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleTriggeringLogicResponse)(nil)).Elem()
}

func (i *automationRuleTriggeringLogicResponsePtrType) ToAutomationRuleTriggeringLogicResponsePtrOutput() AutomationRuleTriggeringLogicResponsePtrOutput {
	return i.ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(context.Background())
}

func (i *automationRuleTriggeringLogicResponsePtrType) ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutomationRuleTriggeringLogicResponsePtrOutput)
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

func (o AutomationRuleTriggeringLogicResponseOutput) ToAutomationRuleTriggeringLogicResponsePtrOutput() AutomationRuleTriggeringLogicResponsePtrOutput {
	return o.ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(context.Background())
}

func (o AutomationRuleTriggeringLogicResponseOutput) ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutomationRuleTriggeringLogicResponse) *AutomationRuleTriggeringLogicResponse {
		return &v
	}).(AutomationRuleTriggeringLogicResponsePtrOutput)
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

type AutomationRuleTriggeringLogicResponsePtrOutput struct{ *pulumi.OutputState }

func (AutomationRuleTriggeringLogicResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutomationRuleTriggeringLogicResponse)(nil)).Elem()
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) ToAutomationRuleTriggeringLogicResponsePtrOutput() AutomationRuleTriggeringLogicResponsePtrOutput {
	return o
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) ToAutomationRuleTriggeringLogicResponsePtrOutputWithContext(ctx context.Context) AutomationRuleTriggeringLogicResponsePtrOutput {
	return o
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) Elem() AutomationRuleTriggeringLogicResponseOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogicResponse) AutomationRuleTriggeringLogicResponse {
		if v != nil {
			return *v
		}
		var ret AutomationRuleTriggeringLogicResponse
		return ret
	}).(AutomationRuleTriggeringLogicResponseOutput)
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) Conditions() AutomationRulePropertyValuesConditionResponseArrayOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogicResponse) []AutomationRulePropertyValuesConditionResponse {
		if v == nil {
			return nil
		}
		return v.Conditions
	}).(AutomationRulePropertyValuesConditionResponseArrayOutput)
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) ExpirationTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogicResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExpirationTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogicResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) TriggersOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogicResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TriggersOn
	}).(pulumi.StringPtrOutput)
}

func (o AutomationRuleTriggeringLogicResponsePtrOutput) TriggersWhen() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutomationRuleTriggeringLogicResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TriggersWhen
	}).(pulumi.StringPtrOutput)
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





type AwsCloudTrailDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesResponseOutput() AwsCloudTrailDataConnectorDataTypesResponseOutput
	ToAwsCloudTrailDataConnectorDataTypesResponseOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesResponseOutput
}

type AwsCloudTrailDataConnectorDataTypesResponseArgs struct {
	Logs AwsCloudTrailDataConnectorDataTypesResponseLogsPtrInput `pulumi:"logs"`
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





type AwsCloudTrailDataConnectorDataTypesResponseLogsInput interface {
	pulumi.Input

	ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutput() AwsCloudTrailDataConnectorDataTypesResponseLogsOutput
	ToAwsCloudTrailDataConnectorDataTypesResponseLogsOutputWithContext(context.Context) AwsCloudTrailDataConnectorDataTypesResponseLogsOutput
}

type AwsCloudTrailDataConnectorDataTypesResponseLogsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
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





type BookmarkTimelineItemResponseInput interface {
	pulumi.Input

	ToBookmarkTimelineItemResponseOutput() BookmarkTimelineItemResponseOutput
	ToBookmarkTimelineItemResponseOutputWithContext(context.Context) BookmarkTimelineItemResponseOutput
}

type BookmarkTimelineItemResponseArgs struct {
	AzureResourceId pulumi.StringInput       `pulumi:"azureResourceId"`
	CreatedBy       UserInfoResponsePtrInput `pulumi:"createdBy"`
	DisplayName     pulumi.StringPtrInput    `pulumi:"displayName"`
	EndTimeUtc      pulumi.StringPtrInput    `pulumi:"endTimeUtc"`
	EventTime       pulumi.StringPtrInput    `pulumi:"eventTime"`
	Kind            pulumi.StringInput       `pulumi:"kind"`
	Labels          pulumi.StringArrayInput  `pulumi:"labels"`
	Notes           pulumi.StringPtrInput    `pulumi:"notes"`
	StartTimeUtc    pulumi.StringPtrInput    `pulumi:"startTimeUtc"`
}

func (BookmarkTimelineItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BookmarkTimelineItemResponse)(nil)).Elem()
}

func (i BookmarkTimelineItemResponseArgs) ToBookmarkTimelineItemResponseOutput() BookmarkTimelineItemResponseOutput {
	return i.ToBookmarkTimelineItemResponseOutputWithContext(context.Background())
}

func (i BookmarkTimelineItemResponseArgs) ToBookmarkTimelineItemResponseOutputWithContext(ctx context.Context) BookmarkTimelineItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BookmarkTimelineItemResponseOutput)
}

type BookmarkTimelineItemResponseOutput struct{ *pulumi.OutputState }

func (BookmarkTimelineItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BookmarkTimelineItemResponse)(nil)).Elem()
}

func (o BookmarkTimelineItemResponseOutput) ToBookmarkTimelineItemResponseOutput() BookmarkTimelineItemResponseOutput {
	return o
}

func (o BookmarkTimelineItemResponseOutput) ToBookmarkTimelineItemResponseOutputWithContext(ctx context.Context) BookmarkTimelineItemResponseOutput {
	return o
}

func (o BookmarkTimelineItemResponseOutput) AzureResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) string { return v.AzureResourceId }).(pulumi.StringOutput)
}

func (o BookmarkTimelineItemResponseOutput) CreatedBy() UserInfoResponsePtrOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) *UserInfoResponse { return v.CreatedBy }).(UserInfoResponsePtrOutput)
}

func (o BookmarkTimelineItemResponseOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) *string { return v.DisplayName }).(pulumi.StringPtrOutput)
}

func (o BookmarkTimelineItemResponseOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) *string { return v.EndTimeUtc }).(pulumi.StringPtrOutput)
}

func (o BookmarkTimelineItemResponseOutput) EventTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) *string { return v.EventTime }).(pulumi.StringPtrOutput)
}

func (o BookmarkTimelineItemResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o BookmarkTimelineItemResponseOutput) Labels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) []string { return v.Labels }).(pulumi.StringArrayOutput)
}

func (o BookmarkTimelineItemResponseOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) *string { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o BookmarkTimelineItemResponseOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BookmarkTimelineItemResponse) *string { return v.StartTimeUtc }).(pulumi.StringPtrOutput)
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





type DataConnectorDataTypeCommonResponseInput interface {
	pulumi.Input

	ToDataConnectorDataTypeCommonResponseOutput() DataConnectorDataTypeCommonResponseOutput
	ToDataConnectorDataTypeCommonResponseOutputWithContext(context.Context) DataConnectorDataTypeCommonResponseOutput
}

type DataConnectorDataTypeCommonResponseArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
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

type EntityInsightItemResponse struct {
	ChartQueryResults []InsightsTableResultResponse               `pulumi:"chartQueryResults"`
	QueryId           *string                                     `pulumi:"queryId"`
	QueryTimeInterval *EntityInsightItemResponseQueryTimeInterval `pulumi:"queryTimeInterval"`
	TableQueryResults *InsightsTableResultResponse                `pulumi:"tableQueryResults"`
}





type EntityInsightItemResponseInput interface {
	pulumi.Input

	ToEntityInsightItemResponseOutput() EntityInsightItemResponseOutput
	ToEntityInsightItemResponseOutputWithContext(context.Context) EntityInsightItemResponseOutput
}

type EntityInsightItemResponseArgs struct {
	ChartQueryResults InsightsTableResultResponseArrayInput              `pulumi:"chartQueryResults"`
	QueryId           pulumi.StringPtrInput                              `pulumi:"queryId"`
	QueryTimeInterval EntityInsightItemResponseQueryTimeIntervalPtrInput `pulumi:"queryTimeInterval"`
	TableQueryResults InsightsTableResultResponsePtrInput                `pulumi:"tableQueryResults"`
}

func (EntityInsightItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInsightItemResponse)(nil)).Elem()
}

func (i EntityInsightItemResponseArgs) ToEntityInsightItemResponseOutput() EntityInsightItemResponseOutput {
	return i.ToEntityInsightItemResponseOutputWithContext(context.Background())
}

func (i EntityInsightItemResponseArgs) ToEntityInsightItemResponseOutputWithContext(ctx context.Context) EntityInsightItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityInsightItemResponseOutput)
}





type EntityInsightItemResponseArrayInput interface {
	pulumi.Input

	ToEntityInsightItemResponseArrayOutput() EntityInsightItemResponseArrayOutput
	ToEntityInsightItemResponseArrayOutputWithContext(context.Context) EntityInsightItemResponseArrayOutput
}

type EntityInsightItemResponseArray []EntityInsightItemResponseInput

func (EntityInsightItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EntityInsightItemResponse)(nil)).Elem()
}

func (i EntityInsightItemResponseArray) ToEntityInsightItemResponseArrayOutput() EntityInsightItemResponseArrayOutput {
	return i.ToEntityInsightItemResponseArrayOutputWithContext(context.Background())
}

func (i EntityInsightItemResponseArray) ToEntityInsightItemResponseArrayOutputWithContext(ctx context.Context) EntityInsightItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityInsightItemResponseArrayOutput)
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





type EntityInsightItemResponseQueryTimeIntervalInput interface {
	pulumi.Input

	ToEntityInsightItemResponseQueryTimeIntervalOutput() EntityInsightItemResponseQueryTimeIntervalOutput
	ToEntityInsightItemResponseQueryTimeIntervalOutputWithContext(context.Context) EntityInsightItemResponseQueryTimeIntervalOutput
}

type EntityInsightItemResponseQueryTimeIntervalArgs struct {
	EndTime   pulumi.StringPtrInput `pulumi:"endTime"`
	StartTime pulumi.StringPtrInput `pulumi:"startTime"`
}

func (EntityInsightItemResponseQueryTimeIntervalArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EntityInsightItemResponseQueryTimeInterval)(nil)).Elem()
}

func (i EntityInsightItemResponseQueryTimeIntervalArgs) ToEntityInsightItemResponseQueryTimeIntervalOutput() EntityInsightItemResponseQueryTimeIntervalOutput {
	return i.ToEntityInsightItemResponseQueryTimeIntervalOutputWithContext(context.Background())
}

func (i EntityInsightItemResponseQueryTimeIntervalArgs) ToEntityInsightItemResponseQueryTimeIntervalOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityInsightItemResponseQueryTimeIntervalOutput)
}

func (i EntityInsightItemResponseQueryTimeIntervalArgs) ToEntityInsightItemResponseQueryTimeIntervalPtrOutput() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return i.ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(context.Background())
}

func (i EntityInsightItemResponseQueryTimeIntervalArgs) ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityInsightItemResponseQueryTimeIntervalOutput).ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(ctx)
}









type EntityInsightItemResponseQueryTimeIntervalPtrInput interface {
	pulumi.Input

	ToEntityInsightItemResponseQueryTimeIntervalPtrOutput() EntityInsightItemResponseQueryTimeIntervalPtrOutput
	ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(context.Context) EntityInsightItemResponseQueryTimeIntervalPtrOutput
}

type entityInsightItemResponseQueryTimeIntervalPtrType EntityInsightItemResponseQueryTimeIntervalArgs

func EntityInsightItemResponseQueryTimeIntervalPtr(v *EntityInsightItemResponseQueryTimeIntervalArgs) EntityInsightItemResponseQueryTimeIntervalPtrInput {
	return (*entityInsightItemResponseQueryTimeIntervalPtrType)(v)
}

func (*entityInsightItemResponseQueryTimeIntervalPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EntityInsightItemResponseQueryTimeInterval)(nil)).Elem()
}

func (i *entityInsightItemResponseQueryTimeIntervalPtrType) ToEntityInsightItemResponseQueryTimeIntervalPtrOutput() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return i.ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(context.Background())
}

func (i *entityInsightItemResponseQueryTimeIntervalPtrType) ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EntityInsightItemResponseQueryTimeIntervalPtrOutput)
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

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToEntityInsightItemResponseQueryTimeIntervalPtrOutput() EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o.ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(context.Background())
}

func (o EntityInsightItemResponseQueryTimeIntervalOutput) ToEntityInsightItemResponseQueryTimeIntervalPtrOutputWithContext(ctx context.Context) EntityInsightItemResponseQueryTimeIntervalPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EntityInsightItemResponseQueryTimeInterval) *EntityInsightItemResponseQueryTimeInterval {
		return &v
	}).(EntityInsightItemResponseQueryTimeIntervalPtrOutput)
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

type GetInsightsErrorResponse struct {
	ErrorMessage string  `pulumi:"errorMessage"`
	Kind         string  `pulumi:"kind"`
	QueryId      *string `pulumi:"queryId"`
}





type GetInsightsErrorResponseInput interface {
	pulumi.Input

	ToGetInsightsErrorResponseOutput() GetInsightsErrorResponseOutput
	ToGetInsightsErrorResponseOutputWithContext(context.Context) GetInsightsErrorResponseOutput
}

type GetInsightsErrorResponseArgs struct {
	ErrorMessage pulumi.StringInput    `pulumi:"errorMessage"`
	Kind         pulumi.StringInput    `pulumi:"kind"`
	QueryId      pulumi.StringPtrInput `pulumi:"queryId"`
}

func (GetInsightsErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsErrorResponse)(nil)).Elem()
}

func (i GetInsightsErrorResponseArgs) ToGetInsightsErrorResponseOutput() GetInsightsErrorResponseOutput {
	return i.ToGetInsightsErrorResponseOutputWithContext(context.Background())
}

func (i GetInsightsErrorResponseArgs) ToGetInsightsErrorResponseOutputWithContext(ctx context.Context) GetInsightsErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInsightsErrorResponseOutput)
}





type GetInsightsErrorResponseArrayInput interface {
	pulumi.Input

	ToGetInsightsErrorResponseArrayOutput() GetInsightsErrorResponseArrayOutput
	ToGetInsightsErrorResponseArrayOutputWithContext(context.Context) GetInsightsErrorResponseArrayOutput
}

type GetInsightsErrorResponseArray []GetInsightsErrorResponseInput

func (GetInsightsErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInsightsErrorResponse)(nil)).Elem()
}

func (i GetInsightsErrorResponseArray) ToGetInsightsErrorResponseArrayOutput() GetInsightsErrorResponseArrayOutput {
	return i.ToGetInsightsErrorResponseArrayOutputWithContext(context.Background())
}

func (i GetInsightsErrorResponseArray) ToGetInsightsErrorResponseArrayOutputWithContext(ctx context.Context) GetInsightsErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInsightsErrorResponseArrayOutput)
}

type GetInsightsErrorResponseOutput struct{ *pulumi.OutputState }

func (GetInsightsErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsErrorResponse)(nil)).Elem()
}

func (o GetInsightsErrorResponseOutput) ToGetInsightsErrorResponseOutput() GetInsightsErrorResponseOutput {
	return o
}

func (o GetInsightsErrorResponseOutput) ToGetInsightsErrorResponseOutputWithContext(ctx context.Context) GetInsightsErrorResponseOutput {
	return o
}

func (o GetInsightsErrorResponseOutput) ErrorMessage() pulumi.StringOutput {
	return o.ApplyT(func(v GetInsightsErrorResponse) string { return v.ErrorMessage }).(pulumi.StringOutput)
}

func (o GetInsightsErrorResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v GetInsightsErrorResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o GetInsightsErrorResponseOutput) QueryId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetInsightsErrorResponse) *string { return v.QueryId }).(pulumi.StringPtrOutput)
}

type GetInsightsErrorResponseArrayOutput struct{ *pulumi.OutputState }

func (GetInsightsErrorResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]GetInsightsErrorResponse)(nil)).Elem()
}

func (o GetInsightsErrorResponseArrayOutput) ToGetInsightsErrorResponseArrayOutput() GetInsightsErrorResponseArrayOutput {
	return o
}

func (o GetInsightsErrorResponseArrayOutput) ToGetInsightsErrorResponseArrayOutputWithContext(ctx context.Context) GetInsightsErrorResponseArrayOutput {
	return o
}

func (o GetInsightsErrorResponseArrayOutput) Index(i pulumi.IntInput) GetInsightsErrorResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) GetInsightsErrorResponse {
		return vs[0].([]GetInsightsErrorResponse)[vs[1].(int)]
	}).(GetInsightsErrorResponseOutput)
}

type GetInsightsResultsMetadataResponse struct {
	Errors     []GetInsightsErrorResponse `pulumi:"errors"`
	TotalCount int                        `pulumi:"totalCount"`
}





type GetInsightsResultsMetadataResponseInput interface {
	pulumi.Input

	ToGetInsightsResultsMetadataResponseOutput() GetInsightsResultsMetadataResponseOutput
	ToGetInsightsResultsMetadataResponseOutputWithContext(context.Context) GetInsightsResultsMetadataResponseOutput
}

type GetInsightsResultsMetadataResponseArgs struct {
	Errors     GetInsightsErrorResponseArrayInput `pulumi:"errors"`
	TotalCount pulumi.IntInput                    `pulumi:"totalCount"`
}

func (GetInsightsResultsMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetInsightsResultsMetadataResponse)(nil)).Elem()
}

func (i GetInsightsResultsMetadataResponseArgs) ToGetInsightsResultsMetadataResponseOutput() GetInsightsResultsMetadataResponseOutput {
	return i.ToGetInsightsResultsMetadataResponseOutputWithContext(context.Background())
}

func (i GetInsightsResultsMetadataResponseArgs) ToGetInsightsResultsMetadataResponseOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInsightsResultsMetadataResponseOutput)
}

func (i GetInsightsResultsMetadataResponseArgs) ToGetInsightsResultsMetadataResponsePtrOutput() GetInsightsResultsMetadataResponsePtrOutput {
	return i.ToGetInsightsResultsMetadataResponsePtrOutputWithContext(context.Background())
}

func (i GetInsightsResultsMetadataResponseArgs) ToGetInsightsResultsMetadataResponsePtrOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInsightsResultsMetadataResponseOutput).ToGetInsightsResultsMetadataResponsePtrOutputWithContext(ctx)
}









type GetInsightsResultsMetadataResponsePtrInput interface {
	pulumi.Input

	ToGetInsightsResultsMetadataResponsePtrOutput() GetInsightsResultsMetadataResponsePtrOutput
	ToGetInsightsResultsMetadataResponsePtrOutputWithContext(context.Context) GetInsightsResultsMetadataResponsePtrOutput
}

type getInsightsResultsMetadataResponsePtrType GetInsightsResultsMetadataResponseArgs

func GetInsightsResultsMetadataResponsePtr(v *GetInsightsResultsMetadataResponseArgs) GetInsightsResultsMetadataResponsePtrInput {
	return (*getInsightsResultsMetadataResponsePtrType)(v)
}

func (*getInsightsResultsMetadataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GetInsightsResultsMetadataResponse)(nil)).Elem()
}

func (i *getInsightsResultsMetadataResponsePtrType) ToGetInsightsResultsMetadataResponsePtrOutput() GetInsightsResultsMetadataResponsePtrOutput {
	return i.ToGetInsightsResultsMetadataResponsePtrOutputWithContext(context.Background())
}

func (i *getInsightsResultsMetadataResponsePtrType) ToGetInsightsResultsMetadataResponsePtrOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GetInsightsResultsMetadataResponsePtrOutput)
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

func (o GetInsightsResultsMetadataResponseOutput) ToGetInsightsResultsMetadataResponsePtrOutput() GetInsightsResultsMetadataResponsePtrOutput {
	return o.ToGetInsightsResultsMetadataResponsePtrOutputWithContext(context.Background())
}

func (o GetInsightsResultsMetadataResponseOutput) ToGetInsightsResultsMetadataResponsePtrOutputWithContext(ctx context.Context) GetInsightsResultsMetadataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GetInsightsResultsMetadataResponse) *GetInsightsResultsMetadataResponse {
		return &v
	}).(GetInsightsResultsMetadataResponsePtrOutput)
}

func (o GetInsightsResultsMetadataResponseOutput) Errors() GetInsightsErrorResponseArrayOutput {
	return o.ApplyT(func(v GetInsightsResultsMetadataResponse) []GetInsightsErrorResponse { return v.Errors }).(GetInsightsErrorResponseArrayOutput)
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

func (o GetInsightsResultsMetadataResponsePtrOutput) Errors() GetInsightsErrorResponseArrayOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) []GetInsightsErrorResponse {
		if v == nil {
			return nil
		}
		return v.Errors
	}).(GetInsightsErrorResponseArrayOutput)
}

func (o GetInsightsResultsMetadataResponsePtrOutput) TotalCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GetInsightsResultsMetadataResponse) *int {
		if v == nil {
			return nil
		}
		return &v.TotalCount
	}).(pulumi.IntPtrOutput)
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





type IncidentInfoResponseInput interface {
	pulumi.Input

	ToIncidentInfoResponseOutput() IncidentInfoResponseOutput
	ToIncidentInfoResponseOutputWithContext(context.Context) IncidentInfoResponseOutput
}

type IncidentInfoResponseArgs struct {
	IncidentId   pulumi.StringPtrInput `pulumi:"incidentId"`
	RelationName pulumi.StringPtrInput `pulumi:"relationName"`
	Severity     pulumi.StringPtrInput `pulumi:"severity"`
	Title        pulumi.StringPtrInput `pulumi:"title"`
}

func (IncidentInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncidentInfoResponse)(nil)).Elem()
}

func (i IncidentInfoResponseArgs) ToIncidentInfoResponseOutput() IncidentInfoResponseOutput {
	return i.ToIncidentInfoResponseOutputWithContext(context.Background())
}

func (i IncidentInfoResponseArgs) ToIncidentInfoResponseOutputWithContext(ctx context.Context) IncidentInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentInfoResponseOutput)
}

func (i IncidentInfoResponseArgs) ToIncidentInfoResponsePtrOutput() IncidentInfoResponsePtrOutput {
	return i.ToIncidentInfoResponsePtrOutputWithContext(context.Background())
}

func (i IncidentInfoResponseArgs) ToIncidentInfoResponsePtrOutputWithContext(ctx context.Context) IncidentInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentInfoResponseOutput).ToIncidentInfoResponsePtrOutputWithContext(ctx)
}









type IncidentInfoResponsePtrInput interface {
	pulumi.Input

	ToIncidentInfoResponsePtrOutput() IncidentInfoResponsePtrOutput
	ToIncidentInfoResponsePtrOutputWithContext(context.Context) IncidentInfoResponsePtrOutput
}

type incidentInfoResponsePtrType IncidentInfoResponseArgs

func IncidentInfoResponsePtr(v *IncidentInfoResponseArgs) IncidentInfoResponsePtrInput {
	return (*incidentInfoResponsePtrType)(v)
}

func (*incidentInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IncidentInfoResponse)(nil)).Elem()
}

func (i *incidentInfoResponsePtrType) ToIncidentInfoResponsePtrOutput() IncidentInfoResponsePtrOutput {
	return i.ToIncidentInfoResponsePtrOutputWithContext(context.Background())
}

func (i *incidentInfoResponsePtrType) ToIncidentInfoResponsePtrOutputWithContext(ctx context.Context) IncidentInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncidentInfoResponsePtrOutput)
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

func (o IncidentInfoResponseOutput) ToIncidentInfoResponsePtrOutput() IncidentInfoResponsePtrOutput {
	return o.ToIncidentInfoResponsePtrOutputWithContext(context.Background())
}

func (o IncidentInfoResponseOutput) ToIncidentInfoResponsePtrOutputWithContext(ctx context.Context) IncidentInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IncidentInfoResponse) *IncidentInfoResponse {
		return &v
	}).(IncidentInfoResponsePtrOutput)
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





type InsightsTableResultResponseInput interface {
	pulumi.Input

	ToInsightsTableResultResponseOutput() InsightsTableResultResponseOutput
	ToInsightsTableResultResponseOutputWithContext(context.Context) InsightsTableResultResponseOutput
}

type InsightsTableResultResponseArgs struct {
	Columns InsightsTableResultResponseColumnsArrayInput `pulumi:"columns"`
	Rows    pulumi.StringArrayArrayInput                 `pulumi:"rows"`
}

func (InsightsTableResultResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InsightsTableResultResponse)(nil)).Elem()
}

func (i InsightsTableResultResponseArgs) ToInsightsTableResultResponseOutput() InsightsTableResultResponseOutput {
	return i.ToInsightsTableResultResponseOutputWithContext(context.Background())
}

func (i InsightsTableResultResponseArgs) ToInsightsTableResultResponseOutputWithContext(ctx context.Context) InsightsTableResultResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightsTableResultResponseOutput)
}

func (i InsightsTableResultResponseArgs) ToInsightsTableResultResponsePtrOutput() InsightsTableResultResponsePtrOutput {
	return i.ToInsightsTableResultResponsePtrOutputWithContext(context.Background())
}

func (i InsightsTableResultResponseArgs) ToInsightsTableResultResponsePtrOutputWithContext(ctx context.Context) InsightsTableResultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightsTableResultResponseOutput).ToInsightsTableResultResponsePtrOutputWithContext(ctx)
}









type InsightsTableResultResponsePtrInput interface {
	pulumi.Input

	ToInsightsTableResultResponsePtrOutput() InsightsTableResultResponsePtrOutput
	ToInsightsTableResultResponsePtrOutputWithContext(context.Context) InsightsTableResultResponsePtrOutput
}

type insightsTableResultResponsePtrType InsightsTableResultResponseArgs

func InsightsTableResultResponsePtr(v *InsightsTableResultResponseArgs) InsightsTableResultResponsePtrInput {
	return (*insightsTableResultResponsePtrType)(v)
}

func (*insightsTableResultResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**InsightsTableResultResponse)(nil)).Elem()
}

func (i *insightsTableResultResponsePtrType) ToInsightsTableResultResponsePtrOutput() InsightsTableResultResponsePtrOutput {
	return i.ToInsightsTableResultResponsePtrOutputWithContext(context.Background())
}

func (i *insightsTableResultResponsePtrType) ToInsightsTableResultResponsePtrOutputWithContext(ctx context.Context) InsightsTableResultResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightsTableResultResponsePtrOutput)
}





type InsightsTableResultResponseArrayInput interface {
	pulumi.Input

	ToInsightsTableResultResponseArrayOutput() InsightsTableResultResponseArrayOutput
	ToInsightsTableResultResponseArrayOutputWithContext(context.Context) InsightsTableResultResponseArrayOutput
}

type InsightsTableResultResponseArray []InsightsTableResultResponseInput

func (InsightsTableResultResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InsightsTableResultResponse)(nil)).Elem()
}

func (i InsightsTableResultResponseArray) ToInsightsTableResultResponseArrayOutput() InsightsTableResultResponseArrayOutput {
	return i.ToInsightsTableResultResponseArrayOutputWithContext(context.Background())
}

func (i InsightsTableResultResponseArray) ToInsightsTableResultResponseArrayOutputWithContext(ctx context.Context) InsightsTableResultResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightsTableResultResponseArrayOutput)
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

func (o InsightsTableResultResponseOutput) ToInsightsTableResultResponsePtrOutput() InsightsTableResultResponsePtrOutput {
	return o.ToInsightsTableResultResponsePtrOutputWithContext(context.Background())
}

func (o InsightsTableResultResponseOutput) ToInsightsTableResultResponsePtrOutputWithContext(ctx context.Context) InsightsTableResultResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InsightsTableResultResponse) *InsightsTableResultResponse {
		return &v
	}).(InsightsTableResultResponsePtrOutput)
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





type InsightsTableResultResponseColumnsInput interface {
	pulumi.Input

	ToInsightsTableResultResponseColumnsOutput() InsightsTableResultResponseColumnsOutput
	ToInsightsTableResultResponseColumnsOutputWithContext(context.Context) InsightsTableResultResponseColumnsOutput
}

type InsightsTableResultResponseColumnsArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (InsightsTableResultResponseColumnsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InsightsTableResultResponseColumns)(nil)).Elem()
}

func (i InsightsTableResultResponseColumnsArgs) ToInsightsTableResultResponseColumnsOutput() InsightsTableResultResponseColumnsOutput {
	return i.ToInsightsTableResultResponseColumnsOutputWithContext(context.Background())
}

func (i InsightsTableResultResponseColumnsArgs) ToInsightsTableResultResponseColumnsOutputWithContext(ctx context.Context) InsightsTableResultResponseColumnsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightsTableResultResponseColumnsOutput)
}





type InsightsTableResultResponseColumnsArrayInput interface {
	pulumi.Input

	ToInsightsTableResultResponseColumnsArrayOutput() InsightsTableResultResponseColumnsArrayOutput
	ToInsightsTableResultResponseColumnsArrayOutputWithContext(context.Context) InsightsTableResultResponseColumnsArrayOutput
}

type InsightsTableResultResponseColumnsArray []InsightsTableResultResponseColumnsInput

func (InsightsTableResultResponseColumnsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InsightsTableResultResponseColumns)(nil)).Elem()
}

func (i InsightsTableResultResponseColumnsArray) ToInsightsTableResultResponseColumnsArrayOutput() InsightsTableResultResponseColumnsArrayOutput {
	return i.ToInsightsTableResultResponseColumnsArrayOutputWithContext(context.Background())
}

func (i InsightsTableResultResponseColumnsArray) ToInsightsTableResultResponseColumnsArrayOutputWithContext(ctx context.Context) InsightsTableResultResponseColumnsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InsightsTableResultResponseColumnsArrayOutput)
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





type MCASDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesResponseOutput() MCASDataConnectorDataTypesResponseOutput
	ToMCASDataConnectorDataTypesResponseOutputWithContext(context.Context) MCASDataConnectorDataTypesResponseOutput
}

type MCASDataConnectorDataTypesResponseArgs struct {
	Alerts        DataConnectorDataTypeCommonResponsePtrInput `pulumi:"alerts"`
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





type OfficeDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseOutput() OfficeDataConnectorDataTypesResponseOutput
	ToOfficeDataConnectorDataTypesResponseOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseOutput
}

type OfficeDataConnectorDataTypesResponseArgs struct {
	Exchange   OfficeDataConnectorDataTypesResponseExchangePtrInput   `pulumi:"exchange"`
	SharePoint OfficeDataConnectorDataTypesResponseSharePointPtrInput `pulumi:"sharePoint"`
	Teams      OfficeDataConnectorDataTypesResponseTeamsPtrInput      `pulumi:"teams"`
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





type OfficeDataConnectorDataTypesResponseExchangeInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseExchangeOutput() OfficeDataConnectorDataTypesResponseExchangeOutput
	ToOfficeDataConnectorDataTypesResponseExchangeOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseExchangeOutput
}

type OfficeDataConnectorDataTypesResponseExchangeArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
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





type OfficeDataConnectorDataTypesResponseSharePointInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseSharePointOutput() OfficeDataConnectorDataTypesResponseSharePointOutput
	ToOfficeDataConnectorDataTypesResponseSharePointOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseSharePointOutput
}

type OfficeDataConnectorDataTypesResponseSharePointArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
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





type OfficeDataConnectorDataTypesResponseTeamsInput interface {
	pulumi.Input

	ToOfficeDataConnectorDataTypesResponseTeamsOutput() OfficeDataConnectorDataTypesResponseTeamsOutput
	ToOfficeDataConnectorDataTypesResponseTeamsOutputWithContext(context.Context) OfficeDataConnectorDataTypesResponseTeamsOutput
}

type OfficeDataConnectorDataTypesResponseTeamsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
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
	Branch      *string          `pulumi:"branch"`
	PathMapping []ContentPathMap `pulumi:"pathMapping"`
	Url         *string          `pulumi:"url"`
}





type RepositoryInput interface {
	pulumi.Input

	ToRepositoryOutput() RepositoryOutput
	ToRepositoryOutputWithContext(context.Context) RepositoryOutput
}

type RepositoryArgs struct {
	Branch      pulumi.StringPtrInput    `pulumi:"branch"`
	PathMapping ContentPathMapArrayInput `pulumi:"pathMapping"`
	Url         pulumi.StringPtrInput    `pulumi:"url"`
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
	Branch      *string                  `pulumi:"branch"`
	PathMapping []ContentPathMapResponse `pulumi:"pathMapping"`
	Url         *string                  `pulumi:"url"`
}





type RepositoryResponseInput interface {
	pulumi.Input

	ToRepositoryResponseOutput() RepositoryResponseOutput
	ToRepositoryResponseOutputWithContext(context.Context) RepositoryResponseOutput
}

type RepositoryResponseArgs struct {
	Branch      pulumi.StringPtrInput            `pulumi:"branch"`
	PathMapping ContentPathMapResponseArrayInput `pulumi:"pathMapping"`
	Url         pulumi.StringPtrInput            `pulumi:"url"`
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





type SecurityAlertTimelineItemResponseInput interface {
	pulumi.Input

	ToSecurityAlertTimelineItemResponseOutput() SecurityAlertTimelineItemResponseOutput
	ToSecurityAlertTimelineItemResponseOutputWithContext(context.Context) SecurityAlertTimelineItemResponseOutput
}

type SecurityAlertTimelineItemResponseArgs struct {
	AlertType       pulumi.StringInput    `pulumi:"alertType"`
	AzureResourceId pulumi.StringInput    `pulumi:"azureResourceId"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	DisplayName     pulumi.StringInput    `pulumi:"displayName"`
	EndTimeUtc      pulumi.StringInput    `pulumi:"endTimeUtc"`
	Kind            pulumi.StringInput    `pulumi:"kind"`
	ProductName     pulumi.StringPtrInput `pulumi:"productName"`
	Severity        pulumi.StringInput    `pulumi:"severity"`
	StartTimeUtc    pulumi.StringInput    `pulumi:"startTimeUtc"`
	TimeGenerated   pulumi.StringInput    `pulumi:"timeGenerated"`
}

func (SecurityAlertTimelineItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertTimelineItemResponse)(nil)).Elem()
}

func (i SecurityAlertTimelineItemResponseArgs) ToSecurityAlertTimelineItemResponseOutput() SecurityAlertTimelineItemResponseOutput {
	return i.ToSecurityAlertTimelineItemResponseOutputWithContext(context.Background())
}

func (i SecurityAlertTimelineItemResponseArgs) ToSecurityAlertTimelineItemResponseOutputWithContext(ctx context.Context) SecurityAlertTimelineItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityAlertTimelineItemResponseOutput)
}

type SecurityAlertTimelineItemResponseOutput struct{ *pulumi.OutputState }

func (SecurityAlertTimelineItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecurityAlertTimelineItemResponse)(nil)).Elem()
}

func (o SecurityAlertTimelineItemResponseOutput) ToSecurityAlertTimelineItemResponseOutput() SecurityAlertTimelineItemResponseOutput {
	return o
}

func (o SecurityAlertTimelineItemResponseOutput) ToSecurityAlertTimelineItemResponseOutputWithContext(ctx context.Context) SecurityAlertTimelineItemResponseOutput {
	return o
}

func (o SecurityAlertTimelineItemResponseOutput) AlertType() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.AlertType }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) AzureResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.AzureResourceId }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) ProductName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) *string { return v.ProductName }).(pulumi.StringPtrOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.Severity }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o SecurityAlertTimelineItemResponseOutput) TimeGenerated() pulumi.StringOutput {
	return o.ApplyT(func(v SecurityAlertTimelineItemResponse) string { return v.TimeGenerated }).(pulumi.StringOutput)
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





type TIDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesResponseOutput() TIDataConnectorDataTypesResponseOutput
	ToTIDataConnectorDataTypesResponseOutputWithContext(context.Context) TIDataConnectorDataTypesResponseOutput
}

type TIDataConnectorDataTypesResponseArgs struct {
	Indicators TIDataConnectorDataTypesResponseIndicatorsPtrInput `pulumi:"indicators"`
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





type TIDataConnectorDataTypesResponseIndicatorsInput interface {
	pulumi.Input

	ToTIDataConnectorDataTypesResponseIndicatorsOutput() TIDataConnectorDataTypesResponseIndicatorsOutput
	ToTIDataConnectorDataTypesResponseIndicatorsOutputWithContext(context.Context) TIDataConnectorDataTypesResponseIndicatorsOutput
}

type TIDataConnectorDataTypesResponseIndicatorsArgs struct {
	State pulumi.StringPtrInput `pulumi:"state"`
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

type TimelineAggregationResponse struct {
	Count int    `pulumi:"count"`
	Kind  string `pulumi:"kind"`
}





type TimelineAggregationResponseInput interface {
	pulumi.Input

	ToTimelineAggregationResponseOutput() TimelineAggregationResponseOutput
	ToTimelineAggregationResponseOutputWithContext(context.Context) TimelineAggregationResponseOutput
}

type TimelineAggregationResponseArgs struct {
	Count pulumi.IntInput    `pulumi:"count"`
	Kind  pulumi.StringInput `pulumi:"kind"`
}

func (TimelineAggregationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineAggregationResponse)(nil)).Elem()
}

func (i TimelineAggregationResponseArgs) ToTimelineAggregationResponseOutput() TimelineAggregationResponseOutput {
	return i.ToTimelineAggregationResponseOutputWithContext(context.Background())
}

func (i TimelineAggregationResponseArgs) ToTimelineAggregationResponseOutputWithContext(ctx context.Context) TimelineAggregationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineAggregationResponseOutput)
}





type TimelineAggregationResponseArrayInput interface {
	pulumi.Input

	ToTimelineAggregationResponseArrayOutput() TimelineAggregationResponseArrayOutput
	ToTimelineAggregationResponseArrayOutputWithContext(context.Context) TimelineAggregationResponseArrayOutput
}

type TimelineAggregationResponseArray []TimelineAggregationResponseInput

func (TimelineAggregationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimelineAggregationResponse)(nil)).Elem()
}

func (i TimelineAggregationResponseArray) ToTimelineAggregationResponseArrayOutput() TimelineAggregationResponseArrayOutput {
	return i.ToTimelineAggregationResponseArrayOutputWithContext(context.Background())
}

func (i TimelineAggregationResponseArray) ToTimelineAggregationResponseArrayOutputWithContext(ctx context.Context) TimelineAggregationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineAggregationResponseArrayOutput)
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





type TimelineErrorResponseInput interface {
	pulumi.Input

	ToTimelineErrorResponseOutput() TimelineErrorResponseOutput
	ToTimelineErrorResponseOutputWithContext(context.Context) TimelineErrorResponseOutput
}

type TimelineErrorResponseArgs struct {
	ErrorMessage pulumi.StringInput    `pulumi:"errorMessage"`
	Kind         pulumi.StringInput    `pulumi:"kind"`
	QueryId      pulumi.StringPtrInput `pulumi:"queryId"`
}

func (TimelineErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineErrorResponse)(nil)).Elem()
}

func (i TimelineErrorResponseArgs) ToTimelineErrorResponseOutput() TimelineErrorResponseOutput {
	return i.ToTimelineErrorResponseOutputWithContext(context.Background())
}

func (i TimelineErrorResponseArgs) ToTimelineErrorResponseOutputWithContext(ctx context.Context) TimelineErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineErrorResponseOutput)
}





type TimelineErrorResponseArrayInput interface {
	pulumi.Input

	ToTimelineErrorResponseArrayOutput() TimelineErrorResponseArrayOutput
	ToTimelineErrorResponseArrayOutputWithContext(context.Context) TimelineErrorResponseArrayOutput
}

type TimelineErrorResponseArray []TimelineErrorResponseInput

func (TimelineErrorResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TimelineErrorResponse)(nil)).Elem()
}

func (i TimelineErrorResponseArray) ToTimelineErrorResponseArrayOutput() TimelineErrorResponseArrayOutput {
	return i.ToTimelineErrorResponseArrayOutputWithContext(context.Background())
}

func (i TimelineErrorResponseArray) ToTimelineErrorResponseArrayOutputWithContext(ctx context.Context) TimelineErrorResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineErrorResponseArrayOutput)
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





type TimelineResultsMetadataResponseInput interface {
	pulumi.Input

	ToTimelineResultsMetadataResponseOutput() TimelineResultsMetadataResponseOutput
	ToTimelineResultsMetadataResponseOutputWithContext(context.Context) TimelineResultsMetadataResponseOutput
}

type TimelineResultsMetadataResponseArgs struct {
	Aggregations TimelineAggregationResponseArrayInput `pulumi:"aggregations"`
	Errors       TimelineErrorResponseArrayInput       `pulumi:"errors"`
	TotalCount   pulumi.IntInput                       `pulumi:"totalCount"`
}

func (TimelineResultsMetadataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TimelineResultsMetadataResponse)(nil)).Elem()
}

func (i TimelineResultsMetadataResponseArgs) ToTimelineResultsMetadataResponseOutput() TimelineResultsMetadataResponseOutput {
	return i.ToTimelineResultsMetadataResponseOutputWithContext(context.Background())
}

func (i TimelineResultsMetadataResponseArgs) ToTimelineResultsMetadataResponseOutputWithContext(ctx context.Context) TimelineResultsMetadataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineResultsMetadataResponseOutput)
}

func (i TimelineResultsMetadataResponseArgs) ToTimelineResultsMetadataResponsePtrOutput() TimelineResultsMetadataResponsePtrOutput {
	return i.ToTimelineResultsMetadataResponsePtrOutputWithContext(context.Background())
}

func (i TimelineResultsMetadataResponseArgs) ToTimelineResultsMetadataResponsePtrOutputWithContext(ctx context.Context) TimelineResultsMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineResultsMetadataResponseOutput).ToTimelineResultsMetadataResponsePtrOutputWithContext(ctx)
}









type TimelineResultsMetadataResponsePtrInput interface {
	pulumi.Input

	ToTimelineResultsMetadataResponsePtrOutput() TimelineResultsMetadataResponsePtrOutput
	ToTimelineResultsMetadataResponsePtrOutputWithContext(context.Context) TimelineResultsMetadataResponsePtrOutput
}

type timelineResultsMetadataResponsePtrType TimelineResultsMetadataResponseArgs

func TimelineResultsMetadataResponsePtr(v *TimelineResultsMetadataResponseArgs) TimelineResultsMetadataResponsePtrInput {
	return (*timelineResultsMetadataResponsePtrType)(v)
}

func (*timelineResultsMetadataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TimelineResultsMetadataResponse)(nil)).Elem()
}

func (i *timelineResultsMetadataResponsePtrType) ToTimelineResultsMetadataResponsePtrOutput() TimelineResultsMetadataResponsePtrOutput {
	return i.ToTimelineResultsMetadataResponsePtrOutputWithContext(context.Background())
}

func (i *timelineResultsMetadataResponsePtrType) ToTimelineResultsMetadataResponsePtrOutputWithContext(ctx context.Context) TimelineResultsMetadataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TimelineResultsMetadataResponsePtrOutput)
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

func (o TimelineResultsMetadataResponseOutput) ToTimelineResultsMetadataResponsePtrOutput() TimelineResultsMetadataResponsePtrOutput {
	return o.ToTimelineResultsMetadataResponsePtrOutputWithContext(context.Background())
}

func (o TimelineResultsMetadataResponseOutput) ToTimelineResultsMetadataResponsePtrOutputWithContext(ctx context.Context) TimelineResultsMetadataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimelineResultsMetadataResponse) *TimelineResultsMetadataResponse {
		return &v
	}).(TimelineResultsMetadataResponsePtrOutput)
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
	ObjectId string `pulumi:"objectId"`
}





type UserInfoInput interface {
	pulumi.Input

	ToUserInfoOutput() UserInfoOutput
	ToUserInfoOutputWithContext(context.Context) UserInfoOutput
}

type UserInfoArgs struct {
	ObjectId pulumi.StringInput `pulumi:"objectId"`
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

func (o UserInfoOutput) ObjectId() pulumi.StringOutput {
	return o.ApplyT(func(v UserInfo) string { return v.ObjectId }).(pulumi.StringOutput)
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
		return &v.ObjectId
	}).(pulumi.StringPtrOutput)
}

type UserInfoResponse struct {
	Email    string  `pulumi:"email"`
	Name     string  `pulumi:"name"`
	ObjectId *string `pulumi:"objectId"`
}





type UserInfoResponseInput interface {
	pulumi.Input

	ToUserInfoResponseOutput() UserInfoResponseOutput
	ToUserInfoResponseOutputWithContext(context.Context) UserInfoResponseOutput
}

type UserInfoResponseArgs struct {
	Email    pulumi.StringInput    `pulumi:"email"`
	Name     pulumi.StringInput    `pulumi:"name"`
	ObjectId pulumi.StringPtrInput `pulumi:"objectId"`
}

func (UserInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserInfoResponse)(nil)).Elem()
}

func (i UserInfoResponseArgs) ToUserInfoResponseOutput() UserInfoResponseOutput {
	return i.ToUserInfoResponseOutputWithContext(context.Background())
}

func (i UserInfoResponseArgs) ToUserInfoResponseOutputWithContext(ctx context.Context) UserInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoResponseOutput)
}

func (i UserInfoResponseArgs) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return i.ToUserInfoResponsePtrOutputWithContext(context.Background())
}

func (i UserInfoResponseArgs) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoResponseOutput).ToUserInfoResponsePtrOutputWithContext(ctx)
}









type UserInfoResponsePtrInput interface {
	pulumi.Input

	ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput
	ToUserInfoResponsePtrOutputWithContext(context.Context) UserInfoResponsePtrOutput
}

type userInfoResponsePtrType UserInfoResponseArgs

func UserInfoResponsePtr(v *UserInfoResponseArgs) UserInfoResponsePtrInput {
	return (*userInfoResponsePtrType)(v)
}

func (*userInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UserInfoResponse)(nil)).Elem()
}

func (i *userInfoResponsePtrType) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return i.ToUserInfoResponsePtrOutputWithContext(context.Background())
}

func (i *userInfoResponsePtrType) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserInfoResponsePtrOutput)
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

func (o UserInfoResponseOutput) ToUserInfoResponsePtrOutput() UserInfoResponsePtrOutput {
	return o.ToUserInfoResponsePtrOutputWithContext(context.Background())
}

func (o UserInfoResponseOutput) ToUserInfoResponsePtrOutputWithContext(ctx context.Context) UserInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserInfoResponse) *UserInfoResponse {
		return &v
	}).(UserInfoResponsePtrOutput)
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
	pulumi.RegisterOutputType(ActivityTimelineItemResponseOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorPtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponsePtrOutput{})
	pulumi.RegisterOutputType(AutomationRuleModifyPropertiesActionOutput{})
	pulumi.RegisterOutputType(AutomationRuleModifyPropertiesActionActionConfigurationOutput{})
	pulumi.RegisterOutputType(AutomationRuleModifyPropertiesActionResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleModifyPropertiesActionResponseActionConfigurationOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionArrayOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionConditionPropertiesOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionResponseOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionResponseArrayOutput{})
	pulumi.RegisterOutputType(AutomationRulePropertyValuesConditionResponseConditionPropertiesOutput{})
	pulumi.RegisterOutputType(AutomationRuleRunPlaybookActionOutput{})
	pulumi.RegisterOutputType(AutomationRuleRunPlaybookActionActionConfigurationOutput{})
	pulumi.RegisterOutputType(AutomationRuleRunPlaybookActionResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleRunPlaybookActionResponseActionConfigurationOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicPtrOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicResponseOutput{})
	pulumi.RegisterOutputType(AutomationRuleTriggeringLogicResponsePtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesLogsPtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsOutput{})
	pulumi.RegisterOutputType(AwsCloudTrailDataConnectorDataTypesResponseLogsPtrOutput{})
	pulumi.RegisterOutputType(BookmarkTimelineItemResponseOutput{})
	pulumi.RegisterOutputType(ClientInfoResponseOutput{})
	pulumi.RegisterOutputType(ClientInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentPathMapOutput{})
	pulumi.RegisterOutputType(ContentPathMapArrayOutput{})
	pulumi.RegisterOutputType(ContentPathMapResponseOutput{})
	pulumi.RegisterOutputType(ContentPathMapResponseArrayOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonPtrOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonResponseOutput{})
	pulumi.RegisterOutputType(DataConnectorDataTypeCommonResponsePtrOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseArrayOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalPtrOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponsePtrOutput{})
	pulumi.RegisterOutputType(IncidentAdditionalDataResponseOutput{})
	pulumi.RegisterOutputType(IncidentAdditionalDataResponsePtrOutput{})
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
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponsePtrOutput{})
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
	pulumi.RegisterOutputType(RepoResponseOutput{})
	pulumi.RegisterOutputType(RepoResponseArrayOutput{})
	pulumi.RegisterOutputType(RepositoryOutput{})
	pulumi.RegisterOutputType(RepositoryPtrOutput{})
	pulumi.RegisterOutputType(RepositoryResponseOutput{})
	pulumi.RegisterOutputType(RepositoryResponsePtrOutput{})
	pulumi.RegisterOutputType(SecurityAlertTimelineItemResponseOutput{})
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
}
