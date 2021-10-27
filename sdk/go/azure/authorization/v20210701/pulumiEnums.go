


package v20210701

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ExemptionCategory string

const (
	// This category of exemptions usually means the scope is not applicable for the policy.
	ExemptionCategoryWaiver = ExemptionCategory("Waiver")
	// This category of exemptions usually means the mitigation actions have been applied to the scope.
	ExemptionCategoryMitigated = ExemptionCategory("Mitigated")
)

func (ExemptionCategory) ElementType() reflect.Type {
	return reflect.TypeOf((*ExemptionCategory)(nil)).Elem()
}

func (e ExemptionCategory) ToExemptionCategoryOutput() ExemptionCategoryOutput {
	return pulumi.ToOutput(e).(ExemptionCategoryOutput)
}

func (e ExemptionCategory) ToExemptionCategoryOutputWithContext(ctx context.Context) ExemptionCategoryOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ExemptionCategoryOutput)
}

func (e ExemptionCategory) ToExemptionCategoryPtrOutput() ExemptionCategoryPtrOutput {
	return e.ToExemptionCategoryPtrOutputWithContext(context.Background())
}

func (e ExemptionCategory) ToExemptionCategoryPtrOutputWithContext(ctx context.Context) ExemptionCategoryPtrOutput {
	return ExemptionCategory(e).ToExemptionCategoryOutputWithContext(ctx).ToExemptionCategoryPtrOutputWithContext(ctx)
}

func (e ExemptionCategory) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExemptionCategory) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ExemptionCategory) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ExemptionCategory) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ExemptionCategoryOutput struct{ *pulumi.OutputState }

func (ExemptionCategoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExemptionCategory)(nil)).Elem()
}

func (o ExemptionCategoryOutput) ToExemptionCategoryOutput() ExemptionCategoryOutput {
	return o
}

func (o ExemptionCategoryOutput) ToExemptionCategoryOutputWithContext(ctx context.Context) ExemptionCategoryOutput {
	return o
}

func (o ExemptionCategoryOutput) ToExemptionCategoryPtrOutput() ExemptionCategoryPtrOutput {
	return o.ToExemptionCategoryPtrOutputWithContext(context.Background())
}

func (o ExemptionCategoryOutput) ToExemptionCategoryPtrOutputWithContext(ctx context.Context) ExemptionCategoryPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExemptionCategory) *ExemptionCategory {
		return &v
	}).(ExemptionCategoryPtrOutput)
}

func (o ExemptionCategoryOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ExemptionCategoryOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExemptionCategory) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ExemptionCategoryOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExemptionCategoryOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ExemptionCategory) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ExemptionCategoryPtrOutput struct{ *pulumi.OutputState }

func (ExemptionCategoryPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExemptionCategory)(nil)).Elem()
}

func (o ExemptionCategoryPtrOutput) ToExemptionCategoryPtrOutput() ExemptionCategoryPtrOutput {
	return o
}

func (o ExemptionCategoryPtrOutput) ToExemptionCategoryPtrOutputWithContext(ctx context.Context) ExemptionCategoryPtrOutput {
	return o
}

func (o ExemptionCategoryPtrOutput) Elem() ExemptionCategoryOutput {
	return o.ApplyT(func(v *ExemptionCategory) ExemptionCategory {
		if v != nil {
			return *v
		}
		var ret ExemptionCategory
		return ret
	}).(ExemptionCategoryOutput)
}

func (o ExemptionCategoryPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ExemptionCategoryPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ExemptionCategory) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ExemptionCategoryInput interface {
	pulumi.Input

	ToExemptionCategoryOutput() ExemptionCategoryOutput
	ToExemptionCategoryOutputWithContext(context.Context) ExemptionCategoryOutput
}

var exemptionCategoryPtrType = reflect.TypeOf((**ExemptionCategory)(nil)).Elem()

type ExemptionCategoryPtrInput interface {
	pulumi.Input

	ToExemptionCategoryPtrOutput() ExemptionCategoryPtrOutput
	ToExemptionCategoryPtrOutputWithContext(context.Context) ExemptionCategoryPtrOutput
}

type exemptionCategoryPtr string

func ExemptionCategoryPtr(v string) ExemptionCategoryPtrInput {
	return (*exemptionCategoryPtr)(&v)
}

func (*exemptionCategoryPtr) ElementType() reflect.Type {
	return exemptionCategoryPtrType
}

func (in *exemptionCategoryPtr) ToExemptionCategoryPtrOutput() ExemptionCategoryPtrOutput {
	return pulumi.ToOutput(in).(ExemptionCategoryPtrOutput)
}

func (in *exemptionCategoryPtr) ToExemptionCategoryPtrOutputWithContext(ctx context.Context) ExemptionCategoryPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ExemptionCategoryPtrOutput)
}

type PricingTier string

const (
	// The pricing tier gives the user ability to use policy exemption feature.
	PricingTierAdvanced = PricingTier("Advanced")
	// The pricing tier gives the user ability to use policy exemption feature. This pricing tier is managed by Azure Security Center.
	PricingTierDefender = PricingTier("Defender")
)

func (PricingTier) ElementType() reflect.Type {
	return reflect.TypeOf((*PricingTier)(nil)).Elem()
}

func (e PricingTier) ToPricingTierOutput() PricingTierOutput {
	return pulumi.ToOutput(e).(PricingTierOutput)
}

func (e PricingTier) ToPricingTierOutputWithContext(ctx context.Context) PricingTierOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PricingTierOutput)
}

func (e PricingTier) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return e.ToPricingTierPtrOutputWithContext(context.Background())
}

func (e PricingTier) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return PricingTier(e).ToPricingTierOutputWithContext(ctx).ToPricingTierPtrOutputWithContext(ctx)
}

func (e PricingTier) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PricingTier) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PricingTier) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PricingTier) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PricingTierOutput struct{ *pulumi.OutputState }

func (PricingTierOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PricingTier)(nil)).Elem()
}

func (o PricingTierOutput) ToPricingTierOutput() PricingTierOutput {
	return o
}

func (o PricingTierOutput) ToPricingTierOutputWithContext(ctx context.Context) PricingTierOutput {
	return o
}

func (o PricingTierOutput) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return o.ToPricingTierPtrOutputWithContext(context.Background())
}

func (o PricingTierOutput) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PricingTier) *PricingTier {
		return &v
	}).(PricingTierPtrOutput)
}

func (o PricingTierOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PricingTierOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PricingTier) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PricingTierOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PricingTierOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PricingTier) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PricingTierPtrOutput struct{ *pulumi.OutputState }

func (PricingTierPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PricingTier)(nil)).Elem()
}

func (o PricingTierPtrOutput) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return o
}

func (o PricingTierPtrOutput) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return o
}

func (o PricingTierPtrOutput) Elem() PricingTierOutput {
	return o.ApplyT(func(v *PricingTier) PricingTier {
		if v != nil {
			return *v
		}
		var ret PricingTier
		return ret
	}).(PricingTierOutput)
}

func (o PricingTierPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PricingTierPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PricingTier) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PricingTierInput interface {
	pulumi.Input

	ToPricingTierOutput() PricingTierOutput
	ToPricingTierOutputWithContext(context.Context) PricingTierOutput
}

var pricingTierPtrType = reflect.TypeOf((**PricingTier)(nil)).Elem()

type PricingTierPtrInput interface {
	pulumi.Input

	ToPricingTierPtrOutput() PricingTierPtrOutput
	ToPricingTierPtrOutputWithContext(context.Context) PricingTierPtrOutput
}

type pricingTierPtr string

func PricingTierPtr(v string) PricingTierPtrInput {
	return (*pricingTierPtr)(&v)
}

func (*pricingTierPtr) ElementType() reflect.Type {
	return pricingTierPtrType
}

func (in *pricingTierPtr) ToPricingTierPtrOutput() PricingTierPtrOutput {
	return pulumi.ToOutput(in).(PricingTierPtrOutput)
}

func (in *pricingTierPtr) ToPricingTierPtrOutputWithContext(ctx context.Context) PricingTierPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PricingTierPtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*ExemptionCategoryInput)(nil)).Elem(), ExemptionCategory("Waiver"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExemptionCategoryPtrInput)(nil)).Elem(), ExemptionCategory("Waiver"))
	pulumi.RegisterInputType(reflect.TypeOf((*PricingTierInput)(nil)).Elem(), PricingTier("Advanced"))
	pulumi.RegisterInputType(reflect.TypeOf((*PricingTierPtrInput)(nil)).Elem(), PricingTier("Advanced"))
	pulumi.RegisterOutputType(ExemptionCategoryOutput{})
	pulumi.RegisterOutputType(ExemptionCategoryPtrOutput{})
	pulumi.RegisterOutputType(PricingTierOutput{})
	pulumi.RegisterOutputType(PricingTierPtrOutput{})
}
