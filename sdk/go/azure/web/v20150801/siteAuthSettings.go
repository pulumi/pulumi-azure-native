


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type SiteAuthSettings struct {
	pulumi.CustomResourceState

	AadClientId                  pulumi.StringPtrOutput   `pulumi:"aadClientId"`
	AdditionalLoginParams        pulumi.StringArrayOutput `pulumi:"additionalLoginParams"`
	AllowedAudiences             pulumi.StringArrayOutput `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls  pulumi.StringArrayOutput `pulumi:"allowedExternalRedirectUrls"`
	ClientId                     pulumi.StringPtrOutput   `pulumi:"clientId"`
	ClientSecret                 pulumi.StringPtrOutput   `pulumi:"clientSecret"`
	DefaultProvider              pulumi.StringPtrOutput   `pulumi:"defaultProvider"`
	Enabled                      pulumi.BoolPtrOutput     `pulumi:"enabled"`
	FacebookAppId                pulumi.StringPtrOutput   `pulumi:"facebookAppId"`
	FacebookAppSecret            pulumi.StringPtrOutput   `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes          pulumi.StringArrayOutput `pulumi:"facebookOAuthScopes"`
	GoogleClientId               pulumi.StringPtrOutput   `pulumi:"googleClientId"`
	GoogleClientSecret           pulumi.StringPtrOutput   `pulumi:"googleClientSecret"`
	GoogleOAuthScopes            pulumi.StringArrayOutput `pulumi:"googleOAuthScopes"`
	HttpApiPrefixPath            pulumi.StringPtrOutput   `pulumi:"httpApiPrefixPath"`
	Issuer                       pulumi.StringPtrOutput   `pulumi:"issuer"`
	MicrosoftAccountClientId     pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes  pulumi.StringArrayOutput `pulumi:"microsoftAccountOAuthScopes"`
	OpenIdIssuer                 pulumi.StringPtrOutput   `pulumi:"openIdIssuer"`
	TokenRefreshExtensionHours   pulumi.Float64PtrOutput  `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled            pulumi.BoolPtrOutput     `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey           pulumi.StringPtrOutput   `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret        pulumi.StringPtrOutput   `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction  pulumi.StringPtrOutput   `pulumi:"unauthenticatedClientAction"`
}


func NewSiteAuthSettings(ctx *pulumi.Context,
	name string, args *SiteAuthSettingsArgs, opts ...pulumi.ResourceOption) (*SiteAuthSettings, error) {
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
			Type: pulumi.String("azure-native:web:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteAuthSettings"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteAuthSettings"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteAuthSettings
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteAuthSettings", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteAuthSettings(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteAuthSettingsState, opts ...pulumi.ResourceOption) (*SiteAuthSettings, error) {
	var resource SiteAuthSettings
	err := ctx.ReadResource("azure-native:web/v20150801:SiteAuthSettings", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteAuthSettingsState struct {
}

type SiteAuthSettingsState struct {
}

func (SiteAuthSettingsState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAuthSettingsState)(nil)).Elem()
}

type siteAuthSettingsArgs struct {
	AadClientId                  *string                        `pulumi:"aadClientId"`
	AdditionalLoginParams        []string                       `pulumi:"additionalLoginParams"`
	AllowedAudiences             []string                       `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls  []string                       `pulumi:"allowedExternalRedirectUrls"`
	ClientId                     *string                        `pulumi:"clientId"`
	ClientSecret                 *string                        `pulumi:"clientSecret"`
	DefaultProvider              *BuiltInAuthenticationProvider `pulumi:"defaultProvider"`
	Enabled                      *bool                          `pulumi:"enabled"`
	FacebookAppId                *string                        `pulumi:"facebookAppId"`
	FacebookAppSecret            *string                        `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes          []string                       `pulumi:"facebookOAuthScopes"`
	GoogleClientId               *string                        `pulumi:"googleClientId"`
	GoogleClientSecret           *string                        `pulumi:"googleClientSecret"`
	GoogleOAuthScopes            []string                       `pulumi:"googleOAuthScopes"`
	HttpApiPrefixPath            *string                        `pulumi:"httpApiPrefixPath"`
	Issuer                       *string                        `pulumi:"issuer"`
	MicrosoftAccountClientId     *string                        `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret *string                        `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes  []string                       `pulumi:"microsoftAccountOAuthScopes"`
	Name                         string                         `pulumi:"name"`
	OpenIdIssuer                 *string                        `pulumi:"openIdIssuer"`
	ResourceGroupName            string                         `pulumi:"resourceGroupName"`
	TokenRefreshExtensionHours   *float64                       `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled            *bool                          `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey           *string                        `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret        *string                        `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction  *UnauthenticatedClientAction   `pulumi:"unauthenticatedClientAction"`
}


type SiteAuthSettingsArgs struct {
	AadClientId                  pulumi.StringPtrInput
	AdditionalLoginParams        pulumi.StringArrayInput
	AllowedAudiences             pulumi.StringArrayInput
	AllowedExternalRedirectUrls  pulumi.StringArrayInput
	ClientId                     pulumi.StringPtrInput
	ClientSecret                 pulumi.StringPtrInput
	DefaultProvider              BuiltInAuthenticationProviderPtrInput
	Enabled                      pulumi.BoolPtrInput
	FacebookAppId                pulumi.StringPtrInput
	FacebookAppSecret            pulumi.StringPtrInput
	FacebookOAuthScopes          pulumi.StringArrayInput
	GoogleClientId               pulumi.StringPtrInput
	GoogleClientSecret           pulumi.StringPtrInput
	GoogleOAuthScopes            pulumi.StringArrayInput
	HttpApiPrefixPath            pulumi.StringPtrInput
	Issuer                       pulumi.StringPtrInput
	MicrosoftAccountClientId     pulumi.StringPtrInput
	MicrosoftAccountClientSecret pulumi.StringPtrInput
	MicrosoftAccountOAuthScopes  pulumi.StringArrayInput
	Name                         pulumi.StringInput
	OpenIdIssuer                 pulumi.StringPtrInput
	ResourceGroupName            pulumi.StringInput
	TokenRefreshExtensionHours   pulumi.Float64PtrInput
	TokenStoreEnabled            pulumi.BoolPtrInput
	TwitterConsumerKey           pulumi.StringPtrInput
	TwitterConsumerSecret        pulumi.StringPtrInput
	UnauthenticatedClientAction  UnauthenticatedClientActionPtrInput
}

func (SiteAuthSettingsArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAuthSettingsArgs)(nil)).Elem()
}

type SiteAuthSettingsInput interface {
	pulumi.Input

	ToSiteAuthSettingsOutput() SiteAuthSettingsOutput
	ToSiteAuthSettingsOutputWithContext(ctx context.Context) SiteAuthSettingsOutput
}

func (*SiteAuthSettings) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAuthSettings)(nil)).Elem()
}

func (i *SiteAuthSettings) ToSiteAuthSettingsOutput() SiteAuthSettingsOutput {
	return i.ToSiteAuthSettingsOutputWithContext(context.Background())
}

func (i *SiteAuthSettings) ToSiteAuthSettingsOutputWithContext(ctx context.Context) SiteAuthSettingsOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAuthSettingsOutput)
}

type SiteAuthSettingsOutput struct{ *pulumi.OutputState }

func (SiteAuthSettingsOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAuthSettings)(nil)).Elem()
}

func (o SiteAuthSettingsOutput) ToSiteAuthSettingsOutput() SiteAuthSettingsOutput {
	return o
}

func (o SiteAuthSettingsOutput) ToSiteAuthSettingsOutputWithContext(ctx context.Context) SiteAuthSettingsOutput {
	return o
}

func init() {
	pulumi.RegisterOutputType(SiteAuthSettingsOutput{})
}
