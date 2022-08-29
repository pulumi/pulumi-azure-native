


package v20190801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupWebApplicationFirewallPolicy(ctx *pulumi.Context, args *LookupWebApplicationFirewallPolicyArgs, opts ...pulumi.InvokeOption) (*LookupWebApplicationFirewallPolicyResult, error) {
	var rv LookupWebApplicationFirewallPolicyResult
	err := ctx.Invoke("azure-native:network/v20190801:getWebApplicationFirewallPolicy", args, &rv, opts...)
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
	Etag                *string                                    `pulumi:"etag"`
	Id                  *string                                    `pulumi:"id"`
	Location            *string                                    `pulumi:"location"`
	ManagedRules        ManagedRulesDefinitionResponse             `pulumi:"managedRules"`
	Name                string                                     `pulumi:"name"`
	PolicySettings      *PolicySettingsResponse                    `pulumi:"policySettings"`
	ProvisioningState   string                                     `pulumi:"provisioningState"`
	ResourceState       string                                     `pulumi:"resourceState"`
	Tags                map[string]string                          `pulumi:"tags"`
	Type                string                                     `pulumi:"type"`
}

func LookupWebApplicationFirewallPolicyOutput(ctx *pulumi.Context, args LookupWebApplicationFirewallPolicyOutputArgs, opts ...pulumi.InvokeOption) LookupWebApplicationFirewallPolicyResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupWebApplicationFirewallPolicyResult, error) {
			args := v.(LookupWebApplicationFirewallPolicyArgs)
			r, err := LookupWebApplicationFirewallPolicy(ctx, &args, opts...)
			var s LookupWebApplicationFirewallPolicyResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupWebApplicationFirewallPolicyResultOutput)
}

type LookupWebApplicationFirewallPolicyOutputArgs struct {
	PolicyName        pulumi.StringInput `pulumi:"policyName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupWebApplicationFirewallPolicyOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebApplicationFirewallPolicyArgs)(nil)).Elem()
}


type LookupWebApplicationFirewallPolicyResultOutput struct{ *pulumi.OutputState }

func (LookupWebApplicationFirewallPolicyResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupWebApplicationFirewallPolicyResult)(nil)).Elem()
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ToLookupWebApplicationFirewallPolicyResultOutput() LookupWebApplicationFirewallPolicyResultOutput {
	return o
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ToLookupWebApplicationFirewallPolicyResultOutputWithContext(ctx context.Context) LookupWebApplicationFirewallPolicyResultOutput {
	return o
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ApplicationGateways() ApplicationGatewayResponseArrayOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) []ApplicationGatewayResponse {
		return v.ApplicationGateways
	}).(ApplicationGatewayResponseArrayOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) CustomRules() WebApplicationFirewallCustomRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) []WebApplicationFirewallCustomRuleResponse {
		return v.CustomRules
	}).(WebApplicationFirewallCustomRuleResponseArrayOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ManagedRules() ManagedRulesDefinitionResponseOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) ManagedRulesDefinitionResponse { return v.ManagedRules }).(ManagedRulesDefinitionResponseOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) PolicySettings() PolicySettingsResponsePtrOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) *PolicySettingsResponse { return v.PolicySettings }).(PolicySettingsResponsePtrOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) ResourceState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.ResourceState }).(pulumi.StringOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupWebApplicationFirewallPolicyResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupWebApplicationFirewallPolicyResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupWebApplicationFirewallPolicyResultOutput{})
}
