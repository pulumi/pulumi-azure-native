


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCircuitPeering(ctx *pulumi.Context, args *LookupExpressRouteCircuitPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitPeeringResult, error) {
	var rv LookupExpressRouteCircuitPeeringResult
	err := ctx.Invoke("azure-native:network/v20220701:getExpressRouteCircuitPeering", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitPeeringArgs struct {
	CircuitName       string `pulumi:"circuitName"`
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteCircuitPeeringResult struct {
	AzureASN                   *int                                          `pulumi:"azureASN"`
	Connections                []ExpressRouteCircuitConnectionResponse       `pulumi:"connections"`
	Etag                       string                                        `pulumi:"etag"`
	ExpressRouteConnection     *ExpressRouteConnectionIdResponse             `pulumi:"expressRouteConnection"`
	GatewayManagerEtag         *string                                       `pulumi:"gatewayManagerEtag"`
	Id                         *string                                       `pulumi:"id"`
	Ipv6PeeringConfig          *Ipv6ExpressRouteCircuitPeeringConfigResponse `pulumi:"ipv6PeeringConfig"`
	LastModifiedBy             string                                        `pulumi:"lastModifiedBy"`
	MicrosoftPeeringConfig     *ExpressRouteCircuitPeeringConfigResponse     `pulumi:"microsoftPeeringConfig"`
	Name                       *string                                       `pulumi:"name"`
	PeerASN                    *float64                                      `pulumi:"peerASN"`
	PeeredConnections          []PeerExpressRouteCircuitConnectionResponse   `pulumi:"peeredConnections"`
	PeeringType                *string                                       `pulumi:"peeringType"`
	PrimaryAzurePort           *string                                       `pulumi:"primaryAzurePort"`
	PrimaryPeerAddressPrefix   *string                                       `pulumi:"primaryPeerAddressPrefix"`
	ProvisioningState          string                                        `pulumi:"provisioningState"`
	RouteFilter                *SubResourceResponse                          `pulumi:"routeFilter"`
	SecondaryAzurePort         *string                                       `pulumi:"secondaryAzurePort"`
	SecondaryPeerAddressPrefix *string                                       `pulumi:"secondaryPeerAddressPrefix"`
	SharedKey                  *string                                       `pulumi:"sharedKey"`
	State                      *string                                       `pulumi:"state"`
	Stats                      *ExpressRouteCircuitStatsResponse             `pulumi:"stats"`
	Type                       string                                        `pulumi:"type"`
	VlanId                     *int                                          `pulumi:"vlanId"`
}

func LookupExpressRouteCircuitPeeringOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitPeeringResult, error) {
			args := v.(LookupExpressRouteCircuitPeeringArgs)
			r, err := LookupExpressRouteCircuitPeering(ctx, &args, opts...)
			var s LookupExpressRouteCircuitPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitPeeringResultOutput)
}

type LookupExpressRouteCircuitPeeringOutputArgs struct {
	CircuitName       pulumi.StringInput `pulumi:"circuitName"`
	PeeringName       pulumi.StringInput `pulumi:"peeringName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitPeeringArgs)(nil)).Elem()
}


type LookupExpressRouteCircuitPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitPeeringResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitPeeringResultOutput) ToLookupExpressRouteCircuitPeeringResultOutput() LookupExpressRouteCircuitPeeringResultOutput {
	return o
}

func (o LookupExpressRouteCircuitPeeringResultOutput) ToLookupExpressRouteCircuitPeeringResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitPeeringResultOutput {
	return o
}

func (o LookupExpressRouteCircuitPeeringResultOutput) AzureASN() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *int { return v.AzureASN }).(pulumi.IntPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Connections() ExpressRouteCircuitConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) []ExpressRouteCircuitConnectionResponse {
		return v.Connections
	}).(ExpressRouteCircuitConnectionResponseArrayOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) ExpressRouteConnection() ExpressRouteConnectionIdResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *ExpressRouteConnectionIdResponse {
		return v.ExpressRouteConnection
	}).(ExpressRouteConnectionIdResponsePtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *float64 { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) PeeredConnections() PeerExpressRouteCircuitConnectionResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) []PeerExpressRouteCircuitConnectionResponse {
		return v.PeeredConnections
	}).(PeerExpressRouteCircuitConnectionResponseArrayOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) PrimaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.PrimaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) RouteFilter() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *SubResourceResponse { return v.RouteFilter }).(SubResourceResponsePtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) SecondaryAzurePort() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.SecondaryAzurePort }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Stats() ExpressRouteCircuitStatsResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *ExpressRouteCircuitStatsResponse { return v.Stats }).(ExpressRouteCircuitStatsResponsePtrOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitPeeringResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitPeeringResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitPeeringResultOutput{})
}
