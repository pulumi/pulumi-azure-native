package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/datashare/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := datashare.NewAccount(ctx, "account", &datashare.AccountArgs{
			AccountName: pulumi.String("Account1"),
			Identity: &datashare.IdentityArgs{
				Type: pulumi.String("SystemAssigned"),
			},
			Location:          pulumi.String("West US 2"),
			ResourceGroupName: pulumi.String("SampleResourceGroup"),
			Tags: pulumi.StringMap{
				"tag1": pulumi.String("Red"),
				"tag2": pulumi.String("White"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
