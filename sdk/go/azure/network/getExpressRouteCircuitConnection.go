


package network

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupExpressRouteCircuitConnection(ctx *pulumi.Context, args *LookupExpressRouteCircuitConnectionArgs, opts ...pulumi.InvokeOption) (*LookupExpressRouteCircuitConnectionResult, error) {
	var rv LookupExpressRouteCircuitConnectionResult
	err := ctx.Invoke("azure-native:network:getExpressRouteCircuitConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupExpressRouteCircuitConnectionArgs struct {
	CircuitName       string `pulumi:"circuitName"`
	ConnectionName    string `pulumi:"connectionName"`
	PeeringName       string `pulumi:"peeringName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupExpressRouteCircuitConnectionResult struct {
	AddressPrefix                  *string                              `pulumi:"addressPrefix"`
	AuthorizationKey               *string                              `pulumi:"authorizationKey"`
	CircuitConnectionStatus        string                               `pulumi:"circuitConnectionStatus"`
	Etag                           string                               `pulumi:"etag"`
	ExpressRouteCircuitPeering     *SubResourceResponse                 `pulumi:"expressRouteCircuitPeering"`
	Id                             *string                              `pulumi:"id"`
	Ipv6CircuitConnectionConfig    *Ipv6CircuitConnectionConfigResponse `pulumi:"ipv6CircuitConnectionConfig"`
	Name                           *string                              `pulumi:"name"`
	PeerExpressRouteCircuitPeering *SubResourceResponse                 `pulumi:"peerExpressRouteCircuitPeering"`
	ProvisioningState              string                               `pulumi:"provisioningState"`
	Type                           string                               `pulumi:"type"`
}

func LookupExpressRouteCircuitConnectionOutput(ctx *pulumi.Context, args LookupExpressRouteCircuitConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupExpressRouteCircuitConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupExpressRouteCircuitConnectionResult, error) {
			args := v.(LookupExpressRouteCircuitConnectionArgs)
			r, err := LookupExpressRouteCircuitConnection(ctx, &args, opts...)
			var s LookupExpressRouteCircuitConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupExpressRouteCircuitConnectionResultOutput)
}

type LookupExpressRouteCircuitConnectionOutputArgs struct {
	CircuitName       pulumi.StringInput `pulumi:"circuitName"`
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	PeeringName       pulumi.StringInput `pulumi:"peeringName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupExpressRouteCircuitConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitConnectionArgs)(nil)).Elem()
}


type LookupExpressRouteCircuitConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupExpressRouteCircuitConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupExpressRouteCircuitConnectionResult)(nil)).Elem()
}

func (o LookupExpressRouteCircuitConnectionResultOutput) ToLookupExpressRouteCircuitConnectionResultOutput() LookupExpressRouteCircuitConnectionResultOutput {
	return o
}

func (o LookupExpressRouteCircuitConnectionResultOutput) ToLookupExpressRouteCircuitConnectionResultOutputWithContext(ctx context.Context) LookupExpressRouteCircuitConnectionResultOutput {
	return o
}

func (o LookupExpressRouteCircuitConnectionResultOutput) AddressPrefix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.AddressPrefix }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) CircuitConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.CircuitConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) ExpressRouteCircuitPeering() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *SubResourceResponse {
		return v.ExpressRouteCircuitPeering
	}).(SubResourceResponsePtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) Ipv6CircuitConnectionConfig() Ipv6CircuitConnectionConfigResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *Ipv6CircuitConnectionConfigResponse {
		return v.Ipv6CircuitConnectionConfig
	}).(Ipv6CircuitConnectionConfigResponsePtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) PeerExpressRouteCircuitPeering() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) *SubResourceResponse {
		return v.PeerExpressRouteCircuitPeering
	}).(SubResourceResponsePtrOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupExpressRouteCircuitConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupExpressRouteCircuitConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupExpressRouteCircuitConnectionResultOutput{})
}
