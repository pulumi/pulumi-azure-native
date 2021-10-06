


package v20190101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type AlertRule struct {
	pulumi.CustomResourceState

	Etag pulumi.StringPtrOutput `pulumi:"etag"`
	Kind pulumi.StringOutput    `pulumi:"kind"`
	Name pulumi.StringOutput    `pulumi:"name"`
	Type pulumi.StringOutput    `pulumi:"type"`
}


func NewAlertRule(ctx *pulumi.Context,
	name string, args *AlertRuleArgs, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
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
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20190101preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights:AlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20200101:AlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20200101:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:securityinsights/v20210301preview:AlertRule"),
		},
		{
			Type: pulumi.String("azure-nextgen:securityinsights/v20210301preview:AlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AlertRule
	err := ctx.RegisterResource("azure-native:securityinsights/v20190101preview:AlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleState, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	var resource AlertRule
	err := ctx.ReadResource("azure-native:securityinsights/v20190101preview:AlertRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type alertRuleState struct {
}

type AlertRuleState struct {
}

func (AlertRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleState)(nil)).Elem()
}

type alertRuleArgs struct {
	Etag                                *string `pulumi:"etag"`
	Kind                                string  `pulumi:"kind"`
	OperationalInsightsResourceProvider string  `pulumi:"operationalInsightsResourceProvider"`
	ResourceGroupName                   string  `pulumi:"resourceGroupName"`
	RuleId                              *string `pulumi:"ruleId"`
	WorkspaceName                       string  `pulumi:"workspaceName"`
}


type AlertRuleArgs struct {
	Etag                                pulumi.StringPtrInput
	Kind                                pulumi.StringInput
	OperationalInsightsResourceProvider pulumi.StringInput
	ResourceGroupName                   pulumi.StringInput
	RuleId                              pulumi.StringPtrInput
	WorkspaceName                       pulumi.StringInput
}

func (AlertRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*alertRuleArgs)(nil)).Elem()
}

type AlertRuleInput interface {
	pulumi.Input

	ToAlertRuleOutput() AlertRuleOutput
	ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput
}

func (*AlertRule) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRule)(nil))
}

func (i *AlertRule) ToAlertRuleOutput() AlertRuleOutput {
	return i.ToAlertRuleOutputWithContext(context.Background())
}

func (i *AlertRule) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleOutput)
}

type AlertRuleOutput struct{ *pulumi.OutputState }

func (AlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlertRule)(nil))
}

func (o AlertRuleOutput) ToAlertRuleOutput() AlertRuleOutput {
	return o
}

func (o AlertRuleOutput) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(AlertRuleOutput{})
}
