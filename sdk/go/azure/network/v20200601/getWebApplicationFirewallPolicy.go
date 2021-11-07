


package v20200601

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebApplicationFirewallPolicy(ctx *pulumi.Context, args *LookupWebApplicationFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupWebApplicationFirewallPolicyResult, error) {
	var rv LookupWebApplicationFirewallPolicyResult
	err := ctx.Invoke("azure-native:network/v20200601:getWebApplicationFirewallPolicy", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupWebApplicationFirewallPolicyArgs struct {
	PolicyName        string `pulumi:"policyName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupWebApplicationFirewallPolicyResult struct {
	ApplicationGateways []ApplicationGatewayResponse               `pulumi:"applicationGateways"`
	CustomRules         []WebApplicationFirewallCustomRuleResponse `pulumi:"customRules"`
	Etag                string                                     `pulumi:"etag"`
	HttpListeners       []SubResourceResponse                      `pulumi:"httpListeners"`
	Id                  *string                                    `pulumi:"id"`
	Location            *string                                    `pulumi:"location"`
	ManagedRules        ManagedRulesDefinitionResponse             `pulumi:"managedRules"`
	Name                string                                     `pulumi:"name"`
	PathBasedRules      []SubResourceResponse                      `pulumi:"pathBasedRules"`
	PolicySettings      *PolicySettingsResponse                    `pulumi:"policySettings"`
	ProvisioningState   string                                     `pulumi:"provisioningState"`
	ResourceState       string                                     `pulumi:"resourceState"`
	Tags                map[string]string                          `pulumi:"tags"`
	Type                string                                     `pulumi:"type"`
}
