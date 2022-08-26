


package v20180601preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListTransactionNodeApiKeys(ctx *pulumi.Context, args *ListTransactionNodeApiKeysArgs, opts ...pulumi.InvokeOption) (*ListTransactionNodeApiKeysResult, error) {
	var rv ListTransactionNodeApiKeysResult
	err := ctx.Invoke("azure-native:blockchain/v20180601preview:listTransactionNodeApiKeys", args, &rv, opts...)
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

func ListTransactionNodeApiKeysOutput(ctx *pulumi.Context, args ListTransactionNodeApiKeysOutputArgs, opts ...pulumi.InvokeOption) ListTransactionNodeApiKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListTransactionNodeApiKeysResult, error) {
			args := v.(ListTransactionNodeApiKeysArgs)
			r, err := ListTransactionNodeApiKeys(ctx, &args, opts...)
			var s ListTransactionNodeApiKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListTransactionNodeApiKeysResultOutput)
}

type ListTransactionNodeApiKeysOutputArgs struct {
	BlockchainMemberName pulumi.StringInput `pulumi:"blockchainMemberName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
	TransactionNodeName  pulumi.StringInput `pulumi:"transactionNodeName"`
}

func (ListTransactionNodeApiKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTransactionNodeApiKeysArgs)(nil)).Elem()
}


type ListTransactionNodeApiKeysResultOutput struct{ *pulumi.OutputState }

func (ListTransactionNodeApiKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListTransactionNodeApiKeysResult)(nil)).Elem()
}

func (o ListTransactionNodeApiKeysResultOutput) ToListTransactionNodeApiKeysResultOutput() ListTransactionNodeApiKeysResultOutput {
	return o
}

func (o ListTransactionNodeApiKeysResultOutput) ToListTransactionNodeApiKeysResultOutputWithContext(ctx context.Context) ListTransactionNodeApiKeysResultOutput {
	return o
}

func (o ListTransactionNodeApiKeysResultOutput) Keys() ApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListTransactionNodeApiKeysResult) []ApiKeyResponse { return v.Keys }).(ApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListTransactionNodeApiKeysResultOutput{})
}
