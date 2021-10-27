


package authorization

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AccessReviewRecurrencePatternType string

const (
	AccessReviewRecurrencePatternTypeWeekly          = AccessReviewRecurrencePatternType("weekly")
	AccessReviewRecurrencePatternTypeAbsoluteMonthly = AccessReviewRecurrencePatternType("absoluteMonthly")
)

func (AccessReviewRecurrencePatternType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewRecurrencePatternType)(nil)).Elem()
}

func (e AccessReviewRecurrencePatternType) ToAccessReviewRecurrencePatternTypeOutput() AccessReviewRecurrencePatternTypeOutput {
	return pulumi.ToOutput(e).(AccessReviewRecurrencePatternTypeOutput)
}

func (e AccessReviewRecurrencePatternType) ToAccessReviewRecurrencePatternTypeOutputWithContext(ctx context.Context) AccessReviewRecurrencePatternTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessReviewRecurrencePatternTypeOutput)
}

func (e AccessReviewRecurrencePatternType) ToAccessReviewRecurrencePatternTypePtrOutput() AccessReviewRecurrencePatternTypePtrOutput {
	return e.ToAccessReviewRecurrencePatternTypePtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrencePatternType) ToAccessReviewRecurrencePatternTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrencePatternTypePtrOutput {
	return AccessReviewRecurrencePatternType(e).ToAccessReviewRecurrencePatternTypeOutputWithContext(ctx).ToAccessReviewRecurrencePatternTypePtrOutputWithContext(ctx)
}

func (e AccessReviewRecurrencePatternType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrencePatternType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrencePatternType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrencePatternType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessReviewRecurrencePatternTypeOutput struct{ *pulumi.OutputState }

func (AccessReviewRecurrencePatternTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewRecurrencePatternType)(nil)).Elem()
}

func (o AccessReviewRecurrencePatternTypeOutput) ToAccessReviewRecurrencePatternTypeOutput() AccessReviewRecurrencePatternTypeOutput {
	return o
}

func (o AccessReviewRecurrencePatternTypeOutput) ToAccessReviewRecurrencePatternTypeOutputWithContext(ctx context.Context) AccessReviewRecurrencePatternTypeOutput {
	return o
}

func (o AccessReviewRecurrencePatternTypeOutput) ToAccessReviewRecurrencePatternTypePtrOutput() AccessReviewRecurrencePatternTypePtrOutput {
	return o.ToAccessReviewRecurrencePatternTypePtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrencePatternTypeOutput) ToAccessReviewRecurrencePatternTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrencePatternTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessReviewRecurrencePatternType) *AccessReviewRecurrencePatternType {
		return &v
	}).(AccessReviewRecurrencePatternTypePtrOutput)
}

func (o AccessReviewRecurrencePatternTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessReviewRecurrencePatternTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewRecurrencePatternType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessReviewRecurrencePatternTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrencePatternTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewRecurrencePatternType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessReviewRecurrencePatternTypePtrOutput struct{ *pulumi.OutputState }

func (AccessReviewRecurrencePatternTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewRecurrencePatternType)(nil)).Elem()
}

func (o AccessReviewRecurrencePatternTypePtrOutput) ToAccessReviewRecurrencePatternTypePtrOutput() AccessReviewRecurrencePatternTypePtrOutput {
	return o
}

func (o AccessReviewRecurrencePatternTypePtrOutput) ToAccessReviewRecurrencePatternTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrencePatternTypePtrOutput {
	return o
}

func (o AccessReviewRecurrencePatternTypePtrOutput) Elem() AccessReviewRecurrencePatternTypeOutput {
	return o.ApplyT(func(v *AccessReviewRecurrencePatternType) AccessReviewRecurrencePatternType {
		if v != nil {
			return *v
		}
		var ret AccessReviewRecurrencePatternType
		return ret
	}).(AccessReviewRecurrencePatternTypeOutput)
}

func (o AccessReviewRecurrencePatternTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrencePatternTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessReviewRecurrencePatternType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessReviewRecurrencePatternTypeInput interface {
	pulumi.Input

	ToAccessReviewRecurrencePatternTypeOutput() AccessReviewRecurrencePatternTypeOutput
	ToAccessReviewRecurrencePatternTypeOutputWithContext(context.Context) AccessReviewRecurrencePatternTypeOutput
}

var accessReviewRecurrencePatternTypePtrType = reflect.TypeOf((**AccessReviewRecurrencePatternType)(nil)).Elem()

type AccessReviewRecurrencePatternTypePtrInput interface {
	pulumi.Input

	ToAccessReviewRecurrencePatternTypePtrOutput() AccessReviewRecurrencePatternTypePtrOutput
	ToAccessReviewRecurrencePatternTypePtrOutputWithContext(context.Context) AccessReviewRecurrencePatternTypePtrOutput
}

type accessReviewRecurrencePatternTypePtr string

func AccessReviewRecurrencePatternTypePtr(v string) AccessReviewRecurrencePatternTypePtrInput {
	return (*accessReviewRecurrencePatternTypePtr)(&v)
}

func (*accessReviewRecurrencePatternTypePtr) ElementType() reflect.Type {
	return accessReviewRecurrencePatternTypePtrType
}

func (in *accessReviewRecurrencePatternTypePtr) ToAccessReviewRecurrencePatternTypePtrOutput() AccessReviewRecurrencePatternTypePtrOutput {
	return pulumi.ToOutput(in).(AccessReviewRecurrencePatternTypePtrOutput)
}

func (in *accessReviewRecurrencePatternTypePtr) ToAccessReviewRecurrencePatternTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrencePatternTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessReviewRecurrencePatternTypePtrOutput)
}

type AccessReviewRecurrenceRangeType string

const (
	AccessReviewRecurrenceRangeTypeEndDate  = AccessReviewRecurrenceRangeType("endDate")
	AccessReviewRecurrenceRangeTypeNoEnd    = AccessReviewRecurrenceRangeType("noEnd")
	AccessReviewRecurrenceRangeTypeNumbered = AccessReviewRecurrenceRangeType("numbered")
)

func (AccessReviewRecurrenceRangeType) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewRecurrenceRangeType)(nil)).Elem()
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypeOutput() AccessReviewRecurrenceRangeTypeOutput {
	return pulumi.ToOutput(e).(AccessReviewRecurrenceRangeTypeOutput)
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypeOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(AccessReviewRecurrenceRangeTypeOutput)
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return e.ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrenceRangeType) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return AccessReviewRecurrenceRangeType(e).ToAccessReviewRecurrenceRangeTypeOutputWithContext(ctx).ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx)
}

func (e AccessReviewRecurrenceRangeType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrenceRangeType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e AccessReviewRecurrenceRangeType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e AccessReviewRecurrenceRangeType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type AccessReviewRecurrenceRangeTypeOutput struct{ *pulumi.OutputState }

func (AccessReviewRecurrenceRangeTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AccessReviewRecurrenceRangeType)(nil)).Elem()
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypeOutput() AccessReviewRecurrenceRangeTypeOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypeOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypeOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return o.ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AccessReviewRecurrenceRangeType) *AccessReviewRecurrenceRangeType {
		return &v
	}).(AccessReviewRecurrenceRangeTypePtrOutput)
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewRecurrenceRangeType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e AccessReviewRecurrenceRangeType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type AccessReviewRecurrenceRangeTypePtrOutput struct{ *pulumi.OutputState }

func (AccessReviewRecurrenceRangeTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AccessReviewRecurrenceRangeType)(nil)).Elem()
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return o
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) Elem() AccessReviewRecurrenceRangeTypeOutput {
	return o.ApplyT(func(v *AccessReviewRecurrenceRangeType) AccessReviewRecurrenceRangeType {
		if v != nil {
			return *v
		}
		var ret AccessReviewRecurrenceRangeType
		return ret
	}).(AccessReviewRecurrenceRangeTypeOutput)
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o AccessReviewRecurrenceRangeTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *AccessReviewRecurrenceRangeType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type AccessReviewRecurrenceRangeTypeInput interface {
	pulumi.Input

	ToAccessReviewRecurrenceRangeTypeOutput() AccessReviewRecurrenceRangeTypeOutput
	ToAccessReviewRecurrenceRangeTypeOutputWithContext(context.Context) AccessReviewRecurrenceRangeTypeOutput
}

var accessReviewRecurrenceRangeTypePtrType = reflect.TypeOf((**AccessReviewRecurrenceRangeType)(nil)).Elem()

type AccessReviewRecurrenceRangeTypePtrInput interface {
	pulumi.Input

	ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput
	ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(context.Context) AccessReviewRecurrenceRangeTypePtrOutput
}

type accessReviewRecurrenceRangeTypePtr string

func AccessReviewRecurrenceRangeTypePtr(v string) AccessReviewRecurrenceRangeTypePtrInput {
	return (*accessReviewRecurrenceRangeTypePtr)(&v)
}

func (*accessReviewRecurrenceRangeTypePtr) ElementType() reflect.Type {
	return accessReviewRecurrenceRangeTypePtrType
}

func (in *accessReviewRecurrenceRangeTypePtr) ToAccessReviewRecurrenceRangeTypePtrOutput() AccessReviewRecurrenceRangeTypePtrOutput {
	return pulumi.ToOutput(in).(AccessReviewRecurrenceRangeTypePtrOutput)
}

func (in *accessReviewRecurrenceRangeTypePtr) ToAccessReviewRecurrenceRangeTypePtrOutputWithContext(ctx context.Context) AccessReviewRecurrenceRangeTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(AccessReviewRecurrenceRangeTypePtrOutput)
}

type DefaultDecisionType string

const (
	DefaultDecisionTypeApprove        = DefaultDecisionType("Approve")
	DefaultDecisionTypeDeny           = DefaultDecisionType("Deny")
	DefaultDecisionTypeRecommendation = DefaultDecisionType("Recommendation")
)

func (DefaultDecisionType) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultDecisionType)(nil)).Elem()
}

func (e DefaultDecisionType) ToDefaultDecisionTypeOutput() DefaultDecisionTypeOutput {
	return pulumi.ToOutput(e).(DefaultDecisionTypeOutput)
}

func (e DefaultDecisionType) ToDefaultDecisionTypeOutputWithContext(ctx context.Context) DefaultDecisionTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(DefaultDecisionTypeOutput)
}

func (e DefaultDecisionType) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return e.ToDefaultDecisionTypePtrOutputWithContext(context.Background())
}

func (e DefaultDecisionType) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return DefaultDecisionType(e).ToDefaultDecisionTypeOutputWithContext(ctx).ToDefaultDecisionTypePtrOutputWithContext(ctx)
}

func (e DefaultDecisionType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultDecisionType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e DefaultDecisionType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e DefaultDecisionType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type DefaultDecisionTypeOutput struct{ *pulumi.OutputState }

func (DefaultDecisionTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DefaultDecisionType)(nil)).Elem()
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypeOutput() DefaultDecisionTypeOutput {
	return o
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypeOutputWithContext(ctx context.Context) DefaultDecisionTypeOutput {
	return o
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return o.ToDefaultDecisionTypePtrOutputWithContext(context.Background())
}

func (o DefaultDecisionTypeOutput) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DefaultDecisionType) *DefaultDecisionType {
		return &v
	}).(DefaultDecisionTypePtrOutput)
}

func (o DefaultDecisionTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o DefaultDecisionTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultDecisionType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o DefaultDecisionTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultDecisionTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e DefaultDecisionType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type DefaultDecisionTypePtrOutput struct{ *pulumi.OutputState }

func (DefaultDecisionTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DefaultDecisionType)(nil)).Elem()
}

func (o DefaultDecisionTypePtrOutput) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return o
}

func (o DefaultDecisionTypePtrOutput) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return o
}

func (o DefaultDecisionTypePtrOutput) Elem() DefaultDecisionTypeOutput {
	return o.ApplyT(func(v *DefaultDecisionType) DefaultDecisionType {
		if v != nil {
			return *v
		}
		var ret DefaultDecisionType
		return ret
	}).(DefaultDecisionTypeOutput)
}

func (o DefaultDecisionTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o DefaultDecisionTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *DefaultDecisionType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type DefaultDecisionTypeInput interface {
	pulumi.Input

	ToDefaultDecisionTypeOutput() DefaultDecisionTypeOutput
	ToDefaultDecisionTypeOutputWithContext(context.Context) DefaultDecisionTypeOutput
}

var defaultDecisionTypePtrType = reflect.TypeOf((**DefaultDecisionType)(nil)).Elem()

type DefaultDecisionTypePtrInput interface {
	pulumi.Input

	ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput
	ToDefaultDecisionTypePtrOutputWithContext(context.Context) DefaultDecisionTypePtrOutput
}

type defaultDecisionTypePtr string

func DefaultDecisionTypePtr(v string) DefaultDecisionTypePtrInput {
	return (*defaultDecisionTypePtr)(&v)
}

func (*defaultDecisionTypePtr) ElementType() reflect.Type {
	return defaultDecisionTypePtrType
}

func (in *defaultDecisionTypePtr) ToDefaultDecisionTypePtrOutput() DefaultDecisionTypePtrOutput {
	return pulumi.ToOutput(in).(DefaultDecisionTypePtrOutput)
}

func (in *defaultDecisionTypePtr) ToDefaultDecisionTypePtrOutputWithContext(ctx context.Context) DefaultDecisionTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(DefaultDecisionTypePtrOutput)
}

type EnforcementMode string

const (
	// The policy effect is enforced during resource creation or update.
	EnforcementModeDefault = EnforcementMode("Default")
	// The policy effect is not enforced during resource creation or update.
	EnforcementModeDoNotEnforce = EnforcementMode("DoNotEnforce")
)

func (EnforcementMode) ElementType() reflect.Type {
	return reflect.TypeOf((*EnforcementMode)(nil)).Elem()
}

func (e EnforcementMode) ToEnforcementModeOutput() EnforcementModeOutput {
	return pulumi.ToOutput(e).(EnforcementModeOutput)
}

func (e EnforcementMode) ToEnforcementModeOutputWithContext(ctx context.Context) EnforcementModeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(EnforcementModeOutput)
}

func (e EnforcementMode) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return e.ToEnforcementModePtrOutputWithContext(context.Background())
}

func (e EnforcementMode) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return EnforcementMode(e).ToEnforcementModeOutputWithContext(ctx).ToEnforcementModePtrOutputWithContext(ctx)
}

func (e EnforcementMode) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnforcementMode) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e EnforcementMode) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e EnforcementMode) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type EnforcementModeOutput struct{ *pulumi.OutputState }

func (EnforcementModeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnforcementMode)(nil)).Elem()
}

func (o EnforcementModeOutput) ToEnforcementModeOutput() EnforcementModeOutput {
	return o
}

func (o EnforcementModeOutput) ToEnforcementModeOutputWithContext(ctx context.Context) EnforcementModeOutput {
	return o
}

func (o EnforcementModeOutput) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return o.ToEnforcementModePtrOutputWithContext(context.Background())
}

func (o EnforcementModeOutput) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnforcementMode) *EnforcementMode {
		return &v
	}).(EnforcementModePtrOutput)
}

func (o EnforcementModeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o EnforcementModeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnforcementMode) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o EnforcementModeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnforcementModeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e EnforcementMode) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type EnforcementModePtrOutput struct{ *pulumi.OutputState }

func (EnforcementModePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnforcementMode)(nil)).Elem()
}

func (o EnforcementModePtrOutput) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return o
}

func (o EnforcementModePtrOutput) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return o
}

func (o EnforcementModePtrOutput) Elem() EnforcementModeOutput {
	return o.ApplyT(func(v *EnforcementMode) EnforcementMode {
		if v != nil {
			return *v
		}
		var ret EnforcementMode
		return ret
	}).(EnforcementModeOutput)
}

func (o EnforcementModePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o EnforcementModePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *EnforcementMode) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type EnforcementModeInput interface {
	pulumi.Input

	ToEnforcementModeOutput() EnforcementModeOutput
	ToEnforcementModeOutputWithContext(context.Context) EnforcementModeOutput
}

var enforcementModePtrType = reflect.TypeOf((**EnforcementMode)(nil)).Elem()

type EnforcementModePtrInput interface {
	pulumi.Input

	ToEnforcementModePtrOutput() EnforcementModePtrOutput
	ToEnforcementModePtrOutputWithContext(context.Context) EnforcementModePtrOutput
}

type enforcementModePtr string

func EnforcementModePtr(v string) EnforcementModePtrInput {
	return (*enforcementModePtr)(&v)
}

func (*enforcementModePtr) ElementType() reflect.Type {
	return enforcementModePtrType
}

func (in *enforcementModePtr) ToEnforcementModePtrOutput() EnforcementModePtrOutput {
	return pulumi.ToOutput(in).(EnforcementModePtrOutput)
}

func (in *enforcementModePtr) ToEnforcementModePtrOutputWithContext(ctx context.Context) EnforcementModePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(EnforcementModePtrOutput)
}

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

type LockLevel string

const (
	LockLevelNotSpecified = LockLevel("NotSpecified")
	LockLevelCanNotDelete = LockLevel("CanNotDelete")
	LockLevelReadOnly     = LockLevel("ReadOnly")
)

func (LockLevel) ElementType() reflect.Type {
	return reflect.TypeOf((*LockLevel)(nil)).Elem()
}

func (e LockLevel) ToLockLevelOutput() LockLevelOutput {
	return pulumi.ToOutput(e).(LockLevelOutput)
}

func (e LockLevel) ToLockLevelOutputWithContext(ctx context.Context) LockLevelOutput {
	return pulumi.ToOutputWithContext(ctx, e).(LockLevelOutput)
}

func (e LockLevel) ToLockLevelPtrOutput() LockLevelPtrOutput {
	return e.ToLockLevelPtrOutputWithContext(context.Background())
}

func (e LockLevel) ToLockLevelPtrOutputWithContext(ctx context.Context) LockLevelPtrOutput {
	return LockLevel(e).ToLockLevelOutputWithContext(ctx).ToLockLevelPtrOutputWithContext(ctx)
}

func (e LockLevel) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e LockLevel) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e LockLevel) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e LockLevel) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type LockLevelOutput struct{ *pulumi.OutputState }

func (LockLevelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LockLevel)(nil)).Elem()
}

func (o LockLevelOutput) ToLockLevelOutput() LockLevelOutput {
	return o
}

func (o LockLevelOutput) ToLockLevelOutputWithContext(ctx context.Context) LockLevelOutput {
	return o
}

func (o LockLevelOutput) ToLockLevelPtrOutput() LockLevelPtrOutput {
	return o.ToLockLevelPtrOutputWithContext(context.Background())
}

func (o LockLevelOutput) ToLockLevelPtrOutputWithContext(ctx context.Context) LockLevelPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LockLevel) *LockLevel {
		return &v
	}).(LockLevelPtrOutput)
}

func (o LockLevelOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o LockLevelOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LockLevel) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o LockLevelOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LockLevelOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e LockLevel) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type LockLevelPtrOutput struct{ *pulumi.OutputState }

func (LockLevelPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LockLevel)(nil)).Elem()
}

func (o LockLevelPtrOutput) ToLockLevelPtrOutput() LockLevelPtrOutput {
	return o
}

func (o LockLevelPtrOutput) ToLockLevelPtrOutputWithContext(ctx context.Context) LockLevelPtrOutput {
	return o
}

func (o LockLevelPtrOutput) Elem() LockLevelOutput {
	return o.ApplyT(func(v *LockLevel) LockLevel {
		if v != nil {
			return *v
		}
		var ret LockLevel
		return ret
	}).(LockLevelOutput)
}

func (o LockLevelPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o LockLevelPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *LockLevel) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type LockLevelInput interface {
	pulumi.Input

	ToLockLevelOutput() LockLevelOutput
	ToLockLevelOutputWithContext(context.Context) LockLevelOutput
}

var lockLevelPtrType = reflect.TypeOf((**LockLevel)(nil)).Elem()

type LockLevelPtrInput interface {
	pulumi.Input

	ToLockLevelPtrOutput() LockLevelPtrOutput
	ToLockLevelPtrOutputWithContext(context.Context) LockLevelPtrOutput
}

type lockLevelPtr string

func LockLevelPtr(v string) LockLevelPtrInput {
	return (*lockLevelPtr)(&v)
}

func (*lockLevelPtr) ElementType() reflect.Type {
	return lockLevelPtrType
}

func (in *lockLevelPtr) ToLockLevelPtrOutput() LockLevelPtrOutput {
	return pulumi.ToOutput(in).(LockLevelPtrOutput)
}

func (in *lockLevelPtr) ToLockLevelPtrOutputWithContext(ctx context.Context) LockLevelPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(LockLevelPtrOutput)
}

type ParameterType string

const (
	ParameterTypeString   = ParameterType("String")
	ParameterTypeArray    = ParameterType("Array")
	ParameterTypeObject   = ParameterType("Object")
	ParameterTypeBoolean  = ParameterType("Boolean")
	ParameterTypeInteger  = ParameterType("Integer")
	ParameterTypeFloat    = ParameterType("Float")
	ParameterTypeDateTime = ParameterType("DateTime")
)

func (ParameterType) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (e ParameterType) ToParameterTypeOutput() ParameterTypeOutput {
	return pulumi.ToOutput(e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ParameterTypeOutput)
}

func (e ParameterType) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return e.ToParameterTypePtrOutputWithContext(context.Background())
}

func (e ParameterType) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return ParameterType(e).ToParameterTypeOutputWithContext(ctx).ToParameterTypePtrOutputWithContext(ctx)
}

func (e ParameterType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ParameterType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ParameterType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ParameterTypeOutput struct{ *pulumi.OutputState }

func (ParameterTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ParameterType)(nil)).Elem()
}

func (o ParameterTypeOutput) ToParameterTypeOutput() ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypeOutputWithContext(ctx context.Context) ParameterTypeOutput {
	return o
}

func (o ParameterTypeOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o.ToParameterTypePtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ParameterType) *ParameterType {
		return &v
	}).(ParameterTypePtrOutput)
}

func (o ParameterTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ParameterTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ParameterType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ParameterTypePtrOutput struct{ *pulumi.OutputState }

func (ParameterTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ParameterType)(nil)).Elem()
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return o
}

func (o ParameterTypePtrOutput) Elem() ParameterTypeOutput {
	return o.ApplyT(func(v *ParameterType) ParameterType {
		if v != nil {
			return *v
		}
		var ret ParameterType
		return ret
	}).(ParameterTypeOutput)
}

func (o ParameterTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ParameterTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ParameterType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ParameterTypeInput interface {
	pulumi.Input

	ToParameterTypeOutput() ParameterTypeOutput
	ToParameterTypeOutputWithContext(context.Context) ParameterTypeOutput
}

var parameterTypePtrType = reflect.TypeOf((**ParameterType)(nil)).Elem()

type ParameterTypePtrInput interface {
	pulumi.Input

	ToParameterTypePtrOutput() ParameterTypePtrOutput
	ToParameterTypePtrOutputWithContext(context.Context) ParameterTypePtrOutput
}

type parameterTypePtr string

func ParameterTypePtr(v string) ParameterTypePtrInput {
	return (*parameterTypePtr)(&v)
}

func (*parameterTypePtr) ElementType() reflect.Type {
	return parameterTypePtrType
}

func (in *parameterTypePtr) ToParameterTypePtrOutput() ParameterTypePtrOutput {
	return pulumi.ToOutput(in).(ParameterTypePtrOutput)
}

func (in *parameterTypePtr) ToParameterTypePtrOutputWithContext(ctx context.Context) ParameterTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ParameterTypePtrOutput)
}

type PolicyType string

const (
	PolicyTypeNotSpecified = PolicyType("NotSpecified")
	PolicyTypeBuiltIn      = PolicyType("BuiltIn")
	PolicyTypeCustom       = PolicyType("Custom")
	PolicyTypeStatic       = PolicyType("Static")
)

func (PolicyType) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (e PolicyType) ToPolicyTypeOutput() PolicyTypeOutput {
	return pulumi.ToOutput(e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PolicyTypeOutput)
}

func (e PolicyType) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return e.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (e PolicyType) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return PolicyType(e).ToPolicyTypeOutputWithContext(ctx).ToPolicyTypePtrOutputWithContext(ctx)
}

func (e PolicyType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PolicyType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PolicyType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PolicyTypeOutput struct{ *pulumi.OutputState }

func (PolicyTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PolicyType)(nil)).Elem()
}

func (o PolicyTypeOutput) ToPolicyTypeOutput() PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypeOutputWithContext(ctx context.Context) PolicyTypeOutput {
	return o
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o.ToPolicyTypePtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PolicyType) *PolicyType {
		return &v
	}).(PolicyTypePtrOutput)
}

func (o PolicyTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PolicyTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PolicyType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PolicyTypePtrOutput struct{ *pulumi.OutputState }

func (PolicyTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PolicyType)(nil)).Elem()
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return o
}

func (o PolicyTypePtrOutput) Elem() PolicyTypeOutput {
	return o.ApplyT(func(v *PolicyType) PolicyType {
		if v != nil {
			return *v
		}
		var ret PolicyType
		return ret
	}).(PolicyTypeOutput)
}

func (o PolicyTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PolicyTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PolicyType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PolicyTypeInput interface {
	pulumi.Input

	ToPolicyTypeOutput() PolicyTypeOutput
	ToPolicyTypeOutputWithContext(context.Context) PolicyTypeOutput
}

var policyTypePtrType = reflect.TypeOf((**PolicyType)(nil)).Elem()

type PolicyTypePtrInput interface {
	pulumi.Input

	ToPolicyTypePtrOutput() PolicyTypePtrOutput
	ToPolicyTypePtrOutputWithContext(context.Context) PolicyTypePtrOutput
}

type policyTypePtr string

func PolicyTypePtr(v string) PolicyTypePtrInput {
	return (*policyTypePtr)(&v)
}

func (*policyTypePtr) ElementType() reflect.Type {
	return policyTypePtrType
}

func (in *policyTypePtr) ToPolicyTypePtrOutput() PolicyTypePtrOutput {
	return pulumi.ToOutput(in).(PolicyTypePtrOutput)
}

func (in *policyTypePtr) ToPolicyTypePtrOutputWithContext(ctx context.Context) PolicyTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PolicyTypePtrOutput)
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

type PrincipalType string

const (
	PrincipalTypeUser             = PrincipalType("User")
	PrincipalTypeGroup            = PrincipalType("Group")
	PrincipalTypeServicePrincipal = PrincipalType("ServicePrincipal")
	PrincipalTypeForeignGroup     = PrincipalType("ForeignGroup")
)

func (PrincipalType) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (e PrincipalType) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return pulumi.ToOutput(e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(PrincipalTypeOutput)
}

func (e PrincipalType) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return e.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return PrincipalType(e).ToPrincipalTypeOutputWithContext(ctx).ToPrincipalTypePtrOutputWithContext(ctx)
}

func (e PrincipalType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e PrincipalType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e PrincipalType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type PrincipalTypeOutput struct{ *pulumi.OutputState }

func (PrincipalTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrincipalType)(nil)).Elem()
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutput() PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypeOutputWithContext(ctx context.Context) PrincipalTypeOutput {
	return o
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o.ToPrincipalTypePtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrincipalType) *PrincipalType {
		return &v
	}).(PrincipalTypePtrOutput)
}

func (o PrincipalTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o PrincipalTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e PrincipalType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type PrincipalTypePtrOutput struct{ *pulumi.OutputState }

func (PrincipalTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrincipalType)(nil)).Elem()
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return o
}

func (o PrincipalTypePtrOutput) Elem() PrincipalTypeOutput {
	return o.ApplyT(func(v *PrincipalType) PrincipalType {
		if v != nil {
			return *v
		}
		var ret PrincipalType
		return ret
	}).(PrincipalTypeOutput)
}

func (o PrincipalTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o PrincipalTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *PrincipalType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type PrincipalTypeInput interface {
	pulumi.Input

	ToPrincipalTypeOutput() PrincipalTypeOutput
	ToPrincipalTypeOutputWithContext(context.Context) PrincipalTypeOutput
}

var principalTypePtrType = reflect.TypeOf((**PrincipalType)(nil)).Elem()

type PrincipalTypePtrInput interface {
	pulumi.Input

	ToPrincipalTypePtrOutput() PrincipalTypePtrOutput
	ToPrincipalTypePtrOutputWithContext(context.Context) PrincipalTypePtrOutput
}

type principalTypePtr string

func PrincipalTypePtr(v string) PrincipalTypePtrInput {
	return (*principalTypePtr)(&v)
}

func (*principalTypePtr) ElementType() reflect.Type {
	return principalTypePtrType
}

func (in *principalTypePtr) ToPrincipalTypePtrOutput() PrincipalTypePtrOutput {
	return pulumi.ToOutput(in).(PrincipalTypePtrOutput)
}

func (in *principalTypePtr) ToPrincipalTypePtrOutputWithContext(ctx context.Context) PrincipalTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(PrincipalTypePtrOutput)
}

type ResourceIdentityType string

const (
	// Indicates that a system assigned identity is associated with the resource.
	ResourceIdentityTypeSystemAssigned = ResourceIdentityType("SystemAssigned")
	// Indicates that no identity is associated with the resource or that the existing identity should be removed.
	ResourceIdentityTypeNone = ResourceIdentityType("None")
)

func (ResourceIdentityType) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return pulumi.ToOutput(e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return pulumi.ToOutputWithContext(ctx, e).(ResourceIdentityTypeOutput)
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return e.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return ResourceIdentityType(e).ToResourceIdentityTypeOutputWithContext(ctx).ToResourceIdentityTypePtrOutputWithContext(ctx)
}

func (e ResourceIdentityType) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e ResourceIdentityType) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e ResourceIdentityType) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type ResourceIdentityTypeOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypeOutputWithContext(ctx context.Context) ResourceIdentityTypeOutput {
	return o
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o.ToResourceIdentityTypePtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentityType) *ResourceIdentityType {
		return &v
	}).(ResourceIdentityTypePtrOutput)
}

func (o ResourceIdentityTypeOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypeOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e ResourceIdentityType) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityTypePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityTypePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return o
}

func (o ResourceIdentityTypePtrOutput) Elem() ResourceIdentityTypeOutput {
	return o.ApplyT(func(v *ResourceIdentityType) ResourceIdentityType {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityType
		return ret
	}).(ResourceIdentityTypeOutput)
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityTypePtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *ResourceIdentityType) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type ResourceIdentityTypeInput interface {
	pulumi.Input

	ToResourceIdentityTypeOutput() ResourceIdentityTypeOutput
	ToResourceIdentityTypeOutputWithContext(context.Context) ResourceIdentityTypeOutput
}

var resourceIdentityTypePtrType = reflect.TypeOf((**ResourceIdentityType)(nil)).Elem()

type ResourceIdentityTypePtrInput interface {
	pulumi.Input

	ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput
	ToResourceIdentityTypePtrOutputWithContext(context.Context) ResourceIdentityTypePtrOutput
}

type resourceIdentityTypePtr string

func ResourceIdentityTypePtr(v string) ResourceIdentityTypePtrInput {
	return (*resourceIdentityTypePtr)(&v)
}

func (*resourceIdentityTypePtr) ElementType() reflect.Type {
	return resourceIdentityTypePtrType
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutput() ResourceIdentityTypePtrOutput {
	return pulumi.ToOutput(in).(ResourceIdentityTypePtrOutput)
}

func (in *resourceIdentityTypePtr) ToResourceIdentityTypePtrOutputWithContext(ctx context.Context) ResourceIdentityTypePtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(ResourceIdentityTypePtrOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrencePatternTypeInput)(nil)).Elem(), AccessReviewRecurrencePatternType("weekly"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrencePatternTypePtrInput)(nil)).Elem(), AccessReviewRecurrencePatternType("weekly"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrenceRangeTypeInput)(nil)).Elem(), AccessReviewRecurrenceRangeType("endDate"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrenceRangeTypePtrInput)(nil)).Elem(), AccessReviewRecurrenceRangeType("endDate"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDecisionTypeInput)(nil)).Elem(), DefaultDecisionType("Approve"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDecisionTypePtrInput)(nil)).Elem(), DefaultDecisionType("Approve"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnforcementModeInput)(nil)).Elem(), EnforcementMode("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*EnforcementModePtrInput)(nil)).Elem(), EnforcementMode("Default"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExemptionCategoryInput)(nil)).Elem(), ExemptionCategory("Waiver"))
	pulumi.RegisterInputType(reflect.TypeOf((*ExemptionCategoryPtrInput)(nil)).Elem(), ExemptionCategory("Waiver"))
	pulumi.RegisterInputType(reflect.TypeOf((*LockLevelInput)(nil)).Elem(), LockLevel("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*LockLevelPtrInput)(nil)).Elem(), LockLevel("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTypeInput)(nil)).Elem(), ParameterType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*ParameterTypePtrInput)(nil)).Elem(), ParameterType("String"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTypeInput)(nil)).Elem(), PolicyType("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*PolicyTypePtrInput)(nil)).Elem(), PolicyType("NotSpecified"))
	pulumi.RegisterInputType(reflect.TypeOf((*PricingTierInput)(nil)).Elem(), PricingTier("Advanced"))
	pulumi.RegisterInputType(reflect.TypeOf((*PricingTierPtrInput)(nil)).Elem(), PricingTier("Advanced"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalTypeInput)(nil)).Elem(), PrincipalType("User"))
	pulumi.RegisterInputType(reflect.TypeOf((*PrincipalTypePtrInput)(nil)).Elem(), PrincipalType("User"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypeInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterInputType(reflect.TypeOf((*ResourceIdentityTypePtrInput)(nil)).Elem(), ResourceIdentityType("SystemAssigned"))
	pulumi.RegisterOutputType(AccessReviewRecurrencePatternTypeOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrencePatternTypePtrOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrenceRangeTypeOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrenceRangeTypePtrOutput{})
	pulumi.RegisterOutputType(DefaultDecisionTypeOutput{})
	pulumi.RegisterOutputType(DefaultDecisionTypePtrOutput{})
	pulumi.RegisterOutputType(EnforcementModeOutput{})
	pulumi.RegisterOutputType(EnforcementModePtrOutput{})
	pulumi.RegisterOutputType(ExemptionCategoryOutput{})
	pulumi.RegisterOutputType(ExemptionCategoryPtrOutput{})
	pulumi.RegisterOutputType(LockLevelOutput{})
	pulumi.RegisterOutputType(LockLevelPtrOutput{})
	pulumi.RegisterOutputType(ParameterTypeOutput{})
	pulumi.RegisterOutputType(ParameterTypePtrOutput{})
	pulumi.RegisterOutputType(PolicyTypeOutput{})
	pulumi.RegisterOutputType(PolicyTypePtrOutput{})
	pulumi.RegisterOutputType(PricingTierOutput{})
	pulumi.RegisterOutputType(PricingTierPtrOutput{})
	pulumi.RegisterOutputType(PrincipalTypeOutput{})
	pulumi.RegisterOutputType(PrincipalTypePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypeOutput{})
	pulumi.RegisterOutputType(ResourceIdentityTypePtrOutput{})
}
