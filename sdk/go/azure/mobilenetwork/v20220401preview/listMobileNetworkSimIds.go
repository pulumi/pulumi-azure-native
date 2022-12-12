


package v20220401preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListMobileNetworkSimIds(ctx *pulumi.Context, args *ListMobileNetworkSimIdsArgs, opts ...pulumi.InvokeOption) (*ListMobileNetworkSimIdsResult, error) {
	var rv ListMobileNetworkSimIdsResult
	err := ctx.Invoke("azure-native:mobilenetwork/v20220401preview:listMobileNetworkSimIds", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListMobileNetworkSimIdsArgs struct {
	MobileNetworkName string `pulumi:"mobileNetworkName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListMobileNetworkSimIdsResult struct {
	NextLink string                `pulumi:"nextLink"`
	Value    []SubResourceResponse `pulumi:"value"`
}

func ListMobileNetworkSimIdsOutput(ctx *pulumi.Context, args ListMobileNetworkSimIdsOutputArgs, opts ...pulumi.InvokeOption) ListMobileNetworkSimIdsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListMobileNetworkSimIdsResult, error) {
			args := v.(ListMobileNetworkSimIdsArgs)
			r, err := ListMobileNetworkSimIds(ctx, &args, opts...)
			var s ListMobileNetworkSimIdsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListMobileNetworkSimIdsResultOutput)
}

type ListMobileNetworkSimIdsOutputArgs struct {
	MobileNetworkName pulumi.StringInput `pulumi:"mobileNetworkName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListMobileNetworkSimIdsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMobileNetworkSimIdsArgs)(nil)).Elem()
}


type ListMobileNetworkSimIdsResultOutput struct{ *pulumi.OutputState }

func (ListMobileNetworkSimIdsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListMobileNetworkSimIdsResult)(nil)).Elem()
}

func (o ListMobileNetworkSimIdsResultOutput) ToListMobileNetworkSimIdsResultOutput() ListMobileNetworkSimIdsResultOutput {
	return o
}

func (o ListMobileNetworkSimIdsResultOutput) ToListMobileNetworkSimIdsResultOutputWithContext(ctx context.Context) ListMobileNetworkSimIdsResultOutput {
	return o
}

func (o ListMobileNetworkSimIdsResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListMobileNetworkSimIdsResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListMobileNetworkSimIdsResultOutput) Value() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ListMobileNetworkSimIdsResult) []SubResourceResponse { return v.Value }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListMobileNetworkSimIdsResultOutput{})
}
