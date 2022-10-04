


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListRemoteRenderingAccountKeys(ctx *pulumi.Context, args *ListRemoteRenderingAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListRemoteRenderingAccountKeysResult, error) {
	var rv ListRemoteRenderingAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality/v20210301preview:listRemoteRenderingAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListRemoteRenderingAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListRemoteRenderingAccountKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListRemoteRenderingAccountKeysOutput(ctx *pulumi.Context, args ListRemoteRenderingAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListRemoteRenderingAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListRemoteRenderingAccountKeysResult, error) {
			args := v.(ListRemoteRenderingAccountKeysArgs)
			r, err := ListRemoteRenderingAccountKeys(ctx, &args, opts...)
			var s ListRemoteRenderingAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListRemoteRenderingAccountKeysResultOutput)
}

type ListRemoteRenderingAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListRemoteRenderingAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemoteRenderingAccountKeysArgs)(nil)).Elem()
}


type ListRemoteRenderingAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListRemoteRenderingAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListRemoteRenderingAccountKeysResult)(nil)).Elem()
}

func (o ListRemoteRenderingAccountKeysResultOutput) ToListRemoteRenderingAccountKeysResultOutput() ListRemoteRenderingAccountKeysResultOutput {
	return o
}

func (o ListRemoteRenderingAccountKeysResultOutput) ToListRemoteRenderingAccountKeysResultOutputWithContext(ctx context.Context) ListRemoteRenderingAccountKeysResultOutput {
	return o
}

func (o ListRemoteRenderingAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemoteRenderingAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListRemoteRenderingAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListRemoteRenderingAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListRemoteRenderingAccountKeysResultOutput{})
}
