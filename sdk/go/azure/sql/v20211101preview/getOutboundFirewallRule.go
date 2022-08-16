


package v20211101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupOutboundFirewallRule(ctx *pulumi.Context, args *LookupOutboundFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupOutboundFirewallRuleResult, error) {
	var rv LookupOutboundFirewallRuleResult
	err := ctx.Invoke("azure-native:sql/v20211101preview:getOutboundFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupOutboundFirewallRuleArgs struct {
	OutboundRuleFqdn  string `pulumi:"outboundRuleFqdn"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupOutboundFirewallRuleResult struct {
	Id                string `pulumi:"id"`
	Name              string `pulumi:"name"`
	ProvisioningState string `pulumi:"provisioningState"`
	Type              string `pulumi:"type"`
}

func LookupOutboundFirewallRuleOutput(ctx *pulumi.Context, args LookupOutboundFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupOutboundFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupOutboundFirewallRuleResult, error) {
			args := v.(LookupOutboundFirewallRuleArgs)
			r, err := LookupOutboundFirewallRule(ctx, &args, opts...)
			var s LookupOutboundFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupOutboundFirewallRuleResultOutput)
}

type LookupOutboundFirewallRuleOutputArgs struct {
	OutboundRuleFqdn  pulumi.StringInput `pulumi:"outboundRuleFqdn"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupOutboundFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundFirewallRuleArgs)(nil)).Elem()
}


type LookupOutboundFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupOutboundFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupOutboundFirewallRuleResult)(nil)).Elem()
}

func (o LookupOutboundFirewallRuleResultOutput) ToLookupOutboundFirewallRuleResultOutput() LookupOutboundFirewallRuleResultOutput {
	return o
}

func (o LookupOutboundFirewallRuleResultOutput) ToLookupOutboundFirewallRuleResultOutputWithContext(ctx context.Context) LookupOutboundFirewallRuleResultOutput {
	return o
}

func (o LookupOutboundFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupOutboundFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupOutboundFirewallRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupOutboundFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupOutboundFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupOutboundFirewallRuleResultOutput{})
}
