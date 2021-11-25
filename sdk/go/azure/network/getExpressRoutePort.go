


package network

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRoutePort(ctx *pulumi.Context, args *LookupExpressRoutePortArgs, opts ...pulumi.InvokeOption) (*LookupExpressRoutePortResult, error) {
	var rv LookupExpressRoutePortResult
	err := ctx.Invoke("azure-native:network:getExpressRoutePort", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRoutePortArgs struct {
	ExpressRoutePortName string `pulumi:"expressRoutePortName"`
	ResourceGroupName    string `pulumi:"resourceGroupName"`
}


type LookupExpressRoutePortResult struct {
	AllocationDate             string                          `pulumi:"allocationDate"`
	BandwidthInGbps            *int                            `pulumi:"bandwidthInGbps"`
	Circuits                   []SubResourceResponse           `pulumi:"circuits"`
	Encapsulation              *string                         `pulumi:"encapsulation"`
	Etag                       string                          `pulumi:"etag"`
	EtherType                  string                          `pulumi:"etherType"`
	Id                         *string                         `pulumi:"id"`
	Identity                   *ManagedServiceIdentityResponse `pulumi:"identity"`
	Links                      []ExpressRouteLinkResponse      `pulumi:"links"`
	Location                   *string                         `pulumi:"location"`
	Mtu                        string                          `pulumi:"mtu"`
	Name                       string                          `pulumi:"name"`
	PeeringLocation            *string                         `pulumi:"peeringLocation"`
	ProvisionedBandwidthInGbps float64                         `pulumi:"provisionedBandwidthInGbps"`
	ProvisioningState          string                          `pulumi:"provisioningState"`
	ResourceGuid               string                          `pulumi:"resourceGuid"`
	Tags                       map[string]string               `pulumi:"tags"`
	Type                       string                          `pulumi:"type"`
}
