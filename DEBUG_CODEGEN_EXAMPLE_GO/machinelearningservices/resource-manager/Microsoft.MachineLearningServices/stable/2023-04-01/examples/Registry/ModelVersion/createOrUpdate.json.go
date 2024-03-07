package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryModelVersion(ctx, "registryModelVersion", &machinelearningservices.RegistryModelVersionArgs{
			ModelName: pulumi.String("string"),
			ModelVersionProperties: &machinelearningservices.ModelVersionTypeArgs{
				Description: pulumi.String("string"),
				Flavors: machinelearningservices.FlavorDataMap{
					"string": &machinelearningservices.FlavorDataArgs{
						Data: pulumi.StringMap{
							"string": pulumi.String("string"),
						},
					},
				},
				IsAnonymous: pulumi.Bool(false),
				ModelType:   pulumi.String("CustomModel"),
				ModelUri:    pulumi.String("string"),
				Properties: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
				Tags: pulumi.StringMap{
					"string": pulumi.String("string"),
				},
			},
			RegistryName:      pulumi.String("my-aml-registry"),
			ResourceGroupName: pulumi.String("test-rg"),
			Version:           pulumi.String("string"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
