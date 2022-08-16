


package v20191101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayLearnedRoutes(ctx *pulumi.Context, args *GetVirtualNetworkGatewayLearnedRoutesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayLearnedRoutesResult, error) {
	var rv GetVirtualNetworkGatewayLearnedRoutesResult
	err := ctx.Invoke("azure-native:network/v20191101:getVirtualNetworkGatewayLearnedRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayLearnedRoutesArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayLearnedRoutesResult struct {
	Value []GatewayRouteResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayLearnedRoutesOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayLearnedRoutesOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayLearnedRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayLearnedRoutesResult, error) {
			args := v.(GetVirtualNetworkGatewayLearnedRoutesArgs)
			r, err := GetVirtualNetworkGatewayLearnedRoutes(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayLearnedRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayLearnedRoutesResultOutput)
}

type GetVirtualNetworkGatewayLearnedRoutesOutputArgs struct {
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayLearnedRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayLearnedRoutesArgs)(nil)).Elem()
}


type GetVirtualNetworkGatewayLearnedRoutesResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayLearnedRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayLearnedRoutesResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayLearnedRoutesResultOutput) ToGetVirtualNetworkGatewayLearnedRoutesResultOutput() GetVirtualNetworkGatewayLearnedRoutesResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayLearnedRoutesResultOutput) ToGetVirtualNetworkGatewayLearnedRoutesResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayLearnedRoutesResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayLearnedRoutesResultOutput) Value() GatewayRouteResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayLearnedRoutesResult) []GatewayRouteResponse { return v.Value }).(GatewayRouteResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayLearnedRoutesResultOutput{})
}
