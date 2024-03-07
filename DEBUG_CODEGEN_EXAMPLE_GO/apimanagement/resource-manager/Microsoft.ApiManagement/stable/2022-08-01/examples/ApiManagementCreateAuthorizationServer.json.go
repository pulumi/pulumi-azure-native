package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorizationServer(ctx, "authorizationServer", &apimanagement.AuthorizationServerArgs{
			AuthorizationEndpoint: pulumi.String("https://www.contoso.com/oauth2/auth"),
			AuthorizationMethods: apimanagement.AuthorizationMethodArray{
				apimanagement.AuthorizationMethodGET,
			},
			Authsid: pulumi.String("newauthServer"),
			BearerTokenSendingMethods: pulumi.StringArray{
				pulumi.String("authorizationHeader"),
			},
			ClientId:                   pulumi.String("1"),
			ClientRegistrationEndpoint: pulumi.String("https://www.contoso.com/apps"),
			ClientSecret:               pulumi.String("2"),
			DefaultScope:               pulumi.String("read write"),
			Description:                pulumi.String("test server"),
			DisplayName:                pulumi.String("test2"),
			GrantTypes: pulumi.StringArray{
				pulumi.String("authorizationCode"),
				pulumi.String("implicit"),
			},
			ResourceGroupName:     pulumi.String("rg1"),
			ResourceOwnerPassword: pulumi.String("pwd"),
			ResourceOwnerUsername: pulumi.String("un"),
			ServiceName:           pulumi.String("apimService1"),
			SupportState:          pulumi.Bool(true),
			TokenEndpoint:         pulumi.String("https://www.contoso.com/oauth2/token"),
			UseInApiDocumentation: pulumi.Bool(true),
			UseInTestConsole:      pulumi.Bool(false),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
