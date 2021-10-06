


package v20190801

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





type ApiDefinitionInfoResponseInput interface {
	pulumi.Input

	ToApiDefinitionInfoResponseOutput() ApiDefinitionInfoResponseOutput
	ToApiDefinitionInfoResponseOutputWithContext(context.Context) ApiDefinitionInfoResponseOutput
}

type ApiDefinitionInfoResponseArgs struct {
	Url pulumi.StringPtrInput `pulumi:"url"`
}

func (ApiDefinitionInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiDefinitionInfoResponse)(nil)).Elem()
}

func (i ApiDefinitionInfoResponseArgs) ToApiDefinitionInfoResponseOutput() ApiDefinitionInfoResponseOutput {
	return i.ToApiDefinitionInfoResponseOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoResponseArgs) ToApiDefinitionInfoResponseOutputWithContext(ctx context.Context) ApiDefinitionInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoResponseOutput)
}

func (i ApiDefinitionInfoResponseArgs) ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput {
	return i.ToApiDefinitionInfoResponsePtrOutputWithContext(context.Background())
}

func (i ApiDefinitionInfoResponseArgs) ToApiDefinitionInfoResponsePtrOutputWithContext(ctx context.Context) ApiDefinitionInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoResponseOutput).ToApiDefinitionInfoResponsePtrOutputWithContext(ctx)
}









type ApiDefinitionInfoResponsePtrInput interface {
	pulumi.Input

	ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput
	ToApiDefinitionInfoResponsePtrOutputWithContext(context.Context) ApiDefinitionInfoResponsePtrOutput
}

type apiDefinitionInfoResponsePtrType ApiDefinitionInfoResponseArgs

func ApiDefinitionInfoResponsePtr(v *ApiDefinitionInfoResponseArgs) ApiDefinitionInfoResponsePtrInput {
	return (*apiDefinitionInfoResponsePtrType)(v)
}

func (*apiDefinitionInfoResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiDefinitionInfoResponse)(nil)).Elem()
}

func (i *apiDefinitionInfoResponsePtrType) ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput {
	return i.ToApiDefinitionInfoResponsePtrOutputWithContext(context.Background())
}

func (i *apiDefinitionInfoResponsePtrType) ToApiDefinitionInfoResponsePtrOutputWithContext(ctx context.Context) ApiDefinitionInfoResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiDefinitionInfoResponsePtrOutput)
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

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponsePtrOutput() ApiDefinitionInfoResponsePtrOutput {
	return o.ToApiDefinitionInfoResponsePtrOutputWithContext(context.Background())
}

func (o ApiDefinitionInfoResponseOutput) ToApiDefinitionInfoResponsePtrOutputWithContext(ctx context.Context) ApiDefinitionInfoResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiDefinitionInfoResponse) *ApiDefinitionInfoResponse {
		return &v
	}).(ApiDefinitionInfoResponsePtrOutput)
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





type ApiManagementConfigResponseInput interface {
	pulumi.Input

	ToApiManagementConfigResponseOutput() ApiManagementConfigResponseOutput
	ToApiManagementConfigResponseOutputWithContext(context.Context) ApiManagementConfigResponseOutput
}

type ApiManagementConfigResponseArgs struct {
	Id pulumi.StringPtrInput `pulumi:"id"`
}

func (ApiManagementConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApiManagementConfigResponse)(nil)).Elem()
}

func (i ApiManagementConfigResponseArgs) ToApiManagementConfigResponseOutput() ApiManagementConfigResponseOutput {
	return i.ToApiManagementConfigResponseOutputWithContext(context.Background())
}

func (i ApiManagementConfigResponseArgs) ToApiManagementConfigResponseOutputWithContext(ctx context.Context) ApiManagementConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigResponseOutput)
}

func (i ApiManagementConfigResponseArgs) ToApiManagementConfigResponsePtrOutput() ApiManagementConfigResponsePtrOutput {
	return i.ToApiManagementConfigResponsePtrOutputWithContext(context.Background())
}

func (i ApiManagementConfigResponseArgs) ToApiManagementConfigResponsePtrOutputWithContext(ctx context.Context) ApiManagementConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigResponseOutput).ToApiManagementConfigResponsePtrOutputWithContext(ctx)
}









type ApiManagementConfigResponsePtrInput interface {
	pulumi.Input

	ToApiManagementConfigResponsePtrOutput() ApiManagementConfigResponsePtrOutput
	ToApiManagementConfigResponsePtrOutputWithContext(context.Context) ApiManagementConfigResponsePtrOutput
}

type apiManagementConfigResponsePtrType ApiManagementConfigResponseArgs

func ApiManagementConfigResponsePtr(v *ApiManagementConfigResponseArgs) ApiManagementConfigResponsePtrInput {
	return (*apiManagementConfigResponsePtrType)(v)
}

func (*apiManagementConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApiManagementConfigResponse)(nil)).Elem()
}

func (i *apiManagementConfigResponsePtrType) ToApiManagementConfigResponsePtrOutput() ApiManagementConfigResponsePtrOutput {
	return i.ToApiManagementConfigResponsePtrOutputWithContext(context.Background())
}

func (i *apiManagementConfigResponsePtrType) ToApiManagementConfigResponsePtrOutputWithContext(ctx context.Context) ApiManagementConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApiManagementConfigResponsePtrOutput)
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

func (o ApiManagementConfigResponseOutput) ToApiManagementConfigResponsePtrOutput() ApiManagementConfigResponsePtrOutput {
	return o.ToApiManagementConfigResponsePtrOutputWithContext(context.Background())
}

func (o ApiManagementConfigResponseOutput) ToApiManagementConfigResponsePtrOutputWithContext(ctx context.Context) ApiManagementConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApiManagementConfigResponse) *ApiManagementConfigResponse {
		return &v
	}).(ApiManagementConfigResponsePtrOutput)
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

type ApplicationLogsConfig struct {
	AzureBlobStorage  *AzureBlobStorageApplicationLogsConfig  `pulumi:"azureBlobStorage"`
	AzureTableStorage *AzureTableStorageApplicationLogsConfig `pulumi:"azureTableStorage"`
	FileSystem        *FileSystemApplicationLogsConfig        `pulumi:"fileSystem"`
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





type ApplicationLogsConfigResponseInput interface {
	pulumi.Input

	ToApplicationLogsConfigResponseOutput() ApplicationLogsConfigResponseOutput
	ToApplicationLogsConfigResponseOutputWithContext(context.Context) ApplicationLogsConfigResponseOutput
}

type ApplicationLogsConfigResponseArgs struct {
	AzureBlobStorage  AzureBlobStorageApplicationLogsConfigResponsePtrInput  `pulumi:"azureBlobStorage"`
	AzureTableStorage AzureTableStorageApplicationLogsConfigResponsePtrInput `pulumi:"azureTableStorage"`
	FileSystem        FileSystemApplicationLogsConfigResponsePtrInput        `pulumi:"fileSystem"`
}

func (ApplicationLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ApplicationLogsConfigResponse)(nil)).Elem()
}

func (i ApplicationLogsConfigResponseArgs) ToApplicationLogsConfigResponseOutput() ApplicationLogsConfigResponseOutput {
	return i.ToApplicationLogsConfigResponseOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigResponseArgs) ToApplicationLogsConfigResponseOutputWithContext(ctx context.Context) ApplicationLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigResponseOutput)
}

func (i ApplicationLogsConfigResponseArgs) ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput {
	return i.ToApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i ApplicationLogsConfigResponseArgs) ToApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) ApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigResponseOutput).ToApplicationLogsConfigResponsePtrOutputWithContext(ctx)
}









type ApplicationLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput
	ToApplicationLogsConfigResponsePtrOutputWithContext(context.Context) ApplicationLogsConfigResponsePtrOutput
}

type applicationLogsConfigResponsePtrType ApplicationLogsConfigResponseArgs

func ApplicationLogsConfigResponsePtr(v *ApplicationLogsConfigResponseArgs) ApplicationLogsConfigResponsePtrInput {
	return (*applicationLogsConfigResponsePtrType)(v)
}

func (*applicationLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ApplicationLogsConfigResponse)(nil)).Elem()
}

func (i *applicationLogsConfigResponsePtrType) ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput {
	return i.ToApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *applicationLogsConfigResponsePtrType) ToApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) ApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ApplicationLogsConfigResponsePtrOutput)
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

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponsePtrOutput() ApplicationLogsConfigResponsePtrOutput {
	return o.ToApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o ApplicationLogsConfigResponseOutput) ToApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) ApplicationLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ApplicationLogsConfigResponse) *ApplicationLogsConfigResponse {
		return &v
	}).(ApplicationLogsConfigResponsePtrOutput)
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

type ArmIdWrapperResponse struct {
	Id string `pulumi:"id"`
}





type ArmIdWrapperResponseInput interface {
	pulumi.Input

	ToArmIdWrapperResponseOutput() ArmIdWrapperResponseOutput
	ToArmIdWrapperResponseOutputWithContext(context.Context) ArmIdWrapperResponseOutput
}

type ArmIdWrapperResponseArgs struct {
	Id pulumi.StringInput `pulumi:"id"`
}

func (ArmIdWrapperResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ArmIdWrapperResponse)(nil)).Elem()
}

func (i ArmIdWrapperResponseArgs) ToArmIdWrapperResponseOutput() ArmIdWrapperResponseOutput {
	return i.ToArmIdWrapperResponseOutputWithContext(context.Background())
}

func (i ArmIdWrapperResponseArgs) ToArmIdWrapperResponseOutputWithContext(ctx context.Context) ArmIdWrapperResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdWrapperResponseOutput)
}

func (i ArmIdWrapperResponseArgs) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return i.ToArmIdWrapperResponsePtrOutputWithContext(context.Background())
}

func (i ArmIdWrapperResponseArgs) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdWrapperResponseOutput).ToArmIdWrapperResponsePtrOutputWithContext(ctx)
}









type ArmIdWrapperResponsePtrInput interface {
	pulumi.Input

	ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput
	ToArmIdWrapperResponsePtrOutputWithContext(context.Context) ArmIdWrapperResponsePtrOutput
}

type armIdWrapperResponsePtrType ArmIdWrapperResponseArgs

func ArmIdWrapperResponsePtr(v *ArmIdWrapperResponseArgs) ArmIdWrapperResponsePtrInput {
	return (*armIdWrapperResponsePtrType)(v)
}

func (*armIdWrapperResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ArmIdWrapperResponse)(nil)).Elem()
}

func (i *armIdWrapperResponsePtrType) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return i.ToArmIdWrapperResponsePtrOutputWithContext(context.Background())
}

func (i *armIdWrapperResponsePtrType) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ArmIdWrapperResponsePtrOutput)
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

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponsePtrOutput() ArmIdWrapperResponsePtrOutput {
	return o.ToArmIdWrapperResponsePtrOutputWithContext(context.Background())
}

func (o ArmIdWrapperResponseOutput) ToArmIdWrapperResponsePtrOutputWithContext(ctx context.Context) ArmIdWrapperResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ArmIdWrapperResponse) *ArmIdWrapperResponse {
		return &v
	}).(ArmIdWrapperResponsePtrOutput)
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





type AutoHealActionsResponseInput interface {
	pulumi.Input

	ToAutoHealActionsResponseOutput() AutoHealActionsResponseOutput
	ToAutoHealActionsResponseOutputWithContext(context.Context) AutoHealActionsResponseOutput
}

type AutoHealActionsResponseArgs struct {
	ActionType              pulumi.StringPtrInput                `pulumi:"actionType"`
	CustomAction            AutoHealCustomActionResponsePtrInput `pulumi:"customAction"`
	MinProcessExecutionTime pulumi.StringPtrInput                `pulumi:"minProcessExecutionTime"`
}

func (AutoHealActionsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealActionsResponse)(nil)).Elem()
}

func (i AutoHealActionsResponseArgs) ToAutoHealActionsResponseOutput() AutoHealActionsResponseOutput {
	return i.ToAutoHealActionsResponseOutputWithContext(context.Background())
}

func (i AutoHealActionsResponseArgs) ToAutoHealActionsResponseOutputWithContext(ctx context.Context) AutoHealActionsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsResponseOutput)
}

func (i AutoHealActionsResponseArgs) ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput {
	return i.ToAutoHealActionsResponsePtrOutputWithContext(context.Background())
}

func (i AutoHealActionsResponseArgs) ToAutoHealActionsResponsePtrOutputWithContext(ctx context.Context) AutoHealActionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsResponseOutput).ToAutoHealActionsResponsePtrOutputWithContext(ctx)
}









type AutoHealActionsResponsePtrInput interface {
	pulumi.Input

	ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput
	ToAutoHealActionsResponsePtrOutputWithContext(context.Context) AutoHealActionsResponsePtrOutput
}

type autoHealActionsResponsePtrType AutoHealActionsResponseArgs

func AutoHealActionsResponsePtr(v *AutoHealActionsResponseArgs) AutoHealActionsResponsePtrInput {
	return (*autoHealActionsResponsePtrType)(v)
}

func (*autoHealActionsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealActionsResponse)(nil)).Elem()
}

func (i *autoHealActionsResponsePtrType) ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput {
	return i.ToAutoHealActionsResponsePtrOutputWithContext(context.Background())
}

func (i *autoHealActionsResponsePtrType) ToAutoHealActionsResponsePtrOutputWithContext(ctx context.Context) AutoHealActionsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealActionsResponsePtrOutput)
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

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponsePtrOutput() AutoHealActionsResponsePtrOutput {
	return o.ToAutoHealActionsResponsePtrOutputWithContext(context.Background())
}

func (o AutoHealActionsResponseOutput) ToAutoHealActionsResponsePtrOutputWithContext(ctx context.Context) AutoHealActionsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealActionsResponse) *AutoHealActionsResponse {
		return &v
	}).(AutoHealActionsResponsePtrOutput)
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





type AutoHealCustomActionResponseInput interface {
	pulumi.Input

	ToAutoHealCustomActionResponseOutput() AutoHealCustomActionResponseOutput
	ToAutoHealCustomActionResponseOutputWithContext(context.Context) AutoHealCustomActionResponseOutput
}

type AutoHealCustomActionResponseArgs struct {
	Exe        pulumi.StringPtrInput `pulumi:"exe"`
	Parameters pulumi.StringPtrInput `pulumi:"parameters"`
}

func (AutoHealCustomActionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealCustomActionResponse)(nil)).Elem()
}

func (i AutoHealCustomActionResponseArgs) ToAutoHealCustomActionResponseOutput() AutoHealCustomActionResponseOutput {
	return i.ToAutoHealCustomActionResponseOutputWithContext(context.Background())
}

func (i AutoHealCustomActionResponseArgs) ToAutoHealCustomActionResponseOutputWithContext(ctx context.Context) AutoHealCustomActionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionResponseOutput)
}

func (i AutoHealCustomActionResponseArgs) ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput {
	return i.ToAutoHealCustomActionResponsePtrOutputWithContext(context.Background())
}

func (i AutoHealCustomActionResponseArgs) ToAutoHealCustomActionResponsePtrOutputWithContext(ctx context.Context) AutoHealCustomActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionResponseOutput).ToAutoHealCustomActionResponsePtrOutputWithContext(ctx)
}









type AutoHealCustomActionResponsePtrInput interface {
	pulumi.Input

	ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput
	ToAutoHealCustomActionResponsePtrOutputWithContext(context.Context) AutoHealCustomActionResponsePtrOutput
}

type autoHealCustomActionResponsePtrType AutoHealCustomActionResponseArgs

func AutoHealCustomActionResponsePtr(v *AutoHealCustomActionResponseArgs) AutoHealCustomActionResponsePtrInput {
	return (*autoHealCustomActionResponsePtrType)(v)
}

func (*autoHealCustomActionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealCustomActionResponse)(nil)).Elem()
}

func (i *autoHealCustomActionResponsePtrType) ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput {
	return i.ToAutoHealCustomActionResponsePtrOutputWithContext(context.Background())
}

func (i *autoHealCustomActionResponsePtrType) ToAutoHealCustomActionResponsePtrOutputWithContext(ctx context.Context) AutoHealCustomActionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealCustomActionResponsePtrOutput)
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

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponsePtrOutput() AutoHealCustomActionResponsePtrOutput {
	return o.ToAutoHealCustomActionResponsePtrOutputWithContext(context.Background())
}

func (o AutoHealCustomActionResponseOutput) ToAutoHealCustomActionResponsePtrOutputWithContext(ctx context.Context) AutoHealCustomActionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealCustomActionResponse) *AutoHealCustomActionResponse {
		return &v
	}).(AutoHealCustomActionResponsePtrOutput)
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





type AutoHealRulesResponseInput interface {
	pulumi.Input

	ToAutoHealRulesResponseOutput() AutoHealRulesResponseOutput
	ToAutoHealRulesResponseOutputWithContext(context.Context) AutoHealRulesResponseOutput
}

type AutoHealRulesResponseArgs struct {
	Actions  AutoHealActionsResponsePtrInput  `pulumi:"actions"`
	Triggers AutoHealTriggersResponsePtrInput `pulumi:"triggers"`
}

func (AutoHealRulesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealRulesResponse)(nil)).Elem()
}

func (i AutoHealRulesResponseArgs) ToAutoHealRulesResponseOutput() AutoHealRulesResponseOutput {
	return i.ToAutoHealRulesResponseOutputWithContext(context.Background())
}

func (i AutoHealRulesResponseArgs) ToAutoHealRulesResponseOutputWithContext(ctx context.Context) AutoHealRulesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesResponseOutput)
}

func (i AutoHealRulesResponseArgs) ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput {
	return i.ToAutoHealRulesResponsePtrOutputWithContext(context.Background())
}

func (i AutoHealRulesResponseArgs) ToAutoHealRulesResponsePtrOutputWithContext(ctx context.Context) AutoHealRulesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesResponseOutput).ToAutoHealRulesResponsePtrOutputWithContext(ctx)
}









type AutoHealRulesResponsePtrInput interface {
	pulumi.Input

	ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput
	ToAutoHealRulesResponsePtrOutputWithContext(context.Context) AutoHealRulesResponsePtrOutput
}

type autoHealRulesResponsePtrType AutoHealRulesResponseArgs

func AutoHealRulesResponsePtr(v *AutoHealRulesResponseArgs) AutoHealRulesResponsePtrInput {
	return (*autoHealRulesResponsePtrType)(v)
}

func (*autoHealRulesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealRulesResponse)(nil)).Elem()
}

func (i *autoHealRulesResponsePtrType) ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput {
	return i.ToAutoHealRulesResponsePtrOutputWithContext(context.Background())
}

func (i *autoHealRulesResponsePtrType) ToAutoHealRulesResponsePtrOutputWithContext(ctx context.Context) AutoHealRulesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealRulesResponsePtrOutput)
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

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponsePtrOutput() AutoHealRulesResponsePtrOutput {
	return o.ToAutoHealRulesResponsePtrOutputWithContext(context.Background())
}

func (o AutoHealRulesResponseOutput) ToAutoHealRulesResponsePtrOutputWithContext(ctx context.Context) AutoHealRulesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealRulesResponse) *AutoHealRulesResponse {
		return &v
	}).(AutoHealRulesResponsePtrOutput)
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
	PrivateBytesInKB *int                      `pulumi:"privateBytesInKB"`
	Requests         *RequestsBasedTrigger     `pulumi:"requests"`
	SlowRequests     *SlowRequestsBasedTrigger `pulumi:"slowRequests"`
	StatusCodes      []StatusCodesBasedTrigger `pulumi:"statusCodes"`
}





type AutoHealTriggersInput interface {
	pulumi.Input

	ToAutoHealTriggersOutput() AutoHealTriggersOutput
	ToAutoHealTriggersOutputWithContext(context.Context) AutoHealTriggersOutput
}

type AutoHealTriggersArgs struct {
	PrivateBytesInKB pulumi.IntPtrInput                `pulumi:"privateBytesInKB"`
	Requests         RequestsBasedTriggerPtrInput      `pulumi:"requests"`
	SlowRequests     SlowRequestsBasedTriggerPtrInput  `pulumi:"slowRequests"`
	StatusCodes      StatusCodesBasedTriggerArrayInput `pulumi:"statusCodes"`
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

func (o AutoHealTriggersOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v AutoHealTriggers) []StatusCodesBasedTrigger { return v.StatusCodes }).(StatusCodesBasedTriggerArrayOutput)
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

func (o AutoHealTriggersPtrOutput) StatusCodes() StatusCodesBasedTriggerArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggers) []StatusCodesBasedTrigger {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerArrayOutput)
}

type AutoHealTriggersResponse struct {
	PrivateBytesInKB *int                              `pulumi:"privateBytesInKB"`
	Requests         *RequestsBasedTriggerResponse     `pulumi:"requests"`
	SlowRequests     *SlowRequestsBasedTriggerResponse `pulumi:"slowRequests"`
	StatusCodes      []StatusCodesBasedTriggerResponse `pulumi:"statusCodes"`
}





type AutoHealTriggersResponseInput interface {
	pulumi.Input

	ToAutoHealTriggersResponseOutput() AutoHealTriggersResponseOutput
	ToAutoHealTriggersResponseOutputWithContext(context.Context) AutoHealTriggersResponseOutput
}

type AutoHealTriggersResponseArgs struct {
	PrivateBytesInKB pulumi.IntPtrInput                        `pulumi:"privateBytesInKB"`
	Requests         RequestsBasedTriggerResponsePtrInput      `pulumi:"requests"`
	SlowRequests     SlowRequestsBasedTriggerResponsePtrInput  `pulumi:"slowRequests"`
	StatusCodes      StatusCodesBasedTriggerResponseArrayInput `pulumi:"statusCodes"`
}

func (AutoHealTriggersResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AutoHealTriggersResponse)(nil)).Elem()
}

func (i AutoHealTriggersResponseArgs) ToAutoHealTriggersResponseOutput() AutoHealTriggersResponseOutput {
	return i.ToAutoHealTriggersResponseOutputWithContext(context.Background())
}

func (i AutoHealTriggersResponseArgs) ToAutoHealTriggersResponseOutputWithContext(ctx context.Context) AutoHealTriggersResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersResponseOutput)
}

func (i AutoHealTriggersResponseArgs) ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput {
	return i.ToAutoHealTriggersResponsePtrOutputWithContext(context.Background())
}

func (i AutoHealTriggersResponseArgs) ToAutoHealTriggersResponsePtrOutputWithContext(ctx context.Context) AutoHealTriggersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersResponseOutput).ToAutoHealTriggersResponsePtrOutputWithContext(ctx)
}









type AutoHealTriggersResponsePtrInput interface {
	pulumi.Input

	ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput
	ToAutoHealTriggersResponsePtrOutputWithContext(context.Context) AutoHealTriggersResponsePtrOutput
}

type autoHealTriggersResponsePtrType AutoHealTriggersResponseArgs

func AutoHealTriggersResponsePtr(v *AutoHealTriggersResponseArgs) AutoHealTriggersResponsePtrInput {
	return (*autoHealTriggersResponsePtrType)(v)
}

func (*autoHealTriggersResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AutoHealTriggersResponse)(nil)).Elem()
}

func (i *autoHealTriggersResponsePtrType) ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput {
	return i.ToAutoHealTriggersResponsePtrOutputWithContext(context.Background())
}

func (i *autoHealTriggersResponsePtrType) ToAutoHealTriggersResponsePtrOutputWithContext(ctx context.Context) AutoHealTriggersResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AutoHealTriggersResponsePtrOutput)
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

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponsePtrOutput() AutoHealTriggersResponsePtrOutput {
	return o.ToAutoHealTriggersResponsePtrOutputWithContext(context.Background())
}

func (o AutoHealTriggersResponseOutput) ToAutoHealTriggersResponsePtrOutputWithContext(ctx context.Context) AutoHealTriggersResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AutoHealTriggersResponse) *AutoHealTriggersResponse {
		return &v
	}).(AutoHealTriggersResponsePtrOutput)
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

func (o AutoHealTriggersResponseOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse { return v.StatusCodes }).(StatusCodesBasedTriggerResponseArrayOutput)
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

func (o AutoHealTriggersResponsePtrOutput) StatusCodes() StatusCodesBasedTriggerResponseArrayOutput {
	return o.ApplyT(func(v *AutoHealTriggersResponse) []StatusCodesBasedTriggerResponse {
		if v == nil {
			return nil
		}
		return v.StatusCodes
	}).(StatusCodesBasedTriggerResponseArrayOutput)
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





type AzureBlobStorageApplicationLogsConfigResponseInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigResponseOutput() AzureBlobStorageApplicationLogsConfigResponseOutput
	ToAzureBlobStorageApplicationLogsConfigResponseOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigResponseOutput
}

type AzureBlobStorageApplicationLogsConfigResponseArgs struct {
	Level           pulumi.StringPtrInput `pulumi:"level"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageApplicationLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (i AzureBlobStorageApplicationLogsConfigResponseArgs) ToAzureBlobStorageApplicationLogsConfigResponseOutput() AzureBlobStorageApplicationLogsConfigResponseOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigResponseOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigResponseArgs) ToAzureBlobStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigResponseOutput)
}

func (i AzureBlobStorageApplicationLogsConfigResponseArgs) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageApplicationLogsConfigResponseArgs) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigResponseOutput).ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx)
}









type AzureBlobStorageApplicationLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput
	ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput
}

type azureBlobStorageApplicationLogsConfigResponsePtrType AzureBlobStorageApplicationLogsConfigResponseArgs

func AzureBlobStorageApplicationLogsConfigResponsePtr(v *AzureBlobStorageApplicationLogsConfigResponseArgs) AzureBlobStorageApplicationLogsConfigResponsePtrInput {
	return (*azureBlobStorageApplicationLogsConfigResponsePtrType)(v)
}

func (*azureBlobStorageApplicationLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (i *azureBlobStorageApplicationLogsConfigResponsePtrType) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return i.ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageApplicationLogsConfigResponsePtrType) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
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

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutput() AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageApplicationLogsConfigResponseOutput) ToAzureBlobStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageApplicationLogsConfigResponse) *AzureBlobStorageApplicationLogsConfigResponse {
		return &v
	}).(AzureBlobStorageApplicationLogsConfigResponsePtrOutput)
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





type AzureBlobStorageHttpLogsConfigResponseInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigResponseOutput() AzureBlobStorageHttpLogsConfigResponseOutput
	ToAzureBlobStorageHttpLogsConfigResponseOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigResponseOutput
}

type AzureBlobStorageHttpLogsConfigResponseArgs struct {
	Enabled         pulumi.BoolPtrInput   `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput    `pulumi:"retentionInDays"`
	SasUrl          pulumi.StringPtrInput `pulumi:"sasUrl"`
}

func (AzureBlobStorageHttpLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (i AzureBlobStorageHttpLogsConfigResponseArgs) ToAzureBlobStorageHttpLogsConfigResponseOutput() AzureBlobStorageHttpLogsConfigResponseOutput {
	return i.ToAzureBlobStorageHttpLogsConfigResponseOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigResponseArgs) ToAzureBlobStorageHttpLogsConfigResponseOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigResponseOutput)
}

func (i AzureBlobStorageHttpLogsConfigResponseArgs) ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i AzureBlobStorageHttpLogsConfigResponseArgs) ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigResponseOutput).ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx)
}









type AzureBlobStorageHttpLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput
	ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput
}

type azureBlobStorageHttpLogsConfigResponsePtrType AzureBlobStorageHttpLogsConfigResponseArgs

func AzureBlobStorageHttpLogsConfigResponsePtr(v *AzureBlobStorageHttpLogsConfigResponseArgs) AzureBlobStorageHttpLogsConfigResponsePtrInput {
	return (*azureBlobStorageHttpLogsConfigResponsePtrType)(v)
}

func (*azureBlobStorageHttpLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureBlobStorageHttpLogsConfigResponse)(nil)).Elem()
}

func (i *azureBlobStorageHttpLogsConfigResponsePtrType) ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return i.ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *azureBlobStorageHttpLogsConfigResponsePtrType) ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
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

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutput() AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o AzureBlobStorageHttpLogsConfigResponseOutput) ToAzureBlobStorageHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureBlobStorageHttpLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureBlobStorageHttpLogsConfigResponse) *AzureBlobStorageHttpLogsConfigResponse {
		return &v
	}).(AzureBlobStorageHttpLogsConfigResponsePtrOutput)
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





type AzureStorageInfoValueResponseInput interface {
	pulumi.Input

	ToAzureStorageInfoValueResponseOutput() AzureStorageInfoValueResponseOutput
	ToAzureStorageInfoValueResponseOutputWithContext(context.Context) AzureStorageInfoValueResponseOutput
}

type AzureStorageInfoValueResponseArgs struct {
	AccessKey   pulumi.StringPtrInput `pulumi:"accessKey"`
	AccountName pulumi.StringPtrInput `pulumi:"accountName"`
	MountPath   pulumi.StringPtrInput `pulumi:"mountPath"`
	ShareName   pulumi.StringPtrInput `pulumi:"shareName"`
	State       pulumi.StringInput    `pulumi:"state"`
	Type        pulumi.StringPtrInput `pulumi:"type"`
}

func (AzureStorageInfoValueResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureStorageInfoValueResponse)(nil)).Elem()
}

func (i AzureStorageInfoValueResponseArgs) ToAzureStorageInfoValueResponseOutput() AzureStorageInfoValueResponseOutput {
	return i.ToAzureStorageInfoValueResponseOutputWithContext(context.Background())
}

func (i AzureStorageInfoValueResponseArgs) ToAzureStorageInfoValueResponseOutputWithContext(ctx context.Context) AzureStorageInfoValueResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageInfoValueResponseOutput)
}





type AzureStorageInfoValueResponseMapInput interface {
	pulumi.Input

	ToAzureStorageInfoValueResponseMapOutput() AzureStorageInfoValueResponseMapOutput
	ToAzureStorageInfoValueResponseMapOutputWithContext(context.Context) AzureStorageInfoValueResponseMapOutput
}

type AzureStorageInfoValueResponseMap map[string]AzureStorageInfoValueResponseInput

func (AzureStorageInfoValueResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]AzureStorageInfoValueResponse)(nil)).Elem()
}

func (i AzureStorageInfoValueResponseMap) ToAzureStorageInfoValueResponseMapOutput() AzureStorageInfoValueResponseMapOutput {
	return i.ToAzureStorageInfoValueResponseMapOutputWithContext(context.Background())
}

func (i AzureStorageInfoValueResponseMap) ToAzureStorageInfoValueResponseMapOutputWithContext(ctx context.Context) AzureStorageInfoValueResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureStorageInfoValueResponseMapOutput)
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





type AzureTableStorageApplicationLogsConfigResponseInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigResponseOutput() AzureTableStorageApplicationLogsConfigResponseOutput
	ToAzureTableStorageApplicationLogsConfigResponseOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigResponseOutput
}

type AzureTableStorageApplicationLogsConfigResponseArgs struct {
	Level  pulumi.StringPtrInput `pulumi:"level"`
	SasUrl pulumi.StringInput    `pulumi:"sasUrl"`
}

func (AzureTableStorageApplicationLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (i AzureTableStorageApplicationLogsConfigResponseArgs) ToAzureTableStorageApplicationLogsConfigResponseOutput() AzureTableStorageApplicationLogsConfigResponseOutput {
	return i.ToAzureTableStorageApplicationLogsConfigResponseOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigResponseArgs) ToAzureTableStorageApplicationLogsConfigResponseOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigResponseOutput)
}

func (i AzureTableStorageApplicationLogsConfigResponseArgs) ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i AzureTableStorageApplicationLogsConfigResponseArgs) ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigResponseOutput).ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx)
}









type AzureTableStorageApplicationLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput
	ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput
}

type azureTableStorageApplicationLogsConfigResponsePtrType AzureTableStorageApplicationLogsConfigResponseArgs

func AzureTableStorageApplicationLogsConfigResponsePtr(v *AzureTableStorageApplicationLogsConfigResponseArgs) AzureTableStorageApplicationLogsConfigResponsePtrInput {
	return (*azureTableStorageApplicationLogsConfigResponsePtrType)(v)
}

func (*azureTableStorageApplicationLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AzureTableStorageApplicationLogsConfigResponse)(nil)).Elem()
}

func (i *azureTableStorageApplicationLogsConfigResponsePtrType) ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return i.ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *azureTableStorageApplicationLogsConfigResponsePtrType) ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
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

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutput() AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o AzureTableStorageApplicationLogsConfigResponseOutput) ToAzureTableStorageApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) AzureTableStorageApplicationLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AzureTableStorageApplicationLogsConfigResponse) *AzureTableStorageApplicationLogsConfigResponse {
		return &v
	}).(AzureTableStorageApplicationLogsConfigResponsePtrOutput)
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





type BackupItemResponseInput interface {
	pulumi.Input

	ToBackupItemResponseOutput() BackupItemResponseOutput
	ToBackupItemResponseOutputWithContext(context.Context) BackupItemResponseOutput
}

type BackupItemResponseArgs struct {
	BackupId             pulumi.IntInput                         `pulumi:"backupId"`
	BlobName             pulumi.StringInput                      `pulumi:"blobName"`
	CorrelationId        pulumi.StringInput                      `pulumi:"correlationId"`
	Created              pulumi.StringInput                      `pulumi:"created"`
	Databases            DatabaseBackupSettingResponseArrayInput `pulumi:"databases"`
	FinishedTimeStamp    pulumi.StringInput                      `pulumi:"finishedTimeStamp"`
	Id                   pulumi.StringInput                      `pulumi:"id"`
	Kind                 pulumi.StringPtrInput                   `pulumi:"kind"`
	LastRestoreTimeStamp pulumi.StringInput                      `pulumi:"lastRestoreTimeStamp"`
	Log                  pulumi.StringInput                      `pulumi:"log"`
	Name                 pulumi.StringInput                      `pulumi:"name"`
	Scheduled            pulumi.BoolInput                        `pulumi:"scheduled"`
	SizeInBytes          pulumi.Float64Input                     `pulumi:"sizeInBytes"`
	Status               pulumi.StringInput                      `pulumi:"status"`
	StorageAccountUrl    pulumi.StringInput                      `pulumi:"storageAccountUrl"`
	Type                 pulumi.StringInput                      `pulumi:"type"`
	WebsiteSizeInBytes   pulumi.Float64Input                     `pulumi:"websiteSizeInBytes"`
}

func (BackupItemResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupItemResponse)(nil)).Elem()
}

func (i BackupItemResponseArgs) ToBackupItemResponseOutput() BackupItemResponseOutput {
	return i.ToBackupItemResponseOutputWithContext(context.Background())
}

func (i BackupItemResponseArgs) ToBackupItemResponseOutputWithContext(ctx context.Context) BackupItemResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupItemResponseOutput)
}





type BackupItemResponseArrayInput interface {
	pulumi.Input

	ToBackupItemResponseArrayOutput() BackupItemResponseArrayOutput
	ToBackupItemResponseArrayOutputWithContext(context.Context) BackupItemResponseArrayOutput
}

type BackupItemResponseArray []BackupItemResponseInput

func (BackupItemResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]BackupItemResponse)(nil)).Elem()
}

func (i BackupItemResponseArray) ToBackupItemResponseArrayOutput() BackupItemResponseArrayOutput {
	return i.ToBackupItemResponseArrayOutputWithContext(context.Background())
}

func (i BackupItemResponseArray) ToBackupItemResponseArrayOutputWithContext(ctx context.Context) BackupItemResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupItemResponseArrayOutput)
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





type BackupScheduleResponseInput interface {
	pulumi.Input

	ToBackupScheduleResponseOutput() BackupScheduleResponseOutput
	ToBackupScheduleResponseOutputWithContext(context.Context) BackupScheduleResponseOutput
}

type BackupScheduleResponseArgs struct {
	FrequencyInterval     pulumi.IntInput       `pulumi:"frequencyInterval"`
	FrequencyUnit         pulumi.StringInput    `pulumi:"frequencyUnit"`
	KeepAtLeastOneBackup  pulumi.BoolInput      `pulumi:"keepAtLeastOneBackup"`
	LastExecutionTime     pulumi.StringInput    `pulumi:"lastExecutionTime"`
	RetentionPeriodInDays pulumi.IntInput       `pulumi:"retentionPeriodInDays"`
	StartTime             pulumi.StringPtrInput `pulumi:"startTime"`
}

func (BackupScheduleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BackupScheduleResponse)(nil)).Elem()
}

func (i BackupScheduleResponseArgs) ToBackupScheduleResponseOutput() BackupScheduleResponseOutput {
	return i.ToBackupScheduleResponseOutputWithContext(context.Background())
}

func (i BackupScheduleResponseArgs) ToBackupScheduleResponseOutputWithContext(ctx context.Context) BackupScheduleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleResponseOutput)
}

func (i BackupScheduleResponseArgs) ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput {
	return i.ToBackupScheduleResponsePtrOutputWithContext(context.Background())
}

func (i BackupScheduleResponseArgs) ToBackupScheduleResponsePtrOutputWithContext(ctx context.Context) BackupScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleResponseOutput).ToBackupScheduleResponsePtrOutputWithContext(ctx)
}









type BackupScheduleResponsePtrInput interface {
	pulumi.Input

	ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput
	ToBackupScheduleResponsePtrOutputWithContext(context.Context) BackupScheduleResponsePtrOutput
}

type backupScheduleResponsePtrType BackupScheduleResponseArgs

func BackupScheduleResponsePtr(v *BackupScheduleResponseArgs) BackupScheduleResponsePtrInput {
	return (*backupScheduleResponsePtrType)(v)
}

func (*backupScheduleResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BackupScheduleResponse)(nil)).Elem()
}

func (i *backupScheduleResponsePtrType) ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput {
	return i.ToBackupScheduleResponsePtrOutputWithContext(context.Background())
}

func (i *backupScheduleResponsePtrType) ToBackupScheduleResponsePtrOutputWithContext(ctx context.Context) BackupScheduleResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BackupScheduleResponsePtrOutput)
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

func (o BackupScheduleResponseOutput) ToBackupScheduleResponsePtrOutput() BackupScheduleResponsePtrOutput {
	return o.ToBackupScheduleResponsePtrOutputWithContext(context.Background())
}

func (o BackupScheduleResponseOutput) ToBackupScheduleResponsePtrOutputWithContext(ctx context.Context) BackupScheduleResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BackupScheduleResponse) *BackupScheduleResponse {
		return &v
	}).(BackupScheduleResponsePtrOutput)
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





type CapabilityResponseInput interface {
	pulumi.Input

	ToCapabilityResponseOutput() CapabilityResponseOutput
	ToCapabilityResponseOutputWithContext(context.Context) CapabilityResponseOutput
}

type CapabilityResponseArgs struct {
	Name   pulumi.StringPtrInput `pulumi:"name"`
	Reason pulumi.StringPtrInput `pulumi:"reason"`
	Value  pulumi.StringPtrInput `pulumi:"value"`
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

type CloningInfoResponse struct {
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





type CloningInfoResponseInput interface {
	pulumi.Input

	ToCloningInfoResponseOutput() CloningInfoResponseOutput
	ToCloningInfoResponseOutputWithContext(context.Context) CloningInfoResponseOutput
}

type CloningInfoResponseArgs struct {
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

func (CloningInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfoResponse)(nil)).Elem()
}

func (i CloningInfoResponseArgs) ToCloningInfoResponseOutput() CloningInfoResponseOutput {
	return i.ToCloningInfoResponseOutputWithContext(context.Background())
}

func (i CloningInfoResponseArgs) ToCloningInfoResponseOutputWithContext(ctx context.Context) CloningInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CloningInfoResponseOutput)
}

type CloningInfoResponseOutput struct{ *pulumi.OutputState }

func (CloningInfoResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*CloningInfoResponse)(nil)).Elem()
}

func (o CloningInfoResponseOutput) ToCloningInfoResponseOutput() CloningInfoResponseOutput {
	return o
}

func (o CloningInfoResponseOutput) ToCloningInfoResponseOutputWithContext(ctx context.Context) CloningInfoResponseOutput {
	return o
}

func (o CloningInfoResponseOutput) AppSettingsOverrides() pulumi.StringMapOutput {
	return o.ApplyT(func(v CloningInfoResponse) map[string]string { return v.AppSettingsOverrides }).(pulumi.StringMapOutput)
}

func (o CloningInfoResponseOutput) CloneCustomHostNames() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.CloneCustomHostNames }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) CloneSourceControl() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.CloneSourceControl }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) ConfigureLoadBalancing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.ConfigureLoadBalancing }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) CorrelationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.CorrelationId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) HostingEnvironment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.HostingEnvironment }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) Overwrite() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *bool { return v.Overwrite }).(pulumi.BoolPtrOutput)
}

func (o CloningInfoResponseOutput) SourceWebAppId() pulumi.StringOutput {
	return o.ApplyT(func(v CloningInfoResponse) string { return v.SourceWebAppId }).(pulumi.StringOutput)
}

func (o CloningInfoResponseOutput) SourceWebAppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.SourceWebAppLocation }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) TrafficManagerProfileId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.TrafficManagerProfileId }).(pulumi.StringPtrOutput)
}

func (o CloningInfoResponseOutput) TrafficManagerProfileName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v CloningInfoResponse) *string { return v.TrafficManagerProfileName }).(pulumi.StringPtrOutput)
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





type ConnStringInfoResponseInput interface {
	pulumi.Input

	ToConnStringInfoResponseOutput() ConnStringInfoResponseOutput
	ToConnStringInfoResponseOutputWithContext(context.Context) ConnStringInfoResponseOutput
}

type ConnStringInfoResponseArgs struct {
	ConnectionString pulumi.StringPtrInput `pulumi:"connectionString"`
	Name             pulumi.StringPtrInput `pulumi:"name"`
	Type             pulumi.StringPtrInput `pulumi:"type"`
}

func (ConnStringInfoResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringInfoResponse)(nil)).Elem()
}

func (i ConnStringInfoResponseArgs) ToConnStringInfoResponseOutput() ConnStringInfoResponseOutput {
	return i.ToConnStringInfoResponseOutputWithContext(context.Background())
}

func (i ConnStringInfoResponseArgs) ToConnStringInfoResponseOutputWithContext(ctx context.Context) ConnStringInfoResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoResponseOutput)
}





type ConnStringInfoResponseArrayInput interface {
	pulumi.Input

	ToConnStringInfoResponseArrayOutput() ConnStringInfoResponseArrayOutput
	ToConnStringInfoResponseArrayOutputWithContext(context.Context) ConnStringInfoResponseArrayOutput
}

type ConnStringInfoResponseArray []ConnStringInfoResponseInput

func (ConnStringInfoResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnStringInfoResponse)(nil)).Elem()
}

func (i ConnStringInfoResponseArray) ToConnStringInfoResponseArrayOutput() ConnStringInfoResponseArrayOutput {
	return i.ToConnStringInfoResponseArrayOutputWithContext(context.Background())
}

func (i ConnStringInfoResponseArray) ToConnStringInfoResponseArrayOutputWithContext(ctx context.Context) ConnStringInfoResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringInfoResponseArrayOutput)
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





type ConnStringValueTypePairResponseInput interface {
	pulumi.Input

	ToConnStringValueTypePairResponseOutput() ConnStringValueTypePairResponseOutput
	ToConnStringValueTypePairResponseOutputWithContext(context.Context) ConnStringValueTypePairResponseOutput
}

type ConnStringValueTypePairResponseArgs struct {
	Type  pulumi.StringInput `pulumi:"type"`
	Value pulumi.StringInput `pulumi:"value"`
}

func (ConnStringValueTypePairResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnStringValueTypePairResponse)(nil)).Elem()
}

func (i ConnStringValueTypePairResponseArgs) ToConnStringValueTypePairResponseOutput() ConnStringValueTypePairResponseOutput {
	return i.ToConnStringValueTypePairResponseOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairResponseArgs) ToConnStringValueTypePairResponseOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairResponseOutput)
}





type ConnStringValueTypePairResponseMapInput interface {
	pulumi.Input

	ToConnStringValueTypePairResponseMapOutput() ConnStringValueTypePairResponseMapOutput
	ToConnStringValueTypePairResponseMapOutputWithContext(context.Context) ConnStringValueTypePairResponseMapOutput
}

type ConnStringValueTypePairResponseMap map[string]ConnStringValueTypePairResponseInput

func (ConnStringValueTypePairResponseMap) ElementType() reflect.Type {
	return reflect.TypeOf((*map[string]ConnStringValueTypePairResponse)(nil)).Elem()
}

func (i ConnStringValueTypePairResponseMap) ToConnStringValueTypePairResponseMapOutput() ConnStringValueTypePairResponseMapOutput {
	return i.ToConnStringValueTypePairResponseMapOutputWithContext(context.Background())
}

func (i ConnStringValueTypePairResponseMap) ToConnStringValueTypePairResponseMapOutputWithContext(ctx context.Context) ConnStringValueTypePairResponseMapOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnStringValueTypePairResponseMapOutput)
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





type CorsSettingsResponseInput interface {
	pulumi.Input

	ToCorsSettingsResponseOutput() CorsSettingsResponseOutput
	ToCorsSettingsResponseOutputWithContext(context.Context) CorsSettingsResponseOutput
}

type CorsSettingsResponseArgs struct {
	AllowedOrigins     pulumi.StringArrayInput `pulumi:"allowedOrigins"`
	SupportCredentials pulumi.BoolPtrInput     `pulumi:"supportCredentials"`
}

func (CorsSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*CorsSettingsResponse)(nil)).Elem()
}

func (i CorsSettingsResponseArgs) ToCorsSettingsResponseOutput() CorsSettingsResponseOutput {
	return i.ToCorsSettingsResponseOutputWithContext(context.Background())
}

func (i CorsSettingsResponseArgs) ToCorsSettingsResponseOutputWithContext(ctx context.Context) CorsSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsResponseOutput)
}

func (i CorsSettingsResponseArgs) ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput {
	return i.ToCorsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i CorsSettingsResponseArgs) ToCorsSettingsResponsePtrOutputWithContext(ctx context.Context) CorsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsResponseOutput).ToCorsSettingsResponsePtrOutputWithContext(ctx)
}









type CorsSettingsResponsePtrInput interface {
	pulumi.Input

	ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput
	ToCorsSettingsResponsePtrOutputWithContext(context.Context) CorsSettingsResponsePtrOutput
}

type corsSettingsResponsePtrType CorsSettingsResponseArgs

func CorsSettingsResponsePtr(v *CorsSettingsResponseArgs) CorsSettingsResponsePtrInput {
	return (*corsSettingsResponsePtrType)(v)
}

func (*corsSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**CorsSettingsResponse)(nil)).Elem()
}

func (i *corsSettingsResponsePtrType) ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput {
	return i.ToCorsSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *corsSettingsResponsePtrType) ToCorsSettingsResponsePtrOutputWithContext(ctx context.Context) CorsSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(CorsSettingsResponsePtrOutput)
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

func (o CorsSettingsResponseOutput) ToCorsSettingsResponsePtrOutput() CorsSettingsResponsePtrOutput {
	return o.ToCorsSettingsResponsePtrOutputWithContext(context.Background())
}

func (o CorsSettingsResponseOutput) ToCorsSettingsResponsePtrOutputWithContext(ctx context.Context) CorsSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v CorsSettingsResponse) *CorsSettingsResponse {
		return &v
	}).(CorsSettingsResponsePtrOutput)
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





type DatabaseBackupSettingResponseInput interface {
	pulumi.Input

	ToDatabaseBackupSettingResponseOutput() DatabaseBackupSettingResponseOutput
	ToDatabaseBackupSettingResponseOutputWithContext(context.Context) DatabaseBackupSettingResponseOutput
}

type DatabaseBackupSettingResponseArgs struct {
	ConnectionString     pulumi.StringPtrInput `pulumi:"connectionString"`
	ConnectionStringName pulumi.StringPtrInput `pulumi:"connectionStringName"`
	DatabaseType         pulumi.StringInput    `pulumi:"databaseType"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
}

func (DatabaseBackupSettingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DatabaseBackupSettingResponse)(nil)).Elem()
}

func (i DatabaseBackupSettingResponseArgs) ToDatabaseBackupSettingResponseOutput() DatabaseBackupSettingResponseOutput {
	return i.ToDatabaseBackupSettingResponseOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingResponseArgs) ToDatabaseBackupSettingResponseOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingResponseOutput)
}





type DatabaseBackupSettingResponseArrayInput interface {
	pulumi.Input

	ToDatabaseBackupSettingResponseArrayOutput() DatabaseBackupSettingResponseArrayOutput
	ToDatabaseBackupSettingResponseArrayOutputWithContext(context.Context) DatabaseBackupSettingResponseArrayOutput
}

type DatabaseBackupSettingResponseArray []DatabaseBackupSettingResponseInput

func (DatabaseBackupSettingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DatabaseBackupSettingResponse)(nil)).Elem()
}

func (i DatabaseBackupSettingResponseArray) ToDatabaseBackupSettingResponseArrayOutput() DatabaseBackupSettingResponseArrayOutput {
	return i.ToDatabaseBackupSettingResponseArrayOutputWithContext(context.Background())
}

func (i DatabaseBackupSettingResponseArray) ToDatabaseBackupSettingResponseArrayOutputWithContext(ctx context.Context) DatabaseBackupSettingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DatabaseBackupSettingResponseArrayOutput)
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





type EnabledConfigResponseInput interface {
	pulumi.Input

	ToEnabledConfigResponseOutput() EnabledConfigResponseOutput
	ToEnabledConfigResponseOutputWithContext(context.Context) EnabledConfigResponseOutput
}

type EnabledConfigResponseArgs struct {
	Enabled pulumi.BoolPtrInput `pulumi:"enabled"`
}

func (EnabledConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnabledConfigResponse)(nil)).Elem()
}

func (i EnabledConfigResponseArgs) ToEnabledConfigResponseOutput() EnabledConfigResponseOutput {
	return i.ToEnabledConfigResponseOutputWithContext(context.Background())
}

func (i EnabledConfigResponseArgs) ToEnabledConfigResponseOutputWithContext(ctx context.Context) EnabledConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigResponseOutput)
}

func (i EnabledConfigResponseArgs) ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput {
	return i.ToEnabledConfigResponsePtrOutputWithContext(context.Background())
}

func (i EnabledConfigResponseArgs) ToEnabledConfigResponsePtrOutputWithContext(ctx context.Context) EnabledConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigResponseOutput).ToEnabledConfigResponsePtrOutputWithContext(ctx)
}









type EnabledConfigResponsePtrInput interface {
	pulumi.Input

	ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput
	ToEnabledConfigResponsePtrOutputWithContext(context.Context) EnabledConfigResponsePtrOutput
}

type enabledConfigResponsePtrType EnabledConfigResponseArgs

func EnabledConfigResponsePtr(v *EnabledConfigResponseArgs) EnabledConfigResponsePtrInput {
	return (*enabledConfigResponsePtrType)(v)
}

func (*enabledConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnabledConfigResponse)(nil)).Elem()
}

func (i *enabledConfigResponsePtrType) ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput {
	return i.ToEnabledConfigResponsePtrOutputWithContext(context.Background())
}

func (i *enabledConfigResponsePtrType) ToEnabledConfigResponsePtrOutputWithContext(ctx context.Context) EnabledConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnabledConfigResponsePtrOutput)
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

func (o EnabledConfigResponseOutput) ToEnabledConfigResponsePtrOutput() EnabledConfigResponsePtrOutput {
	return o.ToEnabledConfigResponsePtrOutputWithContext(context.Background())
}

func (o EnabledConfigResponseOutput) ToEnabledConfigResponsePtrOutputWithContext(ctx context.Context) EnabledConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnabledConfigResponse) *EnabledConfigResponse {
		return &v
	}).(EnabledConfigResponsePtrOutput)
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





type ExperimentsResponseInput interface {
	pulumi.Input

	ToExperimentsResponseOutput() ExperimentsResponseOutput
	ToExperimentsResponseOutputWithContext(context.Context) ExperimentsResponseOutput
}

type ExperimentsResponseArgs struct {
	RampUpRules RampUpRuleResponseArrayInput `pulumi:"rampUpRules"`
}

func (ExperimentsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ExperimentsResponse)(nil)).Elem()
}

func (i ExperimentsResponseArgs) ToExperimentsResponseOutput() ExperimentsResponseOutput {
	return i.ToExperimentsResponseOutputWithContext(context.Background())
}

func (i ExperimentsResponseArgs) ToExperimentsResponseOutputWithContext(ctx context.Context) ExperimentsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsResponseOutput)
}

func (i ExperimentsResponseArgs) ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput {
	return i.ToExperimentsResponsePtrOutputWithContext(context.Background())
}

func (i ExperimentsResponseArgs) ToExperimentsResponsePtrOutputWithContext(ctx context.Context) ExperimentsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsResponseOutput).ToExperimentsResponsePtrOutputWithContext(ctx)
}









type ExperimentsResponsePtrInput interface {
	pulumi.Input

	ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput
	ToExperimentsResponsePtrOutputWithContext(context.Context) ExperimentsResponsePtrOutput
}

type experimentsResponsePtrType ExperimentsResponseArgs

func ExperimentsResponsePtr(v *ExperimentsResponseArgs) ExperimentsResponsePtrInput {
	return (*experimentsResponsePtrType)(v)
}

func (*experimentsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ExperimentsResponse)(nil)).Elem()
}

func (i *experimentsResponsePtrType) ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput {
	return i.ToExperimentsResponsePtrOutputWithContext(context.Background())
}

func (i *experimentsResponsePtrType) ToExperimentsResponsePtrOutputWithContext(ctx context.Context) ExperimentsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ExperimentsResponsePtrOutput)
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

func (o ExperimentsResponseOutput) ToExperimentsResponsePtrOutput() ExperimentsResponsePtrOutput {
	return o.ToExperimentsResponsePtrOutputWithContext(context.Background())
}

func (o ExperimentsResponseOutput) ToExperimentsResponsePtrOutputWithContext(ctx context.Context) ExperimentsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ExperimentsResponse) *ExperimentsResponse {
		return &v
	}).(ExperimentsResponsePtrOutput)
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

type FileSystemApplicationLogsConfig struct {
	Level *LogLevel `pulumi:"level"`
}





type FileSystemApplicationLogsConfigInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigOutput() FileSystemApplicationLogsConfigOutput
	ToFileSystemApplicationLogsConfigOutputWithContext(context.Context) FileSystemApplicationLogsConfigOutput
}

type FileSystemApplicationLogsConfigArgs struct {
	Level LogLevelPtrInput `pulumi:"level"`
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





type FileSystemApplicationLogsConfigResponseInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigResponseOutput() FileSystemApplicationLogsConfigResponseOutput
	ToFileSystemApplicationLogsConfigResponseOutputWithContext(context.Context) FileSystemApplicationLogsConfigResponseOutput
}

type FileSystemApplicationLogsConfigResponseArgs struct {
	Level pulumi.StringPtrInput `pulumi:"level"`
}

func (FileSystemApplicationLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (i FileSystemApplicationLogsConfigResponseArgs) ToFileSystemApplicationLogsConfigResponseOutput() FileSystemApplicationLogsConfigResponseOutput {
	return i.ToFileSystemApplicationLogsConfigResponseOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigResponseArgs) ToFileSystemApplicationLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigResponseOutput)
}

func (i FileSystemApplicationLogsConfigResponseArgs) ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput {
	return i.ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i FileSystemApplicationLogsConfigResponseArgs) ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigResponseOutput).ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx)
}









type FileSystemApplicationLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput
	ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(context.Context) FileSystemApplicationLogsConfigResponsePtrOutput
}

type fileSystemApplicationLogsConfigResponsePtrType FileSystemApplicationLogsConfigResponseArgs

func FileSystemApplicationLogsConfigResponsePtr(v *FileSystemApplicationLogsConfigResponseArgs) FileSystemApplicationLogsConfigResponsePtrInput {
	return (*fileSystemApplicationLogsConfigResponsePtrType)(v)
}

func (*fileSystemApplicationLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemApplicationLogsConfigResponse)(nil)).Elem()
}

func (i *fileSystemApplicationLogsConfigResponsePtrType) ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput {
	return i.ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *fileSystemApplicationLogsConfigResponsePtrType) ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemApplicationLogsConfigResponsePtrOutput)
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

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponsePtrOutput() FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o FileSystemApplicationLogsConfigResponseOutput) ToFileSystemApplicationLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemApplicationLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemApplicationLogsConfigResponse) *FileSystemApplicationLogsConfigResponse {
		return &v
	}).(FileSystemApplicationLogsConfigResponsePtrOutput)
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





type FileSystemHttpLogsConfigResponseInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigResponseOutput() FileSystemHttpLogsConfigResponseOutput
	ToFileSystemHttpLogsConfigResponseOutputWithContext(context.Context) FileSystemHttpLogsConfigResponseOutput
}

type FileSystemHttpLogsConfigResponseArgs struct {
	Enabled         pulumi.BoolPtrInput `pulumi:"enabled"`
	RetentionInDays pulumi.IntPtrInput  `pulumi:"retentionInDays"`
	RetentionInMb   pulumi.IntPtrInput  `pulumi:"retentionInMb"`
}

func (FileSystemHttpLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (i FileSystemHttpLogsConfigResponseArgs) ToFileSystemHttpLogsConfigResponseOutput() FileSystemHttpLogsConfigResponseOutput {
	return i.ToFileSystemHttpLogsConfigResponseOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigResponseArgs) ToFileSystemHttpLogsConfigResponseOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigResponseOutput)
}

func (i FileSystemHttpLogsConfigResponseArgs) ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput {
	return i.ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i FileSystemHttpLogsConfigResponseArgs) ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigResponseOutput).ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx)
}









type FileSystemHttpLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput
	ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(context.Context) FileSystemHttpLogsConfigResponsePtrOutput
}

type fileSystemHttpLogsConfigResponsePtrType FileSystemHttpLogsConfigResponseArgs

func FileSystemHttpLogsConfigResponsePtr(v *FileSystemHttpLogsConfigResponseArgs) FileSystemHttpLogsConfigResponsePtrInput {
	return (*fileSystemHttpLogsConfigResponsePtrType)(v)
}

func (*fileSystemHttpLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FileSystemHttpLogsConfigResponse)(nil)).Elem()
}

func (i *fileSystemHttpLogsConfigResponsePtrType) ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput {
	return i.ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *fileSystemHttpLogsConfigResponsePtrType) ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FileSystemHttpLogsConfigResponsePtrOutput)
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

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponsePtrOutput() FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o FileSystemHttpLogsConfigResponseOutput) ToFileSystemHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) FileSystemHttpLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FileSystemHttpLogsConfigResponse) *FileSystemHttpLogsConfigResponse {
		return &v
	}).(FileSystemHttpLogsConfigResponsePtrOutput)
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





type HandlerMappingResponseInput interface {
	pulumi.Input

	ToHandlerMappingResponseOutput() HandlerMappingResponseOutput
	ToHandlerMappingResponseOutputWithContext(context.Context) HandlerMappingResponseOutput
}

type HandlerMappingResponseArgs struct {
	Arguments       pulumi.StringPtrInput `pulumi:"arguments"`
	Extension       pulumi.StringPtrInput `pulumi:"extension"`
	ScriptProcessor pulumi.StringPtrInput `pulumi:"scriptProcessor"`
}

func (HandlerMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HandlerMappingResponse)(nil)).Elem()
}

func (i HandlerMappingResponseArgs) ToHandlerMappingResponseOutput() HandlerMappingResponseOutput {
	return i.ToHandlerMappingResponseOutputWithContext(context.Background())
}

func (i HandlerMappingResponseArgs) ToHandlerMappingResponseOutputWithContext(ctx context.Context) HandlerMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingResponseOutput)
}





type HandlerMappingResponseArrayInput interface {
	pulumi.Input

	ToHandlerMappingResponseArrayOutput() HandlerMappingResponseArrayOutput
	ToHandlerMappingResponseArrayOutputWithContext(context.Context) HandlerMappingResponseArrayOutput
}

type HandlerMappingResponseArray []HandlerMappingResponseInput

func (HandlerMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HandlerMappingResponse)(nil)).Elem()
}

func (i HandlerMappingResponseArray) ToHandlerMappingResponseArrayOutput() HandlerMappingResponseArrayOutput {
	return i.ToHandlerMappingResponseArrayOutputWithContext(context.Background())
}

func (i HandlerMappingResponseArray) ToHandlerMappingResponseArrayOutputWithContext(ctx context.Context) HandlerMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HandlerMappingResponseArrayOutput)
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





type HostNameSslStateResponseInput interface {
	pulumi.Input

	ToHostNameSslStateResponseOutput() HostNameSslStateResponseOutput
	ToHostNameSslStateResponseOutputWithContext(context.Context) HostNameSslStateResponseOutput
}

type HostNameSslStateResponseArgs struct {
	HostType   pulumi.StringPtrInput `pulumi:"hostType"`
	Name       pulumi.StringPtrInput `pulumi:"name"`
	SslState   pulumi.StringPtrInput `pulumi:"sslState"`
	Thumbprint pulumi.StringPtrInput `pulumi:"thumbprint"`
	ToUpdate   pulumi.BoolPtrInput   `pulumi:"toUpdate"`
	VirtualIP  pulumi.StringPtrInput `pulumi:"virtualIP"`
}

func (HostNameSslStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostNameSslStateResponse)(nil)).Elem()
}

func (i HostNameSslStateResponseArgs) ToHostNameSslStateResponseOutput() HostNameSslStateResponseOutput {
	return i.ToHostNameSslStateResponseOutputWithContext(context.Background())
}

func (i HostNameSslStateResponseArgs) ToHostNameSslStateResponseOutputWithContext(ctx context.Context) HostNameSslStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateResponseOutput)
}





type HostNameSslStateResponseArrayInput interface {
	pulumi.Input

	ToHostNameSslStateResponseArrayOutput() HostNameSslStateResponseArrayOutput
	ToHostNameSslStateResponseArrayOutputWithContext(context.Context) HostNameSslStateResponseArrayOutput
}

type HostNameSslStateResponseArray []HostNameSslStateResponseInput

func (HostNameSslStateResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]HostNameSslStateResponse)(nil)).Elem()
}

func (i HostNameSslStateResponseArray) ToHostNameSslStateResponseArrayOutput() HostNameSslStateResponseArrayOutput {
	return i.ToHostNameSslStateResponseArrayOutputWithContext(context.Background())
}

func (i HostNameSslStateResponseArray) ToHostNameSslStateResponseArrayOutputWithContext(ctx context.Context) HostNameSslStateResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostNameSslStateResponseArrayOutput)
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





type HostingEnvironmentProfileResponseInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput
	ToHostingEnvironmentProfileResponseOutputWithContext(context.Context) HostingEnvironmentProfileResponseOutput
}

type HostingEnvironmentProfileResponseArgs struct {
	Id   pulumi.StringPtrInput `pulumi:"id"`
	Name pulumi.StringInput    `pulumi:"name"`
	Type pulumi.StringInput    `pulumi:"type"`
}

func (HostingEnvironmentProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponseOutput() HostingEnvironmentProfileResponseOutput {
	return i.ToHostingEnvironmentProfileResponseOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponseOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileResponseOutput)
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return i.ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Background())
}

func (i HostingEnvironmentProfileResponseArgs) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileResponseOutput).ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx)
}









type HostingEnvironmentProfileResponsePtrInput interface {
	pulumi.Input

	ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput
	ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Context) HostingEnvironmentProfileResponsePtrOutput
}

type hostingEnvironmentProfileResponsePtrType HostingEnvironmentProfileResponseArgs

func HostingEnvironmentProfileResponsePtr(v *HostingEnvironmentProfileResponseArgs) HostingEnvironmentProfileResponsePtrInput {
	return (*hostingEnvironmentProfileResponsePtrType)(v)
}

func (*hostingEnvironmentProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HostingEnvironmentProfileResponse)(nil)).Elem()
}

func (i *hostingEnvironmentProfileResponsePtrType) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return i.ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Background())
}

func (i *hostingEnvironmentProfileResponsePtrType) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HostingEnvironmentProfileResponsePtrOutput)
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

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponsePtrOutput() HostingEnvironmentProfileResponsePtrOutput {
	return o.ToHostingEnvironmentProfileResponsePtrOutputWithContext(context.Background())
}

func (o HostingEnvironmentProfileResponseOutput) ToHostingEnvironmentProfileResponsePtrOutputWithContext(ctx context.Context) HostingEnvironmentProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HostingEnvironmentProfileResponse) *HostingEnvironmentProfileResponse {
		return &v
	}).(HostingEnvironmentProfileResponsePtrOutput)
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





type HttpLogsConfigResponseInput interface {
	pulumi.Input

	ToHttpLogsConfigResponseOutput() HttpLogsConfigResponseOutput
	ToHttpLogsConfigResponseOutputWithContext(context.Context) HttpLogsConfigResponseOutput
}

type HttpLogsConfigResponseArgs struct {
	AzureBlobStorage AzureBlobStorageHttpLogsConfigResponsePtrInput `pulumi:"azureBlobStorage"`
	FileSystem       FileSystemHttpLogsConfigResponsePtrInput       `pulumi:"fileSystem"`
}

func (HttpLogsConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*HttpLogsConfigResponse)(nil)).Elem()
}

func (i HttpLogsConfigResponseArgs) ToHttpLogsConfigResponseOutput() HttpLogsConfigResponseOutput {
	return i.ToHttpLogsConfigResponseOutputWithContext(context.Background())
}

func (i HttpLogsConfigResponseArgs) ToHttpLogsConfigResponseOutputWithContext(ctx context.Context) HttpLogsConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigResponseOutput)
}

func (i HttpLogsConfigResponseArgs) ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput {
	return i.ToHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i HttpLogsConfigResponseArgs) ToHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) HttpLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigResponseOutput).ToHttpLogsConfigResponsePtrOutputWithContext(ctx)
}









type HttpLogsConfigResponsePtrInput interface {
	pulumi.Input

	ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput
	ToHttpLogsConfigResponsePtrOutputWithContext(context.Context) HttpLogsConfigResponsePtrOutput
}

type httpLogsConfigResponsePtrType HttpLogsConfigResponseArgs

func HttpLogsConfigResponsePtr(v *HttpLogsConfigResponseArgs) HttpLogsConfigResponsePtrInput {
	return (*httpLogsConfigResponsePtrType)(v)
}

func (*httpLogsConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**HttpLogsConfigResponse)(nil)).Elem()
}

func (i *httpLogsConfigResponsePtrType) ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput {
	return i.ToHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (i *httpLogsConfigResponsePtrType) ToHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) HttpLogsConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(HttpLogsConfigResponsePtrOutput)
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

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponsePtrOutput() HttpLogsConfigResponsePtrOutput {
	return o.ToHttpLogsConfigResponsePtrOutputWithContext(context.Background())
}

func (o HttpLogsConfigResponseOutput) ToHttpLogsConfigResponsePtrOutputWithContext(ctx context.Context) HttpLogsConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v HttpLogsConfigResponse) *HttpLogsConfigResponse {
		return &v
	}).(HttpLogsConfigResponsePtrOutput)
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

type IdentifierResponse struct {
	Id    string  `pulumi:"id"`
	Kind  *string `pulumi:"kind"`
	Name  string  `pulumi:"name"`
	Type  string  `pulumi:"type"`
	Value *string `pulumi:"value"`
}





type IdentifierResponseInput interface {
	pulumi.Input

	ToIdentifierResponseOutput() IdentifierResponseOutput
	ToIdentifierResponseOutputWithContext(context.Context) IdentifierResponseOutput
}

type IdentifierResponseArgs struct {
	Id    pulumi.StringInput    `pulumi:"id"`
	Kind  pulumi.StringPtrInput `pulumi:"kind"`
	Name  pulumi.StringInput    `pulumi:"name"`
	Type  pulumi.StringInput    `pulumi:"type"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (IdentifierResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IdentifierResponse)(nil)).Elem()
}

func (i IdentifierResponseArgs) ToIdentifierResponseOutput() IdentifierResponseOutput {
	return i.ToIdentifierResponseOutputWithContext(context.Background())
}

func (i IdentifierResponseArgs) ToIdentifierResponseOutputWithContext(ctx context.Context) IdentifierResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentifierResponseOutput)
}





type IdentifierResponseArrayInput interface {
	pulumi.Input

	ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput
	ToIdentifierResponseArrayOutputWithContext(context.Context) IdentifierResponseArrayOutput
}

type IdentifierResponseArray []IdentifierResponseInput

func (IdentifierResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IdentifierResponse)(nil)).Elem()
}

func (i IdentifierResponseArray) ToIdentifierResponseArrayOutput() IdentifierResponseArrayOutput {
	return i.ToIdentifierResponseArrayOutputWithContext(context.Background())
}

func (i IdentifierResponseArray) ToIdentifierResponseArrayOutputWithContext(ctx context.Context) IdentifierResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IdentifierResponseArrayOutput)
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

type IpSecurityRestriction struct {
	Action               *string      `pulumi:"action"`
	Description          *string      `pulumi:"description"`
	IpAddress            *string      `pulumi:"ipAddress"`
	Name                 *string      `pulumi:"name"`
	Priority             *int         `pulumi:"priority"`
	SubnetMask           *string      `pulumi:"subnetMask"`
	SubnetTrafficTag     *int         `pulumi:"subnetTrafficTag"`
	Tag                  *IpFilterTag `pulumi:"tag"`
	VnetSubnetResourceId *string      `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       *int         `pulumi:"vnetTrafficTag"`
}





type IpSecurityRestrictionInput interface {
	pulumi.Input

	ToIpSecurityRestrictionOutput() IpSecurityRestrictionOutput
	ToIpSecurityRestrictionOutputWithContext(context.Context) IpSecurityRestrictionOutput
}

type IpSecurityRestrictionArgs struct {
	Action               pulumi.StringPtrInput `pulumi:"action"`
	Description          pulumi.StringPtrInput `pulumi:"description"`
	IpAddress            pulumi.StringPtrInput `pulumi:"ipAddress"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	Priority             pulumi.IntPtrInput    `pulumi:"priority"`
	SubnetMask           pulumi.StringPtrInput `pulumi:"subnetMask"`
	SubnetTrafficTag     pulumi.IntPtrInput    `pulumi:"subnetTrafficTag"`
	Tag                  IpFilterTagPtrInput   `pulumi:"tag"`
	VnetSubnetResourceId pulumi.StringPtrInput `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       pulumi.IntPtrInput    `pulumi:"vnetTrafficTag"`
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

func (o IpSecurityRestrictionOutput) Tag() IpFilterTagPtrOutput {
	return o.ApplyT(func(v IpSecurityRestriction) *IpFilterTag { return v.Tag }).(IpFilterTagPtrOutput)
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
	Action               *string `pulumi:"action"`
	Description          *string `pulumi:"description"`
	IpAddress            *string `pulumi:"ipAddress"`
	Name                 *string `pulumi:"name"`
	Priority             *int    `pulumi:"priority"`
	SubnetMask           *string `pulumi:"subnetMask"`
	SubnetTrafficTag     *int    `pulumi:"subnetTrafficTag"`
	Tag                  *string `pulumi:"tag"`
	VnetSubnetResourceId *string `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       *int    `pulumi:"vnetTrafficTag"`
}





type IpSecurityRestrictionResponseInput interface {
	pulumi.Input

	ToIpSecurityRestrictionResponseOutput() IpSecurityRestrictionResponseOutput
	ToIpSecurityRestrictionResponseOutputWithContext(context.Context) IpSecurityRestrictionResponseOutput
}

type IpSecurityRestrictionResponseArgs struct {
	Action               pulumi.StringPtrInput `pulumi:"action"`
	Description          pulumi.StringPtrInput `pulumi:"description"`
	IpAddress            pulumi.StringPtrInput `pulumi:"ipAddress"`
	Name                 pulumi.StringPtrInput `pulumi:"name"`
	Priority             pulumi.IntPtrInput    `pulumi:"priority"`
	SubnetMask           pulumi.StringPtrInput `pulumi:"subnetMask"`
	SubnetTrafficTag     pulumi.IntPtrInput    `pulumi:"subnetTrafficTag"`
	Tag                  pulumi.StringPtrInput `pulumi:"tag"`
	VnetSubnetResourceId pulumi.StringPtrInput `pulumi:"vnetSubnetResourceId"`
	VnetTrafficTag       pulumi.IntPtrInput    `pulumi:"vnetTrafficTag"`
}

func (IpSecurityRestrictionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*IpSecurityRestrictionResponse)(nil)).Elem()
}

func (i IpSecurityRestrictionResponseArgs) ToIpSecurityRestrictionResponseOutput() IpSecurityRestrictionResponseOutput {
	return i.ToIpSecurityRestrictionResponseOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionResponseArgs) ToIpSecurityRestrictionResponseOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionResponseOutput)
}





type IpSecurityRestrictionResponseArrayInput interface {
	pulumi.Input

	ToIpSecurityRestrictionResponseArrayOutput() IpSecurityRestrictionResponseArrayOutput
	ToIpSecurityRestrictionResponseArrayOutputWithContext(context.Context) IpSecurityRestrictionResponseArrayOutput
}

type IpSecurityRestrictionResponseArray []IpSecurityRestrictionResponseInput

func (IpSecurityRestrictionResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]IpSecurityRestrictionResponse)(nil)).Elem()
}

func (i IpSecurityRestrictionResponseArray) ToIpSecurityRestrictionResponseArrayOutput() IpSecurityRestrictionResponseArrayOutput {
	return i.ToIpSecurityRestrictionResponseArrayOutputWithContext(context.Background())
}

func (i IpSecurityRestrictionResponseArray) ToIpSecurityRestrictionResponseArrayOutputWithContext(ctx context.Context) IpSecurityRestrictionResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(IpSecurityRestrictionResponseArrayOutput)
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





type NameValuePairResponseInput interface {
	pulumi.Input

	ToNameValuePairResponseOutput() NameValuePairResponseOutput
	ToNameValuePairResponseOutputWithContext(context.Context) NameValuePairResponseOutput
}

type NameValuePairResponseArgs struct {
	Name  pulumi.StringPtrInput `pulumi:"name"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (NameValuePairResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NameValuePairResponse)(nil)).Elem()
}

func (i NameValuePairResponseArgs) ToNameValuePairResponseOutput() NameValuePairResponseOutput {
	return i.ToNameValuePairResponseOutputWithContext(context.Background())
}

func (i NameValuePairResponseArgs) ToNameValuePairResponseOutputWithContext(ctx context.Context) NameValuePairResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairResponseOutput)
}





type NameValuePairResponseArrayInput interface {
	pulumi.Input

	ToNameValuePairResponseArrayOutput() NameValuePairResponseArrayOutput
	ToNameValuePairResponseArrayOutputWithContext(context.Context) NameValuePairResponseArrayOutput
}

type NameValuePairResponseArray []NameValuePairResponseInput

func (NameValuePairResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NameValuePairResponse)(nil)).Elem()
}

func (i NameValuePairResponseArray) ToNameValuePairResponseArrayOutput() NameValuePairResponseArrayOutput {
	return i.ToNameValuePairResponseArrayOutputWithContext(context.Background())
}

func (i NameValuePairResponseArray) ToNameValuePairResponseArrayOutputWithContext(ctx context.Context) NameValuePairResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NameValuePairResponseArrayOutput)
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

type NetworkAccessControlEntry struct {
	Action       *AccessControlEntryAction `pulumi:"action"`
	Description  *string                   `pulumi:"description"`
	Order        *int                      `pulumi:"order"`
	RemoteSubnet *string                   `pulumi:"remoteSubnet"`
}





type NetworkAccessControlEntryInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput
	ToNetworkAccessControlEntryOutputWithContext(context.Context) NetworkAccessControlEntryOutput
}

type NetworkAccessControlEntryArgs struct {
	Action       AccessControlEntryActionPtrInput `pulumi:"action"`
	Description  pulumi.StringPtrInput            `pulumi:"description"`
	Order        pulumi.IntPtrInput               `pulumi:"order"`
	RemoteSubnet pulumi.StringPtrInput            `pulumi:"remoteSubnet"`
}

func (NetworkAccessControlEntryArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntry)(nil)).Elem()
}

func (i NetworkAccessControlEntryArgs) ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput {
	return i.ToNetworkAccessControlEntryOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryArgs) ToNetworkAccessControlEntryOutputWithContext(ctx context.Context) NetworkAccessControlEntryOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryOutput)
}





type NetworkAccessControlEntryArrayInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput
	ToNetworkAccessControlEntryArrayOutputWithContext(context.Context) NetworkAccessControlEntryArrayOutput
}

type NetworkAccessControlEntryArray []NetworkAccessControlEntryInput

func (NetworkAccessControlEntryArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntry)(nil)).Elem()
}

func (i NetworkAccessControlEntryArray) ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput {
	return i.ToNetworkAccessControlEntryArrayOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryArray) ToNetworkAccessControlEntryArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryArrayOutput)
}

type NetworkAccessControlEntryOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntry)(nil)).Elem()
}

func (o NetworkAccessControlEntryOutput) ToNetworkAccessControlEntryOutput() NetworkAccessControlEntryOutput {
	return o
}

func (o NetworkAccessControlEntryOutput) ToNetworkAccessControlEntryOutputWithContext(ctx context.Context) NetworkAccessControlEntryOutput {
	return o
}

func (o NetworkAccessControlEntryOutput) Action() AccessControlEntryActionPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *AccessControlEntryAction { return v.Action }).(AccessControlEntryActionPtrOutput)
}

func (o NetworkAccessControlEntryOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o NetworkAccessControlEntryOutput) RemoteSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntry) *string { return v.RemoteSubnet }).(pulumi.StringPtrOutput)
}

type NetworkAccessControlEntryArrayOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntry)(nil)).Elem()
}

func (o NetworkAccessControlEntryArrayOutput) ToNetworkAccessControlEntryArrayOutput() NetworkAccessControlEntryArrayOutput {
	return o
}

func (o NetworkAccessControlEntryArrayOutput) ToNetworkAccessControlEntryArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryArrayOutput {
	return o
}

func (o NetworkAccessControlEntryArrayOutput) Index(i pulumi.IntInput) NetworkAccessControlEntryOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkAccessControlEntry {
		return vs[0].([]NetworkAccessControlEntry)[vs[1].(int)]
	}).(NetworkAccessControlEntryOutput)
}

type NetworkAccessControlEntryResponse struct {
	Action       *string `pulumi:"action"`
	Description  *string `pulumi:"description"`
	Order        *int    `pulumi:"order"`
	RemoteSubnet *string `pulumi:"remoteSubnet"`
}





type NetworkAccessControlEntryResponseInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryResponseOutput() NetworkAccessControlEntryResponseOutput
	ToNetworkAccessControlEntryResponseOutputWithContext(context.Context) NetworkAccessControlEntryResponseOutput
}

type NetworkAccessControlEntryResponseArgs struct {
	Action       pulumi.StringPtrInput `pulumi:"action"`
	Description  pulumi.StringPtrInput `pulumi:"description"`
	Order        pulumi.IntPtrInput    `pulumi:"order"`
	RemoteSubnet pulumi.StringPtrInput `pulumi:"remoteSubnet"`
}

func (NetworkAccessControlEntryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (i NetworkAccessControlEntryResponseArgs) ToNetworkAccessControlEntryResponseOutput() NetworkAccessControlEntryResponseOutput {
	return i.ToNetworkAccessControlEntryResponseOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryResponseArgs) ToNetworkAccessControlEntryResponseOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryResponseOutput)
}





type NetworkAccessControlEntryResponseArrayInput interface {
	pulumi.Input

	ToNetworkAccessControlEntryResponseArrayOutput() NetworkAccessControlEntryResponseArrayOutput
	ToNetworkAccessControlEntryResponseArrayOutputWithContext(context.Context) NetworkAccessControlEntryResponseArrayOutput
}

type NetworkAccessControlEntryResponseArray []NetworkAccessControlEntryResponseInput

func (NetworkAccessControlEntryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (i NetworkAccessControlEntryResponseArray) ToNetworkAccessControlEntryResponseArrayOutput() NetworkAccessControlEntryResponseArrayOutput {
	return i.ToNetworkAccessControlEntryResponseArrayOutputWithContext(context.Background())
}

func (i NetworkAccessControlEntryResponseArray) ToNetworkAccessControlEntryResponseArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(NetworkAccessControlEntryResponseArrayOutput)
}

type NetworkAccessControlEntryResponseOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (o NetworkAccessControlEntryResponseOutput) ToNetworkAccessControlEntryResponseOutput() NetworkAccessControlEntryResponseOutput {
	return o
}

func (o NetworkAccessControlEntryResponseOutput) ToNetworkAccessControlEntryResponseOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseOutput {
	return o
}

func (o NetworkAccessControlEntryResponseOutput) Action() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.Action }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) Order() pulumi.IntPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *int { return v.Order }).(pulumi.IntPtrOutput)
}

func (o NetworkAccessControlEntryResponseOutput) RemoteSubnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v NetworkAccessControlEntryResponse) *string { return v.RemoteSubnet }).(pulumi.StringPtrOutput)
}

type NetworkAccessControlEntryResponseArrayOutput struct{ *pulumi.OutputState }

func (NetworkAccessControlEntryResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]NetworkAccessControlEntryResponse)(nil)).Elem()
}

func (o NetworkAccessControlEntryResponseArrayOutput) ToNetworkAccessControlEntryResponseArrayOutput() NetworkAccessControlEntryResponseArrayOutput {
	return o
}

func (o NetworkAccessControlEntryResponseArrayOutput) ToNetworkAccessControlEntryResponseArrayOutputWithContext(ctx context.Context) NetworkAccessControlEntryResponseArrayOutput {
	return o
}

func (o NetworkAccessControlEntryResponseArrayOutput) Index(i pulumi.IntInput) NetworkAccessControlEntryResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) NetworkAccessControlEntryResponse {
		return vs[0].([]NetworkAccessControlEntryResponse)[vs[1].(int)]
	}).(NetworkAccessControlEntryResponseOutput)
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





type PrivateLinkConnectionStateResponseInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput
	ToPrivateLinkConnectionStateResponseOutputWithContext(context.Context) PrivateLinkConnectionStateResponseOutput
}

type PrivateLinkConnectionStateResponseArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkConnectionStateResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (i PrivateLinkConnectionStateResponseArgs) ToPrivateLinkConnectionStateResponseOutput() PrivateLinkConnectionStateResponseOutput {
	return i.ToPrivateLinkConnectionStateResponseOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateResponseArgs) ToPrivateLinkConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateResponseOutput)
}

func (i PrivateLinkConnectionStateResponseArgs) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i PrivateLinkConnectionStateResponseArgs) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateResponseOutput).ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx)
}









type PrivateLinkConnectionStateResponsePtrInput interface {
	pulumi.Input

	ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput
	ToPrivateLinkConnectionStateResponsePtrOutputWithContext(context.Context) PrivateLinkConnectionStateResponsePtrOutput
}

type privateLinkConnectionStateResponsePtrType PrivateLinkConnectionStateResponseArgs

func PrivateLinkConnectionStateResponsePtr(v *PrivateLinkConnectionStateResponseArgs) PrivateLinkConnectionStateResponsePtrInput {
	return (*privateLinkConnectionStateResponsePtrType)(v)
}

func (*privateLinkConnectionStateResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateLinkConnectionStateResponse)(nil)).Elem()
}

func (i *privateLinkConnectionStateResponsePtrType) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return i.ToPrivateLinkConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (i *privateLinkConnectionStateResponsePtrType) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkConnectionStateResponsePtrOutput)
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

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponsePtrOutput() PrivateLinkConnectionStateResponsePtrOutput {
	return o.ToPrivateLinkConnectionStateResponsePtrOutputWithContext(context.Background())
}

func (o PrivateLinkConnectionStateResponseOutput) ToPrivateLinkConnectionStateResponsePtrOutputWithContext(ctx context.Context) PrivateLinkConnectionStateResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PrivateLinkConnectionStateResponse) *PrivateLinkConnectionStateResponse {
		return &v
	}).(PrivateLinkConnectionStateResponsePtrOutput)
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





type PushSettingsResponseInput interface {
	pulumi.Input

	ToPushSettingsResponseOutput() PushSettingsResponseOutput
	ToPushSettingsResponseOutputWithContext(context.Context) PushSettingsResponseOutput
}

type PushSettingsResponseArgs struct {
	DynamicTagsJson   pulumi.StringPtrInput `pulumi:"dynamicTagsJson"`
	Id                pulumi.StringInput    `pulumi:"id"`
	IsPushEnabled     pulumi.BoolInput      `pulumi:"isPushEnabled"`
	Kind              pulumi.StringPtrInput `pulumi:"kind"`
	Name              pulumi.StringInput    `pulumi:"name"`
	TagWhitelistJson  pulumi.StringPtrInput `pulumi:"tagWhitelistJson"`
	TagsRequiringAuth pulumi.StringPtrInput `pulumi:"tagsRequiringAuth"`
	Type              pulumi.StringInput    `pulumi:"type"`
}

func (PushSettingsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PushSettingsResponse)(nil)).Elem()
}

func (i PushSettingsResponseArgs) ToPushSettingsResponseOutput() PushSettingsResponseOutput {
	return i.ToPushSettingsResponseOutputWithContext(context.Background())
}

func (i PushSettingsResponseArgs) ToPushSettingsResponseOutputWithContext(ctx context.Context) PushSettingsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsResponseOutput)
}

func (i PushSettingsResponseArgs) ToPushSettingsResponsePtrOutput() PushSettingsResponsePtrOutput {
	return i.ToPushSettingsResponsePtrOutputWithContext(context.Background())
}

func (i PushSettingsResponseArgs) ToPushSettingsResponsePtrOutputWithContext(ctx context.Context) PushSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsResponseOutput).ToPushSettingsResponsePtrOutputWithContext(ctx)
}









type PushSettingsResponsePtrInput interface {
	pulumi.Input

	ToPushSettingsResponsePtrOutput() PushSettingsResponsePtrOutput
	ToPushSettingsResponsePtrOutputWithContext(context.Context) PushSettingsResponsePtrOutput
}

type pushSettingsResponsePtrType PushSettingsResponseArgs

func PushSettingsResponsePtr(v *PushSettingsResponseArgs) PushSettingsResponsePtrInput {
	return (*pushSettingsResponsePtrType)(v)
}

func (*pushSettingsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**PushSettingsResponse)(nil)).Elem()
}

func (i *pushSettingsResponsePtrType) ToPushSettingsResponsePtrOutput() PushSettingsResponsePtrOutput {
	return i.ToPushSettingsResponsePtrOutputWithContext(context.Background())
}

func (i *pushSettingsResponsePtrType) ToPushSettingsResponsePtrOutputWithContext(ctx context.Context) PushSettingsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PushSettingsResponsePtrOutput)
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

func (o PushSettingsResponseOutput) ToPushSettingsResponsePtrOutput() PushSettingsResponsePtrOutput {
	return o.ToPushSettingsResponsePtrOutputWithContext(context.Background())
}

func (o PushSettingsResponseOutput) ToPushSettingsResponsePtrOutputWithContext(ctx context.Context) PushSettingsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v PushSettingsResponse) *PushSettingsResponse {
		return &v
	}).(PushSettingsResponsePtrOutput)
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





type RampUpRuleResponseInput interface {
	pulumi.Input

	ToRampUpRuleResponseOutput() RampUpRuleResponseOutput
	ToRampUpRuleResponseOutputWithContext(context.Context) RampUpRuleResponseOutput
}

type RampUpRuleResponseArgs struct {
	ActionHostName            pulumi.StringPtrInput  `pulumi:"actionHostName"`
	ChangeDecisionCallbackUrl pulumi.StringPtrInput  `pulumi:"changeDecisionCallbackUrl"`
	ChangeIntervalInMinutes   pulumi.IntPtrInput     `pulumi:"changeIntervalInMinutes"`
	ChangeStep                pulumi.Float64PtrInput `pulumi:"changeStep"`
	MaxReroutePercentage      pulumi.Float64PtrInput `pulumi:"maxReroutePercentage"`
	MinReroutePercentage      pulumi.Float64PtrInput `pulumi:"minReroutePercentage"`
	Name                      pulumi.StringPtrInput  `pulumi:"name"`
	ReroutePercentage         pulumi.Float64PtrInput `pulumi:"reroutePercentage"`
}

func (RampUpRuleResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RampUpRuleResponse)(nil)).Elem()
}

func (i RampUpRuleResponseArgs) ToRampUpRuleResponseOutput() RampUpRuleResponseOutput {
	return i.ToRampUpRuleResponseOutputWithContext(context.Background())
}

func (i RampUpRuleResponseArgs) ToRampUpRuleResponseOutputWithContext(ctx context.Context) RampUpRuleResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleResponseOutput)
}





type RampUpRuleResponseArrayInput interface {
	pulumi.Input

	ToRampUpRuleResponseArrayOutput() RampUpRuleResponseArrayOutput
	ToRampUpRuleResponseArrayOutputWithContext(context.Context) RampUpRuleResponseArrayOutput
}

type RampUpRuleResponseArray []RampUpRuleResponseInput

func (RampUpRuleResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]RampUpRuleResponse)(nil)).Elem()
}

func (i RampUpRuleResponseArray) ToRampUpRuleResponseArrayOutput() RampUpRuleResponseArrayOutput {
	return i.ToRampUpRuleResponseArrayOutputWithContext(context.Background())
}

func (i RampUpRuleResponseArray) ToRampUpRuleResponseArrayOutputWithContext(ctx context.Context) RampUpRuleResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RampUpRuleResponseArrayOutput)
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





type RequestsBasedTriggerResponseInput interface {
	pulumi.Input

	ToRequestsBasedTriggerResponseOutput() RequestsBasedTriggerResponseOutput
	ToRequestsBasedTriggerResponseOutputWithContext(context.Context) RequestsBasedTriggerResponseOutput
}

type RequestsBasedTriggerResponseArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
}

func (RequestsBasedTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*RequestsBasedTriggerResponse)(nil)).Elem()
}

func (i RequestsBasedTriggerResponseArgs) ToRequestsBasedTriggerResponseOutput() RequestsBasedTriggerResponseOutput {
	return i.ToRequestsBasedTriggerResponseOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerResponseArgs) ToRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) RequestsBasedTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerResponseOutput)
}

func (i RequestsBasedTriggerResponseArgs) ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput {
	return i.ToRequestsBasedTriggerResponsePtrOutputWithContext(context.Background())
}

func (i RequestsBasedTriggerResponseArgs) ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) RequestsBasedTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerResponseOutput).ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx)
}









type RequestsBasedTriggerResponsePtrInput interface {
	pulumi.Input

	ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput
	ToRequestsBasedTriggerResponsePtrOutputWithContext(context.Context) RequestsBasedTriggerResponsePtrOutput
}

type requestsBasedTriggerResponsePtrType RequestsBasedTriggerResponseArgs

func RequestsBasedTriggerResponsePtr(v *RequestsBasedTriggerResponseArgs) RequestsBasedTriggerResponsePtrInput {
	return (*requestsBasedTriggerResponsePtrType)(v)
}

func (*requestsBasedTriggerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**RequestsBasedTriggerResponse)(nil)).Elem()
}

func (i *requestsBasedTriggerResponsePtrType) ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput {
	return i.ToRequestsBasedTriggerResponsePtrOutputWithContext(context.Background())
}

func (i *requestsBasedTriggerResponsePtrType) ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) RequestsBasedTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(RequestsBasedTriggerResponsePtrOutput)
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

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponsePtrOutput() RequestsBasedTriggerResponsePtrOutput {
	return o.ToRequestsBasedTriggerResponsePtrOutputWithContext(context.Background())
}

func (o RequestsBasedTriggerResponseOutput) ToRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) RequestsBasedTriggerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v RequestsBasedTriggerResponse) *RequestsBasedTriggerResponse {
		return &v
	}).(RequestsBasedTriggerResponsePtrOutput)
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

type SiteConfig struct {
	AcrUseManagedIdentityCreds       *bool                   `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID         *string                 `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                         *bool                   `pulumi:"alwaysOn"`
	ApiDefinition                    *ApiDefinitionInfo      `pulumi:"apiDefinition"`
	ApiManagementConfig              *ApiManagementConfig    `pulumi:"apiManagementConfig"`
	AppCommandLine                   *string                 `pulumi:"appCommandLine"`
	AppSettings                      []NameValuePair         `pulumi:"appSettings"`
	AutoHealEnabled                  *bool                   `pulumi:"autoHealEnabled"`
	AutoHealRules                    *AutoHealRules          `pulumi:"autoHealRules"`
	AutoSwapSlotName                 *string                 `pulumi:"autoSwapSlotName"`
	ConnectionStrings                []ConnStringInfo        `pulumi:"connectionStrings"`
	Cors                             *CorsSettings           `pulumi:"cors"`
	DefaultDocuments                 []string                `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled      *bool                   `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                     *string                 `pulumi:"documentRoot"`
	Experiments                      *Experiments            `pulumi:"experiments"`
	FtpsState                        *string                 `pulumi:"ftpsState"`
	HandlerMappings                  []HandlerMapping        `pulumi:"handlerMappings"`
	HealthCheckPath                  *string                 `pulumi:"healthCheckPath"`
	Http20Enabled                    *bool                   `pulumi:"http20Enabled"`
	HttpLoggingEnabled               *bool                   `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions           []IpSecurityRestriction `pulumi:"ipSecurityRestrictions"`
	JavaContainer                    *string                 `pulumi:"javaContainer"`
	JavaContainerVersion             *string                 `pulumi:"javaContainerVersion"`
	JavaVersion                      *string                 `pulumi:"javaVersion"`
	Limits                           *SiteLimits             `pulumi:"limits"`
	LinuxFxVersion                   *string                 `pulumi:"linuxFxVersion"`
	LoadBalancing                    *SiteLoadBalancing      `pulumi:"loadBalancing"`
	LocalMySqlEnabled                *bool                   `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit           *int                    `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode              *ManagedPipelineMode    `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId         *int                    `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                    *string                 `pulumi:"minTlsVersion"`
	NetFrameworkVersion              *string                 `pulumi:"netFrameworkVersion"`
	NodeVersion                      *string                 `pulumi:"nodeVersion"`
	NumberOfWorkers                  *int                    `pulumi:"numberOfWorkers"`
	PhpVersion                       *string                 `pulumi:"phpVersion"`
	PowerShellVersion                *string                 `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount           *int                    `pulumi:"preWarmedInstanceCount"`
	PublishingUsername               *string                 `pulumi:"publishingUsername"`
	Push                             *PushSettings           `pulumi:"push"`
	PythonVersion                    *string                 `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled           *bool                   `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion           *string                 `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled            *bool                   `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime     *string                 `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions        []IpSecurityRestriction `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain *bool                   `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmType                          *string                 `pulumi:"scmType"`
	TracingOptions                   *string                 `pulumi:"tracingOptions"`
	Use32BitWorkerProcess            *bool                   `pulumi:"use32BitWorkerProcess"`
	VirtualApplications              []VirtualApplication    `pulumi:"virtualApplications"`
	VnetName                         *string                 `pulumi:"vnetName"`
	WebSocketsEnabled                *bool                   `pulumi:"webSocketsEnabled"`
	WindowsFxVersion                 *string                 `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId        *int                    `pulumi:"xManagedServiceIdentityId"`
}





type SiteConfigInput interface {
	pulumi.Input

	ToSiteConfigOutput() SiteConfigOutput
	ToSiteConfigOutputWithContext(context.Context) SiteConfigOutput
}

type SiteConfigArgs struct {
	AcrUseManagedIdentityCreds       pulumi.BoolPtrInput             `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID         pulumi.StringPtrInput           `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                         pulumi.BoolPtrInput             `pulumi:"alwaysOn"`
	ApiDefinition                    ApiDefinitionInfoPtrInput       `pulumi:"apiDefinition"`
	ApiManagementConfig              ApiManagementConfigPtrInput     `pulumi:"apiManagementConfig"`
	AppCommandLine                   pulumi.StringPtrInput           `pulumi:"appCommandLine"`
	AppSettings                      NameValuePairArrayInput         `pulumi:"appSettings"`
	AutoHealEnabled                  pulumi.BoolPtrInput             `pulumi:"autoHealEnabled"`
	AutoHealRules                    AutoHealRulesPtrInput           `pulumi:"autoHealRules"`
	AutoSwapSlotName                 pulumi.StringPtrInput           `pulumi:"autoSwapSlotName"`
	ConnectionStrings                ConnStringInfoArrayInput        `pulumi:"connectionStrings"`
	Cors                             CorsSettingsPtrInput            `pulumi:"cors"`
	DefaultDocuments                 pulumi.StringArrayInput         `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled      pulumi.BoolPtrInput             `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                     pulumi.StringPtrInput           `pulumi:"documentRoot"`
	Experiments                      ExperimentsPtrInput             `pulumi:"experiments"`
	FtpsState                        pulumi.StringPtrInput           `pulumi:"ftpsState"`
	HandlerMappings                  HandlerMappingArrayInput        `pulumi:"handlerMappings"`
	HealthCheckPath                  pulumi.StringPtrInput           `pulumi:"healthCheckPath"`
	Http20Enabled                    pulumi.BoolPtrInput             `pulumi:"http20Enabled"`
	HttpLoggingEnabled               pulumi.BoolPtrInput             `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions           IpSecurityRestrictionArrayInput `pulumi:"ipSecurityRestrictions"`
	JavaContainer                    pulumi.StringPtrInput           `pulumi:"javaContainer"`
	JavaContainerVersion             pulumi.StringPtrInput           `pulumi:"javaContainerVersion"`
	JavaVersion                      pulumi.StringPtrInput           `pulumi:"javaVersion"`
	Limits                           SiteLimitsPtrInput              `pulumi:"limits"`
	LinuxFxVersion                   pulumi.StringPtrInput           `pulumi:"linuxFxVersion"`
	LoadBalancing                    SiteLoadBalancingPtrInput       `pulumi:"loadBalancing"`
	LocalMySqlEnabled                pulumi.BoolPtrInput             `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit           pulumi.IntPtrInput              `pulumi:"logsDirectorySizeLimit"`
	ManagedPipelineMode              ManagedPipelineModePtrInput     `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId         pulumi.IntPtrInput              `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                    pulumi.StringPtrInput           `pulumi:"minTlsVersion"`
	NetFrameworkVersion              pulumi.StringPtrInput           `pulumi:"netFrameworkVersion"`
	NodeVersion                      pulumi.StringPtrInput           `pulumi:"nodeVersion"`
	NumberOfWorkers                  pulumi.IntPtrInput              `pulumi:"numberOfWorkers"`
	PhpVersion                       pulumi.StringPtrInput           `pulumi:"phpVersion"`
	PowerShellVersion                pulumi.StringPtrInput           `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount           pulumi.IntPtrInput              `pulumi:"preWarmedInstanceCount"`
	PublishingUsername               pulumi.StringPtrInput           `pulumi:"publishingUsername"`
	Push                             PushSettingsPtrInput            `pulumi:"push"`
	PythonVersion                    pulumi.StringPtrInput           `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled           pulumi.BoolPtrInput             `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion           pulumi.StringPtrInput           `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled            pulumi.BoolPtrInput             `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime     pulumi.StringPtrInput           `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions        IpSecurityRestrictionArrayInput `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain pulumi.BoolPtrInput             `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmType                          pulumi.StringPtrInput           `pulumi:"scmType"`
	TracingOptions                   pulumi.StringPtrInput           `pulumi:"tracingOptions"`
	Use32BitWorkerProcess            pulumi.BoolPtrInput             `pulumi:"use32BitWorkerProcess"`
	VirtualApplications              VirtualApplicationArrayInput    `pulumi:"virtualApplications"`
	VnetName                         pulumi.StringPtrInput           `pulumi:"vnetName"`
	WebSocketsEnabled                pulumi.BoolPtrInput             `pulumi:"webSocketsEnabled"`
	WindowsFxVersion                 pulumi.StringPtrInput           `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId        pulumi.IntPtrInput              `pulumi:"xManagedServiceIdentityId"`
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

func (o SiteConfigOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfig) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
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

func (o SiteConfigPtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfig) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
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
	AcrUseManagedIdentityCreds       *bool                           `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID         *string                         `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                         *bool                           `pulumi:"alwaysOn"`
	ApiDefinition                    *ApiDefinitionInfoResponse      `pulumi:"apiDefinition"`
	ApiManagementConfig              *ApiManagementConfigResponse    `pulumi:"apiManagementConfig"`
	AppCommandLine                   *string                         `pulumi:"appCommandLine"`
	AppSettings                      []NameValuePairResponse         `pulumi:"appSettings"`
	AutoHealEnabled                  *bool                           `pulumi:"autoHealEnabled"`
	AutoHealRules                    *AutoHealRulesResponse          `pulumi:"autoHealRules"`
	AutoSwapSlotName                 *string                         `pulumi:"autoSwapSlotName"`
	ConnectionStrings                []ConnStringInfoResponse        `pulumi:"connectionStrings"`
	Cors                             *CorsSettingsResponse           `pulumi:"cors"`
	DefaultDocuments                 []string                        `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled      *bool                           `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                     *string                         `pulumi:"documentRoot"`
	Experiments                      *ExperimentsResponse            `pulumi:"experiments"`
	FtpsState                        *string                         `pulumi:"ftpsState"`
	HandlerMappings                  []HandlerMappingResponse        `pulumi:"handlerMappings"`
	HealthCheckPath                  *string                         `pulumi:"healthCheckPath"`
	Http20Enabled                    *bool                           `pulumi:"http20Enabled"`
	HttpLoggingEnabled               *bool                           `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions           []IpSecurityRestrictionResponse `pulumi:"ipSecurityRestrictions"`
	JavaContainer                    *string                         `pulumi:"javaContainer"`
	JavaContainerVersion             *string                         `pulumi:"javaContainerVersion"`
	JavaVersion                      *string                         `pulumi:"javaVersion"`
	Limits                           *SiteLimitsResponse             `pulumi:"limits"`
	LinuxFxVersion                   *string                         `pulumi:"linuxFxVersion"`
	LoadBalancing                    *string                         `pulumi:"loadBalancing"`
	LocalMySqlEnabled                *bool                           `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit           *int                            `pulumi:"logsDirectorySizeLimit"`
	MachineKey                       SiteMachineKeyResponse          `pulumi:"machineKey"`
	ManagedPipelineMode              *string                         `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId         *int                            `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                    *string                         `pulumi:"minTlsVersion"`
	NetFrameworkVersion              *string                         `pulumi:"netFrameworkVersion"`
	NodeVersion                      *string                         `pulumi:"nodeVersion"`
	NumberOfWorkers                  *int                            `pulumi:"numberOfWorkers"`
	PhpVersion                       *string                         `pulumi:"phpVersion"`
	PowerShellVersion                *string                         `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount           *int                            `pulumi:"preWarmedInstanceCount"`
	PublishingUsername               *string                         `pulumi:"publishingUsername"`
	Push                             *PushSettingsResponse           `pulumi:"push"`
	PythonVersion                    *string                         `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled           *bool                           `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion           *string                         `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled            *bool                           `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime     *string                         `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions        []IpSecurityRestrictionResponse `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain *bool                           `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmType                          *string                         `pulumi:"scmType"`
	TracingOptions                   *string                         `pulumi:"tracingOptions"`
	Use32BitWorkerProcess            *bool                           `pulumi:"use32BitWorkerProcess"`
	VirtualApplications              []VirtualApplicationResponse    `pulumi:"virtualApplications"`
	VnetName                         *string                         `pulumi:"vnetName"`
	WebSocketsEnabled                *bool                           `pulumi:"webSocketsEnabled"`
	WindowsFxVersion                 *string                         `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId        *int                            `pulumi:"xManagedServiceIdentityId"`
}





type SiteConfigResponseInput interface {
	pulumi.Input

	ToSiteConfigResponseOutput() SiteConfigResponseOutput
	ToSiteConfigResponseOutputWithContext(context.Context) SiteConfigResponseOutput
}

type SiteConfigResponseArgs struct {
	AcrUseManagedIdentityCreds       pulumi.BoolPtrInput                     `pulumi:"acrUseManagedIdentityCreds"`
	AcrUserManagedIdentityID         pulumi.StringPtrInput                   `pulumi:"acrUserManagedIdentityID"`
	AlwaysOn                         pulumi.BoolPtrInput                     `pulumi:"alwaysOn"`
	ApiDefinition                    ApiDefinitionInfoResponsePtrInput       `pulumi:"apiDefinition"`
	ApiManagementConfig              ApiManagementConfigResponsePtrInput     `pulumi:"apiManagementConfig"`
	AppCommandLine                   pulumi.StringPtrInput                   `pulumi:"appCommandLine"`
	AppSettings                      NameValuePairResponseArrayInput         `pulumi:"appSettings"`
	AutoHealEnabled                  pulumi.BoolPtrInput                     `pulumi:"autoHealEnabled"`
	AutoHealRules                    AutoHealRulesResponsePtrInput           `pulumi:"autoHealRules"`
	AutoSwapSlotName                 pulumi.StringPtrInput                   `pulumi:"autoSwapSlotName"`
	ConnectionStrings                ConnStringInfoResponseArrayInput        `pulumi:"connectionStrings"`
	Cors                             CorsSettingsResponsePtrInput            `pulumi:"cors"`
	DefaultDocuments                 pulumi.StringArrayInput                 `pulumi:"defaultDocuments"`
	DetailedErrorLoggingEnabled      pulumi.BoolPtrInput                     `pulumi:"detailedErrorLoggingEnabled"`
	DocumentRoot                     pulumi.StringPtrInput                   `pulumi:"documentRoot"`
	Experiments                      ExperimentsResponsePtrInput             `pulumi:"experiments"`
	FtpsState                        pulumi.StringPtrInput                   `pulumi:"ftpsState"`
	HandlerMappings                  HandlerMappingResponseArrayInput        `pulumi:"handlerMappings"`
	HealthCheckPath                  pulumi.StringPtrInput                   `pulumi:"healthCheckPath"`
	Http20Enabled                    pulumi.BoolPtrInput                     `pulumi:"http20Enabled"`
	HttpLoggingEnabled               pulumi.BoolPtrInput                     `pulumi:"httpLoggingEnabled"`
	IpSecurityRestrictions           IpSecurityRestrictionResponseArrayInput `pulumi:"ipSecurityRestrictions"`
	JavaContainer                    pulumi.StringPtrInput                   `pulumi:"javaContainer"`
	JavaContainerVersion             pulumi.StringPtrInput                   `pulumi:"javaContainerVersion"`
	JavaVersion                      pulumi.StringPtrInput                   `pulumi:"javaVersion"`
	Limits                           SiteLimitsResponsePtrInput              `pulumi:"limits"`
	LinuxFxVersion                   pulumi.StringPtrInput                   `pulumi:"linuxFxVersion"`
	LoadBalancing                    pulumi.StringPtrInput                   `pulumi:"loadBalancing"`
	LocalMySqlEnabled                pulumi.BoolPtrInput                     `pulumi:"localMySqlEnabled"`
	LogsDirectorySizeLimit           pulumi.IntPtrInput                      `pulumi:"logsDirectorySizeLimit"`
	MachineKey                       SiteMachineKeyResponseInput             `pulumi:"machineKey"`
	ManagedPipelineMode              pulumi.StringPtrInput                   `pulumi:"managedPipelineMode"`
	ManagedServiceIdentityId         pulumi.IntPtrInput                      `pulumi:"managedServiceIdentityId"`
	MinTlsVersion                    pulumi.StringPtrInput                   `pulumi:"minTlsVersion"`
	NetFrameworkVersion              pulumi.StringPtrInput                   `pulumi:"netFrameworkVersion"`
	NodeVersion                      pulumi.StringPtrInput                   `pulumi:"nodeVersion"`
	NumberOfWorkers                  pulumi.IntPtrInput                      `pulumi:"numberOfWorkers"`
	PhpVersion                       pulumi.StringPtrInput                   `pulumi:"phpVersion"`
	PowerShellVersion                pulumi.StringPtrInput                   `pulumi:"powerShellVersion"`
	PreWarmedInstanceCount           pulumi.IntPtrInput                      `pulumi:"preWarmedInstanceCount"`
	PublishingUsername               pulumi.StringPtrInput                   `pulumi:"publishingUsername"`
	Push                             PushSettingsResponsePtrInput            `pulumi:"push"`
	PythonVersion                    pulumi.StringPtrInput                   `pulumi:"pythonVersion"`
	RemoteDebuggingEnabled           pulumi.BoolPtrInput                     `pulumi:"remoteDebuggingEnabled"`
	RemoteDebuggingVersion           pulumi.StringPtrInput                   `pulumi:"remoteDebuggingVersion"`
	RequestTracingEnabled            pulumi.BoolPtrInput                     `pulumi:"requestTracingEnabled"`
	RequestTracingExpirationTime     pulumi.StringPtrInput                   `pulumi:"requestTracingExpirationTime"`
	ScmIpSecurityRestrictions        IpSecurityRestrictionResponseArrayInput `pulumi:"scmIpSecurityRestrictions"`
	ScmIpSecurityRestrictionsUseMain pulumi.BoolPtrInput                     `pulumi:"scmIpSecurityRestrictionsUseMain"`
	ScmType                          pulumi.StringPtrInput                   `pulumi:"scmType"`
	TracingOptions                   pulumi.StringPtrInput                   `pulumi:"tracingOptions"`
	Use32BitWorkerProcess            pulumi.BoolPtrInput                     `pulumi:"use32BitWorkerProcess"`
	VirtualApplications              VirtualApplicationResponseArrayInput    `pulumi:"virtualApplications"`
	VnetName                         pulumi.StringPtrInput                   `pulumi:"vnetName"`
	WebSocketsEnabled                pulumi.BoolPtrInput                     `pulumi:"webSocketsEnabled"`
	WindowsFxVersion                 pulumi.StringPtrInput                   `pulumi:"windowsFxVersion"`
	XManagedServiceIdentityId        pulumi.IntPtrInput                      `pulumi:"xManagedServiceIdentityId"`
}

func (SiteConfigResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteConfigResponse)(nil)).Elem()
}

func (i SiteConfigResponseArgs) ToSiteConfigResponseOutput() SiteConfigResponseOutput {
	return i.ToSiteConfigResponseOutputWithContext(context.Background())
}

func (i SiteConfigResponseArgs) ToSiteConfigResponseOutputWithContext(ctx context.Context) SiteConfigResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigResponseOutput)
}

func (i SiteConfigResponseArgs) ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput {
	return i.ToSiteConfigResponsePtrOutputWithContext(context.Background())
}

func (i SiteConfigResponseArgs) ToSiteConfigResponsePtrOutputWithContext(ctx context.Context) SiteConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigResponseOutput).ToSiteConfigResponsePtrOutputWithContext(ctx)
}









type SiteConfigResponsePtrInput interface {
	pulumi.Input

	ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput
	ToSiteConfigResponsePtrOutputWithContext(context.Context) SiteConfigResponsePtrOutput
}

type siteConfigResponsePtrType SiteConfigResponseArgs

func SiteConfigResponsePtr(v *SiteConfigResponseArgs) SiteConfigResponsePtrInput {
	return (*siteConfigResponsePtrType)(v)
}

func (*siteConfigResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteConfigResponse)(nil)).Elem()
}

func (i *siteConfigResponsePtrType) ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput {
	return i.ToSiteConfigResponsePtrOutputWithContext(context.Background())
}

func (i *siteConfigResponsePtrType) ToSiteConfigResponsePtrOutputWithContext(ctx context.Context) SiteConfigResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteConfigResponsePtrOutput)
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

func (o SiteConfigResponseOutput) ToSiteConfigResponsePtrOutput() SiteConfigResponsePtrOutput {
	return o.ToSiteConfigResponsePtrOutputWithContext(context.Background())
}

func (o SiteConfigResponseOutput) ToSiteConfigResponsePtrOutputWithContext(ctx context.Context) SiteConfigResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteConfigResponse) *SiteConfigResponse {
		return &v
	}).(SiteConfigResponsePtrOutput)
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

func (o SiteConfigResponseOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteConfigResponse) *bool { return v.WebSocketsEnabled }).(pulumi.BoolPtrOutput)
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

func (o SiteConfigResponsePtrOutput) WebSocketsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteConfigResponse) *bool {
		if v == nil {
			return nil
		}
		return v.WebSocketsEnabled
	}).(pulumi.BoolPtrOutput)
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





type SiteLimitsResponseInput interface {
	pulumi.Input

	ToSiteLimitsResponseOutput() SiteLimitsResponseOutput
	ToSiteLimitsResponseOutputWithContext(context.Context) SiteLimitsResponseOutput
}

type SiteLimitsResponseArgs struct {
	MaxDiskSizeInMb  pulumi.Float64PtrInput `pulumi:"maxDiskSizeInMb"`
	MaxMemoryInMb    pulumi.Float64PtrInput `pulumi:"maxMemoryInMb"`
	MaxPercentageCpu pulumi.Float64PtrInput `pulumi:"maxPercentageCpu"`
}

func (SiteLimitsResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteLimitsResponse)(nil)).Elem()
}

func (i SiteLimitsResponseArgs) ToSiteLimitsResponseOutput() SiteLimitsResponseOutput {
	return i.ToSiteLimitsResponseOutputWithContext(context.Background())
}

func (i SiteLimitsResponseArgs) ToSiteLimitsResponseOutputWithContext(ctx context.Context) SiteLimitsResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsResponseOutput)
}

func (i SiteLimitsResponseArgs) ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput {
	return i.ToSiteLimitsResponsePtrOutputWithContext(context.Background())
}

func (i SiteLimitsResponseArgs) ToSiteLimitsResponsePtrOutputWithContext(ctx context.Context) SiteLimitsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsResponseOutput).ToSiteLimitsResponsePtrOutputWithContext(ctx)
}









type SiteLimitsResponsePtrInput interface {
	pulumi.Input

	ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput
	ToSiteLimitsResponsePtrOutputWithContext(context.Context) SiteLimitsResponsePtrOutput
}

type siteLimitsResponsePtrType SiteLimitsResponseArgs

func SiteLimitsResponsePtr(v *SiteLimitsResponseArgs) SiteLimitsResponsePtrInput {
	return (*siteLimitsResponsePtrType)(v)
}

func (*siteLimitsResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteLimitsResponse)(nil)).Elem()
}

func (i *siteLimitsResponsePtrType) ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput {
	return i.ToSiteLimitsResponsePtrOutputWithContext(context.Background())
}

func (i *siteLimitsResponsePtrType) ToSiteLimitsResponsePtrOutputWithContext(ctx context.Context) SiteLimitsResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteLimitsResponsePtrOutput)
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

func (o SiteLimitsResponseOutput) ToSiteLimitsResponsePtrOutput() SiteLimitsResponsePtrOutput {
	return o.ToSiteLimitsResponsePtrOutputWithContext(context.Background())
}

func (o SiteLimitsResponseOutput) ToSiteLimitsResponsePtrOutputWithContext(ctx context.Context) SiteLimitsResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteLimitsResponse) *SiteLimitsResponse {
		return &v
	}).(SiteLimitsResponsePtrOutput)
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





type SiteMachineKeyResponseInput interface {
	pulumi.Input

	ToSiteMachineKeyResponseOutput() SiteMachineKeyResponseOutput
	ToSiteMachineKeyResponseOutputWithContext(context.Context) SiteMachineKeyResponseOutput
}

type SiteMachineKeyResponseArgs struct {
	Decryption    pulumi.StringPtrInput `pulumi:"decryption"`
	DecryptionKey pulumi.StringPtrInput `pulumi:"decryptionKey"`
	Validation    pulumi.StringPtrInput `pulumi:"validation"`
	ValidationKey pulumi.StringPtrInput `pulumi:"validationKey"`
}

func (SiteMachineKeyResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteMachineKeyResponse)(nil)).Elem()
}

func (i SiteMachineKeyResponseArgs) ToSiteMachineKeyResponseOutput() SiteMachineKeyResponseOutput {
	return i.ToSiteMachineKeyResponseOutputWithContext(context.Background())
}

func (i SiteMachineKeyResponseArgs) ToSiteMachineKeyResponseOutputWithContext(ctx context.Context) SiteMachineKeyResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMachineKeyResponseOutput)
}

func (i SiteMachineKeyResponseArgs) ToSiteMachineKeyResponsePtrOutput() SiteMachineKeyResponsePtrOutput {
	return i.ToSiteMachineKeyResponsePtrOutputWithContext(context.Background())
}

func (i SiteMachineKeyResponseArgs) ToSiteMachineKeyResponsePtrOutputWithContext(ctx context.Context) SiteMachineKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMachineKeyResponseOutput).ToSiteMachineKeyResponsePtrOutputWithContext(ctx)
}









type SiteMachineKeyResponsePtrInput interface {
	pulumi.Input

	ToSiteMachineKeyResponsePtrOutput() SiteMachineKeyResponsePtrOutput
	ToSiteMachineKeyResponsePtrOutputWithContext(context.Context) SiteMachineKeyResponsePtrOutput
}

type siteMachineKeyResponsePtrType SiteMachineKeyResponseArgs

func SiteMachineKeyResponsePtr(v *SiteMachineKeyResponseArgs) SiteMachineKeyResponsePtrInput {
	return (*siteMachineKeyResponsePtrType)(v)
}

func (*siteMachineKeyResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteMachineKeyResponse)(nil)).Elem()
}

func (i *siteMachineKeyResponsePtrType) ToSiteMachineKeyResponsePtrOutput() SiteMachineKeyResponsePtrOutput {
	return i.ToSiteMachineKeyResponsePtrOutputWithContext(context.Background())
}

func (i *siteMachineKeyResponsePtrType) ToSiteMachineKeyResponsePtrOutputWithContext(ctx context.Context) SiteMachineKeyResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteMachineKeyResponsePtrOutput)
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

func (o SiteMachineKeyResponseOutput) ToSiteMachineKeyResponsePtrOutput() SiteMachineKeyResponsePtrOutput {
	return o.ToSiteMachineKeyResponsePtrOutputWithContext(context.Background())
}

func (o SiteMachineKeyResponseOutput) ToSiteMachineKeyResponsePtrOutputWithContext(ctx context.Context) SiteMachineKeyResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SiteMachineKeyResponse) *SiteMachineKeyResponse {
		return &v
	}).(SiteMachineKeyResponsePtrOutput)
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
	Default   *int    `pulumi:"default"`
	Maximum   *int    `pulumi:"maximum"`
	Minimum   *int    `pulumi:"minimum"`
	ScaleType *string `pulumi:"scaleType"`
}





type SkuCapacityInput interface {
	pulumi.Input

	ToSkuCapacityOutput() SkuCapacityOutput
	ToSkuCapacityOutputWithContext(context.Context) SkuCapacityOutput
}

type SkuCapacityArgs struct {
	Default   pulumi.IntPtrInput    `pulumi:"default"`
	Maximum   pulumi.IntPtrInput    `pulumi:"maximum"`
	Minimum   pulumi.IntPtrInput    `pulumi:"minimum"`
	ScaleType pulumi.StringPtrInput `pulumi:"scaleType"`
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
	Default   *int    `pulumi:"default"`
	Maximum   *int    `pulumi:"maximum"`
	Minimum   *int    `pulumi:"minimum"`
	ScaleType *string `pulumi:"scaleType"`
}





type SkuCapacityResponseInput interface {
	pulumi.Input

	ToSkuCapacityResponseOutput() SkuCapacityResponseOutput
	ToSkuCapacityResponseOutputWithContext(context.Context) SkuCapacityResponseOutput
}

type SkuCapacityResponseArgs struct {
	Default   pulumi.IntPtrInput    `pulumi:"default"`
	Maximum   pulumi.IntPtrInput    `pulumi:"maximum"`
	Minimum   pulumi.IntPtrInput    `pulumi:"minimum"`
	ScaleType pulumi.StringPtrInput `pulumi:"scaleType"`
}

func (SkuCapacityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuCapacityResponse)(nil)).Elem()
}

func (i SkuCapacityResponseArgs) ToSkuCapacityResponseOutput() SkuCapacityResponseOutput {
	return i.ToSkuCapacityResponseOutputWithContext(context.Background())
}

func (i SkuCapacityResponseArgs) ToSkuCapacityResponseOutputWithContext(ctx context.Context) SkuCapacityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityResponseOutput)
}

func (i SkuCapacityResponseArgs) ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput {
	return i.ToSkuCapacityResponsePtrOutputWithContext(context.Background())
}

func (i SkuCapacityResponseArgs) ToSkuCapacityResponsePtrOutputWithContext(ctx context.Context) SkuCapacityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityResponseOutput).ToSkuCapacityResponsePtrOutputWithContext(ctx)
}









type SkuCapacityResponsePtrInput interface {
	pulumi.Input

	ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput
	ToSkuCapacityResponsePtrOutputWithContext(context.Context) SkuCapacityResponsePtrOutput
}

type skuCapacityResponsePtrType SkuCapacityResponseArgs

func SkuCapacityResponsePtr(v *SkuCapacityResponseArgs) SkuCapacityResponsePtrInput {
	return (*skuCapacityResponsePtrType)(v)
}

func (*skuCapacityResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuCapacityResponse)(nil)).Elem()
}

func (i *skuCapacityResponsePtrType) ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput {
	return i.ToSkuCapacityResponsePtrOutputWithContext(context.Background())
}

func (i *skuCapacityResponsePtrType) ToSkuCapacityResponsePtrOutputWithContext(ctx context.Context) SkuCapacityResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuCapacityResponsePtrOutput)
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

func (o SkuCapacityResponseOutput) ToSkuCapacityResponsePtrOutput() SkuCapacityResponsePtrOutput {
	return o.ToSkuCapacityResponsePtrOutputWithContext(context.Background())
}

func (o SkuCapacityResponseOutput) ToSkuCapacityResponsePtrOutputWithContext(ctx context.Context) SkuCapacityResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuCapacityResponse) *SkuCapacityResponse {
		return &v
	}).(SkuCapacityResponsePtrOutput)
}

func (o SkuCapacityResponseOutput) Default() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SkuCapacityResponse) *int { return v.Default }).(pulumi.IntPtrOutput)
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





type SkuDescriptionResponseInput interface {
	pulumi.Input

	ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput
	ToSkuDescriptionResponseOutputWithContext(context.Context) SkuDescriptionResponseOutput
}

type SkuDescriptionResponseArgs struct {
	Capabilities CapabilityResponseArrayInput `pulumi:"capabilities"`
	Capacity     pulumi.IntPtrInput           `pulumi:"capacity"`
	Family       pulumi.StringPtrInput        `pulumi:"family"`
	Locations    pulumi.StringArrayInput      `pulumi:"locations"`
	Name         pulumi.StringPtrInput        `pulumi:"name"`
	Size         pulumi.StringPtrInput        `pulumi:"size"`
	SkuCapacity  SkuCapacityResponsePtrInput  `pulumi:"skuCapacity"`
	Tier         pulumi.StringPtrInput        `pulumi:"tier"`
}

func (SkuDescriptionResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuDescriptionResponse)(nil)).Elem()
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponseOutput() SkuDescriptionResponseOutput {
	return i.ToSkuDescriptionResponseOutputWithContext(context.Background())
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponseOutputWithContext(ctx context.Context) SkuDescriptionResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionResponseOutput)
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return i.ToSkuDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i SkuDescriptionResponseArgs) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionResponseOutput).ToSkuDescriptionResponsePtrOutputWithContext(ctx)
}









type SkuDescriptionResponsePtrInput interface {
	pulumi.Input

	ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput
	ToSkuDescriptionResponsePtrOutputWithContext(context.Context) SkuDescriptionResponsePtrOutput
}

type skuDescriptionResponsePtrType SkuDescriptionResponseArgs

func SkuDescriptionResponsePtr(v *SkuDescriptionResponseArgs) SkuDescriptionResponsePtrInput {
	return (*skuDescriptionResponsePtrType)(v)
}

func (*skuDescriptionResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuDescriptionResponse)(nil)).Elem()
}

func (i *skuDescriptionResponsePtrType) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return i.ToSkuDescriptionResponsePtrOutputWithContext(context.Background())
}

func (i *skuDescriptionResponsePtrType) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuDescriptionResponsePtrOutput)
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

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponsePtrOutput() SkuDescriptionResponsePtrOutput {
	return o.ToSkuDescriptionResponsePtrOutputWithContext(context.Background())
}

func (o SkuDescriptionResponseOutput) ToSkuDescriptionResponsePtrOutputWithContext(ctx context.Context) SkuDescriptionResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuDescriptionResponse) *SkuDescriptionResponse {
		return &v
	}).(SkuDescriptionResponsePtrOutput)
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





type SlotSwapStatusResponseInput interface {
	pulumi.Input

	ToSlotSwapStatusResponseOutput() SlotSwapStatusResponseOutput
	ToSlotSwapStatusResponseOutputWithContext(context.Context) SlotSwapStatusResponseOutput
}

type SlotSwapStatusResponseArgs struct {
	DestinationSlotName pulumi.StringInput `pulumi:"destinationSlotName"`
	SourceSlotName      pulumi.StringInput `pulumi:"sourceSlotName"`
	TimestampUtc        pulumi.StringInput `pulumi:"timestampUtc"`
}

func (SlotSwapStatusResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlotSwapStatusResponse)(nil)).Elem()
}

func (i SlotSwapStatusResponseArgs) ToSlotSwapStatusResponseOutput() SlotSwapStatusResponseOutput {
	return i.ToSlotSwapStatusResponseOutputWithContext(context.Background())
}

func (i SlotSwapStatusResponseArgs) ToSlotSwapStatusResponseOutputWithContext(ctx context.Context) SlotSwapStatusResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotSwapStatusResponseOutput)
}

func (i SlotSwapStatusResponseArgs) ToSlotSwapStatusResponsePtrOutput() SlotSwapStatusResponsePtrOutput {
	return i.ToSlotSwapStatusResponsePtrOutputWithContext(context.Background())
}

func (i SlotSwapStatusResponseArgs) ToSlotSwapStatusResponsePtrOutputWithContext(ctx context.Context) SlotSwapStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotSwapStatusResponseOutput).ToSlotSwapStatusResponsePtrOutputWithContext(ctx)
}









type SlotSwapStatusResponsePtrInput interface {
	pulumi.Input

	ToSlotSwapStatusResponsePtrOutput() SlotSwapStatusResponsePtrOutput
	ToSlotSwapStatusResponsePtrOutputWithContext(context.Context) SlotSwapStatusResponsePtrOutput
}

type slotSwapStatusResponsePtrType SlotSwapStatusResponseArgs

func SlotSwapStatusResponsePtr(v *SlotSwapStatusResponseArgs) SlotSwapStatusResponsePtrInput {
	return (*slotSwapStatusResponsePtrType)(v)
}

func (*slotSwapStatusResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlotSwapStatusResponse)(nil)).Elem()
}

func (i *slotSwapStatusResponsePtrType) ToSlotSwapStatusResponsePtrOutput() SlotSwapStatusResponsePtrOutput {
	return i.ToSlotSwapStatusResponsePtrOutputWithContext(context.Background())
}

func (i *slotSwapStatusResponsePtrType) ToSlotSwapStatusResponsePtrOutputWithContext(ctx context.Context) SlotSwapStatusResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlotSwapStatusResponsePtrOutput)
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

func (o SlotSwapStatusResponseOutput) ToSlotSwapStatusResponsePtrOutput() SlotSwapStatusResponsePtrOutput {
	return o.ToSlotSwapStatusResponsePtrOutputWithContext(context.Background())
}

func (o SlotSwapStatusResponseOutput) ToSlotSwapStatusResponsePtrOutputWithContext(ctx context.Context) SlotSwapStatusResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlotSwapStatusResponse) *SlotSwapStatusResponse {
		return &v
	}).(SlotSwapStatusResponsePtrOutput)
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

type SlotSwapStatusResponsePtrOutput struct{ *pulumi.OutputState }

func (SlotSwapStatusResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlotSwapStatusResponse)(nil)).Elem()
}

func (o SlotSwapStatusResponsePtrOutput) ToSlotSwapStatusResponsePtrOutput() SlotSwapStatusResponsePtrOutput {
	return o
}

func (o SlotSwapStatusResponsePtrOutput) ToSlotSwapStatusResponsePtrOutputWithContext(ctx context.Context) SlotSwapStatusResponsePtrOutput {
	return o
}

func (o SlotSwapStatusResponsePtrOutput) Elem() SlotSwapStatusResponseOutput {
	return o.ApplyT(func(v *SlotSwapStatusResponse) SlotSwapStatusResponse {
		if v != nil {
			return *v
		}
		var ret SlotSwapStatusResponse
		return ret
	}).(SlotSwapStatusResponseOutput)
}

func (o SlotSwapStatusResponsePtrOutput) DestinationSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlotSwapStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DestinationSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SlotSwapStatusResponsePtrOutput) SourceSlotName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlotSwapStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SourceSlotName
	}).(pulumi.StringPtrOutput)
}

func (o SlotSwapStatusResponsePtrOutput) TimestampUtc() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlotSwapStatusResponse) *string {
		if v == nil {
			return nil
		}
		return &v.TimestampUtc
	}).(pulumi.StringPtrOutput)
}

type SlowRequestsBasedTrigger struct {
	Count        *int    `pulumi:"count"`
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

type SlowRequestsBasedTriggerResponse struct {
	Count        *int    `pulumi:"count"`
	TimeInterval *string `pulumi:"timeInterval"`
	TimeTaken    *string `pulumi:"timeTaken"`
}





type SlowRequestsBasedTriggerResponseInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerResponseOutput() SlowRequestsBasedTriggerResponseOutput
	ToSlowRequestsBasedTriggerResponseOutputWithContext(context.Context) SlowRequestsBasedTriggerResponseOutput
}

type SlowRequestsBasedTriggerResponseArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	TimeTaken    pulumi.StringPtrInput `pulumi:"timeTaken"`
}

func (SlowRequestsBasedTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (i SlowRequestsBasedTriggerResponseArgs) ToSlowRequestsBasedTriggerResponseOutput() SlowRequestsBasedTriggerResponseOutput {
	return i.ToSlowRequestsBasedTriggerResponseOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerResponseArgs) ToSlowRequestsBasedTriggerResponseOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerResponseOutput)
}

func (i SlowRequestsBasedTriggerResponseArgs) ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput {
	return i.ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(context.Background())
}

func (i SlowRequestsBasedTriggerResponseArgs) ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerResponseOutput).ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx)
}









type SlowRequestsBasedTriggerResponsePtrInput interface {
	pulumi.Input

	ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput
	ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(context.Context) SlowRequestsBasedTriggerResponsePtrOutput
}

type slowRequestsBasedTriggerResponsePtrType SlowRequestsBasedTriggerResponseArgs

func SlowRequestsBasedTriggerResponsePtr(v *SlowRequestsBasedTriggerResponseArgs) SlowRequestsBasedTriggerResponsePtrInput {
	return (*slowRequestsBasedTriggerResponsePtrType)(v)
}

func (*slowRequestsBasedTriggerResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlowRequestsBasedTriggerResponse)(nil)).Elem()
}

func (i *slowRequestsBasedTriggerResponsePtrType) ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput {
	return i.ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(context.Background())
}

func (i *slowRequestsBasedTriggerResponsePtrType) ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlowRequestsBasedTriggerResponsePtrOutput)
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

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponsePtrOutput() SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(context.Background())
}

func (o SlowRequestsBasedTriggerResponseOutput) ToSlowRequestsBasedTriggerResponsePtrOutputWithContext(ctx context.Context) SlowRequestsBasedTriggerResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlowRequestsBasedTriggerResponse) *SlowRequestsBasedTriggerResponse {
		return &v
	}).(SlowRequestsBasedTriggerResponsePtrOutput)
}

func (o SlowRequestsBasedTriggerResponseOutput) Count() pulumi.IntPtrOutput {
	return o.ApplyT(func(v SlowRequestsBasedTriggerResponse) *int { return v.Count }).(pulumi.IntPtrOutput)
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

type StampCapacityResponse struct {
	AvailableCapacity              *float64 `pulumi:"availableCapacity"`
	ComputeMode                    *string  `pulumi:"computeMode"`
	ExcludeFromCapacityAllocation  *bool    `pulumi:"excludeFromCapacityAllocation"`
	IsApplicableForAllComputeModes *bool    `pulumi:"isApplicableForAllComputeModes"`
	IsLinux                        *bool    `pulumi:"isLinux"`
	Name                           *string  `pulumi:"name"`
	SiteMode                       *string  `pulumi:"siteMode"`
	TotalCapacity                  *float64 `pulumi:"totalCapacity"`
	Unit                           *string  `pulumi:"unit"`
	WorkerSize                     *string  `pulumi:"workerSize"`
	WorkerSizeId                   *int     `pulumi:"workerSizeId"`
}





type StampCapacityResponseInput interface {
	pulumi.Input

	ToStampCapacityResponseOutput() StampCapacityResponseOutput
	ToStampCapacityResponseOutputWithContext(context.Context) StampCapacityResponseOutput
}

type StampCapacityResponseArgs struct {
	AvailableCapacity              pulumi.Float64PtrInput `pulumi:"availableCapacity"`
	ComputeMode                    pulumi.StringPtrInput  `pulumi:"computeMode"`
	ExcludeFromCapacityAllocation  pulumi.BoolPtrInput    `pulumi:"excludeFromCapacityAllocation"`
	IsApplicableForAllComputeModes pulumi.BoolPtrInput    `pulumi:"isApplicableForAllComputeModes"`
	IsLinux                        pulumi.BoolPtrInput    `pulumi:"isLinux"`
	Name                           pulumi.StringPtrInput  `pulumi:"name"`
	SiteMode                       pulumi.StringPtrInput  `pulumi:"siteMode"`
	TotalCapacity                  pulumi.Float64PtrInput `pulumi:"totalCapacity"`
	Unit                           pulumi.StringPtrInput  `pulumi:"unit"`
	WorkerSize                     pulumi.StringPtrInput  `pulumi:"workerSize"`
	WorkerSizeId                   pulumi.IntPtrInput     `pulumi:"workerSizeId"`
}

func (StampCapacityResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StampCapacityResponse)(nil)).Elem()
}

func (i StampCapacityResponseArgs) ToStampCapacityResponseOutput() StampCapacityResponseOutput {
	return i.ToStampCapacityResponseOutputWithContext(context.Background())
}

func (i StampCapacityResponseArgs) ToStampCapacityResponseOutputWithContext(ctx context.Context) StampCapacityResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StampCapacityResponseOutput)
}





type StampCapacityResponseArrayInput interface {
	pulumi.Input

	ToStampCapacityResponseArrayOutput() StampCapacityResponseArrayOutput
	ToStampCapacityResponseArrayOutputWithContext(context.Context) StampCapacityResponseArrayOutput
}

type StampCapacityResponseArray []StampCapacityResponseInput

func (StampCapacityResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StampCapacityResponse)(nil)).Elem()
}

func (i StampCapacityResponseArray) ToStampCapacityResponseArrayOutput() StampCapacityResponseArrayOutput {
	return i.ToStampCapacityResponseArrayOutputWithContext(context.Background())
}

func (i StampCapacityResponseArray) ToStampCapacityResponseArrayOutputWithContext(ctx context.Context) StampCapacityResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StampCapacityResponseArrayOutput)
}

type StampCapacityResponseOutput struct{ *pulumi.OutputState }

func (StampCapacityResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*StampCapacityResponse)(nil)).Elem()
}

func (o StampCapacityResponseOutput) ToStampCapacityResponseOutput() StampCapacityResponseOutput {
	return o
}

func (o StampCapacityResponseOutput) ToStampCapacityResponseOutputWithContext(ctx context.Context) StampCapacityResponseOutput {
	return o
}

func (o StampCapacityResponseOutput) AvailableCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *float64 { return v.AvailableCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityResponseOutput) ComputeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.ComputeMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) ExcludeFromCapacityAllocation() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.ExcludeFromCapacityAllocation }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) IsApplicableForAllComputeModes() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.IsApplicableForAllComputeModes }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) IsLinux() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *bool { return v.IsLinux }).(pulumi.BoolPtrOutput)
}

func (o StampCapacityResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) SiteMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.SiteMode }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) TotalCapacity() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *float64 { return v.TotalCapacity }).(pulumi.Float64PtrOutput)
}

func (o StampCapacityResponseOutput) Unit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.Unit }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o StampCapacityResponseOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v StampCapacityResponse) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type StampCapacityResponseArrayOutput struct{ *pulumi.OutputState }

func (StampCapacityResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StampCapacityResponse)(nil)).Elem()
}

func (o StampCapacityResponseArrayOutput) ToStampCapacityResponseArrayOutput() StampCapacityResponseArrayOutput {
	return o
}

func (o StampCapacityResponseArrayOutput) ToStampCapacityResponseArrayOutputWithContext(ctx context.Context) StampCapacityResponseArrayOutput {
	return o
}

func (o StampCapacityResponseArrayOutput) Index(i pulumi.IntInput) StampCapacityResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) StampCapacityResponse {
		return vs[0].([]StampCapacityResponse)[vs[1].(int)]
	}).(StampCapacityResponseOutput)
}

type StaticSiteBuildProperties struct {
	ApiLocation         *string `pulumi:"apiLocation"`
	AppArtifactLocation *string `pulumi:"appArtifactLocation"`
	AppLocation         *string `pulumi:"appLocation"`
}





type StaticSiteBuildPropertiesInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesOutput() StaticSiteBuildPropertiesOutput
	ToStaticSiteBuildPropertiesOutputWithContext(context.Context) StaticSiteBuildPropertiesOutput
}

type StaticSiteBuildPropertiesArgs struct {
	ApiLocation         pulumi.StringPtrInput `pulumi:"apiLocation"`
	AppArtifactLocation pulumi.StringPtrInput `pulumi:"appArtifactLocation"`
	AppLocation         pulumi.StringPtrInput `pulumi:"appLocation"`
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

func (o StaticSiteBuildPropertiesOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.ApiLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppArtifactLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildProperties) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
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

func (o StaticSiteBuildPropertiesPtrOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppLocation
	}).(pulumi.StringPtrOutput)
}

type StaticSiteBuildPropertiesResponse struct {
	ApiLocation         *string `pulumi:"apiLocation"`
	AppArtifactLocation *string `pulumi:"appArtifactLocation"`
	AppLocation         *string `pulumi:"appLocation"`
}





type StaticSiteBuildPropertiesResponseInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesResponseOutput() StaticSiteBuildPropertiesResponseOutput
	ToStaticSiteBuildPropertiesResponseOutputWithContext(context.Context) StaticSiteBuildPropertiesResponseOutput
}

type StaticSiteBuildPropertiesResponseArgs struct {
	ApiLocation         pulumi.StringPtrInput `pulumi:"apiLocation"`
	AppArtifactLocation pulumi.StringPtrInput `pulumi:"appArtifactLocation"`
	AppLocation         pulumi.StringPtrInput `pulumi:"appLocation"`
}

func (StaticSiteBuildPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteBuildPropertiesResponse)(nil)).Elem()
}

func (i StaticSiteBuildPropertiesResponseArgs) ToStaticSiteBuildPropertiesResponseOutput() StaticSiteBuildPropertiesResponseOutput {
	return i.ToStaticSiteBuildPropertiesResponseOutputWithContext(context.Background())
}

func (i StaticSiteBuildPropertiesResponseArgs) ToStaticSiteBuildPropertiesResponseOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesResponseOutput)
}

func (i StaticSiteBuildPropertiesResponseArgs) ToStaticSiteBuildPropertiesResponsePtrOutput() StaticSiteBuildPropertiesResponsePtrOutput {
	return i.ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i StaticSiteBuildPropertiesResponseArgs) ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesResponseOutput).ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(ctx)
}









type StaticSiteBuildPropertiesResponsePtrInput interface {
	pulumi.Input

	ToStaticSiteBuildPropertiesResponsePtrOutput() StaticSiteBuildPropertiesResponsePtrOutput
	ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(context.Context) StaticSiteBuildPropertiesResponsePtrOutput
}

type staticSiteBuildPropertiesResponsePtrType StaticSiteBuildPropertiesResponseArgs

func StaticSiteBuildPropertiesResponsePtr(v *StaticSiteBuildPropertiesResponseArgs) StaticSiteBuildPropertiesResponsePtrInput {
	return (*staticSiteBuildPropertiesResponsePtrType)(v)
}

func (*staticSiteBuildPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**StaticSiteBuildPropertiesResponse)(nil)).Elem()
}

func (i *staticSiteBuildPropertiesResponsePtrType) ToStaticSiteBuildPropertiesResponsePtrOutput() StaticSiteBuildPropertiesResponsePtrOutput {
	return i.ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *staticSiteBuildPropertiesResponsePtrType) ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteBuildPropertiesResponsePtrOutput)
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

func (o StaticSiteBuildPropertiesResponseOutput) ToStaticSiteBuildPropertiesResponsePtrOutput() StaticSiteBuildPropertiesResponsePtrOutput {
	return o.ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o StaticSiteBuildPropertiesResponseOutput) ToStaticSiteBuildPropertiesResponsePtrOutputWithContext(ctx context.Context) StaticSiteBuildPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v StaticSiteBuildPropertiesResponse) *StaticSiteBuildPropertiesResponse {
		return &v
	}).(StaticSiteBuildPropertiesResponsePtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) ApiLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.ApiLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppArtifactLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppArtifactLocation }).(pulumi.StringPtrOutput)
}

func (o StaticSiteBuildPropertiesResponseOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v StaticSiteBuildPropertiesResponse) *string { return v.AppLocation }).(pulumi.StringPtrOutput)
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

func (o StaticSiteBuildPropertiesResponsePtrOutput) AppLocation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *StaticSiteBuildPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppLocation
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





type StaticSiteUserARMResourceResponseInput interface {
	pulumi.Input

	ToStaticSiteUserARMResourceResponseOutput() StaticSiteUserARMResourceResponseOutput
	ToStaticSiteUserARMResourceResponseOutputWithContext(context.Context) StaticSiteUserARMResourceResponseOutput
}

type StaticSiteUserARMResourceResponseArgs struct {
	DisplayName pulumi.StringInput    `pulumi:"displayName"`
	Id          pulumi.StringInput    `pulumi:"id"`
	Kind        pulumi.StringPtrInput `pulumi:"kind"`
	Name        pulumi.StringInput    `pulumi:"name"`
	Provider    pulumi.StringInput    `pulumi:"provider"`
	Roles       pulumi.StringPtrInput `pulumi:"roles"`
	Type        pulumi.StringInput    `pulumi:"type"`
	UserId      pulumi.StringInput    `pulumi:"userId"`
}

func (StaticSiteUserARMResourceResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StaticSiteUserARMResourceResponse)(nil)).Elem()
}

func (i StaticSiteUserARMResourceResponseArgs) ToStaticSiteUserARMResourceResponseOutput() StaticSiteUserARMResourceResponseOutput {
	return i.ToStaticSiteUserARMResourceResponseOutputWithContext(context.Background())
}

func (i StaticSiteUserARMResourceResponseArgs) ToStaticSiteUserARMResourceResponseOutputWithContext(ctx context.Context) StaticSiteUserARMResourceResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteUserARMResourceResponseOutput)
}





type StaticSiteUserARMResourceResponseArrayInput interface {
	pulumi.Input

	ToStaticSiteUserARMResourceResponseArrayOutput() StaticSiteUserARMResourceResponseArrayOutput
	ToStaticSiteUserARMResourceResponseArrayOutputWithContext(context.Context) StaticSiteUserARMResourceResponseArrayOutput
}

type StaticSiteUserARMResourceResponseArray []StaticSiteUserARMResourceResponseInput

func (StaticSiteUserARMResourceResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StaticSiteUserARMResourceResponse)(nil)).Elem()
}

func (i StaticSiteUserARMResourceResponseArray) ToStaticSiteUserARMResourceResponseArrayOutput() StaticSiteUserARMResourceResponseArrayOutput {
	return i.ToStaticSiteUserARMResourceResponseArrayOutputWithContext(context.Background())
}

func (i StaticSiteUserARMResourceResponseArray) ToStaticSiteUserARMResourceResponseArrayOutputWithContext(ctx context.Context) StaticSiteUserARMResourceResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StaticSiteUserARMResourceResponseArrayOutput)
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

type StatusCodesBasedTrigger struct {
	Count        *int    `pulumi:"count"`
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
	Status       *int    `pulumi:"status"`
	SubStatus    *int    `pulumi:"subStatus"`
	TimeInterval *string `pulumi:"timeInterval"`
	Win32Status  *int    `pulumi:"win32Status"`
}





type StatusCodesBasedTriggerResponseInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerResponseOutput() StatusCodesBasedTriggerResponseOutput
	ToStatusCodesBasedTriggerResponseOutputWithContext(context.Context) StatusCodesBasedTriggerResponseOutput
}

type StatusCodesBasedTriggerResponseArgs struct {
	Count        pulumi.IntPtrInput    `pulumi:"count"`
	Status       pulumi.IntPtrInput    `pulumi:"status"`
	SubStatus    pulumi.IntPtrInput    `pulumi:"subStatus"`
	TimeInterval pulumi.StringPtrInput `pulumi:"timeInterval"`
	Win32Status  pulumi.IntPtrInput    `pulumi:"win32Status"`
}

func (StatusCodesBasedTriggerResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (i StatusCodesBasedTriggerResponseArgs) ToStatusCodesBasedTriggerResponseOutput() StatusCodesBasedTriggerResponseOutput {
	return i.ToStatusCodesBasedTriggerResponseOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerResponseArgs) ToStatusCodesBasedTriggerResponseOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerResponseOutput)
}





type StatusCodesBasedTriggerResponseArrayInput interface {
	pulumi.Input

	ToStatusCodesBasedTriggerResponseArrayOutput() StatusCodesBasedTriggerResponseArrayOutput
	ToStatusCodesBasedTriggerResponseArrayOutputWithContext(context.Context) StatusCodesBasedTriggerResponseArrayOutput
}

type StatusCodesBasedTriggerResponseArray []StatusCodesBasedTriggerResponseInput

func (StatusCodesBasedTriggerResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]StatusCodesBasedTriggerResponse)(nil)).Elem()
}

func (i StatusCodesBasedTriggerResponseArray) ToStatusCodesBasedTriggerResponseArrayOutput() StatusCodesBasedTriggerResponseArrayOutput {
	return i.ToStatusCodesBasedTriggerResponseArrayOutputWithContext(context.Background())
}

func (i StatusCodesBasedTriggerResponseArray) ToStatusCodesBasedTriggerResponseArrayOutputWithContext(ctx context.Context) StatusCodesBasedTriggerResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(StatusCodesBasedTriggerResponseArrayOutput)
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





type VirtualApplicationResponseInput interface {
	pulumi.Input

	ToVirtualApplicationResponseOutput() VirtualApplicationResponseOutput
	ToVirtualApplicationResponseOutputWithContext(context.Context) VirtualApplicationResponseOutput
}

type VirtualApplicationResponseArgs struct {
	PhysicalPath       pulumi.StringPtrInput              `pulumi:"physicalPath"`
	PreloadEnabled     pulumi.BoolPtrInput                `pulumi:"preloadEnabled"`
	VirtualDirectories VirtualDirectoryResponseArrayInput `pulumi:"virtualDirectories"`
	VirtualPath        pulumi.StringPtrInput              `pulumi:"virtualPath"`
}

func (VirtualApplicationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualApplicationResponse)(nil)).Elem()
}

func (i VirtualApplicationResponseArgs) ToVirtualApplicationResponseOutput() VirtualApplicationResponseOutput {
	return i.ToVirtualApplicationResponseOutputWithContext(context.Background())
}

func (i VirtualApplicationResponseArgs) ToVirtualApplicationResponseOutputWithContext(ctx context.Context) VirtualApplicationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationResponseOutput)
}





type VirtualApplicationResponseArrayInput interface {
	pulumi.Input

	ToVirtualApplicationResponseArrayOutput() VirtualApplicationResponseArrayOutput
	ToVirtualApplicationResponseArrayOutputWithContext(context.Context) VirtualApplicationResponseArrayOutput
}

type VirtualApplicationResponseArray []VirtualApplicationResponseInput

func (VirtualApplicationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualApplicationResponse)(nil)).Elem()
}

func (i VirtualApplicationResponseArray) ToVirtualApplicationResponseArrayOutput() VirtualApplicationResponseArrayOutput {
	return i.ToVirtualApplicationResponseArrayOutputWithContext(context.Background())
}

func (i VirtualApplicationResponseArray) ToVirtualApplicationResponseArrayOutputWithContext(ctx context.Context) VirtualApplicationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualApplicationResponseArrayOutput)
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





type VirtualDirectoryResponseInput interface {
	pulumi.Input

	ToVirtualDirectoryResponseOutput() VirtualDirectoryResponseOutput
	ToVirtualDirectoryResponseOutputWithContext(context.Context) VirtualDirectoryResponseOutput
}

type VirtualDirectoryResponseArgs struct {
	PhysicalPath pulumi.StringPtrInput `pulumi:"physicalPath"`
	VirtualPath  pulumi.StringPtrInput `pulumi:"virtualPath"`
}

func (VirtualDirectoryResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualDirectoryResponse)(nil)).Elem()
}

func (i VirtualDirectoryResponseArgs) ToVirtualDirectoryResponseOutput() VirtualDirectoryResponseOutput {
	return i.ToVirtualDirectoryResponseOutputWithContext(context.Background())
}

func (i VirtualDirectoryResponseArgs) ToVirtualDirectoryResponseOutputWithContext(ctx context.Context) VirtualDirectoryResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryResponseOutput)
}





type VirtualDirectoryResponseArrayInput interface {
	pulumi.Input

	ToVirtualDirectoryResponseArrayOutput() VirtualDirectoryResponseArrayOutput
	ToVirtualDirectoryResponseArrayOutputWithContext(context.Context) VirtualDirectoryResponseArrayOutput
}

type VirtualDirectoryResponseArray []VirtualDirectoryResponseInput

func (VirtualDirectoryResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualDirectoryResponse)(nil)).Elem()
}

func (i VirtualDirectoryResponseArray) ToVirtualDirectoryResponseArrayOutput() VirtualDirectoryResponseArrayOutput {
	return i.ToVirtualDirectoryResponseArrayOutputWithContext(context.Background())
}

func (i VirtualDirectoryResponseArray) ToVirtualDirectoryResponseArrayOutputWithContext(ctx context.Context) VirtualDirectoryResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualDirectoryResponseArrayOutput)
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

type VirtualIPMappingResponse struct {
	InUse             *bool   `pulumi:"inUse"`
	InternalHttpPort  *int    `pulumi:"internalHttpPort"`
	InternalHttpsPort *int    `pulumi:"internalHttpsPort"`
	ServiceName       *string `pulumi:"serviceName"`
	VirtualIP         *string `pulumi:"virtualIP"`
}





type VirtualIPMappingResponseInput interface {
	pulumi.Input

	ToVirtualIPMappingResponseOutput() VirtualIPMappingResponseOutput
	ToVirtualIPMappingResponseOutputWithContext(context.Context) VirtualIPMappingResponseOutput
}

type VirtualIPMappingResponseArgs struct {
	InUse             pulumi.BoolPtrInput   `pulumi:"inUse"`
	InternalHttpPort  pulumi.IntPtrInput    `pulumi:"internalHttpPort"`
	InternalHttpsPort pulumi.IntPtrInput    `pulumi:"internalHttpsPort"`
	ServiceName       pulumi.StringPtrInput `pulumi:"serviceName"`
	VirtualIP         pulumi.StringPtrInput `pulumi:"virtualIP"`
}

func (VirtualIPMappingResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualIPMappingResponse)(nil)).Elem()
}

func (i VirtualIPMappingResponseArgs) ToVirtualIPMappingResponseOutput() VirtualIPMappingResponseOutput {
	return i.ToVirtualIPMappingResponseOutputWithContext(context.Background())
}

func (i VirtualIPMappingResponseArgs) ToVirtualIPMappingResponseOutputWithContext(ctx context.Context) VirtualIPMappingResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIPMappingResponseOutput)
}





type VirtualIPMappingResponseArrayInput interface {
	pulumi.Input

	ToVirtualIPMappingResponseArrayOutput() VirtualIPMappingResponseArrayOutput
	ToVirtualIPMappingResponseArrayOutputWithContext(context.Context) VirtualIPMappingResponseArrayOutput
}

type VirtualIPMappingResponseArray []VirtualIPMappingResponseInput

func (VirtualIPMappingResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualIPMappingResponse)(nil)).Elem()
}

func (i VirtualIPMappingResponseArray) ToVirtualIPMappingResponseArrayOutput() VirtualIPMappingResponseArrayOutput {
	return i.ToVirtualIPMappingResponseArrayOutputWithContext(context.Background())
}

func (i VirtualIPMappingResponseArray) ToVirtualIPMappingResponseArrayOutputWithContext(ctx context.Context) VirtualIPMappingResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualIPMappingResponseArrayOutput)
}

type VirtualIPMappingResponseOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualIPMappingResponse)(nil)).Elem()
}

func (o VirtualIPMappingResponseOutput) ToVirtualIPMappingResponseOutput() VirtualIPMappingResponseOutput {
	return o
}

func (o VirtualIPMappingResponseOutput) ToVirtualIPMappingResponseOutputWithContext(ctx context.Context) VirtualIPMappingResponseOutput {
	return o
}

func (o VirtualIPMappingResponseOutput) InUse() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *bool { return v.InUse }).(pulumi.BoolPtrOutput)
}

func (o VirtualIPMappingResponseOutput) InternalHttpPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *int { return v.InternalHttpPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingResponseOutput) InternalHttpsPort() pulumi.IntPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *int { return v.InternalHttpsPort }).(pulumi.IntPtrOutput)
}

func (o VirtualIPMappingResponseOutput) ServiceName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *string { return v.ServiceName }).(pulumi.StringPtrOutput)
}

func (o VirtualIPMappingResponseOutput) VirtualIP() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualIPMappingResponse) *string { return v.VirtualIP }).(pulumi.StringPtrOutput)
}

type VirtualIPMappingResponseArrayOutput struct{ *pulumi.OutputState }

func (VirtualIPMappingResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VirtualIPMappingResponse)(nil)).Elem()
}

func (o VirtualIPMappingResponseArrayOutput) ToVirtualIPMappingResponseArrayOutput() VirtualIPMappingResponseArrayOutput {
	return o
}

func (o VirtualIPMappingResponseArrayOutput) ToVirtualIPMappingResponseArrayOutputWithContext(ctx context.Context) VirtualIPMappingResponseArrayOutput {
	return o
}

func (o VirtualIPMappingResponseArrayOutput) Index(i pulumi.IntInput) VirtualIPMappingResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) VirtualIPMappingResponse {
		return vs[0].([]VirtualIPMappingResponse)[vs[1].(int)]
	}).(VirtualIPMappingResponseOutput)
}

type VirtualNetworkProfile struct {
	Id     *string `pulumi:"id"`
	Subnet *string `pulumi:"subnet"`
}





type VirtualNetworkProfileInput interface {
	pulumi.Input

	ToVirtualNetworkProfileOutput() VirtualNetworkProfileOutput
	ToVirtualNetworkProfileOutputWithContext(context.Context) VirtualNetworkProfileOutput
}

type VirtualNetworkProfileArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
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

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileArgs) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileOutput).ToVirtualNetworkProfilePtrOutputWithContext(ctx)
}









type VirtualNetworkProfilePtrInput interface {
	pulumi.Input

	ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput
	ToVirtualNetworkProfilePtrOutputWithContext(context.Context) VirtualNetworkProfilePtrOutput
}

type virtualNetworkProfilePtrType VirtualNetworkProfileArgs

func VirtualNetworkProfilePtr(v *VirtualNetworkProfileArgs) VirtualNetworkProfilePtrInput {
	return (*virtualNetworkProfilePtrType)(v)
}

func (*virtualNetworkProfilePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return i.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkProfilePtrType) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfilePtrOutput)
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

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o.ToVirtualNetworkProfilePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkProfileOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkProfile) *VirtualNetworkProfile {
		return &v
	}).(VirtualNetworkProfilePtrOutput)
}

func (o VirtualNetworkProfileOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfile) *string { return v.Subnet }).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfilePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfilePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfile)(nil)).Elem()
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutput() VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) ToVirtualNetworkProfilePtrOutputWithContext(ctx context.Context) VirtualNetworkProfilePtrOutput {
	return o
}

func (o VirtualNetworkProfilePtrOutput) Elem() VirtualNetworkProfileOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) VirtualNetworkProfile {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfile
		return ret
	}).(VirtualNetworkProfileOutput)
}

func (o VirtualNetworkProfilePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfilePtrOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfile) *string {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(pulumi.StringPtrOutput)
}

type VirtualNetworkProfileResponse struct {
	Id     *string `pulumi:"id"`
	Name   string  `pulumi:"name"`
	Subnet *string `pulumi:"subnet"`
	Type   string  `pulumi:"type"`
}





type VirtualNetworkProfileResponseInput interface {
	pulumi.Input

	ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput
	ToVirtualNetworkProfileResponseOutputWithContext(context.Context) VirtualNetworkProfileResponseOutput
}

type VirtualNetworkProfileResponseArgs struct {
	Id     pulumi.StringPtrInput `pulumi:"id"`
	Name   pulumi.StringInput    `pulumi:"name"`
	Subnet pulumi.StringPtrInput `pulumi:"subnet"`
	Type   pulumi.StringInput    `pulumi:"type"`
}

func (VirtualNetworkProfileResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VirtualNetworkProfileResponse)(nil)).Elem()
}

func (i VirtualNetworkProfileResponseArgs) ToVirtualNetworkProfileResponseOutput() VirtualNetworkProfileResponseOutput {
	return i.ToVirtualNetworkProfileResponseOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileResponseArgs) ToVirtualNetworkProfileResponseOutputWithContext(ctx context.Context) VirtualNetworkProfileResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileResponseOutput)
}

func (i VirtualNetworkProfileResponseArgs) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return i.ToVirtualNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i VirtualNetworkProfileResponseArgs) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileResponseOutput).ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx)
}









type VirtualNetworkProfileResponsePtrInput interface {
	pulumi.Input

	ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput
	ToVirtualNetworkProfileResponsePtrOutputWithContext(context.Context) VirtualNetworkProfileResponsePtrOutput
}

type virtualNetworkProfileResponsePtrType VirtualNetworkProfileResponseArgs

func VirtualNetworkProfileResponsePtr(v *VirtualNetworkProfileResponseArgs) VirtualNetworkProfileResponsePtrInput {
	return (*virtualNetworkProfileResponsePtrType)(v)
}

func (*virtualNetworkProfileResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfileResponse)(nil)).Elem()
}

func (i *virtualNetworkProfileResponsePtrType) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return i.ToVirtualNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (i *virtualNetworkProfileResponsePtrType) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VirtualNetworkProfileResponsePtrOutput)
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

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return o.ToVirtualNetworkProfileResponsePtrOutputWithContext(context.Background())
}

func (o VirtualNetworkProfileResponseOutput) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v VirtualNetworkProfileResponse) *VirtualNetworkProfileResponse {
		return &v
	}).(VirtualNetworkProfileResponsePtrOutput)
}

func (o VirtualNetworkProfileResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v VirtualNetworkProfileResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
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

type VirtualNetworkProfileResponsePtrOutput struct{ *pulumi.OutputState }

func (VirtualNetworkProfileResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**VirtualNetworkProfileResponse)(nil)).Elem()
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutput() VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) ToVirtualNetworkProfileResponsePtrOutputWithContext(ctx context.Context) VirtualNetworkProfileResponsePtrOutput {
	return o
}

func (o VirtualNetworkProfileResponsePtrOutput) Elem() VirtualNetworkProfileResponseOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) VirtualNetworkProfileResponse {
		if v != nil {
			return *v
		}
		var ret VirtualNetworkProfileResponse
		return ret
	}).(VirtualNetworkProfileResponseOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Subnet() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return v.Subnet
	}).(pulumi.StringPtrOutput)
}

func (o VirtualNetworkProfileResponsePtrOutput) Type() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *VirtualNetworkProfileResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Type
	}).(pulumi.StringPtrOutput)
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





type VnetRouteResponseInput interface {
	pulumi.Input

	ToVnetRouteResponseOutput() VnetRouteResponseOutput
	ToVnetRouteResponseOutputWithContext(context.Context) VnetRouteResponseOutput
}

type VnetRouteResponseArgs struct {
	EndAddress   pulumi.StringPtrInput `pulumi:"endAddress"`
	Id           pulumi.StringInput    `pulumi:"id"`
	Kind         pulumi.StringPtrInput `pulumi:"kind"`
	Name         pulumi.StringInput    `pulumi:"name"`
	RouteType    pulumi.StringPtrInput `pulumi:"routeType"`
	StartAddress pulumi.StringPtrInput `pulumi:"startAddress"`
	Type         pulumi.StringInput    `pulumi:"type"`
}

func (VnetRouteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*VnetRouteResponse)(nil)).Elem()
}

func (i VnetRouteResponseArgs) ToVnetRouteResponseOutput() VnetRouteResponseOutput {
	return i.ToVnetRouteResponseOutputWithContext(context.Background())
}

func (i VnetRouteResponseArgs) ToVnetRouteResponseOutputWithContext(ctx context.Context) VnetRouteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteResponseOutput)
}





type VnetRouteResponseArrayInput interface {
	pulumi.Input

	ToVnetRouteResponseArrayOutput() VnetRouteResponseArrayOutput
	ToVnetRouteResponseArrayOutputWithContext(context.Context) VnetRouteResponseArrayOutput
}

type VnetRouteResponseArray []VnetRouteResponseInput

func (VnetRouteResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]VnetRouteResponse)(nil)).Elem()
}

func (i VnetRouteResponseArray) ToVnetRouteResponseArrayOutput() VnetRouteResponseArrayOutput {
	return i.ToVnetRouteResponseArrayOutputWithContext(context.Background())
}

func (i VnetRouteResponseArray) ToVnetRouteResponseArrayOutputWithContext(ctx context.Context) VnetRouteResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(VnetRouteResponseArrayOutput)
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

type WorkerPool struct {
	ComputeMode  *ComputeModeOptions `pulumi:"computeMode"`
	WorkerCount  *int                `pulumi:"workerCount"`
	WorkerSize   *string             `pulumi:"workerSize"`
	WorkerSizeId *int                `pulumi:"workerSizeId"`
}





type WorkerPoolInput interface {
	pulumi.Input

	ToWorkerPoolOutput() WorkerPoolOutput
	ToWorkerPoolOutputWithContext(context.Context) WorkerPoolOutput
}

type WorkerPoolArgs struct {
	ComputeMode  ComputeModeOptionsPtrInput `pulumi:"computeMode"`
	WorkerCount  pulumi.IntPtrInput         `pulumi:"workerCount"`
	WorkerSize   pulumi.StringPtrInput      `pulumi:"workerSize"`
	WorkerSizeId pulumi.IntPtrInput         `pulumi:"workerSizeId"`
}

func (WorkerPoolArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil)).Elem()
}

func (i WorkerPoolArgs) ToWorkerPoolOutput() WorkerPoolOutput {
	return i.ToWorkerPoolOutputWithContext(context.Background())
}

func (i WorkerPoolArgs) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolOutput)
}





type WorkerPoolArrayInput interface {
	pulumi.Input

	ToWorkerPoolArrayOutput() WorkerPoolArrayOutput
	ToWorkerPoolArrayOutputWithContext(context.Context) WorkerPoolArrayOutput
}

type WorkerPoolArray []WorkerPoolInput

func (WorkerPoolArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil)).Elem()
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return i.ToWorkerPoolArrayOutputWithContext(context.Background())
}

func (i WorkerPoolArray) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolArrayOutput)
}

type WorkerPoolOutput struct{ *pulumi.OutputState }

func (WorkerPoolOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPool)(nil)).Elem()
}

func (o WorkerPoolOutput) ToWorkerPoolOutput() WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ToWorkerPoolOutputWithContext(ctx context.Context) WorkerPoolOutput {
	return o
}

func (o WorkerPoolOutput) ComputeMode() ComputeModeOptionsPtrOutput {
	return o.ApplyT(func(v WorkerPool) *ComputeModeOptions { return v.ComputeMode }).(ComputeModeOptionsPtrOutput)
}

func (o WorkerPoolOutput) WorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPool) *int { return v.WorkerCount }).(pulumi.IntPtrOutput)
}

func (o WorkerPoolOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPool) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPool) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type WorkerPoolArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPool)(nil)).Elem()
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutput() WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) ToWorkerPoolArrayOutputWithContext(ctx context.Context) WorkerPoolArrayOutput {
	return o
}

func (o WorkerPoolArrayOutput) Index(i pulumi.IntInput) WorkerPoolOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPool {
		return vs[0].([]WorkerPool)[vs[1].(int)]
	}).(WorkerPoolOutput)
}

type WorkerPoolResponse struct {
	ComputeMode   *string  `pulumi:"computeMode"`
	InstanceNames []string `pulumi:"instanceNames"`
	WorkerCount   *int     `pulumi:"workerCount"`
	WorkerSize    *string  `pulumi:"workerSize"`
	WorkerSizeId  *int     `pulumi:"workerSizeId"`
}





type WorkerPoolResponseInput interface {
	pulumi.Input

	ToWorkerPoolResponseOutput() WorkerPoolResponseOutput
	ToWorkerPoolResponseOutputWithContext(context.Context) WorkerPoolResponseOutput
}

type WorkerPoolResponseArgs struct {
	ComputeMode   pulumi.StringPtrInput   `pulumi:"computeMode"`
	InstanceNames pulumi.StringArrayInput `pulumi:"instanceNames"`
	WorkerCount   pulumi.IntPtrInput      `pulumi:"workerCount"`
	WorkerSize    pulumi.StringPtrInput   `pulumi:"workerSize"`
	WorkerSizeId  pulumi.IntPtrInput      `pulumi:"workerSizeId"`
}

func (WorkerPoolResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPoolResponse)(nil)).Elem()
}

func (i WorkerPoolResponseArgs) ToWorkerPoolResponseOutput() WorkerPoolResponseOutput {
	return i.ToWorkerPoolResponseOutputWithContext(context.Background())
}

func (i WorkerPoolResponseArgs) ToWorkerPoolResponseOutputWithContext(ctx context.Context) WorkerPoolResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolResponseOutput)
}





type WorkerPoolResponseArrayInput interface {
	pulumi.Input

	ToWorkerPoolResponseArrayOutput() WorkerPoolResponseArrayOutput
	ToWorkerPoolResponseArrayOutputWithContext(context.Context) WorkerPoolResponseArrayOutput
}

type WorkerPoolResponseArray []WorkerPoolResponseInput

func (WorkerPoolResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPoolResponse)(nil)).Elem()
}

func (i WorkerPoolResponseArray) ToWorkerPoolResponseArrayOutput() WorkerPoolResponseArrayOutput {
	return i.ToWorkerPoolResponseArrayOutputWithContext(context.Background())
}

func (i WorkerPoolResponseArray) ToWorkerPoolResponseArrayOutputWithContext(ctx context.Context) WorkerPoolResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WorkerPoolResponseArrayOutput)
}

type WorkerPoolResponseOutput struct{ *pulumi.OutputState }

func (WorkerPoolResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WorkerPoolResponse)(nil)).Elem()
}

func (o WorkerPoolResponseOutput) ToWorkerPoolResponseOutput() WorkerPoolResponseOutput {
	return o
}

func (o WorkerPoolResponseOutput) ToWorkerPoolResponseOutputWithContext(ctx context.Context) WorkerPoolResponseOutput {
	return o
}

func (o WorkerPoolResponseOutput) ComputeMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.ComputeMode }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) InstanceNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v WorkerPoolResponse) []string { return v.InstanceNames }).(pulumi.StringArrayOutput)
}

func (o WorkerPoolResponseOutput) WorkerCount() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *int { return v.WorkerCount }).(pulumi.IntPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerSize() pulumi.StringPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *string { return v.WorkerSize }).(pulumi.StringPtrOutput)
}

func (o WorkerPoolResponseOutput) WorkerSizeId() pulumi.IntPtrOutput {
	return o.ApplyT(func(v WorkerPoolResponse) *int { return v.WorkerSizeId }).(pulumi.IntPtrOutput)
}

type WorkerPoolResponseArrayOutput struct{ *pulumi.OutputState }

func (WorkerPoolResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WorkerPoolResponse)(nil)).Elem()
}

func (o WorkerPoolResponseArrayOutput) ToWorkerPoolResponseArrayOutput() WorkerPoolResponseArrayOutput {
	return o
}

func (o WorkerPoolResponseArrayOutput) ToWorkerPoolResponseArrayOutputWithContext(ctx context.Context) WorkerPoolResponseArrayOutput {
	return o
}

func (o WorkerPoolResponseArrayOutput) Index(i pulumi.IntInput) WorkerPoolResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WorkerPoolResponse {
		return vs[0].([]WorkerPoolResponse)[vs[1].(int)]
	}).(WorkerPoolResponseOutput)
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
	pulumi.RegisterOutputType(ApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(ApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponseOutput{})
	pulumi.RegisterOutputType(ArmIdWrapperResponsePtrOutput{})
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
	pulumi.RegisterOutputType(CloningInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnStringInfoOutput{})
	pulumi.RegisterOutputType(ConnStringInfoArrayOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseOutput{})
	pulumi.RegisterOutputType(ConnStringInfoResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairMapOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseOutput{})
	pulumi.RegisterOutputType(ConnStringValueTypePairResponseMapOutput{})
	pulumi.RegisterOutputType(CorsSettingsOutput{})
	pulumi.RegisterOutputType(CorsSettingsPtrOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponseOutput{})
	pulumi.RegisterOutputType(CorsSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingArrayOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseOutput{})
	pulumi.RegisterOutputType(DatabaseBackupSettingResponseArrayOutput{})
	pulumi.RegisterOutputType(EnabledConfigOutput{})
	pulumi.RegisterOutputType(EnabledConfigPtrOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponseOutput{})
	pulumi.RegisterOutputType(EnabledConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(ExperimentsOutput{})
	pulumi.RegisterOutputType(ExperimentsPtrOutput{})
	pulumi.RegisterOutputType(ExperimentsResponseOutput{})
	pulumi.RegisterOutputType(ExperimentsResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemApplicationLogsConfigResponsePtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigPtrOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponseOutput{})
	pulumi.RegisterOutputType(FileSystemHttpLogsConfigResponsePtrOutput{})
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
	pulumi.RegisterOutputType(IdentifierResponseOutput{})
	pulumi.RegisterOutputType(IdentifierResponseArrayOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionArrayOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseOutput{})
	pulumi.RegisterOutputType(IpSecurityRestrictionResponseArrayOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityPtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponsePtrOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseUserAssignedIdentitiesOutput{})
	pulumi.RegisterOutputType(ManagedServiceIdentityResponseUserAssignedIdentitiesMapOutput{})
	pulumi.RegisterOutputType(NameValuePairOutput{})
	pulumi.RegisterOutputType(NameValuePairArrayOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseOutput{})
	pulumi.RegisterOutputType(NameValuePairResponseArrayOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryArrayOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryResponseOutput{})
	pulumi.RegisterOutputType(NetworkAccessControlEntryResponseArrayOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStatePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(PrivateLinkConnectionStateResponsePtrOutput{})
	pulumi.RegisterOutputType(PushSettingsOutput{})
	pulumi.RegisterOutputType(PushSettingsPtrOutput{})
	pulumi.RegisterOutputType(PushSettingsResponseOutput{})
	pulumi.RegisterOutputType(PushSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(RampUpRuleOutput{})
	pulumi.RegisterOutputType(RampUpRuleArrayOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseOutput{})
	pulumi.RegisterOutputType(RampUpRuleResponseArrayOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(RequestsBasedTriggerResponsePtrOutput{})
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
	pulumi.RegisterOutputType(SlotSwapStatusResponsePtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerPtrOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(SlowRequestsBasedTriggerResponsePtrOutput{})
	pulumi.RegisterOutputType(StampCapacityResponseOutput{})
	pulumi.RegisterOutputType(StampCapacityResponseArrayOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesPtrOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteBuildPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(StaticSiteUserARMResourceResponseOutput{})
	pulumi.RegisterOutputType(StaticSiteUserARMResourceResponseArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerArrayOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseOutput{})
	pulumi.RegisterOutputType(StatusCodesBasedTriggerResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplicationOutput{})
	pulumi.RegisterOutputType(VirtualApplicationArrayOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseOutput{})
	pulumi.RegisterOutputType(VirtualApplicationResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryArrayOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseOutput{})
	pulumi.RegisterOutputType(VirtualDirectoryResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingResponseOutput{})
	pulumi.RegisterOutputType(VirtualIPMappingResponseArrayOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfilePtrOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponseOutput{})
	pulumi.RegisterOutputType(VirtualNetworkProfileResponsePtrOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseOutput{})
	pulumi.RegisterOutputType(VnetRouteResponseArrayOutput{})
	pulumi.RegisterOutputType(WorkerPoolOutput{})
	pulumi.RegisterOutputType(WorkerPoolArrayOutput{})
	pulumi.RegisterOutputType(WorkerPoolResponseOutput{})
	pulumi.RegisterOutputType(WorkerPoolResponseArrayOutput{})
}
