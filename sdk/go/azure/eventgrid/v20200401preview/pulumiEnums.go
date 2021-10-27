


package v20200401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdvancedFilterOperatorType string

const (
	AdvancedFilterOperatorTypeNumberIn                  = AdvancedFilterOperatorType("NumberIn")
	AdvancedFilterOperatorTypeNumberNotIn               = AdvancedFilterOperatorType("NumberNotIn")
	AdvancedFilterOperatorTypeNumberLessThan            = AdvancedFilterOperatorType("NumberLessThan")
	AdvancedFilterOperatorTypeNumberGreaterThan         = AdvancedFilterOperatorType("NumberGreaterThan")
	AdvancedFilterOperatorTypeNumberLessThanOrEquals    = AdvancedFilterOperatorType("NumberLessThanOrEquals")
	AdvancedFilterOperatorTypeNumberGreaterThanOrEquals = AdvancedFilterOperatorType("NumberGreaterThanOrEquals")
	AdvancedFilterOperatorTypeBoolEquals                = AdvancedFilterOperatorType("BoolEquals")
	AdvancedFilterOperatorTypeStringIn                  = AdvancedFilterOperatorType("StringIn")
	AdvancedFilterOperatorTypeStringNotIn               = AdvancedFilterOperatorType("StringNotIn")
	AdvancedFilterOperatorTypeStringBeginsWith          = AdvancedFilterOperatorType("StringBeginsWith")
	AdvancedFilterOperatorTypeStringEndsWith            = AdvancedFilterOperatorType("StringEndsWith")
	AdvancedFilterOperatorTypeStringContains            = AdvancedFilterOperatorType("StringContains")
)

func (AdvancedFilterOperatorType) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedFilterOperatorType)(nil)).Elem()
}

func (e AdvancedFilterOperatorType) ToAdvancedFilterOperatorTypeOutput() AdvancedFilterOperatorTypeOutput {
	return pulumi.ToOutput(e).(AdvancedFilterOperatorTypeOutput)
}

func (e AdvancedFilterOperatorType) ToAdvancedFilterOperatorTypeOutputWithContext(ctx context.Context) AdvancedFilterOperatorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AdvancedFilterOperatorTypeOutput)
}

func (e AdvancedFilterOperatorType) ToAdvancedFilterOperatorTypePtrOutput() AdvancedFilterOperatorTypePtrOutput {
	return e.ToAdvancedFilterOperatorTypePtrOutputWithContext(context.Background())
}

func (e AdvancedFilterOperatorType) ToAdvancedFilterOperatorTypePtrOutputWithContext(ctx context.Context) AdvancedFilterOperatorTypePtrOutput {
	return AdvancedFilterOperatorType(e).ToAdvancedFilterOperatorTypeOutputWithContext(ctx).ToAdvancedFilterOperatorTypePtrOutputWithContext(ctx)
}

func (e AdvancedFilterOperatorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdvancedFilterOperatorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AdvancedFilterOperatorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AdvancedFilterOperatorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AdvancedFilterOperatorTypeOutput struct{ *pulumi.OutputState }

func (AdvancedFilterOperatorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedFilterOperatorType)(nil)).Elem()
}

func (o AdvancedFilterOperatorTypeOutput) ToAdvancedFilterOperatorTypeOutput() AdvancedFilterOperatorTypeOutput {
	return o
}

func (o AdvancedFilterOperatorTypeOutput) ToAdvancedFilterOperatorTypeOutputWithContext(ctx context.Context) AdvancedFilterOperatorTypeOutput {
	return o
}

func (o AdvancedFilterOperatorTypeOutput) ToAdvancedFilterOperatorTypePtrOutput() AdvancedFilterOperatorTypePtrOutput {
	return o.ToAdvancedFilterOperatorTypePtrOutputWithContext(context.Background())
}

func (o AdvancedFilterOperatorTypeOutput) ToAdvancedFilterOperatorTypePtrOutputWithContext(ctx context.Context) AdvancedFilterOperatorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdvancedFilterOperatorType) *AdvancedFilterOperatorType {
		return &v
	}).(AdvancedFilterOperatorTypePtrOutput)
}

func (o AdvancedFilterOperatorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AdvancedFilterOperatorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdvancedFilterOperatorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AdvancedFilterOperatorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdvancedFilterOperatorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AdvancedFilterOperatorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AdvancedFilterOperatorTypePtrOutput struct{ *pulumi.OutputState }

func (AdvancedFilterOperatorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedFilterOperatorType)(nil)).Elem()
}

func (o AdvancedFilterOperatorTypePtrOutput) ToAdvancedFilterOperatorTypePtrOutput() AdvancedFilterOperatorTypePtrOutput {
	return o
}

func (o AdvancedFilterOperatorTypePtrOutput) ToAdvancedFilterOperatorTypePtrOutputWithContext(ctx context.Context) AdvancedFilterOperatorTypePtrOutput {
	return o
}

func (o AdvancedFilterOperatorTypePtrOutput) Elem() AdvancedFilterOperatorTypeOutput {
	return o.ApplyT(func(v *AdvancedFilterOperatorType) AdvancedFilterOperatorType {
		if v != nil {
			return *v
		}
		var ret AdvancedFilterOperatorType
		return ret
	}).(AdvancedFilterOperatorTypeOutput)
}

func (o AdvancedFilterOperatorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AdvancedFilterOperatorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AdvancedFilterOperatorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AdvancedFilterOperatorTypeInput interface {
	pulumi.Input

	ToAdvancedFilterOperatorTypeOutput() AdvancedFilterOperatorTypeOutput
	ToAdvancedFilterOperatorTypeOutputWithContext(context.Context) AdvancedFilterOperatorTypeOutput
}

var advancedFilterOperatorTypePtrType = reflect.TypeOf((**AdvancedFilterOperatorType)(nil)).Elem()

type AdvancedFilterOperatorTypePtrInput interface {
	pulumi.Input

	ToAdvancedFilterOperatorTypePtrOutput() AdvancedFilterOperatorTypePtrOutput
	ToAdvancedFilterOperatorTypePtrOutputWithContext(context.Context) AdvancedFilterOperatorTypePtrOutput
}

type advancedFilterOperatorTypePtr string

func AdvancedFilterOperatorTypePtr(v string) AdvancedFilterOperatorTypePtrInput {
	return (*advancedFilterOperatorTypePtr)(&v)
}

func (*advancedFilterOperatorTypePtr) ElementType() reflect.Type {
	return advancedFilterOperatorTypePtrType
}

func (in *advancedFilterOperatorTypePtr) ToAdvancedFilterOperatorTypePtrOutput() AdvancedFilterOperatorTypePtrOutput {
	return pulumi.ToOutput(in).(AdvancedFilterOperatorTypePtrOutput)
}

func (in *advancedFilterOperatorTypePtr) ToAdvancedFilterOperatorTypePtrOutputWithContext(ctx context.Context) AdvancedFilterOperatorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AdvancedFilterOperatorTypePtrOutput)
}

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
	EndpointTypeServiceBusQueue  = EndpointType("ServiceBusQueue")
	EndpointTypeServiceBusTopic  = EndpointType("ServiceBusTopic")
	EndpointTypeAzureFunction    = EndpointType("AzureFunction")
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
	EventDeliverySchemaEventGridSchema       = EventDeliverySchema("EventGridSchema")
	EventDeliverySchemaCustomInputSchema     = EventDeliverySchema("CustomInputSchema")
	EventDeliverySchema_CloudEventSchemaV1_0 = EventDeliverySchema("CloudEventSchemaV1_0")
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

type EventSubscriptionIdentityType string

const (
	EventSubscriptionIdentityTypeSystemAssigned = EventSubscriptionIdentityType("SystemAssigned")
	EventSubscriptionIdentityTypeUserAssigned   = EventSubscriptionIdentityType("UserAssigned")
)

func (EventSubscriptionIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionIdentityType)(nil)).Elem()
}

func (e EventSubscriptionIdentityType) ToEventSubscriptionIdentityTypeOutput() EventSubscriptionIdentityTypeOutput {
	return pulumi.ToOutput(e).(EventSubscriptionIdentityTypeOutput)
}

func (e EventSubscriptionIdentityType) ToEventSubscriptionIdentityTypeOutputWithContext(ctx context.Context) EventSubscriptionIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EventSubscriptionIdentityTypeOutput)
}

func (e EventSubscriptionIdentityType) ToEventSubscriptionIdentityTypePtrOutput() EventSubscriptionIdentityTypePtrOutput {
	return e.ToEventSubscriptionIdentityTypePtrOutputWithContext(context.Background())
}

func (e EventSubscriptionIdentityType) ToEventSubscriptionIdentityTypePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityTypePtrOutput {
	return EventSubscriptionIdentityType(e).ToEventSubscriptionIdentityTypeOutputWithContext(ctx).ToEventSubscriptionIdentityTypePtrOutputWithContext(ctx)
}

func (e EventSubscriptionIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EventSubscriptionIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EventSubscriptionIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EventSubscriptionIdentityTypeOutput struct{ *pulumi.OutputState }

func (EventSubscriptionIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EventSubscriptionIdentityType)(nil)).Elem()
}

func (o EventSubscriptionIdentityTypeOutput) ToEventSubscriptionIdentityTypeOutput() EventSubscriptionIdentityTypeOutput {
	return o
}

func (o EventSubscriptionIdentityTypeOutput) ToEventSubscriptionIdentityTypeOutputWithContext(ctx context.Context) EventSubscriptionIdentityTypeOutput {
	return o
}

func (o EventSubscriptionIdentityTypeOutput) ToEventSubscriptionIdentityTypePtrOutput() EventSubscriptionIdentityTypePtrOutput {
	return o.ToEventSubscriptionIdentityTypePtrOutputWithContext(context.Background())
}

func (o EventSubscriptionIdentityTypeOutput) ToEventSubscriptionIdentityTypePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EventSubscriptionIdentityType) *EventSubscriptionIdentityType {
		return &v
	}).(EventSubscriptionIdentityTypePtrOutput)
}

func (o EventSubscriptionIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EventSubscriptionIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EventSubscriptionIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EventSubscriptionIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EventSubscriptionIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (EventSubscriptionIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EventSubscriptionIdentityType)(nil)).Elem()
}

func (o EventSubscriptionIdentityTypePtrOutput) ToEventSubscriptionIdentityTypePtrOutput() EventSubscriptionIdentityTypePtrOutput {
	return o
}

func (o EventSubscriptionIdentityTypePtrOutput) ToEventSubscriptionIdentityTypePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityTypePtrOutput {
	return o
}

func (o EventSubscriptionIdentityTypePtrOutput) Elem() EventSubscriptionIdentityTypeOutput {
	return o.ApplyT(func(v *EventSubscriptionIdentityType) EventSubscriptionIdentityType {
		if v != nil {
			return *v
		}
		var ret EventSubscriptionIdentityType
		return ret
	}).(EventSubscriptionIdentityTypeOutput)
}

func (o EventSubscriptionIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EventSubscriptionIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EventSubscriptionIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EventSubscriptionIdentityTypeInput interface {
	pulumi.Input

	ToEventSubscriptionIdentityTypeOutput() EventSubscriptionIdentityTypeOutput
	ToEventSubscriptionIdentityTypeOutputWithContext(context.Context) EventSubscriptionIdentityTypeOutput
}

var eventSubscriptionIdentityTypePtrType = reflect.TypeOf((**EventSubscriptionIdentityType)(nil)).Elem()

type EventSubscriptionIdentityTypePtrInput interface {
	pulumi.Input

	ToEventSubscriptionIdentityTypePtrOutput() EventSubscriptionIdentityTypePtrOutput
	ToEventSubscriptionIdentityTypePtrOutputWithContext(context.Context) EventSubscriptionIdentityTypePtrOutput
}

type eventSubscriptionIdentityTypePtr string

func EventSubscriptionIdentityTypePtr(v string) EventSubscriptionIdentityTypePtrInput {
	return (*eventSubscriptionIdentityTypePtr)(&v)
}

func (*eventSubscriptionIdentityTypePtr) ElementType() reflect.Type {
	return eventSubscriptionIdentityTypePtrType
}

func (in *eventSubscriptionIdentityTypePtr) ToEventSubscriptionIdentityTypePtrOutput() EventSubscriptionIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(EventSubscriptionIdentityTypePtrOutput)
}

func (in *eventSubscriptionIdentityTypePtr) ToEventSubscriptionIdentityTypePtrOutputWithContext(ctx context.Context) EventSubscriptionIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EventSubscriptionIdentityTypePtrOutput)
}

type IdentityType string

const (
	IdentityTypeNone                         = IdentityType("None")
	IdentityTypeSystemAssigned               = IdentityType("SystemAssigned")
	IdentityTypeUserAssigned                 = IdentityType("UserAssigned")
	IdentityType_SystemAssigned_UserAssigned = IdentityType("SystemAssigned, UserAssigned")
)

func (IdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (e IdentityType) ToIdentityTypeOutput() IdentityTypeOutput {
	return pulumi.ToOutput(e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IdentityTypeOutput)
}

func (e IdentityType) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return e.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (e IdentityType) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return IdentityType(e).ToIdentityTypeOutputWithContext(ctx).ToIdentityTypePtrOutputWithContext(ctx)
}

func (e IdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IdentityTypeOutput struct{ *pulumi.OutputState }

func (IdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityType)(nil)).Elem()
}

func (o IdentityTypeOutput) ToIdentityTypeOutput() IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypeOutputWithContext(ctx context.Context) IdentityTypeOutput {
	return o
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o.ToIdentityTypePtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityType) *IdentityType {
		return &v
	}).(IdentityTypePtrOutput)
}

func (o IdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IdentityTypePtrOutput struct{ *pulumi.OutputState }

func (IdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityType)(nil)).Elem()
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return o
}

func (o IdentityTypePtrOutput) Elem() IdentityTypeOutput {
	return o.ApplyT(func(v *IdentityType) IdentityType {
		if v != nil {
			return *v
		}
		var ret IdentityType
		return ret
	}).(IdentityTypeOutput)
}

func (o IdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IdentityTypeInput interface {
	pulumi.Input

	ToIdentityTypeOutput() IdentityTypeOutput
	ToIdentityTypeOutputWithContext(context.Context) IdentityTypeOutput
}

var identityTypePtrType = reflect.TypeOf((**IdentityType)(nil)).Elem()

type IdentityTypePtrInput interface {
	pulumi.Input

	ToIdentityTypePtrOutput() IdentityTypePtrOutput
	ToIdentityTypePtrOutputWithContext(context.Context) IdentityTypePtrOutput
}

type identityTypePtr string

func IdentityTypePtr(v string) IdentityTypePtrInput {
	return (*identityTypePtr)(&v)
}

func (*identityTypePtr) ElementType() reflect.Type {
	return identityTypePtrType
}

func (in *identityTypePtr) ToIdentityTypePtrOutput() IdentityTypePtrOutput {
	return pulumi.ToOutput(in).(IdentityTypePtrOutput)
}

func (in *identityTypePtr) ToIdentityTypePtrOutputWithContext(ctx context.Context) IdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IdentityTypePtrOutput)
}

type InputSchema string

const (
	InputSchemaEventGridSchema       = InputSchema("EventGridSchema")
	InputSchemaCustomEventSchema     = InputSchema("CustomEventSchema")
	InputSchema_CloudEventSchemaV1_0 = InputSchema("CloudEventSchemaV1_0")
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

type IpActionType string

const (
	IpActionTypeAllow = IpActionType("Allow")
)

func (IpActionType) ElementType() reflect.Type {
	return reflect.TypeOf((*IpActionType)(nil)).Elem()
}

func (e IpActionType) ToIpActionTypeOutput() IpActionTypeOutput {
	return pulumi.ToOutput(e).(IpActionTypeOutput)
}

func (e IpActionType) ToIpActionTypeOutputWithContext(ctx context.Context) IpActionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(IpActionTypeOutput)
}

func (e IpActionType) ToIpActionTypePtrOutput() IpActionTypePtrOutput {
	return e.ToIpActionTypePtrOutputWithContext(context.Background())
}

func (e IpActionType) ToIpActionTypePtrOutputWithContext(ctx context.Context) IpActionTypePtrOutput {
	return IpActionType(e).ToIpActionTypeOutputWithContext(ctx).ToIpActionTypePtrOutputWithContext(ctx)
}

func (e IpActionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpActionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e IpActionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e IpActionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type IpActionTypeOutput struct{ *pulumi.OutputState }

func (IpActionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpActionType)(nil)).Elem()
}

func (o IpActionTypeOutput) ToIpActionTypeOutput() IpActionTypeOutput {
	return o
}

func (o IpActionTypeOutput) ToIpActionTypeOutputWithContext(ctx context.Context) IpActionTypeOutput {
	return o
}

func (o IpActionTypeOutput) ToIpActionTypePtrOutput() IpActionTypePtrOutput {
	return o.ToIpActionTypePtrOutputWithContext(context.Background())
}

func (o IpActionTypeOutput) ToIpActionTypePtrOutputWithContext(ctx context.Context) IpActionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IpActionType) *IpActionType {
		return &v
	}).(IpActionTypePtrOutput)
}

func (o IpActionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o IpActionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpActionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o IpActionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpActionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e IpActionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type IpActionTypePtrOutput struct{ *pulumi.OutputState }

func (IpActionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IpActionType)(nil)).Elem()
}

func (o IpActionTypePtrOutput) ToIpActionTypePtrOutput() IpActionTypePtrOutput {
	return o
}

func (o IpActionTypePtrOutput) ToIpActionTypePtrOutputWithContext(ctx context.Context) IpActionTypePtrOutput {
	return o
}

func (o IpActionTypePtrOutput) Elem() IpActionTypeOutput {
	return o.ApplyT(func(v *IpActionType) IpActionType {
		if v != nil {
			return *v
		}
		var ret IpActionType
		return ret
	}).(IpActionTypeOutput)
}

func (o IpActionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o IpActionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *IpActionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type IpActionTypeInput interface {
	pulumi.Input

	ToIpActionTypeOutput() IpActionTypeOutput
	ToIpActionTypeOutputWithContext(context.Context) IpActionTypeOutput
}

var ipActionTypePtrType = reflect.TypeOf((**IpActionType)(nil)).Elem()

type IpActionTypePtrInput interface {
	pulumi.Input

	ToIpActionTypePtrOutput() IpActionTypePtrOutput
	ToIpActionTypePtrOutputWithContext(context.Context) IpActionTypePtrOutput
}

type ipActionTypePtr string

func IpActionTypePtr(v string) IpActionTypePtrInput {
	return (*ipActionTypePtr)(&v)
}

func (*ipActionTypePtr) ElementType() reflect.Type {
	return ipActionTypePtrType
}

func (in *ipActionTypePtr) ToIpActionTypePtrOutput() IpActionTypePtrOutput {
	return pulumi.ToOutput(in).(IpActionTypePtrOutput)
}

func (in *ipActionTypePtr) ToIpActionTypePtrOutputWithContext(ctx context.Context) IpActionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(IpActionTypePtrOutput)
}

type PartnerRegistrationVisibilityState string

const (
	PartnerRegistrationVisibilityStateHidden             = PartnerRegistrationVisibilityState("Hidden")
	PartnerRegistrationVisibilityStatePublicPreview      = PartnerRegistrationVisibilityState("PublicPreview")
	PartnerRegistrationVisibilityStateGenerallyAvailable = PartnerRegistrationVisibilityState("GenerallyAvailable")
)

func (PartnerRegistrationVisibilityState) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegistrationVisibilityState)(nil)).Elem()
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStateOutput() PartnerRegistrationVisibilityStateOutput {
	return pulumi.ToOutput(e).(PartnerRegistrationVisibilityStateOutput)
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStateOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PartnerRegistrationVisibilityStateOutput)
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return e.ToPartnerRegistrationVisibilityStatePtrOutputWithContext(context.Background())
}

func (e PartnerRegistrationVisibilityState) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return PartnerRegistrationVisibilityState(e).ToPartnerRegistrationVisibilityStateOutputWithContext(ctx).ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx)
}

func (e PartnerRegistrationVisibilityState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerRegistrationVisibilityState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PartnerRegistrationVisibilityState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PartnerRegistrationVisibilityState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PartnerRegistrationVisibilityStateOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationVisibilityStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PartnerRegistrationVisibilityState)(nil)).Elem()
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStateOutput() PartnerRegistrationVisibilityStateOutput {
	return o
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStateOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStateOutput {
	return o
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return o.ToPartnerRegistrationVisibilityStatePtrOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStateOutput) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PartnerRegistrationVisibilityState) *PartnerRegistrationVisibilityState {
		return &v
	}).(PartnerRegistrationVisibilityStatePtrOutput)
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerRegistrationVisibilityState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PartnerRegistrationVisibilityState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PartnerRegistrationVisibilityStatePtrOutput struct{ *pulumi.OutputState }

func (PartnerRegistrationVisibilityStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PartnerRegistrationVisibilityState)(nil)).Elem()
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return o
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return o
}

func (o PartnerRegistrationVisibilityStatePtrOutput) Elem() PartnerRegistrationVisibilityStateOutput {
	return o.ApplyT(func(v *PartnerRegistrationVisibilityState) PartnerRegistrationVisibilityState {
		if v != nil {
			return *v
		}
		var ret PartnerRegistrationVisibilityState
		return ret
	}).(PartnerRegistrationVisibilityStateOutput)
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PartnerRegistrationVisibilityStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PartnerRegistrationVisibilityState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PartnerRegistrationVisibilityStateInput interface {
	pulumi.Input

	ToPartnerRegistrationVisibilityStateOutput() PartnerRegistrationVisibilityStateOutput
	ToPartnerRegistrationVisibilityStateOutputWithContext(context.Context) PartnerRegistrationVisibilityStateOutput
}

var partnerRegistrationVisibilityStatePtrType = reflect.TypeOf((**PartnerRegistrationVisibilityState)(nil)).Elem()

type PartnerRegistrationVisibilityStatePtrInput interface {
	pulumi.Input

	ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput
	ToPartnerRegistrationVisibilityStatePtrOutputWithContext(context.Context) PartnerRegistrationVisibilityStatePtrOutput
}

type partnerRegistrationVisibilityStatePtr string

func PartnerRegistrationVisibilityStatePtr(v string) PartnerRegistrationVisibilityStatePtrInput {
	return (*partnerRegistrationVisibilityStatePtr)(&v)
}

func (*partnerRegistrationVisibilityStatePtr) ElementType() reflect.Type {
	return partnerRegistrationVisibilityStatePtrType
}

func (in *partnerRegistrationVisibilityStatePtr) ToPartnerRegistrationVisibilityStatePtrOutput() PartnerRegistrationVisibilityStatePtrOutput {
	return pulumi.ToOutput(in).(PartnerRegistrationVisibilityStatePtrOutput)
}

func (in *partnerRegistrationVisibilityStatePtr) ToPartnerRegistrationVisibilityStatePtrOutputWithContext(ctx context.Context) PartnerRegistrationVisibilityStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PartnerRegistrationVisibilityStatePtrOutput)
}

type PersistedConnectionStatus string

const (
	PersistedConnectionStatusPending      = PersistedConnectionStatus("Pending")
	PersistedConnectionStatusApproved     = PersistedConnectionStatus("Approved")
	PersistedConnectionStatusRejected     = PersistedConnectionStatus("Rejected")
	PersistedConnectionStatusDisconnected = PersistedConnectionStatus("Disconnected")
)

func (PersistedConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistedConnectionStatus)(nil)).Elem()
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusOutput() PersistedConnectionStatusOutput {
	return pulumi.ToOutput(e).(PersistedConnectionStatusOutput)
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusOutputWithContext(ctx context.Context) PersistedConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PersistedConnectionStatusOutput)
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return e.ToPersistedConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PersistedConnectionStatus) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return PersistedConnectionStatus(e).ToPersistedConnectionStatusOutputWithContext(ctx).ToPersistedConnectionStatusPtrOutputWithContext(ctx)
}

func (e PersistedConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PersistedConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PersistedConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PersistedConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PersistedConnectionStatusOutput struct{ *pulumi.OutputState }

func (PersistedConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PersistedConnectionStatus)(nil)).Elem()
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusOutput() PersistedConnectionStatusOutput {
	return o
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusOutputWithContext(ctx context.Context) PersistedConnectionStatusOutput {
	return o
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return o.ToPersistedConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusOutput) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PersistedConnectionStatus) *PersistedConnectionStatus {
		return &v
	}).(PersistedConnectionStatusPtrOutput)
}

func (o PersistedConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PersistedConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PersistedConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PersistedConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PersistedConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PersistedConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PersistedConnectionStatus)(nil)).Elem()
}

func (o PersistedConnectionStatusPtrOutput) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return o
}

func (o PersistedConnectionStatusPtrOutput) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return o
}

func (o PersistedConnectionStatusPtrOutput) Elem() PersistedConnectionStatusOutput {
	return o.ApplyT(func(v *PersistedConnectionStatus) PersistedConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PersistedConnectionStatus
		return ret
	}).(PersistedConnectionStatusOutput)
}

func (o PersistedConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PersistedConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PersistedConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PersistedConnectionStatusInput interface {
	pulumi.Input

	ToPersistedConnectionStatusOutput() PersistedConnectionStatusOutput
	ToPersistedConnectionStatusOutputWithContext(context.Context) PersistedConnectionStatusOutput
}

var persistedConnectionStatusPtrType = reflect.TypeOf((**PersistedConnectionStatus)(nil)).Elem()

type PersistedConnectionStatusPtrInput interface {
	pulumi.Input

	ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput
	ToPersistedConnectionStatusPtrOutputWithContext(context.Context) PersistedConnectionStatusPtrOutput
}

type persistedConnectionStatusPtr string

func PersistedConnectionStatusPtr(v string) PersistedConnectionStatusPtrInput {
	return (*persistedConnectionStatusPtr)(&v)
}

func (*persistedConnectionStatusPtr) ElementType() reflect.Type {
	return persistedConnectionStatusPtrType
}

func (in *persistedConnectionStatusPtr) ToPersistedConnectionStatusPtrOutput() PersistedConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PersistedConnectionStatusPtrOutput)
}

func (in *persistedConnectionStatusPtr) ToPersistedConnectionStatusPtrOutputWithContext(ctx context.Context) PersistedConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PersistedConnectionStatusPtrOutput)
}

type PublicNetworkAccess string

const (
	PublicNetworkAccessEnabled  = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled = PublicNetworkAccess("Disabled")
)

func (PublicNetworkAccess) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return pulumi.ToOutput(e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PublicNetworkAccessOutput)
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return e.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return PublicNetworkAccess(e).ToPublicNetworkAccessOutputWithContext(ctx).ToPublicNetworkAccessPtrOutputWithContext(ctx)
}

func (e PublicNetworkAccess) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PublicNetworkAccess) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PublicNetworkAccess) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PublicNetworkAccessOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutput() PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessOutputWithContext(ctx context.Context) PublicNetworkAccessOutput {
	return o
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o.ToPublicNetworkAccessPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PublicNetworkAccess) *PublicNetworkAccess {
		return &v
	}).(PublicNetworkAccessPtrOutput)
}

func (o PublicNetworkAccessOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PublicNetworkAccessOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PublicNetworkAccess) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PublicNetworkAccessPtrOutput struct{ *pulumi.OutputState }

func (PublicNetworkAccessPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return o
}

func (o PublicNetworkAccessPtrOutput) Elem() PublicNetworkAccessOutput {
	return o.ApplyT(func(v *PublicNetworkAccess) PublicNetworkAccess {
		if v != nil {
			return *v
		}
		var ret PublicNetworkAccess
		return ret
	}).(PublicNetworkAccessOutput)
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PublicNetworkAccessPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PublicNetworkAccess) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PublicNetworkAccessInput interface {
	pulumi.Input

	ToPublicNetworkAccessOutput() PublicNetworkAccessOutput
	ToPublicNetworkAccessOutputWithContext(context.Context) PublicNetworkAccessOutput
}

var publicNetworkAccessPtrType = reflect.TypeOf((**PublicNetworkAccess)(nil)).Elem()

type PublicNetworkAccessPtrInput interface {
	pulumi.Input

	ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput
	ToPublicNetworkAccessPtrOutputWithContext(context.Context) PublicNetworkAccessPtrOutput
}

type publicNetworkAccessPtr string

func PublicNetworkAccessPtr(v string) PublicNetworkAccessPtrInput {
	return (*publicNetworkAccessPtr)(&v)
}

func (*publicNetworkAccessPtr) ElementType() reflect.Type {
	return publicNetworkAccessPtrType
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutput() PublicNetworkAccessPtrOutput {
	return pulumi.ToOutput(in).(PublicNetworkAccessPtrOutput)
}

func (in *publicNetworkAccessPtr) ToPublicNetworkAccessPtrOutputWithContext(ctx context.Context) PublicNetworkAccessPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PublicNetworkAccessPtrOutput)
}

type ResourceProvisioningState string

const (
	ResourceProvisioningStateCreating  = ResourceProvisioningState("Creating")
	ResourceProvisioningStateUpdating  = ResourceProvisioningState("Updating")
	ResourceProvisioningStateDeleting  = ResourceProvisioningState("Deleting")
	ResourceProvisioningStateSucceeded = ResourceProvisioningState("Succeeded")
	ResourceProvisioningStateCanceled  = ResourceProvisioningState("Canceled")
	ResourceProvisioningStateFailed    = ResourceProvisioningState("Failed")
)

func (ResourceProvisioningState) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProvisioningState)(nil)).Elem()
}

func (e ResourceProvisioningState) ToResourceProvisioningStateOutput() ResourceProvisioningStateOutput {
	return pulumi.ToOutput(e).(ResourceProvisioningStateOutput)
}

func (e ResourceProvisioningState) ToResourceProvisioningStateOutputWithContext(ctx context.Context) ResourceProvisioningStateOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceProvisioningStateOutput)
}

func (e ResourceProvisioningState) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return e.ToResourceProvisioningStatePtrOutputWithContext(context.Background())
}

func (e ResourceProvisioningState) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return ResourceProvisioningState(e).ToResourceProvisioningStateOutputWithContext(ctx).ToResourceProvisioningStatePtrOutputWithContext(ctx)
}

func (e ResourceProvisioningState) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceProvisioningState) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceProvisioningState) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceProvisioningState) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceProvisioningStateOutput struct{ *pulumi.OutputState }

func (ResourceProvisioningStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceProvisioningState)(nil)).Elem()
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStateOutput() ResourceProvisioningStateOutput {
	return o
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStateOutputWithContext(ctx context.Context) ResourceProvisioningStateOutput {
	return o
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return o.ToResourceProvisioningStatePtrOutputWithContext(context.Background())
}

func (o ResourceProvisioningStateOutput) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceProvisioningState) *ResourceProvisioningState {
		return &v
	}).(ResourceProvisioningStatePtrOutput)
}

func (o ResourceProvisioningStateOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceProvisioningStateOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceProvisioningState) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceProvisioningStateOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceProvisioningStateOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceProvisioningState) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceProvisioningStatePtrOutput struct{ *pulumi.OutputState }

func (ResourceProvisioningStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceProvisioningState)(nil)).Elem()
}

func (o ResourceProvisioningStatePtrOutput) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return o
}

func (o ResourceProvisioningStatePtrOutput) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return o
}

func (o ResourceProvisioningStatePtrOutput) Elem() ResourceProvisioningStateOutput {
	return o.ApplyT(func(v *ResourceProvisioningState) ResourceProvisioningState {
		if v != nil {
			return *v
		}
		var ret ResourceProvisioningState
		return ret
	}).(ResourceProvisioningStateOutput)
}

func (o ResourceProvisioningStatePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceProvisioningStatePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceProvisioningState) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceProvisioningStateInput interface {
	pulumi.Input

	ToResourceProvisioningStateOutput() ResourceProvisioningStateOutput
	ToResourceProvisioningStateOutputWithContext(context.Context) ResourceProvisioningStateOutput
}

var resourceProvisioningStatePtrType = reflect.TypeOf((**ResourceProvisioningState)(nil)).Elem()

type ResourceProvisioningStatePtrInput interface {
	pulumi.Input

	ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput
	ToResourceProvisioningStatePtrOutputWithContext(context.Context) ResourceProvisioningStatePtrOutput
}

type resourceProvisioningStatePtr string

func ResourceProvisioningStatePtr(v string) ResourceProvisioningStatePtrInput {
	return (*resourceProvisioningStatePtr)(&v)
}

func (*resourceProvisioningStatePtr) ElementType() reflect.Type {
	return resourceProvisioningStatePtrType
}

func (in *resourceProvisioningStatePtr) ToResourceProvisioningStatePtrOutput() ResourceProvisioningStatePtrOutput {
	return pulumi.ToOutput(in).(ResourceProvisioningStatePtrOutput)
}

func (in *resourceProvisioningStatePtr) ToResourceProvisioningStatePtrOutputWithContext(ctx context.Context) ResourceProvisioningStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceProvisioningStatePtrOutput)
}

type Sku string

const (
	SkuBasic   = Sku("Basic")
	SkuPremium = Sku("Premium")
)

func (Sku) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (e Sku) ToSkuOutput() SkuOutput {
	return pulumi.ToOutput(e).(SkuOutput)
}

func (e Sku) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuOutput)
}

func (e Sku) ToSkuPtrOutput() SkuPtrOutput {
	return e.ToSkuPtrOutputWithContext(context.Background())
}

func (e Sku) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return Sku(e).ToSkuOutputWithContext(ctx).ToSkuPtrOutputWithContext(ctx)
}

func (e Sku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Sku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Sku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Sku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
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

func (o SkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Sku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Sku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
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

func (o SkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Sku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

var skuPtrType = reflect.TypeOf((**Sku)(nil)).Elem()

type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtr string

func SkuPtr(v string) SkuPtrInput {
	return (*skuPtr)(&v)
}

func (*skuPtr) ElementType() reflect.Type {
	return skuPtrType
}

func (in *skuPtr) ToSkuPtrOutput() SkuPtrOutput {
	return pulumi.ToOutput(in).(SkuPtrOutput)
}

func (in *skuPtr) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedFilterOperatorTypeInput)(nil)).Elem(), AdvancedFilterOperatorType("NumberIn"))
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedFilterOperatorTypePtrInput)(nil)).Elem(), AdvancedFilterOperatorType("NumberIn"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeadLetterEndPointTypeInput)(nil)).Elem(), DeadLetterEndPointType("StorageBlob"))
	pulumi.RegisterInputType(reflect.TypeOf((*DeadLetterEndPointTypePtrInput)(nil)).Elem(), DeadLetterEndPointType("StorageBlob"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypeInput)(nil)).Elem(), EndpointType("WebHook"))
	pulumi.RegisterInputType(reflect.TypeOf((*EndpointTypePtrInput)(nil)).Elem(), EndpointType("WebHook"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventDeliverySchemaInput)(nil)).Elem(), EventDeliverySchema("EventGridSchema"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventDeliverySchemaPtrInput)(nil)).Elem(), EventDeliverySchema("EventGridSchema"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionIdentityTypeInput)(nil)).Elem(), EventSubscriptionIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*EventSubscriptionIdentityTypePtrInput)(nil)).Elem(), EventSubscriptionIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypeInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityTypePtrInput)(nil)).Elem(), IdentityType("None"))
	pulumi.RegisterInputType(reflect.TypeOf((*InputSchemaInput)(nil)).Elem(), InputSchema("EventGridSchema"))
	pulumi.RegisterInputType(reflect.TypeOf((*InputSchemaPtrInput)(nil)).Elem(), InputSchema("EventGridSchema"))
	pulumi.RegisterInputType(reflect.TypeOf((*InputSchemaMappingTypeInput)(nil)).Elem(), InputSchemaMappingType("Json"))
	pulumi.RegisterInputType(reflect.TypeOf((*InputSchemaMappingTypePtrInput)(nil)).Elem(), InputSchemaMappingType("Json"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpActionTypeInput)(nil)).Elem(), IpActionType("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*IpActionTypePtrInput)(nil)).Elem(), IpActionType("Allow"))
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegistrationVisibilityStateInput)(nil)).Elem(), PartnerRegistrationVisibilityState("Hidden"))
	pulumi.RegisterInputType(reflect.TypeOf((*PartnerRegistrationVisibilityStatePtrInput)(nil)).Elem(), PartnerRegistrationVisibilityState("Hidden"))
	pulumi.RegisterInputType(reflect.TypeOf((*PersistedConnectionStatusInput)(nil)).Elem(), PersistedConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PersistedConnectionStatusPtrInput)(nil)).Elem(), PersistedConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*PublicNetworkAccessPtrInput)(nil)).Elem(), PublicNetworkAccess("Enabled"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceProvisioningStateInput)(nil)).Elem(), ResourceProvisioningState("Creating"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceProvisioningStatePtrInput)(nil)).Elem(), ResourceProvisioningState("Creating"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), Sku("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), Sku("Basic"))
	pulumi.RegisterOutputType(AdvancedFilterOperatorTypeOutput{})
	pulumi.RegisterOutputType(AdvancedFilterOperatorTypePtrOutput{})
	pulumi.RegisterOutputType(DeadLetterEndPointTypeOutput{})
	pulumi.RegisterOutputType(DeadLetterEndPointTypePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
	pulumi.RegisterOutputType(EventDeliverySchemaOutput{})
	pulumi.RegisterOutputType(EventDeliverySchemaPtrOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityTypeOutput{})
	pulumi.RegisterOutputType(EventSubscriptionIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(IdentityTypeOutput{})
	pulumi.RegisterOutputType(IdentityTypePtrOutput{})
	pulumi.RegisterOutputType(InputSchemaOutput{})
	pulumi.RegisterOutputType(InputSchemaPtrOutput{})
	pulumi.RegisterOutputType(InputSchemaMappingTypeOutput{})
	pulumi.RegisterOutputType(InputSchemaMappingTypePtrOutput{})
	pulumi.RegisterOutputType(IpActionTypeOutput{})
	pulumi.RegisterOutputType(IpActionTypePtrOutput{})
	pulumi.RegisterOutputType(PartnerRegistrationVisibilityStateOutput{})
	pulumi.RegisterOutputType(PartnerRegistrationVisibilityStatePtrOutput{})
	pulumi.RegisterOutputType(PersistedConnectionStatusOutput{})
	pulumi.RegisterOutputType(PersistedConnectionStatusPtrOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(ResourceProvisioningStateOutput{})
	pulumi.RegisterOutputType(ResourceProvisioningStatePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
}
