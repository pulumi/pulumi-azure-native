


package extendedlocation

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CustomLocationPropertiesAuthentication struct {
	Type  *string `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type CustomLocationPropertiesAuthenticationInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput
	ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationOutput
}

type CustomLocationPropertiesAuthenticationArgs struct {
	Type  pulumi.StringPtrInput `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (CustomLocationPropertiesAuthenticationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return i.ToCustomLocationPropertiesAuthenticationOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput)
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i CustomLocationPropertiesAuthenticationArgs) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationOutput).ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx)
}









type CustomLocationPropertiesAuthenticationPtrInput interface {
	pulumi.Input

	ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput
	ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Context) CustomLocationPropertiesAuthenticationPtrOutput
}

type customLocationPropertiesAuthenticationPtrType CustomLocationPropertiesAuthenticationArgs

func CustomLocationPropertiesAuthenticationPtr(v *CustomLocationPropertiesAuthenticationArgs) CustomLocationPropertiesAuthenticationPtrInput {
	return (*customLocationPropertiesAuthenticationPtrType)(v)
}

func (*customLocationPropertiesAuthenticationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return i.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (i *customLocationPropertiesAuthenticationPtrType) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomLocationPropertiesAuthenticationPtrOutput)
}

type CustomLocationPropertiesAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutput() CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(context.Background())
}

func (o CustomLocationPropertiesAuthenticationOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomLocationPropertiesAuthentication) *CustomLocationPropertiesAuthentication {
		return &v
	}).(CustomLocationPropertiesAuthenticationPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesAuthentication) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutput() CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) ToCustomLocationPropertiesAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Elem() CustomLocationPropertiesAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) CustomLocationPropertiesAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesAuthentication
		return ret
	}).(CustomLocationPropertiesAuthenticationOutput)
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o CustomLocationPropertiesAuthenticationPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthentication struct {
	Type *string `pulumi:"type"`
}

type CustomLocationPropertiesResponseAuthenticationOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutput() CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) ToCustomLocationPropertiesResponseAuthenticationOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomLocationPropertiesResponseAuthentication) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomLocationPropertiesResponseAuthenticationPtrOutput struct{ *pulumi.OutputState }

func (CustomLocationPropertiesResponseAuthenticationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomLocationPropertiesResponseAuthentication)(nil)).Elem()
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutput() CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) ToCustomLocationPropertiesResponseAuthenticationPtrOutputWithContext(ctx context.Context) CustomLocationPropertiesResponseAuthenticationPtrOutput {
	return o
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Elem() CustomLocationPropertiesResponseAuthenticationOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) CustomLocationPropertiesResponseAuthentication {
		if v != nil {
			return *v
		}
		var ret CustomLocationPropertiesResponseAuthentication
		return ret
	}).(CustomLocationPropertiesResponseAuthenticationOutput)
}

func (o CustomLocationPropertiesResponseAuthenticationPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomLocationPropertiesResponseAuthentication) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationOutput{})
	pulumi.RegisterOutputType(CustomLocationPropertiesResponseAuthenticationPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
