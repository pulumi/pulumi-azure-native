


package datalakestore

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:datalakestore:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	AccountName       string `pulumi:"accountName"`
	FirewallRuleName  string `pulumi:"firewallRuleName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type LookupFirewallRuleResult struct {
	EndIpAddress   string `pulumi:"endIpAddress"`
	Id             string `pulumi:"id"`
	Name           string `pulumi:"name"`
	StartIpAddress string `pulumi:"startIpAddress"`
	Type           string `pulumi:"type"`
}

func LookupFirewallRuleOutput(ctx *pulumi.Context, args LookupFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallRuleResult, error) {
			args := v.(LookupFirewallRuleArgs)
			r, err := LookupFirewallRule(ctx, &args, opts...)
			var s LookupFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallRuleResultOutput)
}

type LookupFirewallRuleOutputArgs struct {
	AccountName       pulumi.StringInput `pulumi:"accountName"`
	FirewallRuleName  pulumi.StringInput `pulumi:"firewallRuleName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (LookupFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleArgs)(nil)).Elem()
}


type LookupFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallRuleResult)(nil)).Elem()
}

func (o LookupFirewallRuleResultOutput) ToLookupFirewallRuleResultOutput() LookupFirewallRuleResultOutput {
	return o
}

func (o LookupFirewallRuleResultOutput) ToLookupFirewallRuleResultOutputWithContext(ctx context.Context) LookupFirewallRuleResultOutput {
	return o
}

func (o LookupFirewallRuleResultOutput) EndIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.EndIpAddress }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) StartIpAddress() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.StartIpAddress }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallRuleResultOutput{})
}
