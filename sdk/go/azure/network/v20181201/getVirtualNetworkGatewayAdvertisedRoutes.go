


package v20181201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayAdvertisedRoutes(ctx *pulumi.Context, args *GetVirtualNetworkGatewayAdvertisedRoutesArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayAdvertisedRoutesResult, error) {
	var rv GetVirtualNetworkGatewayAdvertisedRoutesResult
	err := ctx.Invoke("azure-native:network/v20181201:getVirtualNetworkGatewayAdvertisedRoutes", args, &rv, opts...)
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
