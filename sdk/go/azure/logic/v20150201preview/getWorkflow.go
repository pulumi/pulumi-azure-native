


package v20150201preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWorkflow(ctx *pulumi.Context, args *LookupWorkflowArgs, opts ...pulumi.InvokeOption) (*LookupWorkflowResult, error) {
	var rv LookupWorkflowResult
	err := ctx.Invoke("azure-native:logic/v20150201preview:getWorkflow", args, &rv, opts...)
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
	AccessEndpoint    string                               `pulumi:"accessEndpoint"`
	ChangedTime       string                               `pulumi:"changedTime"`
	CreatedTime       string                               `pulumi:"createdTime"`
	Definition        interface{}                          `pulumi:"definition"`
	DefinitionLink    *ContentLinkResponse                 `pulumi:"definitionLink"`
	Id                *string                              `pulumi:"id"`
	Location          *string                              `pulumi:"location"`
	Name              *string                              `pulumi:"name"`
	Parameters        map[string]WorkflowParameterResponse `pulumi:"parameters"`
	ParametersLink    *ContentLinkResponse                 `pulumi:"parametersLink"`
	ProvisioningState string                               `pulumi:"provisioningState"`
	Sku               *SkuResponse                         `pulumi:"sku"`
	State             *string                              `pulumi:"state"`
	Tags              map[string]string                    `pulumi:"tags"`
	Type              *string                              `pulumi:"type"`
	Version           string                               `pulumi:"version"`
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

func (o LookupWorkflowResultOutput) DefinitionLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ContentLinkResponse { return v.DefinitionLink }).(ContentLinkResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Parameters() WorkflowParameterResponseMapOutput {
	return o.ApplyT(func(v LookupWorkflowResult) map[string]WorkflowParameterResponse { return v.Parameters }).(WorkflowParameterResponseMapOutput)
}

func (o LookupWorkflowResultOutput) ParametersLink() ContentLinkResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *ContentLinkResponse { return v.ParametersLink }).(ContentLinkResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWorkflowResultOutput) Sku() SkuResponsePtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *SkuResponse { return v.Sku }).(SkuResponsePtrOutput)
}

func (o LookupWorkflowResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWorkflowResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWorkflowResultOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWorkflowResult) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o LookupWorkflowResultOutput) Version() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWorkflowResult) string { return v.Version }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWorkflowResultOutput{})
}
