


package v20181101

import (
	"context"
	"reflect"

	"github.com/pkg/errors"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

type WebAppAuthSettingsSlot struct {
	pulumi.CustomResourceState

	AdditionalLoginParams             pulumi.StringArrayOutput `pulumi:"additionalLoginParams"`
	AllowedAudiences                  pulumi.StringArrayOutput `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls       pulumi.StringArrayOutput `pulumi:"allowedExternalRedirectUrls"`
	ClientId                          pulumi.StringPtrOutput   `pulumi:"clientId"`
	ClientSecret                      pulumi.StringPtrOutput   `pulumi:"clientSecret"`
	ClientSecretCertificateThumbprint pulumi.StringPtrOutput   `pulumi:"clientSecretCertificateThumbprint"`
	DefaultProvider                   pulumi.StringPtrOutput   `pulumi:"defaultProvider"`
	Enabled                           pulumi.BoolPtrOutput     `pulumi:"enabled"`
	FacebookAppId                     pulumi.StringPtrOutput   `pulumi:"facebookAppId"`
	FacebookAppSecret                 pulumi.StringPtrOutput   `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes               pulumi.StringArrayOutput `pulumi:"facebookOAuthScopes"`
	GoogleClientId                    pulumi.StringPtrOutput   `pulumi:"googleClientId"`
	GoogleClientSecret                pulumi.StringPtrOutput   `pulumi:"googleClientSecret"`
	GoogleOAuthScopes                 pulumi.StringArrayOutput `pulumi:"googleOAuthScopes"`
	Issuer                            pulumi.StringPtrOutput   `pulumi:"issuer"`
	Kind                              pulumi.StringPtrOutput   `pulumi:"kind"`
	MicrosoftAccountClientId          pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret      pulumi.StringPtrOutput   `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes       pulumi.StringArrayOutput `pulumi:"microsoftAccountOAuthScopes"`
	Name                              pulumi.StringOutput      `pulumi:"name"`
	RuntimeVersion                    pulumi.StringPtrOutput   `pulumi:"runtimeVersion"`
	TokenRefreshExtensionHours        pulumi.Float64PtrOutput  `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                 pulumi.BoolPtrOutput     `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                pulumi.StringPtrOutput   `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret             pulumi.StringPtrOutput   `pulumi:"twitterConsumerSecret"`
	Type                              pulumi.StringOutput      `pulumi:"type"`
	UnauthenticatedClientAction       pulumi.StringPtrOutput   `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                    pulumi.BoolPtrOutput     `pulumi:"validateIssuer"`
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
			Type: pulumi.String("azure-native:web/v20201201:WebAppAuthSettingsSlot"),
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
	err := ctx.RegisterResource("azure-native:web/v20181101:WebAppAuthSettingsSlot", name, args, &resource, opts...)
	if err != nil {
		return nil, err
	}
	return &resource, nil
}



func GetWebAppAuthSettingsSlot(ctx *pulumi.Context,
	name string, id pulumi.IDInput, state *WebAppAuthSettingsSlotState, opts ...pulumi.ResourceOption) (*WebAppAuthSettingsSlot, error) {
	var resource WebAppAuthSettingsSlot
	err := ctx.ReadResource("azure-native:web/v20181101:WebAppAuthSettingsSlot", name, id, state, &resource, opts...)
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
	AdditionalLoginParams             []string                       `pulumi:"additionalLoginParams"`
	AllowedAudiences                  []string                       `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls       []string                       `pulumi:"allowedExternalRedirectUrls"`
	ClientId                          *string                        `pulumi:"clientId"`
	ClientSecret                      *string                        `pulumi:"clientSecret"`
	ClientSecretCertificateThumbprint *string                        `pulumi:"clientSecretCertificateThumbprint"`
	DefaultProvider                   *BuiltInAuthenticationProvider `pulumi:"defaultProvider"`
	Enabled                           *bool                          `pulumi:"enabled"`
	FacebookAppId                     *string                        `pulumi:"facebookAppId"`
	FacebookAppSecret                 *string                        `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes               []string                       `pulumi:"facebookOAuthScopes"`
	GoogleClientId                    *string                        `pulumi:"googleClientId"`
	GoogleClientSecret                *string                        `pulumi:"googleClientSecret"`
	GoogleOAuthScopes                 []string                       `pulumi:"googleOAuthScopes"`
	Issuer                            *string                        `pulumi:"issuer"`
	Kind                              *string                        `pulumi:"kind"`
	MicrosoftAccountClientId          *string                        `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret      *string                        `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes       []string                       `pulumi:"microsoftAccountOAuthScopes"`
	Name                              string                         `pulumi:"name"`
	ResourceGroupName                 string                         `pulumi:"resourceGroupName"`
	RuntimeVersion                    *string                        `pulumi:"runtimeVersion"`
	Slot                              string                         `pulumi:"slot"`
	TokenRefreshExtensionHours        *float64                       `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                 *bool                          `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                *string                        `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret             *string                        `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction       *UnauthenticatedClientAction   `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                    *bool                          `pulumi:"validateIssuer"`
}


type WebAppAuthSettingsSlotArgs struct {
	AdditionalLoginParams             pulumi.StringArrayInput
	AllowedAudiences                  pulumi.StringArrayInput
	AllowedExternalRedirectUrls       pulumi.StringArrayInput
	ClientId                          pulumi.StringPtrInput
	ClientSecret                      pulumi.StringPtrInput
	ClientSecretCertificateThumbprint pulumi.StringPtrInput
	DefaultProvider                   BuiltInAuthenticationProviderPtrInput
	Enabled                           pulumi.BoolPtrInput
	FacebookAppId                     pulumi.StringPtrInput
	FacebookAppSecret                 pulumi.StringPtrInput
	FacebookOAuthScopes               pulumi.StringArrayInput
	GoogleClientId                    pulumi.StringPtrInput
	GoogleClientSecret                pulumi.StringPtrInput
	GoogleOAuthScopes                 pulumi.StringArrayInput
	Issuer                            pulumi.StringPtrInput
	Kind                              pulumi.StringPtrInput
	MicrosoftAccountClientId          pulumi.StringPtrInput
	MicrosoftAccountClientSecret      pulumi.StringPtrInput
	MicrosoftAccountOAuthScopes       pulumi.StringArrayInput
	Name                              pulumi.StringInput
	ResourceGroupName                 pulumi.StringInput
	RuntimeVersion                    pulumi.StringPtrInput
	Slot                              pulumi.StringInput
	TokenRefreshExtensionHours        pulumi.Float64PtrInput
	TokenStoreEnabled                 pulumi.BoolPtrInput
	TwitterConsumerKey                pulumi.StringPtrInput
	TwitterConsumerSecret             pulumi.StringPtrInput
	UnauthenticatedClientAction       UnauthenticatedClientActionPtrInput
	ValidateIssuer                    pulumi.BoolPtrInput
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

func (o WebAppAuthSettingsSlotOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
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

func (o WebAppAuthSettingsSlotOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringPtrOutput { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o WebAppAuthSettingsSlotOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v *WebAppAuthSettingsSlot) pulumi.StringArrayOutput { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
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
