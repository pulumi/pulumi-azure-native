


package hybridconnectivity

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CreatedByType string

const (
	CreatedByTypeUser            = CreatedByType("User")
	CreatedByTypeApplication     = CreatedByType("Application")
	CreatedByTypeManagedIdentity = CreatedByType("ManagedIdentity")
	CreatedByTypeKey             = CreatedByType("Key")
)

func (CreatedByType) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatedByType)(nil)).Elem()
}

func (e CreatedByType) ToCreatedByTypeOutput() CreatedByTypeOutput {
	return pulumi.ToOutput(e).(CreatedByTypeOutput)
}

func (e CreatedByType) ToCreatedByTypeOutputWithContext(ctx context.Context) CreatedByTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CreatedByTypeOutput)
}

func (e CreatedByType) ToCreatedByTypePtrOutput() CreatedByTypePtrOutput {
	return e.ToCreatedByTypePtrOutputWithContext(context.Background())
}

func (e CreatedByType) ToCreatedByTypePtrOutputWithContext(ctx context.Context) CreatedByTypePtrOutput {
	return CreatedByType(e).ToCreatedByTypeOutputWithContext(ctx).ToCreatedByTypePtrOutputWithContext(ctx)
}

func (e CreatedByType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreatedByType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CreatedByType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CreatedByType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CreatedByTypeOutput struct{ *pulumi.OutputState }

func (CreatedByTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreatedByType)(nil)).Elem()
}

func (o CreatedByTypeOutput) ToCreatedByTypeOutput() CreatedByTypeOutput {
	return o
}

func (o CreatedByTypeOutput) ToCreatedByTypeOutputWithContext(ctx context.Context) CreatedByTypeOutput {
	return o
}

func (o CreatedByTypeOutput) ToCreatedByTypePtrOutput() CreatedByTypePtrOutput {
	return o.ToCreatedByTypePtrOutputWithContext(context.Background())
}

func (o CreatedByTypeOutput) ToCreatedByTypePtrOutputWithContext(ctx context.Context) CreatedByTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreatedByType) *CreatedByType {
		return &v
	}).(CreatedByTypePtrOutput)
}

func (o CreatedByTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CreatedByTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreatedByType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CreatedByTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreatedByTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CreatedByType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CreatedByTypePtrOutput struct{ *pulumi.OutputState }

func (CreatedByTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreatedByType)(nil)).Elem()
}

func (o CreatedByTypePtrOutput) ToCreatedByTypePtrOutput() CreatedByTypePtrOutput {
	return o
}

func (o CreatedByTypePtrOutput) ToCreatedByTypePtrOutputWithContext(ctx context.Context) CreatedByTypePtrOutput {
	return o
}

func (o CreatedByTypePtrOutput) Elem() CreatedByTypeOutput {
	return o.ApplyT(func(v *CreatedByType) CreatedByType {
		if v != nil {
			return *v
		}
		var ret CreatedByType
		return ret
	}).(CreatedByTypeOutput)
}

func (o CreatedByTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CreatedByTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CreatedByType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CreatedByTypeInput interface {
	pulumi.Input

	ToCreatedByTypeOutput() CreatedByTypeOutput
	ToCreatedByTypeOutputWithContext(context.Context) CreatedByTypeOutput
}

var createdByTypePtrType = reflect.TypeOf((**CreatedByType)(nil)).Elem()

type CreatedByTypePtrInput interface {
	pulumi.Input

	ToCreatedByTypePtrOutput() CreatedByTypePtrOutput
	ToCreatedByTypePtrOutputWithContext(context.Context) CreatedByTypePtrOutput
}

type createdByTypePtr string

func CreatedByTypePtr(v string) CreatedByTypePtrInput {
	return (*createdByTypePtr)(&v)
}

func (*createdByTypePtr) ElementType() reflect.Type {
	return createdByTypePtrType
}

func (in *createdByTypePtr) ToCreatedByTypePtrOutput() CreatedByTypePtrOutput {
	return pulumi.ToOutput(in).(CreatedByTypePtrOutput)
}

func (in *createdByTypePtr) ToCreatedByTypePtrOutputWithContext(ctx context.Context) CreatedByTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CreatedByTypePtrOutput)
}

type Type string

const (
	TypeDefault = Type("default")
	TypeCustom  = Type("custom")
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
	pulumi.RegisterOutputType(CreatedByTypeOutput{})
	pulumi.RegisterOutputType(CreatedByTypePtrOutput{})
	pulumi.RegisterOutputType(TypeOutput{})
	pulumi.RegisterOutputType(TypePtrOutput{})
}
