


package v20160401

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


func LookupRedisFirewallRule(ctx *pulumi.Context, args *LookupRedisFirewallRuleArgs, opts ...pulumi.InvokeOption) (*LookupRedisFirewallRuleResult, error) {
	var rv LookupRedisFirewallRuleResult
	err := ctx.Invoke("azure-native:cache/v20160401:getRedisFirewallRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRedisFirewallRuleArgs struct {
	CacheName         string `pulumi:"cacheName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
}


type LookupRedisFirewallRuleResult struct {
	EndIP   string `pulumi:"endIP"`
	Id      string `pulumi:"id"`
	Name    string `pulumi:"name"`
	StartIP string `pulumi:"startIP"`
	Type    string `pulumi:"type"`
}

func LookupRedisFirewallRuleOutput(ctx *pulumi.Context, args LookupRedisFirewallRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRedisFirewallRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRedisFirewallRuleResult, error) {
			args := v.(LookupRedisFirewallRuleArgs)
			r, err := LookupRedisFirewallRule(ctx, &args, opts...)
			var s LookupRedisFirewallRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRedisFirewallRuleResultOutput)
}

type LookupRedisFirewallRuleOutputArgs struct {
	CacheName         pulumi.StringInput `pulumi:"cacheName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupRedisFirewallRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisFirewallRuleArgs)(nil)).Elem()
}


type LookupRedisFirewallRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRedisFirewallRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRedisFirewallRuleResult)(nil)).Elem()
}

func (o LookupRedisFirewallRuleResultOutput) ToLookupRedisFirewallRuleResultOutput() LookupRedisFirewallRuleResultOutput {
	return o
}

func (o LookupRedisFirewallRuleResultOutput) ToLookupRedisFirewallRuleResultOutputWithContext(ctx context.Context) LookupRedisFirewallRuleResultOutput {
	return o
}

func (o LookupRedisFirewallRuleResultOutput) EndIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisFirewallRuleResult) string { return v.EndIP }).(pulumi.StringOutput)
}

func (o LookupRedisFirewallRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisFirewallRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRedisFirewallRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisFirewallRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRedisFirewallRuleResultOutput) StartIP() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisFirewallRuleResult) string { return v.StartIP }).(pulumi.StringOutput)
}

func (o LookupRedisFirewallRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRedisFirewallRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRedisFirewallRuleResultOutput{})
}
