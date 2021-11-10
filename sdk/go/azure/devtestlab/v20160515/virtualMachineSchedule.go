


package v20160515

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type VirtualMachineSchedule struct {
	pulumi.CustomResourceState

	CreatedDate          pulumi.StringOutput                   `pulumi:"createdDate"`
	DailyRecurrence      DayDetailsResponsePtrOutput           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     HourDetailsResponsePtrOutput          `pulumi:"hourlyRecurrence"`
	Location             pulumi.StringPtrOutput                `pulumi:"location"`
	Name                 pulumi.StringOutput                   `pulumi:"name"`
	NotificationSettings NotificationSettingsResponsePtrOutput `pulumi:"notificationSettings"`
	ProvisioningState    pulumi.StringPtrOutput                `pulumi:"provisioningState"`
	Status               pulumi.StringPtrOutput                `pulumi:"status"`
	Tags                 pulumi.StringMapOutput                `pulumi:"tags"`
	TargetResourceId     pulumi.StringPtrOutput                `pulumi:"targetResourceId"`
	TaskType             pulumi.StringPtrOutput                `pulumi:"taskType"`
	TimeZoneId           pulumi.StringPtrOutput                `pulumi:"timeZoneId"`
	Type                 pulumi.StringOutput                   `pulumi:"type"`
	UniqueIdentifier     pulumi.StringPtrOutput                `pulumi:"uniqueIdentifier"`
	WeeklyRecurrence     WeekDetailsResponsePtrOutput          `pulumi:"weeklyRecurrence"`
}


func NewVirtualMachineSchedule(ctx *pulumi.Context,
	name string, args *VirtualMachineScheduleArgs, opts ...pulumi.ResourceOption) (*VirtualMachineSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabName == nil {
		return nil, errors.New("invalid value for required argument 'LabName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.VirtualMachineName == nil {
		return nil, errors.New("invalid value for required argument 'VirtualMachineName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devtestlab:VirtualMachineSchedule"),
		},
		{
			Type: pulumi.String("azure-native:devtestlab/v20180915:VirtualMachineSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource VirtualMachineSchedule
	err := ctx.RegisterResource("azure-native:devtestlab/v20160515:VirtualMachineSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetVirtualMachineSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *VirtualMachineScheduleState, opts ...pulumi.ResourceOption) (*VirtualMachineSchedule, error) {
	var resource VirtualMachineSchedule
	err := ctx.ReadResource("azure-native:devtestlab/v20160515:VirtualMachineSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type virtualMachineScheduleState struct {
}

type VirtualMachineScheduleState struct {
}

func (VirtualMachineScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScheduleState)(nil)).Elem()
}

type virtualMachineScheduleArgs struct {
	DailyRecurrence      *DayDetails           `pulumi:"dailyRecurrence"`
	HourlyRecurrence     *HourDetails          `pulumi:"hourlyRecurrence"`
	LabName              string                `pulumi:"labName"`
	Location             *string               `pulumi:"location"`
	Name                 *string               `pulumi:"name"`
	NotificationSettings *NotificationSettings `pulumi:"notificationSettings"`
	ProvisioningState    *string               `pulumi:"provisioningState"`
	ResourceGroupName    string                `pulumi:"resourceGroupName"`
	Status               *string               `pulumi:"status"`
	Tags                 map[string]string     `pulumi:"tags"`
	TargetResourceId     *string               `pulumi:"targetResourceId"`
	TaskType             *string               `pulumi:"taskType"`
	TimeZoneId           *string               `pulumi:"timeZoneId"`
	UniqueIdentifier     *string               `pulumi:"uniqueIdentifier"`
	VirtualMachineName   string                `pulumi:"virtualMachineName"`
	WeeklyRecurrence     *WeekDetails          `pulumi:"weeklyRecurrence"`
}


type VirtualMachineScheduleArgs struct {
	DailyRecurrence      DayDetailsPtrInput
	HourlyRecurrence     HourDetailsPtrInput
	LabName              pulumi.StringInput
	Location             pulumi.StringPtrInput
	Name                 pulumi.StringPtrInput
	NotificationSettings NotificationSettingsPtrInput
	ProvisioningState    pulumi.StringPtrInput
	ResourceGroupName    pulumi.StringInput
	Status               pulumi.StringPtrInput
	Tags                 pulumi.StringMapInput
	TargetResourceId     pulumi.StringPtrInput
	TaskType             pulumi.StringPtrInput
	TimeZoneId           pulumi.StringPtrInput
	UniqueIdentifier     pulumi.StringPtrInput
	VirtualMachineName   pulumi.StringInput
	WeeklyRecurrence     WeekDetailsPtrInput
}

func (VirtualMachineScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*virtualMachineScheduleArgs)(nil)).Elem()
}

type VirtualMachineScheduleInput interface {
	pulumi.Input

	ToVirtualMachineScheduleOutput() VirtualMachineScheduleOutput
	ToVirtualMachineScheduleOutputWithContext(ctx context.Context) VirtualMachineScheduleOutput
}

func (*VirtualMachineSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSchedule)(nil))
}

func (i *VirtualMachineSchedule) ToVirtualMachineScheduleOutput() VirtualMachineScheduleOutput {
	return i.ToVirtualMachineScheduleOutputWithContext(context.Background())
}

func (i *VirtualMachineSchedule) ToVirtualMachineScheduleOutputWithContext(ctx context.Context) VirtualMachineScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualMachineScheduleOutput)
}

type VirtualMachineScheduleOutput struct{ *pulumi.OutputState }

func (VirtualMachineScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualMachineSchedule)(nil))
}

func (o VirtualMachineScheduleOutput) ToVirtualMachineScheduleOutput() VirtualMachineScheduleOutput {
	return o
}

func (o VirtualMachineScheduleOutput) ToVirtualMachineScheduleOutputWithContext(ctx context.Context) VirtualMachineScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(VirtualMachineScheduleOutput{})
}
