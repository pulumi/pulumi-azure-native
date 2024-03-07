package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewAuthorization(ctx, "authorization", &apimanagement.AuthorizationArgs{
			AuthorizationId:         pulumi.String("authz1"),
			AuthorizationProviderId: pulumi.String("aadwithclientcred"),
			AuthorizationType:       pulumi.String("OAuth2"),
			OAuth2GrantType:         pulumi.String("AuthorizationCode"),
			Parameters: pulumi.StringMap{
				"clientId":     pulumi.String("53790925-fdd3-4b80-bc7a-4c3aaf25801d"),
				"clientSecret": pulumi.String("FcJkQ3iPSaKAQRA7Ft8Q~fZ1X5vKmqzUAfJagcJ8"),
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
