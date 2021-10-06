


package v20190101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type EventHubEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}





type EventHubEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput
	ToEventHubEventSubscriptionDestinationOutputWithContext(context.Context) EventHubEventSubscriptionDestinationOutput
}

type EventHubEventSubscriptionDestinationArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (EventHubEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestination)(nil)).Elem()
}

func (i EventHubEventSubscriptionDestinationArgs) ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput {
	return i.ToEventHubEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i EventHubEventSubscriptionDestinationArgs) ToEventHubEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSubscriptionDestinationOutput)
}

type EventHubEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (EventHubEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestination)(nil)).Elem()
}

func (o EventHubEventSubscriptionDestinationOutput) ToEventHubEventSubscriptionDestinationOutput() EventHubEventSubscriptionDestinationOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationOutput) ToEventHubEventSubscriptionDestinationOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type EventHubEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}





type EventHubEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput
	ToEventHubEventSubscriptionDestinationResponseOutputWithContext(context.Context) EventHubEventSubscriptionDestinationResponseOutput
}

type EventHubEventSubscriptionDestinationResponseArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (EventHubEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i EventHubEventSubscriptionDestinationResponseArgs) ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput {
	return i.ToEventHubEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i EventHubEventSubscriptionDestinationResponseArgs) ToEventHubEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventHubEventSubscriptionDestinationResponseOutput)
}

type EventHubEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (EventHubEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventHubEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ToEventHubEventSubscriptionDestinationResponseOutput() EventHubEventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ToEventHubEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) EventHubEventSubscriptionDestinationResponseOutput {
	return o
}

func (o EventHubEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o EventHubEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventHubEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type EventSubscriptionFilter struct {
	IncludedEventTypes     []string `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive *bool    `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      *string  `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        *string  `pulumi:"subjectEndsWith"`
}





type EventSubscriptionFilterInput interface {
	pulumi.Input

	ToEventSubscriptionFilterOutput() EventSubscriptionFilterOutput
	ToEventSubscriptionFilterOutputWithContext(context.Context) EventSubscriptionFilterOutput
}

type EventSubscriptionFilterArgs struct {
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
	IncludedEventTypes     []string `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive *bool    `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      *string  `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        *string  `pulumi:"subjectEndsWith"`
}





type EventSubscriptionFilterResponseInput interface {
	pulumi.Input

	ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput
	ToEventSubscriptionFilterResponseOutputWithContext(context.Context) EventSubscriptionFilterResponseOutput
}

type EventSubscriptionFilterResponseArgs struct {
	IncludedEventTypes     pulumi.StringArrayInput `pulumi:"includedEventTypes"`
	IsSubjectCaseSensitive pulumi.BoolPtrInput     `pulumi:"isSubjectCaseSensitive"`
	SubjectBeginsWith      pulumi.StringPtrInput   `pulumi:"subjectBeginsWith"`
	SubjectEndsWith        pulumi.StringPtrInput   `pulumi:"subjectEndsWith"`
}

func (EventSubscriptionFilterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionFilterResponse)(nil)).Elem()
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponseOutput() EventSubscriptionFilterResponseOutput {
	return i.ToEventSubscriptionFilterResponseOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponseOutputWithContext(ctx context.Context) EventSubscriptionFilterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterResponseOutput)
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return i.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (i EventSubscriptionFilterResponseArgs) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterResponseOutput).ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx)
}









type EventSubscriptionFilterResponsePtrInput interface {
	pulumi.Input

	ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput
	ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Context) EventSubscriptionFilterResponsePtrOutput
}

type eventSubscriptionFilterResponsePtrType EventSubscriptionFilterResponseArgs

func EventSubscriptionFilterResponsePtr(v *EventSubscriptionFilterResponseArgs) EventSubscriptionFilterResponsePtrInput {
	return (*eventSubscriptionFilterResponsePtrType)(v)
}

func (*eventSubscriptionFilterResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionFilterResponse)(nil)).Elem()
}

func (i *eventSubscriptionFilterResponsePtrType) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return i.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (i *eventSubscriptionFilterResponsePtrType) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventSubscriptionFilterResponsePtrOutput)
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

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponsePtrOutput() EventSubscriptionFilterResponsePtrOutput {
	return o.ToEventSubscriptionFilterResponsePtrOutputWithContext(context.Background())
}

func (o EventSubscriptionFilterResponseOutput) ToEventSubscriptionFilterResponsePtrOutputWithContext(ctx context.Context) EventSubscriptionFilterResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionFilterResponse) *EventSubscriptionFilterResponse {
		return &v
	}).(EventSubscriptionFilterResponsePtrOutput)
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





type HybridConnectionEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToHybridConnectionEventSubscriptionDestinationOutput() HybridConnectionEventSubscriptionDestinationOutput
	ToHybridConnectionEventSubscriptionDestinationOutputWithContext(context.Context) HybridConnectionEventSubscriptionDestinationOutput
}

type HybridConnectionEventSubscriptionDestinationArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (HybridConnectionEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestination)(nil)).Elem()
}

func (i HybridConnectionEventSubscriptionDestinationArgs) ToHybridConnectionEventSubscriptionDestinationOutput() HybridConnectionEventSubscriptionDestinationOutput {
	return i.ToHybridConnectionEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i HybridConnectionEventSubscriptionDestinationArgs) ToHybridConnectionEventSubscriptionDestinationOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionEventSubscriptionDestinationOutput)
}

type HybridConnectionEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (HybridConnectionEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestination)(nil)).Elem()
}

func (o HybridConnectionEventSubscriptionDestinationOutput) ToHybridConnectionEventSubscriptionDestinationOutput() HybridConnectionEventSubscriptionDestinationOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationOutput) ToHybridConnectionEventSubscriptionDestinationOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o HybridConnectionEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type HybridConnectionEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	ResourceId   *string `pulumi:"resourceId"`
}





type HybridConnectionEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToHybridConnectionEventSubscriptionDestinationResponseOutput() HybridConnectionEventSubscriptionDestinationResponseOutput
	ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(context.Context) HybridConnectionEventSubscriptionDestinationResponseOutput
}

type HybridConnectionEventSubscriptionDestinationResponseArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (HybridConnectionEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i HybridConnectionEventSubscriptionDestinationResponseArgs) ToHybridConnectionEventSubscriptionDestinationResponseOutput() HybridConnectionEventSubscriptionDestinationResponseOutput {
	return i.ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i HybridConnectionEventSubscriptionDestinationResponseArgs) ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridConnectionEventSubscriptionDestinationResponseOutput)
}

type HybridConnectionEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (HybridConnectionEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridConnectionEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) ToHybridConnectionEventSubscriptionDestinationResponseOutput() HybridConnectionEventSubscriptionDestinationResponseOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) ToHybridConnectionEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) HybridConnectionEventSubscriptionDestinationResponseOutput {
	return o
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o HybridConnectionEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridConnectionEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type RetryPolicy struct {
	EventTimeToLiveInMinutes *int `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      *int `pulumi:"maxDeliveryAttempts"`
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





type RetryPolicyResponseInput interface {
	pulumi.Input

	ToRetryPolicyResponseOutput() RetryPolicyResponseOutput
	ToRetryPolicyResponseOutputWithContext(context.Context) RetryPolicyResponseOutput
}

type RetryPolicyResponseArgs struct {
	EventTimeToLiveInMinutes pulumi.IntPtrInput `pulumi:"eventTimeToLiveInMinutes"`
	MaxDeliveryAttempts      pulumi.IntPtrInput `pulumi:"maxDeliveryAttempts"`
}

func (RetryPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetryPolicyResponse)(nil)).Elem()
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponseOutput() RetryPolicyResponseOutput {
	return i.ToRetryPolicyResponseOutputWithContext(context.Background())
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponseOutputWithContext(ctx context.Context) RetryPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponseOutput)
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return i.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RetryPolicyResponseArgs) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponseOutput).ToRetryPolicyResponsePtrOutputWithContext(ctx)
}









type RetryPolicyResponsePtrInput interface {
	pulumi.Input

	ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput
	ToRetryPolicyResponsePtrOutputWithContext(context.Context) RetryPolicyResponsePtrOutput
}

type retryPolicyResponsePtrType RetryPolicyResponseArgs

func RetryPolicyResponsePtr(v *RetryPolicyResponseArgs) RetryPolicyResponsePtrInput {
	return (*retryPolicyResponsePtrType)(v)
}

func (*retryPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetryPolicyResponse)(nil)).Elem()
}

func (i *retryPolicyResponsePtrType) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return i.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *retryPolicyResponsePtrType) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetryPolicyResponsePtrOutput)
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

func (o RetryPolicyResponseOutput) ToRetryPolicyResponsePtrOutput() RetryPolicyResponsePtrOutput {
	return o.ToRetryPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RetryPolicyResponseOutput) ToRetryPolicyResponsePtrOutputWithContext(ctx context.Context) RetryPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetryPolicyResponse) *RetryPolicyResponse {
		return &v
	}).(RetryPolicyResponsePtrOutput)
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





type StorageBlobDeadLetterDestinationResponseInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationResponseOutput() StorageBlobDeadLetterDestinationResponseOutput
	ToStorageBlobDeadLetterDestinationResponseOutputWithContext(context.Context) StorageBlobDeadLetterDestinationResponseOutput
}

type StorageBlobDeadLetterDestinationResponseArgs struct {
	BlobContainerName pulumi.StringPtrInput `pulumi:"blobContainerName"`
	EndpointType      pulumi.StringInput    `pulumi:"endpointType"`
	ResourceId        pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageBlobDeadLetterDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponseOutput() StorageBlobDeadLetterDestinationResponseOutput {
	return i.ToStorageBlobDeadLetterDestinationResponseOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponseOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationResponseOutput)
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return i.ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Background())
}

func (i StorageBlobDeadLetterDestinationResponseArgs) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationResponseOutput).ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx)
}









type StorageBlobDeadLetterDestinationResponsePtrInput interface {
	pulumi.Input

	ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput
	ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput
}

type storageBlobDeadLetterDestinationResponsePtrType StorageBlobDeadLetterDestinationResponseArgs

func StorageBlobDeadLetterDestinationResponsePtr(v *StorageBlobDeadLetterDestinationResponseArgs) StorageBlobDeadLetterDestinationResponsePtrInput {
	return (*storageBlobDeadLetterDestinationResponsePtrType)(v)
}

func (*storageBlobDeadLetterDestinationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageBlobDeadLetterDestinationResponse)(nil)).Elem()
}

func (i *storageBlobDeadLetterDestinationResponsePtrType) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return i.ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Background())
}

func (i *storageBlobDeadLetterDestinationResponsePtrType) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageBlobDeadLetterDestinationResponsePtrOutput)
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

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutput() StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(context.Background())
}

func (o StorageBlobDeadLetterDestinationResponseOutput) ToStorageBlobDeadLetterDestinationResponsePtrOutputWithContext(ctx context.Context) StorageBlobDeadLetterDestinationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageBlobDeadLetterDestinationResponse) *StorageBlobDeadLetterDestinationResponse {
		return &v
	}).(StorageBlobDeadLetterDestinationResponsePtrOutput)
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





type StorageQueueEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToStorageQueueEventSubscriptionDestinationOutput() StorageQueueEventSubscriptionDestinationOutput
	ToStorageQueueEventSubscriptionDestinationOutputWithContext(context.Context) StorageQueueEventSubscriptionDestinationOutput
}

type StorageQueueEventSubscriptionDestinationArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	QueueName    pulumi.StringPtrInput `pulumi:"queueName"`
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageQueueEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestination)(nil)).Elem()
}

func (i StorageQueueEventSubscriptionDestinationArgs) ToStorageQueueEventSubscriptionDestinationOutput() StorageQueueEventSubscriptionDestinationOutput {
	return i.ToStorageQueueEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i StorageQueueEventSubscriptionDestinationArgs) ToStorageQueueEventSubscriptionDestinationOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueEventSubscriptionDestinationOutput)
}

type StorageQueueEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (StorageQueueEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestination)(nil)).Elem()
}

func (o StorageQueueEventSubscriptionDestinationOutput) ToStorageQueueEventSubscriptionDestinationOutput() StorageQueueEventSubscriptionDestinationOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationOutput) ToStorageQueueEventSubscriptionDestinationOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageQueueEventSubscriptionDestinationOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o StorageQueueEventSubscriptionDestinationOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestination) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type StorageQueueEventSubscriptionDestinationResponse struct {
	EndpointType string  `pulumi:"endpointType"`
	QueueName    *string `pulumi:"queueName"`
	ResourceId   *string `pulumi:"resourceId"`
}





type StorageQueueEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToStorageQueueEventSubscriptionDestinationResponseOutput() StorageQueueEventSubscriptionDestinationResponseOutput
	ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(context.Context) StorageQueueEventSubscriptionDestinationResponseOutput
}

type StorageQueueEventSubscriptionDestinationResponseArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	QueueName    pulumi.StringPtrInput `pulumi:"queueName"`
	ResourceId   pulumi.StringPtrInput `pulumi:"resourceId"`
}

func (StorageQueueEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i StorageQueueEventSubscriptionDestinationResponseArgs) ToStorageQueueEventSubscriptionDestinationResponseOutput() StorageQueueEventSubscriptionDestinationResponseOutput {
	return i.ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i StorageQueueEventSubscriptionDestinationResponseArgs) ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageQueueEventSubscriptionDestinationResponseOutput)
}

type StorageQueueEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (StorageQueueEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageQueueEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) ToStorageQueueEventSubscriptionDestinationResponseOutput() StorageQueueEventSubscriptionDestinationResponseOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) ToStorageQueueEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) StorageQueueEventSubscriptionDestinationResponseOutput {
	return o
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

func (o StorageQueueEventSubscriptionDestinationResponseOutput) ResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageQueueEventSubscriptionDestinationResponse) *string { return v.ResourceId }).(pulumi.StringPtrOutput)
}

type WebHookEventSubscriptionDestination struct {
	EndpointType string  `pulumi:"endpointType"`
	EndpointUrl  *string `pulumi:"endpointUrl"`
}





type WebHookEventSubscriptionDestinationInput interface {
	pulumi.Input

	ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput
	ToWebHookEventSubscriptionDestinationOutputWithContext(context.Context) WebHookEventSubscriptionDestinationOutput
}

type WebHookEventSubscriptionDestinationArgs struct {
	EndpointType pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUrl  pulumi.StringPtrInput `pulumi:"endpointUrl"`
}

func (WebHookEventSubscriptionDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestination)(nil)).Elem()
}

func (i WebHookEventSubscriptionDestinationArgs) ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput {
	return i.ToWebHookEventSubscriptionDestinationOutputWithContext(context.Background())
}

func (i WebHookEventSubscriptionDestinationArgs) ToWebHookEventSubscriptionDestinationOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebHookEventSubscriptionDestinationOutput)
}

type WebHookEventSubscriptionDestinationOutput struct{ *pulumi.OutputState }

func (WebHookEventSubscriptionDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestination)(nil)).Elem()
}

func (o WebHookEventSubscriptionDestinationOutput) ToWebHookEventSubscriptionDestinationOutput() WebHookEventSubscriptionDestinationOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationOutput) ToWebHookEventSubscriptionDestinationOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o WebHookEventSubscriptionDestinationOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestination) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

type WebHookEventSubscriptionDestinationResponse struct {
	EndpointBaseUrl string  `pulumi:"endpointBaseUrl"`
	EndpointType    string  `pulumi:"endpointType"`
	EndpointUrl     *string `pulumi:"endpointUrl"`
}





type WebHookEventSubscriptionDestinationResponseInput interface {
	pulumi.Input

	ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput
	ToWebHookEventSubscriptionDestinationResponseOutputWithContext(context.Context) WebHookEventSubscriptionDestinationResponseOutput
}

type WebHookEventSubscriptionDestinationResponseArgs struct {
	EndpointBaseUrl pulumi.StringInput    `pulumi:"endpointBaseUrl"`
	EndpointType    pulumi.StringInput    `pulumi:"endpointType"`
	EndpointUrl     pulumi.StringPtrInput `pulumi:"endpointUrl"`
}

func (WebHookEventSubscriptionDestinationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (i WebHookEventSubscriptionDestinationResponseArgs) ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput {
	return i.ToWebHookEventSubscriptionDestinationResponseOutputWithContext(context.Background())
}

func (i WebHookEventSubscriptionDestinationResponseArgs) ToWebHookEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebHookEventSubscriptionDestinationResponseOutput)
}

type WebHookEventSubscriptionDestinationResponseOutput struct{ *pulumi.OutputState }

func (WebHookEventSubscriptionDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebHookEventSubscriptionDestinationResponse)(nil)).Elem()
}

func (o WebHookEventSubscriptionDestinationResponseOutput) ToWebHookEventSubscriptionDestinationResponseOutput() WebHookEventSubscriptionDestinationResponseOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationResponseOutput) ToWebHookEventSubscriptionDestinationResponseOutputWithContext(ctx context.Context) WebHookEventSubscriptionDestinationResponseOutput {
	return o
}

func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointBaseUrl() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) string { return v.EndpointBaseUrl }).(pulumi.StringOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointType() pulumi.StringOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) string { return v.EndpointType }).(pulumi.StringOutput)
}

func (o WebHookEventSubscriptionDestinationResponseOutput) EndpointUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WebHookEventSubscriptionDestinationResponse) *string { return v.EndpointUrl }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(HybridConnectionEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(HybridConnectionEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(RetryPolicyOutput{})
	pulumi.RegisterOutputType(RetryPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetryPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationPtrOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponseOutput{})
	pulumi.RegisterOutputType(StorageBlobDeadLetterDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageQueueEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(StorageQueueEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationResponseOutput{})
}
