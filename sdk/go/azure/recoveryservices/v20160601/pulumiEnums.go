


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type DayOfWeek string

const (
	DayOfWeekSunday    = DayOfWeek("Sunday")
	DayOfWeekMonday    = DayOfWeek("Monday")
	DayOfWeekTuesday   = DayOfWeek("Tuesday")
	DayOfWeekWednesday = DayOfWeek("Wednesday")
	DayOfWeekThursday  = DayOfWeek("Thursday")
	DayOfWeekFriday    = DayOfWeek("Friday")
	DayOfWeekSaturday  = DayOfWeek("Saturday")
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





type DayOfWeekArrayInput interface {
	pulumi.Input

	ToDayOfWeekArrayOutput() DayOfWeekArrayOutput
	ToDayOfWeekArrayOutputWithContext(context.Context) DayOfWeekArrayOutput
}

type DayOfWeekArray []DayOfWeek

func (DayOfWeekArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayOfWeek)(nil)).Elem()
}

func (i DayOfWeekArray) ToDayOfWeekArrayOutput() DayOfWeekArrayOutput {
	return i.ToDayOfWeekArrayOutputWithContext(context.Background())
}

func (i DayOfWeekArray) ToDayOfWeekArrayOutputWithContext(ctx context.Context) DayOfWeekArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayOfWeekArrayOutput)
}

type DayOfWeekArrayOutput struct{ *pulumi.OutputState }

func (DayOfWeekArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayOfWeek)(nil)).Elem()
}

func (o DayOfWeekArrayOutput) ToDayOfWeekArrayOutput() DayOfWeekArrayOutput {
	return o
}

func (o DayOfWeekArrayOutput) ToDayOfWeekArrayOutputWithContext(ctx context.Context) DayOfWeekArrayOutput {
	return o
}

func (o DayOfWeekArrayOutput) Index(i pulumi.IntInput) DayOfWeekOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DayOfWeek {
		return vs[0].([]DayOfWeek)[vs[1].(int)]
	}).(DayOfWeekOutput)
}

type MonthOfYear string

const (
	MonthOfYearInvalid   = MonthOfYear("Invalid")
	MonthOfYearJanuary   = MonthOfYear("January")
	MonthOfYearFebruary  = MonthOfYear("February")
	MonthOfYearMarch     = MonthOfYear("March")
	MonthOfYearApril     = MonthOfYear("April")
	MonthOfYearMay       = MonthOfYear("May")
	MonthOfYearJune      = MonthOfYear("June")
	MonthOfYearJuly      = MonthOfYear("July")
	MonthOfYearAugust    = MonthOfYear("August")
	MonthOfYearSeptember = MonthOfYear("September")
	MonthOfYearOctober   = MonthOfYear("October")
	MonthOfYearNovember  = MonthOfYear("November")
	MonthOfYearDecember  = MonthOfYear("December")
)

func (MonthOfYear) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthOfYear)(nil)).Elem()
}

func (e MonthOfYear) ToMonthOfYearOutput() MonthOfYearOutput {
	return pulumi.ToOutput(e).(MonthOfYearOutput)
}

func (e MonthOfYear) ToMonthOfYearOutputWithContext(ctx context.Context) MonthOfYearOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MonthOfYearOutput)
}

func (e MonthOfYear) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return e.ToMonthOfYearPtrOutputWithContext(context.Background())
}

func (e MonthOfYear) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return MonthOfYear(e).ToMonthOfYearOutputWithContext(ctx).ToMonthOfYearPtrOutputWithContext(ctx)
}

func (e MonthOfYear) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonthOfYear) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MonthOfYear) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MonthOfYear) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MonthOfYearOutput struct{ *pulumi.OutputState }

func (MonthOfYearOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearOutput) ToMonthOfYearOutput() MonthOfYearOutput {
	return o
}

func (o MonthOfYearOutput) ToMonthOfYearOutputWithContext(ctx context.Context) MonthOfYearOutput {
	return o
}

func (o MonthOfYearOutput) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return o.ToMonthOfYearPtrOutputWithContext(context.Background())
}

func (o MonthOfYearOutput) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthOfYear) *MonthOfYear {
		return &v
	}).(MonthOfYearPtrOutput)
}

func (o MonthOfYearOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MonthOfYearOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonthOfYear) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MonthOfYearOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonthOfYearOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MonthOfYear) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MonthOfYearPtrOutput struct{ *pulumi.OutputState }

func (MonthOfYearPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearPtrOutput) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return o
}

func (o MonthOfYearPtrOutput) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return o
}

func (o MonthOfYearPtrOutput) Elem() MonthOfYearOutput {
	return o.ApplyT(func(v *MonthOfYear) MonthOfYear {
		if v != nil {
			return *v
		}
		var ret MonthOfYear
		return ret
	}).(MonthOfYearOutput)
}

func (o MonthOfYearPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MonthOfYearPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MonthOfYear) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MonthOfYearInput interface {
	pulumi.Input

	ToMonthOfYearOutput() MonthOfYearOutput
	ToMonthOfYearOutputWithContext(context.Context) MonthOfYearOutput
}

var monthOfYearPtrType = reflect.TypeOf((**MonthOfYear)(nil)).Elem()

type MonthOfYearPtrInput interface {
	pulumi.Input

	ToMonthOfYearPtrOutput() MonthOfYearPtrOutput
	ToMonthOfYearPtrOutputWithContext(context.Context) MonthOfYearPtrOutput
}

type monthOfYearPtr string

func MonthOfYearPtr(v string) MonthOfYearPtrInput {
	return (*monthOfYearPtr)(&v)
}

func (*monthOfYearPtr) ElementType() reflect.Type {
	return monthOfYearPtrType
}

func (in *monthOfYearPtr) ToMonthOfYearPtrOutput() MonthOfYearPtrOutput {
	return pulumi.ToOutput(in).(MonthOfYearPtrOutput)
}

func (in *monthOfYearPtr) ToMonthOfYearPtrOutputWithContext(ctx context.Context) MonthOfYearPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MonthOfYearPtrOutput)
}





type MonthOfYearArrayInput interface {
	pulumi.Input

	ToMonthOfYearArrayOutput() MonthOfYearArrayOutput
	ToMonthOfYearArrayOutputWithContext(context.Context) MonthOfYearArrayOutput
}

type MonthOfYearArray []MonthOfYear

func (MonthOfYearArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonthOfYear)(nil)).Elem()
}

func (i MonthOfYearArray) ToMonthOfYearArrayOutput() MonthOfYearArrayOutput {
	return i.ToMonthOfYearArrayOutputWithContext(context.Background())
}

func (i MonthOfYearArray) ToMonthOfYearArrayOutputWithContext(ctx context.Context) MonthOfYearArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthOfYearArrayOutput)
}

type MonthOfYearArrayOutput struct{ *pulumi.OutputState }

func (MonthOfYearArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MonthOfYear)(nil)).Elem()
}

func (o MonthOfYearArrayOutput) ToMonthOfYearArrayOutput() MonthOfYearArrayOutput {
	return o
}

func (o MonthOfYearArrayOutput) ToMonthOfYearArrayOutputWithContext(ctx context.Context) MonthOfYearArrayOutput {
	return o
}

func (o MonthOfYearArrayOutput) Index(i pulumi.IntInput) MonthOfYearOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MonthOfYear {
		return vs[0].([]MonthOfYear)[vs[1].(int)]
	}).(MonthOfYearOutput)
}

type ResourceIdentityType string

const (
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	ResourceIdentityTypeNone           = ResourceIdentityType("None")
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

type RetentionDurationType string

const (
	RetentionDurationTypeInvalid = RetentionDurationType("Invalid")
	RetentionDurationTypeDays    = RetentionDurationType("Days")
	RetentionDurationTypeWeeks   = RetentionDurationType("Weeks")
	RetentionDurationTypeMonths  = RetentionDurationType("Months")
	RetentionDurationTypeYears   = RetentionDurationType("Years")
)

func (RetentionDurationType) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationType)(nil)).Elem()
}

func (e RetentionDurationType) ToRetentionDurationTypeOutput() RetentionDurationTypeOutput {
	return pulumi.ToOutput(e).(RetentionDurationTypeOutput)
}

func (e RetentionDurationType) ToRetentionDurationTypeOutputWithContext(ctx context.Context) RetentionDurationTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RetentionDurationTypeOutput)
}

func (e RetentionDurationType) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return e.ToRetentionDurationTypePtrOutputWithContext(context.Background())
}

func (e RetentionDurationType) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return RetentionDurationType(e).ToRetentionDurationTypeOutputWithContext(ctx).ToRetentionDurationTypePtrOutputWithContext(ctx)
}

func (e RetentionDurationType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionDurationType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionDurationType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RetentionDurationType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RetentionDurationTypeOutput struct{ *pulumi.OutputState }

func (RetentionDurationTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationType)(nil)).Elem()
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypeOutput() RetentionDurationTypeOutput {
	return o
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypeOutputWithContext(ctx context.Context) RetentionDurationTypeOutput {
	return o
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return o.ToRetentionDurationTypePtrOutputWithContext(context.Background())
}

func (o RetentionDurationTypeOutput) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionDurationType) *RetentionDurationType {
		return &v
	}).(RetentionDurationTypePtrOutput)
}

func (o RetentionDurationTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RetentionDurationTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionDurationType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RetentionDurationTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionDurationTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionDurationType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RetentionDurationTypePtrOutput struct{ *pulumi.OutputState }

func (RetentionDurationTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDurationType)(nil)).Elem()
}

func (o RetentionDurationTypePtrOutput) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return o
}

func (o RetentionDurationTypePtrOutput) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return o
}

func (o RetentionDurationTypePtrOutput) Elem() RetentionDurationTypeOutput {
	return o.ApplyT(func(v *RetentionDurationType) RetentionDurationType {
		if v != nil {
			return *v
		}
		var ret RetentionDurationType
		return ret
	}).(RetentionDurationTypeOutput)
}

func (o RetentionDurationTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionDurationTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RetentionDurationType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RetentionDurationTypeInput interface {
	pulumi.Input

	ToRetentionDurationTypeOutput() RetentionDurationTypeOutput
	ToRetentionDurationTypeOutputWithContext(context.Context) RetentionDurationTypeOutput
}

var retentionDurationTypePtrType = reflect.TypeOf((**RetentionDurationType)(nil)).Elem()

type RetentionDurationTypePtrInput interface {
	pulumi.Input

	ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput
	ToRetentionDurationTypePtrOutputWithContext(context.Context) RetentionDurationTypePtrOutput
}

type retentionDurationTypePtr string

func RetentionDurationTypePtr(v string) RetentionDurationTypePtrInput {
	return (*retentionDurationTypePtr)(&v)
}

func (*retentionDurationTypePtr) ElementType() reflect.Type {
	return retentionDurationTypePtrType
}

func (in *retentionDurationTypePtr) ToRetentionDurationTypePtrOutput() RetentionDurationTypePtrOutput {
	return pulumi.ToOutput(in).(RetentionDurationTypePtrOutput)
}

func (in *retentionDurationTypePtr) ToRetentionDurationTypePtrOutputWithContext(ctx context.Context) RetentionDurationTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RetentionDurationTypePtrOutput)
}

type RetentionScheduleFormat string

const (
	RetentionScheduleFormatInvalid = RetentionScheduleFormat("Invalid")
	RetentionScheduleFormatDaily   = RetentionScheduleFormat("Daily")
	RetentionScheduleFormatWeekly  = RetentionScheduleFormat("Weekly")
)

func (RetentionScheduleFormat) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionScheduleFormat)(nil)).Elem()
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatOutput() RetentionScheduleFormatOutput {
	return pulumi.ToOutput(e).(RetentionScheduleFormatOutput)
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatOutputWithContext(ctx context.Context) RetentionScheduleFormatOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RetentionScheduleFormatOutput)
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return e.ToRetentionScheduleFormatPtrOutputWithContext(context.Background())
}

func (e RetentionScheduleFormat) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return RetentionScheduleFormat(e).ToRetentionScheduleFormatOutputWithContext(ctx).ToRetentionScheduleFormatPtrOutputWithContext(ctx)
}

func (e RetentionScheduleFormat) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionScheduleFormat) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RetentionScheduleFormat) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RetentionScheduleFormat) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RetentionScheduleFormatOutput struct{ *pulumi.OutputState }

func (RetentionScheduleFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionScheduleFormat)(nil)).Elem()
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatOutput() RetentionScheduleFormatOutput {
	return o
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatOutputWithContext(ctx context.Context) RetentionScheduleFormatOutput {
	return o
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return o.ToRetentionScheduleFormatPtrOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatOutput) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionScheduleFormat) *RetentionScheduleFormat {
		return &v
	}).(RetentionScheduleFormatPtrOutput)
}

func (o RetentionScheduleFormatOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionScheduleFormat) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RetentionScheduleFormatOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RetentionScheduleFormat) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RetentionScheduleFormatPtrOutput struct{ *pulumi.OutputState }

func (RetentionScheduleFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionScheduleFormat)(nil)).Elem()
}

func (o RetentionScheduleFormatPtrOutput) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return o
}

func (o RetentionScheduleFormatPtrOutput) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return o
}

func (o RetentionScheduleFormatPtrOutput) Elem() RetentionScheduleFormatOutput {
	return o.ApplyT(func(v *RetentionScheduleFormat) RetentionScheduleFormat {
		if v != nil {
			return *v
		}
		var ret RetentionScheduleFormat
		return ret
	}).(RetentionScheduleFormatOutput)
}

func (o RetentionScheduleFormatPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RetentionScheduleFormatPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RetentionScheduleFormat) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RetentionScheduleFormatInput interface {
	pulumi.Input

	ToRetentionScheduleFormatOutput() RetentionScheduleFormatOutput
	ToRetentionScheduleFormatOutputWithContext(context.Context) RetentionScheduleFormatOutput
}

var retentionScheduleFormatPtrType = reflect.TypeOf((**RetentionScheduleFormat)(nil)).Elem()

type RetentionScheduleFormatPtrInput interface {
	pulumi.Input

	ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput
	ToRetentionScheduleFormatPtrOutputWithContext(context.Context) RetentionScheduleFormatPtrOutput
}

type retentionScheduleFormatPtr string

func RetentionScheduleFormatPtr(v string) RetentionScheduleFormatPtrInput {
	return (*retentionScheduleFormatPtr)(&v)
}

func (*retentionScheduleFormatPtr) ElementType() reflect.Type {
	return retentionScheduleFormatPtrType
}

func (in *retentionScheduleFormatPtr) ToRetentionScheduleFormatPtrOutput() RetentionScheduleFormatPtrOutput {
	return pulumi.ToOutput(in).(RetentionScheduleFormatPtrOutput)
}

func (in *retentionScheduleFormatPtr) ToRetentionScheduleFormatPtrOutputWithContext(ctx context.Context) RetentionScheduleFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RetentionScheduleFormatPtrOutput)
}

type ScheduleRunType string

const (
	ScheduleRunTypeInvalid = ScheduleRunType("Invalid")
	ScheduleRunTypeDaily   = ScheduleRunType("Daily")
	ScheduleRunTypeWeekly  = ScheduleRunType("Weekly")
)

func (ScheduleRunType) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRunType)(nil)).Elem()
}

func (e ScheduleRunType) ToScheduleRunTypeOutput() ScheduleRunTypeOutput {
	return pulumi.ToOutput(e).(ScheduleRunTypeOutput)
}

func (e ScheduleRunType) ToScheduleRunTypeOutputWithContext(ctx context.Context) ScheduleRunTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ScheduleRunTypeOutput)
}

func (e ScheduleRunType) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return e.ToScheduleRunTypePtrOutputWithContext(context.Background())
}

func (e ScheduleRunType) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return ScheduleRunType(e).ToScheduleRunTypeOutputWithContext(ctx).ToScheduleRunTypePtrOutputWithContext(ctx)
}

func (e ScheduleRunType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleRunType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ScheduleRunType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ScheduleRunType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ScheduleRunTypeOutput struct{ *pulumi.OutputState }

func (ScheduleRunTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleRunType)(nil)).Elem()
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypeOutput() ScheduleRunTypeOutput {
	return o
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypeOutputWithContext(ctx context.Context) ScheduleRunTypeOutput {
	return o
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return o.ToScheduleRunTypePtrOutputWithContext(context.Background())
}

func (o ScheduleRunTypeOutput) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleRunType) *ScheduleRunType {
		return &v
	}).(ScheduleRunTypePtrOutput)
}

func (o ScheduleRunTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ScheduleRunTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleRunType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ScheduleRunTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleRunTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ScheduleRunType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ScheduleRunTypePtrOutput struct{ *pulumi.OutputState }

func (ScheduleRunTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleRunType)(nil)).Elem()
}

func (o ScheduleRunTypePtrOutput) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return o
}

func (o ScheduleRunTypePtrOutput) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return o
}

func (o ScheduleRunTypePtrOutput) Elem() ScheduleRunTypeOutput {
	return o.ApplyT(func(v *ScheduleRunType) ScheduleRunType {
		if v != nil {
			return *v
		}
		var ret ScheduleRunType
		return ret
	}).(ScheduleRunTypeOutput)
}

func (o ScheduleRunTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ScheduleRunTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ScheduleRunType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ScheduleRunTypeInput interface {
	pulumi.Input

	ToScheduleRunTypeOutput() ScheduleRunTypeOutput
	ToScheduleRunTypeOutputWithContext(context.Context) ScheduleRunTypeOutput
}

var scheduleRunTypePtrType = reflect.TypeOf((**ScheduleRunType)(nil)).Elem()

type ScheduleRunTypePtrInput interface {
	pulumi.Input

	ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput
	ToScheduleRunTypePtrOutputWithContext(context.Context) ScheduleRunTypePtrOutput
}

type scheduleRunTypePtr string

func ScheduleRunTypePtr(v string) ScheduleRunTypePtrInput {
	return (*scheduleRunTypePtr)(&v)
}

func (*scheduleRunTypePtr) ElementType() reflect.Type {
	return scheduleRunTypePtrType
}

func (in *scheduleRunTypePtr) ToScheduleRunTypePtrOutput() ScheduleRunTypePtrOutput {
	return pulumi.ToOutput(in).(ScheduleRunTypePtrOutput)
}

func (in *scheduleRunTypePtr) ToScheduleRunTypePtrOutputWithContext(ctx context.Context) ScheduleRunTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ScheduleRunTypePtrOutput)
}

type SkuName string

const (
	SkuNameStandard = SkuName("Standard")
	SkuNameRS0      = SkuName("RS0")
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

type WeekOfMonth string

const (
	WeekOfMonthFirst  = WeekOfMonth("First")
	WeekOfMonthSecond = WeekOfMonth("Second")
	WeekOfMonthThird  = WeekOfMonth("Third")
	WeekOfMonthFourth = WeekOfMonth("Fourth")
	WeekOfMonthLast   = WeekOfMonth("Last")
)

func (WeekOfMonth) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekOfMonth)(nil)).Elem()
}

func (e WeekOfMonth) ToWeekOfMonthOutput() WeekOfMonthOutput {
	return pulumi.ToOutput(e).(WeekOfMonthOutput)
}

func (e WeekOfMonth) ToWeekOfMonthOutputWithContext(ctx context.Context) WeekOfMonthOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WeekOfMonthOutput)
}

func (e WeekOfMonth) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return e.ToWeekOfMonthPtrOutputWithContext(context.Background())
}

func (e WeekOfMonth) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return WeekOfMonth(e).ToWeekOfMonthOutputWithContext(ctx).ToWeekOfMonthPtrOutputWithContext(ctx)
}

func (e WeekOfMonth) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekOfMonth) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WeekOfMonth) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WeekOfMonth) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WeekOfMonthOutput struct{ *pulumi.OutputState }

func (WeekOfMonthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthOutput) ToWeekOfMonthOutput() WeekOfMonthOutput {
	return o
}

func (o WeekOfMonthOutput) ToWeekOfMonthOutputWithContext(ctx context.Context) WeekOfMonthOutput {
	return o
}

func (o WeekOfMonthOutput) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return o.ToWeekOfMonthPtrOutputWithContext(context.Background())
}

func (o WeekOfMonthOutput) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeekOfMonth) *WeekOfMonth {
		return &v
	}).(WeekOfMonthPtrOutput)
}

func (o WeekOfMonthOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WeekOfMonthOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WeekOfMonth) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WeekOfMonthOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WeekOfMonthOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WeekOfMonth) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WeekOfMonthPtrOutput struct{ *pulumi.OutputState }

func (WeekOfMonthPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthPtrOutput) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return o
}

func (o WeekOfMonthPtrOutput) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return o
}

func (o WeekOfMonthPtrOutput) Elem() WeekOfMonthOutput {
	return o.ApplyT(func(v *WeekOfMonth) WeekOfMonth {
		if v != nil {
			return *v
		}
		var ret WeekOfMonth
		return ret
	}).(WeekOfMonthOutput)
}

func (o WeekOfMonthPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WeekOfMonthPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WeekOfMonth) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WeekOfMonthInput interface {
	pulumi.Input

	ToWeekOfMonthOutput() WeekOfMonthOutput
	ToWeekOfMonthOutputWithContext(context.Context) WeekOfMonthOutput
}

var weekOfMonthPtrType = reflect.TypeOf((**WeekOfMonth)(nil)).Elem()

type WeekOfMonthPtrInput interface {
	pulumi.Input

	ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput
	ToWeekOfMonthPtrOutputWithContext(context.Context) WeekOfMonthPtrOutput
}

type weekOfMonthPtr string

func WeekOfMonthPtr(v string) WeekOfMonthPtrInput {
	return (*weekOfMonthPtr)(&v)
}

func (*weekOfMonthPtr) ElementType() reflect.Type {
	return weekOfMonthPtrType
}

func (in *weekOfMonthPtr) ToWeekOfMonthPtrOutput() WeekOfMonthPtrOutput {
	return pulumi.ToOutput(in).(WeekOfMonthPtrOutput)
}

func (in *weekOfMonthPtr) ToWeekOfMonthPtrOutputWithContext(ctx context.Context) WeekOfMonthPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WeekOfMonthPtrOutput)
}





type WeekOfMonthArrayInput interface {
	pulumi.Input

	ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput
	ToWeekOfMonthArrayOutputWithContext(context.Context) WeekOfMonthArrayOutput
}

type WeekOfMonthArray []WeekOfMonth

func (WeekOfMonthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WeekOfMonth)(nil)).Elem()
}

func (i WeekOfMonthArray) ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput {
	return i.ToWeekOfMonthArrayOutputWithContext(context.Background())
}

func (i WeekOfMonthArray) ToWeekOfMonthArrayOutputWithContext(ctx context.Context) WeekOfMonthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeekOfMonthArrayOutput)
}

type WeekOfMonthArrayOutput struct{ *pulumi.OutputState }

func (WeekOfMonthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WeekOfMonth)(nil)).Elem()
}

func (o WeekOfMonthArrayOutput) ToWeekOfMonthArrayOutput() WeekOfMonthArrayOutput {
	return o
}

func (o WeekOfMonthArrayOutput) ToWeekOfMonthArrayOutputWithContext(ctx context.Context) WeekOfMonthArrayOutput {
	return o
}

func (o WeekOfMonthArrayOutput) Index(i pulumi.IntInput) WeekOfMonthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WeekOfMonth {
		return vs[0].([]WeekOfMonth)[vs[1].(int)]
	}).(WeekOfMonthOutput)
}

func init() {
	pulumi.RegisterOutputType(DayOfWeekOutput{})
	pulumi.RegisterOutputType(DayOfWeekPtrOutput{})
	pulumi.RegisterOutputType(DayOfWeekArrayOutput{})
	pulumi.RegisterOutputType(MonthOfYearOutput{})
	pulumi.RegisterOutputType(MonthOfYearPtrOutput{})
	pulumi.RegisterOutputType(MonthOfYearArrayOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
	pulumi.RegisterOutputType(RetentionDurationTypeOutput{})
	pulumi.RegisterOutputType(RetentionDurationTypePtrOutput{})
	pulumi.RegisterOutputType(RetentionScheduleFormatOutput{})
	pulumi.RegisterOutputType(RetentionScheduleFormatPtrOutput{})
	pulumi.RegisterOutputType(ScheduleRunTypeOutput{})
	pulumi.RegisterOutputType(ScheduleRunTypePtrOutput{})
	pulumi.RegisterOutputType(SkuNameOutput{})
	pulumi.RegisterOutputType(SkuNamePtrOutput{})
	pulumi.RegisterOutputType(WeekOfMonthOutput{})
	pulumi.RegisterOutputType(WeekOfMonthPtrOutput{})
	pulumi.RegisterOutputType(WeekOfMonthArrayOutput{})
}
