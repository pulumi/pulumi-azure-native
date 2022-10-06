


package v20150801

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListSiteAuthSettingsSlot(ctx *pulumi.Context, args *ListSiteAuthSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListSiteAuthSettingsSlotResult, error) {
	var rv ListSiteAuthSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20150801:listSiteAuthSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListSiteAuthSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListSiteAuthSettingsSlotResult struct {
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

func ListSiteAuthSettingsSlotOutput(ctx *pulumi.Context, args ListSiteAuthSettingsSlotOutputArgs, opts ...pulumi.InvokeOption) ListSiteAuthSettingsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListSiteAuthSettingsSlotResult, error) {
			args := v.(ListSiteAuthSettingsSlotArgs)
			r, err := ListSiteAuthSettingsSlot(ctx, &args, opts...)
			var s ListSiteAuthSettingsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListSiteAuthSettingsSlotResultOutput)
}

type ListSiteAuthSettingsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListSiteAuthSettingsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAuthSettingsSlotArgs)(nil)).Elem()
}


type ListSiteAuthSettingsSlotResultOutput struct{ *pulumi.OutputState }

func (ListSiteAuthSettingsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListSiteAuthSettingsSlotResult)(nil)).Elem()
}

func (o ListSiteAuthSettingsSlotResultOutput) ToListSiteAuthSettingsSlotResultOutput() ListSiteAuthSettingsSlotResultOutput {
	return o
}

func (o ListSiteAuthSettingsSlotResultOutput) ToListSiteAuthSettingsSlotResultOutputWithContext(ctx context.Context) ListSiteAuthSettingsSlotResultOutput {
	return o
}

func (o ListSiteAuthSettingsSlotResultOutput) AadClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.AadClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) []string { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) []string { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) []string { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) HttpApiPrefixPath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.HttpApiPrefixPath }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) []string { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) OpenIdIssuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.OpenIdIssuer }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *bool { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

func (o ListSiteAuthSettingsSlotResultOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListSiteAuthSettingsSlotResult) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListSiteAuthSettingsSlotResultOutput{})
}
