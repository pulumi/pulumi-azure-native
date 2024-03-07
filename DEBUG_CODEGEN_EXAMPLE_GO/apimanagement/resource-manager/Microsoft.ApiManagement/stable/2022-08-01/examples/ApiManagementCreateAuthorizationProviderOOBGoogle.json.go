package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorizationProvider(ctx, "authorizationProvider", &apimanagement.AuthorizationProviderArgs{
			AuthorizationProviderId: pulumi.String("google"),
			DisplayName:             pulumi.String("google"),
			IdentityProvider:        pulumi.String("google"),
			Oauth2: apimanagement.AuthorizationProviderOAuth2SettingsResponse{
				GrantTypes: &apimanagement.AuthorizationProviderOAuth2GrantTypesArgs{
					AuthorizationCode: pulumi.StringMap{
						"clientId":     pulumi.String("508791967882-5qv6o2i99a75un7329vlegtk78kr766h.apps.googleusercontent.com"),
						"clientSecret": pulumi.String("qDN0VyVFjU1OsOyT5Kz8ce"),
						"scopes":       pulumi.String("openid https://www.googleapis.com/auth/userinfo.profile https://www.googleapis.com/auth/userinfo.email"),
					},
				},
				RedirectUrl: pulumi.String("https://authorization-manager.consent.azure-apim.net/redirect/apim/apimService1"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
