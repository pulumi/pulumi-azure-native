


package containerregistry

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type TaskRun struct {
	pulumi.CustomResourceState

	ForceUpdateTag    pulumi.StringPtrOutput              `pulumi:"forceUpdateTag"`
	Identity          IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location          pulumi.StringPtrOutput              `pulumi:"location"`
	Name              pulumi.StringOutput                 `pulumi:"name"`
	ProvisioningState pulumi.StringOutput                 `pulumi:"provisioningState"`
	RunRequest        pulumi.AnyOutput                    `pulumi:"runRequest"`
	RunResult         RunResponseOutput                   `pulumi:"runResult"`
	SystemData        SystemDataResponseOutput            `pulumi:"systemData"`
	Type              pulumi.StringOutput                 `pulumi:"type"`
}


func NewTaskRun(ctx *pulumi.Context,
	name string, args *TaskRunArgs, opts ...pulumi.ResourceOption) (*TaskRun, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry/v20190601preview:TaskRun"),
		},
	})
	opts = append(opts, aliases)
	var resource TaskRun
	err := ctx.RegisterResource("azure-native:containerregistry:TaskRun", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTaskRun(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskRunState, opts ...pulumi.ResourceOption) (*TaskRun, error) {
	var resource TaskRun
	err := ctx.ReadResource("azure-native:containerregistry:TaskRun", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type taskRunState struct {
}

type TaskRunState struct {
}

func (TaskRunState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskRunState)(nil)).Elem()
}

type taskRunArgs struct {
	ForceUpdateTag    *string             `pulumi:"forceUpdateTag"`
	Identity          *IdentityProperties `pulumi:"identity"`
	Location          *string             `pulumi:"location"`
	RegistryName      string              `pulumi:"registryName"`
	ResourceGroupName string              `pulumi:"resourceGroupName"`
	RunRequest        interface{}         `pulumi:"runRequest"`
	TaskRunName       *string             `pulumi:"taskRunName"`
}


type TaskRunArgs struct {
	ForceUpdateTag    pulumi.StringPtrInput
	Identity          IdentityPropertiesPtrInput
	Location          pulumi.StringPtrInput
	RegistryName      pulumi.StringInput
	ResourceGroupName pulumi.StringInput
	RunRequest        pulumi.Input
	TaskRunName       pulumi.StringPtrInput
}

func (TaskRunArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskRunArgs)(nil)).Elem()
}

type TaskRunInput interface {
	pulumi.Input

	ToTaskRunOutput() TaskRunOutput
	ToTaskRunOutputWithContext(ctx context.Context) TaskRunOutput
}

func (*TaskRun) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskRun)(nil)).Elem()
}

func (i *TaskRun) ToTaskRunOutput() TaskRunOutput {
	return i.ToTaskRunOutputWithContext(context.Background())
}

func (i *TaskRun) ToTaskRunOutputWithContext(ctx context.Context) TaskRunOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskRunOutput)
}

type TaskRunOutput struct{ *pulumi.OutputState }

func (TaskRunOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TaskRun)(nil)).Elem()
}

func (o TaskRunOutput) ToTaskRunOutput() TaskRunOutput {
	return o
}

func (o TaskRunOutput) ToTaskRunOutputWithContext(ctx context.Context) TaskRunOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(TaskRunOutput{})
}
