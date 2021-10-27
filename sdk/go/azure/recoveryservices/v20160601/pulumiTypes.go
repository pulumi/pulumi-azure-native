


package v20160601

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureIaaSVMProtectionPolicy struct {
	BackupManagementType *string     `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
}





type AzureIaaSVMProtectionPolicyInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectionPolicyOutput() AzureIaaSVMProtectionPolicyOutput
	ToAzureIaaSVMProtectionPolicyOutputWithContext(context.Context) AzureIaaSVMProtectionPolicyOutput
}

type AzureIaaSVMProtectionPolicyArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ProtectedItemsCount  pulumi.IntPtrInput    `pulumi:"protectedItemsCount"`
	RetentionPolicy      pulumi.Input          `pulumi:"retentionPolicy"`
	SchedulePolicy       pulumi.Input          `pulumi:"schedulePolicy"`
}

func (AzureIaaSVMProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicy)(nil)).Elem()
}

func (i AzureIaaSVMProtectionPolicyArgs) ToAzureIaaSVMProtectionPolicyOutput() AzureIaaSVMProtectionPolicyOutput {
	return i.ToAzureIaaSVMProtectionPolicyOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectionPolicyArgs) ToAzureIaaSVMProtectionPolicyOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectionPolicyOutput)
}

type AzureIaaSVMProtectionPolicyOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicy)(nil)).Elem()
}

func (o AzureIaaSVMProtectionPolicyOutput) ToAzureIaaSVMProtectionPolicyOutput() AzureIaaSVMProtectionPolicyOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyOutput) ToAzureIaaSVMProtectionPolicyOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o AzureIaaSVMProtectionPolicyOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicy) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type AzureIaaSVMProtectionPolicyResponse struct {
	BackupManagementType *string     `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
}





type AzureIaaSVMProtectionPolicyResponseInput interface {
	pulumi.Input

	ToAzureIaaSVMProtectionPolicyResponseOutput() AzureIaaSVMProtectionPolicyResponseOutput
	ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(context.Context) AzureIaaSVMProtectionPolicyResponseOutput
}

type AzureIaaSVMProtectionPolicyResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ProtectedItemsCount  pulumi.IntPtrInput    `pulumi:"protectedItemsCount"`
	RetentionPolicy      pulumi.Input          `pulumi:"retentionPolicy"`
	SchedulePolicy       pulumi.Input          `pulumi:"schedulePolicy"`
}

func (AzureIaaSVMProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicyResponse)(nil)).Elem()
}

func (i AzureIaaSVMProtectionPolicyResponseArgs) ToAzureIaaSVMProtectionPolicyResponseOutput() AzureIaaSVMProtectionPolicyResponseOutput {
	return i.ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i AzureIaaSVMProtectionPolicyResponseArgs) ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureIaaSVMProtectionPolicyResponseOutput)
}

type AzureIaaSVMProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (AzureIaaSVMProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureIaaSVMProtectionPolicyResponse)(nil)).Elem()
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ToAzureIaaSVMProtectionPolicyResponseOutput() AzureIaaSVMProtectionPolicyResponseOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ToAzureIaaSVMProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureIaaSVMProtectionPolicyResponseOutput {
	return o
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o AzureIaaSVMProtectionPolicyResponseOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureIaaSVMProtectionPolicyResponse) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type AzureSqlProtectionPolicy struct {
	BackupManagementType *string     `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
}





type AzureSqlProtectionPolicyInput interface {
	pulumi.Input

	ToAzureSqlProtectionPolicyOutput() AzureSqlProtectionPolicyOutput
	ToAzureSqlProtectionPolicyOutputWithContext(context.Context) AzureSqlProtectionPolicyOutput
}

type AzureSqlProtectionPolicyArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ProtectedItemsCount  pulumi.IntPtrInput    `pulumi:"protectedItemsCount"`
	RetentionPolicy      pulumi.Input          `pulumi:"retentionPolicy"`
}

func (AzureSqlProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicy)(nil)).Elem()
}

func (i AzureSqlProtectionPolicyArgs) ToAzureSqlProtectionPolicyOutput() AzureSqlProtectionPolicyOutput {
	return i.ToAzureSqlProtectionPolicyOutputWithContext(context.Background())
}

func (i AzureSqlProtectionPolicyArgs) ToAzureSqlProtectionPolicyOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectionPolicyOutput)
}

type AzureSqlProtectionPolicyOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicy)(nil)).Elem()
}

func (o AzureSqlProtectionPolicyOutput) ToAzureSqlProtectionPolicyOutput() AzureSqlProtectionPolicyOutput {
	return o
}

func (o AzureSqlProtectionPolicyOutput) ToAzureSqlProtectionPolicyOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyOutput {
	return o
}

func (o AzureSqlProtectionPolicyOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureSqlProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

type AzureSqlProtectionPolicyResponse struct {
	BackupManagementType *string     `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
}





type AzureSqlProtectionPolicyResponseInput interface {
	pulumi.Input

	ToAzureSqlProtectionPolicyResponseOutput() AzureSqlProtectionPolicyResponseOutput
	ToAzureSqlProtectionPolicyResponseOutputWithContext(context.Context) AzureSqlProtectionPolicyResponseOutput
}

type AzureSqlProtectionPolicyResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ProtectedItemsCount  pulumi.IntPtrInput    `pulumi:"protectedItemsCount"`
	RetentionPolicy      pulumi.Input          `pulumi:"retentionPolicy"`
}

func (AzureSqlProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicyResponse)(nil)).Elem()
}

func (i AzureSqlProtectionPolicyResponseArgs) ToAzureSqlProtectionPolicyResponseOutput() AzureSqlProtectionPolicyResponseOutput {
	return i.ToAzureSqlProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i AzureSqlProtectionPolicyResponseArgs) ToAzureSqlProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureSqlProtectionPolicyResponseOutput)
}

type AzureSqlProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (AzureSqlProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureSqlProtectionPolicyResponse)(nil)).Elem()
}

func (o AzureSqlProtectionPolicyResponseOutput) ToAzureSqlProtectionPolicyResponseOutput() AzureSqlProtectionPolicyResponseOutput {
	return o
}

func (o AzureSqlProtectionPolicyResponseOutput) ToAzureSqlProtectionPolicyResponseOutputWithContext(ctx context.Context) AzureSqlProtectionPolicyResponseOutput {
	return o
}

func (o AzureSqlProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o AzureSqlProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o AzureSqlProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureSqlProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

type DailyRetentionFormat struct {
	DaysOfTheMonth []Day `pulumi:"daysOfTheMonth"`
}





type DailyRetentionFormatInput interface {
	pulumi.Input

	ToDailyRetentionFormatOutput() DailyRetentionFormatOutput
	ToDailyRetentionFormatOutputWithContext(context.Context) DailyRetentionFormatOutput
}

type DailyRetentionFormatArgs struct {
	DaysOfTheMonth DayArrayInput `pulumi:"daysOfTheMonth"`
}

func (DailyRetentionFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormat)(nil)).Elem()
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatOutput() DailyRetentionFormatOutput {
	return i.ToDailyRetentionFormatOutputWithContext(context.Background())
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatOutputWithContext(ctx context.Context) DailyRetentionFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatOutput)
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return i.ToDailyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i DailyRetentionFormatArgs) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatOutput).ToDailyRetentionFormatPtrOutputWithContext(ctx)
}









type DailyRetentionFormatPtrInput interface {
	pulumi.Input

	ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput
	ToDailyRetentionFormatPtrOutputWithContext(context.Context) DailyRetentionFormatPtrOutput
}

type dailyRetentionFormatPtrType DailyRetentionFormatArgs

func DailyRetentionFormatPtr(v *DailyRetentionFormatArgs) DailyRetentionFormatPtrInput {
	return (*dailyRetentionFormatPtrType)(v)
}

func (*dailyRetentionFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormat)(nil)).Elem()
}

func (i *dailyRetentionFormatPtrType) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return i.ToDailyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i *dailyRetentionFormatPtrType) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatPtrOutput)
}

type DailyRetentionFormatOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormat)(nil)).Elem()
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatOutput() DailyRetentionFormatOutput {
	return o
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatOutputWithContext(ctx context.Context) DailyRetentionFormatOutput {
	return o
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return o.ToDailyRetentionFormatPtrOutputWithContext(context.Background())
}

func (o DailyRetentionFormatOutput) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionFormat) *DailyRetentionFormat {
		return &v
	}).(DailyRetentionFormatPtrOutput)
}

func (o DailyRetentionFormatOutput) DaysOfTheMonth() DayArrayOutput {
	return o.ApplyT(func(v DailyRetentionFormat) []Day { return v.DaysOfTheMonth }).(DayArrayOutput)
}

type DailyRetentionFormatPtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormat)(nil)).Elem()
}

func (o DailyRetentionFormatPtrOutput) ToDailyRetentionFormatPtrOutput() DailyRetentionFormatPtrOutput {
	return o
}

func (o DailyRetentionFormatPtrOutput) ToDailyRetentionFormatPtrOutputWithContext(ctx context.Context) DailyRetentionFormatPtrOutput {
	return o
}

func (o DailyRetentionFormatPtrOutput) Elem() DailyRetentionFormatOutput {
	return o.ApplyT(func(v *DailyRetentionFormat) DailyRetentionFormat {
		if v != nil {
			return *v
		}
		var ret DailyRetentionFormat
		return ret
	}).(DailyRetentionFormatOutput)
}

func (o DailyRetentionFormatPtrOutput) DaysOfTheMonth() DayArrayOutput {
	return o.ApplyT(func(v *DailyRetentionFormat) []Day {
		if v == nil {
			return nil
		}
		return v.DaysOfTheMonth
	}).(DayArrayOutput)
}

type DailyRetentionFormatResponse struct {
	DaysOfTheMonth []DayResponse `pulumi:"daysOfTheMonth"`
}





type DailyRetentionFormatResponseInput interface {
	pulumi.Input

	ToDailyRetentionFormatResponseOutput() DailyRetentionFormatResponseOutput
	ToDailyRetentionFormatResponseOutputWithContext(context.Context) DailyRetentionFormatResponseOutput
}

type DailyRetentionFormatResponseArgs struct {
	DaysOfTheMonth DayResponseArrayInput `pulumi:"daysOfTheMonth"`
}

func (DailyRetentionFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormatResponse)(nil)).Elem()
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponseOutput() DailyRetentionFormatResponseOutput {
	return i.ToDailyRetentionFormatResponseOutputWithContext(context.Background())
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponseOutputWithContext(ctx context.Context) DailyRetentionFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatResponseOutput)
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return i.ToDailyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i DailyRetentionFormatResponseArgs) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatResponseOutput).ToDailyRetentionFormatResponsePtrOutputWithContext(ctx)
}









type DailyRetentionFormatResponsePtrInput interface {
	pulumi.Input

	ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput
	ToDailyRetentionFormatResponsePtrOutputWithContext(context.Context) DailyRetentionFormatResponsePtrOutput
}

type dailyRetentionFormatResponsePtrType DailyRetentionFormatResponseArgs

func DailyRetentionFormatResponsePtr(v *DailyRetentionFormatResponseArgs) DailyRetentionFormatResponsePtrInput {
	return (*dailyRetentionFormatResponsePtrType)(v)
}

func (*dailyRetentionFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormatResponse)(nil)).Elem()
}

func (i *dailyRetentionFormatResponsePtrType) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return i.ToDailyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i *dailyRetentionFormatResponsePtrType) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionFormatResponsePtrOutput)
}

type DailyRetentionFormatResponseOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionFormatResponse)(nil)).Elem()
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponseOutput() DailyRetentionFormatResponseOutput {
	return o
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponseOutputWithContext(ctx context.Context) DailyRetentionFormatResponseOutput {
	return o
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return o.ToDailyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (o DailyRetentionFormatResponseOutput) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionFormatResponse) *DailyRetentionFormatResponse {
		return &v
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o DailyRetentionFormatResponseOutput) DaysOfTheMonth() DayResponseArrayOutput {
	return o.ApplyT(func(v DailyRetentionFormatResponse) []DayResponse { return v.DaysOfTheMonth }).(DayResponseArrayOutput)
}

type DailyRetentionFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionFormatResponse)(nil)).Elem()
}

func (o DailyRetentionFormatResponsePtrOutput) ToDailyRetentionFormatResponsePtrOutput() DailyRetentionFormatResponsePtrOutput {
	return o
}

func (o DailyRetentionFormatResponsePtrOutput) ToDailyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) DailyRetentionFormatResponsePtrOutput {
	return o
}

func (o DailyRetentionFormatResponsePtrOutput) Elem() DailyRetentionFormatResponseOutput {
	return o.ApplyT(func(v *DailyRetentionFormatResponse) DailyRetentionFormatResponse {
		if v != nil {
			return *v
		}
		var ret DailyRetentionFormatResponse
		return ret
	}).(DailyRetentionFormatResponseOutput)
}

func (o DailyRetentionFormatResponsePtrOutput) DaysOfTheMonth() DayResponseArrayOutput {
	return o.ApplyT(func(v *DailyRetentionFormatResponse) []DayResponse {
		if v == nil {
			return nil
		}
		return v.DaysOfTheMonth
	}).(DayResponseArrayOutput)
}

type DailyRetentionSchedule struct {
	RetentionDuration *RetentionDuration `pulumi:"retentionDuration"`
	RetentionTimes    []string           `pulumi:"retentionTimes"`
}





type DailyRetentionScheduleInput interface {
	pulumi.Input

	ToDailyRetentionScheduleOutput() DailyRetentionScheduleOutput
	ToDailyRetentionScheduleOutputWithContext(context.Context) DailyRetentionScheduleOutput
}

type DailyRetentionScheduleArgs struct {
	RetentionDuration RetentionDurationPtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput   `pulumi:"retentionTimes"`
}

func (DailyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionSchedule)(nil)).Elem()
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionScheduleOutput() DailyRetentionScheduleOutput {
	return i.ToDailyRetentionScheduleOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionScheduleOutputWithContext(ctx context.Context) DailyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleOutput)
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return i.ToDailyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleArgs) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleOutput).ToDailyRetentionSchedulePtrOutputWithContext(ctx)
}









type DailyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput
	ToDailyRetentionSchedulePtrOutputWithContext(context.Context) DailyRetentionSchedulePtrOutput
}

type dailyRetentionSchedulePtrType DailyRetentionScheduleArgs

func DailyRetentionSchedulePtr(v *DailyRetentionScheduleArgs) DailyRetentionSchedulePtrInput {
	return (*dailyRetentionSchedulePtrType)(v)
}

func (*dailyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionSchedule)(nil)).Elem()
}

func (i *dailyRetentionSchedulePtrType) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return i.ToDailyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *dailyRetentionSchedulePtrType) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionSchedulePtrOutput)
}

type DailyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (DailyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionSchedule)(nil)).Elem()
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionScheduleOutput() DailyRetentionScheduleOutput {
	return o
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionScheduleOutputWithContext(ctx context.Context) DailyRetentionScheduleOutput {
	return o
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return o.ToDailyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o DailyRetentionScheduleOutput) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionSchedule) *DailyRetentionSchedule {
		return &v
	}).(DailyRetentionSchedulePtrOutput)
}

func (o DailyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v DailyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o DailyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DailyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type DailyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionSchedule)(nil)).Elem()
}

func (o DailyRetentionSchedulePtrOutput) ToDailyRetentionSchedulePtrOutput() DailyRetentionSchedulePtrOutput {
	return o
}

func (o DailyRetentionSchedulePtrOutput) ToDailyRetentionSchedulePtrOutputWithContext(ctx context.Context) DailyRetentionSchedulePtrOutput {
	return o
}

func (o DailyRetentionSchedulePtrOutput) Elem() DailyRetentionScheduleOutput {
	return o.ApplyT(func(v *DailyRetentionSchedule) DailyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret DailyRetentionSchedule
		return ret
	}).(DailyRetentionScheduleOutput)
}

func (o DailyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *DailyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o DailyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DailyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type DailyRetentionScheduleResponse struct {
	RetentionDuration *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionTimes    []string                   `pulumi:"retentionTimes"`
}





type DailyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToDailyRetentionScheduleResponseOutput() DailyRetentionScheduleResponseOutput
	ToDailyRetentionScheduleResponseOutputWithContext(context.Context) DailyRetentionScheduleResponseOutput
}

type DailyRetentionScheduleResponseArgs struct {
	RetentionDuration RetentionDurationResponsePtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput           `pulumi:"retentionTimes"`
}

func (DailyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionScheduleResponse)(nil)).Elem()
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponseOutput() DailyRetentionScheduleResponseOutput {
	return i.ToDailyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponseOutputWithContext(ctx context.Context) DailyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleResponseOutput)
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return i.ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i DailyRetentionScheduleResponseArgs) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleResponseOutput).ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type DailyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput
	ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Context) DailyRetentionScheduleResponsePtrOutput
}

type dailyRetentionScheduleResponsePtrType DailyRetentionScheduleResponseArgs

func DailyRetentionScheduleResponsePtr(v *DailyRetentionScheduleResponseArgs) DailyRetentionScheduleResponsePtrInput {
	return (*dailyRetentionScheduleResponsePtrType)(v)
}

func (*dailyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionScheduleResponse)(nil)).Elem()
}

func (i *dailyRetentionScheduleResponsePtrType) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return i.ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *dailyRetentionScheduleResponsePtrType) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DailyRetentionScheduleResponsePtrOutput)
}

type DailyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (DailyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DailyRetentionScheduleResponse)(nil)).Elem()
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponseOutput() DailyRetentionScheduleResponseOutput {
	return o
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponseOutputWithContext(ctx context.Context) DailyRetentionScheduleResponseOutput {
	return o
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return o.ToDailyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o DailyRetentionScheduleResponseOutput) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DailyRetentionScheduleResponse) *DailyRetentionScheduleResponse {
		return &v
	}).(DailyRetentionScheduleResponsePtrOutput)
}

func (o DailyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v DailyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o DailyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DailyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type DailyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (DailyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DailyRetentionScheduleResponse)(nil)).Elem()
}

func (o DailyRetentionScheduleResponsePtrOutput) ToDailyRetentionScheduleResponsePtrOutput() DailyRetentionScheduleResponsePtrOutput {
	return o
}

func (o DailyRetentionScheduleResponsePtrOutput) ToDailyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) DailyRetentionScheduleResponsePtrOutput {
	return o
}

func (o DailyRetentionScheduleResponsePtrOutput) Elem() DailyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *DailyRetentionScheduleResponse) DailyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret DailyRetentionScheduleResponse
		return ret
	}).(DailyRetentionScheduleResponseOutput)
}

func (o DailyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *DailyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o DailyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DailyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type Day struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}





type DayInput interface {
	pulumi.Input

	ToDayOutput() DayOutput
	ToDayOutputWithContext(context.Context) DayOutput
}

type DayArgs struct {
	Date   pulumi.IntPtrInput  `pulumi:"date"`
	IsLast pulumi.BoolPtrInput `pulumi:"isLast"`
}

func (DayArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Day)(nil)).Elem()
}

func (i DayArgs) ToDayOutput() DayOutput {
	return i.ToDayOutputWithContext(context.Background())
}

func (i DayArgs) ToDayOutputWithContext(ctx context.Context) DayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayOutput)
}





type DayArrayInput interface {
	pulumi.Input

	ToDayArrayOutput() DayArrayOutput
	ToDayArrayOutputWithContext(context.Context) DayArrayOutput
}

type DayArray []DayInput

func (DayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Day)(nil)).Elem()
}

func (i DayArray) ToDayArrayOutput() DayArrayOutput {
	return i.ToDayArrayOutputWithContext(context.Background())
}

func (i DayArray) ToDayArrayOutputWithContext(ctx context.Context) DayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayArrayOutput)
}

type DayOutput struct{ *pulumi.OutputState }

func (DayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Day)(nil)).Elem()
}

func (o DayOutput) ToDayOutput() DayOutput {
	return o
}

func (o DayOutput) ToDayOutputWithContext(ctx context.Context) DayOutput {
	return o
}

func (o DayOutput) Date() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Day) *int { return v.Date }).(pulumi.IntPtrOutput)
}

func (o DayOutput) IsLast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Day) *bool { return v.IsLast }).(pulumi.BoolPtrOutput)
}

type DayArrayOutput struct{ *pulumi.OutputState }

func (DayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Day)(nil)).Elem()
}

func (o DayArrayOutput) ToDayArrayOutput() DayArrayOutput {
	return o
}

func (o DayArrayOutput) ToDayArrayOutputWithContext(ctx context.Context) DayArrayOutput {
	return o
}

func (o DayArrayOutput) Index(i pulumi.IntInput) DayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Day {
		return vs[0].([]Day)[vs[1].(int)]
	}).(DayOutput)
}

type DayResponse struct {
	Date   *int  `pulumi:"date"`
	IsLast *bool `pulumi:"isLast"`
}





type DayResponseInput interface {
	pulumi.Input

	ToDayResponseOutput() DayResponseOutput
	ToDayResponseOutputWithContext(context.Context) DayResponseOutput
}

type DayResponseArgs struct {
	Date   pulumi.IntPtrInput  `pulumi:"date"`
	IsLast pulumi.BoolPtrInput `pulumi:"isLast"`
}

func (DayResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DayResponse)(nil)).Elem()
}

func (i DayResponseArgs) ToDayResponseOutput() DayResponseOutput {
	return i.ToDayResponseOutputWithContext(context.Background())
}

func (i DayResponseArgs) ToDayResponseOutputWithContext(ctx context.Context) DayResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayResponseOutput)
}





type DayResponseArrayInput interface {
	pulumi.Input

	ToDayResponseArrayOutput() DayResponseArrayOutput
	ToDayResponseArrayOutputWithContext(context.Context) DayResponseArrayOutput
}

type DayResponseArray []DayResponseInput

func (DayResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayResponse)(nil)).Elem()
}

func (i DayResponseArray) ToDayResponseArrayOutput() DayResponseArrayOutput {
	return i.ToDayResponseArrayOutputWithContext(context.Background())
}

func (i DayResponseArray) ToDayResponseArrayOutputWithContext(ctx context.Context) DayResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DayResponseArrayOutput)
}

type DayResponseOutput struct{ *pulumi.OutputState }

func (DayResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DayResponse)(nil)).Elem()
}

func (o DayResponseOutput) ToDayResponseOutput() DayResponseOutput {
	return o
}

func (o DayResponseOutput) ToDayResponseOutputWithContext(ctx context.Context) DayResponseOutput {
	return o
}

func (o DayResponseOutput) Date() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DayResponse) *int { return v.Date }).(pulumi.IntPtrOutput)
}

func (o DayResponseOutput) IsLast() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DayResponse) *bool { return v.IsLast }).(pulumi.BoolPtrOutput)
}

type DayResponseArrayOutput struct{ *pulumi.OutputState }

func (DayResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DayResponse)(nil)).Elem()
}

func (o DayResponseArrayOutput) ToDayResponseArrayOutput() DayResponseArrayOutput {
	return o
}

func (o DayResponseArrayOutput) ToDayResponseArrayOutputWithContext(ctx context.Context) DayResponseArrayOutput {
	return o
}

func (o DayResponseArrayOutput) Index(i pulumi.IntInput) DayResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DayResponse {
		return vs[0].([]DayResponse)[vs[1].(int)]
	}).(DayResponseOutput)
}

type IdentityData struct {
	Type string `pulumi:"type"`
}





type IdentityDataInput interface {
	pulumi.Input

	ToIdentityDataOutput() IdentityDataOutput
	ToIdentityDataOutputWithContext(context.Context) IdentityDataOutput
}

type IdentityDataArgs struct {
	Type pulumi.StringInput `pulumi:"type"`
}

func (IdentityDataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityData)(nil)).Elem()
}

func (i IdentityDataArgs) ToIdentityDataOutput() IdentityDataOutput {
	return i.ToIdentityDataOutputWithContext(context.Background())
}

func (i IdentityDataArgs) ToIdentityDataOutputWithContext(ctx context.Context) IdentityDataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataOutput)
}

func (i IdentityDataArgs) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return i.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (i IdentityDataArgs) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataOutput).ToIdentityDataPtrOutputWithContext(ctx)
}









type IdentityDataPtrInput interface {
	pulumi.Input

	ToIdentityDataPtrOutput() IdentityDataPtrOutput
	ToIdentityDataPtrOutputWithContext(context.Context) IdentityDataPtrOutput
}

type identityDataPtrType IdentityDataArgs

func IdentityDataPtr(v *IdentityDataArgs) IdentityDataPtrInput {
	return (*identityDataPtrType)(v)
}

func (*identityDataPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityData)(nil)).Elem()
}

func (i *identityDataPtrType) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return i.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (i *identityDataPtrType) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataPtrOutput)
}

type IdentityDataOutput struct{ *pulumi.OutputState }

func (IdentityDataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityData)(nil)).Elem()
}

func (o IdentityDataOutput) ToIdentityDataOutput() IdentityDataOutput {
	return o
}

func (o IdentityDataOutput) ToIdentityDataOutputWithContext(ctx context.Context) IdentityDataOutput {
	return o
}

func (o IdentityDataOutput) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return o.ToIdentityDataPtrOutputWithContext(context.Background())
}

func (o IdentityDataOutput) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityData) *IdentityData {
		return &v
	}).(IdentityDataPtrOutput)
}

func (o IdentityDataOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityData) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityDataPtrOutput struct{ *pulumi.OutputState }

func (IdentityDataPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityData)(nil)).Elem()
}

func (o IdentityDataPtrOutput) ToIdentityDataPtrOutput() IdentityDataPtrOutput {
	return o
}

func (o IdentityDataPtrOutput) ToIdentityDataPtrOutputWithContext(ctx context.Context) IdentityDataPtrOutput {
	return o
}

func (o IdentityDataPtrOutput) Elem() IdentityDataOutput {
	return o.ApplyT(func(v *IdentityData) IdentityData {
		if v != nil {
			return *v
		}
		var ret IdentityData
		return ret
	}).(IdentityDataOutput)
}

func (o IdentityDataPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityData) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type IdentityDataResponse struct {
	PrincipalId string `pulumi:"principalId"`
	TenantId    string `pulumi:"tenantId"`
	Type        string `pulumi:"type"`
}





type IdentityDataResponseInput interface {
	pulumi.Input

	ToIdentityDataResponseOutput() IdentityDataResponseOutput
	ToIdentityDataResponseOutputWithContext(context.Context) IdentityDataResponseOutput
}

type IdentityDataResponseArgs struct {
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
	TenantId    pulumi.StringInput `pulumi:"tenantId"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (IdentityDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityDataResponse)(nil)).Elem()
}

func (i IdentityDataResponseArgs) ToIdentityDataResponseOutput() IdentityDataResponseOutput {
	return i.ToIdentityDataResponseOutputWithContext(context.Background())
}

func (i IdentityDataResponseArgs) ToIdentityDataResponseOutputWithContext(ctx context.Context) IdentityDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataResponseOutput)
}

func (i IdentityDataResponseArgs) ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput {
	return i.ToIdentityDataResponsePtrOutputWithContext(context.Background())
}

func (i IdentityDataResponseArgs) ToIdentityDataResponsePtrOutputWithContext(ctx context.Context) IdentityDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataResponseOutput).ToIdentityDataResponsePtrOutputWithContext(ctx)
}









type IdentityDataResponsePtrInput interface {
	pulumi.Input

	ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput
	ToIdentityDataResponsePtrOutputWithContext(context.Context) IdentityDataResponsePtrOutput
}

type identityDataResponsePtrType IdentityDataResponseArgs

func IdentityDataResponsePtr(v *IdentityDataResponseArgs) IdentityDataResponsePtrInput {
	return (*identityDataResponsePtrType)(v)
}

func (*identityDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityDataResponse)(nil)).Elem()
}

func (i *identityDataResponsePtrType) ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput {
	return i.ToIdentityDataResponsePtrOutputWithContext(context.Background())
}

func (i *identityDataResponsePtrType) ToIdentityDataResponsePtrOutputWithContext(ctx context.Context) IdentityDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentityDataResponsePtrOutput)
}

type IdentityDataResponseOutput struct{ *pulumi.OutputState }

func (IdentityDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentityDataResponse)(nil)).Elem()
}

func (o IdentityDataResponseOutput) ToIdentityDataResponseOutput() IdentityDataResponseOutput {
	return o
}

func (o IdentityDataResponseOutput) ToIdentityDataResponseOutputWithContext(ctx context.Context) IdentityDataResponseOutput {
	return o
}

func (o IdentityDataResponseOutput) ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput {
	return o.ToIdentityDataResponsePtrOutputWithContext(context.Background())
}

func (o IdentityDataResponseOutput) ToIdentityDataResponsePtrOutputWithContext(ctx context.Context) IdentityDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IdentityDataResponse) *IdentityDataResponse {
		return &v
	}).(IdentityDataResponsePtrOutput)
}

func (o IdentityDataResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o IdentityDataResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentityDataResponse) string { return v.Type }).(pulumi.StringOutput)
}

type IdentityDataResponsePtrOutput struct{ *pulumi.OutputState }

func (IdentityDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IdentityDataResponse)(nil)).Elem()
}

func (o IdentityDataResponsePtrOutput) ToIdentityDataResponsePtrOutput() IdentityDataResponsePtrOutput {
	return o
}

func (o IdentityDataResponsePtrOutput) ToIdentityDataResponsePtrOutputWithContext(ctx context.Context) IdentityDataResponsePtrOutput {
	return o
}

func (o IdentityDataResponsePtrOutput) Elem() IdentityDataResponseOutput {
	return o.ApplyT(func(v *IdentityDataResponse) IdentityDataResponse {
		if v != nil {
			return *v
		}
		var ret IdentityDataResponse
		return ret
	}).(IdentityDataResponseOutput)
}

func (o IdentityDataResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o IdentityDataResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IdentityDataResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type LongTermRetentionPolicy struct {
	DailySchedule       *DailyRetentionSchedule   `pulumi:"dailySchedule"`
	MonthlySchedule     *MonthlyRetentionSchedule `pulumi:"monthlySchedule"`
	RetentionPolicyType *string                   `pulumi:"retentionPolicyType"`
	WeeklySchedule      *WeeklyRetentionSchedule  `pulumi:"weeklySchedule"`
	YearlySchedule      *YearlyRetentionSchedule  `pulumi:"yearlySchedule"`
}





type LongTermRetentionPolicyInput interface {
	pulumi.Input

	ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput
	ToLongTermRetentionPolicyOutputWithContext(context.Context) LongTermRetentionPolicyOutput
}

type LongTermRetentionPolicyArgs struct {
	DailySchedule       DailyRetentionSchedulePtrInput   `pulumi:"dailySchedule"`
	MonthlySchedule     MonthlyRetentionSchedulePtrInput `pulumi:"monthlySchedule"`
	RetentionPolicyType pulumi.StringPtrInput            `pulumi:"retentionPolicyType"`
	WeeklySchedule      WeeklyRetentionSchedulePtrInput  `pulumi:"weeklySchedule"`
	YearlySchedule      YearlyRetentionSchedulePtrInput  `pulumi:"yearlySchedule"`
}

func (LongTermRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicy)(nil)).Elem()
}

func (i LongTermRetentionPolicyArgs) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return i.ToLongTermRetentionPolicyOutputWithContext(context.Background())
}

func (i LongTermRetentionPolicyArgs) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermRetentionPolicyOutput)
}

type LongTermRetentionPolicyOutput struct{ *pulumi.OutputState }

func (LongTermRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicy)(nil)).Elem()
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutput() LongTermRetentionPolicyOutput {
	return o
}

func (o LongTermRetentionPolicyOutput) ToLongTermRetentionPolicyOutputWithContext(ctx context.Context) LongTermRetentionPolicyOutput {
	return o
}

func (o LongTermRetentionPolicyOutput) DailySchedule() DailyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *DailyRetentionSchedule { return v.DailySchedule }).(DailyRetentionSchedulePtrOutput)
}

func (o LongTermRetentionPolicyOutput) MonthlySchedule() MonthlyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *MonthlyRetentionSchedule { return v.MonthlySchedule }).(MonthlyRetentionSchedulePtrOutput)
}

func (o LongTermRetentionPolicyOutput) RetentionPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *string { return v.RetentionPolicyType }).(pulumi.StringPtrOutput)
}

func (o LongTermRetentionPolicyOutput) WeeklySchedule() WeeklyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *WeeklyRetentionSchedule { return v.WeeklySchedule }).(WeeklyRetentionSchedulePtrOutput)
}

func (o LongTermRetentionPolicyOutput) YearlySchedule() YearlyRetentionSchedulePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicy) *YearlyRetentionSchedule { return v.YearlySchedule }).(YearlyRetentionSchedulePtrOutput)
}

type LongTermRetentionPolicyResponse struct {
	DailySchedule       *DailyRetentionScheduleResponse   `pulumi:"dailySchedule"`
	MonthlySchedule     *MonthlyRetentionScheduleResponse `pulumi:"monthlySchedule"`
	RetentionPolicyType *string                           `pulumi:"retentionPolicyType"`
	WeeklySchedule      *WeeklyRetentionScheduleResponse  `pulumi:"weeklySchedule"`
	YearlySchedule      *YearlyRetentionScheduleResponse  `pulumi:"yearlySchedule"`
}





type LongTermRetentionPolicyResponseInput interface {
	pulumi.Input

	ToLongTermRetentionPolicyResponseOutput() LongTermRetentionPolicyResponseOutput
	ToLongTermRetentionPolicyResponseOutputWithContext(context.Context) LongTermRetentionPolicyResponseOutput
}

type LongTermRetentionPolicyResponseArgs struct {
	DailySchedule       DailyRetentionScheduleResponsePtrInput   `pulumi:"dailySchedule"`
	MonthlySchedule     MonthlyRetentionScheduleResponsePtrInput `pulumi:"monthlySchedule"`
	RetentionPolicyType pulumi.StringPtrInput                    `pulumi:"retentionPolicyType"`
	WeeklySchedule      WeeklyRetentionScheduleResponsePtrInput  `pulumi:"weeklySchedule"`
	YearlySchedule      YearlyRetentionScheduleResponsePtrInput  `pulumi:"yearlySchedule"`
}

func (LongTermRetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicyResponse)(nil)).Elem()
}

func (i LongTermRetentionPolicyResponseArgs) ToLongTermRetentionPolicyResponseOutput() LongTermRetentionPolicyResponseOutput {
	return i.ToLongTermRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i LongTermRetentionPolicyResponseArgs) ToLongTermRetentionPolicyResponseOutputWithContext(ctx context.Context) LongTermRetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermRetentionPolicyResponseOutput)
}

type LongTermRetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (LongTermRetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermRetentionPolicyResponse)(nil)).Elem()
}

func (o LongTermRetentionPolicyResponseOutput) ToLongTermRetentionPolicyResponseOutput() LongTermRetentionPolicyResponseOutput {
	return o
}

func (o LongTermRetentionPolicyResponseOutput) ToLongTermRetentionPolicyResponseOutputWithContext(ctx context.Context) LongTermRetentionPolicyResponseOutput {
	return o
}

func (o LongTermRetentionPolicyResponseOutput) DailySchedule() DailyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *DailyRetentionScheduleResponse { return v.DailySchedule }).(DailyRetentionScheduleResponsePtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) MonthlySchedule() MonthlyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *MonthlyRetentionScheduleResponse { return v.MonthlySchedule }).(MonthlyRetentionScheduleResponsePtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) RetentionPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *string { return v.RetentionPolicyType }).(pulumi.StringPtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) WeeklySchedule() WeeklyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *WeeklyRetentionScheduleResponse { return v.WeeklySchedule }).(WeeklyRetentionScheduleResponsePtrOutput)
}

func (o LongTermRetentionPolicyResponseOutput) YearlySchedule() YearlyRetentionScheduleResponsePtrOutput {
	return o.ApplyT(func(v LongTermRetentionPolicyResponse) *YearlyRetentionScheduleResponse { return v.YearlySchedule }).(YearlyRetentionScheduleResponsePtrOutput)
}

type LongTermSchedulePolicy struct {
	SchedulePolicyType *string `pulumi:"schedulePolicyType"`
}





type LongTermSchedulePolicyInput interface {
	pulumi.Input

	ToLongTermSchedulePolicyOutput() LongTermSchedulePolicyOutput
	ToLongTermSchedulePolicyOutputWithContext(context.Context) LongTermSchedulePolicyOutput
}

type LongTermSchedulePolicyArgs struct {
	SchedulePolicyType pulumi.StringPtrInput `pulumi:"schedulePolicyType"`
}

func (LongTermSchedulePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicy)(nil)).Elem()
}

func (i LongTermSchedulePolicyArgs) ToLongTermSchedulePolicyOutput() LongTermSchedulePolicyOutput {
	return i.ToLongTermSchedulePolicyOutputWithContext(context.Background())
}

func (i LongTermSchedulePolicyArgs) ToLongTermSchedulePolicyOutputWithContext(ctx context.Context) LongTermSchedulePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermSchedulePolicyOutput)
}

type LongTermSchedulePolicyOutput struct{ *pulumi.OutputState }

func (LongTermSchedulePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicy)(nil)).Elem()
}

func (o LongTermSchedulePolicyOutput) ToLongTermSchedulePolicyOutput() LongTermSchedulePolicyOutput {
	return o
}

func (o LongTermSchedulePolicyOutput) ToLongTermSchedulePolicyOutputWithContext(ctx context.Context) LongTermSchedulePolicyOutput {
	return o
}

func (o LongTermSchedulePolicyOutput) SchedulePolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LongTermSchedulePolicy) *string { return v.SchedulePolicyType }).(pulumi.StringPtrOutput)
}

type LongTermSchedulePolicyResponse struct {
	SchedulePolicyType *string `pulumi:"schedulePolicyType"`
}





type LongTermSchedulePolicyResponseInput interface {
	pulumi.Input

	ToLongTermSchedulePolicyResponseOutput() LongTermSchedulePolicyResponseOutput
	ToLongTermSchedulePolicyResponseOutputWithContext(context.Context) LongTermSchedulePolicyResponseOutput
}

type LongTermSchedulePolicyResponseArgs struct {
	SchedulePolicyType pulumi.StringPtrInput `pulumi:"schedulePolicyType"`
}

func (LongTermSchedulePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicyResponse)(nil)).Elem()
}

func (i LongTermSchedulePolicyResponseArgs) ToLongTermSchedulePolicyResponseOutput() LongTermSchedulePolicyResponseOutput {
	return i.ToLongTermSchedulePolicyResponseOutputWithContext(context.Background())
}

func (i LongTermSchedulePolicyResponseArgs) ToLongTermSchedulePolicyResponseOutputWithContext(ctx context.Context) LongTermSchedulePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LongTermSchedulePolicyResponseOutput)
}

type LongTermSchedulePolicyResponseOutput struct{ *pulumi.OutputState }

func (LongTermSchedulePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LongTermSchedulePolicyResponse)(nil)).Elem()
}

func (o LongTermSchedulePolicyResponseOutput) ToLongTermSchedulePolicyResponseOutput() LongTermSchedulePolicyResponseOutput {
	return o
}

func (o LongTermSchedulePolicyResponseOutput) ToLongTermSchedulePolicyResponseOutputWithContext(ctx context.Context) LongTermSchedulePolicyResponseOutput {
	return o
}

func (o LongTermSchedulePolicyResponseOutput) SchedulePolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LongTermSchedulePolicyResponse) *string { return v.SchedulePolicyType }).(pulumi.StringPtrOutput)
}

type MabProtectionPolicy struct {
	BackupManagementType *string     `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
}





type MabProtectionPolicyInput interface {
	pulumi.Input

	ToMabProtectionPolicyOutput() MabProtectionPolicyOutput
	ToMabProtectionPolicyOutputWithContext(context.Context) MabProtectionPolicyOutput
}

type MabProtectionPolicyArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ProtectedItemsCount  pulumi.IntPtrInput    `pulumi:"protectedItemsCount"`
	RetentionPolicy      pulumi.Input          `pulumi:"retentionPolicy"`
	SchedulePolicy       pulumi.Input          `pulumi:"schedulePolicy"`
}

func (MabProtectionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicy)(nil)).Elem()
}

func (i MabProtectionPolicyArgs) ToMabProtectionPolicyOutput() MabProtectionPolicyOutput {
	return i.ToMabProtectionPolicyOutputWithContext(context.Background())
}

func (i MabProtectionPolicyArgs) ToMabProtectionPolicyOutputWithContext(ctx context.Context) MabProtectionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabProtectionPolicyOutput)
}

type MabProtectionPolicyOutput struct{ *pulumi.OutputState }

func (MabProtectionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicy)(nil)).Elem()
}

func (o MabProtectionPolicyOutput) ToMabProtectionPolicyOutput() MabProtectionPolicyOutput {
	return o
}

func (o MabProtectionPolicyOutput) ToMabProtectionPolicyOutputWithContext(ctx context.Context) MabProtectionPolicyOutput {
	return o
}

func (o MabProtectionPolicyOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabProtectionPolicy) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o MabProtectionPolicyOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MabProtectionPolicy) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o MabProtectionPolicyOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicy) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o MabProtectionPolicyOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicy) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type MabProtectionPolicyResponse struct {
	BackupManagementType *string     `pulumi:"backupManagementType"`
	ProtectedItemsCount  *int        `pulumi:"protectedItemsCount"`
	RetentionPolicy      interface{} `pulumi:"retentionPolicy"`
	SchedulePolicy       interface{} `pulumi:"schedulePolicy"`
}





type MabProtectionPolicyResponseInput interface {
	pulumi.Input

	ToMabProtectionPolicyResponseOutput() MabProtectionPolicyResponseOutput
	ToMabProtectionPolicyResponseOutputWithContext(context.Context) MabProtectionPolicyResponseOutput
}

type MabProtectionPolicyResponseArgs struct {
	BackupManagementType pulumi.StringPtrInput `pulumi:"backupManagementType"`
	ProtectedItemsCount  pulumi.IntPtrInput    `pulumi:"protectedItemsCount"`
	RetentionPolicy      pulumi.Input          `pulumi:"retentionPolicy"`
	SchedulePolicy       pulumi.Input          `pulumi:"schedulePolicy"`
}

func (MabProtectionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicyResponse)(nil)).Elem()
}

func (i MabProtectionPolicyResponseArgs) ToMabProtectionPolicyResponseOutput() MabProtectionPolicyResponseOutput {
	return i.ToMabProtectionPolicyResponseOutputWithContext(context.Background())
}

func (i MabProtectionPolicyResponseArgs) ToMabProtectionPolicyResponseOutputWithContext(ctx context.Context) MabProtectionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MabProtectionPolicyResponseOutput)
}

type MabProtectionPolicyResponseOutput struct{ *pulumi.OutputState }

func (MabProtectionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MabProtectionPolicyResponse)(nil)).Elem()
}

func (o MabProtectionPolicyResponseOutput) ToMabProtectionPolicyResponseOutput() MabProtectionPolicyResponseOutput {
	return o
}

func (o MabProtectionPolicyResponseOutput) ToMabProtectionPolicyResponseOutputWithContext(ctx context.Context) MabProtectionPolicyResponseOutput {
	return o
}

func (o MabProtectionPolicyResponseOutput) BackupManagementType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) *string { return v.BackupManagementType }).(pulumi.StringPtrOutput)
}

func (o MabProtectionPolicyResponseOutput) ProtectedItemsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) *int { return v.ProtectedItemsCount }).(pulumi.IntPtrOutput)
}

func (o MabProtectionPolicyResponseOutput) RetentionPolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) interface{} { return v.RetentionPolicy }).(pulumi.AnyOutput)
}

func (o MabProtectionPolicyResponseOutput) SchedulePolicy() pulumi.AnyOutput {
	return o.ApplyT(func(v MabProtectionPolicyResponse) interface{} { return v.SchedulePolicy }).(pulumi.AnyOutput)
}

type MonthlyRetentionSchedule struct {
	RetentionDuration           *RetentionDuration       `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormat    `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *RetentionScheduleFormat `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormat   `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                 `pulumi:"retentionTimes"`
}





type MonthlyRetentionScheduleInput interface {
	pulumi.Input

	ToMonthlyRetentionScheduleOutput() MonthlyRetentionScheduleOutput
	ToMonthlyRetentionScheduleOutputWithContext(context.Context) MonthlyRetentionScheduleOutput
}

type MonthlyRetentionScheduleArgs struct {
	RetentionDuration           RetentionDurationPtrInput       `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatPtrInput    `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType RetentionScheduleFormatPtrInput `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatPtrInput   `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput         `pulumi:"retentionTimes"`
}

func (MonthlyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionSchedule)(nil)).Elem()
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionScheduleOutput() MonthlyRetentionScheduleOutput {
	return i.ToMonthlyRetentionScheduleOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionScheduleOutputWithContext(ctx context.Context) MonthlyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleOutput)
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return i.ToMonthlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleArgs) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleOutput).ToMonthlyRetentionSchedulePtrOutputWithContext(ctx)
}









type MonthlyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput
	ToMonthlyRetentionSchedulePtrOutputWithContext(context.Context) MonthlyRetentionSchedulePtrOutput
}

type monthlyRetentionSchedulePtrType MonthlyRetentionScheduleArgs

func MonthlyRetentionSchedulePtr(v *MonthlyRetentionScheduleArgs) MonthlyRetentionSchedulePtrInput {
	return (*monthlyRetentionSchedulePtrType)(v)
}

func (*monthlyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionSchedule)(nil)).Elem()
}

func (i *monthlyRetentionSchedulePtrType) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return i.ToMonthlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *monthlyRetentionSchedulePtrType) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionSchedulePtrOutput)
}

type MonthlyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionSchedule)(nil)).Elem()
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionScheduleOutput() MonthlyRetentionScheduleOutput {
	return o
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionScheduleOutputWithContext(ctx context.Context) MonthlyRetentionScheduleOutput {
	return o
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return o.ToMonthlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o MonthlyRetentionScheduleOutput) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthlyRetentionSchedule) *MonthlyRetentionSchedule {
		return &v
	}).(MonthlyRetentionSchedulePtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *DailyRetentionFormat { return v.RetentionScheduleDaily }).(DailyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionScheduleFormatType() RetentionScheduleFormatPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *RetentionScheduleFormat { return v.RetentionScheduleFormatType }).(RetentionScheduleFormatPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) *WeeklyRetentionFormat { return v.RetentionScheduleWeekly }).(WeeklyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MonthlyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type MonthlyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionSchedule)(nil)).Elem()
}

func (o MonthlyRetentionSchedulePtrOutput) ToMonthlyRetentionSchedulePtrOutput() MonthlyRetentionSchedulePtrOutput {
	return o
}

func (o MonthlyRetentionSchedulePtrOutput) ToMonthlyRetentionSchedulePtrOutputWithContext(ctx context.Context) MonthlyRetentionSchedulePtrOutput {
	return o
}

func (o MonthlyRetentionSchedulePtrOutput) Elem() MonthlyRetentionScheduleOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) MonthlyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret MonthlyRetentionSchedule
		return ret
	}).(MonthlyRetentionScheduleOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *DailyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionScheduleFormatType() RetentionScheduleFormatPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *RetentionScheduleFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(RetentionScheduleFormatPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) *WeeklyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatPtrOutput)
}

func (o MonthlyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonthlyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type MonthlyRetentionScheduleResponse struct {
	RetentionDuration           *RetentionDurationResponse     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormatResponse  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                        `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormatResponse `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                       `pulumi:"retentionTimes"`
}





type MonthlyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToMonthlyRetentionScheduleResponseOutput() MonthlyRetentionScheduleResponseOutput
	ToMonthlyRetentionScheduleResponseOutputWithContext(context.Context) MonthlyRetentionScheduleResponseOutput
}

type MonthlyRetentionScheduleResponseArgs struct {
	RetentionDuration           RetentionDurationResponsePtrInput     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatResponsePtrInput  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType pulumi.StringPtrInput                 `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatResponsePtrInput `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput               `pulumi:"retentionTimes"`
}

func (MonthlyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponseOutput() MonthlyRetentionScheduleResponseOutput {
	return i.ToMonthlyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponseOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleResponseOutput)
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return i.ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i MonthlyRetentionScheduleResponseArgs) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleResponseOutput).ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type MonthlyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput
	ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Context) MonthlyRetentionScheduleResponsePtrOutput
}

type monthlyRetentionScheduleResponsePtrType MonthlyRetentionScheduleResponseArgs

func MonthlyRetentionScheduleResponsePtr(v *MonthlyRetentionScheduleResponseArgs) MonthlyRetentionScheduleResponsePtrInput {
	return (*monthlyRetentionScheduleResponsePtrType)(v)
}

func (*monthlyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (i *monthlyRetentionScheduleResponsePtrType) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return i.ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *monthlyRetentionScheduleResponsePtrType) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MonthlyRetentionScheduleResponsePtrOutput)
}

type MonthlyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponseOutput() MonthlyRetentionScheduleResponseOutput {
	return o
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponseOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponseOutput {
	return o
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return o.ToMonthlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o MonthlyRetentionScheduleResponseOutput) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MonthlyRetentionScheduleResponse) *MonthlyRetentionScheduleResponse {
		return &v
	}).(MonthlyRetentionScheduleResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *DailyRetentionFormatResponse {
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *string { return v.RetentionScheduleFormatType }).(pulumi.StringPtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MonthlyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type MonthlyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (MonthlyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MonthlyRetentionScheduleResponse)(nil)).Elem()
}

func (o MonthlyRetentionScheduleResponsePtrOutput) ToMonthlyRetentionScheduleResponsePtrOutput() MonthlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o MonthlyRetentionScheduleResponsePtrOutput) ToMonthlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) MonthlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o MonthlyRetentionScheduleResponsePtrOutput) Elem() MonthlyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) MonthlyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret MonthlyRetentionScheduleResponse
		return ret
	}).(MonthlyRetentionScheduleResponseOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *DailyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(pulumi.StringPtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o MonthlyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MonthlyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type PrivateEndpointConnectionVaultPropertiesResponse struct {
	Id         string                                 `pulumi:"id"`
	Properties VaultPrivateEndpointConnectionResponse `pulumi:"properties"`
}





type PrivateEndpointConnectionVaultPropertiesResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionVaultPropertiesResponseOutput() PrivateEndpointConnectionVaultPropertiesResponseOutput
	ToPrivateEndpointConnectionVaultPropertiesResponseOutputWithContext(context.Context) PrivateEndpointConnectionVaultPropertiesResponseOutput
}

type PrivateEndpointConnectionVaultPropertiesResponseArgs struct {
	Id         pulumi.StringInput                          `pulumi:"id"`
	Properties VaultPrivateEndpointConnectionResponseInput `pulumi:"properties"`
}

func (PrivateEndpointConnectionVaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionVaultPropertiesResponseArgs) ToPrivateEndpointConnectionVaultPropertiesResponseOutput() PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return i.ToPrivateEndpointConnectionVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionVaultPropertiesResponseArgs) ToPrivateEndpointConnectionVaultPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionVaultPropertiesResponseOutput)
}





type PrivateEndpointConnectionVaultPropertiesResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutput() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput
	ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionVaultPropertiesResponseArrayOutput
}

type PrivateEndpointConnectionVaultPropertiesResponseArray []PrivateEndpointConnectionVaultPropertiesResponseInput

func (PrivateEndpointConnectionVaultPropertiesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionVaultPropertiesResponseArray) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutput() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return i.ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionVaultPropertiesResponseArray) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput)
}

type PrivateEndpointConnectionVaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionVaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) ToPrivateEndpointConnectionVaultPropertiesResponseOutput() PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) ToPrivateEndpointConnectionVaultPropertiesResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionVaultPropertiesResponseOutput) Properties() VaultPrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionVaultPropertiesResponse) VaultPrivateEndpointConnectionResponse {
		return v.Properties
	}).(VaultPrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointConnectionVaultPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionVaultPropertiesResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutput() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) ToPrivateEndpointConnectionVaultPropertiesResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionVaultPropertiesResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionVaultPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionVaultPropertiesResponse {
		return vs[0].([]PrivateEndpointConnectionVaultPropertiesResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionVaultPropertiesResponseOutput)
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}





type PrivateEndpointResponseInput interface {
	pulumi.Input

	ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput
	ToPrivateEndpointResponseOutputWithContext(context.Context) PrivateEndpointResponseOutput
}

type PrivateEndpointResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (PrivateEndpointResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return i.ToPrivateEndpointResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointResponseArgs) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointResponseOutput)
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type RetentionDuration struct {
	Count        *int                   `pulumi:"count"`
	DurationType *RetentionDurationType `pulumi:"durationType"`
}





type RetentionDurationInput interface {
	pulumi.Input

	ToRetentionDurationOutput() RetentionDurationOutput
	ToRetentionDurationOutputWithContext(context.Context) RetentionDurationOutput
}

type RetentionDurationArgs struct {
	Count        pulumi.IntPtrInput            `pulumi:"count"`
	DurationType RetentionDurationTypePtrInput `pulumi:"durationType"`
}

func (RetentionDurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDuration)(nil)).Elem()
}

func (i RetentionDurationArgs) ToRetentionDurationOutput() RetentionDurationOutput {
	return i.ToRetentionDurationOutputWithContext(context.Background())
}

func (i RetentionDurationArgs) ToRetentionDurationOutputWithContext(ctx context.Context) RetentionDurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationOutput)
}

func (i RetentionDurationArgs) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return i.ToRetentionDurationPtrOutputWithContext(context.Background())
}

func (i RetentionDurationArgs) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationOutput).ToRetentionDurationPtrOutputWithContext(ctx)
}









type RetentionDurationPtrInput interface {
	pulumi.Input

	ToRetentionDurationPtrOutput() RetentionDurationPtrOutput
	ToRetentionDurationPtrOutputWithContext(context.Context) RetentionDurationPtrOutput
}

type retentionDurationPtrType RetentionDurationArgs

func RetentionDurationPtr(v *RetentionDurationArgs) RetentionDurationPtrInput {
	return (*retentionDurationPtrType)(v)
}

func (*retentionDurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDuration)(nil)).Elem()
}

func (i *retentionDurationPtrType) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return i.ToRetentionDurationPtrOutputWithContext(context.Background())
}

func (i *retentionDurationPtrType) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationPtrOutput)
}

type RetentionDurationOutput struct{ *pulumi.OutputState }

func (RetentionDurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDuration)(nil)).Elem()
}

func (o RetentionDurationOutput) ToRetentionDurationOutput() RetentionDurationOutput {
	return o
}

func (o RetentionDurationOutput) ToRetentionDurationOutputWithContext(ctx context.Context) RetentionDurationOutput {
	return o
}

func (o RetentionDurationOutput) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return o.ToRetentionDurationPtrOutputWithContext(context.Background())
}

func (o RetentionDurationOutput) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionDuration) *RetentionDuration {
		return &v
	}).(RetentionDurationPtrOutput)
}

func (o RetentionDurationOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionDuration) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RetentionDurationOutput) DurationType() RetentionDurationTypePtrOutput {
	return o.ApplyT(func(v RetentionDuration) *RetentionDurationType { return v.DurationType }).(RetentionDurationTypePtrOutput)
}

type RetentionDurationPtrOutput struct{ *pulumi.OutputState }

func (RetentionDurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDuration)(nil)).Elem()
}

func (o RetentionDurationPtrOutput) ToRetentionDurationPtrOutput() RetentionDurationPtrOutput {
	return o
}

func (o RetentionDurationPtrOutput) ToRetentionDurationPtrOutputWithContext(ctx context.Context) RetentionDurationPtrOutput {
	return o
}

func (o RetentionDurationPtrOutput) Elem() RetentionDurationOutput {
	return o.ApplyT(func(v *RetentionDuration) RetentionDuration {
		if v != nil {
			return *v
		}
		var ret RetentionDuration
		return ret
	}).(RetentionDurationOutput)
}

func (o RetentionDurationPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionDuration) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RetentionDurationPtrOutput) DurationType() RetentionDurationTypePtrOutput {
	return o.ApplyT(func(v *RetentionDuration) *RetentionDurationType {
		if v == nil {
			return nil
		}
		return v.DurationType
	}).(RetentionDurationTypePtrOutput)
}

type RetentionDurationResponse struct {
	Count        *int    `pulumi:"count"`
	DurationType *string `pulumi:"durationType"`
}





type RetentionDurationResponseInput interface {
	pulumi.Input

	ToRetentionDurationResponseOutput() RetentionDurationResponseOutput
	ToRetentionDurationResponseOutputWithContext(context.Context) RetentionDurationResponseOutput
}

type RetentionDurationResponseArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	DurationType pulumi.StringPtrInput `pulumi:"durationType"`
}

func (RetentionDurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationResponse)(nil)).Elem()
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponseOutput() RetentionDurationResponseOutput {
	return i.ToRetentionDurationResponseOutputWithContext(context.Background())
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponseOutputWithContext(ctx context.Context) RetentionDurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationResponseOutput)
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return i.ToRetentionDurationResponsePtrOutputWithContext(context.Background())
}

func (i RetentionDurationResponseArgs) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationResponseOutput).ToRetentionDurationResponsePtrOutputWithContext(ctx)
}









type RetentionDurationResponsePtrInput interface {
	pulumi.Input

	ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput
	ToRetentionDurationResponsePtrOutputWithContext(context.Context) RetentionDurationResponsePtrOutput
}

type retentionDurationResponsePtrType RetentionDurationResponseArgs

func RetentionDurationResponsePtr(v *RetentionDurationResponseArgs) RetentionDurationResponsePtrInput {
	return (*retentionDurationResponsePtrType)(v)
}

func (*retentionDurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDurationResponse)(nil)).Elem()
}

func (i *retentionDurationResponsePtrType) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return i.ToRetentionDurationResponsePtrOutputWithContext(context.Background())
}

func (i *retentionDurationResponsePtrType) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionDurationResponsePtrOutput)
}

type RetentionDurationResponseOutput struct{ *pulumi.OutputState }

func (RetentionDurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionDurationResponse)(nil)).Elem()
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponseOutput() RetentionDurationResponseOutput {
	return o
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponseOutputWithContext(ctx context.Context) RetentionDurationResponseOutput {
	return o
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return o.ToRetentionDurationResponsePtrOutputWithContext(context.Background())
}

func (o RetentionDurationResponseOutput) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionDurationResponse) *RetentionDurationResponse {
		return &v
	}).(RetentionDurationResponsePtrOutput)
}

func (o RetentionDurationResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RetentionDurationResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RetentionDurationResponseOutput) DurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RetentionDurationResponse) *string { return v.DurationType }).(pulumi.StringPtrOutput)
}

type RetentionDurationResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionDurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionDurationResponse)(nil)).Elem()
}

func (o RetentionDurationResponsePtrOutput) ToRetentionDurationResponsePtrOutput() RetentionDurationResponsePtrOutput {
	return o
}

func (o RetentionDurationResponsePtrOutput) ToRetentionDurationResponsePtrOutputWithContext(ctx context.Context) RetentionDurationResponsePtrOutput {
	return o
}

func (o RetentionDurationResponsePtrOutput) Elem() RetentionDurationResponseOutput {
	return o.ApplyT(func(v *RetentionDurationResponse) RetentionDurationResponse {
		if v != nil {
			return *v
		}
		var ret RetentionDurationResponse
		return ret
	}).(RetentionDurationResponseOutput)
}

func (o RetentionDurationResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionDurationResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RetentionDurationResponsePtrOutput) DurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RetentionDurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DurationType
	}).(pulumi.StringPtrOutput)
}

type SimpleRetentionPolicy struct {
	RetentionDuration   *RetentionDuration `pulumi:"retentionDuration"`
	RetentionPolicyType *string            `pulumi:"retentionPolicyType"`
}





type SimpleRetentionPolicyInput interface {
	pulumi.Input

	ToSimpleRetentionPolicyOutput() SimpleRetentionPolicyOutput
	ToSimpleRetentionPolicyOutputWithContext(context.Context) SimpleRetentionPolicyOutput
}

type SimpleRetentionPolicyArgs struct {
	RetentionDuration   RetentionDurationPtrInput `pulumi:"retentionDuration"`
	RetentionPolicyType pulumi.StringPtrInput     `pulumi:"retentionPolicyType"`
}

func (SimpleRetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicy)(nil)).Elem()
}

func (i SimpleRetentionPolicyArgs) ToSimpleRetentionPolicyOutput() SimpleRetentionPolicyOutput {
	return i.ToSimpleRetentionPolicyOutputWithContext(context.Background())
}

func (i SimpleRetentionPolicyArgs) ToSimpleRetentionPolicyOutputWithContext(ctx context.Context) SimpleRetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleRetentionPolicyOutput)
}

type SimpleRetentionPolicyOutput struct{ *pulumi.OutputState }

func (SimpleRetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicy)(nil)).Elem()
}

func (o SimpleRetentionPolicyOutput) ToSimpleRetentionPolicyOutput() SimpleRetentionPolicyOutput {
	return o
}

func (o SimpleRetentionPolicyOutput) ToSimpleRetentionPolicyOutputWithContext(ctx context.Context) SimpleRetentionPolicyOutput {
	return o
}

func (o SimpleRetentionPolicyOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v SimpleRetentionPolicy) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o SimpleRetentionPolicyOutput) RetentionPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleRetentionPolicy) *string { return v.RetentionPolicyType }).(pulumi.StringPtrOutput)
}

type SimpleRetentionPolicyResponse struct {
	RetentionDuration   *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionPolicyType *string                    `pulumi:"retentionPolicyType"`
}





type SimpleRetentionPolicyResponseInput interface {
	pulumi.Input

	ToSimpleRetentionPolicyResponseOutput() SimpleRetentionPolicyResponseOutput
	ToSimpleRetentionPolicyResponseOutputWithContext(context.Context) SimpleRetentionPolicyResponseOutput
}

type SimpleRetentionPolicyResponseArgs struct {
	RetentionDuration   RetentionDurationResponsePtrInput `pulumi:"retentionDuration"`
	RetentionPolicyType pulumi.StringPtrInput             `pulumi:"retentionPolicyType"`
}

func (SimpleRetentionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicyResponse)(nil)).Elem()
}

func (i SimpleRetentionPolicyResponseArgs) ToSimpleRetentionPolicyResponseOutput() SimpleRetentionPolicyResponseOutput {
	return i.ToSimpleRetentionPolicyResponseOutputWithContext(context.Background())
}

func (i SimpleRetentionPolicyResponseArgs) ToSimpleRetentionPolicyResponseOutputWithContext(ctx context.Context) SimpleRetentionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleRetentionPolicyResponseOutput)
}

type SimpleRetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (SimpleRetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleRetentionPolicyResponse)(nil)).Elem()
}

func (o SimpleRetentionPolicyResponseOutput) ToSimpleRetentionPolicyResponseOutput() SimpleRetentionPolicyResponseOutput {
	return o
}

func (o SimpleRetentionPolicyResponseOutput) ToSimpleRetentionPolicyResponseOutputWithContext(ctx context.Context) SimpleRetentionPolicyResponseOutput {
	return o
}

func (o SimpleRetentionPolicyResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v SimpleRetentionPolicyResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o SimpleRetentionPolicyResponseOutput) RetentionPolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleRetentionPolicyResponse) *string { return v.RetentionPolicyType }).(pulumi.StringPtrOutput)
}

type SimpleSchedulePolicy struct {
	SchedulePolicyType      *string          `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []DayOfWeek      `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *ScheduleRunType `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string         `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int             `pulumi:"scheduleWeeklyFrequency"`
}





type SimpleSchedulePolicyInput interface {
	pulumi.Input

	ToSimpleSchedulePolicyOutput() SimpleSchedulePolicyOutput
	ToSimpleSchedulePolicyOutputWithContext(context.Context) SimpleSchedulePolicyOutput
}

type SimpleSchedulePolicyArgs struct {
	SchedulePolicyType      pulumi.StringPtrInput   `pulumi:"schedulePolicyType"`
	ScheduleRunDays         DayOfWeekArrayInput     `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    ScheduleRunTypePtrInput `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        pulumi.StringArrayInput `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency pulumi.IntPtrInput      `pulumi:"scheduleWeeklyFrequency"`
}

func (SimpleSchedulePolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicy)(nil)).Elem()
}

func (i SimpleSchedulePolicyArgs) ToSimpleSchedulePolicyOutput() SimpleSchedulePolicyOutput {
	return i.ToSimpleSchedulePolicyOutputWithContext(context.Background())
}

func (i SimpleSchedulePolicyArgs) ToSimpleSchedulePolicyOutputWithContext(ctx context.Context) SimpleSchedulePolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleSchedulePolicyOutput)
}

type SimpleSchedulePolicyOutput struct{ *pulumi.OutputState }

func (SimpleSchedulePolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicy)(nil)).Elem()
}

func (o SimpleSchedulePolicyOutput) ToSimpleSchedulePolicyOutput() SimpleSchedulePolicyOutput {
	return o
}

func (o SimpleSchedulePolicyOutput) ToSimpleSchedulePolicyOutputWithContext(ctx context.Context) SimpleSchedulePolicyOutput {
	return o
}

func (o SimpleSchedulePolicyOutput) SchedulePolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) *string { return v.SchedulePolicyType }).(pulumi.StringPtrOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleRunDays() DayOfWeekArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) []DayOfWeek { return v.ScheduleRunDays }).(DayOfWeekArrayOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleRunFrequency() ScheduleRunTypePtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) *ScheduleRunType { return v.ScheduleRunFrequency }).(ScheduleRunTypePtrOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleRunTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) []string { return v.ScheduleRunTimes }).(pulumi.StringArrayOutput)
}

func (o SimpleSchedulePolicyOutput) ScheduleWeeklyFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicy) *int { return v.ScheduleWeeklyFrequency }).(pulumi.IntPtrOutput)
}

type SimpleSchedulePolicyResponse struct {
	SchedulePolicyType      *string  `pulumi:"schedulePolicyType"`
	ScheduleRunDays         []string `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    *string  `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        []string `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency *int     `pulumi:"scheduleWeeklyFrequency"`
}





type SimpleSchedulePolicyResponseInput interface {
	pulumi.Input

	ToSimpleSchedulePolicyResponseOutput() SimpleSchedulePolicyResponseOutput
	ToSimpleSchedulePolicyResponseOutputWithContext(context.Context) SimpleSchedulePolicyResponseOutput
}

type SimpleSchedulePolicyResponseArgs struct {
	SchedulePolicyType      pulumi.StringPtrInput   `pulumi:"schedulePolicyType"`
	ScheduleRunDays         pulumi.StringArrayInput `pulumi:"scheduleRunDays"`
	ScheduleRunFrequency    pulumi.StringPtrInput   `pulumi:"scheduleRunFrequency"`
	ScheduleRunTimes        pulumi.StringArrayInput `pulumi:"scheduleRunTimes"`
	ScheduleWeeklyFrequency pulumi.IntPtrInput      `pulumi:"scheduleWeeklyFrequency"`
}

func (SimpleSchedulePolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicyResponse)(nil)).Elem()
}

func (i SimpleSchedulePolicyResponseArgs) ToSimpleSchedulePolicyResponseOutput() SimpleSchedulePolicyResponseOutput {
	return i.ToSimpleSchedulePolicyResponseOutputWithContext(context.Background())
}

func (i SimpleSchedulePolicyResponseArgs) ToSimpleSchedulePolicyResponseOutputWithContext(ctx context.Context) SimpleSchedulePolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SimpleSchedulePolicyResponseOutput)
}

type SimpleSchedulePolicyResponseOutput struct{ *pulumi.OutputState }

func (SimpleSchedulePolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SimpleSchedulePolicyResponse)(nil)).Elem()
}

func (o SimpleSchedulePolicyResponseOutput) ToSimpleSchedulePolicyResponseOutput() SimpleSchedulePolicyResponseOutput {
	return o
}

func (o SimpleSchedulePolicyResponseOutput) ToSimpleSchedulePolicyResponseOutputWithContext(ctx context.Context) SimpleSchedulePolicyResponseOutput {
	return o
}

func (o SimpleSchedulePolicyResponseOutput) SchedulePolicyType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) *string { return v.SchedulePolicyType }).(pulumi.StringPtrOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleRunDays() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) []string { return v.ScheduleRunDays }).(pulumi.StringArrayOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleRunFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) *string { return v.ScheduleRunFrequency }).(pulumi.StringPtrOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleRunTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) []string { return v.ScheduleRunTimes }).(pulumi.StringArrayOutput)
}

func (o SimpleSchedulePolicyResponseOutput) ScheduleWeeklyFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SimpleSchedulePolicyResponse) *int { return v.ScheduleWeeklyFrequency }).(pulumi.IntPtrOutput)
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
}





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
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

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type UpgradeDetailsResponse struct {
	EndTimeUtc         string `pulumi:"endTimeUtc"`
	LastUpdatedTimeUtc string `pulumi:"lastUpdatedTimeUtc"`
	Message            string `pulumi:"message"`
	OperationId        string `pulumi:"operationId"`
	PreviousResourceId string `pulumi:"previousResourceId"`
	StartTimeUtc       string `pulumi:"startTimeUtc"`
	Status             string `pulumi:"status"`
	TriggerType        string `pulumi:"triggerType"`
	UpgradedResourceId string `pulumi:"upgradedResourceId"`
}





type UpgradeDetailsResponseInput interface {
	pulumi.Input

	ToUpgradeDetailsResponseOutput() UpgradeDetailsResponseOutput
	ToUpgradeDetailsResponseOutputWithContext(context.Context) UpgradeDetailsResponseOutput
}

type UpgradeDetailsResponseArgs struct {
	EndTimeUtc         pulumi.StringInput `pulumi:"endTimeUtc"`
	LastUpdatedTimeUtc pulumi.StringInput `pulumi:"lastUpdatedTimeUtc"`
	Message            pulumi.StringInput `pulumi:"message"`
	OperationId        pulumi.StringInput `pulumi:"operationId"`
	PreviousResourceId pulumi.StringInput `pulumi:"previousResourceId"`
	StartTimeUtc       pulumi.StringInput `pulumi:"startTimeUtc"`
	Status             pulumi.StringInput `pulumi:"status"`
	TriggerType        pulumi.StringInput `pulumi:"triggerType"`
	UpgradedResourceId pulumi.StringInput `pulumi:"upgradedResourceId"`
}

func (UpgradeDetailsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeDetailsResponse)(nil)).Elem()
}

func (i UpgradeDetailsResponseArgs) ToUpgradeDetailsResponseOutput() UpgradeDetailsResponseOutput {
	return i.ToUpgradeDetailsResponseOutputWithContext(context.Background())
}

func (i UpgradeDetailsResponseArgs) ToUpgradeDetailsResponseOutputWithContext(ctx context.Context) UpgradeDetailsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeDetailsResponseOutput)
}

func (i UpgradeDetailsResponseArgs) ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput {
	return i.ToUpgradeDetailsResponsePtrOutputWithContext(context.Background())
}

func (i UpgradeDetailsResponseArgs) ToUpgradeDetailsResponsePtrOutputWithContext(ctx context.Context) UpgradeDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeDetailsResponseOutput).ToUpgradeDetailsResponsePtrOutputWithContext(ctx)
}









type UpgradeDetailsResponsePtrInput interface {
	pulumi.Input

	ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput
	ToUpgradeDetailsResponsePtrOutputWithContext(context.Context) UpgradeDetailsResponsePtrOutput
}

type upgradeDetailsResponsePtrType UpgradeDetailsResponseArgs

func UpgradeDetailsResponsePtr(v *UpgradeDetailsResponseArgs) UpgradeDetailsResponsePtrInput {
	return (*upgradeDetailsResponsePtrType)(v)
}

func (*upgradeDetailsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeDetailsResponse)(nil)).Elem()
}

func (i *upgradeDetailsResponsePtrType) ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput {
	return i.ToUpgradeDetailsResponsePtrOutputWithContext(context.Background())
}

func (i *upgradeDetailsResponsePtrType) ToUpgradeDetailsResponsePtrOutputWithContext(ctx context.Context) UpgradeDetailsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UpgradeDetailsResponsePtrOutput)
}

type UpgradeDetailsResponseOutput struct{ *pulumi.OutputState }

func (UpgradeDetailsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UpgradeDetailsResponse)(nil)).Elem()
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponseOutput() UpgradeDetailsResponseOutput {
	return o
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponseOutputWithContext(ctx context.Context) UpgradeDetailsResponseOutput {
	return o
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput {
	return o.ToUpgradeDetailsResponsePtrOutputWithContext(context.Background())
}

func (o UpgradeDetailsResponseOutput) ToUpgradeDetailsResponsePtrOutputWithContext(ctx context.Context) UpgradeDetailsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UpgradeDetailsResponse) *UpgradeDetailsResponse {
		return &v
	}).(UpgradeDetailsResponsePtrOutput)
}

func (o UpgradeDetailsResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) LastUpdatedTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.LastUpdatedTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.Message }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) OperationId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.OperationId }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) PreviousResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.PreviousResourceId }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) TriggerType() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.TriggerType }).(pulumi.StringOutput)
}

func (o UpgradeDetailsResponseOutput) UpgradedResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v UpgradeDetailsResponse) string { return v.UpgradedResourceId }).(pulumi.StringOutput)
}

type UpgradeDetailsResponsePtrOutput struct{ *pulumi.OutputState }

func (UpgradeDetailsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UpgradeDetailsResponse)(nil)).Elem()
}

func (o UpgradeDetailsResponsePtrOutput) ToUpgradeDetailsResponsePtrOutput() UpgradeDetailsResponsePtrOutput {
	return o
}

func (o UpgradeDetailsResponsePtrOutput) ToUpgradeDetailsResponsePtrOutputWithContext(ctx context.Context) UpgradeDetailsResponsePtrOutput {
	return o
}

func (o UpgradeDetailsResponsePtrOutput) Elem() UpgradeDetailsResponseOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) UpgradeDetailsResponse {
		if v != nil {
			return *v
		}
		var ret UpgradeDetailsResponse
		return ret
	}).(UpgradeDetailsResponseOutput)
}

func (o UpgradeDetailsResponsePtrOutput) EndTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) LastUpdatedTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastUpdatedTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) OperationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.OperationId
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) PreviousResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PreviousResourceId
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) StartTimeUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.StartTimeUtc
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Status
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TriggerType
	}).(pulumi.StringPtrOutput)
}

func (o UpgradeDetailsResponsePtrOutput) UpgradedResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UpgradeDetailsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UpgradedResourceId
	}).(pulumi.StringPtrOutput)
}

type VaultPrivateEndpointConnectionResponse struct {
	PrivateEndpoint                   PrivateEndpointResponse                        `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState VaultPrivateLinkServiceConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                                         `pulumi:"provisioningState"`
}





type VaultPrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToVaultPrivateEndpointConnectionResponseOutput() VaultPrivateEndpointConnectionResponseOutput
	ToVaultPrivateEndpointConnectionResponseOutputWithContext(context.Context) VaultPrivateEndpointConnectionResponseOutput
}

type VaultPrivateEndpointConnectionResponseArgs struct {
	PrivateEndpoint                   PrivateEndpointResponseInput                        `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState VaultPrivateLinkServiceConnectionStateResponseInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringInput                                  `pulumi:"provisioningState"`
}

func (VaultPrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i VaultPrivateEndpointConnectionResponseArgs) ToVaultPrivateEndpointConnectionResponseOutput() VaultPrivateEndpointConnectionResponseOutput {
	return i.ToVaultPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i VaultPrivateEndpointConnectionResponseArgs) ToVaultPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) VaultPrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPrivateEndpointConnectionResponseOutput)
}

type VaultPrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (VaultPrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o VaultPrivateEndpointConnectionResponseOutput) ToVaultPrivateEndpointConnectionResponseOutput() VaultPrivateEndpointConnectionResponseOutput {
	return o
}

func (o VaultPrivateEndpointConnectionResponseOutput) ToVaultPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) VaultPrivateEndpointConnectionResponseOutput {
	return o
}

func (o VaultPrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) PrivateEndpointResponse { return v.PrivateEndpoint }).(PrivateEndpointResponseOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) VaultPrivateLinkServiceConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(VaultPrivateLinkServiceConnectionStateResponseOutput)
}

func (o VaultPrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type VaultPrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired string `pulumi:"actionsRequired"`
	Description     string `pulumi:"description"`
	Status          string `pulumi:"status"`
}





type VaultPrivateLinkServiceConnectionStateResponseInput interface {
	pulumi.Input

	ToVaultPrivateLinkServiceConnectionStateResponseOutput() VaultPrivateLinkServiceConnectionStateResponseOutput
	ToVaultPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Context) VaultPrivateLinkServiceConnectionStateResponseOutput
}

type VaultPrivateLinkServiceConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringInput `pulumi:"actionsRequired"`
	Description     pulumi.StringInput `pulumi:"description"`
	Status          pulumi.StringInput `pulumi:"status"`
}

func (VaultPrivateLinkServiceConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (i VaultPrivateLinkServiceConnectionStateResponseArgs) ToVaultPrivateLinkServiceConnectionStateResponseOutput() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return i.ToVaultPrivateLinkServiceConnectionStateResponseOutputWithContext(context.Background())
}

func (i VaultPrivateLinkServiceConnectionStateResponseArgs) ToVaultPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) VaultPrivateLinkServiceConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPrivateLinkServiceConnectionStateResponseOutput)
}

type VaultPrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (VaultPrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ToVaultPrivateLinkServiceConnectionStateResponseOutput() VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ToVaultPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) VaultPrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o VaultPrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPrivateLinkServiceConnectionStateResponse) string { return v.Status }).(pulumi.StringOutput)
}

type VaultPropertiesResponse struct {
	PrivateEndpointConnections          []PrivateEndpointConnectionVaultPropertiesResponse `pulumi:"privateEndpointConnections"`
	PrivateEndpointStateForBackup       string                                             `pulumi:"privateEndpointStateForBackup"`
	PrivateEndpointStateForSiteRecovery string                                             `pulumi:"privateEndpointStateForSiteRecovery"`
	ProvisioningState                   string                                             `pulumi:"provisioningState"`
	UpgradeDetails                      *UpgradeDetailsResponse                            `pulumi:"upgradeDetails"`
}





type VaultPropertiesResponseInput interface {
	pulumi.Input

	ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput
	ToVaultPropertiesResponseOutputWithContext(context.Context) VaultPropertiesResponseOutput
}

type VaultPropertiesResponseArgs struct {
	PrivateEndpointConnections          PrivateEndpointConnectionVaultPropertiesResponseArrayInput `pulumi:"privateEndpointConnections"`
	PrivateEndpointStateForBackup       pulumi.StringInput                                         `pulumi:"privateEndpointStateForBackup"`
	PrivateEndpointStateForSiteRecovery pulumi.StringInput                                         `pulumi:"privateEndpointStateForSiteRecovery"`
	ProvisioningState                   pulumi.StringInput                                         `pulumi:"provisioningState"`
	UpgradeDetails                      UpgradeDetailsResponsePtrInput                             `pulumi:"upgradeDetails"`
}

func (VaultPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return i.ToVaultPropertiesResponseOutputWithContext(context.Background())
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesResponseOutput)
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return i.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i VaultPropertiesResponseArgs) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesResponseOutput).ToVaultPropertiesResponsePtrOutputWithContext(ctx)
}









type VaultPropertiesResponsePtrInput interface {
	pulumi.Input

	ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput
	ToVaultPropertiesResponsePtrOutputWithContext(context.Context) VaultPropertiesResponsePtrOutput
}

type vaultPropertiesResponsePtrType VaultPropertiesResponseArgs

func VaultPropertiesResponsePtr(v *VaultPropertiesResponseArgs) VaultPropertiesResponsePtrInput {
	return (*vaultPropertiesResponsePtrType)(v)
}

func (*vaultPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponse)(nil)).Elem()
}

func (i *vaultPropertiesResponsePtrType) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return i.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *vaultPropertiesResponsePtrType) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VaultPropertiesResponsePtrOutput)
}

type VaultPropertiesResponseOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutput() VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponseOutputWithContext(ctx context.Context) VaultPropertiesResponseOutput {
	return o
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return o.ToVaultPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o VaultPropertiesResponseOutput) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VaultPropertiesResponse) *VaultPropertiesResponse {
		return &v
	}).(VaultPropertiesResponsePtrOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointConnections() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) []PrivateEndpointConnectionVaultPropertiesResponse {
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointStateForBackup() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.PrivateEndpointStateForBackup }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) PrivateEndpointStateForSiteRecovery() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.PrivateEndpointStateForSiteRecovery }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o VaultPropertiesResponseOutput) UpgradeDetails() UpgradeDetailsResponsePtrOutput {
	return o.ApplyT(func(v VaultPropertiesResponse) *UpgradeDetailsResponse { return v.UpgradeDetails }).(UpgradeDetailsResponsePtrOutput)
}

type VaultPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (VaultPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VaultPropertiesResponse)(nil)).Elem()
}

func (o VaultPropertiesResponsePtrOutput) ToVaultPropertiesResponsePtrOutput() VaultPropertiesResponsePtrOutput {
	return o
}

func (o VaultPropertiesResponsePtrOutput) ToVaultPropertiesResponsePtrOutputWithContext(ctx context.Context) VaultPropertiesResponsePtrOutput {
	return o
}

func (o VaultPropertiesResponsePtrOutput) Elem() VaultPropertiesResponseOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) VaultPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret VaultPropertiesResponse
		return ret
	}).(VaultPropertiesResponseOutput)
}

func (o VaultPropertiesResponsePtrOutput) PrivateEndpointConnections() PrivateEndpointConnectionVaultPropertiesResponseArrayOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) []PrivateEndpointConnectionVaultPropertiesResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpointConnections
	}).(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput)
}

func (o VaultPropertiesResponsePtrOutput) PrivateEndpointStateForBackup() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrivateEndpointStateForBackup
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) PrivateEndpointStateForSiteRecovery() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrivateEndpointStateForSiteRecovery
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o VaultPropertiesResponsePtrOutput) UpgradeDetails() UpgradeDetailsResponsePtrOutput {
	return o.ApplyT(func(v *VaultPropertiesResponse) *UpgradeDetailsResponse {
		if v == nil {
			return nil
		}
		return v.UpgradeDetails
	}).(UpgradeDetailsResponsePtrOutput)
}

type WeeklyRetentionFormat struct {
	DaysOfTheWeek   []DayOfWeek   `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth []WeekOfMonth `pulumi:"weeksOfTheMonth"`
}





type WeeklyRetentionFormatInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatOutput() WeeklyRetentionFormatOutput
	ToWeeklyRetentionFormatOutputWithContext(context.Context) WeeklyRetentionFormatOutput
}

type WeeklyRetentionFormatArgs struct {
	DaysOfTheWeek   DayOfWeekArrayInput   `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth WeekOfMonthArrayInput `pulumi:"weeksOfTheMonth"`
}

func (WeeklyRetentionFormatArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormat)(nil)).Elem()
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatOutput() WeeklyRetentionFormatOutput {
	return i.ToWeeklyRetentionFormatOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatOutputWithContext(ctx context.Context) WeeklyRetentionFormatOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatOutput)
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return i.ToWeeklyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatArgs) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatOutput).ToWeeklyRetentionFormatPtrOutputWithContext(ctx)
}









type WeeklyRetentionFormatPtrInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput
	ToWeeklyRetentionFormatPtrOutputWithContext(context.Context) WeeklyRetentionFormatPtrOutput
}

type weeklyRetentionFormatPtrType WeeklyRetentionFormatArgs

func WeeklyRetentionFormatPtr(v *WeeklyRetentionFormatArgs) WeeklyRetentionFormatPtrInput {
	return (*weeklyRetentionFormatPtrType)(v)
}

func (*weeklyRetentionFormatPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormat)(nil)).Elem()
}

func (i *weeklyRetentionFormatPtrType) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return i.ToWeeklyRetentionFormatPtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionFormatPtrType) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatPtrOutput)
}

type WeeklyRetentionFormatOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormat)(nil)).Elem()
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatOutput() WeeklyRetentionFormatOutput {
	return o
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatOutputWithContext(ctx context.Context) WeeklyRetentionFormatOutput {
	return o
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return o.ToWeeklyRetentionFormatPtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionFormatOutput) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionFormat) *WeeklyRetentionFormat {
		return &v
	}).(WeeklyRetentionFormatPtrOutput)
}

func (o WeeklyRetentionFormatOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormat) []DayOfWeek { return v.DaysOfTheWeek }).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionFormatOutput) WeeksOfTheMonth() WeekOfMonthArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormat) []WeekOfMonth { return v.WeeksOfTheMonth }).(WeekOfMonthArrayOutput)
}

type WeeklyRetentionFormatPtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormat)(nil)).Elem()
}

func (o WeeklyRetentionFormatPtrOutput) ToWeeklyRetentionFormatPtrOutput() WeeklyRetentionFormatPtrOutput {
	return o
}

func (o WeeklyRetentionFormatPtrOutput) ToWeeklyRetentionFormatPtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatPtrOutput {
	return o
}

func (o WeeklyRetentionFormatPtrOutput) Elem() WeeklyRetentionFormatOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormat) WeeklyRetentionFormat {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionFormat
		return ret
	}).(WeeklyRetentionFormatOutput)
}

func (o WeeklyRetentionFormatPtrOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormat) []DayOfWeek {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionFormatPtrOutput) WeeksOfTheMonth() WeekOfMonthArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormat) []WeekOfMonth {
		if v == nil {
			return nil
		}
		return v.WeeksOfTheMonth
	}).(WeekOfMonthArrayOutput)
}

type WeeklyRetentionFormatResponse struct {
	DaysOfTheWeek   []string `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth []string `pulumi:"weeksOfTheMonth"`
}





type WeeklyRetentionFormatResponseInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatResponseOutput() WeeklyRetentionFormatResponseOutput
	ToWeeklyRetentionFormatResponseOutputWithContext(context.Context) WeeklyRetentionFormatResponseOutput
}

type WeeklyRetentionFormatResponseArgs struct {
	DaysOfTheWeek   pulumi.StringArrayInput `pulumi:"daysOfTheWeek"`
	WeeksOfTheMonth pulumi.StringArrayInput `pulumi:"weeksOfTheMonth"`
}

func (WeeklyRetentionFormatResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponseOutput() WeeklyRetentionFormatResponseOutput {
	return i.ToWeeklyRetentionFormatResponseOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponseOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatResponseOutput)
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return i.ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionFormatResponseArgs) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatResponseOutput).ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx)
}









type WeeklyRetentionFormatResponsePtrInput interface {
	pulumi.Input

	ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput
	ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Context) WeeklyRetentionFormatResponsePtrOutput
}

type weeklyRetentionFormatResponsePtrType WeeklyRetentionFormatResponseArgs

func WeeklyRetentionFormatResponsePtr(v *WeeklyRetentionFormatResponseArgs) WeeklyRetentionFormatResponsePtrInput {
	return (*weeklyRetentionFormatResponsePtrType)(v)
}

func (*weeklyRetentionFormatResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (i *weeklyRetentionFormatResponsePtrType) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return i.ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionFormatResponsePtrType) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionFormatResponsePtrOutput)
}

type WeeklyRetentionFormatResponseOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponseOutput() WeeklyRetentionFormatResponseOutput {
	return o
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponseOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponseOutput {
	return o
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return o.ToWeeklyRetentionFormatResponsePtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionFormatResponseOutput) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionFormatResponse) *WeeklyRetentionFormatResponse {
		return &v
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o WeeklyRetentionFormatResponseOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormatResponse) []string { return v.DaysOfTheWeek }).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionFormatResponseOutput) WeeksOfTheMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionFormatResponse) []string { return v.WeeksOfTheMonth }).(pulumi.StringArrayOutput)
}

type WeeklyRetentionFormatResponsePtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionFormatResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionFormatResponse)(nil)).Elem()
}

func (o WeeklyRetentionFormatResponsePtrOutput) ToWeeklyRetentionFormatResponsePtrOutput() WeeklyRetentionFormatResponsePtrOutput {
	return o
}

func (o WeeklyRetentionFormatResponsePtrOutput) ToWeeklyRetentionFormatResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionFormatResponsePtrOutput {
	return o
}

func (o WeeklyRetentionFormatResponsePtrOutput) Elem() WeeklyRetentionFormatResponseOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormatResponse) WeeklyRetentionFormatResponse {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionFormatResponse
		return ret
	}).(WeeklyRetentionFormatResponseOutput)
}

func (o WeeklyRetentionFormatResponsePtrOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormatResponse) []string {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionFormatResponsePtrOutput) WeeksOfTheMonth() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionFormatResponse) []string {
		if v == nil {
			return nil
		}
		return v.WeeksOfTheMonth
	}).(pulumi.StringArrayOutput)
}

type WeeklyRetentionSchedule struct {
	DaysOfTheWeek     []DayOfWeek        `pulumi:"daysOfTheWeek"`
	RetentionDuration *RetentionDuration `pulumi:"retentionDuration"`
	RetentionTimes    []string           `pulumi:"retentionTimes"`
}





type WeeklyRetentionScheduleInput interface {
	pulumi.Input

	ToWeeklyRetentionScheduleOutput() WeeklyRetentionScheduleOutput
	ToWeeklyRetentionScheduleOutputWithContext(context.Context) WeeklyRetentionScheduleOutput
}

type WeeklyRetentionScheduleArgs struct {
	DaysOfTheWeek     DayOfWeekArrayInput       `pulumi:"daysOfTheWeek"`
	RetentionDuration RetentionDurationPtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput   `pulumi:"retentionTimes"`
}

func (WeeklyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionSchedule)(nil)).Elem()
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionScheduleOutput() WeeklyRetentionScheduleOutput {
	return i.ToWeeklyRetentionScheduleOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionScheduleOutputWithContext(ctx context.Context) WeeklyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleOutput)
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return i.ToWeeklyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleArgs) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleOutput).ToWeeklyRetentionSchedulePtrOutputWithContext(ctx)
}









type WeeklyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput
	ToWeeklyRetentionSchedulePtrOutputWithContext(context.Context) WeeklyRetentionSchedulePtrOutput
}

type weeklyRetentionSchedulePtrType WeeklyRetentionScheduleArgs

func WeeklyRetentionSchedulePtr(v *WeeklyRetentionScheduleArgs) WeeklyRetentionSchedulePtrInput {
	return (*weeklyRetentionSchedulePtrType)(v)
}

func (*weeklyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionSchedule)(nil)).Elem()
}

func (i *weeklyRetentionSchedulePtrType) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return i.ToWeeklyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionSchedulePtrType) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionSchedulePtrOutput)
}

type WeeklyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionSchedule)(nil)).Elem()
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionScheduleOutput() WeeklyRetentionScheduleOutput {
	return o
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionScheduleOutputWithContext(ctx context.Context) WeeklyRetentionScheduleOutput {
	return o
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return o.ToWeeklyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionScheduleOutput) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionSchedule) *WeeklyRetentionSchedule {
		return &v
	}).(WeeklyRetentionSchedulePtrOutput)
}

func (o WeeklyRetentionScheduleOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionSchedule) []DayOfWeek { return v.DaysOfTheWeek }).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v WeeklyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o WeeklyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type WeeklyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionSchedule)(nil)).Elem()
}

func (o WeeklyRetentionSchedulePtrOutput) ToWeeklyRetentionSchedulePtrOutput() WeeklyRetentionSchedulePtrOutput {
	return o
}

func (o WeeklyRetentionSchedulePtrOutput) ToWeeklyRetentionSchedulePtrOutputWithContext(ctx context.Context) WeeklyRetentionSchedulePtrOutput {
	return o
}

func (o WeeklyRetentionSchedulePtrOutput) Elem() WeeklyRetentionScheduleOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) WeeklyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionSchedule
		return ret
	}).(WeeklyRetentionScheduleOutput)
}

func (o WeeklyRetentionSchedulePtrOutput) DaysOfTheWeek() DayOfWeekArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) []DayOfWeek {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(DayOfWeekArrayOutput)
}

func (o WeeklyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o WeeklyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type WeeklyRetentionScheduleResponse struct {
	DaysOfTheWeek     []string                   `pulumi:"daysOfTheWeek"`
	RetentionDuration *RetentionDurationResponse `pulumi:"retentionDuration"`
	RetentionTimes    []string                   `pulumi:"retentionTimes"`
}





type WeeklyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToWeeklyRetentionScheduleResponseOutput() WeeklyRetentionScheduleResponseOutput
	ToWeeklyRetentionScheduleResponseOutputWithContext(context.Context) WeeklyRetentionScheduleResponseOutput
}

type WeeklyRetentionScheduleResponseArgs struct {
	DaysOfTheWeek     pulumi.StringArrayInput           `pulumi:"daysOfTheWeek"`
	RetentionDuration RetentionDurationResponsePtrInput `pulumi:"retentionDuration"`
	RetentionTimes    pulumi.StringArrayInput           `pulumi:"retentionTimes"`
}

func (WeeklyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponseOutput() WeeklyRetentionScheduleResponseOutput {
	return i.ToWeeklyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponseOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleResponseOutput)
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return i.ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i WeeklyRetentionScheduleResponseArgs) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleResponseOutput).ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type WeeklyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput
	ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Context) WeeklyRetentionScheduleResponsePtrOutput
}

type weeklyRetentionScheduleResponsePtrType WeeklyRetentionScheduleResponseArgs

func WeeklyRetentionScheduleResponsePtr(v *WeeklyRetentionScheduleResponseArgs) WeeklyRetentionScheduleResponsePtrInput {
	return (*weeklyRetentionScheduleResponsePtrType)(v)
}

func (*weeklyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (i *weeklyRetentionScheduleResponsePtrType) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return i.ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *weeklyRetentionScheduleResponsePtrType) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WeeklyRetentionScheduleResponsePtrOutput)
}

type WeeklyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponseOutput() WeeklyRetentionScheduleResponseOutput {
	return o
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponseOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponseOutput {
	return o
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return o.ToWeeklyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o WeeklyRetentionScheduleResponseOutput) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WeeklyRetentionScheduleResponse) *WeeklyRetentionScheduleResponse {
		return &v
	}).(WeeklyRetentionScheduleResponsePtrOutput)
}

func (o WeeklyRetentionScheduleResponseOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionScheduleResponse) []string { return v.DaysOfTheWeek }).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v WeeklyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o WeeklyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WeeklyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type WeeklyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (WeeklyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WeeklyRetentionScheduleResponse)(nil)).Elem()
}

func (o WeeklyRetentionScheduleResponsePtrOutput) ToWeeklyRetentionScheduleResponsePtrOutput() WeeklyRetentionScheduleResponsePtrOutput {
	return o
}

func (o WeeklyRetentionScheduleResponsePtrOutput) ToWeeklyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) WeeklyRetentionScheduleResponsePtrOutput {
	return o
}

func (o WeeklyRetentionScheduleResponsePtrOutput) Elem() WeeklyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) WeeklyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret WeeklyRetentionScheduleResponse
		return ret
	}).(WeeklyRetentionScheduleResponseOutput)
}

func (o WeeklyRetentionScheduleResponsePtrOutput) DaysOfTheWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.DaysOfTheWeek
	}).(pulumi.StringArrayOutput)
}

func (o WeeklyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o WeeklyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WeeklyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type YearlyRetentionSchedule struct {
	MonthsOfYear                []MonthOfYear            `pulumi:"monthsOfYear"`
	RetentionDuration           *RetentionDuration       `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormat    `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *RetentionScheduleFormat `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormat   `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                 `pulumi:"retentionTimes"`
}





type YearlyRetentionScheduleInput interface {
	pulumi.Input

	ToYearlyRetentionScheduleOutput() YearlyRetentionScheduleOutput
	ToYearlyRetentionScheduleOutputWithContext(context.Context) YearlyRetentionScheduleOutput
}

type YearlyRetentionScheduleArgs struct {
	MonthsOfYear                MonthOfYearArrayInput           `pulumi:"monthsOfYear"`
	RetentionDuration           RetentionDurationPtrInput       `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatPtrInput    `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType RetentionScheduleFormatPtrInput `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatPtrInput   `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput         `pulumi:"retentionTimes"`
}

func (YearlyRetentionScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionSchedule)(nil)).Elem()
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionScheduleOutput() YearlyRetentionScheduleOutput {
	return i.ToYearlyRetentionScheduleOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionScheduleOutputWithContext(ctx context.Context) YearlyRetentionScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleOutput)
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return i.ToYearlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleArgs) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleOutput).ToYearlyRetentionSchedulePtrOutputWithContext(ctx)
}









type YearlyRetentionSchedulePtrInput interface {
	pulumi.Input

	ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput
	ToYearlyRetentionSchedulePtrOutputWithContext(context.Context) YearlyRetentionSchedulePtrOutput
}

type yearlyRetentionSchedulePtrType YearlyRetentionScheduleArgs

func YearlyRetentionSchedulePtr(v *YearlyRetentionScheduleArgs) YearlyRetentionSchedulePtrInput {
	return (*yearlyRetentionSchedulePtrType)(v)
}

func (*yearlyRetentionSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionSchedule)(nil)).Elem()
}

func (i *yearlyRetentionSchedulePtrType) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return i.ToYearlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (i *yearlyRetentionSchedulePtrType) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionSchedulePtrOutput)
}

type YearlyRetentionScheduleOutput struct{ *pulumi.OutputState }

func (YearlyRetentionScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionSchedule)(nil)).Elem()
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionScheduleOutput() YearlyRetentionScheduleOutput {
	return o
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionScheduleOutputWithContext(ctx context.Context) YearlyRetentionScheduleOutput {
	return o
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return o.ToYearlyRetentionSchedulePtrOutputWithContext(context.Background())
}

func (o YearlyRetentionScheduleOutput) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v YearlyRetentionSchedule) *YearlyRetentionSchedule {
		return &v
	}).(YearlyRetentionSchedulePtrOutput)
}

func (o YearlyRetentionScheduleOutput) MonthsOfYear() MonthOfYearArrayOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) []MonthOfYear { return v.MonthsOfYear }).(MonthOfYearArrayOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *RetentionDuration { return v.RetentionDuration }).(RetentionDurationPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *DailyRetentionFormat { return v.RetentionScheduleDaily }).(DailyRetentionFormatPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionScheduleFormatType() RetentionScheduleFormatPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *RetentionScheduleFormat { return v.RetentionScheduleFormatType }).(RetentionScheduleFormatPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) *WeeklyRetentionFormat { return v.RetentionScheduleWeekly }).(WeeklyRetentionFormatPtrOutput)
}

func (o YearlyRetentionScheduleOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v YearlyRetentionSchedule) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type YearlyRetentionSchedulePtrOutput struct{ *pulumi.OutputState }

func (YearlyRetentionSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionSchedule)(nil)).Elem()
}

func (o YearlyRetentionSchedulePtrOutput) ToYearlyRetentionSchedulePtrOutput() YearlyRetentionSchedulePtrOutput {
	return o
}

func (o YearlyRetentionSchedulePtrOutput) ToYearlyRetentionSchedulePtrOutputWithContext(ctx context.Context) YearlyRetentionSchedulePtrOutput {
	return o
}

func (o YearlyRetentionSchedulePtrOutput) Elem() YearlyRetentionScheduleOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) YearlyRetentionSchedule {
		if v != nil {
			return *v
		}
		var ret YearlyRetentionSchedule
		return ret
	}).(YearlyRetentionScheduleOutput)
}

func (o YearlyRetentionSchedulePtrOutput) MonthsOfYear() MonthOfYearArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) []MonthOfYear {
		if v == nil {
			return nil
		}
		return v.MonthsOfYear
	}).(MonthOfYearArrayOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionDuration() RetentionDurationPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *RetentionDuration {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionScheduleDaily() DailyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *DailyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionScheduleFormatType() RetentionScheduleFormatPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *RetentionScheduleFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(RetentionScheduleFormatPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) *WeeklyRetentionFormat {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatPtrOutput)
}

func (o YearlyRetentionSchedulePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionSchedule) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

type YearlyRetentionScheduleResponse struct {
	MonthsOfYear                []string                       `pulumi:"monthsOfYear"`
	RetentionDuration           *RetentionDurationResponse     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      *DailyRetentionFormatResponse  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType *string                        `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     *WeeklyRetentionFormatResponse `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              []string                       `pulumi:"retentionTimes"`
}





type YearlyRetentionScheduleResponseInput interface {
	pulumi.Input

	ToYearlyRetentionScheduleResponseOutput() YearlyRetentionScheduleResponseOutput
	ToYearlyRetentionScheduleResponseOutputWithContext(context.Context) YearlyRetentionScheduleResponseOutput
}

type YearlyRetentionScheduleResponseArgs struct {
	MonthsOfYear                pulumi.StringArrayInput               `pulumi:"monthsOfYear"`
	RetentionDuration           RetentionDurationResponsePtrInput     `pulumi:"retentionDuration"`
	RetentionScheduleDaily      DailyRetentionFormatResponsePtrInput  `pulumi:"retentionScheduleDaily"`
	RetentionScheduleFormatType pulumi.StringPtrInput                 `pulumi:"retentionScheduleFormatType"`
	RetentionScheduleWeekly     WeeklyRetentionFormatResponsePtrInput `pulumi:"retentionScheduleWeekly"`
	RetentionTimes              pulumi.StringArrayInput               `pulumi:"retentionTimes"`
}

func (YearlyRetentionScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponseOutput() YearlyRetentionScheduleResponseOutput {
	return i.ToYearlyRetentionScheduleResponseOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponseOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleResponseOutput)
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return i.ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i YearlyRetentionScheduleResponseArgs) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleResponseOutput).ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx)
}









type YearlyRetentionScheduleResponsePtrInput interface {
	pulumi.Input

	ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput
	ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Context) YearlyRetentionScheduleResponsePtrOutput
}

type yearlyRetentionScheduleResponsePtrType YearlyRetentionScheduleResponseArgs

func YearlyRetentionScheduleResponsePtr(v *YearlyRetentionScheduleResponseArgs) YearlyRetentionScheduleResponsePtrInput {
	return (*yearlyRetentionScheduleResponsePtrType)(v)
}

func (*yearlyRetentionScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (i *yearlyRetentionScheduleResponsePtrType) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return i.ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *yearlyRetentionScheduleResponsePtrType) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(YearlyRetentionScheduleResponsePtrOutput)
}

type YearlyRetentionScheduleResponseOutput struct{ *pulumi.OutputState }

func (YearlyRetentionScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponseOutput() YearlyRetentionScheduleResponseOutput {
	return o
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponseOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponseOutput {
	return o
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return o.ToYearlyRetentionScheduleResponsePtrOutputWithContext(context.Background())
}

func (o YearlyRetentionScheduleResponseOutput) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v YearlyRetentionScheduleResponse) *YearlyRetentionScheduleResponse {
		return &v
	}).(YearlyRetentionScheduleResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) MonthsOfYear() pulumi.StringArrayOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) []string { return v.MonthsOfYear }).(pulumi.StringArrayOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *RetentionDurationResponse { return v.RetentionDuration }).(RetentionDurationResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *DailyRetentionFormatResponse { return v.RetentionScheduleDaily }).(DailyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *string { return v.RetentionScheduleFormatType }).(pulumi.StringPtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponseOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v YearlyRetentionScheduleResponse) []string { return v.RetentionTimes }).(pulumi.StringArrayOutput)
}

type YearlyRetentionScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (YearlyRetentionScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**YearlyRetentionScheduleResponse)(nil)).Elem()
}

func (o YearlyRetentionScheduleResponsePtrOutput) ToYearlyRetentionScheduleResponsePtrOutput() YearlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o YearlyRetentionScheduleResponsePtrOutput) ToYearlyRetentionScheduleResponsePtrOutputWithContext(ctx context.Context) YearlyRetentionScheduleResponsePtrOutput {
	return o
}

func (o YearlyRetentionScheduleResponsePtrOutput) Elem() YearlyRetentionScheduleResponseOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) YearlyRetentionScheduleResponse {
		if v != nil {
			return *v
		}
		var ret YearlyRetentionScheduleResponse
		return ret
	}).(YearlyRetentionScheduleResponseOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) MonthsOfYear() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.MonthsOfYear
	}).(pulumi.StringArrayOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionDuration() RetentionDurationResponsePtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *RetentionDurationResponse {
		if v == nil {
			return nil
		}
		return v.RetentionDuration
	}).(RetentionDurationResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionScheduleDaily() DailyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *DailyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleDaily
	}).(DailyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionScheduleFormatType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleFormatType
	}).(pulumi.StringPtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionScheduleWeekly() WeeklyRetentionFormatResponsePtrOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) *WeeklyRetentionFormatResponse {
		if v == nil {
			return nil
		}
		return v.RetentionScheduleWeekly
	}).(WeeklyRetentionFormatResponsePtrOutput)
}

func (o YearlyRetentionScheduleResponsePtrOutput) RetentionTimes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *YearlyRetentionScheduleResponse) []string {
		if v == nil {
			return nil
		}
		return v.RetentionTimes
	}).(pulumi.StringArrayOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AzureIaaSVMProtectionPolicyInput)(nil)).Elem(), AzureIaaSVMProtectionPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureIaaSVMProtectionPolicyResponseInput)(nil)).Elem(), AzureIaaSVMProtectionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSqlProtectionPolicyInput)(nil)).Elem(), AzureSqlProtectionPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AzureSqlProtectionPolicyResponseInput)(nil)).Elem(), AzureSqlProtectionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionFormatInput)(nil)).Elem(), DailyRetentionFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionFormatPtrInput)(nil)).Elem(), DailyRetentionFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionFormatResponseInput)(nil)).Elem(), DailyRetentionFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionFormatResponsePtrInput)(nil)).Elem(), DailyRetentionFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionScheduleInput)(nil)).Elem(), DailyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionSchedulePtrInput)(nil)).Elem(), DailyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionScheduleResponseInput)(nil)).Elem(), DailyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DailyRetentionScheduleResponsePtrInput)(nil)).Elem(), DailyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DayInput)(nil)).Elem(), DayArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DayArrayInput)(nil)).Elem(), DayArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DayResponseInput)(nil)).Elem(), DayResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DayResponseArrayInput)(nil)).Elem(), DayResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityDataInput)(nil)).Elem(), IdentityDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityDataPtrInput)(nil)).Elem(), IdentityDataArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityDataResponseInput)(nil)).Elem(), IdentityDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*IdentityDataResponsePtrInput)(nil)).Elem(), IdentityDataResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LongTermRetentionPolicyInput)(nil)).Elem(), LongTermRetentionPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LongTermRetentionPolicyResponseInput)(nil)).Elem(), LongTermRetentionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LongTermSchedulePolicyInput)(nil)).Elem(), LongTermSchedulePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LongTermSchedulePolicyResponseInput)(nil)).Elem(), LongTermSchedulePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MabProtectionPolicyInput)(nil)).Elem(), MabProtectionPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MabProtectionPolicyResponseInput)(nil)).Elem(), MabProtectionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonthlyRetentionScheduleInput)(nil)).Elem(), MonthlyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonthlyRetentionSchedulePtrInput)(nil)).Elem(), MonthlyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonthlyRetentionScheduleResponseInput)(nil)).Elem(), MonthlyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MonthlyRetentionScheduleResponsePtrInput)(nil)).Elem(), MonthlyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionVaultPropertiesResponseInput)(nil)).Elem(), PrivateEndpointConnectionVaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointConnectionVaultPropertiesResponseArrayInput)(nil)).Elem(), PrivateEndpointConnectionVaultPropertiesResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*PrivateEndpointResponseInput)(nil)).Elem(), PrivateEndpointResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionDurationInput)(nil)).Elem(), RetentionDurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionDurationPtrInput)(nil)).Elem(), RetentionDurationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionDurationResponseInput)(nil)).Elem(), RetentionDurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*RetentionDurationResponsePtrInput)(nil)).Elem(), RetentionDurationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleRetentionPolicyInput)(nil)).Elem(), SimpleRetentionPolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleRetentionPolicyResponseInput)(nil)).Elem(), SimpleRetentionPolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleSchedulePolicyInput)(nil)).Elem(), SimpleSchedulePolicyArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SimpleSchedulePolicyResponseInput)(nil)).Elem(), SimpleSchedulePolicyResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeDetailsResponseInput)(nil)).Elem(), UpgradeDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*UpgradeDetailsResponsePtrInput)(nil)).Elem(), UpgradeDetailsResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPrivateEndpointConnectionResponseInput)(nil)).Elem(), VaultPrivateEndpointConnectionResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPrivateLinkServiceConnectionStateResponseInput)(nil)).Elem(), VaultPrivateLinkServiceConnectionStateResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPropertiesResponseInput)(nil)).Elem(), VaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*VaultPropertiesResponsePtrInput)(nil)).Elem(), VaultPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionFormatInput)(nil)).Elem(), WeeklyRetentionFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionFormatPtrInput)(nil)).Elem(), WeeklyRetentionFormatArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionFormatResponseInput)(nil)).Elem(), WeeklyRetentionFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionFormatResponsePtrInput)(nil)).Elem(), WeeklyRetentionFormatResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionScheduleInput)(nil)).Elem(), WeeklyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionSchedulePtrInput)(nil)).Elem(), WeeklyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionScheduleResponseInput)(nil)).Elem(), WeeklyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WeeklyRetentionScheduleResponsePtrInput)(nil)).Elem(), WeeklyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*YearlyRetentionScheduleInput)(nil)).Elem(), YearlyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*YearlyRetentionSchedulePtrInput)(nil)).Elem(), YearlyRetentionScheduleArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*YearlyRetentionScheduleResponseInput)(nil)).Elem(), YearlyRetentionScheduleResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*YearlyRetentionScheduleResponsePtrInput)(nil)).Elem(), YearlyRetentionScheduleResponseArgs{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectionPolicyOutput{})
	pulumi.RegisterOutputType(AzureIaaSVMProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectionPolicyOutput{})
	pulumi.RegisterOutputType(AzureSqlProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatPtrOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatResponseOutput{})
	pulumi.RegisterOutputType(DailyRetentionFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(DailyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(DailyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(DailyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(DailyRetentionScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(DayOutput{})
	pulumi.RegisterOutputType(DayArrayOutput{})
	pulumi.RegisterOutputType(DayResponseOutput{})
	pulumi.RegisterOutputType(DayResponseArrayOutput{})
	pulumi.RegisterOutputType(IdentityDataOutput{})
	pulumi.RegisterOutputType(IdentityDataPtrOutput{})
	pulumi.RegisterOutputType(IdentityDataResponseOutput{})
	pulumi.RegisterOutputType(IdentityDataResponsePtrOutput{})
	pulumi.RegisterOutputType(LongTermRetentionPolicyOutput{})
	pulumi.RegisterOutputType(LongTermRetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(LongTermSchedulePolicyOutput{})
	pulumi.RegisterOutputType(LongTermSchedulePolicyResponseOutput{})
	pulumi.RegisterOutputType(MabProtectionPolicyOutput{})
	pulumi.RegisterOutputType(MabProtectionPolicyResponseOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(MonthlyRetentionScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionVaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionVaultPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(RetentionDurationOutput{})
	pulumi.RegisterOutputType(RetentionDurationPtrOutput{})
	pulumi.RegisterOutputType(RetentionDurationResponseOutput{})
	pulumi.RegisterOutputType(RetentionDurationResponsePtrOutput{})
	pulumi.RegisterOutputType(SimpleRetentionPolicyOutput{})
	pulumi.RegisterOutputType(SimpleRetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(SimpleSchedulePolicyOutput{})
	pulumi.RegisterOutputType(SimpleSchedulePolicyResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponseOutput{})
	pulumi.RegisterOutputType(UpgradeDetailsResponsePtrOutput{})
	pulumi.RegisterOutputType(VaultPrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(VaultPrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponseOutput{})
	pulumi.RegisterOutputType(VaultPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatPtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatResponseOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionFormatResponsePtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(WeeklyRetentionScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(YearlyRetentionScheduleOutput{})
	pulumi.RegisterOutputType(YearlyRetentionSchedulePtrOutput{})
	pulumi.RegisterOutputType(YearlyRetentionScheduleResponseOutput{})
	pulumi.RegisterOutputType(YearlyRetentionScheduleResponsePtrOutput{})
}
