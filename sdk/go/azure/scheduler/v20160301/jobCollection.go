


package v20160301

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobCollection struct {
	pulumi.CustomResourceState

	Location   pulumi.StringPtrOutput                `pulumi:"location"`
	Name       pulumi.StringPtrOutput                `pulumi:"name"`
	Properties JobCollectionPropertiesResponseOutput `pulumi:"properties"`
	Tags       pulumi.StringMapOutput                `pulumi:"tags"`
	Type       pulumi.StringOutput                   `pulumi:"type"`
}


func NewJobCollection(ctx *pulumi.Context,
	name string, args *JobCollectionArgs, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:scheduler/v20160301:JobCollection"),
		},
		{
			Type: pulumi.String("azure-native:scheduler:JobCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:scheduler:JobCollection"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20140801preview:JobCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:scheduler/v20140801preview:JobCollection"),
		},
		{
			Type: pulumi.String("azure-native:scheduler/v20160101:JobCollection"),
		},
		{
			Type: pulumi.String("azure-nextgen:scheduler/v20160101:JobCollection"),
		},
	})
	opts = append(opts, aliases)
	var resource JobCollection
	err := ctx.RegisterResource("azure-native:scheduler/v20160301:JobCollection", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobCollection(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobCollectionState, opts ...pulumi.ResourceOption) (*JobCollection, error) {
	var resource JobCollection
	err := ctx.ReadResource("azure-native:scheduler/v20160301:JobCollection", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobCollectionState struct {
}

type JobCollectionState struct {
}

func (JobCollectionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCollectionState)(nil)).Elem()
}

type jobCollectionArgs struct {
	JobCollectionName *string                  `pulumi:"jobCollectionName"`
	Location          *string                  `pulumi:"location"`
	Name              *string                  `pulumi:"name"`
	Properties        *JobCollectionProperties `pulumi:"properties"`
	ResourceGroupName string                   `pulumi:"resourceGroupName"`
	Tags              map[string]string        `pulumi:"tags"`
}


type JobCollectionArgs struct {
	JobCollectionName pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Properties        JobCollectionPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	Tags              pulumi.StringMapInput
}

func (JobCollectionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobCollectionArgs)(nil)).Elem()
}

type JobCollectionInput interface {
	pulumi.Input

	ToJobCollectionOutput() JobCollectionOutput
	ToJobCollectionOutputWithContext(ctx context.Context) JobCollectionOutput
}

func (*JobCollection) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollection)(nil))
}

func (i *JobCollection) ToJobCollectionOutput() JobCollectionOutput {
	return i.ToJobCollectionOutputWithContext(context.Background())
}

func (i *JobCollection) ToJobCollectionOutputWithContext(ctx context.Context) JobCollectionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobCollectionOutput)
}

type JobCollectionOutput struct{ *pulumi.OutputState }

func (JobCollectionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*JobCollection)(nil))
}

func (o JobCollectionOutput) ToJobCollectionOutput() JobCollectionOutput {
	return o
}

func (o JobCollectionOutput) ToJobCollectionOutputWithContext(ctx context.Context) JobCollectionOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*JobCollectionInput)(nil)).Elem(), &JobCollection{})
	pulumi.RegisterOutputType(JobCollectionOutput{})
}
