


package v20180601preview

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListLocationConsortiums(ctx *pulumi.Context, args *ListLocationConsortiumsArgs, opts ...pulumi.InvokeOption) (*ListLocationConsortiumsResult, error) {
	var rv ListLocationConsortiumsResult
	err := ctx.Invoke("azure-native:blockchain/v20180601preview:listLocationConsortiums", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListLocationConsortiumsArgs struct {
	LocationName string `pulumi:"locationName"`
}


type ListLocationConsortiumsResult struct {
	Value []ConsortiumResponse `pulumi:"value"`
}
