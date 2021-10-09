


package v20210301

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type AlexaChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Properties  *AlexaChannelProperties `pulumi:"properties"`
}





type AlexaChannelInput interface {
	pulumi.Input

	ToAlexaChannelOutput() AlexaChannelOutput
	ToAlexaChannelOutputWithContext(context.Context) AlexaChannelOutput
}

type AlexaChannelArgs struct {
	ChannelName pulumi.StringInput             `pulumi:"channelName"`
	Properties  AlexaChannelPropertiesPtrInput `pulumi:"properties"`
}

func (AlexaChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannel)(nil)).Elem()
}

func (i AlexaChannelArgs) ToAlexaChannelOutput() AlexaChannelOutput {
	return i.ToAlexaChannelOutputWithContext(context.Background())
}

func (i AlexaChannelArgs) ToAlexaChannelOutputWithContext(ctx context.Context) AlexaChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelOutput)
}

type AlexaChannelOutput struct{ *pulumi.OutputState }

func (AlexaChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannel)(nil)).Elem()
}

func (o AlexaChannelOutput) ToAlexaChannelOutput() AlexaChannelOutput {
	return o
}

func (o AlexaChannelOutput) ToAlexaChannelOutputWithContext(ctx context.Context) AlexaChannelOutput {
	return o
}

func (o AlexaChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v AlexaChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o AlexaChannelOutput) Properties() AlexaChannelPropertiesPtrOutput {
	return o.ApplyT(func(v AlexaChannel) *AlexaChannelProperties { return v.Properties }).(AlexaChannelPropertiesPtrOutput)
}

type AlexaChannelProperties struct {
	AlexaSkillId string `pulumi:"alexaSkillId"`
	IsEnabled    bool   `pulumi:"isEnabled"`
}





type AlexaChannelPropertiesInput interface {
	pulumi.Input

	ToAlexaChannelPropertiesOutput() AlexaChannelPropertiesOutput
	ToAlexaChannelPropertiesOutputWithContext(context.Context) AlexaChannelPropertiesOutput
}

type AlexaChannelPropertiesArgs struct {
	AlexaSkillId pulumi.StringInput `pulumi:"alexaSkillId"`
	IsEnabled    pulumi.BoolInput   `pulumi:"isEnabled"`
}

func (AlexaChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannelProperties)(nil)).Elem()
}

func (i AlexaChannelPropertiesArgs) ToAlexaChannelPropertiesOutput() AlexaChannelPropertiesOutput {
	return i.ToAlexaChannelPropertiesOutputWithContext(context.Background())
}

func (i AlexaChannelPropertiesArgs) ToAlexaChannelPropertiesOutputWithContext(ctx context.Context) AlexaChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelPropertiesOutput)
}

func (i AlexaChannelPropertiesArgs) ToAlexaChannelPropertiesPtrOutput() AlexaChannelPropertiesPtrOutput {
	return i.ToAlexaChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i AlexaChannelPropertiesArgs) ToAlexaChannelPropertiesPtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelPropertiesOutput).ToAlexaChannelPropertiesPtrOutputWithContext(ctx)
}









type AlexaChannelPropertiesPtrInput interface {
	pulumi.Input

	ToAlexaChannelPropertiesPtrOutput() AlexaChannelPropertiesPtrOutput
	ToAlexaChannelPropertiesPtrOutputWithContext(context.Context) AlexaChannelPropertiesPtrOutput
}

type alexaChannelPropertiesPtrType AlexaChannelPropertiesArgs

func AlexaChannelPropertiesPtr(v *AlexaChannelPropertiesArgs) AlexaChannelPropertiesPtrInput {
	return (*alexaChannelPropertiesPtrType)(v)
}

func (*alexaChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlexaChannelProperties)(nil)).Elem()
}

func (i *alexaChannelPropertiesPtrType) ToAlexaChannelPropertiesPtrOutput() AlexaChannelPropertiesPtrOutput {
	return i.ToAlexaChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *alexaChannelPropertiesPtrType) ToAlexaChannelPropertiesPtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelPropertiesPtrOutput)
}

type AlexaChannelPropertiesOutput struct{ *pulumi.OutputState }

func (AlexaChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannelProperties)(nil)).Elem()
}

func (o AlexaChannelPropertiesOutput) ToAlexaChannelPropertiesOutput() AlexaChannelPropertiesOutput {
	return o
}

func (o AlexaChannelPropertiesOutput) ToAlexaChannelPropertiesOutputWithContext(ctx context.Context) AlexaChannelPropertiesOutput {
	return o
}

func (o AlexaChannelPropertiesOutput) ToAlexaChannelPropertiesPtrOutput() AlexaChannelPropertiesPtrOutput {
	return o.ToAlexaChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o AlexaChannelPropertiesOutput) ToAlexaChannelPropertiesPtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlexaChannelProperties) *AlexaChannelProperties {
		return &v
	}).(AlexaChannelPropertiesPtrOutput)
}

func (o AlexaChannelPropertiesOutput) AlexaSkillId() pulumi.StringOutput {
	return o.ApplyT(func(v AlexaChannelProperties) string { return v.AlexaSkillId }).(pulumi.StringOutput)
}

func (o AlexaChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AlexaChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

type AlexaChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (AlexaChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlexaChannelProperties)(nil)).Elem()
}

func (o AlexaChannelPropertiesPtrOutput) ToAlexaChannelPropertiesPtrOutput() AlexaChannelPropertiesPtrOutput {
	return o
}

func (o AlexaChannelPropertiesPtrOutput) ToAlexaChannelPropertiesPtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesPtrOutput {
	return o
}

func (o AlexaChannelPropertiesPtrOutput) Elem() AlexaChannelPropertiesOutput {
	return o.ApplyT(func(v *AlexaChannelProperties) AlexaChannelProperties {
		if v != nil {
			return *v
		}
		var ret AlexaChannelProperties
		return ret
	}).(AlexaChannelPropertiesOutput)
}

func (o AlexaChannelPropertiesPtrOutput) AlexaSkillId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlexaChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AlexaSkillId
	}).(pulumi.StringPtrOutput)
}

func (o AlexaChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AlexaChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type AlexaChannelPropertiesResponse struct {
	AlexaSkillId       string `pulumi:"alexaSkillId"`
	IsEnabled          bool   `pulumi:"isEnabled"`
	ServiceEndpointUri string `pulumi:"serviceEndpointUri"`
	UrlFragment        string `pulumi:"urlFragment"`
}





type AlexaChannelPropertiesResponseInput interface {
	pulumi.Input

	ToAlexaChannelPropertiesResponseOutput() AlexaChannelPropertiesResponseOutput
	ToAlexaChannelPropertiesResponseOutputWithContext(context.Context) AlexaChannelPropertiesResponseOutput
}

type AlexaChannelPropertiesResponseArgs struct {
	AlexaSkillId       pulumi.StringInput `pulumi:"alexaSkillId"`
	IsEnabled          pulumi.BoolInput   `pulumi:"isEnabled"`
	ServiceEndpointUri pulumi.StringInput `pulumi:"serviceEndpointUri"`
	UrlFragment        pulumi.StringInput `pulumi:"urlFragment"`
}

func (AlexaChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannelPropertiesResponse)(nil)).Elem()
}

func (i AlexaChannelPropertiesResponseArgs) ToAlexaChannelPropertiesResponseOutput() AlexaChannelPropertiesResponseOutput {
	return i.ToAlexaChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i AlexaChannelPropertiesResponseArgs) ToAlexaChannelPropertiesResponseOutputWithContext(ctx context.Context) AlexaChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelPropertiesResponseOutput)
}

func (i AlexaChannelPropertiesResponseArgs) ToAlexaChannelPropertiesResponsePtrOutput() AlexaChannelPropertiesResponsePtrOutput {
	return i.ToAlexaChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i AlexaChannelPropertiesResponseArgs) ToAlexaChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelPropertiesResponseOutput).ToAlexaChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type AlexaChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToAlexaChannelPropertiesResponsePtrOutput() AlexaChannelPropertiesResponsePtrOutput
	ToAlexaChannelPropertiesResponsePtrOutputWithContext(context.Context) AlexaChannelPropertiesResponsePtrOutput
}

type alexaChannelPropertiesResponsePtrType AlexaChannelPropertiesResponseArgs

func AlexaChannelPropertiesResponsePtr(v *AlexaChannelPropertiesResponseArgs) AlexaChannelPropertiesResponsePtrInput {
	return (*alexaChannelPropertiesResponsePtrType)(v)
}

func (*alexaChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**AlexaChannelPropertiesResponse)(nil)).Elem()
}

func (i *alexaChannelPropertiesResponsePtrType) ToAlexaChannelPropertiesResponsePtrOutput() AlexaChannelPropertiesResponsePtrOutput {
	return i.ToAlexaChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *alexaChannelPropertiesResponsePtrType) ToAlexaChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelPropertiesResponsePtrOutput)
}

type AlexaChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (AlexaChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannelPropertiesResponse)(nil)).Elem()
}

func (o AlexaChannelPropertiesResponseOutput) ToAlexaChannelPropertiesResponseOutput() AlexaChannelPropertiesResponseOutput {
	return o
}

func (o AlexaChannelPropertiesResponseOutput) ToAlexaChannelPropertiesResponseOutputWithContext(ctx context.Context) AlexaChannelPropertiesResponseOutput {
	return o
}

func (o AlexaChannelPropertiesResponseOutput) ToAlexaChannelPropertiesResponsePtrOutput() AlexaChannelPropertiesResponsePtrOutput {
	return o.ToAlexaChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o AlexaChannelPropertiesResponseOutput) ToAlexaChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v AlexaChannelPropertiesResponse) *AlexaChannelPropertiesResponse {
		return &v
	}).(AlexaChannelPropertiesResponsePtrOutput)
}

func (o AlexaChannelPropertiesResponseOutput) AlexaSkillId() pulumi.StringOutput {
	return o.ApplyT(func(v AlexaChannelPropertiesResponse) string { return v.AlexaSkillId }).(pulumi.StringOutput)
}

func (o AlexaChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v AlexaChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o AlexaChannelPropertiesResponseOutput) ServiceEndpointUri() pulumi.StringOutput {
	return o.ApplyT(func(v AlexaChannelPropertiesResponse) string { return v.ServiceEndpointUri }).(pulumi.StringOutput)
}

func (o AlexaChannelPropertiesResponseOutput) UrlFragment() pulumi.StringOutput {
	return o.ApplyT(func(v AlexaChannelPropertiesResponse) string { return v.UrlFragment }).(pulumi.StringOutput)
}

type AlexaChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (AlexaChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**AlexaChannelPropertiesResponse)(nil)).Elem()
}

func (o AlexaChannelPropertiesResponsePtrOutput) ToAlexaChannelPropertiesResponsePtrOutput() AlexaChannelPropertiesResponsePtrOutput {
	return o
}

func (o AlexaChannelPropertiesResponsePtrOutput) ToAlexaChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) AlexaChannelPropertiesResponsePtrOutput {
	return o
}

func (o AlexaChannelPropertiesResponsePtrOutput) Elem() AlexaChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *AlexaChannelPropertiesResponse) AlexaChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret AlexaChannelPropertiesResponse
		return ret
	}).(AlexaChannelPropertiesResponseOutput)
}

func (o AlexaChannelPropertiesResponsePtrOutput) AlexaSkillId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlexaChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AlexaSkillId
	}).(pulumi.StringPtrOutput)
}

func (o AlexaChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *AlexaChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o AlexaChannelPropertiesResponsePtrOutput) ServiceEndpointUri() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlexaChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceEndpointUri
	}).(pulumi.StringPtrOutput)
}

func (o AlexaChannelPropertiesResponsePtrOutput) UrlFragment() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *AlexaChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UrlFragment
	}).(pulumi.StringPtrOutput)
}

type AlexaChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *AlexaChannelPropertiesResponse `pulumi:"properties"`
}





type AlexaChannelResponseInput interface {
	pulumi.Input

	ToAlexaChannelResponseOutput() AlexaChannelResponseOutput
	ToAlexaChannelResponseOutputWithContext(context.Context) AlexaChannelResponseOutput
}

type AlexaChannelResponseArgs struct {
	ChannelName pulumi.StringInput                     `pulumi:"channelName"`
	Properties  AlexaChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (AlexaChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannelResponse)(nil)).Elem()
}

func (i AlexaChannelResponseArgs) ToAlexaChannelResponseOutput() AlexaChannelResponseOutput {
	return i.ToAlexaChannelResponseOutputWithContext(context.Background())
}

func (i AlexaChannelResponseArgs) ToAlexaChannelResponseOutputWithContext(ctx context.Context) AlexaChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(AlexaChannelResponseOutput)
}

type AlexaChannelResponseOutput struct{ *pulumi.OutputState }

func (AlexaChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*AlexaChannelResponse)(nil)).Elem()
}

func (o AlexaChannelResponseOutput) ToAlexaChannelResponseOutput() AlexaChannelResponseOutput {
	return o
}

func (o AlexaChannelResponseOutput) ToAlexaChannelResponseOutputWithContext(ctx context.Context) AlexaChannelResponseOutput {
	return o
}

func (o AlexaChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v AlexaChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o AlexaChannelResponseOutput) Properties() AlexaChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v AlexaChannelResponse) *AlexaChannelPropertiesResponse { return v.Properties }).(AlexaChannelPropertiesResponsePtrOutput)
}

type BotProperties struct {
	AppPasswordHint                   *string  `pulumi:"appPasswordHint"`
	CmekKeyVaultUrl                   *string  `pulumi:"cmekKeyVaultUrl"`
	Description                       *string  `pulumi:"description"`
	DeveloperAppInsightKey            *string  `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        *string  `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId *string  `pulumi:"developerAppInsightsApplicationId"`
	DisplayName                       string   `pulumi:"displayName"`
	Endpoint                          string   `pulumi:"endpoint"`
	IconUrl                           *string  `pulumi:"iconUrl"`
	IsCmekEnabled                     *bool    `pulumi:"isCmekEnabled"`
	IsIsolated                        *bool    `pulumi:"isIsolated"`
	LuisAppIds                        []string `pulumi:"luisAppIds"`
	LuisKey                           *string  `pulumi:"luisKey"`
	MsaAppId                          string   `pulumi:"msaAppId"`
	OpenWithHint                      *string  `pulumi:"openWithHint"`
	SchemaTransformationVersion       *string  `pulumi:"schemaTransformationVersion"`
}





type BotPropertiesInput interface {
	pulumi.Input

	ToBotPropertiesOutput() BotPropertiesOutput
	ToBotPropertiesOutputWithContext(context.Context) BotPropertiesOutput
}

type BotPropertiesArgs struct {
	AppPasswordHint                   pulumi.StringPtrInput   `pulumi:"appPasswordHint"`
	CmekKeyVaultUrl                   pulumi.StringPtrInput   `pulumi:"cmekKeyVaultUrl"`
	Description                       pulumi.StringPtrInput   `pulumi:"description"`
	DeveloperAppInsightKey            pulumi.StringPtrInput   `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        pulumi.StringPtrInput   `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput   `pulumi:"developerAppInsightsApplicationId"`
	DisplayName                       pulumi.StringInput      `pulumi:"displayName"`
	Endpoint                          pulumi.StringInput      `pulumi:"endpoint"`
	IconUrl                           pulumi.StringPtrInput   `pulumi:"iconUrl"`
	IsCmekEnabled                     pulumi.BoolPtrInput     `pulumi:"isCmekEnabled"`
	IsIsolated                        pulumi.BoolPtrInput     `pulumi:"isIsolated"`
	LuisAppIds                        pulumi.StringArrayInput `pulumi:"luisAppIds"`
	LuisKey                           pulumi.StringPtrInput   `pulumi:"luisKey"`
	MsaAppId                          pulumi.StringInput      `pulumi:"msaAppId"`
	OpenWithHint                      pulumi.StringPtrInput   `pulumi:"openWithHint"`
	SchemaTransformationVersion       pulumi.StringPtrInput   `pulumi:"schemaTransformationVersion"`
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

func (o BotPropertiesOutput) IsIsolated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotProperties) *bool { return v.IsIsolated }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotProperties) []string { return v.LuisAppIds }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.LuisKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) MsaAppId() pulumi.StringOutput {
	return o.ApplyT(func(v BotProperties) string { return v.MsaAppId }).(pulumi.StringOutput)
}

func (o BotPropertiesOutput) OpenWithHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.OpenWithHint }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.SchemaTransformationVersion }).(pulumi.StringPtrOutput)
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

func (o BotPropertiesPtrOutput) IsIsolated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsIsolated
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

func (o BotPropertiesPtrOutput) MsaAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return &v.MsaAppId
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

func (o BotPropertiesPtrOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotProperties) *string {
		if v == nil {
			return nil
		}
		return v.SchemaTransformationVersion
	}).(pulumi.StringPtrOutput)
}

type BotPropertiesResponse struct {
	AppPasswordHint                   *string  `pulumi:"appPasswordHint"`
	CmekKeyVaultUrl                   *string  `pulumi:"cmekKeyVaultUrl"`
	ConfiguredChannels                []string `pulumi:"configuredChannels"`
	Description                       *string  `pulumi:"description"`
	DeveloperAppInsightKey            *string  `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        *string  `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId *string  `pulumi:"developerAppInsightsApplicationId"`
	DisplayName                       string   `pulumi:"displayName"`
	EnabledChannels                   []string `pulumi:"enabledChannels"`
	Endpoint                          string   `pulumi:"endpoint"`
	EndpointVersion                   string   `pulumi:"endpointVersion"`
	IconUrl                           *string  `pulumi:"iconUrl"`
	IsCmekEnabled                     *bool    `pulumi:"isCmekEnabled"`
	IsIsolated                        *bool    `pulumi:"isIsolated"`
	LuisAppIds                        []string `pulumi:"luisAppIds"`
	LuisKey                           *string  `pulumi:"luisKey"`
	MsaAppId                          string   `pulumi:"msaAppId"`
	OpenWithHint                      *string  `pulumi:"openWithHint"`
	SchemaTransformationVersion       *string  `pulumi:"schemaTransformationVersion"`
}





type BotPropertiesResponseInput interface {
	pulumi.Input

	ToBotPropertiesResponseOutput() BotPropertiesResponseOutput
	ToBotPropertiesResponseOutputWithContext(context.Context) BotPropertiesResponseOutput
}

type BotPropertiesResponseArgs struct {
	AppPasswordHint                   pulumi.StringPtrInput   `pulumi:"appPasswordHint"`
	CmekKeyVaultUrl                   pulumi.StringPtrInput   `pulumi:"cmekKeyVaultUrl"`
	ConfiguredChannels                pulumi.StringArrayInput `pulumi:"configuredChannels"`
	Description                       pulumi.StringPtrInput   `pulumi:"description"`
	DeveloperAppInsightKey            pulumi.StringPtrInput   `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        pulumi.StringPtrInput   `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput   `pulumi:"developerAppInsightsApplicationId"`
	DisplayName                       pulumi.StringInput      `pulumi:"displayName"`
	EnabledChannels                   pulumi.StringArrayInput `pulumi:"enabledChannels"`
	Endpoint                          pulumi.StringInput      `pulumi:"endpoint"`
	EndpointVersion                   pulumi.StringInput      `pulumi:"endpointVersion"`
	IconUrl                           pulumi.StringPtrInput   `pulumi:"iconUrl"`
	IsCmekEnabled                     pulumi.BoolPtrInput     `pulumi:"isCmekEnabled"`
	IsIsolated                        pulumi.BoolPtrInput     `pulumi:"isIsolated"`
	LuisAppIds                        pulumi.StringArrayInput `pulumi:"luisAppIds"`
	LuisKey                           pulumi.StringPtrInput   `pulumi:"luisKey"`
	MsaAppId                          pulumi.StringInput      `pulumi:"msaAppId"`
	OpenWithHint                      pulumi.StringPtrInput   `pulumi:"openWithHint"`
	SchemaTransformationVersion       pulumi.StringPtrInput   `pulumi:"schemaTransformationVersion"`
}

func (BotPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*BotPropertiesResponse)(nil)).Elem()
}

func (i BotPropertiesResponseArgs) ToBotPropertiesResponseOutput() BotPropertiesResponseOutput {
	return i.ToBotPropertiesResponseOutputWithContext(context.Background())
}

func (i BotPropertiesResponseArgs) ToBotPropertiesResponseOutputWithContext(ctx context.Context) BotPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotPropertiesResponseOutput)
}

func (i BotPropertiesResponseArgs) ToBotPropertiesResponsePtrOutput() BotPropertiesResponsePtrOutput {
	return i.ToBotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i BotPropertiesResponseArgs) ToBotPropertiesResponsePtrOutputWithContext(ctx context.Context) BotPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotPropertiesResponseOutput).ToBotPropertiesResponsePtrOutputWithContext(ctx)
}









type BotPropertiesResponsePtrInput interface {
	pulumi.Input

	ToBotPropertiesResponsePtrOutput() BotPropertiesResponsePtrOutput
	ToBotPropertiesResponsePtrOutputWithContext(context.Context) BotPropertiesResponsePtrOutput
}

type botPropertiesResponsePtrType BotPropertiesResponseArgs

func BotPropertiesResponsePtr(v *BotPropertiesResponseArgs) BotPropertiesResponsePtrInput {
	return (*botPropertiesResponsePtrType)(v)
}

func (*botPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**BotPropertiesResponse)(nil)).Elem()
}

func (i *botPropertiesResponsePtrType) ToBotPropertiesResponsePtrOutput() BotPropertiesResponsePtrOutput {
	return i.ToBotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *botPropertiesResponsePtrType) ToBotPropertiesResponsePtrOutputWithContext(ctx context.Context) BotPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(BotPropertiesResponsePtrOutput)
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

func (o BotPropertiesResponseOutput) ToBotPropertiesResponsePtrOutput() BotPropertiesResponsePtrOutput {
	return o.ToBotPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o BotPropertiesResponseOutput) ToBotPropertiesResponsePtrOutputWithContext(ctx context.Context) BotPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v BotPropertiesResponse) *BotPropertiesResponse {
		return &v
	}).(BotPropertiesResponsePtrOutput)
}

func (o BotPropertiesResponseOutput) AppPasswordHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.AppPasswordHint }).(pulumi.StringPtrOutput)
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

func (o BotPropertiesResponseOutput) IsIsolated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *bool { return v.IsIsolated }).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponseOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotPropertiesResponse) []string { return v.LuisAppIds }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponseOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.LuisKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) MsaAppId() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.MsaAppId }).(pulumi.StringOutput)
}

func (o BotPropertiesResponseOutput) OpenWithHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.OpenWithHint }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.SchemaTransformationVersion }).(pulumi.StringPtrOutput)
}

type BotPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (BotPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**BotPropertiesResponse)(nil)).Elem()
}

func (o BotPropertiesResponsePtrOutput) ToBotPropertiesResponsePtrOutput() BotPropertiesResponsePtrOutput {
	return o
}

func (o BotPropertiesResponsePtrOutput) ToBotPropertiesResponsePtrOutputWithContext(ctx context.Context) BotPropertiesResponsePtrOutput {
	return o
}

func (o BotPropertiesResponsePtrOutput) Elem() BotPropertiesResponseOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) BotPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret BotPropertiesResponse
		return ret
	}).(BotPropertiesResponseOutput)
}

func (o BotPropertiesResponsePtrOutput) AppPasswordHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppPasswordHint
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) CmekKeyVaultUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CmekKeyVaultUrl
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) ConfiguredChannels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.ConfiguredChannels
	}).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponsePtrOutput) Description() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Description
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) DeveloperAppInsightKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeveloperAppInsightKey
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) DeveloperAppInsightsApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeveloperAppInsightsApiKey
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) DeveloperAppInsightsApplicationId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.DeveloperAppInsightsApplicationId
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) EnabledChannels() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.EnabledChannels
	}).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponsePtrOutput) Endpoint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Endpoint
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) EndpointVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EndpointVersion
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.IconUrl
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) IsCmekEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsCmekEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) IsIsolated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsIsolated
	}).(pulumi.BoolPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) []string {
		if v == nil {
			return nil
		}
		return v.LuisAppIds
	}).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponsePtrOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LuisKey
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) MsaAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.MsaAppId
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) OpenWithHint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.OpenWithHint
	}).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponsePtrOutput) SchemaTransformationVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *BotPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SchemaTransformationVersion
	}).(pulumi.StringPtrOutput)
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





type ConnectionSettingParameterResponseInput interface {
	pulumi.Input

	ToConnectionSettingParameterResponseOutput() ConnectionSettingParameterResponseOutput
	ToConnectionSettingParameterResponseOutputWithContext(context.Context) ConnectionSettingParameterResponseOutput
}

type ConnectionSettingParameterResponseArgs struct {
	Key   pulumi.StringPtrInput `pulumi:"key"`
	Value pulumi.StringPtrInput `pulumi:"value"`
}

func (ConnectionSettingParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingParameterResponse)(nil)).Elem()
}

func (i ConnectionSettingParameterResponseArgs) ToConnectionSettingParameterResponseOutput() ConnectionSettingParameterResponseOutput {
	return i.ToConnectionSettingParameterResponseOutputWithContext(context.Background())
}

func (i ConnectionSettingParameterResponseArgs) ToConnectionSettingParameterResponseOutputWithContext(ctx context.Context) ConnectionSettingParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingParameterResponseOutput)
}





type ConnectionSettingParameterResponseArrayInput interface {
	pulumi.Input

	ToConnectionSettingParameterResponseArrayOutput() ConnectionSettingParameterResponseArrayOutput
	ToConnectionSettingParameterResponseArrayOutputWithContext(context.Context) ConnectionSettingParameterResponseArrayOutput
}

type ConnectionSettingParameterResponseArray []ConnectionSettingParameterResponseInput

func (ConnectionSettingParameterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ConnectionSettingParameterResponse)(nil)).Elem()
}

func (i ConnectionSettingParameterResponseArray) ToConnectionSettingParameterResponseArrayOutput() ConnectionSettingParameterResponseArrayOutput {
	return i.ToConnectionSettingParameterResponseArrayOutputWithContext(context.Background())
}

func (i ConnectionSettingParameterResponseArray) ToConnectionSettingParameterResponseArrayOutputWithContext(ctx context.Context) ConnectionSettingParameterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingParameterResponseArrayOutput)
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
	Parameters                 ConnectionSettingParameterArrayInput `pulumi:"parameters"`
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

func (o ConnectionSettingPropertiesOutput) Parameters() ConnectionSettingParameterArrayOutput {
	return o.ApplyT(func(v ConnectionSettingProperties) []ConnectionSettingParameter { return v.Parameters }).(ConnectionSettingParameterArrayOutput)
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
	Scopes                     *string                              `pulumi:"scopes"`
	ServiceProviderDisplayName *string                              `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          *string                              `pulumi:"serviceProviderId"`
	SettingId                  string                               `pulumi:"settingId"`
}





type ConnectionSettingPropertiesResponseInput interface {
	pulumi.Input

	ToConnectionSettingPropertiesResponseOutput() ConnectionSettingPropertiesResponseOutput
	ToConnectionSettingPropertiesResponseOutputWithContext(context.Context) ConnectionSettingPropertiesResponseOutput
}

type ConnectionSettingPropertiesResponseArgs struct {
	ClientId                   pulumi.StringPtrInput                        `pulumi:"clientId"`
	ClientSecret               pulumi.StringPtrInput                        `pulumi:"clientSecret"`
	Parameters                 ConnectionSettingParameterResponseArrayInput `pulumi:"parameters"`
	Scopes                     pulumi.StringPtrInput                        `pulumi:"scopes"`
	ServiceProviderDisplayName pulumi.StringPtrInput                        `pulumi:"serviceProviderDisplayName"`
	ServiceProviderId          pulumi.StringPtrInput                        `pulumi:"serviceProviderId"`
	SettingId                  pulumi.StringInput                           `pulumi:"settingId"`
}

func (ConnectionSettingPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ConnectionSettingPropertiesResponse)(nil)).Elem()
}

func (i ConnectionSettingPropertiesResponseArgs) ToConnectionSettingPropertiesResponseOutput() ConnectionSettingPropertiesResponseOutput {
	return i.ToConnectionSettingPropertiesResponseOutputWithContext(context.Background())
}

func (i ConnectionSettingPropertiesResponseArgs) ToConnectionSettingPropertiesResponseOutputWithContext(ctx context.Context) ConnectionSettingPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingPropertiesResponseOutput)
}

func (i ConnectionSettingPropertiesResponseArgs) ToConnectionSettingPropertiesResponsePtrOutput() ConnectionSettingPropertiesResponsePtrOutput {
	return i.ToConnectionSettingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ConnectionSettingPropertiesResponseArgs) ToConnectionSettingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingPropertiesResponseOutput).ToConnectionSettingPropertiesResponsePtrOutputWithContext(ctx)
}









type ConnectionSettingPropertiesResponsePtrInput interface {
	pulumi.Input

	ToConnectionSettingPropertiesResponsePtrOutput() ConnectionSettingPropertiesResponsePtrOutput
	ToConnectionSettingPropertiesResponsePtrOutputWithContext(context.Context) ConnectionSettingPropertiesResponsePtrOutput
}

type connectionSettingPropertiesResponsePtrType ConnectionSettingPropertiesResponseArgs

func ConnectionSettingPropertiesResponsePtr(v *ConnectionSettingPropertiesResponseArgs) ConnectionSettingPropertiesResponsePtrInput {
	return (*connectionSettingPropertiesResponsePtrType)(v)
}

func (*connectionSettingPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionSettingPropertiesResponse)(nil)).Elem()
}

func (i *connectionSettingPropertiesResponsePtrType) ToConnectionSettingPropertiesResponsePtrOutput() ConnectionSettingPropertiesResponsePtrOutput {
	return i.ToConnectionSettingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *connectionSettingPropertiesResponsePtrType) ToConnectionSettingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ConnectionSettingPropertiesResponsePtrOutput)
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

func (o ConnectionSettingPropertiesResponseOutput) ToConnectionSettingPropertiesResponsePtrOutput() ConnectionSettingPropertiesResponsePtrOutput {
	return o.ToConnectionSettingPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ConnectionSettingPropertiesResponseOutput) ToConnectionSettingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ConnectionSettingPropertiesResponse) *ConnectionSettingPropertiesResponse {
		return &v
	}).(ConnectionSettingPropertiesResponsePtrOutput)
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

type ConnectionSettingPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ConnectionSettingPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ConnectionSettingPropertiesResponse)(nil)).Elem()
}

func (o ConnectionSettingPropertiesResponsePtrOutput) ToConnectionSettingPropertiesResponsePtrOutput() ConnectionSettingPropertiesResponsePtrOutput {
	return o
}

func (o ConnectionSettingPropertiesResponsePtrOutput) ToConnectionSettingPropertiesResponsePtrOutputWithContext(ctx context.Context) ConnectionSettingPropertiesResponsePtrOutput {
	return o
}

func (o ConnectionSettingPropertiesResponsePtrOutput) Elem() ConnectionSettingPropertiesResponseOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) ConnectionSettingPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ConnectionSettingPropertiesResponse
		return ret
	}).(ConnectionSettingPropertiesResponseOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) Parameters() ConnectionSettingParameterResponseArrayOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) []ConnectionSettingParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ConnectionSettingParameterResponseArrayOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) Scopes() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Scopes
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) ServiceProviderDisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderDisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) ServiceProviderId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ServiceProviderId
	}).(pulumi.StringPtrOutput)
}

func (o ConnectionSettingPropertiesResponsePtrOutput) SettingId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ConnectionSettingPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.SettingId
	}).(pulumi.StringPtrOutput)
}

type DirectLineChannel struct {
	ChannelName string                       `pulumi:"channelName"`
	Properties  *DirectLineChannelProperties `pulumi:"properties"`
}





type DirectLineChannelInput interface {
	pulumi.Input

	ToDirectLineChannelOutput() DirectLineChannelOutput
	ToDirectLineChannelOutputWithContext(context.Context) DirectLineChannelOutput
}

type DirectLineChannelArgs struct {
	ChannelName pulumi.StringInput                  `pulumi:"channelName"`
	Properties  DirectLineChannelPropertiesPtrInput `pulumi:"properties"`
}

func (DirectLineChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannel)(nil)).Elem()
}

func (i DirectLineChannelArgs) ToDirectLineChannelOutput() DirectLineChannelOutput {
	return i.ToDirectLineChannelOutputWithContext(context.Background())
}

func (i DirectLineChannelArgs) ToDirectLineChannelOutputWithContext(ctx context.Context) DirectLineChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelOutput)
}

type DirectLineChannelOutput struct{ *pulumi.OutputState }

func (DirectLineChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannel)(nil)).Elem()
}

func (o DirectLineChannelOutput) ToDirectLineChannelOutput() DirectLineChannelOutput {
	return o
}

func (o DirectLineChannelOutput) ToDirectLineChannelOutputWithContext(ctx context.Context) DirectLineChannelOutput {
	return o
}

func (o DirectLineChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o DirectLineChannelOutput) Properties() DirectLineChannelPropertiesPtrOutput {
	return o.ApplyT(func(v DirectLineChannel) *DirectLineChannelProperties { return v.Properties }).(DirectLineChannelPropertiesPtrOutput)
}

type DirectLineChannelProperties struct {
	Sites []DirectLineSite `pulumi:"sites"`
}





type DirectLineChannelPropertiesInput interface {
	pulumi.Input

	ToDirectLineChannelPropertiesOutput() DirectLineChannelPropertiesOutput
	ToDirectLineChannelPropertiesOutputWithContext(context.Context) DirectLineChannelPropertiesOutput
}

type DirectLineChannelPropertiesArgs struct {
	Sites DirectLineSiteArrayInput `pulumi:"sites"`
}

func (DirectLineChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannelProperties)(nil)).Elem()
}

func (i DirectLineChannelPropertiesArgs) ToDirectLineChannelPropertiesOutput() DirectLineChannelPropertiesOutput {
	return i.ToDirectLineChannelPropertiesOutputWithContext(context.Background())
}

func (i DirectLineChannelPropertiesArgs) ToDirectLineChannelPropertiesOutputWithContext(ctx context.Context) DirectLineChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelPropertiesOutput)
}

func (i DirectLineChannelPropertiesArgs) ToDirectLineChannelPropertiesPtrOutput() DirectLineChannelPropertiesPtrOutput {
	return i.ToDirectLineChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i DirectLineChannelPropertiesArgs) ToDirectLineChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelPropertiesOutput).ToDirectLineChannelPropertiesPtrOutputWithContext(ctx)
}









type DirectLineChannelPropertiesPtrInput interface {
	pulumi.Input

	ToDirectLineChannelPropertiesPtrOutput() DirectLineChannelPropertiesPtrOutput
	ToDirectLineChannelPropertiesPtrOutputWithContext(context.Context) DirectLineChannelPropertiesPtrOutput
}

type directLineChannelPropertiesPtrType DirectLineChannelPropertiesArgs

func DirectLineChannelPropertiesPtr(v *DirectLineChannelPropertiesArgs) DirectLineChannelPropertiesPtrInput {
	return (*directLineChannelPropertiesPtrType)(v)
}

func (*directLineChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineChannelProperties)(nil)).Elem()
}

func (i *directLineChannelPropertiesPtrType) ToDirectLineChannelPropertiesPtrOutput() DirectLineChannelPropertiesPtrOutput {
	return i.ToDirectLineChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *directLineChannelPropertiesPtrType) ToDirectLineChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelPropertiesPtrOutput)
}

type DirectLineChannelPropertiesOutput struct{ *pulumi.OutputState }

func (DirectLineChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannelProperties)(nil)).Elem()
}

func (o DirectLineChannelPropertiesOutput) ToDirectLineChannelPropertiesOutput() DirectLineChannelPropertiesOutput {
	return o
}

func (o DirectLineChannelPropertiesOutput) ToDirectLineChannelPropertiesOutputWithContext(ctx context.Context) DirectLineChannelPropertiesOutput {
	return o
}

func (o DirectLineChannelPropertiesOutput) ToDirectLineChannelPropertiesPtrOutput() DirectLineChannelPropertiesPtrOutput {
	return o.ToDirectLineChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o DirectLineChannelPropertiesOutput) ToDirectLineChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectLineChannelProperties) *DirectLineChannelProperties {
		return &v
	}).(DirectLineChannelPropertiesPtrOutput)
}

func (o DirectLineChannelPropertiesOutput) Sites() DirectLineSiteArrayOutput {
	return o.ApplyT(func(v DirectLineChannelProperties) []DirectLineSite { return v.Sites }).(DirectLineSiteArrayOutput)
}

type DirectLineChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DirectLineChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineChannelProperties)(nil)).Elem()
}

func (o DirectLineChannelPropertiesPtrOutput) ToDirectLineChannelPropertiesPtrOutput() DirectLineChannelPropertiesPtrOutput {
	return o
}

func (o DirectLineChannelPropertiesPtrOutput) ToDirectLineChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesPtrOutput {
	return o
}

func (o DirectLineChannelPropertiesPtrOutput) Elem() DirectLineChannelPropertiesOutput {
	return o.ApplyT(func(v *DirectLineChannelProperties) DirectLineChannelProperties {
		if v != nil {
			return *v
		}
		var ret DirectLineChannelProperties
		return ret
	}).(DirectLineChannelPropertiesOutput)
}

func (o DirectLineChannelPropertiesPtrOutput) Sites() DirectLineSiteArrayOutput {
	return o.ApplyT(func(v *DirectLineChannelProperties) []DirectLineSite {
		if v == nil {
			return nil
		}
		return v.Sites
	}).(DirectLineSiteArrayOutput)
}

type DirectLineChannelPropertiesResponse struct {
	Sites []DirectLineSiteResponse `pulumi:"sites"`
}





type DirectLineChannelPropertiesResponseInput interface {
	pulumi.Input

	ToDirectLineChannelPropertiesResponseOutput() DirectLineChannelPropertiesResponseOutput
	ToDirectLineChannelPropertiesResponseOutputWithContext(context.Context) DirectLineChannelPropertiesResponseOutput
}

type DirectLineChannelPropertiesResponseArgs struct {
	Sites DirectLineSiteResponseArrayInput `pulumi:"sites"`
}

func (DirectLineChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannelPropertiesResponse)(nil)).Elem()
}

func (i DirectLineChannelPropertiesResponseArgs) ToDirectLineChannelPropertiesResponseOutput() DirectLineChannelPropertiesResponseOutput {
	return i.ToDirectLineChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i DirectLineChannelPropertiesResponseArgs) ToDirectLineChannelPropertiesResponseOutputWithContext(ctx context.Context) DirectLineChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelPropertiesResponseOutput)
}

func (i DirectLineChannelPropertiesResponseArgs) ToDirectLineChannelPropertiesResponsePtrOutput() DirectLineChannelPropertiesResponsePtrOutput {
	return i.ToDirectLineChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DirectLineChannelPropertiesResponseArgs) ToDirectLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelPropertiesResponseOutput).ToDirectLineChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type DirectLineChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDirectLineChannelPropertiesResponsePtrOutput() DirectLineChannelPropertiesResponsePtrOutput
	ToDirectLineChannelPropertiesResponsePtrOutputWithContext(context.Context) DirectLineChannelPropertiesResponsePtrOutput
}

type directLineChannelPropertiesResponsePtrType DirectLineChannelPropertiesResponseArgs

func DirectLineChannelPropertiesResponsePtr(v *DirectLineChannelPropertiesResponseArgs) DirectLineChannelPropertiesResponsePtrInput {
	return (*directLineChannelPropertiesResponsePtrType)(v)
}

func (*directLineChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineChannelPropertiesResponse)(nil)).Elem()
}

func (i *directLineChannelPropertiesResponsePtrType) ToDirectLineChannelPropertiesResponsePtrOutput() DirectLineChannelPropertiesResponsePtrOutput {
	return i.ToDirectLineChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *directLineChannelPropertiesResponsePtrType) ToDirectLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelPropertiesResponsePtrOutput)
}

type DirectLineChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DirectLineChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannelPropertiesResponse)(nil)).Elem()
}

func (o DirectLineChannelPropertiesResponseOutput) ToDirectLineChannelPropertiesResponseOutput() DirectLineChannelPropertiesResponseOutput {
	return o
}

func (o DirectLineChannelPropertiesResponseOutput) ToDirectLineChannelPropertiesResponseOutputWithContext(ctx context.Context) DirectLineChannelPropertiesResponseOutput {
	return o
}

func (o DirectLineChannelPropertiesResponseOutput) ToDirectLineChannelPropertiesResponsePtrOutput() DirectLineChannelPropertiesResponsePtrOutput {
	return o.ToDirectLineChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DirectLineChannelPropertiesResponseOutput) ToDirectLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectLineChannelPropertiesResponse) *DirectLineChannelPropertiesResponse {
		return &v
	}).(DirectLineChannelPropertiesResponsePtrOutput)
}

func (o DirectLineChannelPropertiesResponseOutput) Sites() DirectLineSiteResponseArrayOutput {
	return o.ApplyT(func(v DirectLineChannelPropertiesResponse) []DirectLineSiteResponse { return v.Sites }).(DirectLineSiteResponseArrayOutput)
}

type DirectLineChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DirectLineChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineChannelPropertiesResponse)(nil)).Elem()
}

func (o DirectLineChannelPropertiesResponsePtrOutput) ToDirectLineChannelPropertiesResponsePtrOutput() DirectLineChannelPropertiesResponsePtrOutput {
	return o
}

func (o DirectLineChannelPropertiesResponsePtrOutput) ToDirectLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineChannelPropertiesResponsePtrOutput {
	return o
}

func (o DirectLineChannelPropertiesResponsePtrOutput) Elem() DirectLineChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *DirectLineChannelPropertiesResponse) DirectLineChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DirectLineChannelPropertiesResponse
		return ret
	}).(DirectLineChannelPropertiesResponseOutput)
}

func (o DirectLineChannelPropertiesResponsePtrOutput) Sites() DirectLineSiteResponseArrayOutput {
	return o.ApplyT(func(v *DirectLineChannelPropertiesResponse) []DirectLineSiteResponse {
		if v == nil {
			return nil
		}
		return v.Sites
	}).(DirectLineSiteResponseArrayOutput)
}

type DirectLineChannelResponse struct {
	ChannelName string                               `pulumi:"channelName"`
	Properties  *DirectLineChannelPropertiesResponse `pulumi:"properties"`
}





type DirectLineChannelResponseInput interface {
	pulumi.Input

	ToDirectLineChannelResponseOutput() DirectLineChannelResponseOutput
	ToDirectLineChannelResponseOutputWithContext(context.Context) DirectLineChannelResponseOutput
}

type DirectLineChannelResponseArgs struct {
	ChannelName pulumi.StringInput                          `pulumi:"channelName"`
	Properties  DirectLineChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (DirectLineChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannelResponse)(nil)).Elem()
}

func (i DirectLineChannelResponseArgs) ToDirectLineChannelResponseOutput() DirectLineChannelResponseOutput {
	return i.ToDirectLineChannelResponseOutputWithContext(context.Background())
}

func (i DirectLineChannelResponseArgs) ToDirectLineChannelResponseOutputWithContext(ctx context.Context) DirectLineChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineChannelResponseOutput)
}

type DirectLineChannelResponseOutput struct{ *pulumi.OutputState }

func (DirectLineChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineChannelResponse)(nil)).Elem()
}

func (o DirectLineChannelResponseOutput) ToDirectLineChannelResponseOutput() DirectLineChannelResponseOutput {
	return o
}

func (o DirectLineChannelResponseOutput) ToDirectLineChannelResponseOutputWithContext(ctx context.Context) DirectLineChannelResponseOutput {
	return o
}

func (o DirectLineChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o DirectLineChannelResponseOutput) Properties() DirectLineChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v DirectLineChannelResponse) *DirectLineChannelPropertiesResponse { return v.Properties }).(DirectLineChannelPropertiesResponsePtrOutput)
}

type DirectLineSite struct {
	IsEnabled           bool     `pulumi:"isEnabled"`
	IsSecureSiteEnabled *bool    `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled         bool     `pulumi:"isV1Enabled"`
	IsV3Enabled         bool     `pulumi:"isV3Enabled"`
	SiteName            string   `pulumi:"siteName"`
	TrustedOrigins      []string `pulumi:"trustedOrigins"`
}





type DirectLineSiteInput interface {
	pulumi.Input

	ToDirectLineSiteOutput() DirectLineSiteOutput
	ToDirectLineSiteOutputWithContext(context.Context) DirectLineSiteOutput
}

type DirectLineSiteArgs struct {
	IsEnabled           pulumi.BoolInput        `pulumi:"isEnabled"`
	IsSecureSiteEnabled pulumi.BoolPtrInput     `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled         pulumi.BoolInput        `pulumi:"isV1Enabled"`
	IsV3Enabled         pulumi.BoolInput        `pulumi:"isV3Enabled"`
	SiteName            pulumi.StringInput      `pulumi:"siteName"`
	TrustedOrigins      pulumi.StringArrayInput `pulumi:"trustedOrigins"`
}

func (DirectLineSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSite)(nil)).Elem()
}

func (i DirectLineSiteArgs) ToDirectLineSiteOutput() DirectLineSiteOutput {
	return i.ToDirectLineSiteOutputWithContext(context.Background())
}

func (i DirectLineSiteArgs) ToDirectLineSiteOutputWithContext(ctx context.Context) DirectLineSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSiteOutput)
}





type DirectLineSiteArrayInput interface {
	pulumi.Input

	ToDirectLineSiteArrayOutput() DirectLineSiteArrayOutput
	ToDirectLineSiteArrayOutputWithContext(context.Context) DirectLineSiteArrayOutput
}

type DirectLineSiteArray []DirectLineSiteInput

func (DirectLineSiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectLineSite)(nil)).Elem()
}

func (i DirectLineSiteArray) ToDirectLineSiteArrayOutput() DirectLineSiteArrayOutput {
	return i.ToDirectLineSiteArrayOutputWithContext(context.Background())
}

func (i DirectLineSiteArray) ToDirectLineSiteArrayOutputWithContext(ctx context.Context) DirectLineSiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSiteArrayOutput)
}

type DirectLineSiteOutput struct{ *pulumi.OutputState }

func (DirectLineSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSite)(nil)).Elem()
}

func (o DirectLineSiteOutput) ToDirectLineSiteOutput() DirectLineSiteOutput {
	return o
}

func (o DirectLineSiteOutput) ToDirectLineSiteOutputWithContext(ctx context.Context) DirectLineSiteOutput {
	return o
}

func (o DirectLineSiteOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DirectLineSite) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o DirectLineSiteOutput) IsSecureSiteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectLineSite) *bool { return v.IsSecureSiteEnabled }).(pulumi.BoolPtrOutput)
}

func (o DirectLineSiteOutput) IsV1Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DirectLineSite) bool { return v.IsV1Enabled }).(pulumi.BoolOutput)
}

func (o DirectLineSiteOutput) IsV3Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DirectLineSite) bool { return v.IsV3Enabled }).(pulumi.BoolOutput)
}

func (o DirectLineSiteOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSite) string { return v.SiteName }).(pulumi.StringOutput)
}

func (o DirectLineSiteOutput) TrustedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DirectLineSite) []string { return v.TrustedOrigins }).(pulumi.StringArrayOutput)
}

type DirectLineSiteArrayOutput struct{ *pulumi.OutputState }

func (DirectLineSiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectLineSite)(nil)).Elem()
}

func (o DirectLineSiteArrayOutput) ToDirectLineSiteArrayOutput() DirectLineSiteArrayOutput {
	return o
}

func (o DirectLineSiteArrayOutput) ToDirectLineSiteArrayOutputWithContext(ctx context.Context) DirectLineSiteArrayOutput {
	return o
}

func (o DirectLineSiteArrayOutput) Index(i pulumi.IntInput) DirectLineSiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DirectLineSite {
		return vs[0].([]DirectLineSite)[vs[1].(int)]
	}).(DirectLineSiteOutput)
}

type DirectLineSiteResponse struct {
	IsEnabled           bool     `pulumi:"isEnabled"`
	IsSecureSiteEnabled *bool    `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled         bool     `pulumi:"isV1Enabled"`
	IsV3Enabled         bool     `pulumi:"isV3Enabled"`
	Key                 string   `pulumi:"key"`
	Key2                string   `pulumi:"key2"`
	SiteId              string   `pulumi:"siteId"`
	SiteName            string   `pulumi:"siteName"`
	TrustedOrigins      []string `pulumi:"trustedOrigins"`
}





type DirectLineSiteResponseInput interface {
	pulumi.Input

	ToDirectLineSiteResponseOutput() DirectLineSiteResponseOutput
	ToDirectLineSiteResponseOutputWithContext(context.Context) DirectLineSiteResponseOutput
}

type DirectLineSiteResponseArgs struct {
	IsEnabled           pulumi.BoolInput        `pulumi:"isEnabled"`
	IsSecureSiteEnabled pulumi.BoolPtrInput     `pulumi:"isSecureSiteEnabled"`
	IsV1Enabled         pulumi.BoolInput        `pulumi:"isV1Enabled"`
	IsV3Enabled         pulumi.BoolInput        `pulumi:"isV3Enabled"`
	Key                 pulumi.StringInput      `pulumi:"key"`
	Key2                pulumi.StringInput      `pulumi:"key2"`
	SiteId              pulumi.StringInput      `pulumi:"siteId"`
	SiteName            pulumi.StringInput      `pulumi:"siteName"`
	TrustedOrigins      pulumi.StringArrayInput `pulumi:"trustedOrigins"`
}

func (DirectLineSiteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSiteResponse)(nil)).Elem()
}

func (i DirectLineSiteResponseArgs) ToDirectLineSiteResponseOutput() DirectLineSiteResponseOutput {
	return i.ToDirectLineSiteResponseOutputWithContext(context.Background())
}

func (i DirectLineSiteResponseArgs) ToDirectLineSiteResponseOutputWithContext(ctx context.Context) DirectLineSiteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSiteResponseOutput)
}





type DirectLineSiteResponseArrayInput interface {
	pulumi.Input

	ToDirectLineSiteResponseArrayOutput() DirectLineSiteResponseArrayOutput
	ToDirectLineSiteResponseArrayOutputWithContext(context.Context) DirectLineSiteResponseArrayOutput
}

type DirectLineSiteResponseArray []DirectLineSiteResponseInput

func (DirectLineSiteResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectLineSiteResponse)(nil)).Elem()
}

func (i DirectLineSiteResponseArray) ToDirectLineSiteResponseArrayOutput() DirectLineSiteResponseArrayOutput {
	return i.ToDirectLineSiteResponseArrayOutputWithContext(context.Background())
}

func (i DirectLineSiteResponseArray) ToDirectLineSiteResponseArrayOutputWithContext(ctx context.Context) DirectLineSiteResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSiteResponseArrayOutput)
}

type DirectLineSiteResponseOutput struct{ *pulumi.OutputState }

func (DirectLineSiteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSiteResponse)(nil)).Elem()
}

func (o DirectLineSiteResponseOutput) ToDirectLineSiteResponseOutput() DirectLineSiteResponseOutput {
	return o
}

func (o DirectLineSiteResponseOutput) ToDirectLineSiteResponseOutputWithContext(ctx context.Context) DirectLineSiteResponseOutput {
	return o
}

func (o DirectLineSiteResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o DirectLineSiteResponseOutput) IsSecureSiteEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) *bool { return v.IsSecureSiteEnabled }).(pulumi.BoolPtrOutput)
}

func (o DirectLineSiteResponseOutput) IsV1Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) bool { return v.IsV1Enabled }).(pulumi.BoolOutput)
}

func (o DirectLineSiteResponseOutput) IsV3Enabled() pulumi.BoolOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) bool { return v.IsV3Enabled }).(pulumi.BoolOutput)
}

func (o DirectLineSiteResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o DirectLineSiteResponseOutput) Key2() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) string { return v.Key2 }).(pulumi.StringOutput)
}

func (o DirectLineSiteResponseOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) string { return v.SiteId }).(pulumi.StringOutput)
}

func (o DirectLineSiteResponseOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) string { return v.SiteName }).(pulumi.StringOutput)
}

func (o DirectLineSiteResponseOutput) TrustedOrigins() pulumi.StringArrayOutput {
	return o.ApplyT(func(v DirectLineSiteResponse) []string { return v.TrustedOrigins }).(pulumi.StringArrayOutput)
}

type DirectLineSiteResponseArrayOutput struct{ *pulumi.OutputState }

func (DirectLineSiteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]DirectLineSiteResponse)(nil)).Elem()
}

func (o DirectLineSiteResponseArrayOutput) ToDirectLineSiteResponseArrayOutput() DirectLineSiteResponseArrayOutput {
	return o
}

func (o DirectLineSiteResponseArrayOutput) ToDirectLineSiteResponseArrayOutputWithContext(ctx context.Context) DirectLineSiteResponseArrayOutput {
	return o
}

func (o DirectLineSiteResponseArrayOutput) Index(i pulumi.IntInput) DirectLineSiteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) DirectLineSiteResponse {
		return vs[0].([]DirectLineSiteResponse)[vs[1].(int)]
	}).(DirectLineSiteResponseOutput)
}

type DirectLineSpeechChannel struct {
	ChannelName string                             `pulumi:"channelName"`
	Properties  *DirectLineSpeechChannelProperties `pulumi:"properties"`
}





type DirectLineSpeechChannelInput interface {
	pulumi.Input

	ToDirectLineSpeechChannelOutput() DirectLineSpeechChannelOutput
	ToDirectLineSpeechChannelOutputWithContext(context.Context) DirectLineSpeechChannelOutput
}

type DirectLineSpeechChannelArgs struct {
	ChannelName pulumi.StringInput                        `pulumi:"channelName"`
	Properties  DirectLineSpeechChannelPropertiesPtrInput `pulumi:"properties"`
}

func (DirectLineSpeechChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannel)(nil)).Elem()
}

func (i DirectLineSpeechChannelArgs) ToDirectLineSpeechChannelOutput() DirectLineSpeechChannelOutput {
	return i.ToDirectLineSpeechChannelOutputWithContext(context.Background())
}

func (i DirectLineSpeechChannelArgs) ToDirectLineSpeechChannelOutputWithContext(ctx context.Context) DirectLineSpeechChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelOutput)
}

type DirectLineSpeechChannelOutput struct{ *pulumi.OutputState }

func (DirectLineSpeechChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannel)(nil)).Elem()
}

func (o DirectLineSpeechChannelOutput) ToDirectLineSpeechChannelOutput() DirectLineSpeechChannelOutput {
	return o
}

func (o DirectLineSpeechChannelOutput) ToDirectLineSpeechChannelOutputWithContext(ctx context.Context) DirectLineSpeechChannelOutput {
	return o
}

func (o DirectLineSpeechChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSpeechChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o DirectLineSpeechChannelOutput) Properties() DirectLineSpeechChannelPropertiesPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannel) *DirectLineSpeechChannelProperties { return v.Properties }).(DirectLineSpeechChannelPropertiesPtrOutput)
}

type DirectLineSpeechChannelProperties struct {
	CognitiveServiceRegion          string  `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey string  `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             *string `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         *string `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    *bool   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       *bool   `pulumi:"isEnabled"`
}





type DirectLineSpeechChannelPropertiesInput interface {
	pulumi.Input

	ToDirectLineSpeechChannelPropertiesOutput() DirectLineSpeechChannelPropertiesOutput
	ToDirectLineSpeechChannelPropertiesOutputWithContext(context.Context) DirectLineSpeechChannelPropertiesOutput
}

type DirectLineSpeechChannelPropertiesArgs struct {
	CognitiveServiceRegion          pulumi.StringInput    `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey pulumi.StringInput    `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             pulumi.StringPtrInput `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         pulumi.StringPtrInput `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    pulumi.BoolPtrInput   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       pulumi.BoolPtrInput   `pulumi:"isEnabled"`
}

func (DirectLineSpeechChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannelProperties)(nil)).Elem()
}

func (i DirectLineSpeechChannelPropertiesArgs) ToDirectLineSpeechChannelPropertiesOutput() DirectLineSpeechChannelPropertiesOutput {
	return i.ToDirectLineSpeechChannelPropertiesOutputWithContext(context.Background())
}

func (i DirectLineSpeechChannelPropertiesArgs) ToDirectLineSpeechChannelPropertiesOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelPropertiesOutput)
}

func (i DirectLineSpeechChannelPropertiesArgs) ToDirectLineSpeechChannelPropertiesPtrOutput() DirectLineSpeechChannelPropertiesPtrOutput {
	return i.ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i DirectLineSpeechChannelPropertiesArgs) ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelPropertiesOutput).ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(ctx)
}









type DirectLineSpeechChannelPropertiesPtrInput interface {
	pulumi.Input

	ToDirectLineSpeechChannelPropertiesPtrOutput() DirectLineSpeechChannelPropertiesPtrOutput
	ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(context.Context) DirectLineSpeechChannelPropertiesPtrOutput
}

type directLineSpeechChannelPropertiesPtrType DirectLineSpeechChannelPropertiesArgs

func DirectLineSpeechChannelPropertiesPtr(v *DirectLineSpeechChannelPropertiesArgs) DirectLineSpeechChannelPropertiesPtrInput {
	return (*directLineSpeechChannelPropertiesPtrType)(v)
}

func (*directLineSpeechChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineSpeechChannelProperties)(nil)).Elem()
}

func (i *directLineSpeechChannelPropertiesPtrType) ToDirectLineSpeechChannelPropertiesPtrOutput() DirectLineSpeechChannelPropertiesPtrOutput {
	return i.ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *directLineSpeechChannelPropertiesPtrType) ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelPropertiesPtrOutput)
}

type DirectLineSpeechChannelPropertiesOutput struct{ *pulumi.OutputState }

func (DirectLineSpeechChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannelProperties)(nil)).Elem()
}

func (o DirectLineSpeechChannelPropertiesOutput) ToDirectLineSpeechChannelPropertiesOutput() DirectLineSpeechChannelPropertiesOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesOutput) ToDirectLineSpeechChannelPropertiesOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesOutput) ToDirectLineSpeechChannelPropertiesPtrOutput() DirectLineSpeechChannelPropertiesPtrOutput {
	return o.ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o DirectLineSpeechChannelPropertiesOutput) ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectLineSpeechChannelProperties) *DirectLineSpeechChannelProperties {
		return &v
	}).(DirectLineSpeechChannelPropertiesPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesOutput) CognitiveServiceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelProperties) string { return v.CognitiveServiceRegion }).(pulumi.StringOutput)
}

func (o DirectLineSpeechChannelPropertiesOutput) CognitiveServiceSubscriptionKey() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelProperties) string { return v.CognitiveServiceSubscriptionKey }).(pulumi.StringOutput)
}

func (o DirectLineSpeechChannelPropertiesOutput) CustomSpeechModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelProperties) *string { return v.CustomSpeechModelId }).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesOutput) CustomVoiceDeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelProperties) *string { return v.CustomVoiceDeploymentId }).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesOutput) IsDefaultBotForCogSvcAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelProperties) *bool { return v.IsDefaultBotForCogSvcAccount }).(pulumi.BoolPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelProperties) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type DirectLineSpeechChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (DirectLineSpeechChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineSpeechChannelProperties)(nil)).Elem()
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) ToDirectLineSpeechChannelPropertiesPtrOutput() DirectLineSpeechChannelPropertiesPtrOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) ToDirectLineSpeechChannelPropertiesPtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesPtrOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) Elem() DirectLineSpeechChannelPropertiesOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) DirectLineSpeechChannelProperties {
		if v != nil {
			return *v
		}
		var ret DirectLineSpeechChannelProperties
		return ret
	}).(DirectLineSpeechChannelPropertiesOutput)
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) CognitiveServiceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.CognitiveServiceRegion
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) CognitiveServiceSubscriptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.CognitiveServiceSubscriptionKey
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) CustomSpeechModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomSpeechModelId
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) CustomVoiceDeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.CustomVoiceDeploymentId
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) IsDefaultBotForCogSvcAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsDefaultBotForCogSvcAccount
	}).(pulumi.BoolPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type DirectLineSpeechChannelPropertiesResponse struct {
	CognitiveServiceRegion          string  `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey string  `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             *string `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         *string `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    *bool   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       *bool   `pulumi:"isEnabled"`
}





type DirectLineSpeechChannelPropertiesResponseInput interface {
	pulumi.Input

	ToDirectLineSpeechChannelPropertiesResponseOutput() DirectLineSpeechChannelPropertiesResponseOutput
	ToDirectLineSpeechChannelPropertiesResponseOutputWithContext(context.Context) DirectLineSpeechChannelPropertiesResponseOutput
}

type DirectLineSpeechChannelPropertiesResponseArgs struct {
	CognitiveServiceRegion          pulumi.StringInput    `pulumi:"cognitiveServiceRegion"`
	CognitiveServiceSubscriptionKey pulumi.StringInput    `pulumi:"cognitiveServiceSubscriptionKey"`
	CustomSpeechModelId             pulumi.StringPtrInput `pulumi:"customSpeechModelId"`
	CustomVoiceDeploymentId         pulumi.StringPtrInput `pulumi:"customVoiceDeploymentId"`
	IsDefaultBotForCogSvcAccount    pulumi.BoolPtrInput   `pulumi:"isDefaultBotForCogSvcAccount"`
	IsEnabled                       pulumi.BoolPtrInput   `pulumi:"isEnabled"`
}

func (DirectLineSpeechChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannelPropertiesResponse)(nil)).Elem()
}

func (i DirectLineSpeechChannelPropertiesResponseArgs) ToDirectLineSpeechChannelPropertiesResponseOutput() DirectLineSpeechChannelPropertiesResponseOutput {
	return i.ToDirectLineSpeechChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i DirectLineSpeechChannelPropertiesResponseArgs) ToDirectLineSpeechChannelPropertiesResponseOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelPropertiesResponseOutput)
}

func (i DirectLineSpeechChannelPropertiesResponseArgs) ToDirectLineSpeechChannelPropertiesResponsePtrOutput() DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return i.ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i DirectLineSpeechChannelPropertiesResponseArgs) ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelPropertiesResponseOutput).ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type DirectLineSpeechChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToDirectLineSpeechChannelPropertiesResponsePtrOutput() DirectLineSpeechChannelPropertiesResponsePtrOutput
	ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(context.Context) DirectLineSpeechChannelPropertiesResponsePtrOutput
}

type directLineSpeechChannelPropertiesResponsePtrType DirectLineSpeechChannelPropertiesResponseArgs

func DirectLineSpeechChannelPropertiesResponsePtr(v *DirectLineSpeechChannelPropertiesResponseArgs) DirectLineSpeechChannelPropertiesResponsePtrInput {
	return (*directLineSpeechChannelPropertiesResponsePtrType)(v)
}

func (*directLineSpeechChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineSpeechChannelPropertiesResponse)(nil)).Elem()
}

func (i *directLineSpeechChannelPropertiesResponsePtrType) ToDirectLineSpeechChannelPropertiesResponsePtrOutput() DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return i.ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *directLineSpeechChannelPropertiesResponsePtrType) ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelPropertiesResponsePtrOutput)
}

type DirectLineSpeechChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (DirectLineSpeechChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannelPropertiesResponse)(nil)).Elem()
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) ToDirectLineSpeechChannelPropertiesResponseOutput() DirectLineSpeechChannelPropertiesResponseOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) ToDirectLineSpeechChannelPropertiesResponseOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesResponseOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) ToDirectLineSpeechChannelPropertiesResponsePtrOutput() DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return o.ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v DirectLineSpeechChannelPropertiesResponse) *DirectLineSpeechChannelPropertiesResponse {
		return &v
	}).(DirectLineSpeechChannelPropertiesResponsePtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) CognitiveServiceRegion() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelPropertiesResponse) string { return v.CognitiveServiceRegion }).(pulumi.StringOutput)
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) CognitiveServiceSubscriptionKey() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelPropertiesResponse) string { return v.CognitiveServiceSubscriptionKey }).(pulumi.StringOutput)
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) CustomSpeechModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelPropertiesResponse) *string { return v.CustomSpeechModelId }).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) CustomVoiceDeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelPropertiesResponse) *string { return v.CustomVoiceDeploymentId }).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) IsDefaultBotForCogSvcAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelPropertiesResponse) *bool { return v.IsDefaultBotForCogSvcAccount }).(pulumi.BoolPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelPropertiesResponse) *bool { return v.IsEnabled }).(pulumi.BoolPtrOutput)
}

type DirectLineSpeechChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (DirectLineSpeechChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**DirectLineSpeechChannelPropertiesResponse)(nil)).Elem()
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) ToDirectLineSpeechChannelPropertiesResponsePtrOutput() DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) ToDirectLineSpeechChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return o
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) Elem() DirectLineSpeechChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) DirectLineSpeechChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret DirectLineSpeechChannelPropertiesResponse
		return ret
	}).(DirectLineSpeechChannelPropertiesResponseOutput)
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) CognitiveServiceRegion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CognitiveServiceRegion
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) CognitiveServiceSubscriptionKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CognitiveServiceSubscriptionKey
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) CustomSpeechModelId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomSpeechModelId
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) CustomVoiceDeploymentId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CustomVoiceDeploymentId
	}).(pulumi.StringPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) IsDefaultBotForCogSvcAccount() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsDefaultBotForCogSvcAccount
	}).(pulumi.BoolPtrOutput)
}

func (o DirectLineSpeechChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *DirectLineSpeechChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type DirectLineSpeechChannelResponse struct {
	ChannelName string                                     `pulumi:"channelName"`
	Properties  *DirectLineSpeechChannelPropertiesResponse `pulumi:"properties"`
}





type DirectLineSpeechChannelResponseInput interface {
	pulumi.Input

	ToDirectLineSpeechChannelResponseOutput() DirectLineSpeechChannelResponseOutput
	ToDirectLineSpeechChannelResponseOutputWithContext(context.Context) DirectLineSpeechChannelResponseOutput
}

type DirectLineSpeechChannelResponseArgs struct {
	ChannelName pulumi.StringInput                                `pulumi:"channelName"`
	Properties  DirectLineSpeechChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (DirectLineSpeechChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannelResponse)(nil)).Elem()
}

func (i DirectLineSpeechChannelResponseArgs) ToDirectLineSpeechChannelResponseOutput() DirectLineSpeechChannelResponseOutput {
	return i.ToDirectLineSpeechChannelResponseOutputWithContext(context.Background())
}

func (i DirectLineSpeechChannelResponseArgs) ToDirectLineSpeechChannelResponseOutputWithContext(ctx context.Context) DirectLineSpeechChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(DirectLineSpeechChannelResponseOutput)
}

type DirectLineSpeechChannelResponseOutput struct{ *pulumi.OutputState }

func (DirectLineSpeechChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*DirectLineSpeechChannelResponse)(nil)).Elem()
}

func (o DirectLineSpeechChannelResponseOutput) ToDirectLineSpeechChannelResponseOutput() DirectLineSpeechChannelResponseOutput {
	return o
}

func (o DirectLineSpeechChannelResponseOutput) ToDirectLineSpeechChannelResponseOutputWithContext(ctx context.Context) DirectLineSpeechChannelResponseOutput {
	return o
}

func (o DirectLineSpeechChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o DirectLineSpeechChannelResponseOutput) Properties() DirectLineSpeechChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v DirectLineSpeechChannelResponse) *DirectLineSpeechChannelPropertiesResponse {
		return v.Properties
	}).(DirectLineSpeechChannelPropertiesResponsePtrOutput)
}

type EmailChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Properties  *EmailChannelProperties `pulumi:"properties"`
}





type EmailChannelInput interface {
	pulumi.Input

	ToEmailChannelOutput() EmailChannelOutput
	ToEmailChannelOutputWithContext(context.Context) EmailChannelOutput
}

type EmailChannelArgs struct {
	ChannelName pulumi.StringInput             `pulumi:"channelName"`
	Properties  EmailChannelPropertiesPtrInput `pulumi:"properties"`
}

func (EmailChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannel)(nil)).Elem()
}

func (i EmailChannelArgs) ToEmailChannelOutput() EmailChannelOutput {
	return i.ToEmailChannelOutputWithContext(context.Background())
}

func (i EmailChannelArgs) ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelOutput)
}

type EmailChannelOutput struct{ *pulumi.OutputState }

func (EmailChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannel)(nil)).Elem()
}

func (o EmailChannelOutput) ToEmailChannelOutput() EmailChannelOutput {
	return o
}

func (o EmailChannelOutput) ToEmailChannelOutputWithContext(ctx context.Context) EmailChannelOutput {
	return o
}

func (o EmailChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v EmailChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o EmailChannelOutput) Properties() EmailChannelPropertiesPtrOutput {
	return o.ApplyT(func(v EmailChannel) *EmailChannelProperties { return v.Properties }).(EmailChannelPropertiesPtrOutput)
}

type EmailChannelProperties struct {
	EmailAddress string  `pulumi:"emailAddress"`
	IsEnabled    bool    `pulumi:"isEnabled"`
	Password     *string `pulumi:"password"`
}





type EmailChannelPropertiesInput interface {
	pulumi.Input

	ToEmailChannelPropertiesOutput() EmailChannelPropertiesOutput
	ToEmailChannelPropertiesOutputWithContext(context.Context) EmailChannelPropertiesOutput
}

type EmailChannelPropertiesArgs struct {
	EmailAddress pulumi.StringInput    `pulumi:"emailAddress"`
	IsEnabled    pulumi.BoolInput      `pulumi:"isEnabled"`
	Password     pulumi.StringPtrInput `pulumi:"password"`
}

func (EmailChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannelProperties)(nil)).Elem()
}

func (i EmailChannelPropertiesArgs) ToEmailChannelPropertiesOutput() EmailChannelPropertiesOutput {
	return i.ToEmailChannelPropertiesOutputWithContext(context.Background())
}

func (i EmailChannelPropertiesArgs) ToEmailChannelPropertiesOutputWithContext(ctx context.Context) EmailChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPropertiesOutput)
}

func (i EmailChannelPropertiesArgs) ToEmailChannelPropertiesPtrOutput() EmailChannelPropertiesPtrOutput {
	return i.ToEmailChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i EmailChannelPropertiesArgs) ToEmailChannelPropertiesPtrOutputWithContext(ctx context.Context) EmailChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPropertiesOutput).ToEmailChannelPropertiesPtrOutputWithContext(ctx)
}









type EmailChannelPropertiesPtrInput interface {
	pulumi.Input

	ToEmailChannelPropertiesPtrOutput() EmailChannelPropertiesPtrOutput
	ToEmailChannelPropertiesPtrOutputWithContext(context.Context) EmailChannelPropertiesPtrOutput
}

type emailChannelPropertiesPtrType EmailChannelPropertiesArgs

func EmailChannelPropertiesPtr(v *EmailChannelPropertiesArgs) EmailChannelPropertiesPtrInput {
	return (*emailChannelPropertiesPtrType)(v)
}

func (*emailChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannelProperties)(nil)).Elem()
}

func (i *emailChannelPropertiesPtrType) ToEmailChannelPropertiesPtrOutput() EmailChannelPropertiesPtrOutput {
	return i.ToEmailChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *emailChannelPropertiesPtrType) ToEmailChannelPropertiesPtrOutputWithContext(ctx context.Context) EmailChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPropertiesPtrOutput)
}

type EmailChannelPropertiesOutput struct{ *pulumi.OutputState }

func (EmailChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannelProperties)(nil)).Elem()
}

func (o EmailChannelPropertiesOutput) ToEmailChannelPropertiesOutput() EmailChannelPropertiesOutput {
	return o
}

func (o EmailChannelPropertiesOutput) ToEmailChannelPropertiesOutputWithContext(ctx context.Context) EmailChannelPropertiesOutput {
	return o
}

func (o EmailChannelPropertiesOutput) ToEmailChannelPropertiesPtrOutput() EmailChannelPropertiesPtrOutput {
	return o.ToEmailChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o EmailChannelPropertiesOutput) ToEmailChannelPropertiesPtrOutputWithContext(ctx context.Context) EmailChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EmailChannelProperties) *EmailChannelProperties {
		return &v
	}).(EmailChannelPropertiesPtrOutput)
}

func (o EmailChannelPropertiesOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EmailChannelProperties) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o EmailChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v EmailChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o EmailChannelPropertiesOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailChannelProperties) *string { return v.Password }).(pulumi.StringPtrOutput)
}

type EmailChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (EmailChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannelProperties)(nil)).Elem()
}

func (o EmailChannelPropertiesPtrOutput) ToEmailChannelPropertiesPtrOutput() EmailChannelPropertiesPtrOutput {
	return o
}

func (o EmailChannelPropertiesPtrOutput) ToEmailChannelPropertiesPtrOutputWithContext(ctx context.Context) EmailChannelPropertiesPtrOutput {
	return o
}

func (o EmailChannelPropertiesPtrOutput) Elem() EmailChannelPropertiesOutput {
	return o.ApplyT(func(v *EmailChannelProperties) EmailChannelProperties {
		if v != nil {
			return *v
		}
		var ret EmailChannelProperties
		return ret
	}).(EmailChannelPropertiesOutput)
}

func (o EmailChannelPropertiesPtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o EmailChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o EmailChannelPropertiesPtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

type EmailChannelPropertiesResponse struct {
	EmailAddress string  `pulumi:"emailAddress"`
	IsEnabled    bool    `pulumi:"isEnabled"`
	Password     *string `pulumi:"password"`
}





type EmailChannelPropertiesResponseInput interface {
	pulumi.Input

	ToEmailChannelPropertiesResponseOutput() EmailChannelPropertiesResponseOutput
	ToEmailChannelPropertiesResponseOutputWithContext(context.Context) EmailChannelPropertiesResponseOutput
}

type EmailChannelPropertiesResponseArgs struct {
	EmailAddress pulumi.StringInput    `pulumi:"emailAddress"`
	IsEnabled    pulumi.BoolInput      `pulumi:"isEnabled"`
	Password     pulumi.StringPtrInput `pulumi:"password"`
}

func (EmailChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannelPropertiesResponse)(nil)).Elem()
}

func (i EmailChannelPropertiesResponseArgs) ToEmailChannelPropertiesResponseOutput() EmailChannelPropertiesResponseOutput {
	return i.ToEmailChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i EmailChannelPropertiesResponseArgs) ToEmailChannelPropertiesResponseOutputWithContext(ctx context.Context) EmailChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPropertiesResponseOutput)
}

func (i EmailChannelPropertiesResponseArgs) ToEmailChannelPropertiesResponsePtrOutput() EmailChannelPropertiesResponsePtrOutput {
	return i.ToEmailChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i EmailChannelPropertiesResponseArgs) ToEmailChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) EmailChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPropertiesResponseOutput).ToEmailChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type EmailChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToEmailChannelPropertiesResponsePtrOutput() EmailChannelPropertiesResponsePtrOutput
	ToEmailChannelPropertiesResponsePtrOutputWithContext(context.Context) EmailChannelPropertiesResponsePtrOutput
}

type emailChannelPropertiesResponsePtrType EmailChannelPropertiesResponseArgs

func EmailChannelPropertiesResponsePtr(v *EmailChannelPropertiesResponseArgs) EmailChannelPropertiesResponsePtrInput {
	return (*emailChannelPropertiesResponsePtrType)(v)
}

func (*emailChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannelPropertiesResponse)(nil)).Elem()
}

func (i *emailChannelPropertiesResponsePtrType) ToEmailChannelPropertiesResponsePtrOutput() EmailChannelPropertiesResponsePtrOutput {
	return i.ToEmailChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *emailChannelPropertiesResponsePtrType) ToEmailChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) EmailChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelPropertiesResponsePtrOutput)
}

type EmailChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (EmailChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannelPropertiesResponse)(nil)).Elem()
}

func (o EmailChannelPropertiesResponseOutput) ToEmailChannelPropertiesResponseOutput() EmailChannelPropertiesResponseOutput {
	return o
}

func (o EmailChannelPropertiesResponseOutput) ToEmailChannelPropertiesResponseOutputWithContext(ctx context.Context) EmailChannelPropertiesResponseOutput {
	return o
}

func (o EmailChannelPropertiesResponseOutput) ToEmailChannelPropertiesResponsePtrOutput() EmailChannelPropertiesResponsePtrOutput {
	return o.ToEmailChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o EmailChannelPropertiesResponseOutput) ToEmailChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) EmailChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v EmailChannelPropertiesResponse) *EmailChannelPropertiesResponse {
		return &v
	}).(EmailChannelPropertiesResponsePtrOutput)
}

func (o EmailChannelPropertiesResponseOutput) EmailAddress() pulumi.StringOutput {
	return o.ApplyT(func(v EmailChannelPropertiesResponse) string { return v.EmailAddress }).(pulumi.StringOutput)
}

func (o EmailChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v EmailChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o EmailChannelPropertiesResponseOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v EmailChannelPropertiesResponse) *string { return v.Password }).(pulumi.StringPtrOutput)
}

type EmailChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (EmailChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**EmailChannelPropertiesResponse)(nil)).Elem()
}

func (o EmailChannelPropertiesResponsePtrOutput) ToEmailChannelPropertiesResponsePtrOutput() EmailChannelPropertiesResponsePtrOutput {
	return o
}

func (o EmailChannelPropertiesResponsePtrOutput) ToEmailChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) EmailChannelPropertiesResponsePtrOutput {
	return o
}

func (o EmailChannelPropertiesResponsePtrOutput) Elem() EmailChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *EmailChannelPropertiesResponse) EmailChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret EmailChannelPropertiesResponse
		return ret
	}).(EmailChannelPropertiesResponseOutput)
}

func (o EmailChannelPropertiesResponsePtrOutput) EmailAddress() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.EmailAddress
	}).(pulumi.StringPtrOutput)
}

func (o EmailChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *EmailChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o EmailChannelPropertiesResponsePtrOutput) Password() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *EmailChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.Password
	}).(pulumi.StringPtrOutput)
}

type EmailChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *EmailChannelPropertiesResponse `pulumi:"properties"`
}





type EmailChannelResponseInput interface {
	pulumi.Input

	ToEmailChannelResponseOutput() EmailChannelResponseOutput
	ToEmailChannelResponseOutputWithContext(context.Context) EmailChannelResponseOutput
}

type EmailChannelResponseArgs struct {
	ChannelName pulumi.StringInput                     `pulumi:"channelName"`
	Properties  EmailChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (EmailChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannelResponse)(nil)).Elem()
}

func (i EmailChannelResponseArgs) ToEmailChannelResponseOutput() EmailChannelResponseOutput {
	return i.ToEmailChannelResponseOutputWithContext(context.Background())
}

func (i EmailChannelResponseArgs) ToEmailChannelResponseOutputWithContext(ctx context.Context) EmailChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(EmailChannelResponseOutput)
}

type EmailChannelResponseOutput struct{ *pulumi.OutputState }

func (EmailChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*EmailChannelResponse)(nil)).Elem()
}

func (o EmailChannelResponseOutput) ToEmailChannelResponseOutput() EmailChannelResponseOutput {
	return o
}

func (o EmailChannelResponseOutput) ToEmailChannelResponseOutputWithContext(ctx context.Context) EmailChannelResponseOutput {
	return o
}

func (o EmailChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v EmailChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o EmailChannelResponseOutput) Properties() EmailChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v EmailChannelResponse) *EmailChannelPropertiesResponse { return v.Properties }).(EmailChannelPropertiesResponsePtrOutput)
}

type FacebookChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Properties  *FacebookChannelProperties `pulumi:"properties"`
}





type FacebookChannelInput interface {
	pulumi.Input

	ToFacebookChannelOutput() FacebookChannelOutput
	ToFacebookChannelOutputWithContext(context.Context) FacebookChannelOutput
}

type FacebookChannelArgs struct {
	ChannelName pulumi.StringInput                `pulumi:"channelName"`
	Properties  FacebookChannelPropertiesPtrInput `pulumi:"properties"`
}

func (FacebookChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannel)(nil)).Elem()
}

func (i FacebookChannelArgs) ToFacebookChannelOutput() FacebookChannelOutput {
	return i.ToFacebookChannelOutputWithContext(context.Background())
}

func (i FacebookChannelArgs) ToFacebookChannelOutputWithContext(ctx context.Context) FacebookChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelOutput)
}

type FacebookChannelOutput struct{ *pulumi.OutputState }

func (FacebookChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannel)(nil)).Elem()
}

func (o FacebookChannelOutput) ToFacebookChannelOutput() FacebookChannelOutput {
	return o
}

func (o FacebookChannelOutput) ToFacebookChannelOutputWithContext(ctx context.Context) FacebookChannelOutput {
	return o
}

func (o FacebookChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o FacebookChannelOutput) Properties() FacebookChannelPropertiesPtrOutput {
	return o.ApplyT(func(v FacebookChannel) *FacebookChannelProperties { return v.Properties }).(FacebookChannelPropertiesPtrOutput)
}

type FacebookChannelProperties struct {
	AppId     string         `pulumi:"appId"`
	AppSecret *string        `pulumi:"appSecret"`
	IsEnabled bool           `pulumi:"isEnabled"`
	Pages     []FacebookPage `pulumi:"pages"`
}





type FacebookChannelPropertiesInput interface {
	pulumi.Input

	ToFacebookChannelPropertiesOutput() FacebookChannelPropertiesOutput
	ToFacebookChannelPropertiesOutputWithContext(context.Context) FacebookChannelPropertiesOutput
}

type FacebookChannelPropertiesArgs struct {
	AppId     pulumi.StringInput     `pulumi:"appId"`
	AppSecret pulumi.StringPtrInput  `pulumi:"appSecret"`
	IsEnabled pulumi.BoolInput       `pulumi:"isEnabled"`
	Pages     FacebookPageArrayInput `pulumi:"pages"`
}

func (FacebookChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannelProperties)(nil)).Elem()
}

func (i FacebookChannelPropertiesArgs) ToFacebookChannelPropertiesOutput() FacebookChannelPropertiesOutput {
	return i.ToFacebookChannelPropertiesOutputWithContext(context.Background())
}

func (i FacebookChannelPropertiesArgs) ToFacebookChannelPropertiesOutputWithContext(ctx context.Context) FacebookChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelPropertiesOutput)
}

func (i FacebookChannelPropertiesArgs) ToFacebookChannelPropertiesPtrOutput() FacebookChannelPropertiesPtrOutput {
	return i.ToFacebookChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i FacebookChannelPropertiesArgs) ToFacebookChannelPropertiesPtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelPropertiesOutput).ToFacebookChannelPropertiesPtrOutputWithContext(ctx)
}









type FacebookChannelPropertiesPtrInput interface {
	pulumi.Input

	ToFacebookChannelPropertiesPtrOutput() FacebookChannelPropertiesPtrOutput
	ToFacebookChannelPropertiesPtrOutputWithContext(context.Context) FacebookChannelPropertiesPtrOutput
}

type facebookChannelPropertiesPtrType FacebookChannelPropertiesArgs

func FacebookChannelPropertiesPtr(v *FacebookChannelPropertiesArgs) FacebookChannelPropertiesPtrInput {
	return (*facebookChannelPropertiesPtrType)(v)
}

func (*facebookChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FacebookChannelProperties)(nil)).Elem()
}

func (i *facebookChannelPropertiesPtrType) ToFacebookChannelPropertiesPtrOutput() FacebookChannelPropertiesPtrOutput {
	return i.ToFacebookChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *facebookChannelPropertiesPtrType) ToFacebookChannelPropertiesPtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelPropertiesPtrOutput)
}

type FacebookChannelPropertiesOutput struct{ *pulumi.OutputState }

func (FacebookChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannelProperties)(nil)).Elem()
}

func (o FacebookChannelPropertiesOutput) ToFacebookChannelPropertiesOutput() FacebookChannelPropertiesOutput {
	return o
}

func (o FacebookChannelPropertiesOutput) ToFacebookChannelPropertiesOutputWithContext(ctx context.Context) FacebookChannelPropertiesOutput {
	return o
}

func (o FacebookChannelPropertiesOutput) ToFacebookChannelPropertiesPtrOutput() FacebookChannelPropertiesPtrOutput {
	return o.ToFacebookChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o FacebookChannelPropertiesOutput) ToFacebookChannelPropertiesPtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FacebookChannelProperties) *FacebookChannelProperties {
		return &v
	}).(FacebookChannelPropertiesPtrOutput)
}

func (o FacebookChannelPropertiesOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookChannelProperties) string { return v.AppId }).(pulumi.StringOutput)
}

func (o FacebookChannelPropertiesOutput) AppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FacebookChannelProperties) *string { return v.AppSecret }).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FacebookChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o FacebookChannelPropertiesOutput) Pages() FacebookPageArrayOutput {
	return o.ApplyT(func(v FacebookChannelProperties) []FacebookPage { return v.Pages }).(FacebookPageArrayOutput)
}

type FacebookChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (FacebookChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FacebookChannelProperties)(nil)).Elem()
}

func (o FacebookChannelPropertiesPtrOutput) ToFacebookChannelPropertiesPtrOutput() FacebookChannelPropertiesPtrOutput {
	return o
}

func (o FacebookChannelPropertiesPtrOutput) ToFacebookChannelPropertiesPtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesPtrOutput {
	return o
}

func (o FacebookChannelPropertiesPtrOutput) Elem() FacebookChannelPropertiesOutput {
	return o.ApplyT(func(v *FacebookChannelProperties) FacebookChannelProperties {
		if v != nil {
			return *v
		}
		var ret FacebookChannelProperties
		return ret
	}).(FacebookChannelPropertiesOutput)
}

func (o FacebookChannelPropertiesPtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesPtrOutput) AppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.AppSecret
	}).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FacebookChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o FacebookChannelPropertiesPtrOutput) Pages() FacebookPageArrayOutput {
	return o.ApplyT(func(v *FacebookChannelProperties) []FacebookPage {
		if v == nil {
			return nil
		}
		return v.Pages
	}).(FacebookPageArrayOutput)
}

type FacebookChannelPropertiesResponse struct {
	AppId       string                 `pulumi:"appId"`
	AppSecret   *string                `pulumi:"appSecret"`
	CallbackUrl string                 `pulumi:"callbackUrl"`
	IsEnabled   bool                   `pulumi:"isEnabled"`
	Pages       []FacebookPageResponse `pulumi:"pages"`
	VerifyToken string                 `pulumi:"verifyToken"`
}





type FacebookChannelPropertiesResponseInput interface {
	pulumi.Input

	ToFacebookChannelPropertiesResponseOutput() FacebookChannelPropertiesResponseOutput
	ToFacebookChannelPropertiesResponseOutputWithContext(context.Context) FacebookChannelPropertiesResponseOutput
}

type FacebookChannelPropertiesResponseArgs struct {
	AppId       pulumi.StringInput             `pulumi:"appId"`
	AppSecret   pulumi.StringPtrInput          `pulumi:"appSecret"`
	CallbackUrl pulumi.StringInput             `pulumi:"callbackUrl"`
	IsEnabled   pulumi.BoolInput               `pulumi:"isEnabled"`
	Pages       FacebookPageResponseArrayInput `pulumi:"pages"`
	VerifyToken pulumi.StringInput             `pulumi:"verifyToken"`
}

func (FacebookChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannelPropertiesResponse)(nil)).Elem()
}

func (i FacebookChannelPropertiesResponseArgs) ToFacebookChannelPropertiesResponseOutput() FacebookChannelPropertiesResponseOutput {
	return i.ToFacebookChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i FacebookChannelPropertiesResponseArgs) ToFacebookChannelPropertiesResponseOutputWithContext(ctx context.Context) FacebookChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelPropertiesResponseOutput)
}

func (i FacebookChannelPropertiesResponseArgs) ToFacebookChannelPropertiesResponsePtrOutput() FacebookChannelPropertiesResponsePtrOutput {
	return i.ToFacebookChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i FacebookChannelPropertiesResponseArgs) ToFacebookChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelPropertiesResponseOutput).ToFacebookChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type FacebookChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToFacebookChannelPropertiesResponsePtrOutput() FacebookChannelPropertiesResponsePtrOutput
	ToFacebookChannelPropertiesResponsePtrOutputWithContext(context.Context) FacebookChannelPropertiesResponsePtrOutput
}

type facebookChannelPropertiesResponsePtrType FacebookChannelPropertiesResponseArgs

func FacebookChannelPropertiesResponsePtr(v *FacebookChannelPropertiesResponseArgs) FacebookChannelPropertiesResponsePtrInput {
	return (*facebookChannelPropertiesResponsePtrType)(v)
}

func (*facebookChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**FacebookChannelPropertiesResponse)(nil)).Elem()
}

func (i *facebookChannelPropertiesResponsePtrType) ToFacebookChannelPropertiesResponsePtrOutput() FacebookChannelPropertiesResponsePtrOutput {
	return i.ToFacebookChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *facebookChannelPropertiesResponsePtrType) ToFacebookChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelPropertiesResponsePtrOutput)
}

type FacebookChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (FacebookChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannelPropertiesResponse)(nil)).Elem()
}

func (o FacebookChannelPropertiesResponseOutput) ToFacebookChannelPropertiesResponseOutput() FacebookChannelPropertiesResponseOutput {
	return o
}

func (o FacebookChannelPropertiesResponseOutput) ToFacebookChannelPropertiesResponseOutputWithContext(ctx context.Context) FacebookChannelPropertiesResponseOutput {
	return o
}

func (o FacebookChannelPropertiesResponseOutput) ToFacebookChannelPropertiesResponsePtrOutput() FacebookChannelPropertiesResponsePtrOutput {
	return o.ToFacebookChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o FacebookChannelPropertiesResponseOutput) ToFacebookChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v FacebookChannelPropertiesResponse) *FacebookChannelPropertiesResponse {
		return &v
	}).(FacebookChannelPropertiesResponsePtrOutput)
}

func (o FacebookChannelPropertiesResponseOutput) AppId() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookChannelPropertiesResponse) string { return v.AppId }).(pulumi.StringOutput)
}

func (o FacebookChannelPropertiesResponseOutput) AppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FacebookChannelPropertiesResponse) *string { return v.AppSecret }).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesResponseOutput) CallbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookChannelPropertiesResponse) string { return v.CallbackUrl }).(pulumi.StringOutput)
}

func (o FacebookChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v FacebookChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o FacebookChannelPropertiesResponseOutput) Pages() FacebookPageResponseArrayOutput {
	return o.ApplyT(func(v FacebookChannelPropertiesResponse) []FacebookPageResponse { return v.Pages }).(FacebookPageResponseArrayOutput)
}

func (o FacebookChannelPropertiesResponseOutput) VerifyToken() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookChannelPropertiesResponse) string { return v.VerifyToken }).(pulumi.StringOutput)
}

type FacebookChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (FacebookChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**FacebookChannelPropertiesResponse)(nil)).Elem()
}

func (o FacebookChannelPropertiesResponsePtrOutput) ToFacebookChannelPropertiesResponsePtrOutput() FacebookChannelPropertiesResponsePtrOutput {
	return o
}

func (o FacebookChannelPropertiesResponsePtrOutput) ToFacebookChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) FacebookChannelPropertiesResponsePtrOutput {
	return o
}

func (o FacebookChannelPropertiesResponsePtrOutput) Elem() FacebookChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) FacebookChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret FacebookChannelPropertiesResponse
		return ret
	}).(FacebookChannelPropertiesResponseOutput)
}

func (o FacebookChannelPropertiesResponsePtrOutput) AppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AppId
	}).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesResponsePtrOutput) AppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AppSecret
	}).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesResponsePtrOutput) CallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CallbackUrl
	}).(pulumi.StringPtrOutput)
}

func (o FacebookChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o FacebookChannelPropertiesResponsePtrOutput) Pages() FacebookPageResponseArrayOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) []FacebookPageResponse {
		if v == nil {
			return nil
		}
		return v.Pages
	}).(FacebookPageResponseArrayOutput)
}

func (o FacebookChannelPropertiesResponsePtrOutput) VerifyToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *FacebookChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.VerifyToken
	}).(pulumi.StringPtrOutput)
}

type FacebookChannelResponse struct {
	ChannelName string                             `pulumi:"channelName"`
	Properties  *FacebookChannelPropertiesResponse `pulumi:"properties"`
}





type FacebookChannelResponseInput interface {
	pulumi.Input

	ToFacebookChannelResponseOutput() FacebookChannelResponseOutput
	ToFacebookChannelResponseOutputWithContext(context.Context) FacebookChannelResponseOutput
}

type FacebookChannelResponseArgs struct {
	ChannelName pulumi.StringInput                        `pulumi:"channelName"`
	Properties  FacebookChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (FacebookChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannelResponse)(nil)).Elem()
}

func (i FacebookChannelResponseArgs) ToFacebookChannelResponseOutput() FacebookChannelResponseOutput {
	return i.ToFacebookChannelResponseOutputWithContext(context.Background())
}

func (i FacebookChannelResponseArgs) ToFacebookChannelResponseOutputWithContext(ctx context.Context) FacebookChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookChannelResponseOutput)
}

type FacebookChannelResponseOutput struct{ *pulumi.OutputState }

func (FacebookChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookChannelResponse)(nil)).Elem()
}

func (o FacebookChannelResponseOutput) ToFacebookChannelResponseOutput() FacebookChannelResponseOutput {
	return o
}

func (o FacebookChannelResponseOutput) ToFacebookChannelResponseOutputWithContext(ctx context.Context) FacebookChannelResponseOutput {
	return o
}

func (o FacebookChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o FacebookChannelResponseOutput) Properties() FacebookChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v FacebookChannelResponse) *FacebookChannelPropertiesResponse { return v.Properties }).(FacebookChannelPropertiesResponsePtrOutput)
}

type FacebookPage struct {
	AccessToken *string `pulumi:"accessToken"`
	Id          string  `pulumi:"id"`
}





type FacebookPageInput interface {
	pulumi.Input

	ToFacebookPageOutput() FacebookPageOutput
	ToFacebookPageOutputWithContext(context.Context) FacebookPageOutput
}

type FacebookPageArgs struct {
	AccessToken pulumi.StringPtrInput `pulumi:"accessToken"`
	Id          pulumi.StringInput    `pulumi:"id"`
}

func (FacebookPageArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookPage)(nil)).Elem()
}

func (i FacebookPageArgs) ToFacebookPageOutput() FacebookPageOutput {
	return i.ToFacebookPageOutputWithContext(context.Background())
}

func (i FacebookPageArgs) ToFacebookPageOutputWithContext(ctx context.Context) FacebookPageOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookPageOutput)
}





type FacebookPageArrayInput interface {
	pulumi.Input

	ToFacebookPageArrayOutput() FacebookPageArrayOutput
	ToFacebookPageArrayOutputWithContext(context.Context) FacebookPageArrayOutput
}

type FacebookPageArray []FacebookPageInput

func (FacebookPageArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FacebookPage)(nil)).Elem()
}

func (i FacebookPageArray) ToFacebookPageArrayOutput() FacebookPageArrayOutput {
	return i.ToFacebookPageArrayOutputWithContext(context.Background())
}

func (i FacebookPageArray) ToFacebookPageArrayOutputWithContext(ctx context.Context) FacebookPageArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookPageArrayOutput)
}

type FacebookPageOutput struct{ *pulumi.OutputState }

func (FacebookPageOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookPage)(nil)).Elem()
}

func (o FacebookPageOutput) ToFacebookPageOutput() FacebookPageOutput {
	return o
}

func (o FacebookPageOutput) ToFacebookPageOutputWithContext(ctx context.Context) FacebookPageOutput {
	return o
}

func (o FacebookPageOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FacebookPage) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o FacebookPageOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookPage) string { return v.Id }).(pulumi.StringOutput)
}

type FacebookPageArrayOutput struct{ *pulumi.OutputState }

func (FacebookPageArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FacebookPage)(nil)).Elem()
}

func (o FacebookPageArrayOutput) ToFacebookPageArrayOutput() FacebookPageArrayOutput {
	return o
}

func (o FacebookPageArrayOutput) ToFacebookPageArrayOutputWithContext(ctx context.Context) FacebookPageArrayOutput {
	return o
}

func (o FacebookPageArrayOutput) Index(i pulumi.IntInput) FacebookPageOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FacebookPage {
		return vs[0].([]FacebookPage)[vs[1].(int)]
	}).(FacebookPageOutput)
}

type FacebookPageResponse struct {
	AccessToken *string `pulumi:"accessToken"`
	Id          string  `pulumi:"id"`
}





type FacebookPageResponseInput interface {
	pulumi.Input

	ToFacebookPageResponseOutput() FacebookPageResponseOutput
	ToFacebookPageResponseOutputWithContext(context.Context) FacebookPageResponseOutput
}

type FacebookPageResponseArgs struct {
	AccessToken pulumi.StringPtrInput `pulumi:"accessToken"`
	Id          pulumi.StringInput    `pulumi:"id"`
}

func (FacebookPageResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookPageResponse)(nil)).Elem()
}

func (i FacebookPageResponseArgs) ToFacebookPageResponseOutput() FacebookPageResponseOutput {
	return i.ToFacebookPageResponseOutputWithContext(context.Background())
}

func (i FacebookPageResponseArgs) ToFacebookPageResponseOutputWithContext(ctx context.Context) FacebookPageResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookPageResponseOutput)
}





type FacebookPageResponseArrayInput interface {
	pulumi.Input

	ToFacebookPageResponseArrayOutput() FacebookPageResponseArrayOutput
	ToFacebookPageResponseArrayOutputWithContext(context.Context) FacebookPageResponseArrayOutput
}

type FacebookPageResponseArray []FacebookPageResponseInput

func (FacebookPageResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FacebookPageResponse)(nil)).Elem()
}

func (i FacebookPageResponseArray) ToFacebookPageResponseArrayOutput() FacebookPageResponseArrayOutput {
	return i.ToFacebookPageResponseArrayOutputWithContext(context.Background())
}

func (i FacebookPageResponseArray) ToFacebookPageResponseArrayOutputWithContext(ctx context.Context) FacebookPageResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(FacebookPageResponseArrayOutput)
}

type FacebookPageResponseOutput struct{ *pulumi.OutputState }

func (FacebookPageResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*FacebookPageResponse)(nil)).Elem()
}

func (o FacebookPageResponseOutput) ToFacebookPageResponseOutput() FacebookPageResponseOutput {
	return o
}

func (o FacebookPageResponseOutput) ToFacebookPageResponseOutputWithContext(ctx context.Context) FacebookPageResponseOutput {
	return o
}

func (o FacebookPageResponseOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v FacebookPageResponse) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o FacebookPageResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v FacebookPageResponse) string { return v.Id }).(pulumi.StringOutput)
}

type FacebookPageResponseArrayOutput struct{ *pulumi.OutputState }

func (FacebookPageResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]FacebookPageResponse)(nil)).Elem()
}

func (o FacebookPageResponseArrayOutput) ToFacebookPageResponseArrayOutput() FacebookPageResponseArrayOutput {
	return o
}

func (o FacebookPageResponseArrayOutput) ToFacebookPageResponseArrayOutputWithContext(ctx context.Context) FacebookPageResponseArrayOutput {
	return o
}

func (o FacebookPageResponseArrayOutput) Index(i pulumi.IntInput) FacebookPageResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) FacebookPageResponse {
		return vs[0].([]FacebookPageResponse)[vs[1].(int)]
	}).(FacebookPageResponseOutput)
}

type KikChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Properties  *KikChannelProperties `pulumi:"properties"`
}





type KikChannelInput interface {
	pulumi.Input

	ToKikChannelOutput() KikChannelOutput
	ToKikChannelOutputWithContext(context.Context) KikChannelOutput
}

type KikChannelArgs struct {
	ChannelName pulumi.StringInput           `pulumi:"channelName"`
	Properties  KikChannelPropertiesPtrInput `pulumi:"properties"`
}

func (KikChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannel)(nil)).Elem()
}

func (i KikChannelArgs) ToKikChannelOutput() KikChannelOutput {
	return i.ToKikChannelOutputWithContext(context.Background())
}

func (i KikChannelArgs) ToKikChannelOutputWithContext(ctx context.Context) KikChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelOutput)
}

type KikChannelOutput struct{ *pulumi.OutputState }

func (KikChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannel)(nil)).Elem()
}

func (o KikChannelOutput) ToKikChannelOutput() KikChannelOutput {
	return o
}

func (o KikChannelOutput) ToKikChannelOutputWithContext(ctx context.Context) KikChannelOutput {
	return o
}

func (o KikChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v KikChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o KikChannelOutput) Properties() KikChannelPropertiesPtrOutput {
	return o.ApplyT(func(v KikChannel) *KikChannelProperties { return v.Properties }).(KikChannelPropertiesPtrOutput)
}

type KikChannelProperties struct {
	ApiKey      *string `pulumi:"apiKey"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	UserName    string  `pulumi:"userName"`
}





type KikChannelPropertiesInput interface {
	pulumi.Input

	ToKikChannelPropertiesOutput() KikChannelPropertiesOutput
	ToKikChannelPropertiesOutputWithContext(context.Context) KikChannelPropertiesOutput
}

type KikChannelPropertiesArgs struct {
	ApiKey      pulumi.StringPtrInput `pulumi:"apiKey"`
	IsEnabled   pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated pulumi.BoolPtrInput   `pulumi:"isValidated"`
	UserName    pulumi.StringInput    `pulumi:"userName"`
}

func (KikChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannelProperties)(nil)).Elem()
}

func (i KikChannelPropertiesArgs) ToKikChannelPropertiesOutput() KikChannelPropertiesOutput {
	return i.ToKikChannelPropertiesOutputWithContext(context.Background())
}

func (i KikChannelPropertiesArgs) ToKikChannelPropertiesOutputWithContext(ctx context.Context) KikChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelPropertiesOutput)
}

func (i KikChannelPropertiesArgs) ToKikChannelPropertiesPtrOutput() KikChannelPropertiesPtrOutput {
	return i.ToKikChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i KikChannelPropertiesArgs) ToKikChannelPropertiesPtrOutputWithContext(ctx context.Context) KikChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelPropertiesOutput).ToKikChannelPropertiesPtrOutputWithContext(ctx)
}









type KikChannelPropertiesPtrInput interface {
	pulumi.Input

	ToKikChannelPropertiesPtrOutput() KikChannelPropertiesPtrOutput
	ToKikChannelPropertiesPtrOutputWithContext(context.Context) KikChannelPropertiesPtrOutput
}

type kikChannelPropertiesPtrType KikChannelPropertiesArgs

func KikChannelPropertiesPtr(v *KikChannelPropertiesArgs) KikChannelPropertiesPtrInput {
	return (*kikChannelPropertiesPtrType)(v)
}

func (*kikChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KikChannelProperties)(nil)).Elem()
}

func (i *kikChannelPropertiesPtrType) ToKikChannelPropertiesPtrOutput() KikChannelPropertiesPtrOutput {
	return i.ToKikChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *kikChannelPropertiesPtrType) ToKikChannelPropertiesPtrOutputWithContext(ctx context.Context) KikChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelPropertiesPtrOutput)
}

type KikChannelPropertiesOutput struct{ *pulumi.OutputState }

func (KikChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannelProperties)(nil)).Elem()
}

func (o KikChannelPropertiesOutput) ToKikChannelPropertiesOutput() KikChannelPropertiesOutput {
	return o
}

func (o KikChannelPropertiesOutput) ToKikChannelPropertiesOutputWithContext(ctx context.Context) KikChannelPropertiesOutput {
	return o
}

func (o KikChannelPropertiesOutput) ToKikChannelPropertiesPtrOutput() KikChannelPropertiesPtrOutput {
	return o.ToKikChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o KikChannelPropertiesOutput) ToKikChannelPropertiesPtrOutputWithContext(ctx context.Context) KikChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KikChannelProperties) *KikChannelProperties {
		return &v
	}).(KikChannelPropertiesPtrOutput)
}

func (o KikChannelPropertiesOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KikChannelProperties) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o KikChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v KikChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o KikChannelPropertiesOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KikChannelProperties) *bool { return v.IsValidated }).(pulumi.BoolPtrOutput)
}

func (o KikChannelPropertiesOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v KikChannelProperties) string { return v.UserName }).(pulumi.StringOutput)
}

type KikChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (KikChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KikChannelProperties)(nil)).Elem()
}

func (o KikChannelPropertiesPtrOutput) ToKikChannelPropertiesPtrOutput() KikChannelPropertiesPtrOutput {
	return o
}

func (o KikChannelPropertiesPtrOutput) ToKikChannelPropertiesPtrOutputWithContext(ctx context.Context) KikChannelPropertiesPtrOutput {
	return o
}

func (o KikChannelPropertiesPtrOutput) Elem() KikChannelPropertiesOutput {
	return o.ApplyT(func(v *KikChannelProperties) KikChannelProperties {
		if v != nil {
			return *v
		}
		var ret KikChannelProperties
		return ret
	}).(KikChannelPropertiesOutput)
}

func (o KikChannelPropertiesPtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KikChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o KikChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KikChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o KikChannelPropertiesPtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KikChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

func (o KikChannelPropertiesPtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KikChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
}

type KikChannelPropertiesResponse struct {
	ApiKey      *string `pulumi:"apiKey"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	UserName    string  `pulumi:"userName"`
}





type KikChannelPropertiesResponseInput interface {
	pulumi.Input

	ToKikChannelPropertiesResponseOutput() KikChannelPropertiesResponseOutput
	ToKikChannelPropertiesResponseOutputWithContext(context.Context) KikChannelPropertiesResponseOutput
}

type KikChannelPropertiesResponseArgs struct {
	ApiKey      pulumi.StringPtrInput `pulumi:"apiKey"`
	IsEnabled   pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated pulumi.BoolPtrInput   `pulumi:"isValidated"`
	UserName    pulumi.StringInput    `pulumi:"userName"`
}

func (KikChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannelPropertiesResponse)(nil)).Elem()
}

func (i KikChannelPropertiesResponseArgs) ToKikChannelPropertiesResponseOutput() KikChannelPropertiesResponseOutput {
	return i.ToKikChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i KikChannelPropertiesResponseArgs) ToKikChannelPropertiesResponseOutputWithContext(ctx context.Context) KikChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelPropertiesResponseOutput)
}

func (i KikChannelPropertiesResponseArgs) ToKikChannelPropertiesResponsePtrOutput() KikChannelPropertiesResponsePtrOutput {
	return i.ToKikChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i KikChannelPropertiesResponseArgs) ToKikChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) KikChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelPropertiesResponseOutput).ToKikChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type KikChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToKikChannelPropertiesResponsePtrOutput() KikChannelPropertiesResponsePtrOutput
	ToKikChannelPropertiesResponsePtrOutputWithContext(context.Context) KikChannelPropertiesResponsePtrOutput
}

type kikChannelPropertiesResponsePtrType KikChannelPropertiesResponseArgs

func KikChannelPropertiesResponsePtr(v *KikChannelPropertiesResponseArgs) KikChannelPropertiesResponsePtrInput {
	return (*kikChannelPropertiesResponsePtrType)(v)
}

func (*kikChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**KikChannelPropertiesResponse)(nil)).Elem()
}

func (i *kikChannelPropertiesResponsePtrType) ToKikChannelPropertiesResponsePtrOutput() KikChannelPropertiesResponsePtrOutput {
	return i.ToKikChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *kikChannelPropertiesResponsePtrType) ToKikChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) KikChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelPropertiesResponsePtrOutput)
}

type KikChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (KikChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannelPropertiesResponse)(nil)).Elem()
}

func (o KikChannelPropertiesResponseOutput) ToKikChannelPropertiesResponseOutput() KikChannelPropertiesResponseOutput {
	return o
}

func (o KikChannelPropertiesResponseOutput) ToKikChannelPropertiesResponseOutputWithContext(ctx context.Context) KikChannelPropertiesResponseOutput {
	return o
}

func (o KikChannelPropertiesResponseOutput) ToKikChannelPropertiesResponsePtrOutput() KikChannelPropertiesResponsePtrOutput {
	return o.ToKikChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o KikChannelPropertiesResponseOutput) ToKikChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) KikChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v KikChannelPropertiesResponse) *KikChannelPropertiesResponse {
		return &v
	}).(KikChannelPropertiesResponsePtrOutput)
}

func (o KikChannelPropertiesResponseOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v KikChannelPropertiesResponse) *string { return v.ApiKey }).(pulumi.StringPtrOutput)
}

func (o KikChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v KikChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o KikChannelPropertiesResponseOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v KikChannelPropertiesResponse) *bool { return v.IsValidated }).(pulumi.BoolPtrOutput)
}

func (o KikChannelPropertiesResponseOutput) UserName() pulumi.StringOutput {
	return o.ApplyT(func(v KikChannelPropertiesResponse) string { return v.UserName }).(pulumi.StringOutput)
}

type KikChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (KikChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**KikChannelPropertiesResponse)(nil)).Elem()
}

func (o KikChannelPropertiesResponsePtrOutput) ToKikChannelPropertiesResponsePtrOutput() KikChannelPropertiesResponsePtrOutput {
	return o
}

func (o KikChannelPropertiesResponsePtrOutput) ToKikChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) KikChannelPropertiesResponsePtrOutput {
	return o
}

func (o KikChannelPropertiesResponsePtrOutput) Elem() KikChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *KikChannelPropertiesResponse) KikChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret KikChannelPropertiesResponse
		return ret
	}).(KikChannelPropertiesResponseOutput)
}

func (o KikChannelPropertiesResponsePtrOutput) ApiKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KikChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ApiKey
	}).(pulumi.StringPtrOutput)
}

func (o KikChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KikChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o KikChannelPropertiesResponsePtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *KikChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

func (o KikChannelPropertiesResponsePtrOutput) UserName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *KikChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.UserName
	}).(pulumi.StringPtrOutput)
}

type KikChannelResponse struct {
	ChannelName string                        `pulumi:"channelName"`
	Properties  *KikChannelPropertiesResponse `pulumi:"properties"`
}





type KikChannelResponseInput interface {
	pulumi.Input

	ToKikChannelResponseOutput() KikChannelResponseOutput
	ToKikChannelResponseOutputWithContext(context.Context) KikChannelResponseOutput
}

type KikChannelResponseArgs struct {
	ChannelName pulumi.StringInput                   `pulumi:"channelName"`
	Properties  KikChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (KikChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannelResponse)(nil)).Elem()
}

func (i KikChannelResponseArgs) ToKikChannelResponseOutput() KikChannelResponseOutput {
	return i.ToKikChannelResponseOutputWithContext(context.Background())
}

func (i KikChannelResponseArgs) ToKikChannelResponseOutputWithContext(ctx context.Context) KikChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(KikChannelResponseOutput)
}

type KikChannelResponseOutput struct{ *pulumi.OutputState }

func (KikChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*KikChannelResponse)(nil)).Elem()
}

func (o KikChannelResponseOutput) ToKikChannelResponseOutput() KikChannelResponseOutput {
	return o
}

func (o KikChannelResponseOutput) ToKikChannelResponseOutputWithContext(ctx context.Context) KikChannelResponseOutput {
	return o
}

func (o KikChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v KikChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o KikChannelResponseOutput) Properties() KikChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v KikChannelResponse) *KikChannelPropertiesResponse { return v.Properties }).(KikChannelPropertiesResponsePtrOutput)
}

type LineChannel struct {
	ChannelName string                 `pulumi:"channelName"`
	Properties  *LineChannelProperties `pulumi:"properties"`
}





type LineChannelInput interface {
	pulumi.Input

	ToLineChannelOutput() LineChannelOutput
	ToLineChannelOutputWithContext(context.Context) LineChannelOutput
}

type LineChannelArgs struct {
	ChannelName pulumi.StringInput            `pulumi:"channelName"`
	Properties  LineChannelPropertiesPtrInput `pulumi:"properties"`
}

func (LineChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannel)(nil)).Elem()
}

func (i LineChannelArgs) ToLineChannelOutput() LineChannelOutput {
	return i.ToLineChannelOutputWithContext(context.Background())
}

func (i LineChannelArgs) ToLineChannelOutputWithContext(ctx context.Context) LineChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelOutput)
}

type LineChannelOutput struct{ *pulumi.OutputState }

func (LineChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannel)(nil)).Elem()
}

func (o LineChannelOutput) ToLineChannelOutput() LineChannelOutput {
	return o
}

func (o LineChannelOutput) ToLineChannelOutputWithContext(ctx context.Context) LineChannelOutput {
	return o
}

func (o LineChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v LineChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o LineChannelOutput) Properties() LineChannelPropertiesPtrOutput {
	return o.ApplyT(func(v LineChannel) *LineChannelProperties { return v.Properties }).(LineChannelPropertiesPtrOutput)
}

type LineChannelProperties struct {
	LineRegistrations []LineRegistration `pulumi:"lineRegistrations"`
}





type LineChannelPropertiesInput interface {
	pulumi.Input

	ToLineChannelPropertiesOutput() LineChannelPropertiesOutput
	ToLineChannelPropertiesOutputWithContext(context.Context) LineChannelPropertiesOutput
}

type LineChannelPropertiesArgs struct {
	LineRegistrations LineRegistrationArrayInput `pulumi:"lineRegistrations"`
}

func (LineChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannelProperties)(nil)).Elem()
}

func (i LineChannelPropertiesArgs) ToLineChannelPropertiesOutput() LineChannelPropertiesOutput {
	return i.ToLineChannelPropertiesOutputWithContext(context.Background())
}

func (i LineChannelPropertiesArgs) ToLineChannelPropertiesOutputWithContext(ctx context.Context) LineChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelPropertiesOutput)
}

func (i LineChannelPropertiesArgs) ToLineChannelPropertiesPtrOutput() LineChannelPropertiesPtrOutput {
	return i.ToLineChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i LineChannelPropertiesArgs) ToLineChannelPropertiesPtrOutputWithContext(ctx context.Context) LineChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelPropertiesOutput).ToLineChannelPropertiesPtrOutputWithContext(ctx)
}









type LineChannelPropertiesPtrInput interface {
	pulumi.Input

	ToLineChannelPropertiesPtrOutput() LineChannelPropertiesPtrOutput
	ToLineChannelPropertiesPtrOutputWithContext(context.Context) LineChannelPropertiesPtrOutput
}

type lineChannelPropertiesPtrType LineChannelPropertiesArgs

func LineChannelPropertiesPtr(v *LineChannelPropertiesArgs) LineChannelPropertiesPtrInput {
	return (*lineChannelPropertiesPtrType)(v)
}

func (*lineChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LineChannelProperties)(nil)).Elem()
}

func (i *lineChannelPropertiesPtrType) ToLineChannelPropertiesPtrOutput() LineChannelPropertiesPtrOutput {
	return i.ToLineChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *lineChannelPropertiesPtrType) ToLineChannelPropertiesPtrOutputWithContext(ctx context.Context) LineChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelPropertiesPtrOutput)
}

type LineChannelPropertiesOutput struct{ *pulumi.OutputState }

func (LineChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannelProperties)(nil)).Elem()
}

func (o LineChannelPropertiesOutput) ToLineChannelPropertiesOutput() LineChannelPropertiesOutput {
	return o
}

func (o LineChannelPropertiesOutput) ToLineChannelPropertiesOutputWithContext(ctx context.Context) LineChannelPropertiesOutput {
	return o
}

func (o LineChannelPropertiesOutput) ToLineChannelPropertiesPtrOutput() LineChannelPropertiesPtrOutput {
	return o.ToLineChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o LineChannelPropertiesOutput) ToLineChannelPropertiesPtrOutputWithContext(ctx context.Context) LineChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LineChannelProperties) *LineChannelProperties {
		return &v
	}).(LineChannelPropertiesPtrOutput)
}

func (o LineChannelPropertiesOutput) LineRegistrations() LineRegistrationArrayOutput {
	return o.ApplyT(func(v LineChannelProperties) []LineRegistration { return v.LineRegistrations }).(LineRegistrationArrayOutput)
}

type LineChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (LineChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LineChannelProperties)(nil)).Elem()
}

func (o LineChannelPropertiesPtrOutput) ToLineChannelPropertiesPtrOutput() LineChannelPropertiesPtrOutput {
	return o
}

func (o LineChannelPropertiesPtrOutput) ToLineChannelPropertiesPtrOutputWithContext(ctx context.Context) LineChannelPropertiesPtrOutput {
	return o
}

func (o LineChannelPropertiesPtrOutput) Elem() LineChannelPropertiesOutput {
	return o.ApplyT(func(v *LineChannelProperties) LineChannelProperties {
		if v != nil {
			return *v
		}
		var ret LineChannelProperties
		return ret
	}).(LineChannelPropertiesOutput)
}

func (o LineChannelPropertiesPtrOutput) LineRegistrations() LineRegistrationArrayOutput {
	return o.ApplyT(func(v *LineChannelProperties) []LineRegistration {
		if v == nil {
			return nil
		}
		return v.LineRegistrations
	}).(LineRegistrationArrayOutput)
}

type LineChannelPropertiesResponse struct {
	CallbackUrl       string                     `pulumi:"callbackUrl"`
	IsValidated       bool                       `pulumi:"isValidated"`
	LineRegistrations []LineRegistrationResponse `pulumi:"lineRegistrations"`
}





type LineChannelPropertiesResponseInput interface {
	pulumi.Input

	ToLineChannelPropertiesResponseOutput() LineChannelPropertiesResponseOutput
	ToLineChannelPropertiesResponseOutputWithContext(context.Context) LineChannelPropertiesResponseOutput
}

type LineChannelPropertiesResponseArgs struct {
	CallbackUrl       pulumi.StringInput                 `pulumi:"callbackUrl"`
	IsValidated       pulumi.BoolInput                   `pulumi:"isValidated"`
	LineRegistrations LineRegistrationResponseArrayInput `pulumi:"lineRegistrations"`
}

func (LineChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannelPropertiesResponse)(nil)).Elem()
}

func (i LineChannelPropertiesResponseArgs) ToLineChannelPropertiesResponseOutput() LineChannelPropertiesResponseOutput {
	return i.ToLineChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i LineChannelPropertiesResponseArgs) ToLineChannelPropertiesResponseOutputWithContext(ctx context.Context) LineChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelPropertiesResponseOutput)
}

func (i LineChannelPropertiesResponseArgs) ToLineChannelPropertiesResponsePtrOutput() LineChannelPropertiesResponsePtrOutput {
	return i.ToLineChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i LineChannelPropertiesResponseArgs) ToLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) LineChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelPropertiesResponseOutput).ToLineChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type LineChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToLineChannelPropertiesResponsePtrOutput() LineChannelPropertiesResponsePtrOutput
	ToLineChannelPropertiesResponsePtrOutputWithContext(context.Context) LineChannelPropertiesResponsePtrOutput
}

type lineChannelPropertiesResponsePtrType LineChannelPropertiesResponseArgs

func LineChannelPropertiesResponsePtr(v *LineChannelPropertiesResponseArgs) LineChannelPropertiesResponsePtrInput {
	return (*lineChannelPropertiesResponsePtrType)(v)
}

func (*lineChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**LineChannelPropertiesResponse)(nil)).Elem()
}

func (i *lineChannelPropertiesResponsePtrType) ToLineChannelPropertiesResponsePtrOutput() LineChannelPropertiesResponsePtrOutput {
	return i.ToLineChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *lineChannelPropertiesResponsePtrType) ToLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) LineChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelPropertiesResponsePtrOutput)
}

type LineChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (LineChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannelPropertiesResponse)(nil)).Elem()
}

func (o LineChannelPropertiesResponseOutput) ToLineChannelPropertiesResponseOutput() LineChannelPropertiesResponseOutput {
	return o
}

func (o LineChannelPropertiesResponseOutput) ToLineChannelPropertiesResponseOutputWithContext(ctx context.Context) LineChannelPropertiesResponseOutput {
	return o
}

func (o LineChannelPropertiesResponseOutput) ToLineChannelPropertiesResponsePtrOutput() LineChannelPropertiesResponsePtrOutput {
	return o.ToLineChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o LineChannelPropertiesResponseOutput) ToLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) LineChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v LineChannelPropertiesResponse) *LineChannelPropertiesResponse {
		return &v
	}).(LineChannelPropertiesResponsePtrOutput)
}

func (o LineChannelPropertiesResponseOutput) CallbackUrl() pulumi.StringOutput {
	return o.ApplyT(func(v LineChannelPropertiesResponse) string { return v.CallbackUrl }).(pulumi.StringOutput)
}

func (o LineChannelPropertiesResponseOutput) IsValidated() pulumi.BoolOutput {
	return o.ApplyT(func(v LineChannelPropertiesResponse) bool { return v.IsValidated }).(pulumi.BoolOutput)
}

func (o LineChannelPropertiesResponseOutput) LineRegistrations() LineRegistrationResponseArrayOutput {
	return o.ApplyT(func(v LineChannelPropertiesResponse) []LineRegistrationResponse { return v.LineRegistrations }).(LineRegistrationResponseArrayOutput)
}

type LineChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (LineChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**LineChannelPropertiesResponse)(nil)).Elem()
}

func (o LineChannelPropertiesResponsePtrOutput) ToLineChannelPropertiesResponsePtrOutput() LineChannelPropertiesResponsePtrOutput {
	return o
}

func (o LineChannelPropertiesResponsePtrOutput) ToLineChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) LineChannelPropertiesResponsePtrOutput {
	return o
}

func (o LineChannelPropertiesResponsePtrOutput) Elem() LineChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *LineChannelPropertiesResponse) LineChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret LineChannelPropertiesResponse
		return ret
	}).(LineChannelPropertiesResponseOutput)
}

func (o LineChannelPropertiesResponsePtrOutput) CallbackUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *LineChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.CallbackUrl
	}).(pulumi.StringPtrOutput)
}

func (o LineChannelPropertiesResponsePtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *LineChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

func (o LineChannelPropertiesResponsePtrOutput) LineRegistrations() LineRegistrationResponseArrayOutput {
	return o.ApplyT(func(v *LineChannelPropertiesResponse) []LineRegistrationResponse {
		if v == nil {
			return nil
		}
		return v.LineRegistrations
	}).(LineRegistrationResponseArrayOutput)
}

type LineChannelResponse struct {
	ChannelName string                         `pulumi:"channelName"`
	Properties  *LineChannelPropertiesResponse `pulumi:"properties"`
}





type LineChannelResponseInput interface {
	pulumi.Input

	ToLineChannelResponseOutput() LineChannelResponseOutput
	ToLineChannelResponseOutputWithContext(context.Context) LineChannelResponseOutput
}

type LineChannelResponseArgs struct {
	ChannelName pulumi.StringInput                    `pulumi:"channelName"`
	Properties  LineChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (LineChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannelResponse)(nil)).Elem()
}

func (i LineChannelResponseArgs) ToLineChannelResponseOutput() LineChannelResponseOutput {
	return i.ToLineChannelResponseOutputWithContext(context.Background())
}

func (i LineChannelResponseArgs) ToLineChannelResponseOutputWithContext(ctx context.Context) LineChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineChannelResponseOutput)
}

type LineChannelResponseOutput struct{ *pulumi.OutputState }

func (LineChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LineChannelResponse)(nil)).Elem()
}

func (o LineChannelResponseOutput) ToLineChannelResponseOutput() LineChannelResponseOutput {
	return o
}

func (o LineChannelResponseOutput) ToLineChannelResponseOutputWithContext(ctx context.Context) LineChannelResponseOutput {
	return o
}

func (o LineChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v LineChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o LineChannelResponseOutput) Properties() LineChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v LineChannelResponse) *LineChannelPropertiesResponse { return v.Properties }).(LineChannelPropertiesResponsePtrOutput)
}

type LineRegistration struct {
	ChannelAccessToken *string `pulumi:"channelAccessToken"`
	ChannelSecret      *string `pulumi:"channelSecret"`
}





type LineRegistrationInput interface {
	pulumi.Input

	ToLineRegistrationOutput() LineRegistrationOutput
	ToLineRegistrationOutputWithContext(context.Context) LineRegistrationOutput
}

type LineRegistrationArgs struct {
	ChannelAccessToken pulumi.StringPtrInput `pulumi:"channelAccessToken"`
	ChannelSecret      pulumi.StringPtrInput `pulumi:"channelSecret"`
}

func (LineRegistrationArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LineRegistration)(nil)).Elem()
}

func (i LineRegistrationArgs) ToLineRegistrationOutput() LineRegistrationOutput {
	return i.ToLineRegistrationOutputWithContext(context.Background())
}

func (i LineRegistrationArgs) ToLineRegistrationOutputWithContext(ctx context.Context) LineRegistrationOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineRegistrationOutput)
}





type LineRegistrationArrayInput interface {
	pulumi.Input

	ToLineRegistrationArrayOutput() LineRegistrationArrayOutput
	ToLineRegistrationArrayOutputWithContext(context.Context) LineRegistrationArrayOutput
}

type LineRegistrationArray []LineRegistrationInput

func (LineRegistrationArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LineRegistration)(nil)).Elem()
}

func (i LineRegistrationArray) ToLineRegistrationArrayOutput() LineRegistrationArrayOutput {
	return i.ToLineRegistrationArrayOutputWithContext(context.Background())
}

func (i LineRegistrationArray) ToLineRegistrationArrayOutputWithContext(ctx context.Context) LineRegistrationArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineRegistrationArrayOutput)
}

type LineRegistrationOutput struct{ *pulumi.OutputState }

func (LineRegistrationOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LineRegistration)(nil)).Elem()
}

func (o LineRegistrationOutput) ToLineRegistrationOutput() LineRegistrationOutput {
	return o
}

func (o LineRegistrationOutput) ToLineRegistrationOutputWithContext(ctx context.Context) LineRegistrationOutput {
	return o
}

func (o LineRegistrationOutput) ChannelAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LineRegistration) *string { return v.ChannelAccessToken }).(pulumi.StringPtrOutput)
}

func (o LineRegistrationOutput) ChannelSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LineRegistration) *string { return v.ChannelSecret }).(pulumi.StringPtrOutput)
}

type LineRegistrationArrayOutput struct{ *pulumi.OutputState }

func (LineRegistrationArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LineRegistration)(nil)).Elem()
}

func (o LineRegistrationArrayOutput) ToLineRegistrationArrayOutput() LineRegistrationArrayOutput {
	return o
}

func (o LineRegistrationArrayOutput) ToLineRegistrationArrayOutputWithContext(ctx context.Context) LineRegistrationArrayOutput {
	return o
}

func (o LineRegistrationArrayOutput) Index(i pulumi.IntInput) LineRegistrationOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LineRegistration {
		return vs[0].([]LineRegistration)[vs[1].(int)]
	}).(LineRegistrationOutput)
}

type LineRegistrationResponse struct {
	ChannelAccessToken *string `pulumi:"channelAccessToken"`
	ChannelSecret      *string `pulumi:"channelSecret"`
	GeneratedId        string  `pulumi:"generatedId"`
}





type LineRegistrationResponseInput interface {
	pulumi.Input

	ToLineRegistrationResponseOutput() LineRegistrationResponseOutput
	ToLineRegistrationResponseOutputWithContext(context.Context) LineRegistrationResponseOutput
}

type LineRegistrationResponseArgs struct {
	ChannelAccessToken pulumi.StringPtrInput `pulumi:"channelAccessToken"`
	ChannelSecret      pulumi.StringPtrInput `pulumi:"channelSecret"`
	GeneratedId        pulumi.StringInput    `pulumi:"generatedId"`
}

func (LineRegistrationResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*LineRegistrationResponse)(nil)).Elem()
}

func (i LineRegistrationResponseArgs) ToLineRegistrationResponseOutput() LineRegistrationResponseOutput {
	return i.ToLineRegistrationResponseOutputWithContext(context.Background())
}

func (i LineRegistrationResponseArgs) ToLineRegistrationResponseOutputWithContext(ctx context.Context) LineRegistrationResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineRegistrationResponseOutput)
}





type LineRegistrationResponseArrayInput interface {
	pulumi.Input

	ToLineRegistrationResponseArrayOutput() LineRegistrationResponseArrayOutput
	ToLineRegistrationResponseArrayOutputWithContext(context.Context) LineRegistrationResponseArrayOutput
}

type LineRegistrationResponseArray []LineRegistrationResponseInput

func (LineRegistrationResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LineRegistrationResponse)(nil)).Elem()
}

func (i LineRegistrationResponseArray) ToLineRegistrationResponseArrayOutput() LineRegistrationResponseArrayOutput {
	return i.ToLineRegistrationResponseArrayOutputWithContext(context.Background())
}

func (i LineRegistrationResponseArray) ToLineRegistrationResponseArrayOutputWithContext(ctx context.Context) LineRegistrationResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(LineRegistrationResponseArrayOutput)
}

type LineRegistrationResponseOutput struct{ *pulumi.OutputState }

func (LineRegistrationResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*LineRegistrationResponse)(nil)).Elem()
}

func (o LineRegistrationResponseOutput) ToLineRegistrationResponseOutput() LineRegistrationResponseOutput {
	return o
}

func (o LineRegistrationResponseOutput) ToLineRegistrationResponseOutputWithContext(ctx context.Context) LineRegistrationResponseOutput {
	return o
}

func (o LineRegistrationResponseOutput) ChannelAccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LineRegistrationResponse) *string { return v.ChannelAccessToken }).(pulumi.StringPtrOutput)
}

func (o LineRegistrationResponseOutput) ChannelSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v LineRegistrationResponse) *string { return v.ChannelSecret }).(pulumi.StringPtrOutput)
}

func (o LineRegistrationResponseOutput) GeneratedId() pulumi.StringOutput {
	return o.ApplyT(func(v LineRegistrationResponse) string { return v.GeneratedId }).(pulumi.StringOutput)
}

type LineRegistrationResponseArrayOutput struct{ *pulumi.OutputState }

func (LineRegistrationResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]LineRegistrationResponse)(nil)).Elem()
}

func (o LineRegistrationResponseArrayOutput) ToLineRegistrationResponseArrayOutput() LineRegistrationResponseArrayOutput {
	return o
}

func (o LineRegistrationResponseArrayOutput) ToLineRegistrationResponseArrayOutputWithContext(ctx context.Context) LineRegistrationResponseArrayOutput {
	return o
}

func (o LineRegistrationResponseArrayOutput) Index(i pulumi.IntInput) LineRegistrationResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) LineRegistrationResponse {
		return vs[0].([]LineRegistrationResponse)[vs[1].(int)]
	}).(LineRegistrationResponseOutput)
}

type MsTeamsChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Properties  *MsTeamsChannelProperties `pulumi:"properties"`
}





type MsTeamsChannelInput interface {
	pulumi.Input

	ToMsTeamsChannelOutput() MsTeamsChannelOutput
	ToMsTeamsChannelOutputWithContext(context.Context) MsTeamsChannelOutput
}

type MsTeamsChannelArgs struct {
	ChannelName pulumi.StringInput               `pulumi:"channelName"`
	Properties  MsTeamsChannelPropertiesPtrInput `pulumi:"properties"`
}

func (MsTeamsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannel)(nil)).Elem()
}

func (i MsTeamsChannelArgs) ToMsTeamsChannelOutput() MsTeamsChannelOutput {
	return i.ToMsTeamsChannelOutputWithContext(context.Background())
}

func (i MsTeamsChannelArgs) ToMsTeamsChannelOutputWithContext(ctx context.Context) MsTeamsChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelOutput)
}

type MsTeamsChannelOutput struct{ *pulumi.OutputState }

func (MsTeamsChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannel)(nil)).Elem()
}

func (o MsTeamsChannelOutput) ToMsTeamsChannelOutput() MsTeamsChannelOutput {
	return o
}

func (o MsTeamsChannelOutput) ToMsTeamsChannelOutputWithContext(ctx context.Context) MsTeamsChannelOutput {
	return o
}

func (o MsTeamsChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v MsTeamsChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o MsTeamsChannelOutput) Properties() MsTeamsChannelPropertiesPtrOutput {
	return o.ApplyT(func(v MsTeamsChannel) *MsTeamsChannelProperties { return v.Properties }).(MsTeamsChannelPropertiesPtrOutput)
}

type MsTeamsChannelProperties struct {
	CallingWebHook *string `pulumi:"callingWebHook"`
	EnableCalling  *bool   `pulumi:"enableCalling"`
	IsEnabled      bool    `pulumi:"isEnabled"`
}





type MsTeamsChannelPropertiesInput interface {
	pulumi.Input

	ToMsTeamsChannelPropertiesOutput() MsTeamsChannelPropertiesOutput
	ToMsTeamsChannelPropertiesOutputWithContext(context.Context) MsTeamsChannelPropertiesOutput
}

type MsTeamsChannelPropertiesArgs struct {
	CallingWebHook pulumi.StringPtrInput `pulumi:"callingWebHook"`
	EnableCalling  pulumi.BoolPtrInput   `pulumi:"enableCalling"`
	IsEnabled      pulumi.BoolInput      `pulumi:"isEnabled"`
}

func (MsTeamsChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannelProperties)(nil)).Elem()
}

func (i MsTeamsChannelPropertiesArgs) ToMsTeamsChannelPropertiesOutput() MsTeamsChannelPropertiesOutput {
	return i.ToMsTeamsChannelPropertiesOutputWithContext(context.Background())
}

func (i MsTeamsChannelPropertiesArgs) ToMsTeamsChannelPropertiesOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelPropertiesOutput)
}

func (i MsTeamsChannelPropertiesArgs) ToMsTeamsChannelPropertiesPtrOutput() MsTeamsChannelPropertiesPtrOutput {
	return i.ToMsTeamsChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i MsTeamsChannelPropertiesArgs) ToMsTeamsChannelPropertiesPtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelPropertiesOutput).ToMsTeamsChannelPropertiesPtrOutputWithContext(ctx)
}









type MsTeamsChannelPropertiesPtrInput interface {
	pulumi.Input

	ToMsTeamsChannelPropertiesPtrOutput() MsTeamsChannelPropertiesPtrOutput
	ToMsTeamsChannelPropertiesPtrOutputWithContext(context.Context) MsTeamsChannelPropertiesPtrOutput
}

type msTeamsChannelPropertiesPtrType MsTeamsChannelPropertiesArgs

func MsTeamsChannelPropertiesPtr(v *MsTeamsChannelPropertiesArgs) MsTeamsChannelPropertiesPtrInput {
	return (*msTeamsChannelPropertiesPtrType)(v)
}

func (*msTeamsChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MsTeamsChannelProperties)(nil)).Elem()
}

func (i *msTeamsChannelPropertiesPtrType) ToMsTeamsChannelPropertiesPtrOutput() MsTeamsChannelPropertiesPtrOutput {
	return i.ToMsTeamsChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *msTeamsChannelPropertiesPtrType) ToMsTeamsChannelPropertiesPtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelPropertiesPtrOutput)
}

type MsTeamsChannelPropertiesOutput struct{ *pulumi.OutputState }

func (MsTeamsChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannelProperties)(nil)).Elem()
}

func (o MsTeamsChannelPropertiesOutput) ToMsTeamsChannelPropertiesOutput() MsTeamsChannelPropertiesOutput {
	return o
}

func (o MsTeamsChannelPropertiesOutput) ToMsTeamsChannelPropertiesOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesOutput {
	return o
}

func (o MsTeamsChannelPropertiesOutput) ToMsTeamsChannelPropertiesPtrOutput() MsTeamsChannelPropertiesPtrOutput {
	return o.ToMsTeamsChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o MsTeamsChannelPropertiesOutput) ToMsTeamsChannelPropertiesPtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MsTeamsChannelProperties) *MsTeamsChannelProperties {
		return &v
	}).(MsTeamsChannelPropertiesPtrOutput)
}

func (o MsTeamsChannelPropertiesOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsTeamsChannelProperties) *string { return v.CallingWebHook }).(pulumi.StringPtrOutput)
}

func (o MsTeamsChannelPropertiesOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MsTeamsChannelProperties) *bool { return v.EnableCalling }).(pulumi.BoolPtrOutput)
}

func (o MsTeamsChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v MsTeamsChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

type MsTeamsChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (MsTeamsChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MsTeamsChannelProperties)(nil)).Elem()
}

func (o MsTeamsChannelPropertiesPtrOutput) ToMsTeamsChannelPropertiesPtrOutput() MsTeamsChannelPropertiesPtrOutput {
	return o
}

func (o MsTeamsChannelPropertiesPtrOutput) ToMsTeamsChannelPropertiesPtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesPtrOutput {
	return o
}

func (o MsTeamsChannelPropertiesPtrOutput) Elem() MsTeamsChannelPropertiesOutput {
	return o.ApplyT(func(v *MsTeamsChannelProperties) MsTeamsChannelProperties {
		if v != nil {
			return *v
		}
		var ret MsTeamsChannelProperties
		return ret
	}).(MsTeamsChannelPropertiesOutput)
}

func (o MsTeamsChannelPropertiesPtrOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MsTeamsChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.CallingWebHook
	}).(pulumi.StringPtrOutput)
}

func (o MsTeamsChannelPropertiesPtrOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MsTeamsChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCalling
	}).(pulumi.BoolPtrOutput)
}

func (o MsTeamsChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MsTeamsChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type MsTeamsChannelPropertiesResponse struct {
	CallingWebHook *string `pulumi:"callingWebHook"`
	EnableCalling  *bool   `pulumi:"enableCalling"`
	IsEnabled      bool    `pulumi:"isEnabled"`
}





type MsTeamsChannelPropertiesResponseInput interface {
	pulumi.Input

	ToMsTeamsChannelPropertiesResponseOutput() MsTeamsChannelPropertiesResponseOutput
	ToMsTeamsChannelPropertiesResponseOutputWithContext(context.Context) MsTeamsChannelPropertiesResponseOutput
}

type MsTeamsChannelPropertiesResponseArgs struct {
	CallingWebHook pulumi.StringPtrInput `pulumi:"callingWebHook"`
	EnableCalling  pulumi.BoolPtrInput   `pulumi:"enableCalling"`
	IsEnabled      pulumi.BoolInput      `pulumi:"isEnabled"`
}

func (MsTeamsChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannelPropertiesResponse)(nil)).Elem()
}

func (i MsTeamsChannelPropertiesResponseArgs) ToMsTeamsChannelPropertiesResponseOutput() MsTeamsChannelPropertiesResponseOutput {
	return i.ToMsTeamsChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i MsTeamsChannelPropertiesResponseArgs) ToMsTeamsChannelPropertiesResponseOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelPropertiesResponseOutput)
}

func (i MsTeamsChannelPropertiesResponseArgs) ToMsTeamsChannelPropertiesResponsePtrOutput() MsTeamsChannelPropertiesResponsePtrOutput {
	return i.ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i MsTeamsChannelPropertiesResponseArgs) ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelPropertiesResponseOutput).ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type MsTeamsChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToMsTeamsChannelPropertiesResponsePtrOutput() MsTeamsChannelPropertiesResponsePtrOutput
	ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(context.Context) MsTeamsChannelPropertiesResponsePtrOutput
}

type msTeamsChannelPropertiesResponsePtrType MsTeamsChannelPropertiesResponseArgs

func MsTeamsChannelPropertiesResponsePtr(v *MsTeamsChannelPropertiesResponseArgs) MsTeamsChannelPropertiesResponsePtrInput {
	return (*msTeamsChannelPropertiesResponsePtrType)(v)
}

func (*msTeamsChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**MsTeamsChannelPropertiesResponse)(nil)).Elem()
}

func (i *msTeamsChannelPropertiesResponsePtrType) ToMsTeamsChannelPropertiesResponsePtrOutput() MsTeamsChannelPropertiesResponsePtrOutput {
	return i.ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *msTeamsChannelPropertiesResponsePtrType) ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelPropertiesResponsePtrOutput)
}

type MsTeamsChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (MsTeamsChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannelPropertiesResponse)(nil)).Elem()
}

func (o MsTeamsChannelPropertiesResponseOutput) ToMsTeamsChannelPropertiesResponseOutput() MsTeamsChannelPropertiesResponseOutput {
	return o
}

func (o MsTeamsChannelPropertiesResponseOutput) ToMsTeamsChannelPropertiesResponseOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesResponseOutput {
	return o
}

func (o MsTeamsChannelPropertiesResponseOutput) ToMsTeamsChannelPropertiesResponsePtrOutput() MsTeamsChannelPropertiesResponsePtrOutput {
	return o.ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o MsTeamsChannelPropertiesResponseOutput) ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v MsTeamsChannelPropertiesResponse) *MsTeamsChannelPropertiesResponse {
		return &v
	}).(MsTeamsChannelPropertiesResponsePtrOutput)
}

func (o MsTeamsChannelPropertiesResponseOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v MsTeamsChannelPropertiesResponse) *string { return v.CallingWebHook }).(pulumi.StringPtrOutput)
}

func (o MsTeamsChannelPropertiesResponseOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v MsTeamsChannelPropertiesResponse) *bool { return v.EnableCalling }).(pulumi.BoolPtrOutput)
}

func (o MsTeamsChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v MsTeamsChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

type MsTeamsChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (MsTeamsChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**MsTeamsChannelPropertiesResponse)(nil)).Elem()
}

func (o MsTeamsChannelPropertiesResponsePtrOutput) ToMsTeamsChannelPropertiesResponsePtrOutput() MsTeamsChannelPropertiesResponsePtrOutput {
	return o
}

func (o MsTeamsChannelPropertiesResponsePtrOutput) ToMsTeamsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) MsTeamsChannelPropertiesResponsePtrOutput {
	return o
}

func (o MsTeamsChannelPropertiesResponsePtrOutput) Elem() MsTeamsChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *MsTeamsChannelPropertiesResponse) MsTeamsChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret MsTeamsChannelPropertiesResponse
		return ret
	}).(MsTeamsChannelPropertiesResponseOutput)
}

func (o MsTeamsChannelPropertiesResponsePtrOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *MsTeamsChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CallingWebHook
	}).(pulumi.StringPtrOutput)
}

func (o MsTeamsChannelPropertiesResponsePtrOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MsTeamsChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCalling
	}).(pulumi.BoolPtrOutput)
}

func (o MsTeamsChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *MsTeamsChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type MsTeamsChannelResponse struct {
	ChannelName string                            `pulumi:"channelName"`
	Properties  *MsTeamsChannelPropertiesResponse `pulumi:"properties"`
}





type MsTeamsChannelResponseInput interface {
	pulumi.Input

	ToMsTeamsChannelResponseOutput() MsTeamsChannelResponseOutput
	ToMsTeamsChannelResponseOutputWithContext(context.Context) MsTeamsChannelResponseOutput
}

type MsTeamsChannelResponseArgs struct {
	ChannelName pulumi.StringInput                       `pulumi:"channelName"`
	Properties  MsTeamsChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (MsTeamsChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannelResponse)(nil)).Elem()
}

func (i MsTeamsChannelResponseArgs) ToMsTeamsChannelResponseOutput() MsTeamsChannelResponseOutput {
	return i.ToMsTeamsChannelResponseOutputWithContext(context.Background())
}

func (i MsTeamsChannelResponseArgs) ToMsTeamsChannelResponseOutputWithContext(ctx context.Context) MsTeamsChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(MsTeamsChannelResponseOutput)
}

type MsTeamsChannelResponseOutput struct{ *pulumi.OutputState }

func (MsTeamsChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*MsTeamsChannelResponse)(nil)).Elem()
}

func (o MsTeamsChannelResponseOutput) ToMsTeamsChannelResponseOutput() MsTeamsChannelResponseOutput {
	return o
}

func (o MsTeamsChannelResponseOutput) ToMsTeamsChannelResponseOutputWithContext(ctx context.Context) MsTeamsChannelResponseOutput {
	return o
}

func (o MsTeamsChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v MsTeamsChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o MsTeamsChannelResponseOutput) Properties() MsTeamsChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v MsTeamsChannelResponse) *MsTeamsChannelPropertiesResponse { return v.Properties }).(MsTeamsChannelPropertiesResponsePtrOutput)
}

type ServiceProviderParameterResponse struct {
	Default     string `pulumi:"default"`
	Description string `pulumi:"description"`
	DisplayName string `pulumi:"displayName"`
	HelpUrl     string `pulumi:"helpUrl"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
}





type ServiceProviderParameterResponseInput interface {
	pulumi.Input

	ToServiceProviderParameterResponseOutput() ServiceProviderParameterResponseOutput
	ToServiceProviderParameterResponseOutputWithContext(context.Context) ServiceProviderParameterResponseOutput
}

type ServiceProviderParameterResponseArgs struct {
	Default     pulumi.StringInput `pulumi:"default"`
	Description pulumi.StringInput `pulumi:"description"`
	DisplayName pulumi.StringInput `pulumi:"displayName"`
	HelpUrl     pulumi.StringInput `pulumi:"helpUrl"`
	Name        pulumi.StringInput `pulumi:"name"`
	Type        pulumi.StringInput `pulumi:"type"`
}

func (ServiceProviderParameterResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderParameterResponse)(nil)).Elem()
}

func (i ServiceProviderParameterResponseArgs) ToServiceProviderParameterResponseOutput() ServiceProviderParameterResponseOutput {
	return i.ToServiceProviderParameterResponseOutputWithContext(context.Background())
}

func (i ServiceProviderParameterResponseArgs) ToServiceProviderParameterResponseOutputWithContext(ctx context.Context) ServiceProviderParameterResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderParameterResponseOutput)
}





type ServiceProviderParameterResponseArrayInput interface {
	pulumi.Input

	ToServiceProviderParameterResponseArrayOutput() ServiceProviderParameterResponseArrayOutput
	ToServiceProviderParameterResponseArrayOutputWithContext(context.Context) ServiceProviderParameterResponseArrayOutput
}

type ServiceProviderParameterResponseArray []ServiceProviderParameterResponseInput

func (ServiceProviderParameterResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceProviderParameterResponse)(nil)).Elem()
}

func (i ServiceProviderParameterResponseArray) ToServiceProviderParameterResponseArrayOutput() ServiceProviderParameterResponseArrayOutput {
	return i.ToServiceProviderParameterResponseArrayOutputWithContext(context.Background())
}

func (i ServiceProviderParameterResponseArray) ToServiceProviderParameterResponseArrayOutputWithContext(ctx context.Context) ServiceProviderParameterResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderParameterResponseArrayOutput)
}

type ServiceProviderParameterResponseOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderParameterResponse)(nil)).Elem()
}

func (o ServiceProviderParameterResponseOutput) ToServiceProviderParameterResponseOutput() ServiceProviderParameterResponseOutput {
	return o
}

func (o ServiceProviderParameterResponseOutput) ToServiceProviderParameterResponseOutputWithContext(ctx context.Context) ServiceProviderParameterResponseOutput {
	return o
}

func (o ServiceProviderParameterResponseOutput) Default() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Default }).(pulumi.StringOutput)
}

func (o ServiceProviderParameterResponseOutput) Description() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Description }).(pulumi.StringOutput)
}

func (o ServiceProviderParameterResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ServiceProviderParameterResponseOutput) HelpUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.HelpUrl }).(pulumi.StringOutput)
}

func (o ServiceProviderParameterResponseOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Name }).(pulumi.StringOutput)
}

func (o ServiceProviderParameterResponseOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderParameterResponse) string { return v.Type }).(pulumi.StringOutput)
}

type ServiceProviderParameterResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceProviderParameterResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceProviderParameterResponse)(nil)).Elem()
}

func (o ServiceProviderParameterResponseArrayOutput) ToServiceProviderParameterResponseArrayOutput() ServiceProviderParameterResponseArrayOutput {
	return o
}

func (o ServiceProviderParameterResponseArrayOutput) ToServiceProviderParameterResponseArrayOutputWithContext(ctx context.Context) ServiceProviderParameterResponseArrayOutput {
	return o
}

func (o ServiceProviderParameterResponseArrayOutput) Index(i pulumi.IntInput) ServiceProviderParameterResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceProviderParameterResponse {
		return vs[0].([]ServiceProviderParameterResponse)[vs[1].(int)]
	}).(ServiceProviderParameterResponseOutput)
}

type ServiceProviderPropertiesResponse struct {
	DevPortalUrl        string                             `pulumi:"devPortalUrl"`
	DisplayName         string                             `pulumi:"displayName"`
	IconUrl             string                             `pulumi:"iconUrl"`
	Id                  string                             `pulumi:"id"`
	Parameters          []ServiceProviderParameterResponse `pulumi:"parameters"`
	ServiceProviderName string                             `pulumi:"serviceProviderName"`
}





type ServiceProviderPropertiesResponseInput interface {
	pulumi.Input

	ToServiceProviderPropertiesResponseOutput() ServiceProviderPropertiesResponseOutput
	ToServiceProviderPropertiesResponseOutputWithContext(context.Context) ServiceProviderPropertiesResponseOutput
}

type ServiceProviderPropertiesResponseArgs struct {
	DevPortalUrl        pulumi.StringInput                         `pulumi:"devPortalUrl"`
	DisplayName         pulumi.StringInput                         `pulumi:"displayName"`
	IconUrl             pulumi.StringInput                         `pulumi:"iconUrl"`
	Id                  pulumi.StringInput                         `pulumi:"id"`
	Parameters          ServiceProviderParameterResponseArrayInput `pulumi:"parameters"`
	ServiceProviderName pulumi.StringInput                         `pulumi:"serviceProviderName"`
}

func (ServiceProviderPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderPropertiesResponse)(nil)).Elem()
}

func (i ServiceProviderPropertiesResponseArgs) ToServiceProviderPropertiesResponseOutput() ServiceProviderPropertiesResponseOutput {
	return i.ToServiceProviderPropertiesResponseOutputWithContext(context.Background())
}

func (i ServiceProviderPropertiesResponseArgs) ToServiceProviderPropertiesResponseOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderPropertiesResponseOutput)
}

func (i ServiceProviderPropertiesResponseArgs) ToServiceProviderPropertiesResponsePtrOutput() ServiceProviderPropertiesResponsePtrOutput {
	return i.ToServiceProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i ServiceProviderPropertiesResponseArgs) ToServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderPropertiesResponseOutput).ToServiceProviderPropertiesResponsePtrOutputWithContext(ctx)
}









type ServiceProviderPropertiesResponsePtrInput interface {
	pulumi.Input

	ToServiceProviderPropertiesResponsePtrOutput() ServiceProviderPropertiesResponsePtrOutput
	ToServiceProviderPropertiesResponsePtrOutputWithContext(context.Context) ServiceProviderPropertiesResponsePtrOutput
}

type serviceProviderPropertiesResponsePtrType ServiceProviderPropertiesResponseArgs

func ServiceProviderPropertiesResponsePtr(v *ServiceProviderPropertiesResponseArgs) ServiceProviderPropertiesResponsePtrInput {
	return (*serviceProviderPropertiesResponsePtrType)(v)
}

func (*serviceProviderPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceProviderPropertiesResponse)(nil)).Elem()
}

func (i *serviceProviderPropertiesResponsePtrType) ToServiceProviderPropertiesResponsePtrOutput() ServiceProviderPropertiesResponsePtrOutput {
	return i.ToServiceProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *serviceProviderPropertiesResponsePtrType) ToServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderPropertiesResponsePtrOutput)
}

type ServiceProviderPropertiesResponseOutput struct{ *pulumi.OutputState }

func (ServiceProviderPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ServiceProviderPropertiesResponseOutput) ToServiceProviderPropertiesResponseOutput() ServiceProviderPropertiesResponseOutput {
	return o
}

func (o ServiceProviderPropertiesResponseOutput) ToServiceProviderPropertiesResponseOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponseOutput {
	return o
}

func (o ServiceProviderPropertiesResponseOutput) ToServiceProviderPropertiesResponsePtrOutput() ServiceProviderPropertiesResponsePtrOutput {
	return o.ToServiceProviderPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o ServiceProviderPropertiesResponseOutput) ToServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v ServiceProviderPropertiesResponse) *ServiceProviderPropertiesResponse {
		return &v
	}).(ServiceProviderPropertiesResponsePtrOutput)
}

func (o ServiceProviderPropertiesResponseOutput) DevPortalUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.DevPortalUrl }).(pulumi.StringOutput)
}

func (o ServiceProviderPropertiesResponseOutput) DisplayName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.DisplayName }).(pulumi.StringOutput)
}

func (o ServiceProviderPropertiesResponseOutput) IconUrl() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.IconUrl }).(pulumi.StringOutput)
}

func (o ServiceProviderPropertiesResponseOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.Id }).(pulumi.StringOutput)
}

func (o ServiceProviderPropertiesResponseOutput) Parameters() ServiceProviderParameterResponseArrayOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) []ServiceProviderParameterResponse { return v.Parameters }).(ServiceProviderParameterResponseArrayOutput)
}

func (o ServiceProviderPropertiesResponseOutput) ServiceProviderName() pulumi.StringOutput {
	return o.ApplyT(func(v ServiceProviderPropertiesResponse) string { return v.ServiceProviderName }).(pulumi.StringOutput)
}

type ServiceProviderPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (ServiceProviderPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**ServiceProviderPropertiesResponse)(nil)).Elem()
}

func (o ServiceProviderPropertiesResponsePtrOutput) ToServiceProviderPropertiesResponsePtrOutput() ServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ServiceProviderPropertiesResponsePtrOutput) ToServiceProviderPropertiesResponsePtrOutputWithContext(ctx context.Context) ServiceProviderPropertiesResponsePtrOutput {
	return o
}

func (o ServiceProviderPropertiesResponsePtrOutput) Elem() ServiceProviderPropertiesResponseOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) ServiceProviderPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret ServiceProviderPropertiesResponse
		return ret
	}).(ServiceProviderPropertiesResponseOutput)
}

func (o ServiceProviderPropertiesResponsePtrOutput) DevPortalUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DevPortalUrl
	}).(pulumi.StringPtrOutput)
}

func (o ServiceProviderPropertiesResponsePtrOutput) DisplayName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.DisplayName
	}).(pulumi.StringPtrOutput)
}

func (o ServiceProviderPropertiesResponsePtrOutput) IconUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.IconUrl
	}).(pulumi.StringPtrOutput)
}

func (o ServiceProviderPropertiesResponsePtrOutput) Id() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Id
	}).(pulumi.StringPtrOutput)
}

func (o ServiceProviderPropertiesResponsePtrOutput) Parameters() ServiceProviderParameterResponseArrayOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) []ServiceProviderParameterResponse {
		if v == nil {
			return nil
		}
		return v.Parameters
	}).(ServiceProviderParameterResponseArrayOutput)
}

func (o ServiceProviderPropertiesResponsePtrOutput) ServiceProviderName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *ServiceProviderPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.ServiceProviderName
	}).(pulumi.StringPtrOutput)
}

type ServiceProviderResponse struct {
	Properties *ServiceProviderPropertiesResponse `pulumi:"properties"`
}





type ServiceProviderResponseInput interface {
	pulumi.Input

	ToServiceProviderResponseOutput() ServiceProviderResponseOutput
	ToServiceProviderResponseOutputWithContext(context.Context) ServiceProviderResponseOutput
}

type ServiceProviderResponseArgs struct {
	Properties ServiceProviderPropertiesResponsePtrInput `pulumi:"properties"`
}

func (ServiceProviderResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderResponse)(nil)).Elem()
}

func (i ServiceProviderResponseArgs) ToServiceProviderResponseOutput() ServiceProviderResponseOutput {
	return i.ToServiceProviderResponseOutputWithContext(context.Background())
}

func (i ServiceProviderResponseArgs) ToServiceProviderResponseOutputWithContext(ctx context.Context) ServiceProviderResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderResponseOutput)
}





type ServiceProviderResponseArrayInput interface {
	pulumi.Input

	ToServiceProviderResponseArrayOutput() ServiceProviderResponseArrayOutput
	ToServiceProviderResponseArrayOutputWithContext(context.Context) ServiceProviderResponseArrayOutput
}

type ServiceProviderResponseArray []ServiceProviderResponseInput

func (ServiceProviderResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceProviderResponse)(nil)).Elem()
}

func (i ServiceProviderResponseArray) ToServiceProviderResponseArrayOutput() ServiceProviderResponseArrayOutput {
	return i.ToServiceProviderResponseArrayOutputWithContext(context.Background())
}

func (i ServiceProviderResponseArray) ToServiceProviderResponseArrayOutputWithContext(ctx context.Context) ServiceProviderResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(ServiceProviderResponseArrayOutput)
}

type ServiceProviderResponseOutput struct{ *pulumi.OutputState }

func (ServiceProviderResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ServiceProviderResponse)(nil)).Elem()
}

func (o ServiceProviderResponseOutput) ToServiceProviderResponseOutput() ServiceProviderResponseOutput {
	return o
}

func (o ServiceProviderResponseOutput) ToServiceProviderResponseOutputWithContext(ctx context.Context) ServiceProviderResponseOutput {
	return o
}

func (o ServiceProviderResponseOutput) Properties() ServiceProviderPropertiesResponsePtrOutput {
	return o.ApplyT(func(v ServiceProviderResponse) *ServiceProviderPropertiesResponse { return v.Properties }).(ServiceProviderPropertiesResponsePtrOutput)
}

type ServiceProviderResponseArrayOutput struct{ *pulumi.OutputState }

func (ServiceProviderResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]ServiceProviderResponse)(nil)).Elem()
}

func (o ServiceProviderResponseArrayOutput) ToServiceProviderResponseArrayOutput() ServiceProviderResponseArrayOutput {
	return o
}

func (o ServiceProviderResponseArrayOutput) ToServiceProviderResponseArrayOutputWithContext(ctx context.Context) ServiceProviderResponseArrayOutput {
	return o
}

func (o ServiceProviderResponseArrayOutput) Index(i pulumi.IntInput) ServiceProviderResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) ServiceProviderResponse {
		return vs[0].([]ServiceProviderResponse)[vs[1].(int)]
	}).(ServiceProviderResponseOutput)
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





type SkuResponseInput interface {
	pulumi.Input

	ToSkuResponseOutput() SkuResponseOutput
	ToSkuResponseOutputWithContext(context.Context) SkuResponseOutput
}

type SkuResponseArgs struct {
	Name pulumi.StringInput `pulumi:"name"`
	Tier pulumi.StringInput `pulumi:"tier"`
}

func (SkuResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkuResponse)(nil)).Elem()
}

func (i SkuResponseArgs) ToSkuResponseOutput() SkuResponseOutput {
	return i.ToSkuResponseOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponseOutputWithContext(ctx context.Context) SkuResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput)
}

func (i SkuResponseArgs) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i SkuResponseArgs) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponseOutput).ToSkuResponsePtrOutputWithContext(ctx)
}









type SkuResponsePtrInput interface {
	pulumi.Input

	ToSkuResponsePtrOutput() SkuResponsePtrOutput
	ToSkuResponsePtrOutputWithContext(context.Context) SkuResponsePtrOutput
}

type skuResponsePtrType SkuResponseArgs

func SkuResponsePtr(v *SkuResponseArgs) SkuResponsePtrInput {
	return (*skuResponsePtrType)(v)
}

func (*skuResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkuResponse)(nil)).Elem()
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return i.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (i *skuResponsePtrType) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkuResponsePtrOutput)
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

func (o SkuResponseOutput) ToSkuResponsePtrOutput() SkuResponsePtrOutput {
	return o.ToSkuResponsePtrOutputWithContext(context.Background())
}

func (o SkuResponseOutput) ToSkuResponsePtrOutputWithContext(ctx context.Context) SkuResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkuResponse) *SkuResponse {
		return &v
	}).(SkuResponsePtrOutput)
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
	Properties  *SkypeChannelProperties `pulumi:"properties"`
}





type SkypeChannelInput interface {
	pulumi.Input

	ToSkypeChannelOutput() SkypeChannelOutput
	ToSkypeChannelOutputWithContext(context.Context) SkypeChannelOutput
}

type SkypeChannelArgs struct {
	ChannelName pulumi.StringInput             `pulumi:"channelName"`
	Properties  SkypeChannelPropertiesPtrInput `pulumi:"properties"`
}

func (SkypeChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannel)(nil)).Elem()
}

func (i SkypeChannelArgs) ToSkypeChannelOutput() SkypeChannelOutput {
	return i.ToSkypeChannelOutputWithContext(context.Background())
}

func (i SkypeChannelArgs) ToSkypeChannelOutputWithContext(ctx context.Context) SkypeChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelOutput)
}

type SkypeChannelOutput struct{ *pulumi.OutputState }

func (SkypeChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannel)(nil)).Elem()
}

func (o SkypeChannelOutput) ToSkypeChannelOutput() SkypeChannelOutput {
	return o
}

func (o SkypeChannelOutput) ToSkypeChannelOutputWithContext(ctx context.Context) SkypeChannelOutput {
	return o
}

func (o SkypeChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v SkypeChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o SkypeChannelOutput) Properties() SkypeChannelPropertiesPtrOutput {
	return o.ApplyT(func(v SkypeChannel) *SkypeChannelProperties { return v.Properties }).(SkypeChannelPropertiesPtrOutput)
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
	IsEnabled           bool    `pulumi:"isEnabled"`
}





type SkypeChannelPropertiesInput interface {
	pulumi.Input

	ToSkypeChannelPropertiesOutput() SkypeChannelPropertiesOutput
	ToSkypeChannelPropertiesOutputWithContext(context.Context) SkypeChannelPropertiesOutput
}

type SkypeChannelPropertiesArgs struct {
	CallingWebHook      pulumi.StringPtrInput `pulumi:"callingWebHook"`
	EnableCalling       pulumi.BoolPtrInput   `pulumi:"enableCalling"`
	EnableGroups        pulumi.BoolPtrInput   `pulumi:"enableGroups"`
	EnableMediaCards    pulumi.BoolPtrInput   `pulumi:"enableMediaCards"`
	EnableMessaging     pulumi.BoolPtrInput   `pulumi:"enableMessaging"`
	EnableScreenSharing pulumi.BoolPtrInput   `pulumi:"enableScreenSharing"`
	EnableVideo         pulumi.BoolPtrInput   `pulumi:"enableVideo"`
	GroupsMode          pulumi.StringPtrInput `pulumi:"groupsMode"`
	IsEnabled           pulumi.BoolInput      `pulumi:"isEnabled"`
}

func (SkypeChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannelProperties)(nil)).Elem()
}

func (i SkypeChannelPropertiesArgs) ToSkypeChannelPropertiesOutput() SkypeChannelPropertiesOutput {
	return i.ToSkypeChannelPropertiesOutputWithContext(context.Background())
}

func (i SkypeChannelPropertiesArgs) ToSkypeChannelPropertiesOutputWithContext(ctx context.Context) SkypeChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelPropertiesOutput)
}

func (i SkypeChannelPropertiesArgs) ToSkypeChannelPropertiesPtrOutput() SkypeChannelPropertiesPtrOutput {
	return i.ToSkypeChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i SkypeChannelPropertiesArgs) ToSkypeChannelPropertiesPtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelPropertiesOutput).ToSkypeChannelPropertiesPtrOutputWithContext(ctx)
}









type SkypeChannelPropertiesPtrInput interface {
	pulumi.Input

	ToSkypeChannelPropertiesPtrOutput() SkypeChannelPropertiesPtrOutput
	ToSkypeChannelPropertiesPtrOutputWithContext(context.Context) SkypeChannelPropertiesPtrOutput
}

type skypeChannelPropertiesPtrType SkypeChannelPropertiesArgs

func SkypeChannelPropertiesPtr(v *SkypeChannelPropertiesArgs) SkypeChannelPropertiesPtrInput {
	return (*skypeChannelPropertiesPtrType)(v)
}

func (*skypeChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkypeChannelProperties)(nil)).Elem()
}

func (i *skypeChannelPropertiesPtrType) ToSkypeChannelPropertiesPtrOutput() SkypeChannelPropertiesPtrOutput {
	return i.ToSkypeChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *skypeChannelPropertiesPtrType) ToSkypeChannelPropertiesPtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelPropertiesPtrOutput)
}

type SkypeChannelPropertiesOutput struct{ *pulumi.OutputState }

func (SkypeChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannelProperties)(nil)).Elem()
}

func (o SkypeChannelPropertiesOutput) ToSkypeChannelPropertiesOutput() SkypeChannelPropertiesOutput {
	return o
}

func (o SkypeChannelPropertiesOutput) ToSkypeChannelPropertiesOutputWithContext(ctx context.Context) SkypeChannelPropertiesOutput {
	return o
}

func (o SkypeChannelPropertiesOutput) ToSkypeChannelPropertiesPtrOutput() SkypeChannelPropertiesPtrOutput {
	return o.ToSkypeChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o SkypeChannelPropertiesOutput) ToSkypeChannelPropertiesPtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkypeChannelProperties) *SkypeChannelProperties {
		return &v
	}).(SkypeChannelPropertiesPtrOutput)
}

func (o SkypeChannelPropertiesOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *string { return v.CallingWebHook }).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *bool { return v.EnableCalling }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesOutput) EnableGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *bool { return v.EnableGroups }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesOutput) EnableMediaCards() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *bool { return v.EnableMediaCards }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesOutput) EnableMessaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *bool { return v.EnableMessaging }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesOutput) EnableScreenSharing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *bool { return v.EnableScreenSharing }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesOutput) EnableVideo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *bool { return v.EnableVideo }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesOutput) GroupsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkypeChannelProperties) *string { return v.GroupsMode }).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SkypeChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

type SkypeChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SkypeChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkypeChannelProperties)(nil)).Elem()
}

func (o SkypeChannelPropertiesPtrOutput) ToSkypeChannelPropertiesPtrOutput() SkypeChannelPropertiesPtrOutput {
	return o
}

func (o SkypeChannelPropertiesPtrOutput) ToSkypeChannelPropertiesPtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesPtrOutput {
	return o
}

func (o SkypeChannelPropertiesPtrOutput) Elem() SkypeChannelPropertiesOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) SkypeChannelProperties {
		if v != nil {
			return *v
		}
		var ret SkypeChannelProperties
		return ret
	}).(SkypeChannelPropertiesOutput)
}

func (o SkypeChannelPropertiesPtrOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.CallingWebHook
	}).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCalling
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) EnableGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableGroups
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) EnableMediaCards() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableMediaCards
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) EnableMessaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableMessaging
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) EnableScreenSharing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableScreenSharing
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) EnableVideo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.EnableVideo
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) GroupsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.GroupsMode
	}).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
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
	IsEnabled           bool    `pulumi:"isEnabled"`
}





type SkypeChannelPropertiesResponseInput interface {
	pulumi.Input

	ToSkypeChannelPropertiesResponseOutput() SkypeChannelPropertiesResponseOutput
	ToSkypeChannelPropertiesResponseOutputWithContext(context.Context) SkypeChannelPropertiesResponseOutput
}

type SkypeChannelPropertiesResponseArgs struct {
	CallingWebHook      pulumi.StringPtrInput `pulumi:"callingWebHook"`
	EnableCalling       pulumi.BoolPtrInput   `pulumi:"enableCalling"`
	EnableGroups        pulumi.BoolPtrInput   `pulumi:"enableGroups"`
	EnableMediaCards    pulumi.BoolPtrInput   `pulumi:"enableMediaCards"`
	EnableMessaging     pulumi.BoolPtrInput   `pulumi:"enableMessaging"`
	EnableScreenSharing pulumi.BoolPtrInput   `pulumi:"enableScreenSharing"`
	EnableVideo         pulumi.BoolPtrInput   `pulumi:"enableVideo"`
	GroupsMode          pulumi.StringPtrInput `pulumi:"groupsMode"`
	IsEnabled           pulumi.BoolInput      `pulumi:"isEnabled"`
}

func (SkypeChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannelPropertiesResponse)(nil)).Elem()
}

func (i SkypeChannelPropertiesResponseArgs) ToSkypeChannelPropertiesResponseOutput() SkypeChannelPropertiesResponseOutput {
	return i.ToSkypeChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i SkypeChannelPropertiesResponseArgs) ToSkypeChannelPropertiesResponseOutputWithContext(ctx context.Context) SkypeChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelPropertiesResponseOutput)
}

func (i SkypeChannelPropertiesResponseArgs) ToSkypeChannelPropertiesResponsePtrOutput() SkypeChannelPropertiesResponsePtrOutput {
	return i.ToSkypeChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SkypeChannelPropertiesResponseArgs) ToSkypeChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelPropertiesResponseOutput).ToSkypeChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type SkypeChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSkypeChannelPropertiesResponsePtrOutput() SkypeChannelPropertiesResponsePtrOutput
	ToSkypeChannelPropertiesResponsePtrOutputWithContext(context.Context) SkypeChannelPropertiesResponsePtrOutput
}

type skypeChannelPropertiesResponsePtrType SkypeChannelPropertiesResponseArgs

func SkypeChannelPropertiesResponsePtr(v *SkypeChannelPropertiesResponseArgs) SkypeChannelPropertiesResponsePtrInput {
	return (*skypeChannelPropertiesResponsePtrType)(v)
}

func (*skypeChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SkypeChannelPropertiesResponse)(nil)).Elem()
}

func (i *skypeChannelPropertiesResponsePtrType) ToSkypeChannelPropertiesResponsePtrOutput() SkypeChannelPropertiesResponsePtrOutput {
	return i.ToSkypeChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *skypeChannelPropertiesResponsePtrType) ToSkypeChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelPropertiesResponsePtrOutput)
}

type SkypeChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SkypeChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannelPropertiesResponse)(nil)).Elem()
}

func (o SkypeChannelPropertiesResponseOutput) ToSkypeChannelPropertiesResponseOutput() SkypeChannelPropertiesResponseOutput {
	return o
}

func (o SkypeChannelPropertiesResponseOutput) ToSkypeChannelPropertiesResponseOutputWithContext(ctx context.Context) SkypeChannelPropertiesResponseOutput {
	return o
}

func (o SkypeChannelPropertiesResponseOutput) ToSkypeChannelPropertiesResponsePtrOutput() SkypeChannelPropertiesResponsePtrOutput {
	return o.ToSkypeChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SkypeChannelPropertiesResponseOutput) ToSkypeChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SkypeChannelPropertiesResponse) *SkypeChannelPropertiesResponse {
		return &v
	}).(SkypeChannelPropertiesResponsePtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *string { return v.CallingWebHook }).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *bool { return v.EnableCalling }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) EnableGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *bool { return v.EnableGroups }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) EnableMediaCards() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *bool { return v.EnableMediaCards }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) EnableMessaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *bool { return v.EnableMessaging }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) EnableScreenSharing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *bool { return v.EnableScreenSharing }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) EnableVideo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *bool { return v.EnableVideo }).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) GroupsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) *string { return v.GroupsMode }).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SkypeChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

type SkypeChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SkypeChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SkypeChannelPropertiesResponse)(nil)).Elem()
}

func (o SkypeChannelPropertiesResponsePtrOutput) ToSkypeChannelPropertiesResponsePtrOutput() SkypeChannelPropertiesResponsePtrOutput {
	return o
}

func (o SkypeChannelPropertiesResponsePtrOutput) ToSkypeChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SkypeChannelPropertiesResponsePtrOutput {
	return o
}

func (o SkypeChannelPropertiesResponsePtrOutput) Elem() SkypeChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) SkypeChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SkypeChannelPropertiesResponse
		return ret
	}).(SkypeChannelPropertiesResponseOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) CallingWebHook() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.CallingWebHook
	}).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) EnableCalling() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableCalling
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) EnableGroups() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableGroups
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) EnableMediaCards() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableMediaCards
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) EnableMessaging() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableMessaging
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) EnableScreenSharing() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableScreenSharing
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) EnableVideo() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.EnableVideo
	}).(pulumi.BoolPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) GroupsMode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.GroupsMode
	}).(pulumi.StringPtrOutput)
}

func (o SkypeChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SkypeChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

type SkypeChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *SkypeChannelPropertiesResponse `pulumi:"properties"`
}





type SkypeChannelResponseInput interface {
	pulumi.Input

	ToSkypeChannelResponseOutput() SkypeChannelResponseOutput
	ToSkypeChannelResponseOutputWithContext(context.Context) SkypeChannelResponseOutput
}

type SkypeChannelResponseArgs struct {
	ChannelName pulumi.StringInput                     `pulumi:"channelName"`
	Properties  SkypeChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (SkypeChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannelResponse)(nil)).Elem()
}

func (i SkypeChannelResponseArgs) ToSkypeChannelResponseOutput() SkypeChannelResponseOutput {
	return i.ToSkypeChannelResponseOutputWithContext(context.Background())
}

func (i SkypeChannelResponseArgs) ToSkypeChannelResponseOutputWithContext(ctx context.Context) SkypeChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SkypeChannelResponseOutput)
}

type SkypeChannelResponseOutput struct{ *pulumi.OutputState }

func (SkypeChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SkypeChannelResponse)(nil)).Elem()
}

func (o SkypeChannelResponseOutput) ToSkypeChannelResponseOutput() SkypeChannelResponseOutput {
	return o
}

func (o SkypeChannelResponseOutput) ToSkypeChannelResponseOutputWithContext(ctx context.Context) SkypeChannelResponseOutput {
	return o
}

func (o SkypeChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v SkypeChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o SkypeChannelResponseOutput) Properties() SkypeChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SkypeChannelResponse) *SkypeChannelPropertiesResponse { return v.Properties }).(SkypeChannelPropertiesResponsePtrOutput)
}

type SlackChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Properties  *SlackChannelProperties `pulumi:"properties"`
}





type SlackChannelInput interface {
	pulumi.Input

	ToSlackChannelOutput() SlackChannelOutput
	ToSlackChannelOutputWithContext(context.Context) SlackChannelOutput
}

type SlackChannelArgs struct {
	ChannelName pulumi.StringInput             `pulumi:"channelName"`
	Properties  SlackChannelPropertiesPtrInput `pulumi:"properties"`
}

func (SlackChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannel)(nil)).Elem()
}

func (i SlackChannelArgs) ToSlackChannelOutput() SlackChannelOutput {
	return i.ToSlackChannelOutputWithContext(context.Background())
}

func (i SlackChannelArgs) ToSlackChannelOutputWithContext(ctx context.Context) SlackChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelOutput)
}

type SlackChannelOutput struct{ *pulumi.OutputState }

func (SlackChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannel)(nil)).Elem()
}

func (o SlackChannelOutput) ToSlackChannelOutput() SlackChannelOutput {
	return o
}

func (o SlackChannelOutput) ToSlackChannelOutputWithContext(ctx context.Context) SlackChannelOutput {
	return o
}

func (o SlackChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v SlackChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o SlackChannelOutput) Properties() SlackChannelPropertiesPtrOutput {
	return o.ApplyT(func(v SlackChannel) *SlackChannelProperties { return v.Properties }).(SlackChannelPropertiesPtrOutput)
}

type SlackChannelProperties struct {
	ClientId          *string `pulumi:"clientId"`
	ClientSecret      *string `pulumi:"clientSecret"`
	IsEnabled         bool    `pulumi:"isEnabled"`
	LandingPageUrl    *string `pulumi:"landingPageUrl"`
	SigningSecret     *string `pulumi:"signingSecret"`
	VerificationToken *string `pulumi:"verificationToken"`
}





type SlackChannelPropertiesInput interface {
	pulumi.Input

	ToSlackChannelPropertiesOutput() SlackChannelPropertiesOutput
	ToSlackChannelPropertiesOutputWithContext(context.Context) SlackChannelPropertiesOutput
}

type SlackChannelPropertiesArgs struct {
	ClientId          pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret      pulumi.StringPtrInput `pulumi:"clientSecret"`
	IsEnabled         pulumi.BoolInput      `pulumi:"isEnabled"`
	LandingPageUrl    pulumi.StringPtrInput `pulumi:"landingPageUrl"`
	SigningSecret     pulumi.StringPtrInput `pulumi:"signingSecret"`
	VerificationToken pulumi.StringPtrInput `pulumi:"verificationToken"`
}

func (SlackChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelProperties)(nil)).Elem()
}

func (i SlackChannelPropertiesArgs) ToSlackChannelPropertiesOutput() SlackChannelPropertiesOutput {
	return i.ToSlackChannelPropertiesOutputWithContext(context.Background())
}

func (i SlackChannelPropertiesArgs) ToSlackChannelPropertiesOutputWithContext(ctx context.Context) SlackChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelPropertiesOutput)
}

func (i SlackChannelPropertiesArgs) ToSlackChannelPropertiesPtrOutput() SlackChannelPropertiesPtrOutput {
	return i.ToSlackChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i SlackChannelPropertiesArgs) ToSlackChannelPropertiesPtrOutputWithContext(ctx context.Context) SlackChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelPropertiesOutput).ToSlackChannelPropertiesPtrOutputWithContext(ctx)
}









type SlackChannelPropertiesPtrInput interface {
	pulumi.Input

	ToSlackChannelPropertiesPtrOutput() SlackChannelPropertiesPtrOutput
	ToSlackChannelPropertiesPtrOutputWithContext(context.Context) SlackChannelPropertiesPtrOutput
}

type slackChannelPropertiesPtrType SlackChannelPropertiesArgs

func SlackChannelPropertiesPtr(v *SlackChannelPropertiesArgs) SlackChannelPropertiesPtrInput {
	return (*slackChannelPropertiesPtrType)(v)
}

func (*slackChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackChannelProperties)(nil)).Elem()
}

func (i *slackChannelPropertiesPtrType) ToSlackChannelPropertiesPtrOutput() SlackChannelPropertiesPtrOutput {
	return i.ToSlackChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *slackChannelPropertiesPtrType) ToSlackChannelPropertiesPtrOutputWithContext(ctx context.Context) SlackChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelPropertiesPtrOutput)
}

type SlackChannelPropertiesOutput struct{ *pulumi.OutputState }

func (SlackChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelProperties)(nil)).Elem()
}

func (o SlackChannelPropertiesOutput) ToSlackChannelPropertiesOutput() SlackChannelPropertiesOutput {
	return o
}

func (o SlackChannelPropertiesOutput) ToSlackChannelPropertiesOutputWithContext(ctx context.Context) SlackChannelPropertiesOutput {
	return o
}

func (o SlackChannelPropertiesOutput) ToSlackChannelPropertiesPtrOutput() SlackChannelPropertiesPtrOutput {
	return o.ToSlackChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o SlackChannelPropertiesOutput) ToSlackChannelPropertiesPtrOutputWithContext(ctx context.Context) SlackChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlackChannelProperties) *SlackChannelProperties {
		return &v
	}).(SlackChannelPropertiesPtrOutput)
}

func (o SlackChannelPropertiesOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelProperties) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelProperties) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SlackChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o SlackChannelPropertiesOutput) LandingPageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelProperties) *string { return v.LandingPageUrl }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesOutput) SigningSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelProperties) *string { return v.SigningSecret }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesOutput) VerificationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelProperties) *string { return v.VerificationToken }).(pulumi.StringPtrOutput)
}

type SlackChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SlackChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackChannelProperties)(nil)).Elem()
}

func (o SlackChannelPropertiesPtrOutput) ToSlackChannelPropertiesPtrOutput() SlackChannelPropertiesPtrOutput {
	return o
}

func (o SlackChannelPropertiesPtrOutput) ToSlackChannelPropertiesPtrOutputWithContext(ctx context.Context) SlackChannelPropertiesPtrOutput {
	return o
}

func (o SlackChannelPropertiesPtrOutput) Elem() SlackChannelPropertiesOutput {
	return o.ApplyT(func(v *SlackChannelProperties) SlackChannelProperties {
		if v != nil {
			return *v
		}
		var ret SlackChannelProperties
		return ret
	}).(SlackChannelPropertiesOutput)
}

func (o SlackChannelPropertiesPtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesPtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SlackChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SlackChannelPropertiesPtrOutput) LandingPageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.LandingPageUrl
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesPtrOutput) SigningSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.SigningSecret
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesPtrOutput) VerificationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.VerificationToken
	}).(pulumi.StringPtrOutput)
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





type SlackChannelPropertiesResponseInput interface {
	pulumi.Input

	ToSlackChannelPropertiesResponseOutput() SlackChannelPropertiesResponseOutput
	ToSlackChannelPropertiesResponseOutputWithContext(context.Context) SlackChannelPropertiesResponseOutput
}

type SlackChannelPropertiesResponseArgs struct {
	ClientId                pulumi.StringPtrInput `pulumi:"clientId"`
	ClientSecret            pulumi.StringPtrInput `pulumi:"clientSecret"`
	IsEnabled               pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated             pulumi.BoolInput      `pulumi:"isValidated"`
	LandingPageUrl          pulumi.StringPtrInput `pulumi:"landingPageUrl"`
	LastSubmissionId        pulumi.StringInput    `pulumi:"lastSubmissionId"`
	RedirectAction          pulumi.StringInput    `pulumi:"redirectAction"`
	RegisterBeforeOAuthFlow pulumi.BoolInput      `pulumi:"registerBeforeOAuthFlow"`
	SigningSecret           pulumi.StringPtrInput `pulumi:"signingSecret"`
	VerificationToken       pulumi.StringPtrInput `pulumi:"verificationToken"`
}

func (SlackChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelPropertiesResponse)(nil)).Elem()
}

func (i SlackChannelPropertiesResponseArgs) ToSlackChannelPropertiesResponseOutput() SlackChannelPropertiesResponseOutput {
	return i.ToSlackChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i SlackChannelPropertiesResponseArgs) ToSlackChannelPropertiesResponseOutputWithContext(ctx context.Context) SlackChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelPropertiesResponseOutput)
}

func (i SlackChannelPropertiesResponseArgs) ToSlackChannelPropertiesResponsePtrOutput() SlackChannelPropertiesResponsePtrOutput {
	return i.ToSlackChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SlackChannelPropertiesResponseArgs) ToSlackChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SlackChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelPropertiesResponseOutput).ToSlackChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type SlackChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSlackChannelPropertiesResponsePtrOutput() SlackChannelPropertiesResponsePtrOutput
	ToSlackChannelPropertiesResponsePtrOutputWithContext(context.Context) SlackChannelPropertiesResponsePtrOutput
}

type slackChannelPropertiesResponsePtrType SlackChannelPropertiesResponseArgs

func SlackChannelPropertiesResponsePtr(v *SlackChannelPropertiesResponseArgs) SlackChannelPropertiesResponsePtrInput {
	return (*slackChannelPropertiesResponsePtrType)(v)
}

func (*slackChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackChannelPropertiesResponse)(nil)).Elem()
}

func (i *slackChannelPropertiesResponsePtrType) ToSlackChannelPropertiesResponsePtrOutput() SlackChannelPropertiesResponsePtrOutput {
	return i.ToSlackChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *slackChannelPropertiesResponsePtrType) ToSlackChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SlackChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelPropertiesResponsePtrOutput)
}

type SlackChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SlackChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelPropertiesResponse)(nil)).Elem()
}

func (o SlackChannelPropertiesResponseOutput) ToSlackChannelPropertiesResponseOutput() SlackChannelPropertiesResponseOutput {
	return o
}

func (o SlackChannelPropertiesResponseOutput) ToSlackChannelPropertiesResponseOutputWithContext(ctx context.Context) SlackChannelPropertiesResponseOutput {
	return o
}

func (o SlackChannelPropertiesResponseOutput) ToSlackChannelPropertiesResponsePtrOutput() SlackChannelPropertiesResponsePtrOutput {
	return o.ToSlackChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SlackChannelPropertiesResponseOutput) ToSlackChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SlackChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SlackChannelPropertiesResponse) *SlackChannelPropertiesResponse {
		return &v
	}).(SlackChannelPropertiesResponsePtrOutput)
}

func (o SlackChannelPropertiesResponseOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponseOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o SlackChannelPropertiesResponseOutput) IsValidated() pulumi.BoolOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) bool { return v.IsValidated }).(pulumi.BoolOutput)
}

func (o SlackChannelPropertiesResponseOutput) LandingPageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) *string { return v.LandingPageUrl }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponseOutput) LastSubmissionId() pulumi.StringOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) string { return v.LastSubmissionId }).(pulumi.StringOutput)
}

func (o SlackChannelPropertiesResponseOutput) RedirectAction() pulumi.StringOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) string { return v.RedirectAction }).(pulumi.StringOutput)
}

func (o SlackChannelPropertiesResponseOutput) RegisterBeforeOAuthFlow() pulumi.BoolOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) bool { return v.RegisterBeforeOAuthFlow }).(pulumi.BoolOutput)
}

func (o SlackChannelPropertiesResponseOutput) SigningSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) *string { return v.SigningSecret }).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponseOutput) VerificationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SlackChannelPropertiesResponse) *string { return v.VerificationToken }).(pulumi.StringPtrOutput)
}

type SlackChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SlackChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SlackChannelPropertiesResponse)(nil)).Elem()
}

func (o SlackChannelPropertiesResponsePtrOutput) ToSlackChannelPropertiesResponsePtrOutput() SlackChannelPropertiesResponsePtrOutput {
	return o
}

func (o SlackChannelPropertiesResponsePtrOutput) ToSlackChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SlackChannelPropertiesResponsePtrOutput {
	return o
}

func (o SlackChannelPropertiesResponsePtrOutput) Elem() SlackChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) SlackChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SlackChannelPropertiesResponse
		return ret
	}).(SlackChannelPropertiesResponseOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientId
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.ClientSecret
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) LandingPageUrl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.LandingPageUrl
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) LastSubmissionId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.LastSubmissionId
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) RedirectAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.RedirectAction
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) RegisterBeforeOAuthFlow() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.RegisterBeforeOAuthFlow
	}).(pulumi.BoolPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) SigningSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.SigningSecret
	}).(pulumi.StringPtrOutput)
}

func (o SlackChannelPropertiesResponsePtrOutput) VerificationToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SlackChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.VerificationToken
	}).(pulumi.StringPtrOutput)
}

type SlackChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *SlackChannelPropertiesResponse `pulumi:"properties"`
}





type SlackChannelResponseInput interface {
	pulumi.Input

	ToSlackChannelResponseOutput() SlackChannelResponseOutput
	ToSlackChannelResponseOutputWithContext(context.Context) SlackChannelResponseOutput
}

type SlackChannelResponseArgs struct {
	ChannelName pulumi.StringInput                     `pulumi:"channelName"`
	Properties  SlackChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (SlackChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelResponse)(nil)).Elem()
}

func (i SlackChannelResponseArgs) ToSlackChannelResponseOutput() SlackChannelResponseOutput {
	return i.ToSlackChannelResponseOutputWithContext(context.Background())
}

func (i SlackChannelResponseArgs) ToSlackChannelResponseOutputWithContext(ctx context.Context) SlackChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SlackChannelResponseOutput)
}

type SlackChannelResponseOutput struct{ *pulumi.OutputState }

func (SlackChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SlackChannelResponse)(nil)).Elem()
}

func (o SlackChannelResponseOutput) ToSlackChannelResponseOutput() SlackChannelResponseOutput {
	return o
}

func (o SlackChannelResponseOutput) ToSlackChannelResponseOutputWithContext(ctx context.Context) SlackChannelResponseOutput {
	return o
}

func (o SlackChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v SlackChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o SlackChannelResponseOutput) Properties() SlackChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SlackChannelResponse) *SlackChannelPropertiesResponse { return v.Properties }).(SlackChannelPropertiesResponsePtrOutput)
}

type SmsChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Properties  *SmsChannelProperties `pulumi:"properties"`
}





type SmsChannelInput interface {
	pulumi.Input

	ToSmsChannelOutput() SmsChannelOutput
	ToSmsChannelOutputWithContext(context.Context) SmsChannelOutput
}

type SmsChannelArgs struct {
	ChannelName pulumi.StringInput           `pulumi:"channelName"`
	Properties  SmsChannelPropertiesPtrInput `pulumi:"properties"`
}

func (SmsChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannel)(nil)).Elem()
}

func (i SmsChannelArgs) ToSmsChannelOutput() SmsChannelOutput {
	return i.ToSmsChannelOutputWithContext(context.Background())
}

func (i SmsChannelArgs) ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelOutput)
}

type SmsChannelOutput struct{ *pulumi.OutputState }

func (SmsChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannel)(nil)).Elem()
}

func (o SmsChannelOutput) ToSmsChannelOutput() SmsChannelOutput {
	return o
}

func (o SmsChannelOutput) ToSmsChannelOutputWithContext(ctx context.Context) SmsChannelOutput {
	return o
}

func (o SmsChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v SmsChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o SmsChannelOutput) Properties() SmsChannelPropertiesPtrOutput {
	return o.ApplyT(func(v SmsChannel) *SmsChannelProperties { return v.Properties }).(SmsChannelPropertiesPtrOutput)
}

type SmsChannelProperties struct {
	AccountSID  string  `pulumi:"accountSID"`
	AuthToken   *string `pulumi:"authToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	Phone       string  `pulumi:"phone"`
}





type SmsChannelPropertiesInput interface {
	pulumi.Input

	ToSmsChannelPropertiesOutput() SmsChannelPropertiesOutput
	ToSmsChannelPropertiesOutputWithContext(context.Context) SmsChannelPropertiesOutput
}

type SmsChannelPropertiesArgs struct {
	AccountSID  pulumi.StringInput    `pulumi:"accountSID"`
	AuthToken   pulumi.StringPtrInput `pulumi:"authToken"`
	IsEnabled   pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated pulumi.BoolPtrInput   `pulumi:"isValidated"`
	Phone       pulumi.StringInput    `pulumi:"phone"`
}

func (SmsChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannelProperties)(nil)).Elem()
}

func (i SmsChannelPropertiesArgs) ToSmsChannelPropertiesOutput() SmsChannelPropertiesOutput {
	return i.ToSmsChannelPropertiesOutputWithContext(context.Background())
}

func (i SmsChannelPropertiesArgs) ToSmsChannelPropertiesOutputWithContext(ctx context.Context) SmsChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPropertiesOutput)
}

func (i SmsChannelPropertiesArgs) ToSmsChannelPropertiesPtrOutput() SmsChannelPropertiesPtrOutput {
	return i.ToSmsChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i SmsChannelPropertiesArgs) ToSmsChannelPropertiesPtrOutputWithContext(ctx context.Context) SmsChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPropertiesOutput).ToSmsChannelPropertiesPtrOutputWithContext(ctx)
}









type SmsChannelPropertiesPtrInput interface {
	pulumi.Input

	ToSmsChannelPropertiesPtrOutput() SmsChannelPropertiesPtrOutput
	ToSmsChannelPropertiesPtrOutputWithContext(context.Context) SmsChannelPropertiesPtrOutput
}

type smsChannelPropertiesPtrType SmsChannelPropertiesArgs

func SmsChannelPropertiesPtr(v *SmsChannelPropertiesArgs) SmsChannelPropertiesPtrInput {
	return (*smsChannelPropertiesPtrType)(v)
}

func (*smsChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannelProperties)(nil)).Elem()
}

func (i *smsChannelPropertiesPtrType) ToSmsChannelPropertiesPtrOutput() SmsChannelPropertiesPtrOutput {
	return i.ToSmsChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *smsChannelPropertiesPtrType) ToSmsChannelPropertiesPtrOutputWithContext(ctx context.Context) SmsChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPropertiesPtrOutput)
}

type SmsChannelPropertiesOutput struct{ *pulumi.OutputState }

func (SmsChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannelProperties)(nil)).Elem()
}

func (o SmsChannelPropertiesOutput) ToSmsChannelPropertiesOutput() SmsChannelPropertiesOutput {
	return o
}

func (o SmsChannelPropertiesOutput) ToSmsChannelPropertiesOutputWithContext(ctx context.Context) SmsChannelPropertiesOutput {
	return o
}

func (o SmsChannelPropertiesOutput) ToSmsChannelPropertiesPtrOutput() SmsChannelPropertiesPtrOutput {
	return o.ToSmsChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o SmsChannelPropertiesOutput) ToSmsChannelPropertiesPtrOutputWithContext(ctx context.Context) SmsChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SmsChannelProperties) *SmsChannelProperties {
		return &v
	}).(SmsChannelPropertiesPtrOutput)
}

func (o SmsChannelPropertiesOutput) AccountSID() pulumi.StringOutput {
	return o.ApplyT(func(v SmsChannelProperties) string { return v.AccountSID }).(pulumi.StringOutput)
}

func (o SmsChannelPropertiesOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SmsChannelProperties) *string { return v.AuthToken }).(pulumi.StringPtrOutput)
}

func (o SmsChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SmsChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o SmsChannelPropertiesOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SmsChannelProperties) *bool { return v.IsValidated }).(pulumi.BoolPtrOutput)
}

func (o SmsChannelPropertiesOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v SmsChannelProperties) string { return v.Phone }).(pulumi.StringOutput)
}

type SmsChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (SmsChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannelProperties)(nil)).Elem()
}

func (o SmsChannelPropertiesPtrOutput) ToSmsChannelPropertiesPtrOutput() SmsChannelPropertiesPtrOutput {
	return o
}

func (o SmsChannelPropertiesPtrOutput) ToSmsChannelPropertiesPtrOutputWithContext(ctx context.Context) SmsChannelPropertiesPtrOutput {
	return o
}

func (o SmsChannelPropertiesPtrOutput) Elem() SmsChannelPropertiesOutput {
	return o.ApplyT(func(v *SmsChannelProperties) SmsChannelProperties {
		if v != nil {
			return *v
		}
		var ret SmsChannelProperties
		return ret
	}).(SmsChannelPropertiesOutput)
}

func (o SmsChannelPropertiesPtrOutput) AccountSID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.AccountSID
	}).(pulumi.StringPtrOutput)
}

func (o SmsChannelPropertiesPtrOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.AuthToken
	}).(pulumi.StringPtrOutput)
}

func (o SmsChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SmsChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SmsChannelPropertiesPtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SmsChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

func (o SmsChannelPropertiesPtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannelProperties) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

type SmsChannelPropertiesResponse struct {
	AccountSID  string  `pulumi:"accountSID"`
	AuthToken   *string `pulumi:"authToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
	Phone       string  `pulumi:"phone"`
}





type SmsChannelPropertiesResponseInput interface {
	pulumi.Input

	ToSmsChannelPropertiesResponseOutput() SmsChannelPropertiesResponseOutput
	ToSmsChannelPropertiesResponseOutputWithContext(context.Context) SmsChannelPropertiesResponseOutput
}

type SmsChannelPropertiesResponseArgs struct {
	AccountSID  pulumi.StringInput    `pulumi:"accountSID"`
	AuthToken   pulumi.StringPtrInput `pulumi:"authToken"`
	IsEnabled   pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated pulumi.BoolPtrInput   `pulumi:"isValidated"`
	Phone       pulumi.StringInput    `pulumi:"phone"`
}

func (SmsChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannelPropertiesResponse)(nil)).Elem()
}

func (i SmsChannelPropertiesResponseArgs) ToSmsChannelPropertiesResponseOutput() SmsChannelPropertiesResponseOutput {
	return i.ToSmsChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i SmsChannelPropertiesResponseArgs) ToSmsChannelPropertiesResponseOutputWithContext(ctx context.Context) SmsChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPropertiesResponseOutput)
}

func (i SmsChannelPropertiesResponseArgs) ToSmsChannelPropertiesResponsePtrOutput() SmsChannelPropertiesResponsePtrOutput {
	return i.ToSmsChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i SmsChannelPropertiesResponseArgs) ToSmsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SmsChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPropertiesResponseOutput).ToSmsChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type SmsChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToSmsChannelPropertiesResponsePtrOutput() SmsChannelPropertiesResponsePtrOutput
	ToSmsChannelPropertiesResponsePtrOutputWithContext(context.Context) SmsChannelPropertiesResponsePtrOutput
}

type smsChannelPropertiesResponsePtrType SmsChannelPropertiesResponseArgs

func SmsChannelPropertiesResponsePtr(v *SmsChannelPropertiesResponseArgs) SmsChannelPropertiesResponsePtrInput {
	return (*smsChannelPropertiesResponsePtrType)(v)
}

func (*smsChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannelPropertiesResponse)(nil)).Elem()
}

func (i *smsChannelPropertiesResponsePtrType) ToSmsChannelPropertiesResponsePtrOutput() SmsChannelPropertiesResponsePtrOutput {
	return i.ToSmsChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *smsChannelPropertiesResponsePtrType) ToSmsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SmsChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelPropertiesResponsePtrOutput)
}

type SmsChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (SmsChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannelPropertiesResponse)(nil)).Elem()
}

func (o SmsChannelPropertiesResponseOutput) ToSmsChannelPropertiesResponseOutput() SmsChannelPropertiesResponseOutput {
	return o
}

func (o SmsChannelPropertiesResponseOutput) ToSmsChannelPropertiesResponseOutputWithContext(ctx context.Context) SmsChannelPropertiesResponseOutput {
	return o
}

func (o SmsChannelPropertiesResponseOutput) ToSmsChannelPropertiesResponsePtrOutput() SmsChannelPropertiesResponsePtrOutput {
	return o.ToSmsChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o SmsChannelPropertiesResponseOutput) ToSmsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SmsChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v SmsChannelPropertiesResponse) *SmsChannelPropertiesResponse {
		return &v
	}).(SmsChannelPropertiesResponsePtrOutput)
}

func (o SmsChannelPropertiesResponseOutput) AccountSID() pulumi.StringOutput {
	return o.ApplyT(func(v SmsChannelPropertiesResponse) string { return v.AccountSID }).(pulumi.StringOutput)
}

func (o SmsChannelPropertiesResponseOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v SmsChannelPropertiesResponse) *string { return v.AuthToken }).(pulumi.StringPtrOutput)
}

func (o SmsChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v SmsChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o SmsChannelPropertiesResponseOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v SmsChannelPropertiesResponse) *bool { return v.IsValidated }).(pulumi.BoolPtrOutput)
}

func (o SmsChannelPropertiesResponseOutput) Phone() pulumi.StringOutput {
	return o.ApplyT(func(v SmsChannelPropertiesResponse) string { return v.Phone }).(pulumi.StringOutput)
}

type SmsChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (SmsChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SmsChannelPropertiesResponse)(nil)).Elem()
}

func (o SmsChannelPropertiesResponsePtrOutput) ToSmsChannelPropertiesResponsePtrOutput() SmsChannelPropertiesResponsePtrOutput {
	return o
}

func (o SmsChannelPropertiesResponsePtrOutput) ToSmsChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) SmsChannelPropertiesResponsePtrOutput {
	return o
}

func (o SmsChannelPropertiesResponsePtrOutput) Elem() SmsChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *SmsChannelPropertiesResponse) SmsChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret SmsChannelPropertiesResponse
		return ret
	}).(SmsChannelPropertiesResponseOutput)
}

func (o SmsChannelPropertiesResponsePtrOutput) AccountSID() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.AccountSID
	}).(pulumi.StringPtrOutput)
}

func (o SmsChannelPropertiesResponsePtrOutput) AuthToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AuthToken
	}).(pulumi.StringPtrOutput)
}

func (o SmsChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SmsChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o SmsChannelPropertiesResponsePtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SmsChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

func (o SmsChannelPropertiesResponsePtrOutput) Phone() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SmsChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.Phone
	}).(pulumi.StringPtrOutput)
}

type SmsChannelResponse struct {
	ChannelName string                        `pulumi:"channelName"`
	Properties  *SmsChannelPropertiesResponse `pulumi:"properties"`
}





type SmsChannelResponseInput interface {
	pulumi.Input

	ToSmsChannelResponseOutput() SmsChannelResponseOutput
	ToSmsChannelResponseOutputWithContext(context.Context) SmsChannelResponseOutput
}

type SmsChannelResponseArgs struct {
	ChannelName pulumi.StringInput                   `pulumi:"channelName"`
	Properties  SmsChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (SmsChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannelResponse)(nil)).Elem()
}

func (i SmsChannelResponseArgs) ToSmsChannelResponseOutput() SmsChannelResponseOutput {
	return i.ToSmsChannelResponseOutputWithContext(context.Background())
}

func (i SmsChannelResponseArgs) ToSmsChannelResponseOutputWithContext(ctx context.Context) SmsChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SmsChannelResponseOutput)
}

type SmsChannelResponseOutput struct{ *pulumi.OutputState }

func (SmsChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*SmsChannelResponse)(nil)).Elem()
}

func (o SmsChannelResponseOutput) ToSmsChannelResponseOutput() SmsChannelResponseOutput {
	return o
}

func (o SmsChannelResponseOutput) ToSmsChannelResponseOutputWithContext(ctx context.Context) SmsChannelResponseOutput {
	return o
}

func (o SmsChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v SmsChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o SmsChannelResponseOutput) Properties() SmsChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v SmsChannelResponse) *SmsChannelPropertiesResponse { return v.Properties }).(SmsChannelPropertiesResponsePtrOutput)
}

type TelegramChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Properties  *TelegramChannelProperties `pulumi:"properties"`
}





type TelegramChannelInput interface {
	pulumi.Input

	ToTelegramChannelOutput() TelegramChannelOutput
	ToTelegramChannelOutputWithContext(context.Context) TelegramChannelOutput
}

type TelegramChannelArgs struct {
	ChannelName pulumi.StringInput                `pulumi:"channelName"`
	Properties  TelegramChannelPropertiesPtrInput `pulumi:"properties"`
}

func (TelegramChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannel)(nil)).Elem()
}

func (i TelegramChannelArgs) ToTelegramChannelOutput() TelegramChannelOutput {
	return i.ToTelegramChannelOutputWithContext(context.Background())
}

func (i TelegramChannelArgs) ToTelegramChannelOutputWithContext(ctx context.Context) TelegramChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelOutput)
}

type TelegramChannelOutput struct{ *pulumi.OutputState }

func (TelegramChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannel)(nil)).Elem()
}

func (o TelegramChannelOutput) ToTelegramChannelOutput() TelegramChannelOutput {
	return o
}

func (o TelegramChannelOutput) ToTelegramChannelOutputWithContext(ctx context.Context) TelegramChannelOutput {
	return o
}

func (o TelegramChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v TelegramChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o TelegramChannelOutput) Properties() TelegramChannelPropertiesPtrOutput {
	return o.ApplyT(func(v TelegramChannel) *TelegramChannelProperties { return v.Properties }).(TelegramChannelPropertiesPtrOutput)
}

type TelegramChannelProperties struct {
	AccessToken *string `pulumi:"accessToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
}





type TelegramChannelPropertiesInput interface {
	pulumi.Input

	ToTelegramChannelPropertiesOutput() TelegramChannelPropertiesOutput
	ToTelegramChannelPropertiesOutputWithContext(context.Context) TelegramChannelPropertiesOutput
}

type TelegramChannelPropertiesArgs struct {
	AccessToken pulumi.StringPtrInput `pulumi:"accessToken"`
	IsEnabled   pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated pulumi.BoolPtrInput   `pulumi:"isValidated"`
}

func (TelegramChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannelProperties)(nil)).Elem()
}

func (i TelegramChannelPropertiesArgs) ToTelegramChannelPropertiesOutput() TelegramChannelPropertiesOutput {
	return i.ToTelegramChannelPropertiesOutputWithContext(context.Background())
}

func (i TelegramChannelPropertiesArgs) ToTelegramChannelPropertiesOutputWithContext(ctx context.Context) TelegramChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelPropertiesOutput)
}

func (i TelegramChannelPropertiesArgs) ToTelegramChannelPropertiesPtrOutput() TelegramChannelPropertiesPtrOutput {
	return i.ToTelegramChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i TelegramChannelPropertiesArgs) ToTelegramChannelPropertiesPtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelPropertiesOutput).ToTelegramChannelPropertiesPtrOutputWithContext(ctx)
}









type TelegramChannelPropertiesPtrInput interface {
	pulumi.Input

	ToTelegramChannelPropertiesPtrOutput() TelegramChannelPropertiesPtrOutput
	ToTelegramChannelPropertiesPtrOutputWithContext(context.Context) TelegramChannelPropertiesPtrOutput
}

type telegramChannelPropertiesPtrType TelegramChannelPropertiesArgs

func TelegramChannelPropertiesPtr(v *TelegramChannelPropertiesArgs) TelegramChannelPropertiesPtrInput {
	return (*telegramChannelPropertiesPtrType)(v)
}

func (*telegramChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TelegramChannelProperties)(nil)).Elem()
}

func (i *telegramChannelPropertiesPtrType) ToTelegramChannelPropertiesPtrOutput() TelegramChannelPropertiesPtrOutput {
	return i.ToTelegramChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *telegramChannelPropertiesPtrType) ToTelegramChannelPropertiesPtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelPropertiesPtrOutput)
}

type TelegramChannelPropertiesOutput struct{ *pulumi.OutputState }

func (TelegramChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannelProperties)(nil)).Elem()
}

func (o TelegramChannelPropertiesOutput) ToTelegramChannelPropertiesOutput() TelegramChannelPropertiesOutput {
	return o
}

func (o TelegramChannelPropertiesOutput) ToTelegramChannelPropertiesOutputWithContext(ctx context.Context) TelegramChannelPropertiesOutput {
	return o
}

func (o TelegramChannelPropertiesOutput) ToTelegramChannelPropertiesPtrOutput() TelegramChannelPropertiesPtrOutput {
	return o.ToTelegramChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o TelegramChannelPropertiesOutput) ToTelegramChannelPropertiesPtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TelegramChannelProperties) *TelegramChannelProperties {
		return &v
	}).(TelegramChannelPropertiesPtrOutput)
}

func (o TelegramChannelPropertiesOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TelegramChannelProperties) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o TelegramChannelPropertiesOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TelegramChannelProperties) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o TelegramChannelPropertiesOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TelegramChannelProperties) *bool { return v.IsValidated }).(pulumi.BoolPtrOutput)
}

type TelegramChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (TelegramChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TelegramChannelProperties)(nil)).Elem()
}

func (o TelegramChannelPropertiesPtrOutput) ToTelegramChannelPropertiesPtrOutput() TelegramChannelPropertiesPtrOutput {
	return o
}

func (o TelegramChannelPropertiesPtrOutput) ToTelegramChannelPropertiesPtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesPtrOutput {
	return o
}

func (o TelegramChannelPropertiesPtrOutput) Elem() TelegramChannelPropertiesOutput {
	return o.ApplyT(func(v *TelegramChannelProperties) TelegramChannelProperties {
		if v != nil {
			return *v
		}
		var ret TelegramChannelProperties
		return ret
	}).(TelegramChannelPropertiesOutput)
}

func (o TelegramChannelPropertiesPtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TelegramChannelProperties) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o TelegramChannelPropertiesPtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TelegramChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o TelegramChannelPropertiesPtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TelegramChannelProperties) *bool {
		if v == nil {
			return nil
		}
		return v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

type TelegramChannelPropertiesResponse struct {
	AccessToken *string `pulumi:"accessToken"`
	IsEnabled   bool    `pulumi:"isEnabled"`
	IsValidated *bool   `pulumi:"isValidated"`
}





type TelegramChannelPropertiesResponseInput interface {
	pulumi.Input

	ToTelegramChannelPropertiesResponseOutput() TelegramChannelPropertiesResponseOutput
	ToTelegramChannelPropertiesResponseOutputWithContext(context.Context) TelegramChannelPropertiesResponseOutput
}

type TelegramChannelPropertiesResponseArgs struct {
	AccessToken pulumi.StringPtrInput `pulumi:"accessToken"`
	IsEnabled   pulumi.BoolInput      `pulumi:"isEnabled"`
	IsValidated pulumi.BoolPtrInput   `pulumi:"isValidated"`
}

func (TelegramChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannelPropertiesResponse)(nil)).Elem()
}

func (i TelegramChannelPropertiesResponseArgs) ToTelegramChannelPropertiesResponseOutput() TelegramChannelPropertiesResponseOutput {
	return i.ToTelegramChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i TelegramChannelPropertiesResponseArgs) ToTelegramChannelPropertiesResponseOutputWithContext(ctx context.Context) TelegramChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelPropertiesResponseOutput)
}

func (i TelegramChannelPropertiesResponseArgs) ToTelegramChannelPropertiesResponsePtrOutput() TelegramChannelPropertiesResponsePtrOutput {
	return i.ToTelegramChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i TelegramChannelPropertiesResponseArgs) ToTelegramChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelPropertiesResponseOutput).ToTelegramChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type TelegramChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToTelegramChannelPropertiesResponsePtrOutput() TelegramChannelPropertiesResponsePtrOutput
	ToTelegramChannelPropertiesResponsePtrOutputWithContext(context.Context) TelegramChannelPropertiesResponsePtrOutput
}

type telegramChannelPropertiesResponsePtrType TelegramChannelPropertiesResponseArgs

func TelegramChannelPropertiesResponsePtr(v *TelegramChannelPropertiesResponseArgs) TelegramChannelPropertiesResponsePtrInput {
	return (*telegramChannelPropertiesResponsePtrType)(v)
}

func (*telegramChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**TelegramChannelPropertiesResponse)(nil)).Elem()
}

func (i *telegramChannelPropertiesResponsePtrType) ToTelegramChannelPropertiesResponsePtrOutput() TelegramChannelPropertiesResponsePtrOutput {
	return i.ToTelegramChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *telegramChannelPropertiesResponsePtrType) ToTelegramChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelPropertiesResponsePtrOutput)
}

type TelegramChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (TelegramChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannelPropertiesResponse)(nil)).Elem()
}

func (o TelegramChannelPropertiesResponseOutput) ToTelegramChannelPropertiesResponseOutput() TelegramChannelPropertiesResponseOutput {
	return o
}

func (o TelegramChannelPropertiesResponseOutput) ToTelegramChannelPropertiesResponseOutputWithContext(ctx context.Context) TelegramChannelPropertiesResponseOutput {
	return o
}

func (o TelegramChannelPropertiesResponseOutput) ToTelegramChannelPropertiesResponsePtrOutput() TelegramChannelPropertiesResponsePtrOutput {
	return o.ToTelegramChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o TelegramChannelPropertiesResponseOutput) ToTelegramChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v TelegramChannelPropertiesResponse) *TelegramChannelPropertiesResponse {
		return &v
	}).(TelegramChannelPropertiesResponsePtrOutput)
}

func (o TelegramChannelPropertiesResponseOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v TelegramChannelPropertiesResponse) *string { return v.AccessToken }).(pulumi.StringPtrOutput)
}

func (o TelegramChannelPropertiesResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v TelegramChannelPropertiesResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o TelegramChannelPropertiesResponseOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v TelegramChannelPropertiesResponse) *bool { return v.IsValidated }).(pulumi.BoolPtrOutput)
}

type TelegramChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (TelegramChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**TelegramChannelPropertiesResponse)(nil)).Elem()
}

func (o TelegramChannelPropertiesResponsePtrOutput) ToTelegramChannelPropertiesResponsePtrOutput() TelegramChannelPropertiesResponsePtrOutput {
	return o
}

func (o TelegramChannelPropertiesResponsePtrOutput) ToTelegramChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) TelegramChannelPropertiesResponsePtrOutput {
	return o
}

func (o TelegramChannelPropertiesResponsePtrOutput) Elem() TelegramChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *TelegramChannelPropertiesResponse) TelegramChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret TelegramChannelPropertiesResponse
		return ret
	}).(TelegramChannelPropertiesResponseOutput)
}

func (o TelegramChannelPropertiesResponsePtrOutput) AccessToken() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *TelegramChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return v.AccessToken
	}).(pulumi.StringPtrOutput)
}

func (o TelegramChannelPropertiesResponsePtrOutput) IsEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TelegramChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return &v.IsEnabled
	}).(pulumi.BoolPtrOutput)
}

func (o TelegramChannelPropertiesResponsePtrOutput) IsValidated() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *TelegramChannelPropertiesResponse) *bool {
		if v == nil {
			return nil
		}
		return v.IsValidated
	}).(pulumi.BoolPtrOutput)
}

type TelegramChannelResponse struct {
	ChannelName string                             `pulumi:"channelName"`
	Properties  *TelegramChannelPropertiesResponse `pulumi:"properties"`
}





type TelegramChannelResponseInput interface {
	pulumi.Input

	ToTelegramChannelResponseOutput() TelegramChannelResponseOutput
	ToTelegramChannelResponseOutputWithContext(context.Context) TelegramChannelResponseOutput
}

type TelegramChannelResponseArgs struct {
	ChannelName pulumi.StringInput                        `pulumi:"channelName"`
	Properties  TelegramChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (TelegramChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannelResponse)(nil)).Elem()
}

func (i TelegramChannelResponseArgs) ToTelegramChannelResponseOutput() TelegramChannelResponseOutput {
	return i.ToTelegramChannelResponseOutputWithContext(context.Background())
}

func (i TelegramChannelResponseArgs) ToTelegramChannelResponseOutputWithContext(ctx context.Context) TelegramChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(TelegramChannelResponseOutput)
}

type TelegramChannelResponseOutput struct{ *pulumi.OutputState }

func (TelegramChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*TelegramChannelResponse)(nil)).Elem()
}

func (o TelegramChannelResponseOutput) ToTelegramChannelResponseOutput() TelegramChannelResponseOutput {
	return o
}

func (o TelegramChannelResponseOutput) ToTelegramChannelResponseOutputWithContext(ctx context.Context) TelegramChannelResponseOutput {
	return o
}

func (o TelegramChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v TelegramChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o TelegramChannelResponseOutput) Properties() TelegramChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v TelegramChannelResponse) *TelegramChannelPropertiesResponse { return v.Properties }).(TelegramChannelPropertiesResponsePtrOutput)
}

type WebChatChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Properties  *WebChatChannelProperties `pulumi:"properties"`
}





type WebChatChannelInput interface {
	pulumi.Input

	ToWebChatChannelOutput() WebChatChannelOutput
	ToWebChatChannelOutputWithContext(context.Context) WebChatChannelOutput
}

type WebChatChannelArgs struct {
	ChannelName pulumi.StringInput               `pulumi:"channelName"`
	Properties  WebChatChannelPropertiesPtrInput `pulumi:"properties"`
}

func (WebChatChannelArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannel)(nil)).Elem()
}

func (i WebChatChannelArgs) ToWebChatChannelOutput() WebChatChannelOutput {
	return i.ToWebChatChannelOutputWithContext(context.Background())
}

func (i WebChatChannelArgs) ToWebChatChannelOutputWithContext(ctx context.Context) WebChatChannelOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelOutput)
}

type WebChatChannelOutput struct{ *pulumi.OutputState }

func (WebChatChannelOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannel)(nil)).Elem()
}

func (o WebChatChannelOutput) ToWebChatChannelOutput() WebChatChannelOutput {
	return o
}

func (o WebChatChannelOutput) ToWebChatChannelOutputWithContext(ctx context.Context) WebChatChannelOutput {
	return o
}

func (o WebChatChannelOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatChannel) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o WebChatChannelOutput) Properties() WebChatChannelPropertiesPtrOutput {
	return o.ApplyT(func(v WebChatChannel) *WebChatChannelProperties { return v.Properties }).(WebChatChannelPropertiesPtrOutput)
}

type WebChatChannelProperties struct {
	Sites []WebChatSite `pulumi:"sites"`
}





type WebChatChannelPropertiesInput interface {
	pulumi.Input

	ToWebChatChannelPropertiesOutput() WebChatChannelPropertiesOutput
	ToWebChatChannelPropertiesOutputWithContext(context.Context) WebChatChannelPropertiesOutput
}

type WebChatChannelPropertiesArgs struct {
	Sites WebChatSiteArrayInput `pulumi:"sites"`
}

func (WebChatChannelPropertiesArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannelProperties)(nil)).Elem()
}

func (i WebChatChannelPropertiesArgs) ToWebChatChannelPropertiesOutput() WebChatChannelPropertiesOutput {
	return i.ToWebChatChannelPropertiesOutputWithContext(context.Background())
}

func (i WebChatChannelPropertiesArgs) ToWebChatChannelPropertiesOutputWithContext(ctx context.Context) WebChatChannelPropertiesOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelPropertiesOutput)
}

func (i WebChatChannelPropertiesArgs) ToWebChatChannelPropertiesPtrOutput() WebChatChannelPropertiesPtrOutput {
	return i.ToWebChatChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i WebChatChannelPropertiesArgs) ToWebChatChannelPropertiesPtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelPropertiesOutput).ToWebChatChannelPropertiesPtrOutputWithContext(ctx)
}









type WebChatChannelPropertiesPtrInput interface {
	pulumi.Input

	ToWebChatChannelPropertiesPtrOutput() WebChatChannelPropertiesPtrOutput
	ToWebChatChannelPropertiesPtrOutputWithContext(context.Context) WebChatChannelPropertiesPtrOutput
}

type webChatChannelPropertiesPtrType WebChatChannelPropertiesArgs

func WebChatChannelPropertiesPtr(v *WebChatChannelPropertiesArgs) WebChatChannelPropertiesPtrInput {
	return (*webChatChannelPropertiesPtrType)(v)
}

func (*webChatChannelPropertiesPtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebChatChannelProperties)(nil)).Elem()
}

func (i *webChatChannelPropertiesPtrType) ToWebChatChannelPropertiesPtrOutput() WebChatChannelPropertiesPtrOutput {
	return i.ToWebChatChannelPropertiesPtrOutputWithContext(context.Background())
}

func (i *webChatChannelPropertiesPtrType) ToWebChatChannelPropertiesPtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesPtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelPropertiesPtrOutput)
}

type WebChatChannelPropertiesOutput struct{ *pulumi.OutputState }

func (WebChatChannelPropertiesOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannelProperties)(nil)).Elem()
}

func (o WebChatChannelPropertiesOutput) ToWebChatChannelPropertiesOutput() WebChatChannelPropertiesOutput {
	return o
}

func (o WebChatChannelPropertiesOutput) ToWebChatChannelPropertiesOutputWithContext(ctx context.Context) WebChatChannelPropertiesOutput {
	return o
}

func (o WebChatChannelPropertiesOutput) ToWebChatChannelPropertiesPtrOutput() WebChatChannelPropertiesPtrOutput {
	return o.ToWebChatChannelPropertiesPtrOutputWithContext(context.Background())
}

func (o WebChatChannelPropertiesOutput) ToWebChatChannelPropertiesPtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesPtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebChatChannelProperties) *WebChatChannelProperties {
		return &v
	}).(WebChatChannelPropertiesPtrOutput)
}

func (o WebChatChannelPropertiesOutput) Sites() WebChatSiteArrayOutput {
	return o.ApplyT(func(v WebChatChannelProperties) []WebChatSite { return v.Sites }).(WebChatSiteArrayOutput)
}

type WebChatChannelPropertiesPtrOutput struct{ *pulumi.OutputState }

func (WebChatChannelPropertiesPtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebChatChannelProperties)(nil)).Elem()
}

func (o WebChatChannelPropertiesPtrOutput) ToWebChatChannelPropertiesPtrOutput() WebChatChannelPropertiesPtrOutput {
	return o
}

func (o WebChatChannelPropertiesPtrOutput) ToWebChatChannelPropertiesPtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesPtrOutput {
	return o
}

func (o WebChatChannelPropertiesPtrOutput) Elem() WebChatChannelPropertiesOutput {
	return o.ApplyT(func(v *WebChatChannelProperties) WebChatChannelProperties {
		if v != nil {
			return *v
		}
		var ret WebChatChannelProperties
		return ret
	}).(WebChatChannelPropertiesOutput)
}

func (o WebChatChannelPropertiesPtrOutput) Sites() WebChatSiteArrayOutput {
	return o.ApplyT(func(v *WebChatChannelProperties) []WebChatSite {
		if v == nil {
			return nil
		}
		return v.Sites
	}).(WebChatSiteArrayOutput)
}

type WebChatChannelPropertiesResponse struct {
	Sites            []WebChatSiteResponse `pulumi:"sites"`
	WebChatEmbedCode string                `pulumi:"webChatEmbedCode"`
}





type WebChatChannelPropertiesResponseInput interface {
	pulumi.Input

	ToWebChatChannelPropertiesResponseOutput() WebChatChannelPropertiesResponseOutput
	ToWebChatChannelPropertiesResponseOutputWithContext(context.Context) WebChatChannelPropertiesResponseOutput
}

type WebChatChannelPropertiesResponseArgs struct {
	Sites            WebChatSiteResponseArrayInput `pulumi:"sites"`
	WebChatEmbedCode pulumi.StringInput            `pulumi:"webChatEmbedCode"`
}

func (WebChatChannelPropertiesResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannelPropertiesResponse)(nil)).Elem()
}

func (i WebChatChannelPropertiesResponseArgs) ToWebChatChannelPropertiesResponseOutput() WebChatChannelPropertiesResponseOutput {
	return i.ToWebChatChannelPropertiesResponseOutputWithContext(context.Background())
}

func (i WebChatChannelPropertiesResponseArgs) ToWebChatChannelPropertiesResponseOutputWithContext(ctx context.Context) WebChatChannelPropertiesResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelPropertiesResponseOutput)
}

func (i WebChatChannelPropertiesResponseArgs) ToWebChatChannelPropertiesResponsePtrOutput() WebChatChannelPropertiesResponsePtrOutput {
	return i.ToWebChatChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i WebChatChannelPropertiesResponseArgs) ToWebChatChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelPropertiesResponseOutput).ToWebChatChannelPropertiesResponsePtrOutputWithContext(ctx)
}









type WebChatChannelPropertiesResponsePtrInput interface {
	pulumi.Input

	ToWebChatChannelPropertiesResponsePtrOutput() WebChatChannelPropertiesResponsePtrOutput
	ToWebChatChannelPropertiesResponsePtrOutputWithContext(context.Context) WebChatChannelPropertiesResponsePtrOutput
}

type webChatChannelPropertiesResponsePtrType WebChatChannelPropertiesResponseArgs

func WebChatChannelPropertiesResponsePtr(v *WebChatChannelPropertiesResponseArgs) WebChatChannelPropertiesResponsePtrInput {
	return (*webChatChannelPropertiesResponsePtrType)(v)
}

func (*webChatChannelPropertiesResponsePtrType) ElementType() reflect.Type {
	return reflect.TypeOf((**WebChatChannelPropertiesResponse)(nil)).Elem()
}

func (i *webChatChannelPropertiesResponsePtrType) ToWebChatChannelPropertiesResponsePtrOutput() WebChatChannelPropertiesResponsePtrOutput {
	return i.ToWebChatChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (i *webChatChannelPropertiesResponsePtrType) ToWebChatChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesResponsePtrOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelPropertiesResponsePtrOutput)
}

type WebChatChannelPropertiesResponseOutput struct{ *pulumi.OutputState }

func (WebChatChannelPropertiesResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannelPropertiesResponse)(nil)).Elem()
}

func (o WebChatChannelPropertiesResponseOutput) ToWebChatChannelPropertiesResponseOutput() WebChatChannelPropertiesResponseOutput {
	return o
}

func (o WebChatChannelPropertiesResponseOutput) ToWebChatChannelPropertiesResponseOutputWithContext(ctx context.Context) WebChatChannelPropertiesResponseOutput {
	return o
}

func (o WebChatChannelPropertiesResponseOutput) ToWebChatChannelPropertiesResponsePtrOutput() WebChatChannelPropertiesResponsePtrOutput {
	return o.ToWebChatChannelPropertiesResponsePtrOutputWithContext(context.Background())
}

func (o WebChatChannelPropertiesResponseOutput) ToWebChatChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesResponsePtrOutput {
	return o.ApplyTWithContext(ctx, func(_ context.Context, v WebChatChannelPropertiesResponse) *WebChatChannelPropertiesResponse {
		return &v
	}).(WebChatChannelPropertiesResponsePtrOutput)
}

func (o WebChatChannelPropertiesResponseOutput) Sites() WebChatSiteResponseArrayOutput {
	return o.ApplyT(func(v WebChatChannelPropertiesResponse) []WebChatSiteResponse { return v.Sites }).(WebChatSiteResponseArrayOutput)
}

func (o WebChatChannelPropertiesResponseOutput) WebChatEmbedCode() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatChannelPropertiesResponse) string { return v.WebChatEmbedCode }).(pulumi.StringOutput)
}

type WebChatChannelPropertiesResponsePtrOutput struct{ *pulumi.OutputState }

func (WebChatChannelPropertiesResponsePtrOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebChatChannelPropertiesResponse)(nil)).Elem()
}

func (o WebChatChannelPropertiesResponsePtrOutput) ToWebChatChannelPropertiesResponsePtrOutput() WebChatChannelPropertiesResponsePtrOutput {
	return o
}

func (o WebChatChannelPropertiesResponsePtrOutput) ToWebChatChannelPropertiesResponsePtrOutputWithContext(ctx context.Context) WebChatChannelPropertiesResponsePtrOutput {
	return o
}

func (o WebChatChannelPropertiesResponsePtrOutput) Elem() WebChatChannelPropertiesResponseOutput {
	return o.ApplyT(func(v *WebChatChannelPropertiesResponse) WebChatChannelPropertiesResponse {
		if v != nil {
			return *v
		}
		var ret WebChatChannelPropertiesResponse
		return ret
	}).(WebChatChannelPropertiesResponseOutput)
}

func (o WebChatChannelPropertiesResponsePtrOutput) Sites() WebChatSiteResponseArrayOutput {
	return o.ApplyT(func(v *WebChatChannelPropertiesResponse) []WebChatSiteResponse {
		if v == nil {
			return nil
		}
		return v.Sites
	}).(WebChatSiteResponseArrayOutput)
}

func (o WebChatChannelPropertiesResponsePtrOutput) WebChatEmbedCode() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebChatChannelPropertiesResponse) *string {
		if v == nil {
			return nil
		}
		return &v.WebChatEmbedCode
	}).(pulumi.StringPtrOutput)
}

type WebChatChannelResponse struct {
	ChannelName string                            `pulumi:"channelName"`
	Properties  *WebChatChannelPropertiesResponse `pulumi:"properties"`
}





type WebChatChannelResponseInput interface {
	pulumi.Input

	ToWebChatChannelResponseOutput() WebChatChannelResponseOutput
	ToWebChatChannelResponseOutputWithContext(context.Context) WebChatChannelResponseOutput
}

type WebChatChannelResponseArgs struct {
	ChannelName pulumi.StringInput                       `pulumi:"channelName"`
	Properties  WebChatChannelPropertiesResponsePtrInput `pulumi:"properties"`
}

func (WebChatChannelResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannelResponse)(nil)).Elem()
}

func (i WebChatChannelResponseArgs) ToWebChatChannelResponseOutput() WebChatChannelResponseOutput {
	return i.ToWebChatChannelResponseOutputWithContext(context.Background())
}

func (i WebChatChannelResponseArgs) ToWebChatChannelResponseOutputWithContext(ctx context.Context) WebChatChannelResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatChannelResponseOutput)
}

type WebChatChannelResponseOutput struct{ *pulumi.OutputState }

func (WebChatChannelResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatChannelResponse)(nil)).Elem()
}

func (o WebChatChannelResponseOutput) ToWebChatChannelResponseOutput() WebChatChannelResponseOutput {
	return o
}

func (o WebChatChannelResponseOutput) ToWebChatChannelResponseOutputWithContext(ctx context.Context) WebChatChannelResponseOutput {
	return o
}

func (o WebChatChannelResponseOutput) ChannelName() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatChannelResponse) string { return v.ChannelName }).(pulumi.StringOutput)
}

func (o WebChatChannelResponseOutput) Properties() WebChatChannelPropertiesResponsePtrOutput {
	return o.ApplyT(func(v WebChatChannelResponse) *WebChatChannelPropertiesResponse { return v.Properties }).(WebChatChannelPropertiesResponsePtrOutput)
}

type WebChatSite struct {
	EnablePreview bool   `pulumi:"enablePreview"`
	IsEnabled     bool   `pulumi:"isEnabled"`
	SiteName      string `pulumi:"siteName"`
}





type WebChatSiteInput interface {
	pulumi.Input

	ToWebChatSiteOutput() WebChatSiteOutput
	ToWebChatSiteOutputWithContext(context.Context) WebChatSiteOutput
}

type WebChatSiteArgs struct {
	EnablePreview pulumi.BoolInput   `pulumi:"enablePreview"`
	IsEnabled     pulumi.BoolInput   `pulumi:"isEnabled"`
	SiteName      pulumi.StringInput `pulumi:"siteName"`
}

func (WebChatSiteArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatSite)(nil)).Elem()
}

func (i WebChatSiteArgs) ToWebChatSiteOutput() WebChatSiteOutput {
	return i.ToWebChatSiteOutputWithContext(context.Background())
}

func (i WebChatSiteArgs) ToWebChatSiteOutputWithContext(ctx context.Context) WebChatSiteOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatSiteOutput)
}





type WebChatSiteArrayInput interface {
	pulumi.Input

	ToWebChatSiteArrayOutput() WebChatSiteArrayOutput
	ToWebChatSiteArrayOutputWithContext(context.Context) WebChatSiteArrayOutput
}

type WebChatSiteArray []WebChatSiteInput

func (WebChatSiteArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebChatSite)(nil)).Elem()
}

func (i WebChatSiteArray) ToWebChatSiteArrayOutput() WebChatSiteArrayOutput {
	return i.ToWebChatSiteArrayOutputWithContext(context.Background())
}

func (i WebChatSiteArray) ToWebChatSiteArrayOutputWithContext(ctx context.Context) WebChatSiteArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatSiteArrayOutput)
}

type WebChatSiteOutput struct{ *pulumi.OutputState }

func (WebChatSiteOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatSite)(nil)).Elem()
}

func (o WebChatSiteOutput) ToWebChatSiteOutput() WebChatSiteOutput {
	return o
}

func (o WebChatSiteOutput) ToWebChatSiteOutputWithContext(ctx context.Context) WebChatSiteOutput {
	return o
}

func (o WebChatSiteOutput) EnablePreview() pulumi.BoolOutput {
	return o.ApplyT(func(v WebChatSite) bool { return v.EnablePreview }).(pulumi.BoolOutput)
}

func (o WebChatSiteOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v WebChatSite) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o WebChatSiteOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatSite) string { return v.SiteName }).(pulumi.StringOutput)
}

type WebChatSiteArrayOutput struct{ *pulumi.OutputState }

func (WebChatSiteArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebChatSite)(nil)).Elem()
}

func (o WebChatSiteArrayOutput) ToWebChatSiteArrayOutput() WebChatSiteArrayOutput {
	return o
}

func (o WebChatSiteArrayOutput) ToWebChatSiteArrayOutputWithContext(ctx context.Context) WebChatSiteArrayOutput {
	return o
}

func (o WebChatSiteArrayOutput) Index(i pulumi.IntInput) WebChatSiteOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebChatSite {
		return vs[0].([]WebChatSite)[vs[1].(int)]
	}).(WebChatSiteOutput)
}

type WebChatSiteResponse struct {
	EnablePreview bool   `pulumi:"enablePreview"`
	IsEnabled     bool   `pulumi:"isEnabled"`
	Key           string `pulumi:"key"`
	Key2          string `pulumi:"key2"`
	SiteId        string `pulumi:"siteId"`
	SiteName      string `pulumi:"siteName"`
}





type WebChatSiteResponseInput interface {
	pulumi.Input

	ToWebChatSiteResponseOutput() WebChatSiteResponseOutput
	ToWebChatSiteResponseOutputWithContext(context.Context) WebChatSiteResponseOutput
}

type WebChatSiteResponseArgs struct {
	EnablePreview pulumi.BoolInput   `pulumi:"enablePreview"`
	IsEnabled     pulumi.BoolInput   `pulumi:"isEnabled"`
	Key           pulumi.StringInput `pulumi:"key"`
	Key2          pulumi.StringInput `pulumi:"key2"`
	SiteId        pulumi.StringInput `pulumi:"siteId"`
	SiteName      pulumi.StringInput `pulumi:"siteName"`
}

func (WebChatSiteResponseArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatSiteResponse)(nil)).Elem()
}

func (i WebChatSiteResponseArgs) ToWebChatSiteResponseOutput() WebChatSiteResponseOutput {
	return i.ToWebChatSiteResponseOutputWithContext(context.Background())
}

func (i WebChatSiteResponseArgs) ToWebChatSiteResponseOutputWithContext(ctx context.Context) WebChatSiteResponseOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatSiteResponseOutput)
}





type WebChatSiteResponseArrayInput interface {
	pulumi.Input

	ToWebChatSiteResponseArrayOutput() WebChatSiteResponseArrayOutput
	ToWebChatSiteResponseArrayOutputWithContext(context.Context) WebChatSiteResponseArrayOutput
}

type WebChatSiteResponseArray []WebChatSiteResponseInput

func (WebChatSiteResponseArray) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebChatSiteResponse)(nil)).Elem()
}

func (i WebChatSiteResponseArray) ToWebChatSiteResponseArrayOutput() WebChatSiteResponseArrayOutput {
	return i.ToWebChatSiteResponseArrayOutputWithContext(context.Background())
}

func (i WebChatSiteResponseArray) ToWebChatSiteResponseArrayOutputWithContext(ctx context.Context) WebChatSiteResponseArrayOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebChatSiteResponseArrayOutput)
}

type WebChatSiteResponseOutput struct{ *pulumi.OutputState }

func (WebChatSiteResponseOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebChatSiteResponse)(nil)).Elem()
}

func (o WebChatSiteResponseOutput) ToWebChatSiteResponseOutput() WebChatSiteResponseOutput {
	return o
}

func (o WebChatSiteResponseOutput) ToWebChatSiteResponseOutputWithContext(ctx context.Context) WebChatSiteResponseOutput {
	return o
}

func (o WebChatSiteResponseOutput) EnablePreview() pulumi.BoolOutput {
	return o.ApplyT(func(v WebChatSiteResponse) bool { return v.EnablePreview }).(pulumi.BoolOutput)
}

func (o WebChatSiteResponseOutput) IsEnabled() pulumi.BoolOutput {
	return o.ApplyT(func(v WebChatSiteResponse) bool { return v.IsEnabled }).(pulumi.BoolOutput)
}

func (o WebChatSiteResponseOutput) Key() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatSiteResponse) string { return v.Key }).(pulumi.StringOutput)
}

func (o WebChatSiteResponseOutput) Key2() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatSiteResponse) string { return v.Key2 }).(pulumi.StringOutput)
}

func (o WebChatSiteResponseOutput) SiteId() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatSiteResponse) string { return v.SiteId }).(pulumi.StringOutput)
}

func (o WebChatSiteResponseOutput) SiteName() pulumi.StringOutput {
	return o.ApplyT(func(v WebChatSiteResponse) string { return v.SiteName }).(pulumi.StringOutput)
}

type WebChatSiteResponseArrayOutput struct{ *pulumi.OutputState }

func (WebChatSiteResponseArrayOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*[]WebChatSiteResponse)(nil)).Elem()
}

func (o WebChatSiteResponseArrayOutput) ToWebChatSiteResponseArrayOutput() WebChatSiteResponseArrayOutput {
	return o
}

func (o WebChatSiteResponseArrayOutput) ToWebChatSiteResponseArrayOutputWithContext(ctx context.Context) WebChatSiteResponseArrayOutput {
	return o
}

func (o WebChatSiteResponseArrayOutput) Index(i pulumi.IntInput) WebChatSiteResponseOutput {
	return pulumi.All(o, i).ApplyT(func(vs []interface{}) WebChatSiteResponse {
		return vs[0].([]WebChatSiteResponse)[vs[1].(int)]
	}).(WebChatSiteResponseOutput)
}

func init() {
	pulumi.RegisterInputType(reflect.TypeOf((*AlexaChannelInput)(nil)).Elem(), AlexaChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlexaChannelPropertiesInput)(nil)).Elem(), AlexaChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlexaChannelPropertiesPtrInput)(nil)).Elem(), AlexaChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlexaChannelPropertiesResponseInput)(nil)).Elem(), AlexaChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlexaChannelPropertiesResponsePtrInput)(nil)).Elem(), AlexaChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*AlexaChannelResponseInput)(nil)).Elem(), AlexaChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotPropertiesInput)(nil)).Elem(), BotPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotPropertiesPtrInput)(nil)).Elem(), BotPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotPropertiesResponseInput)(nil)).Elem(), BotPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*BotPropertiesResponsePtrInput)(nil)).Elem(), BotPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingParameterInput)(nil)).Elem(), ConnectionSettingParameterArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingParameterArrayInput)(nil)).Elem(), ConnectionSettingParameterArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingParameterResponseInput)(nil)).Elem(), ConnectionSettingParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingParameterResponseArrayInput)(nil)).Elem(), ConnectionSettingParameterResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingPropertiesInput)(nil)).Elem(), ConnectionSettingPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingPropertiesPtrInput)(nil)).Elem(), ConnectionSettingPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingPropertiesResponseInput)(nil)).Elem(), ConnectionSettingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ConnectionSettingPropertiesResponsePtrInput)(nil)).Elem(), ConnectionSettingPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineChannelInput)(nil)).Elem(), DirectLineChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineChannelPropertiesInput)(nil)).Elem(), DirectLineChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineChannelPropertiesPtrInput)(nil)).Elem(), DirectLineChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineChannelPropertiesResponseInput)(nil)).Elem(), DirectLineChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineChannelPropertiesResponsePtrInput)(nil)).Elem(), DirectLineChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineChannelResponseInput)(nil)).Elem(), DirectLineChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSiteInput)(nil)).Elem(), DirectLineSiteArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSiteArrayInput)(nil)).Elem(), DirectLineSiteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSiteResponseInput)(nil)).Elem(), DirectLineSiteResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSiteResponseArrayInput)(nil)).Elem(), DirectLineSiteResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSpeechChannelInput)(nil)).Elem(), DirectLineSpeechChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSpeechChannelPropertiesInput)(nil)).Elem(), DirectLineSpeechChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSpeechChannelPropertiesPtrInput)(nil)).Elem(), DirectLineSpeechChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSpeechChannelPropertiesResponseInput)(nil)).Elem(), DirectLineSpeechChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSpeechChannelPropertiesResponsePtrInput)(nil)).Elem(), DirectLineSpeechChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*DirectLineSpeechChannelResponseInput)(nil)).Elem(), DirectLineSpeechChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelInput)(nil)).Elem(), EmailChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelPropertiesInput)(nil)).Elem(), EmailChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelPropertiesPtrInput)(nil)).Elem(), EmailChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelPropertiesResponseInput)(nil)).Elem(), EmailChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelPropertiesResponsePtrInput)(nil)).Elem(), EmailChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*EmailChannelResponseInput)(nil)).Elem(), EmailChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookChannelInput)(nil)).Elem(), FacebookChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookChannelPropertiesInput)(nil)).Elem(), FacebookChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookChannelPropertiesPtrInput)(nil)).Elem(), FacebookChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookChannelPropertiesResponseInput)(nil)).Elem(), FacebookChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookChannelPropertiesResponsePtrInput)(nil)).Elem(), FacebookChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookChannelResponseInput)(nil)).Elem(), FacebookChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookPageInput)(nil)).Elem(), FacebookPageArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookPageArrayInput)(nil)).Elem(), FacebookPageArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookPageResponseInput)(nil)).Elem(), FacebookPageResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*FacebookPageResponseArrayInput)(nil)).Elem(), FacebookPageResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*KikChannelInput)(nil)).Elem(), KikChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KikChannelPropertiesInput)(nil)).Elem(), KikChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KikChannelPropertiesPtrInput)(nil)).Elem(), KikChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KikChannelPropertiesResponseInput)(nil)).Elem(), KikChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KikChannelPropertiesResponsePtrInput)(nil)).Elem(), KikChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*KikChannelResponseInput)(nil)).Elem(), KikChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineChannelInput)(nil)).Elem(), LineChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineChannelPropertiesInput)(nil)).Elem(), LineChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineChannelPropertiesPtrInput)(nil)).Elem(), LineChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineChannelPropertiesResponseInput)(nil)).Elem(), LineChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineChannelPropertiesResponsePtrInput)(nil)).Elem(), LineChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineChannelResponseInput)(nil)).Elem(), LineChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineRegistrationInput)(nil)).Elem(), LineRegistrationArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineRegistrationArrayInput)(nil)).Elem(), LineRegistrationArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineRegistrationResponseInput)(nil)).Elem(), LineRegistrationResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*LineRegistrationResponseArrayInput)(nil)).Elem(), LineRegistrationResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsTeamsChannelInput)(nil)).Elem(), MsTeamsChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsTeamsChannelPropertiesInput)(nil)).Elem(), MsTeamsChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsTeamsChannelPropertiesPtrInput)(nil)).Elem(), MsTeamsChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsTeamsChannelPropertiesResponseInput)(nil)).Elem(), MsTeamsChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsTeamsChannelPropertiesResponsePtrInput)(nil)).Elem(), MsTeamsChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*MsTeamsChannelResponseInput)(nil)).Elem(), MsTeamsChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceProviderParameterResponseInput)(nil)).Elem(), ServiceProviderParameterResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceProviderParameterResponseArrayInput)(nil)).Elem(), ServiceProviderParameterResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceProviderPropertiesResponseInput)(nil)).Elem(), ServiceProviderPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceProviderPropertiesResponsePtrInput)(nil)).Elem(), ServiceProviderPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceProviderResponseInput)(nil)).Elem(), ServiceProviderResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*ServiceProviderResponseArrayInput)(nil)).Elem(), ServiceProviderResponseArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuPtrInput)(nil)).Elem(), SkuArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponseInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkuResponsePtrInput)(nil)).Elem(), SkuResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkypeChannelInput)(nil)).Elem(), SkypeChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkypeChannelPropertiesInput)(nil)).Elem(), SkypeChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkypeChannelPropertiesPtrInput)(nil)).Elem(), SkypeChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkypeChannelPropertiesResponseInput)(nil)).Elem(), SkypeChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkypeChannelPropertiesResponsePtrInput)(nil)).Elem(), SkypeChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SkypeChannelResponseInput)(nil)).Elem(), SkypeChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelInput)(nil)).Elem(), SlackChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelPropertiesInput)(nil)).Elem(), SlackChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelPropertiesPtrInput)(nil)).Elem(), SlackChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelPropertiesResponseInput)(nil)).Elem(), SlackChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelPropertiesResponsePtrInput)(nil)).Elem(), SlackChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SlackChannelResponseInput)(nil)).Elem(), SlackChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelInput)(nil)).Elem(), SmsChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelPropertiesInput)(nil)).Elem(), SmsChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelPropertiesPtrInput)(nil)).Elem(), SmsChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelPropertiesResponseInput)(nil)).Elem(), SmsChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelPropertiesResponsePtrInput)(nil)).Elem(), SmsChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*SmsChannelResponseInput)(nil)).Elem(), SmsChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TelegramChannelInput)(nil)).Elem(), TelegramChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TelegramChannelPropertiesInput)(nil)).Elem(), TelegramChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TelegramChannelPropertiesPtrInput)(nil)).Elem(), TelegramChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TelegramChannelPropertiesResponseInput)(nil)).Elem(), TelegramChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TelegramChannelPropertiesResponsePtrInput)(nil)).Elem(), TelegramChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*TelegramChannelResponseInput)(nil)).Elem(), TelegramChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatChannelInput)(nil)).Elem(), WebChatChannelArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatChannelPropertiesInput)(nil)).Elem(), WebChatChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatChannelPropertiesPtrInput)(nil)).Elem(), WebChatChannelPropertiesArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatChannelPropertiesResponseInput)(nil)).Elem(), WebChatChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatChannelPropertiesResponsePtrInput)(nil)).Elem(), WebChatChannelPropertiesResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatChannelResponseInput)(nil)).Elem(), WebChatChannelResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatSiteInput)(nil)).Elem(), WebChatSiteArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatSiteArrayInput)(nil)).Elem(), WebChatSiteArray{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatSiteResponseInput)(nil)).Elem(), WebChatSiteResponseArgs{})
	pulumi.RegisterInputType(reflect.TypeOf((*WebChatSiteResponseArrayInput)(nil)).Elem(), WebChatSiteResponseArray{})
	pulumi.RegisterOutputType(AlexaChannelOutput{})
	pulumi.RegisterOutputType(AlexaChannelPropertiesOutput{})
	pulumi.RegisterOutputType(AlexaChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(AlexaChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(AlexaChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(AlexaChannelResponseOutput{})
	pulumi.RegisterOutputType(BotPropertiesOutput{})
	pulumi.RegisterOutputType(BotPropertiesPtrOutput{})
	pulumi.RegisterOutputType(BotPropertiesResponseOutput{})
	pulumi.RegisterOutputType(BotPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterArrayOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterResponseOutput{})
	pulumi.RegisterOutputType(ConnectionSettingParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesPtrOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ConnectionSettingPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DirectLineChannelOutput{})
	pulumi.RegisterOutputType(DirectLineChannelPropertiesOutput{})
	pulumi.RegisterOutputType(DirectLineChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DirectLineChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DirectLineChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DirectLineChannelResponseOutput{})
	pulumi.RegisterOutputType(DirectLineSiteOutput{})
	pulumi.RegisterOutputType(DirectLineSiteArrayOutput{})
	pulumi.RegisterOutputType(DirectLineSiteResponseOutput{})
	pulumi.RegisterOutputType(DirectLineSiteResponseArrayOutput{})
	pulumi.RegisterOutputType(DirectLineSpeechChannelOutput{})
	pulumi.RegisterOutputType(DirectLineSpeechChannelPropertiesOutput{})
	pulumi.RegisterOutputType(DirectLineSpeechChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(DirectLineSpeechChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(DirectLineSpeechChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(DirectLineSpeechChannelResponseOutput{})
	pulumi.RegisterOutputType(EmailChannelOutput{})
	pulumi.RegisterOutputType(EmailChannelPropertiesOutput{})
	pulumi.RegisterOutputType(EmailChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(EmailChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(EmailChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(EmailChannelResponseOutput{})
	pulumi.RegisterOutputType(FacebookChannelOutput{})
	pulumi.RegisterOutputType(FacebookChannelPropertiesOutput{})
	pulumi.RegisterOutputType(FacebookChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(FacebookChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(FacebookChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(FacebookChannelResponseOutput{})
	pulumi.RegisterOutputType(FacebookPageOutput{})
	pulumi.RegisterOutputType(FacebookPageArrayOutput{})
	pulumi.RegisterOutputType(FacebookPageResponseOutput{})
	pulumi.RegisterOutputType(FacebookPageResponseArrayOutput{})
	pulumi.RegisterOutputType(KikChannelOutput{})
	pulumi.RegisterOutputType(KikChannelPropertiesOutput{})
	pulumi.RegisterOutputType(KikChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(KikChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(KikChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(KikChannelResponseOutput{})
	pulumi.RegisterOutputType(LineChannelOutput{})
	pulumi.RegisterOutputType(LineChannelPropertiesOutput{})
	pulumi.RegisterOutputType(LineChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(LineChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(LineChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(LineChannelResponseOutput{})
	pulumi.RegisterOutputType(LineRegistrationOutput{})
	pulumi.RegisterOutputType(LineRegistrationArrayOutput{})
	pulumi.RegisterOutputType(LineRegistrationResponseOutput{})
	pulumi.RegisterOutputType(LineRegistrationResponseArrayOutput{})
	pulumi.RegisterOutputType(MsTeamsChannelOutput{})
	pulumi.RegisterOutputType(MsTeamsChannelPropertiesOutput{})
	pulumi.RegisterOutputType(MsTeamsChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(MsTeamsChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(MsTeamsChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(MsTeamsChannelResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderParameterResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderParameterResponseArrayOutput{})
	pulumi.RegisterOutputType(ServiceProviderPropertiesResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(ServiceProviderResponseOutput{})
	pulumi.RegisterOutputType(ServiceProviderResponseArrayOutput{})
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
	pulumi.RegisterOutputType(SkypeChannelOutput{})
	pulumi.RegisterOutputType(SkypeChannelPropertiesOutput{})
	pulumi.RegisterOutputType(SkypeChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SkypeChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SkypeChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SkypeChannelResponseOutput{})
	pulumi.RegisterOutputType(SlackChannelOutput{})
	pulumi.RegisterOutputType(SlackChannelPropertiesOutput{})
	pulumi.RegisterOutputType(SlackChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SlackChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SlackChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SlackChannelResponseOutput{})
	pulumi.RegisterOutputType(SmsChannelOutput{})
	pulumi.RegisterOutputType(SmsChannelPropertiesOutput{})
	pulumi.RegisterOutputType(SmsChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(SmsChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(SmsChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(SmsChannelResponseOutput{})
	pulumi.RegisterOutputType(TelegramChannelOutput{})
	pulumi.RegisterOutputType(TelegramChannelPropertiesOutput{})
	pulumi.RegisterOutputType(TelegramChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(TelegramChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(TelegramChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(TelegramChannelResponseOutput{})
	pulumi.RegisterOutputType(WebChatChannelOutput{})
	pulumi.RegisterOutputType(WebChatChannelPropertiesOutput{})
	pulumi.RegisterOutputType(WebChatChannelPropertiesPtrOutput{})
	pulumi.RegisterOutputType(WebChatChannelPropertiesResponseOutput{})
	pulumi.RegisterOutputType(WebChatChannelPropertiesResponsePtrOutput{})
	pulumi.RegisterOutputType(WebChatChannelResponseOutput{})
	pulumi.RegisterOutputType(WebChatSiteOutput{})
	pulumi.RegisterOutputType(WebChatSiteArrayOutput{})
	pulumi.RegisterOutputType(WebChatSiteResponseOutput{})
	pulumi.RegisterOutputType(WebChatSiteResponseArrayOutput{})
}
