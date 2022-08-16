


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListDnsResolverByVirtualNetwork(ctx *pulumi.Context, args *ListDnsResolverByVirtualNetworkArgs, opts ...pulumi.InvokeOption) (*ListDnsResolverByVirtualNetworkResult, error) {
	var rv ListDnsResolverByVirtualNetworkResult
	err := ctx.Invoke("azure-native:network:listDnsResolverByVirtualNetwork", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListDnsResolverByVirtualNetworkArgs struct {
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	Top                *int   `pulumi:"top"`
	VirtualNetworkName string `pulumi:"virtualNetworkName"`
}


type ListDnsResolverByVirtualNetworkResult struct {
	NextLink string                `pulumi:"nextLink"`
	Value    []SubResourceResponse `pulumi:"value"`
}

func ListDnsResolverByVirtualNetworkOutput(ctx *pulumi.Context, args ListDnsResolverByVirtualNetworkOutputArgs, opts ...pulumi.InvokeOption) ListDnsResolverByVirtualNetworkResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListDnsResolverByVirtualNetworkResult, error) {
			args := v.(ListDnsResolverByVirtualNetworkArgs)
			r, err := ListDnsResolverByVirtualNetwork(ctx, &args, opts...)
			var s ListDnsResolverByVirtualNetworkResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListDnsResolverByVirtualNetworkResultOutput)
}

type ListDnsResolverByVirtualNetworkOutputArgs struct {
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	Top                pulumi.IntPtrInput `pulumi:"top"`
	VirtualNetworkName pulumi.StringInput `pulumi:"virtualNetworkName"`
}

func (ListDnsResolverByVirtualNetworkOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDnsResolverByVirtualNetworkArgs)(nil)).Elem()
}


type ListDnsResolverByVirtualNetworkResultOutput struct{ *pulumi.OutputState }

func (ListDnsResolverByVirtualNetworkResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListDnsResolverByVirtualNetworkResult)(nil)).Elem()
}

func (o ListDnsResolverByVirtualNetworkResultOutput) ToListDnsResolverByVirtualNetworkResultOutput() ListDnsResolverByVirtualNetworkResultOutput {
	return o
}

func (o ListDnsResolverByVirtualNetworkResultOutput) ToListDnsResolverByVirtualNetworkResultOutputWithContext(ctx context.Context) ListDnsResolverByVirtualNetworkResultOutput {
	return o
}

func (o ListDnsResolverByVirtualNetworkResultOutput) NextLink() pulumi.StringOutput {
	return o.ApplyT(func(v ListDnsResolverByVirtualNetworkResult) string { return v.NextLink }).(pulumi.StringOutput)
}

func (o ListDnsResolverByVirtualNetworkResultOutput) Value() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v ListDnsResolverByVirtualNetworkResult) []SubResourceResponse { return v.Value }).(SubResourceResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(ListDnsResolverByVirtualNetworkResultOutput{})
}
