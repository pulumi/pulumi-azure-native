


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LabelingJob struct {
	pulumi.CustomResourceState

	LabelingJobProperties LabelingJobResponseOutput `pulumi:"labelingJobProperties"`
	Name                  pulumi.StringOutput       `pulumi:"name"`
	SystemData            SystemDataResponseOutput  `pulumi:"systemData"`
	Type                  pulumi.StringOutput       `pulumi:"type"`
}


func NewLabelingJob(ctx *pulumi.Context,
	name string, args *LabelingJobArgs, opts ...pulumi.ResourceOption) (*LabelingJob, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.LabelingJobProperties == nil {
		return nil, errors.New("invalid value for required argument 'LabelingJobProperties'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.LabelingJobProperties = args.LabelingJobProperties.ToLabelingJobTypeOutput().ApplyT(func(v LabelingJobType) LabelingJobType { return *v.Defaults() }).(LabelingJobTypeOutput)
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:machinelearningservices:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20200901preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20210301preview:LabelingJob"),
		},
		{
			Type: pulumi.String("azure-native:machinelearningservices/v20220601preview:LabelingJob"),
		},
	})
	opts = append(opts, aliases)
	var resource LabelingJob
	err := ctx.RegisterResource("azure-native:machinelearningservices/v20221001preview:LabelingJob", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetLabelingJob(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *LabelingJobState, opts ...pulumi.ResourceOption) (*LabelingJob, error) {
	var resource LabelingJob
	err := ctx.ReadResource("azure-native:machinelearningservices/v20221001preview:LabelingJob", name, id, state, &resource, opts...)
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
	Id                    *string         `pulumi:"id"`
	LabelingJobProperties LabelingJobType `pulumi:"labelingJobProperties"`
	ResourceGroupName     string          `pulumi:"resourceGroupName"`
	WorkspaceName         string          `pulumi:"workspaceName"`
}


type LabelingJobArgs struct {
	Id                    pulumi.StringPtrInput
	LabelingJobProperties LabelingJobTypeInput
	ResourceGroupName     pulumi.StringInput
	WorkspaceName         pulumi.StringInput
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
	return reflect.TypeOf((**LabelingJob)(nil)).Elem()
}

func (i *LabelingJob) ToLabelingJobOutput() LabelingJobOutput {
	return i.ToLabelingJobOutputWithContext(context.Background())
}

func (i *LabelingJob) ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LabelingJobOutput)
}

type LabelingJobOutput struct{ *pulumi.OutputState }

func (LabelingJobOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LabelingJob)(nil)).Elem()
}

func (o LabelingJobOutput) ToLabelingJobOutput() LabelingJobOutput {
	return o
}

func (o LabelingJobOutput) ToLabelingJobOutputWithContext(ctx context.Context) LabelingJobOutput {
	return o
}

func (o LabelingJobOutput) LabelingJobProperties() LabelingJobResponseOutput {
	return o.ApplyT(func(v *LabelingJob) LabelingJobResponseOutput { return v.LabelingJobProperties }).(LabelingJobResponseOutput)
}

func (o LabelingJobOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelingJob) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o LabelingJobOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *LabelingJob) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LabelingJobOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *LabelingJob) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LabelingJobOutput{})
}
