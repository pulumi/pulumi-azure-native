


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkPeering(ctx *pulumi.Context, args *LookupVirtualNetworkPeeringArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkPeeringResult, error) {
	var rv LookupVirtualNetworkPeeringResult
	err := ctx.Invoke("azure-native:network/v20210501:getVirtualNetworkPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkPeeringArgs struct {
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkName        string `pulumi:"virtualNetworkName"`
	VirtualNetworkPeeringName string `pulumi:"virtualNetworkPeeringName"`
}


type LookupVirtualNetworkPeeringResult struct {
	AllowForwardedTraffic            *bool                                 `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit              *bool                                 `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess        *bool                                 `pulumi:"allowVirtualNetworkAccess"`
	DoNotVerifyRemoteGateways        *bool                                 `pulumi:"doNotVerifyRemoteGateways"`
	Etag                             string                                `pulumi:"etag"`
	Id                               *string                               `pulumi:"id"`
	Name                             *string                               `pulumi:"name"`
	PeeringState                     *string                               `pulumi:"peeringState"`
	PeeringSyncLevel                 *string                               `pulumi:"peeringSyncLevel"`
	ProvisioningState                string                                `pulumi:"provisioningState"`
	RemoteAddressSpace               *AddressSpaceResponse                 `pulumi:"remoteAddressSpace"`
	RemoteBgpCommunities             *VirtualNetworkBgpCommunitiesResponse `pulumi:"remoteBgpCommunities"`
	RemoteVirtualNetwork             *SubResourceResponse                  `pulumi:"remoteVirtualNetwork"`
	RemoteVirtualNetworkAddressSpace *AddressSpaceResponse                 `pulumi:"remoteVirtualNetworkAddressSpace"`
	RemoteVirtualNetworkEncryption   VirtualNetworkEncryptionResponse      `pulumi:"remoteVirtualNetworkEncryption"`
	ResourceGuid                     string                                `pulumi:"resourceGuid"`
	Type                             *string                               `pulumi:"type"`
	UseRemoteGateways                *bool                                 `pulumi:"useRemoteGateways"`
}
