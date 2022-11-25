


package v20220601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Schedule struct {
	pulumi.CustomResourceState

	Name               pulumi.StringOutput      `pulumi:"name"`
	ScheduleProperties ScheduleResponseOutput   `pulumi:"scheduleProperties"`
	SystemData         SystemDataResponseOutput `pulumi:"systemData"`
	Type               pulumi.StringOutput      `pulumi:"type"`
}


func NewSchedule(ctx *pulumi.Context,
	name string, args *ScheduleArgs, opts ...pulumi.ResourceOption) (*Schedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ScheduleProperties == nil {
		return nil, errors.New("invalid value for required argument 'ScheduleProperties'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.ScheduleProperties = args.ScheduleProperties.ToScheduleTypeOutput().ApplyT(func(v ScheduleType) ScheduleType { return *v.Defaults() }).(ScheduleTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001:Schedule"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20221001preview:Schedule"),
		},
	})
	opts = append(opts, aliases)
	var resource Schedule
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20220601preview:Schedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ScheduleState, opts ...pulumi.ResourceOption) (*Schedule, error) {
	var resource Schedule
	err := ctx.ReadResource("azure-native:machinelearningservices/v20220601preview:Schedule", name, id, state, &resource, opts...)
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
	Name               *string      `pulumi:"name"`
	ResourceGroupName  string       `pulumi:"resourceGroupName"`
	ScheduleProperties ScheduleType `pulumi:"scheduleProperties"`
	WorkspaceName      string       `pulumi:"workspaceName"`
}


type ScheduleArgs struct {
	Name               pulumi.StringPtrInput
	ResourceGroupName  pulumi.StringInput
	ScheduleProperties ScheduleTypeInput
	WorkspaceName      pulumi.StringInput
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

func (o ScheduleOutput) ScheduleProperties() ScheduleResponseOutput {
	return o.ApplyT(func(v *Schedule) ScheduleResponseOutput { return v.ScheduleProperties }).(ScheduleResponseOutput)
}

func (o ScheduleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Schedule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Schedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ScheduleOutput{})
}
