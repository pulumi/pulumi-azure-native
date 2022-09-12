


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountKeys(ctx *pulumi.Context, args *ListStorageAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountKeysResult, error) {
	var rv ListStorageAccountKeysResult
	err := ctx.Invoke("azure-native:storage/v20220501:listStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountKeysArgs struct {
	AccountName       string  `pulumi:"accountName"`
	Expand            *string `pulumi:"expand"`
	ResourceGroupName string  `pulumi:"resourceGroupName"`
}


type ListStorageAccountKeysResult struct {
	Keys []StorageAccountKeyResponse `pulumi:"keys"`
}

func ListStorageAccountKeysOutput(ctx *pulumi.Context, args ListStorageAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListStorageAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListStorageAccountKeysResult, error) {
			args := v.(ListStorageAccountKeysArgs)
			r, err := ListStorageAccountKeys(ctx, &args, opts...)
			var s ListStorageAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListStorageAccountKeysResultOutput)
}

type ListStorageAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput    `pulumi:"accountName"`
	Expand            pulumi.StringPtrInput `pulumi:"expand"`
	ResourceGroupName pulumi.StringInput    `pulumi:"resourceGroupName"`
}

func (ListStorageAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysArgs)(nil)).Elem()
}


type ListStorageAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListStorageAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListStorageAccountKeysResult)(nil)).Elem()
}

func (o ListStorageAccountKeysResultOutput) ToListStorageAccountKeysResultOutput() ListStorageAccountKeysResultOutput {
	return o
}

func (o ListStorageAccountKeysResultOutput) ToListStorageAccountKeysResultOutputWithContext(ctx context.Context) ListStorageAccountKeysResultOutput {
	return o
}

func (o ListStorageAccountKeysResultOutput) Keys() StorageAccountKeyResponseArrayOutput {
	return o.ApplyT(func(v ListStorageAccountKeysResult) []StorageAccountKeyResponse { return v.Keys }).(StorageAccountKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountKeysResultOutput{})
}
