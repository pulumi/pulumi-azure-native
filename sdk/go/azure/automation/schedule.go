


package automation

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.CustomResourceState

	AdvancedSchedule        AdvancedScheduleResponsePtrOutput `pulumi:"advancedSchedule"`
	CreationTime            pulumi.StringPtrOutput            `pulumi:"creationTime"`
	Description             pulumi.StringPtrOutput            `pulumi:"description"`
	ExpiryTime              pulumi.StringPtrOutput            `pulumi:"expiryTime"`
	ExpiryTimeOffsetMinutes pulumi.Float64PtrOutput           `pulumi:"expiryTimeOffsetMinutes"`
	Frequency               pulumi.StringPtrOutput            `pulumi:"frequency"`
	Interval                pulumi.AnyOutput                  `pulumi:"interval"`
	IsEnabled               pulumi.BoolPtrOutput              `pulumi:"isEnabled"`
	LastModifiedTime        pulumi.StringPtrOutput            `pulumi:"lastModifiedTime"`
	Name                    pulumi.StringOutput               `pulumi:"name"`
	NextRun                 pulumi.StringPtrOutput            `pulumi:"nextRun"`
	NextRunOffsetMinutes    pulumi.Float64PtrOutput           `pulumi:"nextRunOffsetMinutes"`
	StartTime               pulumi.StringPtrOutput            `pulumi:"startTime"`
	StartTimeOffsetMinutes  pulumi.Float64Output              `pulumi:"startTimeOffsetMinutes"`
	TimeZone                pulumi.StringPtrOutput            `pulumi:"timeZone"`
	Type                    pulumi.StringOutput               `pulumi:"type"`
}


func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.StartTime == nil {
		return nil, errors.New("invalid value for required argument 'StartTime'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:automation:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20151031:Schedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20151031:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:Schedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20190601:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-nextgen:automation/v20200113preview:Schedule"),
		},
	})
	opts = append(opts, aliases)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:automation:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure-native:automation:Schedule", name, id, state, &resource, opts...)
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
	AdvancedSchedule      *AdvancedSchedule `pulumi:"advancedSchedule"`
	AutomationAccountName string            `pulumi:"automationAccountName"`
	Description           *string           `pulumi:"description"`
	ExpiryTime            *string           `pulumi:"expiryTime"`
	Frequency             string            `pulumi:"frequency"`
	Interval              interface{}       `pulumi:"interval"`
	Name                  string            `pulumi:"name"`
	ResourceGroupName     string            `pulumi:"resourceGroupName"`
	ScheduleName          *string           `pulumi:"scheduleName"`
	StartTime             string            `pulumi:"startTime"`
	TimeZone              *string           `pulumi:"timeZone"`
}


type ScheduleArgs struct {
	AdvancedSchedule      AdvancedSchedulePtrInput
	AutomationAccountName pulumi.StringInput
	Description           pulumi.StringPtrInput
	ExpiryTime            pulumi.StringPtrInput
	Frequency             pulumi.StringInput
	Interval              pulumi.Input
	Name                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	ScheduleName          pulumi.StringPtrInput
	StartTime             pulumi.StringInput
	TimeZone              pulumi.StringPtrInput
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
	return reflect.TypeOf((*Schedule)(nil))
}

func (i *Schedule) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i *Schedule) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil))
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
