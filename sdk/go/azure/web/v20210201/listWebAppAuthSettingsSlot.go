


package v20210201

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAuthSettingsSlot(ctx *pulumi.Context, args *ListWebAppAuthSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppAuthSettingsSlotResult, error) {
	var rv ListWebAppAuthSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20210201:listWebAppAuthSettingsSlot", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAuthSettingsSlotArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
	Slot              string `pulumi:"slot"`
}


type ListWebAppAuthSettingsSlotResult struct {
	AadClaimsAuthorization                  *string  `pulumi:"aadClaimsAuthorization"`
	AdditionalLoginParams                   []string `pulumi:"additionalLoginParams"`
	AllowedAudiences                        []string `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls             []string `pulumi:"allowedExternalRedirectUrls"`
	AuthFilePath                            *string  `pulumi:"authFilePath"`
	ClientId                                *string  `pulumi:"clientId"`
	ClientSecret                            *string  `pulumi:"clientSecret"`
	ClientSecretCertificateThumbprint       *string  `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                 *string  `pulumi:"clientSecretSettingName"`
	ConfigVersion                           *string  `pulumi:"configVersion"`
	DefaultProvider                         *string  `pulumi:"defaultProvider"`
	Enabled                                 *bool    `pulumi:"enabled"`
	FacebookAppId                           *string  `pulumi:"facebookAppId"`
	FacebookAppSecret                       *string  `pulumi:"facebookAppSecret"`
	FacebookAppSecretSettingName            *string  `pulumi:"facebookAppSecretSettingName"`
	FacebookOAuthScopes                     []string `pulumi:"facebookOAuthScopes"`
	GitHubClientId                          *string  `pulumi:"gitHubClientId"`
	GitHubClientSecret                      *string  `pulumi:"gitHubClientSecret"`
	GitHubClientSecretSettingName           *string  `pulumi:"gitHubClientSecretSettingName"`
	GitHubOAuthScopes                       []string `pulumi:"gitHubOAuthScopes"`
	GoogleClientId                          *string  `pulumi:"googleClientId"`
	GoogleClientSecret                      *string  `pulumi:"googleClientSecret"`
	GoogleClientSecretSettingName           *string  `pulumi:"googleClientSecretSettingName"`
	GoogleOAuthScopes                       []string `pulumi:"googleOAuthScopes"`
	Id                                      string   `pulumi:"id"`
	IsAuthFromFile                          *string  `pulumi:"isAuthFromFile"`
	Issuer                                  *string  `pulumi:"issuer"`
	Kind                                    *string  `pulumi:"kind"`
	MicrosoftAccountClientId                *string  `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret            *string  `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountClientSecretSettingName *string  `pulumi:"microsoftAccountClientSecretSettingName"`
	MicrosoftAccountOAuthScopes             []string `pulumi:"microsoftAccountOAuthScopes"`
	Name                                    string   `pulumi:"name"`
	RuntimeVersion                          *string  `pulumi:"runtimeVersion"`
	TokenRefreshExtensionHours              *float64 `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                       *bool    `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                      *string  `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret                   *string  `pulumi:"twitterConsumerSecret"`
	TwitterConsumerSecretSettingName        *string  `pulumi:"twitterConsumerSecretSettingName"`
	Type                                    string   `pulumi:"type"`
	UnauthenticatedClientAction             *string  `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                          *bool    `pulumi:"validateIssuer"`
}

func ListWebAppAuthSettingsSlotOutput(ctx *pulumi.Context, args ListWebAppAuthSettingsSlotOutputArgs, opts ...pulumi.InvokeOption) ListWebAppAuthSettingsSlotResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppAuthSettingsSlotResult, error) {
			args := v.(ListWebAppAuthSettingsSlotArgs)
			r, err := ListWebAppAuthSettingsSlot(ctx, &args, opts...)
			var s ListWebAppAuthSettingsSlotResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppAuthSettingsSlotResultOutput)
}

type ListWebAppAuthSettingsSlotOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
	Slot              pulumi.StringInput `pulumi:"slot"`
}

func (ListWebAppAuthSettingsSlotOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAuthSettingsSlotArgs)(nil)).Elem()
}


type ListWebAppAuthSettingsSlotResultOutput struct{ *pulumi.OutputState }

func (ListWebAppAuthSettingsSlotResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAuthSettingsSlotResult)(nil)).Elem()
}

func (o ListWebAppAuthSettingsSlotResultOutput) ToListWebAppAuthSettingsSlotResultOutput() ListWebAppAuthSettingsSlotResultOutput {
	return o
}

func (o ListWebAppAuthSettingsSlotResultOutput) ToListWebAppAuthSettingsSlotResultOutputWithContext(ctx context.Context) ListWebAppAuthSettingsSlotResultOutput {
	return o
}

func (o ListWebAppAuthSettingsSlotResultOutput) AadClaimsAuthorization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.AadClaimsAuthorization }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) AuthFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.AuthFilePath }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) ConfigVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.ConfigVersion }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) FacebookAppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.FacebookAppSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GitHubClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.GitHubClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GitHubClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.GitHubClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GitHubClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.GitHubClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GitHubOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.GitHubOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GoogleClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.GoogleClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) IsAuthFromFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.IsAuthFromFile }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) MicrosoftAccountClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.MicrosoftAccountClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) []string { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *bool { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) TwitterConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.TwitterConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsSlotResultOutput) ValidateIssuer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsSlotResult) *bool { return v.ValidateIssuer }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppAuthSettingsSlotResultOutput{})
}
