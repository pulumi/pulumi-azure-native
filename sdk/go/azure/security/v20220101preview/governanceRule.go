


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type GovernanceRule struct {
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


func NewGovernanceRule(ctx *pulumi.Context,
	name string, args *GovernanceRuleArgs, opts ...pulumi.ResourceOption) (*GovernanceRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.DisplayName == nil {
		return nil, errors.New("invalid value for required argument 'DisplayName'")
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
	var resource GovernanceRule
	err := ctx.RegisterResource("azure-native:security/v20220101preview:GovernanceRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetGovernanceRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *GovernanceRuleState, opts ...pulumi.ResourceOption) (*GovernanceRule, error) {
	var resource GovernanceRule
	err := ctx.ReadResource("azure-native:security/v20220101preview:GovernanceRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type governanceRuleState struct {
}

type GovernanceRuleState struct {
}

func (GovernanceRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*governanceRuleState)(nil)).Elem()
}

type governanceRuleArgs struct {
	Description                 *string                          `pulumi:"description"`
	DisplayName                 string                           `pulumi:"displayName"`
	ExcludedScopes              []string                         `pulumi:"excludedScopes"`
	GovernanceEmailNotification *GovernanceRuleEmailNotification `pulumi:"governanceEmailNotification"`
	IncludeMemberScopes         *bool                            `pulumi:"includeMemberScopes"`
	IsDisabled                  *bool                            `pulumi:"isDisabled"`
	IsGracePeriod               *bool                            `pulumi:"isGracePeriod"`
	OwnerSource                 GovernanceRuleOwnerSource        `pulumi:"ownerSource"`
	RemediationTimeframe        *string                          `pulumi:"remediationTimeframe"`
	RuleId                      *string                          `pulumi:"ruleId"`
	RulePriority                int                              `pulumi:"rulePriority"`
	RuleType                    string                           `pulumi:"ruleType"`
	SourceResourceType          string                           `pulumi:"sourceResourceType"`
}


type GovernanceRuleArgs struct {
	Description                 pulumi.StringPtrInput
	DisplayName                 pulumi.StringInput
	ExcludedScopes              pulumi.StringArrayInput
	GovernanceEmailNotification GovernanceRuleEmailNotificationPtrInput
	IncludeMemberScopes         pulumi.BoolPtrInput
	IsDisabled                  pulumi.BoolPtrInput
	IsGracePeriod               pulumi.BoolPtrInput
	OwnerSource                 GovernanceRuleOwnerSourceInput
	RemediationTimeframe        pulumi.StringPtrInput
	RuleId                      pulumi.StringPtrInput
	RulePriority                pulumi.IntInput
	RuleType                    pulumi.StringInput
	SourceResourceType          pulumi.StringInput
}

func (GovernanceRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*governanceRuleArgs)(nil)).Elem()
}

type GovernanceRuleInput interface {
	pulumi.Input

	ToGovernanceRuleOutput() GovernanceRuleOutput
	ToGovernanceRuleOutputWithContext(ctx context.Context) GovernanceRuleOutput
}

func (*GovernanceRule) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceRule)(nil)).Elem()
}

func (i *GovernanceRule) ToGovernanceRuleOutput() GovernanceRuleOutput {
	return i.ToGovernanceRuleOutputWithContext(context.Background())
}

func (i *GovernanceRule) ToGovernanceRuleOutputWithContext(ctx context.Context) GovernanceRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GovernanceRuleOutput)
}

type GovernanceRuleOutput struct{ *pulumi.OutputState }

func (GovernanceRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GovernanceRule)(nil)).Elem()
}

func (o GovernanceRuleOutput) ToGovernanceRuleOutput() GovernanceRuleOutput {
	return o
}

func (o GovernanceRuleOutput) ToGovernanceRuleOutputWithContext(ctx context.Context) GovernanceRuleOutput {
	return o
}

func (o GovernanceRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o GovernanceRuleOutput) ExcludedScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringArrayOutput { return v.ExcludedScopes }).(pulumi.StringArrayOutput)
}

func (o GovernanceRuleOutput) GovernanceEmailNotification() GovernanceRuleEmailNotificationResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceRule) GovernanceRuleEmailNotificationResponsePtrOutput {
		return v.GovernanceEmailNotification
	}).(GovernanceRuleEmailNotificationResponsePtrOutput)
}

func (o GovernanceRuleOutput) IncludeMemberScopes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.BoolPtrOutput { return v.IncludeMemberScopes }).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleOutput) IsDisabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.BoolPtrOutput { return v.IsDisabled }).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleOutput) IsGracePeriod() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.BoolPtrOutput { return v.IsGracePeriod }).(pulumi.BoolPtrOutput)
}

func (o GovernanceRuleOutput) Metadata() GovernanceRuleMetadataResponsePtrOutput {
	return o.ApplyT(func(v *GovernanceRule) GovernanceRuleMetadataResponsePtrOutput { return v.Metadata }).(GovernanceRuleMetadataResponsePtrOutput)
}

func (o GovernanceRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o GovernanceRuleOutput) OwnerSource() GovernanceRuleOwnerSourceResponseOutput {
	return o.ApplyT(func(v *GovernanceRule) GovernanceRuleOwnerSourceResponseOutput { return v.OwnerSource }).(GovernanceRuleOwnerSourceResponseOutput)
}

func (o GovernanceRuleOutput) RemediationTimeframe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringPtrOutput { return v.RemediationTimeframe }).(pulumi.StringPtrOutput)
}

func (o GovernanceRuleOutput) RulePriority() pulumi.IntOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.IntOutput { return v.RulePriority }).(pulumi.IntOutput)
}

func (o GovernanceRuleOutput) RuleType() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringOutput { return v.RuleType }).(pulumi.StringOutput)
}

func (o GovernanceRuleOutput) SourceResourceType() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringOutput { return v.SourceResourceType }).(pulumi.StringOutput)
}

func (o GovernanceRuleOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringOutput { return v.TenantId }).(pulumi.StringOutput)
}

func (o GovernanceRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *GovernanceRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GovernanceRuleOutput{})
}
