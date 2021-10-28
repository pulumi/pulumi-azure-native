


package v20210630

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Task struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput   `pulumi:"etag"`
	Name       pulumi.StringOutput      `pulumi:"name"`
	Properties pulumi.AnyOutput         `pulumi:"properties"`
	SystemData SystemDataResponseOutput `pulumi:"systemData"`
	Type       pulumi.StringOutput      `pulumi:"type"`
}


func NewTask(ctx *pulumi.Context,
	name string, args *TaskArgs, opts ...pulumi.ResourceOption) (*Task, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ProjectName == nil {
		return nil, errors.New("invalid value for required argument 'ProjectName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20210630:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20171115preview:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20171115preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180315preview:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180315preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180331preview:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180331preview:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180419:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180419:Task"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20180715preview:Task"),
		},
		{
			Type: pulumi.String("azure-nextgen:datamigration/v20180715preview:Task"),
		},
	})
	opts = append(opts, aliases)
	var resource Task
	err := ctx.RegisterResource("azure-native:datamigration/v20210630:Task", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *TaskState, opts ...pulumi.ResourceOption) (*Task, error) {
	var resource Task
	err := ctx.ReadResource("azure-native:datamigration/v20210630:Task", name, id, state, &resource, opts...)
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
	Etag        *string     `pulumi:"etag"`
	GroupName   string      `pulumi:"groupName"`
	ProjectName string      `pulumi:"projectName"`
	Properties  interface{} `pulumi:"properties"`
	ServiceName string      `pulumi:"serviceName"`
	TaskName    *string     `pulumi:"taskName"`
}


type TaskArgs struct {
	Etag        pulumi.StringPtrInput
	GroupName   pulumi.StringInput
	ProjectName pulumi.StringInput
	Properties  pulumi.Input
	ServiceName pulumi.StringInput
	TaskName    pulumi.StringPtrInput
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
