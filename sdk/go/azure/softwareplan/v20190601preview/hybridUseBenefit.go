


package v20190601preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type HybridUseBenefit struct {
	pulumi.CustomResourceState

	CreatedDate       pulumi.StringOutput `pulumi:"createdDate"`
	Etag              pulumi.IntOutput    `pulumi:"etag"`
	LastUpdatedDate   pulumi.StringOutput `pulumi:"lastUpdatedDate"`
	Name              pulumi.StringOutput `pulumi:"name"`
	ProvisioningState pulumi.StringOutput `pulumi:"provisioningState"`
	Sku               SkuResponseOutput   `pulumi:"sku"`
	Type              pulumi.StringOutput `pulumi:"type"`
}


func NewHybridUseBenefit(ctx *pulumi.Context,
	name string, args *HybridUseBenefitArgs, opts ...pulumi.ResourceOption) (*HybridUseBenefit, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	if args.Sku == nil {
		return nil, errors.New("invalid value for required argument 'Sku'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:softwareplan:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azure-native:softwareplan/v20191201:HybridUseBenefit"),
		},
	})
	opts = append(opts, aliases)
	var resource HybridUseBenefit
	err := ctx.RegisterResource("azure-native:softwareplan/v20190601preview:HybridUseBenefit", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetHybridUseBenefit(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *HybridUseBenefitState, opts ...pulumi.ResourceOption) (*HybridUseBenefit, error) {
	var resource HybridUseBenefit
	err := ctx.ReadResource("azure-native:softwareplan/v20190601preview:HybridUseBenefit", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type hybridUseBenefitState struct {
}

type HybridUseBenefitState struct {
}

func (HybridUseBenefitState) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridUseBenefitState)(nil)).Elem()
}

type hybridUseBenefitArgs struct {
	PlanId *string `pulumi:"planId"`
	Scope  string  `pulumi:"scope"`
	Sku    Sku     `pulumi:"sku"`
}


type HybridUseBenefitArgs struct {
	PlanId pulumi.StringPtrInput
	Scope  pulumi.StringInput
	Sku    SkuInput
}

func (HybridUseBenefitArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*hybridUseBenefitArgs)(nil)).Elem()
}

type HybridUseBenefitInput interface {
	pulumi.Input

	ToHybridUseBenefitOutput() HybridUseBenefitOutput
	ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput
}

func (*HybridUseBenefit) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridUseBenefit)(nil)).Elem()
}

func (i *HybridUseBenefit) ToHybridUseBenefitOutput() HybridUseBenefitOutput {
	return i.ToHybridUseBenefitOutputWithContext(context.Background())
}

func (i *HybridUseBenefit) ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridUseBenefitOutput)
}

type HybridUseBenefitOutput struct{ *pulumi.OutputState }

func (HybridUseBenefitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HybridUseBenefit)(nil)).Elem()
}

func (o HybridUseBenefitOutput) ToHybridUseBenefitOutput() HybridUseBenefitOutput {
	return o
}

func (o HybridUseBenefitOutput) ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput {
	return o
}

func (o HybridUseBenefitOutput) CreatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.CreatedDate }).(pulumi.StringOutput)
}

func (o HybridUseBenefitOutput) Etag() pulumi.IntOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.IntOutput { return v.Etag }).(pulumi.IntOutput)
}

func (o HybridUseBenefitOutput) LastUpdatedDate() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.LastUpdatedDate }).(pulumi.StringOutput)
}

func (o HybridUseBenefitOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o HybridUseBenefitOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o HybridUseBenefitOutput) Sku() SkuResponseOutput {
	return o.ApplyT(func(v *HybridUseBenefit) SkuResponseOutput { return v.Sku }).(SkuResponseOutput)
}

func (o HybridUseBenefitOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *HybridUseBenefit) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(HybridUseBenefitOutput{})
}
