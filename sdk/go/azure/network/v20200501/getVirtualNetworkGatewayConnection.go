


package v20200501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGatewayConnection(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayConnectionResult, error) {
	var rv LookupVirtualNetworkGatewayConnectionResult
	err := ctx.Invoke("azure-native:network/v20200501:getVirtualNetworkGatewayConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkGatewayConnectionArgs struct {
	ResourceGroupName                   string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayConnectionName string `pulumi:"virtualNetworkGatewayConnectionName"`
}


type LookupVirtualNetworkGatewayConnectionResult struct {
	AuthorizationKey               *string                          `pulumi:"authorizationKey"`
	ConnectionProtocol             *string                          `pulumi:"connectionProtocol"`
	ConnectionStatus               string                           `pulumi:"connectionStatus"`
	ConnectionType                 string                           `pulumi:"connectionType"`
	DpdTimeoutSeconds              *int                             `pulumi:"dpdTimeoutSeconds"`
	EgressBytesTransferred         float64                          `pulumi:"egressBytesTransferred"`
	EnableBgp                      *bool                            `pulumi:"enableBgp"`
	Etag                           string                           `pulumi:"etag"`
	ExpressRouteGatewayBypass      *bool                            `pulumi:"expressRouteGatewayBypass"`
	Id                             *string                          `pulumi:"id"`
	IngressBytesTransferred        float64                          `pulumi:"ingressBytesTransferred"`
	IpsecPolicies                  []IpsecPolicyResponse            `pulumi:"ipsecPolicies"`
	LocalNetworkGateway2           *LocalNetworkGatewayResponse     `pulumi:"localNetworkGateway2"`
	Location                       *string                          `pulumi:"location"`
	Name                           string                           `pulumi:"name"`
	Peer                           *SubResourceResponse             `pulumi:"peer"`
	ProvisioningState              string                           `pulumi:"provisioningState"`
	ResourceGuid                   string                           `pulumi:"resourceGuid"`
	RoutingWeight                  *int                             `pulumi:"routingWeight"`
	SharedKey                      *string                          `pulumi:"sharedKey"`
	Tags                           map[string]string                `pulumi:"tags"`
	TrafficSelectorPolicies        []TrafficSelectorPolicyResponse  `pulumi:"trafficSelectorPolicies"`
	TunnelConnectionStatus         []TunnelConnectionHealthResponse `pulumi:"tunnelConnectionStatus"`
	Type                           string                           `pulumi:"type"`
	UseLocalAzureIpAddress         *bool                            `pulumi:"useLocalAzureIpAddress"`
	UsePolicyBasedTrafficSelectors *bool                            `pulumi:"usePolicyBasedTrafficSelectors"`
	VirtualNetworkGateway1         VirtualNetworkGatewayResponse    `pulumi:"virtualNetworkGateway1"`
	VirtualNetworkGateway2         *VirtualNetworkGatewayResponse   `pulumi:"virtualNetworkGateway2"`
}

func LookupVirtualNetworkGatewayConnectionOutput(ctx *pulumi.Context, args LookupVirtualNetworkGatewayConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkGatewayConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkGatewayConnectionResult, error) {
			args := v.(LookupVirtualNetworkGatewayConnectionArgs)
			r, err := LookupVirtualNetworkGatewayConnection(ctx, &args, opts...)
			var s LookupVirtualNetworkGatewayConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkGatewayConnectionResultOutput)
}

type LookupVirtualNetworkGatewayConnectionOutputArgs struct {
	ResourceGroupName                   pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayConnectionName pulumi.StringInput `pulumi:"virtualNetworkGatewayConnectionName"`
}

func (LookupVirtualNetworkGatewayConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayConnectionArgs)(nil)).Elem()
}


type LookupVirtualNetworkGatewayConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkGatewayConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayConnectionResult)(nil)).Elem()
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ToLookupVirtualNetworkGatewayConnectionResultOutput() LookupVirtualNetworkGatewayConnectionResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ToLookupVirtualNetworkGatewayConnectionResultOutputWithContext(ctx context.Context) LookupVirtualNetworkGatewayConnectionResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) AuthorizationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.AuthorizationKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionProtocol() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.ConnectionProtocol }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ConnectionType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ConnectionType }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) DpdTimeoutSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *int { return v.DpdTimeoutSeconds }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ExpressRouteGatewayBypass() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.ExpressRouteGatewayBypass }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) LocalNetworkGateway2() LocalNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *LocalNetworkGatewayResponse {
		return v.LocalNetworkGateway2
	}).(LocalNetworkGatewayResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Peer() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *SubResourceResponse { return v.Peer }).(SubResourceResponsePtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) ResourceGuid() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.ResourceGuid }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) TrafficSelectorPolicies() TrafficSelectorPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []TrafficSelectorPolicyResponse {
		return v.TrafficSelectorPolicies
	}).(TrafficSelectorPolicyResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) TunnelConnectionStatus() TunnelConnectionHealthResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) []TunnelConnectionHealthResponse {
		return v.TunnelConnectionStatus
	}).(TunnelConnectionHealthResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) UseLocalAzureIpAddress() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.UseLocalAzureIpAddress }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) UsePolicyBasedTrafficSelectors() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *bool { return v.UsePolicyBasedTrafficSelectors }).(pulumi.BoolPtrOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) VirtualNetworkGateway1() VirtualNetworkGatewayResponseOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) VirtualNetworkGatewayResponse {
		return v.VirtualNetworkGateway1
	}).(VirtualNetworkGatewayResponseOutput)
}

func (o LookupVirtualNetworkGatewayConnectionResultOutput) VirtualNetworkGateway2() VirtualNetworkGatewayResponsePtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayConnectionResult) *VirtualNetworkGatewayResponse {
		return v.VirtualNetworkGateway2
	}).(VirtualNetworkGatewayResponsePtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkGatewayConnectionResultOutput{})
}
