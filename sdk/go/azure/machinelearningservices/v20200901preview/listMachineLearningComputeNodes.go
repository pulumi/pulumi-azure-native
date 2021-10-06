


package v20200901preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMachineLearningComputeNodes(ctx *pulumi.Context, args *ListMachineLearningComputeNodesArgs, opts ...pulumi.InvokeOption) (*ListMachineLearningComputeNodesResult, error) {
	var rv ListMachineLearningComputeNodesResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20200901preview:listMachineLearningComputeNodes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMachineLearningComputeNodesArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListMachineLearningComputeNodesResult struct {
	ComputeType string                              `pulumi:"computeType"`
	NextLink    string                              `pulumi:"nextLink"`
	Nodes       []AmlComputeNodeInformationResponse `pulumi:"nodes"`
}
