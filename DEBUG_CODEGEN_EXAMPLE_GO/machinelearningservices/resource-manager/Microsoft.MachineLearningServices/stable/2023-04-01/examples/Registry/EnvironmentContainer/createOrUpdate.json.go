package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryEnvironmentContainer(ctx, "registryEnvironmentContainer", &machinelearningservices.RegistryEnvironmentContainerArgs{
			EnvironmentContainerProperties: &machinelearningservices.EnvironmentContainerTypeArgs{
				Description: pulumi.String("string"),
				Properties: pulumi.StringMap{
					"additionalProp1": pulumi.String("string"),
					"additionalProp2": pulumi.String("string"),
					"additionalProp3": pulumi.String("string"),
				},
				Tags: pulumi.StringMap{
					"additionalProp1": pulumi.String("string"),
					"additionalProp2": pulumi.String("string"),
					"additionalProp3": pulumi.String("string"),
				},
			},
			EnvironmentName:   pulumi.String("testEnvironment"),
			RegistryName:      pulumi.String("testregistry"),
			ResourceGroupName: pulumi.String("testrg123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
