


package automation

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

type AzureQueryProperties struct {
	Locations   []string               `pulumi:"locations"`
	Scope       []string               `pulumi:"scope"`
	TagSettings *TagSettingsProperties `pulumi:"tagSettings"`
}





type AzureQueryPropertiesInput interface {
	pulumi.Input

	ToAzureQueryPropertiesOutput() AzureQueryPropertiesOutput
	ToAzureQueryPropertiesOutputWithContext(context.Context) AzureQueryPropertiesOutput
}

type AzureQueryPropertiesArgs struct {
	Locations   pulumi.StringArrayInput       `pulumi:"locations"`
	Scope       pulumi.StringArrayInput       `pulumi:"scope"`
	TagSettings TagSettingsPropertiesPtrInput `pulumi:"tagSettings"`
}

func (AzureQueryPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureQueryProperties)(nil)).Elem()
}

func (i AzureQueryPropertiesArgs) ToAzureQueryPropertiesOutput() AzureQueryPropertiesOutput {
	return i.ToAzureQueryPropertiesOutputWithContext(context.Background())
}

func (i AzureQueryPropertiesArgs) ToAzureQueryPropertiesOutputWithContext(ctx context.Context) AzureQueryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureQueryPropertiesOutput)
}





type AzureQueryPropertiesArrayInput interface {
	pulumi.Input

	ToAzureQueryPropertiesArrayOutput() AzureQueryPropertiesArrayOutput
	ToAzureQueryPropertiesArrayOutputWithContext(context.Context) AzureQueryPropertiesArrayOutput
}

type AzureQueryPropertiesArray []AzureQueryPropertiesInput

func (AzureQueryPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureQueryProperties)(nil)).Elem()
}

func (i AzureQueryPropertiesArray) ToAzureQueryPropertiesArrayOutput() AzureQueryPropertiesArrayOutput {
	return i.ToAzureQueryPropertiesArrayOutputWithContext(context.Background())
}

func (i AzureQueryPropertiesArray) ToAzureQueryPropertiesArrayOutputWithContext(ctx context.Context) AzureQueryPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureQueryPropertiesArrayOutput)
}

type AzureQueryPropertiesOutput struct{ *pulumi.OutputState }

func (AzureQueryPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureQueryProperties)(nil)).Elem()
}

func (o AzureQueryPropertiesOutput) ToAzureQueryPropertiesOutput() AzureQueryPropertiesOutput {
	return o
}

func (o AzureQueryPropertiesOutput) ToAzureQueryPropertiesOutputWithContext(ctx context.Context) AzureQueryPropertiesOutput {
	return o
}

func (o AzureQueryPropertiesOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureQueryProperties) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o AzureQueryPropertiesOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureQueryProperties) []string { return v.Scope }).(pulumi.StringArrayOutput)
}

func (o AzureQueryPropertiesOutput) TagSettings() TagSettingsPropertiesPtrOutput {
	return o.ApplyT(func(v AzureQueryProperties) *TagSettingsProperties { return v.TagSettings }).(TagSettingsPropertiesPtrOutput)
}

type AzureQueryPropertiesArrayOutput struct{ *pulumi.OutputState }

func (AzureQueryPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureQueryProperties)(nil)).Elem()
}

func (o AzureQueryPropertiesArrayOutput) ToAzureQueryPropertiesArrayOutput() AzureQueryPropertiesArrayOutput {
	return o
}

func (o AzureQueryPropertiesArrayOutput) ToAzureQueryPropertiesArrayOutputWithContext(ctx context.Context) AzureQueryPropertiesArrayOutput {
	return o
}

func (o AzureQueryPropertiesArrayOutput) Index(i pulumi.IntInput) AzureQueryPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureQueryProperties {
		return vs[0].([]AzureQueryProperties)[vs[1].(int)]
	}).(AzureQueryPropertiesOutput)
}

type AzureQueryPropertiesResponse struct {
	Locations   []string                       `pulumi:"locations"`
	Scope       []string                       `pulumi:"scope"`
	TagSettings *TagSettingsPropertiesResponse `pulumi:"tagSettings"`
}

type AzureQueryPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AzureQueryPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureQueryPropertiesResponse)(nil)).Elem()
}

func (o AzureQueryPropertiesResponseOutput) ToAzureQueryPropertiesResponseOutput() AzureQueryPropertiesResponseOutput {
	return o
}

func (o AzureQueryPropertiesResponseOutput) ToAzureQueryPropertiesResponseOutputWithContext(ctx context.Context) AzureQueryPropertiesResponseOutput {
	return o
}

func (o AzureQueryPropertiesResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureQueryPropertiesResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o AzureQueryPropertiesResponseOutput) Scope() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AzureQueryPropertiesResponse) []string { return v.Scope }).(pulumi.StringArrayOutput)
}

func (o AzureQueryPropertiesResponseOutput) TagSettings() TagSettingsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AzureQueryPropertiesResponse) *TagSettingsPropertiesResponse { return v.TagSettings }).(TagSettingsPropertiesResponsePtrOutput)
}

type AzureQueryPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureQueryPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureQueryPropertiesResponse)(nil)).Elem()
}

func (o AzureQueryPropertiesResponseArrayOutput) ToAzureQueryPropertiesResponseArrayOutput() AzureQueryPropertiesResponseArrayOutput {
	return o
}

func (o AzureQueryPropertiesResponseArrayOutput) ToAzureQueryPropertiesResponseArrayOutputWithContext(ctx context.Context) AzureQueryPropertiesResponseArrayOutput {
	return o
}

func (o AzureQueryPropertiesResponseArrayOutput) Index(i pulumi.IntInput) AzureQueryPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureQueryPropertiesResponse {
		return vs[0].([]AzureQueryPropertiesResponse)[vs[1].(int)]
	}).(AzureQueryPropertiesResponseOutput)
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

func (o ConnectionTypeAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionTypeAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ConnectionTypeAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
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

type ContentSourceResponse struct {
	Hash    *ContentHashResponse `pulumi:"hash"`
	Type    *string              `pulumi:"type"`
	Value   *string              `pulumi:"value"`
	Version *string              `pulumi:"version"`
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

func (o DscConfigurationAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DscConfigurationAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DscConfigurationAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
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

type EncryptionProperties struct {
	Identity           *EncryptionPropertiesIdentity `pulumi:"identity"`
	KeySource          *EncryptionKeySourceType      `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultProperties           `pulumi:"keyVaultProperties"`
}





type EncryptionPropertiesInput interface {
	pulumi.Input

	ToEncryptionPropertiesOutput() EncryptionPropertiesOutput
	ToEncryptionPropertiesOutputWithContext(context.Context) EncryptionPropertiesOutput
}

type EncryptionPropertiesArgs struct {
	Identity           EncryptionPropertiesIdentityPtrInput `pulumi:"identity"`
	KeySource          EncryptionKeySourceTypePtrInput      `pulumi:"keySource"`
	KeyVaultProperties KeyVaultPropertiesPtrInput           `pulumi:"keyVaultProperties"`
}

func (EncryptionPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return i.ToEncryptionPropertiesOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput)
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesArgs) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesOutput).ToEncryptionPropertiesPtrOutputWithContext(ctx)
}









type EncryptionPropertiesPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput
	ToEncryptionPropertiesPtrOutputWithContext(context.Context) EncryptionPropertiesPtrOutput
}

type encryptionPropertiesPtrType EncryptionPropertiesArgs

func EncryptionPropertiesPtr(v *EncryptionPropertiesArgs) EncryptionPropertiesPtrInput {
	return (*encryptionPropertiesPtrType)(v)
}

func (*encryptionPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return i.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesPtrType) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesPtrOutput)
}

type EncryptionPropertiesOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutput() EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesOutputWithContext(ctx context.Context) EncryptionPropertiesOutput {
	return o
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o.ToEncryptionPropertiesPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionProperties) *EncryptionProperties {
		return &v
	}).(EncryptionPropertiesPtrOutput)
}

func (o EncryptionPropertiesOutput) Identity() EncryptionPropertiesIdentityPtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *EncryptionPropertiesIdentity { return v.Identity }).(EncryptionPropertiesIdentityPtrOutput)
}

func (o EncryptionPropertiesOutput) KeySource() EncryptionKeySourceTypePtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *EncryptionKeySourceType { return v.KeySource }).(EncryptionKeySourceTypePtrOutput)
}

func (o EncryptionPropertiesOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v EncryptionProperties) *KeyVaultProperties { return v.KeyVaultProperties }).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionProperties)(nil)).Elem()
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutput() EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) ToEncryptionPropertiesPtrOutputWithContext(ctx context.Context) EncryptionPropertiesPtrOutput {
	return o
}

func (o EncryptionPropertiesPtrOutput) Elem() EncryptionPropertiesOutput {
	return o.ApplyT(func(v *EncryptionProperties) EncryptionProperties {
		if v != nil {
			return *v
		}
		var ret EncryptionProperties
		return ret
	}).(EncryptionPropertiesOutput)
}

func (o EncryptionPropertiesPtrOutput) Identity() EncryptionPropertiesIdentityPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *EncryptionPropertiesIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EncryptionPropertiesIdentityPtrOutput)
}

func (o EncryptionPropertiesPtrOutput) KeySource() EncryptionKeySourceTypePtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *EncryptionKeySourceType {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(EncryptionKeySourceTypePtrOutput)
}

func (o EncryptionPropertiesPtrOutput) KeyVaultProperties() KeyVaultPropertiesPtrOutput {
	return o.ApplyT(func(v *EncryptionProperties) *KeyVaultProperties {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesPtrOutput)
}

type EncryptionPropertiesIdentity struct {
	UserAssignedIdentity interface{} `pulumi:"userAssignedIdentity"`
}





type EncryptionPropertiesIdentityInput interface {
	pulumi.Input

	ToEncryptionPropertiesIdentityOutput() EncryptionPropertiesIdentityOutput
	ToEncryptionPropertiesIdentityOutputWithContext(context.Context) EncryptionPropertiesIdentityOutput
}

type EncryptionPropertiesIdentityArgs struct {
	UserAssignedIdentity pulumi.Input `pulumi:"userAssignedIdentity"`
}

func (EncryptionPropertiesIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesIdentity)(nil)).Elem()
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityOutput() EncryptionPropertiesIdentityOutput {
	return i.ToEncryptionPropertiesIdentityOutputWithContext(context.Background())
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesIdentityOutput)
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return i.ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Background())
}

func (i EncryptionPropertiesIdentityArgs) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesIdentityOutput).ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx)
}









type EncryptionPropertiesIdentityPtrInput interface {
	pulumi.Input

	ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput
	ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Context) EncryptionPropertiesIdentityPtrOutput
}

type encryptionPropertiesIdentityPtrType EncryptionPropertiesIdentityArgs

func EncryptionPropertiesIdentityPtr(v *EncryptionPropertiesIdentityArgs) EncryptionPropertiesIdentityPtrInput {
	return (*encryptionPropertiesIdentityPtrType)(v)
}

func (*encryptionPropertiesIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesIdentity)(nil)).Elem()
}

func (i *encryptionPropertiesIdentityPtrType) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return i.ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Background())
}

func (i *encryptionPropertiesIdentityPtrType) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EncryptionPropertiesIdentityPtrOutput)
}

type EncryptionPropertiesIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityOutput() EncryptionPropertiesIdentityOutput {
	return o
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityOutput {
	return o
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return o.ToEncryptionPropertiesIdentityPtrOutputWithContext(context.Background())
}

func (o EncryptionPropertiesIdentityOutput) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EncryptionPropertiesIdentity) *EncryptionPropertiesIdentity {
		return &v
	}).(EncryptionPropertiesIdentityPtrOutput)
}

func (o EncryptionPropertiesIdentityOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v EncryptionPropertiesIdentity) interface{} { return v.UserAssignedIdentity }).(pulumi.AnyOutput)
}

type EncryptionPropertiesIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesIdentityPtrOutput) ToEncryptionPropertiesIdentityPtrOutput() EncryptionPropertiesIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesIdentityPtrOutput) ToEncryptionPropertiesIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesIdentityPtrOutput) Elem() EncryptionPropertiesIdentityOutput {
	return o.ApplyT(func(v *EncryptionPropertiesIdentity) EncryptionPropertiesIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesIdentity
		return ret
	}).(EncryptionPropertiesIdentityOutput)
}

func (o EncryptionPropertiesIdentityPtrOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v *EncryptionPropertiesIdentity) interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.AnyOutput)
}

type EncryptionPropertiesResponse struct {
	Identity           *EncryptionPropertiesResponseIdentity `pulumi:"identity"`
	KeySource          *string                               `pulumi:"keySource"`
	KeyVaultProperties *KeyVaultPropertiesResponse           `pulumi:"keyVaultProperties"`
}

type EncryptionPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutput() EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) ToEncryptionPropertiesResponseOutputWithContext(ctx context.Context) EncryptionPropertiesResponseOutput {
	return o
}

func (o EncryptionPropertiesResponseOutput) Identity() EncryptionPropertiesResponseIdentityPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *EncryptionPropertiesResponseIdentity { return v.Identity }).(EncryptionPropertiesResponseIdentityPtrOutput)
}

func (o EncryptionPropertiesResponseOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *string { return v.KeySource }).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesResponseOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponse) *KeyVaultPropertiesResponse { return v.KeyVaultProperties }).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponse)(nil)).Elem()
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutput() EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) ToEncryptionPropertiesResponsePtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponsePtrOutput {
	return o
}

func (o EncryptionPropertiesResponsePtrOutput) Elem() EncryptionPropertiesResponseOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) EncryptionPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesResponse
		return ret
	}).(EncryptionPropertiesResponseOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) Identity() EncryptionPropertiesResponseIdentityPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *EncryptionPropertiesResponseIdentity {
		if v == nil {
			return nil
		}
		return v.Identity
	}).(EncryptionPropertiesResponseIdentityPtrOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) KeySource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeySource
	}).(pulumi.StringPtrOutput)
}

func (o EncryptionPropertiesResponsePtrOutput) KeyVaultProperties() KeyVaultPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponse) *KeyVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.KeyVaultProperties
	}).(KeyVaultPropertiesResponsePtrOutput)
}

type EncryptionPropertiesResponseIdentity struct {
	UserAssignedIdentity interface{} `pulumi:"userAssignedIdentity"`
}

type EncryptionPropertiesResponseIdentityOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EncryptionPropertiesResponseIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesResponseIdentityOutput) ToEncryptionPropertiesResponseIdentityOutput() EncryptionPropertiesResponseIdentityOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityOutput) ToEncryptionPropertiesResponseIdentityOutputWithContext(ctx context.Context) EncryptionPropertiesResponseIdentityOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v EncryptionPropertiesResponseIdentity) interface{} { return v.UserAssignedIdentity }).(pulumi.AnyOutput)
}

type EncryptionPropertiesResponseIdentityPtrOutput struct{ *pulumi.OutputState }

func (EncryptionPropertiesResponseIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EncryptionPropertiesResponseIdentity)(nil)).Elem()
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) ToEncryptionPropertiesResponseIdentityPtrOutput() EncryptionPropertiesResponseIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) ToEncryptionPropertiesResponseIdentityPtrOutputWithContext(ctx context.Context) EncryptionPropertiesResponseIdentityPtrOutput {
	return o
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) Elem() EncryptionPropertiesResponseIdentityOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponseIdentity) EncryptionPropertiesResponseIdentity {
		if v != nil {
			return *v
		}
		var ret EncryptionPropertiesResponseIdentity
		return ret
	}).(EncryptionPropertiesResponseIdentityOutput)
}

func (o EncryptionPropertiesResponseIdentityPtrOutput) UserAssignedIdentity() pulumi.AnyOutput {
	return o.ApplyT(func(v *EncryptionPropertiesResponseIdentity) interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentity
	}).(pulumi.AnyOutput)
}

type ErrorResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}





type ErrorResponseInput interface {
	pulumi.Input

	ToErrorResponseOutput() ErrorResponseOutput
	ToErrorResponseOutputWithContext(context.Context) ErrorResponseOutput
}

type ErrorResponseArgs struct {
	Code    pulumi.StringPtrInput `pulumi:"code"`
	Message pulumi.StringPtrInput `pulumi:"message"`
}

func (ErrorResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (i ErrorResponseArgs) ToErrorResponseOutput() ErrorResponseOutput {
	return i.ToErrorResponseOutputWithContext(context.Background())
}

func (i ErrorResponseArgs) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseOutput)
}

func (i ErrorResponseArgs) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return i.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (i ErrorResponseArgs) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponseOutput).ToErrorResponsePtrOutputWithContext(ctx)
}









type ErrorResponsePtrInput interface {
	pulumi.Input

	ToErrorResponsePtrOutput() ErrorResponsePtrOutput
	ToErrorResponsePtrOutputWithContext(context.Context) ErrorResponsePtrOutput
}

type errorResponsePtrType ErrorResponseArgs

func ErrorResponsePtr(v *ErrorResponseArgs) ErrorResponsePtrInput {
	return (*errorResponsePtrType)(v)
}

func (*errorResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (i *errorResponsePtrType) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return i.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (i *errorResponsePtrType) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ErrorResponsePtrOutput)
}

type ErrorResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponse)(nil)).Elem()
}

func (o ErrorResponseOutput) ToErrorResponseOutput() ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponseOutputWithContext(ctx context.Context) ErrorResponseOutput {
	return o
}

func (o ErrorResponseOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o.ToErrorResponsePtrOutputWithContext(context.Background())
}

func (o ErrorResponseOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ErrorResponse) *ErrorResponse {
		return &v
	}).(ErrorResponsePtrOutput)
}

func (o ErrorResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ErrorResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponse)(nil)).Elem()
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutput() ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) ToErrorResponsePtrOutputWithContext(ctx context.Context) ErrorResponsePtrOutput {
	return o
}

func (o ErrorResponsePtrOutput) Elem() ErrorResponseOutput {
	return o.ApplyT(func(v *ErrorResponse) ErrorResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponse
		return ret
	}).(ErrorResponseOutput)
}

func (o ErrorResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

type ErrorResponseResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
}

type ErrorResponseResponseOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutput() ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) ToErrorResponseResponseOutputWithContext(ctx context.Context) ErrorResponseResponseOutput {
	return o
}

func (o ErrorResponseResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorResponseResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorResponseResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorResponseResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

type ErrorResponseResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorResponseResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorResponseResponse)(nil)).Elem()
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutput() ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) ToErrorResponseResponsePtrOutputWithContext(ctx context.Context) ErrorResponseResponsePtrOutput {
	return o
}

func (o ErrorResponseResponsePtrOutput) Elem() ErrorResponseResponseOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) ErrorResponseResponse {
		if v != nil {
			return *v
		}
		var ret ErrorResponseResponse
		return ret
	}).(ErrorResponseResponseOutput)
}

func (o ErrorResponseResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorResponseResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorResponseResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
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

type HybridRunbookWorkerLegacyResponse struct {
	Ip               *string `pulumi:"ip"`
	LastSeenDateTime *string `pulumi:"lastSeenDateTime"`
	Name             *string `pulumi:"name"`
	RegistrationTime *string `pulumi:"registrationTime"`
}

type HybridRunbookWorkerLegacyResponseOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerLegacyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridRunbookWorkerLegacyResponse)(nil)).Elem()
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToHybridRunbookWorkerLegacyResponseOutput() HybridRunbookWorkerLegacyResponseOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseOutput) ToHybridRunbookWorkerLegacyResponseOutputWithContext(ctx context.Context) HybridRunbookWorkerLegacyResponseOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseOutput) Ip() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.Ip }).(pulumi.StringPtrOutput)
}

func (o HybridRunbookWorkerLegacyResponseOutput) LastSeenDateTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.LastSeenDateTime }).(pulumi.StringPtrOutput)
}

func (o HybridRunbookWorkerLegacyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HybridRunbookWorkerLegacyResponseOutput) RegistrationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HybridRunbookWorkerLegacyResponse) *string { return v.RegistrationTime }).(pulumi.StringPtrOutput)
}

type HybridRunbookWorkerLegacyResponseArrayOutput struct{ *pulumi.OutputState }

func (HybridRunbookWorkerLegacyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HybridRunbookWorkerLegacyResponse)(nil)).Elem()
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToHybridRunbookWorkerLegacyResponseArrayOutput() HybridRunbookWorkerLegacyResponseArrayOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) ToHybridRunbookWorkerLegacyResponseArrayOutputWithContext(ctx context.Context) HybridRunbookWorkerLegacyResponseArrayOutput {
	return o
}

func (o HybridRunbookWorkerLegacyResponseArrayOutput) Index(i pulumi.IntInput) HybridRunbookWorkerLegacyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HybridRunbookWorkerLegacyResponse {
		return vs[0].([]HybridRunbookWorkerLegacyResponse)[vs[1].(int)]
	}).(HybridRunbookWorkerLegacyResponseOutput)
}

type Identity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type IdentityInput interface {
	pulumi.Input

	ToIdentityOutput() IdentityOutput
	ToIdentityOutputWithContext(context.Context) IdentityOutput
}

type IdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (IdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (i IdentityArgs) ToIdentityOutput() IdentityOutput {
	return i.ToIdentityOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput)
}

func (i IdentityArgs) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i IdentityArgs) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityOutput).ToIdentityPtrOutputWithContext(ctx)
}









type IdentityPtrInput interface {
	pulumi.Input

	ToIdentityPtrOutput() IdentityPtrOutput
	ToIdentityPtrOutputWithContext(context.Context) IdentityPtrOutput
}

type identityPtrType IdentityArgs

func IdentityPtr(v *IdentityArgs) IdentityPtrInput {
	return (*identityPtrType)(v)
}

func (*identityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (i *identityPtrType) ToIdentityPtrOutput() IdentityPtrOutput {
	return i.ToIdentityPtrOutputWithContext(context.Background())
}

func (i *identityPtrType) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityPtrOutput)
}

type IdentityOutput struct{ *pulumi.OutputState }

func (IdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Identity)(nil)).Elem()
}

func (o IdentityOutput) ToIdentityOutput() IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityOutputWithContext(ctx context.Context) IdentityOutput {
	return o
}

func (o IdentityOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o.ToIdentityPtrOutputWithContext(context.Background())
}

func (o IdentityOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Identity) *Identity {
		return &v
	}).(IdentityPtrOutput)
}

func (o IdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v Identity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o IdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v Identity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type IdentityPtrOutput struct{ *pulumi.OutputState }

func (IdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Identity)(nil)).Elem()
}

func (o IdentityPtrOutput) ToIdentityPtrOutput() IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) ToIdentityPtrOutputWithContext(ctx context.Context) IdentityPtrOutput {
	return o
}

func (o IdentityPtrOutput) Elem() IdentityOutput {
	return o.ApplyT(func(v *Identity) Identity {
		if v != nil {
			return *v
		}
		var ret Identity
		return ret
	}).(IdentityOutput)
}

func (o IdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *Identity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o IdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *Identity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type IdentityResponse struct {
	PrincipalId            string                                            `pulumi:"principalId"`
	TenantId               string                                            `pulumi:"tenantId"`
	Type                   *string                                           `pulumi:"type"`
	UserAssignedIdentities map[string]IdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}

type IdentityResponseOutput struct{ *pulumi.OutputState }

func (IdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponse)(nil)).Elem()
}

func (o IdentityResponseOutput) ToIdentityResponseOutput() IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) ToIdentityResponseOutputWithContext(ctx context.Context) IdentityResponseOutput {
	return o
}

func (o IdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o IdentityResponseOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityResponse)(nil)).Elem()
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutput() IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) ToIdentityResponsePtrOutputWithContext(ctx context.Context) IdentityResponsePtrOutput {
	return o
}

func (o IdentityResponsePtrOutput) Elem() IdentityResponseOutput {
	return o.ApplyT(func(v *IdentityResponse) IdentityResponse {
		if v != nil {
			return *v
		}
		var ret IdentityResponse
		return ret
	}).(IdentityResponseOutput)
}

func (o IdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o IdentityResponsePtrOutput) UserAssignedIdentities() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *IdentityResponse) map[string]IdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(IdentityResponseUserAssignedIdentitiesMapOutput)
}

type IdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type IdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutput() IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ToIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o IdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type IdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (IdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]IdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutput() IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) ToIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) IdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o IdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) IdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) IdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]IdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(IdentityResponseUserAssignedIdentitiesOutput)
}

type KeyResponse struct {
	KeyName     string `pulumi:"keyName"`
	Permissions string `pulumi:"permissions"`
	Value       string `pulumi:"value"`
}

type KeyVaultProperties struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVersion  *string `pulumi:"keyVersion"`
	KeyvaultUri *string `pulumi:"keyvaultUri"`
}





type KeyVaultPropertiesInput interface {
	pulumi.Input

	ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput
	ToKeyVaultPropertiesOutputWithContext(context.Context) KeyVaultPropertiesOutput
}

type KeyVaultPropertiesArgs struct {
	KeyName     pulumi.StringPtrInput `pulumi:"keyName"`
	KeyVersion  pulumi.StringPtrInput `pulumi:"keyVersion"`
	KeyvaultUri pulumi.StringPtrInput `pulumi:"keyvaultUri"`
}

func (KeyVaultPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return i.ToKeyVaultPropertiesOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput)
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i KeyVaultPropertiesArgs) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesOutput).ToKeyVaultPropertiesPtrOutputWithContext(ctx)
}









type KeyVaultPropertiesPtrInput interface {
	pulumi.Input

	ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput
	ToKeyVaultPropertiesPtrOutputWithContext(context.Context) KeyVaultPropertiesPtrOutput
}

type keyVaultPropertiesPtrType KeyVaultPropertiesArgs

func KeyVaultPropertiesPtr(v *KeyVaultPropertiesArgs) KeyVaultPropertiesPtrInput {
	return (*keyVaultPropertiesPtrType)(v)
}

func (*keyVaultPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return i.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (i *keyVaultPropertiesPtrType) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultPropertiesPtrOutput)
}

type KeyVaultPropertiesOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutput() KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesOutputWithContext(ctx context.Context) KeyVaultPropertiesOutput {
	return o
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o.ToKeyVaultPropertiesPtrOutputWithContext(context.Background())
}

func (o KeyVaultPropertiesOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultProperties) *KeyVaultProperties {
		return &v
	}).(KeyVaultPropertiesPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultProperties) *string { return v.KeyvaultUri }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultProperties)(nil)).Elem()
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutput() KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) ToKeyVaultPropertiesPtrOutputWithContext(ctx context.Context) KeyVaultPropertiesPtrOutput {
	return o
}

func (o KeyVaultPropertiesPtrOutput) Elem() KeyVaultPropertiesOutput {
	return o.ApplyT(func(v *KeyVaultProperties) KeyVaultProperties {
		if v != nil {
			return *v
		}
		var ret KeyVaultProperties
		return ret
	}).(KeyVaultPropertiesOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesPtrOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultProperties) *string {
		if v == nil {
			return nil
		}
		return v.KeyvaultUri
	}).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponse struct {
	KeyName     *string `pulumi:"keyName"`
	KeyVersion  *string `pulumi:"keyVersion"`
	KeyvaultUri *string `pulumi:"keyvaultUri"`
}

type KeyVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutput() KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) ToKeyVaultPropertiesResponseOutputWithContext(ctx context.Context) KeyVaultPropertiesResponseOutput {
	return o
}

func (o KeyVaultPropertiesResponseOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyVersion }).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponseOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultPropertiesResponse) *string { return v.KeyvaultUri }).(pulumi.StringPtrOutput)
}

type KeyVaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultPropertiesResponse)(nil)).Elem()
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutput() KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) ToKeyVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) KeyVaultPropertiesResponsePtrOutput {
	return o
}

func (o KeyVaultPropertiesResponsePtrOutput) Elem() KeyVaultPropertiesResponseOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) KeyVaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultPropertiesResponse
		return ret
	}).(KeyVaultPropertiesResponseOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVersion
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultPropertiesResponsePtrOutput) KeyvaultUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyvaultUri
	}).(pulumi.StringPtrOutput)
}

type LinuxProperties struct {
	ExcludedPackageNameMasks       []string `pulumi:"excludedPackageNameMasks"`
	IncludedPackageClassifications *string  `pulumi:"includedPackageClassifications"`
	IncludedPackageNameMasks       []string `pulumi:"includedPackageNameMasks"`
	RebootSetting                  *string  `pulumi:"rebootSetting"`
}





type LinuxPropertiesInput interface {
	pulumi.Input

	ToLinuxPropertiesOutput() LinuxPropertiesOutput
	ToLinuxPropertiesOutputWithContext(context.Context) LinuxPropertiesOutput
}

type LinuxPropertiesArgs struct {
	ExcludedPackageNameMasks       pulumi.StringArrayInput `pulumi:"excludedPackageNameMasks"`
	IncludedPackageClassifications pulumi.StringPtrInput   `pulumi:"includedPackageClassifications"`
	IncludedPackageNameMasks       pulumi.StringArrayInput `pulumi:"includedPackageNameMasks"`
	RebootSetting                  pulumi.StringPtrInput   `pulumi:"rebootSetting"`
}

func (LinuxPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProperties)(nil)).Elem()
}

func (i LinuxPropertiesArgs) ToLinuxPropertiesOutput() LinuxPropertiesOutput {
	return i.ToLinuxPropertiesOutputWithContext(context.Background())
}

func (i LinuxPropertiesArgs) ToLinuxPropertiesOutputWithContext(ctx context.Context) LinuxPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxPropertiesOutput)
}

func (i LinuxPropertiesArgs) ToLinuxPropertiesPtrOutput() LinuxPropertiesPtrOutput {
	return i.ToLinuxPropertiesPtrOutputWithContext(context.Background())
}

func (i LinuxPropertiesArgs) ToLinuxPropertiesPtrOutputWithContext(ctx context.Context) LinuxPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxPropertiesOutput).ToLinuxPropertiesPtrOutputWithContext(ctx)
}









type LinuxPropertiesPtrInput interface {
	pulumi.Input

	ToLinuxPropertiesPtrOutput() LinuxPropertiesPtrOutput
	ToLinuxPropertiesPtrOutputWithContext(context.Context) LinuxPropertiesPtrOutput
}

type linuxPropertiesPtrType LinuxPropertiesArgs

func LinuxPropertiesPtr(v *LinuxPropertiesArgs) LinuxPropertiesPtrInput {
	return (*linuxPropertiesPtrType)(v)
}

func (*linuxPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProperties)(nil)).Elem()
}

func (i *linuxPropertiesPtrType) ToLinuxPropertiesPtrOutput() LinuxPropertiesPtrOutput {
	return i.ToLinuxPropertiesPtrOutputWithContext(context.Background())
}

func (i *linuxPropertiesPtrType) ToLinuxPropertiesPtrOutputWithContext(ctx context.Context) LinuxPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LinuxPropertiesPtrOutput)
}

type LinuxPropertiesOutput struct{ *pulumi.OutputState }

func (LinuxPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxProperties)(nil)).Elem()
}

func (o LinuxPropertiesOutput) ToLinuxPropertiesOutput() LinuxPropertiesOutput {
	return o
}

func (o LinuxPropertiesOutput) ToLinuxPropertiesOutputWithContext(ctx context.Context) LinuxPropertiesOutput {
	return o
}

func (o LinuxPropertiesOutput) ToLinuxPropertiesPtrOutput() LinuxPropertiesPtrOutput {
	return o.ToLinuxPropertiesPtrOutputWithContext(context.Background())
}

func (o LinuxPropertiesOutput) ToLinuxPropertiesPtrOutputWithContext(ctx context.Context) LinuxPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LinuxProperties) *LinuxProperties {
		return &v
	}).(LinuxPropertiesPtrOutput)
}

func (o LinuxPropertiesOutput) ExcludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LinuxProperties) []string { return v.ExcludedPackageNameMasks }).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesOutput) IncludedPackageClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxProperties) *string { return v.IncludedPackageClassifications }).(pulumi.StringPtrOutput)
}

func (o LinuxPropertiesOutput) IncludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LinuxProperties) []string { return v.IncludedPackageNameMasks }).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxProperties) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type LinuxPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LinuxPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxProperties)(nil)).Elem()
}

func (o LinuxPropertiesPtrOutput) ToLinuxPropertiesPtrOutput() LinuxPropertiesPtrOutput {
	return o
}

func (o LinuxPropertiesPtrOutput) ToLinuxPropertiesPtrOutputWithContext(ctx context.Context) LinuxPropertiesPtrOutput {
	return o
}

func (o LinuxPropertiesPtrOutput) Elem() LinuxPropertiesOutput {
	return o.ApplyT(func(v *LinuxProperties) LinuxProperties {
		if v != nil {
			return *v
		}
		var ret LinuxProperties
		return ret
	}).(LinuxPropertiesOutput)
}

func (o LinuxPropertiesPtrOutput) ExcludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LinuxProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPackageNameMasks
	}).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesPtrOutput) IncludedPackageClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxProperties) *string {
		if v == nil {
			return nil
		}
		return v.IncludedPackageClassifications
	}).(pulumi.StringPtrOutput)
}

func (o LinuxPropertiesPtrOutput) IncludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LinuxProperties) []string {
		if v == nil {
			return nil
		}
		return v.IncludedPackageNameMasks
	}).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesPtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxProperties) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

type LinuxPropertiesResponse struct {
	ExcludedPackageNameMasks       []string `pulumi:"excludedPackageNameMasks"`
	IncludedPackageClassifications *string  `pulumi:"includedPackageClassifications"`
	IncludedPackageNameMasks       []string `pulumi:"includedPackageNameMasks"`
	RebootSetting                  *string  `pulumi:"rebootSetting"`
}

type LinuxPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LinuxPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LinuxPropertiesResponse)(nil)).Elem()
}

func (o LinuxPropertiesResponseOutput) ToLinuxPropertiesResponseOutput() LinuxPropertiesResponseOutput {
	return o
}

func (o LinuxPropertiesResponseOutput) ToLinuxPropertiesResponseOutputWithContext(ctx context.Context) LinuxPropertiesResponseOutput {
	return o
}

func (o LinuxPropertiesResponseOutput) ExcludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LinuxPropertiesResponse) []string { return v.ExcludedPackageNameMasks }).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesResponseOutput) IncludedPackageClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxPropertiesResponse) *string { return v.IncludedPackageClassifications }).(pulumi.StringPtrOutput)
}

func (o LinuxPropertiesResponseOutput) IncludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LinuxPropertiesResponse) []string { return v.IncludedPackageNameMasks }).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesResponseOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LinuxPropertiesResponse) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type LinuxPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LinuxPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LinuxPropertiesResponse)(nil)).Elem()
}

func (o LinuxPropertiesResponsePtrOutput) ToLinuxPropertiesResponsePtrOutput() LinuxPropertiesResponsePtrOutput {
	return o
}

func (o LinuxPropertiesResponsePtrOutput) ToLinuxPropertiesResponsePtrOutputWithContext(ctx context.Context) LinuxPropertiesResponsePtrOutput {
	return o
}

func (o LinuxPropertiesResponsePtrOutput) Elem() LinuxPropertiesResponseOutput {
	return o.ApplyT(func(v *LinuxPropertiesResponse) LinuxPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LinuxPropertiesResponse
		return ret
	}).(LinuxPropertiesResponseOutput)
}

func (o LinuxPropertiesResponsePtrOutput) ExcludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LinuxPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedPackageNameMasks
	}).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesResponsePtrOutput) IncludedPackageClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IncludedPackageClassifications
	}).(pulumi.StringPtrOutput)
}

func (o LinuxPropertiesResponsePtrOutput) IncludedPackageNameMasks() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *LinuxPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.IncludedPackageNameMasks
	}).(pulumi.StringArrayOutput)
}

func (o LinuxPropertiesResponsePtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LinuxPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

type ModuleErrorInfoResponse struct {
	Code    *string `pulumi:"code"`
	Message *string `pulumi:"message"`
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

type NonAzureQueryProperties struct {
	FunctionAlias *string `pulumi:"functionAlias"`
	WorkspaceId   *string `pulumi:"workspaceId"`
}





type NonAzureQueryPropertiesInput interface {
	pulumi.Input

	ToNonAzureQueryPropertiesOutput() NonAzureQueryPropertiesOutput
	ToNonAzureQueryPropertiesOutputWithContext(context.Context) NonAzureQueryPropertiesOutput
}

type NonAzureQueryPropertiesArgs struct {
	FunctionAlias pulumi.StringPtrInput `pulumi:"functionAlias"`
	WorkspaceId   pulumi.StringPtrInput `pulumi:"workspaceId"`
}

func (NonAzureQueryPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NonAzureQueryProperties)(nil)).Elem()
}

func (i NonAzureQueryPropertiesArgs) ToNonAzureQueryPropertiesOutput() NonAzureQueryPropertiesOutput {
	return i.ToNonAzureQueryPropertiesOutputWithContext(context.Background())
}

func (i NonAzureQueryPropertiesArgs) ToNonAzureQueryPropertiesOutputWithContext(ctx context.Context) NonAzureQueryPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonAzureQueryPropertiesOutput)
}





type NonAzureQueryPropertiesArrayInput interface {
	pulumi.Input

	ToNonAzureQueryPropertiesArrayOutput() NonAzureQueryPropertiesArrayOutput
	ToNonAzureQueryPropertiesArrayOutputWithContext(context.Context) NonAzureQueryPropertiesArrayOutput
}

type NonAzureQueryPropertiesArray []NonAzureQueryPropertiesInput

func (NonAzureQueryPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonAzureQueryProperties)(nil)).Elem()
}

func (i NonAzureQueryPropertiesArray) ToNonAzureQueryPropertiesArrayOutput() NonAzureQueryPropertiesArrayOutput {
	return i.ToNonAzureQueryPropertiesArrayOutputWithContext(context.Background())
}

func (i NonAzureQueryPropertiesArray) ToNonAzureQueryPropertiesArrayOutputWithContext(ctx context.Context) NonAzureQueryPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NonAzureQueryPropertiesArrayOutput)
}

type NonAzureQueryPropertiesOutput struct{ *pulumi.OutputState }

func (NonAzureQueryPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonAzureQueryProperties)(nil)).Elem()
}

func (o NonAzureQueryPropertiesOutput) ToNonAzureQueryPropertiesOutput() NonAzureQueryPropertiesOutput {
	return o
}

func (o NonAzureQueryPropertiesOutput) ToNonAzureQueryPropertiesOutputWithContext(ctx context.Context) NonAzureQueryPropertiesOutput {
	return o
}

func (o NonAzureQueryPropertiesOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonAzureQueryProperties) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

func (o NonAzureQueryPropertiesOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonAzureQueryProperties) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

type NonAzureQueryPropertiesArrayOutput struct{ *pulumi.OutputState }

func (NonAzureQueryPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonAzureQueryProperties)(nil)).Elem()
}

func (o NonAzureQueryPropertiesArrayOutput) ToNonAzureQueryPropertiesArrayOutput() NonAzureQueryPropertiesArrayOutput {
	return o
}

func (o NonAzureQueryPropertiesArrayOutput) ToNonAzureQueryPropertiesArrayOutputWithContext(ctx context.Context) NonAzureQueryPropertiesArrayOutput {
	return o
}

func (o NonAzureQueryPropertiesArrayOutput) Index(i pulumi.IntInput) NonAzureQueryPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NonAzureQueryProperties {
		return vs[0].([]NonAzureQueryProperties)[vs[1].(int)]
	}).(NonAzureQueryPropertiesOutput)
}

type NonAzureQueryPropertiesResponse struct {
	FunctionAlias *string `pulumi:"functionAlias"`
	WorkspaceId   *string `pulumi:"workspaceId"`
}

type NonAzureQueryPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NonAzureQueryPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NonAzureQueryPropertiesResponse)(nil)).Elem()
}

func (o NonAzureQueryPropertiesResponseOutput) ToNonAzureQueryPropertiesResponseOutput() NonAzureQueryPropertiesResponseOutput {
	return o
}

func (o NonAzureQueryPropertiesResponseOutput) ToNonAzureQueryPropertiesResponseOutputWithContext(ctx context.Context) NonAzureQueryPropertiesResponseOutput {
	return o
}

func (o NonAzureQueryPropertiesResponseOutput) FunctionAlias() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonAzureQueryPropertiesResponse) *string { return v.FunctionAlias }).(pulumi.StringPtrOutput)
}

func (o NonAzureQueryPropertiesResponseOutput) WorkspaceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NonAzureQueryPropertiesResponse) *string { return v.WorkspaceId }).(pulumi.StringPtrOutput)
}

type NonAzureQueryPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (NonAzureQueryPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NonAzureQueryPropertiesResponse)(nil)).Elem()
}

func (o NonAzureQueryPropertiesResponseArrayOutput) ToNonAzureQueryPropertiesResponseArrayOutput() NonAzureQueryPropertiesResponseArrayOutput {
	return o
}

func (o NonAzureQueryPropertiesResponseArrayOutput) ToNonAzureQueryPropertiesResponseArrayOutputWithContext(ctx context.Context) NonAzureQueryPropertiesResponseArrayOutput {
	return o
}

func (o NonAzureQueryPropertiesResponseArrayOutput) Index(i pulumi.IntInput) NonAzureQueryPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NonAzureQueryPropertiesResponse {
		return vs[0].([]NonAzureQueryPropertiesResponse)[vs[1].(int)]
	}).(NonAzureQueryPropertiesResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	Id                                string                                             `pulumi:"id"`
	Name                              string                                             `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	Type                              string                                             `pulumi:"type"`
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string  `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationProperty struct {
	Name *string `pulumi:"name"`
}





type RunAsCredentialAssociationPropertyInput interface {
	pulumi.Input

	ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput
	ToRunAsCredentialAssociationPropertyOutputWithContext(context.Context) RunAsCredentialAssociationPropertyOutput
}

type RunAsCredentialAssociationPropertyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (RunAsCredentialAssociationPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput {
	return i.ToRunAsCredentialAssociationPropertyOutputWithContext(context.Background())
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyOutput)
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i RunAsCredentialAssociationPropertyArgs) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyOutput).ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx)
}









type RunAsCredentialAssociationPropertyPtrInput interface {
	pulumi.Input

	ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput
	ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Context) RunAsCredentialAssociationPropertyPtrOutput
}

type runAsCredentialAssociationPropertyPtrType RunAsCredentialAssociationPropertyArgs

func RunAsCredentialAssociationPropertyPtr(v *RunAsCredentialAssociationPropertyArgs) RunAsCredentialAssociationPropertyPtrInput {
	return (*runAsCredentialAssociationPropertyPtrType)(v)
}

func (*runAsCredentialAssociationPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (i *runAsCredentialAssociationPropertyPtrType) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return i.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (i *runAsCredentialAssociationPropertyPtrType) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RunAsCredentialAssociationPropertyPtrOutput)
}

type RunAsCredentialAssociationPropertyOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyOutput() RunAsCredentialAssociationPropertyOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return o.ToRunAsCredentialAssociationPropertyPtrOutputWithContext(context.Background())
}

func (o RunAsCredentialAssociationPropertyOutput) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RunAsCredentialAssociationProperty) *RunAsCredentialAssociationProperty {
		return &v
	}).(RunAsCredentialAssociationPropertyPtrOutput)
}

func (o RunAsCredentialAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsCredentialAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyPtrOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationProperty)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToRunAsCredentialAssociationPropertyPtrOutput() RunAsCredentialAssociationPropertyPtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyPtrOutput) ToRunAsCredentialAssociationPropertyPtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyPtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyPtrOutput) Elem() RunAsCredentialAssociationPropertyOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationProperty) RunAsCredentialAssociationProperty {
		if v != nil {
			return *v
		}
		var ret RunAsCredentialAssociationProperty
		return ret
	}).(RunAsCredentialAssociationPropertyOutput)
}

func (o RunAsCredentialAssociationPropertyPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationProperty) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
}

type RunAsCredentialAssociationPropertyResponseOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RunAsCredentialAssociationPropertyResponse)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToRunAsCredentialAssociationPropertyResponseOutput() RunAsCredentialAssociationPropertyResponseOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponseOutput) ToRunAsCredentialAssociationPropertyResponseOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyResponseOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RunAsCredentialAssociationPropertyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type RunAsCredentialAssociationPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (RunAsCredentialAssociationPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RunAsCredentialAssociationPropertyResponse)(nil)).Elem()
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToRunAsCredentialAssociationPropertyResponsePtrOutput() RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) ToRunAsCredentialAssociationPropertyResponsePtrOutputWithContext(ctx context.Context) RunAsCredentialAssociationPropertyResponsePtrOutput {
	return o
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) Elem() RunAsCredentialAssociationPropertyResponseOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationPropertyResponse) RunAsCredentialAssociationPropertyResponse {
		if v != nil {
			return *v
		}
		var ret RunAsCredentialAssociationPropertyResponse
		return ret
	}).(RunAsCredentialAssociationPropertyResponseOutput)
}

func (o RunAsCredentialAssociationPropertyResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RunAsCredentialAssociationPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
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

type SUCScheduleProperties struct {
	AdvancedSchedule        *AdvancedSchedule `pulumi:"advancedSchedule"`
	CreationTime            *string           `pulumi:"creationTime"`
	Description             *string           `pulumi:"description"`
	ExpiryTime              *string           `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes *float64          `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               *string           `pulumi:"frequency"`
	Interval                *float64          `pulumi:"interval"`
	IsEnabled               *bool             `pulumi:"isEnabled"`
	LastModifiedTime        *string           `pulumi:"lastModifiedTime"`
	NextRun                 *string           `pulumi:"nextRun"`
	NextRunOffsetMinutes    *float64          `pulumi:"nextRunOffsetMinutes"`
	StartTime               *string           `pulumi:"startTime"`
	TimeZone                *string           `pulumi:"timeZone"`
}


func (val *SUCScheduleProperties) Defaults() *SUCScheduleProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEnabled) {
		isEnabled_ := false
		tmp.IsEnabled = &isEnabled_
	}
	return &tmp
}





type SUCSchedulePropertiesInput interface {
	pulumi.Input

	ToSUCSchedulePropertiesOutput() SUCSchedulePropertiesOutput
	ToSUCSchedulePropertiesOutputWithContext(context.Context) SUCSchedulePropertiesOutput
}

type SUCSchedulePropertiesArgs struct {
	AdvancedSchedule        AdvancedSchedulePtrInput `pulumi:"advancedSchedule"`
	CreationTime            pulumi.StringPtrInput    `pulumi:"creationTime"`
	Description             pulumi.StringPtrInput    `pulumi:"description"`
	ExpiryTime              pulumi.StringPtrInput    `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes pulumi.Float64PtrInput   `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               pulumi.StringPtrInput    `pulumi:"frequency"`
	Interval                pulumi.Float64PtrInput   `pulumi:"interval"`
	IsEnabled               pulumi.BoolPtrInput      `pulumi:"isEnabled"`
	LastModifiedTime        pulumi.StringPtrInput    `pulumi:"lastModifiedTime"`
	NextRun                 pulumi.StringPtrInput    `pulumi:"nextRun"`
	NextRunOffsetMinutes    pulumi.Float64PtrInput   `pulumi:"nextRunOffsetMinutes"`
	StartTime               pulumi.StringPtrInput    `pulumi:"startTime"`
	TimeZone                pulumi.StringPtrInput    `pulumi:"timeZone"`
}

func (SUCSchedulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SUCScheduleProperties)(nil)).Elem()
}

func (i SUCSchedulePropertiesArgs) ToSUCSchedulePropertiesOutput() SUCSchedulePropertiesOutput {
	return i.ToSUCSchedulePropertiesOutputWithContext(context.Background())
}

func (i SUCSchedulePropertiesArgs) ToSUCSchedulePropertiesOutputWithContext(ctx context.Context) SUCSchedulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SUCSchedulePropertiesOutput)
}

type SUCSchedulePropertiesOutput struct{ *pulumi.OutputState }

func (SUCSchedulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SUCScheduleProperties)(nil)).Elem()
}

func (o SUCSchedulePropertiesOutput) ToSUCSchedulePropertiesOutput() SUCSchedulePropertiesOutput {
	return o
}

func (o SUCSchedulePropertiesOutput) ToSUCSchedulePropertiesOutputWithContext(ctx context.Context) SUCSchedulePropertiesOutput {
	return o
}

func (o SUCSchedulePropertiesOutput) AdvancedSchedule() AdvancedSchedulePtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *AdvancedSchedule { return v.AdvancedSchedule }).(AdvancedSchedulePtrOutput)
}

func (o SUCSchedulePropertiesOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) ExpiryTimeOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *float64 { return v.ExpiryTimeOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SUCSchedulePropertiesOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) Interval() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *float64 { return v.Interval }).(pulumi.Float64PtrOutput)
}

func (o SUCSchedulePropertiesOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SUCSchedulePropertiesOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) NextRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.NextRun }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) NextRunOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *float64 { return v.NextRunOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SUCSchedulePropertiesOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCScheduleProperties) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type SUCSchedulePropertiesResponse struct {
	AdvancedSchedule        *AdvancedScheduleResponse `pulumi:"advancedSchedule"`
	CreationTime            *string                   `pulumi:"creationTime"`
	Description             *string                   `pulumi:"description"`
	ExpiryTime              *string                   `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes *float64                  `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               *string                   `pulumi:"frequency"`
	Interval                *float64                  `pulumi:"interval"`
	IsEnabled               *bool                     `pulumi:"isEnabled"`
	LastModifiedTime        *string                   `pulumi:"lastModifiedTime"`
	NextRun                 *string                   `pulumi:"nextRun"`
	NextRunOffsetMinutes    *float64                  `pulumi:"nextRunOffsetMinutes"`
	StartTime               *string                   `pulumi:"startTime"`
	StartTimeOffsetMinutes  float64                   `pulumi:"startTimeOffsetMinutes"`
	TimeZone                *string                   `pulumi:"timeZone"`
}


func (val *SUCSchedulePropertiesResponse) Defaults() *SUCSchedulePropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEnabled) {
		isEnabled_ := false
		tmp.IsEnabled = &isEnabled_
	}
	return &tmp
}

type SUCSchedulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SUCSchedulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SUCSchedulePropertiesResponse)(nil)).Elem()
}

func (o SUCSchedulePropertiesResponseOutput) ToSUCSchedulePropertiesResponseOutput() SUCSchedulePropertiesResponseOutput {
	return o
}

func (o SUCSchedulePropertiesResponseOutput) ToSUCSchedulePropertiesResponseOutputWithContext(ctx context.Context) SUCSchedulePropertiesResponseOutput {
	return o
}

func (o SUCSchedulePropertiesResponseOutput) AdvancedSchedule() AdvancedScheduleResponsePtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *AdvancedScheduleResponse { return v.AdvancedSchedule }).(AdvancedScheduleResponsePtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) ExpiryTimeOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *float64 { return v.ExpiryTimeOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) Interval() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *float64 { return v.Interval }).(pulumi.Float64PtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) NextRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.NextRun }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) NextRunOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *float64 { return v.NextRunOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o SUCSchedulePropertiesResponseOutput) StartTimeOffsetMinutes() pulumi.Float64Output {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) float64 { return v.StartTimeOffsetMinutes }).(pulumi.Float64Output)
}

func (o SUCSchedulePropertiesResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SUCSchedulePropertiesResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
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

func (o ScheduleAssociationPropertyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleAssociationProperty) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScheduleAssociationPropertyResponse struct {
	Name *string `pulumi:"name"`
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

type SoftwareUpdateConfigurationTasks struct {
	PostTask *TaskProperties `pulumi:"postTask"`
	PreTask  *TaskProperties `pulumi:"preTask"`
}





type SoftwareUpdateConfigurationTasksInput interface {
	pulumi.Input

	ToSoftwareUpdateConfigurationTasksOutput() SoftwareUpdateConfigurationTasksOutput
	ToSoftwareUpdateConfigurationTasksOutputWithContext(context.Context) SoftwareUpdateConfigurationTasksOutput
}

type SoftwareUpdateConfigurationTasksArgs struct {
	PostTask TaskPropertiesPtrInput `pulumi:"postTask"`
	PreTask  TaskPropertiesPtrInput `pulumi:"preTask"`
}

func (SoftwareUpdateConfigurationTasksArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareUpdateConfigurationTasks)(nil)).Elem()
}

func (i SoftwareUpdateConfigurationTasksArgs) ToSoftwareUpdateConfigurationTasksOutput() SoftwareUpdateConfigurationTasksOutput {
	return i.ToSoftwareUpdateConfigurationTasksOutputWithContext(context.Background())
}

func (i SoftwareUpdateConfigurationTasksArgs) ToSoftwareUpdateConfigurationTasksOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareUpdateConfigurationTasksOutput)
}

func (i SoftwareUpdateConfigurationTasksArgs) ToSoftwareUpdateConfigurationTasksPtrOutput() SoftwareUpdateConfigurationTasksPtrOutput {
	return i.ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(context.Background())
}

func (i SoftwareUpdateConfigurationTasksArgs) ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareUpdateConfigurationTasksOutput).ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(ctx)
}









type SoftwareUpdateConfigurationTasksPtrInput interface {
	pulumi.Input

	ToSoftwareUpdateConfigurationTasksPtrOutput() SoftwareUpdateConfigurationTasksPtrOutput
	ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(context.Context) SoftwareUpdateConfigurationTasksPtrOutput
}

type softwareUpdateConfigurationTasksPtrType SoftwareUpdateConfigurationTasksArgs

func SoftwareUpdateConfigurationTasksPtr(v *SoftwareUpdateConfigurationTasksArgs) SoftwareUpdateConfigurationTasksPtrInput {
	return (*softwareUpdateConfigurationTasksPtrType)(v)
}

func (*softwareUpdateConfigurationTasksPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationTasks)(nil)).Elem()
}

func (i *softwareUpdateConfigurationTasksPtrType) ToSoftwareUpdateConfigurationTasksPtrOutput() SoftwareUpdateConfigurationTasksPtrOutput {
	return i.ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(context.Background())
}

func (i *softwareUpdateConfigurationTasksPtrType) ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SoftwareUpdateConfigurationTasksPtrOutput)
}

type SoftwareUpdateConfigurationTasksOutput struct{ *pulumi.OutputState }

func (SoftwareUpdateConfigurationTasksOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareUpdateConfigurationTasks)(nil)).Elem()
}

func (o SoftwareUpdateConfigurationTasksOutput) ToSoftwareUpdateConfigurationTasksOutput() SoftwareUpdateConfigurationTasksOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksOutput) ToSoftwareUpdateConfigurationTasksOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksOutput) ToSoftwareUpdateConfigurationTasksPtrOutput() SoftwareUpdateConfigurationTasksPtrOutput {
	return o.ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(context.Background())
}

func (o SoftwareUpdateConfigurationTasksOutput) ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SoftwareUpdateConfigurationTasks) *SoftwareUpdateConfigurationTasks {
		return &v
	}).(SoftwareUpdateConfigurationTasksPtrOutput)
}

func (o SoftwareUpdateConfigurationTasksOutput) PostTask() TaskPropertiesPtrOutput {
	return o.ApplyT(func(v SoftwareUpdateConfigurationTasks) *TaskProperties { return v.PostTask }).(TaskPropertiesPtrOutput)
}

func (o SoftwareUpdateConfigurationTasksOutput) PreTask() TaskPropertiesPtrOutput {
	return o.ApplyT(func(v SoftwareUpdateConfigurationTasks) *TaskProperties { return v.PreTask }).(TaskPropertiesPtrOutput)
}

type SoftwareUpdateConfigurationTasksPtrOutput struct{ *pulumi.OutputState }

func (SoftwareUpdateConfigurationTasksPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationTasks)(nil)).Elem()
}

func (o SoftwareUpdateConfigurationTasksPtrOutput) ToSoftwareUpdateConfigurationTasksPtrOutput() SoftwareUpdateConfigurationTasksPtrOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksPtrOutput) ToSoftwareUpdateConfigurationTasksPtrOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksPtrOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksPtrOutput) Elem() SoftwareUpdateConfigurationTasksOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationTasks) SoftwareUpdateConfigurationTasks {
		if v != nil {
			return *v
		}
		var ret SoftwareUpdateConfigurationTasks
		return ret
	}).(SoftwareUpdateConfigurationTasksOutput)
}

func (o SoftwareUpdateConfigurationTasksPtrOutput) PostTask() TaskPropertiesPtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationTasks) *TaskProperties {
		if v == nil {
			return nil
		}
		return v.PostTask
	}).(TaskPropertiesPtrOutput)
}

func (o SoftwareUpdateConfigurationTasksPtrOutput) PreTask() TaskPropertiesPtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationTasks) *TaskProperties {
		if v == nil {
			return nil
		}
		return v.PreTask
	}).(TaskPropertiesPtrOutput)
}

type SoftwareUpdateConfigurationTasksResponse struct {
	PostTask *TaskPropertiesResponse `pulumi:"postTask"`
	PreTask  *TaskPropertiesResponse `pulumi:"preTask"`
}

type SoftwareUpdateConfigurationTasksResponseOutput struct{ *pulumi.OutputState }

func (SoftwareUpdateConfigurationTasksResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SoftwareUpdateConfigurationTasksResponse)(nil)).Elem()
}

func (o SoftwareUpdateConfigurationTasksResponseOutput) ToSoftwareUpdateConfigurationTasksResponseOutput() SoftwareUpdateConfigurationTasksResponseOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksResponseOutput) ToSoftwareUpdateConfigurationTasksResponseOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksResponseOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksResponseOutput) PostTask() TaskPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SoftwareUpdateConfigurationTasksResponse) *TaskPropertiesResponse { return v.PostTask }).(TaskPropertiesResponsePtrOutput)
}

func (o SoftwareUpdateConfigurationTasksResponseOutput) PreTask() TaskPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SoftwareUpdateConfigurationTasksResponse) *TaskPropertiesResponse { return v.PreTask }).(TaskPropertiesResponsePtrOutput)
}

type SoftwareUpdateConfigurationTasksResponsePtrOutput struct{ *pulumi.OutputState }

func (SoftwareUpdateConfigurationTasksResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SoftwareUpdateConfigurationTasksResponse)(nil)).Elem()
}

func (o SoftwareUpdateConfigurationTasksResponsePtrOutput) ToSoftwareUpdateConfigurationTasksResponsePtrOutput() SoftwareUpdateConfigurationTasksResponsePtrOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksResponsePtrOutput) ToSoftwareUpdateConfigurationTasksResponsePtrOutputWithContext(ctx context.Context) SoftwareUpdateConfigurationTasksResponsePtrOutput {
	return o
}

func (o SoftwareUpdateConfigurationTasksResponsePtrOutput) Elem() SoftwareUpdateConfigurationTasksResponseOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationTasksResponse) SoftwareUpdateConfigurationTasksResponse {
		if v != nil {
			return *v
		}
		var ret SoftwareUpdateConfigurationTasksResponse
		return ret
	}).(SoftwareUpdateConfigurationTasksResponseOutput)
}

func (o SoftwareUpdateConfigurationTasksResponsePtrOutput) PostTask() TaskPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationTasksResponse) *TaskPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.PostTask
	}).(TaskPropertiesResponsePtrOutput)
}

func (o SoftwareUpdateConfigurationTasksResponsePtrOutput) PreTask() TaskPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *SoftwareUpdateConfigurationTasksResponse) *TaskPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.PreTask
	}).(TaskPropertiesResponsePtrOutput)
}

type SourceControlSecurityTokenProperties struct {
	AccessToken  *string `pulumi:"accessToken"`
	RefreshToken *string `pulumi:"refreshToken"`
	TokenType    *string `pulumi:"tokenType"`
}





type SourceControlSecurityTokenPropertiesInput interface {
	pulumi.Input

	ToSourceControlSecurityTokenPropertiesOutput() SourceControlSecurityTokenPropertiesOutput
	ToSourceControlSecurityTokenPropertiesOutputWithContext(context.Context) SourceControlSecurityTokenPropertiesOutput
}

type SourceControlSecurityTokenPropertiesArgs struct {
	AccessToken  pulumi.StringPtrInput `pulumi:"accessToken"`
	RefreshToken pulumi.StringPtrInput `pulumi:"refreshToken"`
	TokenType    pulumi.StringPtrInput `pulumi:"tokenType"`
}

func (SourceControlSecurityTokenPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlSecurityTokenProperties)(nil)).Elem()
}

func (i SourceControlSecurityTokenPropertiesArgs) ToSourceControlSecurityTokenPropertiesOutput() SourceControlSecurityTokenPropertiesOutput {
	return i.ToSourceControlSecurityTokenPropertiesOutputWithContext(context.Background())
}

func (i SourceControlSecurityTokenPropertiesArgs) ToSourceControlSecurityTokenPropertiesOutputWithContext(ctx context.Context) SourceControlSecurityTokenPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlSecurityTokenPropertiesOutput)
}

func (i SourceControlSecurityTokenPropertiesArgs) ToSourceControlSecurityTokenPropertiesPtrOutput() SourceControlSecurityTokenPropertiesPtrOutput {
	return i.ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(context.Background())
}

func (i SourceControlSecurityTokenPropertiesArgs) ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(ctx context.Context) SourceControlSecurityTokenPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlSecurityTokenPropertiesOutput).ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(ctx)
}









type SourceControlSecurityTokenPropertiesPtrInput interface {
	pulumi.Input

	ToSourceControlSecurityTokenPropertiesPtrOutput() SourceControlSecurityTokenPropertiesPtrOutput
	ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(context.Context) SourceControlSecurityTokenPropertiesPtrOutput
}

type sourceControlSecurityTokenPropertiesPtrType SourceControlSecurityTokenPropertiesArgs

func SourceControlSecurityTokenPropertiesPtr(v *SourceControlSecurityTokenPropertiesArgs) SourceControlSecurityTokenPropertiesPtrInput {
	return (*sourceControlSecurityTokenPropertiesPtrType)(v)
}

func (*sourceControlSecurityTokenPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlSecurityTokenProperties)(nil)).Elem()
}

func (i *sourceControlSecurityTokenPropertiesPtrType) ToSourceControlSecurityTokenPropertiesPtrOutput() SourceControlSecurityTokenPropertiesPtrOutput {
	return i.ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(context.Background())
}

func (i *sourceControlSecurityTokenPropertiesPtrType) ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(ctx context.Context) SourceControlSecurityTokenPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SourceControlSecurityTokenPropertiesPtrOutput)
}

type SourceControlSecurityTokenPropertiesOutput struct{ *pulumi.OutputState }

func (SourceControlSecurityTokenPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SourceControlSecurityTokenProperties)(nil)).Elem()
}

func (o SourceControlSecurityTokenPropertiesOutput) ToSourceControlSecurityTokenPropertiesOutput() SourceControlSecurityTokenPropertiesOutput {
	return o
}

func (o SourceControlSecurityTokenPropertiesOutput) ToSourceControlSecurityTokenPropertiesOutputWithContext(ctx context.Context) SourceControlSecurityTokenPropertiesOutput {
	return o
}

func (o SourceControlSecurityTokenPropertiesOutput) ToSourceControlSecurityTokenPropertiesPtrOutput() SourceControlSecurityTokenPropertiesPtrOutput {
	return o.ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(context.Background())
}

func (o SourceControlSecurityTokenPropertiesOutput) ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(ctx context.Context) SourceControlSecurityTokenPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SourceControlSecurityTokenProperties) *SourceControlSecurityTokenProperties {
		return &v
	}).(SourceControlSecurityTokenPropertiesPtrOutput)
}

func (o SourceControlSecurityTokenPropertiesOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlSecurityTokenProperties) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o SourceControlSecurityTokenPropertiesOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlSecurityTokenProperties) *string { return v.RefreshToken }).(pulumi.StringPtrOutput)
}

func (o SourceControlSecurityTokenPropertiesOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SourceControlSecurityTokenProperties) *string { return v.TokenType }).(pulumi.StringPtrOutput)
}

type SourceControlSecurityTokenPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SourceControlSecurityTokenPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SourceControlSecurityTokenProperties)(nil)).Elem()
}

func (o SourceControlSecurityTokenPropertiesPtrOutput) ToSourceControlSecurityTokenPropertiesPtrOutput() SourceControlSecurityTokenPropertiesPtrOutput {
	return o
}

func (o SourceControlSecurityTokenPropertiesPtrOutput) ToSourceControlSecurityTokenPropertiesPtrOutputWithContext(ctx context.Context) SourceControlSecurityTokenPropertiesPtrOutput {
	return o
}

func (o SourceControlSecurityTokenPropertiesPtrOutput) Elem() SourceControlSecurityTokenPropertiesOutput {
	return o.ApplyT(func(v *SourceControlSecurityTokenProperties) SourceControlSecurityTokenProperties {
		if v != nil {
			return *v
		}
		var ret SourceControlSecurityTokenProperties
		return ret
	}).(SourceControlSecurityTokenPropertiesOutput)
}

func (o SourceControlSecurityTokenPropertiesPtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlSecurityTokenProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlSecurityTokenPropertiesPtrOutput) RefreshToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlSecurityTokenProperties) *string {
		if v == nil {
			return nil
		}
		return v.RefreshToken
	}).(pulumi.StringPtrOutput)
}

func (o SourceControlSecurityTokenPropertiesPtrOutput) TokenType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SourceControlSecurityTokenProperties) *string {
		if v == nil {
			return nil
		}
		return v.TokenType
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

type TagSettingsProperties struct {
	FilterOperator *TagOperators       `pulumi:"filterOperator"`
	Tags           map[string][]string `pulumi:"tags"`
}





type TagSettingsPropertiesInput interface {
	pulumi.Input

	ToTagSettingsPropertiesOutput() TagSettingsPropertiesOutput
	ToTagSettingsPropertiesOutputWithContext(context.Context) TagSettingsPropertiesOutput
}

type TagSettingsPropertiesArgs struct {
	FilterOperator TagOperatorsPtrInput       `pulumi:"filterOperator"`
	Tags           pulumi.StringArrayMapInput `pulumi:"tags"`
}

func (TagSettingsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TagSettingsProperties)(nil)).Elem()
}

func (i TagSettingsPropertiesArgs) ToTagSettingsPropertiesOutput() TagSettingsPropertiesOutput {
	return i.ToTagSettingsPropertiesOutputWithContext(context.Background())
}

func (i TagSettingsPropertiesArgs) ToTagSettingsPropertiesOutputWithContext(ctx context.Context) TagSettingsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagSettingsPropertiesOutput)
}

func (i TagSettingsPropertiesArgs) ToTagSettingsPropertiesPtrOutput() TagSettingsPropertiesPtrOutput {
	return i.ToTagSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (i TagSettingsPropertiesArgs) ToTagSettingsPropertiesPtrOutputWithContext(ctx context.Context) TagSettingsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagSettingsPropertiesOutput).ToTagSettingsPropertiesPtrOutputWithContext(ctx)
}









type TagSettingsPropertiesPtrInput interface {
	pulumi.Input

	ToTagSettingsPropertiesPtrOutput() TagSettingsPropertiesPtrOutput
	ToTagSettingsPropertiesPtrOutputWithContext(context.Context) TagSettingsPropertiesPtrOutput
}

type tagSettingsPropertiesPtrType TagSettingsPropertiesArgs

func TagSettingsPropertiesPtr(v *TagSettingsPropertiesArgs) TagSettingsPropertiesPtrInput {
	return (*tagSettingsPropertiesPtrType)(v)
}

func (*tagSettingsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TagSettingsProperties)(nil)).Elem()
}

func (i *tagSettingsPropertiesPtrType) ToTagSettingsPropertiesPtrOutput() TagSettingsPropertiesPtrOutput {
	return i.ToTagSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (i *tagSettingsPropertiesPtrType) ToTagSettingsPropertiesPtrOutputWithContext(ctx context.Context) TagSettingsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TagSettingsPropertiesPtrOutput)
}

type TagSettingsPropertiesOutput struct{ *pulumi.OutputState }

func (TagSettingsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagSettingsProperties)(nil)).Elem()
}

func (o TagSettingsPropertiesOutput) ToTagSettingsPropertiesOutput() TagSettingsPropertiesOutput {
	return o
}

func (o TagSettingsPropertiesOutput) ToTagSettingsPropertiesOutputWithContext(ctx context.Context) TagSettingsPropertiesOutput {
	return o
}

func (o TagSettingsPropertiesOutput) ToTagSettingsPropertiesPtrOutput() TagSettingsPropertiesPtrOutput {
	return o.ToTagSettingsPropertiesPtrOutputWithContext(context.Background())
}

func (o TagSettingsPropertiesOutput) ToTagSettingsPropertiesPtrOutputWithContext(ctx context.Context) TagSettingsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TagSettingsProperties) *TagSettingsProperties {
		return &v
	}).(TagSettingsPropertiesPtrOutput)
}

func (o TagSettingsPropertiesOutput) FilterOperator() TagOperatorsPtrOutput {
	return o.ApplyT(func(v TagSettingsProperties) *TagOperators { return v.FilterOperator }).(TagOperatorsPtrOutput)
}

func (o TagSettingsPropertiesOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v TagSettingsProperties) map[string][]string { return v.Tags }).(pulumi.StringArrayMapOutput)
}

type TagSettingsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TagSettingsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagSettingsProperties)(nil)).Elem()
}

func (o TagSettingsPropertiesPtrOutput) ToTagSettingsPropertiesPtrOutput() TagSettingsPropertiesPtrOutput {
	return o
}

func (o TagSettingsPropertiesPtrOutput) ToTagSettingsPropertiesPtrOutputWithContext(ctx context.Context) TagSettingsPropertiesPtrOutput {
	return o
}

func (o TagSettingsPropertiesPtrOutput) Elem() TagSettingsPropertiesOutput {
	return o.ApplyT(func(v *TagSettingsProperties) TagSettingsProperties {
		if v != nil {
			return *v
		}
		var ret TagSettingsProperties
		return ret
	}).(TagSettingsPropertiesOutput)
}

func (o TagSettingsPropertiesPtrOutput) FilterOperator() TagOperatorsPtrOutput {
	return o.ApplyT(func(v *TagSettingsProperties) *TagOperators {
		if v == nil {
			return nil
		}
		return v.FilterOperator
	}).(TagOperatorsPtrOutput)
}

func (o TagSettingsPropertiesPtrOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *TagSettingsProperties) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayMapOutput)
}

type TagSettingsPropertiesResponse struct {
	FilterOperator *string             `pulumi:"filterOperator"`
	Tags           map[string][]string `pulumi:"tags"`
}

type TagSettingsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TagSettingsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TagSettingsPropertiesResponse)(nil)).Elem()
}

func (o TagSettingsPropertiesResponseOutput) ToTagSettingsPropertiesResponseOutput() TagSettingsPropertiesResponseOutput {
	return o
}

func (o TagSettingsPropertiesResponseOutput) ToTagSettingsPropertiesResponseOutputWithContext(ctx context.Context) TagSettingsPropertiesResponseOutput {
	return o
}

func (o TagSettingsPropertiesResponseOutput) FilterOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TagSettingsPropertiesResponse) *string { return v.FilterOperator }).(pulumi.StringPtrOutput)
}

func (o TagSettingsPropertiesResponseOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v TagSettingsPropertiesResponse) map[string][]string { return v.Tags }).(pulumi.StringArrayMapOutput)
}

type TagSettingsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TagSettingsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TagSettingsPropertiesResponse)(nil)).Elem()
}

func (o TagSettingsPropertiesResponsePtrOutput) ToTagSettingsPropertiesResponsePtrOutput() TagSettingsPropertiesResponsePtrOutput {
	return o
}

func (o TagSettingsPropertiesResponsePtrOutput) ToTagSettingsPropertiesResponsePtrOutputWithContext(ctx context.Context) TagSettingsPropertiesResponsePtrOutput {
	return o
}

func (o TagSettingsPropertiesResponsePtrOutput) Elem() TagSettingsPropertiesResponseOutput {
	return o.ApplyT(func(v *TagSettingsPropertiesResponse) TagSettingsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TagSettingsPropertiesResponse
		return ret
	}).(TagSettingsPropertiesResponseOutput)
}

func (o TagSettingsPropertiesResponsePtrOutput) FilterOperator() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TagSettingsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.FilterOperator
	}).(pulumi.StringPtrOutput)
}

func (o TagSettingsPropertiesResponsePtrOutput) Tags() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v *TagSettingsPropertiesResponse) map[string][]string {
		if v == nil {
			return nil
		}
		return v.Tags
	}).(pulumi.StringArrayMapOutput)
}

type TargetProperties struct {
	AzureQueries    []AzureQueryProperties    `pulumi:"azureQueries"`
	NonAzureQueries []NonAzureQueryProperties `pulumi:"nonAzureQueries"`
}





type TargetPropertiesInput interface {
	pulumi.Input

	ToTargetPropertiesOutput() TargetPropertiesOutput
	ToTargetPropertiesOutputWithContext(context.Context) TargetPropertiesOutput
}

type TargetPropertiesArgs struct {
	AzureQueries    AzureQueryPropertiesArrayInput    `pulumi:"azureQueries"`
	NonAzureQueries NonAzureQueryPropertiesArrayInput `pulumi:"nonAzureQueries"`
}

func (TargetPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetProperties)(nil)).Elem()
}

func (i TargetPropertiesArgs) ToTargetPropertiesOutput() TargetPropertiesOutput {
	return i.ToTargetPropertiesOutputWithContext(context.Background())
}

func (i TargetPropertiesArgs) ToTargetPropertiesOutputWithContext(ctx context.Context) TargetPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPropertiesOutput)
}

func (i TargetPropertiesArgs) ToTargetPropertiesPtrOutput() TargetPropertiesPtrOutput {
	return i.ToTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i TargetPropertiesArgs) ToTargetPropertiesPtrOutputWithContext(ctx context.Context) TargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPropertiesOutput).ToTargetPropertiesPtrOutputWithContext(ctx)
}









type TargetPropertiesPtrInput interface {
	pulumi.Input

	ToTargetPropertiesPtrOutput() TargetPropertiesPtrOutput
	ToTargetPropertiesPtrOutputWithContext(context.Context) TargetPropertiesPtrOutput
}

type targetPropertiesPtrType TargetPropertiesArgs

func TargetPropertiesPtr(v *TargetPropertiesArgs) TargetPropertiesPtrInput {
	return (*targetPropertiesPtrType)(v)
}

func (*targetPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetProperties)(nil)).Elem()
}

func (i *targetPropertiesPtrType) ToTargetPropertiesPtrOutput() TargetPropertiesPtrOutput {
	return i.ToTargetPropertiesPtrOutputWithContext(context.Background())
}

func (i *targetPropertiesPtrType) ToTargetPropertiesPtrOutputWithContext(ctx context.Context) TargetPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TargetPropertiesPtrOutput)
}

type TargetPropertiesOutput struct{ *pulumi.OutputState }

func (TargetPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetProperties)(nil)).Elem()
}

func (o TargetPropertiesOutput) ToTargetPropertiesOutput() TargetPropertiesOutput {
	return o
}

func (o TargetPropertiesOutput) ToTargetPropertiesOutputWithContext(ctx context.Context) TargetPropertiesOutput {
	return o
}

func (o TargetPropertiesOutput) ToTargetPropertiesPtrOutput() TargetPropertiesPtrOutput {
	return o.ToTargetPropertiesPtrOutputWithContext(context.Background())
}

func (o TargetPropertiesOutput) ToTargetPropertiesPtrOutputWithContext(ctx context.Context) TargetPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TargetProperties) *TargetProperties {
		return &v
	}).(TargetPropertiesPtrOutput)
}

func (o TargetPropertiesOutput) AzureQueries() AzureQueryPropertiesArrayOutput {
	return o.ApplyT(func(v TargetProperties) []AzureQueryProperties { return v.AzureQueries }).(AzureQueryPropertiesArrayOutput)
}

func (o TargetPropertiesOutput) NonAzureQueries() NonAzureQueryPropertiesArrayOutput {
	return o.ApplyT(func(v TargetProperties) []NonAzureQueryProperties { return v.NonAzureQueries }).(NonAzureQueryPropertiesArrayOutput)
}

type TargetPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TargetPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetProperties)(nil)).Elem()
}

func (o TargetPropertiesPtrOutput) ToTargetPropertiesPtrOutput() TargetPropertiesPtrOutput {
	return o
}

func (o TargetPropertiesPtrOutput) ToTargetPropertiesPtrOutputWithContext(ctx context.Context) TargetPropertiesPtrOutput {
	return o
}

func (o TargetPropertiesPtrOutput) Elem() TargetPropertiesOutput {
	return o.ApplyT(func(v *TargetProperties) TargetProperties {
		if v != nil {
			return *v
		}
		var ret TargetProperties
		return ret
	}).(TargetPropertiesOutput)
}

func (o TargetPropertiesPtrOutput) AzureQueries() AzureQueryPropertiesArrayOutput {
	return o.ApplyT(func(v *TargetProperties) []AzureQueryProperties {
		if v == nil {
			return nil
		}
		return v.AzureQueries
	}).(AzureQueryPropertiesArrayOutput)
}

func (o TargetPropertiesPtrOutput) NonAzureQueries() NonAzureQueryPropertiesArrayOutput {
	return o.ApplyT(func(v *TargetProperties) []NonAzureQueryProperties {
		if v == nil {
			return nil
		}
		return v.NonAzureQueries
	}).(NonAzureQueryPropertiesArrayOutput)
}

type TargetPropertiesResponse struct {
	AzureQueries    []AzureQueryPropertiesResponse    `pulumi:"azureQueries"`
	NonAzureQueries []NonAzureQueryPropertiesResponse `pulumi:"nonAzureQueries"`
}

type TargetPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TargetPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TargetPropertiesResponse)(nil)).Elem()
}

func (o TargetPropertiesResponseOutput) ToTargetPropertiesResponseOutput() TargetPropertiesResponseOutput {
	return o
}

func (o TargetPropertiesResponseOutput) ToTargetPropertiesResponseOutputWithContext(ctx context.Context) TargetPropertiesResponseOutput {
	return o
}

func (o TargetPropertiesResponseOutput) AzureQueries() AzureQueryPropertiesResponseArrayOutput {
	return o.ApplyT(func(v TargetPropertiesResponse) []AzureQueryPropertiesResponse { return v.AzureQueries }).(AzureQueryPropertiesResponseArrayOutput)
}

func (o TargetPropertiesResponseOutput) NonAzureQueries() NonAzureQueryPropertiesResponseArrayOutput {
	return o.ApplyT(func(v TargetPropertiesResponse) []NonAzureQueryPropertiesResponse { return v.NonAzureQueries }).(NonAzureQueryPropertiesResponseArrayOutput)
}

type TargetPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TargetPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TargetPropertiesResponse)(nil)).Elem()
}

func (o TargetPropertiesResponsePtrOutput) ToTargetPropertiesResponsePtrOutput() TargetPropertiesResponsePtrOutput {
	return o
}

func (o TargetPropertiesResponsePtrOutput) ToTargetPropertiesResponsePtrOutputWithContext(ctx context.Context) TargetPropertiesResponsePtrOutput {
	return o
}

func (o TargetPropertiesResponsePtrOutput) Elem() TargetPropertiesResponseOutput {
	return o.ApplyT(func(v *TargetPropertiesResponse) TargetPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TargetPropertiesResponse
		return ret
	}).(TargetPropertiesResponseOutput)
}

func (o TargetPropertiesResponsePtrOutput) AzureQueries() AzureQueryPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *TargetPropertiesResponse) []AzureQueryPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.AzureQueries
	}).(AzureQueryPropertiesResponseArrayOutput)
}

func (o TargetPropertiesResponsePtrOutput) NonAzureQueries() NonAzureQueryPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *TargetPropertiesResponse) []NonAzureQueryPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.NonAzureQueries
	}).(NonAzureQueryPropertiesResponseArrayOutput)
}

type TaskProperties struct {
	Parameters map[string]string `pulumi:"parameters"`
	Source     *string           `pulumi:"source"`
}





type TaskPropertiesInput interface {
	pulumi.Input

	ToTaskPropertiesOutput() TaskPropertiesOutput
	ToTaskPropertiesOutputWithContext(context.Context) TaskPropertiesOutput
}

type TaskPropertiesArgs struct {
	Parameters pulumi.StringMapInput `pulumi:"parameters"`
	Source     pulumi.StringPtrInput `pulumi:"source"`
}

func (TaskPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskProperties)(nil)).Elem()
}

func (i TaskPropertiesArgs) ToTaskPropertiesOutput() TaskPropertiesOutput {
	return i.ToTaskPropertiesOutputWithContext(context.Background())
}

func (i TaskPropertiesArgs) ToTaskPropertiesOutputWithContext(ctx context.Context) TaskPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesOutput)
}

func (i TaskPropertiesArgs) ToTaskPropertiesPtrOutput() TaskPropertiesPtrOutput {
	return i.ToTaskPropertiesPtrOutputWithContext(context.Background())
}

func (i TaskPropertiesArgs) ToTaskPropertiesPtrOutputWithContext(ctx context.Context) TaskPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesOutput).ToTaskPropertiesPtrOutputWithContext(ctx)
}









type TaskPropertiesPtrInput interface {
	pulumi.Input

	ToTaskPropertiesPtrOutput() TaskPropertiesPtrOutput
	ToTaskPropertiesPtrOutputWithContext(context.Context) TaskPropertiesPtrOutput
}

type taskPropertiesPtrType TaskPropertiesArgs

func TaskPropertiesPtr(v *TaskPropertiesArgs) TaskPropertiesPtrInput {
	return (*taskPropertiesPtrType)(v)
}

func (*taskPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskProperties)(nil)).Elem()
}

func (i *taskPropertiesPtrType) ToTaskPropertiesPtrOutput() TaskPropertiesPtrOutput {
	return i.ToTaskPropertiesPtrOutputWithContext(context.Background())
}

func (i *taskPropertiesPtrType) ToTaskPropertiesPtrOutputWithContext(ctx context.Context) TaskPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskPropertiesPtrOutput)
}

type TaskPropertiesOutput struct{ *pulumi.OutputState }

func (TaskPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskProperties)(nil)).Elem()
}

func (o TaskPropertiesOutput) ToTaskPropertiesOutput() TaskPropertiesOutput {
	return o
}

func (o TaskPropertiesOutput) ToTaskPropertiesOutputWithContext(ctx context.Context) TaskPropertiesOutput {
	return o
}

func (o TaskPropertiesOutput) ToTaskPropertiesPtrOutput() TaskPropertiesPtrOutput {
	return o.ToTaskPropertiesPtrOutputWithContext(context.Background())
}

func (o TaskPropertiesOutput) ToTaskPropertiesPtrOutputWithContext(ctx context.Context) TaskPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TaskProperties) *TaskProperties {
		return &v
	}).(TaskPropertiesPtrOutput)
}

func (o TaskPropertiesOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v TaskProperties) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o TaskPropertiesOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskProperties) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type TaskPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TaskPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskProperties)(nil)).Elem()
}

func (o TaskPropertiesPtrOutput) ToTaskPropertiesPtrOutput() TaskPropertiesPtrOutput {
	return o
}

func (o TaskPropertiesPtrOutput) ToTaskPropertiesPtrOutputWithContext(ctx context.Context) TaskPropertiesPtrOutput {
	return o
}

func (o TaskPropertiesPtrOutput) Elem() TaskPropertiesOutput {
	return o.ApplyT(func(v *TaskProperties) TaskProperties {
		if v != nil {
			return *v
		}
		var ret TaskProperties
		return ret
	}).(TaskPropertiesOutput)
}

func (o TaskPropertiesPtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TaskProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

func (o TaskPropertiesPtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskProperties) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

type TaskPropertiesResponse struct {
	Parameters map[string]string `pulumi:"parameters"`
	Source     *string           `pulumi:"source"`
}

type TaskPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TaskPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TaskPropertiesResponse)(nil)).Elem()
}

func (o TaskPropertiesResponseOutput) ToTaskPropertiesResponseOutput() TaskPropertiesResponseOutput {
	return o
}

func (o TaskPropertiesResponseOutput) ToTaskPropertiesResponseOutputWithContext(ctx context.Context) TaskPropertiesResponseOutput {
	return o
}

func (o TaskPropertiesResponseOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v TaskPropertiesResponse) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o TaskPropertiesResponseOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TaskPropertiesResponse) *string { return v.Source }).(pulumi.StringPtrOutput)
}

type TaskPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TaskPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskPropertiesResponse)(nil)).Elem()
}

func (o TaskPropertiesResponsePtrOutput) ToTaskPropertiesResponsePtrOutput() TaskPropertiesResponsePtrOutput {
	return o
}

func (o TaskPropertiesResponsePtrOutput) ToTaskPropertiesResponsePtrOutputWithContext(ctx context.Context) TaskPropertiesResponsePtrOutput {
	return o
}

func (o TaskPropertiesResponsePtrOutput) Elem() TaskPropertiesResponseOutput {
	return o.ApplyT(func(v *TaskPropertiesResponse) TaskPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TaskPropertiesResponse
		return ret
	}).(TaskPropertiesResponseOutput)
}

func (o TaskPropertiesResponsePtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *TaskPropertiesResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

func (o TaskPropertiesResponsePtrOutput) Source() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TaskPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Source
	}).(pulumi.StringPtrOutput)
}

type UpdateConfiguration struct {
	AzureVirtualMachines  []string            `pulumi:"azureVirtualMachines"`
	Duration              *string             `pulumi:"duration"`
	Linux                 *LinuxProperties    `pulumi:"linux"`
	NonAzureComputerNames []string            `pulumi:"nonAzureComputerNames"`
	OperatingSystem       OperatingSystemType `pulumi:"operatingSystem"`
	Targets               *TargetProperties   `pulumi:"targets"`
	Windows               *WindowsProperties  `pulumi:"windows"`
}





type UpdateConfigurationInput interface {
	pulumi.Input

	ToUpdateConfigurationOutput() UpdateConfigurationOutput
	ToUpdateConfigurationOutputWithContext(context.Context) UpdateConfigurationOutput
}

type UpdateConfigurationArgs struct {
	AzureVirtualMachines  pulumi.StringArrayInput   `pulumi:"azureVirtualMachines"`
	Duration              pulumi.StringPtrInput     `pulumi:"duration"`
	Linux                 LinuxPropertiesPtrInput   `pulumi:"linux"`
	NonAzureComputerNames pulumi.StringArrayInput   `pulumi:"nonAzureComputerNames"`
	OperatingSystem       OperatingSystemTypeInput  `pulumi:"operatingSystem"`
	Targets               TargetPropertiesPtrInput  `pulumi:"targets"`
	Windows               WindowsPropertiesPtrInput `pulumi:"windows"`
}

func (UpdateConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateConfiguration)(nil)).Elem()
}

func (i UpdateConfigurationArgs) ToUpdateConfigurationOutput() UpdateConfigurationOutput {
	return i.ToUpdateConfigurationOutputWithContext(context.Background())
}

func (i UpdateConfigurationArgs) ToUpdateConfigurationOutputWithContext(ctx context.Context) UpdateConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpdateConfigurationOutput)
}

type UpdateConfigurationOutput struct{ *pulumi.OutputState }

func (UpdateConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateConfiguration)(nil)).Elem()
}

func (o UpdateConfigurationOutput) ToUpdateConfigurationOutput() UpdateConfigurationOutput {
	return o
}

func (o UpdateConfigurationOutput) ToUpdateConfigurationOutputWithContext(ctx context.Context) UpdateConfigurationOutput {
	return o
}

func (o UpdateConfigurationOutput) AzureVirtualMachines() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UpdateConfiguration) []string { return v.AzureVirtualMachines }).(pulumi.StringArrayOutput)
}

func (o UpdateConfigurationOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdateConfiguration) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

func (o UpdateConfigurationOutput) Linux() LinuxPropertiesPtrOutput {
	return o.ApplyT(func(v UpdateConfiguration) *LinuxProperties { return v.Linux }).(LinuxPropertiesPtrOutput)
}

func (o UpdateConfigurationOutput) NonAzureComputerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UpdateConfiguration) []string { return v.NonAzureComputerNames }).(pulumi.StringArrayOutput)
}

func (o UpdateConfigurationOutput) OperatingSystem() OperatingSystemTypeOutput {
	return o.ApplyT(func(v UpdateConfiguration) OperatingSystemType { return v.OperatingSystem }).(OperatingSystemTypeOutput)
}

func (o UpdateConfigurationOutput) Targets() TargetPropertiesPtrOutput {
	return o.ApplyT(func(v UpdateConfiguration) *TargetProperties { return v.Targets }).(TargetPropertiesPtrOutput)
}

func (o UpdateConfigurationOutput) Windows() WindowsPropertiesPtrOutput {
	return o.ApplyT(func(v UpdateConfiguration) *WindowsProperties { return v.Windows }).(WindowsPropertiesPtrOutput)
}

type UpdateConfigurationResponse struct {
	AzureVirtualMachines  []string                   `pulumi:"azureVirtualMachines"`
	Duration              *string                    `pulumi:"duration"`
	Linux                 *LinuxPropertiesResponse   `pulumi:"linux"`
	NonAzureComputerNames []string                   `pulumi:"nonAzureComputerNames"`
	OperatingSystem       string                     `pulumi:"operatingSystem"`
	Targets               *TargetPropertiesResponse  `pulumi:"targets"`
	Windows               *WindowsPropertiesResponse `pulumi:"windows"`
}

type UpdateConfigurationResponseOutput struct{ *pulumi.OutputState }

func (UpdateConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpdateConfigurationResponse)(nil)).Elem()
}

func (o UpdateConfigurationResponseOutput) ToUpdateConfigurationResponseOutput() UpdateConfigurationResponseOutput {
	return o
}

func (o UpdateConfigurationResponseOutput) ToUpdateConfigurationResponseOutputWithContext(ctx context.Context) UpdateConfigurationResponseOutput {
	return o
}

func (o UpdateConfigurationResponseOutput) AzureVirtualMachines() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) []string { return v.AzureVirtualMachines }).(pulumi.StringArrayOutput)
}

func (o UpdateConfigurationResponseOutput) Duration() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) *string { return v.Duration }).(pulumi.StringPtrOutput)
}

func (o UpdateConfigurationResponseOutput) Linux() LinuxPropertiesResponsePtrOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) *LinuxPropertiesResponse { return v.Linux }).(LinuxPropertiesResponsePtrOutput)
}

func (o UpdateConfigurationResponseOutput) NonAzureComputerNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) []string { return v.NonAzureComputerNames }).(pulumi.StringArrayOutput)
}

func (o UpdateConfigurationResponseOutput) OperatingSystem() pulumi.StringOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) string { return v.OperatingSystem }).(pulumi.StringOutput)
}

func (o UpdateConfigurationResponseOutput) Targets() TargetPropertiesResponsePtrOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) *TargetPropertiesResponse { return v.Targets }).(TargetPropertiesResponsePtrOutput)
}

func (o UpdateConfigurationResponseOutput) Windows() WindowsPropertiesResponsePtrOutput {
	return o.ApplyT(func(v UpdateConfigurationResponse) *WindowsPropertiesResponse { return v.Windows }).(WindowsPropertiesResponsePtrOutput)
}

type WindowsProperties struct {
	ExcludedKbNumbers             []string `pulumi:"excludedKbNumbers"`
	IncludedKbNumbers             []string `pulumi:"includedKbNumbers"`
	IncludedUpdateClassifications *string  `pulumi:"includedUpdateClassifications"`
	RebootSetting                 *string  `pulumi:"rebootSetting"`
}





type WindowsPropertiesInput interface {
	pulumi.Input

	ToWindowsPropertiesOutput() WindowsPropertiesOutput
	ToWindowsPropertiesOutputWithContext(context.Context) WindowsPropertiesOutput
}

type WindowsPropertiesArgs struct {
	ExcludedKbNumbers             pulumi.StringArrayInput `pulumi:"excludedKbNumbers"`
	IncludedKbNumbers             pulumi.StringArrayInput `pulumi:"includedKbNumbers"`
	IncludedUpdateClassifications pulumi.StringPtrInput   `pulumi:"includedUpdateClassifications"`
	RebootSetting                 pulumi.StringPtrInput   `pulumi:"rebootSetting"`
}

func (WindowsPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsProperties)(nil)).Elem()
}

func (i WindowsPropertiesArgs) ToWindowsPropertiesOutput() WindowsPropertiesOutput {
	return i.ToWindowsPropertiesOutputWithContext(context.Background())
}

func (i WindowsPropertiesArgs) ToWindowsPropertiesOutputWithContext(ctx context.Context) WindowsPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsPropertiesOutput)
}

func (i WindowsPropertiesArgs) ToWindowsPropertiesPtrOutput() WindowsPropertiesPtrOutput {
	return i.ToWindowsPropertiesPtrOutputWithContext(context.Background())
}

func (i WindowsPropertiesArgs) ToWindowsPropertiesPtrOutputWithContext(ctx context.Context) WindowsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsPropertiesOutput).ToWindowsPropertiesPtrOutputWithContext(ctx)
}









type WindowsPropertiesPtrInput interface {
	pulumi.Input

	ToWindowsPropertiesPtrOutput() WindowsPropertiesPtrOutput
	ToWindowsPropertiesPtrOutputWithContext(context.Context) WindowsPropertiesPtrOutput
}

type windowsPropertiesPtrType WindowsPropertiesArgs

func WindowsPropertiesPtr(v *WindowsPropertiesArgs) WindowsPropertiesPtrInput {
	return (*windowsPropertiesPtrType)(v)
}

func (*windowsPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsProperties)(nil)).Elem()
}

func (i *windowsPropertiesPtrType) ToWindowsPropertiesPtrOutput() WindowsPropertiesPtrOutput {
	return i.ToWindowsPropertiesPtrOutputWithContext(context.Background())
}

func (i *windowsPropertiesPtrType) ToWindowsPropertiesPtrOutputWithContext(ctx context.Context) WindowsPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WindowsPropertiesPtrOutput)
}

type WindowsPropertiesOutput struct{ *pulumi.OutputState }

func (WindowsPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsProperties)(nil)).Elem()
}

func (o WindowsPropertiesOutput) ToWindowsPropertiesOutput() WindowsPropertiesOutput {
	return o
}

func (o WindowsPropertiesOutput) ToWindowsPropertiesOutputWithContext(ctx context.Context) WindowsPropertiesOutput {
	return o
}

func (o WindowsPropertiesOutput) ToWindowsPropertiesPtrOutput() WindowsPropertiesPtrOutput {
	return o.ToWindowsPropertiesPtrOutputWithContext(context.Background())
}

func (o WindowsPropertiesOutput) ToWindowsPropertiesPtrOutputWithContext(ctx context.Context) WindowsPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WindowsProperties) *WindowsProperties {
		return &v
	}).(WindowsPropertiesPtrOutput)
}

func (o WindowsPropertiesOutput) ExcludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsProperties) []string { return v.ExcludedKbNumbers }).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesOutput) IncludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsProperties) []string { return v.IncludedKbNumbers }).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesOutput) IncludedUpdateClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProperties) *string { return v.IncludedUpdateClassifications }).(pulumi.StringPtrOutput)
}

func (o WindowsPropertiesOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsProperties) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type WindowsPropertiesPtrOutput struct{ *pulumi.OutputState }

func (WindowsPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsProperties)(nil)).Elem()
}

func (o WindowsPropertiesPtrOutput) ToWindowsPropertiesPtrOutput() WindowsPropertiesPtrOutput {
	return o
}

func (o WindowsPropertiesPtrOutput) ToWindowsPropertiesPtrOutputWithContext(ctx context.Context) WindowsPropertiesPtrOutput {
	return o
}

func (o WindowsPropertiesPtrOutput) Elem() WindowsPropertiesOutput {
	return o.ApplyT(func(v *WindowsProperties) WindowsProperties {
		if v != nil {
			return *v
		}
		var ret WindowsProperties
		return ret
	}).(WindowsPropertiesOutput)
}

func (o WindowsPropertiesPtrOutput) ExcludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WindowsProperties) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedKbNumbers
	}).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesPtrOutput) IncludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WindowsProperties) []string {
		if v == nil {
			return nil
		}
		return v.IncludedKbNumbers
	}).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesPtrOutput) IncludedUpdateClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProperties) *string {
		if v == nil {
			return nil
		}
		return v.IncludedUpdateClassifications
	}).(pulumi.StringPtrOutput)
}

func (o WindowsPropertiesPtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsProperties) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

type WindowsPropertiesResponse struct {
	ExcludedKbNumbers             []string `pulumi:"excludedKbNumbers"`
	IncludedKbNumbers             []string `pulumi:"includedKbNumbers"`
	IncludedUpdateClassifications *string  `pulumi:"includedUpdateClassifications"`
	RebootSetting                 *string  `pulumi:"rebootSetting"`
}

type WindowsPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WindowsPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WindowsPropertiesResponse)(nil)).Elem()
}

func (o WindowsPropertiesResponseOutput) ToWindowsPropertiesResponseOutput() WindowsPropertiesResponseOutput {
	return o
}

func (o WindowsPropertiesResponseOutput) ToWindowsPropertiesResponseOutputWithContext(ctx context.Context) WindowsPropertiesResponseOutput {
	return o
}

func (o WindowsPropertiesResponseOutput) ExcludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsPropertiesResponse) []string { return v.ExcludedKbNumbers }).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesResponseOutput) IncludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WindowsPropertiesResponse) []string { return v.IncludedKbNumbers }).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesResponseOutput) IncludedUpdateClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsPropertiesResponse) *string { return v.IncludedUpdateClassifications }).(pulumi.StringPtrOutput)
}

func (o WindowsPropertiesResponseOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WindowsPropertiesResponse) *string { return v.RebootSetting }).(pulumi.StringPtrOutput)
}

type WindowsPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WindowsPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WindowsPropertiesResponse)(nil)).Elem()
}

func (o WindowsPropertiesResponsePtrOutput) ToWindowsPropertiesResponsePtrOutput() WindowsPropertiesResponsePtrOutput {
	return o
}

func (o WindowsPropertiesResponsePtrOutput) ToWindowsPropertiesResponsePtrOutputWithContext(ctx context.Context) WindowsPropertiesResponsePtrOutput {
	return o
}

func (o WindowsPropertiesResponsePtrOutput) Elem() WindowsPropertiesResponseOutput {
	return o.ApplyT(func(v *WindowsPropertiesResponse) WindowsPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret WindowsPropertiesResponse
		return ret
	}).(WindowsPropertiesResponseOutput)
}

func (o WindowsPropertiesResponsePtrOutput) ExcludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WindowsPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExcludedKbNumbers
	}).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesResponsePtrOutput) IncludedKbNumbers() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WindowsPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.IncludedKbNumbers
	}).(pulumi.StringArrayOutput)
}

func (o WindowsPropertiesResponsePtrOutput) IncludedUpdateClassifications() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IncludedUpdateClassifications
	}).(pulumi.StringPtrOutput)
}

func (o WindowsPropertiesResponsePtrOutput) RebootSetting() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WindowsPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.RebootSetting
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AdvancedScheduleOutput{})
	pulumi.RegisterOutputType(AdvancedSchedulePtrOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceArrayOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceResponseOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleMonthlyOccurrenceResponseArrayOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleResponseOutput{})
	pulumi.RegisterOutputType(AdvancedScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureQueryPropertiesOutput{})
	pulumi.RegisterOutputType(AzureQueryPropertiesArrayOutput{})
	pulumi.RegisterOutputType(AzureQueryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AzureQueryPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnectionTypeAssociationPropertyOutput{})
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
	pulumi.RegisterOutputType(ContentSourceResponseOutput{})
	pulumi.RegisterOutputType(ContentSourceResponsePtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(DscConfigurationAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterMapOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterResponseOutput{})
	pulumi.RegisterOutputType(DscConfigurationParameterResponseMapOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesIdentityPtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseIdentityOutput{})
	pulumi.RegisterOutputType(EncryptionPropertiesResponseIdentityPtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(FieldDefinitionOutput{})
	pulumi.RegisterOutputType(FieldDefinitionMapOutput{})
	pulumi.RegisterOutputType(FieldDefinitionResponseOutput{})
	pulumi.RegisterOutputType(FieldDefinitionResponseMapOutput{})
	pulumi.RegisterOutputType(HybridRunbookWorkerLegacyResponseOutput{})
	pulumi.RegisterOutputType(HybridRunbookWorkerLegacyResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityOutput{})
	pulumi.RegisterOutputType(IdentityPtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseOutput{})
	pulumi.RegisterOutputType(IdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(IdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(ModuleErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesArrayOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyPtrOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(RunAsCredentialAssociationPropertyResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SUCSchedulePropertiesOutput{})
	pulumi.RegisterOutputType(SUCSchedulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyResponseOutput{})
	pulumi.RegisterOutputType(ScheduleAssociationPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksPtrOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksResponseOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceControlSecurityTokenPropertiesOutput{})
	pulumi.RegisterOutputType(SourceControlSecurityTokenPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TagSettingsPropertiesOutput{})
	pulumi.RegisterOutputType(TagSettingsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TagSettingsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TagSettingsPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TargetPropertiesOutput{})
	pulumi.RegisterOutputType(TargetPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TargetPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TargetPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TaskPropertiesOutput{})
	pulumi.RegisterOutputType(TaskPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TaskPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TaskPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(UpdateConfigurationOutput{})
	pulumi.RegisterOutputType(UpdateConfigurationResponseOutput{})
	pulumi.RegisterOutputType(WindowsPropertiesOutput{})
	pulumi.RegisterOutputType(WindowsPropertiesPtrOutput{})
	pulumi.RegisterOutputType(WindowsPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WindowsPropertiesResponsePtrOutput{})
}
