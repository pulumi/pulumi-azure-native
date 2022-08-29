


package v20171201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type BotProperties struct {
	Description                       *string  `pulumi:"description"`
	DeveloperAppInsightKey            *string  `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        *string  `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId *string  `pulumi:"developerAppInsightsApplicationId"`
	DisplayName                       string   `pulumi:"displayName"`
	Endpoint                          string   `pulumi:"endpoint"`
	IconUrl                           *string  `pulumi:"iconUrl"`
	LuisAppIds                        []string `pulumi:"luisAppIds"`
	LuisKey                           *string  `pulumi:"luisKey"`
	MsaAppId                          string   `pulumi:"msaAppId"`
}





type BotPropertiesInput interface {
	pulumi.Input

	ToBotPropertiesOutput() BotPropertiesOutput
	ToBotPropertiesOutputWithContext(context.Context) BotPropertiesOutput
}

type BotPropertiesArgs struct {
	Description                       pulumi.StringPtrInput   `pulumi:"description"`
	DeveloperAppInsightKey            pulumi.StringPtrInput   `pulumi:"developerAppInsightKey"`
	DeveloperAppInsightsApiKey        pulumi.StringPtrInput   `pulumi:"developerAppInsightsApiKey"`
	DeveloperAppInsightsApplicationId pulumi.StringPtrInput   `pulumi:"developerAppInsightsApplicationId"`
	DisplayName                       pulumi.StringInput      `pulumi:"displayName"`
	Endpoint                          pulumi.StringInput      `pulumi:"endpoint"`
	IconUrl                           pulumi.StringPtrInput   `pulumi:"iconUrl"`
	LuisAppIds                        pulumi.StringArrayInput `pulumi:"luisAppIds"`
	LuisKey                           pulumi.StringPtrInput   `pulumi:"luisKey"`
	MsaAppId                          pulumi.StringInput      `pulumi:"msaAppId"`
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

func (o BotPropertiesOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotProperties) []string { return v.LuisAppIds }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotProperties) *string { return v.LuisKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesOutput) MsaAppId() pulumi.StringOutput {
	return o.ApplyT(func(v BotProperties) string { return v.MsaAppId }).(pulumi.StringOutput)
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

type BotPropertiesResponse struct {
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
	LuisAppIds                        []string `pulumi:"luisAppIds"`
	LuisKey                           *string  `pulumi:"luisKey"`
	MsaAppId                          string   `pulumi:"msaAppId"`
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

func (o BotPropertiesResponseOutput) LuisAppIds() pulumi.StringArrayOutput {
	return o.ApplyT(func(v BotPropertiesResponse) []string { return v.LuisAppIds }).(pulumi.StringArrayOutput)
}

func (o BotPropertiesResponseOutput) LuisKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v BotPropertiesResponse) *string { return v.LuisKey }).(pulumi.StringPtrOutput)
}

func (o BotPropertiesResponseOutput) MsaAppId() pulumi.StringOutput {
	return o.ApplyT(func(v BotPropertiesResponse) string { return v.MsaAppId }).(pulumi.StringOutput)
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
	Properties  *DirectLineChannelProperties `pulumi:"properties"`
}

type DirectLineChannelProperties struct {
	Sites []DirectLineSite `pulumi:"sites"`
}

type DirectLineChannelPropertiesResponse struct {
	Sites []DirectLineSiteResponse `pulumi:"sites"`
}

type DirectLineChannelResponse struct {
	ChannelName string                               `pulumi:"channelName"`
	Properties  *DirectLineChannelPropertiesResponse `pulumi:"properties"`
}

type DirectLineSite struct {
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsV1Enabled bool   `pulumi:"isV1Enabled"`
	IsV3Enabled bool   `pulumi:"isV3Enabled"`
	SiteName    string `pulumi:"siteName"`
}

type DirectLineSiteResponse struct {
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsV1Enabled bool   `pulumi:"isV1Enabled"`
	IsV3Enabled bool   `pulumi:"isV3Enabled"`
	Key         string `pulumi:"key"`
	Key2        string `pulumi:"key2"`
	SiteId      string `pulumi:"siteId"`
	SiteName    string `pulumi:"siteName"`
}

type EmailChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Properties  *EmailChannelProperties `pulumi:"properties"`
}

type EmailChannelProperties struct {
	EmailAddress string `pulumi:"emailAddress"`
	IsEnabled    bool   `pulumi:"isEnabled"`
	Password     string `pulumi:"password"`
}

type EmailChannelPropertiesResponse struct {
	EmailAddress string `pulumi:"emailAddress"`
	IsEnabled    bool   `pulumi:"isEnabled"`
	Password     string `pulumi:"password"`
}

type EmailChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *EmailChannelPropertiesResponse `pulumi:"properties"`
}

type FacebookChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Properties  *FacebookChannelProperties `pulumi:"properties"`
}

type FacebookChannelProperties struct {
	AppId     string         `pulumi:"appId"`
	AppSecret string         `pulumi:"appSecret"`
	IsEnabled bool           `pulumi:"isEnabled"`
	Pages     []FacebookPage `pulumi:"pages"`
}

type FacebookChannelPropertiesResponse struct {
	AppId       string                 `pulumi:"appId"`
	AppSecret   string                 `pulumi:"appSecret"`
	CallbackUrl string                 `pulumi:"callbackUrl"`
	IsEnabled   bool                   `pulumi:"isEnabled"`
	Pages       []FacebookPageResponse `pulumi:"pages"`
	VerifyToken string                 `pulumi:"verifyToken"`
}

type FacebookChannelResponse struct {
	ChannelName string                             `pulumi:"channelName"`
	Properties  *FacebookChannelPropertiesResponse `pulumi:"properties"`
}

type FacebookPage struct {
	AccessToken string `pulumi:"accessToken"`
	Id          string `pulumi:"id"`
}

type FacebookPageResponse struct {
	AccessToken string `pulumi:"accessToken"`
	Id          string `pulumi:"id"`
}

type KikChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Properties  *KikChannelProperties `pulumi:"properties"`
}

type KikChannelProperties struct {
	ApiKey      string `pulumi:"apiKey"`
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsValidated *bool  `pulumi:"isValidated"`
	UserName    string `pulumi:"userName"`
}

type KikChannelPropertiesResponse struct {
	ApiKey      string `pulumi:"apiKey"`
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsValidated *bool  `pulumi:"isValidated"`
	UserName    string `pulumi:"userName"`
}

type KikChannelResponse struct {
	ChannelName string                        `pulumi:"channelName"`
	Properties  *KikChannelPropertiesResponse `pulumi:"properties"`
}

type MsTeamsChannel struct {
	ChannelName string                    `pulumi:"channelName"`
	Properties  *MsTeamsChannelProperties `pulumi:"properties"`
}

type MsTeamsChannelProperties struct {
	CallMode         *string `pulumi:"callMode"`
	EnableCalling    *bool   `pulumi:"enableCalling"`
	EnableMediaCards *bool   `pulumi:"enableMediaCards"`
	EnableMessaging  *bool   `pulumi:"enableMessaging"`
	EnableVideo      *bool   `pulumi:"enableVideo"`
	IsEnabled        bool    `pulumi:"isEnabled"`
}

type MsTeamsChannelPropertiesResponse struct {
	CallMode         *string `pulumi:"callMode"`
	EnableCalling    *bool   `pulumi:"enableCalling"`
	EnableMediaCards *bool   `pulumi:"enableMediaCards"`
	EnableMessaging  *bool   `pulumi:"enableMessaging"`
	EnableVideo      *bool   `pulumi:"enableVideo"`
	IsEnabled        bool    `pulumi:"isEnabled"`
}

type MsTeamsChannelResponse struct {
	ChannelName string                            `pulumi:"channelName"`
	Properties  *MsTeamsChannelPropertiesResponse `pulumi:"properties"`
}

type ServiceProviderParameterResponse struct {
	Default     string `pulumi:"default"`
	Description string `pulumi:"description"`
	DisplayName string `pulumi:"displayName"`
	HelpUrl     string `pulumi:"helpUrl"`
	Name        string `pulumi:"name"`
	Type        string `pulumi:"type"`
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
	IsEnabled           bool    `pulumi:"isEnabled"`
}

type SkypeChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *SkypeChannelPropertiesResponse `pulumi:"properties"`
}

type SlackChannel struct {
	ChannelName string                  `pulumi:"channelName"`
	Properties  *SlackChannelProperties `pulumi:"properties"`
}

type SlackChannelProperties struct {
	ClientId          string  `pulumi:"clientId"`
	ClientSecret      string  `pulumi:"clientSecret"`
	IsEnabled         bool    `pulumi:"isEnabled"`
	LandingPageUrl    *string `pulumi:"landingPageUrl"`
	VerificationToken string  `pulumi:"verificationToken"`
}

type SlackChannelPropertiesResponse struct {
	ClientId                string  `pulumi:"clientId"`
	ClientSecret            string  `pulumi:"clientSecret"`
	IsEnabled               bool    `pulumi:"isEnabled"`
	IsValidated             bool    `pulumi:"isValidated"`
	LandingPageUrl          *string `pulumi:"landingPageUrl"`
	LastSubmissionId        string  `pulumi:"lastSubmissionId"`
	RedirectAction          string  `pulumi:"redirectAction"`
	RegisterBeforeOAuthFlow bool    `pulumi:"registerBeforeOAuthFlow"`
	VerificationToken       string  `pulumi:"verificationToken"`
}

type SlackChannelResponse struct {
	ChannelName string                          `pulumi:"channelName"`
	Properties  *SlackChannelPropertiesResponse `pulumi:"properties"`
}

type SmsChannel struct {
	ChannelName string                `pulumi:"channelName"`
	Properties  *SmsChannelProperties `pulumi:"properties"`
}

type SmsChannelProperties struct {
	AccountSID  string `pulumi:"accountSID"`
	AuthToken   string `pulumi:"authToken"`
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsValidated *bool  `pulumi:"isValidated"`
	Phone       string `pulumi:"phone"`
}

type SmsChannelPropertiesResponse struct {
	AccountSID  string `pulumi:"accountSID"`
	AuthToken   string `pulumi:"authToken"`
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsValidated *bool  `pulumi:"isValidated"`
	Phone       string `pulumi:"phone"`
}

type SmsChannelResponse struct {
	ChannelName string                        `pulumi:"channelName"`
	Properties  *SmsChannelPropertiesResponse `pulumi:"properties"`
}

type TelegramChannel struct {
	ChannelName string                     `pulumi:"channelName"`
	Properties  *TelegramChannelProperties `pulumi:"properties"`
}

type TelegramChannelProperties struct {
	AccessToken string `pulumi:"accessToken"`
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsValidated *bool  `pulumi:"isValidated"`
}

type TelegramChannelPropertiesResponse struct {
	AccessToken string `pulumi:"accessToken"`
	IsEnabled   bool   `pulumi:"isEnabled"`
	IsValidated *bool  `pulumi:"isValidated"`
}

type TelegramChannelResponse struct {
	ChannelName string                             `pulumi:"channelName"`
	Properties  *TelegramChannelPropertiesResponse `pulumi:"properties"`
}

type WebChatChannel struct {
	ChannelName string                    `pulumi:"channelName"`
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
	ChannelName string                            `pulumi:"channelName"`
	Properties  *WebChatChannelPropertiesResponse `pulumi:"properties"`
}

type WebChatSite struct {
	EnablePreview bool   `pulumi:"enablePreview"`
	IsEnabled     bool   `pulumi:"isEnabled"`
	SiteName      string `pulumi:"siteName"`
}

type WebChatSiteResponse struct {
	EnablePreview bool   `pulumi:"enablePreview"`
	IsEnabled     bool   `pulumi:"isEnabled"`
	Key           string `pulumi:"key"`
	Key2          string `pulumi:"key2"`
	SiteId        string `pulumi:"siteId"`
	SiteName      string `pulumi:"siteName"`
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
	pulumi.RegisterOutputType(SkuOutput{})
	pulumi.RegisterOutputType(SkuPtrOutput{})
	pulumi.RegisterOutputType(SkuResponseOutput{})
	pulumi.RegisterOutputType(SkuResponsePtrOutput{})
}
