


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AcsChatChannel struct {
	ChannelName string  `pulumi:"channelName"`
	Etag        *string `pulumi:"etag"`
	Location    *string `pulumi:"location"`
}


func (val *AcsChatChannel) Defaults() *AcsChatChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type AcsChatChannelResponse struct {
	ChannelName       string  `pulumi:"channelName"`
	Etag              *string `pulumi:"etag"`
	Location          *string `pulumi:"location"`
	ProvisioningState string  `pulumi:"provisioningState"`
}


func (val *AcsChatChannelResponse) Defaults() *AcsChatChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type AlexaChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *AlexaChannelProperties `pulumi:"properties"`
}


func (val *AlexaChannel) Defaults() *AlexaChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *AlexaChannelResponse) Defaults() *AlexaChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type BotProperties struct {
	AllSettings                       map[string]string `pulumi:"allSettings"`
	AppPasswordHint                   *string           `pulumi:"appPasswordHint"`
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
	TenantId                          *string           `pulumi:"tenantId"`
}


func (val *BotProperties) Defaults() *BotProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IconUrl) {
		iconUrl_ := ""
		tmp.IconUrl = &iconUrl_
	}
	if isZero(tmp.IsCmekEnabled) {
		isCmekEnabled_ := false
		tmp.IsCmekEnabled = &isCmekEnabled_
	}
	if isZero(tmp.IsStreamingSupported) {
		isStreamingSupported_ := false
		tmp.IsStreamingSupported = &isStreamingSupported_
	}
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
	TenantId                          pulumi.StringPtrInput   `pulumi:"tenantId"`
}


func (val *BotPropertiesArgs) Defaults() *BotPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IconUrl) {
		tmp.IconUrl = pulumi.StringPtr("")
	}
	if isZero(tmp.IsCmekEnabled) {
		tmp.IsCmekEnabled = pulumi.BoolPtr(false)
	}
	if isZero(tmp.IsStreamingSupported) {
		tmp.IsStreamingSupported = pulumi.BoolPtr(false)
	}
	if isZero(tmp.PublicNetworkAccess) {
		tmp.PublicNetworkAccess = pulumi.StringPtr("Enabled")
	}
	return &tmp
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

func (o BotPropertiesOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.TenantId }).(pulumi.StringPtrOutput)
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

func (o BotPropertiesPtrOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.TenantId
	}).(pulumi.StringPtrOutput)
}

type BotPropertiesResponse struct {
	AllSettings                       map[string]string `pulumi:"allSettings"`
	AppPasswordHint                   *string           `pulumi:"appPasswordHint"`
	CmekEncryptionStatus              string            `pulumi:"cmekEncryptionStatus"`
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
	IsDeveloperAppInsightsApiKeySet   bool              `pulumi:"isDeveloperAppInsightsApiKeySet"`
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
	TenantId                          *string           `pulumi:"tenantId"`
}


func (val *BotPropertiesResponse) Defaults() *BotPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IconUrl) {
		iconUrl_ := ""
		tmp.IconUrl = &iconUrl_
	}
	if isZero(tmp.IsCmekEnabled) {
		isCmekEnabled_ := false
		tmp.IsCmekEnabled = &isCmekEnabled_
	}
	if isZero(tmp.IsStreamingSupported) {
		isStreamingSupported_ := false
		tmp.IsStreamingSupported = &isStreamingSupported_
	}
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

func (o BotPropertiesResponseOutput) CmekEncryptionStatus() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.CmekEncryptionStatus }).(pulumi.StringOutput)
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

func (o BotPropertiesResponseOutput) IsDeveloperAppInsightsApiKeySet() pulumi.BoolOutput {
	return o.ApplyT(func(v BotPropertiesResponse) bool { return v.IsDeveloperAppInsightsApiKeySet }).(pulumi.BoolOutput)
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

func (o BotPropertiesResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

type ChannelSettingsResponse struct {
	BotIconUrl            *string        `pulumi:"botIconUrl"`
	BotId                 *string        `pulumi:"botId"`
	ChannelDisplayName    *string        `pulumi:"channelDisplayName"`
	ChannelId             *string        `pulumi:"channelId"`
	DisableLocalAuth      *bool          `pulumi:"disableLocalAuth"`
	ExtensionKey1         *string        `pulumi:"extensionKey1"`
	ExtensionKey2         *string        `pulumi:"extensionKey2"`
	IsEnabled             *bool          `pulumi:"isEnabled"`
	RequireTermsAgreement *bool          `pulumi:"requireTermsAgreement"`
	Sites                 []SiteResponse `pulumi:"sites"`
}


func (val *ChannelSettingsResponse) Defaults() *ChannelSettingsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExtensionKey1) {
		extensionKey1_ := ""
		tmp.ExtensionKey1 = &extensionKey1_
	}
	if isZero(tmp.ExtensionKey2) {
		extensionKey2_ := ""
		tmp.ExtensionKey2 = &extensionKey2_
	}
	return &tmp
}

type ChannelSettingsResponseOutput struct{ *pulumi.OutputState }

func (ChannelSettingsResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ChannelSettingsResponse)(nil)).Elem()
}

func (o ChannelSettingsResponseOutput) ToChannelSettingsResponseOutput() ChannelSettingsResponseOutput {
	return o
}

func (o ChannelSettingsResponseOutput) ToChannelSettingsResponseOutputWithContext(ctx context.Context) ChannelSettingsResponseOutput {
	return o
}

func (o ChannelSettingsResponseOutput) BotIconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *string { return v.BotIconUrl }).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponseOutput) BotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *string { return v.BotId }).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponseOutput) ChannelDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *string { return v.ChannelDisplayName }).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponseOutput) ChannelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *string { return v.ChannelId }).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponseOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *bool { return v.DisableLocalAuth }).(pulumi.BoolPtrOutput)
}

func (o ChannelSettingsResponseOutput) ExtensionKey1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *string { return v.ExtensionKey1 }).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponseOutput) ExtensionKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *string { return v.ExtensionKey2 }).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

func (o ChannelSettingsResponseOutput) RequireTermsAgreement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) *bool { return v.RequireTermsAgreement }).(pulumi.BoolPtrOutput)
}

func (o ChannelSettingsResponseOutput) Sites() SiteResponseArrayOutput {
	return o.ApplyT(func(v ChannelSettingsResponse) []SiteResponse { return v.Sites }).(SiteResponseArrayOutput)
}

type ChannelSettingsResponsePtrOutput struct{ *pulumi.OutputState }

func (ChannelSettingsResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ChannelSettingsResponse)(nil)).Elem()
}

func (o ChannelSettingsResponsePtrOutput) ToChannelSettingsResponsePtrOutput() ChannelSettingsResponsePtrOutput {
	return o
}

func (o ChannelSettingsResponsePtrOutput) ToChannelSettingsResponsePtrOutputWithContext(ctx context.Context) ChannelSettingsResponsePtrOutput {
	return o
}

func (o ChannelSettingsResponsePtrOutput) Elem() ChannelSettingsResponseOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) ChannelSettingsResponse {
		if v != nil {
			return *v
		}
		var ret ChannelSettingsResponse
		return ret
	}).(ChannelSettingsResponseOutput)
}

func (o ChannelSettingsResponsePtrOutput) BotIconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BotIconUrl
	}).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) BotId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.BotId
	}).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) ChannelDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChannelDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) ChannelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ChannelId
	}).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) DisableLocalAuth() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.DisableLocalAuth
	}).(pulumi.BoolPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) ExtensionKey1() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionKey1
	}).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) ExtensionKey2() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *string {
		if v == nil {
			return nil
		}
		return v.ExtensionKey2
	}).(pulumi.StringPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) RequireTermsAgreement() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) *bool {
		if v == nil {
			return nil
		}
		return v.RequireTermsAgreement
	}).(pulumi.BoolPtrOutput)
}

func (o ChannelSettingsResponsePtrOutput) Sites() SiteResponseArrayOutput {
	return o.ApplyT(func(v *ChannelSettingsResponse) []SiteResponse {
		if v == nil {
			return nil
		}
		return v.Sites
	}).(SiteResponseArrayOutput)
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
	Parameters                 []ConnectionSettingParameter `pulumi:"parameters"`
	ProvisioningState          *string                      `pulumi:"provisioningState"`
	Scopes                     *string                      `pulumi:"scopes"`
	ServiceProviderDisplayName *string                      `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          *string                      `pulumi:"serviceProviderId"`
}


func (val *ConnectionSettingProperties) Defaults() *ConnectionSettingProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Scopes) {
		scopes_ := ""
		tmp.Scopes = &scopes_
	}
	return &tmp
}





type ConnectionSettingPropertiesInput interface {
	pulumi.Input

	ToConnectionSettingPropertiesOutput() ConnectionSettingPropertiesOutput
	ToConnectionSettingPropertiesOutputWithContext(context.Context) ConnectionSettingPropertiesOutput
}

type ConnectionSettingPropertiesArgs struct {
	ClientId                   pulumi.StringPtrInput                `pulumi:"clientId"`
	ClientSecret               pulumi.StringPtrInput                `pulumi:"clientSecret"`
	Parameters                 ConnectionSettingParameterArrayInput `pulumi:"parameters"`
	ProvisioningState          pulumi.StringPtrInput                `pulumi:"provisioningState"`
	Scopes                     pulumi.StringPtrInput                `pulumi:"scopes"`
	ServiceProviderDisplayName pulumi.StringPtrInput                `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          pulumi.StringPtrInput                `pulumi:"serviceProviderId"`
}


func (val *ConnectionSettingPropertiesArgs) Defaults() *ConnectionSettingPropertiesArgs {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Scopes) {
		tmp.Scopes = pulumi.StringPtr("")
	}
	return &tmp
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
	Parameters                 []ConnectionSettingParameterResponse `pulumi:"parameters"`
	ProvisioningState          *string                              `pulumi:"provisioningState"`
	Scopes                     *string                              `pulumi:"scopes"`
	ServiceProviderDisplayName *string                              `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          *string                              `pulumi:"serviceProviderId"`
	SettingId                  string                               `pulumi:"settingId"`
}


func (val *ConnectionSettingPropertiesResponse) Defaults() *ConnectionSettingPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Scopes) {
		scopes_ := ""
		tmp.Scopes = &scopes_
	}
	return &tmp
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


func (val *DirectLineChannel) Defaults() *DirectLineChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type DirectLineChannelProperties struct {
	DirectLineEmbedCode *string          `pulumi:"directLineEmbedCode"`
	ExtensionKey1       *string          `pulumi:"extensionKey1"`
	ExtensionKey2       *string          `pulumi:"extensionKey2"`
	Sites               []DirectLineSite `pulumi:"sites"`
}


func (val *DirectLineChannelProperties) Defaults() *DirectLineChannelProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExtensionKey1) {
		extensionKey1_ := ""
		tmp.ExtensionKey1 = &extensionKey1_
	}
	if isZero(tmp.ExtensionKey2) {
		extensionKey2_ := ""
		tmp.ExtensionKey2 = &extensionKey2_
	}
	return &tmp
}

type DirectLineChannelPropertiesResponse struct {
	DirectLineEmbedCode *string                  `pulumi:"directLineEmbedCode"`
	ExtensionKey1       *string                  `pulumi:"extensionKey1"`
	ExtensionKey2       *string                  `pulumi:"extensionKey2"`
	Sites               []DirectLineSiteResponse `pulumi:"sites"`
}


func (val *DirectLineChannelPropertiesResponse) Defaults() *DirectLineChannelPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.ExtensionKey1) {
		extensionKey1_ := ""
		tmp.ExtensionKey1 = &extensionKey1_
	}
	if isZero(tmp.ExtensionKey2) {
		extensionKey2_ := ""
		tmp.ExtensionKey2 = &extensionKey2_
	}
	return &tmp
}

type DirectLineChannelResponse struct {
	ChannelName       string                               `pulumi:"channelName"`
	Etag              *string                              `pulumi:"etag"`
	Location          *string                              `pulumi:"location"`
	Properties        *DirectLineChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                               `pulumi:"provisioningState"`
}


func (val *DirectLineChannelResponse) Defaults() *DirectLineChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type DirectLineSite struct {
	AppId                       *string  `pulumi:"appId"`
	ETag                        *string  `pulumi:"eTag"`
	IsBlockUserUploadEnabled    *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsDetailedLoggingEnabled    *bool    `pulumi:"isDetailedLoggingEnabled"`
	IsEnabled                   bool     `pulumi:"isEnabled"`
	IsEndpointParametersEnabled *bool    `pulumi:"isEndpointParametersEnabled"`
	IsNoStorageEnabled          *bool    `pulumi:"isNoStorageEnabled"`
	IsSecureSiteEnabled         *bool    `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled                 bool     `pulumi:"isV1Enabled"`
	IsV3Enabled                 bool     `pulumi:"isV3Enabled"`
	IsWebChatSpeechEnabled      *bool    `pulumi:"isWebChatSpeechEnabled"`
	IsWebchatPreviewEnabled     *bool    `pulumi:"isWebchatPreviewEnabled"`
	SiteName                    string   `pulumi:"siteName"`
	TenantId                    *string  `pulumi:"tenantId"`
	TrustedOrigins              []string `pulumi:"trustedOrigins"`
}


func (val *DirectLineSite) Defaults() *DirectLineSite {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebChatSpeechEnabled) {
		isWebChatSpeechEnabled_ := false
		tmp.IsWebChatSpeechEnabled = &isWebChatSpeechEnabled_
	}
	if isZero(tmp.IsWebchatPreviewEnabled) {
		isWebchatPreviewEnabled_ := false
		tmp.IsWebchatPreviewEnabled = &isWebchatPreviewEnabled_
	}
	return &tmp
}

type DirectLineSiteResponse struct {
	AppId                       *string  `pulumi:"appId"`
	ETag                        *string  `pulumi:"eTag"`
	IsBlockUserUploadEnabled    *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsDetailedLoggingEnabled    *bool    `pulumi:"isDetailedLoggingEnabled"`
	IsEnabled                   bool     `pulumi:"isEnabled"`
	IsEndpointParametersEnabled *bool    `pulumi:"isEndpointParametersEnabled"`
	IsNoStorageEnabled          *bool    `pulumi:"isNoStorageEnabled"`
	IsSecureSiteEnabled         *bool    `pulumi:"isSecureSiteEnabled"`
	IsTokenEnabled              bool     `pulumi:"isTokenEnabled"`
	IsV1Enabled                 bool     `pulumi:"isV1Enabled"`
	IsV3Enabled                 bool     `pulumi:"isV3Enabled"`
	IsWebChatSpeechEnabled      *bool    `pulumi:"isWebChatSpeechEnabled"`
	IsWebchatPreviewEnabled     *bool    `pulumi:"isWebchatPreviewEnabled"`
	Key                         string   `pulumi:"key"`
	Key2                        string   `pulumi:"key2"`
	SiteId                      string   `pulumi:"siteId"`
	SiteName                    string   `pulumi:"siteName"`
	TenantId                    *string  `pulumi:"tenantId"`
	TrustedOrigins              []string `pulumi:"trustedOrigins"`
}


func (val *DirectLineSiteResponse) Defaults() *DirectLineSiteResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebChatSpeechEnabled) {
		isWebChatSpeechEnabled_ := false
		tmp.IsWebChatSpeechEnabled = &isWebChatSpeechEnabled_
	}
	if isZero(tmp.IsWebchatPreviewEnabled) {
		isWebchatPreviewEnabled_ := false
		tmp.IsWebchatPreviewEnabled = &isWebchatPreviewEnabled_
	}
	return &tmp
}

type DirectLineSpeechChannel struct {
	ChannelName string                             `pulumi:"channelName"`
	Etag        *string                            `pulumi:"etag"`
	Location    *string                            `pulumi:"location"`
	Properties  *DirectLineSpeechChannelProperties `pulumi:"properties"`
}


func (val *DirectLineSpeechChannel) Defaults() *DirectLineSpeechChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type DirectLineSpeechChannelProperties struct {
	CognitiveServiceRegion          *string `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceResourceId      *string `pulumi:"cognitiveServiceResourceId"`
	CognitiveServiceSubscriptionKey *string `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             *string `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         *string `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    *bool   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       *bool   `pulumi:"isEnabled"`
}

type DirectLineSpeechChannelPropertiesResponse struct {
	CognitiveServiceRegion          *string `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceResourceId      *string `pulumi:"cognitiveServiceResourceId"`
	CognitiveServiceSubscriptionKey *string `pulumi:"cognitiveServiceSubscriptionKey"`
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


func (val *DirectLineSpeechChannelResponse) Defaults() *DirectLineSpeechChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type EmailChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *EmailChannelProperties `pulumi:"properties"`
}


func (val *EmailChannel) Defaults() *EmailChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type EmailChannelProperties struct {
	AuthMethod   *float64 `pulumi:"authMethod"`
	EmailAddress string   `pulumi:"emailAddress"`
	IsEnabled    bool     `pulumi:"isEnabled"`
	MagicCode    *string  `pulumi:"magicCode"`
	Password     *string  `pulumi:"password"`
}

type EmailChannelPropertiesResponse struct {
	AuthMethod   *float64 `pulumi:"authMethod"`
	EmailAddress string   `pulumi:"emailAddress"`
	IsEnabled    bool     `pulumi:"isEnabled"`
	MagicCode    *string  `pulumi:"magicCode"`
	Password     *string  `pulumi:"password"`
}

type EmailChannelResponse struct {
	ChannelName       string                          `pulumi:"channelName"`
	Etag              *string                         `pulumi:"etag"`
	Location          *string                         `pulumi:"location"`
	Properties        *EmailChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                          `pulumi:"provisioningState"`
}


func (val *EmailChannelResponse) Defaults() *EmailChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type FacebookChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Etag        *string                    `pulumi:"etag"`
	Location    *string                    `pulumi:"location"`
	Properties  *FacebookChannelProperties `pulumi:"properties"`
}


func (val *FacebookChannel) Defaults() *FacebookChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *FacebookChannelResponse) Defaults() *FacebookChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *KikChannel) Defaults() *KikChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *KikChannelResponse) Defaults() *KikChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type LineChannel struct {
	ChannelName string                 `pulumi:"channelName"`
	Etag        *string                `pulumi:"etag"`
	Location    *string                `pulumi:"location"`
	Properties  *LineChannelProperties `pulumi:"properties"`
}


func (val *LineChannel) Defaults() *LineChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *LineChannelResponse) Defaults() *LineChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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

type M365Extensions struct {
	ChannelName string  `pulumi:"channelName"`
	Etag        *string `pulumi:"etag"`
	Location    *string `pulumi:"location"`
}


func (val *M365Extensions) Defaults() *M365Extensions {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type M365ExtensionsResponse struct {
	ChannelName       string  `pulumi:"channelName"`
	Etag              *string `pulumi:"etag"`
	Location          *string `pulumi:"location"`
	ProvisioningState string  `pulumi:"provisioningState"`
}


func (val *M365ExtensionsResponse) Defaults() *M365ExtensionsResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type MsTeamsChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Etag        *string                   `pulumi:"etag"`
	Location    *string                   `pulumi:"location"`
	Properties  *MsTeamsChannelProperties `pulumi:"properties"`
}


func (val *MsTeamsChannel) Defaults() *MsTeamsChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type MsTeamsChannelProperties struct {
	AcceptedTerms         *bool   `pulumi:"acceptedTerms"`
	CallingWebhook        *string `pulumi:"callingWebhook"`
	DeploymentEnvironment *string `pulumi:"deploymentEnvironment"`
	EnableCalling         *bool   `pulumi:"enableCalling"`
	IncomingCallRoute     *string `pulumi:"incomingCallRoute"`
	IsEnabled             bool    `pulumi:"isEnabled"`
}


func (val *MsTeamsChannelProperties) Defaults() *MsTeamsChannelProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DeploymentEnvironment) {
		deploymentEnvironment_ := "FallbackDeploymentEnvironment"
		tmp.DeploymentEnvironment = &deploymentEnvironment_
	}
	if isZero(tmp.EnableCalling) {
		enableCalling_ := false
		tmp.EnableCalling = &enableCalling_
	}
	return &tmp
}

type MsTeamsChannelPropertiesResponse struct {
	AcceptedTerms         *bool   `pulumi:"acceptedTerms"`
	CallingWebhook        *string `pulumi:"callingWebhook"`
	DeploymentEnvironment *string `pulumi:"deploymentEnvironment"`
	EnableCalling         *bool   `pulumi:"enableCalling"`
	IncomingCallRoute     *string `pulumi:"incomingCallRoute"`
	IsEnabled             bool    `pulumi:"isEnabled"`
}


func (val *MsTeamsChannelPropertiesResponse) Defaults() *MsTeamsChannelPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.DeploymentEnvironment) {
		deploymentEnvironment_ := "FallbackDeploymentEnvironment"
		tmp.DeploymentEnvironment = &deploymentEnvironment_
	}
	if isZero(tmp.EnableCalling) {
		enableCalling_ := false
		tmp.EnableCalling = &enableCalling_
	}
	return &tmp
}

type MsTeamsChannelResponse struct {
	ChannelName       string                            `pulumi:"channelName"`
	Etag              *string                           `pulumi:"etag"`
	Location          *string                           `pulumi:"location"`
	Properties        *MsTeamsChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                            `pulumi:"provisioningState"`
}


func (val *MsTeamsChannelResponse) Defaults() *MsTeamsChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type Omnichannel struct {
	ChannelName string  `pulumi:"channelName"`
	Etag        *string `pulumi:"etag"`
	Location    *string `pulumi:"location"`
}


func (val *Omnichannel) Defaults() *Omnichannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type OmnichannelResponse struct {
	ChannelName       string  `pulumi:"channelName"`
	Etag              *string `pulumi:"etag"`
	Location          *string `pulumi:"location"`
	ProvisioningState string  `pulumi:"provisioningState"`
}


func (val *OmnichannelResponse) Defaults() *OmnichannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type OutlookChannel struct {
	ChannelName string  `pulumi:"channelName"`
	Etag        *string `pulumi:"etag"`
	Location    *string `pulumi:"location"`
}


func (val *OutlookChannel) Defaults() *OutlookChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type OutlookChannelResponse struct {
	ChannelName       string  `pulumi:"channelName"`
	Etag              *string `pulumi:"etag"`
	Location          *string `pulumi:"location"`
	ProvisioningState string  `pulumi:"provisioningState"`
}


func (val *OutlookChannelResponse) Defaults() *OutlookChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type SearchAssistant struct {
	ChannelName string  `pulumi:"channelName"`
	Etag        *string `pulumi:"etag"`
	Location    *string `pulumi:"location"`
}


func (val *SearchAssistant) Defaults() *SearchAssistant {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type SearchAssistantResponse struct {
	ChannelName       string  `pulumi:"channelName"`
	Etag              *string `pulumi:"etag"`
	Location          *string `pulumi:"location"`
	ProvisioningState string  `pulumi:"provisioningState"`
}


func (val *SearchAssistantResponse) Defaults() *SearchAssistantResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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
	IconUrl             *string                            `pulumi:"iconUrl"`
	Id                  string                             `pulumi:"id"`
	Parameters          []ServiceProviderParameterResponse `pulumi:"parameters"`
	ServiceProviderName string                             `pulumi:"serviceProviderName"`
}


func (val *ServiceProviderPropertiesResponse) Defaults() *ServiceProviderPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IconUrl) {
		iconUrl_ := ""
		tmp.IconUrl = &iconUrl_
	}
	return &tmp
}

type ServiceProviderResponse struct {
	Properties *ServiceProviderPropertiesResponse `pulumi:"properties"`
}


func (val *ServiceProviderResponse) Defaults() *ServiceProviderResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type SiteResponse struct {
	AppId                       *string  `pulumi:"appId"`
	ETag                        *string  `pulumi:"eTag"`
	IsBlockUserUploadEnabled    *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsDetailedLoggingEnabled    *bool    `pulumi:"isDetailedLoggingEnabled"`
	IsEnabled                   bool     `pulumi:"isEnabled"`
	IsEndpointParametersEnabled *bool    `pulumi:"isEndpointParametersEnabled"`
	IsNoStorageEnabled          *bool    `pulumi:"isNoStorageEnabled"`
	IsSecureSiteEnabled         *bool    `pulumi:"isSecureSiteEnabled"`
	IsTokenEnabled              bool     `pulumi:"isTokenEnabled"`
	IsV1Enabled                 *bool    `pulumi:"isV1Enabled"`
	IsV3Enabled                 *bool    `pulumi:"isV3Enabled"`
	IsWebChatSpeechEnabled      *bool    `pulumi:"isWebChatSpeechEnabled"`
	IsWebchatPreviewEnabled     *bool    `pulumi:"isWebchatPreviewEnabled"`
	Key                         string   `pulumi:"key"`
	Key2                        string   `pulumi:"key2"`
	SiteId                      string   `pulumi:"siteId"`
	SiteName                    string   `pulumi:"siteName"`
	TenantId                    *string  `pulumi:"tenantId"`
	TrustedOrigins              []string `pulumi:"trustedOrigins"`
}


func (val *SiteResponse) Defaults() *SiteResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebChatSpeechEnabled) {
		isWebChatSpeechEnabled_ := false
		tmp.IsWebChatSpeechEnabled = &isWebChatSpeechEnabled_
	}
	if isZero(tmp.IsWebchatPreviewEnabled) {
		isWebchatPreviewEnabled_ := false
		tmp.IsWebchatPreviewEnabled = &isWebchatPreviewEnabled_
	}
	return &tmp
}

type SiteResponseOutput struct{ *pulumi.OutputState }

func (SiteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SiteResponse)(nil)).Elem()
}

func (o SiteResponseOutput) ToSiteResponseOutput() SiteResponseOutput {
	return o
}

func (o SiteResponseOutput) ToSiteResponseOutputWithContext(ctx context.Context) SiteResponseOutput {
	return o
}

func (o SiteResponseOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteResponse) *string { return v.AppId }).(pulumi.StringPtrOutput)
}

func (o SiteResponseOutput) ETag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteResponse) *string { return v.ETag }).(pulumi.StringPtrOutput)
}

func (o SiteResponseOutput) IsBlockUserUploadEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsBlockUserUploadEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsDetailedLoggingEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsDetailedLoggingEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SiteResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o SiteResponseOutput) IsEndpointParametersEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsEndpointParametersEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsNoStorageEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsNoStorageEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsSecureSiteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsSecureSiteEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsTokenEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SiteResponse) bool { return v.IsTokenEnabled }).(pulumi.BoolOutput)
}

func (o SiteResponseOutput) IsV1Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsV1Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsV3Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsV3Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsWebChatSpeechEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsWebChatSpeechEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) IsWebchatPreviewEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SiteResponse) *bool { return v.IsWebchatPreviewEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v SiteResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o SiteResponseOutput) Key2() pulumi.StringOutput {
	return o.ApplyT(func(v SiteResponse) string { return v.Key2 }).(pulumi.StringOutput)
}

func (o SiteResponseOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v SiteResponse) string { return v.SiteId }).(pulumi.StringOutput)
}

func (o SiteResponseOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v SiteResponse) string { return v.SiteName }).(pulumi.StringOutput)
}

func (o SiteResponseOutput) TenantId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SiteResponse) *string { return v.TenantId }).(pulumi.StringPtrOutput)
}

func (o SiteResponseOutput) TrustedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v SiteResponse) []string { return v.TrustedOrigins }).(pulumi.StringArrayOutput)
}

type SiteResponseArrayOutput struct{ *pulumi.OutputState }

func (SiteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]SiteResponse)(nil)).Elem()
}

func (o SiteResponseArrayOutput) ToSiteResponseArrayOutput() SiteResponseArrayOutput {
	return o
}

func (o SiteResponseArrayOutput) ToSiteResponseArrayOutputWithContext(ctx context.Context) SiteResponseArrayOutput {
	return o
}

func (o SiteResponseArrayOutput) Index(i pulumi.IntInput) SiteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) SiteResponse {
		return vs[0].([]SiteResponse)[vs[1].(int)]
	}).(SiteResponseOutput)
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


func (val *SkypeChannel) Defaults() *SkypeChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
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


func (val *SkypeChannelProperties) Defaults() *SkypeChannelProperties {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableCalling) {
		enableCalling_ := false
		tmp.EnableCalling = &enableCalling_
	}
	return &tmp
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


func (val *SkypeChannelPropertiesResponse) Defaults() *SkypeChannelPropertiesResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.EnableCalling) {
		enableCalling_ := false
		tmp.EnableCalling = &enableCalling_
	}
	return &tmp
}

type SkypeChannelResponse struct {
	ChannelName       string                          `pulumi:"channelName"`
	Etag              *string                         `pulumi:"etag"`
	Location          *string                         `pulumi:"location"`
	Properties        *SkypeChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                          `pulumi:"provisioningState"`
}


func (val *SkypeChannelResponse) Defaults() *SkypeChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	tmp.Properties = tmp.Properties.Defaults()

	return &tmp
}

type SlackChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Etag        *string                 `pulumi:"etag"`
	Location    *string                 `pulumi:"location"`
	Properties  *SlackChannelProperties `pulumi:"properties"`
}


func (val *SlackChannel) Defaults() *SlackChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type SlackChannelProperties struct {
	ClientId          *string `pulumi:"clientId"`
	ClientSecret      *string `pulumi:"clientSecret"`
	IsEnabled         bool    `pulumi:"isEnabled"`
	LandingPageUrl    *string `pulumi:"landingPageUrl"`
	Scopes            *string `pulumi:"scopes"`
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
	Scopes                  *string `pulumi:"scopes"`
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


func (val *SlackChannelResponse) Defaults() *SlackChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type SmsChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Etag        *string               `pulumi:"etag"`
	Location    *string               `pulumi:"location"`
	Properties  *SmsChannelProperties `pulumi:"properties"`
}


func (val *SmsChannel) Defaults() *SmsChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *SmsChannelResponse) Defaults() *SmsChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type TelegramChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Etag        *string                    `pulumi:"etag"`
	Location    *string                    `pulumi:"location"`
	Properties  *TelegramChannelProperties `pulumi:"properties"`
}


func (val *TelegramChannel) Defaults() *TelegramChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *TelegramChannelResponse) Defaults() *TelegramChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type TelephonyChannel struct {
	ChannelName string                      `pulumi:"channelName"`
	Etag        *string                     `pulumi:"etag"`
	Location    *string                     `pulumi:"location"`
	Properties  *TelephonyChannelProperties `pulumi:"properties"`
}


func (val *TelephonyChannel) Defaults() *TelephonyChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type TelephonyChannelProperties struct {
	ApiConfigurations               []TelephonyChannelResourceApiConfiguration `pulumi:"apiConfigurations"`
	CognitiveServiceRegion          *string                                    `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey *string                                    `pulumi:"cognitiveServiceSubscriptionKey"`
	DefaultLocale                   *string                                    `pulumi:"defaultLocale"`
	IsEnabled                       *bool                                      `pulumi:"isEnabled"`
	PhoneNumbers                    []TelephonyPhoneNumbers                    `pulumi:"phoneNumbers"`
	PremiumSKU                      *string                                    `pulumi:"premiumSKU"`
}

type TelephonyChannelPropertiesResponse struct {
	ApiConfigurations               []TelephonyChannelResourceApiConfigurationResponse `pulumi:"apiConfigurations"`
	CognitiveServiceRegion          *string                                            `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey *string                                            `pulumi:"cognitiveServiceSubscriptionKey"`
	DefaultLocale                   *string                                            `pulumi:"defaultLocale"`
	IsEnabled                       *bool                                              `pulumi:"isEnabled"`
	PhoneNumbers                    []TelephonyPhoneNumbersResponse                    `pulumi:"phoneNumbers"`
	PremiumSKU                      *string                                            `pulumi:"premiumSKU"`
}

type TelephonyChannelResourceApiConfiguration struct {
	CognitiveServiceRegion          *string `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceResourceId      *string `pulumi:"cognitiveServiceResourceId"`
	CognitiveServiceSubscriptionKey *string `pulumi:"cognitiveServiceSubscriptionKey"`
	DefaultLocale                   *string `pulumi:"defaultLocale"`
	Id                              *string `pulumi:"id"`
	ProviderName                    *string `pulumi:"providerName"`
}

type TelephonyChannelResourceApiConfigurationResponse struct {
	CognitiveServiceRegion          *string `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceResourceId      *string `pulumi:"cognitiveServiceResourceId"`
	CognitiveServiceSubscriptionKey *string `pulumi:"cognitiveServiceSubscriptionKey"`
	DefaultLocale                   *string `pulumi:"defaultLocale"`
	Id                              *string `pulumi:"id"`
	ProviderName                    *string `pulumi:"providerName"`
}

type TelephonyChannelResponse struct {
	ChannelName       string                              `pulumi:"channelName"`
	Etag              *string                             `pulumi:"etag"`
	Location          *string                             `pulumi:"location"`
	Properties        *TelephonyChannelPropertiesResponse `pulumi:"properties"`
	ProvisioningState string                              `pulumi:"provisioningState"`
}


func (val *TelephonyChannelResponse) Defaults() *TelephonyChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type TelephonyPhoneNumbers struct {
	AcsEndpoint                     *string `pulumi:"acsEndpoint"`
	AcsResourceId                   *string `pulumi:"acsResourceId"`
	AcsSecret                       *string `pulumi:"acsSecret"`
	CognitiveServiceRegion          *string `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceResourceId      *string `pulumi:"cognitiveServiceResourceId"`
	CognitiveServiceSubscriptionKey *string `pulumi:"cognitiveServiceSubscriptionKey"`
	DefaultLocale                   *string `pulumi:"defaultLocale"`
	Id                              *string `pulumi:"id"`
	OfferType                       *string `pulumi:"offerType"`
	PhoneNumber                     *string `pulumi:"phoneNumber"`
}

type TelephonyPhoneNumbersResponse struct {
	AcsEndpoint                     *string `pulumi:"acsEndpoint"`
	AcsResourceId                   *string `pulumi:"acsResourceId"`
	AcsSecret                       *string `pulumi:"acsSecret"`
	CognitiveServiceRegion          *string `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceResourceId      *string `pulumi:"cognitiveServiceResourceId"`
	CognitiveServiceSubscriptionKey *string `pulumi:"cognitiveServiceSubscriptionKey"`
	DefaultLocale                   *string `pulumi:"defaultLocale"`
	Id                              *string `pulumi:"id"`
	OfferType                       *string `pulumi:"offerType"`
	PhoneNumber                     *string `pulumi:"phoneNumber"`
}

type WebChatChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Etag        *string                   `pulumi:"etag"`
	Location    *string                   `pulumi:"location"`
	Properties  *WebChatChannelProperties `pulumi:"properties"`
}


func (val *WebChatChannel) Defaults() *WebChatChannel {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
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


func (val *WebChatChannelResponse) Defaults() *WebChatChannelResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.Location) {
		location_ := "global"
		tmp.Location = &location_
	}
	return &tmp
}

type WebChatSite struct {
	AppId                       *string  `pulumi:"appId"`
	ETag                        *string  `pulumi:"eTag"`
	IsBlockUserUploadEnabled    *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsDetailedLoggingEnabled    *bool    `pulumi:"isDetailedLoggingEnabled"`
	IsEnabled                   bool     `pulumi:"isEnabled"`
	IsEndpointParametersEnabled *bool    `pulumi:"isEndpointParametersEnabled"`
	IsNoStorageEnabled          *bool    `pulumi:"isNoStorageEnabled"`
	IsSecureSiteEnabled         *bool    `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled                 *bool    `pulumi:"isV1Enabled"`
	IsV3Enabled                 *bool    `pulumi:"isV3Enabled"`
	IsWebChatSpeechEnabled      *bool    `pulumi:"isWebChatSpeechEnabled"`
	IsWebchatPreviewEnabled     bool     `pulumi:"isWebchatPreviewEnabled"`
	SiteName                    string   `pulumi:"siteName"`
	TenantId                    *string  `pulumi:"tenantId"`
	TrustedOrigins              []string `pulumi:"trustedOrigins"`
}


func (val *WebChatSite) Defaults() *WebChatSite {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebChatSpeechEnabled) {
		isWebChatSpeechEnabled_ := false
		tmp.IsWebChatSpeechEnabled = &isWebChatSpeechEnabled_
	}
	if isZero(tmp.IsWebchatPreviewEnabled) {
		tmp.IsWebchatPreviewEnabled = false
	}
	return &tmp
}

type WebChatSiteResponse struct {
	AppId                       *string  `pulumi:"appId"`
	ETag                        *string  `pulumi:"eTag"`
	IsBlockUserUploadEnabled    *bool    `pulumi:"isBlockUserUploadEnabled"`
	IsDetailedLoggingEnabled    *bool    `pulumi:"isDetailedLoggingEnabled"`
	IsEnabled                   bool     `pulumi:"isEnabled"`
	IsEndpointParametersEnabled *bool    `pulumi:"isEndpointParametersEnabled"`
	IsNoStorageEnabled          *bool    `pulumi:"isNoStorageEnabled"`
	IsSecureSiteEnabled         *bool    `pulumi:"isSecureSiteEnabled"`
	IsTokenEnabled              bool     `pulumi:"isTokenEnabled"`
	IsV1Enabled                 *bool    `pulumi:"isV1Enabled"`
	IsV3Enabled                 *bool    `pulumi:"isV3Enabled"`
	IsWebChatSpeechEnabled      *bool    `pulumi:"isWebChatSpeechEnabled"`
	IsWebchatPreviewEnabled     bool     `pulumi:"isWebchatPreviewEnabled"`
	Key                         string   `pulumi:"key"`
	Key2                        string   `pulumi:"key2"`
	SiteId                      string   `pulumi:"siteId"`
	SiteName                    string   `pulumi:"siteName"`
	TenantId                    *string  `pulumi:"tenantId"`
	TrustedOrigins              []string `pulumi:"trustedOrigins"`
}


func (val *WebChatSiteResponse) Defaults() *WebChatSiteResponse {
	if val == nil {
		return nil
	}
	tmp := *val
	if isZero(tmp.IsWebChatSpeechEnabled) {
		isWebChatSpeechEnabled_ := false
		tmp.IsWebChatSpeechEnabled = &isWebChatSpeechEnabled_
	}
	if isZero(tmp.IsWebchatPreviewEnabled) {
		tmp.IsWebchatPreviewEnabled = false
	}
	return &tmp
}
func init() {
	pulumi.RegisterOutputType(BotPropertiesOutput{})
	pulumi.RegisterOutputType(BotPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ChannelSettingsResponseOutput{})
	pulumi.RegisterOutputType(ChannelSettingsResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterArrayOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterResponseOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SiteResponseOutput{})
	pulumi.RegisterOutputType(SiteResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
