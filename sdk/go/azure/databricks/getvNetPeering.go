


package databricks

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetvNetPeering(ctx *pulumi.Context, args *GetvNetPeeringArgs, opts ...pulumi.InvokeOption) (*GetvNetPeeringResult, error) {
	var rv GetvNetPeeringResult
	err := ctx.Invoke("azure-native:databricks:getvNetPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetvNetPeeringArgs struct {
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type GetvNetPeeringResult struct {
	AllowForwardedTraffic     *bool                                                                  `pulumi:"allowForwardedTraffic"`
	AllowGatewayTransit       *bool                                                                  `pulumi:"allowGatewayTransit"`
	AllowVirtualNetworkAccess *bool                                                                  `pulumi:"allowVirtualNetworkAccess"`
	DatabricksAddressSpace    *AddressSpaceResponse                                                  `pulumi:"databricksAddressSpace"`
	DatabricksVirtualNetwork  *VirtualNetworkPeeringPropertiesFormatResponseDatabricksVirtualNetwork `pulumi:"databricksVirtualNetwork"`
	Id                        string                                                                 `pulumi:"id"`
	Name                      string                                                                 `pulumi:"name"`
	PeeringState              string                                                                 `pulumi:"peeringState"`
	ProvisioningState         string                                                                 `pulumi:"provisioningState"`
	RemoteAddressSpace        *AddressSpaceResponse                                                  `pulumi:"remoteAddressSpace"`
	RemoteVirtualNetwork      VirtualNetworkPeeringPropertiesFormatResponseRemoteVirtualNetwork      `pulumi:"remoteVirtualNetwork"`
	Type                      string                                                                 `pulumi:"type"`
	UseRemoteGateways         *bool                                                                  `pulumi:"useRemoteGateways"`
}
