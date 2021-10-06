


package labservices

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListGlobalUserLabs(ctx *pulumi.Context, args *ListGlobalUserLabsArgs, opts ...pulumi.InvokeOption) (*ListGlobalUserLabsResult, error) {
	var rv ListGlobalUserLabsResult
	err := ctx.Invoke("azure-native:labservices:listGlobalUserLabs", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListGlobalUserLabsArgs struct {
	UserName string `pulumi:"userName"`
}


type ListGlobalUserLabsResult struct {
	Labs []LabDetailsResponse `pulumi:"labs"`
}
