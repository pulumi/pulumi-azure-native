


package v20170401preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type StreamingJob struct {
	pulumi.CustomResourceState

	Cluster                            ClusterInfoResponsePtrOutput       `pulumi:"cluster"`
	CompatibilityLevel                 pulumi.StringPtrOutput             `pulumi:"compatibilityLevel"`
	ContentStoragePolicy               pulumi.StringPtrOutput             `pulumi:"contentStoragePolicy"`
	CreatedDate                        pulumi.StringOutput                `pulumi:"createdDate"`
	DataLocale                         pulumi.StringPtrOutput             `pulumi:"dataLocale"`
	Etag                               pulumi.StringOutput                `pulumi:"etag"`
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrOutput                `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	EventsOutOfOrderMaxDelayInSeconds  pulumi.IntPtrOutput                `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	EventsOutOfOrderPolicy             pulumi.StringPtrOutput             `pulumi:"eventsOutOfOrderPolicy"`
	Externals                          ExternalResponsePtrOutput          `pulumi:"externals"`
	Functions                          FunctionResponseArrayOutput        `pulumi:"functions"`
	Identity                           IdentityResponsePtrOutput          `pulumi:"identity"`
	Inputs                             InputResponseArrayOutput           `pulumi:"inputs"`
	JobId                              pulumi.StringOutput                `pulumi:"jobId"`
	JobState                           pulumi.StringOutput                `pulumi:"jobState"`
	JobStorageAccount                  JobStorageAccountResponsePtrOutput `pulumi:"jobStorageAccount"`
	JobType                            pulumi.StringPtrOutput             `pulumi:"jobType"`
	LastOutputEventTime                pulumi.StringOutput                `pulumi:"lastOutputEventTime"`
	Location                           pulumi.StringPtrOutput             `pulumi:"location"`
	Name                               pulumi.StringOutput                `pulumi:"name"`
	OutputErrorPolicy                  pulumi.StringPtrOutput             `pulumi:"outputErrorPolicy"`
	OutputStartMode                    pulumi.StringPtrOutput             `pulumi:"outputStartMode"`
	OutputStartTime                    pulumi.StringPtrOutput             `pulumi:"outputStartTime"`
	Outputs                            OutputResponseArrayOutput          `pulumi:"outputs"`
	ProvisioningState                  pulumi.StringOutput                `pulumi:"provisioningState"`
	Sku                                StreamingJobSkuResponsePtrOutput   `pulumi:"sku"`
	Tags                               pulumi.StringMapOutput             `pulumi:"tags"`
	Transformation                     TransformationResponsePtrOutput    `pulumi:"transformation"`
	Type                               pulumi.StringOutput                `pulumi:"type"`
}


func NewStreamingJob(ctx *pulumi.Context,
	name string, args *StreamingJobArgs, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:streamanalytics:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20160301:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20200301:StreamingJob"),
		},
		{
			Type: pulumi.String("azure-native:streamanalytics/v20211001preview:StreamingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource StreamingJob
	err := ctx.RegisterResource("azure-native:streamanalytics/v20170401preview:StreamingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetStreamingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *StreamingJobState, opts ...pulumi.ResourceOption) (*StreamingJob, error) {
	var resource StreamingJob
	err := ctx.ReadResource("azure-native:streamanalytics/v20170401preview:StreamingJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type streamingJobState struct {
}

type StreamingJobState struct {
}

func (StreamingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobState)(nil)).Elem()
}

type streamingJobArgs struct {
	Cluster                            *ClusterInfo       `pulumi:"cluster"`
	CompatibilityLevel                 *string            `pulumi:"compatibilityLevel"`
	ContentStoragePolicy               *string            `pulumi:"contentStoragePolicy"`
	DataLocale                         *string            `pulumi:"dataLocale"`
	EventsLateArrivalMaxDelayInSeconds *int               `pulumi:"eventsLateArrivalMaxDelayInSeconds"`
	EventsOutOfOrderMaxDelayInSeconds  *int               `pulumi:"eventsOutOfOrderMaxDelayInSeconds"`
	EventsOutOfOrderPolicy             *string            `pulumi:"eventsOutOfOrderPolicy"`
	Externals                          *External          `pulumi:"externals"`
	Functions                          []FunctionType     `pulumi:"functions"`
	Identity                           *Identity          `pulumi:"identity"`
	Inputs                             []InputType        `pulumi:"inputs"`
	JobName                            *string            `pulumi:"jobName"`
	JobStorageAccount                  *JobStorageAccount `pulumi:"jobStorageAccount"`
	JobType                            *string            `pulumi:"jobType"`
	Location                           *string            `pulumi:"location"`
	OutputErrorPolicy                  *string            `pulumi:"outputErrorPolicy"`
	OutputStartMode                    *string            `pulumi:"outputStartMode"`
	OutputStartTime                    *string            `pulumi:"outputStartTime"`
	Outputs                            []OutputType       `pulumi:"outputs"`
	ResourceGroupName                  string             `pulumi:"resourceGroupName"`
	Sku                                *StreamingJobSku   `pulumi:"sku"`
	Tags                               map[string]string  `pulumi:"tags"`
	Transformation                     *Transformation    `pulumi:"transformation"`
}


type StreamingJobArgs struct {
	Cluster                            ClusterInfoPtrInput
	CompatibilityLevel                 pulumi.StringPtrInput
	ContentStoragePolicy               pulumi.StringPtrInput
	DataLocale                         pulumi.StringPtrInput
	EventsLateArrivalMaxDelayInSeconds pulumi.IntPtrInput
	EventsOutOfOrderMaxDelayInSeconds  pulumi.IntPtrInput
	EventsOutOfOrderPolicy             pulumi.StringPtrInput
	Externals                          ExternalPtrInput
	Functions                          FunctionTypeArrayInput
	Identity                           IdentityPtrInput
	Inputs                             InputTypeArrayInput
	JobName                            pulumi.StringPtrInput
	JobStorageAccount                  JobStorageAccountPtrInput
	JobType                            pulumi.StringPtrInput
	Location                           pulumi.StringPtrInput
	OutputErrorPolicy                  pulumi.StringPtrInput
	OutputStartMode                    pulumi.StringPtrInput
	OutputStartTime                    pulumi.StringPtrInput
	Outputs                            OutputTypeArrayInput
	ResourceGroupName                  pulumi.StringInput
	Sku                                StreamingJobSkuPtrInput
	Tags                               pulumi.StringMapInput
	Transformation                     TransformationPtrInput
}

func (StreamingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*streamingJobArgs)(nil)).Elem()
}

type StreamingJobInput interface {
	pulumi.Input

	ToStreamingJobOutput() StreamingJobOutput
	ToStreamingJobOutputWithContext(ctx context.Context) StreamingJobOutput
}

func (*StreamingJob) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJob)(nil)).Elem()
}

func (i *StreamingJob) ToStreamingJobOutput() StreamingJobOutput {
	return i.ToStreamingJobOutputWithContext(context.Background())
}

func (i *StreamingJob) ToStreamingJobOutputWithContext(ctx context.Context) StreamingJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StreamingJobOutput)
}

type StreamingJobOutput struct{ *pulumi.OutputState }

func (StreamingJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StreamingJob)(nil)).Elem()
}

func (o StreamingJobOutput) ToStreamingJobOutput() StreamingJobOutput {
	return o
}

func (o StreamingJobOutput) ToStreamingJobOutputWithContext(ctx context.Context) StreamingJobOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(StreamingJobOutput{})
}
