


package v20221005preview

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type MarkupRule struct {
	pulumi.CustomResourceState

	CustomerDetails CustomerMetadataResponseOutput `pulumi:"customerDetails"`
	Description     pulumi.StringPtrOutput         `pulumi:"description"`
	ETag            pulumi.StringPtrOutput         `pulumi:"eTag"`
	EndDate         pulumi.StringPtrOutput         `pulumi:"endDate"`
	Name            pulumi.StringOutput            `pulumi:"name"`
	Percentage      pulumi.Float64Output           `pulumi:"percentage"`
	StartDate       pulumi.StringOutput            `pulumi:"startDate"`
	Type            pulumi.StringOutput            `pulumi:"type"`
}


func NewMarkupRule(ctx *pulumi.Context,
	name string, args *MarkupRuleArgs, opts ...pulumi.ResourceOption) (*MarkupRule, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.BillingAccountId == nil {
		return nil, errors.New("invalid value for required argument 'BillingAccountId'")
	}
	if args.BillingProfileId == nil {
		return nil, errors.New("invalid value for required argument 'BillingProfileId'")
	}
	if args.CustomerDetails == nil {
		return nil, errors.New("invalid value for required argument 'CustomerDetails'")
	}
	if args.Percentage == nil {
		return nil, errors.New("invalid value for required argument 'Percentage'")
	}
	if args.StartDate == nil {
		return nil, errors.New("invalid value for required argument 'StartDate'")
	}
	var resource MarkupRule
	err := ctx.RegisterResource("azure-native:costmanagement/v20221005preview:MarkupRule", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetMarkupRule(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *MarkupRuleState, opts ...pulumi.ResourceOption) (*MarkupRule, error) {
	var resource MarkupRule
	err := ctx.ReadResource("azure-native:costmanagement/v20221005preview:MarkupRule", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type markupRuleState struct {
}

type MarkupRuleState struct {
}

func (MarkupRuleState) ElementType() reflect.Type {
	return reflect.TypeOf((*markupRuleState)(nil)).Elem()
}

type markupRuleArgs struct {
	BillingAccountId string           `pulumi:"billingAccountId"`
	BillingProfileId string           `pulumi:"billingProfileId"`
	CustomerDetails  CustomerMetadata `pulumi:"customerDetails"`
	Description      *string          `pulumi:"description"`
	ETag             *string          `pulumi:"eTag"`
	EndDate          *string          `pulumi:"endDate"`
	Name             *string          `pulumi:"name"`
	Percentage       float64          `pulumi:"percentage"`
	StartDate        string           `pulumi:"startDate"`
}


type MarkupRuleArgs struct {
	BillingAccountId pulumi.StringInput
	BillingProfileId pulumi.StringInput
	CustomerDetails  CustomerMetadataInput
	Description      pulumi.StringPtrInput
	ETag             pulumi.StringPtrInput
	EndDate          pulumi.StringPtrInput
	Name             pulumi.StringPtrInput
	Percentage       pulumi.Float64Input
	StartDate        pulumi.StringInput
}

func (MarkupRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*markupRuleArgs)(nil)).Elem()
}

type MarkupRuleInput interface {
	pulumi.Input

	ToMarkupRuleOutput() MarkupRuleOutput
	ToMarkupRuleOutputWithContext(ctx context.Context) MarkupRuleOutput
}

func (*MarkupRule) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkupRule)(nil)).Elem()
}

func (i *MarkupRule) ToMarkupRuleOutput() MarkupRuleOutput {
	return i.ToMarkupRuleOutputWithContext(context.Background())
}

func (i *MarkupRule) ToMarkupRuleOutputWithContext(ctx context.Context) MarkupRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MarkupRuleOutput)
}

type MarkupRuleOutput struct{ *pulumi.OutputState }

func (MarkupRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MarkupRule)(nil)).Elem()
}

func (o MarkupRuleOutput) ToMarkupRuleOutput() MarkupRuleOutput {
	return o
}

func (o MarkupRuleOutput) ToMarkupRuleOutputWithContext(ctx context.Context) MarkupRuleOutput {
	return o
}

func (o MarkupRuleOutput) CustomerDetails() CustomerMetadataResponseOutput {
	return o.ApplyT(func(v *MarkupRule) CustomerMetadataResponseOutput { return v.CustomerDetails }).(CustomerMetadataResponseOutput)
}

func (o MarkupRuleOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkupRule) pulumi.StringPtrOutput { return v.Description }).(pulumi.StringPtrOutput)
}

func (o MarkupRuleOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkupRule) pulumi.StringPtrOutput { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o MarkupRuleOutput) EndDate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MarkupRule) pulumi.StringPtrOutput { return v.EndDate }).(pulumi.StringPtrOutput)
}

func (o MarkupRuleOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *MarkupRule) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o MarkupRuleOutput) Percentage() pulumi.Float64Output {
	return o.ApplyT(func(v *MarkupRule) pulumi.Float64Output { return v.Percentage }).(pulumi.Float64Output)
}

func (o MarkupRuleOutput) StartDate() pulumi.StringOutput {
	return o.ApplyT(func(v *MarkupRule) pulumi.StringOutput { return v.StartDate }).(pulumi.StringOutput)
}

func (o MarkupRuleOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *MarkupRule) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(MarkupRuleOutput{})
}
