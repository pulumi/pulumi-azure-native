


package v20211001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BudgetOperatorType string

const (
	BudgetOperatorTypeIn = BudgetOperatorType("In")
)

func (BudgetOperatorType) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetOperatorType)(nil)).Elem()
}

func (e BudgetOperatorType) ToBudgetOperatorTypeOutput() BudgetOperatorTypeOutput {
	return pulumi.ToOutput(e).(BudgetOperatorTypeOutput)
}

func (e BudgetOperatorType) ToBudgetOperatorTypeOutputWithContext(ctx context.Context) BudgetOperatorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(BudgetOperatorTypeOutput)
}

func (e BudgetOperatorType) ToBudgetOperatorTypePtrOutput() BudgetOperatorTypePtrOutput {
	return e.ToBudgetOperatorTypePtrOutputWithContext(context.Background())
}

func (e BudgetOperatorType) ToBudgetOperatorTypePtrOutputWithContext(ctx context.Context) BudgetOperatorTypePtrOutput {
	return BudgetOperatorType(e).ToBudgetOperatorTypeOutputWithContext(ctx).ToBudgetOperatorTypePtrOutputWithContext(ctx)
}

func (e BudgetOperatorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e BudgetOperatorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e BudgetOperatorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e BudgetOperatorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type BudgetOperatorTypeOutput struct{ *pulumi.OutputState }

func (BudgetOperatorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BudgetOperatorType)(nil)).Elem()
}

func (o BudgetOperatorTypeOutput) ToBudgetOperatorTypeOutput() BudgetOperatorTypeOutput {
	return o
}

func (o BudgetOperatorTypeOutput) ToBudgetOperatorTypeOutputWithContext(ctx context.Context) BudgetOperatorTypeOutput {
	return o
}

func (o BudgetOperatorTypeOutput) ToBudgetOperatorTypePtrOutput() BudgetOperatorTypePtrOutput {
	return o.ToBudgetOperatorTypePtrOutputWithContext(context.Background())
}

func (o BudgetOperatorTypeOutput) ToBudgetOperatorTypePtrOutputWithContext(ctx context.Context) BudgetOperatorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BudgetOperatorType) *BudgetOperatorType {
		return &v
	}).(BudgetOperatorTypePtrOutput)
}

func (o BudgetOperatorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o BudgetOperatorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BudgetOperatorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o BudgetOperatorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BudgetOperatorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e BudgetOperatorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type BudgetOperatorTypePtrOutput struct{ *pulumi.OutputState }

func (BudgetOperatorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BudgetOperatorType)(nil)).Elem()
}

func (o BudgetOperatorTypePtrOutput) ToBudgetOperatorTypePtrOutput() BudgetOperatorTypePtrOutput {
	return o
}

func (o BudgetOperatorTypePtrOutput) ToBudgetOperatorTypePtrOutputWithContext(ctx context.Context) BudgetOperatorTypePtrOutput {
	return o
}

func (o BudgetOperatorTypePtrOutput) Elem() BudgetOperatorTypeOutput {
	return o.ApplyT(func(v *BudgetOperatorType) BudgetOperatorType {
		if v != nil {
			return *v
		}
		var ret BudgetOperatorType
		return ret
	}).(BudgetOperatorTypeOutput)
}

func (o BudgetOperatorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o BudgetOperatorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *BudgetOperatorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type BudgetOperatorTypeInput interface {
	pulumi.Input

	ToBudgetOperatorTypeOutput() BudgetOperatorTypeOutput
	ToBudgetOperatorTypeOutputWithContext(context.Context) BudgetOperatorTypeOutput
}

var budgetOperatorTypePtrType = reflect.TypeOf((**BudgetOperatorType)(nil)).Elem()

type BudgetOperatorTypePtrInput interface {
	pulumi.Input

	ToBudgetOperatorTypePtrOutput() BudgetOperatorTypePtrOutput
	ToBudgetOperatorTypePtrOutputWithContext(context.Context) BudgetOperatorTypePtrOutput
}

type budgetOperatorTypePtr string

func BudgetOperatorTypePtr(v string) BudgetOperatorTypePtrInput {
	return (*budgetOperatorTypePtr)(&v)
}

func (*budgetOperatorTypePtr) ElementType() reflect.Type {
	return budgetOperatorTypePtrType
}

func (in *budgetOperatorTypePtr) ToBudgetOperatorTypePtrOutput() BudgetOperatorTypePtrOutput {
	return pulumi.ToOutput(in).(BudgetOperatorTypePtrOutput)
}

func (in *budgetOperatorTypePtr) ToBudgetOperatorTypePtrOutputWithContext(ctx context.Context) BudgetOperatorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(BudgetOperatorTypePtrOutput)
}

type CategoryType string

const (
	CategoryTypeCost = CategoryType("Cost")
)

func (CategoryType) ElementType() reflect.Type {
	return reflect.TypeOf((*CategoryType)(nil)).Elem()
}

func (e CategoryType) ToCategoryTypeOutput() CategoryTypeOutput {
	return pulumi.ToOutput(e).(CategoryTypeOutput)
}

func (e CategoryType) ToCategoryTypeOutputWithContext(ctx context.Context) CategoryTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CategoryTypeOutput)
}

func (e CategoryType) ToCategoryTypePtrOutput() CategoryTypePtrOutput {
	return e.ToCategoryTypePtrOutputWithContext(context.Background())
}

func (e CategoryType) ToCategoryTypePtrOutputWithContext(ctx context.Context) CategoryTypePtrOutput {
	return CategoryType(e).ToCategoryTypeOutputWithContext(ctx).ToCategoryTypePtrOutputWithContext(ctx)
}

func (e CategoryType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CategoryType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CategoryType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CategoryType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CategoryTypeOutput struct{ *pulumi.OutputState }

func (CategoryTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CategoryType)(nil)).Elem()
}

func (o CategoryTypeOutput) ToCategoryTypeOutput() CategoryTypeOutput {
	return o
}

func (o CategoryTypeOutput) ToCategoryTypeOutputWithContext(ctx context.Context) CategoryTypeOutput {
	return o
}

func (o CategoryTypeOutput) ToCategoryTypePtrOutput() CategoryTypePtrOutput {
	return o.ToCategoryTypePtrOutputWithContext(context.Background())
}

func (o CategoryTypeOutput) ToCategoryTypePtrOutputWithContext(ctx context.Context) CategoryTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CategoryType) *CategoryType {
		return &v
	}).(CategoryTypePtrOutput)
}

func (o CategoryTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CategoryTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CategoryType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CategoryTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CategoryTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CategoryType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CategoryTypePtrOutput struct{ *pulumi.OutputState }

func (CategoryTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CategoryType)(nil)).Elem()
}

func (o CategoryTypePtrOutput) ToCategoryTypePtrOutput() CategoryTypePtrOutput {
	return o
}

func (o CategoryTypePtrOutput) ToCategoryTypePtrOutputWithContext(ctx context.Context) CategoryTypePtrOutput {
	return o
}

func (o CategoryTypePtrOutput) Elem() CategoryTypeOutput {
	return o.ApplyT(func(v *CategoryType) CategoryType {
		if v != nil {
			return *v
		}
		var ret CategoryType
		return ret
	}).(CategoryTypeOutput)
}

func (o CategoryTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CategoryTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CategoryType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CategoryTypeInput interface {
	pulumi.Input

	ToCategoryTypeOutput() CategoryTypeOutput
	ToCategoryTypeOutputWithContext(context.Context) CategoryTypeOutput
}

var categoryTypePtrType = reflect.TypeOf((**CategoryType)(nil)).Elem()

type CategoryTypePtrInput interface {
	pulumi.Input

	ToCategoryTypePtrOutput() CategoryTypePtrOutput
	ToCategoryTypePtrOutputWithContext(context.Context) CategoryTypePtrOutput
}

type categoryTypePtr string

func CategoryTypePtr(v string) CategoryTypePtrInput {
	return (*categoryTypePtr)(&v)
}

func (*categoryTypePtr) ElementType() reflect.Type {
	return categoryTypePtrType
}

func (in *categoryTypePtr) ToCategoryTypePtrOutput() CategoryTypePtrOutput {
	return pulumi.ToOutput(in).(CategoryTypePtrOutput)
}

func (in *categoryTypePtr) ToCategoryTypePtrOutputWithContext(ctx context.Context) CategoryTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CategoryTypePtrOutput)
}

type CultureCode string

const (
	CultureCode_En_us = CultureCode("en-us")
	CultureCode_Ja_jp = CultureCode("ja-jp")
	CultureCode_Zh_cn = CultureCode("zh-cn")
	CultureCode_De_de = CultureCode("de-de")
	CultureCode_Es_es = CultureCode("es-es")
	CultureCode_Fr_fr = CultureCode("fr-fr")
	CultureCode_It_it = CultureCode("it-it")
	CultureCode_Ko_kr = CultureCode("ko-kr")
	CultureCode_Pt_br = CultureCode("pt-br")
	CultureCode_Ru_ru = CultureCode("ru-ru")
	CultureCode_Zh_tw = CultureCode("zh-tw")
	CultureCode_Cs_cz = CultureCode("cs-cz")
	CultureCode_Pl_pl = CultureCode("pl-pl")
	CultureCode_Tr_tr = CultureCode("tr-tr")
	CultureCode_Da_dk = CultureCode("da-dk")
	CultureCode_En_gb = CultureCode("en-gb")
	CultureCode_Hu_hu = CultureCode("hu-hu")
	CultureCode_Nb_no = CultureCode("nb-no")
	CultureCode_Nl_nl = CultureCode("nl-nl")
	CultureCode_Pt_pt = CultureCode("pt-pt")
	CultureCode_Sv_se = CultureCode("sv-se")
)

func (CultureCode) ElementType() reflect.Type {
	return reflect.TypeOf((*CultureCode)(nil)).Elem()
}

func (e CultureCode) ToCultureCodeOutput() CultureCodeOutput {
	return pulumi.ToOutput(e).(CultureCodeOutput)
}

func (e CultureCode) ToCultureCodeOutputWithContext(ctx context.Context) CultureCodeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(CultureCodeOutput)
}

func (e CultureCode) ToCultureCodePtrOutput() CultureCodePtrOutput {
	return e.ToCultureCodePtrOutputWithContext(context.Background())
}

func (e CultureCode) ToCultureCodePtrOutputWithContext(ctx context.Context) CultureCodePtrOutput {
	return CultureCode(e).ToCultureCodeOutputWithContext(ctx).ToCultureCodePtrOutputWithContext(ctx)
}

func (e CultureCode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e CultureCode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e CultureCode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e CultureCode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type CultureCodeOutput struct{ *pulumi.OutputState }

func (CultureCodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CultureCode)(nil)).Elem()
}

func (o CultureCodeOutput) ToCultureCodeOutput() CultureCodeOutput {
	return o
}

func (o CultureCodeOutput) ToCultureCodeOutputWithContext(ctx context.Context) CultureCodeOutput {
	return o
}

func (o CultureCodeOutput) ToCultureCodePtrOutput() CultureCodePtrOutput {
	return o.ToCultureCodePtrOutputWithContext(context.Background())
}

func (o CultureCodeOutput) ToCultureCodePtrOutputWithContext(ctx context.Context) CultureCodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CultureCode) *CultureCode {
		return &v
	}).(CultureCodePtrOutput)
}

func (o CultureCodeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o CultureCodeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CultureCode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o CultureCodeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CultureCodeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e CultureCode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type CultureCodePtrOutput struct{ *pulumi.OutputState }

func (CultureCodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CultureCode)(nil)).Elem()
}

func (o CultureCodePtrOutput) ToCultureCodePtrOutput() CultureCodePtrOutput {
	return o
}

func (o CultureCodePtrOutput) ToCultureCodePtrOutputWithContext(ctx context.Context) CultureCodePtrOutput {
	return o
}

func (o CultureCodePtrOutput) Elem() CultureCodeOutput {
	return o.ApplyT(func(v *CultureCode) CultureCode {
		if v != nil {
			return *v
		}
		var ret CultureCode
		return ret
	}).(CultureCodeOutput)
}

func (o CultureCodePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o CultureCodePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *CultureCode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type CultureCodeInput interface {
	pulumi.Input

	ToCultureCodeOutput() CultureCodeOutput
	ToCultureCodeOutputWithContext(context.Context) CultureCodeOutput
}

var cultureCodePtrType = reflect.TypeOf((**CultureCode)(nil)).Elem()

type CultureCodePtrInput interface {
	pulumi.Input

	ToCultureCodePtrOutput() CultureCodePtrOutput
	ToCultureCodePtrOutputWithContext(context.Context) CultureCodePtrOutput
}

type cultureCodePtr string

func CultureCodePtr(v string) CultureCodePtrInput {
	return (*cultureCodePtr)(&v)
}

func (*cultureCodePtr) ElementType() reflect.Type {
	return cultureCodePtrType
}

func (in *cultureCodePtr) ToCultureCodePtrOutput() CultureCodePtrOutput {
	return pulumi.ToOutput(in).(CultureCodePtrOutput)
}

func (in *cultureCodePtr) ToCultureCodePtrOutputWithContext(ctx context.Context) CultureCodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(CultureCodePtrOutput)
}

type OperatorType string

const (
	OperatorTypeEqualTo              = OperatorType("EqualTo")
	OperatorTypeGreaterThan          = OperatorType("GreaterThan")
	OperatorTypeGreaterThanOrEqualTo = OperatorType("GreaterThanOrEqualTo")
)

func (OperatorType) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorType)(nil)).Elem()
}

func (e OperatorType) ToOperatorTypeOutput() OperatorTypeOutput {
	return pulumi.ToOutput(e).(OperatorTypeOutput)
}

func (e OperatorType) ToOperatorTypeOutputWithContext(ctx context.Context) OperatorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(OperatorTypeOutput)
}

func (e OperatorType) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return e.ToOperatorTypePtrOutputWithContext(context.Background())
}

func (e OperatorType) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return OperatorType(e).ToOperatorTypeOutputWithContext(ctx).ToOperatorTypePtrOutputWithContext(ctx)
}

func (e OperatorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e OperatorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e OperatorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type OperatorTypeOutput struct{ *pulumi.OutputState }

func (OperatorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*OperatorType)(nil)).Elem()
}

func (o OperatorTypeOutput) ToOperatorTypeOutput() OperatorTypeOutput {
	return o
}

func (o OperatorTypeOutput) ToOperatorTypeOutputWithContext(ctx context.Context) OperatorTypeOutput {
	return o
}

func (o OperatorTypeOutput) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return o.ToOperatorTypePtrOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v OperatorType) *OperatorType {
		return &v
	}).(OperatorTypePtrOutput)
}

func (o OperatorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o OperatorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e OperatorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type OperatorTypePtrOutput struct{ *pulumi.OutputState }

func (OperatorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**OperatorType)(nil)).Elem()
}

func (o OperatorTypePtrOutput) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return o
}

func (o OperatorTypePtrOutput) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return o
}

func (o OperatorTypePtrOutput) Elem() OperatorTypeOutput {
	return o.ApplyT(func(v *OperatorType) OperatorType {
		if v != nil {
			return *v
		}
		var ret OperatorType
		return ret
	}).(OperatorTypeOutput)
}

func (o OperatorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o OperatorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *OperatorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type OperatorTypeInput interface {
	pulumi.Input

	ToOperatorTypeOutput() OperatorTypeOutput
	ToOperatorTypeOutputWithContext(context.Context) OperatorTypeOutput
}

var operatorTypePtrType = reflect.TypeOf((**OperatorType)(nil)).Elem()

type OperatorTypePtrInput interface {
	pulumi.Input

	ToOperatorTypePtrOutput() OperatorTypePtrOutput
	ToOperatorTypePtrOutputWithContext(context.Context) OperatorTypePtrOutput
}

type operatorTypePtr string

func OperatorTypePtr(v string) OperatorTypePtrInput {
	return (*operatorTypePtr)(&v)
}

func (*operatorTypePtr) ElementType() reflect.Type {
	return operatorTypePtrType
}

func (in *operatorTypePtr) ToOperatorTypePtrOutput() OperatorTypePtrOutput {
	return pulumi.ToOutput(in).(OperatorTypePtrOutput)
}

func (in *operatorTypePtr) ToOperatorTypePtrOutputWithContext(ctx context.Context) OperatorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(OperatorTypePtrOutput)
}

type ThresholdType string

const (
	ThresholdTypeActual     = ThresholdType("Actual")
	ThresholdTypeForecasted = ThresholdType("Forecasted")
)

func (ThresholdType) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdType)(nil)).Elem()
}

func (e ThresholdType) ToThresholdTypeOutput() ThresholdTypeOutput {
	return pulumi.ToOutput(e).(ThresholdTypeOutput)
}

func (e ThresholdType) ToThresholdTypeOutputWithContext(ctx context.Context) ThresholdTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ThresholdTypeOutput)
}

func (e ThresholdType) ToThresholdTypePtrOutput() ThresholdTypePtrOutput {
	return e.ToThresholdTypePtrOutputWithContext(context.Background())
}

func (e ThresholdType) ToThresholdTypePtrOutputWithContext(ctx context.Context) ThresholdTypePtrOutput {
	return ThresholdType(e).ToThresholdTypeOutputWithContext(ctx).ToThresholdTypePtrOutputWithContext(ctx)
}

func (e ThresholdType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ThresholdType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ThresholdType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ThresholdType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ThresholdTypeOutput struct{ *pulumi.OutputState }

func (ThresholdTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThresholdType)(nil)).Elem()
}

func (o ThresholdTypeOutput) ToThresholdTypeOutput() ThresholdTypeOutput {
	return o
}

func (o ThresholdTypeOutput) ToThresholdTypeOutputWithContext(ctx context.Context) ThresholdTypeOutput {
	return o
}

func (o ThresholdTypeOutput) ToThresholdTypePtrOutput() ThresholdTypePtrOutput {
	return o.ToThresholdTypePtrOutputWithContext(context.Background())
}

func (o ThresholdTypeOutput) ToThresholdTypePtrOutputWithContext(ctx context.Context) ThresholdTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ThresholdType) *ThresholdType {
		return &v
	}).(ThresholdTypePtrOutput)
}

func (o ThresholdTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ThresholdTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ThresholdType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ThresholdTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ThresholdTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ThresholdType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ThresholdTypePtrOutput struct{ *pulumi.OutputState }

func (ThresholdTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThresholdType)(nil)).Elem()
}

func (o ThresholdTypePtrOutput) ToThresholdTypePtrOutput() ThresholdTypePtrOutput {
	return o
}

func (o ThresholdTypePtrOutput) ToThresholdTypePtrOutputWithContext(ctx context.Context) ThresholdTypePtrOutput {
	return o
}

func (o ThresholdTypePtrOutput) Elem() ThresholdTypeOutput {
	return o.ApplyT(func(v *ThresholdType) ThresholdType {
		if v != nil {
			return *v
		}
		var ret ThresholdType
		return ret
	}).(ThresholdTypeOutput)
}

func (o ThresholdTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ThresholdTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ThresholdType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ThresholdTypeInput interface {
	pulumi.Input

	ToThresholdTypeOutput() ThresholdTypeOutput
	ToThresholdTypeOutputWithContext(context.Context) ThresholdTypeOutput
}

var thresholdTypePtrType = reflect.TypeOf((**ThresholdType)(nil)).Elem()

type ThresholdTypePtrInput interface {
	pulumi.Input

	ToThresholdTypePtrOutput() ThresholdTypePtrOutput
	ToThresholdTypePtrOutputWithContext(context.Context) ThresholdTypePtrOutput
}

type thresholdTypePtr string

func ThresholdTypePtr(v string) ThresholdTypePtrInput {
	return (*thresholdTypePtr)(&v)
}

func (*thresholdTypePtr) ElementType() reflect.Type {
	return thresholdTypePtrType
}

func (in *thresholdTypePtr) ToThresholdTypePtrOutput() ThresholdTypePtrOutput {
	return pulumi.ToOutput(in).(ThresholdTypePtrOutput)
}

func (in *thresholdTypePtr) ToThresholdTypePtrOutputWithContext(ctx context.Context) ThresholdTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ThresholdTypePtrOutput)
}

type TimeGrainType string

const (
	TimeGrainTypeMonthly        = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly      = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually       = TimeGrainType("Annually")
	TimeGrainTypeBillingMonth   = TimeGrainType("BillingMonth")
	TimeGrainTypeBillingQuarter = TimeGrainType("BillingQuarter")
	TimeGrainTypeBillingAnnual  = TimeGrainType("BillingAnnual")
)

func (TimeGrainType) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeGrainType)(nil)).Elem()
}

func (e TimeGrainType) ToTimeGrainTypeOutput() TimeGrainTypeOutput {
	return pulumi.ToOutput(e).(TimeGrainTypeOutput)
}

func (e TimeGrainType) ToTimeGrainTypeOutputWithContext(ctx context.Context) TimeGrainTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeGrainTypeOutput)
}

func (e TimeGrainType) ToTimeGrainTypePtrOutput() TimeGrainTypePtrOutput {
	return e.ToTimeGrainTypePtrOutputWithContext(context.Background())
}

func (e TimeGrainType) ToTimeGrainTypePtrOutputWithContext(ctx context.Context) TimeGrainTypePtrOutput {
	return TimeGrainType(e).ToTimeGrainTypeOutputWithContext(ctx).ToTimeGrainTypePtrOutputWithContext(ctx)
}

func (e TimeGrainType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeGrainType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeGrainType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeGrainType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeGrainTypeOutput struct{ *pulumi.OutputState }

func (TimeGrainTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeGrainType)(nil)).Elem()
}

func (o TimeGrainTypeOutput) ToTimeGrainTypeOutput() TimeGrainTypeOutput {
	return o
}

func (o TimeGrainTypeOutput) ToTimeGrainTypeOutputWithContext(ctx context.Context) TimeGrainTypeOutput {
	return o
}

func (o TimeGrainTypeOutput) ToTimeGrainTypePtrOutput() TimeGrainTypePtrOutput {
	return o.ToTimeGrainTypePtrOutputWithContext(context.Background())
}

func (o TimeGrainTypeOutput) ToTimeGrainTypePtrOutputWithContext(ctx context.Context) TimeGrainTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeGrainType) *TimeGrainType {
		return &v
	}).(TimeGrainTypePtrOutput)
}

func (o TimeGrainTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeGrainTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeGrainType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeGrainTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeGrainTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeGrainType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeGrainTypePtrOutput struct{ *pulumi.OutputState }

func (TimeGrainTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeGrainType)(nil)).Elem()
}

func (o TimeGrainTypePtrOutput) ToTimeGrainTypePtrOutput() TimeGrainTypePtrOutput {
	return o
}

func (o TimeGrainTypePtrOutput) ToTimeGrainTypePtrOutputWithContext(ctx context.Context) TimeGrainTypePtrOutput {
	return o
}

func (o TimeGrainTypePtrOutput) Elem() TimeGrainTypeOutput {
	return o.ApplyT(func(v *TimeGrainType) TimeGrainType {
		if v != nil {
			return *v
		}
		var ret TimeGrainType
		return ret
	}).(TimeGrainTypeOutput)
}

func (o TimeGrainTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeGrainTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeGrainType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TimeGrainTypeInput interface {
	pulumi.Input

	ToTimeGrainTypeOutput() TimeGrainTypeOutput
	ToTimeGrainTypeOutputWithContext(context.Context) TimeGrainTypeOutput
}

var timeGrainTypePtrType = reflect.TypeOf((**TimeGrainType)(nil)).Elem()

type TimeGrainTypePtrInput interface {
	pulumi.Input

	ToTimeGrainTypePtrOutput() TimeGrainTypePtrOutput
	ToTimeGrainTypePtrOutputWithContext(context.Context) TimeGrainTypePtrOutput
}

type timeGrainTypePtr string

func TimeGrainTypePtr(v string) TimeGrainTypePtrInput {
	return (*timeGrainTypePtr)(&v)
}

func (*timeGrainTypePtr) ElementType() reflect.Type {
	return timeGrainTypePtrType
}

func (in *timeGrainTypePtr) ToTimeGrainTypePtrOutput() TimeGrainTypePtrOutput {
	return pulumi.ToOutput(in).(TimeGrainTypePtrOutput)
}

func (in *timeGrainTypePtr) ToTimeGrainTypePtrOutputWithContext(ctx context.Context) TimeGrainTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeGrainTypePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(BudgetOperatorTypeOutput{})
	pulumi.RegisterOutputType(BudgetOperatorTypePtrOutput{})
	pulumi.RegisterOutputType(CategoryTypeOutput{})
	pulumi.RegisterOutputType(CategoryTypePtrOutput{})
	pulumi.RegisterOutputType(CultureCodeOutput{})
	pulumi.RegisterOutputType(CultureCodePtrOutput{})
	pulumi.RegisterOutputType(OperatorTypeOutput{})
	pulumi.RegisterOutputType(OperatorTypePtrOutput{})
	pulumi.RegisterOutputType(ThresholdTypeOutput{})
	pulumi.RegisterOutputType(ThresholdTypePtrOutput{})
	pulumi.RegisterOutputType(TimeGrainTypeOutput{})
	pulumi.RegisterOutputType(TimeGrainTypePtrOutput{})
}
