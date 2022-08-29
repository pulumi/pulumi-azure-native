


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
	if args.Properties != nil {
		args.Properties = args.Properties.ToJobDetailsPtrOutput().ApplyT(func(v *JobDetails) *JobDetails { return v.Defaults() }).(JobDetailsPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:importexport/v20161101:Job"),
		},
		{
			Type: pulumi.String("azure-native:importexport/v20200801:Job"),
		},
		{
			Type: pulumi.String("azure-native:importexport/v20210101:Job"),
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

func (o JobOutput) Identity() IdentityDetailsResponsePtrOutput {
	return o.ApplyT(func(v *Job) IdentityDetailsResponsePtrOutput { return v.Identity }).(IdentityDetailsResponsePtrOutput)
}

func (o JobOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Job) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o JobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobOutput) Properties() JobDetailsResponseOutput {
	return o.ApplyT(func(v *Job) JobDetailsResponseOutput { return v.Properties }).(JobDetailsResponseOutput)
}

func (o JobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *Job) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o JobOutput) Tags() pulumi.AnyOutput {
	return o.ApplyT(func(v *Job) pulumi.AnyOutput { return v.Tags }).(pulumi.AnyOutput)
}

func (o JobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Job) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobOutput{})
}
