


package v20211015

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AnalyticalStorageConfiguration struct {
	SchemaType *string `pulumi:"schemaType"`
}





type AnalyticalStorageConfigurationInput interface {
	pulumi.Input

	ToAnalyticalStorageConfigurationOutput() AnalyticalStorageConfigurationOutput
	ToAnalyticalStorageConfigurationOutputWithContext(context.Context) AnalyticalStorageConfigurationOutput
}

type AnalyticalStorageConfigurationArgs struct {
	SchemaType pulumi.StringPtrInput `pulumi:"schemaType"`
}

func (AnalyticalStorageConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticalStorageConfiguration)(nil)).Elem()
}

func (i AnalyticalStorageConfigurationArgs) ToAnalyticalStorageConfigurationOutput() AnalyticalStorageConfigurationOutput {
	return i.ToAnalyticalStorageConfigurationOutputWithContext(context.Background())
}

func (i AnalyticalStorageConfigurationArgs) ToAnalyticalStorageConfigurationOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticalStorageConfigurationOutput)
}

func (i AnalyticalStorageConfigurationArgs) ToAnalyticalStorageConfigurationPtrOutput() AnalyticalStorageConfigurationPtrOutput {
	return i.ToAnalyticalStorageConfigurationPtrOutputWithContext(context.Background())
}

func (i AnalyticalStorageConfigurationArgs) ToAnalyticalStorageConfigurationPtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticalStorageConfigurationOutput).ToAnalyticalStorageConfigurationPtrOutputWithContext(ctx)
}









type AnalyticalStorageConfigurationPtrInput interface {
	pulumi.Input

	ToAnalyticalStorageConfigurationPtrOutput() AnalyticalStorageConfigurationPtrOutput
	ToAnalyticalStorageConfigurationPtrOutputWithContext(context.Context) AnalyticalStorageConfigurationPtrOutput
}

type analyticalStorageConfigurationPtrType AnalyticalStorageConfigurationArgs

func AnalyticalStorageConfigurationPtr(v *AnalyticalStorageConfigurationArgs) AnalyticalStorageConfigurationPtrInput {
	return (*analyticalStorageConfigurationPtrType)(v)
}

func (*analyticalStorageConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticalStorageConfiguration)(nil)).Elem()
}

func (i *analyticalStorageConfigurationPtrType) ToAnalyticalStorageConfigurationPtrOutput() AnalyticalStorageConfigurationPtrOutput {
	return i.ToAnalyticalStorageConfigurationPtrOutputWithContext(context.Background())
}

func (i *analyticalStorageConfigurationPtrType) ToAnalyticalStorageConfigurationPtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticalStorageConfigurationPtrOutput)
}

type AnalyticalStorageConfigurationOutput struct{ *pulumi.OutputState }

func (AnalyticalStorageConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticalStorageConfiguration)(nil)).Elem()
}

func (o AnalyticalStorageConfigurationOutput) ToAnalyticalStorageConfigurationOutput() AnalyticalStorageConfigurationOutput {
	return o
}

func (o AnalyticalStorageConfigurationOutput) ToAnalyticalStorageConfigurationOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationOutput {
	return o
}

func (o AnalyticalStorageConfigurationOutput) ToAnalyticalStorageConfigurationPtrOutput() AnalyticalStorageConfigurationPtrOutput {
	return o.ToAnalyticalStorageConfigurationPtrOutputWithContext(context.Background())
}

func (o AnalyticalStorageConfigurationOutput) ToAnalyticalStorageConfigurationPtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AnalyticalStorageConfiguration) *AnalyticalStorageConfiguration {
		return &v
	}).(AnalyticalStorageConfigurationPtrOutput)
}

func (o AnalyticalStorageConfigurationOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AnalyticalStorageConfiguration) *string { return v.SchemaType }).(pulumi.StringPtrOutput)
}

type AnalyticalStorageConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AnalyticalStorageConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticalStorageConfiguration)(nil)).Elem()
}

func (o AnalyticalStorageConfigurationPtrOutput) ToAnalyticalStorageConfigurationPtrOutput() AnalyticalStorageConfigurationPtrOutput {
	return o
}

func (o AnalyticalStorageConfigurationPtrOutput) ToAnalyticalStorageConfigurationPtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationPtrOutput {
	return o
}

func (o AnalyticalStorageConfigurationPtrOutput) Elem() AnalyticalStorageConfigurationOutput {
	return o.ApplyT(func(v *AnalyticalStorageConfiguration) AnalyticalStorageConfiguration {
		if v != nil {
			return *v
		}
		var ret AnalyticalStorageConfiguration
		return ret
	}).(AnalyticalStorageConfigurationOutput)
}

func (o AnalyticalStorageConfigurationPtrOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticalStorageConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SchemaType
	}).(pulumi.StringPtrOutput)
}

type AnalyticalStorageConfigurationResponse struct {
	SchemaType *string `pulumi:"schemaType"`
}





type AnalyticalStorageConfigurationResponseInput interface {
	pulumi.Input

	ToAnalyticalStorageConfigurationResponseOutput() AnalyticalStorageConfigurationResponseOutput
	ToAnalyticalStorageConfigurationResponseOutputWithContext(context.Context) AnalyticalStorageConfigurationResponseOutput
}

type AnalyticalStorageConfigurationResponseArgs struct {
	SchemaType pulumi.StringPtrInput `pulumi:"schemaType"`
}

func (AnalyticalStorageConfigurationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticalStorageConfigurationResponse)(nil)).Elem()
}

func (i AnalyticalStorageConfigurationResponseArgs) ToAnalyticalStorageConfigurationResponseOutput() AnalyticalStorageConfigurationResponseOutput {
	return i.ToAnalyticalStorageConfigurationResponseOutputWithContext(context.Background())
}

func (i AnalyticalStorageConfigurationResponseArgs) ToAnalyticalStorageConfigurationResponseOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticalStorageConfigurationResponseOutput)
}

func (i AnalyticalStorageConfigurationResponseArgs) ToAnalyticalStorageConfigurationResponsePtrOutput() AnalyticalStorageConfigurationResponsePtrOutput {
	return i.ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i AnalyticalStorageConfigurationResponseArgs) ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticalStorageConfigurationResponseOutput).ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(ctx)
}









type AnalyticalStorageConfigurationResponsePtrInput interface {
	pulumi.Input

	ToAnalyticalStorageConfigurationResponsePtrOutput() AnalyticalStorageConfigurationResponsePtrOutput
	ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(context.Context) AnalyticalStorageConfigurationResponsePtrOutput
}

type analyticalStorageConfigurationResponsePtrType AnalyticalStorageConfigurationResponseArgs

func AnalyticalStorageConfigurationResponsePtr(v *AnalyticalStorageConfigurationResponseArgs) AnalyticalStorageConfigurationResponsePtrInput {
	return (*analyticalStorageConfigurationResponsePtrType)(v)
}

func (*analyticalStorageConfigurationResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticalStorageConfigurationResponse)(nil)).Elem()
}

func (i *analyticalStorageConfigurationResponsePtrType) ToAnalyticalStorageConfigurationResponsePtrOutput() AnalyticalStorageConfigurationResponsePtrOutput {
	return i.ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(context.Background())
}

func (i *analyticalStorageConfigurationResponsePtrType) ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AnalyticalStorageConfigurationResponsePtrOutput)
}

type AnalyticalStorageConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AnalyticalStorageConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AnalyticalStorageConfigurationResponse)(nil)).Elem()
}

func (o AnalyticalStorageConfigurationResponseOutput) ToAnalyticalStorageConfigurationResponseOutput() AnalyticalStorageConfigurationResponseOutput {
	return o
}

func (o AnalyticalStorageConfigurationResponseOutput) ToAnalyticalStorageConfigurationResponseOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationResponseOutput {
	return o
}

func (o AnalyticalStorageConfigurationResponseOutput) ToAnalyticalStorageConfigurationResponsePtrOutput() AnalyticalStorageConfigurationResponsePtrOutput {
	return o.ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(context.Background())
}

func (o AnalyticalStorageConfigurationResponseOutput) ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AnalyticalStorageConfigurationResponse) *AnalyticalStorageConfigurationResponse {
		return &v
	}).(AnalyticalStorageConfigurationResponsePtrOutput)
}

func (o AnalyticalStorageConfigurationResponseOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AnalyticalStorageConfigurationResponse) *string { return v.SchemaType }).(pulumi.StringPtrOutput)
}

type AnalyticalStorageConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AnalyticalStorageConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AnalyticalStorageConfigurationResponse)(nil)).Elem()
}

func (o AnalyticalStorageConfigurationResponsePtrOutput) ToAnalyticalStorageConfigurationResponsePtrOutput() AnalyticalStorageConfigurationResponsePtrOutput {
	return o
}

func (o AnalyticalStorageConfigurationResponsePtrOutput) ToAnalyticalStorageConfigurationResponsePtrOutputWithContext(ctx context.Context) AnalyticalStorageConfigurationResponsePtrOutput {
	return o
}

func (o AnalyticalStorageConfigurationResponsePtrOutput) Elem() AnalyticalStorageConfigurationResponseOutput {
	return o.ApplyT(func(v *AnalyticalStorageConfigurationResponse) AnalyticalStorageConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AnalyticalStorageConfigurationResponse
		return ret
	}).(AnalyticalStorageConfigurationResponseOutput)
}

func (o AnalyticalStorageConfigurationResponsePtrOutput) SchemaType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AnalyticalStorageConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.SchemaType
	}).(pulumi.StringPtrOutput)
}

type ApiProperties struct {
	ServerVersion *string `pulumi:"serverVersion"`
}





type ApiPropertiesInput interface {
	pulumi.Input

	ToApiPropertiesOutput() ApiPropertiesOutput
	ToApiPropertiesOutputWithContext(context.Context) ApiPropertiesOutput
}

type ApiPropertiesArgs struct {
	ServerVersion pulumi.StringPtrInput `pulumi:"serverVersion"`
}

func (ApiPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (i ApiPropertiesArgs) ToApiPropertiesOutput() ApiPropertiesOutput {
	return i.ToApiPropertiesOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput)
}

func (i ApiPropertiesArgs) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return i.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (i ApiPropertiesArgs) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesOutput).ToApiPropertiesPtrOutputWithContext(ctx)
}









type ApiPropertiesPtrInput interface {
	pulumi.Input

	ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput
	ToApiPropertiesPtrOutputWithContext(context.Context) ApiPropertiesPtrOutput
}

type apiPropertiesPtrType ApiPropertiesArgs

func ApiPropertiesPtr(v *ApiPropertiesArgs) ApiPropertiesPtrInput {
	return (*apiPropertiesPtrType)(v)
}

func (*apiPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiProperties)(nil)).Elem()
}

func (i *apiPropertiesPtrType) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return i.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (i *apiPropertiesPtrType) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesPtrOutput)
}

type ApiPropertiesOutput struct{ *pulumi.OutputState }

func (ApiPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesOutput) ToApiPropertiesOutput() ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesOutputWithContext(ctx context.Context) ApiPropertiesOutput {
	return o
}

func (o ApiPropertiesOutput) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return o.ToApiPropertiesPtrOutputWithContext(context.Background())
}

func (o ApiPropertiesOutput) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiProperties) *ApiProperties {
		return &v
	}).(ApiPropertiesPtrOutput)
}

func (o ApiPropertiesOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiProperties) *string { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

type ApiPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ApiPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiProperties)(nil)).Elem()
}

func (o ApiPropertiesPtrOutput) ToApiPropertiesPtrOutput() ApiPropertiesPtrOutput {
	return o
}

func (o ApiPropertiesPtrOutput) ToApiPropertiesPtrOutputWithContext(ctx context.Context) ApiPropertiesPtrOutput {
	return o
}

func (o ApiPropertiesPtrOutput) Elem() ApiPropertiesOutput {
	return o.ApplyT(func(v *ApiProperties) ApiProperties {
		if v != nil {
			return *v
		}
		var ret ApiProperties
		return ret
	}).(ApiPropertiesOutput)
}

func (o ApiPropertiesPtrOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServerVersion
	}).(pulumi.StringPtrOutput)
}

type ApiPropertiesResponse struct {
	ServerVersion *string `pulumi:"serverVersion"`
}





type ApiPropertiesResponseInput interface {
	pulumi.Input

	ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput
	ToApiPropertiesResponseOutputWithContext(context.Context) ApiPropertiesResponseOutput
}

type ApiPropertiesResponseArgs struct {
	ServerVersion pulumi.StringPtrInput `pulumi:"serverVersion"`
}

func (ApiPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPropertiesResponse)(nil)).Elem()
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput {
	return i.ToApiPropertiesResponseOutputWithContext(context.Background())
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponseOutputWithContext(ctx context.Context) ApiPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesResponseOutput)
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return i.ToApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ApiPropertiesResponseArgs) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesResponseOutput).ToApiPropertiesResponsePtrOutputWithContext(ctx)
}









type ApiPropertiesResponsePtrInput interface {
	pulumi.Input

	ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput
	ToApiPropertiesResponsePtrOutputWithContext(context.Context) ApiPropertiesResponsePtrOutput
}

type apiPropertiesResponsePtrType ApiPropertiesResponseArgs

func ApiPropertiesResponsePtr(v *ApiPropertiesResponseArgs) ApiPropertiesResponsePtrInput {
	return (*apiPropertiesResponsePtrType)(v)
}

func (*apiPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPropertiesResponse)(nil)).Elem()
}

func (i *apiPropertiesResponsePtrType) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return i.ToApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *apiPropertiesResponsePtrType) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiPropertiesResponsePtrOutput)
}

type ApiPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutput() ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponseOutputWithContext(ctx context.Context) ApiPropertiesResponseOutput {
	return o
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return o.ToApiPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ApiPropertiesResponseOutput) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiPropertiesResponse) *ApiPropertiesResponse {
		return &v
	}).(ApiPropertiesResponsePtrOutput)
}

func (o ApiPropertiesResponseOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiPropertiesResponse) *string { return v.ServerVersion }).(pulumi.StringPtrOutput)
}

type ApiPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiPropertiesResponse)(nil)).Elem()
}

func (o ApiPropertiesResponsePtrOutput) ToApiPropertiesResponsePtrOutput() ApiPropertiesResponsePtrOutput {
	return o
}

func (o ApiPropertiesResponsePtrOutput) ToApiPropertiesResponsePtrOutputWithContext(ctx context.Context) ApiPropertiesResponsePtrOutput {
	return o
}

func (o ApiPropertiesResponsePtrOutput) Elem() ApiPropertiesResponseOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) ApiPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ApiPropertiesResponse
		return ret
	}).(ApiPropertiesResponseOutput)
}

func (o ApiPropertiesResponsePtrOutput) ServerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerVersion
	}).(pulumi.StringPtrOutput)
}

type AutoscaleSettings struct {
	MaxThroughput *int `pulumi:"maxThroughput"`
}





type AutoscaleSettingsInput interface {
	pulumi.Input

	ToAutoscaleSettingsOutput() AutoscaleSettingsOutput
	ToAutoscaleSettingsOutputWithContext(context.Context) AutoscaleSettingsOutput
}

type AutoscaleSettingsArgs struct {
	MaxThroughput pulumi.IntPtrInput `pulumi:"maxThroughput"`
}

func (AutoscaleSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettings)(nil)).Elem()
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsOutput() AutoscaleSettingsOutput {
	return i.ToAutoscaleSettingsOutputWithContext(context.Background())
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsOutputWithContext(ctx context.Context) AutoscaleSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsOutput)
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return i.ToAutoscaleSettingsPtrOutputWithContext(context.Background())
}

func (i AutoscaleSettingsArgs) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsOutput).ToAutoscaleSettingsPtrOutputWithContext(ctx)
}









type AutoscaleSettingsPtrInput interface {
	pulumi.Input

	ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput
	ToAutoscaleSettingsPtrOutputWithContext(context.Context) AutoscaleSettingsPtrOutput
}

type autoscaleSettingsPtrType AutoscaleSettingsArgs

func AutoscaleSettingsPtr(v *AutoscaleSettingsArgs) AutoscaleSettingsPtrInput {
	return (*autoscaleSettingsPtrType)(v)
}

func (*autoscaleSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettings)(nil)).Elem()
}

func (i *autoscaleSettingsPtrType) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return i.ToAutoscaleSettingsPtrOutputWithContext(context.Background())
}

func (i *autoscaleSettingsPtrType) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsPtrOutput)
}

type AutoscaleSettingsOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettings)(nil)).Elem()
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsOutput() AutoscaleSettingsOutput {
	return o
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsOutputWithContext(ctx context.Context) AutoscaleSettingsOutput {
	return o
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return o.ToAutoscaleSettingsPtrOutputWithContext(context.Background())
}

func (o AutoscaleSettingsOutput) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoscaleSettings) *AutoscaleSettings {
		return &v
	}).(AutoscaleSettingsPtrOutput)
}

func (o AutoscaleSettingsOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleSettings) *int { return v.MaxThroughput }).(pulumi.IntPtrOutput)
}

type AutoscaleSettingsPtrOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettings)(nil)).Elem()
}

func (o AutoscaleSettingsPtrOutput) ToAutoscaleSettingsPtrOutput() AutoscaleSettingsPtrOutput {
	return o
}

func (o AutoscaleSettingsPtrOutput) ToAutoscaleSettingsPtrOutputWithContext(ctx context.Context) AutoscaleSettingsPtrOutput {
	return o
}

func (o AutoscaleSettingsPtrOutput) Elem() AutoscaleSettingsOutput {
	return o.ApplyT(func(v *AutoscaleSettings) AutoscaleSettings {
		if v != nil {
			return *v
		}
		var ret AutoscaleSettings
		return ret
	}).(AutoscaleSettingsOutput)
}

func (o AutoscaleSettingsPtrOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleSettings) *int {
		if v == nil {
			return nil
		}
		return v.MaxThroughput
	}).(pulumi.IntPtrOutput)
}

type AutoscaleSettingsResponse struct {
	MaxThroughput *int `pulumi:"maxThroughput"`
}





type AutoscaleSettingsResponseInput interface {
	pulumi.Input

	ToAutoscaleSettingsResponseOutput() AutoscaleSettingsResponseOutput
	ToAutoscaleSettingsResponseOutputWithContext(context.Context) AutoscaleSettingsResponseOutput
}

type AutoscaleSettingsResponseArgs struct {
	MaxThroughput pulumi.IntPtrInput `pulumi:"maxThroughput"`
}

func (AutoscaleSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettingsResponse)(nil)).Elem()
}

func (i AutoscaleSettingsResponseArgs) ToAutoscaleSettingsResponseOutput() AutoscaleSettingsResponseOutput {
	return i.ToAutoscaleSettingsResponseOutputWithContext(context.Background())
}

func (i AutoscaleSettingsResponseArgs) ToAutoscaleSettingsResponseOutputWithContext(ctx context.Context) AutoscaleSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsResponseOutput)
}

func (i AutoscaleSettingsResponseArgs) ToAutoscaleSettingsResponsePtrOutput() AutoscaleSettingsResponsePtrOutput {
	return i.ToAutoscaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i AutoscaleSettingsResponseArgs) ToAutoscaleSettingsResponsePtrOutputWithContext(ctx context.Context) AutoscaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsResponseOutput).ToAutoscaleSettingsResponsePtrOutputWithContext(ctx)
}









type AutoscaleSettingsResponsePtrInput interface {
	pulumi.Input

	ToAutoscaleSettingsResponsePtrOutput() AutoscaleSettingsResponsePtrOutput
	ToAutoscaleSettingsResponsePtrOutputWithContext(context.Context) AutoscaleSettingsResponsePtrOutput
}

type autoscaleSettingsResponsePtrType AutoscaleSettingsResponseArgs

func AutoscaleSettingsResponsePtr(v *AutoscaleSettingsResponseArgs) AutoscaleSettingsResponsePtrInput {
	return (*autoscaleSettingsResponsePtrType)(v)
}

func (*autoscaleSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettingsResponse)(nil)).Elem()
}

func (i *autoscaleSettingsResponsePtrType) ToAutoscaleSettingsResponsePtrOutput() AutoscaleSettingsResponsePtrOutput {
	return i.ToAutoscaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *autoscaleSettingsResponsePtrType) ToAutoscaleSettingsResponsePtrOutputWithContext(ctx context.Context) AutoscaleSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoscaleSettingsResponsePtrOutput)
}

type AutoscaleSettingsResponseOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoscaleSettingsResponse)(nil)).Elem()
}

func (o AutoscaleSettingsResponseOutput) ToAutoscaleSettingsResponseOutput() AutoscaleSettingsResponseOutput {
	return o
}

func (o AutoscaleSettingsResponseOutput) ToAutoscaleSettingsResponseOutputWithContext(ctx context.Context) AutoscaleSettingsResponseOutput {
	return o
}

func (o AutoscaleSettingsResponseOutput) ToAutoscaleSettingsResponsePtrOutput() AutoscaleSettingsResponsePtrOutput {
	return o.ToAutoscaleSettingsResponsePtrOutputWithContext(context.Background())
}

func (o AutoscaleSettingsResponseOutput) ToAutoscaleSettingsResponsePtrOutputWithContext(ctx context.Context) AutoscaleSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoscaleSettingsResponse) *AutoscaleSettingsResponse {
		return &v
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o AutoscaleSettingsResponseOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoscaleSettingsResponse) *int { return v.MaxThroughput }).(pulumi.IntPtrOutput)
}

type AutoscaleSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoscaleSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoscaleSettingsResponse)(nil)).Elem()
}

func (o AutoscaleSettingsResponsePtrOutput) ToAutoscaleSettingsResponsePtrOutput() AutoscaleSettingsResponsePtrOutput {
	return o
}

func (o AutoscaleSettingsResponsePtrOutput) ToAutoscaleSettingsResponsePtrOutputWithContext(ctx context.Context) AutoscaleSettingsResponsePtrOutput {
	return o
}

func (o AutoscaleSettingsResponsePtrOutput) Elem() AutoscaleSettingsResponseOutput {
	return o.ApplyT(func(v *AutoscaleSettingsResponse) AutoscaleSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AutoscaleSettingsResponse
		return ret
	}).(AutoscaleSettingsResponseOutput)
}

func (o AutoscaleSettingsResponsePtrOutput) MaxThroughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoscaleSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxThroughput
	}).(pulumi.IntPtrOutput)
}

type BackupPolicyMigrationState struct {
	StartTime  *string `pulumi:"startTime"`
	Status     *string `pulumi:"status"`
	TargetType *string `pulumi:"targetType"`
}





type BackupPolicyMigrationStateInput interface {
	pulumi.Input

	ToBackupPolicyMigrationStateOutput() BackupPolicyMigrationStateOutput
	ToBackupPolicyMigrationStateOutputWithContext(context.Context) BackupPolicyMigrationStateOutput
}

type BackupPolicyMigrationStateArgs struct {
	StartTime  pulumi.StringPtrInput `pulumi:"startTime"`
	Status     pulumi.StringPtrInput `pulumi:"status"`
	TargetType pulumi.StringPtrInput `pulumi:"targetType"`
}

func (BackupPolicyMigrationStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyMigrationState)(nil)).Elem()
}

func (i BackupPolicyMigrationStateArgs) ToBackupPolicyMigrationStateOutput() BackupPolicyMigrationStateOutput {
	return i.ToBackupPolicyMigrationStateOutputWithContext(context.Background())
}

func (i BackupPolicyMigrationStateArgs) ToBackupPolicyMigrationStateOutputWithContext(ctx context.Context) BackupPolicyMigrationStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMigrationStateOutput)
}

func (i BackupPolicyMigrationStateArgs) ToBackupPolicyMigrationStatePtrOutput() BackupPolicyMigrationStatePtrOutput {
	return i.ToBackupPolicyMigrationStatePtrOutputWithContext(context.Background())
}

func (i BackupPolicyMigrationStateArgs) ToBackupPolicyMigrationStatePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMigrationStateOutput).ToBackupPolicyMigrationStatePtrOutputWithContext(ctx)
}









type BackupPolicyMigrationStatePtrInput interface {
	pulumi.Input

	ToBackupPolicyMigrationStatePtrOutput() BackupPolicyMigrationStatePtrOutput
	ToBackupPolicyMigrationStatePtrOutputWithContext(context.Context) BackupPolicyMigrationStatePtrOutput
}

type backupPolicyMigrationStatePtrType BackupPolicyMigrationStateArgs

func BackupPolicyMigrationStatePtr(v *BackupPolicyMigrationStateArgs) BackupPolicyMigrationStatePtrInput {
	return (*backupPolicyMigrationStatePtrType)(v)
}

func (*backupPolicyMigrationStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyMigrationState)(nil)).Elem()
}

func (i *backupPolicyMigrationStatePtrType) ToBackupPolicyMigrationStatePtrOutput() BackupPolicyMigrationStatePtrOutput {
	return i.ToBackupPolicyMigrationStatePtrOutputWithContext(context.Background())
}

func (i *backupPolicyMigrationStatePtrType) ToBackupPolicyMigrationStatePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMigrationStatePtrOutput)
}

type BackupPolicyMigrationStateOutput struct{ *pulumi.OutputState }

func (BackupPolicyMigrationStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyMigrationState)(nil)).Elem()
}

func (o BackupPolicyMigrationStateOutput) ToBackupPolicyMigrationStateOutput() BackupPolicyMigrationStateOutput {
	return o
}

func (o BackupPolicyMigrationStateOutput) ToBackupPolicyMigrationStateOutputWithContext(ctx context.Context) BackupPolicyMigrationStateOutput {
	return o
}

func (o BackupPolicyMigrationStateOutput) ToBackupPolicyMigrationStatePtrOutput() BackupPolicyMigrationStatePtrOutput {
	return o.ToBackupPolicyMigrationStatePtrOutputWithContext(context.Background())
}

func (o BackupPolicyMigrationStateOutput) ToBackupPolicyMigrationStatePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupPolicyMigrationState) *BackupPolicyMigrationState {
		return &v
	}).(BackupPolicyMigrationStatePtrOutput)
}

func (o BackupPolicyMigrationStateOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyMigrationState) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyMigrationState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStateOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyMigrationState) *string { return v.TargetType }).(pulumi.StringPtrOutput)
}

type BackupPolicyMigrationStatePtrOutput struct{ *pulumi.OutputState }

func (BackupPolicyMigrationStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyMigrationState)(nil)).Elem()
}

func (o BackupPolicyMigrationStatePtrOutput) ToBackupPolicyMigrationStatePtrOutput() BackupPolicyMigrationStatePtrOutput {
	return o
}

func (o BackupPolicyMigrationStatePtrOutput) ToBackupPolicyMigrationStatePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStatePtrOutput {
	return o
}

func (o BackupPolicyMigrationStatePtrOutput) Elem() BackupPolicyMigrationStateOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationState) BackupPolicyMigrationState {
		if v != nil {
			return *v
		}
		var ret BackupPolicyMigrationState
		return ret
	}).(BackupPolicyMigrationStateOutput)
}

func (o BackupPolicyMigrationStatePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationState) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStatePtrOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationState) *string {
		if v == nil {
			return nil
		}
		return v.TargetType
	}).(pulumi.StringPtrOutput)
}

type BackupPolicyMigrationStateResponse struct {
	StartTime  *string `pulumi:"startTime"`
	Status     *string `pulumi:"status"`
	TargetType *string `pulumi:"targetType"`
}





type BackupPolicyMigrationStateResponseInput interface {
	pulumi.Input

	ToBackupPolicyMigrationStateResponseOutput() BackupPolicyMigrationStateResponseOutput
	ToBackupPolicyMigrationStateResponseOutputWithContext(context.Context) BackupPolicyMigrationStateResponseOutput
}

type BackupPolicyMigrationStateResponseArgs struct {
	StartTime  pulumi.StringPtrInput `pulumi:"startTime"`
	Status     pulumi.StringPtrInput `pulumi:"status"`
	TargetType pulumi.StringPtrInput `pulumi:"targetType"`
}

func (BackupPolicyMigrationStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyMigrationStateResponse)(nil)).Elem()
}

func (i BackupPolicyMigrationStateResponseArgs) ToBackupPolicyMigrationStateResponseOutput() BackupPolicyMigrationStateResponseOutput {
	return i.ToBackupPolicyMigrationStateResponseOutputWithContext(context.Background())
}

func (i BackupPolicyMigrationStateResponseArgs) ToBackupPolicyMigrationStateResponseOutputWithContext(ctx context.Context) BackupPolicyMigrationStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMigrationStateResponseOutput)
}

func (i BackupPolicyMigrationStateResponseArgs) ToBackupPolicyMigrationStateResponsePtrOutput() BackupPolicyMigrationStateResponsePtrOutput {
	return i.ToBackupPolicyMigrationStateResponsePtrOutputWithContext(context.Background())
}

func (i BackupPolicyMigrationStateResponseArgs) ToBackupPolicyMigrationStateResponsePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMigrationStateResponseOutput).ToBackupPolicyMigrationStateResponsePtrOutputWithContext(ctx)
}









type BackupPolicyMigrationStateResponsePtrInput interface {
	pulumi.Input

	ToBackupPolicyMigrationStateResponsePtrOutput() BackupPolicyMigrationStateResponsePtrOutput
	ToBackupPolicyMigrationStateResponsePtrOutputWithContext(context.Context) BackupPolicyMigrationStateResponsePtrOutput
}

type backupPolicyMigrationStateResponsePtrType BackupPolicyMigrationStateResponseArgs

func BackupPolicyMigrationStateResponsePtr(v *BackupPolicyMigrationStateResponseArgs) BackupPolicyMigrationStateResponsePtrInput {
	return (*backupPolicyMigrationStateResponsePtrType)(v)
}

func (*backupPolicyMigrationStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyMigrationStateResponse)(nil)).Elem()
}

func (i *backupPolicyMigrationStateResponsePtrType) ToBackupPolicyMigrationStateResponsePtrOutput() BackupPolicyMigrationStateResponsePtrOutput {
	return i.ToBackupPolicyMigrationStateResponsePtrOutputWithContext(context.Background())
}

func (i *backupPolicyMigrationStateResponsePtrType) ToBackupPolicyMigrationStateResponsePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupPolicyMigrationStateResponsePtrOutput)
}

type BackupPolicyMigrationStateResponseOutput struct{ *pulumi.OutputState }

func (BackupPolicyMigrationStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupPolicyMigrationStateResponse)(nil)).Elem()
}

func (o BackupPolicyMigrationStateResponseOutput) ToBackupPolicyMigrationStateResponseOutput() BackupPolicyMigrationStateResponseOutput {
	return o
}

func (o BackupPolicyMigrationStateResponseOutput) ToBackupPolicyMigrationStateResponseOutputWithContext(ctx context.Context) BackupPolicyMigrationStateResponseOutput {
	return o
}

func (o BackupPolicyMigrationStateResponseOutput) ToBackupPolicyMigrationStateResponsePtrOutput() BackupPolicyMigrationStateResponsePtrOutput {
	return o.ToBackupPolicyMigrationStateResponsePtrOutputWithContext(context.Background())
}

func (o BackupPolicyMigrationStateResponseOutput) ToBackupPolicyMigrationStateResponsePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupPolicyMigrationStateResponse) *BackupPolicyMigrationStateResponse {
		return &v
	}).(BackupPolicyMigrationStateResponsePtrOutput)
}

func (o BackupPolicyMigrationStateResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyMigrationStateResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyMigrationStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStateResponseOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupPolicyMigrationStateResponse) *string { return v.TargetType }).(pulumi.StringPtrOutput)
}

type BackupPolicyMigrationStateResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupPolicyMigrationStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupPolicyMigrationStateResponse)(nil)).Elem()
}

func (o BackupPolicyMigrationStateResponsePtrOutput) ToBackupPolicyMigrationStateResponsePtrOutput() BackupPolicyMigrationStateResponsePtrOutput {
	return o
}

func (o BackupPolicyMigrationStateResponsePtrOutput) ToBackupPolicyMigrationStateResponsePtrOutputWithContext(ctx context.Context) BackupPolicyMigrationStateResponsePtrOutput {
	return o
}

func (o BackupPolicyMigrationStateResponsePtrOutput) Elem() BackupPolicyMigrationStateResponseOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationStateResponse) BackupPolicyMigrationStateResponse {
		if v != nil {
			return *v
		}
		var ret BackupPolicyMigrationStateResponse
		return ret
	}).(BackupPolicyMigrationStateResponseOutput)
}

func (o BackupPolicyMigrationStateResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

func (o BackupPolicyMigrationStateResponsePtrOutput) TargetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupPolicyMigrationStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.TargetType
	}).(pulumi.StringPtrOutput)
}

type Capability struct {
	Name *string `pulumi:"name"`
}





type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(context.Context) CapabilityOutput
}

type CapabilityArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CapabilityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (i CapabilityArgs) ToCapabilityOutput() CapabilityOutput {
	return i.ToCapabilityOutputWithContext(context.Background())
}

func (i CapabilityArgs) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityOutput)
}





type CapabilityArrayInput interface {
	pulumi.Input

	ToCapabilityArrayOutput() CapabilityArrayOutput
	ToCapabilityArrayOutputWithContext(context.Context) CapabilityArrayOutput
}

type CapabilityArray []CapabilityInput

func (CapabilityArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (i CapabilityArray) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return i.ToCapabilityArrayOutputWithContext(context.Background())
}

func (i CapabilityArray) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityArrayOutput)
}

type CapabilityOutput struct{ *pulumi.OutputState }

func (CapabilityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capability)(nil)).Elem()
}

func (o CapabilityOutput) ToCapabilityOutput() CapabilityOutput {
	return o
}

func (o CapabilityOutput) ToCapabilityOutputWithContext(ctx context.Context) CapabilityOutput {
	return o
}

func (o CapabilityOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CapabilityArrayOutput struct{ *pulumi.OutputState }

func (CapabilityArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Capability)(nil)).Elem()
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutput() CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) ToCapabilityArrayOutputWithContext(ctx context.Context) CapabilityArrayOutput {
	return o
}

func (o CapabilityArrayOutput) Index(i pulumi.IntInput) CapabilityOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Capability {
		return vs[0].([]Capability)[vs[1].(int)]
	}).(CapabilityOutput)
}

type CapabilityResponse struct {
	Name *string `pulumi:"name"`
}





type CapabilityResponseInput interface {
	pulumi.Input

	ToCapabilityResponseOutput() CapabilityResponseOutput
	ToCapabilityResponseOutputWithContext(context.Context) CapabilityResponseOutput
}

type CapabilityResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CapabilityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityResponse)(nil)).Elem()
}

func (i CapabilityResponseArgs) ToCapabilityResponseOutput() CapabilityResponseOutput {
	return i.ToCapabilityResponseOutputWithContext(context.Background())
}

func (i CapabilityResponseArgs) ToCapabilityResponseOutputWithContext(ctx context.Context) CapabilityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityResponseOutput)
}





type CapabilityResponseArrayInput interface {
	pulumi.Input

	ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput
	ToCapabilityResponseArrayOutputWithContext(context.Context) CapabilityResponseArrayOutput
}

type CapabilityResponseArray []CapabilityResponseInput

func (CapabilityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapabilityResponse)(nil)).Elem()
}

func (i CapabilityResponseArray) ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput {
	return i.ToCapabilityResponseArrayOutputWithContext(context.Background())
}

func (i CapabilityResponseArray) ToCapabilityResponseArrayOutputWithContext(ctx context.Context) CapabilityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapabilityResponseArrayOutput)
}

type CapabilityResponseOutput struct{ *pulumi.OutputState }

func (CapabilityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutput() CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) ToCapabilityResponseOutputWithContext(ctx context.Context) CapabilityResponseOutput {
	return o
}

func (o CapabilityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CapabilityResponseArrayOutput struct{ *pulumi.OutputState }

func (CapabilityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CapabilityResponse)(nil)).Elem()
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutput() CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) ToCapabilityResponseArrayOutputWithContext(ctx context.Context) CapabilityResponseArrayOutput {
	return o
}

func (o CapabilityResponseArrayOutput) Index(i pulumi.IntInput) CapabilityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CapabilityResponse {
		return vs[0].([]CapabilityResponse)[vs[1].(int)]
	}).(CapabilityResponseOutput)
}

type Capacity struct {
	TotalThroughputLimit *int `pulumi:"totalThroughputLimit"`
}





type CapacityInput interface {
	pulumi.Input

	ToCapacityOutput() CapacityOutput
	ToCapacityOutputWithContext(context.Context) CapacityOutput
}

type CapacityArgs struct {
	TotalThroughputLimit pulumi.IntPtrInput `pulumi:"totalThroughputLimit"`
}

func (CapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Capacity)(nil)).Elem()
}

func (i CapacityArgs) ToCapacityOutput() CapacityOutput {
	return i.ToCapacityOutputWithContext(context.Background())
}

func (i CapacityArgs) ToCapacityOutputWithContext(ctx context.Context) CapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityOutput)
}

func (i CapacityArgs) ToCapacityPtrOutput() CapacityPtrOutput {
	return i.ToCapacityPtrOutputWithContext(context.Background())
}

func (i CapacityArgs) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityOutput).ToCapacityPtrOutputWithContext(ctx)
}









type CapacityPtrInput interface {
	pulumi.Input

	ToCapacityPtrOutput() CapacityPtrOutput
	ToCapacityPtrOutputWithContext(context.Context) CapacityPtrOutput
}

type capacityPtrType CapacityArgs

func CapacityPtr(v *CapacityArgs) CapacityPtrInput {
	return (*capacityPtrType)(v)
}

func (*capacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Capacity)(nil)).Elem()
}

func (i *capacityPtrType) ToCapacityPtrOutput() CapacityPtrOutput {
	return i.ToCapacityPtrOutputWithContext(context.Background())
}

func (i *capacityPtrType) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityPtrOutput)
}

type CapacityOutput struct{ *pulumi.OutputState }

func (CapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Capacity)(nil)).Elem()
}

func (o CapacityOutput) ToCapacityOutput() CapacityOutput {
	return o
}

func (o CapacityOutput) ToCapacityOutputWithContext(ctx context.Context) CapacityOutput {
	return o
}

func (o CapacityOutput) ToCapacityPtrOutput() CapacityPtrOutput {
	return o.ToCapacityPtrOutputWithContext(context.Background())
}

func (o CapacityOutput) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Capacity) *Capacity {
		return &v
	}).(CapacityPtrOutput)
}

func (o CapacityOutput) TotalThroughputLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Capacity) *int { return v.TotalThroughputLimit }).(pulumi.IntPtrOutput)
}

type CapacityPtrOutput struct{ *pulumi.OutputState }

func (CapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Capacity)(nil)).Elem()
}

func (o CapacityPtrOutput) ToCapacityPtrOutput() CapacityPtrOutput {
	return o
}

func (o CapacityPtrOutput) ToCapacityPtrOutputWithContext(ctx context.Context) CapacityPtrOutput {
	return o
}

func (o CapacityPtrOutput) Elem() CapacityOutput {
	return o.ApplyT(func(v *Capacity) Capacity {
		if v != nil {
			return *v
		}
		var ret Capacity
		return ret
	}).(CapacityOutput)
}

func (o CapacityPtrOutput) TotalThroughputLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Capacity) *int {
		if v == nil {
			return nil
		}
		return v.TotalThroughputLimit
	}).(pulumi.IntPtrOutput)
}

type CapacityResponse struct {
	TotalThroughputLimit *int `pulumi:"totalThroughputLimit"`
}





type CapacityResponseInput interface {
	pulumi.Input

	ToCapacityResponseOutput() CapacityResponseOutput
	ToCapacityResponseOutputWithContext(context.Context) CapacityResponseOutput
}

type CapacityResponseArgs struct {
	TotalThroughputLimit pulumi.IntPtrInput `pulumi:"totalThroughputLimit"`
}

func (CapacityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityResponse)(nil)).Elem()
}

func (i CapacityResponseArgs) ToCapacityResponseOutput() CapacityResponseOutput {
	return i.ToCapacityResponseOutputWithContext(context.Background())
}

func (i CapacityResponseArgs) ToCapacityResponseOutputWithContext(ctx context.Context) CapacityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityResponseOutput)
}

func (i CapacityResponseArgs) ToCapacityResponsePtrOutput() CapacityResponsePtrOutput {
	return i.ToCapacityResponsePtrOutputWithContext(context.Background())
}

func (i CapacityResponseArgs) ToCapacityResponsePtrOutputWithContext(ctx context.Context) CapacityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityResponseOutput).ToCapacityResponsePtrOutputWithContext(ctx)
}









type CapacityResponsePtrInput interface {
	pulumi.Input

	ToCapacityResponsePtrOutput() CapacityResponsePtrOutput
	ToCapacityResponsePtrOutputWithContext(context.Context) CapacityResponsePtrOutput
}

type capacityResponsePtrType CapacityResponseArgs

func CapacityResponsePtr(v *CapacityResponseArgs) CapacityResponsePtrInput {
	return (*capacityResponsePtrType)(v)
}

func (*capacityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityResponse)(nil)).Elem()
}

func (i *capacityResponsePtrType) ToCapacityResponsePtrOutput() CapacityResponsePtrOutput {
	return i.ToCapacityResponsePtrOutputWithContext(context.Background())
}

func (i *capacityResponsePtrType) ToCapacityResponsePtrOutputWithContext(ctx context.Context) CapacityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CapacityResponsePtrOutput)
}

type CapacityResponseOutput struct{ *pulumi.OutputState }

func (CapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CapacityResponse)(nil)).Elem()
}

func (o CapacityResponseOutput) ToCapacityResponseOutput() CapacityResponseOutput {
	return o
}

func (o CapacityResponseOutput) ToCapacityResponseOutputWithContext(ctx context.Context) CapacityResponseOutput {
	return o
}

func (o CapacityResponseOutput) ToCapacityResponsePtrOutput() CapacityResponsePtrOutput {
	return o.ToCapacityResponsePtrOutputWithContext(context.Background())
}

func (o CapacityResponseOutput) ToCapacityResponsePtrOutputWithContext(ctx context.Context) CapacityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CapacityResponse) *CapacityResponse {
		return &v
	}).(CapacityResponsePtrOutput)
}

func (o CapacityResponseOutput) TotalThroughputLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CapacityResponse) *int { return v.TotalThroughputLimit }).(pulumi.IntPtrOutput)
}

type CapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (CapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CapacityResponse)(nil)).Elem()
}

func (o CapacityResponsePtrOutput) ToCapacityResponsePtrOutput() CapacityResponsePtrOutput {
	return o
}

func (o CapacityResponsePtrOutput) ToCapacityResponsePtrOutputWithContext(ctx context.Context) CapacityResponsePtrOutput {
	return o
}

func (o CapacityResponsePtrOutput) Elem() CapacityResponseOutput {
	return o.ApplyT(func(v *CapacityResponse) CapacityResponse {
		if v != nil {
			return *v
		}
		var ret CapacityResponse
		return ret
	}).(CapacityResponseOutput)
}

func (o CapacityResponsePtrOutput) TotalThroughputLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.TotalThroughputLimit
	}).(pulumi.IntPtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type CassandraKeyspaceGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToCassandraKeyspaceGetPropertiesResponseOptionsOutput() CassandraKeyspaceGetPropertiesResponseOptionsOutput
	ToCassandraKeyspaceGetPropertiesResponseOptionsOutputWithContext(context.Context) CassandraKeyspaceGetPropertiesResponseOptionsOutput
}

type CassandraKeyspaceGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (CassandraKeyspaceGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceGetPropertiesResponseOptions)(nil)).Elem()
}

func (i CassandraKeyspaceGetPropertiesResponseOptionsArgs) ToCassandraKeyspaceGetPropertiesResponseOptionsOutput() CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return i.ToCassandraKeyspaceGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i CassandraKeyspaceGetPropertiesResponseOptionsArgs) ToCassandraKeyspaceGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceGetPropertiesResponseOptionsOutput)
}

func (i CassandraKeyspaceGetPropertiesResponseOptionsArgs) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutput() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return i.ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i CassandraKeyspaceGetPropertiesResponseOptionsArgs) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceGetPropertiesResponseOptionsOutput).ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type CassandraKeyspaceGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutput() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput
	ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput
}

type cassandraKeyspaceGetPropertiesResponseOptionsPtrType CassandraKeyspaceGetPropertiesResponseOptionsArgs

func CassandraKeyspaceGetPropertiesResponseOptionsPtr(v *CassandraKeyspaceGetPropertiesResponseOptionsArgs) CassandraKeyspaceGetPropertiesResponseOptionsPtrInput {
	return (*cassandraKeyspaceGetPropertiesResponseOptionsPtrType)(v)
}

func (*cassandraKeyspaceGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *cassandraKeyspaceGetPropertiesResponseOptionsPtrType) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutput() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return i.ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *cassandraKeyspaceGetPropertiesResponseOptionsPtrType) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsOutput() CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutput() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o.ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraKeyspaceGetPropertiesResponseOptions) *CassandraKeyspaceGetPropertiesResponseOptions {
		return &v
	}).(CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutput() CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) ToCassandraKeyspaceGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) Elem() CassandraKeyspaceGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseOptions) CassandraKeyspaceGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret CassandraKeyspaceGetPropertiesResponseOptions
		return ret
	}).(CassandraKeyspaceGetPropertiesResponseOptionsOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}





type CassandraKeyspaceGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToCassandraKeyspaceGetPropertiesResponseResourceOutput() CassandraKeyspaceGetPropertiesResponseResourceOutput
	ToCassandraKeyspaceGetPropertiesResponseResourceOutputWithContext(context.Context) CassandraKeyspaceGetPropertiesResponseResourceOutput
}

type CassandraKeyspaceGetPropertiesResponseResourceArgs struct {
	Etag pulumi.StringInput  `pulumi:"etag"`
	Id   pulumi.StringInput  `pulumi:"id"`
	Rid  pulumi.StringInput  `pulumi:"rid"`
	Ts   pulumi.Float64Input `pulumi:"ts"`
}

func (CassandraKeyspaceGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceGetPropertiesResponseResource)(nil)).Elem()
}

func (i CassandraKeyspaceGetPropertiesResponseResourceArgs) ToCassandraKeyspaceGetPropertiesResponseResourceOutput() CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return i.ToCassandraKeyspaceGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i CassandraKeyspaceGetPropertiesResponseResourceArgs) ToCassandraKeyspaceGetPropertiesResponseResourceOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceGetPropertiesResponseResourceOutput)
}

func (i CassandraKeyspaceGetPropertiesResponseResourceArgs) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutput() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return i.ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i CassandraKeyspaceGetPropertiesResponseResourceArgs) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceGetPropertiesResponseResourceOutput).ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type CassandraKeyspaceGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutput() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput
	ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(context.Context) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput
}

type cassandraKeyspaceGetPropertiesResponseResourcePtrType CassandraKeyspaceGetPropertiesResponseResourceArgs

func CassandraKeyspaceGetPropertiesResponseResourcePtr(v *CassandraKeyspaceGetPropertiesResponseResourceArgs) CassandraKeyspaceGetPropertiesResponseResourcePtrInput {
	return (*cassandraKeyspaceGetPropertiesResponseResourcePtrType)(v)
}

func (*cassandraKeyspaceGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceGetPropertiesResponseResource)(nil)).Elem()
}

func (i *cassandraKeyspaceGetPropertiesResponseResourcePtrType) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutput() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return i.ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *cassandraKeyspaceGetPropertiesResponseResourcePtrType) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceGetPropertiesResponseResourcePtrOutput)
}

type CassandraKeyspaceGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) ToCassandraKeyspaceGetPropertiesResponseResourceOutput() CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) ToCassandraKeyspaceGetPropertiesResponseResourceOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutput() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o.ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraKeyspaceGetPropertiesResponseResource) *CassandraKeyspaceGetPropertiesResponseResource {
		return &v
	}).(CassandraKeyspaceGetPropertiesResponseResourcePtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v CassandraKeyspaceGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type CassandraKeyspaceGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutput() CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) ToCassandraKeyspaceGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Elem() CassandraKeyspaceGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) CassandraKeyspaceGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret CassandraKeyspaceGetPropertiesResponseResource
		return ret
	}).(CassandraKeyspaceGetPropertiesResponseResourceOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o CassandraKeyspaceGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type CassandraKeyspaceResource struct {
	Id string `pulumi:"id"`
}





type CassandraKeyspaceResourceInput interface {
	pulumi.Input

	ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput
	ToCassandraKeyspaceResourceOutputWithContext(context.Context) CassandraKeyspaceResourceOutput
}

type CassandraKeyspaceResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (CassandraKeyspaceResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceResource)(nil)).Elem()
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput {
	return i.ToCassandraKeyspaceResourceOutputWithContext(context.Background())
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourceOutputWithContext(ctx context.Context) CassandraKeyspaceResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceResourceOutput)
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourcePtrOutput() CassandraKeyspaceResourcePtrOutput {
	return i.ToCassandraKeyspaceResourcePtrOutputWithContext(context.Background())
}

func (i CassandraKeyspaceResourceArgs) ToCassandraKeyspaceResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceResourceOutput).ToCassandraKeyspaceResourcePtrOutputWithContext(ctx)
}









type CassandraKeyspaceResourcePtrInput interface {
	pulumi.Input

	ToCassandraKeyspaceResourcePtrOutput() CassandraKeyspaceResourcePtrOutput
	ToCassandraKeyspaceResourcePtrOutputWithContext(context.Context) CassandraKeyspaceResourcePtrOutput
}

type cassandraKeyspaceResourcePtrType CassandraKeyspaceResourceArgs

func CassandraKeyspaceResourcePtr(v *CassandraKeyspaceResourceArgs) CassandraKeyspaceResourcePtrInput {
	return (*cassandraKeyspaceResourcePtrType)(v)
}

func (*cassandraKeyspaceResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceResource)(nil)).Elem()
}

func (i *cassandraKeyspaceResourcePtrType) ToCassandraKeyspaceResourcePtrOutput() CassandraKeyspaceResourcePtrOutput {
	return i.ToCassandraKeyspaceResourcePtrOutputWithContext(context.Background())
}

func (i *cassandraKeyspaceResourcePtrType) ToCassandraKeyspaceResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraKeyspaceResourcePtrOutput)
}

type CassandraKeyspaceResourceOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraKeyspaceResource)(nil)).Elem()
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourceOutput() CassandraKeyspaceResourceOutput {
	return o
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourceOutputWithContext(ctx context.Context) CassandraKeyspaceResourceOutput {
	return o
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourcePtrOutput() CassandraKeyspaceResourcePtrOutput {
	return o.ToCassandraKeyspaceResourcePtrOutputWithContext(context.Background())
}

func (o CassandraKeyspaceResourceOutput) ToCassandraKeyspaceResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraKeyspaceResource) *CassandraKeyspaceResource {
		return &v
	}).(CassandraKeyspaceResourcePtrOutput)
}

func (o CassandraKeyspaceResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraKeyspaceResource) string { return v.Id }).(pulumi.StringOutput)
}

type CassandraKeyspaceResourcePtrOutput struct{ *pulumi.OutputState }

func (CassandraKeyspaceResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraKeyspaceResource)(nil)).Elem()
}

func (o CassandraKeyspaceResourcePtrOutput) ToCassandraKeyspaceResourcePtrOutput() CassandraKeyspaceResourcePtrOutput {
	return o
}

func (o CassandraKeyspaceResourcePtrOutput) ToCassandraKeyspaceResourcePtrOutputWithContext(ctx context.Context) CassandraKeyspaceResourcePtrOutput {
	return o
}

func (o CassandraKeyspaceResourcePtrOutput) Elem() CassandraKeyspaceResourceOutput {
	return o.ApplyT(func(v *CassandraKeyspaceResource) CassandraKeyspaceResource {
		if v != nil {
			return *v
		}
		var ret CassandraKeyspaceResource
		return ret
	}).(CassandraKeyspaceResourceOutput)
}

func (o CassandraKeyspaceResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraKeyspaceResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type CassandraPartitionKey struct {
	Name *string `pulumi:"name"`
}





type CassandraPartitionKeyInput interface {
	pulumi.Input

	ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput
	ToCassandraPartitionKeyOutputWithContext(context.Context) CassandraPartitionKeyOutput
}

type CassandraPartitionKeyArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CassandraPartitionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKey)(nil)).Elem()
}

func (i CassandraPartitionKeyArgs) ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput {
	return i.ToCassandraPartitionKeyOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyArgs) ToCassandraPartitionKeyOutputWithContext(ctx context.Context) CassandraPartitionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyOutput)
}





type CassandraPartitionKeyArrayInput interface {
	pulumi.Input

	ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput
	ToCassandraPartitionKeyArrayOutputWithContext(context.Context) CassandraPartitionKeyArrayOutput
}

type CassandraPartitionKeyArray []CassandraPartitionKeyInput

func (CassandraPartitionKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKey)(nil)).Elem()
}

func (i CassandraPartitionKeyArray) ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput {
	return i.ToCassandraPartitionKeyArrayOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyArray) ToCassandraPartitionKeyArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyArrayOutput)
}

type CassandraPartitionKeyOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKey)(nil)).Elem()
}

func (o CassandraPartitionKeyOutput) ToCassandraPartitionKeyOutput() CassandraPartitionKeyOutput {
	return o
}

func (o CassandraPartitionKeyOutput) ToCassandraPartitionKeyOutputWithContext(ctx context.Context) CassandraPartitionKeyOutput {
	return o
}

func (o CassandraPartitionKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CassandraPartitionKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CassandraPartitionKeyArrayOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKey)(nil)).Elem()
}

func (o CassandraPartitionKeyArrayOutput) ToCassandraPartitionKeyArrayOutput() CassandraPartitionKeyArrayOutput {
	return o
}

func (o CassandraPartitionKeyArrayOutput) ToCassandraPartitionKeyArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyArrayOutput {
	return o
}

func (o CassandraPartitionKeyArrayOutput) Index(i pulumi.IntInput) CassandraPartitionKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraPartitionKey {
		return vs[0].([]CassandraPartitionKey)[vs[1].(int)]
	}).(CassandraPartitionKeyOutput)
}

type CassandraPartitionKeyResponse struct {
	Name *string `pulumi:"name"`
}





type CassandraPartitionKeyResponseInput interface {
	pulumi.Input

	ToCassandraPartitionKeyResponseOutput() CassandraPartitionKeyResponseOutput
	ToCassandraPartitionKeyResponseOutputWithContext(context.Context) CassandraPartitionKeyResponseOutput
}

type CassandraPartitionKeyResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (CassandraPartitionKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKeyResponse)(nil)).Elem()
}

func (i CassandraPartitionKeyResponseArgs) ToCassandraPartitionKeyResponseOutput() CassandraPartitionKeyResponseOutput {
	return i.ToCassandraPartitionKeyResponseOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyResponseArgs) ToCassandraPartitionKeyResponseOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyResponseOutput)
}





type CassandraPartitionKeyResponseArrayInput interface {
	pulumi.Input

	ToCassandraPartitionKeyResponseArrayOutput() CassandraPartitionKeyResponseArrayOutput
	ToCassandraPartitionKeyResponseArrayOutputWithContext(context.Context) CassandraPartitionKeyResponseArrayOutput
}

type CassandraPartitionKeyResponseArray []CassandraPartitionKeyResponseInput

func (CassandraPartitionKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKeyResponse)(nil)).Elem()
}

func (i CassandraPartitionKeyResponseArray) ToCassandraPartitionKeyResponseArrayOutput() CassandraPartitionKeyResponseArrayOutput {
	return i.ToCassandraPartitionKeyResponseArrayOutputWithContext(context.Background())
}

func (i CassandraPartitionKeyResponseArray) ToCassandraPartitionKeyResponseArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraPartitionKeyResponseOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraPartitionKeyResponse)(nil)).Elem()
}

func (o CassandraPartitionKeyResponseOutput) ToCassandraPartitionKeyResponseOutput() CassandraPartitionKeyResponseOutput {
	return o
}

func (o CassandraPartitionKeyResponseOutput) ToCassandraPartitionKeyResponseOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseOutput {
	return o
}

func (o CassandraPartitionKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CassandraPartitionKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type CassandraPartitionKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (CassandraPartitionKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CassandraPartitionKeyResponse)(nil)).Elem()
}

func (o CassandraPartitionKeyResponseArrayOutput) ToCassandraPartitionKeyResponseArrayOutput() CassandraPartitionKeyResponseArrayOutput {
	return o
}

func (o CassandraPartitionKeyResponseArrayOutput) ToCassandraPartitionKeyResponseArrayOutputWithContext(ctx context.Context) CassandraPartitionKeyResponseArrayOutput {
	return o
}

func (o CassandraPartitionKeyResponseArrayOutput) Index(i pulumi.IntInput) CassandraPartitionKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CassandraPartitionKeyResponse {
		return vs[0].([]CassandraPartitionKeyResponse)[vs[1].(int)]
	}).(CassandraPartitionKeyResponseOutput)
}

type CassandraSchema struct {
	ClusterKeys   []ClusterKey            `pulumi:"clusterKeys"`
	Columns       []Column                `pulumi:"columns"`
	PartitionKeys []CassandraPartitionKey `pulumi:"partitionKeys"`
}





type CassandraSchemaInput interface {
	pulumi.Input

	ToCassandraSchemaOutput() CassandraSchemaOutput
	ToCassandraSchemaOutputWithContext(context.Context) CassandraSchemaOutput
}

type CassandraSchemaArgs struct {
	ClusterKeys   ClusterKeyArrayInput            `pulumi:"clusterKeys"`
	Columns       ColumnArrayInput                `pulumi:"columns"`
	PartitionKeys CassandraPartitionKeyArrayInput `pulumi:"partitionKeys"`
}

func (CassandraSchemaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchema)(nil)).Elem()
}

func (i CassandraSchemaArgs) ToCassandraSchemaOutput() CassandraSchemaOutput {
	return i.ToCassandraSchemaOutputWithContext(context.Background())
}

func (i CassandraSchemaArgs) ToCassandraSchemaOutputWithContext(ctx context.Context) CassandraSchemaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaOutput)
}

func (i CassandraSchemaArgs) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return i.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (i CassandraSchemaArgs) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaOutput).ToCassandraSchemaPtrOutputWithContext(ctx)
}









type CassandraSchemaPtrInput interface {
	pulumi.Input

	ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput
	ToCassandraSchemaPtrOutputWithContext(context.Context) CassandraSchemaPtrOutput
}

type cassandraSchemaPtrType CassandraSchemaArgs

func CassandraSchemaPtr(v *CassandraSchemaArgs) CassandraSchemaPtrInput {
	return (*cassandraSchemaPtrType)(v)
}

func (*cassandraSchemaPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchema)(nil)).Elem()
}

func (i *cassandraSchemaPtrType) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return i.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (i *cassandraSchemaPtrType) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaPtrOutput)
}

type CassandraSchemaOutput struct{ *pulumi.OutputState }

func (CassandraSchemaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchema)(nil)).Elem()
}

func (o CassandraSchemaOutput) ToCassandraSchemaOutput() CassandraSchemaOutput {
	return o
}

func (o CassandraSchemaOutput) ToCassandraSchemaOutputWithContext(ctx context.Context) CassandraSchemaOutput {
	return o
}

func (o CassandraSchemaOutput) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return o.ToCassandraSchemaPtrOutputWithContext(context.Background())
}

func (o CassandraSchemaOutput) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraSchema) *CassandraSchema {
		return &v
	}).(CassandraSchemaPtrOutput)
}

func (o CassandraSchemaOutput) ClusterKeys() ClusterKeyArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []ClusterKey { return v.ClusterKeys }).(ClusterKeyArrayOutput)
}

func (o CassandraSchemaOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []Column { return v.Columns }).(ColumnArrayOutput)
}

func (o CassandraSchemaOutput) PartitionKeys() CassandraPartitionKeyArrayOutput {
	return o.ApplyT(func(v CassandraSchema) []CassandraPartitionKey { return v.PartitionKeys }).(CassandraPartitionKeyArrayOutput)
}

type CassandraSchemaPtrOutput struct{ *pulumi.OutputState }

func (CassandraSchemaPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchema)(nil)).Elem()
}

func (o CassandraSchemaPtrOutput) ToCassandraSchemaPtrOutput() CassandraSchemaPtrOutput {
	return o
}

func (o CassandraSchemaPtrOutput) ToCassandraSchemaPtrOutputWithContext(ctx context.Context) CassandraSchemaPtrOutput {
	return o
}

func (o CassandraSchemaPtrOutput) Elem() CassandraSchemaOutput {
	return o.ApplyT(func(v *CassandraSchema) CassandraSchema {
		if v != nil {
			return *v
		}
		var ret CassandraSchema
		return ret
	}).(CassandraSchemaOutput)
}

func (o CassandraSchemaPtrOutput) ClusterKeys() ClusterKeyArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []ClusterKey {
		if v == nil {
			return nil
		}
		return v.ClusterKeys
	}).(ClusterKeyArrayOutput)
}

func (o CassandraSchemaPtrOutput) Columns() ColumnArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []Column {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnArrayOutput)
}

func (o CassandraSchemaPtrOutput) PartitionKeys() CassandraPartitionKeyArrayOutput {
	return o.ApplyT(func(v *CassandraSchema) []CassandraPartitionKey {
		if v == nil {
			return nil
		}
		return v.PartitionKeys
	}).(CassandraPartitionKeyArrayOutput)
}

type CassandraSchemaResponse struct {
	ClusterKeys   []ClusterKeyResponse            `pulumi:"clusterKeys"`
	Columns       []ColumnResponse                `pulumi:"columns"`
	PartitionKeys []CassandraPartitionKeyResponse `pulumi:"partitionKeys"`
}





type CassandraSchemaResponseInput interface {
	pulumi.Input

	ToCassandraSchemaResponseOutput() CassandraSchemaResponseOutput
	ToCassandraSchemaResponseOutputWithContext(context.Context) CassandraSchemaResponseOutput
}

type CassandraSchemaResponseArgs struct {
	ClusterKeys   ClusterKeyResponseArrayInput            `pulumi:"clusterKeys"`
	Columns       ColumnResponseArrayInput                `pulumi:"columns"`
	PartitionKeys CassandraPartitionKeyResponseArrayInput `pulumi:"partitionKeys"`
}

func (CassandraSchemaResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchemaResponse)(nil)).Elem()
}

func (i CassandraSchemaResponseArgs) ToCassandraSchemaResponseOutput() CassandraSchemaResponseOutput {
	return i.ToCassandraSchemaResponseOutputWithContext(context.Background())
}

func (i CassandraSchemaResponseArgs) ToCassandraSchemaResponseOutputWithContext(ctx context.Context) CassandraSchemaResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaResponseOutput)
}

func (i CassandraSchemaResponseArgs) ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput {
	return i.ToCassandraSchemaResponsePtrOutputWithContext(context.Background())
}

func (i CassandraSchemaResponseArgs) ToCassandraSchemaResponsePtrOutputWithContext(ctx context.Context) CassandraSchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaResponseOutput).ToCassandraSchemaResponsePtrOutputWithContext(ctx)
}









type CassandraSchemaResponsePtrInput interface {
	pulumi.Input

	ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput
	ToCassandraSchemaResponsePtrOutputWithContext(context.Context) CassandraSchemaResponsePtrOutput
}

type cassandraSchemaResponsePtrType CassandraSchemaResponseArgs

func CassandraSchemaResponsePtr(v *CassandraSchemaResponseArgs) CassandraSchemaResponsePtrInput {
	return (*cassandraSchemaResponsePtrType)(v)
}

func (*cassandraSchemaResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchemaResponse)(nil)).Elem()
}

func (i *cassandraSchemaResponsePtrType) ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput {
	return i.ToCassandraSchemaResponsePtrOutputWithContext(context.Background())
}

func (i *cassandraSchemaResponsePtrType) ToCassandraSchemaResponsePtrOutputWithContext(ctx context.Context) CassandraSchemaResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraSchemaResponsePtrOutput)
}

type CassandraSchemaResponseOutput struct{ *pulumi.OutputState }

func (CassandraSchemaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraSchemaResponse)(nil)).Elem()
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponseOutput() CassandraSchemaResponseOutput {
	return o
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponseOutputWithContext(ctx context.Context) CassandraSchemaResponseOutput {
	return o
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput {
	return o.ToCassandraSchemaResponsePtrOutputWithContext(context.Background())
}

func (o CassandraSchemaResponseOutput) ToCassandraSchemaResponsePtrOutputWithContext(ctx context.Context) CassandraSchemaResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraSchemaResponse) *CassandraSchemaResponse {
		return &v
	}).(CassandraSchemaResponsePtrOutput)
}

func (o CassandraSchemaResponseOutput) ClusterKeys() ClusterKeyResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []ClusterKeyResponse { return v.ClusterKeys }).(ClusterKeyResponseArrayOutput)
}

func (o CassandraSchemaResponseOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []ColumnResponse { return v.Columns }).(ColumnResponseArrayOutput)
}

func (o CassandraSchemaResponseOutput) PartitionKeys() CassandraPartitionKeyResponseArrayOutput {
	return o.ApplyT(func(v CassandraSchemaResponse) []CassandraPartitionKeyResponse { return v.PartitionKeys }).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraSchemaResponsePtrOutput struct{ *pulumi.OutputState }

func (CassandraSchemaResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraSchemaResponse)(nil)).Elem()
}

func (o CassandraSchemaResponsePtrOutput) ToCassandraSchemaResponsePtrOutput() CassandraSchemaResponsePtrOutput {
	return o
}

func (o CassandraSchemaResponsePtrOutput) ToCassandraSchemaResponsePtrOutputWithContext(ctx context.Context) CassandraSchemaResponsePtrOutput {
	return o
}

func (o CassandraSchemaResponsePtrOutput) Elem() CassandraSchemaResponseOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) CassandraSchemaResponse {
		if v != nil {
			return *v
		}
		var ret CassandraSchemaResponse
		return ret
	}).(CassandraSchemaResponseOutput)
}

func (o CassandraSchemaResponsePtrOutput) ClusterKeys() ClusterKeyResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []ClusterKeyResponse {
		if v == nil {
			return nil
		}
		return v.ClusterKeys
	}).(ClusterKeyResponseArrayOutput)
}

func (o CassandraSchemaResponsePtrOutput) Columns() ColumnResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []ColumnResponse {
		if v == nil {
			return nil
		}
		return v.Columns
	}).(ColumnResponseArrayOutput)
}

func (o CassandraSchemaResponsePtrOutput) PartitionKeys() CassandraPartitionKeyResponseArrayOutput {
	return o.ApplyT(func(v *CassandraSchemaResponse) []CassandraPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKeys
	}).(CassandraPartitionKeyResponseArrayOutput)
}

type CassandraTableGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type CassandraTableGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToCassandraTableGetPropertiesResponseOptionsOutput() CassandraTableGetPropertiesResponseOptionsOutput
	ToCassandraTableGetPropertiesResponseOptionsOutputWithContext(context.Context) CassandraTableGetPropertiesResponseOptionsOutput
}

type CassandraTableGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (CassandraTableGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableGetPropertiesResponseOptions)(nil)).Elem()
}

func (i CassandraTableGetPropertiesResponseOptionsArgs) ToCassandraTableGetPropertiesResponseOptionsOutput() CassandraTableGetPropertiesResponseOptionsOutput {
	return i.ToCassandraTableGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i CassandraTableGetPropertiesResponseOptionsArgs) ToCassandraTableGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableGetPropertiesResponseOptionsOutput)
}

func (i CassandraTableGetPropertiesResponseOptionsArgs) ToCassandraTableGetPropertiesResponseOptionsPtrOutput() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return i.ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i CassandraTableGetPropertiesResponseOptionsArgs) ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableGetPropertiesResponseOptionsOutput).ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type CassandraTableGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToCassandraTableGetPropertiesResponseOptionsPtrOutput() CassandraTableGetPropertiesResponseOptionsPtrOutput
	ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) CassandraTableGetPropertiesResponseOptionsPtrOutput
}

type cassandraTableGetPropertiesResponseOptionsPtrType CassandraTableGetPropertiesResponseOptionsArgs

func CassandraTableGetPropertiesResponseOptionsPtr(v *CassandraTableGetPropertiesResponseOptionsArgs) CassandraTableGetPropertiesResponseOptionsPtrInput {
	return (*cassandraTableGetPropertiesResponseOptionsPtrType)(v)
}

func (*cassandraTableGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *cassandraTableGetPropertiesResponseOptionsPtrType) ToCassandraTableGetPropertiesResponseOptionsPtrOutput() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return i.ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *cassandraTableGetPropertiesResponseOptionsPtrType) ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableGetPropertiesResponseOptionsPtrOutput)
}

type CassandraTableGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) ToCassandraTableGetPropertiesResponseOptionsOutput() CassandraTableGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) ToCassandraTableGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) ToCassandraTableGetPropertiesResponseOptionsPtrOutput() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o.ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraTableGetPropertiesResponseOptions) *CassandraTableGetPropertiesResponseOptions {
		return &v
	}).(CassandraTableGetPropertiesResponseOptionsPtrOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type CassandraTableGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) ToCassandraTableGetPropertiesResponseOptionsPtrOutput() CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) ToCassandraTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) Elem() CassandraTableGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseOptions) CassandraTableGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret CassandraTableGetPropertiesResponseOptions
		return ret
	}).(CassandraTableGetPropertiesResponseOptionsOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type CassandraTableGetPropertiesResponseResource struct {
	AnalyticalStorageTtl *int                     `pulumi:"analyticalStorageTtl"`
	DefaultTtl           *int                     `pulumi:"defaultTtl"`
	Etag                 string                   `pulumi:"etag"`
	Id                   string                   `pulumi:"id"`
	Rid                  string                   `pulumi:"rid"`
	Schema               *CassandraSchemaResponse `pulumi:"schema"`
	Ts                   float64                  `pulumi:"ts"`
}





type CassandraTableGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToCassandraTableGetPropertiesResponseResourceOutput() CassandraTableGetPropertiesResponseResourceOutput
	ToCassandraTableGetPropertiesResponseResourceOutputWithContext(context.Context) CassandraTableGetPropertiesResponseResourceOutput
}

type CassandraTableGetPropertiesResponseResourceArgs struct {
	AnalyticalStorageTtl pulumi.IntPtrInput              `pulumi:"analyticalStorageTtl"`
	DefaultTtl           pulumi.IntPtrInput              `pulumi:"defaultTtl"`
	Etag                 pulumi.StringInput              `pulumi:"etag"`
	Id                   pulumi.StringInput              `pulumi:"id"`
	Rid                  pulumi.StringInput              `pulumi:"rid"`
	Schema               CassandraSchemaResponsePtrInput `pulumi:"schema"`
	Ts                   pulumi.Float64Input             `pulumi:"ts"`
}

func (CassandraTableGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableGetPropertiesResponseResource)(nil)).Elem()
}

func (i CassandraTableGetPropertiesResponseResourceArgs) ToCassandraTableGetPropertiesResponseResourceOutput() CassandraTableGetPropertiesResponseResourceOutput {
	return i.ToCassandraTableGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i CassandraTableGetPropertiesResponseResourceArgs) ToCassandraTableGetPropertiesResponseResourceOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableGetPropertiesResponseResourceOutput)
}

func (i CassandraTableGetPropertiesResponseResourceArgs) ToCassandraTableGetPropertiesResponseResourcePtrOutput() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return i.ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i CassandraTableGetPropertiesResponseResourceArgs) ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableGetPropertiesResponseResourceOutput).ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type CassandraTableGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToCassandraTableGetPropertiesResponseResourcePtrOutput() CassandraTableGetPropertiesResponseResourcePtrOutput
	ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(context.Context) CassandraTableGetPropertiesResponseResourcePtrOutput
}

type cassandraTableGetPropertiesResponseResourcePtrType CassandraTableGetPropertiesResponseResourceArgs

func CassandraTableGetPropertiesResponseResourcePtr(v *CassandraTableGetPropertiesResponseResourceArgs) CassandraTableGetPropertiesResponseResourcePtrInput {
	return (*cassandraTableGetPropertiesResponseResourcePtrType)(v)
}

func (*cassandraTableGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableGetPropertiesResponseResource)(nil)).Elem()
}

func (i *cassandraTableGetPropertiesResponseResourcePtrType) ToCassandraTableGetPropertiesResponseResourcePtrOutput() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return i.ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *cassandraTableGetPropertiesResponseResourcePtrType) ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableGetPropertiesResponseResourcePtrOutput)
}

type CassandraTableGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseResourceOutput) ToCassandraTableGetPropertiesResponseResourceOutput() CassandraTableGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourceOutput) ToCassandraTableGetPropertiesResponseResourceOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourceOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourceOutput) ToCassandraTableGetPropertiesResponseResourcePtrOutput() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o.ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o CassandraTableGetPropertiesResponseResourceOutput) ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraTableGetPropertiesResponseResource) *CassandraTableGetPropertiesResponseResource {
		return &v
	}).(CassandraTableGetPropertiesResponseResourcePtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) *CassandraSchemaResponse { return v.Schema }).(CassandraSchemaResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v CassandraTableGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type CassandraTableGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (CassandraTableGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableGetPropertiesResponseResource)(nil)).Elem()
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) ToCassandraTableGetPropertiesResponseResourcePtrOutput() CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) ToCassandraTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) CassandraTableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Elem() CassandraTableGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) CassandraTableGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret CassandraTableGetPropertiesResponseResource
		return ret
	}).(CassandraTableGetPropertiesResponseResourceOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Schema() CassandraSchemaResponsePtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *CassandraSchemaResponse {
		if v == nil {
			return nil
		}
		return v.Schema
	}).(CassandraSchemaResponsePtrOutput)
}

func (o CassandraTableGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *CassandraTableGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type CassandraTableResource struct {
	AnalyticalStorageTtl *int             `pulumi:"analyticalStorageTtl"`
	DefaultTtl           *int             `pulumi:"defaultTtl"`
	Id                   string           `pulumi:"id"`
	Schema               *CassandraSchema `pulumi:"schema"`
}





type CassandraTableResourceInput interface {
	pulumi.Input

	ToCassandraTableResourceOutput() CassandraTableResourceOutput
	ToCassandraTableResourceOutputWithContext(context.Context) CassandraTableResourceOutput
}

type CassandraTableResourceArgs struct {
	AnalyticalStorageTtl pulumi.IntPtrInput      `pulumi:"analyticalStorageTtl"`
	DefaultTtl           pulumi.IntPtrInput      `pulumi:"defaultTtl"`
	Id                   pulumi.StringInput      `pulumi:"id"`
	Schema               CassandraSchemaPtrInput `pulumi:"schema"`
}

func (CassandraTableResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableResource)(nil)).Elem()
}

func (i CassandraTableResourceArgs) ToCassandraTableResourceOutput() CassandraTableResourceOutput {
	return i.ToCassandraTableResourceOutputWithContext(context.Background())
}

func (i CassandraTableResourceArgs) ToCassandraTableResourceOutputWithContext(ctx context.Context) CassandraTableResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableResourceOutput)
}

func (i CassandraTableResourceArgs) ToCassandraTableResourcePtrOutput() CassandraTableResourcePtrOutput {
	return i.ToCassandraTableResourcePtrOutputWithContext(context.Background())
}

func (i CassandraTableResourceArgs) ToCassandraTableResourcePtrOutputWithContext(ctx context.Context) CassandraTableResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableResourceOutput).ToCassandraTableResourcePtrOutputWithContext(ctx)
}









type CassandraTableResourcePtrInput interface {
	pulumi.Input

	ToCassandraTableResourcePtrOutput() CassandraTableResourcePtrOutput
	ToCassandraTableResourcePtrOutputWithContext(context.Context) CassandraTableResourcePtrOutput
}

type cassandraTableResourcePtrType CassandraTableResourceArgs

func CassandraTableResourcePtr(v *CassandraTableResourceArgs) CassandraTableResourcePtrInput {
	return (*cassandraTableResourcePtrType)(v)
}

func (*cassandraTableResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableResource)(nil)).Elem()
}

func (i *cassandraTableResourcePtrType) ToCassandraTableResourcePtrOutput() CassandraTableResourcePtrOutput {
	return i.ToCassandraTableResourcePtrOutputWithContext(context.Background())
}

func (i *cassandraTableResourcePtrType) ToCassandraTableResourcePtrOutputWithContext(ctx context.Context) CassandraTableResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CassandraTableResourcePtrOutput)
}

type CassandraTableResourceOutput struct{ *pulumi.OutputState }

func (CassandraTableResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CassandraTableResource)(nil)).Elem()
}

func (o CassandraTableResourceOutput) ToCassandraTableResourceOutput() CassandraTableResourceOutput {
	return o
}

func (o CassandraTableResourceOutput) ToCassandraTableResourceOutputWithContext(ctx context.Context) CassandraTableResourceOutput {
	return o
}

func (o CassandraTableResourceOutput) ToCassandraTableResourcePtrOutput() CassandraTableResourcePtrOutput {
	return o.ToCassandraTableResourcePtrOutputWithContext(context.Background())
}

func (o CassandraTableResourceOutput) ToCassandraTableResourcePtrOutputWithContext(ctx context.Context) CassandraTableResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CassandraTableResource) *CassandraTableResource {
		return &v
	}).(CassandraTableResourcePtrOutput)
}

func (o CassandraTableResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CassandraTableResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o CassandraTableResourceOutput) Schema() CassandraSchemaPtrOutput {
	return o.ApplyT(func(v CassandraTableResource) *CassandraSchema { return v.Schema }).(CassandraSchemaPtrOutput)
}

type CassandraTableResourcePtrOutput struct{ *pulumi.OutputState }

func (CassandraTableResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CassandraTableResource)(nil)).Elem()
}

func (o CassandraTableResourcePtrOutput) ToCassandraTableResourcePtrOutput() CassandraTableResourcePtrOutput {
	return o
}

func (o CassandraTableResourcePtrOutput) ToCassandraTableResourcePtrOutputWithContext(ctx context.Context) CassandraTableResourcePtrOutput {
	return o
}

func (o CassandraTableResourcePtrOutput) Elem() CassandraTableResourceOutput {
	return o.ApplyT(func(v *CassandraTableResource) CassandraTableResource {
		if v != nil {
			return *v
		}
		var ret CassandraTableResource
		return ret
	}).(CassandraTableResourceOutput)
}

func (o CassandraTableResourcePtrOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableResource) *int {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CassandraTableResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o CassandraTableResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CassandraTableResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CassandraTableResourcePtrOutput) Schema() CassandraSchemaPtrOutput {
	return o.ApplyT(func(v *CassandraTableResource) *CassandraSchema {
		if v == nil {
			return nil
		}
		return v.Schema
	}).(CassandraSchemaPtrOutput)
}

type Certificate struct {
	Pem *string `pulumi:"pem"`
}





type CertificateInput interface {
	pulumi.Input

	ToCertificateOutput() CertificateOutput
	ToCertificateOutputWithContext(context.Context) CertificateOutput
}

type CertificateArgs struct {
	Pem pulumi.StringPtrInput `pulumi:"pem"`
}

func (CertificateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil)).Elem()
}

func (i CertificateArgs) ToCertificateOutput() CertificateOutput {
	return i.ToCertificateOutputWithContext(context.Background())
}

func (i CertificateArgs) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateOutput)
}





type CertificateArrayInput interface {
	pulumi.Input

	ToCertificateArrayOutput() CertificateArrayOutput
	ToCertificateArrayOutputWithContext(context.Context) CertificateArrayOutput
}

type CertificateArray []CertificateInput

func (CertificateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil)).Elem()
}

func (i CertificateArray) ToCertificateArrayOutput() CertificateArrayOutput {
	return i.ToCertificateArrayOutputWithContext(context.Background())
}

func (i CertificateArray) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateArrayOutput)
}

type CertificateOutput struct{ *pulumi.OutputState }

func (CertificateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Certificate)(nil)).Elem()
}

func (o CertificateOutput) ToCertificateOutput() CertificateOutput {
	return o
}

func (o CertificateOutput) ToCertificateOutputWithContext(ctx context.Context) CertificateOutput {
	return o
}

func (o CertificateOutput) Pem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Certificate) *string { return v.Pem }).(pulumi.StringPtrOutput)
}

type CertificateArrayOutput struct{ *pulumi.OutputState }

func (CertificateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Certificate)(nil)).Elem()
}

func (o CertificateArrayOutput) ToCertificateArrayOutput() CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) ToCertificateArrayOutputWithContext(ctx context.Context) CertificateArrayOutput {
	return o
}

func (o CertificateArrayOutput) Index(i pulumi.IntInput) CertificateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Certificate {
		return vs[0].([]Certificate)[vs[1].(int)]
	}).(CertificateOutput)
}

type CertificateResponse struct {
	Pem *string `pulumi:"pem"`
}





type CertificateResponseInput interface {
	pulumi.Input

	ToCertificateResponseOutput() CertificateResponseOutput
	ToCertificateResponseOutputWithContext(context.Context) CertificateResponseOutput
}

type CertificateResponseArgs struct {
	Pem pulumi.StringPtrInput `pulumi:"pem"`
}

func (CertificateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateResponse)(nil)).Elem()
}

func (i CertificateResponseArgs) ToCertificateResponseOutput() CertificateResponseOutput {
	return i.ToCertificateResponseOutputWithContext(context.Background())
}

func (i CertificateResponseArgs) ToCertificateResponseOutputWithContext(ctx context.Context) CertificateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateResponseOutput)
}





type CertificateResponseArrayInput interface {
	pulumi.Input

	ToCertificateResponseArrayOutput() CertificateResponseArrayOutput
	ToCertificateResponseArrayOutputWithContext(context.Context) CertificateResponseArrayOutput
}

type CertificateResponseArray []CertificateResponseInput

func (CertificateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateResponse)(nil)).Elem()
}

func (i CertificateResponseArray) ToCertificateResponseArrayOutput() CertificateResponseArrayOutput {
	return i.ToCertificateResponseArrayOutputWithContext(context.Background())
}

func (i CertificateResponseArray) ToCertificateResponseArrayOutputWithContext(ctx context.Context) CertificateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CertificateResponseArrayOutput)
}

type CertificateResponseOutput struct{ *pulumi.OutputState }

func (CertificateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CertificateResponse)(nil)).Elem()
}

func (o CertificateResponseOutput) ToCertificateResponseOutput() CertificateResponseOutput {
	return o
}

func (o CertificateResponseOutput) ToCertificateResponseOutputWithContext(ctx context.Context) CertificateResponseOutput {
	return o
}

func (o CertificateResponseOutput) Pem() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CertificateResponse) *string { return v.Pem }).(pulumi.StringPtrOutput)
}

type CertificateResponseArrayOutput struct{ *pulumi.OutputState }

func (CertificateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CertificateResponse)(nil)).Elem()
}

func (o CertificateResponseArrayOutput) ToCertificateResponseArrayOutput() CertificateResponseArrayOutput {
	return o
}

func (o CertificateResponseArrayOutput) ToCertificateResponseArrayOutputWithContext(ctx context.Context) CertificateResponseArrayOutput {
	return o
}

func (o CertificateResponseArrayOutput) Index(i pulumi.IntInput) CertificateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CertificateResponse {
		return vs[0].([]CertificateResponse)[vs[1].(int)]
	}).(CertificateResponseOutput)
}

type ClusterKey struct {
	Name    *string `pulumi:"name"`
	OrderBy *string `pulumi:"orderBy"`
}





type ClusterKeyInput interface {
	pulumi.Input

	ToClusterKeyOutput() ClusterKeyOutput
	ToClusterKeyOutputWithContext(context.Context) ClusterKeyOutput
}

type ClusterKeyArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
}

func (ClusterKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKey)(nil)).Elem()
}

func (i ClusterKeyArgs) ToClusterKeyOutput() ClusterKeyOutput {
	return i.ToClusterKeyOutputWithContext(context.Background())
}

func (i ClusterKeyArgs) ToClusterKeyOutputWithContext(ctx context.Context) ClusterKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyOutput)
}





type ClusterKeyArrayInput interface {
	pulumi.Input

	ToClusterKeyArrayOutput() ClusterKeyArrayOutput
	ToClusterKeyArrayOutputWithContext(context.Context) ClusterKeyArrayOutput
}

type ClusterKeyArray []ClusterKeyInput

func (ClusterKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKey)(nil)).Elem()
}

func (i ClusterKeyArray) ToClusterKeyArrayOutput() ClusterKeyArrayOutput {
	return i.ToClusterKeyArrayOutputWithContext(context.Background())
}

func (i ClusterKeyArray) ToClusterKeyArrayOutputWithContext(ctx context.Context) ClusterKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyArrayOutput)
}

type ClusterKeyOutput struct{ *pulumi.OutputState }

func (ClusterKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKey)(nil)).Elem()
}

func (o ClusterKeyOutput) ToClusterKeyOutput() ClusterKeyOutput {
	return o
}

func (o ClusterKeyOutput) ToClusterKeyOutputWithContext(ctx context.Context) ClusterKeyOutput {
	return o
}

func (o ClusterKeyOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKey) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClusterKeyOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKey) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

type ClusterKeyArrayOutput struct{ *pulumi.OutputState }

func (ClusterKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKey)(nil)).Elem()
}

func (o ClusterKeyArrayOutput) ToClusterKeyArrayOutput() ClusterKeyArrayOutput {
	return o
}

func (o ClusterKeyArrayOutput) ToClusterKeyArrayOutputWithContext(ctx context.Context) ClusterKeyArrayOutput {
	return o
}

func (o ClusterKeyArrayOutput) Index(i pulumi.IntInput) ClusterKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterKey {
		return vs[0].([]ClusterKey)[vs[1].(int)]
	}).(ClusterKeyOutput)
}

type ClusterKeyResponse struct {
	Name    *string `pulumi:"name"`
	OrderBy *string `pulumi:"orderBy"`
}





type ClusterKeyResponseInput interface {
	pulumi.Input

	ToClusterKeyResponseOutput() ClusterKeyResponseOutput
	ToClusterKeyResponseOutputWithContext(context.Context) ClusterKeyResponseOutput
}

type ClusterKeyResponseArgs struct {
	Name    pulumi.StringPtrInput `pulumi:"name"`
	OrderBy pulumi.StringPtrInput `pulumi:"orderBy"`
}

func (ClusterKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKeyResponse)(nil)).Elem()
}

func (i ClusterKeyResponseArgs) ToClusterKeyResponseOutput() ClusterKeyResponseOutput {
	return i.ToClusterKeyResponseOutputWithContext(context.Background())
}

func (i ClusterKeyResponseArgs) ToClusterKeyResponseOutputWithContext(ctx context.Context) ClusterKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyResponseOutput)
}





type ClusterKeyResponseArrayInput interface {
	pulumi.Input

	ToClusterKeyResponseArrayOutput() ClusterKeyResponseArrayOutput
	ToClusterKeyResponseArrayOutputWithContext(context.Context) ClusterKeyResponseArrayOutput
}

type ClusterKeyResponseArray []ClusterKeyResponseInput

func (ClusterKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKeyResponse)(nil)).Elem()
}

func (i ClusterKeyResponseArray) ToClusterKeyResponseArrayOutput() ClusterKeyResponseArrayOutput {
	return i.ToClusterKeyResponseArrayOutputWithContext(context.Background())
}

func (i ClusterKeyResponseArray) ToClusterKeyResponseArrayOutputWithContext(ctx context.Context) ClusterKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterKeyResponseArrayOutput)
}

type ClusterKeyResponseOutput struct{ *pulumi.OutputState }

func (ClusterKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterKeyResponse)(nil)).Elem()
}

func (o ClusterKeyResponseOutput) ToClusterKeyResponseOutput() ClusterKeyResponseOutput {
	return o
}

func (o ClusterKeyResponseOutput) ToClusterKeyResponseOutputWithContext(ctx context.Context) ClusterKeyResponseOutput {
	return o
}

func (o ClusterKeyResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKeyResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ClusterKeyResponseOutput) OrderBy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterKeyResponse) *string { return v.OrderBy }).(pulumi.StringPtrOutput)
}

type ClusterKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (ClusterKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ClusterKeyResponse)(nil)).Elem()
}

func (o ClusterKeyResponseArrayOutput) ToClusterKeyResponseArrayOutput() ClusterKeyResponseArrayOutput {
	return o
}

func (o ClusterKeyResponseArrayOutput) ToClusterKeyResponseArrayOutputWithContext(ctx context.Context) ClusterKeyResponseArrayOutput {
	return o
}

func (o ClusterKeyResponseArrayOutput) Index(i pulumi.IntInput) ClusterKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ClusterKeyResponse {
		return vs[0].([]ClusterKeyResponse)[vs[1].(int)]
	}).(ClusterKeyResponseOutput)
}

type ClusterResourceProperties struct {
	AuthenticationMethod          *string       `pulumi:"authenticationMethod"`
	CassandraAuditLoggingEnabled  *bool         `pulumi:"cassandraAuditLoggingEnabled"`
	CassandraVersion              *string       `pulumi:"cassandraVersion"`
	ClientCertificates            []Certificate `pulumi:"clientCertificates"`
	ClusterNameOverride           *string       `pulumi:"clusterNameOverride"`
	Deallocated                   *bool         `pulumi:"deallocated"`
	DelegatedManagementSubnetId   *string       `pulumi:"delegatedManagementSubnetId"`
	ExternalGossipCertificates    []Certificate `pulumi:"externalGossipCertificates"`
	ExternalSeedNodes             []SeedNode    `pulumi:"externalSeedNodes"`
	HoursBetweenBackups           *int          `pulumi:"hoursBetweenBackups"`
	InitialCassandraAdminPassword *string       `pulumi:"initialCassandraAdminPassword"`
	PrometheusEndpoint            *SeedNode     `pulumi:"prometheusEndpoint"`
	ProvisioningState             *string       `pulumi:"provisioningState"`
	RepairEnabled                 *bool         `pulumi:"repairEnabled"`
	RestoreFromBackupId           *string       `pulumi:"restoreFromBackupId"`
}





type ClusterResourcePropertiesInput interface {
	pulumi.Input

	ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput
	ToClusterResourcePropertiesOutputWithContext(context.Context) ClusterResourcePropertiesOutput
}

type ClusterResourcePropertiesArgs struct {
	AuthenticationMethod          pulumi.StringPtrInput `pulumi:"authenticationMethod"`
	CassandraAuditLoggingEnabled  pulumi.BoolPtrInput   `pulumi:"cassandraAuditLoggingEnabled"`
	CassandraVersion              pulumi.StringPtrInput `pulumi:"cassandraVersion"`
	ClientCertificates            CertificateArrayInput `pulumi:"clientCertificates"`
	ClusterNameOverride           pulumi.StringPtrInput `pulumi:"clusterNameOverride"`
	Deallocated                   pulumi.BoolPtrInput   `pulumi:"deallocated"`
	DelegatedManagementSubnetId   pulumi.StringPtrInput `pulumi:"delegatedManagementSubnetId"`
	ExternalGossipCertificates    CertificateArrayInput `pulumi:"externalGossipCertificates"`
	ExternalSeedNodes             SeedNodeArrayInput    `pulumi:"externalSeedNodes"`
	HoursBetweenBackups           pulumi.IntPtrInput    `pulumi:"hoursBetweenBackups"`
	InitialCassandraAdminPassword pulumi.StringPtrInput `pulumi:"initialCassandraAdminPassword"`
	PrometheusEndpoint            SeedNodePtrInput      `pulumi:"prometheusEndpoint"`
	ProvisioningState             pulumi.StringPtrInput `pulumi:"provisioningState"`
	RepairEnabled                 pulumi.BoolPtrInput   `pulumi:"repairEnabled"`
	RestoreFromBackupId           pulumi.StringPtrInput `pulumi:"restoreFromBackupId"`
}

func (ClusterResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceProperties)(nil)).Elem()
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput {
	return i.ToClusterResourcePropertiesOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesOutputWithContext(ctx context.Context) ClusterResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesOutput)
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return i.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterResourcePropertiesArgs) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesOutput).ToClusterResourcePropertiesPtrOutputWithContext(ctx)
}









type ClusterResourcePropertiesPtrInput interface {
	pulumi.Input

	ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput
	ToClusterResourcePropertiesPtrOutputWithContext(context.Context) ClusterResourcePropertiesPtrOutput
}

type clusterResourcePropertiesPtrType ClusterResourcePropertiesArgs

func ClusterResourcePropertiesPtr(v *ClusterResourcePropertiesArgs) ClusterResourcePropertiesPtrInput {
	return (*clusterResourcePropertiesPtrType)(v)
}

func (*clusterResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceProperties)(nil)).Elem()
}

func (i *clusterResourcePropertiesPtrType) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return i.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterResourcePropertiesPtrType) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourcePropertiesPtrOutput)
}

type ClusterResourcePropertiesOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceProperties)(nil)).Elem()
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesOutput() ClusterResourcePropertiesOutput {
	return o
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesOutputWithContext(ctx context.Context) ClusterResourcePropertiesOutput {
	return o
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return o.ToClusterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterResourcePropertiesOutput) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterResourceProperties) *ClusterResourceProperties {
		return &v
	}).(ClusterResourcePropertiesPtrOutput)
}

func (o ClusterResourcePropertiesOutput) AuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.AuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesOutput) CassandraAuditLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *bool { return v.CassandraAuditLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o ClusterResourcePropertiesOutput) CassandraVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.CassandraVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesOutput) ClientCertificates() CertificateArrayOutput {
	return o.ApplyT(func(v ClusterResourceProperties) []Certificate { return v.ClientCertificates }).(CertificateArrayOutput)
}

func (o ClusterResourcePropertiesOutput) ClusterNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.ClusterNameOverride }).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesOutput) Deallocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *bool { return v.Deallocated }).(pulumi.BoolPtrOutput)
}

func (o ClusterResourcePropertiesOutput) DelegatedManagementSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.DelegatedManagementSubnetId }).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesOutput) ExternalGossipCertificates() CertificateArrayOutput {
	return o.ApplyT(func(v ClusterResourceProperties) []Certificate { return v.ExternalGossipCertificates }).(CertificateArrayOutput)
}

func (o ClusterResourcePropertiesOutput) ExternalSeedNodes() SeedNodeArrayOutput {
	return o.ApplyT(func(v ClusterResourceProperties) []SeedNode { return v.ExternalSeedNodes }).(SeedNodeArrayOutput)
}

func (o ClusterResourcePropertiesOutput) HoursBetweenBackups() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *int { return v.HoursBetweenBackups }).(pulumi.IntPtrOutput)
}

func (o ClusterResourcePropertiesOutput) InitialCassandraAdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.InitialCassandraAdminPassword }).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesOutput) PrometheusEndpoint() SeedNodePtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *SeedNode { return v.PrometheusEndpoint }).(SeedNodePtrOutput)
}

func (o ClusterResourcePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesOutput) RepairEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *bool { return v.RepairEnabled }).(pulumi.BoolPtrOutput)
}

func (o ClusterResourcePropertiesOutput) RestoreFromBackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceProperties) *string { return v.RestoreFromBackupId }).(pulumi.StringPtrOutput)
}

type ClusterResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceProperties)(nil)).Elem()
}

func (o ClusterResourcePropertiesPtrOutput) ToClusterResourcePropertiesPtrOutput() ClusterResourcePropertiesPtrOutput {
	return o
}

func (o ClusterResourcePropertiesPtrOutput) ToClusterResourcePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourcePropertiesPtrOutput {
	return o
}

func (o ClusterResourcePropertiesPtrOutput) Elem() ClusterResourcePropertiesOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) ClusterResourceProperties {
		if v != nil {
			return *v
		}
		var ret ClusterResourceProperties
		return ret
	}).(ClusterResourcePropertiesOutput)
}

func (o ClusterResourcePropertiesPtrOutput) AuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuthenticationMethod
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) CassandraAuditLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.CassandraAuditLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) CassandraVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.CassandraVersion
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) ClientCertificates() CertificateArrayOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) []Certificate {
		if v == nil {
			return nil
		}
		return v.ClientCertificates
	}).(CertificateArrayOutput)
}

func (o ClusterResourcePropertiesPtrOutput) ClusterNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterNameOverride
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) Deallocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Deallocated
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) DelegatedManagementSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedManagementSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) ExternalGossipCertificates() CertificateArrayOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) []Certificate {
		if v == nil {
			return nil
		}
		return v.ExternalGossipCertificates
	}).(CertificateArrayOutput)
}

func (o ClusterResourcePropertiesPtrOutput) ExternalSeedNodes() SeedNodeArrayOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) []SeedNode {
		if v == nil {
			return nil
		}
		return v.ExternalSeedNodes
	}).(SeedNodeArrayOutput)
}

func (o ClusterResourcePropertiesPtrOutput) HoursBetweenBackups() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *int {
		if v == nil {
			return nil
		}
		return v.HoursBetweenBackups
	}).(pulumi.IntPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) InitialCassandraAdminPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.InitialCassandraAdminPassword
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) PrometheusEndpoint() SeedNodePtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *SeedNode {
		if v == nil {
			return nil
		}
		return v.PrometheusEndpoint
	}).(SeedNodePtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) RepairEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.RepairEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterResourcePropertiesPtrOutput) RestoreFromBackupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.RestoreFromBackupId
	}).(pulumi.StringPtrOutput)
}

type ClusterResourceResponseProperties struct {
	AuthenticationMethod         *string               `pulumi:"authenticationMethod"`
	CassandraAuditLoggingEnabled *bool                 `pulumi:"cassandraAuditLoggingEnabled"`
	CassandraVersion             *string               `pulumi:"cassandraVersion"`
	ClientCertificates           []CertificateResponse `pulumi:"clientCertificates"`
	ClusterNameOverride          *string               `pulumi:"clusterNameOverride"`
	Deallocated                  *bool                 `pulumi:"deallocated"`
	DelegatedManagementSubnetId  *string               `pulumi:"delegatedManagementSubnetId"`
	ExternalGossipCertificates   []CertificateResponse `pulumi:"externalGossipCertificates"`
	ExternalSeedNodes            []SeedNodeResponse    `pulumi:"externalSeedNodes"`
	GossipCertificates           []CertificateResponse `pulumi:"gossipCertificates"`
	HoursBetweenBackups          *int                  `pulumi:"hoursBetweenBackups"`
	PrometheusEndpoint           *SeedNodeResponse     `pulumi:"prometheusEndpoint"`
	ProvisioningState            *string               `pulumi:"provisioningState"`
	RepairEnabled                *bool                 `pulumi:"repairEnabled"`
	SeedNodes                    []SeedNodeResponse    `pulumi:"seedNodes"`
}





type ClusterResourceResponsePropertiesInput interface {
	pulumi.Input

	ToClusterResourceResponsePropertiesOutput() ClusterResourceResponsePropertiesOutput
	ToClusterResourceResponsePropertiesOutputWithContext(context.Context) ClusterResourceResponsePropertiesOutput
}

type ClusterResourceResponsePropertiesArgs struct {
	AuthenticationMethod         pulumi.StringPtrInput         `pulumi:"authenticationMethod"`
	CassandraAuditLoggingEnabled pulumi.BoolPtrInput           `pulumi:"cassandraAuditLoggingEnabled"`
	CassandraVersion             pulumi.StringPtrInput         `pulumi:"cassandraVersion"`
	ClientCertificates           CertificateResponseArrayInput `pulumi:"clientCertificates"`
	ClusterNameOverride          pulumi.StringPtrInput         `pulumi:"clusterNameOverride"`
	Deallocated                  pulumi.BoolPtrInput           `pulumi:"deallocated"`
	DelegatedManagementSubnetId  pulumi.StringPtrInput         `pulumi:"delegatedManagementSubnetId"`
	ExternalGossipCertificates   CertificateResponseArrayInput `pulumi:"externalGossipCertificates"`
	ExternalSeedNodes            SeedNodeResponseArrayInput    `pulumi:"externalSeedNodes"`
	GossipCertificates           CertificateResponseArrayInput `pulumi:"gossipCertificates"`
	HoursBetweenBackups          pulumi.IntPtrInput            `pulumi:"hoursBetweenBackups"`
	PrometheusEndpoint           SeedNodeResponsePtrInput      `pulumi:"prometheusEndpoint"`
	ProvisioningState            pulumi.StringPtrInput         `pulumi:"provisioningState"`
	RepairEnabled                pulumi.BoolPtrInput           `pulumi:"repairEnabled"`
	SeedNodes                    SeedNodeResponseArrayInput    `pulumi:"seedNodes"`
}

func (ClusterResourceResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceResponseProperties)(nil)).Elem()
}

func (i ClusterResourceResponsePropertiesArgs) ToClusterResourceResponsePropertiesOutput() ClusterResourceResponsePropertiesOutput {
	return i.ToClusterResourceResponsePropertiesOutputWithContext(context.Background())
}

func (i ClusterResourceResponsePropertiesArgs) ToClusterResourceResponsePropertiesOutputWithContext(ctx context.Context) ClusterResourceResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourceResponsePropertiesOutput)
}

func (i ClusterResourceResponsePropertiesArgs) ToClusterResourceResponsePropertiesPtrOutput() ClusterResourceResponsePropertiesPtrOutput {
	return i.ToClusterResourceResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i ClusterResourceResponsePropertiesArgs) ToClusterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourceResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourceResponsePropertiesOutput).ToClusterResourceResponsePropertiesPtrOutputWithContext(ctx)
}









type ClusterResourceResponsePropertiesPtrInput interface {
	pulumi.Input

	ToClusterResourceResponsePropertiesPtrOutput() ClusterResourceResponsePropertiesPtrOutput
	ToClusterResourceResponsePropertiesPtrOutputWithContext(context.Context) ClusterResourceResponsePropertiesPtrOutput
}

type clusterResourceResponsePropertiesPtrType ClusterResourceResponsePropertiesArgs

func ClusterResourceResponsePropertiesPtr(v *ClusterResourceResponsePropertiesArgs) ClusterResourceResponsePropertiesPtrInput {
	return (*clusterResourceResponsePropertiesPtrType)(v)
}

func (*clusterResourceResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceResponseProperties)(nil)).Elem()
}

func (i *clusterResourceResponsePropertiesPtrType) ToClusterResourceResponsePropertiesPtrOutput() ClusterResourceResponsePropertiesPtrOutput {
	return i.ToClusterResourceResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *clusterResourceResponsePropertiesPtrType) ToClusterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourceResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ClusterResourceResponsePropertiesPtrOutput)
}

type ClusterResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (ClusterResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ClusterResourceResponseProperties)(nil)).Elem()
}

func (o ClusterResourceResponsePropertiesOutput) ToClusterResourceResponsePropertiesOutput() ClusterResourceResponsePropertiesOutput {
	return o
}

func (o ClusterResourceResponsePropertiesOutput) ToClusterResourceResponsePropertiesOutputWithContext(ctx context.Context) ClusterResourceResponsePropertiesOutput {
	return o
}

func (o ClusterResourceResponsePropertiesOutput) ToClusterResourceResponsePropertiesPtrOutput() ClusterResourceResponsePropertiesPtrOutput {
	return o.ToClusterResourceResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o ClusterResourceResponsePropertiesOutput) ToClusterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourceResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ClusterResourceResponseProperties) *ClusterResourceResponseProperties {
		return &v
	}).(ClusterResourceResponsePropertiesPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) AuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *string { return v.AuthenticationMethod }).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) CassandraAuditLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *bool { return v.CassandraAuditLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) CassandraVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *string { return v.CassandraVersion }).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) ClientCertificates() CertificateResponseArrayOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) []CertificateResponse { return v.ClientCertificates }).(CertificateResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesOutput) ClusterNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *string { return v.ClusterNameOverride }).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) Deallocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *bool { return v.Deallocated }).(pulumi.BoolPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) DelegatedManagementSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *string { return v.DelegatedManagementSubnetId }).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) ExternalGossipCertificates() CertificateResponseArrayOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) []CertificateResponse { return v.ExternalGossipCertificates }).(CertificateResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesOutput) ExternalSeedNodes() SeedNodeResponseArrayOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) []SeedNodeResponse { return v.ExternalSeedNodes }).(SeedNodeResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesOutput) GossipCertificates() CertificateResponseArrayOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) []CertificateResponse { return v.GossipCertificates }).(CertificateResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesOutput) HoursBetweenBackups() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *int { return v.HoursBetweenBackups }).(pulumi.IntPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) PrometheusEndpoint() SeedNodeResponsePtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *SeedNodeResponse { return v.PrometheusEndpoint }).(SeedNodeResponsePtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) RepairEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) *bool { return v.RepairEnabled }).(pulumi.BoolPtrOutput)
}

func (o ClusterResourceResponsePropertiesOutput) SeedNodes() SeedNodeResponseArrayOutput {
	return o.ApplyT(func(v ClusterResourceResponseProperties) []SeedNodeResponse { return v.SeedNodes }).(SeedNodeResponseArrayOutput)
}

type ClusterResourceResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (ClusterResourceResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ClusterResourceResponseProperties)(nil)).Elem()
}

func (o ClusterResourceResponsePropertiesPtrOutput) ToClusterResourceResponsePropertiesPtrOutput() ClusterResourceResponsePropertiesPtrOutput {
	return o
}

func (o ClusterResourceResponsePropertiesPtrOutput) ToClusterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) ClusterResourceResponsePropertiesPtrOutput {
	return o
}

func (o ClusterResourceResponsePropertiesPtrOutput) Elem() ClusterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) ClusterResourceResponseProperties {
		if v != nil {
			return *v
		}
		var ret ClusterResourceResponseProperties
		return ret
	}).(ClusterResourceResponsePropertiesOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) AuthenticationMethod() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuthenticationMethod
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) CassandraAuditLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.CassandraAuditLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) CassandraVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.CassandraVersion
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) ClientCertificates() CertificateResponseArrayOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) []CertificateResponse {
		if v == nil {
			return nil
		}
		return v.ClientCertificates
	}).(CertificateResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) ClusterNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClusterNameOverride
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) Deallocated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.Deallocated
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) DelegatedManagementSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedManagementSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) ExternalGossipCertificates() CertificateResponseArrayOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) []CertificateResponse {
		if v == nil {
			return nil
		}
		return v.ExternalGossipCertificates
	}).(CertificateResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) ExternalSeedNodes() SeedNodeResponseArrayOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) []SeedNodeResponse {
		if v == nil {
			return nil
		}
		return v.ExternalSeedNodes
	}).(SeedNodeResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) GossipCertificates() CertificateResponseArrayOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) []CertificateResponse {
		if v == nil {
			return nil
		}
		return v.GossipCertificates
	}).(CertificateResponseArrayOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) HoursBetweenBackups() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.HoursBetweenBackups
	}).(pulumi.IntPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) PrometheusEndpoint() SeedNodeResponsePtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *SeedNodeResponse {
		if v == nil {
			return nil
		}
		return v.PrometheusEndpoint
	}).(SeedNodeResponsePtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) RepairEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.RepairEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ClusterResourceResponsePropertiesPtrOutput) SeedNodes() SeedNodeResponseArrayOutput {
	return o.ApplyT(func(v *ClusterResourceResponseProperties) []SeedNodeResponse {
		if v == nil {
			return nil
		}
		return v.SeedNodes
	}).(SeedNodeResponseArrayOutput)
}

type Column struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ColumnInput interface {
	pulumi.Input

	ToColumnOutput() ColumnOutput
	ToColumnOutputWithContext(context.Context) ColumnOutput
}

type ColumnArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ColumnArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (i ColumnArgs) ToColumnOutput() ColumnOutput {
	return i.ToColumnOutputWithContext(context.Background())
}

func (i ColumnArgs) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnOutput)
}





type ColumnArrayInput interface {
	pulumi.Input

	ToColumnArrayOutput() ColumnArrayOutput
	ToColumnArrayOutputWithContext(context.Context) ColumnArrayOutput
}

type ColumnArray []ColumnInput

func (ColumnArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (i ColumnArray) ToColumnArrayOutput() ColumnArrayOutput {
	return i.ToColumnArrayOutputWithContext(context.Background())
}

func (i ColumnArray) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnArrayOutput)
}

type ColumnOutput struct{ *pulumi.OutputState }

func (ColumnOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Column)(nil)).Elem()
}

func (o ColumnOutput) ToColumnOutput() ColumnOutput {
	return o
}

func (o ColumnOutput) ToColumnOutputWithContext(ctx context.Context) ColumnOutput {
	return o
}

func (o ColumnOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Column) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnArrayOutput struct{ *pulumi.OutputState }

func (ColumnArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Column)(nil)).Elem()
}

func (o ColumnArrayOutput) ToColumnArrayOutput() ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) ToColumnArrayOutputWithContext(ctx context.Context) ColumnArrayOutput {
	return o
}

func (o ColumnArrayOutput) Index(i pulumi.IntInput) ColumnOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Column {
		return vs[0].([]Column)[vs[1].(int)]
	}).(ColumnOutput)
}

type ColumnResponse struct {
	Name *string `pulumi:"name"`
	Type *string `pulumi:"type"`
}





type ColumnResponseInput interface {
	pulumi.Input

	ToColumnResponseOutput() ColumnResponseOutput
	ToColumnResponseOutputWithContext(context.Context) ColumnResponseOutput
}

type ColumnResponseArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ColumnResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnResponse)(nil)).Elem()
}

func (i ColumnResponseArgs) ToColumnResponseOutput() ColumnResponseOutput {
	return i.ToColumnResponseOutputWithContext(context.Background())
}

func (i ColumnResponseArgs) ToColumnResponseOutputWithContext(ctx context.Context) ColumnResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnResponseOutput)
}





type ColumnResponseArrayInput interface {
	pulumi.Input

	ToColumnResponseArrayOutput() ColumnResponseArrayOutput
	ToColumnResponseArrayOutputWithContext(context.Context) ColumnResponseArrayOutput
}

type ColumnResponseArray []ColumnResponseInput

func (ColumnResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnResponse)(nil)).Elem()
}

func (i ColumnResponseArray) ToColumnResponseArrayOutput() ColumnResponseArrayOutput {
	return i.ToColumnResponseArrayOutputWithContext(context.Background())
}

func (i ColumnResponseArray) ToColumnResponseArrayOutputWithContext(ctx context.Context) ColumnResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ColumnResponseArrayOutput)
}

type ColumnResponseOutput struct{ *pulumi.OutputState }

func (ColumnResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseOutput) ToColumnResponseOutput() ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) ToColumnResponseOutputWithContext(ctx context.Context) ColumnResponseOutput {
	return o
}

func (o ColumnResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ColumnResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ColumnResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ColumnResponseArrayOutput struct{ *pulumi.OutputState }

func (ColumnResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ColumnResponse)(nil)).Elem()
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutput() ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) ToColumnResponseArrayOutputWithContext(ctx context.Context) ColumnResponseArrayOutput {
	return o
}

func (o ColumnResponseArrayOutput) Index(i pulumi.IntInput) ColumnResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ColumnResponse {
		return vs[0].([]ColumnResponse)[vs[1].(int)]
	}).(ColumnResponseOutput)
}

type CompositePath struct {
	Order *string `pulumi:"order"`
	Path  *string `pulumi:"path"`
}





type CompositePathInput interface {
	pulumi.Input

	ToCompositePathOutput() CompositePathOutput
	ToCompositePathOutputWithContext(context.Context) CompositePathOutput
}

type CompositePathArgs struct {
	Order pulumi.StringPtrInput `pulumi:"order"`
	Path  pulumi.StringPtrInput `pulumi:"path"`
}

func (CompositePathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePath)(nil)).Elem()
}

func (i CompositePathArgs) ToCompositePathOutput() CompositePathOutput {
	return i.ToCompositePathOutputWithContext(context.Background())
}

func (i CompositePathArgs) ToCompositePathOutputWithContext(ctx context.Context) CompositePathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathOutput)
}





type CompositePathArrayInput interface {
	pulumi.Input

	ToCompositePathArrayOutput() CompositePathArrayOutput
	ToCompositePathArrayOutputWithContext(context.Context) CompositePathArrayOutput
}

type CompositePathArray []CompositePathInput

func (CompositePathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePath)(nil)).Elem()
}

func (i CompositePathArray) ToCompositePathArrayOutput() CompositePathArrayOutput {
	return i.ToCompositePathArrayOutputWithContext(context.Background())
}

func (i CompositePathArray) ToCompositePathArrayOutputWithContext(ctx context.Context) CompositePathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathArrayOutput)
}

type CompositePathOutput struct{ *pulumi.OutputState }

func (CompositePathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePath)(nil)).Elem()
}

func (o CompositePathOutput) ToCompositePathOutput() CompositePathOutput {
	return o
}

func (o CompositePathOutput) ToCompositePathOutputWithContext(ctx context.Context) CompositePathOutput {
	return o
}

func (o CompositePathOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePath) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o CompositePathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type CompositePathArrayOutput struct{ *pulumi.OutputState }

func (CompositePathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePath)(nil)).Elem()
}

func (o CompositePathArrayOutput) ToCompositePathArrayOutput() CompositePathArrayOutput {
	return o
}

func (o CompositePathArrayOutput) ToCompositePathArrayOutputWithContext(ctx context.Context) CompositePathArrayOutput {
	return o
}

func (o CompositePathArrayOutput) Index(i pulumi.IntInput) CompositePathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CompositePath {
		return vs[0].([]CompositePath)[vs[1].(int)]
	}).(CompositePathOutput)
}

type CompositePathResponse struct {
	Order *string `pulumi:"order"`
	Path  *string `pulumi:"path"`
}





type CompositePathResponseInput interface {
	pulumi.Input

	ToCompositePathResponseOutput() CompositePathResponseOutput
	ToCompositePathResponseOutputWithContext(context.Context) CompositePathResponseOutput
}

type CompositePathResponseArgs struct {
	Order pulumi.StringPtrInput `pulumi:"order"`
	Path  pulumi.StringPtrInput `pulumi:"path"`
}

func (CompositePathResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathResponse)(nil)).Elem()
}

func (i CompositePathResponseArgs) ToCompositePathResponseOutput() CompositePathResponseOutput {
	return i.ToCompositePathResponseOutputWithContext(context.Background())
}

func (i CompositePathResponseArgs) ToCompositePathResponseOutputWithContext(ctx context.Context) CompositePathResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathResponseOutput)
}





type CompositePathResponseArrayInput interface {
	pulumi.Input

	ToCompositePathResponseArrayOutput() CompositePathResponseArrayOutput
	ToCompositePathResponseArrayOutputWithContext(context.Context) CompositePathResponseArrayOutput
}

type CompositePathResponseArray []CompositePathResponseInput

func (CompositePathResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePathResponse)(nil)).Elem()
}

func (i CompositePathResponseArray) ToCompositePathResponseArrayOutput() CompositePathResponseArrayOutput {
	return i.ToCompositePathResponseArrayOutputWithContext(context.Background())
}

func (i CompositePathResponseArray) ToCompositePathResponseArrayOutputWithContext(ctx context.Context) CompositePathResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathResponseArrayOutput)
}

type CompositePathResponseOutput struct{ *pulumi.OutputState }

func (CompositePathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathResponse)(nil)).Elem()
}

func (o CompositePathResponseOutput) ToCompositePathResponseOutput() CompositePathResponseOutput {
	return o
}

func (o CompositePathResponseOutput) ToCompositePathResponseOutputWithContext(ctx context.Context) CompositePathResponseOutput {
	return o
}

func (o CompositePathResponseOutput) Order() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePathResponse) *string { return v.Order }).(pulumi.StringPtrOutput)
}

func (o CompositePathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CompositePathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type CompositePathResponseArrayOutput struct{ *pulumi.OutputState }

func (CompositePathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePathResponse)(nil)).Elem()
}

func (o CompositePathResponseArrayOutput) ToCompositePathResponseArrayOutput() CompositePathResponseArrayOutput {
	return o
}

func (o CompositePathResponseArrayOutput) ToCompositePathResponseArrayOutputWithContext(ctx context.Context) CompositePathResponseArrayOutput {
	return o
}

func (o CompositePathResponseArrayOutput) Index(i pulumi.IntInput) CompositePathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CompositePathResponse {
		return vs[0].([]CompositePathResponse)[vs[1].(int)]
	}).(CompositePathResponseOutput)
}

type ConflictResolutionPolicy struct {
	ConflictResolutionPath      *string `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure *string `pulumi:"conflictResolutionProcedure"`
	Mode                        *string `pulumi:"mode"`
}


func (val *ConflictResolutionPolicy) Defaults() *ConflictResolutionPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "LastWriterWins"
		tmp.Mode = &mode_
	}
	return &tmp
}





type ConflictResolutionPolicyInput interface {
	pulumi.Input

	ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput
	ToConflictResolutionPolicyOutputWithContext(context.Context) ConflictResolutionPolicyOutput
}

type ConflictResolutionPolicyArgs struct {
	ConflictResolutionPath      pulumi.StringPtrInput `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure pulumi.StringPtrInput `pulumi:"conflictResolutionProcedure"`
	Mode                        pulumi.StringPtrInput `pulumi:"mode"`
}

func (ConflictResolutionPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicy)(nil)).Elem()
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput {
	return i.ToConflictResolutionPolicyOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyOutputWithContext(ctx context.Context) ConflictResolutionPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyOutput)
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return i.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyArgs) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyOutput).ToConflictResolutionPolicyPtrOutputWithContext(ctx)
}









type ConflictResolutionPolicyPtrInput interface {
	pulumi.Input

	ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput
	ToConflictResolutionPolicyPtrOutputWithContext(context.Context) ConflictResolutionPolicyPtrOutput
}

type conflictResolutionPolicyPtrType ConflictResolutionPolicyArgs

func ConflictResolutionPolicyPtr(v *ConflictResolutionPolicyArgs) ConflictResolutionPolicyPtrInput {
	return (*conflictResolutionPolicyPtrType)(v)
}

func (*conflictResolutionPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicy)(nil)).Elem()
}

func (i *conflictResolutionPolicyPtrType) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return i.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (i *conflictResolutionPolicyPtrType) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyPtrOutput)
}

type ConflictResolutionPolicyOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicy)(nil)).Elem()
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyOutput() ConflictResolutionPolicyOutput {
	return o
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyOutputWithContext(ctx context.Context) ConflictResolutionPolicyOutput {
	return o
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return o.ToConflictResolutionPolicyPtrOutputWithContext(context.Background())
}

func (o ConflictResolutionPolicyOutput) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConflictResolutionPolicy) *ConflictResolutionPolicy {
		return &v
	}).(ConflictResolutionPolicyPtrOutput)
}

func (o ConflictResolutionPolicyOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.ConflictResolutionPath }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.ConflictResolutionProcedure }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicy) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyPtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicy)(nil)).Elem()
}

func (o ConflictResolutionPolicyPtrOutput) ToConflictResolutionPolicyPtrOutput() ConflictResolutionPolicyPtrOutput {
	return o
}

func (o ConflictResolutionPolicyPtrOutput) ToConflictResolutionPolicyPtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyPtrOutput {
	return o
}

func (o ConflictResolutionPolicyPtrOutput) Elem() ConflictResolutionPolicyOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) ConflictResolutionPolicy {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionPolicy
		return ret
	}).(ConflictResolutionPolicyOutput)
}

func (o ConflictResolutionPolicyPtrOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPath
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyPtrOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionProcedure
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyPtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicy) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyResponse struct {
	ConflictResolutionPath      *string `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure *string `pulumi:"conflictResolutionProcedure"`
	Mode                        *string `pulumi:"mode"`
}


func (val *ConflictResolutionPolicyResponse) Defaults() *ConflictResolutionPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Mode) {
		mode_ := "LastWriterWins"
		tmp.Mode = &mode_
	}
	return &tmp
}





type ConflictResolutionPolicyResponseInput interface {
	pulumi.Input

	ToConflictResolutionPolicyResponseOutput() ConflictResolutionPolicyResponseOutput
	ToConflictResolutionPolicyResponseOutputWithContext(context.Context) ConflictResolutionPolicyResponseOutput
}

type ConflictResolutionPolicyResponseArgs struct {
	ConflictResolutionPath      pulumi.StringPtrInput `pulumi:"conflictResolutionPath"`
	ConflictResolutionProcedure pulumi.StringPtrInput `pulumi:"conflictResolutionProcedure"`
	Mode                        pulumi.StringPtrInput `pulumi:"mode"`
}

func (ConflictResolutionPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (i ConflictResolutionPolicyResponseArgs) ToConflictResolutionPolicyResponseOutput() ConflictResolutionPolicyResponseOutput {
	return i.ToConflictResolutionPolicyResponseOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyResponseArgs) ToConflictResolutionPolicyResponseOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyResponseOutput)
}

func (i ConflictResolutionPolicyResponseArgs) ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput {
	return i.ToConflictResolutionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ConflictResolutionPolicyResponseArgs) ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyResponseOutput).ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx)
}









type ConflictResolutionPolicyResponsePtrInput interface {
	pulumi.Input

	ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput
	ToConflictResolutionPolicyResponsePtrOutputWithContext(context.Context) ConflictResolutionPolicyResponsePtrOutput
}

type conflictResolutionPolicyResponsePtrType ConflictResolutionPolicyResponseArgs

func ConflictResolutionPolicyResponsePtr(v *ConflictResolutionPolicyResponseArgs) ConflictResolutionPolicyResponsePtrInput {
	return (*conflictResolutionPolicyResponsePtrType)(v)
}

func (*conflictResolutionPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (i *conflictResolutionPolicyResponsePtrType) ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput {
	return i.ToConflictResolutionPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *conflictResolutionPolicyResponsePtrType) ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConflictResolutionPolicyResponsePtrOutput)
}

type ConflictResolutionPolicyResponseOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponseOutput() ConflictResolutionPolicyResponseOutput {
	return o
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponseOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponseOutput {
	return o
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput {
	return o.ToConflictResolutionPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ConflictResolutionPolicyResponseOutput) ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConflictResolutionPolicyResponse) *ConflictResolutionPolicyResponse {
		return &v
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.ConflictResolutionPath }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.ConflictResolutionProcedure }).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponseOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConflictResolutionPolicyResponse) *string { return v.Mode }).(pulumi.StringPtrOutput)
}

type ConflictResolutionPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConflictResolutionPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConflictResolutionPolicyResponse)(nil)).Elem()
}

func (o ConflictResolutionPolicyResponsePtrOutput) ToConflictResolutionPolicyResponsePtrOutput() ConflictResolutionPolicyResponsePtrOutput {
	return o
}

func (o ConflictResolutionPolicyResponsePtrOutput) ToConflictResolutionPolicyResponsePtrOutputWithContext(ctx context.Context) ConflictResolutionPolicyResponsePtrOutput {
	return o
}

func (o ConflictResolutionPolicyResponsePtrOutput) Elem() ConflictResolutionPolicyResponseOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) ConflictResolutionPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ConflictResolutionPolicyResponse
		return ret
	}).(ConflictResolutionPolicyResponseOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) ConflictResolutionPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPath
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) ConflictResolutionProcedure() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionProcedure
	}).(pulumi.StringPtrOutput)
}

func (o ConflictResolutionPolicyResponsePtrOutput) Mode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConflictResolutionPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Mode
	}).(pulumi.StringPtrOutput)
}

type ConsistencyPolicy struct {
	DefaultConsistencyLevel DefaultConsistencyLevel `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    *int                    `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      *float64                `pulumi:"maxStalenessPrefix"`
}





type ConsistencyPolicyInput interface {
	pulumi.Input

	ToConsistencyPolicyOutput() ConsistencyPolicyOutput
	ToConsistencyPolicyOutputWithContext(context.Context) ConsistencyPolicyOutput
}

type ConsistencyPolicyArgs struct {
	DefaultConsistencyLevel DefaultConsistencyLevelInput `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    pulumi.IntPtrInput           `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      pulumi.Float64PtrInput       `pulumi:"maxStalenessPrefix"`
}

func (ConsistencyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicy)(nil)).Elem()
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyOutput() ConsistencyPolicyOutput {
	return i.ToConsistencyPolicyOutputWithContext(context.Background())
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyOutputWithContext(ctx context.Context) ConsistencyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyOutput)
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return i.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (i ConsistencyPolicyArgs) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyOutput).ToConsistencyPolicyPtrOutputWithContext(ctx)
}









type ConsistencyPolicyPtrInput interface {
	pulumi.Input

	ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput
	ToConsistencyPolicyPtrOutputWithContext(context.Context) ConsistencyPolicyPtrOutput
}

type consistencyPolicyPtrType ConsistencyPolicyArgs

func ConsistencyPolicyPtr(v *ConsistencyPolicyArgs) ConsistencyPolicyPtrInput {
	return (*consistencyPolicyPtrType)(v)
}

func (*consistencyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicy)(nil)).Elem()
}

func (i *consistencyPolicyPtrType) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return i.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (i *consistencyPolicyPtrType) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyPtrOutput)
}

type ConsistencyPolicyOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicy)(nil)).Elem()
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyOutput() ConsistencyPolicyOutput {
	return o
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyOutputWithContext(ctx context.Context) ConsistencyPolicyOutput {
	return o
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return o.ToConsistencyPolicyPtrOutputWithContext(context.Background())
}

func (o ConsistencyPolicyOutput) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsistencyPolicy) *ConsistencyPolicy {
		return &v
	}).(ConsistencyPolicyPtrOutput)
}

func (o ConsistencyPolicyOutput) DefaultConsistencyLevel() DefaultConsistencyLevelOutput {
	return o.ApplyT(func(v ConsistencyPolicy) DefaultConsistencyLevel { return v.DefaultConsistencyLevel }).(DefaultConsistencyLevelOutput)
}

func (o ConsistencyPolicyOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConsistencyPolicy) *int { return v.MaxIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConsistencyPolicy) *float64 { return v.MaxStalenessPrefix }).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyPtrOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicy)(nil)).Elem()
}

func (o ConsistencyPolicyPtrOutput) ToConsistencyPolicyPtrOutput() ConsistencyPolicyPtrOutput {
	return o
}

func (o ConsistencyPolicyPtrOutput) ToConsistencyPolicyPtrOutputWithContext(ctx context.Context) ConsistencyPolicyPtrOutput {
	return o
}

func (o ConsistencyPolicyPtrOutput) Elem() ConsistencyPolicyOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) ConsistencyPolicy {
		if v != nil {
			return *v
		}
		var ret ConsistencyPolicy
		return ret
	}).(ConsistencyPolicyOutput)
}

func (o ConsistencyPolicyPtrOutput) DefaultConsistencyLevel() DefaultConsistencyLevelPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *DefaultConsistencyLevel {
		if v == nil {
			return nil
		}
		return &v.DefaultConsistencyLevel
	}).(DefaultConsistencyLevelPtrOutput)
}

func (o ConsistencyPolicyPtrOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *int {
		if v == nil {
			return nil
		}
		return v.MaxIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyPtrOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicy) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxStalenessPrefix
	}).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyResponse struct {
	DefaultConsistencyLevel string   `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    *int     `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      *float64 `pulumi:"maxStalenessPrefix"`
}





type ConsistencyPolicyResponseInput interface {
	pulumi.Input

	ToConsistencyPolicyResponseOutput() ConsistencyPolicyResponseOutput
	ToConsistencyPolicyResponseOutputWithContext(context.Context) ConsistencyPolicyResponseOutput
}

type ConsistencyPolicyResponseArgs struct {
	DefaultConsistencyLevel pulumi.StringInput     `pulumi:"defaultConsistencyLevel"`
	MaxIntervalInSeconds    pulumi.IntPtrInput     `pulumi:"maxIntervalInSeconds"`
	MaxStalenessPrefix      pulumi.Float64PtrInput `pulumi:"maxStalenessPrefix"`
}

func (ConsistencyPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicyResponse)(nil)).Elem()
}

func (i ConsistencyPolicyResponseArgs) ToConsistencyPolicyResponseOutput() ConsistencyPolicyResponseOutput {
	return i.ToConsistencyPolicyResponseOutputWithContext(context.Background())
}

func (i ConsistencyPolicyResponseArgs) ToConsistencyPolicyResponseOutputWithContext(ctx context.Context) ConsistencyPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyResponseOutput)
}

func (i ConsistencyPolicyResponseArgs) ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput {
	return i.ToConsistencyPolicyResponsePtrOutputWithContext(context.Background())
}

func (i ConsistencyPolicyResponseArgs) ToConsistencyPolicyResponsePtrOutputWithContext(ctx context.Context) ConsistencyPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyResponseOutput).ToConsistencyPolicyResponsePtrOutputWithContext(ctx)
}









type ConsistencyPolicyResponsePtrInput interface {
	pulumi.Input

	ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput
	ToConsistencyPolicyResponsePtrOutputWithContext(context.Context) ConsistencyPolicyResponsePtrOutput
}

type consistencyPolicyResponsePtrType ConsistencyPolicyResponseArgs

func ConsistencyPolicyResponsePtr(v *ConsistencyPolicyResponseArgs) ConsistencyPolicyResponsePtrInput {
	return (*consistencyPolicyResponsePtrType)(v)
}

func (*consistencyPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicyResponse)(nil)).Elem()
}

func (i *consistencyPolicyResponsePtrType) ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput {
	return i.ToConsistencyPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *consistencyPolicyResponsePtrType) ToConsistencyPolicyResponsePtrOutputWithContext(ctx context.Context) ConsistencyPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConsistencyPolicyResponsePtrOutput)
}

type ConsistencyPolicyResponseOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConsistencyPolicyResponse)(nil)).Elem()
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponseOutput() ConsistencyPolicyResponseOutput {
	return o
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponseOutputWithContext(ctx context.Context) ConsistencyPolicyResponseOutput {
	return o
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput {
	return o.ToConsistencyPolicyResponsePtrOutputWithContext(context.Background())
}

func (o ConsistencyPolicyResponseOutput) ToConsistencyPolicyResponsePtrOutputWithContext(ctx context.Context) ConsistencyPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConsistencyPolicyResponse) *ConsistencyPolicyResponse {
		return &v
	}).(ConsistencyPolicyResponsePtrOutput)
}

func (o ConsistencyPolicyResponseOutput) DefaultConsistencyLevel() pulumi.StringOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) string { return v.DefaultConsistencyLevel }).(pulumi.StringOutput)
}

func (o ConsistencyPolicyResponseOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) *int { return v.MaxIntervalInSeconds }).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyResponseOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ConsistencyPolicyResponse) *float64 { return v.MaxStalenessPrefix }).(pulumi.Float64PtrOutput)
}

type ConsistencyPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (ConsistencyPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConsistencyPolicyResponse)(nil)).Elem()
}

func (o ConsistencyPolicyResponsePtrOutput) ToConsistencyPolicyResponsePtrOutput() ConsistencyPolicyResponsePtrOutput {
	return o
}

func (o ConsistencyPolicyResponsePtrOutput) ToConsistencyPolicyResponsePtrOutputWithContext(ctx context.Context) ConsistencyPolicyResponsePtrOutput {
	return o
}

func (o ConsistencyPolicyResponsePtrOutput) Elem() ConsistencyPolicyResponseOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) ConsistencyPolicyResponse {
		if v != nil {
			return *v
		}
		var ret ConsistencyPolicyResponse
		return ret
	}).(ConsistencyPolicyResponseOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) DefaultConsistencyLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DefaultConsistencyLevel
	}).(pulumi.StringPtrOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) MaxIntervalInSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxIntervalInSeconds
	}).(pulumi.IntPtrOutput)
}

func (o ConsistencyPolicyResponsePtrOutput) MaxStalenessPrefix() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ConsistencyPolicyResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxStalenessPrefix
	}).(pulumi.Float64PtrOutput)
}

type ContainerPartitionKey struct {
	Kind    *string  `pulumi:"kind"`
	Paths   []string `pulumi:"paths"`
	Version *int     `pulumi:"version"`
}


func (val *ContainerPartitionKey) Defaults() *ContainerPartitionKey {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type ContainerPartitionKeyInput interface {
	pulumi.Input

	ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput
	ToContainerPartitionKeyOutputWithContext(context.Context) ContainerPartitionKeyOutput
}

type ContainerPartitionKeyArgs struct {
	Kind    pulumi.StringPtrInput   `pulumi:"kind"`
	Paths   pulumi.StringArrayInput `pulumi:"paths"`
	Version pulumi.IntPtrInput      `pulumi:"version"`
}

func (ContainerPartitionKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKey)(nil)).Elem()
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput {
	return i.ToContainerPartitionKeyOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyOutputWithContext(ctx context.Context) ContainerPartitionKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyOutput)
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return i.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyArgs) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyOutput).ToContainerPartitionKeyPtrOutputWithContext(ctx)
}









type ContainerPartitionKeyPtrInput interface {
	pulumi.Input

	ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput
	ToContainerPartitionKeyPtrOutputWithContext(context.Context) ContainerPartitionKeyPtrOutput
}

type containerPartitionKeyPtrType ContainerPartitionKeyArgs

func ContainerPartitionKeyPtr(v *ContainerPartitionKeyArgs) ContainerPartitionKeyPtrInput {
	return (*containerPartitionKeyPtrType)(v)
}

func (*containerPartitionKeyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKey)(nil)).Elem()
}

func (i *containerPartitionKeyPtrType) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return i.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (i *containerPartitionKeyPtrType) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyPtrOutput)
}

type ContainerPartitionKeyOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKey)(nil)).Elem()
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyOutput() ContainerPartitionKeyOutput {
	return o
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyOutputWithContext(ctx context.Context) ContainerPartitionKeyOutput {
	return o
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return o.ToContainerPartitionKeyPtrOutputWithContext(context.Background())
}

func (o ContainerPartitionKeyOutput) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerPartitionKey) *ContainerPartitionKey {
		return &v
	}).(ContainerPartitionKeyPtrOutput)
}

func (o ContainerPartitionKeyOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKey) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerPartitionKey) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKey) *int { return v.Version }).(pulumi.IntPtrOutput)
}

type ContainerPartitionKeyPtrOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKey)(nil)).Elem()
}

func (o ContainerPartitionKeyPtrOutput) ToContainerPartitionKeyPtrOutput() ContainerPartitionKeyPtrOutput {
	return o
}

func (o ContainerPartitionKeyPtrOutput) ToContainerPartitionKeyPtrOutputWithContext(ctx context.Context) ContainerPartitionKeyPtrOutput {
	return o
}

func (o ContainerPartitionKeyPtrOutput) Elem() ContainerPartitionKeyOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) ContainerPartitionKey {
		if v != nil {
			return *v
		}
		var ret ContainerPartitionKey
		return ret
	}).(ContainerPartitionKeyOutput)
}

func (o ContainerPartitionKeyPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyPtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyPtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKey) *int {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.IntPtrOutput)
}

type ContainerPartitionKeyResponse struct {
	Kind      *string  `pulumi:"kind"`
	Paths     []string `pulumi:"paths"`
	SystemKey bool     `pulumi:"systemKey"`
	Version   *int     `pulumi:"version"`
}


func (val *ContainerPartitionKeyResponse) Defaults() *ContainerPartitionKeyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type ContainerPartitionKeyResponseInput interface {
	pulumi.Input

	ToContainerPartitionKeyResponseOutput() ContainerPartitionKeyResponseOutput
	ToContainerPartitionKeyResponseOutputWithContext(context.Context) ContainerPartitionKeyResponseOutput
}

type ContainerPartitionKeyResponseArgs struct {
	Kind      pulumi.StringPtrInput   `pulumi:"kind"`
	Paths     pulumi.StringArrayInput `pulumi:"paths"`
	SystemKey pulumi.BoolInput        `pulumi:"systemKey"`
	Version   pulumi.IntPtrInput      `pulumi:"version"`
}

func (ContainerPartitionKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKeyResponse)(nil)).Elem()
}

func (i ContainerPartitionKeyResponseArgs) ToContainerPartitionKeyResponseOutput() ContainerPartitionKeyResponseOutput {
	return i.ToContainerPartitionKeyResponseOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyResponseArgs) ToContainerPartitionKeyResponseOutputWithContext(ctx context.Context) ContainerPartitionKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyResponseOutput)
}

func (i ContainerPartitionKeyResponseArgs) ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput {
	return i.ToContainerPartitionKeyResponsePtrOutputWithContext(context.Background())
}

func (i ContainerPartitionKeyResponseArgs) ToContainerPartitionKeyResponsePtrOutputWithContext(ctx context.Context) ContainerPartitionKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyResponseOutput).ToContainerPartitionKeyResponsePtrOutputWithContext(ctx)
}









type ContainerPartitionKeyResponsePtrInput interface {
	pulumi.Input

	ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput
	ToContainerPartitionKeyResponsePtrOutputWithContext(context.Context) ContainerPartitionKeyResponsePtrOutput
}

type containerPartitionKeyResponsePtrType ContainerPartitionKeyResponseArgs

func ContainerPartitionKeyResponsePtr(v *ContainerPartitionKeyResponseArgs) ContainerPartitionKeyResponsePtrInput {
	return (*containerPartitionKeyResponsePtrType)(v)
}

func (*containerPartitionKeyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKeyResponse)(nil)).Elem()
}

func (i *containerPartitionKeyResponsePtrType) ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput {
	return i.ToContainerPartitionKeyResponsePtrOutputWithContext(context.Background())
}

func (i *containerPartitionKeyResponsePtrType) ToContainerPartitionKeyResponsePtrOutputWithContext(ctx context.Context) ContainerPartitionKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerPartitionKeyResponsePtrOutput)
}

type ContainerPartitionKeyResponseOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerPartitionKeyResponse)(nil)).Elem()
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponseOutput() ContainerPartitionKeyResponseOutput {
	return o
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponseOutputWithContext(ctx context.Context) ContainerPartitionKeyResponseOutput {
	return o
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput {
	return o.ToContainerPartitionKeyResponsePtrOutputWithContext(context.Background())
}

func (o ContainerPartitionKeyResponseOutput) ToContainerPartitionKeyResponsePtrOutputWithContext(ctx context.Context) ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerPartitionKeyResponse) *ContainerPartitionKeyResponse {
		return &v
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o ContainerPartitionKeyResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyResponseOutput) SystemKey() pulumi.BoolOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) bool { return v.SystemKey }).(pulumi.BoolOutput)
}

func (o ContainerPartitionKeyResponseOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerPartitionKeyResponse) *int { return v.Version }).(pulumi.IntPtrOutput)
}

type ContainerPartitionKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerPartitionKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerPartitionKeyResponse)(nil)).Elem()
}

func (o ContainerPartitionKeyResponsePtrOutput) ToContainerPartitionKeyResponsePtrOutput() ContainerPartitionKeyResponsePtrOutput {
	return o
}

func (o ContainerPartitionKeyResponsePtrOutput) ToContainerPartitionKeyResponsePtrOutputWithContext(ctx context.Context) ContainerPartitionKeyResponsePtrOutput {
	return o
}

func (o ContainerPartitionKeyResponsePtrOutput) Elem() ContainerPartitionKeyResponseOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) ContainerPartitionKeyResponse {
		if v != nil {
			return *v
		}
		var ret ContainerPartitionKeyResponse
		return ret
	}).(ContainerPartitionKeyResponseOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) []string {
		if v == nil {
			return nil
		}
		return v.Paths
	}).(pulumi.StringArrayOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) SystemKey() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.SystemKey
	}).(pulumi.BoolPtrOutput)
}

func (o ContainerPartitionKeyResponsePtrOutput) Version() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ContainerPartitionKeyResponse) *int {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.IntPtrOutput)
}

type ContinuousModeBackupPolicy struct {
	MigrationState *BackupPolicyMigrationState `pulumi:"migrationState"`
	Type           string                      `pulumi:"type"`
}





type ContinuousModeBackupPolicyInput interface {
	pulumi.Input

	ToContinuousModeBackupPolicyOutput() ContinuousModeBackupPolicyOutput
	ToContinuousModeBackupPolicyOutputWithContext(context.Context) ContinuousModeBackupPolicyOutput
}

type ContinuousModeBackupPolicyArgs struct {
	MigrationState BackupPolicyMigrationStatePtrInput `pulumi:"migrationState"`
	Type           pulumi.StringInput                 `pulumi:"type"`
}

func (ContinuousModeBackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousModeBackupPolicy)(nil)).Elem()
}

func (i ContinuousModeBackupPolicyArgs) ToContinuousModeBackupPolicyOutput() ContinuousModeBackupPolicyOutput {
	return i.ToContinuousModeBackupPolicyOutputWithContext(context.Background())
}

func (i ContinuousModeBackupPolicyArgs) ToContinuousModeBackupPolicyOutputWithContext(ctx context.Context) ContinuousModeBackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousModeBackupPolicyOutput)
}

type ContinuousModeBackupPolicyOutput struct{ *pulumi.OutputState }

func (ContinuousModeBackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousModeBackupPolicy)(nil)).Elem()
}

func (o ContinuousModeBackupPolicyOutput) ToContinuousModeBackupPolicyOutput() ContinuousModeBackupPolicyOutput {
	return o
}

func (o ContinuousModeBackupPolicyOutput) ToContinuousModeBackupPolicyOutputWithContext(ctx context.Context) ContinuousModeBackupPolicyOutput {
	return o
}

func (o ContinuousModeBackupPolicyOutput) MigrationState() BackupPolicyMigrationStatePtrOutput {
	return o.ApplyT(func(v ContinuousModeBackupPolicy) *BackupPolicyMigrationState { return v.MigrationState }).(BackupPolicyMigrationStatePtrOutput)
}

func (o ContinuousModeBackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousModeBackupPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type ContinuousModeBackupPolicyResponse struct {
	MigrationState *BackupPolicyMigrationStateResponse `pulumi:"migrationState"`
	Type           string                              `pulumi:"type"`
}





type ContinuousModeBackupPolicyResponseInput interface {
	pulumi.Input

	ToContinuousModeBackupPolicyResponseOutput() ContinuousModeBackupPolicyResponseOutput
	ToContinuousModeBackupPolicyResponseOutputWithContext(context.Context) ContinuousModeBackupPolicyResponseOutput
}

type ContinuousModeBackupPolicyResponseArgs struct {
	MigrationState BackupPolicyMigrationStateResponsePtrInput `pulumi:"migrationState"`
	Type           pulumi.StringInput                         `pulumi:"type"`
}

func (ContinuousModeBackupPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousModeBackupPolicyResponse)(nil)).Elem()
}

func (i ContinuousModeBackupPolicyResponseArgs) ToContinuousModeBackupPolicyResponseOutput() ContinuousModeBackupPolicyResponseOutput {
	return i.ToContinuousModeBackupPolicyResponseOutputWithContext(context.Background())
}

func (i ContinuousModeBackupPolicyResponseArgs) ToContinuousModeBackupPolicyResponseOutputWithContext(ctx context.Context) ContinuousModeBackupPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContinuousModeBackupPolicyResponseOutput)
}

type ContinuousModeBackupPolicyResponseOutput struct{ *pulumi.OutputState }

func (ContinuousModeBackupPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContinuousModeBackupPolicyResponse)(nil)).Elem()
}

func (o ContinuousModeBackupPolicyResponseOutput) ToContinuousModeBackupPolicyResponseOutput() ContinuousModeBackupPolicyResponseOutput {
	return o
}

func (o ContinuousModeBackupPolicyResponseOutput) ToContinuousModeBackupPolicyResponseOutputWithContext(ctx context.Context) ContinuousModeBackupPolicyResponseOutput {
	return o
}

func (o ContinuousModeBackupPolicyResponseOutput) MigrationState() BackupPolicyMigrationStateResponsePtrOutput {
	return o.ApplyT(func(v ContinuousModeBackupPolicyResponse) *BackupPolicyMigrationStateResponse {
		return v.MigrationState
	}).(BackupPolicyMigrationStateResponsePtrOutput)
}

func (o ContinuousModeBackupPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ContinuousModeBackupPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type CorsPolicy struct {
	AllowedHeaders  *string  `pulumi:"allowedHeaders"`
	AllowedMethods  *string  `pulumi:"allowedMethods"`
	AllowedOrigins  string   `pulumi:"allowedOrigins"`
	ExposedHeaders  *string  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds *float64 `pulumi:"maxAgeInSeconds"`
}





type CorsPolicyInput interface {
	pulumi.Input

	ToCorsPolicyOutput() CorsPolicyOutput
	ToCorsPolicyOutputWithContext(context.Context) CorsPolicyOutput
}

type CorsPolicyArgs struct {
	AllowedHeaders  pulumi.StringPtrInput  `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringPtrInput  `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringInput     `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringPtrInput  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.Float64PtrInput `pulumi:"maxAgeInSeconds"`
}

func (CorsPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicy)(nil)).Elem()
}

func (i CorsPolicyArgs) ToCorsPolicyOutput() CorsPolicyOutput {
	return i.ToCorsPolicyOutputWithContext(context.Background())
}

func (i CorsPolicyArgs) ToCorsPolicyOutputWithContext(ctx context.Context) CorsPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsPolicyOutput)
}





type CorsPolicyArrayInput interface {
	pulumi.Input

	ToCorsPolicyArrayOutput() CorsPolicyArrayOutput
	ToCorsPolicyArrayOutputWithContext(context.Context) CorsPolicyArrayOutput
}

type CorsPolicyArray []CorsPolicyInput

func (CorsPolicyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicy)(nil)).Elem()
}

func (i CorsPolicyArray) ToCorsPolicyArrayOutput() CorsPolicyArrayOutput {
	return i.ToCorsPolicyArrayOutputWithContext(context.Background())
}

func (i CorsPolicyArray) ToCorsPolicyArrayOutputWithContext(ctx context.Context) CorsPolicyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsPolicyArrayOutput)
}

type CorsPolicyOutput struct{ *pulumi.OutputState }

func (CorsPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicy)(nil)).Elem()
}

func (o CorsPolicyOutput) ToCorsPolicyOutput() CorsPolicyOutput {
	return o
}

func (o CorsPolicyOutput) ToCorsPolicyOutputWithContext(ctx context.Context) CorsPolicyOutput {
	return o
}

func (o CorsPolicyOutput) AllowedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicy) *string { return v.AllowedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyOutput) AllowedMethods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicy) *string { return v.AllowedMethods }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyOutput) AllowedOrigins() pulumi.StringOutput {
	return o.ApplyT(func(v CorsPolicy) string { return v.AllowedOrigins }).(pulumi.StringOutput)
}

func (o CorsPolicyOutput) ExposedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicy) *string { return v.ExposedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyOutput) MaxAgeInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CorsPolicy) *float64 { return v.MaxAgeInSeconds }).(pulumi.Float64PtrOutput)
}

type CorsPolicyArrayOutput struct{ *pulumi.OutputState }

func (CorsPolicyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicy)(nil)).Elem()
}

func (o CorsPolicyArrayOutput) ToCorsPolicyArrayOutput() CorsPolicyArrayOutput {
	return o
}

func (o CorsPolicyArrayOutput) ToCorsPolicyArrayOutputWithContext(ctx context.Context) CorsPolicyArrayOutput {
	return o
}

func (o CorsPolicyArrayOutput) Index(i pulumi.IntInput) CorsPolicyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsPolicy {
		return vs[0].([]CorsPolicy)[vs[1].(int)]
	}).(CorsPolicyOutput)
}

type CorsPolicyResponse struct {
	AllowedHeaders  *string  `pulumi:"allowedHeaders"`
	AllowedMethods  *string  `pulumi:"allowedMethods"`
	AllowedOrigins  string   `pulumi:"allowedOrigins"`
	ExposedHeaders  *string  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds *float64 `pulumi:"maxAgeInSeconds"`
}





type CorsPolicyResponseInput interface {
	pulumi.Input

	ToCorsPolicyResponseOutput() CorsPolicyResponseOutput
	ToCorsPolicyResponseOutputWithContext(context.Context) CorsPolicyResponseOutput
}

type CorsPolicyResponseArgs struct {
	AllowedHeaders  pulumi.StringPtrInput  `pulumi:"allowedHeaders"`
	AllowedMethods  pulumi.StringPtrInput  `pulumi:"allowedMethods"`
	AllowedOrigins  pulumi.StringInput     `pulumi:"allowedOrigins"`
	ExposedHeaders  pulumi.StringPtrInput  `pulumi:"exposedHeaders"`
	MaxAgeInSeconds pulumi.Float64PtrInput `pulumi:"maxAgeInSeconds"`
}

func (CorsPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicyResponse)(nil)).Elem()
}

func (i CorsPolicyResponseArgs) ToCorsPolicyResponseOutput() CorsPolicyResponseOutput {
	return i.ToCorsPolicyResponseOutputWithContext(context.Background())
}

func (i CorsPolicyResponseArgs) ToCorsPolicyResponseOutputWithContext(ctx context.Context) CorsPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsPolicyResponseOutput)
}





type CorsPolicyResponseArrayInput interface {
	pulumi.Input

	ToCorsPolicyResponseArrayOutput() CorsPolicyResponseArrayOutput
	ToCorsPolicyResponseArrayOutputWithContext(context.Context) CorsPolicyResponseArrayOutput
}

type CorsPolicyResponseArray []CorsPolicyResponseInput

func (CorsPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicyResponse)(nil)).Elem()
}

func (i CorsPolicyResponseArray) ToCorsPolicyResponseArrayOutput() CorsPolicyResponseArrayOutput {
	return i.ToCorsPolicyResponseArrayOutputWithContext(context.Background())
}

func (i CorsPolicyResponseArray) ToCorsPolicyResponseArrayOutputWithContext(ctx context.Context) CorsPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsPolicyResponseArrayOutput)
}

type CorsPolicyResponseOutput struct{ *pulumi.OutputState }

func (CorsPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsPolicyResponse)(nil)).Elem()
}

func (o CorsPolicyResponseOutput) ToCorsPolicyResponseOutput() CorsPolicyResponseOutput {
	return o
}

func (o CorsPolicyResponseOutput) ToCorsPolicyResponseOutputWithContext(ctx context.Context) CorsPolicyResponseOutput {
	return o
}

func (o CorsPolicyResponseOutput) AllowedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *string { return v.AllowedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyResponseOutput) AllowedMethods() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *string { return v.AllowedMethods }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyResponseOutput) AllowedOrigins() pulumi.StringOutput {
	return o.ApplyT(func(v CorsPolicyResponse) string { return v.AllowedOrigins }).(pulumi.StringOutput)
}

func (o CorsPolicyResponseOutput) ExposedHeaders() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *string { return v.ExposedHeaders }).(pulumi.StringPtrOutput)
}

func (o CorsPolicyResponseOutput) MaxAgeInSeconds() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v CorsPolicyResponse) *float64 { return v.MaxAgeInSeconds }).(pulumi.Float64PtrOutput)
}

type CorsPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (CorsPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CorsPolicyResponse)(nil)).Elem()
}

func (o CorsPolicyResponseArrayOutput) ToCorsPolicyResponseArrayOutput() CorsPolicyResponseArrayOutput {
	return o
}

func (o CorsPolicyResponseArrayOutput) ToCorsPolicyResponseArrayOutputWithContext(ctx context.Context) CorsPolicyResponseArrayOutput {
	return o
}

func (o CorsPolicyResponseArrayOutput) Index(i pulumi.IntInput) CorsPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CorsPolicyResponse {
		return vs[0].([]CorsPolicyResponse)[vs[1].(int)]
	}).(CorsPolicyResponseOutput)
}

type CreateUpdateOptions struct {
	AutoscaleSettings *AutoscaleSettings `pulumi:"autoscaleSettings"`
	Throughput        *int               `pulumi:"throughput"`
}





type CreateUpdateOptionsInput interface {
	pulumi.Input

	ToCreateUpdateOptionsOutput() CreateUpdateOptionsOutput
	ToCreateUpdateOptionsOutputWithContext(context.Context) CreateUpdateOptionsOutput
}

type CreateUpdateOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsPtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput        `pulumi:"throughput"`
}

func (CreateUpdateOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateUpdateOptions)(nil)).Elem()
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsOutput() CreateUpdateOptionsOutput {
	return i.ToCreateUpdateOptionsOutputWithContext(context.Background())
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsOutputWithContext(ctx context.Context) CreateUpdateOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateUpdateOptionsOutput)
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return i.ToCreateUpdateOptionsPtrOutputWithContext(context.Background())
}

func (i CreateUpdateOptionsArgs) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateUpdateOptionsOutput).ToCreateUpdateOptionsPtrOutputWithContext(ctx)
}









type CreateUpdateOptionsPtrInput interface {
	pulumi.Input

	ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput
	ToCreateUpdateOptionsPtrOutputWithContext(context.Context) CreateUpdateOptionsPtrOutput
}

type createUpdateOptionsPtrType CreateUpdateOptionsArgs

func CreateUpdateOptionsPtr(v *CreateUpdateOptionsArgs) CreateUpdateOptionsPtrInput {
	return (*createUpdateOptionsPtrType)(v)
}

func (*createUpdateOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateUpdateOptions)(nil)).Elem()
}

func (i *createUpdateOptionsPtrType) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return i.ToCreateUpdateOptionsPtrOutputWithContext(context.Background())
}

func (i *createUpdateOptionsPtrType) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CreateUpdateOptionsPtrOutput)
}

type CreateUpdateOptionsOutput struct{ *pulumi.OutputState }

func (CreateUpdateOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CreateUpdateOptions)(nil)).Elem()
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsOutput() CreateUpdateOptionsOutput {
	return o
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsOutputWithContext(ctx context.Context) CreateUpdateOptionsOutput {
	return o
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return o.ToCreateUpdateOptionsPtrOutputWithContext(context.Background())
}

func (o CreateUpdateOptionsOutput) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CreateUpdateOptions) *CreateUpdateOptions {
		return &v
	}).(CreateUpdateOptionsPtrOutput)
}

func (o CreateUpdateOptionsOutput) AutoscaleSettings() AutoscaleSettingsPtrOutput {
	return o.ApplyT(func(v CreateUpdateOptions) *AutoscaleSettings { return v.AutoscaleSettings }).(AutoscaleSettingsPtrOutput)
}

func (o CreateUpdateOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v CreateUpdateOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type CreateUpdateOptionsPtrOutput struct{ *pulumi.OutputState }

func (CreateUpdateOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CreateUpdateOptions)(nil)).Elem()
}

func (o CreateUpdateOptionsPtrOutput) ToCreateUpdateOptionsPtrOutput() CreateUpdateOptionsPtrOutput {
	return o
}

func (o CreateUpdateOptionsPtrOutput) ToCreateUpdateOptionsPtrOutputWithContext(ctx context.Context) CreateUpdateOptionsPtrOutput {
	return o
}

func (o CreateUpdateOptionsPtrOutput) Elem() CreateUpdateOptionsOutput {
	return o.ApplyT(func(v *CreateUpdateOptions) CreateUpdateOptions {
		if v != nil {
			return *v
		}
		var ret CreateUpdateOptions
		return ret
	}).(CreateUpdateOptionsOutput)
}

func (o CreateUpdateOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsPtrOutput {
	return o.ApplyT(func(v *CreateUpdateOptions) *AutoscaleSettings {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsPtrOutput)
}

func (o CreateUpdateOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *CreateUpdateOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type DataCenterResourceProperties struct {
	AvailabilityZone                   *bool   `pulumi:"availabilityZone"`
	BackupStorageCustomerKeyUri        *string `pulumi:"backupStorageCustomerKeyUri"`
	Base64EncodedCassandraYamlFragment *string `pulumi:"base64EncodedCassandraYamlFragment"`
	DataCenterLocation                 *string `pulumi:"dataCenterLocation"`
	DelegatedSubnetId                  *string `pulumi:"delegatedSubnetId"`
	DiskCapacity                       *int    `pulumi:"diskCapacity"`
	DiskSku                            *string `pulumi:"diskSku"`
	ManagedDiskCustomerKeyUri          *string `pulumi:"managedDiskCustomerKeyUri"`
	NodeCount                          *int    `pulumi:"nodeCount"`
	ProvisioningState                  *string `pulumi:"provisioningState"`
	Sku                                *string `pulumi:"sku"`
}





type DataCenterResourcePropertiesInput interface {
	pulumi.Input

	ToDataCenterResourcePropertiesOutput() DataCenterResourcePropertiesOutput
	ToDataCenterResourcePropertiesOutputWithContext(context.Context) DataCenterResourcePropertiesOutput
}

type DataCenterResourcePropertiesArgs struct {
	AvailabilityZone                   pulumi.BoolPtrInput   `pulumi:"availabilityZone"`
	BackupStorageCustomerKeyUri        pulumi.StringPtrInput `pulumi:"backupStorageCustomerKeyUri"`
	Base64EncodedCassandraYamlFragment pulumi.StringPtrInput `pulumi:"base64EncodedCassandraYamlFragment"`
	DataCenterLocation                 pulumi.StringPtrInput `pulumi:"dataCenterLocation"`
	DelegatedSubnetId                  pulumi.StringPtrInput `pulumi:"delegatedSubnetId"`
	DiskCapacity                       pulumi.IntPtrInput    `pulumi:"diskCapacity"`
	DiskSku                            pulumi.StringPtrInput `pulumi:"diskSku"`
	ManagedDiskCustomerKeyUri          pulumi.StringPtrInput `pulumi:"managedDiskCustomerKeyUri"`
	NodeCount                          pulumi.IntPtrInput    `pulumi:"nodeCount"`
	ProvisioningState                  pulumi.StringPtrInput `pulumi:"provisioningState"`
	Sku                                pulumi.StringPtrInput `pulumi:"sku"`
}

func (DataCenterResourcePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCenterResourceProperties)(nil)).Elem()
}

func (i DataCenterResourcePropertiesArgs) ToDataCenterResourcePropertiesOutput() DataCenterResourcePropertiesOutput {
	return i.ToDataCenterResourcePropertiesOutputWithContext(context.Background())
}

func (i DataCenterResourcePropertiesArgs) ToDataCenterResourcePropertiesOutputWithContext(ctx context.Context) DataCenterResourcePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCenterResourcePropertiesOutput)
}

func (i DataCenterResourcePropertiesArgs) ToDataCenterResourcePropertiesPtrOutput() DataCenterResourcePropertiesPtrOutput {
	return i.ToDataCenterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i DataCenterResourcePropertiesArgs) ToDataCenterResourcePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCenterResourcePropertiesOutput).ToDataCenterResourcePropertiesPtrOutputWithContext(ctx)
}









type DataCenterResourcePropertiesPtrInput interface {
	pulumi.Input

	ToDataCenterResourcePropertiesPtrOutput() DataCenterResourcePropertiesPtrOutput
	ToDataCenterResourcePropertiesPtrOutputWithContext(context.Context) DataCenterResourcePropertiesPtrOutput
}

type dataCenterResourcePropertiesPtrType DataCenterResourcePropertiesArgs

func DataCenterResourcePropertiesPtr(v *DataCenterResourcePropertiesArgs) DataCenterResourcePropertiesPtrInput {
	return (*dataCenterResourcePropertiesPtrType)(v)
}

func (*dataCenterResourcePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCenterResourceProperties)(nil)).Elem()
}

func (i *dataCenterResourcePropertiesPtrType) ToDataCenterResourcePropertiesPtrOutput() DataCenterResourcePropertiesPtrOutput {
	return i.ToDataCenterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (i *dataCenterResourcePropertiesPtrType) ToDataCenterResourcePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourcePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCenterResourcePropertiesPtrOutput)
}

type DataCenterResourcePropertiesOutput struct{ *pulumi.OutputState }

func (DataCenterResourcePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCenterResourceProperties)(nil)).Elem()
}

func (o DataCenterResourcePropertiesOutput) ToDataCenterResourcePropertiesOutput() DataCenterResourcePropertiesOutput {
	return o
}

func (o DataCenterResourcePropertiesOutput) ToDataCenterResourcePropertiesOutputWithContext(ctx context.Context) DataCenterResourcePropertiesOutput {
	return o
}

func (o DataCenterResourcePropertiesOutput) ToDataCenterResourcePropertiesPtrOutput() DataCenterResourcePropertiesPtrOutput {
	return o.ToDataCenterResourcePropertiesPtrOutputWithContext(context.Background())
}

func (o DataCenterResourcePropertiesOutput) ToDataCenterResourcePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourcePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCenterResourceProperties) *DataCenterResourceProperties {
		return &v
	}).(DataCenterResourcePropertiesPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) AvailabilityZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *bool { return v.AvailabilityZone }).(pulumi.BoolPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) BackupStorageCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.BackupStorageCustomerKeyUri }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) Base64EncodedCassandraYamlFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.Base64EncodedCassandraYamlFragment }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) DataCenterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.DataCenterLocation }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) DelegatedSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.DelegatedSubnetId }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) DiskCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *int { return v.DiskCapacity }).(pulumi.IntPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) DiskSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.DiskSku }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) ManagedDiskCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.ManagedDiskCustomerKeyUri }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceProperties) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

type DataCenterResourcePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataCenterResourcePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCenterResourceProperties)(nil)).Elem()
}

func (o DataCenterResourcePropertiesPtrOutput) ToDataCenterResourcePropertiesPtrOutput() DataCenterResourcePropertiesPtrOutput {
	return o
}

func (o DataCenterResourcePropertiesPtrOutput) ToDataCenterResourcePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourcePropertiesPtrOutput {
	return o
}

func (o DataCenterResourcePropertiesPtrOutput) Elem() DataCenterResourcePropertiesOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) DataCenterResourceProperties {
		if v != nil {
			return *v
		}
		var ret DataCenterResourceProperties
		return ret
	}).(DataCenterResourcePropertiesOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) AvailabilityZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AvailabilityZone
	}).(pulumi.BoolPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) BackupStorageCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.BackupStorageCustomerKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) Base64EncodedCassandraYamlFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Base64EncodedCassandraYamlFragment
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) DataCenterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataCenterLocation
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) DelegatedSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) DiskCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *int {
		if v == nil {
			return nil
		}
		return v.DiskCapacity
	}).(pulumi.IntPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) DiskSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiskSku
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) ManagedDiskCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagedDiskCustomerKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourcePropertiesPtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceProperties) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

type DataCenterResourceResponseProperties struct {
	AvailabilityZone                   *bool              `pulumi:"availabilityZone"`
	BackupStorageCustomerKeyUri        *string            `pulumi:"backupStorageCustomerKeyUri"`
	Base64EncodedCassandraYamlFragment *string            `pulumi:"base64EncodedCassandraYamlFragment"`
	DataCenterLocation                 *string            `pulumi:"dataCenterLocation"`
	DelegatedSubnetId                  *string            `pulumi:"delegatedSubnetId"`
	DiskCapacity                       *int               `pulumi:"diskCapacity"`
	DiskSku                            *string            `pulumi:"diskSku"`
	ManagedDiskCustomerKeyUri          *string            `pulumi:"managedDiskCustomerKeyUri"`
	NodeCount                          *int               `pulumi:"nodeCount"`
	ProvisioningState                  *string            `pulumi:"provisioningState"`
	SeedNodes                          []SeedNodeResponse `pulumi:"seedNodes"`
	Sku                                *string            `pulumi:"sku"`
}





type DataCenterResourceResponsePropertiesInput interface {
	pulumi.Input

	ToDataCenterResourceResponsePropertiesOutput() DataCenterResourceResponsePropertiesOutput
	ToDataCenterResourceResponsePropertiesOutputWithContext(context.Context) DataCenterResourceResponsePropertiesOutput
}

type DataCenterResourceResponsePropertiesArgs struct {
	AvailabilityZone                   pulumi.BoolPtrInput        `pulumi:"availabilityZone"`
	BackupStorageCustomerKeyUri        pulumi.StringPtrInput      `pulumi:"backupStorageCustomerKeyUri"`
	Base64EncodedCassandraYamlFragment pulumi.StringPtrInput      `pulumi:"base64EncodedCassandraYamlFragment"`
	DataCenterLocation                 pulumi.StringPtrInput      `pulumi:"dataCenterLocation"`
	DelegatedSubnetId                  pulumi.StringPtrInput      `pulumi:"delegatedSubnetId"`
	DiskCapacity                       pulumi.IntPtrInput         `pulumi:"diskCapacity"`
	DiskSku                            pulumi.StringPtrInput      `pulumi:"diskSku"`
	ManagedDiskCustomerKeyUri          pulumi.StringPtrInput      `pulumi:"managedDiskCustomerKeyUri"`
	NodeCount                          pulumi.IntPtrInput         `pulumi:"nodeCount"`
	ProvisioningState                  pulumi.StringPtrInput      `pulumi:"provisioningState"`
	SeedNodes                          SeedNodeResponseArrayInput `pulumi:"seedNodes"`
	Sku                                pulumi.StringPtrInput      `pulumi:"sku"`
}

func (DataCenterResourceResponsePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCenterResourceResponseProperties)(nil)).Elem()
}

func (i DataCenterResourceResponsePropertiesArgs) ToDataCenterResourceResponsePropertiesOutput() DataCenterResourceResponsePropertiesOutput {
	return i.ToDataCenterResourceResponsePropertiesOutputWithContext(context.Background())
}

func (i DataCenterResourceResponsePropertiesArgs) ToDataCenterResourceResponsePropertiesOutputWithContext(ctx context.Context) DataCenterResourceResponsePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCenterResourceResponsePropertiesOutput)
}

func (i DataCenterResourceResponsePropertiesArgs) ToDataCenterResourceResponsePropertiesPtrOutput() DataCenterResourceResponsePropertiesPtrOutput {
	return i.ToDataCenterResourceResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i DataCenterResourceResponsePropertiesArgs) ToDataCenterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourceResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCenterResourceResponsePropertiesOutput).ToDataCenterResourceResponsePropertiesPtrOutputWithContext(ctx)
}









type DataCenterResourceResponsePropertiesPtrInput interface {
	pulumi.Input

	ToDataCenterResourceResponsePropertiesPtrOutput() DataCenterResourceResponsePropertiesPtrOutput
	ToDataCenterResourceResponsePropertiesPtrOutputWithContext(context.Context) DataCenterResourceResponsePropertiesPtrOutput
}

type dataCenterResourceResponsePropertiesPtrType DataCenterResourceResponsePropertiesArgs

func DataCenterResourceResponsePropertiesPtr(v *DataCenterResourceResponsePropertiesArgs) DataCenterResourceResponsePropertiesPtrInput {
	return (*dataCenterResourceResponsePropertiesPtrType)(v)
}

func (*dataCenterResourceResponsePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCenterResourceResponseProperties)(nil)).Elem()
}

func (i *dataCenterResourceResponsePropertiesPtrType) ToDataCenterResourceResponsePropertiesPtrOutput() DataCenterResourceResponsePropertiesPtrOutput {
	return i.ToDataCenterResourceResponsePropertiesPtrOutputWithContext(context.Background())
}

func (i *dataCenterResourceResponsePropertiesPtrType) ToDataCenterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourceResponsePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DataCenterResourceResponsePropertiesPtrOutput)
}

type DataCenterResourceResponsePropertiesOutput struct{ *pulumi.OutputState }

func (DataCenterResourceResponsePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DataCenterResourceResponseProperties)(nil)).Elem()
}

func (o DataCenterResourceResponsePropertiesOutput) ToDataCenterResourceResponsePropertiesOutput() DataCenterResourceResponsePropertiesOutput {
	return o
}

func (o DataCenterResourceResponsePropertiesOutput) ToDataCenterResourceResponsePropertiesOutputWithContext(ctx context.Context) DataCenterResourceResponsePropertiesOutput {
	return o
}

func (o DataCenterResourceResponsePropertiesOutput) ToDataCenterResourceResponsePropertiesPtrOutput() DataCenterResourceResponsePropertiesPtrOutput {
	return o.ToDataCenterResourceResponsePropertiesPtrOutputWithContext(context.Background())
}

func (o DataCenterResourceResponsePropertiesOutput) ToDataCenterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourceResponsePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DataCenterResourceResponseProperties) *DataCenterResourceResponseProperties {
		return &v
	}).(DataCenterResourceResponsePropertiesPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) AvailabilityZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *bool { return v.AvailabilityZone }).(pulumi.BoolPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) BackupStorageCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.BackupStorageCustomerKeyUri }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) Base64EncodedCassandraYamlFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.Base64EncodedCassandraYamlFragment }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) DataCenterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.DataCenterLocation }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) DelegatedSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.DelegatedSubnetId }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) DiskCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *int { return v.DiskCapacity }).(pulumi.IntPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) DiskSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.DiskSku }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) ManagedDiskCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.ManagedDiskCustomerKeyUri }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *int { return v.NodeCount }).(pulumi.IntPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) SeedNodes() SeedNodeResponseArrayOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) []SeedNodeResponse { return v.SeedNodes }).(SeedNodeResponseArrayOutput)
}

func (o DataCenterResourceResponsePropertiesOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DataCenterResourceResponseProperties) *string { return v.Sku }).(pulumi.StringPtrOutput)
}

type DataCenterResourceResponsePropertiesPtrOutput struct{ *pulumi.OutputState }

func (DataCenterResourceResponsePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DataCenterResourceResponseProperties)(nil)).Elem()
}

func (o DataCenterResourceResponsePropertiesPtrOutput) ToDataCenterResourceResponsePropertiesPtrOutput() DataCenterResourceResponsePropertiesPtrOutput {
	return o
}

func (o DataCenterResourceResponsePropertiesPtrOutput) ToDataCenterResourceResponsePropertiesPtrOutputWithContext(ctx context.Context) DataCenterResourceResponsePropertiesPtrOutput {
	return o
}

func (o DataCenterResourceResponsePropertiesPtrOutput) Elem() DataCenterResourceResponsePropertiesOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) DataCenterResourceResponseProperties {
		if v != nil {
			return *v
		}
		var ret DataCenterResourceResponseProperties
		return ret
	}).(DataCenterResourceResponsePropertiesOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) AvailabilityZone() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *bool {
		if v == nil {
			return nil
		}
		return v.AvailabilityZone
	}).(pulumi.BoolPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) BackupStorageCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.BackupStorageCustomerKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) Base64EncodedCassandraYamlFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Base64EncodedCassandraYamlFragment
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) DataCenterLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DataCenterLocation
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) DelegatedSubnetId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DelegatedSubnetId
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) DiskCapacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.DiskCapacity
	}).(pulumi.IntPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) DiskSku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.DiskSku
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) ManagedDiskCustomerKeyUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManagedDiskCustomerKeyUri
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) NodeCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *int {
		if v == nil {
			return nil
		}
		return v.NodeCount
	}).(pulumi.IntPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) SeedNodes() SeedNodeResponseArrayOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) []SeedNodeResponse {
		if v == nil {
			return nil
		}
		return v.SeedNodes
	}).(SeedNodeResponseArrayOutput)
}

func (o DataCenterResourceResponsePropertiesPtrOutput) Sku() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DataCenterResourceResponseProperties) *string {
		if v == nil {
			return nil
		}
		return v.Sku
	}).(pulumi.StringPtrOutput)
}

type DatabaseAccountConnectionStringResponse struct {
	ConnectionString string `pulumi:"connectionString"`
	Description      string `pulumi:"description"`
}





type DatabaseAccountConnectionStringResponseInput interface {
	pulumi.Input

	ToDatabaseAccountConnectionStringResponseOutput() DatabaseAccountConnectionStringResponseOutput
	ToDatabaseAccountConnectionStringResponseOutputWithContext(context.Context) DatabaseAccountConnectionStringResponseOutput
}

type DatabaseAccountConnectionStringResponseArgs struct {
	ConnectionString pulumi.StringInput `pulumi:"connectionString"`
	Description      pulumi.StringInput `pulumi:"description"`
}

func (DatabaseAccountConnectionStringResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountConnectionStringResponse)(nil)).Elem()
}

func (i DatabaseAccountConnectionStringResponseArgs) ToDatabaseAccountConnectionStringResponseOutput() DatabaseAccountConnectionStringResponseOutput {
	return i.ToDatabaseAccountConnectionStringResponseOutputWithContext(context.Background())
}

func (i DatabaseAccountConnectionStringResponseArgs) ToDatabaseAccountConnectionStringResponseOutputWithContext(ctx context.Context) DatabaseAccountConnectionStringResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountConnectionStringResponseOutput)
}





type DatabaseAccountConnectionStringResponseArrayInput interface {
	pulumi.Input

	ToDatabaseAccountConnectionStringResponseArrayOutput() DatabaseAccountConnectionStringResponseArrayOutput
	ToDatabaseAccountConnectionStringResponseArrayOutputWithContext(context.Context) DatabaseAccountConnectionStringResponseArrayOutput
}

type DatabaseAccountConnectionStringResponseArray []DatabaseAccountConnectionStringResponseInput

func (DatabaseAccountConnectionStringResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseAccountConnectionStringResponse)(nil)).Elem()
}

func (i DatabaseAccountConnectionStringResponseArray) ToDatabaseAccountConnectionStringResponseArrayOutput() DatabaseAccountConnectionStringResponseArrayOutput {
	return i.ToDatabaseAccountConnectionStringResponseArrayOutputWithContext(context.Background())
}

func (i DatabaseAccountConnectionStringResponseArray) ToDatabaseAccountConnectionStringResponseArrayOutputWithContext(ctx context.Context) DatabaseAccountConnectionStringResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseAccountConnectionStringResponseArrayOutput)
}

type DatabaseAccountConnectionStringResponseOutput struct{ *pulumi.OutputState }

func (DatabaseAccountConnectionStringResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseAccountConnectionStringResponse)(nil)).Elem()
}

func (o DatabaseAccountConnectionStringResponseOutput) ToDatabaseAccountConnectionStringResponseOutput() DatabaseAccountConnectionStringResponseOutput {
	return o
}

func (o DatabaseAccountConnectionStringResponseOutput) ToDatabaseAccountConnectionStringResponseOutputWithContext(ctx context.Context) DatabaseAccountConnectionStringResponseOutput {
	return o
}

func (o DatabaseAccountConnectionStringResponseOutput) ConnectionString() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseAccountConnectionStringResponse) string { return v.ConnectionString }).(pulumi.StringOutput)
}

func (o DatabaseAccountConnectionStringResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseAccountConnectionStringResponse) string { return v.Description }).(pulumi.StringOutput)
}

type DatabaseAccountConnectionStringResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseAccountConnectionStringResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseAccountConnectionStringResponse)(nil)).Elem()
}

func (o DatabaseAccountConnectionStringResponseArrayOutput) ToDatabaseAccountConnectionStringResponseArrayOutput() DatabaseAccountConnectionStringResponseArrayOutput {
	return o
}

func (o DatabaseAccountConnectionStringResponseArrayOutput) ToDatabaseAccountConnectionStringResponseArrayOutputWithContext(ctx context.Context) DatabaseAccountConnectionStringResponseArrayOutput {
	return o
}

func (o DatabaseAccountConnectionStringResponseArrayOutput) Index(i pulumi.IntInput) DatabaseAccountConnectionStringResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseAccountConnectionStringResponse {
		return vs[0].([]DatabaseAccountConnectionStringResponse)[vs[1].(int)]
	}).(DatabaseAccountConnectionStringResponseOutput)
}

type DatabaseRestoreResource struct {
	CollectionNames []string `pulumi:"collectionNames"`
	DatabaseName    *string  `pulumi:"databaseName"`
}





type DatabaseRestoreResourceInput interface {
	pulumi.Input

	ToDatabaseRestoreResourceOutput() DatabaseRestoreResourceOutput
	ToDatabaseRestoreResourceOutputWithContext(context.Context) DatabaseRestoreResourceOutput
}

type DatabaseRestoreResourceArgs struct {
	CollectionNames pulumi.StringArrayInput `pulumi:"collectionNames"`
	DatabaseName    pulumi.StringPtrInput   `pulumi:"databaseName"`
}

func (DatabaseRestoreResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseRestoreResource)(nil)).Elem()
}

func (i DatabaseRestoreResourceArgs) ToDatabaseRestoreResourceOutput() DatabaseRestoreResourceOutput {
	return i.ToDatabaseRestoreResourceOutputWithContext(context.Background())
}

func (i DatabaseRestoreResourceArgs) ToDatabaseRestoreResourceOutputWithContext(ctx context.Context) DatabaseRestoreResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRestoreResourceOutput)
}





type DatabaseRestoreResourceArrayInput interface {
	pulumi.Input

	ToDatabaseRestoreResourceArrayOutput() DatabaseRestoreResourceArrayOutput
	ToDatabaseRestoreResourceArrayOutputWithContext(context.Context) DatabaseRestoreResourceArrayOutput
}

type DatabaseRestoreResourceArray []DatabaseRestoreResourceInput

func (DatabaseRestoreResourceArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseRestoreResource)(nil)).Elem()
}

func (i DatabaseRestoreResourceArray) ToDatabaseRestoreResourceArrayOutput() DatabaseRestoreResourceArrayOutput {
	return i.ToDatabaseRestoreResourceArrayOutputWithContext(context.Background())
}

func (i DatabaseRestoreResourceArray) ToDatabaseRestoreResourceArrayOutputWithContext(ctx context.Context) DatabaseRestoreResourceArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRestoreResourceArrayOutput)
}

type DatabaseRestoreResourceOutput struct{ *pulumi.OutputState }

func (DatabaseRestoreResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseRestoreResource)(nil)).Elem()
}

func (o DatabaseRestoreResourceOutput) ToDatabaseRestoreResourceOutput() DatabaseRestoreResourceOutput {
	return o
}

func (o DatabaseRestoreResourceOutput) ToDatabaseRestoreResourceOutputWithContext(ctx context.Context) DatabaseRestoreResourceOutput {
	return o
}

func (o DatabaseRestoreResourceOutput) CollectionNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseRestoreResource) []string { return v.CollectionNames }).(pulumi.StringArrayOutput)
}

func (o DatabaseRestoreResourceOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseRestoreResource) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

type DatabaseRestoreResourceArrayOutput struct{ *pulumi.OutputState }

func (DatabaseRestoreResourceArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseRestoreResource)(nil)).Elem()
}

func (o DatabaseRestoreResourceArrayOutput) ToDatabaseRestoreResourceArrayOutput() DatabaseRestoreResourceArrayOutput {
	return o
}

func (o DatabaseRestoreResourceArrayOutput) ToDatabaseRestoreResourceArrayOutputWithContext(ctx context.Context) DatabaseRestoreResourceArrayOutput {
	return o
}

func (o DatabaseRestoreResourceArrayOutput) Index(i pulumi.IntInput) DatabaseRestoreResourceOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseRestoreResource {
		return vs[0].([]DatabaseRestoreResource)[vs[1].(int)]
	}).(DatabaseRestoreResourceOutput)
}

type DatabaseRestoreResourceResponse struct {
	CollectionNames []string `pulumi:"collectionNames"`
	DatabaseName    *string  `pulumi:"databaseName"`
}





type DatabaseRestoreResourceResponseInput interface {
	pulumi.Input

	ToDatabaseRestoreResourceResponseOutput() DatabaseRestoreResourceResponseOutput
	ToDatabaseRestoreResourceResponseOutputWithContext(context.Context) DatabaseRestoreResourceResponseOutput
}

type DatabaseRestoreResourceResponseArgs struct {
	CollectionNames pulumi.StringArrayInput `pulumi:"collectionNames"`
	DatabaseName    pulumi.StringPtrInput   `pulumi:"databaseName"`
}

func (DatabaseRestoreResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseRestoreResourceResponse)(nil)).Elem()
}

func (i DatabaseRestoreResourceResponseArgs) ToDatabaseRestoreResourceResponseOutput() DatabaseRestoreResourceResponseOutput {
	return i.ToDatabaseRestoreResourceResponseOutputWithContext(context.Background())
}

func (i DatabaseRestoreResourceResponseArgs) ToDatabaseRestoreResourceResponseOutputWithContext(ctx context.Context) DatabaseRestoreResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRestoreResourceResponseOutput)
}





type DatabaseRestoreResourceResponseArrayInput interface {
	pulumi.Input

	ToDatabaseRestoreResourceResponseArrayOutput() DatabaseRestoreResourceResponseArrayOutput
	ToDatabaseRestoreResourceResponseArrayOutputWithContext(context.Context) DatabaseRestoreResourceResponseArrayOutput
}

type DatabaseRestoreResourceResponseArray []DatabaseRestoreResourceResponseInput

func (DatabaseRestoreResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseRestoreResourceResponse)(nil)).Elem()
}

func (i DatabaseRestoreResourceResponseArray) ToDatabaseRestoreResourceResponseArrayOutput() DatabaseRestoreResourceResponseArrayOutput {
	return i.ToDatabaseRestoreResourceResponseArrayOutputWithContext(context.Background())
}

func (i DatabaseRestoreResourceResponseArray) ToDatabaseRestoreResourceResponseArrayOutputWithContext(ctx context.Context) DatabaseRestoreResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseRestoreResourceResponseArrayOutput)
}

type DatabaseRestoreResourceResponseOutput struct{ *pulumi.OutputState }

func (DatabaseRestoreResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseRestoreResourceResponse)(nil)).Elem()
}

func (o DatabaseRestoreResourceResponseOutput) ToDatabaseRestoreResourceResponseOutput() DatabaseRestoreResourceResponseOutput {
	return o
}

func (o DatabaseRestoreResourceResponseOutput) ToDatabaseRestoreResourceResponseOutputWithContext(ctx context.Context) DatabaseRestoreResourceResponseOutput {
	return o
}

func (o DatabaseRestoreResourceResponseOutput) CollectionNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DatabaseRestoreResourceResponse) []string { return v.CollectionNames }).(pulumi.StringArrayOutput)
}

func (o DatabaseRestoreResourceResponseOutput) DatabaseName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseRestoreResourceResponse) *string { return v.DatabaseName }).(pulumi.StringPtrOutput)
}

type DatabaseRestoreResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseRestoreResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseRestoreResourceResponse)(nil)).Elem()
}

func (o DatabaseRestoreResourceResponseArrayOutput) ToDatabaseRestoreResourceResponseArrayOutput() DatabaseRestoreResourceResponseArrayOutput {
	return o
}

func (o DatabaseRestoreResourceResponseArrayOutput) ToDatabaseRestoreResourceResponseArrayOutputWithContext(ctx context.Context) DatabaseRestoreResourceResponseArrayOutput {
	return o
}

func (o DatabaseRestoreResourceResponseArrayOutput) Index(i pulumi.IntInput) DatabaseRestoreResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseRestoreResourceResponse {
		return vs[0].([]DatabaseRestoreResourceResponse)[vs[1].(int)]
	}).(DatabaseRestoreResourceResponseOutput)
}

type ExcludedPath struct {
	Path *string `pulumi:"path"`
}





type ExcludedPathInput interface {
	pulumi.Input

	ToExcludedPathOutput() ExcludedPathOutput
	ToExcludedPathOutputWithContext(context.Context) ExcludedPathOutput
}

type ExcludedPathArgs struct {
	Path pulumi.StringPtrInput `pulumi:"path"`
}

func (ExcludedPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPath)(nil)).Elem()
}

func (i ExcludedPathArgs) ToExcludedPathOutput() ExcludedPathOutput {
	return i.ToExcludedPathOutputWithContext(context.Background())
}

func (i ExcludedPathArgs) ToExcludedPathOutputWithContext(ctx context.Context) ExcludedPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathOutput)
}





type ExcludedPathArrayInput interface {
	pulumi.Input

	ToExcludedPathArrayOutput() ExcludedPathArrayOutput
	ToExcludedPathArrayOutputWithContext(context.Context) ExcludedPathArrayOutput
}

type ExcludedPathArray []ExcludedPathInput

func (ExcludedPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPath)(nil)).Elem()
}

func (i ExcludedPathArray) ToExcludedPathArrayOutput() ExcludedPathArrayOutput {
	return i.ToExcludedPathArrayOutputWithContext(context.Background())
}

func (i ExcludedPathArray) ToExcludedPathArrayOutputWithContext(ctx context.Context) ExcludedPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathArrayOutput)
}

type ExcludedPathOutput struct{ *pulumi.OutputState }

func (ExcludedPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPath)(nil)).Elem()
}

func (o ExcludedPathOutput) ToExcludedPathOutput() ExcludedPathOutput {
	return o
}

func (o ExcludedPathOutput) ToExcludedPathOutputWithContext(ctx context.Context) ExcludedPathOutput {
	return o
}

func (o ExcludedPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ExcludedPathArrayOutput struct{ *pulumi.OutputState }

func (ExcludedPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPath)(nil)).Elem()
}

func (o ExcludedPathArrayOutput) ToExcludedPathArrayOutput() ExcludedPathArrayOutput {
	return o
}

func (o ExcludedPathArrayOutput) ToExcludedPathArrayOutputWithContext(ctx context.Context) ExcludedPathArrayOutput {
	return o
}

func (o ExcludedPathArrayOutput) Index(i pulumi.IntInput) ExcludedPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExcludedPath {
		return vs[0].([]ExcludedPath)[vs[1].(int)]
	}).(ExcludedPathOutput)
}

type ExcludedPathResponse struct {
	Path *string `pulumi:"path"`
}





type ExcludedPathResponseInput interface {
	pulumi.Input

	ToExcludedPathResponseOutput() ExcludedPathResponseOutput
	ToExcludedPathResponseOutputWithContext(context.Context) ExcludedPathResponseOutput
}

type ExcludedPathResponseArgs struct {
	Path pulumi.StringPtrInput `pulumi:"path"`
}

func (ExcludedPathResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPathResponse)(nil)).Elem()
}

func (i ExcludedPathResponseArgs) ToExcludedPathResponseOutput() ExcludedPathResponseOutput {
	return i.ToExcludedPathResponseOutputWithContext(context.Background())
}

func (i ExcludedPathResponseArgs) ToExcludedPathResponseOutputWithContext(ctx context.Context) ExcludedPathResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathResponseOutput)
}





type ExcludedPathResponseArrayInput interface {
	pulumi.Input

	ToExcludedPathResponseArrayOutput() ExcludedPathResponseArrayOutput
	ToExcludedPathResponseArrayOutputWithContext(context.Context) ExcludedPathResponseArrayOutput
}

type ExcludedPathResponseArray []ExcludedPathResponseInput

func (ExcludedPathResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPathResponse)(nil)).Elem()
}

func (i ExcludedPathResponseArray) ToExcludedPathResponseArrayOutput() ExcludedPathResponseArrayOutput {
	return i.ToExcludedPathResponseArrayOutputWithContext(context.Background())
}

func (i ExcludedPathResponseArray) ToExcludedPathResponseArrayOutputWithContext(ctx context.Context) ExcludedPathResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExcludedPathResponseArrayOutput)
}

type ExcludedPathResponseOutput struct{ *pulumi.OutputState }

func (ExcludedPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExcludedPathResponse)(nil)).Elem()
}

func (o ExcludedPathResponseOutput) ToExcludedPathResponseOutput() ExcludedPathResponseOutput {
	return o
}

func (o ExcludedPathResponseOutput) ToExcludedPathResponseOutputWithContext(ctx context.Context) ExcludedPathResponseOutput {
	return o
}

func (o ExcludedPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExcludedPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type ExcludedPathResponseArrayOutput struct{ *pulumi.OutputState }

func (ExcludedPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExcludedPathResponse)(nil)).Elem()
}

func (o ExcludedPathResponseArrayOutput) ToExcludedPathResponseArrayOutput() ExcludedPathResponseArrayOutput {
	return o
}

func (o ExcludedPathResponseArrayOutput) ToExcludedPathResponseArrayOutputWithContext(ctx context.Context) ExcludedPathResponseArrayOutput {
	return o
}

func (o ExcludedPathResponseArrayOutput) Index(i pulumi.IntInput) ExcludedPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExcludedPathResponse {
		return vs[0].([]ExcludedPathResponse)[vs[1].(int)]
	}).(ExcludedPathResponseOutput)
}

type FailoverPolicyResponse struct {
	FailoverPriority *int    `pulumi:"failoverPriority"`
	Id               string  `pulumi:"id"`
	LocationName     *string `pulumi:"locationName"`
}





type FailoverPolicyResponseInput interface {
	pulumi.Input

	ToFailoverPolicyResponseOutput() FailoverPolicyResponseOutput
	ToFailoverPolicyResponseOutputWithContext(context.Context) FailoverPolicyResponseOutput
}

type FailoverPolicyResponseArgs struct {
	FailoverPriority pulumi.IntPtrInput    `pulumi:"failoverPriority"`
	Id               pulumi.StringInput    `pulumi:"id"`
	LocationName     pulumi.StringPtrInput `pulumi:"locationName"`
}

func (FailoverPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverPolicyResponse)(nil)).Elem()
}

func (i FailoverPolicyResponseArgs) ToFailoverPolicyResponseOutput() FailoverPolicyResponseOutput {
	return i.ToFailoverPolicyResponseOutputWithContext(context.Background())
}

func (i FailoverPolicyResponseArgs) ToFailoverPolicyResponseOutputWithContext(ctx context.Context) FailoverPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverPolicyResponseOutput)
}





type FailoverPolicyResponseArrayInput interface {
	pulumi.Input

	ToFailoverPolicyResponseArrayOutput() FailoverPolicyResponseArrayOutput
	ToFailoverPolicyResponseArrayOutputWithContext(context.Context) FailoverPolicyResponseArrayOutput
}

type FailoverPolicyResponseArray []FailoverPolicyResponseInput

func (FailoverPolicyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverPolicyResponse)(nil)).Elem()
}

func (i FailoverPolicyResponseArray) ToFailoverPolicyResponseArrayOutput() FailoverPolicyResponseArrayOutput {
	return i.ToFailoverPolicyResponseArrayOutputWithContext(context.Background())
}

func (i FailoverPolicyResponseArray) ToFailoverPolicyResponseArrayOutputWithContext(ctx context.Context) FailoverPolicyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FailoverPolicyResponseArrayOutput)
}

type FailoverPolicyResponseOutput struct{ *pulumi.OutputState }

func (FailoverPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FailoverPolicyResponse)(nil)).Elem()
}

func (o FailoverPolicyResponseOutput) ToFailoverPolicyResponseOutput() FailoverPolicyResponseOutput {
	return o
}

func (o FailoverPolicyResponseOutput) ToFailoverPolicyResponseOutputWithContext(ctx context.Context) FailoverPolicyResponseOutput {
	return o
}

func (o FailoverPolicyResponseOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o FailoverPolicyResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o FailoverPolicyResponseOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FailoverPolicyResponse) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

type FailoverPolicyResponseArrayOutput struct{ *pulumi.OutputState }

func (FailoverPolicyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FailoverPolicyResponse)(nil)).Elem()
}

func (o FailoverPolicyResponseArrayOutput) ToFailoverPolicyResponseArrayOutput() FailoverPolicyResponseArrayOutput {
	return o
}

func (o FailoverPolicyResponseArrayOutput) ToFailoverPolicyResponseArrayOutputWithContext(ctx context.Context) FailoverPolicyResponseArrayOutput {
	return o
}

func (o FailoverPolicyResponseArrayOutput) Index(i pulumi.IntInput) FailoverPolicyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FailoverPolicyResponse {
		return vs[0].([]FailoverPolicyResponse)[vs[1].(int)]
	}).(FailoverPolicyResponseOutput)
}

type GremlinDatabaseGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type GremlinDatabaseGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToGremlinDatabaseGetPropertiesResponseOptionsOutput() GremlinDatabaseGetPropertiesResponseOptionsOutput
	ToGremlinDatabaseGetPropertiesResponseOptionsOutputWithContext(context.Context) GremlinDatabaseGetPropertiesResponseOptionsOutput
}

type GremlinDatabaseGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (GremlinDatabaseGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (i GremlinDatabaseGetPropertiesResponseOptionsArgs) ToGremlinDatabaseGetPropertiesResponseOptionsOutput() GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return i.ToGremlinDatabaseGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i GremlinDatabaseGetPropertiesResponseOptionsArgs) ToGremlinDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseGetPropertiesResponseOptionsOutput)
}

func (i GremlinDatabaseGetPropertiesResponseOptionsArgs) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutput() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return i.ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i GremlinDatabaseGetPropertiesResponseOptionsArgs) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseGetPropertiesResponseOptionsOutput).ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type GremlinDatabaseGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutput() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput
	ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) GremlinDatabaseGetPropertiesResponseOptionsPtrOutput
}

type gremlinDatabaseGetPropertiesResponseOptionsPtrType GremlinDatabaseGetPropertiesResponseOptionsArgs

func GremlinDatabaseGetPropertiesResponseOptionsPtr(v *GremlinDatabaseGetPropertiesResponseOptionsArgs) GremlinDatabaseGetPropertiesResponseOptionsPtrInput {
	return (*gremlinDatabaseGetPropertiesResponseOptionsPtrType)(v)
}

func (*gremlinDatabaseGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *gremlinDatabaseGetPropertiesResponseOptionsPtrType) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutput() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return i.ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *gremlinDatabaseGetPropertiesResponseOptionsPtrType) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseGetPropertiesResponseOptionsPtrOutput)
}

type GremlinDatabaseGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) ToGremlinDatabaseGetPropertiesResponseOptionsOutput() GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) ToGremlinDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutput() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GremlinDatabaseGetPropertiesResponseOptions) *GremlinDatabaseGetPropertiesResponseOptions {
		return &v
	}).(GremlinDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type GremlinDatabaseGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutput() GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) ToGremlinDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) Elem() GremlinDatabaseGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseOptions) GremlinDatabaseGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret GremlinDatabaseGetPropertiesResponseOptions
		return ret
	}).(GremlinDatabaseGetPropertiesResponseOptionsOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type GremlinDatabaseGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}





type GremlinDatabaseGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToGremlinDatabaseGetPropertiesResponseResourceOutput() GremlinDatabaseGetPropertiesResponseResourceOutput
	ToGremlinDatabaseGetPropertiesResponseResourceOutputWithContext(context.Context) GremlinDatabaseGetPropertiesResponseResourceOutput
}

type GremlinDatabaseGetPropertiesResponseResourceArgs struct {
	Etag pulumi.StringInput  `pulumi:"etag"`
	Id   pulumi.StringInput  `pulumi:"id"`
	Rid  pulumi.StringInput  `pulumi:"rid"`
	Ts   pulumi.Float64Input `pulumi:"ts"`
}

func (GremlinDatabaseGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (i GremlinDatabaseGetPropertiesResponseResourceArgs) ToGremlinDatabaseGetPropertiesResponseResourceOutput() GremlinDatabaseGetPropertiesResponseResourceOutput {
	return i.ToGremlinDatabaseGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i GremlinDatabaseGetPropertiesResponseResourceArgs) ToGremlinDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseGetPropertiesResponseResourceOutput)
}

func (i GremlinDatabaseGetPropertiesResponseResourceArgs) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutput() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return i.ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i GremlinDatabaseGetPropertiesResponseResourceArgs) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseGetPropertiesResponseResourceOutput).ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type GremlinDatabaseGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToGremlinDatabaseGetPropertiesResponseResourcePtrOutput() GremlinDatabaseGetPropertiesResponseResourcePtrOutput
	ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Context) GremlinDatabaseGetPropertiesResponseResourcePtrOutput
}

type gremlinDatabaseGetPropertiesResponseResourcePtrType GremlinDatabaseGetPropertiesResponseResourceArgs

func GremlinDatabaseGetPropertiesResponseResourcePtr(v *GremlinDatabaseGetPropertiesResponseResourceArgs) GremlinDatabaseGetPropertiesResponseResourcePtrInput {
	return (*gremlinDatabaseGetPropertiesResponseResourcePtrType)(v)
}

func (*gremlinDatabaseGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (i *gremlinDatabaseGetPropertiesResponseResourcePtrType) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutput() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return i.ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *gremlinDatabaseGetPropertiesResponseResourcePtrType) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseGetPropertiesResponseResourcePtrOutput)
}

type GremlinDatabaseGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) ToGremlinDatabaseGetPropertiesResponseResourceOutput() GremlinDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) ToGremlinDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutput() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GremlinDatabaseGetPropertiesResponseResource) *GremlinDatabaseGetPropertiesResponseResource {
		return &v
	}).(GremlinDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v GremlinDatabaseGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type GremlinDatabaseGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutput() GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) ToGremlinDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Elem() GremlinDatabaseGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) GremlinDatabaseGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret GremlinDatabaseGetPropertiesResponseResource
		return ret
	}).(GremlinDatabaseGetPropertiesResponseResourceOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o GremlinDatabaseGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type GremlinDatabaseResource struct {
	Id string `pulumi:"id"`
}





type GremlinDatabaseResourceInput interface {
	pulumi.Input

	ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput
	ToGremlinDatabaseResourceOutputWithContext(context.Context) GremlinDatabaseResourceOutput
}

type GremlinDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (GremlinDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseResource)(nil)).Elem()
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput {
	return i.ToGremlinDatabaseResourceOutputWithContext(context.Background())
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourceOutputWithContext(ctx context.Context) GremlinDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseResourceOutput)
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourcePtrOutput() GremlinDatabaseResourcePtrOutput {
	return i.ToGremlinDatabaseResourcePtrOutputWithContext(context.Background())
}

func (i GremlinDatabaseResourceArgs) ToGremlinDatabaseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseResourceOutput).ToGremlinDatabaseResourcePtrOutputWithContext(ctx)
}









type GremlinDatabaseResourcePtrInput interface {
	pulumi.Input

	ToGremlinDatabaseResourcePtrOutput() GremlinDatabaseResourcePtrOutput
	ToGremlinDatabaseResourcePtrOutputWithContext(context.Context) GremlinDatabaseResourcePtrOutput
}

type gremlinDatabaseResourcePtrType GremlinDatabaseResourceArgs

func GremlinDatabaseResourcePtr(v *GremlinDatabaseResourceArgs) GremlinDatabaseResourcePtrInput {
	return (*gremlinDatabaseResourcePtrType)(v)
}

func (*gremlinDatabaseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseResource)(nil)).Elem()
}

func (i *gremlinDatabaseResourcePtrType) ToGremlinDatabaseResourcePtrOutput() GremlinDatabaseResourcePtrOutput {
	return i.ToGremlinDatabaseResourcePtrOutputWithContext(context.Background())
}

func (i *gremlinDatabaseResourcePtrType) ToGremlinDatabaseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinDatabaseResourcePtrOutput)
}

type GremlinDatabaseResourceOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinDatabaseResource)(nil)).Elem()
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourceOutput() GremlinDatabaseResourceOutput {
	return o
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourceOutputWithContext(ctx context.Context) GremlinDatabaseResourceOutput {
	return o
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourcePtrOutput() GremlinDatabaseResourcePtrOutput {
	return o.ToGremlinDatabaseResourcePtrOutputWithContext(context.Background())
}

func (o GremlinDatabaseResourceOutput) ToGremlinDatabaseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GremlinDatabaseResource) *GremlinDatabaseResource {
		return &v
	}).(GremlinDatabaseResourcePtrOutput)
}

func (o GremlinDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type GremlinDatabaseResourcePtrOutput struct{ *pulumi.OutputState }

func (GremlinDatabaseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinDatabaseResource)(nil)).Elem()
}

func (o GremlinDatabaseResourcePtrOutput) ToGremlinDatabaseResourcePtrOutput() GremlinDatabaseResourcePtrOutput {
	return o
}

func (o GremlinDatabaseResourcePtrOutput) ToGremlinDatabaseResourcePtrOutputWithContext(ctx context.Context) GremlinDatabaseResourcePtrOutput {
	return o
}

func (o GremlinDatabaseResourcePtrOutput) Elem() GremlinDatabaseResourceOutput {
	return o.ApplyT(func(v *GremlinDatabaseResource) GremlinDatabaseResource {
		if v != nil {
			return *v
		}
		var ret GremlinDatabaseResource
		return ret
	}).(GremlinDatabaseResourceOutput)
}

func (o GremlinDatabaseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinDatabaseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type GremlinGraphGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type GremlinGraphGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToGremlinGraphGetPropertiesResponseOptionsOutput() GremlinGraphGetPropertiesResponseOptionsOutput
	ToGremlinGraphGetPropertiesResponseOptionsOutputWithContext(context.Context) GremlinGraphGetPropertiesResponseOptionsOutput
}

type GremlinGraphGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (GremlinGraphGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphGetPropertiesResponseOptions)(nil)).Elem()
}

func (i GremlinGraphGetPropertiesResponseOptionsArgs) ToGremlinGraphGetPropertiesResponseOptionsOutput() GremlinGraphGetPropertiesResponseOptionsOutput {
	return i.ToGremlinGraphGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i GremlinGraphGetPropertiesResponseOptionsArgs) ToGremlinGraphGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphGetPropertiesResponseOptionsOutput)
}

func (i GremlinGraphGetPropertiesResponseOptionsArgs) ToGremlinGraphGetPropertiesResponseOptionsPtrOutput() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return i.ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i GremlinGraphGetPropertiesResponseOptionsArgs) ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphGetPropertiesResponseOptionsOutput).ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type GremlinGraphGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToGremlinGraphGetPropertiesResponseOptionsPtrOutput() GremlinGraphGetPropertiesResponseOptionsPtrOutput
	ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) GremlinGraphGetPropertiesResponseOptionsPtrOutput
}

type gremlinGraphGetPropertiesResponseOptionsPtrType GremlinGraphGetPropertiesResponseOptionsArgs

func GremlinGraphGetPropertiesResponseOptionsPtr(v *GremlinGraphGetPropertiesResponseOptionsArgs) GremlinGraphGetPropertiesResponseOptionsPtrInput {
	return (*gremlinGraphGetPropertiesResponseOptionsPtrType)(v)
}

func (*gremlinGraphGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *gremlinGraphGetPropertiesResponseOptionsPtrType) ToGremlinGraphGetPropertiesResponseOptionsPtrOutput() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return i.ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *gremlinGraphGetPropertiesResponseOptionsPtrType) ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphGetPropertiesResponseOptionsPtrOutput)
}

type GremlinGraphGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) ToGremlinGraphGetPropertiesResponseOptionsOutput() GremlinGraphGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) ToGremlinGraphGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) ToGremlinGraphGetPropertiesResponseOptionsPtrOutput() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o.ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GremlinGraphGetPropertiesResponseOptions) *GremlinGraphGetPropertiesResponseOptions {
		return &v
	}).(GremlinGraphGetPropertiesResponseOptionsPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type GremlinGraphGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphGetPropertiesResponseOptions)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) ToGremlinGraphGetPropertiesResponseOptionsPtrOutput() GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) ToGremlinGraphGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) Elem() GremlinGraphGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseOptions) GremlinGraphGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret GremlinGraphGetPropertiesResponseOptions
		return ret
	}).(GremlinGraphGetPropertiesResponseOptionsOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type GremlinGraphGetPropertiesResponseResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                              `pulumi:"defaultTtl"`
	Etag                     string                            `pulumi:"etag"`
	Id                       string                            `pulumi:"id"`
	IndexingPolicy           *IndexingPolicyResponse           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKeyResponse    `pulumi:"partitionKey"`
	Rid                      string                            `pulumi:"rid"`
	Ts                       float64                           `pulumi:"ts"`
	UniqueKeyPolicy          *UniqueKeyPolicyResponse          `pulumi:"uniqueKeyPolicy"`
}


func (val *GremlinGraphGetPropertiesResponseResource) Defaults() *GremlinGraphGetPropertiesResponseResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type GremlinGraphGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToGremlinGraphGetPropertiesResponseResourceOutput() GremlinGraphGetPropertiesResponseResourceOutput
	ToGremlinGraphGetPropertiesResponseResourceOutputWithContext(context.Context) GremlinGraphGetPropertiesResponseResourceOutput
}

type GremlinGraphGetPropertiesResponseResourceArgs struct {
	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput                       `pulumi:"defaultTtl"`
	Etag                     pulumi.StringInput                       `pulumi:"etag"`
	Id                       pulumi.StringInput                       `pulumi:"id"`
	IndexingPolicy           IndexingPolicyResponsePtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyResponsePtrInput    `pulumi:"partitionKey"`
	Rid                      pulumi.StringInput                       `pulumi:"rid"`
	Ts                       pulumi.Float64Input                      `pulumi:"ts"`
	UniqueKeyPolicy          UniqueKeyPolicyResponsePtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (GremlinGraphGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphGetPropertiesResponseResource)(nil)).Elem()
}

func (i GremlinGraphGetPropertiesResponseResourceArgs) ToGremlinGraphGetPropertiesResponseResourceOutput() GremlinGraphGetPropertiesResponseResourceOutput {
	return i.ToGremlinGraphGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i GremlinGraphGetPropertiesResponseResourceArgs) ToGremlinGraphGetPropertiesResponseResourceOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphGetPropertiesResponseResourceOutput)
}

func (i GremlinGraphGetPropertiesResponseResourceArgs) ToGremlinGraphGetPropertiesResponseResourcePtrOutput() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return i.ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i GremlinGraphGetPropertiesResponseResourceArgs) ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphGetPropertiesResponseResourceOutput).ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type GremlinGraphGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToGremlinGraphGetPropertiesResponseResourcePtrOutput() GremlinGraphGetPropertiesResponseResourcePtrOutput
	ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(context.Context) GremlinGraphGetPropertiesResponseResourcePtrOutput
}

type gremlinGraphGetPropertiesResponseResourcePtrType GremlinGraphGetPropertiesResponseResourceArgs

func GremlinGraphGetPropertiesResponseResourcePtr(v *GremlinGraphGetPropertiesResponseResourceArgs) GremlinGraphGetPropertiesResponseResourcePtrInput {
	return (*gremlinGraphGetPropertiesResponseResourcePtrType)(v)
}

func (*gremlinGraphGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphGetPropertiesResponseResource)(nil)).Elem()
}

func (i *gremlinGraphGetPropertiesResponseResourcePtrType) ToGremlinGraphGetPropertiesResponseResourcePtrOutput() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return i.ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *gremlinGraphGetPropertiesResponseResourcePtrType) ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphGetPropertiesResponseResourcePtrOutput)
}

type GremlinGraphGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ToGremlinGraphGetPropertiesResponseResourceOutput() GremlinGraphGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ToGremlinGraphGetPropertiesResponseResourceOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourceOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ToGremlinGraphGetPropertiesResponseResourcePtrOutput() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o.ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GremlinGraphGetPropertiesResponseResource) *GremlinGraphGetPropertiesResponseResource {
		return &v
	}).(GremlinGraphGetPropertiesResponseResourcePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *IndexingPolicyResponse { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

func (o GremlinGraphGetPropertiesResponseResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v GremlinGraphGetPropertiesResponseResource) *UniqueKeyPolicyResponse { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

type GremlinGraphGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (GremlinGraphGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphGetPropertiesResponseResource)(nil)).Elem()
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) ToGremlinGraphGetPropertiesResponseResourcePtrOutput() GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) ToGremlinGraphGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) GremlinGraphGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Elem() GremlinGraphGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) GremlinGraphGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret GremlinGraphGetPropertiesResponseResource
		return ret
	}).(GremlinGraphGetPropertiesResponseResourceOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *IndexingPolicyResponse {
		if v == nil {
			return nil
		}
		return v.IndexingPolicy
	}).(IndexingPolicyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

func (o GremlinGraphGetPropertiesResponseResourcePtrOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *GremlinGraphGetPropertiesResponseResource) *UniqueKeyPolicyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeyPolicy
	}).(UniqueKeyPolicyResponsePtrOutput)
}

type GremlinGraphResource struct {
	ConflictResolutionPolicy *ConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                      `pulumi:"defaultTtl"`
	Id                       string                    `pulumi:"id"`
	IndexingPolicy           *IndexingPolicy           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKey    `pulumi:"partitionKey"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `pulumi:"uniqueKeyPolicy"`
}


func (val *GremlinGraphResource) Defaults() *GremlinGraphResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type GremlinGraphResourceInput interface {
	pulumi.Input

	ToGremlinGraphResourceOutput() GremlinGraphResourceOutput
	ToGremlinGraphResourceOutputWithContext(context.Context) GremlinGraphResourceOutput
}

type GremlinGraphResourceArgs struct {
	ConflictResolutionPolicy ConflictResolutionPolicyPtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput               `pulumi:"defaultTtl"`
	Id                       pulumi.StringInput               `pulumi:"id"`
	IndexingPolicy           IndexingPolicyPtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyPtrInput    `pulumi:"partitionKey"`
	UniqueKeyPolicy          UniqueKeyPolicyPtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (GremlinGraphResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphResource)(nil)).Elem()
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourceOutput() GremlinGraphResourceOutput {
	return i.ToGremlinGraphResourceOutputWithContext(context.Background())
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourceOutputWithContext(ctx context.Context) GremlinGraphResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphResourceOutput)
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourcePtrOutput() GremlinGraphResourcePtrOutput {
	return i.ToGremlinGraphResourcePtrOutputWithContext(context.Background())
}

func (i GremlinGraphResourceArgs) ToGremlinGraphResourcePtrOutputWithContext(ctx context.Context) GremlinGraphResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphResourceOutput).ToGremlinGraphResourcePtrOutputWithContext(ctx)
}









type GremlinGraphResourcePtrInput interface {
	pulumi.Input

	ToGremlinGraphResourcePtrOutput() GremlinGraphResourcePtrOutput
	ToGremlinGraphResourcePtrOutputWithContext(context.Context) GremlinGraphResourcePtrOutput
}

type gremlinGraphResourcePtrType GremlinGraphResourceArgs

func GremlinGraphResourcePtr(v *GremlinGraphResourceArgs) GremlinGraphResourcePtrInput {
	return (*gremlinGraphResourcePtrType)(v)
}

func (*gremlinGraphResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphResource)(nil)).Elem()
}

func (i *gremlinGraphResourcePtrType) ToGremlinGraphResourcePtrOutput() GremlinGraphResourcePtrOutput {
	return i.ToGremlinGraphResourcePtrOutputWithContext(context.Background())
}

func (i *gremlinGraphResourcePtrType) ToGremlinGraphResourcePtrOutputWithContext(ctx context.Context) GremlinGraphResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GremlinGraphResourcePtrOutput)
}

type GremlinGraphResourceOutput struct{ *pulumi.OutputState }

func (GremlinGraphResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GremlinGraphResource)(nil)).Elem()
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourceOutput() GremlinGraphResourceOutput {
	return o
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourceOutputWithContext(ctx context.Context) GremlinGraphResourceOutput {
	return o
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourcePtrOutput() GremlinGraphResourcePtrOutput {
	return o.ToGremlinGraphResourcePtrOutputWithContext(context.Background())
}

func (o GremlinGraphResourceOutput) ToGremlinGraphResourcePtrOutputWithContext(ctx context.Context) GremlinGraphResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GremlinGraphResource) *GremlinGraphResource {
		return &v
	}).(GremlinGraphResourcePtrOutput)
}

func (o GremlinGraphResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *ConflictResolutionPolicy { return v.ConflictResolutionPolicy }).(ConflictResolutionPolicyPtrOutput)
}

func (o GremlinGraphResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o GremlinGraphResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GremlinGraphResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o GremlinGraphResourceOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *IndexingPolicy { return v.IndexingPolicy }).(IndexingPolicyPtrOutput)
}

func (o GremlinGraphResourceOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *ContainerPartitionKey { return v.PartitionKey }).(ContainerPartitionKeyPtrOutput)
}

func (o GremlinGraphResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v GremlinGraphResource) *UniqueKeyPolicy { return v.UniqueKeyPolicy }).(UniqueKeyPolicyPtrOutput)
}

type GremlinGraphResourcePtrOutput struct{ *pulumi.OutputState }

func (GremlinGraphResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GremlinGraphResource)(nil)).Elem()
}

func (o GremlinGraphResourcePtrOutput) ToGremlinGraphResourcePtrOutput() GremlinGraphResourcePtrOutput {
	return o
}

func (o GremlinGraphResourcePtrOutput) ToGremlinGraphResourcePtrOutputWithContext(ctx context.Context) GremlinGraphResourcePtrOutput {
	return o
}

func (o GremlinGraphResourcePtrOutput) Elem() GremlinGraphResourceOutput {
	return o.ApplyT(func(v *GremlinGraphResource) GremlinGraphResource {
		if v != nil {
			return *v
		}
		var ret GremlinGraphResource
		return ret
	}).(GremlinGraphResourceOutput)
}

func (o GremlinGraphResourcePtrOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v *GremlinGraphResource) *ConflictResolutionPolicy {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyPtrOutput)
}

func (o GremlinGraphResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *GremlinGraphResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o GremlinGraphResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GremlinGraphResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o GremlinGraphResourcePtrOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v *GremlinGraphResource) *IndexingPolicy {
		if v == nil {
			return nil
		}
		return v.IndexingPolicy
	}).(IndexingPolicyPtrOutput)
}

func (o GremlinGraphResourcePtrOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v *GremlinGraphResource) *ContainerPartitionKey {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(ContainerPartitionKeyPtrOutput)
}

func (o GremlinGraphResourcePtrOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v *GremlinGraphResource) *UniqueKeyPolicy {
		if v == nil {
			return nil
		}
		return v.UniqueKeyPolicy
	}).(UniqueKeyPolicyPtrOutput)
}

type IncludedPath struct {
	Indexes []Indexes `pulumi:"indexes"`
	Path    *string   `pulumi:"path"`
}





type IncludedPathInput interface {
	pulumi.Input

	ToIncludedPathOutput() IncludedPathOutput
	ToIncludedPathOutputWithContext(context.Context) IncludedPathOutput
}

type IncludedPathArgs struct {
	Indexes IndexesArrayInput     `pulumi:"indexes"`
	Path    pulumi.StringPtrInput `pulumi:"path"`
}

func (IncludedPathArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPath)(nil)).Elem()
}

func (i IncludedPathArgs) ToIncludedPathOutput() IncludedPathOutput {
	return i.ToIncludedPathOutputWithContext(context.Background())
}

func (i IncludedPathArgs) ToIncludedPathOutputWithContext(ctx context.Context) IncludedPathOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathOutput)
}





type IncludedPathArrayInput interface {
	pulumi.Input

	ToIncludedPathArrayOutput() IncludedPathArrayOutput
	ToIncludedPathArrayOutputWithContext(context.Context) IncludedPathArrayOutput
}

type IncludedPathArray []IncludedPathInput

func (IncludedPathArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPath)(nil)).Elem()
}

func (i IncludedPathArray) ToIncludedPathArrayOutput() IncludedPathArrayOutput {
	return i.ToIncludedPathArrayOutputWithContext(context.Background())
}

func (i IncludedPathArray) ToIncludedPathArrayOutputWithContext(ctx context.Context) IncludedPathArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathArrayOutput)
}

type IncludedPathOutput struct{ *pulumi.OutputState }

func (IncludedPathOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPath)(nil)).Elem()
}

func (o IncludedPathOutput) ToIncludedPathOutput() IncludedPathOutput {
	return o
}

func (o IncludedPathOutput) ToIncludedPathOutputWithContext(ctx context.Context) IncludedPathOutput {
	return o
}

func (o IncludedPathOutput) Indexes() IndexesArrayOutput {
	return o.ApplyT(func(v IncludedPath) []Indexes { return v.Indexes }).(IndexesArrayOutput)
}

func (o IncludedPathOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncludedPath) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type IncludedPathArrayOutput struct{ *pulumi.OutputState }

func (IncludedPathArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPath)(nil)).Elem()
}

func (o IncludedPathArrayOutput) ToIncludedPathArrayOutput() IncludedPathArrayOutput {
	return o
}

func (o IncludedPathArrayOutput) ToIncludedPathArrayOutputWithContext(ctx context.Context) IncludedPathArrayOutput {
	return o
}

func (o IncludedPathArrayOutput) Index(i pulumi.IntInput) IncludedPathOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncludedPath {
		return vs[0].([]IncludedPath)[vs[1].(int)]
	}).(IncludedPathOutput)
}

type IncludedPathResponse struct {
	Indexes []IndexesResponse `pulumi:"indexes"`
	Path    *string           `pulumi:"path"`
}





type IncludedPathResponseInput interface {
	pulumi.Input

	ToIncludedPathResponseOutput() IncludedPathResponseOutput
	ToIncludedPathResponseOutputWithContext(context.Context) IncludedPathResponseOutput
}

type IncludedPathResponseArgs struct {
	Indexes IndexesResponseArrayInput `pulumi:"indexes"`
	Path    pulumi.StringPtrInput     `pulumi:"path"`
}

func (IncludedPathResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPathResponse)(nil)).Elem()
}

func (i IncludedPathResponseArgs) ToIncludedPathResponseOutput() IncludedPathResponseOutput {
	return i.ToIncludedPathResponseOutputWithContext(context.Background())
}

func (i IncludedPathResponseArgs) ToIncludedPathResponseOutputWithContext(ctx context.Context) IncludedPathResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathResponseOutput)
}





type IncludedPathResponseArrayInput interface {
	pulumi.Input

	ToIncludedPathResponseArrayOutput() IncludedPathResponseArrayOutput
	ToIncludedPathResponseArrayOutputWithContext(context.Context) IncludedPathResponseArrayOutput
}

type IncludedPathResponseArray []IncludedPathResponseInput

func (IncludedPathResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPathResponse)(nil)).Elem()
}

func (i IncludedPathResponseArray) ToIncludedPathResponseArrayOutput() IncludedPathResponseArrayOutput {
	return i.ToIncludedPathResponseArrayOutputWithContext(context.Background())
}

func (i IncludedPathResponseArray) ToIncludedPathResponseArrayOutputWithContext(ctx context.Context) IncludedPathResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IncludedPathResponseArrayOutput)
}

type IncludedPathResponseOutput struct{ *pulumi.OutputState }

func (IncludedPathResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IncludedPathResponse)(nil)).Elem()
}

func (o IncludedPathResponseOutput) ToIncludedPathResponseOutput() IncludedPathResponseOutput {
	return o
}

func (o IncludedPathResponseOutput) ToIncludedPathResponseOutputWithContext(ctx context.Context) IncludedPathResponseOutput {
	return o
}

func (o IncludedPathResponseOutput) Indexes() IndexesResponseArrayOutput {
	return o.ApplyT(func(v IncludedPathResponse) []IndexesResponse { return v.Indexes }).(IndexesResponseArrayOutput)
}

func (o IncludedPathResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IncludedPathResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

type IncludedPathResponseArrayOutput struct{ *pulumi.OutputState }

func (IncludedPathResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IncludedPathResponse)(nil)).Elem()
}

func (o IncludedPathResponseArrayOutput) ToIncludedPathResponseArrayOutput() IncludedPathResponseArrayOutput {
	return o
}

func (o IncludedPathResponseArrayOutput) ToIncludedPathResponseArrayOutputWithContext(ctx context.Context) IncludedPathResponseArrayOutput {
	return o
}

func (o IncludedPathResponseArrayOutput) Index(i pulumi.IntInput) IncludedPathResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IncludedPathResponse {
		return vs[0].([]IncludedPathResponse)[vs[1].(int)]
	}).(IncludedPathResponseOutput)
}

type Indexes struct {
	DataType  *string `pulumi:"dataType"`
	Kind      *string `pulumi:"kind"`
	Precision *int    `pulumi:"precision"`
}


func (val *Indexes) Defaults() *Indexes {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataType) {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type IndexesInput interface {
	pulumi.Input

	ToIndexesOutput() IndexesOutput
	ToIndexesOutputWithContext(context.Context) IndexesOutput
}

type IndexesArgs struct {
	DataType  pulumi.StringPtrInput `pulumi:"dataType"`
	Kind      pulumi.StringPtrInput `pulumi:"kind"`
	Precision pulumi.IntPtrInput    `pulumi:"precision"`
}

func (IndexesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Indexes)(nil)).Elem()
}

func (i IndexesArgs) ToIndexesOutput() IndexesOutput {
	return i.ToIndexesOutputWithContext(context.Background())
}

func (i IndexesArgs) ToIndexesOutputWithContext(ctx context.Context) IndexesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesOutput)
}





type IndexesArrayInput interface {
	pulumi.Input

	ToIndexesArrayOutput() IndexesArrayOutput
	ToIndexesArrayOutputWithContext(context.Context) IndexesArrayOutput
}

type IndexesArray []IndexesInput

func (IndexesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Indexes)(nil)).Elem()
}

func (i IndexesArray) ToIndexesArrayOutput() IndexesArrayOutput {
	return i.ToIndexesArrayOutputWithContext(context.Background())
}

func (i IndexesArray) ToIndexesArrayOutputWithContext(ctx context.Context) IndexesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesArrayOutput)
}

type IndexesOutput struct{ *pulumi.OutputState }

func (IndexesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Indexes)(nil)).Elem()
}

func (o IndexesOutput) ToIndexesOutput() IndexesOutput {
	return o
}

func (o IndexesOutput) ToIndexesOutputWithContext(ctx context.Context) IndexesOutput {
	return o
}

func (o IndexesOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Indexes) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o IndexesOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Indexes) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IndexesOutput) Precision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Indexes) *int { return v.Precision }).(pulumi.IntPtrOutput)
}

type IndexesArrayOutput struct{ *pulumi.OutputState }

func (IndexesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Indexes)(nil)).Elem()
}

func (o IndexesArrayOutput) ToIndexesArrayOutput() IndexesArrayOutput {
	return o
}

func (o IndexesArrayOutput) ToIndexesArrayOutputWithContext(ctx context.Context) IndexesArrayOutput {
	return o
}

func (o IndexesArrayOutput) Index(i pulumi.IntInput) IndexesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Indexes {
		return vs[0].([]Indexes)[vs[1].(int)]
	}).(IndexesOutput)
}

type IndexesResponse struct {
	DataType  *string `pulumi:"dataType"`
	Kind      *string `pulumi:"kind"`
	Precision *int    `pulumi:"precision"`
}


func (val *IndexesResponse) Defaults() *IndexesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DataType) {
		dataType_ := "String"
		tmp.DataType = &dataType_
	}
	if isZero(tmp.Kind) {
		kind_ := "Hash"
		tmp.Kind = &kind_
	}
	return &tmp
}





type IndexesResponseInput interface {
	pulumi.Input

	ToIndexesResponseOutput() IndexesResponseOutput
	ToIndexesResponseOutputWithContext(context.Context) IndexesResponseOutput
}

type IndexesResponseArgs struct {
	DataType  pulumi.StringPtrInput `pulumi:"dataType"`
	Kind      pulumi.StringPtrInput `pulumi:"kind"`
	Precision pulumi.IntPtrInput    `pulumi:"precision"`
}

func (IndexesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexesResponse)(nil)).Elem()
}

func (i IndexesResponseArgs) ToIndexesResponseOutput() IndexesResponseOutput {
	return i.ToIndexesResponseOutputWithContext(context.Background())
}

func (i IndexesResponseArgs) ToIndexesResponseOutputWithContext(ctx context.Context) IndexesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesResponseOutput)
}





type IndexesResponseArrayInput interface {
	pulumi.Input

	ToIndexesResponseArrayOutput() IndexesResponseArrayOutput
	ToIndexesResponseArrayOutputWithContext(context.Context) IndexesResponseArrayOutput
}

type IndexesResponseArray []IndexesResponseInput

func (IndexesResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IndexesResponse)(nil)).Elem()
}

func (i IndexesResponseArray) ToIndexesResponseArrayOutput() IndexesResponseArrayOutput {
	return i.ToIndexesResponseArrayOutputWithContext(context.Background())
}

func (i IndexesResponseArray) ToIndexesResponseArrayOutputWithContext(ctx context.Context) IndexesResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexesResponseArrayOutput)
}

type IndexesResponseOutput struct{ *pulumi.OutputState }

func (IndexesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexesResponse)(nil)).Elem()
}

func (o IndexesResponseOutput) ToIndexesResponseOutput() IndexesResponseOutput {
	return o
}

func (o IndexesResponseOutput) ToIndexesResponseOutputWithContext(ctx context.Context) IndexesResponseOutput {
	return o
}

func (o IndexesResponseOutput) DataType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *string { return v.DataType }).(pulumi.StringPtrOutput)
}

func (o IndexesResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IndexesResponseOutput) Precision() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IndexesResponse) *int { return v.Precision }).(pulumi.IntPtrOutput)
}

type IndexesResponseArrayOutput struct{ *pulumi.OutputState }

func (IndexesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IndexesResponse)(nil)).Elem()
}

func (o IndexesResponseArrayOutput) ToIndexesResponseArrayOutput() IndexesResponseArrayOutput {
	return o
}

func (o IndexesResponseArrayOutput) ToIndexesResponseArrayOutputWithContext(ctx context.Context) IndexesResponseArrayOutput {
	return o
}

func (o IndexesResponseArrayOutput) Index(i pulumi.IntInput) IndexesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IndexesResponse {
		return vs[0].([]IndexesResponse)[vs[1].(int)]
	}).(IndexesResponseOutput)
}

type IndexingPolicy struct {
	Automatic        *bool             `pulumi:"automatic"`
	CompositeIndexes [][]CompositePath `pulumi:"compositeIndexes"`
	ExcludedPaths    []ExcludedPath    `pulumi:"excludedPaths"`
	IncludedPaths    []IncludedPath    `pulumi:"includedPaths"`
	IndexingMode     *string           `pulumi:"indexingMode"`
	SpatialIndexes   []SpatialSpec     `pulumi:"spatialIndexes"`
}


func (val *IndexingPolicy) Defaults() *IndexingPolicy {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IndexingMode) {
		indexingMode_ := "consistent"
		tmp.IndexingMode = &indexingMode_
	}
	return &tmp
}





type IndexingPolicyInput interface {
	pulumi.Input

	ToIndexingPolicyOutput() IndexingPolicyOutput
	ToIndexingPolicyOutputWithContext(context.Context) IndexingPolicyOutput
}

type IndexingPolicyArgs struct {
	Automatic        pulumi.BoolPtrInput          `pulumi:"automatic"`
	CompositeIndexes CompositePathArrayArrayInput `pulumi:"compositeIndexes"`
	ExcludedPaths    ExcludedPathArrayInput       `pulumi:"excludedPaths"`
	IncludedPaths    IncludedPathArrayInput       `pulumi:"includedPaths"`
	IndexingMode     pulumi.StringPtrInput        `pulumi:"indexingMode"`
	SpatialIndexes   SpatialSpecArrayInput        `pulumi:"spatialIndexes"`
}

func (IndexingPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicy)(nil)).Elem()
}

func (i IndexingPolicyArgs) ToIndexingPolicyOutput() IndexingPolicyOutput {
	return i.ToIndexingPolicyOutputWithContext(context.Background())
}

func (i IndexingPolicyArgs) ToIndexingPolicyOutputWithContext(ctx context.Context) IndexingPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyOutput)
}

func (i IndexingPolicyArgs) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return i.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (i IndexingPolicyArgs) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyOutput).ToIndexingPolicyPtrOutputWithContext(ctx)
}









type IndexingPolicyPtrInput interface {
	pulumi.Input

	ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput
	ToIndexingPolicyPtrOutputWithContext(context.Context) IndexingPolicyPtrOutput
}

type indexingPolicyPtrType IndexingPolicyArgs

func IndexingPolicyPtr(v *IndexingPolicyArgs) IndexingPolicyPtrInput {
	return (*indexingPolicyPtrType)(v)
}

func (*indexingPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicy)(nil)).Elem()
}

func (i *indexingPolicyPtrType) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return i.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (i *indexingPolicyPtrType) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyPtrOutput)
}

type IndexingPolicyOutput struct{ *pulumi.OutputState }

func (IndexingPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicy)(nil)).Elem()
}

func (o IndexingPolicyOutput) ToIndexingPolicyOutput() IndexingPolicyOutput {
	return o
}

func (o IndexingPolicyOutput) ToIndexingPolicyOutputWithContext(ctx context.Context) IndexingPolicyOutput {
	return o
}

func (o IndexingPolicyOutput) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return o.ToIndexingPolicyPtrOutputWithContext(context.Background())
}

func (o IndexingPolicyOutput) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexingPolicy) *IndexingPolicy {
		return &v
	}).(IndexingPolicyPtrOutput)
}

func (o IndexingPolicyOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IndexingPolicy) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyOutput) CompositeIndexes() CompositePathArrayArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) [][]CompositePath { return v.CompositeIndexes }).(CompositePathArrayArrayOutput)
}

func (o IndexingPolicyOutput) ExcludedPaths() ExcludedPathArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []ExcludedPath { return v.ExcludedPaths }).(ExcludedPathArrayOutput)
}

func (o IndexingPolicyOutput) IncludedPaths() IncludedPathArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []IncludedPath { return v.IncludedPaths }).(IncludedPathArrayOutput)
}

func (o IndexingPolicyOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexingPolicy) *string { return v.IndexingMode }).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyOutput) SpatialIndexes() SpatialSpecArrayOutput {
	return o.ApplyT(func(v IndexingPolicy) []SpatialSpec { return v.SpatialIndexes }).(SpatialSpecArrayOutput)
}

type IndexingPolicyPtrOutput struct{ *pulumi.OutputState }

func (IndexingPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicy)(nil)).Elem()
}

func (o IndexingPolicyPtrOutput) ToIndexingPolicyPtrOutput() IndexingPolicyPtrOutput {
	return o
}

func (o IndexingPolicyPtrOutput) ToIndexingPolicyPtrOutputWithContext(ctx context.Context) IndexingPolicyPtrOutput {
	return o
}

func (o IndexingPolicyPtrOutput) Elem() IndexingPolicyOutput {
	return o.ApplyT(func(v *IndexingPolicy) IndexingPolicy {
		if v != nil {
			return *v
		}
		var ret IndexingPolicy
		return ret
	}).(IndexingPolicyOutput)
}

func (o IndexingPolicyPtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IndexingPolicy) *bool {
		if v == nil {
			return nil
		}
		return v.Automatic
	}).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyPtrOutput) CompositeIndexes() CompositePathArrayArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) [][]CompositePath {
		if v == nil {
			return nil
		}
		return v.CompositeIndexes
	}).(CompositePathArrayArrayOutput)
}

func (o IndexingPolicyPtrOutput) ExcludedPaths() ExcludedPathArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []ExcludedPath {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(ExcludedPathArrayOutput)
}

func (o IndexingPolicyPtrOutput) IncludedPaths() IncludedPathArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []IncludedPath {
		if v == nil {
			return nil
		}
		return v.IncludedPaths
	}).(IncludedPathArrayOutput)
}

func (o IndexingPolicyPtrOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexingPolicy) *string {
		if v == nil {
			return nil
		}
		return v.IndexingMode
	}).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyPtrOutput) SpatialIndexes() SpatialSpecArrayOutput {
	return o.ApplyT(func(v *IndexingPolicy) []SpatialSpec {
		if v == nil {
			return nil
		}
		return v.SpatialIndexes
	}).(SpatialSpecArrayOutput)
}

type IndexingPolicyResponse struct {
	Automatic        *bool                     `pulumi:"automatic"`
	CompositeIndexes [][]CompositePathResponse `pulumi:"compositeIndexes"`
	ExcludedPaths    []ExcludedPathResponse    `pulumi:"excludedPaths"`
	IncludedPaths    []IncludedPathResponse    `pulumi:"includedPaths"`
	IndexingMode     *string                   `pulumi:"indexingMode"`
	SpatialIndexes   []SpatialSpecResponse     `pulumi:"spatialIndexes"`
}


func (val *IndexingPolicyResponse) Defaults() *IndexingPolicyResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IndexingMode) {
		indexingMode_ := "consistent"
		tmp.IndexingMode = &indexingMode_
	}
	return &tmp
}





type IndexingPolicyResponseInput interface {
	pulumi.Input

	ToIndexingPolicyResponseOutput() IndexingPolicyResponseOutput
	ToIndexingPolicyResponseOutputWithContext(context.Context) IndexingPolicyResponseOutput
}

type IndexingPolicyResponseArgs struct {
	Automatic        pulumi.BoolPtrInput                  `pulumi:"automatic"`
	CompositeIndexes CompositePathResponseArrayArrayInput `pulumi:"compositeIndexes"`
	ExcludedPaths    ExcludedPathResponseArrayInput       `pulumi:"excludedPaths"`
	IncludedPaths    IncludedPathResponseArrayInput       `pulumi:"includedPaths"`
	IndexingMode     pulumi.StringPtrInput                `pulumi:"indexingMode"`
	SpatialIndexes   SpatialSpecResponseArrayInput        `pulumi:"spatialIndexes"`
}

func (IndexingPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicyResponse)(nil)).Elem()
}

func (i IndexingPolicyResponseArgs) ToIndexingPolicyResponseOutput() IndexingPolicyResponseOutput {
	return i.ToIndexingPolicyResponseOutputWithContext(context.Background())
}

func (i IndexingPolicyResponseArgs) ToIndexingPolicyResponseOutputWithContext(ctx context.Context) IndexingPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyResponseOutput)
}

func (i IndexingPolicyResponseArgs) ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput {
	return i.ToIndexingPolicyResponsePtrOutputWithContext(context.Background())
}

func (i IndexingPolicyResponseArgs) ToIndexingPolicyResponsePtrOutputWithContext(ctx context.Context) IndexingPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyResponseOutput).ToIndexingPolicyResponsePtrOutputWithContext(ctx)
}









type IndexingPolicyResponsePtrInput interface {
	pulumi.Input

	ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput
	ToIndexingPolicyResponsePtrOutputWithContext(context.Context) IndexingPolicyResponsePtrOutput
}

type indexingPolicyResponsePtrType IndexingPolicyResponseArgs

func IndexingPolicyResponsePtr(v *IndexingPolicyResponseArgs) IndexingPolicyResponsePtrInput {
	return (*indexingPolicyResponsePtrType)(v)
}

func (*indexingPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicyResponse)(nil)).Elem()
}

func (i *indexingPolicyResponsePtrType) ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput {
	return i.ToIndexingPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *indexingPolicyResponsePtrType) ToIndexingPolicyResponsePtrOutputWithContext(ctx context.Context) IndexingPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IndexingPolicyResponsePtrOutput)
}

type IndexingPolicyResponseOutput struct{ *pulumi.OutputState }

func (IndexingPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IndexingPolicyResponse)(nil)).Elem()
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponseOutput() IndexingPolicyResponseOutput {
	return o
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponseOutputWithContext(ctx context.Context) IndexingPolicyResponseOutput {
	return o
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput {
	return o.ToIndexingPolicyResponsePtrOutputWithContext(context.Background())
}

func (o IndexingPolicyResponseOutput) ToIndexingPolicyResponsePtrOutputWithContext(ctx context.Context) IndexingPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IndexingPolicyResponse) *IndexingPolicyResponse {
		return &v
	}).(IndexingPolicyResponsePtrOutput)
}

func (o IndexingPolicyResponseOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) *bool { return v.Automatic }).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyResponseOutput) CompositeIndexes() CompositePathResponseArrayArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) [][]CompositePathResponse { return v.CompositeIndexes }).(CompositePathResponseArrayArrayOutput)
}

func (o IndexingPolicyResponseOutput) ExcludedPaths() ExcludedPathResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []ExcludedPathResponse { return v.ExcludedPaths }).(ExcludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponseOutput) IncludedPaths() IncludedPathResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []IncludedPathResponse { return v.IncludedPaths }).(IncludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponseOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) *string { return v.IndexingMode }).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyResponseOutput) SpatialIndexes() SpatialSpecResponseArrayOutput {
	return o.ApplyT(func(v IndexingPolicyResponse) []SpatialSpecResponse { return v.SpatialIndexes }).(SpatialSpecResponseArrayOutput)
}

type IndexingPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (IndexingPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IndexingPolicyResponse)(nil)).Elem()
}

func (o IndexingPolicyResponsePtrOutput) ToIndexingPolicyResponsePtrOutput() IndexingPolicyResponsePtrOutput {
	return o
}

func (o IndexingPolicyResponsePtrOutput) ToIndexingPolicyResponsePtrOutputWithContext(ctx context.Context) IndexingPolicyResponsePtrOutput {
	return o
}

func (o IndexingPolicyResponsePtrOutput) Elem() IndexingPolicyResponseOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) IndexingPolicyResponse {
		if v != nil {
			return *v
		}
		var ret IndexingPolicyResponse
		return ret
	}).(IndexingPolicyResponseOutput)
}

func (o IndexingPolicyResponsePtrOutput) Automatic() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Automatic
	}).(pulumi.BoolPtrOutput)
}

func (o IndexingPolicyResponsePtrOutput) CompositeIndexes() CompositePathResponseArrayArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) [][]CompositePathResponse {
		if v == nil {
			return nil
		}
		return v.CompositeIndexes
	}).(CompositePathResponseArrayArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) ExcludedPaths() ExcludedPathResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []ExcludedPathResponse {
		if v == nil {
			return nil
		}
		return v.ExcludedPaths
	}).(ExcludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) IncludedPaths() IncludedPathResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []IncludedPathResponse {
		if v == nil {
			return nil
		}
		return v.IncludedPaths
	}).(IncludedPathResponseArrayOutput)
}

func (o IndexingPolicyResponsePtrOutput) IndexingMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) *string {
		if v == nil {
			return nil
		}
		return v.IndexingMode
	}).(pulumi.StringPtrOutput)
}

func (o IndexingPolicyResponsePtrOutput) SpatialIndexes() SpatialSpecResponseArrayOutput {
	return o.ApplyT(func(v *IndexingPolicyResponse) []SpatialSpecResponse {
		if v == nil {
			return nil
		}
		return v.SpatialIndexes
	}).(SpatialSpecResponseArrayOutput)
}

type IpAddressOrRange struct {
	IpAddressOrRange *string `pulumi:"ipAddressOrRange"`
}





type IpAddressOrRangeInput interface {
	pulumi.Input

	ToIpAddressOrRangeOutput() IpAddressOrRangeOutput
	ToIpAddressOrRangeOutputWithContext(context.Context) IpAddressOrRangeOutput
}

type IpAddressOrRangeArgs struct {
	IpAddressOrRange pulumi.StringPtrInput `pulumi:"ipAddressOrRange"`
}

func (IpAddressOrRangeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRange)(nil)).Elem()
}

func (i IpAddressOrRangeArgs) ToIpAddressOrRangeOutput() IpAddressOrRangeOutput {
	return i.ToIpAddressOrRangeOutputWithContext(context.Background())
}

func (i IpAddressOrRangeArgs) ToIpAddressOrRangeOutputWithContext(ctx context.Context) IpAddressOrRangeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOrRangeOutput)
}





type IpAddressOrRangeArrayInput interface {
	pulumi.Input

	ToIpAddressOrRangeArrayOutput() IpAddressOrRangeArrayOutput
	ToIpAddressOrRangeArrayOutputWithContext(context.Context) IpAddressOrRangeArrayOutput
}

type IpAddressOrRangeArray []IpAddressOrRangeInput

func (IpAddressOrRangeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRange)(nil)).Elem()
}

func (i IpAddressOrRangeArray) ToIpAddressOrRangeArrayOutput() IpAddressOrRangeArrayOutput {
	return i.ToIpAddressOrRangeArrayOutputWithContext(context.Background())
}

func (i IpAddressOrRangeArray) ToIpAddressOrRangeArrayOutputWithContext(ctx context.Context) IpAddressOrRangeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOrRangeArrayOutput)
}

type IpAddressOrRangeOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRange)(nil)).Elem()
}

func (o IpAddressOrRangeOutput) ToIpAddressOrRangeOutput() IpAddressOrRangeOutput {
	return o
}

func (o IpAddressOrRangeOutput) ToIpAddressOrRangeOutputWithContext(ctx context.Context) IpAddressOrRangeOutput {
	return o
}

func (o IpAddressOrRangeOutput) IpAddressOrRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressOrRange) *string { return v.IpAddressOrRange }).(pulumi.StringPtrOutput)
}

type IpAddressOrRangeArrayOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRange)(nil)).Elem()
}

func (o IpAddressOrRangeArrayOutput) ToIpAddressOrRangeArrayOutput() IpAddressOrRangeArrayOutput {
	return o
}

func (o IpAddressOrRangeArrayOutput) ToIpAddressOrRangeArrayOutputWithContext(ctx context.Context) IpAddressOrRangeArrayOutput {
	return o
}

func (o IpAddressOrRangeArrayOutput) Index(i pulumi.IntInput) IpAddressOrRangeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressOrRange {
		return vs[0].([]IpAddressOrRange)[vs[1].(int)]
	}).(IpAddressOrRangeOutput)
}

type IpAddressOrRangeResponse struct {
	IpAddressOrRange *string `pulumi:"ipAddressOrRange"`
}





type IpAddressOrRangeResponseInput interface {
	pulumi.Input

	ToIpAddressOrRangeResponseOutput() IpAddressOrRangeResponseOutput
	ToIpAddressOrRangeResponseOutputWithContext(context.Context) IpAddressOrRangeResponseOutput
}

type IpAddressOrRangeResponseArgs struct {
	IpAddressOrRange pulumi.StringPtrInput `pulumi:"ipAddressOrRange"`
}

func (IpAddressOrRangeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRangeResponse)(nil)).Elem()
}

func (i IpAddressOrRangeResponseArgs) ToIpAddressOrRangeResponseOutput() IpAddressOrRangeResponseOutput {
	return i.ToIpAddressOrRangeResponseOutputWithContext(context.Background())
}

func (i IpAddressOrRangeResponseArgs) ToIpAddressOrRangeResponseOutputWithContext(ctx context.Context) IpAddressOrRangeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOrRangeResponseOutput)
}





type IpAddressOrRangeResponseArrayInput interface {
	pulumi.Input

	ToIpAddressOrRangeResponseArrayOutput() IpAddressOrRangeResponseArrayOutput
	ToIpAddressOrRangeResponseArrayOutputWithContext(context.Context) IpAddressOrRangeResponseArrayOutput
}

type IpAddressOrRangeResponseArray []IpAddressOrRangeResponseInput

func (IpAddressOrRangeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRangeResponse)(nil)).Elem()
}

func (i IpAddressOrRangeResponseArray) ToIpAddressOrRangeResponseArrayOutput() IpAddressOrRangeResponseArrayOutput {
	return i.ToIpAddressOrRangeResponseArrayOutputWithContext(context.Background())
}

func (i IpAddressOrRangeResponseArray) ToIpAddressOrRangeResponseArrayOutputWithContext(ctx context.Context) IpAddressOrRangeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpAddressOrRangeResponseArrayOutput)
}

type IpAddressOrRangeResponseOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpAddressOrRangeResponse)(nil)).Elem()
}

func (o IpAddressOrRangeResponseOutput) ToIpAddressOrRangeResponseOutput() IpAddressOrRangeResponseOutput {
	return o
}

func (o IpAddressOrRangeResponseOutput) ToIpAddressOrRangeResponseOutputWithContext(ctx context.Context) IpAddressOrRangeResponseOutput {
	return o
}

func (o IpAddressOrRangeResponseOutput) IpAddressOrRange() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpAddressOrRangeResponse) *string { return v.IpAddressOrRange }).(pulumi.StringPtrOutput)
}

type IpAddressOrRangeResponseArrayOutput struct{ *pulumi.OutputState }

func (IpAddressOrRangeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpAddressOrRangeResponse)(nil)).Elem()
}

func (o IpAddressOrRangeResponseArrayOutput) ToIpAddressOrRangeResponseArrayOutput() IpAddressOrRangeResponseArrayOutput {
	return o
}

func (o IpAddressOrRangeResponseArrayOutput) ToIpAddressOrRangeResponseArrayOutputWithContext(ctx context.Context) IpAddressOrRangeResponseArrayOutput {
	return o
}

func (o IpAddressOrRangeResponseArrayOutput) Index(i pulumi.IntInput) IpAddressOrRangeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpAddressOrRangeResponse {
		return vs[0].([]IpAddressOrRangeResponse)[vs[1].(int)]
	}).(IpAddressOrRangeResponseOutput)
}

type Location struct {
	FailoverPriority *int    `pulumi:"failoverPriority"`
	IsZoneRedundant  *bool   `pulumi:"isZoneRedundant"`
	LocationName     *string `pulumi:"locationName"`
}





type LocationInput interface {
	pulumi.Input

	ToLocationOutput() LocationOutput
	ToLocationOutputWithContext(context.Context) LocationOutput
}

type LocationArgs struct {
	FailoverPriority pulumi.IntPtrInput    `pulumi:"failoverPriority"`
	IsZoneRedundant  pulumi.BoolPtrInput   `pulumi:"isZoneRedundant"`
	LocationName     pulumi.StringPtrInput `pulumi:"locationName"`
}

func (LocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (i LocationArgs) ToLocationOutput() LocationOutput {
	return i.ToLocationOutputWithContext(context.Background())
}

func (i LocationArgs) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationOutput)
}





type LocationArrayInput interface {
	pulumi.Input

	ToLocationArrayOutput() LocationArrayOutput
	ToLocationArrayOutputWithContext(context.Context) LocationArrayOutput
}

type LocationArray []LocationInput

func (LocationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Location)(nil)).Elem()
}

func (i LocationArray) ToLocationArrayOutput() LocationArrayOutput {
	return i.ToLocationArrayOutputWithContext(context.Background())
}

func (i LocationArray) ToLocationArrayOutputWithContext(ctx context.Context) LocationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationArrayOutput)
}

type LocationOutput struct{ *pulumi.OutputState }

func (LocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Location)(nil)).Elem()
}

func (o LocationOutput) ToLocationOutput() LocationOutput {
	return o
}

func (o LocationOutput) ToLocationOutputWithContext(ctx context.Context) LocationOutput {
	return o
}

func (o LocationOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Location) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o LocationOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Location) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

func (o LocationOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Location) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

type LocationArrayOutput struct{ *pulumi.OutputState }

func (LocationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Location)(nil)).Elem()
}

func (o LocationArrayOutput) ToLocationArrayOutput() LocationArrayOutput {
	return o
}

func (o LocationArrayOutput) ToLocationArrayOutputWithContext(ctx context.Context) LocationArrayOutput {
	return o
}

func (o LocationArrayOutput) Index(i pulumi.IntInput) LocationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Location {
		return vs[0].([]Location)[vs[1].(int)]
	}).(LocationOutput)
}

type LocationResponse struct {
	DocumentEndpoint  string  `pulumi:"documentEndpoint"`
	FailoverPriority  *int    `pulumi:"failoverPriority"`
	Id                string  `pulumi:"id"`
	IsZoneRedundant   *bool   `pulumi:"isZoneRedundant"`
	LocationName      *string `pulumi:"locationName"`
	ProvisioningState string  `pulumi:"provisioningState"`
}





type LocationResponseInput interface {
	pulumi.Input

	ToLocationResponseOutput() LocationResponseOutput
	ToLocationResponseOutputWithContext(context.Context) LocationResponseOutput
}

type LocationResponseArgs struct {
	DocumentEndpoint  pulumi.StringInput    `pulumi:"documentEndpoint"`
	FailoverPriority  pulumi.IntPtrInput    `pulumi:"failoverPriority"`
	Id                pulumi.StringInput    `pulumi:"id"`
	IsZoneRedundant   pulumi.BoolPtrInput   `pulumi:"isZoneRedundant"`
	LocationName      pulumi.StringPtrInput `pulumi:"locationName"`
	ProvisioningState pulumi.StringInput    `pulumi:"provisioningState"`
}

func (LocationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationResponse)(nil)).Elem()
}

func (i LocationResponseArgs) ToLocationResponseOutput() LocationResponseOutput {
	return i.ToLocationResponseOutputWithContext(context.Background())
}

func (i LocationResponseArgs) ToLocationResponseOutputWithContext(ctx context.Context) LocationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationResponseOutput)
}





type LocationResponseArrayInput interface {
	pulumi.Input

	ToLocationResponseArrayOutput() LocationResponseArrayOutput
	ToLocationResponseArrayOutputWithContext(context.Context) LocationResponseArrayOutput
}

type LocationResponseArray []LocationResponseInput

func (LocationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocationResponse)(nil)).Elem()
}

func (i LocationResponseArray) ToLocationResponseArrayOutput() LocationResponseArrayOutput {
	return i.ToLocationResponseArrayOutputWithContext(context.Background())
}

func (i LocationResponseArray) ToLocationResponseArrayOutputWithContext(ctx context.Context) LocationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LocationResponseArrayOutput)
}

type LocationResponseOutput struct{ *pulumi.OutputState }

func (LocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LocationResponse)(nil)).Elem()
}

func (o LocationResponseOutput) ToLocationResponseOutput() LocationResponseOutput {
	return o
}

func (o LocationResponseOutput) ToLocationResponseOutputWithContext(ctx context.Context) LocationResponseOutput {
	return o
}

func (o LocationResponseOutput) DocumentEndpoint() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.DocumentEndpoint }).(pulumi.StringOutput)
}

func (o LocationResponseOutput) FailoverPriority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LocationResponse) *int { return v.FailoverPriority }).(pulumi.IntPtrOutput)
}

func (o LocationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o LocationResponseOutput) IsZoneRedundant() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v LocationResponse) *bool { return v.IsZoneRedundant }).(pulumi.BoolPtrOutput)
}

func (o LocationResponseOutput) LocationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LocationResponse) *string { return v.LocationName }).(pulumi.StringPtrOutput)
}

func (o LocationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v LocationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

type LocationResponseArrayOutput struct{ *pulumi.OutputState }

func (LocationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LocationResponse)(nil)).Elem()
}

func (o LocationResponseArrayOutput) ToLocationResponseArrayOutput() LocationResponseArrayOutput {
	return o
}

func (o LocationResponseArrayOutput) ToLocationResponseArrayOutputWithContext(ctx context.Context) LocationResponseArrayOutput {
	return o
}

func (o LocationResponseArrayOutput) Index(i pulumi.IntInput) LocationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LocationResponse {
		return vs[0].([]LocationResponse)[vs[1].(int)]
	}).(LocationResponseOutput)
}

type ManagedCassandraManagedServiceIdentity struct {
	Type *string `pulumi:"type"`
}





type ManagedCassandraManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedCassandraManagedServiceIdentityOutput() ManagedCassandraManagedServiceIdentityOutput
	ToManagedCassandraManagedServiceIdentityOutputWithContext(context.Context) ManagedCassandraManagedServiceIdentityOutput
}

type ManagedCassandraManagedServiceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagedCassandraManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCassandraManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedCassandraManagedServiceIdentityArgs) ToManagedCassandraManagedServiceIdentityOutput() ManagedCassandraManagedServiceIdentityOutput {
	return i.ToManagedCassandraManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedCassandraManagedServiceIdentityArgs) ToManagedCassandraManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCassandraManagedServiceIdentityOutput)
}

func (i ManagedCassandraManagedServiceIdentityArgs) ToManagedCassandraManagedServiceIdentityPtrOutput() ManagedCassandraManagedServiceIdentityPtrOutput {
	return i.ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedCassandraManagedServiceIdentityArgs) ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCassandraManagedServiceIdentityOutput).ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedCassandraManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedCassandraManagedServiceIdentityPtrOutput() ManagedCassandraManagedServiceIdentityPtrOutput
	ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedCassandraManagedServiceIdentityPtrOutput
}

type managedCassandraManagedServiceIdentityPtrType ManagedCassandraManagedServiceIdentityArgs

func ManagedCassandraManagedServiceIdentityPtr(v *ManagedCassandraManagedServiceIdentityArgs) ManagedCassandraManagedServiceIdentityPtrInput {
	return (*managedCassandraManagedServiceIdentityPtrType)(v)
}

func (*managedCassandraManagedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCassandraManagedServiceIdentity)(nil)).Elem()
}

func (i *managedCassandraManagedServiceIdentityPtrType) ToManagedCassandraManagedServiceIdentityPtrOutput() ManagedCassandraManagedServiceIdentityPtrOutput {
	return i.ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedCassandraManagedServiceIdentityPtrType) ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCassandraManagedServiceIdentityPtrOutput)
}

type ManagedCassandraManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedCassandraManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCassandraManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedCassandraManagedServiceIdentityOutput) ToManagedCassandraManagedServiceIdentityOutput() ManagedCassandraManagedServiceIdentityOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityOutput) ToManagedCassandraManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityOutput) ToManagedCassandraManagedServiceIdentityPtrOutput() ManagedCassandraManagedServiceIdentityPtrOutput {
	return o.ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedCassandraManagedServiceIdentityOutput) ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedCassandraManagedServiceIdentity) *ManagedCassandraManagedServiceIdentity {
		return &v
	}).(ManagedCassandraManagedServiceIdentityPtrOutput)
}

func (o ManagedCassandraManagedServiceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedCassandraManagedServiceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedCassandraManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedCassandraManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCassandraManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedCassandraManagedServiceIdentityPtrOutput) ToManagedCassandraManagedServiceIdentityPtrOutput() ManagedCassandraManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityPtrOutput) ToManagedCassandraManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityPtrOutput) Elem() ManagedCassandraManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedCassandraManagedServiceIdentity) ManagedCassandraManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedCassandraManagedServiceIdentity
		return ret
	}).(ManagedCassandraManagedServiceIdentityOutput)
}

func (o ManagedCassandraManagedServiceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCassandraManagedServiceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedCassandraManagedServiceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}





type ManagedCassandraManagedServiceIdentityResponseInput interface {
	pulumi.Input

	ToManagedCassandraManagedServiceIdentityResponseOutput() ManagedCassandraManagedServiceIdentityResponseOutput
	ToManagedCassandraManagedServiceIdentityResponseOutputWithContext(context.Context) ManagedCassandraManagedServiceIdentityResponseOutput
}

type ManagedCassandraManagedServiceIdentityResponseArgs struct {
	PrincipalId pulumi.StringInput    `pulumi:"principalId"`
	TenantId    pulumi.StringInput    `pulumi:"tenantId"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (ManagedCassandraManagedServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCassandraManagedServiceIdentityResponse)(nil)).Elem()
}

func (i ManagedCassandraManagedServiceIdentityResponseArgs) ToManagedCassandraManagedServiceIdentityResponseOutput() ManagedCassandraManagedServiceIdentityResponseOutput {
	return i.ToManagedCassandraManagedServiceIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedCassandraManagedServiceIdentityResponseArgs) ToManagedCassandraManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCassandraManagedServiceIdentityResponseOutput)
}

func (i ManagedCassandraManagedServiceIdentityResponseArgs) ToManagedCassandraManagedServiceIdentityResponsePtrOutput() ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedCassandraManagedServiceIdentityResponseArgs) ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCassandraManagedServiceIdentityResponseOutput).ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedCassandraManagedServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedCassandraManagedServiceIdentityResponsePtrOutput() ManagedCassandraManagedServiceIdentityResponsePtrOutput
	ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(context.Context) ManagedCassandraManagedServiceIdentityResponsePtrOutput
}

type managedCassandraManagedServiceIdentityResponsePtrType ManagedCassandraManagedServiceIdentityResponseArgs

func ManagedCassandraManagedServiceIdentityResponsePtr(v *ManagedCassandraManagedServiceIdentityResponseArgs) ManagedCassandraManagedServiceIdentityResponsePtrInput {
	return (*managedCassandraManagedServiceIdentityResponsePtrType)(v)
}

func (*managedCassandraManagedServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCassandraManagedServiceIdentityResponse)(nil)).Elem()
}

func (i *managedCassandraManagedServiceIdentityResponsePtrType) ToManagedCassandraManagedServiceIdentityResponsePtrOutput() ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedCassandraManagedServiceIdentityResponsePtrType) ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedCassandraManagedServiceIdentityResponsePtrOutput)
}

type ManagedCassandraManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedCassandraManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedCassandraManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) ToManagedCassandraManagedServiceIdentityResponseOutput() ManagedCassandraManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) ToManagedCassandraManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) ToManagedCassandraManagedServiceIdentityResponsePtrOutput() ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return o.ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedCassandraManagedServiceIdentityResponse) *ManagedCassandraManagedServiceIdentityResponse {
		return &v
	}).(ManagedCassandraManagedServiceIdentityResponsePtrOutput)
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedCassandraManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedCassandraManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedCassandraManagedServiceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedCassandraManagedServiceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ManagedCassandraManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedCassandraManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedCassandraManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedCassandraManagedServiceIdentityResponsePtrOutput) ToManagedCassandraManagedServiceIdentityResponsePtrOutput() ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityResponsePtrOutput) ToManagedCassandraManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedCassandraManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedCassandraManagedServiceIdentityResponsePtrOutput) Elem() ManagedCassandraManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedCassandraManagedServiceIdentityResponse) ManagedCassandraManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedCassandraManagedServiceIdentityResponse
		return ret
	}).(ManagedCassandraManagedServiceIdentityResponseOutput)
}

func (o ManagedCassandraManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCassandraManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedCassandraManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCassandraManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedCassandraManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedCassandraManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ManagedServiceIdentity struct {
	Type                   *ResourceIdentityType  `pulumi:"type"`
	UserAssignedIdentities map[string]interface{} `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   ResourceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput              `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return i.ToManagedServiceIdentityOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput)
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityArgs) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityOutput).ToManagedServiceIdentityPtrOutputWithContext(ctx)
}









type ManagedServiceIdentityPtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput
	ToManagedServiceIdentityPtrOutputWithContext(context.Context) ManagedServiceIdentityPtrOutput
}

type managedServiceIdentityPtrType ManagedServiceIdentityArgs

func ManagedServiceIdentityPtr(v *ManagedServiceIdentityArgs) ManagedServiceIdentityPtrInput {
	return (*managedServiceIdentityPtrType)(v)
}

func (*managedServiceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return i.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityPtrType) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityPtrOutput)
}

type ManagedServiceIdentityOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityOutputWithContext(ctx context.Context) ManagedServiceIdentityOutput {
	return o
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o.ToManagedServiceIdentityPtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentity) *ManagedServiceIdentity {
		return &v
	}).(ManagedServiceIdentityPtrOutput)
}

func (o ManagedServiceIdentityOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *ResourceIdentityType { return v.Type }).(ResourceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) map[string]interface{} { return v.UserAssignedIdentities }).(pulumi.MapOutput)
}

type ManagedServiceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentity)(nil)).Elem()
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutput() ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) ToManagedServiceIdentityPtrOutputWithContext(ctx context.Context) ManagedServiceIdentityPtrOutput {
	return o
}

func (o ManagedServiceIdentityPtrOutput) Elem() ManagedServiceIdentityOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) ManagedServiceIdentity {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentity
		return ret
	}).(ManagedServiceIdentityOutput)
}

func (o ManagedServiceIdentityPtrOutput) Type() ResourceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *ResourceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ResourceIdentityTypePtrOutput)
}

func (o ManagedServiceIdentityPtrOutput) UserAssignedIdentities() pulumi.MapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) map[string]interface{} {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(pulumi.MapOutput)
}

type ManagedServiceIdentityResponse struct {
	PrincipalId            string                                                          `pulumi:"principalId"`
	TenantId               string                                                          `pulumi:"tenantId"`
	Type                   *string                                                         `pulumi:"type"`
	UserAssignedIdentities map[string]ManagedServiceIdentityResponseUserAssignedIdentities `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityResponseInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput
	ToManagedServiceIdentityResponseOutputWithContext(context.Context) ManagedServiceIdentityResponseOutput
}

type ManagedServiceIdentityResponseArgs struct {
	PrincipalId            pulumi.StringInput                                           `pulumi:"principalId"`
	TenantId               pulumi.StringInput                                           `pulumi:"tenantId"`
	Type                   pulumi.StringPtrInput                                        `pulumi:"type"`
	UserAssignedIdentities ManagedServiceIdentityResponseUserAssignedIdentitiesMapInput `pulumi:"userAssignedIdentities"`
}

func (ManagedServiceIdentityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return i.ToManagedServiceIdentityResponseOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseOutput)
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseArgs) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseOutput).ToManagedServiceIdentityResponsePtrOutputWithContext(ctx)
}









type ManagedServiceIdentityResponsePtrInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput
	ToManagedServiceIdentityResponsePtrOutputWithContext(context.Context) ManagedServiceIdentityResponsePtrOutput
}

type managedServiceIdentityResponsePtrType ManagedServiceIdentityResponseArgs

func ManagedServiceIdentityResponsePtr(v *ManagedServiceIdentityResponseArgs) ManagedServiceIdentityResponsePtrInput {
	return (*managedServiceIdentityResponsePtrType)(v)
}

func (*managedServiceIdentityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (i *managedServiceIdentityResponsePtrType) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return i.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (i *managedServiceIdentityResponsePtrType) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponsePtrOutput)
}

type ManagedServiceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutput() ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponseOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseOutput {
	return o
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o.ToManagedServiceIdentityResponsePtrOutputWithContext(context.Background())
}

func (o ManagedServiceIdentityResponseOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ManagedServiceIdentityResponse) *ManagedServiceIdentityResponse {
		return &v
	}).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]ManagedServiceIdentityResponseUserAssignedIdentities {
		return v.UserAssignedIdentities
	}).(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedServiceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ManagedServiceIdentityResponse)(nil)).Elem()
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutput() ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) ToManagedServiceIdentityResponsePtrOutputWithContext(ctx context.Context) ManagedServiceIdentityResponsePtrOutput {
	return o
}

func (o ManagedServiceIdentityResponsePtrOutput) Elem() ManagedServiceIdentityResponseOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) ManagedServiceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ManagedServiceIdentityResponse
		return ret
	}).(ManagedServiceIdentityResponseOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]ManagedServiceIdentityResponseUserAssignedIdentities {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedServiceIdentityResponseUserAssignedIdentities struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}





type ManagedServiceIdentityResponseUserAssignedIdentitiesInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseUserAssignedIdentitiesOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesOutput
	ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesArgs struct {
	ClientId    pulumi.StringInput `pulumi:"clientId"`
	PrincipalId pulumi.StringInput `pulumi:"principalId"`
}

func (ManagedServiceIdentityResponseUserAssignedIdentitiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesArgs) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return i.ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesArgs) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput)
}





type ManagedServiceIdentityResponseUserAssignedIdentitiesMapInput interface {
	pulumi.Input

	ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput
	ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesMap map[string]ManagedServiceIdentityResponseUserAssignedIdentitiesInput

func (ManagedServiceIdentityResponseUserAssignedIdentitiesMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesMap) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return i.ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(context.Background())
}

func (i ManagedServiceIdentityResponseUserAssignedIdentitiesMap) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput)
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponseUserAssignedIdentities) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponseUserAssignedIdentities) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput struct{ *pulumi.OutputState }

func (ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ManagedServiceIdentityResponseUserAssignedIdentities)(nil)).Elem()
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput() ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) ToManagedServiceIdentityResponseUserAssignedIdentitiesMapOutputWithContext(ctx context.Context) ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput {
	return o
}

func (o ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput) MapIndex(k pulumi.StringInput) ManagedServiceIdentityResponseUserAssignedIdentitiesOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ManagedServiceIdentityResponseUserAssignedIdentities {
		return vs[0].(map[string]ManagedServiceIdentityResponseUserAssignedIdentities)[vs[1].(string)]
	}).(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput)
}

type MongoDBCollectionGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type MongoDBCollectionGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToMongoDBCollectionGetPropertiesResponseOptionsOutput() MongoDBCollectionGetPropertiesResponseOptionsOutput
	ToMongoDBCollectionGetPropertiesResponseOptionsOutputWithContext(context.Context) MongoDBCollectionGetPropertiesResponseOptionsOutput
}

type MongoDBCollectionGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (MongoDBCollectionGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionGetPropertiesResponseOptions)(nil)).Elem()
}

func (i MongoDBCollectionGetPropertiesResponseOptionsArgs) ToMongoDBCollectionGetPropertiesResponseOptionsOutput() MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return i.ToMongoDBCollectionGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i MongoDBCollectionGetPropertiesResponseOptionsArgs) ToMongoDBCollectionGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionGetPropertiesResponseOptionsOutput)
}

func (i MongoDBCollectionGetPropertiesResponseOptionsArgs) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutput() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return i.ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i MongoDBCollectionGetPropertiesResponseOptionsArgs) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionGetPropertiesResponseOptionsOutput).ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type MongoDBCollectionGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutput() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput
	ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput
}

type mongoDBCollectionGetPropertiesResponseOptionsPtrType MongoDBCollectionGetPropertiesResponseOptionsArgs

func MongoDBCollectionGetPropertiesResponseOptionsPtr(v *MongoDBCollectionGetPropertiesResponseOptionsArgs) MongoDBCollectionGetPropertiesResponseOptionsPtrInput {
	return (*mongoDBCollectionGetPropertiesResponseOptionsPtrType)(v)
}

func (*mongoDBCollectionGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *mongoDBCollectionGetPropertiesResponseOptionsPtrType) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutput() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return i.ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *mongoDBCollectionGetPropertiesResponseOptionsPtrType) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput)
}

type MongoDBCollectionGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) ToMongoDBCollectionGetPropertiesResponseOptionsOutput() MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) ToMongoDBCollectionGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutput() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o.ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDBCollectionGetPropertiesResponseOptions) *MongoDBCollectionGetPropertiesResponseOptions {
		return &v
	}).(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type MongoDBCollectionGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutput() MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) ToMongoDBCollectionGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) Elem() MongoDBCollectionGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseOptions) MongoDBCollectionGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret MongoDBCollectionGetPropertiesResponseOptions
		return ret
	}).(MongoDBCollectionGetPropertiesResponseOptionsOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type MongoDBCollectionGetPropertiesResponseResource struct {
	AnalyticalStorageTtl *int                 `pulumi:"analyticalStorageTtl"`
	Etag                 string               `pulumi:"etag"`
	Id                   string               `pulumi:"id"`
	Indexes              []MongoIndexResponse `pulumi:"indexes"`
	Rid                  string               `pulumi:"rid"`
	ShardKey             map[string]string    `pulumi:"shardKey"`
	Ts                   float64              `pulumi:"ts"`
}





type MongoDBCollectionGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToMongoDBCollectionGetPropertiesResponseResourceOutput() MongoDBCollectionGetPropertiesResponseResourceOutput
	ToMongoDBCollectionGetPropertiesResponseResourceOutputWithContext(context.Context) MongoDBCollectionGetPropertiesResponseResourceOutput
}

type MongoDBCollectionGetPropertiesResponseResourceArgs struct {
	AnalyticalStorageTtl pulumi.IntPtrInput           `pulumi:"analyticalStorageTtl"`
	Etag                 pulumi.StringInput           `pulumi:"etag"`
	Id                   pulumi.StringInput           `pulumi:"id"`
	Indexes              MongoIndexResponseArrayInput `pulumi:"indexes"`
	Rid                  pulumi.StringInput           `pulumi:"rid"`
	ShardKey             pulumi.StringMapInput        `pulumi:"shardKey"`
	Ts                   pulumi.Float64Input          `pulumi:"ts"`
}

func (MongoDBCollectionGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionGetPropertiesResponseResource)(nil)).Elem()
}

func (i MongoDBCollectionGetPropertiesResponseResourceArgs) ToMongoDBCollectionGetPropertiesResponseResourceOutput() MongoDBCollectionGetPropertiesResponseResourceOutput {
	return i.ToMongoDBCollectionGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i MongoDBCollectionGetPropertiesResponseResourceArgs) ToMongoDBCollectionGetPropertiesResponseResourceOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionGetPropertiesResponseResourceOutput)
}

func (i MongoDBCollectionGetPropertiesResponseResourceArgs) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutput() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return i.ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i MongoDBCollectionGetPropertiesResponseResourceArgs) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionGetPropertiesResponseResourceOutput).ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type MongoDBCollectionGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToMongoDBCollectionGetPropertiesResponseResourcePtrOutput() MongoDBCollectionGetPropertiesResponseResourcePtrOutput
	ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(context.Context) MongoDBCollectionGetPropertiesResponseResourcePtrOutput
}

type mongoDBCollectionGetPropertiesResponseResourcePtrType MongoDBCollectionGetPropertiesResponseResourceArgs

func MongoDBCollectionGetPropertiesResponseResourcePtr(v *MongoDBCollectionGetPropertiesResponseResourceArgs) MongoDBCollectionGetPropertiesResponseResourcePtrInput {
	return (*mongoDBCollectionGetPropertiesResponseResourcePtrType)(v)
}

func (*mongoDBCollectionGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionGetPropertiesResponseResource)(nil)).Elem()
}

func (i *mongoDBCollectionGetPropertiesResponseResourcePtrType) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutput() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return i.ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *mongoDBCollectionGetPropertiesResponseResourcePtrType) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionGetPropertiesResponseResourcePtrOutput)
}

type MongoDBCollectionGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ToMongoDBCollectionGetPropertiesResponseResourceOutput() MongoDBCollectionGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ToMongoDBCollectionGetPropertiesResponseResourceOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutput() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o.ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDBCollectionGetPropertiesResponseResource) *MongoDBCollectionGetPropertiesResponseResource {
		return &v
	}).(MongoDBCollectionGetPropertiesResponseResourcePtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Indexes() MongoIndexResponseArrayOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) []MongoIndexResponse { return v.Indexes }).(MongoIndexResponseArrayOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) map[string]string { return v.ShardKey }).(pulumi.StringMapOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v MongoDBCollectionGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type MongoDBCollectionGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutput() MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ToMongoDBCollectionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Elem() MongoDBCollectionGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) MongoDBCollectionGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret MongoDBCollectionGetPropertiesResponseResource
		return ret
	}).(MongoDBCollectionGetPropertiesResponseResourceOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Indexes() MongoIndexResponseArrayOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) []MongoIndexResponse {
		if v == nil {
			return nil
		}
		return v.Indexes
	}).(MongoIndexResponseArrayOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) map[string]string {
		if v == nil {
			return nil
		}
		return v.ShardKey
	}).(pulumi.StringMapOutput)
}

func (o MongoDBCollectionGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type MongoDBCollectionResource struct {
	AnalyticalStorageTtl *int              `pulumi:"analyticalStorageTtl"`
	Id                   string            `pulumi:"id"`
	Indexes              []MongoIndex      `pulumi:"indexes"`
	ShardKey             map[string]string `pulumi:"shardKey"`
}





type MongoDBCollectionResourceInput interface {
	pulumi.Input

	ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput
	ToMongoDBCollectionResourceOutputWithContext(context.Context) MongoDBCollectionResourceOutput
}

type MongoDBCollectionResourceArgs struct {
	AnalyticalStorageTtl pulumi.IntPtrInput    `pulumi:"analyticalStorageTtl"`
	Id                   pulumi.StringInput    `pulumi:"id"`
	Indexes              MongoIndexArrayInput  `pulumi:"indexes"`
	ShardKey             pulumi.StringMapInput `pulumi:"shardKey"`
}

func (MongoDBCollectionResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionResource)(nil)).Elem()
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput {
	return i.ToMongoDBCollectionResourceOutputWithContext(context.Background())
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourceOutputWithContext(ctx context.Context) MongoDBCollectionResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionResourceOutput)
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourcePtrOutput() MongoDBCollectionResourcePtrOutput {
	return i.ToMongoDBCollectionResourcePtrOutputWithContext(context.Background())
}

func (i MongoDBCollectionResourceArgs) ToMongoDBCollectionResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionResourceOutput).ToMongoDBCollectionResourcePtrOutputWithContext(ctx)
}









type MongoDBCollectionResourcePtrInput interface {
	pulumi.Input

	ToMongoDBCollectionResourcePtrOutput() MongoDBCollectionResourcePtrOutput
	ToMongoDBCollectionResourcePtrOutputWithContext(context.Context) MongoDBCollectionResourcePtrOutput
}

type mongoDBCollectionResourcePtrType MongoDBCollectionResourceArgs

func MongoDBCollectionResourcePtr(v *MongoDBCollectionResourceArgs) MongoDBCollectionResourcePtrInput {
	return (*mongoDBCollectionResourcePtrType)(v)
}

func (*mongoDBCollectionResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionResource)(nil)).Elem()
}

func (i *mongoDBCollectionResourcePtrType) ToMongoDBCollectionResourcePtrOutput() MongoDBCollectionResourcePtrOutput {
	return i.ToMongoDBCollectionResourcePtrOutputWithContext(context.Background())
}

func (i *mongoDBCollectionResourcePtrType) ToMongoDBCollectionResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBCollectionResourcePtrOutput)
}

type MongoDBCollectionResourceOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBCollectionResource)(nil)).Elem()
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourceOutput() MongoDBCollectionResourceOutput {
	return o
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourceOutputWithContext(ctx context.Context) MongoDBCollectionResourceOutput {
	return o
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourcePtrOutput() MongoDBCollectionResourcePtrOutput {
	return o.ToMongoDBCollectionResourcePtrOutputWithContext(context.Background())
}

func (o MongoDBCollectionResourceOutput) ToMongoDBCollectionResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDBCollectionResource) *MongoDBCollectionResource {
		return &v
	}).(MongoDBCollectionResourcePtrOutput)
}

func (o MongoDBCollectionResourceOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) *int { return v.AnalyticalStorageTtl }).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBCollectionResourceOutput) Indexes() MongoIndexArrayOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) []MongoIndex { return v.Indexes }).(MongoIndexArrayOutput)
}

func (o MongoDBCollectionResourceOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v MongoDBCollectionResource) map[string]string { return v.ShardKey }).(pulumi.StringMapOutput)
}

type MongoDBCollectionResourcePtrOutput struct{ *pulumi.OutputState }

func (MongoDBCollectionResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBCollectionResource)(nil)).Elem()
}

func (o MongoDBCollectionResourcePtrOutput) ToMongoDBCollectionResourcePtrOutput() MongoDBCollectionResourcePtrOutput {
	return o
}

func (o MongoDBCollectionResourcePtrOutput) ToMongoDBCollectionResourcePtrOutputWithContext(ctx context.Context) MongoDBCollectionResourcePtrOutput {
	return o
}

func (o MongoDBCollectionResourcePtrOutput) Elem() MongoDBCollectionResourceOutput {
	return o.ApplyT(func(v *MongoDBCollectionResource) MongoDBCollectionResource {
		if v != nil {
			return *v
		}
		var ret MongoDBCollectionResource
		return ret
	}).(MongoDBCollectionResourceOutput)
}

func (o MongoDBCollectionResourcePtrOutput) AnalyticalStorageTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionResource) *int {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.IntPtrOutput)
}

func (o MongoDBCollectionResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBCollectionResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBCollectionResourcePtrOutput) Indexes() MongoIndexArrayOutput {
	return o.ApplyT(func(v *MongoDBCollectionResource) []MongoIndex {
		if v == nil {
			return nil
		}
		return v.Indexes
	}).(MongoIndexArrayOutput)
}

func (o MongoDBCollectionResourcePtrOutput) ShardKey() pulumi.StringMapOutput {
	return o.ApplyT(func(v *MongoDBCollectionResource) map[string]string {
		if v == nil {
			return nil
		}
		return v.ShardKey
	}).(pulumi.StringMapOutput)
}

type MongoDBDatabaseGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type MongoDBDatabaseGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToMongoDBDatabaseGetPropertiesResponseOptionsOutput() MongoDBDatabaseGetPropertiesResponseOptionsOutput
	ToMongoDBDatabaseGetPropertiesResponseOptionsOutputWithContext(context.Context) MongoDBDatabaseGetPropertiesResponseOptionsOutput
}

type MongoDBDatabaseGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (MongoDBDatabaseGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (i MongoDBDatabaseGetPropertiesResponseOptionsArgs) ToMongoDBDatabaseGetPropertiesResponseOptionsOutput() MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return i.ToMongoDBDatabaseGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i MongoDBDatabaseGetPropertiesResponseOptionsArgs) ToMongoDBDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseGetPropertiesResponseOptionsOutput)
}

func (i MongoDBDatabaseGetPropertiesResponseOptionsArgs) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutput() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return i.ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i MongoDBDatabaseGetPropertiesResponseOptionsArgs) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseGetPropertiesResponseOptionsOutput).ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type MongoDBDatabaseGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutput() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput
	ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput
}

type mongoDBDatabaseGetPropertiesResponseOptionsPtrType MongoDBDatabaseGetPropertiesResponseOptionsArgs

func MongoDBDatabaseGetPropertiesResponseOptionsPtr(v *MongoDBDatabaseGetPropertiesResponseOptionsArgs) MongoDBDatabaseGetPropertiesResponseOptionsPtrInput {
	return (*mongoDBDatabaseGetPropertiesResponseOptionsPtrType)(v)
}

func (*mongoDBDatabaseGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *mongoDBDatabaseGetPropertiesResponseOptionsPtrType) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutput() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return i.ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *mongoDBDatabaseGetPropertiesResponseOptionsPtrType) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput)
}

type MongoDBDatabaseGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsOutput() MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutput() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDBDatabaseGetPropertiesResponseOptions) *MongoDBDatabaseGetPropertiesResponseOptions {
		return &v
	}).(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutput() MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) ToMongoDBDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) Elem() MongoDBDatabaseGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseOptions) MongoDBDatabaseGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret MongoDBDatabaseGetPropertiesResponseOptions
		return ret
	}).(MongoDBDatabaseGetPropertiesResponseOptionsOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type MongoDBDatabaseGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}





type MongoDBDatabaseGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToMongoDBDatabaseGetPropertiesResponseResourceOutput() MongoDBDatabaseGetPropertiesResponseResourceOutput
	ToMongoDBDatabaseGetPropertiesResponseResourceOutputWithContext(context.Context) MongoDBDatabaseGetPropertiesResponseResourceOutput
}

type MongoDBDatabaseGetPropertiesResponseResourceArgs struct {
	Etag pulumi.StringInput  `pulumi:"etag"`
	Id   pulumi.StringInput  `pulumi:"id"`
	Rid  pulumi.StringInput  `pulumi:"rid"`
	Ts   pulumi.Float64Input `pulumi:"ts"`
}

func (MongoDBDatabaseGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (i MongoDBDatabaseGetPropertiesResponseResourceArgs) ToMongoDBDatabaseGetPropertiesResponseResourceOutput() MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return i.ToMongoDBDatabaseGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i MongoDBDatabaseGetPropertiesResponseResourceArgs) ToMongoDBDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseGetPropertiesResponseResourceOutput)
}

func (i MongoDBDatabaseGetPropertiesResponseResourceArgs) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutput() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return i.ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i MongoDBDatabaseGetPropertiesResponseResourceArgs) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseGetPropertiesResponseResourceOutput).ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type MongoDBDatabaseGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutput() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput
	ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Context) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput
}

type mongoDBDatabaseGetPropertiesResponseResourcePtrType MongoDBDatabaseGetPropertiesResponseResourceArgs

func MongoDBDatabaseGetPropertiesResponseResourcePtr(v *MongoDBDatabaseGetPropertiesResponseResourceArgs) MongoDBDatabaseGetPropertiesResponseResourcePtrInput {
	return (*mongoDBDatabaseGetPropertiesResponseResourcePtrType)(v)
}

func (*mongoDBDatabaseGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (i *mongoDBDatabaseGetPropertiesResponseResourcePtrType) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutput() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return i.ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *mongoDBDatabaseGetPropertiesResponseResourcePtrType) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput)
}

type MongoDBDatabaseGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) ToMongoDBDatabaseGetPropertiesResponseResourceOutput() MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) ToMongoDBDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutput() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDBDatabaseGetPropertiesResponseResource) *MongoDBDatabaseGetPropertiesResponseResource {
		return &v
	}).(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v MongoDBDatabaseGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type MongoDBDatabaseGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutput() MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) ToMongoDBDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Elem() MongoDBDatabaseGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) MongoDBDatabaseGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret MongoDBDatabaseGetPropertiesResponseResource
		return ret
	}).(MongoDBDatabaseGetPropertiesResponseResourceOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o MongoDBDatabaseGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type MongoDBDatabaseResource struct {
	Id string `pulumi:"id"`
}





type MongoDBDatabaseResourceInput interface {
	pulumi.Input

	ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput
	ToMongoDBDatabaseResourceOutputWithContext(context.Context) MongoDBDatabaseResourceOutput
}

type MongoDBDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (MongoDBDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseResource)(nil)).Elem()
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput {
	return i.ToMongoDBDatabaseResourceOutputWithContext(context.Background())
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseResourceOutput)
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourcePtrOutput() MongoDBDatabaseResourcePtrOutput {
	return i.ToMongoDBDatabaseResourcePtrOutputWithContext(context.Background())
}

func (i MongoDBDatabaseResourceArgs) ToMongoDBDatabaseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseResourceOutput).ToMongoDBDatabaseResourcePtrOutputWithContext(ctx)
}









type MongoDBDatabaseResourcePtrInput interface {
	pulumi.Input

	ToMongoDBDatabaseResourcePtrOutput() MongoDBDatabaseResourcePtrOutput
	ToMongoDBDatabaseResourcePtrOutputWithContext(context.Context) MongoDBDatabaseResourcePtrOutput
}

type mongoDBDatabaseResourcePtrType MongoDBDatabaseResourceArgs

func MongoDBDatabaseResourcePtr(v *MongoDBDatabaseResourceArgs) MongoDBDatabaseResourcePtrInput {
	return (*mongoDBDatabaseResourcePtrType)(v)
}

func (*mongoDBDatabaseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseResource)(nil)).Elem()
}

func (i *mongoDBDatabaseResourcePtrType) ToMongoDBDatabaseResourcePtrOutput() MongoDBDatabaseResourcePtrOutput {
	return i.ToMongoDBDatabaseResourcePtrOutputWithContext(context.Background())
}

func (i *mongoDBDatabaseResourcePtrType) ToMongoDBDatabaseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoDBDatabaseResourcePtrOutput)
}

type MongoDBDatabaseResourceOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoDBDatabaseResource)(nil)).Elem()
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourceOutput() MongoDBDatabaseResourceOutput {
	return o
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourceOutputWithContext(ctx context.Context) MongoDBDatabaseResourceOutput {
	return o
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourcePtrOutput() MongoDBDatabaseResourcePtrOutput {
	return o.ToMongoDBDatabaseResourcePtrOutputWithContext(context.Background())
}

func (o MongoDBDatabaseResourceOutput) ToMongoDBDatabaseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoDBDatabaseResource) *MongoDBDatabaseResource {
		return &v
	}).(MongoDBDatabaseResourcePtrOutput)
}

func (o MongoDBDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v MongoDBDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type MongoDBDatabaseResourcePtrOutput struct{ *pulumi.OutputState }

func (MongoDBDatabaseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoDBDatabaseResource)(nil)).Elem()
}

func (o MongoDBDatabaseResourcePtrOutput) ToMongoDBDatabaseResourcePtrOutput() MongoDBDatabaseResourcePtrOutput {
	return o
}

func (o MongoDBDatabaseResourcePtrOutput) ToMongoDBDatabaseResourcePtrOutputWithContext(ctx context.Context) MongoDBDatabaseResourcePtrOutput {
	return o
}

func (o MongoDBDatabaseResourcePtrOutput) Elem() MongoDBDatabaseResourceOutput {
	return o.ApplyT(func(v *MongoDBDatabaseResource) MongoDBDatabaseResource {
		if v != nil {
			return *v
		}
		var ret MongoDBDatabaseResource
		return ret
	}).(MongoDBDatabaseResourceOutput)
}

func (o MongoDBDatabaseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MongoDBDatabaseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type MongoIndex struct {
	Key     *MongoIndexKeys    `pulumi:"key"`
	Options *MongoIndexOptions `pulumi:"options"`
}





type MongoIndexInput interface {
	pulumi.Input

	ToMongoIndexOutput() MongoIndexOutput
	ToMongoIndexOutputWithContext(context.Context) MongoIndexOutput
}

type MongoIndexArgs struct {
	Key     MongoIndexKeysPtrInput    `pulumi:"key"`
	Options MongoIndexOptionsPtrInput `pulumi:"options"`
}

func (MongoIndexArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndex)(nil)).Elem()
}

func (i MongoIndexArgs) ToMongoIndexOutput() MongoIndexOutput {
	return i.ToMongoIndexOutputWithContext(context.Background())
}

func (i MongoIndexArgs) ToMongoIndexOutputWithContext(ctx context.Context) MongoIndexOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOutput)
}





type MongoIndexArrayInput interface {
	pulumi.Input

	ToMongoIndexArrayOutput() MongoIndexArrayOutput
	ToMongoIndexArrayOutputWithContext(context.Context) MongoIndexArrayOutput
}

type MongoIndexArray []MongoIndexInput

func (MongoIndexArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndex)(nil)).Elem()
}

func (i MongoIndexArray) ToMongoIndexArrayOutput() MongoIndexArrayOutput {
	return i.ToMongoIndexArrayOutputWithContext(context.Background())
}

func (i MongoIndexArray) ToMongoIndexArrayOutputWithContext(ctx context.Context) MongoIndexArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexArrayOutput)
}

type MongoIndexOutput struct{ *pulumi.OutputState }

func (MongoIndexOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndex)(nil)).Elem()
}

func (o MongoIndexOutput) ToMongoIndexOutput() MongoIndexOutput {
	return o
}

func (o MongoIndexOutput) ToMongoIndexOutputWithContext(ctx context.Context) MongoIndexOutput {
	return o
}

func (o MongoIndexOutput) Key() MongoIndexKeysPtrOutput {
	return o.ApplyT(func(v MongoIndex) *MongoIndexKeys { return v.Key }).(MongoIndexKeysPtrOutput)
}

func (o MongoIndexOutput) Options() MongoIndexOptionsPtrOutput {
	return o.ApplyT(func(v MongoIndex) *MongoIndexOptions { return v.Options }).(MongoIndexOptionsPtrOutput)
}

type MongoIndexArrayOutput struct{ *pulumi.OutputState }

func (MongoIndexArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndex)(nil)).Elem()
}

func (o MongoIndexArrayOutput) ToMongoIndexArrayOutput() MongoIndexArrayOutput {
	return o
}

func (o MongoIndexArrayOutput) ToMongoIndexArrayOutputWithContext(ctx context.Context) MongoIndexArrayOutput {
	return o
}

func (o MongoIndexArrayOutput) Index(i pulumi.IntInput) MongoIndexOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MongoIndex {
		return vs[0].([]MongoIndex)[vs[1].(int)]
	}).(MongoIndexOutput)
}

type MongoIndexKeys struct {
	Keys []string `pulumi:"keys"`
}





type MongoIndexKeysInput interface {
	pulumi.Input

	ToMongoIndexKeysOutput() MongoIndexKeysOutput
	ToMongoIndexKeysOutputWithContext(context.Context) MongoIndexKeysOutput
}

type MongoIndexKeysArgs struct {
	Keys pulumi.StringArrayInput `pulumi:"keys"`
}

func (MongoIndexKeysArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeys)(nil)).Elem()
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysOutput() MongoIndexKeysOutput {
	return i.ToMongoIndexKeysOutputWithContext(context.Background())
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysOutputWithContext(ctx context.Context) MongoIndexKeysOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysOutput)
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return i.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (i MongoIndexKeysArgs) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysOutput).ToMongoIndexKeysPtrOutputWithContext(ctx)
}









type MongoIndexKeysPtrInput interface {
	pulumi.Input

	ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput
	ToMongoIndexKeysPtrOutputWithContext(context.Context) MongoIndexKeysPtrOutput
}

type mongoIndexKeysPtrType MongoIndexKeysArgs

func MongoIndexKeysPtr(v *MongoIndexKeysArgs) MongoIndexKeysPtrInput {
	return (*mongoIndexKeysPtrType)(v)
}

func (*mongoIndexKeysPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeys)(nil)).Elem()
}

func (i *mongoIndexKeysPtrType) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return i.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (i *mongoIndexKeysPtrType) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysPtrOutput)
}

type MongoIndexKeysOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeys)(nil)).Elem()
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysOutput() MongoIndexKeysOutput {
	return o
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysOutputWithContext(ctx context.Context) MongoIndexKeysOutput {
	return o
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return o.ToMongoIndexKeysPtrOutputWithContext(context.Background())
}

func (o MongoIndexKeysOutput) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexKeys) *MongoIndexKeys {
		return &v
	}).(MongoIndexKeysPtrOutput)
}

func (o MongoIndexKeysOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoIndexKeys) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

type MongoIndexKeysPtrOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeys)(nil)).Elem()
}

func (o MongoIndexKeysPtrOutput) ToMongoIndexKeysPtrOutput() MongoIndexKeysPtrOutput {
	return o
}

func (o MongoIndexKeysPtrOutput) ToMongoIndexKeysPtrOutputWithContext(ctx context.Context) MongoIndexKeysPtrOutput {
	return o
}

func (o MongoIndexKeysPtrOutput) Elem() MongoIndexKeysOutput {
	return o.ApplyT(func(v *MongoIndexKeys) MongoIndexKeys {
		if v != nil {
			return *v
		}
		var ret MongoIndexKeys
		return ret
	}).(MongoIndexKeysOutput)
}

func (o MongoIndexKeysPtrOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoIndexKeys) []string {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.StringArrayOutput)
}

type MongoIndexKeysResponse struct {
	Keys []string `pulumi:"keys"`
}





type MongoIndexKeysResponseInput interface {
	pulumi.Input

	ToMongoIndexKeysResponseOutput() MongoIndexKeysResponseOutput
	ToMongoIndexKeysResponseOutputWithContext(context.Context) MongoIndexKeysResponseOutput
}

type MongoIndexKeysResponseArgs struct {
	Keys pulumi.StringArrayInput `pulumi:"keys"`
}

func (MongoIndexKeysResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeysResponse)(nil)).Elem()
}

func (i MongoIndexKeysResponseArgs) ToMongoIndexKeysResponseOutput() MongoIndexKeysResponseOutput {
	return i.ToMongoIndexKeysResponseOutputWithContext(context.Background())
}

func (i MongoIndexKeysResponseArgs) ToMongoIndexKeysResponseOutputWithContext(ctx context.Context) MongoIndexKeysResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysResponseOutput)
}

func (i MongoIndexKeysResponseArgs) ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput {
	return i.ToMongoIndexKeysResponsePtrOutputWithContext(context.Background())
}

func (i MongoIndexKeysResponseArgs) ToMongoIndexKeysResponsePtrOutputWithContext(ctx context.Context) MongoIndexKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysResponseOutput).ToMongoIndexKeysResponsePtrOutputWithContext(ctx)
}









type MongoIndexKeysResponsePtrInput interface {
	pulumi.Input

	ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput
	ToMongoIndexKeysResponsePtrOutputWithContext(context.Context) MongoIndexKeysResponsePtrOutput
}

type mongoIndexKeysResponsePtrType MongoIndexKeysResponseArgs

func MongoIndexKeysResponsePtr(v *MongoIndexKeysResponseArgs) MongoIndexKeysResponsePtrInput {
	return (*mongoIndexKeysResponsePtrType)(v)
}

func (*mongoIndexKeysResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeysResponse)(nil)).Elem()
}

func (i *mongoIndexKeysResponsePtrType) ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput {
	return i.ToMongoIndexKeysResponsePtrOutputWithContext(context.Background())
}

func (i *mongoIndexKeysResponsePtrType) ToMongoIndexKeysResponsePtrOutputWithContext(ctx context.Context) MongoIndexKeysResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexKeysResponsePtrOutput)
}

type MongoIndexKeysResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexKeysResponse)(nil)).Elem()
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponseOutput() MongoIndexKeysResponseOutput {
	return o
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponseOutputWithContext(ctx context.Context) MongoIndexKeysResponseOutput {
	return o
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput {
	return o.ToMongoIndexKeysResponsePtrOutputWithContext(context.Background())
}

func (o MongoIndexKeysResponseOutput) ToMongoIndexKeysResponsePtrOutputWithContext(ctx context.Context) MongoIndexKeysResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexKeysResponse) *MongoIndexKeysResponse {
		return &v
	}).(MongoIndexKeysResponsePtrOutput)
}

func (o MongoIndexKeysResponseOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v MongoIndexKeysResponse) []string { return v.Keys }).(pulumi.StringArrayOutput)
}

type MongoIndexKeysResponsePtrOutput struct{ *pulumi.OutputState }

func (MongoIndexKeysResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexKeysResponse)(nil)).Elem()
}

func (o MongoIndexKeysResponsePtrOutput) ToMongoIndexKeysResponsePtrOutput() MongoIndexKeysResponsePtrOutput {
	return o
}

func (o MongoIndexKeysResponsePtrOutput) ToMongoIndexKeysResponsePtrOutputWithContext(ctx context.Context) MongoIndexKeysResponsePtrOutput {
	return o
}

func (o MongoIndexKeysResponsePtrOutput) Elem() MongoIndexKeysResponseOutput {
	return o.ApplyT(func(v *MongoIndexKeysResponse) MongoIndexKeysResponse {
		if v != nil {
			return *v
		}
		var ret MongoIndexKeysResponse
		return ret
	}).(MongoIndexKeysResponseOutput)
}

func (o MongoIndexKeysResponsePtrOutput) Keys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *MongoIndexKeysResponse) []string {
		if v == nil {
			return nil
		}
		return v.Keys
	}).(pulumi.StringArrayOutput)
}

type MongoIndexOptions struct {
	ExpireAfterSeconds *int  `pulumi:"expireAfterSeconds"`
	Unique             *bool `pulumi:"unique"`
}





type MongoIndexOptionsInput interface {
	pulumi.Input

	ToMongoIndexOptionsOutput() MongoIndexOptionsOutput
	ToMongoIndexOptionsOutputWithContext(context.Context) MongoIndexOptionsOutput
}

type MongoIndexOptionsArgs struct {
	ExpireAfterSeconds pulumi.IntPtrInput  `pulumi:"expireAfterSeconds"`
	Unique             pulumi.BoolPtrInput `pulumi:"unique"`
}

func (MongoIndexOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptions)(nil)).Elem()
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsOutput() MongoIndexOptionsOutput {
	return i.ToMongoIndexOptionsOutputWithContext(context.Background())
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsOutputWithContext(ctx context.Context) MongoIndexOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsOutput)
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return i.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (i MongoIndexOptionsArgs) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsOutput).ToMongoIndexOptionsPtrOutputWithContext(ctx)
}









type MongoIndexOptionsPtrInput interface {
	pulumi.Input

	ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput
	ToMongoIndexOptionsPtrOutputWithContext(context.Context) MongoIndexOptionsPtrOutput
}

type mongoIndexOptionsPtrType MongoIndexOptionsArgs

func MongoIndexOptionsPtr(v *MongoIndexOptionsArgs) MongoIndexOptionsPtrInput {
	return (*mongoIndexOptionsPtrType)(v)
}

func (*mongoIndexOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptions)(nil)).Elem()
}

func (i *mongoIndexOptionsPtrType) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return i.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (i *mongoIndexOptionsPtrType) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsPtrOutput)
}

type MongoIndexOptionsOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptions)(nil)).Elem()
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsOutput() MongoIndexOptionsOutput {
	return o
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsOutputWithContext(ctx context.Context) MongoIndexOptionsOutput {
	return o
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return o.ToMongoIndexOptionsPtrOutputWithContext(context.Background())
}

func (o MongoIndexOptionsOutput) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexOptions) *MongoIndexOptions {
		return &v
	}).(MongoIndexOptionsPtrOutput)
}

func (o MongoIndexOptionsOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoIndexOptions) *int { return v.ExpireAfterSeconds }).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MongoIndexOptions) *bool { return v.Unique }).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsPtrOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptions)(nil)).Elem()
}

func (o MongoIndexOptionsPtrOutput) ToMongoIndexOptionsPtrOutput() MongoIndexOptionsPtrOutput {
	return o
}

func (o MongoIndexOptionsPtrOutput) ToMongoIndexOptionsPtrOutputWithContext(ctx context.Context) MongoIndexOptionsPtrOutput {
	return o
}

func (o MongoIndexOptionsPtrOutput) Elem() MongoIndexOptionsOutput {
	return o.ApplyT(func(v *MongoIndexOptions) MongoIndexOptions {
		if v != nil {
			return *v
		}
		var ret MongoIndexOptions
		return ret
	}).(MongoIndexOptionsOutput)
}

func (o MongoIndexOptionsPtrOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptions) *int {
		if v == nil {
			return nil
		}
		return v.ExpireAfterSeconds
	}).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsPtrOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptions) *bool {
		if v == nil {
			return nil
		}
		return v.Unique
	}).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsResponse struct {
	ExpireAfterSeconds *int  `pulumi:"expireAfterSeconds"`
	Unique             *bool `pulumi:"unique"`
}





type MongoIndexOptionsResponseInput interface {
	pulumi.Input

	ToMongoIndexOptionsResponseOutput() MongoIndexOptionsResponseOutput
	ToMongoIndexOptionsResponseOutputWithContext(context.Context) MongoIndexOptionsResponseOutput
}

type MongoIndexOptionsResponseArgs struct {
	ExpireAfterSeconds pulumi.IntPtrInput  `pulumi:"expireAfterSeconds"`
	Unique             pulumi.BoolPtrInput `pulumi:"unique"`
}

func (MongoIndexOptionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptionsResponse)(nil)).Elem()
}

func (i MongoIndexOptionsResponseArgs) ToMongoIndexOptionsResponseOutput() MongoIndexOptionsResponseOutput {
	return i.ToMongoIndexOptionsResponseOutputWithContext(context.Background())
}

func (i MongoIndexOptionsResponseArgs) ToMongoIndexOptionsResponseOutputWithContext(ctx context.Context) MongoIndexOptionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsResponseOutput)
}

func (i MongoIndexOptionsResponseArgs) ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput {
	return i.ToMongoIndexOptionsResponsePtrOutputWithContext(context.Background())
}

func (i MongoIndexOptionsResponseArgs) ToMongoIndexOptionsResponsePtrOutputWithContext(ctx context.Context) MongoIndexOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsResponseOutput).ToMongoIndexOptionsResponsePtrOutputWithContext(ctx)
}









type MongoIndexOptionsResponsePtrInput interface {
	pulumi.Input

	ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput
	ToMongoIndexOptionsResponsePtrOutputWithContext(context.Context) MongoIndexOptionsResponsePtrOutput
}

type mongoIndexOptionsResponsePtrType MongoIndexOptionsResponseArgs

func MongoIndexOptionsResponsePtr(v *MongoIndexOptionsResponseArgs) MongoIndexOptionsResponsePtrInput {
	return (*mongoIndexOptionsResponsePtrType)(v)
}

func (*mongoIndexOptionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptionsResponse)(nil)).Elem()
}

func (i *mongoIndexOptionsResponsePtrType) ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput {
	return i.ToMongoIndexOptionsResponsePtrOutputWithContext(context.Background())
}

func (i *mongoIndexOptionsResponsePtrType) ToMongoIndexOptionsResponsePtrOutputWithContext(ctx context.Context) MongoIndexOptionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexOptionsResponsePtrOutput)
}

type MongoIndexOptionsResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexOptionsResponse)(nil)).Elem()
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponseOutput() MongoIndexOptionsResponseOutput {
	return o
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponseOutputWithContext(ctx context.Context) MongoIndexOptionsResponseOutput {
	return o
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput {
	return o.ToMongoIndexOptionsResponsePtrOutputWithContext(context.Background())
}

func (o MongoIndexOptionsResponseOutput) ToMongoIndexOptionsResponsePtrOutputWithContext(ctx context.Context) MongoIndexOptionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MongoIndexOptionsResponse) *MongoIndexOptionsResponse {
		return &v
	}).(MongoIndexOptionsResponsePtrOutput)
}

func (o MongoIndexOptionsResponseOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v MongoIndexOptionsResponse) *int { return v.ExpireAfterSeconds }).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsResponseOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MongoIndexOptionsResponse) *bool { return v.Unique }).(pulumi.BoolPtrOutput)
}

type MongoIndexOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (MongoIndexOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MongoIndexOptionsResponse)(nil)).Elem()
}

func (o MongoIndexOptionsResponsePtrOutput) ToMongoIndexOptionsResponsePtrOutput() MongoIndexOptionsResponsePtrOutput {
	return o
}

func (o MongoIndexOptionsResponsePtrOutput) ToMongoIndexOptionsResponsePtrOutputWithContext(ctx context.Context) MongoIndexOptionsResponsePtrOutput {
	return o
}

func (o MongoIndexOptionsResponsePtrOutput) Elem() MongoIndexOptionsResponseOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) MongoIndexOptionsResponse {
		if v != nil {
			return *v
		}
		var ret MongoIndexOptionsResponse
		return ret
	}).(MongoIndexOptionsResponseOutput)
}

func (o MongoIndexOptionsResponsePtrOutput) ExpireAfterSeconds() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) *int {
		if v == nil {
			return nil
		}
		return v.ExpireAfterSeconds
	}).(pulumi.IntPtrOutput)
}

func (o MongoIndexOptionsResponsePtrOutput) Unique() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MongoIndexOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Unique
	}).(pulumi.BoolPtrOutput)
}

type MongoIndexResponse struct {
	Key     *MongoIndexKeysResponse    `pulumi:"key"`
	Options *MongoIndexOptionsResponse `pulumi:"options"`
}





type MongoIndexResponseInput interface {
	pulumi.Input

	ToMongoIndexResponseOutput() MongoIndexResponseOutput
	ToMongoIndexResponseOutputWithContext(context.Context) MongoIndexResponseOutput
}

type MongoIndexResponseArgs struct {
	Key     MongoIndexKeysResponsePtrInput    `pulumi:"key"`
	Options MongoIndexOptionsResponsePtrInput `pulumi:"options"`
}

func (MongoIndexResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexResponse)(nil)).Elem()
}

func (i MongoIndexResponseArgs) ToMongoIndexResponseOutput() MongoIndexResponseOutput {
	return i.ToMongoIndexResponseOutputWithContext(context.Background())
}

func (i MongoIndexResponseArgs) ToMongoIndexResponseOutputWithContext(ctx context.Context) MongoIndexResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexResponseOutput)
}





type MongoIndexResponseArrayInput interface {
	pulumi.Input

	ToMongoIndexResponseArrayOutput() MongoIndexResponseArrayOutput
	ToMongoIndexResponseArrayOutputWithContext(context.Context) MongoIndexResponseArrayOutput
}

type MongoIndexResponseArray []MongoIndexResponseInput

func (MongoIndexResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndexResponse)(nil)).Elem()
}

func (i MongoIndexResponseArray) ToMongoIndexResponseArrayOutput() MongoIndexResponseArrayOutput {
	return i.ToMongoIndexResponseArrayOutputWithContext(context.Background())
}

func (i MongoIndexResponseArray) ToMongoIndexResponseArrayOutputWithContext(ctx context.Context) MongoIndexResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MongoIndexResponseArrayOutput)
}

type MongoIndexResponseOutput struct{ *pulumi.OutputState }

func (MongoIndexResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MongoIndexResponse)(nil)).Elem()
}

func (o MongoIndexResponseOutput) ToMongoIndexResponseOutput() MongoIndexResponseOutput {
	return o
}

func (o MongoIndexResponseOutput) ToMongoIndexResponseOutputWithContext(ctx context.Context) MongoIndexResponseOutput {
	return o
}

func (o MongoIndexResponseOutput) Key() MongoIndexKeysResponsePtrOutput {
	return o.ApplyT(func(v MongoIndexResponse) *MongoIndexKeysResponse { return v.Key }).(MongoIndexKeysResponsePtrOutput)
}

func (o MongoIndexResponseOutput) Options() MongoIndexOptionsResponsePtrOutput {
	return o.ApplyT(func(v MongoIndexResponse) *MongoIndexOptionsResponse { return v.Options }).(MongoIndexOptionsResponsePtrOutput)
}

type MongoIndexResponseArrayOutput struct{ *pulumi.OutputState }

func (MongoIndexResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MongoIndexResponse)(nil)).Elem()
}

func (o MongoIndexResponseArrayOutput) ToMongoIndexResponseArrayOutput() MongoIndexResponseArrayOutput {
	return o
}

func (o MongoIndexResponseArrayOutput) ToMongoIndexResponseArrayOutputWithContext(ctx context.Context) MongoIndexResponseArrayOutput {
	return o
}

func (o MongoIndexResponseArrayOutput) Index(i pulumi.IntInput) MongoIndexResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MongoIndexResponse {
		return vs[0].([]MongoIndexResponse)[vs[1].(int)]
	}).(MongoIndexResponseOutput)
}

type PeriodicModeBackupPolicy struct {
	MigrationState         *BackupPolicyMigrationState `pulumi:"migrationState"`
	PeriodicModeProperties *PeriodicModeProperties     `pulumi:"periodicModeProperties"`
	Type                   string                      `pulumi:"type"`
}





type PeriodicModeBackupPolicyInput interface {
	pulumi.Input

	ToPeriodicModeBackupPolicyOutput() PeriodicModeBackupPolicyOutput
	ToPeriodicModeBackupPolicyOutputWithContext(context.Context) PeriodicModeBackupPolicyOutput
}

type PeriodicModeBackupPolicyArgs struct {
	MigrationState         BackupPolicyMigrationStatePtrInput `pulumi:"migrationState"`
	PeriodicModeProperties PeriodicModePropertiesPtrInput     `pulumi:"periodicModeProperties"`
	Type                   pulumi.StringInput                 `pulumi:"type"`
}

func (PeriodicModeBackupPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModeBackupPolicy)(nil)).Elem()
}

func (i PeriodicModeBackupPolicyArgs) ToPeriodicModeBackupPolicyOutput() PeriodicModeBackupPolicyOutput {
	return i.ToPeriodicModeBackupPolicyOutputWithContext(context.Background())
}

func (i PeriodicModeBackupPolicyArgs) ToPeriodicModeBackupPolicyOutputWithContext(ctx context.Context) PeriodicModeBackupPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModeBackupPolicyOutput)
}

type PeriodicModeBackupPolicyOutput struct{ *pulumi.OutputState }

func (PeriodicModeBackupPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModeBackupPolicy)(nil)).Elem()
}

func (o PeriodicModeBackupPolicyOutput) ToPeriodicModeBackupPolicyOutput() PeriodicModeBackupPolicyOutput {
	return o
}

func (o PeriodicModeBackupPolicyOutput) ToPeriodicModeBackupPolicyOutputWithContext(ctx context.Context) PeriodicModeBackupPolicyOutput {
	return o
}

func (o PeriodicModeBackupPolicyOutput) MigrationState() BackupPolicyMigrationStatePtrOutput {
	return o.ApplyT(func(v PeriodicModeBackupPolicy) *BackupPolicyMigrationState { return v.MigrationState }).(BackupPolicyMigrationStatePtrOutput)
}

func (o PeriodicModeBackupPolicyOutput) PeriodicModeProperties() PeriodicModePropertiesPtrOutput {
	return o.ApplyT(func(v PeriodicModeBackupPolicy) *PeriodicModeProperties { return v.PeriodicModeProperties }).(PeriodicModePropertiesPtrOutput)
}

func (o PeriodicModeBackupPolicyOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PeriodicModeBackupPolicy) string { return v.Type }).(pulumi.StringOutput)
}

type PeriodicModeBackupPolicyResponse struct {
	MigrationState         *BackupPolicyMigrationStateResponse `pulumi:"migrationState"`
	PeriodicModeProperties *PeriodicModePropertiesResponse     `pulumi:"periodicModeProperties"`
	Type                   string                              `pulumi:"type"`
}





type PeriodicModeBackupPolicyResponseInput interface {
	pulumi.Input

	ToPeriodicModeBackupPolicyResponseOutput() PeriodicModeBackupPolicyResponseOutput
	ToPeriodicModeBackupPolicyResponseOutputWithContext(context.Context) PeriodicModeBackupPolicyResponseOutput
}

type PeriodicModeBackupPolicyResponseArgs struct {
	MigrationState         BackupPolicyMigrationStateResponsePtrInput `pulumi:"migrationState"`
	PeriodicModeProperties PeriodicModePropertiesResponsePtrInput     `pulumi:"periodicModeProperties"`
	Type                   pulumi.StringInput                         `pulumi:"type"`
}

func (PeriodicModeBackupPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModeBackupPolicyResponse)(nil)).Elem()
}

func (i PeriodicModeBackupPolicyResponseArgs) ToPeriodicModeBackupPolicyResponseOutput() PeriodicModeBackupPolicyResponseOutput {
	return i.ToPeriodicModeBackupPolicyResponseOutputWithContext(context.Background())
}

func (i PeriodicModeBackupPolicyResponseArgs) ToPeriodicModeBackupPolicyResponseOutputWithContext(ctx context.Context) PeriodicModeBackupPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModeBackupPolicyResponseOutput)
}

type PeriodicModeBackupPolicyResponseOutput struct{ *pulumi.OutputState }

func (PeriodicModeBackupPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModeBackupPolicyResponse)(nil)).Elem()
}

func (o PeriodicModeBackupPolicyResponseOutput) ToPeriodicModeBackupPolicyResponseOutput() PeriodicModeBackupPolicyResponseOutput {
	return o
}

func (o PeriodicModeBackupPolicyResponseOutput) ToPeriodicModeBackupPolicyResponseOutputWithContext(ctx context.Context) PeriodicModeBackupPolicyResponseOutput {
	return o
}

func (o PeriodicModeBackupPolicyResponseOutput) MigrationState() BackupPolicyMigrationStateResponsePtrOutput {
	return o.ApplyT(func(v PeriodicModeBackupPolicyResponse) *BackupPolicyMigrationStateResponse { return v.MigrationState }).(BackupPolicyMigrationStateResponsePtrOutput)
}

func (o PeriodicModeBackupPolicyResponseOutput) PeriodicModeProperties() PeriodicModePropertiesResponsePtrOutput {
	return o.ApplyT(func(v PeriodicModeBackupPolicyResponse) *PeriodicModePropertiesResponse {
		return v.PeriodicModeProperties
	}).(PeriodicModePropertiesResponsePtrOutput)
}

func (o PeriodicModeBackupPolicyResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PeriodicModeBackupPolicyResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PeriodicModeProperties struct {
	BackupIntervalInMinutes        *int    `pulumi:"backupIntervalInMinutes"`
	BackupRetentionIntervalInHours *int    `pulumi:"backupRetentionIntervalInHours"`
	BackupStorageRedundancy        *string `pulumi:"backupStorageRedundancy"`
}





type PeriodicModePropertiesInput interface {
	pulumi.Input

	ToPeriodicModePropertiesOutput() PeriodicModePropertiesOutput
	ToPeriodicModePropertiesOutputWithContext(context.Context) PeriodicModePropertiesOutput
}

type PeriodicModePropertiesArgs struct {
	BackupIntervalInMinutes        pulumi.IntPtrInput    `pulumi:"backupIntervalInMinutes"`
	BackupRetentionIntervalInHours pulumi.IntPtrInput    `pulumi:"backupRetentionIntervalInHours"`
	BackupStorageRedundancy        pulumi.StringPtrInput `pulumi:"backupStorageRedundancy"`
}

func (PeriodicModePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModeProperties)(nil)).Elem()
}

func (i PeriodicModePropertiesArgs) ToPeriodicModePropertiesOutput() PeriodicModePropertiesOutput {
	return i.ToPeriodicModePropertiesOutputWithContext(context.Background())
}

func (i PeriodicModePropertiesArgs) ToPeriodicModePropertiesOutputWithContext(ctx context.Context) PeriodicModePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModePropertiesOutput)
}

func (i PeriodicModePropertiesArgs) ToPeriodicModePropertiesPtrOutput() PeriodicModePropertiesPtrOutput {
	return i.ToPeriodicModePropertiesPtrOutputWithContext(context.Background())
}

func (i PeriodicModePropertiesArgs) ToPeriodicModePropertiesPtrOutputWithContext(ctx context.Context) PeriodicModePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModePropertiesOutput).ToPeriodicModePropertiesPtrOutputWithContext(ctx)
}









type PeriodicModePropertiesPtrInput interface {
	pulumi.Input

	ToPeriodicModePropertiesPtrOutput() PeriodicModePropertiesPtrOutput
	ToPeriodicModePropertiesPtrOutputWithContext(context.Context) PeriodicModePropertiesPtrOutput
}

type periodicModePropertiesPtrType PeriodicModePropertiesArgs

func PeriodicModePropertiesPtr(v *PeriodicModePropertiesArgs) PeriodicModePropertiesPtrInput {
	return (*periodicModePropertiesPtrType)(v)
}

func (*periodicModePropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicModeProperties)(nil)).Elem()
}

func (i *periodicModePropertiesPtrType) ToPeriodicModePropertiesPtrOutput() PeriodicModePropertiesPtrOutput {
	return i.ToPeriodicModePropertiesPtrOutputWithContext(context.Background())
}

func (i *periodicModePropertiesPtrType) ToPeriodicModePropertiesPtrOutputWithContext(ctx context.Context) PeriodicModePropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModePropertiesPtrOutput)
}

type PeriodicModePropertiesOutput struct{ *pulumi.OutputState }

func (PeriodicModePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModeProperties)(nil)).Elem()
}

func (o PeriodicModePropertiesOutput) ToPeriodicModePropertiesOutput() PeriodicModePropertiesOutput {
	return o
}

func (o PeriodicModePropertiesOutput) ToPeriodicModePropertiesOutputWithContext(ctx context.Context) PeriodicModePropertiesOutput {
	return o
}

func (o PeriodicModePropertiesOutput) ToPeriodicModePropertiesPtrOutput() PeriodicModePropertiesPtrOutput {
	return o.ToPeriodicModePropertiesPtrOutputWithContext(context.Background())
}

func (o PeriodicModePropertiesOutput) ToPeriodicModePropertiesPtrOutputWithContext(ctx context.Context) PeriodicModePropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeriodicModeProperties) *PeriodicModeProperties {
		return &v
	}).(PeriodicModePropertiesPtrOutput)
}

func (o PeriodicModePropertiesOutput) BackupIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PeriodicModeProperties) *int { return v.BackupIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesOutput) BackupRetentionIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PeriodicModeProperties) *int { return v.BackupRetentionIntervalInHours }).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesOutput) BackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeriodicModeProperties) *string { return v.BackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

type PeriodicModePropertiesPtrOutput struct{ *pulumi.OutputState }

func (PeriodicModePropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicModeProperties)(nil)).Elem()
}

func (o PeriodicModePropertiesPtrOutput) ToPeriodicModePropertiesPtrOutput() PeriodicModePropertiesPtrOutput {
	return o
}

func (o PeriodicModePropertiesPtrOutput) ToPeriodicModePropertiesPtrOutputWithContext(ctx context.Context) PeriodicModePropertiesPtrOutput {
	return o
}

func (o PeriodicModePropertiesPtrOutput) Elem() PeriodicModePropertiesOutput {
	return o.ApplyT(func(v *PeriodicModeProperties) PeriodicModeProperties {
		if v != nil {
			return *v
		}
		var ret PeriodicModeProperties
		return ret
	}).(PeriodicModePropertiesOutput)
}

func (o PeriodicModePropertiesPtrOutput) BackupIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeriodicModeProperties) *int {
		if v == nil {
			return nil
		}
		return v.BackupIntervalInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesPtrOutput) BackupRetentionIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeriodicModeProperties) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionIntervalInHours
	}).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesPtrOutput) BackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicModeProperties) *string {
		if v == nil {
			return nil
		}
		return v.BackupStorageRedundancy
	}).(pulumi.StringPtrOutput)
}

type PeriodicModePropertiesResponse struct {
	BackupIntervalInMinutes        *int    `pulumi:"backupIntervalInMinutes"`
	BackupRetentionIntervalInHours *int    `pulumi:"backupRetentionIntervalInHours"`
	BackupStorageRedundancy        *string `pulumi:"backupStorageRedundancy"`
}





type PeriodicModePropertiesResponseInput interface {
	pulumi.Input

	ToPeriodicModePropertiesResponseOutput() PeriodicModePropertiesResponseOutput
	ToPeriodicModePropertiesResponseOutputWithContext(context.Context) PeriodicModePropertiesResponseOutput
}

type PeriodicModePropertiesResponseArgs struct {
	BackupIntervalInMinutes        pulumi.IntPtrInput    `pulumi:"backupIntervalInMinutes"`
	BackupRetentionIntervalInHours pulumi.IntPtrInput    `pulumi:"backupRetentionIntervalInHours"`
	BackupStorageRedundancy        pulumi.StringPtrInput `pulumi:"backupStorageRedundancy"`
}

func (PeriodicModePropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModePropertiesResponse)(nil)).Elem()
}

func (i PeriodicModePropertiesResponseArgs) ToPeriodicModePropertiesResponseOutput() PeriodicModePropertiesResponseOutput {
	return i.ToPeriodicModePropertiesResponseOutputWithContext(context.Background())
}

func (i PeriodicModePropertiesResponseArgs) ToPeriodicModePropertiesResponseOutputWithContext(ctx context.Context) PeriodicModePropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModePropertiesResponseOutput)
}

func (i PeriodicModePropertiesResponseArgs) ToPeriodicModePropertiesResponsePtrOutput() PeriodicModePropertiesResponsePtrOutput {
	return i.ToPeriodicModePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i PeriodicModePropertiesResponseArgs) ToPeriodicModePropertiesResponsePtrOutputWithContext(ctx context.Context) PeriodicModePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModePropertiesResponseOutput).ToPeriodicModePropertiesResponsePtrOutputWithContext(ctx)
}









type PeriodicModePropertiesResponsePtrInput interface {
	pulumi.Input

	ToPeriodicModePropertiesResponsePtrOutput() PeriodicModePropertiesResponsePtrOutput
	ToPeriodicModePropertiesResponsePtrOutputWithContext(context.Context) PeriodicModePropertiesResponsePtrOutput
}

type periodicModePropertiesResponsePtrType PeriodicModePropertiesResponseArgs

func PeriodicModePropertiesResponsePtr(v *PeriodicModePropertiesResponseArgs) PeriodicModePropertiesResponsePtrInput {
	return (*periodicModePropertiesResponsePtrType)(v)
}

func (*periodicModePropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicModePropertiesResponse)(nil)).Elem()
}

func (i *periodicModePropertiesResponsePtrType) ToPeriodicModePropertiesResponsePtrOutput() PeriodicModePropertiesResponsePtrOutput {
	return i.ToPeriodicModePropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *periodicModePropertiesResponsePtrType) ToPeriodicModePropertiesResponsePtrOutputWithContext(ctx context.Context) PeriodicModePropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PeriodicModePropertiesResponsePtrOutput)
}

type PeriodicModePropertiesResponseOutput struct{ *pulumi.OutputState }

func (PeriodicModePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PeriodicModePropertiesResponse)(nil)).Elem()
}

func (o PeriodicModePropertiesResponseOutput) ToPeriodicModePropertiesResponseOutput() PeriodicModePropertiesResponseOutput {
	return o
}

func (o PeriodicModePropertiesResponseOutput) ToPeriodicModePropertiesResponseOutputWithContext(ctx context.Context) PeriodicModePropertiesResponseOutput {
	return o
}

func (o PeriodicModePropertiesResponseOutput) ToPeriodicModePropertiesResponsePtrOutput() PeriodicModePropertiesResponsePtrOutput {
	return o.ToPeriodicModePropertiesResponsePtrOutputWithContext(context.Background())
}

func (o PeriodicModePropertiesResponseOutput) ToPeriodicModePropertiesResponsePtrOutputWithContext(ctx context.Context) PeriodicModePropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PeriodicModePropertiesResponse) *PeriodicModePropertiesResponse {
		return &v
	}).(PeriodicModePropertiesResponsePtrOutput)
}

func (o PeriodicModePropertiesResponseOutput) BackupIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PeriodicModePropertiesResponse) *int { return v.BackupIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesResponseOutput) BackupRetentionIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v PeriodicModePropertiesResponse) *int { return v.BackupRetentionIntervalInHours }).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesResponseOutput) BackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PeriodicModePropertiesResponse) *string { return v.BackupStorageRedundancy }).(pulumi.StringPtrOutput)
}

type PeriodicModePropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (PeriodicModePropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PeriodicModePropertiesResponse)(nil)).Elem()
}

func (o PeriodicModePropertiesResponsePtrOutput) ToPeriodicModePropertiesResponsePtrOutput() PeriodicModePropertiesResponsePtrOutput {
	return o
}

func (o PeriodicModePropertiesResponsePtrOutput) ToPeriodicModePropertiesResponsePtrOutputWithContext(ctx context.Context) PeriodicModePropertiesResponsePtrOutput {
	return o
}

func (o PeriodicModePropertiesResponsePtrOutput) Elem() PeriodicModePropertiesResponseOutput {
	return o.ApplyT(func(v *PeriodicModePropertiesResponse) PeriodicModePropertiesResponse {
		if v != nil {
			return *v
		}
		var ret PeriodicModePropertiesResponse
		return ret
	}).(PeriodicModePropertiesResponseOutput)
}

func (o PeriodicModePropertiesResponsePtrOutput) BackupIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeriodicModePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.BackupIntervalInMinutes
	}).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesResponsePtrOutput) BackupRetentionIntervalInHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *PeriodicModePropertiesResponse) *int {
		if v == nil {
			return nil
		}
		return v.BackupRetentionIntervalInHours
	}).(pulumi.IntPtrOutput)
}

func (o PeriodicModePropertiesResponsePtrOutput) BackupStorageRedundancy() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PeriodicModePropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.BackupStorageRedundancy
	}).(pulumi.StringPtrOutput)
}

type Permission struct {
	DataActions    []string `pulumi:"dataActions"`
	NotDataActions []string `pulumi:"notDataActions"`
}





type PermissionInput interface {
	pulumi.Input

	ToPermissionOutput() PermissionOutput
	ToPermissionOutputWithContext(context.Context) PermissionOutput
}

type PermissionArgs struct {
	DataActions    pulumi.StringArrayInput `pulumi:"dataActions"`
	NotDataActions pulumi.StringArrayInput `pulumi:"notDataActions"`
}

func (PermissionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (i PermissionArgs) ToPermissionOutput() PermissionOutput {
	return i.ToPermissionOutputWithContext(context.Background())
}

func (i PermissionArgs) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionOutput)
}





type PermissionArrayInput interface {
	pulumi.Input

	ToPermissionArrayOutput() PermissionArrayOutput
	ToPermissionArrayOutputWithContext(context.Context) PermissionArrayOutput
}

type PermissionArray []PermissionInput

func (PermissionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (i PermissionArray) ToPermissionArrayOutput() PermissionArrayOutput {
	return i.ToPermissionArrayOutputWithContext(context.Background())
}

func (i PermissionArray) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionArrayOutput)
}

type PermissionOutput struct{ *pulumi.OutputState }

func (PermissionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Permission)(nil)).Elem()
}

func (o PermissionOutput) ToPermissionOutput() PermissionOutput {
	return o
}

func (o PermissionOutput) ToPermissionOutputWithContext(ctx context.Context) PermissionOutput {
	return o
}

func (o PermissionOutput) DataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.DataActions }).(pulumi.StringArrayOutput)
}

func (o PermissionOutput) NotDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Permission) []string { return v.NotDataActions }).(pulumi.StringArrayOutput)
}

type PermissionArrayOutput struct{ *pulumi.OutputState }

func (PermissionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Permission)(nil)).Elem()
}

func (o PermissionArrayOutput) ToPermissionArrayOutput() PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) ToPermissionArrayOutputWithContext(ctx context.Context) PermissionArrayOutput {
	return o
}

func (o PermissionArrayOutput) Index(i pulumi.IntInput) PermissionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Permission {
		return vs[0].([]Permission)[vs[1].(int)]
	}).(PermissionOutput)
}

type PermissionResponse struct {
	DataActions    []string `pulumi:"dataActions"`
	NotDataActions []string `pulumi:"notDataActions"`
}





type PermissionResponseInput interface {
	pulumi.Input

	ToPermissionResponseOutput() PermissionResponseOutput
	ToPermissionResponseOutputWithContext(context.Context) PermissionResponseOutput
}

type PermissionResponseArgs struct {
	DataActions    pulumi.StringArrayInput `pulumi:"dataActions"`
	NotDataActions pulumi.StringArrayInput `pulumi:"notDataActions"`
}

func (PermissionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionResponse)(nil)).Elem()
}

func (i PermissionResponseArgs) ToPermissionResponseOutput() PermissionResponseOutput {
	return i.ToPermissionResponseOutputWithContext(context.Background())
}

func (i PermissionResponseArgs) ToPermissionResponseOutputWithContext(ctx context.Context) PermissionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionResponseOutput)
}





type PermissionResponseArrayInput interface {
	pulumi.Input

	ToPermissionResponseArrayOutput() PermissionResponseArrayOutput
	ToPermissionResponseArrayOutputWithContext(context.Context) PermissionResponseArrayOutput
}

type PermissionResponseArray []PermissionResponseInput

func (PermissionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionResponse)(nil)).Elem()
}

func (i PermissionResponseArray) ToPermissionResponseArrayOutput() PermissionResponseArrayOutput {
	return i.ToPermissionResponseArrayOutputWithContext(context.Background())
}

func (i PermissionResponseArray) ToPermissionResponseArrayOutputWithContext(ctx context.Context) PermissionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PermissionResponseArrayOutput)
}

type PermissionResponseOutput struct{ *pulumi.OutputState }

func (PermissionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseOutput) ToPermissionResponseOutput() PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) ToPermissionResponseOutputWithContext(ctx context.Context) PermissionResponseOutput {
	return o
}

func (o PermissionResponseOutput) DataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.DataActions }).(pulumi.StringArrayOutput)
}

func (o PermissionResponseOutput) NotDataActions() pulumi.StringArrayOutput {
	return o.ApplyT(func(v PermissionResponse) []string { return v.NotDataActions }).(pulumi.StringArrayOutput)
}

type PermissionResponseArrayOutput struct{ *pulumi.OutputState }

func (PermissionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PermissionResponse)(nil)).Elem()
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutput() PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) ToPermissionResponseArrayOutputWithContext(ctx context.Context) PermissionResponseArrayOutput {
	return o
}

func (o PermissionResponseArrayOutput) Index(i pulumi.IntInput) PermissionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PermissionResponse {
		return vs[0].([]PermissionResponse)[vs[1].(int)]
	}).(PermissionResponseOutput)
}

type PrivateEndpointConnectionResponse struct {
	GroupId                           *string                                            `pulumi:"groupId"`
	Id                                string                                             `pulumi:"id"`
	Name                              string                                             `pulumi:"name"`
	PrivateEndpoint                   *PrivateEndpointPropertyResponse                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionStatePropertyResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 *string                                            `pulumi:"provisioningState"`
	Type                              string                                             `pulumi:"type"`
}





type PrivateEndpointConnectionResponseInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput
	ToPrivateEndpointConnectionResponseOutputWithContext(context.Context) PrivateEndpointConnectionResponseOutput
}

type PrivateEndpointConnectionResponseArgs struct {
	GroupId                           pulumi.StringPtrInput                                     `pulumi:"groupId"`
	Id                                pulumi.StringInput                                        `pulumi:"id"`
	Name                              pulumi.StringInput                                        `pulumi:"name"`
	PrivateEndpoint                   PrivateEndpointPropertyResponsePtrInput                   `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState PrivateLinkServiceConnectionStatePropertyResponsePtrInput `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 pulumi.StringPtrInput                                     `pulumi:"provisioningState"`
	Type                              pulumi.StringInput                                        `pulumi:"type"`
}

func (PrivateEndpointConnectionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return i.ToPrivateEndpointConnectionResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArgs) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseOutput)
}





type PrivateEndpointConnectionResponseArrayInput interface {
	pulumi.Input

	ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput
	ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Context) PrivateEndpointConnectionResponseArrayOutput
}

type PrivateEndpointConnectionResponseArray []PrivateEndpointConnectionResponseInput

func (PrivateEndpointConnectionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return i.ToPrivateEndpointConnectionResponseArrayOutputWithContext(context.Background())
}

func (i PrivateEndpointConnectionResponseArray) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointConnectionResponseArrayOutput)
}

type PrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutput() PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) ToPrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseOutput {
	return o
}

func (o PrivateEndpointConnectionResponseOutput) GroupId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.GroupId }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateEndpoint() PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateEndpointPropertyResponse { return v.PrivateEndpoint }).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o PrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (PrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]PrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutput() PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) ToPrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) PrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o PrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) PrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) PrivateEndpointConnectionResponse {
		return vs[0].([]PrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(PrivateEndpointConnectionResponseOutput)
}

type PrivateEndpointProperty struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput
	ToPrivateEndpointPropertyOutputWithContext(context.Context) PrivateEndpointPropertyOutput
}

type PrivateEndpointPropertyArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return i.ToPrivateEndpointPropertyOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput)
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyArgs) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyOutput).ToPrivateEndpointPropertyPtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyPtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput
	ToPrivateEndpointPropertyPtrOutputWithContext(context.Context) PrivateEndpointPropertyPtrOutput
}

type privateEndpointPropertyPtrType PrivateEndpointPropertyArgs

func PrivateEndpointPropertyPtr(v *PrivateEndpointPropertyArgs) PrivateEndpointPropertyPtrInput {
	return (*privateEndpointPropertyPtrType)(v)
}

func (*privateEndpointPropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return i.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyPtrType) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyPtrOutput)
}

type PrivateEndpointPropertyOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutput() PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyOutputWithContext(ctx context.Context) PrivateEndpointPropertyOutput {
	return o
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o.ToPrivateEndpointPropertyPtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointProperty) *PrivateEndpointProperty {
		return &v
	}).(PrivateEndpointPropertyPtrOutput)
}

func (o PrivateEndpointPropertyOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointProperty) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointProperty)(nil)).Elem()
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutput() PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) ToPrivateEndpointPropertyPtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyPtrOutput {
	return o
}

func (o PrivateEndpointPropertyPtrOutput) Elem() PrivateEndpointPropertyOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) PrivateEndpointProperty {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointProperty
		return ret
	}).(PrivateEndpointPropertyOutput)
}

func (o PrivateEndpointPropertyPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointProperty) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponse struct {
	Id *string `pulumi:"id"`
}





type PrivateEndpointPropertyResponseInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput
	ToPrivateEndpointPropertyResponseOutputWithContext(context.Context) PrivateEndpointPropertyResponseOutput
}

type PrivateEndpointPropertyResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (PrivateEndpointPropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return i.ToPrivateEndpointPropertyResponseOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput)
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateEndpointPropertyResponseArgs) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponseOutput).ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx)
}









type PrivateEndpointPropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput
	ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Context) PrivateEndpointPropertyResponsePtrOutput
}

type privateEndpointPropertyResponsePtrType PrivateEndpointPropertyResponseArgs

func PrivateEndpointPropertyResponsePtr(v *PrivateEndpointPropertyResponseArgs) PrivateEndpointPropertyResponsePtrInput {
	return (*privateEndpointPropertyResponsePtrType)(v)
}

func (*privateEndpointPropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return i.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateEndpointPropertyResponsePtrType) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateEndpointPropertyResponsePtrOutput)
}

type PrivateEndpointPropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutput() PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponseOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponseOutput {
	return o
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o.ToPrivateEndpointPropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateEndpointPropertyResponseOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateEndpointPropertyResponse) *PrivateEndpointPropertyResponse {
		return &v
	}).(PrivateEndpointPropertyResponsePtrOutput)
}

func (o PrivateEndpointPropertyResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateEndpointPropertyResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type PrivateEndpointPropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointPropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointPropertyResponse)(nil)).Elem()
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutput() PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) ToPrivateEndpointPropertyResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointPropertyResponsePtrOutput {
	return o
}

func (o PrivateEndpointPropertyResponsePtrOutput) Elem() PrivateEndpointPropertyResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) PrivateEndpointPropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointPropertyResponse
		return ret
	}).(PrivateEndpointPropertyResponseOutput)
}

func (o PrivateEndpointPropertyResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointPropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateProperty struct {
	Description *string `pulumi:"description"`
	Status      *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput
	ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyOutput
}

type PrivateLinkServiceConnectionStatePropertyArgs struct {
	Description pulumi.StringPtrInput `pulumi:"description"`
	Status      pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyArgs) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyOutput).ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyPtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput
	ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput
}

type privateLinkServiceConnectionStatePropertyPtrType PrivateLinkServiceConnectionStatePropertyArgs

func PrivateLinkServiceConnectionStatePropertyPtr(v *PrivateLinkServiceConnectionStatePropertyArgs) PrivateLinkServiceConnectionStatePropertyPtrInput {
	return (*privateLinkServiceConnectionStatePropertyPtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyPtrType) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutput() PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStateProperty) *PrivateLinkServiceConnectionStateProperty {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateProperty) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyPtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStateProperty)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutput() PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) ToPrivateLinkServiceConnectionStatePropertyPtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyPtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) PrivateLinkServiceConnectionStateProperty {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStateProperty
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyPtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStateProperty) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponse struct {
	ActionsRequired string  `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStatePropertyResponseInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput
	ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput
}

type PrivateLinkServiceConnectionStatePropertyResponseArgs struct {
	ActionsRequired pulumi.StringInput    `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStatePropertyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStatePropertyResponseArgs) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponseOutput).ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx)
}









type PrivateLinkServiceConnectionStatePropertyResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
	ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput
}

type privateLinkServiceConnectionStatePropertyResponsePtrType PrivateLinkServiceConnectionStatePropertyResponseArgs

func PrivateLinkServiceConnectionStatePropertyResponsePtr(v *PrivateLinkServiceConnectionStatePropertyResponseArgs) PrivateLinkServiceConnectionStatePropertyResponsePtrInput {
	return (*privateLinkServiceConnectionStatePropertyResponsePtrType)(v)
}

func (*privateLinkServiceConnectionStatePropertyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return i.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkServiceConnectionStatePropertyResponsePtrType) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutput() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkServiceConnectionStatePropertyResponse) *PrivateLinkServiceConnectionStatePropertyResponse {
		return &v
	}).(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) ActionsRequired() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) string { return v.ActionsRequired }).(pulumi.StringOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStatePropertyResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStatePropertyResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkServiceConnectionStatePropertyResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutput() PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ToPrivateLinkServiceConnectionStatePropertyResponsePtrOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStatePropertyResponsePtrOutput {
	return o
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Elem() PrivateLinkServiceConnectionStatePropertyResponseOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) PrivateLinkServiceConnectionStatePropertyResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkServiceConnectionStatePropertyResponse
		return ret
	}).(PrivateLinkServiceConnectionStatePropertyResponseOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStatePropertyResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkServiceConnectionStatePropertyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type RestoreParameters struct {
	DatabasesToRestore    []DatabaseRestoreResource `pulumi:"databasesToRestore"`
	RestoreMode           *string                   `pulumi:"restoreMode"`
	RestoreSource         *string                   `pulumi:"restoreSource"`
	RestoreTimestampInUtc *string                   `pulumi:"restoreTimestampInUtc"`
}





type RestoreParametersInput interface {
	pulumi.Input

	ToRestoreParametersOutput() RestoreParametersOutput
	ToRestoreParametersOutputWithContext(context.Context) RestoreParametersOutput
}

type RestoreParametersArgs struct {
	DatabasesToRestore    DatabaseRestoreResourceArrayInput `pulumi:"databasesToRestore"`
	RestoreMode           pulumi.StringPtrInput             `pulumi:"restoreMode"`
	RestoreSource         pulumi.StringPtrInput             `pulumi:"restoreSource"`
	RestoreTimestampInUtc pulumi.StringPtrInput             `pulumi:"restoreTimestampInUtc"`
}

func (RestoreParametersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoreParameters)(nil)).Elem()
}

func (i RestoreParametersArgs) ToRestoreParametersOutput() RestoreParametersOutput {
	return i.ToRestoreParametersOutputWithContext(context.Background())
}

func (i RestoreParametersArgs) ToRestoreParametersOutputWithContext(ctx context.Context) RestoreParametersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreParametersOutput)
}

func (i RestoreParametersArgs) ToRestoreParametersPtrOutput() RestoreParametersPtrOutput {
	return i.ToRestoreParametersPtrOutputWithContext(context.Background())
}

func (i RestoreParametersArgs) ToRestoreParametersPtrOutputWithContext(ctx context.Context) RestoreParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreParametersOutput).ToRestoreParametersPtrOutputWithContext(ctx)
}









type RestoreParametersPtrInput interface {
	pulumi.Input

	ToRestoreParametersPtrOutput() RestoreParametersPtrOutput
	ToRestoreParametersPtrOutputWithContext(context.Context) RestoreParametersPtrOutput
}

type restoreParametersPtrType RestoreParametersArgs

func RestoreParametersPtr(v *RestoreParametersArgs) RestoreParametersPtrInput {
	return (*restoreParametersPtrType)(v)
}

func (*restoreParametersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoreParameters)(nil)).Elem()
}

func (i *restoreParametersPtrType) ToRestoreParametersPtrOutput() RestoreParametersPtrOutput {
	return i.ToRestoreParametersPtrOutputWithContext(context.Background())
}

func (i *restoreParametersPtrType) ToRestoreParametersPtrOutputWithContext(ctx context.Context) RestoreParametersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreParametersPtrOutput)
}

type RestoreParametersOutput struct{ *pulumi.OutputState }

func (RestoreParametersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoreParameters)(nil)).Elem()
}

func (o RestoreParametersOutput) ToRestoreParametersOutput() RestoreParametersOutput {
	return o
}

func (o RestoreParametersOutput) ToRestoreParametersOutputWithContext(ctx context.Context) RestoreParametersOutput {
	return o
}

func (o RestoreParametersOutput) ToRestoreParametersPtrOutput() RestoreParametersPtrOutput {
	return o.ToRestoreParametersPtrOutputWithContext(context.Background())
}

func (o RestoreParametersOutput) ToRestoreParametersPtrOutputWithContext(ctx context.Context) RestoreParametersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestoreParameters) *RestoreParameters {
		return &v
	}).(RestoreParametersPtrOutput)
}

func (o RestoreParametersOutput) DatabasesToRestore() DatabaseRestoreResourceArrayOutput {
	return o.ApplyT(func(v RestoreParameters) []DatabaseRestoreResource { return v.DatabasesToRestore }).(DatabaseRestoreResourceArrayOutput)
}

func (o RestoreParametersOutput) RestoreMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoreParameters) *string { return v.RestoreMode }).(pulumi.StringPtrOutput)
}

func (o RestoreParametersOutput) RestoreSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoreParameters) *string { return v.RestoreSource }).(pulumi.StringPtrOutput)
}

func (o RestoreParametersOutput) RestoreTimestampInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoreParameters) *string { return v.RestoreTimestampInUtc }).(pulumi.StringPtrOutput)
}

type RestoreParametersPtrOutput struct{ *pulumi.OutputState }

func (RestoreParametersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoreParameters)(nil)).Elem()
}

func (o RestoreParametersPtrOutput) ToRestoreParametersPtrOutput() RestoreParametersPtrOutput {
	return o
}

func (o RestoreParametersPtrOutput) ToRestoreParametersPtrOutputWithContext(ctx context.Context) RestoreParametersPtrOutput {
	return o
}

func (o RestoreParametersPtrOutput) Elem() RestoreParametersOutput {
	return o.ApplyT(func(v *RestoreParameters) RestoreParameters {
		if v != nil {
			return *v
		}
		var ret RestoreParameters
		return ret
	}).(RestoreParametersOutput)
}

func (o RestoreParametersPtrOutput) DatabasesToRestore() DatabaseRestoreResourceArrayOutput {
	return o.ApplyT(func(v *RestoreParameters) []DatabaseRestoreResource {
		if v == nil {
			return nil
		}
		return v.DatabasesToRestore
	}).(DatabaseRestoreResourceArrayOutput)
}

func (o RestoreParametersPtrOutput) RestoreMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoreParameters) *string {
		if v == nil {
			return nil
		}
		return v.RestoreMode
	}).(pulumi.StringPtrOutput)
}

func (o RestoreParametersPtrOutput) RestoreSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoreParameters) *string {
		if v == nil {
			return nil
		}
		return v.RestoreSource
	}).(pulumi.StringPtrOutput)
}

func (o RestoreParametersPtrOutput) RestoreTimestampInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoreParameters) *string {
		if v == nil {
			return nil
		}
		return v.RestoreTimestampInUtc
	}).(pulumi.StringPtrOutput)
}

type RestoreParametersResponse struct {
	DatabasesToRestore    []DatabaseRestoreResourceResponse `pulumi:"databasesToRestore"`
	RestoreMode           *string                           `pulumi:"restoreMode"`
	RestoreSource         *string                           `pulumi:"restoreSource"`
	RestoreTimestampInUtc *string                           `pulumi:"restoreTimestampInUtc"`
}





type RestoreParametersResponseInput interface {
	pulumi.Input

	ToRestoreParametersResponseOutput() RestoreParametersResponseOutput
	ToRestoreParametersResponseOutputWithContext(context.Context) RestoreParametersResponseOutput
}

type RestoreParametersResponseArgs struct {
	DatabasesToRestore    DatabaseRestoreResourceResponseArrayInput `pulumi:"databasesToRestore"`
	RestoreMode           pulumi.StringPtrInput                     `pulumi:"restoreMode"`
	RestoreSource         pulumi.StringPtrInput                     `pulumi:"restoreSource"`
	RestoreTimestampInUtc pulumi.StringPtrInput                     `pulumi:"restoreTimestampInUtc"`
}

func (RestoreParametersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoreParametersResponse)(nil)).Elem()
}

func (i RestoreParametersResponseArgs) ToRestoreParametersResponseOutput() RestoreParametersResponseOutput {
	return i.ToRestoreParametersResponseOutputWithContext(context.Background())
}

func (i RestoreParametersResponseArgs) ToRestoreParametersResponseOutputWithContext(ctx context.Context) RestoreParametersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreParametersResponseOutput)
}

func (i RestoreParametersResponseArgs) ToRestoreParametersResponsePtrOutput() RestoreParametersResponsePtrOutput {
	return i.ToRestoreParametersResponsePtrOutputWithContext(context.Background())
}

func (i RestoreParametersResponseArgs) ToRestoreParametersResponsePtrOutputWithContext(ctx context.Context) RestoreParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreParametersResponseOutput).ToRestoreParametersResponsePtrOutputWithContext(ctx)
}









type RestoreParametersResponsePtrInput interface {
	pulumi.Input

	ToRestoreParametersResponsePtrOutput() RestoreParametersResponsePtrOutput
	ToRestoreParametersResponsePtrOutputWithContext(context.Context) RestoreParametersResponsePtrOutput
}

type restoreParametersResponsePtrType RestoreParametersResponseArgs

func RestoreParametersResponsePtr(v *RestoreParametersResponseArgs) RestoreParametersResponsePtrInput {
	return (*restoreParametersResponsePtrType)(v)
}

func (*restoreParametersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoreParametersResponse)(nil)).Elem()
}

func (i *restoreParametersResponsePtrType) ToRestoreParametersResponsePtrOutput() RestoreParametersResponsePtrOutput {
	return i.ToRestoreParametersResponsePtrOutputWithContext(context.Background())
}

func (i *restoreParametersResponsePtrType) ToRestoreParametersResponsePtrOutputWithContext(ctx context.Context) RestoreParametersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RestoreParametersResponsePtrOutput)
}

type RestoreParametersResponseOutput struct{ *pulumi.OutputState }

func (RestoreParametersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RestoreParametersResponse)(nil)).Elem()
}

func (o RestoreParametersResponseOutput) ToRestoreParametersResponseOutput() RestoreParametersResponseOutput {
	return o
}

func (o RestoreParametersResponseOutput) ToRestoreParametersResponseOutputWithContext(ctx context.Context) RestoreParametersResponseOutput {
	return o
}

func (o RestoreParametersResponseOutput) ToRestoreParametersResponsePtrOutput() RestoreParametersResponsePtrOutput {
	return o.ToRestoreParametersResponsePtrOutputWithContext(context.Background())
}

func (o RestoreParametersResponseOutput) ToRestoreParametersResponsePtrOutputWithContext(ctx context.Context) RestoreParametersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RestoreParametersResponse) *RestoreParametersResponse {
		return &v
	}).(RestoreParametersResponsePtrOutput)
}

func (o RestoreParametersResponseOutput) DatabasesToRestore() DatabaseRestoreResourceResponseArrayOutput {
	return o.ApplyT(func(v RestoreParametersResponse) []DatabaseRestoreResourceResponse { return v.DatabasesToRestore }).(DatabaseRestoreResourceResponseArrayOutput)
}

func (o RestoreParametersResponseOutput) RestoreMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoreParametersResponse) *string { return v.RestoreMode }).(pulumi.StringPtrOutput)
}

func (o RestoreParametersResponseOutput) RestoreSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoreParametersResponse) *string { return v.RestoreSource }).(pulumi.StringPtrOutput)
}

func (o RestoreParametersResponseOutput) RestoreTimestampInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RestoreParametersResponse) *string { return v.RestoreTimestampInUtc }).(pulumi.StringPtrOutput)
}

type RestoreParametersResponsePtrOutput struct{ *pulumi.OutputState }

func (RestoreParametersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RestoreParametersResponse)(nil)).Elem()
}

func (o RestoreParametersResponsePtrOutput) ToRestoreParametersResponsePtrOutput() RestoreParametersResponsePtrOutput {
	return o
}

func (o RestoreParametersResponsePtrOutput) ToRestoreParametersResponsePtrOutputWithContext(ctx context.Context) RestoreParametersResponsePtrOutput {
	return o
}

func (o RestoreParametersResponsePtrOutput) Elem() RestoreParametersResponseOutput {
	return o.ApplyT(func(v *RestoreParametersResponse) RestoreParametersResponse {
		if v != nil {
			return *v
		}
		var ret RestoreParametersResponse
		return ret
	}).(RestoreParametersResponseOutput)
}

func (o RestoreParametersResponsePtrOutput) DatabasesToRestore() DatabaseRestoreResourceResponseArrayOutput {
	return o.ApplyT(func(v *RestoreParametersResponse) []DatabaseRestoreResourceResponse {
		if v == nil {
			return nil
		}
		return v.DatabasesToRestore
	}).(DatabaseRestoreResourceResponseArrayOutput)
}

func (o RestoreParametersResponsePtrOutput) RestoreMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoreParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.RestoreMode
	}).(pulumi.StringPtrOutput)
}

func (o RestoreParametersResponsePtrOutput) RestoreSource() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoreParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.RestoreSource
	}).(pulumi.StringPtrOutput)
}

func (o RestoreParametersResponsePtrOutput) RestoreTimestampInUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RestoreParametersResponse) *string {
		if v == nil {
			return nil
		}
		return v.RestoreTimestampInUtc
	}).(pulumi.StringPtrOutput)
}

type SeedNode struct {
	IpAddress *string `pulumi:"ipAddress"`
}





type SeedNodeInput interface {
	pulumi.Input

	ToSeedNodeOutput() SeedNodeOutput
	ToSeedNodeOutputWithContext(context.Context) SeedNodeOutput
}

type SeedNodeArgs struct {
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (SeedNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SeedNode)(nil)).Elem()
}

func (i SeedNodeArgs) ToSeedNodeOutput() SeedNodeOutput {
	return i.ToSeedNodeOutputWithContext(context.Background())
}

func (i SeedNodeArgs) ToSeedNodeOutputWithContext(ctx context.Context) SeedNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeOutput)
}

func (i SeedNodeArgs) ToSeedNodePtrOutput() SeedNodePtrOutput {
	return i.ToSeedNodePtrOutputWithContext(context.Background())
}

func (i SeedNodeArgs) ToSeedNodePtrOutputWithContext(ctx context.Context) SeedNodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeOutput).ToSeedNodePtrOutputWithContext(ctx)
}









type SeedNodePtrInput interface {
	pulumi.Input

	ToSeedNodePtrOutput() SeedNodePtrOutput
	ToSeedNodePtrOutputWithContext(context.Context) SeedNodePtrOutput
}

type seedNodePtrType SeedNodeArgs

func SeedNodePtr(v *SeedNodeArgs) SeedNodePtrInput {
	return (*seedNodePtrType)(v)
}

func (*seedNodePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SeedNode)(nil)).Elem()
}

func (i *seedNodePtrType) ToSeedNodePtrOutput() SeedNodePtrOutput {
	return i.ToSeedNodePtrOutputWithContext(context.Background())
}

func (i *seedNodePtrType) ToSeedNodePtrOutputWithContext(ctx context.Context) SeedNodePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodePtrOutput)
}





type SeedNodeArrayInput interface {
	pulumi.Input

	ToSeedNodeArrayOutput() SeedNodeArrayOutput
	ToSeedNodeArrayOutputWithContext(context.Context) SeedNodeArrayOutput
}

type SeedNodeArray []SeedNodeInput

func (SeedNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SeedNode)(nil)).Elem()
}

func (i SeedNodeArray) ToSeedNodeArrayOutput() SeedNodeArrayOutput {
	return i.ToSeedNodeArrayOutputWithContext(context.Background())
}

func (i SeedNodeArray) ToSeedNodeArrayOutputWithContext(ctx context.Context) SeedNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeArrayOutput)
}

type SeedNodeOutput struct{ *pulumi.OutputState }

func (SeedNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SeedNode)(nil)).Elem()
}

func (o SeedNodeOutput) ToSeedNodeOutput() SeedNodeOutput {
	return o
}

func (o SeedNodeOutput) ToSeedNodeOutputWithContext(ctx context.Context) SeedNodeOutput {
	return o
}

func (o SeedNodeOutput) ToSeedNodePtrOutput() SeedNodePtrOutput {
	return o.ToSeedNodePtrOutputWithContext(context.Background())
}

func (o SeedNodeOutput) ToSeedNodePtrOutputWithContext(ctx context.Context) SeedNodePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SeedNode) *SeedNode {
		return &v
	}).(SeedNodePtrOutput)
}

func (o SeedNodeOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SeedNode) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type SeedNodePtrOutput struct{ *pulumi.OutputState }

func (SeedNodePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SeedNode)(nil)).Elem()
}

func (o SeedNodePtrOutput) ToSeedNodePtrOutput() SeedNodePtrOutput {
	return o
}

func (o SeedNodePtrOutput) ToSeedNodePtrOutputWithContext(ctx context.Context) SeedNodePtrOutput {
	return o
}

func (o SeedNodePtrOutput) Elem() SeedNodeOutput {
	return o.ApplyT(func(v *SeedNode) SeedNode {
		if v != nil {
			return *v
		}
		var ret SeedNode
		return ret
	}).(SeedNodeOutput)
}

func (o SeedNodePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SeedNode) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

type SeedNodeArrayOutput struct{ *pulumi.OutputState }

func (SeedNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SeedNode)(nil)).Elem()
}

func (o SeedNodeArrayOutput) ToSeedNodeArrayOutput() SeedNodeArrayOutput {
	return o
}

func (o SeedNodeArrayOutput) ToSeedNodeArrayOutputWithContext(ctx context.Context) SeedNodeArrayOutput {
	return o
}

func (o SeedNodeArrayOutput) Index(i pulumi.IntInput) SeedNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SeedNode {
		return vs[0].([]SeedNode)[vs[1].(int)]
	}).(SeedNodeOutput)
}

type SeedNodeResponse struct {
	IpAddress *string `pulumi:"ipAddress"`
}





type SeedNodeResponseInput interface {
	pulumi.Input

	ToSeedNodeResponseOutput() SeedNodeResponseOutput
	ToSeedNodeResponseOutputWithContext(context.Context) SeedNodeResponseOutput
}

type SeedNodeResponseArgs struct {
	IpAddress pulumi.StringPtrInput `pulumi:"ipAddress"`
}

func (SeedNodeResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SeedNodeResponse)(nil)).Elem()
}

func (i SeedNodeResponseArgs) ToSeedNodeResponseOutput() SeedNodeResponseOutput {
	return i.ToSeedNodeResponseOutputWithContext(context.Background())
}

func (i SeedNodeResponseArgs) ToSeedNodeResponseOutputWithContext(ctx context.Context) SeedNodeResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeResponseOutput)
}

func (i SeedNodeResponseArgs) ToSeedNodeResponsePtrOutput() SeedNodeResponsePtrOutput {
	return i.ToSeedNodeResponsePtrOutputWithContext(context.Background())
}

func (i SeedNodeResponseArgs) ToSeedNodeResponsePtrOutputWithContext(ctx context.Context) SeedNodeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeResponseOutput).ToSeedNodeResponsePtrOutputWithContext(ctx)
}









type SeedNodeResponsePtrInput interface {
	pulumi.Input

	ToSeedNodeResponsePtrOutput() SeedNodeResponsePtrOutput
	ToSeedNodeResponsePtrOutputWithContext(context.Context) SeedNodeResponsePtrOutput
}

type seedNodeResponsePtrType SeedNodeResponseArgs

func SeedNodeResponsePtr(v *SeedNodeResponseArgs) SeedNodeResponsePtrInput {
	return (*seedNodeResponsePtrType)(v)
}

func (*seedNodeResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SeedNodeResponse)(nil)).Elem()
}

func (i *seedNodeResponsePtrType) ToSeedNodeResponsePtrOutput() SeedNodeResponsePtrOutput {
	return i.ToSeedNodeResponsePtrOutputWithContext(context.Background())
}

func (i *seedNodeResponsePtrType) ToSeedNodeResponsePtrOutputWithContext(ctx context.Context) SeedNodeResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeResponsePtrOutput)
}





type SeedNodeResponseArrayInput interface {
	pulumi.Input

	ToSeedNodeResponseArrayOutput() SeedNodeResponseArrayOutput
	ToSeedNodeResponseArrayOutputWithContext(context.Context) SeedNodeResponseArrayOutput
}

type SeedNodeResponseArray []SeedNodeResponseInput

func (SeedNodeResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SeedNodeResponse)(nil)).Elem()
}

func (i SeedNodeResponseArray) ToSeedNodeResponseArrayOutput() SeedNodeResponseArrayOutput {
	return i.ToSeedNodeResponseArrayOutputWithContext(context.Background())
}

func (i SeedNodeResponseArray) ToSeedNodeResponseArrayOutputWithContext(ctx context.Context) SeedNodeResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SeedNodeResponseArrayOutput)
}

type SeedNodeResponseOutput struct{ *pulumi.OutputState }

func (SeedNodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SeedNodeResponse)(nil)).Elem()
}

func (o SeedNodeResponseOutput) ToSeedNodeResponseOutput() SeedNodeResponseOutput {
	return o
}

func (o SeedNodeResponseOutput) ToSeedNodeResponseOutputWithContext(ctx context.Context) SeedNodeResponseOutput {
	return o
}

func (o SeedNodeResponseOutput) ToSeedNodeResponsePtrOutput() SeedNodeResponsePtrOutput {
	return o.ToSeedNodeResponsePtrOutputWithContext(context.Background())
}

func (o SeedNodeResponseOutput) ToSeedNodeResponsePtrOutputWithContext(ctx context.Context) SeedNodeResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SeedNodeResponse) *SeedNodeResponse {
		return &v
	}).(SeedNodeResponsePtrOutput)
}

func (o SeedNodeResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SeedNodeResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

type SeedNodeResponsePtrOutput struct{ *pulumi.OutputState }

func (SeedNodeResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SeedNodeResponse)(nil)).Elem()
}

func (o SeedNodeResponsePtrOutput) ToSeedNodeResponsePtrOutput() SeedNodeResponsePtrOutput {
	return o
}

func (o SeedNodeResponsePtrOutput) ToSeedNodeResponsePtrOutputWithContext(ctx context.Context) SeedNodeResponsePtrOutput {
	return o
}

func (o SeedNodeResponsePtrOutput) Elem() SeedNodeResponseOutput {
	return o.ApplyT(func(v *SeedNodeResponse) SeedNodeResponse {
		if v != nil {
			return *v
		}
		var ret SeedNodeResponse
		return ret
	}).(SeedNodeResponseOutput)
}

func (o SeedNodeResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SeedNodeResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

type SeedNodeResponseArrayOutput struct{ *pulumi.OutputState }

func (SeedNodeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SeedNodeResponse)(nil)).Elem()
}

func (o SeedNodeResponseArrayOutput) ToSeedNodeResponseArrayOutput() SeedNodeResponseArrayOutput {
	return o
}

func (o SeedNodeResponseArrayOutput) ToSeedNodeResponseArrayOutputWithContext(ctx context.Context) SeedNodeResponseArrayOutput {
	return o
}

func (o SeedNodeResponseArrayOutput) Index(i pulumi.IntInput) SeedNodeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SeedNodeResponse {
		return vs[0].([]SeedNodeResponse)[vs[1].(int)]
	}).(SeedNodeResponseOutput)
}

type SpatialSpec struct {
	Path  *string  `pulumi:"path"`
	Types []string `pulumi:"types"`
}





type SpatialSpecInput interface {
	pulumi.Input

	ToSpatialSpecOutput() SpatialSpecOutput
	ToSpatialSpecOutputWithContext(context.Context) SpatialSpecOutput
}

type SpatialSpecArgs struct {
	Path  pulumi.StringPtrInput   `pulumi:"path"`
	Types pulumi.StringArrayInput `pulumi:"types"`
}

func (SpatialSpecArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpec)(nil)).Elem()
}

func (i SpatialSpecArgs) ToSpatialSpecOutput() SpatialSpecOutput {
	return i.ToSpatialSpecOutputWithContext(context.Background())
}

func (i SpatialSpecArgs) ToSpatialSpecOutputWithContext(ctx context.Context) SpatialSpecOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialSpecOutput)
}





type SpatialSpecArrayInput interface {
	pulumi.Input

	ToSpatialSpecArrayOutput() SpatialSpecArrayOutput
	ToSpatialSpecArrayOutputWithContext(context.Context) SpatialSpecArrayOutput
}

type SpatialSpecArray []SpatialSpecInput

func (SpatialSpecArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpec)(nil)).Elem()
}

func (i SpatialSpecArray) ToSpatialSpecArrayOutput() SpatialSpecArrayOutput {
	return i.ToSpatialSpecArrayOutputWithContext(context.Background())
}

func (i SpatialSpecArray) ToSpatialSpecArrayOutputWithContext(ctx context.Context) SpatialSpecArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialSpecArrayOutput)
}

type SpatialSpecOutput struct{ *pulumi.OutputState }

func (SpatialSpecOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpec)(nil)).Elem()
}

func (o SpatialSpecOutput) ToSpatialSpecOutput() SpatialSpecOutput {
	return o
}

func (o SpatialSpecOutput) ToSpatialSpecOutputWithContext(ctx context.Context) SpatialSpecOutput {
	return o
}

func (o SpatialSpecOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SpatialSpec) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SpatialSpecOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SpatialSpec) []string { return v.Types }).(pulumi.StringArrayOutput)
}

type SpatialSpecArrayOutput struct{ *pulumi.OutputState }

func (SpatialSpecArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpec)(nil)).Elem()
}

func (o SpatialSpecArrayOutput) ToSpatialSpecArrayOutput() SpatialSpecArrayOutput {
	return o
}

func (o SpatialSpecArrayOutput) ToSpatialSpecArrayOutputWithContext(ctx context.Context) SpatialSpecArrayOutput {
	return o
}

func (o SpatialSpecArrayOutput) Index(i pulumi.IntInput) SpatialSpecOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpatialSpec {
		return vs[0].([]SpatialSpec)[vs[1].(int)]
	}).(SpatialSpecOutput)
}

type SpatialSpecResponse struct {
	Path  *string  `pulumi:"path"`
	Types []string `pulumi:"types"`
}





type SpatialSpecResponseInput interface {
	pulumi.Input

	ToSpatialSpecResponseOutput() SpatialSpecResponseOutput
	ToSpatialSpecResponseOutputWithContext(context.Context) SpatialSpecResponseOutput
}

type SpatialSpecResponseArgs struct {
	Path  pulumi.StringPtrInput   `pulumi:"path"`
	Types pulumi.StringArrayInput `pulumi:"types"`
}

func (SpatialSpecResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpecResponse)(nil)).Elem()
}

func (i SpatialSpecResponseArgs) ToSpatialSpecResponseOutput() SpatialSpecResponseOutput {
	return i.ToSpatialSpecResponseOutputWithContext(context.Background())
}

func (i SpatialSpecResponseArgs) ToSpatialSpecResponseOutputWithContext(ctx context.Context) SpatialSpecResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialSpecResponseOutput)
}





type SpatialSpecResponseArrayInput interface {
	pulumi.Input

	ToSpatialSpecResponseArrayOutput() SpatialSpecResponseArrayOutput
	ToSpatialSpecResponseArrayOutputWithContext(context.Context) SpatialSpecResponseArrayOutput
}

type SpatialSpecResponseArray []SpatialSpecResponseInput

func (SpatialSpecResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpecResponse)(nil)).Elem()
}

func (i SpatialSpecResponseArray) ToSpatialSpecResponseArrayOutput() SpatialSpecResponseArrayOutput {
	return i.ToSpatialSpecResponseArrayOutputWithContext(context.Background())
}

func (i SpatialSpecResponseArray) ToSpatialSpecResponseArrayOutputWithContext(ctx context.Context) SpatialSpecResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SpatialSpecResponseArrayOutput)
}

type SpatialSpecResponseOutput struct{ *pulumi.OutputState }

func (SpatialSpecResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SpatialSpecResponse)(nil)).Elem()
}

func (o SpatialSpecResponseOutput) ToSpatialSpecResponseOutput() SpatialSpecResponseOutput {
	return o
}

func (o SpatialSpecResponseOutput) ToSpatialSpecResponseOutputWithContext(ctx context.Context) SpatialSpecResponseOutput {
	return o
}

func (o SpatialSpecResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SpatialSpecResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SpatialSpecResponseOutput) Types() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SpatialSpecResponse) []string { return v.Types }).(pulumi.StringArrayOutput)
}

type SpatialSpecResponseArrayOutput struct{ *pulumi.OutputState }

func (SpatialSpecResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SpatialSpecResponse)(nil)).Elem()
}

func (o SpatialSpecResponseArrayOutput) ToSpatialSpecResponseArrayOutput() SpatialSpecResponseArrayOutput {
	return o
}

func (o SpatialSpecResponseArrayOutput) ToSpatialSpecResponseArrayOutputWithContext(ctx context.Context) SpatialSpecResponseArrayOutput {
	return o
}

func (o SpatialSpecResponseArrayOutput) Index(i pulumi.IntInput) SpatialSpecResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SpatialSpecResponse {
		return vs[0].([]SpatialSpecResponse)[vs[1].(int)]
	}).(SpatialSpecResponseOutput)
}

type SqlContainerGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type SqlContainerGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToSqlContainerGetPropertiesResponseOptionsOutput() SqlContainerGetPropertiesResponseOptionsOutput
	ToSqlContainerGetPropertiesResponseOptionsOutputWithContext(context.Context) SqlContainerGetPropertiesResponseOptionsOutput
}

type SqlContainerGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (SqlContainerGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerGetPropertiesResponseOptions)(nil)).Elem()
}

func (i SqlContainerGetPropertiesResponseOptionsArgs) ToSqlContainerGetPropertiesResponseOptionsOutput() SqlContainerGetPropertiesResponseOptionsOutput {
	return i.ToSqlContainerGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i SqlContainerGetPropertiesResponseOptionsArgs) ToSqlContainerGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerGetPropertiesResponseOptionsOutput)
}

func (i SqlContainerGetPropertiesResponseOptionsArgs) ToSqlContainerGetPropertiesResponseOptionsPtrOutput() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return i.ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i SqlContainerGetPropertiesResponseOptionsArgs) ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerGetPropertiesResponseOptionsOutput).ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type SqlContainerGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToSqlContainerGetPropertiesResponseOptionsPtrOutput() SqlContainerGetPropertiesResponseOptionsPtrOutput
	ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) SqlContainerGetPropertiesResponseOptionsPtrOutput
}

type sqlContainerGetPropertiesResponseOptionsPtrType SqlContainerGetPropertiesResponseOptionsArgs

func SqlContainerGetPropertiesResponseOptionsPtr(v *SqlContainerGetPropertiesResponseOptionsArgs) SqlContainerGetPropertiesResponseOptionsPtrInput {
	return (*sqlContainerGetPropertiesResponseOptionsPtrType)(v)
}

func (*sqlContainerGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *sqlContainerGetPropertiesResponseOptionsPtrType) ToSqlContainerGetPropertiesResponseOptionsPtrOutput() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return i.ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *sqlContainerGetPropertiesResponseOptionsPtrType) ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerGetPropertiesResponseOptionsPtrOutput)
}

type SqlContainerGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) ToSqlContainerGetPropertiesResponseOptionsOutput() SqlContainerGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) ToSqlContainerGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) ToSqlContainerGetPropertiesResponseOptionsPtrOutput() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o.ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlContainerGetPropertiesResponseOptions) *SqlContainerGetPropertiesResponseOptions {
		return &v
	}).(SqlContainerGetPropertiesResponseOptionsPtrOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type SqlContainerGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) ToSqlContainerGetPropertiesResponseOptionsPtrOutput() SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) ToSqlContainerGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) Elem() SqlContainerGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseOptions) SqlContainerGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret SqlContainerGetPropertiesResponseOptions
		return ret
	}).(SqlContainerGetPropertiesResponseOptionsOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type SqlContainerGetPropertiesResponseResource struct {
	AnalyticalStorageTtl     *float64                          `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy *ConflictResolutionPolicyResponse `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                              `pulumi:"defaultTtl"`
	Etag                     string                            `pulumi:"etag"`
	Id                       string                            `pulumi:"id"`
	IndexingPolicy           *IndexingPolicyResponse           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKeyResponse    `pulumi:"partitionKey"`
	Rid                      string                            `pulumi:"rid"`
	Ts                       float64                           `pulumi:"ts"`
	UniqueKeyPolicy          *UniqueKeyPolicyResponse          `pulumi:"uniqueKeyPolicy"`
}


func (val *SqlContainerGetPropertiesResponseResource) Defaults() *SqlContainerGetPropertiesResponseResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type SqlContainerGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToSqlContainerGetPropertiesResponseResourceOutput() SqlContainerGetPropertiesResponseResourceOutput
	ToSqlContainerGetPropertiesResponseResourceOutputWithContext(context.Context) SqlContainerGetPropertiesResponseResourceOutput
}

type SqlContainerGetPropertiesResponseResourceArgs struct {
	AnalyticalStorageTtl     pulumi.Float64PtrInput                   `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy ConflictResolutionPolicyResponsePtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput                       `pulumi:"defaultTtl"`
	Etag                     pulumi.StringInput                       `pulumi:"etag"`
	Id                       pulumi.StringInput                       `pulumi:"id"`
	IndexingPolicy           IndexingPolicyResponsePtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyResponsePtrInput    `pulumi:"partitionKey"`
	Rid                      pulumi.StringInput                       `pulumi:"rid"`
	Ts                       pulumi.Float64Input                      `pulumi:"ts"`
	UniqueKeyPolicy          UniqueKeyPolicyResponsePtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (SqlContainerGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerGetPropertiesResponseResource)(nil)).Elem()
}

func (i SqlContainerGetPropertiesResponseResourceArgs) ToSqlContainerGetPropertiesResponseResourceOutput() SqlContainerGetPropertiesResponseResourceOutput {
	return i.ToSqlContainerGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i SqlContainerGetPropertiesResponseResourceArgs) ToSqlContainerGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerGetPropertiesResponseResourceOutput)
}

func (i SqlContainerGetPropertiesResponseResourceArgs) ToSqlContainerGetPropertiesResponseResourcePtrOutput() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i SqlContainerGetPropertiesResponseResourceArgs) ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerGetPropertiesResponseResourceOutput).ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type SqlContainerGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToSqlContainerGetPropertiesResponseResourcePtrOutput() SqlContainerGetPropertiesResponseResourcePtrOutput
	ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(context.Context) SqlContainerGetPropertiesResponseResourcePtrOutput
}

type sqlContainerGetPropertiesResponseResourcePtrType SqlContainerGetPropertiesResponseResourceArgs

func SqlContainerGetPropertiesResponseResourcePtr(v *SqlContainerGetPropertiesResponseResourceArgs) SqlContainerGetPropertiesResponseResourcePtrInput {
	return (*sqlContainerGetPropertiesResponseResourcePtrType)(v)
}

func (*sqlContainerGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerGetPropertiesResponseResource)(nil)).Elem()
}

func (i *sqlContainerGetPropertiesResponseResourcePtrType) ToSqlContainerGetPropertiesResponseResourcePtrOutput() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *sqlContainerGetPropertiesResponseResourcePtrType) ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerGetPropertiesResponseResourcePtrOutput)
}

type SqlContainerGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ToSqlContainerGetPropertiesResponseResourceOutput() SqlContainerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ToSqlContainerGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ToSqlContainerGetPropertiesResponseResourcePtrOutput() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o.ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlContainerGetPropertiesResponseResource) *SqlContainerGetPropertiesResponseResource {
		return &v
	}).(SqlContainerGetPropertiesResponseResourcePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *float64 { return v.AnalyticalStorageTtl }).(pulumi.Float64PtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *IndexingPolicyResponse { return v.IndexingPolicy }).(IndexingPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

func (o SqlContainerGetPropertiesResponseResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v SqlContainerGetPropertiesResponseResource) *UniqueKeyPolicyResponse { return v.UniqueKeyPolicy }).(UniqueKeyPolicyResponsePtrOutput)
}

type SqlContainerGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlContainerGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) ToSqlContainerGetPropertiesResponseResourcePtrOutput() SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) ToSqlContainerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlContainerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Elem() SqlContainerGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) SqlContainerGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlContainerGetPropertiesResponseResource
		return ret
	}).(SqlContainerGetPropertiesResponseResourceOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) ConflictResolutionPolicy() ConflictResolutionPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *ConflictResolutionPolicyResponse {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) IndexingPolicy() IndexingPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *IndexingPolicyResponse {
		if v == nil {
			return nil
		}
		return v.IndexingPolicy
	}).(IndexingPolicyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) PartitionKey() ContainerPartitionKeyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *ContainerPartitionKeyResponse {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(ContainerPartitionKeyResponsePtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

func (o SqlContainerGetPropertiesResponseResourcePtrOutput) UniqueKeyPolicy() UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyT(func(v *SqlContainerGetPropertiesResponseResource) *UniqueKeyPolicyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeyPolicy
	}).(UniqueKeyPolicyResponsePtrOutput)
}

type SqlContainerResource struct {
	AnalyticalStorageTtl     *float64                  `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy *ConflictResolutionPolicy `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               *int                      `pulumi:"defaultTtl"`
	Id                       string                    `pulumi:"id"`
	IndexingPolicy           *IndexingPolicy           `pulumi:"indexingPolicy"`
	PartitionKey             *ContainerPartitionKey    `pulumi:"partitionKey"`
	UniqueKeyPolicy          *UniqueKeyPolicy          `pulumi:"uniqueKeyPolicy"`
}


func (val *SqlContainerResource) Defaults() *SqlContainerResource {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.ConflictResolutionPolicy = tmp.ConflictResolutionPolicy.Defaults()

	tmp.IndexingPolicy = tmp.IndexingPolicy.Defaults()

	tmp.PartitionKey = tmp.PartitionKey.Defaults()

	return &tmp
}





type SqlContainerResourceInput interface {
	pulumi.Input

	ToSqlContainerResourceOutput() SqlContainerResourceOutput
	ToSqlContainerResourceOutputWithContext(context.Context) SqlContainerResourceOutput
}

type SqlContainerResourceArgs struct {
	AnalyticalStorageTtl     pulumi.Float64PtrInput           `pulumi:"analyticalStorageTtl"`
	ConflictResolutionPolicy ConflictResolutionPolicyPtrInput `pulumi:"conflictResolutionPolicy"`
	DefaultTtl               pulumi.IntPtrInput               `pulumi:"defaultTtl"`
	Id                       pulumi.StringInput               `pulumi:"id"`
	IndexingPolicy           IndexingPolicyPtrInput           `pulumi:"indexingPolicy"`
	PartitionKey             ContainerPartitionKeyPtrInput    `pulumi:"partitionKey"`
	UniqueKeyPolicy          UniqueKeyPolicyPtrInput          `pulumi:"uniqueKeyPolicy"`
}

func (SqlContainerResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerResource)(nil)).Elem()
}

func (i SqlContainerResourceArgs) ToSqlContainerResourceOutput() SqlContainerResourceOutput {
	return i.ToSqlContainerResourceOutputWithContext(context.Background())
}

func (i SqlContainerResourceArgs) ToSqlContainerResourceOutputWithContext(ctx context.Context) SqlContainerResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerResourceOutput)
}

func (i SqlContainerResourceArgs) ToSqlContainerResourcePtrOutput() SqlContainerResourcePtrOutput {
	return i.ToSqlContainerResourcePtrOutputWithContext(context.Background())
}

func (i SqlContainerResourceArgs) ToSqlContainerResourcePtrOutputWithContext(ctx context.Context) SqlContainerResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerResourceOutput).ToSqlContainerResourcePtrOutputWithContext(ctx)
}









type SqlContainerResourcePtrInput interface {
	pulumi.Input

	ToSqlContainerResourcePtrOutput() SqlContainerResourcePtrOutput
	ToSqlContainerResourcePtrOutputWithContext(context.Context) SqlContainerResourcePtrOutput
}

type sqlContainerResourcePtrType SqlContainerResourceArgs

func SqlContainerResourcePtr(v *SqlContainerResourceArgs) SqlContainerResourcePtrInput {
	return (*sqlContainerResourcePtrType)(v)
}

func (*sqlContainerResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerResource)(nil)).Elem()
}

func (i *sqlContainerResourcePtrType) ToSqlContainerResourcePtrOutput() SqlContainerResourcePtrOutput {
	return i.ToSqlContainerResourcePtrOutputWithContext(context.Background())
}

func (i *sqlContainerResourcePtrType) ToSqlContainerResourcePtrOutputWithContext(ctx context.Context) SqlContainerResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlContainerResourcePtrOutput)
}

type SqlContainerResourceOutput struct{ *pulumi.OutputState }

func (SqlContainerResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlContainerResource)(nil)).Elem()
}

func (o SqlContainerResourceOutput) ToSqlContainerResourceOutput() SqlContainerResourceOutput {
	return o
}

func (o SqlContainerResourceOutput) ToSqlContainerResourceOutputWithContext(ctx context.Context) SqlContainerResourceOutput {
	return o
}

func (o SqlContainerResourceOutput) ToSqlContainerResourcePtrOutput() SqlContainerResourcePtrOutput {
	return o.ToSqlContainerResourcePtrOutputWithContext(context.Background())
}

func (o SqlContainerResourceOutput) ToSqlContainerResourcePtrOutputWithContext(ctx context.Context) SqlContainerResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlContainerResource) *SqlContainerResource {
		return &v
	}).(SqlContainerResourcePtrOutput)
}

func (o SqlContainerResourceOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *float64 { return v.AnalyticalStorageTtl }).(pulumi.Float64PtrOutput)
}

func (o SqlContainerResourceOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *ConflictResolutionPolicy { return v.ConflictResolutionPolicy }).(ConflictResolutionPolicyPtrOutput)
}

func (o SqlContainerResourceOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *int { return v.DefaultTtl }).(pulumi.IntPtrOutput)
}

func (o SqlContainerResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlContainerResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlContainerResourceOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *IndexingPolicy { return v.IndexingPolicy }).(IndexingPolicyPtrOutput)
}

func (o SqlContainerResourceOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *ContainerPartitionKey { return v.PartitionKey }).(ContainerPartitionKeyPtrOutput)
}

func (o SqlContainerResourceOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v SqlContainerResource) *UniqueKeyPolicy { return v.UniqueKeyPolicy }).(UniqueKeyPolicyPtrOutput)
}

type SqlContainerResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlContainerResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlContainerResource)(nil)).Elem()
}

func (o SqlContainerResourcePtrOutput) ToSqlContainerResourcePtrOutput() SqlContainerResourcePtrOutput {
	return o
}

func (o SqlContainerResourcePtrOutput) ToSqlContainerResourcePtrOutputWithContext(ctx context.Context) SqlContainerResourcePtrOutput {
	return o
}

func (o SqlContainerResourcePtrOutput) Elem() SqlContainerResourceOutput {
	return o.ApplyT(func(v *SqlContainerResource) SqlContainerResource {
		if v != nil {
			return *v
		}
		var ret SqlContainerResource
		return ret
	}).(SqlContainerResourceOutput)
}

func (o SqlContainerResourcePtrOutput) AnalyticalStorageTtl() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *float64 {
		if v == nil {
			return nil
		}
		return v.AnalyticalStorageTtl
	}).(pulumi.Float64PtrOutput)
}

func (o SqlContainerResourcePtrOutput) ConflictResolutionPolicy() ConflictResolutionPolicyPtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *ConflictResolutionPolicy {
		if v == nil {
			return nil
		}
		return v.ConflictResolutionPolicy
	}).(ConflictResolutionPolicyPtrOutput)
}

func (o SqlContainerResourcePtrOutput) DefaultTtl() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *int {
		if v == nil {
			return nil
		}
		return v.DefaultTtl
	}).(pulumi.IntPtrOutput)
}

func (o SqlContainerResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlContainerResourcePtrOutput) IndexingPolicy() IndexingPolicyPtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *IndexingPolicy {
		if v == nil {
			return nil
		}
		return v.IndexingPolicy
	}).(IndexingPolicyPtrOutput)
}

func (o SqlContainerResourcePtrOutput) PartitionKey() ContainerPartitionKeyPtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *ContainerPartitionKey {
		if v == nil {
			return nil
		}
		return v.PartitionKey
	}).(ContainerPartitionKeyPtrOutput)
}

func (o SqlContainerResourcePtrOutput) UniqueKeyPolicy() UniqueKeyPolicyPtrOutput {
	return o.ApplyT(func(v *SqlContainerResource) *UniqueKeyPolicy {
		if v == nil {
			return nil
		}
		return v.UniqueKeyPolicy
	}).(UniqueKeyPolicyPtrOutput)
}

type SqlDatabaseGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type SqlDatabaseGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToSqlDatabaseGetPropertiesResponseOptionsOutput() SqlDatabaseGetPropertiesResponseOptionsOutput
	ToSqlDatabaseGetPropertiesResponseOptionsOutputWithContext(context.Context) SqlDatabaseGetPropertiesResponseOptionsOutput
}

type SqlDatabaseGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (SqlDatabaseGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (i SqlDatabaseGetPropertiesResponseOptionsArgs) ToSqlDatabaseGetPropertiesResponseOptionsOutput() SqlDatabaseGetPropertiesResponseOptionsOutput {
	return i.ToSqlDatabaseGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i SqlDatabaseGetPropertiesResponseOptionsArgs) ToSqlDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseGetPropertiesResponseOptionsOutput)
}

func (i SqlDatabaseGetPropertiesResponseOptionsArgs) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutput() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return i.ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i SqlDatabaseGetPropertiesResponseOptionsArgs) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseGetPropertiesResponseOptionsOutput).ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type SqlDatabaseGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToSqlDatabaseGetPropertiesResponseOptionsPtrOutput() SqlDatabaseGetPropertiesResponseOptionsPtrOutput
	ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) SqlDatabaseGetPropertiesResponseOptionsPtrOutput
}

type sqlDatabaseGetPropertiesResponseOptionsPtrType SqlDatabaseGetPropertiesResponseOptionsArgs

func SqlDatabaseGetPropertiesResponseOptionsPtr(v *SqlDatabaseGetPropertiesResponseOptionsArgs) SqlDatabaseGetPropertiesResponseOptionsPtrInput {
	return (*sqlDatabaseGetPropertiesResponseOptionsPtrType)(v)
}

func (*sqlDatabaseGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *sqlDatabaseGetPropertiesResponseOptionsPtrType) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutput() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return i.ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *sqlDatabaseGetPropertiesResponseOptionsPtrType) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseGetPropertiesResponseOptionsPtrOutput)
}

type SqlDatabaseGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) ToSqlDatabaseGetPropertiesResponseOptionsOutput() SqlDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) ToSqlDatabaseGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutput() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlDatabaseGetPropertiesResponseOptions) *SqlDatabaseGetPropertiesResponseOptions {
		return &v
	}).(SqlDatabaseGetPropertiesResponseOptionsPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse { return v.AutoscaleSettings }).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type SqlDatabaseGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseGetPropertiesResponseOptions)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutput() SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) ToSqlDatabaseGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) Elem() SqlDatabaseGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseOptions) SqlDatabaseGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseGetPropertiesResponseOptions
		return ret
	}).(SqlDatabaseGetPropertiesResponseOptionsOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type SqlDatabaseGetPropertiesResponseResource struct {
	Colls *string `pulumi:"colls"`
	Etag  string  `pulumi:"etag"`
	Id    string  `pulumi:"id"`
	Rid   string  `pulumi:"rid"`
	Ts    float64 `pulumi:"ts"`
	Users *string `pulumi:"users"`
}





type SqlDatabaseGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToSqlDatabaseGetPropertiesResponseResourceOutput() SqlDatabaseGetPropertiesResponseResourceOutput
	ToSqlDatabaseGetPropertiesResponseResourceOutputWithContext(context.Context) SqlDatabaseGetPropertiesResponseResourceOutput
}

type SqlDatabaseGetPropertiesResponseResourceArgs struct {
	Colls pulumi.StringPtrInput `pulumi:"colls"`
	Etag  pulumi.StringInput    `pulumi:"etag"`
	Id    pulumi.StringInput    `pulumi:"id"`
	Rid   pulumi.StringInput    `pulumi:"rid"`
	Ts    pulumi.Float64Input   `pulumi:"ts"`
	Users pulumi.StringPtrInput `pulumi:"users"`
}

func (SqlDatabaseGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (i SqlDatabaseGetPropertiesResponseResourceArgs) ToSqlDatabaseGetPropertiesResponseResourceOutput() SqlDatabaseGetPropertiesResponseResourceOutput {
	return i.ToSqlDatabaseGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i SqlDatabaseGetPropertiesResponseResourceArgs) ToSqlDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseGetPropertiesResponseResourceOutput)
}

func (i SqlDatabaseGetPropertiesResponseResourceArgs) ToSqlDatabaseGetPropertiesResponseResourcePtrOutput() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i SqlDatabaseGetPropertiesResponseResourceArgs) ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseGetPropertiesResponseResourceOutput).ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type SqlDatabaseGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToSqlDatabaseGetPropertiesResponseResourcePtrOutput() SqlDatabaseGetPropertiesResponseResourcePtrOutput
	ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Context) SqlDatabaseGetPropertiesResponseResourcePtrOutput
}

type sqlDatabaseGetPropertiesResponseResourcePtrType SqlDatabaseGetPropertiesResponseResourceArgs

func SqlDatabaseGetPropertiesResponseResourcePtr(v *SqlDatabaseGetPropertiesResponseResourceArgs) SqlDatabaseGetPropertiesResponseResourcePtrInput {
	return (*sqlDatabaseGetPropertiesResponseResourcePtrType)(v)
}

func (*sqlDatabaseGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (i *sqlDatabaseGetPropertiesResponseResourcePtrType) ToSqlDatabaseGetPropertiesResponseResourcePtrOutput() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *sqlDatabaseGetPropertiesResponseResourcePtrType) ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseGetPropertiesResponseResourcePtrOutput)
}

type SqlDatabaseGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) ToSqlDatabaseGetPropertiesResponseResourceOutput() SqlDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) ToSqlDatabaseGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) ToSqlDatabaseGetPropertiesResponseResourcePtrOutput() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlDatabaseGetPropertiesResponseResource) *SqlDatabaseGetPropertiesResponseResource {
		return &v
	}).(SqlDatabaseGetPropertiesResponseResourcePtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Colls() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) *string { return v.Colls }).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

func (o SqlDatabaseGetPropertiesResponseResourceOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlDatabaseGetPropertiesResponseResource) *string { return v.Users }).(pulumi.StringPtrOutput)
}

type SqlDatabaseGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) ToSqlDatabaseGetPropertiesResponseResourcePtrOutput() SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) ToSqlDatabaseGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Elem() SqlDatabaseGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) SqlDatabaseGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseGetPropertiesResponseResource
		return ret
	}).(SqlDatabaseGetPropertiesResponseResourceOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Colls() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Colls
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

func (o SqlDatabaseGetPropertiesResponseResourcePtrOutput) Users() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Users
	}).(pulumi.StringPtrOutput)
}

type SqlDatabaseResource struct {
	Id string `pulumi:"id"`
}





type SqlDatabaseResourceInput interface {
	pulumi.Input

	ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput
	ToSqlDatabaseResourceOutputWithContext(context.Context) SqlDatabaseResourceOutput
}

type SqlDatabaseResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (SqlDatabaseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResource)(nil)).Elem()
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput {
	return i.ToSqlDatabaseResourceOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourceOutputWithContext(ctx context.Context) SqlDatabaseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceOutput)
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourcePtrOutput() SqlDatabaseResourcePtrOutput {
	return i.ToSqlDatabaseResourcePtrOutputWithContext(context.Background())
}

func (i SqlDatabaseResourceArgs) ToSqlDatabaseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourceOutput).ToSqlDatabaseResourcePtrOutputWithContext(ctx)
}









type SqlDatabaseResourcePtrInput interface {
	pulumi.Input

	ToSqlDatabaseResourcePtrOutput() SqlDatabaseResourcePtrOutput
	ToSqlDatabaseResourcePtrOutputWithContext(context.Context) SqlDatabaseResourcePtrOutput
}

type sqlDatabaseResourcePtrType SqlDatabaseResourceArgs

func SqlDatabaseResourcePtr(v *SqlDatabaseResourceArgs) SqlDatabaseResourcePtrInput {
	return (*sqlDatabaseResourcePtrType)(v)
}

func (*sqlDatabaseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseResource)(nil)).Elem()
}

func (i *sqlDatabaseResourcePtrType) ToSqlDatabaseResourcePtrOutput() SqlDatabaseResourcePtrOutput {
	return i.ToSqlDatabaseResourcePtrOutputWithContext(context.Background())
}

func (i *sqlDatabaseResourcePtrType) ToSqlDatabaseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlDatabaseResourcePtrOutput)
}

type SqlDatabaseResourceOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlDatabaseResource)(nil)).Elem()
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourceOutput() SqlDatabaseResourceOutput {
	return o
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourceOutputWithContext(ctx context.Context) SqlDatabaseResourceOutput {
	return o
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourcePtrOutput() SqlDatabaseResourcePtrOutput {
	return o.ToSqlDatabaseResourcePtrOutputWithContext(context.Background())
}

func (o SqlDatabaseResourceOutput) ToSqlDatabaseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlDatabaseResource) *SqlDatabaseResource {
		return &v
	}).(SqlDatabaseResourcePtrOutput)
}

func (o SqlDatabaseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlDatabaseResource) string { return v.Id }).(pulumi.StringOutput)
}

type SqlDatabaseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlDatabaseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlDatabaseResource)(nil)).Elem()
}

func (o SqlDatabaseResourcePtrOutput) ToSqlDatabaseResourcePtrOutput() SqlDatabaseResourcePtrOutput {
	return o
}

func (o SqlDatabaseResourcePtrOutput) ToSqlDatabaseResourcePtrOutputWithContext(ctx context.Context) SqlDatabaseResourcePtrOutput {
	return o
}

func (o SqlDatabaseResourcePtrOutput) Elem() SqlDatabaseResourceOutput {
	return o.ApplyT(func(v *SqlDatabaseResource) SqlDatabaseResource {
		if v != nil {
			return *v
		}
		var ret SqlDatabaseResource
		return ret
	}).(SqlDatabaseResourceOutput)
}

func (o SqlDatabaseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlDatabaseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SqlStoredProcedureGetPropertiesResponseResource struct {
	Body *string `pulumi:"body"`
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}





type SqlStoredProcedureGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToSqlStoredProcedureGetPropertiesResponseResourceOutput() SqlStoredProcedureGetPropertiesResponseResourceOutput
	ToSqlStoredProcedureGetPropertiesResponseResourceOutputWithContext(context.Context) SqlStoredProcedureGetPropertiesResponseResourceOutput
}

type SqlStoredProcedureGetPropertiesResponseResourceArgs struct {
	Body pulumi.StringPtrInput `pulumi:"body"`
	Etag pulumi.StringInput    `pulumi:"etag"`
	Id   pulumi.StringInput    `pulumi:"id"`
	Rid  pulumi.StringInput    `pulumi:"rid"`
	Ts   pulumi.Float64Input   `pulumi:"ts"`
}

func (SqlStoredProcedureGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureGetPropertiesResponseResource)(nil)).Elem()
}

func (i SqlStoredProcedureGetPropertiesResponseResourceArgs) ToSqlStoredProcedureGetPropertiesResponseResourceOutput() SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return i.ToSqlStoredProcedureGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i SqlStoredProcedureGetPropertiesResponseResourceArgs) ToSqlStoredProcedureGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureGetPropertiesResponseResourceOutput)
}

func (i SqlStoredProcedureGetPropertiesResponseResourceArgs) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutput() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i SqlStoredProcedureGetPropertiesResponseResourceArgs) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureGetPropertiesResponseResourceOutput).ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type SqlStoredProcedureGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutput() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput
	ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(context.Context) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput
}

type sqlStoredProcedureGetPropertiesResponseResourcePtrType SqlStoredProcedureGetPropertiesResponseResourceArgs

func SqlStoredProcedureGetPropertiesResponseResourcePtr(v *SqlStoredProcedureGetPropertiesResponseResourceArgs) SqlStoredProcedureGetPropertiesResponseResourcePtrInput {
	return (*sqlStoredProcedureGetPropertiesResponseResourcePtrType)(v)
}

func (*sqlStoredProcedureGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedureGetPropertiesResponseResource)(nil)).Elem()
}

func (i *sqlStoredProcedureGetPropertiesResponseResourcePtrType) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutput() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *sqlStoredProcedureGetPropertiesResponseResourcePtrType) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput)
}

type SqlStoredProcedureGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) ToSqlStoredProcedureGetPropertiesResponseResourceOutput() SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) ToSqlStoredProcedureGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutput() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o.ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlStoredProcedureGetPropertiesResponseResource) *SqlStoredProcedureGetPropertiesResponseResource {
		return &v
	}).(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlStoredProcedureGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type SqlStoredProcedureGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedureGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutput() SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) ToSqlStoredProcedureGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Elem() SqlStoredProcedureGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) SqlStoredProcedureGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlStoredProcedureGetPropertiesResponseResource
		return ret
	}).(SqlStoredProcedureGetPropertiesResponseResourceOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type SqlStoredProcedureResource struct {
	Body *string `pulumi:"body"`
	Id   string  `pulumi:"id"`
}





type SqlStoredProcedureResourceInput interface {
	pulumi.Input

	ToSqlStoredProcedureResourceOutput() SqlStoredProcedureResourceOutput
	ToSqlStoredProcedureResourceOutputWithContext(context.Context) SqlStoredProcedureResourceOutput
}

type SqlStoredProcedureResourceArgs struct {
	Body pulumi.StringPtrInput `pulumi:"body"`
	Id   pulumi.StringInput    `pulumi:"id"`
}

func (SqlStoredProcedureResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureResource)(nil)).Elem()
}

func (i SqlStoredProcedureResourceArgs) ToSqlStoredProcedureResourceOutput() SqlStoredProcedureResourceOutput {
	return i.ToSqlStoredProcedureResourceOutputWithContext(context.Background())
}

func (i SqlStoredProcedureResourceArgs) ToSqlStoredProcedureResourceOutputWithContext(ctx context.Context) SqlStoredProcedureResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureResourceOutput)
}

func (i SqlStoredProcedureResourceArgs) ToSqlStoredProcedureResourcePtrOutput() SqlStoredProcedureResourcePtrOutput {
	return i.ToSqlStoredProcedureResourcePtrOutputWithContext(context.Background())
}

func (i SqlStoredProcedureResourceArgs) ToSqlStoredProcedureResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureResourceOutput).ToSqlStoredProcedureResourcePtrOutputWithContext(ctx)
}









type SqlStoredProcedureResourcePtrInput interface {
	pulumi.Input

	ToSqlStoredProcedureResourcePtrOutput() SqlStoredProcedureResourcePtrOutput
	ToSqlStoredProcedureResourcePtrOutputWithContext(context.Context) SqlStoredProcedureResourcePtrOutput
}

type sqlStoredProcedureResourcePtrType SqlStoredProcedureResourceArgs

func SqlStoredProcedureResourcePtr(v *SqlStoredProcedureResourceArgs) SqlStoredProcedureResourcePtrInput {
	return (*sqlStoredProcedureResourcePtrType)(v)
}

func (*sqlStoredProcedureResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedureResource)(nil)).Elem()
}

func (i *sqlStoredProcedureResourcePtrType) ToSqlStoredProcedureResourcePtrOutput() SqlStoredProcedureResourcePtrOutput {
	return i.ToSqlStoredProcedureResourcePtrOutputWithContext(context.Background())
}

func (i *sqlStoredProcedureResourcePtrType) ToSqlStoredProcedureResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStoredProcedureResourcePtrOutput)
}

type SqlStoredProcedureResourceOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStoredProcedureResource)(nil)).Elem()
}

func (o SqlStoredProcedureResourceOutput) ToSqlStoredProcedureResourceOutput() SqlStoredProcedureResourceOutput {
	return o
}

func (o SqlStoredProcedureResourceOutput) ToSqlStoredProcedureResourceOutputWithContext(ctx context.Context) SqlStoredProcedureResourceOutput {
	return o
}

func (o SqlStoredProcedureResourceOutput) ToSqlStoredProcedureResourcePtrOutput() SqlStoredProcedureResourcePtrOutput {
	return o.ToSqlStoredProcedureResourcePtrOutputWithContext(context.Background())
}

func (o SqlStoredProcedureResourceOutput) ToSqlStoredProcedureResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlStoredProcedureResource) *SqlStoredProcedureResource {
		return &v
	}).(SqlStoredProcedureResourcePtrOutput)
}

func (o SqlStoredProcedureResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlStoredProcedureResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlStoredProcedureResource) string { return v.Id }).(pulumi.StringOutput)
}

type SqlStoredProcedureResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlStoredProcedureResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStoredProcedureResource)(nil)).Elem()
}

func (o SqlStoredProcedureResourcePtrOutput) ToSqlStoredProcedureResourcePtrOutput() SqlStoredProcedureResourcePtrOutput {
	return o
}

func (o SqlStoredProcedureResourcePtrOutput) ToSqlStoredProcedureResourcePtrOutputWithContext(ctx context.Context) SqlStoredProcedureResourcePtrOutput {
	return o
}

func (o SqlStoredProcedureResourcePtrOutput) Elem() SqlStoredProcedureResourceOutput {
	return o.ApplyT(func(v *SqlStoredProcedureResource) SqlStoredProcedureResource {
		if v != nil {
			return *v
		}
		var ret SqlStoredProcedureResource
		return ret
	}).(SqlStoredProcedureResourceOutput)
}

func (o SqlStoredProcedureResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlStoredProcedureResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStoredProcedureResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type SqlTriggerGetPropertiesResponseResource struct {
	Body             *string `pulumi:"body"`
	Etag             string  `pulumi:"etag"`
	Id               string  `pulumi:"id"`
	Rid              string  `pulumi:"rid"`
	TriggerOperation *string `pulumi:"triggerOperation"`
	TriggerType      *string `pulumi:"triggerType"`
	Ts               float64 `pulumi:"ts"`
}





type SqlTriggerGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToSqlTriggerGetPropertiesResponseResourceOutput() SqlTriggerGetPropertiesResponseResourceOutput
	ToSqlTriggerGetPropertiesResponseResourceOutputWithContext(context.Context) SqlTriggerGetPropertiesResponseResourceOutput
}

type SqlTriggerGetPropertiesResponseResourceArgs struct {
	Body             pulumi.StringPtrInput `pulumi:"body"`
	Etag             pulumi.StringInput    `pulumi:"etag"`
	Id               pulumi.StringInput    `pulumi:"id"`
	Rid              pulumi.StringInput    `pulumi:"rid"`
	TriggerOperation pulumi.StringPtrInput `pulumi:"triggerOperation"`
	TriggerType      pulumi.StringPtrInput `pulumi:"triggerType"`
	Ts               pulumi.Float64Input   `pulumi:"ts"`
}

func (SqlTriggerGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerGetPropertiesResponseResource)(nil)).Elem()
}

func (i SqlTriggerGetPropertiesResponseResourceArgs) ToSqlTriggerGetPropertiesResponseResourceOutput() SqlTriggerGetPropertiesResponseResourceOutput {
	return i.ToSqlTriggerGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i SqlTriggerGetPropertiesResponseResourceArgs) ToSqlTriggerGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerGetPropertiesResponseResourceOutput)
}

func (i SqlTriggerGetPropertiesResponseResourceArgs) ToSqlTriggerGetPropertiesResponseResourcePtrOutput() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i SqlTriggerGetPropertiesResponseResourceArgs) ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerGetPropertiesResponseResourceOutput).ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type SqlTriggerGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToSqlTriggerGetPropertiesResponseResourcePtrOutput() SqlTriggerGetPropertiesResponseResourcePtrOutput
	ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(context.Context) SqlTriggerGetPropertiesResponseResourcePtrOutput
}

type sqlTriggerGetPropertiesResponseResourcePtrType SqlTriggerGetPropertiesResponseResourceArgs

func SqlTriggerGetPropertiesResponseResourcePtr(v *SqlTriggerGetPropertiesResponseResourceArgs) SqlTriggerGetPropertiesResponseResourcePtrInput {
	return (*sqlTriggerGetPropertiesResponseResourcePtrType)(v)
}

func (*sqlTriggerGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlTriggerGetPropertiesResponseResource)(nil)).Elem()
}

func (i *sqlTriggerGetPropertiesResponseResourcePtrType) ToSqlTriggerGetPropertiesResponseResourcePtrOutput() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *sqlTriggerGetPropertiesResponseResourcePtrType) ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerGetPropertiesResponseResourcePtrOutput)
}

type SqlTriggerGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlTriggerGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) ToSqlTriggerGetPropertiesResponseResourceOutput() SqlTriggerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) ToSqlTriggerGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) ToSqlTriggerGetPropertiesResponseResourcePtrOutput() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o.ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlTriggerGetPropertiesResponseResource) *SqlTriggerGetPropertiesResponseResource {
		return &v
	}).(SqlTriggerGetPropertiesResponseResourcePtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) *string { return v.TriggerOperation }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) *string { return v.TriggerType }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlTriggerGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type SqlTriggerGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlTriggerGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlTriggerGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) ToSqlTriggerGetPropertiesResponseResourcePtrOutput() SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) ToSqlTriggerGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlTriggerGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Elem() SqlTriggerGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) SqlTriggerGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlTriggerGetPropertiesResponseResource
		return ret
	}).(SqlTriggerGetPropertiesResponseResourceOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.TriggerOperation
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.TriggerType
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlTriggerGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type SqlTriggerResource struct {
	Body             *string `pulumi:"body"`
	Id               string  `pulumi:"id"`
	TriggerOperation *string `pulumi:"triggerOperation"`
	TriggerType      *string `pulumi:"triggerType"`
}





type SqlTriggerResourceInput interface {
	pulumi.Input

	ToSqlTriggerResourceOutput() SqlTriggerResourceOutput
	ToSqlTriggerResourceOutputWithContext(context.Context) SqlTriggerResourceOutput
}

type SqlTriggerResourceArgs struct {
	Body             pulumi.StringPtrInput `pulumi:"body"`
	Id               pulumi.StringInput    `pulumi:"id"`
	TriggerOperation pulumi.StringPtrInput `pulumi:"triggerOperation"`
	TriggerType      pulumi.StringPtrInput `pulumi:"triggerType"`
}

func (SqlTriggerResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerResource)(nil)).Elem()
}

func (i SqlTriggerResourceArgs) ToSqlTriggerResourceOutput() SqlTriggerResourceOutput {
	return i.ToSqlTriggerResourceOutputWithContext(context.Background())
}

func (i SqlTriggerResourceArgs) ToSqlTriggerResourceOutputWithContext(ctx context.Context) SqlTriggerResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerResourceOutput)
}

func (i SqlTriggerResourceArgs) ToSqlTriggerResourcePtrOutput() SqlTriggerResourcePtrOutput {
	return i.ToSqlTriggerResourcePtrOutputWithContext(context.Background())
}

func (i SqlTriggerResourceArgs) ToSqlTriggerResourcePtrOutputWithContext(ctx context.Context) SqlTriggerResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerResourceOutput).ToSqlTriggerResourcePtrOutputWithContext(ctx)
}









type SqlTriggerResourcePtrInput interface {
	pulumi.Input

	ToSqlTriggerResourcePtrOutput() SqlTriggerResourcePtrOutput
	ToSqlTriggerResourcePtrOutputWithContext(context.Context) SqlTriggerResourcePtrOutput
}

type sqlTriggerResourcePtrType SqlTriggerResourceArgs

func SqlTriggerResourcePtr(v *SqlTriggerResourceArgs) SqlTriggerResourcePtrInput {
	return (*sqlTriggerResourcePtrType)(v)
}

func (*sqlTriggerResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlTriggerResource)(nil)).Elem()
}

func (i *sqlTriggerResourcePtrType) ToSqlTriggerResourcePtrOutput() SqlTriggerResourcePtrOutput {
	return i.ToSqlTriggerResourcePtrOutputWithContext(context.Background())
}

func (i *sqlTriggerResourcePtrType) ToSqlTriggerResourcePtrOutputWithContext(ctx context.Context) SqlTriggerResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlTriggerResourcePtrOutput)
}

type SqlTriggerResourceOutput struct{ *pulumi.OutputState }

func (SqlTriggerResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlTriggerResource)(nil)).Elem()
}

func (o SqlTriggerResourceOutput) ToSqlTriggerResourceOutput() SqlTriggerResourceOutput {
	return o
}

func (o SqlTriggerResourceOutput) ToSqlTriggerResourceOutputWithContext(ctx context.Context) SqlTriggerResourceOutput {
	return o
}

func (o SqlTriggerResourceOutput) ToSqlTriggerResourcePtrOutput() SqlTriggerResourcePtrOutput {
	return o.ToSqlTriggerResourcePtrOutputWithContext(context.Background())
}

func (o SqlTriggerResourceOutput) ToSqlTriggerResourcePtrOutputWithContext(ctx context.Context) SqlTriggerResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlTriggerResource) *SqlTriggerResource {
		return &v
	}).(SqlTriggerResourcePtrOutput)
}

func (o SqlTriggerResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlTriggerResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlTriggerResourceOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerResource) *string { return v.TriggerOperation }).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourceOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlTriggerResource) *string { return v.TriggerType }).(pulumi.StringPtrOutput)
}

type SqlTriggerResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlTriggerResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlTriggerResource)(nil)).Elem()
}

func (o SqlTriggerResourcePtrOutput) ToSqlTriggerResourcePtrOutput() SqlTriggerResourcePtrOutput {
	return o
}

func (o SqlTriggerResourcePtrOutput) ToSqlTriggerResourcePtrOutputWithContext(ctx context.Context) SqlTriggerResourcePtrOutput {
	return o
}

func (o SqlTriggerResourcePtrOutput) Elem() SqlTriggerResourceOutput {
	return o.ApplyT(func(v *SqlTriggerResource) SqlTriggerResource {
		if v != nil {
			return *v
		}
		var ret SqlTriggerResource
		return ret
	}).(SqlTriggerResourceOutput)
}

func (o SqlTriggerResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourcePtrOutput) TriggerOperation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerResource) *string {
		if v == nil {
			return nil
		}
		return v.TriggerOperation
	}).(pulumi.StringPtrOutput)
}

func (o SqlTriggerResourcePtrOutput) TriggerType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlTriggerResource) *string {
		if v == nil {
			return nil
		}
		return v.TriggerType
	}).(pulumi.StringPtrOutput)
}

type SqlUserDefinedFunctionGetPropertiesResponseResource struct {
	Body *string `pulumi:"body"`
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}





type SqlUserDefinedFunctionGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutput() SqlUserDefinedFunctionGetPropertiesResponseResourceOutput
	ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutputWithContext(context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourceOutput
}

type SqlUserDefinedFunctionGetPropertiesResponseResourceArgs struct {
	Body pulumi.StringPtrInput `pulumi:"body"`
	Etag pulumi.StringInput    `pulumi:"etag"`
	Id   pulumi.StringInput    `pulumi:"id"`
	Rid  pulumi.StringInput    `pulumi:"rid"`
	Ts   pulumi.Float64Input   `pulumi:"ts"`
}

func (SqlUserDefinedFunctionGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionGetPropertiesResponseResource)(nil)).Elem()
}

func (i SqlUserDefinedFunctionGetPropertiesResponseResourceArgs) ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutput() SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return i.ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i SqlUserDefinedFunctionGetPropertiesResponseResourceArgs) ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionGetPropertiesResponseResourceOutput)
}

func (i SqlUserDefinedFunctionGetPropertiesResponseResourceArgs) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i SqlUserDefinedFunctionGetPropertiesResponseResourceArgs) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionGetPropertiesResponseResourceOutput).ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type SqlUserDefinedFunctionGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput
	ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput
}

type sqlUserDefinedFunctionGetPropertiesResponseResourcePtrType SqlUserDefinedFunctionGetPropertiesResponseResourceArgs

func SqlUserDefinedFunctionGetPropertiesResponseResourcePtr(v *SqlUserDefinedFunctionGetPropertiesResponseResourceArgs) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrInput {
	return (*sqlUserDefinedFunctionGetPropertiesResponseResourcePtrType)(v)
}

func (*sqlUserDefinedFunctionGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUserDefinedFunctionGetPropertiesResponseResource)(nil)).Elem()
}

func (i *sqlUserDefinedFunctionGetPropertiesResponseResourcePtrType) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return i.ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *sqlUserDefinedFunctionGetPropertiesResponseResourcePtrType) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput)
}

type SqlUserDefinedFunctionGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutput() SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o.ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlUserDefinedFunctionGetPropertiesResponseResource) *SqlUserDefinedFunctionGetPropertiesResponseResource {
		return &v
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v SqlUserDefinedFunctionGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUserDefinedFunctionGetPropertiesResponseResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput() SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) ToSqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Elem() SqlUserDefinedFunctionGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) SqlUserDefinedFunctionGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret SqlUserDefinedFunctionGetPropertiesResponseResource
		return ret
	}).(SqlUserDefinedFunctionGetPropertiesResponseResourceOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type SqlUserDefinedFunctionResource struct {
	Body *string `pulumi:"body"`
	Id   string  `pulumi:"id"`
}





type SqlUserDefinedFunctionResourceInput interface {
	pulumi.Input

	ToSqlUserDefinedFunctionResourceOutput() SqlUserDefinedFunctionResourceOutput
	ToSqlUserDefinedFunctionResourceOutputWithContext(context.Context) SqlUserDefinedFunctionResourceOutput
}

type SqlUserDefinedFunctionResourceArgs struct {
	Body pulumi.StringPtrInput `pulumi:"body"`
	Id   pulumi.StringInput    `pulumi:"id"`
}

func (SqlUserDefinedFunctionResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionResource)(nil)).Elem()
}

func (i SqlUserDefinedFunctionResourceArgs) ToSqlUserDefinedFunctionResourceOutput() SqlUserDefinedFunctionResourceOutput {
	return i.ToSqlUserDefinedFunctionResourceOutputWithContext(context.Background())
}

func (i SqlUserDefinedFunctionResourceArgs) ToSqlUserDefinedFunctionResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionResourceOutput)
}

func (i SqlUserDefinedFunctionResourceArgs) ToSqlUserDefinedFunctionResourcePtrOutput() SqlUserDefinedFunctionResourcePtrOutput {
	return i.ToSqlUserDefinedFunctionResourcePtrOutputWithContext(context.Background())
}

func (i SqlUserDefinedFunctionResourceArgs) ToSqlUserDefinedFunctionResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionResourceOutput).ToSqlUserDefinedFunctionResourcePtrOutputWithContext(ctx)
}









type SqlUserDefinedFunctionResourcePtrInput interface {
	pulumi.Input

	ToSqlUserDefinedFunctionResourcePtrOutput() SqlUserDefinedFunctionResourcePtrOutput
	ToSqlUserDefinedFunctionResourcePtrOutputWithContext(context.Context) SqlUserDefinedFunctionResourcePtrOutput
}

type sqlUserDefinedFunctionResourcePtrType SqlUserDefinedFunctionResourceArgs

func SqlUserDefinedFunctionResourcePtr(v *SqlUserDefinedFunctionResourceArgs) SqlUserDefinedFunctionResourcePtrInput {
	return (*sqlUserDefinedFunctionResourcePtrType)(v)
}

func (*sqlUserDefinedFunctionResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUserDefinedFunctionResource)(nil)).Elem()
}

func (i *sqlUserDefinedFunctionResourcePtrType) ToSqlUserDefinedFunctionResourcePtrOutput() SqlUserDefinedFunctionResourcePtrOutput {
	return i.ToSqlUserDefinedFunctionResourcePtrOutputWithContext(context.Background())
}

func (i *sqlUserDefinedFunctionResourcePtrType) ToSqlUserDefinedFunctionResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlUserDefinedFunctionResourcePtrOutput)
}

type SqlUserDefinedFunctionResourceOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlUserDefinedFunctionResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionResourceOutput) ToSqlUserDefinedFunctionResourceOutput() SqlUserDefinedFunctionResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionResourceOutput) ToSqlUserDefinedFunctionResourceOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourceOutput {
	return o
}

func (o SqlUserDefinedFunctionResourceOutput) ToSqlUserDefinedFunctionResourcePtrOutput() SqlUserDefinedFunctionResourcePtrOutput {
	return o.ToSqlUserDefinedFunctionResourcePtrOutputWithContext(context.Background())
}

func (o SqlUserDefinedFunctionResourceOutput) ToSqlUserDefinedFunctionResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlUserDefinedFunctionResource) *SqlUserDefinedFunctionResource {
		return &v
	}).(SqlUserDefinedFunctionResourcePtrOutput)
}

func (o SqlUserDefinedFunctionResourceOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionResource) *string { return v.Body }).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v SqlUserDefinedFunctionResource) string { return v.Id }).(pulumi.StringOutput)
}

type SqlUserDefinedFunctionResourcePtrOutput struct{ *pulumi.OutputState }

func (SqlUserDefinedFunctionResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlUserDefinedFunctionResource)(nil)).Elem()
}

func (o SqlUserDefinedFunctionResourcePtrOutput) ToSqlUserDefinedFunctionResourcePtrOutput() SqlUserDefinedFunctionResourcePtrOutput {
	return o
}

func (o SqlUserDefinedFunctionResourcePtrOutput) ToSqlUserDefinedFunctionResourcePtrOutputWithContext(ctx context.Context) SqlUserDefinedFunctionResourcePtrOutput {
	return o
}

func (o SqlUserDefinedFunctionResourcePtrOutput) Elem() SqlUserDefinedFunctionResourceOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionResource) SqlUserDefinedFunctionResource {
		if v != nil {
			return *v
		}
		var ret SqlUserDefinedFunctionResource
		return ret
	}).(SqlUserDefinedFunctionResourceOutput)
}

func (o SqlUserDefinedFunctionResourcePtrOutput) Body() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionResource) *string {
		if v == nil {
			return nil
		}
		return v.Body
	}).(pulumi.StringPtrOutput)
}

func (o SqlUserDefinedFunctionResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlUserDefinedFunctionResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
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

type TableGetPropertiesResponseOptions struct {
	AutoscaleSettings *AutoscaleSettingsResponse `pulumi:"autoscaleSettings"`
	Throughput        *int                       `pulumi:"throughput"`
}





type TableGetPropertiesResponseOptionsInput interface {
	pulumi.Input

	ToTableGetPropertiesResponseOptionsOutput() TableGetPropertiesResponseOptionsOutput
	ToTableGetPropertiesResponseOptionsOutputWithContext(context.Context) TableGetPropertiesResponseOptionsOutput
}

type TableGetPropertiesResponseOptionsArgs struct {
	AutoscaleSettings AutoscaleSettingsResponsePtrInput `pulumi:"autoscaleSettings"`
	Throughput        pulumi.IntPtrInput                `pulumi:"throughput"`
}

func (TableGetPropertiesResponseOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGetPropertiesResponseOptions)(nil)).Elem()
}

func (i TableGetPropertiesResponseOptionsArgs) ToTableGetPropertiesResponseOptionsOutput() TableGetPropertiesResponseOptionsOutput {
	return i.ToTableGetPropertiesResponseOptionsOutputWithContext(context.Background())
}

func (i TableGetPropertiesResponseOptionsArgs) ToTableGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGetPropertiesResponseOptionsOutput)
}

func (i TableGetPropertiesResponseOptionsArgs) ToTableGetPropertiesResponseOptionsPtrOutput() TableGetPropertiesResponseOptionsPtrOutput {
	return i.ToTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i TableGetPropertiesResponseOptionsArgs) ToTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGetPropertiesResponseOptionsOutput).ToTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx)
}









type TableGetPropertiesResponseOptionsPtrInput interface {
	pulumi.Input

	ToTableGetPropertiesResponseOptionsPtrOutput() TableGetPropertiesResponseOptionsPtrOutput
	ToTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Context) TableGetPropertiesResponseOptionsPtrOutput
}

type tableGetPropertiesResponseOptionsPtrType TableGetPropertiesResponseOptionsArgs

func TableGetPropertiesResponseOptionsPtr(v *TableGetPropertiesResponseOptionsArgs) TableGetPropertiesResponseOptionsPtrInput {
	return (*tableGetPropertiesResponseOptionsPtrType)(v)
}

func (*tableGetPropertiesResponseOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGetPropertiesResponseOptions)(nil)).Elem()
}

func (i *tableGetPropertiesResponseOptionsPtrType) ToTableGetPropertiesResponseOptionsPtrOutput() TableGetPropertiesResponseOptionsPtrOutput {
	return i.ToTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (i *tableGetPropertiesResponseOptionsPtrType) ToTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGetPropertiesResponseOptionsPtrOutput)
}

type TableGetPropertiesResponseOptionsOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o TableGetPropertiesResponseOptionsOutput) ToTableGetPropertiesResponseOptionsOutput() TableGetPropertiesResponseOptionsOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsOutput) ToTableGetPropertiesResponseOptionsOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsOutput) ToTableGetPropertiesResponseOptionsPtrOutput() TableGetPropertiesResponseOptionsPtrOutput {
	return o.ToTableGetPropertiesResponseOptionsPtrOutputWithContext(context.Background())
}

func (o TableGetPropertiesResponseOptionsOutput) ToTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableGetPropertiesResponseOptions) *TableGetPropertiesResponseOptions {
		return &v
	}).(TableGetPropertiesResponseOptionsPtrOutput)
}

func (o TableGetPropertiesResponseOptionsOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseOptions) *AutoscaleSettingsResponse { return v.AutoscaleSettings }).(AutoscaleSettingsResponsePtrOutput)
}

func (o TableGetPropertiesResponseOptionsOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseOptions) *int { return v.Throughput }).(pulumi.IntPtrOutput)
}

type TableGetPropertiesResponseOptionsPtrOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGetPropertiesResponseOptions)(nil)).Elem()
}

func (o TableGetPropertiesResponseOptionsPtrOutput) ToTableGetPropertiesResponseOptionsPtrOutput() TableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsPtrOutput) ToTableGetPropertiesResponseOptionsPtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseOptionsPtrOutput {
	return o
}

func (o TableGetPropertiesResponseOptionsPtrOutput) Elem() TableGetPropertiesResponseOptionsOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseOptions) TableGetPropertiesResponseOptions {
		if v != nil {
			return *v
		}
		var ret TableGetPropertiesResponseOptions
		return ret
	}).(TableGetPropertiesResponseOptionsOutput)
}

func (o TableGetPropertiesResponseOptionsPtrOutput) AutoscaleSettings() AutoscaleSettingsResponsePtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseOptions) *AutoscaleSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AutoscaleSettings
	}).(AutoscaleSettingsResponsePtrOutput)
}

func (o TableGetPropertiesResponseOptionsPtrOutput) Throughput() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseOptions) *int {
		if v == nil {
			return nil
		}
		return v.Throughput
	}).(pulumi.IntPtrOutput)
}

type TableGetPropertiesResponseResource struct {
	Etag string  `pulumi:"etag"`
	Id   string  `pulumi:"id"`
	Rid  string  `pulumi:"rid"`
	Ts   float64 `pulumi:"ts"`
}





type TableGetPropertiesResponseResourceInput interface {
	pulumi.Input

	ToTableGetPropertiesResponseResourceOutput() TableGetPropertiesResponseResourceOutput
	ToTableGetPropertiesResponseResourceOutputWithContext(context.Context) TableGetPropertiesResponseResourceOutput
}

type TableGetPropertiesResponseResourceArgs struct {
	Etag pulumi.StringInput  `pulumi:"etag"`
	Id   pulumi.StringInput  `pulumi:"id"`
	Rid  pulumi.StringInput  `pulumi:"rid"`
	Ts   pulumi.Float64Input `pulumi:"ts"`
}

func (TableGetPropertiesResponseResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGetPropertiesResponseResource)(nil)).Elem()
}

func (i TableGetPropertiesResponseResourceArgs) ToTableGetPropertiesResponseResourceOutput() TableGetPropertiesResponseResourceOutput {
	return i.ToTableGetPropertiesResponseResourceOutputWithContext(context.Background())
}

func (i TableGetPropertiesResponseResourceArgs) ToTableGetPropertiesResponseResourceOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGetPropertiesResponseResourceOutput)
}

func (i TableGetPropertiesResponseResourceArgs) ToTableGetPropertiesResponseResourcePtrOutput() TableGetPropertiesResponseResourcePtrOutput {
	return i.ToTableGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i TableGetPropertiesResponseResourceArgs) ToTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGetPropertiesResponseResourceOutput).ToTableGetPropertiesResponseResourcePtrOutputWithContext(ctx)
}









type TableGetPropertiesResponseResourcePtrInput interface {
	pulumi.Input

	ToTableGetPropertiesResponseResourcePtrOutput() TableGetPropertiesResponseResourcePtrOutput
	ToTableGetPropertiesResponseResourcePtrOutputWithContext(context.Context) TableGetPropertiesResponseResourcePtrOutput
}

type tableGetPropertiesResponseResourcePtrType TableGetPropertiesResponseResourceArgs

func TableGetPropertiesResponseResourcePtr(v *TableGetPropertiesResponseResourceArgs) TableGetPropertiesResponseResourcePtrInput {
	return (*tableGetPropertiesResponseResourcePtrType)(v)
}

func (*tableGetPropertiesResponseResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGetPropertiesResponseResource)(nil)).Elem()
}

func (i *tableGetPropertiesResponseResourcePtrType) ToTableGetPropertiesResponseResourcePtrOutput() TableGetPropertiesResponseResourcePtrOutput {
	return i.ToTableGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (i *tableGetPropertiesResponseResourcePtrType) ToTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableGetPropertiesResponseResourcePtrOutput)
}

type TableGetPropertiesResponseResourceOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableGetPropertiesResponseResource)(nil)).Elem()
}

func (o TableGetPropertiesResponseResourceOutput) ToTableGetPropertiesResponseResourceOutput() TableGetPropertiesResponseResourceOutput {
	return o
}

func (o TableGetPropertiesResponseResourceOutput) ToTableGetPropertiesResponseResourceOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourceOutput {
	return o
}

func (o TableGetPropertiesResponseResourceOutput) ToTableGetPropertiesResponseResourcePtrOutput() TableGetPropertiesResponseResourcePtrOutput {
	return o.ToTableGetPropertiesResponseResourcePtrOutputWithContext(context.Background())
}

func (o TableGetPropertiesResponseResourceOutput) ToTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableGetPropertiesResponseResource) *TableGetPropertiesResponseResource {
		return &v
	}).(TableGetPropertiesResponseResourcePtrOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Etag() pulumi.StringOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) string { return v.Etag }).(pulumi.StringOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) string { return v.Id }).(pulumi.StringOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Rid() pulumi.StringOutput {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) string { return v.Rid }).(pulumi.StringOutput)
}

func (o TableGetPropertiesResponseResourceOutput) Ts() pulumi.Float64Output {
	return o.ApplyT(func(v TableGetPropertiesResponseResource) float64 { return v.Ts }).(pulumi.Float64Output)
}

type TableGetPropertiesResponseResourcePtrOutput struct{ *pulumi.OutputState }

func (TableGetPropertiesResponseResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableGetPropertiesResponseResource)(nil)).Elem()
}

func (o TableGetPropertiesResponseResourcePtrOutput) ToTableGetPropertiesResponseResourcePtrOutput() TableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o TableGetPropertiesResponseResourcePtrOutput) ToTableGetPropertiesResponseResourcePtrOutputWithContext(ctx context.Context) TableGetPropertiesResponseResourcePtrOutput {
	return o
}

func (o TableGetPropertiesResponseResourcePtrOutput) Elem() TableGetPropertiesResponseResourceOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) TableGetPropertiesResponseResource {
		if v != nil {
			return *v
		}
		var ret TableGetPropertiesResponseResource
		return ret
	}).(TableGetPropertiesResponseResourceOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Etag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Etag
	}).(pulumi.StringPtrOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Rid() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *string {
		if v == nil {
			return nil
		}
		return &v.Rid
	}).(pulumi.StringPtrOutput)
}

func (o TableGetPropertiesResponseResourcePtrOutput) Ts() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *TableGetPropertiesResponseResource) *float64 {
		if v == nil {
			return nil
		}
		return &v.Ts
	}).(pulumi.Float64PtrOutput)
}

type TableResource struct {
	Id string `pulumi:"id"`
}





type TableResourceInput interface {
	pulumi.Input

	ToTableResourceOutput() TableResourceOutput
	ToTableResourceOutputWithContext(context.Context) TableResourceOutput
}

type TableResourceArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (TableResourceArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TableResource)(nil)).Elem()
}

func (i TableResourceArgs) ToTableResourceOutput() TableResourceOutput {
	return i.ToTableResourceOutputWithContext(context.Background())
}

func (i TableResourceArgs) ToTableResourceOutputWithContext(ctx context.Context) TableResourceOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourceOutput)
}

func (i TableResourceArgs) ToTableResourcePtrOutput() TableResourcePtrOutput {
	return i.ToTableResourcePtrOutputWithContext(context.Background())
}

func (i TableResourceArgs) ToTableResourcePtrOutputWithContext(ctx context.Context) TableResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourceOutput).ToTableResourcePtrOutputWithContext(ctx)
}









type TableResourcePtrInput interface {
	pulumi.Input

	ToTableResourcePtrOutput() TableResourcePtrOutput
	ToTableResourcePtrOutputWithContext(context.Context) TableResourcePtrOutput
}

type tableResourcePtrType TableResourceArgs

func TableResourcePtr(v *TableResourceArgs) TableResourcePtrInput {
	return (*tableResourcePtrType)(v)
}

func (*tableResourcePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TableResource)(nil)).Elem()
}

func (i *tableResourcePtrType) ToTableResourcePtrOutput() TableResourcePtrOutput {
	return i.ToTableResourcePtrOutputWithContext(context.Background())
}

func (i *tableResourcePtrType) ToTableResourcePtrOutputWithContext(ctx context.Context) TableResourcePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TableResourcePtrOutput)
}

type TableResourceOutput struct{ *pulumi.OutputState }

func (TableResourceOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TableResource)(nil)).Elem()
}

func (o TableResourceOutput) ToTableResourceOutput() TableResourceOutput {
	return o
}

func (o TableResourceOutput) ToTableResourceOutputWithContext(ctx context.Context) TableResourceOutput {
	return o
}

func (o TableResourceOutput) ToTableResourcePtrOutput() TableResourcePtrOutput {
	return o.ToTableResourcePtrOutputWithContext(context.Background())
}

func (o TableResourceOutput) ToTableResourcePtrOutputWithContext(ctx context.Context) TableResourcePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TableResource) *TableResource {
		return &v
	}).(TableResourcePtrOutput)
}

func (o TableResourceOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v TableResource) string { return v.Id }).(pulumi.StringOutput)
}

type TableResourcePtrOutput struct{ *pulumi.OutputState }

func (TableResourcePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TableResource)(nil)).Elem()
}

func (o TableResourcePtrOutput) ToTableResourcePtrOutput() TableResourcePtrOutput {
	return o
}

func (o TableResourcePtrOutput) ToTableResourcePtrOutputWithContext(ctx context.Context) TableResourcePtrOutput {
	return o
}

func (o TableResourcePtrOutput) Elem() TableResourceOutput {
	return o.ApplyT(func(v *TableResource) TableResource {
		if v != nil {
			return *v
		}
		var ret TableResource
		return ret
	}).(TableResourceOutput)
}

func (o TableResourcePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TableResource) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type UniqueKey struct {
	Paths []string `pulumi:"paths"`
}





type UniqueKeyInput interface {
	pulumi.Input

	ToUniqueKeyOutput() UniqueKeyOutput
	ToUniqueKeyOutputWithContext(context.Context) UniqueKeyOutput
}

type UniqueKeyArgs struct {
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (UniqueKeyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKey)(nil)).Elem()
}

func (i UniqueKeyArgs) ToUniqueKeyOutput() UniqueKeyOutput {
	return i.ToUniqueKeyOutputWithContext(context.Background())
}

func (i UniqueKeyArgs) ToUniqueKeyOutputWithContext(ctx context.Context) UniqueKeyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyOutput)
}





type UniqueKeyArrayInput interface {
	pulumi.Input

	ToUniqueKeyArrayOutput() UniqueKeyArrayOutput
	ToUniqueKeyArrayOutputWithContext(context.Context) UniqueKeyArrayOutput
}

type UniqueKeyArray []UniqueKeyInput

func (UniqueKeyArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKey)(nil)).Elem()
}

func (i UniqueKeyArray) ToUniqueKeyArrayOutput() UniqueKeyArrayOutput {
	return i.ToUniqueKeyArrayOutputWithContext(context.Background())
}

func (i UniqueKeyArray) ToUniqueKeyArrayOutputWithContext(ctx context.Context) UniqueKeyArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyArrayOutput)
}

type UniqueKeyOutput struct{ *pulumi.OutputState }

func (UniqueKeyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKey)(nil)).Elem()
}

func (o UniqueKeyOutput) ToUniqueKeyOutput() UniqueKeyOutput {
	return o
}

func (o UniqueKeyOutput) ToUniqueKeyOutputWithContext(ctx context.Context) UniqueKeyOutput {
	return o
}

func (o UniqueKeyOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UniqueKey) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type UniqueKeyArrayOutput struct{ *pulumi.OutputState }

func (UniqueKeyArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKey)(nil)).Elem()
}

func (o UniqueKeyArrayOutput) ToUniqueKeyArrayOutput() UniqueKeyArrayOutput {
	return o
}

func (o UniqueKeyArrayOutput) ToUniqueKeyArrayOutputWithContext(ctx context.Context) UniqueKeyArrayOutput {
	return o
}

func (o UniqueKeyArrayOutput) Index(i pulumi.IntInput) UniqueKeyOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UniqueKey {
		return vs[0].([]UniqueKey)[vs[1].(int)]
	}).(UniqueKeyOutput)
}

type UniqueKeyPolicy struct {
	UniqueKeys []UniqueKey `pulumi:"uniqueKeys"`
}





type UniqueKeyPolicyInput interface {
	pulumi.Input

	ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput
	ToUniqueKeyPolicyOutputWithContext(context.Context) UniqueKeyPolicyOutput
}

type UniqueKeyPolicyArgs struct {
	UniqueKeys UniqueKeyArrayInput `pulumi:"uniqueKeys"`
}

func (UniqueKeyPolicyArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicy)(nil)).Elem()
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput {
	return i.ToUniqueKeyPolicyOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyOutputWithContext(ctx context.Context) UniqueKeyPolicyOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyOutput)
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return i.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyArgs) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyOutput).ToUniqueKeyPolicyPtrOutputWithContext(ctx)
}









type UniqueKeyPolicyPtrInput interface {
	pulumi.Input

	ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput
	ToUniqueKeyPolicyPtrOutputWithContext(context.Context) UniqueKeyPolicyPtrOutput
}

type uniqueKeyPolicyPtrType UniqueKeyPolicyArgs

func UniqueKeyPolicyPtr(v *UniqueKeyPolicyArgs) UniqueKeyPolicyPtrInput {
	return (*uniqueKeyPolicyPtrType)(v)
}

func (*uniqueKeyPolicyPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicy)(nil)).Elem()
}

func (i *uniqueKeyPolicyPtrType) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return i.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (i *uniqueKeyPolicyPtrType) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyPtrOutput)
}

type UniqueKeyPolicyOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicy)(nil)).Elem()
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyOutput() UniqueKeyPolicyOutput {
	return o
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyOutputWithContext(ctx context.Context) UniqueKeyPolicyOutput {
	return o
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return o.ToUniqueKeyPolicyPtrOutputWithContext(context.Background())
}

func (o UniqueKeyPolicyOutput) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UniqueKeyPolicy) *UniqueKeyPolicy {
		return &v
	}).(UniqueKeyPolicyPtrOutput)
}

func (o UniqueKeyPolicyOutput) UniqueKeys() UniqueKeyArrayOutput {
	return o.ApplyT(func(v UniqueKeyPolicy) []UniqueKey { return v.UniqueKeys }).(UniqueKeyArrayOutput)
}

type UniqueKeyPolicyPtrOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicy)(nil)).Elem()
}

func (o UniqueKeyPolicyPtrOutput) ToUniqueKeyPolicyPtrOutput() UniqueKeyPolicyPtrOutput {
	return o
}

func (o UniqueKeyPolicyPtrOutput) ToUniqueKeyPolicyPtrOutputWithContext(ctx context.Context) UniqueKeyPolicyPtrOutput {
	return o
}

func (o UniqueKeyPolicyPtrOutput) Elem() UniqueKeyPolicyOutput {
	return o.ApplyT(func(v *UniqueKeyPolicy) UniqueKeyPolicy {
		if v != nil {
			return *v
		}
		var ret UniqueKeyPolicy
		return ret
	}).(UniqueKeyPolicyOutput)
}

func (o UniqueKeyPolicyPtrOutput) UniqueKeys() UniqueKeyArrayOutput {
	return o.ApplyT(func(v *UniqueKeyPolicy) []UniqueKey {
		if v == nil {
			return nil
		}
		return v.UniqueKeys
	}).(UniqueKeyArrayOutput)
}

type UniqueKeyPolicyResponse struct {
	UniqueKeys []UniqueKeyResponse `pulumi:"uniqueKeys"`
}





type UniqueKeyPolicyResponseInput interface {
	pulumi.Input

	ToUniqueKeyPolicyResponseOutput() UniqueKeyPolicyResponseOutput
	ToUniqueKeyPolicyResponseOutputWithContext(context.Context) UniqueKeyPolicyResponseOutput
}

type UniqueKeyPolicyResponseArgs struct {
	UniqueKeys UniqueKeyResponseArrayInput `pulumi:"uniqueKeys"`
}

func (UniqueKeyPolicyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicyResponse)(nil)).Elem()
}

func (i UniqueKeyPolicyResponseArgs) ToUniqueKeyPolicyResponseOutput() UniqueKeyPolicyResponseOutput {
	return i.ToUniqueKeyPolicyResponseOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyResponseArgs) ToUniqueKeyPolicyResponseOutputWithContext(ctx context.Context) UniqueKeyPolicyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyResponseOutput)
}

func (i UniqueKeyPolicyResponseArgs) ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput {
	return i.ToUniqueKeyPolicyResponsePtrOutputWithContext(context.Background())
}

func (i UniqueKeyPolicyResponseArgs) ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx context.Context) UniqueKeyPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyResponseOutput).ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx)
}









type UniqueKeyPolicyResponsePtrInput interface {
	pulumi.Input

	ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput
	ToUniqueKeyPolicyResponsePtrOutputWithContext(context.Context) UniqueKeyPolicyResponsePtrOutput
}

type uniqueKeyPolicyResponsePtrType UniqueKeyPolicyResponseArgs

func UniqueKeyPolicyResponsePtr(v *UniqueKeyPolicyResponseArgs) UniqueKeyPolicyResponsePtrInput {
	return (*uniqueKeyPolicyResponsePtrType)(v)
}

func (*uniqueKeyPolicyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicyResponse)(nil)).Elem()
}

func (i *uniqueKeyPolicyResponsePtrType) ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput {
	return i.ToUniqueKeyPolicyResponsePtrOutputWithContext(context.Background())
}

func (i *uniqueKeyPolicyResponsePtrType) ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx context.Context) UniqueKeyPolicyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyPolicyResponsePtrOutput)
}

type UniqueKeyPolicyResponseOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyPolicyResponse)(nil)).Elem()
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponseOutput() UniqueKeyPolicyResponseOutput {
	return o
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponseOutputWithContext(ctx context.Context) UniqueKeyPolicyResponseOutput {
	return o
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput {
	return o.ToUniqueKeyPolicyResponsePtrOutputWithContext(context.Background())
}

func (o UniqueKeyPolicyResponseOutput) ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx context.Context) UniqueKeyPolicyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v UniqueKeyPolicyResponse) *UniqueKeyPolicyResponse {
		return &v
	}).(UniqueKeyPolicyResponsePtrOutput)
}

func (o UniqueKeyPolicyResponseOutput) UniqueKeys() UniqueKeyResponseArrayOutput {
	return o.ApplyT(func(v UniqueKeyPolicyResponse) []UniqueKeyResponse { return v.UniqueKeys }).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyPolicyResponsePtrOutput struct{ *pulumi.OutputState }

func (UniqueKeyPolicyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UniqueKeyPolicyResponse)(nil)).Elem()
}

func (o UniqueKeyPolicyResponsePtrOutput) ToUniqueKeyPolicyResponsePtrOutput() UniqueKeyPolicyResponsePtrOutput {
	return o
}

func (o UniqueKeyPolicyResponsePtrOutput) ToUniqueKeyPolicyResponsePtrOutputWithContext(ctx context.Context) UniqueKeyPolicyResponsePtrOutput {
	return o
}

func (o UniqueKeyPolicyResponsePtrOutput) Elem() UniqueKeyPolicyResponseOutput {
	return o.ApplyT(func(v *UniqueKeyPolicyResponse) UniqueKeyPolicyResponse {
		if v != nil {
			return *v
		}
		var ret UniqueKeyPolicyResponse
		return ret
	}).(UniqueKeyPolicyResponseOutput)
}

func (o UniqueKeyPolicyResponsePtrOutput) UniqueKeys() UniqueKeyResponseArrayOutput {
	return o.ApplyT(func(v *UniqueKeyPolicyResponse) []UniqueKeyResponse {
		if v == nil {
			return nil
		}
		return v.UniqueKeys
	}).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyResponse struct {
	Paths []string `pulumi:"paths"`
}





type UniqueKeyResponseInput interface {
	pulumi.Input

	ToUniqueKeyResponseOutput() UniqueKeyResponseOutput
	ToUniqueKeyResponseOutputWithContext(context.Context) UniqueKeyResponseOutput
}

type UniqueKeyResponseArgs struct {
	Paths pulumi.StringArrayInput `pulumi:"paths"`
}

func (UniqueKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyResponse)(nil)).Elem()
}

func (i UniqueKeyResponseArgs) ToUniqueKeyResponseOutput() UniqueKeyResponseOutput {
	return i.ToUniqueKeyResponseOutputWithContext(context.Background())
}

func (i UniqueKeyResponseArgs) ToUniqueKeyResponseOutputWithContext(ctx context.Context) UniqueKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyResponseOutput)
}





type UniqueKeyResponseArrayInput interface {
	pulumi.Input

	ToUniqueKeyResponseArrayOutput() UniqueKeyResponseArrayOutput
	ToUniqueKeyResponseArrayOutputWithContext(context.Context) UniqueKeyResponseArrayOutput
}

type UniqueKeyResponseArray []UniqueKeyResponseInput

func (UniqueKeyResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKeyResponse)(nil)).Elem()
}

func (i UniqueKeyResponseArray) ToUniqueKeyResponseArrayOutput() UniqueKeyResponseArrayOutput {
	return i.ToUniqueKeyResponseArrayOutputWithContext(context.Background())
}

func (i UniqueKeyResponseArray) ToUniqueKeyResponseArrayOutputWithContext(ctx context.Context) UniqueKeyResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(UniqueKeyResponseArrayOutput)
}

type UniqueKeyResponseOutput struct{ *pulumi.OutputState }

func (UniqueKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UniqueKeyResponse)(nil)).Elem()
}

func (o UniqueKeyResponseOutput) ToUniqueKeyResponseOutput() UniqueKeyResponseOutput {
	return o
}

func (o UniqueKeyResponseOutput) ToUniqueKeyResponseOutputWithContext(ctx context.Context) UniqueKeyResponseOutput {
	return o
}

func (o UniqueKeyResponseOutput) Paths() pulumi.StringArrayOutput {
	return o.ApplyT(func(v UniqueKeyResponse) []string { return v.Paths }).(pulumi.StringArrayOutput)
}

type UniqueKeyResponseArrayOutput struct{ *pulumi.OutputState }

func (UniqueKeyResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]UniqueKeyResponse)(nil)).Elem()
}

func (o UniqueKeyResponseArrayOutput) ToUniqueKeyResponseArrayOutput() UniqueKeyResponseArrayOutput {
	return o
}

func (o UniqueKeyResponseArrayOutput) ToUniqueKeyResponseArrayOutputWithContext(ctx context.Context) UniqueKeyResponseArrayOutput {
	return o
}

func (o UniqueKeyResponseArrayOutput) Index(i pulumi.IntInput) UniqueKeyResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) UniqueKeyResponse {
		return vs[0].([]UniqueKeyResponse)[vs[1].(int)]
	}).(UniqueKeyResponseOutput)
}

type VirtualNetworkRule struct {
	Id                               *string `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint *bool   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}





type VirtualNetworkRuleInput interface {
	pulumi.Input

	ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput
	ToVirtualNetworkRuleOutputWithContext(context.Context) VirtualNetworkRuleOutput
}

type VirtualNetworkRuleArgs struct {
	Id                               pulumi.StringPtrInput `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}

func (VirtualNetworkRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return i.ToVirtualNetworkRuleOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArgs) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleOutput)
}





type VirtualNetworkRuleArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput
	ToVirtualNetworkRuleArrayOutputWithContext(context.Context) VirtualNetworkRuleArrayOutput
}

type VirtualNetworkRuleArray []VirtualNetworkRuleInput

func (VirtualNetworkRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return i.ToVirtualNetworkRuleArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleArray) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleArrayOutput)
}

type VirtualNetworkRuleOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutput() VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) ToVirtualNetworkRuleOutputWithContext(ctx context.Context) VirtualNetworkRuleOutput {
	return o
}

func (o VirtualNetworkRuleOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleOutput) IgnoreMissingVNetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRule) *bool { return v.IgnoreMissingVNetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkRuleArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRule)(nil)).Elem()
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutput() VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) ToVirtualNetworkRuleArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleArrayOutput {
	return o
}

func (o VirtualNetworkRuleArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRule {
		return vs[0].([]VirtualNetworkRule)[vs[1].(int)]
	}).(VirtualNetworkRuleOutput)
}

type VirtualNetworkRuleResponse struct {
	Id                               *string `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint *bool   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}





type VirtualNetworkRuleResponseInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput
	ToVirtualNetworkRuleResponseOutputWithContext(context.Context) VirtualNetworkRuleResponseOutput
}

type VirtualNetworkRuleResponseArgs struct {
	Id                               pulumi.StringPtrInput `pulumi:"id"`
	IgnoreMissingVNetServiceEndpoint pulumi.BoolPtrInput   `pulumi:"ignoreMissingVNetServiceEndpoint"`
}

func (VirtualNetworkRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return i.ToVirtualNetworkRuleResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArgs) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseOutput)
}





type VirtualNetworkRuleResponseArrayInput interface {
	pulumi.Input

	ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput
	ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Context) VirtualNetworkRuleResponseArrayOutput
}

type VirtualNetworkRuleResponseArray []VirtualNetworkRuleResponseInput

func (VirtualNetworkRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return i.ToVirtualNetworkRuleResponseArrayOutputWithContext(context.Background())
}

func (i VirtualNetworkRuleResponseArray) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkRuleResponseArrayOutput)
}

type VirtualNetworkRuleResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutput() VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) ToVirtualNetworkRuleResponseOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseOutput {
	return o
}

func (o VirtualNetworkRuleResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkRuleResponseOutput) IgnoreMissingVNetServiceEndpoint() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualNetworkRuleResponse) *bool { return v.IgnoreMissingVNetServiceEndpoint }).(pulumi.BoolPtrOutput)
}

type VirtualNetworkRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualNetworkRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualNetworkRuleResponse)(nil)).Elem()
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutput() VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) ToVirtualNetworkRuleResponseArrayOutputWithContext(ctx context.Context) VirtualNetworkRuleResponseArrayOutput {
	return o
}

func (o VirtualNetworkRuleResponseArrayOutput) Index(i pulumi.IntInput) VirtualNetworkRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualNetworkRuleResponse {
		return vs[0].([]VirtualNetworkRuleResponse)[vs[1].(int)]
	}).(VirtualNetworkRuleResponseOutput)
}

type CompositePathArrayArray []CompositePathArrayInput

func (CompositePathArrayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathArray)(nil)).Elem()
}

func (i CompositePathArrayArray) ToCompositePathArrayArrayOutput() CompositePathArrayArrayOutput {
	return i.ToCompositePathArrayArrayOutputWithContext(context.Background())
}

func (i CompositePathArrayArray) ToCompositePathArrayArrayOutputWithContext(ctx context.Context) CompositePathArrayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathArrayArrayOutput)
}

type CompositePathArrayArrayOutput struct{ *pulumi.OutputState }

func (CompositePathArrayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePathArray)(nil)).Elem()
}

func (o CompositePathArrayArrayOutput) ToCompositePathArrayArrayOutput() CompositePathArrayArrayOutput {
	return o
}

func (o CompositePathArrayArrayOutput) ToCompositePathArrayArrayOutputWithContext(ctx context.Context) CompositePathArrayArrayOutput {
	return o
}

func (o CompositePathArrayArrayOutput) Index(i pulumi.IntInput) CompositePathArrayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CompositePathArray {
		return vs[0].([]CompositePathArray)[vs[1].(int)]
	}).(CompositePathArrayOutput)
}





type CompositePathArrayArrayInput interface {
	pulumi.Input

	ToCompositePathArrayArrayOutput() CompositePathArrayArrayOutput
	ToCompositePathArrayArrayOutputWithContext(context.Context) CompositePathArrayArrayOutput
}

type CompositePathResponseArrayArray []CompositePathResponseArrayInput

func (CompositePathResponseArrayArray) ElementType() reflect.Type {
	return reflect.TypeOf((*CompositePathResponseArray)(nil)).Elem()
}

func (i CompositePathResponseArrayArray) ToCompositePathResponseArrayArrayOutput() CompositePathResponseArrayArrayOutput {
	return i.ToCompositePathResponseArrayArrayOutputWithContext(context.Background())
}

func (i CompositePathResponseArrayArray) ToCompositePathResponseArrayArrayOutputWithContext(ctx context.Context) CompositePathResponseArrayArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CompositePathResponseArrayArrayOutput)
}

type CompositePathResponseArrayArrayOutput struct{ *pulumi.OutputState }

func (CompositePathResponseArrayArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]CompositePathResponseArray)(nil)).Elem()
}

func (o CompositePathResponseArrayArrayOutput) ToCompositePathResponseArrayArrayOutput() CompositePathResponseArrayArrayOutput {
	return o
}

func (o CompositePathResponseArrayArrayOutput) ToCompositePathResponseArrayArrayOutputWithContext(ctx context.Context) CompositePathResponseArrayArrayOutput {
	return o
}

func (o CompositePathResponseArrayArrayOutput) Index(i pulumi.IntInput) CompositePathResponseArrayOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) CompositePathResponseArray {
		return vs[0].([]CompositePathResponseArray)[vs[1].(int)]
	}).(CompositePathResponseArrayOutput)
}





type CompositePathResponseArrayArrayInput interface {
	pulumi.Input

	ToCompositePathResponseArrayArrayOutput() CompositePathResponseArrayArrayOutput
	ToCompositePathResponseArrayArrayOutputWithContext(context.Context) CompositePathResponseArrayArrayOutput
}

func init() {
	pulumi.RegisterOutputType(AnalyticalStorageConfigurationOutput{})
	pulumi.RegisterOutputType(AnalyticalStorageConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AnalyticalStorageConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AnalyticalStorageConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiPropertiesOutput{})
	pulumi.RegisterOutputType(ApiPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ApiPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsPtrOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsResponseOutput{})
	pulumi.RegisterOutputType(AutoscaleSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(BackupPolicyMigrationStateOutput{})
	pulumi.RegisterOutputType(BackupPolicyMigrationStatePtrOutput{})
	pulumi.RegisterOutputType(BackupPolicyMigrationStateResponseOutput{})
	pulumi.RegisterOutputType(BackupPolicyMigrationStateResponsePtrOutput{})
	pulumi.RegisterOutputType(CapabilityOutput{})
	pulumi.RegisterOutputType(CapabilityArrayOutput{})
	pulumi.RegisterOutputType(CapabilityResponseOutput{})
	pulumi.RegisterOutputType(CapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(CapacityOutput{})
	pulumi.RegisterOutputType(CapacityPtrOutput{})
	pulumi.RegisterOutputType(CapacityResponseOutput{})
	pulumi.RegisterOutputType(CapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceResourceOutput{})
	pulumi.RegisterOutputType(CassandraKeyspaceResourcePtrOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyArrayOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyResponseOutput{})
	pulumi.RegisterOutputType(CassandraPartitionKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(CassandraSchemaOutput{})
	pulumi.RegisterOutputType(CassandraSchemaPtrOutput{})
	pulumi.RegisterOutputType(CassandraSchemaResponseOutput{})
	pulumi.RegisterOutputType(CassandraSchemaResponsePtrOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(CassandraTableGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(CassandraTableResourceOutput{})
	pulumi.RegisterOutputType(CassandraTableResourcePtrOutput{})
	pulumi.RegisterOutputType(CertificateOutput{})
	pulumi.RegisterOutputType(CertificateArrayOutput{})
	pulumi.RegisterOutputType(CertificateResponseOutput{})
	pulumi.RegisterOutputType(CertificateResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterKeyOutput{})
	pulumi.RegisterOutputType(ClusterKeyArrayOutput{})
	pulumi.RegisterOutputType(ClusterKeyResponseOutput{})
	pulumi.RegisterOutputType(ClusterKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesOutput{})
	pulumi.RegisterOutputType(ClusterResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ClusterResourceResponsePropertiesOutput{})
	pulumi.RegisterOutputType(ClusterResourceResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(ColumnOutput{})
	pulumi.RegisterOutputType(ColumnArrayOutput{})
	pulumi.RegisterOutputType(ColumnResponseOutput{})
	pulumi.RegisterOutputType(ColumnResponseArrayOutput{})
	pulumi.RegisterOutputType(CompositePathOutput{})
	pulumi.RegisterOutputType(CompositePathArrayOutput{})
	pulumi.RegisterOutputType(CompositePathResponseOutput{})
	pulumi.RegisterOutputType(CompositePathResponseArrayOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyResponseOutput{})
	pulumi.RegisterOutputType(ConflictResolutionPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyPtrOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyResponseOutput{})
	pulumi.RegisterOutputType(ConsistencyPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyPtrOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyResponseOutput{})
	pulumi.RegisterOutputType(ContainerPartitionKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(ContinuousModeBackupPolicyOutput{})
	pulumi.RegisterOutputType(ContinuousModeBackupPolicyResponseOutput{})
	pulumi.RegisterOutputType(CorsPolicyOutput{})
	pulumi.RegisterOutputType(CorsPolicyArrayOutput{})
	pulumi.RegisterOutputType(CorsPolicyResponseOutput{})
	pulumi.RegisterOutputType(CorsPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(CreateUpdateOptionsOutput{})
	pulumi.RegisterOutputType(CreateUpdateOptionsPtrOutput{})
	pulumi.RegisterOutputType(DataCenterResourcePropertiesOutput{})
	pulumi.RegisterOutputType(DataCenterResourcePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DataCenterResourceResponsePropertiesOutput{})
	pulumi.RegisterOutputType(DataCenterResourceResponsePropertiesPtrOutput{})
	pulumi.RegisterOutputType(DatabaseAccountConnectionStringResponseOutput{})
	pulumi.RegisterOutputType(DatabaseAccountConnectionStringResponseArrayOutput{})
	pulumi.RegisterOutputType(DatabaseRestoreResourceOutput{})
	pulumi.RegisterOutputType(DatabaseRestoreResourceArrayOutput{})
	pulumi.RegisterOutputType(DatabaseRestoreResourceResponseOutput{})
	pulumi.RegisterOutputType(DatabaseRestoreResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(ExcludedPathOutput{})
	pulumi.RegisterOutputType(ExcludedPathArrayOutput{})
	pulumi.RegisterOutputType(ExcludedPathResponseOutput{})
	pulumi.RegisterOutputType(ExcludedPathResponseArrayOutput{})
	pulumi.RegisterOutputType(FailoverPolicyResponseOutput{})
	pulumi.RegisterOutputType(FailoverPolicyResponseArrayOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseResourceOutput{})
	pulumi.RegisterOutputType(GremlinDatabaseResourcePtrOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(GremlinGraphGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(GremlinGraphResourceOutput{})
	pulumi.RegisterOutputType(GremlinGraphResourcePtrOutput{})
	pulumi.RegisterOutputType(IncludedPathOutput{})
	pulumi.RegisterOutputType(IncludedPathArrayOutput{})
	pulumi.RegisterOutputType(IncludedPathResponseOutput{})
	pulumi.RegisterOutputType(IncludedPathResponseArrayOutput{})
	pulumi.RegisterOutputType(IndexesOutput{})
	pulumi.RegisterOutputType(IndexesArrayOutput{})
	pulumi.RegisterOutputType(IndexesResponseOutput{})
	pulumi.RegisterOutputType(IndexesResponseArrayOutput{})
	pulumi.RegisterOutputType(IndexingPolicyOutput{})
	pulumi.RegisterOutputType(IndexingPolicyPtrOutput{})
	pulumi.RegisterOutputType(IndexingPolicyResponseOutput{})
	pulumi.RegisterOutputType(IndexingPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeArrayOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeResponseOutput{})
	pulumi.RegisterOutputType(IpAddressOrRangeResponseArrayOutput{})
	pulumi.RegisterOutputType(LocationOutput{})
	pulumi.RegisterOutputType(LocationArrayOutput{})
	pulumi.RegisterOutputType(LocationResponseOutput{})
	pulumi.RegisterOutputType(LocationResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedCassandraManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedCassandraManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedCassandraManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedCassandraManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionResourceOutput{})
	pulumi.RegisterOutputType(MongoDBCollectionResourcePtrOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseResourceOutput{})
	pulumi.RegisterOutputType(MongoDBDatabaseResourcePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOutput{})
	pulumi.RegisterOutputType(MongoIndexArrayOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysPtrOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexKeysResponsePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsPtrOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(MongoIndexResponseOutput{})
	pulumi.RegisterOutputType(MongoIndexResponseArrayOutput{})
	pulumi.RegisterOutputType(PeriodicModeBackupPolicyOutput{})
	pulumi.RegisterOutputType(PeriodicModeBackupPolicyResponseOutput{})
	pulumi.RegisterOutputType(PeriodicModePropertiesOutput{})
	pulumi.RegisterOutputType(PeriodicModePropertiesPtrOutput{})
	pulumi.RegisterOutputType(PeriodicModePropertiesResponseOutput{})
	pulumi.RegisterOutputType(PeriodicModePropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(PermissionOutput{})
	pulumi.RegisterOutputType(PermissionArrayOutput{})
	pulumi.RegisterOutputType(PermissionResponseOutput{})
	pulumi.RegisterOutputType(PermissionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointConnectionResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointPropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyPtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStatePropertyResponsePtrOutput{})
	pulumi.RegisterOutputType(RestoreParametersOutput{})
	pulumi.RegisterOutputType(RestoreParametersPtrOutput{})
	pulumi.RegisterOutputType(RestoreParametersResponseOutput{})
	pulumi.RegisterOutputType(RestoreParametersResponsePtrOutput{})
	pulumi.RegisterOutputType(SeedNodeOutput{})
	pulumi.RegisterOutputType(SeedNodePtrOutput{})
	pulumi.RegisterOutputType(SeedNodeArrayOutput{})
	pulumi.RegisterOutputType(SeedNodeResponseOutput{})
	pulumi.RegisterOutputType(SeedNodeResponsePtrOutput{})
	pulumi.RegisterOutputType(SeedNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(SpatialSpecOutput{})
	pulumi.RegisterOutputType(SpatialSpecArrayOutput{})
	pulumi.RegisterOutputType(SpatialSpecResponseOutput{})
	pulumi.RegisterOutputType(SpatialSpecResponseArrayOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlContainerGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlContainerResourceOutput{})
	pulumi.RegisterOutputType(SqlContainerResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlDatabaseGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourceOutput{})
	pulumi.RegisterOutputType(SqlDatabaseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureResourceOutput{})
	pulumi.RegisterOutputType(SqlStoredProcedureResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlTriggerGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlTriggerGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlTriggerResourceOutput{})
	pulumi.RegisterOutputType(SqlTriggerResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionResourceOutput{})
	pulumi.RegisterOutputType(SqlUserDefinedFunctionResourcePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(SystemDataResponsePtrOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseOptionsOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseOptionsPtrOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseResourceOutput{})
	pulumi.RegisterOutputType(TableGetPropertiesResponseResourcePtrOutput{})
	pulumi.RegisterOutputType(TableResourceOutput{})
	pulumi.RegisterOutputType(TableResourcePtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyOutput{})
	pulumi.RegisterOutputType(UniqueKeyArrayOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyPtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyResponseOutput{})
	pulumi.RegisterOutputType(UniqueKeyPolicyResponsePtrOutput{})
	pulumi.RegisterOutputType(UniqueKeyResponseOutput{})
	pulumi.RegisterOutputType(UniqueKeyResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(CompositePathArrayArrayOutput{})
	pulumi.RegisterOutputType(CompositePathResponseArrayArrayOutput{})
}
