


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListObjectAnchorsAccountKeys(ctx *pulumi.Context, args *ListObjectAnchorsAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListObjectAnchorsAccountKeysResult, error) {
	var rv ListObjectAnchorsAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality/v20210301preview:listObjectAnchorsAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListObjectAnchorsAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListObjectAnchorsAccountKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListObjectAnchorsAccountKeysOutput(ctx *pulumi.Context, args ListObjectAnchorsAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListObjectAnchorsAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListObjectAnchorsAccountKeysResult, error) {
			args := v.(ListObjectAnchorsAccountKeysArgs)
			r, err := ListObjectAnchorsAccountKeys(ctx, &args, opts...)
			var s ListObjectAnchorsAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListObjectAnchorsAccountKeysResultOutput)
}

type ListObjectAnchorsAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListObjectAnchorsAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListObjectAnchorsAccountKeysArgs)(nil)).Elem()
}


type ListObjectAnchorsAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListObjectAnchorsAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListObjectAnchorsAccountKeysResult)(nil)).Elem()
}

func (o ListObjectAnchorsAccountKeysResultOutput) ToListObjectAnchorsAccountKeysResultOutput() ListObjectAnchorsAccountKeysResultOutput {
	return o
}

func (o ListObjectAnchorsAccountKeysResultOutput) ToListObjectAnchorsAccountKeysResultOutputWithContext(ctx context.Context) ListObjectAnchorsAccountKeysResultOutput {
	return o
}

func (o ListObjectAnchorsAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListObjectAnchorsAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListObjectAnchorsAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListObjectAnchorsAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListObjectAnchorsAccountKeysResultOutput{})
}
