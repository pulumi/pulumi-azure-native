


package v20210401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNotebookKeys(ctx *pulumi.Context, args *ListNotebookKeysArgs, opts ...pulumi.InvokeOption) (*ListNotebookKeysResult, error) {
	var rv ListNotebookKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20210401:listNotebookKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListNotebookKeysResult struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}
