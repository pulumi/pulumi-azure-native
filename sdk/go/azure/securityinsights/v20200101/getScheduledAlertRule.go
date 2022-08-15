


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupScheduledAlertRule(ctx *pulumi.Context, args *LookupScheduledAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupScheduledAlertRuleResult, error) {
	var rv LookupScheduledAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getScheduledAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupScheduledAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupScheduledAlertRuleResult struct {
	AlertRuleTemplateName *string  `pulumi:"alertRuleTemplateName"`
	Description           *string  `pulumi:"description"`
	DisplayName           string   `pulumi:"displayName"`
	Enabled               bool     `pulumi:"enabled"`
	Etag                  *string  `pulumi:"etag"`
	Id                    string   `pulumi:"id"`
	Kind                  string   `pulumi:"kind"`
	LastModifiedUtc       string   `pulumi:"lastModifiedUtc"`
	Name                  string   `pulumi:"name"`
	Query                 string   `pulumi:"query"`
	QueryFrequency        string   `pulumi:"queryFrequency"`
	QueryPeriod           string   `pulumi:"queryPeriod"`
	Severity              string   `pulumi:"severity"`
	SuppressionDuration   string   `pulumi:"suppressionDuration"`
	SuppressionEnabled    bool     `pulumi:"suppressionEnabled"`
	Tactics               []string `pulumi:"tactics"`
	TriggerOperator       string   `pulumi:"triggerOperator"`
	TriggerThreshold      int      `pulumi:"triggerThreshold"`
	Type                  string   `pulumi:"type"`
}

func LookupScheduledAlertRuleOutput(ctx *pulumi.Context, args LookupScheduledAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupScheduledAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupScheduledAlertRuleResult, error) {
			args := v.(LookupScheduledAlertRuleArgs)
			r, err := LookupScheduledAlertRule(ctx, &args, opts...)
			var s LookupScheduledAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupScheduledAlertRuleResultOutput)
}

type LookupScheduledAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupScheduledAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledAlertRuleArgs)(nil)).Elem()
}


type LookupScheduledAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupScheduledAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupScheduledAlertRuleResult)(nil)).Elem()
}

func (o LookupScheduledAlertRuleResultOutput) ToLookupScheduledAlertRuleResultOutput() LookupScheduledAlertRuleResultOutput {
	return o
}

func (o LookupScheduledAlertRuleResultOutput) ToLookupScheduledAlertRuleResultOutputWithContext(ctx context.Context) LookupScheduledAlertRuleResultOutput {
	return o
}

func (o LookupScheduledAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) *string { return v.AlertRuleTemplateName }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Query() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.Query }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) QueryFrequency() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.QueryFrequency }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) QueryPeriod() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.QueryPeriod }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) SuppressionDuration() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.SuppressionDuration }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) SuppressionEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) bool { return v.SuppressionEnabled }).(pulumi.BoolOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupScheduledAlertRuleResultOutput) TriggerOperator() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.TriggerOperator }).(pulumi.StringOutput)
}

func (o LookupScheduledAlertRuleResultOutput) TriggerThreshold() pulumi.IntOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) int { return v.TriggerThreshold }).(pulumi.IntOutput)
}

func (o LookupScheduledAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupScheduledAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupScheduledAlertRuleResultOutput{})
}
