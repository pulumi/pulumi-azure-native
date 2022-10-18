


package machinelearningservices

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListStorageAccountKeys(ctx *pulumi.Context, args *ListStorageAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListStorageAccountKeysResult, error) {
	var rv ListStorageAccountKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices:listStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListStorageAccountKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListStorageAccountKeysResult struct {
	UserStorageKey string `pulumi:"userStorageKey"`
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
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
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

func (o ListStorageAccountKeysResultOutput) UserStorageKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListStorageAccountKeysResult) string { return v.UserStorageKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListStorageAccountKeysResultOutput{})
}
