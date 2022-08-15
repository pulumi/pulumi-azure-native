


package labservices

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput                `pulumi:"name"`
	Notes             pulumi.StringPtrOutput             `pulumi:"notes"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	RecurrencePattern RecurrencePatternResponsePtrOutput `pulumi:"recurrencePattern"`
	StartAt           pulumi.StringPtrOutput             `pulumi:"startAt"`
	StopAt            pulumi.StringOutput                `pulumi:"stopAt"`
	SystemData        SystemDataResponseOutput           `pulumi:"systemData"`
	TimeZoneId        pulumi.StringOutput                `pulumi:"timeZoneId"`
	Type              pulumi.StringOutput                `pulumi:"type"`
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
	if args.StopAt == nil {
		return nil, errors.New("invalid value for required argument 'StopAt'")
	}
	if args.TimeZoneId == nil {
		return nil, errors.New("invalid value for required argument 'TimeZoneId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:labservices/v20211001preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20211115preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:labservices/v20220801:Schedule"),
		},
	})
	opts = append(opts, aliases)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:labservices:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure-native:labservices:Schedule", name, id, state, &resource, opts...)
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
	LabName           string             `pulumi:"labName"`
	Notes             *string            `pulumi:"notes"`
	RecurrencePattern *RecurrencePattern `pulumi:"recurrencePattern"`
	ResourceGroupName string             `pulumi:"resourceGroupName"`
	ScheduleName      *string            `pulumi:"scheduleName"`
	StartAt           *string            `pulumi:"startAt"`
	StopAt            string             `pulumi:"stopAt"`
	TimeZoneId        string             `pulumi:"timeZoneId"`
}


type ScheduleArgs struct {
	LabName           pulumi.StringInput
	Notes             pulumi.StringPtrInput
	RecurrencePattern RecurrencePatternPtrInput
	ResourceGroupName pulumi.StringInput
	ScheduleName      pulumi.StringPtrInput
	StartAt           pulumi.StringPtrInput
	StopAt            pulumi.StringInput
	TimeZoneId        pulumi.StringInput
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

func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduleOutput) Notes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.Notes }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScheduleOutput) RecurrencePattern() RecurrencePatternResponsePtrOutput {
	return o.ApplyT(func(v *Schedule) RecurrencePatternResponsePtrOutput { return v.RecurrencePattern }).(RecurrencePatternResponsePtrOutput)
}

func (o ScheduleOutput) StartAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.StartAt }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) StopAt() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.StopAt }).(pulumi.StringOutput)
}

func (o ScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Schedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduleOutput) TimeZoneId() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.TimeZoneId }).(pulumi.StringOutput)
}

func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
