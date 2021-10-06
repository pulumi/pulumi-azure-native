


package v20171101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayLearnedRoutes(ctx *pulumi.Context, args *GetVirtualNetworkGatewayLearnedRoutesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayLearnedRoutesResult, error) {
	var rv GetVirtualNetworkGatewayLearnedRoutesResult
	err := ctx.Invoke("azure-native:network/v20171101:getVirtualNetworkGatewayLearnedRoutes", args, &rv, opts...)
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
