


package v20200901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LabelingJob struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                 `pulumi:"name"`
	Properties LabelingJobPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput            `pulumi:"systemData"`
	Type       pulumi.StringOutput                 `pulumi:"type"`
}


func NewLabelingJob(ctx *pulumi.Context,
	name string, args *LabelingJobArgs, opts ...pulumi.ResourceOption) (*LabelingJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20200901preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-nextgen:machinelearningservices/v20210301preview:LabelingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource LabelingJob
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20200901preview:LabelingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLabelingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelingJobState, opts ...pulumi.ResourceOption) (*LabelingJob, error) {
	var resource LabelingJob
	err := ctx.ReadResource("azure-native:machinelearningservices/v20200901preview:LabelingJob", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type labelingJobState struct {
}

type LabelingJobState struct {
}

func (LabelingJobState) ElementType() reflect.Type {
	return reflect.TypeOf((*labelingJobState)(nil)).Elem()
}

type labelingJobArgs struct {
	LabelingJobId     *string                `pulumi:"labelingJobId"`
	Properties        *LabelingJobProperties `pulumi:"properties"`
	ResourceGroupName string                 `pulumi:"resourceGroupName"`
	WorkspaceName     string                 `pulumi:"workspaceName"`
}


type LabelingJobArgs struct {
	LabelingJobId     pulumi.StringPtrInput
	Properties        LabelingJobPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	WorkspaceName     pulumi.StringInput
}

func (LabelingJobArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*labelingJobArgs)(nil)).Elem()
}

type LabelingJobInput interface {
	pulumi.Input

	ToLabelingJobOutput() LabelingJobOutput
	ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput
}

func (*LabelingJob) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJob)(nil))
}

func (i *LabelingJob) ToLabelingJobOutput() LabelingJobOutput {
	return i.ToLabelingJobOutputWithContext(context.Background())
}

func (i *LabelingJob) ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobOutput)
}

type LabelingJobOutput struct{ *pulumi.OutputState }

func (LabelingJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LabelingJob)(nil))
}

func (o LabelingJobOutput) ToLabelingJobOutput() LabelingJobOutput {
	return o
}

func (o LabelingJobOutput) ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*LabelingJobInput)(nil)).Elem(), &LabelingJob{})
	pulumi.RegisterOutputType(LabelingJobOutput{})
}
