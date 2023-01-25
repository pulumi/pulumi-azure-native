


package v20220801preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AADAuthenticationSettings struct {
	ClientId *string `pulumi:"clientId"`
}





type AADAuthenticationSettingsInput interface {
	pulumi.Input

	ToAADAuthenticationSettingsOutput() AADAuthenticationSettingsOutput
	ToAADAuthenticationSettingsOutputWithContext(context.Context) AADAuthenticationSettingsOutput
}

type AADAuthenticationSettingsArgs struct {
	ClientId pulumi.StringPtrInput `pulumi:"clientId"`
}

func (AADAuthenticationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AADAuthenticationSettings)(nil)).Elem()
}

func (i AADAuthenticationSettingsArgs) ToAADAuthenticationSettingsOutput() AADAuthenticationSettingsOutput {
	return i.ToAADAuthenticationSettingsOutputWithContext(context.Background())
}

func (i AADAuthenticationSettingsArgs) ToAADAuthenticationSettingsOutputWithContext(ctx context.Context) AADAuthenticationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAuthenticationSettingsOutput)
}

func (i AADAuthenticationSettingsArgs) ToAADAuthenticationSettingsPtrOutput() AADAuthenticationSettingsPtrOutput {
	return i.ToAADAuthenticationSettingsPtrOutputWithContext(context.Background())
}

func (i AADAuthenticationSettingsArgs) ToAADAuthenticationSettingsPtrOutputWithContext(ctx context.Context) AADAuthenticationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAuthenticationSettingsOutput).ToAADAuthenticationSettingsPtrOutputWithContext(ctx)
}









type AADAuthenticationSettingsPtrInput interface {
	pulumi.Input

	ToAADAuthenticationSettingsPtrOutput() AADAuthenticationSettingsPtrOutput
	ToAADAuthenticationSettingsPtrOutputWithContext(context.Context) AADAuthenticationSettingsPtrOutput
}

type aadauthenticationSettingsPtrType AADAuthenticationSettingsArgs

func AADAuthenticationSettingsPtr(v *AADAuthenticationSettingsArgs) AADAuthenticationSettingsPtrInput {
	return (*aadauthenticationSettingsPtrType)(v)
}

func (*aadauthenticationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AADAuthenticationSettings)(nil)).Elem()
}

func (i *aadauthenticationSettingsPtrType) ToAADAuthenticationSettingsPtrOutput() AADAuthenticationSettingsPtrOutput {
	return i.ToAADAuthenticationSettingsPtrOutputWithContext(context.Background())
}

func (i *aadauthenticationSettingsPtrType) ToAADAuthenticationSettingsPtrOutputWithContext(ctx context.Context) AADAuthenticationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AADAuthenticationSettingsPtrOutput)
}

type AADAuthenticationSettingsOutput struct{ *pulumi.OutputState }

func (AADAuthenticationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADAuthenticationSettings)(nil)).Elem()
}

func (o AADAuthenticationSettingsOutput) ToAADAuthenticationSettingsOutput() AADAuthenticationSettingsOutput {
	return o
}

func (o AADAuthenticationSettingsOutput) ToAADAuthenticationSettingsOutputWithContext(ctx context.Context) AADAuthenticationSettingsOutput {
	return o
}

func (o AADAuthenticationSettingsOutput) ToAADAuthenticationSettingsPtrOutput() AADAuthenticationSettingsPtrOutput {
	return o.ToAADAuthenticationSettingsPtrOutputWithContext(context.Background())
}

func (o AADAuthenticationSettingsOutput) ToAADAuthenticationSettingsPtrOutputWithContext(ctx context.Context) AADAuthenticationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AADAuthenticationSettings) *AADAuthenticationSettings {
		return &v
	}).(AADAuthenticationSettingsPtrOutput)
}

func (o AADAuthenticationSettingsOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADAuthenticationSettings) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type AADAuthenticationSettingsPtrOutput struct{ *pulumi.OutputState }

func (AADAuthenticationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADAuthenticationSettings)(nil)).Elem()
}

func (o AADAuthenticationSettingsPtrOutput) ToAADAuthenticationSettingsPtrOutput() AADAuthenticationSettingsPtrOutput {
	return o
}

func (o AADAuthenticationSettingsPtrOutput) ToAADAuthenticationSettingsPtrOutputWithContext(ctx context.Context) AADAuthenticationSettingsPtrOutput {
	return o
}

func (o AADAuthenticationSettingsPtrOutput) Elem() AADAuthenticationSettingsOutput {
	return o.ApplyT(func(v *AADAuthenticationSettings) AADAuthenticationSettings {
		if v != nil {
			return *v
		}
		var ret AADAuthenticationSettings
		return ret
	}).(AADAuthenticationSettingsOutput)
}

func (o AADAuthenticationSettingsPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADAuthenticationSettings) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type AADAuthenticationSettingsResponse struct {
	ClientId *string `pulumi:"clientId"`
}

type AADAuthenticationSettingsResponseOutput struct{ *pulumi.OutputState }

func (AADAuthenticationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AADAuthenticationSettingsResponse)(nil)).Elem()
}

func (o AADAuthenticationSettingsResponseOutput) ToAADAuthenticationSettingsResponseOutput() AADAuthenticationSettingsResponseOutput {
	return o
}

func (o AADAuthenticationSettingsResponseOutput) ToAADAuthenticationSettingsResponseOutputWithContext(ctx context.Context) AADAuthenticationSettingsResponseOutput {
	return o
}

func (o AADAuthenticationSettingsResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AADAuthenticationSettingsResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

type AADAuthenticationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AADAuthenticationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AADAuthenticationSettingsResponse)(nil)).Elem()
}

func (o AADAuthenticationSettingsResponsePtrOutput) ToAADAuthenticationSettingsResponsePtrOutput() AADAuthenticationSettingsResponsePtrOutput {
	return o
}

func (o AADAuthenticationSettingsResponsePtrOutput) ToAADAuthenticationSettingsResponsePtrOutputWithContext(ctx context.Context) AADAuthenticationSettingsResponsePtrOutput {
	return o
}

func (o AADAuthenticationSettingsResponsePtrOutput) Elem() AADAuthenticationSettingsResponseOutput {
	return o.ApplyT(func(v *AADAuthenticationSettingsResponse) AADAuthenticationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AADAuthenticationSettingsResponse
		return ret
	}).(AADAuthenticationSettingsResponseOutput)
}

func (o AADAuthenticationSettingsResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AADAuthenticationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

type AdditionalFeaturesServerConfigurations struct {
	IsRServicesEnabled *bool `pulumi:"isRServicesEnabled"`
}





type AdditionalFeaturesServerConfigurationsInput interface {
	pulumi.Input

	ToAdditionalFeaturesServerConfigurationsOutput() AdditionalFeaturesServerConfigurationsOutput
	ToAdditionalFeaturesServerConfigurationsOutputWithContext(context.Context) AdditionalFeaturesServerConfigurationsOutput
}

type AdditionalFeaturesServerConfigurationsArgs struct {
	IsRServicesEnabled pulumi.BoolPtrInput `pulumi:"isRServicesEnabled"`
}

func (AdditionalFeaturesServerConfigurationsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalFeaturesServerConfigurations)(nil)).Elem()
}

func (i AdditionalFeaturesServerConfigurationsArgs) ToAdditionalFeaturesServerConfigurationsOutput() AdditionalFeaturesServerConfigurationsOutput {
	return i.ToAdditionalFeaturesServerConfigurationsOutputWithContext(context.Background())
}

func (i AdditionalFeaturesServerConfigurationsArgs) ToAdditionalFeaturesServerConfigurationsOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalFeaturesServerConfigurationsOutput)
}

func (i AdditionalFeaturesServerConfigurationsArgs) ToAdditionalFeaturesServerConfigurationsPtrOutput() AdditionalFeaturesServerConfigurationsPtrOutput {
	return i.ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(context.Background())
}

func (i AdditionalFeaturesServerConfigurationsArgs) ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalFeaturesServerConfigurationsOutput).ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(ctx)
}









type AdditionalFeaturesServerConfigurationsPtrInput interface {
	pulumi.Input

	ToAdditionalFeaturesServerConfigurationsPtrOutput() AdditionalFeaturesServerConfigurationsPtrOutput
	ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(context.Context) AdditionalFeaturesServerConfigurationsPtrOutput
}

type additionalFeaturesServerConfigurationsPtrType AdditionalFeaturesServerConfigurationsArgs

func AdditionalFeaturesServerConfigurationsPtr(v *AdditionalFeaturesServerConfigurationsArgs) AdditionalFeaturesServerConfigurationsPtrInput {
	return (*additionalFeaturesServerConfigurationsPtrType)(v)
}

func (*additionalFeaturesServerConfigurationsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalFeaturesServerConfigurations)(nil)).Elem()
}

func (i *additionalFeaturesServerConfigurationsPtrType) ToAdditionalFeaturesServerConfigurationsPtrOutput() AdditionalFeaturesServerConfigurationsPtrOutput {
	return i.ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(context.Background())
}

func (i *additionalFeaturesServerConfigurationsPtrType) ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AdditionalFeaturesServerConfigurationsPtrOutput)
}

type AdditionalFeaturesServerConfigurationsOutput struct{ *pulumi.OutputState }

func (AdditionalFeaturesServerConfigurationsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalFeaturesServerConfigurations)(nil)).Elem()
}

func (o AdditionalFeaturesServerConfigurationsOutput) ToAdditionalFeaturesServerConfigurationsOutput() AdditionalFeaturesServerConfigurationsOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsOutput) ToAdditionalFeaturesServerConfigurationsOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsOutput) ToAdditionalFeaturesServerConfigurationsPtrOutput() AdditionalFeaturesServerConfigurationsPtrOutput {
	return o.ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(context.Background())
}

func (o AdditionalFeaturesServerConfigurationsOutput) ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AdditionalFeaturesServerConfigurations) *AdditionalFeaturesServerConfigurations {
		return &v
	}).(AdditionalFeaturesServerConfigurationsPtrOutput)
}

func (o AdditionalFeaturesServerConfigurationsOutput) IsRServicesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalFeaturesServerConfigurations) *bool { return v.IsRServicesEnabled }).(pulumi.BoolPtrOutput)
}

type AdditionalFeaturesServerConfigurationsPtrOutput struct{ *pulumi.OutputState }

func (AdditionalFeaturesServerConfigurationsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalFeaturesServerConfigurations)(nil)).Elem()
}

func (o AdditionalFeaturesServerConfigurationsPtrOutput) ToAdditionalFeaturesServerConfigurationsPtrOutput() AdditionalFeaturesServerConfigurationsPtrOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsPtrOutput) ToAdditionalFeaturesServerConfigurationsPtrOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsPtrOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsPtrOutput) Elem() AdditionalFeaturesServerConfigurationsOutput {
	return o.ApplyT(func(v *AdditionalFeaturesServerConfigurations) AdditionalFeaturesServerConfigurations {
		if v != nil {
			return *v
		}
		var ret AdditionalFeaturesServerConfigurations
		return ret
	}).(AdditionalFeaturesServerConfigurationsOutput)
}

func (o AdditionalFeaturesServerConfigurationsPtrOutput) IsRServicesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdditionalFeaturesServerConfigurations) *bool {
		if v == nil {
			return nil
		}
		return v.IsRServicesEnabled
	}).(pulumi.BoolPtrOutput)
}

type AdditionalFeaturesServerConfigurationsResponse struct {
	IsRServicesEnabled *bool `pulumi:"isRServicesEnabled"`
}

type AdditionalFeaturesServerConfigurationsResponseOutput struct{ *pulumi.OutputState }

func (AdditionalFeaturesServerConfigurationsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AdditionalFeaturesServerConfigurationsResponse)(nil)).Elem()
}

func (o AdditionalFeaturesServerConfigurationsResponseOutput) ToAdditionalFeaturesServerConfigurationsResponseOutput() AdditionalFeaturesServerConfigurationsResponseOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsResponseOutput) ToAdditionalFeaturesServerConfigurationsResponseOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsResponseOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsResponseOutput) IsRServicesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AdditionalFeaturesServerConfigurationsResponse) *bool { return v.IsRServicesEnabled }).(pulumi.BoolPtrOutput)
}

type AdditionalFeaturesServerConfigurationsResponsePtrOutput struct{ *pulumi.OutputState }

func (AdditionalFeaturesServerConfigurationsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AdditionalFeaturesServerConfigurationsResponse)(nil)).Elem()
}

func (o AdditionalFeaturesServerConfigurationsResponsePtrOutput) ToAdditionalFeaturesServerConfigurationsResponsePtrOutput() AdditionalFeaturesServerConfigurationsResponsePtrOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsResponsePtrOutput) ToAdditionalFeaturesServerConfigurationsResponsePtrOutputWithContext(ctx context.Context) AdditionalFeaturesServerConfigurationsResponsePtrOutput {
	return o
}

func (o AdditionalFeaturesServerConfigurationsResponsePtrOutput) Elem() AdditionalFeaturesServerConfigurationsResponseOutput {
	return o.ApplyT(func(v *AdditionalFeaturesServerConfigurationsResponse) AdditionalFeaturesServerConfigurationsResponse {
		if v != nil {
			return *v
		}
		var ret AdditionalFeaturesServerConfigurationsResponse
		return ret
	}).(AdditionalFeaturesServerConfigurationsResponseOutput)
}

func (o AdditionalFeaturesServerConfigurationsResponsePtrOutput) IsRServicesEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AdditionalFeaturesServerConfigurationsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsRServicesEnabled
	}).(pulumi.BoolPtrOutput)
}

type AgConfiguration struct {
	Replicas []AgReplica `pulumi:"replicas"`
}





type AgConfigurationInput interface {
	pulumi.Input

	ToAgConfigurationOutput() AgConfigurationOutput
	ToAgConfigurationOutputWithContext(context.Context) AgConfigurationOutput
}

type AgConfigurationArgs struct {
	Replicas AgReplicaArrayInput `pulumi:"replicas"`
}

func (AgConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgConfiguration)(nil)).Elem()
}

func (i AgConfigurationArgs) ToAgConfigurationOutput() AgConfigurationOutput {
	return i.ToAgConfigurationOutputWithContext(context.Background())
}

func (i AgConfigurationArgs) ToAgConfigurationOutputWithContext(ctx context.Context) AgConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgConfigurationOutput)
}

func (i AgConfigurationArgs) ToAgConfigurationPtrOutput() AgConfigurationPtrOutput {
	return i.ToAgConfigurationPtrOutputWithContext(context.Background())
}

func (i AgConfigurationArgs) ToAgConfigurationPtrOutputWithContext(ctx context.Context) AgConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgConfigurationOutput).ToAgConfigurationPtrOutputWithContext(ctx)
}









type AgConfigurationPtrInput interface {
	pulumi.Input

	ToAgConfigurationPtrOutput() AgConfigurationPtrOutput
	ToAgConfigurationPtrOutputWithContext(context.Context) AgConfigurationPtrOutput
}

type agConfigurationPtrType AgConfigurationArgs

func AgConfigurationPtr(v *AgConfigurationArgs) AgConfigurationPtrInput {
	return (*agConfigurationPtrType)(v)
}

func (*agConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AgConfiguration)(nil)).Elem()
}

func (i *agConfigurationPtrType) ToAgConfigurationPtrOutput() AgConfigurationPtrOutput {
	return i.ToAgConfigurationPtrOutputWithContext(context.Background())
}

func (i *agConfigurationPtrType) ToAgConfigurationPtrOutputWithContext(ctx context.Context) AgConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgConfigurationPtrOutput)
}

type AgConfigurationOutput struct{ *pulumi.OutputState }

func (AgConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgConfiguration)(nil)).Elem()
}

func (o AgConfigurationOutput) ToAgConfigurationOutput() AgConfigurationOutput {
	return o
}

func (o AgConfigurationOutput) ToAgConfigurationOutputWithContext(ctx context.Context) AgConfigurationOutput {
	return o
}

func (o AgConfigurationOutput) ToAgConfigurationPtrOutput() AgConfigurationPtrOutput {
	return o.ToAgConfigurationPtrOutputWithContext(context.Background())
}

func (o AgConfigurationOutput) ToAgConfigurationPtrOutputWithContext(ctx context.Context) AgConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AgConfiguration) *AgConfiguration {
		return &v
	}).(AgConfigurationPtrOutput)
}

func (o AgConfigurationOutput) Replicas() AgReplicaArrayOutput {
	return o.ApplyT(func(v AgConfiguration) []AgReplica { return v.Replicas }).(AgReplicaArrayOutput)
}

type AgConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AgConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgConfiguration)(nil)).Elem()
}

func (o AgConfigurationPtrOutput) ToAgConfigurationPtrOutput() AgConfigurationPtrOutput {
	return o
}

func (o AgConfigurationPtrOutput) ToAgConfigurationPtrOutputWithContext(ctx context.Context) AgConfigurationPtrOutput {
	return o
}

func (o AgConfigurationPtrOutput) Elem() AgConfigurationOutput {
	return o.ApplyT(func(v *AgConfiguration) AgConfiguration {
		if v != nil {
			return *v
		}
		var ret AgConfiguration
		return ret
	}).(AgConfigurationOutput)
}

func (o AgConfigurationPtrOutput) Replicas() AgReplicaArrayOutput {
	return o.ApplyT(func(v *AgConfiguration) []AgReplica {
		if v == nil {
			return nil
		}
		return v.Replicas
	}).(AgReplicaArrayOutput)
}

type AgConfigurationResponse struct {
	Replicas []AgReplicaResponse `pulumi:"replicas"`
}

type AgConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AgConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgConfigurationResponse)(nil)).Elem()
}

func (o AgConfigurationResponseOutput) ToAgConfigurationResponseOutput() AgConfigurationResponseOutput {
	return o
}

func (o AgConfigurationResponseOutput) ToAgConfigurationResponseOutputWithContext(ctx context.Context) AgConfigurationResponseOutput {
	return o
}

func (o AgConfigurationResponseOutput) Replicas() AgReplicaResponseArrayOutput {
	return o.ApplyT(func(v AgConfigurationResponse) []AgReplicaResponse { return v.Replicas }).(AgReplicaResponseArrayOutput)
}

type AgConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AgConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AgConfigurationResponse)(nil)).Elem()
}

func (o AgConfigurationResponsePtrOutput) ToAgConfigurationResponsePtrOutput() AgConfigurationResponsePtrOutput {
	return o
}

func (o AgConfigurationResponsePtrOutput) ToAgConfigurationResponsePtrOutputWithContext(ctx context.Context) AgConfigurationResponsePtrOutput {
	return o
}

func (o AgConfigurationResponsePtrOutput) Elem() AgConfigurationResponseOutput {
	return o.ApplyT(func(v *AgConfigurationResponse) AgConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AgConfigurationResponse
		return ret
	}).(AgConfigurationResponseOutput)
}

func (o AgConfigurationResponsePtrOutput) Replicas() AgReplicaResponseArrayOutput {
	return o.ApplyT(func(v *AgConfigurationResponse) []AgReplicaResponse {
		if v == nil {
			return nil
		}
		return v.Replicas
	}).(AgReplicaResponseArrayOutput)
}

type AgReplica struct {
	Commit                      *string `pulumi:"commit"`
	Failover                    *string `pulumi:"failover"`
	ReadableSecondary           *string `pulumi:"readableSecondary"`
	Role                        *string `pulumi:"role"`
	SqlVirtualMachineInstanceId *string `pulumi:"sqlVirtualMachineInstanceId"`
}





type AgReplicaInput interface {
	pulumi.Input

	ToAgReplicaOutput() AgReplicaOutput
	ToAgReplicaOutputWithContext(context.Context) AgReplicaOutput
}

type AgReplicaArgs struct {
	Commit                      pulumi.StringPtrInput `pulumi:"commit"`
	Failover                    pulumi.StringPtrInput `pulumi:"failover"`
	ReadableSecondary           pulumi.StringPtrInput `pulumi:"readableSecondary"`
	Role                        pulumi.StringPtrInput `pulumi:"role"`
	SqlVirtualMachineInstanceId pulumi.StringPtrInput `pulumi:"sqlVirtualMachineInstanceId"`
}

func (AgReplicaArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AgReplica)(nil)).Elem()
}

func (i AgReplicaArgs) ToAgReplicaOutput() AgReplicaOutput {
	return i.ToAgReplicaOutputWithContext(context.Background())
}

func (i AgReplicaArgs) ToAgReplicaOutputWithContext(ctx context.Context) AgReplicaOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgReplicaOutput)
}





type AgReplicaArrayInput interface {
	pulumi.Input

	ToAgReplicaArrayOutput() AgReplicaArrayOutput
	ToAgReplicaArrayOutputWithContext(context.Context) AgReplicaArrayOutput
}

type AgReplicaArray []AgReplicaInput

func (AgReplicaArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AgReplica)(nil)).Elem()
}

func (i AgReplicaArray) ToAgReplicaArrayOutput() AgReplicaArrayOutput {
	return i.ToAgReplicaArrayOutputWithContext(context.Background())
}

func (i AgReplicaArray) ToAgReplicaArrayOutputWithContext(ctx context.Context) AgReplicaArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AgReplicaArrayOutput)
}

type AgReplicaOutput struct{ *pulumi.OutputState }

func (AgReplicaOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgReplica)(nil)).Elem()
}

func (o AgReplicaOutput) ToAgReplicaOutput() AgReplicaOutput {
	return o
}

func (o AgReplicaOutput) ToAgReplicaOutputWithContext(ctx context.Context) AgReplicaOutput {
	return o
}

func (o AgReplicaOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplica) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o AgReplicaOutput) Failover() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplica) *string { return v.Failover }).(pulumi.StringPtrOutput)
}

func (o AgReplicaOutput) ReadableSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplica) *string { return v.ReadableSecondary }).(pulumi.StringPtrOutput)
}

func (o AgReplicaOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplica) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o AgReplicaOutput) SqlVirtualMachineInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplica) *string { return v.SqlVirtualMachineInstanceId }).(pulumi.StringPtrOutput)
}

type AgReplicaArrayOutput struct{ *pulumi.OutputState }

func (AgReplicaArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AgReplica)(nil)).Elem()
}

func (o AgReplicaArrayOutput) ToAgReplicaArrayOutput() AgReplicaArrayOutput {
	return o
}

func (o AgReplicaArrayOutput) ToAgReplicaArrayOutputWithContext(ctx context.Context) AgReplicaArrayOutput {
	return o
}

func (o AgReplicaArrayOutput) Index(i pulumi.IntInput) AgReplicaOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AgReplica {
		return vs[0].([]AgReplica)[vs[1].(int)]
	}).(AgReplicaOutput)
}

type AgReplicaResponse struct {
	Commit                      *string `pulumi:"commit"`
	Failover                    *string `pulumi:"failover"`
	ReadableSecondary           *string `pulumi:"readableSecondary"`
	Role                        *string `pulumi:"role"`
	SqlVirtualMachineInstanceId *string `pulumi:"sqlVirtualMachineInstanceId"`
}

type AgReplicaResponseOutput struct{ *pulumi.OutputState }

func (AgReplicaResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AgReplicaResponse)(nil)).Elem()
}

func (o AgReplicaResponseOutput) ToAgReplicaResponseOutput() AgReplicaResponseOutput {
	return o
}

func (o AgReplicaResponseOutput) ToAgReplicaResponseOutputWithContext(ctx context.Context) AgReplicaResponseOutput {
	return o
}

func (o AgReplicaResponseOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplicaResponse) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o AgReplicaResponseOutput) Failover() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplicaResponse) *string { return v.Failover }).(pulumi.StringPtrOutput)
}

func (o AgReplicaResponseOutput) ReadableSecondary() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplicaResponse) *string { return v.ReadableSecondary }).(pulumi.StringPtrOutput)
}

func (o AgReplicaResponseOutput) Role() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplicaResponse) *string { return v.Role }).(pulumi.StringPtrOutput)
}

func (o AgReplicaResponseOutput) SqlVirtualMachineInstanceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AgReplicaResponse) *string { return v.SqlVirtualMachineInstanceId }).(pulumi.StringPtrOutput)
}

type AgReplicaResponseArrayOutput struct{ *pulumi.OutputState }

func (AgReplicaResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AgReplicaResponse)(nil)).Elem()
}

func (o AgReplicaResponseArrayOutput) ToAgReplicaResponseArrayOutput() AgReplicaResponseArrayOutput {
	return o
}

func (o AgReplicaResponseArrayOutput) ToAgReplicaResponseArrayOutputWithContext(ctx context.Context) AgReplicaResponseArrayOutput {
	return o
}

func (o AgReplicaResponseArrayOutput) Index(i pulumi.IntInput) AgReplicaResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AgReplicaResponse {
		return vs[0].([]AgReplicaResponse)[vs[1].(int)]
	}).(AgReplicaResponseOutput)
}

type AssessmentSettings struct {
	Enable         *bool     `pulumi:"enable"`
	RunImmediately *bool     `pulumi:"runImmediately"`
	Schedule       *Schedule `pulumi:"schedule"`
}





type AssessmentSettingsInput interface {
	pulumi.Input

	ToAssessmentSettingsOutput() AssessmentSettingsOutput
	ToAssessmentSettingsOutputWithContext(context.Context) AssessmentSettingsOutput
}

type AssessmentSettingsArgs struct {
	Enable         pulumi.BoolPtrInput `pulumi:"enable"`
	RunImmediately pulumi.BoolPtrInput `pulumi:"runImmediately"`
	Schedule       SchedulePtrInput    `pulumi:"schedule"`
}

func (AssessmentSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentSettings)(nil)).Elem()
}

func (i AssessmentSettingsArgs) ToAssessmentSettingsOutput() AssessmentSettingsOutput {
	return i.ToAssessmentSettingsOutputWithContext(context.Background())
}

func (i AssessmentSettingsArgs) ToAssessmentSettingsOutputWithContext(ctx context.Context) AssessmentSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentSettingsOutput)
}

func (i AssessmentSettingsArgs) ToAssessmentSettingsPtrOutput() AssessmentSettingsPtrOutput {
	return i.ToAssessmentSettingsPtrOutputWithContext(context.Background())
}

func (i AssessmentSettingsArgs) ToAssessmentSettingsPtrOutputWithContext(ctx context.Context) AssessmentSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentSettingsOutput).ToAssessmentSettingsPtrOutputWithContext(ctx)
}









type AssessmentSettingsPtrInput interface {
	pulumi.Input

	ToAssessmentSettingsPtrOutput() AssessmentSettingsPtrOutput
	ToAssessmentSettingsPtrOutputWithContext(context.Context) AssessmentSettingsPtrOutput
}

type assessmentSettingsPtrType AssessmentSettingsArgs

func AssessmentSettingsPtr(v *AssessmentSettingsArgs) AssessmentSettingsPtrInput {
	return (*assessmentSettingsPtrType)(v)
}

func (*assessmentSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentSettings)(nil)).Elem()
}

func (i *assessmentSettingsPtrType) ToAssessmentSettingsPtrOutput() AssessmentSettingsPtrOutput {
	return i.ToAssessmentSettingsPtrOutputWithContext(context.Background())
}

func (i *assessmentSettingsPtrType) ToAssessmentSettingsPtrOutputWithContext(ctx context.Context) AssessmentSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AssessmentSettingsPtrOutput)
}

type AssessmentSettingsOutput struct{ *pulumi.OutputState }

func (AssessmentSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentSettings)(nil)).Elem()
}

func (o AssessmentSettingsOutput) ToAssessmentSettingsOutput() AssessmentSettingsOutput {
	return o
}

func (o AssessmentSettingsOutput) ToAssessmentSettingsOutputWithContext(ctx context.Context) AssessmentSettingsOutput {
	return o
}

func (o AssessmentSettingsOutput) ToAssessmentSettingsPtrOutput() AssessmentSettingsPtrOutput {
	return o.ToAssessmentSettingsPtrOutputWithContext(context.Background())
}

func (o AssessmentSettingsOutput) ToAssessmentSettingsPtrOutputWithContext(ctx context.Context) AssessmentSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AssessmentSettings) *AssessmentSettings {
		return &v
	}).(AssessmentSettingsPtrOutput)
}

func (o AssessmentSettingsOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AssessmentSettings) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsOutput) RunImmediately() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AssessmentSettings) *bool { return v.RunImmediately }).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsOutput) Schedule() SchedulePtrOutput {
	return o.ApplyT(func(v AssessmentSettings) *Schedule { return v.Schedule }).(SchedulePtrOutput)
}

type AssessmentSettingsPtrOutput struct{ *pulumi.OutputState }

func (AssessmentSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentSettings)(nil)).Elem()
}

func (o AssessmentSettingsPtrOutput) ToAssessmentSettingsPtrOutput() AssessmentSettingsPtrOutput {
	return o
}

func (o AssessmentSettingsPtrOutput) ToAssessmentSettingsPtrOutputWithContext(ctx context.Context) AssessmentSettingsPtrOutput {
	return o
}

func (o AssessmentSettingsPtrOutput) Elem() AssessmentSettingsOutput {
	return o.ApplyT(func(v *AssessmentSettings) AssessmentSettings {
		if v != nil {
			return *v
		}
		var ret AssessmentSettings
		return ret
	}).(AssessmentSettingsOutput)
}

func (o AssessmentSettingsPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsPtrOutput) RunImmediately() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentSettings) *bool {
		if v == nil {
			return nil
		}
		return v.RunImmediately
	}).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsPtrOutput) Schedule() SchedulePtrOutput {
	return o.ApplyT(func(v *AssessmentSettings) *Schedule {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(SchedulePtrOutput)
}

type AssessmentSettingsResponse struct {
	Enable         *bool             `pulumi:"enable"`
	RunImmediately *bool             `pulumi:"runImmediately"`
	Schedule       *ScheduleResponse `pulumi:"schedule"`
}

type AssessmentSettingsResponseOutput struct{ *pulumi.OutputState }

func (AssessmentSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AssessmentSettingsResponse)(nil)).Elem()
}

func (o AssessmentSettingsResponseOutput) ToAssessmentSettingsResponseOutput() AssessmentSettingsResponseOutput {
	return o
}

func (o AssessmentSettingsResponseOutput) ToAssessmentSettingsResponseOutputWithContext(ctx context.Context) AssessmentSettingsResponseOutput {
	return o
}

func (o AssessmentSettingsResponseOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AssessmentSettingsResponse) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsResponseOutput) RunImmediately() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AssessmentSettingsResponse) *bool { return v.RunImmediately }).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsResponseOutput) Schedule() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v AssessmentSettingsResponse) *ScheduleResponse { return v.Schedule }).(ScheduleResponsePtrOutput)
}

type AssessmentSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AssessmentSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AssessmentSettingsResponse)(nil)).Elem()
}

func (o AssessmentSettingsResponsePtrOutput) ToAssessmentSettingsResponsePtrOutput() AssessmentSettingsResponsePtrOutput {
	return o
}

func (o AssessmentSettingsResponsePtrOutput) ToAssessmentSettingsResponsePtrOutputWithContext(ctx context.Context) AssessmentSettingsResponsePtrOutput {
	return o
}

func (o AssessmentSettingsResponsePtrOutput) Elem() AssessmentSettingsResponseOutput {
	return o.ApplyT(func(v *AssessmentSettingsResponse) AssessmentSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AssessmentSettingsResponse
		return ret
	}).(AssessmentSettingsResponseOutput)
}

func (o AssessmentSettingsResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsResponsePtrOutput) RunImmediately() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AssessmentSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RunImmediately
	}).(pulumi.BoolPtrOutput)
}

func (o AssessmentSettingsResponsePtrOutput) Schedule() ScheduleResponsePtrOutput {
	return o.ApplyT(func(v *AssessmentSettingsResponse) *ScheduleResponse {
		if v == nil {
			return nil
		}
		return v.Schedule
	}).(ScheduleResponsePtrOutput)
}

type AutoBackupSettings struct {
	BackupScheduleType    *string  `pulumi:"backupScheduleType"`
	BackupSystemDbs       *bool    `pulumi:"backupSystemDbs"`
	DaysOfWeek            []string `pulumi:"daysOfWeek"`
	Enable                *bool    `pulumi:"enable"`
	EnableEncryption      *bool    `pulumi:"enableEncryption"`
	FullBackupFrequency   *string  `pulumi:"fullBackupFrequency"`
	FullBackupStartTime   *int     `pulumi:"fullBackupStartTime"`
	FullBackupWindowHours *int     `pulumi:"fullBackupWindowHours"`
	LogBackupFrequency    *int     `pulumi:"logBackupFrequency"`
	Password              *string  `pulumi:"password"`
	RetentionPeriod       *int     `pulumi:"retentionPeriod"`
	StorageAccessKey      *string  `pulumi:"storageAccessKey"`
	StorageAccountUrl     *string  `pulumi:"storageAccountUrl"`
	StorageContainerName  *string  `pulumi:"storageContainerName"`
}





type AutoBackupSettingsInput interface {
	pulumi.Input

	ToAutoBackupSettingsOutput() AutoBackupSettingsOutput
	ToAutoBackupSettingsOutputWithContext(context.Context) AutoBackupSettingsOutput
}

type AutoBackupSettingsArgs struct {
	BackupScheduleType    pulumi.StringPtrInput   `pulumi:"backupScheduleType"`
	BackupSystemDbs       pulumi.BoolPtrInput     `pulumi:"backupSystemDbs"`
	DaysOfWeek            pulumi.StringArrayInput `pulumi:"daysOfWeek"`
	Enable                pulumi.BoolPtrInput     `pulumi:"enable"`
	EnableEncryption      pulumi.BoolPtrInput     `pulumi:"enableEncryption"`
	FullBackupFrequency   pulumi.StringPtrInput   `pulumi:"fullBackupFrequency"`
	FullBackupStartTime   pulumi.IntPtrInput      `pulumi:"fullBackupStartTime"`
	FullBackupWindowHours pulumi.IntPtrInput      `pulumi:"fullBackupWindowHours"`
	LogBackupFrequency    pulumi.IntPtrInput      `pulumi:"logBackupFrequency"`
	Password              pulumi.StringPtrInput   `pulumi:"password"`
	RetentionPeriod       pulumi.IntPtrInput      `pulumi:"retentionPeriod"`
	StorageAccessKey      pulumi.StringPtrInput   `pulumi:"storageAccessKey"`
	StorageAccountUrl     pulumi.StringPtrInput   `pulumi:"storageAccountUrl"`
	StorageContainerName  pulumi.StringPtrInput   `pulumi:"storageContainerName"`
}

func (AutoBackupSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoBackupSettings)(nil)).Elem()
}

func (i AutoBackupSettingsArgs) ToAutoBackupSettingsOutput() AutoBackupSettingsOutput {
	return i.ToAutoBackupSettingsOutputWithContext(context.Background())
}

func (i AutoBackupSettingsArgs) ToAutoBackupSettingsOutputWithContext(ctx context.Context) AutoBackupSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoBackupSettingsOutput)
}

func (i AutoBackupSettingsArgs) ToAutoBackupSettingsPtrOutput() AutoBackupSettingsPtrOutput {
	return i.ToAutoBackupSettingsPtrOutputWithContext(context.Background())
}

func (i AutoBackupSettingsArgs) ToAutoBackupSettingsPtrOutputWithContext(ctx context.Context) AutoBackupSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoBackupSettingsOutput).ToAutoBackupSettingsPtrOutputWithContext(ctx)
}









type AutoBackupSettingsPtrInput interface {
	pulumi.Input

	ToAutoBackupSettingsPtrOutput() AutoBackupSettingsPtrOutput
	ToAutoBackupSettingsPtrOutputWithContext(context.Context) AutoBackupSettingsPtrOutput
}

type autoBackupSettingsPtrType AutoBackupSettingsArgs

func AutoBackupSettingsPtr(v *AutoBackupSettingsArgs) AutoBackupSettingsPtrInput {
	return (*autoBackupSettingsPtrType)(v)
}

func (*autoBackupSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoBackupSettings)(nil)).Elem()
}

func (i *autoBackupSettingsPtrType) ToAutoBackupSettingsPtrOutput() AutoBackupSettingsPtrOutput {
	return i.ToAutoBackupSettingsPtrOutputWithContext(context.Background())
}

func (i *autoBackupSettingsPtrType) ToAutoBackupSettingsPtrOutputWithContext(ctx context.Context) AutoBackupSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoBackupSettingsPtrOutput)
}

type AutoBackupSettingsOutput struct{ *pulumi.OutputState }

func (AutoBackupSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoBackupSettings)(nil)).Elem()
}

func (o AutoBackupSettingsOutput) ToAutoBackupSettingsOutput() AutoBackupSettingsOutput {
	return o
}

func (o AutoBackupSettingsOutput) ToAutoBackupSettingsOutputWithContext(ctx context.Context) AutoBackupSettingsOutput {
	return o
}

func (o AutoBackupSettingsOutput) ToAutoBackupSettingsPtrOutput() AutoBackupSettingsPtrOutput {
	return o.ToAutoBackupSettingsPtrOutputWithContext(context.Background())
}

func (o AutoBackupSettingsOutput) ToAutoBackupSettingsPtrOutputWithContext(ctx context.Context) AutoBackupSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoBackupSettings) *AutoBackupSettings {
		return &v
	}).(AutoBackupSettingsPtrOutput)
}

func (o AutoBackupSettingsOutput) BackupScheduleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *string { return v.BackupScheduleType }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsOutput) BackupSystemDbs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *bool { return v.BackupSystemDbs }).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AutoBackupSettings) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o AutoBackupSettingsOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsOutput) EnableEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *bool { return v.EnableEncryption }).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsOutput) FullBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *string { return v.FullBackupFrequency }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsOutput) FullBackupStartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *int { return v.FullBackupStartTime }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsOutput) FullBackupWindowHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *int { return v.FullBackupWindowHours }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsOutput) LogBackupFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *int { return v.LogBackupFrequency }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *int { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsOutput) StorageAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *string { return v.StorageAccessKey }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsOutput) StorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettings) *string { return v.StorageContainerName }).(pulumi.StringPtrOutput)
}

type AutoBackupSettingsPtrOutput struct{ *pulumi.OutputState }

func (AutoBackupSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoBackupSettings)(nil)).Elem()
}

func (o AutoBackupSettingsPtrOutput) ToAutoBackupSettingsPtrOutput() AutoBackupSettingsPtrOutput {
	return o
}

func (o AutoBackupSettingsPtrOutput) ToAutoBackupSettingsPtrOutputWithContext(ctx context.Context) AutoBackupSettingsPtrOutput {
	return o
}

func (o AutoBackupSettingsPtrOutput) Elem() AutoBackupSettingsOutput {
	return o.ApplyT(func(v *AutoBackupSettings) AutoBackupSettings {
		if v != nil {
			return *v
		}
		var ret AutoBackupSettings
		return ret
	}).(AutoBackupSettingsOutput)
}

func (o AutoBackupSettingsPtrOutput) BackupScheduleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *string {
		if v == nil {
			return nil
		}
		return v.BackupScheduleType
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) BackupSystemDbs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *bool {
		if v == nil {
			return nil
		}
		return v.BackupSystemDbs
	}).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoBackupSettings) []string {
		if v == nil {
			return nil
		}
		return v.DaysOfWeek
	}).(pulumi.StringArrayOutput)
}

func (o AutoBackupSettingsPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) EnableEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *bool {
		if v == nil {
			return nil
		}
		return v.EnableEncryption
	}).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) FullBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *string {
		if v == nil {
			return nil
		}
		return v.FullBackupFrequency
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) FullBackupStartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *int {
		if v == nil {
			return nil
		}
		return v.FullBackupStartTime
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) FullBackupWindowHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *int {
		if v == nil {
			return nil
		}
		return v.FullBackupWindowHours
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) LogBackupFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *int {
		if v == nil {
			return nil
		}
		return v.LogBackupFrequency
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *int {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) StorageAccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccessKey
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountUrl
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsPtrOutput) StorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettings) *string {
		if v == nil {
			return nil
		}
		return v.StorageContainerName
	}).(pulumi.StringPtrOutput)
}

type AutoBackupSettingsResponse struct {
	BackupScheduleType    *string  `pulumi:"backupScheduleType"`
	BackupSystemDbs       *bool    `pulumi:"backupSystemDbs"`
	DaysOfWeek            []string `pulumi:"daysOfWeek"`
	Enable                *bool    `pulumi:"enable"`
	EnableEncryption      *bool    `pulumi:"enableEncryption"`
	FullBackupFrequency   *string  `pulumi:"fullBackupFrequency"`
	FullBackupStartTime   *int     `pulumi:"fullBackupStartTime"`
	FullBackupWindowHours *int     `pulumi:"fullBackupWindowHours"`
	LogBackupFrequency    *int     `pulumi:"logBackupFrequency"`
	RetentionPeriod       *int     `pulumi:"retentionPeriod"`
	StorageAccountUrl     *string  `pulumi:"storageAccountUrl"`
	StorageContainerName  *string  `pulumi:"storageContainerName"`
}

type AutoBackupSettingsResponseOutput struct{ *pulumi.OutputState }

func (AutoBackupSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoBackupSettingsResponse)(nil)).Elem()
}

func (o AutoBackupSettingsResponseOutput) ToAutoBackupSettingsResponseOutput() AutoBackupSettingsResponseOutput {
	return o
}

func (o AutoBackupSettingsResponseOutput) ToAutoBackupSettingsResponseOutputWithContext(ctx context.Context) AutoBackupSettingsResponseOutput {
	return o
}

func (o AutoBackupSettingsResponseOutput) BackupScheduleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *string { return v.BackupScheduleType }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) BackupSystemDbs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *bool { return v.BackupSystemDbs }).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) []string { return v.DaysOfWeek }).(pulumi.StringArrayOutput)
}

func (o AutoBackupSettingsResponseOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) EnableEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *bool { return v.EnableEncryption }).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) FullBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *string { return v.FullBackupFrequency }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) FullBackupStartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *int { return v.FullBackupStartTime }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) FullBackupWindowHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *int { return v.FullBackupWindowHours }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) LogBackupFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *int { return v.LogBackupFrequency }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *int { return v.RetentionPeriod }).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsResponseOutput) StorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoBackupSettingsResponse) *string { return v.StorageContainerName }).(pulumi.StringPtrOutput)
}

type AutoBackupSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoBackupSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoBackupSettingsResponse)(nil)).Elem()
}

func (o AutoBackupSettingsResponsePtrOutput) ToAutoBackupSettingsResponsePtrOutput() AutoBackupSettingsResponsePtrOutput {
	return o
}

func (o AutoBackupSettingsResponsePtrOutput) ToAutoBackupSettingsResponsePtrOutputWithContext(ctx context.Context) AutoBackupSettingsResponsePtrOutput {
	return o
}

func (o AutoBackupSettingsResponsePtrOutput) Elem() AutoBackupSettingsResponseOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) AutoBackupSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AutoBackupSettingsResponse
		return ret
	}).(AutoBackupSettingsResponseOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) BackupScheduleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BackupScheduleType
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) BackupSystemDbs() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.BackupSystemDbs
	}).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) DaysOfWeek() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.DaysOfWeek
	}).(pulumi.StringArrayOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) EnableEncryption() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableEncryption
	}).(pulumi.BoolPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) FullBackupFrequency() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.FullBackupFrequency
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) FullBackupStartTime() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.FullBackupStartTime
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) FullBackupWindowHours() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.FullBackupWindowHours
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) LogBackupFrequency() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogBackupFrequency
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) RetentionPeriod() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionPeriod
	}).(pulumi.IntPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountUrl
	}).(pulumi.StringPtrOutput)
}

func (o AutoBackupSettingsResponsePtrOutput) StorageContainerName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoBackupSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageContainerName
	}).(pulumi.StringPtrOutput)
}

type AutoPatchingSettings struct {
	DayOfWeek                     *DayOfWeek `pulumi:"dayOfWeek"`
	Enable                        *bool      `pulumi:"enable"`
	MaintenanceWindowDuration     *int       `pulumi:"maintenanceWindowDuration"`
	MaintenanceWindowStartingHour *int       `pulumi:"maintenanceWindowStartingHour"`
}





type AutoPatchingSettingsInput interface {
	pulumi.Input

	ToAutoPatchingSettingsOutput() AutoPatchingSettingsOutput
	ToAutoPatchingSettingsOutputWithContext(context.Context) AutoPatchingSettingsOutput
}

type AutoPatchingSettingsArgs struct {
	DayOfWeek                     DayOfWeekPtrInput   `pulumi:"dayOfWeek"`
	Enable                        pulumi.BoolPtrInput `pulumi:"enable"`
	MaintenanceWindowDuration     pulumi.IntPtrInput  `pulumi:"maintenanceWindowDuration"`
	MaintenanceWindowStartingHour pulumi.IntPtrInput  `pulumi:"maintenanceWindowStartingHour"`
}

func (AutoPatchingSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPatchingSettings)(nil)).Elem()
}

func (i AutoPatchingSettingsArgs) ToAutoPatchingSettingsOutput() AutoPatchingSettingsOutput {
	return i.ToAutoPatchingSettingsOutputWithContext(context.Background())
}

func (i AutoPatchingSettingsArgs) ToAutoPatchingSettingsOutputWithContext(ctx context.Context) AutoPatchingSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPatchingSettingsOutput)
}

func (i AutoPatchingSettingsArgs) ToAutoPatchingSettingsPtrOutput() AutoPatchingSettingsPtrOutput {
	return i.ToAutoPatchingSettingsPtrOutputWithContext(context.Background())
}

func (i AutoPatchingSettingsArgs) ToAutoPatchingSettingsPtrOutputWithContext(ctx context.Context) AutoPatchingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPatchingSettingsOutput).ToAutoPatchingSettingsPtrOutputWithContext(ctx)
}









type AutoPatchingSettingsPtrInput interface {
	pulumi.Input

	ToAutoPatchingSettingsPtrOutput() AutoPatchingSettingsPtrOutput
	ToAutoPatchingSettingsPtrOutputWithContext(context.Context) AutoPatchingSettingsPtrOutput
}

type autoPatchingSettingsPtrType AutoPatchingSettingsArgs

func AutoPatchingSettingsPtr(v *AutoPatchingSettingsArgs) AutoPatchingSettingsPtrInput {
	return (*autoPatchingSettingsPtrType)(v)
}

func (*autoPatchingSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPatchingSettings)(nil)).Elem()
}

func (i *autoPatchingSettingsPtrType) ToAutoPatchingSettingsPtrOutput() AutoPatchingSettingsPtrOutput {
	return i.ToAutoPatchingSettingsPtrOutputWithContext(context.Background())
}

func (i *autoPatchingSettingsPtrType) ToAutoPatchingSettingsPtrOutputWithContext(ctx context.Context) AutoPatchingSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoPatchingSettingsPtrOutput)
}

type AutoPatchingSettingsOutput struct{ *pulumi.OutputState }

func (AutoPatchingSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPatchingSettings)(nil)).Elem()
}

func (o AutoPatchingSettingsOutput) ToAutoPatchingSettingsOutput() AutoPatchingSettingsOutput {
	return o
}

func (o AutoPatchingSettingsOutput) ToAutoPatchingSettingsOutputWithContext(ctx context.Context) AutoPatchingSettingsOutput {
	return o
}

func (o AutoPatchingSettingsOutput) ToAutoPatchingSettingsPtrOutput() AutoPatchingSettingsPtrOutput {
	return o.ToAutoPatchingSettingsPtrOutputWithContext(context.Background())
}

func (o AutoPatchingSettingsOutput) ToAutoPatchingSettingsPtrOutputWithContext(ctx context.Context) AutoPatchingSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoPatchingSettings) *AutoPatchingSettings {
		return &v
	}).(AutoPatchingSettingsPtrOutput)
}

func (o AutoPatchingSettingsOutput) DayOfWeek() DayOfWeekPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettings) *DayOfWeek { return v.DayOfWeek }).(DayOfWeekPtrOutput)
}

func (o AutoPatchingSettingsOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettings) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o AutoPatchingSettingsOutput) MaintenanceWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettings) *int { return v.MaintenanceWindowDuration }).(pulumi.IntPtrOutput)
}

func (o AutoPatchingSettingsOutput) MaintenanceWindowStartingHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettings) *int { return v.MaintenanceWindowStartingHour }).(pulumi.IntPtrOutput)
}

type AutoPatchingSettingsPtrOutput struct{ *pulumi.OutputState }

func (AutoPatchingSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPatchingSettings)(nil)).Elem()
}

func (o AutoPatchingSettingsPtrOutput) ToAutoPatchingSettingsPtrOutput() AutoPatchingSettingsPtrOutput {
	return o
}

func (o AutoPatchingSettingsPtrOutput) ToAutoPatchingSettingsPtrOutputWithContext(ctx context.Context) AutoPatchingSettingsPtrOutput {
	return o
}

func (o AutoPatchingSettingsPtrOutput) Elem() AutoPatchingSettingsOutput {
	return o.ApplyT(func(v *AutoPatchingSettings) AutoPatchingSettings {
		if v != nil {
			return *v
		}
		var ret AutoPatchingSettings
		return ret
	}).(AutoPatchingSettingsOutput)
}

func (o AutoPatchingSettingsPtrOutput) DayOfWeek() DayOfWeekPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettings) *DayOfWeek {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(DayOfWeekPtrOutput)
}

func (o AutoPatchingSettingsPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o AutoPatchingSettingsPtrOutput) MaintenanceWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettings) *int {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowDuration
	}).(pulumi.IntPtrOutput)
}

func (o AutoPatchingSettingsPtrOutput) MaintenanceWindowStartingHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettings) *int {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowStartingHour
	}).(pulumi.IntPtrOutput)
}

type AutoPatchingSettingsResponse struct {
	DayOfWeek                     *string `pulumi:"dayOfWeek"`
	Enable                        *bool   `pulumi:"enable"`
	MaintenanceWindowDuration     *int    `pulumi:"maintenanceWindowDuration"`
	MaintenanceWindowStartingHour *int    `pulumi:"maintenanceWindowStartingHour"`
}

type AutoPatchingSettingsResponseOutput struct{ *pulumi.OutputState }

func (AutoPatchingSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoPatchingSettingsResponse)(nil)).Elem()
}

func (o AutoPatchingSettingsResponseOutput) ToAutoPatchingSettingsResponseOutput() AutoPatchingSettingsResponseOutput {
	return o
}

func (o AutoPatchingSettingsResponseOutput) ToAutoPatchingSettingsResponseOutputWithContext(ctx context.Context) AutoPatchingSettingsResponseOutput {
	return o
}

func (o AutoPatchingSettingsResponseOutput) DayOfWeek() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettingsResponse) *string { return v.DayOfWeek }).(pulumi.StringPtrOutput)
}

func (o AutoPatchingSettingsResponseOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettingsResponse) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o AutoPatchingSettingsResponseOutput) MaintenanceWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettingsResponse) *int { return v.MaintenanceWindowDuration }).(pulumi.IntPtrOutput)
}

func (o AutoPatchingSettingsResponseOutput) MaintenanceWindowStartingHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoPatchingSettingsResponse) *int { return v.MaintenanceWindowStartingHour }).(pulumi.IntPtrOutput)
}

type AutoPatchingSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoPatchingSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoPatchingSettingsResponse)(nil)).Elem()
}

func (o AutoPatchingSettingsResponsePtrOutput) ToAutoPatchingSettingsResponsePtrOutput() AutoPatchingSettingsResponsePtrOutput {
	return o
}

func (o AutoPatchingSettingsResponsePtrOutput) ToAutoPatchingSettingsResponsePtrOutputWithContext(ctx context.Context) AutoPatchingSettingsResponsePtrOutput {
	return o
}

func (o AutoPatchingSettingsResponsePtrOutput) Elem() AutoPatchingSettingsResponseOutput {
	return o.ApplyT(func(v *AutoPatchingSettingsResponse) AutoPatchingSettingsResponse {
		if v != nil {
			return *v
		}
		var ret AutoPatchingSettingsResponse
		return ret
	}).(AutoPatchingSettingsResponseOutput)
}

func (o AutoPatchingSettingsResponsePtrOutput) DayOfWeek() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.StringPtrOutput)
}

func (o AutoPatchingSettingsResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o AutoPatchingSettingsResponsePtrOutput) MaintenanceWindowDuration() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowDuration
	}).(pulumi.IntPtrOutput)
}

func (o AutoPatchingSettingsResponsePtrOutput) MaintenanceWindowStartingHour() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoPatchingSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaintenanceWindowStartingHour
	}).(pulumi.IntPtrOutput)
}

type KeyVaultCredentialSettings struct {
	AzureKeyVaultUrl       *string `pulumi:"azureKeyVaultUrl"`
	CredentialName         *string `pulumi:"credentialName"`
	Enable                 *bool   `pulumi:"enable"`
	ServicePrincipalName   *string `pulumi:"servicePrincipalName"`
	ServicePrincipalSecret *string `pulumi:"servicePrincipalSecret"`
}





type KeyVaultCredentialSettingsInput interface {
	pulumi.Input

	ToKeyVaultCredentialSettingsOutput() KeyVaultCredentialSettingsOutput
	ToKeyVaultCredentialSettingsOutputWithContext(context.Context) KeyVaultCredentialSettingsOutput
}

type KeyVaultCredentialSettingsArgs struct {
	AzureKeyVaultUrl       pulumi.StringPtrInput `pulumi:"azureKeyVaultUrl"`
	CredentialName         pulumi.StringPtrInput `pulumi:"credentialName"`
	Enable                 pulumi.BoolPtrInput   `pulumi:"enable"`
	ServicePrincipalName   pulumi.StringPtrInput `pulumi:"servicePrincipalName"`
	ServicePrincipalSecret pulumi.StringPtrInput `pulumi:"servicePrincipalSecret"`
}

func (KeyVaultCredentialSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCredentialSettings)(nil)).Elem()
}

func (i KeyVaultCredentialSettingsArgs) ToKeyVaultCredentialSettingsOutput() KeyVaultCredentialSettingsOutput {
	return i.ToKeyVaultCredentialSettingsOutputWithContext(context.Background())
}

func (i KeyVaultCredentialSettingsArgs) ToKeyVaultCredentialSettingsOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCredentialSettingsOutput)
}

func (i KeyVaultCredentialSettingsArgs) ToKeyVaultCredentialSettingsPtrOutput() KeyVaultCredentialSettingsPtrOutput {
	return i.ToKeyVaultCredentialSettingsPtrOutputWithContext(context.Background())
}

func (i KeyVaultCredentialSettingsArgs) ToKeyVaultCredentialSettingsPtrOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCredentialSettingsOutput).ToKeyVaultCredentialSettingsPtrOutputWithContext(ctx)
}









type KeyVaultCredentialSettingsPtrInput interface {
	pulumi.Input

	ToKeyVaultCredentialSettingsPtrOutput() KeyVaultCredentialSettingsPtrOutput
	ToKeyVaultCredentialSettingsPtrOutputWithContext(context.Context) KeyVaultCredentialSettingsPtrOutput
}

type keyVaultCredentialSettingsPtrType KeyVaultCredentialSettingsArgs

func KeyVaultCredentialSettingsPtr(v *KeyVaultCredentialSettingsArgs) KeyVaultCredentialSettingsPtrInput {
	return (*keyVaultCredentialSettingsPtrType)(v)
}

func (*keyVaultCredentialSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCredentialSettings)(nil)).Elem()
}

func (i *keyVaultCredentialSettingsPtrType) ToKeyVaultCredentialSettingsPtrOutput() KeyVaultCredentialSettingsPtrOutput {
	return i.ToKeyVaultCredentialSettingsPtrOutputWithContext(context.Background())
}

func (i *keyVaultCredentialSettingsPtrType) ToKeyVaultCredentialSettingsPtrOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KeyVaultCredentialSettingsPtrOutput)
}

type KeyVaultCredentialSettingsOutput struct{ *pulumi.OutputState }

func (KeyVaultCredentialSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCredentialSettings)(nil)).Elem()
}

func (o KeyVaultCredentialSettingsOutput) ToKeyVaultCredentialSettingsOutput() KeyVaultCredentialSettingsOutput {
	return o
}

func (o KeyVaultCredentialSettingsOutput) ToKeyVaultCredentialSettingsOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsOutput {
	return o
}

func (o KeyVaultCredentialSettingsOutput) ToKeyVaultCredentialSettingsPtrOutput() KeyVaultCredentialSettingsPtrOutput {
	return o.ToKeyVaultCredentialSettingsPtrOutputWithContext(context.Background())
}

func (o KeyVaultCredentialSettingsOutput) ToKeyVaultCredentialSettingsPtrOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KeyVaultCredentialSettings) *KeyVaultCredentialSettings {
		return &v
	}).(KeyVaultCredentialSettingsPtrOutput)
}

func (o KeyVaultCredentialSettingsOutput) AzureKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettings) *string { return v.AzureKeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsOutput) CredentialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettings) *string { return v.CredentialName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettings) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o KeyVaultCredentialSettingsOutput) ServicePrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettings) *string { return v.ServicePrincipalName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsOutput) ServicePrincipalSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettings) *string { return v.ServicePrincipalSecret }).(pulumi.StringPtrOutput)
}

type KeyVaultCredentialSettingsPtrOutput struct{ *pulumi.OutputState }

func (KeyVaultCredentialSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCredentialSettings)(nil)).Elem()
}

func (o KeyVaultCredentialSettingsPtrOutput) ToKeyVaultCredentialSettingsPtrOutput() KeyVaultCredentialSettingsPtrOutput {
	return o
}

func (o KeyVaultCredentialSettingsPtrOutput) ToKeyVaultCredentialSettingsPtrOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsPtrOutput {
	return o
}

func (o KeyVaultCredentialSettingsPtrOutput) Elem() KeyVaultCredentialSettingsOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettings) KeyVaultCredentialSettings {
		if v != nil {
			return *v
		}
		var ret KeyVaultCredentialSettings
		return ret
	}).(KeyVaultCredentialSettingsOutput)
}

func (o KeyVaultCredentialSettingsPtrOutput) AzureKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettings) *string {
		if v == nil {
			return nil
		}
		return v.AzureKeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsPtrOutput) CredentialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettings) *string {
		if v == nil {
			return nil
		}
		return v.CredentialName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsPtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettings) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o KeyVaultCredentialSettingsPtrOutput) ServicePrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettings) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsPtrOutput) ServicePrincipalSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettings) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalSecret
	}).(pulumi.StringPtrOutput)
}

type KeyVaultCredentialSettingsResponse struct {
	AzureKeyVaultUrl     *string `pulumi:"azureKeyVaultUrl"`
	CredentialName       *string `pulumi:"credentialName"`
	Enable               *bool   `pulumi:"enable"`
	ServicePrincipalName *string `pulumi:"servicePrincipalName"`
}

type KeyVaultCredentialSettingsResponseOutput struct{ *pulumi.OutputState }

func (KeyVaultCredentialSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KeyVaultCredentialSettingsResponse)(nil)).Elem()
}

func (o KeyVaultCredentialSettingsResponseOutput) ToKeyVaultCredentialSettingsResponseOutput() KeyVaultCredentialSettingsResponseOutput {
	return o
}

func (o KeyVaultCredentialSettingsResponseOutput) ToKeyVaultCredentialSettingsResponseOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsResponseOutput {
	return o
}

func (o KeyVaultCredentialSettingsResponseOutput) AzureKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettingsResponse) *string { return v.AzureKeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsResponseOutput) CredentialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettingsResponse) *string { return v.CredentialName }).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsResponseOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettingsResponse) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o KeyVaultCredentialSettingsResponseOutput) ServicePrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KeyVaultCredentialSettingsResponse) *string { return v.ServicePrincipalName }).(pulumi.StringPtrOutput)
}

type KeyVaultCredentialSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (KeyVaultCredentialSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KeyVaultCredentialSettingsResponse)(nil)).Elem()
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) ToKeyVaultCredentialSettingsResponsePtrOutput() KeyVaultCredentialSettingsResponsePtrOutput {
	return o
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) ToKeyVaultCredentialSettingsResponsePtrOutputWithContext(ctx context.Context) KeyVaultCredentialSettingsResponsePtrOutput {
	return o
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) Elem() KeyVaultCredentialSettingsResponseOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettingsResponse) KeyVaultCredentialSettingsResponse {
		if v != nil {
			return *v
		}
		var ret KeyVaultCredentialSettingsResponse
		return ret
	}).(KeyVaultCredentialSettingsResponseOutput)
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) AzureKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.AzureKeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) CredentialName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.CredentialName
	}).(pulumi.StringPtrOutput)
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o KeyVaultCredentialSettingsResponsePtrOutput) ServicePrincipalName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KeyVaultCredentialSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServicePrincipalName
	}).(pulumi.StringPtrOutput)
}

type LoadBalancerConfiguration struct {
	LoadBalancerResourceId     *string           `pulumi:"loadBalancerResourceId"`
	PrivateIpAddress           *PrivateIPAddress `pulumi:"privateIpAddress"`
	ProbePort                  *int              `pulumi:"probePort"`
	PublicIpAddressResourceId  *string           `pulumi:"publicIpAddressResourceId"`
	SqlVirtualMachineInstances []string          `pulumi:"sqlVirtualMachineInstances"`
}





type LoadBalancerConfigurationInput interface {
	pulumi.Input

	ToLoadBalancerConfigurationOutput() LoadBalancerConfigurationOutput
	ToLoadBalancerConfigurationOutputWithContext(context.Context) LoadBalancerConfigurationOutput
}

type LoadBalancerConfigurationArgs struct {
	LoadBalancerResourceId     pulumi.StringPtrInput    `pulumi:"loadBalancerResourceId"`
	PrivateIpAddress           PrivateIPAddressPtrInput `pulumi:"privateIpAddress"`
	ProbePort                  pulumi.IntPtrInput       `pulumi:"probePort"`
	PublicIpAddressResourceId  pulumi.StringPtrInput    `pulumi:"publicIpAddressResourceId"`
	SqlVirtualMachineInstances pulumi.StringArrayInput  `pulumi:"sqlVirtualMachineInstances"`
}

func (LoadBalancerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfiguration)(nil)).Elem()
}

func (i LoadBalancerConfigurationArgs) ToLoadBalancerConfigurationOutput() LoadBalancerConfigurationOutput {
	return i.ToLoadBalancerConfigurationOutputWithContext(context.Background())
}

func (i LoadBalancerConfigurationArgs) ToLoadBalancerConfigurationOutputWithContext(ctx context.Context) LoadBalancerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigurationOutput)
}





type LoadBalancerConfigurationArrayInput interface {
	pulumi.Input

	ToLoadBalancerConfigurationArrayOutput() LoadBalancerConfigurationArrayOutput
	ToLoadBalancerConfigurationArrayOutputWithContext(context.Context) LoadBalancerConfigurationArrayOutput
}

type LoadBalancerConfigurationArray []LoadBalancerConfigurationInput

func (LoadBalancerConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerConfiguration)(nil)).Elem()
}

func (i LoadBalancerConfigurationArray) ToLoadBalancerConfigurationArrayOutput() LoadBalancerConfigurationArrayOutput {
	return i.ToLoadBalancerConfigurationArrayOutputWithContext(context.Background())
}

func (i LoadBalancerConfigurationArray) ToLoadBalancerConfigurationArrayOutputWithContext(ctx context.Context) LoadBalancerConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LoadBalancerConfigurationArrayOutput)
}

type LoadBalancerConfigurationOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfiguration)(nil)).Elem()
}

func (o LoadBalancerConfigurationOutput) ToLoadBalancerConfigurationOutput() LoadBalancerConfigurationOutput {
	return o
}

func (o LoadBalancerConfigurationOutput) ToLoadBalancerConfigurationOutputWithContext(ctx context.Context) LoadBalancerConfigurationOutput {
	return o
}

func (o LoadBalancerConfigurationOutput) LoadBalancerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) *string { return v.LoadBalancerResourceId }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigurationOutput) PrivateIpAddress() PrivateIPAddressPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) *PrivateIPAddress { return v.PrivateIpAddress }).(PrivateIPAddressPtrOutput)
}

func (o LoadBalancerConfigurationOutput) ProbePort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) *int { return v.ProbePort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerConfigurationOutput) PublicIpAddressResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) *string { return v.PublicIpAddressResourceId }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigurationOutput) SqlVirtualMachineInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoadBalancerConfiguration) []string { return v.SqlVirtualMachineInstances }).(pulumi.StringArrayOutput)
}

type LoadBalancerConfigurationArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerConfiguration)(nil)).Elem()
}

func (o LoadBalancerConfigurationArrayOutput) ToLoadBalancerConfigurationArrayOutput() LoadBalancerConfigurationArrayOutput {
	return o
}

func (o LoadBalancerConfigurationArrayOutput) ToLoadBalancerConfigurationArrayOutputWithContext(ctx context.Context) LoadBalancerConfigurationArrayOutput {
	return o
}

func (o LoadBalancerConfigurationArrayOutput) Index(i pulumi.IntInput) LoadBalancerConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerConfiguration {
		return vs[0].([]LoadBalancerConfiguration)[vs[1].(int)]
	}).(LoadBalancerConfigurationOutput)
}

type LoadBalancerConfigurationResponse struct {
	LoadBalancerResourceId     *string                   `pulumi:"loadBalancerResourceId"`
	PrivateIpAddress           *PrivateIPAddressResponse `pulumi:"privateIpAddress"`
	ProbePort                  *int                      `pulumi:"probePort"`
	PublicIpAddressResourceId  *string                   `pulumi:"publicIpAddressResourceId"`
	SqlVirtualMachineInstances []string                  `pulumi:"sqlVirtualMachineInstances"`
}

type LoadBalancerConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LoadBalancerConfigurationResponse)(nil)).Elem()
}

func (o LoadBalancerConfigurationResponseOutput) ToLoadBalancerConfigurationResponseOutput() LoadBalancerConfigurationResponseOutput {
	return o
}

func (o LoadBalancerConfigurationResponseOutput) ToLoadBalancerConfigurationResponseOutputWithContext(ctx context.Context) LoadBalancerConfigurationResponseOutput {
	return o
}

func (o LoadBalancerConfigurationResponseOutput) LoadBalancerResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) *string { return v.LoadBalancerResourceId }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigurationResponseOutput) PrivateIpAddress() PrivateIPAddressResponsePtrOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) *PrivateIPAddressResponse { return v.PrivateIpAddress }).(PrivateIPAddressResponsePtrOutput)
}

func (o LoadBalancerConfigurationResponseOutput) ProbePort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) *int { return v.ProbePort }).(pulumi.IntPtrOutput)
}

func (o LoadBalancerConfigurationResponseOutput) PublicIpAddressResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) *string { return v.PublicIpAddressResourceId }).(pulumi.StringPtrOutput)
}

func (o LoadBalancerConfigurationResponseOutput) SqlVirtualMachineInstances() pulumi.StringArrayOutput {
	return o.ApplyT(func(v LoadBalancerConfigurationResponse) []string { return v.SqlVirtualMachineInstances }).(pulumi.StringArrayOutput)
}

type LoadBalancerConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (LoadBalancerConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LoadBalancerConfigurationResponse)(nil)).Elem()
}

func (o LoadBalancerConfigurationResponseArrayOutput) ToLoadBalancerConfigurationResponseArrayOutput() LoadBalancerConfigurationResponseArrayOutput {
	return o
}

func (o LoadBalancerConfigurationResponseArrayOutput) ToLoadBalancerConfigurationResponseArrayOutputWithContext(ctx context.Context) LoadBalancerConfigurationResponseArrayOutput {
	return o
}

func (o LoadBalancerConfigurationResponseArrayOutput) Index(i pulumi.IntInput) LoadBalancerConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LoadBalancerConfigurationResponse {
		return vs[0].([]LoadBalancerConfigurationResponse)[vs[1].(int)]
	}).(LoadBalancerConfigurationResponseOutput)
}

type MultiSubnetIpConfiguration struct {
	PrivateIpAddress          PrivateIPAddress `pulumi:"privateIpAddress"`
	SqlVirtualMachineInstance string           `pulumi:"sqlVirtualMachineInstance"`
}





type MultiSubnetIpConfigurationInput interface {
	pulumi.Input

	ToMultiSubnetIpConfigurationOutput() MultiSubnetIpConfigurationOutput
	ToMultiSubnetIpConfigurationOutputWithContext(context.Context) MultiSubnetIpConfigurationOutput
}

type MultiSubnetIpConfigurationArgs struct {
	PrivateIpAddress          PrivateIPAddressInput `pulumi:"privateIpAddress"`
	SqlVirtualMachineInstance pulumi.StringInput    `pulumi:"sqlVirtualMachineInstance"`
}

func (MultiSubnetIpConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiSubnetIpConfiguration)(nil)).Elem()
}

func (i MultiSubnetIpConfigurationArgs) ToMultiSubnetIpConfigurationOutput() MultiSubnetIpConfigurationOutput {
	return i.ToMultiSubnetIpConfigurationOutputWithContext(context.Background())
}

func (i MultiSubnetIpConfigurationArgs) ToMultiSubnetIpConfigurationOutputWithContext(ctx context.Context) MultiSubnetIpConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiSubnetIpConfigurationOutput)
}





type MultiSubnetIpConfigurationArrayInput interface {
	pulumi.Input

	ToMultiSubnetIpConfigurationArrayOutput() MultiSubnetIpConfigurationArrayOutput
	ToMultiSubnetIpConfigurationArrayOutputWithContext(context.Context) MultiSubnetIpConfigurationArrayOutput
}

type MultiSubnetIpConfigurationArray []MultiSubnetIpConfigurationInput

func (MultiSubnetIpConfigurationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MultiSubnetIpConfiguration)(nil)).Elem()
}

func (i MultiSubnetIpConfigurationArray) ToMultiSubnetIpConfigurationArrayOutput() MultiSubnetIpConfigurationArrayOutput {
	return i.ToMultiSubnetIpConfigurationArrayOutputWithContext(context.Background())
}

func (i MultiSubnetIpConfigurationArray) ToMultiSubnetIpConfigurationArrayOutputWithContext(ctx context.Context) MultiSubnetIpConfigurationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MultiSubnetIpConfigurationArrayOutput)
}

type MultiSubnetIpConfigurationOutput struct{ *pulumi.OutputState }

func (MultiSubnetIpConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiSubnetIpConfiguration)(nil)).Elem()
}

func (o MultiSubnetIpConfigurationOutput) ToMultiSubnetIpConfigurationOutput() MultiSubnetIpConfigurationOutput {
	return o
}

func (o MultiSubnetIpConfigurationOutput) ToMultiSubnetIpConfigurationOutputWithContext(ctx context.Context) MultiSubnetIpConfigurationOutput {
	return o
}

func (o MultiSubnetIpConfigurationOutput) PrivateIpAddress() PrivateIPAddressOutput {
	return o.ApplyT(func(v MultiSubnetIpConfiguration) PrivateIPAddress { return v.PrivateIpAddress }).(PrivateIPAddressOutput)
}

func (o MultiSubnetIpConfigurationOutput) SqlVirtualMachineInstance() pulumi.StringOutput {
	return o.ApplyT(func(v MultiSubnetIpConfiguration) string { return v.SqlVirtualMachineInstance }).(pulumi.StringOutput)
}

type MultiSubnetIpConfigurationArrayOutput struct{ *pulumi.OutputState }

func (MultiSubnetIpConfigurationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MultiSubnetIpConfiguration)(nil)).Elem()
}

func (o MultiSubnetIpConfigurationArrayOutput) ToMultiSubnetIpConfigurationArrayOutput() MultiSubnetIpConfigurationArrayOutput {
	return o
}

func (o MultiSubnetIpConfigurationArrayOutput) ToMultiSubnetIpConfigurationArrayOutputWithContext(ctx context.Context) MultiSubnetIpConfigurationArrayOutput {
	return o
}

func (o MultiSubnetIpConfigurationArrayOutput) Index(i pulumi.IntInput) MultiSubnetIpConfigurationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MultiSubnetIpConfiguration {
		return vs[0].([]MultiSubnetIpConfiguration)[vs[1].(int)]
	}).(MultiSubnetIpConfigurationOutput)
}

type MultiSubnetIpConfigurationResponse struct {
	PrivateIpAddress          PrivateIPAddressResponse `pulumi:"privateIpAddress"`
	SqlVirtualMachineInstance string                   `pulumi:"sqlVirtualMachineInstance"`
}

type MultiSubnetIpConfigurationResponseOutput struct{ *pulumi.OutputState }

func (MultiSubnetIpConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MultiSubnetIpConfigurationResponse)(nil)).Elem()
}

func (o MultiSubnetIpConfigurationResponseOutput) ToMultiSubnetIpConfigurationResponseOutput() MultiSubnetIpConfigurationResponseOutput {
	return o
}

func (o MultiSubnetIpConfigurationResponseOutput) ToMultiSubnetIpConfigurationResponseOutputWithContext(ctx context.Context) MultiSubnetIpConfigurationResponseOutput {
	return o
}

func (o MultiSubnetIpConfigurationResponseOutput) PrivateIpAddress() PrivateIPAddressResponseOutput {
	return o.ApplyT(func(v MultiSubnetIpConfigurationResponse) PrivateIPAddressResponse { return v.PrivateIpAddress }).(PrivateIPAddressResponseOutput)
}

func (o MultiSubnetIpConfigurationResponseOutput) SqlVirtualMachineInstance() pulumi.StringOutput {
	return o.ApplyT(func(v MultiSubnetIpConfigurationResponse) string { return v.SqlVirtualMachineInstance }).(pulumi.StringOutput)
}

type MultiSubnetIpConfigurationResponseArrayOutput struct{ *pulumi.OutputState }

func (MultiSubnetIpConfigurationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]MultiSubnetIpConfigurationResponse)(nil)).Elem()
}

func (o MultiSubnetIpConfigurationResponseArrayOutput) ToMultiSubnetIpConfigurationResponseArrayOutput() MultiSubnetIpConfigurationResponseArrayOutput {
	return o
}

func (o MultiSubnetIpConfigurationResponseArrayOutput) ToMultiSubnetIpConfigurationResponseArrayOutputWithContext(ctx context.Context) MultiSubnetIpConfigurationResponseArrayOutput {
	return o
}

func (o MultiSubnetIpConfigurationResponseArrayOutput) Index(i pulumi.IntInput) MultiSubnetIpConfigurationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) MultiSubnetIpConfigurationResponse {
		return vs[0].([]MultiSubnetIpConfigurationResponse)[vs[1].(int)]
	}).(MultiSubnetIpConfigurationResponseOutput)
}

type PrivateIPAddress struct {
	IpAddress        *string `pulumi:"ipAddress"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
}





type PrivateIPAddressInput interface {
	pulumi.Input

	ToPrivateIPAddressOutput() PrivateIPAddressOutput
	ToPrivateIPAddressOutputWithContext(context.Context) PrivateIPAddressOutput
}

type PrivateIPAddressArgs struct {
	IpAddress        pulumi.StringPtrInput `pulumi:"ipAddress"`
	SubnetResourceId pulumi.StringPtrInput `pulumi:"subnetResourceId"`
}

func (PrivateIPAddressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateIPAddress)(nil)).Elem()
}

func (i PrivateIPAddressArgs) ToPrivateIPAddressOutput() PrivateIPAddressOutput {
	return i.ToPrivateIPAddressOutputWithContext(context.Background())
}

func (i PrivateIPAddressArgs) ToPrivateIPAddressOutputWithContext(ctx context.Context) PrivateIPAddressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateIPAddressOutput)
}

func (i PrivateIPAddressArgs) ToPrivateIPAddressPtrOutput() PrivateIPAddressPtrOutput {
	return i.ToPrivateIPAddressPtrOutputWithContext(context.Background())
}

func (i PrivateIPAddressArgs) ToPrivateIPAddressPtrOutputWithContext(ctx context.Context) PrivateIPAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateIPAddressOutput).ToPrivateIPAddressPtrOutputWithContext(ctx)
}









type PrivateIPAddressPtrInput interface {
	pulumi.Input

	ToPrivateIPAddressPtrOutput() PrivateIPAddressPtrOutput
	ToPrivateIPAddressPtrOutputWithContext(context.Context) PrivateIPAddressPtrOutput
}

type privateIPAddressPtrType PrivateIPAddressArgs

func PrivateIPAddressPtr(v *PrivateIPAddressArgs) PrivateIPAddressPtrInput {
	return (*privateIPAddressPtrType)(v)
}

func (*privateIPAddressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateIPAddress)(nil)).Elem()
}

func (i *privateIPAddressPtrType) ToPrivateIPAddressPtrOutput() PrivateIPAddressPtrOutput {
	return i.ToPrivateIPAddressPtrOutputWithContext(context.Background())
}

func (i *privateIPAddressPtrType) ToPrivateIPAddressPtrOutputWithContext(ctx context.Context) PrivateIPAddressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateIPAddressPtrOutput)
}

type PrivateIPAddressOutput struct{ *pulumi.OutputState }

func (PrivateIPAddressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateIPAddress)(nil)).Elem()
}

func (o PrivateIPAddressOutput) ToPrivateIPAddressOutput() PrivateIPAddressOutput {
	return o
}

func (o PrivateIPAddressOutput) ToPrivateIPAddressOutputWithContext(ctx context.Context) PrivateIPAddressOutput {
	return o
}

func (o PrivateIPAddressOutput) ToPrivateIPAddressPtrOutput() PrivateIPAddressPtrOutput {
	return o.ToPrivateIPAddressPtrOutputWithContext(context.Background())
}

func (o PrivateIPAddressOutput) ToPrivateIPAddressPtrOutputWithContext(ctx context.Context) PrivateIPAddressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateIPAddress) *PrivateIPAddress {
		return &v
	}).(PrivateIPAddressPtrOutput)
}

func (o PrivateIPAddressOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateIPAddress) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PrivateIPAddressOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateIPAddress) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

type PrivateIPAddressPtrOutput struct{ *pulumi.OutputState }

func (PrivateIPAddressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateIPAddress)(nil)).Elem()
}

func (o PrivateIPAddressPtrOutput) ToPrivateIPAddressPtrOutput() PrivateIPAddressPtrOutput {
	return o
}

func (o PrivateIPAddressPtrOutput) ToPrivateIPAddressPtrOutputWithContext(ctx context.Context) PrivateIPAddressPtrOutput {
	return o
}

func (o PrivateIPAddressPtrOutput) Elem() PrivateIPAddressOutput {
	return o.ApplyT(func(v *PrivateIPAddress) PrivateIPAddress {
		if v != nil {
			return *v
		}
		var ret PrivateIPAddress
		return ret
	}).(PrivateIPAddressOutput)
}

func (o PrivateIPAddressPtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateIPAddress) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o PrivateIPAddressPtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateIPAddress) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

type PrivateIPAddressResponse struct {
	IpAddress        *string `pulumi:"ipAddress"`
	SubnetResourceId *string `pulumi:"subnetResourceId"`
}

type PrivateIPAddressResponseOutput struct{ *pulumi.OutputState }

func (PrivateIPAddressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateIPAddressResponse)(nil)).Elem()
}

func (o PrivateIPAddressResponseOutput) ToPrivateIPAddressResponseOutput() PrivateIPAddressResponseOutput {
	return o
}

func (o PrivateIPAddressResponseOutput) ToPrivateIPAddressResponseOutputWithContext(ctx context.Context) PrivateIPAddressResponseOutput {
	return o
}

func (o PrivateIPAddressResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateIPAddressResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o PrivateIPAddressResponseOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateIPAddressResponse) *string { return v.SubnetResourceId }).(pulumi.StringPtrOutput)
}

type PrivateIPAddressResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateIPAddressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateIPAddressResponse)(nil)).Elem()
}

func (o PrivateIPAddressResponsePtrOutput) ToPrivateIPAddressResponsePtrOutput() PrivateIPAddressResponsePtrOutput {
	return o
}

func (o PrivateIPAddressResponsePtrOutput) ToPrivateIPAddressResponsePtrOutputWithContext(ctx context.Context) PrivateIPAddressResponsePtrOutput {
	return o
}

func (o PrivateIPAddressResponsePtrOutput) Elem() PrivateIPAddressResponseOutput {
	return o.ApplyT(func(v *PrivateIPAddressResponse) PrivateIPAddressResponse {
		if v != nil {
			return *v
		}
		var ret PrivateIPAddressResponse
		return ret
	}).(PrivateIPAddressResponseOutput)
}

func (o PrivateIPAddressResponsePtrOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.IpAddress
	}).(pulumi.StringPtrOutput)
}

func (o PrivateIPAddressResponsePtrOutput) SubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateIPAddressResponse) *string {
		if v == nil {
			return nil
		}
		return v.SubnetResourceId
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentity struct {
	Type *string `pulumi:"type"`
}





type ResourceIdentityInput interface {
	pulumi.Input

	ToResourceIdentityOutput() ResourceIdentityOutput
	ToResourceIdentityOutputWithContext(context.Context) ResourceIdentityOutput
}

type ResourceIdentityArgs struct {
	Type pulumi.StringPtrInput `pulumi:"type"`
}

func (ResourceIdentityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (i ResourceIdentityArgs) ToResourceIdentityOutput() ResourceIdentityOutput {
	return i.ToResourceIdentityOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput)
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i ResourceIdentityArgs) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityOutput).ToResourceIdentityPtrOutputWithContext(ctx)
}









type ResourceIdentityPtrInput interface {
	pulumi.Input

	ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput
	ToResourceIdentityPtrOutputWithContext(context.Context) ResourceIdentityPtrOutput
}

type resourceIdentityPtrType ResourceIdentityArgs

func ResourceIdentityPtr(v *ResourceIdentityArgs) ResourceIdentityPtrInput {
	return (*resourceIdentityPtrType)(v)
}

func (*resourceIdentityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return i.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (i *resourceIdentityPtrType) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceIdentityPtrOutput)
}

type ResourceIdentityOutput struct{ *pulumi.OutputState }

func (ResourceIdentityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityOutput) ToResourceIdentityOutput() ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityOutputWithContext(ctx context.Context) ResourceIdentityOutput {
	return o
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o.ToResourceIdentityPtrOutputWithContext(context.Background())
}

func (o ResourceIdentityOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceIdentity) *ResourceIdentity {
		return &v
	}).(ResourceIdentityPtrOutput)
}

func (o ResourceIdentityOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentity) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityPtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentity)(nil)).Elem()
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutput() ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) ToResourceIdentityPtrOutputWithContext(ctx context.Context) ResourceIdentityPtrOutput {
	return o
}

func (o ResourceIdentityPtrOutput) Elem() ResourceIdentityOutput {
	return o.ApplyT(func(v *ResourceIdentity) ResourceIdentity {
		if v != nil {
			return *v
		}
		var ret ResourceIdentity
		return ret
	}).(ResourceIdentityOutput)
}

func (o ResourceIdentityPtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentity) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponse struct {
	PrincipalId string  `pulumi:"principalId"`
	TenantId    string  `pulumi:"tenantId"`
	Type        *string `pulumi:"type"`
}

type ResourceIdentityResponseOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutput() ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) ToResourceIdentityResponseOutputWithContext(ctx context.Context) ResourceIdentityResponseOutput {
	return o
}

func (o ResourceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ResourceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResourceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ResourceIdentityResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceIdentityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceIdentityResponse)(nil)).Elem()
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutput() ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) ToResourceIdentityResponsePtrOutputWithContext(ctx context.Context) ResourceIdentityResponsePtrOutput {
	return o
}

func (o ResourceIdentityResponsePtrOutput) Elem() ResourceIdentityResponseOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) ResourceIdentityResponse {
		if v != nil {
			return *v
		}
		var ret ResourceIdentityResponse
		return ret
	}).(ResourceIdentityResponseOutput)
}

func (o ResourceIdentityResponsePtrOutput) PrincipalId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PrincipalId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TenantId
	}).(pulumi.StringPtrOutput)
}

func (o ResourceIdentityResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ResourceIdentityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type SQLInstanceSettings struct {
	Collation                          *string `pulumi:"collation"`
	IsIfiEnabled                       *bool   `pulumi:"isIfiEnabled"`
	IsLpimEnabled                      *bool   `pulumi:"isLpimEnabled"`
	IsOptimizeForAdHocWorkloadsEnabled *bool   `pulumi:"isOptimizeForAdHocWorkloadsEnabled"`
	MaxDop                             *int    `pulumi:"maxDop"`
	MaxServerMemoryMB                  *int    `pulumi:"maxServerMemoryMB"`
	MinServerMemoryMB                  *int    `pulumi:"minServerMemoryMB"`
}





type SQLInstanceSettingsInput interface {
	pulumi.Input

	ToSQLInstanceSettingsOutput() SQLInstanceSettingsOutput
	ToSQLInstanceSettingsOutputWithContext(context.Context) SQLInstanceSettingsOutput
}

type SQLInstanceSettingsArgs struct {
	Collation                          pulumi.StringPtrInput `pulumi:"collation"`
	IsIfiEnabled                       pulumi.BoolPtrInput   `pulumi:"isIfiEnabled"`
	IsLpimEnabled                      pulumi.BoolPtrInput   `pulumi:"isLpimEnabled"`
	IsOptimizeForAdHocWorkloadsEnabled pulumi.BoolPtrInput   `pulumi:"isOptimizeForAdHocWorkloadsEnabled"`
	MaxDop                             pulumi.IntPtrInput    `pulumi:"maxDop"`
	MaxServerMemoryMB                  pulumi.IntPtrInput    `pulumi:"maxServerMemoryMB"`
	MinServerMemoryMB                  pulumi.IntPtrInput    `pulumi:"minServerMemoryMB"`
}

func (SQLInstanceSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLInstanceSettings)(nil)).Elem()
}

func (i SQLInstanceSettingsArgs) ToSQLInstanceSettingsOutput() SQLInstanceSettingsOutput {
	return i.ToSQLInstanceSettingsOutputWithContext(context.Background())
}

func (i SQLInstanceSettingsArgs) ToSQLInstanceSettingsOutputWithContext(ctx context.Context) SQLInstanceSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLInstanceSettingsOutput)
}

func (i SQLInstanceSettingsArgs) ToSQLInstanceSettingsPtrOutput() SQLInstanceSettingsPtrOutput {
	return i.ToSQLInstanceSettingsPtrOutputWithContext(context.Background())
}

func (i SQLInstanceSettingsArgs) ToSQLInstanceSettingsPtrOutputWithContext(ctx context.Context) SQLInstanceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLInstanceSettingsOutput).ToSQLInstanceSettingsPtrOutputWithContext(ctx)
}









type SQLInstanceSettingsPtrInput interface {
	pulumi.Input

	ToSQLInstanceSettingsPtrOutput() SQLInstanceSettingsPtrOutput
	ToSQLInstanceSettingsPtrOutputWithContext(context.Context) SQLInstanceSettingsPtrOutput
}

type sqlinstanceSettingsPtrType SQLInstanceSettingsArgs

func SQLInstanceSettingsPtr(v *SQLInstanceSettingsArgs) SQLInstanceSettingsPtrInput {
	return (*sqlinstanceSettingsPtrType)(v)
}

func (*sqlinstanceSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLInstanceSettings)(nil)).Elem()
}

func (i *sqlinstanceSettingsPtrType) ToSQLInstanceSettingsPtrOutput() SQLInstanceSettingsPtrOutput {
	return i.ToSQLInstanceSettingsPtrOutputWithContext(context.Background())
}

func (i *sqlinstanceSettingsPtrType) ToSQLInstanceSettingsPtrOutputWithContext(ctx context.Context) SQLInstanceSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLInstanceSettingsPtrOutput)
}

type SQLInstanceSettingsOutput struct{ *pulumi.OutputState }

func (SQLInstanceSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLInstanceSettings)(nil)).Elem()
}

func (o SQLInstanceSettingsOutput) ToSQLInstanceSettingsOutput() SQLInstanceSettingsOutput {
	return o
}

func (o SQLInstanceSettingsOutput) ToSQLInstanceSettingsOutputWithContext(ctx context.Context) SQLInstanceSettingsOutput {
	return o
}

func (o SQLInstanceSettingsOutput) ToSQLInstanceSettingsPtrOutput() SQLInstanceSettingsPtrOutput {
	return o.ToSQLInstanceSettingsPtrOutputWithContext(context.Background())
}

func (o SQLInstanceSettingsOutput) ToSQLInstanceSettingsPtrOutputWithContext(ctx context.Context) SQLInstanceSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SQLInstanceSettings) *SQLInstanceSettings {
		return &v
	}).(SQLInstanceSettingsPtrOutput)
}

func (o SQLInstanceSettingsOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SQLInstanceSettingsOutput) IsIfiEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *bool { return v.IsIfiEnabled }).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsOutput) IsLpimEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *bool { return v.IsLpimEnabled }).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsOutput) IsOptimizeForAdHocWorkloadsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *bool { return v.IsOptimizeForAdHocWorkloadsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsOutput) MaxDop() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *int { return v.MaxDop }).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsOutput) MaxServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *int { return v.MaxServerMemoryMB }).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsOutput) MinServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettings) *int { return v.MinServerMemoryMB }).(pulumi.IntPtrOutput)
}

type SQLInstanceSettingsPtrOutput struct{ *pulumi.OutputState }

func (SQLInstanceSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLInstanceSettings)(nil)).Elem()
}

func (o SQLInstanceSettingsPtrOutput) ToSQLInstanceSettingsPtrOutput() SQLInstanceSettingsPtrOutput {
	return o
}

func (o SQLInstanceSettingsPtrOutput) ToSQLInstanceSettingsPtrOutputWithContext(ctx context.Context) SQLInstanceSettingsPtrOutput {
	return o
}

func (o SQLInstanceSettingsPtrOutput) Elem() SQLInstanceSettingsOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) SQLInstanceSettings {
		if v != nil {
			return *v
		}
		var ret SQLInstanceSettings
		return ret
	}).(SQLInstanceSettingsOutput)
}

func (o SQLInstanceSettingsPtrOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *string {
		if v == nil {
			return nil
		}
		return v.Collation
	}).(pulumi.StringPtrOutput)
}

func (o SQLInstanceSettingsPtrOutput) IsIfiEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IsIfiEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsPtrOutput) IsLpimEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IsLpimEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsPtrOutput) IsOptimizeForAdHocWorkloadsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *bool {
		if v == nil {
			return nil
		}
		return v.IsOptimizeForAdHocWorkloadsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsPtrOutput) MaxDop() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *int {
		if v == nil {
			return nil
		}
		return v.MaxDop
	}).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsPtrOutput) MaxServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *int {
		if v == nil {
			return nil
		}
		return v.MaxServerMemoryMB
	}).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsPtrOutput) MinServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettings) *int {
		if v == nil {
			return nil
		}
		return v.MinServerMemoryMB
	}).(pulumi.IntPtrOutput)
}

type SQLInstanceSettingsResponse struct {
	Collation                          *string `pulumi:"collation"`
	IsIfiEnabled                       *bool   `pulumi:"isIfiEnabled"`
	IsLpimEnabled                      *bool   `pulumi:"isLpimEnabled"`
	IsOptimizeForAdHocWorkloadsEnabled *bool   `pulumi:"isOptimizeForAdHocWorkloadsEnabled"`
	MaxDop                             *int    `pulumi:"maxDop"`
	MaxServerMemoryMB                  *int    `pulumi:"maxServerMemoryMB"`
	MinServerMemoryMB                  *int    `pulumi:"minServerMemoryMB"`
}

type SQLInstanceSettingsResponseOutput struct{ *pulumi.OutputState }

func (SQLInstanceSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLInstanceSettingsResponse)(nil)).Elem()
}

func (o SQLInstanceSettingsResponseOutput) ToSQLInstanceSettingsResponseOutput() SQLInstanceSettingsResponseOutput {
	return o
}

func (o SQLInstanceSettingsResponseOutput) ToSQLInstanceSettingsResponseOutputWithContext(ctx context.Context) SQLInstanceSettingsResponseOutput {
	return o
}

func (o SQLInstanceSettingsResponseOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *string { return v.Collation }).(pulumi.StringPtrOutput)
}

func (o SQLInstanceSettingsResponseOutput) IsIfiEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *bool { return v.IsIfiEnabled }).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsResponseOutput) IsLpimEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *bool { return v.IsLpimEnabled }).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsResponseOutput) IsOptimizeForAdHocWorkloadsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *bool { return v.IsOptimizeForAdHocWorkloadsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsResponseOutput) MaxDop() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *int { return v.MaxDop }).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsResponseOutput) MaxServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *int { return v.MaxServerMemoryMB }).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsResponseOutput) MinServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLInstanceSettingsResponse) *int { return v.MinServerMemoryMB }).(pulumi.IntPtrOutput)
}

type SQLInstanceSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SQLInstanceSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLInstanceSettingsResponse)(nil)).Elem()
}

func (o SQLInstanceSettingsResponsePtrOutput) ToSQLInstanceSettingsResponsePtrOutput() SQLInstanceSettingsResponsePtrOutput {
	return o
}

func (o SQLInstanceSettingsResponsePtrOutput) ToSQLInstanceSettingsResponsePtrOutputWithContext(ctx context.Context) SQLInstanceSettingsResponsePtrOutput {
	return o
}

func (o SQLInstanceSettingsResponsePtrOutput) Elem() SQLInstanceSettingsResponseOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) SQLInstanceSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SQLInstanceSettingsResponse
		return ret
	}).(SQLInstanceSettingsResponseOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) Collation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Collation
	}).(pulumi.StringPtrOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) IsIfiEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsIfiEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) IsLpimEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsLpimEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) IsOptimizeForAdHocWorkloadsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsOptimizeForAdHocWorkloadsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) MaxDop() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxDop
	}).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) MaxServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxServerMemoryMB
	}).(pulumi.IntPtrOutput)
}

func (o SQLInstanceSettingsResponsePtrOutput) MinServerMemoryMB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLInstanceSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinServerMemoryMB
	}).(pulumi.IntPtrOutput)
}

type SQLStorageSettings struct {
	DefaultFilePath *string `pulumi:"defaultFilePath"`
	Luns            []int   `pulumi:"luns"`
}





type SQLStorageSettingsInput interface {
	pulumi.Input

	ToSQLStorageSettingsOutput() SQLStorageSettingsOutput
	ToSQLStorageSettingsOutputWithContext(context.Context) SQLStorageSettingsOutput
}

type SQLStorageSettingsArgs struct {
	DefaultFilePath pulumi.StringPtrInput `pulumi:"defaultFilePath"`
	Luns            pulumi.IntArrayInput  `pulumi:"luns"`
}

func (SQLStorageSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLStorageSettings)(nil)).Elem()
}

func (i SQLStorageSettingsArgs) ToSQLStorageSettingsOutput() SQLStorageSettingsOutput {
	return i.ToSQLStorageSettingsOutputWithContext(context.Background())
}

func (i SQLStorageSettingsArgs) ToSQLStorageSettingsOutputWithContext(ctx context.Context) SQLStorageSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLStorageSettingsOutput)
}

func (i SQLStorageSettingsArgs) ToSQLStorageSettingsPtrOutput() SQLStorageSettingsPtrOutput {
	return i.ToSQLStorageSettingsPtrOutputWithContext(context.Background())
}

func (i SQLStorageSettingsArgs) ToSQLStorageSettingsPtrOutputWithContext(ctx context.Context) SQLStorageSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLStorageSettingsOutput).ToSQLStorageSettingsPtrOutputWithContext(ctx)
}









type SQLStorageSettingsPtrInput interface {
	pulumi.Input

	ToSQLStorageSettingsPtrOutput() SQLStorageSettingsPtrOutput
	ToSQLStorageSettingsPtrOutputWithContext(context.Context) SQLStorageSettingsPtrOutput
}

type sqlstorageSettingsPtrType SQLStorageSettingsArgs

func SQLStorageSettingsPtr(v *SQLStorageSettingsArgs) SQLStorageSettingsPtrInput {
	return (*sqlstorageSettingsPtrType)(v)
}

func (*sqlstorageSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLStorageSettings)(nil)).Elem()
}

func (i *sqlstorageSettingsPtrType) ToSQLStorageSettingsPtrOutput() SQLStorageSettingsPtrOutput {
	return i.ToSQLStorageSettingsPtrOutputWithContext(context.Background())
}

func (i *sqlstorageSettingsPtrType) ToSQLStorageSettingsPtrOutputWithContext(ctx context.Context) SQLStorageSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLStorageSettingsPtrOutput)
}

type SQLStorageSettingsOutput struct{ *pulumi.OutputState }

func (SQLStorageSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLStorageSettings)(nil)).Elem()
}

func (o SQLStorageSettingsOutput) ToSQLStorageSettingsOutput() SQLStorageSettingsOutput {
	return o
}

func (o SQLStorageSettingsOutput) ToSQLStorageSettingsOutputWithContext(ctx context.Context) SQLStorageSettingsOutput {
	return o
}

func (o SQLStorageSettingsOutput) ToSQLStorageSettingsPtrOutput() SQLStorageSettingsPtrOutput {
	return o.ToSQLStorageSettingsPtrOutputWithContext(context.Background())
}

func (o SQLStorageSettingsOutput) ToSQLStorageSettingsPtrOutputWithContext(ctx context.Context) SQLStorageSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SQLStorageSettings) *SQLStorageSettings {
		return &v
	}).(SQLStorageSettingsPtrOutput)
}

func (o SQLStorageSettingsOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLStorageSettings) *string { return v.DefaultFilePath }).(pulumi.StringPtrOutput)
}

func (o SQLStorageSettingsOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SQLStorageSettings) []int { return v.Luns }).(pulumi.IntArrayOutput)
}

type SQLStorageSettingsPtrOutput struct{ *pulumi.OutputState }

func (SQLStorageSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLStorageSettings)(nil)).Elem()
}

func (o SQLStorageSettingsPtrOutput) ToSQLStorageSettingsPtrOutput() SQLStorageSettingsPtrOutput {
	return o
}

func (o SQLStorageSettingsPtrOutput) ToSQLStorageSettingsPtrOutputWithContext(ctx context.Context) SQLStorageSettingsPtrOutput {
	return o
}

func (o SQLStorageSettingsPtrOutput) Elem() SQLStorageSettingsOutput {
	return o.ApplyT(func(v *SQLStorageSettings) SQLStorageSettings {
		if v != nil {
			return *v
		}
		var ret SQLStorageSettings
		return ret
	}).(SQLStorageSettingsOutput)
}

func (o SQLStorageSettingsPtrOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLStorageSettings) *string {
		if v == nil {
			return nil
		}
		return v.DefaultFilePath
	}).(pulumi.StringPtrOutput)
}

func (o SQLStorageSettingsPtrOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SQLStorageSettings) []int {
		if v == nil {
			return nil
		}
		return v.Luns
	}).(pulumi.IntArrayOutput)
}

type SQLStorageSettingsResponse struct {
	DefaultFilePath *string `pulumi:"defaultFilePath"`
	Luns            []int   `pulumi:"luns"`
}

type SQLStorageSettingsResponseOutput struct{ *pulumi.OutputState }

func (SQLStorageSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLStorageSettingsResponse)(nil)).Elem()
}

func (o SQLStorageSettingsResponseOutput) ToSQLStorageSettingsResponseOutput() SQLStorageSettingsResponseOutput {
	return o
}

func (o SQLStorageSettingsResponseOutput) ToSQLStorageSettingsResponseOutputWithContext(ctx context.Context) SQLStorageSettingsResponseOutput {
	return o
}

func (o SQLStorageSettingsResponseOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLStorageSettingsResponse) *string { return v.DefaultFilePath }).(pulumi.StringPtrOutput)
}

func (o SQLStorageSettingsResponseOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SQLStorageSettingsResponse) []int { return v.Luns }).(pulumi.IntArrayOutput)
}

type SQLStorageSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SQLStorageSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLStorageSettingsResponse)(nil)).Elem()
}

func (o SQLStorageSettingsResponsePtrOutput) ToSQLStorageSettingsResponsePtrOutput() SQLStorageSettingsResponsePtrOutput {
	return o
}

func (o SQLStorageSettingsResponsePtrOutput) ToSQLStorageSettingsResponsePtrOutputWithContext(ctx context.Context) SQLStorageSettingsResponsePtrOutput {
	return o
}

func (o SQLStorageSettingsResponsePtrOutput) Elem() SQLStorageSettingsResponseOutput {
	return o.ApplyT(func(v *SQLStorageSettingsResponse) SQLStorageSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SQLStorageSettingsResponse
		return ret
	}).(SQLStorageSettingsResponseOutput)
}

func (o SQLStorageSettingsResponsePtrOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLStorageSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultFilePath
	}).(pulumi.StringPtrOutput)
}

func (o SQLStorageSettingsResponsePtrOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SQLStorageSettingsResponse) []int {
		if v == nil {
			return nil
		}
		return v.Luns
	}).(pulumi.IntArrayOutput)
}

type SQLTempDbSettings struct {
	DataFileCount     *int    `pulumi:"dataFileCount"`
	DataFileSize      *int    `pulumi:"dataFileSize"`
	DataGrowth        *int    `pulumi:"dataGrowth"`
	DefaultFilePath   *string `pulumi:"defaultFilePath"`
	LogFileSize       *int    `pulumi:"logFileSize"`
	LogGrowth         *int    `pulumi:"logGrowth"`
	Luns              []int   `pulumi:"luns"`
	PersistFolder     *bool   `pulumi:"persistFolder"`
	PersistFolderPath *string `pulumi:"persistFolderPath"`
}





type SQLTempDbSettingsInput interface {
	pulumi.Input

	ToSQLTempDbSettingsOutput() SQLTempDbSettingsOutput
	ToSQLTempDbSettingsOutputWithContext(context.Context) SQLTempDbSettingsOutput
}

type SQLTempDbSettingsArgs struct {
	DataFileCount     pulumi.IntPtrInput    `pulumi:"dataFileCount"`
	DataFileSize      pulumi.IntPtrInput    `pulumi:"dataFileSize"`
	DataGrowth        pulumi.IntPtrInput    `pulumi:"dataGrowth"`
	DefaultFilePath   pulumi.StringPtrInput `pulumi:"defaultFilePath"`
	LogFileSize       pulumi.IntPtrInput    `pulumi:"logFileSize"`
	LogGrowth         pulumi.IntPtrInput    `pulumi:"logGrowth"`
	Luns              pulumi.IntArrayInput  `pulumi:"luns"`
	PersistFolder     pulumi.BoolPtrInput   `pulumi:"persistFolder"`
	PersistFolderPath pulumi.StringPtrInput `pulumi:"persistFolderPath"`
}

func (SQLTempDbSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLTempDbSettings)(nil)).Elem()
}

func (i SQLTempDbSettingsArgs) ToSQLTempDbSettingsOutput() SQLTempDbSettingsOutput {
	return i.ToSQLTempDbSettingsOutputWithContext(context.Background())
}

func (i SQLTempDbSettingsArgs) ToSQLTempDbSettingsOutputWithContext(ctx context.Context) SQLTempDbSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLTempDbSettingsOutput)
}

func (i SQLTempDbSettingsArgs) ToSQLTempDbSettingsPtrOutput() SQLTempDbSettingsPtrOutput {
	return i.ToSQLTempDbSettingsPtrOutputWithContext(context.Background())
}

func (i SQLTempDbSettingsArgs) ToSQLTempDbSettingsPtrOutputWithContext(ctx context.Context) SQLTempDbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLTempDbSettingsOutput).ToSQLTempDbSettingsPtrOutputWithContext(ctx)
}









type SQLTempDbSettingsPtrInput interface {
	pulumi.Input

	ToSQLTempDbSettingsPtrOutput() SQLTempDbSettingsPtrOutput
	ToSQLTempDbSettingsPtrOutputWithContext(context.Context) SQLTempDbSettingsPtrOutput
}

type sqltempDbSettingsPtrType SQLTempDbSettingsArgs

func SQLTempDbSettingsPtr(v *SQLTempDbSettingsArgs) SQLTempDbSettingsPtrInput {
	return (*sqltempDbSettingsPtrType)(v)
}

func (*sqltempDbSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLTempDbSettings)(nil)).Elem()
}

func (i *sqltempDbSettingsPtrType) ToSQLTempDbSettingsPtrOutput() SQLTempDbSettingsPtrOutput {
	return i.ToSQLTempDbSettingsPtrOutputWithContext(context.Background())
}

func (i *sqltempDbSettingsPtrType) ToSQLTempDbSettingsPtrOutputWithContext(ctx context.Context) SQLTempDbSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SQLTempDbSettingsPtrOutput)
}

type SQLTempDbSettingsOutput struct{ *pulumi.OutputState }

func (SQLTempDbSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLTempDbSettings)(nil)).Elem()
}

func (o SQLTempDbSettingsOutput) ToSQLTempDbSettingsOutput() SQLTempDbSettingsOutput {
	return o
}

func (o SQLTempDbSettingsOutput) ToSQLTempDbSettingsOutputWithContext(ctx context.Context) SQLTempDbSettingsOutput {
	return o
}

func (o SQLTempDbSettingsOutput) ToSQLTempDbSettingsPtrOutput() SQLTempDbSettingsPtrOutput {
	return o.ToSQLTempDbSettingsPtrOutputWithContext(context.Background())
}

func (o SQLTempDbSettingsOutput) ToSQLTempDbSettingsPtrOutputWithContext(ctx context.Context) SQLTempDbSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SQLTempDbSettings) *SQLTempDbSettings {
		return &v
	}).(SQLTempDbSettingsPtrOutput)
}

func (o SQLTempDbSettingsOutput) DataFileCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *int { return v.DataFileCount }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsOutput) DataFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *int { return v.DataFileSize }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsOutput) DataGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *int { return v.DataGrowth }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *string { return v.DefaultFilePath }).(pulumi.StringPtrOutput)
}

func (o SQLTempDbSettingsOutput) LogFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *int { return v.LogFileSize }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsOutput) LogGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *int { return v.LogGrowth }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SQLTempDbSettings) []int { return v.Luns }).(pulumi.IntArrayOutput)
}

func (o SQLTempDbSettingsOutput) PersistFolder() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *bool { return v.PersistFolder }).(pulumi.BoolPtrOutput)
}

func (o SQLTempDbSettingsOutput) PersistFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettings) *string { return v.PersistFolderPath }).(pulumi.StringPtrOutput)
}

type SQLTempDbSettingsPtrOutput struct{ *pulumi.OutputState }

func (SQLTempDbSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLTempDbSettings)(nil)).Elem()
}

func (o SQLTempDbSettingsPtrOutput) ToSQLTempDbSettingsPtrOutput() SQLTempDbSettingsPtrOutput {
	return o
}

func (o SQLTempDbSettingsPtrOutput) ToSQLTempDbSettingsPtrOutputWithContext(ctx context.Context) SQLTempDbSettingsPtrOutput {
	return o
}

func (o SQLTempDbSettingsPtrOutput) Elem() SQLTempDbSettingsOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) SQLTempDbSettings {
		if v != nil {
			return *v
		}
		var ret SQLTempDbSettings
		return ret
	}).(SQLTempDbSettingsOutput)
}

func (o SQLTempDbSettingsPtrOutput) DataFileCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.DataFileCount
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) DataFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.DataFileSize
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) DataGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.DataGrowth
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *string {
		if v == nil {
			return nil
		}
		return v.DefaultFilePath
	}).(pulumi.StringPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) LogFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.LogFileSize
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) LogGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *int {
		if v == nil {
			return nil
		}
		return v.LogGrowth
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) []int {
		if v == nil {
			return nil
		}
		return v.Luns
	}).(pulumi.IntArrayOutput)
}

func (o SQLTempDbSettingsPtrOutput) PersistFolder() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *bool {
		if v == nil {
			return nil
		}
		return v.PersistFolder
	}).(pulumi.BoolPtrOutput)
}

func (o SQLTempDbSettingsPtrOutput) PersistFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettings) *string {
		if v == nil {
			return nil
		}
		return v.PersistFolderPath
	}).(pulumi.StringPtrOutput)
}

type SQLTempDbSettingsResponse struct {
	DataFileCount     *int    `pulumi:"dataFileCount"`
	DataFileSize      *int    `pulumi:"dataFileSize"`
	DataGrowth        *int    `pulumi:"dataGrowth"`
	DefaultFilePath   *string `pulumi:"defaultFilePath"`
	LogFileSize       *int    `pulumi:"logFileSize"`
	LogGrowth         *int    `pulumi:"logGrowth"`
	Luns              []int   `pulumi:"luns"`
	PersistFolder     *bool   `pulumi:"persistFolder"`
	PersistFolderPath *string `pulumi:"persistFolderPath"`
}

type SQLTempDbSettingsResponseOutput struct{ *pulumi.OutputState }

func (SQLTempDbSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SQLTempDbSettingsResponse)(nil)).Elem()
}

func (o SQLTempDbSettingsResponseOutput) ToSQLTempDbSettingsResponseOutput() SQLTempDbSettingsResponseOutput {
	return o
}

func (o SQLTempDbSettingsResponseOutput) ToSQLTempDbSettingsResponseOutputWithContext(ctx context.Context) SQLTempDbSettingsResponseOutput {
	return o
}

func (o SQLTempDbSettingsResponseOutput) DataFileCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *int { return v.DataFileCount }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) DataFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *int { return v.DataFileSize }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) DataGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *int { return v.DataGrowth }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *string { return v.DefaultFilePath }).(pulumi.StringPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) LogFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *int { return v.LogFileSize }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) LogGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *int { return v.LogGrowth }).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) []int { return v.Luns }).(pulumi.IntArrayOutput)
}

func (o SQLTempDbSettingsResponseOutput) PersistFolder() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *bool { return v.PersistFolder }).(pulumi.BoolPtrOutput)
}

func (o SQLTempDbSettingsResponseOutput) PersistFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SQLTempDbSettingsResponse) *string { return v.PersistFolderPath }).(pulumi.StringPtrOutput)
}

type SQLTempDbSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SQLTempDbSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SQLTempDbSettingsResponse)(nil)).Elem()
}

func (o SQLTempDbSettingsResponsePtrOutput) ToSQLTempDbSettingsResponsePtrOutput() SQLTempDbSettingsResponsePtrOutput {
	return o
}

func (o SQLTempDbSettingsResponsePtrOutput) ToSQLTempDbSettingsResponsePtrOutputWithContext(ctx context.Context) SQLTempDbSettingsResponsePtrOutput {
	return o
}

func (o SQLTempDbSettingsResponsePtrOutput) Elem() SQLTempDbSettingsResponseOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) SQLTempDbSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SQLTempDbSettingsResponse
		return ret
	}).(SQLTempDbSettingsResponseOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) DataFileCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DataFileCount
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) DataFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DataFileSize
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) DataGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DataGrowth
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) DefaultFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DefaultFilePath
	}).(pulumi.StringPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) LogFileSize() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogFileSize
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) LogGrowth() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogGrowth
	}).(pulumi.IntPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) Luns() pulumi.IntArrayOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) []int {
		if v == nil {
			return nil
		}
		return v.Luns
	}).(pulumi.IntArrayOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) PersistFolder() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.PersistFolder
	}).(pulumi.BoolPtrOutput)
}

func (o SQLTempDbSettingsResponsePtrOutput) PersistFolderPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SQLTempDbSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.PersistFolderPath
	}).(pulumi.StringPtrOutput)
}

type Schedule struct {
	DayOfWeek         *AssessmentDayOfWeek `pulumi:"dayOfWeek"`
	Enable            *bool                `pulumi:"enable"`
	MonthlyOccurrence *int                 `pulumi:"monthlyOccurrence"`
	StartTime         *string              `pulumi:"startTime"`
	WeeklyInterval    *int                 `pulumi:"weeklyInterval"`
}





type ScheduleInput interface {
	pulumi.Input

	ToScheduleOutput() ScheduleOutput
	ToScheduleOutputWithContext(context.Context) ScheduleOutput
}

type ScheduleArgs struct {
	DayOfWeek         AssessmentDayOfWeekPtrInput `pulumi:"dayOfWeek"`
	Enable            pulumi.BoolPtrInput         `pulumi:"enable"`
	MonthlyOccurrence pulumi.IntPtrInput          `pulumi:"monthlyOccurrence"`
	StartTime         pulumi.StringPtrInput       `pulumi:"startTime"`
	WeeklyInterval    pulumi.IntPtrInput          `pulumi:"weeklyInterval"`
}

func (ScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (i ScheduleArgs) ToScheduleOutput() ScheduleOutput {
	return i.ToScheduleOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput)
}

func (i ScheduleArgs) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i ScheduleArgs) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScheduleOutput).ToSchedulePtrOutputWithContext(ctx)
}









type SchedulePtrInput interface {
	pulumi.Input

	ToSchedulePtrOutput() SchedulePtrOutput
	ToSchedulePtrOutputWithContext(context.Context) SchedulePtrOutput
}

type schedulePtrType ScheduleArgs

func SchedulePtr(v *ScheduleArgs) SchedulePtrInput {
	return (*schedulePtrType)(v)
}

func (*schedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (i *schedulePtrType) ToSchedulePtrOutput() SchedulePtrOutput {
	return i.ToSchedulePtrOutputWithContext(context.Background())
}

func (i *schedulePtrType) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SchedulePtrOutput)
}

type ScheduleOutput struct{ *pulumi.OutputState }

func (ScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Schedule)(nil)).Elem()
}

func (o ScheduleOutput) ToScheduleOutput() ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToScheduleOutputWithContext(ctx context.Context) ScheduleOutput {
	return o
}

func (o ScheduleOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o.ToSchedulePtrOutputWithContext(context.Background())
}

func (o ScheduleOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Schedule) *Schedule {
		return &v
	}).(SchedulePtrOutput)
}

func (o ScheduleOutput) DayOfWeek() AssessmentDayOfWeekPtrOutput {
	return o.ApplyT(func(v Schedule) *AssessmentDayOfWeek { return v.DayOfWeek }).(AssessmentDayOfWeekPtrOutput)
}

func (o ScheduleOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Schedule) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o ScheduleOutput) MonthlyOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Schedule) *int { return v.MonthlyOccurrence }).(pulumi.IntPtrOutput)
}

func (o ScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Schedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ScheduleOutput) WeeklyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Schedule) *int { return v.WeeklyInterval }).(pulumi.IntPtrOutput)
}

type SchedulePtrOutput struct{ *pulumi.OutputState }

func (SchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Schedule)(nil)).Elem()
}

func (o SchedulePtrOutput) ToSchedulePtrOutput() SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) ToSchedulePtrOutputWithContext(ctx context.Context) SchedulePtrOutput {
	return o
}

func (o SchedulePtrOutput) Elem() ScheduleOutput {
	return o.ApplyT(func(v *Schedule) Schedule {
		if v != nil {
			return *v
		}
		var ret Schedule
		return ret
	}).(ScheduleOutput)
}

func (o SchedulePtrOutput) DayOfWeek() AssessmentDayOfWeekPtrOutput {
	return o.ApplyT(func(v *Schedule) *AssessmentDayOfWeek {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(AssessmentDayOfWeekPtrOutput)
}

func (o SchedulePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Schedule) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o SchedulePtrOutput) MonthlyOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrence
	}).(pulumi.IntPtrOutput)
}

func (o SchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Schedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o SchedulePtrOutput) WeeklyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Schedule) *int {
		if v == nil {
			return nil
		}
		return v.WeeklyInterval
	}).(pulumi.IntPtrOutput)
}

type ScheduleResponse struct {
	DayOfWeek         *string `pulumi:"dayOfWeek"`
	Enable            *bool   `pulumi:"enable"`
	MonthlyOccurrence *int    `pulumi:"monthlyOccurrence"`
	StartTime         *string `pulumi:"startTime"`
	WeeklyInterval    *int    `pulumi:"weeklyInterval"`
}

type ScheduleResponseOutput struct{ *pulumi.OutputState }

func (ScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponseOutput) ToScheduleResponseOutput() ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) ToScheduleResponseOutputWithContext(ctx context.Context) ScheduleResponseOutput {
	return o
}

func (o ScheduleResponseOutput) DayOfWeek() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.DayOfWeek }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *bool { return v.Enable }).(pulumi.BoolPtrOutput)
}

func (o ScheduleResponseOutput) MonthlyOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *int { return v.MonthlyOccurrence }).(pulumi.IntPtrOutput)
}

func (o ScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ScheduleResponseOutput) WeeklyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScheduleResponse) *int { return v.WeeklyInterval }).(pulumi.IntPtrOutput)
}

type ScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScheduleResponse)(nil)).Elem()
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutput() ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) ToScheduleResponsePtrOutputWithContext(ctx context.Context) ScheduleResponsePtrOutput {
	return o
}

func (o ScheduleResponsePtrOutput) Elem() ScheduleResponseOutput {
	return o.ApplyT(func(v *ScheduleResponse) ScheduleResponse {
		if v != nil {
			return *v
		}
		var ret ScheduleResponse
		return ret
	}).(ScheduleResponseOutput)
}

func (o ScheduleResponsePtrOutput) DayOfWeek() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.DayOfWeek
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) Enable() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enable
	}).(pulumi.BoolPtrOutput)
}

func (o ScheduleResponsePtrOutput) MonthlyOccurrence() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.MonthlyOccurrence
	}).(pulumi.IntPtrOutput)
}

func (o ScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o ScheduleResponsePtrOutput) WeeklyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return v.WeeklyInterval
	}).(pulumi.IntPtrOutput)
}

type ServerConfigurationsManagementSettings struct {
	AdditionalFeaturesServerConfigurations *AdditionalFeaturesServerConfigurations `pulumi:"additionalFeaturesServerConfigurations"`
	AzureAdAuthenticationSettings          *AADAuthenticationSettings              `pulumi:"azureAdAuthenticationSettings"`
	SqlConnectivityUpdateSettings          *SqlConnectivityUpdateSettings          `pulumi:"sqlConnectivityUpdateSettings"`
	SqlInstanceSettings                    *SQLInstanceSettings                    `pulumi:"sqlInstanceSettings"`
	SqlStorageUpdateSettings               *SqlStorageUpdateSettings               `pulumi:"sqlStorageUpdateSettings"`
	SqlWorkloadTypeUpdateSettings          *SqlWorkloadTypeUpdateSettings          `pulumi:"sqlWorkloadTypeUpdateSettings"`
}





type ServerConfigurationsManagementSettingsInput interface {
	pulumi.Input

	ToServerConfigurationsManagementSettingsOutput() ServerConfigurationsManagementSettingsOutput
	ToServerConfigurationsManagementSettingsOutputWithContext(context.Context) ServerConfigurationsManagementSettingsOutput
}

type ServerConfigurationsManagementSettingsArgs struct {
	AdditionalFeaturesServerConfigurations AdditionalFeaturesServerConfigurationsPtrInput `pulumi:"additionalFeaturesServerConfigurations"`
	AzureAdAuthenticationSettings          AADAuthenticationSettingsPtrInput              `pulumi:"azureAdAuthenticationSettings"`
	SqlConnectivityUpdateSettings          SqlConnectivityUpdateSettingsPtrInput          `pulumi:"sqlConnectivityUpdateSettings"`
	SqlInstanceSettings                    SQLInstanceSettingsPtrInput                    `pulumi:"sqlInstanceSettings"`
	SqlStorageUpdateSettings               SqlStorageUpdateSettingsPtrInput               `pulumi:"sqlStorageUpdateSettings"`
	SqlWorkloadTypeUpdateSettings          SqlWorkloadTypeUpdateSettingsPtrInput          `pulumi:"sqlWorkloadTypeUpdateSettings"`
}

func (ServerConfigurationsManagementSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerConfigurationsManagementSettings)(nil)).Elem()
}

func (i ServerConfigurationsManagementSettingsArgs) ToServerConfigurationsManagementSettingsOutput() ServerConfigurationsManagementSettingsOutput {
	return i.ToServerConfigurationsManagementSettingsOutputWithContext(context.Background())
}

func (i ServerConfigurationsManagementSettingsArgs) ToServerConfigurationsManagementSettingsOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerConfigurationsManagementSettingsOutput)
}

func (i ServerConfigurationsManagementSettingsArgs) ToServerConfigurationsManagementSettingsPtrOutput() ServerConfigurationsManagementSettingsPtrOutput {
	return i.ToServerConfigurationsManagementSettingsPtrOutputWithContext(context.Background())
}

func (i ServerConfigurationsManagementSettingsArgs) ToServerConfigurationsManagementSettingsPtrOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerConfigurationsManagementSettingsOutput).ToServerConfigurationsManagementSettingsPtrOutputWithContext(ctx)
}









type ServerConfigurationsManagementSettingsPtrInput interface {
	pulumi.Input

	ToServerConfigurationsManagementSettingsPtrOutput() ServerConfigurationsManagementSettingsPtrOutput
	ToServerConfigurationsManagementSettingsPtrOutputWithContext(context.Context) ServerConfigurationsManagementSettingsPtrOutput
}

type serverConfigurationsManagementSettingsPtrType ServerConfigurationsManagementSettingsArgs

func ServerConfigurationsManagementSettingsPtr(v *ServerConfigurationsManagementSettingsArgs) ServerConfigurationsManagementSettingsPtrInput {
	return (*serverConfigurationsManagementSettingsPtrType)(v)
}

func (*serverConfigurationsManagementSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerConfigurationsManagementSettings)(nil)).Elem()
}

func (i *serverConfigurationsManagementSettingsPtrType) ToServerConfigurationsManagementSettingsPtrOutput() ServerConfigurationsManagementSettingsPtrOutput {
	return i.ToServerConfigurationsManagementSettingsPtrOutputWithContext(context.Background())
}

func (i *serverConfigurationsManagementSettingsPtrType) ToServerConfigurationsManagementSettingsPtrOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServerConfigurationsManagementSettingsPtrOutput)
}

type ServerConfigurationsManagementSettingsOutput struct{ *pulumi.OutputState }

func (ServerConfigurationsManagementSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerConfigurationsManagementSettings)(nil)).Elem()
}

func (o ServerConfigurationsManagementSettingsOutput) ToServerConfigurationsManagementSettingsOutput() ServerConfigurationsManagementSettingsOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsOutput) ToServerConfigurationsManagementSettingsOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsOutput) ToServerConfigurationsManagementSettingsPtrOutput() ServerConfigurationsManagementSettingsPtrOutput {
	return o.ToServerConfigurationsManagementSettingsPtrOutputWithContext(context.Background())
}

func (o ServerConfigurationsManagementSettingsOutput) ToServerConfigurationsManagementSettingsPtrOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServerConfigurationsManagementSettings) *ServerConfigurationsManagementSettings {
		return &v
	}).(ServerConfigurationsManagementSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsOutput) AdditionalFeaturesServerConfigurations() AdditionalFeaturesServerConfigurationsPtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettings) *AdditionalFeaturesServerConfigurations {
		return v.AdditionalFeaturesServerConfigurations
	}).(AdditionalFeaturesServerConfigurationsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsOutput) AzureAdAuthenticationSettings() AADAuthenticationSettingsPtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettings) *AADAuthenticationSettings {
		return v.AzureAdAuthenticationSettings
	}).(AADAuthenticationSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsOutput) SqlConnectivityUpdateSettings() SqlConnectivityUpdateSettingsPtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettings) *SqlConnectivityUpdateSettings {
		return v.SqlConnectivityUpdateSettings
	}).(SqlConnectivityUpdateSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsOutput) SqlInstanceSettings() SQLInstanceSettingsPtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettings) *SQLInstanceSettings { return v.SqlInstanceSettings }).(SQLInstanceSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsOutput) SqlStorageUpdateSettings() SqlStorageUpdateSettingsPtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettings) *SqlStorageUpdateSettings {
		return v.SqlStorageUpdateSettings
	}).(SqlStorageUpdateSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsOutput) SqlWorkloadTypeUpdateSettings() SqlWorkloadTypeUpdateSettingsPtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettings) *SqlWorkloadTypeUpdateSettings {
		return v.SqlWorkloadTypeUpdateSettings
	}).(SqlWorkloadTypeUpdateSettingsPtrOutput)
}

type ServerConfigurationsManagementSettingsPtrOutput struct{ *pulumi.OutputState }

func (ServerConfigurationsManagementSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerConfigurationsManagementSettings)(nil)).Elem()
}

func (o ServerConfigurationsManagementSettingsPtrOutput) ToServerConfigurationsManagementSettingsPtrOutput() ServerConfigurationsManagementSettingsPtrOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsPtrOutput) ToServerConfigurationsManagementSettingsPtrOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsPtrOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsPtrOutput) Elem() ServerConfigurationsManagementSettingsOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) ServerConfigurationsManagementSettings {
		if v != nil {
			return *v
		}
		var ret ServerConfigurationsManagementSettings
		return ret
	}).(ServerConfigurationsManagementSettingsOutput)
}

func (o ServerConfigurationsManagementSettingsPtrOutput) AdditionalFeaturesServerConfigurations() AdditionalFeaturesServerConfigurationsPtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) *AdditionalFeaturesServerConfigurations {
		if v == nil {
			return nil
		}
		return v.AdditionalFeaturesServerConfigurations
	}).(AdditionalFeaturesServerConfigurationsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsPtrOutput) AzureAdAuthenticationSettings() AADAuthenticationSettingsPtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) *AADAuthenticationSettings {
		if v == nil {
			return nil
		}
		return v.AzureAdAuthenticationSettings
	}).(AADAuthenticationSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsPtrOutput) SqlConnectivityUpdateSettings() SqlConnectivityUpdateSettingsPtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) *SqlConnectivityUpdateSettings {
		if v == nil {
			return nil
		}
		return v.SqlConnectivityUpdateSettings
	}).(SqlConnectivityUpdateSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsPtrOutput) SqlInstanceSettings() SQLInstanceSettingsPtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) *SQLInstanceSettings {
		if v == nil {
			return nil
		}
		return v.SqlInstanceSettings
	}).(SQLInstanceSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsPtrOutput) SqlStorageUpdateSettings() SqlStorageUpdateSettingsPtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) *SqlStorageUpdateSettings {
		if v == nil {
			return nil
		}
		return v.SqlStorageUpdateSettings
	}).(SqlStorageUpdateSettingsPtrOutput)
}

func (o ServerConfigurationsManagementSettingsPtrOutput) SqlWorkloadTypeUpdateSettings() SqlWorkloadTypeUpdateSettingsPtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettings) *SqlWorkloadTypeUpdateSettings {
		if v == nil {
			return nil
		}
		return v.SqlWorkloadTypeUpdateSettings
	}).(SqlWorkloadTypeUpdateSettingsPtrOutput)
}

type ServerConfigurationsManagementSettingsResponse struct {
	AdditionalFeaturesServerConfigurations *AdditionalFeaturesServerConfigurationsResponse `pulumi:"additionalFeaturesServerConfigurations"`
	AzureAdAuthenticationSettings          *AADAuthenticationSettingsResponse              `pulumi:"azureAdAuthenticationSettings"`
	SqlConnectivityUpdateSettings          *SqlConnectivityUpdateSettingsResponse          `pulumi:"sqlConnectivityUpdateSettings"`
	SqlInstanceSettings                    *SQLInstanceSettingsResponse                    `pulumi:"sqlInstanceSettings"`
	SqlStorageUpdateSettings               *SqlStorageUpdateSettingsResponse               `pulumi:"sqlStorageUpdateSettings"`
	SqlWorkloadTypeUpdateSettings          *SqlWorkloadTypeUpdateSettingsResponse          `pulumi:"sqlWorkloadTypeUpdateSettings"`
}

type ServerConfigurationsManagementSettingsResponseOutput struct{ *pulumi.OutputState }

func (ServerConfigurationsManagementSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServerConfigurationsManagementSettingsResponse)(nil)).Elem()
}

func (o ServerConfigurationsManagementSettingsResponseOutput) ToServerConfigurationsManagementSettingsResponseOutput() ServerConfigurationsManagementSettingsResponseOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsResponseOutput) ToServerConfigurationsManagementSettingsResponseOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsResponseOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsResponseOutput) AdditionalFeaturesServerConfigurations() AdditionalFeaturesServerConfigurationsResponsePtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettingsResponse) *AdditionalFeaturesServerConfigurationsResponse {
		return v.AdditionalFeaturesServerConfigurations
	}).(AdditionalFeaturesServerConfigurationsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponseOutput) AzureAdAuthenticationSettings() AADAuthenticationSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettingsResponse) *AADAuthenticationSettingsResponse {
		return v.AzureAdAuthenticationSettings
	}).(AADAuthenticationSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponseOutput) SqlConnectivityUpdateSettings() SqlConnectivityUpdateSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettingsResponse) *SqlConnectivityUpdateSettingsResponse {
		return v.SqlConnectivityUpdateSettings
	}).(SqlConnectivityUpdateSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponseOutput) SqlInstanceSettings() SQLInstanceSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettingsResponse) *SQLInstanceSettingsResponse {
		return v.SqlInstanceSettings
	}).(SQLInstanceSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponseOutput) SqlStorageUpdateSettings() SqlStorageUpdateSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettingsResponse) *SqlStorageUpdateSettingsResponse {
		return v.SqlStorageUpdateSettings
	}).(SqlStorageUpdateSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponseOutput) SqlWorkloadTypeUpdateSettings() SqlWorkloadTypeUpdateSettingsResponsePtrOutput {
	return o.ApplyT(func(v ServerConfigurationsManagementSettingsResponse) *SqlWorkloadTypeUpdateSettingsResponse {
		return v.SqlWorkloadTypeUpdateSettings
	}).(SqlWorkloadTypeUpdateSettingsResponsePtrOutput)
}

type ServerConfigurationsManagementSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ServerConfigurationsManagementSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServerConfigurationsManagementSettingsResponse)(nil)).Elem()
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) ToServerConfigurationsManagementSettingsResponsePtrOutput() ServerConfigurationsManagementSettingsResponsePtrOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) ToServerConfigurationsManagementSettingsResponsePtrOutputWithContext(ctx context.Context) ServerConfigurationsManagementSettingsResponsePtrOutput {
	return o
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) Elem() ServerConfigurationsManagementSettingsResponseOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) ServerConfigurationsManagementSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ServerConfigurationsManagementSettingsResponse
		return ret
	}).(ServerConfigurationsManagementSettingsResponseOutput)
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) AdditionalFeaturesServerConfigurations() AdditionalFeaturesServerConfigurationsResponsePtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) *AdditionalFeaturesServerConfigurationsResponse {
		if v == nil {
			return nil
		}
		return v.AdditionalFeaturesServerConfigurations
	}).(AdditionalFeaturesServerConfigurationsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) AzureAdAuthenticationSettings() AADAuthenticationSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) *AADAuthenticationSettingsResponse {
		if v == nil {
			return nil
		}
		return v.AzureAdAuthenticationSettings
	}).(AADAuthenticationSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) SqlConnectivityUpdateSettings() SqlConnectivityUpdateSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) *SqlConnectivityUpdateSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlConnectivityUpdateSettings
	}).(SqlConnectivityUpdateSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) SqlInstanceSettings() SQLInstanceSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) *SQLInstanceSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlInstanceSettings
	}).(SQLInstanceSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) SqlStorageUpdateSettings() SqlStorageUpdateSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) *SqlStorageUpdateSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlStorageUpdateSettings
	}).(SqlStorageUpdateSettingsResponsePtrOutput)
}

func (o ServerConfigurationsManagementSettingsResponsePtrOutput) SqlWorkloadTypeUpdateSettings() SqlWorkloadTypeUpdateSettingsResponsePtrOutput {
	return o.ApplyT(func(v *ServerConfigurationsManagementSettingsResponse) *SqlWorkloadTypeUpdateSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlWorkloadTypeUpdateSettings
	}).(SqlWorkloadTypeUpdateSettingsResponsePtrOutput)
}

type SqlConnectivityUpdateSettings struct {
	ConnectivityType      *string `pulumi:"connectivityType"`
	Port                  *int    `pulumi:"port"`
	SqlAuthUpdatePassword *string `pulumi:"sqlAuthUpdatePassword"`
	SqlAuthUpdateUserName *string `pulumi:"sqlAuthUpdateUserName"`
}





type SqlConnectivityUpdateSettingsInput interface {
	pulumi.Input

	ToSqlConnectivityUpdateSettingsOutput() SqlConnectivityUpdateSettingsOutput
	ToSqlConnectivityUpdateSettingsOutputWithContext(context.Context) SqlConnectivityUpdateSettingsOutput
}

type SqlConnectivityUpdateSettingsArgs struct {
	ConnectivityType      pulumi.StringPtrInput `pulumi:"connectivityType"`
	Port                  pulumi.IntPtrInput    `pulumi:"port"`
	SqlAuthUpdatePassword pulumi.StringPtrInput `pulumi:"sqlAuthUpdatePassword"`
	SqlAuthUpdateUserName pulumi.StringPtrInput `pulumi:"sqlAuthUpdateUserName"`
}

func (SqlConnectivityUpdateSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectivityUpdateSettings)(nil)).Elem()
}

func (i SqlConnectivityUpdateSettingsArgs) ToSqlConnectivityUpdateSettingsOutput() SqlConnectivityUpdateSettingsOutput {
	return i.ToSqlConnectivityUpdateSettingsOutputWithContext(context.Background())
}

func (i SqlConnectivityUpdateSettingsArgs) ToSqlConnectivityUpdateSettingsOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectivityUpdateSettingsOutput)
}

func (i SqlConnectivityUpdateSettingsArgs) ToSqlConnectivityUpdateSettingsPtrOutput() SqlConnectivityUpdateSettingsPtrOutput {
	return i.ToSqlConnectivityUpdateSettingsPtrOutputWithContext(context.Background())
}

func (i SqlConnectivityUpdateSettingsArgs) ToSqlConnectivityUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectivityUpdateSettingsOutput).ToSqlConnectivityUpdateSettingsPtrOutputWithContext(ctx)
}









type SqlConnectivityUpdateSettingsPtrInput interface {
	pulumi.Input

	ToSqlConnectivityUpdateSettingsPtrOutput() SqlConnectivityUpdateSettingsPtrOutput
	ToSqlConnectivityUpdateSettingsPtrOutputWithContext(context.Context) SqlConnectivityUpdateSettingsPtrOutput
}

type sqlConnectivityUpdateSettingsPtrType SqlConnectivityUpdateSettingsArgs

func SqlConnectivityUpdateSettingsPtr(v *SqlConnectivityUpdateSettingsArgs) SqlConnectivityUpdateSettingsPtrInput {
	return (*sqlConnectivityUpdateSettingsPtrType)(v)
}

func (*sqlConnectivityUpdateSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectivityUpdateSettings)(nil)).Elem()
}

func (i *sqlConnectivityUpdateSettingsPtrType) ToSqlConnectivityUpdateSettingsPtrOutput() SqlConnectivityUpdateSettingsPtrOutput {
	return i.ToSqlConnectivityUpdateSettingsPtrOutputWithContext(context.Background())
}

func (i *sqlConnectivityUpdateSettingsPtrType) ToSqlConnectivityUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlConnectivityUpdateSettingsPtrOutput)
}

type SqlConnectivityUpdateSettingsOutput struct{ *pulumi.OutputState }

func (SqlConnectivityUpdateSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectivityUpdateSettings)(nil)).Elem()
}

func (o SqlConnectivityUpdateSettingsOutput) ToSqlConnectivityUpdateSettingsOutput() SqlConnectivityUpdateSettingsOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsOutput) ToSqlConnectivityUpdateSettingsOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsOutput) ToSqlConnectivityUpdateSettingsPtrOutput() SqlConnectivityUpdateSettingsPtrOutput {
	return o.ToSqlConnectivityUpdateSettingsPtrOutputWithContext(context.Background())
}

func (o SqlConnectivityUpdateSettingsOutput) ToSqlConnectivityUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlConnectivityUpdateSettings) *SqlConnectivityUpdateSettings {
		return &v
	}).(SqlConnectivityUpdateSettingsPtrOutput)
}

func (o SqlConnectivityUpdateSettingsOutput) ConnectivityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectivityUpdateSettings) *string { return v.ConnectivityType }).(pulumi.StringPtrOutput)
}

func (o SqlConnectivityUpdateSettingsOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlConnectivityUpdateSettings) *int { return v.Port }).(pulumi.IntPtrOutput)
}

func (o SqlConnectivityUpdateSettingsOutput) SqlAuthUpdatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectivityUpdateSettings) *string { return v.SqlAuthUpdatePassword }).(pulumi.StringPtrOutput)
}

func (o SqlConnectivityUpdateSettingsOutput) SqlAuthUpdateUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectivityUpdateSettings) *string { return v.SqlAuthUpdateUserName }).(pulumi.StringPtrOutput)
}

type SqlConnectivityUpdateSettingsPtrOutput struct{ *pulumi.OutputState }

func (SqlConnectivityUpdateSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectivityUpdateSettings)(nil)).Elem()
}

func (o SqlConnectivityUpdateSettingsPtrOutput) ToSqlConnectivityUpdateSettingsPtrOutput() SqlConnectivityUpdateSettingsPtrOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsPtrOutput) ToSqlConnectivityUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsPtrOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsPtrOutput) Elem() SqlConnectivityUpdateSettingsOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettings) SqlConnectivityUpdateSettings {
		if v != nil {
			return *v
		}
		var ret SqlConnectivityUpdateSettings
		return ret
	}).(SqlConnectivityUpdateSettingsOutput)
}

func (o SqlConnectivityUpdateSettingsPtrOutput) ConnectivityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettings) *string {
		if v == nil {
			return nil
		}
		return v.ConnectivityType
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectivityUpdateSettingsPtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettings) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

func (o SqlConnectivityUpdateSettingsPtrOutput) SqlAuthUpdatePassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettings) *string {
		if v == nil {
			return nil
		}
		return v.SqlAuthUpdatePassword
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectivityUpdateSettingsPtrOutput) SqlAuthUpdateUserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettings) *string {
		if v == nil {
			return nil
		}
		return v.SqlAuthUpdateUserName
	}).(pulumi.StringPtrOutput)
}

type SqlConnectivityUpdateSettingsResponse struct {
	ConnectivityType *string `pulumi:"connectivityType"`
	Port             *int    `pulumi:"port"`
}

type SqlConnectivityUpdateSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlConnectivityUpdateSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlConnectivityUpdateSettingsResponse)(nil)).Elem()
}

func (o SqlConnectivityUpdateSettingsResponseOutput) ToSqlConnectivityUpdateSettingsResponseOutput() SqlConnectivityUpdateSettingsResponseOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsResponseOutput) ToSqlConnectivityUpdateSettingsResponseOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsResponseOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsResponseOutput) ConnectivityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlConnectivityUpdateSettingsResponse) *string { return v.ConnectivityType }).(pulumi.StringPtrOutput)
}

func (o SqlConnectivityUpdateSettingsResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlConnectivityUpdateSettingsResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type SqlConnectivityUpdateSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlConnectivityUpdateSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlConnectivityUpdateSettingsResponse)(nil)).Elem()
}

func (o SqlConnectivityUpdateSettingsResponsePtrOutput) ToSqlConnectivityUpdateSettingsResponsePtrOutput() SqlConnectivityUpdateSettingsResponsePtrOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsResponsePtrOutput) ToSqlConnectivityUpdateSettingsResponsePtrOutputWithContext(ctx context.Context) SqlConnectivityUpdateSettingsResponsePtrOutput {
	return o
}

func (o SqlConnectivityUpdateSettingsResponsePtrOutput) Elem() SqlConnectivityUpdateSettingsResponseOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettingsResponse) SqlConnectivityUpdateSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SqlConnectivityUpdateSettingsResponse
		return ret
	}).(SqlConnectivityUpdateSettingsResponseOutput)
}

func (o SqlConnectivityUpdateSettingsResponsePtrOutput) ConnectivityType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ConnectivityType
	}).(pulumi.StringPtrOutput)
}

func (o SqlConnectivityUpdateSettingsResponsePtrOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlConnectivityUpdateSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.Port
	}).(pulumi.IntPtrOutput)
}

type SqlStorageUpdateSettings struct {
	DiskConfigurationType *string `pulumi:"diskConfigurationType"`
	DiskCount             *int    `pulumi:"diskCount"`
	StartingDeviceId      *int    `pulumi:"startingDeviceId"`
}





type SqlStorageUpdateSettingsInput interface {
	pulumi.Input

	ToSqlStorageUpdateSettingsOutput() SqlStorageUpdateSettingsOutput
	ToSqlStorageUpdateSettingsOutputWithContext(context.Context) SqlStorageUpdateSettingsOutput
}

type SqlStorageUpdateSettingsArgs struct {
	DiskConfigurationType pulumi.StringPtrInput `pulumi:"diskConfigurationType"`
	DiskCount             pulumi.IntPtrInput    `pulumi:"diskCount"`
	StartingDeviceId      pulumi.IntPtrInput    `pulumi:"startingDeviceId"`
}

func (SqlStorageUpdateSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStorageUpdateSettings)(nil)).Elem()
}

func (i SqlStorageUpdateSettingsArgs) ToSqlStorageUpdateSettingsOutput() SqlStorageUpdateSettingsOutput {
	return i.ToSqlStorageUpdateSettingsOutputWithContext(context.Background())
}

func (i SqlStorageUpdateSettingsArgs) ToSqlStorageUpdateSettingsOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStorageUpdateSettingsOutput)
}

func (i SqlStorageUpdateSettingsArgs) ToSqlStorageUpdateSettingsPtrOutput() SqlStorageUpdateSettingsPtrOutput {
	return i.ToSqlStorageUpdateSettingsPtrOutputWithContext(context.Background())
}

func (i SqlStorageUpdateSettingsArgs) ToSqlStorageUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStorageUpdateSettingsOutput).ToSqlStorageUpdateSettingsPtrOutputWithContext(ctx)
}









type SqlStorageUpdateSettingsPtrInput interface {
	pulumi.Input

	ToSqlStorageUpdateSettingsPtrOutput() SqlStorageUpdateSettingsPtrOutput
	ToSqlStorageUpdateSettingsPtrOutputWithContext(context.Context) SqlStorageUpdateSettingsPtrOutput
}

type sqlStorageUpdateSettingsPtrType SqlStorageUpdateSettingsArgs

func SqlStorageUpdateSettingsPtr(v *SqlStorageUpdateSettingsArgs) SqlStorageUpdateSettingsPtrInput {
	return (*sqlStorageUpdateSettingsPtrType)(v)
}

func (*sqlStorageUpdateSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStorageUpdateSettings)(nil)).Elem()
}

func (i *sqlStorageUpdateSettingsPtrType) ToSqlStorageUpdateSettingsPtrOutput() SqlStorageUpdateSettingsPtrOutput {
	return i.ToSqlStorageUpdateSettingsPtrOutputWithContext(context.Background())
}

func (i *sqlStorageUpdateSettingsPtrType) ToSqlStorageUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlStorageUpdateSettingsPtrOutput)
}

type SqlStorageUpdateSettingsOutput struct{ *pulumi.OutputState }

func (SqlStorageUpdateSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStorageUpdateSettings)(nil)).Elem()
}

func (o SqlStorageUpdateSettingsOutput) ToSqlStorageUpdateSettingsOutput() SqlStorageUpdateSettingsOutput {
	return o
}

func (o SqlStorageUpdateSettingsOutput) ToSqlStorageUpdateSettingsOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsOutput {
	return o
}

func (o SqlStorageUpdateSettingsOutput) ToSqlStorageUpdateSettingsPtrOutput() SqlStorageUpdateSettingsPtrOutput {
	return o.ToSqlStorageUpdateSettingsPtrOutputWithContext(context.Background())
}

func (o SqlStorageUpdateSettingsOutput) ToSqlStorageUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlStorageUpdateSettings) *SqlStorageUpdateSettings {
		return &v
	}).(SqlStorageUpdateSettingsPtrOutput)
}

func (o SqlStorageUpdateSettingsOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlStorageUpdateSettings) *string { return v.DiskConfigurationType }).(pulumi.StringPtrOutput)
}

func (o SqlStorageUpdateSettingsOutput) DiskCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlStorageUpdateSettings) *int { return v.DiskCount }).(pulumi.IntPtrOutput)
}

func (o SqlStorageUpdateSettingsOutput) StartingDeviceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlStorageUpdateSettings) *int { return v.StartingDeviceId }).(pulumi.IntPtrOutput)
}

type SqlStorageUpdateSettingsPtrOutput struct{ *pulumi.OutputState }

func (SqlStorageUpdateSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStorageUpdateSettings)(nil)).Elem()
}

func (o SqlStorageUpdateSettingsPtrOutput) ToSqlStorageUpdateSettingsPtrOutput() SqlStorageUpdateSettingsPtrOutput {
	return o
}

func (o SqlStorageUpdateSettingsPtrOutput) ToSqlStorageUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsPtrOutput {
	return o
}

func (o SqlStorageUpdateSettingsPtrOutput) Elem() SqlStorageUpdateSettingsOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettings) SqlStorageUpdateSettings {
		if v != nil {
			return *v
		}
		var ret SqlStorageUpdateSettings
		return ret
	}).(SqlStorageUpdateSettingsOutput)
}

func (o SqlStorageUpdateSettingsPtrOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettings) *string {
		if v == nil {
			return nil
		}
		return v.DiskConfigurationType
	}).(pulumi.StringPtrOutput)
}

func (o SqlStorageUpdateSettingsPtrOutput) DiskCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettings) *int {
		if v == nil {
			return nil
		}
		return v.DiskCount
	}).(pulumi.IntPtrOutput)
}

func (o SqlStorageUpdateSettingsPtrOutput) StartingDeviceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettings) *int {
		if v == nil {
			return nil
		}
		return v.StartingDeviceId
	}).(pulumi.IntPtrOutput)
}

type SqlStorageUpdateSettingsResponse struct {
	DiskConfigurationType *string `pulumi:"diskConfigurationType"`
	DiskCount             *int    `pulumi:"diskCount"`
	StartingDeviceId      *int    `pulumi:"startingDeviceId"`
}

type SqlStorageUpdateSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlStorageUpdateSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlStorageUpdateSettingsResponse)(nil)).Elem()
}

func (o SqlStorageUpdateSettingsResponseOutput) ToSqlStorageUpdateSettingsResponseOutput() SqlStorageUpdateSettingsResponseOutput {
	return o
}

func (o SqlStorageUpdateSettingsResponseOutput) ToSqlStorageUpdateSettingsResponseOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsResponseOutput {
	return o
}

func (o SqlStorageUpdateSettingsResponseOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlStorageUpdateSettingsResponse) *string { return v.DiskConfigurationType }).(pulumi.StringPtrOutput)
}

func (o SqlStorageUpdateSettingsResponseOutput) DiskCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlStorageUpdateSettingsResponse) *int { return v.DiskCount }).(pulumi.IntPtrOutput)
}

func (o SqlStorageUpdateSettingsResponseOutput) StartingDeviceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SqlStorageUpdateSettingsResponse) *int { return v.StartingDeviceId }).(pulumi.IntPtrOutput)
}

type SqlStorageUpdateSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlStorageUpdateSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlStorageUpdateSettingsResponse)(nil)).Elem()
}

func (o SqlStorageUpdateSettingsResponsePtrOutput) ToSqlStorageUpdateSettingsResponsePtrOutput() SqlStorageUpdateSettingsResponsePtrOutput {
	return o
}

func (o SqlStorageUpdateSettingsResponsePtrOutput) ToSqlStorageUpdateSettingsResponsePtrOutputWithContext(ctx context.Context) SqlStorageUpdateSettingsResponsePtrOutput {
	return o
}

func (o SqlStorageUpdateSettingsResponsePtrOutput) Elem() SqlStorageUpdateSettingsResponseOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettingsResponse) SqlStorageUpdateSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SqlStorageUpdateSettingsResponse
		return ret
	}).(SqlStorageUpdateSettingsResponseOutput)
}

func (o SqlStorageUpdateSettingsResponsePtrOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskConfigurationType
	}).(pulumi.StringPtrOutput)
}

func (o SqlStorageUpdateSettingsResponsePtrOutput) DiskCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.DiskCount
	}).(pulumi.IntPtrOutput)
}

func (o SqlStorageUpdateSettingsResponsePtrOutput) StartingDeviceId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SqlStorageUpdateSettingsResponse) *int {
		if v == nil {
			return nil
		}
		return v.StartingDeviceId
	}).(pulumi.IntPtrOutput)
}

type SqlWorkloadTypeUpdateSettings struct {
	SqlWorkloadType *string `pulumi:"sqlWorkloadType"`
}





type SqlWorkloadTypeUpdateSettingsInput interface {
	pulumi.Input

	ToSqlWorkloadTypeUpdateSettingsOutput() SqlWorkloadTypeUpdateSettingsOutput
	ToSqlWorkloadTypeUpdateSettingsOutputWithContext(context.Context) SqlWorkloadTypeUpdateSettingsOutput
}

type SqlWorkloadTypeUpdateSettingsArgs struct {
	SqlWorkloadType pulumi.StringPtrInput `pulumi:"sqlWorkloadType"`
}

func (SqlWorkloadTypeUpdateSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlWorkloadTypeUpdateSettings)(nil)).Elem()
}

func (i SqlWorkloadTypeUpdateSettingsArgs) ToSqlWorkloadTypeUpdateSettingsOutput() SqlWorkloadTypeUpdateSettingsOutput {
	return i.ToSqlWorkloadTypeUpdateSettingsOutputWithContext(context.Background())
}

func (i SqlWorkloadTypeUpdateSettingsArgs) ToSqlWorkloadTypeUpdateSettingsOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlWorkloadTypeUpdateSettingsOutput)
}

func (i SqlWorkloadTypeUpdateSettingsArgs) ToSqlWorkloadTypeUpdateSettingsPtrOutput() SqlWorkloadTypeUpdateSettingsPtrOutput {
	return i.ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(context.Background())
}

func (i SqlWorkloadTypeUpdateSettingsArgs) ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlWorkloadTypeUpdateSettingsOutput).ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(ctx)
}









type SqlWorkloadTypeUpdateSettingsPtrInput interface {
	pulumi.Input

	ToSqlWorkloadTypeUpdateSettingsPtrOutput() SqlWorkloadTypeUpdateSettingsPtrOutput
	ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(context.Context) SqlWorkloadTypeUpdateSettingsPtrOutput
}

type sqlWorkloadTypeUpdateSettingsPtrType SqlWorkloadTypeUpdateSettingsArgs

func SqlWorkloadTypeUpdateSettingsPtr(v *SqlWorkloadTypeUpdateSettingsArgs) SqlWorkloadTypeUpdateSettingsPtrInput {
	return (*sqlWorkloadTypeUpdateSettingsPtrType)(v)
}

func (*sqlWorkloadTypeUpdateSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlWorkloadTypeUpdateSettings)(nil)).Elem()
}

func (i *sqlWorkloadTypeUpdateSettingsPtrType) ToSqlWorkloadTypeUpdateSettingsPtrOutput() SqlWorkloadTypeUpdateSettingsPtrOutput {
	return i.ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(context.Background())
}

func (i *sqlWorkloadTypeUpdateSettingsPtrType) ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SqlWorkloadTypeUpdateSettingsPtrOutput)
}

type SqlWorkloadTypeUpdateSettingsOutput struct{ *pulumi.OutputState }

func (SqlWorkloadTypeUpdateSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlWorkloadTypeUpdateSettings)(nil)).Elem()
}

func (o SqlWorkloadTypeUpdateSettingsOutput) ToSqlWorkloadTypeUpdateSettingsOutput() SqlWorkloadTypeUpdateSettingsOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsOutput) ToSqlWorkloadTypeUpdateSettingsOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsOutput) ToSqlWorkloadTypeUpdateSettingsPtrOutput() SqlWorkloadTypeUpdateSettingsPtrOutput {
	return o.ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(context.Background())
}

func (o SqlWorkloadTypeUpdateSettingsOutput) ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SqlWorkloadTypeUpdateSettings) *SqlWorkloadTypeUpdateSettings {
		return &v
	}).(SqlWorkloadTypeUpdateSettingsPtrOutput)
}

func (o SqlWorkloadTypeUpdateSettingsOutput) SqlWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlWorkloadTypeUpdateSettings) *string { return v.SqlWorkloadType }).(pulumi.StringPtrOutput)
}

type SqlWorkloadTypeUpdateSettingsPtrOutput struct{ *pulumi.OutputState }

func (SqlWorkloadTypeUpdateSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlWorkloadTypeUpdateSettings)(nil)).Elem()
}

func (o SqlWorkloadTypeUpdateSettingsPtrOutput) ToSqlWorkloadTypeUpdateSettingsPtrOutput() SqlWorkloadTypeUpdateSettingsPtrOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsPtrOutput) ToSqlWorkloadTypeUpdateSettingsPtrOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsPtrOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsPtrOutput) Elem() SqlWorkloadTypeUpdateSettingsOutput {
	return o.ApplyT(func(v *SqlWorkloadTypeUpdateSettings) SqlWorkloadTypeUpdateSettings {
		if v != nil {
			return *v
		}
		var ret SqlWorkloadTypeUpdateSettings
		return ret
	}).(SqlWorkloadTypeUpdateSettingsOutput)
}

func (o SqlWorkloadTypeUpdateSettingsPtrOutput) SqlWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlWorkloadTypeUpdateSettings) *string {
		if v == nil {
			return nil
		}
		return v.SqlWorkloadType
	}).(pulumi.StringPtrOutput)
}

type SqlWorkloadTypeUpdateSettingsResponse struct {
	SqlWorkloadType *string `pulumi:"sqlWorkloadType"`
}

type SqlWorkloadTypeUpdateSettingsResponseOutput struct{ *pulumi.OutputState }

func (SqlWorkloadTypeUpdateSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SqlWorkloadTypeUpdateSettingsResponse)(nil)).Elem()
}

func (o SqlWorkloadTypeUpdateSettingsResponseOutput) ToSqlWorkloadTypeUpdateSettingsResponseOutput() SqlWorkloadTypeUpdateSettingsResponseOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsResponseOutput) ToSqlWorkloadTypeUpdateSettingsResponseOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsResponseOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsResponseOutput) SqlWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SqlWorkloadTypeUpdateSettingsResponse) *string { return v.SqlWorkloadType }).(pulumi.StringPtrOutput)
}

type SqlWorkloadTypeUpdateSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (SqlWorkloadTypeUpdateSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SqlWorkloadTypeUpdateSettingsResponse)(nil)).Elem()
}

func (o SqlWorkloadTypeUpdateSettingsResponsePtrOutput) ToSqlWorkloadTypeUpdateSettingsResponsePtrOutput() SqlWorkloadTypeUpdateSettingsResponsePtrOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsResponsePtrOutput) ToSqlWorkloadTypeUpdateSettingsResponsePtrOutputWithContext(ctx context.Context) SqlWorkloadTypeUpdateSettingsResponsePtrOutput {
	return o
}

func (o SqlWorkloadTypeUpdateSettingsResponsePtrOutput) Elem() SqlWorkloadTypeUpdateSettingsResponseOutput {
	return o.ApplyT(func(v *SqlWorkloadTypeUpdateSettingsResponse) SqlWorkloadTypeUpdateSettingsResponse {
		if v != nil {
			return *v
		}
		var ret SqlWorkloadTypeUpdateSettingsResponse
		return ret
	}).(SqlWorkloadTypeUpdateSettingsResponseOutput)
}

func (o SqlWorkloadTypeUpdateSettingsResponsePtrOutput) SqlWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SqlWorkloadTypeUpdateSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlWorkloadType
	}).(pulumi.StringPtrOutput)
}

type StorageConfigurationSettings struct {
	DiskConfigurationType *string             `pulumi:"diskConfigurationType"`
	SqlDataSettings       *SQLStorageSettings `pulumi:"sqlDataSettings"`
	SqlLogSettings        *SQLStorageSettings `pulumi:"sqlLogSettings"`
	SqlSystemDbOnDataDisk *bool               `pulumi:"sqlSystemDbOnDataDisk"`
	SqlTempDbSettings     *SQLTempDbSettings  `pulumi:"sqlTempDbSettings"`
	StorageWorkloadType   *string             `pulumi:"storageWorkloadType"`
}





type StorageConfigurationSettingsInput interface {
	pulumi.Input

	ToStorageConfigurationSettingsOutput() StorageConfigurationSettingsOutput
	ToStorageConfigurationSettingsOutputWithContext(context.Context) StorageConfigurationSettingsOutput
}

type StorageConfigurationSettingsArgs struct {
	DiskConfigurationType pulumi.StringPtrInput      `pulumi:"diskConfigurationType"`
	SqlDataSettings       SQLStorageSettingsPtrInput `pulumi:"sqlDataSettings"`
	SqlLogSettings        SQLStorageSettingsPtrInput `pulumi:"sqlLogSettings"`
	SqlSystemDbOnDataDisk pulumi.BoolPtrInput        `pulumi:"sqlSystemDbOnDataDisk"`
	SqlTempDbSettings     SQLTempDbSettingsPtrInput  `pulumi:"sqlTempDbSettings"`
	StorageWorkloadType   pulumi.StringPtrInput      `pulumi:"storageWorkloadType"`
}

func (StorageConfigurationSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageConfigurationSettings)(nil)).Elem()
}

func (i StorageConfigurationSettingsArgs) ToStorageConfigurationSettingsOutput() StorageConfigurationSettingsOutput {
	return i.ToStorageConfigurationSettingsOutputWithContext(context.Background())
}

func (i StorageConfigurationSettingsArgs) ToStorageConfigurationSettingsOutputWithContext(ctx context.Context) StorageConfigurationSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageConfigurationSettingsOutput)
}

func (i StorageConfigurationSettingsArgs) ToStorageConfigurationSettingsPtrOutput() StorageConfigurationSettingsPtrOutput {
	return i.ToStorageConfigurationSettingsPtrOutputWithContext(context.Background())
}

func (i StorageConfigurationSettingsArgs) ToStorageConfigurationSettingsPtrOutputWithContext(ctx context.Context) StorageConfigurationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageConfigurationSettingsOutput).ToStorageConfigurationSettingsPtrOutputWithContext(ctx)
}









type StorageConfigurationSettingsPtrInput interface {
	pulumi.Input

	ToStorageConfigurationSettingsPtrOutput() StorageConfigurationSettingsPtrOutput
	ToStorageConfigurationSettingsPtrOutputWithContext(context.Context) StorageConfigurationSettingsPtrOutput
}

type storageConfigurationSettingsPtrType StorageConfigurationSettingsArgs

func StorageConfigurationSettingsPtr(v *StorageConfigurationSettingsArgs) StorageConfigurationSettingsPtrInput {
	return (*storageConfigurationSettingsPtrType)(v)
}

func (*storageConfigurationSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageConfigurationSettings)(nil)).Elem()
}

func (i *storageConfigurationSettingsPtrType) ToStorageConfigurationSettingsPtrOutput() StorageConfigurationSettingsPtrOutput {
	return i.ToStorageConfigurationSettingsPtrOutputWithContext(context.Background())
}

func (i *storageConfigurationSettingsPtrType) ToStorageConfigurationSettingsPtrOutputWithContext(ctx context.Context) StorageConfigurationSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StorageConfigurationSettingsPtrOutput)
}

type StorageConfigurationSettingsOutput struct{ *pulumi.OutputState }

func (StorageConfigurationSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageConfigurationSettings)(nil)).Elem()
}

func (o StorageConfigurationSettingsOutput) ToStorageConfigurationSettingsOutput() StorageConfigurationSettingsOutput {
	return o
}

func (o StorageConfigurationSettingsOutput) ToStorageConfigurationSettingsOutputWithContext(ctx context.Context) StorageConfigurationSettingsOutput {
	return o
}

func (o StorageConfigurationSettingsOutput) ToStorageConfigurationSettingsPtrOutput() StorageConfigurationSettingsPtrOutput {
	return o.ToStorageConfigurationSettingsPtrOutputWithContext(context.Background())
}

func (o StorageConfigurationSettingsOutput) ToStorageConfigurationSettingsPtrOutputWithContext(ctx context.Context) StorageConfigurationSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StorageConfigurationSettings) *StorageConfigurationSettings {
		return &v
	}).(StorageConfigurationSettingsPtrOutput)
}

func (o StorageConfigurationSettingsOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettings) *string { return v.DiskConfigurationType }).(pulumi.StringPtrOutput)
}

func (o StorageConfigurationSettingsOutput) SqlDataSettings() SQLStorageSettingsPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettings) *SQLStorageSettings { return v.SqlDataSettings }).(SQLStorageSettingsPtrOutput)
}

func (o StorageConfigurationSettingsOutput) SqlLogSettings() SQLStorageSettingsPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettings) *SQLStorageSettings { return v.SqlLogSettings }).(SQLStorageSettingsPtrOutput)
}

func (o StorageConfigurationSettingsOutput) SqlSystemDbOnDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettings) *bool { return v.SqlSystemDbOnDataDisk }).(pulumi.BoolPtrOutput)
}

func (o StorageConfigurationSettingsOutput) SqlTempDbSettings() SQLTempDbSettingsPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettings) *SQLTempDbSettings { return v.SqlTempDbSettings }).(SQLTempDbSettingsPtrOutput)
}

func (o StorageConfigurationSettingsOutput) StorageWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettings) *string { return v.StorageWorkloadType }).(pulumi.StringPtrOutput)
}

type StorageConfigurationSettingsPtrOutput struct{ *pulumi.OutputState }

func (StorageConfigurationSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageConfigurationSettings)(nil)).Elem()
}

func (o StorageConfigurationSettingsPtrOutput) ToStorageConfigurationSettingsPtrOutput() StorageConfigurationSettingsPtrOutput {
	return o
}

func (o StorageConfigurationSettingsPtrOutput) ToStorageConfigurationSettingsPtrOutputWithContext(ctx context.Context) StorageConfigurationSettingsPtrOutput {
	return o
}

func (o StorageConfigurationSettingsPtrOutput) Elem() StorageConfigurationSettingsOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) StorageConfigurationSettings {
		if v != nil {
			return *v
		}
		var ret StorageConfigurationSettings
		return ret
	}).(StorageConfigurationSettingsOutput)
}

func (o StorageConfigurationSettingsPtrOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) *string {
		if v == nil {
			return nil
		}
		return v.DiskConfigurationType
	}).(pulumi.StringPtrOutput)
}

func (o StorageConfigurationSettingsPtrOutput) SqlDataSettings() SQLStorageSettingsPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) *SQLStorageSettings {
		if v == nil {
			return nil
		}
		return v.SqlDataSettings
	}).(SQLStorageSettingsPtrOutput)
}

func (o StorageConfigurationSettingsPtrOutput) SqlLogSettings() SQLStorageSettingsPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) *SQLStorageSettings {
		if v == nil {
			return nil
		}
		return v.SqlLogSettings
	}).(SQLStorageSettingsPtrOutput)
}

func (o StorageConfigurationSettingsPtrOutput) SqlSystemDbOnDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SqlSystemDbOnDataDisk
	}).(pulumi.BoolPtrOutput)
}

func (o StorageConfigurationSettingsPtrOutput) SqlTempDbSettings() SQLTempDbSettingsPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) *SQLTempDbSettings {
		if v == nil {
			return nil
		}
		return v.SqlTempDbSettings
	}).(SQLTempDbSettingsPtrOutput)
}

func (o StorageConfigurationSettingsPtrOutput) StorageWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettings) *string {
		if v == nil {
			return nil
		}
		return v.StorageWorkloadType
	}).(pulumi.StringPtrOutput)
}

type StorageConfigurationSettingsResponse struct {
	DiskConfigurationType *string                     `pulumi:"diskConfigurationType"`
	SqlDataSettings       *SQLStorageSettingsResponse `pulumi:"sqlDataSettings"`
	SqlLogSettings        *SQLStorageSettingsResponse `pulumi:"sqlLogSettings"`
	SqlSystemDbOnDataDisk *bool                       `pulumi:"sqlSystemDbOnDataDisk"`
	SqlTempDbSettings     *SQLTempDbSettingsResponse  `pulumi:"sqlTempDbSettings"`
	StorageWorkloadType   *string                     `pulumi:"storageWorkloadType"`
}

type StorageConfigurationSettingsResponseOutput struct{ *pulumi.OutputState }

func (StorageConfigurationSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StorageConfigurationSettingsResponse)(nil)).Elem()
}

func (o StorageConfigurationSettingsResponseOutput) ToStorageConfigurationSettingsResponseOutput() StorageConfigurationSettingsResponseOutput {
	return o
}

func (o StorageConfigurationSettingsResponseOutput) ToStorageConfigurationSettingsResponseOutputWithContext(ctx context.Context) StorageConfigurationSettingsResponseOutput {
	return o
}

func (o StorageConfigurationSettingsResponseOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettingsResponse) *string { return v.DiskConfigurationType }).(pulumi.StringPtrOutput)
}

func (o StorageConfigurationSettingsResponseOutput) SqlDataSettings() SQLStorageSettingsResponsePtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettingsResponse) *SQLStorageSettingsResponse { return v.SqlDataSettings }).(SQLStorageSettingsResponsePtrOutput)
}

func (o StorageConfigurationSettingsResponseOutput) SqlLogSettings() SQLStorageSettingsResponsePtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettingsResponse) *SQLStorageSettingsResponse { return v.SqlLogSettings }).(SQLStorageSettingsResponsePtrOutput)
}

func (o StorageConfigurationSettingsResponseOutput) SqlSystemDbOnDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettingsResponse) *bool { return v.SqlSystemDbOnDataDisk }).(pulumi.BoolPtrOutput)
}

func (o StorageConfigurationSettingsResponseOutput) SqlTempDbSettings() SQLTempDbSettingsResponsePtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettingsResponse) *SQLTempDbSettingsResponse { return v.SqlTempDbSettings }).(SQLTempDbSettingsResponsePtrOutput)
}

func (o StorageConfigurationSettingsResponseOutput) StorageWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StorageConfigurationSettingsResponse) *string { return v.StorageWorkloadType }).(pulumi.StringPtrOutput)
}

type StorageConfigurationSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (StorageConfigurationSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StorageConfigurationSettingsResponse)(nil)).Elem()
}

func (o StorageConfigurationSettingsResponsePtrOutput) ToStorageConfigurationSettingsResponsePtrOutput() StorageConfigurationSettingsResponsePtrOutput {
	return o
}

func (o StorageConfigurationSettingsResponsePtrOutput) ToStorageConfigurationSettingsResponsePtrOutputWithContext(ctx context.Context) StorageConfigurationSettingsResponsePtrOutput {
	return o
}

func (o StorageConfigurationSettingsResponsePtrOutput) Elem() StorageConfigurationSettingsResponseOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) StorageConfigurationSettingsResponse {
		if v != nil {
			return *v
		}
		var ret StorageConfigurationSettingsResponse
		return ret
	}).(StorageConfigurationSettingsResponseOutput)
}

func (o StorageConfigurationSettingsResponsePtrOutput) DiskConfigurationType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DiskConfigurationType
	}).(pulumi.StringPtrOutput)
}

func (o StorageConfigurationSettingsResponsePtrOutput) SqlDataSettings() SQLStorageSettingsResponsePtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) *SQLStorageSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlDataSettings
	}).(SQLStorageSettingsResponsePtrOutput)
}

func (o StorageConfigurationSettingsResponsePtrOutput) SqlLogSettings() SQLStorageSettingsResponsePtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) *SQLStorageSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlLogSettings
	}).(SQLStorageSettingsResponsePtrOutput)
}

func (o StorageConfigurationSettingsResponsePtrOutput) SqlSystemDbOnDataDisk() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SqlSystemDbOnDataDisk
	}).(pulumi.BoolPtrOutput)
}

func (o StorageConfigurationSettingsResponsePtrOutput) SqlTempDbSettings() SQLTempDbSettingsResponsePtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) *SQLTempDbSettingsResponse {
		if v == nil {
			return nil
		}
		return v.SqlTempDbSettings
	}).(SQLTempDbSettingsResponsePtrOutput)
}

func (o StorageConfigurationSettingsResponsePtrOutput) StorageWorkloadType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StorageConfigurationSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageWorkloadType
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

type TroubleshootingAdditionalPropertiesResponse struct {
	UnhealthyReplicaInfo *UnhealthyReplicaInfoResponse `pulumi:"unhealthyReplicaInfo"`
}

type TroubleshootingAdditionalPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TroubleshootingAdditionalPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TroubleshootingAdditionalPropertiesResponse)(nil)).Elem()
}

func (o TroubleshootingAdditionalPropertiesResponseOutput) ToTroubleshootingAdditionalPropertiesResponseOutput() TroubleshootingAdditionalPropertiesResponseOutput {
	return o
}

func (o TroubleshootingAdditionalPropertiesResponseOutput) ToTroubleshootingAdditionalPropertiesResponseOutputWithContext(ctx context.Context) TroubleshootingAdditionalPropertiesResponseOutput {
	return o
}

func (o TroubleshootingAdditionalPropertiesResponseOutput) UnhealthyReplicaInfo() UnhealthyReplicaInfoResponsePtrOutput {
	return o.ApplyT(func(v TroubleshootingAdditionalPropertiesResponse) *UnhealthyReplicaInfoResponse {
		return v.UnhealthyReplicaInfo
	}).(UnhealthyReplicaInfoResponsePtrOutput)
}

type TroubleshootingStatusResponse struct {
	EndTimeUtc              string                                      `pulumi:"endTimeUtc"`
	LastTriggerTimeUtc      string                                      `pulumi:"lastTriggerTimeUtc"`
	Properties              TroubleshootingAdditionalPropertiesResponse `pulumi:"properties"`
	RootCause               string                                      `pulumi:"rootCause"`
	StartTimeUtc            string                                      `pulumi:"startTimeUtc"`
	TroubleshootingScenario string                                      `pulumi:"troubleshootingScenario"`
}


func (val *TroubleshootingStatusResponse) Defaults() *TroubleshootingStatusResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.TroubleshootingScenario) {
		tmp.TroubleshootingScenario = "UnhealthyReplica"
	}
	return &tmp
}

type TroubleshootingStatusResponseOutput struct{ *pulumi.OutputState }

func (TroubleshootingStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TroubleshootingStatusResponse)(nil)).Elem()
}

func (o TroubleshootingStatusResponseOutput) ToTroubleshootingStatusResponseOutput() TroubleshootingStatusResponseOutput {
	return o
}

func (o TroubleshootingStatusResponseOutput) ToTroubleshootingStatusResponseOutputWithContext(ctx context.Context) TroubleshootingStatusResponseOutput {
	return o
}

func (o TroubleshootingStatusResponseOutput) EndTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v TroubleshootingStatusResponse) string { return v.EndTimeUtc }).(pulumi.StringOutput)
}

func (o TroubleshootingStatusResponseOutput) LastTriggerTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v TroubleshootingStatusResponse) string { return v.LastTriggerTimeUtc }).(pulumi.StringOutput)
}

func (o TroubleshootingStatusResponseOutput) Properties() TroubleshootingAdditionalPropertiesResponseOutput {
	return o.ApplyT(func(v TroubleshootingStatusResponse) TroubleshootingAdditionalPropertiesResponse { return v.Properties }).(TroubleshootingAdditionalPropertiesResponseOutput)
}

func (o TroubleshootingStatusResponseOutput) RootCause() pulumi.StringOutput {
	return o.ApplyT(func(v TroubleshootingStatusResponse) string { return v.RootCause }).(pulumi.StringOutput)
}

func (o TroubleshootingStatusResponseOutput) StartTimeUtc() pulumi.StringOutput {
	return o.ApplyT(func(v TroubleshootingStatusResponse) string { return v.StartTimeUtc }).(pulumi.StringOutput)
}

func (o TroubleshootingStatusResponseOutput) TroubleshootingScenario() pulumi.StringOutput {
	return o.ApplyT(func(v TroubleshootingStatusResponse) string { return v.TroubleshootingScenario }).(pulumi.StringOutput)
}

type UnhealthyReplicaInfoResponse struct {
	AvailabilityGroupName *string `pulumi:"availabilityGroupName"`
}

type UnhealthyReplicaInfoResponseOutput struct{ *pulumi.OutputState }

func (UnhealthyReplicaInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UnhealthyReplicaInfoResponse)(nil)).Elem()
}

func (o UnhealthyReplicaInfoResponseOutput) ToUnhealthyReplicaInfoResponseOutput() UnhealthyReplicaInfoResponseOutput {
	return o
}

func (o UnhealthyReplicaInfoResponseOutput) ToUnhealthyReplicaInfoResponseOutputWithContext(ctx context.Context) UnhealthyReplicaInfoResponseOutput {
	return o
}

func (o UnhealthyReplicaInfoResponseOutput) AvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v UnhealthyReplicaInfoResponse) *string { return v.AvailabilityGroupName }).(pulumi.StringPtrOutput)
}

type UnhealthyReplicaInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (UnhealthyReplicaInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**UnhealthyReplicaInfoResponse)(nil)).Elem()
}

func (o UnhealthyReplicaInfoResponsePtrOutput) ToUnhealthyReplicaInfoResponsePtrOutput() UnhealthyReplicaInfoResponsePtrOutput {
	return o
}

func (o UnhealthyReplicaInfoResponsePtrOutput) ToUnhealthyReplicaInfoResponsePtrOutputWithContext(ctx context.Context) UnhealthyReplicaInfoResponsePtrOutput {
	return o
}

func (o UnhealthyReplicaInfoResponsePtrOutput) Elem() UnhealthyReplicaInfoResponseOutput {
	return o.ApplyT(func(v *UnhealthyReplicaInfoResponse) UnhealthyReplicaInfoResponse {
		if v != nil {
			return *v
		}
		var ret UnhealthyReplicaInfoResponse
		return ret
	}).(UnhealthyReplicaInfoResponseOutput)
}

func (o UnhealthyReplicaInfoResponsePtrOutput) AvailabilityGroupName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *UnhealthyReplicaInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.AvailabilityGroupName
	}).(pulumi.StringPtrOutput)
}

type WsfcDomainCredentials struct {
	ClusterBootstrapAccountPassword *string `pulumi:"clusterBootstrapAccountPassword"`
	ClusterOperatorAccountPassword  *string `pulumi:"clusterOperatorAccountPassword"`
	SqlServiceAccountPassword       *string `pulumi:"sqlServiceAccountPassword"`
}





type WsfcDomainCredentialsInput interface {
	pulumi.Input

	ToWsfcDomainCredentialsOutput() WsfcDomainCredentialsOutput
	ToWsfcDomainCredentialsOutputWithContext(context.Context) WsfcDomainCredentialsOutput
}

type WsfcDomainCredentialsArgs struct {
	ClusterBootstrapAccountPassword pulumi.StringPtrInput `pulumi:"clusterBootstrapAccountPassword"`
	ClusterOperatorAccountPassword  pulumi.StringPtrInput `pulumi:"clusterOperatorAccountPassword"`
	SqlServiceAccountPassword       pulumi.StringPtrInput `pulumi:"sqlServiceAccountPassword"`
}

func (WsfcDomainCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WsfcDomainCredentials)(nil)).Elem()
}

func (i WsfcDomainCredentialsArgs) ToWsfcDomainCredentialsOutput() WsfcDomainCredentialsOutput {
	return i.ToWsfcDomainCredentialsOutputWithContext(context.Background())
}

func (i WsfcDomainCredentialsArgs) ToWsfcDomainCredentialsOutputWithContext(ctx context.Context) WsfcDomainCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsfcDomainCredentialsOutput)
}

func (i WsfcDomainCredentialsArgs) ToWsfcDomainCredentialsPtrOutput() WsfcDomainCredentialsPtrOutput {
	return i.ToWsfcDomainCredentialsPtrOutputWithContext(context.Background())
}

func (i WsfcDomainCredentialsArgs) ToWsfcDomainCredentialsPtrOutputWithContext(ctx context.Context) WsfcDomainCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsfcDomainCredentialsOutput).ToWsfcDomainCredentialsPtrOutputWithContext(ctx)
}









type WsfcDomainCredentialsPtrInput interface {
	pulumi.Input

	ToWsfcDomainCredentialsPtrOutput() WsfcDomainCredentialsPtrOutput
	ToWsfcDomainCredentialsPtrOutputWithContext(context.Context) WsfcDomainCredentialsPtrOutput
}

type wsfcDomainCredentialsPtrType WsfcDomainCredentialsArgs

func WsfcDomainCredentialsPtr(v *WsfcDomainCredentialsArgs) WsfcDomainCredentialsPtrInput {
	return (*wsfcDomainCredentialsPtrType)(v)
}

func (*wsfcDomainCredentialsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WsfcDomainCredentials)(nil)).Elem()
}

func (i *wsfcDomainCredentialsPtrType) ToWsfcDomainCredentialsPtrOutput() WsfcDomainCredentialsPtrOutput {
	return i.ToWsfcDomainCredentialsPtrOutputWithContext(context.Background())
}

func (i *wsfcDomainCredentialsPtrType) ToWsfcDomainCredentialsPtrOutputWithContext(ctx context.Context) WsfcDomainCredentialsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsfcDomainCredentialsPtrOutput)
}

type WsfcDomainCredentialsOutput struct{ *pulumi.OutputState }

func (WsfcDomainCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsfcDomainCredentials)(nil)).Elem()
}

func (o WsfcDomainCredentialsOutput) ToWsfcDomainCredentialsOutput() WsfcDomainCredentialsOutput {
	return o
}

func (o WsfcDomainCredentialsOutput) ToWsfcDomainCredentialsOutputWithContext(ctx context.Context) WsfcDomainCredentialsOutput {
	return o
}

func (o WsfcDomainCredentialsOutput) ToWsfcDomainCredentialsPtrOutput() WsfcDomainCredentialsPtrOutput {
	return o.ToWsfcDomainCredentialsPtrOutputWithContext(context.Background())
}

func (o WsfcDomainCredentialsOutput) ToWsfcDomainCredentialsPtrOutputWithContext(ctx context.Context) WsfcDomainCredentialsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WsfcDomainCredentials) *WsfcDomainCredentials {
		return &v
	}).(WsfcDomainCredentialsPtrOutput)
}

func (o WsfcDomainCredentialsOutput) ClusterBootstrapAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainCredentials) *string { return v.ClusterBootstrapAccountPassword }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsOutput) ClusterOperatorAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainCredentials) *string { return v.ClusterOperatorAccountPassword }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsOutput) SqlServiceAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainCredentials) *string { return v.SqlServiceAccountPassword }).(pulumi.StringPtrOutput)
}

type WsfcDomainCredentialsPtrOutput struct{ *pulumi.OutputState }

func (WsfcDomainCredentialsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsfcDomainCredentials)(nil)).Elem()
}

func (o WsfcDomainCredentialsPtrOutput) ToWsfcDomainCredentialsPtrOutput() WsfcDomainCredentialsPtrOutput {
	return o
}

func (o WsfcDomainCredentialsPtrOutput) ToWsfcDomainCredentialsPtrOutputWithContext(ctx context.Context) WsfcDomainCredentialsPtrOutput {
	return o
}

func (o WsfcDomainCredentialsPtrOutput) Elem() WsfcDomainCredentialsOutput {
	return o.ApplyT(func(v *WsfcDomainCredentials) WsfcDomainCredentials {
		if v != nil {
			return *v
		}
		var ret WsfcDomainCredentials
		return ret
	}).(WsfcDomainCredentialsOutput)
}

func (o WsfcDomainCredentialsPtrOutput) ClusterBootstrapAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainCredentials) *string {
		if v == nil {
			return nil
		}
		return v.ClusterBootstrapAccountPassword
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsPtrOutput) ClusterOperatorAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainCredentials) *string {
		if v == nil {
			return nil
		}
		return v.ClusterOperatorAccountPassword
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsPtrOutput) SqlServiceAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainCredentials) *string {
		if v == nil {
			return nil
		}
		return v.SqlServiceAccountPassword
	}).(pulumi.StringPtrOutput)
}

type WsfcDomainCredentialsResponse struct {
	ClusterBootstrapAccountPassword *string `pulumi:"clusterBootstrapAccountPassword"`
	ClusterOperatorAccountPassword  *string `pulumi:"clusterOperatorAccountPassword"`
	SqlServiceAccountPassword       *string `pulumi:"sqlServiceAccountPassword"`
}

type WsfcDomainCredentialsResponseOutput struct{ *pulumi.OutputState }

func (WsfcDomainCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsfcDomainCredentialsResponse)(nil)).Elem()
}

func (o WsfcDomainCredentialsResponseOutput) ToWsfcDomainCredentialsResponseOutput() WsfcDomainCredentialsResponseOutput {
	return o
}

func (o WsfcDomainCredentialsResponseOutput) ToWsfcDomainCredentialsResponseOutputWithContext(ctx context.Context) WsfcDomainCredentialsResponseOutput {
	return o
}

func (o WsfcDomainCredentialsResponseOutput) ClusterBootstrapAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainCredentialsResponse) *string { return v.ClusterBootstrapAccountPassword }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsResponseOutput) ClusterOperatorAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainCredentialsResponse) *string { return v.ClusterOperatorAccountPassword }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsResponseOutput) SqlServiceAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainCredentialsResponse) *string { return v.SqlServiceAccountPassword }).(pulumi.StringPtrOutput)
}

type WsfcDomainCredentialsResponsePtrOutput struct{ *pulumi.OutputState }

func (WsfcDomainCredentialsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsfcDomainCredentialsResponse)(nil)).Elem()
}

func (o WsfcDomainCredentialsResponsePtrOutput) ToWsfcDomainCredentialsResponsePtrOutput() WsfcDomainCredentialsResponsePtrOutput {
	return o
}

func (o WsfcDomainCredentialsResponsePtrOutput) ToWsfcDomainCredentialsResponsePtrOutputWithContext(ctx context.Context) WsfcDomainCredentialsResponsePtrOutput {
	return o
}

func (o WsfcDomainCredentialsResponsePtrOutput) Elem() WsfcDomainCredentialsResponseOutput {
	return o.ApplyT(func(v *WsfcDomainCredentialsResponse) WsfcDomainCredentialsResponse {
		if v != nil {
			return *v
		}
		var ret WsfcDomainCredentialsResponse
		return ret
	}).(WsfcDomainCredentialsResponseOutput)
}

func (o WsfcDomainCredentialsResponsePtrOutput) ClusterBootstrapAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterBootstrapAccountPassword
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsResponsePtrOutput) ClusterOperatorAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterOperatorAccountPassword
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainCredentialsResponsePtrOutput) SqlServiceAccountPassword() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainCredentialsResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlServiceAccountPassword
	}).(pulumi.StringPtrOutput)
}

type WsfcDomainProfile struct {
	ClusterBootstrapAccount  *string `pulumi:"clusterBootstrapAccount"`
	ClusterOperatorAccount   *string `pulumi:"clusterOperatorAccount"`
	ClusterSubnetType        *string `pulumi:"clusterSubnetType"`
	DomainFqdn               *string `pulumi:"domainFqdn"`
	FileShareWitnessPath     *string `pulumi:"fileShareWitnessPath"`
	OuPath                   *string `pulumi:"ouPath"`
	SqlServiceAccount        *string `pulumi:"sqlServiceAccount"`
	StorageAccountPrimaryKey *string `pulumi:"storageAccountPrimaryKey"`
	StorageAccountUrl        *string `pulumi:"storageAccountUrl"`
}





type WsfcDomainProfileInput interface {
	pulumi.Input

	ToWsfcDomainProfileOutput() WsfcDomainProfileOutput
	ToWsfcDomainProfileOutputWithContext(context.Context) WsfcDomainProfileOutput
}

type WsfcDomainProfileArgs struct {
	ClusterBootstrapAccount  pulumi.StringPtrInput `pulumi:"clusterBootstrapAccount"`
	ClusterOperatorAccount   pulumi.StringPtrInput `pulumi:"clusterOperatorAccount"`
	ClusterSubnetType        pulumi.StringPtrInput `pulumi:"clusterSubnetType"`
	DomainFqdn               pulumi.StringPtrInput `pulumi:"domainFqdn"`
	FileShareWitnessPath     pulumi.StringPtrInput `pulumi:"fileShareWitnessPath"`
	OuPath                   pulumi.StringPtrInput `pulumi:"ouPath"`
	SqlServiceAccount        pulumi.StringPtrInput `pulumi:"sqlServiceAccount"`
	StorageAccountPrimaryKey pulumi.StringPtrInput `pulumi:"storageAccountPrimaryKey"`
	StorageAccountUrl        pulumi.StringPtrInput `pulumi:"storageAccountUrl"`
}

func (WsfcDomainProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WsfcDomainProfile)(nil)).Elem()
}

func (i WsfcDomainProfileArgs) ToWsfcDomainProfileOutput() WsfcDomainProfileOutput {
	return i.ToWsfcDomainProfileOutputWithContext(context.Background())
}

func (i WsfcDomainProfileArgs) ToWsfcDomainProfileOutputWithContext(ctx context.Context) WsfcDomainProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsfcDomainProfileOutput)
}

func (i WsfcDomainProfileArgs) ToWsfcDomainProfilePtrOutput() WsfcDomainProfilePtrOutput {
	return i.ToWsfcDomainProfilePtrOutputWithContext(context.Background())
}

func (i WsfcDomainProfileArgs) ToWsfcDomainProfilePtrOutputWithContext(ctx context.Context) WsfcDomainProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsfcDomainProfileOutput).ToWsfcDomainProfilePtrOutputWithContext(ctx)
}









type WsfcDomainProfilePtrInput interface {
	pulumi.Input

	ToWsfcDomainProfilePtrOutput() WsfcDomainProfilePtrOutput
	ToWsfcDomainProfilePtrOutputWithContext(context.Context) WsfcDomainProfilePtrOutput
}

type wsfcDomainProfilePtrType WsfcDomainProfileArgs

func WsfcDomainProfilePtr(v *WsfcDomainProfileArgs) WsfcDomainProfilePtrInput {
	return (*wsfcDomainProfilePtrType)(v)
}

func (*wsfcDomainProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WsfcDomainProfile)(nil)).Elem()
}

func (i *wsfcDomainProfilePtrType) ToWsfcDomainProfilePtrOutput() WsfcDomainProfilePtrOutput {
	return i.ToWsfcDomainProfilePtrOutputWithContext(context.Background())
}

func (i *wsfcDomainProfilePtrType) ToWsfcDomainProfilePtrOutputWithContext(ctx context.Context) WsfcDomainProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WsfcDomainProfilePtrOutput)
}

type WsfcDomainProfileOutput struct{ *pulumi.OutputState }

func (WsfcDomainProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsfcDomainProfile)(nil)).Elem()
}

func (o WsfcDomainProfileOutput) ToWsfcDomainProfileOutput() WsfcDomainProfileOutput {
	return o
}

func (o WsfcDomainProfileOutput) ToWsfcDomainProfileOutputWithContext(ctx context.Context) WsfcDomainProfileOutput {
	return o
}

func (o WsfcDomainProfileOutput) ToWsfcDomainProfilePtrOutput() WsfcDomainProfilePtrOutput {
	return o.ToWsfcDomainProfilePtrOutputWithContext(context.Background())
}

func (o WsfcDomainProfileOutput) ToWsfcDomainProfilePtrOutputWithContext(ctx context.Context) WsfcDomainProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WsfcDomainProfile) *WsfcDomainProfile {
		return &v
	}).(WsfcDomainProfilePtrOutput)
}

func (o WsfcDomainProfileOutput) ClusterBootstrapAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.ClusterBootstrapAccount }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) ClusterOperatorAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.ClusterOperatorAccount }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) ClusterSubnetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.ClusterSubnetType }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) DomainFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.DomainFqdn }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) FileShareWitnessPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.FileShareWitnessPath }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) OuPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.OuPath }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) SqlServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.SqlServiceAccount }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) StorageAccountPrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.StorageAccountPrimaryKey }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfile) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

type WsfcDomainProfilePtrOutput struct{ *pulumi.OutputState }

func (WsfcDomainProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsfcDomainProfile)(nil)).Elem()
}

func (o WsfcDomainProfilePtrOutput) ToWsfcDomainProfilePtrOutput() WsfcDomainProfilePtrOutput {
	return o
}

func (o WsfcDomainProfilePtrOutput) ToWsfcDomainProfilePtrOutputWithContext(ctx context.Context) WsfcDomainProfilePtrOutput {
	return o
}

func (o WsfcDomainProfilePtrOutput) Elem() WsfcDomainProfileOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) WsfcDomainProfile {
		if v != nil {
			return *v
		}
		var ret WsfcDomainProfile
		return ret
	}).(WsfcDomainProfileOutput)
}

func (o WsfcDomainProfilePtrOutput) ClusterBootstrapAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClusterBootstrapAccount
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) ClusterOperatorAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClusterOperatorAccount
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) ClusterSubnetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.ClusterSubnetType
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) DomainFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.DomainFqdn
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) FileShareWitnessPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.FileShareWitnessPath
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) OuPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.OuPath
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) SqlServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.SqlServiceAccount
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) StorageAccountPrimaryKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountPrimaryKey
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfilePtrOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfile) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountUrl
	}).(pulumi.StringPtrOutput)
}

type WsfcDomainProfileResponse struct {
	ClusterBootstrapAccount *string `pulumi:"clusterBootstrapAccount"`
	ClusterOperatorAccount  *string `pulumi:"clusterOperatorAccount"`
	ClusterSubnetType       *string `pulumi:"clusterSubnetType"`
	DomainFqdn              *string `pulumi:"domainFqdn"`
	FileShareWitnessPath    *string `pulumi:"fileShareWitnessPath"`
	OuPath                  *string `pulumi:"ouPath"`
	SqlServiceAccount       *string `pulumi:"sqlServiceAccount"`
	StorageAccountUrl       *string `pulumi:"storageAccountUrl"`
}

type WsfcDomainProfileResponseOutput struct{ *pulumi.OutputState }

func (WsfcDomainProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WsfcDomainProfileResponse)(nil)).Elem()
}

func (o WsfcDomainProfileResponseOutput) ToWsfcDomainProfileResponseOutput() WsfcDomainProfileResponseOutput {
	return o
}

func (o WsfcDomainProfileResponseOutput) ToWsfcDomainProfileResponseOutputWithContext(ctx context.Context) WsfcDomainProfileResponseOutput {
	return o
}

func (o WsfcDomainProfileResponseOutput) ClusterBootstrapAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.ClusterBootstrapAccount }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) ClusterOperatorAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.ClusterOperatorAccount }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) ClusterSubnetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.ClusterSubnetType }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) DomainFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.DomainFqdn }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) FileShareWitnessPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.FileShareWitnessPath }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) OuPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.OuPath }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) SqlServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.SqlServiceAccount }).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponseOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WsfcDomainProfileResponse) *string { return v.StorageAccountUrl }).(pulumi.StringPtrOutput)
}

type WsfcDomainProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (WsfcDomainProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WsfcDomainProfileResponse)(nil)).Elem()
}

func (o WsfcDomainProfileResponsePtrOutput) ToWsfcDomainProfileResponsePtrOutput() WsfcDomainProfileResponsePtrOutput {
	return o
}

func (o WsfcDomainProfileResponsePtrOutput) ToWsfcDomainProfileResponsePtrOutputWithContext(ctx context.Context) WsfcDomainProfileResponsePtrOutput {
	return o
}

func (o WsfcDomainProfileResponsePtrOutput) Elem() WsfcDomainProfileResponseOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) WsfcDomainProfileResponse {
		if v != nil {
			return *v
		}
		var ret WsfcDomainProfileResponse
		return ret
	}).(WsfcDomainProfileResponseOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) ClusterBootstrapAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterBootstrapAccount
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) ClusterOperatorAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterOperatorAccount
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) ClusterSubnetType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClusterSubnetType
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) DomainFqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.DomainFqdn
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) FileShareWitnessPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.FileShareWitnessPath
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) OuPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.OuPath
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) SqlServiceAccount() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.SqlServiceAccount
	}).(pulumi.StringPtrOutput)
}

func (o WsfcDomainProfileResponsePtrOutput) StorageAccountUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WsfcDomainProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.StorageAccountUrl
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AADAuthenticationSettingsOutput{})
	pulumi.RegisterOutputType(AADAuthenticationSettingsPtrOutput{})
	pulumi.RegisterOutputType(AADAuthenticationSettingsResponseOutput{})
	pulumi.RegisterOutputType(AADAuthenticationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AdditionalFeaturesServerConfigurationsOutput{})
	pulumi.RegisterOutputType(AdditionalFeaturesServerConfigurationsPtrOutput{})
	pulumi.RegisterOutputType(AdditionalFeaturesServerConfigurationsResponseOutput{})
	pulumi.RegisterOutputType(AdditionalFeaturesServerConfigurationsResponsePtrOutput{})
	pulumi.RegisterOutputType(AgConfigurationOutput{})
	pulumi.RegisterOutputType(AgConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AgConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AgConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(AgReplicaOutput{})
	pulumi.RegisterOutputType(AgReplicaArrayOutput{})
	pulumi.RegisterOutputType(AgReplicaResponseOutput{})
	pulumi.RegisterOutputType(AgReplicaResponseArrayOutput{})
	pulumi.RegisterOutputType(AssessmentSettingsOutput{})
	pulumi.RegisterOutputType(AssessmentSettingsPtrOutput{})
	pulumi.RegisterOutputType(AssessmentSettingsResponseOutput{})
	pulumi.RegisterOutputType(AssessmentSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoBackupSettingsOutput{})
	pulumi.RegisterOutputType(AutoBackupSettingsPtrOutput{})
	pulumi.RegisterOutputType(AutoBackupSettingsResponseOutput{})
	pulumi.RegisterOutputType(AutoBackupSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoPatchingSettingsOutput{})
	pulumi.RegisterOutputType(AutoPatchingSettingsPtrOutput{})
	pulumi.RegisterOutputType(AutoPatchingSettingsResponseOutput{})
	pulumi.RegisterOutputType(AutoPatchingSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(KeyVaultCredentialSettingsOutput{})
	pulumi.RegisterOutputType(KeyVaultCredentialSettingsPtrOutput{})
	pulumi.RegisterOutputType(KeyVaultCredentialSettingsResponseOutput{})
	pulumi.RegisterOutputType(KeyVaultCredentialSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationArrayOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LoadBalancerConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(MultiSubnetIpConfigurationOutput{})
	pulumi.RegisterOutputType(MultiSubnetIpConfigurationArrayOutput{})
	pulumi.RegisterOutputType(MultiSubnetIpConfigurationResponseOutput{})
	pulumi.RegisterOutputType(MultiSubnetIpConfigurationResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateIPAddressOutput{})
	pulumi.RegisterOutputType(PrivateIPAddressPtrOutput{})
	pulumi.RegisterOutputType(PrivateIPAddressResponseOutput{})
	pulumi.RegisterOutputType(PrivateIPAddressResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityOutput{})
	pulumi.RegisterOutputType(ResourceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ResourceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(SQLInstanceSettingsOutput{})
	pulumi.RegisterOutputType(SQLInstanceSettingsPtrOutput{})
	pulumi.RegisterOutputType(SQLInstanceSettingsResponseOutput{})
	pulumi.RegisterOutputType(SQLInstanceSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SQLStorageSettingsOutput{})
	pulumi.RegisterOutputType(SQLStorageSettingsPtrOutput{})
	pulumi.RegisterOutputType(SQLStorageSettingsResponseOutput{})
	pulumi.RegisterOutputType(SQLStorageSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SQLTempDbSettingsOutput{})
	pulumi.RegisterOutputType(SQLTempDbSettingsPtrOutput{})
	pulumi.RegisterOutputType(SQLTempDbSettingsResponseOutput{})
	pulumi.RegisterOutputType(SQLTempDbSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ScheduleOutput{})
	pulumi.RegisterOutputType(SchedulePtrOutput{})
	pulumi.RegisterOutputType(ScheduleResponseOutput{})
	pulumi.RegisterOutputType(ScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(ServerConfigurationsManagementSettingsOutput{})
	pulumi.RegisterOutputType(ServerConfigurationsManagementSettingsPtrOutput{})
	pulumi.RegisterOutputType(ServerConfigurationsManagementSettingsResponseOutput{})
	pulumi.RegisterOutputType(ServerConfigurationsManagementSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlConnectivityUpdateSettingsOutput{})
	pulumi.RegisterOutputType(SqlConnectivityUpdateSettingsPtrOutput{})
	pulumi.RegisterOutputType(SqlConnectivityUpdateSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlConnectivityUpdateSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlStorageUpdateSettingsOutput{})
	pulumi.RegisterOutputType(SqlStorageUpdateSettingsPtrOutput{})
	pulumi.RegisterOutputType(SqlStorageUpdateSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlStorageUpdateSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SqlWorkloadTypeUpdateSettingsOutput{})
	pulumi.RegisterOutputType(SqlWorkloadTypeUpdateSettingsPtrOutput{})
	pulumi.RegisterOutputType(SqlWorkloadTypeUpdateSettingsResponseOutput{})
	pulumi.RegisterOutputType(SqlWorkloadTypeUpdateSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(StorageConfigurationSettingsOutput{})
	pulumi.RegisterOutputType(StorageConfigurationSettingsPtrOutput{})
	pulumi.RegisterOutputType(StorageConfigurationSettingsResponseOutput{})
	pulumi.RegisterOutputType(StorageConfigurationSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(SystemDataResponseOutput{})
	pulumi.RegisterOutputType(TroubleshootingAdditionalPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TroubleshootingStatusResponseOutput{})
	pulumi.RegisterOutputType(UnhealthyReplicaInfoResponseOutput{})
	pulumi.RegisterOutputType(UnhealthyReplicaInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(WsfcDomainCredentialsOutput{})
	pulumi.RegisterOutputType(WsfcDomainCredentialsPtrOutput{})
	pulumi.RegisterOutputType(WsfcDomainCredentialsResponseOutput{})
	pulumi.RegisterOutputType(WsfcDomainCredentialsResponsePtrOutput{})
	pulumi.RegisterOutputType(WsfcDomainProfileOutput{})
	pulumi.RegisterOutputType(WsfcDomainProfilePtrOutput{})
	pulumi.RegisterOutputType(WsfcDomainProfileResponseOutput{})
	pulumi.RegisterOutputType(WsfcDomainProfileResponsePtrOutput{})
}
