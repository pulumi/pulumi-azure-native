package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewDataContainer(ctx, "dataContainer", &machinelearningservices.DataContainerArgs{
			DataContainerProperties: &machinelearningservices.DataContainerTypeArgs{
				DataType:    pulumi.String("UriFile"),
				Description: pulumi.String("string"),
				Properties: pulumi.StringMap{
					"properties1": pulumi.String("value1"),
					"properties2": pulumi.String("value2"),
				},
				Tags: pulumi.StringMap{
					"tag1": pulumi.String("value1"),
					"tag2": pulumi.String("value2"),
				},
			},
			Name:              pulumi.String("datacontainer123"),
			ResourceGroupName: pulumi.String("testrg123"),
			WorkspaceName:     pulumi.String("workspace123"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
