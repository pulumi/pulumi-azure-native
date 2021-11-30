


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workflow struct {
	pulumi.CustomResourceState

	AccessEndpoint     pulumi.StringOutput                `pulumi:"accessEndpoint"`
	ChangedTime        pulumi.StringOutput                `pulumi:"changedTime"`
	CreatedTime        pulumi.StringOutput                `pulumi:"createdTime"`
	Definition         pulumi.AnyOutput                   `pulumi:"definition"`
	IntegrationAccount ResourceReferenceResponsePtrOutput `pulumi:"integrationAccount"`
	Location           pulumi.StringPtrOutput             `pulumi:"location"`
	Name               pulumi.StringOutput                `pulumi:"name"`
	Parameters         WorkflowParameterResponseMapOutput `pulumi:"parameters"`
	ProvisioningState  pulumi.StringOutput                `pulumi:"provisioningState"`
	Sku                SkuResponsePtrOutput               `pulumi:"sku"`
	State              pulumi.StringPtrOutput             `pulumi:"state"`
	Tags               pulumi.StringMapOutput             `pulumi:"tags"`
	Type               pulumi.StringOutput                `pulumi:"type"`
	Version            pulumi.StringOutput                `pulumi:"version"`
}


func NewWorkflow(ctx *pulumi.Context,
	name string, args *WorkflowArgs, opts ...pulumi.ResourceOption) (*Workflow, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logic:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:logic/v20180701preview:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:logic/v20180701preview:Workflow", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type workflowState struct {
}

type WorkflowState struct {
}

func (WorkflowState) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowState)(nil)).Elem()
}

type workflowArgs struct {
	Definition         interface{}                  `pulumi:"definition"`
	IntegrationAccount *ResourceReference           `pulumi:"integrationAccount"`
	Location           *string                      `pulumi:"location"`
	Parameters         map[string]WorkflowParameter `pulumi:"parameters"`
	ResourceGroupName  string                       `pulumi:"resourceGroupName"`
	Sku                *Sku                         `pulumi:"sku"`
	State              *string                      `pulumi:"state"`
	Tags               map[string]string            `pulumi:"tags"`
	WorkflowName       *string                      `pulumi:"workflowName"`
}


type WorkflowArgs struct {
	Definition         pulumi.Input
	IntegrationAccount ResourceReferencePtrInput
	Location           pulumi.StringPtrInput
	Parameters         WorkflowParameterMapInput
	ResourceGroupName  pulumi.StringInput
	Sku                SkuPtrInput
	State              pulumi.StringPtrInput
	Tags               pulumi.StringMapInput
	WorkflowName       pulumi.StringPtrInput
}

func (WorkflowArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*workflowArgs)(nil)).Elem()
}

type WorkflowInput interface {
	pulumi.Input

	ToWorkflowOutput() WorkflowOutput
	ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput
}

func (*Workflow) ElementType() reflect.Type {
	return reflect.TypeOf((*Workflow)(nil))
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Workflow)(nil))
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}
