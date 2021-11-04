


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SensorType string

const (
	SensorTypeOt         = SensorType("Ot")
	SensorTypeEnterprise = SensorType("Enterprise")
)

func (SensorType) ElementType() reflect.Type {
	return reflect.TypeOf((*SensorType)(nil)).Elem()
}

func (e SensorType) ToSensorTypeOutput() SensorTypeOutput {
	return pulumi.ToOutput(e).(SensorTypeOutput)
}

func (e SensorType) ToSensorTypeOutputWithContext(ctx context.Context) SensorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SensorTypeOutput)
}

func (e SensorType) ToSensorTypePtrOutput() SensorTypePtrOutput {
	return e.ToSensorTypePtrOutputWithContext(context.Background())
}

func (e SensorType) ToSensorTypePtrOutputWithContext(ctx context.Context) SensorTypePtrOutput {
	return SensorType(e).ToSensorTypeOutputWithContext(ctx).ToSensorTypePtrOutputWithContext(ctx)
}

func (e SensorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SensorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SensorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SensorTypeOutput struct{ *pulumi.OutputState }

func (SensorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SensorType)(nil)).Elem()
}

func (o SensorTypeOutput) ToSensorTypeOutput() SensorTypeOutput {
	return o
}

func (o SensorTypeOutput) ToSensorTypeOutputWithContext(ctx context.Context) SensorTypeOutput {
	return o
}

func (o SensorTypeOutput) ToSensorTypePtrOutput() SensorTypePtrOutput {
	return o.ToSensorTypePtrOutputWithContext(context.Background())
}

func (o SensorTypeOutput) ToSensorTypePtrOutputWithContext(ctx context.Context) SensorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SensorType) *SensorType {
		return &v
	}).(SensorTypePtrOutput)
}

func (o SensorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SensorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SensorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SensorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SensorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SensorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SensorTypePtrOutput struct{ *pulumi.OutputState }

func (SensorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SensorType)(nil)).Elem()
}

func (o SensorTypePtrOutput) ToSensorTypePtrOutput() SensorTypePtrOutput {
	return o
}

func (o SensorTypePtrOutput) ToSensorTypePtrOutputWithContext(ctx context.Context) SensorTypePtrOutput {
	return o
}

func (o SensorTypePtrOutput) Elem() SensorTypeOutput {
	return o.ApplyT(func(v *SensorType) SensorType {
		if v != nil {
			return *v
		}
		var ret SensorType
		return ret
	}).(SensorTypeOutput)
}

func (o SensorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SensorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SensorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SensorTypeInput interface {
	pulumi.Input

	ToSensorTypeOutput() SensorTypeOutput
	ToSensorTypeOutputWithContext(context.Context) SensorTypeOutput
}

var sensorTypePtrType = reflect.TypeOf((**SensorType)(nil)).Elem()

type SensorTypePtrInput interface {
	pulumi.Input

	ToSensorTypePtrOutput() SensorTypePtrOutput
	ToSensorTypePtrOutputWithContext(context.Context) SensorTypePtrOutput
}

type sensorTypePtr string

func SensorTypePtr(v string) SensorTypePtrInput {
	return (*sensorTypePtr)(&v)
}

func (*sensorTypePtr) ElementType() reflect.Type {
	return sensorTypePtrType
}

func (in *sensorTypePtr) ToSensorTypePtrOutput() SensorTypePtrOutput {
	return pulumi.ToOutput(in).(SensorTypePtrOutput)
}

func (in *sensorTypePtr) ToSensorTypePtrOutputWithContext(ctx context.Context) SensorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SensorTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SensorTypeOutput{})
	pulumi.RegisterOutputType(SensorTypePtrOutput{})
}
