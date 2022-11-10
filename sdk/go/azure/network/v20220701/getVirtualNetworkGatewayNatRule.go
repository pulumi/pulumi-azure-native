


package v20220701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGatewayNatRule(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayNatRuleResult, error) {
	var rv LookupVirtualNetworkGatewayNatRuleResult
	err := ctx.Invoke("azure-native:network/v20220701:getVirtualNetworkGatewayNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupVirtualNetworkGatewayNatRuleArgs struct {
	NatRuleName               string `pulumi:"natRuleName"`
	ResourceGroupName         string `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName string `pulumi:"virtualNetworkGatewayName"`
}


type LookupVirtualNetworkGatewayNatRuleResult struct {
	Etag              string                      `pulumi:"etag"`
	ExternalMappings  []VpnNatRuleMappingResponse `pulumi:"externalMappings"`
	Id                *string                     `pulumi:"id"`
	InternalMappings  []VpnNatRuleMappingResponse `pulumi:"internalMappings"`
	IpConfigurationId *string                     `pulumi:"ipConfigurationId"`
	Mode              *string                     `pulumi:"mode"`
	Name              *string                     `pulumi:"name"`
	ProvisioningState string                      `pulumi:"provisioningState"`
	Type              string                      `pulumi:"type"`
}

func LookupVirtualNetworkGatewayNatRuleOutput(ctx *pulumi.Context, args LookupVirtualNetworkGatewayNatRuleOutputArgs, opts ...pulumi.InvokeOption) LookupVirtualNetworkGatewayNatRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupVirtualNetworkGatewayNatRuleResult, error) {
			args := v.(LookupVirtualNetworkGatewayNatRuleArgs)
			r, err := LookupVirtualNetworkGatewayNatRule(ctx, &args, opts...)
			var s LookupVirtualNetworkGatewayNatRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupVirtualNetworkGatewayNatRuleResultOutput)
}

type LookupVirtualNetworkGatewayNatRuleOutputArgs struct {
	NatRuleName               pulumi.StringInput `pulumi:"natRuleName"`
	ResourceGroupName         pulumi.StringInput `pulumi:"resourceGroupName"`
	VirtualNetworkGatewayName pulumi.StringInput `pulumi:"virtualNetworkGatewayName"`
}

func (LookupVirtualNetworkGatewayNatRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayNatRuleArgs)(nil)).Elem()
}


type LookupVirtualNetworkGatewayNatRuleResultOutput struct{ *pulumi.OutputState }

func (LookupVirtualNetworkGatewayNatRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupVirtualNetworkGatewayNatRuleResult)(nil)).Elem()
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ToLookupVirtualNetworkGatewayNatRuleResultOutput() LookupVirtualNetworkGatewayNatRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ToLookupVirtualNetworkGatewayNatRuleResultOutputWithContext(ctx context.Context) LookupVirtualNetworkGatewayNatRuleResultOutput {
	return o
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ExternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) []VpnNatRuleMappingResponse {
		return v.ExternalMappings
	}).(VpnNatRuleMappingResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) InternalMappings() VpnNatRuleMappingResponseArrayOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) []VpnNatRuleMappingResponse {
		return v.InternalMappings
	}).(VpnNatRuleMappingResponseArrayOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) IpConfigurationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.IpConfigurationId }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupVirtualNetworkGatewayNatRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupVirtualNetworkGatewayNatRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupVirtualNetworkGatewayNatRuleResultOutput{})
}
