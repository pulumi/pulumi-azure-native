


package v20180901

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Task struct {
	pulumi.CustomResourceState

	AgentConfiguration AgentPropertiesResponsePtrOutput   `pulumi:"agentConfiguration"`
	CreationDate       pulumi.StringOutput                `pulumi:"creationDate"`
	Credentials        CredentialsResponsePtrOutput       `pulumi:"credentials"`
	Location           pulumi.StringOutput                `pulumi:"location"`
	Name               pulumi.StringOutput                `pulumi:"name"`
	Platform           PlatformPropertiesResponseOutput   `pulumi:"platform"`
	ProvisioningState  pulumi.StringOutput                `pulumi:"provisioningState"`
	Status             pulumi.StringPtrOutput             `pulumi:"status"`
	Step               pulumi.AnyOutput                   `pulumi:"step"`
	Tags               pulumi.StringMapOutput             `pulumi:"tags"`
	Timeout            pulumi.IntPtrOutput                `pulumi:"timeout"`
	Trigger            TriggerPropertiesResponsePtrOutput `pulumi:"trigger"`
	Type               pulumi.StringOutput                `pulumi:"type"`
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
	if isZero(args.Timeout) {
		args.Timeout = pulumi.IntPtr(3600)
	}
	if args.Trigger != nil {
		args.Trigger = args.Trigger.ToTriggerPropertiesPtrOutput().ApplyT(func(v *TriggerProperties) *TriggerProperties { return v.Defaults() }).(TriggerPropertiesPtrOutput)
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:containerregistry:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20180201preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190401:Task"),
		},
		{
			Type: pulumi.String("azure-native:containerregistry/v20190601preview:Task"),
		},
	})
	opts = append(opts, aliases)
	var resource Task
	err := ctx.RegisterResource("azure-native:containerregistry/v20180901:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("azure-native:containerregistry/v20180901:Task", name, id, state, &resource, opts...)
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
	AgentConfiguration *AgentProperties   `pulumi:"agentConfiguration"`
	Credentials        *Credentials       `pulumi:"credentials"`
	Location           *string            `pulumi:"location"`
	Platform           PlatformProperties `pulumi:"platform"`
	RegistryName       string             `pulumi:"registryName"`
	ResourceGroupName  string             `pulumi:"resourceGroupName"`
	Status             *string            `pulumi:"status"`
	Step               TaskStepProperties `pulumi:"step"`
	Tags               map[string]string  `pulumi:"tags"`
	TaskName           *string            `pulumi:"taskName"`
	Timeout            *int               `pulumi:"timeout"`
	Trigger            *TriggerProperties `pulumi:"trigger"`
}


type TaskArgs struct {
	AgentConfiguration AgentPropertiesPtrInput
	Credentials        CredentialsPtrInput
	Location           pulumi.StringPtrInput
	Platform           PlatformPropertiesInput
	RegistryName       pulumi.StringInput
	ResourceGroupName  pulumi.StringInput
	Status             pulumi.StringPtrInput
	Step               TaskStepPropertiesInput
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
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (i *Task) ToTaskOutput() TaskOutput {
	return i.ToTaskOutputWithContext(context.Background())
}

func (i *Task) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TaskOutput)
}

type TaskOutput struct{ *pulumi.OutputState }

func (TaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Task)(nil)).Elem()
}

func (o TaskOutput) ToTaskOutput() TaskOutput {
	return o
}

func (o TaskOutput) ToTaskOutputWithContext(ctx context.Context) TaskOutput {
	return o
}

func (o TaskOutput) AgentConfiguration() AgentPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Task) AgentPropertiesResponsePtrOutput { return v.AgentConfiguration }).(AgentPropertiesResponsePtrOutput)
}

func (o TaskOutput) CreationDate() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.CreationDate }).(pulumi.StringOutput)
}

func (o TaskOutput) Credentials() CredentialsResponsePtrOutput {
	return o.ApplyT(func(v *Task) CredentialsResponsePtrOutput { return v.Credentials }).(CredentialsResponsePtrOutput)
}

func (o TaskOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o TaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o TaskOutput) Platform() PlatformPropertiesResponseOutput {
	return o.ApplyT(func(v *Task) PlatformPropertiesResponseOutput { return v.Platform }).(PlatformPropertiesResponseOutput)
}

func (o TaskOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o TaskOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Task) pulumi.StringPtrOutput { return v.Status }).(pulumi.StringPtrOutput)
}

func (o TaskOutput) Step() pulumi.AnyOutput {
	return o.ApplyT(func(v *Task) pulumi.AnyOutput { return v.Step }).(pulumi.AnyOutput)
}

func (o TaskOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Task) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o TaskOutput) Timeout() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Task) pulumi.IntPtrOutput { return v.Timeout }).(pulumi.IntPtrOutput)
}

func (o TaskOutput) Trigger() TriggerPropertiesResponsePtrOutput {
	return o.ApplyT(func(v *Task) TriggerPropertiesResponsePtrOutput { return v.Trigger }).(TriggerPropertiesResponsePtrOutput)
}

func (o TaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Task) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(TaskOutput{})
}
