


package v20160301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CompatibilityLevel string

const (
	CompatibilityLevel_1_0 = CompatibilityLevel("1.0")
)

func (CompatibilityLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*CompatibilityLevel)(nil)).Elem()
}

func (e CompatibilityLevel) ToCompatibilityLevelOutput() CompatibilityLevelOutput {
	return pulumi.ToOutput(e).(CompatibilityLevelOutput)
}

func (e CompatibilityLevel) ToCompatibilityLevelOutputWithContext(ctx context.Context) CompatibilityLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CompatibilityLevelOutput)
}

func (e CompatibilityLevel) ToCompatibilityLevelPtrOutput() CompatibilityLevelPtrOutput {
	return e.ToCompatibilityLevelPtrOutputWithContext(context.Background())
}

func (e CompatibilityLevel) ToCompatibilityLevelPtrOutputWithContext(ctx context.Context) CompatibilityLevelPtrOutput {
	return CompatibilityLevel(e).ToCompatibilityLevelOutputWithContext(ctx).ToCompatibilityLevelPtrOutputWithContext(ctx)
}

func (e CompatibilityLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompatibilityLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CompatibilityLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CompatibilityLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CompatibilityLevelOutput struct{ *pulumi.OutputState }

func (CompatibilityLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompatibilityLevel)(nil)).Elem()
}

func (o CompatibilityLevelOutput) ToCompatibilityLevelOutput() CompatibilityLevelOutput {
	return o
}

func (o CompatibilityLevelOutput) ToCompatibilityLevelOutputWithContext(ctx context.Context) CompatibilityLevelOutput {
	return o
}

func (o CompatibilityLevelOutput) ToCompatibilityLevelPtrOutput() CompatibilityLevelPtrOutput {
	return o.ToCompatibilityLevelPtrOutputWithContext(context.Background())
}

func (o CompatibilityLevelOutput) ToCompatibilityLevelPtrOutputWithContext(ctx context.Context) CompatibilityLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CompatibilityLevel) *CompatibilityLevel {
		return &v
	}).(CompatibilityLevelPtrOutput)
}

func (o CompatibilityLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CompatibilityLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CompatibilityLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CompatibilityLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CompatibilityLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CompatibilityLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CompatibilityLevelPtrOutput struct{ *pulumi.OutputState }

func (CompatibilityLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CompatibilityLevel)(nil)).Elem()
}

func (o CompatibilityLevelPtrOutput) ToCompatibilityLevelPtrOutput() CompatibilityLevelPtrOutput {
	return o
}

func (o CompatibilityLevelPtrOutput) ToCompatibilityLevelPtrOutputWithContext(ctx context.Context) CompatibilityLevelPtrOutput {
	return o
}

func (o CompatibilityLevelPtrOutput) Elem() CompatibilityLevelOutput {
	return o.ApplyT(func(v *CompatibilityLevel) CompatibilityLevel {
		if v != nil {
			return *v
		}
		var ret CompatibilityLevel
		return ret
	}).(CompatibilityLevelOutput)
}

func (o CompatibilityLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CompatibilityLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CompatibilityLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CompatibilityLevelInput interface {
	pulumi.Input

	ToCompatibilityLevelOutput() CompatibilityLevelOutput
	ToCompatibilityLevelOutputWithContext(context.Context) CompatibilityLevelOutput
}

var compatibilityLevelPtrType = reflect.TypeOf((**CompatibilityLevel)(nil)).Elem()

type CompatibilityLevelPtrInput interface {
	pulumi.Input

	ToCompatibilityLevelPtrOutput() CompatibilityLevelPtrOutput
	ToCompatibilityLevelPtrOutputWithContext(context.Context) CompatibilityLevelPtrOutput
}

type compatibilityLevelPtr string

func CompatibilityLevelPtr(v string) CompatibilityLevelPtrInput {
	return (*compatibilityLevelPtr)(&v)
}

func (*compatibilityLevelPtr) ElementType() reflect.Type {
	return compatibilityLevelPtrType
}

func (in *compatibilityLevelPtr) ToCompatibilityLevelPtrOutput() CompatibilityLevelPtrOutput {
	return pulumi.ToOutput(in).(CompatibilityLevelPtrOutput)
}

func (in *compatibilityLevelPtr) ToCompatibilityLevelPtrOutputWithContext(ctx context.Context) CompatibilityLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CompatibilityLevelPtrOutput)
}

type Encoding string

const (
	EncodingUTF8 = Encoding("UTF8")
)

func (Encoding) ElementType() reflect.Type {
	return reflect.TypeOf((*Encoding)(nil)).Elem()
}

func (e Encoding) ToEncodingOutput() EncodingOutput {
	return pulumi.ToOutput(e).(EncodingOutput)
}

func (e Encoding) ToEncodingOutputWithContext(ctx context.Context) EncodingOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EncodingOutput)
}

func (e Encoding) ToEncodingPtrOutput() EncodingPtrOutput {
	return e.ToEncodingPtrOutputWithContext(context.Background())
}

func (e Encoding) ToEncodingPtrOutputWithContext(ctx context.Context) EncodingPtrOutput {
	return Encoding(e).ToEncodingOutputWithContext(ctx).ToEncodingPtrOutputWithContext(ctx)
}

func (e Encoding) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Encoding) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Encoding) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Encoding) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EncodingOutput struct{ *pulumi.OutputState }

func (EncodingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Encoding)(nil)).Elem()
}

func (o EncodingOutput) ToEncodingOutput() EncodingOutput {
	return o
}

func (o EncodingOutput) ToEncodingOutputWithContext(ctx context.Context) EncodingOutput {
	return o
}

func (o EncodingOutput) ToEncodingPtrOutput() EncodingPtrOutput {
	return o.ToEncodingPtrOutputWithContext(context.Background())
}

func (o EncodingOutput) ToEncodingPtrOutputWithContext(ctx context.Context) EncodingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Encoding) *Encoding {
		return &v
	}).(EncodingPtrOutput)
}

func (o EncodingOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EncodingOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Encoding) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EncodingOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncodingOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Encoding) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EncodingPtrOutput struct{ *pulumi.OutputState }

func (EncodingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Encoding)(nil)).Elem()
}

func (o EncodingPtrOutput) ToEncodingPtrOutput() EncodingPtrOutput {
	return o
}

func (o EncodingPtrOutput) ToEncodingPtrOutputWithContext(ctx context.Context) EncodingPtrOutput {
	return o
}

func (o EncodingPtrOutput) Elem() EncodingOutput {
	return o.ApplyT(func(v *Encoding) Encoding {
		if v != nil {
			return *v
		}
		var ret Encoding
		return ret
	}).(EncodingOutput)
}

func (o EncodingPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EncodingPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Encoding) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EncodingInput interface {
	pulumi.Input

	ToEncodingOutput() EncodingOutput
	ToEncodingOutputWithContext(context.Context) EncodingOutput
}

var encodingPtrType = reflect.TypeOf((**Encoding)(nil)).Elem()

type EncodingPtrInput interface {
	pulumi.Input

	ToEncodingPtrOutput() EncodingPtrOutput
	ToEncodingPtrOutputWithContext(context.Context) EncodingPtrOutput
}

type encodingPtr string

func EncodingPtr(v string) EncodingPtrInput {
	return (*encodingPtr)(&v)
}

func (*encodingPtr) ElementType() reflect.Type {
	return encodingPtrType
}

func (in *encodingPtr) ToEncodingPtrOutput() EncodingPtrOutput {
	return pulumi.ToOutput(in).(EncodingPtrOutput)
}

func (in *encodingPtr) ToEncodingPtrOutputWithContext(ctx context.Context) EncodingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EncodingPtrOutput)
}

type EventSerializationType string

const (
	EventSerializationTypeCsv  = EventSerializationType("Csv")
	EventSerializationTypeAvro = EventSerializationType("Avro")
	EventSerializationTypeJson = EventSerializationType("Json")
)

func (EventSerializationType) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSerializationType)(nil)).Elem()
}

func (e EventSerializationType) ToEventSerializationTypeOutput() EventSerializationTypeOutput {
	return pulumi.ToOutput(e).(EventSerializationTypeOutput)
}

func (e EventSerializationType) ToEventSerializationTypeOutputWithContext(ctx context.Context) EventSerializationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSerializationTypeOutput)
}

func (e EventSerializationType) ToEventSerializationTypePtrOutput() EventSerializationTypePtrOutput {
	return e.ToEventSerializationTypePtrOutputWithContext(context.Background())
}

func (e EventSerializationType) ToEventSerializationTypePtrOutputWithContext(ctx context.Context) EventSerializationTypePtrOutput {
	return EventSerializationType(e).ToEventSerializationTypeOutputWithContext(ctx).ToEventSerializationTypePtrOutputWithContext(ctx)
}

func (e EventSerializationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSerializationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSerializationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSerializationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSerializationTypeOutput struct{ *pulumi.OutputState }

func (EventSerializationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSerializationType)(nil)).Elem()
}

func (o EventSerializationTypeOutput) ToEventSerializationTypeOutput() EventSerializationTypeOutput {
	return o
}

func (o EventSerializationTypeOutput) ToEventSerializationTypeOutputWithContext(ctx context.Context) EventSerializationTypeOutput {
	return o
}

func (o EventSerializationTypeOutput) ToEventSerializationTypePtrOutput() EventSerializationTypePtrOutput {
	return o.ToEventSerializationTypePtrOutputWithContext(context.Background())
}

func (o EventSerializationTypeOutput) ToEventSerializationTypePtrOutputWithContext(ctx context.Context) EventSerializationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSerializationType) *EventSerializationType {
		return &v
	}).(EventSerializationTypePtrOutput)
}

func (o EventSerializationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSerializationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSerializationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSerializationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSerializationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSerializationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSerializationTypePtrOutput struct{ *pulumi.OutputState }

func (EventSerializationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSerializationType)(nil)).Elem()
}

func (o EventSerializationTypePtrOutput) ToEventSerializationTypePtrOutput() EventSerializationTypePtrOutput {
	return o
}

func (o EventSerializationTypePtrOutput) ToEventSerializationTypePtrOutputWithContext(ctx context.Context) EventSerializationTypePtrOutput {
	return o
}

func (o EventSerializationTypePtrOutput) Elem() EventSerializationTypeOutput {
	return o.ApplyT(func(v *EventSerializationType) EventSerializationType {
		if v != nil {
			return *v
		}
		var ret EventSerializationType
		return ret
	}).(EventSerializationTypeOutput)
}

func (o EventSerializationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSerializationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSerializationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventSerializationTypeInput interface {
	pulumi.Input

	ToEventSerializationTypeOutput() EventSerializationTypeOutput
	ToEventSerializationTypeOutputWithContext(context.Context) EventSerializationTypeOutput
}

var eventSerializationTypePtrType = reflect.TypeOf((**EventSerializationType)(nil)).Elem()

type EventSerializationTypePtrInput interface {
	pulumi.Input

	ToEventSerializationTypePtrOutput() EventSerializationTypePtrOutput
	ToEventSerializationTypePtrOutputWithContext(context.Context) EventSerializationTypePtrOutput
}

type eventSerializationTypePtr string

func EventSerializationTypePtr(v string) EventSerializationTypePtrInput {
	return (*eventSerializationTypePtr)(&v)
}

func (*eventSerializationTypePtr) ElementType() reflect.Type {
	return eventSerializationTypePtrType
}

func (in *eventSerializationTypePtr) ToEventSerializationTypePtrOutput() EventSerializationTypePtrOutput {
	return pulumi.ToOutput(in).(EventSerializationTypePtrOutput)
}

func (in *eventSerializationTypePtr) ToEventSerializationTypePtrOutputWithContext(ctx context.Context) EventSerializationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSerializationTypePtrOutput)
}

type EventsOutOfOrderPolicy string

const (
	EventsOutOfOrderPolicyAdjust = EventsOutOfOrderPolicy("Adjust")
	EventsOutOfOrderPolicyDrop   = EventsOutOfOrderPolicy("Drop")
)

func (EventsOutOfOrderPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*EventsOutOfOrderPolicy)(nil)).Elem()
}

func (e EventsOutOfOrderPolicy) ToEventsOutOfOrderPolicyOutput() EventsOutOfOrderPolicyOutput {
	return pulumi.ToOutput(e).(EventsOutOfOrderPolicyOutput)
}

func (e EventsOutOfOrderPolicy) ToEventsOutOfOrderPolicyOutputWithContext(ctx context.Context) EventsOutOfOrderPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventsOutOfOrderPolicyOutput)
}

func (e EventsOutOfOrderPolicy) ToEventsOutOfOrderPolicyPtrOutput() EventsOutOfOrderPolicyPtrOutput {
	return e.ToEventsOutOfOrderPolicyPtrOutputWithContext(context.Background())
}

func (e EventsOutOfOrderPolicy) ToEventsOutOfOrderPolicyPtrOutputWithContext(ctx context.Context) EventsOutOfOrderPolicyPtrOutput {
	return EventsOutOfOrderPolicy(e).ToEventsOutOfOrderPolicyOutputWithContext(ctx).ToEventsOutOfOrderPolicyPtrOutputWithContext(ctx)
}

func (e EventsOutOfOrderPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventsOutOfOrderPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventsOutOfOrderPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventsOutOfOrderPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventsOutOfOrderPolicyOutput struct{ *pulumi.OutputState }

func (EventsOutOfOrderPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventsOutOfOrderPolicy)(nil)).Elem()
}

func (o EventsOutOfOrderPolicyOutput) ToEventsOutOfOrderPolicyOutput() EventsOutOfOrderPolicyOutput {
	return o
}

func (o EventsOutOfOrderPolicyOutput) ToEventsOutOfOrderPolicyOutputWithContext(ctx context.Context) EventsOutOfOrderPolicyOutput {
	return o
}

func (o EventsOutOfOrderPolicyOutput) ToEventsOutOfOrderPolicyPtrOutput() EventsOutOfOrderPolicyPtrOutput {
	return o.ToEventsOutOfOrderPolicyPtrOutputWithContext(context.Background())
}

func (o EventsOutOfOrderPolicyOutput) ToEventsOutOfOrderPolicyPtrOutputWithContext(ctx context.Context) EventsOutOfOrderPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventsOutOfOrderPolicy) *EventsOutOfOrderPolicy {
		return &v
	}).(EventsOutOfOrderPolicyPtrOutput)
}

func (o EventsOutOfOrderPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventsOutOfOrderPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventsOutOfOrderPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventsOutOfOrderPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventsOutOfOrderPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventsOutOfOrderPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventsOutOfOrderPolicyPtrOutput struct{ *pulumi.OutputState }

func (EventsOutOfOrderPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventsOutOfOrderPolicy)(nil)).Elem()
}

func (o EventsOutOfOrderPolicyPtrOutput) ToEventsOutOfOrderPolicyPtrOutput() EventsOutOfOrderPolicyPtrOutput {
	return o
}

func (o EventsOutOfOrderPolicyPtrOutput) ToEventsOutOfOrderPolicyPtrOutputWithContext(ctx context.Context) EventsOutOfOrderPolicyPtrOutput {
	return o
}

func (o EventsOutOfOrderPolicyPtrOutput) Elem() EventsOutOfOrderPolicyOutput {
	return o.ApplyT(func(v *EventsOutOfOrderPolicy) EventsOutOfOrderPolicy {
		if v != nil {
			return *v
		}
		var ret EventsOutOfOrderPolicy
		return ret
	}).(EventsOutOfOrderPolicyOutput)
}

func (o EventsOutOfOrderPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventsOutOfOrderPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventsOutOfOrderPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventsOutOfOrderPolicyInput interface {
	pulumi.Input

	ToEventsOutOfOrderPolicyOutput() EventsOutOfOrderPolicyOutput
	ToEventsOutOfOrderPolicyOutputWithContext(context.Context) EventsOutOfOrderPolicyOutput
}

var eventsOutOfOrderPolicyPtrType = reflect.TypeOf((**EventsOutOfOrderPolicy)(nil)).Elem()

type EventsOutOfOrderPolicyPtrInput interface {
	pulumi.Input

	ToEventsOutOfOrderPolicyPtrOutput() EventsOutOfOrderPolicyPtrOutput
	ToEventsOutOfOrderPolicyPtrOutputWithContext(context.Context) EventsOutOfOrderPolicyPtrOutput
}

type eventsOutOfOrderPolicyPtr string

func EventsOutOfOrderPolicyPtr(v string) EventsOutOfOrderPolicyPtrInput {
	return (*eventsOutOfOrderPolicyPtr)(&v)
}

func (*eventsOutOfOrderPolicyPtr) ElementType() reflect.Type {
	return eventsOutOfOrderPolicyPtrType
}

func (in *eventsOutOfOrderPolicyPtr) ToEventsOutOfOrderPolicyPtrOutput() EventsOutOfOrderPolicyPtrOutput {
	return pulumi.ToOutput(in).(EventsOutOfOrderPolicyPtrOutput)
}

func (in *eventsOutOfOrderPolicyPtr) ToEventsOutOfOrderPolicyPtrOutputWithContext(ctx context.Context) EventsOutOfOrderPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventsOutOfOrderPolicyPtrOutput)
}

type JsonOutputSerializationFormat string

const (
	JsonOutputSerializationFormatLineSeparated = JsonOutputSerializationFormat("LineSeparated")
	JsonOutputSerializationFormatArray         = JsonOutputSerializationFormat("Array")
)

func (JsonOutputSerializationFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonOutputSerializationFormat)(nil)).Elem()
}

func (e JsonOutputSerializationFormat) ToJsonOutputSerializationFormatOutput() JsonOutputSerializationFormatOutput {
	return pulumi.ToOutput(e).(JsonOutputSerializationFormatOutput)
}

func (e JsonOutputSerializationFormat) ToJsonOutputSerializationFormatOutputWithContext(ctx context.Context) JsonOutputSerializationFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(JsonOutputSerializationFormatOutput)
}

func (e JsonOutputSerializationFormat) ToJsonOutputSerializationFormatPtrOutput() JsonOutputSerializationFormatPtrOutput {
	return e.ToJsonOutputSerializationFormatPtrOutputWithContext(context.Background())
}

func (e JsonOutputSerializationFormat) ToJsonOutputSerializationFormatPtrOutputWithContext(ctx context.Context) JsonOutputSerializationFormatPtrOutput {
	return JsonOutputSerializationFormat(e).ToJsonOutputSerializationFormatOutputWithContext(ctx).ToJsonOutputSerializationFormatPtrOutputWithContext(ctx)
}

func (e JsonOutputSerializationFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e JsonOutputSerializationFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e JsonOutputSerializationFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e JsonOutputSerializationFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type JsonOutputSerializationFormatOutput struct{ *pulumi.OutputState }

func (JsonOutputSerializationFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JsonOutputSerializationFormat)(nil)).Elem()
}

func (o JsonOutputSerializationFormatOutput) ToJsonOutputSerializationFormatOutput() JsonOutputSerializationFormatOutput {
	return o
}

func (o JsonOutputSerializationFormatOutput) ToJsonOutputSerializationFormatOutputWithContext(ctx context.Context) JsonOutputSerializationFormatOutput {
	return o
}

func (o JsonOutputSerializationFormatOutput) ToJsonOutputSerializationFormatPtrOutput() JsonOutputSerializationFormatPtrOutput {
	return o.ToJsonOutputSerializationFormatPtrOutputWithContext(context.Background())
}

func (o JsonOutputSerializationFormatOutput) ToJsonOutputSerializationFormatPtrOutputWithContext(ctx context.Context) JsonOutputSerializationFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v JsonOutputSerializationFormat) *JsonOutputSerializationFormat {
		return &v
	}).(JsonOutputSerializationFormatPtrOutput)
}

func (o JsonOutputSerializationFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o JsonOutputSerializationFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JsonOutputSerializationFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o JsonOutputSerializationFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JsonOutputSerializationFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e JsonOutputSerializationFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type JsonOutputSerializationFormatPtrOutput struct{ *pulumi.OutputState }

func (JsonOutputSerializationFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JsonOutputSerializationFormat)(nil)).Elem()
}

func (o JsonOutputSerializationFormatPtrOutput) ToJsonOutputSerializationFormatPtrOutput() JsonOutputSerializationFormatPtrOutput {
	return o
}

func (o JsonOutputSerializationFormatPtrOutput) ToJsonOutputSerializationFormatPtrOutputWithContext(ctx context.Context) JsonOutputSerializationFormatPtrOutput {
	return o
}

func (o JsonOutputSerializationFormatPtrOutput) Elem() JsonOutputSerializationFormatOutput {
	return o.ApplyT(func(v *JsonOutputSerializationFormat) JsonOutputSerializationFormat {
		if v != nil {
			return *v
		}
		var ret JsonOutputSerializationFormat
		return ret
	}).(JsonOutputSerializationFormatOutput)
}

func (o JsonOutputSerializationFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o JsonOutputSerializationFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *JsonOutputSerializationFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type JsonOutputSerializationFormatInput interface {
	pulumi.Input

	ToJsonOutputSerializationFormatOutput() JsonOutputSerializationFormatOutput
	ToJsonOutputSerializationFormatOutputWithContext(context.Context) JsonOutputSerializationFormatOutput
}

var jsonOutputSerializationFormatPtrType = reflect.TypeOf((**JsonOutputSerializationFormat)(nil)).Elem()

type JsonOutputSerializationFormatPtrInput interface {
	pulumi.Input

	ToJsonOutputSerializationFormatPtrOutput() JsonOutputSerializationFormatPtrOutput
	ToJsonOutputSerializationFormatPtrOutputWithContext(context.Context) JsonOutputSerializationFormatPtrOutput
}

type jsonOutputSerializationFormatPtr string

func JsonOutputSerializationFormatPtr(v string) JsonOutputSerializationFormatPtrInput {
	return (*jsonOutputSerializationFormatPtr)(&v)
}

func (*jsonOutputSerializationFormatPtr) ElementType() reflect.Type {
	return jsonOutputSerializationFormatPtrType
}

func (in *jsonOutputSerializationFormatPtr) ToJsonOutputSerializationFormatPtrOutput() JsonOutputSerializationFormatPtrOutput {
	return pulumi.ToOutput(in).(JsonOutputSerializationFormatPtrOutput)
}

func (in *jsonOutputSerializationFormatPtr) ToJsonOutputSerializationFormatPtrOutputWithContext(ctx context.Context) JsonOutputSerializationFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(JsonOutputSerializationFormatPtrOutput)
}

type OutputErrorPolicy string

const (
	OutputErrorPolicyStop = OutputErrorPolicy("Stop")
	OutputErrorPolicyDrop = OutputErrorPolicy("Drop")
)

func (OutputErrorPolicy) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputErrorPolicy)(nil)).Elem()
}

func (e OutputErrorPolicy) ToOutputErrorPolicyOutput() OutputErrorPolicyOutput {
	return pulumi.ToOutput(e).(OutputErrorPolicyOutput)
}

func (e OutputErrorPolicy) ToOutputErrorPolicyOutputWithContext(ctx context.Context) OutputErrorPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OutputErrorPolicyOutput)
}

func (e OutputErrorPolicy) ToOutputErrorPolicyPtrOutput() OutputErrorPolicyPtrOutput {
	return e.ToOutputErrorPolicyPtrOutputWithContext(context.Background())
}

func (e OutputErrorPolicy) ToOutputErrorPolicyPtrOutputWithContext(ctx context.Context) OutputErrorPolicyPtrOutput {
	return OutputErrorPolicy(e).ToOutputErrorPolicyOutputWithContext(ctx).ToOutputErrorPolicyPtrOutputWithContext(ctx)
}

func (e OutputErrorPolicy) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputErrorPolicy) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputErrorPolicy) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OutputErrorPolicy) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OutputErrorPolicyOutput struct{ *pulumi.OutputState }

func (OutputErrorPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputErrorPolicy)(nil)).Elem()
}

func (o OutputErrorPolicyOutput) ToOutputErrorPolicyOutput() OutputErrorPolicyOutput {
	return o
}

func (o OutputErrorPolicyOutput) ToOutputErrorPolicyOutputWithContext(ctx context.Context) OutputErrorPolicyOutput {
	return o
}

func (o OutputErrorPolicyOutput) ToOutputErrorPolicyPtrOutput() OutputErrorPolicyPtrOutput {
	return o.ToOutputErrorPolicyPtrOutputWithContext(context.Background())
}

func (o OutputErrorPolicyOutput) ToOutputErrorPolicyPtrOutputWithContext(ctx context.Context) OutputErrorPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputErrorPolicy) *OutputErrorPolicy {
		return &v
	}).(OutputErrorPolicyPtrOutput)
}

func (o OutputErrorPolicyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutputErrorPolicyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputErrorPolicy) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutputErrorPolicyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputErrorPolicyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputErrorPolicy) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutputErrorPolicyPtrOutput struct{ *pulumi.OutputState }

func (OutputErrorPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputErrorPolicy)(nil)).Elem()
}

func (o OutputErrorPolicyPtrOutput) ToOutputErrorPolicyPtrOutput() OutputErrorPolicyPtrOutput {
	return o
}

func (o OutputErrorPolicyPtrOutput) ToOutputErrorPolicyPtrOutputWithContext(ctx context.Context) OutputErrorPolicyPtrOutput {
	return o
}

func (o OutputErrorPolicyPtrOutput) Elem() OutputErrorPolicyOutput {
	return o.ApplyT(func(v *OutputErrorPolicy) OutputErrorPolicy {
		if v != nil {
			return *v
		}
		var ret OutputErrorPolicy
		return ret
	}).(OutputErrorPolicyOutput)
}

func (o OutputErrorPolicyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputErrorPolicyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutputErrorPolicy) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OutputErrorPolicyInput interface {
	pulumi.Input

	ToOutputErrorPolicyOutput() OutputErrorPolicyOutput
	ToOutputErrorPolicyOutputWithContext(context.Context) OutputErrorPolicyOutput
}

var outputErrorPolicyPtrType = reflect.TypeOf((**OutputErrorPolicy)(nil)).Elem()

type OutputErrorPolicyPtrInput interface {
	pulumi.Input

	ToOutputErrorPolicyPtrOutput() OutputErrorPolicyPtrOutput
	ToOutputErrorPolicyPtrOutputWithContext(context.Context) OutputErrorPolicyPtrOutput
}

type outputErrorPolicyPtr string

func OutputErrorPolicyPtr(v string) OutputErrorPolicyPtrInput {
	return (*outputErrorPolicyPtr)(&v)
}

func (*outputErrorPolicyPtr) ElementType() reflect.Type {
	return outputErrorPolicyPtrType
}

func (in *outputErrorPolicyPtr) ToOutputErrorPolicyPtrOutput() OutputErrorPolicyPtrOutput {
	return pulumi.ToOutput(in).(OutputErrorPolicyPtrOutput)
}

func (in *outputErrorPolicyPtr) ToOutputErrorPolicyPtrOutputWithContext(ctx context.Context) OutputErrorPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OutputErrorPolicyPtrOutput)
}

type OutputStartMode string

const (
	OutputStartModeJobStartTime        = OutputStartMode("JobStartTime")
	OutputStartModeCustomTime          = OutputStartMode("CustomTime")
	OutputStartModeLastOutputEventTime = OutputStartMode("LastOutputEventTime")
)

func (OutputStartMode) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputStartMode)(nil)).Elem()
}

func (e OutputStartMode) ToOutputStartModeOutput() OutputStartModeOutput {
	return pulumi.ToOutput(e).(OutputStartModeOutput)
}

func (e OutputStartMode) ToOutputStartModeOutputWithContext(ctx context.Context) OutputStartModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OutputStartModeOutput)
}

func (e OutputStartMode) ToOutputStartModePtrOutput() OutputStartModePtrOutput {
	return e.ToOutputStartModePtrOutputWithContext(context.Background())
}

func (e OutputStartMode) ToOutputStartModePtrOutputWithContext(ctx context.Context) OutputStartModePtrOutput {
	return OutputStartMode(e).ToOutputStartModeOutputWithContext(ctx).ToOutputStartModePtrOutputWithContext(ctx)
}

func (e OutputStartMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputStartMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OutputStartMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OutputStartMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OutputStartModeOutput struct{ *pulumi.OutputState }

func (OutputStartModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OutputStartMode)(nil)).Elem()
}

func (o OutputStartModeOutput) ToOutputStartModeOutput() OutputStartModeOutput {
	return o
}

func (o OutputStartModeOutput) ToOutputStartModeOutputWithContext(ctx context.Context) OutputStartModeOutput {
	return o
}

func (o OutputStartModeOutput) ToOutputStartModePtrOutput() OutputStartModePtrOutput {
	return o.ToOutputStartModePtrOutputWithContext(context.Background())
}

func (o OutputStartModeOutput) ToOutputStartModePtrOutputWithContext(ctx context.Context) OutputStartModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OutputStartMode) *OutputStartMode {
		return &v
	}).(OutputStartModePtrOutput)
}

func (o OutputStartModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OutputStartModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputStartMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OutputStartModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputStartModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OutputStartMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OutputStartModePtrOutput struct{ *pulumi.OutputState }

func (OutputStartModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OutputStartMode)(nil)).Elem()
}

func (o OutputStartModePtrOutput) ToOutputStartModePtrOutput() OutputStartModePtrOutput {
	return o
}

func (o OutputStartModePtrOutput) ToOutputStartModePtrOutputWithContext(ctx context.Context) OutputStartModePtrOutput {
	return o
}

func (o OutputStartModePtrOutput) Elem() OutputStartModeOutput {
	return o.ApplyT(func(v *OutputStartMode) OutputStartMode {
		if v != nil {
			return *v
		}
		var ret OutputStartMode
		return ret
	}).(OutputStartModeOutput)
}

func (o OutputStartModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OutputStartModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OutputStartMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OutputStartModeInput interface {
	pulumi.Input

	ToOutputStartModeOutput() OutputStartModeOutput
	ToOutputStartModeOutputWithContext(context.Context) OutputStartModeOutput
}

var outputStartModePtrType = reflect.TypeOf((**OutputStartMode)(nil)).Elem()

type OutputStartModePtrInput interface {
	pulumi.Input

	ToOutputStartModePtrOutput() OutputStartModePtrOutput
	ToOutputStartModePtrOutputWithContext(context.Context) OutputStartModePtrOutput
}

type outputStartModePtr string

func OutputStartModePtr(v string) OutputStartModePtrInput {
	return (*outputStartModePtr)(&v)
}

func (*outputStartModePtr) ElementType() reflect.Type {
	return outputStartModePtrType
}

func (in *outputStartModePtr) ToOutputStartModePtrOutput() OutputStartModePtrOutput {
	return pulumi.ToOutput(in).(OutputStartModePtrOutput)
}

func (in *outputStartModePtr) ToOutputStartModePtrOutputWithContext(ctx context.Context) OutputStartModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OutputStartModePtrOutput)
}

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CompatibilityLevelOutput{})
	pulumi.RegisterOutputType(CompatibilityLevelPtrOutput{})
	pulumi.RegisterOutputType(EncodingOutput{})
	pulumi.RegisterOutputType(EncodingPtrOutput{})
	pulumi.RegisterOutputType(EventSerializationTypeOutput{})
	pulumi.RegisterOutputType(EventSerializationTypePtrOutput{})
	pulumi.RegisterOutputType(EventsOutOfOrderPolicyOutput{})
	pulumi.RegisterOutputType(EventsOutOfOrderPolicyPtrOutput{})
	pulumi.RegisterOutputType(JsonOutputSerializationFormatOutput{})
	pulumi.RegisterOutputType(JsonOutputSerializationFormatPtrOutput{})
	pulumi.RegisterOutputType(OutputErrorPolicyOutput{})
	pulumi.RegisterOutputType(OutputErrorPolicyPtrOutput{})
	pulumi.RegisterOutputType(OutputStartModeOutput{})
	pulumi.RegisterOutputType(OutputStartModePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
