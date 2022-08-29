


package v20210101

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAuthSettings(ctx *pulumi.Context, args *ListWebAppAuthSettingsArgs, opts ...pulumi.InvokeOption) (*ListWebAppAuthSettingsResult, error) {
	var rv ListWebAppAuthSettingsResult
	err := ctx.Invoke("azure-native:web/v20210101:listWebAppAuthSettings", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

type ListWebAppAuthSettingsArgs struct {
	Name              string `pulumi:"name"`
	ResourceGroupName string `pulumi:"resourceGroupName"`
}


type ListWebAppAuthSettingsResult struct {
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

func ListWebAppAuthSettingsOutput(ctx *pulumi.Context, args ListWebAppAuthSettingsOutputArgs, opts ...pulumi.InvokeOption) ListWebAppAuthSettingsResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (ListWebAppAuthSettingsResult, error) {
			args := v.(ListWebAppAuthSettingsArgs)
			r, err := ListWebAppAuthSettings(ctx, &args, opts...)
			var s ListWebAppAuthSettingsResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(ListWebAppAuthSettingsResultOutput)
}

type ListWebAppAuthSettingsOutputArgs struct {
	Name              pulumi.StringInput `pulumi:"name"`
	ResourceGroupName pulumi.StringInput `pulumi:"resourceGroupName"`
}

func (ListWebAppAuthSettingsOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAuthSettingsArgs)(nil)).Elem()
}


type ListWebAppAuthSettingsResultOutput struct{ *pulumi.OutputState }

func (ListWebAppAuthSettingsResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*ListWebAppAuthSettingsResult)(nil)).Elem()
}

func (o ListWebAppAuthSettingsResultOutput) ToListWebAppAuthSettingsResultOutput() ListWebAppAuthSettingsResultOutput {
	return o
}

func (o ListWebAppAuthSettingsResultOutput) ToListWebAppAuthSettingsResultOutputWithContext(ctx context.Context) ListWebAppAuthSettingsResultOutput {
	return o
}

func (o ListWebAppAuthSettingsResultOutput) AadClaimsAuthorization() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.AadClaimsAuthorization }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) AdditionalLoginParams() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.AdditionalLoginParams }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) AllowedAudiences() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.AllowedAudiences }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) AllowedExternalRedirectUrls() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.AllowedExternalRedirectUrls }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) AuthFilePath() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.AuthFilePath }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) ClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) ClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) ClientSecretCertificateThumbprint() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientSecretCertificateThumbprint }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) ClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) ConfigVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.ConfigVersion }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) DefaultProvider() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.DefaultProvider }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) Enabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *bool { return v.Enabled }).(pulumi.BoolPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) FacebookAppId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.FacebookAppId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) FacebookAppSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.FacebookAppSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) FacebookAppSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.FacebookAppSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) FacebookOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.FacebookOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GitHubClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GitHubClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GitHubClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GitHubClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GitHubClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GitHubClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GitHubOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.GitHubOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GoogleClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GoogleClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GoogleClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GoogleClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GoogleClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.GoogleClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) GoogleOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.GoogleOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o ListWebAppAuthSettingsResultOutput) IsAuthFromFile() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.IsAuthFromFile }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) Issuer() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.Issuer }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) Kind() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.Kind }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountClientId() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.MicrosoftAccountClientId }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountClientSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.MicrosoftAccountClientSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountClientSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.MicrosoftAccountClientSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) MicrosoftAccountOAuthScopes() pulumi.StringArrayOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) []string { return v.MicrosoftAccountOAuthScopes }).(pulumi.StringArrayOutput)
}

func (o ListWebAppAuthSettingsResultOutput) Name() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) string { return v.Name }).(pulumi.StringOutput)
}

func (o ListWebAppAuthSettingsResultOutput) RuntimeVersion() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.RuntimeVersion }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) TokenRefreshExtensionHours() pulumi.Float64PtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *float64 { return v.TokenRefreshExtensionHours }).(pulumi.Float64PtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) TokenStoreEnabled() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *bool { return v.TokenStoreEnabled }).(pulumi.BoolPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) TwitterConsumerKey() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.TwitterConsumerKey }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) TwitterConsumerSecret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.TwitterConsumerSecret }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) TwitterConsumerSecretSettingName() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.TwitterConsumerSecretSettingName }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) Type() pulumi.StringOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) string { return v.Type }).(pulumi.StringOutput)
}

func (o ListWebAppAuthSettingsResultOutput) UnauthenticatedClientAction() pulumi.StringPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *string { return v.UnauthenticatedClientAction }).(pulumi.StringPtrOutput)
}

func (o ListWebAppAuthSettingsResultOutput) ValidateIssuer() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v ListWebAppAuthSettingsResult) *bool { return v.ValidateIssuer }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(ListWebAppAuthSettingsResultOutput{})
}
