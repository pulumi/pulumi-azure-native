


package v20190601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCrossConnectionPeering(ctx *pulumi.Context, args *LookupExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCrossConnectionPeeringResult, error) {
	var rv LookupExpressRouteCrossConnectionPeeringResult
	err := ctx.Invoke("azure-native:network/v20190601:getExpressRouteCrossConnectionPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCrossConnectionPeeringArgs struct {
	CrossConnectionName string `pulumi:"crossConnectionName"`
	PeeringName         string `pulumi:"peeringName"`
	ResourceGroupName   string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteCrossConnectionPeeringResult struct {
	AzureASN                   int                                           `pulumi:"azureASN"`
	Etag                       string                                        `pulumi:"etag"`
	GatewayManagerEtag         *string                                       `pulumi:"gatewayManagerEtag"`
	Id                         *string                                       `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             *string                                       `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                                       `pulumi:"name"`
	PeerASN                    *float64                                      `pulumi:"peerASN"`
	PeeringType                *string                                       `pulumi:"peeringType"`
	PrimaryAzurePort           string                                        `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                                       `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          string                                        `pulumi:"provisioningState"`
	SecondaryAzurePort         string                                        `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                                       `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                                       `pulumi:"sharedKey"`
	State                      *string                                       `pulumi:"state"`
	VlanId                     *int                                          `pulumi:"vlanId"`
}
