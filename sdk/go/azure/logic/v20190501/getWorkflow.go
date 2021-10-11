


package v20190501

import (
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
