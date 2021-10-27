


package devtestlab

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
	if args.Status == nil {
		args.Status = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:devtestlab:ServiceFabricSchedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:ServiceFabricSchedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:devtestlab/v20180915:ServiceFabricSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceFabricSchedule
	err := ctx.RegisterResource("azure-native:devtestlab:ServiceFabricSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceFabricSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceFabricScheduleState, opts ...pulumi.ResourceOption) (*ServiceFabricSchedule, error) {
	var resource ServiceFabricSchedule
	err := ctx.ReadResource("azure-native:devtestlab:ServiceFabricSchedule", name, id, state, &resource, opts...)
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
	return reflect.TypeOf((*ServiceFabricSchedule)(nil))
}

func (i *ServiceFabricSchedule) ToServiceFabricScheduleOutput() ServiceFabricScheduleOutput {
	return i.ToServiceFabricScheduleOutputWithContext(context.Background())
}

func (i *ServiceFabricSchedule) ToServiceFabricScheduleOutputWithContext(ctx context.Context) ServiceFabricScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceFabricScheduleOutput)
}

type ServiceFabricScheduleOutput struct{ *pulumi.OutputState }

func (ServiceFabricScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceFabricSchedule)(nil))
}

func (o ServiceFabricScheduleOutput) ToServiceFabricScheduleOutput() ServiceFabricScheduleOutput {
	return o
}

func (o ServiceFabricScheduleOutput) ToServiceFabricScheduleOutputWithContext(ctx context.Context) ServiceFabricScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceFabricScheduleInput)(nil)).Elem(), &ServiceFabricSchedule{})
	pulumi.RegisterOutputType(ServiceFabricScheduleOutput{})
}
