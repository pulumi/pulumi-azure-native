


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SubscriptionFeatureRegistrationStateEnum string

const (
	SubscriptionFeatureRegistrationStateEnumNotSpecified  = SubscriptionFeatureRegistrationStateEnum("NotSpecified")
	SubscriptionFeatureRegistrationStateEnumNotRegistered = SubscriptionFeatureRegistrationStateEnum("NotRegistered")
	SubscriptionFeatureRegistrationStateEnumPending       = SubscriptionFeatureRegistrationStateEnum("Pending")
	SubscriptionFeatureRegistrationStateEnumRegistering   = SubscriptionFeatureRegistrationStateEnum("Registering")
	SubscriptionFeatureRegistrationStateEnumRegistered    = SubscriptionFeatureRegistrationStateEnum("Registered")
	SubscriptionFeatureRegistrationStateEnumUnregistering = SubscriptionFeatureRegistrationStateEnum("Unregistering")
	SubscriptionFeatureRegistrationStateEnumUnregistered  = SubscriptionFeatureRegistrationStateEnum("Unregistered")
)

func (SubscriptionFeatureRegistrationStateEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistrationStateEnum)(nil)).Elem()
}

func (e SubscriptionFeatureRegistrationStateEnum) ToSubscriptionFeatureRegistrationStateEnumOutput() SubscriptionFeatureRegistrationStateEnumOutput {
	return pulumi.ToOutput(e).(SubscriptionFeatureRegistrationStateEnumOutput)
}

func (e SubscriptionFeatureRegistrationStateEnum) ToSubscriptionFeatureRegistrationStateEnumOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationStateEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SubscriptionFeatureRegistrationStateEnumOutput)
}

func (e SubscriptionFeatureRegistrationStateEnum) ToSubscriptionFeatureRegistrationStateEnumPtrOutput() SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return e.ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(context.Background())
}

func (e SubscriptionFeatureRegistrationStateEnum) ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return SubscriptionFeatureRegistrationStateEnum(e).ToSubscriptionFeatureRegistrationStateEnumOutputWithContext(ctx).ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(ctx)
}

func (e SubscriptionFeatureRegistrationStateEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriptionFeatureRegistrationStateEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SubscriptionFeatureRegistrationStateEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SubscriptionFeatureRegistrationStateEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SubscriptionFeatureRegistrationStateEnumOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationStateEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionFeatureRegistrationStateEnum)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToSubscriptionFeatureRegistrationStateEnumOutput() SubscriptionFeatureRegistrationStateEnumOutput {
	return o
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToSubscriptionFeatureRegistrationStateEnumOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationStateEnumOutput {
	return o
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToSubscriptionFeatureRegistrationStateEnumPtrOutput() SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return o.ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(context.Background())
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SubscriptionFeatureRegistrationStateEnum) *SubscriptionFeatureRegistrationStateEnum {
		return &v
	}).(SubscriptionFeatureRegistrationStateEnumPtrOutput)
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionFeatureRegistrationStateEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionFeatureRegistrationStateEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SubscriptionFeatureRegistrationStateEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SubscriptionFeatureRegistrationStateEnumPtrOutput struct{ *pulumi.OutputState }

func (SubscriptionFeatureRegistrationStateEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SubscriptionFeatureRegistrationStateEnum)(nil)).Elem()
}

func (o SubscriptionFeatureRegistrationStateEnumPtrOutput) ToSubscriptionFeatureRegistrationStateEnumPtrOutput() SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return o
}

func (o SubscriptionFeatureRegistrationStateEnumPtrOutput) ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return o
}

func (o SubscriptionFeatureRegistrationStateEnumPtrOutput) Elem() SubscriptionFeatureRegistrationStateEnumOutput {
	return o.ApplyT(func(v *SubscriptionFeatureRegistrationStateEnum) SubscriptionFeatureRegistrationStateEnum {
		if v != nil {
			return *v
		}
		var ret SubscriptionFeatureRegistrationStateEnum
		return ret
	}).(SubscriptionFeatureRegistrationStateEnumOutput)
}

func (o SubscriptionFeatureRegistrationStateEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SubscriptionFeatureRegistrationStateEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SubscriptionFeatureRegistrationStateEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SubscriptionFeatureRegistrationStateEnumInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationStateEnumOutput() SubscriptionFeatureRegistrationStateEnumOutput
	ToSubscriptionFeatureRegistrationStateEnumOutputWithContext(context.Context) SubscriptionFeatureRegistrationStateEnumOutput
}

var subscriptionFeatureRegistrationStateEnumPtrType = reflect.TypeOf((**SubscriptionFeatureRegistrationStateEnum)(nil)).Elem()

type SubscriptionFeatureRegistrationStateEnumPtrInput interface {
	pulumi.Input

	ToSubscriptionFeatureRegistrationStateEnumPtrOutput() SubscriptionFeatureRegistrationStateEnumPtrOutput
	ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(context.Context) SubscriptionFeatureRegistrationStateEnumPtrOutput
}

type subscriptionFeatureRegistrationStateEnumPtr string

func SubscriptionFeatureRegistrationStateEnumPtr(v string) SubscriptionFeatureRegistrationStateEnumPtrInput {
	return (*subscriptionFeatureRegistrationStateEnumPtr)(&v)
}

func (*subscriptionFeatureRegistrationStateEnumPtr) ElementType() reflect.Type {
	return subscriptionFeatureRegistrationStateEnumPtrType
}

func (in *subscriptionFeatureRegistrationStateEnumPtr) ToSubscriptionFeatureRegistrationStateEnumPtrOutput() SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return pulumi.ToOutput(in).(SubscriptionFeatureRegistrationStateEnumPtrOutput)
}

func (in *subscriptionFeatureRegistrationStateEnumPtr) ToSubscriptionFeatureRegistrationStateEnumPtrOutputWithContext(ctx context.Context) SubscriptionFeatureRegistrationStateEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SubscriptionFeatureRegistrationStateEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationStateEnumOutput{})
	pulumi.RegisterOutputType(SubscriptionFeatureRegistrationStateEnumPtrOutput{})
}
