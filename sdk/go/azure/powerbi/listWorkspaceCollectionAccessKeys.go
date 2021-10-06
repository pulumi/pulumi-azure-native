


package powerbi

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceCollectionAccessKeys(ctx *pulumi.Context, args *ListWorkspaceCollectionAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceCollectionAccessKeysResult, error) {
	var rv ListWorkspaceCollectionAccessKeysResult
	err := ctx.Invoke("azure-native:powerbi:listWorkspaceCollectionAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceCollectionAccessKeysArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type ListWorkspaceCollectionAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}
