


package v20210315preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HostType string

const (
	HostTypeKubernetes = HostType("Kubernetes")
)

func (HostType) ElementType() reflect.Type {
	return reflect.TypeOf((*HostType)(nil)).Elem()
}

func (e HostType) ToHostTypeOutput() HostTypeOutput {
	return pulumi.ToOutput(e).(HostTypeOutput)
}

func (e HostType) ToHostTypeOutputWithContext(ctx context.Context) HostTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(HostTypeOutput)
}

func (e HostType) ToHostTypePtrOutput() HostTypePtrOutput {
	return e.ToHostTypePtrOutputWithContext(context.Background())
}

func (e HostType) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return HostType(e).ToHostTypeOutputWithContext(ctx).ToHostTypePtrOutputWithContext(ctx)
}

func (e HostType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e HostType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e HostType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type HostTypeOutput struct{ *pulumi.OutputState }

func (HostTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostType)(nil)).Elem()
}

func (o HostTypeOutput) ToHostTypeOutput() HostTypeOutput {
	return o
}

func (o HostTypeOutput) ToHostTypeOutputWithContext(ctx context.Context) HostTypeOutput {
	return o
}

func (o HostTypeOutput) ToHostTypePtrOutput() HostTypePtrOutput {
	return o.ToHostTypePtrOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostType) *HostType {
		return &v
	}).(HostTypePtrOutput)
}

func (o HostTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o HostTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e HostType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type HostTypePtrOutput struct{ *pulumi.OutputState }

func (HostTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostType)(nil)).Elem()
}

func (o HostTypePtrOutput) ToHostTypePtrOutput() HostTypePtrOutput {
	return o
}

func (o HostTypePtrOutput) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return o
}

func (o HostTypePtrOutput) Elem() HostTypeOutput {
	return o.ApplyT(func(v *HostType) HostType {
		if v != nil {
			return *v
		}
		var ret HostType
		return ret
	}).(HostTypeOutput)
}

func (o HostTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o HostTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *HostType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type HostTypeInput interface {
	pulumi.Input

	ToHostTypeOutput() HostTypeOutput
	ToHostTypeOutputWithContext(context.Context) HostTypeOutput
}

var hostTypePtrType = reflect.TypeOf((**HostType)(nil)).Elem()

type HostTypePtrInput interface {
	pulumi.Input

	ToHostTypePtrOutput() HostTypePtrOutput
	ToHostTypePtrOutputWithContext(context.Context) HostTypePtrOutput
}

type hostTypePtr string

func HostTypePtr(v string) HostTypePtrInput {
	return (*hostTypePtr)(&v)
}

func (*hostTypePtr) ElementType() reflect.Type {
	return hostTypePtrType
}

func (in *hostTypePtr) ToHostTypePtrOutput() HostTypePtrOutput {
	return pulumi.ToOutput(in).(HostTypePtrOutput)
}

func (in *hostTypePtr) ToHostTypePtrOutputWithContext(ctx context.Context) HostTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(HostTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HostTypeInput)(nil)).Elem(), HostType("Kubernetes"))
	pulumi.RegisterInputType(reflect.TypeOf((*HostTypePtrInput)(nil)).Elem(), HostType("Kubernetes"))
	pulumi.RegisterOutputType(HostTypeOutput{})
	pulumi.RegisterOutputType(HostTypePtrOutput{})
}
