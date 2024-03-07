package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/maps/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := maps.NewAccount(ctx, "account", &maps.AccountArgs{
			AccountName: pulumi.String("myMapsAccount"),
			Kind:        pulumi.String("Gen1"),
			Location:    pulumi.String("global"),
			Properties: &maps.MapsAccountPropertiesArgs{
				DisableLocalAuth: pulumi.Bool(false),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &maps.SkuArgs{
				Name: pulumi.String("S0"),
			},
			Tags: pulumi.StringMap{
				"test": pulumi.String("true"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
