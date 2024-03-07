package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewRegistryDataVersion(ctx, "registryDataVersion", &machinelearningservices.RegistryDataVersionArgs{
			DataVersionBaseProperties: machinelearningservices.MLTableData{
				DataType:    "mltable",
				DataUri:     "string",
				Description: "string",
				IsAnonymous: false,
				IsArchived:  false,
				Properties: map[string]interface{}{
					"string": "string",
				},
				ReferencedUris: []string{
					"string",
				},
				Tags: map[string]interface{}{
					"string": "string",
				},
			},
			Name:              pulumi.String("string"),
			RegistryName:      pulumi.String("registryName"),
			ResourceGroupName: pulumi.String("test-rg"),
			Version:           pulumi.String("string"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
