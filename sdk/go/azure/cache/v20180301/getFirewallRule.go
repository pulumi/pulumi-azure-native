


package v20180301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupFirewallRule(ctx *pulumi.Context, args *LookupFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupFirewallRuleResult, error) {
	var rv LookupFirewallRuleResult
	err := ctx.Invoke("azure-native:cache/v20180301:getFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallRuleArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupFirewallRuleResult struct {
	EndIP   string `pulumi:"endIP"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	StartIP string `pulumi:"startIP"`
	Type    string `pulumi:"type"`
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
	CacheName         pulumi.StringInput `pulumi:"cacheName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
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

func (o LookupFirewallRuleResultOutput) EndIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.EndIP }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) StartIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.StartIP }).(pulumi.StringOutput)
}

func (o LookupFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallRuleResultOutput{})
}
