


package v20150201preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workflow struct {
	pulumi.CustomResourceState

	AccessEndpoint    pulumi.StringOutput                `pulumi:"accessEndpoint"`
	ChangedTime       pulumi.StringOutput                `pulumi:"changedTime"`
	CreatedTime       pulumi.StringOutput                `pulumi:"createdTime"`
	Definition        pulumi.AnyOutput                   `pulumi:"definition"`
	DefinitionLink    ContentLinkResponsePtrOutput       `pulumi:"definitionLink"`
	Location          pulumi.StringPtrOutput             `pulumi:"location"`
	Name              pulumi.StringPtrOutput             `pulumi:"name"`
	Parameters        WorkflowParameterResponseMapOutput `pulumi:"parameters"`
	ParametersLink    ContentLinkResponsePtrOutput       `pulumi:"parametersLink"`
	ProvisioningState pulumi.StringOutput                `pulumi:"provisioningState"`
	Sku               SkuResponsePtrOutput               `pulumi:"sku"`
	State             pulumi.StringPtrOutput             `pulumi:"state"`
	Tags              pulumi.StringMapOutput             `pulumi:"tags"`
	Type              pulumi.StringPtrOutput             `pulumi:"type"`
	Version           pulumi.StringOutput                `pulumi:"version"`
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
			Type: pulumi.String("azure-nextgen:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Workflow"),
		},
		{
			Type: pulumi.String("azure-nextgen:logic/v20190501:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:logic/v20150201preview:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:logic/v20150201preview:Workflow", name, id, state, &resource, opts...)
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
	Definition        interface{}                  `pulumi:"definition"`
	DefinitionLink    *ContentLink                 `pulumi:"definitionLink"`
	Id                *string                      `pulumi:"id"`
	Location          *string                      `pulumi:"location"`
	Name              *string                      `pulumi:"name"`
	Parameters        map[string]WorkflowParameter `pulumi:"parameters"`
	ParametersLink    *ContentLink                 `pulumi:"parametersLink"`
	ResourceGroupName string                       `pulumi:"resourceGroupName"`
	Sku               *Sku                         `pulumi:"sku"`
	State             *WorkflowStateEnum           `pulumi:"state"`
	Tags              map[string]string            `pulumi:"tags"`
	Type              *string                      `pulumi:"type"`
	WorkflowName      *string                      `pulumi:"workflowName"`
}


type WorkflowArgs struct {
	Definition        pulumi.Input
	DefinitionLink    ContentLinkPtrInput
	Id                pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringPtrInput
	Parameters        WorkflowParameterMapInput
	ParametersLink    ContentLinkPtrInput
	ResourceGroupName pulumi.StringInput
	Sku               SkuPtrInput
	State             WorkflowStateEnumPtrInput
	Tags              pulumi.StringMapInput
	Type              pulumi.StringPtrInput
	WorkflowName      pulumi.StringPtrInput
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
