


package blockchain

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBlockchainMemberApiKeys(ctx *pulumi.Context, args *ListBlockchainMemberApiKeysArgs, opts ...pulumi.InvokeOption) (*ListBlockchainMemberApiKeysResult, error) {
	var rv ListBlockchainMemberApiKeysResult
	err := ctx.Invoke("azure-native:blockchain:listBlockchainMemberApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBlockchainMemberApiKeysArgs struct {
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type ListBlockchainMemberApiKeysResult struct {
	Keys []ApiKeyResponse `pulumi:"keys"`
}
