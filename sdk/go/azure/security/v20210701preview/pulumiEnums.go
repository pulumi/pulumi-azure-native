


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ImplementationEffortEnum string

const (
	ImplementationEffortEnumHigh     = ImplementationEffortEnum("High")
	ImplementationEffortEnumModerate = ImplementationEffortEnum("Moderate")
	ImplementationEffortEnumLow      = ImplementationEffortEnum("Low")
)

func (ImplementationEffortEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*ImplementationEffortEnum)(nil)).Elem()
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumOutput() ImplementationEffortEnumOutput {
	return pulumi.ToOutput(e).(ImplementationEffortEnumOutput)
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumOutputWithContext(ctx context.Context) ImplementationEffortEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ImplementationEffortEnumOutput)
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return e.ToImplementationEffortEnumPtrOutputWithContext(context.Background())
}

func (e ImplementationEffortEnum) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return ImplementationEffortEnum(e).ToImplementationEffortEnumOutputWithContext(ctx).ToImplementationEffortEnumPtrOutputWithContext(ctx)
}

func (e ImplementationEffortEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffortEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ImplementationEffortEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ImplementationEffortEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ImplementationEffortEnumOutput struct{ *pulumi.OutputState }

func (ImplementationEffortEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImplementationEffortEnum)(nil)).Elem()
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumOutput() ImplementationEffortEnumOutput {
	return o
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumOutputWithContext(ctx context.Context) ImplementationEffortEnumOutput {
	return o
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return o.ToImplementationEffortEnumPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumOutput) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImplementationEffortEnum) *ImplementationEffortEnum {
		return &v
	}).(ImplementationEffortEnumPtrOutput)
}

func (o ImplementationEffortEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImplementationEffortEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ImplementationEffortEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ImplementationEffortEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ImplementationEffortEnumPtrOutput struct{ *pulumi.OutputState }

func (ImplementationEffortEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImplementationEffortEnum)(nil)).Elem()
}

func (o ImplementationEffortEnumPtrOutput) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return o
}

func (o ImplementationEffortEnumPtrOutput) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return o
}

func (o ImplementationEffortEnumPtrOutput) Elem() ImplementationEffortEnumOutput {
	return o.ApplyT(func(v *ImplementationEffortEnum) ImplementationEffortEnum {
		if v != nil {
			return *v
		}
		var ret ImplementationEffortEnum
		return ret
	}).(ImplementationEffortEnumOutput)
}

func (o ImplementationEffortEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ImplementationEffortEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ImplementationEffortEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ImplementationEffortEnumInput interface {
	pulumi.Input

	ToImplementationEffortEnumOutput() ImplementationEffortEnumOutput
	ToImplementationEffortEnumOutputWithContext(context.Context) ImplementationEffortEnumOutput
}

var implementationEffortEnumPtrType = reflect.TypeOf((**ImplementationEffortEnum)(nil)).Elem()

type ImplementationEffortEnumPtrInput interface {
	pulumi.Input

	ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput
	ToImplementationEffortEnumPtrOutputWithContext(context.Context) ImplementationEffortEnumPtrOutput
}

type implementationEffortEnumPtr string

func ImplementationEffortEnumPtr(v string) ImplementationEffortEnumPtrInput {
	return (*implementationEffortEnumPtr)(&v)
}

func (*implementationEffortEnumPtr) ElementType() reflect.Type {
	return implementationEffortEnumPtrType
}

func (in *implementationEffortEnumPtr) ToImplementationEffortEnumPtrOutput() ImplementationEffortEnumPtrOutput {
	return pulumi.ToOutput(in).(ImplementationEffortEnumPtrOutput)
}

func (in *implementationEffortEnumPtr) ToImplementationEffortEnumPtrOutputWithContext(ctx context.Context) ImplementationEffortEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ImplementationEffortEnumPtrOutput)
}

type SeverityEnum string

const (
	SeverityEnumHigh   = SeverityEnum("High")
	SeverityEnumMedium = SeverityEnum("Medium")
	SeverityEnumLow    = SeverityEnum("Low")
)

func (SeverityEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SeverityEnum)(nil)).Elem()
}

func (e SeverityEnum) ToSeverityEnumOutput() SeverityEnumOutput {
	return pulumi.ToOutput(e).(SeverityEnumOutput)
}

func (e SeverityEnum) ToSeverityEnumOutputWithContext(ctx context.Context) SeverityEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SeverityEnumOutput)
}

func (e SeverityEnum) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return e.ToSeverityEnumPtrOutputWithContext(context.Background())
}

func (e SeverityEnum) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return SeverityEnum(e).ToSeverityEnumOutputWithContext(ctx).ToSeverityEnumPtrOutputWithContext(ctx)
}

func (e SeverityEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SeverityEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SeverityEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SeverityEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SeverityEnumOutput struct{ *pulumi.OutputState }

func (SeverityEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SeverityEnum)(nil)).Elem()
}

func (o SeverityEnumOutput) ToSeverityEnumOutput() SeverityEnumOutput {
	return o
}

func (o SeverityEnumOutput) ToSeverityEnumOutputWithContext(ctx context.Context) SeverityEnumOutput {
	return o
}

func (o SeverityEnumOutput) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return o.ToSeverityEnumPtrOutputWithContext(context.Background())
}

func (o SeverityEnumOutput) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SeverityEnum) *SeverityEnum {
		return &v
	}).(SeverityEnumPtrOutput)
}

func (o SeverityEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SeverityEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SeverityEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SeverityEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SeverityEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SeverityEnumPtrOutput struct{ *pulumi.OutputState }

func (SeverityEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SeverityEnum)(nil)).Elem()
}

func (o SeverityEnumPtrOutput) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return o
}

func (o SeverityEnumPtrOutput) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return o
}

func (o SeverityEnumPtrOutput) Elem() SeverityEnumOutput {
	return o.ApplyT(func(v *SeverityEnum) SeverityEnum {
		if v != nil {
			return *v
		}
		var ret SeverityEnum
		return ret
	}).(SeverityEnumOutput)
}

func (o SeverityEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SeverityEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SeverityEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SeverityEnumInput interface {
	pulumi.Input

	ToSeverityEnumOutput() SeverityEnumOutput
	ToSeverityEnumOutputWithContext(context.Context) SeverityEnumOutput
}

var severityEnumPtrType = reflect.TypeOf((**SeverityEnum)(nil)).Elem()

type SeverityEnumPtrInput interface {
	pulumi.Input

	ToSeverityEnumPtrOutput() SeverityEnumPtrOutput
	ToSeverityEnumPtrOutputWithContext(context.Context) SeverityEnumPtrOutput
}

type severityEnumPtr string

func SeverityEnumPtr(v string) SeverityEnumPtrInput {
	return (*severityEnumPtr)(&v)
}

func (*severityEnumPtr) ElementType() reflect.Type {
	return severityEnumPtrType
}

func (in *severityEnumPtr) ToSeverityEnumPtrOutput() SeverityEnumPtrOutput {
	return pulumi.ToOutput(in).(SeverityEnumPtrOutput)
}

func (in *severityEnumPtr) ToSeverityEnumPtrOutputWithContext(ctx context.Context) SeverityEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SeverityEnumPtrOutput)
}

type SupportedCloudEnum string

const (
	SupportedCloudEnumAWS = SupportedCloudEnum("AWS")
)

func (SupportedCloudEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCloudEnum)(nil)).Elem()
}

func (e SupportedCloudEnum) ToSupportedCloudEnumOutput() SupportedCloudEnumOutput {
	return pulumi.ToOutput(e).(SupportedCloudEnumOutput)
}

func (e SupportedCloudEnum) ToSupportedCloudEnumOutputWithContext(ctx context.Context) SupportedCloudEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SupportedCloudEnumOutput)
}

func (e SupportedCloudEnum) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return e.ToSupportedCloudEnumPtrOutputWithContext(context.Background())
}

func (e SupportedCloudEnum) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return SupportedCloudEnum(e).ToSupportedCloudEnumOutputWithContext(ctx).ToSupportedCloudEnumPtrOutputWithContext(ctx)
}

func (e SupportedCloudEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedCloudEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SupportedCloudEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SupportedCloudEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SupportedCloudEnumOutput struct{ *pulumi.OutputState }

func (SupportedCloudEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SupportedCloudEnum)(nil)).Elem()
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumOutput() SupportedCloudEnumOutput {
	return o
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumOutputWithContext(ctx context.Context) SupportedCloudEnumOutput {
	return o
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return o.ToSupportedCloudEnumPtrOutputWithContext(context.Background())
}

func (o SupportedCloudEnumOutput) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SupportedCloudEnum) *SupportedCloudEnum {
		return &v
	}).(SupportedCloudEnumPtrOutput)
}

func (o SupportedCloudEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SupportedCloudEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedCloudEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SupportedCloudEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedCloudEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SupportedCloudEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SupportedCloudEnumPtrOutput struct{ *pulumi.OutputState }

func (SupportedCloudEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SupportedCloudEnum)(nil)).Elem()
}

func (o SupportedCloudEnumPtrOutput) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return o
}

func (o SupportedCloudEnumPtrOutput) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return o
}

func (o SupportedCloudEnumPtrOutput) Elem() SupportedCloudEnumOutput {
	return o.ApplyT(func(v *SupportedCloudEnum) SupportedCloudEnum {
		if v != nil {
			return *v
		}
		var ret SupportedCloudEnum
		return ret
	}).(SupportedCloudEnumOutput)
}

func (o SupportedCloudEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SupportedCloudEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SupportedCloudEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SupportedCloudEnumInput interface {
	pulumi.Input

	ToSupportedCloudEnumOutput() SupportedCloudEnumOutput
	ToSupportedCloudEnumOutputWithContext(context.Context) SupportedCloudEnumOutput
}

var supportedCloudEnumPtrType = reflect.TypeOf((**SupportedCloudEnum)(nil)).Elem()

type SupportedCloudEnumPtrInput interface {
	pulumi.Input

	ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput
	ToSupportedCloudEnumPtrOutputWithContext(context.Context) SupportedCloudEnumPtrOutput
}

type supportedCloudEnumPtr string

func SupportedCloudEnumPtr(v string) SupportedCloudEnumPtrInput {
	return (*supportedCloudEnumPtr)(&v)
}

func (*supportedCloudEnumPtr) ElementType() reflect.Type {
	return supportedCloudEnumPtrType
}

func (in *supportedCloudEnumPtr) ToSupportedCloudEnumPtrOutput() SupportedCloudEnumPtrOutput {
	return pulumi.ToOutput(in).(SupportedCloudEnumPtrOutput)
}

func (in *supportedCloudEnumPtr) ToSupportedCloudEnumPtrOutputWithContext(ctx context.Context) SupportedCloudEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SupportedCloudEnumPtrOutput)
}

type UserImpactEnum string

const (
	UserImpactEnumHigh     = UserImpactEnum("High")
	UserImpactEnumModerate = UserImpactEnum("Moderate")
	UserImpactEnumLow      = UserImpactEnum("Low")
)

func (UserImpactEnum) ElementType() reflect.Type {
	return reflect.TypeOf((*UserImpactEnum)(nil)).Elem()
}

func (e UserImpactEnum) ToUserImpactEnumOutput() UserImpactEnumOutput {
	return pulumi.ToOutput(e).(UserImpactEnumOutput)
}

func (e UserImpactEnum) ToUserImpactEnumOutputWithContext(ctx context.Context) UserImpactEnumOutput {
	return pulumi.ToOutputWithContext(ctx, e).(UserImpactEnumOutput)
}

func (e UserImpactEnum) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return e.ToUserImpactEnumPtrOutputWithContext(context.Background())
}

func (e UserImpactEnum) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return UserImpactEnum(e).ToUserImpactEnumOutputWithContext(ctx).ToUserImpactEnumPtrOutputWithContext(ctx)
}

func (e UserImpactEnum) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpactEnum) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e UserImpactEnum) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e UserImpactEnum) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type UserImpactEnumOutput struct{ *pulumi.OutputState }

func (UserImpactEnumOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserImpactEnum)(nil)).Elem()
}

func (o UserImpactEnumOutput) ToUserImpactEnumOutput() UserImpactEnumOutput {
	return o
}

func (o UserImpactEnumOutput) ToUserImpactEnumOutputWithContext(ctx context.Context) UserImpactEnumOutput {
	return o
}

func (o UserImpactEnumOutput) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return o.ToUserImpactEnumPtrOutputWithContext(context.Background())
}

func (o UserImpactEnumOutput) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UserImpactEnum) *UserImpactEnum {
		return &v
	}).(UserImpactEnumPtrOutput)
}

func (o UserImpactEnumOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o UserImpactEnumOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserImpactEnum) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o UserImpactEnumOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserImpactEnumOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e UserImpactEnum) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type UserImpactEnumPtrOutput struct{ *pulumi.OutputState }

func (UserImpactEnumPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UserImpactEnum)(nil)).Elem()
}

func (o UserImpactEnumPtrOutput) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return o
}

func (o UserImpactEnumPtrOutput) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return o
}

func (o UserImpactEnumPtrOutput) Elem() UserImpactEnumOutput {
	return o.ApplyT(func(v *UserImpactEnum) UserImpactEnum {
		if v != nil {
			return *v
		}
		var ret UserImpactEnum
		return ret
	}).(UserImpactEnumOutput)
}

func (o UserImpactEnumPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o UserImpactEnumPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *UserImpactEnum) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type UserImpactEnumInput interface {
	pulumi.Input

	ToUserImpactEnumOutput() UserImpactEnumOutput
	ToUserImpactEnumOutputWithContext(context.Context) UserImpactEnumOutput
}

var userImpactEnumPtrType = reflect.TypeOf((**UserImpactEnum)(nil)).Elem()

type UserImpactEnumPtrInput interface {
	pulumi.Input

	ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput
	ToUserImpactEnumPtrOutputWithContext(context.Context) UserImpactEnumPtrOutput
}

type userImpactEnumPtr string

func UserImpactEnumPtr(v string) UserImpactEnumPtrInput {
	return (*userImpactEnumPtr)(&v)
}

func (*userImpactEnumPtr) ElementType() reflect.Type {
	return userImpactEnumPtrType
}

func (in *userImpactEnumPtr) ToUserImpactEnumPtrOutput() UserImpactEnumPtrOutput {
	return pulumi.ToOutput(in).(UserImpactEnumPtrOutput)
}

func (in *userImpactEnumPtr) ToUserImpactEnumPtrOutputWithContext(ctx context.Context) UserImpactEnumPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(UserImpactEnumPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ImplementationEffortEnumOutput{})
	pulumi.RegisterOutputType(ImplementationEffortEnumPtrOutput{})
	pulumi.RegisterOutputType(SeverityEnumOutput{})
	pulumi.RegisterOutputType(SeverityEnumPtrOutput{})
	pulumi.RegisterOutputType(SupportedCloudEnumOutput{})
	pulumi.RegisterOutputType(SupportedCloudEnumPtrOutput{})
	pulumi.RegisterOutputType(UserImpactEnumOutput{})
	pulumi.RegisterOutputType(UserImpactEnumPtrOutput{})
}
