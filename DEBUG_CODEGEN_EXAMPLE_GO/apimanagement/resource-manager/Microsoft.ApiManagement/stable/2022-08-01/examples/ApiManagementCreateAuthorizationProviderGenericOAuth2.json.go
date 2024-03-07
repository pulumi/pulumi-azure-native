package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorizationProvider(ctx, "authorizationProvider", &apimanagement.AuthorizationProviderArgs{
			AuthorizationProviderId: pulumi.String("eventbrite"),
			DisplayName:             pulumi.String("eventbrite"),
			IdentityProvider:        pulumi.String("oauth2"),
			Oauth2: apimanagement.AuthorizationProviderOAuth2SettingsResponse{
				GrantTypes: &apimanagement.AuthorizationProviderOAuth2GrantTypesArgs{
					AuthorizationCode: pulumi.StringMap{
						"authorizationUrl": pulumi.String("https://www.eventbrite.com/oauth/authorize"),
						"clientId":         pulumi.String("ZYIJTBTABHOUQQDLZY"),
						"clientSecret":     pulumi.String("Q3iPSaKQ~fZFcJk5vKmqzUAfJagcJ8"),
						"refreshUrl":       pulumi.String("https://www.eventbrite.com/oauth/token"),
						"scopes":           pulumi.String(""),
						"tokenUrl":         pulumi.String("https://www.eventbrite.com/oauth/token"),
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
