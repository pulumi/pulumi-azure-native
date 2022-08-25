


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSiteBackups(ctx *pulumi.Context, args *ListWebAppSiteBackupsArgs, opts ...pulumi.InvokeOption) (*ListWebAppSiteBackupsResult, error) {
	var rv ListWebAppSiteBackupsResult
	err := ctx.Invoke("azure-native:web/v20210101:listWebAppSiteBackups", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSiteBackupsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppSiteBackupsResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []BackupItemResponse `pulumi:"value"`
}

func ListWebAppSiteBackupsOutput(ctx *pulumi.Context, args ListWebAppSiteBackupsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSiteBackupsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSiteBackupsResult, error) {
			args := v.(ListWebAppSiteBackupsArgs)
			r, err := ListWebAppSiteBackups(ctx, &args, opts...)
			var s ListWebAppSiteBackupsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSiteBackupsResultOutput)
}

type ListWebAppSiteBackupsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppSiteBackupsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsArgs)(nil)).Elem()
}


type ListWebAppSiteBackupsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSiteBackupsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsResult)(nil)).Elem()
}

func (o ListWebAppSiteBackupsResultOutput) ToListWebAppSiteBackupsResultOutput() ListWebAppSiteBackupsResultOutput {
	return o
}

func (o ListWebAppSiteBackupsResultOutput) ToListWebAppSiteBackupsResultOutputWithContext(ctx context.Context) ListWebAppSiteBackupsResultOutput {
	return o
}

func (o ListWebAppSiteBackupsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListWebAppSiteBackupsResultOutput) Value() BackupItemResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsResult) []BackupItemResponse { return v.Value }).(BackupItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSiteBackupsResultOutput{})
}
