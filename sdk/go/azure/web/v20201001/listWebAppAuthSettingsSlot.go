


package v20201001

import (
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func ListWebAppAuthSettingsSlot(ctx *pulumi.Context, args *ListWebAppAuthSettingsSlotArgs, opts ...pulumi.InvokeOption) (*ListWebAppAuthSettingsSlotResult, error) {
	var rv ListWebAppAuthSettingsSlotResult
	err := ctx.Invoke("azure-native:web/v20201001:listWebAppAuthSettingsSlot", args, &rv, opts...)
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
	AadClaimsAuthorization                  *string            `pulumi:"aadClaimsAuthorization"`
	AdditionalLoginParams                   []string           `pulumi:"additionalLoginParams"`
	AllowedAudiences                        []string           `pulumi:"allowedAudiences"`
	AllowedExternalRedirectUrls             []string           `pulumi:"allowedExternalRedirectUrls"`
	AuthFilePath                            *string            `pulumi:"authFilePath"`
	ClientId                                *string            `pulumi:"clientId"`
	ClientSecret                            *string            `pulumi:"clientSecret"`
	ClientSecretCertificateThumbprint       *string            `pulumi:"clientSecretCertificateThumbprint"`
	ClientSecretSettingName                 *string            `pulumi:"clientSecretSettingName"`
	DefaultProvider                         *string            `pulumi:"defaultProvider"`
	Enabled                                 *bool              `pulumi:"enabled"`
	FacebookAppId                           *string            `pulumi:"facebookAppId"`
	FacebookAppSecret                       *string            `pulumi:"facebookAppSecret"`
	FacebookAppSecretSettingName            *string            `pulumi:"facebookAppSecretSettingName"`
	FacebookOAuthScopes                     []string           `pulumi:"facebookOAuthScopes"`
	GitHubClientId                          *string            `pulumi:"gitHubClientId"`
	GitHubClientSecret                      *string            `pulumi:"gitHubClientSecret"`
	GitHubClientSecretSettingName           *string            `pulumi:"gitHubClientSecretSettingName"`
	GitHubOAuthScopes                       []string           `pulumi:"gitHubOAuthScopes"`
	GoogleClientId                          *string            `pulumi:"googleClientId"`
	GoogleClientSecret                      *string            `pulumi:"googleClientSecret"`
	GoogleClientSecretSettingName           *string            `pulumi:"googleClientSecretSettingName"`
	GoogleOAuthScopes                       []string           `pulumi:"googleOAuthScopes"`
	Id                                      string             `pulumi:"id"`
	IsAuthFromFile                          *string            `pulumi:"isAuthFromFile"`
	Issuer                                  *string            `pulumi:"issuer"`
	Kind                                    *string            `pulumi:"kind"`
	MicrosoftAccountClientId                *string            `pulumi:"microsoftAccountClientId"`
	MicrosoftAccountClientSecret            *string            `pulumi:"microsoftAccountClientSecret"`
	MicrosoftAccountClientSecretSettingName *string            `pulumi:"microsoftAccountClientSecretSettingName"`
	MicrosoftAccountOAuthScopes             []string           `pulumi:"microsoftAccountOAuthScopes"`
	Name                                    string             `pulumi:"name"`
	RuntimeVersion                          *string            `pulumi:"runtimeVersion"`
	SystemData                              SystemDataResponse `pulumi:"systemData"`
	TokenRefreshExtensionHours              *float64           `pulumi:"tokenRefreshExtensionHours"`
	TokenStoreEnabled                       *bool              `pulumi:"tokenStoreEnabled"`
	TwitterConsumerKey                      *string            `pulumi:"twitterConsumerKey"`
	TwitterConsumerSecret                   *string            `pulumi:"twitterConsumerSecret"`
	TwitterConsumerSecretSettingName        *string            `pulumi:"twitterConsumerSecretSettingName"`
	Type                                    string             `pulumi:"type"`
	UnauthenticatedClientAction             *string            `pulumi:"unauthenticatedClientAction"`
	ValidateIssuer                          *bool              `pulumi:"validateIssuer"`
}
