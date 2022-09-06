


package v20211201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Amount struct {
	Currency *string  `pulumi:"currency"`
	Value    *float64 `pulumi:"value"`
}





type AmountInput interface {
	pulumi.Input

	ToAmountOutput() AmountOutput
	ToAmountOutputWithContext(context.Context) AmountOutput
}

type AmountArgs struct {
	Currency pulumi.StringPtrInput  `pulumi:"currency"`
	Value    pulumi.Float64PtrInput `pulumi:"value"`
}

func (AmountArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Amount)(nil)).Elem()
}

func (i AmountArgs) ToAmountOutput() AmountOutput {
	return i.ToAmountOutputWithContext(context.Background())
}

func (i AmountArgs) ToAmountOutputWithContext(ctx context.Context) AmountOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AmountOutput)
}

type AmountOutput struct{ *pulumi.OutputState }

func (AmountOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Amount)(nil)).Elem()
}

func (o AmountOutput) ToAmountOutput() AmountOutput {
	return o
}

func (o AmountOutput) ToAmountOutputWithContext(ctx context.Context) AmountOutput {
	return o
}

func (o AmountOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Amount) *string { return v.Currency }).(pulumi.StringPtrOutput)
}

func (o AmountOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v Amount) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
}

type AmountResponse struct {
	Currency *string  `pulumi:"currency"`
	Value    *float64 `pulumi:"value"`
}

type AmountResponseOutput struct{ *pulumi.OutputState }

func (AmountResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AmountResponse)(nil)).Elem()
}

func (o AmountResponseOutput) ToAmountResponseOutput() AmountResponseOutput {
	return o
}

func (o AmountResponseOutput) ToAmountResponseOutputWithContext(ctx context.Context) AmountResponseOutput {
	return o
}

func (o AmountResponseOutput) Currency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AmountResponse) *string { return v.Currency }).(pulumi.StringPtrOutput)
}

func (o AmountResponseOutput) Value() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v AmountResponse) *float64 { return v.Value }).(pulumi.Float64PtrOutput)
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
	pulumi.RegisterOutputType(AmountOutput{})
	pulumi.RegisterOutputType(AmountResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
