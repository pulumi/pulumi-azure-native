


package v20180301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMachineLearningCompute(ctx *pulumi.Context, args *LookupMachineLearningComputeArgs, opts ...pulumi.InvokeOption) (*LookupMachineLearningComputeResult, error) {
	var rv LookupMachineLearningComputeResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20180301preview:getMachineLearningCompute", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMachineLearningComputeArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMachineLearningComputeResult struct {
	Id         string            `pulumi:"id"`
	Identity   *IdentityResponse `pulumi:"identity"`
	Location   *string           `pulumi:"location"`
	Name       string            `pulumi:"name"`
	Properties interface{}       `pulumi:"properties"`
	Tags       map[string]string `pulumi:"tags"`
	Type       string            `pulumi:"type"`
}
