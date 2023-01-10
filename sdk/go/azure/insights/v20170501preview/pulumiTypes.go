


package v20170501preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type LogSettings struct {
	Category        *string          `pulumi:"category"`
	Enabled         bool             `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicy `pulumi:"retentionPolicy"`
}





type LogSettingsInput interface {
	pulumi.Input

	ToLogSettingsOutput() LogSettingsOutput
	ToLogSettingsOutputWithContext(context.Context) LogSettingsOutput
}

type LogSettingsArgs struct {
	Category        pulumi.StringPtrInput   `pulumi:"category"`
	Enabled         pulumi.BoolInput        `pulumi:"enabled"`
	RetentionPolicy RetentionPolicyPtrInput `pulumi:"retentionPolicy"`
}

func (LogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettings)(nil)).Elem()
}

func (i LogSettingsArgs) ToLogSettingsOutput() LogSettingsOutput {
	return i.ToLogSettingsOutputWithContext(context.Background())
}

func (i LogSettingsArgs) ToLogSettingsOutputWithContext(ctx context.Context) LogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsOutput)
}





type LogSettingsArrayInput interface {
	pulumi.Input

	ToLogSettingsArrayOutput() LogSettingsArrayOutput
	ToLogSettingsArrayOutputWithContext(context.Context) LogSettingsArrayOutput
}

type LogSettingsArray []LogSettingsInput

func (LogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettings)(nil)).Elem()
}

func (i LogSettingsArray) ToLogSettingsArrayOutput() LogSettingsArrayOutput {
	return i.ToLogSettingsArrayOutputWithContext(context.Background())
}

func (i LogSettingsArray) ToLogSettingsArrayOutputWithContext(ctx context.Context) LogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogSettingsArrayOutput)
}

type LogSettingsOutput struct{ *pulumi.OutputState }

func (LogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettings)(nil)).Elem()
}

func (o LogSettingsOutput) ToLogSettingsOutput() LogSettingsOutput {
	return o
}

func (o LogSettingsOutput) ToLogSettingsOutputWithContext(ctx context.Context) LogSettingsOutput {
	return o
}

func (o LogSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LogSettingsOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v LogSettings) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

type LogSettingsArrayOutput struct{ *pulumi.OutputState }

func (LogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettings)(nil)).Elem()
}

func (o LogSettingsArrayOutput) ToLogSettingsArrayOutput() LogSettingsArrayOutput {
	return o
}

func (o LogSettingsArrayOutput) ToLogSettingsArrayOutputWithContext(ctx context.Context) LogSettingsArrayOutput {
	return o
}

func (o LogSettingsArrayOutput) Index(i pulumi.IntInput) LogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSettings {
		return vs[0].([]LogSettings)[vs[1].(int)]
	}).(LogSettingsOutput)
}

type LogSettingsResponse struct {
	Category        *string                  `pulumi:"category"`
	Enabled         bool                     `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicyResponse `pulumi:"retentionPolicy"`
}

type LogSettingsResponseOutput struct{ *pulumi.OutputState }

func (LogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogSettingsResponse)(nil)).Elem()
}

func (o LogSettingsResponseOutput) ToLogSettingsResponseOutput() LogSettingsResponseOutput {
	return o
}

func (o LogSettingsResponseOutput) ToLogSettingsResponseOutputWithContext(ctx context.Context) LogSettingsResponseOutput {
	return o
}

func (o LogSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o LogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v LogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o LogSettingsResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v LogSettingsResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

type LogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (LogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LogSettingsResponse)(nil)).Elem()
}

func (o LogSettingsResponseArrayOutput) ToLogSettingsResponseArrayOutput() LogSettingsResponseArrayOutput {
	return o
}

func (o LogSettingsResponseArrayOutput) ToLogSettingsResponseArrayOutputWithContext(ctx context.Context) LogSettingsResponseArrayOutput {
	return o
}

func (o LogSettingsResponseArrayOutput) Index(i pulumi.IntInput) LogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LogSettingsResponse {
		return vs[0].([]LogSettingsResponse)[vs[1].(int)]
	}).(LogSettingsResponseOutput)
}

type MetricSettings struct {
	Category        *string          `pulumi:"category"`
	Enabled         bool             `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicy `pulumi:"retentionPolicy"`
	TimeGrain       *string          `pulumi:"timeGrain"`
}





type MetricSettingsInput interface {
	pulumi.Input

	ToMetricSettingsOutput() MetricSettingsOutput
	ToMetricSettingsOutputWithContext(context.Context) MetricSettingsOutput
}

type MetricSettingsArgs struct {
	Category        pulumi.StringPtrInput   `pulumi:"category"`
	Enabled         pulumi.BoolInput        `pulumi:"enabled"`
	RetentionPolicy RetentionPolicyPtrInput `pulumi:"retentionPolicy"`
	TimeGrain       pulumi.StringPtrInput   `pulumi:"timeGrain"`
}

func (MetricSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettings)(nil)).Elem()
}

func (i MetricSettingsArgs) ToMetricSettingsOutput() MetricSettingsOutput {
	return i.ToMetricSettingsOutputWithContext(context.Background())
}

func (i MetricSettingsArgs) ToMetricSettingsOutputWithContext(ctx context.Context) MetricSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricSettingsOutput)
}





type MetricSettingsArrayInput interface {
	pulumi.Input

	ToMetricSettingsArrayOutput() MetricSettingsArrayOutput
	ToMetricSettingsArrayOutputWithContext(context.Context) MetricSettingsArrayOutput
}

type MetricSettingsArray []MetricSettingsInput

func (MetricSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettings)(nil)).Elem()
}

func (i MetricSettingsArray) ToMetricSettingsArrayOutput() MetricSettingsArrayOutput {
	return i.ToMetricSettingsArrayOutputWithContext(context.Background())
}

func (i MetricSettingsArray) ToMetricSettingsArrayOutputWithContext(ctx context.Context) MetricSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MetricSettingsArrayOutput)
}

type MetricSettingsOutput struct{ *pulumi.OutputState }

func (MetricSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettings)(nil)).Elem()
}

func (o MetricSettingsOutput) ToMetricSettingsOutput() MetricSettingsOutput {
	return o
}

func (o MetricSettingsOutput) ToMetricSettingsOutputWithContext(ctx context.Context) MetricSettingsOutput {
	return o
}

func (o MetricSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o MetricSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v MetricSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MetricSettingsOutput) RetentionPolicy() RetentionPolicyPtrOutput {
	return o.ApplyT(func(v MetricSettings) *RetentionPolicy { return v.RetentionPolicy }).(RetentionPolicyPtrOutput)
}

func (o MetricSettingsOutput) TimeGrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettings) *string { return v.TimeGrain }).(pulumi.StringPtrOutput)
}

type MetricSettingsArrayOutput struct{ *pulumi.OutputState }

func (MetricSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettings)(nil)).Elem()
}

func (o MetricSettingsArrayOutput) ToMetricSettingsArrayOutput() MetricSettingsArrayOutput {
	return o
}

func (o MetricSettingsArrayOutput) ToMetricSettingsArrayOutputWithContext(ctx context.Context) MetricSettingsArrayOutput {
	return o
}

func (o MetricSettingsArrayOutput) Index(i pulumi.IntInput) MetricSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricSettings {
		return vs[0].([]MetricSettings)[vs[1].(int)]
	}).(MetricSettingsOutput)
}

type MetricSettingsResponse struct {
	Category        *string                  `pulumi:"category"`
	Enabled         bool                     `pulumi:"enabled"`
	RetentionPolicy *RetentionPolicyResponse `pulumi:"retentionPolicy"`
	TimeGrain       *string                  `pulumi:"timeGrain"`
}

type MetricSettingsResponseOutput struct{ *pulumi.OutputState }

func (MetricSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MetricSettingsResponse)(nil)).Elem()
}

func (o MetricSettingsResponseOutput) ToMetricSettingsResponseOutput() MetricSettingsResponseOutput {
	return o
}

func (o MetricSettingsResponseOutput) ToMetricSettingsResponseOutputWithContext(ctx context.Context) MetricSettingsResponseOutput {
	return o
}

func (o MetricSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o MetricSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v MetricSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

func (o MetricSettingsResponseOutput) RetentionPolicy() RetentionPolicyResponsePtrOutput {
	return o.ApplyT(func(v MetricSettingsResponse) *RetentionPolicyResponse { return v.RetentionPolicy }).(RetentionPolicyResponsePtrOutput)
}

func (o MetricSettingsResponseOutput) TimeGrain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MetricSettingsResponse) *string { return v.TimeGrain }).(pulumi.StringPtrOutput)
}

type MetricSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (MetricSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MetricSettingsResponse)(nil)).Elem()
}

func (o MetricSettingsResponseArrayOutput) ToMetricSettingsResponseArrayOutput() MetricSettingsResponseArrayOutput {
	return o
}

func (o MetricSettingsResponseArrayOutput) ToMetricSettingsResponseArrayOutputWithContext(ctx context.Context) MetricSettingsResponseArrayOutput {
	return o
}

func (o MetricSettingsResponseArrayOutput) Index(i pulumi.IntInput) MetricSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MetricSettingsResponse {
		return vs[0].([]MetricSettingsResponse)[vs[1].(int)]
	}).(MetricSettingsResponseOutput)
}

type RetentionPolicy struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}





type RetentionPolicyInput interface {
	pulumi.Input

	ToRetentionPolicyOutput() RetentionPolicyOutput
	ToRetentionPolicyOutputWithContext(context.Context) RetentionPolicyOutput
}

type RetentionPolicyArgs struct {
	Days    pulumi.IntInput  `pulumi:"days"`
	Enabled pulumi.BoolInput `pulumi:"enabled"`
}

func (RetentionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return i.ToRetentionPolicyOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput)
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i RetentionPolicyArgs) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyOutput).ToRetentionPolicyPtrOutputWithContext(ctx)
}









type RetentionPolicyPtrInput interface {
	pulumi.Input

	ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput
	ToRetentionPolicyPtrOutputWithContext(context.Context) RetentionPolicyPtrOutput
}

type retentionPolicyPtrType RetentionPolicyArgs

func RetentionPolicyPtr(v *RetentionPolicyArgs) RetentionPolicyPtrInput {
	return (*retentionPolicyPtrType)(v)
}

func (*retentionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return i.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (i *retentionPolicyPtrType) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RetentionPolicyPtrOutput)
}

type RetentionPolicyOutput struct{ *pulumi.OutputState }

func (RetentionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutput() RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyOutputWithContext(ctx context.Context) RetentionPolicyOutput {
	return o
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o.ToRetentionPolicyPtrOutputWithContext(context.Background())
}

func (o RetentionPolicyOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RetentionPolicy) *RetentionPolicy {
		return &v
	}).(RetentionPolicyPtrOutput)
}

func (o RetentionPolicyOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicy) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicy) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyPtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicy)(nil)).Elem()
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutput() RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) ToRetentionPolicyPtrOutputWithContext(ctx context.Context) RetentionPolicyPtrOutput {
	return o
}

func (o RetentionPolicyPtrOutput) Elem() RetentionPolicyOutput {
	return o.ApplyT(func(v *RetentionPolicy) RetentionPolicy {
		if v != nil {
			return *v
		}
		var ret RetentionPolicy
		return ret
	}).(RetentionPolicyOutput)
}

func (o RetentionPolicyPtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicy) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type RetentionPolicyResponse struct {
	Days    int  `pulumi:"days"`
	Enabled bool `pulumi:"enabled"`
}

type RetentionPolicyResponseOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutput() RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) ToRetentionPolicyResponseOutputWithContext(ctx context.Context) RetentionPolicyResponseOutput {
	return o
}

func (o RetentionPolicyResponseOutput) Days() pulumi.IntOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) int { return v.Days }).(pulumi.IntOutput)
}

func (o RetentionPolicyResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v RetentionPolicyResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type RetentionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (RetentionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RetentionPolicyResponse)(nil)).Elem()
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutput() RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) ToRetentionPolicyResponsePtrOutputWithContext(ctx context.Context) RetentionPolicyResponsePtrOutput {
	return o
}

func (o RetentionPolicyResponsePtrOutput) Elem() RetentionPolicyResponseOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) RetentionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret RetentionPolicyResponse
		return ret
	}).(RetentionPolicyResponseOutput)
}

func (o RetentionPolicyResponsePtrOutput) Days() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return &v.Days
	}).(pulumi.IntPtrOutput)
}

func (o RetentionPolicyResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *RetentionPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type SubscriptionLogSettings struct {
	Category *string `pulumi:"category"`
	Enabled  bool    `pulumi:"enabled"`
}





type SubscriptionLogSettingsInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput
	ToSubscriptionLogSettingsOutputWithContext(context.Context) SubscriptionLogSettingsOutput
}

type SubscriptionLogSettingsArgs struct {
	Category pulumi.StringPtrInput `pulumi:"category"`
	Enabled  pulumi.BoolInput      `pulumi:"enabled"`
}

func (SubscriptionLogSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettings)(nil)).Elem()
}

func (i SubscriptionLogSettingsArgs) ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput {
	return i.ToSubscriptionLogSettingsOutputWithContext(context.Background())
}

func (i SubscriptionLogSettingsArgs) ToSubscriptionLogSettingsOutputWithContext(ctx context.Context) SubscriptionLogSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionLogSettingsOutput)
}





type SubscriptionLogSettingsArrayInput interface {
	pulumi.Input

	ToSubscriptionLogSettingsArrayOutput() SubscriptionLogSettingsArrayOutput
	ToSubscriptionLogSettingsArrayOutputWithContext(context.Context) SubscriptionLogSettingsArrayOutput
}

type SubscriptionLogSettingsArray []SubscriptionLogSettingsInput

func (SubscriptionLogSettingsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettings)(nil)).Elem()
}

func (i SubscriptionLogSettingsArray) ToSubscriptionLogSettingsArrayOutput() SubscriptionLogSettingsArrayOutput {
	return i.ToSubscriptionLogSettingsArrayOutputWithContext(context.Background())
}

func (i SubscriptionLogSettingsArray) ToSubscriptionLogSettingsArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SubscriptionLogSettingsArrayOutput)
}

type SubscriptionLogSettingsOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettings)(nil)).Elem()
}

func (o SubscriptionLogSettingsOutput) ToSubscriptionLogSettingsOutput() SubscriptionLogSettingsOutput {
	return o
}

func (o SubscriptionLogSettingsOutput) ToSubscriptionLogSettingsOutputWithContext(ctx context.Context) SubscriptionLogSettingsOutput {
	return o
}

func (o SubscriptionLogSettingsOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionLogSettings) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o SubscriptionLogSettingsOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SubscriptionLogSettings) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type SubscriptionLogSettingsArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettings)(nil)).Elem()
}

func (o SubscriptionLogSettingsArrayOutput) ToSubscriptionLogSettingsArrayOutput() SubscriptionLogSettingsArrayOutput {
	return o
}

func (o SubscriptionLogSettingsArrayOutput) ToSubscriptionLogSettingsArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsArrayOutput {
	return o
}

func (o SubscriptionLogSettingsArrayOutput) Index(i pulumi.IntInput) SubscriptionLogSettingsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionLogSettings {
		return vs[0].([]SubscriptionLogSettings)[vs[1].(int)]
	}).(SubscriptionLogSettingsOutput)
}

type SubscriptionLogSettingsResponse struct {
	Category *string `pulumi:"category"`
	Enabled  bool    `pulumi:"enabled"`
}

type SubscriptionLogSettingsResponseOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SubscriptionLogSettingsResponse)(nil)).Elem()
}

func (o SubscriptionLogSettingsResponseOutput) ToSubscriptionLogSettingsResponseOutput() SubscriptionLogSettingsResponseOutput {
	return o
}

func (o SubscriptionLogSettingsResponseOutput) ToSubscriptionLogSettingsResponseOutputWithContext(ctx context.Context) SubscriptionLogSettingsResponseOutput {
	return o
}

func (o SubscriptionLogSettingsResponseOutput) Category() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SubscriptionLogSettingsResponse) *string { return v.Category }).(pulumi.StringPtrOutput)
}

func (o SubscriptionLogSettingsResponseOutput) Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SubscriptionLogSettingsResponse) bool { return v.Enabled }).(pulumi.BoolOutput)
}

type SubscriptionLogSettingsResponseArrayOutput struct{ *pulumi.OutputState }

func (SubscriptionLogSettingsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SubscriptionLogSettingsResponse)(nil)).Elem()
}

func (o SubscriptionLogSettingsResponseArrayOutput) ToSubscriptionLogSettingsResponseArrayOutput() SubscriptionLogSettingsResponseArrayOutput {
	return o
}

func (o SubscriptionLogSettingsResponseArrayOutput) ToSubscriptionLogSettingsResponseArrayOutputWithContext(ctx context.Context) SubscriptionLogSettingsResponseArrayOutput {
	return o
}

func (o SubscriptionLogSettingsResponseArrayOutput) Index(i pulumi.IntInput) SubscriptionLogSettingsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SubscriptionLogSettingsResponse {
		return vs[0].([]SubscriptionLogSettingsResponse)[vs[1].(int)]
	}).(SubscriptionLogSettingsResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(LogSettingsOutput{})
	pulumi.RegisterOutputType(LogSettingsArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(MetricSettingsOutput{})
	pulumi.RegisterOutputType(MetricSettingsArrayOutput{})
	pulumi.RegisterOutputType(MetricSettingsResponseOutput{})
	pulumi.RegisterOutputType(MetricSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsArrayOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseOutput{})
	pulumi.RegisterOutputType(SubscriptionLogSettingsResponseArrayOutput{})
}
