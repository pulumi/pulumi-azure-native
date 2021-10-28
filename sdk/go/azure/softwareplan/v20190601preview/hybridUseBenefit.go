


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
			Type: pulumi.String("azure-nextgen:softwareplan/v20190601preview:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azure-native:softwareplan:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azure-nextgen:softwareplan:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azure-native:softwareplan/v20191201:HybridUseBenefit"),
		},
		{
			Type: pulumi.String("azure-nextgen:softwareplan/v20191201:HybridUseBenefit"),
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
	return reflect.TypeOf((*HybridUseBenefit)(nil))
}

func (i *HybridUseBenefit) ToHybridUseBenefitOutput() HybridUseBenefitOutput {
	return i.ToHybridUseBenefitOutputWithContext(context.Background())
}

func (i *HybridUseBenefit) ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HybridUseBenefitOutput)
}

type HybridUseBenefitOutput struct{ *pulumi.OutputState }

func (HybridUseBenefitOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HybridUseBenefit)(nil))
}

func (o HybridUseBenefitOutput) ToHybridUseBenefitOutput() HybridUseBenefitOutput {
	return o
}

func (o HybridUseBenefitOutput) ToHybridUseBenefitOutputWithContext(ctx context.Context) HybridUseBenefitOutput {
	return o
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*HybridUseBenefitInput)(nil)).Elem(), &HybridUseBenefit{})
	pulumi.RegisterOutputType(HybridUseBenefitOutput{})
}
