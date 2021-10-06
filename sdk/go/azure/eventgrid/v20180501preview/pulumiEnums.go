


package v20180501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DeadLetterEndPointType string

const (
	DeadLetterEndPointTypeStorageBlob = DeadLetterEndPointType("StorageBlob")
)

func (DeadLetterEndPointType) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterEndPointType)(nil)).Elem()
}

func (e DeadLetterEndPointType) ToDeadLetterEndPointTypeOutput() DeadLetterEndPointTypeOutput {
	return pulumi.ToOutput(e).(DeadLetterEndPointTypeOutput)
}

func (e DeadLetterEndPointType) ToDeadLetterEndPointTypeOutputWithContext(ctx context.Context) DeadLetterEndPointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DeadLetterEndPointTypeOutput)
}

func (e DeadLetterEndPointType) ToDeadLetterEndPointTypePtrOutput() DeadLetterEndPointTypePtrOutput {
	return e.ToDeadLetterEndPointTypePtrOutputWithContext(context.Background())
}

func (e DeadLetterEndPointType) ToDeadLetterEndPointTypePtrOutputWithContext(ctx context.Context) DeadLetterEndPointTypePtrOutput {
	return DeadLetterEndPointType(e).ToDeadLetterEndPointTypeOutputWithContext(ctx).ToDeadLetterEndPointTypePtrOutputWithContext(ctx)
}

func (e DeadLetterEndPointType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeadLetterEndPointType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DeadLetterEndPointType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DeadLetterEndPointType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DeadLetterEndPointTypeOutput struct{ *pulumi.OutputState }

func (DeadLetterEndPointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DeadLetterEndPointType)(nil)).Elem()
}

func (o DeadLetterEndPointTypeOutput) ToDeadLetterEndPointTypeOutput() DeadLetterEndPointTypeOutput {
	return o
}

func (o DeadLetterEndPointTypeOutput) ToDeadLetterEndPointTypeOutputWithContext(ctx context.Context) DeadLetterEndPointTypeOutput {
	return o
}

func (o DeadLetterEndPointTypeOutput) ToDeadLetterEndPointTypePtrOutput() DeadLetterEndPointTypePtrOutput {
	return o.ToDeadLetterEndPointTypePtrOutputWithContext(context.Background())
}

func (o DeadLetterEndPointTypeOutput) ToDeadLetterEndPointTypePtrOutputWithContext(ctx context.Context) DeadLetterEndPointTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DeadLetterEndPointType) *DeadLetterEndPointType {
		return &v
	}).(DeadLetterEndPointTypePtrOutput)
}

func (o DeadLetterEndPointTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DeadLetterEndPointTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeadLetterEndPointType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DeadLetterEndPointTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeadLetterEndPointTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DeadLetterEndPointType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DeadLetterEndPointTypePtrOutput struct{ *pulumi.OutputState }

func (DeadLetterEndPointTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DeadLetterEndPointType)(nil)).Elem()
}

func (o DeadLetterEndPointTypePtrOutput) ToDeadLetterEndPointTypePtrOutput() DeadLetterEndPointTypePtrOutput {
	return o
}

func (o DeadLetterEndPointTypePtrOutput) ToDeadLetterEndPointTypePtrOutputWithContext(ctx context.Context) DeadLetterEndPointTypePtrOutput {
	return o
}

func (o DeadLetterEndPointTypePtrOutput) Elem() DeadLetterEndPointTypeOutput {
	return o.ApplyT(func(v *DeadLetterEndPointType) DeadLetterEndPointType {
		if v != nil {
			return *v
		}
		var ret DeadLetterEndPointType
		return ret
	}).(DeadLetterEndPointTypeOutput)
}

func (o DeadLetterEndPointTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DeadLetterEndPointTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DeadLetterEndPointType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DeadLetterEndPointTypeInput interface {
	pulumi.Input

	ToDeadLetterEndPointTypeOutput() DeadLetterEndPointTypeOutput
	ToDeadLetterEndPointTypeOutputWithContext(context.Context) DeadLetterEndPointTypeOutput
}

var deadLetterEndPointTypePtrType = reflect.TypeOf((**DeadLetterEndPointType)(nil)).Elem()

type DeadLetterEndPointTypePtrInput interface {
	pulumi.Input

	ToDeadLetterEndPointTypePtrOutput() DeadLetterEndPointTypePtrOutput
	ToDeadLetterEndPointTypePtrOutputWithContext(context.Context) DeadLetterEndPointTypePtrOutput
}

type deadLetterEndPointTypePtr string

func DeadLetterEndPointTypePtr(v string) DeadLetterEndPointTypePtrInput {
	return (*deadLetterEndPointTypePtr)(&v)
}

func (*deadLetterEndPointTypePtr) ElementType() reflect.Type {
	return deadLetterEndPointTypePtrType
}

func (in *deadLetterEndPointTypePtr) ToDeadLetterEndPointTypePtrOutput() DeadLetterEndPointTypePtrOutput {
	return pulumi.ToOutput(in).(DeadLetterEndPointTypePtrOutput)
}

func (in *deadLetterEndPointTypePtr) ToDeadLetterEndPointTypePtrOutputWithContext(ctx context.Context) DeadLetterEndPointTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DeadLetterEndPointTypePtrOutput)
}

type EndpointType string

const (
	EndpointTypeWebHook          = EndpointType("WebHook")
	EndpointTypeEventHub         = EndpointType("EventHub")
	EndpointTypeStorageQueue     = EndpointType("StorageQueue")
	EndpointTypeHybridConnection = EndpointType("HybridConnection")
)

func (EndpointType) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (e EndpointType) ToEndpointTypeOutput() EndpointTypeOutput {
	return pulumi.ToOutput(e).(EndpointTypeOutput)
}

func (e EndpointType) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EndpointTypeOutput)
}

func (e EndpointType) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return e.ToEndpointTypePtrOutputWithContext(context.Background())
}

func (e EndpointType) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return EndpointType(e).ToEndpointTypeOutputWithContext(ctx).ToEndpointTypePtrOutputWithContext(ctx)
}

func (e EndpointType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EndpointType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EndpointType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EndpointTypeOutput struct{ *pulumi.OutputState }

func (EndpointTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointType)(nil)).Elem()
}

func (o EndpointTypeOutput) ToEndpointTypeOutput() EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypeOutputWithContext(ctx context.Context) EndpointTypeOutput {
	return o
}

func (o EndpointTypeOutput) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return o.ToEndpointTypePtrOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EndpointType) *EndpointType {
		return &v
	}).(EndpointTypePtrOutput)
}

func (o EndpointTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EndpointTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EndpointType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EndpointTypePtrOutput struct{ *pulumi.OutputState }

func (EndpointTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EndpointType)(nil)).Elem()
}

func (o EndpointTypePtrOutput) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return o
}

func (o EndpointTypePtrOutput) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return o
}

func (o EndpointTypePtrOutput) Elem() EndpointTypeOutput {
	return o.ApplyT(func(v *EndpointType) EndpointType {
		if v != nil {
			return *v
		}
		var ret EndpointType
		return ret
	}).(EndpointTypeOutput)
}

func (o EndpointTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EndpointTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EndpointType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EndpointTypeInput interface {
	pulumi.Input

	ToEndpointTypeOutput() EndpointTypeOutput
	ToEndpointTypeOutputWithContext(context.Context) EndpointTypeOutput
}

var endpointTypePtrType = reflect.TypeOf((**EndpointType)(nil)).Elem()

type EndpointTypePtrInput interface {
	pulumi.Input

	ToEndpointTypePtrOutput() EndpointTypePtrOutput
	ToEndpointTypePtrOutputWithContext(context.Context) EndpointTypePtrOutput
}

type endpointTypePtr string

func EndpointTypePtr(v string) EndpointTypePtrInput {
	return (*endpointTypePtr)(&v)
}

func (*endpointTypePtr) ElementType() reflect.Type {
	return endpointTypePtrType
}

func (in *endpointTypePtr) ToEndpointTypePtrOutput() EndpointTypePtrOutput {
	return pulumi.ToOutput(in).(EndpointTypePtrOutput)
}

func (in *endpointTypePtr) ToEndpointTypePtrOutputWithContext(ctx context.Context) EndpointTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EndpointTypePtrOutput)
}

type EventDeliverySchema string

const (
	EventDeliverySchemaEventGridSchema     = EventDeliverySchema("EventGridSchema")
	EventDeliverySchemaInputEventSchema    = EventDeliverySchema("InputEventSchema")
	EventDeliverySchemaCloudEventV01Schema = EventDeliverySchema("CloudEventV01Schema")
)

func (EventDeliverySchema) ElementType() reflect.Type {
	return reflect.TypeOf((*EventDeliverySchema)(nil)).Elem()
}

func (e EventDeliverySchema) ToEventDeliverySchemaOutput() EventDeliverySchemaOutput {
	return pulumi.ToOutput(e).(EventDeliverySchemaOutput)
}

func (e EventDeliverySchema) ToEventDeliverySchemaOutputWithContext(ctx context.Context) EventDeliverySchemaOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventDeliverySchemaOutput)
}

func (e EventDeliverySchema) ToEventDeliverySchemaPtrOutput() EventDeliverySchemaPtrOutput {
	return e.ToEventDeliverySchemaPtrOutputWithContext(context.Background())
}

func (e EventDeliverySchema) ToEventDeliverySchemaPtrOutputWithContext(ctx context.Context) EventDeliverySchemaPtrOutput {
	return EventDeliverySchema(e).ToEventDeliverySchemaOutputWithContext(ctx).ToEventDeliverySchemaPtrOutputWithContext(ctx)
}

func (e EventDeliverySchema) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventDeliverySchema) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventDeliverySchema) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventDeliverySchema) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventDeliverySchemaOutput struct{ *pulumi.OutputState }

func (EventDeliverySchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventDeliverySchema)(nil)).Elem()
}

func (o EventDeliverySchemaOutput) ToEventDeliverySchemaOutput() EventDeliverySchemaOutput {
	return o
}

func (o EventDeliverySchemaOutput) ToEventDeliverySchemaOutputWithContext(ctx context.Context) EventDeliverySchemaOutput {
	return o
}

func (o EventDeliverySchemaOutput) ToEventDeliverySchemaPtrOutput() EventDeliverySchemaPtrOutput {
	return o.ToEventDeliverySchemaPtrOutputWithContext(context.Background())
}

func (o EventDeliverySchemaOutput) ToEventDeliverySchemaPtrOutputWithContext(ctx context.Context) EventDeliverySchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventDeliverySchema) *EventDeliverySchema {
		return &v
	}).(EventDeliverySchemaPtrOutput)
}

func (o EventDeliverySchemaOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventDeliverySchemaOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventDeliverySchema) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventDeliverySchemaOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventDeliverySchemaOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventDeliverySchema) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventDeliverySchemaPtrOutput struct{ *pulumi.OutputState }

func (EventDeliverySchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventDeliverySchema)(nil)).Elem()
}

func (o EventDeliverySchemaPtrOutput) ToEventDeliverySchemaPtrOutput() EventDeliverySchemaPtrOutput {
	return o
}

func (o EventDeliverySchemaPtrOutput) ToEventDeliverySchemaPtrOutputWithContext(ctx context.Context) EventDeliverySchemaPtrOutput {
	return o
}

func (o EventDeliverySchemaPtrOutput) Elem() EventDeliverySchemaOutput {
	return o.ApplyT(func(v *EventDeliverySchema) EventDeliverySchema {
		if v != nil {
			return *v
		}
		var ret EventDeliverySchema
		return ret
	}).(EventDeliverySchemaOutput)
}

func (o EventDeliverySchemaPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventDeliverySchemaPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventDeliverySchema) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventDeliverySchemaInput interface {
	pulumi.Input

	ToEventDeliverySchemaOutput() EventDeliverySchemaOutput
	ToEventDeliverySchemaOutputWithContext(context.Context) EventDeliverySchemaOutput
}

var eventDeliverySchemaPtrType = reflect.TypeOf((**EventDeliverySchema)(nil)).Elem()

type EventDeliverySchemaPtrInput interface {
	pulumi.Input

	ToEventDeliverySchemaPtrOutput() EventDeliverySchemaPtrOutput
	ToEventDeliverySchemaPtrOutputWithContext(context.Context) EventDeliverySchemaPtrOutput
}

type eventDeliverySchemaPtr string

func EventDeliverySchemaPtr(v string) EventDeliverySchemaPtrInput {
	return (*eventDeliverySchemaPtr)(&v)
}

func (*eventDeliverySchemaPtr) ElementType() reflect.Type {
	return eventDeliverySchemaPtrType
}

func (in *eventDeliverySchemaPtr) ToEventDeliverySchemaPtrOutput() EventDeliverySchemaPtrOutput {
	return pulumi.ToOutput(in).(EventDeliverySchemaPtrOutput)
}

func (in *eventDeliverySchemaPtr) ToEventDeliverySchemaPtrOutputWithContext(ctx context.Context) EventDeliverySchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventDeliverySchemaPtrOutput)
}

type InputSchema string

const (
	InputSchemaEventGridSchema     = InputSchema("EventGridSchema")
	InputSchemaCustomEventSchema   = InputSchema("CustomEventSchema")
	InputSchemaCloudEventV01Schema = InputSchema("CloudEventV01Schema")
)

func (InputSchema) ElementType() reflect.Type {
	return reflect.TypeOf((*InputSchema)(nil)).Elem()
}

func (e InputSchema) ToInputSchemaOutput() InputSchemaOutput {
	return pulumi.ToOutput(e).(InputSchemaOutput)
}

func (e InputSchema) ToInputSchemaOutputWithContext(ctx context.Context) InputSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InputSchemaOutput)
}

func (e InputSchema) ToInputSchemaPtrOutput() InputSchemaPtrOutput {
	return e.ToInputSchemaPtrOutputWithContext(context.Background())
}

func (e InputSchema) ToInputSchemaPtrOutputWithContext(ctx context.Context) InputSchemaPtrOutput {
	return InputSchema(e).ToInputSchemaOutputWithContext(ctx).ToInputSchemaPtrOutputWithContext(ctx)
}

func (e InputSchema) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InputSchema) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InputSchema) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InputSchema) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InputSchemaOutput struct{ *pulumi.OutputState }

func (InputSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputSchema)(nil)).Elem()
}

func (o InputSchemaOutput) ToInputSchemaOutput() InputSchemaOutput {
	return o
}

func (o InputSchemaOutput) ToInputSchemaOutputWithContext(ctx context.Context) InputSchemaOutput {
	return o
}

func (o InputSchemaOutput) ToInputSchemaPtrOutput() InputSchemaPtrOutput {
	return o.ToInputSchemaPtrOutputWithContext(context.Background())
}

func (o InputSchemaOutput) ToInputSchemaPtrOutputWithContext(ctx context.Context) InputSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputSchema) *InputSchema {
		return &v
	}).(InputSchemaPtrOutput)
}

func (o InputSchemaOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InputSchemaOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InputSchema) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InputSchemaOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InputSchemaOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InputSchema) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InputSchemaPtrOutput struct{ *pulumi.OutputState }

func (InputSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputSchema)(nil)).Elem()
}

func (o InputSchemaPtrOutput) ToInputSchemaPtrOutput() InputSchemaPtrOutput {
	return o
}

func (o InputSchemaPtrOutput) ToInputSchemaPtrOutputWithContext(ctx context.Context) InputSchemaPtrOutput {
	return o
}

func (o InputSchemaPtrOutput) Elem() InputSchemaOutput {
	return o.ApplyT(func(v *InputSchema) InputSchema {
		if v != nil {
			return *v
		}
		var ret InputSchema
		return ret
	}).(InputSchemaOutput)
}

func (o InputSchemaPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InputSchemaPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InputSchema) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InputSchemaInput interface {
	pulumi.Input

	ToInputSchemaOutput() InputSchemaOutput
	ToInputSchemaOutputWithContext(context.Context) InputSchemaOutput
}

var inputSchemaPtrType = reflect.TypeOf((**InputSchema)(nil)).Elem()

type InputSchemaPtrInput interface {
	pulumi.Input

	ToInputSchemaPtrOutput() InputSchemaPtrOutput
	ToInputSchemaPtrOutputWithContext(context.Context) InputSchemaPtrOutput
}

type inputSchemaPtr string

func InputSchemaPtr(v string) InputSchemaPtrInput {
	return (*inputSchemaPtr)(&v)
}

func (*inputSchemaPtr) ElementType() reflect.Type {
	return inputSchemaPtrType
}

func (in *inputSchemaPtr) ToInputSchemaPtrOutput() InputSchemaPtrOutput {
	return pulumi.ToOutput(in).(InputSchemaPtrOutput)
}

func (in *inputSchemaPtr) ToInputSchemaPtrOutputWithContext(ctx context.Context) InputSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InputSchemaPtrOutput)
}

type InputSchemaMappingType string

const (
	InputSchemaMappingTypeJson = InputSchemaMappingType("Json")
)

func (InputSchemaMappingType) ElementType() reflect.Type {
	return reflect.TypeOf((*InputSchemaMappingType)(nil)).Elem()
}

func (e InputSchemaMappingType) ToInputSchemaMappingTypeOutput() InputSchemaMappingTypeOutput {
	return pulumi.ToOutput(e).(InputSchemaMappingTypeOutput)
}

func (e InputSchemaMappingType) ToInputSchemaMappingTypeOutputWithContext(ctx context.Context) InputSchemaMappingTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(InputSchemaMappingTypeOutput)
}

func (e InputSchemaMappingType) ToInputSchemaMappingTypePtrOutput() InputSchemaMappingTypePtrOutput {
	return e.ToInputSchemaMappingTypePtrOutputWithContext(context.Background())
}

func (e InputSchemaMappingType) ToInputSchemaMappingTypePtrOutputWithContext(ctx context.Context) InputSchemaMappingTypePtrOutput {
	return InputSchemaMappingType(e).ToInputSchemaMappingTypeOutputWithContext(ctx).ToInputSchemaMappingTypePtrOutputWithContext(ctx)
}

func (e InputSchemaMappingType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e InputSchemaMappingType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e InputSchemaMappingType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e InputSchemaMappingType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type InputSchemaMappingTypeOutput struct{ *pulumi.OutputState }

func (InputSchemaMappingTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*InputSchemaMappingType)(nil)).Elem()
}

func (o InputSchemaMappingTypeOutput) ToInputSchemaMappingTypeOutput() InputSchemaMappingTypeOutput {
	return o
}

func (o InputSchemaMappingTypeOutput) ToInputSchemaMappingTypeOutputWithContext(ctx context.Context) InputSchemaMappingTypeOutput {
	return o
}

func (o InputSchemaMappingTypeOutput) ToInputSchemaMappingTypePtrOutput() InputSchemaMappingTypePtrOutput {
	return o.ToInputSchemaMappingTypePtrOutputWithContext(context.Background())
}

func (o InputSchemaMappingTypeOutput) ToInputSchemaMappingTypePtrOutputWithContext(ctx context.Context) InputSchemaMappingTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v InputSchemaMappingType) *InputSchemaMappingType {
		return &v
	}).(InputSchemaMappingTypePtrOutput)
}

func (o InputSchemaMappingTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o InputSchemaMappingTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InputSchemaMappingType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o InputSchemaMappingTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InputSchemaMappingTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e InputSchemaMappingType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type InputSchemaMappingTypePtrOutput struct{ *pulumi.OutputState }

func (InputSchemaMappingTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**InputSchemaMappingType)(nil)).Elem()
}

func (o InputSchemaMappingTypePtrOutput) ToInputSchemaMappingTypePtrOutput() InputSchemaMappingTypePtrOutput {
	return o
}

func (o InputSchemaMappingTypePtrOutput) ToInputSchemaMappingTypePtrOutputWithContext(ctx context.Context) InputSchemaMappingTypePtrOutput {
	return o
}

func (o InputSchemaMappingTypePtrOutput) Elem() InputSchemaMappingTypeOutput {
	return o.ApplyT(func(v *InputSchemaMappingType) InputSchemaMappingType {
		if v != nil {
			return *v
		}
		var ret InputSchemaMappingType
		return ret
	}).(InputSchemaMappingTypeOutput)
}

func (o InputSchemaMappingTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o InputSchemaMappingTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *InputSchemaMappingType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type InputSchemaMappingTypeInput interface {
	pulumi.Input

	ToInputSchemaMappingTypeOutput() InputSchemaMappingTypeOutput
	ToInputSchemaMappingTypeOutputWithContext(context.Context) InputSchemaMappingTypeOutput
}

var inputSchemaMappingTypePtrType = reflect.TypeOf((**InputSchemaMappingType)(nil)).Elem()

type InputSchemaMappingTypePtrInput interface {
	pulumi.Input

	ToInputSchemaMappingTypePtrOutput() InputSchemaMappingTypePtrOutput
	ToInputSchemaMappingTypePtrOutputWithContext(context.Context) InputSchemaMappingTypePtrOutput
}

type inputSchemaMappingTypePtr string

func InputSchemaMappingTypePtr(v string) InputSchemaMappingTypePtrInput {
	return (*inputSchemaMappingTypePtr)(&v)
}

func (*inputSchemaMappingTypePtr) ElementType() reflect.Type {
	return inputSchemaMappingTypePtrType
}

func (in *inputSchemaMappingTypePtr) ToInputSchemaMappingTypePtrOutput() InputSchemaMappingTypePtrOutput {
	return pulumi.ToOutput(in).(InputSchemaMappingTypePtrOutput)
}

func (in *inputSchemaMappingTypePtr) ToInputSchemaMappingTypePtrOutputWithContext(ctx context.Context) InputSchemaMappingTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(InputSchemaMappingTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(DeadLetterEndPointTypeOutput{})
	pulumi.RegisterOutputType(DeadLetterEndPointTypePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
	pulumi.RegisterOutputType(EventDeliverySchemaOutput{})
	pulumi.RegisterOutputType(EventDeliverySchemaPtrOutput{})
	pulumi.RegisterOutputType(InputSchemaOutput{})
	pulumi.RegisterOutputType(InputSchemaPtrOutput{})
	pulumi.RegisterOutputType(InputSchemaMappingTypeOutput{})
	pulumi.RegisterOutputType(InputSchemaMappingTypePtrOutput{})
}
