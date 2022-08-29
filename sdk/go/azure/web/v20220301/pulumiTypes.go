


package v20220301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type ApiDefinitionInfo struct {
	Url *string `pulumi:"url"`
}





type ApiDefinitionInfoInput interface {
	pulumi.Input

	ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput
	ToApiDefinitionInfoOutputWithContext(context.Context) ApiDefinitionInfoOutput
}

type ApiDefinitionInfoArgs struct {
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (ApiDefinitionInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfo)(nil)).Elem()
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput {
	return i.ToApiDefinitionInfoOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoOutputWithContext(ctx context.Context) ApiDefinitionInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoOutput)
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return i.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoArgs) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoOutput).ToApiDefinitionInfoPtrOutputWithContext(ctx)
}









type ApiDefinitionInfoPtrInput interface {
	pulumi.Input

	ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput
	ToApiDefinitionInfoPtrOutputWithContext(context.Context) ApiDefinitionInfoPtrOutput
}

type apiDefinitionInfoPtrType ApiDefinitionInfoArgs

func ApiDefinitionInfoPtr(v *ApiDefinitionInfoArgs) ApiDefinitionInfoPtrInput {
	return (*apiDefinitionInfoPtrType)(v)
}

func (*apiDefinitionInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfo)(nil)).Elem()
}

func (i *apiDefinitionInfoPtrType) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return i.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (i *apiDefinitionInfoPtrType) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoPtrOutput)
}

type ApiDefinitionInfoOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfo)(nil)).Elem()
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoOutput() ApiDefinitionInfoOutput {
	return o
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoOutputWithContext(ctx context.Context) ApiDefinitionInfoOutput {
	return o
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return o.ToApiDefinitionInfoPtrOutputWithContext(context.Background())
}

func (o ApiDefinitionInfoOutput) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiDefinitionInfo) *ApiDefinitionInfo {
		return &v
	}).(ApiDefinitionInfoPtrOutput)
}

func (o ApiDefinitionInfoOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDefinitionInfo) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoPtrOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfo)(nil)).Elem()
}

func (o ApiDefinitionInfoPtrOutput) ToApiDefinitionInfoPtrOutput() ApiDefinitionInfoPtrOutput {
	return o
}

func (o ApiDefinitionInfoPtrOutput) ToApiDefinitionInfoPtrOutputWithContext(ctx context.Context) ApiDefinitionInfoPtrOutput {
	return o
}

func (o ApiDefinitionInfoPtrOutput) Elem() ApiDefinitionInfoOutput {
	return o.ApplyT(func(v *ApiDefinitionInfo) ApiDefinitionInfo {
		if v != nil {
			return *v
		}
		var ret ApiDefinitionInfo
		return ret
	}).(ApiDefinitionInfoOutput)
}

func (o ApiDefinitionInfoPtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinitionInfo) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoResponse struct {
	Url *string `pulumi:"url"`
}

type ApiDefinitionInfoResponseOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfoResponse)(nil)).Elem()
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponseOutput() ApiDefinitionInfoResponseOutput {
	return o
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponseOutputWithContext(ctx context.Context) ApiDefinitionInfoResponseOutput {
	return o
}

func (o ApiDefinitionInfoResponseOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiDefinitionInfoResponse) *string { return v.Url }).(pulumi.StringPtrOutput)
}

type ApiDefinitionInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiDefinitionInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfoResponse)(nil)).Elem()
}

func (o ApiDefinitionInfoResponsePtrOutput) ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput {
	return o
}

func (o ApiDefinitionInfoResponsePtrOutput) ToApiDefinitionInfoResponsePtrOutputWithContext(ctx context.Context) ApiDefinitionInfoResponsePtrOutput {
	return o
}

func (o ApiDefinitionInfoResponsePtrOutput) Elem() ApiDefinitionInfoResponseOutput {
	return o.ApplyT(func(v *ApiDefinitionInfoResponse) ApiDefinitionInfoResponse {
		if v != nil {
			return *v
		}
		var ret ApiDefinitionInfoResponse
		return ret
	}).(ApiDefinitionInfoResponseOutput)
}

func (o ApiDefinitionInfoResponsePtrOutput) Url() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiDefinitionInfoResponse) *string {
		if v == nil {
			return nil
		}
		return v.Url
	}).(pulumi.StringPtrOutput)
}

type ApiManagementConfig struct {
	Id *string `pulumi:"id"`
}





type ApiManagementConfigInput interface {
	pulumi.Input

	ToApiManagementConfigOutput() ApiManagementConfigOutput
	ToApiManagementConfigOutputWithContext(context.Context) ApiManagementConfigOutput
}

type ApiManagementConfigArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiManagementConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfig)(nil)).Elem()
}

func (i ApiManagementConfigArgs) ToApiManagementConfigOutput() ApiManagementConfigOutput {
	return i.ToApiManagementConfigOutputWithContext(context.Background())
}

func (i ApiManagementConfigArgs) ToApiManagementConfigOutputWithContext(ctx context.Context) ApiManagementConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigOutput)
}

func (i ApiManagementConfigArgs) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return i.ToApiManagementConfigPtrOutputWithContext(context.Background())
}

func (i ApiManagementConfigArgs) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigOutput).ToApiManagementConfigPtrOutputWithContext(ctx)
}









type ApiManagementConfigPtrInput interface {
	pulumi.Input

	ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput
	ToApiManagementConfigPtrOutputWithContext(context.Context) ApiManagementConfigPtrOutput
}

type apiManagementConfigPtrType ApiManagementConfigArgs

func ApiManagementConfigPtr(v *ApiManagementConfigArgs) ApiManagementConfigPtrInput {
	return (*apiManagementConfigPtrType)(v)
}

func (*apiManagementConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfig)(nil)).Elem()
}

func (i *apiManagementConfigPtrType) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return i.ToApiManagementConfigPtrOutputWithContext(context.Background())
}

func (i *apiManagementConfigPtrType) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigPtrOutput)
}

type ApiManagementConfigOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfig)(nil)).Elem()
}

func (o ApiManagementConfigOutput) ToApiManagementConfigOutput() ApiManagementConfigOutput {
	return o
}

func (o ApiManagementConfigOutput) ToApiManagementConfigOutputWithContext(ctx context.Context) ApiManagementConfigOutput {
	return o
}

func (o ApiManagementConfigOutput) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return o.ToApiManagementConfigPtrOutputWithContext(context.Background())
}

func (o ApiManagementConfigOutput) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementConfig) *ApiManagementConfig {
		return &v
	}).(ApiManagementConfigPtrOutput)
}

func (o ApiManagementConfigOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiManagementConfig) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiManagementConfigPtrOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfig)(nil)).Elem()
}

func (o ApiManagementConfigPtrOutput) ToApiManagementConfigPtrOutput() ApiManagementConfigPtrOutput {
	return o
}

func (o ApiManagementConfigPtrOutput) ToApiManagementConfigPtrOutputWithContext(ctx context.Context) ApiManagementConfigPtrOutput {
	return o
}

func (o ApiManagementConfigPtrOutput) Elem() ApiManagementConfigOutput {
	return o.ApplyT(func(v *ApiManagementConfig) ApiManagementConfig {
		if v != nil {
			return *v
		}
		var ret ApiManagementConfig
		return ret
	}).(ApiManagementConfigOutput)
}

func (o ApiManagementConfigPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementConfig) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type ApiManagementConfigResponse struct {
	Id *string `pulumi:"id"`
}

type ApiManagementConfigResponseOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfigResponse)(nil)).Elem()
}

func (o ApiManagementConfigResponseOutput) ToApiManagementConfigResponseOutput() ApiManagementConfigResponseOutput {
	return o
}

func (o ApiManagementConfigResponseOutput) ToApiManagementConfigResponseOutputWithContext(ctx context.Context) ApiManagementConfigResponseOutput {
	return o
}

func (o ApiManagementConfigResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ApiManagementConfigResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type ApiManagementConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ApiManagementConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfigResponse)(nil)).Elem()
}

func (o ApiManagementConfigResponsePtrOutput) ToApiManagementConfigResponsePtrOutput() ApiManagementConfigResponsePtrOutput {
	return o
}

func (o ApiManagementConfigResponsePtrOutput) ToApiManagementConfigResponsePtrOutputWithContext(ctx context.Context) ApiManagementConfigResponsePtrOutput {
	return o
}

func (o ApiManagementConfigResponsePtrOutput) Elem() ApiManagementConfigResponseOutput {
	return o.ApplyT(func(v *ApiManagementConfigResponse) ApiManagementConfigResponse {
		if v != nil {
			return *v
		}
		var ret ApiManagementConfigResponse
		return ret
	}).(ApiManagementConfigResponseOutput)
}

func (o ApiManagementConfigResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ApiManagementConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

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

type ApplicationLogsConfig struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfig  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfig `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfig        `pulumi:"fileSystem"`
}


func (val *ApplicationLogsConfig) Defaults() *ApplicationLogsConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FileSystem = tmp.FileSystem.Defaults()

	return &tmp
}





type ApplicationLogsConfigInput interface {
	pulumi.Input

	ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput
	ToApplicationLogsConfigOutputWithContext(context.Context) ApplicationLogsConfigOutput
}

type ApplicationLogsConfigArgs struct {
	AzureBlobStorage  AzureBlobStorageApplicationLogsConfigPtrInput  `pulumi:"azureBlobStorage"`
	AzureTableStorage AzureTableStorageApplicationLogsConfigPtrInput `pulumi:"azureTableStorage"`
	FileSystem        FileSystemApplicationLogsConfigPtrInput        `pulumi:"fileSystem"`
}


func (val *ApplicationLogsConfigArgs) Defaults() *ApplicationLogsConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
}
func (ApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfig)(nil)).Elem()
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput {
	return i.ToApplicationLogsConfigOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigOutputWithContext(ctx context.Context) ApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigOutput)
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return i.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigArgs) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigOutput).ToApplicationLogsConfigPtrOutputWithContext(ctx)
}









type ApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput
	ToApplicationLogsConfigPtrOutputWithContext(context.Context) ApplicationLogsConfigPtrOutput
}

type applicationLogsConfigPtrType ApplicationLogsConfigArgs

func ApplicationLogsConfigPtr(v *ApplicationLogsConfigArgs) ApplicationLogsConfigPtrInput {
	return (*applicationLogsConfigPtrType)(v)
}

func (*applicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfig)(nil)).Elem()
}

func (i *applicationLogsConfigPtrType) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return i.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *applicationLogsConfigPtrType) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfig)(nil)).Elem()
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigOutput() ApplicationLogsConfigOutput {
	return o
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigOutputWithContext(ctx context.Context) ApplicationLogsConfigOutput {
	return o
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return o.ToApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o ApplicationLogsConfigOutput) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationLogsConfig) *ApplicationLogsConfig {
		return &v
	}).(ApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig { return v.AzureBlobStorage }).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig { return v.AzureTableStorage }).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigOutput) FileSystem() FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfig) *FileSystemApplicationLogsConfig { return v.FileSystem }).(FileSystemApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfig)(nil)).Elem()
}

func (o ApplicationLogsConfigPtrOutput) ToApplicationLogsConfigPtrOutput() ApplicationLogsConfigPtrOutput {
	return o
}

func (o ApplicationLogsConfigPtrOutput) ToApplicationLogsConfigPtrOutputWithContext(ctx context.Context) ApplicationLogsConfigPtrOutput {
	return o
}

func (o ApplicationLogsConfigPtrOutput) Elem() ApplicationLogsConfigOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) ApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret ApplicationLogsConfig
		return ret
	}).(ApplicationLogsConfigOutput)
}

func (o ApplicationLogsConfigPtrOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigPtrOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o ApplicationLogsConfigPtrOutput) FileSystem() FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfig) *FileSystemApplicationLogsConfig {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemApplicationLogsConfigPtrOutput)
}

type ApplicationLogsConfigResponse struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfigResponse  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfigResponse `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfigResponse        `pulumi:"fileSystem"`
}


func (val *ApplicationLogsConfigResponse) Defaults() *ApplicationLogsConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.FileSystem = tmp.FileSystem.Defaults()

	return &tmp
}

type ApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfigResponse)(nil)).Elem()
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponseOutput() ApplicationLogsConfigResponseOutput {
	return o
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponseOutputWithContext(ctx context.Context) ApplicationLogsConfigResponseOutput {
	return o
}

func (o ApplicationLogsConfigResponseOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponseOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponseOutput) FileSystem() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v ApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse { return v.FileSystem }).(FileSystemApplicationLogsConfigResponsePtrOutput)
}

type ApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (ApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfigResponse)(nil)).Elem()
}

func (o ApplicationLogsConfigResponsePtrOutput) ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o ApplicationLogsConfigResponsePtrOutput) ToApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) ApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o ApplicationLogsConfigResponsePtrOutput) Elem() ApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) ApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret ApplicationLogsConfigResponse
		return ret
	}).(ApplicationLogsConfigResponseOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) AzureBlobStorage() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) AzureTableStorage() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureTableStorage
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
}

func (o ApplicationLogsConfigResponsePtrOutput) FileSystem() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *ApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemApplicationLogsConfigResponsePtrOutput)
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

type ArmIdWrapperResponse struct {
	Id string `pulumi:"id"`
}

type ArmIdWrapperResponseOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutput() ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponseOutputWithContext(ctx context.Context) ArmIdWrapperResponseOutput {
	return o
}

func (o ArmIdWrapperResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ArmIdWrapperResponse) string { return v.Id }).(pulumi.StringOutput)
}

type ArmIdWrapperResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmIdWrapperResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdWrapperResponse)(nil)).Elem()
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return o
}

func (o ArmIdWrapperResponsePtrOutput) Elem() ArmIdWrapperResponseOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) ArmIdWrapperResponse {
		if v != nil {
			return *v
		}
		var ret ArmIdWrapperResponse
		return ret
	}).(ArmIdWrapperResponseOutput)
}

func (o ArmIdWrapperResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmIdWrapperResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type ArmPlanResponse struct {
	Name          *string `pulumi:"name"`
	Product       *string `pulumi:"product"`
	PromotionCode *string `pulumi:"promotionCode"`
	Publisher     *string `pulumi:"publisher"`
	Version       *string `pulumi:"version"`
}

type ArmPlanResponseOutput struct{ *pulumi.OutputState }

func (ArmPlanResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmPlanResponse)(nil)).Elem()
}

func (o ArmPlanResponseOutput) ToArmPlanResponseOutput() ArmPlanResponseOutput {
	return o
}

func (o ArmPlanResponseOutput) ToArmPlanResponseOutputWithContext(ctx context.Context) ArmPlanResponseOutput {
	return o
}

func (o ArmPlanResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Product }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.PromotionCode }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Publisher }).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponseOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ArmPlanResponse) *string { return v.Version }).(pulumi.StringPtrOutput)
}

type ArmPlanResponsePtrOutput struct{ *pulumi.OutputState }

func (ArmPlanResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmPlanResponse)(nil)).Elem()
}

func (o ArmPlanResponsePtrOutput) ToArmPlanResponsePtrOutput() ArmPlanResponsePtrOutput {
	return o
}

func (o ArmPlanResponsePtrOutput) ToArmPlanResponsePtrOutputWithContext(ctx context.Context) ArmPlanResponsePtrOutput {
	return o
}

func (o ArmPlanResponsePtrOutput) Elem() ArmPlanResponseOutput {
	return o.ApplyT(func(v *ArmPlanResponse) ArmPlanResponse {
		if v != nil {
			return *v
		}
		var ret ArmPlanResponse
		return ret
	}).(ArmPlanResponseOutput)
}

func (o ArmPlanResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Product() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Product
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) PromotionCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.PromotionCode
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Publisher() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Publisher
	}).(pulumi.StringPtrOutput)
}

func (o ArmPlanResponsePtrOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ArmPlanResponse) *string {
		if v == nil {
			return nil
		}
		return v.Version
	}).(pulumi.StringPtrOutput)
}

type AseV3NetworkingConfiguration struct {
	AllowNewPrivateEndpointConnections *bool   `pulumi:"allowNewPrivateEndpointConnections"`
	FtpEnabled                         *bool   `pulumi:"ftpEnabled"`
	InboundIpAddressOverride           *string `pulumi:"inboundIpAddressOverride"`
	Kind                               *string `pulumi:"kind"`
	RemoteDebugEnabled                 *bool   `pulumi:"remoteDebugEnabled"`
}





type AseV3NetworkingConfigurationInput interface {
	pulumi.Input

	ToAseV3NetworkingConfigurationOutput() AseV3NetworkingConfigurationOutput
	ToAseV3NetworkingConfigurationOutputWithContext(context.Context) AseV3NetworkingConfigurationOutput
}

type AseV3NetworkingConfigurationArgs struct {
	AllowNewPrivateEndpointConnections pulumi.BoolPtrInput   `pulumi:"allowNewPrivateEndpointConnections"`
	FtpEnabled                         pulumi.BoolPtrInput   `pulumi:"ftpEnabled"`
	InboundIpAddressOverride           pulumi.StringPtrInput `pulumi:"inboundIpAddressOverride"`
	Kind                               pulumi.StringPtrInput `pulumi:"kind"`
	RemoteDebugEnabled                 pulumi.BoolPtrInput   `pulumi:"remoteDebugEnabled"`
}

func (AseV3NetworkingConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AseV3NetworkingConfiguration)(nil)).Elem()
}

func (i AseV3NetworkingConfigurationArgs) ToAseV3NetworkingConfigurationOutput() AseV3NetworkingConfigurationOutput {
	return i.ToAseV3NetworkingConfigurationOutputWithContext(context.Background())
}

func (i AseV3NetworkingConfigurationArgs) ToAseV3NetworkingConfigurationOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AseV3NetworkingConfigurationOutput)
}

func (i AseV3NetworkingConfigurationArgs) ToAseV3NetworkingConfigurationPtrOutput() AseV3NetworkingConfigurationPtrOutput {
	return i.ToAseV3NetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (i AseV3NetworkingConfigurationArgs) ToAseV3NetworkingConfigurationPtrOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AseV3NetworkingConfigurationOutput).ToAseV3NetworkingConfigurationPtrOutputWithContext(ctx)
}









type AseV3NetworkingConfigurationPtrInput interface {
	pulumi.Input

	ToAseV3NetworkingConfigurationPtrOutput() AseV3NetworkingConfigurationPtrOutput
	ToAseV3NetworkingConfigurationPtrOutputWithContext(context.Context) AseV3NetworkingConfigurationPtrOutput
}

type aseV3NetworkingConfigurationPtrType AseV3NetworkingConfigurationArgs

func AseV3NetworkingConfigurationPtr(v *AseV3NetworkingConfigurationArgs) AseV3NetworkingConfigurationPtrInput {
	return (*aseV3NetworkingConfigurationPtrType)(v)
}

func (*aseV3NetworkingConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AseV3NetworkingConfiguration)(nil)).Elem()
}

func (i *aseV3NetworkingConfigurationPtrType) ToAseV3NetworkingConfigurationPtrOutput() AseV3NetworkingConfigurationPtrOutput {
	return i.ToAseV3NetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (i *aseV3NetworkingConfigurationPtrType) ToAseV3NetworkingConfigurationPtrOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AseV3NetworkingConfigurationPtrOutput)
}

type AseV3NetworkingConfigurationOutput struct{ *pulumi.OutputState }

func (AseV3NetworkingConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AseV3NetworkingConfiguration)(nil)).Elem()
}

func (o AseV3NetworkingConfigurationOutput) ToAseV3NetworkingConfigurationOutput() AseV3NetworkingConfigurationOutput {
	return o
}

func (o AseV3NetworkingConfigurationOutput) ToAseV3NetworkingConfigurationOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationOutput {
	return o
}

func (o AseV3NetworkingConfigurationOutput) ToAseV3NetworkingConfigurationPtrOutput() AseV3NetworkingConfigurationPtrOutput {
	return o.ToAseV3NetworkingConfigurationPtrOutputWithContext(context.Background())
}

func (o AseV3NetworkingConfigurationOutput) ToAseV3NetworkingConfigurationPtrOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AseV3NetworkingConfiguration) *AseV3NetworkingConfiguration {
		return &v
	}).(AseV3NetworkingConfigurationPtrOutput)
}

func (o AseV3NetworkingConfigurationOutput) AllowNewPrivateEndpointConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfiguration) *bool { return v.AllowNewPrivateEndpointConnections }).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationOutput) FtpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfiguration) *bool { return v.FtpEnabled }).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationOutput) InboundIpAddressOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfiguration) *string { return v.InboundIpAddressOverride }).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfiguration) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationOutput) RemoteDebugEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfiguration) *bool { return v.RemoteDebugEnabled }).(pulumi.BoolPtrOutput)
}

type AseV3NetworkingConfigurationPtrOutput struct{ *pulumi.OutputState }

func (AseV3NetworkingConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AseV3NetworkingConfiguration)(nil)).Elem()
}

func (o AseV3NetworkingConfigurationPtrOutput) ToAseV3NetworkingConfigurationPtrOutput() AseV3NetworkingConfigurationPtrOutput {
	return o
}

func (o AseV3NetworkingConfigurationPtrOutput) ToAseV3NetworkingConfigurationPtrOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationPtrOutput {
	return o
}

func (o AseV3NetworkingConfigurationPtrOutput) Elem() AseV3NetworkingConfigurationOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfiguration) AseV3NetworkingConfiguration {
		if v != nil {
			return *v
		}
		var ret AseV3NetworkingConfiguration
		return ret
	}).(AseV3NetworkingConfigurationOutput)
}

func (o AseV3NetworkingConfigurationPtrOutput) AllowNewPrivateEndpointConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.AllowNewPrivateEndpointConnections
	}).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationPtrOutput) FtpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.FtpEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationPtrOutput) InboundIpAddressOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.InboundIpAddressOverride
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationPtrOutput) RemoteDebugEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebugEnabled
	}).(pulumi.BoolPtrOutput)
}

type AseV3NetworkingConfigurationResponse struct {
	AllowNewPrivateEndpointConnections *bool    `pulumi:"allowNewPrivateEndpointConnections"`
	ExternalInboundIpAddresses         []string `pulumi:"externalInboundIpAddresses"`
	FtpEnabled                         *bool    `pulumi:"ftpEnabled"`
	Id                                 string   `pulumi:"id"`
	InboundIpAddressOverride           *string  `pulumi:"inboundIpAddressOverride"`
	InternalInboundIpAddresses         []string `pulumi:"internalInboundIpAddresses"`
	Kind                               *string  `pulumi:"kind"`
	LinuxOutboundIpAddresses           []string `pulumi:"linuxOutboundIpAddresses"`
	Name                               string   `pulumi:"name"`
	RemoteDebugEnabled                 *bool    `pulumi:"remoteDebugEnabled"`
	Type                               string   `pulumi:"type"`
	WindowsOutboundIpAddresses         []string `pulumi:"windowsOutboundIpAddresses"`
}

type AseV3NetworkingConfigurationResponseOutput struct{ *pulumi.OutputState }

func (AseV3NetworkingConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AseV3NetworkingConfigurationResponse)(nil)).Elem()
}

func (o AseV3NetworkingConfigurationResponseOutput) ToAseV3NetworkingConfigurationResponseOutput() AseV3NetworkingConfigurationResponseOutput {
	return o
}

func (o AseV3NetworkingConfigurationResponseOutput) ToAseV3NetworkingConfigurationResponseOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationResponseOutput {
	return o
}

func (o AseV3NetworkingConfigurationResponseOutput) AllowNewPrivateEndpointConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) *bool { return v.AllowNewPrivateEndpointConnections }).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) ExternalInboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) []string { return v.ExternalInboundIpAddresses }).(pulumi.StringArrayOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) FtpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) *bool { return v.FtpEnabled }).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) InboundIpAddressOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) *string { return v.InboundIpAddressOverride }).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) InternalInboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) []string { return v.InternalInboundIpAddresses }).(pulumi.StringArrayOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) LinuxOutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) []string { return v.LinuxOutboundIpAddresses }).(pulumi.StringArrayOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) RemoteDebugEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) *bool { return v.RemoteDebugEnabled }).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o AseV3NetworkingConfigurationResponseOutput) WindowsOutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v AseV3NetworkingConfigurationResponse) []string { return v.WindowsOutboundIpAddresses }).(pulumi.StringArrayOutput)
}

type AseV3NetworkingConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (AseV3NetworkingConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AseV3NetworkingConfigurationResponse)(nil)).Elem()
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) ToAseV3NetworkingConfigurationResponsePtrOutput() AseV3NetworkingConfigurationResponsePtrOutput {
	return o
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) ToAseV3NetworkingConfigurationResponsePtrOutputWithContext(ctx context.Context) AseV3NetworkingConfigurationResponsePtrOutput {
	return o
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) Elem() AseV3NetworkingConfigurationResponseOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) AseV3NetworkingConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret AseV3NetworkingConfigurationResponse
		return ret
	}).(AseV3NetworkingConfigurationResponseOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) AllowNewPrivateEndpointConnections() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AllowNewPrivateEndpointConnections
	}).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) ExternalInboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.ExternalInboundIpAddresses
	}).(pulumi.StringArrayOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) FtpEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.FtpEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) InboundIpAddressOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.InboundIpAddressOverride
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) InternalInboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.InternalInboundIpAddresses
	}).(pulumi.StringArrayOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) LinuxOutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.LinuxOutboundIpAddresses
	}).(pulumi.StringArrayOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) RemoteDebugEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebugEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

func (o AseV3NetworkingConfigurationResponsePtrOutput) WindowsOutboundIpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *AseV3NetworkingConfigurationResponse) []string {
		if v == nil {
			return nil
		}
		return v.WindowsOutboundIpAddresses
	}).(pulumi.StringArrayOutput)
}

type AutoHealActions struct {
	ActionType              *AutoHealActionType   `pulumi:"actionType"`
	CustomAction            *AutoHealCustomAction `pulumi:"customAction"`
	MinProcessExecutionTime *string               `pulumi:"minProcessExecutionTime"`
}





type AutoHealActionsInput interface {
	pulumi.Input

	ToAutoHealActionsOutput() AutoHealActionsOutput
	ToAutoHealActionsOutputWithContext(context.Context) AutoHealActionsOutput
}

type AutoHealActionsArgs struct {
	ActionType              AutoHealActionTypePtrInput   `pulumi:"actionType"`
	CustomAction            AutoHealCustomActionPtrInput `pulumi:"customAction"`
	MinProcessExecutionTime pulumi.StringPtrInput        `pulumi:"minProcessExecutionTime"`
}

func (AutoHealActionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActions)(nil)).Elem()
}

func (i AutoHealActionsArgs) ToAutoHealActionsOutput() AutoHealActionsOutput {
	return i.ToAutoHealActionsOutputWithContext(context.Background())
}

func (i AutoHealActionsArgs) ToAutoHealActionsOutputWithContext(ctx context.Context) AutoHealActionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsOutput)
}

func (i AutoHealActionsArgs) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return i.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (i AutoHealActionsArgs) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsOutput).ToAutoHealActionsPtrOutputWithContext(ctx)
}









type AutoHealActionsPtrInput interface {
	pulumi.Input

	ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput
	ToAutoHealActionsPtrOutputWithContext(context.Context) AutoHealActionsPtrOutput
}

type autoHealActionsPtrType AutoHealActionsArgs

func AutoHealActionsPtr(v *AutoHealActionsArgs) AutoHealActionsPtrInput {
	return (*autoHealActionsPtrType)(v)
}

func (*autoHealActionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActions)(nil)).Elem()
}

func (i *autoHealActionsPtrType) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return i.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (i *autoHealActionsPtrType) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsPtrOutput)
}

type AutoHealActionsOutput struct{ *pulumi.OutputState }

func (AutoHealActionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActions)(nil)).Elem()
}

func (o AutoHealActionsOutput) ToAutoHealActionsOutput() AutoHealActionsOutput {
	return o
}

func (o AutoHealActionsOutput) ToAutoHealActionsOutputWithContext(ctx context.Context) AutoHealActionsOutput {
	return o
}

func (o AutoHealActionsOutput) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return o.ToAutoHealActionsPtrOutputWithContext(context.Background())
}

func (o AutoHealActionsOutput) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActions) *AutoHealActions {
		return &v
	}).(AutoHealActionsPtrOutput)
}

func (o AutoHealActionsOutput) ActionType() AutoHealActionTypePtrOutput {
	return o.ApplyT(func(v AutoHealActions) *AutoHealActionType { return v.ActionType }).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionsOutput) CustomAction() AutoHealCustomActionPtrOutput {
	return o.ApplyT(func(v AutoHealActions) *AutoHealCustomAction { return v.CustomAction }).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealActionsOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActions) *string { return v.MinProcessExecutionTime }).(pulumi.StringPtrOutput)
}

type AutoHealActionsPtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActions)(nil)).Elem()
}

func (o AutoHealActionsPtrOutput) ToAutoHealActionsPtrOutput() AutoHealActionsPtrOutput {
	return o
}

func (o AutoHealActionsPtrOutput) ToAutoHealActionsPtrOutputWithContext(ctx context.Context) AutoHealActionsPtrOutput {
	return o
}

func (o AutoHealActionsPtrOutput) Elem() AutoHealActionsOutput {
	return o.ApplyT(func(v *AutoHealActions) AutoHealActions {
		if v != nil {
			return *v
		}
		var ret AutoHealActions
		return ret
	}).(AutoHealActionsOutput)
}

func (o AutoHealActionsPtrOutput) ActionType() AutoHealActionTypePtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *AutoHealActionType {
		if v == nil {
			return nil
		}
		return v.ActionType
	}).(AutoHealActionTypePtrOutput)
}

func (o AutoHealActionsPtrOutput) CustomAction() AutoHealCustomActionPtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *AutoHealCustomAction {
		if v == nil {
			return nil
		}
		return v.CustomAction
	}).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealActionsPtrOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActions) *string {
		if v == nil {
			return nil
		}
		return v.MinProcessExecutionTime
	}).(pulumi.StringPtrOutput)
}

type AutoHealActionsResponse struct {
	ActionType              *string                       `pulumi:"actionType"`
	CustomAction            *AutoHealCustomActionResponse `pulumi:"customAction"`
	MinProcessExecutionTime *string                       `pulumi:"minProcessExecutionTime"`
}

type AutoHealActionsResponseOutput struct{ *pulumi.OutputState }

func (AutoHealActionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionsResponse)(nil)).Elem()
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponseOutput() AutoHealActionsResponseOutput {
	return o
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponseOutputWithContext(ctx context.Context) AutoHealActionsResponseOutput {
	return o
}

func (o AutoHealActionsResponseOutput) ActionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *string { return v.ActionType }).(pulumi.StringPtrOutput)
}

func (o AutoHealActionsResponseOutput) CustomAction() AutoHealCustomActionResponsePtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *AutoHealCustomActionResponse { return v.CustomAction }).(AutoHealCustomActionResponsePtrOutput)
}

func (o AutoHealActionsResponseOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealActionsResponse) *string { return v.MinProcessExecutionTime }).(pulumi.StringPtrOutput)
}

type AutoHealActionsResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealActionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionsResponse)(nil)).Elem()
}

func (o AutoHealActionsResponsePtrOutput) ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput {
	return o
}

func (o AutoHealActionsResponsePtrOutput) ToAutoHealActionsResponsePtrOutputWithContext(ctx context.Context) AutoHealActionsResponsePtrOutput {
	return o
}

func (o AutoHealActionsResponsePtrOutput) Elem() AutoHealActionsResponseOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) AutoHealActionsResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealActionsResponse
		return ret
	}).(AutoHealActionsResponseOutput)
}

func (o AutoHealActionsResponsePtrOutput) ActionType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionType
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealActionsResponsePtrOutput) CustomAction() AutoHealCustomActionResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *AutoHealCustomActionResponse {
		if v == nil {
			return nil
		}
		return v.CustomAction
	}).(AutoHealCustomActionResponsePtrOutput)
}

func (o AutoHealActionsResponsePtrOutput) MinProcessExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealActionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinProcessExecutionTime
	}).(pulumi.StringPtrOutput)
}

type AutoHealCustomAction struct {
	Exe        *string `pulumi:"exe"`
	Parameters *string `pulumi:"parameters"`
}





type AutoHealCustomActionInput interface {
	pulumi.Input

	ToAutoHealCustomActionOutput() AutoHealCustomActionOutput
	ToAutoHealCustomActionOutputWithContext(context.Context) AutoHealCustomActionOutput
}

type AutoHealCustomActionArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Parameters pulumi.StringPtrInput `pulumi:"parameters"`
}

func (AutoHealCustomActionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomAction)(nil)).Elem()
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionOutput() AutoHealCustomActionOutput {
	return i.ToAutoHealCustomActionOutputWithContext(context.Background())
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionOutputWithContext(ctx context.Context) AutoHealCustomActionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionOutput)
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return i.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (i AutoHealCustomActionArgs) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionOutput).ToAutoHealCustomActionPtrOutputWithContext(ctx)
}









type AutoHealCustomActionPtrInput interface {
	pulumi.Input

	ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput
	ToAutoHealCustomActionPtrOutputWithContext(context.Context) AutoHealCustomActionPtrOutput
}

type autoHealCustomActionPtrType AutoHealCustomActionArgs

func AutoHealCustomActionPtr(v *AutoHealCustomActionArgs) AutoHealCustomActionPtrInput {
	return (*autoHealCustomActionPtrType)(v)
}

func (*autoHealCustomActionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomAction)(nil)).Elem()
}

func (i *autoHealCustomActionPtrType) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return i.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (i *autoHealCustomActionPtrType) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionPtrOutput)
}

type AutoHealCustomActionOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomAction)(nil)).Elem()
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionOutput() AutoHealCustomActionOutput {
	return o
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionOutputWithContext(ctx context.Context) AutoHealCustomActionOutput {
	return o
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return o.ToAutoHealCustomActionPtrOutputWithContext(context.Background())
}

func (o AutoHealCustomActionOutput) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealCustomAction) *AutoHealCustomAction {
		return &v
	}).(AutoHealCustomActionPtrOutput)
}

func (o AutoHealCustomActionOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomAction) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomAction) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionPtrOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomAction)(nil)).Elem()
}

func (o AutoHealCustomActionPtrOutput) ToAutoHealCustomActionPtrOutput() AutoHealCustomActionPtrOutput {
	return o
}

func (o AutoHealCustomActionPtrOutput) ToAutoHealCustomActionPtrOutputWithContext(ctx context.Context) AutoHealCustomActionPtrOutput {
	return o
}

func (o AutoHealCustomActionPtrOutput) Elem() AutoHealCustomActionOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) AutoHealCustomAction {
		if v != nil {
			return *v
		}
		var ret AutoHealCustomAction
		return ret
	}).(AutoHealCustomActionOutput)
}

func (o AutoHealCustomActionPtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionPtrOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomAction) *string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionResponse struct {
	Exe        *string `pulumi:"exe"`
	Parameters *string `pulumi:"parameters"`
}

type AutoHealCustomActionResponseOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomActionResponse)(nil)).Elem()
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponseOutput() AutoHealCustomActionResponseOutput {
	return o
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponseOutputWithContext(ctx context.Context) AutoHealCustomActionResponseOutput {
	return o
}

func (o AutoHealCustomActionResponseOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomActionResponse) *string { return v.Exe }).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionResponseOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AutoHealCustomActionResponse) *string { return v.Parameters }).(pulumi.StringPtrOutput)
}

type AutoHealCustomActionResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealCustomActionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomActionResponse)(nil)).Elem()
}

func (o AutoHealCustomActionResponsePtrOutput) ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput {
	return o
}

func (o AutoHealCustomActionResponsePtrOutput) ToAutoHealCustomActionResponsePtrOutputWithContext(ctx context.Context) AutoHealCustomActionResponsePtrOutput {
	return o
}

func (o AutoHealCustomActionResponsePtrOutput) Elem() AutoHealCustomActionResponseOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) AutoHealCustomActionResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealCustomActionResponse
		return ret
	}).(AutoHealCustomActionResponseOutput)
}

func (o AutoHealCustomActionResponsePtrOutput) Exe() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Exe
	}).(pulumi.StringPtrOutput)
}

func (o AutoHealCustomActionResponsePtrOutput) Parameters() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AutoHealCustomActionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringPtrOutput)
}

type AutoHealRules struct {
	Actions  *AutoHealActions  `pulumi:"actions"`
	Triggers *AutoHealTriggers `pulumi:"triggers"`
}





type AutoHealRulesInput interface {
	pulumi.Input

	ToAutoHealRulesOutput() AutoHealRulesOutput
	ToAutoHealRulesOutputWithContext(context.Context) AutoHealRulesOutput
}

type AutoHealRulesArgs struct {
	Actions  AutoHealActionsPtrInput  `pulumi:"actions"`
	Triggers AutoHealTriggersPtrInput `pulumi:"triggers"`
}

func (AutoHealRulesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRules)(nil)).Elem()
}

func (i AutoHealRulesArgs) ToAutoHealRulesOutput() AutoHealRulesOutput {
	return i.ToAutoHealRulesOutputWithContext(context.Background())
}

func (i AutoHealRulesArgs) ToAutoHealRulesOutputWithContext(ctx context.Context) AutoHealRulesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesOutput)
}

func (i AutoHealRulesArgs) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return i.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (i AutoHealRulesArgs) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesOutput).ToAutoHealRulesPtrOutputWithContext(ctx)
}









type AutoHealRulesPtrInput interface {
	pulumi.Input

	ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput
	ToAutoHealRulesPtrOutputWithContext(context.Context) AutoHealRulesPtrOutput
}

type autoHealRulesPtrType AutoHealRulesArgs

func AutoHealRulesPtr(v *AutoHealRulesArgs) AutoHealRulesPtrInput {
	return (*autoHealRulesPtrType)(v)
}

func (*autoHealRulesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRules)(nil)).Elem()
}

func (i *autoHealRulesPtrType) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return i.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (i *autoHealRulesPtrType) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesPtrOutput)
}

type AutoHealRulesOutput struct{ *pulumi.OutputState }

func (AutoHealRulesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRules)(nil)).Elem()
}

func (o AutoHealRulesOutput) ToAutoHealRulesOutput() AutoHealRulesOutput {
	return o
}

func (o AutoHealRulesOutput) ToAutoHealRulesOutputWithContext(ctx context.Context) AutoHealRulesOutput {
	return o
}

func (o AutoHealRulesOutput) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return o.ToAutoHealRulesPtrOutputWithContext(context.Background())
}

func (o AutoHealRulesOutput) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealRules) *AutoHealRules {
		return &v
	}).(AutoHealRulesPtrOutput)
}

func (o AutoHealRulesOutput) Actions() AutoHealActionsPtrOutput {
	return o.ApplyT(func(v AutoHealRules) *AutoHealActions { return v.Actions }).(AutoHealActionsPtrOutput)
}

func (o AutoHealRulesOutput) Triggers() AutoHealTriggersPtrOutput {
	return o.ApplyT(func(v AutoHealRules) *AutoHealTriggers { return v.Triggers }).(AutoHealTriggersPtrOutput)
}

type AutoHealRulesPtrOutput struct{ *pulumi.OutputState }

func (AutoHealRulesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRules)(nil)).Elem()
}

func (o AutoHealRulesPtrOutput) ToAutoHealRulesPtrOutput() AutoHealRulesPtrOutput {
	return o
}

func (o AutoHealRulesPtrOutput) ToAutoHealRulesPtrOutputWithContext(ctx context.Context) AutoHealRulesPtrOutput {
	return o
}

func (o AutoHealRulesPtrOutput) Elem() AutoHealRulesOutput {
	return o.ApplyT(func(v *AutoHealRules) AutoHealRules {
		if v != nil {
			return *v
		}
		var ret AutoHealRules
		return ret
	}).(AutoHealRulesOutput)
}

func (o AutoHealRulesPtrOutput) Actions() AutoHealActionsPtrOutput {
	return o.ApplyT(func(v *AutoHealRules) *AutoHealActions {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(AutoHealActionsPtrOutput)
}

func (o AutoHealRulesPtrOutput) Triggers() AutoHealTriggersPtrOutput {
	return o.ApplyT(func(v *AutoHealRules) *AutoHealTriggers {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(AutoHealTriggersPtrOutput)
}

type AutoHealRulesResponse struct {
	Actions  *AutoHealActionsResponse  `pulumi:"actions"`
	Triggers *AutoHealTriggersResponse `pulumi:"triggers"`
}

type AutoHealRulesResponseOutput struct{ *pulumi.OutputState }

func (AutoHealRulesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRulesResponse)(nil)).Elem()
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponseOutput() AutoHealRulesResponseOutput {
	return o
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponseOutputWithContext(ctx context.Context) AutoHealRulesResponseOutput {
	return o
}

func (o AutoHealRulesResponseOutput) Actions() AutoHealActionsResponsePtrOutput {
	return o.ApplyT(func(v AutoHealRulesResponse) *AutoHealActionsResponse { return v.Actions }).(AutoHealActionsResponsePtrOutput)
}

func (o AutoHealRulesResponseOutput) Triggers() AutoHealTriggersResponsePtrOutput {
	return o.ApplyT(func(v AutoHealRulesResponse) *AutoHealTriggersResponse { return v.Triggers }).(AutoHealTriggersResponsePtrOutput)
}

type AutoHealRulesResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealRulesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRulesResponse)(nil)).Elem()
}

func (o AutoHealRulesResponsePtrOutput) ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput {
	return o
}

func (o AutoHealRulesResponsePtrOutput) ToAutoHealRulesResponsePtrOutputWithContext(ctx context.Context) AutoHealRulesResponsePtrOutput {
	return o
}

func (o AutoHealRulesResponsePtrOutput) Elem() AutoHealRulesResponseOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) AutoHealRulesResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealRulesResponse
		return ret
	}).(AutoHealRulesResponseOutput)
}

func (o AutoHealRulesResponsePtrOutput) Actions() AutoHealActionsResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) *AutoHealActionsResponse {
		if v == nil {
			return nil
		}
		return v.Actions
	}).(AutoHealActionsResponsePtrOutput)
}

func (o AutoHealRulesResponsePtrOutput) Triggers() AutoHealTriggersResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealRulesResponse) *AutoHealTriggersResponse {
		if v == nil {
			return nil
		}
		return v.Triggers
	}).(AutoHealTriggersResponsePtrOutput)
}

type AutoHealTriggers struct {
	PrivateBytesInKB     *int                           `pulumi:"privateBytesInKB"`
	Requests             *RequestsBasedTrigger          `pulumi:"requests"`
	SlowRequests         *SlowRequestsBasedTrigger      `pulumi:"slowRequests"`
	SlowRequestsWithPath []SlowRequestsBasedTrigger     `pulumi:"slowRequestsWithPath"`
	StatusCodes          []StatusCodesBasedTrigger      `pulumi:"statusCodes"`
	StatusCodesRange     []StatusCodesRangeBasedTrigger `pulumi:"statusCodesRange"`
}





type AutoHealTriggersInput interface {
	pulumi.Input

	ToAutoHealTriggersOutput() AutoHealTriggersOutput
	ToAutoHealTriggersOutputWithContext(context.Context) AutoHealTriggersOutput
}

type AutoHealTriggersArgs struct {
	PrivateBytesInKB     pulumi.IntPtrInput                     `pulumi:"privateBytesInKB"`
	Requests             RequestsBasedTriggerPtrInput           `pulumi:"requests"`
	SlowRequests         SlowRequestsBasedTriggerPtrInput       `pulumi:"slowRequests"`
	SlowRequestsWithPath SlowRequestsBasedTriggerArrayInput     `pulumi:"slowRequestsWithPath"`
	StatusCodes          StatusCodesBasedTriggerArrayInput      `pulumi:"statusCodes"`
	StatusCodesRange     StatusCodesRangeBasedTriggerArrayInput `pulumi:"statusCodesRange"`
}

func (AutoHealTriggersArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggers)(nil)).Elem()
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersOutput() AutoHealTriggersOutput {
	return i.ToAutoHealTriggersOutputWithContext(context.Background())
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersOutputWithContext(ctx context.Context) AutoHealTriggersOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersOutput)
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return i.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (i AutoHealTriggersArgs) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersOutput).ToAutoHealTriggersPtrOutputWithContext(ctx)
}









type AutoHealTriggersPtrInput interface {
	pulumi.Input

	ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput
	ToAutoHealTriggersPtrOutputWithContext(context.Context) AutoHealTriggersPtrOutput
}

type autoHealTriggersPtrType AutoHealTriggersArgs

func AutoHealTriggersPtr(v *AutoHealTriggersArgs) AutoHealTriggersPtrInput {
	return (*autoHealTriggersPtrType)(v)
}

func (*autoHealTriggersPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggers)(nil)).Elem()
}

func (i *autoHealTriggersPtrType) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return i.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (i *autoHealTriggersPtrType) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersPtrOutput)
}

type AutoHealTriggersOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggers)(nil)).Elem()
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersOutput() AutoHealTriggersOutput {
	return o
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersOutputWithContext(ctx context.Context) AutoHealTriggersOutput {
	return o
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return o.ToAutoHealTriggersPtrOutputWithContext(context.Background())
}

func (o AutoHealTriggersOutput) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealTriggers) *AutoHealTriggers {
		return &v
	}).(AutoHealTriggersPtrOutput)
}

func (o AutoHealTriggersOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *int { return v.PrivateBytesInKB }).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersOutput) Requests() RequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *RequestsBasedTrigger { return v.Requests }).(RequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersOutput) SlowRequests() SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v AutoHealTriggers) *SlowRequestsBasedTrigger { return v.SlowRequests }).(SlowRequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []SlowRequestsBasedTrigger { return v.SlowRequestsWithPath }).(SlowRequestsBasedTriggerArrayOutput)
}

func (o AutoHealTriggersOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []StatusCodesBasedTrigger { return v.StatusCodes }).(StatusCodesBasedTriggerArrayOutput)
}

func (o AutoHealTriggersOutput) StatusCodesRange() StatusCodesRangeBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []StatusCodesRangeBasedTrigger { return v.StatusCodesRange }).(StatusCodesRangeBasedTriggerArrayOutput)
}

type AutoHealTriggersPtrOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggers)(nil)).Elem()
}

func (o AutoHealTriggersPtrOutput) ToAutoHealTriggersPtrOutput() AutoHealTriggersPtrOutput {
	return o
}

func (o AutoHealTriggersPtrOutput) ToAutoHealTriggersPtrOutputWithContext(ctx context.Context) AutoHealTriggersPtrOutput {
	return o
}

func (o AutoHealTriggersPtrOutput) Elem() AutoHealTriggersOutput {
	return o.ApplyT(func(v *AutoHealTriggers) AutoHealTriggers {
		if v != nil {
			return *v
		}
		var ret AutoHealTriggers
		return ret
	}).(AutoHealTriggersOutput)
}

func (o AutoHealTriggersPtrOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *int {
		if v == nil {
			return nil
		}
		return v.PrivateBytesInKB
	}).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersPtrOutput) Requests() RequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *RequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(RequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersPtrOutput) SlowRequests() SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggers) *SlowRequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.SlowRequests
	}).(SlowRequestsBasedTriggerPtrOutput)
}

func (o AutoHealTriggersPtrOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []SlowRequestsBasedTrigger {
		if v == nil {
			return nil
		}
		return v.SlowRequestsWithPath
	}).(SlowRequestsBasedTriggerArrayOutput)
}

func (o AutoHealTriggersPtrOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []StatusCodesBasedTrigger {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerArrayOutput)
}

func (o AutoHealTriggersPtrOutput) StatusCodesRange() StatusCodesRangeBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []StatusCodesRangeBasedTrigger {
		if v == nil {
			return nil
		}
		return v.StatusCodesRange
	}).(StatusCodesRangeBasedTriggerArrayOutput)
}

type AutoHealTriggersResponse struct {
	PrivateBytesInKB     *int                                   `pulumi:"privateBytesInKB"`
	Requests             *RequestsBasedTriggerResponse          `pulumi:"requests"`
	SlowRequests         *SlowRequestsBasedTriggerResponse      `pulumi:"slowRequests"`
	SlowRequestsWithPath []SlowRequestsBasedTriggerResponse     `pulumi:"slowRequestsWithPath"`
	StatusCodes          []StatusCodesBasedTriggerResponse      `pulumi:"statusCodes"`
	StatusCodesRange     []StatusCodesRangeBasedTriggerResponse `pulumi:"statusCodesRange"`
}

type AutoHealTriggersResponseOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggersResponse)(nil)).Elem()
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponseOutput() AutoHealTriggersResponseOutput {
	return o
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponseOutputWithContext(ctx context.Context) AutoHealTriggersResponseOutput {
	return o
}

func (o AutoHealTriggersResponseOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *int { return v.PrivateBytesInKB }).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersResponseOutput) Requests() RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *RequestsBasedTriggerResponse { return v.Requests }).(RequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponseOutput) SlowRequests() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) *SlowRequestsBasedTriggerResponse { return v.SlowRequests }).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponseOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []SlowRequestsBasedTriggerResponse { return v.SlowRequestsWithPath }).(SlowRequestsBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponseOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse { return v.StatusCodes }).(StatusCodesBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponseOutput) StatusCodesRange() StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []StatusCodesRangeBasedTriggerResponse { return v.StatusCodesRange }).(StatusCodesRangeBasedTriggerResponseArrayOutput)
}

type AutoHealTriggersResponsePtrOutput struct{ *pulumi.OutputState }

func (AutoHealTriggersResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggersResponse)(nil)).Elem()
}

func (o AutoHealTriggersResponsePtrOutput) ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput {
	return o
}

func (o AutoHealTriggersResponsePtrOutput) ToAutoHealTriggersResponsePtrOutputWithContext(ctx context.Context) AutoHealTriggersResponsePtrOutput {
	return o
}

func (o AutoHealTriggersResponsePtrOutput) Elem() AutoHealTriggersResponseOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) AutoHealTriggersResponse {
		if v != nil {
			return *v
		}
		var ret AutoHealTriggersResponse
		return ret
	}).(AutoHealTriggersResponseOutput)
}

func (o AutoHealTriggersResponsePtrOutput) PrivateBytesInKB() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *int {
		if v == nil {
			return nil
		}
		return v.PrivateBytesInKB
	}).(pulumi.IntPtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) Requests() RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *RequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.Requests
	}).(RequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) SlowRequests() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) *SlowRequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.SlowRequests
	}).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o AutoHealTriggersResponsePtrOutput) SlowRequestsWithPath() SlowRequestsBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []SlowRequestsBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.SlowRequestsWithPath
	}).(SlowRequestsBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponsePtrOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerResponseArrayOutput)
}

func (o AutoHealTriggersResponsePtrOutput) StatusCodesRange() StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []StatusCodesRangeBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.StatusCodesRange
	}).(StatusCodesRangeBasedTriggerResponseArrayOutput)
}

type AzureBlobStorageApplicationLogsConfig struct {
	Level           *LogLevel `pulumi:"level"`
	RetentionInDays *int      `pulumi:"retentionInDays"`
	SasUrl          *string   `pulumi:"sasUrl"`
}





type AzureBlobStorageApplicationLogsConfigInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput
	ToAzureBlobStorageApplicationLogsConfigOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigOutput
}

type AzureBlobStorageApplicationLogsConfigArgs struct {
	Level           LogLevelPtrInput      `pulumi:"level"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigOutput)
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigArgs) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigOutput).ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx)
}









type AzureBlobStorageApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput
	ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput
}

type azureBlobStorageApplicationLogsConfigPtrType AzureBlobStorageApplicationLogsConfigArgs

func AzureBlobStorageApplicationLogsConfigPtr(v *AzureBlobStorageApplicationLogsConfigArgs) AzureBlobStorageApplicationLogsConfigPtrInput {
	return (*azureBlobStorageApplicationLogsConfigPtrType)(v)
}

func (*azureBlobStorageApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (i *azureBlobStorageApplicationLogsConfigPtrType) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageApplicationLogsConfigPtrType) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigOutput() AzureBlobStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageApplicationLogsConfigOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageApplicationLogsConfig) *AzureBlobStorageApplicationLogsConfig {
		return &v
	}).(AzureBlobStorageApplicationLogsConfigPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutput() AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) ToAzureBlobStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) Elem() AzureBlobStorageApplicationLogsConfigOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) AzureBlobStorageApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageApplicationLogsConfig
		return ret
	}).(AzureBlobStorageApplicationLogsConfigOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigResponse struct {
	Level           *string `pulumi:"level"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}

type AzureBlobStorageApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponseOutput() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageApplicationLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) Elem() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) AzureBlobStorageApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageApplicationLogsConfigResponse
		return ret
	}).(AzureBlobStorageApplicationLogsConfigResponseOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageApplicationLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfig struct {
	Enabled         *bool   `pulumi:"enabled"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}





type AzureBlobStorageHttpLogsConfigInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput
	ToAzureBlobStorageHttpLogsConfigOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigOutput
}

type AzureBlobStorageHttpLogsConfigArgs struct {
	Enabled         pulumi.BoolPtrInput   `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageHttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput {
	return i.ToAzureBlobStorageHttpLogsConfigOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigOutput)
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigArgs) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigOutput).ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx)
}









type AzureBlobStorageHttpLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput
	ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigPtrOutput
}

type azureBlobStorageHttpLogsConfigPtrType AzureBlobStorageHttpLogsConfigArgs

func AzureBlobStorageHttpLogsConfigPtr(v *AzureBlobStorageHttpLogsConfigArgs) AzureBlobStorageHttpLogsConfigPtrInput {
	return (*azureBlobStorageHttpLogsConfigPtrType)(v)
}

func (*azureBlobStorageHttpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (i *azureBlobStorageHttpLogsConfigPtrType) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageHttpLogsConfigPtrType) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

type AzureBlobStorageHttpLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigOutput() AzureBlobStorageHttpLogsConfigOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageHttpLogsConfigOutput) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageHttpLogsConfig) *AzureBlobStorageHttpLogsConfig {
		return &v
	}).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfig) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfig)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) ToAzureBlobStorageHttpLogsConfigPtrOutput() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) ToAzureBlobStorageHttpLogsConfigPtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigPtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) Elem() AzureBlobStorageHttpLogsConfigOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) AzureBlobStorageHttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageHttpLogsConfig
		return ret
	}).(AzureBlobStorageHttpLogsConfigOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfig) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigResponse struct {
	Enabled         *bool   `pulumi:"enabled"`
	RetentionInDays *int    `pulumi:"retentionInDays"`
	SasUrl          *string `pulumi:"sasUrl"`
}

type AzureBlobStorageHttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponseOutput() AzureBlobStorageHttpLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponseOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureBlobStorageHttpLogsConfigResponse) *string { return v.SasUrl }).(pulumi.StringPtrOutput)
}

type AzureBlobStorageHttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureBlobStorageHttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) Elem() AzureBlobStorageHttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) AzureBlobStorageHttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureBlobStorageHttpLogsConfigResponse
		return ret
	}).(AzureBlobStorageHttpLogsConfigResponseOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o AzureBlobStorageHttpLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureBlobStorageHttpLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureResourceErrorInfoResponse struct {
	Code    string                           `pulumi:"code"`
	Details []AzureResourceErrorInfoResponse `pulumi:"details"`
	Message string                           `pulumi:"message"`
}

type AzureResourceErrorInfoResponseOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponseOutput) ToAzureResourceErrorInfoResponseOutput() AzureResourceErrorInfoResponseOutput {
	return o
}

func (o AzureResourceErrorInfoResponseOutput) ToAzureResourceErrorInfoResponseOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponseOutput {
	return o
}

func (o AzureResourceErrorInfoResponseOutput) Code() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) string { return v.Code }).(pulumi.StringOutput)
}

func (o AzureResourceErrorInfoResponseOutput) Details() AzureResourceErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) []AzureResourceErrorInfoResponse { return v.Details }).(AzureResourceErrorInfoResponseArrayOutput)
}

func (o AzureResourceErrorInfoResponseOutput) Message() pulumi.StringOutput {
	return o.ApplyT(func(v AzureResourceErrorInfoResponse) string { return v.Message }).(pulumi.StringOutput)
}

type AzureResourceErrorInfoResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponsePtrOutput) ToAzureResourceErrorInfoResponsePtrOutput() AzureResourceErrorInfoResponsePtrOutput {
	return o
}

func (o AzureResourceErrorInfoResponsePtrOutput) ToAzureResourceErrorInfoResponsePtrOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponsePtrOutput {
	return o
}

func (o AzureResourceErrorInfoResponsePtrOutput) Elem() AzureResourceErrorInfoResponseOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) AzureResourceErrorInfoResponse {
		if v != nil {
			return *v
		}
		var ret AzureResourceErrorInfoResponse
		return ret
	}).(AzureResourceErrorInfoResponseOutput)
}

func (o AzureResourceErrorInfoResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Code
	}).(pulumi.StringPtrOutput)
}

func (o AzureResourceErrorInfoResponsePtrOutput) Details() AzureResourceErrorInfoResponseArrayOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) []AzureResourceErrorInfoResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(AzureResourceErrorInfoResponseArrayOutput)
}

func (o AzureResourceErrorInfoResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureResourceErrorInfoResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Message
	}).(pulumi.StringPtrOutput)
}

type AzureResourceErrorInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (AzureResourceErrorInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]AzureResourceErrorInfoResponse)(nil)).Elem()
}

func (o AzureResourceErrorInfoResponseArrayOutput) ToAzureResourceErrorInfoResponseArrayOutput() AzureResourceErrorInfoResponseArrayOutput {
	return o
}

func (o AzureResourceErrorInfoResponseArrayOutput) ToAzureResourceErrorInfoResponseArrayOutputWithContext(ctx context.Context) AzureResourceErrorInfoResponseArrayOutput {
	return o
}

func (o AzureResourceErrorInfoResponseArrayOutput) Index(i pulumi.IntInput) AzureResourceErrorInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) AzureResourceErrorInfoResponse {
		return vs[0].([]AzureResourceErrorInfoResponse)[vs[1].(int)]
	}).(AzureResourceErrorInfoResponseOutput)
}

type AzureStorageInfoValue struct {
	AccessKey   *string           `pulumi:"accessKey"`
	AccountName *string           `pulumi:"accountName"`
	MountPath   *string           `pulumi:"mountPath"`
	ShareName   *string           `pulumi:"shareName"`
	Type        *AzureStorageType `pulumi:"type"`
}





type AzureStorageInfoValueInput interface {
	pulumi.Input

	ToAzureStorageInfoValueOutput() AzureStorageInfoValueOutput
	ToAzureStorageInfoValueOutputWithContext(context.Context) AzureStorageInfoValueOutput
}

type AzureStorageInfoValueArgs struct {
	AccessKey   pulumi.StringPtrInput    `pulumi:"accessKey"`
	AccountName pulumi.StringPtrInput    `pulumi:"accountName"`
	MountPath   pulumi.StringPtrInput    `pulumi:"mountPath"`
	ShareName   pulumi.StringPtrInput    `pulumi:"shareName"`
	Type        AzureStorageTypePtrInput `pulumi:"type"`
}

func (AzureStorageInfoValueArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValue)(nil)).Elem()
}

func (i AzureStorageInfoValueArgs) ToAzureStorageInfoValueOutput() AzureStorageInfoValueOutput {
	return i.ToAzureStorageInfoValueOutputWithContext(context.Background())
}

func (i AzureStorageInfoValueArgs) ToAzureStorageInfoValueOutputWithContext(ctx context.Context) AzureStorageInfoValueOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageInfoValueOutput)
}





type AzureStorageInfoValueMapInput interface {
	pulumi.Input

	ToAzureStorageInfoValueMapOutput() AzureStorageInfoValueMapOutput
	ToAzureStorageInfoValueMapOutputWithContext(context.Context) AzureStorageInfoValueMapOutput
}

type AzureStorageInfoValueMap map[string]AzureStorageInfoValueInput

func (AzureStorageInfoValueMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValue)(nil)).Elem()
}

func (i AzureStorageInfoValueMap) ToAzureStorageInfoValueMapOutput() AzureStorageInfoValueMapOutput {
	return i.ToAzureStorageInfoValueMapOutputWithContext(context.Background())
}

func (i AzureStorageInfoValueMap) ToAzureStorageInfoValueMapOutputWithContext(ctx context.Context) AzureStorageInfoValueMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageInfoValueMapOutput)
}

type AzureStorageInfoValueOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValue)(nil)).Elem()
}

func (o AzureStorageInfoValueOutput) ToAzureStorageInfoValueOutput() AzureStorageInfoValueOutput {
	return o
}

func (o AzureStorageInfoValueOutput) ToAzureStorageInfoValueOutputWithContext(ctx context.Context) AzureStorageInfoValueOutput {
	return o
}

func (o AzureStorageInfoValueOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *string { return v.ShareName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueOutput) Type() AzureStorageTypePtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValue) *AzureStorageType { return v.Type }).(AzureStorageTypePtrOutput)
}

type AzureStorageInfoValueMapOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValue)(nil)).Elem()
}

func (o AzureStorageInfoValueMapOutput) ToAzureStorageInfoValueMapOutput() AzureStorageInfoValueMapOutput {
	return o
}

func (o AzureStorageInfoValueMapOutput) ToAzureStorageInfoValueMapOutputWithContext(ctx context.Context) AzureStorageInfoValueMapOutput {
	return o
}

func (o AzureStorageInfoValueMapOutput) MapIndex(k pulumi.StringInput) AzureStorageInfoValueOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AzureStorageInfoValue {
		return vs[0].(map[string]AzureStorageInfoValue)[vs[1].(string)]
	}).(AzureStorageInfoValueOutput)
}

type AzureStorageInfoValueResponse struct {
	AccessKey   *string `pulumi:"accessKey"`
	AccountName *string `pulumi:"accountName"`
	MountPath   *string `pulumi:"mountPath"`
	ShareName   *string `pulumi:"shareName"`
	State       string  `pulumi:"state"`
	Type        *string `pulumi:"type"`
}

type AzureStorageInfoValueResponseOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValueResponse)(nil)).Elem()
}

func (o AzureStorageInfoValueResponseOutput) ToAzureStorageInfoValueResponseOutput() AzureStorageInfoValueResponseOutput {
	return o
}

func (o AzureStorageInfoValueResponseOutput) ToAzureStorageInfoValueResponseOutputWithContext(ctx context.Context) AzureStorageInfoValueResponseOutput {
	return o
}

func (o AzureStorageInfoValueResponseOutput) AccessKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.AccessKey }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) AccountName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.AccountName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) MountPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.MountPath }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) ShareName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.ShareName }).(pulumi.StringPtrOutput)
}

func (o AzureStorageInfoValueResponseOutput) State() pulumi.StringOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) string { return v.State }).(pulumi.StringOutput)
}

func (o AzureStorageInfoValueResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureStorageInfoValueResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type AzureStorageInfoValueResponseMapOutput struct{ *pulumi.OutputState }

func (AzureStorageInfoValueResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValueResponse)(nil)).Elem()
}

func (o AzureStorageInfoValueResponseMapOutput) ToAzureStorageInfoValueResponseMapOutput() AzureStorageInfoValueResponseMapOutput {
	return o
}

func (o AzureStorageInfoValueResponseMapOutput) ToAzureStorageInfoValueResponseMapOutputWithContext(ctx context.Context) AzureStorageInfoValueResponseMapOutput {
	return o
}

func (o AzureStorageInfoValueResponseMapOutput) MapIndex(k pulumi.StringInput) AzureStorageInfoValueResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) AzureStorageInfoValueResponse {
		return vs[0].(map[string]AzureStorageInfoValueResponse)[vs[1].(string)]
	}).(AzureStorageInfoValueResponseOutput)
}

type AzureTableStorageApplicationLogsConfig struct {
	Level  *LogLevel `pulumi:"level"`
	SasUrl string    `pulumi:"sasUrl"`
}





type AzureTableStorageApplicationLogsConfigInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput
	ToAzureTableStorageApplicationLogsConfigOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigOutput
}

type AzureTableStorageApplicationLogsConfigArgs struct {
	Level  LogLevelPtrInput   `pulumi:"level"`
	SasUrl pulumi.StringInput `pulumi:"sasUrl"`
}

func (AzureTableStorageApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput {
	return i.ToAzureTableStorageApplicationLogsConfigOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigOutput)
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigArgs) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigOutput).ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx)
}









type AzureTableStorageApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput
	ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigPtrOutput
}

type azureTableStorageApplicationLogsConfigPtrType AzureTableStorageApplicationLogsConfigArgs

func AzureTableStorageApplicationLogsConfigPtr(v *AzureTableStorageApplicationLogsConfigArgs) AzureTableStorageApplicationLogsConfigPtrInput {
	return (*azureTableStorageApplicationLogsConfigPtrType)(v)
}

func (*azureTableStorageApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (i *azureTableStorageApplicationLogsConfigPtrType) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *azureTableStorageApplicationLogsConfigPtrType) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

type AzureTableStorageApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigOutput() AzureTableStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o AzureTableStorageApplicationLogsConfigOutput) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureTableStorageApplicationLogsConfig) *AzureTableStorageApplicationLogsConfig {
		return &v
	}).(AzureTableStorageApplicationLogsConfigPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigOutput) SasUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfig) string { return v.SasUrl }).(pulumi.StringOutput)
}

type AzureTableStorageApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfig)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) ToAzureTableStorageApplicationLogsConfigPtrOutput() AzureTableStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) ToAzureTableStorageApplicationLogsConfigPtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigPtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) Elem() AzureTableStorageApplicationLogsConfigOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) AzureTableStorageApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret AzureTableStorageApplicationLogsConfig
		return ret
	}).(AzureTableStorageApplicationLogsConfigOutput)
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigPtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfig) *string {
		if v == nil {
			return nil
		}
		return &v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type AzureTableStorageApplicationLogsConfigResponse struct {
	Level  *string `pulumi:"level"`
	SasUrl string  `pulumi:"sasUrl"`
}

type AzureTableStorageApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponseOutput() AzureTableStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponseOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) SasUrl() pulumi.StringOutput {
	return o.ApplyT(func(v AzureTableStorageApplicationLogsConfigResponse) string { return v.SasUrl }).(pulumi.StringOutput)
}

type AzureTableStorageApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (AzureTableStorageApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) Elem() AzureTableStorageApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) AzureTableStorageApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret AzureTableStorageApplicationLogsConfigResponse
		return ret
	}).(AzureTableStorageApplicationLogsConfigResponseOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

func (o AzureTableStorageApplicationLogsConfigResponsePtrOutput) SasUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AzureTableStorageApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SasUrl
	}).(pulumi.StringPtrOutput)
}

type BackupItemResponse struct {
	BackupId             int                             `pulumi:"backupId"`
	BlobName             string                          `pulumi:"blobName"`
	CorrelationId        string                          `pulumi:"correlationId"`
	Created              string                          `pulumi:"created"`
	Databases            []DatabaseBackupSettingResponse `pulumi:"databases"`
	FinishedTimeStamp    string                          `pulumi:"finishedTimeStamp"`
	Id                   string                          `pulumi:"id"`
	Kind                 *string                         `pulumi:"kind"`
	LastRestoreTimeStamp string                          `pulumi:"lastRestoreTimeStamp"`
	Log                  string                          `pulumi:"log"`
	Name                 string                          `pulumi:"name"`
	Scheduled            bool                            `pulumi:"scheduled"`
	SizeInBytes          float64                         `pulumi:"sizeInBytes"`
	Status               string                          `pulumi:"status"`
	StorageAccountUrl    string                          `pulumi:"storageAccountUrl"`
	Type                 string                          `pulumi:"type"`
	WebsiteSizeInBytes   float64                         `pulumi:"websiteSizeInBytes"`
}

type BackupItemResponseOutput struct{ *pulumi.OutputState }

func (BackupItemResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupItemResponse)(nil)).Elem()
}

func (o BackupItemResponseOutput) ToBackupItemResponseOutput() BackupItemResponseOutput {
	return o
}

func (o BackupItemResponseOutput) ToBackupItemResponseOutputWithContext(ctx context.Context) BackupItemResponseOutput {
	return o
}

func (o BackupItemResponseOutput) BackupId() pulumi.IntOutput {
	return o.ApplyT(func(v BackupItemResponse) int { return v.BackupId }).(pulumi.IntOutput)
}

func (o BackupItemResponseOutput) BlobName() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.BlobName }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) CorrelationId() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.CorrelationId }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Created() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Created }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Databases() DatabaseBackupSettingResponseArrayOutput {
	return o.ApplyT(func(v BackupItemResponse) []DatabaseBackupSettingResponse { return v.Databases }).(DatabaseBackupSettingResponseArrayOutput)
}

func (o BackupItemResponseOutput) FinishedTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.FinishedTimeStamp }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupItemResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o BackupItemResponseOutput) LastRestoreTimeStamp() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.LastRestoreTimeStamp }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Log() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Log }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Scheduled() pulumi.BoolOutput {
	return o.ApplyT(func(v BackupItemResponse) bool { return v.Scheduled }).(pulumi.BoolOutput)
}

func (o BackupItemResponseOutput) SizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v BackupItemResponse) float64 { return v.SizeInBytes }).(pulumi.Float64Output)
}

func (o BackupItemResponseOutput) Status() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Status }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) StorageAccountUrl() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.StorageAccountUrl }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v BackupItemResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o BackupItemResponseOutput) WebsiteSizeInBytes() pulumi.Float64Output {
	return o.ApplyT(func(v BackupItemResponse) float64 { return v.WebsiteSizeInBytes }).(pulumi.Float64Output)
}

type BackupItemResponseArrayOutput struct{ *pulumi.OutputState }

func (BackupItemResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupItemResponse)(nil)).Elem()
}

func (o BackupItemResponseArrayOutput) ToBackupItemResponseArrayOutput() BackupItemResponseArrayOutput {
	return o
}

func (o BackupItemResponseArrayOutput) ToBackupItemResponseArrayOutputWithContext(ctx context.Context) BackupItemResponseArrayOutput {
	return o
}

func (o BackupItemResponseArrayOutput) Index(i pulumi.IntInput) BackupItemResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) BackupItemResponse {
		return vs[0].([]BackupItemResponse)[vs[1].(int)]
	}).(BackupItemResponseOutput)
}

type BackupSchedule struct {
	FrequencyInterval     int           `pulumi:"frequencyInterval"`
	FrequencyUnit         FrequencyUnit `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  bool          `pulumi:"keepAtLeastOneBackup"`
	RetentionPeriodInDays int           `pulumi:"retentionPeriodInDays"`
	StartTime             *string       `pulumi:"startTime"`
}


func (val *BackupSchedule) Defaults() *BackupSchedule {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FrequencyInterval) {
		tmp.FrequencyInterval = 7
	}
	if isZero(tmp.FrequencyUnit) {
		tmp.FrequencyUnit = FrequencyUnit("Day")
	}
	if isZero(tmp.KeepAtLeastOneBackup) {
		tmp.KeepAtLeastOneBackup = true
	}
	if isZero(tmp.RetentionPeriodInDays) {
		tmp.RetentionPeriodInDays = 30
	}
	return &tmp
}





type BackupScheduleInput interface {
	pulumi.Input

	ToBackupScheduleOutput() BackupScheduleOutput
	ToBackupScheduleOutputWithContext(context.Context) BackupScheduleOutput
}

type BackupScheduleArgs struct {
	FrequencyInterval     pulumi.IntInput       `pulumi:"frequencyInterval"`
	FrequencyUnit         FrequencyUnitInput    `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  pulumi.BoolInput      `pulumi:"keepAtLeastOneBackup"`
	RetentionPeriodInDays pulumi.IntInput       `pulumi:"retentionPeriodInDays"`
	StartTime             pulumi.StringPtrInput `pulumi:"startTime"`
}


func (val *BackupScheduleArgs) Defaults() *BackupScheduleArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FrequencyInterval) {
		tmp.FrequencyInterval = pulumi.Int(7)
	}
	if isZero(tmp.FrequencyUnit) {
		tmp.FrequencyUnit = FrequencyUnit("Day")
	}
	if isZero(tmp.KeepAtLeastOneBackup) {
		tmp.KeepAtLeastOneBackup = pulumi.Bool(true)
	}
	if isZero(tmp.RetentionPeriodInDays) {
		tmp.RetentionPeriodInDays = pulumi.Int(30)
	}
	return &tmp
}
func (BackupScheduleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupSchedule)(nil)).Elem()
}

func (i BackupScheduleArgs) ToBackupScheduleOutput() BackupScheduleOutput {
	return i.ToBackupScheduleOutputWithContext(context.Background())
}

func (i BackupScheduleArgs) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput)
}

func (i BackupScheduleArgs) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return i.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (i BackupScheduleArgs) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleOutput).ToBackupSchedulePtrOutputWithContext(ctx)
}









type BackupSchedulePtrInput interface {
	pulumi.Input

	ToBackupSchedulePtrOutput() BackupSchedulePtrOutput
	ToBackupSchedulePtrOutputWithContext(context.Context) BackupSchedulePtrOutput
}

type backupSchedulePtrType BackupScheduleArgs

func BackupSchedulePtr(v *BackupScheduleArgs) BackupSchedulePtrInput {
	return (*backupSchedulePtrType)(v)
}

func (*backupSchedulePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (i *backupSchedulePtrType) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return i.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (i *backupSchedulePtrType) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupSchedulePtrOutput)
}

type BackupScheduleOutput struct{ *pulumi.OutputState }

func (BackupScheduleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupSchedule)(nil)).Elem()
}

func (o BackupScheduleOutput) ToBackupScheduleOutput() BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupScheduleOutputWithContext(ctx context.Context) BackupScheduleOutput {
	return o
}

func (o BackupScheduleOutput) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return o.ToBackupSchedulePtrOutputWithContext(context.Background())
}

func (o BackupScheduleOutput) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupSchedule) *BackupSchedule {
		return &v
	}).(BackupSchedulePtrOutput)
}

func (o BackupScheduleOutput) FrequencyInterval() pulumi.IntOutput {
	return o.ApplyT(func(v BackupSchedule) int { return v.FrequencyInterval }).(pulumi.IntOutput)
}

func (o BackupScheduleOutput) FrequencyUnit() FrequencyUnitOutput {
	return o.ApplyT(func(v BackupSchedule) FrequencyUnit { return v.FrequencyUnit }).(FrequencyUnitOutput)
}

func (o BackupScheduleOutput) KeepAtLeastOneBackup() pulumi.BoolOutput {
	return o.ApplyT(func(v BackupSchedule) bool { return v.KeepAtLeastOneBackup }).(pulumi.BoolOutput)
}

func (o BackupScheduleOutput) RetentionPeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v BackupSchedule) int { return v.RetentionPeriodInDays }).(pulumi.IntOutput)
}

func (o BackupScheduleOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupSchedule) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type BackupSchedulePtrOutput struct{ *pulumi.OutputState }

func (BackupSchedulePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupSchedule)(nil)).Elem()
}

func (o BackupSchedulePtrOutput) ToBackupSchedulePtrOutput() BackupSchedulePtrOutput {
	return o
}

func (o BackupSchedulePtrOutput) ToBackupSchedulePtrOutputWithContext(ctx context.Context) BackupSchedulePtrOutput {
	return o
}

func (o BackupSchedulePtrOutput) Elem() BackupScheduleOutput {
	return o.ApplyT(func(v *BackupSchedule) BackupSchedule {
		if v != nil {
			return *v
		}
		var ret BackupSchedule
		return ret
	}).(BackupScheduleOutput)
}

func (o BackupSchedulePtrOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInterval
	}).(pulumi.IntPtrOutput)
}

func (o BackupSchedulePtrOutput) FrequencyUnit() FrequencyUnitPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *FrequencyUnit {
		if v == nil {
			return nil
		}
		return &v.FrequencyUnit
	}).(FrequencyUnitPtrOutput)
}

func (o BackupSchedulePtrOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepAtLeastOneBackup
	}).(pulumi.BoolPtrOutput)
}

func (o BackupSchedulePtrOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *int {
		if v == nil {
			return nil
		}
		return &v.RetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupSchedulePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupSchedule) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type BackupScheduleResponse struct {
	FrequencyInterval     int     `pulumi:"frequencyInterval"`
	FrequencyUnit         string  `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  bool    `pulumi:"keepAtLeastOneBackup"`
	LastExecutionTime     string  `pulumi:"lastExecutionTime"`
	RetentionPeriodInDays int     `pulumi:"retentionPeriodInDays"`
	StartTime             *string `pulumi:"startTime"`
}


func (val *BackupScheduleResponse) Defaults() *BackupScheduleResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.FrequencyInterval) {
		tmp.FrequencyInterval = 7
	}
	if isZero(tmp.FrequencyUnit) {
		tmp.FrequencyUnit = "Day"
	}
	if isZero(tmp.KeepAtLeastOneBackup) {
		tmp.KeepAtLeastOneBackup = true
	}
	if isZero(tmp.RetentionPeriodInDays) {
		tmp.RetentionPeriodInDays = 30
	}
	return &tmp
}

type BackupScheduleResponseOutput struct{ *pulumi.OutputState }

func (BackupScheduleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupScheduleResponse)(nil)).Elem()
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponseOutput() BackupScheduleResponseOutput {
	return o
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponseOutputWithContext(ctx context.Context) BackupScheduleResponseOutput {
	return o
}

func (o BackupScheduleResponseOutput) FrequencyInterval() pulumi.IntOutput {
	return o.ApplyT(func(v BackupScheduleResponse) int { return v.FrequencyInterval }).(pulumi.IntOutput)
}

func (o BackupScheduleResponseOutput) FrequencyUnit() pulumi.StringOutput {
	return o.ApplyT(func(v BackupScheduleResponse) string { return v.FrequencyUnit }).(pulumi.StringOutput)
}

func (o BackupScheduleResponseOutput) KeepAtLeastOneBackup() pulumi.BoolOutput {
	return o.ApplyT(func(v BackupScheduleResponse) bool { return v.KeepAtLeastOneBackup }).(pulumi.BoolOutput)
}

func (o BackupScheduleResponseOutput) LastExecutionTime() pulumi.StringOutput {
	return o.ApplyT(func(v BackupScheduleResponse) string { return v.LastExecutionTime }).(pulumi.StringOutput)
}

func (o BackupScheduleResponseOutput) RetentionPeriodInDays() pulumi.IntOutput {
	return o.ApplyT(func(v BackupScheduleResponse) int { return v.RetentionPeriodInDays }).(pulumi.IntOutput)
}

func (o BackupScheduleResponseOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BackupScheduleResponse) *string { return v.StartTime }).(pulumi.StringPtrOutput)
}

type BackupScheduleResponsePtrOutput struct{ *pulumi.OutputState }

func (BackupScheduleResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupScheduleResponse)(nil)).Elem()
}

func (o BackupScheduleResponsePtrOutput) ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput {
	return o
}

func (o BackupScheduleResponsePtrOutput) ToBackupScheduleResponsePtrOutputWithContext(ctx context.Context) BackupScheduleResponsePtrOutput {
	return o
}

func (o BackupScheduleResponsePtrOutput) Elem() BackupScheduleResponseOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) BackupScheduleResponse {
		if v != nil {
			return *v
		}
		var ret BackupScheduleResponse
		return ret
	}).(BackupScheduleResponseOutput)
}

func (o BackupScheduleResponsePtrOutput) FrequencyInterval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.FrequencyInterval
	}).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) FrequencyUnit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.FrequencyUnit
	}).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) KeepAtLeastOneBackup() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.KeepAtLeastOneBackup
	}).(pulumi.BoolPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) LastExecutionTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastExecutionTime
	}).(pulumi.StringPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) RetentionPeriodInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *int {
		if v == nil {
			return nil
		}
		return &v.RetentionPeriodInDays
	}).(pulumi.IntPtrOutput)
}

func (o BackupScheduleResponsePtrOutput) StartTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BackupScheduleResponse) *string {
		if v == nil {
			return nil
		}
		return v.StartTime
	}).(pulumi.StringPtrOutput)
}

type Capability struct {
	Name   *string `pulumi:"name"`
	Reason *string `pulumi:"reason"`
	Value  *string `pulumi:"value"`
}





type CapabilityInput interface {
	pulumi.Input

	ToCapabilityOutput() CapabilityOutput
	ToCapabilityOutputWithContext(context.Context) CapabilityOutput
}

type CapabilityArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	Value  pulumi.StringPtrInput `pulumi:"value"`
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

func (o CapabilityOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o CapabilityOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v Capability) *string { return v.Value }).(pulumi.StringPtrOutput)
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
	Name   *string `pulumi:"name"`
	Reason *string `pulumi:"reason"`
	Value  *string `pulumi:"value"`
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

func (o CapabilityResponseOutput) Reason() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Reason }).(pulumi.StringPtrOutput)
}

func (o CapabilityResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CapabilityResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
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

type CloningInfo struct {
	AppSettingsOverrides      map[string]string `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      *bool             `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        *bool             `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    *bool             `pulumi:"configureLoadBalancing"`
	CorrelationId             *string           `pulumi:"correlationId"`
	HostingEnvironment        *string           `pulumi:"hostingEnvironment"`
	Overwrite                 *bool             `pulumi:"overwrite"`
	SourceWebAppId            string            `pulumi:"sourceWebAppId"`
	SourceWebAppLocation      *string           `pulumi:"sourceWebAppLocation"`
	TrafficManagerProfileId   *string           `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName *string           `pulumi:"trafficManagerProfileName"`
}





type CloningInfoInput interface {
	pulumi.Input

	ToCloningInfoOutput() CloningInfoOutput
	ToCloningInfoOutputWithContext(context.Context) CloningInfoOutput
}

type CloningInfoArgs struct {
	AppSettingsOverrides      pulumi.StringMapInput `pulumi:"appSettingsOverrides"`
	CloneCustomHostNames      pulumi.BoolPtrInput   `pulumi:"cloneCustomHostNames"`
	CloneSourceControl        pulumi.BoolPtrInput   `pulumi:"cloneSourceControl"`
	ConfigureLoadBalancing    pulumi.BoolPtrInput   `pulumi:"configureLoadBalancing"`
	CorrelationId             pulumi.StringPtrInput `pulumi:"correlationId"`
	HostingEnvironment        pulumi.StringPtrInput `pulumi:"hostingEnvironment"`
	Overwrite                 pulumi.BoolPtrInput   `pulumi:"overwrite"`
	SourceWebAppId            pulumi.StringInput    `pulumi:"sourceWebAppId"`
	SourceWebAppLocation      pulumi.StringPtrInput `pulumi:"sourceWebAppLocation"`
	TrafficManagerProfileId   pulumi.StringPtrInput `pulumi:"trafficManagerProfileId"`
	TrafficManagerProfileName pulumi.StringPtrInput `pulumi:"trafficManagerProfileName"`
}

func (CloningInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfo)(nil)).Elem()
}

func (i CloningInfoArgs) ToCloningInfoOutput() CloningInfoOutput {
	return i.ToCloningInfoOutputWithContext(context.Background())
}

func (i CloningInfoArgs) ToCloningInfoOutputWithContext(ctx context.Context) CloningInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoOutput)
}

func (i CloningInfoArgs) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return i.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (i CloningInfoArgs) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoOutput).ToCloningInfoPtrOutputWithContext(ctx)
}









type CloningInfoPtrInput interface {
	pulumi.Input

	ToCloningInfoPtrOutput() CloningInfoPtrOutput
	ToCloningInfoPtrOutputWithContext(context.Context) CloningInfoPtrOutput
}

type cloningInfoPtrType CloningInfoArgs

func CloningInfoPtr(v *CloningInfoArgs) CloningInfoPtrInput {
	return (*cloningInfoPtrType)(v)
}

func (*cloningInfoPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfo)(nil)).Elem()
}

func (i *cloningInfoPtrType) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return i.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (i *cloningInfoPtrType) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoPtrOutput)
}

type CloningInfoOutput struct{ *pulumi.OutputState }

func (CloningInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfo)(nil)).Elem()
}

func (o CloningInfoOutput) ToCloningInfoOutput() CloningInfoOutput {
	return o
}

func (o CloningInfoOutput) ToCloningInfoOutputWithContext(ctx context.Context) CloningInfoOutput {
	return o
}

func (o CloningInfoOutput) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return o.ToCloningInfoPtrOutputWithContext(context.Background())
}

func (o CloningInfoOutput) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CloningInfo) *CloningInfo {
		return &v
	}).(CloningInfoPtrOutput)
}

func (o CloningInfoOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v CloningInfo) map[string]string { return v.AppSettingsOverrides }).(pulumi.StringMapOutput)
}

func (o CloningInfoOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.CloneCustomHostNames }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.CloneSourceControl }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.ConfigureLoadBalancing }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfo) *bool { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoOutput) SourceWebAppId() pulumi.StringOutput {
	return o.ApplyT(func(v CloningInfo) string { return v.SourceWebAppId }).(pulumi.StringOutput)
}

func (o CloningInfoOutput) SourceWebAppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.SourceWebAppLocation }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.TrafficManagerProfileId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfo) *string { return v.TrafficManagerProfileName }).(pulumi.StringPtrOutput)
}

type CloningInfoPtrOutput struct{ *pulumi.OutputState }

func (CloningInfoPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CloningInfo)(nil)).Elem()
}

func (o CloningInfoPtrOutput) ToCloningInfoPtrOutput() CloningInfoPtrOutput {
	return o
}

func (o CloningInfoPtrOutput) ToCloningInfoPtrOutputWithContext(ctx context.Context) CloningInfoPtrOutput {
	return o
}

func (o CloningInfoPtrOutput) Elem() CloningInfoOutput {
	return o.ApplyT(func(v *CloningInfo) CloningInfo {
		if v != nil {
			return *v
		}
		var ret CloningInfo
		return ret
	}).(CloningInfoOutput)
}

func (o CloningInfoPtrOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v *CloningInfo) map[string]string {
		if v == nil {
			return nil
		}
		return v.AppSettingsOverrides
	}).(pulumi.StringMapOutput)
}

func (o CloningInfoPtrOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.CloneCustomHostNames
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.CloneSourceControl
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.ConfigureLoadBalancing
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.CorrelationId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.HostingEnvironment
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *bool {
		if v == nil {
			return nil
		}
		return v.Overwrite
	}).(pulumi.BoolPtrOutput)
}

func (o CloningInfoPtrOutput) SourceWebAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return &v.SourceWebAppId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) SourceWebAppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.SourceWebAppLocation
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileId
	}).(pulumi.StringPtrOutput)
}

func (o CloningInfoPtrOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CloningInfo) *string {
		if v == nil {
			return nil
		}
		return v.TrafficManagerProfileName
	}).(pulumi.StringPtrOutput)
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


func (val *ConfigurationArgs) Defaults() *ConfigurationArgs {
	if val == nil {
		return nil
	}
	tmp := *val

	return &tmp
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

type ConnStringInfo struct {
	ConnectionString *string               `pulumi:"connectionString"`
	Name             *string               `pulumi:"name"`
	Type             *ConnectionStringType `pulumi:"type"`
}





type ConnStringInfoInput interface {
	pulumi.Input

	ToConnStringInfoOutput() ConnStringInfoOutput
	ToConnStringInfoOutputWithContext(context.Context) ConnStringInfoOutput
}

type ConnStringInfoArgs struct {
	ConnectionString pulumi.StringPtrInput        `pulumi:"connectionString"`
	Name             pulumi.StringPtrInput        `pulumi:"name"`
	Type             ConnectionStringTypePtrInput `pulumi:"type"`
}

func (ConnStringInfoArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfo)(nil)).Elem()
}

func (i ConnStringInfoArgs) ToConnStringInfoOutput() ConnStringInfoOutput {
	return i.ToConnStringInfoOutputWithContext(context.Background())
}

func (i ConnStringInfoArgs) ToConnStringInfoOutputWithContext(ctx context.Context) ConnStringInfoOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoOutput)
}





type ConnStringInfoArrayInput interface {
	pulumi.Input

	ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput
	ToConnStringInfoArrayOutputWithContext(context.Context) ConnStringInfoArrayOutput
}

type ConnStringInfoArray []ConnStringInfoInput

func (ConnStringInfoArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfo)(nil)).Elem()
}

func (i ConnStringInfoArray) ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput {
	return i.ToConnStringInfoArrayOutputWithContext(context.Background())
}

func (i ConnStringInfoArray) ToConnStringInfoArrayOutputWithContext(ctx context.Context) ConnStringInfoArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoArrayOutput)
}

type ConnStringInfoOutput struct{ *pulumi.OutputState }

func (ConnStringInfoOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfo)(nil)).Elem()
}

func (o ConnStringInfoOutput) ToConnStringInfoOutput() ConnStringInfoOutput {
	return o
}

func (o ConnStringInfoOutput) ToConnStringInfoOutputWithContext(ctx context.Context) ConnStringInfoOutput {
	return o
}

func (o ConnStringInfoOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoOutput) Type() ConnectionStringTypePtrOutput {
	return o.ApplyT(func(v ConnStringInfo) *ConnectionStringType { return v.Type }).(ConnectionStringTypePtrOutput)
}

type ConnStringInfoArrayOutput struct{ *pulumi.OutputState }

func (ConnStringInfoArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfo)(nil)).Elem()
}

func (o ConnStringInfoArrayOutput) ToConnStringInfoArrayOutput() ConnStringInfoArrayOutput {
	return o
}

func (o ConnStringInfoArrayOutput) ToConnStringInfoArrayOutputWithContext(ctx context.Context) ConnStringInfoArrayOutput {
	return o
}

func (o ConnStringInfoArrayOutput) Index(i pulumi.IntInput) ConnStringInfoOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnStringInfo {
		return vs[0].([]ConnStringInfo)[vs[1].(int)]
	}).(ConnStringInfoOutput)
}

type ConnStringInfoResponse struct {
	ConnectionString *string `pulumi:"connectionString"`
	Name             *string `pulumi:"name"`
	Type             *string `pulumi:"type"`
}

type ConnStringInfoResponseOutput struct{ *pulumi.OutputState }

func (ConnStringInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfoResponse)(nil)).Elem()
}

func (o ConnStringInfoResponseOutput) ToConnStringInfoResponseOutput() ConnStringInfoResponseOutput {
	return o
}

func (o ConnStringInfoResponseOutput) ToConnStringInfoResponseOutputWithContext(ctx context.Context) ConnStringInfoResponseOutput {
	return o
}

func (o ConnStringInfoResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnStringInfoResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnStringInfoResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

type ConnStringInfoResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnStringInfoResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfoResponse)(nil)).Elem()
}

func (o ConnStringInfoResponseArrayOutput) ToConnStringInfoResponseArrayOutput() ConnStringInfoResponseArrayOutput {
	return o
}

func (o ConnStringInfoResponseArrayOutput) ToConnStringInfoResponseArrayOutputWithContext(ctx context.Context) ConnStringInfoResponseArrayOutput {
	return o
}

func (o ConnStringInfoResponseArrayOutput) Index(i pulumi.IntInput) ConnStringInfoResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnStringInfoResponse {
		return vs[0].([]ConnStringInfoResponse)[vs[1].(int)]
	}).(ConnStringInfoResponseOutput)
}

type ConnStringValueTypePair struct {
	Type  ConnectionStringType `pulumi:"type"`
	Value string               `pulumi:"value"`
}





type ConnStringValueTypePairInput interface {
	pulumi.Input

	ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput
	ToConnStringValueTypePairOutputWithContext(context.Context) ConnStringValueTypePairOutput
}

type ConnStringValueTypePairArgs struct {
	Type  ConnectionStringTypeInput `pulumi:"type"`
	Value pulumi.StringInput        `pulumi:"value"`
}

func (ConnStringValueTypePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePair)(nil)).Elem()
}

func (i ConnStringValueTypePairArgs) ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput {
	return i.ToConnStringValueTypePairOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairArgs) ToConnStringValueTypePairOutputWithContext(ctx context.Context) ConnStringValueTypePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairOutput)
}





type ConnStringValueTypePairMapInput interface {
	pulumi.Input

	ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput
	ToConnStringValueTypePairMapOutputWithContext(context.Context) ConnStringValueTypePairMapOutput
}

type ConnStringValueTypePairMap map[string]ConnStringValueTypePairInput

func (ConnStringValueTypePairMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePair)(nil)).Elem()
}

func (i ConnStringValueTypePairMap) ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput {
	return i.ToConnStringValueTypePairMapOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairMap) ToConnStringValueTypePairMapOutputWithContext(ctx context.Context) ConnStringValueTypePairMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairMapOutput)
}

type ConnStringValueTypePairOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePair)(nil)).Elem()
}

func (o ConnStringValueTypePairOutput) ToConnStringValueTypePairOutput() ConnStringValueTypePairOutput {
	return o
}

func (o ConnStringValueTypePairOutput) ToConnStringValueTypePairOutputWithContext(ctx context.Context) ConnStringValueTypePairOutput {
	return o
}

func (o ConnStringValueTypePairOutput) Type() ConnectionStringTypeOutput {
	return o.ApplyT(func(v ConnStringValueTypePair) ConnectionStringType { return v.Type }).(ConnectionStringTypeOutput)
}

func (o ConnStringValueTypePairOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePair) string { return v.Value }).(pulumi.StringOutput)
}

type ConnStringValueTypePairMapOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePair)(nil)).Elem()
}

func (o ConnStringValueTypePairMapOutput) ToConnStringValueTypePairMapOutput() ConnStringValueTypePairMapOutput {
	return o
}

func (o ConnStringValueTypePairMapOutput) ToConnStringValueTypePairMapOutputWithContext(ctx context.Context) ConnStringValueTypePairMapOutput {
	return o
}

func (o ConnStringValueTypePairMapOutput) MapIndex(k pulumi.StringInput) ConnStringValueTypePairOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnStringValueTypePair {
		return vs[0].(map[string]ConnStringValueTypePair)[vs[1].(string)]
	}).(ConnStringValueTypePairOutput)
}

type ConnStringValueTypePairResponse struct {
	Type  string `pulumi:"type"`
	Value string `pulumi:"value"`
}

type ConnStringValueTypePairResponseOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePairResponse)(nil)).Elem()
}

func (o ConnStringValueTypePairResponseOutput) ToConnStringValueTypePairResponseOutput() ConnStringValueTypePairResponseOutput {
	return o
}

func (o ConnStringValueTypePairResponseOutput) ToConnStringValueTypePairResponseOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseOutput {
	return o
}

func (o ConnStringValueTypePairResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePairResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o ConnStringValueTypePairResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ConnStringValueTypePairResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ConnStringValueTypePairResponseMapOutput struct{ *pulumi.OutputState }

func (ConnStringValueTypePairResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePairResponse)(nil)).Elem()
}

func (o ConnStringValueTypePairResponseMapOutput) ToConnStringValueTypePairResponseMapOutput() ConnStringValueTypePairResponseMapOutput {
	return o
}

func (o ConnStringValueTypePairResponseMapOutput) ToConnStringValueTypePairResponseMapOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseMapOutput {
	return o
}

func (o ConnStringValueTypePairResponseMapOutput) MapIndex(k pulumi.StringInput) ConnStringValueTypePairResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) ConnStringValueTypePairResponse {
		return vs[0].(map[string]ConnStringValueTypePairResponse)[vs[1].(string)]
	}).(ConnStringValueTypePairResponseOutput)
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

type ContainerAppSecretResponse struct {
	Name  string `pulumi:"name"`
	Value string `pulumi:"value"`
}

type ContainerAppSecretResponseOutput struct{ *pulumi.OutputState }

func (ContainerAppSecretResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppSecretResponse)(nil)).Elem()
}

func (o ContainerAppSecretResponseOutput) ToContainerAppSecretResponseOutput() ContainerAppSecretResponseOutput {
	return o
}

func (o ContainerAppSecretResponseOutput) ToContainerAppSecretResponseOutputWithContext(ctx context.Context) ContainerAppSecretResponseOutput {
	return o
}

func (o ContainerAppSecretResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppSecretResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ContainerAppSecretResponseOutput) Value() pulumi.StringOutput {
	return o.ApplyT(func(v ContainerAppSecretResponse) string { return v.Value }).(pulumi.StringOutput)
}

type ContainerAppSecretResponseArrayOutput struct{ *pulumi.OutputState }

func (ContainerAppSecretResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ContainerAppSecretResponse)(nil)).Elem()
}

func (o ContainerAppSecretResponseArrayOutput) ToContainerAppSecretResponseArrayOutput() ContainerAppSecretResponseArrayOutput {
	return o
}

func (o ContainerAppSecretResponseArrayOutput) ToContainerAppSecretResponseArrayOutputWithContext(ctx context.Context) ContainerAppSecretResponseArrayOutput {
	return o
}

func (o ContainerAppSecretResponseArrayOutput) Index(i pulumi.IntInput) ContainerAppSecretResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ContainerAppSecretResponse {
		return vs[0].([]ContainerAppSecretResponse)[vs[1].(int)]
	}).(ContainerAppSecretResponseOutput)
}

type ContainerAppsConfiguration struct {
	AppSubnetResourceId          *string `pulumi:"appSubnetResourceId"`
	ControlPlaneSubnetResourceId *string `pulumi:"controlPlaneSubnetResourceId"`
	DaprAIInstrumentationKey     *string `pulumi:"daprAIInstrumentationKey"`
	DockerBridgeCidr             *string `pulumi:"dockerBridgeCidr"`
	PlatformReservedCidr         *string `pulumi:"platformReservedCidr"`
	PlatformReservedDnsIP        *string `pulumi:"platformReservedDnsIP"`
}





type ContainerAppsConfigurationInput interface {
	pulumi.Input

	ToContainerAppsConfigurationOutput() ContainerAppsConfigurationOutput
	ToContainerAppsConfigurationOutputWithContext(context.Context) ContainerAppsConfigurationOutput
}

type ContainerAppsConfigurationArgs struct {
	AppSubnetResourceId          pulumi.StringPtrInput `pulumi:"appSubnetResourceId"`
	ControlPlaneSubnetResourceId pulumi.StringPtrInput `pulumi:"controlPlaneSubnetResourceId"`
	DaprAIInstrumentationKey     pulumi.StringPtrInput `pulumi:"daprAIInstrumentationKey"`
	DockerBridgeCidr             pulumi.StringPtrInput `pulumi:"dockerBridgeCidr"`
	PlatformReservedCidr         pulumi.StringPtrInput `pulumi:"platformReservedCidr"`
	PlatformReservedDnsIP        pulumi.StringPtrInput `pulumi:"platformReservedDnsIP"`
}

func (ContainerAppsConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppsConfiguration)(nil)).Elem()
}

func (i ContainerAppsConfigurationArgs) ToContainerAppsConfigurationOutput() ContainerAppsConfigurationOutput {
	return i.ToContainerAppsConfigurationOutputWithContext(context.Background())
}

func (i ContainerAppsConfigurationArgs) ToContainerAppsConfigurationOutputWithContext(ctx context.Context) ContainerAppsConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsConfigurationOutput)
}

func (i ContainerAppsConfigurationArgs) ToContainerAppsConfigurationPtrOutput() ContainerAppsConfigurationPtrOutput {
	return i.ToContainerAppsConfigurationPtrOutputWithContext(context.Background())
}

func (i ContainerAppsConfigurationArgs) ToContainerAppsConfigurationPtrOutputWithContext(ctx context.Context) ContainerAppsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsConfigurationOutput).ToContainerAppsConfigurationPtrOutputWithContext(ctx)
}









type ContainerAppsConfigurationPtrInput interface {
	pulumi.Input

	ToContainerAppsConfigurationPtrOutput() ContainerAppsConfigurationPtrOutput
	ToContainerAppsConfigurationPtrOutputWithContext(context.Context) ContainerAppsConfigurationPtrOutput
}

type containerAppsConfigurationPtrType ContainerAppsConfigurationArgs

func ContainerAppsConfigurationPtr(v *ContainerAppsConfigurationArgs) ContainerAppsConfigurationPtrInput {
	return (*containerAppsConfigurationPtrType)(v)
}

func (*containerAppsConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsConfiguration)(nil)).Elem()
}

func (i *containerAppsConfigurationPtrType) ToContainerAppsConfigurationPtrOutput() ContainerAppsConfigurationPtrOutput {
	return i.ToContainerAppsConfigurationPtrOutputWithContext(context.Background())
}

func (i *containerAppsConfigurationPtrType) ToContainerAppsConfigurationPtrOutputWithContext(ctx context.Context) ContainerAppsConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ContainerAppsConfigurationPtrOutput)
}

type ContainerAppsConfigurationOutput struct{ *pulumi.OutputState }

func (ContainerAppsConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppsConfiguration)(nil)).Elem()
}

func (o ContainerAppsConfigurationOutput) ToContainerAppsConfigurationOutput() ContainerAppsConfigurationOutput {
	return o
}

func (o ContainerAppsConfigurationOutput) ToContainerAppsConfigurationOutputWithContext(ctx context.Context) ContainerAppsConfigurationOutput {
	return o
}

func (o ContainerAppsConfigurationOutput) ToContainerAppsConfigurationPtrOutput() ContainerAppsConfigurationPtrOutput {
	return o.ToContainerAppsConfigurationPtrOutputWithContext(context.Background())
}

func (o ContainerAppsConfigurationOutput) ToContainerAppsConfigurationPtrOutputWithContext(ctx context.Context) ContainerAppsConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ContainerAppsConfiguration) *ContainerAppsConfiguration {
		return &v
	}).(ContainerAppsConfigurationPtrOutput)
}

func (o ContainerAppsConfigurationOutput) AppSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfiguration) *string { return v.AppSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationOutput) ControlPlaneSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfiguration) *string { return v.ControlPlaneSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfiguration) *string { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfiguration) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfiguration) *string { return v.PlatformReservedCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfiguration) *string { return v.PlatformReservedDnsIP }).(pulumi.StringPtrOutput)
}

type ContainerAppsConfigurationPtrOutput struct{ *pulumi.OutputState }

func (ContainerAppsConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsConfiguration)(nil)).Elem()
}

func (o ContainerAppsConfigurationPtrOutput) ToContainerAppsConfigurationPtrOutput() ContainerAppsConfigurationPtrOutput {
	return o
}

func (o ContainerAppsConfigurationPtrOutput) ToContainerAppsConfigurationPtrOutputWithContext(ctx context.Context) ContainerAppsConfigurationPtrOutput {
	return o
}

func (o ContainerAppsConfigurationPtrOutput) Elem() ContainerAppsConfigurationOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) ContainerAppsConfiguration {
		if v != nil {
			return *v
		}
		var ret ContainerAppsConfiguration
		return ret
	}).(ContainerAppsConfigurationOutput)
}

func (o ContainerAppsConfigurationPtrOutput) AppSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.AppSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationPtrOutput) ControlPlaneSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ControlPlaneSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationPtrOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DaprAIInstrumentationKey
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationPtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationPtrOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationPtrOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedDnsIP
	}).(pulumi.StringPtrOutput)
}

type ContainerAppsConfigurationResponse struct {
	AppSubnetResourceId          *string `pulumi:"appSubnetResourceId"`
	ControlPlaneSubnetResourceId *string `pulumi:"controlPlaneSubnetResourceId"`
	DaprAIInstrumentationKey     *string `pulumi:"daprAIInstrumentationKey"`
	DockerBridgeCidr             *string `pulumi:"dockerBridgeCidr"`
	PlatformReservedCidr         *string `pulumi:"platformReservedCidr"`
	PlatformReservedDnsIP        *string `pulumi:"platformReservedDnsIP"`
}

type ContainerAppsConfigurationResponseOutput struct{ *pulumi.OutputState }

func (ContainerAppsConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ContainerAppsConfigurationResponse)(nil)).Elem()
}

func (o ContainerAppsConfigurationResponseOutput) ToContainerAppsConfigurationResponseOutput() ContainerAppsConfigurationResponseOutput {
	return o
}

func (o ContainerAppsConfigurationResponseOutput) ToContainerAppsConfigurationResponseOutputWithContext(ctx context.Context) ContainerAppsConfigurationResponseOutput {
	return o
}

func (o ContainerAppsConfigurationResponseOutput) AppSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfigurationResponse) *string { return v.AppSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponseOutput) ControlPlaneSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfigurationResponse) *string { return v.ControlPlaneSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponseOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfigurationResponse) *string { return v.DaprAIInstrumentationKey }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponseOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfigurationResponse) *string { return v.DockerBridgeCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponseOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfigurationResponse) *string { return v.PlatformReservedCidr }).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponseOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ContainerAppsConfigurationResponse) *string { return v.PlatformReservedDnsIP }).(pulumi.StringPtrOutput)
}

type ContainerAppsConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (ContainerAppsConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ContainerAppsConfigurationResponse)(nil)).Elem()
}

func (o ContainerAppsConfigurationResponsePtrOutput) ToContainerAppsConfigurationResponsePtrOutput() ContainerAppsConfigurationResponsePtrOutput {
	return o
}

func (o ContainerAppsConfigurationResponsePtrOutput) ToContainerAppsConfigurationResponsePtrOutputWithContext(ctx context.Context) ContainerAppsConfigurationResponsePtrOutput {
	return o
}

func (o ContainerAppsConfigurationResponsePtrOutput) Elem() ContainerAppsConfigurationResponseOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) ContainerAppsConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret ContainerAppsConfigurationResponse
		return ret
	}).(ContainerAppsConfigurationResponseOutput)
}

func (o ContainerAppsConfigurationResponsePtrOutput) AppSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponsePtrOutput) ControlPlaneSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ControlPlaneSubnetResourceId
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponsePtrOutput) DaprAIInstrumentationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DaprAIInstrumentationKey
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponsePtrOutput) DockerBridgeCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DockerBridgeCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponsePtrOutput) PlatformReservedCidr() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedCidr
	}).(pulumi.StringPtrOutput)
}

func (o ContainerAppsConfigurationResponsePtrOutput) PlatformReservedDnsIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ContainerAppsConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.PlatformReservedDnsIP
	}).(pulumi.StringPtrOutput)
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

type CorsSettings struct {
	AllowedOrigins     []string `pulumi:"allowedOrigins"`
	SupportCredentials *bool    `pulumi:"supportCredentials"`
}





type CorsSettingsInput interface {
	pulumi.Input

	ToCorsSettingsOutput() CorsSettingsOutput
	ToCorsSettingsOutputWithContext(context.Context) CorsSettingsOutput
}

type CorsSettingsArgs struct {
	AllowedOrigins     pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	SupportCredentials pulumi.BoolPtrInput     `pulumi:"supportCredentials"`
}

func (CorsSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettings)(nil)).Elem()
}

func (i CorsSettingsArgs) ToCorsSettingsOutput() CorsSettingsOutput {
	return i.ToCorsSettingsOutputWithContext(context.Background())
}

func (i CorsSettingsArgs) ToCorsSettingsOutputWithContext(ctx context.Context) CorsSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsOutput)
}

func (i CorsSettingsArgs) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return i.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (i CorsSettingsArgs) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsOutput).ToCorsSettingsPtrOutputWithContext(ctx)
}









type CorsSettingsPtrInput interface {
	pulumi.Input

	ToCorsSettingsPtrOutput() CorsSettingsPtrOutput
	ToCorsSettingsPtrOutputWithContext(context.Context) CorsSettingsPtrOutput
}

type corsSettingsPtrType CorsSettingsArgs

func CorsSettingsPtr(v *CorsSettingsArgs) CorsSettingsPtrInput {
	return (*corsSettingsPtrType)(v)
}

func (*corsSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettings)(nil)).Elem()
}

func (i *corsSettingsPtrType) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return i.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (i *corsSettingsPtrType) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsPtrOutput)
}

type CorsSettingsOutput struct{ *pulumi.OutputState }

func (CorsSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettings)(nil)).Elem()
}

func (o CorsSettingsOutput) ToCorsSettingsOutput() CorsSettingsOutput {
	return o
}

func (o CorsSettingsOutput) ToCorsSettingsOutputWithContext(ctx context.Context) CorsSettingsOutput {
	return o
}

func (o CorsSettingsOutput) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return o.ToCorsSettingsPtrOutputWithContext(context.Background())
}

func (o CorsSettingsOutput) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsSettings) *CorsSettings {
		return &v
	}).(CorsSettingsPtrOutput)
}

func (o CorsSettingsOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsSettings) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsSettingsOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorsSettings) *bool { return v.SupportCredentials }).(pulumi.BoolPtrOutput)
}

type CorsSettingsPtrOutput struct{ *pulumi.OutputState }

func (CorsSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettings)(nil)).Elem()
}

func (o CorsSettingsPtrOutput) ToCorsSettingsPtrOutput() CorsSettingsPtrOutput {
	return o
}

func (o CorsSettingsPtrOutput) ToCorsSettingsPtrOutputWithContext(ctx context.Context) CorsSettingsPtrOutput {
	return o
}

func (o CorsSettingsPtrOutput) Elem() CorsSettingsOutput {
	return o.ApplyT(func(v *CorsSettings) CorsSettings {
		if v != nil {
			return *v
		}
		var ret CorsSettings
		return ret
	}).(CorsSettingsOutput)
}

func (o CorsSettingsPtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CorsSettings) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

func (o CorsSettingsPtrOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorsSettings) *bool {
		if v == nil {
			return nil
		}
		return v.SupportCredentials
	}).(pulumi.BoolPtrOutput)
}

type CorsSettingsResponse struct {
	AllowedOrigins     []string `pulumi:"allowedOrigins"`
	SupportCredentials *bool    `pulumi:"supportCredentials"`
}

type CorsSettingsResponseOutput struct{ *pulumi.OutputState }

func (CorsSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettingsResponse)(nil)).Elem()
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponseOutput() CorsSettingsResponseOutput {
	return o
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponseOutputWithContext(ctx context.Context) CorsSettingsResponseOutput {
	return o
}

func (o CorsSettingsResponseOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v CorsSettingsResponse) []string { return v.AllowedOrigins }).(pulumi.StringArrayOutput)
}

func (o CorsSettingsResponseOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CorsSettingsResponse) *bool { return v.SupportCredentials }).(pulumi.BoolPtrOutput)
}

type CorsSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (CorsSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettingsResponse)(nil)).Elem()
}

func (o CorsSettingsResponsePtrOutput) ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput {
	return o
}

func (o CorsSettingsResponsePtrOutput) ToCorsSettingsResponsePtrOutputWithContext(ctx context.Context) CorsSettingsResponsePtrOutput {
	return o
}

func (o CorsSettingsResponsePtrOutput) Elem() CorsSettingsResponseOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) CorsSettingsResponse {
		if v != nil {
			return *v
		}
		var ret CorsSettingsResponse
		return ret
	}).(CorsSettingsResponseOutput)
}

func (o CorsSettingsResponsePtrOutput) AllowedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) []string {
		if v == nil {
			return nil
		}
		return v.AllowedOrigins
	}).(pulumi.StringArrayOutput)
}

func (o CorsSettingsResponsePtrOutput) SupportCredentials() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *CorsSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SupportCredentials
	}).(pulumi.BoolPtrOutput)
}

type CustomDnsSuffixConfiguration struct {
	CertificateUrl            *string `pulumi:"certificateUrl"`
	DnsSuffix                 *string `pulumi:"dnsSuffix"`
	KeyVaultReferenceIdentity *string `pulumi:"keyVaultReferenceIdentity"`
	Kind                      *string `pulumi:"kind"`
}





type CustomDnsSuffixConfigurationInput interface {
	pulumi.Input

	ToCustomDnsSuffixConfigurationOutput() CustomDnsSuffixConfigurationOutput
	ToCustomDnsSuffixConfigurationOutputWithContext(context.Context) CustomDnsSuffixConfigurationOutput
}

type CustomDnsSuffixConfigurationArgs struct {
	CertificateUrl            pulumi.StringPtrInput `pulumi:"certificateUrl"`
	DnsSuffix                 pulumi.StringPtrInput `pulumi:"dnsSuffix"`
	KeyVaultReferenceIdentity pulumi.StringPtrInput `pulumi:"keyVaultReferenceIdentity"`
	Kind                      pulumi.StringPtrInput `pulumi:"kind"`
}

func (CustomDnsSuffixConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDnsSuffixConfiguration)(nil)).Elem()
}

func (i CustomDnsSuffixConfigurationArgs) ToCustomDnsSuffixConfigurationOutput() CustomDnsSuffixConfigurationOutput {
	return i.ToCustomDnsSuffixConfigurationOutputWithContext(context.Background())
}

func (i CustomDnsSuffixConfigurationArgs) ToCustomDnsSuffixConfigurationOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDnsSuffixConfigurationOutput)
}

func (i CustomDnsSuffixConfigurationArgs) ToCustomDnsSuffixConfigurationPtrOutput() CustomDnsSuffixConfigurationPtrOutput {
	return i.ToCustomDnsSuffixConfigurationPtrOutputWithContext(context.Background())
}

func (i CustomDnsSuffixConfigurationArgs) ToCustomDnsSuffixConfigurationPtrOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDnsSuffixConfigurationOutput).ToCustomDnsSuffixConfigurationPtrOutputWithContext(ctx)
}









type CustomDnsSuffixConfigurationPtrInput interface {
	pulumi.Input

	ToCustomDnsSuffixConfigurationPtrOutput() CustomDnsSuffixConfigurationPtrOutput
	ToCustomDnsSuffixConfigurationPtrOutputWithContext(context.Context) CustomDnsSuffixConfigurationPtrOutput
}

type customDnsSuffixConfigurationPtrType CustomDnsSuffixConfigurationArgs

func CustomDnsSuffixConfigurationPtr(v *CustomDnsSuffixConfigurationArgs) CustomDnsSuffixConfigurationPtrInput {
	return (*customDnsSuffixConfigurationPtrType)(v)
}

func (*customDnsSuffixConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDnsSuffixConfiguration)(nil)).Elem()
}

func (i *customDnsSuffixConfigurationPtrType) ToCustomDnsSuffixConfigurationPtrOutput() CustomDnsSuffixConfigurationPtrOutput {
	return i.ToCustomDnsSuffixConfigurationPtrOutputWithContext(context.Background())
}

func (i *customDnsSuffixConfigurationPtrType) ToCustomDnsSuffixConfigurationPtrOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CustomDnsSuffixConfigurationPtrOutput)
}

type CustomDnsSuffixConfigurationOutput struct{ *pulumi.OutputState }

func (CustomDnsSuffixConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDnsSuffixConfiguration)(nil)).Elem()
}

func (o CustomDnsSuffixConfigurationOutput) ToCustomDnsSuffixConfigurationOutput() CustomDnsSuffixConfigurationOutput {
	return o
}

func (o CustomDnsSuffixConfigurationOutput) ToCustomDnsSuffixConfigurationOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationOutput {
	return o
}

func (o CustomDnsSuffixConfigurationOutput) ToCustomDnsSuffixConfigurationPtrOutput() CustomDnsSuffixConfigurationPtrOutput {
	return o.ToCustomDnsSuffixConfigurationPtrOutputWithContext(context.Background())
}

func (o CustomDnsSuffixConfigurationOutput) ToCustomDnsSuffixConfigurationPtrOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CustomDnsSuffixConfiguration) *CustomDnsSuffixConfiguration {
		return &v
	}).(CustomDnsSuffixConfigurationPtrOutput)
}

func (o CustomDnsSuffixConfigurationOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfiguration) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfiguration) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfiguration) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfiguration) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

type CustomDnsSuffixConfigurationPtrOutput struct{ *pulumi.OutputState }

func (CustomDnsSuffixConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDnsSuffixConfiguration)(nil)).Elem()
}

func (o CustomDnsSuffixConfigurationPtrOutput) ToCustomDnsSuffixConfigurationPtrOutput() CustomDnsSuffixConfigurationPtrOutput {
	return o
}

func (o CustomDnsSuffixConfigurationPtrOutput) ToCustomDnsSuffixConfigurationPtrOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationPtrOutput {
	return o
}

func (o CustomDnsSuffixConfigurationPtrOutput) Elem() CustomDnsSuffixConfigurationOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfiguration) CustomDnsSuffixConfiguration {
		if v != nil {
			return *v
		}
		var ret CustomDnsSuffixConfiguration
		return ret
	}).(CustomDnsSuffixConfigurationOutput)
}

func (o CustomDnsSuffixConfigurationPtrOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.CertificateUrl
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationPtrOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.DnsSuffix
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationPtrOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

type CustomDnsSuffixConfigurationResponse struct {
	CertificateUrl            *string `pulumi:"certificateUrl"`
	DnsSuffix                 *string `pulumi:"dnsSuffix"`
	Id                        string  `pulumi:"id"`
	KeyVaultReferenceIdentity *string `pulumi:"keyVaultReferenceIdentity"`
	Kind                      *string `pulumi:"kind"`
	Name                      string  `pulumi:"name"`
	ProvisioningDetails       string  `pulumi:"provisioningDetails"`
	ProvisioningState         string  `pulumi:"provisioningState"`
	Type                      string  `pulumi:"type"`
}

type CustomDnsSuffixConfigurationResponseOutput struct{ *pulumi.OutputState }

func (CustomDnsSuffixConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CustomDnsSuffixConfigurationResponse)(nil)).Elem()
}

func (o CustomDnsSuffixConfigurationResponseOutput) ToCustomDnsSuffixConfigurationResponseOutput() CustomDnsSuffixConfigurationResponseOutput {
	return o
}

func (o CustomDnsSuffixConfigurationResponseOutput) ToCustomDnsSuffixConfigurationResponseOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationResponseOutput {
	return o
}

func (o CustomDnsSuffixConfigurationResponseOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) *string { return v.CertificateUrl }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) *string { return v.DnsSuffix }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) ProvisioningDetails() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) string { return v.ProvisioningDetails }).(pulumi.StringOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o CustomDnsSuffixConfigurationResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v CustomDnsSuffixConfigurationResponse) string { return v.Type }).(pulumi.StringOutput)
}

type CustomDnsSuffixConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (CustomDnsSuffixConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**CustomDnsSuffixConfigurationResponse)(nil)).Elem()
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) ToCustomDnsSuffixConfigurationResponsePtrOutput() CustomDnsSuffixConfigurationResponsePtrOutput {
	return o
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) ToCustomDnsSuffixConfigurationResponsePtrOutputWithContext(ctx context.Context) CustomDnsSuffixConfigurationResponsePtrOutput {
	return o
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) Elem() CustomDnsSuffixConfigurationResponseOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) CustomDnsSuffixConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret CustomDnsSuffixConfigurationResponse
		return ret
	}).(CustomDnsSuffixConfigurationResponseOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) CertificateUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.CertificateUrl
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) DnsSuffix() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.DnsSuffix
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) ProvisioningDetails() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningDetails
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o CustomDnsSuffixConfigurationResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *CustomDnsSuffixConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
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

type DatabaseBackupSetting struct {
	ConnectionString     *string `pulumi:"connectionString"`
	ConnectionStringName *string `pulumi:"connectionStringName"`
	DatabaseType         string  `pulumi:"databaseType"`
	Name                 *string `pulumi:"name"`
}





type DatabaseBackupSettingInput interface {
	pulumi.Input

	ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput
	ToDatabaseBackupSettingOutputWithContext(context.Context) DatabaseBackupSettingOutput
}

type DatabaseBackupSettingArgs struct {
	ConnectionString     pulumi.StringPtrInput `pulumi:"connectionString"`
	ConnectionStringName pulumi.StringPtrInput `pulumi:"connectionStringName"`
	DatabaseType         pulumi.StringInput    `pulumi:"databaseType"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
}

func (DatabaseBackupSettingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSetting)(nil)).Elem()
}

func (i DatabaseBackupSettingArgs) ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput {
	return i.ToDatabaseBackupSettingOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingArgs) ToDatabaseBackupSettingOutputWithContext(ctx context.Context) DatabaseBackupSettingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingOutput)
}





type DatabaseBackupSettingArrayInput interface {
	pulumi.Input

	ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput
	ToDatabaseBackupSettingArrayOutputWithContext(context.Context) DatabaseBackupSettingArrayOutput
}

type DatabaseBackupSettingArray []DatabaseBackupSettingInput

func (DatabaseBackupSettingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSetting)(nil)).Elem()
}

func (i DatabaseBackupSettingArray) ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput {
	return i.ToDatabaseBackupSettingArrayOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingArray) ToDatabaseBackupSettingArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingArrayOutput)
}

type DatabaseBackupSettingOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSetting)(nil)).Elem()
}

func (o DatabaseBackupSettingOutput) ToDatabaseBackupSettingOutput() DatabaseBackupSettingOutput {
	return o
}

func (o DatabaseBackupSettingOutput) ToDatabaseBackupSettingOutputWithContext(ctx context.Context) DatabaseBackupSettingOutput {
	return o
}

func (o DatabaseBackupSettingOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) ConnectionStringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.ConnectionStringName }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) string { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o DatabaseBackupSettingOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSetting) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatabaseBackupSettingArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSetting)(nil)).Elem()
}

func (o DatabaseBackupSettingArrayOutput) ToDatabaseBackupSettingArrayOutput() DatabaseBackupSettingArrayOutput {
	return o
}

func (o DatabaseBackupSettingArrayOutput) ToDatabaseBackupSettingArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingArrayOutput {
	return o
}

func (o DatabaseBackupSettingArrayOutput) Index(i pulumi.IntInput) DatabaseBackupSettingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseBackupSetting {
		return vs[0].([]DatabaseBackupSetting)[vs[1].(int)]
	}).(DatabaseBackupSettingOutput)
}

type DatabaseBackupSettingResponse struct {
	ConnectionString     *string `pulumi:"connectionString"`
	ConnectionStringName *string `pulumi:"connectionStringName"`
	DatabaseType         string  `pulumi:"databaseType"`
	Name                 *string `pulumi:"name"`
}

type DatabaseBackupSettingResponseOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSettingResponse)(nil)).Elem()
}

func (o DatabaseBackupSettingResponseOutput) ToDatabaseBackupSettingResponseOutput() DatabaseBackupSettingResponseOutput {
	return o
}

func (o DatabaseBackupSettingResponseOutput) ToDatabaseBackupSettingResponseOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseOutput {
	return o
}

func (o DatabaseBackupSettingResponseOutput) ConnectionString() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.ConnectionString }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) ConnectionStringName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.ConnectionStringName }).(pulumi.StringPtrOutput)
}

func (o DatabaseBackupSettingResponseOutput) DatabaseType() pulumi.StringOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) string { return v.DatabaseType }).(pulumi.StringOutput)
}

func (o DatabaseBackupSettingResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DatabaseBackupSettingResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

type DatabaseBackupSettingResponseArrayOutput struct{ *pulumi.OutputState }

func (DatabaseBackupSettingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSettingResponse)(nil)).Elem()
}

func (o DatabaseBackupSettingResponseArrayOutput) ToDatabaseBackupSettingResponseArrayOutput() DatabaseBackupSettingResponseArrayOutput {
	return o
}

func (o DatabaseBackupSettingResponseArrayOutput) ToDatabaseBackupSettingResponseArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseArrayOutput {
	return o
}

func (o DatabaseBackupSettingResponseArrayOutput) Index(i pulumi.IntInput) DatabaseBackupSettingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DatabaseBackupSettingResponse {
		return vs[0].([]DatabaseBackupSettingResponse)[vs[1].(int)]
	}).(DatabaseBackupSettingResponseOutput)
}

type EnabledConfig struct {
	Enabled *bool `pulumi:"enabled"`
}





type EnabledConfigInput interface {
	pulumi.Input

	ToEnabledConfigOutput() EnabledConfigOutput
	ToEnabledConfigOutputWithContext(context.Context) EnabledConfigOutput
}

type EnabledConfigArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EnabledConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfig)(nil)).Elem()
}

func (i EnabledConfigArgs) ToEnabledConfigOutput() EnabledConfigOutput {
	return i.ToEnabledConfigOutputWithContext(context.Background())
}

func (i EnabledConfigArgs) ToEnabledConfigOutputWithContext(ctx context.Context) EnabledConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigOutput)
}

func (i EnabledConfigArgs) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return i.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (i EnabledConfigArgs) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigOutput).ToEnabledConfigPtrOutputWithContext(ctx)
}









type EnabledConfigPtrInput interface {
	pulumi.Input

	ToEnabledConfigPtrOutput() EnabledConfigPtrOutput
	ToEnabledConfigPtrOutputWithContext(context.Context) EnabledConfigPtrOutput
}

type enabledConfigPtrType EnabledConfigArgs

func EnabledConfigPtr(v *EnabledConfigArgs) EnabledConfigPtrInput {
	return (*enabledConfigPtrType)(v)
}

func (*enabledConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfig)(nil)).Elem()
}

func (i *enabledConfigPtrType) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return i.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (i *enabledConfigPtrType) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigPtrOutput)
}

type EnabledConfigOutput struct{ *pulumi.OutputState }

func (EnabledConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfig)(nil)).Elem()
}

func (o EnabledConfigOutput) ToEnabledConfigOutput() EnabledConfigOutput {
	return o
}

func (o EnabledConfigOutput) ToEnabledConfigOutputWithContext(ctx context.Context) EnabledConfigOutput {
	return o
}

func (o EnabledConfigOutput) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return o.ToEnabledConfigPtrOutputWithContext(context.Background())
}

func (o EnabledConfigOutput) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledConfig) *EnabledConfig {
		return &v
	}).(EnabledConfigPtrOutput)
}

func (o EnabledConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnabledConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EnabledConfigPtrOutput struct{ *pulumi.OutputState }

func (EnabledConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfig)(nil)).Elem()
}

func (o EnabledConfigPtrOutput) ToEnabledConfigPtrOutput() EnabledConfigPtrOutput {
	return o
}

func (o EnabledConfigPtrOutput) ToEnabledConfigPtrOutputWithContext(ctx context.Context) EnabledConfigPtrOutput {
	return o
}

func (o EnabledConfigPtrOutput) Elem() EnabledConfigOutput {
	return o.ApplyT(func(v *EnabledConfig) EnabledConfig {
		if v != nil {
			return *v
		}
		var ret EnabledConfig
		return ret
	}).(EnabledConfigOutput)
}

func (o EnabledConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

type EnabledConfigResponse struct {
	Enabled *bool `pulumi:"enabled"`
}

type EnabledConfigResponseOutput struct{ *pulumi.OutputState }

func (EnabledConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfigResponse)(nil)).Elem()
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponseOutput() EnabledConfigResponseOutput {
	return o
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponseOutputWithContext(ctx context.Context) EnabledConfigResponseOutput {
	return o
}

func (o EnabledConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v EnabledConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

type EnabledConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (EnabledConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfigResponse)(nil)).Elem()
}

func (o EnabledConfigResponsePtrOutput) ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput {
	return o
}

func (o EnabledConfigResponsePtrOutput) ToEnabledConfigResponsePtrOutputWithContext(ctx context.Context) EnabledConfigResponsePtrOutput {
	return o
}

func (o EnabledConfigResponsePtrOutput) Elem() EnabledConfigResponseOutput {
	return o.ApplyT(func(v *EnabledConfigResponse) EnabledConfigResponse {
		if v != nil {
			return *v
		}
		var ret EnabledConfigResponse
		return ret
	}).(EnabledConfigResponseOutput)
}

func (o EnabledConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EnabledConfigResponse) *bool {
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

type ErrorEntityResponse struct {
	Code            *string               `pulumi:"code"`
	Details         []ErrorEntityResponse `pulumi:"details"`
	ExtendedCode    *string               `pulumi:"extendedCode"`
	InnerErrors     []ErrorEntityResponse `pulumi:"innerErrors"`
	Message         *string               `pulumi:"message"`
	MessageTemplate *string               `pulumi:"messageTemplate"`
	Parameters      []string              `pulumi:"parameters"`
	Target          *string               `pulumi:"target"`
}

type ErrorEntityResponseOutput struct{ *pulumi.OutputState }

func (ErrorEntityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ErrorEntityResponse)(nil)).Elem()
}

func (o ErrorEntityResponseOutput) ToErrorEntityResponseOutput() ErrorEntityResponseOutput {
	return o
}

func (o ErrorEntityResponseOutput) ToErrorEntityResponseOutputWithContext(ctx context.Context) ErrorEntityResponseOutput {
	return o
}

func (o ErrorEntityResponseOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.Code }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) Details() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v ErrorEntityResponse) []ErrorEntityResponse { return v.Details }).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponseOutput) ExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.ExtendedCode }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) InnerErrors() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v ErrorEntityResponse) []ErrorEntityResponse { return v.InnerErrors }).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponseOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.Message }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) MessageTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.MessageTemplate }).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponseOutput) Parameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ErrorEntityResponse) []string { return v.Parameters }).(pulumi.StringArrayOutput)
}

func (o ErrorEntityResponseOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ErrorEntityResponse) *string { return v.Target }).(pulumi.StringPtrOutput)
}

type ErrorEntityResponsePtrOutput struct{ *pulumi.OutputState }

func (ErrorEntityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ErrorEntityResponse)(nil)).Elem()
}

func (o ErrorEntityResponsePtrOutput) ToErrorEntityResponsePtrOutput() ErrorEntityResponsePtrOutput {
	return o
}

func (o ErrorEntityResponsePtrOutput) ToErrorEntityResponsePtrOutputWithContext(ctx context.Context) ErrorEntityResponsePtrOutput {
	return o
}

func (o ErrorEntityResponsePtrOutput) Elem() ErrorEntityResponseOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) ErrorEntityResponse {
		if v != nil {
			return *v
		}
		var ret ErrorEntityResponse
		return ret
	}).(ErrorEntityResponseOutput)
}

func (o ErrorEntityResponsePtrOutput) Code() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Code
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) Details() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) []ErrorEntityResponse {
		if v == nil {
			return nil
		}
		return v.Details
	}).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponsePtrOutput) ExtendedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtendedCode
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) InnerErrors() ErrorEntityResponseArrayOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) []ErrorEntityResponse {
		if v == nil {
			return nil
		}
		return v.InnerErrors
	}).(ErrorEntityResponseArrayOutput)
}

func (o ErrorEntityResponsePtrOutput) Message() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Message
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) MessageTemplate() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.MessageTemplate
	}).(pulumi.StringPtrOutput)
}

func (o ErrorEntityResponsePtrOutput) Parameters() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) []string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringArrayOutput)
}

func (o ErrorEntityResponsePtrOutput) Target() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ErrorEntityResponse) *string {
		if v == nil {
			return nil
		}
		return v.Target
	}).(pulumi.StringPtrOutput)
}

type ErrorEntityResponseArrayOutput struct{ *pulumi.OutputState }

func (ErrorEntityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ErrorEntityResponse)(nil)).Elem()
}

func (o ErrorEntityResponseArrayOutput) ToErrorEntityResponseArrayOutput() ErrorEntityResponseArrayOutput {
	return o
}

func (o ErrorEntityResponseArrayOutput) ToErrorEntityResponseArrayOutputWithContext(ctx context.Context) ErrorEntityResponseArrayOutput {
	return o
}

func (o ErrorEntityResponseArrayOutput) Index(i pulumi.IntInput) ErrorEntityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ErrorEntityResponse {
		return vs[0].([]ErrorEntityResponse)[vs[1].(int)]
	}).(ErrorEntityResponseOutput)
}

type Experiments struct {
	RampUpRules []RampUpRule `pulumi:"rampUpRules"`
}





type ExperimentsInput interface {
	pulumi.Input

	ToExperimentsOutput() ExperimentsOutput
	ToExperimentsOutputWithContext(context.Context) ExperimentsOutput
}

type ExperimentsArgs struct {
	RampUpRules RampUpRuleArrayInput `pulumi:"rampUpRules"`
}

func (ExperimentsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Experiments)(nil)).Elem()
}

func (i ExperimentsArgs) ToExperimentsOutput() ExperimentsOutput {
	return i.ToExperimentsOutputWithContext(context.Background())
}

func (i ExperimentsArgs) ToExperimentsOutputWithContext(ctx context.Context) ExperimentsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsOutput)
}

func (i ExperimentsArgs) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return i.ToExperimentsPtrOutputWithContext(context.Background())
}

func (i ExperimentsArgs) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsOutput).ToExperimentsPtrOutputWithContext(ctx)
}









type ExperimentsPtrInput interface {
	pulumi.Input

	ToExperimentsPtrOutput() ExperimentsPtrOutput
	ToExperimentsPtrOutputWithContext(context.Context) ExperimentsPtrOutput
}

type experimentsPtrType ExperimentsArgs

func ExperimentsPtr(v *ExperimentsArgs) ExperimentsPtrInput {
	return (*experimentsPtrType)(v)
}

func (*experimentsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiments)(nil)).Elem()
}

func (i *experimentsPtrType) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return i.ToExperimentsPtrOutputWithContext(context.Background())
}

func (i *experimentsPtrType) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsPtrOutput)
}

type ExperimentsOutput struct{ *pulumi.OutputState }

func (ExperimentsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Experiments)(nil)).Elem()
}

func (o ExperimentsOutput) ToExperimentsOutput() ExperimentsOutput {
	return o
}

func (o ExperimentsOutput) ToExperimentsOutputWithContext(ctx context.Context) ExperimentsOutput {
	return o
}

func (o ExperimentsOutput) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return o.ToExperimentsPtrOutputWithContext(context.Background())
}

func (o ExperimentsOutput) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Experiments) *Experiments {
		return &v
	}).(ExperimentsPtrOutput)
}

func (o ExperimentsOutput) RampUpRules() RampUpRuleArrayOutput {
	return o.ApplyT(func(v Experiments) []RampUpRule { return v.RampUpRules }).(RampUpRuleArrayOutput)
}

type ExperimentsPtrOutput struct{ *pulumi.OutputState }

func (ExperimentsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Experiments)(nil)).Elem()
}

func (o ExperimentsPtrOutput) ToExperimentsPtrOutput() ExperimentsPtrOutput {
	return o
}

func (o ExperimentsPtrOutput) ToExperimentsPtrOutputWithContext(ctx context.Context) ExperimentsPtrOutput {
	return o
}

func (o ExperimentsPtrOutput) Elem() ExperimentsOutput {
	return o.ApplyT(func(v *Experiments) Experiments {
		if v != nil {
			return *v
		}
		var ret Experiments
		return ret
	}).(ExperimentsOutput)
}

func (o ExperimentsPtrOutput) RampUpRules() RampUpRuleArrayOutput {
	return o.ApplyT(func(v *Experiments) []RampUpRule {
		if v == nil {
			return nil
		}
		return v.RampUpRules
	}).(RampUpRuleArrayOutput)
}

type ExperimentsResponse struct {
	RampUpRules []RampUpRuleResponse `pulumi:"rampUpRules"`
}

type ExperimentsResponseOutput struct{ *pulumi.OutputState }

func (ExperimentsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentsResponse)(nil)).Elem()
}

func (o ExperimentsResponseOutput) ToExperimentsResponseOutput() ExperimentsResponseOutput {
	return o
}

func (o ExperimentsResponseOutput) ToExperimentsResponseOutputWithContext(ctx context.Context) ExperimentsResponseOutput {
	return o
}

func (o ExperimentsResponseOutput) RampUpRules() RampUpRuleResponseArrayOutput {
	return o.ApplyT(func(v ExperimentsResponse) []RampUpRuleResponse { return v.RampUpRules }).(RampUpRuleResponseArrayOutput)
}

type ExperimentsResponsePtrOutput struct{ *pulumi.OutputState }

func (ExperimentsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentsResponse)(nil)).Elem()
}

func (o ExperimentsResponsePtrOutput) ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput {
	return o
}

func (o ExperimentsResponsePtrOutput) ToExperimentsResponsePtrOutputWithContext(ctx context.Context) ExperimentsResponsePtrOutput {
	return o
}

func (o ExperimentsResponsePtrOutput) Elem() ExperimentsResponseOutput {
	return o.ApplyT(func(v *ExperimentsResponse) ExperimentsResponse {
		if v != nil {
			return *v
		}
		var ret ExperimentsResponse
		return ret
	}).(ExperimentsResponseOutput)
}

func (o ExperimentsResponsePtrOutput) RampUpRules() RampUpRuleResponseArrayOutput {
	return o.ApplyT(func(v *ExperimentsResponse) []RampUpRuleResponse {
		if v == nil {
			return nil
		}
		return v.RampUpRules
	}).(RampUpRuleResponseArrayOutput)
}

type ExpressionResponse struct {
	Error          *AzureResourceErrorInfoResponse `pulumi:"error"`
	Subexpressions []ExpressionResponse            `pulumi:"subexpressions"`
	Text           *string                         `pulumi:"text"`
	Value          interface{}                     `pulumi:"value"`
}

type ExpressionResponseOutput struct{ *pulumi.OutputState }

func (ExpressionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionResponse)(nil)).Elem()
}

func (o ExpressionResponseOutput) ToExpressionResponseOutput() ExpressionResponseOutput {
	return o
}

func (o ExpressionResponseOutput) ToExpressionResponseOutputWithContext(ctx context.Context) ExpressionResponseOutput {
	return o
}

func (o ExpressionResponseOutput) Error() AzureResourceErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ExpressionResponse) *AzureResourceErrorInfoResponse { return v.Error }).(AzureResourceErrorInfoResponsePtrOutput)
}

func (o ExpressionResponseOutput) Subexpressions() ExpressionResponseArrayOutput {
	return o.ApplyT(func(v ExpressionResponse) []ExpressionResponse { return v.Subexpressions }).(ExpressionResponseArrayOutput)
}

func (o ExpressionResponseOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionResponse) *string { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ExpressionResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ExpressionResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ExpressionResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressionResponse)(nil)).Elem()
}

func (o ExpressionResponseArrayOutput) ToExpressionResponseArrayOutput() ExpressionResponseArrayOutput {
	return o
}

func (o ExpressionResponseArrayOutput) ToExpressionResponseArrayOutputWithContext(ctx context.Context) ExpressionResponseArrayOutput {
	return o
}

func (o ExpressionResponseArrayOutput) Index(i pulumi.IntInput) ExpressionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressionResponse {
		return vs[0].([]ExpressionResponse)[vs[1].(int)]
	}).(ExpressionResponseOutput)
}

type ExpressionRootResponse struct {
	Error          *AzureResourceErrorInfoResponse `pulumi:"error"`
	Path           *string                         `pulumi:"path"`
	Subexpressions []ExpressionResponse            `pulumi:"subexpressions"`
	Text           *string                         `pulumi:"text"`
	Value          interface{}                     `pulumi:"value"`
}

type ExpressionRootResponseOutput struct{ *pulumi.OutputState }

func (ExpressionRootResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ExpressionRootResponse)(nil)).Elem()
}

func (o ExpressionRootResponseOutput) ToExpressionRootResponseOutput() ExpressionRootResponseOutput {
	return o
}

func (o ExpressionRootResponseOutput) ToExpressionRootResponseOutputWithContext(ctx context.Context) ExpressionRootResponseOutput {
	return o
}

func (o ExpressionRootResponseOutput) Error() AzureResourceErrorInfoResponsePtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *AzureResourceErrorInfoResponse { return v.Error }).(AzureResourceErrorInfoResponsePtrOutput)
}

func (o ExpressionRootResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o ExpressionRootResponseOutput) Subexpressions() ExpressionResponseArrayOutput {
	return o.ApplyT(func(v ExpressionRootResponse) []ExpressionResponse { return v.Subexpressions }).(ExpressionResponseArrayOutput)
}

func (o ExpressionRootResponseOutput) Text() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ExpressionRootResponse) *string { return v.Text }).(pulumi.StringPtrOutput)
}

func (o ExpressionRootResponseOutput) Value() pulumi.AnyOutput {
	return o.ApplyT(func(v ExpressionRootResponse) interface{} { return v.Value }).(pulumi.AnyOutput)
}

type ExpressionRootResponseArrayOutput struct{ *pulumi.OutputState }

func (ExpressionRootResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ExpressionRootResponse)(nil)).Elem()
}

func (o ExpressionRootResponseArrayOutput) ToExpressionRootResponseArrayOutput() ExpressionRootResponseArrayOutput {
	return o
}

func (o ExpressionRootResponseArrayOutput) ToExpressionRootResponseArrayOutputWithContext(ctx context.Context) ExpressionRootResponseArrayOutput {
	return o
}

func (o ExpressionRootResponseArrayOutput) Index(i pulumi.IntInput) ExpressionRootResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ExpressionRootResponse {
		return vs[0].([]ExpressionRootResponse)[vs[1].(int)]
	}).(ExpressionRootResponseOutput)
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

type FileSystemApplicationLogsConfig struct {
	Level *LogLevel `pulumi:"level"`
}


func (val *FileSystemApplicationLogsConfig) Defaults() *FileSystemApplicationLogsConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := LogLevel("Off")
		tmp.Level = &level_
	}
	return &tmp
}





type FileSystemApplicationLogsConfigInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput
	ToFileSystemApplicationLogsConfigOutputWithContext(context.Context) FileSystemApplicationLogsConfigOutput
}

type FileSystemApplicationLogsConfigArgs struct {
	Level LogLevelPtrInput `pulumi:"level"`
}


func (val *FileSystemApplicationLogsConfigArgs) Defaults() *FileSystemApplicationLogsConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		tmp.Level = LogLevel("Off")
	}
	return &tmp
}
func (FileSystemApplicationLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput {
	return i.ToFileSystemApplicationLogsConfigOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigOutput)
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return i.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigArgs) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigOutput).ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx)
}









type FileSystemApplicationLogsConfigPtrInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput
	ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Context) FileSystemApplicationLogsConfigPtrOutput
}

type fileSystemApplicationLogsConfigPtrType FileSystemApplicationLogsConfigArgs

func FileSystemApplicationLogsConfigPtr(v *FileSystemApplicationLogsConfigArgs) FileSystemApplicationLogsConfigPtrInput {
	return (*fileSystemApplicationLogsConfigPtrType)(v)
}

func (*fileSystemApplicationLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (i *fileSystemApplicationLogsConfigPtrType) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return i.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (i *fileSystemApplicationLogsConfigPtrType) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigPtrOutput)
}

type FileSystemApplicationLogsConfigOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput {
	return o
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigOutput {
	return o
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return o.ToFileSystemApplicationLogsConfigPtrOutputWithContext(context.Background())
}

func (o FileSystemApplicationLogsConfigOutput) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemApplicationLogsConfig) *FileSystemApplicationLogsConfig {
		return &v
	}).(FileSystemApplicationLogsConfigPtrOutput)
}

func (o FileSystemApplicationLogsConfigOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v FileSystemApplicationLogsConfig) *LogLevel { return v.Level }).(LogLevelPtrOutput)
}

type FileSystemApplicationLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfig)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigPtrOutput) ToFileSystemApplicationLogsConfigPtrOutput() FileSystemApplicationLogsConfigPtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigPtrOutput) ToFileSystemApplicationLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigPtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigPtrOutput) Elem() FileSystemApplicationLogsConfigOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfig) FileSystemApplicationLogsConfig {
		if v != nil {
			return *v
		}
		var ret FileSystemApplicationLogsConfig
		return ret
	}).(FileSystemApplicationLogsConfigOutput)
}

func (o FileSystemApplicationLogsConfigPtrOutput) Level() LogLevelPtrOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfig) *LogLevel {
		if v == nil {
			return nil
		}
		return v.Level
	}).(LogLevelPtrOutput)
}

type FileSystemApplicationLogsConfigResponse struct {
	Level *string `pulumi:"level"`
}


func (val *FileSystemApplicationLogsConfigResponse) Defaults() *FileSystemApplicationLogsConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Level) {
		level_ := "Off"
		tmp.Level = &level_
	}
	return &tmp
}

type FileSystemApplicationLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponseOutput() FileSystemApplicationLogsConfigResponseOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponseOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponseOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FileSystemApplicationLogsConfigResponse) *string { return v.Level }).(pulumi.StringPtrOutput)
}

type FileSystemApplicationLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemApplicationLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) Elem() FileSystemApplicationLogsConfigResponseOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfigResponse) FileSystemApplicationLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemApplicationLogsConfigResponse
		return ret
	}).(FileSystemApplicationLogsConfigResponseOutput)
}

func (o FileSystemApplicationLogsConfigResponsePtrOutput) Level() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FileSystemApplicationLogsConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.Level
	}).(pulumi.StringPtrOutput)
}

type FileSystemHttpLogsConfig struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
	RetentionInMb   *int  `pulumi:"retentionInMb"`
}





type FileSystemHttpLogsConfigInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput
	ToFileSystemHttpLogsConfigOutputWithContext(context.Context) FileSystemHttpLogsConfigOutput
}

type FileSystemHttpLogsConfigArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput  `pulumi:"retentionInDays"`
	RetentionInMb   pulumi.IntPtrInput  `pulumi:"retentionInMb"`
}

func (FileSystemHttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfig)(nil)).Elem()
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput {
	return i.ToFileSystemHttpLogsConfigOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigOutput)
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return i.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigArgs) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigOutput).ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx)
}









type FileSystemHttpLogsConfigPtrInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput
	ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Context) FileSystemHttpLogsConfigPtrOutput
}

type fileSystemHttpLogsConfigPtrType FileSystemHttpLogsConfigArgs

func FileSystemHttpLogsConfigPtr(v *FileSystemHttpLogsConfigArgs) FileSystemHttpLogsConfigPtrInput {
	return (*fileSystemHttpLogsConfigPtrType)(v)
}

func (*fileSystemHttpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfig)(nil)).Elem()
}

func (i *fileSystemHttpLogsConfigPtrType) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return i.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *fileSystemHttpLogsConfigPtrType) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigPtrOutput)
}

type FileSystemHttpLogsConfigOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfig)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigOutput() FileSystemHttpLogsConfigOutput {
	return o
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigOutput {
	return o
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return o.ToFileSystemHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o FileSystemHttpLogsConfigOutput) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemHttpLogsConfig) *FileSystemHttpLogsConfig {
		return &v
	}).(FileSystemHttpLogsConfigPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfig) *int { return v.RetentionInMb }).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfig)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigPtrOutput) ToFileSystemHttpLogsConfigPtrOutput() FileSystemHttpLogsConfigPtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigPtrOutput) ToFileSystemHttpLogsConfigPtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigPtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigPtrOutput) Elem() FileSystemHttpLogsConfigOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) FileSystemHttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret FileSystemHttpLogsConfig
		return ret
	}).(FileSystemHttpLogsConfigOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigPtrOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfig) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInMb
	}).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigResponse struct {
	Enabled         *bool `pulumi:"enabled"`
	RetentionInDays *int  `pulumi:"retentionInDays"`
	RetentionInMb   *int  `pulumi:"retentionInMb"`
}

type FileSystemHttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponseOutput() FileSystemHttpLogsConfigResponseOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponseOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponseOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigResponseOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *int { return v.RetentionInDays }).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigResponseOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v FileSystemHttpLogsConfigResponse) *int { return v.RetentionInMb }).(pulumi.IntPtrOutput)
}

type FileSystemHttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (FileSystemHttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponsePtrOutput {
	return o
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) Elem() FileSystemHttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) FileSystemHttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret FileSystemHttpLogsConfigResponse
		return ret
	}).(FileSystemHttpLogsConfigResponseOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) RetentionInDays() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInDays
	}).(pulumi.IntPtrOutput)
}

func (o FileSystemHttpLogsConfigResponsePtrOutput) RetentionInMb() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *FileSystemHttpLogsConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.RetentionInMb
	}).(pulumi.IntPtrOutput)
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

type GitHubActionCodeConfiguration struct {
	RuntimeStack   *string `pulumi:"runtimeStack"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}





type GitHubActionCodeConfigurationInput interface {
	pulumi.Input

	ToGitHubActionCodeConfigurationOutput() GitHubActionCodeConfigurationOutput
	ToGitHubActionCodeConfigurationOutputWithContext(context.Context) GitHubActionCodeConfigurationOutput
}

type GitHubActionCodeConfigurationArgs struct {
	RuntimeStack   pulumi.StringPtrInput `pulumi:"runtimeStack"`
	RuntimeVersion pulumi.StringPtrInput `pulumi:"runtimeVersion"`
}

func (GitHubActionCodeConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionCodeConfiguration)(nil)).Elem()
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationOutput() GitHubActionCodeConfigurationOutput {
	return i.ToGitHubActionCodeConfigurationOutputWithContext(context.Background())
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionCodeConfigurationOutput)
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return i.ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Background())
}

func (i GitHubActionCodeConfigurationArgs) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionCodeConfigurationOutput).ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx)
}









type GitHubActionCodeConfigurationPtrInput interface {
	pulumi.Input

	ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput
	ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Context) GitHubActionCodeConfigurationPtrOutput
}

type gitHubActionCodeConfigurationPtrType GitHubActionCodeConfigurationArgs

func GitHubActionCodeConfigurationPtr(v *GitHubActionCodeConfigurationArgs) GitHubActionCodeConfigurationPtrInput {
	return (*gitHubActionCodeConfigurationPtrType)(v)
}

func (*gitHubActionCodeConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionCodeConfiguration)(nil)).Elem()
}

func (i *gitHubActionCodeConfigurationPtrType) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return i.ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Background())
}

func (i *gitHubActionCodeConfigurationPtrType) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionCodeConfigurationPtrOutput)
}

type GitHubActionCodeConfigurationOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionCodeConfiguration)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationOutput() GitHubActionCodeConfigurationOutput {
	return o
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationOutput {
	return o
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return o.ToGitHubActionCodeConfigurationPtrOutputWithContext(context.Background())
}

func (o GitHubActionCodeConfigurationOutput) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubActionCodeConfiguration) *GitHubActionCodeConfiguration {
		return &v
	}).(GitHubActionCodeConfigurationPtrOutput)
}

func (o GitHubActionCodeConfigurationOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfiguration) *string { return v.RuntimeStack }).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfiguration) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type GitHubActionCodeConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionCodeConfiguration)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationPtrOutput) ToGitHubActionCodeConfigurationPtrOutput() GitHubActionCodeConfigurationPtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationPtrOutput) ToGitHubActionCodeConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationPtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationPtrOutput) Elem() GitHubActionCodeConfigurationOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfiguration) GitHubActionCodeConfiguration {
		if v != nil {
			return *v
		}
		var ret GitHubActionCodeConfiguration
		return ret
	}).(GitHubActionCodeConfigurationOutput)
}

func (o GitHubActionCodeConfigurationPtrOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeStack
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationPtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GitHubActionCodeConfigurationResponse struct {
	RuntimeStack   *string `pulumi:"runtimeStack"`
	RuntimeVersion *string `pulumi:"runtimeVersion"`
}

type GitHubActionCodeConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionCodeConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationResponseOutput) ToGitHubActionCodeConfigurationResponseOutput() GitHubActionCodeConfigurationResponseOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponseOutput) ToGitHubActionCodeConfigurationResponseOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationResponseOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponseOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfigurationResponse) *string { return v.RuntimeStack }).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationResponseOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionCodeConfigurationResponse) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

type GitHubActionCodeConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubActionCodeConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionCodeConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) ToGitHubActionCodeConfigurationResponsePtrOutput() GitHubActionCodeConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) ToGitHubActionCodeConfigurationResponsePtrOutputWithContext(ctx context.Context) GitHubActionCodeConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) Elem() GitHubActionCodeConfigurationResponseOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfigurationResponse) GitHubActionCodeConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GitHubActionCodeConfigurationResponse
		return ret
	}).(GitHubActionCodeConfigurationResponseOutput)
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) RuntimeStack() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeStack
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionCodeConfigurationResponsePtrOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionCodeConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.RuntimeVersion
	}).(pulumi.StringPtrOutput)
}

type GitHubActionConfiguration struct {
	CodeConfiguration      *GitHubActionCodeConfiguration      `pulumi:"codeConfiguration"`
	ContainerConfiguration *GitHubActionContainerConfiguration `pulumi:"containerConfiguration"`
	GenerateWorkflowFile   *bool                               `pulumi:"generateWorkflowFile"`
	IsLinux                *bool                               `pulumi:"isLinux"`
}





type GitHubActionConfigurationInput interface {
	pulumi.Input

	ToGitHubActionConfigurationOutput() GitHubActionConfigurationOutput
	ToGitHubActionConfigurationOutputWithContext(context.Context) GitHubActionConfigurationOutput
}

type GitHubActionConfigurationArgs struct {
	CodeConfiguration      GitHubActionCodeConfigurationPtrInput      `pulumi:"codeConfiguration"`
	ContainerConfiguration GitHubActionContainerConfigurationPtrInput `pulumi:"containerConfiguration"`
	GenerateWorkflowFile   pulumi.BoolPtrInput                        `pulumi:"generateWorkflowFile"`
	IsLinux                pulumi.BoolPtrInput                        `pulumi:"isLinux"`
}

func (GitHubActionConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionConfiguration)(nil)).Elem()
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationOutput() GitHubActionConfigurationOutput {
	return i.ToGitHubActionConfigurationOutputWithContext(context.Background())
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationOutputWithContext(ctx context.Context) GitHubActionConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionConfigurationOutput)
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return i.ToGitHubActionConfigurationPtrOutputWithContext(context.Background())
}

func (i GitHubActionConfigurationArgs) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionConfigurationOutput).ToGitHubActionConfigurationPtrOutputWithContext(ctx)
}









type GitHubActionConfigurationPtrInput interface {
	pulumi.Input

	ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput
	ToGitHubActionConfigurationPtrOutputWithContext(context.Context) GitHubActionConfigurationPtrOutput
}

type gitHubActionConfigurationPtrType GitHubActionConfigurationArgs

func GitHubActionConfigurationPtr(v *GitHubActionConfigurationArgs) GitHubActionConfigurationPtrInput {
	return (*gitHubActionConfigurationPtrType)(v)
}

func (*gitHubActionConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionConfiguration)(nil)).Elem()
}

func (i *gitHubActionConfigurationPtrType) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return i.ToGitHubActionConfigurationPtrOutputWithContext(context.Background())
}

func (i *gitHubActionConfigurationPtrType) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionConfigurationPtrOutput)
}

type GitHubActionConfigurationOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionConfiguration)(nil)).Elem()
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationOutput() GitHubActionConfigurationOutput {
	return o
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationOutputWithContext(ctx context.Context) GitHubActionConfigurationOutput {
	return o
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return o.ToGitHubActionConfigurationPtrOutputWithContext(context.Background())
}

func (o GitHubActionConfigurationOutput) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubActionConfiguration) *GitHubActionConfiguration {
		return &v
	}).(GitHubActionConfigurationPtrOutput)
}

func (o GitHubActionConfigurationOutput) CodeConfiguration() GitHubActionCodeConfigurationPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *GitHubActionCodeConfiguration { return v.CodeConfiguration }).(GitHubActionCodeConfigurationPtrOutput)
}

func (o GitHubActionConfigurationOutput) ContainerConfiguration() GitHubActionContainerConfigurationPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *GitHubActionContainerConfiguration { return v.ContainerConfiguration }).(GitHubActionContainerConfigurationPtrOutput)
}

func (o GitHubActionConfigurationOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *bool { return v.GenerateWorkflowFile }).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfiguration) *bool { return v.IsLinux }).(pulumi.BoolPtrOutput)
}

type GitHubActionConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionConfiguration)(nil)).Elem()
}

func (o GitHubActionConfigurationPtrOutput) ToGitHubActionConfigurationPtrOutput() GitHubActionConfigurationPtrOutput {
	return o
}

func (o GitHubActionConfigurationPtrOutput) ToGitHubActionConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionConfigurationPtrOutput {
	return o
}

func (o GitHubActionConfigurationPtrOutput) Elem() GitHubActionConfigurationOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) GitHubActionConfiguration {
		if v != nil {
			return *v
		}
		var ret GitHubActionConfiguration
		return ret
	}).(GitHubActionConfigurationOutput)
}

func (o GitHubActionConfigurationPtrOutput) CodeConfiguration() GitHubActionCodeConfigurationPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *GitHubActionCodeConfiguration {
		if v == nil {
			return nil
		}
		return v.CodeConfiguration
	}).(GitHubActionCodeConfigurationPtrOutput)
}

func (o GitHubActionConfigurationPtrOutput) ContainerConfiguration() GitHubActionContainerConfigurationPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *GitHubActionContainerConfiguration {
		if v == nil {
			return nil
		}
		return v.ContainerConfiguration
	}).(GitHubActionContainerConfigurationPtrOutput)
}

func (o GitHubActionConfigurationPtrOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.GenerateWorkflowFile
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationPtrOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfiguration) *bool {
		if v == nil {
			return nil
		}
		return v.IsLinux
	}).(pulumi.BoolPtrOutput)
}

type GitHubActionConfigurationResponse struct {
	CodeConfiguration      *GitHubActionCodeConfigurationResponse      `pulumi:"codeConfiguration"`
	ContainerConfiguration *GitHubActionContainerConfigurationResponse `pulumi:"containerConfiguration"`
	GenerateWorkflowFile   *bool                                       `pulumi:"generateWorkflowFile"`
	IsLinux                *bool                                       `pulumi:"isLinux"`
}

type GitHubActionConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionConfigurationResponseOutput) ToGitHubActionConfigurationResponseOutput() GitHubActionConfigurationResponseOutput {
	return o
}

func (o GitHubActionConfigurationResponseOutput) ToGitHubActionConfigurationResponseOutputWithContext(ctx context.Context) GitHubActionConfigurationResponseOutput {
	return o
}

func (o GitHubActionConfigurationResponseOutput) CodeConfiguration() GitHubActionCodeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *GitHubActionCodeConfigurationResponse {
		return v.CodeConfiguration
	}).(GitHubActionCodeConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponseOutput) ContainerConfiguration() GitHubActionContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *GitHubActionContainerConfigurationResponse {
		return v.ContainerConfiguration
	}).(GitHubActionContainerConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponseOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *bool { return v.GenerateWorkflowFile }).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationResponseOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GitHubActionConfigurationResponse) *bool { return v.IsLinux }).(pulumi.BoolPtrOutput)
}

type GitHubActionConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubActionConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionConfigurationResponsePtrOutput) ToGitHubActionConfigurationResponsePtrOutput() GitHubActionConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionConfigurationResponsePtrOutput) ToGitHubActionConfigurationResponsePtrOutputWithContext(ctx context.Context) GitHubActionConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionConfigurationResponsePtrOutput) Elem() GitHubActionConfigurationResponseOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) GitHubActionConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GitHubActionConfigurationResponse
		return ret
	}).(GitHubActionConfigurationResponseOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) CodeConfiguration() GitHubActionCodeConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *GitHubActionCodeConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.CodeConfiguration
	}).(GitHubActionCodeConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) ContainerConfiguration() GitHubActionContainerConfigurationResponsePtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *GitHubActionContainerConfigurationResponse {
		if v == nil {
			return nil
		}
		return v.ContainerConfiguration
	}).(GitHubActionContainerConfigurationResponsePtrOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) GenerateWorkflowFile() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.GenerateWorkflowFile
	}).(pulumi.BoolPtrOutput)
}

func (o GitHubActionConfigurationResponsePtrOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *GitHubActionConfigurationResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsLinux
	}).(pulumi.BoolPtrOutput)
}

type GitHubActionContainerConfiguration struct {
	ImageName *string `pulumi:"imageName"`
	Password  *string `pulumi:"password"`
	ServerUrl *string `pulumi:"serverUrl"`
	Username  *string `pulumi:"username"`
}





type GitHubActionContainerConfigurationInput interface {
	pulumi.Input

	ToGitHubActionContainerConfigurationOutput() GitHubActionContainerConfigurationOutput
	ToGitHubActionContainerConfigurationOutputWithContext(context.Context) GitHubActionContainerConfigurationOutput
}

type GitHubActionContainerConfigurationArgs struct {
	ImageName pulumi.StringPtrInput `pulumi:"imageName"`
	Password  pulumi.StringPtrInput `pulumi:"password"`
	ServerUrl pulumi.StringPtrInput `pulumi:"serverUrl"`
	Username  pulumi.StringPtrInput `pulumi:"username"`
}

func (GitHubActionContainerConfigurationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionContainerConfiguration)(nil)).Elem()
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationOutput() GitHubActionContainerConfigurationOutput {
	return i.ToGitHubActionContainerConfigurationOutputWithContext(context.Background())
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionContainerConfigurationOutput)
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return i.ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i GitHubActionContainerConfigurationArgs) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionContainerConfigurationOutput).ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx)
}









type GitHubActionContainerConfigurationPtrInput interface {
	pulumi.Input

	ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput
	ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Context) GitHubActionContainerConfigurationPtrOutput
}

type gitHubActionContainerConfigurationPtrType GitHubActionContainerConfigurationArgs

func GitHubActionContainerConfigurationPtr(v *GitHubActionContainerConfigurationArgs) GitHubActionContainerConfigurationPtrInput {
	return (*gitHubActionContainerConfigurationPtrType)(v)
}

func (*gitHubActionContainerConfigurationPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionContainerConfiguration)(nil)).Elem()
}

func (i *gitHubActionContainerConfigurationPtrType) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return i.ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Background())
}

func (i *gitHubActionContainerConfigurationPtrType) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(GitHubActionContainerConfigurationPtrOutput)
}

type GitHubActionContainerConfigurationOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionContainerConfiguration)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationOutput() GitHubActionContainerConfigurationOutput {
	return o
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationOutput {
	return o
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return o.ToGitHubActionContainerConfigurationPtrOutputWithContext(context.Background())
}

func (o GitHubActionContainerConfigurationOutput) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v GitHubActionContainerConfiguration) *GitHubActionContainerConfiguration {
		return &v
	}).(GitHubActionContainerConfigurationPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfiguration) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GitHubActionContainerConfigurationPtrOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionContainerConfiguration)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationPtrOutput) ToGitHubActionContainerConfigurationPtrOutput() GitHubActionContainerConfigurationPtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationPtrOutput) ToGitHubActionContainerConfigurationPtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationPtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationPtrOutput) Elem() GitHubActionContainerConfigurationOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) GitHubActionContainerConfiguration {
		if v != nil {
			return *v
		}
		var ret GitHubActionContainerConfiguration
		return ret
	}).(GitHubActionContainerConfigurationOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationPtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfiguration) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type GitHubActionContainerConfigurationResponse struct {
	ImageName *string `pulumi:"imageName"`
	Password  *string `pulumi:"password"`
	ServerUrl *string `pulumi:"serverUrl"`
	Username  *string `pulumi:"username"`
}

type GitHubActionContainerConfigurationResponseOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GitHubActionContainerConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationResponseOutput) ToGitHubActionContainerConfigurationResponseOutput() GitHubActionContainerConfigurationResponseOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponseOutput) ToGitHubActionContainerConfigurationResponseOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationResponseOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponseOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.ImageName }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponseOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.ServerUrl }).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponseOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GitHubActionContainerConfigurationResponse) *string { return v.Username }).(pulumi.StringPtrOutput)
}

type GitHubActionContainerConfigurationResponsePtrOutput struct{ *pulumi.OutputState }

func (GitHubActionContainerConfigurationResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**GitHubActionContainerConfigurationResponse)(nil)).Elem()
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ToGitHubActionContainerConfigurationResponsePtrOutput() GitHubActionContainerConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ToGitHubActionContainerConfigurationResponsePtrOutputWithContext(ctx context.Context) GitHubActionContainerConfigurationResponsePtrOutput {
	return o
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) Elem() GitHubActionContainerConfigurationResponseOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) GitHubActionContainerConfigurationResponse {
		if v != nil {
			return *v
		}
		var ret GitHubActionContainerConfigurationResponse
		return ret
	}).(GitHubActionContainerConfigurationResponseOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ImageName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ImageName
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) ServerUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServerUrl
	}).(pulumi.StringPtrOutput)
}

func (o GitHubActionContainerConfigurationResponsePtrOutput) Username() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *GitHubActionContainerConfigurationResponse) *string {
		if v == nil {
			return nil
		}
		return v.Username
	}).(pulumi.StringPtrOutput)
}

type HandlerMapping struct {
	Arguments       *string `pulumi:"arguments"`
	Extension       *string `pulumi:"extension"`
	ScriptProcessor *string `pulumi:"scriptProcessor"`
}





type HandlerMappingInput interface {
	pulumi.Input

	ToHandlerMappingOutput() HandlerMappingOutput
	ToHandlerMappingOutputWithContext(context.Context) HandlerMappingOutput
}

type HandlerMappingArgs struct {
	Arguments       pulumi.StringPtrInput `pulumi:"arguments"`
	Extension       pulumi.StringPtrInput `pulumi:"extension"`
	ScriptProcessor pulumi.StringPtrInput `pulumi:"scriptProcessor"`
}

func (HandlerMappingArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMapping)(nil)).Elem()
}

func (i HandlerMappingArgs) ToHandlerMappingOutput() HandlerMappingOutput {
	return i.ToHandlerMappingOutputWithContext(context.Background())
}

func (i HandlerMappingArgs) ToHandlerMappingOutputWithContext(ctx context.Context) HandlerMappingOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingOutput)
}





type HandlerMappingArrayInput interface {
	pulumi.Input

	ToHandlerMappingArrayOutput() HandlerMappingArrayOutput
	ToHandlerMappingArrayOutputWithContext(context.Context) HandlerMappingArrayOutput
}

type HandlerMappingArray []HandlerMappingInput

func (HandlerMappingArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMapping)(nil)).Elem()
}

func (i HandlerMappingArray) ToHandlerMappingArrayOutput() HandlerMappingArrayOutput {
	return i.ToHandlerMappingArrayOutputWithContext(context.Background())
}

func (i HandlerMappingArray) ToHandlerMappingArrayOutputWithContext(ctx context.Context) HandlerMappingArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingArrayOutput)
}

type HandlerMappingOutput struct{ *pulumi.OutputState }

func (HandlerMappingOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMapping)(nil)).Elem()
}

func (o HandlerMappingOutput) ToHandlerMappingOutput() HandlerMappingOutput {
	return o
}

func (o HandlerMappingOutput) ToHandlerMappingOutputWithContext(ctx context.Context) HandlerMappingOutput {
	return o
}

func (o HandlerMappingOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingOutput) ScriptProcessor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMapping) *string { return v.ScriptProcessor }).(pulumi.StringPtrOutput)
}

type HandlerMappingArrayOutput struct{ *pulumi.OutputState }

func (HandlerMappingArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMapping)(nil)).Elem()
}

func (o HandlerMappingArrayOutput) ToHandlerMappingArrayOutput() HandlerMappingArrayOutput {
	return o
}

func (o HandlerMappingArrayOutput) ToHandlerMappingArrayOutputWithContext(ctx context.Context) HandlerMappingArrayOutput {
	return o
}

func (o HandlerMappingArrayOutput) Index(i pulumi.IntInput) HandlerMappingOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HandlerMapping {
		return vs[0].([]HandlerMapping)[vs[1].(int)]
	}).(HandlerMappingOutput)
}

type HandlerMappingResponse struct {
	Arguments       *string `pulumi:"arguments"`
	Extension       *string `pulumi:"extension"`
	ScriptProcessor *string `pulumi:"scriptProcessor"`
}

type HandlerMappingResponseOutput struct{ *pulumi.OutputState }

func (HandlerMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMappingResponse)(nil)).Elem()
}

func (o HandlerMappingResponseOutput) ToHandlerMappingResponseOutput() HandlerMappingResponseOutput {
	return o
}

func (o HandlerMappingResponseOutput) ToHandlerMappingResponseOutputWithContext(ctx context.Context) HandlerMappingResponseOutput {
	return o
}

func (o HandlerMappingResponseOutput) Arguments() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.Arguments }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingResponseOutput) Extension() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.Extension }).(pulumi.StringPtrOutput)
}

func (o HandlerMappingResponseOutput) ScriptProcessor() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HandlerMappingResponse) *string { return v.ScriptProcessor }).(pulumi.StringPtrOutput)
}

type HandlerMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (HandlerMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMappingResponse)(nil)).Elem()
}

func (o HandlerMappingResponseArrayOutput) ToHandlerMappingResponseArrayOutput() HandlerMappingResponseArrayOutput {
	return o
}

func (o HandlerMappingResponseArrayOutput) ToHandlerMappingResponseArrayOutputWithContext(ctx context.Context) HandlerMappingResponseArrayOutput {
	return o
}

func (o HandlerMappingResponseArrayOutput) Index(i pulumi.IntInput) HandlerMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HandlerMappingResponse {
		return vs[0].([]HandlerMappingResponse)[vs[1].(int)]
	}).(HandlerMappingResponseOutput)
}

type HostNameSslState struct {
	HostType   *HostType `pulumi:"hostType"`
	Name       *string   `pulumi:"name"`
	SslState   *SslState `pulumi:"sslState"`
	Thumbprint *string   `pulumi:"thumbprint"`
	ToUpdate   *bool     `pulumi:"toUpdate"`
	VirtualIP  *string   `pulumi:"virtualIP"`
}





type HostNameSslStateInput interface {
	pulumi.Input

	ToHostNameSslStateOutput() HostNameSslStateOutput
	ToHostNameSslStateOutputWithContext(context.Context) HostNameSslStateOutput
}

type HostNameSslStateArgs struct {
	HostType   HostTypePtrInput      `pulumi:"hostType"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	SslState   SslStatePtrInput      `pulumi:"sslState"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
	ToUpdate   pulumi.BoolPtrInput   `pulumi:"toUpdate"`
	VirtualIP  pulumi.StringPtrInput `pulumi:"virtualIP"`
}

func (HostNameSslStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslState)(nil)).Elem()
}

func (i HostNameSslStateArgs) ToHostNameSslStateOutput() HostNameSslStateOutput {
	return i.ToHostNameSslStateOutputWithContext(context.Background())
}

func (i HostNameSslStateArgs) ToHostNameSslStateOutputWithContext(ctx context.Context) HostNameSslStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateOutput)
}





type HostNameSslStateArrayInput interface {
	pulumi.Input

	ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput
	ToHostNameSslStateArrayOutputWithContext(context.Context) HostNameSslStateArrayOutput
}

type HostNameSslStateArray []HostNameSslStateInput

func (HostNameSslStateArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslState)(nil)).Elem()
}

func (i HostNameSslStateArray) ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput {
	return i.ToHostNameSslStateArrayOutputWithContext(context.Background())
}

func (i HostNameSslStateArray) ToHostNameSslStateArrayOutputWithContext(ctx context.Context) HostNameSslStateArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateArrayOutput)
}

type HostNameSslStateOutput struct{ *pulumi.OutputState }

func (HostNameSslStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslState)(nil)).Elem()
}

func (o HostNameSslStateOutput) ToHostNameSslStateOutput() HostNameSslStateOutput {
	return o
}

func (o HostNameSslStateOutput) ToHostNameSslStateOutputWithContext(ctx context.Context) HostNameSslStateOutput {
	return o
}

func (o HostNameSslStateOutput) HostType() HostTypePtrOutput {
	return o.ApplyT(func(v HostNameSslState) *HostType { return v.HostType }).(HostTypePtrOutput)
}

func (o HostNameSslStateOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateOutput) SslState() SslStatePtrOutput {
	return o.ApplyT(func(v HostNameSslState) *SslState { return v.SslState }).(SslStatePtrOutput)
}

func (o HostNameSslStateOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateOutput) ToUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *bool { return v.ToUpdate }).(pulumi.BoolPtrOutput)
}

func (o HostNameSslStateOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslState) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type HostNameSslStateArrayOutput struct{ *pulumi.OutputState }

func (HostNameSslStateArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslState)(nil)).Elem()
}

func (o HostNameSslStateArrayOutput) ToHostNameSslStateArrayOutput() HostNameSslStateArrayOutput {
	return o
}

func (o HostNameSslStateArrayOutput) ToHostNameSslStateArrayOutputWithContext(ctx context.Context) HostNameSslStateArrayOutput {
	return o
}

func (o HostNameSslStateArrayOutput) Index(i pulumi.IntInput) HostNameSslStateOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameSslState {
		return vs[0].([]HostNameSslState)[vs[1].(int)]
	}).(HostNameSslStateOutput)
}

type HostNameSslStateResponse struct {
	HostType   *string `pulumi:"hostType"`
	Name       *string `pulumi:"name"`
	SslState   *string `pulumi:"sslState"`
	Thumbprint *string `pulumi:"thumbprint"`
	ToUpdate   *bool   `pulumi:"toUpdate"`
	VirtualIP  *string `pulumi:"virtualIP"`
}

type HostNameSslStateResponseOutput struct{ *pulumi.OutputState }

func (HostNameSslStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslStateResponse)(nil)).Elem()
}

func (o HostNameSslStateResponseOutput) ToHostNameSslStateResponseOutput() HostNameSslStateResponseOutput {
	return o
}

func (o HostNameSslStateResponseOutput) ToHostNameSslStateResponseOutputWithContext(ctx context.Context) HostNameSslStateResponseOutput {
	return o
}

func (o HostNameSslStateResponseOutput) HostType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.HostType }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) SslState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.SslState }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) Thumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.Thumbprint }).(pulumi.StringPtrOutput)
}

func (o HostNameSslStateResponseOutput) ToUpdate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *bool { return v.ToUpdate }).(pulumi.BoolPtrOutput)
}

func (o HostNameSslStateResponseOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostNameSslStateResponse) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type HostNameSslStateResponseArrayOutput struct{ *pulumi.OutputState }

func (HostNameSslStateResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslStateResponse)(nil)).Elem()
}

func (o HostNameSslStateResponseArrayOutput) ToHostNameSslStateResponseArrayOutput() HostNameSslStateResponseArrayOutput {
	return o
}

func (o HostNameSslStateResponseArrayOutput) ToHostNameSslStateResponseArrayOutputWithContext(ctx context.Context) HostNameSslStateResponseArrayOutput {
	return o
}

func (o HostNameSslStateResponseArrayOutput) Index(i pulumi.IntInput) HostNameSslStateResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) HostNameSslStateResponse {
		return vs[0].([]HostNameSslStateResponse)[vs[1].(int)]
	}).(HostNameSslStateResponseOutput)
}

type HostingEnvironmentProfile struct {
	Id *string `pulumi:"id"`
}





type HostingEnvironmentProfileInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput
	ToHostingEnvironmentProfileOutputWithContext(context.Context) HostingEnvironmentProfileOutput
}

type HostingEnvironmentProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (HostingEnvironmentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return i.ToHostingEnvironmentProfileOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput)
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileArgs) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileOutput).ToHostingEnvironmentProfilePtrOutputWithContext(ctx)
}









type HostingEnvironmentProfilePtrInput interface {
	pulumi.Input

	ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput
	ToHostingEnvironmentProfilePtrOutputWithContext(context.Context) HostingEnvironmentProfilePtrOutput
}

type hostingEnvironmentProfilePtrType HostingEnvironmentProfileArgs

func HostingEnvironmentProfilePtr(v *HostingEnvironmentProfileArgs) HostingEnvironmentProfilePtrInput {
	return (*hostingEnvironmentProfilePtrType)(v)
}

func (*hostingEnvironmentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return i.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i *hostingEnvironmentProfilePtrType) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfilePtrOutput)
}

type HostingEnvironmentProfileOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutput() HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfileOutputWithContext(ctx context.Context) HostingEnvironmentProfileOutput {
	return o
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o.ToHostingEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentProfileOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentProfile) *HostingEnvironmentProfile {
		return &v
	}).(HostingEnvironmentProfilePtrOutput)
}

func (o HostingEnvironmentProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfilePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfile)(nil)).Elem()
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutput() HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) ToHostingEnvironmentProfilePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfilePtrOutput {
	return o
}

func (o HostingEnvironmentProfilePtrOutput) Elem() HostingEnvironmentProfileOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) HostingEnvironmentProfile {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfile
		return ret
	}).(HostingEnvironmentProfileOutput)
}

func (o HostingEnvironmentProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type HostingEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type HostingEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return o
}

func (o HostingEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o HostingEnvironmentProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v HostingEnvironmentProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type HostingEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (HostingEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o
}

func (o HostingEnvironmentProfileResponsePtrOutput) Elem() HostingEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) HostingEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret HostingEnvironmentProfileResponse
		return ret
	}).(HostingEnvironmentProfileResponseOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o HostingEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *HostingEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type HttpLogsConfig struct {
	AzureBlobStorage *AzureBlobStorageHttpLogsConfig `pulumi:"azureBlobStorage"`
	FileSystem       *FileSystemHttpLogsConfig       `pulumi:"fileSystem"`
}





type HttpLogsConfigInput interface {
	pulumi.Input

	ToHttpLogsConfigOutput() HttpLogsConfigOutput
	ToHttpLogsConfigOutputWithContext(context.Context) HttpLogsConfigOutput
}

type HttpLogsConfigArgs struct {
	AzureBlobStorage AzureBlobStorageHttpLogsConfigPtrInput `pulumi:"azureBlobStorage"`
	FileSystem       FileSystemHttpLogsConfigPtrInput       `pulumi:"fileSystem"`
}

func (HttpLogsConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfig)(nil)).Elem()
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigOutput() HttpLogsConfigOutput {
	return i.ToHttpLogsConfigOutputWithContext(context.Background())
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigOutputWithContext(ctx context.Context) HttpLogsConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigOutput)
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return i.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i HttpLogsConfigArgs) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigOutput).ToHttpLogsConfigPtrOutputWithContext(ctx)
}









type HttpLogsConfigPtrInput interface {
	pulumi.Input

	ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput
	ToHttpLogsConfigPtrOutputWithContext(context.Context) HttpLogsConfigPtrOutput
}

type httpLogsConfigPtrType HttpLogsConfigArgs

func HttpLogsConfigPtr(v *HttpLogsConfigArgs) HttpLogsConfigPtrInput {
	return (*httpLogsConfigPtrType)(v)
}

func (*httpLogsConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfig)(nil)).Elem()
}

func (i *httpLogsConfigPtrType) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return i.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (i *httpLogsConfigPtrType) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigPtrOutput)
}

type HttpLogsConfigOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfig)(nil)).Elem()
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigOutput() HttpLogsConfigOutput {
	return o
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigOutputWithContext(ctx context.Context) HttpLogsConfigOutput {
	return o
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return o.ToHttpLogsConfigPtrOutputWithContext(context.Background())
}

func (o HttpLogsConfigOutput) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpLogsConfig) *HttpLogsConfig {
		return &v
	}).(HttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v HttpLogsConfig) *AzureBlobStorageHttpLogsConfig { return v.AzureBlobStorage }).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigOutput) FileSystem() FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v HttpLogsConfig) *FileSystemHttpLogsConfig { return v.FileSystem }).(FileSystemHttpLogsConfigPtrOutput)
}

type HttpLogsConfigPtrOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfig)(nil)).Elem()
}

func (o HttpLogsConfigPtrOutput) ToHttpLogsConfigPtrOutput() HttpLogsConfigPtrOutput {
	return o
}

func (o HttpLogsConfigPtrOutput) ToHttpLogsConfigPtrOutputWithContext(ctx context.Context) HttpLogsConfigPtrOutput {
	return o
}

func (o HttpLogsConfigPtrOutput) Elem() HttpLogsConfigOutput {
	return o.ApplyT(func(v *HttpLogsConfig) HttpLogsConfig {
		if v != nil {
			return *v
		}
		var ret HttpLogsConfig
		return ret
	}).(HttpLogsConfigOutput)
}

func (o HttpLogsConfigPtrOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v *HttpLogsConfig) *AzureBlobStorageHttpLogsConfig {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageHttpLogsConfigPtrOutput)
}

func (o HttpLogsConfigPtrOutput) FileSystem() FileSystemHttpLogsConfigPtrOutput {
	return o.ApplyT(func(v *HttpLogsConfig) *FileSystemHttpLogsConfig {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemHttpLogsConfigPtrOutput)
}

type HttpLogsConfigResponse struct {
	AzureBlobStorage *AzureBlobStorageHttpLogsConfigResponse `pulumi:"azureBlobStorage"`
	FileSystem       *FileSystemHttpLogsConfigResponse       `pulumi:"fileSystem"`
}

type HttpLogsConfigResponseOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfigResponse)(nil)).Elem()
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponseOutput() HttpLogsConfigResponseOutput {
	return o
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponseOutputWithContext(ctx context.Context) HttpLogsConfigResponseOutput {
	return o
}

func (o HttpLogsConfigResponseOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v HttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse { return v.AzureBlobStorage }).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
}

func (o HttpLogsConfigResponseOutput) FileSystem() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v HttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse { return v.FileSystem }).(FileSystemHttpLogsConfigResponsePtrOutput)
}

type HttpLogsConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (HttpLogsConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfigResponse)(nil)).Elem()
}

func (o HttpLogsConfigResponsePtrOutput) ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput {
	return o
}

func (o HttpLogsConfigResponsePtrOutput) ToHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) HttpLogsConfigResponsePtrOutput {
	return o
}

func (o HttpLogsConfigResponsePtrOutput) Elem() HttpLogsConfigResponseOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) HttpLogsConfigResponse {
		if v != nil {
			return *v
		}
		var ret HttpLogsConfigResponse
		return ret
	}).(HttpLogsConfigResponseOutput)
}

func (o HttpLogsConfigResponsePtrOutput) AzureBlobStorage() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.AzureBlobStorage
	}).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
}

func (o HttpLogsConfigResponsePtrOutput) FileSystem() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyT(func(v *HttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse {
		if v == nil {
			return nil
		}
		return v.FileSystem
	}).(FileSystemHttpLogsConfigResponsePtrOutput)
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

type IdentifierResponse struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}

type IdentifierResponseOutput struct{ *pulumi.OutputState }

func (IdentifierResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentifierResponse)(nil)).Elem()
}

func (o IdentifierResponseOutput) ToIdentifierResponseOutput() IdentifierResponseOutput {
	return o
}

func (o IdentifierResponseOutput) ToIdentifierResponseOutputWithContext(ctx context.Context) IdentifierResponseOutput {
	return o
}

func (o IdentifierResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentifierResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o IdentifierResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v IdentifierResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o IdentifierResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IdentifierResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type IdentifierResponseArrayOutput struct{ *pulumi.OutputState }

func (IdentifierResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentifierResponse)(nil)).Elem()
}

func (o IdentifierResponseArrayOutput) ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput {
	return o
}

func (o IdentifierResponseArrayOutput) ToIdentifierResponseArrayOutputWithContext(ctx context.Context) IdentifierResponseArrayOutput {
	return o
}

func (o IdentifierResponseArrayOutput) Index(i pulumi.IntInput) IdentifierResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IdentifierResponse {
		return vs[0].([]IdentifierResponse)[vs[1].(int)]
	}).(IdentifierResponseOutput)
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


func (val *IngressArgs) Defaults() *IngressArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.External) {
		tmp.External = pulumi.BoolPtr(false)
	}
	return &tmp
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

type IpSecurityRestriction struct {
	Action               *string             `pulumi:"action"`
	Description          *string             `pulumi:"description"`
	Headers              map[string][]string `pulumi:"headers"`
	IpAddress            *string             `pulumi:"ipAddress"`
	Name                 *string             `pulumi:"name"`
	Priority             *int                `pulumi:"priority"`
	SubnetMask           *string             `pulumi:"subnetMask"`
	SubnetTrafficTag     *int                `pulumi:"subnetTrafficTag"`
	Tag                  *string             `pulumi:"tag"`
	VnetSubnetResourceId *string             `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       *int                `pulumi:"vnetTrafficTag"`
}





type IpSecurityRestrictionInput interface {
	pulumi.Input

	ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput
	ToIpSecurityRestrictionOutputWithContext(context.Context) IpSecurityRestrictionOutput
}

type IpSecurityRestrictionArgs struct {
	Action               pulumi.StringPtrInput      `pulumi:"action"`
	Description          pulumi.StringPtrInput      `pulumi:"description"`
	Headers              pulumi.StringArrayMapInput `pulumi:"headers"`
	IpAddress            pulumi.StringPtrInput      `pulumi:"ipAddress"`
	Name                 pulumi.StringPtrInput      `pulumi:"name"`
	Priority             pulumi.IntPtrInput         `pulumi:"priority"`
	SubnetMask           pulumi.StringPtrInput      `pulumi:"subnetMask"`
	SubnetTrafficTag     pulumi.IntPtrInput         `pulumi:"subnetTrafficTag"`
	Tag                  pulumi.StringPtrInput      `pulumi:"tag"`
	VnetSubnetResourceId pulumi.StringPtrInput      `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       pulumi.IntPtrInput         `pulumi:"vnetTrafficTag"`
}

func (IpSecurityRestrictionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestriction)(nil)).Elem()
}

func (i IpSecurityRestrictionArgs) ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput {
	return i.ToIpSecurityRestrictionOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionArgs) ToIpSecurityRestrictionOutputWithContext(ctx context.Context) IpSecurityRestrictionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionOutput)
}





type IpSecurityRestrictionArrayInput interface {
	pulumi.Input

	ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput
	ToIpSecurityRestrictionArrayOutputWithContext(context.Context) IpSecurityRestrictionArrayOutput
}

type IpSecurityRestrictionArray []IpSecurityRestrictionInput

func (IpSecurityRestrictionArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestriction)(nil)).Elem()
}

func (i IpSecurityRestrictionArray) ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput {
	return i.ToIpSecurityRestrictionArrayOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionArray) ToIpSecurityRestrictionArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionArrayOutput)
}

type IpSecurityRestrictionOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestriction)(nil)).Elem()
}

func (o IpSecurityRestrictionOutput) ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput {
	return o
}

func (o IpSecurityRestrictionOutput) ToIpSecurityRestrictionOutputWithContext(ctx context.Context) IpSecurityRestrictionOutput {
	return o
}

func (o IpSecurityRestrictionOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Headers() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v IpSecurityRestriction) map[string][]string { return v.Headers }).(pulumi.StringArrayMapOutput)
}

func (o IpSecurityRestrictionOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) SubnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *int { return v.SubnetTrafficTag }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) VnetSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *string { return v.VnetSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionOutput) VnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *int { return v.VnetTrafficTag }).(pulumi.IntPtrOutput)
}

type IpSecurityRestrictionArrayOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestriction)(nil)).Elem()
}

func (o IpSecurityRestrictionArrayOutput) ToIpSecurityRestrictionArrayOutput() IpSecurityRestrictionArrayOutput {
	return o
}

func (o IpSecurityRestrictionArrayOutput) ToIpSecurityRestrictionArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionArrayOutput {
	return o
}

func (o IpSecurityRestrictionArrayOutput) Index(i pulumi.IntInput) IpSecurityRestrictionOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecurityRestriction {
		return vs[0].([]IpSecurityRestriction)[vs[1].(int)]
	}).(IpSecurityRestrictionOutput)
}

type IpSecurityRestrictionResponse struct {
	Action               *string             `pulumi:"action"`
	Description          *string             `pulumi:"description"`
	Headers              map[string][]string `pulumi:"headers"`
	IpAddress            *string             `pulumi:"ipAddress"`
	Name                 *string             `pulumi:"name"`
	Priority             *int                `pulumi:"priority"`
	SubnetMask           *string             `pulumi:"subnetMask"`
	SubnetTrafficTag     *int                `pulumi:"subnetTrafficTag"`
	Tag                  *string             `pulumi:"tag"`
	VnetSubnetResourceId *string             `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       *int                `pulumi:"vnetTrafficTag"`
}

type IpSecurityRestrictionResponseOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestrictionResponse)(nil)).Elem()
}

func (o IpSecurityRestrictionResponseOutput) ToIpSecurityRestrictionResponseOutput() IpSecurityRestrictionResponseOutput {
	return o
}

func (o IpSecurityRestrictionResponseOutput) ToIpSecurityRestrictionResponseOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseOutput {
	return o
}

func (o IpSecurityRestrictionResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Headers() pulumi.StringArrayMapOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) map[string][]string { return v.Headers }).(pulumi.StringArrayMapOutput)
}

func (o IpSecurityRestrictionResponseOutput) IpAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.IpAddress }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Priority() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *int { return v.Priority }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) SubnetMask() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.SubnetMask }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) SubnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *int { return v.SubnetTrafficTag }).(pulumi.IntPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) VnetSubnetResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *string { return v.VnetSubnetResourceId }).(pulumi.StringPtrOutput)
}

func (o IpSecurityRestrictionResponseOutput) VnetTrafficTag() pulumi.IntPtrOutput {
	return o.ApplyT(func(v IpSecurityRestrictionResponse) *int { return v.VnetTrafficTag }).(pulumi.IntPtrOutput)
}

type IpSecurityRestrictionResponseArrayOutput struct{ *pulumi.OutputState }

func (IpSecurityRestrictionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestrictionResponse)(nil)).Elem()
}

func (o IpSecurityRestrictionResponseArrayOutput) ToIpSecurityRestrictionResponseArrayOutput() IpSecurityRestrictionResponseArrayOutput {
	return o
}

func (o IpSecurityRestrictionResponseArrayOutput) ToIpSecurityRestrictionResponseArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseArrayOutput {
	return o
}

func (o IpSecurityRestrictionResponseArrayOutput) Index(i pulumi.IntInput) IpSecurityRestrictionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) IpSecurityRestrictionResponse {
		return vs[0].([]IpSecurityRestrictionResponse)[vs[1].(int)]
	}).(IpSecurityRestrictionResponseOutput)
}

type KubeEnvironmentProfile struct {
	Id *string `pulumi:"id"`
}





type KubeEnvironmentProfileInput interface {
	pulumi.Input

	ToKubeEnvironmentProfileOutput() KubeEnvironmentProfileOutput
	ToKubeEnvironmentProfileOutputWithContext(context.Context) KubeEnvironmentProfileOutput
}

type KubeEnvironmentProfileArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (KubeEnvironmentProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironmentProfile)(nil)).Elem()
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfileOutput() KubeEnvironmentProfileOutput {
	return i.ToKubeEnvironmentProfileOutputWithContext(context.Background())
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfileOutputWithContext(ctx context.Context) KubeEnvironmentProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentProfileOutput)
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return i.ToKubeEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i KubeEnvironmentProfileArgs) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentProfileOutput).ToKubeEnvironmentProfilePtrOutputWithContext(ctx)
}









type KubeEnvironmentProfilePtrInput interface {
	pulumi.Input

	ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput
	ToKubeEnvironmentProfilePtrOutputWithContext(context.Context) KubeEnvironmentProfilePtrOutput
}

type kubeEnvironmentProfilePtrType KubeEnvironmentProfileArgs

func KubeEnvironmentProfilePtr(v *KubeEnvironmentProfileArgs) KubeEnvironmentProfilePtrInput {
	return (*kubeEnvironmentProfilePtrType)(v)
}

func (*kubeEnvironmentProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironmentProfile)(nil)).Elem()
}

func (i *kubeEnvironmentProfilePtrType) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return i.ToKubeEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (i *kubeEnvironmentProfilePtrType) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KubeEnvironmentProfilePtrOutput)
}

type KubeEnvironmentProfileOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironmentProfile)(nil)).Elem()
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfileOutput() KubeEnvironmentProfileOutput {
	return o
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfileOutputWithContext(ctx context.Context) KubeEnvironmentProfileOutput {
	return o
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return o.ToKubeEnvironmentProfilePtrOutputWithContext(context.Background())
}

func (o KubeEnvironmentProfileOutput) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KubeEnvironmentProfile) *KubeEnvironmentProfile {
		return &v
	}).(KubeEnvironmentProfilePtrOutput)
}

func (o KubeEnvironmentProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeEnvironmentProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

type KubeEnvironmentProfilePtrOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironmentProfile)(nil)).Elem()
}

func (o KubeEnvironmentProfilePtrOutput) ToKubeEnvironmentProfilePtrOutput() KubeEnvironmentProfilePtrOutput {
	return o
}

func (o KubeEnvironmentProfilePtrOutput) ToKubeEnvironmentProfilePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfilePtrOutput {
	return o
}

func (o KubeEnvironmentProfilePtrOutput) Elem() KubeEnvironmentProfileOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfile) KubeEnvironmentProfile {
		if v != nil {
			return *v
		}
		var ret KubeEnvironmentProfile
		return ret
	}).(KubeEnvironmentProfileOutput)
}

func (o KubeEnvironmentProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

type KubeEnvironmentProfileResponse struct {
	Id   *string `pulumi:"id"`
	Name string  `pulumi:"name"`
	Type string  `pulumi:"type"`
}

type KubeEnvironmentProfileResponseOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KubeEnvironmentProfileResponse)(nil)).Elem()
}

func (o KubeEnvironmentProfileResponseOutput) ToKubeEnvironmentProfileResponseOutput() KubeEnvironmentProfileResponseOutput {
	return o
}

func (o KubeEnvironmentProfileResponseOutput) ToKubeEnvironmentProfileResponseOutputWithContext(ctx context.Context) KubeEnvironmentProfileResponseOutput {
	return o
}

func (o KubeEnvironmentProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KubeEnvironmentProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v KubeEnvironmentProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o KubeEnvironmentProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v KubeEnvironmentProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type KubeEnvironmentProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (KubeEnvironmentProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KubeEnvironmentProfileResponse)(nil)).Elem()
}

func (o KubeEnvironmentProfileResponsePtrOutput) ToKubeEnvironmentProfileResponsePtrOutput() KubeEnvironmentProfileResponsePtrOutput {
	return o
}

func (o KubeEnvironmentProfileResponsePtrOutput) ToKubeEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) KubeEnvironmentProfileResponsePtrOutput {
	return o
}

func (o KubeEnvironmentProfileResponsePtrOutput) Elem() KubeEnvironmentProfileResponseOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) KubeEnvironmentProfileResponse {
		if v != nil {
			return *v
		}
		var ret KubeEnvironmentProfileResponse
		return ret
	}).(KubeEnvironmentProfileResponseOutput)
}

func (o KubeEnvironmentProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o KubeEnvironmentProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KubeEnvironmentProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
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

type ManagedServiceIdentity struct {
	Type                   *ManagedServiceIdentityType `pulumi:"type"`
	UserAssignedIdentities map[string]interface{}      `pulumi:"userAssignedIdentities"`
}





type ManagedServiceIdentityInput interface {
	pulumi.Input

	ToManagedServiceIdentityOutput() ManagedServiceIdentityOutput
	ToManagedServiceIdentityOutputWithContext(context.Context) ManagedServiceIdentityOutput
}

type ManagedServiceIdentityArgs struct {
	Type                   ManagedServiceIdentityTypePtrInput `pulumi:"type"`
	UserAssignedIdentities pulumi.MapInput                    `pulumi:"userAssignedIdentities"`
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

func (o ManagedServiceIdentityOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentity) *ManagedServiceIdentityType { return v.Type }).(ManagedServiceIdentityTypePtrOutput)
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

func (o ManagedServiceIdentityPtrOutput) Type() ManagedServiceIdentityTypePtrOutput {
	return o.ApplyT(func(v *ManagedServiceIdentity) *ManagedServiceIdentityType {
		if v == nil {
			return nil
		}
		return v.Type
	}).(ManagedServiceIdentityTypePtrOutput)
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
	PrincipalId            string                                  `pulumi:"principalId"`
	TenantId               string                                  `pulumi:"tenantId"`
	Type                   *string                                 `pulumi:"type"`
	UserAssignedIdentities map[string]UserAssignedIdentityResponse `pulumi:"userAssignedIdentities"`
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

func (o ManagedServiceIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) TenantId() pulumi.StringOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) string { return v.TenantId }).(pulumi.StringOutput)
}

func (o ManagedServiceIdentityResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ManagedServiceIdentityResponseOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
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

func (o ManagedServiceIdentityResponsePtrOutput) UserAssignedIdentities() UserAssignedIdentityResponseMapOutput {
	return o.ApplyT(func(v *ManagedServiceIdentityResponse) map[string]UserAssignedIdentityResponse {
		if v == nil {
			return nil
		}
		return v.UserAssignedIdentities
	}).(UserAssignedIdentityResponseMapOutput)
}

type NameValuePair struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}





type NameValuePairInput interface {
	pulumi.Input

	ToNameValuePairOutput() NameValuePairOutput
	ToNameValuePairOutputWithContext(context.Context) NameValuePairOutput
}

type NameValuePairArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (NameValuePairArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (i NameValuePairArgs) ToNameValuePairOutput() NameValuePairOutput {
	return i.ToNameValuePairOutputWithContext(context.Background())
}

func (i NameValuePairArgs) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairOutput)
}





type NameValuePairArrayInput interface {
	pulumi.Input

	ToNameValuePairArrayOutput() NameValuePairArrayOutput
	ToNameValuePairArrayOutputWithContext(context.Context) NameValuePairArrayOutput
}

type NameValuePairArray []NameValuePairInput

func (NameValuePairArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (i NameValuePairArray) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return i.ToNameValuePairArrayOutputWithContext(context.Background())
}

func (i NameValuePairArray) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairArrayOutput)
}

type NameValuePairOutput struct{ *pulumi.OutputState }

func (NameValuePairOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePair)(nil)).Elem()
}

func (o NameValuePairOutput) ToNameValuePairOutput() NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) ToNameValuePairOutputWithContext(ctx context.Context) NameValuePairOutput {
	return o
}

func (o NameValuePairOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePair) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePair)(nil)).Elem()
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutput() NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) ToNameValuePairArrayOutputWithContext(ctx context.Context) NameValuePairArrayOutput {
	return o
}

func (o NameValuePairArrayOutput) Index(i pulumi.IntInput) NameValuePairOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePair {
		return vs[0].([]NameValuePair)[vs[1].(int)]
	}).(NameValuePairOutput)
}

type NameValuePairResponse struct {
	Name  *string `pulumi:"name"`
	Value *string `pulumi:"value"`
}

type NameValuePairResponseOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutput() NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) ToNameValuePairResponseOutputWithContext(ctx context.Context) NameValuePairResponseOutput {
	return o
}

func (o NameValuePairResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o NameValuePairResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NameValuePairResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type NameValuePairResponseArrayOutput struct{ *pulumi.OutputState }

func (NameValuePairResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePairResponse)(nil)).Elem()
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutput() NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) ToNameValuePairResponseArrayOutputWithContext(ctx context.Context) NameValuePairResponseArrayOutput {
	return o
}

func (o NameValuePairResponseArrayOutput) Index(i pulumi.IntInput) NameValuePairResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NameValuePairResponse {
		return vs[0].([]NameValuePairResponse)[vs[1].(int)]
	}).(NameValuePairResponseOutput)
}

type PrivateLinkConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput
	ToPrivateLinkConnectionStateOutputWithContext(context.Context) PrivateLinkConnectionStateOutput
}

type PrivateLinkConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return i.ToPrivateLinkConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput)
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateArgs) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateOutput).ToPrivateLinkConnectionStatePtrOutputWithContext(ctx)
}









type PrivateLinkConnectionStatePtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput
	ToPrivateLinkConnectionStatePtrOutputWithContext(context.Context) PrivateLinkConnectionStatePtrOutput
}

type privateLinkConnectionStatePtrType PrivateLinkConnectionStateArgs

func PrivateLinkConnectionStatePtr(v *PrivateLinkConnectionStateArgs) PrivateLinkConnectionStatePtrInput {
	return (*privateLinkConnectionStatePtrType)(v)
}

func (*privateLinkConnectionStatePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return i.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionStatePtrType) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStatePtrOutput)
}

type PrivateLinkConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutput() PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStateOutputWithContext(ctx context.Context) PrivateLinkConnectionStateOutput {
	return o
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o.ToPrivateLinkConnectionStatePtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionStateOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionState) *PrivateLinkConnectionState {
		return &v
	}).(PrivateLinkConnectionStatePtrOutput)
}

func (o PrivateLinkConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStatePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStatePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionState)(nil)).Elem()
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutput() PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) ToPrivateLinkConnectionStatePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStatePtrOutput {
	return o
}

func (o PrivateLinkConnectionStatePtrOutput) Elem() PrivateLinkConnectionStateOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) PrivateLinkConnectionState {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionState
		return ret
	}).(PrivateLinkConnectionStateOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStatePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionState) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkConnectionStateResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateLinkConnectionStateResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return o
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Elem() PrivateLinkConnectionStateResponseOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) PrivateLinkConnectionStateResponse {
		if v != nil {
			return *v
		}
		var ret PrivateLinkConnectionStateResponse
		return ret
	}).(PrivateLinkConnectionStateResponseOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.ActionsRequired
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o PrivateLinkConnectionStateResponsePtrOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateLinkConnectionStateResponse) *string {
		if v == nil {
			return nil
		}
		return v.Status
	}).(pulumi.StringPtrOutput)
}

type PushSettings struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
}





type PushSettingsInput interface {
	pulumi.Input

	ToPushSettingsOutput() PushSettingsOutput
	ToPushSettingsOutputWithContext(context.Context) PushSettingsOutput
}

type PushSettingsArgs struct {
	DynamicTagsJson   pulumi.StringPtrInput `pulumi:"dynamicTagsJson"`
	IsPushEnabled     pulumi.BoolInput      `pulumi:"isPushEnabled"`
	Kind              pulumi.StringPtrInput `pulumi:"kind"`
	TagWhitelistJson  pulumi.StringPtrInput `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth pulumi.StringPtrInput `pulumi:"tagsRequiringAuth"`
}

func (PushSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettings)(nil)).Elem()
}

func (i PushSettingsArgs) ToPushSettingsOutput() PushSettingsOutput {
	return i.ToPushSettingsOutputWithContext(context.Background())
}

func (i PushSettingsArgs) ToPushSettingsOutputWithContext(ctx context.Context) PushSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsOutput)
}

func (i PushSettingsArgs) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return i.ToPushSettingsPtrOutputWithContext(context.Background())
}

func (i PushSettingsArgs) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsOutput).ToPushSettingsPtrOutputWithContext(ctx)
}









type PushSettingsPtrInput interface {
	pulumi.Input

	ToPushSettingsPtrOutput() PushSettingsPtrOutput
	ToPushSettingsPtrOutputWithContext(context.Context) PushSettingsPtrOutput
}

type pushSettingsPtrType PushSettingsArgs

func PushSettingsPtr(v *PushSettingsArgs) PushSettingsPtrInput {
	return (*pushSettingsPtrType)(v)
}

func (*pushSettingsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettings)(nil)).Elem()
}

func (i *pushSettingsPtrType) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return i.ToPushSettingsPtrOutputWithContext(context.Background())
}

func (i *pushSettingsPtrType) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsPtrOutput)
}

type PushSettingsOutput struct{ *pulumi.OutputState }

func (PushSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettings)(nil)).Elem()
}

func (o PushSettingsOutput) ToPushSettingsOutput() PushSettingsOutput {
	return o
}

func (o PushSettingsOutput) ToPushSettingsOutputWithContext(ctx context.Context) PushSettingsOutput {
	return o
}

func (o PushSettingsOutput) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return o.ToPushSettingsPtrOutputWithContext(context.Background())
}

func (o PushSettingsOutput) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PushSettings) *PushSettings {
		return &v
	}).(PushSettingsPtrOutput)
}

func (o PushSettingsOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v PushSettings) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o PushSettingsOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PushSettingsOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettings) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

type PushSettingsPtrOutput struct{ *pulumi.OutputState }

func (PushSettingsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettings)(nil)).Elem()
}

func (o PushSettingsPtrOutput) ToPushSettingsPtrOutput() PushSettingsPtrOutput {
	return o
}

func (o PushSettingsPtrOutput) ToPushSettingsPtrOutputWithContext(ctx context.Context) PushSettingsPtrOutput {
	return o
}

func (o PushSettingsPtrOutput) Elem() PushSettingsOutput {
	return o.ApplyT(func(v *PushSettings) PushSettings {
		if v != nil {
			return *v
		}
		var ret PushSettings
		return ret
	}).(PushSettingsOutput)
}

func (o PushSettingsPtrOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.DynamicTagsJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsPtrOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PushSettings) *bool {
		if v == nil {
			return nil
		}
		return &v.IsPushEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PushSettingsPtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsPtrOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.TagWhitelistJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsPtrOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettings) *string {
		if v == nil {
			return nil
		}
		return v.TagsRequiringAuth
	}).(pulumi.StringPtrOutput)
}

type PushSettingsResponse struct {
	DynamicTagsJson   *string `pulumi:"dynamicTagsJson"`
	Id                string  `pulumi:"id"`
	IsPushEnabled     bool    `pulumi:"isPushEnabled"`
	Kind              *string `pulumi:"kind"`
	Name              string  `pulumi:"name"`
	TagWhitelistJson  *string `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth *string `pulumi:"tagsRequiringAuth"`
	Type              string  `pulumi:"type"`
}

type PushSettingsResponseOutput struct{ *pulumi.OutputState }

func (PushSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettingsResponse)(nil)).Elem()
}

func (o PushSettingsResponseOutput) ToPushSettingsResponseOutput() PushSettingsResponseOutput {
	return o
}

func (o PushSettingsResponseOutput) ToPushSettingsResponseOutputWithContext(ctx context.Context) PushSettingsResponseOutput {
	return o
}

func (o PushSettingsResponseOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.DynamicTagsJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PushSettingsResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o PushSettingsResponseOutput) IsPushEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v PushSettingsResponse) bool { return v.IsPushEnabled }).(pulumi.BoolOutput)
}

func (o PushSettingsResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v PushSettingsResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o PushSettingsResponseOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.TagWhitelistJson }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PushSettingsResponse) *string { return v.TagsRequiringAuth }).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v PushSettingsResponse) string { return v.Type }).(pulumi.StringOutput)
}

type PushSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (PushSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettingsResponse)(nil)).Elem()
}

func (o PushSettingsResponsePtrOutput) ToPushSettingsResponsePtrOutput() PushSettingsResponsePtrOutput {
	return o
}

func (o PushSettingsResponsePtrOutput) ToPushSettingsResponsePtrOutputWithContext(ctx context.Context) PushSettingsResponsePtrOutput {
	return o
}

func (o PushSettingsResponsePtrOutput) Elem() PushSettingsResponseOutput {
	return o.ApplyT(func(v *PushSettingsResponse) PushSettingsResponse {
		if v != nil {
			return *v
		}
		var ret PushSettingsResponse
		return ret
	}).(PushSettingsResponseOutput)
}

func (o PushSettingsResponsePtrOutput) DynamicTagsJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.DynamicTagsJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) IsPushEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsPushEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) TagWhitelistJson() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TagWhitelistJson
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) TagsRequiringAuth() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TagsRequiringAuth
	}).(pulumi.StringPtrOutput)
}

func (o PushSettingsResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PushSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
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

type RampUpRule struct {
	ActionHostName            *string  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl *string  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   *int     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                *float64 `pulumi:"changeStep"`
	MaxReroutePercentage      *float64 `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      *float64 `pulumi:"minReroutePercentage"`
	Name                      *string  `pulumi:"name"`
	ReroutePercentage         *float64 `pulumi:"reroutePercentage"`
}





type RampUpRuleInput interface {
	pulumi.Input

	ToRampUpRuleOutput() RampUpRuleOutput
	ToRampUpRuleOutputWithContext(context.Context) RampUpRuleOutput
}

type RampUpRuleArgs struct {
	ActionHostName            pulumi.StringPtrInput  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl pulumi.StringPtrInput  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   pulumi.IntPtrInput     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                pulumi.Float64PtrInput `pulumi:"changeStep"`
	MaxReroutePercentage      pulumi.Float64PtrInput `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      pulumi.Float64PtrInput `pulumi:"minReroutePercentage"`
	Name                      pulumi.StringPtrInput  `pulumi:"name"`
	ReroutePercentage         pulumi.Float64PtrInput `pulumi:"reroutePercentage"`
}

func (RampUpRuleArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRule)(nil)).Elem()
}

func (i RampUpRuleArgs) ToRampUpRuleOutput() RampUpRuleOutput {
	return i.ToRampUpRuleOutputWithContext(context.Background())
}

func (i RampUpRuleArgs) ToRampUpRuleOutputWithContext(ctx context.Context) RampUpRuleOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleOutput)
}





type RampUpRuleArrayInput interface {
	pulumi.Input

	ToRampUpRuleArrayOutput() RampUpRuleArrayOutput
	ToRampUpRuleArrayOutputWithContext(context.Context) RampUpRuleArrayOutput
}

type RampUpRuleArray []RampUpRuleInput

func (RampUpRuleArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRule)(nil)).Elem()
}

func (i RampUpRuleArray) ToRampUpRuleArrayOutput() RampUpRuleArrayOutput {
	return i.ToRampUpRuleArrayOutputWithContext(context.Background())
}

func (i RampUpRuleArray) ToRampUpRuleArrayOutputWithContext(ctx context.Context) RampUpRuleArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleArrayOutput)
}

type RampUpRuleOutput struct{ *pulumi.OutputState }

func (RampUpRuleOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRule)(nil)).Elem()
}

func (o RampUpRuleOutput) ToRampUpRuleOutput() RampUpRuleOutput {
	return o
}

func (o RampUpRuleOutput) ToRampUpRuleOutputWithContext(ctx context.Context) RampUpRuleOutput {
	return o
}

func (o RampUpRuleOutput) ActionHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.ActionHostName }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ChangeDecisionCallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.ChangeDecisionCallbackUrl }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ChangeIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RampUpRule) *int { return v.ChangeIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o RampUpRuleOutput) ChangeStep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.ChangeStep }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) MaxReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.MaxReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) MinReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.MinReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRule) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleOutput) ReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRule) *float64 { return v.ReroutePercentage }).(pulumi.Float64PtrOutput)
}

type RampUpRuleArrayOutput struct{ *pulumi.OutputState }

func (RampUpRuleArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRule)(nil)).Elem()
}

func (o RampUpRuleArrayOutput) ToRampUpRuleArrayOutput() RampUpRuleArrayOutput {
	return o
}

func (o RampUpRuleArrayOutput) ToRampUpRuleArrayOutputWithContext(ctx context.Context) RampUpRuleArrayOutput {
	return o
}

func (o RampUpRuleArrayOutput) Index(i pulumi.IntInput) RampUpRuleOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RampUpRule {
		return vs[0].([]RampUpRule)[vs[1].(int)]
	}).(RampUpRuleOutput)
}

type RampUpRuleResponse struct {
	ActionHostName            *string  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl *string  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   *int     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                *float64 `pulumi:"changeStep"`
	MaxReroutePercentage      *float64 `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      *float64 `pulumi:"minReroutePercentage"`
	Name                      *string  `pulumi:"name"`
	ReroutePercentage         *float64 `pulumi:"reroutePercentage"`
}

type RampUpRuleResponseOutput struct{ *pulumi.OutputState }

func (RampUpRuleResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRuleResponse)(nil)).Elem()
}

func (o RampUpRuleResponseOutput) ToRampUpRuleResponseOutput() RampUpRuleResponseOutput {
	return o
}

func (o RampUpRuleResponseOutput) ToRampUpRuleResponseOutputWithContext(ctx context.Context) RampUpRuleResponseOutput {
	return o
}

func (o RampUpRuleResponseOutput) ActionHostName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.ActionHostName }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeDecisionCallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.ChangeDecisionCallbackUrl }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeIntervalInMinutes() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *int { return v.ChangeIntervalInMinutes }).(pulumi.IntPtrOutput)
}

func (o RampUpRuleResponseOutput) ChangeStep() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.ChangeStep }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) MaxReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.MaxReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) MinReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.MinReroutePercentage }).(pulumi.Float64PtrOutput)
}

func (o RampUpRuleResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o RampUpRuleResponseOutput) ReroutePercentage() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v RampUpRuleResponse) *float64 { return v.ReroutePercentage }).(pulumi.Float64PtrOutput)
}

type RampUpRuleResponseArrayOutput struct{ *pulumi.OutputState }

func (RampUpRuleResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRuleResponse)(nil)).Elem()
}

func (o RampUpRuleResponseArrayOutput) ToRampUpRuleResponseArrayOutput() RampUpRuleResponseArrayOutput {
	return o
}

func (o RampUpRuleResponseArrayOutput) ToRampUpRuleResponseArrayOutputWithContext(ctx context.Context) RampUpRuleResponseArrayOutput {
	return o
}

func (o RampUpRuleResponseArrayOutput) Index(i pulumi.IntInput) RampUpRuleResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) RampUpRuleResponse {
		return vs[0].([]RampUpRuleResponse)[vs[1].(int)]
	}).(RampUpRuleResponseOutput)
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

type RemotePrivateEndpointConnectionResponse struct {
	Id                                string                              `pulumi:"id"`
	IpAddresses                       []string                            `pulumi:"ipAddresses"`
	Kind                              *string                             `pulumi:"kind"`
	Name                              string                              `pulumi:"name"`
	PrivateEndpoint                   *ArmIdWrapperResponse               `pulumi:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionStateResponse `pulumi:"privateLinkServiceConnectionState"`
	ProvisioningState                 string                              `pulumi:"provisioningState"`
	Type                              string                              `pulumi:"type"`
}

type RemotePrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutput() RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) ToRemotePrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) []string { return v.IpAddresses }).(pulumi.StringArrayOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *ArmIdWrapperResponse { return v.PrivateEndpoint }).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) *PrivateLinkConnectionStateResponse {
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o RemotePrivateEndpointConnectionResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v RemotePrivateEndpointConnectionResponse) string { return v.Type }).(pulumi.StringOutput)
}

type RemotePrivateEndpointConnectionResponsePtrOutput struct{ *pulumi.OutputState }

func (RemotePrivateEndpointConnectionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) ToRemotePrivateEndpointConnectionResponsePtrOutput() RemotePrivateEndpointConnectionResponsePtrOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) ToRemotePrivateEndpointConnectionResponsePtrOutputWithContext(ctx context.Context) RemotePrivateEndpointConnectionResponsePtrOutput {
	return o
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Elem() RemotePrivateEndpointConnectionResponseOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) RemotePrivateEndpointConnectionResponse {
		if v != nil {
			return *v
		}
		var ret RemotePrivateEndpointConnectionResponse
		return ret
	}).(RemotePrivateEndpointConnectionResponseOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) IpAddresses() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) []string {
		if v == nil {
			return nil
		}
		return v.IpAddresses
	}).(pulumi.StringArrayOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Kind
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) PrivateEndpoint() ArmIdWrapperResponsePtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *ArmIdWrapperResponse {
		if v == nil {
			return nil
		}
		return v.PrivateEndpoint
	}).(ArmIdWrapperResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) PrivateLinkServiceConnectionState() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *PrivateLinkConnectionStateResponse {
		if v == nil {
			return nil
		}
		return v.PrivateLinkServiceConnectionState
	}).(PrivateLinkConnectionStateResponsePtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o RemotePrivateEndpointConnectionResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RemotePrivateEndpointConnectionResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
}

type RequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
}





type RequestsBasedTriggerInput interface {
	pulumi.Input

	ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput
	ToRequestsBasedTriggerOutputWithContext(context.Context) RequestsBasedTriggerOutput
}

type RequestsBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
}

func (RequestsBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTrigger)(nil)).Elem()
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput {
	return i.ToRequestsBasedTriggerOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerOutputWithContext(ctx context.Context) RequestsBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerOutput)
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return i.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerArgs) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerOutput).ToRequestsBasedTriggerPtrOutputWithContext(ctx)
}









type RequestsBasedTriggerPtrInput interface {
	pulumi.Input

	ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput
	ToRequestsBasedTriggerPtrOutputWithContext(context.Context) RequestsBasedTriggerPtrOutput
}

type requestsBasedTriggerPtrType RequestsBasedTriggerArgs

func RequestsBasedTriggerPtr(v *RequestsBasedTriggerArgs) RequestsBasedTriggerPtrInput {
	return (*requestsBasedTriggerPtrType)(v)
}

func (*requestsBasedTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTrigger)(nil)).Elem()
}

func (i *requestsBasedTriggerPtrType) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return i.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i *requestsBasedTriggerPtrType) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerPtrOutput)
}

type RequestsBasedTriggerOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTrigger)(nil)).Elem()
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerOutput() RequestsBasedTriggerOutput {
	return o
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerOutputWithContext(ctx context.Context) RequestsBasedTriggerOutput {
	return o
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return o.ToRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (o RequestsBasedTriggerOutput) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestsBasedTrigger) *RequestsBasedTrigger {
		return &v
	}).(RequestsBasedTriggerPtrOutput)
}

func (o RequestsBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RequestsBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestsBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerPtrOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTrigger)(nil)).Elem()
}

func (o RequestsBasedTriggerPtrOutput) ToRequestsBasedTriggerPtrOutput() RequestsBasedTriggerPtrOutput {
	return o
}

func (o RequestsBasedTriggerPtrOutput) ToRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) RequestsBasedTriggerPtrOutput {
	return o
}

func (o RequestsBasedTriggerPtrOutput) Elem() RequestsBasedTriggerOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) RequestsBasedTrigger {
		if v != nil {
			return *v
		}
		var ret RequestsBasedTrigger
		return ret
	}).(RequestsBasedTriggerOutput)
}

func (o RequestsBasedTriggerPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerPtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
}

type RequestsBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTriggerResponse)(nil)).Elem()
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponseOutput() RequestsBasedTriggerResponseOutput {
	return o
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) RequestsBasedTriggerResponseOutput {
	return o
}

func (o RequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v RequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v RequestsBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type RequestsBasedTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (RequestsBasedTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTriggerResponse)(nil)).Elem()
}

func (o RequestsBasedTriggerResponsePtrOutput) ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o RequestsBasedTriggerResponsePtrOutput) ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) RequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o RequestsBasedTriggerResponsePtrOutput) Elem() RequestsBasedTriggerResponseOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) RequestsBasedTriggerResponse {
		if v != nil {
			return *v
		}
		var ret RequestsBasedTriggerResponse
		return ret
	}).(RequestsBasedTriggerResponseOutput)
}

func (o RequestsBasedTriggerResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o RequestsBasedTriggerResponsePtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *RequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

type ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse struct {
	Error      *ErrorEntityResponse                     `pulumi:"error"`
	Id         *string                                  `pulumi:"id"`
	Identity   *ManagedServiceIdentityResponse          `pulumi:"identity"`
	Location   *string                                  `pulumi:"location"`
	Name       *string                                  `pulumi:"name"`
	Plan       *ArmPlanResponse                         `pulumi:"plan"`
	Properties *RemotePrivateEndpointConnectionResponse `pulumi:"properties"`
	Sku        *SkuDescriptionResponse                  `pulumi:"sku"`
	Status     *string                                  `pulumi:"status"`
	Tags       map[string]string                        `pulumi:"tags"`
	Type       *string                                  `pulumi:"type"`
	Zones      []string                                 `pulumi:"zones"`
}

type ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutputWithContext(ctx context.Context) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Error() ErrorEntityResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *ErrorEntityResponse {
		return v.Error
	}).(ErrorEntityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Identity() ManagedServiceIdentityResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *ManagedServiceIdentityResponse {
		return v.Identity
	}).(ManagedServiceIdentityResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Location() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Location }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Plan() ArmPlanResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *ArmPlanResponse { return v.Plan }).(ArmPlanResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Properties() RemotePrivateEndpointConnectionResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *RemotePrivateEndpointConnectionResponse {
		return v.Properties
	}).(RemotePrivateEndpointConnectionResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Sku() SkuDescriptionResponsePtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *SkuDescriptionResponse {
		return v.Sku
	}).(SkuDescriptionResponsePtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Tags() pulumi.StringMapOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) map[string]string {
		return v.Tags
	}).(pulumi.StringMapOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) *string { return v.Type }).(pulumi.StringPtrOutput)
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput) Zones() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse) []string { return v.Zones }).(pulumi.StringArrayOutput)
}

type ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput struct{ *pulumi.OutputState }

func (ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse)(nil)).Elem()
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput() ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) ToResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutputWithContext(ctx context.Context) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput {
	return o
}

func (o ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput) Index(i pulumi.IntInput) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse {
		return vs[0].([]ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponse)[vs[1].(int)]
	}).(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput)
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

type SiteConfig struct {
	AcrUseManagedIdentityCreds             *bool                            `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID               *string                          `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                               *bool                            `pulumi:"alwaysOn"`
	ApiDefinition                          *ApiDefinitionInfo               `pulumi:"apiDefinition"`
	ApiManagementConfig                    *ApiManagementConfig             `pulumi:"apiManagementConfig"`
	AppCommandLine                         *string                          `pulumi:"appCommandLine"`
	AppSettings                            []NameValuePair                  `pulumi:"appSettings"`
	AutoHealEnabled                        *bool                            `pulumi:"autoHealEnabled"`
	AutoHealRules                          *AutoHealRules                   `pulumi:"autoHealRules"`
	AutoSwapSlotName                       *string                          `pulumi:"autoSwapSlotName"`
	AzureStorageAccounts                   map[string]AzureStorageInfoValue `pulumi:"azureStorageAccounts"`
	ConnectionStrings                      []ConnStringInfo                 `pulumi:"connectionStrings"`
	Cors                                   *CorsSettings                    `pulumi:"cors"`
	DefaultDocuments                       []string                         `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled            *bool                            `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                           *string                          `pulumi:"documentRoot"`
	Experiments                            *Experiments                     `pulumi:"experiments"`
	FtpsState                              *string                          `pulumi:"ftpsState"`
	FunctionAppScaleLimit                  *int                             `pulumi:"functionAppScaleLimit"`
	FunctionsRuntimeScaleMonitoringEnabled *bool                            `pulumi:"functionsRuntimeScaleMonitoringEnabled"`
	HandlerMappings                        []HandlerMapping                 `pulumi:"handlerMappings"`
	HealthCheckPath                        *string                          `pulumi:"healthCheckPath"`
	Http20Enabled                          *bool                            `pulumi:"http20Enabled"`
	HttpLoggingEnabled                     *bool                            `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions                 []IpSecurityRestriction          `pulumi:"ipSecurityRestrictions"`
	JavaContainer                          *string                          `pulumi:"javaContainer"`
	JavaContainerVersion                   *string                          `pulumi:"javaContainerVersion"`
	JavaVersion                            *string                          `pulumi:"javaVersion"`
	KeyVaultReferenceIdentity              *string                          `pulumi:"keyVaultReferenceIdentity"`
	Limits                                 *SiteLimits                      `pulumi:"limits"`
	LinuxFxVersion                         *string                          `pulumi:"linuxFxVersion"`
	LoadBalancing                          *SiteLoadBalancing               `pulumi:"loadBalancing"`
	LocalMySqlEnabled                      *bool                            `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit                 *int                             `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode                    *ManagedPipelineMode             `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId               *int                             `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                          *string                          `pulumi:"minTlsVersion"`
	MinimumElasticInstanceCount            *int                             `pulumi:"minimumElasticInstanceCount"`
	NetFrameworkVersion                    *string                          `pulumi:"netFrameworkVersion"`
	NodeVersion                            *string                          `pulumi:"nodeVersion"`
	NumberOfWorkers                        *int                             `pulumi:"numberOfWorkers"`
	PhpVersion                             *string                          `pulumi:"phpVersion"`
	PowerShellVersion                      *string                          `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount                 *int                             `pulumi:"preWarmedInstanceCount"`
	PublicNetworkAccess                    *string                          `pulumi:"publicNetworkAccess"`
	PublishingUsername                     *string                          `pulumi:"publishingUsername"`
	Push                                   *PushSettings                    `pulumi:"push"`
	PythonVersion                          *string                          `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled                 *bool                            `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion                 *string                          `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled                  *bool                            `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime           *string                          `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions              []IpSecurityRestriction          `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain       *bool                            `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmMinTlsVersion                       *string                          `pulumi:"scmMinTlsVersion"`
	ScmType                                *string                          `pulumi:"scmType"`
	TracingOptions                         *string                          `pulumi:"tracingOptions"`
	Use32BitWorkerProcess                  *bool                            `pulumi:"use32BitWorkerProcess"`
	VirtualApplications                    []VirtualApplication             `pulumi:"virtualApplications"`
	VnetName                               *string                          `pulumi:"vnetName"`
	VnetPrivatePortsCount                  *int                             `pulumi:"vnetPrivatePortsCount"`
	VnetRouteAllEnabled                    *bool                            `pulumi:"vnetRouteAllEnabled"`
	WebSocketsEnabled                      *bool                            `pulumi:"webSocketsEnabled"`
	WebsiteTimeZone                        *string                          `pulumi:"websiteTimeZone"`
	WindowsFxVersion                       *string                          `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId              *int                             `pulumi:"xManagedServiceIdentityId"`
}


func (val *SiteConfig) Defaults() *SiteConfig {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Http20Enabled) {
		http20Enabled_ := true
		tmp.Http20Enabled = &http20Enabled_
	}
	if isZero(tmp.LocalMySqlEnabled) {
		localMySqlEnabled_ := false
		tmp.LocalMySqlEnabled = &localMySqlEnabled_
	}
	if isZero(tmp.NetFrameworkVersion) {
		netFrameworkVersion_ := "v4.6"
		tmp.NetFrameworkVersion = &netFrameworkVersion_
	}
	return &tmp
}





type SiteConfigInput interface {
	pulumi.Input

	ToSiteConfigOutput() SiteConfigOutput
	ToSiteConfigOutputWithContext(context.Context) SiteConfigOutput
}

type SiteConfigArgs struct {
	AcrUseManagedIdentityCreds             pulumi.BoolPtrInput             `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID               pulumi.StringPtrInput           `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                               pulumi.BoolPtrInput             `pulumi:"alwaysOn"`
	ApiDefinition                          ApiDefinitionInfoPtrInput       `pulumi:"apiDefinition"`
	ApiManagementConfig                    ApiManagementConfigPtrInput     `pulumi:"apiManagementConfig"`
	AppCommandLine                         pulumi.StringPtrInput           `pulumi:"appCommandLine"`
	AppSettings                            NameValuePairArrayInput         `pulumi:"appSettings"`
	AutoHealEnabled                        pulumi.BoolPtrInput             `pulumi:"autoHealEnabled"`
	AutoHealRules                          AutoHealRulesPtrInput           `pulumi:"autoHealRules"`
	AutoSwapSlotName                       pulumi.StringPtrInput           `pulumi:"autoSwapSlotName"`
	AzureStorageAccounts                   AzureStorageInfoValueMapInput   `pulumi:"azureStorageAccounts"`
	ConnectionStrings                      ConnStringInfoArrayInput        `pulumi:"connectionStrings"`
	Cors                                   CorsSettingsPtrInput            `pulumi:"cors"`
	DefaultDocuments                       pulumi.StringArrayInput         `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled            pulumi.BoolPtrInput             `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                           pulumi.StringPtrInput           `pulumi:"documentRoot"`
	Experiments                            ExperimentsPtrInput             `pulumi:"experiments"`
	FtpsState                              pulumi.StringPtrInput           `pulumi:"ftpsState"`
	FunctionAppScaleLimit                  pulumi.IntPtrInput              `pulumi:"functionAppScaleLimit"`
	FunctionsRuntimeScaleMonitoringEnabled pulumi.BoolPtrInput             `pulumi:"functionsRuntimeScaleMonitoringEnabled"`
	HandlerMappings                        HandlerMappingArrayInput        `pulumi:"handlerMappings"`
	HealthCheckPath                        pulumi.StringPtrInput           `pulumi:"healthCheckPath"`
	Http20Enabled                          pulumi.BoolPtrInput             `pulumi:"http20Enabled"`
	HttpLoggingEnabled                     pulumi.BoolPtrInput             `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions                 IpSecurityRestrictionArrayInput `pulumi:"ipSecurityRestrictions"`
	JavaContainer                          pulumi.StringPtrInput           `pulumi:"javaContainer"`
	JavaContainerVersion                   pulumi.StringPtrInput           `pulumi:"javaContainerVersion"`
	JavaVersion                            pulumi.StringPtrInput           `pulumi:"javaVersion"`
	KeyVaultReferenceIdentity              pulumi.StringPtrInput           `pulumi:"keyVaultReferenceIdentity"`
	Limits                                 SiteLimitsPtrInput              `pulumi:"limits"`
	LinuxFxVersion                         pulumi.StringPtrInput           `pulumi:"linuxFxVersion"`
	LoadBalancing                          SiteLoadBalancingPtrInput       `pulumi:"loadBalancing"`
	LocalMySqlEnabled                      pulumi.BoolPtrInput             `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit                 pulumi.IntPtrInput              `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode                    ManagedPipelineModePtrInput     `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId               pulumi.IntPtrInput              `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                          pulumi.StringPtrInput           `pulumi:"minTlsVersion"`
	MinimumElasticInstanceCount            pulumi.IntPtrInput              `pulumi:"minimumElasticInstanceCount"`
	NetFrameworkVersion                    pulumi.StringPtrInput           `pulumi:"netFrameworkVersion"`
	NodeVersion                            pulumi.StringPtrInput           `pulumi:"nodeVersion"`
	NumberOfWorkers                        pulumi.IntPtrInput              `pulumi:"numberOfWorkers"`
	PhpVersion                             pulumi.StringPtrInput           `pulumi:"phpVersion"`
	PowerShellVersion                      pulumi.StringPtrInput           `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount                 pulumi.IntPtrInput              `pulumi:"preWarmedInstanceCount"`
	PublicNetworkAccess                    pulumi.StringPtrInput           `pulumi:"publicNetworkAccess"`
	PublishingUsername                     pulumi.StringPtrInput           `pulumi:"publishingUsername"`
	Push                                   PushSettingsPtrInput            `pulumi:"push"`
	PythonVersion                          pulumi.StringPtrInput           `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled                 pulumi.BoolPtrInput             `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion                 pulumi.StringPtrInput           `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled                  pulumi.BoolPtrInput             `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime           pulumi.StringPtrInput           `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions              IpSecurityRestrictionArrayInput `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain       pulumi.BoolPtrInput             `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmMinTlsVersion                       pulumi.StringPtrInput           `pulumi:"scmMinTlsVersion"`
	ScmType                                pulumi.StringPtrInput           `pulumi:"scmType"`
	TracingOptions                         pulumi.StringPtrInput           `pulumi:"tracingOptions"`
	Use32BitWorkerProcess                  pulumi.BoolPtrInput             `pulumi:"use32BitWorkerProcess"`
	VirtualApplications                    VirtualApplicationArrayInput    `pulumi:"virtualApplications"`
	VnetName                               pulumi.StringPtrInput           `pulumi:"vnetName"`
	VnetPrivatePortsCount                  pulumi.IntPtrInput              `pulumi:"vnetPrivatePortsCount"`
	VnetRouteAllEnabled                    pulumi.BoolPtrInput             `pulumi:"vnetRouteAllEnabled"`
	WebSocketsEnabled                      pulumi.BoolPtrInput             `pulumi:"webSocketsEnabled"`
	WebsiteTimeZone                        pulumi.StringPtrInput           `pulumi:"websiteTimeZone"`
	WindowsFxVersion                       pulumi.StringPtrInput           `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId              pulumi.IntPtrInput              `pulumi:"xManagedServiceIdentityId"`
}


func (val *SiteConfigArgs) Defaults() *SiteConfigArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Http20Enabled) {
		tmp.Http20Enabled = pulumi.BoolPtr(true)
	}
	if isZero(tmp.LocalMySqlEnabled) {
		tmp.LocalMySqlEnabled = pulumi.BoolPtr(false)
	}
	if isZero(tmp.NetFrameworkVersion) {
		tmp.NetFrameworkVersion = pulumi.StringPtr("v4.6")
	}
	return &tmp
}
func (SiteConfigArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfig)(nil)).Elem()
}

func (i SiteConfigArgs) ToSiteConfigOutput() SiteConfigOutput {
	return i.ToSiteConfigOutputWithContext(context.Background())
}

func (i SiteConfigArgs) ToSiteConfigOutputWithContext(ctx context.Context) SiteConfigOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigOutput)
}

func (i SiteConfigArgs) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return i.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (i SiteConfigArgs) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigOutput).ToSiteConfigPtrOutputWithContext(ctx)
}









type SiteConfigPtrInput interface {
	pulumi.Input

	ToSiteConfigPtrOutput() SiteConfigPtrOutput
	ToSiteConfigPtrOutputWithContext(context.Context) SiteConfigPtrOutput
}

type siteConfigPtrType SiteConfigArgs

func SiteConfigPtr(v *SiteConfigArgs) SiteConfigPtrInput {
	return (*siteConfigPtrType)(v)
}

func (*siteConfigPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfig)(nil)).Elem()
}

func (i *siteConfigPtrType) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return i.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (i *siteConfigPtrType) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigPtrOutput)
}

type SiteConfigOutput struct{ *pulumi.OutputState }

func (SiteConfigOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfig)(nil)).Elem()
}

func (o SiteConfigOutput) ToSiteConfigOutput() SiteConfigOutput {
	return o
}

func (o SiteConfigOutput) ToSiteConfigOutputWithContext(ctx context.Context) SiteConfigOutput {
	return o
}

func (o SiteConfigOutput) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return o.ToSiteConfigPtrOutputWithContext(context.Background())
}

func (o SiteConfigOutput) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteConfig) *SiteConfig {
		return &v
	}).(SiteConfigPtrOutput)
}

func (o SiteConfigOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AcrUseManagedIdentityCreds }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AcrUserManagedIdentityID }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AlwaysOn }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) ApiDefinition() ApiDefinitionInfoPtrOutput {
	return o.ApplyT(func(v SiteConfig) *ApiDefinitionInfo { return v.ApiDefinition }).(ApiDefinitionInfoPtrOutput)
}

func (o SiteConfigOutput) ApiManagementConfig() ApiManagementConfigPtrOutput {
	return o.ApplyT(func(v SiteConfig) *ApiManagementConfig { return v.ApiManagementConfig }).(ApiManagementConfigPtrOutput)
}

func (o SiteConfigOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AppCommandLine }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AppSettings() NameValuePairArrayOutput {
	return o.ApplyT(func(v SiteConfig) []NameValuePair { return v.AppSettings }).(NameValuePairArrayOutput)
}

func (o SiteConfigOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.AutoHealEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) AutoHealRules() AutoHealRulesPtrOutput {
	return o.ApplyT(func(v SiteConfig) *AutoHealRules { return v.AutoHealRules }).(AutoHealRulesPtrOutput)
}

func (o SiteConfigOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.AutoSwapSlotName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) AzureStorageAccounts() AzureStorageInfoValueMapOutput {
	return o.ApplyT(func(v SiteConfig) map[string]AzureStorageInfoValue { return v.AzureStorageAccounts }).(AzureStorageInfoValueMapOutput)
}

func (o SiteConfigOutput) ConnectionStrings() ConnStringInfoArrayOutput {
	return o.ApplyT(func(v SiteConfig) []ConnStringInfo { return v.ConnectionStrings }).(ConnStringInfoArrayOutput)
}

func (o SiteConfigOutput) Cors() CorsSettingsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *CorsSettings { return v.Cors }).(CorsSettingsPtrOutput)
}

func (o SiteConfigOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteConfig) []string { return v.DefaultDocuments }).(pulumi.StringArrayOutput)
}

func (o SiteConfigOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.DetailedErrorLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Experiments() ExperimentsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *Experiments { return v.Experiments }).(ExperimentsPtrOutput)
}

func (o SiteConfigOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.FtpsState }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.FunctionAppScaleLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.FunctionsRuntimeScaleMonitoringEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) HandlerMappings() HandlerMappingArrayOutput {
	return o.ApplyT(func(v SiteConfig) []HandlerMapping { return v.HandlerMappings }).(HandlerMappingArrayOutput)
}

func (o SiteConfigOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.Http20Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.HttpLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) IpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v SiteConfig) []IpSecurityRestriction { return v.IpSecurityRestrictions }).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaContainer }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaContainerVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.JavaVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Limits() SiteLimitsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *SiteLimits { return v.Limits }).(SiteLimitsPtrOutput)
}

func (o SiteConfigOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.LinuxFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) LoadBalancing() SiteLoadBalancingPtrOutput {
	return o.ApplyT(func(v SiteConfig) *SiteLoadBalancing { return v.LoadBalancing }).(SiteLoadBalancingPtrOutput)
}

func (o SiteConfigOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.LocalMySqlEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.LogsDirectorySizeLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) ManagedPipelineMode() ManagedPipelineModePtrOutput {
	return o.ApplyT(func(v SiteConfig) *ManagedPipelineMode { return v.ManagedPipelineMode }).(ManagedPipelineModePtrOutput)
}

func (o SiteConfigOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.ManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.MinimumElasticInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.NetFrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.NodeVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PhpVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PowerShellVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.PreWarmedInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PublishingUsername }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Push() PushSettingsPtrOutput {
	return o.ApplyT(func(v SiteConfig) *PushSettings { return v.Push }).(PushSettingsPtrOutput)
}

func (o SiteConfigOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.PythonVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.RemoteDebuggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.RemoteDebuggingVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.RequestTracingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.RequestTracingExpirationTime }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v SiteConfig) []IpSecurityRestriction { return v.ScmIpSecurityRestrictions }).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.ScmIpSecurityRestrictionsUseMain }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.ScmMinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.ScmType }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.TracingOptions }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.Use32BitWorkerProcess }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) VirtualApplications() VirtualApplicationArrayOutput {
	return o.ApplyT(func(v SiteConfig) []VirtualApplication { return v.VirtualApplications }).(VirtualApplicationArrayOutput)
}

func (o SiteConfigOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.VnetPrivatePortsCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.VnetRouteAllEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.WebsiteTimeZone }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfig) *string { return v.WindowsFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfig) *int { return v.XManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

type SiteConfigPtrOutput struct{ *pulumi.OutputState }

func (SiteConfigPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfig)(nil)).Elem()
}

func (o SiteConfigPtrOutput) ToSiteConfigPtrOutput() SiteConfigPtrOutput {
	return o
}

func (o SiteConfigPtrOutput) ToSiteConfigPtrOutputWithContext(ctx context.Context) SiteConfigPtrOutput {
	return o
}

func (o SiteConfigPtrOutput) Elem() SiteConfigOutput {
	return o.ApplyT(func(v *SiteConfig) SiteConfig {
		if v != nil {
			return *v
		}
		var ret SiteConfig
		return ret
	}).(SiteConfigOutput)
}

func (o SiteConfigPtrOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AcrUseManagedIdentityCreds
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AcrUserManagedIdentityID
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AlwaysOn
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) ApiDefinition() ApiDefinitionInfoPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ApiDefinitionInfo {
		if v == nil {
			return nil
		}
		return v.ApiDefinition
	}).(ApiDefinitionInfoPtrOutput)
}

func (o SiteConfigPtrOutput) ApiManagementConfig() ApiManagementConfigPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ApiManagementConfig {
		if v == nil {
			return nil
		}
		return v.ApiManagementConfig
	}).(ApiManagementConfigPtrOutput)
}

func (o SiteConfigPtrOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AppCommandLine
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AppSettings() NameValuePairArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []NameValuePair {
		if v == nil {
			return nil
		}
		return v.AppSettings
	}).(NameValuePairArrayOutput)
}

func (o SiteConfigPtrOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.AutoHealEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) AutoHealRules() AutoHealRulesPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *AutoHealRules {
		if v == nil {
			return nil
		}
		return v.AutoHealRules
	}).(AutoHealRulesPtrOutput)
}

func (o SiteConfigPtrOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.AutoSwapSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) AzureStorageAccounts() AzureStorageInfoValueMapOutput {
	return o.ApplyT(func(v *SiteConfig) map[string]AzureStorageInfoValue {
		if v == nil {
			return nil
		}
		return v.AzureStorageAccounts
	}).(AzureStorageInfoValueMapOutput)
}

func (o SiteConfigPtrOutput) ConnectionStrings() ConnStringInfoArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []ConnStringInfo {
		if v == nil {
			return nil
		}
		return v.ConnectionStrings
	}).(ConnStringInfoArrayOutput)
}

func (o SiteConfigPtrOutput) Cors() CorsSettingsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *CorsSettings {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsSettingsPtrOutput)
}

func (o SiteConfigPtrOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []string {
		if v == nil {
			return nil
		}
		return v.DefaultDocuments
	}).(pulumi.StringArrayOutput)
}

func (o SiteConfigPtrOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.DetailedErrorLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.DocumentRoot
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Experiments() ExperimentsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *Experiments {
		if v == nil {
			return nil
		}
		return v.Experiments
	}).(ExperimentsPtrOutput)
}

func (o SiteConfigPtrOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.FtpsState
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.FunctionAppScaleLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.FunctionsRuntimeScaleMonitoringEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) HandlerMappings() HandlerMappingArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []HandlerMapping {
		if v == nil {
			return nil
		}
		return v.HandlerMappings
	}).(HandlerMappingArrayOutput)
}

func (o SiteConfigPtrOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckPath
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Http20Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.HttpLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) IpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []IpSecurityRestriction {
		if v == nil {
			return nil
		}
		return v.IpSecurityRestrictions
	}).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigPtrOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainer
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainerVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.JavaVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Limits() SiteLimitsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *SiteLimits {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(SiteLimitsPtrOutput)
}

func (o SiteConfigPtrOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.LinuxFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) LoadBalancing() SiteLoadBalancingPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *SiteLoadBalancing {
		if v == nil {
			return nil
		}
		return v.LoadBalancing
	}).(SiteLoadBalancingPtrOutput)
}

func (o SiteConfigPtrOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.LocalMySqlEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.LogsDirectorySizeLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) ManagedPipelineMode() ManagedPipelineModePtrOutput {
	return o.ApplyT(func(v *SiteConfig) *ManagedPipelineMode {
		if v == nil {
			return nil
		}
		return v.ManagedPipelineMode
	}).(ManagedPipelineModePtrOutput)
}

func (o SiteConfigPtrOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.ManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.MinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.MinimumElasticInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.NetFrameworkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.NodeVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.NumberOfWorkers
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PhpVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PowerShellVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.PreWarmedInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PublishingUsername
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Push() PushSettingsPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *PushSettings {
		if v == nil {
			return nil
		}
		return v.Push
	}).(PushSettingsPtrOutput)
}

func (o SiteConfigPtrOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.PythonVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.RequestTracingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.RequestTracingExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []IpSecurityRestriction {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictions
	}).(IpSecurityRestrictionArrayOutput)
}

func (o SiteConfigPtrOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictionsUseMain
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.ScmMinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.ScmType
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.TracingOptions
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.Use32BitWorkerProcess
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) VirtualApplications() VirtualApplicationArrayOutput {
	return o.ApplyT(func(v *SiteConfig) []VirtualApplication {
		if v == nil {
			return nil
		}
		return v.VirtualApplications
	}).(VirtualApplicationArrayOutput)
}

func (o SiteConfigPtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.VnetPrivatePortsCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigPtrOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.VnetRouteAllEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigPtrOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteTimeZone
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *string {
		if v == nil {
			return nil
		}
		return v.WindowsFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigPtrOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *int {
		if v == nil {
			return nil
		}
		return v.XManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

type SiteConfigResponse struct {
	AcrUseManagedIdentityCreds             *bool                                    `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID               *string                                  `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                               *bool                                    `pulumi:"alwaysOn"`
	ApiDefinition                          *ApiDefinitionInfoResponse               `pulumi:"apiDefinition"`
	ApiManagementConfig                    *ApiManagementConfigResponse             `pulumi:"apiManagementConfig"`
	AppCommandLine                         *string                                  `pulumi:"appCommandLine"`
	AppSettings                            []NameValuePairResponse                  `pulumi:"appSettings"`
	AutoHealEnabled                        *bool                                    `pulumi:"autoHealEnabled"`
	AutoHealRules                          *AutoHealRulesResponse                   `pulumi:"autoHealRules"`
	AutoSwapSlotName                       *string                                  `pulumi:"autoSwapSlotName"`
	AzureStorageAccounts                   map[string]AzureStorageInfoValueResponse `pulumi:"azureStorageAccounts"`
	ConnectionStrings                      []ConnStringInfoResponse                 `pulumi:"connectionStrings"`
	Cors                                   *CorsSettingsResponse                    `pulumi:"cors"`
	DefaultDocuments                       []string                                 `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled            *bool                                    `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                           *string                                  `pulumi:"documentRoot"`
	Experiments                            *ExperimentsResponse                     `pulumi:"experiments"`
	FtpsState                              *string                                  `pulumi:"ftpsState"`
	FunctionAppScaleLimit                  *int                                     `pulumi:"functionAppScaleLimit"`
	FunctionsRuntimeScaleMonitoringEnabled *bool                                    `pulumi:"functionsRuntimeScaleMonitoringEnabled"`
	HandlerMappings                        []HandlerMappingResponse                 `pulumi:"handlerMappings"`
	HealthCheckPath                        *string                                  `pulumi:"healthCheckPath"`
	Http20Enabled                          *bool                                    `pulumi:"http20Enabled"`
	HttpLoggingEnabled                     *bool                                    `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions                 []IpSecurityRestrictionResponse          `pulumi:"ipSecurityRestrictions"`
	JavaContainer                          *string                                  `pulumi:"javaContainer"`
	JavaContainerVersion                   *string                                  `pulumi:"javaContainerVersion"`
	JavaVersion                            *string                                  `pulumi:"javaVersion"`
	KeyVaultReferenceIdentity              *string                                  `pulumi:"keyVaultReferenceIdentity"`
	Limits                                 *SiteLimitsResponse                      `pulumi:"limits"`
	LinuxFxVersion                         *string                                  `pulumi:"linuxFxVersion"`
	LoadBalancing                          *string                                  `pulumi:"loadBalancing"`
	LocalMySqlEnabled                      *bool                                    `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit                 *int                                     `pulumi:"logsDirectorySizeLimit"`
	MachineKey                             SiteMachineKeyResponse                   `pulumi:"machineKey"`
	ManagedPipelineMode                    *string                                  `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId               *int                                     `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                          *string                                  `pulumi:"minTlsVersion"`
	MinimumElasticInstanceCount            *int                                     `pulumi:"minimumElasticInstanceCount"`
	NetFrameworkVersion                    *string                                  `pulumi:"netFrameworkVersion"`
	NodeVersion                            *string                                  `pulumi:"nodeVersion"`
	NumberOfWorkers                        *int                                     `pulumi:"numberOfWorkers"`
	PhpVersion                             *string                                  `pulumi:"phpVersion"`
	PowerShellVersion                      *string                                  `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount                 *int                                     `pulumi:"preWarmedInstanceCount"`
	PublicNetworkAccess                    *string                                  `pulumi:"publicNetworkAccess"`
	PublishingUsername                     *string                                  `pulumi:"publishingUsername"`
	Push                                   *PushSettingsResponse                    `pulumi:"push"`
	PythonVersion                          *string                                  `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled                 *bool                                    `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion                 *string                                  `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled                  *bool                                    `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime           *string                                  `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions              []IpSecurityRestrictionResponse          `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain       *bool                                    `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmMinTlsVersion                       *string                                  `pulumi:"scmMinTlsVersion"`
	ScmType                                *string                                  `pulumi:"scmType"`
	TracingOptions                         *string                                  `pulumi:"tracingOptions"`
	Use32BitWorkerProcess                  *bool                                    `pulumi:"use32BitWorkerProcess"`
	VirtualApplications                    []VirtualApplicationResponse             `pulumi:"virtualApplications"`
	VnetName                               *string                                  `pulumi:"vnetName"`
	VnetPrivatePortsCount                  *int                                     `pulumi:"vnetPrivatePortsCount"`
	VnetRouteAllEnabled                    *bool                                    `pulumi:"vnetRouteAllEnabled"`
	WebSocketsEnabled                      *bool                                    `pulumi:"webSocketsEnabled"`
	WebsiteTimeZone                        *string                                  `pulumi:"websiteTimeZone"`
	WindowsFxVersion                       *string                                  `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId              *int                                     `pulumi:"xManagedServiceIdentityId"`
}


func (val *SiteConfigResponse) Defaults() *SiteConfigResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Http20Enabled) {
		http20Enabled_ := true
		tmp.Http20Enabled = &http20Enabled_
	}
	if isZero(tmp.LocalMySqlEnabled) {
		localMySqlEnabled_ := false
		tmp.LocalMySqlEnabled = &localMySqlEnabled_
	}
	if isZero(tmp.NetFrameworkVersion) {
		netFrameworkVersion_ := "v4.6"
		tmp.NetFrameworkVersion = &netFrameworkVersion_
	}
	return &tmp
}

type SiteConfigResponseOutput struct{ *pulumi.OutputState }

func (SiteConfigResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfigResponse)(nil)).Elem()
}

func (o SiteConfigResponseOutput) ToSiteConfigResponseOutput() SiteConfigResponseOutput {
	return o
}

func (o SiteConfigResponseOutput) ToSiteConfigResponseOutputWithContext(ctx context.Context) SiteConfigResponseOutput {
	return o
}

func (o SiteConfigResponseOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AcrUseManagedIdentityCreds }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AcrUserManagedIdentityID }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AlwaysOn }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) ApiDefinition() ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ApiDefinitionInfoResponse { return v.ApiDefinition }).(ApiDefinitionInfoResponsePtrOutput)
}

func (o SiteConfigResponseOutput) ApiManagementConfig() ApiManagementConfigResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ApiManagementConfigResponse { return v.ApiManagementConfig }).(ApiManagementConfigResponsePtrOutput)
}

func (o SiteConfigResponseOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AppCommandLine }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AppSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []NameValuePairResponse { return v.AppSettings }).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponseOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.AutoHealEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) AutoHealRules() AutoHealRulesResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *AutoHealRulesResponse { return v.AutoHealRules }).(AutoHealRulesResponsePtrOutput)
}

func (o SiteConfigResponseOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.AutoSwapSlotName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) AzureStorageAccounts() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v SiteConfigResponse) map[string]AzureStorageInfoValueResponse { return v.AzureStorageAccounts }).(AzureStorageInfoValueResponseMapOutput)
}

func (o SiteConfigResponseOutput) ConnectionStrings() ConnStringInfoResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []ConnStringInfoResponse { return v.ConnectionStrings }).(ConnStringInfoResponseArrayOutput)
}

func (o SiteConfigResponseOutput) Cors() CorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *CorsSettingsResponse { return v.Cors }).(CorsSettingsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []string { return v.DefaultDocuments }).(pulumi.StringArrayOutput)
}

func (o SiteConfigResponseOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.DetailedErrorLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.DocumentRoot }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Experiments() ExperimentsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *ExperimentsResponse { return v.Experiments }).(ExperimentsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.FtpsState }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.FunctionAppScaleLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.FunctionsRuntimeScaleMonitoringEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) HandlerMappings() HandlerMappingResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []HandlerMappingResponse { return v.HandlerMappings }).(HandlerMappingResponseArrayOutput)
}

func (o SiteConfigResponseOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.HealthCheckPath }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.Http20Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.HttpLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) IpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []IpSecurityRestrictionResponse { return v.IpSecurityRestrictions }).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponseOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaContainer }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaContainerVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.JavaVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.KeyVaultReferenceIdentity }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Limits() SiteLimitsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *SiteLimitsResponse { return v.Limits }).(SiteLimitsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.LinuxFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) LoadBalancing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.LoadBalancing }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.LocalMySqlEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.LogsDirectorySizeLimit }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) MachineKey() SiteMachineKeyResponseOutput {
	return o.ApplyT(func(v SiteConfigResponse) SiteMachineKeyResponse { return v.MachineKey }).(SiteMachineKeyResponseOutput)
}

func (o SiteConfigResponseOutput) ManagedPipelineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ManagedPipelineMode }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.ManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.MinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.MinimumElasticInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.NetFrameworkVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.NodeVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.NumberOfWorkers }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PhpVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PowerShellVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.PreWarmedInstanceCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PublishingUsername }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Push() PushSettingsResponsePtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *PushSettingsResponse { return v.Push }).(PushSettingsResponsePtrOutput)
}

func (o SiteConfigResponseOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.PythonVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.RemoteDebuggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.RemoteDebuggingVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.RequestTracingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.RequestTracingExpirationTime }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []IpSecurityRestrictionResponse { return v.ScmIpSecurityRestrictions }).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponseOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.ScmIpSecurityRestrictionsUseMain }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ScmMinTlsVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.ScmType }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.TracingOptions }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.Use32BitWorkerProcess }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) VirtualApplications() VirtualApplicationResponseArrayOutput {
	return o.ApplyT(func(v SiteConfigResponse) []VirtualApplicationResponse { return v.VirtualApplications }).(VirtualApplicationResponseArrayOutput)
}

func (o SiteConfigResponseOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.VnetName }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.VnetPrivatePortsCount }).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponseOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.VnetRouteAllEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponseOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.WebsiteTimeZone }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *string { return v.WindowsFxVersion }).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponseOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *int { return v.XManagedServiceIdentityId }).(pulumi.IntPtrOutput)
}

type SiteConfigResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteConfigResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfigResponse)(nil)).Elem()
}

func (o SiteConfigResponsePtrOutput) ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput {
	return o
}

func (o SiteConfigResponsePtrOutput) ToSiteConfigResponsePtrOutputWithContext(ctx context.Context) SiteConfigResponsePtrOutput {
	return o
}

func (o SiteConfigResponsePtrOutput) Elem() SiteConfigResponseOutput {
	return o.ApplyT(func(v *SiteConfigResponse) SiteConfigResponse {
		if v != nil {
			return *v
		}
		var ret SiteConfigResponse
		return ret
	}).(SiteConfigResponseOutput)
}

func (o SiteConfigResponsePtrOutput) AcrUseManagedIdentityCreds() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AcrUseManagedIdentityCreds
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AcrUserManagedIdentityID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AcrUserManagedIdentityID
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AlwaysOn() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AlwaysOn
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ApiDefinition() ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ApiDefinitionInfoResponse {
		if v == nil {
			return nil
		}
		return v.ApiDefinition
	}).(ApiDefinitionInfoResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) ApiManagementConfig() ApiManagementConfigResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ApiManagementConfigResponse {
		if v == nil {
			return nil
		}
		return v.ApiManagementConfig
	}).(ApiManagementConfigResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) AppCommandLine() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppCommandLine
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AppSettings() NameValuePairResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []NameValuePairResponse {
		if v == nil {
			return nil
		}
		return v.AppSettings
	}).(NameValuePairResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) AutoHealEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.AutoHealEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AutoHealRules() AutoHealRulesResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *AutoHealRulesResponse {
		if v == nil {
			return nil
		}
		return v.AutoHealRules
	}).(AutoHealRulesResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) AutoSwapSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.AutoSwapSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) AzureStorageAccounts() AzureStorageInfoValueResponseMapOutput {
	return o.ApplyT(func(v *SiteConfigResponse) map[string]AzureStorageInfoValueResponse {
		if v == nil {
			return nil
		}
		return v.AzureStorageAccounts
	}).(AzureStorageInfoValueResponseMapOutput)
}

func (o SiteConfigResponsePtrOutput) ConnectionStrings() ConnStringInfoResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []ConnStringInfoResponse {
		if v == nil {
			return nil
		}
		return v.ConnectionStrings
	}).(ConnStringInfoResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) Cors() CorsSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *CorsSettingsResponse {
		if v == nil {
			return nil
		}
		return v.Cors
	}).(CorsSettingsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) DefaultDocuments() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []string {
		if v == nil {
			return nil
		}
		return v.DefaultDocuments
	}).(pulumi.StringArrayOutput)
}

func (o SiteConfigResponsePtrOutput) DetailedErrorLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DetailedErrorLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) DocumentRoot() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.DocumentRoot
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Experiments() ExperimentsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *ExperimentsResponse {
		if v == nil {
			return nil
		}
		return v.Experiments
	}).(ExperimentsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) FtpsState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.FtpsState
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) FunctionAppScaleLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.FunctionAppScaleLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) FunctionsRuntimeScaleMonitoringEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.FunctionsRuntimeScaleMonitoringEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) HandlerMappings() HandlerMappingResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []HandlerMappingResponse {
		if v == nil {
			return nil
		}
		return v.HandlerMappings
	}).(HandlerMappingResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) HealthCheckPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.HealthCheckPath
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Http20Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Http20Enabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) HttpLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.HttpLoggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) IpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []IpSecurityRestrictionResponse {
		if v == nil {
			return nil
		}
		return v.IpSecurityRestrictions
	}).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) JavaContainer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainer
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) JavaContainerVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaContainerVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) JavaVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.JavaVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) KeyVaultReferenceIdentity() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.KeyVaultReferenceIdentity
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Limits() SiteLimitsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *SiteLimitsResponse {
		if v == nil {
			return nil
		}
		return v.Limits
	}).(SiteLimitsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) LinuxFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.LinuxFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LoadBalancing() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.LoadBalancing
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LocalMySqlEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.LocalMySqlEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) LogsDirectorySizeLimit() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.LogsDirectorySizeLimit
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) MachineKey() SiteMachineKeyResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *SiteMachineKeyResponse {
		if v == nil {
			return nil
		}
		return &v.MachineKey
	}).(SiteMachineKeyResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) ManagedPipelineMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ManagedPipelineMode
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.ManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) MinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.MinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) MinimumElasticInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.MinimumElasticInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NetFrameworkVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NetFrameworkVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NodeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.NodeVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) NumberOfWorkers() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.NumberOfWorkers
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PhpVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PhpVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PowerShellVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PowerShellVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PreWarmedInstanceCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.PreWarmedInstanceCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) PublishingUsername() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PublishingUsername
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Push() PushSettingsResponsePtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *PushSettingsResponse {
		if v == nil {
			return nil
		}
		return v.Push
	}).(PushSettingsResponsePtrOutput)
}

func (o SiteConfigResponsePtrOutput) PythonVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.PythonVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RemoteDebuggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RemoteDebuggingVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RemoteDebuggingVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RequestTracingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequestTracingEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) RequestTracingExpirationTime() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.RequestTracingExpirationTime
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmIpSecurityRestrictions() IpSecurityRestrictionResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []IpSecurityRestrictionResponse {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictions
	}).(IpSecurityRestrictionResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) ScmIpSecurityRestrictionsUseMain() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.ScmIpSecurityRestrictionsUseMain
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmMinTlsVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScmMinTlsVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) ScmType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScmType
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) TracingOptions() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.TracingOptions
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) Use32BitWorkerProcess() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.Use32BitWorkerProcess
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VirtualApplications() VirtualApplicationResponseArrayOutput {
	return o.ApplyT(func(v *SiteConfigResponse) []VirtualApplicationResponse {
		if v == nil {
			return nil
		}
		return v.VirtualApplications
	}).(VirtualApplicationResponseArrayOutput)
}

func (o SiteConfigResponsePtrOutput) VnetName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.VnetName
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VnetPrivatePortsCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.VnetPrivatePortsCount
	}).(pulumi.IntPtrOutput)
}

func (o SiteConfigResponsePtrOutput) VnetRouteAllEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.VnetRouteAllEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WebsiteTimeZone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WebsiteTimeZone
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) WindowsFxVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *string {
		if v == nil {
			return nil
		}
		return v.WindowsFxVersion
	}).(pulumi.StringPtrOutput)
}

func (o SiteConfigResponsePtrOutput) XManagedServiceIdentityId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *int {
		if v == nil {
			return nil
		}
		return v.XManagedServiceIdentityId
	}).(pulumi.IntPtrOutput)
}

type SiteLimits struct {
	MaxDiskSizeInMb  *float64 `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    *float64 `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu *float64 `pulumi:"maxPercentageCpu"`
}





type SiteLimitsInput interface {
	pulumi.Input

	ToSiteLimitsOutput() SiteLimitsOutput
	ToSiteLimitsOutputWithContext(context.Context) SiteLimitsOutput
}

type SiteLimitsArgs struct {
	MaxDiskSizeInMb  pulumi.Float64PtrInput `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    pulumi.Float64PtrInput `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu pulumi.Float64PtrInput `pulumi:"maxPercentageCpu"`
}

func (SiteLimitsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimits)(nil)).Elem()
}

func (i SiteLimitsArgs) ToSiteLimitsOutput() SiteLimitsOutput {
	return i.ToSiteLimitsOutputWithContext(context.Background())
}

func (i SiteLimitsArgs) ToSiteLimitsOutputWithContext(ctx context.Context) SiteLimitsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsOutput)
}

func (i SiteLimitsArgs) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return i.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (i SiteLimitsArgs) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsOutput).ToSiteLimitsPtrOutputWithContext(ctx)
}









type SiteLimitsPtrInput interface {
	pulumi.Input

	ToSiteLimitsPtrOutput() SiteLimitsPtrOutput
	ToSiteLimitsPtrOutputWithContext(context.Context) SiteLimitsPtrOutput
}

type siteLimitsPtrType SiteLimitsArgs

func SiteLimitsPtr(v *SiteLimitsArgs) SiteLimitsPtrInput {
	return (*siteLimitsPtrType)(v)
}

func (*siteLimitsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimits)(nil)).Elem()
}

func (i *siteLimitsPtrType) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return i.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (i *siteLimitsPtrType) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsPtrOutput)
}

type SiteLimitsOutput struct{ *pulumi.OutputState }

func (SiteLimitsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimits)(nil)).Elem()
}

func (o SiteLimitsOutput) ToSiteLimitsOutput() SiteLimitsOutput {
	return o
}

func (o SiteLimitsOutput) ToSiteLimitsOutputWithContext(ctx context.Context) SiteLimitsOutput {
	return o
}

func (o SiteLimitsOutput) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return o.ToSiteLimitsPtrOutputWithContext(context.Background())
}

func (o SiteLimitsOutput) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLimits) *SiteLimits {
		return &v
	}).(SiteLimitsPtrOutput)
}

func (o SiteLimitsOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxDiskSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxMemoryInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimits) *float64 { return v.MaxPercentageCpu }).(pulumi.Float64PtrOutput)
}

type SiteLimitsPtrOutput struct{ *pulumi.OutputState }

func (SiteLimitsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimits)(nil)).Elem()
}

func (o SiteLimitsPtrOutput) ToSiteLimitsPtrOutput() SiteLimitsPtrOutput {
	return o
}

func (o SiteLimitsPtrOutput) ToSiteLimitsPtrOutputWithContext(ctx context.Context) SiteLimitsPtrOutput {
	return o
}

func (o SiteLimitsPtrOutput) Elem() SiteLimitsOutput {
	return o.ApplyT(func(v *SiteLimits) SiteLimits {
		if v != nil {
			return *v
		}
		var ret SiteLimits
		return ret
	}).(SiteLimitsOutput)
}

func (o SiteLimitsPtrOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDiskSizeInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsPtrOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxMemoryInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsPtrOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimits) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentageCpu
	}).(pulumi.Float64PtrOutput)
}

type SiteLimitsResponse struct {
	MaxDiskSizeInMb  *float64 `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    *float64 `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu *float64 `pulumi:"maxPercentageCpu"`
}

type SiteLimitsResponseOutput struct{ *pulumi.OutputState }

func (SiteLimitsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimitsResponse)(nil)).Elem()
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponseOutput() SiteLimitsResponseOutput {
	return o
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponseOutputWithContext(ctx context.Context) SiteLimitsResponseOutput {
	return o
}

func (o SiteLimitsResponseOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxDiskSizeInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponseOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxMemoryInMb }).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponseOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v SiteLimitsResponse) *float64 { return v.MaxPercentageCpu }).(pulumi.Float64PtrOutput)
}

type SiteLimitsResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteLimitsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimitsResponse)(nil)).Elem()
}

func (o SiteLimitsResponsePtrOutput) ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput {
	return o
}

func (o SiteLimitsResponsePtrOutput) ToSiteLimitsResponsePtrOutputWithContext(ctx context.Context) SiteLimitsResponsePtrOutput {
	return o
}

func (o SiteLimitsResponsePtrOutput) Elem() SiteLimitsResponseOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) SiteLimitsResponse {
		if v != nil {
			return *v
		}
		var ret SiteLimitsResponse
		return ret
	}).(SiteLimitsResponseOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxDiskSizeInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxDiskSizeInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxMemoryInMb() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxMemoryInMb
	}).(pulumi.Float64PtrOutput)
}

func (o SiteLimitsResponsePtrOutput) MaxPercentageCpu() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteLimitsResponse) *float64 {
		if v == nil {
			return nil
		}
		return v.MaxPercentageCpu
	}).(pulumi.Float64PtrOutput)
}

type SiteMachineKeyResponse struct {
	Decryption    *string `pulumi:"decryption"`
	DecryptionKey *string `pulumi:"decryptionKey"`
	Validation    *string `pulumi:"validation"`
	ValidationKey *string `pulumi:"validationKey"`
}

type SiteMachineKeyResponseOutput struct{ *pulumi.OutputState }

func (SiteMachineKeyResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMachineKeyResponse)(nil)).Elem()
}

func (o SiteMachineKeyResponseOutput) ToSiteMachineKeyResponseOutput() SiteMachineKeyResponseOutput {
	return o
}

func (o SiteMachineKeyResponseOutput) ToSiteMachineKeyResponseOutputWithContext(ctx context.Context) SiteMachineKeyResponseOutput {
	return o
}

func (o SiteMachineKeyResponseOutput) Decryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.Decryption }).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponseOutput) DecryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.DecryptionKey }).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponseOutput) Validation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.Validation }).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponseOutput) ValidationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteMachineKeyResponse) *string { return v.ValidationKey }).(pulumi.StringPtrOutput)
}

type SiteMachineKeyResponsePtrOutput struct{ *pulumi.OutputState }

func (SiteMachineKeyResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMachineKeyResponse)(nil)).Elem()
}

func (o SiteMachineKeyResponsePtrOutput) ToSiteMachineKeyResponsePtrOutput() SiteMachineKeyResponsePtrOutput {
	return o
}

func (o SiteMachineKeyResponsePtrOutput) ToSiteMachineKeyResponsePtrOutputWithContext(ctx context.Context) SiteMachineKeyResponsePtrOutput {
	return o
}

func (o SiteMachineKeyResponsePtrOutput) Elem() SiteMachineKeyResponseOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) SiteMachineKeyResponse {
		if v != nil {
			return *v
		}
		var ret SiteMachineKeyResponse
		return ret
	}).(SiteMachineKeyResponseOutput)
}

func (o SiteMachineKeyResponsePtrOutput) Decryption() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Decryption
	}).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponsePtrOutput) DecryptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.DecryptionKey
	}).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponsePtrOutput) Validation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.Validation
	}).(pulumi.StringPtrOutput)
}

func (o SiteMachineKeyResponsePtrOutput) ValidationKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteMachineKeyResponse) *string {
		if v == nil {
			return nil
		}
		return v.ValidationKey
	}).(pulumi.StringPtrOutput)
}

type SkuCapacity struct {
	Default        *int    `pulumi:"default"`
	ElasticMaximum *int    `pulumi:"elasticMaximum"`
	Maximum        *int    `pulumi:"maximum"`
	Minimum        *int    `pulumi:"minimum"`
	ScaleType      *string `pulumi:"scaleType"`
}





type SkuCapacityInput interface {
	pulumi.Input

	ToSkuCapacityOutput() SkuCapacityOutput
	ToSkuCapacityOutputWithContext(context.Context) SkuCapacityOutput
}

type SkuCapacityArgs struct {
	Default        pulumi.IntPtrInput    `pulumi:"default"`
	ElasticMaximum pulumi.IntPtrInput    `pulumi:"elasticMaximum"`
	Maximum        pulumi.IntPtrInput    `pulumi:"maximum"`
	Minimum        pulumi.IntPtrInput    `pulumi:"minimum"`
	ScaleType      pulumi.StringPtrInput `pulumi:"scaleType"`
}

func (SkuCapacityArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacity)(nil)).Elem()
}

func (i SkuCapacityArgs) ToSkuCapacityOutput() SkuCapacityOutput {
	return i.ToSkuCapacityOutputWithContext(context.Background())
}

func (i SkuCapacityArgs) ToSkuCapacityOutputWithContext(ctx context.Context) SkuCapacityOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityOutput)
}

func (i SkuCapacityArgs) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return i.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (i SkuCapacityArgs) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityOutput).ToSkuCapacityPtrOutputWithContext(ctx)
}









type SkuCapacityPtrInput interface {
	pulumi.Input

	ToSkuCapacityPtrOutput() SkuCapacityPtrOutput
	ToSkuCapacityPtrOutputWithContext(context.Context) SkuCapacityPtrOutput
}

type skuCapacityPtrType SkuCapacityArgs

func SkuCapacityPtr(v *SkuCapacityArgs) SkuCapacityPtrInput {
	return (*skuCapacityPtrType)(v)
}

func (*skuCapacityPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacity)(nil)).Elem()
}

func (i *skuCapacityPtrType) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return i.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (i *skuCapacityPtrType) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityPtrOutput)
}

type SkuCapacityOutput struct{ *pulumi.OutputState }

func (SkuCapacityOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacity)(nil)).Elem()
}

func (o SkuCapacityOutput) ToSkuCapacityOutput() SkuCapacityOutput {
	return o
}

func (o SkuCapacityOutput) ToSkuCapacityOutputWithContext(ctx context.Context) SkuCapacityOutput {
	return o
}

func (o SkuCapacityOutput) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return o.ToSkuCapacityPtrOutputWithContext(context.Background())
}

func (o SkuCapacityOutput) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuCapacity) *SkuCapacity {
		return &v
	}).(SkuCapacityPtrOutput)
}

func (o SkuCapacityOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Default }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.ElasticMaximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Maximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *int { return v.Minimum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapacity) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type SkuCapacityPtrOutput struct{ *pulumi.OutputState }

func (SkuCapacityPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacity)(nil)).Elem()
}

func (o SkuCapacityPtrOutput) ToSkuCapacityPtrOutput() SkuCapacityPtrOutput {
	return o
}

func (o SkuCapacityPtrOutput) ToSkuCapacityPtrOutputWithContext(ctx context.Context) SkuCapacityPtrOutput {
	return o
}

func (o SkuCapacityPtrOutput) Elem() SkuCapacityOutput {
	return o.ApplyT(func(v *SkuCapacity) SkuCapacity {
		if v != nil {
			return *v
		}
		var ret SkuCapacity
		return ret
	}).(SkuCapacityOutput)
}

func (o SkuCapacityPtrOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.ElasticMaximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *int {
		if v == nil {
			return nil
		}
		return v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityPtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuCapacity) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type SkuCapacityResponse struct {
	Default        *int    `pulumi:"default"`
	ElasticMaximum *int    `pulumi:"elasticMaximum"`
	Maximum        *int    `pulumi:"maximum"`
	Minimum        *int    `pulumi:"minimum"`
	ScaleType      *string `pulumi:"scaleType"`
}

type SkuCapacityResponseOutput struct{ *pulumi.OutputState }

func (SkuCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacityResponse)(nil)).Elem()
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponseOutput() SkuCapacityResponseOutput {
	return o
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponseOutputWithContext(ctx context.Context) SkuCapacityResponseOutput {
	return o
}

func (o SkuCapacityResponseOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Default }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.ElasticMaximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Maximum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Minimum }).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponseOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *string { return v.ScaleType }).(pulumi.StringPtrOutput)
}

type SkuCapacityResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuCapacityResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacityResponse)(nil)).Elem()
}

func (o SkuCapacityResponsePtrOutput) ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput {
	return o
}

func (o SkuCapacityResponsePtrOutput) ToSkuCapacityResponsePtrOutputWithContext(ctx context.Context) SkuCapacityResponsePtrOutput {
	return o
}

func (o SkuCapacityResponsePtrOutput) Elem() SkuCapacityResponseOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) SkuCapacityResponse {
		if v != nil {
			return *v
		}
		var ret SkuCapacityResponse
		return ret
	}).(SkuCapacityResponseOutput)
}

func (o SkuCapacityResponsePtrOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Default
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) ElasticMaximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.ElasticMaximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) Maximum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Maximum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) Minimum() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *int {
		if v == nil {
			return nil
		}
		return v.Minimum
	}).(pulumi.IntPtrOutput)
}

func (o SkuCapacityResponsePtrOutput) ScaleType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuCapacityResponse) *string {
		if v == nil {
			return nil
		}
		return v.ScaleType
	}).(pulumi.StringPtrOutput)
}

type SkuDescription struct {
	Capabilities []Capability `pulumi:"capabilities"`
	Capacity     *int         `pulumi:"capacity"`
	Family       *string      `pulumi:"family"`
	Locations    []string     `pulumi:"locations"`
	Name         *string      `pulumi:"name"`
	Size         *string      `pulumi:"size"`
	SkuCapacity  *SkuCapacity `pulumi:"skuCapacity"`
	Tier         *string      `pulumi:"tier"`
}





type SkuDescriptionInput interface {
	pulumi.Input

	ToSkuDescriptionOutput() SkuDescriptionOutput
	ToSkuDescriptionOutputWithContext(context.Context) SkuDescriptionOutput
}

type SkuDescriptionArgs struct {
	Capabilities CapabilityArrayInput    `pulumi:"capabilities"`
	Capacity     pulumi.IntPtrInput      `pulumi:"capacity"`
	Family       pulumi.StringPtrInput   `pulumi:"family"`
	Locations    pulumi.StringArrayInput `pulumi:"locations"`
	Name         pulumi.StringPtrInput   `pulumi:"name"`
	Size         pulumi.StringPtrInput   `pulumi:"size"`
	SkuCapacity  SkuCapacityPtrInput     `pulumi:"skuCapacity"`
	Tier         pulumi.StringPtrInput   `pulumi:"tier"`
}

func (SkuDescriptionArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return i.ToSkuDescriptionOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput)
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i SkuDescriptionArgs) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionOutput).ToSkuDescriptionPtrOutputWithContext(ctx)
}









type SkuDescriptionPtrInput interface {
	pulumi.Input

	ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput
	ToSkuDescriptionPtrOutputWithContext(context.Context) SkuDescriptionPtrOutput
}

type skuDescriptionPtrType SkuDescriptionArgs

func SkuDescriptionPtr(v *SkuDescriptionArgs) SkuDescriptionPtrInput {
	return (*skuDescriptionPtrType)(v)
}

func (*skuDescriptionPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return i.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (i *skuDescriptionPtrType) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionPtrOutput)
}

type SkuDescriptionOutput struct{ *pulumi.OutputState }

func (SkuDescriptionOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutput() SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionOutputWithContext(ctx context.Context) SkuDescriptionOutput {
	return o
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o.ToSkuDescriptionPtrOutputWithContext(context.Background())
}

func (o SkuDescriptionOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescription) *SkuDescription {
		return &v
	}).(SkuDescriptionPtrOutput)
}

func (o SkuDescriptionOutput) Capabilities() CapabilityArrayOutput {
	return o.ApplyT(func(v SkuDescription) []Capability { return v.Capabilities }).(CapabilityArrayOutput)
}

func (o SkuDescriptionOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescription) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuDescription) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionOutput) SkuCapacity() SkuCapacityPtrOutput {
	return o.ApplyT(func(v SkuDescription) *SkuCapacity { return v.SkuCapacity }).(SkuCapacityPtrOutput)
}

func (o SkuDescriptionOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescription) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionPtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescription)(nil)).Elem()
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutput() SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) ToSkuDescriptionPtrOutputWithContext(ctx context.Context) SkuDescriptionPtrOutput {
	return o
}

func (o SkuDescriptionPtrOutput) Elem() SkuDescriptionOutput {
	return o.ApplyT(func(v *SkuDescription) SkuDescription {
		if v != nil {
			return *v
		}
		var ret SkuDescription
		return ret
	}).(SkuDescriptionOutput)
}

func (o SkuDescriptionPtrOutput) Capabilities() CapabilityArrayOutput {
	return o.ApplyT(func(v *SkuDescription) []Capability {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(CapabilityArrayOutput)
}

func (o SkuDescriptionPtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionPtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuDescription) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionPtrOutput) SkuCapacity() SkuCapacityPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *SkuCapacity {
		if v == nil {
			return nil
		}
		return v.SkuCapacity
	}).(SkuCapacityPtrOutput)
}

func (o SkuDescriptionPtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescription) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponse struct {
	Capabilities []CapabilityResponse `pulumi:"capabilities"`
	Capacity     *int                 `pulumi:"capacity"`
	Family       *string              `pulumi:"family"`
	Locations    []string             `pulumi:"locations"`
	Name         *string              `pulumi:"name"`
	Size         *string              `pulumi:"size"`
	SkuCapacity  *SkuCapacityResponse `pulumi:"skuCapacity"`
	Tier         *string              `pulumi:"tier"`
}

type SkuDescriptionResponseOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return o
}

func (o SkuDescriptionResponseOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) []CapabilityResponse { return v.Capabilities }).(CapabilityResponseArrayOutput)
}

func (o SkuDescriptionResponseOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *int { return v.Capacity }).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponseOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Family }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) []string { return v.Locations }).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Size }).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponseOutput) SkuCapacity() SkuCapacityResponsePtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *SkuCapacityResponse { return v.SkuCapacity }).(SkuCapacityResponsePtrOutput)
}

func (o SkuDescriptionResponseOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkuDescriptionResponse) *string { return v.Tier }).(pulumi.StringPtrOutput)
}

type SkuDescriptionResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuDescriptionResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o
}

func (o SkuDescriptionResponsePtrOutput) Elem() SkuDescriptionResponseOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) SkuDescriptionResponse {
		if v != nil {
			return *v
		}
		var ret SkuDescriptionResponse
		return ret
	}).(SkuDescriptionResponseOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capabilities() CapabilityResponseArrayOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) []CapabilityResponse {
		if v == nil {
			return nil
		}
		return v.Capabilities
	}).(CapabilityResponseArrayOutput)
}

func (o SkuDescriptionResponsePtrOutput) Capacity() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *int {
		if v == nil {
			return nil
		}
		return v.Capacity
	}).(pulumi.IntPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Family() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Family
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Locations() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) []string {
		if v == nil {
			return nil
		}
		return v.Locations
	}).(pulumi.StringArrayOutput)
}

func (o SkuDescriptionResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Size() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Size
	}).(pulumi.StringPtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) SkuCapacity() SkuCapacityResponsePtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *SkuCapacityResponse {
		if v == nil {
			return nil
		}
		return v.SkuCapacity
	}).(SkuCapacityResponsePtrOutput)
}

func (o SkuDescriptionResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuDescriptionResponse) *string {
		if v == nil {
			return nil
		}
		return v.Tier
	}).(pulumi.StringPtrOutput)
}

type SlotSwapStatusResponse struct {
	DestinationSlotName string `pulumi:"destinationSlotName"`
	SourceSlotName      string `pulumi:"sourceSlotName"`
	TimestampUtc        string `pulumi:"timestampUtc"`
}

type SlotSwapStatusResponseOutput struct{ *pulumi.OutputState }

func (SlotSwapStatusResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotSwapStatusResponse)(nil)).Elem()
}

func (o SlotSwapStatusResponseOutput) ToSlotSwapStatusResponseOutput() SlotSwapStatusResponseOutput {
	return o
}

func (o SlotSwapStatusResponseOutput) ToSlotSwapStatusResponseOutputWithContext(ctx context.Context) SlotSwapStatusResponseOutput {
	return o
}

func (o SlotSwapStatusResponseOutput) DestinationSlotName() pulumi.StringOutput {
	return o.ApplyT(func(v SlotSwapStatusResponse) string { return v.DestinationSlotName }).(pulumi.StringOutput)
}

func (o SlotSwapStatusResponseOutput) SourceSlotName() pulumi.StringOutput {
	return o.ApplyT(func(v SlotSwapStatusResponse) string { return v.SourceSlotName }).(pulumi.StringOutput)
}

func (o SlotSwapStatusResponseOutput) TimestampUtc() pulumi.StringOutput {
	return o.ApplyT(func(v SlotSwapStatusResponse) string { return v.TimestampUtc }).(pulumi.StringOutput)
}

type SlowRequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}





type SlowRequestsBasedTriggerInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput
	ToSlowRequestsBasedTriggerOutputWithContext(context.Context) SlowRequestsBasedTriggerOutput
}

type SlowRequestsBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Path         pulumi.StringPtrInput `pulumi:"path"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	TimeTaken    pulumi.StringPtrInput `pulumi:"timeTaken"`
}

func (SlowRequestsBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput {
	return i.ToSlowRequestsBasedTriggerOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerOutput)
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return i.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArgs) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerOutput).ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx)
}









type SlowRequestsBasedTriggerPtrInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput
	ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Context) SlowRequestsBasedTriggerPtrOutput
}

type slowRequestsBasedTriggerPtrType SlowRequestsBasedTriggerArgs

func SlowRequestsBasedTriggerPtr(v *SlowRequestsBasedTriggerArgs) SlowRequestsBasedTriggerPtrInput {
	return (*slowRequestsBasedTriggerPtrType)(v)
}

func (*slowRequestsBasedTriggerPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i *slowRequestsBasedTriggerPtrType) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return i.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (i *slowRequestsBasedTriggerPtrType) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerPtrOutput)
}





type SlowRequestsBasedTriggerArrayInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerArrayOutput() SlowRequestsBasedTriggerArrayOutput
	ToSlowRequestsBasedTriggerArrayOutputWithContext(context.Context) SlowRequestsBasedTriggerArrayOutput
}

type SlowRequestsBasedTriggerArray []SlowRequestsBasedTriggerInput

func (SlowRequestsBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlowRequestsBasedTrigger)(nil)).Elem()
}

func (i SlowRequestsBasedTriggerArray) ToSlowRequestsBasedTriggerArrayOutput() SlowRequestsBasedTriggerArrayOutput {
	return i.ToSlowRequestsBasedTriggerArrayOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerArray) ToSlowRequestsBasedTriggerArrayOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerArrayOutput)
}

type SlowRequestsBasedTriggerOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerOutput() SlowRequestsBasedTriggerOutput {
	return o
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerOutput {
	return o
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return o.ToSlowRequestsBasedTriggerPtrOutputWithContext(context.Background())
}

func (o SlowRequestsBasedTriggerOutput) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlowRequestsBasedTrigger) *SlowRequestsBasedTrigger {
		return &v
	}).(SlowRequestsBasedTriggerPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTrigger) *string { return v.TimeTaken }).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerPtrOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerPtrOutput) ToSlowRequestsBasedTriggerPtrOutput() SlowRequestsBasedTriggerPtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerPtrOutput) ToSlowRequestsBasedTriggerPtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerPtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerPtrOutput) Elem() SlowRequestsBasedTriggerOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) SlowRequestsBasedTrigger {
		if v != nil {
			return *v
		}
		var ret SlowRequestsBasedTrigger
		return ret
	}).(SlowRequestsBasedTriggerOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerPtrOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTrigger) *string {
		if v == nil {
			return nil
		}
		return v.TimeTaken
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlowRequestsBasedTrigger)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerArrayOutput) ToSlowRequestsBasedTriggerArrayOutput() SlowRequestsBasedTriggerArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerArrayOutput) ToSlowRequestsBasedTriggerArrayOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerArrayOutput) Index(i pulumi.IntInput) SlowRequestsBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SlowRequestsBasedTrigger {
		return vs[0].([]SlowRequestsBasedTrigger)[vs[1].(int)]
	}).(SlowRequestsBasedTriggerOutput)
}

type SlowRequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}

type SlowRequestsBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponseOutput() SlowRequestsBasedTriggerResponseOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponseOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *string { return v.TimeTaken }).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerResponsePtrOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponsePtrOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Elem() SlowRequestsBasedTriggerResponseOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) SlowRequestsBasedTriggerResponse {
		if v != nil {
			return *v
		}
		var ret SlowRequestsBasedTriggerResponse
		return ret
	}).(SlowRequestsBasedTriggerResponseOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *int {
		if v == nil {
			return nil
		}
		return v.Count
	}).(pulumi.IntPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.Path
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeInterval
	}).(pulumi.StringPtrOutput)
}

func (o SlowRequestsBasedTriggerResponsePtrOutput) TimeTaken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlowRequestsBasedTriggerResponse) *string {
		if v == nil {
			return nil
		}
		return v.TimeTaken
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (SlowRequestsBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (o SlowRequestsBasedTriggerResponseArrayOutput) ToSlowRequestsBasedTriggerResponseArrayOutput() SlowRequestsBasedTriggerResponseArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseArrayOutput) ToSlowRequestsBasedTriggerResponseArrayOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponseArrayOutput {
	return o
}

func (o SlowRequestsBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) SlowRequestsBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SlowRequestsBasedTriggerResponse {
		return vs[0].([]SlowRequestsBasedTriggerResponse)[vs[1].(int)]
	}).(SlowRequestsBasedTriggerResponseOutput)
}

type StaticSiteBuildProperties struct {
	ApiBuildCommand                    *string `pulumi:"apiBuildCommand"`
	ApiLocation                        *string `pulumi:"apiLocation"`
	AppArtifactLocation                *string `pulumi:"appArtifactLocation"`
	AppBuildCommand                    *string `pulumi:"appBuildCommand"`
	AppLocation                        *string `pulumi:"appLocation"`
	GithubActionSecretNameOverride     *string `pulumi:"githubActionSecretNameOverride"`
	OutputLocation                     *string `pulumi:"outputLocation"`
	SkipGithubActionWorkflowGeneration *bool   `pulumi:"skipGithubActionWorkflowGeneration"`
}





type StaticSiteBuildPropertiesInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput
	ToStaticSiteBuildPropertiesOutputWithContext(context.Context) StaticSiteBuildPropertiesOutput
}

type StaticSiteBuildPropertiesArgs struct {
	ApiBuildCommand                    pulumi.StringPtrInput `pulumi:"apiBuildCommand"`
	ApiLocation                        pulumi.StringPtrInput `pulumi:"apiLocation"`
	AppArtifactLocation                pulumi.StringPtrInput `pulumi:"appArtifactLocation"`
	AppBuildCommand                    pulumi.StringPtrInput `pulumi:"appBuildCommand"`
	AppLocation                        pulumi.StringPtrInput `pulumi:"appLocation"`
	GithubActionSecretNameOverride     pulumi.StringPtrInput `pulumi:"githubActionSecretNameOverride"`
	OutputLocation                     pulumi.StringPtrInput `pulumi:"outputLocation"`
	SkipGithubActionWorkflowGeneration pulumi.BoolPtrInput   `pulumi:"skipGithubActionWorkflowGeneration"`
}

func (StaticSiteBuildPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildProperties)(nil)).Elem()
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput {
	return i.ToStaticSiteBuildPropertiesOutputWithContext(context.Background())
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesOutput)
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return i.ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Background())
}

func (i StaticSiteBuildPropertiesArgs) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesOutput).ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx)
}









type StaticSiteBuildPropertiesPtrInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput
	ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Context) StaticSiteBuildPropertiesPtrOutput
}

type staticSiteBuildPropertiesPtrType StaticSiteBuildPropertiesArgs

func StaticSiteBuildPropertiesPtr(v *StaticSiteBuildPropertiesArgs) StaticSiteBuildPropertiesPtrInput {
	return (*staticSiteBuildPropertiesPtrType)(v)
}

func (*staticSiteBuildPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildProperties)(nil)).Elem()
}

func (i *staticSiteBuildPropertiesPtrType) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return i.ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Background())
}

func (i *staticSiteBuildPropertiesPtrType) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesPtrOutput)
}

type StaticSiteBuildPropertiesOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildProperties)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput {
	return o
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesOutput {
	return o
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return o.ToStaticSiteBuildPropertiesPtrOutputWithContext(context.Background())
}

func (o StaticSiteBuildPropertiesOutput) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StaticSiteBuildProperties) *StaticSiteBuildProperties {
		return &v
	}).(StaticSiteBuildPropertiesPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.ApiBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.ApiLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppArtifactLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.GithubActionSecretNameOverride }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.OutputLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *bool { return v.SkipGithubActionWorkflowGeneration }).(pulumi.BoolPtrOutput)
}

type StaticSiteBuildPropertiesPtrOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildProperties)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesPtrOutput) ToStaticSiteBuildPropertiesPtrOutput() StaticSiteBuildPropertiesPtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesPtrOutput) ToStaticSiteBuildPropertiesPtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesPtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesPtrOutput) Elem() StaticSiteBuildPropertiesOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) StaticSiteBuildProperties {
		if v != nil {
			return *v
		}
		var ret StaticSiteBuildProperties
		return ret
	}).(StaticSiteBuildPropertiesOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppArtifactLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.GithubActionSecretNameOverride
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.OutputLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesPtrOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *bool {
		if v == nil {
			return nil
		}
		return v.SkipGithubActionWorkflowGeneration
	}).(pulumi.BoolPtrOutput)
}

type StaticSiteBuildPropertiesResponse struct {
	ApiBuildCommand                    *string `pulumi:"apiBuildCommand"`
	ApiLocation                        *string `pulumi:"apiLocation"`
	AppArtifactLocation                *string `pulumi:"appArtifactLocation"`
	AppBuildCommand                    *string `pulumi:"appBuildCommand"`
	AppLocation                        *string `pulumi:"appLocation"`
	GithubActionSecretNameOverride     *string `pulumi:"githubActionSecretNameOverride"`
	OutputLocation                     *string `pulumi:"outputLocation"`
	SkipGithubActionWorkflowGeneration *bool   `pulumi:"skipGithubActionWorkflowGeneration"`
}

type StaticSiteBuildPropertiesResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildPropertiesResponse)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesResponseOutput) ToStaticSiteBuildPropertiesResponseOutput() StaticSiteBuildPropertiesResponseOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponseOutput) ToStaticSiteBuildPropertiesResponseOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponseOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponseOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.ApiBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.ApiLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppArtifactLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppBuildCommand }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.GithubActionSecretNameOverride }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.OutputLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *bool { return v.SkipGithubActionWorkflowGeneration }).(pulumi.BoolPtrOutput)
}

type StaticSiteBuildPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (StaticSiteBuildPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildPropertiesResponse)(nil)).Elem()
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ToStaticSiteBuildPropertiesResponsePtrOutput() StaticSiteBuildPropertiesResponsePtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponsePtrOutput {
	return o
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) Elem() StaticSiteBuildPropertiesResponseOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) StaticSiteBuildPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret StaticSiteBuildPropertiesResponse
		return ret
	}).(StaticSiteBuildPropertiesResponseOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ApiBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppArtifactLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppBuildCommand() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppBuildCommand
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) GithubActionSecretNameOverride() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GithubActionSecretNameOverride
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) OutputLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.OutputLocation
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponsePtrOutput) SkipGithubActionWorkflowGeneration() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.SkipGithubActionWorkflowGeneration
	}).(pulumi.BoolPtrOutput)
}

type StaticSiteLinkedBackendResponse struct {
	BackendResourceId *string `pulumi:"backendResourceId"`
	CreatedOn         string  `pulumi:"createdOn"`
	ProvisioningState string  `pulumi:"provisioningState"`
	Region            *string `pulumi:"region"`
}

type StaticSiteLinkedBackendResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteLinkedBackendResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteLinkedBackendResponse)(nil)).Elem()
}

func (o StaticSiteLinkedBackendResponseOutput) ToStaticSiteLinkedBackendResponseOutput() StaticSiteLinkedBackendResponseOutput {
	return o
}

func (o StaticSiteLinkedBackendResponseOutput) ToStaticSiteLinkedBackendResponseOutputWithContext(ctx context.Context) StaticSiteLinkedBackendResponseOutput {
	return o
}

func (o StaticSiteLinkedBackendResponseOutput) BackendResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteLinkedBackendResponse) *string { return v.BackendResourceId }).(pulumi.StringPtrOutput)
}

func (o StaticSiteLinkedBackendResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteLinkedBackendResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteLinkedBackendResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o StaticSiteLinkedBackendResponseOutput) Region() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteLinkedBackendResponse) *string { return v.Region }).(pulumi.StringPtrOutput)
}

type StaticSiteLinkedBackendResponseArrayOutput struct{ *pulumi.OutputState }

func (StaticSiteLinkedBackendResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSiteLinkedBackendResponse)(nil)).Elem()
}

func (o StaticSiteLinkedBackendResponseArrayOutput) ToStaticSiteLinkedBackendResponseArrayOutput() StaticSiteLinkedBackendResponseArrayOutput {
	return o
}

func (o StaticSiteLinkedBackendResponseArrayOutput) ToStaticSiteLinkedBackendResponseArrayOutputWithContext(ctx context.Context) StaticSiteLinkedBackendResponseArrayOutput {
	return o
}

func (o StaticSiteLinkedBackendResponseArrayOutput) Index(i pulumi.IntInput) StaticSiteLinkedBackendResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticSiteLinkedBackendResponse {
		return vs[0].([]StaticSiteLinkedBackendResponse)[vs[1].(int)]
	}).(StaticSiteLinkedBackendResponseOutput)
}

type StaticSiteTemplateOptions struct {
	Description           *string `pulumi:"description"`
	IsPrivate             *bool   `pulumi:"isPrivate"`
	Owner                 *string `pulumi:"owner"`
	RepositoryName        *string `pulumi:"repositoryName"`
	TemplateRepositoryUrl *string `pulumi:"templateRepositoryUrl"`
}





type StaticSiteTemplateOptionsInput interface {
	pulumi.Input

	ToStaticSiteTemplateOptionsOutput() StaticSiteTemplateOptionsOutput
	ToStaticSiteTemplateOptionsOutputWithContext(context.Context) StaticSiteTemplateOptionsOutput
}

type StaticSiteTemplateOptionsArgs struct {
	Description           pulumi.StringPtrInput `pulumi:"description"`
	IsPrivate             pulumi.BoolPtrInput   `pulumi:"isPrivate"`
	Owner                 pulumi.StringPtrInput `pulumi:"owner"`
	RepositoryName        pulumi.StringPtrInput `pulumi:"repositoryName"`
	TemplateRepositoryUrl pulumi.StringPtrInput `pulumi:"templateRepositoryUrl"`
}

func (StaticSiteTemplateOptionsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteTemplateOptions)(nil)).Elem()
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsOutput() StaticSiteTemplateOptionsOutput {
	return i.ToStaticSiteTemplateOptionsOutputWithContext(context.Background())
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteTemplateOptionsOutput)
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return i.ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Background())
}

func (i StaticSiteTemplateOptionsArgs) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteTemplateOptionsOutput).ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx)
}









type StaticSiteTemplateOptionsPtrInput interface {
	pulumi.Input

	ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput
	ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Context) StaticSiteTemplateOptionsPtrOutput
}

type staticSiteTemplateOptionsPtrType StaticSiteTemplateOptionsArgs

func StaticSiteTemplateOptionsPtr(v *StaticSiteTemplateOptionsArgs) StaticSiteTemplateOptionsPtrInput {
	return (*staticSiteTemplateOptionsPtrType)(v)
}

func (*staticSiteTemplateOptionsPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteTemplateOptions)(nil)).Elem()
}

func (i *staticSiteTemplateOptionsPtrType) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return i.ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Background())
}

func (i *staticSiteTemplateOptionsPtrType) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteTemplateOptionsPtrOutput)
}

type StaticSiteTemplateOptionsOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteTemplateOptions)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsOutput() StaticSiteTemplateOptionsOutput {
	return o
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsOutput {
	return o
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return o.ToStaticSiteTemplateOptionsPtrOutputWithContext(context.Background())
}

func (o StaticSiteTemplateOptionsOutput) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StaticSiteTemplateOptions) *StaticSiteTemplateOptions {
		return &v
	}).(StaticSiteTemplateOptionsPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *bool { return v.IsPrivate }).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptions) *string { return v.TemplateRepositoryUrl }).(pulumi.StringPtrOutput)
}

type StaticSiteTemplateOptionsPtrOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteTemplateOptions)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsPtrOutput) ToStaticSiteTemplateOptionsPtrOutput() StaticSiteTemplateOptionsPtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsPtrOutput) ToStaticSiteTemplateOptionsPtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsPtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsPtrOutput) Elem() StaticSiteTemplateOptionsOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) StaticSiteTemplateOptions {
		if v != nil {
			return *v
		}
		var ret StaticSiteTemplateOptions
		return ret
	}).(StaticSiteTemplateOptionsOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *bool {
		if v == nil {
			return nil
		}
		return v.IsPrivate
	}).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.Owner
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsPtrOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptions) *string {
		if v == nil {
			return nil
		}
		return v.TemplateRepositoryUrl
	}).(pulumi.StringPtrOutput)
}

type StaticSiteTemplateOptionsResponse struct {
	Description           *string `pulumi:"description"`
	IsPrivate             *bool   `pulumi:"isPrivate"`
	Owner                 *string `pulumi:"owner"`
	RepositoryName        *string `pulumi:"repositoryName"`
	TemplateRepositoryUrl *string `pulumi:"templateRepositoryUrl"`
}

type StaticSiteTemplateOptionsResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteTemplateOptionsResponse)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsResponseOutput) ToStaticSiteTemplateOptionsResponseOutput() StaticSiteTemplateOptionsResponseOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponseOutput) ToStaticSiteTemplateOptionsResponseOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsResponseOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *bool { return v.IsPrivate }).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.Owner }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.RepositoryName }).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponseOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteTemplateOptionsResponse) *string { return v.TemplateRepositoryUrl }).(pulumi.StringPtrOutput)
}

type StaticSiteTemplateOptionsResponsePtrOutput struct{ *pulumi.OutputState }

func (StaticSiteTemplateOptionsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteTemplateOptionsResponse)(nil)).Elem()
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) ToStaticSiteTemplateOptionsResponsePtrOutput() StaticSiteTemplateOptionsResponsePtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) ToStaticSiteTemplateOptionsResponsePtrOutputWithContext(ctx context.Context) StaticSiteTemplateOptionsResponsePtrOutput {
	return o
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) Elem() StaticSiteTemplateOptionsResponseOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) StaticSiteTemplateOptionsResponse {
		if v != nil {
			return *v
		}
		var ret StaticSiteTemplateOptionsResponse
		return ret
	}).(StaticSiteTemplateOptionsResponseOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) IsPrivate() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsPrivate
	}).(pulumi.BoolPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) Owner() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.Owner
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) RepositoryName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.RepositoryName
	}).(pulumi.StringPtrOutput)
}

func (o StaticSiteTemplateOptionsResponsePtrOutput) TemplateRepositoryUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteTemplateOptionsResponse) *string {
		if v == nil {
			return nil
		}
		return v.TemplateRepositoryUrl
	}).(pulumi.StringPtrOutput)
}

type StaticSiteUserARMResourceResponse struct {
	DisplayName string  `pulumi:"displayName"`
	Id          string  `pulumi:"id"`
	Kind        *string `pulumi:"kind"`
	Name        string  `pulumi:"name"`
	Provider    string  `pulumi:"provider"`
	Roles       *string `pulumi:"roles"`
	Type        string  `pulumi:"type"`
	UserId      string  `pulumi:"userId"`
}

type StaticSiteUserARMResourceResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteUserARMResourceResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteUserARMResourceResponse)(nil)).Elem()
}

func (o StaticSiteUserARMResourceResponseOutput) ToStaticSiteUserARMResourceResponseOutput() StaticSiteUserARMResourceResponseOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseOutput) ToStaticSiteUserARMResourceResponseOutputWithContext(ctx context.Context) StaticSiteUserARMResourceResponseOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Provider() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Provider }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Roles() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) *string { return v.Roles }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.Type }).(pulumi.StringOutput)
}

func (o StaticSiteUserARMResourceResponseOutput) UserId() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserARMResourceResponse) string { return v.UserId }).(pulumi.StringOutput)
}

type StaticSiteUserARMResourceResponseArrayOutput struct{ *pulumi.OutputState }

func (StaticSiteUserARMResourceResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSiteUserARMResourceResponse)(nil)).Elem()
}

func (o StaticSiteUserARMResourceResponseArrayOutput) ToStaticSiteUserARMResourceResponseArrayOutput() StaticSiteUserARMResourceResponseArrayOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseArrayOutput) ToStaticSiteUserARMResourceResponseArrayOutputWithContext(ctx context.Context) StaticSiteUserARMResourceResponseArrayOutput {
	return o
}

func (o StaticSiteUserARMResourceResponseArrayOutput) Index(i pulumi.IntInput) StaticSiteUserARMResourceResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticSiteUserARMResourceResponse {
		return vs[0].([]StaticSiteUserARMResourceResponse)[vs[1].(int)]
	}).(StaticSiteUserARMResourceResponseOutput)
}

type StaticSiteUserProvidedFunctionAppResponse struct {
	CreatedOn             string  `pulumi:"createdOn"`
	FunctionAppRegion     *string `pulumi:"functionAppRegion"`
	FunctionAppResourceId *string `pulumi:"functionAppResourceId"`
	Id                    string  `pulumi:"id"`
	Kind                  *string `pulumi:"kind"`
	Name                  string  `pulumi:"name"`
	Type                  string  `pulumi:"type"`
}

type StaticSiteUserProvidedFunctionAppResponseOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteUserProvidedFunctionAppResponse)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) ToStaticSiteUserProvidedFunctionAppResponseOutput() StaticSiteUserProvidedFunctionAppResponseOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) ToStaticSiteUserProvidedFunctionAppResponseOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppResponseOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) CreatedOn() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.CreatedOn }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) FunctionAppRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) *string { return v.FunctionAppRegion }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) FunctionAppResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) *string { return v.FunctionAppResourceId }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o StaticSiteUserProvidedFunctionAppResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v StaticSiteUserProvidedFunctionAppResponse) string { return v.Type }).(pulumi.StringOutput)
}

type StaticSiteUserProvidedFunctionAppResponseArrayOutput struct{ *pulumi.OutputState }

func (StaticSiteUserProvidedFunctionAppResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSiteUserProvidedFunctionAppResponse)(nil)).Elem()
}

func (o StaticSiteUserProvidedFunctionAppResponseArrayOutput) ToStaticSiteUserProvidedFunctionAppResponseArrayOutput() StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseArrayOutput) ToStaticSiteUserProvidedFunctionAppResponseArrayOutputWithContext(ctx context.Context) StaticSiteUserProvidedFunctionAppResponseArrayOutput {
	return o
}

func (o StaticSiteUserProvidedFunctionAppResponseArrayOutput) Index(i pulumi.IntInput) StaticSiteUserProvidedFunctionAppResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StaticSiteUserProvidedFunctionAppResponse {
		return vs[0].([]StaticSiteUserProvidedFunctionAppResponse)[vs[1].(int)]
	}).(StaticSiteUserProvidedFunctionAppResponseOutput)
}

type StatusCodesBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}





type StatusCodesBasedTriggerInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput
	ToStatusCodesBasedTriggerOutputWithContext(context.Context) StatusCodesBasedTriggerOutput
}

type StatusCodesBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Path         pulumi.StringPtrInput `pulumi:"path"`
	Status       pulumi.IntPtrInput    `pulumi:"status"`
	SubStatus    pulumi.IntPtrInput    `pulumi:"subStatus"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	Win32Status  pulumi.IntPtrInput    `pulumi:"win32Status"`
}

func (StatusCodesBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTrigger)(nil)).Elem()
}

func (i StatusCodesBasedTriggerArgs) ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput {
	return i.ToStatusCodesBasedTriggerOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerArgs) ToStatusCodesBasedTriggerOutputWithContext(ctx context.Context) StatusCodesBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerOutput)
}





type StatusCodesBasedTriggerArrayInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput
	ToStatusCodesBasedTriggerArrayOutputWithContext(context.Context) StatusCodesBasedTriggerArrayOutput
}

type StatusCodesBasedTriggerArray []StatusCodesBasedTriggerInput

func (StatusCodesBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTrigger)(nil)).Elem()
}

func (i StatusCodesBasedTriggerArray) ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput {
	return i.ToStatusCodesBasedTriggerArrayOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerArray) ToStatusCodesBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerArrayOutput)
}

type StatusCodesBasedTriggerOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTrigger)(nil)).Elem()
}

func (o StatusCodesBasedTriggerOutput) ToStatusCodesBasedTriggerOutput() StatusCodesBasedTriggerOutput {
	return o
}

func (o StatusCodesBasedTriggerOutput) ToStatusCodesBasedTriggerOutputWithContext(ctx context.Context) StatusCodesBasedTriggerOutput {
	return o
}

func (o StatusCodesBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) SubStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.SubStatus }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerOutput) Win32Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTrigger) *int { return v.Win32Status }).(pulumi.IntPtrOutput)
}

type StatusCodesBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTrigger)(nil)).Elem()
}

func (o StatusCodesBasedTriggerArrayOutput) ToStatusCodesBasedTriggerArrayOutput() StatusCodesBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerArrayOutput) ToStatusCodesBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerArrayOutput) Index(i pulumi.IntInput) StatusCodesBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesBasedTrigger {
		return vs[0].([]StatusCodesBasedTrigger)[vs[1].(int)]
	}).(StatusCodesBasedTriggerOutput)
}

type StatusCodesBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}

type StatusCodesBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesBasedTriggerResponseOutput) ToStatusCodesBasedTriggerResponseOutput() StatusCodesBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseOutput) ToStatusCodesBasedTriggerResponseOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Status }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) SubStatus() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.SubStatus }).(pulumi.IntPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

func (o StatusCodesBasedTriggerResponseOutput) Win32Status() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesBasedTriggerResponse) *int { return v.Win32Status }).(pulumi.IntPtrOutput)
}

type StatusCodesBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesBasedTriggerResponseArrayOutput) ToStatusCodesBasedTriggerResponseArrayOutput() StatusCodesBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseArrayOutput) ToStatusCodesBasedTriggerResponseArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) StatusCodesBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesBasedTriggerResponse {
		return vs[0].([]StatusCodesBasedTriggerResponse)[vs[1].(int)]
	}).(StatusCodesBasedTriggerResponseOutput)
}

type StatusCodesRangeBasedTrigger struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	StatusCodes  *string `pulumi:"statusCodes"`
	TimeInterval *string `pulumi:"timeInterval"`
}





type StatusCodesRangeBasedTriggerInput interface {
	pulumi.Input

	ToStatusCodesRangeBasedTriggerOutput() StatusCodesRangeBasedTriggerOutput
	ToStatusCodesRangeBasedTriggerOutputWithContext(context.Context) StatusCodesRangeBasedTriggerOutput
}

type StatusCodesRangeBasedTriggerArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Path         pulumi.StringPtrInput `pulumi:"path"`
	StatusCodes  pulumi.StringPtrInput `pulumi:"statusCodes"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
}

func (StatusCodesRangeBasedTriggerArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (i StatusCodesRangeBasedTriggerArgs) ToStatusCodesRangeBasedTriggerOutput() StatusCodesRangeBasedTriggerOutput {
	return i.ToStatusCodesRangeBasedTriggerOutputWithContext(context.Background())
}

func (i StatusCodesRangeBasedTriggerArgs) ToStatusCodesRangeBasedTriggerOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesRangeBasedTriggerOutput)
}





type StatusCodesRangeBasedTriggerArrayInput interface {
	pulumi.Input

	ToStatusCodesRangeBasedTriggerArrayOutput() StatusCodesRangeBasedTriggerArrayOutput
	ToStatusCodesRangeBasedTriggerArrayOutputWithContext(context.Context) StatusCodesRangeBasedTriggerArrayOutput
}

type StatusCodesRangeBasedTriggerArray []StatusCodesRangeBasedTriggerInput

func (StatusCodesRangeBasedTriggerArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (i StatusCodesRangeBasedTriggerArray) ToStatusCodesRangeBasedTriggerArrayOutput() StatusCodesRangeBasedTriggerArrayOutput {
	return i.ToStatusCodesRangeBasedTriggerArrayOutputWithContext(context.Background())
}

func (i StatusCodesRangeBasedTriggerArray) ToStatusCodesRangeBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesRangeBasedTriggerArrayOutput)
}

type StatusCodesRangeBasedTriggerOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerOutput) ToStatusCodesRangeBasedTriggerOutput() StatusCodesRangeBasedTriggerOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerOutput) ToStatusCodesRangeBasedTriggerOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesRangeBasedTriggerOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerOutput) StatusCodes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *string { return v.StatusCodes }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTrigger) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type StatusCodesRangeBasedTriggerArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesRangeBasedTrigger)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerArrayOutput) ToStatusCodesRangeBasedTriggerArrayOutput() StatusCodesRangeBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerArrayOutput) ToStatusCodesRangeBasedTriggerArrayOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerArrayOutput) Index(i pulumi.IntInput) StatusCodesRangeBasedTriggerOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesRangeBasedTrigger {
		return vs[0].([]StatusCodesRangeBasedTrigger)[vs[1].(int)]
	}).(StatusCodesRangeBasedTriggerOutput)
}

type StatusCodesRangeBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	Path         *string `pulumi:"path"`
	StatusCodes  *string `pulumi:"statusCodes"`
	TimeInterval *string `pulumi:"timeInterval"`
}

type StatusCodesRangeBasedTriggerResponseOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesRangeBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerResponseOutput) ToStatusCodesRangeBasedTriggerResponseOutput() StatusCodesRangeBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseOutput) ToStatusCodesRangeBasedTriggerResponseOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerResponseOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
}

func (o StatusCodesRangeBasedTriggerResponseOutput) Path() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *string { return v.Path }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerResponseOutput) StatusCodes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *string { return v.StatusCodes }).(pulumi.StringPtrOutput)
}

func (o StatusCodesRangeBasedTriggerResponseOutput) TimeInterval() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StatusCodesRangeBasedTriggerResponse) *string { return v.TimeInterval }).(pulumi.StringPtrOutput)
}

type StatusCodesRangeBasedTriggerResponseArrayOutput struct{ *pulumi.OutputState }

func (StatusCodesRangeBasedTriggerResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesRangeBasedTriggerResponse)(nil)).Elem()
}

func (o StatusCodesRangeBasedTriggerResponseArrayOutput) ToStatusCodesRangeBasedTriggerResponseArrayOutput() StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseArrayOutput) ToStatusCodesRangeBasedTriggerResponseArrayOutputWithContext(ctx context.Context) StatusCodesRangeBasedTriggerResponseArrayOutput {
	return o
}

func (o StatusCodesRangeBasedTriggerResponseArrayOutput) Index(i pulumi.IntInput) StatusCodesRangeBasedTriggerResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StatusCodesRangeBasedTriggerResponse {
		return vs[0].([]StatusCodesRangeBasedTriggerResponse)[vs[1].(int)]
	}).(StatusCodesRangeBasedTriggerResponseOutput)
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


func (val *TrafficWeightArgs) Defaults() *TrafficWeightArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.LatestRevision) {
		tmp.LatestRevision = pulumi.BoolPtr(false)
	}
	return &tmp
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

type UserAssignedIdentityResponse struct {
	ClientId    string `pulumi:"clientId"`
	PrincipalId string `pulumi:"principalId"`
}

type UserAssignedIdentityResponseOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutput() UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ToUserAssignedIdentityResponseOutputWithContext(ctx context.Context) UserAssignedIdentityResponseOutput {
	return o
}

func (o UserAssignedIdentityResponseOutput) ClientId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.ClientId }).(pulumi.StringOutput)
}

func (o UserAssignedIdentityResponseOutput) PrincipalId() pulumi.StringOutput {
	return o.ApplyT(func(v UserAssignedIdentityResponse) string { return v.PrincipalId }).(pulumi.StringOutput)
}

type UserAssignedIdentityResponseMapOutput struct{ *pulumi.OutputState }

func (UserAssignedIdentityResponseMapOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]UserAssignedIdentityResponse)(nil)).Elem()
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutput() UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) ToUserAssignedIdentityResponseMapOutputWithContext(ctx context.Context) UserAssignedIdentityResponseMapOutput {
	return o
}

func (o UserAssignedIdentityResponseMapOutput) MapIndex(k pulumi.StringInput) UserAssignedIdentityResponseOutput {
	return pulumi.All(o, k).ApplyT(func(vs []interface{}) UserAssignedIdentityResponse {
		return vs[0].(map[string]UserAssignedIdentityResponse)[vs[1].(string)]
	}).(UserAssignedIdentityResponseOutput)
}

type VirtualApplication struct {
	PhysicalPath       *string            `pulumi:"physicalPath"`
	PreloadEnabled     *bool              `pulumi:"preloadEnabled"`
	VirtualDirectories []VirtualDirectory `pulumi:"virtualDirectories"`
	VirtualPath        *string            `pulumi:"virtualPath"`
}





type VirtualApplicationInput interface {
	pulumi.Input

	ToVirtualApplicationOutput() VirtualApplicationOutput
	ToVirtualApplicationOutputWithContext(context.Context) VirtualApplicationOutput
}

type VirtualApplicationArgs struct {
	PhysicalPath       pulumi.StringPtrInput      `pulumi:"physicalPath"`
	PreloadEnabled     pulumi.BoolPtrInput        `pulumi:"preloadEnabled"`
	VirtualDirectories VirtualDirectoryArrayInput `pulumi:"virtualDirectories"`
	VirtualPath        pulumi.StringPtrInput      `pulumi:"virtualPath"`
}

func (VirtualApplicationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplication)(nil)).Elem()
}

func (i VirtualApplicationArgs) ToVirtualApplicationOutput() VirtualApplicationOutput {
	return i.ToVirtualApplicationOutputWithContext(context.Background())
}

func (i VirtualApplicationArgs) ToVirtualApplicationOutputWithContext(ctx context.Context) VirtualApplicationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationOutput)
}





type VirtualApplicationArrayInput interface {
	pulumi.Input

	ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput
	ToVirtualApplicationArrayOutputWithContext(context.Context) VirtualApplicationArrayOutput
}

type VirtualApplicationArray []VirtualApplicationInput

func (VirtualApplicationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplication)(nil)).Elem()
}

func (i VirtualApplicationArray) ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput {
	return i.ToVirtualApplicationArrayOutputWithContext(context.Background())
}

func (i VirtualApplicationArray) ToVirtualApplicationArrayOutputWithContext(ctx context.Context) VirtualApplicationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationArrayOutput)
}

type VirtualApplicationOutput struct{ *pulumi.OutputState }

func (VirtualApplicationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplication)(nil)).Elem()
}

func (o VirtualApplicationOutput) ToVirtualApplicationOutput() VirtualApplicationOutput {
	return o
}

func (o VirtualApplicationOutput) ToVirtualApplicationOutputWithContext(ctx context.Context) VirtualApplicationOutput {
	return o
}

func (o VirtualApplicationOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualApplicationOutput) PreloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *bool { return v.PreloadEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualApplicationOutput) VirtualDirectories() VirtualDirectoryArrayOutput {
	return o.ApplyT(func(v VirtualApplication) []VirtualDirectory { return v.VirtualDirectories }).(VirtualDirectoryArrayOutput)
}

func (o VirtualApplicationOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplication) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualApplicationArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplicationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplication)(nil)).Elem()
}

func (o VirtualApplicationArrayOutput) ToVirtualApplicationArrayOutput() VirtualApplicationArrayOutput {
	return o
}

func (o VirtualApplicationArrayOutput) ToVirtualApplicationArrayOutputWithContext(ctx context.Context) VirtualApplicationArrayOutput {
	return o
}

func (o VirtualApplicationArrayOutput) Index(i pulumi.IntInput) VirtualApplicationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplication {
		return vs[0].([]VirtualApplication)[vs[1].(int)]
	}).(VirtualApplicationOutput)
}

type VirtualApplicationResponse struct {
	PhysicalPath       *string                    `pulumi:"physicalPath"`
	PreloadEnabled     *bool                      `pulumi:"preloadEnabled"`
	VirtualDirectories []VirtualDirectoryResponse `pulumi:"virtualDirectories"`
	VirtualPath        *string                    `pulumi:"virtualPath"`
}

type VirtualApplicationResponseOutput struct{ *pulumi.OutputState }

func (VirtualApplicationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplicationResponse)(nil)).Elem()
}

func (o VirtualApplicationResponseOutput) ToVirtualApplicationResponseOutput() VirtualApplicationResponseOutput {
	return o
}

func (o VirtualApplicationResponseOutput) ToVirtualApplicationResponseOutputWithContext(ctx context.Context) VirtualApplicationResponseOutput {
	return o
}

func (o VirtualApplicationResponseOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualApplicationResponseOutput) PreloadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *bool { return v.PreloadEnabled }).(pulumi.BoolPtrOutput)
}

func (o VirtualApplicationResponseOutput) VirtualDirectories() VirtualDirectoryResponseArrayOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) []VirtualDirectoryResponse { return v.VirtualDirectories }).(VirtualDirectoryResponseArrayOutput)
}

func (o VirtualApplicationResponseOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualApplicationResponse) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualApplicationResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualApplicationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplicationResponse)(nil)).Elem()
}

func (o VirtualApplicationResponseArrayOutput) ToVirtualApplicationResponseArrayOutput() VirtualApplicationResponseArrayOutput {
	return o
}

func (o VirtualApplicationResponseArrayOutput) ToVirtualApplicationResponseArrayOutputWithContext(ctx context.Context) VirtualApplicationResponseArrayOutput {
	return o
}

func (o VirtualApplicationResponseArrayOutput) Index(i pulumi.IntInput) VirtualApplicationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualApplicationResponse {
		return vs[0].([]VirtualApplicationResponse)[vs[1].(int)]
	}).(VirtualApplicationResponseOutput)
}

type VirtualDirectory struct {
	PhysicalPath *string `pulumi:"physicalPath"`
	VirtualPath  *string `pulumi:"virtualPath"`
}





type VirtualDirectoryInput interface {
	pulumi.Input

	ToVirtualDirectoryOutput() VirtualDirectoryOutput
	ToVirtualDirectoryOutputWithContext(context.Context) VirtualDirectoryOutput
}

type VirtualDirectoryArgs struct {
	PhysicalPath pulumi.StringPtrInput `pulumi:"physicalPath"`
	VirtualPath  pulumi.StringPtrInput `pulumi:"virtualPath"`
}

func (VirtualDirectoryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectory)(nil)).Elem()
}

func (i VirtualDirectoryArgs) ToVirtualDirectoryOutput() VirtualDirectoryOutput {
	return i.ToVirtualDirectoryOutputWithContext(context.Background())
}

func (i VirtualDirectoryArgs) ToVirtualDirectoryOutputWithContext(ctx context.Context) VirtualDirectoryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryOutput)
}





type VirtualDirectoryArrayInput interface {
	pulumi.Input

	ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput
	ToVirtualDirectoryArrayOutputWithContext(context.Context) VirtualDirectoryArrayOutput
}

type VirtualDirectoryArray []VirtualDirectoryInput

func (VirtualDirectoryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectory)(nil)).Elem()
}

func (i VirtualDirectoryArray) ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput {
	return i.ToVirtualDirectoryArrayOutputWithContext(context.Background())
}

func (i VirtualDirectoryArray) ToVirtualDirectoryArrayOutputWithContext(ctx context.Context) VirtualDirectoryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryArrayOutput)
}

type VirtualDirectoryOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectory)(nil)).Elem()
}

func (o VirtualDirectoryOutput) ToVirtualDirectoryOutput() VirtualDirectoryOutput {
	return o
}

func (o VirtualDirectoryOutput) ToVirtualDirectoryOutputWithContext(ctx context.Context) VirtualDirectoryOutput {
	return o
}

func (o VirtualDirectoryOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectory) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualDirectoryOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectory) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualDirectoryArrayOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectory)(nil)).Elem()
}

func (o VirtualDirectoryArrayOutput) ToVirtualDirectoryArrayOutput() VirtualDirectoryArrayOutput {
	return o
}

func (o VirtualDirectoryArrayOutput) ToVirtualDirectoryArrayOutputWithContext(ctx context.Context) VirtualDirectoryArrayOutput {
	return o
}

func (o VirtualDirectoryArrayOutput) Index(i pulumi.IntInput) VirtualDirectoryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDirectory {
		return vs[0].([]VirtualDirectory)[vs[1].(int)]
	}).(VirtualDirectoryOutput)
}

type VirtualDirectoryResponse struct {
	PhysicalPath *string `pulumi:"physicalPath"`
	VirtualPath  *string `pulumi:"virtualPath"`
}

type VirtualDirectoryResponseOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectoryResponse)(nil)).Elem()
}

func (o VirtualDirectoryResponseOutput) ToVirtualDirectoryResponseOutput() VirtualDirectoryResponseOutput {
	return o
}

func (o VirtualDirectoryResponseOutput) ToVirtualDirectoryResponseOutputWithContext(ctx context.Context) VirtualDirectoryResponseOutput {
	return o
}

func (o VirtualDirectoryResponseOutput) PhysicalPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectoryResponse) *string { return v.PhysicalPath }).(pulumi.StringPtrOutput)
}

func (o VirtualDirectoryResponseOutput) VirtualPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualDirectoryResponse) *string { return v.VirtualPath }).(pulumi.StringPtrOutput)
}

type VirtualDirectoryResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualDirectoryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectoryResponse)(nil)).Elem()
}

func (o VirtualDirectoryResponseArrayOutput) ToVirtualDirectoryResponseArrayOutput() VirtualDirectoryResponseArrayOutput {
	return o
}

func (o VirtualDirectoryResponseArrayOutput) ToVirtualDirectoryResponseArrayOutputWithContext(ctx context.Context) VirtualDirectoryResponseArrayOutput {
	return o
}

func (o VirtualDirectoryResponseArrayOutput) Index(i pulumi.IntInput) VirtualDirectoryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualDirectoryResponse {
		return vs[0].([]VirtualDirectoryResponse)[vs[1].(int)]
	}).(VirtualDirectoryResponseOutput)
}

type VirtualNetworkProfile struct {
	Id     string  `pulumi:"id"`
	Subnet *string `pulumi:"subnet"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	Id     pulumi.StringInput    `pulumi:"id"`
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
}

func (VirtualNetworkProfileArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return i.ToVirtualNetworkProfileOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput)
}

type VirtualNetworkProfileOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfileOutputWithContext(ctx context.Context) VirtualNetworkProfileOutput {
	return o
}

func (o VirtualNetworkProfileOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	Id     string  `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Subnet *string `pulumi:"subnet"`
	Type   string  `pulumi:"type"`
}

type VirtualNetworkProfileResponseOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return o
}

func (o VirtualNetworkProfileResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VirtualNetworkProfileResponseOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VnetRouteResponse struct {
	EndAddress   *string `pulumi:"endAddress"`
	Id           string  `pulumi:"id"`
	Kind         *string `pulumi:"kind"`
	Name         string  `pulumi:"name"`
	RouteType    *string `pulumi:"routeType"`
	StartAddress *string `pulumi:"startAddress"`
	Type         string  `pulumi:"type"`
}

type VnetRouteResponseOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutput() VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) ToVnetRouteResponseOutputWithContext(ctx context.Context) VnetRouteResponseOutput {
	return o
}

func (o VnetRouteResponseOutput) EndAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.EndAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o VnetRouteResponseOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o VnetRouteResponseOutput) RouteType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.RouteType }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) StartAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VnetRouteResponse) *string { return v.StartAddress }).(pulumi.StringPtrOutput)
}

func (o VnetRouteResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v VnetRouteResponse) string { return v.Type }).(pulumi.StringOutput)
}

type VnetRouteResponseArrayOutput struct{ *pulumi.OutputState }

func (VnetRouteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnetRouteResponse)(nil)).Elem()
}

func (o VnetRouteResponseArrayOutput) ToVnetRouteResponseArrayOutput() VnetRouteResponseArrayOutput {
	return o
}

func (o VnetRouteResponseArrayOutput) ToVnetRouteResponseArrayOutputWithContext(ctx context.Context) VnetRouteResponseArrayOutput {
	return o
}

func (o VnetRouteResponseArrayOutput) Index(i pulumi.IntInput) VnetRouteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VnetRouteResponse {
		return vs[0].([]VnetRouteResponse)[vs[1].(int)]
	}).(VnetRouteResponseOutput)
}

type WorkflowTriggerListCallbackUrlQueriesResponse struct {
	ApiVersion *string `pulumi:"apiVersion"`
	Se         *string `pulumi:"se"`
	Sig        *string `pulumi:"sig"`
	Sp         *string `pulumi:"sp"`
	Sv         *string `pulumi:"sv"`
}

type WorkflowTriggerListCallbackUrlQueriesResponseOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerListCallbackUrlQueriesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkflowTriggerListCallbackUrlQueriesResponse)(nil)).Elem()
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ToWorkflowTriggerListCallbackUrlQueriesResponseOutput() WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ToWorkflowTriggerListCallbackUrlQueriesResponseOutputWithContext(ctx context.Context) WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.ApiVersion }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Se() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Se }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sig }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sp }).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponseOutput) Sv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkflowTriggerListCallbackUrlQueriesResponse) *string { return v.Sv }).(pulumi.StringPtrOutput)
}

type WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput struct{ *pulumi.OutputState }

func (WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WorkflowTriggerListCallbackUrlQueriesResponse)(nil)).Elem()
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ToWorkflowTriggerListCallbackUrlQueriesResponsePtrOutput() WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ToWorkflowTriggerListCallbackUrlQueriesResponsePtrOutputWithContext(ctx context.Context) WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput {
	return o
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Elem() WorkflowTriggerListCallbackUrlQueriesResponseOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) WorkflowTriggerListCallbackUrlQueriesResponse {
		if v != nil {
			return *v
		}
		var ret WorkflowTriggerListCallbackUrlQueriesResponse
		return ret
	}).(WorkflowTriggerListCallbackUrlQueriesResponseOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) ApiVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiVersion
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Se() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Se
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sig() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sig
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sp() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sp
	}).(pulumi.StringPtrOutput)
}

func (o WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput) Sv() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WorkflowTriggerListCallbackUrlQueriesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Sv
	}).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ApiDefinitionInfoOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoPtrOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoResponseOutput{})
	pulumi.RegisterOutputType(ApiDefinitionInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigPtrOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigResponseOutput{})
	pulumi.RegisterOutputType(ApiManagementConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AppLogsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ArcConfigurationOutput{})
	pulumi.RegisterOutputType(ArcConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ArcConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ArcConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponseOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmPlanResponseOutput{})
	pulumi.RegisterOutputType(ArmPlanResponsePtrOutput{})
	pulumi.RegisterOutputType(AseV3NetworkingConfigurationOutput{})
	pulumi.RegisterOutputType(AseV3NetworkingConfigurationPtrOutput{})
	pulumi.RegisterOutputType(AseV3NetworkingConfigurationResponseOutput{})
	pulumi.RegisterOutputType(AseV3NetworkingConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionsOutput{})
	pulumi.RegisterOutputType(AutoHealActionsPtrOutput{})
	pulumi.RegisterOutputType(AutoHealActionsResponseOutput{})
	pulumi.RegisterOutputType(AutoHealActionsResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionPtrOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionResponseOutput{})
	pulumi.RegisterOutputType(AutoHealCustomActionResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealRulesOutput{})
	pulumi.RegisterOutputType(AutoHealRulesPtrOutput{})
	pulumi.RegisterOutputType(AutoHealRulesResponseOutput{})
	pulumi.RegisterOutputType(AutoHealRulesResponsePtrOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersPtrOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersResponseOutput{})
	pulumi.RegisterOutputType(AutoHealTriggersResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureBlobStorageHttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponseOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponsePtrOutput{})
	pulumi.RegisterOutputType(AzureResourceErrorInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueMapOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueResponseOutput{})
	pulumi.RegisterOutputType(AzureStorageInfoValueResponseMapOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(AzureTableStorageApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(BackupItemResponseOutput{})
	pulumi.RegisterOutputType(BackupItemResponseArrayOutput{})
	pulumi.RegisterOutputType(BackupScheduleOutput{})
	pulumi.RegisterOutputType(BackupSchedulePtrOutput{})
	pulumi.RegisterOutputType(BackupScheduleResponseOutput{})
	pulumi.RegisterOutputType(BackupScheduleResponsePtrOutput{})
	pulumi.RegisterOutputType(CapabilityOutput{})
	pulumi.RegisterOutputType(CapabilityArrayOutput{})
	pulumi.RegisterOutputType(CapabilityResponseOutput{})
	pulumi.RegisterOutputType(CapabilityResponseArrayOutput{})
	pulumi.RegisterOutputType(CloningInfoOutput{})
	pulumi.RegisterOutputType(CloningInfoPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationOutput{})
	pulumi.RegisterOutputType(ConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnStringInfoOutput{})
	pulumi.RegisterOutputType(ConnStringInfoArrayOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairMapOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseMapOutput{})
	pulumi.RegisterOutputType(ContainerOutput{})
	pulumi.RegisterOutputType(ContainerArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppSecretResponseOutput{})
	pulumi.RegisterOutputType(ContainerAppSecretResponseArrayOutput{})
	pulumi.RegisterOutputType(ContainerAppsConfigurationOutput{})
	pulumi.RegisterOutputType(ContainerAppsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(ContainerAppsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(ContainerAppsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResourcesOutput{})
	pulumi.RegisterOutputType(ContainerResourcesPtrOutput{})
	pulumi.RegisterOutputType(ContainerResourcesResponseOutput{})
	pulumi.RegisterOutputType(ContainerResourcesResponsePtrOutput{})
	pulumi.RegisterOutputType(ContainerResponseOutput{})
	pulumi.RegisterOutputType(ContainerResponseArrayOutput{})
	pulumi.RegisterOutputType(CorsSettingsOutput{})
	pulumi.RegisterOutputType(CorsSettingsPtrOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponseOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(CustomDnsSuffixConfigurationOutput{})
	pulumi.RegisterOutputType(CustomDnsSuffixConfigurationPtrOutput{})
	pulumi.RegisterOutputType(CustomDnsSuffixConfigurationResponseOutput{})
	pulumi.RegisterOutputType(CustomDnsSuffixConfigurationResponsePtrOutput{})
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
	pulumi.RegisterOutputType(DatabaseBackupSettingOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingArrayOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(EnabledConfigOutput{})
	pulumi.RegisterOutputType(EnabledConfigPtrOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponseOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(EnvironmentVarOutput{})
	pulumi.RegisterOutputType(EnvironmentVarArrayOutput{})
	pulumi.RegisterOutputType(EnvironmentVarResponseOutput{})
	pulumi.RegisterOutputType(EnvironmentVarResponseArrayOutput{})
	pulumi.RegisterOutputType(ErrorEntityResponseOutput{})
	pulumi.RegisterOutputType(ErrorEntityResponsePtrOutput{})
	pulumi.RegisterOutputType(ErrorEntityResponseArrayOutput{})
	pulumi.RegisterOutputType(ExperimentsOutput{})
	pulumi.RegisterOutputType(ExperimentsPtrOutput{})
	pulumi.RegisterOutputType(ExperimentsResponseOutput{})
	pulumi.RegisterOutputType(ExperimentsResponsePtrOutput{})
	pulumi.RegisterOutputType(ExpressionResponseOutput{})
	pulumi.RegisterOutputType(ExpressionResponseArrayOutput{})
	pulumi.RegisterOutputType(ExpressionRootResponseOutput{})
	pulumi.RegisterOutputType(ExpressionRootResponseArrayOutput{})
	pulumi.RegisterOutputType(ExtendedLocationOutput{})
	pulumi.RegisterOutputType(ExtendedLocationPtrOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponseOutput{})
	pulumi.RegisterOutputType(ExtendedLocationResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationPtrOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationResponseOutput{})
	pulumi.RegisterOutputType(FrontEndConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GitHubActionCodeConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GitHubActionConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationPtrOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationResponseOutput{})
	pulumi.RegisterOutputType(GitHubActionContainerConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(HandlerMappingOutput{})
	pulumi.RegisterOutputType(HandlerMappingArrayOutput{})
	pulumi.RegisterOutputType(HandlerMappingResponseOutput{})
	pulumi.RegisterOutputType(HandlerMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(HostNameSslStateOutput{})
	pulumi.RegisterOutputType(HostNameSslStateArrayOutput{})
	pulumi.RegisterOutputType(HostNameSslStateResponseOutput{})
	pulumi.RegisterOutputType(HostNameSslStateResponseArrayOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfilePtrOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(HostingEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(HttpLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleOutput{})
	pulumi.RegisterOutputType(HttpScaleRulePtrOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(HttpScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(IdentifierResponseOutput{})
	pulumi.RegisterOutputType(IdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(IngressOutput{})
	pulumi.RegisterOutputType(IngressPtrOutput{})
	pulumi.RegisterOutputType(IngressResponseOutput{})
	pulumi.RegisterOutputType(IngressResponsePtrOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionArrayOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseArrayOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfileOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfilePtrOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfileResponseOutput{})
	pulumi.RegisterOutputType(KubeEnvironmentProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationPtrOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponseOutput{})
	pulumi.RegisterOutputType(LogAnalyticsConfigurationResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(NameValuePairOutput{})
	pulumi.RegisterOutputType(NameValuePairArrayOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PushSettingsOutput{})
	pulumi.RegisterOutputType(PushSettingsPtrOutput{})
	pulumi.RegisterOutputType(PushSettingsResponseOutput{})
	pulumi.RegisterOutputType(PushSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleOutput{})
	pulumi.RegisterOutputType(QueueScaleRulePtrOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleResponseOutput{})
	pulumi.RegisterOutputType(QueueScaleRuleResponsePtrOutput{})
	pulumi.RegisterOutputType(RampUpRuleOutput{})
	pulumi.RegisterOutputType(RampUpRuleArrayOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsArrayOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsResponseOutput{})
	pulumi.RegisterOutputType(RegistryCredentialsResponseArrayOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(RemotePrivateEndpointConnectionResponsePtrOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseOutput{})
	pulumi.RegisterOutputType(ResponseMessageEnvelopeRemotePrivateEndpointConnectionResponseArrayOutput{})
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
	pulumi.RegisterOutputType(SiteConfigOutput{})
	pulumi.RegisterOutputType(SiteConfigPtrOutput{})
	pulumi.RegisterOutputType(SiteConfigResponseOutput{})
	pulumi.RegisterOutputType(SiteConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteLimitsOutput{})
	pulumi.RegisterOutputType(SiteLimitsPtrOutput{})
	pulumi.RegisterOutputType(SiteLimitsResponseOutput{})
	pulumi.RegisterOutputType(SiteLimitsResponsePtrOutput{})
	pulumi.RegisterOutputType(SiteMachineKeyResponseOutput{})
	pulumi.RegisterOutputType(SiteMachineKeyResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuCapacityOutput{})
	pulumi.RegisterOutputType(SkuCapacityPtrOutput{})
	pulumi.RegisterOutputType(SkuCapacityResponseOutput{})
	pulumi.RegisterOutputType(SkuCapacityResponsePtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionOutput{})
	pulumi.RegisterOutputType(SkuDescriptionPtrOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponseOutput{})
	pulumi.RegisterOutputType(SkuDescriptionResponsePtrOutput{})
	pulumi.RegisterOutputType(SlotSwapStatusResponseOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StaticSiteLinkedBackendResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteLinkedBackendResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsPtrOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteTemplateOptionsResponsePtrOutput{})
	pulumi.RegisterOutputType(StaticSiteUserARMResourceResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteUserARMResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteUserProvidedFunctionAppResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(StatusCodesRangeBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(TemplateOutput{})
	pulumi.RegisterOutputType(TemplatePtrOutput{})
	pulumi.RegisterOutputType(TemplateResponseOutput{})
	pulumi.RegisterOutputType(TemplateResponsePtrOutput{})
	pulumi.RegisterOutputType(TrafficWeightOutput{})
	pulumi.RegisterOutputType(TrafficWeightArrayOutput{})
	pulumi.RegisterOutputType(TrafficWeightResponseOutput{})
	pulumi.RegisterOutputType(TrafficWeightResponseArrayOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseOutput{})
	pulumi.RegisterOutputType(UserAssignedIdentityResponseMapOutput{})
	pulumi.RegisterOutputType(VirtualApplicationOutput{})
	pulumi.RegisterOutputType(VirtualApplicationArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerListCallbackUrlQueriesResponseOutput{})
	pulumi.RegisterOutputType(WorkflowTriggerListCallbackUrlQueriesResponsePtrOutput{})
}
