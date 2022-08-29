


package v20170515preview

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

type ScheduleProperties struct {
	AdvancedSchedule        *AdvancedSchedule `pulumi:"advancedSchedule"`
	CreationTime            *string           `pulumi:"creationTime"`
	Description             *string           `pulumi:"description"`
	ExpiryTime              *string           `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes *float64          `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               *string           `pulumi:"frequency"`
	Interval                *int              `pulumi:"interval"`
	IsEnabled               *bool             `pulumi:"isEnabled"`
	LastModifiedTime        *string           `pulumi:"lastModifiedTime"`
	NextRun                 *string           `pulumi:"nextRun"`
	NextRunOffsetMinutes    *float64          `pulumi:"nextRunOffsetMinutes"`
	StartTime               *string           `pulumi:"startTime"`
	TimeZone                *string           `pulumi:"timeZone"`
}


func (val *ScheduleProperties) Defaults() *ScheduleProperties {
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





type SchedulePropertiesInput interface {
	pulumi.Input

	ToSchedulePropertiesOutput() SchedulePropertiesOutput
	ToSchedulePropertiesOutputWithContext(context.Context) SchedulePropertiesOutput
}

type SchedulePropertiesArgs struct {
	AdvancedSchedule        AdvancedSchedulePtrInput `pulumi:"advancedSchedule"`
	CreationTime            pulumi.StringPtrInput    `pulumi:"creationTime"`
	Description             pulumi.StringPtrInput    `pulumi:"description"`
	ExpiryTime              pulumi.StringPtrInput    `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes pulumi.Float64PtrInput   `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               pulumi.StringPtrInput    `pulumi:"frequency"`
	Interval                pulumi.IntPtrInput       `pulumi:"interval"`
	IsEnabled               pulumi.BoolPtrInput      `pulumi:"isEnabled"`
	LastModifiedTime        pulumi.StringPtrInput    `pulumi:"lastModifiedTime"`
	NextRun                 pulumi.StringPtrInput    `pulumi:"nextRun"`
	NextRunOffsetMinutes    pulumi.Float64PtrInput   `pulumi:"nextRunOffsetMinutes"`
	StartTime               pulumi.StringPtrInput    `pulumi:"startTime"`
	TimeZone                pulumi.StringPtrInput    `pulumi:"timeZone"`
}


func (val *SchedulePropertiesArgs) Defaults() *SchedulePropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsEnabled) {
		tmp.IsEnabled = pulumi.BoolPtr(false)
	}
	return &tmp
}
func (SchedulePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleProperties)(nil)).Elem()
}

func (i SchedulePropertiesArgs) ToSchedulePropertiesOutput() SchedulePropertiesOutput {
	return i.ToSchedulePropertiesOutputWithContext(context.Background())
}

func (i SchedulePropertiesArgs) ToSchedulePropertiesOutputWithContext(ctx context.Context) SchedulePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePropertiesOutput)
}

type SchedulePropertiesOutput struct{ *pulumi.OutputState }

func (SchedulePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleProperties)(nil)).Elem()
}

func (o SchedulePropertiesOutput) ToSchedulePropertiesOutput() SchedulePropertiesOutput {
	return o
}

func (o SchedulePropertiesOutput) ToSchedulePropertiesOutputWithContext(ctx context.Context) SchedulePropertiesOutput {
	return o
}

func (o SchedulePropertiesOutput) AdvancedSchedule() AdvancedSchedulePtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *AdvancedSchedule { return v.AdvancedSchedule }).(AdvancedSchedulePtrOutput)
}

func (o SchedulePropertiesOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) ExpiryTimeOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *float64 { return v.ExpiryTimeOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SchedulePropertiesOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SchedulePropertiesOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) NextRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.NextRun }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) NextRunOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *float64 { return v.NextRunOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SchedulePropertiesOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
}

type SchedulePropertiesResponse struct {
	AdvancedSchedule        *AdvancedScheduleResponse `pulumi:"advancedSchedule"`
	CreationTime            *string                   `pulumi:"creationTime"`
	Description             *string                   `pulumi:"description"`
	ExpiryTime              *string                   `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes *float64                  `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               *string                   `pulumi:"frequency"`
	Interval                *int                      `pulumi:"interval"`
	IsEnabled               *bool                     `pulumi:"isEnabled"`
	LastModifiedTime        *string                   `pulumi:"lastModifiedTime"`
	NextRun                 *string                   `pulumi:"nextRun"`
	NextRunOffsetMinutes    *float64                  `pulumi:"nextRunOffsetMinutes"`
	StartTime               *string                   `pulumi:"startTime"`
	StartTimeOffsetMinutes  float64                   `pulumi:"startTimeOffsetMinutes"`
	TimeZone                *string                   `pulumi:"timeZone"`
}


func (val *SchedulePropertiesResponse) Defaults() *SchedulePropertiesResponse {
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

type SchedulePropertiesResponseOutput struct{ *pulumi.OutputState }

func (SchedulePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SchedulePropertiesResponse)(nil)).Elem()
}

func (o SchedulePropertiesResponseOutput) ToSchedulePropertiesResponseOutput() SchedulePropertiesResponseOutput {
	return o
}

func (o SchedulePropertiesResponseOutput) ToSchedulePropertiesResponseOutputWithContext(ctx context.Context) SchedulePropertiesResponseOutput {
	return o
}

func (o SchedulePropertiesResponseOutput) AdvancedSchedule() AdvancedScheduleResponsePtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *AdvancedScheduleResponse { return v.AdvancedSchedule }).(AdvancedScheduleResponsePtrOutput)
}

func (o SchedulePropertiesResponseOutput) CreationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.CreationTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) ExpiryTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.ExpiryTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) ExpiryTimeOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *float64 { return v.ExpiryTimeOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SchedulePropertiesResponseOutput) Frequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.Frequency }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SchedulePropertiesResponseOutput) LastModifiedTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.LastModifiedTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) NextRun() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.NextRun }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) NextRunOffsetMinutes() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *float64 { return v.NextRunOffsetMinutes }).(pulumi.Float64PtrOutput)
}

func (o SchedulePropertiesResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o SchedulePropertiesResponseOutput) StartTimeOffsetMinutes() pulumi.Float64Output {
	return o.ApplyT(func(v SchedulePropertiesResponse) float64 { return v.StartTimeOffsetMinutes }).(pulumi.Float64Output)
}

func (o SchedulePropertiesResponseOutput) TimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *string { return v.TimeZone }).(pulumi.StringPtrOutput)
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
	pulumi.RegisterOutputType(ErrorResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponseOutput{})
	pulumi.RegisterOutputType(ErrorResponseResponsePtrOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LinuxPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesArrayOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesResponseOutput{})
	pulumi.RegisterOutputType(NonAzureQueryPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(SchedulePropertiesOutput{})
	pulumi.RegisterOutputType(SchedulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksPtrOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksResponseOutput{})
	pulumi.RegisterOutputType(SoftwareUpdateConfigurationTasksResponsePtrOutput{})
	pulumi.RegisterOutputType(SourceControlSecurityTokenPropertiesOutput{})
	pulumi.RegisterOutputType(SourceControlSecurityTokenPropertiesPtrOutput{})
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
