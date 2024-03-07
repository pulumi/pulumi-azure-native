package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewEnvironmentContainer(ctx, "environmentContainer", &machinelearningservices.EnvironmentContainerArgs{
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
			Name:              pulumi.String("testEnvironment"),
			ResourceGroupName: pulumi.String("testrg123"),
			WorkspaceName:     pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
