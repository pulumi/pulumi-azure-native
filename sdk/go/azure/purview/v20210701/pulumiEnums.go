


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PublicNetworkAccess string

const (
	PublicNetworkAccessNotSpecified = PublicNetworkAccess("NotSpecified")
	PublicNetworkAccessEnabled      = PublicNetworkAccess("Enabled")
	PublicNetworkAccessDisabled     = PublicNetworkAccess("Disabled")
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

type Status string

const (
	StatusUnknown      = Status("Unknown")
	StatusPending      = Status("Pending")
	StatusApproved     = Status("Approved")
	StatusRejected     = Status("Rejected")
	StatusDisconnected = Status("Disconnected")
)

func (Status) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (e Status) ToStatusOutput() StatusOutput {
	return pulumi.ToOutput(e).(StatusOutput)
}

func (e Status) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusOutput)
}

func (e Status) ToStatusPtrOutput() StatusPtrOutput {
	return e.ToStatusPtrOutputWithContext(context.Background())
}

func (e Status) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return Status(e).ToStatusOutputWithContext(ctx).ToStatusPtrOutputWithContext(ctx)
}

func (e Status) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Status) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Status) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusOutput struct{ *pulumi.OutputState }

func (StatusOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Status)(nil)).Elem()
}

func (o StatusOutput) ToStatusOutput() StatusOutput {
	return o
}

func (o StatusOutput) ToStatusOutputWithContext(ctx context.Context) StatusOutput {
	return o
}

func (o StatusOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o.ToStatusPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Status) *Status {
		return &v
	}).(StatusPtrOutput)
}

func (o StatusOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Status) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusPtrOutput struct{ *pulumi.OutputState }

func (StatusPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Status)(nil)).Elem()
}

func (o StatusPtrOutput) ToStatusPtrOutput() StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return o
}

func (o StatusPtrOutput) Elem() StatusOutput {
	return o.ApplyT(func(v *Status) Status {
		if v != nil {
			return *v
		}
		var ret Status
		return ret
	}).(StatusOutput)
}

func (o StatusPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Status) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusInput interface {
	pulumi.Input

	ToStatusOutput() StatusOutput
	ToStatusOutputWithContext(context.Context) StatusOutput
}

var statusPtrType = reflect.TypeOf((**Status)(nil)).Elem()

type StatusPtrInput interface {
	pulumi.Input

	ToStatusPtrOutput() StatusPtrOutput
	ToStatusPtrOutputWithContext(context.Context) StatusPtrOutput
}

type statusPtr string

func StatusPtr(v string) StatusPtrInput {
	return (*statusPtr)(&v)
}

func (*statusPtr) ElementType() reflect.Type {
	return statusPtrType
}

func (in *statusPtr) ToStatusPtrOutput() StatusPtrOutput {
	return pulumi.ToOutput(in).(StatusPtrOutput)
}

func (in *statusPtr) ToStatusPtrOutputWithContext(ctx context.Context) StatusPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusPtrOutput)
}

type Type string

const (
	TypeSystemAssigned = Type("SystemAssigned")
)

func (Type) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (e Type) ToTypeOutput() TypeOutput {
	return pulumi.ToOutput(e).(TypeOutput)
}

func (e Type) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TypeOutput)
}

func (e Type) ToTypePtrOutput() TypePtrOutput {
	return e.ToTypePtrOutputWithContext(context.Background())
}

func (e Type) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return Type(e).ToTypeOutputWithContext(ctx).ToTypePtrOutputWithContext(ctx)
}

func (e Type) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e Type) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e Type) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TypeOutput struct{ *pulumi.OutputState }

func (TypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Type)(nil)).Elem()
}

func (o TypeOutput) ToTypeOutput() TypeOutput {
	return o
}

func (o TypeOutput) ToTypeOutputWithContext(ctx context.Context) TypeOutput {
	return o
}

func (o TypeOutput) ToTypePtrOutput() TypePtrOutput {
	return o.ToTypePtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Type) *Type {
		return &v
	}).(TypePtrOutput)
}

func (o TypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e Type) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TypePtrOutput struct{ *pulumi.OutputState }

func (TypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Type)(nil)).Elem()
}

func (o TypePtrOutput) ToTypePtrOutput() TypePtrOutput {
	return o
}

func (o TypePtrOutput) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return o
}

func (o TypePtrOutput) Elem() TypeOutput {
	return o.ApplyT(func(v *Type) Type {
		if v != nil {
			return *v
		}
		var ret Type
		return ret
	}).(TypeOutput)
}

func (o TypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *Type) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TypeInput interface {
	pulumi.Input

	ToTypeOutput() TypeOutput
	ToTypeOutputWithContext(context.Context) TypeOutput
}

var typePtrType = reflect.TypeOf((**Type)(nil)).Elem()

type TypePtrInput interface {
	pulumi.Input

	ToTypePtrOutput() TypePtrOutput
	ToTypePtrOutputWithContext(context.Context) TypePtrOutput
}

type typePtr string

func TypePtr(v string) TypePtrInput {
	return (*typePtr)(&v)
}

func (*typePtr) ElementType() reflect.Type {
	return typePtrType
}

func (in *typePtr) ToTypePtrOutput() TypePtrOutput {
	return pulumi.ToOutput(in).(TypePtrOutput)
}

func (in *typePtr) ToTypePtrOutputWithContext(ctx context.Context) TypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PublicNetworkAccessOutput{})
	pulumi.RegisterOutputType(PublicNetworkAccessPtrOutput{})
	pulumi.RegisterOutputType(StatusOutput{})
	pulumi.RegisterOutputType(StatusPtrOutput{})
	pulumi.RegisterOutputType(TypeOutput{})
	pulumi.RegisterOutputType(TypePtrOutput{})
}
