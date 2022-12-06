


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SecurityConnectorGovernanceRule struct {
	pulumi.CustomResourceState

	Description                 pulumi.StringPtrOutput                           `pulumi:"description"`
	DisplayName                 pulumi.StringOutput                              `pulumi:"displayName"`
	ExcludedScopes              pulumi.StringArrayOutput                         `pulumi:"excludedScopes"`
	GovernanceEmailNotification GovernanceRuleEmailNotificationResponsePtrOutput `pulumi:"governanceEmailNotification"`
	IncludeMemberScopes         pulumi.BoolPtrOutput                             `pulumi:"includeMemberScopes"`
	IsDisabled                  pulumi.BoolPtrOutput                             `pulumi:"isDisabled"`
	IsGracePeriod               pulumi.BoolPtrOutput                             `pulumi:"isGracePeriod"`
	Metadata                    GovernanceRuleMetadataResponsePtrOutput          `pulumi:"metadata"`
	Name                        pulumi.StringOutput                              `pulumi:"name"`
	OwnerSource                 GovernanceRuleOwnerSourceResponseOutput          `pulumi:"ownerSource"`
	RemediationTimeframe        pulumi.StringPtrOutput                           `pulumi:"remediationTimeframe"`
	RulePriority                pulumi.IntOutput                                 `pulumi:"rulePriority"`
	RuleType                    pulumi.StringOutput                              `pulumi:"ruleType"`
	SourceResourceType          pulumi.StringOutput                              `pulumi:"sourceResourceType"`
	TenantId                    pulumi.StringOutput                              `pulumi:"tenantId"`
	Type                        pulumi.StringOutput                              `pulumi:"type"`
}


func NewSecurityConnectorGovernanceRule(ctx *pulumi.Context,
	name string, args *SecurityConnectorGovernanceRuleArgs, opts ...pulumi.ResourceOption) (*SecurityConnectorGovernanceRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.OwnerSource == nil {
		return nil, errors.New("invalid value for required argument 'OwnerSource'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.RulePriority == nil {
		return nil, errors.New("invalid value for required argument 'RulePriority'")
	}
	if args.RuleType == nil {
		return nil, errors.New("invalid value for required argument 'RuleType'")
	}
	if args.SecurityConnectorName == nil {
		return nil, errors.New("invalid value for required argument 'SecurityConnectorName'")
	}
	if args.SourceResourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceResourceType'")
	}
	var resource SecurityConnectorGovernanceRule
	err := ctx.RegisterResource("azure-native:security/v20220101preview:SecurityConnectorGovernanceRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSecurityConnectorGovernanceRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SecurityConnectorGovernanceRuleState, opts ...pulumi.ResourceOption) (*SecurityConnectorGovernanceRule, error) {
	var resource SecurityConnectorGovernanceRule
	err := ctx.ReadResource("azure-native:security/v20220101preview:SecurityConnectorGovernanceRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type securityConnectorGovernanceRuleState struct {
}

type SecurityConnectorGovernanceRuleState struct {
}

func (SecurityConnectorGovernanceRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorGovernanceRuleState)(nil)).Elem()
}

type securityConnectorGovernanceRuleArgs struct {
	Description                 *string                          `pulumi:"description"`
	DisplayName                 string                           `pulumi:"displayName"`
	ExcludedScopes              []string                         `pulumi:"excludedScopes"`
	GovernanceEmailNotification *GovernanceRuleEmailNotification `pulumi:"governanceEmailNotification"`
	IncludeMemberScopes         *bool                            `pulumi:"includeMemberScopes"`
	IsDisabled                  *bool                            `pulumi:"isDisabled"`
	IsGracePeriod               *bool                            `pulumi:"isGracePeriod"`
	OwnerSource                 GovernanceRuleOwnerSource        `pulumi:"ownerSource"`
	RemediationTimeframe        *string                          `pulumi:"remediationTimeframe"`
	ResourceGroupName           string                           `pulumi:"resourceGroupName"`
	RuleId                      *string                          `pulumi:"ruleId"`
	RulePriority                int                              `pulumi:"rulePriority"`
	RuleType                    string                           `pulumi:"ruleType"`
	SecurityConnectorName       string                           `pulumi:"securityConnectorName"`
	SourceResourceType          string                           `pulumi:"sourceResourceType"`
}


type SecurityConnectorGovernanceRuleArgs struct {
	Description                 pulumi.StringPtrInput
	DisplayName                 pulumi.StringInput
	ExcludedScopes              pulumi.StringArrayInput
	GovernanceEmailNotification GovernanceRuleEmailNotificationPtrInput
	IncludeMemberScopes         pulumi.BoolPtrInput
	IsDisabled                  pulumi.BoolPtrInput
	IsGracePeriod               pulumi.BoolPtrInput
	OwnerSource                 GovernanceRuleOwnerSourceInput
	RemediationTimeframe        pulumi.StringPtrInput
	ResourceGroupName           pulumi.StringInput
	RuleId                      pulumi.StringPtrInput
	RulePriority                pulumi.IntInput
	RuleType                    pulumi.StringInput
	SecurityConnectorName       pulumi.StringInput
	SourceResourceType          pulumi.StringInput
}

func (SecurityConnectorGovernanceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*securityConnectorGovernanceRuleArgs)(nil)).Elem()
}

type SecurityConnectorGovernanceRuleInput interface {
	pulumi.Input

	ToSecurityConnectorGovernanceRuleOutput() SecurityConnectorGovernanceRuleOutput
	ToSecurityConnectorGovernanceRuleOutputWithContext(ctx context.Context) SecurityConnectorGovernanceRuleOutput
}

func (*SecurityConnectorGovernanceRule) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorGovernanceRule)(nil)).Elem()
}

func (i *SecurityConnectorGovernanceRule) ToSecurityConnectorGovernanceRuleOutput() SecurityConnectorGovernanceRuleOutput {
	return i.ToSecurityConnectorGovernanceRuleOutputWithContext(context.Background())
}

func (i *SecurityConnectorGovernanceRule) ToSecurityConnectorGovernanceRuleOutputWithContext(ctx context.Context) SecurityConnectorGovernanceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecurityConnectorGovernanceRuleOutput)
}

type SecurityConnectorGovernanceRuleOutput struct{ *pulumi.OutputState }

func (SecurityConnectorGovernanceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SecurityConnectorGovernanceRule)(nil)).Elem()
}

func (o SecurityConnectorGovernanceRuleOutput) ToSecurityConnectorGovernanceRuleOutput() SecurityConnectorGovernanceRuleOutput {
	return o
}

func (o SecurityConnectorGovernanceRuleOutput) ToSecurityConnectorGovernanceRuleOutputWithContext(ctx context.Context) SecurityConnectorGovernanceRuleOutput {
	return o
}

func (o SecurityConnectorGovernanceRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringArrayOutput { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) GovernanceRuleEmailNotificationResponsePtrOutput {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) IncludeMemberScopes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.BoolPtrOutput { return v.IncludeMemberScopes }).(pulumi.BoolPtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.BoolPtrOutput { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) Metadata() GovernanceRuleMetadataResponsePtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) GovernanceRuleMetadataResponsePtrOutput { return v.Metadata }).(GovernanceRuleMetadataResponsePtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) GovernanceRuleOwnerSourceResponseOutput { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringPtrOutput { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.IntOutput { return v.RulePriority }).(pulumi.IntOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o SecurityConnectorGovernanceRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *SecurityConnectorGovernanceRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(SecurityConnectorGovernanceRuleOutput{})
}
