package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appconfiguration/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appconfiguration.NewConfigurationStore(ctx, "configurationStore", &appconfiguration.ConfigurationStoreArgs{
			ConfigStoreName:   pulumi.String("contoso"),
			Location:          pulumi.String("westus"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &appconfiguration.SkuArgs{
				Name: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"myTag": pulumi.String("myTagValue"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
