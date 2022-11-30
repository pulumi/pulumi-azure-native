


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupRule(ctx *pulumi.Context, args *LookupRuleArgs, opts ...pulumi.InvokeOption) (*LookupRuleResult, error) {
	var rv LookupRuleResult
	err := ctx.Invoke("azure-native:cdn/v20221101preview:getRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupRuleArgs struct {
	ProfileName       string `pulumi:"profileName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleName          string `pulumi:"ruleName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupRuleResult struct {
	Actions                 []interface{}      `pulumi:"actions"`
	Conditions              []interface{}      `pulumi:"conditions"`
	DeploymentStatus        string             `pulumi:"deploymentStatus"`
	Id                      string             `pulumi:"id"`
	MatchProcessingBehavior *string            `pulumi:"matchProcessingBehavior"`
	Name                    string             `pulumi:"name"`
	Order                   int                `pulumi:"order"`
	ProvisioningState       string             `pulumi:"provisioningState"`
	RuleSetName             string             `pulumi:"ruleSetName"`
	SystemData              SystemDataResponse `pulumi:"systemData"`
	Type                    string             `pulumi:"type"`
}

func LookupRuleOutput(ctx *pulumi.Context, args LookupRuleOutputArgs, opts ...pulumi.InvokeOption) LookupRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupRuleResult, error) {
			args := v.(LookupRuleArgs)
			r, err := LookupRule(ctx, &args, opts...)
			var s LookupRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupRuleResultOutput)
}

type LookupRuleOutputArgs struct {
	ProfileName       pulumi.StringInput `pulumi:"profileName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleName          pulumi.StringInput `pulumi:"ruleName"`
	RuleSetName       pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleArgs)(nil)).Elem()
}


type LookupRuleResultOutput struct{ *pulumi.OutputState }

func (LookupRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupRuleResult)(nil)).Elem()
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutput() LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) ToLookupRuleResultOutputWithContext(ctx context.Context) LookupRuleResultOutput {
	return o
}

func (o LookupRuleResultOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupRuleResult) []interface{} { return v.Actions }).(pulumi.ArrayOutput)
}

func (o LookupRuleResultOutput) Conditions() pulumi.ArrayOutput {
	return o.ApplyT(func(v LookupRuleResult) []interface{} { return v.Conditions }).(pulumi.ArrayOutput)
}

func (o LookupRuleResultOutput) DeploymentStatus() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.DeploymentStatus }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) MatchProcessingBehavior() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupRuleResult) *string { return v.MatchProcessingBehavior }).(pulumi.StringPtrOutput)
}

func (o LookupRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) Order() pulumi.IntOutput {
	return o.ApplyT(func(v LookupRuleResult) int { return v.Order }).(pulumi.IntOutput)
}

func (o LookupRuleResultOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) RuleSetName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.RuleSetName }).(pulumi.StringOutput)
}

func (o LookupRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupRuleResultOutput{})
}
