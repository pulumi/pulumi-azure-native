


package v20201201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAuthSettingsSlot struct {
	pulumi.CustomResourceState

	AadClaimsAuthorization                  pulumi.StringPtrOutput   `pulumi:"aadClaimsAuthorization"`
	AdditionalLoginParams                   pulumi.StringArrayOutput `pulumi:"additionalLoginParams"`
	AllowedAudiences                        pulumi.StringArrayOutput `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls             pulumi.StringArrayOutput `pulumi:"allowedExternalRedirectUrls"`
	AuthFilePath                            pulumi.StringPtrOutput   `pulumi:"authFilePath"`
	ClientId                                pulumi.StringPtrOutput   `pulumi:"clientId"`
	ClientSecret                            pulumi.StringPtrOutput   `pulumi:"clientSecret"`
	ClientSecretCertificateThumbprint       pulumi.StringPtrOutput   `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                 pulumi.StringPtrOutput   `pulumi:"clientSecretSettingName"`
	ConfigVersion                           pulumi.StringPtrOutput   `pulumi:"configVersion"`
	DefaultProvider                         pulumi.StringPtrOutput   `pulumi:"defaultProvider"`
	Enabled                                 pulumi.BoolPtrOutput     `pulumi:"enabled"`
	FacebookAppId                           pulumi.StringPtrOutput   `pulumi:"facebookAppId"`
	FacebookAppSecret                       pulumi.StringPtrOutput   `pulumi:"facebookAppSecret"`
	FacebookAppSecretSettingName            pulumi.StringPtrOutput   `pulumi:"facebookAppSecretSettingName"`
	FacebookOAuthScopes                     pulumi.StringArrayOutput `pulumi:"facebookOAuthScopes"`
	GitHubClientId                          pulumi.StringPtrOutput   `pulumi:"gitHubClientId"`
	GitHubClientSecret                      pulumi.StringPtrOutput   `pulumi:"gitHubClientSecret"`
	GitHubClientSecretSettingName           pulumi.StringPtrOutput   `pulumi:"gitHubClientSecretSettingName"`
	GitHubOAuthScopes                       pulumi.StringArrayOutput `pulumi:"gitHubOAuthScopes"`
	GoogleClientId                          pulumi.StringPtrOutput   `pulumi:"googleClientId"`
	GoogleClientSecret                      pulumi.StringPtrOutput   `pulumi:"googleClientSecret"`
	GoogleClientSecretSettingName           pulumi.StringPtrOutput   `pulumi:"googleClientSecretSettingName"`
	GoogleOAuthScopes                       pulumi.StringArrayOutput `pulumi:"googleOAuthScopes"`
	IsAuthFromFile                          pulumi.StringPtrOutput   `pulumi:"isAuthFromFile"`
	Issuer                                  pulumi.StringPtrOutput   `pulumi:"issuer"`
	Kind                                    pulumi.StringPtrOutput   `pulumi:"kind"`
	MicrosoftAccountClientId                pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret            pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountClientSecretSettingName pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientSecretSettingName"`
	MicrosoftAccountOAuthScopes             pulumi.StringArrayOutput `pulumi:"microsoftAccountOAuthScopes"`
	Name                                    pulumi.StringOutput      `pulumi:"name"`
	RuntimeVersion                          pulumi.StringPtrOutput   `pulumi:"runtimeVersion"`
	TokenRefreshExtensionHours              pulumi.Float64PtrOutput  `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                       pulumi.BoolPtrOutput     `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                      pulumi.StringPtrOutput   `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret                   pulumi.StringPtrOutput   `pulumi:"twitterConsumerSecret"`
	TwitterConsumerSecretSettingName        pulumi.StringPtrOutput   `pulumi:"twitterConsumerSecretSettingName"`
	Type                                    pulumi.StringOutput      `pulumi:"type"`
	UnauthenticatedClientAction             pulumi.StringPtrOutput   `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                          pulumi.BoolPtrOutput     `pulumi:"validateIssuer"`
}


func NewWebAppAuthSettingsSlot(ctx *pulumi.Context,
	name string, args *WebAppAuthSettingsSlotArgs, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsSlot, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	if args.Slot == nil {
		return nil, errors.New("invalid value for required argument 'Slot'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:WebAppAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:WebAppAuthSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppAuthSettingsSlot
	err := ctx.RegisterResource("azure-native:web/v20201201:WebAppAuthSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppAuthSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAuthSettingsSlotState, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsSlot, error) {
	var resource WebAppAuthSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20201201:WebAppAuthSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppAuthSettingsSlotState struct {
}

type WebAppAuthSettingsSlotState struct {
}

func (WebAppAuthSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsSlotState)(nil)).Elem()
}

type webAppAuthSettingsSlotArgs struct {
	AadClaimsAuthorization                  *string                        `pulumi:"aadClaimsAuthorization"`
	AdditionalLoginParams                   []string                       `pulumi:"additionalLoginParams"`
	AllowedAudiences                        []string                       `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls             []string                       `pulumi:"allowedExternalRedirectUrls"`
	AuthFilePath                            *string                        `pulumi:"authFilePath"`
	ClientId                                *string                        `pulumi:"clientId"`
	ClientSecret                            *string                        `pulumi:"clientSecret"`
	ClientSecretCertificateThumbprint       *string                        `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                 *string                        `pulumi:"clientSecretSettingName"`
	ConfigVersion                           *string                        `pulumi:"configVersion"`
	DefaultProvider                         *BuiltInAuthenticationProvider `pulumi:"defaultProvider"`
	Enabled                                 *bool                          `pulumi:"enabled"`
	FacebookAppId                           *string                        `pulumi:"facebookAppId"`
	FacebookAppSecret                       *string                        `pulumi:"facebookAppSecret"`
	FacebookAppSecretSettingName            *string                        `pulumi:"facebookAppSecretSettingName"`
	FacebookOAuthScopes                     []string                       `pulumi:"facebookOAuthScopes"`
	GitHubClientId                          *string                        `pulumi:"gitHubClientId"`
	GitHubClientSecret                      *string                        `pulumi:"gitHubClientSecret"`
	GitHubClientSecretSettingName           *string                        `pulumi:"gitHubClientSecretSettingName"`
	GitHubOAuthScopes                       []string                       `pulumi:"gitHubOAuthScopes"`
	GoogleClientId                          *string                        `pulumi:"googleClientId"`
	GoogleClientSecret                      *string                        `pulumi:"googleClientSecret"`
	GoogleClientSecretSettingName           *string                        `pulumi:"googleClientSecretSettingName"`
	GoogleOAuthScopes                       []string                       `pulumi:"googleOAuthScopes"`
	IsAuthFromFile                          *string                        `pulumi:"isAuthFromFile"`
	Issuer                                  *string                        `pulumi:"issuer"`
	Kind                                    *string                        `pulumi:"kind"`
	MicrosoftAccountClientId                *string                        `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret            *string                        `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountClientSecretSettingName *string                        `pulumi:"microsoftAccountClientSecretSettingName"`
	MicrosoftAccountOAuthScopes             []string                       `pulumi:"microsoftAccountOAuthScopes"`
	Name                                    string                         `pulumi:"name"`
	ResourceGroupName                       string                         `pulumi:"resourceGroupName"`
	RuntimeVersion                          *string                        `pulumi:"runtimeVersion"`
	Slot                                    string                         `pulumi:"slot"`
	TokenRefreshExtensionHours              *float64                       `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                       *bool                          `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                      *string                        `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret                   *string                        `pulumi:"twitterConsumerSecret"`
	TwitterConsumerSecretSettingName        *string                        `pulumi:"twitterConsumerSecretSettingName"`
	UnauthenticatedClientAction             *UnauthenticatedClientAction   `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                          *bool                          `pulumi:"validateIssuer"`
}


type WebAppAuthSettingsSlotArgs struct {
	AadClaimsAuthorization                  pulumi.StringPtrInput
	AdditionalLoginParams                   pulumi.StringArrayInput
	AllowedAudiences                        pulumi.StringArrayInput
	AllowedExternalRedirectUrls             pulumi.StringArrayInput
	AuthFilePath                            pulumi.StringPtrInput
	ClientId                                pulumi.StringPtrInput
	ClientSecret                            pulumi.StringPtrInput
	ClientSecretCertificateThumbprint       pulumi.StringPtrInput
	ClientSecretSettingName                 pulumi.StringPtrInput
	ConfigVersion                           pulumi.StringPtrInput
	DefaultProvider                         BuiltInAuthenticationProviderPtrInput
	Enabled                                 pulumi.BoolPtrInput
	FacebookAppId                           pulumi.StringPtrInput
	FacebookAppSecret                       pulumi.StringPtrInput
	FacebookAppSecretSettingName            pulumi.StringPtrInput
	FacebookOAuthScopes                     pulumi.StringArrayInput
	GitHubClientId                          pulumi.StringPtrInput
	GitHubClientSecret                      pulumi.StringPtrInput
	GitHubClientSecretSettingName           pulumi.StringPtrInput
	GitHubOAuthScopes                       pulumi.StringArrayInput
	GoogleClientId                          pulumi.StringPtrInput
	GoogleClientSecret                      pulumi.StringPtrInput
	GoogleClientSecretSettingName           pulumi.StringPtrInput
	GoogleOAuthScopes                       pulumi.StringArrayInput
	IsAuthFromFile                          pulumi.StringPtrInput
	Issuer                                  pulumi.StringPtrInput
	Kind                                    pulumi.StringPtrInput
	MicrosoftAccountClientId                pulumi.StringPtrInput
	MicrosoftAccountClientSecret            pulumi.StringPtrInput
	MicrosoftAccountClientSecretSettingName pulumi.StringPtrInput
	MicrosoftAccountOAuthScopes             pulumi.StringArrayInput
	Name                                    pulumi.StringInput
	ResourceGroupName                       pulumi.StringInput
	RuntimeVersion                          pulumi.StringPtrInput
	Slot                                    pulumi.StringInput
	TokenRefreshExtensionHours              pulumi.Float64PtrInput
	TokenStoreEnabled                       pulumi.BoolPtrInput
	TwitterConsumerKey                      pulumi.StringPtrInput
	TwitterConsumerSecret                   pulumi.StringPtrInput
	TwitterConsumerSecretSettingName        pulumi.StringPtrInput
	UnauthenticatedClientAction             UnauthenticatedClientActionPtrInput
	ValidateIssuer                          pulumi.BoolPtrInput
}

func (WebAppAuthSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsSlotArgs)(nil)).Elem()
}

type WebAppAuthSettingsSlotInput interface {
	pulumi.Input

	ToWebAppAuthSettingsSlotOutput() WebAppAuthSettingsSlotOutput
	ToWebAppAuthSettingsSlotOutputWithContext(ctx context.Context) WebAppAuthSettingsSlotOutput
}

func (*WebAppAuthSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsSlot)(nil)).Elem()
}

func (i *WebAppAuthSettingsSlot) ToWebAppAuthSettingsSlotOutput() WebAppAuthSettingsSlotOutput {
	return i.ToWebAppAuthSettingsSlotOutputWithContext(context.Background())
}

func (i *WebAppAuthSettingsSlot) ToWebAppAuthSettingsSlotOutputWithContext(ctx context.Context) WebAppAuthSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAuthSettingsSlotOutput)
}

type WebAppAuthSettingsSlotOutput struct{ *pulumi.OutputState }

func (WebAppAuthSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**WebAppAuthSettingsSlot)(nil)).Elem()
}

func (o WebAppAuthSettingsSlotOutput) ToWebAppAuthSettingsSlotOutput() WebAppAuthSettingsSlotOutput {
	return o
}

func (o WebAppAuthSettingsSlotOutput) ToWebAppAuthSettingsSlotOutputWithContext(ctx context.Context) WebAppAuthSettingsSlotOutput {
	return o
}

func (o WebAppAuthSettingsSlotOutput) AadClaimsAuthorization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.AadClaimsAuthorization }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) AuthFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.AuthFilePath }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) ConfigVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.ConfigVersion }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) FacebookAppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.FacebookAppSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) GitHubClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GitHubClientId }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GitHubClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GitHubClientSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GitHubClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GitHubClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GitHubOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.GitHubOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) IsAuthFromFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.IsAuthFromFile }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) MicrosoftAccountClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput {
		return v.MicrosoftAccountClientSecretSettingName
	}).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringOutput { return v.Name }).(pulumi.StringOutput)
}

func (o WebAppAuthSettingsSlotOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.Float64PtrOutput { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.BoolPtrOutput { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) TwitterConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.TwitterConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringOutput { return v.Type }).(pulumi.StringOutput)
}

func (o WebAppAuthSettingsSlotOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) ValidateIssuer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.BoolPtrOutput { return v.ValidateIssuer }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(WebAppAuthSettingsSlotOutput{})
}
