


package v20220101preview

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

type ScopeType string

const (
	ScopeTypeCluster   = ScopeType("cluster")
	ScopeTypeNamespace = ScopeType("namespace")
)

func (ScopeType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeType)(nil)).Elem()
}

func (e ScopeType) ToScopeTypeOutput() ScopeTypeOutput {
	return pulumi.ToOutput(e).(ScopeTypeOutput)
}

func (e ScopeType) ToScopeTypeOutputWithContext(ctx context.Context) ScopeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScopeTypeOutput)
}

func (e ScopeType) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return e.ToScopeTypePtrOutputWithContext(context.Background())
}

func (e ScopeType) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return ScopeType(e).ToScopeTypeOutputWithContext(ctx).ToScopeTypePtrOutputWithContext(ctx)
}

func (e ScopeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScopeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScopeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScopeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScopeTypeOutput struct{ *pulumi.OutputState }

func (ScopeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScopeType)(nil)).Elem()
}

func (o ScopeTypeOutput) ToScopeTypeOutput() ScopeTypeOutput {
	return o
}

func (o ScopeTypeOutput) ToScopeTypeOutputWithContext(ctx context.Context) ScopeTypeOutput {
	return o
}

func (o ScopeTypeOutput) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return o.ToScopeTypePtrOutputWithContext(context.Background())
}

func (o ScopeTypeOutput) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScopeType) *ScopeType {
		return &v
	}).(ScopeTypePtrOutput)
}

func (o ScopeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScopeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScopeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScopeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScopeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScopeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScopeTypePtrOutput struct{ *pulumi.OutputState }

func (ScopeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScopeType)(nil)).Elem()
}

func (o ScopeTypePtrOutput) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return o
}

func (o ScopeTypePtrOutput) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return o
}

func (o ScopeTypePtrOutput) Elem() ScopeTypeOutput {
	return o.ApplyT(func(v *ScopeType) ScopeType {
		if v != nil {
			return *v
		}
		var ret ScopeType
		return ret
	}).(ScopeTypeOutput)
}

func (o ScopeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScopeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScopeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScopeTypeInput interface {
	pulumi.Input

	ToScopeTypeOutput() ScopeTypeOutput
	ToScopeTypeOutputWithContext(context.Context) ScopeTypeOutput
}

var scopeTypePtrType = reflect.TypeOf((**ScopeType)(nil)).Elem()

type ScopeTypePtrInput interface {
	pulumi.Input

	ToScopeTypePtrOutput() ScopeTypePtrOutput
	ToScopeTypePtrOutputWithContext(context.Context) ScopeTypePtrOutput
}

type scopeTypePtr string

func ScopeTypePtr(v string) ScopeTypePtrInput {
	return (*scopeTypePtr)(&v)
}

func (*scopeTypePtr) ElementType() reflect.Type {
	return scopeTypePtrType
}

func (in *scopeTypePtr) ToScopeTypePtrOutput() ScopeTypePtrOutput {
	return pulumi.ToOutput(in).(ScopeTypePtrOutput)
}

func (in *scopeTypePtr) ToScopeTypePtrOutputWithContext(ctx context.Context) ScopeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScopeTypePtrOutput)
}

type SourceKindType string

const (
	SourceKindTypeGitRepository = SourceKindType("GitRepository")
	SourceKindTypeBucket        = SourceKindType("Bucket")
)

func (SourceKindType) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceKindType)(nil)).Elem()
}

func (e SourceKindType) ToSourceKindTypeOutput() SourceKindTypeOutput {
	return pulumi.ToOutput(e).(SourceKindTypeOutput)
}

func (e SourceKindType) ToSourceKindTypeOutputWithContext(ctx context.Context) SourceKindTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SourceKindTypeOutput)
}

func (e SourceKindType) ToSourceKindTypePtrOutput() SourceKindTypePtrOutput {
	return e.ToSourceKindTypePtrOutputWithContext(context.Background())
}

func (e SourceKindType) ToSourceKindTypePtrOutputWithContext(ctx context.Context) SourceKindTypePtrOutput {
	return SourceKindType(e).ToSourceKindTypeOutputWithContext(ctx).ToSourceKindTypePtrOutputWithContext(ctx)
}

func (e SourceKindType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceKindType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SourceKindType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SourceKindType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SourceKindTypeOutput struct{ *pulumi.OutputState }

func (SourceKindTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceKindType)(nil)).Elem()
}

func (o SourceKindTypeOutput) ToSourceKindTypeOutput() SourceKindTypeOutput {
	return o
}

func (o SourceKindTypeOutput) ToSourceKindTypeOutputWithContext(ctx context.Context) SourceKindTypeOutput {
	return o
}

func (o SourceKindTypeOutput) ToSourceKindTypePtrOutput() SourceKindTypePtrOutput {
	return o.ToSourceKindTypePtrOutputWithContext(context.Background())
}

func (o SourceKindTypeOutput) ToSourceKindTypePtrOutputWithContext(ctx context.Context) SourceKindTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceKindType) *SourceKindType {
		return &v
	}).(SourceKindTypePtrOutput)
}

func (o SourceKindTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SourceKindTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceKindType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SourceKindTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceKindTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SourceKindType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SourceKindTypePtrOutput struct{ *pulumi.OutputState }

func (SourceKindTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceKindType)(nil)).Elem()
}

func (o SourceKindTypePtrOutput) ToSourceKindTypePtrOutput() SourceKindTypePtrOutput {
	return o
}

func (o SourceKindTypePtrOutput) ToSourceKindTypePtrOutputWithContext(ctx context.Context) SourceKindTypePtrOutput {
	return o
}

func (o SourceKindTypePtrOutput) Elem() SourceKindTypeOutput {
	return o.ApplyT(func(v *SourceKindType) SourceKindType {
		if v != nil {
			return *v
		}
		var ret SourceKindType
		return ret
	}).(SourceKindTypeOutput)
}

func (o SourceKindTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SourceKindTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SourceKindType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SourceKindTypeInput interface {
	pulumi.Input

	ToSourceKindTypeOutput() SourceKindTypeOutput
	ToSourceKindTypeOutputWithContext(context.Context) SourceKindTypeOutput
}

var sourceKindTypePtrType = reflect.TypeOf((**SourceKindType)(nil)).Elem()

type SourceKindTypePtrInput interface {
	pulumi.Input

	ToSourceKindTypePtrOutput() SourceKindTypePtrOutput
	ToSourceKindTypePtrOutputWithContext(context.Context) SourceKindTypePtrOutput
}

type sourceKindTypePtr string

func SourceKindTypePtr(v string) SourceKindTypePtrInput {
	return (*sourceKindTypePtr)(&v)
}

func (*sourceKindTypePtr) ElementType() reflect.Type {
	return sourceKindTypePtrType
}

func (in *sourceKindTypePtr) ToSourceKindTypePtrOutput() SourceKindTypePtrOutput {
	return pulumi.ToOutput(in).(SourceKindTypePtrOutput)
}

func (in *sourceKindTypePtr) ToSourceKindTypePtrOutputWithContext(ctx context.Context) SourceKindTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SourceKindTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LevelTypeOutput{})
	pulumi.RegisterOutputType(LevelTypePtrOutput{})
	pulumi.RegisterOutputType(OperatorScopeTypeOutput{})
	pulumi.RegisterOutputType(OperatorScopeTypePtrOutput{})
	pulumi.RegisterOutputType(OperatorTypeOutput{})
	pulumi.RegisterOutputType(OperatorTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(ScopeTypeOutput{})
	pulumi.RegisterOutputType(ScopeTypePtrOutput{})
	pulumi.RegisterOutputType(SourceKindTypeOutput{})
	pulumi.RegisterOutputType(SourceKindTypePtrOutput{})
}
