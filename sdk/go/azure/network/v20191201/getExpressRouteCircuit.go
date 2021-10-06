


package v20191201

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCircuit(ctx *pulumi.Context, args *LookupExpressRouteCircuitArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitResult, error) {
	var rv LookupExpressRouteCircuitResult
	err := ctx.Invoke("azure-native:network/v20191201:getExpressRouteCircuit", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitArgs struct {
	CircuitName       string `pulumi:"circuitName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteCircuitResult struct {
	AllowClassicOperations           *bool                                                 `pulumi:"allowClassicOperations"`
	Authorizations                   []ExpressRouteCircuitAuthorizationResponse            `pulumi:"authorizations"`
	BandwidthInGbps                  *float64                                              `pulumi:"bandwidthInGbps"`
	CircuitProvisioningState         *string                                               `pulumi:"circuitProvisioningState"`
	Etag                             string                                                `pulumi:"etag"`
	ExpressRoutePort                 *SubResourceResponse                                  `pulumi:"expressRoutePort"`
	GatewayManagerEtag               *string                                               `pulumi:"gatewayManagerEtag"`
	GlobalReachEnabled               *bool                                                 `pulumi:"globalReachEnabled"`
	Id                               *string                                               `pulumi:"id"`
	Location                         *string                                               `pulumi:"location"`
	Name                             string                                                `pulumi:"name"`
	Peerings                         []ExpressRouteCircuitPeeringResponse                  `pulumi:"peerings"`
	ProvisioningState                string                                                `pulumi:"provisioningState"`
	ServiceKey                       *string                                               `pulumi:"serviceKey"`
	ServiceProviderNotes             *string                                               `pulumi:"serviceProviderNotes"`
	ServiceProviderProperties        *ExpressRouteCircuitServiceProviderPropertiesResponse `pulumi:"serviceProviderProperties"`
	ServiceProviderProvisioningState *string                                               `pulumi:"serviceProviderProvisioningState"`
	Sku                              *ExpressRouteCircuitSkuResponse                       `pulumi:"sku"`
	Stag                             int                                                   `pulumi:"stag"`
	Tags                             map[string]string                                     `pulumi:"tags"`
	Type                             string                                                `pulumi:"type"`
}
