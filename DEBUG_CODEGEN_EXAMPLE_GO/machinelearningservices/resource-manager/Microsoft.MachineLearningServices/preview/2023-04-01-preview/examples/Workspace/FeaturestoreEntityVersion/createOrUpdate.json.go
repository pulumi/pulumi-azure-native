package main

import (
	"github.com/pulumi/pulumi-azure-native-sdk/machinelearningservices/v2"
	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func main() {
	pulumi.Run(func(ctx *pulumi.Context) error {
		_, err := machinelearningservices.NewFeaturestoreEntityVersion(ctx, "featurestoreEntityVersion", &machinelearningservices.FeaturestoreEntityVersionArgs{
			FeaturestoreEntityVersionProperties: &machinelearningservices.FeaturestoreEntityVersionTypeArgs{
				Description: pulumi.String("string"),
				IndexColumns: machinelearningservices.IndexColumnArray{
					&machinelearningservices.IndexColumnArgs{
						ColumnName: pulumi.String("string"),
						DataType:   pulumi.String("Datetime"),
					},
				},
				IsAnonymous: pulumi.Bool(false),
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
			Version:           pulumi.String("string"),
			WorkspaceName:     pulumi.String("my-aml-workspace"),
		})
		if err != nil {
			return err
		}
		return nil
	})
}
