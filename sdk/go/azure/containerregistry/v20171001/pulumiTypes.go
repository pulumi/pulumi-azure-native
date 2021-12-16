


package v20171001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActorResponse struct {
	Name *string `pulumi:"name"`
}





type ActorResponseInput interface {
	pulumi.Input

	ToActorResponseOutput() ActorResponseOutput
	ToActorResponseOutputWithContext(context.Context) ActorResponseOutput
}

type ActorResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ActorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActorResponse)(nil)).Elem()
}

func (i ActorResponseArgs) ToActorResponseOutput() ActorResponseOutput {
	return i.ToActorResponseOutputWithContext(context.Background())
}

func (i ActorResponseArgs) ToActorResponseOutputWithContext(ctx context.Context) ActorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActorResponseOutput)
}

func (i ActorResponseArgs) ToActorResponsePtrOutput() ActorResponsePtrOutput {
	return i.ToActorResponsePtrOutputWithContext(context.Background())
}

func (i ActorResponseArgs) ToActorResponsePtrOutputWithContext(ctx context.Context) ActorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActorResponseOutput).ToActorResponsePtrOutputWithContext(ctx)
}









type ActorResponsePtrInput interface {
	pulumi.Input

	ToActorResponsePtrOutput() ActorResponsePtrOutput
	ToActorResponsePtrOutputWithContext(context.Context) ActorResponsePtrOutput
}

type actorResponsePtrType ActorResponseArgs

func ActorResponsePtr(v *ActorResponseArgs) ActorResponsePtrInput {
	return (*actorResponsePtrType)(v)
}

func (*actorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActorResponse)(nil)).Elem()
}

func (i *actorResponsePtrType) ToActorResponsePtrOutput() ActorResponsePtrOutput {
	return i.ToActorResponsePtrOutputWithContext(context.Background())
}

func (i *actorResponsePtrType) ToActorResponsePtrOutputWithContext(ctx context.Context) ActorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActorResponsePtrOutput)
}

type ActorResponseOutput struct{ *pulumi.OutputState }

func (ActorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActorResponse)(nil)).Elem()
}

func (o ActorResponseOutput) ToActorResponseOutput() ActorResponseOutput {
	return o
}

func (o ActorResponseOutput) ToActorResponseOutputWithContext(ctx context.Context) ActorResponseOutput {
	return o
}

func (o ActorResponseOutput) ToActorResponsePtrOutput() ActorResponsePtrOutput {
	return o.ToActorResponsePtrOutputWithContext(context.Background())
}

func (o ActorResponseOutput) ToActorResponsePtrOutputWithContext(ctx context.Context) ActorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActorResponse) *ActorResponse {
		return &v
	}).(ActorResponsePtrOutput)
}

func (o ActorResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ActorResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ActorResponsePtrOutput struct{ *pulumi.OutputState }

func (ActorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActorResponse)(nil)).Elem()
}

func (o ActorResponsePtrOutput) ToActorResponsePtrOutput() ActorResponsePtrOutput {
	return o
}

func (o ActorResponsePtrOutput) ToActorResponsePtrOutputWithContext(ctx context.Context) ActorResponsePtrOutput {
	return o
}

func (o ActorResponsePtrOutput) Elem() ActorResponseOutput {
	return o.ApplyT(func(v *ActorResponse) ActorResponse {
		if v != nil {
			return *v
		}
		var ret ActorResponse
		return ret
	}).(ActorResponseOutput)
}

func (o ActorResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type EventContentResponse struct {
	Action    *string          `pulumi:"action"`
	Actor     *ActorResponse   `pulumi:"actor"`
	Id        *string          `pulumi:"id"`
	Request   *RequestResponse `pulumi:"request"`
	Source    *SourceResponse  `pulumi:"source"`
	Target    *TargetResponse  `pulumi:"target"`
	Timestamp *string          `pulumi:"timestamp"`
}





type EventContentResponseInput interface {
	pulumi.Input

	ToEventContentResponseOutput() EventContentResponseOutput
	ToEventContentResponseOutputWithContext(context.Context) EventContentResponseOutput
}

type EventContentResponseArgs struct {
	Action    pulumi.StringPtrInput   `pulumi:"action"`
	Actor     ActorResponsePtrInput   `pulumi:"actor"`
	Id        pulumi.StringPtrInput   `pulumi:"id"`
	Request   RequestResponsePtrInput `pulumi:"request"`
	Source    SourceResponsePtrInput  `pulumi:"source"`
	Target    TargetResponsePtrInput  `pulumi:"target"`
	Timestamp pulumi.StringPtrInput   `pulumi:"timestamp"`
}

func (EventContentResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventContentResponse)(nil)).Elem()
}

func (i EventContentResponseArgs) ToEventContentResponseOutput() EventContentResponseOutput {
	return i.ToEventContentResponseOutputWithContext(context.Background())
}

func (i EventContentResponseArgs) ToEventContentResponseOutputWithContext(ctx context.Context) EventContentResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventContentResponseOutput)
}

func (i EventContentResponseArgs) ToEventContentResponsePtrOutput() EventContentResponsePtrOutput {
	return i.ToEventContentResponsePtrOutputWithContext(context.Background())
}

func (i EventContentResponseArgs) ToEventContentResponsePtrOutputWithContext(ctx context.Context) EventContentResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventContentResponseOutput).ToEventContentResponsePtrOutputWithContext(ctx)
}









type EventContentResponsePtrInput interface {
	pulumi.Input

	ToEventContentResponsePtrOutput() EventContentResponsePtrOutput
	ToEventContentResponsePtrOutputWithContext(context.Context) EventContentResponsePtrOutput
}

type eventContentResponsePtrType EventContentResponseArgs

func EventContentResponsePtr(v *EventContentResponseArgs) EventContentResponsePtrInput {
	return (*eventContentResponsePtrType)(v)
}

func (*eventContentResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventContentResponse)(nil)).Elem()
}

func (i *eventContentResponsePtrType) ToEventContentResponsePtrOutput() EventContentResponsePtrOutput {
	return i.ToEventContentResponsePtrOutputWithContext(context.Background())
}

func (i *eventContentResponsePtrType) ToEventContentResponsePtrOutputWithContext(ctx context.Context) EventContentResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventContentResponsePtrOutput)
}

type EventContentResponseOutput struct{ *pulumi.OutputState }

func (EventContentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventContentResponse)(nil)).Elem()
}

func (o EventContentResponseOutput) ToEventContentResponseOutput() EventContentResponseOutput {
	return o
}

func (o EventContentResponseOutput) ToEventContentResponseOutputWithContext(ctx context.Context) EventContentResponseOutput {
	return o
}

func (o EventContentResponseOutput) ToEventContentResponsePtrOutput() EventContentResponsePtrOutput {
	return o.ToEventContentResponsePtrOutputWithContext(context.Background())
}

func (o EventContentResponseOutput) ToEventContentResponsePtrOutputWithContext(ctx context.Context) EventContentResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventContentResponse) *EventContentResponse {
		return &v
	}).(EventContentResponsePtrOutput)
}

func (o EventContentResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventContentResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o EventContentResponseOutput) Actor() ActorResponsePtrOutput {
	return o.ApplyT(func(v EventContentResponse) *ActorResponse { return v.Actor }).(ActorResponsePtrOutput)
}

func (o EventContentResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventContentResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o EventContentResponseOutput) Request() RequestResponsePtrOutput {
	return o.ApplyT(func(v EventContentResponse) *RequestResponse { return v.Request }).(RequestResponsePtrOutput)
}

func (o EventContentResponseOutput) Source() SourceResponsePtrOutput {
	return o.ApplyT(func(v EventContentResponse) *SourceResponse { return v.Source }).(SourceResponsePtrOutput)
}

func (o EventContentResponseOutput) Target() TargetResponsePtrOutput {
	return o.ApplyT(func(v EventContentResponse) *TargetResponse { return v.Target }).(TargetResponsePtrOutput)
}

func (o EventContentResponseOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventContentResponse) *string { return v.Timestamp }).(pulumi.StringPtrOutput)
}

type EventContentResponsePtrOutput struct{ *pulumi.OutputState }

func (EventContentResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventContentResponse)(nil)).Elem()
}

func (o EventContentResponsePtrOutput) ToEventContentResponsePtrOutput() EventContentResponsePtrOutput {
	return o
}

func (o EventContentResponsePtrOutput) ToEventContentResponsePtrOutputWithContext(ctx context.Context) EventContentResponsePtrOutput {
	return o
}

func (o EventContentResponsePtrOutput) Elem() EventContentResponseOutput {
	return o.ApplyT(func(v *EventContentResponse) EventContentResponse {
		if v != nil {
			return *v
		}
		var ret EventContentResponse
		return ret
	}).(EventContentResponseOutput)
}

func (o EventContentResponsePtrOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *string {
		if v == nil {
			return nil
		}
		return v.Action
	}).(pulumi.StringPtrOutput)
}

func (o EventContentResponsePtrOutput) Actor() ActorResponsePtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *ActorResponse {
		if v == nil {
			return nil
		}
		return v.Actor
	}).(ActorResponsePtrOutput)
}

func (o EventContentResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o EventContentResponsePtrOutput) Request() RequestResponsePtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *RequestResponse {
		if v == nil {
			return nil
		}
		return v.Request
	}).(RequestResponsePtrOutput)
}

func (o EventContentResponsePtrOutput) Source() SourceResponsePtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *SourceResponse {
		if v == nil {
			return nil
		}
		return v.Source
	}).(SourceResponsePtrOutput)
}

func (o EventContentResponsePtrOutput) Target() TargetResponsePtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *TargetResponse {
		if v == nil {
			return nil
		}
		return v.Target
	}).(TargetResponsePtrOutput)
}

func (o EventContentResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventContentResponse) *string {
		if v == nil {
			return nil
		}
		return v.Timestamp
	}).(pulumi.StringPtrOutput)
}

type EventRequestMessageResponse struct {
	Content    *EventContentResponse `pulumi:"content"`
	Headers    map[string]string     `pulumi:"headers"`
	Method     *string               `pulumi:"method"`
	RequestUri *string               `pulumi:"requestUri"`
	Version    *string               `pulumi:"version"`
}





type EventRequestMessageResponseInput interface {
	pulumi.Input

	ToEventRequestMessageResponseOutput() EventRequestMessageResponseOutput
	ToEventRequestMessageResponseOutputWithContext(context.Context) EventRequestMessageResponseOutput
}

type EventRequestMessageResponseArgs struct {
	Content    EventContentResponsePtrInput `pulumi:"content"`
	Headers    pulumi.StringMapInput        `pulumi:"headers"`
	Method     pulumi.StringPtrInput        `pulumi:"method"`
	RequestUri pulumi.StringPtrInput        `pulumi:"requestUri"`
	Version    pulumi.StringPtrInput        `pulumi:"version"`
}

func (EventRequestMessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventRequestMessageResponse)(nil)).Elem()
}

func (i EventRequestMessageResponseArgs) ToEventRequestMessageResponseOutput() EventRequestMessageResponseOutput {
	return i.ToEventRequestMessageResponseOutputWithContext(context.Background())
}

func (i EventRequestMessageResponseArgs) ToEventRequestMessageResponseOutputWithContext(ctx context.Context) EventRequestMessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventRequestMessageResponseOutput)
}

func (i EventRequestMessageResponseArgs) ToEventRequestMessageResponsePtrOutput() EventRequestMessageResponsePtrOutput {
	return i.ToEventRequestMessageResponsePtrOutputWithContext(context.Background())
}

func (i EventRequestMessageResponseArgs) ToEventRequestMessageResponsePtrOutputWithContext(ctx context.Context) EventRequestMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventRequestMessageResponseOutput).ToEventRequestMessageResponsePtrOutputWithContext(ctx)
}









type EventRequestMessageResponsePtrInput interface {
	pulumi.Input

	ToEventRequestMessageResponsePtrOutput() EventRequestMessageResponsePtrOutput
	ToEventRequestMessageResponsePtrOutputWithContext(context.Context) EventRequestMessageResponsePtrOutput
}

type eventRequestMessageResponsePtrType EventRequestMessageResponseArgs

func EventRequestMessageResponsePtr(v *EventRequestMessageResponseArgs) EventRequestMessageResponsePtrInput {
	return (*eventRequestMessageResponsePtrType)(v)
}

func (*eventRequestMessageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventRequestMessageResponse)(nil)).Elem()
}

func (i *eventRequestMessageResponsePtrType) ToEventRequestMessageResponsePtrOutput() EventRequestMessageResponsePtrOutput {
	return i.ToEventRequestMessageResponsePtrOutputWithContext(context.Background())
}

func (i *eventRequestMessageResponsePtrType) ToEventRequestMessageResponsePtrOutputWithContext(ctx context.Context) EventRequestMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventRequestMessageResponsePtrOutput)
}

type EventRequestMessageResponseOutput struct{ *pulumi.OutputState }

func (EventRequestMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventRequestMessageResponse)(nil)).Elem()
}

func (o EventRequestMessageResponseOutput) ToEventRequestMessageResponseOutput() EventRequestMessageResponseOutput {
	return o
}

func (o EventRequestMessageResponseOutput) ToEventRequestMessageResponseOutputWithContext(ctx context.Context) EventRequestMessageResponseOutput {
	return o
}

func (o EventRequestMessageResponseOutput) ToEventRequestMessageResponsePtrOutput() EventRequestMessageResponsePtrOutput {
	return o.ToEventRequestMessageResponsePtrOutputWithContext(context.Background())
}

func (o EventRequestMessageResponseOutput) ToEventRequestMessageResponsePtrOutputWithContext(ctx context.Context) EventRequestMessageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventRequestMessageResponse) *EventRequestMessageResponse {
		return &v
	}).(EventRequestMessageResponsePtrOutput)
}

func (o EventRequestMessageResponseOutput) Content() EventContentResponsePtrOutput {
	return o.ApplyT(func(v EventRequestMessageResponse) *EventContentResponse { return v.Content }).(EventContentResponsePtrOutput)
}

func (o EventRequestMessageResponseOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v EventRequestMessageResponse) map[string]string { return v.Headers }).(pulumi.StringMapOutput)
}

func (o EventRequestMessageResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventRequestMessageResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o EventRequestMessageResponseOutput) RequestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventRequestMessageResponse) *string { return v.RequestUri }).(pulumi.StringPtrOutput)
}

func (o EventRequestMessageResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventRequestMessageResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type EventRequestMessageResponsePtrOutput struct{ *pulumi.OutputState }

func (EventRequestMessageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventRequestMessageResponse)(nil)).Elem()
}

func (o EventRequestMessageResponsePtrOutput) ToEventRequestMessageResponsePtrOutput() EventRequestMessageResponsePtrOutput {
	return o
}

func (o EventRequestMessageResponsePtrOutput) ToEventRequestMessageResponsePtrOutputWithContext(ctx context.Context) EventRequestMessageResponsePtrOutput {
	return o
}

func (o EventRequestMessageResponsePtrOutput) Elem() EventRequestMessageResponseOutput {
	return o.ApplyT(func(v *EventRequestMessageResponse) EventRequestMessageResponse {
		if v != nil {
			return *v
		}
		var ret EventRequestMessageResponse
		return ret
	}).(EventRequestMessageResponseOutput)
}

func (o EventRequestMessageResponsePtrOutput) Content() EventContentResponsePtrOutput {
	return o.ApplyT(func(v *EventRequestMessageResponse) *EventContentResponse {
		if v == nil {
			return nil
		}
		return v.Content
	}).(EventContentResponsePtrOutput)
}

func (o EventRequestMessageResponsePtrOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventRequestMessageResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringMapOutput)
}

func (o EventRequestMessageResponsePtrOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventRequestMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Method
	}).(pulumi.StringPtrOutput)
}

func (o EventRequestMessageResponsePtrOutput) RequestUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventRequestMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestUri
	}).(pulumi.StringPtrOutput)
}

func (o EventRequestMessageResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventRequestMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type EventResponse struct {
	EventRequestMessage  *EventRequestMessageResponse  `pulumi:"eventRequestMessage"`
	EventResponseMessage *EventResponseMessageResponse `pulumi:"eventResponseMessage"`
	Id                   *string                       `pulumi:"id"`
}





type EventResponseInput interface {
	pulumi.Input

	ToEventResponseOutput() EventResponseOutput
	ToEventResponseOutputWithContext(context.Context) EventResponseOutput
}

type EventResponseArgs struct {
	EventRequestMessage  EventRequestMessageResponsePtrInput  `pulumi:"eventRequestMessage"`
	EventResponseMessage EventResponseMessageResponsePtrInput `pulumi:"eventResponseMessage"`
	Id                   pulumi.StringPtrInput                `pulumi:"id"`
}

func (EventResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponse)(nil)).Elem()
}

func (i EventResponseArgs) ToEventResponseOutput() EventResponseOutput {
	return i.ToEventResponseOutputWithContext(context.Background())
}

func (i EventResponseArgs) ToEventResponseOutputWithContext(ctx context.Context) EventResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseOutput)
}





type EventResponseArrayInput interface {
	pulumi.Input

	ToEventResponseArrayOutput() EventResponseArrayOutput
	ToEventResponseArrayOutputWithContext(context.Context) EventResponseArrayOutput
}

type EventResponseArray []EventResponseInput

func (EventResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventResponse)(nil)).Elem()
}

func (i EventResponseArray) ToEventResponseArrayOutput() EventResponseArrayOutput {
	return i.ToEventResponseArrayOutputWithContext(context.Background())
}

func (i EventResponseArray) ToEventResponseArrayOutputWithContext(ctx context.Context) EventResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseArrayOutput)
}

type EventResponseOutput struct{ *pulumi.OutputState }

func (EventResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponse)(nil)).Elem()
}

func (o EventResponseOutput) ToEventResponseOutput() EventResponseOutput {
	return o
}

func (o EventResponseOutput) ToEventResponseOutputWithContext(ctx context.Context) EventResponseOutput {
	return o
}

func (o EventResponseOutput) EventRequestMessage() EventRequestMessageResponsePtrOutput {
	return o.ApplyT(func(v EventResponse) *EventRequestMessageResponse { return v.EventRequestMessage }).(EventRequestMessageResponsePtrOutput)
}

func (o EventResponseOutput) EventResponseMessage() EventResponseMessageResponsePtrOutput {
	return o.ApplyT(func(v EventResponse) *EventResponseMessageResponse { return v.EventResponseMessage }).(EventResponseMessageResponsePtrOutput)
}

func (o EventResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type EventResponseArrayOutput struct{ *pulumi.OutputState }

func (EventResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EventResponse)(nil)).Elem()
}

func (o EventResponseArrayOutput) ToEventResponseArrayOutput() EventResponseArrayOutput {
	return o
}

func (o EventResponseArrayOutput) ToEventResponseArrayOutputWithContext(ctx context.Context) EventResponseArrayOutput {
	return o
}

func (o EventResponseArrayOutput) Index(i pulumi.IntInput) EventResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EventResponse {
		return vs[0].([]EventResponse)[vs[1].(int)]
	}).(EventResponseOutput)
}

type EventResponseMessageResponse struct {
	Content      *string           `pulumi:"content"`
	Headers      map[string]string `pulumi:"headers"`
	ReasonPhrase *string           `pulumi:"reasonPhrase"`
	StatusCode   *string           `pulumi:"statusCode"`
	Version      *string           `pulumi:"version"`
}





type EventResponseMessageResponseInput interface {
	pulumi.Input

	ToEventResponseMessageResponseOutput() EventResponseMessageResponseOutput
	ToEventResponseMessageResponseOutputWithContext(context.Context) EventResponseMessageResponseOutput
}

type EventResponseMessageResponseArgs struct {
	Content      pulumi.StringPtrInput `pulumi:"content"`
	Headers      pulumi.StringMapInput `pulumi:"headers"`
	ReasonPhrase pulumi.StringPtrInput `pulumi:"reasonPhrase"`
	StatusCode   pulumi.StringPtrInput `pulumi:"statusCode"`
	Version      pulumi.StringPtrInput `pulumi:"version"`
}

func (EventResponseMessageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponseMessageResponse)(nil)).Elem()
}

func (i EventResponseMessageResponseArgs) ToEventResponseMessageResponseOutput() EventResponseMessageResponseOutput {
	return i.ToEventResponseMessageResponseOutputWithContext(context.Background())
}

func (i EventResponseMessageResponseArgs) ToEventResponseMessageResponseOutputWithContext(ctx context.Context) EventResponseMessageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseMessageResponseOutput)
}

func (i EventResponseMessageResponseArgs) ToEventResponseMessageResponsePtrOutput() EventResponseMessageResponsePtrOutput {
	return i.ToEventResponseMessageResponsePtrOutputWithContext(context.Background())
}

func (i EventResponseMessageResponseArgs) ToEventResponseMessageResponsePtrOutputWithContext(ctx context.Context) EventResponseMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseMessageResponseOutput).ToEventResponseMessageResponsePtrOutputWithContext(ctx)
}









type EventResponseMessageResponsePtrInput interface {
	pulumi.Input

	ToEventResponseMessageResponsePtrOutput() EventResponseMessageResponsePtrOutput
	ToEventResponseMessageResponsePtrOutputWithContext(context.Context) EventResponseMessageResponsePtrOutput
}

type eventResponseMessageResponsePtrType EventResponseMessageResponseArgs

func EventResponseMessageResponsePtr(v *EventResponseMessageResponseArgs) EventResponseMessageResponsePtrInput {
	return (*eventResponseMessageResponsePtrType)(v)
}

func (*eventResponseMessageResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EventResponseMessageResponse)(nil)).Elem()
}

func (i *eventResponseMessageResponsePtrType) ToEventResponseMessageResponsePtrOutput() EventResponseMessageResponsePtrOutput {
	return i.ToEventResponseMessageResponsePtrOutputWithContext(context.Background())
}

func (i *eventResponseMessageResponsePtrType) ToEventResponseMessageResponsePtrOutputWithContext(ctx context.Context) EventResponseMessageResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EventResponseMessageResponsePtrOutput)
}

type EventResponseMessageResponseOutput struct{ *pulumi.OutputState }

func (EventResponseMessageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventResponseMessageResponse)(nil)).Elem()
}

func (o EventResponseMessageResponseOutput) ToEventResponseMessageResponseOutput() EventResponseMessageResponseOutput {
	return o
}

func (o EventResponseMessageResponseOutput) ToEventResponseMessageResponseOutputWithContext(ctx context.Context) EventResponseMessageResponseOutput {
	return o
}

func (o EventResponseMessageResponseOutput) ToEventResponseMessageResponsePtrOutput() EventResponseMessageResponsePtrOutput {
	return o.ToEventResponseMessageResponsePtrOutputWithContext(context.Background())
}

func (o EventResponseMessageResponseOutput) ToEventResponseMessageResponsePtrOutputWithContext(ctx context.Context) EventResponseMessageResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventResponseMessageResponse) *EventResponseMessageResponse {
		return &v
	}).(EventResponseMessageResponsePtrOutput)
}

func (o EventResponseMessageResponseOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponseMessageResponse) *string { return v.Content }).(pulumi.StringPtrOutput)
}

func (o EventResponseMessageResponseOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v EventResponseMessageResponse) map[string]string { return v.Headers }).(pulumi.StringMapOutput)
}

func (o EventResponseMessageResponseOutput) ReasonPhrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponseMessageResponse) *string { return v.ReasonPhrase }).(pulumi.StringPtrOutput)
}

func (o EventResponseMessageResponseOutput) StatusCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponseMessageResponse) *string { return v.StatusCode }).(pulumi.StringPtrOutput)
}

func (o EventResponseMessageResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EventResponseMessageResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type EventResponseMessageResponsePtrOutput struct{ *pulumi.OutputState }

func (EventResponseMessageResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventResponseMessageResponse)(nil)).Elem()
}

func (o EventResponseMessageResponsePtrOutput) ToEventResponseMessageResponsePtrOutput() EventResponseMessageResponsePtrOutput {
	return o
}

func (o EventResponseMessageResponsePtrOutput) ToEventResponseMessageResponsePtrOutputWithContext(ctx context.Context) EventResponseMessageResponsePtrOutput {
	return o
}

func (o EventResponseMessageResponsePtrOutput) Elem() EventResponseMessageResponseOutput {
	return o.ApplyT(func(v *EventResponseMessageResponse) EventResponseMessageResponse {
		if v != nil {
			return *v
		}
		var ret EventResponseMessageResponse
		return ret
	}).(EventResponseMessageResponseOutput)
}

func (o EventResponseMessageResponsePtrOutput) Content() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventResponseMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Content
	}).(pulumi.StringPtrOutput)
}

func (o EventResponseMessageResponsePtrOutput) Headers() pulumi.StringMapOutput {
	return o.ApplyT(func(v *EventResponseMessageResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Headers
	}).(pulumi.StringMapOutput)
}

func (o EventResponseMessageResponsePtrOutput) ReasonPhrase() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventResponseMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.ReasonPhrase
	}).(pulumi.StringPtrOutput)
}

func (o EventResponseMessageResponsePtrOutput) StatusCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventResponseMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.StatusCode
	}).(pulumi.StringPtrOutput)
}

func (o EventResponseMessageResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EventResponseMessageResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type IPRule struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRule) Defaults() *IPRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleInput interface {
	pulumi.Input

	ToIPRuleOutput() IPRuleOutput
	ToIPRuleOutputWithContext(context.Context) IPRuleOutput
}

type IPRuleArgs struct {
	Action           pulumi.StringPtrInput `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput    `pulumi:"iPAddressOrRange"`
}

func (IPRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (i IPRuleArgs) ToIPRuleOutput() IPRuleOutput {
	return i.ToIPRuleOutputWithContext(context.Background())
}

func (i IPRuleArgs) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleOutput)
}





type IPRuleArrayInput interface {
	pulumi.Input

	ToIPRuleArrayOutput() IPRuleArrayOutput
	ToIPRuleArrayOutputWithContext(context.Context) IPRuleArrayOutput
}

type IPRuleArray []IPRuleInput

func (IPRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (i IPRuleArray) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return i.ToIPRuleArrayOutputWithContext(context.Background())
}

func (i IPRuleArray) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleArrayOutput)
}

type IPRuleOutput struct{ *pulumi.OutputState }

func (IPRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRule)(nil)).Elem()
}

func (o IPRuleOutput) ToIPRuleOutput() IPRuleOutput {
	return o
}

func (o IPRuleOutput) ToIPRuleOutputWithContext(ctx context.Context) IPRuleOutput {
	return o
}

func (o IPRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRule) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleArrayOutput struct{ *pulumi.OutputState }

func (IPRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRule)(nil)).Elem()
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutput() IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) ToIPRuleArrayOutputWithContext(ctx context.Context) IPRuleArrayOutput {
	return o
}

func (o IPRuleArrayOutput) Index(i pulumi.IntInput) IPRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRule {
		return vs[0].([]IPRule)[vs[1].(int)]
	}).(IPRuleOutput)
}

type IPRuleResponse struct {
	Action           *string `pulumi:"action"`
	IPAddressOrRange string  `pulumi:"iPAddressOrRange"`
}


func (val *IPRuleResponse) Defaults() *IPRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type IPRuleResponseInput interface {
	pulumi.Input

	ToIPRuleResponseOutput() IPRuleResponseOutput
	ToIPRuleResponseOutputWithContext(context.Context) IPRuleResponseOutput
}

type IPRuleResponseArgs struct {
	Action           pulumi.StringPtrInput `pulumi:"action"`
	IPAddressOrRange pulumi.StringInput    `pulumi:"iPAddressOrRange"`
}

func (IPRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return i.ToIPRuleResponseOutputWithContext(context.Background())
}

func (i IPRuleResponseArgs) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseOutput)
}





type IPRuleResponseArrayInput interface {
	pulumi.Input

	ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput
	ToIPRuleResponseArrayOutputWithContext(context.Context) IPRuleResponseArrayOutput
}

type IPRuleResponseArray []IPRuleResponseInput

func (IPRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return i.ToIPRuleResponseArrayOutputWithContext(context.Background())
}

func (i IPRuleResponseArray) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IPRuleResponseArrayOutput)
}

type IPRuleResponseOutput struct{ *pulumi.OutputState }

func (IPRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutput() IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) ToIPRuleResponseOutputWithContext(ctx context.Context) IPRuleResponseOutput {
	return o
}

func (o IPRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IPRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IPRuleResponseOutput) IPAddressOrRange() pulumi.StringOutput {
	return o.ApplyT(func(v IPRuleResponse) string { return v.IPAddressOrRange }).(pulumi.StringOutput)
}

type IPRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (IPRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IPRuleResponse)(nil)).Elem()
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutput() IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) ToIPRuleResponseArrayOutputWithContext(ctx context.Context) IPRuleResponseArrayOutput {
	return o
}

func (o IPRuleResponseArrayOutput) Index(i pulumi.IntInput) IPRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IPRuleResponse {
		return vs[0].([]IPRuleResponse)[vs[1].(int)]
	}).(IPRuleResponseOutput)
}

type NetworkRuleSet struct {
	DefaultAction       string               `pulumi:"defaultAction"`
	IpRules             []IPRule             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRule `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSet) Defaults() *NetworkRuleSet {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}





type NetworkRuleSetInput interface {
	pulumi.Input

	ToNetworkRuleSetOutput() NetworkRuleSetOutput
	ToNetworkRuleSetOutputWithContext(context.Context) NetworkRuleSetOutput
}

type NetworkRuleSetArgs struct {
	DefaultAction       pulumi.StringInput           `pulumi:"defaultAction"`
	IpRules             IPRuleArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return i.ToNetworkRuleSetOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput)
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetArgs) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetOutput).ToNetworkRuleSetPtrOutputWithContext(ctx)
}









type NetworkRuleSetPtrInput interface {
	pulumi.Input

	ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput
	ToNetworkRuleSetPtrOutputWithContext(context.Context) NetworkRuleSetPtrOutput
}

type networkRuleSetPtrType NetworkRuleSetArgs

func NetworkRuleSetPtr(v *NetworkRuleSetArgs) NetworkRuleSetPtrInput {
	return (*networkRuleSetPtrType)(v)
}

func (*networkRuleSetPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return i.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (i *networkRuleSetPtrType) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetPtrOutput)
}

type NetworkRuleSetOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutput() NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetOutputWithContext(ctx context.Context) NetworkRuleSetOutput {
	return o
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o.ToNetworkRuleSetPtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSet) *NetworkRuleSet {
		return &v
	}).(NetworkRuleSetPtrOutput)
}

func (o NetworkRuleSetOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSet) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []IPRule { return v.IpRules }).(IPRuleArrayOutput)
}

func (o NetworkRuleSetOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v NetworkRuleSet) []VirtualNetworkRule { return v.VirtualNetworkRules }).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetPtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSet)(nil)).Elem()
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutput() NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) ToNetworkRuleSetPtrOutputWithContext(ctx context.Context) NetworkRuleSetPtrOutput {
	return o
}

func (o NetworkRuleSetPtrOutput) Elem() NetworkRuleSetOutput {
	return o.ApplyT(func(v *NetworkRuleSet) NetworkRuleSet {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSet
		return ret
	}).(NetworkRuleSetOutput)
}

func (o NetworkRuleSetPtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSet) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetPtrOutput) IpRules() IPRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []IPRule {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleArrayOutput)
}

func (o NetworkRuleSetPtrOutput) VirtualNetworkRules() VirtualNetworkRuleArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSet) []VirtualNetworkRule {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleArrayOutput)
}

type NetworkRuleSetResponse struct {
	DefaultAction       string                       `pulumi:"defaultAction"`
	IpRules             []IPRuleResponse             `pulumi:"ipRules"`
	VirtualNetworkRules []VirtualNetworkRuleResponse `pulumi:"virtualNetworkRules"`
}


func (val *NetworkRuleSetResponse) Defaults() *NetworkRuleSetResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DefaultAction) {
		tmp.DefaultAction = "Allow"
	}
	return &tmp
}





type NetworkRuleSetResponseInput interface {
	pulumi.Input

	ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput
	ToNetworkRuleSetResponseOutputWithContext(context.Context) NetworkRuleSetResponseOutput
}

type NetworkRuleSetResponseArgs struct {
	DefaultAction       pulumi.StringInput                   `pulumi:"defaultAction"`
	IpRules             IPRuleResponseArrayInput             `pulumi:"ipRules"`
	VirtualNetworkRules VirtualNetworkRuleResponseArrayInput `pulumi:"virtualNetworkRules"`
}

func (NetworkRuleSetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return i.ToNetworkRuleSetResponseOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput)
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i NetworkRuleSetResponseArgs) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponseOutput).ToNetworkRuleSetResponsePtrOutputWithContext(ctx)
}









type NetworkRuleSetResponsePtrInput interface {
	pulumi.Input

	ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput
	ToNetworkRuleSetResponsePtrOutputWithContext(context.Context) NetworkRuleSetResponsePtrOutput
}

type networkRuleSetResponsePtrType NetworkRuleSetResponseArgs

func NetworkRuleSetResponsePtr(v *NetworkRuleSetResponseArgs) NetworkRuleSetResponsePtrInput {
	return (*networkRuleSetResponsePtrType)(v)
}

func (*networkRuleSetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return i.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (i *networkRuleSetResponsePtrType) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRuleSetResponsePtrOutput)
}

type NetworkRuleSetResponseOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutput() NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponseOutputWithContext(ctx context.Context) NetworkRuleSetResponseOutput {
	return o
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o.ToNetworkRuleSetResponsePtrOutputWithContext(context.Background())
}

func (o NetworkRuleSetResponseOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NetworkRuleSetResponse) *NetworkRuleSetResponse {
		return &v
	}).(NetworkRuleSetResponsePtrOutput)
}

func (o NetworkRuleSetResponseOutput) DefaultAction() pulumi.StringOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) string { return v.DefaultAction }).(pulumi.StringOutput)
}

func (o NetworkRuleSetResponseOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []IPRuleResponse { return v.IpRules }).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponseOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v NetworkRuleSetResponse) []VirtualNetworkRuleResponse { return v.VirtualNetworkRules }).(VirtualNetworkRuleResponseArrayOutput)
}

type NetworkRuleSetResponsePtrOutput struct{ *pulumi.OutputState }

func (NetworkRuleSetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NetworkRuleSetResponse)(nil)).Elem()
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutput() NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) ToNetworkRuleSetResponsePtrOutputWithContext(ctx context.Context) NetworkRuleSetResponsePtrOutput {
	return o
}

func (o NetworkRuleSetResponsePtrOutput) Elem() NetworkRuleSetResponseOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) NetworkRuleSetResponse {
		if v != nil {
			return *v
		}
		var ret NetworkRuleSetResponse
		return ret
	}).(NetworkRuleSetResponseOutput)
}

func (o NetworkRuleSetResponsePtrOutput) DefaultAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultAction
	}).(pulumi.StringPtrOutput)
}

func (o NetworkRuleSetResponsePtrOutput) IpRules() IPRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []IPRuleResponse {
		if v == nil {
			return nil
		}
		return v.IpRules
	}).(IPRuleResponseArrayOutput)
}

func (o NetworkRuleSetResponsePtrOutput) VirtualNetworkRules() VirtualNetworkRuleResponseArrayOutput {
	return o.ApplyT(func(v *NetworkRuleSetResponse) []VirtualNetworkRuleResponse {
		if v == nil {
			return nil
		}
		return v.VirtualNetworkRules
	}).(VirtualNetworkRuleResponseArrayOutput)
}

type RegistryPasswordResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type RegistryPasswordResponseInput interface {
	pulumi.Input

	ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput
	ToRegistryPasswordResponseOutputWithContext(context.Context) RegistryPasswordResponseOutput
}

type RegistryPasswordResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (RegistryPasswordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPasswordResponse)(nil)).Elem()
}

func (i RegistryPasswordResponseArgs) ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput {
	return i.ToRegistryPasswordResponseOutputWithContext(context.Background())
}

func (i RegistryPasswordResponseArgs) ToRegistryPasswordResponseOutputWithContext(ctx context.Context) RegistryPasswordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPasswordResponseOutput)
}





type RegistryPasswordResponseArrayInput interface {
	pulumi.Input

	ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput
	ToRegistryPasswordResponseArrayOutputWithContext(context.Context) RegistryPasswordResponseArrayOutput
}

type RegistryPasswordResponseArray []RegistryPasswordResponseInput

func (RegistryPasswordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryPasswordResponse)(nil)).Elem()
}

func (i RegistryPasswordResponseArray) ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput {
	return i.ToRegistryPasswordResponseArrayOutputWithContext(context.Background())
}

func (i RegistryPasswordResponseArray) ToRegistryPasswordResponseArrayOutputWithContext(ctx context.Context) RegistryPasswordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryPasswordResponseArrayOutput)
}

type RegistryPasswordResponseOutput struct{ *pulumi.OutputState }

func (RegistryPasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryPasswordResponse)(nil)).Elem()
}

func (o RegistryPasswordResponseOutput) ToRegistryPasswordResponseOutput() RegistryPasswordResponseOutput {
	return o
}

func (o RegistryPasswordResponseOutput) ToRegistryPasswordResponseOutputWithContext(ctx context.Context) RegistryPasswordResponseOutput {
	return o
}

func (o RegistryPasswordResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryPasswordResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RegistryPasswordResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryPasswordResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type RegistryPasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (RegistryPasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryPasswordResponse)(nil)).Elem()
}

func (o RegistryPasswordResponseArrayOutput) ToRegistryPasswordResponseArrayOutput() RegistryPasswordResponseArrayOutput {
	return o
}

func (o RegistryPasswordResponseArrayOutput) ToRegistryPasswordResponseArrayOutputWithContext(ctx context.Context) RegistryPasswordResponseArrayOutput {
	return o
}

func (o RegistryPasswordResponseArrayOutput) Index(i pulumi.IntInput) RegistryPasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryPasswordResponse {
		return vs[0].([]RegistryPasswordResponse)[vs[1].(int)]
	}).(RegistryPasswordResponseOutput)
}

type RequestResponse struct {
	Addr      *string `pulumi:"addr"`
	Host      *string `pulumi:"host"`
	Id        *string `pulumi:"id"`
	Method    *string `pulumi:"method"`
	Useragent *string `pulumi:"useragent"`
}





type RequestResponseInput interface {
	pulumi.Input

	ToRequestResponseOutput() RequestResponseOutput
	ToRequestResponseOutputWithContext(context.Context) RequestResponseOutput
}

type RequestResponseArgs struct {
	Addr      pulumi.StringPtrInput `pulumi:"addr"`
	Host      pulumi.StringPtrInput `pulumi:"host"`
	Id        pulumi.StringPtrInput `pulumi:"id"`
	Method    pulumi.StringPtrInput `pulumi:"method"`
	Useragent pulumi.StringPtrInput `pulumi:"useragent"`
}

func (RequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestResponse)(nil)).Elem()
}

func (i RequestResponseArgs) ToRequestResponseOutput() RequestResponseOutput {
	return i.ToRequestResponseOutputWithContext(context.Background())
}

func (i RequestResponseArgs) ToRequestResponseOutputWithContext(ctx context.Context) RequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestResponseOutput)
}

func (i RequestResponseArgs) ToRequestResponsePtrOutput() RequestResponsePtrOutput {
	return i.ToRequestResponsePtrOutputWithContext(context.Background())
}

func (i RequestResponseArgs) ToRequestResponsePtrOutputWithContext(ctx context.Context) RequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestResponseOutput).ToRequestResponsePtrOutputWithContext(ctx)
}









type RequestResponsePtrInput interface {
	pulumi.Input

	ToRequestResponsePtrOutput() RequestResponsePtrOutput
	ToRequestResponsePtrOutputWithContext(context.Context) RequestResponsePtrOutput
}

type requestResponsePtrType RequestResponseArgs

func RequestResponsePtr(v *RequestResponseArgs) RequestResponsePtrInput {
	return (*requestResponsePtrType)(v)
}

func (*requestResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestResponse)(nil)).Elem()
}

func (i *requestResponsePtrType) ToRequestResponsePtrOutput() RequestResponsePtrOutput {
	return i.ToRequestResponsePtrOutputWithContext(context.Background())
}

func (i *requestResponsePtrType) ToRequestResponsePtrOutputWithContext(ctx context.Context) RequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestResponsePtrOutput)
}

type RequestResponseOutput struct{ *pulumi.OutputState }

func (RequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestResponse)(nil)).Elem()
}

func (o RequestResponseOutput) ToRequestResponseOutput() RequestResponseOutput {
	return o
}

func (o RequestResponseOutput) ToRequestResponseOutputWithContext(ctx context.Context) RequestResponseOutput {
	return o
}

func (o RequestResponseOutput) ToRequestResponsePtrOutput() RequestResponsePtrOutput {
	return o.ToRequestResponsePtrOutputWithContext(context.Background())
}

func (o RequestResponseOutput) ToRequestResponsePtrOutputWithContext(ctx context.Context) RequestResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestResponse) *RequestResponse {
		return &v
	}).(RequestResponsePtrOutput)
}

func (o RequestResponseOutput) Addr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestResponse) *string { return v.Addr }).(pulumi.StringPtrOutput)
}

func (o RequestResponseOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestResponse) *string { return v.Host }).(pulumi.StringPtrOutput)
}

func (o RequestResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o RequestResponseOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestResponse) *string { return v.Method }).(pulumi.StringPtrOutput)
}

func (o RequestResponseOutput) Useragent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestResponse) *string { return v.Useragent }).(pulumi.StringPtrOutput)
}

type RequestResponsePtrOutput struct{ *pulumi.OutputState }

func (RequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestResponse)(nil)).Elem()
}

func (o RequestResponsePtrOutput) ToRequestResponsePtrOutput() RequestResponsePtrOutput {
	return o
}

func (o RequestResponsePtrOutput) ToRequestResponsePtrOutputWithContext(ctx context.Context) RequestResponsePtrOutput {
	return o
}

func (o RequestResponsePtrOutput) Elem() RequestResponseOutput {
	return o.ApplyT(func(v *RequestResponse) RequestResponse {
		if v != nil {
			return *v
		}
		var ret RequestResponse
		return ret
	}).(RequestResponseOutput)
}

func (o RequestResponsePtrOutput) Addr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Addr
	}).(pulumi.StringPtrOutput)
}

func (o RequestResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Host
	}).(pulumi.StringPtrOutput)
}

func (o RequestResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RequestResponsePtrOutput) Method() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Method
	}).(pulumi.StringPtrOutput)
}

func (o RequestResponsePtrOutput) Useragent() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.Useragent
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SourceResponse struct {
	Addr       *string `pulumi:"addr"`
	InstanceID *string `pulumi:"instanceID"`
}





type SourceResponseInput interface {
	pulumi.Input

	ToSourceResponseOutput() SourceResponseOutput
	ToSourceResponseOutputWithContext(context.Context) SourceResponseOutput
}

type SourceResponseArgs struct {
	Addr       pulumi.StringPtrInput `pulumi:"addr"`
	InstanceID pulumi.StringPtrInput `pulumi:"instanceID"`
}

func (SourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceResponse)(nil)).Elem()
}

func (i SourceResponseArgs) ToSourceResponseOutput() SourceResponseOutput {
	return i.ToSourceResponseOutputWithContext(context.Background())
}

func (i SourceResponseArgs) ToSourceResponseOutputWithContext(ctx context.Context) SourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceResponseOutput)
}

func (i SourceResponseArgs) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return i.ToSourceResponsePtrOutputWithContext(context.Background())
}

func (i SourceResponseArgs) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceResponseOutput).ToSourceResponsePtrOutputWithContext(ctx)
}









type SourceResponsePtrInput interface {
	pulumi.Input

	ToSourceResponsePtrOutput() SourceResponsePtrOutput
	ToSourceResponsePtrOutputWithContext(context.Context) SourceResponsePtrOutput
}

type sourceResponsePtrType SourceResponseArgs

func SourceResponsePtr(v *SourceResponseArgs) SourceResponsePtrInput {
	return (*sourceResponsePtrType)(v)
}

func (*sourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceResponse)(nil)).Elem()
}

func (i *sourceResponsePtrType) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return i.ToSourceResponsePtrOutputWithContext(context.Background())
}

func (i *sourceResponsePtrType) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceResponsePtrOutput)
}

type SourceResponseOutput struct{ *pulumi.OutputState }

func (SourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceResponse)(nil)).Elem()
}

func (o SourceResponseOutput) ToSourceResponseOutput() SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) ToSourceResponseOutputWithContext(ctx context.Context) SourceResponseOutput {
	return o
}

func (o SourceResponseOutput) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return o.ToSourceResponsePtrOutputWithContext(context.Background())
}

func (o SourceResponseOutput) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceResponse) *SourceResponse {
		return &v
	}).(SourceResponsePtrOutput)
}

func (o SourceResponseOutput) Addr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.Addr }).(pulumi.StringPtrOutput)
}

func (o SourceResponseOutput) InstanceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceResponse) *string { return v.InstanceID }).(pulumi.StringPtrOutput)
}

type SourceResponsePtrOutput struct{ *pulumi.OutputState }

func (SourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceResponse)(nil)).Elem()
}

func (o SourceResponsePtrOutput) ToSourceResponsePtrOutput() SourceResponsePtrOutput {
	return o
}

func (o SourceResponsePtrOutput) ToSourceResponsePtrOutputWithContext(ctx context.Context) SourceResponsePtrOutput {
	return o
}

func (o SourceResponsePtrOutput) Elem() SourceResponseOutput {
	return o.ApplyT(func(v *SourceResponse) SourceResponse {
		if v != nil {
			return *v
		}
		var ret SourceResponse
		return ret
	}).(SourceResponseOutput)
}

func (o SourceResponsePtrOutput) Addr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Addr
	}).(pulumi.StringPtrOutput)
}

func (o SourceResponsePtrOutput) InstanceID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.InstanceID
	}).(pulumi.StringPtrOutput)
}

type StatusResponse struct {
	DisplayStatus string `pulumi:"displayStatus"`
	Message       string `pulumi:"message"`
	Timestamp     string `pulumi:"timestamp"`
}





type StatusResponseInput interface {
	pulumi.Input

	ToStatusResponseOutput() StatusResponseOutput
	ToStatusResponseOutputWithContext(context.Context) StatusResponseOutput
}

type StatusResponseArgs struct {
	DisplayStatus pulumi.StringInput `pulumi:"displayStatus"`
	Message       pulumi.StringInput `pulumi:"message"`
	Timestamp     pulumi.StringInput `pulumi:"timestamp"`
}

func (StatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusResponse)(nil)).Elem()
}

func (i StatusResponseArgs) ToStatusResponseOutput() StatusResponseOutput {
	return i.ToStatusResponseOutputWithContext(context.Background())
}

func (i StatusResponseArgs) ToStatusResponseOutputWithContext(ctx context.Context) StatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusResponseOutput)
}

func (i StatusResponseArgs) ToStatusResponsePtrOutput() StatusResponsePtrOutput {
	return i.ToStatusResponsePtrOutputWithContext(context.Background())
}

func (i StatusResponseArgs) ToStatusResponsePtrOutputWithContext(ctx context.Context) StatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusResponseOutput).ToStatusResponsePtrOutputWithContext(ctx)
}









type StatusResponsePtrInput interface {
	pulumi.Input

	ToStatusResponsePtrOutput() StatusResponsePtrOutput
	ToStatusResponsePtrOutputWithContext(context.Context) StatusResponsePtrOutput
}

type statusResponsePtrType StatusResponseArgs

func StatusResponsePtr(v *StatusResponseArgs) StatusResponsePtrInput {
	return (*statusResponsePtrType)(v)
}

func (*statusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusResponse)(nil)).Elem()
}

func (i *statusResponsePtrType) ToStatusResponsePtrOutput() StatusResponsePtrOutput {
	return i.ToStatusResponsePtrOutputWithContext(context.Background())
}

func (i *statusResponsePtrType) ToStatusResponsePtrOutputWithContext(ctx context.Context) StatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusResponsePtrOutput)
}

type StatusResponseOutput struct{ *pulumi.OutputState }

func (StatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusResponse)(nil)).Elem()
}

func (o StatusResponseOutput) ToStatusResponseOutput() StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) ToStatusResponseOutputWithContext(ctx context.Context) StatusResponseOutput {
	return o
}

func (o StatusResponseOutput) ToStatusResponsePtrOutput() StatusResponsePtrOutput {
	return o.ToStatusResponsePtrOutputWithContext(context.Background())
}

func (o StatusResponseOutput) ToStatusResponsePtrOutputWithContext(ctx context.Context) StatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusResponse) *StatusResponse {
		return &v
	}).(StatusResponsePtrOutput)
}

func (o StatusResponseOutput) DisplayStatus() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.DisplayStatus }).(pulumi.StringOutput)
}

func (o StatusResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o StatusResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v StatusResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

type StatusResponsePtrOutput struct{ *pulumi.OutputState }

func (StatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusResponse)(nil)).Elem()
}

func (o StatusResponsePtrOutput) ToStatusResponsePtrOutput() StatusResponsePtrOutput {
	return o
}

func (o StatusResponsePtrOutput) ToStatusResponsePtrOutputWithContext(ctx context.Context) StatusResponsePtrOutput {
	return o
}

func (o StatusResponsePtrOutput) Elem() StatusResponseOutput {
	return o.ApplyT(func(v *StatusResponse) StatusResponse {
		if v != nil {
			return *v
		}
		var ret StatusResponse
		return ret
	}).(StatusResponseOutput)
}

func (o StatusResponsePtrOutput) DisplayStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayStatus
	}).(pulumi.StringPtrOutput)
}

func (o StatusResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o StatusResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Timestamp
	}).(pulumi.StringPtrOutput)
}

type StorageAccountProperties struct {
	Id string `pulumi:"id"`
}





type StorageAccountPropertiesInput interface {
	pulumi.Input

	ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput
	ToStorageAccountPropertiesOutputWithContext(context.Context) StorageAccountPropertiesOutput
}

type StorageAccountPropertiesArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (StorageAccountPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return i.ToStorageAccountPropertiesOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput)
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesArgs) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesOutput).ToStorageAccountPropertiesPtrOutputWithContext(ctx)
}









type StorageAccountPropertiesPtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput
	ToStorageAccountPropertiesPtrOutputWithContext(context.Context) StorageAccountPropertiesPtrOutput
}

type storageAccountPropertiesPtrType StorageAccountPropertiesArgs

func StorageAccountPropertiesPtr(v *StorageAccountPropertiesArgs) StorageAccountPropertiesPtrInput {
	return (*storageAccountPropertiesPtrType)(v)
}

func (*storageAccountPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return i.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesPtrType) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesPtrOutput)
}

type StorageAccountPropertiesOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutput() StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesOutputWithContext(ctx context.Context) StorageAccountPropertiesOutput {
	return o
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o.ToStorageAccountPropertiesPtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountProperties) *StorageAccountProperties {
		return &v
	}).(StorageAccountPropertiesPtrOutput)
}

func (o StorageAccountPropertiesOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountProperties) string { return v.Id }).(pulumi.StringOutput)
}

type StorageAccountPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountProperties)(nil)).Elem()
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutput() StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) ToStorageAccountPropertiesPtrOutputWithContext(ctx context.Context) StorageAccountPropertiesPtrOutput {
	return o
}

func (o StorageAccountPropertiesPtrOutput) Elem() StorageAccountPropertiesOutput {
	return o.ApplyT(func(v *StorageAccountProperties) StorageAccountProperties {
		if v != nil {
			return *v
		}
		var ret StorageAccountProperties
		return ret
	}).(StorageAccountPropertiesOutput)
}

func (o StorageAccountPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type StorageAccountPropertiesResponse struct {
	Id string `pulumi:"id"`
}





type StorageAccountPropertiesResponseInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput
	ToStorageAccountPropertiesResponseOutputWithContext(context.Context) StorageAccountPropertiesResponseOutput
}

type StorageAccountPropertiesResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (StorageAccountPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return i.ToStorageAccountPropertiesResponseOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponseOutput)
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return i.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i StorageAccountPropertiesResponseArgs) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponseOutput).ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx)
}









type StorageAccountPropertiesResponsePtrInput interface {
	pulumi.Input

	ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput
	ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Context) StorageAccountPropertiesResponsePtrOutput
}

type storageAccountPropertiesResponsePtrType StorageAccountPropertiesResponseArgs

func StorageAccountPropertiesResponsePtr(v *StorageAccountPropertiesResponseArgs) StorageAccountPropertiesResponsePtrInput {
	return (*storageAccountPropertiesResponsePtrType)(v)
}

func (*storageAccountPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (i *storageAccountPropertiesResponsePtrType) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return i.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *storageAccountPropertiesResponsePtrType) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageAccountPropertiesResponsePtrOutput)
}

type StorageAccountPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutput() StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponseOutputWithContext(ctx context.Context) StorageAccountPropertiesResponseOutput {
	return o
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o.ToStorageAccountPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o StorageAccountPropertiesResponseOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageAccountPropertiesResponse) *StorageAccountPropertiesResponse {
		return &v
	}).(StorageAccountPropertiesResponsePtrOutput)
}

func (o StorageAccountPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StorageAccountPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

type StorageAccountPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageAccountPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageAccountPropertiesResponse)(nil)).Elem()
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutput() StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) ToStorageAccountPropertiesResponsePtrOutputWithContext(ctx context.Context) StorageAccountPropertiesResponsePtrOutput {
	return o
}

func (o StorageAccountPropertiesResponsePtrOutput) Elem() StorageAccountPropertiesResponseOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) StorageAccountPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StorageAccountPropertiesResponse
		return ret
	}).(StorageAccountPropertiesResponseOutput)
}

func (o StorageAccountPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageAccountPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type TargetResponse struct {
	Digest     *string  `pulumi:"digest"`
	Length     *float64 `pulumi:"length"`
	MediaType  *string  `pulumi:"mediaType"`
	Name       *string  `pulumi:"name"`
	Repository *string  `pulumi:"repository"`
	Size       *float64 `pulumi:"size"`
	Tag        *string  `pulumi:"tag"`
	Url        *string  `pulumi:"url"`
	Version    *string  `pulumi:"version"`
}





type TargetResponseInput interface {
	pulumi.Input

	ToTargetResponseOutput() TargetResponseOutput
	ToTargetResponseOutputWithContext(context.Context) TargetResponseOutput
}

type TargetResponseArgs struct {
	Digest     pulumi.StringPtrInput  `pulumi:"digest"`
	Length     pulumi.Float64PtrInput `pulumi:"length"`
	MediaType  pulumi.StringPtrInput  `pulumi:"mediaType"`
	Name       pulumi.StringPtrInput  `pulumi:"name"`
	Repository pulumi.StringPtrInput  `pulumi:"repository"`
	Size       pulumi.Float64PtrInput `pulumi:"size"`
	Tag        pulumi.StringPtrInput  `pulumi:"tag"`
	Url        pulumi.StringPtrInput  `pulumi:"url"`
	Version    pulumi.StringPtrInput  `pulumi:"version"`
}

func (TargetResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetResponse)(nil)).Elem()
}

func (i TargetResponseArgs) ToTargetResponseOutput() TargetResponseOutput {
	return i.ToTargetResponseOutputWithContext(context.Background())
}

func (i TargetResponseArgs) ToTargetResponseOutputWithContext(ctx context.Context) TargetResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetResponseOutput)
}

func (i TargetResponseArgs) ToTargetResponsePtrOutput() TargetResponsePtrOutput {
	return i.ToTargetResponsePtrOutputWithContext(context.Background())
}

func (i TargetResponseArgs) ToTargetResponsePtrOutputWithContext(ctx context.Context) TargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetResponseOutput).ToTargetResponsePtrOutputWithContext(ctx)
}









type TargetResponsePtrInput interface {
	pulumi.Input

	ToTargetResponsePtrOutput() TargetResponsePtrOutput
	ToTargetResponsePtrOutputWithContext(context.Context) TargetResponsePtrOutput
}

type targetResponsePtrType TargetResponseArgs

func TargetResponsePtr(v *TargetResponseArgs) TargetResponsePtrInput {
	return (*targetResponsePtrType)(v)
}

func (*targetResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetResponse)(nil)).Elem()
}

func (i *targetResponsePtrType) ToTargetResponsePtrOutput() TargetResponsePtrOutput {
	return i.ToTargetResponsePtrOutputWithContext(context.Background())
}

func (i *targetResponsePtrType) ToTargetResponsePtrOutputWithContext(ctx context.Context) TargetResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetResponsePtrOutput)
}

type TargetResponseOutput struct{ *pulumi.OutputState }

func (TargetResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetResponse)(nil)).Elem()
}

func (o TargetResponseOutput) ToTargetResponseOutput() TargetResponseOutput {
	return o
}

func (o TargetResponseOutput) ToTargetResponseOutputWithContext(ctx context.Context) TargetResponseOutput {
	return o
}

func (o TargetResponseOutput) ToTargetResponsePtrOutput() TargetResponsePtrOutput {
	return o.ToTargetResponsePtrOutputWithContext(context.Background())
}

func (o TargetResponseOutput) ToTargetResponsePtrOutputWithContext(ctx context.Context) TargetResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetResponse) *TargetResponse {
		return &v
	}).(TargetResponsePtrOutput)
}

func (o TargetResponseOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.Digest }).(pulumi.StringPtrOutput)
}

func (o TargetResponseOutput) Length() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v TargetResponse) *float64 { return v.Length }).(pulumi.Float64PtrOutput)
}

func (o TargetResponseOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.MediaType }).(pulumi.StringPtrOutput)
}

func (o TargetResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TargetResponseOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.Repository }).(pulumi.StringPtrOutput)
}

func (o TargetResponseOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v TargetResponse) *float64 { return v.Size }).(pulumi.Float64PtrOutput)
}

func (o TargetResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o TargetResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

func (o TargetResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TargetResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type TargetResponsePtrOutput struct{ *pulumi.OutputState }

func (TargetResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetResponse)(nil)).Elem()
}

func (o TargetResponsePtrOutput) ToTargetResponsePtrOutput() TargetResponsePtrOutput {
	return o
}

func (o TargetResponsePtrOutput) ToTargetResponsePtrOutputWithContext(ctx context.Context) TargetResponsePtrOutput {
	return o
}

func (o TargetResponsePtrOutput) Elem() TargetResponseOutput {
	return o.ApplyT(func(v *TargetResponse) TargetResponse {
		if v != nil {
			return *v
		}
		var ret TargetResponse
		return ret
	}).(TargetResponseOutput)
}

func (o TargetResponsePtrOutput) Digest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Digest
	}).(pulumi.StringPtrOutput)
}

func (o TargetResponsePtrOutput) Length() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *TargetResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Length
	}).(pulumi.Float64PtrOutput)
}

func (o TargetResponsePtrOutput) MediaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.MediaType
	}).(pulumi.StringPtrOutput)
}

func (o TargetResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o TargetResponsePtrOutput) Repository() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Repository
	}).(pulumi.StringPtrOutput)
}

func (o TargetResponsePtrOutput) Size() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *TargetResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.Float64PtrOutput)
}

func (o TargetResponsePtrOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tag
	}).(pulumi.StringPtrOutput)
}

func (o TargetResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

func (o TargetResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TargetResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkRule struct {
	Action                   *string `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRule) Defaults() *VirtualNetworkRule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	VirtualNetworkResourceId pulumi.StringInput    `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRule) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Action                   *string `pulumi:"action"`
	VirtualNetworkResourceId string  `pulumi:"virtualNetworkResourceId"`
}


func (val *VirtualNetworkRuleResponse) Defaults() *VirtualNetworkRuleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Action) {
		action_ := "Allow"
		tmp.Action = &action_
	}
	return &tmp
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Action                   pulumi.StringPtrInput `pulumi:"action"`
	VirtualNetworkResourceId pulumi.StringInput    `pulumi:"virtualNetworkResourceId"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) VirtualNetworkResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) string { return v.VirtualNetworkResourceId }).(pulumi.StringOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ActorResponseOutput{})
	pulumi.RegisterOutputType(ActorResponsePtrOutput{})
	pulumi.RegisterOutputType(EventContentResponseOutput{})
	pulumi.RegisterOutputType(EventContentResponsePtrOutput{})
	pulumi.RegisterOutputType(EventRequestMessageResponseOutput{})
	pulumi.RegisterOutputType(EventRequestMessageResponsePtrOutput{})
	pulumi.RegisterOutputType(EventResponseOutput{})
	pulumi.RegisterOutputType(EventResponseArrayOutput{})
	pulumi.RegisterOutputType(EventResponseMessageResponseOutput{})
	pulumi.RegisterOutputType(EventResponseMessageResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistryPasswordResponseOutput{})
	pulumi.RegisterOutputType(RegistryPasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(RequestResponseOutput{})
	pulumi.RegisterOutputType(RequestResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
	pulumi.RegisterOutputType(SourceResponsePtrOutput{})
	pulumi.RegisterOutputType(StatusResponseOutput{})
	pulumi.RegisterOutputType(StatusResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StorageAccountPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetResponseOutput{})
	pulumi.RegisterOutputType(TargetResponsePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
