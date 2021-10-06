


package v20190601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAgentPool(ctx *pulumi.Context, args *LookupAgentPoolArgs, opts ...pulumi.InvokeOption) (*LookupAgentPoolResult, error) {
	var rv LookupAgentPoolResult
	err := ctx.Invoke("azure-native:containerregistry/v20190601preview:getAgentPool", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAgentPoolArgs struct {
	AgentPoolName     string `pulumi:"agentPoolName"`
	RegistryName      string `pulumi:"registryName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}



type LookupAgentPoolResult struct {
	Count                          *int               `pulumi:"count"`
	Id                             string             `pulumi:"id"`
	Location                       string             `pulumi:"location"`
	Name                           string             `pulumi:"name"`
	Os                             *string            `pulumi:"os"`
	ProvisioningState              string             `pulumi:"provisioningState"`
	SystemData                     SystemDataResponse `pulumi:"systemData"`
	Tags                           map[string]string  `pulumi:"tags"`
	Tier                           *string            `pulumi:"tier"`
	Type                           string             `pulumi:"type"`
	VirtualNetworkSubnetResourceId *string            `pulumi:"virtualNetworkSubnetResourceId"`
}
