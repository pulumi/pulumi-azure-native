


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserOperationStatus(ctx *pulumi.Context, args *GetGlobalUserOperationStatusArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserOperationStatusResult, error) {
	var rv GetGlobalUserOperationStatusResult
	err := ctx.Invoke("azure-native:labservices:getGlobalUserOperationStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserOperationStatusArgs struct {
	OperationUrl string `pulumi:"operationUrl"`
	UserName     string `pulumi:"userName"`
}


type GetGlobalUserOperationStatusResult struct {
	Status string `pulumi:"status"`
}
