


package v20160401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DayOfWeek string

const (
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekEveryday  = DayOfWeek("Everyday")
	DayOfWeekWeekend   = DayOfWeek("Weekend")
)

func (DayOfWeek) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (e DayOfWeek) ToDayOfWeekOutput() DayOfWeekOutput {
	return pulumi.ToOutput(e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DayOfWeekOutput)
}

func (e DayOfWeek) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return e.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return DayOfWeek(e).ToDayOfWeekOutputWithContext(ctx).ToDayOfWeekPtrOutputWithContext(ctx)
}

func (e DayOfWeek) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DayOfWeek) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DayOfWeek) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DayOfWeekOutput struct{ *pulumi.OutputState }

func (DayOfWeekOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekOutput) ToDayOfWeekOutput() DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekOutputWithContext(ctx context.Context) DayOfWeekOutput {
	return o
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o.ToDayOfWeekPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DayOfWeek) *DayOfWeek {
		return &v
	}).(DayOfWeekPtrOutput)
}

func (o DayOfWeekOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DayOfWeekOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DayOfWeek) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DayOfWeekPtrOutput struct{ *pulumi.OutputState }

func (DayOfWeekPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return o
}

func (o DayOfWeekPtrOutput) Elem() DayOfWeekOutput {
	return o.ApplyT(func(v *DayOfWeek) DayOfWeek {
		if v != nil {
			return *v
		}
		var ret DayOfWeek
		return ret
	}).(DayOfWeekOutput)
}

func (o DayOfWeekPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DayOfWeekPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DayOfWeek) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DayOfWeekInput interface {
	pulumi.Input

	ToDayOfWeekOutput() DayOfWeekOutput
	ToDayOfWeekOutputWithContext(context.Context) DayOfWeekOutput
}

var dayOfWeekPtrType = reflect.TypeOf((**DayOfWeek)(nil)).Elem()

type DayOfWeekPtrInput interface {
	pulumi.Input

	ToDayOfWeekPtrOutput() DayOfWeekPtrOutput
	ToDayOfWeekPtrOutputWithContext(context.Context) DayOfWeekPtrOutput
}

type dayOfWeekPtr string

func DayOfWeekPtr(v string) DayOfWeekPtrInput {
	return (*dayOfWeekPtr)(&v)
}

func (*dayOfWeekPtr) ElementType() reflect.Type {
	return dayOfWeekPtrType
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutput() DayOfWeekPtrOutput {
	return pulumi.ToOutput(in).(DayOfWeekPtrOutput)
}

func (in *dayOfWeekPtr) ToDayOfWeekPtrOutputWithContext(ctx context.Context) DayOfWeekPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DayOfWeekPtrOutput)
}

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
	pulumi.RegisterInputType(reflect.TypeOf((*DayOfWeekInput)(nil)).Elem(), DayOfWeek("Monday"))
	pulumi.RegisterInputType(reflect.TypeOf((*DayOfWeekPtrInput)(nil)).Elem(), DayOfWeek("Monday"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuFamilyInput)(nil)).Elem(), SkuFamily("C"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuFamilyPtrInput)(nil)).Elem(), SkuFamily("C"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNameInput)(nil)).Elem(), SkuName("Basic"))
	pulumi.RegisterInputType(reflect.TypeOf((*SkuNamePtrInput)(nil)).Elem(), SkuName("Basic"))
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(SkuFamilyOutput{})
	pulumi.RegisterOutputType(SkuFamilyPtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
}
