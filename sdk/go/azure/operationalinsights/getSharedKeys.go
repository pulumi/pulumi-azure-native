


package operationalinsights

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetSharedKeys(ctx *pulumi.Context, args *GetSharedKeysArgs, opts ...pulumi.InvokeOption) (*GetSharedKeysResult, error) {
	var rv GetSharedKeysResult
	err := ctx.Invoke("azure-native:operationalinsights:getSharedKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetSharedKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetSharedKeysResult struct {
	PrimarySharedKey   *string `pulumi:"primarySharedKey"`
	SecondarySharedKey *string `pulumi:"secondarySharedKey"`
}
