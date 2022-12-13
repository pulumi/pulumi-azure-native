


package blockchain

import (
	"context"
	"reflect"

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

func ListBlockchainMemberApiKeysOutput(ctx *pulumi.Context, args ListBlockchainMemberApiKeysOutputArgs, opts ...pulumi.InvokeOption) ListBlockchainMemberApiKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBlockchainMemberApiKeysResult, error) {
			args := v.(ListBlockchainMemberApiKeysArgs)
			r, err := ListBlockchainMemberApiKeys(ctx, &args, opts...)
			var s ListBlockchainMemberApiKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBlockchainMemberApiKeysResultOutput)
}

type ListBlockchainMemberApiKeysOutputArgs struct {
	BlockchainMemberName pulumi.StringInput `pulumi:"blockchainMemberName"`
	ResourceGroupName    pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBlockchainMemberApiKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBlockchainMemberApiKeysArgs)(nil)).Elem()
}


type ListBlockchainMemberApiKeysResultOutput struct{ *pulumi.OutputState }

func (ListBlockchainMemberApiKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBlockchainMemberApiKeysResult)(nil)).Elem()
}

func (o ListBlockchainMemberApiKeysResultOutput) ToListBlockchainMemberApiKeysResultOutput() ListBlockchainMemberApiKeysResultOutput {
	return o
}

func (o ListBlockchainMemberApiKeysResultOutput) ToListBlockchainMemberApiKeysResultOutputWithContext(ctx context.Context) ListBlockchainMemberApiKeysResultOutput {
	return o
}

func (o ListBlockchainMemberApiKeysResultOutput) Keys() ApiKeyResponseArrayOutput {
	return o.ApplyT(func(v ListBlockchainMemberApiKeysResult) []ApiKeyResponse { return v.Keys }).(ApiKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBlockchainMemberApiKeysResultOutput{})
}
