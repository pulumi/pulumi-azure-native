


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
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	RuleId                              *string `pulumi:"ruleId"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type FusionAlertRuleArgs struct {
	AlertRuleTemplateName               pulumi.StringInput
	Enabled                             pulumi.BoolInput
	Etag                                pulumi.StringPtrInput
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
	return reflect.TypeOf((*FusionAlertRule)(nil))
}

func (i *FusionAlertRule) ToFusionAlertRuleOutput() FusionAlertRuleOutput {
	return i.ToFusionAlertRuleOutputWithContext(context.Background())
}

func (i *FusionAlertRule) ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FusionAlertRuleOutput)
}

type FusionAlertRuleOutput struct{ *pulumi.OutputState }

func (FusionAlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FusionAlertRule)(nil))
}

func (o FusionAlertRuleOutput) ToFusionAlertRuleOutput() FusionAlertRuleOutput {
	return o
}

func (o FusionAlertRuleOutput) ToFusionAlertRuleOutputWithContext(ctx context.Context) FusionAlertRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(FusionAlertRuleOutput{})
}
