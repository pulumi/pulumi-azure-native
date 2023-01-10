


package v20180715preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ServiceTask struct {
	pulumi.CustomResourceState

	Etag       pulumi.StringPtrOutput `pulumi:"etag"`
	Name       pulumi.StringOutput    `pulumi:"name"`
	Properties pulumi.AnyOutput       `pulumi:"properties"`
	Type       pulumi.StringOutput    `pulumi:"type"`
}


func NewServiceTask(ctx *pulumi.Context,
	name string, args *ServiceTaskArgs, opts ...pulumi.ResourceOption) (*ServiceTask, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.GroupName == nil {
		return nil, errors.New("invalid value for required argument 'GroupName'")
	}
	if args.ServiceName == nil {
		return nil, errors.New("invalid value for required argument 'ServiceName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:datamigration/v20210630:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20211030preview:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220130preview:ServiceTask"),
		},
		{
			Type: pulumi.String("azure-native:datamigration/v20220330preview:ServiceTask"),
		},
	})
	opts = append(opts, aliases)
	var resource ServiceTask
	err := ctx.RegisterResource("azure-native:datamigration/v20180715preview:ServiceTask", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetServiceTask(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ServiceTaskState, opts ...pulumi.ResourceOption) (*ServiceTask, error) {
	var resource ServiceTask
	err := ctx.ReadResource("azure-native:datamigration/v20180715preview:ServiceTask", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type serviceTaskState struct {
}

type ServiceTaskState struct {
}

func (ServiceTaskState) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTaskState)(nil)).Elem()
}

type serviceTaskArgs struct {
	GroupName   string      `pulumi:"groupName"`
	Properties  interface{} `pulumi:"properties"`
	ServiceName string      `pulumi:"serviceName"`
	TaskName    *string     `pulumi:"taskName"`
}


type ServiceTaskArgs struct {
	GroupName   pulumi.StringInput
	Properties  pulumi.Input
	ServiceName pulumi.StringInput
	TaskName    pulumi.StringPtrInput
}

func (ServiceTaskArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*serviceTaskArgs)(nil)).Elem()
}

type ServiceTaskInput interface {
	pulumi.Input

	ToServiceTaskOutput() ServiceTaskOutput
	ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput
}

func (*ServiceTask) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTask)(nil)).Elem()
}

func (i *ServiceTask) ToServiceTaskOutput() ServiceTaskOutput {
	return i.ToServiceTaskOutputWithContext(context.Background())
}

func (i *ServiceTask) ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceTaskOutput)
}

type ServiceTaskOutput struct{ *pulumi.OutputState }

func (ServiceTaskOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceTask)(nil)).Elem()
}

func (o ServiceTaskOutput) ToServiceTaskOutput() ServiceTaskOutput {
	return o
}

func (o ServiceTaskOutput) ToServiceTaskOutputWithContext(ctx context.Context) ServiceTaskOutput {
	return o
}

func (o ServiceTaskOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ServiceTaskOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceTaskOutput) Properties() pulumi.AnyOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.AnyOutput { return v.Properties }).(pulumi.AnyOutput)
}

func (o ServiceTaskOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ServiceTask) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ServiceTaskOutput{})
}
