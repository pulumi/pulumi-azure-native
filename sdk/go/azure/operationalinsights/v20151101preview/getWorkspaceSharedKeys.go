


package v20151101preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetWorkspaceSharedKeys(ctx *pulumi.Context, args *GetWorkspaceSharedKeysArgs, opts ...pulumi.InvokeOption) (*GetWorkspaceSharedKeysResult, error) {
	var rv GetWorkspaceSharedKeysResult
	err := ctx.Invoke("azure-native:operationalinsights/v20151101preview:getWorkspaceSharedKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetWorkspaceSharedKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetWorkspaceSharedKeysResult struct {
	PrimarySharedKey   *string `pulumi:"primarySharedKey"`
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
}
