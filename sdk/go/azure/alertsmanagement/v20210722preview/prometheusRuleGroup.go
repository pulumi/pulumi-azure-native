


package v20210722preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrometheusRuleGroup struct {
	pulumi.CustomResourceState

	ClusterName pulumi.StringPtrOutput            `pulumi:"clusterName"`
	Description pulumi.StringPtrOutput            `pulumi:"description"`
	Enabled     pulumi.BoolPtrOutput              `pulumi:"enabled"`
	Interval    pulumi.StringPtrOutput            `pulumi:"interval"`
	Location    pulumi.StringOutput               `pulumi:"location"`
	Name        pulumi.StringOutput               `pulumi:"name"`
	Rules       PrometheusRuleResponseArrayOutput `pulumi:"rules"`
	Scopes      pulumi.StringArrayOutput          `pulumi:"scopes"`
	SystemData  SystemDataResponseOutput          `pulumi:"systemData"`
	Tags        pulumi.StringMapOutput            `pulumi:"tags"`
	Type        pulumi.StringOutput               `pulumi:"type"`
}


func NewPrometheusRuleGroup(ctx *pulumi.Context,
	name string, args *PrometheusRuleGroupArgs, opts ...pulumi.ResourceOption) (*PrometheusRuleGroup, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Rules == nil {
		return nil, errors.New("invalid value for required argument 'Rules'")
	}
	if args.Scopes == nil {
		return nil, errors.New("invalid value for required argument 'Scopes'")
	}
	var resource PrometheusRuleGroup
	err := ctx.RegisterResource("azure-native:alertsmanagement/v20210722preview:PrometheusRuleGroup", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPrometheusRuleGroup(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PrometheusRuleGroupState, opts ...pulumi.ResourceOption) (*PrometheusRuleGroup, error) {
	var resource PrometheusRuleGroup
	err := ctx.ReadResource("azure-native:alertsmanagement/v20210722preview:PrometheusRuleGroup", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type prometheusRuleGroupState struct {
}

type PrometheusRuleGroupState struct {
}

func (PrometheusRuleGroupState) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusRuleGroupState)(nil)).Elem()
}

type prometheusRuleGroupArgs struct {
	ClusterName       *string           `pulumi:"clusterName"`
	Description       *string           `pulumi:"description"`
	Enabled           *bool             `pulumi:"enabled"`
	Interval          *string           `pulumi:"interval"`
	Location          *string           `pulumi:"location"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RuleGroupName     *string           `pulumi:"ruleGroupName"`
	Rules             []PrometheusRule  `pulumi:"rules"`
	Scopes            []string          `pulumi:"scopes"`
	Tags              map[string]string `pulumi:"tags"`
}


type PrometheusRuleGroupArgs struct {
	ClusterName       pulumi.StringPtrInput
	Description       pulumi.StringPtrInput
	Enabled           pulumi.BoolPtrInput
	Interval          pulumi.StringPtrInput
	Location          pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RuleGroupName     pulumi.StringPtrInput
	Rules             PrometheusRuleArrayInput
	Scopes            pulumi.StringArrayInput
	Tags              pulumi.StringMapInput
}

func (PrometheusRuleGroupArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*prometheusRuleGroupArgs)(nil)).Elem()
}

type PrometheusRuleGroupInput interface {
	pulumi.Input

	ToPrometheusRuleGroupOutput() PrometheusRuleGroupOutput
	ToPrometheusRuleGroupOutputWithContext(ctx context.Context) PrometheusRuleGroupOutput
}

func (*PrometheusRuleGroup) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleGroup)(nil)).Elem()
}

func (i *PrometheusRuleGroup) ToPrometheusRuleGroupOutput() PrometheusRuleGroupOutput {
	return i.ToPrometheusRuleGroupOutputWithContext(context.Background())
}

func (i *PrometheusRuleGroup) ToPrometheusRuleGroupOutputWithContext(ctx context.Context) PrometheusRuleGroupOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleGroupOutput)
}

type PrometheusRuleGroupOutput struct{ *pulumi.OutputState }

func (PrometheusRuleGroupOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleGroup)(nil)).Elem()
}

func (o PrometheusRuleGroupOutput) ToPrometheusRuleGroupOutput() PrometheusRuleGroupOutput {
	return o
}

func (o PrometheusRuleGroupOutput) ToPrometheusRuleGroupOutputWithContext(ctx context.Context) PrometheusRuleGroupOutput {
	return o
}

func (o PrometheusRuleGroupOutput) ClusterName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringPtrOutput { return v.ClusterName }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleGroupOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleGroupOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o PrometheusRuleGroupOutput) Interval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringPtrOutput { return v.Interval }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleGroupOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o PrometheusRuleGroupOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o PrometheusRuleGroupOutput) Rules() PrometheusRuleResponseArrayOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) PrometheusRuleResponseArrayOutput { return v.Rules }).(PrometheusRuleResponseArrayOutput)
}

func (o PrometheusRuleGroupOutput) Scopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringArrayOutput { return v.Scopes }).(pulumi.StringArrayOutput)
}

func (o PrometheusRuleGroupOutput) SystemData() SystemDataResponseOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) SystemDataResponseOutput { return v.SystemData }).(SystemDataResponseOutput)
}

func (o PrometheusRuleGroupOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o PrometheusRuleGroupOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *PrometheusRuleGroup) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(PrometheusRuleGroupOutput{})
}
