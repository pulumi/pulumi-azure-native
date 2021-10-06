


package v20201101

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupHubVirtualNetworkConnection(ctx *pulumi.Context, args *LookupHubVirtualNetworkConnectionArgs, opts ...pulumi.InvokeOption) (*LookupHubVirtualNetworkConnectionResult, error) {
	var rv LookupHubVirtualNetworkConnectionResult
	err := ctx.Invoke("azure-native:network/v20201101:getHubVirtualNetworkConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupHubVirtualNetworkConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupHubVirtualNetworkConnectionResult struct {
	AllowHubToRemoteVnetTransit         *bool                         `pulumi:"allowHubToRemoteVnetTransit"`
	AllowRemoteVnetToUseHubVnetGateways *bool                         `pulumi:"allowRemoteVnetToUseHubVnetGateways"`
	EnableInternetSecurity              *bool                         `pulumi:"enableInternetSecurity"`
	Etag                                string                        `pulumi:"etag"`
	Id                                  *string                       `pulumi:"id"`
	Name                                *string                       `pulumi:"name"`
	ProvisioningState                   string                        `pulumi:"provisioningState"`
	RemoteVirtualNetwork                *SubResourceResponse          `pulumi:"remoteVirtualNetwork"`
	RoutingConfiguration                *RoutingConfigurationResponse `pulumi:"routingConfiguration"`
}
