


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupManagementGroupGovernanceRule(ctx *pulumi.Context, args *LookupManagementGroupGovernanceRuleArgs, opts ...pulumi.InvokeOption) (*LookupManagementGroupGovernanceRuleResult, error) {
	var rv LookupManagementGroupGovernanceRuleResult
	err := ctx.Invoke("azure-native:security/v20220101preview:getManagementGroupGovernanceRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupManagementGroupGovernanceRuleArgs struct {
	ManagementGroupId string `pulumi:"managementGroupId"`
	RuleId            string `pulumi:"ruleId"`
}


type LookupManagementGroupGovernanceRuleResult struct {
	Description                 *string                                  `pulumi:"description"`
	DisplayName                 string                                   `pulumi:"displayName"`
	ExcludedScopes              []string                                 `pulumi:"excludedScopes"`
	GovernanceEmailNotification *GovernanceRuleEmailNotificationResponse `pulumi:"governanceEmailNotification"`
	Id                          string                                   `pulumi:"id"`
	IncludeMemberScopes         *bool                                    `pulumi:"includeMemberScopes"`
	IsDisabled                  *bool                                    `pulumi:"isDisabled"`
	IsGracePeriod               *bool                                    `pulumi:"isGracePeriod"`
	Metadata                    *GovernanceRuleMetadataResponse          `pulumi:"metadata"`
	Name                        string                                   `pulumi:"name"`
	OwnerSource                 GovernanceRuleOwnerSourceResponse        `pulumi:"ownerSource"`
	RemediationTimeframe        *string                                  `pulumi:"remediationTimeframe"`
	RulePriority                int                                      `pulumi:"rulePriority"`
	RuleType                    string                                   `pulumi:"ruleType"`
	SourceResourceType          string                                   `pulumi:"sourceResourceType"`
	TenantId                    string                                   `pulumi:"tenantId"`
	Type                        string                                   `pulumi:"type"`
}

func LookupManagementGroupGovernanceRuleOutput(ctx *pulumi.Context, args LookupManagementGroupGovernanceRuleOutputArgs, opts ...pulumi.InvokeOption) LookupManagementGroupGovernanceRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupManagementGroupGovernanceRuleResult, error) {
			args := v.(LookupManagementGroupGovernanceRuleArgs)
			r, err := LookupManagementGroupGovernanceRule(ctx, &args, opts...)
			var s LookupManagementGroupGovernanceRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupManagementGroupGovernanceRuleResultOutput)
}

type LookupManagementGroupGovernanceRuleOutputArgs struct {
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
}

func (LookupManagementGroupGovernanceRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupGovernanceRuleArgs)(nil)).Elem()
}


type LookupManagementGroupGovernanceRuleResultOutput struct{ *pulumi.OutputState }

func (LookupManagementGroupGovernanceRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupManagementGroupGovernanceRuleResult)(nil)).Elem()
}

func (o LookupManagementGroupGovernanceRuleResultOutput) ToLookupManagementGroupGovernanceRuleResultOutput() LookupManagementGroupGovernanceRuleResultOutput {
	return o
}

func (o LookupManagementGroupGovernanceRuleResultOutput) ToLookupManagementGroupGovernanceRuleResultOutputWithContext(ctx context.Context) LookupManagementGroupGovernanceRuleResultOutput {
	return o
}

func (o LookupManagementGroupGovernanceRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) []string { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *GovernanceRuleEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) IncludeMemberScopes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *bool { return v.IncludeMemberScopes }).(pulumi.BoolPtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) Metadata() GovernanceRuleMetadataResponsePtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *GovernanceRuleMetadataResponse { return v.Metadata }).(GovernanceRuleMetadataResponsePtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) GovernanceRuleOwnerSourceResponse {
		return v.OwnerSource
	}).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) *string { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) int { return v.RulePriority }).(pulumi.IntOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o LookupManagementGroupGovernanceRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupManagementGroupGovernanceRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupManagementGroupGovernanceRuleResultOutput{})
}
