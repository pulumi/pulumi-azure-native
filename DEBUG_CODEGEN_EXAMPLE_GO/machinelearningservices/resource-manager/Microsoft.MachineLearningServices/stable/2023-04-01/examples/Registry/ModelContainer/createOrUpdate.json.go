package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryModelContainer(ctx, "registryModelContainer", &machinelearningservices.RegistryModelContainerArgs{
			ModelContainerProperties: &machinelearningservices.ModelContainerTypeArgs{
				Description: pulumi.String("Model container description"),
				Tags: pulumi.StringMap{
					"tag1": pulumi.String("value1"),
					"tag2": pulumi.String("value2"),
				},
			},
			ModelName:         pulumi.String("testContainer"),
			RegistryName:      pulumi.String("registry123"),
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
