


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





type AlertsDataTypeOfDataConnectorResponseInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorResponseOutput() AlertsDataTypeOfDataConnectorResponseOutput
	ToAlertsDataTypeOfDataConnectorResponseOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorResponseOutput
}

type AlertsDataTypeOfDataConnectorResponseArgs struct {
	Alerts AlertsDataTypeOfDataConnectorResponseAlertsInput `pulumi:"alerts"`
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





type AlertsDataTypeOfDataConnectorResponseAlertsInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorResponseAlertsOutput() AlertsDataTypeOfDataConnectorResponseAlertsOutput
	ToAlertsDataTypeOfDataConnectorResponseAlertsOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorResponseAlertsOutput
}

type AlertsDataTypeOfDataConnectorResponseAlertsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (AlertsDataTypeOfDataConnectorResponseAlertsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponseAlerts)(nil)).Elem()
}

func (i AlertsDataTypeOfDataConnectorResponseAlertsArgs) ToAlertsDataTypeOfDataConnectorResponseAlertsOutput() AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return i.ToAlertsDataTypeOfDataConnectorResponseAlertsOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorResponseAlertsArgs) ToAlertsDataTypeOfDataConnectorResponseAlertsOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseAlertsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorResponseAlertsOutput)
}

func (i AlertsDataTypeOfDataConnectorResponseAlertsArgs) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutput() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(context.Background())
}

func (i AlertsDataTypeOfDataConnectorResponseAlertsArgs) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorResponseAlertsOutput).ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(ctx)
}









type AlertsDataTypeOfDataConnectorResponseAlertsPtrInput interface {
	pulumi.Input

	ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutput() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput
	ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(context.Context) AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput
}

type alertsDataTypeOfDataConnectorResponseAlertsPtrType AlertsDataTypeOfDataConnectorResponseAlertsArgs

func AlertsDataTypeOfDataConnectorResponseAlertsPtr(v *AlertsDataTypeOfDataConnectorResponseAlertsArgs) AlertsDataTypeOfDataConnectorResponseAlertsPtrInput {
	return (*alertsDataTypeOfDataConnectorResponseAlertsPtrType)(v)
}

func (*alertsDataTypeOfDataConnectorResponseAlertsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertsDataTypeOfDataConnectorResponseAlerts)(nil)).Elem()
}

func (i *alertsDataTypeOfDataConnectorResponseAlertsPtrType) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutput() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return i.ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(context.Background())
}

func (i *alertsDataTypeOfDataConnectorResponseAlertsPtrType) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput)
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

func (o AlertsDataTypeOfDataConnectorResponseAlertsOutput) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutput() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return o.ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(context.Background())
}

func (o AlertsDataTypeOfDataConnectorResponseAlertsOutput) ToAlertsDataTypeOfDataConnectorResponseAlertsPtrOutputWithContext(ctx context.Context) AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlertsDataTypeOfDataConnectorResponseAlerts) *AlertsDataTypeOfDataConnectorResponseAlerts {
		return &v
	}).(AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput)
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





type GroupingConfigurationResponseInput interface {
	pulumi.Input

	ToGroupingConfigurationResponseOutput() GroupingConfigurationResponseOutput
	ToGroupingConfigurationResponseOutputWithContext(context.Context) GroupingConfigurationResponseOutput
}

type GroupingConfigurationResponseArgs struct {
	Enabled                pulumi.BoolInput        `pulumi:"enabled"`
	EntitiesMatchingMethod pulumi.StringInput      `pulumi:"entitiesMatchingMethod"`
	GroupByEntities        pulumi.StringArrayInput `pulumi:"groupByEntities"`
	LookbackDuration       pulumi.StringInput      `pulumi:"lookbackDuration"`
	ReopenClosedIncident   pulumi.BoolInput        `pulumi:"reopenClosedIncident"`
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

func (o MCASDataConnectorDataTypesOutput) Alerts() AlertsDataTypeOfDataConnectorAlertsOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) AlertsDataTypeOfDataConnectorAlerts { return v.Alerts }).(AlertsDataTypeOfDataConnectorAlertsOutput)
}

func (o MCASDataConnectorDataTypesOutput) DiscoveryLogs() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o.ApplyT(func(v MCASDataConnectorDataTypes) *MCASDataConnectorDataTypesDiscoveryLogs { return v.DiscoveryLogs }).(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput)
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

func (o MCASDataConnectorDataTypesPtrOutput) Alerts() AlertsDataTypeOfDataConnectorAlertsPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypes) *AlertsDataTypeOfDataConnectorAlerts {
		if v == nil {
			return nil
		}
		return &v.Alerts
	}).(AlertsDataTypeOfDataConnectorAlertsPtrOutput)
}

func (o MCASDataConnectorDataTypesPtrOutput) DiscoveryLogs() MCASDataConnectorDataTypesDiscoveryLogsPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypes) *MCASDataConnectorDataTypesDiscoveryLogs {
		if v == nil {
			return nil
		}
		return v.DiscoveryLogs
	}).(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput)
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





type MCASDataConnectorDataTypesResponseInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesResponseOutput() MCASDataConnectorDataTypesResponseOutput
	ToMCASDataConnectorDataTypesResponseOutputWithContext(context.Context) MCASDataConnectorDataTypesResponseOutput
}

type MCASDataConnectorDataTypesResponseArgs struct {
	Alerts        AlertsDataTypeOfDataConnectorResponseAlertsInput        `pulumi:"alerts"`
	DiscoveryLogs MCASDataConnectorDataTypesResponseDiscoveryLogsPtrInput `pulumi:"discoveryLogs"`
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

func (o MCASDataConnectorDataTypesResponsePtrOutput) Alerts() AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponse) *AlertsDataTypeOfDataConnectorResponseAlerts {
		if v == nil {
			return nil
		}
		return &v.Alerts
	}).(AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput)
}

func (o MCASDataConnectorDataTypesResponsePtrOutput) DiscoveryLogs() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return o.ApplyT(func(v *MCASDataConnectorDataTypesResponse) *MCASDataConnectorDataTypesResponseDiscoveryLogs {
		if v == nil {
			return nil
		}
		return v.DiscoveryLogs
	}).(MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput)
}

type MCASDataConnectorDataTypesResponseDiscoveryLogs struct {
	State string `pulumi:"state"`
}





type MCASDataConnectorDataTypesResponseDiscoveryLogsInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsOutput
	ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutputWithContext(context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsOutput
}

type MCASDataConnectorDataTypesResponseDiscoveryLogsArgs struct {
	State pulumi.StringInput `pulumi:"state"`
}

func (MCASDataConnectorDataTypesResponseDiscoveryLogsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MCASDataConnectorDataTypesResponseDiscoveryLogs)(nil)).Elem()
}

func (i MCASDataConnectorDataTypesResponseDiscoveryLogsArgs) ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsOutput {
	return i.ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesResponseDiscoveryLogsArgs) ToMCASDataConnectorDataTypesResponseDiscoveryLogsOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesResponseDiscoveryLogsOutput)
}

func (i MCASDataConnectorDataTypesResponseDiscoveryLogsArgs) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return i.ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(context.Background())
}

func (i MCASDataConnectorDataTypesResponseDiscoveryLogsArgs) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesResponseDiscoveryLogsOutput).ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(ctx)
}









type MCASDataConnectorDataTypesResponseDiscoveryLogsPtrInput interface {
	pulumi.Input

	ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput
	ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput
}

type mcasdataConnectorDataTypesResponseDiscoveryLogsPtrType MCASDataConnectorDataTypesResponseDiscoveryLogsArgs

func MCASDataConnectorDataTypesResponseDiscoveryLogsPtr(v *MCASDataConnectorDataTypesResponseDiscoveryLogsArgs) MCASDataConnectorDataTypesResponseDiscoveryLogsPtrInput {
	return (*mcasdataConnectorDataTypesResponseDiscoveryLogsPtrType)(v)
}

func (*mcasdataConnectorDataTypesResponseDiscoveryLogsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MCASDataConnectorDataTypesResponseDiscoveryLogs)(nil)).Elem()
}

func (i *mcasdataConnectorDataTypesResponseDiscoveryLogsPtrType) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return i.ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(context.Background())
}

func (i *mcasdataConnectorDataTypesResponseDiscoveryLogsPtrType) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput)
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

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsOutput) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput() MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return o.ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(context.Background())
}

func (o MCASDataConnectorDataTypesResponseDiscoveryLogsOutput) ToMCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutputWithContext(ctx context.Context) MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MCASDataConnectorDataTypesResponseDiscoveryLogs) *MCASDataConnectorDataTypesResponseDiscoveryLogs {
		return &v
	}).(MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput)
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
	pulumi.RegisterInputType(reflect.TypeOf((*ActivityTimelineItemResponseInput)(nil)).Elem(), ActivityTimelineItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorPtrInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorAlertsInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorAlertsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorAlertsPtrInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorAlertsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponseInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponsePtrInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponseAlertsInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorResponseAlertsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlertsDataTypeOfDataConnectorResponseAlertsPtrInput)(nil)).Elem(), AlertsDataTypeOfDataConnectorResponseAlertsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleModifyPropertiesActionInput)(nil)).Elem(), AutomationRuleModifyPropertiesActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleModifyPropertiesActionActionConfigurationInput)(nil)).Elem(), AutomationRuleModifyPropertiesActionActionConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleModifyPropertiesActionResponseInput)(nil)).Elem(), AutomationRuleModifyPropertiesActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleModifyPropertiesActionResponseActionConfigurationInput)(nil)).Elem(), AutomationRuleModifyPropertiesActionResponseActionConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRulePropertyValuesConditionInput)(nil)).Elem(), AutomationRulePropertyValuesConditionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRulePropertyValuesConditionArrayInput)(nil)).Elem(), AutomationRulePropertyValuesConditionArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRulePropertyValuesConditionConditionPropertiesInput)(nil)).Elem(), AutomationRulePropertyValuesConditionConditionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRulePropertyValuesConditionResponseInput)(nil)).Elem(), AutomationRulePropertyValuesConditionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRulePropertyValuesConditionResponseArrayInput)(nil)).Elem(), AutomationRulePropertyValuesConditionResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRulePropertyValuesConditionResponseConditionPropertiesInput)(nil)).Elem(), AutomationRulePropertyValuesConditionResponseConditionPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleRunPlaybookActionInput)(nil)).Elem(), AutomationRuleRunPlaybookActionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleRunPlaybookActionActionConfigurationInput)(nil)).Elem(), AutomationRuleRunPlaybookActionActionConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleRunPlaybookActionResponseInput)(nil)).Elem(), AutomationRuleRunPlaybookActionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleRunPlaybookActionResponseActionConfigurationInput)(nil)).Elem(), AutomationRuleRunPlaybookActionResponseActionConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleTriggeringLogicInput)(nil)).Elem(), AutomationRuleTriggeringLogicArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleTriggeringLogicPtrInput)(nil)).Elem(), AutomationRuleTriggeringLogicArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleTriggeringLogicResponseInput)(nil)).Elem(), AutomationRuleTriggeringLogicResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AutomationRuleTriggeringLogicResponsePtrInput)(nil)).Elem(), AutomationRuleTriggeringLogicResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesPtrInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesLogsInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesLogsPtrInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponseInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponseLogsInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesResponseLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AwsCloudTrailDataConnectorDataTypesResponseLogsPtrInput)(nil)).Elem(), AwsCloudTrailDataConnectorDataTypesResponseLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BookmarkTimelineItemResponseInput)(nil)).Elem(), BookmarkTimelineItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientInfoResponseInput)(nil)).Elem(), ClientInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ClientInfoResponsePtrInput)(nil)).Elem(), ClientInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesPtrInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponseInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponsePtrInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrInput)(nil)).Elem(), Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityInsightItemResponseInput)(nil)).Elem(), EntityInsightItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityInsightItemResponseArrayInput)(nil)).Elem(), EntityInsightItemResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityInsightItemResponseQueryTimeIntervalInput)(nil)).Elem(), EntityInsightItemResponseQueryTimeIntervalArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EntityInsightItemResponseQueryTimeIntervalPtrInput)(nil)).Elem(), EntityInsightItemResponseQueryTimeIntervalArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventGroupingSettingsInput)(nil)).Elem(), EventGroupingSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventGroupingSettingsPtrInput)(nil)).Elem(), EventGroupingSettingsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventGroupingSettingsResponseInput)(nil)).Elem(), EventGroupingSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventGroupingSettingsResponsePtrInput)(nil)).Elem(), EventGroupingSettingsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInsightsErrorResponseInput)(nil)).Elem(), GetInsightsErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInsightsErrorResponseArrayInput)(nil)).Elem(), GetInsightsErrorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInsightsResultsMetadataResponseInput)(nil)).Elem(), GetInsightsResultsMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GetInsightsResultsMetadataResponsePtrInput)(nil)).Elem(), GetInsightsResultsMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupingConfigurationInput)(nil)).Elem(), GroupingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupingConfigurationPtrInput)(nil)).Elem(), GroupingConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupingConfigurationResponseInput)(nil)).Elem(), GroupingConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*GroupingConfigurationResponsePtrInput)(nil)).Elem(), GroupingConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentAdditionalDataResponseInput)(nil)).Elem(), IncidentAdditionalDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentAdditionalDataResponsePtrInput)(nil)).Elem(), IncidentAdditionalDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentConfigurationInput)(nil)).Elem(), IncidentConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentConfigurationPtrInput)(nil)).Elem(), IncidentConfigurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentConfigurationResponseInput)(nil)).Elem(), IncidentConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentConfigurationResponsePtrInput)(nil)).Elem(), IncidentConfigurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentInfoInput)(nil)).Elem(), IncidentInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentInfoPtrInput)(nil)).Elem(), IncidentInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentInfoResponseInput)(nil)).Elem(), IncidentInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentInfoResponsePtrInput)(nil)).Elem(), IncidentInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentLabelInput)(nil)).Elem(), IncidentLabelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentLabelArrayInput)(nil)).Elem(), IncidentLabelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentLabelResponseInput)(nil)).Elem(), IncidentLabelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentLabelResponseArrayInput)(nil)).Elem(), IncidentLabelResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentOwnerInfoInput)(nil)).Elem(), IncidentOwnerInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentOwnerInfoPtrInput)(nil)).Elem(), IncidentOwnerInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentOwnerInfoResponseInput)(nil)).Elem(), IncidentOwnerInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IncidentOwnerInfoResponsePtrInput)(nil)).Elem(), IncidentOwnerInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightsTableResultResponseInput)(nil)).Elem(), InsightsTableResultResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightsTableResultResponsePtrInput)(nil)).Elem(), InsightsTableResultResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightsTableResultResponseArrayInput)(nil)).Elem(), InsightsTableResultResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightsTableResultResponseColumnsInput)(nil)).Elem(), InsightsTableResultResponseColumnsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*InsightsTableResultResponseColumnsArrayInput)(nil)).Elem(), InsightsTableResultResponseColumnsArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesInput)(nil)).Elem(), MCASDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesPtrInput)(nil)).Elem(), MCASDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesDiscoveryLogsInput)(nil)).Elem(), MCASDataConnectorDataTypesDiscoveryLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesDiscoveryLogsPtrInput)(nil)).Elem(), MCASDataConnectorDataTypesDiscoveryLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesResponseInput)(nil)).Elem(), MCASDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), MCASDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesResponseDiscoveryLogsInput)(nil)).Elem(), MCASDataConnectorDataTypesResponseDiscoveryLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MCASDataConnectorDataTypesResponseDiscoveryLogsPtrInput)(nil)).Elem(), MCASDataConnectorDataTypesResponseDiscoveryLogsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesInput)(nil)).Elem(), MSTIDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesPtrInput)(nil)).Elem(), MSTIDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesBingSafetyPhishingURLInput)(nil)).Elem(), MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesBingSafetyPhishingURLPtrInput)(nil)).Elem(), MSTIDataConnectorDataTypesBingSafetyPhishingURLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedInput)(nil)).Elem(), MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedPtrInput)(nil)).Elem(), MSTIDataConnectorDataTypesMicrosoftEmergingThreatFeedArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesResponseInput)(nil)).Elem(), MSTIDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), MSTIDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLInput)(nil)).Elem(), MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLPtrInput)(nil)).Elem(), MSTIDataConnectorDataTypesResponseBingSafetyPhishingURLArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedInput)(nil)).Elem(), MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedPtrInput)(nil)).Elem(), MSTIDataConnectorDataTypesResponseMicrosoftEmergingThreatFeedArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesInput)(nil)).Elem(), MTPDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesPtrInput)(nil)).Elem(), MTPDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesIncidentsInput)(nil)).Elem(), MTPDataConnectorDataTypesIncidentsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesIncidentsPtrInput)(nil)).Elem(), MTPDataConnectorDataTypesIncidentsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesResponseInput)(nil)).Elem(), MTPDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), MTPDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesResponseIncidentsInput)(nil)).Elem(), MTPDataConnectorDataTypesResponseIncidentsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MTPDataConnectorDataTypesResponseIncidentsPtrInput)(nil)).Elem(), MTPDataConnectorDataTypesResponseIncidentsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesInput)(nil)).Elem(), OfficeDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesPtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesExchangeInput)(nil)).Elem(), OfficeDataConnectorDataTypesExchangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesExchangePtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesExchangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseExchangeInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseExchangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseExchangePtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseExchangeArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseSharePointInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseSharePointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseSharePointPtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseSharePointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseTeamsInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseTeamsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesResponseTeamsPtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesResponseTeamsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesSharePointInput)(nil)).Elem(), OfficeDataConnectorDataTypesSharePointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesSharePointPtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesSharePointArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesTeamsInput)(nil)).Elem(), OfficeDataConnectorDataTypesTeamsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*OfficeDataConnectorDataTypesTeamsPtrInput)(nil)).Elem(), OfficeDataConnectorDataTypesTeamsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SecurityAlertTimelineItemResponseInput)(nil)).Elem(), SecurityAlertTimelineItemResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesInput)(nil)).Elem(), TIDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesPtrInput)(nil)).Elem(), TIDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesIndicatorsInput)(nil)).Elem(), TIDataConnectorDataTypesIndicatorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesIndicatorsPtrInput)(nil)).Elem(), TIDataConnectorDataTypesIndicatorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesResponseInput)(nil)).Elem(), TIDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), TIDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesResponseIndicatorsInput)(nil)).Elem(), TIDataConnectorDataTypesResponseIndicatorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TIDataConnectorDataTypesResponseIndicatorsPtrInput)(nil)).Elem(), TIDataConnectorDataTypesResponseIndicatorsArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceExternalReferenceInput)(nil)).Elem(), ThreatIntelligenceExternalReferenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceExternalReferenceArrayInput)(nil)).Elem(), ThreatIntelligenceExternalReferenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceGranularMarkingModelInput)(nil)).Elem(), ThreatIntelligenceGranularMarkingModelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceGranularMarkingModelArrayInput)(nil)).Elem(), ThreatIntelligenceGranularMarkingModelArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceKillChainPhaseInput)(nil)).Elem(), ThreatIntelligenceKillChainPhaseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceKillChainPhaseArrayInput)(nil)).Elem(), ThreatIntelligenceKillChainPhaseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceParsedPatternInput)(nil)).Elem(), ThreatIntelligenceParsedPatternArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceParsedPatternArrayInput)(nil)).Elem(), ThreatIntelligenceParsedPatternArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceParsedPatternTypeValueInput)(nil)).Elem(), ThreatIntelligenceParsedPatternTypeValueArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceParsedPatternTypeValueArrayInput)(nil)).Elem(), ThreatIntelligenceParsedPatternTypeValueArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesPtrInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponseInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponsePtrInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponseTaxiiClientInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesResponseTaxiiClientArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesTaxiiClientInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesTaxiiClientArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TiTaxiiDataConnectorDataTypesTaxiiClientPtrInput)(nil)).Elem(), TiTaxiiDataConnectorDataTypesTaxiiClientArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimelineAggregationResponseInput)(nil)).Elem(), TimelineAggregationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimelineAggregationResponseArrayInput)(nil)).Elem(), TimelineAggregationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimelineErrorResponseInput)(nil)).Elem(), TimelineErrorResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimelineErrorResponseArrayInput)(nil)).Elem(), TimelineErrorResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimelineResultsMetadataResponseInput)(nil)).Elem(), TimelineResultsMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TimelineResultsMetadataResponsePtrInput)(nil)).Elem(), TimelineResultsMetadataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoInput)(nil)).Elem(), UserInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoPtrInput)(nil)).Elem(), UserInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoResponseInput)(nil)).Elem(), UserInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UserInfoResponsePtrInput)(nil)).Elem(), UserInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistUserInfoInput)(nil)).Elem(), WatchlistUserInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistUserInfoPtrInput)(nil)).Elem(), WatchlistUserInfoArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistUserInfoResponseInput)(nil)).Elem(), WatchlistUserInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WatchlistUserInfoResponsePtrInput)(nil)).Elem(), WatchlistUserInfoResponseArgs{})
	pulumi.RegisterOutputType(ActivityTimelineItemResponseOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorPtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorAlertsOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorAlertsPtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponsePtrOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseAlertsOutput{})
	pulumi.RegisterOutputType(AlertsDataTypeOfDataConnectorResponseAlertsPtrOutput{})
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
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesDynamics365CdsActivitiesPtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesOutput{})
	pulumi.RegisterOutputType(Dynamics365DataConnectorDataTypesResponseDynamics365CdsActivitiesPtrOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseArrayOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalOutput{})
	pulumi.RegisterOutputType(EntityInsightItemResponseQueryTimeIntervalPtrOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsPtrOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsResponseOutput{})
	pulumi.RegisterOutputType(EventGroupingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsErrorResponseArrayOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponseOutput{})
	pulumi.RegisterOutputType(GetInsightsResultsMetadataResponsePtrOutput{})
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
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesDiscoveryLogsOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesDiscoveryLogsPtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseDiscoveryLogsOutput{})
	pulumi.RegisterOutputType(MCASDataConnectorDataTypesResponseDiscoveryLogsPtrOutput{})
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
	pulumi.RegisterOutputType(SecurityAlertTimelineItemResponseOutput{})
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
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesPtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponsePtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseTaxiiClientOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesResponseTaxiiClientPtrOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesTaxiiClientOutput{})
	pulumi.RegisterOutputType(TiTaxiiDataConnectorDataTypesTaxiiClientPtrOutput{})
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
