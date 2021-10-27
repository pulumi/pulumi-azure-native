


package v20180501preview

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

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrencePatternTypeInput)(nil)).Elem(), AccessReviewRecurrencePatternType("weekly"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrencePatternTypePtrInput)(nil)).Elem(), AccessReviewRecurrencePatternType("weekly"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrenceRangeTypeInput)(nil)).Elem(), AccessReviewRecurrenceRangeType("endDate"))
	pulumi.RegisterInputType(reflect.TypeOf((*AccessReviewRecurrenceRangeTypePtrInput)(nil)).Elem(), AccessReviewRecurrenceRangeType("endDate"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDecisionTypeInput)(nil)).Elem(), DefaultDecisionType("Approve"))
	pulumi.RegisterInputType(reflect.TypeOf((*DefaultDecisionTypePtrInput)(nil)).Elem(), DefaultDecisionType("Approve"))
	pulumi.RegisterOutputType(AccessReviewRecurrencePatternTypeOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrencePatternTypePtrOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrenceRangeTypeOutput{})
	pulumi.RegisterOutputType(AccessReviewRecurrenceRangeTypePtrOutput{})
	pulumi.RegisterOutputType(DefaultDecisionTypeOutput{})
	pulumi.RegisterOutputType(DefaultDecisionTypePtrOutput{})
}
