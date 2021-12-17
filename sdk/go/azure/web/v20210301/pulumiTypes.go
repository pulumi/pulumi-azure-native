


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AppLogsConfiguration struct {
	Destination               *string                    `pulumi:"destination"`
	LogAnalyticsConfiguration *LogAnalyticsConfiguration `pulumi:"logAnalyticsConfiguration"`
}





type AppLogsConfigurationInput interface {
	pulumi.Input

	ToAppLogsConfigurationOutput() AppLogsConfigurationOutput
	ToAppLogsConfigurationOutputWithContext(context.Context) AppLogsConfigurationOutput
}

type AppLogsConfigurationArgs struct {
	Destination               pulumi.StringPtrInput             `pulumi:"destination"`
	LogAnalyticsConfiguration LogAnalyticsConfigurationPtrInput `pulumi:"logAnalyticsConfiguration"`
}

func (AppLogsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfiguration)(nil)).Elem()
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationOutput() AppLogsConfigurationOutput {
	return i.ToAppLogsConfigurationOutputWithContext(context.Background())
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationOutputWithContext(ctx context.Context) AppLogsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationOutput)
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return i.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (i AppLogsConfigurationArgs) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationOutput).ToAppLogsConfigurationPtrOutputWithContext(ctx)
}









type AppLogsConfigurationPtrInput interface {
	pulumi.Input

	ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput
	ToAppLogsConfigurationPtrOutputWithContext(context.Context) AppLogsConfigurationPtrOutput
}

type appLogsConfigurationPtrType AppLogsConfigurationArgs

func AppLogsConfigurationPtr(v *AppLogsConfigurationArgs) AppLogsConfigurationPtrInput {
	return (*appLogsConfigurationPtrType)(v)
}

func (*appLogsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfiguration)(nil)).Elem()
}

func (i *appLogsConfigurationPtrType) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return i.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (i *appLogsConfigurationPtrType) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AppLogsConfigurationPtrOutput)
}

type AppLogsConfigurationOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfiguration)(nil)).Elem()
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationOutput() AppLogsConfigurationOutput {
	return o
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationOutputWithContext(ctx context.Context) AppLogsConfigurationOutput {
	return o
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return o.ToAppLogsConfigurationPtrOutputWithContext(context.Background())
}

func (o AppLogsConfigurationOutput) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AppLogsConfiguration) *AppLogsConfiguration {
		return &v
	}).(AppLogsConfigurationPtrOutput)
}

func (o AppLogsConfigurationOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppLogsConfiguration) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v AppLogsConfiguration) *LogAnalyticsConfiguration { return v.LogAnalyticsConfiguration }).(LogAnalyticsConfigurationPtrOutput)
}

type AppLogsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfiguration)(nil)).Elem()
}

func (o AppLogsConfigurationPtrOutput) ToAppLogsConfigurationPtrOutput() AppLogsConfigurationPtrOutput {
	return o
}

func (o AppLogsConfigurationPtrOutput) ToAppLogsConfigurationPtrOutputWithContext(ctx context.Context) AppLogsConfigurationPtrOutput {
	return o
}

func (o AppLogsConfigurationPtrOutput) Elem() AppLogsConfigurationOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) AppLogsConfiguration {
		if v != nil {
			return *v
		}
		var ret AppLogsConfiguration
		return ret
	}).(AppLogsConfigurationOutput)
}

func (o AppLogsConfigurationPtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationPtrOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationPtrOutput {
	return o.ApplyT(func(v *AppLogsConfiguration) *LogAnalyticsConfiguration {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationPtrOutput)
}

type AppLogsConfigurationResponse struct {
	Destination               *string                            `pulumi:"destination"`
	LogAnalyticsConfiguration *LogAnalyticsConfigurationResponse `pulumi:"logAnalyticsConfiguration"`
}

type AppLogsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AppLogsConfigurationResponse)(nil)).Elem()
}

func (o AppLogsConfigurationResponseOutput) ToAppLogsConfigurationResponseOutput() AppLogsConfigurationResponseOutput {
	return o
}

func (o AppLogsConfigurationResponseOutput) ToAppLogsConfigurationResponseOutputWithContext(ctx context.Context) AppLogsConfigurationResponseOutput {
	return o
}

func (o AppLogsConfigurationResponseOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AppLogsConfigurationResponse) *string { return v.Destination }).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationResponseOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v AppLogsConfigurationResponse) *LogAnalyticsConfigurationResponse {
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationResponsePtrOutput)
}

type AppLogsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AppLogsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AppLogsConfigurationResponse)(nil)).Elem()
}

func (o AppLogsConfigurationResponsePtrOutput) ToAppLogsConfigurationResponsePtrOutput() AppLogsConfigurationResponsePtrOutput {
	return o
}

func (o AppLogsConfigurationResponsePtrOutput) ToAppLogsConfigurationResponsePtrOutputWithContext(ctx context.Context) AppLogsConfigurationResponsePtrOutput {
	return o
}

func (o AppLogsConfigurationResponsePtrOutput) Elem() AppLogsConfigurationResponseOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) AppLogsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AppLogsConfigurationResponse
		return ret
	}).(AppLogsConfigurationResponseOutput)
}

func (o AppLogsConfigurationResponsePtrOutput) Destination() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Destination
	}).(pulumi.StringPtrOutput)
}

func (o AppLogsConfigurationResponsePtrOutput) LogAnalyticsConfiguration() LogAnalyticsConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *AppLogsConfigurationResponse) *LogAnalyticsConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.LogAnalyticsConfiguration
	}).(LogAnalyticsConfigurationResponsePtrOutput)
}

type ArcConfiguration struct {
	ArtifactStorageAccessMode    *string                `pulumi:"artifactStorageAccessMode"`
	ArtifactStorageClassName     *string                `pulumi:"artifactStorageClassName"`
	ArtifactStorageMountPath     *string                `pulumi:"artifactStorageMountPath"`
	ArtifactStorageNodeName      *string                `pulumi:"artifactStorageNodeName"`
	ArtifactsStorageType         *StorageType           `pulumi:"artifactsStorageType"`
	FrontEndServiceConfiguration *FrontEndConfiguration `pulumi:"frontEndServiceConfiguration"`
	KubeConfig                   *string                `pulumi:"kubeConfig"`
}





type ArcConfigurationInput interface {
	pulumi.Input

	ToArcConfigurationOutput() ArcConfigurationOutput
	ToArcConfigurationOutputWithContext(context.Context) ArcConfigurationOutput
}

type ArcConfigurationArgs struct {
	ArtifactStorageAccessMode    pulumi.StringPtrInput         `pulumi:"artifactStorageAccessMode"`
	ArtifactStorageClassName     pulumi.StringPtrInput         `pulumi:"artifactStorageClassName"`
	ArtifactStorageMountPath     pulumi.StringPtrInput         `pulumi:"artifactStorageMountPath"`
	ArtifactStorageNodeName      pulumi.StringPtrInput         `pulumi:"artifactStorageNodeName"`
	ArtifactsStorageType         StorageTypePtrInput           `pulumi:"artifactsStorageType"`
	FrontEndServiceConfiguration FrontEndConfigurationPtrInput `pulumi:"frontEndServiceConfiguration"`
	KubeConfig                   pulumi.StringPtrInput         `pulumi:"kubeConfig"`
}

func (ArcConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConfiguration)(nil)).Elem()
}

func (i ArcConfigurationArgs) ToArcConfigurationOutput() ArcConfigurationOutput {
	return i.ToArcConfigurationOutputWithContext(context.Background())
}

func (i ArcConfigurationArgs) ToArcConfigurationOutputWithContext(ctx context.Context) ArcConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConfigurationOutput)
}

func (i ArcConfigurationArgs) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return i.ToArcConfigurationPtrOutputWithContext(context.Background())
}

func (i ArcConfigurationArgs) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConfigurationOutput).ToArcConfigurationPtrOutputWithContext(ctx)
}









type ArcConfigurationPtrInput interface {
	pulumi.Input

	ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput
	ToArcConfigurationPtrOutputWithContext(context.Context) ArcConfigurationPtrOutput
}

type arcConfigurationPtrType ArcConfigurationArgs

func ArcConfigurationPtr(v *ArcConfigurationArgs) ArcConfigurationPtrInput {
	return (*arcConfigurationPtrType)(v)
}

func (*arcConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcConfiguration)(nil)).Elem()
}

func (i *arcConfigurationPtrType) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return i.ToArcConfigurationPtrOutputWithContext(context.Background())
}

func (i *arcConfigurationPtrType) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArcConfigurationPtrOutput)
}

type ArcConfigurationOutput struct{ *pulumi.OutputState }

func (ArcConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConfiguration)(nil)).Elem()
}

func (o ArcConfigurationOutput) ToArcConfigurationOutput() ArcConfigurationOutput {
	return o
}

func (o ArcConfigurationOutput) ToArcConfigurationOutputWithContext(ctx context.Context) ArcConfigurationOutput {
	return o
}

func (o ArcConfigurationOutput) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return o.ToArcConfigurationPtrOutputWithContext(context.Background())
}

func (o ArcConfigurationOutput) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArcConfiguration) *ArcConfiguration {
		return &v
	}).(ArcConfigurationPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageAccessMode }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageClassName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageMountPath }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.ArtifactStorageNodeName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationOutput) ArtifactsStorageType() StorageTypePtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *StorageType { return v.ArtifactsStorageType }).(StorageTypePtrOutput)
}

func (o ArcConfigurationOutput) FrontEndServiceConfiguration() FrontEndConfigurationPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *FrontEndConfiguration { return v.FrontEndServiceConfiguration }).(FrontEndConfigurationPtrOutput)
}

func (o ArcConfigurationOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfiguration) *string { return v.KubeConfig }).(pulumi.StringPtrOutput)
}

type ArcConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ArcConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcConfiguration)(nil)).Elem()
}

func (o ArcConfigurationPtrOutput) ToArcConfigurationPtrOutput() ArcConfigurationPtrOutput {
	return o
}

func (o ArcConfigurationPtrOutput) ToArcConfigurationPtrOutputWithContext(ctx context.Context) ArcConfigurationPtrOutput {
	return o
}

func (o ArcConfigurationPtrOutput) Elem() ArcConfigurationOutput {
	return o.ApplyT(func(v *ArcConfiguration) ArcConfiguration {
		if v != nil {
			return *v
		}
		var ret ArcConfiguration
		return ret
	}).(ArcConfigurationOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageAccessMode
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageClassName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageMountPath
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageNodeName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationPtrOutput) ArtifactsStorageType() StorageTypePtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *StorageType {
		if v == nil {
			return nil
		}
		return v.ArtifactsStorageType
	}).(StorageTypePtrOutput)
}

func (o ArcConfigurationPtrOutput) FrontEndServiceConfiguration() FrontEndConfigurationPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *FrontEndConfiguration {
		if v == nil {
			return nil
		}
		return v.FrontEndServiceConfiguration
	}).(FrontEndConfigurationPtrOutput)
}

func (o ArcConfigurationPtrOutput) KubeConfig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.KubeConfig
	}).(pulumi.StringPtrOutput)
}

type ArcConfigurationResponse struct {
	ArtifactStorageAccessMode    *string                        `pulumi:"artifactStorageAccessMode"`
	ArtifactStorageClassName     *string                        `pulumi:"artifactStorageClassName"`
	ArtifactStorageMountPath     *string                        `pulumi:"artifactStorageMountPath"`
	ArtifactStorageNodeName      *string                        `pulumi:"artifactStorageNodeName"`
	ArtifactsStorageType         *string                        `pulumi:"artifactsStorageType"`
	FrontEndServiceConfiguration *FrontEndConfigurationResponse `pulumi:"frontEndServiceConfiguration"`
}

type ArcConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ArcConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArcConfigurationResponse)(nil)).Elem()
}

func (o ArcConfigurationResponseOutput) ToArcConfigurationResponseOutput() ArcConfigurationResponseOutput {
	return o
}

func (o ArcConfigurationResponseOutput) ToArcConfigurationResponseOutputWithContext(ctx context.Context) ArcConfigurationResponseOutput {
	return o
}

func (o ArcConfigurationResponseOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageAccessMode }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageClassName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageMountPath }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactStorageNodeName }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) ArtifactsStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *string { return v.ArtifactsStorageType }).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponseOutput) FrontEndServiceConfiguration() FrontEndConfigurationResponsePtrOutput {
	return o.ApplyT(func(v ArcConfigurationResponse) *FrontEndConfigurationResponse { return v.FrontEndServiceConfiguration }).(FrontEndConfigurationResponsePtrOutput)
}

type ArcConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ArcConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArcConfigurationResponse)(nil)).Elem()
}

func (o ArcConfigurationResponsePtrOutput) ToArcConfigurationResponsePtrOutput() ArcConfigurationResponsePtrOutput {
	return o
}

func (o ArcConfigurationResponsePtrOutput) ToArcConfigurationResponsePtrOutputWithContext(ctx context.Context) ArcConfigurationResponsePtrOutput {
	return o
}

func (o ArcConfigurationResponsePtrOutput) Elem() ArcConfigurationResponseOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) ArcConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ArcConfigurationResponse
		return ret
	}).(ArcConfigurationResponseOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageAccessMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageAccessMode
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageClassName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageClassName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageMountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageMountPath
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactStorageNodeName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactStorageNodeName
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) ArtifactsStorageType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ArtifactsStorageType
	}).(pulumi.StringPtrOutput)
}

func (o ArcConfigurationResponsePtrOutput) FrontEndServiceConfiguration() FrontEndConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *ArcConfigurationResponse) *FrontEndConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.FrontEndServiceConfiguration
	}).(FrontEndConfigurationResponsePtrOutput)
}

type Configuration struct {
	ActiveRevisionsMode *string               `pulumi:"activeRevisionsMode"`
	Ingress             *Ingress              `pulumi:"ingress"`
	Registries          []RegistryCredentials `pulumi:"registries"`
	Secrets             []Secret              `pulumi:"secrets"`
}


func (val *Configuration) Defaults() *Configuration {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Ingress = tmp.Ingress.Defaults()

	return &tmp
}





type ConfigurationInput interface {
	pulumi.Input

	ToConfigurationOutput() ConfigurationOutput
	ToConfigurationOutputWithContext(context.Context) ConfigurationOutput
}

type ConfigurationArgs struct {
	ActiveRevisionsMode pulumi.StringPtrInput         `pulumi:"activeRevisionsMode"`
	Ingress             IngressPtrInput               `pulumi:"ingress"`
	Registries          RegistryCredentialsArrayInput `pulumi:"registries"`
	Secrets             SecretArrayInput              `pulumi:"secrets"`
}

func (ConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil)).Elem()
}

func (i ConfigurationArgs) ToConfigurationOutput() ConfigurationOutput {
	return i.ToConfigurationOutputWithContext(context.Background())
}

func (i ConfigurationArgs) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput)
}

func (i ConfigurationArgs) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i ConfigurationArgs) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationOutput).ToConfigurationPtrOutputWithContext(ctx)
}









type ConfigurationPtrInput interface {
	pulumi.Input

	ToConfigurationPtrOutput() ConfigurationPtrOutput
	ToConfigurationPtrOutputWithContext(context.Context) ConfigurationPtrOutput
}

type configurationPtrType ConfigurationArgs

func ConfigurationPtr(v *ConfigurationArgs) ConfigurationPtrInput {
	return (*configurationPtrType)(v)
}

func (*configurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (i *configurationPtrType) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return i.ToConfigurationPtrOutputWithContext(context.Background())
}

func (i *configurationPtrType) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConfigurationPtrOutput)
}

type ConfigurationOutput struct{ *pulumi.OutputState }

func (ConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Configuration)(nil)).Elem()
}

func (o ConfigurationOutput) ToConfigurationOutput() ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationOutputWithContext(ctx context.Context) ConfigurationOutput {
	return o
}

func (o ConfigurationOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o.ToConfigurationPtrOutputWithContext(context.Background())
}

func (o ConfigurationOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Configuration) *Configuration {
		return &v
	}).(ConfigurationPtrOutput)
}

func (o ConfigurationOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Configuration) *string { return v.ActiveRevisionsMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationOutput) Ingress() IngressPtrOutput {
	return o.ApplyT(func(v Configuration) *Ingress { return v.Ingress }).(IngressPtrOutput)
}

func (o ConfigurationOutput) Registries() RegistryCredentialsArrayOutput {
	return o.ApplyT(func(v Configuration) []RegistryCredentials { return v.Registries }).(RegistryCredentialsArrayOutput)
}

func (o ConfigurationOutput) Secrets() SecretArrayOutput {
	return o.ApplyT(func(v Configuration) []Secret { return v.Secrets }).(SecretArrayOutput)
}

type ConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Configuration)(nil)).Elem()
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutput() ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) ToConfigurationPtrOutputWithContext(ctx context.Context) ConfigurationPtrOutput {
	return o
}

func (o ConfigurationPtrOutput) Elem() ConfigurationOutput {
	return o.ApplyT(func(v *Configuration) Configuration {
		if v != nil {
			return *v
		}
		var ret Configuration
		return ret
	}).(ConfigurationOutput)
}

func (o ConfigurationPtrOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Configuration) *string {
		if v == nil {
			return nil
		}
		return v.ActiveRevisionsMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationPtrOutput) Ingress() IngressPtrOutput {
	return o.ApplyT(func(v *Configuration) *Ingress {
		if v == nil {
			return nil
		}
		return v.Ingress
	}).(IngressPtrOutput)
}

func (o ConfigurationPtrOutput) Registries() RegistryCredentialsArrayOutput {
	return o.ApplyT(func(v *Configuration) []RegistryCredentials {
		if v == nil {
			return nil
		}
		return v.Registries
	}).(RegistryCredentialsArrayOutput)
}

func (o ConfigurationPtrOutput) Secrets() SecretArrayOutput {
	return o.ApplyT(func(v *Configuration) []Secret {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(SecretArrayOutput)
}

type ConfigurationResponse struct {
	ActiveRevisionsMode *string                       `pulumi:"activeRevisionsMode"`
	Ingress             *IngressResponse              `pulumi:"ingress"`
	Registries          []RegistryCredentialsResponse `pulumi:"registries"`
	Secrets             []SecretResponse              `pulumi:"secrets"`
}


func (val *ConfigurationResponse) Defaults() *ConfigurationResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Ingress = tmp.Ingress.Defaults()

	return &tmp
}

type ConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConfigurationResponse)(nil)).Elem()
}

func (o ConfigurationResponseOutput) ToConfigurationResponseOutput() ConfigurationResponseOutput {
	return o
}

func (o ConfigurationResponseOutput) ToConfigurationResponseOutputWithContext(ctx context.Context) ConfigurationResponseOutput {
	return o
}

func (o ConfigurationResponseOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConfigurationResponse) *string { return v.ActiveRevisionsMode }).(pulumi.StringPtrOutput)
}

func (o ConfigurationResponseOutput) Ingress() IngressResponsePtrOutput {
	return o.ApplyT(func(v ConfigurationResponse) *IngressResponse { return v.Ingress }).(IngressResponsePtrOutput)
}

func (o ConfigurationResponseOutput) Registries() RegistryCredentialsResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []RegistryCredentialsResponse { return v.Registries }).(RegistryCredentialsResponseArrayOutput)
}

func (o ConfigurationResponseOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v ConfigurationResponse) []SecretResponse { return v.Secrets }).(SecretResponseArrayOutput)
}

type ConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConfigurationResponse)(nil)).Elem()
}

func (o ConfigurationResponsePtrOutput) ToConfigurationResponsePtrOutput() ConfigurationResponsePtrOutput {
	return o
}

func (o ConfigurationResponsePtrOutput) ToConfigurationResponsePtrOutputWithContext(ctx context.Context) ConfigurationResponsePtrOutput {
	return o
}

func (o ConfigurationResponsePtrOutput) Elem() ConfigurationResponseOutput {
	return o.ApplyT(func(v *ConfigurationResponse) ConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ConfigurationResponse
		return ret
	}).(ConfigurationResponseOutput)
}

func (o ConfigurationResponsePtrOutput) ActiveRevisionsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActiveRevisionsMode
	}).(pulumi.StringPtrOutput)
}

func (o ConfigurationResponsePtrOutput) Ingress() IngressResponsePtrOutput {
	return o.ApplyT(func(v *ConfigurationResponse) *IngressResponse {
		if v == nil {
			return nil
		}
		return v.Ingress
	}).(IngressResponsePtrOutput)
}

func (o ConfigurationResponsePtrOutput) Registries() RegistryCredentialsResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationResponse) []RegistryCredentialsResponse {
		if v == nil {
			return nil
		}
		return v.Registries
	}).(RegistryCredentialsResponseArrayOutput)
}

func (o ConfigurationResponsePtrOutput) Secrets() SecretResponseArrayOutput {
	return o.ApplyT(func(v *ConfigurationResponse) []SecretResponse {
		if v == nil {
			return nil
		}
		return v.Secrets
	}).(SecretResponseArrayOutput)
}

type Container struct {
	Args      []string            `pulumi:"args"`
	Command   []string            `pulumi:"command"`
	Env       []EnvironmentVar    `pulumi:"env"`
	Image     *string             `pulumi:"image"`
	Name      *string             `pulumi:"name"`
	Resources *ContainerResources `pulumi:"resources"`
}





type ContainerInput interface {
	pulumi.Input

	ToContainerOutput() ContainerOutput
	ToContainerOutputWithContext(context.Context) ContainerOutput
}

type ContainerArgs struct {
	Args      pulumi.StringArrayInput    `pulumi:"args"`
	Command   pulumi.StringArrayInput    `pulumi:"command"`
	Env       EnvironmentVarArrayInput   `pulumi:"env"`
	Image     pulumi.StringPtrInput      `pulumi:"image"`
	Name      pulumi.StringPtrInput      `pulumi:"name"`
	Resources ContainerResourcesPtrInput `pulumi:"resources"`
}

func (ContainerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (i ContainerArgs) ToContainerOutput() ContainerOutput {
	return i.ToContainerOutputWithContext(context.Background())
}

func (i ContainerArgs) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerOutput)
}





type ContainerArrayInput interface {
	pulumi.Input

	ToContainerArrayOutput() ContainerArrayOutput
	ToContainerArrayOutputWithContext(context.Context) ContainerArrayOutput
}

type ContainerArray []ContainerInput

func (ContainerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil)).Elem()
}

func (i ContainerArray) ToContainerArrayOutput() ContainerArrayOutput {
	return i.ToContainerArrayOutputWithContext(context.Background())
}

func (i ContainerArray) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerArrayOutput)
}

type ContainerOutput struct{ *pulumi.OutputState }

func (ContainerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Container)(nil)).Elem()
}

func (o ContainerOutput) ToContainerOutput() ContainerOutput {
	return o
}

func (o ContainerOutput) ToContainerOutputWithContext(ctx context.Context) ContainerOutput {
	return o
}

func (o ContainerOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Container) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o ContainerOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v Container) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o ContainerOutput) Env() EnvironmentVarArrayOutput {
	return o.ApplyT(func(v Container) []EnvironmentVar { return v.Env }).(EnvironmentVarArrayOutput)
}

func (o ContainerOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Container) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContainerOutput) Resources() ContainerResourcesPtrOutput {
	return o.ApplyT(func(v Container) *ContainerResources { return v.Resources }).(ContainerResourcesPtrOutput)
}

type ContainerArrayOutput struct{ *pulumi.OutputState }

func (ContainerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Container)(nil)).Elem()
}

func (o ContainerArrayOutput) ToContainerArrayOutput() ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) ToContainerArrayOutputWithContext(ctx context.Context) ContainerArrayOutput {
	return o
}

func (o ContainerArrayOutput) Index(i pulumi.IntInput) ContainerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Container {
		return vs[0].([]Container)[vs[1].(int)]
	}).(ContainerOutput)
}

type ContainerResources struct {
	Cpu    *float64 `pulumi:"cpu"`
	Memory *string  `pulumi:"memory"`
}





type ContainerResourcesInput interface {
	pulumi.Input

	ToContainerResourcesOutput() ContainerResourcesOutput
	ToContainerResourcesOutputWithContext(context.Context) ContainerResourcesOutput
}

type ContainerResourcesArgs struct {
	Cpu    pulumi.Float64PtrInput `pulumi:"cpu"`
	Memory pulumi.StringPtrInput  `pulumi:"memory"`
}

func (ContainerResourcesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResources)(nil)).Elem()
}

func (i ContainerResourcesArgs) ToContainerResourcesOutput() ContainerResourcesOutput {
	return i.ToContainerResourcesOutputWithContext(context.Background())
}

func (i ContainerResourcesArgs) ToContainerResourcesOutputWithContext(ctx context.Context) ContainerResourcesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourcesOutput)
}

func (i ContainerResourcesArgs) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return i.ToContainerResourcesPtrOutputWithContext(context.Background())
}

func (i ContainerResourcesArgs) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourcesOutput).ToContainerResourcesPtrOutputWithContext(ctx)
}









type ContainerResourcesPtrInput interface {
	pulumi.Input

	ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput
	ToContainerResourcesPtrOutputWithContext(context.Context) ContainerResourcesPtrOutput
}

type containerResourcesPtrType ContainerResourcesArgs

func ContainerResourcesPtr(v *ContainerResourcesArgs) ContainerResourcesPtrInput {
	return (*containerResourcesPtrType)(v)
}

func (*containerResourcesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResources)(nil)).Elem()
}

func (i *containerResourcesPtrType) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return i.ToContainerResourcesPtrOutputWithContext(context.Background())
}

func (i *containerResourcesPtrType) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerResourcesPtrOutput)
}

type ContainerResourcesOutput struct{ *pulumi.OutputState }

func (ContainerResourcesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResources)(nil)).Elem()
}

func (o ContainerResourcesOutput) ToContainerResourcesOutput() ContainerResourcesOutput {
	return o
}

func (o ContainerResourcesOutput) ToContainerResourcesOutputWithContext(ctx context.Context) ContainerResourcesOutput {
	return o
}

func (o ContainerResourcesOutput) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return o.ToContainerResourcesPtrOutputWithContext(context.Background())
}

func (o ContainerResourcesOutput) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerResources) *ContainerResources {
		return &v
	}).(ContainerResourcesPtrOutput)
}

func (o ContainerResourcesOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResources) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResources) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type ContainerResourcesPtrOutput struct{ *pulumi.OutputState }

func (ContainerResourcesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResources)(nil)).Elem()
}

func (o ContainerResourcesPtrOutput) ToContainerResourcesPtrOutput() ContainerResourcesPtrOutput {
	return o
}

func (o ContainerResourcesPtrOutput) ToContainerResourcesPtrOutputWithContext(ctx context.Context) ContainerResourcesPtrOutput {
	return o
}

func (o ContainerResourcesPtrOutput) Elem() ContainerResourcesOutput {
	return o.ApplyT(func(v *ContainerResources) ContainerResources {
		if v != nil {
			return *v
		}
		var ret ContainerResources
		return ret
	}).(ContainerResourcesOutput)
}

func (o ContainerResourcesPtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResources) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesPtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerResources) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type ContainerResourcesResponse struct {
	Cpu    *float64 `pulumi:"cpu"`
	Memory *string  `pulumi:"memory"`
}

type ContainerResourcesResponseOutput struct{ *pulumi.OutputState }

func (ContainerResourcesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResourcesResponse)(nil)).Elem()
}

func (o ContainerResourcesResponseOutput) ToContainerResourcesResponseOutput() ContainerResourcesResponseOutput {
	return o
}

func (o ContainerResourcesResponseOutput) ToContainerResourcesResponseOutputWithContext(ctx context.Context) ContainerResourcesResponseOutput {
	return o
}

func (o ContainerResourcesResponseOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ContainerResourcesResponse) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesResponseOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResourcesResponse) *string { return v.Memory }).(pulumi.StringPtrOutput)
}

type ContainerResourcesResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerResourcesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerResourcesResponse)(nil)).Elem()
}

func (o ContainerResourcesResponsePtrOutput) ToContainerResourcesResponsePtrOutput() ContainerResourcesResponsePtrOutput {
	return o
}

func (o ContainerResourcesResponsePtrOutput) ToContainerResourcesResponsePtrOutputWithContext(ctx context.Context) ContainerResourcesResponsePtrOutput {
	return o
}

func (o ContainerResourcesResponsePtrOutput) Elem() ContainerResourcesResponseOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) ContainerResourcesResponse {
		if v != nil {
			return *v
		}
		var ret ContainerResourcesResponse
		return ret
	}).(ContainerResourcesResponseOutput)
}

func (o ContainerResourcesResponsePtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ContainerResourcesResponsePtrOutput) Memory() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerResourcesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Memory
	}).(pulumi.StringPtrOutput)
}

type ContainerResponse struct {
	Args      []string                    `pulumi:"args"`
	Command   []string                    `pulumi:"command"`
	Env       []EnvironmentVarResponse    `pulumi:"env"`
	Image     *string                     `pulumi:"image"`
	Name      *string                     `pulumi:"name"`
	Resources *ContainerResourcesResponse `pulumi:"resources"`
}

type ContainerResponseOutput struct{ *pulumi.OutputState }

func (ContainerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerResponse)(nil)).Elem()
}

func (o ContainerResponseOutput) ToContainerResponseOutput() ContainerResponseOutput {
	return o
}

func (o ContainerResponseOutput) ToContainerResponseOutputWithContext(ctx context.Context) ContainerResponseOutput {
	return o
}

func (o ContainerResponseOutput) Args() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []string { return v.Args }).(pulumi.StringArrayOutput)
}

func (o ContainerResponseOutput) Command() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []string { return v.Command }).(pulumi.StringArrayOutput)
}

func (o ContainerResponseOutput) Env() EnvironmentVarResponseArrayOutput {
	return o.ApplyT(func(v ContainerResponse) []EnvironmentVarResponse { return v.Env }).(EnvironmentVarResponseArrayOutput)
}

func (o ContainerResponseOutput) Image() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResponse) *string { return v.Image }).(pulumi.StringPtrOutput)
}

func (o ContainerResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContainerResponseOutput) Resources() ContainerResourcesResponsePtrOutput {
	return o.ApplyT(func(v ContainerResponse) *ContainerResourcesResponse { return v.Resources }).(ContainerResourcesResponsePtrOutput)
}

type ContainerResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerResponse)(nil)).Elem()
}

func (o ContainerResponseArrayOutput) ToContainerResponseArrayOutput() ContainerResponseArrayOutput {
	return o
}

func (o ContainerResponseArrayOutput) ToContainerResponseArrayOutputWithContext(ctx context.Context) ContainerResponseArrayOutput {
	return o
}

func (o ContainerResponseArrayOutput) Index(i pulumi.IntInput) ContainerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerResponse {
		return vs[0].([]ContainerResponse)[vs[1].(int)]
	}).(ContainerResponseOutput)
}

type CustomScaleRule struct {
	Auth     []ScaleRuleAuth   `pulumi:"auth"`
	Metadata map[string]string `pulumi:"metadata"`
	Type     *string           `pulumi:"type"`
}





type CustomScaleRuleInput interface {
	pulumi.Input

	ToCustomScaleRuleOutput() CustomScaleRuleOutput
	ToCustomScaleRuleOutputWithContext(context.Context) CustomScaleRuleOutput
}

type CustomScaleRuleArgs struct {
	Auth     ScaleRuleAuthArrayInput `pulumi:"auth"`
	Metadata pulumi.StringMapInput   `pulumi:"metadata"`
	Type     pulumi.StringPtrInput   `pulumi:"type"`
}

func (CustomScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomScaleRule)(nil)).Elem()
}

func (i CustomScaleRuleArgs) ToCustomScaleRuleOutput() CustomScaleRuleOutput {
	return i.ToCustomScaleRuleOutputWithContext(context.Background())
}

func (i CustomScaleRuleArgs) ToCustomScaleRuleOutputWithContext(ctx context.Context) CustomScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomScaleRuleOutput)
}

func (i CustomScaleRuleArgs) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return i.ToCustomScaleRulePtrOutputWithContext(context.Background())
}

func (i CustomScaleRuleArgs) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomScaleRuleOutput).ToCustomScaleRulePtrOutputWithContext(ctx)
}









type CustomScaleRulePtrInput interface {
	pulumi.Input

	ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput
	ToCustomScaleRulePtrOutputWithContext(context.Context) CustomScaleRulePtrOutput
}

type customScaleRulePtrType CustomScaleRuleArgs

func CustomScaleRulePtr(v *CustomScaleRuleArgs) CustomScaleRulePtrInput {
	return (*customScaleRulePtrType)(v)
}

func (*customScaleRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomScaleRule)(nil)).Elem()
}

func (i *customScaleRulePtrType) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return i.ToCustomScaleRulePtrOutputWithContext(context.Background())
}

func (i *customScaleRulePtrType) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomScaleRulePtrOutput)
}

type CustomScaleRuleOutput struct{ *pulumi.OutputState }

func (CustomScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomScaleRule)(nil)).Elem()
}

func (o CustomScaleRuleOutput) ToCustomScaleRuleOutput() CustomScaleRuleOutput {
	return o
}

func (o CustomScaleRuleOutput) ToCustomScaleRuleOutputWithContext(ctx context.Context) CustomScaleRuleOutput {
	return o
}

func (o CustomScaleRuleOutput) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return o.ToCustomScaleRulePtrOutputWithContext(context.Background())
}

func (o CustomScaleRuleOutput) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomScaleRule) *CustomScaleRule {
		return &v
	}).(CustomScaleRulePtrOutput)
}

func (o CustomScaleRuleOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v CustomScaleRule) []ScaleRuleAuth { return v.Auth }).(ScaleRuleAuthArrayOutput)
}

func (o CustomScaleRuleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v CustomScaleRule) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o CustomScaleRuleOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomScaleRule) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomScaleRulePtrOutput struct{ *pulumi.OutputState }

func (CustomScaleRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomScaleRule)(nil)).Elem()
}

func (o CustomScaleRulePtrOutput) ToCustomScaleRulePtrOutput() CustomScaleRulePtrOutput {
	return o
}

func (o CustomScaleRulePtrOutput) ToCustomScaleRulePtrOutputWithContext(ctx context.Context) CustomScaleRulePtrOutput {
	return o
}

func (o CustomScaleRulePtrOutput) Elem() CustomScaleRuleOutput {
	return o.ApplyT(func(v *CustomScaleRule) CustomScaleRule {
		if v != nil {
			return *v
		}
		var ret CustomScaleRule
		return ret
	}).(CustomScaleRuleOutput)
}

func (o CustomScaleRulePtrOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v *CustomScaleRule) []ScaleRuleAuth {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthArrayOutput)
}

func (o CustomScaleRulePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomScaleRule) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o CustomScaleRulePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomScaleRule) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type CustomScaleRuleResponse struct {
	Auth     []ScaleRuleAuthResponse `pulumi:"auth"`
	Metadata map[string]string       `pulumi:"metadata"`
	Type     *string                 `pulumi:"type"`
}

type CustomScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (CustomScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomScaleRuleResponse)(nil)).Elem()
}

func (o CustomScaleRuleResponseOutput) ToCustomScaleRuleResponseOutput() CustomScaleRuleResponseOutput {
	return o
}

func (o CustomScaleRuleResponseOutput) ToCustomScaleRuleResponseOutputWithContext(ctx context.Context) CustomScaleRuleResponseOutput {
	return o
}

func (o CustomScaleRuleResponseOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v CustomScaleRuleResponse) []ScaleRuleAuthResponse { return v.Auth }).(ScaleRuleAuthResponseArrayOutput)
}

func (o CustomScaleRuleResponseOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v CustomScaleRuleResponse) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

func (o CustomScaleRuleResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomScaleRuleResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type CustomScaleRuleResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomScaleRuleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomScaleRuleResponse)(nil)).Elem()
}

func (o CustomScaleRuleResponsePtrOutput) ToCustomScaleRuleResponsePtrOutput() CustomScaleRuleResponsePtrOutput {
	return o
}

func (o CustomScaleRuleResponsePtrOutput) ToCustomScaleRuleResponsePtrOutputWithContext(ctx context.Context) CustomScaleRuleResponsePtrOutput {
	return o
}

func (o CustomScaleRuleResponsePtrOutput) Elem() CustomScaleRuleResponseOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) CustomScaleRuleResponse {
		if v != nil {
			return *v
		}
		var ret CustomScaleRuleResponse
		return ret
	}).(CustomScaleRuleResponseOutput)
}

func (o CustomScaleRuleResponsePtrOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) []ScaleRuleAuthResponse {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthResponseArrayOutput)
}

func (o CustomScaleRuleResponsePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

func (o CustomScaleRuleResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomScaleRuleResponse) *string {
		if v == nil {
			return nil
		}
		return v.Type
	}).(pulumi.StringPtrOutput)
}

type Dapr struct {
	AppId      *string         `pulumi:"appId"`
	AppPort    *int            `pulumi:"appPort"`
	Components []DaprComponent `pulumi:"components"`
	Enabled    *bool           `pulumi:"enabled"`
}





type DaprInput interface {
	pulumi.Input

	ToDaprOutput() DaprOutput
	ToDaprOutputWithContext(context.Context) DaprOutput
}

type DaprArgs struct {
	AppId      pulumi.StringPtrInput   `pulumi:"appId"`
	AppPort    pulumi.IntPtrInput      `pulumi:"appPort"`
	Components DaprComponentArrayInput `pulumi:"components"`
	Enabled    pulumi.BoolPtrInput     `pulumi:"enabled"`
}

func (DaprArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Dapr)(nil)).Elem()
}

func (i DaprArgs) ToDaprOutput() DaprOutput {
	return i.ToDaprOutputWithContext(context.Background())
}

func (i DaprArgs) ToDaprOutputWithContext(ctx context.Context) DaprOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprOutput)
}

func (i DaprArgs) ToDaprPtrOutput() DaprPtrOutput {
	return i.ToDaprPtrOutputWithContext(context.Background())
}

func (i DaprArgs) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprOutput).ToDaprPtrOutputWithContext(ctx)
}









type DaprPtrInput interface {
	pulumi.Input

	ToDaprPtrOutput() DaprPtrOutput
	ToDaprPtrOutputWithContext(context.Context) DaprPtrOutput
}

type daprPtrType DaprArgs

func DaprPtr(v *DaprArgs) DaprPtrInput {
	return (*daprPtrType)(v)
}

func (*daprPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Dapr)(nil)).Elem()
}

func (i *daprPtrType) ToDaprPtrOutput() DaprPtrOutput {
	return i.ToDaprPtrOutputWithContext(context.Background())
}

func (i *daprPtrType) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprPtrOutput)
}

type DaprOutput struct{ *pulumi.OutputState }

func (DaprOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Dapr)(nil)).Elem()
}

func (o DaprOutput) ToDaprOutput() DaprOutput {
	return o
}

func (o DaprOutput) ToDaprOutputWithContext(ctx context.Context) DaprOutput {
	return o
}

func (o DaprOutput) ToDaprPtrOutput() DaprPtrOutput {
	return o.ToDaprPtrOutputWithContext(context.Background())
}

func (o DaprOutput) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Dapr) *Dapr {
		return &v
	}).(DaprPtrOutput)
}

func (o DaprOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Dapr) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o DaprOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Dapr) *int { return v.AppPort }).(pulumi.IntPtrOutput)
}

func (o DaprOutput) Components() DaprComponentArrayOutput {
	return o.ApplyT(func(v Dapr) []DaprComponent { return v.Components }).(DaprComponentArrayOutput)
}

func (o DaprOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Dapr) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DaprPtrOutput struct{ *pulumi.OutputState }

func (DaprPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Dapr)(nil)).Elem()
}

func (o DaprPtrOutput) ToDaprPtrOutput() DaprPtrOutput {
	return o
}

func (o DaprPtrOutput) ToDaprPtrOutputWithContext(ctx context.Context) DaprPtrOutput {
	return o
}

func (o DaprPtrOutput) Elem() DaprOutput {
	return o.ApplyT(func(v *Dapr) Dapr {
		if v != nil {
			return *v
		}
		var ret Dapr
		return ret
	}).(DaprOutput)
}

func (o DaprPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Dapr) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o DaprPtrOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Dapr) *int {
		if v == nil {
			return nil
		}
		return v.AppPort
	}).(pulumi.IntPtrOutput)
}

func (o DaprPtrOutput) Components() DaprComponentArrayOutput {
	return o.ApplyT(func(v *Dapr) []DaprComponent {
		if v == nil {
			return nil
		}
		return v.Components
	}).(DaprComponentArrayOutput)
}

func (o DaprPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Dapr) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type DaprComponent struct {
	Metadata []DaprMetadata `pulumi:"metadata"`
	Name     *string        `pulumi:"name"`
	Type     *string        `pulumi:"type"`
	Version  *string        `pulumi:"version"`
}





type DaprComponentInput interface {
	pulumi.Input

	ToDaprComponentOutput() DaprComponentOutput
	ToDaprComponentOutputWithContext(context.Context) DaprComponentOutput
}

type DaprComponentArgs struct {
	Metadata DaprMetadataArrayInput `pulumi:"metadata"`
	Name     pulumi.StringPtrInput  `pulumi:"name"`
	Type     pulumi.StringPtrInput  `pulumi:"type"`
	Version  pulumi.StringPtrInput  `pulumi:"version"`
}

func (DaprComponentArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprComponent)(nil)).Elem()
}

func (i DaprComponentArgs) ToDaprComponentOutput() DaprComponentOutput {
	return i.ToDaprComponentOutputWithContext(context.Background())
}

func (i DaprComponentArgs) ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprComponentOutput)
}





type DaprComponentArrayInput interface {
	pulumi.Input

	ToDaprComponentArrayOutput() DaprComponentArrayOutput
	ToDaprComponentArrayOutputWithContext(context.Context) DaprComponentArrayOutput
}

type DaprComponentArray []DaprComponentInput

func (DaprComponentArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprComponent)(nil)).Elem()
}

func (i DaprComponentArray) ToDaprComponentArrayOutput() DaprComponentArrayOutput {
	return i.ToDaprComponentArrayOutputWithContext(context.Background())
}

func (i DaprComponentArray) ToDaprComponentArrayOutputWithContext(ctx context.Context) DaprComponentArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprComponentArrayOutput)
}

type DaprComponentOutput struct{ *pulumi.OutputState }

func (DaprComponentOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprComponent)(nil)).Elem()
}

func (o DaprComponentOutput) ToDaprComponentOutput() DaprComponentOutput {
	return o
}

func (o DaprComponentOutput) ToDaprComponentOutputWithContext(ctx context.Context) DaprComponentOutput {
	return o
}

func (o DaprComponentOutput) Metadata() DaprMetadataArrayOutput {
	return o.ApplyT(func(v DaprComponent) []DaprMetadata { return v.Metadata }).(DaprMetadataArrayOutput)
}

func (o DaprComponentOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprComponent) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DaprComponentOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprComponent) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o DaprComponentOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprComponent) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DaprComponentArrayOutput struct{ *pulumi.OutputState }

func (DaprComponentArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprComponent)(nil)).Elem()
}

func (o DaprComponentArrayOutput) ToDaprComponentArrayOutput() DaprComponentArrayOutput {
	return o
}

func (o DaprComponentArrayOutput) ToDaprComponentArrayOutputWithContext(ctx context.Context) DaprComponentArrayOutput {
	return o
}

func (o DaprComponentArrayOutput) Index(i pulumi.IntInput) DaprComponentOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprComponent {
		return vs[0].([]DaprComponent)[vs[1].(int)]
	}).(DaprComponentOutput)
}

type DaprComponentResponse struct {
	Metadata []DaprMetadataResponse `pulumi:"metadata"`
	Name     *string                `pulumi:"name"`
	Type     *string                `pulumi:"type"`
	Version  *string                `pulumi:"version"`
}

type DaprComponentResponseOutput struct{ *pulumi.OutputState }

func (DaprComponentResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprComponentResponse)(nil)).Elem()
}

func (o DaprComponentResponseOutput) ToDaprComponentResponseOutput() DaprComponentResponseOutput {
	return o
}

func (o DaprComponentResponseOutput) ToDaprComponentResponseOutputWithContext(ctx context.Context) DaprComponentResponseOutput {
	return o
}

func (o DaprComponentResponseOutput) Metadata() DaprMetadataResponseArrayOutput {
	return o.ApplyT(func(v DaprComponentResponse) []DaprMetadataResponse { return v.Metadata }).(DaprMetadataResponseArrayOutput)
}

func (o DaprComponentResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprComponentResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DaprComponentResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprComponentResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o DaprComponentResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprComponentResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type DaprComponentResponseArrayOutput struct{ *pulumi.OutputState }

func (DaprComponentResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprComponentResponse)(nil)).Elem()
}

func (o DaprComponentResponseArrayOutput) ToDaprComponentResponseArrayOutput() DaprComponentResponseArrayOutput {
	return o
}

func (o DaprComponentResponseArrayOutput) ToDaprComponentResponseArrayOutputWithContext(ctx context.Context) DaprComponentResponseArrayOutput {
	return o
}

func (o DaprComponentResponseArrayOutput) Index(i pulumi.IntInput) DaprComponentResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprComponentResponse {
		return vs[0].([]DaprComponentResponse)[vs[1].(int)]
	}).(DaprComponentResponseOutput)
}

type DaprMetadata struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}





type DaprMetadataInput interface {
	pulumi.Input

	ToDaprMetadataOutput() DaprMetadataOutput
	ToDaprMetadataOutputWithContext(context.Context) DaprMetadataOutput
}

type DaprMetadataArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	SecretRef pulumi.StringPtrInput `pulumi:"secretRef"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (DaprMetadataArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprMetadata)(nil)).Elem()
}

func (i DaprMetadataArgs) ToDaprMetadataOutput() DaprMetadataOutput {
	return i.ToDaprMetadataOutputWithContext(context.Background())
}

func (i DaprMetadataArgs) ToDaprMetadataOutputWithContext(ctx context.Context) DaprMetadataOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprMetadataOutput)
}





type DaprMetadataArrayInput interface {
	pulumi.Input

	ToDaprMetadataArrayOutput() DaprMetadataArrayOutput
	ToDaprMetadataArrayOutputWithContext(context.Context) DaprMetadataArrayOutput
}

type DaprMetadataArray []DaprMetadataInput

func (DaprMetadataArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprMetadata)(nil)).Elem()
}

func (i DaprMetadataArray) ToDaprMetadataArrayOutput() DaprMetadataArrayOutput {
	return i.ToDaprMetadataArrayOutputWithContext(context.Background())
}

func (i DaprMetadataArray) ToDaprMetadataArrayOutputWithContext(ctx context.Context) DaprMetadataArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DaprMetadataArrayOutput)
}

type DaprMetadataOutput struct{ *pulumi.OutputState }

func (DaprMetadataOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprMetadata)(nil)).Elem()
}

func (o DaprMetadataOutput) ToDaprMetadataOutput() DaprMetadataOutput {
	return o
}

func (o DaprMetadataOutput) ToDaprMetadataOutputWithContext(ctx context.Context) DaprMetadataOutput {
	return o
}

func (o DaprMetadataOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadata) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadata) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadata) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DaprMetadataArrayOutput struct{ *pulumi.OutputState }

func (DaprMetadataArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprMetadata)(nil)).Elem()
}

func (o DaprMetadataArrayOutput) ToDaprMetadataArrayOutput() DaprMetadataArrayOutput {
	return o
}

func (o DaprMetadataArrayOutput) ToDaprMetadataArrayOutputWithContext(ctx context.Context) DaprMetadataArrayOutput {
	return o
}

func (o DaprMetadataArrayOutput) Index(i pulumi.IntInput) DaprMetadataOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprMetadata {
		return vs[0].([]DaprMetadata)[vs[1].(int)]
	}).(DaprMetadataOutput)
}

type DaprMetadataResponse struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}

type DaprMetadataResponseOutput struct{ *pulumi.OutputState }

func (DaprMetadataResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprMetadataResponse)(nil)).Elem()
}

func (o DaprMetadataResponseOutput) ToDaprMetadataResponseOutput() DaprMetadataResponseOutput {
	return o
}

func (o DaprMetadataResponseOutput) ToDaprMetadataResponseOutputWithContext(ctx context.Context) DaprMetadataResponseOutput {
	return o
}

func (o DaprMetadataResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadataResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataResponseOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadataResponse) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o DaprMetadataResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprMetadataResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type DaprMetadataResponseArrayOutput struct{ *pulumi.OutputState }

func (DaprMetadataResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DaprMetadataResponse)(nil)).Elem()
}

func (o DaprMetadataResponseArrayOutput) ToDaprMetadataResponseArrayOutput() DaprMetadataResponseArrayOutput {
	return o
}

func (o DaprMetadataResponseArrayOutput) ToDaprMetadataResponseArrayOutputWithContext(ctx context.Context) DaprMetadataResponseArrayOutput {
	return o
}

func (o DaprMetadataResponseArrayOutput) Index(i pulumi.IntInput) DaprMetadataResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DaprMetadataResponse {
		return vs[0].([]DaprMetadataResponse)[vs[1].(int)]
	}).(DaprMetadataResponseOutput)
}

type DaprResponse struct {
	AppId      *string                 `pulumi:"appId"`
	AppPort    *int                    `pulumi:"appPort"`
	Components []DaprComponentResponse `pulumi:"components"`
	Enabled    *bool                   `pulumi:"enabled"`
}

type DaprResponseOutput struct{ *pulumi.OutputState }

func (DaprResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DaprResponse)(nil)).Elem()
}

func (o DaprResponseOutput) ToDaprResponseOutput() DaprResponseOutput {
	return o
}

func (o DaprResponseOutput) ToDaprResponseOutputWithContext(ctx context.Context) DaprResponseOutput {
	return o
}

func (o DaprResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DaprResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o DaprResponseOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v DaprResponse) *int { return v.AppPort }).(pulumi.IntPtrOutput)
}

func (o DaprResponseOutput) Components() DaprComponentResponseArrayOutput {
	return o.ApplyT(func(v DaprResponse) []DaprComponentResponse { return v.Components }).(DaprComponentResponseArrayOutput)
}

func (o DaprResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DaprResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type DaprResponsePtrOutput struct{ *pulumi.OutputState }

func (DaprResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DaprResponse)(nil)).Elem()
}

func (o DaprResponsePtrOutput) ToDaprResponsePtrOutput() DaprResponsePtrOutput {
	return o
}

func (o DaprResponsePtrOutput) ToDaprResponsePtrOutputWithContext(ctx context.Context) DaprResponsePtrOutput {
	return o
}

func (o DaprResponsePtrOutput) Elem() DaprResponseOutput {
	return o.ApplyT(func(v *DaprResponse) DaprResponse {
		if v != nil {
			return *v
		}
		var ret DaprResponse
		return ret
	}).(DaprResponseOutput)
}

func (o DaprResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o DaprResponsePtrOutput) AppPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *int {
		if v == nil {
			return nil
		}
		return v.AppPort
	}).(pulumi.IntPtrOutput)
}

func (o DaprResponsePtrOutput) Components() DaprComponentResponseArrayOutput {
	return o.ApplyT(func(v *DaprResponse) []DaprComponentResponse {
		if v == nil {
			return nil
		}
		return v.Components
	}).(DaprComponentResponseArrayOutput)
}

func (o DaprResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DaprResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EnvironmentVar struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}





type EnvironmentVarInput interface {
	pulumi.Input

	ToEnvironmentVarOutput() EnvironmentVarOutput
	ToEnvironmentVarOutputWithContext(context.Context) EnvironmentVarOutput
}

type EnvironmentVarArgs struct {
	Name      pulumi.StringPtrInput `pulumi:"name"`
	SecretRef pulumi.StringPtrInput `pulumi:"secretRef"`
	Value     pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentVarArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVar)(nil)).Elem()
}

func (i EnvironmentVarArgs) ToEnvironmentVarOutput() EnvironmentVarOutput {
	return i.ToEnvironmentVarOutputWithContext(context.Background())
}

func (i EnvironmentVarArgs) ToEnvironmentVarOutputWithContext(ctx context.Context) EnvironmentVarOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVarOutput)
}





type EnvironmentVarArrayInput interface {
	pulumi.Input

	ToEnvironmentVarArrayOutput() EnvironmentVarArrayOutput
	ToEnvironmentVarArrayOutputWithContext(context.Context) EnvironmentVarArrayOutput
}

type EnvironmentVarArray []EnvironmentVarInput

func (EnvironmentVarArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVar)(nil)).Elem()
}

func (i EnvironmentVarArray) ToEnvironmentVarArrayOutput() EnvironmentVarArrayOutput {
	return i.ToEnvironmentVarArrayOutputWithContext(context.Background())
}

func (i EnvironmentVarArray) ToEnvironmentVarArrayOutputWithContext(ctx context.Context) EnvironmentVarArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVarArrayOutput)
}

type EnvironmentVarOutput struct{ *pulumi.OutputState }

func (EnvironmentVarOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVar)(nil)).Elem()
}

func (o EnvironmentVarOutput) ToEnvironmentVarOutput() EnvironmentVarOutput {
	return o
}

func (o EnvironmentVarOutput) ToEnvironmentVarOutputWithContext(ctx context.Context) EnvironmentVarOutput {
	return o
}

func (o EnvironmentVarOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVar) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVar) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVar) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVarArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVarArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVar)(nil)).Elem()
}

func (o EnvironmentVarArrayOutput) ToEnvironmentVarArrayOutput() EnvironmentVarArrayOutput {
	return o
}

func (o EnvironmentVarArrayOutput) ToEnvironmentVarArrayOutputWithContext(ctx context.Context) EnvironmentVarArrayOutput {
	return o
}

func (o EnvironmentVarArrayOutput) Index(i pulumi.IntInput) EnvironmentVarOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVar {
		return vs[0].([]EnvironmentVar)[vs[1].(int)]
	}).(EnvironmentVarOutput)
}

type EnvironmentVarResponse struct {
	Name      *string `pulumi:"name"`
	SecretRef *string `pulumi:"secretRef"`
	Value     *string `pulumi:"value"`
}

type EnvironmentVarResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentVarResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVarResponse)(nil)).Elem()
}

func (o EnvironmentVarResponseOutput) ToEnvironmentVarResponseOutput() EnvironmentVarResponseOutput {
	return o
}

func (o EnvironmentVarResponseOutput) ToEnvironmentVarResponseOutputWithContext(ctx context.Context) EnvironmentVarResponseOutput {
	return o
}

func (o EnvironmentVarResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVarResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarResponseOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVarResponse) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVarResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVarResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVarResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVarResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVarResponse)(nil)).Elem()
}

func (o EnvironmentVarResponseArrayOutput) ToEnvironmentVarResponseArrayOutput() EnvironmentVarResponseArrayOutput {
	return o
}

func (o EnvironmentVarResponseArrayOutput) ToEnvironmentVarResponseArrayOutputWithContext(ctx context.Context) EnvironmentVarResponseArrayOutput {
	return o
}

func (o EnvironmentVarResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentVarResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVarResponse {
		return vs[0].([]EnvironmentVarResponse)[vs[1].(int)]
	}).(EnvironmentVarResponseOutput)
}

type ExtendedLocation struct {
	Name *string `pulumi:"name"`
}





type ExtendedLocationInput interface {
	pulumi.Input

	ToExtendedLocationOutput() ExtendedLocationOutput
	ToExtendedLocationOutputWithContext(context.Context) ExtendedLocationOutput
}

type ExtendedLocationArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (ExtendedLocationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (i ExtendedLocationArgs) ToExtendedLocationOutput() ExtendedLocationOutput {
	return i.ToExtendedLocationOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput)
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i ExtendedLocationArgs) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationOutput).ToExtendedLocationPtrOutputWithContext(ctx)
}









type ExtendedLocationPtrInput interface {
	pulumi.Input

	ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput
	ToExtendedLocationPtrOutputWithContext(context.Context) ExtendedLocationPtrOutput
}

type extendedLocationPtrType ExtendedLocationArgs

func ExtendedLocationPtr(v *ExtendedLocationArgs) ExtendedLocationPtrInput {
	return (*extendedLocationPtrType)(v)
}

func (*extendedLocationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return i.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (i *extendedLocationPtrType) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExtendedLocationPtrOutput)
}

type ExtendedLocationOutput struct{ *pulumi.OutputState }

func (ExtendedLocationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationOutput) ToExtendedLocationOutput() ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationOutputWithContext(ctx context.Context) ExtendedLocationOutput {
	return o
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o.ToExtendedLocationPtrOutputWithContext(context.Background())
}

func (o ExtendedLocationOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExtendedLocation) *ExtendedLocation {
		return &v
	}).(ExtendedLocationPtrOutput)
}

func (o ExtendedLocationOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocation) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ExtendedLocationPtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocation)(nil)).Elem()
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutput() ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) ToExtendedLocationPtrOutputWithContext(ctx context.Context) ExtendedLocationPtrOutput {
	return o
}

func (o ExtendedLocationPtrOutput) Elem() ExtendedLocationOutput {
	return o.ApplyT(func(v *ExtendedLocation) ExtendedLocation {
		if v != nil {
			return *v
		}
		var ret ExtendedLocation
		return ret
	}).(ExtendedLocationOutput)
}

func (o ExtendedLocationPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocation) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

type ExtendedLocationResponse struct {
	Name *string `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type ExtendedLocationResponseOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutput() ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) ToExtendedLocationResponseOutputWithContext(ctx context.Context) ExtendedLocationResponseOutput {
	return o
}

func (o ExtendedLocationResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ExtendedLocationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ExtendedLocationResponsePtrOutput struct{ *pulumi.OutputState }

func (ExtendedLocationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExtendedLocationResponse)(nil)).Elem()
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutput() ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) ToExtendedLocationResponsePtrOutputWithContext(ctx context.Context) ExtendedLocationResponsePtrOutput {
	return o
}

func (o ExtendedLocationResponsePtrOutput) Elem() ExtendedLocationResponseOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) ExtendedLocationResponse {
		if v != nil {
			return *v
		}
		var ret ExtendedLocationResponse
		return ret
	}).(ExtendedLocationResponseOutput)
}

func (o ExtendedLocationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ExtendedLocationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ExtendedLocationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type FrontEndConfiguration struct {
	Kind *FrontEndServiceType `pulumi:"kind"`
}





type FrontEndConfigurationInput interface {
	pulumi.Input

	ToFrontEndConfigurationOutput() FrontEndConfigurationOutput
	ToFrontEndConfigurationOutputWithContext(context.Context) FrontEndConfigurationOutput
}

type FrontEndConfigurationArgs struct {
	Kind FrontEndServiceTypePtrInput `pulumi:"kind"`
}

func (FrontEndConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndConfiguration)(nil)).Elem()
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationOutput() FrontEndConfigurationOutput {
	return i.ToFrontEndConfigurationOutputWithContext(context.Background())
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationOutputWithContext(ctx context.Context) FrontEndConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontEndConfigurationOutput)
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return i.ToFrontEndConfigurationPtrOutputWithContext(context.Background())
}

func (i FrontEndConfigurationArgs) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontEndConfigurationOutput).ToFrontEndConfigurationPtrOutputWithContext(ctx)
}









type FrontEndConfigurationPtrInput interface {
	pulumi.Input

	ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput
	ToFrontEndConfigurationPtrOutputWithContext(context.Context) FrontEndConfigurationPtrOutput
}

type frontEndConfigurationPtrType FrontEndConfigurationArgs

func FrontEndConfigurationPtr(v *FrontEndConfigurationArgs) FrontEndConfigurationPtrInput {
	return (*frontEndConfigurationPtrType)(v)
}

func (*frontEndConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndConfiguration)(nil)).Elem()
}

func (i *frontEndConfigurationPtrType) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return i.ToFrontEndConfigurationPtrOutputWithContext(context.Background())
}

func (i *frontEndConfigurationPtrType) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FrontEndConfigurationPtrOutput)
}

type FrontEndConfigurationOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndConfiguration)(nil)).Elem()
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationOutput() FrontEndConfigurationOutput {
	return o
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationOutputWithContext(ctx context.Context) FrontEndConfigurationOutput {
	return o
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return o.ToFrontEndConfigurationPtrOutputWithContext(context.Background())
}

func (o FrontEndConfigurationOutput) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FrontEndConfiguration) *FrontEndConfiguration {
		return &v
	}).(FrontEndConfigurationPtrOutput)
}

func (o FrontEndConfigurationOutput) Kind() FrontEndServiceTypePtrOutput {
	return o.ApplyT(func(v FrontEndConfiguration) *FrontEndServiceType { return v.Kind }).(FrontEndServiceTypePtrOutput)
}

type FrontEndConfigurationPtrOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndConfiguration)(nil)).Elem()
}

func (o FrontEndConfigurationPtrOutput) ToFrontEndConfigurationPtrOutput() FrontEndConfigurationPtrOutput {
	return o
}

func (o FrontEndConfigurationPtrOutput) ToFrontEndConfigurationPtrOutputWithContext(ctx context.Context) FrontEndConfigurationPtrOutput {
	return o
}

func (o FrontEndConfigurationPtrOutput) Elem() FrontEndConfigurationOutput {
	return o.ApplyT(func(v *FrontEndConfiguration) FrontEndConfiguration {
		if v != nil {
			return *v
		}
		var ret FrontEndConfiguration
		return ret
	}).(FrontEndConfigurationOutput)
}

func (o FrontEndConfigurationPtrOutput) Kind() FrontEndServiceTypePtrOutput {
	return o.ApplyT(func(v *FrontEndConfiguration) *FrontEndServiceType {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(FrontEndServiceTypePtrOutput)
}

type FrontEndConfigurationResponse struct {
	Kind *string `pulumi:"kind"`
}

type FrontEndConfigurationResponseOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FrontEndConfigurationResponse)(nil)).Elem()
}

func (o FrontEndConfigurationResponseOutput) ToFrontEndConfigurationResponseOutput() FrontEndConfigurationResponseOutput {
	return o
}

func (o FrontEndConfigurationResponseOutput) ToFrontEndConfigurationResponseOutputWithContext(ctx context.Context) FrontEndConfigurationResponseOutput {
	return o
}

func (o FrontEndConfigurationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FrontEndConfigurationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type FrontEndConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (FrontEndConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FrontEndConfigurationResponse)(nil)).Elem()
}

func (o FrontEndConfigurationResponsePtrOutput) ToFrontEndConfigurationResponsePtrOutput() FrontEndConfigurationResponsePtrOutput {
	return o
}

func (o FrontEndConfigurationResponsePtrOutput) ToFrontEndConfigurationResponsePtrOutputWithContext(ctx context.Context) FrontEndConfigurationResponsePtrOutput {
	return o
}

func (o FrontEndConfigurationResponsePtrOutput) Elem() FrontEndConfigurationResponseOutput {
	return o.ApplyT(func(v *FrontEndConfigurationResponse) FrontEndConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret FrontEndConfigurationResponse
		return ret
	}).(FrontEndConfigurationResponseOutput)
}

func (o FrontEndConfigurationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FrontEndConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

type HttpScaleRule struct {
	Auth     []ScaleRuleAuth   `pulumi:"auth"`
	Metadata map[string]string `pulumi:"metadata"`
}





type HttpScaleRuleInput interface {
	pulumi.Input

	ToHttpScaleRuleOutput() HttpScaleRuleOutput
	ToHttpScaleRuleOutputWithContext(context.Context) HttpScaleRuleOutput
}

type HttpScaleRuleArgs struct {
	Auth     ScaleRuleAuthArrayInput `pulumi:"auth"`
	Metadata pulumi.StringMapInput   `pulumi:"metadata"`
}

func (HttpScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpScaleRule)(nil)).Elem()
}

func (i HttpScaleRuleArgs) ToHttpScaleRuleOutput() HttpScaleRuleOutput {
	return i.ToHttpScaleRuleOutputWithContext(context.Background())
}

func (i HttpScaleRuleArgs) ToHttpScaleRuleOutputWithContext(ctx context.Context) HttpScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpScaleRuleOutput)
}

func (i HttpScaleRuleArgs) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return i.ToHttpScaleRulePtrOutputWithContext(context.Background())
}

func (i HttpScaleRuleArgs) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpScaleRuleOutput).ToHttpScaleRulePtrOutputWithContext(ctx)
}









type HttpScaleRulePtrInput interface {
	pulumi.Input

	ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput
	ToHttpScaleRulePtrOutputWithContext(context.Context) HttpScaleRulePtrOutput
}

type httpScaleRulePtrType HttpScaleRuleArgs

func HttpScaleRulePtr(v *HttpScaleRuleArgs) HttpScaleRulePtrInput {
	return (*httpScaleRulePtrType)(v)
}

func (*httpScaleRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpScaleRule)(nil)).Elem()
}

func (i *httpScaleRulePtrType) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return i.ToHttpScaleRulePtrOutputWithContext(context.Background())
}

func (i *httpScaleRulePtrType) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpScaleRulePtrOutput)
}

type HttpScaleRuleOutput struct{ *pulumi.OutputState }

func (HttpScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpScaleRule)(nil)).Elem()
}

func (o HttpScaleRuleOutput) ToHttpScaleRuleOutput() HttpScaleRuleOutput {
	return o
}

func (o HttpScaleRuleOutput) ToHttpScaleRuleOutputWithContext(ctx context.Context) HttpScaleRuleOutput {
	return o
}

func (o HttpScaleRuleOutput) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return o.ToHttpScaleRulePtrOutputWithContext(context.Background())
}

func (o HttpScaleRuleOutput) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpScaleRule) *HttpScaleRule {
		return &v
	}).(HttpScaleRulePtrOutput)
}

func (o HttpScaleRuleOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v HttpScaleRule) []ScaleRuleAuth { return v.Auth }).(ScaleRuleAuthArrayOutput)
}

func (o HttpScaleRuleOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v HttpScaleRule) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

type HttpScaleRulePtrOutput struct{ *pulumi.OutputState }

func (HttpScaleRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpScaleRule)(nil)).Elem()
}

func (o HttpScaleRulePtrOutput) ToHttpScaleRulePtrOutput() HttpScaleRulePtrOutput {
	return o
}

func (o HttpScaleRulePtrOutput) ToHttpScaleRulePtrOutputWithContext(ctx context.Context) HttpScaleRulePtrOutput {
	return o
}

func (o HttpScaleRulePtrOutput) Elem() HttpScaleRuleOutput {
	return o.ApplyT(func(v *HttpScaleRule) HttpScaleRule {
		if v != nil {
			return *v
		}
		var ret HttpScaleRule
		return ret
	}).(HttpScaleRuleOutput)
}

func (o HttpScaleRulePtrOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v *HttpScaleRule) []ScaleRuleAuth {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthArrayOutput)
}

func (o HttpScaleRulePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpScaleRule) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

type HttpScaleRuleResponse struct {
	Auth     []ScaleRuleAuthResponse `pulumi:"auth"`
	Metadata map[string]string       `pulumi:"metadata"`
}

type HttpScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (HttpScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpScaleRuleResponse)(nil)).Elem()
}

func (o HttpScaleRuleResponseOutput) ToHttpScaleRuleResponseOutput() HttpScaleRuleResponseOutput {
	return o
}

func (o HttpScaleRuleResponseOutput) ToHttpScaleRuleResponseOutputWithContext(ctx context.Context) HttpScaleRuleResponseOutput {
	return o
}

func (o HttpScaleRuleResponseOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v HttpScaleRuleResponse) []ScaleRuleAuthResponse { return v.Auth }).(ScaleRuleAuthResponseArrayOutput)
}

func (o HttpScaleRuleResponseOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v HttpScaleRuleResponse) map[string]string { return v.Metadata }).(pulumi.StringMapOutput)
}

type HttpScaleRuleResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpScaleRuleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpScaleRuleResponse)(nil)).Elem()
}

func (o HttpScaleRuleResponsePtrOutput) ToHttpScaleRuleResponsePtrOutput() HttpScaleRuleResponsePtrOutput {
	return o
}

func (o HttpScaleRuleResponsePtrOutput) ToHttpScaleRuleResponsePtrOutputWithContext(ctx context.Context) HttpScaleRuleResponsePtrOutput {
	return o
}

func (o HttpScaleRuleResponsePtrOutput) Elem() HttpScaleRuleResponseOutput {
	return o.ApplyT(func(v *HttpScaleRuleResponse) HttpScaleRuleResponse {
		if v != nil {
			return *v
		}
		var ret HttpScaleRuleResponse
		return ret
	}).(HttpScaleRuleResponseOutput)
}

func (o HttpScaleRuleResponsePtrOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v *HttpScaleRuleResponse) []ScaleRuleAuthResponse {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthResponseArrayOutput)
}

func (o HttpScaleRuleResponsePtrOutput) Metadata() pulumi.StringMapOutput {
	return o.ApplyT(func(v *HttpScaleRuleResponse) map[string]string {
		if v == nil {
			return nil
		}
		return v.Metadata
	}).(pulumi.StringMapOutput)
}

type Ingress struct {
	AllowInsecure *bool           `pulumi:"allowInsecure"`
	External      *bool           `pulumi:"external"`
	TargetPort    *int            `pulumi:"targetPort"`
	Traffic       []TrafficWeight `pulumi:"traffic"`
	Transport     *string         `pulumi:"transport"`
}


func (val *Ingress) Defaults() *Ingress {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.External) {
		external_ := false
		tmp.External = &external_
	}
	return &tmp
}





type IngressInput interface {
	pulumi.Input

	ToIngressOutput() IngressOutput
	ToIngressOutputWithContext(context.Context) IngressOutput
}

type IngressArgs struct {
	AllowInsecure pulumi.BoolPtrInput     `pulumi:"allowInsecure"`
	External      pulumi.BoolPtrInput     `pulumi:"external"`
	TargetPort    pulumi.IntPtrInput      `pulumi:"targetPort"`
	Traffic       TrafficWeightArrayInput `pulumi:"traffic"`
	Transport     pulumi.StringPtrInput   `pulumi:"transport"`
}

func (IngressArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Ingress)(nil)).Elem()
}

func (i IngressArgs) ToIngressOutput() IngressOutput {
	return i.ToIngressOutputWithContext(context.Background())
}

func (i IngressArgs) ToIngressOutputWithContext(ctx context.Context) IngressOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressOutput)
}

func (i IngressArgs) ToIngressPtrOutput() IngressPtrOutput {
	return i.ToIngressPtrOutputWithContext(context.Background())
}

func (i IngressArgs) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressOutput).ToIngressPtrOutputWithContext(ctx)
}









type IngressPtrInput interface {
	pulumi.Input

	ToIngressPtrOutput() IngressPtrOutput
	ToIngressPtrOutputWithContext(context.Context) IngressPtrOutput
}

type ingressPtrType IngressArgs

func IngressPtr(v *IngressArgs) IngressPtrInput {
	return (*ingressPtrType)(v)
}

func (*ingressPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingress)(nil)).Elem()
}

func (i *ingressPtrType) ToIngressPtrOutput() IngressPtrOutput {
	return i.ToIngressPtrOutputWithContext(context.Background())
}

func (i *ingressPtrType) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressPtrOutput)
}

type IngressOutput struct{ *pulumi.OutputState }

func (IngressOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Ingress)(nil)).Elem()
}

func (o IngressOutput) ToIngressOutput() IngressOutput {
	return o
}

func (o IngressOutput) ToIngressOutputWithContext(ctx context.Context) IngressOutput {
	return o
}

func (o IngressOutput) ToIngressPtrOutput() IngressPtrOutput {
	return o.ToIngressPtrOutputWithContext(context.Background())
}

func (o IngressOutput) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Ingress) *Ingress {
		return &v
	}).(IngressPtrOutput)
}

func (o IngressOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ingress) *bool { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

func (o IngressOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v Ingress) *bool { return v.External }).(pulumi.BoolPtrOutput)
}

func (o IngressOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Ingress) *int { return v.TargetPort }).(pulumi.IntPtrOutput)
}

func (o IngressOutput) Traffic() TrafficWeightArrayOutput {
	return o.ApplyT(func(v Ingress) []TrafficWeight { return v.Traffic }).(TrafficWeightArrayOutput)
}

func (o IngressOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Ingress) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

type IngressPtrOutput struct{ *pulumi.OutputState }

func (IngressPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Ingress)(nil)).Elem()
}

func (o IngressPtrOutput) ToIngressPtrOutput() IngressPtrOutput {
	return o
}

func (o IngressPtrOutput) ToIngressPtrOutputWithContext(ctx context.Context) IngressPtrOutput {
	return o
}

func (o IngressPtrOutput) Elem() IngressOutput {
	return o.ApplyT(func(v *Ingress) Ingress {
		if v != nil {
			return *v
		}
		var ret Ingress
		return ret
	}).(IngressOutput)
}

func (o IngressPtrOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ingress) *bool {
		if v == nil {
			return nil
		}
		return v.AllowInsecure
	}).(pulumi.BoolPtrOutput)
}

func (o IngressPtrOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *Ingress) *bool {
		if v == nil {
			return nil
		}
		return v.External
	}).(pulumi.BoolPtrOutput)
}

func (o IngressPtrOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Ingress) *int {
		if v == nil {
			return nil
		}
		return v.TargetPort
	}).(pulumi.IntPtrOutput)
}

func (o IngressPtrOutput) Traffic() TrafficWeightArrayOutput {
	return o.ApplyT(func(v *Ingress) []TrafficWeight {
		if v == nil {
			return nil
		}
		return v.Traffic
	}).(TrafficWeightArrayOutput)
}

func (o IngressPtrOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Ingress) *string {
		if v == nil {
			return nil
		}
		return v.Transport
	}).(pulumi.StringPtrOutput)
}

type IngressResponse struct {
	AllowInsecure *bool                   `pulumi:"allowInsecure"`
	External      *bool                   `pulumi:"external"`
	Fqdn          string                  `pulumi:"fqdn"`
	TargetPort    *int                    `pulumi:"targetPort"`
	Traffic       []TrafficWeightResponse `pulumi:"traffic"`
	Transport     *string                 `pulumi:"transport"`
}


func (val *IngressResponse) Defaults() *IngressResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.External) {
		external_ := false
		tmp.External = &external_
	}
	return &tmp
}

type IngressResponseOutput struct{ *pulumi.OutputState }

func (IngressResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressResponse)(nil)).Elem()
}

func (o IngressResponseOutput) ToIngressResponseOutput() IngressResponseOutput {
	return o
}

func (o IngressResponseOutput) ToIngressResponseOutputWithContext(ctx context.Context) IngressResponseOutput {
	return o
}

func (o IngressResponseOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngressResponse) *bool { return v.AllowInsecure }).(pulumi.BoolPtrOutput)
}

func (o IngressResponseOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v IngressResponse) *bool { return v.External }).(pulumi.BoolPtrOutput)
}

func (o IngressResponseOutput) Fqdn() pulumi.StringOutput {
	return o.ApplyT(func(v IngressResponse) string { return v.Fqdn }).(pulumi.StringOutput)
}

func (o IngressResponseOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IngressResponse) *int { return v.TargetPort }).(pulumi.IntPtrOutput)
}

func (o IngressResponseOutput) Traffic() TrafficWeightResponseArrayOutput {
	return o.ApplyT(func(v IngressResponse) []TrafficWeightResponse { return v.Traffic }).(TrafficWeightResponseArrayOutput)
}

func (o IngressResponseOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressResponse) *string { return v.Transport }).(pulumi.StringPtrOutput)
}

type IngressResponsePtrOutput struct{ *pulumi.OutputState }

func (IngressResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressResponse)(nil)).Elem()
}

func (o IngressResponsePtrOutput) ToIngressResponsePtrOutput() IngressResponsePtrOutput {
	return o
}

func (o IngressResponsePtrOutput) ToIngressResponsePtrOutputWithContext(ctx context.Context) IngressResponsePtrOutput {
	return o
}

func (o IngressResponsePtrOutput) Elem() IngressResponseOutput {
	return o.ApplyT(func(v *IngressResponse) IngressResponse {
		if v != nil {
			return *v
		}
		var ret IngressResponse
		return ret
	}).(IngressResponseOutput)
}

func (o IngressResponsePtrOutput) AllowInsecure() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowInsecure
	}).(pulumi.BoolPtrOutput)
}

func (o IngressResponsePtrOutput) External() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *bool {
		if v == nil {
			return nil
		}
		return v.External
	}).(pulumi.BoolPtrOutput)
}

func (o IngressResponsePtrOutput) Fqdn() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Fqdn
	}).(pulumi.StringPtrOutput)
}

func (o IngressResponsePtrOutput) TargetPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *int {
		if v == nil {
			return nil
		}
		return v.TargetPort
	}).(pulumi.IntPtrOutput)
}

func (o IngressResponsePtrOutput) Traffic() TrafficWeightResponseArrayOutput {
	return o.ApplyT(func(v *IngressResponse) []TrafficWeightResponse {
		if v == nil {
			return nil
		}
		return v.Traffic
	}).(TrafficWeightResponseArrayOutput)
}

func (o IngressResponsePtrOutput) Transport() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressResponse) *string {
		if v == nil {
			return nil
		}
		return v.Transport
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfiguration struct {
	CustomerId *string `pulumi:"customerId"`
	SharedKey  *string `pulumi:"sharedKey"`
}





type LogAnalyticsConfigurationInput interface {
	pulumi.Input

	ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput
	ToLogAnalyticsConfigurationOutputWithContext(context.Context) LogAnalyticsConfigurationOutput
}

type LogAnalyticsConfigurationArgs struct {
	CustomerId pulumi.StringPtrInput `pulumi:"customerId"`
	SharedKey  pulumi.StringPtrInput `pulumi:"sharedKey"`
}

func (LogAnalyticsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfiguration)(nil)).Elem()
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput {
	return i.ToLogAnalyticsConfigurationOutputWithContext(context.Background())
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationOutputWithContext(ctx context.Context) LogAnalyticsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationOutput)
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return i.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i LogAnalyticsConfigurationArgs) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationOutput).ToLogAnalyticsConfigurationPtrOutputWithContext(ctx)
}









type LogAnalyticsConfigurationPtrInput interface {
	pulumi.Input

	ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput
	ToLogAnalyticsConfigurationPtrOutputWithContext(context.Context) LogAnalyticsConfigurationPtrOutput
}

type logAnalyticsConfigurationPtrType LogAnalyticsConfigurationArgs

func LogAnalyticsConfigurationPtr(v *LogAnalyticsConfigurationArgs) LogAnalyticsConfigurationPtrInput {
	return (*logAnalyticsConfigurationPtrType)(v)
}

func (*logAnalyticsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfiguration)(nil)).Elem()
}

func (i *logAnalyticsConfigurationPtrType) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return i.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (i *logAnalyticsConfigurationPtrType) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LogAnalyticsConfigurationPtrOutput)
}

type LogAnalyticsConfigurationOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfiguration)(nil)).Elem()
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationOutput() LogAnalyticsConfigurationOutput {
	return o
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationOutputWithContext(ctx context.Context) LogAnalyticsConfigurationOutput {
	return o
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return o.ToLogAnalyticsConfigurationPtrOutputWithContext(context.Background())
}

func (o LogAnalyticsConfigurationOutput) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LogAnalyticsConfiguration) *LogAnalyticsConfiguration {
		return &v
	}).(LogAnalyticsConfigurationPtrOutput)
}

func (o LogAnalyticsConfigurationOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfiguration) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsConfigurationOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfiguration) *string { return v.SharedKey }).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfiguration)(nil)).Elem()
}

func (o LogAnalyticsConfigurationPtrOutput) ToLogAnalyticsConfigurationPtrOutput() LogAnalyticsConfigurationPtrOutput {
	return o
}

func (o LogAnalyticsConfigurationPtrOutput) ToLogAnalyticsConfigurationPtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationPtrOutput {
	return o
}

func (o LogAnalyticsConfigurationPtrOutput) Elem() LogAnalyticsConfigurationOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) LogAnalyticsConfiguration {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsConfiguration
		return ret
	}).(LogAnalyticsConfigurationOutput)
}

func (o LogAnalyticsConfigurationPtrOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CustomerId
	}).(pulumi.StringPtrOutput)
}

func (o LogAnalyticsConfigurationPtrOutput) SharedKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.SharedKey
	}).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationResponse struct {
	CustomerId *string `pulumi:"customerId"`
}

type LogAnalyticsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LogAnalyticsConfigurationResponse)(nil)).Elem()
}

func (o LogAnalyticsConfigurationResponseOutput) ToLogAnalyticsConfigurationResponseOutput() LogAnalyticsConfigurationResponseOutput {
	return o
}

func (o LogAnalyticsConfigurationResponseOutput) ToLogAnalyticsConfigurationResponseOutputWithContext(ctx context.Context) LogAnalyticsConfigurationResponseOutput {
	return o
}

func (o LogAnalyticsConfigurationResponseOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LogAnalyticsConfigurationResponse) *string { return v.CustomerId }).(pulumi.StringPtrOutput)
}

type LogAnalyticsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (LogAnalyticsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LogAnalyticsConfigurationResponse)(nil)).Elem()
}

func (o LogAnalyticsConfigurationResponsePtrOutput) ToLogAnalyticsConfigurationResponsePtrOutput() LogAnalyticsConfigurationResponsePtrOutput {
	return o
}

func (o LogAnalyticsConfigurationResponsePtrOutput) ToLogAnalyticsConfigurationResponsePtrOutputWithContext(ctx context.Context) LogAnalyticsConfigurationResponsePtrOutput {
	return o
}

func (o LogAnalyticsConfigurationResponsePtrOutput) Elem() LogAnalyticsConfigurationResponseOutput {
	return o.ApplyT(func(v *LogAnalyticsConfigurationResponse) LogAnalyticsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret LogAnalyticsConfigurationResponse
		return ret
	}).(LogAnalyticsConfigurationResponseOutput)
}

func (o LogAnalyticsConfigurationResponsePtrOutput) CustomerId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LogAnalyticsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomerId
	}).(pulumi.StringPtrOutput)
}

type QueueScaleRule struct {
	Auth        []ScaleRuleAuth `pulumi:"auth"`
	QueueLength *int            `pulumi:"queueLength"`
	QueueName   *string         `pulumi:"queueName"`
}





type QueueScaleRuleInput interface {
	pulumi.Input

	ToQueueScaleRuleOutput() QueueScaleRuleOutput
	ToQueueScaleRuleOutputWithContext(context.Context) QueueScaleRuleOutput
}

type QueueScaleRuleArgs struct {
	Auth        ScaleRuleAuthArrayInput `pulumi:"auth"`
	QueueLength pulumi.IntPtrInput      `pulumi:"queueLength"`
	QueueName   pulumi.StringPtrInput   `pulumi:"queueName"`
}

func (QueueScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueScaleRule)(nil)).Elem()
}

func (i QueueScaleRuleArgs) ToQueueScaleRuleOutput() QueueScaleRuleOutput {
	return i.ToQueueScaleRuleOutputWithContext(context.Background())
}

func (i QueueScaleRuleArgs) ToQueueScaleRuleOutputWithContext(ctx context.Context) QueueScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueScaleRuleOutput)
}

func (i QueueScaleRuleArgs) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return i.ToQueueScaleRulePtrOutputWithContext(context.Background())
}

func (i QueueScaleRuleArgs) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueScaleRuleOutput).ToQueueScaleRulePtrOutputWithContext(ctx)
}









type QueueScaleRulePtrInput interface {
	pulumi.Input

	ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput
	ToQueueScaleRulePtrOutputWithContext(context.Context) QueueScaleRulePtrOutput
}

type queueScaleRulePtrType QueueScaleRuleArgs

func QueueScaleRulePtr(v *QueueScaleRuleArgs) QueueScaleRulePtrInput {
	return (*queueScaleRulePtrType)(v)
}

func (*queueScaleRulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueScaleRule)(nil)).Elem()
}

func (i *queueScaleRulePtrType) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return i.ToQueueScaleRulePtrOutputWithContext(context.Background())
}

func (i *queueScaleRulePtrType) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(QueueScaleRulePtrOutput)
}

type QueueScaleRuleOutput struct{ *pulumi.OutputState }

func (QueueScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueScaleRule)(nil)).Elem()
}

func (o QueueScaleRuleOutput) ToQueueScaleRuleOutput() QueueScaleRuleOutput {
	return o
}

func (o QueueScaleRuleOutput) ToQueueScaleRuleOutputWithContext(ctx context.Context) QueueScaleRuleOutput {
	return o
}

func (o QueueScaleRuleOutput) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return o.ToQueueScaleRulePtrOutputWithContext(context.Background())
}

func (o QueueScaleRuleOutput) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v QueueScaleRule) *QueueScaleRule {
		return &v
	}).(QueueScaleRulePtrOutput)
}

func (o QueueScaleRuleOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v QueueScaleRule) []ScaleRuleAuth { return v.Auth }).(ScaleRuleAuthArrayOutput)
}

func (o QueueScaleRuleOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueScaleRule) *int { return v.QueueLength }).(pulumi.IntPtrOutput)
}

func (o QueueScaleRuleOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueScaleRule) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

type QueueScaleRulePtrOutput struct{ *pulumi.OutputState }

func (QueueScaleRulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueScaleRule)(nil)).Elem()
}

func (o QueueScaleRulePtrOutput) ToQueueScaleRulePtrOutput() QueueScaleRulePtrOutput {
	return o
}

func (o QueueScaleRulePtrOutput) ToQueueScaleRulePtrOutputWithContext(ctx context.Context) QueueScaleRulePtrOutput {
	return o
}

func (o QueueScaleRulePtrOutput) Elem() QueueScaleRuleOutput {
	return o.ApplyT(func(v *QueueScaleRule) QueueScaleRule {
		if v != nil {
			return *v
		}
		var ret QueueScaleRule
		return ret
	}).(QueueScaleRuleOutput)
}

func (o QueueScaleRulePtrOutput) Auth() ScaleRuleAuthArrayOutput {
	return o.ApplyT(func(v *QueueScaleRule) []ScaleRuleAuth {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthArrayOutput)
}

func (o QueueScaleRulePtrOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueScaleRule) *int {
		if v == nil {
			return nil
		}
		return v.QueueLength
	}).(pulumi.IntPtrOutput)
}

func (o QueueScaleRulePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueScaleRule) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

type QueueScaleRuleResponse struct {
	Auth        []ScaleRuleAuthResponse `pulumi:"auth"`
	QueueLength *int                    `pulumi:"queueLength"`
	QueueName   *string                 `pulumi:"queueName"`
}

type QueueScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (QueueScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*QueueScaleRuleResponse)(nil)).Elem()
}

func (o QueueScaleRuleResponseOutput) ToQueueScaleRuleResponseOutput() QueueScaleRuleResponseOutput {
	return o
}

func (o QueueScaleRuleResponseOutput) ToQueueScaleRuleResponseOutputWithContext(ctx context.Context) QueueScaleRuleResponseOutput {
	return o
}

func (o QueueScaleRuleResponseOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v QueueScaleRuleResponse) []ScaleRuleAuthResponse { return v.Auth }).(ScaleRuleAuthResponseArrayOutput)
}

func (o QueueScaleRuleResponseOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v QueueScaleRuleResponse) *int { return v.QueueLength }).(pulumi.IntPtrOutput)
}

func (o QueueScaleRuleResponseOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v QueueScaleRuleResponse) *string { return v.QueueName }).(pulumi.StringPtrOutput)
}

type QueueScaleRuleResponsePtrOutput struct{ *pulumi.OutputState }

func (QueueScaleRuleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**QueueScaleRuleResponse)(nil)).Elem()
}

func (o QueueScaleRuleResponsePtrOutput) ToQueueScaleRuleResponsePtrOutput() QueueScaleRuleResponsePtrOutput {
	return o
}

func (o QueueScaleRuleResponsePtrOutput) ToQueueScaleRuleResponsePtrOutputWithContext(ctx context.Context) QueueScaleRuleResponsePtrOutput {
	return o
}

func (o QueueScaleRuleResponsePtrOutput) Elem() QueueScaleRuleResponseOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) QueueScaleRuleResponse {
		if v != nil {
			return *v
		}
		var ret QueueScaleRuleResponse
		return ret
	}).(QueueScaleRuleResponseOutput)
}

func (o QueueScaleRuleResponsePtrOutput) Auth() ScaleRuleAuthResponseArrayOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) []ScaleRuleAuthResponse {
		if v == nil {
			return nil
		}
		return v.Auth
	}).(ScaleRuleAuthResponseArrayOutput)
}

func (o QueueScaleRuleResponsePtrOutput) QueueLength() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) *int {
		if v == nil {
			return nil
		}
		return v.QueueLength
	}).(pulumi.IntPtrOutput)
}

func (o QueueScaleRuleResponsePtrOutput) QueueName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *QueueScaleRuleResponse) *string {
		if v == nil {
			return nil
		}
		return v.QueueName
	}).(pulumi.StringPtrOutput)
}

type RegistryCredentials struct {
	PasswordSecretRef *string `pulumi:"passwordSecretRef"`
	Server            *string `pulumi:"server"`
	Username          *string `pulumi:"username"`
}





type RegistryCredentialsInput interface {
	pulumi.Input

	ToRegistryCredentialsOutput() RegistryCredentialsOutput
	ToRegistryCredentialsOutputWithContext(context.Context) RegistryCredentialsOutput
}

type RegistryCredentialsArgs struct {
	PasswordSecretRef pulumi.StringPtrInput `pulumi:"passwordSecretRef"`
	Server            pulumi.StringPtrInput `pulumi:"server"`
	Username          pulumi.StringPtrInput `pulumi:"username"`
}

func (RegistryCredentialsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryCredentials)(nil)).Elem()
}

func (i RegistryCredentialsArgs) ToRegistryCredentialsOutput() RegistryCredentialsOutput {
	return i.ToRegistryCredentialsOutputWithContext(context.Background())
}

func (i RegistryCredentialsArgs) ToRegistryCredentialsOutputWithContext(ctx context.Context) RegistryCredentialsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCredentialsOutput)
}





type RegistryCredentialsArrayInput interface {
	pulumi.Input

	ToRegistryCredentialsArrayOutput() RegistryCredentialsArrayOutput
	ToRegistryCredentialsArrayOutputWithContext(context.Context) RegistryCredentialsArrayOutput
}

type RegistryCredentialsArray []RegistryCredentialsInput

func (RegistryCredentialsArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryCredentials)(nil)).Elem()
}

func (i RegistryCredentialsArray) ToRegistryCredentialsArrayOutput() RegistryCredentialsArrayOutput {
	return i.ToRegistryCredentialsArrayOutputWithContext(context.Background())
}

func (i RegistryCredentialsArray) ToRegistryCredentialsArrayOutputWithContext(ctx context.Context) RegistryCredentialsArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RegistryCredentialsArrayOutput)
}

type RegistryCredentialsOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryCredentials)(nil)).Elem()
}

func (o RegistryCredentialsOutput) ToRegistryCredentialsOutput() RegistryCredentialsOutput {
	return o
}

func (o RegistryCredentialsOutput) ToRegistryCredentialsOutputWithContext(ctx context.Context) RegistryCredentialsOutput {
	return o
}

func (o RegistryCredentialsOutput) PasswordSecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.PasswordSecretRef }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentials) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type RegistryCredentialsArrayOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryCredentials)(nil)).Elem()
}

func (o RegistryCredentialsArrayOutput) ToRegistryCredentialsArrayOutput() RegistryCredentialsArrayOutput {
	return o
}

func (o RegistryCredentialsArrayOutput) ToRegistryCredentialsArrayOutputWithContext(ctx context.Context) RegistryCredentialsArrayOutput {
	return o
}

func (o RegistryCredentialsArrayOutput) Index(i pulumi.IntInput) RegistryCredentialsOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryCredentials {
		return vs[0].([]RegistryCredentials)[vs[1].(int)]
	}).(RegistryCredentialsOutput)
}

type RegistryCredentialsResponse struct {
	PasswordSecretRef *string `pulumi:"passwordSecretRef"`
	Server            *string `pulumi:"server"`
	Username          *string `pulumi:"username"`
}

type RegistryCredentialsResponseOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RegistryCredentialsResponse)(nil)).Elem()
}

func (o RegistryCredentialsResponseOutput) ToRegistryCredentialsResponseOutput() RegistryCredentialsResponseOutput {
	return o
}

func (o RegistryCredentialsResponseOutput) ToRegistryCredentialsResponseOutputWithContext(ctx context.Context) RegistryCredentialsResponseOutput {
	return o
}

func (o RegistryCredentialsResponseOutput) PasswordSecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.PasswordSecretRef }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsResponseOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.Server }).(pulumi.StringPtrOutput)
}

func (o RegistryCredentialsResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RegistryCredentialsResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type RegistryCredentialsResponseArrayOutput struct{ *pulumi.OutputState }

func (RegistryCredentialsResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RegistryCredentialsResponse)(nil)).Elem()
}

func (o RegistryCredentialsResponseArrayOutput) ToRegistryCredentialsResponseArrayOutput() RegistryCredentialsResponseArrayOutput {
	return o
}

func (o RegistryCredentialsResponseArrayOutput) ToRegistryCredentialsResponseArrayOutputWithContext(ctx context.Context) RegistryCredentialsResponseArrayOutput {
	return o
}

func (o RegistryCredentialsResponseArrayOutput) Index(i pulumi.IntInput) RegistryCredentialsResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RegistryCredentialsResponse {
		return vs[0].([]RegistryCredentialsResponse)[vs[1].(int)]
	}).(RegistryCredentialsResponseOutput)
}

type Scale struct {
	MaxReplicas *int        `pulumi:"maxReplicas"`
	MinReplicas *int        `pulumi:"minReplicas"`
	Rules       []ScaleRule `pulumi:"rules"`
}





type ScaleInput interface {
	pulumi.Input

	ToScaleOutput() ScaleOutput
	ToScaleOutputWithContext(context.Context) ScaleOutput
}

type ScaleArgs struct {
	MaxReplicas pulumi.IntPtrInput  `pulumi:"maxReplicas"`
	MinReplicas pulumi.IntPtrInput  `pulumi:"minReplicas"`
	Rules       ScaleRuleArrayInput `pulumi:"rules"`
}

func (ScaleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Scale)(nil)).Elem()
}

func (i ScaleArgs) ToScaleOutput() ScaleOutput {
	return i.ToScaleOutputWithContext(context.Background())
}

func (i ScaleArgs) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleOutput)
}

func (i ScaleArgs) ToScalePtrOutput() ScalePtrOutput {
	return i.ToScalePtrOutputWithContext(context.Background())
}

func (i ScaleArgs) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleOutput).ToScalePtrOutputWithContext(ctx)
}









type ScalePtrInput interface {
	pulumi.Input

	ToScalePtrOutput() ScalePtrOutput
	ToScalePtrOutputWithContext(context.Context) ScalePtrOutput
}

type scalePtrType ScaleArgs

func ScalePtr(v *ScaleArgs) ScalePtrInput {
	return (*scalePtrType)(v)
}

func (*scalePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Scale)(nil)).Elem()
}

func (i *scalePtrType) ToScalePtrOutput() ScalePtrOutput {
	return i.ToScalePtrOutputWithContext(context.Background())
}

func (i *scalePtrType) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScalePtrOutput)
}

type ScaleOutput struct{ *pulumi.OutputState }

func (ScaleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Scale)(nil)).Elem()
}

func (o ScaleOutput) ToScaleOutput() ScaleOutput {
	return o
}

func (o ScaleOutput) ToScaleOutputWithContext(ctx context.Context) ScaleOutput {
	return o
}

func (o ScaleOutput) ToScalePtrOutput() ScalePtrOutput {
	return o.ToScalePtrOutputWithContext(context.Background())
}

func (o ScaleOutput) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Scale) *Scale {
		return &v
	}).(ScalePtrOutput)
}

func (o ScaleOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Scale) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Scale) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleOutput) Rules() ScaleRuleArrayOutput {
	return o.ApplyT(func(v Scale) []ScaleRule { return v.Rules }).(ScaleRuleArrayOutput)
}

type ScalePtrOutput struct{ *pulumi.OutputState }

func (ScalePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Scale)(nil)).Elem()
}

func (o ScalePtrOutput) ToScalePtrOutput() ScalePtrOutput {
	return o
}

func (o ScalePtrOutput) ToScalePtrOutputWithContext(ctx context.Context) ScalePtrOutput {
	return o
}

func (o ScalePtrOutput) Elem() ScaleOutput {
	return o.ApplyT(func(v *Scale) Scale {
		if v != nil {
			return *v
		}
		var ret Scale
		return ret
	}).(ScaleOutput)
}

func (o ScalePtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Scale) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScalePtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *Scale) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScalePtrOutput) Rules() ScaleRuleArrayOutput {
	return o.ApplyT(func(v *Scale) []ScaleRule {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ScaleRuleArrayOutput)
}

type ScaleResponse struct {
	MaxReplicas *int                `pulumi:"maxReplicas"`
	MinReplicas *int                `pulumi:"minReplicas"`
	Rules       []ScaleRuleResponse `pulumi:"rules"`
}

type ScaleResponseOutput struct{ *pulumi.OutputState }

func (ScaleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleResponse)(nil)).Elem()
}

func (o ScaleResponseOutput) ToScaleResponseOutput() ScaleResponseOutput {
	return o
}

func (o ScaleResponseOutput) ToScaleResponseOutputWithContext(ctx context.Context) ScaleResponseOutput {
	return o
}

func (o ScaleResponseOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleResponse) *int { return v.MaxReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleResponseOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ScaleResponse) *int { return v.MinReplicas }).(pulumi.IntPtrOutput)
}

func (o ScaleResponseOutput) Rules() ScaleRuleResponseArrayOutput {
	return o.ApplyT(func(v ScaleResponse) []ScaleRuleResponse { return v.Rules }).(ScaleRuleResponseArrayOutput)
}

type ScaleResponsePtrOutput struct{ *pulumi.OutputState }

func (ScaleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ScaleResponse)(nil)).Elem()
}

func (o ScaleResponsePtrOutput) ToScaleResponsePtrOutput() ScaleResponsePtrOutput {
	return o
}

func (o ScaleResponsePtrOutput) ToScaleResponsePtrOutputWithContext(ctx context.Context) ScaleResponsePtrOutput {
	return o
}

func (o ScaleResponsePtrOutput) Elem() ScaleResponseOutput {
	return o.ApplyT(func(v *ScaleResponse) ScaleResponse {
		if v != nil {
			return *v
		}
		var ret ScaleResponse
		return ret
	}).(ScaleResponseOutput)
}

func (o ScaleResponsePtrOutput) MaxReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleResponse) *int {
		if v == nil {
			return nil
		}
		return v.MaxReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScaleResponsePtrOutput) MinReplicas() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *ScaleResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinReplicas
	}).(pulumi.IntPtrOutput)
}

func (o ScaleResponsePtrOutput) Rules() ScaleRuleResponseArrayOutput {
	return o.ApplyT(func(v *ScaleResponse) []ScaleRuleResponse {
		if v == nil {
			return nil
		}
		return v.Rules
	}).(ScaleRuleResponseArrayOutput)
}

type ScaleRule struct {
	AzureQueue *QueueScaleRule  `pulumi:"azureQueue"`
	Custom     *CustomScaleRule `pulumi:"custom"`
	Http       *HttpScaleRule   `pulumi:"http"`
	Name       *string          `pulumi:"name"`
}





type ScaleRuleInput interface {
	pulumi.Input

	ToScaleRuleOutput() ScaleRuleOutput
	ToScaleRuleOutputWithContext(context.Context) ScaleRuleOutput
}

type ScaleRuleArgs struct {
	AzureQueue QueueScaleRulePtrInput  `pulumi:"azureQueue"`
	Custom     CustomScaleRulePtrInput `pulumi:"custom"`
	Http       HttpScaleRulePtrInput   `pulumi:"http"`
	Name       pulumi.StringPtrInput   `pulumi:"name"`
}

func (ScaleRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArgs) ToScaleRuleOutput() ScaleRuleOutput {
	return i.ToScaleRuleOutputWithContext(context.Background())
}

func (i ScaleRuleArgs) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleOutput)
}





type ScaleRuleArrayInput interface {
	pulumi.Input

	ToScaleRuleArrayOutput() ScaleRuleArrayOutput
	ToScaleRuleArrayOutputWithContext(context.Context) ScaleRuleArrayOutput
}

type ScaleRuleArray []ScaleRuleInput

func (ScaleRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRule)(nil)).Elem()
}

func (i ScaleRuleArray) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return i.ToScaleRuleArrayOutputWithContext(context.Background())
}

func (i ScaleRuleArray) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleArrayOutput)
}

type ScaleRuleOutput struct{ *pulumi.OutputState }

func (ScaleRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRule)(nil)).Elem()
}

func (o ScaleRuleOutput) ToScaleRuleOutput() ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) ToScaleRuleOutputWithContext(ctx context.Context) ScaleRuleOutput {
	return o
}

func (o ScaleRuleOutput) AzureQueue() QueueScaleRulePtrOutput {
	return o.ApplyT(func(v ScaleRule) *QueueScaleRule { return v.AzureQueue }).(QueueScaleRulePtrOutput)
}

func (o ScaleRuleOutput) Custom() CustomScaleRulePtrOutput {
	return o.ApplyT(func(v ScaleRule) *CustomScaleRule { return v.Custom }).(CustomScaleRulePtrOutput)
}

func (o ScaleRuleOutput) Http() HttpScaleRulePtrOutput {
	return o.ApplyT(func(v ScaleRule) *HttpScaleRule { return v.Http }).(HttpScaleRulePtrOutput)
}

func (o ScaleRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScaleRuleArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRule)(nil)).Elem()
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutput() ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) ToScaleRuleArrayOutputWithContext(ctx context.Context) ScaleRuleArrayOutput {
	return o
}

func (o ScaleRuleArrayOutput) Index(i pulumi.IntInput) ScaleRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRule {
		return vs[0].([]ScaleRule)[vs[1].(int)]
	}).(ScaleRuleOutput)
}

type ScaleRuleAuth struct {
	SecretRef        *string `pulumi:"secretRef"`
	TriggerParameter *string `pulumi:"triggerParameter"`
}





type ScaleRuleAuthInput interface {
	pulumi.Input

	ToScaleRuleAuthOutput() ScaleRuleAuthOutput
	ToScaleRuleAuthOutputWithContext(context.Context) ScaleRuleAuthOutput
}

type ScaleRuleAuthArgs struct {
	SecretRef        pulumi.StringPtrInput `pulumi:"secretRef"`
	TriggerParameter pulumi.StringPtrInput `pulumi:"triggerParameter"`
}

func (ScaleRuleAuthArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleAuth)(nil)).Elem()
}

func (i ScaleRuleAuthArgs) ToScaleRuleAuthOutput() ScaleRuleAuthOutput {
	return i.ToScaleRuleAuthOutputWithContext(context.Background())
}

func (i ScaleRuleAuthArgs) ToScaleRuleAuthOutputWithContext(ctx context.Context) ScaleRuleAuthOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleAuthOutput)
}





type ScaleRuleAuthArrayInput interface {
	pulumi.Input

	ToScaleRuleAuthArrayOutput() ScaleRuleAuthArrayOutput
	ToScaleRuleAuthArrayOutputWithContext(context.Context) ScaleRuleAuthArrayOutput
}

type ScaleRuleAuthArray []ScaleRuleAuthInput

func (ScaleRuleAuthArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleAuth)(nil)).Elem()
}

func (i ScaleRuleAuthArray) ToScaleRuleAuthArrayOutput() ScaleRuleAuthArrayOutput {
	return i.ToScaleRuleAuthArrayOutputWithContext(context.Background())
}

func (i ScaleRuleAuthArray) ToScaleRuleAuthArrayOutputWithContext(ctx context.Context) ScaleRuleAuthArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ScaleRuleAuthArrayOutput)
}

type ScaleRuleAuthOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleAuth)(nil)).Elem()
}

func (o ScaleRuleAuthOutput) ToScaleRuleAuthOutput() ScaleRuleAuthOutput {
	return o
}

func (o ScaleRuleAuthOutput) ToScaleRuleAuthOutputWithContext(ctx context.Context) ScaleRuleAuthOutput {
	return o
}

func (o ScaleRuleAuthOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuth) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o ScaleRuleAuthOutput) TriggerParameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuth) *string { return v.TriggerParameter }).(pulumi.StringPtrOutput)
}

type ScaleRuleAuthArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleAuth)(nil)).Elem()
}

func (o ScaleRuleAuthArrayOutput) ToScaleRuleAuthArrayOutput() ScaleRuleAuthArrayOutput {
	return o
}

func (o ScaleRuleAuthArrayOutput) ToScaleRuleAuthArrayOutputWithContext(ctx context.Context) ScaleRuleAuthArrayOutput {
	return o
}

func (o ScaleRuleAuthArrayOutput) Index(i pulumi.IntInput) ScaleRuleAuthOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleAuth {
		return vs[0].([]ScaleRuleAuth)[vs[1].(int)]
	}).(ScaleRuleAuthOutput)
}

type ScaleRuleAuthResponse struct {
	SecretRef        *string `pulumi:"secretRef"`
	TriggerParameter *string `pulumi:"triggerParameter"`
}

type ScaleRuleAuthResponseOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleAuthResponse)(nil)).Elem()
}

func (o ScaleRuleAuthResponseOutput) ToScaleRuleAuthResponseOutput() ScaleRuleAuthResponseOutput {
	return o
}

func (o ScaleRuleAuthResponseOutput) ToScaleRuleAuthResponseOutputWithContext(ctx context.Context) ScaleRuleAuthResponseOutput {
	return o
}

func (o ScaleRuleAuthResponseOutput) SecretRef() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuthResponse) *string { return v.SecretRef }).(pulumi.StringPtrOutput)
}

func (o ScaleRuleAuthResponseOutput) TriggerParameter() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleAuthResponse) *string { return v.TriggerParameter }).(pulumi.StringPtrOutput)
}

type ScaleRuleAuthResponseArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleAuthResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleAuthResponse)(nil)).Elem()
}

func (o ScaleRuleAuthResponseArrayOutput) ToScaleRuleAuthResponseArrayOutput() ScaleRuleAuthResponseArrayOutput {
	return o
}

func (o ScaleRuleAuthResponseArrayOutput) ToScaleRuleAuthResponseArrayOutputWithContext(ctx context.Context) ScaleRuleAuthResponseArrayOutput {
	return o
}

func (o ScaleRuleAuthResponseArrayOutput) Index(i pulumi.IntInput) ScaleRuleAuthResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleAuthResponse {
		return vs[0].([]ScaleRuleAuthResponse)[vs[1].(int)]
	}).(ScaleRuleAuthResponseOutput)
}

type ScaleRuleResponse struct {
	AzureQueue *QueueScaleRuleResponse  `pulumi:"azureQueue"`
	Custom     *CustomScaleRuleResponse `pulumi:"custom"`
	Http       *HttpScaleRuleResponse   `pulumi:"http"`
	Name       *string                  `pulumi:"name"`
}

type ScaleRuleResponseOutput struct{ *pulumi.OutputState }

func (ScaleRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ScaleRuleResponse)(nil)).Elem()
}

func (o ScaleRuleResponseOutput) ToScaleRuleResponseOutput() ScaleRuleResponseOutput {
	return o
}

func (o ScaleRuleResponseOutput) ToScaleRuleResponseOutputWithContext(ctx context.Context) ScaleRuleResponseOutput {
	return o
}

func (o ScaleRuleResponseOutput) AzureQueue() QueueScaleRuleResponsePtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *QueueScaleRuleResponse { return v.AzureQueue }).(QueueScaleRuleResponsePtrOutput)
}

func (o ScaleRuleResponseOutput) Custom() CustomScaleRuleResponsePtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *CustomScaleRuleResponse { return v.Custom }).(CustomScaleRuleResponsePtrOutput)
}

func (o ScaleRuleResponseOutput) Http() HttpScaleRuleResponsePtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *HttpScaleRuleResponse { return v.Http }).(HttpScaleRuleResponsePtrOutput)
}

func (o ScaleRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ScaleRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type ScaleRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (ScaleRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ScaleRuleResponse)(nil)).Elem()
}

func (o ScaleRuleResponseArrayOutput) ToScaleRuleResponseArrayOutput() ScaleRuleResponseArrayOutput {
	return o
}

func (o ScaleRuleResponseArrayOutput) ToScaleRuleResponseArrayOutputWithContext(ctx context.Context) ScaleRuleResponseArrayOutput {
	return o
}

func (o ScaleRuleResponseArrayOutput) Index(i pulumi.IntInput) ScaleRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ScaleRuleResponse {
		return vs[0].([]ScaleRuleResponse)[vs[1].(int)]
	}).(ScaleRuleResponseOutput)
}

type Secret struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SecretInput interface {
	pulumi.Input

	ToSecretOutput() SecretOutput
	ToSecretOutputWithContext(context.Context) SecretOutput
}

type SecretArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SecretArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil)).Elem()
}

func (i SecretArgs) ToSecretOutput() SecretOutput {
	return i.ToSecretOutputWithContext(context.Background())
}

func (i SecretArgs) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretOutput)
}





type SecretArrayInput interface {
	pulumi.Input

	ToSecretArrayOutput() SecretArrayOutput
	ToSecretArrayOutputWithContext(context.Context) SecretArrayOutput
}

type SecretArray []SecretInput

func (SecretArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Secret)(nil)).Elem()
}

func (i SecretArray) ToSecretArrayOutput() SecretArrayOutput {
	return i.ToSecretArrayOutputWithContext(context.Background())
}

func (i SecretArray) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SecretArrayOutput)
}

type SecretOutput struct{ *pulumi.OutputState }

func (SecretOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Secret)(nil)).Elem()
}

func (o SecretOutput) ToSecretOutput() SecretOutput {
	return o
}

func (o SecretOutput) ToSecretOutputWithContext(ctx context.Context) SecretOutput {
	return o
}

func (o SecretOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Secret) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SecretOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Secret) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SecretArrayOutput struct{ *pulumi.OutputState }

func (SecretArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Secret)(nil)).Elem()
}

func (o SecretArrayOutput) ToSecretArrayOutput() SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) ToSecretArrayOutputWithContext(ctx context.Context) SecretArrayOutput {
	return o
}

func (o SecretArrayOutput) Index(i pulumi.IntInput) SecretOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Secret {
		return vs[0].([]Secret)[vs[1].(int)]
	}).(SecretOutput)
}

type SecretResponse struct {
	Name *string `pulumi:"name"`
}

type SecretResponseOutput struct{ *pulumi.OutputState }

func (SecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SecretResponse)(nil)).Elem()
}

func (o SecretResponseOutput) ToSecretResponseOutput() SecretResponseOutput {
	return o
}

func (o SecretResponseOutput) ToSecretResponseOutputWithContext(ctx context.Context) SecretResponseOutput {
	return o
}

func (o SecretResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SecretResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type SecretResponseArrayOutput struct{ *pulumi.OutputState }

func (SecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SecretResponse)(nil)).Elem()
}

func (o SecretResponseArrayOutput) ToSecretResponseArrayOutput() SecretResponseArrayOutput {
	return o
}

func (o SecretResponseArrayOutput) ToSecretResponseArrayOutputWithContext(ctx context.Context) SecretResponseArrayOutput {
	return o
}

func (o SecretResponseArrayOutput) Index(i pulumi.IntInput) SecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SecretResponse {
		return vs[0].([]SecretResponse)[vs[1].(int)]
	}).(SecretResponseOutput)
}

type Template struct {
	Containers     []Container `pulumi:"containers"`
	Dapr           *Dapr       `pulumi:"dapr"`
	RevisionSuffix *string     `pulumi:"revisionSuffix"`
	Scale          *Scale      `pulumi:"scale"`
}





type TemplateInput interface {
	pulumi.Input

	ToTemplateOutput() TemplateOutput
	ToTemplateOutputWithContext(context.Context) TemplateOutput
}

type TemplateArgs struct {
	Containers     ContainerArrayInput   `pulumi:"containers"`
	Dapr           DaprPtrInput          `pulumi:"dapr"`
	RevisionSuffix pulumi.StringPtrInput `pulumi:"revisionSuffix"`
	Scale          ScalePtrInput         `pulumi:"scale"`
}

func (TemplateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil)).Elem()
}

func (i TemplateArgs) ToTemplateOutput() TemplateOutput {
	return i.ToTemplateOutputWithContext(context.Background())
}

func (i TemplateArgs) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput)
}

func (i TemplateArgs) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i TemplateArgs) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplateOutput).ToTemplatePtrOutputWithContext(ctx)
}









type TemplatePtrInput interface {
	pulumi.Input

	ToTemplatePtrOutput() TemplatePtrOutput
	ToTemplatePtrOutputWithContext(context.Context) TemplatePtrOutput
}

type templatePtrType TemplateArgs

func TemplatePtr(v *TemplateArgs) TemplatePtrInput {
	return (*templatePtrType)(v)
}

func (*templatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (i *templatePtrType) ToTemplatePtrOutput() TemplatePtrOutput {
	return i.ToTemplatePtrOutputWithContext(context.Background())
}

func (i *templatePtrType) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TemplatePtrOutput)
}

type TemplateOutput struct{ *pulumi.OutputState }

func (TemplateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Template)(nil)).Elem()
}

func (o TemplateOutput) ToTemplateOutput() TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplateOutputWithContext(ctx context.Context) TemplateOutput {
	return o
}

func (o TemplateOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o.ToTemplatePtrOutputWithContext(context.Background())
}

func (o TemplateOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Template) *Template {
		return &v
	}).(TemplatePtrOutput)
}

func (o TemplateOutput) Containers() ContainerArrayOutput {
	return o.ApplyT(func(v Template) []Container { return v.Containers }).(ContainerArrayOutput)
}

func (o TemplateOutput) Dapr() DaprPtrOutput {
	return o.ApplyT(func(v Template) *Dapr { return v.Dapr }).(DaprPtrOutput)
}

func (o TemplateOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Template) *string { return v.RevisionSuffix }).(pulumi.StringPtrOutput)
}

func (o TemplateOutput) Scale() ScalePtrOutput {
	return o.ApplyT(func(v Template) *Scale { return v.Scale }).(ScalePtrOutput)
}

type TemplatePtrOutput struct{ *pulumi.OutputState }

func (TemplatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Template)(nil)).Elem()
}

func (o TemplatePtrOutput) ToTemplatePtrOutput() TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) ToTemplatePtrOutputWithContext(ctx context.Context) TemplatePtrOutput {
	return o
}

func (o TemplatePtrOutput) Elem() TemplateOutput {
	return o.ApplyT(func(v *Template) Template {
		if v != nil {
			return *v
		}
		var ret Template
		return ret
	}).(TemplateOutput)
}

func (o TemplatePtrOutput) Containers() ContainerArrayOutput {
	return o.ApplyT(func(v *Template) []Container {
		if v == nil {
			return nil
		}
		return v.Containers
	}).(ContainerArrayOutput)
}

func (o TemplatePtrOutput) Dapr() DaprPtrOutput {
	return o.ApplyT(func(v *Template) *Dapr {
		if v == nil {
			return nil
		}
		return v.Dapr
	}).(DaprPtrOutput)
}

func (o TemplatePtrOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Template) *string {
		if v == nil {
			return nil
		}
		return v.RevisionSuffix
	}).(pulumi.StringPtrOutput)
}

func (o TemplatePtrOutput) Scale() ScalePtrOutput {
	return o.ApplyT(func(v *Template) *Scale {
		if v == nil {
			return nil
		}
		return v.Scale
	}).(ScalePtrOutput)
}

type TemplateResponse struct {
	Containers     []ContainerResponse `pulumi:"containers"`
	Dapr           *DaprResponse       `pulumi:"dapr"`
	RevisionSuffix *string             `pulumi:"revisionSuffix"`
	Scale          *ScaleResponse      `pulumi:"scale"`
}

type TemplateResponseOutput struct{ *pulumi.OutputState }

func (TemplateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TemplateResponse)(nil)).Elem()
}

func (o TemplateResponseOutput) ToTemplateResponseOutput() TemplateResponseOutput {
	return o
}

func (o TemplateResponseOutput) ToTemplateResponseOutputWithContext(ctx context.Context) TemplateResponseOutput {
	return o
}

func (o TemplateResponseOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v TemplateResponse) []ContainerResponse { return v.Containers }).(ContainerResponseArrayOutput)
}

func (o TemplateResponseOutput) Dapr() DaprResponsePtrOutput {
	return o.ApplyT(func(v TemplateResponse) *DaprResponse { return v.Dapr }).(DaprResponsePtrOutput)
}

func (o TemplateResponseOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TemplateResponse) *string { return v.RevisionSuffix }).(pulumi.StringPtrOutput)
}

func (o TemplateResponseOutput) Scale() ScaleResponsePtrOutput {
	return o.ApplyT(func(v TemplateResponse) *ScaleResponse { return v.Scale }).(ScaleResponsePtrOutput)
}

type TemplateResponsePtrOutput struct{ *pulumi.OutputState }

func (TemplateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TemplateResponse)(nil)).Elem()
}

func (o TemplateResponsePtrOutput) ToTemplateResponsePtrOutput() TemplateResponsePtrOutput {
	return o
}

func (o TemplateResponsePtrOutput) ToTemplateResponsePtrOutputWithContext(ctx context.Context) TemplateResponsePtrOutput {
	return o
}

func (o TemplateResponsePtrOutput) Elem() TemplateResponseOutput {
	return o.ApplyT(func(v *TemplateResponse) TemplateResponse {
		if v != nil {
			return *v
		}
		var ret TemplateResponse
		return ret
	}).(TemplateResponseOutput)
}

func (o TemplateResponsePtrOutput) Containers() ContainerResponseArrayOutput {
	return o.ApplyT(func(v *TemplateResponse) []ContainerResponse {
		if v == nil {
			return nil
		}
		return v.Containers
	}).(ContainerResponseArrayOutput)
}

func (o TemplateResponsePtrOutput) Dapr() DaprResponsePtrOutput {
	return o.ApplyT(func(v *TemplateResponse) *DaprResponse {
		if v == nil {
			return nil
		}
		return v.Dapr
	}).(DaprResponsePtrOutput)
}

func (o TemplateResponsePtrOutput) RevisionSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TemplateResponse) *string {
		if v == nil {
			return nil
		}
		return v.RevisionSuffix
	}).(pulumi.StringPtrOutput)
}

func (o TemplateResponsePtrOutput) Scale() ScaleResponsePtrOutput {
	return o.ApplyT(func(v *TemplateResponse) *ScaleResponse {
		if v == nil {
			return nil
		}
		return v.Scale
	}).(ScaleResponsePtrOutput)
}

type TrafficWeight struct {
	LatestRevision *bool   `pulumi:"latestRevision"`
	RevisionName   *string `pulumi:"revisionName"`
	Weight         *int    `pulumi:"weight"`
}


func (val *TrafficWeight) Defaults() *TrafficWeight {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LatestRevision) {
		latestRevision_ := false
		tmp.LatestRevision = &latestRevision_
	}
	return &tmp
}





type TrafficWeightInput interface {
	pulumi.Input

	ToTrafficWeightOutput() TrafficWeightOutput
	ToTrafficWeightOutputWithContext(context.Context) TrafficWeightOutput
}

type TrafficWeightArgs struct {
	LatestRevision pulumi.BoolPtrInput   `pulumi:"latestRevision"`
	RevisionName   pulumi.StringPtrInput `pulumi:"revisionName"`
	Weight         pulumi.IntPtrInput    `pulumi:"weight"`
}

func (TrafficWeightArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficWeight)(nil)).Elem()
}

func (i TrafficWeightArgs) ToTrafficWeightOutput() TrafficWeightOutput {
	return i.ToTrafficWeightOutputWithContext(context.Background())
}

func (i TrafficWeightArgs) ToTrafficWeightOutputWithContext(ctx context.Context) TrafficWeightOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficWeightOutput)
}





type TrafficWeightArrayInput interface {
	pulumi.Input

	ToTrafficWeightArrayOutput() TrafficWeightArrayOutput
	ToTrafficWeightArrayOutputWithContext(context.Context) TrafficWeightArrayOutput
}

type TrafficWeightArray []TrafficWeightInput

func (TrafficWeightArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficWeight)(nil)).Elem()
}

func (i TrafficWeightArray) ToTrafficWeightArrayOutput() TrafficWeightArrayOutput {
	return i.ToTrafficWeightArrayOutputWithContext(context.Background())
}

func (i TrafficWeightArray) ToTrafficWeightArrayOutputWithContext(ctx context.Context) TrafficWeightArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TrafficWeightArrayOutput)
}

type TrafficWeightOutput struct{ *pulumi.OutputState }

func (TrafficWeightOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficWeight)(nil)).Elem()
}

func (o TrafficWeightOutput) ToTrafficWeightOutput() TrafficWeightOutput {
	return o
}

func (o TrafficWeightOutput) ToTrafficWeightOutputWithContext(ctx context.Context) TrafficWeightOutput {
	return o
}

func (o TrafficWeightOutput) LatestRevision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *bool { return v.LatestRevision }).(pulumi.BoolPtrOutput)
}

func (o TrafficWeightOutput) RevisionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *string { return v.RevisionName }).(pulumi.StringPtrOutput)
}

func (o TrafficWeightOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TrafficWeight) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type TrafficWeightArrayOutput struct{ *pulumi.OutputState }

func (TrafficWeightArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficWeight)(nil)).Elem()
}

func (o TrafficWeightArrayOutput) ToTrafficWeightArrayOutput() TrafficWeightArrayOutput {
	return o
}

func (o TrafficWeightArrayOutput) ToTrafficWeightArrayOutputWithContext(ctx context.Context) TrafficWeightArrayOutput {
	return o
}

func (o TrafficWeightArrayOutput) Index(i pulumi.IntInput) TrafficWeightOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficWeight {
		return vs[0].([]TrafficWeight)[vs[1].(int)]
	}).(TrafficWeightOutput)
}

type TrafficWeightResponse struct {
	LatestRevision *bool   `pulumi:"latestRevision"`
	RevisionName   *string `pulumi:"revisionName"`
	Weight         *int    `pulumi:"weight"`
}


func (val *TrafficWeightResponse) Defaults() *TrafficWeightResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LatestRevision) {
		latestRevision_ := false
		tmp.LatestRevision = &latestRevision_
	}
	return &tmp
}

type TrafficWeightResponseOutput struct{ *pulumi.OutputState }

func (TrafficWeightResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TrafficWeightResponse)(nil)).Elem()
}

func (o TrafficWeightResponseOutput) ToTrafficWeightResponseOutput() TrafficWeightResponseOutput {
	return o
}

func (o TrafficWeightResponseOutput) ToTrafficWeightResponseOutputWithContext(ctx context.Context) TrafficWeightResponseOutput {
	return o
}

func (o TrafficWeightResponseOutput) LatestRevision() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *bool { return v.LatestRevision }).(pulumi.BoolPtrOutput)
}

func (o TrafficWeightResponseOutput) RevisionName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *string { return v.RevisionName }).(pulumi.StringPtrOutput)
}

func (o TrafficWeightResponseOutput) Weight() pulumi.IntPtrOutput {
	return o.ApplyT(func(v TrafficWeightResponse) *int { return v.Weight }).(pulumi.IntPtrOutput)
}

type TrafficWeightResponseArrayOutput struct{ *pulumi.OutputState }

func (TrafficWeightResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]TrafficWeightResponse)(nil)).Elem()
}

func (o TrafficWeightResponseArrayOutput) ToTrafficWeightResponseArrayOutput() TrafficWeightResponseArrayOutput {
	return o
}

func (o TrafficWeightResponseArrayOutput) ToTrafficWeightResponseArrayOutputWithContext(ctx context.Context) TrafficWeightResponseArrayOutput {
	return o
}

func (o TrafficWeightResponseArrayOutput) Index(i pulumi.IntInput) TrafficWeightResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) TrafficWeightResponse {
		return vs[0].([]TrafficWeightResponse)[vs[1].(int)]
	}).(TrafficWeightResponseOutput)
}

func init() {
	pulumi.RegisterOutputType(AppLogsConfigurationOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ArcConfigurationOutput{})
	pulumi.RegisterOutputType(ArcConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ArcConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ArcConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerResourcesOutput{})
	pulumi.RegisterOutputType(ContainerResourcesPtrOutput{})
	pulumi.RegisterOutputType(ContainerResourcesResponseOutput{})
	pulumi.RegisterOutputType(ContainerResourcesResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResponseOutput{})
	pulumi.RegisterOutputType(ContainerResponseArrayOutput{})
	pulumi.RegisterOutputType(CustomScaleRuleOutput{})
	pulumi.RegisterOutputType(CustomScaleRulePtrOutput{})
	pulumi.RegisterOutputType(CustomScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(CustomScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(DaprOutput{})
	pulumi.RegisterOutputType(DaprPtrOutput{})
	pulumi.RegisterOutputType(DaprComponentOutput{})
	pulumi.RegisterOutputType(DaprComponentArrayOutput{})
	pulumi.RegisterOutputType(DaprComponentResponseOutput{})
	pulumi.RegisterOutputType(DaprComponentResponseArrayOutput{})
	pulumi.RegisterOutputType(DaprMetadataOutput{})
	pulumi.RegisterOutputType(DaprMetadataArrayOutput{})
	pulumi.RegisterOutputType(DaprMetadataResponseOutput{})
	pulumi.RegisterOutputType(DaprMetadataResponseArrayOutput{})
	pulumi.RegisterOutputType(DaprResponseOutput{})
	pulumi.RegisterOutputType(DaprResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentVarOutput{})
	pulumi.RegisterOutputType(EnvironmentVarArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVarResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVarResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleOutput{})
	pulumi.RegisterOutputType(HttpScaleRulePtrOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(IngressOutput{})
	pulumi.RegisterOutputType(IngressPtrOutput{})
	pulumi.RegisterOutputType(IngressResponseOutput{})
	pulumi.RegisterOutputType(IngressResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleOutput{})
	pulumi.RegisterOutputType(QueueScaleRulePtrOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsArrayOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsResponseOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsResponseArrayOutput{})
	pulumi.RegisterOutputType(ScaleOutput{})
	pulumi.RegisterOutputType(ScalePtrOutput{})
	pulumi.RegisterOutputType(ScaleResponseOutput{})
	pulumi.RegisterOutputType(ScaleResponsePtrOutput{})
	pulumi.RegisterOutputType(ScaleRuleOutput{})
	pulumi.RegisterOutputType(ScaleRuleArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleAuthResponseArrayOutput{})
	pulumi.RegisterOutputType(ScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(ScaleRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(SecretOutput{})
	pulumi.RegisterOutputType(SecretArrayOutput{})
	pulumi.RegisterOutputType(SecretResponseOutput{})
	pulumi.RegisterOutputType(SecretResponseArrayOutput{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplatePtrOutput{})
	pulumi.RegisterOutputType(TemplateResponseOutput{})
	pulumi.RegisterOutputType(TemplateResponsePtrOutput{})
	pulumi.RegisterOutputType(TrafficWeightOutput{})
	pulumi.RegisterOutputType(TrafficWeightArrayOutput{})
	pulumi.RegisterOutputType(TrafficWeightResponseOutput{})
	pulumi.RegisterOutputType(TrafficWeightResponseArrayOutput{})
}
