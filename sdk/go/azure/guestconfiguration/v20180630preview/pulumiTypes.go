


package v20180630preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ConfigurationParameterResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
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

func (o ConfigurationParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ConfigurationParameterResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ConfigurationParameterResponse) string { return v.Value }).(pulumi.StringOutput)
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


func (val *ConfigurationSetting) Defaults() *ConfigurationSetting {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ConfigurationModeFrequencyMins) {
		configurationModeFrequencyMins_ := 15.0
		tmp.ConfigurationModeFrequencyMins = &configurationModeFrequencyMins_
	}
	if isZero(tmp.RefreshFrequencyMins) {
		refreshFrequencyMins_ := 30.0
		tmp.RefreshFrequencyMins = &refreshFrequencyMins_
	}
	return &tmp
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


func (val *ConfigurationSettingResponse) Defaults() *ConfigurationSettingResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ConfigurationModeFrequencyMins) {
		configurationModeFrequencyMins_ := 15.0
		tmp.ConfigurationModeFrequencyMins = &configurationModeFrequencyMins_
	}
	if isZero(tmp.RefreshFrequencyMins) {
		refreshFrequencyMins_ := 30.0
		tmp.RefreshFrequencyMins = &refreshFrequencyMins_
	}
	return &tmp
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


func (val *GuestConfigurationAssignmentProperties) Defaults() *GuestConfigurationAssignmentProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.GuestConfiguration = tmp.GuestConfiguration.Defaults()

	return &tmp
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
	ProvisioningState           string                                `pulumi:"provisioningState"`
}


func (val *GuestConfigurationAssignmentPropertiesResponse) Defaults() *GuestConfigurationAssignmentPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.GuestConfiguration = tmp.GuestConfiguration.Defaults()

	return &tmp
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

func (o GuestConfigurationAssignmentPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationAssignmentPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type GuestConfigurationNavigation struct {
	ConfigurationSetting *ConfigurationSetting `pulumi:"configurationSetting"`
	Kind                 *string               `pulumi:"kind"`
	Name                 *string               `pulumi:"name"`
	Version              *string               `pulumi:"version"`
}


func (val *GuestConfigurationNavigation) Defaults() *GuestConfigurationNavigation {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConfigurationSetting = tmp.ConfigurationSetting.Defaults()

	return &tmp
}





type GuestConfigurationNavigationInput interface {
	pulumi.Input

	ToGuestConfigurationNavigationOutput() GuestConfigurationNavigationOutput
	ToGuestConfigurationNavigationOutputWithContext(context.Context) GuestConfigurationNavigationOutput
}

type GuestConfigurationNavigationArgs struct {
	ConfigurationSetting ConfigurationSettingPtrInput `pulumi:"configurationSetting"`
	Kind                 pulumi.StringPtrInput        `pulumi:"kind"`
	Name                 pulumi.StringPtrInput        `pulumi:"name"`
	Version              pulumi.StringPtrInput        `pulumi:"version"`
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

func (o GuestConfigurationNavigationOutput) ConfigurationSetting() ConfigurationSettingPtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigation) *ConfigurationSetting { return v.ConfigurationSetting }).(ConfigurationSettingPtrOutput)
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

func (o GuestConfigurationNavigationPtrOutput) ConfigurationSetting() ConfigurationSettingPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigation) *ConfigurationSetting {
		if v == nil {
			return nil
		}
		return v.ConfigurationSetting
	}).(ConfigurationSettingPtrOutput)
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
	ConfigurationParameter []ConfigurationParameterResponse `pulumi:"configurationParameter"`
	ConfigurationSetting   *ConfigurationSettingResponse    `pulumi:"configurationSetting"`
	ContentHash            string                           `pulumi:"contentHash"`
	ContentUri             string                           `pulumi:"contentUri"`
	Kind                   *string                          `pulumi:"kind"`
	Name                   *string                          `pulumi:"name"`
	Version                *string                          `pulumi:"version"`
}


func (val *GuestConfigurationNavigationResponse) Defaults() *GuestConfigurationNavigationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConfigurationSetting = tmp.ConfigurationSetting.Defaults()

	return &tmp
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

func (o GuestConfigurationNavigationResponseOutput) ConfigurationParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		return v.ConfigurationParameter
	}).(ConfigurationParameterResponseArrayOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ConfigurationSetting() ConfigurationSettingResponsePtrOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) *ConfigurationSettingResponse {
		return v.ConfigurationSetting
	}).(ConfigurationSettingResponsePtrOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentHash() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) string { return v.ContentHash }).(pulumi.StringOutput)
}

func (o GuestConfigurationNavigationResponseOutput) ContentUri() pulumi.StringOutput {
	return o.ApplyT(func(v GuestConfigurationNavigationResponse) string { return v.ContentUri }).(pulumi.StringOutput)
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

func (o GuestConfigurationNavigationResponsePtrOutput) ConfigurationParameter() ConfigurationParameterResponseArrayOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) []ConfigurationParameterResponse {
		if v == nil {
			return nil
		}
		return v.ConfigurationParameter
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
		return &v.ContentHash
	}).(pulumi.StringPtrOutput)
}

func (o GuestConfigurationNavigationResponsePtrOutput) ContentUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GuestConfigurationNavigationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ContentUri
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

func init() {
	pulumi.RegisterOutputType(ConfigurationParameterResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationSettingResponsePtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationAssignmentPropertiesResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationPtrOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponseOutput{})
	pulumi.RegisterOutputType(GuestConfigurationNavigationResponsePtrOutput{})
}
