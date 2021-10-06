


package v20200601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNotebookWorkspaceConnectionInfo(ctx *pulumi.Context, args *ListNotebookWorkspaceConnectionInfoArgs, opts ...pulumi.InvokeOption) (*ListNotebookWorkspaceConnectionInfoResult, error) {
	var rv ListNotebookWorkspaceConnectionInfoResult
	err := ctx.Invoke("azure-native:documentdb/v20200601preview:listNotebookWorkspaceConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookWorkspaceConnectionInfoArgs struct {
	AccountName           string `pulumi:"accountName"`
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListNotebookWorkspaceConnectionInfoResult struct {
	AuthToken              string `pulumi:"authToken"`
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
}
