


package iotcentral

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppSku string

const (
	AppSkuF1  = AppSku("F1")
	AppSkuS1  = AppSku("S1")
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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AppSkuInput)(nil)).Elem(), AppSku("F1"))
	pulumi.RegisterInputType(reflect.TypeOf((*AppSkuPtrInput)(nil)).Elem(), AppSku("F1"))
	pulumi.RegisterOutputType(AppSkuOutput{})
	pulumi.RegisterOutputType(AppSkuPtrOutput{})
}
