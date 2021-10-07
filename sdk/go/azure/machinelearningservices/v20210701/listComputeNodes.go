


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListComputeNodes(ctx *pulumi.Context, args *ListComputeNodesArgs, opts ...pulumi.InvokeOption) (*ListComputeNodesResult, error) {
	var rv ListComputeNodesResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210701:listComputeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListComputeNodesArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListComputeNodesResult struct {
	NextLink string                              `pulumi:"nextLink"`
	Nodes    []AmlComputeNodeInformationResponse `pulumi:"nodes"`
}
