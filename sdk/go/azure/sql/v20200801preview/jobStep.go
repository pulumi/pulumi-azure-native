


package v20200801preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobStep struct {
	pulumi.CustomResourceState

	Action           JobStepActionResponseOutput              `pulumi:"action"`
	Credential       pulumi.StringOutput                      `pulumi:"credential"`
	ExecutionOptions JobStepExecutionOptionsResponsePtrOutput `pulumi:"executionOptions"`
	Name             pulumi.StringOutput                      `pulumi:"name"`
	Output           JobStepOutputResponsePtrOutput           `pulumi:"output"`
	StepId           pulumi.IntPtrOutput                      `pulumi:"stepId"`
	TargetGroup      pulumi.StringOutput                      `pulumi:"targetGroup"`
	Type             pulumi.StringOutput                      `pulumi:"type"`
}


func NewJobStep(ctx *pulumi.Context,
	name string, args *JobStepArgs, opts ...pulumi.ResourceOption) (*JobStep, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Action == nil {
		return nil, errors.New("invalid value for required argument 'Action'")
	}
	if args.Credential == nil {
		return nil, errors.New("invalid value for required argument 'Credential'")
	}
	if args.JobAgentName == nil {
		return nil, errors.New("invalid value for required argument 'JobAgentName'")
	}
	if args.JobName == nil {
		return nil, errors.New("invalid value for required argument 'JobName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.ServerName == nil {
		return nil, errors.New("invalid value for required argument 'ServerName'")
	}
	if args.TargetGroup == nil {
		return nil, errors.New("invalid value for required argument 'TargetGroup'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:sql/v20200801preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql:JobStep"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20170301preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20170301preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20200202preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20200202preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20201101preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20201101preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210201preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210201preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-native:sql/v20210501preview:JobStep"),
		},
		{
			Type: pulumi.String("azure-nextgen:sql/v20210501preview:JobStep"),
		},
	})
	opts = append(opts, aliases)
	var resource JobStep
	err := ctx.RegisterResource("azure-native:sql/v20200801preview:JobStep", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobStep(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobStepState, opts ...pulumi.ResourceOption) (*JobStep, error) {
	var resource JobStep
	err := ctx.ReadResource("azure-native:sql/v20200801preview:JobStep", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobStepState struct {
}

type JobStepState struct {
}

func (JobStepState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobStepState)(nil)).Elem()
}

type jobStepArgs struct {
	Action            JobStepAction            `pulumi:"action"`
	Credential        string                   `pulumi:"credential"`
	ExecutionOptions  *JobStepExecutionOptions `pulumi:"executionOptions"`
	JobAgentName      string                   `pulumi:"jobAgentName"`
	JobName           string                   `pulumi:"jobName"`
	Output            *JobStepOutputType       `pulumi:"output"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	ServerName        string                   `pulumi:"serverName"`
	StepId            *int                     `pulumi:"stepId"`
	StepName          *string                  `pulumi:"stepName"`
	TargetGroup       string                   `pulumi:"targetGroup"`
}


type JobStepArgs struct {
	Action            JobStepActionInput
	Credential        pulumi.StringInput
	ExecutionOptions  JobStepExecutionOptionsPtrInput
	JobAgentName      pulumi.StringInput
	JobName           pulumi.StringInput
	Output            JobStepOutputTypePtrInput
	ResourceGroupName pulumi.StringInput
	ServerName        pulumi.StringInput
	StepId            pulumi.IntPtrInput
	StepName          pulumi.StringPtrInput
	TargetGroup       pulumi.StringInput
}

func (JobStepArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobStepArgs)(nil)).Elem()
}

type JobStepInput interface {
	pulumi.Input

	ToJobStepOutput() JobStepOutput
	ToJobStepOutputWithContext(ctx context.Context) JobStepOutput
}

func (*JobStep) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStep)(nil))
}

func (i *JobStep) ToJobStepOutput() JobStepOutput {
	return i.ToJobStepOutputWithContext(context.Background())
}

func (i *JobStep) ToJobStepOutputWithContext(ctx context.Context) JobStepOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobStepOutput)
}

type JobStepOutput struct{ *pulumi.OutputState }

func (JobStepOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobStep)(nil))
}

func (o JobStepOutput) ToJobStepOutput() JobStepOutput {
	return o
}

func (o JobStepOutput) ToJobStepOutputWithContext(ctx context.Context) JobStepOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobStepInput)(nil)).Elem(), &JobStep{})
	pulumi.RegisterOutputType(JobStepOutput{})
}
