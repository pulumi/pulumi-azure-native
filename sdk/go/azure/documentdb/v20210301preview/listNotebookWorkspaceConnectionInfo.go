


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListNotebookWorkspaceConnectionInfo(ctx *pulumi.Context, args *ListNotebookWorkspaceConnectionInfoArgs, opts ...pulumi.InvokeOption) (*ListNotebookWorkspaceConnectionInfoResult, error) {
	var rv ListNotebookWorkspaceConnectionInfoResult
	err := ctx.Invoke("azure-native:documentdb/v20210301preview:listNotebookWorkspaceConnectionInfo", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListNotebookWorkspaceConnectionInfoArgs struct {
	AccountName           string `pulumi:"accountName"`
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type ListNotebookWorkspaceConnectionInfoResult struct {
	AuthToken              string `pulumi:"authToken"`
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
}

func ListNotebookWorkspaceConnectionInfoOutput(ctx *pulumi.Context, args ListNotebookWorkspaceConnectionInfoOutputArgs, opts ...pulumi.InvokeOption) ListNotebookWorkspaceConnectionInfoResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListNotebookWorkspaceConnectionInfoResult, error) {
			args := v.(ListNotebookWorkspaceConnectionInfoArgs)
			r, err := ListNotebookWorkspaceConnectionInfo(ctx, &args, opts...)
			var s ListNotebookWorkspaceConnectionInfoResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListNotebookWorkspaceConnectionInfoResultOutput)
}

type ListNotebookWorkspaceConnectionInfoOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	NotebookWorkspaceName pulumi.StringInput `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListNotebookWorkspaceConnectionInfoOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookWorkspaceConnectionInfoArgs)(nil)).Elem()
}


type ListNotebookWorkspaceConnectionInfoResultOutput struct{ *pulumi.OutputState }

func (ListNotebookWorkspaceConnectionInfoResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListNotebookWorkspaceConnectionInfoResult)(nil)).Elem()
}

func (o ListNotebookWorkspaceConnectionInfoResultOutput) ToListNotebookWorkspaceConnectionInfoResultOutput() ListNotebookWorkspaceConnectionInfoResultOutput {
	return o
}

func (o ListNotebookWorkspaceConnectionInfoResultOutput) ToListNotebookWorkspaceConnectionInfoResultOutputWithContext(ctx context.Context) ListNotebookWorkspaceConnectionInfoResultOutput {
	return o
}

func (o ListNotebookWorkspaceConnectionInfoResultOutput) AuthToken() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookWorkspaceConnectionInfoResult) string { return v.AuthToken }).(pulumi.StringOutput)
}

func (o ListNotebookWorkspaceConnectionInfoResultOutput) NotebookServerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v ListNotebookWorkspaceConnectionInfoResult) string { return v.NotebookServerEndpoint }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListNotebookWorkspaceConnectionInfoResultOutput{})
}
