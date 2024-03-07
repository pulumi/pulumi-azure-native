package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/automanage/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := automanage.NewAccount(ctx, "account", &automanage.AccountArgs{
			AccountName: pulumi.String("account"),
			Identity: &automanage.AccountIdentityArgs{
				Type: automanage.ResourceIdentityTypeSystemAssigned,
			},
			Location:          pulumi.String("East US"),
			ResourceGroupName: pulumi.String("resourceGroup"),
			Tags: pulumi.StringMap{
				"Organization": pulumi.String("Administration"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
