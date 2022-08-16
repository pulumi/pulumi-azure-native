


package v20150801

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)


type SiteAuthSettingsSlot struct {
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


func NewSiteAuthSettingsSlot(ctx *pulumi.Context,
	name string, args *SiteAuthSettingsSlotArgs, opts ...pulumi.ResourceOption) (*SiteAuthSettingsSlot, error) {
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
			Type: pulumi.String("azure-native:web:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20160801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20180201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20181101:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20190801:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200601:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20200901:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201001:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20201201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210101:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210115:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210201:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20210301:SiteAuthSettingsSlot"),
		},
		{
			Type: pulumi.String("azure-native:web/v20220301:SiteAuthSettingsSlot"),
		},
	})
	opts = append(opts, aliases)
	var resource SiteAuthSettingsSlot
	err := ctx.RegisterResource("azure-native:web/v20150801:SiteAuthSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetSiteAuthSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *SiteAuthSettingsSlotState, opts ...pulumi.ResourceOption) (*SiteAuthSettingsSlot, error) {
	var resource SiteAuthSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20150801:SiteAuthSettingsSlot", name, id, state, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}


type siteAuthSettingsSlotState struct {
}

type SiteAuthSettingsSlotState struct {
}

func (SiteAuthSettingsSlotState) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAuthSettingsSlotState)(nil)).Elem()
}

type siteAuthSettingsSlotArgs struct {
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
	Slot                         string                         `pulumi:"slot"`
	TokenRefreshExtensionHours   *float64                       `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled            *bool                          `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey           *string                        `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret        *string                        `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction  *UnauthenticatedClientAction   `pulumi:"unauthenticatedClientAction"`
}


type SiteAuthSettingsSlotArgs struct {
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
	Slot                         pulumi.StringInput
	TokenRefreshExtensionHours   pulumi.Float64PtrInput
	TokenStoreEnabled            pulumi.BoolPtrInput
	TwitterConsumerKey           pulumi.StringPtrInput
	TwitterConsumerSecret        pulumi.StringPtrInput
	UnauthenticatedClientAction  UnauthenticatedClientActionPtrInput
}

func (SiteAuthSettingsSlotArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*siteAuthSettingsSlotArgs)(nil)).Elem()
}

type SiteAuthSettingsSlotInput interface {
	pulumi.Input

	ToSiteAuthSettingsSlotOutput() SiteAuthSettingsSlotOutput
	ToSiteAuthSettingsSlotOutputWithContext(ctx context.Context) SiteAuthSettingsSlotOutput
}

func (*SiteAuthSettingsSlot) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAuthSettingsSlot)(nil)).Elem()
}

func (i *SiteAuthSettingsSlot) ToSiteAuthSettingsSlotOutput() SiteAuthSettingsSlotOutput {
	return i.ToSiteAuthSettingsSlotOutputWithContext(context.Background())
}

func (i *SiteAuthSettingsSlot) ToSiteAuthSettingsSlotOutputWithContext(ctx context.Context) SiteAuthSettingsSlotOutput {
	return pulumi.ToOutputWithContext(ctx, i).(SiteAuthSettingsSlotOutput)
}

type SiteAuthSettingsSlotOutput struct{ *pulumi.OutputState }

func (SiteAuthSettingsSlotOutput) ElementType() reflect.Type {
	return reflect.TypeOf((**SiteAuthSettingsSlot)(nil)).Elem()
}

func (o SiteAuthSettingsSlotOutput) ToSiteAuthSettingsSlotOutput() SiteAuthSettingsSlotOutput {
	return o
}

func (o SiteAuthSettingsSlotOutput) ToSiteAuthSettingsSlotOutputWithContext(ctx context.Context) SiteAuthSettingsSlotOutput {
	return o
}

func (o SiteAuthSettingsSlotOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringArrayOutput { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o SiteAuthSettingsSlotOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringArrayOutput { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o SiteAuthSettingsSlotOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringArrayOutput { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o SiteAuthSettingsSlotOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.BoolPtrOutput { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringArrayOutput { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o SiteAuthSettingsSlotOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringArrayOutput { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o SiteAuthSettingsSlotOutput) HttpApiPrefixPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.HttpApiPrefixPath }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringArrayOutput { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o SiteAuthSettingsSlotOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.Float64PtrOutput { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

func (o SiteAuthSettingsSlotOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.BoolPtrOutput { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

func (o SiteAuthSettingsSlotOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *SiteAuthSettingsSlot) pulumi.StringPtrOutput { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(SiteAuthSettingsSlotOutput{})
}
