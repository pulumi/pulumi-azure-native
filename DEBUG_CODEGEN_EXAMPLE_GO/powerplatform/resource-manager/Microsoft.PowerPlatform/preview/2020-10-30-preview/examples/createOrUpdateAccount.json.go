package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/powerplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := powerplatform.NewAccount(ctx, "account", &powerplatform.AccountArgs{
			AccountName:       pulumi.String("account"),
			Description:       pulumi.String("Description of the account."),
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
