


package v20210901

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LevelType string

const (
	LevelTypeError       = LevelType("Error")
	LevelTypeWarning     = LevelType("Warning")
	LevelTypeInformation = LevelType("Information")
)

func (LevelType) ElementType() reflect.Type {
	return reflect.TypeOf((*LevelType)(nil)).Elem()
}

func (e LevelType) ToLevelTypeOutput() LevelTypeOutput {
	return pulumi.ToOutput(e).(LevelTypeOutput)
}

func (e LevelType) ToLevelTypeOutputWithContext(ctx context.Context) LevelTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LevelTypeOutput)
}

func (e LevelType) ToLevelTypePtrOutput() LevelTypePtrOutput {
	return e.ToLevelTypePtrOutputWithContext(context.Background())
}

func (e LevelType) ToLevelTypePtrOutputWithContext(ctx context.Context) LevelTypePtrOutput {
	return LevelType(e).ToLevelTypeOutputWithContext(ctx).ToLevelTypePtrOutputWithContext(ctx)
}

func (e LevelType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LevelType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LevelType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LevelType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LevelTypeOutput struct{ *pulumi.OutputState }

func (LevelTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LevelType)(nil)).Elem()
}

func (o LevelTypeOutput) ToLevelTypeOutput() LevelTypeOutput {
	return o
}

func (o LevelTypeOutput) ToLevelTypeOutputWithContext(ctx context.Context) LevelTypeOutput {
	return o
}

func (o LevelTypeOutput) ToLevelTypePtrOutput() LevelTypePtrOutput {
	return o.ToLevelTypePtrOutputWithContext(context.Background())
}

func (o LevelTypeOutput) ToLevelTypePtrOutputWithContext(ctx context.Context) LevelTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LevelType) *LevelType {
		return &v
	}).(LevelTypePtrOutput)
}

func (o LevelTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LevelTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LevelType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LevelTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LevelTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LevelType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LevelTypePtrOutput struct{ *pulumi.OutputState }

func (LevelTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LevelType)(nil)).Elem()
}

func (o LevelTypePtrOutput) ToLevelTypePtrOutput() LevelTypePtrOutput {
	return o
}

func (o LevelTypePtrOutput) ToLevelTypePtrOutputWithContext(ctx context.Context) LevelTypePtrOutput {
	return o
}

func (o LevelTypePtrOutput) Elem() LevelTypeOutput {
	return o.ApplyT(func(v *LevelType) LevelType {
		if v != nil {
			return *v
		}
		var ret LevelType
		return ret
	}).(LevelTypeOutput)
}

func (o LevelTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LevelTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LevelType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LevelTypeInput interface {
	pulumi.Input

	ToLevelTypeOutput() LevelTypeOutput
	ToLevelTypeOutputWithContext(context.Context) LevelTypeOutput
}

var levelTypePtrType = reflect.TypeOf((**LevelType)(nil)).Elem()

type LevelTypePtrInput interface {
	pulumi.Input

	ToLevelTypePtrOutput() LevelTypePtrOutput
	ToLevelTypePtrOutputWithContext(context.Context) LevelTypePtrOutput
}

type levelTypePtr string

func LevelTypePtr(v string) LevelTypePtrInput {
	return (*levelTypePtr)(&v)
}

func (*levelTypePtr) ElementType() reflect.Type {
	return levelTypePtrType
}

func (in *levelTypePtr) ToLevelTypePtrOutput() LevelTypePtrOutput {
	return pulumi.ToOutput(in).(LevelTypePtrOutput)
}

func (in *levelTypePtr) ToLevelTypePtrOutputWithContext(ctx context.Context) LevelTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LevelTypePtrOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LevelTypeOutput{})
	pulumi.RegisterOutputType(LevelTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
