


package v20170401preview

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

func init() {
	pulumi.RegisterOutputType(LogSettingsOutput{})
	pulumi.RegisterOutputType(LogSettingsArrayOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseOutput{})
	pulumi.RegisterOutputType(LogSettingsResponseArrayOutput{})
	pulumi.RegisterOutputType(RetentionPolicyOutput{})
	pulumi.RegisterOutputType(RetentionPolicyPtrOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponseOutput{})
	pulumi.RegisterOutputType(RetentionPolicyResponsePtrOutput{})
}
