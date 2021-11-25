


package v20190901

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualHub(ctx *pulumi.Context, args *LookupVirtualHubArgs, opts ...pulumi.InvokeOption) (*LookupVirtualHubResult, error) {
	var rv LookupVirtualHubResult
	err := ctx.Invoke("azure-native:network/v20190901:getVirtualHub", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualHubArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	VirtualHubName    string `pulumi:"virtualHubName"`
}


type LookupVirtualHubResult struct {
	AddressPrefix             *string                               `pulumi:"addressPrefix"`
	AzureFirewall             *SubResourceResponse                  `pulumi:"azureFirewall"`
	Etag                      string                                `pulumi:"etag"`
	ExpressRouteGateway       *SubResourceResponse                  `pulumi:"expressRouteGateway"`
	Id                        *string                               `pulumi:"id"`
	Location                  string                                `pulumi:"location"`
	Name                      string                                `pulumi:"name"`
	P2SVpnGateway             *SubResourceResponse                  `pulumi:"p2SVpnGateway"`
	ProvisioningState         string                                `pulumi:"provisioningState"`
	RouteTable                *VirtualHubRouteTableResponse         `pulumi:"routeTable"`
	SecurityProviderName      *string                               `pulumi:"securityProviderName"`
	Sku                       *string                               `pulumi:"sku"`
	Tags                      map[string]string                     `pulumi:"tags"`
	Type                      string                                `pulumi:"type"`
	VirtualHubRouteTableV2s   []VirtualHubRouteTableV2Response      `pulumi:"virtualHubRouteTableV2s"`
	VirtualNetworkConnections []HubVirtualNetworkConnectionResponse `pulumi:"virtualNetworkConnections"`
	VirtualWan                *SubResourceResponse                  `pulumi:"virtualWan"`
	VpnGateway                *SubResourceResponse                  `pulumi:"vpnGateway"`
}
