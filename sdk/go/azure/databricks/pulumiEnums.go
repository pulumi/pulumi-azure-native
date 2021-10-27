


package databricks

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type KeySource string

const (
	KeySourceDefault             = KeySource("Default")
	KeySource_Microsoft_Keyvault = KeySource("Microsoft.Keyvault")
)

func (KeySource) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySource)(nil)).Elem()
}

func (e KeySource) ToKeySourceOutput() KeySourceOutput {
	return pulumi.ToOutput(e).(KeySourceOutput)
}

func (e KeySource) ToKeySourceOutputWithContext(ctx context.Context) KeySourceOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KeySourceOutput)
}

func (e KeySource) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return e.ToKeySourcePtrOutputWithContext(context.Background())
}

func (e KeySource) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return KeySource(e).ToKeySourceOutputWithContext(ctx).ToKeySourcePtrOutputWithContext(ctx)
}

func (e KeySource) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KeySource) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KeySource) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KeySourceOutput struct{ *pulumi.OutputState }

func (KeySourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeySource)(nil)).Elem()
}

func (o KeySourceOutput) ToKeySourceOutput() KeySourceOutput {
	return o
}

func (o KeySourceOutput) ToKeySourceOutputWithContext(ctx context.Context) KeySourceOutput {
	return o
}

func (o KeySourceOutput) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return o.ToKeySourcePtrOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeySource) *KeySource {
		return &v
	}).(KeySourcePtrOutput)
}

func (o KeySourceOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySource) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KeySourceOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySourceOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KeySource) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KeySourcePtrOutput struct{ *pulumi.OutputState }

func (KeySourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeySource)(nil)).Elem()
}

func (o KeySourcePtrOutput) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return o
}

func (o KeySourcePtrOutput) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return o
}

func (o KeySourcePtrOutput) Elem() KeySourceOutput {
	return o.ApplyT(func(v *KeySource) KeySource {
		if v != nil {
			return *v
		}
		var ret KeySource
		return ret
	}).(KeySourceOutput)
}

func (o KeySourcePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KeySourcePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KeySource) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KeySourceInput interface {
	pulumi.Input

	ToKeySourceOutput() KeySourceOutput
	ToKeySourceOutputWithContext(context.Context) KeySourceOutput
}

var keySourcePtrType = reflect.TypeOf((**KeySource)(nil)).Elem()

type KeySourcePtrInput interface {
	pulumi.Input

	ToKeySourcePtrOutput() KeySourcePtrOutput
	ToKeySourcePtrOutputWithContext(context.Context) KeySourcePtrOutput
}

type keySourcePtr string

func KeySourcePtr(v string) KeySourcePtrInput {
	return (*keySourcePtr)(&v)
}

func (*keySourcePtr) ElementType() reflect.Type {
	return keySourcePtrType
}

func (in *keySourcePtr) ToKeySourcePtrOutput() KeySourcePtrOutput {
	return pulumi.ToOutput(in).(KeySourcePtrOutput)
}

func (in *keySourcePtr) ToKeySourcePtrOutputWithContext(ctx context.Context) KeySourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KeySourcePtrOutput)
}

type PrivateLinkServiceConnectionStatus string

const (
	PrivateLinkServiceConnectionStatusPending      = PrivateLinkServiceConnectionStatus("Pending")
	PrivateLinkServiceConnectionStatusApproved     = PrivateLinkServiceConnectionStatus("Approved")
	PrivateLinkServiceConnectionStatusRejected     = PrivateLinkServiceConnectionStatus("Rejected")
	PrivateLinkServiceConnectionStatusDisconnected = PrivateLinkServiceConnectionStatus("Disconnected")
)

func (PrivateLinkServiceConnectionStatus) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutput(e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrivateLinkServiceConnectionStatusOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return e.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return PrivateLinkServiceConnectionStatus(e).ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx).ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrivateLinkServiceConnectionStatus) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrivateLinkServiceConnectionStatusOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatus) *PrivateLinkServiceConnectionStatus {
		return &v
	}).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrivateLinkServiceConnectionStatus) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatusPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) Elem() PrivateLinkServiceConnectionStatusOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatus) PrivateLinkServiceConnectionStatus {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatus
		return ret
	}).(PrivateLinkServiceConnectionStatusOutput)
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrivateLinkServiceConnectionStatus) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrivateLinkServiceConnectionStatusInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusOutput() PrivateLinkServiceConnectionStatusOutput
	ToPrivateLinkServiceConnectionStatusOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusOutput
}

var privateLinkServiceConnectionStatusPtrType = reflect.TypeOf((**PrivateLinkServiceConnectionStatus)(nil)).Elem()

type PrivateLinkServiceConnectionStatusPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput
	ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatusPtrOutput
}

type privateLinkServiceConnectionStatusPtr string

func PrivateLinkServiceConnectionStatusPtr(v string) PrivateLinkServiceConnectionStatusPtrInput {
	return (*privateLinkServiceConnectionStatusPtr)(&v)
}

func (*privateLinkServiceConnectionStatusPtr) ElementType() reflect.Type {
	return privateLinkServiceConnectionStatusPtrType
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutput() PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutput(in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func (in *privateLinkServiceConnectionStatusPtr) ToPrivateLinkServiceConnectionStatusPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrivateLinkServiceConnectionStatusPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*KeySourceInput)(nil)).Elem(), KeySource("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*KeySourcePtrInput)(nil)).Elem(), KeySource("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateLinkServiceConnectionStatusPtrInput)(nil)).Elem(), PrivateLinkServiceConnectionStatus("Pending"))
	pulumi.RegisterOutputType(KeySourceOutput{})
	pulumi.RegisterOutputType(KeySourcePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatusPtrOutput{})
}
