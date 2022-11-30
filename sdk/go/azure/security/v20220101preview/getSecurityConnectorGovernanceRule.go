


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func LookupSecurityConnectorGovernanceRule(ctx *pulumi.Context, args *LookupSecurityConnectorGovernanceRuleArgs, opts ...pulumi.InvokeOption) (*LookupSecurityConnectorGovernanceRuleResult, error) {
	var rv LookupSecurityConnectorGovernanceRuleResult
	err := ctx.Invoke("azure-native:security/v20220101preview:getSecurityConnectorGovernanceRule", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type LookupSecurityConnectorGovernanceRuleArgs struct {
	ResourceGroupName     string `pulumi:"resourceGroupName"`
	RuleId                string `pulumi:"ruleId"`
	SecurityConnectorName string `pulumi:"securityConnectorName"`
}


type LookupSecurityConnectorGovernanceRuleResult struct {
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

func LookupSecurityConnectorGovernanceRuleOutput(ctx *pulumi.Context, args LookupSecurityConnectorGovernanceRuleOutputArgs, opts ...pulumi.InvokeOption) LookupSecurityConnectorGovernanceRuleResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (LookupSecurityConnectorGovernanceRuleResult, error) {
			args := v.(LookupSecurityConnectorGovernanceRuleArgs)
			r, err := LookupSecurityConnectorGovernanceRule(ctx, &args, opts...)
			var s LookupSecurityConnectorGovernanceRuleResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(LookupSecurityConnectorGovernanceRuleResultOutput)
}

type LookupSecurityConnectorGovernanceRuleOutputArgs struct {
	ResourceGroupName     pulumi.StringInput `pulumi:"resourceGroupName"`
	RuleId                pulumi.StringInput `pulumi:"ruleId"`
	SecurityConnectorName pulumi.StringInput `pulumi:"securityConnectorName"`
}

func (LookupSecurityConnectorGovernanceRuleOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorGovernanceRuleArgs)(nil)).Elem()
}


type LookupSecurityConnectorGovernanceRuleResultOutput struct{ *pulumi.OutputState }

func (LookupSecurityConnectorGovernanceRuleResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LookupSecurityConnectorGovernanceRuleResult)(nil)).Elem()
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) ToLookupSecurityConnectorGovernanceRuleResultOutput() LookupSecurityConnectorGovernanceRuleResultOutput {
	return o
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) ToLookupSecurityConnectorGovernanceRuleResultOutputWithContext(ctx context.Context) LookupSecurityConnectorGovernanceRuleResultOutput {
	return o
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) *GovernanceRuleEmailNotificationResponse {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) *bool { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) *bool { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) GovernanceRuleOwnerSourceResponse {
		return v.OwnerSource
	}).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) *string { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) int { return v.RulePriority }).(pulumi.IntOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) string { return v.RuleType }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) string { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o LookupSecurityConnectorGovernanceRuleResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v LookupSecurityConnectorGovernanceRuleResult) string { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(LookupSecurityConnectorGovernanceRuleResultOutput{})
}
