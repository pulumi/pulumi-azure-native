


package authorization

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PolicyPricing struct {
	pulumi.CustomResourceState

	Name              pulumi.StringOutput      `pulumi:"name"`
	PricingTier       pulumi.StringOutput      `pulumi:"pricingTier"`
	ProvisioningState pulumi.StringOutput      `pulumi:"provisioningState"`
	SystemData        SystemDataResponseOutput `pulumi:"systemData"`
	Type              pulumi.StringOutput      `pulumi:"type"`
}


func NewPolicyPricing(ctx *pulumi.Context,
	name string, args *PolicyPricingArgs, opts ...pulumi.ResourceOption) (*PolicyPricing, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.PricingTier == nil {
		return nil, errors.New("invalid value for required argument 'PricingTier'")
	}
	if args.Scope == nil {
		return nil, errors.New("invalid value for required argument 'Scope'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-nextgen:authorization:PolicyPricing"),
		},
		{
			Type: pulumi.String("azure-native:authorization/v20210701:PolicyPricing"),
		},
		{
			Type: pulumi.String("azure-nextgen:authorization/v20210701:PolicyPricing"),
		},
	})
	opts = append(opts, aliases)
	var resource PolicyPricing
	err := ctx.RegisterResource("azure-native:authorization:PolicyPricing", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetPolicyPricing(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *PolicyPricingState, opts ...pulumi.ResourceOption) (*PolicyPricing, error) {
	var resource PolicyPricing
	err := ctx.ReadResource("azure-native:authorization:PolicyPricing", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type policyPricingState struct {
}

type PolicyPricingState struct {
}

func (PolicyPricingState) ElementType() reflect.Type {
	return reflect.TypeOf((*policyPricingState)(nil)).Elem()
}

type policyPricingArgs struct {
	PolicyPricingName *string `pulumi:"policyPricingName"`
	PricingTier       string  `pulumi:"pricingTier"`
	Scope             string  `pulumi:"scope"`
}


type PolicyPricingArgs struct {
	PolicyPricingName pulumi.StringPtrInput
	PricingTier       pulumi.StringInput
	Scope             pulumi.StringInput
}

func (PolicyPricingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*policyPricingArgs)(nil)).Elem()
}

type PolicyPricingInput interface {
	pulumi.Input

	ToPolicyPricingOutput() PolicyPricingOutput
	ToPolicyPricingOutputWithContext(ctx context.Context) PolicyPricingOutput
}

func (*PolicyPricing) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPricing)(nil))
}

func (i *PolicyPricing) ToPolicyPricingOutput() PolicyPricingOutput {
	return i.ToPolicyPricingOutputWithContext(context.Background())
}

func (i *PolicyPricing) ToPolicyPricingOutputWithContext(ctx context.Context) PolicyPricingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PolicyPricingOutput)
}

type PolicyPricingOutput struct{ *pulumi.OutputState }

func (PolicyPricingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyPricing)(nil))
}

func (o PolicyPricingOutput) ToPolicyPricingOutput() PolicyPricingOutput {
	return o
}

func (o PolicyPricingOutput) ToPolicyPricingOutputWithContext(ctx context.Context) PolicyPricingOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(PolicyPricingOutput{})
}
