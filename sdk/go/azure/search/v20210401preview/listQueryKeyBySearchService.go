


package v20210401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListQueryKeyBySearchService(ctx *pulumi.Context, args *ListQueryKeyBySearchServiceArgs, opts ...pulumi.InvokeOption) (*ListQueryKeyBySearchServiceResult, error) {
	var rv ListQueryKeyBySearchServiceResult
	err := ctx.Invoke("azure-native:search/v20210401preview:listQueryKeyBySearchService", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListQueryKeyBySearchServiceArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	SearchServiceName string `pulumi:"searchServiceName"`
}


type ListQueryKeyBySearchServiceResult struct {
	NextLink string             `pulumi:"nextLink"`
	Value    []QueryKeyResponse `pulumi:"value"`
}

func ListQueryKeyBySearchServiceOutput(ctx *pulumi.Context, args ListQueryKeyBySearchServiceOutputArgs, opts ...pulumi.InvokeOption) ListQueryKeyBySearchServiceResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListQueryKeyBySearchServiceResult, error) {
			args := v.(ListQueryKeyBySearchServiceArgs)
			r, err := ListQueryKeyBySearchService(ctx, &args, opts...)
			var s ListQueryKeyBySearchServiceResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListQueryKeyBySearchServiceResultOutput)
}

type ListQueryKeyBySearchServiceOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	SearchServiceName pulumi.StringInput `pulumi:"searchServiceName"`
}

func (ListQueryKeyBySearchServiceOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQueryKeyBySearchServiceArgs)(nil)).Elem()
}


type ListQueryKeyBySearchServiceResultOutput struct{ *pulumi.OutputState }

func (ListQueryKeyBySearchServiceResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListQueryKeyBySearchServiceResult)(nil)).Elem()
}

func (o ListQueryKeyBySearchServiceResultOutput) ToListQueryKeyBySearchServiceResultOutput() ListQueryKeyBySearchServiceResultOutput {
	return o
}

func (o ListQueryKeyBySearchServiceResultOutput) ToListQueryKeyBySearchServiceResultOutputWithContext(ctx context.Context) ListQueryKeyBySearchServiceResultOutput {
	return o
}

func (o ListQueryKeyBySearchServiceResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListQueryKeyBySearchServiceResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListQueryKeyBySearchServiceResultOutput) Value() QueryKeyResponseArrayOutput {
	return o.ApplyT(func(v ListQueryKeyBySearchServiceResult) []QueryKeyResponse { return v.Value }).(QueryKeyResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListQueryKeyBySearchServiceResultOutput{})
}
