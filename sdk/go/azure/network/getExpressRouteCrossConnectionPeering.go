


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCrossConnectionPeering(ctx *pulumi.Context, args *LookupExpressRouteCrossConnectionPeeringArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCrossConnectionPeeringResult, error) {
	var rv LookupExpressRouteCrossConnectionPeeringResult
	err := ctx.Invoke("azure-native:network:getExpressRouteCrossConnectionPeering", args, &rv, opts...)
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
	LastModifiedBy             string                                        `pulumi:"lastModifiedBy"`
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

func LookupExpressRouteCrossConnectionPeeringOutput(ctx *pulumi.Context, args LookupExpressRouteCrossConnectionPeeringOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCrossConnectionPeeringResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCrossConnectionPeeringResult, error) {
			args := v.(LookupExpressRouteCrossConnectionPeeringArgs)
			r, err := LookupExpressRouteCrossConnectionPeering(ctx, &args, opts...)
			var s LookupExpressRouteCrossConnectionPeeringResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCrossConnectionPeeringResultOutput)
}

type LookupExpressRouteCrossConnectionPeeringOutputArgs struct {
	CrossConnectionName pulumi.StringInput `pulumi:"crossConnectionName"`
	PeeringName         pulumi.StringInput `pulumi:"peeringName"`
	ResourceGroupName   pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCrossConnectionPeeringOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCrossConnectionPeeringArgs)(nil)).Elem()
}


type LookupExpressRouteCrossConnectionPeeringResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCrossConnectionPeeringResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCrossConnectionPeeringResult)(nil)).Elem()
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) ToLookupExpressRouteCrossConnectionPeeringResultOutput() LookupExpressRouteCrossConnectionPeeringResultOutput {
	return o
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) ToLookupExpressRouteCrossConnectionPeeringResultOutputWithContext(ctx context.Context) LookupExpressRouteCrossConnectionPeeringResultOutput {
	return o
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) AzureASN() pulumi.IntOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) int { return v.AzureASN }).(pulumi.IntOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Ipv6PeeringConfig() Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *Ipv6ExpressRouteCircuitPeeringConfigResponse {
		return v.Ipv6PeeringConfig
	}).(Ipv6ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) LastModifiedBy() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.LastModifiedBy }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) MicrosoftPeeringConfig() ExpressRouteCircuitPeeringConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *ExpressRouteCircuitPeeringConfigResponse {
		return v.MicrosoftPeeringConfig
	}).(ExpressRouteCircuitPeeringConfigResponsePtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PeerASN() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *float64 { return v.PeerASN }).(pulumi.Float64PtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PeeringType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.PeeringType }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PrimaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.PrimaryAzurePort }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) PrimaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.PrimaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) SecondaryAzurePort() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) string { return v.SecondaryAzurePort }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) SecondaryPeerAddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.SecondaryPeerAddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *string { return v.State }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCrossConnectionPeeringResultOutput) VlanId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCrossConnectionPeeringResult) *int { return v.VlanId }).(pulumi.IntPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCrossConnectionPeeringResultOutput{})
}
