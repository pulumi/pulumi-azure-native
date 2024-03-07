package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewModelVersion(ctx, "modelVersion", &machinelearningservices.ModelVersionArgs{
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
			Name:              pulumi.String("string"),
			ResourceGroupName: pulumi.String("test-rg"),
			Version:           pulumi.String("string"),
			WorkspaceName:     pulumi.String("my-aml-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
