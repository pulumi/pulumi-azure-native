


package insights

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlertRule struct {
	pulumi.CustomResourceState

	Action            pulumi.AnyOutput       `pulumi:"action"`
	Actions           pulumi.ArrayOutput     `pulumi:"actions"`
	Condition         pulumi.AnyOutput       `pulumi:"condition"`
	Description       pulumi.StringPtrOutput `pulumi:"description"`
	IsEnabled         pulumi.BoolOutput      `pulumi:"isEnabled"`
	LastUpdatedTime   pulumi.StringOutput    `pulumi:"lastUpdatedTime"`
	Location          pulumi.StringOutput    `pulumi:"location"`
	Name              pulumi.StringOutput    `pulumi:"name"`
	ProvisioningState pulumi.StringPtrOutput `pulumi:"provisioningState"`
	Tags              pulumi.StringMapOutput `pulumi:"tags"`
	Type              pulumi.StringOutput    `pulumi:"type"`
}


func NewAlertRule(ctx *pulumi.Context,
	name string, args *AlertRuleArgs, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Condition == nil {
		return nil, errors.New("invalid value for required argument 'Condition'")
	}
	if args.IsEnabled == nil {
		return nil, errors.New("invalid value for required argument 'IsEnabled'")
	}
	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:insights/v20140401:AlertRule"),
		},
		{
			Type: pulumi.String("azure-native:insights/v20160301:AlertRule"),
		},
	})
	opts = append(opts, aliases)
	var resource AlertRule
	err := ctx.RegisterResource("azure-native:insights:AlertRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetAlertRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *AlertRuleState, opts ...pulumi.ResourceOption) (*AlertRule, error) {
	var resource AlertRule
	err := ctx.ReadResource("azure-native:insights:AlertRule", name, id, state, &resource, opts...)
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
	Action            interface{}       `pulumi:"action"`
	Actions           []interface{}     `pulumi:"actions"`
	Condition         interface{}       `pulumi:"condition"`
	Description       *string           `pulumi:"description"`
	IsEnabled         bool              `pulumi:"isEnabled"`
	Location          *string           `pulumi:"location"`
	Name              string            `pulumi:"name"`
	ProvisioningState *string           `pulumi:"provisioningState"`
	ResourceGroupName string            `pulumi:"resourceGroupName"`
	RuleName          *string           `pulumi:"ruleName"`
	Tags              map[string]string `pulumi:"tags"`
}


type AlertRuleArgs struct {
	Action            pulumi.Input
	Actions           pulumi.ArrayInput
	Condition         pulumi.Input
	Description       pulumi.StringPtrInput
	IsEnabled         pulumi.BoolInput
	Location          pulumi.StringPtrInput
	Name              pulumi.StringInput
	ProvisioningState pulumi.StringPtrInput
	ResourceGroupName pulumi.StringInput
	RuleName          pulumi.StringPtrInput
	Tags              pulumi.StringMapInput
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
	return reflect.TypeOf((**AlertRule)(nil)).Elem()
}

func (i *AlertRule) ToAlertRuleOutput() AlertRuleOutput {
	return i.ToAlertRuleOutputWithContext(context.Background())
}

func (i *AlertRule) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlertRuleOutput)
}

type AlertRuleOutput struct{ *pulumi.OutputState }

func (AlertRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlertRule)(nil)).Elem()
}

func (o AlertRuleOutput) ToAlertRuleOutput() AlertRuleOutput {
	return o
}

func (o AlertRuleOutput) ToAlertRuleOutputWithContext(ctx context.Context) AlertRuleOutput {
	return o
}

func (o AlertRuleOutput) Action() pulumi.AnyOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.AnyOutput { return v.Action }).(pulumi.AnyOutput)
}

func (o AlertRuleOutput) Actions() pulumi.ArrayOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.ArrayOutput { return v.Actions }).(pulumi.ArrayOutput)
}

func (o AlertRuleOutput) Condition() pulumi.AnyOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.AnyOutput { return v.Condition }).(pulumi.AnyOutput)
}

func (o AlertRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AlertRuleOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.BoolOutput { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AlertRuleOutput) LastUpdatedTime() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.LastUpdatedTime }).(pulumi.StringOutput)
}

func (o AlertRuleOutput) Location() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Location }).(pulumi.StringOutput)
}

func (o AlertRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o AlertRuleOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringPtrOutput { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o AlertRuleOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringMapOutput { return v.Tags }).(pulumi.StringMapOutput)
}

func (o AlertRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *AlertRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(AlertRuleOutput{})
}
