


package v20220615

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConditionOperator string

const (
	ConditionOperatorEquals             = ConditionOperator("Equals")
	ConditionOperatorGreaterThan        = ConditionOperator("GreaterThan")
	ConditionOperatorGreaterThanOrEqual = ConditionOperator("GreaterThanOrEqual")
	ConditionOperatorLessThan           = ConditionOperator("LessThan")
	ConditionOperatorLessThanOrEqual    = ConditionOperator("LessThanOrEqual")
)

type DimensionOperator string

const (
	DimensionOperatorInclude = DimensionOperator("Include")
	DimensionOperatorExclude = DimensionOperator("Exclude")
)

type Kind string

const (
	KindLogAlert    = Kind("LogAlert")
	KindLogToMetric = Kind("LogToMetric")
)

type TimeAggregation string

const (
	TimeAggregationCount   = TimeAggregation("Count")
	TimeAggregationAverage = TimeAggregation("Average")
	TimeAggregationMinimum = TimeAggregation("Minimum")
	TimeAggregationMaximum = TimeAggregation("Maximum")
	TimeAggregationTotal   = TimeAggregation("Total")
)

type WebTestKind string

const (
	WebTestKindPing      = WebTestKind("ping")
	WebTestKindMultistep = WebTestKind("multistep")
	WebTestKindStandard  = WebTestKind("standard")
)

func (WebTestKind) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestKind)(nil)).Elem()
}

func (e WebTestKind) ToWebTestKindOutput() WebTestKindOutput {
	return pulumi.ToOutput(e).(WebTestKindOutput)
}

func (e WebTestKind) ToWebTestKindOutputWithContext(ctx context.Context) WebTestKindOutput {
	return pulumi.ToOutputWithContext(ctx, e).(WebTestKindOutput)
}

func (e WebTestKind) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return e.ToWebTestKindPtrOutputWithContext(context.Background())
}

func (e WebTestKind) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return WebTestKind(e).ToWebTestKindOutputWithContext(ctx).ToWebTestKindPtrOutputWithContext(ctx)
}

func (e WebTestKind) ToStringOutput() pulumi.StringOutput {
	return pulumi.ToOutput(pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebTestKind) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return pulumi.ToOutputWithContext(ctx, pulumi.String(e)).(pulumi.StringOutput)
}

func (e WebTestKind) ToStringPtrOutput() pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringPtrOutputWithContext(context.Background())
}

func (e WebTestKind) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return pulumi.String(e).ToStringOutputWithContext(ctx).ToStringPtrOutputWithContext(ctx)
}

type WebTestKindOutput struct{ *pulumi.OutputState }

func (WebTestKindOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebTestKind)(nil)).Elem()
}

func (o WebTestKindOutput) ToWebTestKindOutput() WebTestKindOutput {
	return o
}

func (o WebTestKindOutput) ToWebTestKindOutputWithContext(ctx context.Context) WebTestKindOutput {
	return o
}

func (o WebTestKindOutput) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return o.ToWebTestKindPtrOutputWithContext(context.Background())
}

func (o WebTestKindOutput) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebTestKind) *WebTestKind {
		return &v
	}).(WebTestKindPtrOutput)
}

func (o WebTestKindOutput) ToStringOutput() pulumi.StringOutput {
	return o.ToStringOutputWithContext(context.Background())
}

func (o WebTestKindOutput) ToStringOutputWithContext(ctx context.Context) pulumi.StringOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebTestKind) string {
		return string(e)
	}).(pulumi.StringOutput)
}

func (o WebTestKindOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebTestKindOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e WebTestKind) *string {
		v := string(e)
		return &v
	}).(pulumi.StringPtrOutput)
}

type WebTestKindPtrOutput struct{ *pulumi.OutputState }

func (WebTestKindPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebTestKind)(nil)).Elem()
}

func (o WebTestKindPtrOutput) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return o
}

func (o WebTestKindPtrOutput) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return o
}

func (o WebTestKindPtrOutput) Elem() WebTestKindOutput {
	return o.ApplyT(func(v *WebTestKind) WebTestKind {
		if v != nil {
			return *v
		}
		var ret WebTestKind
		return ret
	}).(WebTestKindOutput)
}

func (o WebTestKindPtrOutput) ToStringPtrOutput() pulumi.StringPtrOutput {
	return o.ToStringPtrOutputWithContext(context.Background())
}

func (o WebTestKindPtrOutput) ToStringPtrOutputWithContext(ctx context.Context) pulumi.StringPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, e *WebTestKind) *string {
		if e == nil {
			return nil
		}
		v := string(*e)
		return &v
	}).(pulumi.StringPtrOutput)
}





type WebTestKindInput interface {
	pulumi.Input

	ToWebTestKindOutput() WebTestKindOutput
	ToWebTestKindOutputWithContext(context.Context) WebTestKindOutput
}

var webTestKindPtrType = reflect.TypeOf((**WebTestKind)(nil)).Elem()

type WebTestKindPtrInput interface {
	pulumi.Input

	ToWebTestKindPtrOutput() WebTestKindPtrOutput
	ToWebTestKindPtrOutputWithContext(context.Context) WebTestKindPtrOutput
}

type webTestKindPtr string

func WebTestKindPtr(v string) WebTestKindPtrInput {
	return (*webTestKindPtr)(&v)
}

func (*webTestKindPtr) ElementType() reflect.Type {
	return webTestKindPtrType
}

func (in *webTestKindPtr) ToWebTestKindPtrOutput() WebTestKindPtrOutput {
	return pulumi.ToOutput(in).(WebTestKindPtrOutput)
}

func (in *webTestKindPtr) ToWebTestKindPtrOutputWithContext(ctx context.Context) WebTestKindPtrOutput {
	return pulumi.ToOutputWithContext(ctx, in).(WebTestKindPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebTestKindOutput{})
	pulumi.RegisterOutputType(WebTestKindPtrOutput{})
}
