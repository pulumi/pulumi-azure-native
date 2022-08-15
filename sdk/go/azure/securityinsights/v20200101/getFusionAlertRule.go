


package v20200101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupFusionAlertRule(ctx *pulumi.Context, args *LookupFusionAlertRuleArgs, opts ...pulumi.InvokeOption) (*LookupFusionAlertRuleResult, error) {
	var rv LookupFusionAlertRuleResult
	err := ctx.Invoke("azure-native:securityinsights/v20200101:getFusionAlertRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupFusionAlertRuleArgs struct {
	ResourceGroupName string `pulumi:"resourceGroupName"`
	RuleId            string `pulumi:"ruleId"`
	WorkspaceName     string `pulumi:"workspaceName"`
}


type LookupFusionAlertRuleResult struct {
	AlertRuleTemplateName string   `pulumi:"alertRuleTemplateName"`
	Description           string   `pulumi:"description"`
	DisplayName           string   `pulumi:"displayName"`
	Enabled               bool     `pulumi:"enabled"`
	Etag                  *string  `pulumi:"etag"`
	Id                    string   `pulumi:"id"`
	Kind                  string   `pulumi:"kind"`
	LastModifiedUtc       string   `pulumi:"lastModifiedUtc"`
	Name                  string   `pulumi:"name"`
	Severity              string   `pulumi:"severity"`
	Tactics               []string `pulumi:"tactics"`
	Type                  string   `pulumi:"type"`
}

func LookupFusionAlertRuleOutput(ctx *pulumi.Context, args LookupFusionAlertRuleOutputArgs, opts ...pulumi.InvokeOption) LookupFusionAlertRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupFusionAlertRuleResult, error) {
			args := v.(LookupFusionAlertRuleArgs)
			r, err := LookupFusionAlertRule(ctx, &args, opts...)
			var s LookupFusionAlertRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupFusionAlertRuleResultOutput)
}

type LookupFusionAlertRuleOutputArgs struct {
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
	WorkspaceName     pulumi.StringInput `pulumi:"workspaceName"`
}

func (LookupFusionAlertRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFusionAlertRuleArgs)(nil)).Elem()
}


type LookupFusionAlertRuleResultOutput struct{ *pulumi.OutputState }

func (LookupFusionAlertRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupFusionAlertRuleResult)(nil)).Elem()
}

func (o LookupFusionAlertRuleResultOutput) ToLookupFusionAlertRuleResultOutput() LookupFusionAlertRuleResultOutput {
	return o
}

func (o LookupFusionAlertRuleResultOutput) ToLookupFusionAlertRuleResultOutputWithContext(ctx context.Context) LookupFusionAlertRuleResultOutput {
	return o
}

func (o LookupFusionAlertRuleResultOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Description }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LookupFusionAlertRuleResultOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) *string { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o LookupFusionAlertRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Kind }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Severity }).(pulumi.StringOutput)
}

func (o LookupFusionAlertRuleResultOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) []string { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o LookupFusionAlertRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupFusionAlertRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupFusionAlertRuleResultOutput{})
}
