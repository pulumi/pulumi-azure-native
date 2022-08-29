


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FileDestination struct {
	FileFormats []string `pulumi:"fileFormats"`
}





type FileDestinationInput interface {
	pulumi.Input

	ToFileDestinationOutput() FileDestinationOutput
	ToFileDestinationOutputWithContext(context.Context) FileDestinationOutput
}

type FileDestinationArgs struct {
	FileFormats pulumi.StringArrayInput `pulumi:"fileFormats"`
}

func (FileDestinationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileDestination)(nil)).Elem()
}

func (i FileDestinationArgs) ToFileDestinationOutput() FileDestinationOutput {
	return i.ToFileDestinationOutputWithContext(context.Background())
}

func (i FileDestinationArgs) ToFileDestinationOutputWithContext(ctx context.Context) FileDestinationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileDestinationOutput)
}

func (i FileDestinationArgs) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return i.ToFileDestinationPtrOutputWithContext(context.Background())
}

func (i FileDestinationArgs) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileDestinationOutput).ToFileDestinationPtrOutputWithContext(ctx)
}









type FileDestinationPtrInput interface {
	pulumi.Input

	ToFileDestinationPtrOutput() FileDestinationPtrOutput
	ToFileDestinationPtrOutputWithContext(context.Context) FileDestinationPtrOutput
}

type fileDestinationPtrType FileDestinationArgs

func FileDestinationPtr(v *FileDestinationArgs) FileDestinationPtrInput {
	return (*fileDestinationPtrType)(v)
}

func (*fileDestinationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileDestination)(nil)).Elem()
}

func (i *fileDestinationPtrType) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return i.ToFileDestinationPtrOutputWithContext(context.Background())
}

func (i *fileDestinationPtrType) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileDestinationPtrOutput)
}

type FileDestinationOutput struct{ *pulumi.OutputState }

func (FileDestinationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileDestination)(nil)).Elem()
}

func (o FileDestinationOutput) ToFileDestinationOutput() FileDestinationOutput {
	return o
}

func (o FileDestinationOutput) ToFileDestinationOutputWithContext(ctx context.Context) FileDestinationOutput {
	return o
}

func (o FileDestinationOutput) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return o.ToFileDestinationPtrOutputWithContext(context.Background())
}

func (o FileDestinationOutput) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileDestination) *FileDestination {
		return &v
	}).(FileDestinationPtrOutput)
}

func (o FileDestinationOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FileDestination) []string { return v.FileFormats }).(pulumi.StringArrayOutput)
}

type FileDestinationPtrOutput struct{ *pulumi.OutputState }

func (FileDestinationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileDestination)(nil)).Elem()
}

func (o FileDestinationPtrOutput) ToFileDestinationPtrOutput() FileDestinationPtrOutput {
	return o
}

func (o FileDestinationPtrOutput) ToFileDestinationPtrOutputWithContext(ctx context.Context) FileDestinationPtrOutput {
	return o
}

func (o FileDestinationPtrOutput) Elem() FileDestinationOutput {
	return o.ApplyT(func(v *FileDestination) FileDestination {
		if v != nil {
			return *v
		}
		var ret FileDestination
		return ret
	}).(FileDestinationOutput)
}

func (o FileDestinationPtrOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FileDestination) []string {
		if v == nil {
			return nil
		}
		return v.FileFormats
	}).(pulumi.StringArrayOutput)
}

type FileDestinationResponse struct {
	FileFormats []string `pulumi:"fileFormats"`
}

type FileDestinationResponseOutput struct{ *pulumi.OutputState }

func (FileDestinationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileDestinationResponse)(nil)).Elem()
}

func (o FileDestinationResponseOutput) ToFileDestinationResponseOutput() FileDestinationResponseOutput {
	return o
}

func (o FileDestinationResponseOutput) ToFileDestinationResponseOutputWithContext(ctx context.Context) FileDestinationResponseOutput {
	return o
}

func (o FileDestinationResponseOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v FileDestinationResponse) []string { return v.FileFormats }).(pulumi.StringArrayOutput)
}

type FileDestinationResponsePtrOutput struct{ *pulumi.OutputState }

func (FileDestinationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileDestinationResponse)(nil)).Elem()
}

func (o FileDestinationResponsePtrOutput) ToFileDestinationResponsePtrOutput() FileDestinationResponsePtrOutput {
	return o
}

func (o FileDestinationResponsePtrOutput) ToFileDestinationResponsePtrOutputWithContext(ctx context.Context) FileDestinationResponsePtrOutput {
	return o
}

func (o FileDestinationResponsePtrOutput) Elem() FileDestinationResponseOutput {
	return o.ApplyT(func(v *FileDestinationResponse) FileDestinationResponse {
		if v != nil {
			return *v
		}
		var ret FileDestinationResponse
		return ret
	}).(FileDestinationResponseOutput)
}

func (o FileDestinationResponsePtrOutput) FileFormats() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FileDestinationResponse) []string {
		if v == nil {
			return nil
		}
		return v.FileFormats
	}).(pulumi.StringArrayOutput)
}

type NotificationProperties struct {
	Message *string  `pulumi:"message"`
	Subject string   `pulumi:"subject"`
	To      []string `pulumi:"to"`
}





type NotificationPropertiesInput interface {
	pulumi.Input

	ToNotificationPropertiesOutput() NotificationPropertiesOutput
	ToNotificationPropertiesOutputWithContext(context.Context) NotificationPropertiesOutput
}

type NotificationPropertiesArgs struct {
	Message pulumi.StringPtrInput   `pulumi:"message"`
	Subject pulumi.StringInput      `pulumi:"subject"`
	To      pulumi.StringArrayInput `pulumi:"to"`
}

func (NotificationPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationProperties)(nil)).Elem()
}

func (i NotificationPropertiesArgs) ToNotificationPropertiesOutput() NotificationPropertiesOutput {
	return i.ToNotificationPropertiesOutputWithContext(context.Background())
}

func (i NotificationPropertiesArgs) ToNotificationPropertiesOutputWithContext(ctx context.Context) NotificationPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NotificationPropertiesOutput)
}

type NotificationPropertiesOutput struct{ *pulumi.OutputState }

func (NotificationPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationProperties)(nil)).Elem()
}

func (o NotificationPropertiesOutput) ToNotificationPropertiesOutput() NotificationPropertiesOutput {
	return o
}

func (o NotificationPropertiesOutput) ToNotificationPropertiesOutputWithContext(ctx context.Context) NotificationPropertiesOutput {
	return o
}

func (o NotificationPropertiesOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationProperties) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o NotificationPropertiesOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationProperties) string { return v.Subject }).(pulumi.StringOutput)
}

func (o NotificationPropertiesOutput) To() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationProperties) []string { return v.To }).(pulumi.StringArrayOutput)
}

type NotificationPropertiesResponse struct {
	Message *string  `pulumi:"message"`
	Subject string   `pulumi:"subject"`
	To      []string `pulumi:"to"`
}

type NotificationPropertiesResponseOutput struct{ *pulumi.OutputState }

func (NotificationPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NotificationPropertiesResponse)(nil)).Elem()
}

func (o NotificationPropertiesResponseOutput) ToNotificationPropertiesResponseOutput() NotificationPropertiesResponseOutput {
	return o
}

func (o NotificationPropertiesResponseOutput) ToNotificationPropertiesResponseOutputWithContext(ctx context.Context) NotificationPropertiesResponseOutput {
	return o
}

func (o NotificationPropertiesResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NotificationPropertiesResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o NotificationPropertiesResponseOutput) Subject() pulumi.StringOutput {
	return o.ApplyT(func(v NotificationPropertiesResponse) string { return v.Subject }).(pulumi.StringOutput)
}

func (o NotificationPropertiesResponseOutput) To() pulumi.StringArrayOutput {
	return o.ApplyT(func(v NotificationPropertiesResponse) []string { return v.To }).(pulumi.StringArrayOutput)
}

type ScheduleProperties struct {
	DayOfMonth   *int     `pulumi:"dayOfMonth"`
	DaysOfWeek   []string `pulumi:"daysOfWeek"`
	EndDate      string   `pulumi:"endDate"`
	Frequency    string   `pulumi:"frequency"`
	HourOfDay    *int     `pulumi:"hourOfDay"`
	StartDate    string   `pulumi:"startDate"`
	WeeksOfMonth []string `pulumi:"weeksOfMonth"`
}





type SchedulePropertiesInput interface {
	pulumi.Input

	ToSchedulePropertiesOutput() SchedulePropertiesOutput
	ToSchedulePropertiesOutputWithContext(context.Context) SchedulePropertiesOutput
}

type SchedulePropertiesArgs struct {
	DayOfMonth   pulumi.IntPtrInput      `pulumi:"dayOfMonth"`
	DaysOfWeek   pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	EndDate      pulumi.StringInput      `pulumi:"endDate"`
	Frequency    pulumi.StringInput      `pulumi:"frequency"`
	HourOfDay    pulumi.IntPtrInput      `pulumi:"hourOfDay"`
	StartDate    pulumi.StringInput      `pulumi:"startDate"`
	WeeksOfMonth pulumi.StringArrayInput `pulumi:"weeksOfMonth"`
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

func (o SchedulePropertiesOutput) DayOfMonth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *int { return v.DayOfMonth }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleProperties) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o SchedulePropertiesOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleProperties) string { return v.EndDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleProperties) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o SchedulePropertiesOutput) HourOfDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleProperties) *int { return v.HourOfDay }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v ScheduleProperties) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesOutput) WeeksOfMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ScheduleProperties) []string { return v.WeeksOfMonth }).(pulumi.StringArrayOutput)
}

type SchedulePropertiesResponse struct {
	DayOfMonth   *int     `pulumi:"dayOfMonth"`
	DaysOfWeek   []string `pulumi:"daysOfWeek"`
	EndDate      string   `pulumi:"endDate"`
	Frequency    string   `pulumi:"frequency"`
	HourOfDay    *int     `pulumi:"hourOfDay"`
	StartDate    string   `pulumi:"startDate"`
	WeeksOfMonth []string `pulumi:"weeksOfMonth"`
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

func (o SchedulePropertiesResponseOutput) DayOfMonth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *int { return v.DayOfMonth }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesResponseOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o SchedulePropertiesResponseOutput) EndDate() pulumi.StringOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) string { return v.EndDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesResponseOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) string { return v.Frequency }).(pulumi.StringOutput)
}

func (o SchedulePropertiesResponseOutput) HourOfDay() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) *int { return v.HourOfDay }).(pulumi.IntPtrOutput)
}

func (o SchedulePropertiesResponseOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) string { return v.StartDate }).(pulumi.StringOutput)
}

func (o SchedulePropertiesResponseOutput) WeeksOfMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SchedulePropertiesResponse) []string { return v.WeeksOfMonth }).(pulumi.StringArrayOutput)
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
	pulumi.RegisterOutputType(FileDestinationOutput{})
	pulumi.RegisterOutputType(FileDestinationPtrOutput{})
	pulumi.RegisterOutputType(FileDestinationResponseOutput{})
	pulumi.RegisterOutputType(FileDestinationResponsePtrOutput{})
	pulumi.RegisterOutputType(NotificationPropertiesOutput{})
	pulumi.RegisterOutputType(NotificationPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SchedulePropertiesOutput{})
	pulumi.RegisterOutputType(SchedulePropertiesResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
}
