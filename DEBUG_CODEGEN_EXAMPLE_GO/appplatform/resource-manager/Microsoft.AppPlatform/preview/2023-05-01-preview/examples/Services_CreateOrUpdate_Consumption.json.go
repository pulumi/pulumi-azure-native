package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/appplatform/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := appplatform.NewService(ctx, "service", &appplatform.ServiceArgs{
			Location: pulumi.String("eastus"),
			Properties: &appplatform.ClusterResourcePropertiesArgs{
				ManagedEnvironmentId: pulumi.String("/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.App/managedEnvironments/myenvironment"),
			},
			ResourceGroupName: pulumi.String("myResourceGroup"),
			ServiceName:       pulumi.String("myservice"),
			Sku: &appplatform.SkuArgs{
				Name: pulumi.String("S0"),
				Tier: pulumi.String("StandardGen2"),
			},
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
