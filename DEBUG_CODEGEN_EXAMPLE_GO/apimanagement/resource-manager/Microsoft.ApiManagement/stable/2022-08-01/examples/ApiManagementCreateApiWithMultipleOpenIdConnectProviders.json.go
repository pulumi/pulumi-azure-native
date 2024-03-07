package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewApi(ctx, "api", &apimanagement.ApiArgs{
			ApiId: pulumi.String("tempgroup"),
			AuthenticationSettings: apimanagement.AuthenticationSettingsContractResponse{
				OpenidAuthenticationSettings: apimanagement.OpenIdAuthenticationSettingsContractArray{
					&apimanagement.OpenIdAuthenticationSettingsContractArgs{
						BearerTokenSendingMethods: pulumi.StringArray{
							pulumi.String("authorizationHeader"),
						},
						OpenidProviderId: pulumi.String("openidProviderId2283"),
					},
					&apimanagement.OpenIdAuthenticationSettingsContractArgs{
						BearerTokenSendingMethods: pulumi.StringArray{
							pulumi.String("authorizationHeader"),
						},
						OpenidProviderId: pulumi.String("openidProviderId2284"),
					},
				},
			},
			Description: pulumi.String("apidescription5200"),
			DisplayName: pulumi.String("apiname1463"),
			Path:        pulumi.String("newapiPath"),
			Protocols: pulumi.StringArray{
				pulumi.String("https"),
				pulumi.String("http"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			ServiceUrl:        pulumi.String("http://newechoapi.cloudapp.net/api"),
			SubscriptionKeyParameterNames: &apimanagement.SubscriptionKeyParameterNamesContractArgs{
				Header: pulumi.String("header4520"),
				Query:  pulumi.String("query3037"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
