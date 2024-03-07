package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewEnvironmentSpecificationVersion(ctx, "environmentSpecificationVersion", &machinelearningservices.EnvironmentSpecificationVersionArgs{
			Name: pulumi.String("testEnvironment"),
			Properties: machinelearningservices.EnvironmentSpecificationVersionResponse{
				CondaFile:   pulumi.String("channels:\n- defaults\ndependencies:\n- python=3.7.7\nname: my-env"),
				Description: pulumi.String("string"),
				Docker: machinelearningservices.DockerBuild{
					DockerSpecificationType: "Build",
					Dockerfile:              "FROM myimage",
				},
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
			ResourceGroupName: pulumi.String("testrg123"),
			Version:           pulumi.String("1"),
			WorkspaceName:     pulumi.String("testworkspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
