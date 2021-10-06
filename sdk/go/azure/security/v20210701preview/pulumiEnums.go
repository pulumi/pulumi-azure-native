


package v20210701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

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

func init() {
	pulumi.RegisterOutputType(SeverityEnumOutput{})
	pulumi.RegisterOutputType(SeverityEnumPtrOutput{})
	pulumi.RegisterOutputType(SupportedCloudEnumOutput{})
	pulumi.RegisterOutputType(SupportedCloudEnumPtrOutput{})
}
