


package v20201001

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSubAccountTagRule(ctx *pulumi.Context, args *LookupSubAccountTagRuleArgs, opts ...pulumi.InvokeOption) (*LookupSubAccountTagRuleResult, error) {
	var rv LookupSubAccountTagRuleResult
	err := ctx.Invoke("azure-native:logz/v20201001:getSubAccountTagRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSubAccountTagRuleArgs struct {
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
	SubAccountName    string `pulumi:"subAccountName"`
}


type LookupSubAccountTagRuleResult struct {
	Id         string                               `pulumi:"id"`
	Name       string                               `pulumi:"name"`
	Properties MonitoringTagRulesPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                   `pulumi:"systemData"`
	Type       string                               `pulumi:"type"`
}

func LookupSubAccountTagRuleOutput(ctx *pulumi.Context, args LookupSubAccountTagRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSubAccountTagRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSubAccountTagRuleResult, error) {
			args := v.(LookupSubAccountTagRuleArgs)
			r, err := LookupSubAccountTagRule(ctx, &args, opts...)
			var s LookupSubAccountTagRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSubAccountTagRuleResultOutput)
}

type LookupSubAccountTagRuleOutputArgs struct {
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleSetName       pulumi.StringInput `pulumi:"ruleSetName"`
	SubAccountName    pulumi.StringInput `pulumi:"subAccountName"`
}

func (LookupSubAccountTagRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubAccountTagRuleArgs)(nil)).Elem()
}


type LookupSubAccountTagRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSubAccountTagRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSubAccountTagRuleResult)(nil)).Elem()
}

func (o LookupSubAccountTagRuleResultOutput) ToLookupSubAccountTagRuleResultOutput() LookupSubAccountTagRuleResultOutput {
	return o
}

func (o LookupSubAccountTagRuleResultOutput) ToLookupSubAccountTagRuleResultOutputWithContext(ctx context.Context) LookupSubAccountTagRuleResultOutput {
	return o
}

func (o LookupSubAccountTagRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountTagRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSubAccountTagRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountTagRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSubAccountTagRuleResultOutput) Properties() MonitoringTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupSubAccountTagRuleResult) MonitoringTagRulesPropertiesResponse { return v.Properties }).(MonitoringTagRulesPropertiesResponseOutput)
}

func (o LookupSubAccountTagRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupSubAccountTagRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupSubAccountTagRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSubAccountTagRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSubAccountTagRuleResultOutput{})
}
