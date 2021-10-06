


package v20190701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetP2sVpnGatewayP2sVpnConnectionHealth(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthResult
	err := ctx.Invoke("azure-native:network/v20190701:getP2sVpnGatewayP2sVpnConnectionHealth", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetP2sVpnGatewayP2sVpnConnectionHealthArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type GetP2sVpnGatewayP2sVpnConnectionHealthResult struct {
	CustomRoutes              *AddressSpaceResponse             `pulumi:"customRoutes"`
	Etag                      string                            `pulumi:"etag"`
	Id                        *string                           `pulumi:"id"`
	Location                  string                            `pulumi:"location"`
	Name                      string                            `pulumi:"name"`
	P2SVpnServerConfiguration *SubResourceResponse              `pulumi:"p2SVpnServerConfiguration"`
	ProvisioningState         string                            `pulumi:"provisioningState"`
	Tags                      map[string]string                 `pulumi:"tags"`
	Type                      string                            `pulumi:"type"`
	VirtualHub                *SubResourceResponse              `pulumi:"virtualHub"`
	VpnClientAddressPool      *AddressSpaceResponse             `pulumi:"vpnClientAddressPool"`
	VpnClientConnectionHealth VpnClientConnectionHealthResponse `pulumi:"vpnClientConnectionHealth"`
	VpnGatewayScaleUnit       *int                              `pulumi:"vpnGatewayScaleUnit"`
}
