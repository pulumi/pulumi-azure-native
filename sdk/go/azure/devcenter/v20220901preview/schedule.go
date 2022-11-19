


package v20220901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.CustomResourceState

	Frequency         pulumi.StringOutput      `pulumi:"frequency"`
	Name              pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	State             pulumi.StringPtrOutput   `pulumi:"state"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Time              pulumi.StringOutput      `pulumi:"time"`
	TimeZone          pulumi.StringOutput      `pulumi:"timeZone"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Frequency == nil {
		return nil, errors.New("invalid value for required argument 'Frequency'")
	}
	if args.PoolName == nil {
		return nil, errors.New("invalid value for required argument 'PoolName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Time == nil {
		return nil, errors.New("invalid value for required argument 'Time'")
	}
	if args.TimeZone == nil {
		return nil, errors.New("invalid value for required argument 'TimeZone'")
	}
	if args.Type == nil {
		return nil, errors.New("invalid value for required argument 'Type'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:devcenter:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20220801preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221012preview:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:devcenter/v20221111preview:Schedule"),
		},
	})
	opts = append(opts, aliases)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:devcenter/v20220901preview:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure-native:devcenter/v20220901preview:Schedule", name, id, state, &resource, opts...)
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
	Frequency         string  `pulumi:"frequency"`
	PoolName          string  `pulumi:"poolName"`
	ProjectName       string  `pulumi:"projectName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	ScheduleName      *string `pulumi:"scheduleName"`
	State             *string `pulumi:"state"`
	Time              string  `pulumi:"time"`
	TimeZone          string  `pulumi:"timeZone"`
	Top               *int    `pulumi:"top"`
	Type              string  `pulumi:"type"`
}


type ScheduleArgs struct {
	Frequency         pulumi.StringInput
	PoolName          pulumi.StringInput
	ProjectName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	ScheduleName      pulumi.StringPtrInput
	State             pulumi.StringPtrInput
	Time              pulumi.StringInput
	TimeZone          pulumi.StringInput
	Top               pulumi.IntPtrInput
	Type              pulumi.StringInput
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

func (o ScheduleOutput) Frequency() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Frequency }).(pulumi.StringOutput)
}

func (o ScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ScheduleOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o ScheduleOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Schedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduleOutput) Time() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Time }).(pulumi.StringOutput)
}

func (o ScheduleOutput) TimeZone() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.TimeZone }).(pulumi.StringOutput)
}

func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
