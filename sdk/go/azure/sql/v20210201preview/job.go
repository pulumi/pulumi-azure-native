


package v20210201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Job struct {
	pulumi.CustomResourceState

	Description pulumi.StringPtrOutput       `pulumi:"description"`
	Name        pulumi.StringOutput          `pulumi:"name"`
	Schedule    JobScheduleResponsePtrOutput `pulumi:"schedule"`
	Type        pulumi.StringOutput          `pulumi:"type"`
	Version     pulumi.IntOutput             `pulumi:"version"`
}


func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if isZero(args.Description) {
		args.Description = pulumi.StringPtr("")
	}
	if args.Schedule != nil {
		args.Schedule = args.Schedule.ToJobSchedulePtrOutput().ApplyT(func(v *JobSchedule) *JobSchedule { return v.Defaults() }).(JobSchedulePtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:sql:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200801preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:Job"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:sql/v20210201preview:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:sql/v20210201preview:Job", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobState struct {
}

type JobState struct {
}

func (JobState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobState)(nil)).Elem()
}

type jobArgs struct {
	Description       *string      `pulumi:"description"`
	JobAgentName      string       `pulumi:"jobAgentName"`
	JobName           *string      `pulumi:"jobName"`
	ResourceGroupName string       `pulumi:"resourceGroupName"`
	Schedule          *JobSchedule `pulumi:"schedule"`
	ServerName        string       `pulumi:"serverName"`
}


type JobArgs struct {
	Description       pulumi.StringPtrInput
	JobAgentName      pulumi.StringInput
	JobName           pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	Schedule          JobSchedulePtrInput
	ServerName        pulumi.StringInput
}

func (JobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobArgs)(nil)).Elem()
}

type JobInput interface {
	pulumi.Input

	ToJobOutput() JobOutput
	ToJobOutputWithContext(ctx context.Context) JobOutput
}

func (*Job) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Job)(nil)).Elem()
}

func (o JobOutput) ToJobOutput() JobOutput {
	return o
}

func (o JobOutput) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
