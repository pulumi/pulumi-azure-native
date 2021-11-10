


package v20210201

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAuthSettings struct {
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


func NewWebAppAuthSettings(ctx *pulumi.Context,
	name string, args *WebAppAuthSettingsArgs, opts ...pulumi.ResourceOption) (*WebAppAuthSettings, error) {
	if args == nil {
		return nil, errors.New("missing one or more required arguments")
	}

	if args.Name == nil {
		return nil, errors.New("invalid value for required argument 'Name'")
	}
	if args.ResourceGroupName == nil {
		return nil, errors.New("invalid value for required argument 'ResourceGroupName'")
	}
	aliases := pulumi.Aliases([]pulumi.Alias{
		{
			Type: pulumi.String("azure-native:web:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20150801:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:WebAppAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:WebAppAuthSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource WebAppAuthSettings
	err := ctx.RegisterResource("azure-native:web/v20210201:WebAppAuthSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppAuthSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAuthSettingsState, opts ...pulumi.ResourceOption) (*WebAppAuthSettings, error) {
	var resource WebAppAuthSettings
	err := ctx.ReadResource("azure-native:web/v20210201:WebAppAuthSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type webAppAuthSettingsState struct {
}

type WebAppAuthSettingsState struct {
}

func (WebAppAuthSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsState)(nil)).Elem()
}

type webAppAuthSettingsArgs struct {
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
	TokenRefreshExtensionHours              *float64                       `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                       *bool                          `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                      *string                        `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret                   *string                        `pulumi:"twitterConsumerSecret"`
	TwitterConsumerSecretSettingName        *string                        `pulumi:"twitterConsumerSecretSettingName"`
	UnauthenticatedClientAction             *UnauthenticatedClientAction   `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                          *bool                          `pulumi:"validateIssuer"`
}


type WebAppAuthSettingsArgs struct {
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
	TokenRefreshExtensionHours              pulumi.Float64PtrInput
	TokenStoreEnabled                       pulumi.BoolPtrInput
	TwitterConsumerKey                      pulumi.StringPtrInput
	TwitterConsumerSecret                   pulumi.StringPtrInput
	TwitterConsumerSecretSettingName        pulumi.StringPtrInput
	UnauthenticatedClientAction             UnauthenticatedClientActionPtrInput
	ValidateIssuer                          pulumi.BoolPtrInput
}

func (WebAppAuthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*webAppAuthSettingsArgs)(nil)).Elem()
}

type WebAppAuthSettingsInput interface {
	pulumi.Input

	ToWebAppAuthSettingsOutput() WebAppAuthSettingsOutput
	ToWebAppAuthSettingsOutputWithContext(ctx context.Context) WebAppAuthSettingsOutput
}

func (*WebAppAuthSettings) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppAuthSettings)(nil))
}

func (i *WebAppAuthSettings) ToWebAppAuthSettingsOutput() WebAppAuthSettingsOutput {
	return i.ToWebAppAuthSettingsOutputWithContext(context.Background())
}

func (i *WebAppAuthSettings) ToWebAppAuthSettingsOutputWithContext(ctx context.Context) WebAppAuthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(WebAppAuthSettingsOutput)
}

type WebAppAuthSettingsOutput struct{ *pulumi.OutputState }

func (WebAppAuthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*WebAppAuthSettings)(nil))
}

func (o WebAppAuthSettingsOutput) ToWebAppAuthSettingsOutput() WebAppAuthSettingsOutput {
	return o
}

func (o WebAppAuthSettingsOutput) ToWebAppAuthSettingsOutputWithContext(ctx context.Context) WebAppAuthSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(WebAppAuthSettingsOutput{})
}
