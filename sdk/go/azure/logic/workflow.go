


package logic

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type Workflow struct {
	pulumi.CustomResourceState

	AccessControl                 FlowAccessControlConfigurationResponsePtrOutput `pulumi:"accessControl"`
	AccessEndpoint                pulumi.StringOutput                             `pulumi:"accessEndpoint"`
	ChangedTime                   pulumi.StringOutput                             `pulumi:"changedTime"`
	CreatedTime                   pulumi.StringOutput                             `pulumi:"createdTime"`
	Definition                    pulumi.AnyOutput                                `pulumi:"definition"`
	EndpointsConfiguration        FlowEndpointsConfigurationResponsePtrOutput     `pulumi:"endpointsConfiguration"`
	Identity                      ManagedServiceIdentityResponsePtrOutput         `pulumi:"identity"`
	IntegrationAccount            ResourceReferenceResponsePtrOutput              `pulumi:"integrationAccount"`
	IntegrationServiceEnvironment ResourceReferenceResponsePtrOutput              `pulumi:"integrationServiceEnvironment"`
	Location                      pulumi.StringPtrOutput                          `pulumi:"location"`
	Name                          pulumi.StringOutput                             `pulumi:"name"`
	Parameters                    WorkflowParameterResponseMapOutput              `pulumi:"parameters"`
	ProvisioningState             pulumi.StringOutput                             `pulumi:"provisioningState"`
	Sku                           SkuResponseOutput                               `pulumi:"sku"`
	State                         pulumi.StringPtrOutput                          `pulumi:"state"`
	Tags                          pulumi.StringMapOutput                          `pulumi:"tags"`
	Type                          pulumi.StringOutput                             `pulumi:"type"`
	Version                       pulumi.StringOutput                             `pulumi:"version"`
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
			Type: pulumi.String("azure-native:logic/v20150201preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20160601:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20180701preview:Workflow"),
		},
		{
			Type: pulumi.String("azure-native:logic/v20190501:Workflow"),
		},
	})
	opts = append(opts, aliases)
	var resource Workflow
	err := ctx.RegisterResource("azure-native:logic:Workflow", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWorkflow(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WorkflowState, opts ...pulumi.ResourceOption) (*Workflow, error) {
	var resource Workflow
	err := ctx.ReadResource("azure-native:logic:Workflow", name, id, state, &resource, opts...)
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
	AccessControl                 *FlowAccessControlConfiguration `pulumi:"accessControl"`
	Definition                    interface{}                     `pulumi:"definition"`
	EndpointsConfiguration        *FlowEndpointsConfiguration     `pulumi:"endpointsConfiguration"`
	Identity                      *ManagedServiceIdentity         `pulumi:"identity"`
	IntegrationAccount            *ResourceReference              `pulumi:"integrationAccount"`
	IntegrationServiceEnvironment *ResourceReference              `pulumi:"integrationServiceEnvironment"`
	Location                      *string                         `pulumi:"location"`
	Parameters                    map[string]WorkflowParameter    `pulumi:"parameters"`
	ResourceGroupName             string                          `pulumi:"resourceGroupName"`
	State                         *string                         `pulumi:"state"`
	Tags                          map[string]string               `pulumi:"tags"`
	WorkflowName                  *string                         `pulumi:"workflowName"`
}


type WorkflowArgs struct {
	AccessControl                 FlowAccessControlConfigurationPtrInput
	Definition                    pulumi.Input
	EndpointsConfiguration        FlowEndpointsConfigurationPtrInput
	Identity                      ManagedServiceIdentityPtrInput
	IntegrationAccount            ResourceReferencePtrInput
	IntegrationServiceEnvironment ResourceReferencePtrInput
	Location                      pulumi.StringPtrInput
	Parameters                    WorkflowParameterMapInput
	ResourceGroupName             pulumi.StringInput
	State                         pulumi.StringPtrInput
	Tags                          pulumi.StringMapInput
	WorkflowName                  pulumi.StringPtrInput
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
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (i *Workflow) ToWorkflowOutput() WorkflowOutput {
	return i.ToWorkflowOutputWithContext(context.Background())
}

func (i *Workflow) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkflowOutput)
}

type WorkflowOutput struct{ *pulumi.OutputState }

func (WorkflowOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Workflow)(nil)).Elem()
}

func (o WorkflowOutput) ToWorkflowOutput() WorkflowOutput {
	return o
}

func (o WorkflowOutput) ToWorkflowOutputWithContext(ctx context.Context) WorkflowOutput {
	return o
}

func (o WorkflowOutput) AccessControl() FlowAccessControlConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) FlowAccessControlConfigurationResponsePtrOutput { return v.AccessControl }).(FlowAccessControlConfigurationResponsePtrOutput)
}

func (o WorkflowOutput) AccessEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.AccessEndpoint }).(pulumi.StringOutput)
}

func (o WorkflowOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o WorkflowOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Definition() pulumi.AnyOutput {
	return o.ApplyT(func(v *Workflow) pulumi.AnyOutput { return v.Definition }).(pulumi.AnyOutput)
}

func (o WorkflowOutput) EndpointsConfiguration() FlowEndpointsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) FlowEndpointsConfigurationResponsePtrOutput { return v.EndpointsConfiguration }).(FlowEndpointsConfigurationResponsePtrOutput)
}

func (o WorkflowOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ManagedServiceIdentityResponsePtrOutput { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o WorkflowOutput) IntegrationAccount() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ResourceReferenceResponsePtrOutput { return v.IntegrationAccount }).(ResourceReferenceResponsePtrOutput)
}

func (o WorkflowOutput) IntegrationServiceEnvironment() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v *Workflow) ResourceReferenceResponsePtrOutput { return v.IntegrationServiceEnvironment }).(ResourceReferenceResponsePtrOutput)
}

func (o WorkflowOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.Location }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Parameters() WorkflowParameterResponseMapOutput {
	return o.ApplyT(func(v *Workflow) WorkflowParameterResponseMapOutput { return v.Parameters }).(WorkflowParameterResponseMapOutput)
}

func (o WorkflowOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *Workflow) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o WorkflowOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringPtrOutput { return v.State }).(pulumi.StringPtrOutput)
}

func (o WorkflowOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o WorkflowOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WorkflowOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v *Workflow) pulumi.StringOutput { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(WorkflowOutput{})
}
