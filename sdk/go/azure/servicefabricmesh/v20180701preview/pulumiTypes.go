


package v20180701preview

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AzureInternalMonitoringPipelineSinkDescription struct {
	AccountName      *string     `pulumi:"accountName"`
	AutoKeyConfigUrl *string     `pulumi:"autoKeyConfigUrl"`
	Description      *string     `pulumi:"description"`
	FluentdConfigUrl interface{} `pulumi:"fluentdConfigUrl"`
	Kind             string      `pulumi:"kind"`
	MaConfigUrl      *string     `pulumi:"maConfigUrl"`
	Name             *string     `pulumi:"name"`
	Namespace        *string     `pulumi:"namespace"`
}





type AzureInternalMonitoringPipelineSinkDescriptionInput interface {
	pulumi.Input

	ToAzureInternalMonitoringPipelineSinkDescriptionOutput() AzureInternalMonitoringPipelineSinkDescriptionOutput
	ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(context.Context) AzureInternalMonitoringPipelineSinkDescriptionOutput
}

type AzureInternalMonitoringPipelineSinkDescriptionArgs struct {
	AccountName      pulumi.StringPtrInput `pulumi:"accountName"`
	AutoKeyConfigUrl pulumi.StringPtrInput `pulumi:"autoKeyConfigUrl"`
	Description      pulumi.StringPtrInput `pulumi:"description"`
	FluentdConfigUrl pulumi.Input          `pulumi:"fluentdConfigUrl"`
	Kind             pulumi.StringInput    `pulumi:"kind"`
	MaConfigUrl      pulumi.StringPtrInput `pulumi:"maConfigUrl"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	Namespace        pulumi.StringPtrInput `pulumi:"namespace"`
}

func (AzureInternalMonitoringPipelineSinkDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArgs) ToAzureInternalMonitoringPipelineSinkDescriptionOutput() AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return i.ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(context.Background())
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArgs) ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureInternalMonitoringPipelineSinkDescriptionOutput)
}





type AzureInternalMonitoringPipelineSinkDescriptionArrayInput interface {
	pulumi.Input

	ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput
	ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(context.Context) AzureInternalMonitoringPipelineSinkDescriptionArrayOutput
}

type AzureInternalMonitoringPipelineSinkDescriptionArray []AzureInternalMonitoringPipelineSinkDescriptionInput

func (AzureInternalMonitoringPipelineSinkDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArray) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return i.ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(context.Background())
}

func (i AzureInternalMonitoringPipelineSinkDescriptionArray) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) ToAzureInternalMonitoringPipelineSinkDescriptionOutput() AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) ToAzureInternalMonitoringPipelineSinkDescriptionOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) AutoKeyConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.AutoKeyConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) FluentdConfigUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) interface{} { return v.FluentdConfigUrl }).(pulumi.AnyOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) MaConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.MaConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescription) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionArrayOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureInternalMonitoringPipelineSinkDescription)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionArrayOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionArrayOutput) Index(i pulumi.IntInput) AzureInternalMonitoringPipelineSinkDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureInternalMonitoringPipelineSinkDescription {
		return vs[0].([]AzureInternalMonitoringPipelineSinkDescription)[vs[1].(int)]
	}).(AzureInternalMonitoringPipelineSinkDescriptionOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionResponse struct {
	AccountName      *string     `pulumi:"accountName"`
	AutoKeyConfigUrl *string     `pulumi:"autoKeyConfigUrl"`
	Description      *string     `pulumi:"description"`
	FluentdConfigUrl interface{} `pulumi:"fluentdConfigUrl"`
	Kind             string      `pulumi:"kind"`
	MaConfigUrl      *string     `pulumi:"maConfigUrl"`
	Name             *string     `pulumi:"name"`
	Namespace        *string     `pulumi:"namespace"`
}

type AzureInternalMonitoringPipelineSinkDescriptionResponseOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureInternalMonitoringPipelineSinkDescriptionResponse)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseOutput() AzureInternalMonitoringPipelineSinkDescriptionResponseOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionResponseOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) AutoKeyConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.AutoKeyConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) FluentdConfigUrl() pulumi.AnyOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) interface{} { return v.FluentdConfigUrl }).(pulumi.AnyOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Kind() pulumi.StringOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) string { return v.Kind }).(pulumi.StringOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) MaConfigUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.MaConfigUrl }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureInternalMonitoringPipelineSinkDescriptionResponse) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

type AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureInternalMonitoringPipelineSinkDescriptionResponse)(nil)).Elem()
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput() AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) ToAzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutputWithContext(ctx context.Context) AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o
}

func (o AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput) Index(i pulumi.IntInput) AzureInternalMonitoringPipelineSinkDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureInternalMonitoringPipelineSinkDescriptionResponse {
		return vs[0].([]AzureInternalMonitoringPipelineSinkDescriptionResponse)[vs[1].(int)]
	}).(AzureInternalMonitoringPipelineSinkDescriptionResponseOutput)
}

type ContainerCodePackageProperties struct {
	Commands                []string                 `pulumi:"commands"`
	Diagnostics             *DiagnosticsRef          `pulumi:"diagnostics"`
	Endpoints               []EndpointProperties     `pulumi:"endpoints"`
	Entrypoint              *string                  `pulumi:"entrypoint"`
	EnvironmentVariables    []EnvironmentVariable    `pulumi:"environmentVariables"`
	Image                   string                   `pulumi:"image"`
	ImageRegistryCredential *ImageRegistryCredential `pulumi:"imageRegistryCredential"`
	Labels                  []ContainerLabel         `pulumi:"labels"`
	Name                    string                   `pulumi:"name"`
	Resources               ResourceRequirements     `pulumi:"resources"`
	Settings                []Setting                `pulumi:"settings"`
	VolumeRefs              []ContainerVolume        `pulumi:"volumeRefs"`
}





type ContainerCodePackagePropertiesInput interface {
	pulumi.Input

	ToContainerCodePackagePropertiesOutput() ContainerCodePackagePropertiesOutput
	ToContainerCodePackagePropertiesOutputWithContext(context.Context) ContainerCodePackagePropertiesOutput
}

type ContainerCodePackagePropertiesArgs struct {
	Commands                pulumi.StringArrayInput         `pulumi:"commands"`
	Diagnostics             DiagnosticsRefPtrInput          `pulumi:"diagnostics"`
	Endpoints               EndpointPropertiesArrayInput    `pulumi:"endpoints"`
	Entrypoint              pulumi.StringPtrInput           `pulumi:"entrypoint"`
	EnvironmentVariables    EnvironmentVariableArrayInput   `pulumi:"environmentVariables"`
	Image                   pulumi.StringInput              `pulumi:"image"`
	ImageRegistryCredential ImageRegistryCredentialPtrInput `pulumi:"imageRegistryCredential"`
	Labels                  ContainerLabelArrayInput        `pulumi:"labels"`
	Name                    pulumi.StringInput              `pulumi:"name"`
	Resources               ResourceRequirementsInput       `pulumi:"resources"`
	Settings                SettingArrayInput               `pulumi:"settings"`
	VolumeRefs              ContainerVolumeArrayInput       `pulumi:"volumeRefs"`
}

func (ContainerCodePackagePropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerCodePackageProperties)(nil)).Elem()
}

func (i ContainerCodePackagePropertiesArgs) ToContainerCodePackagePropertiesOutput() ContainerCodePackagePropertiesOutput {
	return i.ToContainerCodePackagePropertiesOutputWithContext(context.Background())
}

func (i ContainerCodePackagePropertiesArgs) ToContainerCodePackagePropertiesOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCodePackagePropertiesOutput)
}





type ContainerCodePackagePropertiesArrayInput interface {
	pulumi.Input

	ToContainerCodePackagePropertiesArrayOutput() ContainerCodePackagePropertiesArrayOutput
	ToContainerCodePackagePropertiesArrayOutputWithContext(context.Context) ContainerCodePackagePropertiesArrayOutput
}

type ContainerCodePackagePropertiesArray []ContainerCodePackagePropertiesInput

func (ContainerCodePackagePropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerCodePackageProperties)(nil)).Elem()
}

func (i ContainerCodePackagePropertiesArray) ToContainerCodePackagePropertiesArrayOutput() ContainerCodePackagePropertiesArrayOutput {
	return i.ToContainerCodePackagePropertiesArrayOutputWithContext(context.Background())
}

func (i ContainerCodePackagePropertiesArray) ToContainerCodePackagePropertiesArrayOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerCodePackagePropertiesArrayOutput)
}

type ContainerCodePackagePropertiesOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerCodePackageProperties)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesOutput) ToContainerCodePackagePropertiesOutput() ContainerCodePackagePropertiesOutput {
	return o
}

func (o ContainerCodePackagePropertiesOutput) ToContainerCodePackagePropertiesOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesOutput {
	return o
}

func (o ContainerCodePackagePropertiesOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Diagnostics() DiagnosticsRefPtrOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) *DiagnosticsRef { return v.Diagnostics }).(DiagnosticsRefPtrOutput)
}

func (o ContainerCodePackagePropertiesOutput) Endpoints() EndpointPropertiesArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []EndpointProperties { return v.Endpoints }).(EndpointPropertiesArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Entrypoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) *string { return v.Entrypoint }).(pulumi.StringPtrOutput)
}

func (o ContainerCodePackagePropertiesOutput) EnvironmentVariables() EnvironmentVariableArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []EnvironmentVariable { return v.EnvironmentVariables }).(EnvironmentVariableArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) string { return v.Image }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesOutput) ImageRegistryCredential() ImageRegistryCredentialPtrOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) *ImageRegistryCredential { return v.ImageRegistryCredential }).(ImageRegistryCredentialPtrOutput)
}

func (o ContainerCodePackagePropertiesOutput) Labels() ContainerLabelArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []ContainerLabel { return v.Labels }).(ContainerLabelArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesOutput) Resources() ResourceRequirementsOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) ResourceRequirements { return v.Resources }).(ResourceRequirementsOutput)
}

func (o ContainerCodePackagePropertiesOutput) Settings() SettingArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []Setting { return v.Settings }).(SettingArrayOutput)
}

func (o ContainerCodePackagePropertiesOutput) VolumeRefs() ContainerVolumeArrayOutput {
	return o.ApplyT(func(v ContainerCodePackageProperties) []ContainerVolume { return v.VolumeRefs }).(ContainerVolumeArrayOutput)
}

type ContainerCodePackagePropertiesArrayOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerCodePackageProperties)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesArrayOutput) ToContainerCodePackagePropertiesArrayOutput() ContainerCodePackagePropertiesArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesArrayOutput) ToContainerCodePackagePropertiesArrayOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesArrayOutput) Index(i pulumi.IntInput) ContainerCodePackagePropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerCodePackageProperties {
		return vs[0].([]ContainerCodePackageProperties)[vs[1].(int)]
	}).(ContainerCodePackagePropertiesOutput)
}

type ContainerCodePackagePropertiesResponse struct {
	Commands                []string                         `pulumi:"commands"`
	Diagnostics             *DiagnosticsRefResponse          `pulumi:"diagnostics"`
	Endpoints               []EndpointPropertiesResponse     `pulumi:"endpoints"`
	Entrypoint              *string                          `pulumi:"entrypoint"`
	EnvironmentVariables    []EnvironmentVariableResponse    `pulumi:"environmentVariables"`
	Image                   string                           `pulumi:"image"`
	ImageRegistryCredential *ImageRegistryCredentialResponse `pulumi:"imageRegistryCredential"`
	InstanceView            ContainerInstanceViewResponse    `pulumi:"instanceView"`
	Labels                  []ContainerLabelResponse         `pulumi:"labels"`
	Name                    string                           `pulumi:"name"`
	Resources               ResourceRequirementsResponse     `pulumi:"resources"`
	Settings                []SettingResponse                `pulumi:"settings"`
	VolumeRefs              []ContainerVolumeResponse        `pulumi:"volumeRefs"`
}

type ContainerCodePackagePropertiesResponseOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerCodePackagePropertiesResponse)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesResponseOutput) ToContainerCodePackagePropertiesResponseOutput() ContainerCodePackagePropertiesResponseOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseOutput) ToContainerCodePackagePropertiesResponseOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesResponseOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseOutput) Commands() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []string { return v.Commands }).(pulumi.StringArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Diagnostics() DiagnosticsRefResponsePtrOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) *DiagnosticsRefResponse { return v.Diagnostics }).(DiagnosticsRefResponsePtrOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Endpoints() EndpointPropertiesResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []EndpointPropertiesResponse { return v.Endpoints }).(EndpointPropertiesResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Entrypoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) *string { return v.Entrypoint }).(pulumi.StringPtrOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) EnvironmentVariables() EnvironmentVariableResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []EnvironmentVariableResponse {
		return v.EnvironmentVariables
	}).(EnvironmentVariableResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Image() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) string { return v.Image }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) ImageRegistryCredential() ImageRegistryCredentialResponsePtrOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) *ImageRegistryCredentialResponse {
		return v.ImageRegistryCredential
	}).(ImageRegistryCredentialResponsePtrOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) InstanceView() ContainerInstanceViewResponseOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) ContainerInstanceViewResponse { return v.InstanceView }).(ContainerInstanceViewResponseOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Labels() ContainerLabelResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []ContainerLabelResponse { return v.Labels }).(ContainerLabelResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Resources() ResourceRequirementsResponseOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) ResourceRequirementsResponse { return v.Resources }).(ResourceRequirementsResponseOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) Settings() SettingResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []SettingResponse { return v.Settings }).(SettingResponseArrayOutput)
}

func (o ContainerCodePackagePropertiesResponseOutput) VolumeRefs() ContainerVolumeResponseArrayOutput {
	return o.ApplyT(func(v ContainerCodePackagePropertiesResponse) []ContainerVolumeResponse { return v.VolumeRefs }).(ContainerVolumeResponseArrayOutput)
}

type ContainerCodePackagePropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerCodePackagePropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerCodePackagePropertiesResponse)(nil)).Elem()
}

func (o ContainerCodePackagePropertiesResponseArrayOutput) ToContainerCodePackagePropertiesResponseArrayOutput() ContainerCodePackagePropertiesResponseArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseArrayOutput) ToContainerCodePackagePropertiesResponseArrayOutputWithContext(ctx context.Context) ContainerCodePackagePropertiesResponseArrayOutput {
	return o
}

func (o ContainerCodePackagePropertiesResponseArrayOutput) Index(i pulumi.IntInput) ContainerCodePackagePropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerCodePackagePropertiesResponse {
		return vs[0].([]ContainerCodePackagePropertiesResponse)[vs[1].(int)]
	}).(ContainerCodePackagePropertiesResponseOutput)
}

type ContainerEventResponse struct {
	Count          *int    `pulumi:"count"`
	FirstTimestamp *string `pulumi:"firstTimestamp"`
	LastTimestamp  *string `pulumi:"lastTimestamp"`
	Message        *string `pulumi:"message"`
	Name           *string `pulumi:"name"`
	Type           *string `pulumi:"type"`
}

type ContainerEventResponseOutput struct{ *pulumi.OutputState }

func (ContainerEventResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerEventResponse)(nil)).Elem()
}

func (o ContainerEventResponseOutput) ToContainerEventResponseOutput() ContainerEventResponseOutput {
	return o
}

func (o ContainerEventResponseOutput) ToContainerEventResponseOutputWithContext(ctx context.Context) ContainerEventResponseOutput {
	return o
}

func (o ContainerEventResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o ContainerEventResponseOutput) FirstTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.FirstTimestamp }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) LastTimestamp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.LastTimestamp }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ContainerEventResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerEventResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ContainerEventResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerEventResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerEventResponse)(nil)).Elem()
}

func (o ContainerEventResponseArrayOutput) ToContainerEventResponseArrayOutput() ContainerEventResponseArrayOutput {
	return o
}

func (o ContainerEventResponseArrayOutput) ToContainerEventResponseArrayOutputWithContext(ctx context.Context) ContainerEventResponseArrayOutput {
	return o
}

func (o ContainerEventResponseArrayOutput) Index(i pulumi.IntInput) ContainerEventResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerEventResponse {
		return vs[0].([]ContainerEventResponse)[vs[1].(int)]
	}).(ContainerEventResponseOutput)
}

type ContainerInstanceViewResponse struct {
	CurrentState  *ContainerStateResponse  `pulumi:"currentState"`
	Events        []ContainerEventResponse `pulumi:"events"`
	PreviousState *ContainerStateResponse  `pulumi:"previousState"`
	RestartCount  *int                     `pulumi:"restartCount"`
}

type ContainerInstanceViewResponseOutput struct{ *pulumi.OutputState }

func (ContainerInstanceViewResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerInstanceViewResponse)(nil)).Elem()
}

func (o ContainerInstanceViewResponseOutput) ToContainerInstanceViewResponseOutput() ContainerInstanceViewResponseOutput {
	return o
}

func (o ContainerInstanceViewResponseOutput) ToContainerInstanceViewResponseOutputWithContext(ctx context.Context) ContainerInstanceViewResponseOutput {
	return o
}

func (o ContainerInstanceViewResponseOutput) CurrentState() ContainerStateResponsePtrOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) *ContainerStateResponse { return v.CurrentState }).(ContainerStateResponsePtrOutput)
}

func (o ContainerInstanceViewResponseOutput) Events() ContainerEventResponseArrayOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) []ContainerEventResponse { return v.Events }).(ContainerEventResponseArrayOutput)
}

func (o ContainerInstanceViewResponseOutput) PreviousState() ContainerStateResponsePtrOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) *ContainerStateResponse { return v.PreviousState }).(ContainerStateResponsePtrOutput)
}

func (o ContainerInstanceViewResponseOutput) RestartCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ContainerInstanceViewResponse) *int { return v.RestartCount }).(pulumi.IntPtrOutput)
}

type ContainerLabel struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}





type ContainerLabelInput interface {
	pulumi.Input

	ToContainerLabelOutput() ContainerLabelOutput
	ToContainerLabelOutputWithContext(context.Context) ContainerLabelOutput
}

type ContainerLabelArgs struct {
	Name  pulumi.StringInput `pulumi:"name"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ContainerLabelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerLabel)(nil)).Elem()
}

func (i ContainerLabelArgs) ToContainerLabelOutput() ContainerLabelOutput {
	return i.ToContainerLabelOutputWithContext(context.Background())
}

func (i ContainerLabelArgs) ToContainerLabelOutputWithContext(ctx context.Context) ContainerLabelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerLabelOutput)
}





type ContainerLabelArrayInput interface {
	pulumi.Input

	ToContainerLabelArrayOutput() ContainerLabelArrayOutput
	ToContainerLabelArrayOutputWithContext(context.Context) ContainerLabelArrayOutput
}

type ContainerLabelArray []ContainerLabelInput

func (ContainerLabelArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerLabel)(nil)).Elem()
}

func (i ContainerLabelArray) ToContainerLabelArrayOutput() ContainerLabelArrayOutput {
	return i.ToContainerLabelArrayOutputWithContext(context.Background())
}

func (i ContainerLabelArray) ToContainerLabelArrayOutputWithContext(ctx context.Context) ContainerLabelArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerLabelArrayOutput)
}

type ContainerLabelOutput struct{ *pulumi.OutputState }

func (ContainerLabelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerLabel)(nil)).Elem()
}

func (o ContainerLabelOutput) ToContainerLabelOutput() ContainerLabelOutput {
	return o
}

func (o ContainerLabelOutput) ToContainerLabelOutputWithContext(ctx context.Context) ContainerLabelOutput {
	return o
}

func (o ContainerLabelOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabel) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerLabelOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabel) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerLabelArrayOutput struct{ *pulumi.OutputState }

func (ContainerLabelArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerLabel)(nil)).Elem()
}

func (o ContainerLabelArrayOutput) ToContainerLabelArrayOutput() ContainerLabelArrayOutput {
	return o
}

func (o ContainerLabelArrayOutput) ToContainerLabelArrayOutputWithContext(ctx context.Context) ContainerLabelArrayOutput {
	return o
}

func (o ContainerLabelArrayOutput) Index(i pulumi.IntInput) ContainerLabelOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerLabel {
		return vs[0].([]ContainerLabel)[vs[1].(int)]
	}).(ContainerLabelOutput)
}

type ContainerLabelResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ContainerLabelResponseOutput struct{ *pulumi.OutputState }

func (ContainerLabelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerLabelResponse)(nil)).Elem()
}

func (o ContainerLabelResponseOutput) ToContainerLabelResponseOutput() ContainerLabelResponseOutput {
	return o
}

func (o ContainerLabelResponseOutput) ToContainerLabelResponseOutputWithContext(ctx context.Context) ContainerLabelResponseOutput {
	return o
}

func (o ContainerLabelResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabelResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerLabelResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerLabelResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerLabelResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerLabelResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerLabelResponse)(nil)).Elem()
}

func (o ContainerLabelResponseArrayOutput) ToContainerLabelResponseArrayOutput() ContainerLabelResponseArrayOutput {
	return o
}

func (o ContainerLabelResponseArrayOutput) ToContainerLabelResponseArrayOutputWithContext(ctx context.Context) ContainerLabelResponseArrayOutput {
	return o
}

func (o ContainerLabelResponseArrayOutput) Index(i pulumi.IntInput) ContainerLabelResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerLabelResponse {
		return vs[0].([]ContainerLabelResponse)[vs[1].(int)]
	}).(ContainerLabelResponseOutput)
}

type ContainerStateResponse struct {
	DetailStatus *string `pulumi:"detailStatus"`
	ExitCode     *string `pulumi:"exitCode"`
	FinishTime   *string `pulumi:"finishTime"`
	StartTime    *string `pulumi:"startTime"`
	State        *string `pulumi:"state"`
}

type ContainerStateResponseOutput struct{ *pulumi.OutputState }

func (ContainerStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerStateResponse)(nil)).Elem()
}

func (o ContainerStateResponseOutput) ToContainerStateResponseOutput() ContainerStateResponseOutput {
	return o
}

func (o ContainerStateResponseOutput) ToContainerStateResponseOutputWithContext(ctx context.Context) ContainerStateResponseOutput {
	return o
}

func (o ContainerStateResponseOutput) DetailStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.DetailStatus }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) ExitCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.ExitCode }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.FinishTime }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerStateResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type ContainerStateResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerStateResponse)(nil)).Elem()
}

func (o ContainerStateResponsePtrOutput) ToContainerStateResponsePtrOutput() ContainerStateResponsePtrOutput {
	return o
}

func (o ContainerStateResponsePtrOutput) ToContainerStateResponsePtrOutputWithContext(ctx context.Context) ContainerStateResponsePtrOutput {
	return o
}

func (o ContainerStateResponsePtrOutput) Elem() ContainerStateResponseOutput {
	return o.ApplyT(func(v *ContainerStateResponse) ContainerStateResponse {
		if v != nil {
			return *v
		}
		var ret ContainerStateResponse
		return ret
	}).(ContainerStateResponseOutput)
}

func (o ContainerStateResponsePtrOutput) DetailStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.DetailStatus
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) ExitCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExitCode
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) FinishTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.FinishTime
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

func (o ContainerStateResponsePtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type ContainerVolume struct {
	DestinationPath string `pulumi:"destinationPath"`
	Name            string `pulumi:"name"`
	ReadOnly        *bool  `pulumi:"readOnly"`
}





type ContainerVolumeInput interface {
	pulumi.Input

	ToContainerVolumeOutput() ContainerVolumeOutput
	ToContainerVolumeOutputWithContext(context.Context) ContainerVolumeOutput
}

type ContainerVolumeArgs struct {
	DestinationPath pulumi.StringInput  `pulumi:"destinationPath"`
	Name            pulumi.StringInput  `pulumi:"name"`
	ReadOnly        pulumi.BoolPtrInput `pulumi:"readOnly"`
}

func (ContainerVolumeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerVolume)(nil)).Elem()
}

func (i ContainerVolumeArgs) ToContainerVolumeOutput() ContainerVolumeOutput {
	return i.ToContainerVolumeOutputWithContext(context.Background())
}

func (i ContainerVolumeArgs) ToContainerVolumeOutputWithContext(ctx context.Context) ContainerVolumeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerVolumeOutput)
}





type ContainerVolumeArrayInput interface {
	pulumi.Input

	ToContainerVolumeArrayOutput() ContainerVolumeArrayOutput
	ToContainerVolumeArrayOutputWithContext(context.Context) ContainerVolumeArrayOutput
}

type ContainerVolumeArray []ContainerVolumeInput

func (ContainerVolumeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerVolume)(nil)).Elem()
}

func (i ContainerVolumeArray) ToContainerVolumeArrayOutput() ContainerVolumeArrayOutput {
	return i.ToContainerVolumeArrayOutputWithContext(context.Background())
}

func (i ContainerVolumeArray) ToContainerVolumeArrayOutputWithContext(ctx context.Context) ContainerVolumeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerVolumeArrayOutput)
}

type ContainerVolumeOutput struct{ *pulumi.OutputState }

func (ContainerVolumeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerVolume)(nil)).Elem()
}

func (o ContainerVolumeOutput) ToContainerVolumeOutput() ContainerVolumeOutput {
	return o
}

func (o ContainerVolumeOutput) ToContainerVolumeOutputWithContext(ctx context.Context) ContainerVolumeOutput {
	return o
}

func (o ContainerVolumeOutput) DestinationPath() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerVolume) string { return v.DestinationPath }).(pulumi.StringOutput)
}

func (o ContainerVolumeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerVolume) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerVolumeOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ContainerVolume) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type ContainerVolumeArrayOutput struct{ *pulumi.OutputState }

func (ContainerVolumeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerVolume)(nil)).Elem()
}

func (o ContainerVolumeArrayOutput) ToContainerVolumeArrayOutput() ContainerVolumeArrayOutput {
	return o
}

func (o ContainerVolumeArrayOutput) ToContainerVolumeArrayOutputWithContext(ctx context.Context) ContainerVolumeArrayOutput {
	return o
}

func (o ContainerVolumeArrayOutput) Index(i pulumi.IntInput) ContainerVolumeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerVolume {
		return vs[0].([]ContainerVolume)[vs[1].(int)]
	}).(ContainerVolumeOutput)
}

type ContainerVolumeResponse struct {
	DestinationPath string `pulumi:"destinationPath"`
	Name            string `pulumi:"name"`
	ReadOnly        *bool  `pulumi:"readOnly"`
}

type ContainerVolumeResponseOutput struct{ *pulumi.OutputState }

func (ContainerVolumeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerVolumeResponse)(nil)).Elem()
}

func (o ContainerVolumeResponseOutput) ToContainerVolumeResponseOutput() ContainerVolumeResponseOutput {
	return o
}

func (o ContainerVolumeResponseOutput) ToContainerVolumeResponseOutputWithContext(ctx context.Context) ContainerVolumeResponseOutput {
	return o
}

func (o ContainerVolumeResponseOutput) DestinationPath() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerVolumeResponse) string { return v.DestinationPath }).(pulumi.StringOutput)
}

func (o ContainerVolumeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerVolumeResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerVolumeResponseOutput) ReadOnly() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ContainerVolumeResponse) *bool { return v.ReadOnly }).(pulumi.BoolPtrOutput)
}

type ContainerVolumeResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerVolumeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerVolumeResponse)(nil)).Elem()
}

func (o ContainerVolumeResponseArrayOutput) ToContainerVolumeResponseArrayOutput() ContainerVolumeResponseArrayOutput {
	return o
}

func (o ContainerVolumeResponseArrayOutput) ToContainerVolumeResponseArrayOutputWithContext(ctx context.Context) ContainerVolumeResponseArrayOutput {
	return o
}

func (o ContainerVolumeResponseArrayOutput) Index(i pulumi.IntInput) ContainerVolumeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerVolumeResponse {
		return vs[0].([]ContainerVolumeResponse)[vs[1].(int)]
	}).(ContainerVolumeResponseOutput)
}

type DiagnosticsDescription struct {
	DefaultSinkRefs []string                                         `pulumi:"defaultSinkRefs"`
	Enabled         *bool                                            `pulumi:"enabled"`
	Sinks           []AzureInternalMonitoringPipelineSinkDescription `pulumi:"sinks"`
}





type DiagnosticsDescriptionInput interface {
	pulumi.Input

	ToDiagnosticsDescriptionOutput() DiagnosticsDescriptionOutput
	ToDiagnosticsDescriptionOutputWithContext(context.Context) DiagnosticsDescriptionOutput
}

type DiagnosticsDescriptionArgs struct {
	DefaultSinkRefs pulumi.StringArrayInput                                  `pulumi:"defaultSinkRefs"`
	Enabled         pulumi.BoolPtrInput                                      `pulumi:"enabled"`
	Sinks           AzureInternalMonitoringPipelineSinkDescriptionArrayInput `pulumi:"sinks"`
}

func (DiagnosticsDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsDescription)(nil)).Elem()
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionOutput() DiagnosticsDescriptionOutput {
	return i.ToDiagnosticsDescriptionOutputWithContext(context.Background())
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionOutputWithContext(ctx context.Context) DiagnosticsDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsDescriptionOutput)
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return i.ToDiagnosticsDescriptionPtrOutputWithContext(context.Background())
}

func (i DiagnosticsDescriptionArgs) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsDescriptionOutput).ToDiagnosticsDescriptionPtrOutputWithContext(ctx)
}









type DiagnosticsDescriptionPtrInput interface {
	pulumi.Input

	ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput
	ToDiagnosticsDescriptionPtrOutputWithContext(context.Context) DiagnosticsDescriptionPtrOutput
}

type diagnosticsDescriptionPtrType DiagnosticsDescriptionArgs

func DiagnosticsDescriptionPtr(v *DiagnosticsDescriptionArgs) DiagnosticsDescriptionPtrInput {
	return (*diagnosticsDescriptionPtrType)(v)
}

func (*diagnosticsDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsDescription)(nil)).Elem()
}

func (i *diagnosticsDescriptionPtrType) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return i.ToDiagnosticsDescriptionPtrOutputWithContext(context.Background())
}

func (i *diagnosticsDescriptionPtrType) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsDescriptionPtrOutput)
}

type DiagnosticsDescriptionOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsDescription)(nil)).Elem()
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionOutput() DiagnosticsDescriptionOutput {
	return o
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionOutputWithContext(ctx context.Context) DiagnosticsDescriptionOutput {
	return o
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return o.ToDiagnosticsDescriptionPtrOutputWithContext(context.Background())
}

func (o DiagnosticsDescriptionOutput) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsDescription) *DiagnosticsDescription {
		return &v
	}).(DiagnosticsDescriptionPtrOutput)
}

func (o DiagnosticsDescriptionOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescription) []string { return v.DefaultSinkRefs }).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsDescription) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescription) []AzureInternalMonitoringPipelineSinkDescription { return v.Sinks }).(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput)
}

type DiagnosticsDescriptionPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsDescription)(nil)).Elem()
}

func (o DiagnosticsDescriptionPtrOutput) ToDiagnosticsDescriptionPtrOutput() DiagnosticsDescriptionPtrOutput {
	return o
}

func (o DiagnosticsDescriptionPtrOutput) ToDiagnosticsDescriptionPtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionPtrOutput {
	return o
}

func (o DiagnosticsDescriptionPtrOutput) Elem() DiagnosticsDescriptionOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) DiagnosticsDescription {
		if v != nil {
			return *v
		}
		var ret DiagnosticsDescription
		return ret
	}).(DiagnosticsDescriptionOutput)
}

func (o DiagnosticsDescriptionPtrOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) []string {
		if v == nil {
			return nil
		}
		return v.DefaultSinkRefs
	}).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionPtrOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescription) []AzureInternalMonitoringPipelineSinkDescription {
		if v == nil {
			return nil
		}
		return v.Sinks
	}).(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput)
}

type DiagnosticsDescriptionResponse struct {
	DefaultSinkRefs []string                                                 `pulumi:"defaultSinkRefs"`
	Enabled         *bool                                                    `pulumi:"enabled"`
	Sinks           []AzureInternalMonitoringPipelineSinkDescriptionResponse `pulumi:"sinks"`
}

type DiagnosticsDescriptionResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsDescriptionResponse)(nil)).Elem()
}

func (o DiagnosticsDescriptionResponseOutput) ToDiagnosticsDescriptionResponseOutput() DiagnosticsDescriptionResponseOutput {
	return o
}

func (o DiagnosticsDescriptionResponseOutput) ToDiagnosticsDescriptionResponseOutputWithContext(ctx context.Context) DiagnosticsDescriptionResponseOutput {
	return o
}

func (o DiagnosticsDescriptionResponseOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescriptionResponse) []string { return v.DefaultSinkRefs }).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsDescriptionResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionResponseOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o.ApplyT(func(v DiagnosticsDescriptionResponse) []AzureInternalMonitoringPipelineSinkDescriptionResponse {
		return v.Sinks
	}).(AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput)
}

type DiagnosticsDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsDescriptionResponse)(nil)).Elem()
}

func (o DiagnosticsDescriptionResponsePtrOutput) ToDiagnosticsDescriptionResponsePtrOutput() DiagnosticsDescriptionResponsePtrOutput {
	return o
}

func (o DiagnosticsDescriptionResponsePtrOutput) ToDiagnosticsDescriptionResponsePtrOutputWithContext(ctx context.Context) DiagnosticsDescriptionResponsePtrOutput {
	return o
}

func (o DiagnosticsDescriptionResponsePtrOutput) Elem() DiagnosticsDescriptionResponseOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) DiagnosticsDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsDescriptionResponse
		return ret
	}).(DiagnosticsDescriptionResponseOutput)
}

func (o DiagnosticsDescriptionResponsePtrOutput) DefaultSinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.DefaultSinkRefs
	}).(pulumi.StringArrayOutput)
}

func (o DiagnosticsDescriptionResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsDescriptionResponsePtrOutput) Sinks() AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput {
	return o.ApplyT(func(v *DiagnosticsDescriptionResponse) []AzureInternalMonitoringPipelineSinkDescriptionResponse {
		if v == nil {
			return nil
		}
		return v.Sinks
	}).(AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput)
}

type DiagnosticsRef struct {
	Enabled  *bool    `pulumi:"enabled"`
	SinkRefs []string `pulumi:"sinkRefs"`
}





type DiagnosticsRefInput interface {
	pulumi.Input

	ToDiagnosticsRefOutput() DiagnosticsRefOutput
	ToDiagnosticsRefOutputWithContext(context.Context) DiagnosticsRefOutput
}

type DiagnosticsRefArgs struct {
	Enabled  pulumi.BoolPtrInput     `pulumi:"enabled"`
	SinkRefs pulumi.StringArrayInput `pulumi:"sinkRefs"`
}

func (DiagnosticsRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsRef)(nil)).Elem()
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefOutput() DiagnosticsRefOutput {
	return i.ToDiagnosticsRefOutputWithContext(context.Background())
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefOutputWithContext(ctx context.Context) DiagnosticsRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsRefOutput)
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return i.ToDiagnosticsRefPtrOutputWithContext(context.Background())
}

func (i DiagnosticsRefArgs) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsRefOutput).ToDiagnosticsRefPtrOutputWithContext(ctx)
}









type DiagnosticsRefPtrInput interface {
	pulumi.Input

	ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput
	ToDiagnosticsRefPtrOutputWithContext(context.Context) DiagnosticsRefPtrOutput
}

type diagnosticsRefPtrType DiagnosticsRefArgs

func DiagnosticsRefPtr(v *DiagnosticsRefArgs) DiagnosticsRefPtrInput {
	return (*diagnosticsRefPtrType)(v)
}

func (*diagnosticsRefPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsRef)(nil)).Elem()
}

func (i *diagnosticsRefPtrType) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return i.ToDiagnosticsRefPtrOutputWithContext(context.Background())
}

func (i *diagnosticsRefPtrType) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DiagnosticsRefPtrOutput)
}

type DiagnosticsRefOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsRef)(nil)).Elem()
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefOutput() DiagnosticsRefOutput {
	return o
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefOutputWithContext(ctx context.Context) DiagnosticsRefOutput {
	return o
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return o.ToDiagnosticsRefPtrOutputWithContext(context.Background())
}

func (o DiagnosticsRefOutput) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DiagnosticsRef) *DiagnosticsRef {
		return &v
	}).(DiagnosticsRefPtrOutput)
}

func (o DiagnosticsRefOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsRef) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsRef) []string { return v.SinkRefs }).(pulumi.StringArrayOutput)
}

type DiagnosticsRefPtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsRef)(nil)).Elem()
}

func (o DiagnosticsRefPtrOutput) ToDiagnosticsRefPtrOutput() DiagnosticsRefPtrOutput {
	return o
}

func (o DiagnosticsRefPtrOutput) ToDiagnosticsRefPtrOutputWithContext(ctx context.Context) DiagnosticsRefPtrOutput {
	return o
}

func (o DiagnosticsRefPtrOutput) Elem() DiagnosticsRefOutput {
	return o.ApplyT(func(v *DiagnosticsRef) DiagnosticsRef {
		if v != nil {
			return *v
		}
		var ret DiagnosticsRef
		return ret
	}).(DiagnosticsRefOutput)
}

func (o DiagnosticsRefPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsRef) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefPtrOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsRef) []string {
		if v == nil {
			return nil
		}
		return v.SinkRefs
	}).(pulumi.StringArrayOutput)
}

type DiagnosticsRefResponse struct {
	Enabled  *bool    `pulumi:"enabled"`
	SinkRefs []string `pulumi:"sinkRefs"`
}

type DiagnosticsRefResponseOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DiagnosticsRefResponse)(nil)).Elem()
}

func (o DiagnosticsRefResponseOutput) ToDiagnosticsRefResponseOutput() DiagnosticsRefResponseOutput {
	return o
}

func (o DiagnosticsRefResponseOutput) ToDiagnosticsRefResponseOutputWithContext(ctx context.Context) DiagnosticsRefResponseOutput {
	return o
}

func (o DiagnosticsRefResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DiagnosticsRefResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefResponseOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DiagnosticsRefResponse) []string { return v.SinkRefs }).(pulumi.StringArrayOutput)
}

type DiagnosticsRefResponsePtrOutput struct{ *pulumi.OutputState }

func (DiagnosticsRefResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DiagnosticsRefResponse)(nil)).Elem()
}

func (o DiagnosticsRefResponsePtrOutput) ToDiagnosticsRefResponsePtrOutput() DiagnosticsRefResponsePtrOutput {
	return o
}

func (o DiagnosticsRefResponsePtrOutput) ToDiagnosticsRefResponsePtrOutputWithContext(ctx context.Context) DiagnosticsRefResponsePtrOutput {
	return o
}

func (o DiagnosticsRefResponsePtrOutput) Elem() DiagnosticsRefResponseOutput {
	return o.ApplyT(func(v *DiagnosticsRefResponse) DiagnosticsRefResponse {
		if v != nil {
			return *v
		}
		var ret DiagnosticsRefResponse
		return ret
	}).(DiagnosticsRefResponseOutput)
}

func (o DiagnosticsRefResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DiagnosticsRefResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o DiagnosticsRefResponsePtrOutput) SinkRefs() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *DiagnosticsRefResponse) []string {
		if v == nil {
			return nil
		}
		return v.SinkRefs
	}).(pulumi.StringArrayOutput)
}

type EndpointProperties struct {
	Name string `pulumi:"name"`
	Port *int   `pulumi:"port"`
}





type EndpointPropertiesInput interface {
	pulumi.Input

	ToEndpointPropertiesOutput() EndpointPropertiesOutput
	ToEndpointPropertiesOutputWithContext(context.Context) EndpointPropertiesOutput
}

type EndpointPropertiesArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Port pulumi.IntPtrInput `pulumi:"port"`
}

func (EndpointPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointProperties)(nil)).Elem()
}

func (i EndpointPropertiesArgs) ToEndpointPropertiesOutput() EndpointPropertiesOutput {
	return i.ToEndpointPropertiesOutputWithContext(context.Background())
}

func (i EndpointPropertiesArgs) ToEndpointPropertiesOutputWithContext(ctx context.Context) EndpointPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesOutput)
}





type EndpointPropertiesArrayInput interface {
	pulumi.Input

	ToEndpointPropertiesArrayOutput() EndpointPropertiesArrayOutput
	ToEndpointPropertiesArrayOutputWithContext(context.Context) EndpointPropertiesArrayOutput
}

type EndpointPropertiesArray []EndpointPropertiesInput

func (EndpointPropertiesArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointProperties)(nil)).Elem()
}

func (i EndpointPropertiesArray) ToEndpointPropertiesArrayOutput() EndpointPropertiesArrayOutput {
	return i.ToEndpointPropertiesArrayOutputWithContext(context.Background())
}

func (i EndpointPropertiesArray) ToEndpointPropertiesArrayOutputWithContext(ctx context.Context) EndpointPropertiesArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EndpointPropertiesArrayOutput)
}

type EndpointPropertiesOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointProperties)(nil)).Elem()
}

func (o EndpointPropertiesOutput) ToEndpointPropertiesOutput() EndpointPropertiesOutput {
	return o
}

func (o EndpointPropertiesOutput) ToEndpointPropertiesOutputWithContext(ctx context.Context) EndpointPropertiesOutput {
	return o
}

func (o EndpointPropertiesOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointProperties) string { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointPropertiesOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointProperties) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type EndpointPropertiesArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointProperties)(nil)).Elem()
}

func (o EndpointPropertiesArrayOutput) ToEndpointPropertiesArrayOutput() EndpointPropertiesArrayOutput {
	return o
}

func (o EndpointPropertiesArrayOutput) ToEndpointPropertiesArrayOutputWithContext(ctx context.Context) EndpointPropertiesArrayOutput {
	return o
}

func (o EndpointPropertiesArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointProperties {
		return vs[0].([]EndpointProperties)[vs[1].(int)]
	}).(EndpointPropertiesOutput)
}

type EndpointPropertiesResponse struct {
	Name string `pulumi:"name"`
	Port *int   `pulumi:"port"`
}

type EndpointPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EndpointPropertiesResponse)(nil)).Elem()
}

func (o EndpointPropertiesResponseOutput) ToEndpointPropertiesResponseOutput() EndpointPropertiesResponseOutput {
	return o
}

func (o EndpointPropertiesResponseOutput) ToEndpointPropertiesResponseOutputWithContext(ctx context.Context) EndpointPropertiesResponseOutput {
	return o
}

func (o EndpointPropertiesResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EndpointPropertiesResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EndpointPropertiesResponseOutput) Port() pulumi.IntPtrOutput {
	return o.ApplyT(func(v EndpointPropertiesResponse) *int { return v.Port }).(pulumi.IntPtrOutput)
}

type EndpointPropertiesResponseArrayOutput struct{ *pulumi.OutputState }

func (EndpointPropertiesResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EndpointPropertiesResponse)(nil)).Elem()
}

func (o EndpointPropertiesResponseArrayOutput) ToEndpointPropertiesResponseArrayOutput() EndpointPropertiesResponseArrayOutput {
	return o
}

func (o EndpointPropertiesResponseArrayOutput) ToEndpointPropertiesResponseArrayOutputWithContext(ctx context.Context) EndpointPropertiesResponseArrayOutput {
	return o
}

func (o EndpointPropertiesResponseArrayOutput) Index(i pulumi.IntInput) EndpointPropertiesResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EndpointPropertiesResponse {
		return vs[0].([]EndpointPropertiesResponse)[vs[1].(int)]
	}).(EndpointPropertiesResponseOutput)
}

type EnvironmentVariable struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type EnvironmentVariableInput interface {
	pulumi.Input

	ToEnvironmentVariableOutput() EnvironmentVariableOutput
	ToEnvironmentVariableOutputWithContext(context.Context) EnvironmentVariableOutput
}

type EnvironmentVariableArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (EnvironmentVariableArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArgs) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return i.ToEnvironmentVariableOutputWithContext(context.Background())
}

func (i EnvironmentVariableArgs) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableOutput)
}





type EnvironmentVariableArrayInput interface {
	pulumi.Input

	ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput
	ToEnvironmentVariableArrayOutputWithContext(context.Context) EnvironmentVariableArrayOutput
}

type EnvironmentVariableArray []EnvironmentVariableInput

func (EnvironmentVariableArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariable)(nil)).Elem()
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return i.ToEnvironmentVariableArrayOutputWithContext(context.Background())
}

func (i EnvironmentVariableArray) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnvironmentVariableArrayOutput)
}

type EnvironmentVariableOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutput() EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) ToEnvironmentVariableOutputWithContext(ctx context.Context) EnvironmentVariableOutput {
	return o
}

func (o EnvironmentVariableOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariable) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVariableOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariable) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVariableArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariable)(nil)).Elem()
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutput() EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) ToEnvironmentVariableArrayOutputWithContext(ctx context.Context) EnvironmentVariableArrayOutput {
	return o
}

func (o EnvironmentVariableArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVariable {
		return vs[0].([]EnvironmentVariable)[vs[1].(int)]
	}).(EnvironmentVariableOutput)
}

type EnvironmentVariableResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type EnvironmentVariableResponseOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnvironmentVariableResponse)(nil)).Elem()
}

func (o EnvironmentVariableResponseOutput) ToEnvironmentVariableResponseOutput() EnvironmentVariableResponseOutput {
	return o
}

func (o EnvironmentVariableResponseOutput) ToEnvironmentVariableResponseOutputWithContext(ctx context.Context) EnvironmentVariableResponseOutput {
	return o
}

func (o EnvironmentVariableResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o EnvironmentVariableResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnvironmentVariableResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type EnvironmentVariableResponseArrayOutput struct{ *pulumi.OutputState }

func (EnvironmentVariableResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnvironmentVariableResponse)(nil)).Elem()
}

func (o EnvironmentVariableResponseArrayOutput) ToEnvironmentVariableResponseArrayOutput() EnvironmentVariableResponseArrayOutput {
	return o
}

func (o EnvironmentVariableResponseArrayOutput) ToEnvironmentVariableResponseArrayOutputWithContext(ctx context.Context) EnvironmentVariableResponseArrayOutput {
	return o
}

func (o EnvironmentVariableResponseArrayOutput) Index(i pulumi.IntInput) EnvironmentVariableResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnvironmentVariableResponse {
		return vs[0].([]EnvironmentVariableResponse)[vs[1].(int)]
	}).(EnvironmentVariableResponseOutput)
}

type ImageRegistryCredential struct {
	Password *string `pulumi:"password"`
	Server   string  `pulumi:"server"`
	Username string  `pulumi:"username"`
}





type ImageRegistryCredentialInput interface {
	pulumi.Input

	ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput
	ToImageRegistryCredentialOutputWithContext(context.Context) ImageRegistryCredentialOutput
}

type ImageRegistryCredentialArgs struct {
	Password pulumi.StringPtrInput `pulumi:"password"`
	Server   pulumi.StringInput    `pulumi:"server"`
	Username pulumi.StringInput    `pulumi:"username"`
}

func (ImageRegistryCredentialArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return i.ToImageRegistryCredentialOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput)
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return i.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (i ImageRegistryCredentialArgs) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialOutput).ToImageRegistryCredentialPtrOutputWithContext(ctx)
}









type ImageRegistryCredentialPtrInput interface {
	pulumi.Input

	ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput
	ToImageRegistryCredentialPtrOutputWithContext(context.Context) ImageRegistryCredentialPtrOutput
}

type imageRegistryCredentialPtrType ImageRegistryCredentialArgs

func ImageRegistryCredentialPtr(v *ImageRegistryCredentialArgs) ImageRegistryCredentialPtrInput {
	return (*imageRegistryCredentialPtrType)(v)
}

func (*imageRegistryCredentialPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredential)(nil)).Elem()
}

func (i *imageRegistryCredentialPtrType) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return i.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (i *imageRegistryCredentialPtrType) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ImageRegistryCredentialPtrOutput)
}

type ImageRegistryCredentialOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutput() ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialOutputWithContext(ctx context.Context) ImageRegistryCredentialOutput {
	return o
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return o.ToImageRegistryCredentialPtrOutputWithContext(context.Background())
}

func (o ImageRegistryCredentialOutput) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ImageRegistryCredential) *ImageRegistryCredential {
		return &v
	}).(ImageRegistryCredentialPtrOutput)
}

func (o ImageRegistryCredentialOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredential) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredential) string { return v.Server }).(pulumi.StringOutput)
}

func (o ImageRegistryCredentialOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredential) string { return v.Username }).(pulumi.StringOutput)
}

type ImageRegistryCredentialPtrOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredential)(nil)).Elem()
}

func (o ImageRegistryCredentialPtrOutput) ToImageRegistryCredentialPtrOutput() ImageRegistryCredentialPtrOutput {
	return o
}

func (o ImageRegistryCredentialPtrOutput) ToImageRegistryCredentialPtrOutputWithContext(ctx context.Context) ImageRegistryCredentialPtrOutput {
	return o
}

func (o ImageRegistryCredentialPtrOutput) Elem() ImageRegistryCredentialOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) ImageRegistryCredential {
		if v != nil {
			return *v
		}
		var ret ImageRegistryCredential
		return ret
	}).(ImageRegistryCredentialOutput)
}

func (o ImageRegistryCredentialPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialPtrOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return &v.Server
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredential) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type ImageRegistryCredentialResponse struct {
	Password *string `pulumi:"password"`
	Server   string  `pulumi:"server"`
	Username string  `pulumi:"username"`
}

type ImageRegistryCredentialResponseOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutput() ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) ToImageRegistryCredentialResponseOutputWithContext(ctx context.Context) ImageRegistryCredentialResponseOutput {
	return o
}

func (o ImageRegistryCredentialResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponseOutput) Server() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) string { return v.Server }).(pulumi.StringOutput)
}

func (o ImageRegistryCredentialResponseOutput) Username() pulumi.StringOutput {
	return o.ApplyT(func(v ImageRegistryCredentialResponse) string { return v.Username }).(pulumi.StringOutput)
}

type ImageRegistryCredentialResponsePtrOutput struct{ *pulumi.OutputState }

func (ImageRegistryCredentialResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ImageRegistryCredentialResponse)(nil)).Elem()
}

func (o ImageRegistryCredentialResponsePtrOutput) ToImageRegistryCredentialResponsePtrOutput() ImageRegistryCredentialResponsePtrOutput {
	return o
}

func (o ImageRegistryCredentialResponsePtrOutput) ToImageRegistryCredentialResponsePtrOutputWithContext(ctx context.Context) ImageRegistryCredentialResponsePtrOutput {
	return o
}

func (o ImageRegistryCredentialResponsePtrOutput) Elem() ImageRegistryCredentialResponseOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) ImageRegistryCredentialResponse {
		if v != nil {
			return *v
		}
		var ret ImageRegistryCredentialResponse
		return ret
	}).(ImageRegistryCredentialResponseOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Server() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Server
	}).(pulumi.StringPtrOutput)
}

func (o ImageRegistryCredentialResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ImageRegistryCredentialResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Username
	}).(pulumi.StringPtrOutput)
}

type IngressConfig struct {
	Layer4   []Layer4IngressConfig `pulumi:"layer4"`
	QosLevel *string               `pulumi:"qosLevel"`
}





type IngressConfigInput interface {
	pulumi.Input

	ToIngressConfigOutput() IngressConfigOutput
	ToIngressConfigOutputWithContext(context.Context) IngressConfigOutput
}

type IngressConfigArgs struct {
	Layer4   Layer4IngressConfigArrayInput `pulumi:"layer4"`
	QosLevel pulumi.StringPtrInput         `pulumi:"qosLevel"`
}

func (IngressConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressConfig)(nil)).Elem()
}

func (i IngressConfigArgs) ToIngressConfigOutput() IngressConfigOutput {
	return i.ToIngressConfigOutputWithContext(context.Background())
}

func (i IngressConfigArgs) ToIngressConfigOutputWithContext(ctx context.Context) IngressConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressConfigOutput)
}

func (i IngressConfigArgs) ToIngressConfigPtrOutput() IngressConfigPtrOutput {
	return i.ToIngressConfigPtrOutputWithContext(context.Background())
}

func (i IngressConfigArgs) ToIngressConfigPtrOutputWithContext(ctx context.Context) IngressConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressConfigOutput).ToIngressConfigPtrOutputWithContext(ctx)
}









type IngressConfigPtrInput interface {
	pulumi.Input

	ToIngressConfigPtrOutput() IngressConfigPtrOutput
	ToIngressConfigPtrOutputWithContext(context.Context) IngressConfigPtrOutput
}

type ingressConfigPtrType IngressConfigArgs

func IngressConfigPtr(v *IngressConfigArgs) IngressConfigPtrInput {
	return (*ingressConfigPtrType)(v)
}

func (*ingressConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressConfig)(nil)).Elem()
}

func (i *ingressConfigPtrType) ToIngressConfigPtrOutput() IngressConfigPtrOutput {
	return i.ToIngressConfigPtrOutputWithContext(context.Background())
}

func (i *ingressConfigPtrType) ToIngressConfigPtrOutputWithContext(ctx context.Context) IngressConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IngressConfigPtrOutput)
}

type IngressConfigOutput struct{ *pulumi.OutputState }

func (IngressConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressConfig)(nil)).Elem()
}

func (o IngressConfigOutput) ToIngressConfigOutput() IngressConfigOutput {
	return o
}

func (o IngressConfigOutput) ToIngressConfigOutputWithContext(ctx context.Context) IngressConfigOutput {
	return o
}

func (o IngressConfigOutput) ToIngressConfigPtrOutput() IngressConfigPtrOutput {
	return o.ToIngressConfigPtrOutputWithContext(context.Background())
}

func (o IngressConfigOutput) ToIngressConfigPtrOutputWithContext(ctx context.Context) IngressConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v IngressConfig) *IngressConfig {
		return &v
	}).(IngressConfigPtrOutput)
}

func (o IngressConfigOutput) Layer4() Layer4IngressConfigArrayOutput {
	return o.ApplyT(func(v IngressConfig) []Layer4IngressConfig { return v.Layer4 }).(Layer4IngressConfigArrayOutput)
}

func (o IngressConfigOutput) QosLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressConfig) *string { return v.QosLevel }).(pulumi.StringPtrOutput)
}

type IngressConfigPtrOutput struct{ *pulumi.OutputState }

func (IngressConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressConfig)(nil)).Elem()
}

func (o IngressConfigPtrOutput) ToIngressConfigPtrOutput() IngressConfigPtrOutput {
	return o
}

func (o IngressConfigPtrOutput) ToIngressConfigPtrOutputWithContext(ctx context.Context) IngressConfigPtrOutput {
	return o
}

func (o IngressConfigPtrOutput) Elem() IngressConfigOutput {
	return o.ApplyT(func(v *IngressConfig) IngressConfig {
		if v != nil {
			return *v
		}
		var ret IngressConfig
		return ret
	}).(IngressConfigOutput)
}

func (o IngressConfigPtrOutput) Layer4() Layer4IngressConfigArrayOutput {
	return o.ApplyT(func(v *IngressConfig) []Layer4IngressConfig {
		if v == nil {
			return nil
		}
		return v.Layer4
	}).(Layer4IngressConfigArrayOutput)
}

func (o IngressConfigPtrOutput) QosLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressConfig) *string {
		if v == nil {
			return nil
		}
		return v.QosLevel
	}).(pulumi.StringPtrOutput)
}

type IngressConfigResponse struct {
	Layer4          []Layer4IngressConfigResponse `pulumi:"layer4"`
	PublicIPAddress string                        `pulumi:"publicIPAddress"`
	QosLevel        *string                       `pulumi:"qosLevel"`
}

type IngressConfigResponseOutput struct{ *pulumi.OutputState }

func (IngressConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IngressConfigResponse)(nil)).Elem()
}

func (o IngressConfigResponseOutput) ToIngressConfigResponseOutput() IngressConfigResponseOutput {
	return o
}

func (o IngressConfigResponseOutput) ToIngressConfigResponseOutputWithContext(ctx context.Context) IngressConfigResponseOutput {
	return o
}

func (o IngressConfigResponseOutput) Layer4() Layer4IngressConfigResponseArrayOutput {
	return o.ApplyT(func(v IngressConfigResponse) []Layer4IngressConfigResponse { return v.Layer4 }).(Layer4IngressConfigResponseArrayOutput)
}

func (o IngressConfigResponseOutput) PublicIPAddress() pulumi.StringOutput {
	return o.ApplyT(func(v IngressConfigResponse) string { return v.PublicIPAddress }).(pulumi.StringOutput)
}

func (o IngressConfigResponseOutput) QosLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IngressConfigResponse) *string { return v.QosLevel }).(pulumi.StringPtrOutput)
}

type IngressConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (IngressConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**IngressConfigResponse)(nil)).Elem()
}

func (o IngressConfigResponsePtrOutput) ToIngressConfigResponsePtrOutput() IngressConfigResponsePtrOutput {
	return o
}

func (o IngressConfigResponsePtrOutput) ToIngressConfigResponsePtrOutputWithContext(ctx context.Context) IngressConfigResponsePtrOutput {
	return o
}

func (o IngressConfigResponsePtrOutput) Elem() IngressConfigResponseOutput {
	return o.ApplyT(func(v *IngressConfigResponse) IngressConfigResponse {
		if v != nil {
			return *v
		}
		var ret IngressConfigResponse
		return ret
	}).(IngressConfigResponseOutput)
}

func (o IngressConfigResponsePtrOutput) Layer4() Layer4IngressConfigResponseArrayOutput {
	return o.ApplyT(func(v *IngressConfigResponse) []Layer4IngressConfigResponse {
		if v == nil {
			return nil
		}
		return v.Layer4
	}).(Layer4IngressConfigResponseArrayOutput)
}

func (o IngressConfigResponsePtrOutput) PublicIPAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.PublicIPAddress
	}).(pulumi.StringPtrOutput)
}

func (o IngressConfigResponsePtrOutput) QosLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *IngressConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.QosLevel
	}).(pulumi.StringPtrOutput)
}

type Layer4IngressConfig struct {
	ApplicationName *string `pulumi:"applicationName"`
	EndpointName    *string `pulumi:"endpointName"`
	Name            *string `pulumi:"name"`
	PublicPort      *int    `pulumi:"publicPort"`
	ServiceName     *string `pulumi:"serviceName"`
}





type Layer4IngressConfigInput interface {
	pulumi.Input

	ToLayer4IngressConfigOutput() Layer4IngressConfigOutput
	ToLayer4IngressConfigOutputWithContext(context.Context) Layer4IngressConfigOutput
}

type Layer4IngressConfigArgs struct {
	ApplicationName pulumi.StringPtrInput `pulumi:"applicationName"`
	EndpointName    pulumi.StringPtrInput `pulumi:"endpointName"`
	Name            pulumi.StringPtrInput `pulumi:"name"`
	PublicPort      pulumi.IntPtrInput    `pulumi:"publicPort"`
	ServiceName     pulumi.StringPtrInput `pulumi:"serviceName"`
}

func (Layer4IngressConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Layer4IngressConfig)(nil)).Elem()
}

func (i Layer4IngressConfigArgs) ToLayer4IngressConfigOutput() Layer4IngressConfigOutput {
	return i.ToLayer4IngressConfigOutputWithContext(context.Background())
}

func (i Layer4IngressConfigArgs) ToLayer4IngressConfigOutputWithContext(ctx context.Context) Layer4IngressConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4IngressConfigOutput)
}





type Layer4IngressConfigArrayInput interface {
	pulumi.Input

	ToLayer4IngressConfigArrayOutput() Layer4IngressConfigArrayOutput
	ToLayer4IngressConfigArrayOutputWithContext(context.Context) Layer4IngressConfigArrayOutput
}

type Layer4IngressConfigArray []Layer4IngressConfigInput

func (Layer4IngressConfigArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Layer4IngressConfig)(nil)).Elem()
}

func (i Layer4IngressConfigArray) ToLayer4IngressConfigArrayOutput() Layer4IngressConfigArrayOutput {
	return i.ToLayer4IngressConfigArrayOutputWithContext(context.Background())
}

func (i Layer4IngressConfigArray) ToLayer4IngressConfigArrayOutputWithContext(ctx context.Context) Layer4IngressConfigArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(Layer4IngressConfigArrayOutput)
}

type Layer4IngressConfigOutput struct{ *pulumi.OutputState }

func (Layer4IngressConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Layer4IngressConfig)(nil)).Elem()
}

func (o Layer4IngressConfigOutput) ToLayer4IngressConfigOutput() Layer4IngressConfigOutput {
	return o
}

func (o Layer4IngressConfigOutput) ToLayer4IngressConfigOutputWithContext(ctx context.Context) Layer4IngressConfigOutput {
	return o
}

func (o Layer4IngressConfigOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfig) *string { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

func (o Layer4IngressConfigOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfig) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o Layer4IngressConfigOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfig) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o Layer4IngressConfigOutput) PublicPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfig) *int { return v.PublicPort }).(pulumi.IntPtrOutput)
}

func (o Layer4IngressConfigOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfig) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

type Layer4IngressConfigArrayOutput struct{ *pulumi.OutputState }

func (Layer4IngressConfigArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Layer4IngressConfig)(nil)).Elem()
}

func (o Layer4IngressConfigArrayOutput) ToLayer4IngressConfigArrayOutput() Layer4IngressConfigArrayOutput {
	return o
}

func (o Layer4IngressConfigArrayOutput) ToLayer4IngressConfigArrayOutputWithContext(ctx context.Context) Layer4IngressConfigArrayOutput {
	return o
}

func (o Layer4IngressConfigArrayOutput) Index(i pulumi.IntInput) Layer4IngressConfigOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Layer4IngressConfig {
		return vs[0].([]Layer4IngressConfig)[vs[1].(int)]
	}).(Layer4IngressConfigOutput)
}

type Layer4IngressConfigResponse struct {
	ApplicationName *string `pulumi:"applicationName"`
	EndpointName    *string `pulumi:"endpointName"`
	Name            *string `pulumi:"name"`
	PublicPort      *int    `pulumi:"publicPort"`
	ServiceName     *string `pulumi:"serviceName"`
}

type Layer4IngressConfigResponseOutput struct{ *pulumi.OutputState }

func (Layer4IngressConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Layer4IngressConfigResponse)(nil)).Elem()
}

func (o Layer4IngressConfigResponseOutput) ToLayer4IngressConfigResponseOutput() Layer4IngressConfigResponseOutput {
	return o
}

func (o Layer4IngressConfigResponseOutput) ToLayer4IngressConfigResponseOutputWithContext(ctx context.Context) Layer4IngressConfigResponseOutput {
	return o
}

func (o Layer4IngressConfigResponseOutput) ApplicationName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfigResponse) *string { return v.ApplicationName }).(pulumi.StringPtrOutput)
}

func (o Layer4IngressConfigResponseOutput) EndpointName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfigResponse) *string { return v.EndpointName }).(pulumi.StringPtrOutput)
}

func (o Layer4IngressConfigResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfigResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o Layer4IngressConfigResponseOutput) PublicPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfigResponse) *int { return v.PublicPort }).(pulumi.IntPtrOutput)
}

func (o Layer4IngressConfigResponseOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Layer4IngressConfigResponse) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

type Layer4IngressConfigResponseArrayOutput struct{ *pulumi.OutputState }

func (Layer4IngressConfigResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Layer4IngressConfigResponse)(nil)).Elem()
}

func (o Layer4IngressConfigResponseArrayOutput) ToLayer4IngressConfigResponseArrayOutput() Layer4IngressConfigResponseArrayOutput {
	return o
}

func (o Layer4IngressConfigResponseArrayOutput) ToLayer4IngressConfigResponseArrayOutputWithContext(ctx context.Context) Layer4IngressConfigResponseArrayOutput {
	return o
}

func (o Layer4IngressConfigResponseArrayOutput) Index(i pulumi.IntInput) Layer4IngressConfigResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Layer4IngressConfigResponse {
		return vs[0].([]Layer4IngressConfigResponse)[vs[1].(int)]
	}).(Layer4IngressConfigResponseOutput)
}

type NetworkRef struct {
	Name *string `pulumi:"name"`
}





type NetworkRefInput interface {
	pulumi.Input

	ToNetworkRefOutput() NetworkRefOutput
	ToNetworkRefOutputWithContext(context.Context) NetworkRefOutput
}

type NetworkRefArgs struct {
	Name pulumi.StringPtrInput `pulumi:"name"`
}

func (NetworkRefArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRef)(nil)).Elem()
}

func (i NetworkRefArgs) ToNetworkRefOutput() NetworkRefOutput {
	return i.ToNetworkRefOutputWithContext(context.Background())
}

func (i NetworkRefArgs) ToNetworkRefOutputWithContext(ctx context.Context) NetworkRefOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRefOutput)
}





type NetworkRefArrayInput interface {
	pulumi.Input

	ToNetworkRefArrayOutput() NetworkRefArrayOutput
	ToNetworkRefArrayOutputWithContext(context.Context) NetworkRefArrayOutput
}

type NetworkRefArray []NetworkRefInput

func (NetworkRefArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRef)(nil)).Elem()
}

func (i NetworkRefArray) ToNetworkRefArrayOutput() NetworkRefArrayOutput {
	return i.ToNetworkRefArrayOutputWithContext(context.Background())
}

func (i NetworkRefArray) ToNetworkRefArrayOutputWithContext(ctx context.Context) NetworkRefArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkRefArrayOutput)
}

type NetworkRefOutput struct{ *pulumi.OutputState }

func (NetworkRefOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRef)(nil)).Elem()
}

func (o NetworkRefOutput) ToNetworkRefOutput() NetworkRefOutput {
	return o
}

func (o NetworkRefOutput) ToNetworkRefOutputWithContext(ctx context.Context) NetworkRefOutput {
	return o
}

func (o NetworkRefOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRef) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type NetworkRefArrayOutput struct{ *pulumi.OutputState }

func (NetworkRefArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRef)(nil)).Elem()
}

func (o NetworkRefArrayOutput) ToNetworkRefArrayOutput() NetworkRefArrayOutput {
	return o
}

func (o NetworkRefArrayOutput) ToNetworkRefArrayOutputWithContext(ctx context.Context) NetworkRefArrayOutput {
	return o
}

func (o NetworkRefArrayOutput) Index(i pulumi.IntInput) NetworkRefOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRef {
		return vs[0].([]NetworkRef)[vs[1].(int)]
	}).(NetworkRefOutput)
}

type NetworkRefResponse struct {
	Name *string `pulumi:"name"`
}

type NetworkRefResponseOutput struct{ *pulumi.OutputState }

func (NetworkRefResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkRefResponse)(nil)).Elem()
}

func (o NetworkRefResponseOutput) ToNetworkRefResponseOutput() NetworkRefResponseOutput {
	return o
}

func (o NetworkRefResponseOutput) ToNetworkRefResponseOutputWithContext(ctx context.Context) NetworkRefResponseOutput {
	return o
}

func (o NetworkRefResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkRefResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type NetworkRefResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkRefResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkRefResponse)(nil)).Elem()
}

func (o NetworkRefResponseArrayOutput) ToNetworkRefResponseArrayOutput() NetworkRefResponseArrayOutput {
	return o
}

func (o NetworkRefResponseArrayOutput) ToNetworkRefResponseArrayOutputWithContext(ctx context.Context) NetworkRefResponseArrayOutput {
	return o
}

func (o NetworkRefResponseArrayOutput) Index(i pulumi.IntInput) NetworkRefResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkRefResponse {
		return vs[0].([]NetworkRefResponse)[vs[1].(int)]
	}).(NetworkRefResponseOutput)
}

type ResourceLimits struct {
	Cpu        *float64 `pulumi:"cpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}





type ResourceLimitsInput interface {
	pulumi.Input

	ToResourceLimitsOutput() ResourceLimitsOutput
	ToResourceLimitsOutputWithContext(context.Context) ResourceLimitsOutput
}

type ResourceLimitsArgs struct {
	Cpu        pulumi.Float64PtrInput `pulumi:"cpu"`
	MemoryInGB pulumi.Float64PtrInput `pulumi:"memoryInGB"`
}

func (ResourceLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (i ResourceLimitsArgs) ToResourceLimitsOutput() ResourceLimitsOutput {
	return i.ToResourceLimitsOutputWithContext(context.Background())
}

func (i ResourceLimitsArgs) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsOutput)
}

func (i ResourceLimitsArgs) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return i.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (i ResourceLimitsArgs) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsOutput).ToResourceLimitsPtrOutputWithContext(ctx)
}









type ResourceLimitsPtrInput interface {
	pulumi.Input

	ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput
	ToResourceLimitsPtrOutputWithContext(context.Context) ResourceLimitsPtrOutput
}

type resourceLimitsPtrType ResourceLimitsArgs

func ResourceLimitsPtr(v *ResourceLimitsArgs) ResourceLimitsPtrInput {
	return (*resourceLimitsPtrType)(v)
}

func (*resourceLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimits)(nil)).Elem()
}

func (i *resourceLimitsPtrType) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return i.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (i *resourceLimitsPtrType) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceLimitsPtrOutput)
}

type ResourceLimitsOutput struct{ *pulumi.OutputState }

func (ResourceLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsOutput) ToResourceLimitsOutput() ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsOutputWithContext(ctx context.Context) ResourceLimitsOutput {
	return o
}

func (o ResourceLimitsOutput) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return o.ToResourceLimitsPtrOutputWithContext(context.Background())
}

func (o ResourceLimitsOutput) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ResourceLimits) *ResourceLimits {
		return &v
	}).(ResourceLimitsPtrOutput)
}

func (o ResourceLimitsOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimits) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimits) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ResourceLimitsPtrOutput struct{ *pulumi.OutputState }

func (ResourceLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimits)(nil)).Elem()
}

func (o ResourceLimitsPtrOutput) ToResourceLimitsPtrOutput() ResourceLimitsPtrOutput {
	return o
}

func (o ResourceLimitsPtrOutput) ToResourceLimitsPtrOutputWithContext(ctx context.Context) ResourceLimitsPtrOutput {
	return o
}

func (o ResourceLimitsPtrOutput) Elem() ResourceLimitsOutput {
	return o.ApplyT(func(v *ResourceLimits) ResourceLimits {
		if v != nil {
			return *v
		}
		var ret ResourceLimits
		return ret
	}).(ResourceLimitsOutput)
}

func (o ResourceLimitsPtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsPtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ResourceLimitsResponse struct {
	Cpu        *float64 `pulumi:"cpu"`
	MemoryInGB *float64 `pulumi:"memoryInGB"`
}

type ResourceLimitsResponseOutput struct{ *pulumi.OutputState }

func (ResourceLimitsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceLimitsResponse)(nil)).Elem()
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponseOutput() ResourceLimitsResponseOutput {
	return o
}

func (o ResourceLimitsResponseOutput) ToResourceLimitsResponseOutputWithContext(ctx context.Context) ResourceLimitsResponseOutput {
	return o
}

func (o ResourceLimitsResponseOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *float64 { return v.Cpu }).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsResponseOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ResourceLimitsResponse) *float64 { return v.MemoryInGB }).(pulumi.Float64PtrOutput)
}

type ResourceLimitsResponsePtrOutput struct{ *pulumi.OutputState }

func (ResourceLimitsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ResourceLimitsResponse)(nil)).Elem()
}

func (o ResourceLimitsResponsePtrOutput) ToResourceLimitsResponsePtrOutput() ResourceLimitsResponsePtrOutput {
	return o
}

func (o ResourceLimitsResponsePtrOutput) ToResourceLimitsResponsePtrOutputWithContext(ctx context.Context) ResourceLimitsResponsePtrOutput {
	return o
}

func (o ResourceLimitsResponsePtrOutput) Elem() ResourceLimitsResponseOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) ResourceLimitsResponse {
		if v != nil {
			return *v
		}
		var ret ResourceLimitsResponse
		return ret
	}).(ResourceLimitsResponseOutput)
}

func (o ResourceLimitsResponsePtrOutput) Cpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.Cpu
	}).(pulumi.Float64PtrOutput)
}

func (o ResourceLimitsResponsePtrOutput) MemoryInGB() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *ResourceLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MemoryInGB
	}).(pulumi.Float64PtrOutput)
}

type ResourceRequests struct {
	Cpu        float64 `pulumi:"cpu"`
	MemoryInGB float64 `pulumi:"memoryInGB"`
}





type ResourceRequestsInput interface {
	pulumi.Input

	ToResourceRequestsOutput() ResourceRequestsOutput
	ToResourceRequestsOutputWithContext(context.Context) ResourceRequestsOutput
}

type ResourceRequestsArgs struct {
	Cpu        pulumi.Float64Input `pulumi:"cpu"`
	MemoryInGB pulumi.Float64Input `pulumi:"memoryInGB"`
}

func (ResourceRequestsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (i ResourceRequestsArgs) ToResourceRequestsOutput() ResourceRequestsOutput {
	return i.ToResourceRequestsOutputWithContext(context.Background())
}

func (i ResourceRequestsArgs) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequestsOutput)
}

type ResourceRequestsOutput struct{ *pulumi.OutputState }

func (ResourceRequestsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequests)(nil)).Elem()
}

func (o ResourceRequestsOutput) ToResourceRequestsOutput() ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) ToResourceRequestsOutputWithContext(ctx context.Context) ResourceRequestsOutput {
	return o
}

func (o ResourceRequestsOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o ResourceRequestsOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequests) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequestsResponse struct {
	Cpu        float64 `pulumi:"cpu"`
	MemoryInGB float64 `pulumi:"memoryInGB"`
}

type ResourceRequestsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequestsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequestsResponse)(nil)).Elem()
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutput() ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) ToResourceRequestsResponseOutputWithContext(ctx context.Context) ResourceRequestsResponseOutput {
	return o
}

func (o ResourceRequestsResponseOutput) Cpu() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequestsResponse) float64 { return v.Cpu }).(pulumi.Float64Output)
}

func (o ResourceRequestsResponseOutput) MemoryInGB() pulumi.Float64Output {
	return o.ApplyT(func(v ResourceRequestsResponse) float64 { return v.MemoryInGB }).(pulumi.Float64Output)
}

type ResourceRequirements struct {
	Limits   *ResourceLimits  `pulumi:"limits"`
	Requests ResourceRequests `pulumi:"requests"`
}





type ResourceRequirementsInput interface {
	pulumi.Input

	ToResourceRequirementsOutput() ResourceRequirementsOutput
	ToResourceRequirementsOutputWithContext(context.Context) ResourceRequirementsOutput
}

type ResourceRequirementsArgs struct {
	Limits   ResourceLimitsPtrInput `pulumi:"limits"`
	Requests ResourceRequestsInput  `pulumi:"requests"`
}

func (ResourceRequirementsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirements)(nil)).Elem()
}

func (i ResourceRequirementsArgs) ToResourceRequirementsOutput() ResourceRequirementsOutput {
	return i.ToResourceRequirementsOutputWithContext(context.Background())
}

func (i ResourceRequirementsArgs) ToResourceRequirementsOutputWithContext(ctx context.Context) ResourceRequirementsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ResourceRequirementsOutput)
}

type ResourceRequirementsOutput struct{ *pulumi.OutputState }

func (ResourceRequirementsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirements)(nil)).Elem()
}

func (o ResourceRequirementsOutput) ToResourceRequirementsOutput() ResourceRequirementsOutput {
	return o
}

func (o ResourceRequirementsOutput) ToResourceRequirementsOutputWithContext(ctx context.Context) ResourceRequirementsOutput {
	return o
}

func (o ResourceRequirementsOutput) Limits() ResourceLimitsPtrOutput {
	return o.ApplyT(func(v ResourceRequirements) *ResourceLimits { return v.Limits }).(ResourceLimitsPtrOutput)
}

func (o ResourceRequirementsOutput) Requests() ResourceRequestsOutput {
	return o.ApplyT(func(v ResourceRequirements) ResourceRequests { return v.Requests }).(ResourceRequestsOutput)
}

type ResourceRequirementsResponse struct {
	Limits   *ResourceLimitsResponse  `pulumi:"limits"`
	Requests ResourceRequestsResponse `pulumi:"requests"`
}

type ResourceRequirementsResponseOutput struct{ *pulumi.OutputState }

func (ResourceRequirementsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResourceRequirementsResponse)(nil)).Elem()
}

func (o ResourceRequirementsResponseOutput) ToResourceRequirementsResponseOutput() ResourceRequirementsResponseOutput {
	return o
}

func (o ResourceRequirementsResponseOutput) ToResourceRequirementsResponseOutputWithContext(ctx context.Context) ResourceRequirementsResponseOutput {
	return o
}

func (o ResourceRequirementsResponseOutput) Limits() ResourceLimitsResponsePtrOutput {
	return o.ApplyT(func(v ResourceRequirementsResponse) *ResourceLimitsResponse { return v.Limits }).(ResourceLimitsResponsePtrOutput)
}

func (o ResourceRequirementsResponseOutput) Requests() ResourceRequestsResponseOutput {
	return o.ApplyT(func(v ResourceRequirementsResponse) ResourceRequestsResponse { return v.Requests }).(ResourceRequestsResponseOutput)
}

type ServiceResourceDescription struct {
	CodePackages []ContainerCodePackageProperties `pulumi:"codePackages"`
	Description  *string                          `pulumi:"description"`
	Diagnostics  *DiagnosticsRef                  `pulumi:"diagnostics"`
	HealthState  *string                          `pulumi:"healthState"`
	Name         *string                          `pulumi:"name"`
	NetworkRefs  []NetworkRef                     `pulumi:"networkRefs"`
	OsType       string                           `pulumi:"osType"`
	ReplicaCount *int                             `pulumi:"replicaCount"`
}





type ServiceResourceDescriptionInput interface {
	pulumi.Input

	ToServiceResourceDescriptionOutput() ServiceResourceDescriptionOutput
	ToServiceResourceDescriptionOutputWithContext(context.Context) ServiceResourceDescriptionOutput
}

type ServiceResourceDescriptionArgs struct {
	CodePackages ContainerCodePackagePropertiesArrayInput `pulumi:"codePackages"`
	Description  pulumi.StringPtrInput                    `pulumi:"description"`
	Diagnostics  DiagnosticsRefPtrInput                   `pulumi:"diagnostics"`
	HealthState  pulumi.StringPtrInput                    `pulumi:"healthState"`
	Name         pulumi.StringPtrInput                    `pulumi:"name"`
	NetworkRefs  NetworkRefArrayInput                     `pulumi:"networkRefs"`
	OsType       pulumi.StringInput                       `pulumi:"osType"`
	ReplicaCount pulumi.IntPtrInput                       `pulumi:"replicaCount"`
}

func (ServiceResourceDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceDescription)(nil)).Elem()
}

func (i ServiceResourceDescriptionArgs) ToServiceResourceDescriptionOutput() ServiceResourceDescriptionOutput {
	return i.ToServiceResourceDescriptionOutputWithContext(context.Background())
}

func (i ServiceResourceDescriptionArgs) ToServiceResourceDescriptionOutputWithContext(ctx context.Context) ServiceResourceDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResourceDescriptionOutput)
}





type ServiceResourceDescriptionArrayInput interface {
	pulumi.Input

	ToServiceResourceDescriptionArrayOutput() ServiceResourceDescriptionArrayOutput
	ToServiceResourceDescriptionArrayOutputWithContext(context.Context) ServiceResourceDescriptionArrayOutput
}

type ServiceResourceDescriptionArray []ServiceResourceDescriptionInput

func (ServiceResourceDescriptionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceDescription)(nil)).Elem()
}

func (i ServiceResourceDescriptionArray) ToServiceResourceDescriptionArrayOutput() ServiceResourceDescriptionArrayOutput {
	return i.ToServiceResourceDescriptionArrayOutputWithContext(context.Background())
}

func (i ServiceResourceDescriptionArray) ToServiceResourceDescriptionArrayOutputWithContext(ctx context.Context) ServiceResourceDescriptionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceResourceDescriptionArrayOutput)
}

type ServiceResourceDescriptionOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceDescription)(nil)).Elem()
}

func (o ServiceResourceDescriptionOutput) ToServiceResourceDescriptionOutput() ServiceResourceDescriptionOutput {
	return o
}

func (o ServiceResourceDescriptionOutput) ToServiceResourceDescriptionOutputWithContext(ctx context.Context) ServiceResourceDescriptionOutput {
	return o
}

func (o ServiceResourceDescriptionOutput) CodePackages() ContainerCodePackagePropertiesArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescription) []ContainerCodePackageProperties { return v.CodePackages }).(ContainerCodePackagePropertiesArrayOutput)
}

func (o ServiceResourceDescriptionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionOutput) Diagnostics() DiagnosticsRefPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *DiagnosticsRef { return v.Diagnostics }).(DiagnosticsRefPtrOutput)
}

func (o ServiceResourceDescriptionOutput) HealthState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *string { return v.HealthState }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionOutput) NetworkRefs() NetworkRefArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescription) []NetworkRef { return v.NetworkRefs }).(NetworkRefArrayOutput)
}

func (o ServiceResourceDescriptionOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescription) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescription) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

type ServiceResourceDescriptionArrayOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceDescription)(nil)).Elem()
}

func (o ServiceResourceDescriptionArrayOutput) ToServiceResourceDescriptionArrayOutput() ServiceResourceDescriptionArrayOutput {
	return o
}

func (o ServiceResourceDescriptionArrayOutput) ToServiceResourceDescriptionArrayOutputWithContext(ctx context.Context) ServiceResourceDescriptionArrayOutput {
	return o
}

func (o ServiceResourceDescriptionArrayOutput) Index(i pulumi.IntInput) ServiceResourceDescriptionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResourceDescription {
		return vs[0].([]ServiceResourceDescription)[vs[1].(int)]
	}).(ServiceResourceDescriptionOutput)
}

type ServiceResourceDescriptionResponse struct {
	CodePackages []ContainerCodePackagePropertiesResponse `pulumi:"codePackages"`
	Description  *string                                  `pulumi:"description"`
	Diagnostics  *DiagnosticsRefResponse                  `pulumi:"diagnostics"`
	HealthState  *string                                  `pulumi:"healthState"`
	Id           string                                   `pulumi:"id"`
	Name         *string                                  `pulumi:"name"`
	NetworkRefs  []NetworkRefResponse                     `pulumi:"networkRefs"`
	OsType       string                                   `pulumi:"osType"`
	ReplicaCount *int                                     `pulumi:"replicaCount"`
	Status       string                                   `pulumi:"status"`
	Type         string                                   `pulumi:"type"`
}

type ServiceResourceDescriptionResponseOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceResourceDescriptionResponse)(nil)).Elem()
}

func (o ServiceResourceDescriptionResponseOutput) ToServiceResourceDescriptionResponseOutput() ServiceResourceDescriptionResponseOutput {
	return o
}

func (o ServiceResourceDescriptionResponseOutput) ToServiceResourceDescriptionResponseOutputWithContext(ctx context.Context) ServiceResourceDescriptionResponseOutput {
	return o
}

func (o ServiceResourceDescriptionResponseOutput) CodePackages() ContainerCodePackagePropertiesResponseArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) []ContainerCodePackagePropertiesResponse {
		return v.CodePackages
	}).(ContainerCodePackagePropertiesResponseArrayOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Diagnostics() DiagnosticsRefResponsePtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *DiagnosticsRefResponse { return v.Diagnostics }).(DiagnosticsRefResponsePtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) HealthState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *string { return v.HealthState }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) NetworkRefs() NetworkRefResponseArrayOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) []NetworkRefResponse { return v.NetworkRefs }).(NetworkRefResponseArrayOutput)
}

func (o ServiceResourceDescriptionResponseOutput) OsType() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.OsType }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) ReplicaCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) *int { return v.ReplicaCount }).(pulumi.IntPtrOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o ServiceResourceDescriptionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceResourceDescriptionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceResourceDescriptionResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceResourceDescriptionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceResourceDescriptionResponse)(nil)).Elem()
}

func (o ServiceResourceDescriptionResponseArrayOutput) ToServiceResourceDescriptionResponseArrayOutput() ServiceResourceDescriptionResponseArrayOutput {
	return o
}

func (o ServiceResourceDescriptionResponseArrayOutput) ToServiceResourceDescriptionResponseArrayOutputWithContext(ctx context.Context) ServiceResourceDescriptionResponseArrayOutput {
	return o
}

func (o ServiceResourceDescriptionResponseArrayOutput) Index(i pulumi.IntInput) ServiceResourceDescriptionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceResourceDescriptionResponse {
		return vs[0].([]ServiceResourceDescriptionResponse)[vs[1].(int)]
	}).(ServiceResourceDescriptionResponseOutput)
}

type Setting struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type SettingInput interface {
	pulumi.Input

	ToSettingOutput() SettingOutput
	ToSettingOutputWithContext(context.Context) SettingOutput
}

type SettingArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (SettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Setting)(nil)).Elem()
}

func (i SettingArgs) ToSettingOutput() SettingOutput {
	return i.ToSettingOutputWithContext(context.Background())
}

func (i SettingArgs) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingOutput)
}





type SettingArrayInput interface {
	pulumi.Input

	ToSettingArrayOutput() SettingArrayOutput
	ToSettingArrayOutputWithContext(context.Context) SettingArrayOutput
}

type SettingArray []SettingInput

func (SettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Setting)(nil)).Elem()
}

func (i SettingArray) ToSettingArrayOutput() SettingArrayOutput {
	return i.ToSettingArrayOutputWithContext(context.Background())
}

func (i SettingArray) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SettingArrayOutput)
}

type SettingOutput struct{ *pulumi.OutputState }

func (SettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Setting)(nil)).Elem()
}

func (o SettingOutput) ToSettingOutput() SettingOutput {
	return o
}

func (o SettingOutput) ToSettingOutputWithContext(ctx context.Context) SettingOutput {
	return o
}

func (o SettingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Setting) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SettingOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Setting) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SettingArrayOutput struct{ *pulumi.OutputState }

func (SettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]Setting)(nil)).Elem()
}

func (o SettingArrayOutput) ToSettingArrayOutput() SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) ToSettingArrayOutputWithContext(ctx context.Context) SettingArrayOutput {
	return o
}

func (o SettingArrayOutput) Index(i pulumi.IntInput) SettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) Setting {
		return vs[0].([]Setting)[vs[1].(int)]
	}).(SettingOutput)
}

type SettingResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type SettingResponseOutput struct{ *pulumi.OutputState }

func (SettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SettingResponse)(nil)).Elem()
}

func (o SettingResponseOutput) ToSettingResponseOutput() SettingResponseOutput {
	return o
}

func (o SettingResponseOutput) ToSettingResponseOutputWithContext(ctx context.Context) SettingResponseOutput {
	return o
}

func (o SettingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SettingResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SettingResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type SettingResponseArrayOutput struct{ *pulumi.OutputState }

func (SettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SettingResponse)(nil)).Elem()
}

func (o SettingResponseArrayOutput) ToSettingResponseArrayOutput() SettingResponseArrayOutput {
	return o
}

func (o SettingResponseArrayOutput) ToSettingResponseArrayOutputWithContext(ctx context.Context) SettingResponseArrayOutput {
	return o
}

func (o SettingResponseArrayOutput) Index(i pulumi.IntInput) SettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SettingResponse {
		return vs[0].([]SettingResponse)[vs[1].(int)]
	}).(SettingResponseOutput)
}

type VolumeProviderParametersAzureFile struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName string  `pulumi:"accountName"`
	ShareName   string  `pulumi:"shareName"`
}





type VolumeProviderParametersAzureFileInput interface {
	pulumi.Input

	ToVolumeProviderParametersAzureFileOutput() VolumeProviderParametersAzureFileOutput
	ToVolumeProviderParametersAzureFileOutputWithContext(context.Context) VolumeProviderParametersAzureFileOutput
}

type VolumeProviderParametersAzureFileArgs struct {
	AccountKey  pulumi.StringPtrInput `pulumi:"accountKey"`
	AccountName pulumi.StringInput    `pulumi:"accountName"`
	ShareName   pulumi.StringInput    `pulumi:"shareName"`
}

func (VolumeProviderParametersAzureFileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFileOutput() VolumeProviderParametersAzureFileOutput {
	return i.ToVolumeProviderParametersAzureFileOutputWithContext(context.Background())
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFileOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeProviderParametersAzureFileOutput)
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return i.ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Background())
}

func (i VolumeProviderParametersAzureFileArgs) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeProviderParametersAzureFileOutput).ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx)
}









type VolumeProviderParametersAzureFilePtrInput interface {
	pulumi.Input

	ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput
	ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Context) VolumeProviderParametersAzureFilePtrOutput
}

type volumeProviderParametersAzureFilePtrType VolumeProviderParametersAzureFileArgs

func VolumeProviderParametersAzureFilePtr(v *VolumeProviderParametersAzureFileArgs) VolumeProviderParametersAzureFilePtrInput {
	return (*volumeProviderParametersAzureFilePtrType)(v)
}

func (*volumeProviderParametersAzureFilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (i *volumeProviderParametersAzureFilePtrType) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return i.ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Background())
}

func (i *volumeProviderParametersAzureFilePtrType) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VolumeProviderParametersAzureFilePtrOutput)
}

type VolumeProviderParametersAzureFileOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFileOutput() VolumeProviderParametersAzureFileOutput {
	return o
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFileOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileOutput {
	return o
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return o.ToVolumeProviderParametersAzureFilePtrOutputWithContext(context.Background())
}

func (o VolumeProviderParametersAzureFileOutput) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VolumeProviderParametersAzureFile) *VolumeProviderParametersAzureFile {
		return &v
	}).(VolumeProviderParametersAzureFilePtrOutput)
}

func (o VolumeProviderParametersAzureFileOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFile) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFile) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o VolumeProviderParametersAzureFileOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFile) string { return v.ShareName }).(pulumi.StringOutput)
}

type VolumeProviderParametersAzureFilePtrOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProviderParametersAzureFile)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFilePtrOutput) ToVolumeProviderParametersAzureFilePtrOutput() VolumeProviderParametersAzureFilePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFilePtrOutput) ToVolumeProviderParametersAzureFilePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFilePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFilePtrOutput) Elem() VolumeProviderParametersAzureFileOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) VolumeProviderParametersAzureFile {
		if v != nil {
			return *v
		}
		var ret VolumeProviderParametersAzureFile
		return ret
	}).(VolumeProviderParametersAzureFileOutput)
}

func (o VolumeProviderParametersAzureFilePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFilePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) *string {
		if v == nil {
			return nil
		}
		return &v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFilePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFile) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

type VolumeProviderParametersAzureFileResponse struct {
	AccountKey  *string `pulumi:"accountKey"`
	AccountName string  `pulumi:"accountName"`
	ShareName   string  `pulumi:"shareName"`
}

type VolumeProviderParametersAzureFileResponseOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VolumeProviderParametersAzureFileResponse)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFileResponseOutput) ToVolumeProviderParametersAzureFileResponseOutput() VolumeProviderParametersAzureFileResponseOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponseOutput) ToVolumeProviderParametersAzureFileResponseOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileResponseOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponseOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFileResponse) *string { return v.AccountKey }).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileResponseOutput) AccountName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFileResponse) string { return v.AccountName }).(pulumi.StringOutput)
}

func (o VolumeProviderParametersAzureFileResponseOutput) ShareName() pulumi.StringOutput {
	return o.ApplyT(func(v VolumeProviderParametersAzureFileResponse) string { return v.ShareName }).(pulumi.StringOutput)
}

type VolumeProviderParametersAzureFileResponsePtrOutput struct{ *pulumi.OutputState }

func (VolumeProviderParametersAzureFileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VolumeProviderParametersAzureFileResponse)(nil)).Elem()
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) ToVolumeProviderParametersAzureFileResponsePtrOutput() VolumeProviderParametersAzureFileResponsePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) ToVolumeProviderParametersAzureFileResponsePtrOutputWithContext(ctx context.Context) VolumeProviderParametersAzureFileResponsePtrOutput {
	return o
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) Elem() VolumeProviderParametersAzureFileResponseOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) VolumeProviderParametersAzureFileResponse {
		if v != nil {
			return *v
		}
		var ret VolumeProviderParametersAzureFileResponse
		return ret
	}).(VolumeProviderParametersAzureFileResponseOutput)
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) AccountKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccountKey
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AccountName
	}).(pulumi.StringPtrOutput)
}

func (o VolumeProviderParametersAzureFileResponsePtrOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VolumeProviderParametersAzureFileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ShareName
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionArrayOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionResponseOutput{})
	pulumi.RegisterOutputType(AzureInternalMonitoringPipelineSinkDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesArrayOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesResponseOutput{})
	pulumi.RegisterOutputType(ContainerCodePackagePropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerEventResponseOutput{})
	pulumi.RegisterOutputType(ContainerEventResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerInstanceViewResponseOutput{})
	pulumi.RegisterOutputType(ContainerLabelOutput{})
	pulumi.RegisterOutputType(ContainerLabelArrayOutput{})
	pulumi.RegisterOutputType(ContainerLabelResponseOutput{})
	pulumi.RegisterOutputType(ContainerLabelResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerStateResponseOutput{})
	pulumi.RegisterOutputType(ContainerStateResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerVolumeOutput{})
	pulumi.RegisterOutputType(ContainerVolumeArrayOutput{})
	pulumi.RegisterOutputType(ContainerVolumeResponseOutput{})
	pulumi.RegisterOutputType(ContainerVolumeResponseArrayOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefPtrOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefResponseOutput{})
	pulumi.RegisterOutputType(DiagnosticsRefResponsePtrOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesArrayOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EndpointPropertiesResponseArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVariableResponseArrayOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialPtrOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponseOutput{})
	pulumi.RegisterOutputType(ImageRegistryCredentialResponsePtrOutput{})
	pulumi.RegisterOutputType(IngressConfigOutput{})
	pulumi.RegisterOutputType(IngressConfigPtrOutput{})
	pulumi.RegisterOutputType(IngressConfigResponseOutput{})
	pulumi.RegisterOutputType(IngressConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(Layer4IngressConfigOutput{})
	pulumi.RegisterOutputType(Layer4IngressConfigArrayOutput{})
	pulumi.RegisterOutputType(Layer4IngressConfigResponseOutput{})
	pulumi.RegisterOutputType(Layer4IngressConfigResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkRefOutput{})
	pulumi.RegisterOutputType(NetworkRefArrayOutput{})
	pulumi.RegisterOutputType(NetworkRefResponseOutput{})
	pulumi.RegisterOutputType(NetworkRefResponseArrayOutput{})
	pulumi.RegisterOutputType(ResourceLimitsOutput{})
	pulumi.RegisterOutputType(ResourceLimitsPtrOutput{})
	pulumi.RegisterOutputType(ResourceLimitsResponseOutput{})
	pulumi.RegisterOutputType(ResourceLimitsResponsePtrOutput{})
	pulumi.RegisterOutputType(ResourceRequestsOutput{})
	pulumi.RegisterOutputType(ResourceRequestsResponseOutput{})
	pulumi.RegisterOutputType(ResourceRequirementsOutput{})
	pulumi.RegisterOutputType(ResourceRequirementsResponseOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionArrayOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionResponseOutput{})
	pulumi.RegisterOutputType(ServiceResourceDescriptionResponseArrayOutput{})
	pulumi.RegisterOutputType(SettingOutput{})
	pulumi.RegisterOutputType(SettingArrayOutput{})
	pulumi.RegisterOutputType(SettingResponseOutput{})
	pulumi.RegisterOutputType(SettingResponseArrayOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFileOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFilePtrOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFileResponseOutput{})
	pulumi.RegisterOutputType(VolumeProviderParametersAzureFileResponsePtrOutput{})
}
