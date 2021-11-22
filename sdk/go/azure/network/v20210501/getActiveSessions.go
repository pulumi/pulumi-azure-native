


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetActiveSessions(ctx *pulumi.Context, args *GetActiveSessionsArgs, opts ...pulumi.InvokeOption) (*GetActiveSessionsResult, error) {
	var rv GetActiveSessionsResult
	err := ctx.Invoke("azure-native:network/v20210501:getActiveSessions", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetActiveSessionsArgs struct {
	BastionHostName   string `pulumi:"bastionHostName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetActiveSessionsResult struct {
	NextLink *string                        `pulumi:"nextLink"`
	Value    []BastionActiveSessionResponse `pulumi:"value"`
}
