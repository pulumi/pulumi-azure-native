


package v20210315

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotebookWorkspace(ctx *pulumi.Context, args *LookupNotebookWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupNotebookWorkspaceResult, error) {
	var rv LookupNotebookWorkspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20210315:getNotebookWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotebookWorkspaceArgs struct {
	AccountName           string `pulumi:"accountName"`
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNotebookWorkspaceResult struct {
	Id                     string `pulumi:"id"`
	Name                   string `pulumi:"name"`
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
	Status                 string `pulumi:"status"`
	Type                   string `pulumi:"type"`
}
