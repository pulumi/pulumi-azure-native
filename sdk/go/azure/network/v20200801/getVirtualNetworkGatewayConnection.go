


package v20200801

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGatewayConnection(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayConnectionResult, error) {
	var rv LookupVirtualNetworkGatewayConnectionResult
	err := ctx.Invoke("azure-native:network/v20200801:getVirtualNetworkGatewayConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkGatewayConnectionArgs struct {
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayConnectionName string `pulumi:"virtualNetworkGatewayConnectionName"`
}


type LookupVirtualNetworkGatewayConnectionResult struct {
	AuthorizationKey               *string                          `pulumi:"authorizationKey"`
	ConnectionMode                 *string                          `pulumi:"connectionMode"`
	ConnectionProtocol             *string                          `pulumi:"connectionProtocol"`
	ConnectionStatus               string                           `pulumi:"connectionStatus"`
	ConnectionType                 string                           `pulumi:"connectionType"`
	DpdTimeoutSeconds              *int                             `pulumi:"dpdTimeoutSeconds"`
	EgressBytesTransferred         float64                          `pulumi:"egressBytesTransferred"`
	EnableBgp                      *bool                            `pulumi:"enableBgp"`
	Etag                           string                           `pulumi:"etag"`
	ExpressRouteGatewayBypass      *bool                            `pulumi:"expressRouteGatewayBypass"`
	Id                             *string                          `pulumi:"id"`
	IngressBytesTransferred        float64                          `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  []IpsecPolicyResponse            `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2           *LocalNetworkGatewayResponse     `pulumi:"localNetworkGateway2"`
	Location                       *string                          `pulumi:"location"`
	Name                           string                           `pulumi:"name"`
	Peer                           *SubResourceResponse             `pulumi:"peer"`
	ProvisioningState              string                           `pulumi:"provisioningState"`
	ResourceGuid                   string                           `pulumi:"resourceGuid"`
	RoutingWeight                  *int                             `pulumi:"routingWeight"`
	SharedKey                      *string                          `pulumi:"sharedKey"`
	Tags                           map[string]string                `pulumi:"tags"`
	TrafficSelectorPolicies        []TrafficSelectorPolicyResponse  `pulumi:"trafficSelectorPolicies"`
	TunnelConnectionStatus         []TunnelConnectionHealthResponse `pulumi:"tunnelConnectionStatus"`
	Type                           string                           `pulumi:"type"`
	UseLocalAzureIpAddress         *bool                            `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                            `pulumi:"usePolicyBasedTrafficSelectors"`
	VirtualNetworkGateway1         VirtualNetworkGatewayResponse    `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2         *VirtualNetworkGatewayResponse   `pulumi:"virtualNetworkGateway2"`
}
