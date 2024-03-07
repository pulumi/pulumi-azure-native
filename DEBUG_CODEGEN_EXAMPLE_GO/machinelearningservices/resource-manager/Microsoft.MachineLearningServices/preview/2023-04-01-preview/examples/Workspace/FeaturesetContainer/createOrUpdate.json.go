package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewFeaturesetContainerEntity(ctx, "featuresetContainerEntity", &machinelearningservices.FeaturesetContainerEntityArgs{
			FeaturesetContainerProperties: &machinelearningservices.FeaturesetContainerArgs{
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
			ResourceGroupName: pulumi.String("test-rg"),
			WorkspaceName:     pulumi.String("my-aml-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
