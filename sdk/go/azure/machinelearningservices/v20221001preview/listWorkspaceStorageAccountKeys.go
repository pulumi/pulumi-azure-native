


package v20221001preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceStorageAccountKeys(ctx *pulumi.Context, args *ListWorkspaceStorageAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceStorageAccountKeysResult, error) {
	var rv ListWorkspaceStorageAccountKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001preview:listWorkspaceStorageAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceStorageAccountKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceStorageAccountKeysResult struct {
	UserStorageKey string `pulumi:"userStorageKey"`
}

func ListWorkspaceStorageAccountKeysOutput(ctx *pulumi.Context, args ListWorkspaceStorageAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceStorageAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceStorageAccountKeysResult, error) {
			args := v.(ListWorkspaceStorageAccountKeysArgs)
			r, err := ListWorkspaceStorageAccountKeys(ctx, &args, opts...)
			var s ListWorkspaceStorageAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceStorageAccountKeysResultOutput)
}

type ListWorkspaceStorageAccountKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceStorageAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceStorageAccountKeysArgs)(nil)).Elem()
}

type ListWorkspaceStorageAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceStorageAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceStorageAccountKeysResult)(nil)).Elem()
}

func (o ListWorkspaceStorageAccountKeysResultOutput) ToListWorkspaceStorageAccountKeysResultOutput() ListWorkspaceStorageAccountKeysResultOutput {
	return o
}

func (o ListWorkspaceStorageAccountKeysResultOutput) ToListWorkspaceStorageAccountKeysResultOutputWithContext(ctx context.Context) ListWorkspaceStorageAccountKeysResultOutput {
	return o
}

func (o ListWorkspaceStorageAccountKeysResultOutput) UserStorageKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceStorageAccountKeysResult) string { return v.UserStorageKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceStorageAccountKeysResultOutput{})
}
