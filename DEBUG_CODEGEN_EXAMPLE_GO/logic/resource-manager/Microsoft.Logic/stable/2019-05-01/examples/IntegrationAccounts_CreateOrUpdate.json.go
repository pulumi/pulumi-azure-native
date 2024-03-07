package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/logic/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := logic.NewIntegrationAccount(ctx, "integrationAccount", &logic.IntegrationAccountArgs{
			IntegrationAccountName: pulumi.String("testIntegrationAccount"),
			Location:               pulumi.String("westus"),
			ResourceGroupName:      pulumi.String("testResourceGroup"),
			Sku: &logic.IntegrationAccountSkuArgs{
				Name: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
