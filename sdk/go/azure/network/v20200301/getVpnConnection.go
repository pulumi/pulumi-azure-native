


package v20200301

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnConnection(ctx *pulumi.Context, args *LookupVpnConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionResult, error) {
	var rv LookupVpnConnectionResult
	err := ctx.Invoke("azure-native:network/v20200301:getVpnConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVpnConnectionResult struct {
	ConnectionBandwidth            *int                            `pulumi:"connectionBandwidth"`
	ConnectionStatus               string                          `pulumi:"connectionStatus"`
	DpdTimeoutSeconds              *int                            `pulumi:"dpdTimeoutSeconds"`
	EgressBytesTransferred         float64                         `pulumi:"egressBytesTransferred"`
	EnableBgp                      *bool                           `pulumi:"enableBgp"`
	EnableInternetSecurity         *bool                           `pulumi:"enableInternetSecurity"`
	EnableRateLimiting             *bool                           `pulumi:"enableRateLimiting"`
	Etag                           string                          `pulumi:"etag"`
	Id                             *string                         `pulumi:"id"`
	IngressBytesTransferred        float64                         `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  []IpsecPolicyResponse           `pulumi:"ipsecPolicies"`
	Name                           *string                         `pulumi:"name"`
	ProvisioningState              string                          `pulumi:"provisioningState"`
	RemoteVpnSite                  *SubResourceResponse            `pulumi:"remoteVpnSite"`
	RoutingWeight                  *int                            `pulumi:"routingWeight"`
	SharedKey                      *string                         `pulumi:"sharedKey"`
	UseLocalAzureIpAddress         *bool                           `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                           `pulumi:"usePolicyBasedTrafficSelectors"`
	VpnConnectionProtocolType      *string                         `pulumi:"vpnConnectionProtocolType"`
	VpnLinkConnections             []VpnSiteLinkConnectionResponse `pulumi:"vpnLinkConnections"`
}
