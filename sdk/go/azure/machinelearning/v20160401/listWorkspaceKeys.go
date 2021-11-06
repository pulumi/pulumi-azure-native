


package v20160401

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceKeys(ctx *pulumi.Context, args *ListWorkspaceKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceKeysResult, error) {
	var rv ListWorkspaceKeysResult
	err := ctx.Invoke("azure-native:machinelearning/v20160401:listWorkspaceKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListWorkspaceKeysResult struct {
	PrimaryToken   *string `pulumi:"primaryToken"`
	SecondaryToken *string `pulumi:"secondaryToken"`
}
