


package v20210301preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceStorageAccountKeys(ctx *pulumi.Context, args *ListWorkspaceStorageAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceStorageAccountKeysResult, error) {
	var rv ListWorkspaceStorageAccountKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210301preview:listWorkspaceStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceStorageAccountKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceStorageAccountKeysResult struct {
	UserStorageKey string `pulumi:"userStorageKey"`
}
