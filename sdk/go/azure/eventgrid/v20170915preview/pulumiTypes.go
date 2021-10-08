


package v20170915preview

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
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubEventSubscriptionDestinationInput)(nil)).Elem(), EventHubEventSubscriptionDestinationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventHubEventSubscriptionDestinationResponseInput)(nil)).Elem(), EventHubEventSubscriptionDestinationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionFilterInput)(nil)).Elem(), EventSubscriptionFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionFilterPtrInput)(nil)).Elem(), EventSubscriptionFilterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionFilterResponseInput)(nil)).Elem(), EventSubscriptionFilterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionFilterResponsePtrInput)(nil)).Elem(), EventSubscriptionFilterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebHookEventSubscriptionDestinationInput)(nil)).Elem(), WebHookEventSubscriptionDestinationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebHookEventSubscriptionDestinationResponseInput)(nil)).Elem(), WebHookEventSubscriptionDestinationResponseArgs{})
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(EventHubEventSubscriptionDestinationResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponseOutput{})
	pulumi.RegisterOutputType(EventSubscriptionFilterResponsePtrOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationOutput{})
	pulumi.RegisterOutputType(WebHookEventSubscriptionDestinationResponseOutput{})
}
