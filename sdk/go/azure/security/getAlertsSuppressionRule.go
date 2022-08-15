


package security

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupAlertsSuppressionRule(ctx *pulumi.Context, args *LookupAlertsSuppressionRuleArgs, opts ...pulumi.InvokeOption) (*LookupAlertsSuppressionRuleResult, error) {
	var rv LookupAlertsSuppressionRuleResult
	err := ctx.Invoke("azure-native:security:getAlertsSuppressionRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupAlertsSuppressionRuleArgs struct {
	AlertsSuppressionRuleName string `pulumi:"alertsSuppressionRuleName"`
}


type LookupAlertsSuppressionRuleResult struct {
	AlertType              string                          `pulumi:"alertType"`
	Comment                *string                         `pulumi:"comment"`
	ExpirationDateUtc      *string                         `pulumi:"expirationDateUtc"`
	Id                     string                          `pulumi:"id"`
	LastModifiedUtc        string                          `pulumi:"lastModifiedUtc"`
	Name                   string                          `pulumi:"name"`
	Reason                 string                          `pulumi:"reason"`
	State                  string                          `pulumi:"state"`
	SuppressionAlertsScope *SuppressionAlertsScopeResponse `pulumi:"suppressionAlertsScope"`
	Type                   string                          `pulumi:"type"`
}

func LookupAlertsSuppressionRuleOutput(ctx *pulumi.Context, args LookupAlertsSuppressionRuleOutputArgs, opts ...pulumi.InvokeOption) LookupAlertsSuppressionRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupAlertsSuppressionRuleResult, error) {
			args := v.(LookupAlertsSuppressionRuleArgs)
			r, err := LookupAlertsSuppressionRule(ctx, &args, opts...)
			var s LookupAlertsSuppressionRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupAlertsSuppressionRuleResultOutput)
}

type LookupAlertsSuppressionRuleOutputArgs struct {
	AlertsSuppressionRuleName pulumi.StringInput `pulumi:"alertsSuppressionRuleName"`
}

func (LookupAlertsSuppressionRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertsSuppressionRuleArgs)(nil)).Elem()
}


type LookupAlertsSuppressionRuleResultOutput struct{ *pulumi.OutputState }

func (LookupAlertsSuppressionRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupAlertsSuppressionRuleResult)(nil)).Elem()
}

func (o LookupAlertsSuppressionRuleResultOutput) ToLookupAlertsSuppressionRuleResultOutput() LookupAlertsSuppressionRuleResultOutput {
	return o
}

func (o LookupAlertsSuppressionRuleResultOutput) ToLookupAlertsSuppressionRuleResultOutputWithContext(ctx context.Context) LookupAlertsSuppressionRuleResultOutput {
	return o
}

func (o LookupAlertsSuppressionRuleResultOutput) AlertType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.AlertType }).(pulumi.StringOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) Comment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) *string { return v.Comment }).(pulumi.StringPtrOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) ExpirationDateUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) *string { return v.ExpirationDateUtc }).(pulumi.StringPtrOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) Reason() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Reason }).(pulumi.StringOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.State }).(pulumi.StringOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) SuppressionAlertsScope() SuppressionAlertsScopeResponsePtrOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) *SuppressionAlertsScopeResponse {
		return v.SuppressionAlertsScope
	}).(SuppressionAlertsScopeResponsePtrOutput)
}

func (o LookupAlertsSuppressionRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupAlertsSuppressionRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupAlertsSuppressionRuleResultOutput{})
}
