


package v20190101preview

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
	Tactics               pulumi.StringArrayOutput `pulumi:"tactics"`
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
	if args.OperationalInsightsResourceProvider == nil {
		return nil, errors.New("invalid value for required argument 'OperationalInsightsResourceProvider'")
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
			Type: pulumi.String("azure-native:securityinsights/v20200101:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:MLBehaviorAnalyticsAlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210901preview:MLBehaviorAnalyticsAlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource MLBehaviorAnalyticsAlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:MLBehaviorAnalyticsAlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMLBehaviorAnalyticsAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MLBehaviorAnalyticsAlertRuleState, opts ...pulumi.ResourceOption) (*MLBehaviorAnalyticsAlertRule, error) {
	var resource MLBehaviorAnalyticsAlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:MLBehaviorAnalyticsAlertRule", name, id, state, &resource, opts...)
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
	AlertRuleTemplateName               string  `pulumi:"alertRuleTemplateName"`
	Enabled                             bool    `pulumi:"enabled"`
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	RuleId                              *string `pulumi:"ruleId"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type MLBehaviorAnalyticsAlertRuleArgs struct {
	AlertRuleTemplateName               pulumi.StringInput
	Enabled                             pulumi.BoolInput
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	RuleId                              pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
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

func init() {
	pulumi.RegisterOutputType(MLBehaviorAnalyticsAlertRuleOutput{})
}
