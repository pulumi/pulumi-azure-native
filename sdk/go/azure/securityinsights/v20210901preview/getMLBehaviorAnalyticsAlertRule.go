


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupMLBehaviorAnalyticsAlertRule(ctx *pulumi.Context, args *LookupMLBehaviorAnalyticsAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupMLBehaviorAnalyticsAlertRuleResult, error) {
	var rv LookupMLBehaviorAnalyticsAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20210901preview:getMLBehaviorAnalyticsAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupMLBehaviorAnalyticsAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupMLBehaviorAnalyticsAlertRuleResult struct {
	AlertRuleTemplateName string             `pulumi:"alertRuleTemplateName"`
	Description           string             `pulumi:"description"`
	DisplayName           string             `pulumi:"displayName"`
	Enabled               bool               `pulumi:"enabled"`
	Etag                  *string            `pulumi:"etag"`
	Id                    string             `pulumi:"id"`
	Kind                  string             `pulumi:"kind"`
	LastModifiedUtc       string             `pulumi:"lastModifiedUtc"`
	Name                  string             `pulumi:"name"`
	Severity              string             `pulumi:"severity"`
	SystemData            SystemDataResponse `pulumi:"systemData"`
	Tactics               []string           `pulumi:"tactics"`
	Type                  string             `pulumi:"type"`
}

func LookupMLBehaviorAnalyticsAlertRuleOutput(ctx *pulumi.Context, args LookupMLBehaviorAnalyticsAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupMLBehaviorAnalyticsAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupMLBehaviorAnalyticsAlertRuleResult, error) {
			args := v.(LookupMLBehaviorAnalyticsAlertRuleArgs)
			r, err := LookupMLBehaviorAnalyticsAlertRule(ctx, &args, opts...)
			var s LookupMLBehaviorAnalyticsAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupMLBehaviorAnalyticsAlertRuleResultOutput)
}

type LookupMLBehaviorAnalyticsAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupMLBehaviorAnalyticsAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMLBehaviorAnalyticsAlertRuleArgs)(nil)).Elem()
}


type LookupMLBehaviorAnalyticsAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupMLBehaviorAnalyticsAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupMLBehaviorAnalyticsAlertRuleResult)(nil)).Elem()
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) ToLookupMLBehaviorAnalyticsAlertRuleResultOutput() LookupMLBehaviorAnalyticsAlertRuleResultOutput {
	return o
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) ToLookupMLBehaviorAnalyticsAlertRuleResultOutputWithContext(ctx context.Context) LookupMLBehaviorAnalyticsAlertRuleResultOutput {
	return o
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupMLBehaviorAnalyticsAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupMLBehaviorAnalyticsAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupMLBehaviorAnalyticsAlertRuleResultOutput{})
}
