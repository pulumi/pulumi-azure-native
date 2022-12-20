


package v20221001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWorkspaceNotebookKeys(ctx *pulumi.Context, args *ListWorkspaceNotebookKeysArgs, opts ...pulumi.InvokeOption) (*ListWorkspaceNotebookKeysResult, error) {
	var rv ListWorkspaceNotebookKeysResult
	err := ctx.Invoke("azure-native:machinelearningservices/v20221001:listWorkspaceNotebookKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWorkspaceNotebookKeysArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}

type ListWorkspaceNotebookKeysResult struct {
	PrimaryAccessKey   string `pulumi:"primaryAccessKey"`
	SecondaryAccessKey string `pulumi:"secondaryAccessKey"`
}

func ListWorkspaceNotebookKeysOutput(ctx *pulumi.Context, args ListWorkspaceNotebookKeysOutputArgs, opts ...pulumi.InvokeOption) ListWorkspaceNotebookKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWorkspaceNotebookKeysResult, error) {
			args := v.(ListWorkspaceNotebookKeysArgs)
			r, err := ListWorkspaceNotebookKeys(ctx, &args, opts...)
			var s ListWorkspaceNotebookKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWorkspaceNotebookKeysResultOutput)
}

type ListWorkspaceNotebookKeysOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListWorkspaceNotebookKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookKeysArgs)(nil)).Elem()
}

type ListWorkspaceNotebookKeysResultOutput struct{ *pulumi.OutputState }

func (ListWorkspaceNotebookKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWorkspaceNotebookKeysResult)(nil)).Elem()
}

func (o ListWorkspaceNotebookKeysResultOutput) ToListWorkspaceNotebookKeysResultOutput() ListWorkspaceNotebookKeysResultOutput {
	return o
}

func (o ListWorkspaceNotebookKeysResultOutput) ToListWorkspaceNotebookKeysResultOutputWithContext(ctx context.Context) ListWorkspaceNotebookKeysResultOutput {
	return o
}

func (o ListWorkspaceNotebookKeysResultOutput) PrimaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookKeysResult) string { return v.PrimaryAccessKey }).(pulumi.StringOutput)
}

func (o ListWorkspaceNotebookKeysResultOutput) SecondaryAccessKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListWorkspaceNotebookKeysResult) string { return v.SecondaryAccessKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWorkspaceNotebookKeysResultOutput{})
}
