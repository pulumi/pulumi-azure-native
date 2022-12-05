


package v20221101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupThreatIntelligenceAlertRule(ctx *pulumi.Context, args *LookupThreatIntelligenceAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupThreatIntelligenceAlertRuleResult, error) {
	var rv LookupThreatIntelligenceAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20221101preview:getThreatIntelligenceAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupThreatIntelligenceAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupThreatIntelligenceAlertRuleResult struct {
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
	Techniques            []string           `pulumi:"techniques"`
	Type                  string             `pulumi:"type"`
}

func LookupThreatIntelligenceAlertRuleOutput(ctx *pulumi.Context, args LookupThreatIntelligenceAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupThreatIntelligenceAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupThreatIntelligenceAlertRuleResult, error) {
			args := v.(LookupThreatIntelligenceAlertRuleArgs)
			r, err := LookupThreatIntelligenceAlertRule(ctx, &args, opts...)
			var s LookupThreatIntelligenceAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupThreatIntelligenceAlertRuleResultOutput)
}

type LookupThreatIntelligenceAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupThreatIntelligenceAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceAlertRuleArgs)(nil)).Elem()
}


type LookupThreatIntelligenceAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupThreatIntelligenceAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupThreatIntelligenceAlertRuleResult)(nil)).Elem()
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) ToLookupThreatIntelligenceAlertRuleResultOutput() LookupThreatIntelligenceAlertRuleResultOutput {
	return o
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) ToLookupThreatIntelligenceAlertRuleResultOutputWithContext(ctx context.Context) LookupThreatIntelligenceAlertRuleResultOutput {
	return o
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) SystemDataResponse { return v.SystemData }).(SystemDataResponseOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) []string { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o LookupThreatIntelligenceAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupThreatIntelligenceAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupThreatIntelligenceAlertRuleResultOutput{})
}
