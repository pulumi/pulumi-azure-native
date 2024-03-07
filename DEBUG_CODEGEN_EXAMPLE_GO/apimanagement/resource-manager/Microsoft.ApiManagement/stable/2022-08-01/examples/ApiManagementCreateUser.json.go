package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/apimanagement/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := apimanagement.NewUser(ctx, "user", &apimanagement.UserArgs{
			Confirmation:      pulumi.String("signup"),
			Email:             pulumi.String("foobar@outlook.com"),
			FirstName:         pulumi.String("foo"),
			LastName:          pulumi.String("bar"),
			ResourceGroupName: pulumi.String("rg1"),
			ServiceName:       pulumi.String("apimService1"),
			UserId:            pulumi.String("5931a75ae4bbd512288c680b"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
