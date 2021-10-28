


package v20210301preview

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
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
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
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20190101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:ThreatIntelligenceAlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:ThreatIntelligenceAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource ThreatIntelligenceAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20210301preview:ThreatIntelligenceAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetThreatIntelligenceAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *ThreatIntelligenceAlertRuleState, opts ...pulumi.ResourceOption) (*ThreatIntelligenceAlertRule, error) {
	var resource ThreatIntelligenceAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20210301preview:ThreatIntelligenceAlertRule", name, id, state, &resource, opts...)
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
	AlertRuleTemplateName               string  `pulumi:"alertRuleTemplateName"`
	Enabled                             bool    `pulumi:"enabled"`
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	RuleId                              *string `pulumi:"ruleId"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type ThreatIntelligenceAlertRuleArgs struct {
	AlertRuleTemplateName               pulumi.StringInput
	Enabled                             pulumi.BoolInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	RuleId                              pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
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
	return reflect.TypeOf((*ThreatIntelligenceAlertRule)(nil))
}

func (i *ThreatIntelligenceAlertRule) ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput {
	return i.ToThreatIntelligenceAlertRuleOutputWithContext(context.Background())
}

func (i *ThreatIntelligenceAlertRule) ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ThreatIntelligenceAlertRuleOutput)
}

type ThreatIntelligenceAlertRuleOutput struct{ *pulumi.OutputState }

func (ThreatIntelligenceAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ThreatIntelligenceAlertRule)(nil))
}

func (o ThreatIntelligenceAlertRuleOutput) ToThreatIntelligenceAlertRuleOutput() ThreatIntelligenceAlertRuleOutput {
	return o
}

func (o ThreatIntelligenceAlertRuleOutput) ToThreatIntelligenceAlertRuleOutputWithContext(ctx context.Context) ThreatIntelligenceAlertRuleOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ThreatIntelligenceAlertRuleInput)(nil)).Elem(), &ThreatIntelligenceAlertRule{})
	pulumi.RegisterOutputType(ThreatIntelligenceAlertRuleOutput{})
}
