


package devtestlab

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.CustomResourceState

	CreatedDate          pulumi.StringOutput                   `pulumi:"createdDate"`
	DailyRecurrence      DayDetailsResponsePtrOutput           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     HourDetailsResponsePtrOutput          `pulumi:"hourlyRecurrence"`
	Location             pulumi.StringPtrOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	ProvisioningState    pulumi.StringOutput                   `pulumi:"provisioningState"`
	Status               pulumi.StringPtrOutput                `pulumi:"status"`
	Tags                 pulumi.StringMapOutput                `pulumi:"tags"`
	TargetResourceId     pulumi.StringPtrOutput                `pulumi:"targetResourceId"`
	TaskType             pulumi.StringPtrOutput                `pulumi:"taskType"`
	TimeZoneId           pulumi.StringPtrOutput                `pulumi:"timeZoneId"`
	Type                 pulumi.StringOutput                   `pulumi:"type"`
	UniqueIdentifier     pulumi.StringOutput                   `pulumi:"uniqueIdentifier"`
	WeeklyRecurrence     WeekDetailsResponsePtrOutput          `pulumi:"weeklyRecurrence"`
}


func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.NotificationSettings != nil {
		args.NotificationSettings = args.NotificationSettings.ToNotificationSettingsPtrOutput().ApplyT(func(v *NotificationSettings) *NotificationSettings { return v.Defaults() }).(NotificationSettingsPtrOutput)
	}
	if isZero(args.Status) {
		args.Status = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab/v20150521preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:Schedule"),
		},
	})
	opts = append(opts, aliases)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:devtestlab:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure-native:devtestlab:Schedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type scheduleState struct {
}

type ScheduleState struct {
}

func (ScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleState)(nil)).Elem()
}

type scheduleArgs struct {
	DailyRecurrence      *DayDetails           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetails          `pulumi:"hourlyRecurrence"`
	LabName              string                `pulumi:"labName"`
	Location             *string               `pulumi:"location"`
	Name                 *string               `pulumi:"name"`
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	Status               *string               `pulumi:"status"`
	Tags                 map[string]string     `pulumi:"tags"`
	TargetResourceId     *string               `pulumi:"targetResourceId"`
	TaskType             *string               `pulumi:"taskType"`
	TimeZoneId           *string               `pulumi:"timeZoneId"`
	WeeklyRecurrence     *WeekDetails          `pulumi:"weeklyRecurrence"`
}


type ScheduleArgs struct {
	DailyRecurrence      DayDetailsPtrInput
	HourlyRecurrence     HourDetailsPtrInput
	LabName              pulumi.StringInput
	Location             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NotificationSettings NotificationSettingsPtrInput
	ResourceGroupName    pulumi.StringInput
	Status               pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	TargetResourceId     pulumi.StringPtrInput
	TaskType             pulumi.StringPtrInput
	TimeZoneId           pulumi.StringPtrInput
	WeeklyRecurrence     WeekDetailsPtrInput
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*scheduleArgs)(nil)).Elem()
}

type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput
}

func (*Schedule) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i *Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o ScheduleOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) DayDetailsResponsePtrOutput { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o ScheduleOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) HourDetailsResponsePtrOutput { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o ScheduleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduleOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o ScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ScheduleOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ScheduleOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o ScheduleOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) WeekDetailsResponsePtrOutput { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
