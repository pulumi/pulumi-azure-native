


package v20190401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccumulatedType string

const (
	AccumulatedTypeTrue  = AccumulatedType("true")
	AccumulatedTypeFalse = AccumulatedType("false")
)

func (AccumulatedType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccumulatedType)(nil)).Elem()
}

func (e AccumulatedType) ToAccumulatedTypeOutput() AccumulatedTypeOutput {
	return pulumi.ToOutput(e).(AccumulatedTypeOutput)
}

func (e AccumulatedType) ToAccumulatedTypeOutputWithContext(ctx context.Context) AccumulatedTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccumulatedTypeOutput)
}

func (e AccumulatedType) ToAccumulatedTypePtrOutput() AccumulatedTypePtrOutput {
	return e.ToAccumulatedTypePtrOutputWithContext(context.Background())
}

func (e AccumulatedType) ToAccumulatedTypePtrOutputWithContext(ctx context.Context) AccumulatedTypePtrOutput {
	return AccumulatedType(e).ToAccumulatedTypeOutputWithContext(ctx).ToAccumulatedTypePtrOutputWithContext(ctx)
}

func (e AccumulatedType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccumulatedType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccumulatedType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccumulatedType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccumulatedTypeOutput struct{ *pulumi.OutputState }

func (AccumulatedTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccumulatedType)(nil)).Elem()
}

func (o AccumulatedTypeOutput) ToAccumulatedTypeOutput() AccumulatedTypeOutput {
	return o
}

func (o AccumulatedTypeOutput) ToAccumulatedTypeOutputWithContext(ctx context.Context) AccumulatedTypeOutput {
	return o
}

func (o AccumulatedTypeOutput) ToAccumulatedTypePtrOutput() AccumulatedTypePtrOutput {
	return o.ToAccumulatedTypePtrOutputWithContext(context.Background())
}

func (o AccumulatedTypeOutput) ToAccumulatedTypePtrOutputWithContext(ctx context.Context) AccumulatedTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccumulatedType) *AccumulatedType {
		return &v
	}).(AccumulatedTypePtrOutput)
}

func (o AccumulatedTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccumulatedTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccumulatedType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccumulatedTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccumulatedTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccumulatedType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccumulatedTypePtrOutput struct{ *pulumi.OutputState }

func (AccumulatedTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccumulatedType)(nil)).Elem()
}

func (o AccumulatedTypePtrOutput) ToAccumulatedTypePtrOutput() AccumulatedTypePtrOutput {
	return o
}

func (o AccumulatedTypePtrOutput) ToAccumulatedTypePtrOutputWithContext(ctx context.Context) AccumulatedTypePtrOutput {
	return o
}

func (o AccumulatedTypePtrOutput) Elem() AccumulatedTypeOutput {
	return o.ApplyT(func(v *AccumulatedType) AccumulatedType {
		if v != nil {
			return *v
		}
		var ret AccumulatedType
		return ret
	}).(AccumulatedTypeOutput)
}

func (o AccumulatedTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccumulatedTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccumulatedType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccumulatedTypeInput interface {
	pulumi.Input

	ToAccumulatedTypeOutput() AccumulatedTypeOutput
	ToAccumulatedTypeOutputWithContext(context.Context) AccumulatedTypeOutput
}

var accumulatedTypePtrType = reflect.TypeOf((**AccumulatedType)(nil)).Elem()

type AccumulatedTypePtrInput interface {
	pulumi.Input

	ToAccumulatedTypePtrOutput() AccumulatedTypePtrOutput
	ToAccumulatedTypePtrOutputWithContext(context.Context) AccumulatedTypePtrOutput
}

type accumulatedTypePtr string

func AccumulatedTypePtr(v string) AccumulatedTypePtrInput {
	return (*accumulatedTypePtr)(&v)
}

func (*accumulatedTypePtr) ElementType() reflect.Type {
	return accumulatedTypePtrType
}

func (in *accumulatedTypePtr) ToAccumulatedTypePtrOutput() AccumulatedTypePtrOutput {
	return pulumi.ToOutput(in).(AccumulatedTypePtrOutput)
}

func (in *accumulatedTypePtr) ToAccumulatedTypePtrOutputWithContext(ctx context.Context) AccumulatedTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccumulatedTypePtrOutput)
}

type CategoryType string

const (
	CategoryTypeCost  = CategoryType("Cost")
	CategoryTypeUsage = CategoryType("Usage")
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

type ChartType string

const (
	ChartTypeArea          = ChartType("Area")
	ChartTypeLine          = ChartType("Line")
	ChartTypeStackedColumn = ChartType("StackedColumn")
	ChartTypeGroupedColumn = ChartType("GroupedColumn")
	ChartTypeTable         = ChartType("Table")
)

func (ChartType) ElementType() reflect.Type {
	return reflect.TypeOf((*ChartType)(nil)).Elem()
}

func (e ChartType) ToChartTypeOutput() ChartTypeOutput {
	return pulumi.ToOutput(e).(ChartTypeOutput)
}

func (e ChartType) ToChartTypeOutputWithContext(ctx context.Context) ChartTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ChartTypeOutput)
}

func (e ChartType) ToChartTypePtrOutput() ChartTypePtrOutput {
	return e.ToChartTypePtrOutputWithContext(context.Background())
}

func (e ChartType) ToChartTypePtrOutputWithContext(ctx context.Context) ChartTypePtrOutput {
	return ChartType(e).ToChartTypeOutputWithContext(ctx).ToChartTypePtrOutputWithContext(ctx)
}

func (e ChartType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChartType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ChartType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ChartType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ChartTypeOutput struct{ *pulumi.OutputState }

func (ChartTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChartType)(nil)).Elem()
}

func (o ChartTypeOutput) ToChartTypeOutput() ChartTypeOutput {
	return o
}

func (o ChartTypeOutput) ToChartTypeOutputWithContext(ctx context.Context) ChartTypeOutput {
	return o
}

func (o ChartTypeOutput) ToChartTypePtrOutput() ChartTypePtrOutput {
	return o.ToChartTypePtrOutputWithContext(context.Background())
}

func (o ChartTypeOutput) ToChartTypePtrOutputWithContext(ctx context.Context) ChartTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ChartType) *ChartType {
		return &v
	}).(ChartTypePtrOutput)
}

func (o ChartTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ChartTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChartType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ChartTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChartTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ChartType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ChartTypePtrOutput struct{ *pulumi.OutputState }

func (ChartTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChartType)(nil)).Elem()
}

func (o ChartTypePtrOutput) ToChartTypePtrOutput() ChartTypePtrOutput {
	return o
}

func (o ChartTypePtrOutput) ToChartTypePtrOutputWithContext(ctx context.Context) ChartTypePtrOutput {
	return o
}

func (o ChartTypePtrOutput) Elem() ChartTypeOutput {
	return o.ApplyT(func(v *ChartType) ChartType {
		if v != nil {
			return *v
		}
		var ret ChartType
		return ret
	}).(ChartTypeOutput)
}

func (o ChartTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ChartTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ChartType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ChartTypeInput interface {
	pulumi.Input

	ToChartTypeOutput() ChartTypeOutput
	ToChartTypeOutputWithContext(context.Context) ChartTypeOutput
}

var chartTypePtrType = reflect.TypeOf((**ChartType)(nil)).Elem()

type ChartTypePtrInput interface {
	pulumi.Input

	ToChartTypePtrOutput() ChartTypePtrOutput
	ToChartTypePtrOutputWithContext(context.Context) ChartTypePtrOutput
}

type chartTypePtr string

func ChartTypePtr(v string) ChartTypePtrInput {
	return (*chartTypePtr)(&v)
}

func (*chartTypePtr) ElementType() reflect.Type {
	return chartTypePtrType
}

func (in *chartTypePtr) ToChartTypePtrOutput() ChartTypePtrOutput {
	return pulumi.ToOutput(in).(ChartTypePtrOutput)
}

func (in *chartTypePtr) ToChartTypePtrOutputWithContext(ctx context.Context) ChartTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ChartTypePtrOutput)
}

type FunctionType string

const (
	FunctionTypeSum = FunctionType("Sum")
)

func (FunctionType) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (e FunctionType) ToFunctionTypeOutput() FunctionTypeOutput {
	return pulumi.ToOutput(e).(FunctionTypeOutput)
}

func (e FunctionType) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FunctionTypeOutput)
}

func (e FunctionType) ToFunctionTypePtrOutput() FunctionTypePtrOutput {
	return e.ToFunctionTypePtrOutputWithContext(context.Background())
}

func (e FunctionType) ToFunctionTypePtrOutputWithContext(ctx context.Context) FunctionTypePtrOutput {
	return FunctionType(e).ToFunctionTypeOutputWithContext(ctx).ToFunctionTypePtrOutputWithContext(ctx)
}

func (e FunctionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FunctionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FunctionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FunctionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FunctionTypeOutput struct{ *pulumi.OutputState }

func (FunctionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FunctionType)(nil)).Elem()
}

func (o FunctionTypeOutput) ToFunctionTypeOutput() FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) ToFunctionTypeOutputWithContext(ctx context.Context) FunctionTypeOutput {
	return o
}

func (o FunctionTypeOutput) ToFunctionTypePtrOutput() FunctionTypePtrOutput {
	return o.ToFunctionTypePtrOutputWithContext(context.Background())
}

func (o FunctionTypeOutput) ToFunctionTypePtrOutputWithContext(ctx context.Context) FunctionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FunctionType) *FunctionType {
		return &v
	}).(FunctionTypePtrOutput)
}

func (o FunctionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FunctionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FunctionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FunctionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FunctionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FunctionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FunctionTypePtrOutput struct{ *pulumi.OutputState }

func (FunctionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FunctionType)(nil)).Elem()
}

func (o FunctionTypePtrOutput) ToFunctionTypePtrOutput() FunctionTypePtrOutput {
	return o
}

func (o FunctionTypePtrOutput) ToFunctionTypePtrOutputWithContext(ctx context.Context) FunctionTypePtrOutput {
	return o
}

func (o FunctionTypePtrOutput) Elem() FunctionTypeOutput {
	return o.ApplyT(func(v *FunctionType) FunctionType {
		if v != nil {
			return *v
		}
		var ret FunctionType
		return ret
	}).(FunctionTypeOutput)
}

func (o FunctionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FunctionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FunctionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FunctionTypeInput interface {
	pulumi.Input

	ToFunctionTypeOutput() FunctionTypeOutput
	ToFunctionTypeOutputWithContext(context.Context) FunctionTypeOutput
}

var functionTypePtrType = reflect.TypeOf((**FunctionType)(nil)).Elem()

type FunctionTypePtrInput interface {
	pulumi.Input

	ToFunctionTypePtrOutput() FunctionTypePtrOutput
	ToFunctionTypePtrOutputWithContext(context.Context) FunctionTypePtrOutput
}

type functionTypePtr string

func FunctionTypePtr(v string) FunctionTypePtrInput {
	return (*functionTypePtr)(&v)
}

func (*functionTypePtr) ElementType() reflect.Type {
	return functionTypePtrType
}

func (in *functionTypePtr) ToFunctionTypePtrOutput() FunctionTypePtrOutput {
	return pulumi.ToOutput(in).(FunctionTypePtrOutput)
}

func (in *functionTypePtr) ToFunctionTypePtrOutputWithContext(ctx context.Context) FunctionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FunctionTypePtrOutput)
}

type GranularityType string

const (
	GranularityTypeDaily   = GranularityType("Daily")
	GranularityTypeMonthly = GranularityType("Monthly")
)

func (GranularityType) ElementType() reflect.Type {
	return reflect.TypeOf((*GranularityType)(nil)).Elem()
}

func (e GranularityType) ToGranularityTypeOutput() GranularityTypeOutput {
	return pulumi.ToOutput(e).(GranularityTypeOutput)
}

func (e GranularityType) ToGranularityTypeOutputWithContext(ctx context.Context) GranularityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(GranularityTypeOutput)
}

func (e GranularityType) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return e.ToGranularityTypePtrOutputWithContext(context.Background())
}

func (e GranularityType) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return GranularityType(e).ToGranularityTypeOutputWithContext(ctx).ToGranularityTypePtrOutputWithContext(ctx)
}

func (e GranularityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e GranularityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e GranularityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e GranularityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type GranularityTypeOutput struct{ *pulumi.OutputState }

func (GranularityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GranularityType)(nil)).Elem()
}

func (o GranularityTypeOutput) ToGranularityTypeOutput() GranularityTypeOutput {
	return o
}

func (o GranularityTypeOutput) ToGranularityTypeOutputWithContext(ctx context.Context) GranularityTypeOutput {
	return o
}

func (o GranularityTypeOutput) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return o.ToGranularityTypePtrOutputWithContext(context.Background())
}

func (o GranularityTypeOutput) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GranularityType) *GranularityType {
		return &v
	}).(GranularityTypePtrOutput)
}

func (o GranularityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o GranularityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GranularityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o GranularityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GranularityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e GranularityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type GranularityTypePtrOutput struct{ *pulumi.OutputState }

func (GranularityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GranularityType)(nil)).Elem()
}

func (o GranularityTypePtrOutput) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return o
}

func (o GranularityTypePtrOutput) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return o
}

func (o GranularityTypePtrOutput) Elem() GranularityTypeOutput {
	return o.ApplyT(func(v *GranularityType) GranularityType {
		if v != nil {
			return *v
		}
		var ret GranularityType
		return ret
	}).(GranularityTypeOutput)
}

func (o GranularityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o GranularityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *GranularityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type GranularityTypeInput interface {
	pulumi.Input

	ToGranularityTypeOutput() GranularityTypeOutput
	ToGranularityTypeOutputWithContext(context.Context) GranularityTypeOutput
}

var granularityTypePtrType = reflect.TypeOf((**GranularityType)(nil)).Elem()

type GranularityTypePtrInput interface {
	pulumi.Input

	ToGranularityTypePtrOutput() GranularityTypePtrOutput
	ToGranularityTypePtrOutputWithContext(context.Context) GranularityTypePtrOutput
}

type granularityTypePtr string

func GranularityTypePtr(v string) GranularityTypePtrInput {
	return (*granularityTypePtr)(&v)
}

func (*granularityTypePtr) ElementType() reflect.Type {
	return granularityTypePtrType
}

func (in *granularityTypePtr) ToGranularityTypePtrOutput() GranularityTypePtrOutput {
	return pulumi.ToOutput(in).(GranularityTypePtrOutput)
}

func (in *granularityTypePtr) ToGranularityTypePtrOutputWithContext(ctx context.Context) GranularityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(GranularityTypePtrOutput)
}

type KpiTypeType string

const (
	KpiTypeTypeForecast = KpiTypeType("Forecast")
	KpiTypeTypeBudget   = KpiTypeType("Budget")
)

func (KpiTypeType) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiTypeType)(nil)).Elem()
}

func (e KpiTypeType) ToKpiTypeTypeOutput() KpiTypeTypeOutput {
	return pulumi.ToOutput(e).(KpiTypeTypeOutput)
}

func (e KpiTypeType) ToKpiTypeTypeOutputWithContext(ctx context.Context) KpiTypeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(KpiTypeTypeOutput)
}

func (e KpiTypeType) ToKpiTypeTypePtrOutput() KpiTypeTypePtrOutput {
	return e.ToKpiTypeTypePtrOutputWithContext(context.Background())
}

func (e KpiTypeType) ToKpiTypeTypePtrOutputWithContext(ctx context.Context) KpiTypeTypePtrOutput {
	return KpiTypeType(e).ToKpiTypeTypeOutputWithContext(ctx).ToKpiTypeTypePtrOutputWithContext(ctx)
}

func (e KpiTypeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e KpiTypeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e KpiTypeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e KpiTypeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type KpiTypeTypeOutput struct{ *pulumi.OutputState }

func (KpiTypeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KpiTypeType)(nil)).Elem()
}

func (o KpiTypeTypeOutput) ToKpiTypeTypeOutput() KpiTypeTypeOutput {
	return o
}

func (o KpiTypeTypeOutput) ToKpiTypeTypeOutputWithContext(ctx context.Context) KpiTypeTypeOutput {
	return o
}

func (o KpiTypeTypeOutput) ToKpiTypeTypePtrOutput() KpiTypeTypePtrOutput {
	return o.ToKpiTypeTypePtrOutputWithContext(context.Background())
}

func (o KpiTypeTypeOutput) ToKpiTypeTypePtrOutputWithContext(ctx context.Context) KpiTypeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KpiTypeType) *KpiTypeType {
		return &v
	}).(KpiTypeTypePtrOutput)
}

func (o KpiTypeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o KpiTypeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KpiTypeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o KpiTypeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KpiTypeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e KpiTypeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type KpiTypeTypePtrOutput struct{ *pulumi.OutputState }

func (KpiTypeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KpiTypeType)(nil)).Elem()
}

func (o KpiTypeTypePtrOutput) ToKpiTypeTypePtrOutput() KpiTypeTypePtrOutput {
	return o
}

func (o KpiTypeTypePtrOutput) ToKpiTypeTypePtrOutputWithContext(ctx context.Context) KpiTypeTypePtrOutput {
	return o
}

func (o KpiTypeTypePtrOutput) Elem() KpiTypeTypeOutput {
	return o.ApplyT(func(v *KpiTypeType) KpiTypeType {
		if v != nil {
			return *v
		}
		var ret KpiTypeType
		return ret
	}).(KpiTypeTypeOutput)
}

func (o KpiTypeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o KpiTypeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *KpiTypeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type KpiTypeTypeInput interface {
	pulumi.Input

	ToKpiTypeTypeOutput() KpiTypeTypeOutput
	ToKpiTypeTypeOutputWithContext(context.Context) KpiTypeTypeOutput
}

var kpiTypeTypePtrType = reflect.TypeOf((**KpiTypeType)(nil)).Elem()

type KpiTypeTypePtrInput interface {
	pulumi.Input

	ToKpiTypeTypePtrOutput() KpiTypeTypePtrOutput
	ToKpiTypeTypePtrOutputWithContext(context.Context) KpiTypeTypePtrOutput
}

type kpiTypeTypePtr string

func KpiTypeTypePtr(v string) KpiTypeTypePtrInput {
	return (*kpiTypeTypePtr)(&v)
}

func (*kpiTypeTypePtr) ElementType() reflect.Type {
	return kpiTypeTypePtrType
}

func (in *kpiTypeTypePtr) ToKpiTypeTypePtrOutput() KpiTypeTypePtrOutput {
	return pulumi.ToOutput(in).(KpiTypeTypePtrOutput)
}

func (in *kpiTypeTypePtr) ToKpiTypeTypePtrOutputWithContext(ctx context.Context) KpiTypeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(KpiTypeTypePtrOutput)
}

type MetricType string

const (
	MetricTypeActualCost    = MetricType("ActualCost")
	MetricTypeAmortizedCost = MetricType("AmortizedCost")
	MetricTypeAHUB          = MetricType("AHUB")
)

func (MetricType) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricType)(nil)).Elem()
}

func (e MetricType) ToMetricTypeOutput() MetricTypeOutput {
	return pulumi.ToOutput(e).(MetricTypeOutput)
}

func (e MetricType) ToMetricTypeOutputWithContext(ctx context.Context) MetricTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(MetricTypeOutput)
}

func (e MetricType) ToMetricTypePtrOutput() MetricTypePtrOutput {
	return e.ToMetricTypePtrOutputWithContext(context.Background())
}

func (e MetricType) ToMetricTypePtrOutputWithContext(ctx context.Context) MetricTypePtrOutput {
	return MetricType(e).ToMetricTypeOutputWithContext(ctx).ToMetricTypePtrOutputWithContext(ctx)
}

func (e MetricType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e MetricType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e MetricType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e MetricType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type MetricTypeOutput struct{ *pulumi.OutputState }

func (MetricTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricType)(nil)).Elem()
}

func (o MetricTypeOutput) ToMetricTypeOutput() MetricTypeOutput {
	return o
}

func (o MetricTypeOutput) ToMetricTypeOutputWithContext(ctx context.Context) MetricTypeOutput {
	return o
}

func (o MetricTypeOutput) ToMetricTypePtrOutput() MetricTypePtrOutput {
	return o.ToMetricTypePtrOutputWithContext(context.Background())
}

func (o MetricTypeOutput) ToMetricTypePtrOutputWithContext(ctx context.Context) MetricTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MetricType) *MetricType {
		return &v
	}).(MetricTypePtrOutput)
}

func (o MetricTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o MetricTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MetricType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o MetricTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetricTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e MetricType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type MetricTypePtrOutput struct{ *pulumi.OutputState }

func (MetricTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricType)(nil)).Elem()
}

func (o MetricTypePtrOutput) ToMetricTypePtrOutput() MetricTypePtrOutput {
	return o
}

func (o MetricTypePtrOutput) ToMetricTypePtrOutputWithContext(ctx context.Context) MetricTypePtrOutput {
	return o
}

func (o MetricTypePtrOutput) Elem() MetricTypeOutput {
	return o.ApplyT(func(v *MetricType) MetricType {
		if v != nil {
			return *v
		}
		var ret MetricType
		return ret
	}).(MetricTypeOutput)
}

func (o MetricTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o MetricTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *MetricType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type MetricTypeInput interface {
	pulumi.Input

	ToMetricTypeOutput() MetricTypeOutput
	ToMetricTypeOutputWithContext(context.Context) MetricTypeOutput
}

var metricTypePtrType = reflect.TypeOf((**MetricType)(nil)).Elem()

type MetricTypePtrInput interface {
	pulumi.Input

	ToMetricTypePtrOutput() MetricTypePtrOutput
	ToMetricTypePtrOutputWithContext(context.Context) MetricTypePtrOutput
}

type metricTypePtr string

func MetricTypePtr(v string) MetricTypePtrInput {
	return (*metricTypePtr)(&v)
}

func (*metricTypePtr) ElementType() reflect.Type {
	return metricTypePtrType
}

func (in *metricTypePtr) ToMetricTypePtrOutput() MetricTypePtrOutput {
	return pulumi.ToOutput(in).(MetricTypePtrOutput)
}

func (in *metricTypePtr) ToMetricTypePtrOutputWithContext(ctx context.Context) MetricTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(MetricTypePtrOutput)
}

type NotificationOperatorType string

const (
	NotificationOperatorTypeEqualTo              = NotificationOperatorType("EqualTo")
	NotificationOperatorTypeGreaterThan          = NotificationOperatorType("GreaterThan")
	NotificationOperatorTypeGreaterThanOrEqualTo = NotificationOperatorType("GreaterThanOrEqualTo")
)

func (NotificationOperatorType) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationOperatorType)(nil)).Elem()
}

func (e NotificationOperatorType) ToNotificationOperatorTypeOutput() NotificationOperatorTypeOutput {
	return pulumi.ToOutput(e).(NotificationOperatorTypeOutput)
}

func (e NotificationOperatorType) ToNotificationOperatorTypeOutputWithContext(ctx context.Context) NotificationOperatorTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(NotificationOperatorTypeOutput)
}

func (e NotificationOperatorType) ToNotificationOperatorTypePtrOutput() NotificationOperatorTypePtrOutput {
	return e.ToNotificationOperatorTypePtrOutputWithContext(context.Background())
}

func (e NotificationOperatorType) ToNotificationOperatorTypePtrOutputWithContext(ctx context.Context) NotificationOperatorTypePtrOutput {
	return NotificationOperatorType(e).ToNotificationOperatorTypeOutputWithContext(ctx).ToNotificationOperatorTypePtrOutputWithContext(ctx)
}

func (e NotificationOperatorType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationOperatorType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e NotificationOperatorType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e NotificationOperatorType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type NotificationOperatorTypeOutput struct{ *pulumi.OutputState }

func (NotificationOperatorTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationOperatorType)(nil)).Elem()
}

func (o NotificationOperatorTypeOutput) ToNotificationOperatorTypeOutput() NotificationOperatorTypeOutput {
	return o
}

func (o NotificationOperatorTypeOutput) ToNotificationOperatorTypeOutputWithContext(ctx context.Context) NotificationOperatorTypeOutput {
	return o
}

func (o NotificationOperatorTypeOutput) ToNotificationOperatorTypePtrOutput() NotificationOperatorTypePtrOutput {
	return o.ToNotificationOperatorTypePtrOutputWithContext(context.Background())
}

func (o NotificationOperatorTypeOutput) ToNotificationOperatorTypePtrOutputWithContext(ctx context.Context) NotificationOperatorTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v NotificationOperatorType) *NotificationOperatorType {
		return &v
	}).(NotificationOperatorTypePtrOutput)
}

func (o NotificationOperatorTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o NotificationOperatorTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationOperatorType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o NotificationOperatorTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationOperatorTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e NotificationOperatorType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type NotificationOperatorTypePtrOutput struct{ *pulumi.OutputState }

func (NotificationOperatorTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**NotificationOperatorType)(nil)).Elem()
}

func (o NotificationOperatorTypePtrOutput) ToNotificationOperatorTypePtrOutput() NotificationOperatorTypePtrOutput {
	return o
}

func (o NotificationOperatorTypePtrOutput) ToNotificationOperatorTypePtrOutputWithContext(ctx context.Context) NotificationOperatorTypePtrOutput {
	return o
}

func (o NotificationOperatorTypePtrOutput) Elem() NotificationOperatorTypeOutput {
	return o.ApplyT(func(v *NotificationOperatorType) NotificationOperatorType {
		if v != nil {
			return *v
		}
		var ret NotificationOperatorType
		return ret
	}).(NotificationOperatorTypeOutput)
}

func (o NotificationOperatorTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o NotificationOperatorTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *NotificationOperatorType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type NotificationOperatorTypeInput interface {
	pulumi.Input

	ToNotificationOperatorTypeOutput() NotificationOperatorTypeOutput
	ToNotificationOperatorTypeOutputWithContext(context.Context) NotificationOperatorTypeOutput
}

var notificationOperatorTypePtrType = reflect.TypeOf((**NotificationOperatorType)(nil)).Elem()

type NotificationOperatorTypePtrInput interface {
	pulumi.Input

	ToNotificationOperatorTypePtrOutput() NotificationOperatorTypePtrOutput
	ToNotificationOperatorTypePtrOutputWithContext(context.Context) NotificationOperatorTypePtrOutput
}

type notificationOperatorTypePtr string

func NotificationOperatorTypePtr(v string) NotificationOperatorTypePtrInput {
	return (*notificationOperatorTypePtr)(&v)
}

func (*notificationOperatorTypePtr) ElementType() reflect.Type {
	return notificationOperatorTypePtrType
}

func (in *notificationOperatorTypePtr) ToNotificationOperatorTypePtrOutput() NotificationOperatorTypePtrOutput {
	return pulumi.ToOutput(in).(NotificationOperatorTypePtrOutput)
}

func (in *notificationOperatorTypePtr) ToNotificationOperatorTypePtrOutputWithContext(ctx context.Context) NotificationOperatorTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(NotificationOperatorTypePtrOutput)
}

type OperatorType string

const (
	OperatorTypeIn       = OperatorType("In")
	OperatorTypeContains = OperatorType("Contains")
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

type PivotTypeType string

const (
	PivotTypeTypeDimension = PivotTypeType("Dimension")
	PivotTypeTypeTagKey    = PivotTypeType("TagKey")
)

func (PivotTypeType) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotTypeType)(nil)).Elem()
}

func (e PivotTypeType) ToPivotTypeTypeOutput() PivotTypeTypeOutput {
	return pulumi.ToOutput(e).(PivotTypeTypeOutput)
}

func (e PivotTypeType) ToPivotTypeTypeOutputWithContext(ctx context.Context) PivotTypeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PivotTypeTypeOutput)
}

func (e PivotTypeType) ToPivotTypeTypePtrOutput() PivotTypeTypePtrOutput {
	return e.ToPivotTypeTypePtrOutputWithContext(context.Background())
}

func (e PivotTypeType) ToPivotTypeTypePtrOutputWithContext(ctx context.Context) PivotTypeTypePtrOutput {
	return PivotTypeType(e).ToPivotTypeTypeOutputWithContext(ctx).ToPivotTypeTypePtrOutputWithContext(ctx)
}

func (e PivotTypeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PivotTypeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PivotTypeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PivotTypeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PivotTypeTypeOutput struct{ *pulumi.OutputState }

func (PivotTypeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PivotTypeType)(nil)).Elem()
}

func (o PivotTypeTypeOutput) ToPivotTypeTypeOutput() PivotTypeTypeOutput {
	return o
}

func (o PivotTypeTypeOutput) ToPivotTypeTypeOutputWithContext(ctx context.Context) PivotTypeTypeOutput {
	return o
}

func (o PivotTypeTypeOutput) ToPivotTypeTypePtrOutput() PivotTypeTypePtrOutput {
	return o.ToPivotTypeTypePtrOutputWithContext(context.Background())
}

func (o PivotTypeTypeOutput) ToPivotTypeTypePtrOutputWithContext(ctx context.Context) PivotTypeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PivotTypeType) *PivotTypeType {
		return &v
	}).(PivotTypeTypePtrOutput)
}

func (o PivotTypeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PivotTypeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PivotTypeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PivotTypeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PivotTypeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PivotTypeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PivotTypeTypePtrOutput struct{ *pulumi.OutputState }

func (PivotTypeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PivotTypeType)(nil)).Elem()
}

func (o PivotTypeTypePtrOutput) ToPivotTypeTypePtrOutput() PivotTypeTypePtrOutput {
	return o
}

func (o PivotTypeTypePtrOutput) ToPivotTypeTypePtrOutputWithContext(ctx context.Context) PivotTypeTypePtrOutput {
	return o
}

func (o PivotTypeTypePtrOutput) Elem() PivotTypeTypeOutput {
	return o.ApplyT(func(v *PivotTypeType) PivotTypeType {
		if v != nil {
			return *v
		}
		var ret PivotTypeType
		return ret
	}).(PivotTypeTypeOutput)
}

func (o PivotTypeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PivotTypeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PivotTypeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PivotTypeTypeInput interface {
	pulumi.Input

	ToPivotTypeTypeOutput() PivotTypeTypeOutput
	ToPivotTypeTypeOutputWithContext(context.Context) PivotTypeTypeOutput
}

var pivotTypeTypePtrType = reflect.TypeOf((**PivotTypeType)(nil)).Elem()

type PivotTypeTypePtrInput interface {
	pulumi.Input

	ToPivotTypeTypePtrOutput() PivotTypeTypePtrOutput
	ToPivotTypeTypePtrOutputWithContext(context.Context) PivotTypeTypePtrOutput
}

type pivotTypeTypePtr string

func PivotTypeTypePtr(v string) PivotTypeTypePtrInput {
	return (*pivotTypeTypePtr)(&v)
}

func (*pivotTypeTypePtr) ElementType() reflect.Type {
	return pivotTypeTypePtrType
}

func (in *pivotTypeTypePtr) ToPivotTypeTypePtrOutput() PivotTypeTypePtrOutput {
	return pulumi.ToOutput(in).(PivotTypeTypePtrOutput)
}

func (in *pivotTypeTypePtr) ToPivotTypeTypePtrOutputWithContext(ctx context.Context) PivotTypeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PivotTypeTypePtrOutput)
}

type ReportConfigColumnType string

const (
	ReportConfigColumnTypeTag       = ReportConfigColumnType("Tag")
	ReportConfigColumnTypeDimension = ReportConfigColumnType("Dimension")
)

func (ReportConfigColumnType) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigColumnType)(nil)).Elem()
}

func (e ReportConfigColumnType) ToReportConfigColumnTypeOutput() ReportConfigColumnTypeOutput {
	return pulumi.ToOutput(e).(ReportConfigColumnTypeOutput)
}

func (e ReportConfigColumnType) ToReportConfigColumnTypeOutputWithContext(ctx context.Context) ReportConfigColumnTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReportConfigColumnTypeOutput)
}

func (e ReportConfigColumnType) ToReportConfigColumnTypePtrOutput() ReportConfigColumnTypePtrOutput {
	return e.ToReportConfigColumnTypePtrOutputWithContext(context.Background())
}

func (e ReportConfigColumnType) ToReportConfigColumnTypePtrOutputWithContext(ctx context.Context) ReportConfigColumnTypePtrOutput {
	return ReportConfigColumnType(e).ToReportConfigColumnTypeOutputWithContext(ctx).ToReportConfigColumnTypePtrOutputWithContext(ctx)
}

func (e ReportConfigColumnType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReportConfigColumnType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReportConfigColumnType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReportConfigColumnType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReportConfigColumnTypeOutput struct{ *pulumi.OutputState }

func (ReportConfigColumnTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportConfigColumnType)(nil)).Elem()
}

func (o ReportConfigColumnTypeOutput) ToReportConfigColumnTypeOutput() ReportConfigColumnTypeOutput {
	return o
}

func (o ReportConfigColumnTypeOutput) ToReportConfigColumnTypeOutputWithContext(ctx context.Context) ReportConfigColumnTypeOutput {
	return o
}

func (o ReportConfigColumnTypeOutput) ToReportConfigColumnTypePtrOutput() ReportConfigColumnTypePtrOutput {
	return o.ToReportConfigColumnTypePtrOutputWithContext(context.Background())
}

func (o ReportConfigColumnTypeOutput) ToReportConfigColumnTypePtrOutputWithContext(ctx context.Context) ReportConfigColumnTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportConfigColumnType) *ReportConfigColumnType {
		return &v
	}).(ReportConfigColumnTypePtrOutput)
}

func (o ReportConfigColumnTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReportConfigColumnTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReportConfigColumnType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReportConfigColumnTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReportConfigColumnTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReportConfigColumnType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReportConfigColumnTypePtrOutput struct{ *pulumi.OutputState }

func (ReportConfigColumnTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportConfigColumnType)(nil)).Elem()
}

func (o ReportConfigColumnTypePtrOutput) ToReportConfigColumnTypePtrOutput() ReportConfigColumnTypePtrOutput {
	return o
}

func (o ReportConfigColumnTypePtrOutput) ToReportConfigColumnTypePtrOutputWithContext(ctx context.Context) ReportConfigColumnTypePtrOutput {
	return o
}

func (o ReportConfigColumnTypePtrOutput) Elem() ReportConfigColumnTypeOutput {
	return o.ApplyT(func(v *ReportConfigColumnType) ReportConfigColumnType {
		if v != nil {
			return *v
		}
		var ret ReportConfigColumnType
		return ret
	}).(ReportConfigColumnTypeOutput)
}

func (o ReportConfigColumnTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReportConfigColumnTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReportConfigColumnType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReportConfigColumnTypeInput interface {
	pulumi.Input

	ToReportConfigColumnTypeOutput() ReportConfigColumnTypeOutput
	ToReportConfigColumnTypeOutputWithContext(context.Context) ReportConfigColumnTypeOutput
}

var reportConfigColumnTypePtrType = reflect.TypeOf((**ReportConfigColumnType)(nil)).Elem()

type ReportConfigColumnTypePtrInput interface {
	pulumi.Input

	ToReportConfigColumnTypePtrOutput() ReportConfigColumnTypePtrOutput
	ToReportConfigColumnTypePtrOutputWithContext(context.Context) ReportConfigColumnTypePtrOutput
}

type reportConfigColumnTypePtr string

func ReportConfigColumnTypePtr(v string) ReportConfigColumnTypePtrInput {
	return (*reportConfigColumnTypePtr)(&v)
}

func (*reportConfigColumnTypePtr) ElementType() reflect.Type {
	return reportConfigColumnTypePtrType
}

func (in *reportConfigColumnTypePtr) ToReportConfigColumnTypePtrOutput() ReportConfigColumnTypePtrOutput {
	return pulumi.ToOutput(in).(ReportConfigColumnTypePtrOutput)
}

func (in *reportConfigColumnTypePtr) ToReportConfigColumnTypePtrOutputWithContext(ctx context.Context) ReportConfigColumnTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReportConfigColumnTypePtrOutput)
}

type ReportType string

const (
	ReportTypeUsage = ReportType("Usage")
)

func (ReportType) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportType)(nil)).Elem()
}

func (e ReportType) ToReportTypeOutput() ReportTypeOutput {
	return pulumi.ToOutput(e).(ReportTypeOutput)
}

func (e ReportType) ToReportTypeOutputWithContext(ctx context.Context) ReportTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ReportTypeOutput)
}

func (e ReportType) ToReportTypePtrOutput() ReportTypePtrOutput {
	return e.ToReportTypePtrOutputWithContext(context.Background())
}

func (e ReportType) ToReportTypePtrOutputWithContext(ctx context.Context) ReportTypePtrOutput {
	return ReportType(e).ToReportTypeOutputWithContext(ctx).ToReportTypePtrOutputWithContext(ctx)
}

func (e ReportType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReportType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ReportType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ReportType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ReportTypeOutput struct{ *pulumi.OutputState }

func (ReportTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ReportType)(nil)).Elem()
}

func (o ReportTypeOutput) ToReportTypeOutput() ReportTypeOutput {
	return o
}

func (o ReportTypeOutput) ToReportTypeOutputWithContext(ctx context.Context) ReportTypeOutput {
	return o
}

func (o ReportTypeOutput) ToReportTypePtrOutput() ReportTypePtrOutput {
	return o.ToReportTypePtrOutputWithContext(context.Background())
}

func (o ReportTypeOutput) ToReportTypePtrOutputWithContext(ctx context.Context) ReportTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ReportType) *ReportType {
		return &v
	}).(ReportTypePtrOutput)
}

func (o ReportTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ReportTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReportType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ReportTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReportTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ReportType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ReportTypePtrOutput struct{ *pulumi.OutputState }

func (ReportTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ReportType)(nil)).Elem()
}

func (o ReportTypePtrOutput) ToReportTypePtrOutput() ReportTypePtrOutput {
	return o
}

func (o ReportTypePtrOutput) ToReportTypePtrOutputWithContext(ctx context.Context) ReportTypePtrOutput {
	return o
}

func (o ReportTypePtrOutput) Elem() ReportTypeOutput {
	return o.ApplyT(func(v *ReportType) ReportType {
		if v != nil {
			return *v
		}
		var ret ReportType
		return ret
	}).(ReportTypeOutput)
}

func (o ReportTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ReportTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ReportType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ReportTypeInput interface {
	pulumi.Input

	ToReportTypeOutput() ReportTypeOutput
	ToReportTypeOutputWithContext(context.Context) ReportTypeOutput
}

var reportTypePtrType = reflect.TypeOf((**ReportType)(nil)).Elem()

type ReportTypePtrInput interface {
	pulumi.Input

	ToReportTypePtrOutput() ReportTypePtrOutput
	ToReportTypePtrOutputWithContext(context.Context) ReportTypePtrOutput
}

type reportTypePtr string

func ReportTypePtr(v string) ReportTypePtrInput {
	return (*reportTypePtr)(&v)
}

func (*reportTypePtr) ElementType() reflect.Type {
	return reportTypePtrType
}

func (in *reportTypePtr) ToReportTypePtrOutput() ReportTypePtrOutput {
	return pulumi.ToOutput(in).(ReportTypePtrOutput)
}

func (in *reportTypePtr) ToReportTypePtrOutputWithContext(ctx context.Context) ReportTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ReportTypePtrOutput)
}

type TimeGrainType string

const (
	TimeGrainTypeMonthly   = TimeGrainType("Monthly")
	TimeGrainTypeQuarterly = TimeGrainType("Quarterly")
	TimeGrainTypeAnnually  = TimeGrainType("Annually")
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

type TimeframeType string

const (
	TimeframeTypeWeekToDate  = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate = TimeframeType("MonthToDate")
	TimeframeTypeYearToDate  = TimeframeType("YearToDate")
	TimeframeTypeCustom      = TimeframeType("Custom")
)

func (TimeframeType) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeframeType)(nil)).Elem()
}

func (e TimeframeType) ToTimeframeTypeOutput() TimeframeTypeOutput {
	return pulumi.ToOutput(e).(TimeframeTypeOutput)
}

func (e TimeframeType) ToTimeframeTypeOutputWithContext(ctx context.Context) TimeframeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(TimeframeTypeOutput)
}

func (e TimeframeType) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return e.ToTimeframeTypePtrOutputWithContext(context.Background())
}

func (e TimeframeType) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return TimeframeType(e).ToTimeframeTypeOutputWithContext(ctx).ToTimeframeTypePtrOutputWithContext(ctx)
}

func (e TimeframeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeframeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e TimeframeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e TimeframeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type TimeframeTypeOutput struct{ *pulumi.OutputState }

func (TimeframeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TimeframeType)(nil)).Elem()
}

func (o TimeframeTypeOutput) ToTimeframeTypeOutput() TimeframeTypeOutput {
	return o
}

func (o TimeframeTypeOutput) ToTimeframeTypeOutputWithContext(ctx context.Context) TimeframeTypeOutput {
	return o
}

func (o TimeframeTypeOutput) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return o.ToTimeframeTypePtrOutputWithContext(context.Background())
}

func (o TimeframeTypeOutput) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TimeframeType) *TimeframeType {
		return &v
	}).(TimeframeTypePtrOutput)
}

func (o TimeframeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o TimeframeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeframeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o TimeframeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeframeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e TimeframeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type TimeframeTypePtrOutput struct{ *pulumi.OutputState }

func (TimeframeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TimeframeType)(nil)).Elem()
}

func (o TimeframeTypePtrOutput) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return o
}

func (o TimeframeTypePtrOutput) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return o
}

func (o TimeframeTypePtrOutput) Elem() TimeframeTypeOutput {
	return o.ApplyT(func(v *TimeframeType) TimeframeType {
		if v != nil {
			return *v
		}
		var ret TimeframeType
		return ret
	}).(TimeframeTypeOutput)
}

func (o TimeframeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o TimeframeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *TimeframeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type TimeframeTypeInput interface {
	pulumi.Input

	ToTimeframeTypeOutput() TimeframeTypeOutput
	ToTimeframeTypeOutputWithContext(context.Context) TimeframeTypeOutput
}

var timeframeTypePtrType = reflect.TypeOf((**TimeframeType)(nil)).Elem()

type TimeframeTypePtrInput interface {
	pulumi.Input

	ToTimeframeTypePtrOutput() TimeframeTypePtrOutput
	ToTimeframeTypePtrOutputWithContext(context.Context) TimeframeTypePtrOutput
}

type timeframeTypePtr string

func TimeframeTypePtr(v string) TimeframeTypePtrInput {
	return (*timeframeTypePtr)(&v)
}

func (*timeframeTypePtr) ElementType() reflect.Type {
	return timeframeTypePtrType
}

func (in *timeframeTypePtr) ToTimeframeTypePtrOutput() TimeframeTypePtrOutput {
	return pulumi.ToOutput(in).(TimeframeTypePtrOutput)
}

func (in *timeframeTypePtr) ToTimeframeTypePtrOutputWithContext(ctx context.Context) TimeframeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(TimeframeTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccumulatedTypeInput)(nil)).Elem(), AccumulatedType("true"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccumulatedTypePtrInput)(nil)).Elem(), AccumulatedType("true"))
	pulumi.RegisterInputType(reflect.TypeOf((*CategoryTypeInput)(nil)).Elem(), CategoryType("Cost"))
	pulumi.RegisterInputType(reflect.TypeOf((*CategoryTypePtrInput)(nil)).Elem(), CategoryType("Cost"))
	pulumi.RegisterInputType(reflect.TypeOf((*ChartTypeInput)(nil)).Elem(), ChartType("Area"))
	pulumi.RegisterInputType(reflect.TypeOf((*ChartTypePtrInput)(nil)).Elem(), ChartType("Area"))
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTypeInput)(nil)).Elem(), FunctionType("Sum"))
	pulumi.RegisterInputType(reflect.TypeOf((*FunctionTypePtrInput)(nil)).Elem(), FunctionType("Sum"))
	pulumi.RegisterInputType(reflect.TypeOf((*GranularityTypeInput)(nil)).Elem(), GranularityType("Daily"))
	pulumi.RegisterInputType(reflect.TypeOf((*GranularityTypePtrInput)(nil)).Elem(), GranularityType("Daily"))
	pulumi.RegisterInputType(reflect.TypeOf((*KpiTypeTypeInput)(nil)).Elem(), KpiTypeType("Forecast"))
	pulumi.RegisterInputType(reflect.TypeOf((*KpiTypeTypePtrInput)(nil)).Elem(), KpiTypeType("Forecast"))
	pulumi.RegisterInputType(reflect.TypeOf((*MetricTypeInput)(nil)).Elem(), MetricType("ActualCost"))
	pulumi.RegisterInputType(reflect.TypeOf((*MetricTypePtrInput)(nil)).Elem(), MetricType("ActualCost"))
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationOperatorTypeInput)(nil)).Elem(), NotificationOperatorType("EqualTo"))
	pulumi.RegisterInputType(reflect.TypeOf((*NotificationOperatorTypePtrInput)(nil)).Elem(), NotificationOperatorType("EqualTo"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorTypeInput)(nil)).Elem(), OperatorType("In"))
	pulumi.RegisterInputType(reflect.TypeOf((*OperatorTypePtrInput)(nil)).Elem(), OperatorType("In"))
	pulumi.RegisterInputType(reflect.TypeOf((*PivotTypeTypeInput)(nil)).Elem(), PivotTypeType("Dimension"))
	pulumi.RegisterInputType(reflect.TypeOf((*PivotTypeTypePtrInput)(nil)).Elem(), PivotTypeType("Dimension"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReportConfigColumnTypeInput)(nil)).Elem(), ReportConfigColumnType("Tag"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReportConfigColumnTypePtrInput)(nil)).Elem(), ReportConfigColumnType("Tag"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReportTypeInput)(nil)).Elem(), ReportType("Usage"))
	pulumi.RegisterInputType(reflect.TypeOf((*ReportTypePtrInput)(nil)).Elem(), ReportType("Usage"))
	pulumi.RegisterInputType(reflect.TypeOf((*TimeGrainTypeInput)(nil)).Elem(), TimeGrainType("Monthly"))
	pulumi.RegisterInputType(reflect.TypeOf((*TimeGrainTypePtrInput)(nil)).Elem(), TimeGrainType("Monthly"))
	pulumi.RegisterInputType(reflect.TypeOf((*TimeframeTypeInput)(nil)).Elem(), TimeframeType("WeekToDate"))
	pulumi.RegisterInputType(reflect.TypeOf((*TimeframeTypePtrInput)(nil)).Elem(), TimeframeType("WeekToDate"))
	pulumi.RegisterOutputType(AccumulatedTypeOutput{})
	pulumi.RegisterOutputType(AccumulatedTypePtrOutput{})
	pulumi.RegisterOutputType(CategoryTypeOutput{})
	pulumi.RegisterOutputType(CategoryTypePtrOutput{})
	pulumi.RegisterOutputType(ChartTypeOutput{})
	pulumi.RegisterOutputType(ChartTypePtrOutput{})
	pulumi.RegisterOutputType(FunctionTypeOutput{})
	pulumi.RegisterOutputType(FunctionTypePtrOutput{})
	pulumi.RegisterOutputType(GranularityTypeOutput{})
	pulumi.RegisterOutputType(GranularityTypePtrOutput{})
	pulumi.RegisterOutputType(KpiTypeTypeOutput{})
	pulumi.RegisterOutputType(KpiTypeTypePtrOutput{})
	pulumi.RegisterOutputType(MetricTypeOutput{})
	pulumi.RegisterOutputType(MetricTypePtrOutput{})
	pulumi.RegisterOutputType(NotificationOperatorTypeOutput{})
	pulumi.RegisterOutputType(NotificationOperatorTypePtrOutput{})
	pulumi.RegisterOutputType(OperatorTypeOutput{})
	pulumi.RegisterOutputType(OperatorTypePtrOutput{})
	pulumi.RegisterOutputType(PivotTypeTypeOutput{})
	pulumi.RegisterOutputType(PivotTypeTypePtrOutput{})
	pulumi.RegisterOutputType(ReportConfigColumnTypeOutput{})
	pulumi.RegisterOutputType(ReportConfigColumnTypePtrOutput{})
	pulumi.RegisterOutputType(ReportTypeOutput{})
	pulumi.RegisterOutputType(ReportTypePtrOutput{})
	pulumi.RegisterOutputType(TimeGrainTypeOutput{})
	pulumi.RegisterOutputType(TimeGrainTypePtrOutput{})
	pulumi.RegisterOutputType(TimeframeTypeOutput{})
	pulumi.RegisterOutputType(TimeframeTypePtrOutput{})
}
