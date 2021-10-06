


package v20150201preview

import (
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
