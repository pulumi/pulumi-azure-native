


package sql

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupIPv6FirewallRule(ctx *pulumi.Context, args *LookupIPv6FirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupIPv6FirewallRuleResult, error) {
	var rv LookupIPv6FirewallRuleResult
	err := ctx.Invoke("azure-native:sql:getIPv6FirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupIPv6FirewallRuleArgs struct {
	FirewallRuleName  string `pulumi:"firewallRuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	ServerName        string `pulumi:"serverName"`
}


type LookupIPv6FirewallRuleResult struct {
	EndIPv6Address   *string `pulumi:"endIPv6Address"`
	Id               string  `pulumi:"id"`
	Name             *string `pulumi:"name"`
	StartIPv6Address *string `pulumi:"startIPv6Address"`
	Type             string  `pulumi:"type"`
}

func LookupIPv6FirewallRuleOutput(ctx *pulumi.Context, args LookupIPv6FirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupIPv6FirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupIPv6FirewallRuleResult, error) {
			args := v.(LookupIPv6FirewallRuleArgs)
			r, err := LookupIPv6FirewallRule(ctx, &args, opts...)
			var s LookupIPv6FirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupIPv6FirewallRuleResultOutput)
}

type LookupIPv6FirewallRuleOutputArgs struct {
	FirewallRuleName  pulumi.StringInput `pulumi:"firewallRuleName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	ServerName        pulumi.StringInput `pulumi:"serverName"`
}

func (LookupIPv6FirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPv6FirewallRuleArgs)(nil)).Elem()
}


type LookupIPv6FirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupIPv6FirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupIPv6FirewallRuleResult)(nil)).Elem()
}

func (o LookupIPv6FirewallRuleResultOutput) ToLookupIPv6FirewallRuleResultOutput() LookupIPv6FirewallRuleResultOutput {
	return o
}

func (o LookupIPv6FirewallRuleResultOutput) ToLookupIPv6FirewallRuleResultOutputWithContext(ctx context.Context) LookupIPv6FirewallRuleResultOutput {
	return o
}

func (o LookupIPv6FirewallRuleResultOutput) EndIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) *string { return v.EndIPv6Address }).(pulumi.StringPtrOutput)
}

func (o LookupIPv6FirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupIPv6FirewallRuleResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupIPv6FirewallRuleResultOutput) StartIPv6Address() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) *string { return v.StartIPv6Address }).(pulumi.StringPtrOutput)
}

func (o LookupIPv6FirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupIPv6FirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupIPv6FirewallRuleResultOutput{})
}
