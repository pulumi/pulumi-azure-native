


package botservice

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlexaChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *AlexaChannelProperties `pulumi:"properties"`
}

type AlexaChannelProperties struct {
	AlexaSkillId string `pulumi:"alexaSkillId"`
	IsEnabled    bool   `pulumi:"isEnabled"`
}

type AlexaChannelPropertiesResponse struct {
	AlexaSkillId       string `pulumi:"alexaSkillId"`
	IsEnabled          bool   `pulumi:"isEnabled"`
	ServiceEndpointUri string `pulumi:"serviceEndpointUri"`
	UrlFragment        string `pulumi:"urlFragment"`
}

type AlexaChannelResponse struct {
	ChannelName       string                          `pulumi:"channelName"`
	Etag              *string                         `pulumi:"etag"`
	Location          *string                         `pulumi:"location"`
	Properties        *AlexaChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                          `pulumi:"provisioningState"`
}

type BotProperties struct {
	AllSettings                       map[string]string `pulumi:"allSettings"`
	AppPasswordHint                   *string           `pulumi:"appPasswordHint"`
	CmekEncryptionStatus              *string           `pulumi:"cmekEncryptionStatus"`
	CmekKeyVaultUrl                   *string           `pulumi:"cmekKeyVaultUrl"`
	Description                       *string           `pulumi:"description"`
	DeveloperAppInsightKey            *string           `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        *string           `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId *string           `pulumi:"developerAppInsightsApplicationId"`
	DisableLocalAuth                  *bool             `pulumi:"disableLocalAuth"`
	DisplayName                       string            `pulumi:"displayName"`
	Endpoint                          string            `pulumi:"endpoint"`
	IconUrl                           *string           `pulumi:"iconUrl"`
	IsCmekEnabled                     *bool             `pulumi:"isCmekEnabled"`
	IsDeveloperAppInsightsApiKeySet   *bool             `pulumi:"isDeveloperAppInsightsApiKeySet"`
	IsStreamingSupported              *bool             `pulumi:"isStreamingSupported"`
	LuisAppIds                        []string          `pulumi:"luisAppIds"`
	LuisKey                           *string           `pulumi:"luisKey"`
	ManifestUrl                       *string           `pulumi:"manifestUrl"`
	MsaAppId                          string            `pulumi:"msaAppId"`
	MsaAppMSIResourceId               *string           `pulumi:"msaAppMSIResourceId"`
	MsaAppTenantId                    *string           `pulumi:"msaAppTenantId"`
	MsaAppType                        *string           `pulumi:"msaAppType"`
	OpenWithHint                      *string           `pulumi:"openWithHint"`
	Parameters                        map[string]string `pulumi:"parameters"`
	PublicNetworkAccess               *string           `pulumi:"publicNetworkAccess"`
	PublishingCredentials             *string           `pulumi:"publishingCredentials"`
	SchemaTransformationVersion       *string           `pulumi:"schemaTransformationVersion"`
	StorageResourceId                 *string           `pulumi:"storageResourceId"`
}


func (val *BotProperties) Defaults() *BotProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}





type BotPropertiesInput interface {
	pulumi.Input

	ToBotPropertiesOutput() BotPropertiesOutput
	ToBotPropertiesOutputWithContext(context.Context) BotPropertiesOutput
}

type BotPropertiesArgs struct {
	AllSettings                       pulumi.StringMapInput   `pulumi:"allSettings"`
	AppPasswordHint                   pulumi.StringPtrInput   `pulumi:"appPasswordHint"`
	CmekEncryptionStatus              pulumi.StringPtrInput   `pulumi:"cmekEncryptionStatus"`
	CmekKeyVaultUrl                   pulumi.StringPtrInput   `pulumi:"cmekKeyVaultUrl"`
	Description                       pulumi.StringPtrInput   `pulumi:"description"`
	DeveloperAppInsightKey            pulumi.StringPtrInput   `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        pulumi.StringPtrInput   `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput   `pulumi:"developerAppInsightsApplicationId"`
	DisableLocalAuth                  pulumi.BoolPtrInput     `pulumi:"disableLocalAuth"`
	DisplayName                       pulumi.StringInput      `pulumi:"displayName"`
	Endpoint                          pulumi.StringInput      `pulumi:"endpoint"`
	IconUrl                           pulumi.StringPtrInput   `pulumi:"iconUrl"`
	IsCmekEnabled                     pulumi.BoolPtrInput     `pulumi:"isCmekEnabled"`
	IsDeveloperAppInsightsApiKeySet   pulumi.BoolPtrInput     `pulumi:"isDeveloperAppInsightsApiKeySet"`
	IsStreamingSupported              pulumi.BoolPtrInput     `pulumi:"isStreamingSupported"`
	LuisAppIds                        pulumi.StringArrayInput `pulumi:"luisAppIds"`
	LuisKey                           pulumi.StringPtrInput   `pulumi:"luisKey"`
	ManifestUrl                       pulumi.StringPtrInput   `pulumi:"manifestUrl"`
	MsaAppId                          pulumi.StringInput      `pulumi:"msaAppId"`
	MsaAppMSIResourceId               pulumi.StringPtrInput   `pulumi:"msaAppMSIResourceId"`
	MsaAppTenantId                    pulumi.StringPtrInput   `pulumi:"msaAppTenantId"`
	MsaAppType                        pulumi.StringPtrInput   `pulumi:"msaAppType"`
	OpenWithHint                      pulumi.StringPtrInput   `pulumi:"openWithHint"`
	Parameters                        pulumi.StringMapInput   `pulumi:"parameters"`
	PublicNetworkAccess               pulumi.StringPtrInput   `pulumi:"publicNetworkAccess"`
	PublishingCredentials             pulumi.StringPtrInput   `pulumi:"publishingCredentials"`
	SchemaTransformationVersion       pulumi.StringPtrInput   `pulumi:"schemaTransformationVersion"`
	StorageResourceId                 pulumi.StringPtrInput   `pulumi:"storageResourceId"`
}

func (BotPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BotProperties)(nil)).Elem()
}

func (i BotPropertiesArgs) ToBotPropertiesOutput() BotPropertiesOutput {
	return i.ToBotPropertiesOutputWithContext(context.Background())
}

func (i BotPropertiesArgs) ToBotPropertiesOutputWithContext(ctx context.Context) BotPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotPropertiesOutput)
}

func (i BotPropertiesArgs) ToBotPropertiesPtrOutput() BotPropertiesPtrOutput {
	return i.ToBotPropertiesPtrOutputWithContext(context.Background())
}

func (i BotPropertiesArgs) ToBotPropertiesPtrOutputWithContext(ctx context.Context) BotPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotPropertiesOutput).ToBotPropertiesPtrOutputWithContext(ctx)
}









type BotPropertiesPtrInput interface {
	pulumi.Input

	ToBotPropertiesPtrOutput() BotPropertiesPtrOutput
	ToBotPropertiesPtrOutputWithContext(context.Context) BotPropertiesPtrOutput
}

type botPropertiesPtrType BotPropertiesArgs

func BotPropertiesPtr(v *BotPropertiesArgs) BotPropertiesPtrInput {
	return (*botPropertiesPtrType)(v)
}

func (*botPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BotProperties)(nil)).Elem()
}

func (i *botPropertiesPtrType) ToBotPropertiesPtrOutput() BotPropertiesPtrOutput {
	return i.ToBotPropertiesPtrOutputWithContext(context.Background())
}

func (i *botPropertiesPtrType) ToBotPropertiesPtrOutputWithContext(ctx context.Context) BotPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotPropertiesPtrOutput)
}

type BotPropertiesOutput struct{ *pulumi.OutputState }

func (BotPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BotProperties)(nil)).Elem()
}

func (o BotPropertiesOutput) ToBotPropertiesOutput() BotPropertiesOutput {
	return o
}

func (o BotPropertiesOutput) ToBotPropertiesOutputWithContext(ctx context.Context) BotPropertiesOutput {
	return o
}

func (o BotPropertiesOutput) ToBotPropertiesPtrOutput() BotPropertiesPtrOutput {
	return o.ToBotPropertiesPtrOutputWithContext(context.Background())
}

func (o BotPropertiesOutput) ToBotPropertiesPtrOutputWithContext(ctx context.Context) BotPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BotProperties) *BotProperties {
		return &v
	}).(BotPropertiesPtrOutput)
}

func (o BotPropertiesOutput) AllSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v BotProperties) map[string]string { return v.AllSettings }).(pulumi.StringMapOutput)
}

func (o BotPropertiesOutput) AppPasswordHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.AppPasswordHint }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) CmekEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.CmekEncryptionStatus }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) CmekKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.CmekKeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) DeveloperAppInsightKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.DeveloperAppInsightKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) DeveloperAppInsightsApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.DeveloperAppInsightsApiKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) DeveloperAppInsightsApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.DeveloperAppInsightsApplicationId }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotProperties) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v BotProperties) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o BotPropertiesOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v BotProperties) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o BotPropertiesOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) IsCmekEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotProperties) *bool { return v.IsCmekEnabled }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesOutput) IsDeveloperAppInsightsApiKeySet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotProperties) *bool { return v.IsDeveloperAppInsightsApiKeySet }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesOutput) IsStreamingSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotProperties) *bool { return v.IsStreamingSupported }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotProperties) []string { return v.LuisAppIds }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.LuisKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) ManifestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.ManifestUrl }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) MsaAppId() pulumi.StringOutput {
	return o.ApplyT(func(v BotProperties) string { return v.MsaAppId }).(pulumi.StringOutput)
}

func (o BotPropertiesOutput) MsaAppMSIResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.MsaAppMSIResourceId }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) MsaAppTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.MsaAppTenantId }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) MsaAppType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.MsaAppType }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) OpenWithHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.OpenWithHint }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v BotProperties) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o BotPropertiesOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) PublishingCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.PublishingCredentials }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.SchemaTransformationVersion }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) StorageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.StorageResourceId }).(pulumi.StringPtrOutput)
}

type BotPropertiesPtrOutput struct{ *pulumi.OutputState }

func (BotPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BotProperties)(nil)).Elem()
}

func (o BotPropertiesPtrOutput) ToBotPropertiesPtrOutput() BotPropertiesPtrOutput {
	return o
}

func (o BotPropertiesPtrOutput) ToBotPropertiesPtrOutputWithContext(ctx context.Context) BotPropertiesPtrOutput {
	return o
}

func (o BotPropertiesPtrOutput) Elem() BotPropertiesOutput {
	return o.ApplyT(func(v *BotProperties) BotProperties {
		if v != nil {
			return *v
		}
		var ret BotProperties
		return ret
	}).(BotPropertiesOutput)
}

func (o BotPropertiesPtrOutput) AllSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BotProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.AllSettings
	}).(pulumi.StringMapOutput)
}

func (o BotPropertiesPtrOutput) AppPasswordHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppPasswordHint
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) CmekEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.CmekEncryptionStatus
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) CmekKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.CmekKeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) DeveloperAppInsightKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.DeveloperAppInsightKey
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) DeveloperAppInsightsApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.DeveloperAppInsightsApiKey
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) DeveloperAppInsightsApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.DeveloperAppInsightsApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotProperties) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesPtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.IconUrl
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) IsCmekEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsCmekEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesPtrOutput) IsDeveloperAppInsightsApiKeySet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsDeveloperAppInsightsApiKeySet
	}).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesPtrOutput) IsStreamingSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsStreamingSupported
	}).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesPtrOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BotProperties) []string {
		if v == nil {
			return nil
		}
		return v.LuisAppIds
	}).(pulumi.StringArrayOutput)
}

func (o BotPropertiesPtrOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.LuisKey
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) ManifestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.ManifestUrl
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) MsaAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return &v.MsaAppId
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) MsaAppMSIResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.MsaAppMSIResourceId
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) MsaAppTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.MsaAppTenantId
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) MsaAppType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.MsaAppType
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) OpenWithHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.OpenWithHint
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v *BotProperties) map[string]string {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(pulumi.StringMapOutput)
}

func (o BotPropertiesPtrOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublicNetworkAccess
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) PublishingCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.PublishingCredentials
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.SchemaTransformationVersion
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesPtrOutput) StorageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.StorageResourceId
	}).(pulumi.StringPtrOutput)
}

type BotPropertiesResponse struct {
	AllSettings                       map[string]string `pulumi:"allSettings"`
	AppPasswordHint                   *string           `pulumi:"appPasswordHint"`
	CmekEncryptionStatus              *string           `pulumi:"cmekEncryptionStatus"`
	CmekKeyVaultUrl                   *string           `pulumi:"cmekKeyVaultUrl"`
	ConfiguredChannels                []string          `pulumi:"configuredChannels"`
	Description                       *string           `pulumi:"description"`
	DeveloperAppInsightKey            *string           `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        *string           `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId *string           `pulumi:"developerAppInsightsApplicationId"`
	DisableLocalAuth                  *bool             `pulumi:"disableLocalAuth"`
	DisplayName                       string            `pulumi:"displayName"`
	EnabledChannels                   []string          `pulumi:"enabledChannels"`
	Endpoint                          string            `pulumi:"endpoint"`
	EndpointVersion                   string            `pulumi:"endpointVersion"`
	IconUrl                           *string           `pulumi:"iconUrl"`
	IsCmekEnabled                     *bool             `pulumi:"isCmekEnabled"`
	IsDeveloperAppInsightsApiKeySet   *bool             `pulumi:"isDeveloperAppInsightsApiKeySet"`
	IsStreamingSupported              *bool             `pulumi:"isStreamingSupported"`
	LuisAppIds                        []string          `pulumi:"luisAppIds"`
	LuisKey                           *string           `pulumi:"luisKey"`
	ManifestUrl                       *string           `pulumi:"manifestUrl"`
	MigrationToken                    string            `pulumi:"migrationToken"`
	MsaAppId                          string            `pulumi:"msaAppId"`
	MsaAppMSIResourceId               *string           `pulumi:"msaAppMSIResourceId"`
	MsaAppTenantId                    *string           `pulumi:"msaAppTenantId"`
	MsaAppType                        *string           `pulumi:"msaAppType"`
	OpenWithHint                      *string           `pulumi:"openWithHint"`
	Parameters                        map[string]string `pulumi:"parameters"`
	ProvisioningState                 string            `pulumi:"provisioningState"`
	PublicNetworkAccess               *string           `pulumi:"publicNetworkAccess"`
	PublishingCredentials             *string           `pulumi:"publishingCredentials"`
	SchemaTransformationVersion       *string           `pulumi:"schemaTransformationVersion"`
	StorageResourceId                 *string           `pulumi:"storageResourceId"`
}


func (val *BotPropertiesResponse) Defaults() *BotPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.PublicNetworkAccess) {
		publicNetworkAccess_ := "Enabled"
		tmp.PublicNetworkAccess = &publicNetworkAccess_
	}
	return &tmp
}

type BotPropertiesResponseOutput struct{ *pulumi.OutputState }

func (BotPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*BotPropertiesResponse)(nil)).Elem()
}

func (o BotPropertiesResponseOutput) ToBotPropertiesResponseOutput() BotPropertiesResponseOutput {
	return o
}

func (o BotPropertiesResponseOutput) ToBotPropertiesResponseOutputWithContext(ctx context.Context) BotPropertiesResponseOutput {
	return o
}

func (o BotPropertiesResponseOutput) AllSettings() pulumi.StringMapOutput {
	return o.ApplyT(func(v BotPropertiesResponse) map[string]string { return v.AllSettings }).(pulumi.StringMapOutput)
}

func (o BotPropertiesResponseOutput) AppPasswordHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.AppPasswordHint }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) CmekEncryptionStatus() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.CmekEncryptionStatus }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) CmekKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.CmekKeyVaultUrl }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) ConfiguredChannels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotPropertiesResponse) []string { return v.ConfiguredChannels }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) DeveloperAppInsightKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.DeveloperAppInsightKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) DeveloperAppInsightsApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.DeveloperAppInsightsApiKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) DeveloperAppInsightsApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.DeveloperAppInsightsApplicationId }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) EnabledChannels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotPropertiesResponse) []string { return v.EnabledChannels }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponseOutput) Endpoint() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.Endpoint }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) EndpointVersion() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.EndpointVersion }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.IconUrl }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) IsCmekEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *bool { return v.IsCmekEnabled }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponseOutput) IsDeveloperAppInsightsApiKeySet() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *bool { return v.IsDeveloperAppInsightsApiKeySet }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponseOutput) IsStreamingSupported() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *bool { return v.IsStreamingSupported }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponseOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotPropertiesResponse) []string { return v.LuisAppIds }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponseOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.LuisKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) ManifestUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.ManifestUrl }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) MigrationToken() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.MigrationToken }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) MsaAppId() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.MsaAppId }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) MsaAppMSIResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.MsaAppMSIResourceId }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) MsaAppTenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.MsaAppTenantId }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) MsaAppType() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.MsaAppType }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) OpenWithHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.OpenWithHint }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) Parameters() pulumi.StringMapOutput {
	return o.ApplyT(func(v BotPropertiesResponse) map[string]string { return v.Parameters }).(pulumi.StringMapOutput)
}

func (o BotPropertiesResponseOutput) ProvisioningState() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.ProvisioningState }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) PublicNetworkAccess() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.PublicNetworkAccess }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) PublishingCredentials() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.PublishingCredentials }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.SchemaTransformationVersion }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) StorageResourceId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.StorageResourceId }).(pulumi.StringPtrOutput)
}

type ChannelSettingsResponse struct {
	BotIconUrl         *string        `pulumi:"botIconUrl"`
	BotId              *string        `pulumi:"botId"`
	ChannelDisplayName *string        `pulumi:"channelDisplayName"`
	ChannelId          *string        `pulumi:"channelId"`
	DisableLocalAuth   *bool          `pulumi:"disableLocalAuth"`
	ExtensionKey1      *string        `pulumi:"extensionKey1"`
	ExtensionKey2      *string        `pulumi:"extensionKey2"`
	IsEnabled          *bool          `pulumi:"isEnabled"`
	Sites              []SiteResponse `pulumi:"sites"`
}

type ConnectionSettingParameter struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}





type ConnectionSettingParameterInput interface {
	pulumi.Input

	ToConnectionSettingParameterOutput() ConnectionSettingParameterOutput
	ToConnectionSettingParameterOutputWithContext(context.Context) ConnectionSettingParameterOutput
}

type ConnectionSettingParameterArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConnectionSettingParameterArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingParameter)(nil)).Elem()
}

func (i ConnectionSettingParameterArgs) ToConnectionSettingParameterOutput() ConnectionSettingParameterOutput {
	return i.ToConnectionSettingParameterOutputWithContext(context.Background())
}

func (i ConnectionSettingParameterArgs) ToConnectionSettingParameterOutputWithContext(ctx context.Context) ConnectionSettingParameterOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingParameterOutput)
}





type ConnectionSettingParameterArrayInput interface {
	pulumi.Input

	ToConnectionSettingParameterArrayOutput() ConnectionSettingParameterArrayOutput
	ToConnectionSettingParameterArrayOutputWithContext(context.Context) ConnectionSettingParameterArrayOutput
}

type ConnectionSettingParameterArray []ConnectionSettingParameterInput

func (ConnectionSettingParameterArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionSettingParameter)(nil)).Elem()
}

func (i ConnectionSettingParameterArray) ToConnectionSettingParameterArrayOutput() ConnectionSettingParameterArrayOutput {
	return i.ToConnectionSettingParameterArrayOutputWithContext(context.Background())
}

func (i ConnectionSettingParameterArray) ToConnectionSettingParameterArrayOutputWithContext(ctx context.Context) ConnectionSettingParameterArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingParameterArrayOutput)
}

type ConnectionSettingParameterOutput struct{ *pulumi.OutputState }

func (ConnectionSettingParameterOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingParameter)(nil)).Elem()
}

func (o ConnectionSettingParameterOutput) ToConnectionSettingParameterOutput() ConnectionSettingParameterOutput {
	return o
}

func (o ConnectionSettingParameterOutput) ToConnectionSettingParameterOutputWithContext(ctx context.Context) ConnectionSettingParameterOutput {
	return o
}

func (o ConnectionSettingParameterOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingParameter) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingParameterOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingParameter) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConnectionSettingParameterArrayOutput struct{ *pulumi.OutputState }

func (ConnectionSettingParameterArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionSettingParameter)(nil)).Elem()
}

func (o ConnectionSettingParameterArrayOutput) ToConnectionSettingParameterArrayOutput() ConnectionSettingParameterArrayOutput {
	return o
}

func (o ConnectionSettingParameterArrayOutput) ToConnectionSettingParameterArrayOutputWithContext(ctx context.Context) ConnectionSettingParameterArrayOutput {
	return o
}

func (o ConnectionSettingParameterArrayOutput) Index(i pulumi.IntInput) ConnectionSettingParameterOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionSettingParameter {
		return vs[0].([]ConnectionSettingParameter)[vs[1].(int)]
	}).(ConnectionSettingParameterOutput)
}

type ConnectionSettingParameterResponse struct {
	Key   *string `pulumi:"key"`
	Value *string `pulumi:"value"`
}

type ConnectionSettingParameterResponseOutput struct{ *pulumi.OutputState }

func (ConnectionSettingParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingParameterResponse)(nil)).Elem()
}

func (o ConnectionSettingParameterResponseOutput) ToConnectionSettingParameterResponseOutput() ConnectionSettingParameterResponseOutput {
	return o
}

func (o ConnectionSettingParameterResponseOutput) ToConnectionSettingParameterResponseOutputWithContext(ctx context.Context) ConnectionSettingParameterResponseOutput {
	return o
}

func (o ConnectionSettingParameterResponseOutput) Key() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingParameterResponse) *string { return v.Key }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingParameterResponseOutput) Value() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingParameterResponse) *string { return v.Value }).(pulumi.StringPtrOutput)
}

type ConnectionSettingParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ConnectionSettingParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionSettingParameterResponse)(nil)).Elem()
}

func (o ConnectionSettingParameterResponseArrayOutput) ToConnectionSettingParameterResponseArrayOutput() ConnectionSettingParameterResponseArrayOutput {
	return o
}

func (o ConnectionSettingParameterResponseArrayOutput) ToConnectionSettingParameterResponseArrayOutputWithContext(ctx context.Context) ConnectionSettingParameterResponseArrayOutput {
	return o
}

func (o ConnectionSettingParameterResponseArrayOutput) Index(i pulumi.IntInput) ConnectionSettingParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ConnectionSettingParameterResponse {
		return vs[0].([]ConnectionSettingParameterResponse)[vs[1].(int)]
	}).(ConnectionSettingParameterResponseOutput)
}

type ConnectionSettingProperties struct {
	ClientId                   *string                      `pulumi:"clientId"`
	ClientSecret               *string                      `pulumi:"clientSecret"`
	Id                         *string                      `pulumi:"id"`
	Name                       *string                      `pulumi:"name"`
	Parameters                 []ConnectionSettingParameter `pulumi:"parameters"`
	ProvisioningState          *string                      `pulumi:"provisioningState"`
	Scopes                     *string                      `pulumi:"scopes"`
	ServiceProviderDisplayName *string                      `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          *string                      `pulumi:"serviceProviderId"`
}





type ConnectionSettingPropertiesInput interface {
	pulumi.Input

	ToConnectionSettingPropertiesOutput() ConnectionSettingPropertiesOutput
	ToConnectionSettingPropertiesOutputWithContext(context.Context) ConnectionSettingPropertiesOutput
}

type ConnectionSettingPropertiesArgs struct {
	ClientId                   pulumi.StringPtrInput                `pulumi:"clientId"`
	ClientSecret               pulumi.StringPtrInput                `pulumi:"clientSecret"`
	Id                         pulumi.StringPtrInput                `pulumi:"id"`
	Name                       pulumi.StringPtrInput                `pulumi:"name"`
	Parameters                 ConnectionSettingParameterArrayInput `pulumi:"parameters"`
	ProvisioningState          pulumi.StringPtrInput                `pulumi:"provisioningState"`
	Scopes                     pulumi.StringPtrInput                `pulumi:"scopes"`
	ServiceProviderDisplayName pulumi.StringPtrInput                `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          pulumi.StringPtrInput                `pulumi:"serviceProviderId"`
}

func (ConnectionSettingPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingProperties)(nil)).Elem()
}

func (i ConnectionSettingPropertiesArgs) ToConnectionSettingPropertiesOutput() ConnectionSettingPropertiesOutput {
	return i.ToConnectionSettingPropertiesOutputWithContext(context.Background())
}

func (i ConnectionSettingPropertiesArgs) ToConnectionSettingPropertiesOutputWithContext(ctx context.Context) ConnectionSettingPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingPropertiesOutput)
}

func (i ConnectionSettingPropertiesArgs) ToConnectionSettingPropertiesPtrOutput() ConnectionSettingPropertiesPtrOutput {
	return i.ToConnectionSettingPropertiesPtrOutputWithContext(context.Background())
}

func (i ConnectionSettingPropertiesArgs) ToConnectionSettingPropertiesPtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingPropertiesOutput).ToConnectionSettingPropertiesPtrOutputWithContext(ctx)
}









type ConnectionSettingPropertiesPtrInput interface {
	pulumi.Input

	ToConnectionSettingPropertiesPtrOutput() ConnectionSettingPropertiesPtrOutput
	ToConnectionSettingPropertiesPtrOutputWithContext(context.Context) ConnectionSettingPropertiesPtrOutput
}

type connectionSettingPropertiesPtrType ConnectionSettingPropertiesArgs

func ConnectionSettingPropertiesPtr(v *ConnectionSettingPropertiesArgs) ConnectionSettingPropertiesPtrInput {
	return (*connectionSettingPropertiesPtrType)(v)
}

func (*connectionSettingPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionSettingProperties)(nil)).Elem()
}

func (i *connectionSettingPropertiesPtrType) ToConnectionSettingPropertiesPtrOutput() ConnectionSettingPropertiesPtrOutput {
	return i.ToConnectionSettingPropertiesPtrOutputWithContext(context.Background())
}

func (i *connectionSettingPropertiesPtrType) ToConnectionSettingPropertiesPtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingPropertiesPtrOutput)
}

type ConnectionSettingPropertiesOutput struct{ *pulumi.OutputState }

func (ConnectionSettingPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingProperties)(nil)).Elem()
}

func (o ConnectionSettingPropertiesOutput) ToConnectionSettingPropertiesOutput() ConnectionSettingPropertiesOutput {
	return o
}

func (o ConnectionSettingPropertiesOutput) ToConnectionSettingPropertiesOutputWithContext(ctx context.Context) ConnectionSettingPropertiesOutput {
	return o
}

func (o ConnectionSettingPropertiesOutput) ToConnectionSettingPropertiesPtrOutput() ConnectionSettingPropertiesPtrOutput {
	return o.ToConnectionSettingPropertiesPtrOutputWithContext(context.Background())
}

func (o ConnectionSettingPropertiesOutput) ToConnectionSettingPropertiesPtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionSettingProperties) *ConnectionSettingProperties {
		return &v
	}).(ConnectionSettingPropertiesPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) Parameters() ConnectionSettingParameterArrayOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) []ConnectionSettingParameter { return v.Parameters }).(ConnectionSettingParameterArrayOutput)
}

func (o ConnectionSettingPropertiesOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) Scopes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.Scopes }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) ServiceProviderDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.ServiceProviderDisplayName }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesOutput) ServiceProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) *string { return v.ServiceProviderId }).(pulumi.StringPtrOutput)
}

type ConnectionSettingPropertiesPtrOutput struct{ *pulumi.OutputState }

func (ConnectionSettingPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionSettingProperties)(nil)).Elem()
}

func (o ConnectionSettingPropertiesPtrOutput) ToConnectionSettingPropertiesPtrOutput() ConnectionSettingPropertiesPtrOutput {
	return o
}

func (o ConnectionSettingPropertiesPtrOutput) ToConnectionSettingPropertiesPtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesPtrOutput {
	return o
}

func (o ConnectionSettingPropertiesPtrOutput) Elem() ConnectionSettingPropertiesOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) ConnectionSettingProperties {
		if v != nil {
			return *v
		}
		var ret ConnectionSettingProperties
		return ret
	}).(ConnectionSettingPropertiesOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.Name
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) Parameters() ConnectionSettingParameterArrayOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) []ConnectionSettingParameter {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ConnectionSettingParameterArrayOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.ProvisioningState
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) Scopes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) ServiceProviderDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesPtrOutput) ServiceProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingProperties) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderId
	}).(pulumi.StringPtrOutput)
}

type ConnectionSettingPropertiesResponse struct {
	ClientId                   *string                              `pulumi:"clientId"`
	ClientSecret               *string                              `pulumi:"clientSecret"`
	Id                         *string                              `pulumi:"id"`
	Name                       *string                              `pulumi:"name"`
	Parameters                 []ConnectionSettingParameterResponse `pulumi:"parameters"`
	ProvisioningState          *string                              `pulumi:"provisioningState"`
	Scopes                     *string                              `pulumi:"scopes"`
	ServiceProviderDisplayName *string                              `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          *string                              `pulumi:"serviceProviderId"`
	SettingId                  string                               `pulumi:"settingId"`
}

type ConnectionSettingPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ConnectionSettingPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingPropertiesResponse)(nil)).Elem()
}

func (o ConnectionSettingPropertiesResponseOutput) ToConnectionSettingPropertiesResponseOutput() ConnectionSettingPropertiesResponseOutput {
	return o
}

func (o ConnectionSettingPropertiesResponseOutput) ToConnectionSettingPropertiesResponseOutputWithContext(ctx context.Context) ConnectionSettingPropertiesResponseOutput {
	return o
}

func (o ConnectionSettingPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.Id }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) Parameters() ConnectionSettingParameterResponseArrayOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) []ConnectionSettingParameterResponse { return v.Parameters }).(ConnectionSettingParameterResponseArrayOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) ProvisioningState() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.ProvisioningState }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) Scopes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.Scopes }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) ServiceProviderDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.ServiceProviderDisplayName }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) ServiceProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) *string { return v.ServiceProviderId }).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponseOutput) SettingId() pulumi.StringOutput {
	return o.ApplyT(func(v ConnectionSettingPropertiesResponse) string { return v.SettingId }).(pulumi.StringOutput)
}

type DirectLineChannel struct {
	ChannelName string                       `pulumi:"channelName"`
	Etag        *string                      `pulumi:"etag"`
	Location    *string                      `pulumi:"location"`
	Properties  *DirectLineChannelProperties `pulumi:"properties"`
}

type DirectLineChannelProperties struct {
	DirectLineEmbedCode *string          `pulumi:"directLineEmbedCode"`
	Sites               []DirectLineSite `pulumi:"sites"`
}

type DirectLineChannelPropertiesResponse struct {
	DirectLineEmbedCode *string                  `pulumi:"directLineEmbedCode"`
	Sites               []DirectLineSiteResponse `pulumi:"sites"`
}

type DirectLineChannelResponse struct {
	ChannelName       string                               `pulumi:"channelName"`
	Etag              *string                              `pulumi:"etag"`
	Location          *string                              `pulumi:"location"`
	Properties        *DirectLineChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                               `pulumi:"provisioningState"`
}

type DirectLineSite struct {
	IsBlockUserUploadEnabled *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsEnabled                bool     `pulumi:"isEnabled"`
	IsSecureSiteEnabled      *bool    `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled              bool     `pulumi:"isV1Enabled"`
	IsV3Enabled              bool     `pulumi:"isV3Enabled"`
	SiteName                 string   `pulumi:"siteName"`
	TrustedOrigins           []string `pulumi:"trustedOrigins"`
}

type DirectLineSiteResponse struct {
	IsBlockUserUploadEnabled *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsEnabled                bool     `pulumi:"isEnabled"`
	IsSecureSiteEnabled      *bool    `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled              bool     `pulumi:"isV1Enabled"`
	IsV3Enabled              bool     `pulumi:"isV3Enabled"`
	Key                      string   `pulumi:"key"`
	Key2                     string   `pulumi:"key2"`
	SiteId                   string   `pulumi:"siteId"`
	SiteName                 string   `pulumi:"siteName"`
	TrustedOrigins           []string `pulumi:"trustedOrigins"`
}

type DirectLineSpeechChannel struct {
	ChannelName string                             `pulumi:"channelName"`
	Etag        *string                            `pulumi:"etag"`
	Location    *string                            `pulumi:"location"`
	Properties  *DirectLineSpeechChannelProperties `pulumi:"properties"`
}

type DirectLineSpeechChannelProperties struct {
	CognitiveServiceRegion          string  `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey string  `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             *string `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         *string `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    *bool   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       *bool   `pulumi:"isEnabled"`
}

type DirectLineSpeechChannelPropertiesResponse struct {
	CognitiveServiceRegion          string  `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey string  `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             *string `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         *string `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    *bool   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       *bool   `pulumi:"isEnabled"`
}

type DirectLineSpeechChannelResponse struct {
	ChannelName       string                                     `pulumi:"channelName"`
	Etag              *string                                    `pulumi:"etag"`
	Location          *string                                    `pulumi:"location"`
	Properties        *DirectLineSpeechChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                                     `pulumi:"provisioningState"`
}

type EmailChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *EmailChannelProperties `pulumi:"properties"`
}

type EmailChannelProperties struct {
	EmailAddress string  `pulumi:"emailAddress"`
	IsEnabled    bool    `pulumi:"isEnabled"`
	Password     *string `pulumi:"password"`
}

type EmailChannelPropertiesResponse struct {
	EmailAddress string  `pulumi:"emailAddress"`
	IsEnabled    bool    `pulumi:"isEnabled"`
	Password     *string `pulumi:"password"`
}

type EmailChannelResponse struct {
	ChannelName       string                          `pulumi:"channelName"`
	Etag              *string                         `pulumi:"etag"`
	Location          *string                         `pulumi:"location"`
	Properties        *EmailChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                          `pulumi:"provisioningState"`
}

type EnterpriseChannelNode struct {
	AzureLocation string  `pulumi:"azureLocation"`
	AzureSku      string  `pulumi:"azureSku"`
	Name          string  `pulumi:"name"`
	State         *string `pulumi:"state"`
}





type EnterpriseChannelNodeInput interface {
	pulumi.Input

	ToEnterpriseChannelNodeOutput() EnterpriseChannelNodeOutput
	ToEnterpriseChannelNodeOutputWithContext(context.Context) EnterpriseChannelNodeOutput
}

type EnterpriseChannelNodeArgs struct {
	AzureLocation pulumi.StringInput    `pulumi:"azureLocation"`
	AzureSku      pulumi.StringInput    `pulumi:"azureSku"`
	Name          pulumi.StringInput    `pulumi:"name"`
	State         pulumi.StringPtrInput `pulumi:"state"`
}

func (EnterpriseChannelNodeArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelNode)(nil)).Elem()
}

func (i EnterpriseChannelNodeArgs) ToEnterpriseChannelNodeOutput() EnterpriseChannelNodeOutput {
	return i.ToEnterpriseChannelNodeOutputWithContext(context.Background())
}

func (i EnterpriseChannelNodeArgs) ToEnterpriseChannelNodeOutputWithContext(ctx context.Context) EnterpriseChannelNodeOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseChannelNodeOutput)
}





type EnterpriseChannelNodeArrayInput interface {
	pulumi.Input

	ToEnterpriseChannelNodeArrayOutput() EnterpriseChannelNodeArrayOutput
	ToEnterpriseChannelNodeArrayOutputWithContext(context.Context) EnterpriseChannelNodeArrayOutput
}

type EnterpriseChannelNodeArray []EnterpriseChannelNodeInput

func (EnterpriseChannelNodeArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnterpriseChannelNode)(nil)).Elem()
}

func (i EnterpriseChannelNodeArray) ToEnterpriseChannelNodeArrayOutput() EnterpriseChannelNodeArrayOutput {
	return i.ToEnterpriseChannelNodeArrayOutputWithContext(context.Background())
}

func (i EnterpriseChannelNodeArray) ToEnterpriseChannelNodeArrayOutputWithContext(ctx context.Context) EnterpriseChannelNodeArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseChannelNodeArrayOutput)
}

type EnterpriseChannelNodeOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelNodeOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelNode)(nil)).Elem()
}

func (o EnterpriseChannelNodeOutput) ToEnterpriseChannelNodeOutput() EnterpriseChannelNodeOutput {
	return o
}

func (o EnterpriseChannelNodeOutput) ToEnterpriseChannelNodeOutputWithContext(ctx context.Context) EnterpriseChannelNodeOutput {
	return o
}

func (o EnterpriseChannelNodeOutput) AzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNode) string { return v.AzureLocation }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeOutput) AzureSku() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNode) string { return v.AzureSku }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNode) string { return v.Name }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseChannelNode) *string { return v.State }).(pulumi.StringPtrOutput)
}

type EnterpriseChannelNodeArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelNodeArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnterpriseChannelNode)(nil)).Elem()
}

func (o EnterpriseChannelNodeArrayOutput) ToEnterpriseChannelNodeArrayOutput() EnterpriseChannelNodeArrayOutput {
	return o
}

func (o EnterpriseChannelNodeArrayOutput) ToEnterpriseChannelNodeArrayOutputWithContext(ctx context.Context) EnterpriseChannelNodeArrayOutput {
	return o
}

func (o EnterpriseChannelNodeArrayOutput) Index(i pulumi.IntInput) EnterpriseChannelNodeOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnterpriseChannelNode {
		return vs[0].([]EnterpriseChannelNode)[vs[1].(int)]
	}).(EnterpriseChannelNodeOutput)
}

type EnterpriseChannelNodeResponse struct {
	AzureLocation string  `pulumi:"azureLocation"`
	AzureSku      string  `pulumi:"azureSku"`
	Id            string  `pulumi:"id"`
	Name          string  `pulumi:"name"`
	State         *string `pulumi:"state"`
}

type EnterpriseChannelNodeResponseOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelNodeResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelNodeResponse)(nil)).Elem()
}

func (o EnterpriseChannelNodeResponseOutput) ToEnterpriseChannelNodeResponseOutput() EnterpriseChannelNodeResponseOutput {
	return o
}

func (o EnterpriseChannelNodeResponseOutput) ToEnterpriseChannelNodeResponseOutputWithContext(ctx context.Context) EnterpriseChannelNodeResponseOutput {
	return o
}

func (o EnterpriseChannelNodeResponseOutput) AzureLocation() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNodeResponse) string { return v.AzureLocation }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeResponseOutput) AzureSku() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNodeResponse) string { return v.AzureSku }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNodeResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v EnterpriseChannelNodeResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o EnterpriseChannelNodeResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseChannelNodeResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type EnterpriseChannelNodeResponseArrayOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelNodeResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]EnterpriseChannelNodeResponse)(nil)).Elem()
}

func (o EnterpriseChannelNodeResponseArrayOutput) ToEnterpriseChannelNodeResponseArrayOutput() EnterpriseChannelNodeResponseArrayOutput {
	return o
}

func (o EnterpriseChannelNodeResponseArrayOutput) ToEnterpriseChannelNodeResponseArrayOutputWithContext(ctx context.Context) EnterpriseChannelNodeResponseArrayOutput {
	return o
}

func (o EnterpriseChannelNodeResponseArrayOutput) Index(i pulumi.IntInput) EnterpriseChannelNodeResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) EnterpriseChannelNodeResponse {
		return vs[0].([]EnterpriseChannelNodeResponse)[vs[1].(int)]
	}).(EnterpriseChannelNodeResponseOutput)
}

type EnterpriseChannelProperties struct {
	Nodes []EnterpriseChannelNode `pulumi:"nodes"`
	State *string                 `pulumi:"state"`
}





type EnterpriseChannelPropertiesInput interface {
	pulumi.Input

	ToEnterpriseChannelPropertiesOutput() EnterpriseChannelPropertiesOutput
	ToEnterpriseChannelPropertiesOutputWithContext(context.Context) EnterpriseChannelPropertiesOutput
}

type EnterpriseChannelPropertiesArgs struct {
	Nodes EnterpriseChannelNodeArrayInput `pulumi:"nodes"`
	State pulumi.StringPtrInput           `pulumi:"state"`
}

func (EnterpriseChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelProperties)(nil)).Elem()
}

func (i EnterpriseChannelPropertiesArgs) ToEnterpriseChannelPropertiesOutput() EnterpriseChannelPropertiesOutput {
	return i.ToEnterpriseChannelPropertiesOutputWithContext(context.Background())
}

func (i EnterpriseChannelPropertiesArgs) ToEnterpriseChannelPropertiesOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseChannelPropertiesOutput)
}

func (i EnterpriseChannelPropertiesArgs) ToEnterpriseChannelPropertiesPtrOutput() EnterpriseChannelPropertiesPtrOutput {
	return i.ToEnterpriseChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i EnterpriseChannelPropertiesArgs) ToEnterpriseChannelPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseChannelPropertiesOutput).ToEnterpriseChannelPropertiesPtrOutputWithContext(ctx)
}









type EnterpriseChannelPropertiesPtrInput interface {
	pulumi.Input

	ToEnterpriseChannelPropertiesPtrOutput() EnterpriseChannelPropertiesPtrOutput
	ToEnterpriseChannelPropertiesPtrOutputWithContext(context.Context) EnterpriseChannelPropertiesPtrOutput
}

type enterpriseChannelPropertiesPtrType EnterpriseChannelPropertiesArgs

func EnterpriseChannelPropertiesPtr(v *EnterpriseChannelPropertiesArgs) EnterpriseChannelPropertiesPtrInput {
	return (*enterpriseChannelPropertiesPtrType)(v)
}

func (*enterpriseChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseChannelProperties)(nil)).Elem()
}

func (i *enterpriseChannelPropertiesPtrType) ToEnterpriseChannelPropertiesPtrOutput() EnterpriseChannelPropertiesPtrOutput {
	return i.ToEnterpriseChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *enterpriseChannelPropertiesPtrType) ToEnterpriseChannelPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EnterpriseChannelPropertiesPtrOutput)
}

type EnterpriseChannelPropertiesOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelProperties)(nil)).Elem()
}

func (o EnterpriseChannelPropertiesOutput) ToEnterpriseChannelPropertiesOutput() EnterpriseChannelPropertiesOutput {
	return o
}

func (o EnterpriseChannelPropertiesOutput) ToEnterpriseChannelPropertiesOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesOutput {
	return o
}

func (o EnterpriseChannelPropertiesOutput) ToEnterpriseChannelPropertiesPtrOutput() EnterpriseChannelPropertiesPtrOutput {
	return o.ToEnterpriseChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o EnterpriseChannelPropertiesOutput) ToEnterpriseChannelPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EnterpriseChannelProperties) *EnterpriseChannelProperties {
		return &v
	}).(EnterpriseChannelPropertiesPtrOutput)
}

func (o EnterpriseChannelPropertiesOutput) Nodes() EnterpriseChannelNodeArrayOutput {
	return o.ApplyT(func(v EnterpriseChannelProperties) []EnterpriseChannelNode { return v.Nodes }).(EnterpriseChannelNodeArrayOutput)
}

func (o EnterpriseChannelPropertiesOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseChannelProperties) *string { return v.State }).(pulumi.StringPtrOutput)
}

type EnterpriseChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EnterpriseChannelProperties)(nil)).Elem()
}

func (o EnterpriseChannelPropertiesPtrOutput) ToEnterpriseChannelPropertiesPtrOutput() EnterpriseChannelPropertiesPtrOutput {
	return o
}

func (o EnterpriseChannelPropertiesPtrOutput) ToEnterpriseChannelPropertiesPtrOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesPtrOutput {
	return o
}

func (o EnterpriseChannelPropertiesPtrOutput) Elem() EnterpriseChannelPropertiesOutput {
	return o.ApplyT(func(v *EnterpriseChannelProperties) EnterpriseChannelProperties {
		if v != nil {
			return *v
		}
		var ret EnterpriseChannelProperties
		return ret
	}).(EnterpriseChannelPropertiesOutput)
}

func (o EnterpriseChannelPropertiesPtrOutput) Nodes() EnterpriseChannelNodeArrayOutput {
	return o.ApplyT(func(v *EnterpriseChannelProperties) []EnterpriseChannelNode {
		if v == nil {
			return nil
		}
		return v.Nodes
	}).(EnterpriseChannelNodeArrayOutput)
}

func (o EnterpriseChannelPropertiesPtrOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EnterpriseChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.State
	}).(pulumi.StringPtrOutput)
}

type EnterpriseChannelPropertiesResponse struct {
	Nodes []EnterpriseChannelNodeResponse `pulumi:"nodes"`
	State *string                         `pulumi:"state"`
}

type EnterpriseChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EnterpriseChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EnterpriseChannelPropertiesResponse)(nil)).Elem()
}

func (o EnterpriseChannelPropertiesResponseOutput) ToEnterpriseChannelPropertiesResponseOutput() EnterpriseChannelPropertiesResponseOutput {
	return o
}

func (o EnterpriseChannelPropertiesResponseOutput) ToEnterpriseChannelPropertiesResponseOutputWithContext(ctx context.Context) EnterpriseChannelPropertiesResponseOutput {
	return o
}

func (o EnterpriseChannelPropertiesResponseOutput) Nodes() EnterpriseChannelNodeResponseArrayOutput {
	return o.ApplyT(func(v EnterpriseChannelPropertiesResponse) []EnterpriseChannelNodeResponse { return v.Nodes }).(EnterpriseChannelNodeResponseArrayOutput)
}

func (o EnterpriseChannelPropertiesResponseOutput) State() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EnterpriseChannelPropertiesResponse) *string { return v.State }).(pulumi.StringPtrOutput)
}

type FacebookChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Etag        *string                    `pulumi:"etag"`
	Location    *string                    `pulumi:"location"`
	Properties  *FacebookChannelProperties `pulumi:"properties"`
}

type FacebookChannelProperties struct {
	AppId     string         `pulumi:"appId"`
	AppSecret *string        `pulumi:"appSecret"`
	IsEnabled bool           `pulumi:"isEnabled"`
	Pages     []FacebookPage `pulumi:"pages"`
}

type FacebookChannelPropertiesResponse struct {
	AppId       string                 `pulumi:"appId"`
	AppSecret   *string                `pulumi:"appSecret"`
	CallbackUrl string                 `pulumi:"callbackUrl"`
	IsEnabled   bool                   `pulumi:"isEnabled"`
	Pages       []FacebookPageResponse `pulumi:"pages"`
	VerifyToken string                 `pulumi:"verifyToken"`
}

type FacebookChannelResponse struct {
	ChannelName       string                             `pulumi:"channelName"`
	Etag              *string                            `pulumi:"etag"`
	Location          *string                            `pulumi:"location"`
	Properties        *FacebookChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                             `pulumi:"provisioningState"`
}

type FacebookPage struct {
	AccessToken *string `pulumi:"accessToken"`
	Id          string  `pulumi:"id"`
}

type FacebookPageResponse struct {
	AccessToken *string `pulumi:"accessToken"`
	Id          string  `pulumi:"id"`
}

type KikChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Etag        *string               `pulumi:"etag"`
	Location    *string               `pulumi:"location"`
	Properties  *KikChannelProperties `pulumi:"properties"`
}

type KikChannelProperties struct {
	ApiKey      *string `pulumi:"apiKey"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	UserName    string  `pulumi:"userName"`
}

type KikChannelPropertiesResponse struct {
	ApiKey      *string `pulumi:"apiKey"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	UserName    string  `pulumi:"userName"`
}

type KikChannelResponse struct {
	ChannelName       string                        `pulumi:"channelName"`
	Etag              *string                       `pulumi:"etag"`
	Location          *string                       `pulumi:"location"`
	Properties        *KikChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                        `pulumi:"provisioningState"`
}

type LineChannel struct {
	ChannelName string                 `pulumi:"channelName"`
	Etag        *string                `pulumi:"etag"`
	Location    *string                `pulumi:"location"`
	Properties  *LineChannelProperties `pulumi:"properties"`
}

type LineChannelProperties struct {
	LineRegistrations []LineRegistration `pulumi:"lineRegistrations"`
}

type LineChannelPropertiesResponse struct {
	CallbackUrl       string                     `pulumi:"callbackUrl"`
	IsValidated       bool                       `pulumi:"isValidated"`
	LineRegistrations []LineRegistrationResponse `pulumi:"lineRegistrations"`
}

type LineChannelResponse struct {
	ChannelName       string                         `pulumi:"channelName"`
	Etag              *string                        `pulumi:"etag"`
	Location          *string                        `pulumi:"location"`
	Properties        *LineChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                         `pulumi:"provisioningState"`
}

type LineRegistration struct {
	ChannelAccessToken *string `pulumi:"channelAccessToken"`
	ChannelSecret      *string `pulumi:"channelSecret"`
}

type LineRegistrationResponse struct {
	ChannelAccessToken *string `pulumi:"channelAccessToken"`
	ChannelSecret      *string `pulumi:"channelSecret"`
	GeneratedId        string  `pulumi:"generatedId"`
}

type MsTeamsChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Etag        *string                   `pulumi:"etag"`
	Location    *string                   `pulumi:"location"`
	Properties  *MsTeamsChannelProperties `pulumi:"properties"`
}

type MsTeamsChannelProperties struct {
	AcceptedTerms         *bool   `pulumi:"acceptedTerms"`
	CallingWebHook        *string `pulumi:"callingWebHook"`
	DeploymentEnvironment *string `pulumi:"deploymentEnvironment"`
	EnableCalling         *bool   `pulumi:"enableCalling"`
	IncomingCallRoute     *string `pulumi:"incomingCallRoute"`
	IsEnabled             bool    `pulumi:"isEnabled"`
}

type MsTeamsChannelPropertiesResponse struct {
	AcceptedTerms         *bool   `pulumi:"acceptedTerms"`
	CallingWebHook        *string `pulumi:"callingWebHook"`
	DeploymentEnvironment *string `pulumi:"deploymentEnvironment"`
	EnableCalling         *bool   `pulumi:"enableCalling"`
	IncomingCallRoute     *string `pulumi:"incomingCallRoute"`
	IsEnabled             bool    `pulumi:"isEnabled"`
}

type MsTeamsChannelResponse struct {
	ChannelName       string                            `pulumi:"channelName"`
	Etag              *string                           `pulumi:"etag"`
	Location          *string                           `pulumi:"location"`
	Properties        *MsTeamsChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                            `pulumi:"provisioningState"`
}

type PrivateEndpointResponse struct {
	Id string `pulumi:"id"`
}

type PrivateEndpointResponseOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutput() PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) ToPrivateEndpointResponseOutputWithContext(ctx context.Context) PrivateEndpointResponseOutput {
	return o
}

func (o PrivateEndpointResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v PrivateEndpointResponse) string { return v.Id }).(pulumi.StringOutput)
}

type PrivateEndpointResponsePtrOutput struct{ *pulumi.OutputState }

func (PrivateEndpointResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**PrivateEndpointResponse)(nil)).Elem()
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutput() PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) ToPrivateEndpointResponsePtrOutputWithContext(ctx context.Context) PrivateEndpointResponsePtrOutput {
	return o
}

func (o PrivateEndpointResponsePtrOutput) Elem() PrivateEndpointResponseOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) PrivateEndpointResponse {
		if v != nil {
			return *v
		}
		var ret PrivateEndpointResponse
		return ret
	}).(PrivateEndpointResponseOutput)
}

func (o PrivateEndpointResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *PrivateEndpointResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionState struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}





type PrivateLinkServiceConnectionStateInput interface {
	pulumi.Input

	ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput
	ToPrivateLinkServiceConnectionStateOutputWithContext(context.Context) PrivateLinkServiceConnectionStateOutput
}

type PrivateLinkServiceConnectionStateArgs struct {
	ActionsRequired pulumi.StringPtrInput `pulumi:"actionsRequired"`
	Description     pulumi.StringPtrInput `pulumi:"description"`
	Status          pulumi.StringPtrInput `pulumi:"status"`
}

func (PrivateLinkServiceConnectionStateArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return i.ToPrivateLinkServiceConnectionStateOutputWithContext(context.Background())
}

func (i PrivateLinkServiceConnectionStateArgs) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return pulumi.ToOutputWithContext(ctx, i).(PrivateLinkServiceConnectionStateOutput)
}

type PrivateLinkServiceConnectionStateOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionState)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutput() PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ToPrivateLinkServiceConnectionStateOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionState) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type PrivateLinkServiceConnectionStateResponse struct {
	ActionsRequired *string `pulumi:"actionsRequired"`
	Description     *string `pulumi:"description"`
	Status          *string `pulumi:"status"`
}

type PrivateLinkServiceConnectionStateResponseOutput struct{ *pulumi.OutputState }

func (PrivateLinkServiceConnectionStateResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*PrivateLinkServiceConnectionStateResponse)(nil)).Elem()
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutput() PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ToPrivateLinkServiceConnectionStateResponseOutputWithContext(ctx context.Context) PrivateLinkServiceConnectionStateResponseOutput {
	return o
}

func (o PrivateLinkServiceConnectionStateResponseOutput) ActionsRequired() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.ActionsRequired }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Description }).(pulumi.StringPtrOutput)
}

func (o PrivateLinkServiceConnectionStateResponseOutput) Status() pulumi.StringPtrOutput {
	return o.ApplyT(func(v PrivateLinkServiceConnectionStateResponse) *string { return v.Status }).(pulumi.StringPtrOutput)
}

type ServiceProviderParameterResponse struct {
	Default     string                                   `pulumi:"default"`
	Description string                                   `pulumi:"description"`
	DisplayName string                                   `pulumi:"displayName"`
	HelpUrl     string                                   `pulumi:"helpUrl"`
	Metadata    ServiceProviderParameterResponseMetadata `pulumi:"metadata"`
	Name        string                                   `pulumi:"name"`
	Type        string                                   `pulumi:"type"`
}

type ServiceProviderParameterResponseConstraints struct {
	Required *bool `pulumi:"required"`
}

type ServiceProviderParameterResponseMetadata struct {
	Constraints *ServiceProviderParameterResponseConstraints `pulumi:"constraints"`
}

type ServiceProviderPropertiesResponse struct {
	DevPortalUrl        string                             `pulumi:"devPortalUrl"`
	DisplayName         string                             `pulumi:"displayName"`
	IconUrl             string                             `pulumi:"iconUrl"`
	Id                  string                             `pulumi:"id"`
	Parameters          []ServiceProviderParameterResponse `pulumi:"parameters"`
	ServiceProviderName string                             `pulumi:"serviceProviderName"`
}

type ServiceProviderResponse struct {
	Properties *ServiceProviderPropertiesResponse `pulumi:"properties"`
}

type SiteResponse struct {
	ETag                     *string  `pulumi:"eTag"`
	IsBlockUserUploadEnabled *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsEnabled                bool     `pulumi:"isEnabled"`
	IsSecureSiteEnabled      *bool    `pulumi:"isSecureSiteEnabled"`
	IsTokenEnabled           *bool    `pulumi:"isTokenEnabled"`
	IsV1Enabled              bool     `pulumi:"isV1Enabled"`
	IsV3Enabled              bool     `pulumi:"isV3Enabled"`
	IsWebchatPreviewEnabled  bool     `pulumi:"isWebchatPreviewEnabled"`
	Key                      string   `pulumi:"key"`
	Key2                     string   `pulumi:"key2"`
	SiteId                   string   `pulumi:"siteId"`
	SiteName                 string   `pulumi:"siteName"`
	TrustedOrigins           []string `pulumi:"trustedOrigins"`
}


func (val *SiteResponse) Defaults() *SiteResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebchatPreviewEnabled) {
		tmp.IsWebchatPreviewEnabled = false
	}
	return &tmp
}

type Sku struct {
	Name string `pulumi:"name"`
}





type SkuInput interface {
	pulumi.Input

	ToSkuOutput() SkuOutput
	ToSkuOutputWithContext(context.Context) SkuOutput
}

type SkuArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
}

func (SkuArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (i SkuArgs) ToSkuOutput() SkuOutput {
	return i.ToSkuOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput)
}

func (i SkuArgs) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i SkuArgs) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuOutput).ToSkuPtrOutputWithContext(ctx)
}









type SkuPtrInput interface {
	pulumi.Input

	ToSkuPtrOutput() SkuPtrOutput
	ToSkuPtrOutputWithContext(context.Context) SkuPtrOutput
}

type skuPtrType SkuArgs

func SkuPtr(v *SkuArgs) SkuPtrInput {
	return (*skuPtrType)(v)
}

func (*skuPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (i *skuPtrType) ToSkuPtrOutput() SkuPtrOutput {
	return i.ToSkuPtrOutputWithContext(context.Background())
}

func (i *skuPtrType) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuPtrOutput)
}

type SkuOutput struct{ *pulumi.OutputState }

func (SkuOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*Sku)(nil)).Elem()
}

func (o SkuOutput) ToSkuOutput() SkuOutput {
	return o
}

func (o SkuOutput) ToSkuOutputWithContext(ctx context.Context) SkuOutput {
	return o
}

func (o SkuOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o.ToSkuPtrOutputWithContext(context.Background())
}

func (o SkuOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v Sku) *Sku {
		return &v
	}).(SkuPtrOutput)
}

func (o SkuOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v Sku) string { return v.Name }).(pulumi.StringOutput)
}

type SkuPtrOutput struct{ *pulumi.OutputState }

func (SkuPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**Sku)(nil)).Elem()
}

func (o SkuPtrOutput) ToSkuPtrOutput() SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) ToSkuPtrOutputWithContext(ctx context.Context) SkuPtrOutput {
	return o
}

func (o SkuPtrOutput) Elem() SkuOutput {
	return o.ApplyT(func(v *Sku) Sku {
		if v != nil {
			return *v
		}
		var ret Sku
		return ret
	}).(SkuOutput)
}

func (o SkuPtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *Sku) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

type SkuResponse struct {
	Name string `pulumi:"name"`
	Tier string `pulumi:"tier"`
}

type SkuResponseOutput struct{ *pulumi.OutputState }

func (SkuResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (o SkuResponseOutput) ToSkuResponseOutput() SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return o
}

func (o SkuResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o SkuResponseOutput) Tier() pulumi.StringOutput {
	return o.ApplyT(func(v SkuResponse) string { return v.Tier }).(pulumi.StringOutput)
}

type SkuResponsePtrOutput struct{ *pulumi.OutputState }

func (SkuResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o
}

func (o SkuResponsePtrOutput) Elem() SkuResponseOutput {
	return o.ApplyT(func(v *SkuResponse) SkuResponse {
		if v != nil {
			return *v
		}
		var ret SkuResponse
		return ret
	}).(SkuResponseOutput)
}

func (o SkuResponsePtrOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Name
	}).(pulumi.StringPtrOutput)
}

func (o SkuResponsePtrOutput) Tier() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkuResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Tier
	}).(pulumi.StringPtrOutput)
}

type SkypeChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *SkypeChannelProperties `pulumi:"properties"`
}

type SkypeChannelProperties struct {
	CallingWebHook      *string `pulumi:"callingWebHook"`
	EnableCalling       *bool   `pulumi:"enableCalling"`
	EnableGroups        *bool   `pulumi:"enableGroups"`
	EnableMediaCards    *bool   `pulumi:"enableMediaCards"`
	EnableMessaging     *bool   `pulumi:"enableMessaging"`
	EnableScreenSharing *bool   `pulumi:"enableScreenSharing"`
	EnableVideo         *bool   `pulumi:"enableVideo"`
	GroupsMode          *string `pulumi:"groupsMode"`
	IncomingCallRoute   *string `pulumi:"incomingCallRoute"`
	IsEnabled           bool    `pulumi:"isEnabled"`
}

type SkypeChannelPropertiesResponse struct {
	CallingWebHook      *string `pulumi:"callingWebHook"`
	EnableCalling       *bool   `pulumi:"enableCalling"`
	EnableGroups        *bool   `pulumi:"enableGroups"`
	EnableMediaCards    *bool   `pulumi:"enableMediaCards"`
	EnableMessaging     *bool   `pulumi:"enableMessaging"`
	EnableScreenSharing *bool   `pulumi:"enableScreenSharing"`
	EnableVideo         *bool   `pulumi:"enableVideo"`
	GroupsMode          *string `pulumi:"groupsMode"`
	IncomingCallRoute   *string `pulumi:"incomingCallRoute"`
	IsEnabled           bool    `pulumi:"isEnabled"`
}

type SkypeChannelResponse struct {
	ChannelName       string                          `pulumi:"channelName"`
	Etag              *string                         `pulumi:"etag"`
	Location          *string                         `pulumi:"location"`
	Properties        *SkypeChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                          `pulumi:"provisioningState"`
}

type SlackChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *SlackChannelProperties `pulumi:"properties"`
}

type SlackChannelProperties struct {
	ClientId          *string `pulumi:"clientId"`
	ClientSecret      *string `pulumi:"clientSecret"`
	IsEnabled         bool    `pulumi:"isEnabled"`
	LandingPageUrl    *string `pulumi:"landingPageUrl"`
	SigningSecret     *string `pulumi:"signingSecret"`
	VerificationToken *string `pulumi:"verificationToken"`
}

type SlackChannelPropertiesResponse struct {
	ClientId                *string `pulumi:"clientId"`
	ClientSecret            *string `pulumi:"clientSecret"`
	IsEnabled               bool    `pulumi:"isEnabled"`
	IsValidated             bool    `pulumi:"isValidated"`
	LandingPageUrl          *string `pulumi:"landingPageUrl"`
	LastSubmissionId        string  `pulumi:"lastSubmissionId"`
	RedirectAction          string  `pulumi:"redirectAction"`
	RegisterBeforeOAuthFlow bool    `pulumi:"registerBeforeOAuthFlow"`
	SigningSecret           *string `pulumi:"signingSecret"`
	VerificationToken       *string `pulumi:"verificationToken"`
}

type SlackChannelResponse struct {
	ChannelName       string                          `pulumi:"channelName"`
	Etag              *string                         `pulumi:"etag"`
	Location          *string                         `pulumi:"location"`
	Properties        *SlackChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                          `pulumi:"provisioningState"`
}

type SmsChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Etag        *string               `pulumi:"etag"`
	Location    *string               `pulumi:"location"`
	Properties  *SmsChannelProperties `pulumi:"properties"`
}

type SmsChannelProperties struct {
	AccountSID  string  `pulumi:"accountSID"`
	AuthToken   *string `pulumi:"authToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	Phone       string  `pulumi:"phone"`
}

type SmsChannelPropertiesResponse struct {
	AccountSID  string  `pulumi:"accountSID"`
	AuthToken   *string `pulumi:"authToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	Phone       string  `pulumi:"phone"`
}

type SmsChannelResponse struct {
	ChannelName       string                        `pulumi:"channelName"`
	Etag              *string                       `pulumi:"etag"`
	Location          *string                       `pulumi:"location"`
	Properties        *SmsChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                        `pulumi:"provisioningState"`
}

type TelegramChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Etag        *string                    `pulumi:"etag"`
	Location    *string                    `pulumi:"location"`
	Properties  *TelegramChannelProperties `pulumi:"properties"`
}

type TelegramChannelProperties struct {
	AccessToken *string `pulumi:"accessToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
}

type TelegramChannelPropertiesResponse struct {
	AccessToken *string `pulumi:"accessToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
}

type TelegramChannelResponse struct {
	ChannelName       string                             `pulumi:"channelName"`
	Etag              *string                            `pulumi:"etag"`
	Location          *string                            `pulumi:"location"`
	Properties        *TelegramChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                             `pulumi:"provisioningState"`
}

type WebChatChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Etag        *string                   `pulumi:"etag"`
	Location    *string                   `pulumi:"location"`
	Properties  *WebChatChannelProperties `pulumi:"properties"`
}

type WebChatChannelProperties struct {
	Sites []WebChatSite `pulumi:"sites"`
}

type WebChatChannelPropertiesResponse struct {
	Sites            []WebChatSiteResponse `pulumi:"sites"`
	WebChatEmbedCode string                `pulumi:"webChatEmbedCode"`
}

type WebChatChannelResponse struct {
	ChannelName       string                            `pulumi:"channelName"`
	Etag              *string                           `pulumi:"etag"`
	Location          *string                           `pulumi:"location"`
	Properties        *WebChatChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                            `pulumi:"provisioningState"`
}

type WebChatSite struct {
	IsEnabled               bool   `pulumi:"isEnabled"`
	IsWebchatPreviewEnabled bool   `pulumi:"isWebchatPreviewEnabled"`
	SiteName                string `pulumi:"siteName"`
}


func (val *WebChatSite) Defaults() *WebChatSite {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebchatPreviewEnabled) {
		tmp.IsWebchatPreviewEnabled = false
	}
	return &tmp
}

type WebChatSiteResponse struct {
	IsEnabled               bool   `pulumi:"isEnabled"`
	IsWebchatPreviewEnabled bool   `pulumi:"isWebchatPreviewEnabled"`
	Key                     string `pulumi:"key"`
	Key2                    string `pulumi:"key2"`
	SiteId                  string `pulumi:"siteId"`
	SiteName                string `pulumi:"siteName"`
}


func (val *WebChatSiteResponse) Defaults() *WebChatSiteResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebchatPreviewEnabled) {
		tmp.IsWebchatPreviewEnabled = false
	}
	return &tmp
}
func init() {
	pulumi.RegisterOutputType(BotPropertiesOutput{})
	pulumi.RegisterOutputType(BotPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterArrayOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterResponseOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelNodeOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelNodeArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelNodeResponseOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelNodeResponseArrayOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelPropertiesOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EnterpriseChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponseOutput{})
	pulumi.RegisterOutputType(PrivateEndpointResponsePtrOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateOutput{})
	pulumi.RegisterOutputType(PrivateLinkServiceConnectionStateResponseOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
