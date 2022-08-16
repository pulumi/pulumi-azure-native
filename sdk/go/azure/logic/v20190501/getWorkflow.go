


package v20190501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkflow(ctx *pulumi.Context, args *LookupWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowResult, error) {
	var rv LookupWorkflowResult
	err := ctx.Invoke("azure-native:logic/v20190501:getWorkflow", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWorkflowArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkflowName      string `pulumi:"workflowName"`
}


type LookupWorkflowResult struct {
	AccessControl                 *FlowAccessControlConfigurationResponse `pulumi:"accessControl"`
	AccessEndpoint                string                                  `pulumi:"accessEndpoint"`
	ChangedTime                   string                                  `pulumi:"changedTime"`
	CreatedTime                   string                                  `pulumi:"createdTime"`
	Definition                    interface{}                             `pulumi:"definition"`
	EndpointsConfiguration        *FlowEndpointsConfigurationResponse     `pulumi:"endpointsConfiguration"`
	Id                            string                                  `pulumi:"id"`
	Identity                      *ManagedServiceIdentityResponse         `pulumi:"identity"`
	IntegrationAccount            *ResourceReferenceResponse              `pulumi:"integrationAccount"`
	IntegrationServiceEnvironment *ResourceReferenceResponse              `pulumi:"integrationServiceEnvironment"`
	Location                      *string                                 `pulumi:"location"`
	Name                          string                                  `pulumi:"name"`
	Parameters                    map[string]WorkflowParameterResponse    `pulumi:"parameters"`
	ProvisioningState             string                                  `pulumi:"provisioningState"`
	Sku                           SkuResponse                             `pulumi:"sku"`
	State                         *string                                 `pulumi:"state"`
	Tags                          map[string]string                       `pulumi:"tags"`
	Type                          string                                  `pulumi:"type"`
	Version                       string                                  `pulumi:"version"`
}

func LookupWorkflowOutput(ctx *pulumi.Context, args LookupWorkflowOutputArgs, opts ...pulumi.InvokeOption) LookupWorkflowResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWorkflowResult, error) {
			args := v.(LookupWorkflowArgs)
			r, err := LookupWorkflow(ctx, &args, opts...)
			var s LookupWorkflowResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWorkflowResultOutput)
}

type LookupWorkflowOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkflowName      pulumi.StringInput `pulumi:"workflowName"`
}

func (LookupWorkflowOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowArgs)(nil)).Elem()
}


type LookupWorkflowResultOutput struct{ *pulumi.OutputState }

func (LookupWorkflowResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWorkflowResult)(nil)).Elem()
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutput() LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) ToLookupWorkflowResultOutputWithContext(ctx context.Context) LookupWorkflowResultOutput {
	return o
}

func (o LookupWorkflowResultOutput) AccessControl() FlowAccessControlConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *FlowAccessControlConfigurationResponse { return v.AccessControl }).(FlowAccessControlConfigurationResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) AccessEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.AccessEndpoint }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) ChangedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.ChangedTime }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) CreatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.CreatedTime }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Definition() pulumi.AnyOutput {
	return o.ApplyT(func(v LookupWorkflowResult) interface{} { return v.Definition }).(pulumi.AnyOutput)
}

func (o LookupWorkflowResultOutput) EndpointsConfiguration() FlowEndpointsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *FlowEndpointsConfigurationResponse { return v.EndpointsConfiguration }).(FlowEndpointsConfigurationResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ManagedServiceIdentityResponse { return v.Identity }).(ManagedServiceIdentityResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) IntegrationAccount() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ResourceReferenceResponse { return v.IntegrationAccount }).(ResourceReferenceResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) IntegrationServiceEnvironment() ResourceReferenceResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ResourceReferenceResponse { return v.IntegrationServiceEnvironment }).(ResourceReferenceResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Parameters() WorkflowParameterResponseMapOutput {
	return o.ApplyT(func(v LookupWorkflowResult) map[string]WorkflowParameterResponse { return v.Parameters }).(WorkflowParameterResponseMapOutput)
}

func (o LookupWorkflowResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v LookupWorkflowResult) SkuResponse { return v.Sku }).(SkuResponseOutput)
}

func (o LookupWorkflowResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkflowResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkflowResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowResultOutput{})
}
