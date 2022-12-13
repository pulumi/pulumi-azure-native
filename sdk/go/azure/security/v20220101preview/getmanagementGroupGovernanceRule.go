


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetmanagementGroupGovernanceRule(ctx *pulumi.Context, args *GetmanagementGroupGovernanceRuleArgs, opts ...pulumi.InvokeOption) (*GetmanagementGroupGovernanceRuleResult, error) {
	var rv GetmanagementGroupGovernanceRuleResult
	err := ctx.Invoke("azure-native:security/v20220101preview:getmanagementGroupGovernanceRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type GetmanagementGroupGovernanceRuleArgs struct {
	ManagementGroupId string `pulumi:"managementGroupId"`
	RuleId            string `pulumi:"ruleId"`
}


type GetmanagementGroupGovernanceRuleResult struct {
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

func GetmanagementGroupGovernanceRuleOutput(ctx *pulumi.Context, args GetmanagementGroupGovernanceRuleOutputArgs, opts ...pulumi.InvokeOption) GetmanagementGroupGovernanceRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetmanagementGroupGovernanceRuleResult, error) {
			args := v.(GetmanagementGroupGovernanceRuleArgs)
			r, err := GetmanagementGroupGovernanceRule(ctx, &args, opts...)
			var s GetmanagementGroupGovernanceRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetmanagementGroupGovernanceRuleResultOutput)
}

type GetmanagementGroupGovernanceRuleOutputArgs struct {
	ManagementGroupId pulumi.StringInput `pulumi:"managementGroupId"`
	RuleId            pulumi.StringInput `pulumi:"ruleId"`
}

func (GetmanagementGroupGovernanceRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmanagementGroupGovernanceRuleArgs)(nil)).Elem()
}


type GetmanagementGroupGovernanceRuleResultOutput struct{ *pulumi.OutputState }

func (GetmanagementGroupGovernanceRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetmanagementGroupGovernanceRuleResult)(nil)).Elem()
}

func (o GetmanagementGroupGovernanceRuleResultOutput) ToGetmanagementGroupGovernanceRuleResultOutput() GetmanagementGroupGovernanceRuleResultOutput {
	return o
}

func (o GetmanagementGroupGovernanceRuleResultOutput) ToGetmanagementGroupGovernanceRuleResultOutputWithContext(ctx context.Context) GetmanagementGroupGovernanceRuleResultOutput {
	return o
}

func (o GetmanagementGroupGovernanceRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) []string { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *GovernanceRuleEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) IncludeMemberScopes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *bool { return v.IncludeMemberScopes }).(pulumi.BoolPtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) Metadata() GovernanceRuleMetadataResponsePtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *GovernanceRuleMetadataResponse { return v.Metadata }).(GovernanceRuleMetadataResponsePtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) GovernanceRuleOwnerSourceResponse { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) *string { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) int { return v.RulePriority }).(pulumi.IntOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o GetmanagementGroupGovernanceRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v GetmanagementGroupGovernanceRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetmanagementGroupGovernanceRuleResultOutput{})
}
