


package v20180915

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GlobalSchedule struct {
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


func NewGlobalSchedule(ctx *pulumi.Context,
	name string, args *GlobalScheduleArgs, opts ...pulumi.ResourceOption) (*GlobalSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Status == nil {
		args.Status = pulumi.StringPtr("Disabled")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:GlobalSchedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20160515:GlobalSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource GlobalSchedule
	err := ctx.RegisterResource("azure-native:devtestlab/v20180915:GlobalSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGlobalSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GlobalScheduleState, opts ...pulumi.ResourceOption) (*GlobalSchedule, error) {
	var resource GlobalSchedule
	err := ctx.ReadResource("azure-native:devtestlab/v20180915:GlobalSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type globalScheduleState struct {
}

type GlobalScheduleState struct {
}

func (GlobalScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*globalScheduleState)(nil)).Elem()
}

type globalScheduleArgs struct {
	DailyRecurrence      *DayDetails           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetails          `pulumi:"hourlyRecurrence"`
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


type GlobalScheduleArgs struct {
	DailyRecurrence      DayDetailsPtrInput
	HourlyRecurrence     HourDetailsPtrInput
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

func (GlobalScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*globalScheduleArgs)(nil)).Elem()
}

type GlobalScheduleInput interface {
	pulumi.Input

	ToGlobalScheduleOutput() GlobalScheduleOutput
	ToGlobalScheduleOutputWithContext(ctx context.Context) GlobalScheduleOutput
}

func (*GlobalSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalSchedule)(nil))
}

func (i *GlobalSchedule) ToGlobalScheduleOutput() GlobalScheduleOutput {
	return i.ToGlobalScheduleOutputWithContext(context.Background())
}

func (i *GlobalSchedule) ToGlobalScheduleOutputWithContext(ctx context.Context) GlobalScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GlobalScheduleOutput)
}

type GlobalScheduleOutput struct{ *pulumi.OutputState }

func (GlobalScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GlobalSchedule)(nil))
}

func (o GlobalScheduleOutput) ToGlobalScheduleOutput() GlobalScheduleOutput {
	return o
}

func (o GlobalScheduleOutput) ToGlobalScheduleOutputWithContext(ctx context.Context) GlobalScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(GlobalScheduleOutput{})
}
