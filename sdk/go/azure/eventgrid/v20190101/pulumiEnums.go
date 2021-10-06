


package v20190101

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

func init() {
	pulumi.RegisterOutputType(DeadLetterEndPointTypeOutput{})
	pulumi.RegisterOutputType(DeadLetterEndPointTypePtrOutput{})
	pulumi.RegisterOutputType(EndpointTypeOutput{})
	pulumi.RegisterOutputType(EndpointTypePtrOutput{})
}
