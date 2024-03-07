package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfiguration.NewConfigurationStore(ctx, "configurationStore", &appconfiguration.ConfigurationStoreArgs{
			ConfigStoreName:   pulumi.String("contoso"),
			DisableLocalAuth:  pulumi.Bool(true),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &appconfiguration.SkuArgs{
				Name: pulumi.String("Standard"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
