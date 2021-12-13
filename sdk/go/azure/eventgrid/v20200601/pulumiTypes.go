


package v20200601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureFunctionEventSubscriptionDestination struct {
	EndpointType                  string  `pulumi:"endpointType"`
	MaxEventsPerBatch             *int    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes *int    `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    *string `pulumi:"resourceId"`
}


func (val *AzureFunctionEventSubscriptionDestination) Defaults() *AzureFunctionEventSubscriptionDestination {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}

type AzureFunctionEventSubscriptionDestinationResponse struct {
	EndpointType                  string  `pulumi:"endpointType"`
	MaxEventsPerBatch             *int    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes *int    `pulumi:"preferredBatchSizeInKilobytes"`
	ResourceId                    *string `pulumi:"resourceId"`
}


func (val *AzureFunctionEventSubscriptionDestinationResponse) Defaults() *AzureFunctionEventSubscriptionDestinationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}

type BoolEqualsAdvancedFilter struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
	Value        *bool   `pulumi:"value"`
}

type BoolEqualsAdvancedFilterResponse struct {
	Key          *string `pulumi:"key"`
	OperatorType string  `pulumi:"operatorType"`
	Value        *bool   `pulumi:"value"`
}

type ConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type ConnectionStateInput interface {
	pulumi.Input

	ToConnectionStateOutput() ConnectionStateOutput
	ToConnectionStateOutputWithContext(context.Context) ConnectionStateOutput
}

type ConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (ConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (i ConnectionStateArgs) ToConnectionStateOutput() ConnectionStateOutput {
	return i.ToConnectionStateOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput)
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i ConnectionStateArgs) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStateOutput).ToConnectionStatePtrOutputWithContext(ctx)
}









type ConnectionStatePtrInput interface {
	pulumi.Input

	ToConnectionStatePtrOutput() ConnectionStatePtrOutput
	ToConnectionStatePtrOutputWithContext(context.Context) ConnectionStatePtrOutput
}

type connectionStatePtrType ConnectionStateArgs

func ConnectionStatePtr(v *ConnectionStateArgs) ConnectionStatePtrInput {
	return (*connectionStatePtrType)(v)
}

func (*connectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return i.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (i *connectionStatePtrType) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionStatePtrOutput)
}

type ConnectionStateOutput struct{ *pulumi.OutputState }

func (ConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionState)(nil)).Elem()
}

func (o ConnectionStateOutput) ToConnectionStateOutput() ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStateOutputWithContext(ctx context.Context) ConnectionStateOutput {
	return o
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o.ToConnectionStatePtrOutputWithContext(context.Background())
}

func (o ConnectionStateOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionState) *ConnectionState {
		return &v
	}).(ConnectionStatePtrOutput)
}

func (o ConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionState)(nil)).Elem()
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutput() ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) ToConnectionStatePtrOutputWithContext(ctx context.Context) ConnectionStatePtrOutput {
	return o
}

func (o ConnectionStatePtrOutput) Elem() ConnectionStateOutput {
	return o.ApplyT(func(v *ConnectionState) ConnectionState {
		if v != nil {
			return *v
		}
		var ret ConnectionState
		return ret
	}).(ConnectionStateOutput)
}

func (o ConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type ConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutput() ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ToConnectionStateResponseOutputWithContext(ctx context.Context) ConnectionStateResponseOutput {
	return o
}

func (o ConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionStateResponse)(nil)).Elem()
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutput() ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) ToConnectionStateResponsePtrOutputWithContext(ctx context.Context) ConnectionStateResponsePtrOutput {
	return o
}

func (o ConnectionStateResponsePtrOutput) Elem() ConnectionStateResponseOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) ConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionStateResponse
		return ret
	}).(ConnectionStateResponseOutput)
}

func (o ConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type EventHubEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type EventHubEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type EventSubscriptionFilter struct {
	AdvancedFilters        []interface{} `pulumi:"advancedFilters"`
	IncludedEventTypes     []string      `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive *bool         `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      *string       `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        *string       `pulumi:"subjectEndsWith"`
}


func (val *EventSubscriptionFilter) Defaults() *EventSubscriptionFilter {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		isSubjectCaseSensitive_ := false
		tmp.IsSubjectCaseSensitive = &isSubjectCaseSensitive_
	}
	return &tmp
}





type EventSubscriptionFilterInput interface {
	pulumi.Input

	ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput
	ToEventSubscriptionFilterOutputWithContext(context.Context) EventSubscriptionFilterOutput
}

type EventSubscriptionFilterArgs struct {
	AdvancedFilters        pulumi.ArrayInput       `pulumi:"advancedFilters"`
	IncludedEventTypes     pulumi.StringArrayInput `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive pulumi.BoolPtrInput     `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      pulumi.StringPtrInput   `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        pulumi.StringPtrInput   `pulumi:"subjectEndsWith"`
}

func (EventSubscriptionFilterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return i.ToEventSubscriptionFilterOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput)
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterArgs) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterOutput).ToEventSubscriptionFilterPtrOutputWithContext(ctx)
}









type EventSubscriptionFilterPtrInput interface {
	pulumi.Input

	ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput
	ToEventSubscriptionFilterPtrOutputWithContext(context.Context) EventSubscriptionFilterPtrOutput
}

type eventSubscriptionFilterPtrType EventSubscriptionFilterArgs

func EventSubscriptionFilterPtr(v *EventSubscriptionFilterArgs) EventSubscriptionFilterPtrInput {
	return (*eventSubscriptionFilterPtrType)(v)
}

func (*eventSubscriptionFilterPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return i.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionFilterPtrType) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterPtrOutput)
}

type EventSubscriptionFilterOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilter)(nil)).Elem()
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput {
	return o
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterOutputWithContext(ctx context.Context) EventSubscriptionFilterOutput {
	return o
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return o.ToEventSubscriptionFilterPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionFilterOutput) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionFilter) *EventSubscriptionFilter {
		return &v
	}).(EventSubscriptionFilterPtrOutput)
}

func (o EventSubscriptionFilterOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) []interface{} { return v.AdvancedFilters }).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilter) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterPtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilter)(nil)).Elem()
}

func (o EventSubscriptionFilterPtrOutput) ToEventSubscriptionFilterPtrOutput() EventSubscriptionFilterPtrOutput {
	return o
}

func (o EventSubscriptionFilterPtrOutput) ToEventSubscriptionFilterPtrOutputWithContext(ctx context.Context) EventSubscriptionFilterPtrOutput {
	return o
}

func (o EventSubscriptionFilterPtrOutput) Elem() EventSubscriptionFilterOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) EventSubscriptionFilter {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionFilter
		return ret
	}).(EventSubscriptionFilterOutput)
}

func (o EventSubscriptionFilterPtrOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) []interface{} {
		if v == nil {
			return nil
		}
		return v.AdvancedFilters
	}).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterPtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterPtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterPtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilter) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponse struct {
	AdvancedFilters        []interface{} `pulumi:"advancedFilters"`
	IncludedEventTypes     []string      `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive *bool         `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      *string       `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        *string       `pulumi:"subjectEndsWith"`
}


func (val *EventSubscriptionFilterResponse) Defaults() *EventSubscriptionFilterResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsSubjectCaseSensitive) {
		isSubjectCaseSensitive_ := false
		tmp.IsSubjectCaseSensitive = &isSubjectCaseSensitive_
	}
	return &tmp
}

type EventSubscriptionFilterResponseOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponseOutputWithContext(ctx context.Context) EventSubscriptionFilterResponseOutput {
	return o
}

func (o EventSubscriptionFilterResponseOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) []interface{} { return v.AdvancedFilters }).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterResponseOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) []string { return v.IncludedEventTypes }).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterResponseOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *bool { return v.IsSubjectCaseSensitive }).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectBeginsWith }).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterResponseOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventSubscriptionFilterResponse) *string { return v.SubjectEndsWith }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilterResponsePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionFilterResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilterResponse)(nil)).Elem()
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o
}

func (o EventSubscriptionFilterResponsePtrOutput) Elem() EventSubscriptionFilterResponseOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) EventSubscriptionFilterResponse {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionFilterResponse
		return ret
	}).(EventSubscriptionFilterResponseOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) AdvancedFilters() pulumi.ArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) []interface{} {
		if v == nil {
			return nil
		}
		return v.AdvancedFilters
	}).(pulumi.ArrayOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) IncludedEventTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) []string {
		if v == nil {
			return nil
		}
		return v.IncludedEventTypes
	}).(pulumi.StringArrayOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) IsSubjectCaseSensitive() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsSubjectCaseSensitive
	}).(pulumi.BoolPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) SubjectBeginsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectBeginsWith
	}).(pulumi.StringPtrOutput)
}

func (o EventSubscriptionFilterResponsePtrOutput) SubjectEndsWith() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventSubscriptionFilterResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubjectEndsWith
	}).(pulumi.StringPtrOutput)
}

type HybridConnectionEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type HybridConnectionEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type InboundIpRule struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}





type InboundIpRuleInput interface {
	pulumi.Input

	ToInboundIpRuleOutput() InboundIpRuleOutput
	ToInboundIpRuleOutputWithContext(context.Context) InboundIpRuleOutput
}

type InboundIpRuleArgs struct {
	Action pulumi.StringPtrInput `pulumi:"action"`
	IpMask pulumi.StringPtrInput `pulumi:"ipMask"`
}

func (InboundIpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRule)(nil)).Elem()
}

func (i InboundIpRuleArgs) ToInboundIpRuleOutput() InboundIpRuleOutput {
	return i.ToInboundIpRuleOutputWithContext(context.Background())
}

func (i InboundIpRuleArgs) ToInboundIpRuleOutputWithContext(ctx context.Context) InboundIpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundIpRuleOutput)
}





type InboundIpRuleArrayInput interface {
	pulumi.Input

	ToInboundIpRuleArrayOutput() InboundIpRuleArrayOutput
	ToInboundIpRuleArrayOutputWithContext(context.Context) InboundIpRuleArrayOutput
}

type InboundIpRuleArray []InboundIpRuleInput

func (InboundIpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRule)(nil)).Elem()
}

func (i InboundIpRuleArray) ToInboundIpRuleArrayOutput() InboundIpRuleArrayOutput {
	return i.ToInboundIpRuleArrayOutputWithContext(context.Background())
}

func (i InboundIpRuleArray) ToInboundIpRuleArrayOutputWithContext(ctx context.Context) InboundIpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(InboundIpRuleArrayOutput)
}

type InboundIpRuleOutput struct{ *pulumi.OutputState }

func (InboundIpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRule)(nil)).Elem()
}

func (o InboundIpRuleOutput) ToInboundIpRuleOutput() InboundIpRuleOutput {
	return o
}

func (o InboundIpRuleOutput) ToInboundIpRuleOutputWithContext(ctx context.Context) InboundIpRuleOutput {
	return o
}

func (o InboundIpRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o InboundIpRuleOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRule) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type InboundIpRuleArrayOutput struct{ *pulumi.OutputState }

func (InboundIpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRule)(nil)).Elem()
}

func (o InboundIpRuleArrayOutput) ToInboundIpRuleArrayOutput() InboundIpRuleArrayOutput {
	return o
}

func (o InboundIpRuleArrayOutput) ToInboundIpRuleArrayOutputWithContext(ctx context.Context) InboundIpRuleArrayOutput {
	return o
}

func (o InboundIpRuleArrayOutput) Index(i pulumi.IntInput) InboundIpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundIpRule {
		return vs[0].([]InboundIpRule)[vs[1].(int)]
	}).(InboundIpRuleOutput)
}

type InboundIpRuleResponse struct {
	Action *string `pulumi:"action"`
	IpMask *string `pulumi:"ipMask"`
}

type InboundIpRuleResponseOutput struct{ *pulumi.OutputState }

func (InboundIpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InboundIpRuleResponse)(nil)).Elem()
}

func (o InboundIpRuleResponseOutput) ToInboundIpRuleResponseOutput() InboundIpRuleResponseOutput {
	return o
}

func (o InboundIpRuleResponseOutput) ToInboundIpRuleResponseOutputWithContext(ctx context.Context) InboundIpRuleResponseOutput {
	return o
}

func (o InboundIpRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o InboundIpRuleResponseOutput) IpMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v InboundIpRuleResponse) *string { return v.IpMask }).(pulumi.StringPtrOutput)
}

type InboundIpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (InboundIpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]InboundIpRuleResponse)(nil)).Elem()
}

func (o InboundIpRuleResponseArrayOutput) ToInboundIpRuleResponseArrayOutput() InboundIpRuleResponseArrayOutput {
	return o
}

func (o InboundIpRuleResponseArrayOutput) ToInboundIpRuleResponseArrayOutputWithContext(ctx context.Context) InboundIpRuleResponseArrayOutput {
	return o
}

func (o InboundIpRuleResponseArrayOutput) Index(i pulumi.IntInput) InboundIpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) InboundIpRuleResponse {
		return vs[0].([]InboundIpRuleResponse)[vs[1].(int)]
	}).(InboundIpRuleResponseOutput)
}

type JsonField struct {
	SourceField *string `pulumi:"sourceField"`
}





type JsonFieldInput interface {
	pulumi.Input

	ToJsonFieldOutput() JsonFieldOutput
	ToJsonFieldOutputWithContext(context.Context) JsonFieldOutput
}

type JsonFieldArgs struct {
	SourceField pulumi.StringPtrInput `pulumi:"sourceField"`
}

func (JsonFieldArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonField)(nil)).Elem()
}

func (i JsonFieldArgs) ToJsonFieldOutput() JsonFieldOutput {
	return i.ToJsonFieldOutputWithContext(context.Background())
}

func (i JsonFieldArgs) ToJsonFieldOutputWithContext(ctx context.Context) JsonFieldOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldOutput)
}

func (i JsonFieldArgs) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return i.ToJsonFieldPtrOutputWithContext(context.Background())
}

func (i JsonFieldArgs) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldOutput).ToJsonFieldPtrOutputWithContext(ctx)
}









type JsonFieldPtrInput interface {
	pulumi.Input

	ToJsonFieldPtrOutput() JsonFieldPtrOutput
	ToJsonFieldPtrOutputWithContext(context.Context) JsonFieldPtrOutput
}

type jsonFieldPtrType JsonFieldArgs

func JsonFieldPtr(v *JsonFieldArgs) JsonFieldPtrInput {
	return (*jsonFieldPtrType)(v)
}

func (*jsonFieldPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonField)(nil)).Elem()
}

func (i *jsonFieldPtrType) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return i.ToJsonFieldPtrOutputWithContext(context.Background())
}

func (i *jsonFieldPtrType) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldPtrOutput)
}

type JsonFieldOutput struct{ *pulumi.OutputState }

func (JsonFieldOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonField)(nil)).Elem()
}

func (o JsonFieldOutput) ToJsonFieldOutput() JsonFieldOutput {
	return o
}

func (o JsonFieldOutput) ToJsonFieldOutputWithContext(ctx context.Context) JsonFieldOutput {
	return o
}

func (o JsonFieldOutput) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return o.ToJsonFieldPtrOutputWithContext(context.Background())
}

func (o JsonFieldOutput) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonField) *JsonField {
		return &v
	}).(JsonFieldPtrOutput)
}

func (o JsonFieldOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonField) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldPtrOutput struct{ *pulumi.OutputState }

func (JsonFieldPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonField)(nil)).Elem()
}

func (o JsonFieldPtrOutput) ToJsonFieldPtrOutput() JsonFieldPtrOutput {
	return o
}

func (o JsonFieldPtrOutput) ToJsonFieldPtrOutputWithContext(ctx context.Context) JsonFieldPtrOutput {
	return o
}

func (o JsonFieldPtrOutput) Elem() JsonFieldOutput {
	return o.ApplyT(func(v *JsonField) JsonField {
		if v != nil {
			return *v
		}
		var ret JsonField
		return ret
	}).(JsonFieldOutput)
}

func (o JsonFieldPtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonField) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonFieldResponse struct {
	SourceField *string `pulumi:"sourceField"`
}

type JsonFieldResponseOutput struct{ *pulumi.OutputState }

func (JsonFieldResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldResponse)(nil)).Elem()
}

func (o JsonFieldResponseOutput) ToJsonFieldResponseOutput() JsonFieldResponseOutput {
	return o
}

func (o JsonFieldResponseOutput) ToJsonFieldResponseOutputWithContext(ctx context.Context) JsonFieldResponseOutput {
	return o
}

func (o JsonFieldResponseOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldResponse) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldResponsePtrOutput struct{ *pulumi.OutputState }

func (JsonFieldResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldResponse)(nil)).Elem()
}

func (o JsonFieldResponsePtrOutput) ToJsonFieldResponsePtrOutput() JsonFieldResponsePtrOutput {
	return o
}

func (o JsonFieldResponsePtrOutput) ToJsonFieldResponsePtrOutputWithContext(ctx context.Context) JsonFieldResponsePtrOutput {
	return o
}

func (o JsonFieldResponsePtrOutput) Elem() JsonFieldResponseOutput {
	return o.ApplyT(func(v *JsonFieldResponse) JsonFieldResponse {
		if v != nil {
			return *v
		}
		var ret JsonFieldResponse
		return ret
	}).(JsonFieldResponseOutput)
}

func (o JsonFieldResponsePtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefault struct {
	DefaultValue *string `pulumi:"defaultValue"`
	SourceField  *string `pulumi:"sourceField"`
}





type JsonFieldWithDefaultInput interface {
	pulumi.Input

	ToJsonFieldWithDefaultOutput() JsonFieldWithDefaultOutput
	ToJsonFieldWithDefaultOutputWithContext(context.Context) JsonFieldWithDefaultOutput
}

type JsonFieldWithDefaultArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	SourceField  pulumi.StringPtrInput `pulumi:"sourceField"`
}

func (JsonFieldWithDefaultArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefault)(nil)).Elem()
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultOutput() JsonFieldWithDefaultOutput {
	return i.ToJsonFieldWithDefaultOutputWithContext(context.Background())
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultOutputWithContext(ctx context.Context) JsonFieldWithDefaultOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultOutput)
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return i.ToJsonFieldWithDefaultPtrOutputWithContext(context.Background())
}

func (i JsonFieldWithDefaultArgs) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultOutput).ToJsonFieldWithDefaultPtrOutputWithContext(ctx)
}









type JsonFieldWithDefaultPtrInput interface {
	pulumi.Input

	ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput
	ToJsonFieldWithDefaultPtrOutputWithContext(context.Context) JsonFieldWithDefaultPtrOutput
}

type jsonFieldWithDefaultPtrType JsonFieldWithDefaultArgs

func JsonFieldWithDefaultPtr(v *JsonFieldWithDefaultArgs) JsonFieldWithDefaultPtrInput {
	return (*jsonFieldWithDefaultPtrType)(v)
}

func (*jsonFieldWithDefaultPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefault)(nil)).Elem()
}

func (i *jsonFieldWithDefaultPtrType) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return i.ToJsonFieldWithDefaultPtrOutputWithContext(context.Background())
}

func (i *jsonFieldWithDefaultPtrType) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonFieldWithDefaultPtrOutput)
}

type JsonFieldWithDefaultOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefault)(nil)).Elem()
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultOutput() JsonFieldWithDefaultOutput {
	return o
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultOutputWithContext(ctx context.Context) JsonFieldWithDefaultOutput {
	return o
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return o.ToJsonFieldWithDefaultPtrOutputWithContext(context.Background())
}

func (o JsonFieldWithDefaultOutput) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonFieldWithDefault) *JsonFieldWithDefault {
		return &v
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonFieldWithDefaultOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefault) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefault) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefaultPtrOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefault)(nil)).Elem()
}

func (o JsonFieldWithDefaultPtrOutput) ToJsonFieldWithDefaultPtrOutput() JsonFieldWithDefaultPtrOutput {
	return o
}

func (o JsonFieldWithDefaultPtrOutput) ToJsonFieldWithDefaultPtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultPtrOutput {
	return o
}

func (o JsonFieldWithDefaultPtrOutput) Elem() JsonFieldWithDefaultOutput {
	return o.ApplyT(func(v *JsonFieldWithDefault) JsonFieldWithDefault {
		if v != nil {
			return *v
		}
		var ret JsonFieldWithDefault
		return ret
	}).(JsonFieldWithDefaultOutput)
}

func (o JsonFieldWithDefaultPtrOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefault) *string {
		if v == nil {
			return nil
		}
		return v.DefaultValue
	}).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultPtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefault) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefaultResponse struct {
	DefaultValue *string `pulumi:"defaultValue"`
	SourceField  *string `pulumi:"sourceField"`
}

type JsonFieldWithDefaultResponseOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonFieldWithDefaultResponse)(nil)).Elem()
}

func (o JsonFieldWithDefaultResponseOutput) ToJsonFieldWithDefaultResponseOutput() JsonFieldWithDefaultResponseOutput {
	return o
}

func (o JsonFieldWithDefaultResponseOutput) ToJsonFieldWithDefaultResponseOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponseOutput {
	return o
}

func (o JsonFieldWithDefaultResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefaultResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultResponseOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v JsonFieldWithDefaultResponse) *string { return v.SourceField }).(pulumi.StringPtrOutput)
}

type JsonFieldWithDefaultResponsePtrOutput struct{ *pulumi.OutputState }

func (JsonFieldWithDefaultResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonFieldWithDefaultResponse)(nil)).Elem()
}

func (o JsonFieldWithDefaultResponsePtrOutput) ToJsonFieldWithDefaultResponsePtrOutput() JsonFieldWithDefaultResponsePtrOutput {
	return o
}

func (o JsonFieldWithDefaultResponsePtrOutput) ToJsonFieldWithDefaultResponsePtrOutputWithContext(ctx context.Context) JsonFieldWithDefaultResponsePtrOutput {
	return o
}

func (o JsonFieldWithDefaultResponsePtrOutput) Elem() JsonFieldWithDefaultResponseOutput {
	return o.ApplyT(func(v *JsonFieldWithDefaultResponse) JsonFieldWithDefaultResponse {
		if v != nil {
			return *v
		}
		var ret JsonFieldWithDefaultResponse
		return ret
	}).(JsonFieldWithDefaultResponseOutput)
}

func (o JsonFieldWithDefaultResponsePtrOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefaultResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultValue
	}).(pulumi.StringPtrOutput)
}

func (o JsonFieldWithDefaultResponsePtrOutput) SourceField() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonFieldWithDefaultResponse) *string {
		if v == nil {
			return nil
		}
		return v.SourceField
	}).(pulumi.StringPtrOutput)
}

type JsonInputSchemaMapping struct {
	DataVersion            *JsonFieldWithDefault `pulumi:"dataVersion"`
	EventTime              *JsonField            `pulumi:"eventTime"`
	EventType              *JsonFieldWithDefault `pulumi:"eventType"`
	Id                     *JsonField            `pulumi:"id"`
	InputSchemaMappingType string                `pulumi:"inputSchemaMappingType"`
	Subject                *JsonFieldWithDefault `pulumi:"subject"`
	Topic                  *JsonField            `pulumi:"topic"`
}





type JsonInputSchemaMappingInput interface {
	pulumi.Input

	ToJsonInputSchemaMappingOutput() JsonInputSchemaMappingOutput
	ToJsonInputSchemaMappingOutputWithContext(context.Context) JsonInputSchemaMappingOutput
}

type JsonInputSchemaMappingArgs struct {
	DataVersion            JsonFieldWithDefaultPtrInput `pulumi:"dataVersion"`
	EventTime              JsonFieldPtrInput            `pulumi:"eventTime"`
	EventType              JsonFieldWithDefaultPtrInput `pulumi:"eventType"`
	Id                     JsonFieldPtrInput            `pulumi:"id"`
	InputSchemaMappingType pulumi.StringInput           `pulumi:"inputSchemaMappingType"`
	Subject                JsonFieldWithDefaultPtrInput `pulumi:"subject"`
	Topic                  JsonFieldPtrInput            `pulumi:"topic"`
}

func (JsonInputSchemaMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMapping)(nil)).Elem()
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingOutput() JsonInputSchemaMappingOutput {
	return i.ToJsonInputSchemaMappingOutputWithContext(context.Background())
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingOutputWithContext(ctx context.Context) JsonInputSchemaMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingOutput)
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return i.ToJsonInputSchemaMappingPtrOutputWithContext(context.Background())
}

func (i JsonInputSchemaMappingArgs) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingOutput).ToJsonInputSchemaMappingPtrOutputWithContext(ctx)
}









type JsonInputSchemaMappingPtrInput interface {
	pulumi.Input

	ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput
	ToJsonInputSchemaMappingPtrOutputWithContext(context.Context) JsonInputSchemaMappingPtrOutput
}

type jsonInputSchemaMappingPtrType JsonInputSchemaMappingArgs

func JsonInputSchemaMappingPtr(v *JsonInputSchemaMappingArgs) JsonInputSchemaMappingPtrInput {
	return (*jsonInputSchemaMappingPtrType)(v)
}

func (*jsonInputSchemaMappingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMapping)(nil)).Elem()
}

func (i *jsonInputSchemaMappingPtrType) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return i.ToJsonInputSchemaMappingPtrOutputWithContext(context.Background())
}

func (i *jsonInputSchemaMappingPtrType) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JsonInputSchemaMappingPtrOutput)
}

type JsonInputSchemaMappingOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMapping)(nil)).Elem()
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingOutput() JsonInputSchemaMappingOutput {
	return o
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingOutputWithContext(ctx context.Context) JsonInputSchemaMappingOutput {
	return o
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return o.ToJsonInputSchemaMappingPtrOutputWithContext(context.Background())
}

func (o JsonInputSchemaMappingOutput) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonInputSchemaMapping) *JsonInputSchemaMapping {
		return &v
	}).(JsonInputSchemaMappingPtrOutput)
}

func (o JsonInputSchemaMappingOutput) DataVersion() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonFieldWithDefault { return v.DataVersion }).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingOutput) EventTime() JsonFieldPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonField { return v.EventTime }).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingOutput) EventType() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonFieldWithDefault { return v.EventType }).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingOutput) Id() JsonFieldPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonField { return v.Id }).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingOutput) InputSchemaMappingType() pulumi.StringOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) string { return v.InputSchemaMappingType }).(pulumi.StringOutput)
}

func (o JsonInputSchemaMappingOutput) Subject() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonFieldWithDefault { return v.Subject }).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingOutput) Topic() JsonFieldPtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMapping) *JsonField { return v.Topic }).(JsonFieldPtrOutput)
}

type JsonInputSchemaMappingPtrOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMapping)(nil)).Elem()
}

func (o JsonInputSchemaMappingPtrOutput) ToJsonInputSchemaMappingPtrOutput() JsonInputSchemaMappingPtrOutput {
	return o
}

func (o JsonInputSchemaMappingPtrOutput) ToJsonInputSchemaMappingPtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingPtrOutput {
	return o
}

func (o JsonInputSchemaMappingPtrOutput) Elem() JsonInputSchemaMappingOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) JsonInputSchemaMapping {
		if v != nil {
			return *v
		}
		var ret JsonInputSchemaMapping
		return ret
	}).(JsonInputSchemaMappingOutput)
}

func (o JsonInputSchemaMappingPtrOutput) DataVersion() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonFieldWithDefault {
		if v == nil {
			return nil
		}
		return v.DataVersion
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) EventTime() JsonFieldPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonField {
		if v == nil {
			return nil
		}
		return v.EventTime
	}).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) EventType() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonFieldWithDefault {
		if v == nil {
			return nil
		}
		return v.EventType
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) Id() JsonFieldPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonField {
		if v == nil {
			return nil
		}
		return v.Id
	}).(JsonFieldPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) InputSchemaMappingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *string {
		if v == nil {
			return nil
		}
		return &v.InputSchemaMappingType
	}).(pulumi.StringPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) Subject() JsonFieldWithDefaultPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonFieldWithDefault {
		if v == nil {
			return nil
		}
		return v.Subject
	}).(JsonFieldWithDefaultPtrOutput)
}

func (o JsonInputSchemaMappingPtrOutput) Topic() JsonFieldPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMapping) *JsonField {
		if v == nil {
			return nil
		}
		return v.Topic
	}).(JsonFieldPtrOutput)
}

type JsonInputSchemaMappingResponse struct {
	DataVersion            *JsonFieldWithDefaultResponse `pulumi:"dataVersion"`
	EventTime              *JsonFieldResponse            `pulumi:"eventTime"`
	EventType              *JsonFieldWithDefaultResponse `pulumi:"eventType"`
	Id                     *JsonFieldResponse            `pulumi:"id"`
	InputSchemaMappingType string                        `pulumi:"inputSchemaMappingType"`
	Subject                *JsonFieldWithDefaultResponse `pulumi:"subject"`
	Topic                  *JsonFieldResponse            `pulumi:"topic"`
}

type JsonInputSchemaMappingResponseOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonInputSchemaMappingResponse)(nil)).Elem()
}

func (o JsonInputSchemaMappingResponseOutput) ToJsonInputSchemaMappingResponseOutput() JsonInputSchemaMappingResponseOutput {
	return o
}

func (o JsonInputSchemaMappingResponseOutput) ToJsonInputSchemaMappingResponseOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponseOutput {
	return o
}

func (o JsonInputSchemaMappingResponseOutput) DataVersion() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse { return v.DataVersion }).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) EventTime() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldResponse { return v.EventTime }).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) EventType() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse { return v.EventType }).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) Id() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldResponse { return v.Id }).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) InputSchemaMappingType() pulumi.StringOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) string { return v.InputSchemaMappingType }).(pulumi.StringOutput)
}

func (o JsonInputSchemaMappingResponseOutput) Subject() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse { return v.Subject }).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponseOutput) Topic() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v JsonInputSchemaMappingResponse) *JsonFieldResponse { return v.Topic }).(JsonFieldResponsePtrOutput)
}

type JsonInputSchemaMappingResponsePtrOutput struct{ *pulumi.OutputState }

func (JsonInputSchemaMappingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonInputSchemaMappingResponse)(nil)).Elem()
}

func (o JsonInputSchemaMappingResponsePtrOutput) ToJsonInputSchemaMappingResponsePtrOutput() JsonInputSchemaMappingResponsePtrOutput {
	return o
}

func (o JsonInputSchemaMappingResponsePtrOutput) ToJsonInputSchemaMappingResponsePtrOutputWithContext(ctx context.Context) JsonInputSchemaMappingResponsePtrOutput {
	return o
}

func (o JsonInputSchemaMappingResponsePtrOutput) Elem() JsonInputSchemaMappingResponseOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) JsonInputSchemaMappingResponse {
		if v != nil {
			return *v
		}
		var ret JsonInputSchemaMappingResponse
		return ret
	}).(JsonInputSchemaMappingResponseOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) DataVersion() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse {
		if v == nil {
			return nil
		}
		return v.DataVersion
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) EventTime() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldResponse {
		if v == nil {
			return nil
		}
		return v.EventTime
	}).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) EventType() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse {
		if v == nil {
			return nil
		}
		return v.EventType
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) Id() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldResponse {
		if v == nil {
			return nil
		}
		return v.Id
	}).(JsonFieldResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) InputSchemaMappingType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *string {
		if v == nil {
			return nil
		}
		return &v.InputSchemaMappingType
	}).(pulumi.StringPtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) Subject() JsonFieldWithDefaultResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldWithDefaultResponse {
		if v == nil {
			return nil
		}
		return v.Subject
	}).(JsonFieldWithDefaultResponsePtrOutput)
}

func (o JsonInputSchemaMappingResponsePtrOutput) Topic() JsonFieldResponsePtrOutput {
	return o.ApplyT(func(v *JsonInputSchemaMappingResponse) *JsonFieldResponse {
		if v == nil {
			return nil
		}
		return v.Topic
	}).(JsonFieldResponsePtrOutput)
}

type NumberGreaterThanAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberGreaterThanAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberGreaterThanOrEqualsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberGreaterThanOrEqualsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberInAdvancedFilter struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberInAdvancedFilterResponse struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberLessThanAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberLessThanAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberLessThanOrEqualsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberLessThanOrEqualsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Value        *float64 `pulumi:"value"`
}

type NumberNotInAdvancedFilter struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type NumberNotInAdvancedFilterResponse struct {
	Key          *string   `pulumi:"key"`
	OperatorType string    `pulumi:"operatorType"`
	Values       []float64 `pulumi:"values"`
}

type PrivateEndpoint struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointInput interface {
	pulumi.Input

	ToPrivateEndpointOutput() PrivateEndpointOutput
	ToPrivateEndpointOutputWithContext(context.Context) PrivateEndpointOutput
}

type PrivateEndpointArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return i.ToPrivateEndpointOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput)
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointArgs) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointOutput).ToPrivateEndpointPtrOutputWithContext(ctx)
}









type PrivateEndpointPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput
	ToPrivateEndpointPtrOutputWithContext(context.Context) PrivateEndpointPtrOutput
}

type privateEndpointPtrType PrivateEndpointArgs

func PrivateEndpointPtr(v *PrivateEndpointArgs) PrivateEndpointPtrInput {
	return (*privateEndpointPtrType)(v)
}

func (*privateEndpointPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return i.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPtrType) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPtrOutput)
}

type PrivateEndpointOutput struct{ *pulumi.OutputState }

func (PrivateEndpointOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutput() PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointOutputWithContext(ctx context.Context) PrivateEndpointOutput {
	return o
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o.ToPrivateEndpointPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpoint) *PrivateEndpoint {
		return &v
	}).(PrivateEndpointPtrOutput)
}

func (o PrivateEndpointOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpoint) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpoint)(nil)).Elem()
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutput() PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) ToPrivateEndpointPtrOutputWithContext(ctx context.Context) PrivateEndpointPtrOutput {
	return o
}

func (o PrivateEndpointPtrOutput) Elem() PrivateEndpointOutput {
	return o.ApplyT(func(v *PrivateEndpoint) PrivateEndpoint {
		if v != nil {
			return *v
		}
		var ret PrivateEndpoint
		return ret
	}).(PrivateEndpointOutput)
}

func (o PrivateEndpointPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpoint) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupIds                          []string                 `pulumi:"groupIds"`
	Id                                string                   `pulumi:"id"`
	Name                              string                   `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                  `pulumi:"provisioningState"`
	Type                              string                   `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) GroupIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) []string { return v.GroupIds }).(pulumi.StringArrayOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() ConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *ConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(ConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type RetryPolicy struct {
	EventTimeToLiveInMinutes *int `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      *int `pulumi:"maxDeliveryAttempts"`
}


func (val *RetryPolicy) Defaults() *RetryPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventTimeToLiveInMinutes) {
		eventTimeToLiveInMinutes_ := 1440
		tmp.EventTimeToLiveInMinutes = &eventTimeToLiveInMinutes_
	}
	if isZero(tmp.MaxDeliveryAttempts) {
		maxDeliveryAttempts_ := 30
		tmp.MaxDeliveryAttempts = &maxDeliveryAttempts_
	}
	return &tmp
}





type RetryPolicyInput interface {
	pulumi.Input

	ToRetryPolicyOutput() RetryPolicyOutput
	ToRetryPolicyOutputWithContext(context.Context) RetryPolicyOutput
}

type RetryPolicyArgs struct {
	EventTimeToLiveInMinutes pulumi.IntPtrInput `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      pulumi.IntPtrInput `pulumi:"maxDeliveryAttempts"`
}

func (RetryPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicy)(nil)).Elem()
}

func (i RetryPolicyArgs) ToRetryPolicyOutput() RetryPolicyOutput {
	return i.ToRetryPolicyOutputWithContext(context.Background())
}

func (i RetryPolicyArgs) ToRetryPolicyOutputWithContext(ctx context.Context) RetryPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyOutput)
}

func (i RetryPolicyArgs) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return i.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (i RetryPolicyArgs) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyOutput).ToRetryPolicyPtrOutputWithContext(ctx)
}









type RetryPolicyPtrInput interface {
	pulumi.Input

	ToRetryPolicyPtrOutput() RetryPolicyPtrOutput
	ToRetryPolicyPtrOutputWithContext(context.Context) RetryPolicyPtrOutput
}

type retryPolicyPtrType RetryPolicyArgs

func RetryPolicyPtr(v *RetryPolicyArgs) RetryPolicyPtrInput {
	return (*retryPolicyPtrType)(v)
}

func (*retryPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicy)(nil)).Elem()
}

func (i *retryPolicyPtrType) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return i.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (i *retryPolicyPtrType) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyPtrOutput)
}

type RetryPolicyOutput struct{ *pulumi.OutputState }

func (RetryPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicy)(nil)).Elem()
}

func (o RetryPolicyOutput) ToRetryPolicyOutput() RetryPolicyOutput {
	return o
}

func (o RetryPolicyOutput) ToRetryPolicyOutputWithContext(ctx context.Context) RetryPolicyOutput {
	return o
}

func (o RetryPolicyOutput) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return o.ToRetryPolicyPtrOutputWithContext(context.Background())
}

func (o RetryPolicyOutput) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetryPolicy) *RetryPolicy {
		return &v
	}).(RetryPolicyPtrOutput)
}

func (o RetryPolicyOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicy) *int { return v.EventTimeToLiveInMinutes }).(pulumi.IntPtrOutput)
}

func (o RetryPolicyOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicy) *int { return v.MaxDeliveryAttempts }).(pulumi.IntPtrOutput)
}

type RetryPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetryPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicy)(nil)).Elem()
}

func (o RetryPolicyPtrOutput) ToRetryPolicyPtrOutput() RetryPolicyPtrOutput {
	return o
}

func (o RetryPolicyPtrOutput) ToRetryPolicyPtrOutputWithContext(ctx context.Context) RetryPolicyPtrOutput {
	return o
}

func (o RetryPolicyPtrOutput) Elem() RetryPolicyOutput {
	return o.ApplyT(func(v *RetryPolicy) RetryPolicy {
		if v != nil {
			return *v
		}
		var ret RetryPolicy
		return ret
	}).(RetryPolicyOutput)
}

func (o RetryPolicyPtrOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *int {
		if v == nil {
			return nil
		}
		return v.EventTimeToLiveInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o RetryPolicyPtrOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryAttempts
	}).(pulumi.IntPtrOutput)
}

type RetryPolicyResponse struct {
	EventTimeToLiveInMinutes *int `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      *int `pulumi:"maxDeliveryAttempts"`
}


func (val *RetryPolicyResponse) Defaults() *RetryPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EventTimeToLiveInMinutes) {
		eventTimeToLiveInMinutes_ := 1440
		tmp.EventTimeToLiveInMinutes = &eventTimeToLiveInMinutes_
	}
	if isZero(tmp.MaxDeliveryAttempts) {
		maxDeliveryAttempts_ := 30
		tmp.MaxDeliveryAttempts = &maxDeliveryAttempts_
	}
	return &tmp
}

type RetryPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetryPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicyResponse)(nil)).Elem()
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponseOutput() RetryPolicyResponseOutput {
	return o
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponseOutputWithContext(ctx context.Context) RetryPolicyResponseOutput {
	return o
}

func (o RetryPolicyResponseOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *int { return v.EventTimeToLiveInMinutes }).(pulumi.IntPtrOutput)
}

func (o RetryPolicyResponseOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetryPolicyResponse) *int { return v.MaxDeliveryAttempts }).(pulumi.IntPtrOutput)
}

type RetryPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetryPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicyResponse)(nil)).Elem()
}

func (o RetryPolicyResponsePtrOutput) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return o
}

func (o RetryPolicyResponsePtrOutput) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return o
}

func (o RetryPolicyResponsePtrOutput) Elem() RetryPolicyResponseOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) RetryPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetryPolicyResponse
		return ret
	}).(RetryPolicyResponseOutput)
}

func (o RetryPolicyResponsePtrOutput) EventTimeToLiveInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.EventTimeToLiveInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o RetryPolicyResponsePtrOutput) MaxDeliveryAttempts() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetryPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDeliveryAttempts
	}).(pulumi.IntPtrOutput)
}

type ServiceBusQueueEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type ServiceBusQueueEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type ServiceBusTopicEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type ServiceBusTopicEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}

type StorageBlobDeadLetterDestination struct {
	BlobContainerName *string `pulumi:"blobContainerName"`
	EndpointType      string  `pulumi:"endpointType"`
	ResourceId        *string `pulumi:"resourceId"`
}





type StorageBlobDeadLetterDestinationInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationOutput() StorageBlobDeadLetterDestinationOutput
	ToStorageBlobDeadLetterDestinationOutputWithContext(context.Context) StorageBlobDeadLetterDestinationOutput
}

type StorageBlobDeadLetterDestinationArgs struct {
	BlobContainerName pulumi.StringPtrInput `pulumi:"blobContainerName"`
	EndpointType      pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId        pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageBlobDeadLetterDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationOutput() StorageBlobDeadLetterDestinationOutput {
	return i.ToStorageBlobDeadLetterDestinationOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationOutput)
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return i.ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationArgs) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationOutput).ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx)
}









type StorageBlobDeadLetterDestinationPtrInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput
	ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Context) StorageBlobDeadLetterDestinationPtrOutput
}

type storageBlobDeadLetterDestinationPtrType StorageBlobDeadLetterDestinationArgs

func StorageBlobDeadLetterDestinationPtr(v *StorageBlobDeadLetterDestinationArgs) StorageBlobDeadLetterDestinationPtrInput {
	return (*storageBlobDeadLetterDestinationPtrType)(v)
}

func (*storageBlobDeadLetterDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (i *storageBlobDeadLetterDestinationPtrType) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return i.ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Background())
}

func (i *storageBlobDeadLetterDestinationPtrType) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationPtrOutput)
}

type StorageBlobDeadLetterDestinationOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationOutput() StorageBlobDeadLetterDestinationOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return o.ToStorageBlobDeadLetterDestinationPtrOutputWithContext(context.Background())
}

func (o StorageBlobDeadLetterDestinationOutput) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBlobDeadLetterDestination) *StorageBlobDeadLetterDestination {
		return &v
	}).(StorageBlobDeadLetterDestinationPtrOutput)
}

func (o StorageBlobDeadLetterDestinationOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestination) *string { return v.BlobContainerName }).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageBlobDeadLetterDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestinationPtrOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestination)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationPtrOutput) ToStorageBlobDeadLetterDestinationPtrOutput() StorageBlobDeadLetterDestinationPtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationPtrOutput) ToStorageBlobDeadLetterDestinationPtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationPtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationPtrOutput) Elem() StorageBlobDeadLetterDestinationOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) StorageBlobDeadLetterDestination {
		if v != nil {
			return *v
		}
		var ret StorageBlobDeadLetterDestination
		return ret
	}).(StorageBlobDeadLetterDestinationOutput)
}

func (o StorageBlobDeadLetterDestinationPtrOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainerName
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationPtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationPtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestination) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestinationResponse struct {
	BlobContainerName *string `pulumi:"blobContainerName"`
	EndpointType      string  `pulumi:"endpointType"`
	ResourceId        *string `pulumi:"resourceId"`
}

type StorageBlobDeadLetterDestinationResponseOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponseOutput() StorageBlobDeadLetterDestinationResponseOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponseOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponseOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponseOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestinationResponse) *string { return v.BlobContainerName }).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageBlobDeadLetterDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageBlobDeadLetterDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageBlobDeadLetterDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) Elem() StorageBlobDeadLetterDestinationResponseOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) StorageBlobDeadLetterDestinationResponse {
		if v != nil {
			return *v
		}
		var ret StorageBlobDeadLetterDestinationResponse
		return ret
	}).(StorageBlobDeadLetterDestinationResponseOutput)
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) BlobContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.BlobContainerName
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) EndpointType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointType
	}).(pulumi.StringPtrOutput)
}

func (o StorageBlobDeadLetterDestinationResponsePtrOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageBlobDeadLetterDestinationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ResourceId
	}).(pulumi.StringPtrOutput)
}

type StorageQueueEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	QueueName    *string `pulumi:"queueName"`
	ResourceId   *string `pulumi:"resourceId"`
}

type StorageQueueEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	QueueName    *string `pulumi:"queueName"`
	ResourceId   *string `pulumi:"resourceId"`
}

type StringBeginsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringBeginsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringContainsAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringContainsAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringEndsWithAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringEndsWithAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringInAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringInAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotInAdvancedFilter struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
}

type StringNotInAdvancedFilterResponse struct {
	Key          *string  `pulumi:"key"`
	OperatorType string   `pulumi:"operatorType"`
	Values       []string `pulumi:"values"`
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

type WebHookEventSubscriptionDestination struct {
	AzureActiveDirectoryApplicationIdOrUri *string `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           *string `pulumi:"azureActiveDirectoryTenantId"`
	EndpointType                           string  `pulumi:"endpointType"`
	EndpointUrl                            *string `pulumi:"endpointUrl"`
	MaxEventsPerBatch                      *int    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes          *int    `pulumi:"preferredBatchSizeInKilobytes"`
}


func (val *WebHookEventSubscriptionDestination) Defaults() *WebHookEventSubscriptionDestination {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}

type WebHookEventSubscriptionDestinationResponse struct {
	AzureActiveDirectoryApplicationIdOrUri *string `pulumi:"azureActiveDirectoryApplicationIdOrUri"`
	AzureActiveDirectoryTenantId           *string `pulumi:"azureActiveDirectoryTenantId"`
	EndpointBaseUrl                        string  `pulumi:"endpointBaseUrl"`
	EndpointType                           string  `pulumi:"endpointType"`
	EndpointUrl                            *string `pulumi:"endpointUrl"`
	MaxEventsPerBatch                      *int    `pulumi:"maxEventsPerBatch"`
	PreferredBatchSizeInKilobytes          *int    `pulumi:"preferredBatchSizeInKilobytes"`
}


func (val *WebHookEventSubscriptionDestinationResponse) Defaults() *WebHookEventSubscriptionDestinationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.MaxEventsPerBatch) {
		maxEventsPerBatch_ := 1
		tmp.MaxEventsPerBatch = &maxEventsPerBatch_
	}
	if isZero(tmp.PreferredBatchSizeInKilobytes) {
		preferredBatchSizeInKilobytes_ := 64
		tmp.PreferredBatchSizeInKilobytes = &preferredBatchSizeInKilobytes_
	}
	return &tmp
}
func init() {
	pulumi.RegisterOutputType(ConnectionStateOutput{})
	pulumi.RegisterOutputType(ConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(ConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(InboundIpRuleOutput{})
	pulumi.RegisterOutputType(InboundIpRuleArrayOutput{})
	pulumi.RegisterOutputType(InboundIpRuleResponseOutput{})
	pulumi.RegisterOutputType(InboundIpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(JsonFieldOutput{})
	pulumi.RegisterOutputType(JsonFieldPtrOutput{})
	pulumi.RegisterOutputType(JsonFieldResponseOutput{})
	pulumi.RegisterOutputType(JsonFieldResponsePtrOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultPtrOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultResponseOutput{})
	pulumi.RegisterOutputType(JsonFieldWithDefaultResponsePtrOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingPtrOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingResponseOutput{})
	pulumi.RegisterOutputType(JsonInputSchemaMappingResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyOutput{})
	pulumi.RegisterOutputType(RetryPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationPtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponseOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
