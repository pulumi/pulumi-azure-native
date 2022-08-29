


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupCostAllocationRule(ctx *pulumi.Context, args *LookupCostAllocationRuleArgs, opts ...pulumi.InvokeOption) (*LookupCostAllocationRuleResult, error) {
	var rv LookupCostAllocationRuleResult
	err := ctx.Invoke("azure-native:costmanagement:getCostAllocationRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupCostAllocationRuleArgs struct {
	BillingAccountId string `pulumi:"billingAccountId"`
	RuleName         string `pulumi:"ruleName"`
}


type LookupCostAllocationRuleResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties CostAllocationRulePropertiesResponse `pulumi:"properties"`
	Type       string                               `pulumi:"type"`
}

func LookupCostAllocationRuleOutput(ctx *pulumi.Context, args LookupCostAllocationRuleOutputArgs, opts ...pulumi.InvokeOption) LookupCostAllocationRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupCostAllocationRuleResult, error) {
			args := v.(LookupCostAllocationRuleArgs)
			r, err := LookupCostAllocationRule(ctx, &args, opts...)
			var s LookupCostAllocationRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupCostAllocationRuleResultOutput)
}

type LookupCostAllocationRuleOutputArgs struct {
	BillingAccountId pulumi.StringInput `pulumi:"billingAccountId"`
	RuleName         pulumi.StringInput `pulumi:"ruleName"`
}

func (LookupCostAllocationRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCostAllocationRuleArgs)(nil)).Elem()
}


type LookupCostAllocationRuleResultOutput struct{ *pulumi.OutputState }

func (LookupCostAllocationRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupCostAllocationRuleResult)(nil)).Elem()
}

func (o LookupCostAllocationRuleResultOutput) ToLookupCostAllocationRuleResultOutput() LookupCostAllocationRuleResultOutput {
	return o
}

func (o LookupCostAllocationRuleResultOutput) ToLookupCostAllocationRuleResultOutputWithContext(ctx context.Context) LookupCostAllocationRuleResultOutput {
	return o
}

func (o LookupCostAllocationRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupCostAllocationRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupCostAllocationRuleResultOutput) Properties() CostAllocationRulePropertiesResponseOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) CostAllocationRulePropertiesResponse { return v.Properties }).(CostAllocationRulePropertiesResponseOutput)
}

func (o LookupCostAllocationRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupCostAllocationRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupCostAllocationRuleResultOutput{})
}
