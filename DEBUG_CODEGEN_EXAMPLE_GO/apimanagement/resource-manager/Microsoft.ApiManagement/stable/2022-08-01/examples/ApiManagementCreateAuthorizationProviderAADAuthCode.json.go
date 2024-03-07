package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorizationProvider(ctx, "authorizationProvider", &apimanagement.AuthorizationProviderArgs{
			AuthorizationProviderId: pulumi.String("aadwithauthcode"),
			DisplayName:             pulumi.String("aadwithauthcode"),
			IdentityProvider:        pulumi.String("aad"),
			Oauth2: apimanagement.AuthorizationProviderOAuth2SettingsResponse{
				GrantTypes: &apimanagement.AuthorizationProviderOAuth2GrantTypesArgs{
					AuthorizationCode: pulumi.StringMap{
						"clientId":     pulumi.String("59790825-fdd3-4b10-bc7a-4c3aaf25801d"),
						"clientSecret": pulumi.String("Q3iPSaKQ~fZFcJk5vKmqzUAfJagcJ8"),
						"resourceUri":  pulumi.String("https://graph.microsoft.com"),
						"scopes":       pulumi.String("User.Read.All Group.Read.All"),
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
