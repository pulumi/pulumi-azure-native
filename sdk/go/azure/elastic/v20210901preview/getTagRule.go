


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupTagRule(ctx *pulumi.Context, args *LookupTagRuleArgs, opts ...pulumi.InvokeOption) (*LookupTagRuleResult, error) {
	var rv LookupTagRuleResult
	err := ctx.Invoke("azure-native:elastic/v20210901preview:getTagRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupTagRuleArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupTagRuleResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties MonitoringTagRulesPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                   `pulumi:"systemData"`
	Type       string                               `pulumi:"type"`
}

func LookupTagRuleOutput(ctx *pulumi.Context, args LookupTagRuleOutputArgs, opts ...pulumi.InvokeOption) LookupTagRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupTagRuleResult, error) {
			args := v.(LookupTagRuleArgs)
			r, err := LookupTagRule(ctx, &args, opts...)
			var s LookupTagRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupTagRuleResultOutput)
}

type LookupTagRuleOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleSetName       pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupTagRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagRuleArgs)(nil)).Elem()
}


type LookupTagRuleResultOutput struct{ *pulumi.OutputState }

func (LookupTagRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupTagRuleResult)(nil)).Elem()
}

func (o LookupTagRuleResultOutput) ToLookupTagRuleResultOutput() LookupTagRuleResultOutput {
	return o
}

func (o LookupTagRuleResultOutput) ToLookupTagRuleResultOutputWithContext(ctx context.Context) LookupTagRuleResultOutput {
	return o
}

func (o LookupTagRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupTagRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupTagRuleResultOutput) Properties() MonitoringTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupTagRuleResult) MonitoringTagRulesPropertiesResponse { return v.Properties }).(MonitoringTagRulesPropertiesResponseOutput)
}

func (o LookupTagRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupTagRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupTagRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupTagRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupTagRuleResultOutput{})
}
