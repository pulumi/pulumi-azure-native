


package v20180701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVpnConnection(ctx *pulumi.Context, args *LookupVpnConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionResult, error) {
	var rv LookupVpnConnectionResult
	err := ctx.Invoke("azure-native:network/v20180701:getVpnConnection", args, &rv, opts...)
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
	ConnectionBandwidthInMbps int                   `pulumi:"connectionBandwidthInMbps"`
	ConnectionStatus          string                `pulumi:"connectionStatus"`
	EgressBytesTransferred    float64               `pulumi:"egressBytesTransferred"`
	EnableBgp                 *bool                 `pulumi:"enableBgp"`
	Etag                      string                `pulumi:"etag"`
	Id                        *string               `pulumi:"id"`
	IngressBytesTransferred   float64               `pulumi:"ingressBytesTransferred"`
	IpsecPolicies             []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	Name                      *string               `pulumi:"name"`
	ProvisioningState         string                `pulumi:"provisioningState"`
	RemoteVpnSite             *SubResourceResponse  `pulumi:"remoteVpnSite"`
	RoutingWeight             *int                  `pulumi:"routingWeight"`
	SharedKey                 *string               `pulumi:"sharedKey"`
}
