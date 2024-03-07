package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryComponentContainer(ctx, "registryComponentContainer", &machinelearningservices.RegistryComponentContainerArgs{
			ComponentContainerProperties: &machinelearningservices.ComponentContainerTypeArgs{
				Description: pulumi.String("string"),
				Properties: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
				Tags: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
			},
			ComponentName:     pulumi.String("string"),
			RegistryName:      pulumi.String("my-aml-registry"),
			ResourceGroupName: pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
