


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceFabricSchedule struct {
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


func NewServiceFabricSchedule(ctx *pulumi.Context,
	name string, args *ServiceFabricScheduleArgs, opts ...pulumi.ResourceOption) (*ServiceFabricSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServiceFabricName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceFabricName'")
	}
	if args.UserName == nil {
		return nil, errors.New("invalid value for required argument 'UserName'")
	}
	if args.NotificationSettings != nil {
		args.NotificationSettings = args.NotificationSettings.ToNotificationSettingsPtrOutput().ApplyT(func(v *NotificationSettings) *NotificationSettings { return v.Defaults() }).(NotificationSettingsPtrOutput)
	}
	if isZero(args.Status) {
		args.Status = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:ServiceFabricSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceFabricSchedule
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:ServiceFabricSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceFabricSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceFabricScheduleState, opts ...pulumi.ResourceOption) (*ServiceFabricSchedule, error) {
	var resource ServiceFabricSchedule
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:ServiceFabricSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceFabricScheduleState struct {
}

type ServiceFabricScheduleState struct {
}

func (ServiceFabricScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricScheduleState)(nil)).Elem()
}

type serviceFabricScheduleArgs struct {
	DailyRecurrence      *DayDetails           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetails          `pulumi:"hourlyRecurrence"`
	LabName              string                `pulumi:"labName"`
	Location             *string               `pulumi:"location"`
	Name                 *string               `pulumi:"name"`
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	ServiceFabricName    string                `pulumi:"serviceFabricName"`
	Status               *string               `pulumi:"status"`
	Tags                 map[string]string     `pulumi:"tags"`
	TargetResourceId     *string               `pulumi:"targetResourceId"`
	TaskType             *string               `pulumi:"taskType"`
	TimeZoneId           *string               `pulumi:"timeZoneId"`
	UserName             string                `pulumi:"userName"`
	WeeklyRecurrence     *WeekDetails          `pulumi:"weeklyRecurrence"`
}


type ServiceFabricScheduleArgs struct {
	DailyRecurrence      DayDetailsPtrInput
	HourlyRecurrence     HourDetailsPtrInput
	LabName              pulumi.StringInput
	Location             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NotificationSettings NotificationSettingsPtrInput
	ResourceGroupName    pulumi.StringInput
	ServiceFabricName    pulumi.StringInput
	Status               pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	TargetResourceId     pulumi.StringPtrInput
	TaskType             pulumi.StringPtrInput
	TimeZoneId           pulumi.StringPtrInput
	UserName             pulumi.StringInput
	WeeklyRecurrence     WeekDetailsPtrInput
}

func (ServiceFabricScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceFabricScheduleArgs)(nil)).Elem()
}

type ServiceFabricScheduleInput interface {
	pulumi.Input

	ToServiceFabricScheduleOutput() ServiceFabricScheduleOutput
	ToServiceFabricScheduleOutputWithContext(ctx context.Context) ServiceFabricScheduleOutput
}

func (*ServiceFabricSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceFabricSchedule)(nil)).Elem()
}

func (i *ServiceFabricSchedule) ToServiceFabricScheduleOutput() ServiceFabricScheduleOutput {
	return i.ToServiceFabricScheduleOutputWithContext(context.Background())
}

func (i *ServiceFabricSchedule) ToServiceFabricScheduleOutputWithContext(ctx context.Context) ServiceFabricScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceFabricScheduleOutput)
}

type ServiceFabricScheduleOutput struct{ *pulumi.OutputState }

func (ServiceFabricScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceFabricSchedule)(nil)).Elem()
}

func (o ServiceFabricScheduleOutput) ToServiceFabricScheduleOutput() ServiceFabricScheduleOutput {
	return o
}

func (o ServiceFabricScheduleOutput) ToServiceFabricScheduleOutputWithContext(ctx context.Context) ServiceFabricScheduleOutput {
	return o
}

func (o ServiceFabricScheduleOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o ServiceFabricScheduleOutput) DailyRecurrence() DayDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) DayDetailsResponsePtrOutput { return v.DailyRecurrence }).(DayDetailsResponsePtrOutput)
}

func (o ServiceFabricScheduleOutput) HourlyRecurrence() HourDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) HourDetailsResponsePtrOutput { return v.HourlyRecurrence }).(HourDetailsResponsePtrOutput)
}

func (o ServiceFabricScheduleOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ServiceFabricScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceFabricScheduleOutput) NotificationSettings() NotificationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) NotificationSettingsResponsePtrOutput { return v.NotificationSettings }).(NotificationSettingsResponsePtrOutput)
}

func (o ServiceFabricScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ServiceFabricScheduleOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ServiceFabricScheduleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o ServiceFabricScheduleOutput) TargetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringPtrOutput { return v.TargetResourceId }).(pulumi.StringPtrOutput)
}

func (o ServiceFabricScheduleOutput) TaskType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringPtrOutput { return v.TaskType }).(pulumi.StringPtrOutput)
}

func (o ServiceFabricScheduleOutput) TimeZoneId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringPtrOutput { return v.TimeZoneId }).(pulumi.StringPtrOutput)
}

func (o ServiceFabricScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o ServiceFabricScheduleOutput) UniqueIdentifier() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) pulumi.StringOutput { return v.UniqueIdentifier }).(pulumi.StringOutput)
}

func (o ServiceFabricScheduleOutput) WeeklyRecurrence() WeekDetailsResponsePtrOutput {
	return o.ApplyT(func(v *ServiceFabricSchedule) WeekDetailsResponsePtrOutput { return v.WeeklyRecurrence }).(WeekDetailsResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceFabricScheduleOutput{})
}
