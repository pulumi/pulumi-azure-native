


package v20180401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupVpnConnection(ctx *pulumi.Context, args *LookupVpnConnectionArgs, opts ...pulumi.InvokeOption) (*LookupVpnConnectionResult, error) {
	var rv LookupVpnConnectionResult
	err := ctx.Invoke("azure-native:network/v20180401:getVpnConnection", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVpnConnectionArgs struct {
	ConnectionName    string `pulumi:"connectionName"`
	GatewayName       string `pulumi:"gatewayName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupVpnConnectionResult struct {
	ConnectionBandwidth     int                   `pulumi:"connectionBandwidth"`
	ConnectionStatus        string                `pulumi:"connectionStatus"`
	EgressBytesTransferred  float64               `pulumi:"egressBytesTransferred"`
	EnableBgp               *bool                 `pulumi:"enableBgp"`
	Etag                    string                `pulumi:"etag"`
	Id                      *string               `pulumi:"id"`
	IngressBytesTransferred float64               `pulumi:"ingressBytesTransferred"`
	IpsecPolicies           []IpsecPolicyResponse `pulumi:"ipsecPolicies"`
	Name                    *string               `pulumi:"name"`
	ProvisioningState       string                `pulumi:"provisioningState"`
	RemoteVpnSite           *SubResourceResponse  `pulumi:"remoteVpnSite"`
	RoutingWeight           *int                  `pulumi:"routingWeight"`
	SharedKey               *string               `pulumi:"sharedKey"`
}

func LookupVpnConnectionOutput(ctx *pulumi.Context, args LookupVpnConnectionOutputArgs, opts ...pulumi.InvokeOption) LookupVpnConnectionResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVpnConnectionResult, error) {
			args := v.(LookupVpnConnectionArgs)
			r, err := LookupVpnConnection(ctx, &args, opts...)
			var s LookupVpnConnectionResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVpnConnectionResultOutput)
}

type LookupVpnConnectionOutputArgs struct {
	ConnectionName    pulumi.StringInput `pulumi:"connectionName"`
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupVpnConnectionOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionArgs)(nil)).Elem()
}


type LookupVpnConnectionResultOutput struct{ *pulumi.OutputState }

func (LookupVpnConnectionResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVpnConnectionResult)(nil)).Elem()
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutput() LookupVpnConnectionResultOutput {
	return o
}

func (o LookupVpnConnectionResultOutput) ToLookupVpnConnectionResultOutputWithContext(ctx context.Context) LookupVpnConnectionResultOutput {
	return o
}

func (o LookupVpnConnectionResultOutput) ConnectionBandwidth() pulumi.IntOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) int { return v.ConnectionBandwidth }).(pulumi.IntOutput)
}

func (o LookupVpnConnectionResultOutput) ConnectionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.ConnectionStatus }).(pulumi.StringOutput)
}

func (o LookupVpnConnectionResultOutput) EgressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVpnConnectionResult) float64 { return v.EgressBytesTransferred }).(pulumi.Float64Output)
}

func (o LookupVpnConnectionResultOutput) EnableBgp() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *bool { return v.EnableBgp }).(pulumi.BoolPtrOutput)
}

func (o LookupVpnConnectionResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVpnConnectionResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVpnConnectionResultOutput) IngressBytesTransferred() pulumi.Float64Output {
	return o.ApplyT(func(v LookupVpnConnectionResult) float64 { return v.IngressBytesTransferred }).(pulumi.Float64Output)
}

func (o LookupVpnConnectionResultOutput) IpsecPolicies() IpsecPolicyResponseArrayOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) []IpsecPolicyResponse { return v.IpsecPolicies }).(IpsecPolicyResponseArrayOutput)
}

func (o LookupVpnConnectionResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVpnConnectionResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVpnConnectionResultOutput) RemoteVpnSite() SubResourceResponsePtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *SubResourceResponse { return v.RemoteVpnSite }).(SubResourceResponsePtrOutput)
}

func (o LookupVpnConnectionResultOutput) RoutingWeight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *int { return v.RoutingWeight }).(pulumi.IntPtrOutput)
}

func (o LookupVpnConnectionResultOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVpnConnectionResult) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVpnConnectionResultOutput{})
}
