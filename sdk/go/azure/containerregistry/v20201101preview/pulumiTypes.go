


package v20201101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ActivationPropertiesResponse struct {
	Status string `pulumi:"status"`
}





type ActivationPropertiesResponseInput interface {
	pulumi.Input

	ToActivationPropertiesResponseOutput() ActivationPropertiesResponseOutput
	ToActivationPropertiesResponseOutputWithContext(context.Context) ActivationPropertiesResponseOutput
}

type ActivationPropertiesResponseArgs struct {
	Status pulumi.StringInput `pulumi:"status"`
}

func (ActivationPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivationPropertiesResponse)(nil)).Elem()
}

func (i ActivationPropertiesResponseArgs) ToActivationPropertiesResponseOutput() ActivationPropertiesResponseOutput {
	return i.ToActivationPropertiesResponseOutputWithContext(context.Background())
}

func (i ActivationPropertiesResponseArgs) ToActivationPropertiesResponseOutputWithContext(ctx context.Context) ActivationPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationPropertiesResponseOutput)
}

func (i ActivationPropertiesResponseArgs) ToActivationPropertiesResponsePtrOutput() ActivationPropertiesResponsePtrOutput {
	return i.ToActivationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ActivationPropertiesResponseArgs) ToActivationPropertiesResponsePtrOutputWithContext(ctx context.Context) ActivationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationPropertiesResponseOutput).ToActivationPropertiesResponsePtrOutputWithContext(ctx)
}









type ActivationPropertiesResponsePtrInput interface {
	pulumi.Input

	ToActivationPropertiesResponsePtrOutput() ActivationPropertiesResponsePtrOutput
	ToActivationPropertiesResponsePtrOutputWithContext(context.Context) ActivationPropertiesResponsePtrOutput
}

type activationPropertiesResponsePtrType ActivationPropertiesResponseArgs

func ActivationPropertiesResponsePtr(v *ActivationPropertiesResponseArgs) ActivationPropertiesResponsePtrInput {
	return (*activationPropertiesResponsePtrType)(v)
}

func (*activationPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivationPropertiesResponse)(nil)).Elem()
}

func (i *activationPropertiesResponsePtrType) ToActivationPropertiesResponsePtrOutput() ActivationPropertiesResponsePtrOutput {
	return i.ToActivationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *activationPropertiesResponsePtrType) ToActivationPropertiesResponsePtrOutputWithContext(ctx context.Context) ActivationPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ActivationPropertiesResponsePtrOutput)
}

type ActivationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ActivationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ActivationPropertiesResponse)(nil)).Elem()
}

func (o ActivationPropertiesResponseOutput) ToActivationPropertiesResponseOutput() ActivationPropertiesResponseOutput {
	return o
}

func (o ActivationPropertiesResponseOutput) ToActivationPropertiesResponseOutputWithContext(ctx context.Context) ActivationPropertiesResponseOutput {
	return o
}

func (o ActivationPropertiesResponseOutput) ToActivationPropertiesResponsePtrOutput() ActivationPropertiesResponsePtrOutput {
	return o.ToActivationPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ActivationPropertiesResponseOutput) ToActivationPropertiesResponsePtrOutputWithContext(ctx context.Context) ActivationPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ActivationPropertiesResponse) *ActivationPropertiesResponse {
		return &v
	}).(ActivationPropertiesResponsePtrOutput)
}

func (o ActivationPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ActivationPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type ActivationPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ActivationPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ActivationPropertiesResponse)(nil)).Elem()
}

func (o ActivationPropertiesResponsePtrOutput) ToActivationPropertiesResponsePtrOutput() ActivationPropertiesResponsePtrOutput {
	return o
}

func (o ActivationPropertiesResponsePtrOutput) ToActivationPropertiesResponsePtrOutputWithContext(ctx context.Context) ActivationPropertiesResponsePtrOutput {
	return o
}

func (o ActivationPropertiesResponsePtrOutput) Elem() ActivationPropertiesResponseOutput {
	return o.ApplyT(func(v *ActivationPropertiesResponse) ActivationPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ActivationPropertiesResponse
		return ret
	}).(ActivationPropertiesResponseOutput)
}

func (o ActivationPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ActivationPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

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

type EncryptionProperty struct {
	KeyVaultProperties *KeyVaultProperties `pulumi:"keyVaultProperties"`
	Status             *string             `pulumi:"status"`
}





type EncryptionPropertyInput interface {
	pulumi.Input

	ToEncryptionPropertyOutput() EncryptionPropertyOutput
	ToEncryptionPropertyOutputWithContext(context.Context) EncryptionPropertyOutput
}

type EncryptionPropertyArgs struct {
	KeyVaultProperties KeyVaultPropertiesPtrInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringPtrInput      `pulumi:"status"`
}

func (EncryptionPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperty)(nil)).Elem()
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyOutput() EncryptionPropertyOutput {
	return i.ToEncryptionPropertyOutputWithContext(context.Background())
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyOutputWithContext(ctx context.Context) EncryptionPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyOutput)
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return i.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertyArgs) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyOutput).ToEncryptionPropertyPtrOutputWithContext(ctx)
}









type EncryptionPropertyPtrInput interface {
	pulumi.Input

	ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput
	ToEncryptionPropertyPtrOutputWithContext(context.Context) EncryptionPropertyPtrOutput
}

type encryptionPropertyPtrType EncryptionPropertyArgs

func EncryptionPropertyPtr(v *EncryptionPropertyArgs) EncryptionPropertyPtrInput {
	return (*encryptionPropertyPtrType)(v)
}

func (*encryptionPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperty)(nil)).Elem()
}

func (i *encryptionPropertyPtrType) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return i.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertyPtrType) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyPtrOutput)
}

type EncryptionPropertyOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperty)(nil)).Elem()
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyOutput() EncryptionPropertyOutput {
	return o
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyOutputWithContext(ctx context.Context) EncryptionPropertyOutput {
	return o
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return o.ToEncryptionPropertyPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertyOutput) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperty) *EncryptionProperty {
		return &v
	}).(EncryptionPropertyPtrOutput)
}

func (o EncryptionPropertyOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v EncryptionProperty) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPropertyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionProperty) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type EncryptionPropertyPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperty)(nil)).Elem()
}

func (o EncryptionPropertyPtrOutput) ToEncryptionPropertyPtrOutput() EncryptionPropertyPtrOutput {
	return o
}

func (o EncryptionPropertyPtrOutput) ToEncryptionPropertyPtrOutputWithContext(ctx context.Context) EncryptionPropertyPtrOutput {
	return o
}

func (o EncryptionPropertyPtrOutput) Elem() EncryptionPropertyOutput {
	return o.ApplyT(func(v *EncryptionProperty) EncryptionProperty {
		if v != nil {
			return *v
		}
		var ret EncryptionProperty
		return ret
	}).(EncryptionPropertyOutput)
}

func (o EncryptionPropertyPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

func (o EncryptionPropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionProperty) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type EncryptionPropertyResponse struct {
	KeyVaultProperties *KeyVaultPropertiesResponse `pulumi:"keyVaultProperties"`
	Status             *string                     `pulumi:"status"`
}





type EncryptionPropertyResponseInput interface {
	pulumi.Input

	ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput
	ToEncryptionPropertyResponseOutputWithContext(context.Context) EncryptionPropertyResponseOutput
}

type EncryptionPropertyResponseArgs struct {
	KeyVaultProperties KeyVaultPropertiesResponsePtrInput `pulumi:"keyVaultProperties"`
	Status             pulumi.StringPtrInput              `pulumi:"status"`
}

func (EncryptionPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertyResponse)(nil)).Elem()
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput {
	return i.ToEncryptionPropertyResponseOutputWithContext(context.Background())
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponseOutputWithContext(ctx context.Context) EncryptionPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponseOutput)
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return i.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (i EncryptionPropertyResponseArgs) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponseOutput).ToEncryptionPropertyResponsePtrOutputWithContext(ctx)
}









type EncryptionPropertyResponsePtrInput interface {
	pulumi.Input

	ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput
	ToEncryptionPropertyResponsePtrOutputWithContext(context.Context) EncryptionPropertyResponsePtrOutput
}

type encryptionPropertyResponsePtrType EncryptionPropertyResponseArgs

func EncryptionPropertyResponsePtr(v *EncryptionPropertyResponseArgs) EncryptionPropertyResponsePtrInput {
	return (*encryptionPropertyResponsePtrType)(v)
}

func (*encryptionPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertyResponse)(nil)).Elem()
}

func (i *encryptionPropertyResponsePtrType) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return i.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *encryptionPropertyResponsePtrType) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertyResponsePtrOutput)
}

type EncryptionPropertyResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertyResponse)(nil)).Elem()
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponseOutput() EncryptionPropertyResponseOutput {
	return o
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponseOutputWithContext(ctx context.Context) EncryptionPropertyResponseOutput {
	return o
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return o.ToEncryptionPropertyResponsePtrOutputWithContext(context.Background())
}

func (o EncryptionPropertyResponseOutput) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertyResponse) *EncryptionPropertyResponse {
		return &v
	}).(EncryptionPropertyResponsePtrOutput)
}

func (o EncryptionPropertyResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionPropertyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type EncryptionPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertyResponse)(nil)).Elem()
}

func (o EncryptionPropertyResponsePtrOutput) ToEncryptionPropertyResponsePtrOutput() EncryptionPropertyResponsePtrOutput {
	return o
}

func (o EncryptionPropertyResponsePtrOutput) ToEncryptionPropertyResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertyResponsePtrOutput {
	return o
}

func (o EncryptionPropertyResponsePtrOutput) Elem() EncryptionPropertyResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) EncryptionPropertyResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertyResponse
		return ret
	}).(EncryptionPropertyResponseOutput)
}

func (o EncryptionPropertyResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o EncryptionPropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
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

type ExportPipelineTargetProperties struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}





type ExportPipelineTargetPropertiesInput interface {
	pulumi.Input

	ToExportPipelineTargetPropertiesOutput() ExportPipelineTargetPropertiesOutput
	ToExportPipelineTargetPropertiesOutputWithContext(context.Context) ExportPipelineTargetPropertiesOutput
}

type ExportPipelineTargetPropertiesArgs struct {
	KeyVaultUri pulumi.StringInput    `pulumi:"keyVaultUri"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
}

func (ExportPipelineTargetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetProperties)(nil)).Elem()
}

func (i ExportPipelineTargetPropertiesArgs) ToExportPipelineTargetPropertiesOutput() ExportPipelineTargetPropertiesOutput {
	return i.ToExportPipelineTargetPropertiesOutputWithContext(context.Background())
}

func (i ExportPipelineTargetPropertiesArgs) ToExportPipelineTargetPropertiesOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesOutput)
}

func (i ExportPipelineTargetPropertiesArgs) ToExportPipelineTargetPropertiesPtrOutput() ExportPipelineTargetPropertiesPtrOutput {
	return i.ToExportPipelineTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i ExportPipelineTargetPropertiesArgs) ToExportPipelineTargetPropertiesPtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesOutput).ToExportPipelineTargetPropertiesPtrOutputWithContext(ctx)
}









type ExportPipelineTargetPropertiesPtrInput interface {
	pulumi.Input

	ToExportPipelineTargetPropertiesPtrOutput() ExportPipelineTargetPropertiesPtrOutput
	ToExportPipelineTargetPropertiesPtrOutputWithContext(context.Context) ExportPipelineTargetPropertiesPtrOutput
}

type exportPipelineTargetPropertiesPtrType ExportPipelineTargetPropertiesArgs

func ExportPipelineTargetPropertiesPtr(v *ExportPipelineTargetPropertiesArgs) ExportPipelineTargetPropertiesPtrInput {
	return (*exportPipelineTargetPropertiesPtrType)(v)
}

func (*exportPipelineTargetPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportPipelineTargetProperties)(nil)).Elem()
}

func (i *exportPipelineTargetPropertiesPtrType) ToExportPipelineTargetPropertiesPtrOutput() ExportPipelineTargetPropertiesPtrOutput {
	return i.ToExportPipelineTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i *exportPipelineTargetPropertiesPtrType) ToExportPipelineTargetPropertiesPtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesPtrOutput)
}

type ExportPipelineTargetPropertiesOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetProperties)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesOutput) ToExportPipelineTargetPropertiesOutput() ExportPipelineTargetPropertiesOutput {
	return o
}

func (o ExportPipelineTargetPropertiesOutput) ToExportPipelineTargetPropertiesOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesOutput {
	return o
}

func (o ExportPipelineTargetPropertiesOutput) ToExportPipelineTargetPropertiesPtrOutput() ExportPipelineTargetPropertiesPtrOutput {
	return o.ToExportPipelineTargetPropertiesPtrOutputWithContext(context.Background())
}

func (o ExportPipelineTargetPropertiesOutput) ToExportPipelineTargetPropertiesPtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportPipelineTargetProperties) *ExportPipelineTargetProperties {
		return &v
	}).(ExportPipelineTargetPropertiesPtrOutput)
}

func (o ExportPipelineTargetPropertiesOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ExportPipelineTargetProperties) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ExportPipelineTargetPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ExportPipelineTargetPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportPipelineTargetProperties)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesPtrOutput) ToExportPipelineTargetPropertiesPtrOutput() ExportPipelineTargetPropertiesPtrOutput {
	return o
}

func (o ExportPipelineTargetPropertiesPtrOutput) ToExportPipelineTargetPropertiesPtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesPtrOutput {
	return o
}

func (o ExportPipelineTargetPropertiesPtrOutput) Elem() ExportPipelineTargetPropertiesOutput {
	return o.ApplyT(func(v *ExportPipelineTargetProperties) ExportPipelineTargetProperties {
		if v != nil {
			return *v
		}
		var ret ExportPipelineTargetProperties
		return ret
	}).(ExportPipelineTargetPropertiesOutput)
}

func (o ExportPipelineTargetPropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ExportPipelineTargetPropertiesResponse struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}





type ExportPipelineTargetPropertiesResponseInput interface {
	pulumi.Input

	ToExportPipelineTargetPropertiesResponseOutput() ExportPipelineTargetPropertiesResponseOutput
	ToExportPipelineTargetPropertiesResponseOutputWithContext(context.Context) ExportPipelineTargetPropertiesResponseOutput
}

type ExportPipelineTargetPropertiesResponseArgs struct {
	KeyVaultUri pulumi.StringInput    `pulumi:"keyVaultUri"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
}

func (ExportPipelineTargetPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetPropertiesResponse)(nil)).Elem()
}

func (i ExportPipelineTargetPropertiesResponseArgs) ToExportPipelineTargetPropertiesResponseOutput() ExportPipelineTargetPropertiesResponseOutput {
	return i.ToExportPipelineTargetPropertiesResponseOutputWithContext(context.Background())
}

func (i ExportPipelineTargetPropertiesResponseArgs) ToExportPipelineTargetPropertiesResponseOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesResponseOutput)
}

func (i ExportPipelineTargetPropertiesResponseArgs) ToExportPipelineTargetPropertiesResponsePtrOutput() ExportPipelineTargetPropertiesResponsePtrOutput {
	return i.ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ExportPipelineTargetPropertiesResponseArgs) ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesResponseOutput).ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(ctx)
}









type ExportPipelineTargetPropertiesResponsePtrInput interface {
	pulumi.Input

	ToExportPipelineTargetPropertiesResponsePtrOutput() ExportPipelineTargetPropertiesResponsePtrOutput
	ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(context.Context) ExportPipelineTargetPropertiesResponsePtrOutput
}

type exportPipelineTargetPropertiesResponsePtrType ExportPipelineTargetPropertiesResponseArgs

func ExportPipelineTargetPropertiesResponsePtr(v *ExportPipelineTargetPropertiesResponseArgs) ExportPipelineTargetPropertiesResponsePtrInput {
	return (*exportPipelineTargetPropertiesResponsePtrType)(v)
}

func (*exportPipelineTargetPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportPipelineTargetPropertiesResponse)(nil)).Elem()
}

func (i *exportPipelineTargetPropertiesResponsePtrType) ToExportPipelineTargetPropertiesResponsePtrOutput() ExportPipelineTargetPropertiesResponsePtrOutput {
	return i.ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *exportPipelineTargetPropertiesResponsePtrType) ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExportPipelineTargetPropertiesResponsePtrOutput)
}

type ExportPipelineTargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportPipelineTargetPropertiesResponse)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesResponseOutput) ToExportPipelineTargetPropertiesResponseOutput() ExportPipelineTargetPropertiesResponseOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponseOutput) ToExportPipelineTargetPropertiesResponseOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponseOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponseOutput) ToExportPipelineTargetPropertiesResponsePtrOutput() ExportPipelineTargetPropertiesResponsePtrOutput {
	return o.ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ExportPipelineTargetPropertiesResponseOutput) ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportPipelineTargetPropertiesResponse) *ExportPipelineTargetPropertiesResponse {
		return &v
	}).(ExportPipelineTargetPropertiesResponsePtrOutput)
}

func (o ExportPipelineTargetPropertiesResponseOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ExportPipelineTargetPropertiesResponse) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ExportPipelineTargetPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExportPipelineTargetPropertiesResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ExportPipelineTargetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ExportPipelineTargetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportPipelineTargetPropertiesResponse)(nil)).Elem()
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) ToExportPipelineTargetPropertiesResponsePtrOutput() ExportPipelineTargetPropertiesResponsePtrOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) ToExportPipelineTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) ExportPipelineTargetPropertiesResponsePtrOutput {
	return o
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) Elem() ExportPipelineTargetPropertiesResponseOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) ExportPipelineTargetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ExportPipelineTargetPropertiesResponse
		return ret
	}).(ExportPipelineTargetPropertiesResponseOutput)
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ExportPipelineTargetPropertiesResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExportPipelineTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
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

type IdentityProperties struct {
	PrincipalId            *string                           `pulumi:"principalId"`
	TenantId               *string                           `pulumi:"tenantId"`
	Type                   *ResourceIdentityType             `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityProperties `pulumi:"userAssignedIdentities"`
}





type IdentityPropertiesInput interface {
	pulumi.Input

	ToIdentityPropertiesOutput() IdentityPropertiesOutput
	ToIdentityPropertiesOutputWithContext(context.Context) IdentityPropertiesOutput
}

type IdentityPropertiesArgs struct {
	PrincipalId            pulumi.StringPtrInput          `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput          `pulumi:"tenantId"`
	Type                   ResourceIdentityTypePtrInput   `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return i.ToIdentityPropertiesOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput)
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesArgs) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesOutput).ToIdentityPropertiesPtrOutputWithContext(ctx)
}









type IdentityPropertiesPtrInput interface {
	pulumi.Input

	ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput
	ToIdentityPropertiesPtrOutputWithContext(context.Context) IdentityPropertiesPtrOutput
}

type identityPropertiesPtrType IdentityPropertiesArgs

func IdentityPropertiesPtr(v *IdentityPropertiesArgs) IdentityPropertiesPtrInput {
	return (*identityPropertiesPtrType)(v)
}

func (*identityPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return i.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (i *identityPropertiesPtrType) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesPtrOutput)
}

type IdentityPropertiesOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutput() IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesOutputWithContext(ctx context.Context) IdentityPropertiesOutput {
	return o
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o.ToIdentityPropertiesPtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityProperties) *IdentityProperties {
		return &v
	}).(IdentityPropertiesPtrOutput)
}

func (o IdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v IdentityProperties) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPropertiesOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v IdentityProperties) map[string]UserIdentityProperties { return v.UserAssignedIdentities }).(UserIdentityPropertiesMapOutput)
}

type IdentityPropertiesPtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityProperties)(nil)).Elem()
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutput() IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) ToIdentityPropertiesPtrOutputWithContext(ctx context.Context) IdentityPropertiesPtrOutput {
	return o
}

func (o IdentityPropertiesPtrOutput) Elem() IdentityPropertiesOutput {
	return o.ApplyT(func(v *IdentityProperties) IdentityProperties {
		if v != nil {
			return *v
		}
		var ret IdentityProperties
		return ret
	}).(IdentityPropertiesOutput)
}

func (o IdentityPropertiesPtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *IdentityProperties) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPropertiesPtrOutput) UserAssignedIdentities() UserIdentityPropertiesMapOutput {
	return o.ApplyT(func(v *IdentityProperties) map[string]UserIdentityProperties {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesMapOutput)
}

type IdentityPropertiesResponse struct {
	PrincipalId            *string                                   `pulumi:"principalId"`
	TenantId               *string                                   `pulumi:"tenantId"`
	Type                   *string                                   `pulumi:"type"`
	UserAssignedIdentities map[string]UserIdentityPropertiesResponse `pulumi:"userAssignedIdentities"`
}





type IdentityPropertiesResponseInput interface {
	pulumi.Input

	ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput
	ToIdentityPropertiesResponseOutputWithContext(context.Context) IdentityPropertiesResponseOutput
}

type IdentityPropertiesResponseArgs struct {
	PrincipalId            pulumi.StringPtrInput                  `pulumi:"principalId"`
	TenantId               pulumi.StringPtrInput                  `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                  `pulumi:"type"`
	UserAssignedIdentities UserIdentityPropertiesResponseMapInput `pulumi:"userAssignedIdentities"`
}

func (IdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return i.ToIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput)
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i IdentityPropertiesResponseArgs) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponseOutput).ToIdentityPropertiesResponsePtrOutputWithContext(ctx)
}









type IdentityPropertiesResponsePtrInput interface {
	pulumi.Input

	ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput
	ToIdentityPropertiesResponsePtrOutputWithContext(context.Context) IdentityPropertiesResponsePtrOutput
}

type identityPropertiesResponsePtrType IdentityPropertiesResponseArgs

func IdentityPropertiesResponsePtr(v *IdentityPropertiesResponseArgs) IdentityPropertiesResponsePtrInput {
	return (*identityPropertiesResponsePtrType)(v)
}

func (*identityPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return i.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *identityPropertiesResponsePtrType) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPropertiesResponsePtrOutput)
}

type IdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutput() IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponseOutputWithContext(ctx context.Context) IdentityPropertiesResponseOutput {
	return o
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o.ToIdentityPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o IdentityPropertiesResponseOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityPropertiesResponse) *IdentityPropertiesResponse {
		return &v
	}).(IdentityPropertiesResponsePtrOutput)
}

func (o IdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponseOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v IdentityPropertiesResponse) map[string]UserIdentityPropertiesResponse {
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type IdentityPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityPropertiesResponse)(nil)).Elem()
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutput() IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) ToIdentityPropertiesResponsePtrOutputWithContext(ctx context.Context) IdentityPropertiesResponsePtrOutput {
	return o
}

func (o IdentityPropertiesResponsePtrOutput) Elem() IdentityPropertiesResponseOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) IdentityPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret IdentityPropertiesResponse
		return ret
	}).(IdentityPropertiesResponseOutput)
}

func (o IdentityPropertiesResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityPropertiesResponsePtrOutput) UserAssignedIdentities() UserIdentityPropertiesResponseMapOutput {
	return o.ApplyT(func(v *IdentityPropertiesResponse) map[string]UserIdentityPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserIdentityPropertiesResponseMapOutput)
}

type ImportPipelineSourceProperties struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}


func (val *ImportPipelineSourceProperties) Defaults() *ImportPipelineSourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlobContainer"
		tmp.Type = &type_
	}
	return &tmp
}





type ImportPipelineSourcePropertiesInput interface {
	pulumi.Input

	ToImportPipelineSourcePropertiesOutput() ImportPipelineSourcePropertiesOutput
	ToImportPipelineSourcePropertiesOutputWithContext(context.Context) ImportPipelineSourcePropertiesOutput
}

type ImportPipelineSourcePropertiesArgs struct {
	KeyVaultUri pulumi.StringInput    `pulumi:"keyVaultUri"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
}

func (ImportPipelineSourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourceProperties)(nil)).Elem()
}

func (i ImportPipelineSourcePropertiesArgs) ToImportPipelineSourcePropertiesOutput() ImportPipelineSourcePropertiesOutput {
	return i.ToImportPipelineSourcePropertiesOutputWithContext(context.Background())
}

func (i ImportPipelineSourcePropertiesArgs) ToImportPipelineSourcePropertiesOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesOutput)
}

func (i ImportPipelineSourcePropertiesArgs) ToImportPipelineSourcePropertiesPtrOutput() ImportPipelineSourcePropertiesPtrOutput {
	return i.ToImportPipelineSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ImportPipelineSourcePropertiesArgs) ToImportPipelineSourcePropertiesPtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesOutput).ToImportPipelineSourcePropertiesPtrOutputWithContext(ctx)
}









type ImportPipelineSourcePropertiesPtrInput interface {
	pulumi.Input

	ToImportPipelineSourcePropertiesPtrOutput() ImportPipelineSourcePropertiesPtrOutput
	ToImportPipelineSourcePropertiesPtrOutputWithContext(context.Context) ImportPipelineSourcePropertiesPtrOutput
}

type importPipelineSourcePropertiesPtrType ImportPipelineSourcePropertiesArgs

func ImportPipelineSourcePropertiesPtr(v *ImportPipelineSourcePropertiesArgs) ImportPipelineSourcePropertiesPtrInput {
	return (*importPipelineSourcePropertiesPtrType)(v)
}

func (*importPipelineSourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipelineSourceProperties)(nil)).Elem()
}

func (i *importPipelineSourcePropertiesPtrType) ToImportPipelineSourcePropertiesPtrOutput() ImportPipelineSourcePropertiesPtrOutput {
	return i.ToImportPipelineSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *importPipelineSourcePropertiesPtrType) ToImportPipelineSourcePropertiesPtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesPtrOutput)
}

type ImportPipelineSourcePropertiesOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourceProperties)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesOutput) ToImportPipelineSourcePropertiesOutput() ImportPipelineSourcePropertiesOutput {
	return o
}

func (o ImportPipelineSourcePropertiesOutput) ToImportPipelineSourcePropertiesOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesOutput {
	return o
}

func (o ImportPipelineSourcePropertiesOutput) ToImportPipelineSourcePropertiesPtrOutput() ImportPipelineSourcePropertiesPtrOutput {
	return o.ToImportPipelineSourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ImportPipelineSourcePropertiesOutput) ToImportPipelineSourcePropertiesPtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImportPipelineSourceProperties) *ImportPipelineSourceProperties {
		return &v
	}).(ImportPipelineSourcePropertiesPtrOutput)
}

func (o ImportPipelineSourcePropertiesOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ImportPipelineSourceProperties) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ImportPipelineSourcePropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourceProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourceProperties) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ImportPipelineSourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipelineSourceProperties)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesPtrOutput) ToImportPipelineSourcePropertiesPtrOutput() ImportPipelineSourcePropertiesPtrOutput {
	return o
}

func (o ImportPipelineSourcePropertiesPtrOutput) ToImportPipelineSourcePropertiesPtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesPtrOutput {
	return o
}

func (o ImportPipelineSourcePropertiesPtrOutput) Elem() ImportPipelineSourcePropertiesOutput {
	return o.ApplyT(func(v *ImportPipelineSourceProperties) ImportPipelineSourceProperties {
		if v != nil {
			return *v
		}
		var ret ImportPipelineSourceProperties
		return ret
	}).(ImportPipelineSourcePropertiesOutput)
}

func (o ImportPipelineSourcePropertiesPtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourceProperties) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type ImportPipelineSourcePropertiesResponse struct {
	KeyVaultUri string  `pulumi:"keyVaultUri"`
	Type        *string `pulumi:"type"`
	Uri         *string `pulumi:"uri"`
}


func (val *ImportPipelineSourcePropertiesResponse) Defaults() *ImportPipelineSourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlobContainer"
		tmp.Type = &type_
	}
	return &tmp
}





type ImportPipelineSourcePropertiesResponseInput interface {
	pulumi.Input

	ToImportPipelineSourcePropertiesResponseOutput() ImportPipelineSourcePropertiesResponseOutput
	ToImportPipelineSourcePropertiesResponseOutputWithContext(context.Context) ImportPipelineSourcePropertiesResponseOutput
}

type ImportPipelineSourcePropertiesResponseArgs struct {
	KeyVaultUri pulumi.StringInput    `pulumi:"keyVaultUri"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
}

func (ImportPipelineSourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourcePropertiesResponse)(nil)).Elem()
}

func (i ImportPipelineSourcePropertiesResponseArgs) ToImportPipelineSourcePropertiesResponseOutput() ImportPipelineSourcePropertiesResponseOutput {
	return i.ToImportPipelineSourcePropertiesResponseOutputWithContext(context.Background())
}

func (i ImportPipelineSourcePropertiesResponseArgs) ToImportPipelineSourcePropertiesResponseOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesResponseOutput)
}

func (i ImportPipelineSourcePropertiesResponseArgs) ToImportPipelineSourcePropertiesResponsePtrOutput() ImportPipelineSourcePropertiesResponsePtrOutput {
	return i.ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ImportPipelineSourcePropertiesResponseArgs) ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesResponseOutput).ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(ctx)
}









type ImportPipelineSourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToImportPipelineSourcePropertiesResponsePtrOutput() ImportPipelineSourcePropertiesResponsePtrOutput
	ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(context.Context) ImportPipelineSourcePropertiesResponsePtrOutput
}

type importPipelineSourcePropertiesResponsePtrType ImportPipelineSourcePropertiesResponseArgs

func ImportPipelineSourcePropertiesResponsePtr(v *ImportPipelineSourcePropertiesResponseArgs) ImportPipelineSourcePropertiesResponsePtrInput {
	return (*importPipelineSourcePropertiesResponsePtrType)(v)
}

func (*importPipelineSourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipelineSourcePropertiesResponse)(nil)).Elem()
}

func (i *importPipelineSourcePropertiesResponsePtrType) ToImportPipelineSourcePropertiesResponsePtrOutput() ImportPipelineSourcePropertiesResponsePtrOutput {
	return i.ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *importPipelineSourcePropertiesResponsePtrType) ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImportPipelineSourcePropertiesResponsePtrOutput)
}

type ImportPipelineSourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImportPipelineSourcePropertiesResponse)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesResponseOutput) ToImportPipelineSourcePropertiesResponseOutput() ImportPipelineSourcePropertiesResponseOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponseOutput) ToImportPipelineSourcePropertiesResponseOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponseOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponseOutput) ToImportPipelineSourcePropertiesResponsePtrOutput() ImportPipelineSourcePropertiesResponsePtrOutput {
	return o.ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ImportPipelineSourcePropertiesResponseOutput) ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImportPipelineSourcePropertiesResponse) *ImportPipelineSourcePropertiesResponse {
		return &v
	}).(ImportPipelineSourcePropertiesResponsePtrOutput)
}

func (o ImportPipelineSourcePropertiesResponseOutput) KeyVaultUri() pulumi.StringOutput {
	return o.ApplyT(func(v ImportPipelineSourcePropertiesResponse) string { return v.KeyVaultUri }).(pulumi.StringOutput)
}

func (o ImportPipelineSourcePropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourcePropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImportPipelineSourcePropertiesResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

type ImportPipelineSourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ImportPipelineSourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImportPipelineSourcePropertiesResponse)(nil)).Elem()
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) ToImportPipelineSourcePropertiesResponsePtrOutput() ImportPipelineSourcePropertiesResponsePtrOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) ToImportPipelineSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) ImportPipelineSourcePropertiesResponsePtrOutput {
	return o
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) Elem() ImportPipelineSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) ImportPipelineSourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ImportPipelineSourcePropertiesResponse
		return ret
	}).(ImportPipelineSourcePropertiesResponseOutput)
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) KeyVaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.KeyVaultUri
	}).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ImportPipelineSourcePropertiesResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImportPipelineSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

type KeyVaultProperties struct {
	Identity      *string `pulumi:"identity"`
	KeyIdentifier *string `pulumi:"keyIdentifier"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	Identity      pulumi.StringPtrInput `pulumi:"identity"`
	KeyIdentifier pulumi.StringPtrInput `pulumi:"keyIdentifier"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	Identity                 *string `pulumi:"identity"`
	KeyIdentifier            *string `pulumi:"keyIdentifier"`
	KeyRotationEnabled       bool    `pulumi:"keyRotationEnabled"`
	LastKeyRotationTimestamp string  `pulumi:"lastKeyRotationTimestamp"`
	VersionedKeyIdentifier   string  `pulumi:"versionedKeyIdentifier"`
}





type KeyVaultPropertiesResponseInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput
	ToKeyVaultPropertiesResponseOutputWithContext(context.Context) KeyVaultPropertiesResponseOutput
}

type KeyVaultPropertiesResponseArgs struct {
	Identity                 pulumi.StringPtrInput `pulumi:"identity"`
	KeyIdentifier            pulumi.StringPtrInput `pulumi:"keyIdentifier"`
	KeyRotationEnabled       pulumi.BoolInput      `pulumi:"keyRotationEnabled"`
	LastKeyRotationTimestamp pulumi.StringInput    `pulumi:"lastKeyRotationTimestamp"`
	VersionedKeyIdentifier   pulumi.StringInput    `pulumi:"versionedKeyIdentifier"`
}

func (KeyVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return i.ToKeyVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput)
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesResponseArgs) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponseOutput).ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type KeyVaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput
	ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Context) KeyVaultPropertiesResponsePtrOutput
}

type keyVaultPropertiesResponsePtrType KeyVaultPropertiesResponseArgs

func KeyVaultPropertiesResponsePtr(v *KeyVaultPropertiesResponseArgs) KeyVaultPropertiesResponsePtrInput {
	return (*keyVaultPropertiesResponsePtrType)(v)
}

func (*keyVaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return i.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesResponsePtrType) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesResponsePtrOutput)
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o.ToKeyVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultPropertiesResponse) *KeyVaultPropertiesResponse {
		return &v
	}).(KeyVaultPropertiesResponsePtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.Identity }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyIdentifier }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyRotationEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) bool { return v.KeyRotationEnabled }).(pulumi.BoolOutput)
}

func (o KeyVaultPropertiesResponseOutput) LastKeyRotationTimestamp() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.LastKeyRotationTimestamp }).(pulumi.StringOutput)
}

func (o KeyVaultPropertiesResponseOutput) VersionedKeyIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) string { return v.VersionedKeyIdentifier }).(pulumi.StringOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) Identity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyIdentifier
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyRotationEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.KeyRotationEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) LastKeyRotationTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastKeyRotationTimestamp
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) VersionedKeyIdentifier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VersionedKeyIdentifier
	}).(pulumi.StringPtrOutput)
}

type LoggingProperties struct {
	AuditLogStatus *string `pulumi:"auditLogStatus"`
	LogLevel       *string `pulumi:"logLevel"`
}


func (val *LoggingProperties) Defaults() *LoggingProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AuditLogStatus) {
		auditLogStatus_ := "Disabled"
		tmp.AuditLogStatus = &auditLogStatus_
	}
	if isZero(tmp.LogLevel) {
		logLevel_ := "Information"
		tmp.LogLevel = &logLevel_
	}
	return &tmp
}





type LoggingPropertiesInput interface {
	pulumi.Input

	ToLoggingPropertiesOutput() LoggingPropertiesOutput
	ToLoggingPropertiesOutputWithContext(context.Context) LoggingPropertiesOutput
}

type LoggingPropertiesArgs struct {
	AuditLogStatus pulumi.StringPtrInput `pulumi:"auditLogStatus"`
	LogLevel       pulumi.StringPtrInput `pulumi:"logLevel"`
}

func (LoggingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingProperties)(nil)).Elem()
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesOutput() LoggingPropertiesOutput {
	return i.ToLoggingPropertiesOutputWithContext(context.Background())
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesOutputWithContext(ctx context.Context) LoggingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesOutput)
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return i.ToLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (i LoggingPropertiesArgs) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesOutput).ToLoggingPropertiesPtrOutputWithContext(ctx)
}









type LoggingPropertiesPtrInput interface {
	pulumi.Input

	ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput
	ToLoggingPropertiesPtrOutputWithContext(context.Context) LoggingPropertiesPtrOutput
}

type loggingPropertiesPtrType LoggingPropertiesArgs

func LoggingPropertiesPtr(v *LoggingPropertiesArgs) LoggingPropertiesPtrInput {
	return (*loggingPropertiesPtrType)(v)
}

func (*loggingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingProperties)(nil)).Elem()
}

func (i *loggingPropertiesPtrType) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return i.ToLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (i *loggingPropertiesPtrType) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesPtrOutput)
}

type LoggingPropertiesOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingProperties)(nil)).Elem()
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesOutput() LoggingPropertiesOutput {
	return o
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesOutputWithContext(ctx context.Context) LoggingPropertiesOutput {
	return o
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return o.ToLoggingPropertiesPtrOutputWithContext(context.Background())
}

func (o LoggingPropertiesOutput) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggingProperties) *LoggingProperties {
		return &v
	}).(LoggingPropertiesPtrOutput)
}

func (o LoggingPropertiesOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingProperties) *string { return v.AuditLogStatus }).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingProperties) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

type LoggingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingProperties)(nil)).Elem()
}

func (o LoggingPropertiesPtrOutput) ToLoggingPropertiesPtrOutput() LoggingPropertiesPtrOutput {
	return o
}

func (o LoggingPropertiesPtrOutput) ToLoggingPropertiesPtrOutputWithContext(ctx context.Context) LoggingPropertiesPtrOutput {
	return o
}

func (o LoggingPropertiesPtrOutput) Elem() LoggingPropertiesOutput {
	return o.ApplyT(func(v *LoggingProperties) LoggingProperties {
		if v != nil {
			return *v
		}
		var ret LoggingProperties
		return ret
	}).(LoggingPropertiesOutput)
}

func (o LoggingPropertiesPtrOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuditLogStatus
	}).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesPtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingProperties) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

type LoggingPropertiesResponse struct {
	AuditLogStatus *string `pulumi:"auditLogStatus"`
	LogLevel       *string `pulumi:"logLevel"`
}


func (val *LoggingPropertiesResponse) Defaults() *LoggingPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.AuditLogStatus) {
		auditLogStatus_ := "Disabled"
		tmp.AuditLogStatus = &auditLogStatus_
	}
	if isZero(tmp.LogLevel) {
		logLevel_ := "Information"
		tmp.LogLevel = &logLevel_
	}
	return &tmp
}





type LoggingPropertiesResponseInput interface {
	pulumi.Input

	ToLoggingPropertiesResponseOutput() LoggingPropertiesResponseOutput
	ToLoggingPropertiesResponseOutputWithContext(context.Context) LoggingPropertiesResponseOutput
}

type LoggingPropertiesResponseArgs struct {
	AuditLogStatus pulumi.StringPtrInput `pulumi:"auditLogStatus"`
	LogLevel       pulumi.StringPtrInput `pulumi:"logLevel"`
}

func (LoggingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingPropertiesResponse)(nil)).Elem()
}

func (i LoggingPropertiesResponseArgs) ToLoggingPropertiesResponseOutput() LoggingPropertiesResponseOutput {
	return i.ToLoggingPropertiesResponseOutputWithContext(context.Background())
}

func (i LoggingPropertiesResponseArgs) ToLoggingPropertiesResponseOutputWithContext(ctx context.Context) LoggingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesResponseOutput)
}

func (i LoggingPropertiesResponseArgs) ToLoggingPropertiesResponsePtrOutput() LoggingPropertiesResponsePtrOutput {
	return i.ToLoggingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i LoggingPropertiesResponseArgs) ToLoggingPropertiesResponsePtrOutputWithContext(ctx context.Context) LoggingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesResponseOutput).ToLoggingPropertiesResponsePtrOutputWithContext(ctx)
}









type LoggingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToLoggingPropertiesResponsePtrOutput() LoggingPropertiesResponsePtrOutput
	ToLoggingPropertiesResponsePtrOutputWithContext(context.Context) LoggingPropertiesResponsePtrOutput
}

type loggingPropertiesResponsePtrType LoggingPropertiesResponseArgs

func LoggingPropertiesResponsePtr(v *LoggingPropertiesResponseArgs) LoggingPropertiesResponsePtrInput {
	return (*loggingPropertiesResponsePtrType)(v)
}

func (*loggingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingPropertiesResponse)(nil)).Elem()
}

func (i *loggingPropertiesResponsePtrType) ToLoggingPropertiesResponsePtrOutput() LoggingPropertiesResponsePtrOutput {
	return i.ToLoggingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *loggingPropertiesResponsePtrType) ToLoggingPropertiesResponsePtrOutputWithContext(ctx context.Context) LoggingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoggingPropertiesResponsePtrOutput)
}

type LoggingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoggingPropertiesResponse)(nil)).Elem()
}

func (o LoggingPropertiesResponseOutput) ToLoggingPropertiesResponseOutput() LoggingPropertiesResponseOutput {
	return o
}

func (o LoggingPropertiesResponseOutput) ToLoggingPropertiesResponseOutputWithContext(ctx context.Context) LoggingPropertiesResponseOutput {
	return o
}

func (o LoggingPropertiesResponseOutput) ToLoggingPropertiesResponsePtrOutput() LoggingPropertiesResponsePtrOutput {
	return o.ToLoggingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o LoggingPropertiesResponseOutput) ToLoggingPropertiesResponsePtrOutputWithContext(ctx context.Context) LoggingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoggingPropertiesResponse) *LoggingPropertiesResponse {
		return &v
	}).(LoggingPropertiesResponsePtrOutput)
}

func (o LoggingPropertiesResponseOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingPropertiesResponse) *string { return v.AuditLogStatus }).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesResponseOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoggingPropertiesResponse) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

type LoggingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoggingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoggingPropertiesResponse)(nil)).Elem()
}

func (o LoggingPropertiesResponsePtrOutput) ToLoggingPropertiesResponsePtrOutput() LoggingPropertiesResponsePtrOutput {
	return o
}

func (o LoggingPropertiesResponsePtrOutput) ToLoggingPropertiesResponsePtrOutputWithContext(ctx context.Context) LoggingPropertiesResponsePtrOutput {
	return o
}

func (o LoggingPropertiesResponsePtrOutput) Elem() LoggingPropertiesResponseOutput {
	return o.ApplyT(func(v *LoggingPropertiesResponse) LoggingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LoggingPropertiesResponse
		return ret
	}).(LoggingPropertiesResponseOutput)
}

func (o LoggingPropertiesResponsePtrOutput) AuditLogStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuditLogStatus
	}).(pulumi.StringPtrOutput)
}

func (o LoggingPropertiesResponsePtrOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoggingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LogLevel
	}).(pulumi.StringPtrOutput)
}

type LoginServerPropertiesResponse struct {
	Host string                `pulumi:"host"`
	Tls  TlsPropertiesResponse `pulumi:"tls"`
}





type LoginServerPropertiesResponseInput interface {
	pulumi.Input

	ToLoginServerPropertiesResponseOutput() LoginServerPropertiesResponseOutput
	ToLoginServerPropertiesResponseOutputWithContext(context.Context) LoginServerPropertiesResponseOutput
}

type LoginServerPropertiesResponseArgs struct {
	Host pulumi.StringInput         `pulumi:"host"`
	Tls  TlsPropertiesResponseInput `pulumi:"tls"`
}

func (LoginServerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginServerPropertiesResponse)(nil)).Elem()
}

func (i LoginServerPropertiesResponseArgs) ToLoginServerPropertiesResponseOutput() LoginServerPropertiesResponseOutput {
	return i.ToLoginServerPropertiesResponseOutputWithContext(context.Background())
}

func (i LoginServerPropertiesResponseArgs) ToLoginServerPropertiesResponseOutputWithContext(ctx context.Context) LoginServerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginServerPropertiesResponseOutput)
}

func (i LoginServerPropertiesResponseArgs) ToLoginServerPropertiesResponsePtrOutput() LoginServerPropertiesResponsePtrOutput {
	return i.ToLoginServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i LoginServerPropertiesResponseArgs) ToLoginServerPropertiesResponsePtrOutputWithContext(ctx context.Context) LoginServerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginServerPropertiesResponseOutput).ToLoginServerPropertiesResponsePtrOutputWithContext(ctx)
}









type LoginServerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToLoginServerPropertiesResponsePtrOutput() LoginServerPropertiesResponsePtrOutput
	ToLoginServerPropertiesResponsePtrOutputWithContext(context.Context) LoginServerPropertiesResponsePtrOutput
}

type loginServerPropertiesResponsePtrType LoginServerPropertiesResponseArgs

func LoginServerPropertiesResponsePtr(v *LoginServerPropertiesResponseArgs) LoginServerPropertiesResponsePtrInput {
	return (*loginServerPropertiesResponsePtrType)(v)
}

func (*loginServerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginServerPropertiesResponse)(nil)).Elem()
}

func (i *loginServerPropertiesResponsePtrType) ToLoginServerPropertiesResponsePtrOutput() LoginServerPropertiesResponsePtrOutput {
	return i.ToLoginServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *loginServerPropertiesResponsePtrType) ToLoginServerPropertiesResponsePtrOutputWithContext(ctx context.Context) LoginServerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoginServerPropertiesResponsePtrOutput)
}

type LoginServerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LoginServerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoginServerPropertiesResponse)(nil)).Elem()
}

func (o LoginServerPropertiesResponseOutput) ToLoginServerPropertiesResponseOutput() LoginServerPropertiesResponseOutput {
	return o
}

func (o LoginServerPropertiesResponseOutput) ToLoginServerPropertiesResponseOutputWithContext(ctx context.Context) LoginServerPropertiesResponseOutput {
	return o
}

func (o LoginServerPropertiesResponseOutput) ToLoginServerPropertiesResponsePtrOutput() LoginServerPropertiesResponsePtrOutput {
	return o.ToLoginServerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o LoginServerPropertiesResponseOutput) ToLoginServerPropertiesResponsePtrOutputWithContext(ctx context.Context) LoginServerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LoginServerPropertiesResponse) *LoginServerPropertiesResponse {
		return &v
	}).(LoginServerPropertiesResponsePtrOutput)
}

func (o LoginServerPropertiesResponseOutput) Host() pulumi.StringOutput {
	return o.ApplyT(func(v LoginServerPropertiesResponse) string { return v.Host }).(pulumi.StringOutput)
}

func (o LoginServerPropertiesResponseOutput) Tls() TlsPropertiesResponseOutput {
	return o.ApplyT(func(v LoginServerPropertiesResponse) TlsPropertiesResponse { return v.Tls }).(TlsPropertiesResponseOutput)
}

type LoginServerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LoginServerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LoginServerPropertiesResponse)(nil)).Elem()
}

func (o LoginServerPropertiesResponsePtrOutput) ToLoginServerPropertiesResponsePtrOutput() LoginServerPropertiesResponsePtrOutput {
	return o
}

func (o LoginServerPropertiesResponsePtrOutput) ToLoginServerPropertiesResponsePtrOutputWithContext(ctx context.Context) LoginServerPropertiesResponsePtrOutput {
	return o
}

func (o LoginServerPropertiesResponsePtrOutput) Elem() LoginServerPropertiesResponseOutput {
	return o.ApplyT(func(v *LoginServerPropertiesResponse) LoginServerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LoginServerPropertiesResponse
		return ret
	}).(LoginServerPropertiesResponseOutput)
}

func (o LoginServerPropertiesResponsePtrOutput) Host() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LoginServerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Host
	}).(pulumi.StringPtrOutput)
}

func (o LoginServerPropertiesResponsePtrOutput) Tls() TlsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *LoginServerPropertiesResponse) *TlsPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.Tls
	}).(TlsPropertiesResponsePtrOutput)
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

type ParentProperties struct {
	Id             *string        `pulumi:"id"`
	SyncProperties SyncProperties `pulumi:"syncProperties"`
}





type ParentPropertiesInput interface {
	pulumi.Input

	ToParentPropertiesOutput() ParentPropertiesOutput
	ToParentPropertiesOutputWithContext(context.Context) ParentPropertiesOutput
}

type ParentPropertiesArgs struct {
	Id             pulumi.StringPtrInput `pulumi:"id"`
	SyncProperties SyncPropertiesInput   `pulumi:"syncProperties"`
}

func (ParentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentProperties)(nil)).Elem()
}

func (i ParentPropertiesArgs) ToParentPropertiesOutput() ParentPropertiesOutput {
	return i.ToParentPropertiesOutputWithContext(context.Background())
}

func (i ParentPropertiesArgs) ToParentPropertiesOutputWithContext(ctx context.Context) ParentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesOutput)
}

func (i ParentPropertiesArgs) ToParentPropertiesPtrOutput() ParentPropertiesPtrOutput {
	return i.ToParentPropertiesPtrOutputWithContext(context.Background())
}

func (i ParentPropertiesArgs) ToParentPropertiesPtrOutputWithContext(ctx context.Context) ParentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesOutput).ToParentPropertiesPtrOutputWithContext(ctx)
}









type ParentPropertiesPtrInput interface {
	pulumi.Input

	ToParentPropertiesPtrOutput() ParentPropertiesPtrOutput
	ToParentPropertiesPtrOutputWithContext(context.Context) ParentPropertiesPtrOutput
}

type parentPropertiesPtrType ParentPropertiesArgs

func ParentPropertiesPtr(v *ParentPropertiesArgs) ParentPropertiesPtrInput {
	return (*parentPropertiesPtrType)(v)
}

func (*parentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentProperties)(nil)).Elem()
}

func (i *parentPropertiesPtrType) ToParentPropertiesPtrOutput() ParentPropertiesPtrOutput {
	return i.ToParentPropertiesPtrOutputWithContext(context.Background())
}

func (i *parentPropertiesPtrType) ToParentPropertiesPtrOutputWithContext(ctx context.Context) ParentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesPtrOutput)
}

type ParentPropertiesOutput struct{ *pulumi.OutputState }

func (ParentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentProperties)(nil)).Elem()
}

func (o ParentPropertiesOutput) ToParentPropertiesOutput() ParentPropertiesOutput {
	return o
}

func (o ParentPropertiesOutput) ToParentPropertiesOutputWithContext(ctx context.Context) ParentPropertiesOutput {
	return o
}

func (o ParentPropertiesOutput) ToParentPropertiesPtrOutput() ParentPropertiesPtrOutput {
	return o.ToParentPropertiesPtrOutputWithContext(context.Background())
}

func (o ParentPropertiesOutput) ToParentPropertiesPtrOutputWithContext(ctx context.Context) ParentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParentProperties) *ParentProperties {
		return &v
	}).(ParentPropertiesPtrOutput)
}

func (o ParentPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParentPropertiesOutput) SyncProperties() SyncPropertiesOutput {
	return o.ApplyT(func(v ParentProperties) SyncProperties { return v.SyncProperties }).(SyncPropertiesOutput)
}

type ParentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ParentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentProperties)(nil)).Elem()
}

func (o ParentPropertiesPtrOutput) ToParentPropertiesPtrOutput() ParentPropertiesPtrOutput {
	return o
}

func (o ParentPropertiesPtrOutput) ToParentPropertiesPtrOutputWithContext(ctx context.Context) ParentPropertiesPtrOutput {
	return o
}

func (o ParentPropertiesPtrOutput) Elem() ParentPropertiesOutput {
	return o.ApplyT(func(v *ParentProperties) ParentProperties {
		if v != nil {
			return *v
		}
		var ret ParentProperties
		return ret
	}).(ParentPropertiesOutput)
}

func (o ParentPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ParentPropertiesPtrOutput) SyncProperties() SyncPropertiesPtrOutput {
	return o.ApplyT(func(v *ParentProperties) *SyncProperties {
		if v == nil {
			return nil
		}
		return &v.SyncProperties
	}).(SyncPropertiesPtrOutput)
}

type ParentPropertiesResponse struct {
	Id             *string                `pulumi:"id"`
	SyncProperties SyncPropertiesResponse `pulumi:"syncProperties"`
}





type ParentPropertiesResponseInput interface {
	pulumi.Input

	ToParentPropertiesResponseOutput() ParentPropertiesResponseOutput
	ToParentPropertiesResponseOutputWithContext(context.Context) ParentPropertiesResponseOutput
}

type ParentPropertiesResponseArgs struct {
	Id             pulumi.StringPtrInput       `pulumi:"id"`
	SyncProperties SyncPropertiesResponseInput `pulumi:"syncProperties"`
}

func (ParentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentPropertiesResponse)(nil)).Elem()
}

func (i ParentPropertiesResponseArgs) ToParentPropertiesResponseOutput() ParentPropertiesResponseOutput {
	return i.ToParentPropertiesResponseOutputWithContext(context.Background())
}

func (i ParentPropertiesResponseArgs) ToParentPropertiesResponseOutputWithContext(ctx context.Context) ParentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesResponseOutput)
}

func (i ParentPropertiesResponseArgs) ToParentPropertiesResponsePtrOutput() ParentPropertiesResponsePtrOutput {
	return i.ToParentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ParentPropertiesResponseArgs) ToParentPropertiesResponsePtrOutputWithContext(ctx context.Context) ParentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesResponseOutput).ToParentPropertiesResponsePtrOutputWithContext(ctx)
}









type ParentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToParentPropertiesResponsePtrOutput() ParentPropertiesResponsePtrOutput
	ToParentPropertiesResponsePtrOutputWithContext(context.Context) ParentPropertiesResponsePtrOutput
}

type parentPropertiesResponsePtrType ParentPropertiesResponseArgs

func ParentPropertiesResponsePtr(v *ParentPropertiesResponseArgs) ParentPropertiesResponsePtrInput {
	return (*parentPropertiesResponsePtrType)(v)
}

func (*parentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentPropertiesResponse)(nil)).Elem()
}

func (i *parentPropertiesResponsePtrType) ToParentPropertiesResponsePtrOutput() ParentPropertiesResponsePtrOutput {
	return i.ToParentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *parentPropertiesResponsePtrType) ToParentPropertiesResponsePtrOutputWithContext(ctx context.Context) ParentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ParentPropertiesResponsePtrOutput)
}

type ParentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ParentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParentPropertiesResponse)(nil)).Elem()
}

func (o ParentPropertiesResponseOutput) ToParentPropertiesResponseOutput() ParentPropertiesResponseOutput {
	return o
}

func (o ParentPropertiesResponseOutput) ToParentPropertiesResponseOutputWithContext(ctx context.Context) ParentPropertiesResponseOutput {
	return o
}

func (o ParentPropertiesResponseOutput) ToParentPropertiesResponsePtrOutput() ParentPropertiesResponsePtrOutput {
	return o.ToParentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ParentPropertiesResponseOutput) ToParentPropertiesResponsePtrOutputWithContext(ctx context.Context) ParentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParentPropertiesResponse) *ParentPropertiesResponse {
		return &v
	}).(ParentPropertiesResponsePtrOutput)
}

func (o ParentPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ParentPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ParentPropertiesResponseOutput) SyncProperties() SyncPropertiesResponseOutput {
	return o.ApplyT(func(v ParentPropertiesResponse) SyncPropertiesResponse { return v.SyncProperties }).(SyncPropertiesResponseOutput)
}

type ParentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ParentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParentPropertiesResponse)(nil)).Elem()
}

func (o ParentPropertiesResponsePtrOutput) ToParentPropertiesResponsePtrOutput() ParentPropertiesResponsePtrOutput {
	return o
}

func (o ParentPropertiesResponsePtrOutput) ToParentPropertiesResponsePtrOutputWithContext(ctx context.Context) ParentPropertiesResponsePtrOutput {
	return o
}

func (o ParentPropertiesResponsePtrOutput) Elem() ParentPropertiesResponseOutput {
	return o.ApplyT(func(v *ParentPropertiesResponse) ParentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ParentPropertiesResponse
		return ret
	}).(ParentPropertiesResponseOutput)
}

func (o ParentPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ParentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ParentPropertiesResponsePtrOutput) SyncProperties() SyncPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *ParentPropertiesResponse) *SyncPropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.SyncProperties
	}).(SyncPropertiesResponsePtrOutput)
}

type PipelineRunRequest struct {
	Artifacts          []string                     `pulumi:"artifacts"`
	CatalogDigest      *string                      `pulumi:"catalogDigest"`
	PipelineResourceId *string                      `pulumi:"pipelineResourceId"`
	Source             *PipelineRunSourceProperties `pulumi:"source"`
	Target             *PipelineRunTargetProperties `pulumi:"target"`
}


func (val *PipelineRunRequest) Defaults() *PipelineRunRequest {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = tmp.Source.Defaults()

	tmp.Target = tmp.Target.Defaults()

	return &tmp
}





type PipelineRunRequestInput interface {
	pulumi.Input

	ToPipelineRunRequestOutput() PipelineRunRequestOutput
	ToPipelineRunRequestOutputWithContext(context.Context) PipelineRunRequestOutput
}

type PipelineRunRequestArgs struct {
	Artifacts          pulumi.StringArrayInput             `pulumi:"artifacts"`
	CatalogDigest      pulumi.StringPtrInput               `pulumi:"catalogDigest"`
	PipelineResourceId pulumi.StringPtrInput               `pulumi:"pipelineResourceId"`
	Source             PipelineRunSourcePropertiesPtrInput `pulumi:"source"`
	Target             PipelineRunTargetPropertiesPtrInput `pulumi:"target"`
}

func (PipelineRunRequestArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequest)(nil)).Elem()
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestOutput() PipelineRunRequestOutput {
	return i.ToPipelineRunRequestOutputWithContext(context.Background())
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestOutputWithContext(ctx context.Context) PipelineRunRequestOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestOutput)
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return i.ToPipelineRunRequestPtrOutputWithContext(context.Background())
}

func (i PipelineRunRequestArgs) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestOutput).ToPipelineRunRequestPtrOutputWithContext(ctx)
}









type PipelineRunRequestPtrInput interface {
	pulumi.Input

	ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput
	ToPipelineRunRequestPtrOutputWithContext(context.Context) PipelineRunRequestPtrOutput
}

type pipelineRunRequestPtrType PipelineRunRequestArgs

func PipelineRunRequestPtr(v *PipelineRunRequestArgs) PipelineRunRequestPtrInput {
	return (*pipelineRunRequestPtrType)(v)
}

func (*pipelineRunRequestPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequest)(nil)).Elem()
}

func (i *pipelineRunRequestPtrType) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return i.ToPipelineRunRequestPtrOutputWithContext(context.Background())
}

func (i *pipelineRunRequestPtrType) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestPtrOutput)
}

type PipelineRunRequestOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequest)(nil)).Elem()
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestOutput() PipelineRunRequestOutput {
	return o
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestOutputWithContext(ctx context.Context) PipelineRunRequestOutput {
	return o
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return o.ToPipelineRunRequestPtrOutputWithContext(context.Background())
}

func (o PipelineRunRequestOutput) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunRequest) *PipelineRunRequest {
		return &v
	}).(PipelineRunRequestPtrOutput)
}

func (o PipelineRunRequestOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PipelineRunRequest) []string { return v.Artifacts }).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *string { return v.CatalogDigest }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *string { return v.PipelineResourceId }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestOutput) Source() PipelineRunSourcePropertiesPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *PipelineRunSourceProperties { return v.Source }).(PipelineRunSourcePropertiesPtrOutput)
}

func (o PipelineRunRequestOutput) Target() PipelineRunTargetPropertiesPtrOutput {
	return o.ApplyT(func(v PipelineRunRequest) *PipelineRunTargetProperties { return v.Target }).(PipelineRunTargetPropertiesPtrOutput)
}

type PipelineRunRequestPtrOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequest)(nil)).Elem()
}

func (o PipelineRunRequestPtrOutput) ToPipelineRunRequestPtrOutput() PipelineRunRequestPtrOutput {
	return o
}

func (o PipelineRunRequestPtrOutput) ToPipelineRunRequestPtrOutputWithContext(ctx context.Context) PipelineRunRequestPtrOutput {
	return o
}

func (o PipelineRunRequestPtrOutput) Elem() PipelineRunRequestOutput {
	return o.ApplyT(func(v *PipelineRunRequest) PipelineRunRequest {
		if v != nil {
			return *v
		}
		var ret PipelineRunRequest
		return ret
	}).(PipelineRunRequestOutput)
}

func (o PipelineRunRequestPtrOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PipelineRunRequest) []string {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestPtrOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *string {
		if v == nil {
			return nil
		}
		return v.CatalogDigest
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestPtrOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *string {
		if v == nil {
			return nil
		}
		return v.PipelineResourceId
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestPtrOutput) Source() PipelineRunSourcePropertiesPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *PipelineRunSourceProperties {
		if v == nil {
			return nil
		}
		return v.Source
	}).(PipelineRunSourcePropertiesPtrOutput)
}

func (o PipelineRunRequestPtrOutput) Target() PipelineRunTargetPropertiesPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequest) *PipelineRunTargetProperties {
		if v == nil {
			return nil
		}
		return v.Target
	}).(PipelineRunTargetPropertiesPtrOutput)
}

type PipelineRunRequestResponse struct {
	Artifacts          []string                             `pulumi:"artifacts"`
	CatalogDigest      *string                              `pulumi:"catalogDigest"`
	PipelineResourceId *string                              `pulumi:"pipelineResourceId"`
	Source             *PipelineRunSourcePropertiesResponse `pulumi:"source"`
	Target             *PipelineRunTargetPropertiesResponse `pulumi:"target"`
}


func (val *PipelineRunRequestResponse) Defaults() *PipelineRunRequestResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = tmp.Source.Defaults()

	tmp.Target = tmp.Target.Defaults()

	return &tmp
}





type PipelineRunRequestResponseInput interface {
	pulumi.Input

	ToPipelineRunRequestResponseOutput() PipelineRunRequestResponseOutput
	ToPipelineRunRequestResponseOutputWithContext(context.Context) PipelineRunRequestResponseOutput
}

type PipelineRunRequestResponseArgs struct {
	Artifacts          pulumi.StringArrayInput                     `pulumi:"artifacts"`
	CatalogDigest      pulumi.StringPtrInput                       `pulumi:"catalogDigest"`
	PipelineResourceId pulumi.StringPtrInput                       `pulumi:"pipelineResourceId"`
	Source             PipelineRunSourcePropertiesResponsePtrInput `pulumi:"source"`
	Target             PipelineRunTargetPropertiesResponsePtrInput `pulumi:"target"`
}

func (PipelineRunRequestResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequestResponse)(nil)).Elem()
}

func (i PipelineRunRequestResponseArgs) ToPipelineRunRequestResponseOutput() PipelineRunRequestResponseOutput {
	return i.ToPipelineRunRequestResponseOutputWithContext(context.Background())
}

func (i PipelineRunRequestResponseArgs) ToPipelineRunRequestResponseOutputWithContext(ctx context.Context) PipelineRunRequestResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestResponseOutput)
}

func (i PipelineRunRequestResponseArgs) ToPipelineRunRequestResponsePtrOutput() PipelineRunRequestResponsePtrOutput {
	return i.ToPipelineRunRequestResponsePtrOutputWithContext(context.Background())
}

func (i PipelineRunRequestResponseArgs) ToPipelineRunRequestResponsePtrOutputWithContext(ctx context.Context) PipelineRunRequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestResponseOutput).ToPipelineRunRequestResponsePtrOutputWithContext(ctx)
}









type PipelineRunRequestResponsePtrInput interface {
	pulumi.Input

	ToPipelineRunRequestResponsePtrOutput() PipelineRunRequestResponsePtrOutput
	ToPipelineRunRequestResponsePtrOutputWithContext(context.Context) PipelineRunRequestResponsePtrOutput
}

type pipelineRunRequestResponsePtrType PipelineRunRequestResponseArgs

func PipelineRunRequestResponsePtr(v *PipelineRunRequestResponseArgs) PipelineRunRequestResponsePtrInput {
	return (*pipelineRunRequestResponsePtrType)(v)
}

func (*pipelineRunRequestResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequestResponse)(nil)).Elem()
}

func (i *pipelineRunRequestResponsePtrType) ToPipelineRunRequestResponsePtrOutput() PipelineRunRequestResponsePtrOutput {
	return i.ToPipelineRunRequestResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineRunRequestResponsePtrType) ToPipelineRunRequestResponsePtrOutputWithContext(ctx context.Context) PipelineRunRequestResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunRequestResponsePtrOutput)
}

type PipelineRunRequestResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunRequestResponse)(nil)).Elem()
}

func (o PipelineRunRequestResponseOutput) ToPipelineRunRequestResponseOutput() PipelineRunRequestResponseOutput {
	return o
}

func (o PipelineRunRequestResponseOutput) ToPipelineRunRequestResponseOutputWithContext(ctx context.Context) PipelineRunRequestResponseOutput {
	return o
}

func (o PipelineRunRequestResponseOutput) ToPipelineRunRequestResponsePtrOutput() PipelineRunRequestResponsePtrOutput {
	return o.ToPipelineRunRequestResponsePtrOutputWithContext(context.Background())
}

func (o PipelineRunRequestResponseOutput) ToPipelineRunRequestResponsePtrOutputWithContext(ctx context.Context) PipelineRunRequestResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunRequestResponse) *PipelineRunRequestResponse {
		return &v
	}).(PipelineRunRequestResponsePtrOutput)
}

func (o PipelineRunRequestResponseOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) []string { return v.Artifacts }).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestResponseOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *string { return v.CatalogDigest }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponseOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *string { return v.PipelineResourceId }).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponseOutput) Source() PipelineRunSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *PipelineRunSourcePropertiesResponse { return v.Source }).(PipelineRunSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunRequestResponseOutput) Target() PipelineRunTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunRequestResponse) *PipelineRunTargetPropertiesResponse { return v.Target }).(PipelineRunTargetPropertiesResponsePtrOutput)
}

type PipelineRunRequestResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunRequestResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunRequestResponse)(nil)).Elem()
}

func (o PipelineRunRequestResponsePtrOutput) ToPipelineRunRequestResponsePtrOutput() PipelineRunRequestResponsePtrOutput {
	return o
}

func (o PipelineRunRequestResponsePtrOutput) ToPipelineRunRequestResponsePtrOutputWithContext(ctx context.Context) PipelineRunRequestResponsePtrOutput {
	return o
}

func (o PipelineRunRequestResponsePtrOutput) Elem() PipelineRunRequestResponseOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) PipelineRunRequestResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunRequestResponse
		return ret
	}).(PipelineRunRequestResponseOutput)
}

func (o PipelineRunRequestResponsePtrOutput) Artifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) []string {
		if v == nil {
			return nil
		}
		return v.Artifacts
	}).(pulumi.StringArrayOutput)
}

func (o PipelineRunRequestResponsePtrOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.CatalogDigest
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponsePtrOutput) PipelineResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *string {
		if v == nil {
			return nil
		}
		return v.PipelineResourceId
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunRequestResponsePtrOutput) Source() PipelineRunSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *PipelineRunSourcePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Source
	}).(PipelineRunSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunRequestResponsePtrOutput) Target() PipelineRunTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunRequestResponse) *PipelineRunTargetPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Target
	}).(PipelineRunTargetPropertiesResponsePtrOutput)
}

type PipelineRunResponseResponse struct {
	CatalogDigest           *string                                 `pulumi:"catalogDigest"`
	FinishTime              *string                                 `pulumi:"finishTime"`
	ImportedArtifacts       []string                                `pulumi:"importedArtifacts"`
	PipelineRunErrorMessage *string                                 `pulumi:"pipelineRunErrorMessage"`
	Progress                *ProgressPropertiesResponse             `pulumi:"progress"`
	Source                  *ImportPipelineSourcePropertiesResponse `pulumi:"source"`
	StartTime               *string                                 `pulumi:"startTime"`
	Status                  *string                                 `pulumi:"status"`
	Target                  *ExportPipelineTargetPropertiesResponse `pulumi:"target"`
	Trigger                 *PipelineTriggerDescriptorResponse      `pulumi:"trigger"`
}


func (val *PipelineRunResponseResponse) Defaults() *PipelineRunResponseResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Source = tmp.Source.Defaults()

	return &tmp
}





type PipelineRunResponseResponseInput interface {
	pulumi.Input

	ToPipelineRunResponseResponseOutput() PipelineRunResponseResponseOutput
	ToPipelineRunResponseResponseOutputWithContext(context.Context) PipelineRunResponseResponseOutput
}

type PipelineRunResponseResponseArgs struct {
	CatalogDigest           pulumi.StringPtrInput                          `pulumi:"catalogDigest"`
	FinishTime              pulumi.StringPtrInput                          `pulumi:"finishTime"`
	ImportedArtifacts       pulumi.StringArrayInput                        `pulumi:"importedArtifacts"`
	PipelineRunErrorMessage pulumi.StringPtrInput                          `pulumi:"pipelineRunErrorMessage"`
	Progress                ProgressPropertiesResponsePtrInput             `pulumi:"progress"`
	Source                  ImportPipelineSourcePropertiesResponsePtrInput `pulumi:"source"`
	StartTime               pulumi.StringPtrInput                          `pulumi:"startTime"`
	Status                  pulumi.StringPtrInput                          `pulumi:"status"`
	Target                  ExportPipelineTargetPropertiesResponsePtrInput `pulumi:"target"`
	Trigger                 PipelineTriggerDescriptorResponsePtrInput      `pulumi:"trigger"`
}

func (PipelineRunResponseResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunResponseResponse)(nil)).Elem()
}

func (i PipelineRunResponseResponseArgs) ToPipelineRunResponseResponseOutput() PipelineRunResponseResponseOutput {
	return i.ToPipelineRunResponseResponseOutputWithContext(context.Background())
}

func (i PipelineRunResponseResponseArgs) ToPipelineRunResponseResponseOutputWithContext(ctx context.Context) PipelineRunResponseResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunResponseResponseOutput)
}

func (i PipelineRunResponseResponseArgs) ToPipelineRunResponseResponsePtrOutput() PipelineRunResponseResponsePtrOutput {
	return i.ToPipelineRunResponseResponsePtrOutputWithContext(context.Background())
}

func (i PipelineRunResponseResponseArgs) ToPipelineRunResponseResponsePtrOutputWithContext(ctx context.Context) PipelineRunResponseResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunResponseResponseOutput).ToPipelineRunResponseResponsePtrOutputWithContext(ctx)
}









type PipelineRunResponseResponsePtrInput interface {
	pulumi.Input

	ToPipelineRunResponseResponsePtrOutput() PipelineRunResponseResponsePtrOutput
	ToPipelineRunResponseResponsePtrOutputWithContext(context.Context) PipelineRunResponseResponsePtrOutput
}

type pipelineRunResponseResponsePtrType PipelineRunResponseResponseArgs

func PipelineRunResponseResponsePtr(v *PipelineRunResponseResponseArgs) PipelineRunResponseResponsePtrInput {
	return (*pipelineRunResponseResponsePtrType)(v)
}

func (*pipelineRunResponseResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunResponseResponse)(nil)).Elem()
}

func (i *pipelineRunResponseResponsePtrType) ToPipelineRunResponseResponsePtrOutput() PipelineRunResponseResponsePtrOutput {
	return i.ToPipelineRunResponseResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineRunResponseResponsePtrType) ToPipelineRunResponseResponsePtrOutputWithContext(ctx context.Context) PipelineRunResponseResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunResponseResponsePtrOutput)
}

type PipelineRunResponseResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunResponseResponse)(nil)).Elem()
}

func (o PipelineRunResponseResponseOutput) ToPipelineRunResponseResponseOutput() PipelineRunResponseResponseOutput {
	return o
}

func (o PipelineRunResponseResponseOutput) ToPipelineRunResponseResponseOutputWithContext(ctx context.Context) PipelineRunResponseResponseOutput {
	return o
}

func (o PipelineRunResponseResponseOutput) ToPipelineRunResponseResponsePtrOutput() PipelineRunResponseResponsePtrOutput {
	return o.ToPipelineRunResponseResponsePtrOutputWithContext(context.Background())
}

func (o PipelineRunResponseResponseOutput) ToPipelineRunResponseResponsePtrOutputWithContext(ctx context.Context) PipelineRunResponseResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunResponseResponse) *PipelineRunResponseResponse {
		return &v
	}).(PipelineRunResponseResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.CatalogDigest }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.FinishTime }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) ImportedArtifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) []string { return v.ImportedArtifacts }).(pulumi.StringArrayOutput)
}

func (o PipelineRunResponseResponseOutput) PipelineRunErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.PipelineRunErrorMessage }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) Progress() ProgressPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *ProgressPropertiesResponse { return v.Progress }).(ProgressPropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) Source() ImportPipelineSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *ImportPipelineSourcePropertiesResponse { return v.Source }).(ImportPipelineSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponseOutput) Target() ExportPipelineTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *ExportPipelineTargetPropertiesResponse { return v.Target }).(ExportPipelineTargetPropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponseOutput) Trigger() PipelineTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v PipelineRunResponseResponse) *PipelineTriggerDescriptorResponse { return v.Trigger }).(PipelineTriggerDescriptorResponsePtrOutput)
}

type PipelineRunResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunResponseResponse)(nil)).Elem()
}

func (o PipelineRunResponseResponsePtrOutput) ToPipelineRunResponseResponsePtrOutput() PipelineRunResponseResponsePtrOutput {
	return o
}

func (o PipelineRunResponseResponsePtrOutput) ToPipelineRunResponseResponsePtrOutputWithContext(ctx context.Context) PipelineRunResponseResponsePtrOutput {
	return o
}

func (o PipelineRunResponseResponsePtrOutput) Elem() PipelineRunResponseResponseOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) PipelineRunResponseResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunResponseResponse
		return ret
	}).(PipelineRunResponseResponseOutput)
}

func (o PipelineRunResponseResponsePtrOutput) CatalogDigest() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.CatalogDigest
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.FinishTime
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) ImportedArtifacts() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) []string {
		if v == nil {
			return nil
		}
		return v.ImportedArtifacts
	}).(pulumi.StringArrayOutput)
}

func (o PipelineRunResponseResponsePtrOutput) PipelineRunErrorMessage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.PipelineRunErrorMessage
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) Progress() ProgressPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *ProgressPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Progress
	}).(ProgressPropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) Source() ImportPipelineSourcePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *ImportPipelineSourcePropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Source
	}).(ImportPipelineSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) Target() ExportPipelineTargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *ExportPipelineTargetPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.Target
	}).(ExportPipelineTargetPropertiesResponsePtrOutput)
}

func (o PipelineRunResponseResponsePtrOutput) Trigger() PipelineTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v *PipelineRunResponseResponse) *PipelineTriggerDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.Trigger
	}).(PipelineTriggerDescriptorResponsePtrOutput)
}

type PipelineRunSourceProperties struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunSourceProperties) Defaults() *PipelineRunSourceProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}





type PipelineRunSourcePropertiesInput interface {
	pulumi.Input

	ToPipelineRunSourcePropertiesOutput() PipelineRunSourcePropertiesOutput
	ToPipelineRunSourcePropertiesOutputWithContext(context.Context) PipelineRunSourcePropertiesOutput
}

type PipelineRunSourcePropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PipelineRunSourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourceProperties)(nil)).Elem()
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesOutput() PipelineRunSourcePropertiesOutput {
	return i.ToPipelineRunSourcePropertiesOutputWithContext(context.Background())
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesOutput)
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return i.ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineRunSourcePropertiesArgs) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesOutput).ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx)
}









type PipelineRunSourcePropertiesPtrInput interface {
	pulumi.Input

	ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput
	ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Context) PipelineRunSourcePropertiesPtrOutput
}

type pipelineRunSourcePropertiesPtrType PipelineRunSourcePropertiesArgs

func PipelineRunSourcePropertiesPtr(v *PipelineRunSourcePropertiesArgs) PipelineRunSourcePropertiesPtrInput {
	return (*pipelineRunSourcePropertiesPtrType)(v)
}

func (*pipelineRunSourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourceProperties)(nil)).Elem()
}

func (i *pipelineRunSourcePropertiesPtrType) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return i.ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineRunSourcePropertiesPtrType) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesPtrOutput)
}

type PipelineRunSourcePropertiesOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourceProperties)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesOutput() PipelineRunSourcePropertiesOutput {
	return o
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesOutput {
	return o
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return o.ToPipelineRunSourcePropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineRunSourcePropertiesOutput) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunSourceProperties) *PipelineRunSourceProperties {
		return &v
	}).(PipelineRunSourcePropertiesPtrOutput)
}

func (o PipelineRunSourcePropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourceProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourceProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunSourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourceProperties)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesPtrOutput) ToPipelineRunSourcePropertiesPtrOutput() PipelineRunSourcePropertiesPtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesPtrOutput) ToPipelineRunSourcePropertiesPtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesPtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesPtrOutput) Elem() PipelineRunSourcePropertiesOutput {
	return o.ApplyT(func(v *PipelineRunSourceProperties) PipelineRunSourceProperties {
		if v != nil {
			return *v
		}
		var ret PipelineRunSourceProperties
		return ret
	}).(PipelineRunSourcePropertiesOutput)
}

func (o PipelineRunSourcePropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineRunSourcePropertiesResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunSourcePropertiesResponse) Defaults() *PipelineRunSourcePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}





type PipelineRunSourcePropertiesResponseInput interface {
	pulumi.Input

	ToPipelineRunSourcePropertiesResponseOutput() PipelineRunSourcePropertiesResponseOutput
	ToPipelineRunSourcePropertiesResponseOutputWithContext(context.Context) PipelineRunSourcePropertiesResponseOutput
}

type PipelineRunSourcePropertiesResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PipelineRunSourcePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourcePropertiesResponse)(nil)).Elem()
}

func (i PipelineRunSourcePropertiesResponseArgs) ToPipelineRunSourcePropertiesResponseOutput() PipelineRunSourcePropertiesResponseOutput {
	return i.ToPipelineRunSourcePropertiesResponseOutputWithContext(context.Background())
}

func (i PipelineRunSourcePropertiesResponseArgs) ToPipelineRunSourcePropertiesResponseOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesResponseOutput)
}

func (i PipelineRunSourcePropertiesResponseArgs) ToPipelineRunSourcePropertiesResponsePtrOutput() PipelineRunSourcePropertiesResponsePtrOutput {
	return i.ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PipelineRunSourcePropertiesResponseArgs) ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesResponseOutput).ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(ctx)
}









type PipelineRunSourcePropertiesResponsePtrInput interface {
	pulumi.Input

	ToPipelineRunSourcePropertiesResponsePtrOutput() PipelineRunSourcePropertiesResponsePtrOutput
	ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(context.Context) PipelineRunSourcePropertiesResponsePtrOutput
}

type pipelineRunSourcePropertiesResponsePtrType PipelineRunSourcePropertiesResponseArgs

func PipelineRunSourcePropertiesResponsePtr(v *PipelineRunSourcePropertiesResponseArgs) PipelineRunSourcePropertiesResponsePtrInput {
	return (*pipelineRunSourcePropertiesResponsePtrType)(v)
}

func (*pipelineRunSourcePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourcePropertiesResponse)(nil)).Elem()
}

func (i *pipelineRunSourcePropertiesResponsePtrType) ToPipelineRunSourcePropertiesResponsePtrOutput() PipelineRunSourcePropertiesResponsePtrOutput {
	return i.ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineRunSourcePropertiesResponsePtrType) ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunSourcePropertiesResponsePtrOutput)
}

type PipelineRunSourcePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunSourcePropertiesResponse)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesResponseOutput) ToPipelineRunSourcePropertiesResponseOutput() PipelineRunSourcePropertiesResponseOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponseOutput) ToPipelineRunSourcePropertiesResponseOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponseOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponseOutput) ToPipelineRunSourcePropertiesResponsePtrOutput() PipelineRunSourcePropertiesResponsePtrOutput {
	return o.ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PipelineRunSourcePropertiesResponseOutput) ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunSourcePropertiesResponse) *PipelineRunSourcePropertiesResponse {
		return &v
	}).(PipelineRunSourcePropertiesResponsePtrOutput)
}

func (o PipelineRunSourcePropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourcePropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunSourcePropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunSourcePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunSourcePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunSourcePropertiesResponse)(nil)).Elem()
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) ToPipelineRunSourcePropertiesResponsePtrOutput() PipelineRunSourcePropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) ToPipelineRunSourcePropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunSourcePropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) Elem() PipelineRunSourcePropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineRunSourcePropertiesResponse) PipelineRunSourcePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunSourcePropertiesResponse
		return ret
	}).(PipelineRunSourcePropertiesResponseOutput)
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunSourcePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunSourcePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineRunTargetProperties struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunTargetProperties) Defaults() *PipelineRunTargetProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}





type PipelineRunTargetPropertiesInput interface {
	pulumi.Input

	ToPipelineRunTargetPropertiesOutput() PipelineRunTargetPropertiesOutput
	ToPipelineRunTargetPropertiesOutputWithContext(context.Context) PipelineRunTargetPropertiesOutput
}

type PipelineRunTargetPropertiesArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PipelineRunTargetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetProperties)(nil)).Elem()
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesOutput() PipelineRunTargetPropertiesOutput {
	return i.ToPipelineRunTargetPropertiesOutputWithContext(context.Background())
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesOutput)
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return i.ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineRunTargetPropertiesArgs) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesOutput).ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx)
}









type PipelineRunTargetPropertiesPtrInput interface {
	pulumi.Input

	ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput
	ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Context) PipelineRunTargetPropertiesPtrOutput
}

type pipelineRunTargetPropertiesPtrType PipelineRunTargetPropertiesArgs

func PipelineRunTargetPropertiesPtr(v *PipelineRunTargetPropertiesArgs) PipelineRunTargetPropertiesPtrInput {
	return (*pipelineRunTargetPropertiesPtrType)(v)
}

func (*pipelineRunTargetPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetProperties)(nil)).Elem()
}

func (i *pipelineRunTargetPropertiesPtrType) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return i.ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineRunTargetPropertiesPtrType) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesPtrOutput)
}

type PipelineRunTargetPropertiesOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetProperties)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesOutput() PipelineRunTargetPropertiesOutput {
	return o
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesOutput {
	return o
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return o.ToPipelineRunTargetPropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineRunTargetPropertiesOutput) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunTargetProperties) *PipelineRunTargetProperties {
		return &v
	}).(PipelineRunTargetPropertiesPtrOutput)
}

func (o PipelineRunTargetPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetProperties) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunTargetPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetProperties)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesPtrOutput) ToPipelineRunTargetPropertiesPtrOutput() PipelineRunTargetPropertiesPtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesPtrOutput) ToPipelineRunTargetPropertiesPtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesPtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesPtrOutput) Elem() PipelineRunTargetPropertiesOutput {
	return o.ApplyT(func(v *PipelineRunTargetProperties) PipelineRunTargetProperties {
		if v != nil {
			return *v
		}
		var ret PipelineRunTargetProperties
		return ret
	}).(PipelineRunTargetPropertiesOutput)
}

func (o PipelineRunTargetPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetProperties) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineRunTargetPropertiesResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}


func (val *PipelineRunTargetPropertiesResponse) Defaults() *PipelineRunTargetPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Type) {
		type_ := "AzureStorageBlob"
		tmp.Type = &type_
	}
	return &tmp
}





type PipelineRunTargetPropertiesResponseInput interface {
	pulumi.Input

	ToPipelineRunTargetPropertiesResponseOutput() PipelineRunTargetPropertiesResponseOutput
	ToPipelineRunTargetPropertiesResponseOutputWithContext(context.Context) PipelineRunTargetPropertiesResponseOutput
}

type PipelineRunTargetPropertiesResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (PipelineRunTargetPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetPropertiesResponse)(nil)).Elem()
}

func (i PipelineRunTargetPropertiesResponseArgs) ToPipelineRunTargetPropertiesResponseOutput() PipelineRunTargetPropertiesResponseOutput {
	return i.ToPipelineRunTargetPropertiesResponseOutputWithContext(context.Background())
}

func (i PipelineRunTargetPropertiesResponseArgs) ToPipelineRunTargetPropertiesResponseOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesResponseOutput)
}

func (i PipelineRunTargetPropertiesResponseArgs) ToPipelineRunTargetPropertiesResponsePtrOutput() PipelineRunTargetPropertiesResponsePtrOutput {
	return i.ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PipelineRunTargetPropertiesResponseArgs) ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesResponseOutput).ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(ctx)
}









type PipelineRunTargetPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPipelineRunTargetPropertiesResponsePtrOutput() PipelineRunTargetPropertiesResponsePtrOutput
	ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(context.Context) PipelineRunTargetPropertiesResponsePtrOutput
}

type pipelineRunTargetPropertiesResponsePtrType PipelineRunTargetPropertiesResponseArgs

func PipelineRunTargetPropertiesResponsePtr(v *PipelineRunTargetPropertiesResponseArgs) PipelineRunTargetPropertiesResponsePtrInput {
	return (*pipelineRunTargetPropertiesResponsePtrType)(v)
}

func (*pipelineRunTargetPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetPropertiesResponse)(nil)).Elem()
}

func (i *pipelineRunTargetPropertiesResponsePtrType) ToPipelineRunTargetPropertiesResponsePtrOutput() PipelineRunTargetPropertiesResponsePtrOutput {
	return i.ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineRunTargetPropertiesResponsePtrType) ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineRunTargetPropertiesResponsePtrOutput)
}

type PipelineRunTargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineRunTargetPropertiesResponse)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesResponseOutput) ToPipelineRunTargetPropertiesResponseOutput() PipelineRunTargetPropertiesResponseOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponseOutput) ToPipelineRunTargetPropertiesResponseOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponseOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponseOutput) ToPipelineRunTargetPropertiesResponsePtrOutput() PipelineRunTargetPropertiesResponsePtrOutput {
	return o.ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PipelineRunTargetPropertiesResponseOutput) ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineRunTargetPropertiesResponse) *PipelineRunTargetPropertiesResponse {
		return &v
	}).(PipelineRunTargetPropertiesResponsePtrOutput)
}

func (o PipelineRunTargetPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineRunTargetPropertiesResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type PipelineRunTargetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineRunTargetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineRunTargetPropertiesResponse)(nil)).Elem()
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) ToPipelineRunTargetPropertiesResponsePtrOutput() PipelineRunTargetPropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) ToPipelineRunTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineRunTargetPropertiesResponsePtrOutput {
	return o
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) Elem() PipelineRunTargetPropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineRunTargetPropertiesResponse) PipelineRunTargetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineRunTargetPropertiesResponse
		return ret
	}).(PipelineRunTargetPropertiesResponseOutput)
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PipelineRunTargetPropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineRunTargetPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerDescriptorResponse struct {
	Timestamp *string `pulumi:"timestamp"`
}





type PipelineSourceTriggerDescriptorResponseInput interface {
	pulumi.Input

	ToPipelineSourceTriggerDescriptorResponseOutput() PipelineSourceTriggerDescriptorResponseOutput
	ToPipelineSourceTriggerDescriptorResponseOutputWithContext(context.Context) PipelineSourceTriggerDescriptorResponseOutput
}

type PipelineSourceTriggerDescriptorResponseArgs struct {
	Timestamp pulumi.StringPtrInput `pulumi:"timestamp"`
}

func (PipelineSourceTriggerDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerDescriptorResponse)(nil)).Elem()
}

func (i PipelineSourceTriggerDescriptorResponseArgs) ToPipelineSourceTriggerDescriptorResponseOutput() PipelineSourceTriggerDescriptorResponseOutput {
	return i.ToPipelineSourceTriggerDescriptorResponseOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerDescriptorResponseArgs) ToPipelineSourceTriggerDescriptorResponseOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerDescriptorResponseOutput)
}

func (i PipelineSourceTriggerDescriptorResponseArgs) ToPipelineSourceTriggerDescriptorResponsePtrOutput() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return i.ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerDescriptorResponseArgs) ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerDescriptorResponseOutput).ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(ctx)
}









type PipelineSourceTriggerDescriptorResponsePtrInput interface {
	pulumi.Input

	ToPipelineSourceTriggerDescriptorResponsePtrOutput() PipelineSourceTriggerDescriptorResponsePtrOutput
	ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(context.Context) PipelineSourceTriggerDescriptorResponsePtrOutput
}

type pipelineSourceTriggerDescriptorResponsePtrType PipelineSourceTriggerDescriptorResponseArgs

func PipelineSourceTriggerDescriptorResponsePtr(v *PipelineSourceTriggerDescriptorResponseArgs) PipelineSourceTriggerDescriptorResponsePtrInput {
	return (*pipelineSourceTriggerDescriptorResponsePtrType)(v)
}

func (*pipelineSourceTriggerDescriptorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerDescriptorResponse)(nil)).Elem()
}

func (i *pipelineSourceTriggerDescriptorResponsePtrType) ToPipelineSourceTriggerDescriptorResponsePtrOutput() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return i.ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineSourceTriggerDescriptorResponsePtrType) ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerDescriptorResponsePtrOutput)
}

type PipelineSourceTriggerDescriptorResponseOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerDescriptorResponseOutput) ToPipelineSourceTriggerDescriptorResponseOutput() PipelineSourceTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponseOutput) ToPipelineSourceTriggerDescriptorResponseOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponseOutput) ToPipelineSourceTriggerDescriptorResponsePtrOutput() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o.ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (o PipelineSourceTriggerDescriptorResponseOutput) ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineSourceTriggerDescriptorResponse) *PipelineSourceTriggerDescriptorResponse {
		return &v
	}).(PipelineSourceTriggerDescriptorResponsePtrOutput)
}

func (o PipelineSourceTriggerDescriptorResponseOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PipelineSourceTriggerDescriptorResponse) *string { return v.Timestamp }).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) ToPipelineSourceTriggerDescriptorResponsePtrOutput() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) ToPipelineSourceTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) Elem() PipelineSourceTriggerDescriptorResponseOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerDescriptorResponse) PipelineSourceTriggerDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret PipelineSourceTriggerDescriptorResponse
		return ret
	}).(PipelineSourceTriggerDescriptorResponseOutput)
}

func (o PipelineSourceTriggerDescriptorResponsePtrOutput) Timestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerDescriptorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Timestamp
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerProperties struct {
	Status string `pulumi:"status"`
}


func (val *PipelineSourceTriggerProperties) Defaults() *PipelineSourceTriggerProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
}





type PipelineSourceTriggerPropertiesInput interface {
	pulumi.Input

	ToPipelineSourceTriggerPropertiesOutput() PipelineSourceTriggerPropertiesOutput
	ToPipelineSourceTriggerPropertiesOutputWithContext(context.Context) PipelineSourceTriggerPropertiesOutput
}

type PipelineSourceTriggerPropertiesArgs struct {
	Status pulumi.StringInput `pulumi:"status"`
}

func (PipelineSourceTriggerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerProperties)(nil)).Elem()
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesOutput() PipelineSourceTriggerPropertiesOutput {
	return i.ToPipelineSourceTriggerPropertiesOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesOutput)
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return i.ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerPropertiesArgs) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesOutput).ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx)
}









type PipelineSourceTriggerPropertiesPtrInput interface {
	pulumi.Input

	ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput
	ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Context) PipelineSourceTriggerPropertiesPtrOutput
}

type pipelineSourceTriggerPropertiesPtrType PipelineSourceTriggerPropertiesArgs

func PipelineSourceTriggerPropertiesPtr(v *PipelineSourceTriggerPropertiesArgs) PipelineSourceTriggerPropertiesPtrInput {
	return (*pipelineSourceTriggerPropertiesPtrType)(v)
}

func (*pipelineSourceTriggerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerProperties)(nil)).Elem()
}

func (i *pipelineSourceTriggerPropertiesPtrType) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return i.ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineSourceTriggerPropertiesPtrType) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesPtrOutput)
}

type PipelineSourceTriggerPropertiesOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerProperties)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesOutput() PipelineSourceTriggerPropertiesOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return o.ToPipelineSourceTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineSourceTriggerPropertiesOutput) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineSourceTriggerProperties) *PipelineSourceTriggerProperties {
		return &v
	}).(PipelineSourceTriggerPropertiesPtrOutput)
}

func (o PipelineSourceTriggerPropertiesOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineSourceTriggerProperties) string { return v.Status }).(pulumi.StringOutput)
}

type PipelineSourceTriggerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerProperties)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesPtrOutput) ToPipelineSourceTriggerPropertiesPtrOutput() PipelineSourceTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesPtrOutput) ToPipelineSourceTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesPtrOutput) Elem() PipelineSourceTriggerPropertiesOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerProperties) PipelineSourceTriggerProperties {
		if v != nil {
			return *v
		}
		var ret PipelineSourceTriggerProperties
		return ret
	}).(PipelineSourceTriggerPropertiesOutput)
}

func (o PipelineSourceTriggerPropertiesPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PipelineSourceTriggerPropertiesResponse struct {
	Status string `pulumi:"status"`
}


func (val *PipelineSourceTriggerPropertiesResponse) Defaults() *PipelineSourceTriggerPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		tmp.Status = "Enabled"
	}
	return &tmp
}





type PipelineSourceTriggerPropertiesResponseInput interface {
	pulumi.Input

	ToPipelineSourceTriggerPropertiesResponseOutput() PipelineSourceTriggerPropertiesResponseOutput
	ToPipelineSourceTriggerPropertiesResponseOutputWithContext(context.Context) PipelineSourceTriggerPropertiesResponseOutput
}

type PipelineSourceTriggerPropertiesResponseArgs struct {
	Status pulumi.StringInput `pulumi:"status"`
}

func (PipelineSourceTriggerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerPropertiesResponse)(nil)).Elem()
}

func (i PipelineSourceTriggerPropertiesResponseArgs) ToPipelineSourceTriggerPropertiesResponseOutput() PipelineSourceTriggerPropertiesResponseOutput {
	return i.ToPipelineSourceTriggerPropertiesResponseOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerPropertiesResponseArgs) ToPipelineSourceTriggerPropertiesResponseOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesResponseOutput)
}

func (i PipelineSourceTriggerPropertiesResponseArgs) ToPipelineSourceTriggerPropertiesResponsePtrOutput() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return i.ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PipelineSourceTriggerPropertiesResponseArgs) ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesResponseOutput).ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(ctx)
}









type PipelineSourceTriggerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPipelineSourceTriggerPropertiesResponsePtrOutput() PipelineSourceTriggerPropertiesResponsePtrOutput
	ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(context.Context) PipelineSourceTriggerPropertiesResponsePtrOutput
}

type pipelineSourceTriggerPropertiesResponsePtrType PipelineSourceTriggerPropertiesResponseArgs

func PipelineSourceTriggerPropertiesResponsePtr(v *PipelineSourceTriggerPropertiesResponseArgs) PipelineSourceTriggerPropertiesResponsePtrInput {
	return (*pipelineSourceTriggerPropertiesResponsePtrType)(v)
}

func (*pipelineSourceTriggerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerPropertiesResponse)(nil)).Elem()
}

func (i *pipelineSourceTriggerPropertiesResponsePtrType) ToPipelineSourceTriggerPropertiesResponsePtrOutput() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return i.ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineSourceTriggerPropertiesResponsePtrType) ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineSourceTriggerPropertiesResponsePtrOutput)
}

type PipelineSourceTriggerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineSourceTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesResponseOutput) ToPipelineSourceTriggerPropertiesResponseOutput() PipelineSourceTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponseOutput) ToPipelineSourceTriggerPropertiesResponseOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponseOutput) ToPipelineSourceTriggerPropertiesResponsePtrOutput() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o.ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PipelineSourceTriggerPropertiesResponseOutput) ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineSourceTriggerPropertiesResponse) *PipelineSourceTriggerPropertiesResponse {
		return &v
	}).(PipelineSourceTriggerPropertiesResponsePtrOutput)
}

func (o PipelineSourceTriggerPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v PipelineSourceTriggerPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type PipelineSourceTriggerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineSourceTriggerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineSourceTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) ToPipelineSourceTriggerPropertiesResponsePtrOutput() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) ToPipelineSourceTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) Elem() PipelineSourceTriggerPropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerPropertiesResponse) PipelineSourceTriggerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineSourceTriggerPropertiesResponse
		return ret
	}).(PipelineSourceTriggerPropertiesResponseOutput)
}

func (o PipelineSourceTriggerPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PipelineSourceTriggerPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type PipelineTriggerDescriptorResponse struct {
	SourceTrigger *PipelineSourceTriggerDescriptorResponse `pulumi:"sourceTrigger"`
}





type PipelineTriggerDescriptorResponseInput interface {
	pulumi.Input

	ToPipelineTriggerDescriptorResponseOutput() PipelineTriggerDescriptorResponseOutput
	ToPipelineTriggerDescriptorResponseOutputWithContext(context.Context) PipelineTriggerDescriptorResponseOutput
}

type PipelineTriggerDescriptorResponseArgs struct {
	SourceTrigger PipelineSourceTriggerDescriptorResponsePtrInput `pulumi:"sourceTrigger"`
}

func (PipelineTriggerDescriptorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerDescriptorResponse)(nil)).Elem()
}

func (i PipelineTriggerDescriptorResponseArgs) ToPipelineTriggerDescriptorResponseOutput() PipelineTriggerDescriptorResponseOutput {
	return i.ToPipelineTriggerDescriptorResponseOutputWithContext(context.Background())
}

func (i PipelineTriggerDescriptorResponseArgs) ToPipelineTriggerDescriptorResponseOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerDescriptorResponseOutput)
}

func (i PipelineTriggerDescriptorResponseArgs) ToPipelineTriggerDescriptorResponsePtrOutput() PipelineTriggerDescriptorResponsePtrOutput {
	return i.ToPipelineTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i PipelineTriggerDescriptorResponseArgs) ToPipelineTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerDescriptorResponseOutput).ToPipelineTriggerDescriptorResponsePtrOutputWithContext(ctx)
}









type PipelineTriggerDescriptorResponsePtrInput interface {
	pulumi.Input

	ToPipelineTriggerDescriptorResponsePtrOutput() PipelineTriggerDescriptorResponsePtrOutput
	ToPipelineTriggerDescriptorResponsePtrOutputWithContext(context.Context) PipelineTriggerDescriptorResponsePtrOutput
}

type pipelineTriggerDescriptorResponsePtrType PipelineTriggerDescriptorResponseArgs

func PipelineTriggerDescriptorResponsePtr(v *PipelineTriggerDescriptorResponseArgs) PipelineTriggerDescriptorResponsePtrInput {
	return (*pipelineTriggerDescriptorResponsePtrType)(v)
}

func (*pipelineTriggerDescriptorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerDescriptorResponse)(nil)).Elem()
}

func (i *pipelineTriggerDescriptorResponsePtrType) ToPipelineTriggerDescriptorResponsePtrOutput() PipelineTriggerDescriptorResponsePtrOutput {
	return i.ToPipelineTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineTriggerDescriptorResponsePtrType) ToPipelineTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerDescriptorResponsePtrOutput)
}

type PipelineTriggerDescriptorResponseOutput struct{ *pulumi.OutputState }

func (PipelineTriggerDescriptorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineTriggerDescriptorResponseOutput) ToPipelineTriggerDescriptorResponseOutput() PipelineTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineTriggerDescriptorResponseOutput) ToPipelineTriggerDescriptorResponseOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponseOutput {
	return o
}

func (o PipelineTriggerDescriptorResponseOutput) ToPipelineTriggerDescriptorResponsePtrOutput() PipelineTriggerDescriptorResponsePtrOutput {
	return o.ToPipelineTriggerDescriptorResponsePtrOutputWithContext(context.Background())
}

func (o PipelineTriggerDescriptorResponseOutput) ToPipelineTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineTriggerDescriptorResponse) *PipelineTriggerDescriptorResponse {
		return &v
	}).(PipelineTriggerDescriptorResponsePtrOutput)
}

func (o PipelineTriggerDescriptorResponseOutput) SourceTrigger() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v PipelineTriggerDescriptorResponse) *PipelineSourceTriggerDescriptorResponse {
		return v.SourceTrigger
	}).(PipelineSourceTriggerDescriptorResponsePtrOutput)
}

type PipelineTriggerDescriptorResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineTriggerDescriptorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerDescriptorResponse)(nil)).Elem()
}

func (o PipelineTriggerDescriptorResponsePtrOutput) ToPipelineTriggerDescriptorResponsePtrOutput() PipelineTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineTriggerDescriptorResponsePtrOutput) ToPipelineTriggerDescriptorResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerDescriptorResponsePtrOutput {
	return o
}

func (o PipelineTriggerDescriptorResponsePtrOutput) Elem() PipelineTriggerDescriptorResponseOutput {
	return o.ApplyT(func(v *PipelineTriggerDescriptorResponse) PipelineTriggerDescriptorResponse {
		if v != nil {
			return *v
		}
		var ret PipelineTriggerDescriptorResponse
		return ret
	}).(PipelineTriggerDescriptorResponseOutput)
}

func (o PipelineTriggerDescriptorResponsePtrOutput) SourceTrigger() PipelineSourceTriggerDescriptorResponsePtrOutput {
	return o.ApplyT(func(v *PipelineTriggerDescriptorResponse) *PipelineSourceTriggerDescriptorResponse {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(PipelineSourceTriggerDescriptorResponsePtrOutput)
}

type PipelineTriggerProperties struct {
	SourceTrigger *PipelineSourceTriggerProperties `pulumi:"sourceTrigger"`
}


func (val *PipelineTriggerProperties) Defaults() *PipelineTriggerProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceTrigger = tmp.SourceTrigger.Defaults()

	return &tmp
}





type PipelineTriggerPropertiesInput interface {
	pulumi.Input

	ToPipelineTriggerPropertiesOutput() PipelineTriggerPropertiesOutput
	ToPipelineTriggerPropertiesOutputWithContext(context.Context) PipelineTriggerPropertiesOutput
}

type PipelineTriggerPropertiesArgs struct {
	SourceTrigger PipelineSourceTriggerPropertiesPtrInput `pulumi:"sourceTrigger"`
}

func (PipelineTriggerPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerProperties)(nil)).Elem()
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesOutput() PipelineTriggerPropertiesOutput {
	return i.ToPipelineTriggerPropertiesOutputWithContext(context.Background())
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesOutputWithContext(ctx context.Context) PipelineTriggerPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesOutput)
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return i.ToPipelineTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i PipelineTriggerPropertiesArgs) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesOutput).ToPipelineTriggerPropertiesPtrOutputWithContext(ctx)
}









type PipelineTriggerPropertiesPtrInput interface {
	pulumi.Input

	ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput
	ToPipelineTriggerPropertiesPtrOutputWithContext(context.Context) PipelineTriggerPropertiesPtrOutput
}

type pipelineTriggerPropertiesPtrType PipelineTriggerPropertiesArgs

func PipelineTriggerPropertiesPtr(v *PipelineTriggerPropertiesArgs) PipelineTriggerPropertiesPtrInput {
	return (*pipelineTriggerPropertiesPtrType)(v)
}

func (*pipelineTriggerPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerProperties)(nil)).Elem()
}

func (i *pipelineTriggerPropertiesPtrType) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return i.ToPipelineTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (i *pipelineTriggerPropertiesPtrType) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesPtrOutput)
}

type PipelineTriggerPropertiesOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerProperties)(nil)).Elem()
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesOutput() PipelineTriggerPropertiesOutput {
	return o
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesOutputWithContext(ctx context.Context) PipelineTriggerPropertiesOutput {
	return o
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return o.ToPipelineTriggerPropertiesPtrOutputWithContext(context.Background())
}

func (o PipelineTriggerPropertiesOutput) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineTriggerProperties) *PipelineTriggerProperties {
		return &v
	}).(PipelineTriggerPropertiesPtrOutput)
}

func (o PipelineTriggerPropertiesOutput) SourceTrigger() PipelineSourceTriggerPropertiesPtrOutput {
	return o.ApplyT(func(v PipelineTriggerProperties) *PipelineSourceTriggerProperties { return v.SourceTrigger }).(PipelineSourceTriggerPropertiesPtrOutput)
}

type PipelineTriggerPropertiesPtrOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerProperties)(nil)).Elem()
}

func (o PipelineTriggerPropertiesPtrOutput) ToPipelineTriggerPropertiesPtrOutput() PipelineTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineTriggerPropertiesPtrOutput) ToPipelineTriggerPropertiesPtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesPtrOutput {
	return o
}

func (o PipelineTriggerPropertiesPtrOutput) Elem() PipelineTriggerPropertiesOutput {
	return o.ApplyT(func(v *PipelineTriggerProperties) PipelineTriggerProperties {
		if v != nil {
			return *v
		}
		var ret PipelineTriggerProperties
		return ret
	}).(PipelineTriggerPropertiesOutput)
}

func (o PipelineTriggerPropertiesPtrOutput) SourceTrigger() PipelineSourceTriggerPropertiesPtrOutput {
	return o.ApplyT(func(v *PipelineTriggerProperties) *PipelineSourceTriggerProperties {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(PipelineSourceTriggerPropertiesPtrOutput)
}

type PipelineTriggerPropertiesResponse struct {
	SourceTrigger *PipelineSourceTriggerPropertiesResponse `pulumi:"sourceTrigger"`
}


func (val *PipelineTriggerPropertiesResponse) Defaults() *PipelineTriggerPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.SourceTrigger = tmp.SourceTrigger.Defaults()

	return &tmp
}





type PipelineTriggerPropertiesResponseInput interface {
	pulumi.Input

	ToPipelineTriggerPropertiesResponseOutput() PipelineTriggerPropertiesResponseOutput
	ToPipelineTriggerPropertiesResponseOutputWithContext(context.Context) PipelineTriggerPropertiesResponseOutput
}

type PipelineTriggerPropertiesResponseArgs struct {
	SourceTrigger PipelineSourceTriggerPropertiesResponsePtrInput `pulumi:"sourceTrigger"`
}

func (PipelineTriggerPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerPropertiesResponse)(nil)).Elem()
}

func (i PipelineTriggerPropertiesResponseArgs) ToPipelineTriggerPropertiesResponseOutput() PipelineTriggerPropertiesResponseOutput {
	return i.ToPipelineTriggerPropertiesResponseOutputWithContext(context.Background())
}

func (i PipelineTriggerPropertiesResponseArgs) ToPipelineTriggerPropertiesResponseOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesResponseOutput)
}

func (i PipelineTriggerPropertiesResponseArgs) ToPipelineTriggerPropertiesResponsePtrOutput() PipelineTriggerPropertiesResponsePtrOutput {
	return i.ToPipelineTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PipelineTriggerPropertiesResponseArgs) ToPipelineTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesResponseOutput).ToPipelineTriggerPropertiesResponsePtrOutputWithContext(ctx)
}









type PipelineTriggerPropertiesResponsePtrInput interface {
	pulumi.Input

	ToPipelineTriggerPropertiesResponsePtrOutput() PipelineTriggerPropertiesResponsePtrOutput
	ToPipelineTriggerPropertiesResponsePtrOutputWithContext(context.Context) PipelineTriggerPropertiesResponsePtrOutput
}

type pipelineTriggerPropertiesResponsePtrType PipelineTriggerPropertiesResponseArgs

func PipelineTriggerPropertiesResponsePtr(v *PipelineTriggerPropertiesResponseArgs) PipelineTriggerPropertiesResponsePtrInput {
	return (*pipelineTriggerPropertiesResponsePtrType)(v)
}

func (*pipelineTriggerPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerPropertiesResponse)(nil)).Elem()
}

func (i *pipelineTriggerPropertiesResponsePtrType) ToPipelineTriggerPropertiesResponsePtrOutput() PipelineTriggerPropertiesResponsePtrOutput {
	return i.ToPipelineTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *pipelineTriggerPropertiesResponsePtrType) ToPipelineTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PipelineTriggerPropertiesResponsePtrOutput)
}

type PipelineTriggerPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PipelineTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineTriggerPropertiesResponseOutput) ToPipelineTriggerPropertiesResponseOutput() PipelineTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineTriggerPropertiesResponseOutput) ToPipelineTriggerPropertiesResponseOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponseOutput {
	return o
}

func (o PipelineTriggerPropertiesResponseOutput) ToPipelineTriggerPropertiesResponsePtrOutput() PipelineTriggerPropertiesResponsePtrOutput {
	return o.ToPipelineTriggerPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PipelineTriggerPropertiesResponseOutput) ToPipelineTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PipelineTriggerPropertiesResponse) *PipelineTriggerPropertiesResponse {
		return &v
	}).(PipelineTriggerPropertiesResponsePtrOutput)
}

func (o PipelineTriggerPropertiesResponseOutput) SourceTrigger() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v PipelineTriggerPropertiesResponse) *PipelineSourceTriggerPropertiesResponse {
		return v.SourceTrigger
	}).(PipelineSourceTriggerPropertiesResponsePtrOutput)
}

type PipelineTriggerPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PipelineTriggerPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PipelineTriggerPropertiesResponse)(nil)).Elem()
}

func (o PipelineTriggerPropertiesResponsePtrOutput) ToPipelineTriggerPropertiesResponsePtrOutput() PipelineTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineTriggerPropertiesResponsePtrOutput) ToPipelineTriggerPropertiesResponsePtrOutputWithContext(ctx context.Context) PipelineTriggerPropertiesResponsePtrOutput {
	return o
}

func (o PipelineTriggerPropertiesResponsePtrOutput) Elem() PipelineTriggerPropertiesResponseOutput {
	return o.ApplyT(func(v *PipelineTriggerPropertiesResponse) PipelineTriggerPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PipelineTriggerPropertiesResponse
		return ret
	}).(PipelineTriggerPropertiesResponseOutput)
}

func (o PipelineTriggerPropertiesResponsePtrOutput) SourceTrigger() PipelineSourceTriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *PipelineTriggerPropertiesResponse) *PipelineSourceTriggerPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.SourceTrigger
	}).(PipelineSourceTriggerPropertiesResponsePtrOutput)
}

type Policies struct {
	QuarantinePolicy *QuarantinePolicy `pulumi:"quarantinePolicy"`
	RetentionPolicy  *RetentionPolicy  `pulumi:"retentionPolicy"`
	TrustPolicy      *TrustPolicy      `pulumi:"trustPolicy"`
}


func (val *Policies) Defaults() *Policies {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.QuarantinePolicy = tmp.QuarantinePolicy.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	tmp.TrustPolicy = tmp.TrustPolicy.Defaults()

	return &tmp
}





type PoliciesInput interface {
	pulumi.Input

	ToPoliciesOutput() PoliciesOutput
	ToPoliciesOutputWithContext(context.Context) PoliciesOutput
}

type PoliciesArgs struct {
	QuarantinePolicy QuarantinePolicyPtrInput `pulumi:"quarantinePolicy"`
	RetentionPolicy  RetentionPolicyPtrInput  `pulumi:"retentionPolicy"`
	TrustPolicy      TrustPolicyPtrInput      `pulumi:"trustPolicy"`
}

func (PoliciesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Policies)(nil)).Elem()
}

func (i PoliciesArgs) ToPoliciesOutput() PoliciesOutput {
	return i.ToPoliciesOutputWithContext(context.Background())
}

func (i PoliciesArgs) ToPoliciesOutputWithContext(ctx context.Context) PoliciesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesOutput)
}

func (i PoliciesArgs) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return i.ToPoliciesPtrOutputWithContext(context.Background())
}

func (i PoliciesArgs) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesOutput).ToPoliciesPtrOutputWithContext(ctx)
}









type PoliciesPtrInput interface {
	pulumi.Input

	ToPoliciesPtrOutput() PoliciesPtrOutput
	ToPoliciesPtrOutputWithContext(context.Context) PoliciesPtrOutput
}

type policiesPtrType PoliciesArgs

func PoliciesPtr(v *PoliciesArgs) PoliciesPtrInput {
	return (*policiesPtrType)(v)
}

func (*policiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Policies)(nil)).Elem()
}

func (i *policiesPtrType) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return i.ToPoliciesPtrOutputWithContext(context.Background())
}

func (i *policiesPtrType) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesPtrOutput)
}

type PoliciesOutput struct{ *pulumi.OutputState }

func (PoliciesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Policies)(nil)).Elem()
}

func (o PoliciesOutput) ToPoliciesOutput() PoliciesOutput {
	return o
}

func (o PoliciesOutput) ToPoliciesOutputWithContext(ctx context.Context) PoliciesOutput {
	return o
}

func (o PoliciesOutput) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return o.ToPoliciesPtrOutputWithContext(context.Background())
}

func (o PoliciesOutput) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Policies) *Policies {
		return &v
	}).(PoliciesPtrOutput)
}

func (o PoliciesOutput) QuarantinePolicy() QuarantinePolicyPtrOutput {
	return o.ApplyT(func(v Policies) *QuarantinePolicy { return v.QuarantinePolicy }).(QuarantinePolicyPtrOutput)
}

func (o PoliciesOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v Policies) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

func (o PoliciesOutput) TrustPolicy() TrustPolicyPtrOutput {
	return o.ApplyT(func(v Policies) *TrustPolicy { return v.TrustPolicy }).(TrustPolicyPtrOutput)
}

type PoliciesPtrOutput struct{ *pulumi.OutputState }

func (PoliciesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Policies)(nil)).Elem()
}

func (o PoliciesPtrOutput) ToPoliciesPtrOutput() PoliciesPtrOutput {
	return o
}

func (o PoliciesPtrOutput) ToPoliciesPtrOutputWithContext(ctx context.Context) PoliciesPtrOutput {
	return o
}

func (o PoliciesPtrOutput) Elem() PoliciesOutput {
	return o.ApplyT(func(v *Policies) Policies {
		if v != nil {
			return *v
		}
		var ret Policies
		return ret
	}).(PoliciesOutput)
}

func (o PoliciesPtrOutput) QuarantinePolicy() QuarantinePolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *QuarantinePolicy {
		if v == nil {
			return nil
		}
		return v.QuarantinePolicy
	}).(QuarantinePolicyPtrOutput)
}

func (o PoliciesPtrOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *RetentionPolicy {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(RetentionPolicyPtrOutput)
}

func (o PoliciesPtrOutput) TrustPolicy() TrustPolicyPtrOutput {
	return o.ApplyT(func(v *Policies) *TrustPolicy {
		if v == nil {
			return nil
		}
		return v.TrustPolicy
	}).(TrustPolicyPtrOutput)
}

type PoliciesResponse struct {
	QuarantinePolicy *QuarantinePolicyResponse `pulumi:"quarantinePolicy"`
	RetentionPolicy  *RetentionPolicyResponse  `pulumi:"retentionPolicy"`
	TrustPolicy      *TrustPolicyResponse      `pulumi:"trustPolicy"`
}


func (val *PoliciesResponse) Defaults() *PoliciesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.QuarantinePolicy = tmp.QuarantinePolicy.Defaults()

	tmp.RetentionPolicy = tmp.RetentionPolicy.Defaults()

	tmp.TrustPolicy = tmp.TrustPolicy.Defaults()

	return &tmp
}





type PoliciesResponseInput interface {
	pulumi.Input

	ToPoliciesResponseOutput() PoliciesResponseOutput
	ToPoliciesResponseOutputWithContext(context.Context) PoliciesResponseOutput
}

type PoliciesResponseArgs struct {
	QuarantinePolicy QuarantinePolicyResponsePtrInput `pulumi:"quarantinePolicy"`
	RetentionPolicy  RetentionPolicyResponsePtrInput  `pulumi:"retentionPolicy"`
	TrustPolicy      TrustPolicyResponsePtrInput      `pulumi:"trustPolicy"`
}

func (PoliciesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PoliciesResponse)(nil)).Elem()
}

func (i PoliciesResponseArgs) ToPoliciesResponseOutput() PoliciesResponseOutput {
	return i.ToPoliciesResponseOutputWithContext(context.Background())
}

func (i PoliciesResponseArgs) ToPoliciesResponseOutputWithContext(ctx context.Context) PoliciesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesResponseOutput)
}

func (i PoliciesResponseArgs) ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput {
	return i.ToPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i PoliciesResponseArgs) ToPoliciesResponsePtrOutputWithContext(ctx context.Context) PoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesResponseOutput).ToPoliciesResponsePtrOutputWithContext(ctx)
}









type PoliciesResponsePtrInput interface {
	pulumi.Input

	ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput
	ToPoliciesResponsePtrOutputWithContext(context.Context) PoliciesResponsePtrOutput
}

type policiesResponsePtrType PoliciesResponseArgs

func PoliciesResponsePtr(v *PoliciesResponseArgs) PoliciesResponsePtrInput {
	return (*policiesResponsePtrType)(v)
}

func (*policiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PoliciesResponse)(nil)).Elem()
}

func (i *policiesResponsePtrType) ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput {
	return i.ToPoliciesResponsePtrOutputWithContext(context.Background())
}

func (i *policiesResponsePtrType) ToPoliciesResponsePtrOutputWithContext(ctx context.Context) PoliciesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PoliciesResponsePtrOutput)
}

type PoliciesResponseOutput struct{ *pulumi.OutputState }

func (PoliciesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PoliciesResponse)(nil)).Elem()
}

func (o PoliciesResponseOutput) ToPoliciesResponseOutput() PoliciesResponseOutput {
	return o
}

func (o PoliciesResponseOutput) ToPoliciesResponseOutputWithContext(ctx context.Context) PoliciesResponseOutput {
	return o
}

func (o PoliciesResponseOutput) ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput {
	return o.ToPoliciesResponsePtrOutputWithContext(context.Background())
}

func (o PoliciesResponseOutput) ToPoliciesResponsePtrOutputWithContext(ctx context.Context) PoliciesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PoliciesResponse) *PoliciesResponse {
		return &v
	}).(PoliciesResponsePtrOutput)
}

func (o PoliciesResponseOutput) QuarantinePolicy() QuarantinePolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *QuarantinePolicyResponse { return v.QuarantinePolicy }).(QuarantinePolicyResponsePtrOutput)
}

func (o PoliciesResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

func (o PoliciesResponseOutput) TrustPolicy() TrustPolicyResponsePtrOutput {
	return o.ApplyT(func(v PoliciesResponse) *TrustPolicyResponse { return v.TrustPolicy }).(TrustPolicyResponsePtrOutput)
}

type PoliciesResponsePtrOutput struct{ *pulumi.OutputState }

func (PoliciesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PoliciesResponse)(nil)).Elem()
}

func (o PoliciesResponsePtrOutput) ToPoliciesResponsePtrOutput() PoliciesResponsePtrOutput {
	return o
}

func (o PoliciesResponsePtrOutput) ToPoliciesResponsePtrOutputWithContext(ctx context.Context) PoliciesResponsePtrOutput {
	return o
}

func (o PoliciesResponsePtrOutput) Elem() PoliciesResponseOutput {
	return o.ApplyT(func(v *PoliciesResponse) PoliciesResponse {
		if v != nil {
			return *v
		}
		var ret PoliciesResponse
		return ret
	}).(PoliciesResponseOutput)
}

func (o PoliciesResponsePtrOutput) QuarantinePolicy() QuarantinePolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *QuarantinePolicyResponse {
		if v == nil {
			return nil
		}
		return v.QuarantinePolicy
	}).(QuarantinePolicyResponsePtrOutput)
}

func (o PoliciesResponsePtrOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *RetentionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.RetentionPolicy
	}).(RetentionPolicyResponsePtrOutput)
}

func (o PoliciesResponsePtrOutput) TrustPolicy() TrustPolicyResponsePtrOutput {
	return o.ApplyT(func(v *PoliciesResponse) *TrustPolicyResponse {
		if v == nil {
			return nil
		}
		return v.TrustPolicy
	}).(TrustPolicyResponsePtrOutput)
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
	Id                                string                                     `pulumi:"id"`
	Name                              string                                     `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                     `pulumi:"provisioningState"`
	SystemData                        SystemDataResponse                         `pulumi:"systemData"`
	Type                              string                                     `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	Id                                pulumi.StringInput                                `pulumi:"id"`
	Name                              pulumi.StringInput                                `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStateResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                `pulumi:"provisioningState"`
	SystemData                        SystemDataResponseInput                           `pulumi:"systemData"`
	Type                              pulumi.StringInput                                `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
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

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
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





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput).ToPrivateEndpointResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput
	ToPrivateEndpointResponsePtrOutputWithContext(context.Context) PrivateEndpointResponsePtrOutput
}

type privateEndpointResponsePtrType PrivateEndpointResponseArgs

func PrivateEndpointResponsePtr(v *PrivateEndpointResponseArgs) PrivateEndpointResponsePtrInput {
	return (*privateEndpointResponsePtrType)(v)
}

func (*privateEndpointResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return i.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointResponsePtrType) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponsePtrOutput)
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

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o.ToPrivateEndpointResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointResponse) *PrivateEndpointResponse {
		return &v
	}).(PrivateEndpointResponsePtrOutput)
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

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput).ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput
	ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePtrOutput
}

type privateLinkServiceConnectionStatePtrType PrivateLinkServiceConnectionStateArgs

func PrivateLinkServiceConnectionStatePtr(v *PrivateLinkServiceConnectionStateArgs) PrivateLinkServiceConnectionStatePtrInput {
	return (*privateLinkServiceConnectionStatePtrType)(v)
}

func (*privateLinkServiceConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePtrType) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePtrOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionState) *PrivateLinkServiceConnectionState {
		return &v
	}).(PrivateLinkServiceConnectionStatePtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutput() PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ToPrivateLinkServiceConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Elem() PrivateLinkServiceConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) PrivateLinkServiceConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionState
		return ret
	}).(PrivateLinkServiceConnectionStateOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput
	ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponseOutput
}

type PrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return i.ToPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateResponseArgs) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponseOutput).ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput
	ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput
}

type privateLinkServiceConnectionStateResponsePtrType PrivateLinkServiceConnectionStateResponseArgs

func PrivateLinkServiceConnectionStateResponsePtr(v *PrivateLinkServiceConnectionStateResponseArgs) PrivateLinkServiceConnectionStateResponsePtrInput {
	return (*privateLinkServiceConnectionStateResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStateResponsePtrType) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateResponse) *PrivateLinkServiceConnectionStateResponse {
		return &v
	}).(PrivateLinkServiceConnectionStateResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutput() PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ToPrivateLinkServiceConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Elem() PrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) PrivateLinkServiceConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateResponse
		return ret
	}).(PrivateLinkServiceConnectionStateResponseOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type ProgressPropertiesResponse struct {
	Percentage *string `pulumi:"percentage"`
}





type ProgressPropertiesResponseInput interface {
	pulumi.Input

	ToProgressPropertiesResponseOutput() ProgressPropertiesResponseOutput
	ToProgressPropertiesResponseOutputWithContext(context.Context) ProgressPropertiesResponseOutput
}

type ProgressPropertiesResponseArgs struct {
	Percentage pulumi.StringPtrInput `pulumi:"percentage"`
}

func (ProgressPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ProgressPropertiesResponse)(nil)).Elem()
}

func (i ProgressPropertiesResponseArgs) ToProgressPropertiesResponseOutput() ProgressPropertiesResponseOutput {
	return i.ToProgressPropertiesResponseOutputWithContext(context.Background())
}

func (i ProgressPropertiesResponseArgs) ToProgressPropertiesResponseOutputWithContext(ctx context.Context) ProgressPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProgressPropertiesResponseOutput)
}

func (i ProgressPropertiesResponseArgs) ToProgressPropertiesResponsePtrOutput() ProgressPropertiesResponsePtrOutput {
	return i.ToProgressPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ProgressPropertiesResponseArgs) ToProgressPropertiesResponsePtrOutputWithContext(ctx context.Context) ProgressPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProgressPropertiesResponseOutput).ToProgressPropertiesResponsePtrOutputWithContext(ctx)
}









type ProgressPropertiesResponsePtrInput interface {
	pulumi.Input

	ToProgressPropertiesResponsePtrOutput() ProgressPropertiesResponsePtrOutput
	ToProgressPropertiesResponsePtrOutputWithContext(context.Context) ProgressPropertiesResponsePtrOutput
}

type progressPropertiesResponsePtrType ProgressPropertiesResponseArgs

func ProgressPropertiesResponsePtr(v *ProgressPropertiesResponseArgs) ProgressPropertiesResponsePtrInput {
	return (*progressPropertiesResponsePtrType)(v)
}

func (*progressPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ProgressPropertiesResponse)(nil)).Elem()
}

func (i *progressPropertiesResponsePtrType) ToProgressPropertiesResponsePtrOutput() ProgressPropertiesResponsePtrOutput {
	return i.ToProgressPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *progressPropertiesResponsePtrType) ToProgressPropertiesResponsePtrOutputWithContext(ctx context.Context) ProgressPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ProgressPropertiesResponsePtrOutput)
}

type ProgressPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ProgressPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ProgressPropertiesResponse)(nil)).Elem()
}

func (o ProgressPropertiesResponseOutput) ToProgressPropertiesResponseOutput() ProgressPropertiesResponseOutput {
	return o
}

func (o ProgressPropertiesResponseOutput) ToProgressPropertiesResponseOutputWithContext(ctx context.Context) ProgressPropertiesResponseOutput {
	return o
}

func (o ProgressPropertiesResponseOutput) ToProgressPropertiesResponsePtrOutput() ProgressPropertiesResponsePtrOutput {
	return o.ToProgressPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ProgressPropertiesResponseOutput) ToProgressPropertiesResponsePtrOutputWithContext(ctx context.Context) ProgressPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ProgressPropertiesResponse) *ProgressPropertiesResponse {
		return &v
	}).(ProgressPropertiesResponsePtrOutput)
}

func (o ProgressPropertiesResponseOutput) Percentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ProgressPropertiesResponse) *string { return v.Percentage }).(pulumi.StringPtrOutput)
}

type ProgressPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ProgressPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ProgressPropertiesResponse)(nil)).Elem()
}

func (o ProgressPropertiesResponsePtrOutput) ToProgressPropertiesResponsePtrOutput() ProgressPropertiesResponsePtrOutput {
	return o
}

func (o ProgressPropertiesResponsePtrOutput) ToProgressPropertiesResponsePtrOutputWithContext(ctx context.Context) ProgressPropertiesResponsePtrOutput {
	return o
}

func (o ProgressPropertiesResponsePtrOutput) Elem() ProgressPropertiesResponseOutput {
	return o.ApplyT(func(v *ProgressPropertiesResponse) ProgressPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ProgressPropertiesResponse
		return ret
	}).(ProgressPropertiesResponseOutput)
}

func (o ProgressPropertiesResponsePtrOutput) Percentage() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ProgressPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Percentage
	}).(pulumi.StringPtrOutput)
}

type QuarantinePolicy struct {
	Status *string `pulumi:"status"`
}


func (val *QuarantinePolicy) Defaults() *QuarantinePolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type QuarantinePolicyInput interface {
	pulumi.Input

	ToQuarantinePolicyOutput() QuarantinePolicyOutput
	ToQuarantinePolicyOutputWithContext(context.Context) QuarantinePolicyOutput
}

type QuarantinePolicyArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (QuarantinePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicy)(nil)).Elem()
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyOutput() QuarantinePolicyOutput {
	return i.ToQuarantinePolicyOutputWithContext(context.Background())
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyOutputWithContext(ctx context.Context) QuarantinePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyOutput)
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return i.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (i QuarantinePolicyArgs) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyOutput).ToQuarantinePolicyPtrOutputWithContext(ctx)
}









type QuarantinePolicyPtrInput interface {
	pulumi.Input

	ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput
	ToQuarantinePolicyPtrOutputWithContext(context.Context) QuarantinePolicyPtrOutput
}

type quarantinePolicyPtrType QuarantinePolicyArgs

func QuarantinePolicyPtr(v *QuarantinePolicyArgs) QuarantinePolicyPtrInput {
	return (*quarantinePolicyPtrType)(v)
}

func (*quarantinePolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicy)(nil)).Elem()
}

func (i *quarantinePolicyPtrType) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return i.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (i *quarantinePolicyPtrType) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyPtrOutput)
}

type QuarantinePolicyOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicy)(nil)).Elem()
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyOutput() QuarantinePolicyOutput {
	return o
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyOutputWithContext(ctx context.Context) QuarantinePolicyOutput {
	return o
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return o.ToQuarantinePolicyPtrOutputWithContext(context.Background())
}

func (o QuarantinePolicyOutput) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuarantinePolicy) *QuarantinePolicy {
		return &v
	}).(QuarantinePolicyPtrOutput)
}

func (o QuarantinePolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuarantinePolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type QuarantinePolicyPtrOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicy)(nil)).Elem()
}

func (o QuarantinePolicyPtrOutput) ToQuarantinePolicyPtrOutput() QuarantinePolicyPtrOutput {
	return o
}

func (o QuarantinePolicyPtrOutput) ToQuarantinePolicyPtrOutputWithContext(ctx context.Context) QuarantinePolicyPtrOutput {
	return o
}

func (o QuarantinePolicyPtrOutput) Elem() QuarantinePolicyOutput {
	return o.ApplyT(func(v *QuarantinePolicy) QuarantinePolicy {
		if v != nil {
			return *v
		}
		var ret QuarantinePolicy
		return ret
	}).(QuarantinePolicyOutput)
}

func (o QuarantinePolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuarantinePolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type QuarantinePolicyResponse struct {
	Status *string `pulumi:"status"`
}


func (val *QuarantinePolicyResponse) Defaults() *QuarantinePolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type QuarantinePolicyResponseInput interface {
	pulumi.Input

	ToQuarantinePolicyResponseOutput() QuarantinePolicyResponseOutput
	ToQuarantinePolicyResponseOutputWithContext(context.Context) QuarantinePolicyResponseOutput
}

type QuarantinePolicyResponseArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (QuarantinePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicyResponse)(nil)).Elem()
}

func (i QuarantinePolicyResponseArgs) ToQuarantinePolicyResponseOutput() QuarantinePolicyResponseOutput {
	return i.ToQuarantinePolicyResponseOutputWithContext(context.Background())
}

func (i QuarantinePolicyResponseArgs) ToQuarantinePolicyResponseOutputWithContext(ctx context.Context) QuarantinePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyResponseOutput)
}

func (i QuarantinePolicyResponseArgs) ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput {
	return i.ToQuarantinePolicyResponsePtrOutputWithContext(context.Background())
}

func (i QuarantinePolicyResponseArgs) ToQuarantinePolicyResponsePtrOutputWithContext(ctx context.Context) QuarantinePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyResponseOutput).ToQuarantinePolicyResponsePtrOutputWithContext(ctx)
}









type QuarantinePolicyResponsePtrInput interface {
	pulumi.Input

	ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput
	ToQuarantinePolicyResponsePtrOutputWithContext(context.Context) QuarantinePolicyResponsePtrOutput
}

type quarantinePolicyResponsePtrType QuarantinePolicyResponseArgs

func QuarantinePolicyResponsePtr(v *QuarantinePolicyResponseArgs) QuarantinePolicyResponsePtrInput {
	return (*quarantinePolicyResponsePtrType)(v)
}

func (*quarantinePolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicyResponse)(nil)).Elem()
}

func (i *quarantinePolicyResponsePtrType) ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput {
	return i.ToQuarantinePolicyResponsePtrOutputWithContext(context.Background())
}

func (i *quarantinePolicyResponsePtrType) ToQuarantinePolicyResponsePtrOutputWithContext(ctx context.Context) QuarantinePolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QuarantinePolicyResponsePtrOutput)
}

type QuarantinePolicyResponseOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QuarantinePolicyResponse)(nil)).Elem()
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponseOutput() QuarantinePolicyResponseOutput {
	return o
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponseOutputWithContext(ctx context.Context) QuarantinePolicyResponseOutput {
	return o
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput {
	return o.ToQuarantinePolicyResponsePtrOutputWithContext(context.Background())
}

func (o QuarantinePolicyResponseOutput) ToQuarantinePolicyResponsePtrOutputWithContext(ctx context.Context) QuarantinePolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QuarantinePolicyResponse) *QuarantinePolicyResponse {
		return &v
	}).(QuarantinePolicyResponsePtrOutput)
}

func (o QuarantinePolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QuarantinePolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type QuarantinePolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (QuarantinePolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QuarantinePolicyResponse)(nil)).Elem()
}

func (o QuarantinePolicyResponsePtrOutput) ToQuarantinePolicyResponsePtrOutput() QuarantinePolicyResponsePtrOutput {
	return o
}

func (o QuarantinePolicyResponsePtrOutput) ToQuarantinePolicyResponsePtrOutputWithContext(ctx context.Context) QuarantinePolicyResponsePtrOutput {
	return o
}

func (o QuarantinePolicyResponsePtrOutput) Elem() QuarantinePolicyResponseOutput {
	return o.ApplyT(func(v *QuarantinePolicyResponse) QuarantinePolicyResponse {
		if v != nil {
			return *v
		}
		var ret QuarantinePolicyResponse
		return ret
	}).(QuarantinePolicyResponseOutput)
}

func (o QuarantinePolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QuarantinePolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
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

type RetentionPolicy struct {
	Days   *int    `pulumi:"days"`
	Status *string `pulumi:"status"`
}


func (val *RetentionPolicy) Defaults() *RetentionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Days) {
		days_ := 7
		tmp.Days = &days_
	}
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days   pulumi.IntPtrInput    `pulumi:"days"`
	Status pulumi.StringPtrInput `pulumi:"status"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}









type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicy) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RetentionPolicyResponse struct {
	Days            *int    `pulumi:"days"`
	LastUpdatedTime string  `pulumi:"lastUpdatedTime"`
	Status          *string `pulumi:"status"`
}


func (val *RetentionPolicyResponse) Defaults() *RetentionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Days) {
		days_ := 7
		tmp.Days = &days_
	}
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	return &tmp
}





type RetentionPolicyResponseInput interface {
	pulumi.Input

	ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput
	ToRetentionPolicyResponseOutputWithContext(context.Context) RetentionPolicyResponseOutput
}

type RetentionPolicyResponseArgs struct {
	Days            pulumi.IntPtrInput    `pulumi:"days"`
	LastUpdatedTime pulumi.StringInput    `pulumi:"lastUpdatedTime"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (RetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return i.ToRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponseOutput)
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return i.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i RetentionPolicyResponseArgs) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponseOutput).ToRetentionPolicyResponsePtrOutputWithContext(ctx)
}









type RetentionPolicyResponsePtrInput interface {
	pulumi.Input

	ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput
	ToRetentionPolicyResponsePtrOutputWithContext(context.Context) RetentionPolicyResponsePtrOutput
}

type retentionPolicyResponsePtrType RetentionPolicyResponseArgs

func RetentionPolicyResponsePtr(v *RetentionPolicyResponseArgs) RetentionPolicyResponsePtrInput {
	return (*retentionPolicyResponsePtrType)(v)
}

func (*retentionPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (i *retentionPolicyResponsePtrType) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return i.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *retentionPolicyResponsePtrType) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyResponsePtrOutput)
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o.ToRetentionPolicyResponsePtrOutputWithContext(context.Background())
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicyResponse) *RetentionPolicyResponse {
		return &v
	}).(RetentionPolicyResponsePtrOutput)
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) *int { return v.Days }).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponseOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) string { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o RetentionPolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) LastUpdatedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTime
	}).(pulumi.StringPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
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

type StatusDetailPropertiesResponse struct {
	Code          string `pulumi:"code"`
	CorrelationId string `pulumi:"correlationId"`
	Description   string `pulumi:"description"`
	Timestamp     string `pulumi:"timestamp"`
	Type          string `pulumi:"type"`
}





type StatusDetailPropertiesResponseInput interface {
	pulumi.Input

	ToStatusDetailPropertiesResponseOutput() StatusDetailPropertiesResponseOutput
	ToStatusDetailPropertiesResponseOutputWithContext(context.Context) StatusDetailPropertiesResponseOutput
}

type StatusDetailPropertiesResponseArgs struct {
	Code          pulumi.StringInput `pulumi:"code"`
	CorrelationId pulumi.StringInput `pulumi:"correlationId"`
	Description   pulumi.StringInput `pulumi:"description"`
	Timestamp     pulumi.StringInput `pulumi:"timestamp"`
	Type          pulumi.StringInput `pulumi:"type"`
}

func (StatusDetailPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusDetailPropertiesResponse)(nil)).Elem()
}

func (i StatusDetailPropertiesResponseArgs) ToStatusDetailPropertiesResponseOutput() StatusDetailPropertiesResponseOutput {
	return i.ToStatusDetailPropertiesResponseOutputWithContext(context.Background())
}

func (i StatusDetailPropertiesResponseArgs) ToStatusDetailPropertiesResponseOutputWithContext(ctx context.Context) StatusDetailPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusDetailPropertiesResponseOutput)
}





type StatusDetailPropertiesResponseArrayInput interface {
	pulumi.Input

	ToStatusDetailPropertiesResponseArrayOutput() StatusDetailPropertiesResponseArrayOutput
	ToStatusDetailPropertiesResponseArrayOutputWithContext(context.Context) StatusDetailPropertiesResponseArrayOutput
}

type StatusDetailPropertiesResponseArray []StatusDetailPropertiesResponseInput

func (StatusDetailPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusDetailPropertiesResponse)(nil)).Elem()
}

func (i StatusDetailPropertiesResponseArray) ToStatusDetailPropertiesResponseArrayOutput() StatusDetailPropertiesResponseArrayOutput {
	return i.ToStatusDetailPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i StatusDetailPropertiesResponseArray) ToStatusDetailPropertiesResponseArrayOutputWithContext(ctx context.Context) StatusDetailPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusDetailPropertiesResponseArrayOutput)
}

type StatusDetailPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StatusDetailPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusDetailPropertiesResponse)(nil)).Elem()
}

func (o StatusDetailPropertiesResponseOutput) ToStatusDetailPropertiesResponseOutput() StatusDetailPropertiesResponseOutput {
	return o
}

func (o StatusDetailPropertiesResponseOutput) ToStatusDetailPropertiesResponseOutputWithContext(ctx context.Context) StatusDetailPropertiesResponseOutput {
	return o
}

func (o StatusDetailPropertiesResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) Timestamp() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Timestamp }).(pulumi.StringOutput)
}

func (o StatusDetailPropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StatusDetailPropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StatusDetailPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusDetailPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusDetailPropertiesResponse)(nil)).Elem()
}

func (o StatusDetailPropertiesResponseArrayOutput) ToStatusDetailPropertiesResponseArrayOutput() StatusDetailPropertiesResponseArrayOutput {
	return o
}

func (o StatusDetailPropertiesResponseArrayOutput) ToStatusDetailPropertiesResponseArrayOutputWithContext(ctx context.Context) StatusDetailPropertiesResponseArrayOutput {
	return o
}

func (o StatusDetailPropertiesResponseArrayOutput) Index(i pulumi.IntInput) StatusDetailPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusDetailPropertiesResponse {
		return vs[0].([]StatusDetailPropertiesResponse)[vs[1].(int)]
	}).(StatusDetailPropertiesResponseOutput)
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

type SyncProperties struct {
	MessageTtl string  `pulumi:"messageTtl"`
	Schedule   *string `pulumi:"schedule"`
	SyncWindow *string `pulumi:"syncWindow"`
	TokenId    string  `pulumi:"tokenId"`
}





type SyncPropertiesInput interface {
	pulumi.Input

	ToSyncPropertiesOutput() SyncPropertiesOutput
	ToSyncPropertiesOutputWithContext(context.Context) SyncPropertiesOutput
}

type SyncPropertiesArgs struct {
	MessageTtl pulumi.StringInput    `pulumi:"messageTtl"`
	Schedule   pulumi.StringPtrInput `pulumi:"schedule"`
	SyncWindow pulumi.StringPtrInput `pulumi:"syncWindow"`
	TokenId    pulumi.StringInput    `pulumi:"tokenId"`
}

func (SyncPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProperties)(nil)).Elem()
}

func (i SyncPropertiesArgs) ToSyncPropertiesOutput() SyncPropertiesOutput {
	return i.ToSyncPropertiesOutputWithContext(context.Background())
}

func (i SyncPropertiesArgs) ToSyncPropertiesOutputWithContext(ctx context.Context) SyncPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesOutput)
}

func (i SyncPropertiesArgs) ToSyncPropertiesPtrOutput() SyncPropertiesPtrOutput {
	return i.ToSyncPropertiesPtrOutputWithContext(context.Background())
}

func (i SyncPropertiesArgs) ToSyncPropertiesPtrOutputWithContext(ctx context.Context) SyncPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesOutput).ToSyncPropertiesPtrOutputWithContext(ctx)
}









type SyncPropertiesPtrInput interface {
	pulumi.Input

	ToSyncPropertiesPtrOutput() SyncPropertiesPtrOutput
	ToSyncPropertiesPtrOutputWithContext(context.Context) SyncPropertiesPtrOutput
}

type syncPropertiesPtrType SyncPropertiesArgs

func SyncPropertiesPtr(v *SyncPropertiesArgs) SyncPropertiesPtrInput {
	return (*syncPropertiesPtrType)(v)
}

func (*syncPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncProperties)(nil)).Elem()
}

func (i *syncPropertiesPtrType) ToSyncPropertiesPtrOutput() SyncPropertiesPtrOutput {
	return i.ToSyncPropertiesPtrOutputWithContext(context.Background())
}

func (i *syncPropertiesPtrType) ToSyncPropertiesPtrOutputWithContext(ctx context.Context) SyncPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesPtrOutput)
}

type SyncPropertiesOutput struct{ *pulumi.OutputState }

func (SyncPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncProperties)(nil)).Elem()
}

func (o SyncPropertiesOutput) ToSyncPropertiesOutput() SyncPropertiesOutput {
	return o
}

func (o SyncPropertiesOutput) ToSyncPropertiesOutputWithContext(ctx context.Context) SyncPropertiesOutput {
	return o
}

func (o SyncPropertiesOutput) ToSyncPropertiesPtrOutput() SyncPropertiesPtrOutput {
	return o.ToSyncPropertiesPtrOutputWithContext(context.Background())
}

func (o SyncPropertiesOutput) ToSyncPropertiesPtrOutputWithContext(ctx context.Context) SyncPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncProperties) *SyncProperties {
		return &v
	}).(SyncPropertiesPtrOutput)
}

func (o SyncPropertiesOutput) MessageTtl() pulumi.StringOutput {
	return o.ApplyT(func(v SyncProperties) string { return v.MessageTtl }).(pulumi.StringOutput)
}

func (o SyncPropertiesOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProperties) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesOutput) SyncWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncProperties) *string { return v.SyncWindow }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v SyncProperties) string { return v.TokenId }).(pulumi.StringOutput)
}

type SyncPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SyncPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncProperties)(nil)).Elem()
}

func (o SyncPropertiesPtrOutput) ToSyncPropertiesPtrOutput() SyncPropertiesPtrOutput {
	return o
}

func (o SyncPropertiesPtrOutput) ToSyncPropertiesPtrOutputWithContext(ctx context.Context) SyncPropertiesPtrOutput {
	return o
}

func (o SyncPropertiesPtrOutput) Elem() SyncPropertiesOutput {
	return o.ApplyT(func(v *SyncProperties) SyncProperties {
		if v != nil {
			return *v
		}
		var ret SyncProperties
		return ret
	}).(SyncPropertiesOutput)
}

func (o SyncPropertiesPtrOutput) MessageTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProperties) *string {
		if v == nil {
			return nil
		}
		return &v.MessageTtl
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesPtrOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProperties) *string {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesPtrOutput) SyncWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProperties) *string {
		if v == nil {
			return nil
		}
		return v.SyncWindow
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesPtrOutput) TokenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncProperties) *string {
		if v == nil {
			return nil
		}
		return &v.TokenId
	}).(pulumi.StringPtrOutput)
}

type SyncPropertiesResponse struct {
	GatewayEndpoint string  `pulumi:"gatewayEndpoint"`
	LastSyncTime    string  `pulumi:"lastSyncTime"`
	MessageTtl      string  `pulumi:"messageTtl"`
	Schedule        *string `pulumi:"schedule"`
	SyncWindow      *string `pulumi:"syncWindow"`
	TokenId         string  `pulumi:"tokenId"`
}





type SyncPropertiesResponseInput interface {
	pulumi.Input

	ToSyncPropertiesResponseOutput() SyncPropertiesResponseOutput
	ToSyncPropertiesResponseOutputWithContext(context.Context) SyncPropertiesResponseOutput
}

type SyncPropertiesResponseArgs struct {
	GatewayEndpoint pulumi.StringInput    `pulumi:"gatewayEndpoint"`
	LastSyncTime    pulumi.StringInput    `pulumi:"lastSyncTime"`
	MessageTtl      pulumi.StringInput    `pulumi:"messageTtl"`
	Schedule        pulumi.StringPtrInput `pulumi:"schedule"`
	SyncWindow      pulumi.StringPtrInput `pulumi:"syncWindow"`
	TokenId         pulumi.StringInput    `pulumi:"tokenId"`
}

func (SyncPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncPropertiesResponse)(nil)).Elem()
}

func (i SyncPropertiesResponseArgs) ToSyncPropertiesResponseOutput() SyncPropertiesResponseOutput {
	return i.ToSyncPropertiesResponseOutputWithContext(context.Background())
}

func (i SyncPropertiesResponseArgs) ToSyncPropertiesResponseOutputWithContext(ctx context.Context) SyncPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesResponseOutput)
}

func (i SyncPropertiesResponseArgs) ToSyncPropertiesResponsePtrOutput() SyncPropertiesResponsePtrOutput {
	return i.ToSyncPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SyncPropertiesResponseArgs) ToSyncPropertiesResponsePtrOutputWithContext(ctx context.Context) SyncPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesResponseOutput).ToSyncPropertiesResponsePtrOutputWithContext(ctx)
}









type SyncPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSyncPropertiesResponsePtrOutput() SyncPropertiesResponsePtrOutput
	ToSyncPropertiesResponsePtrOutputWithContext(context.Context) SyncPropertiesResponsePtrOutput
}

type syncPropertiesResponsePtrType SyncPropertiesResponseArgs

func SyncPropertiesResponsePtr(v *SyncPropertiesResponseArgs) SyncPropertiesResponsePtrInput {
	return (*syncPropertiesResponsePtrType)(v)
}

func (*syncPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncPropertiesResponse)(nil)).Elem()
}

func (i *syncPropertiesResponsePtrType) ToSyncPropertiesResponsePtrOutput() SyncPropertiesResponsePtrOutput {
	return i.ToSyncPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *syncPropertiesResponsePtrType) ToSyncPropertiesResponsePtrOutputWithContext(ctx context.Context) SyncPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SyncPropertiesResponsePtrOutput)
}

type SyncPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SyncPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SyncPropertiesResponse)(nil)).Elem()
}

func (o SyncPropertiesResponseOutput) ToSyncPropertiesResponseOutput() SyncPropertiesResponseOutput {
	return o
}

func (o SyncPropertiesResponseOutput) ToSyncPropertiesResponseOutputWithContext(ctx context.Context) SyncPropertiesResponseOutput {
	return o
}

func (o SyncPropertiesResponseOutput) ToSyncPropertiesResponsePtrOutput() SyncPropertiesResponsePtrOutput {
	return o.ToSyncPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SyncPropertiesResponseOutput) ToSyncPropertiesResponsePtrOutputWithContext(ctx context.Context) SyncPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SyncPropertiesResponse) *SyncPropertiesResponse {
		return &v
	}).(SyncPropertiesResponsePtrOutput)
}

func (o SyncPropertiesResponseOutput) GatewayEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.GatewayEndpoint }).(pulumi.StringOutput)
}

func (o SyncPropertiesResponseOutput) LastSyncTime() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.LastSyncTime }).(pulumi.StringOutput)
}

func (o SyncPropertiesResponseOutput) MessageTtl() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.MessageTtl }).(pulumi.StringOutput)
}

func (o SyncPropertiesResponseOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) *string { return v.Schedule }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponseOutput) SyncWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) *string { return v.SyncWindow }).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponseOutput) TokenId() pulumi.StringOutput {
	return o.ApplyT(func(v SyncPropertiesResponse) string { return v.TokenId }).(pulumi.StringOutput)
}

type SyncPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SyncPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SyncPropertiesResponse)(nil)).Elem()
}

func (o SyncPropertiesResponsePtrOutput) ToSyncPropertiesResponsePtrOutput() SyncPropertiesResponsePtrOutput {
	return o
}

func (o SyncPropertiesResponsePtrOutput) ToSyncPropertiesResponsePtrOutputWithContext(ctx context.Context) SyncPropertiesResponsePtrOutput {
	return o
}

func (o SyncPropertiesResponsePtrOutput) Elem() SyncPropertiesResponseOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) SyncPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SyncPropertiesResponse
		return ret
	}).(SyncPropertiesResponseOutput)
}

func (o SyncPropertiesResponsePtrOutput) GatewayEndpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.GatewayEndpoint
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponsePtrOutput) LastSyncTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSyncTime
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponsePtrOutput) MessageTtl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MessageTtl
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponsePtrOutput) Schedule() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponsePtrOutput) SyncWindow() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SyncWindow
	}).(pulumi.StringPtrOutput)
}

func (o SyncPropertiesResponsePtrOutput) TokenId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SyncPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TokenId
	}).(pulumi.StringPtrOutput)
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

type TlsCertificatePropertiesResponse struct {
	Location string `pulumi:"location"`
	Type     string `pulumi:"type"`
}





type TlsCertificatePropertiesResponseInput interface {
	pulumi.Input

	ToTlsCertificatePropertiesResponseOutput() TlsCertificatePropertiesResponseOutput
	ToTlsCertificatePropertiesResponseOutputWithContext(context.Context) TlsCertificatePropertiesResponseOutput
}

type TlsCertificatePropertiesResponseArgs struct {
	Location pulumi.StringInput `pulumi:"location"`
	Type     pulumi.StringInput `pulumi:"type"`
}

func (TlsCertificatePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsCertificatePropertiesResponse)(nil)).Elem()
}

func (i TlsCertificatePropertiesResponseArgs) ToTlsCertificatePropertiesResponseOutput() TlsCertificatePropertiesResponseOutput {
	return i.ToTlsCertificatePropertiesResponseOutputWithContext(context.Background())
}

func (i TlsCertificatePropertiesResponseArgs) ToTlsCertificatePropertiesResponseOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsCertificatePropertiesResponseOutput)
}

func (i TlsCertificatePropertiesResponseArgs) ToTlsCertificatePropertiesResponsePtrOutput() TlsCertificatePropertiesResponsePtrOutput {
	return i.ToTlsCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TlsCertificatePropertiesResponseArgs) ToTlsCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsCertificatePropertiesResponseOutput).ToTlsCertificatePropertiesResponsePtrOutputWithContext(ctx)
}









type TlsCertificatePropertiesResponsePtrInput interface {
	pulumi.Input

	ToTlsCertificatePropertiesResponsePtrOutput() TlsCertificatePropertiesResponsePtrOutput
	ToTlsCertificatePropertiesResponsePtrOutputWithContext(context.Context) TlsCertificatePropertiesResponsePtrOutput
}

type tlsCertificatePropertiesResponsePtrType TlsCertificatePropertiesResponseArgs

func TlsCertificatePropertiesResponsePtr(v *TlsCertificatePropertiesResponseArgs) TlsCertificatePropertiesResponsePtrInput {
	return (*tlsCertificatePropertiesResponsePtrType)(v)
}

func (*tlsCertificatePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsCertificatePropertiesResponse)(nil)).Elem()
}

func (i *tlsCertificatePropertiesResponsePtrType) ToTlsCertificatePropertiesResponsePtrOutput() TlsCertificatePropertiesResponsePtrOutput {
	return i.ToTlsCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *tlsCertificatePropertiesResponsePtrType) ToTlsCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsCertificatePropertiesResponsePtrOutput)
}

type TlsCertificatePropertiesResponseOutput struct{ *pulumi.OutputState }

func (TlsCertificatePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsCertificatePropertiesResponse)(nil)).Elem()
}

func (o TlsCertificatePropertiesResponseOutput) ToTlsCertificatePropertiesResponseOutput() TlsCertificatePropertiesResponseOutput {
	return o
}

func (o TlsCertificatePropertiesResponseOutput) ToTlsCertificatePropertiesResponseOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponseOutput {
	return o
}

func (o TlsCertificatePropertiesResponseOutput) ToTlsCertificatePropertiesResponsePtrOutput() TlsCertificatePropertiesResponsePtrOutput {
	return o.ToTlsCertificatePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TlsCertificatePropertiesResponseOutput) ToTlsCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TlsCertificatePropertiesResponse) *TlsCertificatePropertiesResponse {
		return &v
	}).(TlsCertificatePropertiesResponsePtrOutput)
}

func (o TlsCertificatePropertiesResponseOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v TlsCertificatePropertiesResponse) string { return v.Location }).(pulumi.StringOutput)
}

func (o TlsCertificatePropertiesResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v TlsCertificatePropertiesResponse) string { return v.Type }).(pulumi.StringOutput)
}

type TlsCertificatePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TlsCertificatePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsCertificatePropertiesResponse)(nil)).Elem()
}

func (o TlsCertificatePropertiesResponsePtrOutput) ToTlsCertificatePropertiesResponsePtrOutput() TlsCertificatePropertiesResponsePtrOutput {
	return o
}

func (o TlsCertificatePropertiesResponsePtrOutput) ToTlsCertificatePropertiesResponsePtrOutputWithContext(ctx context.Context) TlsCertificatePropertiesResponsePtrOutput {
	return o
}

func (o TlsCertificatePropertiesResponsePtrOutput) Elem() TlsCertificatePropertiesResponseOutput {
	return o.ApplyT(func(v *TlsCertificatePropertiesResponse) TlsCertificatePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TlsCertificatePropertiesResponse
		return ret
	}).(TlsCertificatePropertiesResponseOutput)
}

func (o TlsCertificatePropertiesResponsePtrOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsCertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Location
	}).(pulumi.StringPtrOutput)
}

func (o TlsCertificatePropertiesResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsCertificatePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type TlsPropertiesResponse struct {
	Certificate TlsCertificatePropertiesResponse `pulumi:"certificate"`
	Status      string                           `pulumi:"status"`
}





type TlsPropertiesResponseInput interface {
	pulumi.Input

	ToTlsPropertiesResponseOutput() TlsPropertiesResponseOutput
	ToTlsPropertiesResponseOutputWithContext(context.Context) TlsPropertiesResponseOutput
}

type TlsPropertiesResponseArgs struct {
	Certificate TlsCertificatePropertiesResponseInput `pulumi:"certificate"`
	Status      pulumi.StringInput                    `pulumi:"status"`
}

func (TlsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsPropertiesResponse)(nil)).Elem()
}

func (i TlsPropertiesResponseArgs) ToTlsPropertiesResponseOutput() TlsPropertiesResponseOutput {
	return i.ToTlsPropertiesResponseOutputWithContext(context.Background())
}

func (i TlsPropertiesResponseArgs) ToTlsPropertiesResponseOutputWithContext(ctx context.Context) TlsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsPropertiesResponseOutput)
}

func (i TlsPropertiesResponseArgs) ToTlsPropertiesResponsePtrOutput() TlsPropertiesResponsePtrOutput {
	return i.ToTlsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TlsPropertiesResponseArgs) ToTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) TlsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsPropertiesResponseOutput).ToTlsPropertiesResponsePtrOutputWithContext(ctx)
}









type TlsPropertiesResponsePtrInput interface {
	pulumi.Input

	ToTlsPropertiesResponsePtrOutput() TlsPropertiesResponsePtrOutput
	ToTlsPropertiesResponsePtrOutputWithContext(context.Context) TlsPropertiesResponsePtrOutput
}

type tlsPropertiesResponsePtrType TlsPropertiesResponseArgs

func TlsPropertiesResponsePtr(v *TlsPropertiesResponseArgs) TlsPropertiesResponsePtrInput {
	return (*tlsPropertiesResponsePtrType)(v)
}

func (*tlsPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsPropertiesResponse)(nil)).Elem()
}

func (i *tlsPropertiesResponsePtrType) ToTlsPropertiesResponsePtrOutput() TlsPropertiesResponsePtrOutput {
	return i.ToTlsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *tlsPropertiesResponsePtrType) ToTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) TlsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TlsPropertiesResponsePtrOutput)
}

type TlsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TlsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TlsPropertiesResponse)(nil)).Elem()
}

func (o TlsPropertiesResponseOutput) ToTlsPropertiesResponseOutput() TlsPropertiesResponseOutput {
	return o
}

func (o TlsPropertiesResponseOutput) ToTlsPropertiesResponseOutputWithContext(ctx context.Context) TlsPropertiesResponseOutput {
	return o
}

func (o TlsPropertiesResponseOutput) ToTlsPropertiesResponsePtrOutput() TlsPropertiesResponsePtrOutput {
	return o.ToTlsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TlsPropertiesResponseOutput) ToTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) TlsPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TlsPropertiesResponse) *TlsPropertiesResponse {
		return &v
	}).(TlsPropertiesResponsePtrOutput)
}

func (o TlsPropertiesResponseOutput) Certificate() TlsCertificatePropertiesResponseOutput {
	return o.ApplyT(func(v TlsPropertiesResponse) TlsCertificatePropertiesResponse { return v.Certificate }).(TlsCertificatePropertiesResponseOutput)
}

func (o TlsPropertiesResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v TlsPropertiesResponse) string { return v.Status }).(pulumi.StringOutput)
}

type TlsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TlsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TlsPropertiesResponse)(nil)).Elem()
}

func (o TlsPropertiesResponsePtrOutput) ToTlsPropertiesResponsePtrOutput() TlsPropertiesResponsePtrOutput {
	return o
}

func (o TlsPropertiesResponsePtrOutput) ToTlsPropertiesResponsePtrOutputWithContext(ctx context.Context) TlsPropertiesResponsePtrOutput {
	return o
}

func (o TlsPropertiesResponsePtrOutput) Elem() TlsPropertiesResponseOutput {
	return o.ApplyT(func(v *TlsPropertiesResponse) TlsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TlsPropertiesResponse
		return ret
	}).(TlsPropertiesResponseOutput)
}

func (o TlsPropertiesResponsePtrOutput) Certificate() TlsCertificatePropertiesResponsePtrOutput {
	return o.ApplyT(func(v *TlsPropertiesResponse) *TlsCertificatePropertiesResponse {
		if v == nil {
			return nil
		}
		return &v.Certificate
	}).(TlsCertificatePropertiesResponsePtrOutput)
}

func (o TlsPropertiesResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TlsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

type TokenCertificate struct {
	EncodedPemCertificate *string `pulumi:"encodedPemCertificate"`
	Expiry                *string `pulumi:"expiry"`
	Name                  *string `pulumi:"name"`
	Thumbprint            *string `pulumi:"thumbprint"`
}





type TokenCertificateInput interface {
	pulumi.Input

	ToTokenCertificateOutput() TokenCertificateOutput
	ToTokenCertificateOutputWithContext(context.Context) TokenCertificateOutput
}

type TokenCertificateArgs struct {
	EncodedPemCertificate pulumi.StringPtrInput `pulumi:"encodedPemCertificate"`
	Expiry                pulumi.StringPtrInput `pulumi:"expiry"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint            pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (TokenCertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificate)(nil)).Elem()
}

func (i TokenCertificateArgs) ToTokenCertificateOutput() TokenCertificateOutput {
	return i.ToTokenCertificateOutputWithContext(context.Background())
}

func (i TokenCertificateArgs) ToTokenCertificateOutputWithContext(ctx context.Context) TokenCertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateOutput)
}





type TokenCertificateArrayInput interface {
	pulumi.Input

	ToTokenCertificateArrayOutput() TokenCertificateArrayOutput
	ToTokenCertificateArrayOutputWithContext(context.Context) TokenCertificateArrayOutput
}

type TokenCertificateArray []TokenCertificateInput

func (TokenCertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificate)(nil)).Elem()
}

func (i TokenCertificateArray) ToTokenCertificateArrayOutput() TokenCertificateArrayOutput {
	return i.ToTokenCertificateArrayOutputWithContext(context.Background())
}

func (i TokenCertificateArray) ToTokenCertificateArrayOutputWithContext(ctx context.Context) TokenCertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateArrayOutput)
}

type TokenCertificateOutput struct{ *pulumi.OutputState }

func (TokenCertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificate)(nil)).Elem()
}

func (o TokenCertificateOutput) ToTokenCertificateOutput() TokenCertificateOutput {
	return o
}

func (o TokenCertificateOutput) ToTokenCertificateOutputWithContext(ctx context.Context) TokenCertificateOutput {
	return o
}

func (o TokenCertificateOutput) EncodedPemCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.EncodedPemCertificate }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificate) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type TokenCertificateArrayOutput struct{ *pulumi.OutputState }

func (TokenCertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificate)(nil)).Elem()
}

func (o TokenCertificateArrayOutput) ToTokenCertificateArrayOutput() TokenCertificateArrayOutput {
	return o
}

func (o TokenCertificateArrayOutput) ToTokenCertificateArrayOutputWithContext(ctx context.Context) TokenCertificateArrayOutput {
	return o
}

func (o TokenCertificateArrayOutput) Index(i pulumi.IntInput) TokenCertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenCertificate {
		return vs[0].([]TokenCertificate)[vs[1].(int)]
	}).(TokenCertificateOutput)
}

type TokenCertificateResponse struct {
	EncodedPemCertificate *string `pulumi:"encodedPemCertificate"`
	Expiry                *string `pulumi:"expiry"`
	Name                  *string `pulumi:"name"`
	Thumbprint            *string `pulumi:"thumbprint"`
}





type TokenCertificateResponseInput interface {
	pulumi.Input

	ToTokenCertificateResponseOutput() TokenCertificateResponseOutput
	ToTokenCertificateResponseOutputWithContext(context.Context) TokenCertificateResponseOutput
}

type TokenCertificateResponseArgs struct {
	EncodedPemCertificate pulumi.StringPtrInput `pulumi:"encodedPemCertificate"`
	Expiry                pulumi.StringPtrInput `pulumi:"expiry"`
	Name                  pulumi.StringPtrInput `pulumi:"name"`
	Thumbprint            pulumi.StringPtrInput `pulumi:"thumbprint"`
}

func (TokenCertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateResponse)(nil)).Elem()
}

func (i TokenCertificateResponseArgs) ToTokenCertificateResponseOutput() TokenCertificateResponseOutput {
	return i.ToTokenCertificateResponseOutputWithContext(context.Background())
}

func (i TokenCertificateResponseArgs) ToTokenCertificateResponseOutputWithContext(ctx context.Context) TokenCertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateResponseOutput)
}





type TokenCertificateResponseArrayInput interface {
	pulumi.Input

	ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput
	ToTokenCertificateResponseArrayOutputWithContext(context.Context) TokenCertificateResponseArrayOutput
}

type TokenCertificateResponseArray []TokenCertificateResponseInput

func (TokenCertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificateResponse)(nil)).Elem()
}

func (i TokenCertificateResponseArray) ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput {
	return i.ToTokenCertificateResponseArrayOutputWithContext(context.Background())
}

func (i TokenCertificateResponseArray) ToTokenCertificateResponseArrayOutputWithContext(ctx context.Context) TokenCertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCertificateResponseArrayOutput)
}

type TokenCertificateResponseOutput struct{ *pulumi.OutputState }

func (TokenCertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCertificateResponse)(nil)).Elem()
}

func (o TokenCertificateResponseOutput) ToTokenCertificateResponseOutput() TokenCertificateResponseOutput {
	return o
}

func (o TokenCertificateResponseOutput) ToTokenCertificateResponseOutputWithContext(ctx context.Context) TokenCertificateResponseOutput {
	return o
}

func (o TokenCertificateResponseOutput) EncodedPemCertificate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.EncodedPemCertificate }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenCertificateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenCertificateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

type TokenCertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenCertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenCertificateResponse)(nil)).Elem()
}

func (o TokenCertificateResponseArrayOutput) ToTokenCertificateResponseArrayOutput() TokenCertificateResponseArrayOutput {
	return o
}

func (o TokenCertificateResponseArrayOutput) ToTokenCertificateResponseArrayOutputWithContext(ctx context.Context) TokenCertificateResponseArrayOutput {
	return o
}

func (o TokenCertificateResponseArrayOutput) Index(i pulumi.IntInput) TokenCertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenCertificateResponse {
		return vs[0].([]TokenCertificateResponse)[vs[1].(int)]
	}).(TokenCertificateResponseOutput)
}

type TokenCredentialsProperties struct {
	Certificates []TokenCertificate `pulumi:"certificates"`
	Passwords    []TokenPassword    `pulumi:"passwords"`
}





type TokenCredentialsPropertiesInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput
	ToTokenCredentialsPropertiesOutputWithContext(context.Context) TokenCredentialsPropertiesOutput
}

type TokenCredentialsPropertiesArgs struct {
	Certificates TokenCertificateArrayInput `pulumi:"certificates"`
	Passwords    TokenPasswordArrayInput    `pulumi:"passwords"`
}

func (TokenCredentialsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsProperties)(nil)).Elem()
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput {
	return i.ToTokenCredentialsPropertiesOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesOutputWithContext(ctx context.Context) TokenCredentialsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesOutput)
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return i.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesArgs) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesOutput).ToTokenCredentialsPropertiesPtrOutputWithContext(ctx)
}









type TokenCredentialsPropertiesPtrInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput
	ToTokenCredentialsPropertiesPtrOutputWithContext(context.Context) TokenCredentialsPropertiesPtrOutput
}

type tokenCredentialsPropertiesPtrType TokenCredentialsPropertiesArgs

func TokenCredentialsPropertiesPtr(v *TokenCredentialsPropertiesArgs) TokenCredentialsPropertiesPtrInput {
	return (*tokenCredentialsPropertiesPtrType)(v)
}

func (*tokenCredentialsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsProperties)(nil)).Elem()
}

func (i *tokenCredentialsPropertiesPtrType) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return i.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (i *tokenCredentialsPropertiesPtrType) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesPtrOutput)
}

type TokenCredentialsPropertiesOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsProperties)(nil)).Elem()
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesOutput() TokenCredentialsPropertiesOutput {
	return o
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesOutputWithContext(ctx context.Context) TokenCredentialsPropertiesOutput {
	return o
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return o.ToTokenCredentialsPropertiesPtrOutputWithContext(context.Background())
}

func (o TokenCredentialsPropertiesOutput) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCredentialsProperties) *TokenCredentialsProperties {
		return &v
	}).(TokenCredentialsPropertiesPtrOutput)
}

func (o TokenCredentialsPropertiesOutput) Certificates() TokenCertificateArrayOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) []TokenCertificate { return v.Certificates }).(TokenCertificateArrayOutput)
}

func (o TokenCredentialsPropertiesOutput) Passwords() TokenPasswordArrayOutput {
	return o.ApplyT(func(v TokenCredentialsProperties) []TokenPassword { return v.Passwords }).(TokenPasswordArrayOutput)
}

type TokenCredentialsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsProperties)(nil)).Elem()
}

func (o TokenCredentialsPropertiesPtrOutput) ToTokenCredentialsPropertiesPtrOutput() TokenCredentialsPropertiesPtrOutput {
	return o
}

func (o TokenCredentialsPropertiesPtrOutput) ToTokenCredentialsPropertiesPtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesPtrOutput {
	return o
}

func (o TokenCredentialsPropertiesPtrOutput) Elem() TokenCredentialsPropertiesOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) TokenCredentialsProperties {
		if v != nil {
			return *v
		}
		var ret TokenCredentialsProperties
		return ret
	}).(TokenCredentialsPropertiesOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) Certificates() TokenCertificateArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) []TokenCertificate {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(TokenCertificateArrayOutput)
}

func (o TokenCredentialsPropertiesPtrOutput) Passwords() TokenPasswordArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsProperties) []TokenPassword {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(TokenPasswordArrayOutput)
}

type TokenCredentialsPropertiesResponse struct {
	Certificates []TokenCertificateResponse `pulumi:"certificates"`
	Passwords    []TokenPasswordResponse    `pulumi:"passwords"`
}





type TokenCredentialsPropertiesResponseInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput
	ToTokenCredentialsPropertiesResponseOutputWithContext(context.Context) TokenCredentialsPropertiesResponseOutput
}

type TokenCredentialsPropertiesResponseArgs struct {
	Certificates TokenCertificateResponseArrayInput `pulumi:"certificates"`
	Passwords    TokenPasswordResponseArrayInput    `pulumi:"passwords"`
}

func (TokenCredentialsPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput {
	return i.ToTokenCredentialsPropertiesResponseOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponseOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesResponseOutput)
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return i.ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TokenCredentialsPropertiesResponseArgs) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesResponseOutput).ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx)
}









type TokenCredentialsPropertiesResponsePtrInput interface {
	pulumi.Input

	ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput
	ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Context) TokenCredentialsPropertiesResponsePtrOutput
}

type tokenCredentialsPropertiesResponsePtrType TokenCredentialsPropertiesResponseArgs

func TokenCredentialsPropertiesResponsePtr(v *TokenCredentialsPropertiesResponseArgs) TokenCredentialsPropertiesResponsePtrInput {
	return (*tokenCredentialsPropertiesResponsePtrType)(v)
}

func (*tokenCredentialsPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (i *tokenCredentialsPropertiesResponsePtrType) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return i.ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *tokenCredentialsPropertiesResponsePtrType) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenCredentialsPropertiesResponsePtrOutput)
}

type TokenCredentialsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponseOutput() TokenCredentialsPropertiesResponseOutput {
	return o
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponseOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponseOutput {
	return o
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return o.ToTokenCredentialsPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TokenCredentialsPropertiesResponseOutput) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TokenCredentialsPropertiesResponse) *TokenCredentialsPropertiesResponse {
		return &v
	}).(TokenCredentialsPropertiesResponsePtrOutput)
}

func (o TokenCredentialsPropertiesResponseOutput) Certificates() TokenCertificateResponseArrayOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) []TokenCertificateResponse { return v.Certificates }).(TokenCertificateResponseArrayOutput)
}

func (o TokenCredentialsPropertiesResponseOutput) Passwords() TokenPasswordResponseArrayOutput {
	return o.ApplyT(func(v TokenCredentialsPropertiesResponse) []TokenPasswordResponse { return v.Passwords }).(TokenPasswordResponseArrayOutput)
}

type TokenCredentialsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TokenCredentialsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TokenCredentialsPropertiesResponse)(nil)).Elem()
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ToTokenCredentialsPropertiesResponsePtrOutput() TokenCredentialsPropertiesResponsePtrOutput {
	return o
}

func (o TokenCredentialsPropertiesResponsePtrOutput) ToTokenCredentialsPropertiesResponsePtrOutputWithContext(ctx context.Context) TokenCredentialsPropertiesResponsePtrOutput {
	return o
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Elem() TokenCredentialsPropertiesResponseOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) TokenCredentialsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TokenCredentialsPropertiesResponse
		return ret
	}).(TokenCredentialsPropertiesResponseOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Certificates() TokenCertificateResponseArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) []TokenCertificateResponse {
		if v == nil {
			return nil
		}
		return v.Certificates
	}).(TokenCertificateResponseArrayOutput)
}

func (o TokenCredentialsPropertiesResponsePtrOutput) Passwords() TokenPasswordResponseArrayOutput {
	return o.ApplyT(func(v *TokenCredentialsPropertiesResponse) []TokenPasswordResponse {
		if v == nil {
			return nil
		}
		return v.Passwords
	}).(TokenPasswordResponseArrayOutput)
}

type TokenPassword struct {
	CreationTime *string `pulumi:"creationTime"`
	Expiry       *string `pulumi:"expiry"`
	Name         *string `pulumi:"name"`
}





type TokenPasswordInput interface {
	pulumi.Input

	ToTokenPasswordOutput() TokenPasswordOutput
	ToTokenPasswordOutputWithContext(context.Context) TokenPasswordOutput
}

type TokenPasswordArgs struct {
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	Expiry       pulumi.StringPtrInput `pulumi:"expiry"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
}

func (TokenPasswordArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPassword)(nil)).Elem()
}

func (i TokenPasswordArgs) ToTokenPasswordOutput() TokenPasswordOutput {
	return i.ToTokenPasswordOutputWithContext(context.Background())
}

func (i TokenPasswordArgs) ToTokenPasswordOutputWithContext(ctx context.Context) TokenPasswordOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordOutput)
}





type TokenPasswordArrayInput interface {
	pulumi.Input

	ToTokenPasswordArrayOutput() TokenPasswordArrayOutput
	ToTokenPasswordArrayOutputWithContext(context.Context) TokenPasswordArrayOutput
}

type TokenPasswordArray []TokenPasswordInput

func (TokenPasswordArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPassword)(nil)).Elem()
}

func (i TokenPasswordArray) ToTokenPasswordArrayOutput() TokenPasswordArrayOutput {
	return i.ToTokenPasswordArrayOutputWithContext(context.Background())
}

func (i TokenPasswordArray) ToTokenPasswordArrayOutputWithContext(ctx context.Context) TokenPasswordArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordArrayOutput)
}

type TokenPasswordOutput struct{ *pulumi.OutputState }

func (TokenPasswordOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPassword)(nil)).Elem()
}

func (o TokenPasswordOutput) ToTokenPasswordOutput() TokenPasswordOutput {
	return o
}

func (o TokenPasswordOutput) ToTokenPasswordOutputWithContext(ctx context.Context) TokenPasswordOutput {
	return o
}

func (o TokenPasswordOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPassword) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type TokenPasswordArrayOutput struct{ *pulumi.OutputState }

func (TokenPasswordArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPassword)(nil)).Elem()
}

func (o TokenPasswordArrayOutput) ToTokenPasswordArrayOutput() TokenPasswordArrayOutput {
	return o
}

func (o TokenPasswordArrayOutput) ToTokenPasswordArrayOutputWithContext(ctx context.Context) TokenPasswordArrayOutput {
	return o
}

func (o TokenPasswordArrayOutput) Index(i pulumi.IntInput) TokenPasswordOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenPassword {
		return vs[0].([]TokenPassword)[vs[1].(int)]
	}).(TokenPasswordOutput)
}

type TokenPasswordResponse struct {
	CreationTime *string `pulumi:"creationTime"`
	Expiry       *string `pulumi:"expiry"`
	Name         *string `pulumi:"name"`
	Value        string  `pulumi:"value"`
}





type TokenPasswordResponseInput interface {
	pulumi.Input

	ToTokenPasswordResponseOutput() TokenPasswordResponseOutput
	ToTokenPasswordResponseOutputWithContext(context.Context) TokenPasswordResponseOutput
}

type TokenPasswordResponseArgs struct {
	CreationTime pulumi.StringPtrInput `pulumi:"creationTime"`
	Expiry       pulumi.StringPtrInput `pulumi:"expiry"`
	Name         pulumi.StringPtrInput `pulumi:"name"`
	Value        pulumi.StringInput    `pulumi:"value"`
}

func (TokenPasswordResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordResponse)(nil)).Elem()
}

func (i TokenPasswordResponseArgs) ToTokenPasswordResponseOutput() TokenPasswordResponseOutput {
	return i.ToTokenPasswordResponseOutputWithContext(context.Background())
}

func (i TokenPasswordResponseArgs) ToTokenPasswordResponseOutputWithContext(ctx context.Context) TokenPasswordResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordResponseOutput)
}





type TokenPasswordResponseArrayInput interface {
	pulumi.Input

	ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput
	ToTokenPasswordResponseArrayOutputWithContext(context.Context) TokenPasswordResponseArrayOutput
}

type TokenPasswordResponseArray []TokenPasswordResponseInput

func (TokenPasswordResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPasswordResponse)(nil)).Elem()
}

func (i TokenPasswordResponseArray) ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput {
	return i.ToTokenPasswordResponseArrayOutputWithContext(context.Background())
}

func (i TokenPasswordResponseArray) ToTokenPasswordResponseArrayOutputWithContext(ctx context.Context) TokenPasswordResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TokenPasswordResponseArrayOutput)
}

type TokenPasswordResponseOutput struct{ *pulumi.OutputState }

func (TokenPasswordResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TokenPasswordResponse)(nil)).Elem()
}

func (o TokenPasswordResponseOutput) ToTokenPasswordResponseOutput() TokenPasswordResponseOutput {
	return o
}

func (o TokenPasswordResponseOutput) ToTokenPasswordResponseOutputWithContext(ctx context.Context) TokenPasswordResponseOutput {
	return o
}

func (o TokenPasswordResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Expiry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.Expiry }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TokenPasswordResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o TokenPasswordResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v TokenPasswordResponse) string { return v.Value }).(pulumi.StringOutput)
}

type TokenPasswordResponseArrayOutput struct{ *pulumi.OutputState }

func (TokenPasswordResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TokenPasswordResponse)(nil)).Elem()
}

func (o TokenPasswordResponseArrayOutput) ToTokenPasswordResponseArrayOutput() TokenPasswordResponseArrayOutput {
	return o
}

func (o TokenPasswordResponseArrayOutput) ToTokenPasswordResponseArrayOutputWithContext(ctx context.Context) TokenPasswordResponseArrayOutput {
	return o
}

func (o TokenPasswordResponseArrayOutput) Index(i pulumi.IntInput) TokenPasswordResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TokenPasswordResponse {
		return vs[0].([]TokenPasswordResponse)[vs[1].(int)]
	}).(TokenPasswordResponseOutput)
}

type TrustPolicy struct {
	Status *string `pulumi:"status"`
	Type   *string `pulumi:"type"`
}


func (val *TrustPolicy) Defaults() *TrustPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	if isZero(tmp.Type) {
		type_ := "Notary"
		tmp.Type = &type_
	}
	return &tmp
}





type TrustPolicyInput interface {
	pulumi.Input

	ToTrustPolicyOutput() TrustPolicyOutput
	ToTrustPolicyOutputWithContext(context.Context) TrustPolicyOutput
}

type TrustPolicyArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
}

func (TrustPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicy)(nil)).Elem()
}

func (i TrustPolicyArgs) ToTrustPolicyOutput() TrustPolicyOutput {
	return i.ToTrustPolicyOutputWithContext(context.Background())
}

func (i TrustPolicyArgs) ToTrustPolicyOutputWithContext(ctx context.Context) TrustPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyOutput)
}

func (i TrustPolicyArgs) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return i.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (i TrustPolicyArgs) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyOutput).ToTrustPolicyPtrOutputWithContext(ctx)
}









type TrustPolicyPtrInput interface {
	pulumi.Input

	ToTrustPolicyPtrOutput() TrustPolicyPtrOutput
	ToTrustPolicyPtrOutputWithContext(context.Context) TrustPolicyPtrOutput
}

type trustPolicyPtrType TrustPolicyArgs

func TrustPolicyPtr(v *TrustPolicyArgs) TrustPolicyPtrInput {
	return (*trustPolicyPtrType)(v)
}

func (*trustPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicy)(nil)).Elem()
}

func (i *trustPolicyPtrType) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return i.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (i *trustPolicyPtrType) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyPtrOutput)
}

type TrustPolicyOutput struct{ *pulumi.OutputState }

func (TrustPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicy)(nil)).Elem()
}

func (o TrustPolicyOutput) ToTrustPolicyOutput() TrustPolicyOutput {
	return o
}

func (o TrustPolicyOutput) ToTrustPolicyOutputWithContext(ctx context.Context) TrustPolicyOutput {
	return o
}

func (o TrustPolicyOutput) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return o.ToTrustPolicyPtrOutputWithContext(context.Background())
}

func (o TrustPolicyOutput) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustPolicy) *TrustPolicy {
		return &v
	}).(TrustPolicyPtrOutput)
}

func (o TrustPolicyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicy) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrustPolicyOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicy) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TrustPolicyPtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicy)(nil)).Elem()
}

func (o TrustPolicyPtrOutput) ToTrustPolicyPtrOutput() TrustPolicyPtrOutput {
	return o
}

func (o TrustPolicyPtrOutput) ToTrustPolicyPtrOutputWithContext(ctx context.Context) TrustPolicyPtrOutput {
	return o
}

func (o TrustPolicyPtrOutput) Elem() TrustPolicyOutput {
	return o.ApplyT(func(v *TrustPolicy) TrustPolicy {
		if v != nil {
			return *v
		}
		var ret TrustPolicy
		return ret
	}).(TrustPolicyOutput)
}

func (o TrustPolicyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o TrustPolicyPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type TrustPolicyResponse struct {
	Status *string `pulumi:"status"`
	Type   *string `pulumi:"type"`
}


func (val *TrustPolicyResponse) Defaults() *TrustPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Status) {
		status_ := "disabled"
		tmp.Status = &status_
	}
	if isZero(tmp.Type) {
		type_ := "Notary"
		tmp.Type = &type_
	}
	return &tmp
}





type TrustPolicyResponseInput interface {
	pulumi.Input

	ToTrustPolicyResponseOutput() TrustPolicyResponseOutput
	ToTrustPolicyResponseOutputWithContext(context.Context) TrustPolicyResponseOutput
}

type TrustPolicyResponseArgs struct {
	Status pulumi.StringPtrInput `pulumi:"status"`
	Type   pulumi.StringPtrInput `pulumi:"type"`
}

func (TrustPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicyResponse)(nil)).Elem()
}

func (i TrustPolicyResponseArgs) ToTrustPolicyResponseOutput() TrustPolicyResponseOutput {
	return i.ToTrustPolicyResponseOutputWithContext(context.Background())
}

func (i TrustPolicyResponseArgs) ToTrustPolicyResponseOutputWithContext(ctx context.Context) TrustPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyResponseOutput)
}

func (i TrustPolicyResponseArgs) ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput {
	return i.ToTrustPolicyResponsePtrOutputWithContext(context.Background())
}

func (i TrustPolicyResponseArgs) ToTrustPolicyResponsePtrOutputWithContext(ctx context.Context) TrustPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyResponseOutput).ToTrustPolicyResponsePtrOutputWithContext(ctx)
}









type TrustPolicyResponsePtrInput interface {
	pulumi.Input

	ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput
	ToTrustPolicyResponsePtrOutputWithContext(context.Context) TrustPolicyResponsePtrOutput
}

type trustPolicyResponsePtrType TrustPolicyResponseArgs

func TrustPolicyResponsePtr(v *TrustPolicyResponseArgs) TrustPolicyResponsePtrInput {
	return (*trustPolicyResponsePtrType)(v)
}

func (*trustPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicyResponse)(nil)).Elem()
}

func (i *trustPolicyResponsePtrType) ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput {
	return i.ToTrustPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *trustPolicyResponsePtrType) ToTrustPolicyResponsePtrOutputWithContext(ctx context.Context) TrustPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrustPolicyResponsePtrOutput)
}

type TrustPolicyResponseOutput struct{ *pulumi.OutputState }

func (TrustPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrustPolicyResponse)(nil)).Elem()
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponseOutput() TrustPolicyResponseOutput {
	return o
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponseOutputWithContext(ctx context.Context) TrustPolicyResponseOutput {
	return o
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput {
	return o.ToTrustPolicyResponsePtrOutputWithContext(context.Background())
}

func (o TrustPolicyResponseOutput) ToTrustPolicyResponsePtrOutputWithContext(ctx context.Context) TrustPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TrustPolicyResponse) *TrustPolicyResponse {
		return &v
	}).(TrustPolicyResponsePtrOutput)
}

func (o TrustPolicyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TrustPolicyResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrustPolicyResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type TrustPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (TrustPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TrustPolicyResponse)(nil)).Elem()
}

func (o TrustPolicyResponsePtrOutput) ToTrustPolicyResponsePtrOutput() TrustPolicyResponsePtrOutput {
	return o
}

func (o TrustPolicyResponsePtrOutput) ToTrustPolicyResponsePtrOutputWithContext(ctx context.Context) TrustPolicyResponsePtrOutput {
	return o
}

func (o TrustPolicyResponsePtrOutput) Elem() TrustPolicyResponseOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) TrustPolicyResponse {
		if v != nil {
			return *v
		}
		var ret TrustPolicyResponse
		return ret
	}).(TrustPolicyResponseOutput)
}

func (o TrustPolicyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o TrustPolicyResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TrustPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type UserIdentityProperties struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesInput interface {
	pulumi.Input

	ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput
	ToUserIdentityPropertiesOutputWithContext(context.Context) UserIdentityPropertiesOutput
}

type UserIdentityPropertiesArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return i.ToUserIdentityPropertiesOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesArgs) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesOutput)
}





type UserIdentityPropertiesMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput
	ToUserIdentityPropertiesMapOutputWithContext(context.Context) UserIdentityPropertiesMapOutput
}

type UserIdentityPropertiesMap map[string]UserIdentityPropertiesInput

func (UserIdentityPropertiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return i.ToUserIdentityPropertiesMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesMap) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesMapOutput)
}

type UserIdentityPropertiesOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutput() UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ToUserIdentityPropertiesOutputWithContext(ctx context.Context) UserIdentityPropertiesOutput {
	return o
}

func (o UserIdentityPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityProperties) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityProperties)(nil)).Elem()
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutput() UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) ToUserIdentityPropertiesMapOutputWithContext(ctx context.Context) UserIdentityPropertiesMapOutput {
	return o
}

func (o UserIdentityPropertiesMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityProperties {
		return vs[0].(map[string]UserIdentityProperties)[vs[1].(string)]
	}).(UserIdentityPropertiesOutput)
}

type UserIdentityPropertiesResponse struct {
	ClientId    *string `pulumi:"clientId"`
	PrincipalId *string `pulumi:"principalId"`
}





type UserIdentityPropertiesResponseInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput
	ToUserIdentityPropertiesResponseOutputWithContext(context.Context) UserIdentityPropertiesResponseOutput
}

type UserIdentityPropertiesResponseArgs struct {
	ClientId    pulumi.StringPtrInput `pulumi:"clientId"`
	PrincipalId pulumi.StringPtrInput `pulumi:"principalId"`
}

func (UserIdentityPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return i.ToUserIdentityPropertiesResponseOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseArgs) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseOutput)
}





type UserIdentityPropertiesResponseMapInput interface {
	pulumi.Input

	ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput
	ToUserIdentityPropertiesResponseMapOutputWithContext(context.Context) UserIdentityPropertiesResponseMapOutput
}

type UserIdentityPropertiesResponseMap map[string]UserIdentityPropertiesResponseInput

func (UserIdentityPropertiesResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return i.ToUserIdentityPropertiesResponseMapOutputWithContext(context.Background())
}

func (i UserIdentityPropertiesResponseMap) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UserIdentityPropertiesResponseMapOutput)
}

type UserIdentityPropertiesResponseOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutput() UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ToUserIdentityPropertiesResponseOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseOutput {
	return o
}

func (o UserIdentityPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o UserIdentityPropertiesResponseOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UserIdentityPropertiesResponse) *string { return v.PrincipalId }).(pulumi.StringPtrOutput)
}

type UserIdentityPropertiesResponseMapOutput struct{ *pulumi.OutputState }

func (UserIdentityPropertiesResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserIdentityPropertiesResponse)(nil)).Elem()
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutput() UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) ToUserIdentityPropertiesResponseMapOutputWithContext(ctx context.Context) UserIdentityPropertiesResponseMapOutput {
	return o
}

func (o UserIdentityPropertiesResponseMapOutput) MapIndex(k pulumi.StringInput) UserIdentityPropertiesResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserIdentityPropertiesResponse {
		return vs[0].(map[string]UserIdentityPropertiesResponse)[vs[1].(string)]
	}).(UserIdentityPropertiesResponseOutput)
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
	pulumi.RegisterOutputType(ActivationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ActivationPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ActorResponseOutput{})
	pulumi.RegisterOutputType(ActorResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(EventContentResponseOutput{})
	pulumi.RegisterOutputType(EventContentResponsePtrOutput{})
	pulumi.RegisterOutputType(EventRequestMessageResponseOutput{})
	pulumi.RegisterOutputType(EventRequestMessageResponsePtrOutput{})
	pulumi.RegisterOutputType(EventResponseOutput{})
	pulumi.RegisterOutputType(EventResponseArrayOutput{})
	pulumi.RegisterOutputType(EventResponseMessageResponseOutput{})
	pulumi.RegisterOutputType(EventResponseMessageResponsePtrOutput{})
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesOutput{})
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ExportPipelineTargetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(IPRuleOutput{})
	pulumi.RegisterOutputType(IPRuleArrayOutput{})
	pulumi.RegisterOutputType(IPRuleResponseOutput{})
	pulumi.RegisterOutputType(IPRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesPtrOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(IdentityPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ImportPipelineSourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LoggingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LoginServerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LoginServerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetPtrOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponseOutput{})
	pulumi.RegisterOutputType(NetworkRuleSetResponsePtrOutput{})
	pulumi.RegisterOutputType(ParentPropertiesOutput{})
	pulumi.RegisterOutputType(ParentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ParentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ParentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunRequestResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunResponseResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunSourcePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineRunTargetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineSourceTriggerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineTriggerDescriptorResponseOutput{})
	pulumi.RegisterOutputType(PipelineTriggerDescriptorResponsePtrOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesPtrOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PipelineTriggerPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PoliciesOutput{})
	pulumi.RegisterOutputType(PoliciesPtrOutput{})
	pulumi.RegisterOutputType(PoliciesResponseOutput{})
	pulumi.RegisterOutputType(PoliciesResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ProgressPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ProgressPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyPtrOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyResponseOutput{})
	pulumi.RegisterOutputType(QuarantinePolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistryPasswordResponseOutput{})
	pulumi.RegisterOutputType(RegistryPasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(RequestResponseOutput{})
	pulumi.RegisterOutputType(RequestResponsePtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceResponseOutput{})
	pulumi.RegisterOutputType(SourceResponsePtrOutput{})
	pulumi.RegisterOutputType(StatusDetailPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StatusDetailPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusResponseOutput{})
	pulumi.RegisterOutputType(StatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SyncPropertiesOutput{})
	pulumi.RegisterOutputType(SyncPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SyncPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SyncPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetResponseOutput{})
	pulumi.RegisterOutputType(TargetResponsePtrOutput{})
	pulumi.RegisterOutputType(TlsCertificatePropertiesResponseOutput{})
	pulumi.RegisterOutputType(TlsCertificatePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TlsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TlsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenCertificateOutput{})
	pulumi.RegisterOutputType(TokenCertificateArrayOutput{})
	pulumi.RegisterOutputType(TokenCertificateResponseOutput{})
	pulumi.RegisterOutputType(TokenCertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TokenCredentialsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TokenPasswordOutput{})
	pulumi.RegisterOutputType(TokenPasswordArrayOutput{})
	pulumi.RegisterOutputType(TokenPasswordResponseOutput{})
	pulumi.RegisterOutputType(TokenPasswordResponseArrayOutput{})
	pulumi.RegisterOutputType(TrustPolicyOutput{})
	pulumi.RegisterOutputType(TrustPolicyPtrOutput{})
	pulumi.RegisterOutputType(TrustPolicyResponseOutput{})
	pulumi.RegisterOutputType(TrustPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesMapOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseOutput{})
	pulumi.RegisterOutputType(UserIdentityPropertiesResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
}
