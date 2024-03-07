package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryCodeContainer(ctx, "registryCodeContainer", &machinelearningservices.RegistryCodeContainerArgs{
			CodeContainerProperties: &machinelearningservices.CodeContainerTypeArgs{
				Description: pulumi.String("string"),
				Tags: pulumi.StringMap{
					"tag1": pulumi.String("value1"),
					"tag2": pulumi.String("value2"),
				},
			},
			CodeName:          pulumi.String("testContainer"),
			RegistryName:      pulumi.String("testregistry"),
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
