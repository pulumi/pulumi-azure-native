


package synapse

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetIntegrationRuntimeStatus(ctx *pulumi.Context, args *GetIntegrationRuntimeStatusArgs, opts ...pulumi.InvokeOption) (*GetIntegrationRuntimeStatusResult, error) {
	var rv GetIntegrationRuntimeStatusResult
	err := ctx.Invoke("azure-native:synapse:getIntegrationRuntimeStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetIntegrationRuntimeStatusArgs struct {
	IntegrationRuntimeName string `pulumi:"integrationRuntimeName"`
	ResourceGroupName      string `pulumi:"resourceGroupName"`
	WorkspaceName          string `pulumi:"workspaceName"`
}


type GetIntegrationRuntimeStatusResult struct {
	Name       string      `pulumi:"name"`
	Properties interface{} `pulumi:"properties"`
}
