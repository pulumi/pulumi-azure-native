


package v20210301preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type FusionAlertRule struct {
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


func NewFusionAlertRule(ctx *pulumi.Context,
	name string, args *FusionAlertRuleArgs, opts ...pulumi.ResourceOption) (*FusionAlertRule, error) {
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
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.WorkspaceName == nil {
		return nil, errors.New("invalid value for required argument 'WorkspaceName'")
	}
	args.Kind = pulumi.String("Fusion")
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:securityinsights:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20211001preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220101preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220401preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220501preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220601preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220701preview:FusionAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20220801preview:FusionAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource FusionAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:FusionAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetFusionAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *FusionAlertRuleState, opts ...pulumi.ResourceOption) (*FusionAlertRule, error) {
	var resource FusionAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:FusionAlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type fusionAlertRuleState struct {
}

type FusionAlertRuleState struct {
}

func (FusionAlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAlertRuleState)(nil)).Elem()
}

type fusionAlertRuleArgs struct {
	AlertRuleTemplateName               string  `pulumi:"alertRuleTemplateName"`
	Enabled                             bool    `pulumi:"enabled"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	RuleId                              *string `pulumi:"ruleId"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type FusionAlertRuleArgs struct {
	AlertRuleTemplateName               pulumi.StringInput
	Enabled                             pulumi.BoolInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	RuleId                              pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (FusionAlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*fusionAlertRuleArgs)(nil)).Elem()
}

type FusionAlertRuleInput interface {
	pulumi.Input

	ToFusionAlertRuleOutput() FusionAlertRuleOutput
	ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput
}

func (*FusionAlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAlertRule)(nil)).Elem()
}

func (i *FusionAlertRule) ToFusionAlertRuleOutput() FusionAlertRuleOutput {
	return i.ToFusionAlertRuleOutputWithContext(context.Background())
}

func (i *FusionAlertRule) ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAlertRuleOutput)
}

type FusionAlertRuleOutput struct{ *pulumi.OutputState }

func (FusionAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FusionAlertRule)(nil)).Elem()
}

func (o FusionAlertRuleOutput) ToFusionAlertRuleOutput() FusionAlertRuleOutput {
	return o
}

func (o FusionAlertRuleOutput) ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput {
	return o
}

func (o FusionAlertRuleOutput) AlertRuleTemplateName() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.AlertRuleTemplateName }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Description }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.DisplayName }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.BoolOutput { return v.Enabled }).(pulumi.BoolOutput)
}

func (o FusionAlertRuleOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringPtrOutput { return v.Etag }).(pulumi.StringPtrOutput)
}

func (o FusionAlertRuleOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Kind }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) LastModifiedUtc() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.LastModifiedUtc }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) Severity() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Severity }).(pulumi.StringOutput)
}

func (o FusionAlertRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *FusionAlertRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o FusionAlertRuleOutput) Tactics() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringArrayOutput { return v.Tactics }).(pulumi.StringArrayOutput)
}

func (o FusionAlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *FusionAlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(FusionAlertRuleOutput{})
}
