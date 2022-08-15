


package batch

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListBatchAccountKeys(ctx *pulumi.Context, args *ListBatchAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListBatchAccountKeysResult, error) {
	var rv ListBatchAccountKeysResult
	err := ctx.Invoke("azure-native:batch:listBatchAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListBatchAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListBatchAccountKeysResult struct {
	AccountName string `pulumi:"accountName"`
	Primary     string `pulumi:"primary"`
	Secondary   string `pulumi:"secondary"`
}

func ListBatchAccountKeysOutput(ctx *pulumi.Context, args ListBatchAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListBatchAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListBatchAccountKeysResult, error) {
			args := v.(ListBatchAccountKeysArgs)
			r, err := ListBatchAccountKeys(ctx, &args, opts...)
			var s ListBatchAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListBatchAccountKeysResultOutput)
}

type ListBatchAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListBatchAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBatchAccountKeysArgs)(nil)).Elem()
}


type ListBatchAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListBatchAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListBatchAccountKeysResult)(nil)).Elem()
}

func (o ListBatchAccountKeysResultOutput) ToListBatchAccountKeysResultOutput() ListBatchAccountKeysResultOutput {
	return o
}

func (o ListBatchAccountKeysResultOutput) ToListBatchAccountKeysResultOutputWithContext(ctx context.Context) ListBatchAccountKeysResultOutput {
	return o
}

func (o ListBatchAccountKeysResultOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v ListBatchAccountKeysResult) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o ListBatchAccountKeysResultOutput) Primary() pulumi.StringOutput {
	return o.ApplyT(func(v ListBatchAccountKeysResult) string { return v.Primary }).(pulumi.StringOutput)
}

func (o ListBatchAccountKeysResultOutput) Secondary() pulumi.StringOutput {
	return o.ApplyT(func(v ListBatchAccountKeysResult) string { return v.Secondary }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListBatchAccountKeysResultOutput{})
}
