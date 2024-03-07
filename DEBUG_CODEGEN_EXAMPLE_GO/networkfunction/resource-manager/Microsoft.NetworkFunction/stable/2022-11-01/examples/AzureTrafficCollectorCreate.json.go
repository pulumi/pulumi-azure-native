package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/networkfunction/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := networkfunction.NewAzureTrafficCollector(ctx, "azureTrafficCollector", &networkfunction.AzureTrafficCollectorArgs{
			AzureTrafficCollectorName: pulumi.String("atc"),
			Location:                  pulumi.String("West US"),
			ResourceGroupName:         pulumi.String("rg1"),
			Tags: pulumi.StringMap{
				"key1": pulumi.String("value1"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
