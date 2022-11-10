


package powerbi

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceCollectionAccessKeys(ctx *pulumi.Context, args *ListWorkspaceCollectionAccessKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceCollectionAccessKeysResult, error) {
	var rv ListWorkspaceCollectionAccessKeysResult
	err := ctx.Invoke("azure-native:powerbi:listWorkspaceCollectionAccessKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceCollectionAccessKeysArgs struct {
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	WorkspaceCollectionName string `pulumi:"workspaceCollectionName"`
}

type ListWorkspaceCollectionAccessKeysResult struct {
	Key1 *string `pulumi:"key1"`
	Key2 *string `pulumi:"key2"`
}

func ListWorkspaceCollectionAccessKeysOutput(ctx *pulumi.Context, args ListWorkspaceCollectionAccessKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceCollectionAccessKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceCollectionAccessKeysResult, error) {
			args := v.(ListWorkspaceCollectionAccessKeysArgs)
			r, err := ListWorkspaceCollectionAccessKeys(ctx, &args, opts...)
			var s ListWorkspaceCollectionAccessKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceCollectionAccessKeysResultOutput)
}

type ListWorkspaceCollectionAccessKeysOutputArgs struct {
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceCollectionName pulumi.StringInput `pulumi:"workspaceCollectionName"`
}

func (ListWorkspaceCollectionAccessKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceCollectionAccessKeysArgs)(nil)).Elem()
}

type ListWorkspaceCollectionAccessKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceCollectionAccessKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceCollectionAccessKeysResult)(nil)).Elem()
}

func (o ListWorkspaceCollectionAccessKeysResultOutput) ToListWorkspaceCollectionAccessKeysResultOutput() ListWorkspaceCollectionAccessKeysResultOutput {
	return o
}

func (o ListWorkspaceCollectionAccessKeysResultOutput) ToListWorkspaceCollectionAccessKeysResultOutputWithContext(ctx context.Context) ListWorkspaceCollectionAccessKeysResultOutput {
	return o
}

func (o ListWorkspaceCollectionAccessKeysResultOutput) Key1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceCollectionAccessKeysResult) *string { return v.Key1 }).(pulumi.StringPtrOutput)
}

func (o ListWorkspaceCollectionAccessKeysResultOutput) Key2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWorkspaceCollectionAccessKeysResult) *string { return v.Key2 }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceCollectionAccessKeysResultOutput{})
}
