


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type JobDefinition struct {
	pulumi.CustomResourceState

	AgentName              pulumi.StringPtrOutput   `pulumi:"agentName"`
	AgentResourceId        pulumi.StringOutput      `pulumi:"agentResourceId"`
	CopyMode               pulumi.StringOutput      `pulumi:"copyMode"`
	Description            pulumi.StringPtrOutput   `pulumi:"description"`
	LatestJobRunName       pulumi.StringOutput      `pulumi:"latestJobRunName"`
	LatestJobRunResourceId pulumi.StringOutput      `pulumi:"latestJobRunResourceId"`
	LatestJobRunStatus     pulumi.StringOutput      `pulumi:"latestJobRunStatus"`
	Name                   pulumi.StringOutput      `pulumi:"name"`
	ProvisioningState      pulumi.StringOutput      `pulumi:"provisioningState"`
	SourceName             pulumi.StringOutput      `pulumi:"sourceName"`
	SourceResourceId       pulumi.StringOutput      `pulumi:"sourceResourceId"`
	SourceSubpath          pulumi.StringPtrOutput   `pulumi:"sourceSubpath"`
	SystemData             SystemDataResponseOutput `pulumi:"systemData"`
	TargetName             pulumi.StringOutput      `pulumi:"targetName"`
	TargetResourceId       pulumi.StringOutput      `pulumi:"targetResourceId"`
	TargetSubpath          pulumi.StringPtrOutput   `pulumi:"targetSubpath"`
	Type                   pulumi.StringOutput      `pulumi:"type"`
}


func NewJobDefinition(ctx *pulumi.Context,
	name string, args *JobDefinitionArgs, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.CopyMode == nil {
		return nil, errors.New("invalid value for required argument 'CopyMode'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.SourceName == nil {
		return nil, errors.New("invalid value for required argument 'SourceName'")
	}
	if args.StorageMoverName == nil {
		return nil, errors.New("invalid value for required argument 'StorageMoverName'")
	}
	if args.TargetName == nil {
		return nil, errors.New("invalid value for required argument 'TargetName'")
	}
	var resource JobDefinition
	err := ctx.RegisterResource("azure-native:storagemover/v20220701preview:JobDefinition", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetJobDefinition(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *JobDefinitionState, opts ...pulumi.ResourceOption) (*JobDefinition, error) {
	var resource JobDefinition
	err := ctx.ReadResource("azure-native:storagemover/v20220701preview:JobDefinition", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type jobDefinitionState struct {
}

type JobDefinitionState struct {
}

func (JobDefinitionState) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionState)(nil)).Elem()
}

type jobDefinitionArgs struct {
	AgentName         *string `pulumi:"agentName"`
	CopyMode          string  `pulumi:"copyMode"`
	Description       *string `pulumi:"description"`
	JobDefinitionName *string `pulumi:"jobDefinitionName"`
	ProjectName       string  `pulumi:"projectName"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	SourceName        string  `pulumi:"sourceName"`
	SourceSubpath     *string `pulumi:"sourceSubpath"`
	StorageMoverName  string  `pulumi:"storageMoverName"`
	TargetName        string  `pulumi:"targetName"`
	TargetSubpath     *string `pulumi:"targetSubpath"`
}


type JobDefinitionArgs struct {
	AgentName         pulumi.StringPtrInput
	CopyMode          pulumi.StringInput
	Description       pulumi.StringPtrInput
	JobDefinitionName pulumi.StringPtrInput
	ProjectName       pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	SourceName        pulumi.StringInput
	SourceSubpath     pulumi.StringPtrInput
	StorageMoverName  pulumi.StringInput
	TargetName        pulumi.StringInput
	TargetSubpath     pulumi.StringPtrInput
}

func (JobDefinitionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*jobDefinitionArgs)(nil)).Elem()
}

type JobDefinitionInput interface {
	pulumi.Input

	ToJobDefinitionOutput() JobDefinitionOutput
	ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput
}

func (*JobDefinition) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (i *JobDefinition) ToJobDefinitionOutput() JobDefinitionOutput {
	return i.ToJobDefinitionOutputWithContext(context.Background())
}

func (i *JobDefinition) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(JobDefinitionOutput)
}

type JobDefinitionOutput struct{ *pulumi.OutputState }

func (JobDefinitionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**JobDefinition)(nil)).Elem()
}

func (o JobDefinitionOutput) ToJobDefinitionOutput() JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) ToJobDefinitionOutputWithContext(ctx context.Context) JobDefinitionOutput {
	return o
}

func (o JobDefinitionOutput) AgentName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.AgentName }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) AgentResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.AgentResourceId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) CopyMode() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.CopyMode }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) LatestJobRunName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.LatestJobRunName }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) LatestJobRunResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.LatestJobRunResourceId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) LatestJobRunStatus() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.LatestJobRunStatus }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) SourceName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.SourceName }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) SourceResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.SourceResourceId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) SourceSubpath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.SourceSubpath }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *JobDefinition) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o JobDefinitionOutput) TargetName() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.TargetName }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) TargetResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.TargetResourceId }).(pulumi.StringOutput)
}

func (o JobDefinitionOutput) TargetSubpath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringPtrOutput { return v.TargetSubpath }).(pulumi.StringPtrOutput)
}

func (o JobDefinitionOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *JobDefinition) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(JobDefinitionOutput{})
}
