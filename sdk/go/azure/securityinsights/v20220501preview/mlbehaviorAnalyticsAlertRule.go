


package v20220501preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MLBehaviorAnalyticsAlertRule struct {
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
	Techniques            pulumi.StringArrayOutput `pulumi:"techniques"`
	Type                  pulumi.StringOutput      `pulumi:"type"`
}


func NewMLBehaviorAnalyticsAlertRule(ctx *pulumi.Context,
	name string, args *MLBehaviorAnalyticsAlertRuleArgs, opts ...pulumi.ResourceOption) (*MLBehaviorAnalyticsAlertRule, error) {
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
	args.Kind = pulumi.String("MLBehaviorAnalytics")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220901preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221001preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20221101preview:MLBehaviorAnalyticsAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource MLBehaviorAnalyticsAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20220501preview:MLBehaviorAnalyticsAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMLBehaviorAnalyticsAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLBehaviorAnalyticsAlertRuleState, opts ...pulumi.ResourceOption) (*MLBehaviorAnalyticsAlertRule, error) {
	var resource MLBehaviorAnalyticsAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20220501preview:MLBehaviorAnalyticsAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type mlbehaviorAnalyticsAlertRuleState struct {
}

type MLBehaviorAnalyticsAlertRuleState struct {
}

func (MLBehaviorAnalyticsAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*mlbehaviorAnalyticsAlertRuleState)(nil)).Elem()
}

type mlbehaviorAnalyticsAlertRuleArgs struct {
	AlertRuleTemplateName string  `pulumi:"alertRuleTemplateName"`
	Enabled               bool    `pulumi:"enabled"`
	Kind                  string  `pulumi:"kind"`
	ResourceGroupName     string  `pulumi:"resourceGroupName"`
	RuleId                *string `pulumi:"ruleId"`
	WorkspaceName         string  `pulumi:"workspaceName"`
}


type MLBehaviorAnalyticsAlertRuleArgs struct {
	AlertRuleTemplateName pulumi.StringInput
	Enabled               pulumi.BoolInput
	Kind                  pulumi.StringInput
	ResourceGroupName     pulumi.StringInput
	RuleId                pulumi.StringPtrInput
	WorkspaceName         pulumi.StringInput
}

func (MLBehaviorAnalyticsAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*mlbehaviorAnalyticsAlertRuleArgs)(nil)).Elem()
}

type MLBehaviorAnalyticsAlertRuleInput interface {
	pulumi.Input

	ToMLBehaviorAnalyticsAlertRuleOutput() MLBehaviorAnalyticsAlertRuleOutput
	ToMLBehaviorAnalyticsAlertRuleOutputWithContext(ctx context.Context) MLBehaviorAnalyticsAlertRuleOutput
}

func (*MLBehaviorAnalyticsAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MLBehaviorAnalyticsAlertRule)(nil)).Elem()
}

func (i *MLBehaviorAnalyticsAlertRule) ToMLBehaviorAnalyticsAlertRuleOutput() MLBehaviorAnalyticsAlertRuleOutput {
	return i.ToMLBehaviorAnalyticsAlertRuleOutputWithContext(context.Background())
}

func (i *MLBehaviorAnalyticsAlertRule) ToMLBehaviorAnalyticsAlertRuleOutputWithContext(ctx context.Context) MLBehaviorAnalyticsAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MLBehaviorAnalyticsAlertRuleOutput)
}

type MLBehaviorAnalyticsAlertRuleOutput struct{ *pulumi.OutputState }

func (MLBehaviorAnalyticsAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MLBehaviorAnalyticsAlertRule)(nil)).Elem()
}

func (o MLBehaviorAnalyticsAlertRuleOutput) ToMLBehaviorAnalyticsAlertRuleOutput() MLBehaviorAnalyticsAlertRuleOutput {
	return o
}

func (o MLBehaviorAnalyticsAlertRuleOutput) ToMLBehaviorAnalyticsAlertRuleOutputWithContext(ctx context.Context) MLBehaviorAnalyticsAlertRuleOutput {
	return o
}

func (o MLBehaviorAnalyticsAlertRuleOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Techniques() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringArrayOutput { return v.Techniques }).(pulumi.StringArrayOutput)
}

func (o MLBehaviorAnalyticsAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MLBehaviorAnalyticsAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MLBehaviorAnalyticsAlertRuleOutput{})
}
