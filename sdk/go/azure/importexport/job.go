


package importexport

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Job struct {
	pulumi.CustomResourceState

	Identity   IdentityDetailsResponsePtrOutput `pulumi:"identity"`
	Location   pulumi.StringPtrOutput           `pulumi:"location"`
	Name       pulumi.StringOutput              `pulumi:"name"`
	Properties JobDetailsResponseOutput         `pulumi:"properties"`
	SystemData SystemDataResponseOutput         `pulumi:"systemData"`
	Tags       pulumi.AnyOutput                 `pulumi:"tags"`
	Type       pulumi.StringOutput              `pulumi:"type"`
}


func NewJob(ctx *pulumi.Context,
	name string, args *JobArgs, opts ...pulumi.ResourceOption) (*Job, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:importexport:Job"),
		},
		{
			Type: pulumi.String("azure-native:importexport/v20161101:Job"),
		},
		{
			Type: pulumi.String("azure-nextgen:importexport/v20161101:Job"),
		},
		{
			Type: pulumi.String("azure-native:importexport/v20200801:Job"),
		},
		{
			Type: pulumi.String("azure-nextgen:importexport/v20200801:Job"),
		},
		{
			Type: pulumi.String("azure-native:importexport/v20210101:Job"),
		},
		{
			Type: pulumi.String("azure-nextgen:importexport/v20210101:Job"),
		},
	})
	opts = append(opts, aliases)
	var resource Job
	err := ctx.RegisterResource("azure-native:importexport:Job", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobState, opts ...pulumi.ResourceOption) (*Job, error) {
	var resource Job
	err := ctx.ReadResource("azure-native:importexport:Job", name, id, state, &resource, opts...)
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
	JobName           *string     `pulumi:"jobName"`
	Location          *string     `pulumi:"location"`
	Properties        *JobDetails `pulumi:"properties"`
	ResourceGroupName string      `pulumi:"resourceGroupName"`
	Tags              interface{} `pulumi:"tags"`
}


type JobArgs struct {
	JobName           pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Properties        JobDetailsPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.Input
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
	return reflect.TypeOf((*Job)(nil))
}

func (i *Job) ToJobOutput() JobOutput {
	return i.ToJobOutputWithContext(context.Background())
}

func (i *Job) ToJobOutputWithContext(ctx context.Context) JobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobOutput)
}

type JobOutput struct{ *pulumi.OutputState }

func (JobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Job)(nil))
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
