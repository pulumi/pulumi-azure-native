


package v20200801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFirewallPolicyRuleCollectionGroup(ctx *pulumi.Context, args *LookupFirewallPolicyRuleCollectionGroupArgs, opts ...pulumi.InvokeOption) (*LookupFirewallPolicyRuleCollectionGroupResult, error) {
	var rv LookupFirewallPolicyRuleCollectionGroupResult
	err := ctx.Invoke("azure-native:network/v20200801:getFirewallPolicyRuleCollectionGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFirewallPolicyRuleCollectionGroupArgs struct {
	FirewallPolicyName      string `pulumi:"firewallPolicyName"`
	ResourceGroupName       string `pulumi:"resourceGroupName"`
	RuleCollectionGroupName string `pulumi:"ruleCollectionGroupName"`
}


type LookupFirewallPolicyRuleCollectionGroupResult struct {
	Etag              string        `pulumi:"etag"`
	Id                *string       `pulumi:"id"`
	Name              *string       `pulumi:"name"`
	Priority          *int          `pulumi:"priority"`
	ProvisioningState string        `pulumi:"provisioningState"`
	RuleCollections   []interface{} `pulumi:"ruleCollections"`
	Type              string        `pulumi:"type"`
}

func LookupFirewallPolicyRuleCollectionGroupOutput(ctx *pulumi.Context, args LookupFirewallPolicyRuleCollectionGroupOutputArgs, opts ...pulumi.InvokeOption) LookupFirewallPolicyRuleCollectionGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFirewallPolicyRuleCollectionGroupResult, error) {
			args := v.(LookupFirewallPolicyRuleCollectionGroupArgs)
			r, err := LookupFirewallPolicyRuleCollectionGroup(ctx, &args, opts...)
			var s LookupFirewallPolicyRuleCollectionGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFirewallPolicyRuleCollectionGroupResultOutput)
}

type LookupFirewallPolicyRuleCollectionGroupOutputArgs struct {
	FirewallPolicyName      pulumi.StringInput `pulumi:"firewallPolicyName"`
	ResourceGroupName       pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleCollectionGroupName pulumi.StringInput `pulumi:"ruleCollectionGroupName"`
}

func (LookupFirewallPolicyRuleCollectionGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyRuleCollectionGroupArgs)(nil)).Elem()
}


type LookupFirewallPolicyRuleCollectionGroupResultOutput struct{ *pulumi.OutputState }

func (LookupFirewallPolicyRuleCollectionGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFirewallPolicyRuleCollectionGroupResult)(nil)).Elem()
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ToLookupFirewallPolicyRuleCollectionGroupResultOutput() LookupFirewallPolicyRuleCollectionGroupResultOutput {
	return o
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ToLookupFirewallPolicyRuleCollectionGroupResultOutputWithContext(ctx context.Context) LookupFirewallPolicyRuleCollectionGroupResultOutput {
	return o
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) string { return v.Etag }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) RuleCollections() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) []interface{} { return v.RuleCollections }).(pulumi.ArrayOutput)
}

func (o LookupFirewallPolicyRuleCollectionGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFirewallPolicyRuleCollectionGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFirewallPolicyRuleCollectionGroupResultOutput{})
}
