package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryDataContainer(ctx, "registryDataContainer", &machinelearningservices.RegistryDataContainerArgs{
			DataContainerProperties: &machinelearningservices.DataContainerTypeArgs{
				DataType:    pulumi.String("uri_folder"),
				Description: pulumi.String("string"),
				IsArchived:  pulumi.Bool(false),
				Properties: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
				Tags: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
			},
			Name:              pulumi.String("string"),
			RegistryName:      pulumi.String("registryName"),
			ResourceGroupName: pulumi.String("test-rg"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
