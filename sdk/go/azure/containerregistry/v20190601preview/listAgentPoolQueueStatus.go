


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListAgentPoolQueueStatus(ctx *pulumi.Context, args *ListAgentPoolQueueStatusArgs, opts ...pulumi.InvokeOption) (*ListAgentPoolQueueStatusResult, error) {
	var rv ListAgentPoolQueueStatusResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:listAgentPoolQueueStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListAgentPoolQueueStatusArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListAgentPoolQueueStatusResult struct {
	Count *int `pulumi:"count"`
}
