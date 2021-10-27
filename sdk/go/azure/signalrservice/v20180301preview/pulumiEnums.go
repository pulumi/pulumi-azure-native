


package v20180301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SignalRSkuTier string

const (
	SignalRSkuTierFree     = SignalRSkuTier("Free")
	SignalRSkuTierBasic    = SignalRSkuTier("Basic")
	SignalRSkuTierStandard = SignalRSkuTier("Standard")
	SignalRSkuTierPremium  = SignalRSkuTier("Premium")
)

func (SignalRSkuTier) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSkuTier)(nil)).Elem()
}

func (e SignalRSkuTier) ToSignalRSkuTierOutput() SignalRSkuTierOutput {
	return pulumi.ToOutput(e).(SignalRSkuTierOutput)
}

func (e SignalRSkuTier) ToSignalRSkuTierOutputWithContext(ctx context.Context) SignalRSkuTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SignalRSkuTierOutput)
}

func (e SignalRSkuTier) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return e.ToSignalRSkuTierPtrOutputWithContext(context.Background())
}

func (e SignalRSkuTier) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return SignalRSkuTier(e).ToSignalRSkuTierOutputWithContext(ctx).ToSignalRSkuTierPtrOutputWithContext(ctx)
}

func (e SignalRSkuTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRSkuTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SignalRSkuTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SignalRSkuTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SignalRSkuTierOutput struct{ *pulumi.OutputState }

func (SignalRSkuTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SignalRSkuTier)(nil)).Elem()
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierOutput() SignalRSkuTierOutput {
	return o
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierOutputWithContext(ctx context.Context) SignalRSkuTierOutput {
	return o
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return o.ToSignalRSkuTierPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SignalRSkuTier) *SignalRSkuTier {
		return &v
	}).(SignalRSkuTierPtrOutput)
}

func (o SignalRSkuTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRSkuTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SignalRSkuTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SignalRSkuTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SignalRSkuTierPtrOutput struct{ *pulumi.OutputState }

func (SignalRSkuTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SignalRSkuTier)(nil)).Elem()
}

func (o SignalRSkuTierPtrOutput) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return o
}

func (o SignalRSkuTierPtrOutput) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return o
}

func (o SignalRSkuTierPtrOutput) Elem() SignalRSkuTierOutput {
	return o.ApplyT(func(v *SignalRSkuTier) SignalRSkuTier {
		if v != nil {
			return *v
		}
		var ret SignalRSkuTier
		return ret
	}).(SignalRSkuTierOutput)
}

func (o SignalRSkuTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SignalRSkuTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SignalRSkuTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SignalRSkuTierInput interface {
	pulumi.Input

	ToSignalRSkuTierOutput() SignalRSkuTierOutput
	ToSignalRSkuTierOutputWithContext(context.Context) SignalRSkuTierOutput
}

var signalRSkuTierPtrType = reflect.TypeOf((**SignalRSkuTier)(nil)).Elem()

type SignalRSkuTierPtrInput interface {
	pulumi.Input

	ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput
	ToSignalRSkuTierPtrOutputWithContext(context.Context) SignalRSkuTierPtrOutput
}

type signalRSkuTierPtr string

func SignalRSkuTierPtr(v string) SignalRSkuTierPtrInput {
	return (*signalRSkuTierPtr)(&v)
}

func (*signalRSkuTierPtr) ElementType() reflect.Type {
	return signalRSkuTierPtrType
}

func (in *signalRSkuTierPtr) ToSignalRSkuTierPtrOutput() SignalRSkuTierPtrOutput {
	return pulumi.ToOutput(in).(SignalRSkuTierPtrOutput)
}

func (in *signalRSkuTierPtr) ToSignalRSkuTierPtrOutputWithContext(ctx context.Context) SignalRSkuTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SignalRSkuTierPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*SignalRSkuTierInput)(nil)).Elem(), SignalRSkuTier("Free"))
	pulumi.RegisterInputType(reflect.TypeOf((*SignalRSkuTierPtrInput)(nil)).Elem(), SignalRSkuTier("Free"))
	pulumi.RegisterOutputType(SignalRSkuTierOutput{})
	pulumi.RegisterOutputType(SignalRSkuTierPtrOutput{})
}
