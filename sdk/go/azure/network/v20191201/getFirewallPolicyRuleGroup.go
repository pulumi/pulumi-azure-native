


package v20191201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallPolicyRuleGroup(ctx *pulumi.Context, args *LookupFirewallPolicyRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyRuleGroupResult, error) {
	var rv LookupFirewallPolicyRuleGroupResult
	err := ctx.Invoke("azure-native:network/v20191201:getFirewallPolicyRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyRuleGroupArgs struct {
	FirewallPolicyName string `pulumi:"firewallPolicyName"`
	ResourceGroupName  string `pulumi:"resourceGroupName"`
	RuleGroupName      string `pulumi:"ruleGroupName"`
}


type LookupFirewallPolicyRuleGroupResult struct {
	Etag              string        `pulumi:"etag"`
	Id                *string       `pulumi:"id"`
	Name              *string       `pulumi:"name"`
	Priority          *int          `pulumi:"priority"`
	ProvisioningState string        `pulumi:"provisioningState"`
	Rules             []interface{} `pulumi:"rules"`
	Type              string        `pulumi:"type"`
}

func LookupFirewallPolicyRuleGroupOutput(ctx *pulumi.Context, args LookupFirewallPolicyRuleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicyRuleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicyRuleGroupResult, error) {
			args := v.(LookupFirewallPolicyRuleGroupArgs)
			r, err := LookupFirewallPolicyRuleGroup(ctx, &args, opts...)
			var s LookupFirewallPolicyRuleGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallPolicyRuleGroupResultOutput)
}

type LookupFirewallPolicyRuleGroupOutputArgs struct {
	FirewallPolicyName pulumi.StringInput `pulumi:"firewallPolicyName"`
	ResourceGroupName  pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleGroupName      pulumi.StringInput `pulumi:"ruleGroupName"`
}

func (LookupFirewallPolicyRuleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyRuleGroupArgs)(nil)).Elem()
}


type LookupFirewallPolicyRuleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicyRuleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyRuleGroupResult)(nil)).Elem()
}

func (o LookupFirewallPolicyRuleGroupResultOutput) ToLookupFirewallPolicyRuleGroupResultOutput() LookupFirewallPolicyRuleGroupResultOutput {
	return o
}

func (o LookupFirewallPolicyRuleGroupResultOutput) ToLookupFirewallPolicyRuleGroupResultOutputWithContext(ctx context.Context) LookupFirewallPolicyRuleGroupResultOutput {
	return o
}

func (o LookupFirewallPolicyRuleGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyRuleGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyRuleGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyRuleGroupResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupFirewallPolicyRuleGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyRuleGroupResultOutput) Rules() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) []interface{} { return v.Rules }).(pulumi.ArrayOutput)
}

func (o LookupFirewallPolicyRuleGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicyRuleGroupResultOutput{})
}
