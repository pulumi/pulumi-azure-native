


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListComputeKeys(ctx *pulumi.Context, args *ListComputeKeysArgs, opts ...pulumi.InvokeOption) (*ListComputeKeysResult, error) {
	var rv ListComputeKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listComputeKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListComputeKeysArgs struct {
	ComputeName       string `pulumi:"computeName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListComputeKeysResult struct {
	ComputeType string `pulumi:"computeType"`
}
