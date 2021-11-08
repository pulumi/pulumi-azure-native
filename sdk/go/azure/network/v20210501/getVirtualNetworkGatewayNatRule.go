


package v20210501

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupVirtualNetworkGatewayNatRule(ctx *pulumi.Context, args *LookupVirtualNetworkGatewayNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupVirtualNetworkGatewayNatRuleResult, error) {
	var rv LookupVirtualNetworkGatewayNatRuleResult
	err := ctx.Invoke("azure-native:network/v20210501:getVirtualNetworkGatewayNatRule", args, &rv, opts...)
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
