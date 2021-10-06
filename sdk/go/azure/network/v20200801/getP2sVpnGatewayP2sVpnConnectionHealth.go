


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetP2sVpnGatewayP2sVpnConnectionHealth(ctx *pulumi.Context, args *GetP2sVpnGatewayP2sVpnConnectionHealthArgs, opts ...pulumi.InvokeOption) (*GetP2sVpnGatewayP2sVpnConnectionHealthResult, error) {
	var rv GetP2sVpnGatewayP2sVpnConnectionHealthResult
	err := ctx.Invoke("azure-native:network/v20200801:getP2sVpnGatewayP2sVpnConnectionHealth", args, &rv, opts...)
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
	CustomDnsServers            []string                             `pulumi:"customDnsServers"`
	Etag                        string                               `pulumi:"etag"`
	Id                          *string                              `pulumi:"id"`
	IsRoutingPreferenceInternet *bool                                `pulumi:"isRoutingPreferenceInternet"`
	Location                    string                               `pulumi:"location"`
	Name                        string                               `pulumi:"name"`
	P2SConnectionConfigurations []P2SConnectionConfigurationResponse `pulumi:"p2SConnectionConfigurations"`
	ProvisioningState           string                               `pulumi:"provisioningState"`
	Tags                        map[string]string                    `pulumi:"tags"`
	Type                        string                               `pulumi:"type"`
	VirtualHub                  *SubResourceResponse                 `pulumi:"virtualHub"`
	VpnClientConnectionHealth   VpnClientConnectionHealthResponse    `pulumi:"vpnClientConnectionHealth"`
	VpnGatewayScaleUnit         *int                                 `pulumi:"vpnGatewayScaleUnit"`
	VpnServerConfiguration      *SubResourceResponse                 `pulumi:"vpnServerConfiguration"`
}
