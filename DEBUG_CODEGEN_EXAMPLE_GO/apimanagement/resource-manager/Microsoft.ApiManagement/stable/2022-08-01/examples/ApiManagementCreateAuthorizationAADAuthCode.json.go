package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorization(ctx, "authorization", &apimanagement.AuthorizationArgs{
			AuthorizationId:         pulumi.String("authz2"),
			AuthorizationProviderId: pulumi.String("aadwithauthcode"),
			AuthorizationType:       pulumi.String("OAuth2"),
			OAuth2GrantType:         pulumi.String("AuthorizationCode"),
			ResourceGroupName:       pulumi.String("rg1"),
			ServiceName:             pulumi.String("apimService1"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
