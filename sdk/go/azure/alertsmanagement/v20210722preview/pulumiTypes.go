


package v20210722preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type PrometheusRule struct {
	Actions              []PrometheusRuleGroupAction         `pulumi:"actions"`
	Alert                *string                             `pulumi:"alert"`
	Annotations          map[string]string                   `pulumi:"annotations"`
	Expression           *string                             `pulumi:"expression"`
	For                  *string                             `pulumi:"for"`
	Labels               map[string]string                   `pulumi:"labels"`
	Record               *string                             `pulumi:"record"`
	ResolveConfiguration *PrometheusRuleResolveConfiguration `pulumi:"resolveConfiguration"`
	Severity             *int                                `pulumi:"severity"`
}





type PrometheusRuleInput interface {
	pulumi.Input

	ToPrometheusRuleOutput() PrometheusRuleOutput
	ToPrometheusRuleOutputWithContext(context.Context) PrometheusRuleOutput
}

type PrometheusRuleArgs struct {
	Actions              PrometheusRuleGroupActionArrayInput        `pulumi:"actions"`
	Alert                pulumi.StringPtrInput                      `pulumi:"alert"`
	Annotations          pulumi.StringMapInput                      `pulumi:"annotations"`
	Expression           pulumi.StringPtrInput                      `pulumi:"expression"`
	For                  pulumi.StringPtrInput                      `pulumi:"for"`
	Labels               pulumi.StringMapInput                      `pulumi:"labels"`
	Record               pulumi.StringPtrInput                      `pulumi:"record"`
	ResolveConfiguration PrometheusRuleResolveConfigurationPtrInput `pulumi:"resolveConfiguration"`
	Severity             pulumi.IntPtrInput                         `pulumi:"severity"`
}

func (PrometheusRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRule)(nil)).Elem()
}

func (i PrometheusRuleArgs) ToPrometheusRuleOutput() PrometheusRuleOutput {
	return i.ToPrometheusRuleOutputWithContext(context.Background())
}

func (i PrometheusRuleArgs) ToPrometheusRuleOutputWithContext(ctx context.Context) PrometheusRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleOutput)
}





type PrometheusRuleArrayInput interface {
	pulumi.Input

	ToPrometheusRuleArrayOutput() PrometheusRuleArrayOutput
	ToPrometheusRuleArrayOutputWithContext(context.Context) PrometheusRuleArrayOutput
}

type PrometheusRuleArray []PrometheusRuleInput

func (PrometheusRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRule)(nil)).Elem()
}

func (i PrometheusRuleArray) ToPrometheusRuleArrayOutput() PrometheusRuleArrayOutput {
	return i.ToPrometheusRuleArrayOutputWithContext(context.Background())
}

func (i PrometheusRuleArray) ToPrometheusRuleArrayOutputWithContext(ctx context.Context) PrometheusRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleArrayOutput)
}

type PrometheusRuleOutput struct{ *pulumi.OutputState }

func (PrometheusRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRule)(nil)).Elem()
}

func (o PrometheusRuleOutput) ToPrometheusRuleOutput() PrometheusRuleOutput {
	return o
}

func (o PrometheusRuleOutput) ToPrometheusRuleOutputWithContext(ctx context.Context) PrometheusRuleOutput {
	return o
}

func (o PrometheusRuleOutput) Actions() PrometheusRuleGroupActionArrayOutput {
	return o.ApplyT(func(v PrometheusRule) []PrometheusRuleGroupAction { return v.Actions }).(PrometheusRuleGroupActionArrayOutput)
}

func (o PrometheusRuleOutput) Alert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRule) *string { return v.Alert }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrometheusRule) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o PrometheusRuleOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRule) *string { return v.Expression }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleOutput) For() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRule) *string { return v.For }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrometheusRule) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o PrometheusRuleOutput) Record() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRule) *string { return v.Record }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleOutput) ResolveConfiguration() PrometheusRuleResolveConfigurationPtrOutput {
	return o.ApplyT(func(v PrometheusRule) *PrometheusRuleResolveConfiguration { return v.ResolveConfiguration }).(PrometheusRuleResolveConfigurationPtrOutput)
}

func (o PrometheusRuleOutput) Severity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PrometheusRule) *int { return v.Severity }).(pulumi.IntPtrOutput)
}

type PrometheusRuleArrayOutput struct{ *pulumi.OutputState }

func (PrometheusRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRule)(nil)).Elem()
}

func (o PrometheusRuleArrayOutput) ToPrometheusRuleArrayOutput() PrometheusRuleArrayOutput {
	return o
}

func (o PrometheusRuleArrayOutput) ToPrometheusRuleArrayOutputWithContext(ctx context.Context) PrometheusRuleArrayOutput {
	return o
}

func (o PrometheusRuleArrayOutput) Index(i pulumi.IntInput) PrometheusRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrometheusRule {
		return vs[0].([]PrometheusRule)[vs[1].(int)]
	}).(PrometheusRuleOutput)
}

type PrometheusRuleGroupAction struct {
	ActionGroupId    *string           `pulumi:"actionGroupId"`
	ActionProperties map[string]string `pulumi:"actionProperties"`
}





type PrometheusRuleGroupActionInput interface {
	pulumi.Input

	ToPrometheusRuleGroupActionOutput() PrometheusRuleGroupActionOutput
	ToPrometheusRuleGroupActionOutputWithContext(context.Context) PrometheusRuleGroupActionOutput
}

type PrometheusRuleGroupActionArgs struct {
	ActionGroupId    pulumi.StringPtrInput `pulumi:"actionGroupId"`
	ActionProperties pulumi.StringMapInput `pulumi:"actionProperties"`
}

func (PrometheusRuleGroupActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleGroupAction)(nil)).Elem()
}

func (i PrometheusRuleGroupActionArgs) ToPrometheusRuleGroupActionOutput() PrometheusRuleGroupActionOutput {
	return i.ToPrometheusRuleGroupActionOutputWithContext(context.Background())
}

func (i PrometheusRuleGroupActionArgs) ToPrometheusRuleGroupActionOutputWithContext(ctx context.Context) PrometheusRuleGroupActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleGroupActionOutput)
}





type PrometheusRuleGroupActionArrayInput interface {
	pulumi.Input

	ToPrometheusRuleGroupActionArrayOutput() PrometheusRuleGroupActionArrayOutput
	ToPrometheusRuleGroupActionArrayOutputWithContext(context.Context) PrometheusRuleGroupActionArrayOutput
}

type PrometheusRuleGroupActionArray []PrometheusRuleGroupActionInput

func (PrometheusRuleGroupActionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRuleGroupAction)(nil)).Elem()
}

func (i PrometheusRuleGroupActionArray) ToPrometheusRuleGroupActionArrayOutput() PrometheusRuleGroupActionArrayOutput {
	return i.ToPrometheusRuleGroupActionArrayOutputWithContext(context.Background())
}

func (i PrometheusRuleGroupActionArray) ToPrometheusRuleGroupActionArrayOutputWithContext(ctx context.Context) PrometheusRuleGroupActionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleGroupActionArrayOutput)
}

type PrometheusRuleGroupActionOutput struct{ *pulumi.OutputState }

func (PrometheusRuleGroupActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleGroupAction)(nil)).Elem()
}

func (o PrometheusRuleGroupActionOutput) ToPrometheusRuleGroupActionOutput() PrometheusRuleGroupActionOutput {
	return o
}

func (o PrometheusRuleGroupActionOutput) ToPrometheusRuleGroupActionOutputWithContext(ctx context.Context) PrometheusRuleGroupActionOutput {
	return o
}

func (o PrometheusRuleGroupActionOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleGroupAction) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleGroupActionOutput) ActionProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrometheusRuleGroupAction) map[string]string { return v.ActionProperties }).(pulumi.StringMapOutput)
}

type PrometheusRuleGroupActionArrayOutput struct{ *pulumi.OutputState }

func (PrometheusRuleGroupActionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRuleGroupAction)(nil)).Elem()
}

func (o PrometheusRuleGroupActionArrayOutput) ToPrometheusRuleGroupActionArrayOutput() PrometheusRuleGroupActionArrayOutput {
	return o
}

func (o PrometheusRuleGroupActionArrayOutput) ToPrometheusRuleGroupActionArrayOutputWithContext(ctx context.Context) PrometheusRuleGroupActionArrayOutput {
	return o
}

func (o PrometheusRuleGroupActionArrayOutput) Index(i pulumi.IntInput) PrometheusRuleGroupActionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrometheusRuleGroupAction {
		return vs[0].([]PrometheusRuleGroupAction)[vs[1].(int)]
	}).(PrometheusRuleGroupActionOutput)
}

type PrometheusRuleGroupActionResponse struct {
	ActionGroupId    *string           `pulumi:"actionGroupId"`
	ActionProperties map[string]string `pulumi:"actionProperties"`
}





type PrometheusRuleGroupActionResponseInput interface {
	pulumi.Input

	ToPrometheusRuleGroupActionResponseOutput() PrometheusRuleGroupActionResponseOutput
	ToPrometheusRuleGroupActionResponseOutputWithContext(context.Context) PrometheusRuleGroupActionResponseOutput
}

type PrometheusRuleGroupActionResponseArgs struct {
	ActionGroupId    pulumi.StringPtrInput `pulumi:"actionGroupId"`
	ActionProperties pulumi.StringMapInput `pulumi:"actionProperties"`
}

func (PrometheusRuleGroupActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleGroupActionResponse)(nil)).Elem()
}

func (i PrometheusRuleGroupActionResponseArgs) ToPrometheusRuleGroupActionResponseOutput() PrometheusRuleGroupActionResponseOutput {
	return i.ToPrometheusRuleGroupActionResponseOutputWithContext(context.Background())
}

func (i PrometheusRuleGroupActionResponseArgs) ToPrometheusRuleGroupActionResponseOutputWithContext(ctx context.Context) PrometheusRuleGroupActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleGroupActionResponseOutput)
}





type PrometheusRuleGroupActionResponseArrayInput interface {
	pulumi.Input

	ToPrometheusRuleGroupActionResponseArrayOutput() PrometheusRuleGroupActionResponseArrayOutput
	ToPrometheusRuleGroupActionResponseArrayOutputWithContext(context.Context) PrometheusRuleGroupActionResponseArrayOutput
}

type PrometheusRuleGroupActionResponseArray []PrometheusRuleGroupActionResponseInput

func (PrometheusRuleGroupActionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRuleGroupActionResponse)(nil)).Elem()
}

func (i PrometheusRuleGroupActionResponseArray) ToPrometheusRuleGroupActionResponseArrayOutput() PrometheusRuleGroupActionResponseArrayOutput {
	return i.ToPrometheusRuleGroupActionResponseArrayOutputWithContext(context.Background())
}

func (i PrometheusRuleGroupActionResponseArray) ToPrometheusRuleGroupActionResponseArrayOutputWithContext(ctx context.Context) PrometheusRuleGroupActionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleGroupActionResponseArrayOutput)
}

type PrometheusRuleGroupActionResponseOutput struct{ *pulumi.OutputState }

func (PrometheusRuleGroupActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleGroupActionResponse)(nil)).Elem()
}

func (o PrometheusRuleGroupActionResponseOutput) ToPrometheusRuleGroupActionResponseOutput() PrometheusRuleGroupActionResponseOutput {
	return o
}

func (o PrometheusRuleGroupActionResponseOutput) ToPrometheusRuleGroupActionResponseOutputWithContext(ctx context.Context) PrometheusRuleGroupActionResponseOutput {
	return o
}

func (o PrometheusRuleGroupActionResponseOutput) ActionGroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleGroupActionResponse) *string { return v.ActionGroupId }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleGroupActionResponseOutput) ActionProperties() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrometheusRuleGroupActionResponse) map[string]string { return v.ActionProperties }).(pulumi.StringMapOutput)
}

type PrometheusRuleGroupActionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrometheusRuleGroupActionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRuleGroupActionResponse)(nil)).Elem()
}

func (o PrometheusRuleGroupActionResponseArrayOutput) ToPrometheusRuleGroupActionResponseArrayOutput() PrometheusRuleGroupActionResponseArrayOutput {
	return o
}

func (o PrometheusRuleGroupActionResponseArrayOutput) ToPrometheusRuleGroupActionResponseArrayOutputWithContext(ctx context.Context) PrometheusRuleGroupActionResponseArrayOutput {
	return o
}

func (o PrometheusRuleGroupActionResponseArrayOutput) Index(i pulumi.IntInput) PrometheusRuleGroupActionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrometheusRuleGroupActionResponse {
		return vs[0].([]PrometheusRuleGroupActionResponse)[vs[1].(int)]
	}).(PrometheusRuleGroupActionResponseOutput)
}

type PrometheusRuleResolveConfiguration struct {
	AutoResolved  *bool   `pulumi:"autoResolved"`
	TimeToResolve *string `pulumi:"timeToResolve"`
}





type PrometheusRuleResolveConfigurationInput interface {
	pulumi.Input

	ToPrometheusRuleResolveConfigurationOutput() PrometheusRuleResolveConfigurationOutput
	ToPrometheusRuleResolveConfigurationOutputWithContext(context.Context) PrometheusRuleResolveConfigurationOutput
}

type PrometheusRuleResolveConfigurationArgs struct {
	AutoResolved  pulumi.BoolPtrInput   `pulumi:"autoResolved"`
	TimeToResolve pulumi.StringPtrInput `pulumi:"timeToResolve"`
}

func (PrometheusRuleResolveConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleResolveConfiguration)(nil)).Elem()
}

func (i PrometheusRuleResolveConfigurationArgs) ToPrometheusRuleResolveConfigurationOutput() PrometheusRuleResolveConfigurationOutput {
	return i.ToPrometheusRuleResolveConfigurationOutputWithContext(context.Background())
}

func (i PrometheusRuleResolveConfigurationArgs) ToPrometheusRuleResolveConfigurationOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResolveConfigurationOutput)
}

func (i PrometheusRuleResolveConfigurationArgs) ToPrometheusRuleResolveConfigurationPtrOutput() PrometheusRuleResolveConfigurationPtrOutput {
	return i.ToPrometheusRuleResolveConfigurationPtrOutputWithContext(context.Background())
}

func (i PrometheusRuleResolveConfigurationArgs) ToPrometheusRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResolveConfigurationOutput).ToPrometheusRuleResolveConfigurationPtrOutputWithContext(ctx)
}









type PrometheusRuleResolveConfigurationPtrInput interface {
	pulumi.Input

	ToPrometheusRuleResolveConfigurationPtrOutput() PrometheusRuleResolveConfigurationPtrOutput
	ToPrometheusRuleResolveConfigurationPtrOutputWithContext(context.Context) PrometheusRuleResolveConfigurationPtrOutput
}

type prometheusRuleResolveConfigurationPtrType PrometheusRuleResolveConfigurationArgs

func PrometheusRuleResolveConfigurationPtr(v *PrometheusRuleResolveConfigurationArgs) PrometheusRuleResolveConfigurationPtrInput {
	return (*prometheusRuleResolveConfigurationPtrType)(v)
}

func (*prometheusRuleResolveConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleResolveConfiguration)(nil)).Elem()
}

func (i *prometheusRuleResolveConfigurationPtrType) ToPrometheusRuleResolveConfigurationPtrOutput() PrometheusRuleResolveConfigurationPtrOutput {
	return i.ToPrometheusRuleResolveConfigurationPtrOutputWithContext(context.Background())
}

func (i *prometheusRuleResolveConfigurationPtrType) ToPrometheusRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResolveConfigurationPtrOutput)
}

type PrometheusRuleResolveConfigurationOutput struct{ *pulumi.OutputState }

func (PrometheusRuleResolveConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleResolveConfiguration)(nil)).Elem()
}

func (o PrometheusRuleResolveConfigurationOutput) ToPrometheusRuleResolveConfigurationOutput() PrometheusRuleResolveConfigurationOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationOutput) ToPrometheusRuleResolveConfigurationOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationOutput) ToPrometheusRuleResolveConfigurationPtrOutput() PrometheusRuleResolveConfigurationPtrOutput {
	return o.ToPrometheusRuleResolveConfigurationPtrOutputWithContext(context.Background())
}

func (o PrometheusRuleResolveConfigurationOutput) ToPrometheusRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrometheusRuleResolveConfiguration) *PrometheusRuleResolveConfiguration {
		return &v
	}).(PrometheusRuleResolveConfigurationPtrOutput)
}

func (o PrometheusRuleResolveConfigurationOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResolveConfiguration) *bool { return v.AutoResolved }).(pulumi.BoolPtrOutput)
}

func (o PrometheusRuleResolveConfigurationOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResolveConfiguration) *string { return v.TimeToResolve }).(pulumi.StringPtrOutput)
}

type PrometheusRuleResolveConfigurationPtrOutput struct{ *pulumi.OutputState }

func (PrometheusRuleResolveConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleResolveConfiguration)(nil)).Elem()
}

func (o PrometheusRuleResolveConfigurationPtrOutput) ToPrometheusRuleResolveConfigurationPtrOutput() PrometheusRuleResolveConfigurationPtrOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationPtrOutput) ToPrometheusRuleResolveConfigurationPtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationPtrOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationPtrOutput) Elem() PrometheusRuleResolveConfigurationOutput {
	return o.ApplyT(func(v *PrometheusRuleResolveConfiguration) PrometheusRuleResolveConfiguration {
		if v != nil {
			return *v
		}
		var ret PrometheusRuleResolveConfiguration
		return ret
	}).(PrometheusRuleResolveConfigurationOutput)
}

func (o PrometheusRuleResolveConfigurationPtrOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleResolveConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.AutoResolved
	}).(pulumi.BoolPtrOutput)
}

func (o PrometheusRuleResolveConfigurationPtrOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleResolveConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.TimeToResolve
	}).(pulumi.StringPtrOutput)
}

type PrometheusRuleResolveConfigurationResponse struct {
	AutoResolved  *bool   `pulumi:"autoResolved"`
	TimeToResolve *string `pulumi:"timeToResolve"`
}





type PrometheusRuleResolveConfigurationResponseInput interface {
	pulumi.Input

	ToPrometheusRuleResolveConfigurationResponseOutput() PrometheusRuleResolveConfigurationResponseOutput
	ToPrometheusRuleResolveConfigurationResponseOutputWithContext(context.Context) PrometheusRuleResolveConfigurationResponseOutput
}

type PrometheusRuleResolveConfigurationResponseArgs struct {
	AutoResolved  pulumi.BoolPtrInput   `pulumi:"autoResolved"`
	TimeToResolve pulumi.StringPtrInput `pulumi:"timeToResolve"`
}

func (PrometheusRuleResolveConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleResolveConfigurationResponse)(nil)).Elem()
}

func (i PrometheusRuleResolveConfigurationResponseArgs) ToPrometheusRuleResolveConfigurationResponseOutput() PrometheusRuleResolveConfigurationResponseOutput {
	return i.ToPrometheusRuleResolveConfigurationResponseOutputWithContext(context.Background())
}

func (i PrometheusRuleResolveConfigurationResponseArgs) ToPrometheusRuleResolveConfigurationResponseOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResolveConfigurationResponseOutput)
}

func (i PrometheusRuleResolveConfigurationResponseArgs) ToPrometheusRuleResolveConfigurationResponsePtrOutput() PrometheusRuleResolveConfigurationResponsePtrOutput {
	return i.ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i PrometheusRuleResolveConfigurationResponseArgs) ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResolveConfigurationResponseOutput).ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(ctx)
}









type PrometheusRuleResolveConfigurationResponsePtrInput interface {
	pulumi.Input

	ToPrometheusRuleResolveConfigurationResponsePtrOutput() PrometheusRuleResolveConfigurationResponsePtrOutput
	ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(context.Context) PrometheusRuleResolveConfigurationResponsePtrOutput
}

type prometheusRuleResolveConfigurationResponsePtrType PrometheusRuleResolveConfigurationResponseArgs

func PrometheusRuleResolveConfigurationResponsePtr(v *PrometheusRuleResolveConfigurationResponseArgs) PrometheusRuleResolveConfigurationResponsePtrInput {
	return (*prometheusRuleResolveConfigurationResponsePtrType)(v)
}

func (*prometheusRuleResolveConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleResolveConfigurationResponse)(nil)).Elem()
}

func (i *prometheusRuleResolveConfigurationResponsePtrType) ToPrometheusRuleResolveConfigurationResponsePtrOutput() PrometheusRuleResolveConfigurationResponsePtrOutput {
	return i.ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *prometheusRuleResolveConfigurationResponsePtrType) ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResolveConfigurationResponsePtrOutput)
}

type PrometheusRuleResolveConfigurationResponseOutput struct{ *pulumi.OutputState }

func (PrometheusRuleResolveConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleResolveConfigurationResponse)(nil)).Elem()
}

func (o PrometheusRuleResolveConfigurationResponseOutput) ToPrometheusRuleResolveConfigurationResponseOutput() PrometheusRuleResolveConfigurationResponseOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationResponseOutput) ToPrometheusRuleResolveConfigurationResponseOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationResponseOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationResponseOutput) ToPrometheusRuleResolveConfigurationResponsePtrOutput() PrometheusRuleResolveConfigurationResponsePtrOutput {
	return o.ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o PrometheusRuleResolveConfigurationResponseOutput) ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrometheusRuleResolveConfigurationResponse) *PrometheusRuleResolveConfigurationResponse {
		return &v
	}).(PrometheusRuleResolveConfigurationResponsePtrOutput)
}

func (o PrometheusRuleResolveConfigurationResponseOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResolveConfigurationResponse) *bool { return v.AutoResolved }).(pulumi.BoolPtrOutput)
}

func (o PrometheusRuleResolveConfigurationResponseOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResolveConfigurationResponse) *string { return v.TimeToResolve }).(pulumi.StringPtrOutput)
}

type PrometheusRuleResolveConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (PrometheusRuleResolveConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrometheusRuleResolveConfigurationResponse)(nil)).Elem()
}

func (o PrometheusRuleResolveConfigurationResponsePtrOutput) ToPrometheusRuleResolveConfigurationResponsePtrOutput() PrometheusRuleResolveConfigurationResponsePtrOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationResponsePtrOutput) ToPrometheusRuleResolveConfigurationResponsePtrOutputWithContext(ctx context.Context) PrometheusRuleResolveConfigurationResponsePtrOutput {
	return o
}

func (o PrometheusRuleResolveConfigurationResponsePtrOutput) Elem() PrometheusRuleResolveConfigurationResponseOutput {
	return o.ApplyT(func(v *PrometheusRuleResolveConfigurationResponse) PrometheusRuleResolveConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret PrometheusRuleResolveConfigurationResponse
		return ret
	}).(PrometheusRuleResolveConfigurationResponseOutput)
}

func (o PrometheusRuleResolveConfigurationResponsePtrOutput) AutoResolved() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleResolveConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoResolved
	}).(pulumi.BoolPtrOutput)
}

func (o PrometheusRuleResolveConfigurationResponsePtrOutput) TimeToResolve() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrometheusRuleResolveConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeToResolve
	}).(pulumi.StringPtrOutput)
}

type PrometheusRuleResponse struct {
	Actions              []PrometheusRuleGroupActionResponse         `pulumi:"actions"`
	Alert                *string                                     `pulumi:"alert"`
	Annotations          map[string]string                           `pulumi:"annotations"`
	Expression           *string                                     `pulumi:"expression"`
	For                  *string                                     `pulumi:"for"`
	Labels               map[string]string                           `pulumi:"labels"`
	Record               *string                                     `pulumi:"record"`
	ResolveConfiguration *PrometheusRuleResolveConfigurationResponse `pulumi:"resolveConfiguration"`
	Severity             *int                                        `pulumi:"severity"`
}





type PrometheusRuleResponseInput interface {
	pulumi.Input

	ToPrometheusRuleResponseOutput() PrometheusRuleResponseOutput
	ToPrometheusRuleResponseOutputWithContext(context.Context) PrometheusRuleResponseOutput
}

type PrometheusRuleResponseArgs struct {
	Actions              PrometheusRuleGroupActionResponseArrayInput        `pulumi:"actions"`
	Alert                pulumi.StringPtrInput                              `pulumi:"alert"`
	Annotations          pulumi.StringMapInput                              `pulumi:"annotations"`
	Expression           pulumi.StringPtrInput                              `pulumi:"expression"`
	For                  pulumi.StringPtrInput                              `pulumi:"for"`
	Labels               pulumi.StringMapInput                              `pulumi:"labels"`
	Record               pulumi.StringPtrInput                              `pulumi:"record"`
	ResolveConfiguration PrometheusRuleResolveConfigurationResponsePtrInput `pulumi:"resolveConfiguration"`
	Severity             pulumi.IntPtrInput                                 `pulumi:"severity"`
}

func (PrometheusRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleResponse)(nil)).Elem()
}

func (i PrometheusRuleResponseArgs) ToPrometheusRuleResponseOutput() PrometheusRuleResponseOutput {
	return i.ToPrometheusRuleResponseOutputWithContext(context.Background())
}

func (i PrometheusRuleResponseArgs) ToPrometheusRuleResponseOutputWithContext(ctx context.Context) PrometheusRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResponseOutput)
}





type PrometheusRuleResponseArrayInput interface {
	pulumi.Input

	ToPrometheusRuleResponseArrayOutput() PrometheusRuleResponseArrayOutput
	ToPrometheusRuleResponseArrayOutputWithContext(context.Context) PrometheusRuleResponseArrayOutput
}

type PrometheusRuleResponseArray []PrometheusRuleResponseInput

func (PrometheusRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRuleResponse)(nil)).Elem()
}

func (i PrometheusRuleResponseArray) ToPrometheusRuleResponseArrayOutput() PrometheusRuleResponseArrayOutput {
	return i.ToPrometheusRuleResponseArrayOutputWithContext(context.Background())
}

func (i PrometheusRuleResponseArray) ToPrometheusRuleResponseArrayOutputWithContext(ctx context.Context) PrometheusRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrometheusRuleResponseArrayOutput)
}

type PrometheusRuleResponseOutput struct{ *pulumi.OutputState }

func (PrometheusRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrometheusRuleResponse)(nil)).Elem()
}

func (o PrometheusRuleResponseOutput) ToPrometheusRuleResponseOutput() PrometheusRuleResponseOutput {
	return o
}

func (o PrometheusRuleResponseOutput) ToPrometheusRuleResponseOutputWithContext(ctx context.Context) PrometheusRuleResponseOutput {
	return o
}

func (o PrometheusRuleResponseOutput) Actions() PrometheusRuleGroupActionResponseArrayOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) []PrometheusRuleGroupActionResponse { return v.Actions }).(PrometheusRuleGroupActionResponseArrayOutput)
}

func (o PrometheusRuleResponseOutput) Alert() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) *string { return v.Alert }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleResponseOutput) Annotations() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) map[string]string { return v.Annotations }).(pulumi.StringMapOutput)
}

func (o PrometheusRuleResponseOutput) Expression() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) *string { return v.Expression }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleResponseOutput) For() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) *string { return v.For }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleResponseOutput) Labels() pulumi.StringMapOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) map[string]string { return v.Labels }).(pulumi.StringMapOutput)
}

func (o PrometheusRuleResponseOutput) Record() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) *string { return v.Record }).(pulumi.StringPtrOutput)
}

func (o PrometheusRuleResponseOutput) ResolveConfiguration() PrometheusRuleResolveConfigurationResponsePtrOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) *PrometheusRuleResolveConfigurationResponse {
		return v.ResolveConfiguration
	}).(PrometheusRuleResolveConfigurationResponsePtrOutput)
}

func (o PrometheusRuleResponseOutput) Severity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PrometheusRuleResponse) *int { return v.Severity }).(pulumi.IntPtrOutput)
}

type PrometheusRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (PrometheusRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrometheusRuleResponse)(nil)).Elem()
}

func (o PrometheusRuleResponseArrayOutput) ToPrometheusRuleResponseArrayOutput() PrometheusRuleResponseArrayOutput {
	return o
}

func (o PrometheusRuleResponseArrayOutput) ToPrometheusRuleResponseArrayOutputWithContext(ctx context.Context) PrometheusRuleResponseArrayOutput {
	return o
}

func (o PrometheusRuleResponseArrayOutput) Index(i pulumi.IntInput) PrometheusRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrometheusRuleResponse {
		return vs[0].([]PrometheusRuleResponse)[vs[1].(int)]
	}).(PrometheusRuleResponseOutput)
}

type SystemDataResponse struct {
	CreatedAt          *string `pulumi:"createdAt"`
	CreatedBy          *string `pulumi:"createdBy"`
	CreatedByType      *string `pulumi:"createdByType"`
	LastModifiedAt     *string `pulumi:"lastModifiedAt"`
	LastModifiedBy     *string `pulumi:"lastModifiedBy"`
	LastModifiedByType *string `pulumi:"lastModifiedByType"`
}





type SystemDataResponseInput interface {
	pulumi.Input

	ToSystemDataResponseOutput() SystemDataResponseOutput
	ToSystemDataResponseOutputWithContext(context.Context) SystemDataResponseOutput
}

type SystemDataResponseArgs struct {
	CreatedAt          pulumi.StringPtrInput `pulumi:"createdAt"`
	CreatedBy          pulumi.StringPtrInput `pulumi:"createdBy"`
	CreatedByType      pulumi.StringPtrInput `pulumi:"createdByType"`
	LastModifiedAt     pulumi.StringPtrInput `pulumi:"lastModifiedAt"`
	LastModifiedBy     pulumi.StringPtrInput `pulumi:"lastModifiedBy"`
	LastModifiedByType pulumi.StringPtrInput `pulumi:"lastModifiedByType"`
}

func (SystemDataResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return i.ToSystemDataResponseOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput)
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i SystemDataResponseArgs) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponseOutput).ToSystemDataResponsePtrOutputWithContext(ctx)
}









type SystemDataResponsePtrInput interface {
	pulumi.Input

	ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput
	ToSystemDataResponsePtrOutputWithContext(context.Context) SystemDataResponsePtrOutput
}

type systemDataResponsePtrType SystemDataResponseArgs

func SystemDataResponsePtr(v *SystemDataResponseArgs) SystemDataResponsePtrInput {
	return (*systemDataResponsePtrType)(v)
}

func (*systemDataResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return i.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (i *systemDataResponsePtrType) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SystemDataResponsePtrOutput)
}

type SystemDataResponseOutput struct{ *pulumi.OutputState }

func (SystemDataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutput() SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponseOutputWithContext(ctx context.Context) SystemDataResponseOutput {
	return o
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o.ToSystemDataResponsePtrOutputWithContext(context.Background())
}

func (o SystemDataResponseOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SystemDataResponse) *SystemDataResponse {
		return &v
	}).(SystemDataResponsePtrOutput)
}

func (o SystemDataResponseOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.CreatedByType }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedAt }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedBy }).(pulumi.StringPtrOutput)
}

func (o SystemDataResponseOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SystemDataResponse) *string { return v.LastModifiedByType }).(pulumi.StringPtrOutput)
}

type SystemDataResponsePtrOutput struct{ *pulumi.OutputState }

func (SystemDataResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SystemDataResponse)(nil)).Elem()
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutput() SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) ToSystemDataResponsePtrOutputWithContext(ctx context.Context) SystemDataResponsePtrOutput {
	return o
}

func (o SystemDataResponsePtrOutput) Elem() SystemDataResponseOutput {
	return o.ApplyT(func(v *SystemDataResponse) SystemDataResponse {
		if v != nil {
			return *v
		}
		var ret SystemDataResponse
		return ret
	}).(SystemDataResponseOutput)
}

func (o SystemDataResponsePtrOutput) CreatedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) CreatedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.CreatedByType
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedAt() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedAt
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedBy
	}).(pulumi.StringPtrOutput)
}

func (o SystemDataResponsePtrOutput) LastModifiedByType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SystemDataResponse) *string {
		if v == nil {
			return nil
		}
		return v.LastModifiedByType
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(PrometheusRuleOutput{})
	pulumi.RegisterOutputType(PrometheusRuleArrayOutput{})
	pulumi.RegisterOutputType(PrometheusRuleGroupActionOutput{})
	pulumi.RegisterOutputType(PrometheusRuleGroupActionArrayOutput{})
	pulumi.RegisterOutputType(PrometheusRuleGroupActionResponseOutput{})
	pulumi.RegisterOutputType(PrometheusRuleGroupActionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrometheusRuleResolveConfigurationOutput{})
	pulumi.RegisterOutputType(PrometheusRuleResolveConfigurationPtrOutput{})
	pulumi.RegisterOutputType(PrometheusRuleResolveConfigurationResponseOutput{})
	pulumi.RegisterOutputType(PrometheusRuleResolveConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(PrometheusRuleResponseOutput{})
	pulumi.RegisterOutputType(PrometheusRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
}
