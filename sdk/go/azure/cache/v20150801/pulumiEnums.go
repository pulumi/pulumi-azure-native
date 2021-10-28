


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SkuFamily string

const (
	SkuFamilyC = SkuFamily("C")
	SkuFamilyP = SkuFamily("P")
)

func (SkuFamily) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuFamily)(nil)).Elem()
}

func (e SkuFamily) ToSkuFamilyOutput() SkuFamilyOutput {
	return pulumi.ToOutput(e).(SkuFamilyOutput)
}

func (e SkuFamily) ToSkuFamilyOutputWithContext(ctx context.Context) SkuFamilyOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuFamilyOutput)
}

func (e SkuFamily) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return e.ToSkuFamilyPtrOutputWithContext(context.Background())
}

func (e SkuFamily) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return SkuFamily(e).ToSkuFamilyOutputWithContext(ctx).ToSkuFamilyPtrOutputWithContext(ctx)
}

func (e SkuFamily) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuFamily) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuFamily) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuFamily) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuFamilyOutput struct{ *pulumi.OutputState }

func (SkuFamilyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuFamily)(nil)).Elem()
}

func (o SkuFamilyOutput) ToSkuFamilyOutput() SkuFamilyOutput {
	return o
}

func (o SkuFamilyOutput) ToSkuFamilyOutputWithContext(ctx context.Context) SkuFamilyOutput {
	return o
}

func (o SkuFamilyOutput) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return o.ToSkuFamilyPtrOutputWithContext(context.Background())
}

func (o SkuFamilyOutput) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuFamily) *SkuFamily {
		return &v
	}).(SkuFamilyPtrOutput)
}

func (o SkuFamilyOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuFamilyOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuFamily) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuFamilyOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuFamilyOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuFamily) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuFamilyPtrOutput struct{ *pulumi.OutputState }

func (SkuFamilyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuFamily)(nil)).Elem()
}

func (o SkuFamilyPtrOutput) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return o
}

func (o SkuFamilyPtrOutput) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return o
}

func (o SkuFamilyPtrOutput) Elem() SkuFamilyOutput {
	return o.ApplyT(func(v *SkuFamily) SkuFamily {
		if v != nil {
			return *v
		}
		var ret SkuFamily
		return ret
	}).(SkuFamilyOutput)
}

func (o SkuFamilyPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuFamilyPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuFamily) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuFamilyInput interface {
	pulumi.Input

	ToSkuFamilyOutput() SkuFamilyOutput
	ToSkuFamilyOutputWithContext(context.Context) SkuFamilyOutput
}

var skuFamilyPtrType = reflect.TypeOf((**SkuFamily)(nil)).Elem()

type SkuFamilyPtrInput interface {
	pulumi.Input

	ToSkuFamilyPtrOutput() SkuFamilyPtrOutput
	ToSkuFamilyPtrOutputWithContext(context.Context) SkuFamilyPtrOutput
}

type skuFamilyPtr string

func SkuFamilyPtr(v string) SkuFamilyPtrInput {
	return (*skuFamilyPtr)(&v)
}

func (*skuFamilyPtr) ElementType() reflect.Type {
	return skuFamilyPtrType
}

func (in *skuFamilyPtr) ToSkuFamilyPtrOutput() SkuFamilyPtrOutput {
	return pulumi.ToOutput(in).(SkuFamilyPtrOutput)
}

func (in *skuFamilyPtr) ToSkuFamilyPtrOutputWithContext(ctx context.Context) SkuFamilyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuFamilyPtrOutput)
}

type SkuName string

const (
	SkuNameBasic    = SkuName("Basic")
	SkuNameStandard = SkuName("Standard")
	SkuNamePremium  = SkuName("Premium")
)

func (SkuName) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (e SkuName) ToSkuNameOutput() SkuNameOutput {
	return pulumi.ToOutput(e).(SkuNameOutput)
}

func (e SkuName) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SkuNameOutput)
}

func (e SkuName) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return e.ToSkuNamePtrOutputWithContext(context.Background())
}

func (e SkuName) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return SkuName(e).ToSkuNameOutputWithContext(ctx).ToSkuNamePtrOutputWithContext(ctx)
}

func (e SkuName) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SkuName) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SkuName) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SkuNameOutput struct{ *pulumi.OutputState }

func (SkuNameOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuName)(nil)).Elem()
}

func (o SkuNameOutput) ToSkuNameOutput() SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNameOutputWithContext(ctx context.Context) SkuNameOutput {
	return o
}

func (o SkuNameOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o.ToSkuNamePtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuName) *SkuName {
		return &v
	}).(SkuNamePtrOutput)
}

func (o SkuNameOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SkuNameOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNameOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SkuName) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SkuNamePtrOutput struct{ *pulumi.OutputState }

func (SkuNamePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuName)(nil)).Elem()
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return o
}

func (o SkuNamePtrOutput) Elem() SkuNameOutput {
	return o.ApplyT(func(v *SkuName) SkuName {
		if v != nil {
			return *v
		}
		var ret SkuName
		return ret
	}).(SkuNameOutput)
}

func (o SkuNamePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SkuNamePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SkuName) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SkuNameInput interface {
	pulumi.Input

	ToSkuNameOutput() SkuNameOutput
	ToSkuNameOutputWithContext(context.Context) SkuNameOutput
}

var skuNamePtrType = reflect.TypeOf((**SkuName)(nil)).Elem()

type SkuNamePtrInput interface {
	pulumi.Input

	ToSkuNamePtrOutput() SkuNamePtrOutput
	ToSkuNamePtrOutputWithContext(context.Context) SkuNamePtrOutput
}

type skuNamePtr string

func SkuNamePtr(v string) SkuNamePtrInput {
	return (*skuNamePtr)(&v)
}

func (*skuNamePtr) ElementType() reflect.Type {
	return skuNamePtrType
}

func (in *skuNamePtr) ToSkuNamePtrOutput() SkuNamePtrOutput {
	return pulumi.ToOutput(in).(SkuNamePtrOutput)
}

func (in *skuNamePtr) ToSkuNamePtrOutputWithContext(ctx context.Context) SkuNamePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SkuNamePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SkuFamilyOutput{})
	pulumi.RegisterOutputType(SkuFamilyPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
