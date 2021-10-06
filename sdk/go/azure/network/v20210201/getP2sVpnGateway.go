


package v20210201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupP2sVpnGateway(ctx *pulumi.Context, args *LookupP2sVpnGatewayArgs, opts ...pulumi.InvokeOption) (*LookupP2sVpnGatewayResult, error) {
	var rv LookupP2sVpnGatewayResult
	err := ctx.Invoke("azure-native:network/v20210201:getP2sVpnGateway", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupP2sVpnGatewayArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupP2sVpnGatewayResult struct {
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
