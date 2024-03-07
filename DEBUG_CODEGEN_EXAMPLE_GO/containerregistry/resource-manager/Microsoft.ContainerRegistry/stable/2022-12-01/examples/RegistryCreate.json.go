package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/containerregistry/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := containerregistry.NewRegistry(ctx, "registry", &containerregistry.RegistryArgs{
			AdminUserEnabled:  pulumi.Bool(true),
			Location:          pulumi.String("westus"),
			RegistryName:      pulumi.String("myRegistry"),
			ResourceGroupName: pulumi.String("myResourceGroup"),
			Sku: &containerregistry.SkuArgs{
				Name: pulumi.String("Standard"),
			},
			Tags: pulumi.StringMap{
				"key": pulumi.String("value"),
			},
		})
		if err != nil {
			return err
		}
		return nil
	})
}
