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
				Openid: &apimanagement.OpenIdAuthenticationSettingsContractArgs{
					BearerTokenSendingMethods: pulumi.StringArray{
						pulumi.String("authorizationHeader"),
					},
					OpenidProviderId: pulumi.String("testopenid"),
				},
			},
			Description: pulumi.String("This is a sample server Petstore server.  You can find out more about Swagger at [http://swagger.io](http://swagger.io) or on [irc.freenode.net, #swagger](http://swagger.io/irc/).  For this sample, you can use the api key `special-key` to test the authorization filters."),
			DisplayName: pulumi.String("Swagger Petstore"),
			Path:        pulumi.String("petstore"),
			Protocols: pulumi.StringArray{
				pulumi.String("https"),
			},
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			ServiceUrl:        pulumi.String("http://petstore.swagger.io/v2"),
			SubscriptionKeyParameterNames: &apimanagement.SubscriptionKeyParameterNamesContractArgs{
				Header: pulumi.String("Ocp-Apim-Subscription-Key"),
				Query:  pulumi.String("subscription-key"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
