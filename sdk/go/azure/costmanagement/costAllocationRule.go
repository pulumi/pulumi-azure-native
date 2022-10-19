


package costmanagement

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type CostAllocationRule struct {
	pulumi.CustomResourceState

	Name       pulumi.StringOutput                        `pulumi:"name"`
	Properties CostAllocationRulePropertiesResponseOutput `pulumi:"properties"`
	Type       pulumi.StringOutput                        `pulumi:"type"`
}


func NewCostAllocationRule(ctx *pulumi.Context,
	name string, args *CostAllocationRuleArgs, opts ...pulumi.ResourceOption) (*CostAllocationRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:costmanagement/v20200301preview:CostAllocationRule"),
		},
	})
	opts = append(opts, aliases)
	var resource CostAllocationRule
	err := ctx.RegisterResource("azure-native:costmanagement:CostAllocationRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetCostAllocationRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *CostAllocationRuleState, opts ...pulumi.ResourceOption) (*CostAllocationRule, error) {
	var resource CostAllocationRule
	err := ctx.ReadResource("azure-native:costmanagement:CostAllocationRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type costAllocationRuleState struct {
}

type CostAllocationRuleState struct {
}

func (CostAllocationRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*costAllocationRuleState)(nil)).Elem()
}

type costAllocationRuleArgs struct {
	BillingAccountId string                        `pulumi:"billingAccountId"`
	Properties       *CostAllocationRuleProperties `pulumi:"properties"`
	RuleName         *string                       `pulumi:"ruleName"`
}


type CostAllocationRuleArgs struct {
	BillingAccountId pulumi.StringInput
	Properties       CostAllocationRulePropertiesPtrInput
	RuleName         pulumi.StringPtrInput
}

func (CostAllocationRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*costAllocationRuleArgs)(nil)).Elem()
}

type CostAllocationRuleInput interface {
	pulumi.Input

	ToCostAllocationRuleOutput() CostAllocationRuleOutput
	ToCostAllocationRuleOutputWithContext(ctx context.Context) CostAllocationRuleOutput
}

func (*CostAllocationRule) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRule)(nil)).Elem()
}

func (i *CostAllocationRule) ToCostAllocationRuleOutput() CostAllocationRuleOutput {
	return i.ToCostAllocationRuleOutputWithContext(context.Background())
}

func (i *CostAllocationRule) ToCostAllocationRuleOutputWithContext(ctx context.Context) CostAllocationRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CostAllocationRuleOutput)
}

type CostAllocationRuleOutput struct{ *pulumi.OutputState }

func (CostAllocationRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CostAllocationRule)(nil)).Elem()
}

func (o CostAllocationRuleOutput) ToCostAllocationRuleOutput() CostAllocationRuleOutput {
	return o
}

func (o CostAllocationRuleOutput) ToCostAllocationRuleOutputWithContext(ctx context.Context) CostAllocationRuleOutput {
	return o
}

func (o CostAllocationRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *CostAllocationRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o CostAllocationRuleOutput) Properties() CostAllocationRulePropertiesResponseOutput {
	return o.ApplyT(func(v *CostAllocationRule) CostAllocationRulePropertiesResponseOutput { return v.Properties }).(CostAllocationRulePropertiesResponseOutput)
}

func (o CostAllocationRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *CostAllocationRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(CostAllocationRuleOutput{})
}
