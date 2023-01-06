


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ManagementGroupGovernanceRule struct {
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


func NewManagementGroupGovernanceRule(ctx *pulumi.Context,
	name string, args *ManagementGroupGovernanceRuleArgs, opts ...pulumi.ResourceOption) (*ManagementGroupGovernanceRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
	}
	if args.ManagementGroupId == nil {
		return nil, errors.New("invalid value for required argument 'ManagementGroupId'")
	}
	if args.OwnerSource == nil {
		return nil, errors.New("invalid value for required argument 'OwnerSource'")
	}
	if args.RulePriority == nil {
		return nil, errors.New("invalid value for required argument 'RulePriority'")
	}
	if args.RuleType == nil {
		return nil, errors.New("invalid value for required argument 'RuleType'")
	}
	if args.SourceResourceType == nil {
		return nil, errors.New("invalid value for required argument 'SourceResourceType'")
	}
	var resource ManagementGroupGovernanceRule
	err := ctx.RegisterResource("azure-native:security/v20220101preview:ManagementGroupGovernanceRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetManagementGroupGovernanceRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ManagementGroupGovernanceRuleState, opts ...pulumi.ResourceOption) (*ManagementGroupGovernanceRule, error) {
	var resource ManagementGroupGovernanceRule
	err := ctx.ReadResource("azure-native:security/v20220101preview:ManagementGroupGovernanceRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type managementGroupGovernanceRuleState struct {
}

type ManagementGroupGovernanceRuleState struct {
}

func (ManagementGroupGovernanceRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupGovernanceRuleState)(nil)).Elem()
}

type managementGroupGovernanceRuleArgs struct {
	Description                 *string                          `pulumi:"description"`
	DisplayName                 string                           `pulumi:"displayName"`
	ExcludedScopes              []string                         `pulumi:"excludedScopes"`
	GovernanceEmailNotification *GovernanceRuleEmailNotification `pulumi:"governanceEmailNotification"`
	IncludeMemberScopes         *bool                            `pulumi:"includeMemberScopes"`
	IsDisabled                  *bool                            `pulumi:"isDisabled"`
	IsGracePeriod               *bool                            `pulumi:"isGracePeriod"`
	ManagementGroupId           string                           `pulumi:"managementGroupId"`
	OwnerSource                 GovernanceRuleOwnerSource        `pulumi:"ownerSource"`
	RemediationTimeframe        *string                          `pulumi:"remediationTimeframe"`
	RuleId                      *string                          `pulumi:"ruleId"`
	RulePriority                int                              `pulumi:"rulePriority"`
	RuleType                    string                           `pulumi:"ruleType"`
	SourceResourceType          string                           `pulumi:"sourceResourceType"`
}


type ManagementGroupGovernanceRuleArgs struct {
	Description                 pulumi.StringPtrInput
	DisplayName                 pulumi.StringInput
	ExcludedScopes              pulumi.StringArrayInput
	GovernanceEmailNotification GovernanceRuleEmailNotificationPtrInput
	IncludeMemberScopes         pulumi.BoolPtrInput
	IsDisabled                  pulumi.BoolPtrInput
	IsGracePeriod               pulumi.BoolPtrInput
	ManagementGroupId           pulumi.StringInput
	OwnerSource                 GovernanceRuleOwnerSourceInput
	RemediationTimeframe        pulumi.StringPtrInput
	RuleId                      pulumi.StringPtrInput
	RulePriority                pulumi.IntInput
	RuleType                    pulumi.StringInput
	SourceResourceType          pulumi.StringInput
}

func (ManagementGroupGovernanceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*managementGroupGovernanceRuleArgs)(nil)).Elem()
}

type ManagementGroupGovernanceRuleInput interface {
	pulumi.Input

	ToManagementGroupGovernanceRuleOutput() ManagementGroupGovernanceRuleOutput
	ToManagementGroupGovernanceRuleOutputWithContext(ctx context.Context) ManagementGroupGovernanceRuleOutput
}

func (*ManagementGroupGovernanceRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupGovernanceRule)(nil)).Elem()
}

func (i *ManagementGroupGovernanceRule) ToManagementGroupGovernanceRuleOutput() ManagementGroupGovernanceRuleOutput {
	return i.ToManagementGroupGovernanceRuleOutputWithContext(context.Background())
}

func (i *ManagementGroupGovernanceRule) ToManagementGroupGovernanceRuleOutputWithContext(ctx context.Context) ManagementGroupGovernanceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagementGroupGovernanceRuleOutput)
}

type ManagementGroupGovernanceRuleOutput struct{ *pulumi.OutputState }

func (ManagementGroupGovernanceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagementGroupGovernanceRule)(nil)).Elem()
}

func (o ManagementGroupGovernanceRuleOutput) ToManagementGroupGovernanceRuleOutput() ManagementGroupGovernanceRuleOutput {
	return o
}

func (o ManagementGroupGovernanceRuleOutput) ToManagementGroupGovernanceRuleOutputWithContext(ctx context.Context) ManagementGroupGovernanceRuleOutput {
	return o
}

func (o ManagementGroupGovernanceRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ManagementGroupGovernanceRuleOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringArrayOutput { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

func (o ManagementGroupGovernanceRuleOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) GovernanceRuleEmailNotificationResponsePtrOutput {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) IncludeMemberScopes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.BoolPtrOutput { return v.IncludeMemberScopes }).(pulumi.BoolPtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.BoolPtrOutput { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) Metadata() GovernanceRuleMetadataResponsePtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) GovernanceRuleMetadataResponsePtrOutput { return v.Metadata }).(GovernanceRuleMetadataResponsePtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ManagementGroupGovernanceRuleOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) GovernanceRuleOwnerSourceResponseOutput { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o ManagementGroupGovernanceRuleOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringPtrOutput { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o ManagementGroupGovernanceRuleOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.IntOutput { return v.RulePriority }).(pulumi.IntOutput)
}

func (o ManagementGroupGovernanceRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

func (o ManagementGroupGovernanceRuleOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringOutput { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o ManagementGroupGovernanceRuleOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagementGroupGovernanceRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ManagementGroupGovernanceRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ManagementGroupGovernanceRuleOutput{})
}
