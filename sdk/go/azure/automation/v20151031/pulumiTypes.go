


package v20151031

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AdvancedSchedule struct {
	MonthDays          []int                               `pulumi:"monthDays"`
	MonthlyOccurrences []AdvancedScheduleMonthlyOccurrence `pulumi:"monthlyOccurrences"`
	WeekDays           []string                            `pulumi:"weekDays"`
}





type AdvancedScheduleInput interface {
	pulumi.Input

	ToAdvancedScheduleOutput() AdvancedScheduleOutput
	ToAdvancedScheduleOutputWithContext(context.Context) AdvancedScheduleOutput
}

type AdvancedScheduleArgs struct {
	MonthDays          pulumi.IntArrayInput                        `pulumi:"monthDays"`
	MonthlyOccurrences AdvancedScheduleMonthlyOccurrenceArrayInput `pulumi:"monthlyOccurrences"`
	WeekDays           pulumi.StringArrayInput                     `pulumi:"weekDays"`
}

func (AdvancedScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedSchedule)(nil)).Elem()
}

func (i AdvancedScheduleArgs) ToAdvancedScheduleOutput() AdvancedScheduleOutput {
	return i.ToAdvancedScheduleOutputWithContext(context.Background())
}

func (i AdvancedScheduleArgs) ToAdvancedScheduleOutputWithContext(ctx context.Context) AdvancedScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleOutput)
}

func (i AdvancedScheduleArgs) ToAdvancedSchedulePtrOutput() AdvancedSchedulePtrOutput {
	return i.ToAdvancedSchedulePtrOutputWithContext(context.Background())
}

func (i AdvancedScheduleArgs) ToAdvancedSchedulePtrOutputWithContext(ctx context.Context) AdvancedSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleOutput).ToAdvancedSchedulePtrOutputWithContext(ctx)
}









type AdvancedSchedulePtrInput interface {
	pulumi.Input

	ToAdvancedSchedulePtrOutput() AdvancedSchedulePtrOutput
	ToAdvancedSchedulePtrOutputWithContext(context.Context) AdvancedSchedulePtrOutput
}

type advancedSchedulePtrType AdvancedScheduleArgs

func AdvancedSchedulePtr(v *AdvancedScheduleArgs) AdvancedSchedulePtrInput {
	return (*advancedSchedulePtrType)(v)
}

func (*advancedSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedSchedule)(nil)).Elem()
}

func (i *advancedSchedulePtrType) ToAdvancedSchedulePtrOutput() AdvancedSchedulePtrOutput {
	return i.ToAdvancedSchedulePtrOutputWithContext(context.Background())
}

func (i *advancedSchedulePtrType) ToAdvancedSchedulePtrOutputWithContext(ctx context.Context) AdvancedSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedSchedulePtrOutput)
}

type AdvancedScheduleOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedSchedule)(nil)).Elem()
}

func (o AdvancedScheduleOutput) ToAdvancedScheduleOutput() AdvancedScheduleOutput {
	return o
}

func (o AdvancedScheduleOutput) ToAdvancedScheduleOutputWithContext(ctx context.Context) AdvancedScheduleOutput {
	return o
}

func (o AdvancedScheduleOutput) ToAdvancedSchedulePtrOutput() AdvancedSchedulePtrOutput {
	return o.ToAdvancedSchedulePtrOutputWithContext(context.Background())
}

func (o AdvancedScheduleOutput) ToAdvancedSchedulePtrOutputWithContext(ctx context.Context) AdvancedSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdvancedSchedule) *AdvancedSchedule {
		return &v
	}).(AdvancedSchedulePtrOutput)
}

func (o AdvancedScheduleOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v AdvancedSchedule) []int { return v.MonthDays }).(pulumi.IntArrayOutput)
}

func (o AdvancedScheduleOutput) MonthlyOccurrences() AdvancedScheduleMonthlyOccurrenceArrayOutput {
	return o.ApplyT(func(v AdvancedSchedule) []AdvancedScheduleMonthlyOccurrence { return v.MonthlyOccurrences }).(AdvancedScheduleMonthlyOccurrenceArrayOutput)
}

func (o AdvancedScheduleOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdvancedSchedule) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

type AdvancedSchedulePtrOutput struct{ *pulumi.OutputState }

func (AdvancedSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedSchedule)(nil)).Elem()
}

func (o AdvancedSchedulePtrOutput) ToAdvancedSchedulePtrOutput() AdvancedSchedulePtrOutput {
	return o
}

func (o AdvancedSchedulePtrOutput) ToAdvancedSchedulePtrOutputWithContext(ctx context.Context) AdvancedSchedulePtrOutput {
	return o
}

func (o AdvancedSchedulePtrOutput) Elem() AdvancedScheduleOutput {
	return o.ApplyT(func(v *AdvancedSchedule) AdvancedSchedule {
		if v != nil {
			return *v
		}
		var ret AdvancedSchedule
		return ret
	}).(AdvancedScheduleOutput)
}

func (o AdvancedSchedulePtrOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *AdvancedSchedule) []int {
		if v == nil {
			return nil
		}
		return v.MonthDays
	}).(pulumi.IntArrayOutput)
}

func (o AdvancedSchedulePtrOutput) MonthlyOccurrences() AdvancedScheduleMonthlyOccurrenceArrayOutput {
	return o.ApplyT(func(v *AdvancedSchedule) []AdvancedScheduleMonthlyOccurrence {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrences
	}).(AdvancedScheduleMonthlyOccurrenceArrayOutput)
}

func (o AdvancedSchedulePtrOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdvancedSchedule) []string {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(pulumi.StringArrayOutput)
}

type AdvancedScheduleMonthlyOccurrence struct {
	Day        *string `pulumi:"day"`
	Occurrence *int    `pulumi:"occurrence"`
}





type AdvancedScheduleMonthlyOccurrenceInput interface {
	pulumi.Input

	ToAdvancedScheduleMonthlyOccurrenceOutput() AdvancedScheduleMonthlyOccurrenceOutput
	ToAdvancedScheduleMonthlyOccurrenceOutputWithContext(context.Context) AdvancedScheduleMonthlyOccurrenceOutput
}

type AdvancedScheduleMonthlyOccurrenceArgs struct {
	Day        pulumi.StringPtrInput `pulumi:"day"`
	Occurrence pulumi.IntPtrInput    `pulumi:"occurrence"`
}

func (AdvancedScheduleMonthlyOccurrenceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i AdvancedScheduleMonthlyOccurrenceArgs) ToAdvancedScheduleMonthlyOccurrenceOutput() AdvancedScheduleMonthlyOccurrenceOutput {
	return i.ToAdvancedScheduleMonthlyOccurrenceOutputWithContext(context.Background())
}

func (i AdvancedScheduleMonthlyOccurrenceArgs) ToAdvancedScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleMonthlyOccurrenceOutput)
}





type AdvancedScheduleMonthlyOccurrenceArrayInput interface {
	pulumi.Input

	ToAdvancedScheduleMonthlyOccurrenceArrayOutput() AdvancedScheduleMonthlyOccurrenceArrayOutput
	ToAdvancedScheduleMonthlyOccurrenceArrayOutputWithContext(context.Context) AdvancedScheduleMonthlyOccurrenceArrayOutput
}

type AdvancedScheduleMonthlyOccurrenceArray []AdvancedScheduleMonthlyOccurrenceInput

func (AdvancedScheduleMonthlyOccurrenceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdvancedScheduleMonthlyOccurrence)(nil)).Elem()
}

func (i AdvancedScheduleMonthlyOccurrenceArray) ToAdvancedScheduleMonthlyOccurrenceArrayOutput() AdvancedScheduleMonthlyOccurrenceArrayOutput {
	return i.ToAdvancedScheduleMonthlyOccurrenceArrayOutputWithContext(context.Background())
}

func (i AdvancedScheduleMonthlyOccurrenceArray) ToAdvancedScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleMonthlyOccurrenceArrayOutput)
}

type AdvancedScheduleMonthlyOccurrenceOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleMonthlyOccurrenceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o AdvancedScheduleMonthlyOccurrenceOutput) ToAdvancedScheduleMonthlyOccurrenceOutput() AdvancedScheduleMonthlyOccurrenceOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceOutput) ToAdvancedScheduleMonthlyOccurrenceOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdvancedScheduleMonthlyOccurrence) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o AdvancedScheduleMonthlyOccurrenceOutput) Occurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AdvancedScheduleMonthlyOccurrence) *int { return v.Occurrence }).(pulumi.IntPtrOutput)
}

type AdvancedScheduleMonthlyOccurrenceArrayOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleMonthlyOccurrenceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdvancedScheduleMonthlyOccurrence)(nil)).Elem()
}

func (o AdvancedScheduleMonthlyOccurrenceArrayOutput) ToAdvancedScheduleMonthlyOccurrenceArrayOutput() AdvancedScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceArrayOutput) ToAdvancedScheduleMonthlyOccurrenceArrayOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceArrayOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceArrayOutput) Index(i pulumi.IntInput) AdvancedScheduleMonthlyOccurrenceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdvancedScheduleMonthlyOccurrence {
		return vs[0].([]AdvancedScheduleMonthlyOccurrence)[vs[1].(int)]
	}).(AdvancedScheduleMonthlyOccurrenceOutput)
}

type AdvancedScheduleMonthlyOccurrenceResponse struct {
	Day        *string `pulumi:"day"`
	Occurrence *int    `pulumi:"occurrence"`
}





type AdvancedScheduleMonthlyOccurrenceResponseInput interface {
	pulumi.Input

	ToAdvancedScheduleMonthlyOccurrenceResponseOutput() AdvancedScheduleMonthlyOccurrenceResponseOutput
	ToAdvancedScheduleMonthlyOccurrenceResponseOutputWithContext(context.Context) AdvancedScheduleMonthlyOccurrenceResponseOutput
}

type AdvancedScheduleMonthlyOccurrenceResponseArgs struct {
	Day        pulumi.StringPtrInput `pulumi:"day"`
	Occurrence pulumi.IntPtrInput    `pulumi:"occurrence"`
}

func (AdvancedScheduleMonthlyOccurrenceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (i AdvancedScheduleMonthlyOccurrenceResponseArgs) ToAdvancedScheduleMonthlyOccurrenceResponseOutput() AdvancedScheduleMonthlyOccurrenceResponseOutput {
	return i.ToAdvancedScheduleMonthlyOccurrenceResponseOutputWithContext(context.Background())
}

func (i AdvancedScheduleMonthlyOccurrenceResponseArgs) ToAdvancedScheduleMonthlyOccurrenceResponseOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleMonthlyOccurrenceResponseOutput)
}





type AdvancedScheduleMonthlyOccurrenceResponseArrayInput interface {
	pulumi.Input

	ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutput() AdvancedScheduleMonthlyOccurrenceResponseArrayOutput
	ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutputWithContext(context.Context) AdvancedScheduleMonthlyOccurrenceResponseArrayOutput
}

type AdvancedScheduleMonthlyOccurrenceResponseArray []AdvancedScheduleMonthlyOccurrenceResponseInput

func (AdvancedScheduleMonthlyOccurrenceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdvancedScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (i AdvancedScheduleMonthlyOccurrenceResponseArray) ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutput() AdvancedScheduleMonthlyOccurrenceResponseArrayOutput {
	return i.ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutputWithContext(context.Background())
}

func (i AdvancedScheduleMonthlyOccurrenceResponseArray) ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleMonthlyOccurrenceResponseArrayOutput)
}

type AdvancedScheduleMonthlyOccurrenceResponseOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleMonthlyOccurrenceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (o AdvancedScheduleMonthlyOccurrenceResponseOutput) ToAdvancedScheduleMonthlyOccurrenceResponseOutput() AdvancedScheduleMonthlyOccurrenceResponseOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceResponseOutput) ToAdvancedScheduleMonthlyOccurrenceResponseOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceResponseOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceResponseOutput) Day() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AdvancedScheduleMonthlyOccurrenceResponse) *string { return v.Day }).(pulumi.StringPtrOutput)
}

func (o AdvancedScheduleMonthlyOccurrenceResponseOutput) Occurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AdvancedScheduleMonthlyOccurrenceResponse) *int { return v.Occurrence }).(pulumi.IntPtrOutput)
}

type AdvancedScheduleMonthlyOccurrenceResponseArrayOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleMonthlyOccurrenceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AdvancedScheduleMonthlyOccurrenceResponse)(nil)).Elem()
}

func (o AdvancedScheduleMonthlyOccurrenceResponseArrayOutput) ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutput() AdvancedScheduleMonthlyOccurrenceResponseArrayOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceResponseArrayOutput) ToAdvancedScheduleMonthlyOccurrenceResponseArrayOutputWithContext(ctx context.Context) AdvancedScheduleMonthlyOccurrenceResponseArrayOutput {
	return o
}

func (o AdvancedScheduleMonthlyOccurrenceResponseArrayOutput) Index(i pulumi.IntInput) AdvancedScheduleMonthlyOccurrenceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AdvancedScheduleMonthlyOccurrenceResponse {
		return vs[0].([]AdvancedScheduleMonthlyOccurrenceResponse)[vs[1].(int)]
	}).(AdvancedScheduleMonthlyOccurrenceResponseOutput)
}

type AdvancedScheduleResponse struct {
	MonthDays          []int                                       `pulumi:"monthDays"`
	MonthlyOccurrences []AdvancedScheduleMonthlyOccurrenceResponse `pulumi:"monthlyOccurrences"`
	WeekDays           []string                                    `pulumi:"weekDays"`
}





type AdvancedScheduleResponseInput interface {
	pulumi.Input

	ToAdvancedScheduleResponseOutput() AdvancedScheduleResponseOutput
	ToAdvancedScheduleResponseOutputWithContext(context.Context) AdvancedScheduleResponseOutput
}

type AdvancedScheduleResponseArgs struct {
	MonthDays          pulumi.IntArrayInput                                `pulumi:"monthDays"`
	MonthlyOccurrences AdvancedScheduleMonthlyOccurrenceResponseArrayInput `pulumi:"monthlyOccurrences"`
	WeekDays           pulumi.StringArrayInput                             `pulumi:"weekDays"`
}

func (AdvancedScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedScheduleResponse)(nil)).Elem()
}

func (i AdvancedScheduleResponseArgs) ToAdvancedScheduleResponseOutput() AdvancedScheduleResponseOutput {
	return i.ToAdvancedScheduleResponseOutputWithContext(context.Background())
}

func (i AdvancedScheduleResponseArgs) ToAdvancedScheduleResponseOutputWithContext(ctx context.Context) AdvancedScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleResponseOutput)
}

func (i AdvancedScheduleResponseArgs) ToAdvancedScheduleResponsePtrOutput() AdvancedScheduleResponsePtrOutput {
	return i.ToAdvancedScheduleResponsePtrOutputWithContext(context.Background())
}

func (i AdvancedScheduleResponseArgs) ToAdvancedScheduleResponsePtrOutputWithContext(ctx context.Context) AdvancedScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleResponseOutput).ToAdvancedScheduleResponsePtrOutputWithContext(ctx)
}









type AdvancedScheduleResponsePtrInput interface {
	pulumi.Input

	ToAdvancedScheduleResponsePtrOutput() AdvancedScheduleResponsePtrOutput
	ToAdvancedScheduleResponsePtrOutputWithContext(context.Context) AdvancedScheduleResponsePtrOutput
}

type advancedScheduleResponsePtrType AdvancedScheduleResponseArgs

func AdvancedScheduleResponsePtr(v *AdvancedScheduleResponseArgs) AdvancedScheduleResponsePtrInput {
	return (*advancedScheduleResponsePtrType)(v)
}

func (*advancedScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedScheduleResponse)(nil)).Elem()
}

func (i *advancedScheduleResponsePtrType) ToAdvancedScheduleResponsePtrOutput() AdvancedScheduleResponsePtrOutput {
	return i.ToAdvancedScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *advancedScheduleResponsePtrType) ToAdvancedScheduleResponsePtrOutputWithContext(ctx context.Context) AdvancedScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdvancedScheduleResponsePtrOutput)
}

type AdvancedScheduleResponseOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdvancedScheduleResponse)(nil)).Elem()
}

func (o AdvancedScheduleResponseOutput) ToAdvancedScheduleResponseOutput() AdvancedScheduleResponseOutput {
	return o
}

func (o AdvancedScheduleResponseOutput) ToAdvancedScheduleResponseOutputWithContext(ctx context.Context) AdvancedScheduleResponseOutput {
	return o
}

func (o AdvancedScheduleResponseOutput) ToAdvancedScheduleResponsePtrOutput() AdvancedScheduleResponsePtrOutput {
	return o.ToAdvancedScheduleResponsePtrOutputWithContext(context.Background())
}

func (o AdvancedScheduleResponseOutput) ToAdvancedScheduleResponsePtrOutputWithContext(ctx context.Context) AdvancedScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdvancedScheduleResponse) *AdvancedScheduleResponse {
		return &v
	}).(AdvancedScheduleResponsePtrOutput)
}

func (o AdvancedScheduleResponseOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v AdvancedScheduleResponse) []int { return v.MonthDays }).(pulumi.IntArrayOutput)
}

func (o AdvancedScheduleResponseOutput) MonthlyOccurrences() AdvancedScheduleMonthlyOccurrenceResponseArrayOutput {
	return o.ApplyT(func(v AdvancedScheduleResponse) []AdvancedScheduleMonthlyOccurrenceResponse {
		return v.MonthlyOccurrences
	}).(AdvancedScheduleMonthlyOccurrenceResponseArrayOutput)
}

func (o AdvancedScheduleResponseOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AdvancedScheduleResponse) []string { return v.WeekDays }).(pulumi.StringArrayOutput)
}

type AdvancedScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (AdvancedScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdvancedScheduleResponse)(nil)).Elem()
}

func (o AdvancedScheduleResponsePtrOutput) ToAdvancedScheduleResponsePtrOutput() AdvancedScheduleResponsePtrOutput {
	return o
}

func (o AdvancedScheduleResponsePtrOutput) ToAdvancedScheduleResponsePtrOutputWithContext(ctx context.Context) AdvancedScheduleResponsePtrOutput {
	return o
}

func (o AdvancedScheduleResponsePtrOutput) Elem() AdvancedScheduleResponseOutput {
	return o.ApplyT(func(v *AdvancedScheduleResponse) AdvancedScheduleResponse {
		if v != nil {
			return *v
		}
		var ret AdvancedScheduleResponse
		return ret
	}).(AdvancedScheduleResponseOutput)
}

func (o AdvancedScheduleResponsePtrOutput) MonthDays() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *AdvancedScheduleResponse) []int {
		if v == nil {
			return nil
		}
		return v.MonthDays
	}).(pulumi.IntArrayOutput)
}

func (o AdvancedScheduleResponsePtrOutput) MonthlyOccurrences() AdvancedScheduleMonthlyOccurrenceResponseArrayOutput {
	return o.ApplyT(func(v *AdvancedScheduleResponse) []AdvancedScheduleMonthlyOccurrenceResponse {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrences
	}).(AdvancedScheduleMonthlyOccurrenceResponseArrayOutput)
}

func (o AdvancedScheduleResponsePtrOutput) WeekDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AdvancedScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeekDays
	}).(pulumi.StringArrayOutput)
}

type ConnectionTypeAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type ConnectionTypeAssociationPropertyInput interface {
	pulumi.Input

	ToConnectionTypeAssociationPropertyOutput() ConnectionTypeAssociationPropertyOutput
	ToConnectionTypeAssociationPropertyOutputWithContext(context.Context) ConnectionTypeAssociationPropertyOutput
}

type ConnectionTypeAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ConnectionTypeAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionTypeAssociationProperty)(nil)).Elem()
}

func (i ConnectionTypeAssociationPropertyArgs) ToConnectionTypeAssociationPropertyOutput() ConnectionTypeAssociationPropertyOutput {
	return i.ToConnectionTypeAssociationPropertyOutputWithContext(context.Background())
}

func (i ConnectionTypeAssociationPropertyArgs) ToConnectionTypeAssociationPropertyOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeAssociationPropertyOutput)
}

func (i ConnectionTypeAssociationPropertyArgs) ToConnectionTypeAssociationPropertyPtrOutput() ConnectionTypeAssociationPropertyPtrOutput {
	return i.ToConnectionTypeAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i ConnectionTypeAssociationPropertyArgs) ToConnectionTypeAssociationPropertyPtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeAssociationPropertyOutput).ToConnectionTypeAssociationPropertyPtrOutputWithContext(ctx)
}









type ConnectionTypeAssociationPropertyPtrInput interface {
	pulumi.Input

	ToConnectionTypeAssociationPropertyPtrOutput() ConnectionTypeAssociationPropertyPtrOutput
	ToConnectionTypeAssociationPropertyPtrOutputWithContext(context.Context) ConnectionTypeAssociationPropertyPtrOutput
}

type connectionTypeAssociationPropertyPtrType ConnectionTypeAssociationPropertyArgs

func ConnectionTypeAssociationPropertyPtr(v *ConnectionTypeAssociationPropertyArgs) ConnectionTypeAssociationPropertyPtrInput {
	return (*connectionTypeAssociationPropertyPtrType)(v)
}

func (*connectionTypeAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionTypeAssociationProperty)(nil)).Elem()
}

func (i *connectionTypeAssociationPropertyPtrType) ToConnectionTypeAssociationPropertyPtrOutput() ConnectionTypeAssociationPropertyPtrOutput {
	return i.ToConnectionTypeAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *connectionTypeAssociationPropertyPtrType) ToConnectionTypeAssociationPropertyPtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeAssociationPropertyPtrOutput)
}

type ConnectionTypeAssociationPropertyOutput struct{ *pulumi.OutputState }

func (ConnectionTypeAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionTypeAssociationProperty)(nil)).Elem()
}

func (o ConnectionTypeAssociationPropertyOutput) ToConnectionTypeAssociationPropertyOutput() ConnectionTypeAssociationPropertyOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyOutput) ToConnectionTypeAssociationPropertyOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyOutput) ToConnectionTypeAssociationPropertyPtrOutput() ConnectionTypeAssociationPropertyPtrOutput {
	return o.ToConnectionTypeAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o ConnectionTypeAssociationPropertyOutput) ToConnectionTypeAssociationPropertyPtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionTypeAssociationProperty) *ConnectionTypeAssociationProperty {
		return &v
	}).(ConnectionTypeAssociationPropertyPtrOutput)
}

func (o ConnectionTypeAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionTypeAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ConnectionTypeAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (ConnectionTypeAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionTypeAssociationProperty)(nil)).Elem()
}

func (o ConnectionTypeAssociationPropertyPtrOutput) ToConnectionTypeAssociationPropertyPtrOutput() ConnectionTypeAssociationPropertyPtrOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyPtrOutput) ToConnectionTypeAssociationPropertyPtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyPtrOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyPtrOutput) Elem() ConnectionTypeAssociationPropertyOutput {
	return o.ApplyT(func(v *ConnectionTypeAssociationProperty) ConnectionTypeAssociationProperty {
		if v != nil {
			return *v
		}
		var ret ConnectionTypeAssociationProperty
		return ret
	}).(ConnectionTypeAssociationPropertyOutput)
}

func (o ConnectionTypeAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionTypeAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ConnectionTypeAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}





type ConnectionTypeAssociationPropertyResponseInput interface {
	pulumi.Input

	ToConnectionTypeAssociationPropertyResponseOutput() ConnectionTypeAssociationPropertyResponseOutput
	ToConnectionTypeAssociationPropertyResponseOutputWithContext(context.Context) ConnectionTypeAssociationPropertyResponseOutput
}

type ConnectionTypeAssociationPropertyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ConnectionTypeAssociationPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionTypeAssociationPropertyResponse)(nil)).Elem()
}

func (i ConnectionTypeAssociationPropertyResponseArgs) ToConnectionTypeAssociationPropertyResponseOutput() ConnectionTypeAssociationPropertyResponseOutput {
	return i.ToConnectionTypeAssociationPropertyResponseOutputWithContext(context.Background())
}

func (i ConnectionTypeAssociationPropertyResponseArgs) ToConnectionTypeAssociationPropertyResponseOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeAssociationPropertyResponseOutput)
}

func (i ConnectionTypeAssociationPropertyResponseArgs) ToConnectionTypeAssociationPropertyResponsePtrOutput() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return i.ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionTypeAssociationPropertyResponseArgs) ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeAssociationPropertyResponseOutput).ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(ctx)
}









type ConnectionTypeAssociationPropertyResponsePtrInput interface {
	pulumi.Input

	ToConnectionTypeAssociationPropertyResponsePtrOutput() ConnectionTypeAssociationPropertyResponsePtrOutput
	ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(context.Context) ConnectionTypeAssociationPropertyResponsePtrOutput
}

type connectionTypeAssociationPropertyResponsePtrType ConnectionTypeAssociationPropertyResponseArgs

func ConnectionTypeAssociationPropertyResponsePtr(v *ConnectionTypeAssociationPropertyResponseArgs) ConnectionTypeAssociationPropertyResponsePtrInput {
	return (*connectionTypeAssociationPropertyResponsePtrType)(v)
}

func (*connectionTypeAssociationPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionTypeAssociationPropertyResponse)(nil)).Elem()
}

func (i *connectionTypeAssociationPropertyResponsePtrType) ToConnectionTypeAssociationPropertyResponsePtrOutput() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return i.ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *connectionTypeAssociationPropertyResponsePtrType) ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionTypeAssociationPropertyResponsePtrOutput)
}

type ConnectionTypeAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (ConnectionTypeAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionTypeAssociationPropertyResponse)(nil)).Elem()
}

func (o ConnectionTypeAssociationPropertyResponseOutput) ToConnectionTypeAssociationPropertyResponseOutput() ConnectionTypeAssociationPropertyResponseOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyResponseOutput) ToConnectionTypeAssociationPropertyResponseOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyResponseOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyResponseOutput) ToConnectionTypeAssociationPropertyResponsePtrOutput() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o.ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionTypeAssociationPropertyResponseOutput) ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionTypeAssociationPropertyResponse) *ConnectionTypeAssociationPropertyResponse {
		return &v
	}).(ConnectionTypeAssociationPropertyResponsePtrOutput)
}

func (o ConnectionTypeAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionTypeAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ConnectionTypeAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionTypeAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionTypeAssociationPropertyResponse)(nil)).Elem()
}

func (o ConnectionTypeAssociationPropertyResponsePtrOutput) ToConnectionTypeAssociationPropertyResponsePtrOutput() ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyResponsePtrOutput) ToConnectionTypeAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ConnectionTypeAssociationPropertyResponsePtrOutput {
	return o
}

func (o ConnectionTypeAssociationPropertyResponsePtrOutput) Elem() ConnectionTypeAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *ConnectionTypeAssociationPropertyResponse) ConnectionTypeAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionTypeAssociationPropertyResponse
		return ret
	}).(ConnectionTypeAssociationPropertyResponseOutput)
}

func (o ConnectionTypeAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionTypeAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ContentHash struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}





type ContentHashInput interface {
	pulumi.Input

	ToContentHashOutput() ContentHashOutput
	ToContentHashOutputWithContext(context.Context) ContentHashOutput
}

type ContentHashArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (ContentHashArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (i ContentHashArgs) ToContentHashOutput() ContentHashOutput {
	return i.ToContentHashOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput)
}

func (i ContentHashArgs) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i ContentHashArgs) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashOutput).ToContentHashPtrOutputWithContext(ctx)
}









type ContentHashPtrInput interface {
	pulumi.Input

	ToContentHashPtrOutput() ContentHashPtrOutput
	ToContentHashPtrOutputWithContext(context.Context) ContentHashPtrOutput
}

type contentHashPtrType ContentHashArgs

func ContentHashPtr(v *ContentHashArgs) ContentHashPtrInput {
	return (*contentHashPtrType)(v)
}

func (*contentHashPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (i *contentHashPtrType) ToContentHashPtrOutput() ContentHashPtrOutput {
	return i.ToContentHashPtrOutputWithContext(context.Background())
}

func (i *contentHashPtrType) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashPtrOutput)
}

type ContentHashOutput struct{ *pulumi.OutputState }

func (ContentHashOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHash)(nil)).Elem()
}

func (o ContentHashOutput) ToContentHashOutput() ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashOutputWithContext(ctx context.Context) ContentHashOutput {
	return o
}

func (o ContentHashOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o.ToContentHashPtrOutputWithContext(context.Background())
}

func (o ContentHashOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentHash) *ContentHash {
		return &v
	}).(ContentHashPtrOutput)
}

func (o ContentHashOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ContentHashOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHash) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashPtrOutput struct{ *pulumi.OutputState }

func (ContentHashPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHash)(nil)).Elem()
}

func (o ContentHashPtrOutput) ToContentHashPtrOutput() ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) ToContentHashPtrOutputWithContext(ctx context.Context) ContentHashPtrOutput {
	return o
}

func (o ContentHashPtrOutput) Elem() ContentHashOutput {
	return o.ApplyT(func(v *ContentHash) ContentHash {
		if v != nil {
			return *v
		}
		var ret ContentHash
		return ret
	}).(ContentHashOutput)
}

func (o ContentHashPtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashPtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHash) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentHashResponse struct {
	Algorithm string `pulumi:"algorithm"`
	Value     string `pulumi:"value"`
}





type ContentHashResponseInput interface {
	pulumi.Input

	ToContentHashResponseOutput() ContentHashResponseOutput
	ToContentHashResponseOutputWithContext(context.Context) ContentHashResponseOutput
}

type ContentHashResponseArgs struct {
	Algorithm pulumi.StringInput `pulumi:"algorithm"`
	Value     pulumi.StringInput `pulumi:"value"`
}

func (ContentHashResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (i ContentHashResponseArgs) ToContentHashResponseOutput() ContentHashResponseOutput {
	return i.ToContentHashResponseOutputWithContext(context.Background())
}

func (i ContentHashResponseArgs) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashResponseOutput)
}

func (i ContentHashResponseArgs) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return i.ToContentHashResponsePtrOutputWithContext(context.Background())
}

func (i ContentHashResponseArgs) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashResponseOutput).ToContentHashResponsePtrOutputWithContext(ctx)
}









type ContentHashResponsePtrInput interface {
	pulumi.Input

	ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput
	ToContentHashResponsePtrOutputWithContext(context.Context) ContentHashResponsePtrOutput
}

type contentHashResponsePtrType ContentHashResponseArgs

func ContentHashResponsePtr(v *ContentHashResponseArgs) ContentHashResponsePtrInput {
	return (*contentHashResponsePtrType)(v)
}

func (*contentHashResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (i *contentHashResponsePtrType) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return i.ToContentHashResponsePtrOutputWithContext(context.Background())
}

func (i *contentHashResponsePtrType) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentHashResponsePtrOutput)
}

type ContentHashResponseOutput struct{ *pulumi.OutputState }

func (ContentHashResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponseOutput) ToContentHashResponseOutput() ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponseOutputWithContext(ctx context.Context) ContentHashResponseOutput {
	return o
}

func (o ContentHashResponseOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o.ToContentHashResponsePtrOutputWithContext(context.Background())
}

func (o ContentHashResponseOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentHashResponse) *ContentHashResponse {
		return &v
	}).(ContentHashResponsePtrOutput)
}

func (o ContentHashResponseOutput) Algorithm() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHashResponse) string { return v.Algorithm }).(pulumi.StringOutput)
}

func (o ContentHashResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContentHashResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContentHashResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentHashResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentHashResponse)(nil)).Elem()
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutput() ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) ToContentHashResponsePtrOutputWithContext(ctx context.Context) ContentHashResponsePtrOutput {
	return o
}

func (o ContentHashResponsePtrOutput) Elem() ContentHashResponseOutput {
	return o.ApplyT(func(v *ContentHashResponse) ContentHashResponse {
		if v != nil {
			return *v
		}
		var ret ContentHashResponse
		return ret
	}).(ContentHashResponseOutput)
}

func (o ContentHashResponsePtrOutput) Algorithm() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Algorithm
	}).(pulumi.StringPtrOutput)
}

func (o ContentHashResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentHashResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Value
	}).(pulumi.StringPtrOutput)
}

type ContentLink struct {
	ContentHash *ContentHash `pulumi:"contentHash"`
	Uri         *string      `pulumi:"uri"`
	Version     *string      `pulumi:"version"`
}





type ContentLinkInput interface {
	pulumi.Input

	ToContentLinkOutput() ContentLinkOutput
	ToContentLinkOutputWithContext(context.Context) ContentLinkOutput
}

type ContentLinkArgs struct {
	ContentHash ContentHashPtrInput   `pulumi:"contentHash"`
	Uri         pulumi.StringPtrInput `pulumi:"uri"`
	Version     pulumi.StringPtrInput `pulumi:"version"`
}

func (ContentLinkArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (i ContentLinkArgs) ToContentLinkOutput() ContentLinkOutput {
	return i.ToContentLinkOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput)
}

func (i ContentLinkArgs) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i ContentLinkArgs) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkOutput).ToContentLinkPtrOutputWithContext(ctx)
}









type ContentLinkPtrInput interface {
	pulumi.Input

	ToContentLinkPtrOutput() ContentLinkPtrOutput
	ToContentLinkPtrOutputWithContext(context.Context) ContentLinkPtrOutput
}

type contentLinkPtrType ContentLinkArgs

func ContentLinkPtr(v *ContentLinkArgs) ContentLinkPtrInput {
	return (*contentLinkPtrType)(v)
}

func (*contentLinkPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (i *contentLinkPtrType) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return i.ToContentLinkPtrOutputWithContext(context.Background())
}

func (i *contentLinkPtrType) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkPtrOutput)
}

type ContentLinkOutput struct{ *pulumi.OutputState }

func (ContentLinkOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLink)(nil)).Elem()
}

func (o ContentLinkOutput) ToContentLinkOutput() ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkOutputWithContext(ctx context.Context) ContentLinkOutput {
	return o
}

func (o ContentLinkOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o.ToContentLinkPtrOutputWithContext(context.Background())
}

func (o ContentLinkOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLink) *ContentLink {
		return &v
	}).(ContentLinkPtrOutput)
}

func (o ContentLinkOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentLink) *ContentHash { return v.ContentHash }).(ContentHashPtrOutput)
}

func (o ContentLinkOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func (o ContentLinkOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLink) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentLinkPtrOutput struct{ *pulumi.OutputState }

func (ContentLinkPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLink)(nil)).Elem()
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutput() ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) ToContentLinkPtrOutputWithContext(ctx context.Context) ContentLinkPtrOutput {
	return o
}

func (o ContentLinkPtrOutput) Elem() ContentLinkOutput {
	return o.ApplyT(func(v *ContentLink) ContentLink {
		if v != nil {
			return *v
		}
		var ret ContentLink
		return ret
	}).(ContentLinkOutput)
}

func (o ContentLinkPtrOutput) ContentHash() ContentHashPtrOutput {
	return o.ApplyT(func(v *ContentLink) *ContentHash {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashPtrOutput)
}

func (o ContentLinkPtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLink) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ContentLinkResponse struct {
	ContentHash *ContentHashResponse `pulumi:"contentHash"`
	Uri         *string              `pulumi:"uri"`
	Version     *string              `pulumi:"version"`
}





type ContentLinkResponseInput interface {
	pulumi.Input

	ToContentLinkResponseOutput() ContentLinkResponseOutput
	ToContentLinkResponseOutputWithContext(context.Context) ContentLinkResponseOutput
}

type ContentLinkResponseArgs struct {
	ContentHash ContentHashResponsePtrInput `pulumi:"contentHash"`
	Uri         pulumi.StringPtrInput       `pulumi:"uri"`
	Version     pulumi.StringPtrInput       `pulumi:"version"`
}

func (ContentLinkResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (i ContentLinkResponseArgs) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return i.ToContentLinkResponseOutputWithContext(context.Background())
}

func (i ContentLinkResponseArgs) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkResponseOutput)
}

func (i ContentLinkResponseArgs) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return i.ToContentLinkResponsePtrOutputWithContext(context.Background())
}

func (i ContentLinkResponseArgs) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkResponseOutput).ToContentLinkResponsePtrOutputWithContext(ctx)
}









type ContentLinkResponsePtrInput interface {
	pulumi.Input

	ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput
	ToContentLinkResponsePtrOutputWithContext(context.Context) ContentLinkResponsePtrOutput
}

type contentLinkResponsePtrType ContentLinkResponseArgs

func ContentLinkResponsePtr(v *ContentLinkResponseArgs) ContentLinkResponsePtrInput {
	return (*contentLinkResponsePtrType)(v)
}

func (*contentLinkResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (i *contentLinkResponsePtrType) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return i.ToContentLinkResponsePtrOutputWithContext(context.Background())
}

func (i *contentLinkResponsePtrType) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentLinkResponsePtrOutput)
}

type ContentLinkResponseOutput struct{ *pulumi.OutputState }

func (ContentLinkResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutput() ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponseOutputWithContext(ctx context.Context) ContentLinkResponseOutput {
	return o
}

func (o ContentLinkResponseOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o.ToContentLinkResponsePtrOutputWithContext(context.Background())
}

func (o ContentLinkResponseOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentLinkResponse) *ContentLinkResponse {
		return &v
	}).(ContentLinkResponsePtrOutput)
}

func (o ContentLinkResponseOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *ContentHashResponse { return v.ContentHash }).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponseOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Uri }).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentLinkResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentLinkResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentLinkResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentLinkResponse)(nil)).Elem()
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutput() ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) ToContentLinkResponsePtrOutputWithContext(ctx context.Context) ContentLinkResponsePtrOutput {
	return o
}

func (o ContentLinkResponsePtrOutput) Elem() ContentLinkResponseOutput {
	return o.ApplyT(func(v *ContentLinkResponse) ContentLinkResponse {
		if v != nil {
			return *v
		}
		var ret ContentLinkResponse
		return ret
	}).(ContentLinkResponseOutput)
}

func (o ContentLinkResponsePtrOutput) ContentHash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *ContentHashResponse {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(ContentHashResponsePtrOutput)
}

func (o ContentLinkResponsePtrOutput) Uri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Uri
	}).(pulumi.StringPtrOutput)
}

func (o ContentLinkResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentLinkResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ContentSource struct {
	Hash    *ContentHash `pulumi:"hash"`
	Type    *string      `pulumi:"type"`
	Value   *string      `pulumi:"value"`
	Version *string      `pulumi:"version"`
}





type ContentSourceInput interface {
	pulumi.Input

	ToContentSourceOutput() ContentSourceOutput
	ToContentSourceOutputWithContext(context.Context) ContentSourceOutput
}

type ContentSourceArgs struct {
	Hash    ContentHashPtrInput   `pulumi:"hash"`
	Type    pulumi.StringPtrInput `pulumi:"type"`
	Value   pulumi.StringPtrInput `pulumi:"value"`
	Version pulumi.StringPtrInput `pulumi:"version"`
}

func (ContentSourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSource)(nil)).Elem()
}

func (i ContentSourceArgs) ToContentSourceOutput() ContentSourceOutput {
	return i.ToContentSourceOutputWithContext(context.Background())
}

func (i ContentSourceArgs) ToContentSourceOutputWithContext(ctx context.Context) ContentSourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceOutput)
}

func (i ContentSourceArgs) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return i.ToContentSourcePtrOutputWithContext(context.Background())
}

func (i ContentSourceArgs) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceOutput).ToContentSourcePtrOutputWithContext(ctx)
}









type ContentSourcePtrInput interface {
	pulumi.Input

	ToContentSourcePtrOutput() ContentSourcePtrOutput
	ToContentSourcePtrOutputWithContext(context.Context) ContentSourcePtrOutput
}

type contentSourcePtrType ContentSourceArgs

func ContentSourcePtr(v *ContentSourceArgs) ContentSourcePtrInput {
	return (*contentSourcePtrType)(v)
}

func (*contentSourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSource)(nil)).Elem()
}

func (i *contentSourcePtrType) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return i.ToContentSourcePtrOutputWithContext(context.Background())
}

func (i *contentSourcePtrType) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourcePtrOutput)
}

type ContentSourceOutput struct{ *pulumi.OutputState }

func (ContentSourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSource)(nil)).Elem()
}

func (o ContentSourceOutput) ToContentSourceOutput() ContentSourceOutput {
	return o
}

func (o ContentSourceOutput) ToContentSourceOutputWithContext(ctx context.Context) ContentSourceOutput {
	return o
}

func (o ContentSourceOutput) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return o.ToContentSourcePtrOutputWithContext(context.Background())
}

func (o ContentSourceOutput) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentSource) *ContentSource {
		return &v
	}).(ContentSourcePtrOutput)
}

func (o ContentSourceOutput) Hash() ContentHashPtrOutput {
	return o.ApplyT(func(v ContentSource) *ContentHash { return v.Hash }).(ContentHashPtrOutput)
}

func (o ContentSourceOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ContentSourceOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o ContentSourceOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSource) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentSourcePtrOutput struct{ *pulumi.OutputState }

func (ContentSourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSource)(nil)).Elem()
}

func (o ContentSourcePtrOutput) ToContentSourcePtrOutput() ContentSourcePtrOutput {
	return o
}

func (o ContentSourcePtrOutput) ToContentSourcePtrOutputWithContext(ctx context.Context) ContentSourcePtrOutput {
	return o
}

func (o ContentSourcePtrOutput) Elem() ContentSourceOutput {
	return o.ApplyT(func(v *ContentSource) ContentSource {
		if v != nil {
			return *v
		}
		var ret ContentSource
		return ret
	}).(ContentSourceOutput)
}

func (o ContentSourcePtrOutput) Hash() ContentHashPtrOutput {
	return o.ApplyT(func(v *ContentSource) *ContentHash {
		if v == nil {
			return nil
		}
		return v.Hash
	}).(ContentHashPtrOutput)
}

func (o ContentSourcePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSource) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ContentSourcePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSource) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

func (o ContentSourcePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSource) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type ContentSourceResponse struct {
	Hash    *ContentHashResponse `pulumi:"hash"`
	Type    *string              `pulumi:"type"`
	Value   *string              `pulumi:"value"`
	Version *string              `pulumi:"version"`
}





type ContentSourceResponseInput interface {
	pulumi.Input

	ToContentSourceResponseOutput() ContentSourceResponseOutput
	ToContentSourceResponseOutputWithContext(context.Context) ContentSourceResponseOutput
}

type ContentSourceResponseArgs struct {
	Hash    ContentHashResponsePtrInput `pulumi:"hash"`
	Type    pulumi.StringPtrInput       `pulumi:"type"`
	Value   pulumi.StringPtrInput       `pulumi:"value"`
	Version pulumi.StringPtrInput       `pulumi:"version"`
}

func (ContentSourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSourceResponse)(nil)).Elem()
}

func (i ContentSourceResponseArgs) ToContentSourceResponseOutput() ContentSourceResponseOutput {
	return i.ToContentSourceResponseOutputWithContext(context.Background())
}

func (i ContentSourceResponseArgs) ToContentSourceResponseOutputWithContext(ctx context.Context) ContentSourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceResponseOutput)
}

func (i ContentSourceResponseArgs) ToContentSourceResponsePtrOutput() ContentSourceResponsePtrOutput {
	return i.ToContentSourceResponsePtrOutputWithContext(context.Background())
}

func (i ContentSourceResponseArgs) ToContentSourceResponsePtrOutputWithContext(ctx context.Context) ContentSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceResponseOutput).ToContentSourceResponsePtrOutputWithContext(ctx)
}









type ContentSourceResponsePtrInput interface {
	pulumi.Input

	ToContentSourceResponsePtrOutput() ContentSourceResponsePtrOutput
	ToContentSourceResponsePtrOutputWithContext(context.Context) ContentSourceResponsePtrOutput
}

type contentSourceResponsePtrType ContentSourceResponseArgs

func ContentSourceResponsePtr(v *ContentSourceResponseArgs) ContentSourceResponsePtrInput {
	return (*contentSourceResponsePtrType)(v)
}

func (*contentSourceResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSourceResponse)(nil)).Elem()
}

func (i *contentSourceResponsePtrType) ToContentSourceResponsePtrOutput() ContentSourceResponsePtrOutput {
	return i.ToContentSourceResponsePtrOutputWithContext(context.Background())
}

func (i *contentSourceResponsePtrType) ToContentSourceResponsePtrOutputWithContext(ctx context.Context) ContentSourceResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContentSourceResponsePtrOutput)
}

type ContentSourceResponseOutput struct{ *pulumi.OutputState }

func (ContentSourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContentSourceResponse)(nil)).Elem()
}

func (o ContentSourceResponseOutput) ToContentSourceResponseOutput() ContentSourceResponseOutput {
	return o
}

func (o ContentSourceResponseOutput) ToContentSourceResponseOutputWithContext(ctx context.Context) ContentSourceResponseOutput {
	return o
}

func (o ContentSourceResponseOutput) ToContentSourceResponsePtrOutput() ContentSourceResponsePtrOutput {
	return o.ToContentSourceResponsePtrOutputWithContext(context.Background())
}

func (o ContentSourceResponseOutput) ToContentSourceResponsePtrOutputWithContext(ctx context.Context) ContentSourceResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContentSourceResponse) *ContentSourceResponse {
		return &v
	}).(ContentSourceResponsePtrOutput)
}

func (o ContentSourceResponseOutput) Hash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v ContentSourceResponse) *ContentHashResponse { return v.Hash }).(ContentHashResponsePtrOutput)
}

func (o ContentSourceResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSourceResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ContentSourceResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSourceResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

func (o ContentSourceResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContentSourceResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ContentSourceResponsePtrOutput struct{ *pulumi.OutputState }

func (ContentSourceResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContentSourceResponse)(nil)).Elem()
}

func (o ContentSourceResponsePtrOutput) ToContentSourceResponsePtrOutput() ContentSourceResponsePtrOutput {
	return o
}

func (o ContentSourceResponsePtrOutput) ToContentSourceResponsePtrOutputWithContext(ctx context.Context) ContentSourceResponsePtrOutput {
	return o
}

func (o ContentSourceResponsePtrOutput) Elem() ContentSourceResponseOutput {
	return o.ApplyT(func(v *ContentSourceResponse) ContentSourceResponse {
		if v != nil {
			return *v
		}
		var ret ContentSourceResponse
		return ret
	}).(ContentSourceResponseOutput)
}

func (o ContentSourceResponsePtrOutput) Hash() ContentHashResponsePtrOutput {
	return o.ApplyT(func(v *ContentSourceResponse) *ContentHashResponse {
		if v == nil {
			return nil
		}
		return v.Hash
	}).(ContentHashResponsePtrOutput)
}

func (o ContentSourceResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ContentSourceResponsePtrOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Value
	}).(pulumi.StringPtrOutput)
}

func (o ContentSourceResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContentSourceResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type DscConfigurationAssociationPropertyInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput
	ToDscConfigurationAssociationPropertyOutputWithContext(context.Context) DscConfigurationAssociationPropertyOutput
}

type DscConfigurationAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DscConfigurationAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationProperty)(nil)).Elem()
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput {
	return i.ToDscConfigurationAssociationPropertyOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyOutput)
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return i.ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyArgs) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyOutput).ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx)
}









type DscConfigurationAssociationPropertyPtrInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput
	ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Context) DscConfigurationAssociationPropertyPtrOutput
}

type dscConfigurationAssociationPropertyPtrType DscConfigurationAssociationPropertyArgs

func DscConfigurationAssociationPropertyPtr(v *DscConfigurationAssociationPropertyArgs) DscConfigurationAssociationPropertyPtrInput {
	return (*dscConfigurationAssociationPropertyPtrType)(v)
}

func (*dscConfigurationAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationProperty)(nil)).Elem()
}

func (i *dscConfigurationAssociationPropertyPtrType) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return i.ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *dscConfigurationAssociationPropertyPtrType) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyPtrOutput)
}

type DscConfigurationAssociationPropertyOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationProperty)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyOutput() DscConfigurationAssociationPropertyOutput {
	return o
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyOutput {
	return o
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return o.ToDscConfigurationAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o DscConfigurationAssociationPropertyOutput) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DscConfigurationAssociationProperty) *DscConfigurationAssociationProperty {
		return &v
	}).(DscConfigurationAssociationPropertyPtrOutput)
}

func (o DscConfigurationAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationProperty)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyPtrOutput) ToDscConfigurationAssociationPropertyPtrOutput() DscConfigurationAssociationPropertyPtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyPtrOutput) ToDscConfigurationAssociationPropertyPtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyPtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyPtrOutput) Elem() DscConfigurationAssociationPropertyOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationProperty) DscConfigurationAssociationProperty {
		if v != nil {
			return *v
		}
		var ret DscConfigurationAssociationProperty
		return ret
	}).(DscConfigurationAssociationPropertyOutput)
}

func (o DscConfigurationAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}





type DscConfigurationAssociationPropertyResponseInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput
	ToDscConfigurationAssociationPropertyResponseOutputWithContext(context.Context) DscConfigurationAssociationPropertyResponseOutput
}

type DscConfigurationAssociationPropertyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (DscConfigurationAssociationPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput {
	return i.ToDscConfigurationAssociationPropertyResponseOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponseOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyResponseOutput)
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return i.ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i DscConfigurationAssociationPropertyResponseArgs) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyResponseOutput).ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx)
}









type DscConfigurationAssociationPropertyResponsePtrInput interface {
	pulumi.Input

	ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput
	ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Context) DscConfigurationAssociationPropertyResponsePtrOutput
}

type dscConfigurationAssociationPropertyResponsePtrType DscConfigurationAssociationPropertyResponseArgs

func DscConfigurationAssociationPropertyResponsePtr(v *DscConfigurationAssociationPropertyResponseArgs) DscConfigurationAssociationPropertyResponsePtrInput {
	return (*dscConfigurationAssociationPropertyResponsePtrType)(v)
}

func (*dscConfigurationAssociationPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (i *dscConfigurationAssociationPropertyResponsePtrType) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return i.ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *dscConfigurationAssociationPropertyResponsePtrType) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

type DscConfigurationAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponseOutput() DscConfigurationAssociationPropertyResponseOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponseOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponseOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (o DscConfigurationAssociationPropertyResponseOutput) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DscConfigurationAssociationPropertyResponse) *DscConfigurationAssociationPropertyResponse {
		return &v
	}).(DscConfigurationAssociationPropertyResponsePtrOutput)
}

func (o DscConfigurationAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (DscConfigurationAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DscConfigurationAssociationPropertyResponse)(nil)).Elem()
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) ToDscConfigurationAssociationPropertyResponsePtrOutput() DscConfigurationAssociationPropertyResponsePtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) ToDscConfigurationAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) DscConfigurationAssociationPropertyResponsePtrOutput {
	return o
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) Elem() DscConfigurationAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationPropertyResponse) DscConfigurationAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret DscConfigurationAssociationPropertyResponse
		return ret
	}).(DscConfigurationAssociationPropertyResponseOutput)
}

func (o DscConfigurationAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DscConfigurationAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type DscConfigurationParameter struct {
	DefaultValue *string `pulumi:"defaultValue"`
	IsMandatory  *bool   `pulumi:"isMandatory"`
	Position     *int    `pulumi:"position"`
	Type         *string `pulumi:"type"`
}





type DscConfigurationParameterInput interface {
	pulumi.Input

	ToDscConfigurationParameterOutput() DscConfigurationParameterOutput
	ToDscConfigurationParameterOutputWithContext(context.Context) DscConfigurationParameterOutput
}

type DscConfigurationParameterArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	IsMandatory  pulumi.BoolPtrInput   `pulumi:"isMandatory"`
	Position     pulumi.IntPtrInput    `pulumi:"position"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (DscConfigurationParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationParameter)(nil)).Elem()
}

func (i DscConfigurationParameterArgs) ToDscConfigurationParameterOutput() DscConfigurationParameterOutput {
	return i.ToDscConfigurationParameterOutputWithContext(context.Background())
}

func (i DscConfigurationParameterArgs) ToDscConfigurationParameterOutputWithContext(ctx context.Context) DscConfigurationParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationParameterOutput)
}





type DscConfigurationParameterMapInput interface {
	pulumi.Input

	ToDscConfigurationParameterMapOutput() DscConfigurationParameterMapOutput
	ToDscConfigurationParameterMapOutputWithContext(context.Context) DscConfigurationParameterMapOutput
}

type DscConfigurationParameterMap map[string]DscConfigurationParameterInput

func (DscConfigurationParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DscConfigurationParameter)(nil)).Elem()
}

func (i DscConfigurationParameterMap) ToDscConfigurationParameterMapOutput() DscConfigurationParameterMapOutput {
	return i.ToDscConfigurationParameterMapOutputWithContext(context.Background())
}

func (i DscConfigurationParameterMap) ToDscConfigurationParameterMapOutputWithContext(ctx context.Context) DscConfigurationParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationParameterMapOutput)
}

type DscConfigurationParameterOutput struct{ *pulumi.OutputState }

func (DscConfigurationParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationParameter)(nil)).Elem()
}

func (o DscConfigurationParameterOutput) ToDscConfigurationParameterOutput() DscConfigurationParameterOutput {
	return o
}

func (o DscConfigurationParameterOutput) ToDscConfigurationParameterOutputWithContext(ctx context.Context) DscConfigurationParameterOutput {
	return o
}

func (o DscConfigurationParameterOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameter) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationParameterOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameter) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o DscConfigurationParameterOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameter) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func (o DscConfigurationParameterOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameter) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DscConfigurationParameterMapOutput struct{ *pulumi.OutputState }

func (DscConfigurationParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DscConfigurationParameter)(nil)).Elem()
}

func (o DscConfigurationParameterMapOutput) ToDscConfigurationParameterMapOutput() DscConfigurationParameterMapOutput {
	return o
}

func (o DscConfigurationParameterMapOutput) ToDscConfigurationParameterMapOutputWithContext(ctx context.Context) DscConfigurationParameterMapOutput {
	return o
}

func (o DscConfigurationParameterMapOutput) MapIndex(k pulumi.StringInput) DscConfigurationParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DscConfigurationParameter {
		return vs[0].(map[string]DscConfigurationParameter)[vs[1].(string)]
	}).(DscConfigurationParameterOutput)
}

type DscConfigurationParameterResponse struct {
	DefaultValue *string `pulumi:"defaultValue"`
	IsMandatory  *bool   `pulumi:"isMandatory"`
	Position     *int    `pulumi:"position"`
	Type         *string `pulumi:"type"`
}





type DscConfigurationParameterResponseInput interface {
	pulumi.Input

	ToDscConfigurationParameterResponseOutput() DscConfigurationParameterResponseOutput
	ToDscConfigurationParameterResponseOutputWithContext(context.Context) DscConfigurationParameterResponseOutput
}

type DscConfigurationParameterResponseArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	IsMandatory  pulumi.BoolPtrInput   `pulumi:"isMandatory"`
	Position     pulumi.IntPtrInput    `pulumi:"position"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (DscConfigurationParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationParameterResponse)(nil)).Elem()
}

func (i DscConfigurationParameterResponseArgs) ToDscConfigurationParameterResponseOutput() DscConfigurationParameterResponseOutput {
	return i.ToDscConfigurationParameterResponseOutputWithContext(context.Background())
}

func (i DscConfigurationParameterResponseArgs) ToDscConfigurationParameterResponseOutputWithContext(ctx context.Context) DscConfigurationParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationParameterResponseOutput)
}





type DscConfigurationParameterResponseMapInput interface {
	pulumi.Input

	ToDscConfigurationParameterResponseMapOutput() DscConfigurationParameterResponseMapOutput
	ToDscConfigurationParameterResponseMapOutputWithContext(context.Context) DscConfigurationParameterResponseMapOutput
}

type DscConfigurationParameterResponseMap map[string]DscConfigurationParameterResponseInput

func (DscConfigurationParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DscConfigurationParameterResponse)(nil)).Elem()
}

func (i DscConfigurationParameterResponseMap) ToDscConfigurationParameterResponseMapOutput() DscConfigurationParameterResponseMapOutput {
	return i.ToDscConfigurationParameterResponseMapOutputWithContext(context.Background())
}

func (i DscConfigurationParameterResponseMap) ToDscConfigurationParameterResponseMapOutputWithContext(ctx context.Context) DscConfigurationParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DscConfigurationParameterResponseMapOutput)
}

type DscConfigurationParameterResponseOutput struct{ *pulumi.OutputState }

func (DscConfigurationParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DscConfigurationParameterResponse)(nil)).Elem()
}

func (o DscConfigurationParameterResponseOutput) ToDscConfigurationParameterResponseOutput() DscConfigurationParameterResponseOutput {
	return o
}

func (o DscConfigurationParameterResponseOutput) ToDscConfigurationParameterResponseOutputWithContext(ctx context.Context) DscConfigurationParameterResponseOutput {
	return o
}

func (o DscConfigurationParameterResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameterResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o DscConfigurationParameterResponseOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameterResponse) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o DscConfigurationParameterResponseOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameterResponse) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func (o DscConfigurationParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type DscConfigurationParameterResponseMapOutput struct{ *pulumi.OutputState }

func (DscConfigurationParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]DscConfigurationParameterResponse)(nil)).Elem()
}

func (o DscConfigurationParameterResponseMapOutput) ToDscConfigurationParameterResponseMapOutput() DscConfigurationParameterResponseMapOutput {
	return o
}

func (o DscConfigurationParameterResponseMapOutput) ToDscConfigurationParameterResponseMapOutputWithContext(ctx context.Context) DscConfigurationParameterResponseMapOutput {
	return o
}

func (o DscConfigurationParameterResponseMapOutput) MapIndex(k pulumi.StringInput) DscConfigurationParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) DscConfigurationParameterResponse {
		return vs[0].(map[string]DscConfigurationParameterResponse)[vs[1].(string)]
	}).(DscConfigurationParameterResponseOutput)
}

type FieldDefinition struct {
	IsEncrypted *bool  `pulumi:"isEncrypted"`
	IsOptional  *bool  `pulumi:"isOptional"`
	Type        string `pulumi:"type"`
}





type FieldDefinitionInput interface {
	pulumi.Input

	ToFieldDefinitionOutput() FieldDefinitionOutput
	ToFieldDefinitionOutputWithContext(context.Context) FieldDefinitionOutput
}

type FieldDefinitionArgs struct {
	IsEncrypted pulumi.BoolPtrInput `pulumi:"isEncrypted"`
	IsOptional  pulumi.BoolPtrInput `pulumi:"isOptional"`
	Type        pulumi.StringInput  `pulumi:"type"`
}

func (FieldDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldDefinition)(nil)).Elem()
}

func (i FieldDefinitionArgs) ToFieldDefinitionOutput() FieldDefinitionOutput {
	return i.ToFieldDefinitionOutputWithContext(context.Background())
}

func (i FieldDefinitionArgs) ToFieldDefinitionOutputWithContext(ctx context.Context) FieldDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldDefinitionOutput)
}





type FieldDefinitionMapInput interface {
	pulumi.Input

	ToFieldDefinitionMapOutput() FieldDefinitionMapOutput
	ToFieldDefinitionMapOutputWithContext(context.Context) FieldDefinitionMapOutput
}

type FieldDefinitionMap map[string]FieldDefinitionInput

func (FieldDefinitionMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FieldDefinition)(nil)).Elem()
}

func (i FieldDefinitionMap) ToFieldDefinitionMapOutput() FieldDefinitionMapOutput {
	return i.ToFieldDefinitionMapOutputWithContext(context.Background())
}

func (i FieldDefinitionMap) ToFieldDefinitionMapOutputWithContext(ctx context.Context) FieldDefinitionMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldDefinitionMapOutput)
}

type FieldDefinitionOutput struct{ *pulumi.OutputState }

func (FieldDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldDefinition)(nil)).Elem()
}

func (o FieldDefinitionOutput) ToFieldDefinitionOutput() FieldDefinitionOutput {
	return o
}

func (o FieldDefinitionOutput) ToFieldDefinitionOutputWithContext(ctx context.Context) FieldDefinitionOutput {
	return o
}

func (o FieldDefinitionOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FieldDefinition) *bool { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

func (o FieldDefinitionOutput) IsOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FieldDefinition) *bool { return v.IsOptional }).(pulumi.BoolPtrOutput)
}

func (o FieldDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FieldDefinition) string { return v.Type }).(pulumi.StringOutput)
}

type FieldDefinitionMapOutput struct{ *pulumi.OutputState }

func (FieldDefinitionMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FieldDefinition)(nil)).Elem()
}

func (o FieldDefinitionMapOutput) ToFieldDefinitionMapOutput() FieldDefinitionMapOutput {
	return o
}

func (o FieldDefinitionMapOutput) ToFieldDefinitionMapOutputWithContext(ctx context.Context) FieldDefinitionMapOutput {
	return o
}

func (o FieldDefinitionMapOutput) MapIndex(k pulumi.StringInput) FieldDefinitionOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FieldDefinition {
		return vs[0].(map[string]FieldDefinition)[vs[1].(string)]
	}).(FieldDefinitionOutput)
}

type FieldDefinitionResponse struct {
	IsEncrypted *bool  `pulumi:"isEncrypted"`
	IsOptional  *bool  `pulumi:"isOptional"`
	Type        string `pulumi:"type"`
}





type FieldDefinitionResponseInput interface {
	pulumi.Input

	ToFieldDefinitionResponseOutput() FieldDefinitionResponseOutput
	ToFieldDefinitionResponseOutputWithContext(context.Context) FieldDefinitionResponseOutput
}

type FieldDefinitionResponseArgs struct {
	IsEncrypted pulumi.BoolPtrInput `pulumi:"isEncrypted"`
	IsOptional  pulumi.BoolPtrInput `pulumi:"isOptional"`
	Type        pulumi.StringInput  `pulumi:"type"`
}

func (FieldDefinitionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldDefinitionResponse)(nil)).Elem()
}

func (i FieldDefinitionResponseArgs) ToFieldDefinitionResponseOutput() FieldDefinitionResponseOutput {
	return i.ToFieldDefinitionResponseOutputWithContext(context.Background())
}

func (i FieldDefinitionResponseArgs) ToFieldDefinitionResponseOutputWithContext(ctx context.Context) FieldDefinitionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldDefinitionResponseOutput)
}





type FieldDefinitionResponseMapInput interface {
	pulumi.Input

	ToFieldDefinitionResponseMapOutput() FieldDefinitionResponseMapOutput
	ToFieldDefinitionResponseMapOutputWithContext(context.Context) FieldDefinitionResponseMapOutput
}

type FieldDefinitionResponseMap map[string]FieldDefinitionResponseInput

func (FieldDefinitionResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FieldDefinitionResponse)(nil)).Elem()
}

func (i FieldDefinitionResponseMap) ToFieldDefinitionResponseMapOutput() FieldDefinitionResponseMapOutput {
	return i.ToFieldDefinitionResponseMapOutputWithContext(context.Background())
}

func (i FieldDefinitionResponseMap) ToFieldDefinitionResponseMapOutputWithContext(ctx context.Context) FieldDefinitionResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FieldDefinitionResponseMapOutput)
}

type FieldDefinitionResponseOutput struct{ *pulumi.OutputState }

func (FieldDefinitionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FieldDefinitionResponse)(nil)).Elem()
}

func (o FieldDefinitionResponseOutput) ToFieldDefinitionResponseOutput() FieldDefinitionResponseOutput {
	return o
}

func (o FieldDefinitionResponseOutput) ToFieldDefinitionResponseOutputWithContext(ctx context.Context) FieldDefinitionResponseOutput {
	return o
}

func (o FieldDefinitionResponseOutput) IsEncrypted() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FieldDefinitionResponse) *bool { return v.IsEncrypted }).(pulumi.BoolPtrOutput)
}

func (o FieldDefinitionResponseOutput) IsOptional() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FieldDefinitionResponse) *bool { return v.IsOptional }).(pulumi.BoolPtrOutput)
}

func (o FieldDefinitionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v FieldDefinitionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type FieldDefinitionResponseMapOutput struct{ *pulumi.OutputState }

func (FieldDefinitionResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]FieldDefinitionResponse)(nil)).Elem()
}

func (o FieldDefinitionResponseMapOutput) ToFieldDefinitionResponseMapOutput() FieldDefinitionResponseMapOutput {
	return o
}

func (o FieldDefinitionResponseMapOutput) ToFieldDefinitionResponseMapOutputWithContext(ctx context.Context) FieldDefinitionResponseMapOutput {
	return o
}

func (o FieldDefinitionResponseMapOutput) MapIndex(k pulumi.StringInput) FieldDefinitionResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) FieldDefinitionResponse {
		return vs[0].(map[string]FieldDefinitionResponse)[vs[1].(string)]
	}).(FieldDefinitionResponseOutput)
}

type KeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}





type KeyResponseInput interface {
	pulumi.Input

	ToKeyResponseOutput() KeyResponseOutput
	ToKeyResponseOutputWithContext(context.Context) KeyResponseOutput
}

type KeyResponseArgs struct {
	KeyName     pulumi.StringInput `pulumi:"keyName"`
	Permissions pulumi.StringInput `pulumi:"permissions"`
	Value       pulumi.StringInput `pulumi:"value"`
}

func (KeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyResponse)(nil)).Elem()
}

func (i KeyResponseArgs) ToKeyResponseOutput() KeyResponseOutput {
	return i.ToKeyResponseOutputWithContext(context.Background())
}

func (i KeyResponseArgs) ToKeyResponseOutputWithContext(ctx context.Context) KeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyResponseOutput)
}





type KeyResponseArrayInput interface {
	pulumi.Input

	ToKeyResponseArrayOutput() KeyResponseArrayOutput
	ToKeyResponseArrayOutputWithContext(context.Context) KeyResponseArrayOutput
}

type KeyResponseArray []KeyResponseInput

func (KeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyResponse)(nil)).Elem()
}

func (i KeyResponseArray) ToKeyResponseArrayOutput() KeyResponseArrayOutput {
	return i.ToKeyResponseArrayOutputWithContext(context.Background())
}

func (i KeyResponseArray) ToKeyResponseArrayOutputWithContext(ctx context.Context) KeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyResponseArrayOutput)
}

type KeyResponseOutput struct{ *pulumi.OutputState }

func (KeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyResponse)(nil)).Elem()
}

func (o KeyResponseOutput) ToKeyResponseOutput() KeyResponseOutput {
	return o
}

func (o KeyResponseOutput) ToKeyResponseOutputWithContext(ctx context.Context) KeyResponseOutput {
	return o
}

func (o KeyResponseOutput) KeyName() pulumi.StringOutput {
	return o.ApplyT(func(v KeyResponse) string { return v.KeyName }).(pulumi.StringOutput)
}

func (o KeyResponseOutput) Permissions() pulumi.StringOutput {
	return o.ApplyT(func(v KeyResponse) string { return v.Permissions }).(pulumi.StringOutput)
}

func (o KeyResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v KeyResponse) string { return v.Value }).(pulumi.StringOutput)
}

type KeyResponseArrayOutput struct{ *pulumi.OutputState }

func (KeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]KeyResponse)(nil)).Elem()
}

func (o KeyResponseArrayOutput) ToKeyResponseArrayOutput() KeyResponseArrayOutput {
	return o
}

func (o KeyResponseArrayOutput) ToKeyResponseArrayOutputWithContext(ctx context.Context) KeyResponseArrayOutput {
	return o
}

func (o KeyResponseArrayOutput) Index(i pulumi.IntInput) KeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) KeyResponse {
		return vs[0].([]KeyResponse)[vs[1].(int)]
	}).(KeyResponseOutput)
}

type ModuleErrorInfoResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type ModuleErrorInfoResponseInput interface {
	pulumi.Input

	ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput
	ToModuleErrorInfoResponseOutputWithContext(context.Context) ModuleErrorInfoResponseOutput
}

type ModuleErrorInfoResponseArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (ModuleErrorInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleErrorInfoResponse)(nil)).Elem()
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput {
	return i.ToModuleErrorInfoResponseOutputWithContext(context.Background())
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponseOutputWithContext(ctx context.Context) ModuleErrorInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleErrorInfoResponseOutput)
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return i.ToModuleErrorInfoResponsePtrOutputWithContext(context.Background())
}

func (i ModuleErrorInfoResponseArgs) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleErrorInfoResponseOutput).ToModuleErrorInfoResponsePtrOutputWithContext(ctx)
}









type ModuleErrorInfoResponsePtrInput interface {
	pulumi.Input

	ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput
	ToModuleErrorInfoResponsePtrOutputWithContext(context.Context) ModuleErrorInfoResponsePtrOutput
}

type moduleErrorInfoResponsePtrType ModuleErrorInfoResponseArgs

func ModuleErrorInfoResponsePtr(v *ModuleErrorInfoResponseArgs) ModuleErrorInfoResponsePtrInput {
	return (*moduleErrorInfoResponsePtrType)(v)
}

func (*moduleErrorInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleErrorInfoResponse)(nil)).Elem()
}

func (i *moduleErrorInfoResponsePtrType) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return i.ToModuleErrorInfoResponsePtrOutputWithContext(context.Background())
}

func (i *moduleErrorInfoResponsePtrType) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ModuleErrorInfoResponsePtrOutput)
}

type ModuleErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (ModuleErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ModuleErrorInfoResponse)(nil)).Elem()
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponseOutput() ModuleErrorInfoResponseOutput {
	return o
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponseOutputWithContext(ctx context.Context) ModuleErrorInfoResponseOutput {
	return o
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return o.ToModuleErrorInfoResponsePtrOutputWithContext(context.Background())
}

func (o ModuleErrorInfoResponseOutput) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ModuleErrorInfoResponse) *ModuleErrorInfoResponse {
		return &v
	}).(ModuleErrorInfoResponsePtrOutput)
}

func (o ModuleErrorInfoResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleErrorInfoResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ModuleErrorInfoResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ModuleErrorInfoResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ModuleErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ModuleErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ModuleErrorInfoResponse)(nil)).Elem()
}

func (o ModuleErrorInfoResponsePtrOutput) ToModuleErrorInfoResponsePtrOutput() ModuleErrorInfoResponsePtrOutput {
	return o
}

func (o ModuleErrorInfoResponsePtrOutput) ToModuleErrorInfoResponsePtrOutputWithContext(ctx context.Context) ModuleErrorInfoResponsePtrOutput {
	return o
}

func (o ModuleErrorInfoResponsePtrOutput) Elem() ModuleErrorInfoResponseOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) ModuleErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret ModuleErrorInfoResponse
		return ret
	}).(ModuleErrorInfoResponseOutput)
}

func (o ModuleErrorInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ModuleErrorInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ModuleErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type RunbookAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type RunbookAssociationPropertyInput interface {
	pulumi.Input

	ToRunbookAssociationPropertyOutput() RunbookAssociationPropertyOutput
	ToRunbookAssociationPropertyOutputWithContext(context.Context) RunbookAssociationPropertyOutput
}

type RunbookAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (RunbookAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookAssociationProperty)(nil)).Elem()
}

func (i RunbookAssociationPropertyArgs) ToRunbookAssociationPropertyOutput() RunbookAssociationPropertyOutput {
	return i.ToRunbookAssociationPropertyOutputWithContext(context.Background())
}

func (i RunbookAssociationPropertyArgs) ToRunbookAssociationPropertyOutputWithContext(ctx context.Context) RunbookAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookAssociationPropertyOutput)
}

func (i RunbookAssociationPropertyArgs) ToRunbookAssociationPropertyPtrOutput() RunbookAssociationPropertyPtrOutput {
	return i.ToRunbookAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i RunbookAssociationPropertyArgs) ToRunbookAssociationPropertyPtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookAssociationPropertyOutput).ToRunbookAssociationPropertyPtrOutputWithContext(ctx)
}









type RunbookAssociationPropertyPtrInput interface {
	pulumi.Input

	ToRunbookAssociationPropertyPtrOutput() RunbookAssociationPropertyPtrOutput
	ToRunbookAssociationPropertyPtrOutputWithContext(context.Context) RunbookAssociationPropertyPtrOutput
}

type runbookAssociationPropertyPtrType RunbookAssociationPropertyArgs

func RunbookAssociationPropertyPtr(v *RunbookAssociationPropertyArgs) RunbookAssociationPropertyPtrInput {
	return (*runbookAssociationPropertyPtrType)(v)
}

func (*runbookAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookAssociationProperty)(nil)).Elem()
}

func (i *runbookAssociationPropertyPtrType) ToRunbookAssociationPropertyPtrOutput() RunbookAssociationPropertyPtrOutput {
	return i.ToRunbookAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *runbookAssociationPropertyPtrType) ToRunbookAssociationPropertyPtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookAssociationPropertyPtrOutput)
}

type RunbookAssociationPropertyOutput struct{ *pulumi.OutputState }

func (RunbookAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookAssociationProperty)(nil)).Elem()
}

func (o RunbookAssociationPropertyOutput) ToRunbookAssociationPropertyOutput() RunbookAssociationPropertyOutput {
	return o
}

func (o RunbookAssociationPropertyOutput) ToRunbookAssociationPropertyOutputWithContext(ctx context.Context) RunbookAssociationPropertyOutput {
	return o
}

func (o RunbookAssociationPropertyOutput) ToRunbookAssociationPropertyPtrOutput() RunbookAssociationPropertyPtrOutput {
	return o.ToRunbookAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o RunbookAssociationPropertyOutput) ToRunbookAssociationPropertyPtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookAssociationProperty) *RunbookAssociationProperty {
		return &v
	}).(RunbookAssociationPropertyPtrOutput)
}

func (o RunbookAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunbookAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (RunbookAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookAssociationProperty)(nil)).Elem()
}

func (o RunbookAssociationPropertyPtrOutput) ToRunbookAssociationPropertyPtrOutput() RunbookAssociationPropertyPtrOutput {
	return o
}

func (o RunbookAssociationPropertyPtrOutput) ToRunbookAssociationPropertyPtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyPtrOutput {
	return o
}

func (o RunbookAssociationPropertyPtrOutput) Elem() RunbookAssociationPropertyOutput {
	return o.ApplyT(func(v *RunbookAssociationProperty) RunbookAssociationProperty {
		if v != nil {
			return *v
		}
		var ret RunbookAssociationProperty
		return ret
	}).(RunbookAssociationPropertyOutput)
}

func (o RunbookAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type RunbookAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}





type RunbookAssociationPropertyResponseInput interface {
	pulumi.Input

	ToRunbookAssociationPropertyResponseOutput() RunbookAssociationPropertyResponseOutput
	ToRunbookAssociationPropertyResponseOutputWithContext(context.Context) RunbookAssociationPropertyResponseOutput
}

type RunbookAssociationPropertyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (RunbookAssociationPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookAssociationPropertyResponse)(nil)).Elem()
}

func (i RunbookAssociationPropertyResponseArgs) ToRunbookAssociationPropertyResponseOutput() RunbookAssociationPropertyResponseOutput {
	return i.ToRunbookAssociationPropertyResponseOutputWithContext(context.Background())
}

func (i RunbookAssociationPropertyResponseArgs) ToRunbookAssociationPropertyResponseOutputWithContext(ctx context.Context) RunbookAssociationPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookAssociationPropertyResponseOutput)
}

func (i RunbookAssociationPropertyResponseArgs) ToRunbookAssociationPropertyResponsePtrOutput() RunbookAssociationPropertyResponsePtrOutput {
	return i.ToRunbookAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i RunbookAssociationPropertyResponseArgs) ToRunbookAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookAssociationPropertyResponseOutput).ToRunbookAssociationPropertyResponsePtrOutputWithContext(ctx)
}









type RunbookAssociationPropertyResponsePtrInput interface {
	pulumi.Input

	ToRunbookAssociationPropertyResponsePtrOutput() RunbookAssociationPropertyResponsePtrOutput
	ToRunbookAssociationPropertyResponsePtrOutputWithContext(context.Context) RunbookAssociationPropertyResponsePtrOutput
}

type runbookAssociationPropertyResponsePtrType RunbookAssociationPropertyResponseArgs

func RunbookAssociationPropertyResponsePtr(v *RunbookAssociationPropertyResponseArgs) RunbookAssociationPropertyResponsePtrInput {
	return (*runbookAssociationPropertyResponsePtrType)(v)
}

func (*runbookAssociationPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookAssociationPropertyResponse)(nil)).Elem()
}

func (i *runbookAssociationPropertyResponsePtrType) ToRunbookAssociationPropertyResponsePtrOutput() RunbookAssociationPropertyResponsePtrOutput {
	return i.ToRunbookAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *runbookAssociationPropertyResponsePtrType) ToRunbookAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookAssociationPropertyResponsePtrOutput)
}

type RunbookAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (RunbookAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookAssociationPropertyResponse)(nil)).Elem()
}

func (o RunbookAssociationPropertyResponseOutput) ToRunbookAssociationPropertyResponseOutput() RunbookAssociationPropertyResponseOutput {
	return o
}

func (o RunbookAssociationPropertyResponseOutput) ToRunbookAssociationPropertyResponseOutputWithContext(ctx context.Context) RunbookAssociationPropertyResponseOutput {
	return o
}

func (o RunbookAssociationPropertyResponseOutput) ToRunbookAssociationPropertyResponsePtrOutput() RunbookAssociationPropertyResponsePtrOutput {
	return o.ToRunbookAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (o RunbookAssociationPropertyResponseOutput) ToRunbookAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookAssociationPropertyResponse) *RunbookAssociationPropertyResponse {
		return &v
	}).(RunbookAssociationPropertyResponsePtrOutput)
}

func (o RunbookAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunbookAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (RunbookAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookAssociationPropertyResponse)(nil)).Elem()
}

func (o RunbookAssociationPropertyResponsePtrOutput) ToRunbookAssociationPropertyResponsePtrOutput() RunbookAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunbookAssociationPropertyResponsePtrOutput) ToRunbookAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunbookAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunbookAssociationPropertyResponsePtrOutput) Elem() RunbookAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *RunbookAssociationPropertyResponse) RunbookAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret RunbookAssociationPropertyResponse
		return ret
	}).(RunbookAssociationPropertyResponseOutput)
}

func (o RunbookAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type RunbookDraft struct {
	CreationTime     *string                     `pulumi:"creationTime"`
	DraftContentLink *ContentLink                `pulumi:"draftContentLink"`
	InEdit           *bool                       `pulumi:"inEdit"`
	LastModifiedTime *string                     `pulumi:"lastModifiedTime"`
	OutputTypes      []string                    `pulumi:"outputTypes"`
	Parameters       map[string]RunbookParameter `pulumi:"parameters"`
}





type RunbookDraftInput interface {
	pulumi.Input

	ToRunbookDraftOutput() RunbookDraftOutput
	ToRunbookDraftOutputWithContext(context.Context) RunbookDraftOutput
}

type RunbookDraftArgs struct {
	CreationTime     pulumi.StringPtrInput    `pulumi:"creationTime"`
	DraftContentLink ContentLinkPtrInput      `pulumi:"draftContentLink"`
	InEdit           pulumi.BoolPtrInput      `pulumi:"inEdit"`
	LastModifiedTime pulumi.StringPtrInput    `pulumi:"lastModifiedTime"`
	OutputTypes      pulumi.StringArrayInput  `pulumi:"outputTypes"`
	Parameters       RunbookParameterMapInput `pulumi:"parameters"`
}

func (RunbookDraftArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraft)(nil)).Elem()
}

func (i RunbookDraftArgs) ToRunbookDraftOutput() RunbookDraftOutput {
	return i.ToRunbookDraftOutputWithContext(context.Background())
}

func (i RunbookDraftArgs) ToRunbookDraftOutputWithContext(ctx context.Context) RunbookDraftOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftOutput)
}

func (i RunbookDraftArgs) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return i.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (i RunbookDraftArgs) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftOutput).ToRunbookDraftPtrOutputWithContext(ctx)
}









type RunbookDraftPtrInput interface {
	pulumi.Input

	ToRunbookDraftPtrOutput() RunbookDraftPtrOutput
	ToRunbookDraftPtrOutputWithContext(context.Context) RunbookDraftPtrOutput
}

type runbookDraftPtrType RunbookDraftArgs

func RunbookDraftPtr(v *RunbookDraftArgs) RunbookDraftPtrInput {
	return (*runbookDraftPtrType)(v)
}

func (*runbookDraftPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraft)(nil)).Elem()
}

func (i *runbookDraftPtrType) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return i.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (i *runbookDraftPtrType) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftPtrOutput)
}

type RunbookDraftOutput struct{ *pulumi.OutputState }

func (RunbookDraftOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraft)(nil)).Elem()
}

func (o RunbookDraftOutput) ToRunbookDraftOutput() RunbookDraftOutput {
	return o
}

func (o RunbookDraftOutput) ToRunbookDraftOutputWithContext(ctx context.Context) RunbookDraftOutput {
	return o
}

func (o RunbookDraftOutput) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return o.ToRunbookDraftPtrOutputWithContext(context.Background())
}

func (o RunbookDraftOutput) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookDraft) *RunbookDraft {
		return &v
	}).(RunbookDraftPtrOutput)
}

func (o RunbookDraftOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftOutput) DraftContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *ContentLink { return v.DraftContentLink }).(ContentLinkPtrOutput)
}

func (o RunbookDraftOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *bool { return v.InEdit }).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraft) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunbookDraft) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

func (o RunbookDraftOutput) Parameters() RunbookParameterMapOutput {
	return o.ApplyT(func(v RunbookDraft) map[string]RunbookParameter { return v.Parameters }).(RunbookParameterMapOutput)
}

type RunbookDraftPtrOutput struct{ *pulumi.OutputState }

func (RunbookDraftPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraft)(nil)).Elem()
}

func (o RunbookDraftPtrOutput) ToRunbookDraftPtrOutput() RunbookDraftPtrOutput {
	return o
}

func (o RunbookDraftPtrOutput) ToRunbookDraftPtrOutputWithContext(ctx context.Context) RunbookDraftPtrOutput {
	return o
}

func (o RunbookDraftPtrOutput) Elem() RunbookDraftOutput {
	return o.ApplyT(func(v *RunbookDraft) RunbookDraft {
		if v != nil {
			return *v
		}
		var ret RunbookDraft
		return ret
	}).(RunbookDraftOutput)
}

func (o RunbookDraftPtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftPtrOutput) DraftContentLink() ContentLinkPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *ContentLink {
		if v == nil {
			return nil
		}
		return v.DraftContentLink
	}).(ContentLinkPtrOutput)
}

func (o RunbookDraftPtrOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *bool {
		if v == nil {
			return nil
		}
		return v.InEdit
	}).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftPtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraft) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftPtrOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunbookDraft) []string {
		if v == nil {
			return nil
		}
		return v.OutputTypes
	}).(pulumi.StringArrayOutput)
}

func (o RunbookDraftPtrOutput) Parameters() RunbookParameterMapOutput {
	return o.ApplyT(func(v *RunbookDraft) map[string]RunbookParameter {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(RunbookParameterMapOutput)
}

type RunbookDraftResponse struct {
	CreationTime     *string                             `pulumi:"creationTime"`
	DraftContentLink *ContentLinkResponse                `pulumi:"draftContentLink"`
	InEdit           *bool                               `pulumi:"inEdit"`
	LastModifiedTime *string                             `pulumi:"lastModifiedTime"`
	OutputTypes      []string                            `pulumi:"outputTypes"`
	Parameters       map[string]RunbookParameterResponse `pulumi:"parameters"`
}





type RunbookDraftResponseInput interface {
	pulumi.Input

	ToRunbookDraftResponseOutput() RunbookDraftResponseOutput
	ToRunbookDraftResponseOutputWithContext(context.Context) RunbookDraftResponseOutput
}

type RunbookDraftResponseArgs struct {
	CreationTime     pulumi.StringPtrInput            `pulumi:"creationTime"`
	DraftContentLink ContentLinkResponsePtrInput      `pulumi:"draftContentLink"`
	InEdit           pulumi.BoolPtrInput              `pulumi:"inEdit"`
	LastModifiedTime pulumi.StringPtrInput            `pulumi:"lastModifiedTime"`
	OutputTypes      pulumi.StringArrayInput          `pulumi:"outputTypes"`
	Parameters       RunbookParameterResponseMapInput `pulumi:"parameters"`
}

func (RunbookDraftResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraftResponse)(nil)).Elem()
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponseOutput() RunbookDraftResponseOutput {
	return i.ToRunbookDraftResponseOutputWithContext(context.Background())
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponseOutputWithContext(ctx context.Context) RunbookDraftResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftResponseOutput)
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return i.ToRunbookDraftResponsePtrOutputWithContext(context.Background())
}

func (i RunbookDraftResponseArgs) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftResponseOutput).ToRunbookDraftResponsePtrOutputWithContext(ctx)
}









type RunbookDraftResponsePtrInput interface {
	pulumi.Input

	ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput
	ToRunbookDraftResponsePtrOutputWithContext(context.Context) RunbookDraftResponsePtrOutput
}

type runbookDraftResponsePtrType RunbookDraftResponseArgs

func RunbookDraftResponsePtr(v *RunbookDraftResponseArgs) RunbookDraftResponsePtrInput {
	return (*runbookDraftResponsePtrType)(v)
}

func (*runbookDraftResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraftResponse)(nil)).Elem()
}

func (i *runbookDraftResponsePtrType) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return i.ToRunbookDraftResponsePtrOutputWithContext(context.Background())
}

func (i *runbookDraftResponsePtrType) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookDraftResponsePtrOutput)
}

type RunbookDraftResponseOutput struct{ *pulumi.OutputState }

func (RunbookDraftResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookDraftResponse)(nil)).Elem()
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponseOutput() RunbookDraftResponseOutput {
	return o
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponseOutputWithContext(ctx context.Context) RunbookDraftResponseOutput {
	return o
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return o.ToRunbookDraftResponsePtrOutputWithContext(context.Background())
}

func (o RunbookDraftResponseOutput) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunbookDraftResponse) *RunbookDraftResponse {
		return &v
	}).(RunbookDraftResponsePtrOutput)
}

func (o RunbookDraftResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponseOutput) DraftContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *ContentLinkResponse { return v.DraftContentLink }).(ContentLinkResponsePtrOutput)
}

func (o RunbookDraftResponseOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *bool { return v.InEdit }).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftResponseOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookDraftResponse) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponseOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RunbookDraftResponse) []string { return v.OutputTypes }).(pulumi.StringArrayOutput)
}

func (o RunbookDraftResponseOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v RunbookDraftResponse) map[string]RunbookParameterResponse { return v.Parameters }).(RunbookParameterResponseMapOutput)
}

type RunbookDraftResponsePtrOutput struct{ *pulumi.OutputState }

func (RunbookDraftResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunbookDraftResponse)(nil)).Elem()
}

func (o RunbookDraftResponsePtrOutput) ToRunbookDraftResponsePtrOutput() RunbookDraftResponsePtrOutput {
	return o
}

func (o RunbookDraftResponsePtrOutput) ToRunbookDraftResponsePtrOutputWithContext(ctx context.Context) RunbookDraftResponsePtrOutput {
	return o
}

func (o RunbookDraftResponsePtrOutput) Elem() RunbookDraftResponseOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) RunbookDraftResponse {
		if v != nil {
			return *v
		}
		var ret RunbookDraftResponse
		return ret
	}).(RunbookDraftResponseOutput)
}

func (o RunbookDraftResponsePtrOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreationTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponsePtrOutput) DraftContentLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *ContentLinkResponse {
		if v == nil {
			return nil
		}
		return v.DraftContentLink
	}).(ContentLinkResponsePtrOutput)
}

func (o RunbookDraftResponsePtrOutput) InEdit() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *bool {
		if v == nil {
			return nil
		}
		return v.InEdit
	}).(pulumi.BoolPtrOutput)
}

func (o RunbookDraftResponsePtrOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedTime
	}).(pulumi.StringPtrOutput)
}

func (o RunbookDraftResponsePtrOutput) OutputTypes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) []string {
		if v == nil {
			return nil
		}
		return v.OutputTypes
	}).(pulumi.StringArrayOutput)
}

func (o RunbookDraftResponsePtrOutput) Parameters() RunbookParameterResponseMapOutput {
	return o.ApplyT(func(v *RunbookDraftResponse) map[string]RunbookParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(RunbookParameterResponseMapOutput)
}

type RunbookParameter struct {
	DefaultValue *string `pulumi:"defaultValue"`
	IsMandatory  *bool   `pulumi:"isMandatory"`
	Position     *int    `pulumi:"position"`
	Type         *string `pulumi:"type"`
}





type RunbookParameterInput interface {
	pulumi.Input

	ToRunbookParameterOutput() RunbookParameterOutput
	ToRunbookParameterOutputWithContext(context.Context) RunbookParameterOutput
}

type RunbookParameterArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	IsMandatory  pulumi.BoolPtrInput   `pulumi:"isMandatory"`
	Position     pulumi.IntPtrInput    `pulumi:"position"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (RunbookParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameter)(nil)).Elem()
}

func (i RunbookParameterArgs) ToRunbookParameterOutput() RunbookParameterOutput {
	return i.ToRunbookParameterOutputWithContext(context.Background())
}

func (i RunbookParameterArgs) ToRunbookParameterOutputWithContext(ctx context.Context) RunbookParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterOutput)
}





type RunbookParameterMapInput interface {
	pulumi.Input

	ToRunbookParameterMapOutput() RunbookParameterMapOutput
	ToRunbookParameterMapOutputWithContext(context.Context) RunbookParameterMapOutput
}

type RunbookParameterMap map[string]RunbookParameterInput

func (RunbookParameterMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameter)(nil)).Elem()
}

func (i RunbookParameterMap) ToRunbookParameterMapOutput() RunbookParameterMapOutput {
	return i.ToRunbookParameterMapOutputWithContext(context.Background())
}

func (i RunbookParameterMap) ToRunbookParameterMapOutputWithContext(ctx context.Context) RunbookParameterMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterMapOutput)
}

type RunbookParameterOutput struct{ *pulumi.OutputState }

func (RunbookParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameter)(nil)).Elem()
}

func (o RunbookParameterOutput) ToRunbookParameterOutput() RunbookParameterOutput {
	return o
}

func (o RunbookParameterOutput) ToRunbookParameterOutputWithContext(ctx context.Context) RunbookParameterOutput {
	return o
}

func (o RunbookParameterOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o RunbookParameterOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o RunbookParameterOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func (o RunbookParameterOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameter) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RunbookParameterMapOutput struct{ *pulumi.OutputState }

func (RunbookParameterMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameter)(nil)).Elem()
}

func (o RunbookParameterMapOutput) ToRunbookParameterMapOutput() RunbookParameterMapOutput {
	return o
}

func (o RunbookParameterMapOutput) ToRunbookParameterMapOutputWithContext(ctx context.Context) RunbookParameterMapOutput {
	return o
}

func (o RunbookParameterMapOutput) MapIndex(k pulumi.StringInput) RunbookParameterOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunbookParameter {
		return vs[0].(map[string]RunbookParameter)[vs[1].(string)]
	}).(RunbookParameterOutput)
}

type RunbookParameterResponse struct {
	DefaultValue *string `pulumi:"defaultValue"`
	IsMandatory  *bool   `pulumi:"isMandatory"`
	Position     *int    `pulumi:"position"`
	Type         *string `pulumi:"type"`
}





type RunbookParameterResponseInput interface {
	pulumi.Input

	ToRunbookParameterResponseOutput() RunbookParameterResponseOutput
	ToRunbookParameterResponseOutputWithContext(context.Context) RunbookParameterResponseOutput
}

type RunbookParameterResponseArgs struct {
	DefaultValue pulumi.StringPtrInput `pulumi:"defaultValue"`
	IsMandatory  pulumi.BoolPtrInput   `pulumi:"isMandatory"`
	Position     pulumi.IntPtrInput    `pulumi:"position"`
	Type         pulumi.StringPtrInput `pulumi:"type"`
}

func (RunbookParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameterResponse)(nil)).Elem()
}

func (i RunbookParameterResponseArgs) ToRunbookParameterResponseOutput() RunbookParameterResponseOutput {
	return i.ToRunbookParameterResponseOutputWithContext(context.Background())
}

func (i RunbookParameterResponseArgs) ToRunbookParameterResponseOutputWithContext(ctx context.Context) RunbookParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterResponseOutput)
}





type RunbookParameterResponseMapInput interface {
	pulumi.Input

	ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput
	ToRunbookParameterResponseMapOutputWithContext(context.Context) RunbookParameterResponseMapOutput
}

type RunbookParameterResponseMap map[string]RunbookParameterResponseInput

func (RunbookParameterResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameterResponse)(nil)).Elem()
}

func (i RunbookParameterResponseMap) ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput {
	return i.ToRunbookParameterResponseMapOutputWithContext(context.Background())
}

func (i RunbookParameterResponseMap) ToRunbookParameterResponseMapOutputWithContext(ctx context.Context) RunbookParameterResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunbookParameterResponseMapOutput)
}

type RunbookParameterResponseOutput struct{ *pulumi.OutputState }

func (RunbookParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunbookParameterResponse)(nil)).Elem()
}

func (o RunbookParameterResponseOutput) ToRunbookParameterResponseOutput() RunbookParameterResponseOutput {
	return o
}

func (o RunbookParameterResponseOutput) ToRunbookParameterResponseOutputWithContext(ctx context.Context) RunbookParameterResponseOutput {
	return o
}

func (o RunbookParameterResponseOutput) DefaultValue() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *string { return v.DefaultValue }).(pulumi.StringPtrOutput)
}

func (o RunbookParameterResponseOutput) IsMandatory() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *bool { return v.IsMandatory }).(pulumi.BoolPtrOutput)
}

func (o RunbookParameterResponseOutput) Position() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *int { return v.Position }).(pulumi.IntPtrOutput)
}

func (o RunbookParameterResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunbookParameterResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type RunbookParameterResponseMapOutput struct{ *pulumi.OutputState }

func (RunbookParameterResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]RunbookParameterResponse)(nil)).Elem()
}

func (o RunbookParameterResponseMapOutput) ToRunbookParameterResponseMapOutput() RunbookParameterResponseMapOutput {
	return o
}

func (o RunbookParameterResponseMapOutput) ToRunbookParameterResponseMapOutputWithContext(ctx context.Context) RunbookParameterResponseMapOutput {
	return o
}

func (o RunbookParameterResponseMapOutput) MapIndex(k pulumi.StringInput) RunbookParameterResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) RunbookParameterResponse {
		return vs[0].(map[string]RunbookParameterResponse)[vs[1].(string)]
	}).(RunbookParameterResponseOutput)
}

type ScheduleAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type ScheduleAssociationPropertyInput interface {
	pulumi.Input

	ToScheduleAssociationPropertyOutput() ScheduleAssociationPropertyOutput
	ToScheduleAssociationPropertyOutputWithContext(context.Context) ScheduleAssociationPropertyOutput
}

type ScheduleAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ScheduleAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleAssociationProperty)(nil)).Elem()
}

func (i ScheduleAssociationPropertyArgs) ToScheduleAssociationPropertyOutput() ScheduleAssociationPropertyOutput {
	return i.ToScheduleAssociationPropertyOutputWithContext(context.Background())
}

func (i ScheduleAssociationPropertyArgs) ToScheduleAssociationPropertyOutputWithContext(ctx context.Context) ScheduleAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleAssociationPropertyOutput)
}

func (i ScheduleAssociationPropertyArgs) ToScheduleAssociationPropertyPtrOutput() ScheduleAssociationPropertyPtrOutput {
	return i.ToScheduleAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i ScheduleAssociationPropertyArgs) ToScheduleAssociationPropertyPtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleAssociationPropertyOutput).ToScheduleAssociationPropertyPtrOutputWithContext(ctx)
}









type ScheduleAssociationPropertyPtrInput interface {
	pulumi.Input

	ToScheduleAssociationPropertyPtrOutput() ScheduleAssociationPropertyPtrOutput
	ToScheduleAssociationPropertyPtrOutputWithContext(context.Context) ScheduleAssociationPropertyPtrOutput
}

type scheduleAssociationPropertyPtrType ScheduleAssociationPropertyArgs

func ScheduleAssociationPropertyPtr(v *ScheduleAssociationPropertyArgs) ScheduleAssociationPropertyPtrInput {
	return (*scheduleAssociationPropertyPtrType)(v)
}

func (*scheduleAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleAssociationProperty)(nil)).Elem()
}

func (i *scheduleAssociationPropertyPtrType) ToScheduleAssociationPropertyPtrOutput() ScheduleAssociationPropertyPtrOutput {
	return i.ToScheduleAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *scheduleAssociationPropertyPtrType) ToScheduleAssociationPropertyPtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleAssociationPropertyPtrOutput)
}

type ScheduleAssociationPropertyOutput struct{ *pulumi.OutputState }

func (ScheduleAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleAssociationProperty)(nil)).Elem()
}

func (o ScheduleAssociationPropertyOutput) ToScheduleAssociationPropertyOutput() ScheduleAssociationPropertyOutput {
	return o
}

func (o ScheduleAssociationPropertyOutput) ToScheduleAssociationPropertyOutputWithContext(ctx context.Context) ScheduleAssociationPropertyOutput {
	return o
}

func (o ScheduleAssociationPropertyOutput) ToScheduleAssociationPropertyPtrOutput() ScheduleAssociationPropertyPtrOutput {
	return o.ToScheduleAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o ScheduleAssociationPropertyOutput) ToScheduleAssociationPropertyPtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleAssociationProperty) *ScheduleAssociationProperty {
		return &v
	}).(ScheduleAssociationPropertyPtrOutput)
}

func (o ScheduleAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScheduleAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (ScheduleAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleAssociationProperty)(nil)).Elem()
}

func (o ScheduleAssociationPropertyPtrOutput) ToScheduleAssociationPropertyPtrOutput() ScheduleAssociationPropertyPtrOutput {
	return o
}

func (o ScheduleAssociationPropertyPtrOutput) ToScheduleAssociationPropertyPtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyPtrOutput {
	return o
}

func (o ScheduleAssociationPropertyPtrOutput) Elem() ScheduleAssociationPropertyOutput {
	return o.ApplyT(func(v *ScheduleAssociationProperty) ScheduleAssociationProperty {
		if v != nil {
			return *v
		}
		var ret ScheduleAssociationProperty
		return ret
	}).(ScheduleAssociationPropertyOutput)
}

func (o ScheduleAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ScheduleAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}





type ScheduleAssociationPropertyResponseInput interface {
	pulumi.Input

	ToScheduleAssociationPropertyResponseOutput() ScheduleAssociationPropertyResponseOutput
	ToScheduleAssociationPropertyResponseOutputWithContext(context.Context) ScheduleAssociationPropertyResponseOutput
}

type ScheduleAssociationPropertyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ScheduleAssociationPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleAssociationPropertyResponse)(nil)).Elem()
}

func (i ScheduleAssociationPropertyResponseArgs) ToScheduleAssociationPropertyResponseOutput() ScheduleAssociationPropertyResponseOutput {
	return i.ToScheduleAssociationPropertyResponseOutputWithContext(context.Background())
}

func (i ScheduleAssociationPropertyResponseArgs) ToScheduleAssociationPropertyResponseOutputWithContext(ctx context.Context) ScheduleAssociationPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleAssociationPropertyResponseOutput)
}

func (i ScheduleAssociationPropertyResponseArgs) ToScheduleAssociationPropertyResponsePtrOutput() ScheduleAssociationPropertyResponsePtrOutput {
	return i.ToScheduleAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i ScheduleAssociationPropertyResponseArgs) ToScheduleAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleAssociationPropertyResponseOutput).ToScheduleAssociationPropertyResponsePtrOutputWithContext(ctx)
}









type ScheduleAssociationPropertyResponsePtrInput interface {
	pulumi.Input

	ToScheduleAssociationPropertyResponsePtrOutput() ScheduleAssociationPropertyResponsePtrOutput
	ToScheduleAssociationPropertyResponsePtrOutputWithContext(context.Context) ScheduleAssociationPropertyResponsePtrOutput
}

type scheduleAssociationPropertyResponsePtrType ScheduleAssociationPropertyResponseArgs

func ScheduleAssociationPropertyResponsePtr(v *ScheduleAssociationPropertyResponseArgs) ScheduleAssociationPropertyResponsePtrInput {
	return (*scheduleAssociationPropertyResponsePtrType)(v)
}

func (*scheduleAssociationPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleAssociationPropertyResponse)(nil)).Elem()
}

func (i *scheduleAssociationPropertyResponsePtrType) ToScheduleAssociationPropertyResponsePtrOutput() ScheduleAssociationPropertyResponsePtrOutput {
	return i.ToScheduleAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *scheduleAssociationPropertyResponsePtrType) ToScheduleAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleAssociationPropertyResponsePtrOutput)
}

type ScheduleAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (ScheduleAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleAssociationPropertyResponse)(nil)).Elem()
}

func (o ScheduleAssociationPropertyResponseOutput) ToScheduleAssociationPropertyResponseOutput() ScheduleAssociationPropertyResponseOutput {
	return o
}

func (o ScheduleAssociationPropertyResponseOutput) ToScheduleAssociationPropertyResponseOutputWithContext(ctx context.Context) ScheduleAssociationPropertyResponseOutput {
	return o
}

func (o ScheduleAssociationPropertyResponseOutput) ToScheduleAssociationPropertyResponsePtrOutput() ScheduleAssociationPropertyResponsePtrOutput {
	return o.ToScheduleAssociationPropertyResponsePtrOutputWithContext(context.Background())
}

func (o ScheduleAssociationPropertyResponseOutput) ToScheduleAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ScheduleAssociationPropertyResponse) *ScheduleAssociationPropertyResponse {
		return &v
	}).(ScheduleAssociationPropertyResponsePtrOutput)
}

func (o ScheduleAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScheduleAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleAssociationPropertyResponse)(nil)).Elem()
}

func (o ScheduleAssociationPropertyResponsePtrOutput) ToScheduleAssociationPropertyResponsePtrOutput() ScheduleAssociationPropertyResponsePtrOutput {
	return o
}

func (o ScheduleAssociationPropertyResponsePtrOutput) ToScheduleAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) ScheduleAssociationPropertyResponsePtrOutput {
	return o
}

func (o ScheduleAssociationPropertyResponsePtrOutput) Elem() ScheduleAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *ScheduleAssociationPropertyResponse) ScheduleAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleAssociationPropertyResponse
		return ret
	}).(ScheduleAssociationPropertyResponseOutput)
}

func (o ScheduleAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type Sku struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Sku) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Sku) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Sku) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Capacity *int    `pulumi:"capacity"`
	Family   *string `pulumi:"family"`
	Name     string  `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Capacity pulumi.IntPtrInput    `pulumi:"capacity"`
	Family   pulumi.StringPtrInput `pulumi:"family"`
	Name     pulumi.StringInput    `pulumi:"name"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
}

func (o SkuResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleInput)(nil)).Elem(), AdvancedScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedSchedulePtrInput)(nil)).Elem(), AdvancedScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleMonthlyOccurrenceInput)(nil)).Elem(), AdvancedScheduleMonthlyOccurrenceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleMonthlyOccurrenceArrayInput)(nil)).Elem(), AdvancedScheduleMonthlyOccurrenceArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleMonthlyOccurrenceResponseInput)(nil)).Elem(), AdvancedScheduleMonthlyOccurrenceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleMonthlyOccurrenceResponseArrayInput)(nil)).Elem(), AdvancedScheduleMonthlyOccurrenceResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleResponseInput)(nil)).Elem(), AdvancedScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AdvancedScheduleResponsePtrInput)(nil)).Elem(), AdvancedScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionTypeAssociationPropertyInput)(nil)).Elem(), ConnectionTypeAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionTypeAssociationPropertyPtrInput)(nil)).Elem(), ConnectionTypeAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionTypeAssociationPropertyResponseInput)(nil)).Elem(), ConnectionTypeAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionTypeAssociationPropertyResponsePtrInput)(nil)).Elem(), ConnectionTypeAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashInput)(nil)).Elem(), ContentHashArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashPtrInput)(nil)).Elem(), ContentHashArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashResponseInput)(nil)).Elem(), ContentHashResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentHashResponsePtrInput)(nil)).Elem(), ContentHashResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkInput)(nil)).Elem(), ContentLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkPtrInput)(nil)).Elem(), ContentLinkArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkResponseInput)(nil)).Elem(), ContentLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentLinkResponsePtrInput)(nil)).Elem(), ContentLinkResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentSourceInput)(nil)).Elem(), ContentSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentSourcePtrInput)(nil)).Elem(), ContentSourceArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentSourceResponseInput)(nil)).Elem(), ContentSourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ContentSourceResponsePtrInput)(nil)).Elem(), ContentSourceResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationAssociationPropertyInput)(nil)).Elem(), DscConfigurationAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationAssociationPropertyPtrInput)(nil)).Elem(), DscConfigurationAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationAssociationPropertyResponseInput)(nil)).Elem(), DscConfigurationAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationAssociationPropertyResponsePtrInput)(nil)).Elem(), DscConfigurationAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationParameterInput)(nil)).Elem(), DscConfigurationParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationParameterMapInput)(nil)).Elem(), DscConfigurationParameterMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationParameterResponseInput)(nil)).Elem(), DscConfigurationParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DscConfigurationParameterResponseMapInput)(nil)).Elem(), DscConfigurationParameterResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*FieldDefinitionInput)(nil)).Elem(), FieldDefinitionArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FieldDefinitionMapInput)(nil)).Elem(), FieldDefinitionMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*FieldDefinitionResponseInput)(nil)).Elem(), FieldDefinitionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FieldDefinitionResponseMapInput)(nil)).Elem(), FieldDefinitionResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyResponseInput)(nil)).Elem(), KeyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KeyResponseArrayInput)(nil)).Elem(), KeyResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleErrorInfoResponseInput)(nil)).Elem(), ModuleErrorInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ModuleErrorInfoResponsePtrInput)(nil)).Elem(), ModuleErrorInfoResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookAssociationPropertyInput)(nil)).Elem(), RunbookAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookAssociationPropertyPtrInput)(nil)).Elem(), RunbookAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookAssociationPropertyResponseInput)(nil)).Elem(), RunbookAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookAssociationPropertyResponsePtrInput)(nil)).Elem(), RunbookAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftInput)(nil)).Elem(), RunbookDraftArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftPtrInput)(nil)).Elem(), RunbookDraftArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftResponseInput)(nil)).Elem(), RunbookDraftResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookDraftResponsePtrInput)(nil)).Elem(), RunbookDraftResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterInput)(nil)).Elem(), RunbookParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterMapInput)(nil)).Elem(), RunbookParameterMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterResponseInput)(nil)).Elem(), RunbookParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RunbookParameterResponseMapInput)(nil)).Elem(), RunbookParameterResponseMap{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleAssociationPropertyInput)(nil)).Elem(), ScheduleAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleAssociationPropertyPtrInput)(nil)).Elem(), ScheduleAssociationPropertyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleAssociationPropertyResponseInput)(nil)).Elem(), ScheduleAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ScheduleAssociationPropertyResponsePtrInput)(nil)).Elem(), ScheduleAssociationPropertyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterOutputType(AdvancedScheduleOutput{})
	pulumi.RegisterOutputType(AdvancedSchedulePtrOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceArrayOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceResponseOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceResponseArrayOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleResponseOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionTypeAssociationPropertyOutput{})
	pulumi.RegisterOutputType(ConnectionTypeAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(ConnectionTypeAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(ConnectionTypeAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentHashOutput{})
	pulumi.RegisterOutputType(ContentHashPtrOutput{})
	pulumi.RegisterOutputType(ContentHashResponseOutput{})
	pulumi.RegisterOutputType(ContentHashResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentLinkOutput{})
	pulumi.RegisterOutputType(ContentLinkPtrOutput{})
	pulumi.RegisterOutputType(ContentLinkResponseOutput{})
	pulumi.RegisterOutputType(ContentLinkResponsePtrOutput{})
	pulumi.RegisterOutputType(ContentSourceOutput{})
	pulumi.RegisterOutputType(ContentSourcePtrOutput{})
	pulumi.RegisterOutputType(ContentSourceResponseOutput{})
	pulumi.RegisterOutputType(ContentSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterMapOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterResponseOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterResponseMapOutput{})
	pulumi.RegisterOutputType(FieldDefinitionOutput{})
	pulumi.RegisterOutputType(FieldDefinitionMapOutput{})
	pulumi.RegisterOutputType(FieldDefinitionResponseOutput{})
	pulumi.RegisterOutputType(FieldDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(KeyResponseOutput{})
	pulumi.RegisterOutputType(KeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookAssociationPropertyOutput{})
	pulumi.RegisterOutputType(RunbookAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(RunbookAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(RunbookAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookDraftOutput{})
	pulumi.RegisterOutputType(RunbookDraftPtrOutput{})
	pulumi.RegisterOutputType(RunbookDraftResponseOutput{})
	pulumi.RegisterOutputType(RunbookDraftResponsePtrOutput{})
	pulumi.RegisterOutputType(RunbookParameterOutput{})
	pulumi.RegisterOutputType(RunbookParameterMapOutput{})
	pulumi.RegisterOutputType(RunbookParameterResponseOutput{})
	pulumi.RegisterOutputType(RunbookParameterResponseMapOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
