


package v20180601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGatewayConnection(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayConnectionResult, error) {
	var rv LookupVirtualNetworkGatewayConnectionResult
	err := ctx.Invoke("azure-native:network/v20180601:getVirtualNetworkGatewayConnection", args, &rv, opts...)
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
	ConnectionStatus               string                           `pulumi:"connectionStatus"`
	ConnectionType                 string                           `pulumi:"connectionType"`
	EgressBytesTransferred         float64                          `pulumi:"egressBytesTransferred"`
	EnableBgp                      *bool                            `pulumi:"enableBgp"`
	Etag                           *string                          `pulumi:"etag"`
	Id                             *string                          `pulumi:"id"`
	IngressBytesTransferred        float64                          `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  []IpsecPolicyResponse            `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2           *LocalNetworkGatewayResponse     `pulumi:"localNetworkGateway2"`
	Location                       *string                          `pulumi:"location"`
	Name                           string                           `pulumi:"name"`
	Peer                           *SubResourceResponse             `pulumi:"peer"`
	ProvisioningState              string                           `pulumi:"provisioningState"`
	ResourceGuid                   *string                          `pulumi:"resourceGuid"`
	RoutingWeight                  *int                             `pulumi:"routingWeight"`
	SharedKey                      *string                          `pulumi:"sharedKey"`
	Tags                           map[string]string                `pulumi:"tags"`
	TunnelConnectionStatus         []TunnelConnectionHealthResponse `pulumi:"tunnelConnectionStatus"`
	Type                           string                           `pulumi:"type"`
	UsePolicyBasedTrafficSelectors *bool                            `pulumi:"usePolicyBasedTrafficSelectors"`
	VirtualNetworkGateway1         VirtualNetworkGatewayResponse    `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2         *VirtualNetworkGatewayResponse   `pulumi:"virtualNetworkGateway2"`
}
