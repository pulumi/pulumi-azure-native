package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/azureplaywrightservice/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := azureplaywrightservice.NewAccount(ctx, "account", &azureplaywrightservice.AccountArgs{
			Location:          pulumi.String("westus"),
			Name:              pulumi.String("myPlaywrightAccount"),
			RegionalAffinity:  pulumi.String("Enabled"),
			ResourceGroupName: pulumi.String("dummyrg"),
			Tags: pulumi.StringMap{
				"Team": pulumi.String("Dev Exp"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
