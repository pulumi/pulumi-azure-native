


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSpatialAnchorsAccountKeys(ctx *pulumi.Context, args *ListSpatialAnchorsAccountKeysArgs, opts ...pulumi.InvokeOption) (*ListSpatialAnchorsAccountKeysResult, error) {
	var rv ListSpatialAnchorsAccountKeysResult
	err := ctx.Invoke("azure-native:mixedreality/v20200501:listSpatialAnchorsAccountKeys", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSpatialAnchorsAccountKeysArgs struct {
	AccountName       string `pulumi:"accountName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSpatialAnchorsAccountKeysResult struct {
	PrimaryKey   string `pulumi:"primaryKey"`
	SecondaryKey string `pulumi:"secondaryKey"`
}

func ListSpatialAnchorsAccountKeysOutput(ctx *pulumi.Context, args ListSpatialAnchorsAccountKeysOutputArgs, opts ...pulumi.InvokeOption) ListSpatialAnchorsAccountKeysResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSpatialAnchorsAccountKeysResult, error) {
			args := v.(ListSpatialAnchorsAccountKeysArgs)
			r, err := ListSpatialAnchorsAccountKeys(ctx, &args, opts...)
			var s ListSpatialAnchorsAccountKeysResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSpatialAnchorsAccountKeysResultOutput)
}

type ListSpatialAnchorsAccountKeysOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSpatialAnchorsAccountKeysOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpatialAnchorsAccountKeysArgs)(nil)).Elem()
}


type ListSpatialAnchorsAccountKeysResultOutput struct{ *pulumi.OutputState }

func (ListSpatialAnchorsAccountKeysResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSpatialAnchorsAccountKeysResult)(nil)).Elem()
}

func (o ListSpatialAnchorsAccountKeysResultOutput) ToListSpatialAnchorsAccountKeysResultOutput() ListSpatialAnchorsAccountKeysResultOutput {
	return o
}

func (o ListSpatialAnchorsAccountKeysResultOutput) ToListSpatialAnchorsAccountKeysResultOutputWithContext(ctx context.Context) ListSpatialAnchorsAccountKeysResultOutput {
	return o
}

func (o ListSpatialAnchorsAccountKeysResultOutput) PrimaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListSpatialAnchorsAccountKeysResult) string { return v.PrimaryKey }).(pulumi.StringOutput)
}

func (o ListSpatialAnchorsAccountKeysResultOutput) SecondaryKey() pulumi.StringOutput {
	return o.ApplyT(func(v ListSpatialAnchorsAccountKeysResult) string { return v.SecondaryKey }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSpatialAnchorsAccountKeysResultOutput{})
}
