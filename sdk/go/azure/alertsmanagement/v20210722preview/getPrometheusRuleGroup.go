


package v20210722preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupPrometheusRuleGroup(ctx *pulumi.Context, args *LookupPrometheusRuleGroupArgs, opts ...pulumi.InvokeOption) (*LookupPrometheusRuleGroupResult, error) {
	var rv LookupPrometheusRuleGroupResult
	err := ctx.Invoke("azure-native:alertsmanagement/v20210722preview:getPrometheusRuleGroup", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupPrometheusRuleGroupArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleGroupName     string `pulumi:"ruleGroupName"`
}


type LookupPrometheusRuleGroupResult struct {
	ClusterName *string                  `pulumi:"clusterName"`
	Description *string                  `pulumi:"description"`
	Enabled     *bool                    `pulumi:"enabled"`
	Id          string                   `pulumi:"id"`
	Interval    *string                  `pulumi:"interval"`
	Location    string                   `pulumi:"location"`
	Name        string                   `pulumi:"name"`
	Rules       []PrometheusRuleResponse `pulumi:"rules"`
	Scopes      []string                 `pulumi:"scopes"`
	SystemData  SystemDataResponse       `pulumi:"systemData"`
	Tags        map[string]string        `pulumi:"tags"`
	Type        string                   `pulumi:"type"`
}

func LookupPrometheusRuleGroupOutput(ctx *pulumi.Context, args LookupPrometheusRuleGroupOutputArgs, opts ...pulumi.InvokeOption) LookupPrometheusRuleGroupResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupPrometheusRuleGroupResult, error) {
			args := v.(LookupPrometheusRuleGroupArgs)
			r, err := LookupPrometheusRuleGroup(ctx, &args, opts...)
			var s LookupPrometheusRuleGroupResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupPrometheusRuleGroupResultOutput)
}

type LookupPrometheusRuleGroupOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleGroupName     pulumi.StringInput `pulumi:"ruleGroupName"`
}

func (LookupPrometheusRuleGroupOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrometheusRuleGroupArgs)(nil)).Elem()
}


type LookupPrometheusRuleGroupResultOutput struct{ *pulumi.OutputState }

func (LookupPrometheusRuleGroupResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupPrometheusRuleGroupResult)(nil)).Elem()
}

func (o LookupPrometheusRuleGroupResultOutput) ToLookupPrometheusRuleGroupResultOutput() LookupPrometheusRuleGroupResultOutput {
	return o
}

func (o LookupPrometheusRuleGroupResultOutput) ToLookupPrometheusRuleGroupResultOutputWithContext(ctx context.Context) LookupPrometheusRuleGroupResultOutput {
	return o
}

func (o LookupPrometheusRuleGroupResultOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *string { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) *string { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Location }).(pulumi.StringOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Rules() PrometheusRuleResponseArrayOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) []PrometheusRuleResponse { return v.Rules }).(PrometheusRuleResponseArrayOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) []string { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) map[string]string { return v.Tags }).(pulumi.StringMapOutput)
}

func (o LookupPrometheusRuleGroupResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupPrometheusRuleGroupResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupPrometheusRuleGroupResultOutput{})
}
