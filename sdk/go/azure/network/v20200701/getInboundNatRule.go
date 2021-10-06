


package v20200701

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupInboundNatRule(ctx *pulumi.Context, args *LookupInboundNatRuleArgs, opts ...pulumi.InvokeOption) (*LookupInboundNatRuleResult, error) {
	var rv LookupInboundNatRuleResult
	err := ctx.Invoke("azure-native:network/v20200701:getInboundNatRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupInboundNatRuleArgs struct {
	Expand             *string `pulumi:"expand"`
	InboundNatRuleName string  `pulumi:"inboundNatRuleName"`
	LoadBalancerName   string  `pulumi:"loadBalancerName"`
	ResourceGroupName  string  `pulumi:"resourceGroupName"`
}


type LookupInboundNatRuleResult struct {
	BackendIPConfiguration  NetworkInterfaceIPConfigurationResponse `pulumi:"backendIPConfiguration"`
	BackendPort             *int                                    `pulumi:"backendPort"`
	EnableFloatingIP        *bool                                   `pulumi:"enableFloatingIP"`
	EnableTcpReset          *bool                                   `pulumi:"enableTcpReset"`
	Etag                    string                                  `pulumi:"etag"`
	FrontendIPConfiguration *SubResourceResponse                    `pulumi:"frontendIPConfiguration"`
	FrontendPort            *int                                    `pulumi:"frontendPort"`
	Id                      *string                                 `pulumi:"id"`
	IdleTimeoutInMinutes    *int                                    `pulumi:"idleTimeoutInMinutes"`
	Name                    *string                                 `pulumi:"name"`
	Protocol                *string                                 `pulumi:"protocol"`
	ProvisioningState       string                                  `pulumi:"provisioningState"`
	Type                    string                                  `pulumi:"type"`
}
