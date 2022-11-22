


package v20210901preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ThreatIntelligenceAlertRule struct {
	pulumi.CustomResourceState

	AlertRuleTemplateName pulumi.StringOutput      `pulumi:"alertRuleTemplateName"`
	Description           pulumi.StringOutput      `pulumi:"description"`
	DisplayName           pulumi.StringOutput      `pulumi:"displayName"`
	Enabled               pulumi.BoolOutput        `pulumi:"enabled"`
	Etag                  pulumi.StringPtrOutput   `pulumi:"etag"`
	Kind                  pulumi.StringOutput      `pulumi:"kind"`
	LastModifiedUtc       pulumi.StringOutput      `pulumi:"lastModifiedUtc"`
	Name                  pulumi.StringOutput      `pulumi:"name"`
	Severity              pulumi.StringOutput      `pulumi:"severity"`
	SystemData            SystemDataResponseOutput `pulumi:"systemData"`
	Tactics               pulumi.StringArrayOutput `pulumi:"tactics"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewThreatIntelligenceAlertRule(ctx *pulumi.Context,
	name string, args *ThreatIntelligenceAlertRuleArgs, opts ...pulumi.ResourceOption) (*ThreatIntelligenceAlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.AlertRuleTemplateName == nil {
		return nil, errors.New("invalid value for required argument 'AlertRuleTemplateName'")
	}
	if args.Enabled == nil {
		return nil, errors.New("invalid value for required argument 'Enabled'")
	}
	if args.Kind == nil {
		return nil, errors.New("invalid value for required argument 'Kind'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("ThreatIntelligence")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:ThreatIntelligenceAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ThreatIntelligenceAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20210901preview:ThreatIntelligenceAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetThreatIntelligenceAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelligenceAlertRuleState, opts ...pulumi.ResourceOption) (*ThreatIntelligenceAlertRule, error) {
	var resource ThreatIntelligenceAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20210901preview:ThreatIntelligenceAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type threatIntelligenceAlertRuleState struct {
}

type ThreatIntelligenceAlertRuleState struct {
}

func (ThreatIntelligenceAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceAlertRuleState)(nil)).Elem()
}

type threatIntelligenceAlertRuleArgs struct {
	AlertRuleTemplateName string  `pulumi:"alertRuleTemplateName"`
	Enabled               bool    `pulumi:"enabled"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	RuleId                *string `pulumi:"ruleId"`
	WorkspaceName         string  `pulumi:"workspaceName"`
}


type ThreatIntelligenceAlertRuleArgs struct {
	AlertRuleTemplateName pulumi.StringInput
	Enabled               pulumi.BoolInput
	Kind                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	RuleId                pulumi.StringPtrInput
	WorkspaceName         pulumi.StringInput
}

func (ThreatIntelligenceAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*threatIntelligenceAlertRuleArgs)(nil)).Elem()
}

type ThreatIntelligenceAlertRuleInput interface {
	pulumi.Input

	ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput
	ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput
}

func (*ThreatIntelligenceAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceAlertRule)(nil)).Elem()
}

func (i *ThreatIntelligenceAlertRule) ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput {
	return i.ToThreatIntelligenceAlertRuleOutputWithContext(context.Background())
}

func (i *ThreatIntelligenceAlertRule) ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceAlertRuleOutput)
}

type ThreatIntelligenceAlertRuleOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ThreatIntelligenceAlertRule)(nil)).Elem()
}

func (o ThreatIntelligenceAlertRuleOutput) ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput {
	return o
}

func (o ThreatIntelligenceAlertRuleOutput) ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput {
	return o
}

func (o ThreatIntelligenceAlertRuleOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o ThreatIntelligenceAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *ThreatIntelligenceAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(ThreatIntelligenceAlertRuleOutput{})
}
