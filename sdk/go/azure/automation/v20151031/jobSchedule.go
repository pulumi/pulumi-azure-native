


package v20151031

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobSchedule struct {
	pulumi.CustomResourceState

	JobScheduleId pulumi.StringPtrOutput                       `pulumi:"jobScheduleId"`
	Name          pulumi.StringOutput                          `pulumi:"name"`
	Parameters    pulumi.StringMapOutput                       `pulumi:"parameters"`
	RunOn         pulumi.StringPtrOutput                       `pulumi:"runOn"`
	Runbook       RunbookAssociationPropertyResponsePtrOutput  `pulumi:"runbook"`
	Schedule      ScheduleAssociationPropertyResponsePtrOutput `pulumi:"schedule"`
	Type          pulumi.StringOutput                          `pulumi:"type"`
}


func NewJobSchedule(ctx *pulumi.Context,
	name string, args *JobScheduleArgs, opts ...pulumi.ResourceOption) (*JobSchedule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AutomationAccountName == nil {
		return nil, errors.New("invalid value for required argument 'AutomationAccountName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Runbook == nil {
		return nil, errors.New("invalid value for required argument 'Runbook'")
	}
	if args.Schedule == nil {
		return nil, errors.New("invalid value for required argument 'Schedule'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:automation:JobSchedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20190601:JobSchedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20200113preview:JobSchedule"),
		},
		{
			Type: pulumi.String("azure-native:automation/v20220808:JobSchedule"),
		},
	})
	opts = append(opts, aliases)
	var resource JobSchedule
	err := ctx.RegisterResource("azure-native:automation/v20151031:JobSchedule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobSchedule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobScheduleState, opts ...pulumi.ResourceOption) (*JobSchedule, error) {
	var resource JobSchedule
	err := ctx.ReadResource("azure-native:automation/v20151031:JobSchedule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobScheduleState struct {
}

type JobScheduleState struct {
}

func (JobScheduleState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobScheduleState)(nil)).Elem()
}

type jobScheduleArgs struct {
	AutomationAccountName string                      `pulumi:"automationAccountName"`
	JobScheduleId         *string                     `pulumi:"jobScheduleId"`
	Parameters            map[string]string           `pulumi:"parameters"`
	ResourceGroupName     string                      `pulumi:"resourceGroupName"`
	RunOn                 *string                     `pulumi:"runOn"`
	Runbook               RunbookAssociationProperty  `pulumi:"runbook"`
	Schedule              ScheduleAssociationProperty `pulumi:"schedule"`
}


type JobScheduleArgs struct {
	AutomationAccountName pulumi.StringInput
	JobScheduleId         pulumi.StringPtrInput
	Parameters            pulumi.StringMapInput
	ResourceGroupName     pulumi.StringInput
	RunOn                 pulumi.StringPtrInput
	Runbook               RunbookAssociationPropertyInput
	Schedule              ScheduleAssociationPropertyInput
}

func (JobScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobScheduleArgs)(nil)).Elem()
}

type JobScheduleInput interface {
	pulumi.Input

	ToJobScheduleOutput() JobScheduleOutput
	ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput
}

func (*JobSchedule) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil)).Elem()
}

func (i *JobSchedule) ToJobScheduleOutput() JobScheduleOutput {
	return i.ToJobScheduleOutputWithContext(context.Background())
}

func (i *JobSchedule) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobScheduleOutput)
}

type JobScheduleOutput struct{ *pulumi.OutputState }

func (JobScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobSchedule)(nil)).Elem()
}

func (o JobScheduleOutput) ToJobScheduleOutput() JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) ToJobScheduleOutputWithContext(ctx context.Context) JobScheduleOutput {
	return o
}

func (o JobScheduleOutput) JobScheduleId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) pulumi.StringPtrOutput { return v.JobScheduleId }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobSchedule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobScheduleOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *JobSchedule) pulumi.StringMapOutput { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o JobScheduleOutput) RunOn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobSchedule) pulumi.StringPtrOutput { return v.RunOn }).(pulumi.StringPtrOutput)
}

func (o JobScheduleOutput) Runbook() RunbookAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *JobSchedule) RunbookAssociationPropertyResponsePtrOutput { return v.Runbook }).(RunbookAssociationPropertyResponsePtrOutput)
}

func (o JobScheduleOutput) Schedule() ScheduleAssociationPropertyResponsePtrOutput {
	return o.ApplyT(func(v *JobSchedule) ScheduleAssociationPropertyResponsePtrOutput { return v.Schedule }).(ScheduleAssociationPropertyResponsePtrOutput)
}

func (o JobScheduleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobSchedule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobScheduleOutput{})
}
