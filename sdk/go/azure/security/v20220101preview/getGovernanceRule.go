


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupGovernanceRule(ctx *pulumi.Context, args *LookupGovernanceRuleArgs, opts ...pulumi.InvokeOption) (*LookupGovernanceRuleResult, error) {
	var rv LookupGovernanceRuleResult
	err := ctx.Invoke("azure-native:security/v20220101preview:getGovernanceRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupGovernanceRuleArgs struct {
	RuleId string `pulumi:"ruleId"`
}


type LookupGovernanceRuleResult struct {
	Description                 *string                                  `pulumi:"description"`
	DisplayName                 string                                   `pulumi:"displayName"`
	GovernanceEmailNotification *GovernanceRuleEmailNotificationResponse `pulumi:"governanceEmailNotification"`
	Id                          string                                   `pulumi:"id"`
	IsDisabled                  *bool                                    `pulumi:"isDisabled"`
	IsGracePeriod               *bool                                    `pulumi:"isGracePeriod"`
	Name                        string                                   `pulumi:"name"`
	OwnerSource                 GovernanceRuleOwnerSourceResponse        `pulumi:"ownerSource"`
	RemediationTimeframe        *string                                  `pulumi:"remediationTimeframe"`
	RulePriority                int                                      `pulumi:"rulePriority"`
	RuleType                    string                                   `pulumi:"ruleType"`
	SourceResourceType          string                                   `pulumi:"sourceResourceType"`
	Type                        string                                   `pulumi:"type"`
}

func LookupGovernanceRuleOutput(ctx *pulumi.Context, args LookupGovernanceRuleOutputArgs, opts ...pulumi.InvokeOption) LookupGovernanceRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupGovernanceRuleResult, error) {
			args := v.(LookupGovernanceRuleArgs)
			r, err := LookupGovernanceRule(ctx, &args, opts...)
			var s LookupGovernanceRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupGovernanceRuleResultOutput)
}

type LookupGovernanceRuleOutputArgs struct {
	RuleId pulumi.StringInput `pulumi:"ruleId"`
}

func (LookupGovernanceRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceRuleArgs)(nil)).Elem()
}


type LookupGovernanceRuleResultOutput struct{ *pulumi.OutputState }

func (LookupGovernanceRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupGovernanceRuleResult)(nil)).Elem()
}

func (o LookupGovernanceRuleResultOutput) ToLookupGovernanceRuleResultOutput() LookupGovernanceRuleResultOutput {
	return o
}

func (o LookupGovernanceRuleResultOutput) ToLookupGovernanceRuleResultOutputWithContext(ctx context.Context) LookupGovernanceRuleResultOutput {
	return o
}

func (o LookupGovernanceRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupGovernanceRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupGovernanceRuleResultOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *GovernanceRuleEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o LookupGovernanceRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupGovernanceRuleResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupGovernanceRuleResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o LookupGovernanceRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupGovernanceRuleResultOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) GovernanceRuleOwnerSourceResponse { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o LookupGovernanceRuleResultOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) *string { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o LookupGovernanceRuleResultOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) int { return v.RulePriority }).(pulumi.IntOutput)
}

func (o LookupGovernanceRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o LookupGovernanceRuleResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o LookupGovernanceRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupGovernanceRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupGovernanceRuleResultOutput{})
}
