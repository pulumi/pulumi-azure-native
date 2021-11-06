


package v20191101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMachineLearningComputeKeys(ctx *pulumi.Context, args *ListMachineLearningComputeKeysArgs, opts ...pulumi.InvokeOption) (*ListMachineLearningComputeKeysResult, error) {
	var rv ListMachineLearningComputeKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20191101:listMachineLearningComputeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMachineLearningComputeKeysArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListMachineLearningComputeKeysResult struct {
	ComputeType string `pulumi:"computeType"`
}
