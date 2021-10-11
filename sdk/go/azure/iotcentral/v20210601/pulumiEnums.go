


package v20210601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSku string

const (
	AppSkuST0 = AppSku("ST0")
	AppSkuST1 = AppSku("ST1")
	AppSkuST2 = AppSku("ST2")
)

func (AppSku) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSku)(nil)).Elem()
}

func (e AppSku) ToAppSkuOutput() AppSkuOutput {
	return pulumi.ToOutput(e).(AppSkuOutput)
}

func (e AppSku) ToAppSkuOutputWithContext(ctx context.Context) AppSkuOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AppSkuOutput)
}

func (e AppSku) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return e.ToAppSkuPtrOutputWithContext(context.Background())
}

func (e AppSku) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return AppSku(e).ToAppSkuOutputWithContext(ctx).ToAppSkuPtrOutputWithContext(ctx)
}

func (e AppSku) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppSku) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AppSku) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AppSku) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AppSkuOutput struct{ *pulumi.OutputState }

func (AppSkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppSku)(nil)).Elem()
}

func (o AppSkuOutput) ToAppSkuOutput() AppSkuOutput {
	return o
}

func (o AppSkuOutput) ToAppSkuOutputWithContext(ctx context.Context) AppSkuOutput {
	return o
}

func (o AppSkuOutput) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return o.ToAppSkuPtrOutputWithContext(context.Background())
}

func (o AppSkuOutput) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppSku) *AppSku {
		return &v
	}).(AppSkuPtrOutput)
}

func (o AppSkuOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AppSkuOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppSku) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AppSkuOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppSkuOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AppSku) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AppSkuPtrOutput struct{ *pulumi.OutputState }

func (AppSkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppSku)(nil)).Elem()
}

func (o AppSkuPtrOutput) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return o
}

func (o AppSkuPtrOutput) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return o
}

func (o AppSkuPtrOutput) Elem() AppSkuOutput {
	return o.ApplyT(func(v *AppSku) AppSku {
		if v != nil {
			return *v
		}
		var ret AppSku
		return ret
	}).(AppSkuOutput)
}

func (o AppSkuPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AppSkuPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AppSku) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AppSkuInput interface {
	pulumi.Input

	ToAppSkuOutput() AppSkuOutput
	ToAppSkuOutputWithContext(context.Context) AppSkuOutput
}

var appSkuPtrType = reflect.TypeOf((**AppSku)(nil)).Elem()

type AppSkuPtrInput interface {
	pulumi.Input

	ToAppSkuPtrOutput() AppSkuPtrOutput
	ToAppSkuPtrOutputWithContext(context.Context) AppSkuPtrOutput
}

type appSkuPtr string

func AppSkuPtr(v string) AppSkuPtrInput {
	return (*appSkuPtr)(&v)
}

func (*appSkuPtr) ElementType() reflect.Type {
	return appSkuPtrType
}

func (in *appSkuPtr) ToAppSkuPtrOutput() AppSkuPtrOutput {
	return pulumi.ToOutput(in).(AppSkuPtrOutput)
}

func (in *appSkuPtr) ToAppSkuPtrOutputWithContext(ctx context.Context) AppSkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AppSkuPtrOutput)
}

type SystemAssignedServiceIdentityType string

const (
	SystemAssignedServiceIdentityTypeNone           = SystemAssignedServiceIdentityType("None")
	SystemAssignedServiceIdentityTypeSystemAssigned = SystemAssignedServiceIdentityType("SystemAssigned")
)

func (SystemAssignedServiceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput {
	return pulumi.ToOutput(e).(SystemAssignedServiceIdentityTypeOutput)
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SystemAssignedServiceIdentityTypeOutput)
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return e.ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (e SystemAssignedServiceIdentityType) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return SystemAssignedServiceIdentityType(e).ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx).ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx)
}

func (e SystemAssignedServiceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemAssignedServiceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SystemAssignedServiceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SystemAssignedServiceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SystemAssignedServiceIdentityTypeOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypeOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypeOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return o.ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemAssignedServiceIdentityType) *SystemAssignedServiceIdentityType {
		return &v
	}).(SystemAssignedServiceIdentityTypePtrOutput)
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemAssignedServiceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SystemAssignedServiceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SystemAssignedServiceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (SystemAssignedServiceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemAssignedServiceIdentityType)(nil)).Elem()
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return o
}

func (o SystemAssignedServiceIdentityTypePtrOutput) Elem() SystemAssignedServiceIdentityTypeOutput {
	return o.ApplyT(func(v *SystemAssignedServiceIdentityType) SystemAssignedServiceIdentityType {
		if v != nil {
			return *v
		}
		var ret SystemAssignedServiceIdentityType
		return ret
	}).(SystemAssignedServiceIdentityTypeOutput)
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SystemAssignedServiceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SystemAssignedServiceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SystemAssignedServiceIdentityTypeInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityTypeOutput() SystemAssignedServiceIdentityTypeOutput
	ToSystemAssignedServiceIdentityTypeOutputWithContext(context.Context) SystemAssignedServiceIdentityTypeOutput
}

var systemAssignedServiceIdentityTypePtrType = reflect.TypeOf((**SystemAssignedServiceIdentityType)(nil)).Elem()

type SystemAssignedServiceIdentityTypePtrInput interface {
	pulumi.Input

	ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput
	ToSystemAssignedServiceIdentityTypePtrOutputWithContext(context.Context) SystemAssignedServiceIdentityTypePtrOutput
}

type systemAssignedServiceIdentityTypePtr string

func SystemAssignedServiceIdentityTypePtr(v string) SystemAssignedServiceIdentityTypePtrInput {
	return (*systemAssignedServiceIdentityTypePtr)(&v)
}

func (*systemAssignedServiceIdentityTypePtr) ElementType() reflect.Type {
	return systemAssignedServiceIdentityTypePtrType
}

func (in *systemAssignedServiceIdentityTypePtr) ToSystemAssignedServiceIdentityTypePtrOutput() SystemAssignedServiceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(SystemAssignedServiceIdentityTypePtrOutput)
}

func (in *systemAssignedServiceIdentityTypePtr) ToSystemAssignedServiceIdentityTypePtrOutputWithContext(ctx context.Context) SystemAssignedServiceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SystemAssignedServiceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AppSkuOutput{})
	pulumi.RegisterOutputType(AppSkuPtrOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityTypeOutput{})
	pulumi.RegisterOutputType(SystemAssignedServiceIdentityTypePtrOutput{})
}
