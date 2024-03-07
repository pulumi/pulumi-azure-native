package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewDataVersion(ctx, "dataVersion", &machinelearningservices.DataVersionArgs{
			DataVersionBaseProperties: machinelearningservices.UriFileDataVersion{
				DataType:    "uri_file",
				DataUri:     "string",
				Description: "string",
				IsAnonymous: false,
				Properties: map[string]interface{}{
					"string": "string",
				},
				Tags: map[string]interface{}{
					"string": "string",
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
