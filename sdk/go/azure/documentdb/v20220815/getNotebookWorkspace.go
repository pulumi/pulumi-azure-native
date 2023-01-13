


package v20220815

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNotebookWorkspace(ctx *pulumi.Context, args *LookupNotebookWorkspaceArgs, opts ...pulumi.InvokeOption) (*LookupNotebookWorkspaceResult, error) {
	var rv LookupNotebookWorkspaceResult
	err := ctx.Invoke("azure-native:documentdb/v20220815:getNotebookWorkspace", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNotebookWorkspaceArgs struct {
	AccountName           string `pulumi:"accountName"`
	NotebookWorkspaceName string `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     string `pulumi:"resourceGroupName"`
}


type LookupNotebookWorkspaceResult struct {
	Id                     string `pulumi:"id"`
	Name                   string `pulumi:"name"`
	NotebookServerEndpoint string `pulumi:"notebookServerEndpoint"`
	Status                 string `pulumi:"status"`
	Type                   string `pulumi:"type"`
}

func LookupNotebookWorkspaceOutput(ctx *pulumi.Context, args LookupNotebookWorkspaceOutputArgs, opts ...pulumi.InvokeOption) LookupNotebookWorkspaceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNotebookWorkspaceResult, error) {
			args := v.(LookupNotebookWorkspaceArgs)
			r, err := LookupNotebookWorkspace(ctx, &args, opts...)
			var s LookupNotebookWorkspaceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNotebookWorkspaceResultOutput)
}

type LookupNotebookWorkspaceOutputArgs struct {
	AccountName           pulumi.StringInput `pulumi:"accountName"`
	NotebookWorkspaceName pulumi.StringInput `pulumi:"notebookWorkspaceName"`
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNotebookWorkspaceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookWorkspaceArgs)(nil)).Elem()
}


type LookupNotebookWorkspaceResultOutput struct{ *pulumi.OutputState }

func (LookupNotebookWorkspaceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNotebookWorkspaceResult)(nil)).Elem()
}

func (o LookupNotebookWorkspaceResultOutput) ToLookupNotebookWorkspaceResultOutput() LookupNotebookWorkspaceResultOutput {
	return o
}

func (o LookupNotebookWorkspaceResultOutput) ToLookupNotebookWorkspaceResultOutputWithContext(ctx context.Context) LookupNotebookWorkspaceResultOutput {
	return o
}

func (o LookupNotebookWorkspaceResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupNotebookWorkspaceResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupNotebookWorkspaceResultOutput) NotebookServerEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.NotebookServerEndpoint }).(pulumi.StringOutput)
}

func (o LookupNotebookWorkspaceResultOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Status }).(pulumi.StringOutput)
}

func (o LookupNotebookWorkspaceResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNotebookWorkspaceResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNotebookWorkspaceResultOutput{})
}
