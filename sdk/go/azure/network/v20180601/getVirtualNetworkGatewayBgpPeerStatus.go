


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetVirtualNetworkGatewayBgpPeerStatus(ctx *pulumi.Context, args *GetVirtualNetworkGatewayBgpPeerStatusArgs, opts ...pulumi.InvokeOption) (*GetVirtualNetworkGatewayBgpPeerStatusResult, error) {
	var rv GetVirtualNetworkGatewayBgpPeerStatusResult
	err := ctx.Invoke("azure-native:network/v20180601:getVirtualNetworkGatewayBgpPeerStatus", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetVirtualNetworkGatewayBgpPeerStatusArgs struct {
	Peer                      *string `pulumi:"peer"`
	ResourceGroupName         string  `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string  `pulumi:"virtualNetworkGatewayName"`
}


type GetVirtualNetworkGatewayBgpPeerStatusResult struct {
	Value []BgpPeerStatusResponse `pulumi:"value"`
}
