


package v20201001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type OperatorScopeType string

const (
	OperatorScopeTypeCluster   = OperatorScopeType("cluster")
	OperatorScopeTypeNamespace = OperatorScopeType("namespace")
)

func (OperatorScopeType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorScopeType)(nil)).Elem()
}

func (e OperatorScopeType) ToOperatorScopeTypeOutput() OperatorScopeTypeOutput {
	return pulumi.ToOutput(e).(OperatorScopeTypeOutput)
}

func (e OperatorScopeType) ToOperatorScopeTypeOutputWithContext(ctx context.Context) OperatorScopeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorScopeTypeOutput)
}

func (e OperatorScopeType) ToOperatorScopeTypePtrOutput() OperatorScopeTypePtrOutput {
	return e.ToOperatorScopeTypePtrOutputWithContext(context.Background())
}

func (e OperatorScopeType) ToOperatorScopeTypePtrOutputWithContext(ctx context.Context) OperatorScopeTypePtrOutput {
	return OperatorScopeType(e).ToOperatorScopeTypeOutputWithContext(ctx).ToOperatorScopeTypePtrOutputWithContext(ctx)
}

func (e OperatorScopeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorScopeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorScopeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatorScopeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorScopeTypeOutput struct{ *pulumi.OutputState }

func (OperatorScopeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorScopeType)(nil)).Elem()
}

func (o OperatorScopeTypeOutput) ToOperatorScopeTypeOutput() OperatorScopeTypeOutput {
	return o
}

func (o OperatorScopeTypeOutput) ToOperatorScopeTypeOutputWithContext(ctx context.Context) OperatorScopeTypeOutput {
	return o
}

func (o OperatorScopeTypeOutput) ToOperatorScopeTypePtrOutput() OperatorScopeTypePtrOutput {
	return o.ToOperatorScopeTypePtrOutputWithContext(context.Background())
}

func (o OperatorScopeTypeOutput) ToOperatorScopeTypePtrOutputWithContext(ctx context.Context) OperatorScopeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatorScopeType) *OperatorScopeType {
		return &v
	}).(OperatorScopeTypePtrOutput)
}

func (o OperatorScopeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorScopeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorScopeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorScopeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorScopeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorScopeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorScopeTypePtrOutput struct{ *pulumi.OutputState }

func (OperatorScopeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatorScopeType)(nil)).Elem()
}

func (o OperatorScopeTypePtrOutput) ToOperatorScopeTypePtrOutput() OperatorScopeTypePtrOutput {
	return o
}

func (o OperatorScopeTypePtrOutput) ToOperatorScopeTypePtrOutputWithContext(ctx context.Context) OperatorScopeTypePtrOutput {
	return o
}

func (o OperatorScopeTypePtrOutput) Elem() OperatorScopeTypeOutput {
	return o.ApplyT(func(v *OperatorScopeType) OperatorScopeType {
		if v != nil {
			return *v
		}
		var ret OperatorScopeType
		return ret
	}).(OperatorScopeTypeOutput)
}

func (o OperatorScopeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorScopeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatorScopeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatorScopeTypeInput interface {
	pulumi.Input

	ToOperatorScopeTypeOutput() OperatorScopeTypeOutput
	ToOperatorScopeTypeOutputWithContext(context.Context) OperatorScopeTypeOutput
}

var operatorScopeTypePtrType = reflect.TypeOf((**OperatorScopeType)(nil)).Elem()

type OperatorScopeTypePtrInput interface {
	pulumi.Input

	ToOperatorScopeTypePtrOutput() OperatorScopeTypePtrOutput
	ToOperatorScopeTypePtrOutputWithContext(context.Context) OperatorScopeTypePtrOutput
}

type operatorScopeTypePtr string

func OperatorScopeTypePtr(v string) OperatorScopeTypePtrInput {
	return (*operatorScopeTypePtr)(&v)
}

func (*operatorScopeTypePtr) ElementType() reflect.Type {
	return operatorScopeTypePtrType
}

func (in *operatorScopeTypePtr) ToOperatorScopeTypePtrOutput() OperatorScopeTypePtrOutput {
	return pulumi.ToOutput(in).(OperatorScopeTypePtrOutput)
}

func (in *operatorScopeTypePtr) ToOperatorScopeTypePtrOutputWithContext(ctx context.Context) OperatorScopeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorScopeTypePtrOutput)
}

type OperatorType string

const (
	OperatorTypeFlux = OperatorType("Flux")
)

func (OperatorType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorType)(nil)).Elem()
}

func (e OperatorType) ToOperatorTypeOutput() OperatorTypeOutput {
	return pulumi.ToOutput(e).(OperatorTypeOutput)
}

func (e OperatorType) ToOperatorTypeOutputWithContext(ctx context.Context) OperatorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorTypeOutput)
}

func (e OperatorType) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return e.ToOperatorTypePtrOutputWithContext(context.Background())
}

func (e OperatorType) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return OperatorType(e).ToOperatorTypeOutputWithContext(ctx).ToOperatorTypePtrOutputWithContext(ctx)
}

func (e OperatorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorTypeOutput struct{ *pulumi.OutputState }

func (OperatorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorType)(nil)).Elem()
}

func (o OperatorTypeOutput) ToOperatorTypeOutput() OperatorTypeOutput {
	return o
}

func (o OperatorTypeOutput) ToOperatorTypeOutputWithContext(ctx context.Context) OperatorTypeOutput {
	return o
}

func (o OperatorTypeOutput) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return o.ToOperatorTypePtrOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatorType) *OperatorType {
		return &v
	}).(OperatorTypePtrOutput)
}

func (o OperatorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorTypePtrOutput struct{ *pulumi.OutputState }

func (OperatorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatorType)(nil)).Elem()
}

func (o OperatorTypePtrOutput) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return o
}

func (o OperatorTypePtrOutput) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return o
}

func (o OperatorTypePtrOutput) Elem() OperatorTypeOutput {
	return o.ApplyT(func(v *OperatorType) OperatorType {
		if v != nil {
			return *v
		}
		var ret OperatorType
		return ret
	}).(OperatorTypeOutput)
}

func (o OperatorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatorTypeInput interface {
	pulumi.Input

	ToOperatorTypeOutput() OperatorTypeOutput
	ToOperatorTypeOutputWithContext(context.Context) OperatorTypeOutput
}

var operatorTypePtrType = reflect.TypeOf((**OperatorType)(nil)).Elem()

type OperatorTypePtrInput interface {
	pulumi.Input

	ToOperatorTypePtrOutput() OperatorTypePtrOutput
	ToOperatorTypePtrOutputWithContext(context.Context) OperatorTypePtrOutput
}

type operatorTypePtr string

func OperatorTypePtr(v string) OperatorTypePtrInput {
	return (*operatorTypePtr)(&v)
}

func (*operatorTypePtr) ElementType() reflect.Type {
	return operatorTypePtrType
}

func (in *operatorTypePtr) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return pulumi.ToOutput(in).(OperatorTypePtrOutput)
}

func (in *operatorTypePtr) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorScopeTypeInput)(nil)).Elem(), OperatorScopeType("cluster"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorScopeTypePtrInput)(nil)).Elem(), OperatorScopeType("cluster"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorTypeInput)(nil)).Elem(), OperatorType("Flux"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorTypePtrInput)(nil)).Elem(), OperatorType("Flux"))
	pulumi.RegisterOutputType(OperatorScopeTypeOutput{})
	pulumi.RegisterOutputType(OperatorScopeTypePtrOutput{})
	pulumi.RegisterOutputType(OperatorTypeOutput{})
	pulumi.RegisterOutputType(OperatorTypePtrOutput{})
}
