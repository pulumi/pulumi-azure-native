


package v20181015

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetGlobalUserOperationBatchStatus(ctx *pulumi.Context, args *GetGlobalUserOperationBatchStatusArgs, opts ...pulumi.InvokeOption) (*GetGlobalUserOperationBatchStatusResult, error) {
	var rv GetGlobalUserOperationBatchStatusResult
	err := ctx.Invoke("azure-native:labservices/v20181015:getGlobalUserOperationBatchStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetGlobalUserOperationBatchStatusArgs struct {
	Urls     []string `pulumi:"urls"`
	UserName string   `pulumi:"userName"`
}


type GetGlobalUserOperationBatchStatusResult struct {
	Items []OperationBatchStatusResponseItemResponse `pulumi:"items"`
}
