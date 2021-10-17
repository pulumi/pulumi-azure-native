


package v20210701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceNotebookKeys(ctx *pulumi.Context, args *ListWorkspaceNotebookKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceNotebookKeysResult, error) {
	var rv ListWorkspaceNotebookKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210701:listWorkspaceNotebookKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceNotebookKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceNotebookKeysResult struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}
