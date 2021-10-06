


package v20181120

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationParameter struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ConfigurationParameterInput interface {
	pulumi.Input

	ToConfigurationParameterOutput() ConfigurationParameterOutput
	ToConfigurationParameterOutputWithContext(context.Context) ConfigurationParameterOutput
}

type ConfigurationParameterArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConfigurationParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameter)(nil)).Elem()
}

func (i ConfigurationParameterArgs) ToConfigurationParameterOutput() ConfigurationParameterOutput {
	return i.ToConfigurationParameterOutputWithContext(context.Background())
}

func (i ConfigurationParameterArgs) ToConfigurationParameterOutputWithContext(ctx context.Context) ConfigurationParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterOutput)
}





type ConfigurationParameterArrayInput interface {
	pulumi.Input

	ToConfigurationParameterArrayOutput() ConfigurationParameterArrayOutput
	ToConfigurationParameterArrayOutputWithContext(context.Context) ConfigurationParameterArrayOutput
}

type ConfigurationParameterArray []ConfigurationParameterInput

func (ConfigurationParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameter)(nil)).Elem()
}

func (i ConfigurationParameterArray) ToConfigurationParameterArrayOutput() ConfigurationParameterArrayOutput {
	return i.ToConfigurationParameterArrayOutputWithContext(context.Background())
}

func (i ConfigurationParameterArray) ToConfigurationParameterArrayOutputWithContext(ctx context.Context) ConfigurationParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterArrayOutput)
}

type ConfigurationParameterOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameter)(nil)).Elem()
}

func (o ConfigurationParameterOutput) ToConfigurationParameterOutput() ConfigurationParameterOutput {
	return o
}

func (o ConfigurationParameterOutput) ToConfigurationParameterOutputWithContext(ctx context.Context) ConfigurationParameterOutput {
	return o
}

func (o ConfigurationParameterOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameter) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConfigurationParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConfigurationParameterArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameter)(nil)).Elem()
}

func (o ConfigurationParameterArrayOutput) ToConfigurationParameterArrayOutput() ConfigurationParameterArrayOutput {
	return o
}

func (o ConfigurationParameterArrayOutput) ToConfigurationParameterArrayOutputWithContext(ctx context.Context) ConfigurationParameterArrayOutput {
	return o
}

func (o ConfigurationParameterArrayOutput) Index(i pulumi.IntInput) ConfigurationParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationParameter {
		return vs[0].([]ConfigurationParameter)[vs[1].(int)]
	}).(ConfigurationParameterOutput)
}

type ConfigurationParameterResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type ConfigurationParameterResponseInput interface {
	pulumi.Input

	ToConfigurationParameterResponseOutput() ConfigurationParameterResponseOutput
	ToConfigurationParameterResponseOutputWithContext(context.Context) ConfigurationParameterResponseOutput
}

type ConfigurationParameterResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConfigurationParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameterResponse)(nil)).Elem()
}

func (i ConfigurationParameterResponseArgs) ToConfigurationParameterResponseOutput() ConfigurationParameterResponseOutput {
	return i.ToConfigurationParameterResponseOutputWithContext(context.Background())
}

func (i ConfigurationParameterResponseArgs) ToConfigurationParameterResponseOutputWithContext(ctx context.Context) ConfigurationParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterResponseOutput)
}





type ConfigurationParameterResponseArrayInput interface {
	pulumi.Input

	ToConfigurationParameterResponseArrayOutput() ConfigurationParameterResponseArrayOutput
	ToConfigurationParameterResponseArrayOutputWithContext(context.Context) ConfigurationParameterResponseArrayOutput
}

type ConfigurationParameterResponseArray []ConfigurationParameterResponseInput

func (ConfigurationParameterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameterResponse)(nil)).Elem()
}

func (i ConfigurationParameterResponseArray) ToConfigurationParameterResponseArrayOutput() ConfigurationParameterResponseArrayOutput {
	return i.ToConfigurationParameterResponseArrayOutputWithContext(context.Background())
}

func (i ConfigurationParameterResponseArray) ToConfigurationParameterResponseArrayOutputWithContext(ctx context.Context) ConfigurationParameterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationParameterResponseArrayOutput)
}

type ConfigurationParameterResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationParameterResponse)(nil)).Elem()
}

func (o ConfigurationParameterResponseOutput) ToConfigurationParameterResponseOutput() ConfigurationParameterResponseOutput {
	return o
}

func (o ConfigurationParameterResponseOutput) ToConfigurationParameterResponseOutputWithContext(ctx context.Context) ConfigurationParameterResponseOutput {
	return o
}

func (o ConfigurationParameterResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameterResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConfigurationParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConfigurationParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ConfigurationParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConfigurationParameterResponse)(nil)).Elem()
}

func (o ConfigurationParameterResponseArrayOutput) ToConfigurationParameterResponseArrayOutput() ConfigurationParameterResponseArrayOutput {
	return o
}

func (o ConfigurationParameterResponseArrayOutput) ToConfigurationParameterResponseArrayOutputWithContext(ctx context.Context) ConfigurationParameterResponseArrayOutput {
	return o
}

func (o ConfigurationParameterResponseArrayOutput) Index(i pulumi.IntInput) ConfigurationParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConfigurationParameterResponse {
		return vs[0].([]ConfigurationParameterResponse)[vs[1].(int)]
	}).(ConfigurationParameterResponseOutput)
}

type ConfigurationSetting struct {
	ActionAfterReboot              *string  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           *bool    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              *string  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins *float64 `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 *bool    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           *float64 `pulumi:"refreshFrequencyMins"`
}





type ConfigurationSettingInput interface {
	pulumi.Input

	ToConfigurationSettingOutput() ConfigurationSettingOutput
	ToConfigurationSettingOutputWithContext(context.Context) ConfigurationSettingOutput
}

type ConfigurationSettingArgs struct {
	ActionAfterReboot              pulumi.StringPtrInput  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           pulumi.BoolPtrInput    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              pulumi.StringPtrInput  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins pulumi.Float64PtrInput `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 pulumi.BoolPtrInput    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           pulumi.Float64PtrInput `pulumi:"refreshFrequencyMins"`
}

func (ConfigurationSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSetting)(nil)).Elem()
}

func (i ConfigurationSettingArgs) ToConfigurationSettingOutput() ConfigurationSettingOutput {
	return i.ToConfigurationSettingOutputWithContext(context.Background())
}

func (i ConfigurationSettingArgs) ToConfigurationSettingOutputWithContext(ctx context.Context) ConfigurationSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingOutput)
}

func (i ConfigurationSettingArgs) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return i.ToConfigurationSettingPtrOutputWithContext(context.Background())
}

func (i ConfigurationSettingArgs) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingOutput).ToConfigurationSettingPtrOutputWithContext(ctx)
}









type ConfigurationSettingPtrInput interface {
	pulumi.Input

	ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput
	ToConfigurationSettingPtrOutputWithContext(context.Context) ConfigurationSettingPtrOutput
}

type configurationSettingPtrType ConfigurationSettingArgs

func ConfigurationSettingPtr(v *ConfigurationSettingArgs) ConfigurationSettingPtrInput {
	return (*configurationSettingPtrType)(v)
}

func (*configurationSettingPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSetting)(nil)).Elem()
}

func (i *configurationSettingPtrType) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return i.ToConfigurationSettingPtrOutputWithContext(context.Background())
}

func (i *configurationSettingPtrType) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingPtrOutput)
}

type ConfigurationSettingOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSetting)(nil)).Elem()
}

func (o ConfigurationSettingOutput) ToConfigurationSettingOutput() ConfigurationSettingOutput {
	return o
}

func (o ConfigurationSettingOutput) ToConfigurationSettingOutputWithContext(ctx context.Context) ConfigurationSettingOutput {
	return o
}

func (o ConfigurationSettingOutput) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return o.ToConfigurationSettingPtrOutputWithContext(context.Background())
}

func (o ConfigurationSettingOutput) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationSetting) *ConfigurationSetting {
		return &v
	}).(ConfigurationSettingPtrOutput)
}

func (o ConfigurationSettingOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *string { return v.ActionAfterReboot }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *bool { return v.AllowModuleOverwrite }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *string { return v.ConfigurationMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *float64 { return v.ConfigurationModeFrequencyMins }).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *bool { return v.RebootIfNeeded }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSetting) *float64 { return v.RefreshFrequencyMins }).(pulumi.Float64PtrOutput)
}

type ConfigurationSettingPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSetting)(nil)).Elem()
}

func (o ConfigurationSettingPtrOutput) ToConfigurationSettingPtrOutput() ConfigurationSettingPtrOutput {
	return o
}

func (o ConfigurationSettingPtrOutput) ToConfigurationSettingPtrOutputWithContext(ctx context.Context) ConfigurationSettingPtrOutput {
	return o
}

func (o ConfigurationSettingPtrOutput) Elem() ConfigurationSettingOutput {
	return o.ApplyT(func(v *ConfigurationSetting) ConfigurationSetting {
		if v != nil {
			return *v
		}
		var ret ConfigurationSetting
		return ret
	}).(ConfigurationSettingOutput)
}

func (o ConfigurationSettingPtrOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *string {
		if v == nil {
			return nil
		}
		return v.ActionAfterReboot
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingPtrOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *bool {
		if v == nil {
			return nil
		}
		return v.AllowModuleOverwrite
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingPtrOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingPtrOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *float64 {
		if v == nil {
			return nil
		}
		return v.ConfigurationModeFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingPtrOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *bool {
		if v == nil {
			return nil
		}
		return v.RebootIfNeeded
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingPtrOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSetting) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

type ConfigurationSettingResponse struct {
	ActionAfterReboot              *string  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           *bool    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              *string  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins *float64 `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 *bool    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           *float64 `pulumi:"refreshFrequencyMins"`
}





type ConfigurationSettingResponseInput interface {
	pulumi.Input

	ToConfigurationSettingResponseOutput() ConfigurationSettingResponseOutput
	ToConfigurationSettingResponseOutputWithContext(context.Context) ConfigurationSettingResponseOutput
}

type ConfigurationSettingResponseArgs struct {
	ActionAfterReboot              pulumi.StringPtrInput  `pulumi:"actionAfterReboot"`
	AllowModuleOverwrite           pulumi.BoolPtrInput    `pulumi:"allowModuleOverwrite"`
	ConfigurationMode              pulumi.StringPtrInput  `pulumi:"configurationMode"`
	ConfigurationModeFrequencyMins pulumi.Float64PtrInput `pulumi:"configurationModeFrequencyMins"`
	RebootIfNeeded                 pulumi.BoolPtrInput    `pulumi:"rebootIfNeeded"`
	RefreshFrequencyMins           pulumi.Float64PtrInput `pulumi:"refreshFrequencyMins"`
}

func (ConfigurationSettingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSettingResponse)(nil)).Elem()
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponseOutput() ConfigurationSettingResponseOutput {
	return i.ToConfigurationSettingResponseOutputWithContext(context.Background())
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponseOutputWithContext(ctx context.Context) ConfigurationSettingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingResponseOutput)
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return i.ToConfigurationSettingResponsePtrOutputWithContext(context.Background())
}

func (i ConfigurationSettingResponseArgs) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingResponseOutput).ToConfigurationSettingResponsePtrOutputWithContext(ctx)
}









type ConfigurationSettingResponsePtrInput interface {
	pulumi.Input

	ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput
	ToConfigurationSettingResponsePtrOutputWithContext(context.Context) ConfigurationSettingResponsePtrOutput
}

type configurationSettingResponsePtrType ConfigurationSettingResponseArgs

func ConfigurationSettingResponsePtr(v *ConfigurationSettingResponseArgs) ConfigurationSettingResponsePtrInput {
	return (*configurationSettingResponsePtrType)(v)
}

func (*configurationSettingResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSettingResponse)(nil)).Elem()
}

func (i *configurationSettingResponsePtrType) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return i.ToConfigurationSettingResponsePtrOutputWithContext(context.Background())
}

func (i *configurationSettingResponsePtrType) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationSettingResponsePtrOutput)
}

type ConfigurationSettingResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationSettingResponse)(nil)).Elem()
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponseOutput() ConfigurationSettingResponseOutput {
	return o
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponseOutputWithContext(ctx context.Context) ConfigurationSettingResponseOutput {
	return o
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return o.ToConfigurationSettingResponsePtrOutputWithContext(context.Background())
}

func (o ConfigurationSettingResponseOutput) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConfigurationSettingResponse) *ConfigurationSettingResponse {
		return &v
	}).(ConfigurationSettingResponsePtrOutput)
}

func (o ConfigurationSettingResponseOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *string { return v.ActionAfterReboot }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponseOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *bool { return v.AllowModuleOverwrite }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponseOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *string { return v.ConfigurationMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponseOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *float64 { return v.ConfigurationModeFrequencyMins }).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingResponseOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *bool { return v.RebootIfNeeded }).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponseOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConfigurationSettingResponse) *float64 { return v.RefreshFrequencyMins }).(pulumi.Float64PtrOutput)
}

type ConfigurationSettingResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationSettingResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationSettingResponse)(nil)).Elem()
}

func (o ConfigurationSettingResponsePtrOutput) ToConfigurationSettingResponsePtrOutput() ConfigurationSettingResponsePtrOutput {
	return o
}

func (o ConfigurationSettingResponsePtrOutput) ToConfigurationSettingResponsePtrOutputWithContext(ctx context.Context) ConfigurationSettingResponsePtrOutput {
	return o
}

func (o ConfigurationSettingResponsePtrOutput) Elem() ConfigurationSettingResponseOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) ConfigurationSettingResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationSettingResponse
		return ret
	}).(ConfigurationSettingResponseOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ActionAfterReboot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionAfterReboot
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) AllowModuleOverwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowModuleOverwrite
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ConfigurationMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConfigurationMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) ConfigurationModeFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.ConfigurationModeFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) RebootIfNeeded() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RebootIfNeeded
	}).(pulumi.BoolPtrOutput)
}

func (o ConfigurationSettingResponsePtrOutput) RefreshFrequencyMins() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConfigurationSettingResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.RefreshFrequencyMins
	}).(pulumi.Float64PtrOutput)
}

type GuestConfigurationAssignmentProperties struct {
	Context            *string                       `pulumi:"context"`
	GuestConfiguration *GuestConfigurationNavigation `pulumi:"guestConfiguration"`
}





type GuestConfigurationAssignmentPropertiesInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesOutput() GuestConfigurationAssignmentPropertiesOutput
	ToGuestConfigurationAssignmentPropertiesOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesOutput
}

type GuestConfigurationAssignmentPropertiesArgs struct {
	Context            pulumi.StringPtrInput                `pulumi:"context"`
	GuestConfiguration GuestConfigurationNavigationPtrInput `pulumi:"guestConfiguration"`
}

func (GuestConfigurationAssignmentPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesOutput() GuestConfigurationAssignmentPropertiesOutput {
	return i.ToGuestConfigurationAssignmentPropertiesOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesOutput)
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesArgs) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesOutput).ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx)
}









type GuestConfigurationAssignmentPropertiesPtrInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput
	ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesPtrOutput
}

type guestConfigurationAssignmentPropertiesPtrType GuestConfigurationAssignmentPropertiesArgs

func GuestConfigurationAssignmentPropertiesPtr(v *GuestConfigurationAssignmentPropertiesArgs) GuestConfigurationAssignmentPropertiesPtrInput {
	return (*guestConfigurationAssignmentPropertiesPtrType)(v)
}

func (*guestConfigurationAssignmentPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (i *guestConfigurationAssignmentPropertiesPtrType) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (i *guestConfigurationAssignmentPropertiesPtrType) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesPtrOutput)
}

type GuestConfigurationAssignmentPropertiesOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesOutput() GuestConfigurationAssignmentPropertiesOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return o.ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(context.Background())
}

func (o GuestConfigurationAssignmentPropertiesOutput) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationAssignmentProperties) *GuestConfigurationAssignmentProperties {
		return &v
	}).(GuestConfigurationAssignmentPropertiesPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentProperties) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesOutput) GuestConfiguration() GuestConfigurationNavigationPtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentProperties) *GuestConfigurationNavigation {
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationPtrOutput)
}

type GuestConfigurationAssignmentPropertiesPtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentProperties)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) ToGuestConfigurationAssignmentPropertiesPtrOutput() GuestConfigurationAssignmentPropertiesPtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) ToGuestConfigurationAssignmentPropertiesPtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesPtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) Elem() GuestConfigurationAssignmentPropertiesOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentProperties) GuestConfigurationAssignmentProperties {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationAssignmentProperties
		return ret
	}).(GuestConfigurationAssignmentPropertiesOutput)
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentProperties) *string {
		if v == nil {
			return nil
		}
		return v.Context
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesPtrOutput) GuestConfiguration() GuestConfigurationNavigationPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentProperties) *GuestConfigurationNavigation {
		if v == nil {
			return nil
		}
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationPtrOutput)
}

type GuestConfigurationAssignmentPropertiesResponse struct {
	AssignmentHash              string                                `pulumi:"assignmentHash"`
	ComplianceStatus            string                                `pulumi:"complianceStatus"`
	Context                     *string                               `pulumi:"context"`
	GuestConfiguration          *GuestConfigurationNavigationResponse `pulumi:"guestConfiguration"`
	LastComplianceStatusChecked string                                `pulumi:"lastComplianceStatusChecked"`
	LatestReportId              string                                `pulumi:"latestReportId"`
	ParameterHash               string                                `pulumi:"parameterHash"`
	ProvisioningState           string                                `pulumi:"provisioningState"`
	VmssVMList                  []VMSSVMInfoResponse                  `pulumi:"vmssVMList"`
}





type GuestConfigurationAssignmentPropertiesResponseInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesResponseOutput() GuestConfigurationAssignmentPropertiesResponseOutput
	ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesResponseOutput
}

type GuestConfigurationAssignmentPropertiesResponseArgs struct {
	AssignmentHash              pulumi.StringInput                           `pulumi:"assignmentHash"`
	ComplianceStatus            pulumi.StringInput                           `pulumi:"complianceStatus"`
	Context                     pulumi.StringPtrInput                        `pulumi:"context"`
	GuestConfiguration          GuestConfigurationNavigationResponsePtrInput `pulumi:"guestConfiguration"`
	LastComplianceStatusChecked pulumi.StringInput                           `pulumi:"lastComplianceStatusChecked"`
	LatestReportId              pulumi.StringInput                           `pulumi:"latestReportId"`
	ParameterHash               pulumi.StringInput                           `pulumi:"parameterHash"`
	ProvisioningState           pulumi.StringInput                           `pulumi:"provisioningState"`
	VmssVMList                  VMSSVMInfoResponseArrayInput                 `pulumi:"vmssVMList"`
}

func (GuestConfigurationAssignmentPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponseOutput() GuestConfigurationAssignmentPropertiesResponseOutput {
	return i.ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i GuestConfigurationAssignmentPropertiesResponseArgs) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesResponseOutput).ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx)
}









type GuestConfigurationAssignmentPropertiesResponsePtrInput interface {
	pulumi.Input

	ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput
	ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput
}

type guestConfigurationAssignmentPropertiesResponsePtrType GuestConfigurationAssignmentPropertiesResponseArgs

func GuestConfigurationAssignmentPropertiesResponsePtr(v *GuestConfigurationAssignmentPropertiesResponseArgs) GuestConfigurationAssignmentPropertiesResponsePtrInput {
	return (*guestConfigurationAssignmentPropertiesResponsePtrType)(v)
}

func (*guestConfigurationAssignmentPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (i *guestConfigurationAssignmentPropertiesResponsePtrType) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return i.ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *guestConfigurationAssignmentPropertiesResponsePtrType) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationAssignmentPropertiesResponsePtrOutput)
}

type GuestConfigurationAssignmentPropertiesResponseOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponseOutput() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponseOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponseOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o.ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationAssignmentPropertiesResponse) *GuestConfigurationAssignmentPropertiesResponse {
		return &v
	}).(GuestConfigurationAssignmentPropertiesResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) AssignmentHash() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.AssignmentHash }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ComplianceStatus }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) *string { return v.Context }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) GuestConfiguration() GuestConfigurationNavigationResponsePtrOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) *GuestConfigurationNavigationResponse {
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) LastComplianceStatusChecked() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.LastComplianceStatusChecked }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) LatestReportId() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.LatestReportId }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ParameterHash() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ParameterHash }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponseOutput) VmssVMList() VMSSVMInfoResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) []VMSSVMInfoResponse { return v.VmssVMList }).(VMSSVMInfoResponseArrayOutput)
}

type GuestConfigurationAssignmentPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationAssignmentPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationAssignmentPropertiesResponse)(nil)).Elem()
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutput() GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ToGuestConfigurationAssignmentPropertiesResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationAssignmentPropertiesResponsePtrOutput {
	return o
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) Elem() GuestConfigurationAssignmentPropertiesResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) GuestConfigurationAssignmentPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationAssignmentPropertiesResponse
		return ret
	}).(GuestConfigurationAssignmentPropertiesResponseOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) AssignmentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AssignmentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ComplianceStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ComplianceStatus
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) Context() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Context
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) GuestConfiguration() GuestConfigurationNavigationResponsePtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *GuestConfigurationNavigationResponse {
		if v == nil {
			return nil
		}
		return v.GuestConfiguration
	}).(GuestConfigurationNavigationResponsePtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) LastComplianceStatusChecked() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastComplianceStatusChecked
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) LatestReportId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LatestReportId
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ParameterHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ParameterHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationAssignmentPropertiesResponsePtrOutput) VmssVMList() VMSSVMInfoResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationAssignmentPropertiesResponse) []VMSSVMInfoResponse {
		if v == nil {
			return nil
		}
		return v.VmssVMList
	}).(VMSSVMInfoResponseArrayOutput)
}

type GuestConfigurationNavigation struct {
	AssignmentType                  *string                  `pulumi:"assignmentType"`
	ConfigurationParameter          []ConfigurationParameter `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter []ConfigurationParameter `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            *ConfigurationSetting    `pulumi:"configurationSetting"`
	ContentHash                     *string                  `pulumi:"contentHash"`
	ContentUri                      *string                  `pulumi:"contentUri"`
	Kind                            *string                  `pulumi:"kind"`
	Name                            *string                  `pulumi:"name"`
	Version                         *string                  `pulumi:"version"`
}





type GuestConfigurationNavigationInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput
	ToGuestConfigurationNavigationOutputWithContext(context.Context) GuestConfigurationNavigationOutput
}

type GuestConfigurationNavigationArgs struct {
	AssignmentType                  pulumi.StringPtrInput            `pulumi:"assignmentType"`
	ConfigurationParameter          ConfigurationParameterArrayInput `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter ConfigurationParameterArrayInput `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            ConfigurationSettingPtrInput     `pulumi:"configurationSetting"`
	ContentHash                     pulumi.StringPtrInput            `pulumi:"contentHash"`
	ContentUri                      pulumi.StringPtrInput            `pulumi:"contentUri"`
	Kind                            pulumi.StringPtrInput            `pulumi:"kind"`
	Name                            pulumi.StringPtrInput            `pulumi:"name"`
	Version                         pulumi.StringPtrInput            `pulumi:"version"`
}

func (GuestConfigurationNavigationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigation)(nil)).Elem()
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput {
	return i.ToGuestConfigurationNavigationOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationOutputWithContext(ctx context.Context) GuestConfigurationNavigationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationOutput)
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return i.ToGuestConfigurationNavigationPtrOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationArgs) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationOutput).ToGuestConfigurationNavigationPtrOutputWithContext(ctx)
}









type GuestConfigurationNavigationPtrInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput
	ToGuestConfigurationNavigationPtrOutputWithContext(context.Context) GuestConfigurationNavigationPtrOutput
}

type guestConfigurationNavigationPtrType GuestConfigurationNavigationArgs

func GuestConfigurationNavigationPtr(v *GuestConfigurationNavigationArgs) GuestConfigurationNavigationPtrInput {
	return (*guestConfigurationNavigationPtrType)(v)
}

func (*guestConfigurationNavigationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigation)(nil)).Elem()
}

func (i *guestConfigurationNavigationPtrType) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return i.ToGuestConfigurationNavigationPtrOutputWithContext(context.Background())
}

func (i *guestConfigurationNavigationPtrType) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationPtrOutput)
}

type GuestConfigurationNavigationOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigation)(nil)).Elem()
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput {
	return o
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationOutputWithContext(ctx context.Context) GuestConfigurationNavigationOutput {
	return o
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return o.ToGuestConfigurationNavigationPtrOutputWithContext(context.Background())
}

func (o GuestConfigurationNavigationOutput) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationNavigation) *GuestConfigurationNavigation {
		return &v
	}).(GuestConfigurationNavigationPtrOutput)
}

func (o GuestConfigurationNavigationOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.AssignmentType }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) ConfigurationParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) []ConfigurationParameter { return v.ConfigurationParameter }).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationOutput) ConfigurationProtectedParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) []ConfigurationParameter {
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationOutput) ConfigurationSetting() ConfigurationSettingPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *ConfigurationSetting { return v.ConfigurationSetting }).(ConfigurationSettingPtrOutput)
}

func (o GuestConfigurationNavigationOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.ContentHash }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.ContentUri }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GuestConfigurationNavigationPtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigation)(nil)).Elem()
}

func (o GuestConfigurationNavigationPtrOutput) ToGuestConfigurationNavigationPtrOutput() GuestConfigurationNavigationPtrOutput {
	return o
}

func (o GuestConfigurationNavigationPtrOutput) ToGuestConfigurationNavigationPtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationPtrOutput {
	return o
}

func (o GuestConfigurationNavigationPtrOutput) Elem() GuestConfigurationNavigationOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) GuestConfigurationNavigation {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationNavigation
		return ret
	}).(GuestConfigurationNavigationOutput)
}

func (o GuestConfigurationNavigationPtrOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.AssignmentType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ConfigurationParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) []ConfigurationParameter {
		if v == nil {
			return nil
		}
		return v.ConfigurationParameter
	}).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ConfigurationProtectedParameter() ConfigurationParameterArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) []ConfigurationParameter {
		if v == nil {
			return nil
		}
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterArrayOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ConfigurationSetting() ConfigurationSettingPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *ConfigurationSetting {
		if v == nil {
			return nil
		}
		return v.ConfigurationSetting
	}).(ConfigurationSettingPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.ContentUri
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationPtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type GuestConfigurationNavigationResponse struct {
	AssignmentType                  *string                          `pulumi:"assignmentType"`
	ConfigurationParameter          []ConfigurationParameterResponse `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter []ConfigurationParameterResponse `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            *ConfigurationSettingResponse    `pulumi:"configurationSetting"`
	ContentHash                     *string                          `pulumi:"contentHash"`
	ContentType                     string                           `pulumi:"contentType"`
	ContentUri                      *string                          `pulumi:"contentUri"`
	Kind                            *string                          `pulumi:"kind"`
	Name                            *string                          `pulumi:"name"`
	Version                         *string                          `pulumi:"version"`
}





type GuestConfigurationNavigationResponseInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationResponseOutput() GuestConfigurationNavigationResponseOutput
	ToGuestConfigurationNavigationResponseOutputWithContext(context.Context) GuestConfigurationNavigationResponseOutput
}

type GuestConfigurationNavigationResponseArgs struct {
	AssignmentType                  pulumi.StringPtrInput                    `pulumi:"assignmentType"`
	ConfigurationParameter          ConfigurationParameterResponseArrayInput `pulumi:"configurationParameter"`
	ConfigurationProtectedParameter ConfigurationParameterResponseArrayInput `pulumi:"configurationProtectedParameter"`
	ConfigurationSetting            ConfigurationSettingResponsePtrInput     `pulumi:"configurationSetting"`
	ContentHash                     pulumi.StringPtrInput                    `pulumi:"contentHash"`
	ContentType                     pulumi.StringInput                       `pulumi:"contentType"`
	ContentUri                      pulumi.StringPtrInput                    `pulumi:"contentUri"`
	Kind                            pulumi.StringPtrInput                    `pulumi:"kind"`
	Name                            pulumi.StringPtrInput                    `pulumi:"name"`
	Version                         pulumi.StringPtrInput                    `pulumi:"version"`
}

func (GuestConfigurationNavigationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponseOutput() GuestConfigurationNavigationResponseOutput {
	return i.ToGuestConfigurationNavigationResponseOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponseOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationResponseOutput)
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return i.ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Background())
}

func (i GuestConfigurationNavigationResponseArgs) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationResponseOutput).ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx)
}









type GuestConfigurationNavigationResponsePtrInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput
	ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Context) GuestConfigurationNavigationResponsePtrOutput
}

type guestConfigurationNavigationResponsePtrType GuestConfigurationNavigationResponseArgs

func GuestConfigurationNavigationResponsePtr(v *GuestConfigurationNavigationResponseArgs) GuestConfigurationNavigationResponsePtrInput {
	return (*guestConfigurationNavigationResponsePtrType)(v)
}

func (*guestConfigurationNavigationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (i *guestConfigurationNavigationResponsePtrType) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return i.ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Background())
}

func (i *guestConfigurationNavigationResponsePtrType) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GuestConfigurationNavigationResponsePtrOutput)
}

type GuestConfigurationNavigationResponseOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponseOutput() GuestConfigurationNavigationResponseOutput {
	return o
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponseOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponseOutput {
	return o
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return o.ToGuestConfigurationNavigationResponsePtrOutputWithContext(context.Background())
}

func (o GuestConfigurationNavigationResponseOutput) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GuestConfigurationNavigationResponse) *GuestConfigurationNavigationResponse {
		return &v
	}).(GuestConfigurationNavigationResponsePtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.AssignmentType }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		return v.ConfigurationParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationProtectedParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationSetting() ConfigurationSettingResponsePtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *ConfigurationSettingResponse {
		return v.ConfigurationSetting
	}).(ConfigurationSettingResponsePtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.ContentHash }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentType() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) string { return v.ContentType }).(pulumi.StringOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.ContentUri }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type GuestConfigurationNavigationResponsePtrOutput struct{ *pulumi.OutputState }

func (GuestConfigurationNavigationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GuestConfigurationNavigationResponse)(nil)).Elem()
}

func (o GuestConfigurationNavigationResponsePtrOutput) ToGuestConfigurationNavigationResponsePtrOutput() GuestConfigurationNavigationResponsePtrOutput {
	return o
}

func (o GuestConfigurationNavigationResponsePtrOutput) ToGuestConfigurationNavigationResponsePtrOutputWithContext(ctx context.Context) GuestConfigurationNavigationResponsePtrOutput {
	return o
}

func (o GuestConfigurationNavigationResponsePtrOutput) Elem() GuestConfigurationNavigationResponseOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) GuestConfigurationNavigationResponse {
		if v != nil {
			return *v
		}
		var ret GuestConfigurationNavigationResponse
		return ret
	}).(GuestConfigurationNavigationResponseOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) AssignmentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AssignmentType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationProtectedParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationProtectedParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationSetting() ConfigurationSettingResponsePtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *ConfigurationSettingResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationSetting
	}).(ConfigurationSettingResponsePtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentHash() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContentType
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ContentUri
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type VMSSVMInfoResponse struct {
	ComplianceStatus      string `pulumi:"complianceStatus"`
	LastComplianceChecked string `pulumi:"lastComplianceChecked"`
	LatestReportId        string `pulumi:"latestReportId"`
	VmId                  string `pulumi:"vmId"`
	VmResourceId          string `pulumi:"vmResourceId"`
}





type VMSSVMInfoResponseInput interface {
	pulumi.Input

	ToVMSSVMInfoResponseOutput() VMSSVMInfoResponseOutput
	ToVMSSVMInfoResponseOutputWithContext(context.Context) VMSSVMInfoResponseOutput
}

type VMSSVMInfoResponseArgs struct {
	ComplianceStatus      pulumi.StringInput `pulumi:"complianceStatus"`
	LastComplianceChecked pulumi.StringInput `pulumi:"lastComplianceChecked"`
	LatestReportId        pulumi.StringInput `pulumi:"latestReportId"`
	VmId                  pulumi.StringInput `pulumi:"vmId"`
	VmResourceId          pulumi.StringInput `pulumi:"vmResourceId"`
}

func (VMSSVMInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSVMInfoResponse)(nil)).Elem()
}

func (i VMSSVMInfoResponseArgs) ToVMSSVMInfoResponseOutput() VMSSVMInfoResponseOutput {
	return i.ToVMSSVMInfoResponseOutputWithContext(context.Background())
}

func (i VMSSVMInfoResponseArgs) ToVMSSVMInfoResponseOutputWithContext(ctx context.Context) VMSSVMInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSVMInfoResponseOutput)
}





type VMSSVMInfoResponseArrayInput interface {
	pulumi.Input

	ToVMSSVMInfoResponseArrayOutput() VMSSVMInfoResponseArrayOutput
	ToVMSSVMInfoResponseArrayOutputWithContext(context.Context) VMSSVMInfoResponseArrayOutput
}

type VMSSVMInfoResponseArray []VMSSVMInfoResponseInput

func (VMSSVMInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSVMInfoResponse)(nil)).Elem()
}

func (i VMSSVMInfoResponseArray) ToVMSSVMInfoResponseArrayOutput() VMSSVMInfoResponseArrayOutput {
	return i.ToVMSSVMInfoResponseArrayOutputWithContext(context.Background())
}

func (i VMSSVMInfoResponseArray) ToVMSSVMInfoResponseArrayOutputWithContext(ctx context.Context) VMSSVMInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VMSSVMInfoResponseArrayOutput)
}

type VMSSVMInfoResponseOutput struct{ *pulumi.OutputState }

func (VMSSVMInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VMSSVMInfoResponse)(nil)).Elem()
}

func (o VMSSVMInfoResponseOutput) ToVMSSVMInfoResponseOutput() VMSSVMInfoResponseOutput {
	return o
}

func (o VMSSVMInfoResponseOutput) ToVMSSVMInfoResponseOutputWithContext(ctx context.Context) VMSSVMInfoResponseOutput {
	return o
}

func (o VMSSVMInfoResponseOutput) ComplianceStatus() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.ComplianceStatus }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) LastComplianceChecked() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.LastComplianceChecked }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) LatestReportId() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.LatestReportId }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) VmId() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.VmId }).(pulumi.StringOutput)
}

func (o VMSSVMInfoResponseOutput) VmResourceId() pulumi.StringOutput {
	return o.ApplyT(func(v VMSSVMInfoResponse) string { return v.VmResourceId }).(pulumi.StringOutput)
}

type VMSSVMInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (VMSSVMInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VMSSVMInfoResponse)(nil)).Elem()
}

func (o VMSSVMInfoResponseArrayOutput) ToVMSSVMInfoResponseArrayOutput() VMSSVMInfoResponseArrayOutput {
	return o
}

func (o VMSSVMInfoResponseArrayOutput) ToVMSSVMInfoResponseArrayOutputWithContext(ctx context.Context) VMSSVMInfoResponseArrayOutput {
	return o
}

func (o VMSSVMInfoResponseArrayOutput) Index(i pulumi.IntInput) VMSSVMInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VMSSVMInfoResponse {
		return vs[0].([]VMSSVMInfoResponse)[vs[1].(int)]
	}).(VMSSVMInfoResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(ConfigurationParameterOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponsePtrOutput{})
	pulumi.RegisterOutputType(VMSSVMInfoResponseOutput{})
	pulumi.RegisterOutputType(VMSSVMInfoResponseArrayOutput{})
}
