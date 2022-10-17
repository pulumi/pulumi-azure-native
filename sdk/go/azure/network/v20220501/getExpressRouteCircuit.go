


package v20220501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCircuit(ctx *pulumi.Context, args *LookupExpressRouteCircuitArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitResult, error) {
	var rv LookupExpressRouteCircuitResult
	err := ctx.Invoke("azure-native:network/v20220501:getExpressRouteCircuit", args, &rv, opts...)
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
	AuthorizationKey                 *string                                               `pulumi:"authorizationKey"`
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

func LookupExpressRouteCircuitOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitResult, error) {
			args := v.(LookupExpressRouteCircuitArgs)
			r, err := LookupExpressRouteCircuit(ctx, &args, opts...)
			var s LookupExpressRouteCircuitResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitResultOutput)
}

type LookupExpressRouteCircuitOutputArgs struct {
	CircuitName       pulumi.StringInput `pulumi:"circuitName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitArgs)(nil)).Elem()
}


type LookupExpressRouteCircuitResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitResultOutput) ToLookupExpressRouteCircuitResultOutput() LookupExpressRouteCircuitResultOutput {
	return o
}

func (o LookupExpressRouteCircuitResultOutput) ToLookupExpressRouteCircuitResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitResultOutput {
	return o
}

func (o LookupExpressRouteCircuitResultOutput) AllowClassicOperations() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *bool { return v.AllowClassicOperations }).(pulumi.BoolPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Authorizations() ExpressRouteCircuitAuthorizationResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) []ExpressRouteCircuitAuthorizationResponse {
		return v.Authorizations
	}).(ExpressRouteCircuitAuthorizationResponseArrayOutput)
}

func (o LookupExpressRouteCircuitResultOutput) BandwidthInGbps() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *float64 { return v.BandwidthInGbps }).(pulumi.Float64PtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) CircuitProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.CircuitProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitResultOutput) ExpressRoutePort() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *SubResourceResponse { return v.ExpressRoutePort }).(SubResourceResponsePtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) GatewayManagerEtag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.GatewayManagerEtag }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) GlobalReachEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *bool { return v.GlobalReachEnabled }).(pulumi.BoolPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Peerings() ExpressRouteCircuitPeeringResponseArrayOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) []ExpressRouteCircuitPeeringResponse { return v.Peerings }).(ExpressRouteCircuitPeeringResponseArrayOutput)
}

func (o LookupExpressRouteCircuitResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitResultOutput) ServiceKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.ServiceKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) ServiceProviderNotes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.ServiceProviderNotes }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) ServiceProviderProperties() ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *ExpressRouteCircuitServiceProviderPropertiesResponse {
		return v.ServiceProviderProperties
	}).(ExpressRouteCircuitServiceProviderPropertiesResponsePtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) ServiceProviderProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *string { return v.ServiceProviderProvisioningState }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Sku() ExpressRouteCircuitSkuResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) *ExpressRouteCircuitSkuResponse { return v.Sku }).(ExpressRouteCircuitSkuResponsePtrOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Stag() pulumi.IntOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) int { return v.Stag }).(pulumi.IntOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupExpressRouteCircuitResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitResultOutput{})
}
