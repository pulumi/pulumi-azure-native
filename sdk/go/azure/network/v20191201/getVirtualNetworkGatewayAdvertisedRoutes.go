


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayAdvertisedRoutes(ctx *pulumi.Context, args *GetVirtualNetworkGatewayAdvertisedRoutesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayAdvertisedRoutesResult, error) {
	var rv GetVirtualNetworkGatewayAdvertisedRoutesResult
	err := ctx.Invoke("azure-native:network/v20191201:getVirtualNetworkGatewayAdvertisedRoutes", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayAdvertisedRoutesArgs struct {
	Peer                      string `pulumi:"peer"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayAdvertisedRoutesResult struct {
	Value []GatewayRouteResponse `pulumi:"value"`
}

func GetVirtualNetworkGatewayAdvertisedRoutesOutput(ctx *pulumi.Context, args GetVirtualNetworkGatewayAdvertisedRoutesOutputArgs, opts ...pulumi.InvokeOption) GetVirtualNetworkGatewayAdvertisedRoutesResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetVirtualNetworkGatewayAdvertisedRoutesResult, error) {
			args := v.(GetVirtualNetworkGatewayAdvertisedRoutesArgs)
			r, err := GetVirtualNetworkGatewayAdvertisedRoutes(ctx, &args, opts...)
			var s GetVirtualNetworkGatewayAdvertisedRoutesResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetVirtualNetworkGatewayAdvertisedRoutesResultOutput)
}

type GetVirtualNetworkGatewayAdvertisedRoutesOutputArgs struct {
	Peer                      pulumi.StringInput `pulumi:"peer"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (GetVirtualNetworkGatewayAdvertisedRoutesOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayAdvertisedRoutesArgs)(nil)).Elem()
}


type GetVirtualNetworkGatewayAdvertisedRoutesResultOutput struct{ *pulumi.OutputState }

func (GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetVirtualNetworkGatewayAdvertisedRoutesResult)(nil)).Elem()
}

func (o GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) ToGetVirtualNetworkGatewayAdvertisedRoutesResultOutput() GetVirtualNetworkGatewayAdvertisedRoutesResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) ToGetVirtualNetworkGatewayAdvertisedRoutesResultOutputWithContext(ctx context.Context) GetVirtualNetworkGatewayAdvertisedRoutesResultOutput {
	return o
}

func (o GetVirtualNetworkGatewayAdvertisedRoutesResultOutput) Value() GatewayRouteResponseArrayOutput {
	return o.ApplyT(func(v GetVirtualNetworkGatewayAdvertisedRoutesResult) []GatewayRouteResponse { return v.Value }).(GatewayRouteResponseArrayOutput)
}

func init() {
	pulumi.RegisterOutputType(GetVirtualNetworkGatewayAdvertisedRoutesResultOutput{})
}
