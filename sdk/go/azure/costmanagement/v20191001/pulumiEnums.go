


package v20191001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExportType string

const (
	ExportTypeUsage = ExportType("Usage")
)

func (ExportType) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportType)(nil)).Elem()
}

func (e ExportType) ToExportTypeOutput() ExportTypeOutput {
	return pulumi.ToOutput(e).(ExportTypeOutput)
}

func (e ExportType) ToExportTypeOutputWithContext(ctx context.Context) ExportTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExportTypeOutput)
}

func (e ExportType) ToExportTypePtrOutput() ExportTypePtrOutput {
	return e.ToExportTypePtrOutputWithContext(context.Background())
}

func (e ExportType) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return ExportType(e).ToExportTypeOutputWithContext(ctx).ToExportTypePtrOutputWithContext(ctx)
}

func (e ExportType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExportType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExportType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExportTypeOutput struct{ *pulumi.OutputState }

func (ExportTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExportType)(nil)).Elem()
}

func (o ExportTypeOutput) ToExportTypeOutput() ExportTypeOutput {
	return o
}

func (o ExportTypeOutput) ToExportTypeOutputWithContext(ctx context.Context) ExportTypeOutput {
	return o
}

func (o ExportTypeOutput) ToExportTypePtrOutput() ExportTypePtrOutput {
	return o.ToExportTypePtrOutputWithContext(context.Background())
}

func (o ExportTypeOutput) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExportType) *ExportType {
		return &v
	}).(ExportTypePtrOutput)
}

func (o ExportTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExportTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExportTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExportType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExportTypePtrOutput struct{ *pulumi.OutputState }

func (ExportTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExportType)(nil)).Elem()
}

func (o ExportTypePtrOutput) ToExportTypePtrOutput() ExportTypePtrOutput {
	return o
}

func (o ExportTypePtrOutput) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return o
}

func (o ExportTypePtrOutput) Elem() ExportTypeOutput {
	return o.ApplyT(func(v *ExportType) ExportType {
		if v != nil {
			return *v
		}
		var ret ExportType
		return ret
	}).(ExportTypeOutput)
}

func (o ExportTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExportTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExportType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExportTypeInput interface {
	pulumi.Input

	ToExportTypeOutput() ExportTypeOutput
	ToExportTypeOutputWithContext(context.Context) ExportTypeOutput
}

var exportTypePtrType = reflect.TypeOf((**ExportType)(nil)).Elem()

type ExportTypePtrInput interface {
	pulumi.Input

	ToExportTypePtrOutput() ExportTypePtrOutput
	ToExportTypePtrOutputWithContext(context.Context) ExportTypePtrOutput
}

type exportTypePtr string

func ExportTypePtr(v string) ExportTypePtrInput {
	return (*exportTypePtr)(&v)
}

func (*exportTypePtr) ElementType() reflect.Type {
	return exportTypePtrType
}

func (in *exportTypePtr) ToExportTypePtrOutput() ExportTypePtrOutput {
	return pulumi.ToOutput(in).(ExportTypePtrOutput)
}

func (in *exportTypePtr) ToExportTypePtrOutputWithContext(ctx context.Context) ExportTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExportTypePtrOutput)
}

type FormatType string

const (
	FormatTypeCsv = FormatType("Csv")
)

func (FormatType) ElementType() reflect.Type {
	return reflect.TypeOf((*FormatType)(nil)).Elem()
}

func (e FormatType) ToFormatTypeOutput() FormatTypeOutput {
	return pulumi.ToOutput(e).(FormatTypeOutput)
}

func (e FormatType) ToFormatTypeOutputWithContext(ctx context.Context) FormatTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(FormatTypeOutput)
}

func (e FormatType) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return e.ToFormatTypePtrOutputWithContext(context.Background())
}

func (e FormatType) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return FormatType(e).ToFormatTypeOutputWithContext(ctx).ToFormatTypePtrOutputWithContext(ctx)
}

func (e FormatType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e FormatType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e FormatType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e FormatType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type FormatTypeOutput struct{ *pulumi.OutputState }

func (FormatTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FormatType)(nil)).Elem()
}

func (o FormatTypeOutput) ToFormatTypeOutput() FormatTypeOutput {
	return o
}

func (o FormatTypeOutput) ToFormatTypeOutputWithContext(ctx context.Context) FormatTypeOutput {
	return o
}

func (o FormatTypeOutput) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return o.ToFormatTypePtrOutputWithContext(context.Background())
}

func (o FormatTypeOutput) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FormatType) *FormatType {
		return &v
	}).(FormatTypePtrOutput)
}

func (o FormatTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o FormatTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FormatType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o FormatTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FormatTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e FormatType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type FormatTypePtrOutput struct{ *pulumi.OutputState }

func (FormatTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FormatType)(nil)).Elem()
}

func (o FormatTypePtrOutput) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return o
}

func (o FormatTypePtrOutput) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return o
}

func (o FormatTypePtrOutput) Elem() FormatTypeOutput {
	return o.ApplyT(func(v *FormatType) FormatType {
		if v != nil {
			return *v
		}
		var ret FormatType
		return ret
	}).(FormatTypeOutput)
}

func (o FormatTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o FormatTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *FormatType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type FormatTypeInput interface {
	pulumi.Input

	ToFormatTypeOutput() FormatTypeOutput
	ToFormatTypeOutputWithContext(context.Context) FormatTypeOutput
}

var formatTypePtrType = reflect.TypeOf((**FormatType)(nil)).Elem()

type FormatTypePtrInput interface {
	pulumi.Input

	ToFormatTypePtrOutput() FormatTypePtrOutput
	ToFormatTypePtrOutputWithContext(context.Context) FormatTypePtrOutput
}

type formatTypePtr string

func FormatTypePtr(v string) FormatTypePtrInput {
	return (*formatTypePtr)(&v)
}

func (*formatTypePtr) ElementType() reflect.Type {
	return formatTypePtrType
}

func (in *formatTypePtr) ToFormatTypePtrOutput() FormatTypePtrOutput {
	return pulumi.ToOutput(in).(FormatTypePtrOutput)
}

func (in *formatTypePtr) ToFormatTypePtrOutputWithContext(ctx context.Context) FormatTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(FormatTypePtrOutput)
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
	GranularityTypeDaily  = GranularityType("Daily")
	GranularityTypeHourly = GranularityType("Hourly")
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

type OperatorType string

const (
	OperatorTypeIn = OperatorType("In")
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

type QueryColumnType string

const (
	QueryColumnTypeTag       = QueryColumnType("Tag")
	QueryColumnTypeDimension = QueryColumnType("Dimension")
)

func (QueryColumnType) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryColumnType)(nil)).Elem()
}

func (e QueryColumnType) ToQueryColumnTypeOutput() QueryColumnTypeOutput {
	return pulumi.ToOutput(e).(QueryColumnTypeOutput)
}

func (e QueryColumnType) ToQueryColumnTypeOutputWithContext(ctx context.Context) QueryColumnTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(QueryColumnTypeOutput)
}

func (e QueryColumnType) ToQueryColumnTypePtrOutput() QueryColumnTypePtrOutput {
	return e.ToQueryColumnTypePtrOutputWithContext(context.Background())
}

func (e QueryColumnType) ToQueryColumnTypePtrOutputWithContext(ctx context.Context) QueryColumnTypePtrOutput {
	return QueryColumnType(e).ToQueryColumnTypeOutputWithContext(ctx).ToQueryColumnTypePtrOutputWithContext(ctx)
}

func (e QueryColumnType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e QueryColumnType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e QueryColumnType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e QueryColumnType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type QueryColumnTypeOutput struct{ *pulumi.OutputState }

func (QueryColumnTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueryColumnType)(nil)).Elem()
}

func (o QueryColumnTypeOutput) ToQueryColumnTypeOutput() QueryColumnTypeOutput {
	return o
}

func (o QueryColumnTypeOutput) ToQueryColumnTypeOutputWithContext(ctx context.Context) QueryColumnTypeOutput {
	return o
}

func (o QueryColumnTypeOutput) ToQueryColumnTypePtrOutput() QueryColumnTypePtrOutput {
	return o.ToQueryColumnTypePtrOutputWithContext(context.Background())
}

func (o QueryColumnTypeOutput) ToQueryColumnTypePtrOutputWithContext(ctx context.Context) QueryColumnTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueryColumnType) *QueryColumnType {
		return &v
	}).(QueryColumnTypePtrOutput)
}

func (o QueryColumnTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o QueryColumnTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QueryColumnType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o QueryColumnTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QueryColumnTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e QueryColumnType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type QueryColumnTypePtrOutput struct{ *pulumi.OutputState }

func (QueryColumnTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueryColumnType)(nil)).Elem()
}

func (o QueryColumnTypePtrOutput) ToQueryColumnTypePtrOutput() QueryColumnTypePtrOutput {
	return o
}

func (o QueryColumnTypePtrOutput) ToQueryColumnTypePtrOutputWithContext(ctx context.Context) QueryColumnTypePtrOutput {
	return o
}

func (o QueryColumnTypePtrOutput) Elem() QueryColumnTypeOutput {
	return o.ApplyT(func(v *QueryColumnType) QueryColumnType {
		if v != nil {
			return *v
		}
		var ret QueryColumnType
		return ret
	}).(QueryColumnTypeOutput)
}

func (o QueryColumnTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o QueryColumnTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *QueryColumnType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type QueryColumnTypeInput interface {
	pulumi.Input

	ToQueryColumnTypeOutput() QueryColumnTypeOutput
	ToQueryColumnTypeOutputWithContext(context.Context) QueryColumnTypeOutput
}

var queryColumnTypePtrType = reflect.TypeOf((**QueryColumnType)(nil)).Elem()

type QueryColumnTypePtrInput interface {
	pulumi.Input

	ToQueryColumnTypePtrOutput() QueryColumnTypePtrOutput
	ToQueryColumnTypePtrOutputWithContext(context.Context) QueryColumnTypePtrOutput
}

type queryColumnTypePtr string

func QueryColumnTypePtr(v string) QueryColumnTypePtrInput {
	return (*queryColumnTypePtr)(&v)
}

func (*queryColumnTypePtr) ElementType() reflect.Type {
	return queryColumnTypePtrType
}

func (in *queryColumnTypePtr) ToQueryColumnTypePtrOutput() QueryColumnTypePtrOutput {
	return pulumi.ToOutput(in).(QueryColumnTypePtrOutput)
}

func (in *queryColumnTypePtr) ToQueryColumnTypePtrOutputWithContext(ctx context.Context) QueryColumnTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(QueryColumnTypePtrOutput)
}

type RecurrenceType string

const (
	RecurrenceTypeDaily    = RecurrenceType("Daily")
	RecurrenceTypeWeekly   = RecurrenceType("Weekly")
	RecurrenceTypeMonthly  = RecurrenceType("Monthly")
	RecurrenceTypeAnnually = RecurrenceType("Annually")
)

func (RecurrenceType) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceType)(nil)).Elem()
}

func (e RecurrenceType) ToRecurrenceTypeOutput() RecurrenceTypeOutput {
	return pulumi.ToOutput(e).(RecurrenceTypeOutput)
}

func (e RecurrenceType) ToRecurrenceTypeOutputWithContext(ctx context.Context) RecurrenceTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(RecurrenceTypeOutput)
}

func (e RecurrenceType) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return e.ToRecurrenceTypePtrOutputWithContext(context.Background())
}

func (e RecurrenceType) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return RecurrenceType(e).ToRecurrenceTypeOutputWithContext(ctx).ToRecurrenceTypePtrOutputWithContext(ctx)
}

func (e RecurrenceType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e RecurrenceType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e RecurrenceType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type RecurrenceTypeOutput struct{ *pulumi.OutputState }

func (RecurrenceTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RecurrenceType)(nil)).Elem()
}

func (o RecurrenceTypeOutput) ToRecurrenceTypeOutput() RecurrenceTypeOutput {
	return o
}

func (o RecurrenceTypeOutput) ToRecurrenceTypeOutputWithContext(ctx context.Context) RecurrenceTypeOutput {
	return o
}

func (o RecurrenceTypeOutput) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return o.ToRecurrenceTypePtrOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RecurrenceType) *RecurrenceType {
		return &v
	}).(RecurrenceTypePtrOutput)
}

func (o RecurrenceTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o RecurrenceTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e RecurrenceType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type RecurrenceTypePtrOutput struct{ *pulumi.OutputState }

func (RecurrenceTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RecurrenceType)(nil)).Elem()
}

func (o RecurrenceTypePtrOutput) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return o
}

func (o RecurrenceTypePtrOutput) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return o
}

func (o RecurrenceTypePtrOutput) Elem() RecurrenceTypeOutput {
	return o.ApplyT(func(v *RecurrenceType) RecurrenceType {
		if v != nil {
			return *v
		}
		var ret RecurrenceType
		return ret
	}).(RecurrenceTypeOutput)
}

func (o RecurrenceTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o RecurrenceTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *RecurrenceType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type RecurrenceTypeInput interface {
	pulumi.Input

	ToRecurrenceTypeOutput() RecurrenceTypeOutput
	ToRecurrenceTypeOutputWithContext(context.Context) RecurrenceTypeOutput
}

var recurrenceTypePtrType = reflect.TypeOf((**RecurrenceType)(nil)).Elem()

type RecurrenceTypePtrInput interface {
	pulumi.Input

	ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput
	ToRecurrenceTypePtrOutputWithContext(context.Context) RecurrenceTypePtrOutput
}

type recurrenceTypePtr string

func RecurrenceTypePtr(v string) RecurrenceTypePtrInput {
	return (*recurrenceTypePtr)(&v)
}

func (*recurrenceTypePtr) ElementType() reflect.Type {
	return recurrenceTypePtrType
}

func (in *recurrenceTypePtr) ToRecurrenceTypePtrOutput() RecurrenceTypePtrOutput {
	return pulumi.ToOutput(in).(RecurrenceTypePtrOutput)
}

func (in *recurrenceTypePtr) ToRecurrenceTypePtrOutputWithContext(ctx context.Context) RecurrenceTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(RecurrenceTypePtrOutput)
}

type SortDirection string

const (
	SortDirectionAscending  = SortDirection("Ascending")
	SortDirectionDescending = SortDirection("Descending")
)

func (SortDirection) ElementType() reflect.Type {
	return reflect.TypeOf((*SortDirection)(nil)).Elem()
}

func (e SortDirection) ToSortDirectionOutput() SortDirectionOutput {
	return pulumi.ToOutput(e).(SortDirectionOutput)
}

func (e SortDirection) ToSortDirectionOutputWithContext(ctx context.Context) SortDirectionOutput {
	return pulumi.ToOutputWithContext(ctx, e).(SortDirectionOutput)
}

func (e SortDirection) ToSortDirectionPtrOutput() SortDirectionPtrOutput {
	return e.ToSortDirectionPtrOutputWithContext(context.Background())
}

func (e SortDirection) ToSortDirectionPtrOutputWithContext(ctx context.Context) SortDirectionPtrOutput {
	return SortDirection(e).ToSortDirectionOutputWithContext(ctx).ToSortDirectionPtrOutputWithContext(ctx)
}

func (e SortDirection) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e SortDirection) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e SortDirection) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e SortDirection) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type SortDirectionOutput struct{ *pulumi.OutputState }

func (SortDirectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SortDirection)(nil)).Elem()
}

func (o SortDirectionOutput) ToSortDirectionOutput() SortDirectionOutput {
	return o
}

func (o SortDirectionOutput) ToSortDirectionOutputWithContext(ctx context.Context) SortDirectionOutput {
	return o
}

func (o SortDirectionOutput) ToSortDirectionPtrOutput() SortDirectionPtrOutput {
	return o.ToSortDirectionPtrOutputWithContext(context.Background())
}

func (o SortDirectionOutput) ToSortDirectionPtrOutputWithContext(ctx context.Context) SortDirectionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SortDirection) *SortDirection {
		return &v
	}).(SortDirectionPtrOutput)
}

func (o SortDirectionOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o SortDirectionOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SortDirection) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o SortDirectionOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SortDirectionOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e SortDirection) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type SortDirectionPtrOutput struct{ *pulumi.OutputState }

func (SortDirectionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SortDirection)(nil)).Elem()
}

func (o SortDirectionPtrOutput) ToSortDirectionPtrOutput() SortDirectionPtrOutput {
	return o
}

func (o SortDirectionPtrOutput) ToSortDirectionPtrOutputWithContext(ctx context.Context) SortDirectionPtrOutput {
	return o
}

func (o SortDirectionPtrOutput) Elem() SortDirectionOutput {
	return o.ApplyT(func(v *SortDirection) SortDirection {
		if v != nil {
			return *v
		}
		var ret SortDirection
		return ret
	}).(SortDirectionOutput)
}

func (o SortDirectionPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o SortDirectionPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *SortDirection) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type SortDirectionInput interface {
	pulumi.Input

	ToSortDirectionOutput() SortDirectionOutput
	ToSortDirectionOutputWithContext(context.Context) SortDirectionOutput
}

var sortDirectionPtrType = reflect.TypeOf((**SortDirection)(nil)).Elem()

type SortDirectionPtrInput interface {
	pulumi.Input

	ToSortDirectionPtrOutput() SortDirectionPtrOutput
	ToSortDirectionPtrOutputWithContext(context.Context) SortDirectionPtrOutput
}

type sortDirectionPtr string

func SortDirectionPtr(v string) SortDirectionPtrInput {
	return (*sortDirectionPtr)(&v)
}

func (*sortDirectionPtr) ElementType() reflect.Type {
	return sortDirectionPtrType
}

func (in *sortDirectionPtr) ToSortDirectionPtrOutput() SortDirectionPtrOutput {
	return pulumi.ToOutput(in).(SortDirectionPtrOutput)
}

func (in *sortDirectionPtr) ToSortDirectionPtrOutputWithContext(ctx context.Context) SortDirectionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(SortDirectionPtrOutput)
}

type StatusType string

const (
	StatusTypeActive   = StatusType("Active")
	StatusTypeInactive = StatusType("Inactive")
)

func (StatusType) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusType)(nil)).Elem()
}

func (e StatusType) ToStatusTypeOutput() StatusTypeOutput {
	return pulumi.ToOutput(e).(StatusTypeOutput)
}

func (e StatusType) ToStatusTypeOutputWithContext(ctx context.Context) StatusTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(StatusTypeOutput)
}

func (e StatusType) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return e.ToStatusTypePtrOutputWithContext(context.Background())
}

func (e StatusType) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return StatusType(e).ToStatusTypeOutputWithContext(ctx).ToStatusTypePtrOutputWithContext(ctx)
}

func (e StatusType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e StatusType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e StatusType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type StatusTypeOutput struct{ *pulumi.OutputState }

func (StatusTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusType)(nil)).Elem()
}

func (o StatusTypeOutput) ToStatusTypeOutput() StatusTypeOutput {
	return o
}

func (o StatusTypeOutput) ToStatusTypeOutputWithContext(ctx context.Context) StatusTypeOutput {
	return o
}

func (o StatusTypeOutput) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return o.ToStatusTypePtrOutputWithContext(context.Background())
}

func (o StatusTypeOutput) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StatusType) *StatusType {
		return &v
	}).(StatusTypePtrOutput)
}

func (o StatusTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o StatusTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o StatusTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e StatusType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type StatusTypePtrOutput struct{ *pulumi.OutputState }

func (StatusTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StatusType)(nil)).Elem()
}

func (o StatusTypePtrOutput) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return o
}

func (o StatusTypePtrOutput) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return o
}

func (o StatusTypePtrOutput) Elem() StatusTypeOutput {
	return o.ApplyT(func(v *StatusType) StatusType {
		if v != nil {
			return *v
		}
		var ret StatusType
		return ret
	}).(StatusTypeOutput)
}

func (o StatusTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o StatusTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *StatusType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type StatusTypeInput interface {
	pulumi.Input

	ToStatusTypeOutput() StatusTypeOutput
	ToStatusTypeOutputWithContext(context.Context) StatusTypeOutput
}

var statusTypePtrType = reflect.TypeOf((**StatusType)(nil)).Elem()

type StatusTypePtrInput interface {
	pulumi.Input

	ToStatusTypePtrOutput() StatusTypePtrOutput
	ToStatusTypePtrOutputWithContext(context.Context) StatusTypePtrOutput
}

type statusTypePtr string

func StatusTypePtr(v string) StatusTypePtrInput {
	return (*statusTypePtr)(&v)
}

func (*statusTypePtr) ElementType() reflect.Type {
	return statusTypePtrType
}

func (in *statusTypePtr) ToStatusTypePtrOutput() StatusTypePtrOutput {
	return pulumi.ToOutput(in).(StatusTypePtrOutput)
}

func (in *statusTypePtr) ToStatusTypePtrOutputWithContext(ctx context.Context) StatusTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(StatusTypePtrOutput)
}

type TimeframeType string

const (
	TimeframeTypeWeekToDate   = TimeframeType("WeekToDate")
	TimeframeTypeMonthToDate  = TimeframeType("MonthToDate")
	TimeframeTypeYearToDate   = TimeframeType("YearToDate")
	TimeframeTypeTheLastWeek  = TimeframeType("TheLastWeek")
	TimeframeTypeTheLastMonth = TimeframeType("TheLastMonth")
	TimeframeTypeTheLastYear  = TimeframeType("TheLastYear")
	TimeframeTypeCustom       = TimeframeType("Custom")
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
	pulumi.RegisterOutputType(ExportTypeOutput{})
	pulumi.RegisterOutputType(ExportTypePtrOutput{})
	pulumi.RegisterOutputType(FormatTypeOutput{})
	pulumi.RegisterOutputType(FormatTypePtrOutput{})
	pulumi.RegisterOutputType(FunctionTypeOutput{})
	pulumi.RegisterOutputType(FunctionTypePtrOutput{})
	pulumi.RegisterOutputType(GranularityTypeOutput{})
	pulumi.RegisterOutputType(GranularityTypePtrOutput{})
	pulumi.RegisterOutputType(OperatorTypeOutput{})
	pulumi.RegisterOutputType(OperatorTypePtrOutput{})
	pulumi.RegisterOutputType(QueryColumnTypeOutput{})
	pulumi.RegisterOutputType(QueryColumnTypePtrOutput{})
	pulumi.RegisterOutputType(RecurrenceTypeOutput{})
	pulumi.RegisterOutputType(RecurrenceTypePtrOutput{})
	pulumi.RegisterOutputType(SortDirectionOutput{})
	pulumi.RegisterOutputType(SortDirectionPtrOutput{})
	pulumi.RegisterOutputType(StatusTypeOutput{})
	pulumi.RegisterOutputType(StatusTypePtrOutput{})
	pulumi.RegisterOutputType(TimeframeTypeOutput{})
	pulumi.RegisterOutputType(TimeframeTypePtrOutput{})
}
