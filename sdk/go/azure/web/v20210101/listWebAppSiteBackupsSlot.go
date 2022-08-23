


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppSiteBackupsSlot(ctx *pulumi.Context, args *ListWebAppSiteBackupsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppSiteBackupsSlotResult, error) {
	var rv ListWebAppSiteBackupsSlotResult
	err := ctx.Invoke("azure-native:web/v20210101:listWebAppSiteBackupsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppSiteBackupsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppSiteBackupsSlotResult struct {
	NextLink string               `pulumi:"nextLink"`
	Value    []BackupItemResponse `pulumi:"value"`
}

func ListWebAppSiteBackupsSlotOutput(ctx *pulumi.Context, args ListWebAppSiteBackupsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppSiteBackupsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppSiteBackupsSlotResult, error) {
			args := v.(ListWebAppSiteBackupsSlotArgs)
			r, err := ListWebAppSiteBackupsSlot(ctx, &args, opts...)
			var s ListWebAppSiteBackupsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppSiteBackupsSlotResultOutput)
}

type ListWebAppSiteBackupsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppSiteBackupsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsSlotArgs)(nil)).Elem()
}


type ListWebAppSiteBackupsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppSiteBackupsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppSiteBackupsSlotResult)(nil)).Elem()
}

func (o ListWebAppSiteBackupsSlotResultOutput) ToListWebAppSiteBackupsSlotResultOutput() ListWebAppSiteBackupsSlotResultOutput {
	return o
}

func (o ListWebAppSiteBackupsSlotResultOutput) ToListWebAppSiteBackupsSlotResultOutputWithContext(ctx context.Context) ListWebAppSiteBackupsSlotResultOutput {
	return o
}

func (o ListWebAppSiteBackupsSlotResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsSlotResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListWebAppSiteBackupsSlotResultOutput) Value() BackupItemResponseArrayOutput {
	return o.ApplyT(func(v ListWebAppSiteBackupsSlotResult) []BackupItemResponse { return v.Value }).(BackupItemResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppSiteBackupsSlotResultOutput{})
}
