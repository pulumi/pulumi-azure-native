


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteAuthSettings(ctx *pulumi.Context, args *ListSiteAuthSettingsArgs, opts ...pulumi.InvokeOption) (*ListSiteAuthSettingsResult, error) {
	var rv ListSiteAuthSettingsResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAuthSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAuthSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListSiteAuthSettingsResult struct {
	AadClientId                  *string  `pulumi:"aadClientId"`
	AdditionalLoginParams        []string `pulumi:"additionalLoginParams"`
	AllowedAudiences             []string `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls  []string `pulumi:"allowedExternalRedirectUrls"`
	ClientId                     *string  `pulumi:"clientId"`
	ClientSecret                 *string  `pulumi:"clientSecret"`
	DefaultProvider              *string  `pulumi:"defaultProvider"`
	Enabled                      *bool    `pulumi:"enabled"`
	FacebookAppId                *string  `pulumi:"facebookAppId"`
	FacebookAppSecret            *string  `pulumi:"facebookAppSecret"`
	FacebookOAuthScopes          []string `pulumi:"facebookOAuthScopes"`
	GoogleClientId               *string  `pulumi:"googleClientId"`
	GoogleClientSecret           *string  `pulumi:"googleClientSecret"`
	GoogleOAuthScopes            []string `pulumi:"googleOAuthScopes"`
	HttpApiPrefixPath            *string  `pulumi:"httpApiPrefixPath"`
	Issuer                       *string  `pulumi:"issuer"`
	MicrosoftAccountClientId     *string  `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret *string  `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountOAuthScopes  []string `pulumi:"microsoftAccountOAuthScopes"`
	OpenIdIssuer                 *string  `pulumi:"openIdIssuer"`
	TokenRefreshExtensionHours   *float64 `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled            *bool    `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey           *string  `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret        *string  `pulumi:"twitterConsumerSecret"`
	UnauthenticatedClientAction  *string  `pulumi:"unauthenticatedClientAction"`
}

func ListSiteAuthSettingsOutput(ctx *pulumi.Context, args ListSiteAuthSettingsOutputArgs, opts ...pulumi.InvokeOption) ListSiteAuthSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteAuthSettingsResult, error) {
			args := v.(ListSiteAuthSettingsArgs)
			r, err := ListSiteAuthSettings(ctx, &args, opts...)
			var s ListSiteAuthSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteAuthSettingsResultOutput)
}

type ListSiteAuthSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListSiteAuthSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAuthSettingsArgs)(nil)).Elem()
}


type ListSiteAuthSettingsResultOutput struct{ *pulumi.OutputState }

func (ListSiteAuthSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAuthSettingsResult)(nil)).Elem()
}

func (o ListSiteAuthSettingsResultOutput) ToListSiteAuthSettingsResultOutput() ListSiteAuthSettingsResultOutput {
	return o
}

func (o ListSiteAuthSettingsResultOutput) ToListSiteAuthSettingsResultOutputWithContext(ctx context.Context) ListSiteAuthSettingsResultOutput {
	return o
}

func (o ListSiteAuthSettingsResultOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) []string { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsResultOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsResultOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) []string { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsResultOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) []string { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsResultOutput) HttpApiPrefixPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.HttpApiPrefixPath }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) []string { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsResultOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *bool { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsResultOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsResult) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteAuthSettingsResultOutput{})
}
