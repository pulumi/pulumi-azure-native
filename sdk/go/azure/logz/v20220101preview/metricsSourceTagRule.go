


package v20220101preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MetricsSourceTagRule struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                     `pulumi:"name"`
	Properties MetricsTagRulesPropertiesResponseOutput `pulumi:"properties"`
	SystemData SystemDataResponseOutput                `pulumi:"systemData"`
	Type       pulumi.StringOutput                     `pulumi:"type"`
}


func NewMetricsSourceTagRule(ctx *pulumi.Context,
	name string, args *MetricsSourceTagRuleArgs, opts ...pulumi.ResourceOption) (*MetricsSourceTagRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.MetricsSourceName == nil {
		return nil, errors.New("invalid value for required argument 'MetricsSourceName'")
	}
	if args.MonitorName == nil {
		return nil, errors.New("invalid value for required argument 'MonitorName'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:logz:MetricsSourceTagRule"),
		},
	})
	opts = append(opts, aliases)
	var resource MetricsSourceTagRule
	err := ctx.RegisterResource("azure-native:logz/v20220101preview:MetricsSourceTagRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMetricsSourceTagRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MetricsSourceTagRuleState, opts ...pulumi.ResourceOption) (*MetricsSourceTagRule, error) {
	var resource MetricsSourceTagRule
	err := ctx.ReadResource("azure-native:logz/v20220101preview:MetricsSourceTagRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type metricsSourceTagRuleState struct {
}

type MetricsSourceTagRuleState struct {
}

func (MetricsSourceTagRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsSourceTagRuleState)(nil)).Elem()
}

type metricsSourceTagRuleArgs struct {
	MetricsSourceName string                     `pulumi:"metricsSourceName"`
	MonitorName       string                     `pulumi:"monitorName"`
	Properties        *MetricsTagRulesProperties `pulumi:"properties"`
	ResourceGroupName string                     `pulumi:"resourceGroupName"`
	RuleSetName       *string                    `pulumi:"ruleSetName"`
}


type MetricsSourceTagRuleArgs struct {
	MetricsSourceName pulumi.StringInput
	MonitorName       pulumi.StringInput
	Properties        MetricsTagRulesPropertiesPtrInput
	ResourceGroupName pulumi.StringInput
	RuleSetName       pulumi.StringPtrInput
}

func (MetricsSourceTagRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*metricsSourceTagRuleArgs)(nil)).Elem()
}

type MetricsSourceTagRuleInput interface {
	pulumi.Input

	ToMetricsSourceTagRuleOutput() MetricsSourceTagRuleOutput
	ToMetricsSourceTagRuleOutputWithContext(ctx context.Context) MetricsSourceTagRuleOutput
}

func (*MetricsSourceTagRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsSourceTagRule)(nil)).Elem()
}

func (i *MetricsSourceTagRule) ToMetricsSourceTagRuleOutput() MetricsSourceTagRuleOutput {
	return i.ToMetricsSourceTagRuleOutputWithContext(context.Background())
}

func (i *MetricsSourceTagRule) ToMetricsSourceTagRuleOutputWithContext(ctx context.Context) MetricsSourceTagRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricsSourceTagRuleOutput)
}

type MetricsSourceTagRuleOutput struct{ *pulumi.OutputState }

func (MetricsSourceTagRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MetricsSourceTagRule)(nil)).Elem()
}

func (o MetricsSourceTagRuleOutput) ToMetricsSourceTagRuleOutput() MetricsSourceTagRuleOutput {
	return o
}

func (o MetricsSourceTagRuleOutput) ToMetricsSourceTagRuleOutputWithContext(ctx context.Context) MetricsSourceTagRuleOutput {
	return o
}

func (o MetricsSourceTagRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MetricsSourceTagRuleOutput) Properties() MetricsTagRulesPropertiesResponseOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) MetricsTagRulesPropertiesResponseOutput { return v.Properties }).(MetricsTagRulesPropertiesResponseOutput)
}

func (o MetricsSourceTagRuleOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o MetricsSourceTagRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MetricsSourceTagRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MetricsSourceTagRuleOutput{})
}
