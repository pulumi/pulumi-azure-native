


package v20220701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSourceControlRepositories(ctx *pulumi.Context, args *ListSourceControlRepositoriesArgs, opts ...pulumi.InvokeOption) (*ListSourceControlRepositoriesResult, error) {
	var rv ListSourceControlRepositoriesResult
	err := ctx.Invoke("azure-native:securityinsights/v20220701preview:listSourceControlRepositories", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSourceControlRepositoriesArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type ListSourceControlRepositoriesResult struct {
	NextLink string         `pulumi:"nextLink"`
	Value    []RepoResponse `pulumi:"value"`
}

func ListSourceControlRepositoriesOutput(ctx *pulumi.Context, args ListSourceControlRepositoriesOutputArgs, opts ...pulumi.InvokeOption) ListSourceControlRepositoriesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSourceControlRepositoriesResult, error) {
			args := v.(ListSourceControlRepositoriesArgs)
			r, err := ListSourceControlRepositories(ctx, &args, opts...)
			var s ListSourceControlRepositoriesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSourceControlRepositoriesResultOutput)
}

type ListSourceControlRepositoriesOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (ListSourceControlRepositoriesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSourceControlRepositoriesArgs)(nil)).Elem()
}


type ListSourceControlRepositoriesResultOutput struct{ *pulumi.OutputState }

func (ListSourceControlRepositoriesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSourceControlRepositoriesResult)(nil)).Elem()
}

func (o ListSourceControlRepositoriesResultOutput) ToListSourceControlRepositoriesResultOutput() ListSourceControlRepositoriesResultOutput {
	return o
}

func (o ListSourceControlRepositoriesResultOutput) ToListSourceControlRepositoriesResultOutputWithContext(ctx context.Context) ListSourceControlRepositoriesResultOutput {
	return o
}

func (o ListSourceControlRepositoriesResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListSourceControlRepositoriesResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListSourceControlRepositoriesResultOutput) Value() RepoResponseArrayOutput {
	return o.ApplyT(func(v ListSourceControlRepositoriesResult) []RepoResponse { return v.Value }).(RepoResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSourceControlRepositoriesResultOutput{})
}
