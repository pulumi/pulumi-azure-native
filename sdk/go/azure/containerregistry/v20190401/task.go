


package v20190401

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Task struct {
	pulumi.CustomResourceState

	AgentConfiguration AgentPropertiesResponsePtrOutput    `pulumi:"agentConfiguration"`
	CreationDate       pulumi.StringOutput                 `pulumi:"creationDate"`
	Credentials        CredentialsResponsePtrOutput        `pulumi:"credentials"`
	Identity           IdentityPropertiesResponsePtrOutput `pulumi:"identity"`
	Location           pulumi.StringOutput                 `pulumi:"location"`
	Name               pulumi.StringOutput                 `pulumi:"name"`
	Platform           PlatformPropertiesResponseOutput    `pulumi:"platform"`
	ProvisioningState  pulumi.StringOutput                 `pulumi:"provisioningState"`
	Status             pulumi.StringPtrOutput              `pulumi:"status"`
	Step               pulumi.AnyOutput                    `pulumi:"step"`
	Tags               pulumi.StringMapOutput              `pulumi:"tags"`
	Timeout            pulumi.IntPtrOutput                 `pulumi:"timeout"`
	Trigger            TriggerPropertiesResponsePtrOutput  `pulumi:"trigger"`
	Type               pulumi.StringOutput                 `pulumi:"type"`
}


func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Platform == nil {
		return nil, errors.New("invalid value for required argument 'Platform'")
	}
	if args.RegistryName == nil {
		return nil, errors.New("invalid value for required argument 'RegistryName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Step == nil {
		return nil, errors.New("invalid value for required argument 'Step'")
	}
	if args.Timeout == nil {
		args.Timeout = pulumi.IntPtr(3600)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20190401:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20180201preview:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20180201preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20180901:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20180901:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190601preview:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:containerregistry/v20190601preview:Task"),
		},
	})
	opts = append(opts, aliases)
	var resource Task
	err := ctx.RegisterResource("azure-native:containerregistry/v20190401:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("azure-native:containerregistry/v20190401:Task", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type taskState struct {
}

type TaskState struct {
}

func (TaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*taskState)(nil)).Elem()
}

type taskArgs struct {
	AgentConfiguration *AgentProperties    `pulumi:"agentConfiguration"`
	Credentials        *Credentials        `pulumi:"credentials"`
	Identity           *IdentityProperties `pulumi:"identity"`
	Location           *string             `pulumi:"location"`
	Platform           PlatformProperties  `pulumi:"platform"`
	RegistryName       string              `pulumi:"registryName"`
	ResourceGroupName  string              `pulumi:"resourceGroupName"`
	Status             *string             `pulumi:"status"`
	Step               interface{}         `pulumi:"step"`
	Tags               map[string]string   `pulumi:"tags"`
	TaskName           *string             `pulumi:"taskName"`
	Timeout            *int                `pulumi:"timeout"`
	Trigger            *TriggerProperties  `pulumi:"trigger"`
}


type TaskArgs struct {
	AgentConfiguration AgentPropertiesPtrInput
	Credentials        CredentialsPtrInput
	Identity           IdentityPropertiesPtrInput
	Location           pulumi.StringPtrInput
	Platform           PlatformPropertiesInput
	RegistryName       pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	Status             pulumi.StringPtrInput
	Step               pulumi.Input
	Tags               pulumi.StringMapInput
	TaskName           pulumi.StringPtrInput
	Timeout            pulumi.IntPtrInput
	Trigger            TriggerPropertiesPtrInput
}

func (TaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*taskArgs)(nil)).Elem()
}

type TaskInput interface {
	pulumi.Input

	ToTaskOutput() TaskOutput
	ToTaskOutputWithContext(ctx context.Context) TaskOutput
}

func (*Task) ElementType() reflect.Type {
	return reflect.TypeOf((*Task)(nil))
}

func (i *Task) ToTaskOutput() TaskOutput {
	return i.ToTaskOutputWithContext(context.Background())
}

func (i *Task) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOutput)
}

type TaskOutput struct{ *pulumi.OutputState }

func (TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Task)(nil))
}

func (o TaskOutput) ToTaskOutput() TaskOutput {
	return o
}

func (o TaskOutput) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*TaskInput)(nil)).Elem(), &Task{})
	pulumi.RegisterOutputType(TaskOutput{})
}
