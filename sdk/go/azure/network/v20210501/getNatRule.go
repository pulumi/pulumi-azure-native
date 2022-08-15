


package v20210501

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupNatRule(ctx *pulumi.Context, args *LookupNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupNatRuleResult, error) {
	var rv LookupNatRuleResult
	err := ctx.Invoke("azure-native:network/v20210501:getNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupNatRuleArgs struct {
	GatewayName       string `pulumi:"gatewayName"`
	NatRuleName       string `pulumi:"natRuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupNatRuleResult struct {
	EgressVpnSiteLinkConnections  []SubResourceResponse       `pulumi:"egressVpnSiteLinkConnections"`
	Etag                          string                      `pulumi:"etag"`
	ExternalMappings              []VpnNatRuleMappingResponse `pulumi:"externalMappings"`
	Id                            *string                     `pulumi:"id"`
	IngressVpnSiteLinkConnections []SubResourceResponse       `pulumi:"ingressVpnSiteLinkConnections"`
	InternalMappings              []VpnNatRuleMappingResponse `pulumi:"internalMappings"`
	IpConfigurationId             *string                     `pulumi:"ipConfigurationId"`
	Mode                          *string                     `pulumi:"mode"`
	Name                          *string                     `pulumi:"name"`
	ProvisioningState             string                      `pulumi:"provisioningState"`
	Type                          string                      `pulumi:"type"`
}

func LookupNatRuleOutput(ctx *pulumi.Context, args LookupNatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupNatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupNatRuleResult, error) {
			args := v.(LookupNatRuleArgs)
			r, err := LookupNatRule(ctx, &args, opts...)
			var s LookupNatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupNatRuleResultOutput)
}

type LookupNatRuleOutputArgs struct {
	GatewayName       pulumi.StringInput `pulumi:"gatewayName"`
	NatRuleName       pulumi.StringInput `pulumi:"natRuleName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupNatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatRuleArgs)(nil)).Elem()
}


type LookupNatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupNatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupNatRuleResult)(nil)).Elem()
}

func (o LookupNatRuleResultOutput) ToLookupNatRuleResultOutput() LookupNatRuleResultOutput {
	return o
}

func (o LookupNatRuleResultOutput) ToLookupNatRuleResultOutputWithContext(ctx context.Context) LookupNatRuleResultOutput {
	return o
}

func (o LookupNatRuleResultOutput) EgressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatRuleResult) []SubResourceResponse { return v.EgressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o LookupNatRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupNatRuleResultOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupNatRuleResult) []VpnNatRuleMappingResponse { return v.ExternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o LookupNatRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupNatRuleResultOutput) IngressVpnSiteLinkConnections() SubResourceResponseArrayOutput {
	return o.ApplyT(func(v LookupNatRuleResult) []SubResourceResponse { return v.IngressVpnSiteLinkConnections }).(SubResourceResponseArrayOutput)
}

func (o LookupNatRuleResultOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupNatRuleResult) []VpnNatRuleMappingResponse { return v.InternalMappings }).(VpnNatRuleMappingResponseArrayOutput)
}

func (o LookupNatRuleResultOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatRuleResult) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupNatRuleResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatRuleResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupNatRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupNatRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupNatRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupNatRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupNatRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupNatRuleResultOutput{})
}
