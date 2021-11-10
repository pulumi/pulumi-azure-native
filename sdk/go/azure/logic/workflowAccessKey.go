


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WorkflowAccessKey struct {
	pulumi.CustomResourceState

	Name      pulumi.StringOutput    `pulumi:"name"`
	NotAfter  pulumi.StringPtrOutput `pulumi:"notAfter"`
	NotBefore pulumi.StringPtrOutput `pulumi:"notBefore"`
	Type      pulumi.StringOutput    `pulumi:"type"`
}


func NewWorkflowAccessKey(ctx *pulumi.Context,
	name string, args *WorkflowAccessKeyArgs, opts ...pulumi.ResourceOption) (*WorkflowAccessKey, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkflowName == nil {
		return nil, errors.New("invalid value for required argument 'WorkflowName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic/v20150201preview:WorkflowAccessKey"),
		},
	})
	opts = append(opts, aliases)
	var resource WorkflowAccessKey
	err := ctx.RegisterResource("azure-native:logic:WorkflowAccessKey", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkflowAccessKey(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowAccessKeyState, opts ...pulumi.ResourceOption) (*WorkflowAccessKey, error) {
	var resource WorkflowAccessKey
	err := ctx.ReadResource("azure-native:logic:WorkflowAccessKey", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workflowAccessKeyState struct {
}

type WorkflowAccessKeyState struct {
}

func (WorkflowAccessKeyState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowAccessKeyState)(nil)).Elem()
}

type workflowAccessKeyArgs struct {
	AccessKeyName     *string `pulumi:"accessKeyName"`
	Id                *string `pulumi:"id"`
	NotAfter          *string `pulumi:"notAfter"`
	NotBefore         *string `pulumi:"notBefore"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
	WorkflowName      string  `pulumi:"workflowName"`
}


type WorkflowAccessKeyArgs struct {
	AccessKeyName     pulumi.StringPtrInput
	Id                pulumi.StringPtrInput
	NotAfter          pulumi.StringPtrInput
	NotBefore         pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	WorkflowName      pulumi.StringInput
}

func (WorkflowAccessKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowAccessKeyArgs)(nil)).Elem()
}

type WorkflowAccessKeyInput interface {
	pulumi.Input

	ToWorkflowAccessKeyOutput() WorkflowAccessKeyOutput
	ToWorkflowAccessKeyOutputWithContext(ctx context.Context) WorkflowAccessKeyOutput
}

func (*WorkflowAccessKey) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowAccessKey)(nil))
}

func (i *WorkflowAccessKey) ToWorkflowAccessKeyOutput() WorkflowAccessKeyOutput {
	return i.ToWorkflowAccessKeyOutputWithContext(context.Background())
}

func (i *WorkflowAccessKey) ToWorkflowAccessKeyOutputWithContext(ctx context.Context) WorkflowAccessKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowAccessKeyOutput)
}

type WorkflowAccessKeyOutput struct{ *pulumi.OutputState }

func (WorkflowAccessKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowAccessKey)(nil))
}

func (o WorkflowAccessKeyOutput) ToWorkflowAccessKeyOutput() WorkflowAccessKeyOutput {
	return o
}

func (o WorkflowAccessKeyOutput) ToWorkflowAccessKeyOutputWithContext(ctx context.Context) WorkflowAccessKeyOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkflowAccessKeyOutput{})
}
