


package v20180202

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetProjectKeys(ctx *pulumi.Context, args *GetProjectKeysArgs, opts ...pulumi.InvokeOption) (*GetProjectKeysResult, error) {
	var rv GetProjectKeysResult
	err := ctx.Invoke("azure-native:migrate/v20180202:getProjectKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetProjectKeysArgs struct {
	ProjectName       string `pulumi:"projectName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetProjectKeysResult struct {
	WorkspaceId  string `pulumi:"workspaceId"`
	WorkspaceKey string `pulumi:"workspaceKey"`
}
