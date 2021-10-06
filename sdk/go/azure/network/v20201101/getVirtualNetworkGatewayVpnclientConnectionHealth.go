


package v20201101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayVpnclientConnectionHealth(ctx *pulumi.Context, args *GetVirtualNetworkGatewayVpnclientConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayVpnclientConnectionHealthResult, error) {
	var rv GetVirtualNetworkGatewayVpnclientConnectionHealthResult
	err := ctx.Invoke("azure-native:network/v20201101:getVirtualNetworkGatewayVpnclientConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayVpnclientConnectionHealthArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayVpnclientConnectionHealthResult struct {
	Value []VpnClientConnectionHealthDetailResponse `pulumi:"value"`
}
