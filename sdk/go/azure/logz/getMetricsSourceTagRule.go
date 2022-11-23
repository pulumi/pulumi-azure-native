


package logz

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMetricsSourceTagRule(ctx *pulumi.Context, args *LookupMetricsSourceTagRuleArgs, opts ...pulumi.InvokeOption) (*LookupMetricsSourceTagRuleResult, error) {
	var rv LookupMetricsSourceTagRuleResult
	err := ctx.Invoke("azure-native:logz:getMetricsSourceTagRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMetricsSourceTagRuleArgs struct {
	MetricsSourceName string `pulumi:"metricsSourceName"`
	MonitorName       string `pulumi:"monitorName"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleSetName       string `pulumi:"ruleSetName"`
}


type LookupMetricsSourceTagRuleResult struct {
	Id         string                            `pulumi:"id"`
	Name       string                            `pulumi:"name"`
	Properties MetricsTagRulesPropertiesResponse `pulumi:"properties"`
	SystemData SystemDataResponse                `pulumi:"systemData"`
	Type       string                            `pulumi:"type"`
}

func LookupMetricsSourceTagRuleOutput(ctx *pulumi.Context, args LookupMetricsSourceTagRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMetricsSourceTagRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMetricsSourceTagRuleResult, error) {
			args := v.(LookupMetricsSourceTagRuleArgs)
			r, err := LookupMetricsSourceTagRule(ctx, &args, opts...)
			var s LookupMetricsSourceTagRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMetricsSourceTagRuleResultOutput)
}

type LookupMetricsSourceTagRuleOutputArgs struct {
	MetricsSourceName pulumi.StringInput `pulumi:"metricsSourceName"`
	MonitorName       pulumi.StringInput `pulumi:"monitorName"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleSetName       pulumi.StringInput `pulumi:"ruleSetName"`
}

func (LookupMetricsSourceTagRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceTagRuleArgs)(nil)).Elem()
}


type LookupMetricsSourceTagRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMetricsSourceTagRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMetricsSourceTagRuleResult)(nil)).Elem()
}

func (o LookupMetricsSourceTagRuleResultOutput) ToLookupMetricsSourceTagRuleResultOutput() LookupMetricsSourceTagRuleResultOutput {
	return o
}

func (o LookupMetricsSourceTagRuleResultOutput) ToLookupMetricsSourceTagRuleResultOutputWithContext(ctx context.Context) LookupMetricsSourceTagRuleResultOutput {
	return o
}

func (o LookupMetricsSourceTagRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMetricsSourceTagRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMetricsSourceTagRuleResultOutput) Properties() MetricsTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) MetricsTagRulesPropertiesResponse { return v.Properties }).(MetricsTagRulesPropertiesResponseOutput)
}

func (o LookupMetricsSourceTagRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMetricsSourceTagRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMetricsSourceTagRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMetricsSourceTagRuleResultOutput{})
}
