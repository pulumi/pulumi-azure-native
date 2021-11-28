


package blockchain

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTransactionNodeApiKeys(ctx *pulumi.Context, args *ListTransactionNodeApiKeysArgs, opts ...pulumi.InvokeOption) (*ListTransactionNodeApiKeysResult, error) {
	var rv ListTransactionNodeApiKeysResult
	err := ctx.Invoke("azure-native:blockchain:listTransactionNodeApiKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListTransactionNodeApiKeysArgs struct {
	BlockchainMemberName string `pulumi:"blockchainMemberName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
	TransactionNodeName  string `pulumi:"transactionNodeName"`
}


type ListTransactionNodeApiKeysResult struct {
	Keys []ApiKeyResponse `pulumi:"keys"`
}
